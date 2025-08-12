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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"_rollupManager\",\"type\":\"address\"},{\"internalType\":\"contractIAggLayerGateway\",\"name\":\"_aggLayerGateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AggchainManagerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainSignersHashNotInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyAggchainSignersArray\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalAccInputHashDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndicesNotInDescendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggLayerGatewayAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggchainDataLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggchainType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAggchainManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAggchainManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingVKeyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyVKeyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerURLCannotBeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdTooHighAfterRemoval\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultGatewayAlreadyDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultGatewayAlreadyEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroValueAggchainVKey\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAggchainManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"AcceptAggchainManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldVKeyManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVKeyManager\",\"type\":\"address\"}],\"name\":\"AcceptVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"AddAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainSignersHash\",\"type\":\"bytes32\"}],\"name\":\"AggchainSignersHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DisableUseDefaultGatewayFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EnableUseDefaultGatewayFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"OnVerifyPessimisticECDSAMultisig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"aggchainSigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newThreshold\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainSignersHash\",\"type\":\"bytes32\"}],\"name\":\"SignersAndThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentAggchainManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAggchainManager\",\"type\":\"address\"}],\"name\":\"TransferAggchainManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentVKeyManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingVKeyManager\",\"type\":\"address\"}],\"name\":\"TransferVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousAggchainVKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"UpdateAggchainVKey\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AGGCHAIN_ECDSA_MULTISIG_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AGGCHAIN_TYPE\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONSENSUS_TYPE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAggchainManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptVKeyManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"addOwnedAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggLayerGateway\",\"outputs\":[{\"internalType\":\"contractIAggLayerGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggchainManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"aggchainSigners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggchainSignersHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableUseDefaultGatewayFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableUseDefaultGatewayFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"getAggchainHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"getAggchainParamsAndVKeySelector\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggchainSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggchainSignersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"getAggchainTypeFromSelector\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"getAggchainVKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"aggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes2\",\"name\":\"aggchainVKeyVersion\",\"type\":\"bytes2\"},{\"internalType\":\"bytes2\",\"name\":\"aggchainType\",\"type\":\"bytes2\"}],\"name\":\"getAggchainVKeySelector\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"getAggchainVKeyVersionFromSelector\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"initAggchainManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initializeBytesAggchain\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"onVerifyPessimistic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"ownedAggchainVKeys\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"ownedAggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAggchainManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingVKeyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signerToURLs\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"transferAggchainManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVKeyManager\",\"type\":\"address\"}],\"name\":\"transferVKeyManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"updatedAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"updateOwnedAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structAggchainBase.RemoveSignerInfo[]\",\"name\":\"_signersToRemove\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structAggchainBase.SignerInfo[]\",\"name\":\"_signersToAdd\",\"type\":\"tuple[]\"},{\"internalType\":\"uint32\",\"name\":\"_newThreshold\",\"type\":\"uint32\"}],\"name\":\"updateSignersAndThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useDefaultGateway\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vKeyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x610120604052348015610010575f5ffd5b5060405161392a38038061392a83398101604081905261002f916101c4565b6001600160a01b0380861660a05280851660805280841660c052821660e05284848484848484848461005f6100f0565b505050506001600160a01b038116158061008057506001600160a01b038516155b8061009257506001600160a01b038416155b806100a457506001600160a01b038316155b806100b657506001600160a01b038216155b156100d45760405163f6b2911f60e01b815260040160405180910390fd5b6001600160a01b03166101005250610235975050505050505050565b5f54610100900460ff161561015b5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff90811610156101ab575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146101c1575f5ffd5b50565b5f5f5f5f5f60a086880312156101d8575f5ffd5b85516101e3816101ad565b60208701519095506101f4816101ad565b6040870151909450610205816101ad565b6060870151909350610216816101ad565b6080870151909250610227816101ad565b809150509295509295909350565b60805160a05160c05160e051610100516136a36102875f395f818161086b0152610b4301525f818161064d01528181611cfd0152611e9b01525f61084401525f61095601525f6109a501526136a35ff3fe608060405234801561000f575f5ffd5b506004361061037d575f3560e01c80636ff512cc116101d4578063c754c7ed11610109578063e279984e116100a9578063e90a340911610079578063e90a3409146109e3578063effb8479146109f4578063f851a44014610a13578063ff90407914610a38575f5ffd5b8063e279984e14610980578063e46761c4146109a0578063e631476c146109c7578063e7a7ed02146109cf575f5ffd5b8063cea5a4c0116100e4578063cea5a4c014610929578063cfa8ed4714610931578063d02103ca14610951578063dc8c424914610978575f5ffd5b8063c754c7ed146108e6578063c89e42df1461090e578063ca69e7dc14610921575f5ffd5b80639ee4afa311610174578063ada8f9191161014f578063ada8f9191461088d578063b3a326f7146108a0578063bdfbed7e146108b3578063bfb193b6146108c6575f5ffd5b80639ee4afa31461082c578063a3c573eb1461083f578063ab0475cf14610866575f5ffd5b8063759664d4116101af578063759664d4146107b25780637df73e27146107ee57806385018182146108115780638c3d730114610824575f5ffd5b80636ff512cc1461076c578063712570221461077f5780637388c43614610792575f5ffd5b80633cbc795b116102b5578063542028d5116102555780636a55f66c116102255780636a55f66c1461070b5780636b8616ce1461071e5780636e05d2cd1461073d5780636e7fbce914610746575f5ffd5b8063542028d51461068f578063546350c41461069757806354fd4d50146106aa578063680196ed146106e3575f5ffd5b8063439fab9111610290578063439fab91146105fc578063456052671461060f57806349b7b80214610648578063527570f11461066f575f5ffd5b80633cbc795b1461059a5780633e1e0121146105d757806342cde4e8146105ec575f5ffd5b806326f9b76d1161032057806335acd6c2116102fb57806335acd6c21461054c578063368c822c1461055f57806336cd6b5b146105675780633c351e101461057a575f5ffd5b806326f9b76d146104b25780632c111c0614610519578063314eb17b14610539575f5ffd5b806317b7a9f01161035b57806317b7a9f0146103c657806319451a8f146103cf5780631d0b435e146103e2578063267822471461046d575f5ffd5b806301fcf6a014610381578063107bf28c146103a757806315981b29146103bc575b5f5ffd5b61039461038f366004612b44565b610a5d565b6040519081526020015b60405180910390f35b6103af610bc7565b60405161039e9190612b66565b6103c4610c53565b005b61039460455481565b6103c46103dd366004612bb9565b610d31565b61043c6103f0366004612c12565b60101c7dffff00000000000000000000000000000000000000000000000000000000167fffff000000000000000000000000000000000000000000000000000000000000919091161790565b6040517fffffffff00000000000000000000000000000000000000000000000000000000909116815260200161039e565b60015461048d9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161039e565b6104e86104c0366004612b44565b60101b7fffff0000000000000000000000000000000000000000000000000000000000001690565b6040517fffff000000000000000000000000000000000000000000000000000000000000909116815260200161039e565b60085461048d9073ffffffffffffffffffffffffffffffffffffffff1681565b6103c4610547366004612bb9565b610e92565b61048d61055a366004612c43565b610fc8565b6103c4610ffd565b6103af610575366004612c7b565b6110d9565b60095461048d9073ffffffffffffffffffffffffffffffffffffffff1681565b6009546105c29074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff909116815260200161039e565b6105df6110f1565b60405161039e9190612c96565b6044546105c29063ffffffff1681565b6103c461060a366004612dea565b61115e565b60075461062f9068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff909116815260200161039e565b61048d7f000000000000000000000000000000000000000000000000000000000000000081565b60405461048d9073ffffffffffffffffffffffffffffffffffffffff1681565b6103af6115c6565b6103c46106a5366004612e92565b6115d3565b60408051808201909152600681527f76312e302e30000000000000000000000000000000000000000000000000000060208201526103af565b6106f66106f1366004612dea565b611896565b6040805192835260208301919091520161039e565b610394610719366004612dea565b611973565b61039461072c366004612f41565b60066020525f908152604090205481565b61039460055481565b6104e87e0200000000000000000000000000000000000000000000000000000000000081565b6103c461077a366004612c7b565b611a1f565b6103c461078d366004612f86565b611ae8565b603f5461048d9073ffffffffffffffffffffffffffffffffffffffff1681565b6103af6040518060400160405280600681526020017f76312e302e30000000000000000000000000000000000000000000000000000081525081565b6108016107fc366004612c7b565b611b1a565b604051901515815260200161039e565b6103c461081f366004612c7b565b611b56565b6103c4611c28565b6103c461083a366004613030565b611cfb565b61048d7f000000000000000000000000000000000000000000000000000000000000000081565b61048d7f000000000000000000000000000000000000000000000000000000000000000081565b6103c461089b366004612c7b565b611dd0565b6103c46108ae366004612c7b565b611e99565b6103c46108c1366004612c7b565b611fcf565b603e5461048d9073ffffffffffffffffffffffffffffffffffffffff1681565b60075461062f90700100000000000000000000000000000000900467ffffffffffffffff1681565b6103c461091c36600461309e565b6120e9565b604254610394565b6105c2600181565b60025461048d9073ffffffffffffffffffffffffffffffffffffffff1681565b61048d7f000000000000000000000000000000000000000000000000000000000000000081565b6103c461217b565b603d5461048d9073ffffffffffffffffffffffffffffffffffffffff1681565b61048d7f000000000000000000000000000000000000000000000000000000000000000081565b6103c4612272565b60075461062f9067ffffffffffffffff1681565b6104e86109f1366004612b44565b90565b610394610a02366004612b44565b60416020525f908152604090205481565b5f5461048d9062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b603e546108019074010000000000000000000000000000000000000000900460ff1681565b603e545f9074010000000000000000000000000000000000000000900460ff1615158103610af357507fffffffff0000000000000000000000000000000000000000000000000000000081165f9081526041602052604090205480610aee576040517f925e5a3a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b919050565b6040517f6cabfdab0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000831660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690636cabfdab90602401602060405180830381865afa158015610b9d573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bc191906130d0565b92915050565b60048054610bd4906130e7565b80601f0160208091040260200160405190810160405280929190818152602001828054610c00906130e7565b8015610c4b5780601f10610c2257610100808354040283529160200191610c4b565b820191905f5260205f20905b815481529060010190602001808311610c2e57829003601f168201915b505050505081565b60405473ffffffffffffffffffffffffffffffffffffffff163314610ca4576040517f3ac87ac900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603f80546040805473ffffffffffffffffffffffffffffffffffffffff8082167fffffffffffffffffffffffff000000000000000000000000000000000000000080861682179096559490911682558151921680835260208301939093527f67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f91015b60405180910390a150565b603d5473ffffffffffffffffffffffffffffffffffffffff163314610d82576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80610db9576040517fe1dbcf2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f9081526041602052604090205415610e21576040517fe3cc761000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f81815260416020908152604091829020849055815192835282018390527f6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a7991015b60405180910390a15050565b603d5473ffffffffffffffffffffffffffffffffffffffff163314610ee3576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f90815260416020526040902054610f4a576040517ff360deaf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f818152604160209081526040918290208054908590558251938452908301819052908201839052907f0aa5f73c189fb0b0a7cc98ae5fa89dfc16595480396208483518178435ed5b4f9060600160405180910390a1505050565b60428181548110610fd7575f80fd5b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b603e5473ffffffffffffffffffffffffffffffffffffffff16331461104e576040517f05882cf000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d8054603e805473ffffffffffffffffffffffffffffffffffffffff8082167fffffffffffffffffffffffff0000000000000000000000000000000000000000808616821790965594909116909155604080519190921680825260208201939093527fbb3b066505f14a628f4ba4187a046abd4dd17e96591d7a9ed31c91c79322ffe29101610d26565b60436020525f908152604090208054610bd4906130e7565b6060604280548060200260200160405190810160405280929190818152602001828054801561115457602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611129575b5050505050905090565b603f5473ffffffffffffffffffffffffffffffffffffffff1633146111af576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f805460ff16907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00815c168217905d505f54600290610100900460ff161580156111ff57505f5460ff8083169116105b611290576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff80841691909117610100178255815c16900361140a575f5f5f5f5f5f5f5f5f8a8060200190518101906112ea919061319f565b985098509850985098509850985098509850881561136d577fffffffff00000000000000000000000000000000000000000000000000000000871615158061133157508715155b15611368576040517fadc06ae700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113eb565b7fffff000000000000000000000000000000000000000000000000000000000000601088901b167e02000000000000000000000000000000000000000000000000000000000000146113eb576040517f4570795000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113fc85858585858e8e8e8e612381565b50505050505050505061156c565b60ff5f5c1660010361153a575f5f5f5f8580602001905181019061142e9190613271565b935093509350935083156114a7577fffffffff00000000000000000000000000000000000000000000000000000000821615158061146b57508215155b156114a2576040517fadc06ae700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611525565b7fffff000000000000000000000000000000000000000000000000000000000000601083901b167e0200000000000000000000000000000000000000000000000000000000000014611525576040517f4570795000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611531848484846124c8565b5050505061156c565b6040517fadc06ae700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001610e86565b60038054610bd4906130e7565b603f5473ffffffffffffffffffffffffffffffffffffffff163314611624576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018411156116c5575f5b61163a6001866132ee565b8110156116c357858561164e836001613301565b81811061165d5761165d613314565b9050604002016020013586868381811061167957611679613314565b9050604002016020013510156116bb576040517fb9a11d3100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60010161162f565b505b5f5b848110156117235761171b8686838181106116e4576116e4613314565b6116fa9260206040909202019081019150612c7b565b87878481811061170c5761170c613314565b9050604002016020013561261e565b6001016116c7565b505f5b828110156117d6576117ce84848381811061174357611743613314565b90506020028101906117559190613341565b611763906020810190612c7b565b85858481811061177557611775613314565b90506020028101906117879190613341565b61179590602081019061337d565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061281192505050565b600101611726565b5060425463ffffffff82161115611819576040517faabd5a0900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604480547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff8316179055611851612962565b7fa9031663ce9b6bcb10b14a169f269c3a2ab3685eba97d673e3af691473bf59f9604282604554604051611887939291906133de565b60405180910390a15050505050565b5f5f82516020146118d3576040517f3063965400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f838060200190518101906118e8919061344b565b90507fffff000000000000000000000000000000000000000000000000000000000000601082901b167e0200000000000000000000000000000000000000000000000000000000000014611968576040517f4570795000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505f93849350915050565b6045545f906119ae576040517fdd41f1ef00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f6119b984611896565b6045546040517c010000000000000000000000000000000000000000000000000000000060208201526024810184905260448101839052606481019190915291935091506084016040516020818303038152906040528051906020012092505050919050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314611a75576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc090602001610d26565b6040517ff57ac68300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81165f9081526043602052604081208054829190611b4c906130e7565b9050119050919050565b603d5473ffffffffffffffffffffffffffffffffffffffff163314611ba7576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217909255603d546040805191909316815260208101919091527fc54ae01017d0b80bd8af833f66387d6eb547dc16c8206faf13d0b72764aab8b29101610d26565b60015473ffffffffffffffffffffffffffffffffffffffff163314611c79576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001545f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff1673ffffffffffffffffffffffffffffffffffffffff9092166201000081029290921790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e906020015b60405180910390a1565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314611d6a576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208114611da4576040517f3063965400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517fc118263690a4306c74bd1bc80b55962addc2d9e61619ac0b2c2883badbbd01d8905f90a15050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314611e26576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001610d26565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314611f08576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116611f55576040517fd6bdac3f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8316908117909155604080515f815260208101929092527f67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f9101610d26565b603f5473ffffffffffffffffffffffffffffffffffffffff163314612020576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff811661206d576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169182178355603f5483519116815260208101919091527fa3d8e5d045432398be30f83ce7c35a7bfc220c1b66cc5bf3f4dd4d539d93fab69101610d26565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff16331461213f576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600361214b82826134b1565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b2081604051610d269190612b66565b603d5473ffffffffffffffffffffffffffffffffffffffff1633146121cc576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e5474010000000000000000000000000000000000000000900460ff16612220576040517f62de044500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690556040517f67dd1717a1952df380cb73eecb312e949df6d6a086bd7f88669005341972528e905f90a1565b603d5473ffffffffffffffffffffffffffffffffffffffff1633146122c3576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e5474010000000000000000000000000000000000000000900460ff1615612318576040517f93be805100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790556040517fb6563aed80fde357e737eb0d19f246a58cb6bfd469933d05701ecbad0f2dca84905f90a1565b5f54610100900460ff16612417576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401611287565b73ffffffffffffffffffffffffffffffffffffffff8916158061244e575073ffffffffffffffffffffffffffffffffffffffff8816155b8061246d575073ffffffffffffffffffffffffffffffffffffffff8116155b156124a4576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6124b189898989896129e4565b6124bd848484846124c8565b505050505050505050565b5f54610100900460ff1661255e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401611287565b603e80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000095151595909502949094179093557fffffffff00000000000000000000000000000000000000000000000000000000165f90815260416020526040902055603d80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909216919091179055565b6042548110612659576040517fd244b30700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff166042828154811061268357612683613314565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff16146126db576040517fd244b30700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82165f90815260436020526040812061270891612ac5565b60428054612718906001906132ee565b8154811061272857612728613314565b5f918252602090912001546042805473ffffffffffffffffffffffffffffffffffffffff909216918390811061276057612760613314565b905f5260205f20015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060428054806127b6576127b66135c8565b5f8281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690550190555050565b73ffffffffffffffffffffffffffffffffffffffff821661285e576040517f7b3a0df600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80515f03612898576040517f8715f5fb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6128a182611b1a565b156128d8576040517f38615ecc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60428054600181019091557f38dfe4635b27babeca8be38d3b448cb5161a639b899a14825ba9c8d7892eb8c30180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84169081179091555f90815260436020526040902061295d82826134b1565b505050565b60445460405161297e9163ffffffff16906042906020016135f5565b604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018152908290528051602091820120604581905582527f43a1dd43d2705e069faf1da3afc2772475650d6642c2d71bd620e0fd2f7c3ee89101611cf1565b5f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff166201000073ffffffffffffffffffffffffffffffffffffffff8881169190910291909117909155600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000169186169190911790556003612a6b83826134b1565b506004612a7882826134b1565b5050600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9390931692909217909155505050565b508054612ad1906130e7565b5f825580601f10612ae0575050565b601f0160209004905f5260205f2090810190612afc9190612aff565b50565b5b80821115612b13575f8155600101612b00565b5090565b7fffffffff0000000000000000000000000000000000000000000000000000000081168114612afc575f5ffd5b5f60208284031215612b54575f5ffd5b8135612b5f81612b17565b9392505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b5f5f60408385031215612bca575f5ffd5b8235612bd581612b17565b946020939093013593505050565b80357fffff00000000000000000000000000000000000000000000000000000000000081168114610aee575f5ffd5b5f5f60408385031215612c23575f5ffd5b612c2c83612be3565b9150612c3a60208401612be3565b90509250929050565b5f60208284031215612c53575f5ffd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff81168114612afc575f5ffd5b5f60208284031215612c8b575f5ffd5b8135612b5f81612c5a565b602080825282518282018190525f918401906040840190835b81811015612ce357835173ffffffffffffffffffffffffffffffffffffffff16835260209384019390920191600101612caf565b509095945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715612d6257612d62612cee565b604052919050565b5f67ffffffffffffffff821115612d8357612d83612cee565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b5f612dc1612dbc84612d6a565b612d1b565b9050828152838383011115612dd4575f5ffd5b828260208301375f602084830101529392505050565b5f60208284031215612dfa575f5ffd5b813567ffffffffffffffff811115612e10575f5ffd5b8201601f81018413612e20575f5ffd5b612e2f84823560208401612daf565b949350505050565b5f5f83601f840112612e47575f5ffd5b50813567ffffffffffffffff811115612e5e575f5ffd5b6020830191508360208260051b8501011115612e78575f5ffd5b9250929050565b803563ffffffff81168114610aee575f5ffd5b5f5f5f5f5f60608688031215612ea6575f5ffd5b853567ffffffffffffffff811115612ebc575f5ffd5b8601601f81018813612ecc575f5ffd5b803567ffffffffffffffff811115612ee2575f5ffd5b8860208260061b8401011115612ef6575f5ffd5b60209182019650945086013567ffffffffffffffff811115612f16575f5ffd5b612f2288828901612e37565b9094509250612f35905060408701612e7f565b90509295509295909350565b5f60208284031215612f51575f5ffd5b813567ffffffffffffffff81168114612b5f575f5ffd5b5f82601f830112612f77575f5ffd5b612b5f83833560208501612daf565b5f5f5f5f5f5f60c08789031215612f9b575f5ffd5b8635612fa681612c5a565b95506020870135612fb681612c5a565b9450612fc460408801612e7f565b93506060870135612fd481612c5a565b9250608087013567ffffffffffffffff811115612fef575f5ffd5b612ffb89828a01612f68565b92505060a087013567ffffffffffffffff811115613017575f5ffd5b61302389828a01612f68565b9150509295509295509295565b5f5f60208385031215613041575f5ffd5b823567ffffffffffffffff811115613057575f5ffd5b8301601f81018513613067575f5ffd5b803567ffffffffffffffff81111561307d575f5ffd5b85602082840101111561308e575f5ffd5b6020919091019590945092505050565b5f602082840312156130ae575f5ffd5b813567ffffffffffffffff8111156130c4575f5ffd5b612e2f84828501612f68565b5f602082840312156130e0575f5ffd5b5051919050565b600181811c908216806130fb57607f821691505b602082108103613132577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b80518015158114610aee575f5ffd5b8051610aee81612c5a565b5f82601f830112613161575f5ffd5b815161316f612dbc82612d6a565b818152846020838601011115613183575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f5f5f5f5f5f5f5f5f6101208a8c0312156131b8575f5ffd5b6131c18a613138565b60208b015160408c0151919a5098506131d981612b17565b96506131e760608b01613147565b95506131f560808b01613147565b945061320360a08b01613147565b935061321160c08b01613147565b925060e08a015167ffffffffffffffff81111561322c575f5ffd5b6132388c828d01613152565b9250506101008a015167ffffffffffffffff811115613255575f5ffd5b6132618c828d01613152565b9150509295985092959850929598565b5f5f5f5f60808587031215613284575f5ffd5b61328d85613138565b6020860151604087015191955093506132a581612b17565b60608601519092506132b681612c5a565b939692955090935050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b81810381811115610bc157610bc16132c1565b80820180821115610bc157610bc16132c1565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f82357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc1833603018112613373575f5ffd5b9190910192915050565b5f5f83357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126133b0575f5ffd5b83018035915067ffffffffffffffff8211156133ca575f5ffd5b602001915036819003821315612e78575f5ffd5b606080825284549082018190525f8581526020812090916080840190835b8181101561343057835473ffffffffffffffffffffffffffffffffffffffff168352600193840193602090930192016133fc565b505063ffffffff959095166020840152505060400152919050565b5f6020828403121561345b575f5ffd5b8151612b5f81612b17565b601f82111561295d57805f5260205f20601f840160051c8101602085101561348b5750805b601f840160051c820191505b818110156134aa575f8155600101613497565b5050505050565b815167ffffffffffffffff8111156134cb576134cb612cee565b6134df816134d984546130e7565b84613466565b6020601f821160018114613530575f83156134fa5750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b1784556134aa565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b8281101561357d578785015182556020948501946001909201910161355d565b50848210156135b957868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7fffffffff000000000000000000000000000000000000000000000000000000008360e01b1681525f600482018354845f5260205f205f5b8281101561366157815473ffffffffffffffffffffffffffffffffffffffff1684526020909301926001918201910161362d565b5091969550505050505056fea264697066735822122043cb80eabe1e2c40b22bf0c31e73775c998d860048947f946c62c725e79e655464736f6c634300081c0033",
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

// PendingVKeyManager is a free data retrieval call binding the contract method 0xbfb193b6.
//
// Solidity: function pendingVKeyManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) PendingVKeyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "pendingVKeyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingVKeyManager is a free data retrieval call binding the contract method 0xbfb193b6.
//
// Solidity: function pendingVKeyManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) PendingVKeyManager() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.PendingVKeyManager(&_Aggchainecdsamultisig.CallOpts)
}

// PendingVKeyManager is a free data retrieval call binding the contract method 0xbfb193b6.
//
// Solidity: function pendingVKeyManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) PendingVKeyManager() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.PendingVKeyManager(&_Aggchainecdsamultisig.CallOpts)
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
// Solidity: function threshold() view returns(uint32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) Threshold(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) Threshold() (uint32, error) {
	return _Aggchainecdsamultisig.Contract.Threshold(&_Aggchainecdsamultisig.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) Threshold() (uint32, error) {
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

// UseDefaultGateway is a free data retrieval call binding the contract method 0xff904079.
//
// Solidity: function useDefaultGateway() view returns(bool)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) UseDefaultGateway(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "useDefaultGateway")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseDefaultGateway is a free data retrieval call binding the contract method 0xff904079.
//
// Solidity: function useDefaultGateway() view returns(bool)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) UseDefaultGateway() (bool, error) {
	return _Aggchainecdsamultisig.Contract.UseDefaultGateway(&_Aggchainecdsamultisig.CallOpts)
}

// UseDefaultGateway is a free data retrieval call binding the contract method 0xff904079.
//
// Solidity: function useDefaultGateway() view returns(bool)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) UseDefaultGateway() (bool, error) {
	return _Aggchainecdsamultisig.Contract.UseDefaultGateway(&_Aggchainecdsamultisig.CallOpts)
}

// VKeyManager is a free data retrieval call binding the contract method 0xe279984e.
//
// Solidity: function vKeyManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) VKeyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "vKeyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VKeyManager is a free data retrieval call binding the contract method 0xe279984e.
//
// Solidity: function vKeyManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) VKeyManager() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.VKeyManager(&_Aggchainecdsamultisig.CallOpts)
}

// VKeyManager is a free data retrieval call binding the contract method 0xe279984e.
//
// Solidity: function vKeyManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) VKeyManager() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.VKeyManager(&_Aggchainecdsamultisig.CallOpts)
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

// AcceptVKeyManagerRole is a paid mutator transaction binding the contract method 0x368c822c.
//
// Solidity: function acceptVKeyManagerRole() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) AcceptVKeyManagerRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "acceptVKeyManagerRole")
}

// AcceptVKeyManagerRole is a paid mutator transaction binding the contract method 0x368c822c.
//
// Solidity: function acceptVKeyManagerRole() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) AcceptVKeyManagerRole() (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.AcceptVKeyManagerRole(&_Aggchainecdsamultisig.TransactOpts)
}

// AcceptVKeyManagerRole is a paid mutator transaction binding the contract method 0x368c822c.
//
// Solidity: function acceptVKeyManagerRole() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) AcceptVKeyManagerRole() (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.AcceptVKeyManagerRole(&_Aggchainecdsamultisig.TransactOpts)
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

// DisableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xdc8c4249.
//
// Solidity: function disableUseDefaultGatewayFlag() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) DisableUseDefaultGatewayFlag(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "disableUseDefaultGatewayFlag")
}

// DisableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xdc8c4249.
//
// Solidity: function disableUseDefaultGatewayFlag() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) DisableUseDefaultGatewayFlag() (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.DisableUseDefaultGatewayFlag(&_Aggchainecdsamultisig.TransactOpts)
}

// DisableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xdc8c4249.
//
// Solidity: function disableUseDefaultGatewayFlag() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) DisableUseDefaultGatewayFlag() (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.DisableUseDefaultGatewayFlag(&_Aggchainecdsamultisig.TransactOpts)
}

// EnableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xe631476c.
//
// Solidity: function enableUseDefaultGatewayFlag() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) EnableUseDefaultGatewayFlag(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "enableUseDefaultGatewayFlag")
}

// EnableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xe631476c.
//
// Solidity: function enableUseDefaultGatewayFlag() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) EnableUseDefaultGatewayFlag() (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.EnableUseDefaultGatewayFlag(&_Aggchainecdsamultisig.TransactOpts)
}

// EnableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xe631476c.
//
// Solidity: function enableUseDefaultGatewayFlag() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) EnableUseDefaultGatewayFlag() (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.EnableUseDefaultGatewayFlag(&_Aggchainecdsamultisig.TransactOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializeBytesAggchain) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) Initialize(opts *bind.TransactOpts, initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "initialize", initializeBytesAggchain)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializeBytesAggchain) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) Initialize(initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.Initialize(&_Aggchainecdsamultisig.TransactOpts, initializeBytesAggchain)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializeBytesAggchain) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) Initialize(initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.Initialize(&_Aggchainecdsamultisig.TransactOpts, initializeBytesAggchain)
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

// TransferVKeyManagerRole is a paid mutator transaction binding the contract method 0x85018182.
//
// Solidity: function transferVKeyManagerRole(address newVKeyManager) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) TransferVKeyManagerRole(opts *bind.TransactOpts, newVKeyManager common.Address) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "transferVKeyManagerRole", newVKeyManager)
}

// TransferVKeyManagerRole is a paid mutator transaction binding the contract method 0x85018182.
//
// Solidity: function transferVKeyManagerRole(address newVKeyManager) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) TransferVKeyManagerRole(newVKeyManager common.Address) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.TransferVKeyManagerRole(&_Aggchainecdsamultisig.TransactOpts, newVKeyManager)
}

// TransferVKeyManagerRole is a paid mutator transaction binding the contract method 0x85018182.
//
// Solidity: function transferVKeyManagerRole(address newVKeyManager) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) TransferVKeyManagerRole(newVKeyManager common.Address) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.TransferVKeyManagerRole(&_Aggchainecdsamultisig.TransactOpts, newVKeyManager)
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

// UpdateSignersAndThreshold is a paid mutator transaction binding the contract method 0x546350c4.
//
// Solidity: function updateSignersAndThreshold((address,uint256)[] _signersToRemove, (address,string)[] _signersToAdd, uint32 _newThreshold) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) UpdateSignersAndThreshold(opts *bind.TransactOpts, _signersToRemove []AggchainBaseRemoveSignerInfo, _signersToAdd []AggchainBaseSignerInfo, _newThreshold uint32) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "updateSignersAndThreshold", _signersToRemove, _signersToAdd, _newThreshold)
}

// UpdateSignersAndThreshold is a paid mutator transaction binding the contract method 0x546350c4.
//
// Solidity: function updateSignersAndThreshold((address,uint256)[] _signersToRemove, (address,string)[] _signersToAdd, uint32 _newThreshold) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) UpdateSignersAndThreshold(_signersToRemove []AggchainBaseRemoveSignerInfo, _signersToAdd []AggchainBaseSignerInfo, _newThreshold uint32) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.UpdateSignersAndThreshold(&_Aggchainecdsamultisig.TransactOpts, _signersToRemove, _signersToAdd, _newThreshold)
}

// UpdateSignersAndThreshold is a paid mutator transaction binding the contract method 0x546350c4.
//
// Solidity: function updateSignersAndThreshold((address,uint256)[] _signersToRemove, (address,string)[] _signersToAdd, uint32 _newThreshold) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) UpdateSignersAndThreshold(_signersToRemove []AggchainBaseRemoveSignerInfo, _signersToAdd []AggchainBaseSignerInfo, _newThreshold uint32) (*types.Transaction, error) {
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

// AggchainecdsamultisigAggchainSignersHashUpdatedIterator is returned from FilterAggchainSignersHashUpdated and is used to iterate over the raw logs and unpacked data for AggchainSignersHashUpdated events raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigAggchainSignersHashUpdatedIterator struct {
	Event *AggchainecdsamultisigAggchainSignersHashUpdated // Event containing the contract specifics and raw log

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
func (it *AggchainecdsamultisigAggchainSignersHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsamultisigAggchainSignersHashUpdated)
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
		it.Event = new(AggchainecdsamultisigAggchainSignersHashUpdated)
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
func (it *AggchainecdsamultisigAggchainSignersHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsamultisigAggchainSignersHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsamultisigAggchainSignersHashUpdated represents a AggchainSignersHashUpdated event raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigAggchainSignersHashUpdated struct {
	NewAggchainSignersHash [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterAggchainSignersHashUpdated is a free log retrieval operation binding the contract event 0x43a1dd43d2705e069faf1da3afc2772475650d6642c2d71bd620e0fd2f7c3ee8.
//
// Solidity: event AggchainSignersHashUpdated(bytes32 newAggchainSignersHash)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterAggchainSignersHashUpdated(opts *bind.FilterOpts) (*AggchainecdsamultisigAggchainSignersHashUpdatedIterator, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "AggchainSignersHashUpdated")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigAggchainSignersHashUpdatedIterator{contract: _Aggchainecdsamultisig.contract, event: "AggchainSignersHashUpdated", logs: logs, sub: sub}, nil
}

// WatchAggchainSignersHashUpdated is a free log subscription operation binding the contract event 0x43a1dd43d2705e069faf1da3afc2772475650d6642c2d71bd620e0fd2f7c3ee8.
//
// Solidity: event AggchainSignersHashUpdated(bytes32 newAggchainSignersHash)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) WatchAggchainSignersHashUpdated(opts *bind.WatchOpts, sink chan<- *AggchainecdsamultisigAggchainSignersHashUpdated) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.WatchLogs(opts, "AggchainSignersHashUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsamultisigAggchainSignersHashUpdated)
				if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "AggchainSignersHashUpdated", log); err != nil {
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

// ParseAggchainSignersHashUpdated is a log parse operation binding the contract event 0x43a1dd43d2705e069faf1da3afc2772475650d6642c2d71bd620e0fd2f7c3ee8.
//
// Solidity: event AggchainSignersHashUpdated(bytes32 newAggchainSignersHash)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) ParseAggchainSignersHashUpdated(log types.Log) (*AggchainecdsamultisigAggchainSignersHashUpdated, error) {
	event := new(AggchainecdsamultisigAggchainSignersHashUpdated)
	if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "AggchainSignersHashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsamultisigDisableUseDefaultGatewayFlagIterator is returned from FilterDisableUseDefaultGatewayFlag and is used to iterate over the raw logs and unpacked data for DisableUseDefaultGatewayFlag events raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigDisableUseDefaultGatewayFlagIterator struct {
	Event *AggchainecdsamultisigDisableUseDefaultGatewayFlag // Event containing the contract specifics and raw log

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
func (it *AggchainecdsamultisigDisableUseDefaultGatewayFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsamultisigDisableUseDefaultGatewayFlag)
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
		it.Event = new(AggchainecdsamultisigDisableUseDefaultGatewayFlag)
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
func (it *AggchainecdsamultisigDisableUseDefaultGatewayFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsamultisigDisableUseDefaultGatewayFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsamultisigDisableUseDefaultGatewayFlag represents a DisableUseDefaultGatewayFlag event raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigDisableUseDefaultGatewayFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisableUseDefaultGatewayFlag is a free log retrieval operation binding the contract event 0x67dd1717a1952df380cb73eecb312e949df6d6a086bd7f88669005341972528e.
//
// Solidity: event DisableUseDefaultGatewayFlag()
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterDisableUseDefaultGatewayFlag(opts *bind.FilterOpts) (*AggchainecdsamultisigDisableUseDefaultGatewayFlagIterator, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "DisableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigDisableUseDefaultGatewayFlagIterator{contract: _Aggchainecdsamultisig.contract, event: "DisableUseDefaultGatewayFlag", logs: logs, sub: sub}, nil
}

// WatchDisableUseDefaultGatewayFlag is a free log subscription operation binding the contract event 0x67dd1717a1952df380cb73eecb312e949df6d6a086bd7f88669005341972528e.
//
// Solidity: event DisableUseDefaultGatewayFlag()
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) WatchDisableUseDefaultGatewayFlag(opts *bind.WatchOpts, sink chan<- *AggchainecdsamultisigDisableUseDefaultGatewayFlag) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.WatchLogs(opts, "DisableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsamultisigDisableUseDefaultGatewayFlag)
				if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "DisableUseDefaultGatewayFlag", log); err != nil {
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
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) ParseDisableUseDefaultGatewayFlag(log types.Log) (*AggchainecdsamultisigDisableUseDefaultGatewayFlag, error) {
	event := new(AggchainecdsamultisigDisableUseDefaultGatewayFlag)
	if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "DisableUseDefaultGatewayFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsamultisigEnableUseDefaultGatewayFlagIterator is returned from FilterEnableUseDefaultGatewayFlag and is used to iterate over the raw logs and unpacked data for EnableUseDefaultGatewayFlag events raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigEnableUseDefaultGatewayFlagIterator struct {
	Event *AggchainecdsamultisigEnableUseDefaultGatewayFlag // Event containing the contract specifics and raw log

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
func (it *AggchainecdsamultisigEnableUseDefaultGatewayFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsamultisigEnableUseDefaultGatewayFlag)
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
		it.Event = new(AggchainecdsamultisigEnableUseDefaultGatewayFlag)
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
func (it *AggchainecdsamultisigEnableUseDefaultGatewayFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsamultisigEnableUseDefaultGatewayFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsamultisigEnableUseDefaultGatewayFlag represents a EnableUseDefaultGatewayFlag event raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigEnableUseDefaultGatewayFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEnableUseDefaultGatewayFlag is a free log retrieval operation binding the contract event 0xb6563aed80fde357e737eb0d19f246a58cb6bfd469933d05701ecbad0f2dca84.
//
// Solidity: event EnableUseDefaultGatewayFlag()
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterEnableUseDefaultGatewayFlag(opts *bind.FilterOpts) (*AggchainecdsamultisigEnableUseDefaultGatewayFlagIterator, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "EnableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigEnableUseDefaultGatewayFlagIterator{contract: _Aggchainecdsamultisig.contract, event: "EnableUseDefaultGatewayFlag", logs: logs, sub: sub}, nil
}

// WatchEnableUseDefaultGatewayFlag is a free log subscription operation binding the contract event 0xb6563aed80fde357e737eb0d19f246a58cb6bfd469933d05701ecbad0f2dca84.
//
// Solidity: event EnableUseDefaultGatewayFlag()
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) WatchEnableUseDefaultGatewayFlag(opts *bind.WatchOpts, sink chan<- *AggchainecdsamultisigEnableUseDefaultGatewayFlag) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.WatchLogs(opts, "EnableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsamultisigEnableUseDefaultGatewayFlag)
				if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "EnableUseDefaultGatewayFlag", log); err != nil {
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
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) ParseEnableUseDefaultGatewayFlag(log types.Log) (*AggchainecdsamultisigEnableUseDefaultGatewayFlag, error) {
	event := new(AggchainecdsamultisigEnableUseDefaultGatewayFlag)
	if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "EnableUseDefaultGatewayFlag", log); err != nil {
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
	NewThreshold           uint32
	NewAggchainSignersHash [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSignersAndThresholdUpdated is a free log retrieval operation binding the contract event 0xa9031663ce9b6bcb10b14a169f269c3a2ab3685eba97d673e3af691473bf59f9.
//
// Solidity: event SignersAndThresholdUpdated(address[] aggchainSigners, uint32 newThreshold, bytes32 newAggchainSignersHash)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterSignersAndThresholdUpdated(opts *bind.FilterOpts) (*AggchainecdsamultisigSignersAndThresholdUpdatedIterator, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "SignersAndThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigSignersAndThresholdUpdatedIterator{contract: _Aggchainecdsamultisig.contract, event: "SignersAndThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchSignersAndThresholdUpdated is a free log subscription operation binding the contract event 0xa9031663ce9b6bcb10b14a169f269c3a2ab3685eba97d673e3af691473bf59f9.
//
// Solidity: event SignersAndThresholdUpdated(address[] aggchainSigners, uint32 newThreshold, bytes32 newAggchainSignersHash)
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

// ParseSignersAndThresholdUpdated is a log parse operation binding the contract event 0xa9031663ce9b6bcb10b14a169f269c3a2ab3685eba97d673e3af691473bf59f9.
//
// Solidity: event SignersAndThresholdUpdated(address[] aggchainSigners, uint32 newThreshold, bytes32 newAggchainSignersHash)
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

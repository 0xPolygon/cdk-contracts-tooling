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

// IAggchainSignersRemoveSignerInfo is an auto generated low-level Go binding around an user-defined struct.
type IAggchainSignersRemoveSignerInfo struct {
	Addr  common.Address
	Index *big.Int
}

// IAggchainSignersSignerInfo is an auto generated low-level Go binding around an user-defined struct.
type IAggchainSignersSignerInfo struct {
	Addr common.Address
	Url  string
}

// AggchainecdsamultisigMetaData contains all meta data concerning the Aggchainecdsamultisig contract.
var AggchainecdsamultisigMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"_rollupManager\",\"type\":\"address\"},{\"internalType\":\"contractIAggLayerGateway\",\"name\":\"_aggLayerGateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AdminCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainManagerAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainManagerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainSignersHashNotInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainSignersTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ConflictingDefaultSignersConfiguration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalAccInputHashDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FunctionNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndicesNotInDescendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggchainDataLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggchainType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitAggchainVKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MetadataArrayLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAggchainManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAggchainMetadataManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAggchainManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerURLCannotBeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultSignersAlreadyDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultSignersAlreadyEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultVkeysAlreadyDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultVkeysAlreadyEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroValueAggchainVKey\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAggchainManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"AcceptAggchainManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"AddAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"AggchainMetadataSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DisableUseDefaultSignersFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DisableUseDefaultVkeysFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EnableUseDefaultSignersFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EnableUseDefaultVkeysFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"OnVerifyPessimisticECDSAMultisig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAggchainMetadataManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAggchainMetadataManager\",\"type\":\"address\"}],\"name\":\"SetAggchainMetadataManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"aggchainSigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainMultisigHash\",\"type\":\"bytes32\"}],\"name\":\"SignersAndThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentAggchainManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAggchainManager\",\"type\":\"address\"}],\"name\":\"TransferAggchainManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousAggchainVKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"UpdateAggchainVKey\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AGGCHAIN_ECDSA_MULTISIG_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AGGCHAIN_TYPE\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONSENSUS_TYPE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_AGGCHAIN_SIGNERS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_legacypendingVKeyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_legacyvKeyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAggchainManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"addOwnedAggchainVKey\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggLayerGateway\",\"outputs\":[{\"internalType\":\"contractIAggLayerGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggchainManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"aggchainMetadata\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggchainMetadataManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggchainMultisigHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"aggchainSigners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"name\":\"batchSetAggchainMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableUseDefaultSignersFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableUseDefaultVkeysFlag\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableUseDefaultSignersFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableUseDefaultVkeysFlag\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"getAggchainHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggchainMultisigHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggchainSignerInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structIAggchainSigners.SignerInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggchainSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggchainSignersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"getAggchainTypeFromSelector\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"name\":\"getAggchainVKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"aggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes2\",\"name\":\"aggchainVKeyVersion\",\"type\":\"bytes2\"},{\"internalType\":\"bytes2\",\"name\":\"aggchainType\",\"type\":\"bytes2\"}],\"name\":\"getAggchainVKeySelector\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"getAggchainVKeyVersionFromSelector\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"getVKeyAndAggchainParams\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"initAggchainManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trustedSequencer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_trustedSequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_useDefaultSigners\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structIAggchainSigners.SignerInfo[]\",\"name\":\"_signersToAdd\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_newThreshold\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrateFromLegacyConsensus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"onVerifyPessimistic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"ownedAggchainVKeys\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"ownedAggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAggchainManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setAggchainMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAggchainMetadataManager\",\"type\":\"address\"}],\"name\":\"setAggchainMetadataManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signerToURLs\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"transferAggchainManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"updateOwnedAggchainVKey\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structIAggchainSigners.RemoveSignerInfo[]\",\"name\":\"_signersToRemove\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structIAggchainSigners.SignerInfo[]\",\"name\":\"_signersToAdd\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_newThreshold\",\"type\":\"uint256\"}],\"name\":\"updateSignersAndThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useDefaultSigners\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useDefaultVkeys\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x610120604052348015610010575f5ffd5b506040516144be3803806144be83398101604081905261002f916101c4565b6001600160a01b0380861660a05280851660805280841660c052821660e05284848484848484848461005f6100f0565b505050506001600160a01b038116158061008057506001600160a01b038516155b8061009257506001600160a01b038416155b806100a457506001600160a01b038316155b806100b657506001600160a01b038216155b156100d45760405163f6b2911f60e01b815260040160405180910390fd5b6001600160a01b03166101005250610235975050505050505050565b5f54610100900460ff161561015b5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff90811610156101ab575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146101c1575f5ffd5b50565b5f5f5f5f5f60a086880312156101d8575f5ffd5b85516101e3816101ad565b60208701519095506101f4816101ad565b6040870151909450610205816101ad565b6060870151909350610216816101ad565b6080870151909250610227816101ad565b809150509295509295909350565b60805160a05160c05160e0516101005161420d6102b15f395f8181610958015281816113740152818161166701528181611c3a0152818161248d01528181612547015261265c01525f818161072d01528181610c1f01528181611db1015261202401525f61091e01525f610a3301525f610a82015261420d5ff3fe608060405234801561000f575f5ffd5b5060043610610441575f3560e01c80636e7fbce911610243578063c754c7ed11610148578063e75235b8116100c3578063effb847911610093578063f851a44011610079578063f851a44014610b03578063fc5014d614610b28578063fd7d249314610b4d575f5ffd5b8063effb847914610ad1578063f51f563a14610af0575f5ffd5b8063e75235b814610aa4578063e7a7ed0214610aac578063e90a340914610ac0578063efe6c9f41461049d575f5ffd5b8063cea5a4c011610118578063d02103ca116100fe578063d02103ca14610a2e578063d9c2853914610a55578063e46761c414610a7d575f5ffd5b8063cea5a4c014610a06578063cfa8ed4714610a0e575f5ffd5b8063c754c7ed146109bb578063c89e42df146109e3578063ca69e7dc146109f6578063cce7d0df146109fe575f5ffd5b80638c3d7301116101d8578063ab0475cf116101a8578063b3a326f71161018e578063b3a326f71461098d578063bdfbed7e146109a0578063be647d03146109b3575f5ffd5b8063ab0475cf14610953578063ada8f9191461097a575f5ffd5b80638c3d7301146108fe5780639ee4afa314610906578063a3c573eb14610919578063a8d31bd914610940575f5ffd5b806374f0b0c11161021357806374f0b0c114610887578063750a6e72146108a7578063759664d4146108af5780637df73e27146108eb575f5ffd5b80636e7fbce91461083a5780636ff512cc1461084157806371257022146108545780637388c43614610867575f5ffd5b806339b7ec1611610349578063527570f1116102de5780635ecaca2b116102ae5780636a55f66c116102945780636a55f66c146107ff5780636b8616ce146108125780636e05d2cd14610831575f5ffd5b80635ecaca2b146107cc578063697427f6146107ec575f5ffd5b8063527570f114610758578063542028d51461077857806354fd4d501461078057806359a03e0f146107b9575f5ffd5b806342cde4e81161031957806342cde4e8146106e657806345605267146106ef57806349b7b802146107285780634a5db0c11461074f575f5ffd5b806339b7ec16146106545780633c351e10146106745780633cbc795b146106945780633e1e0121146106d1575f5ffd5b806319451a8f116103d95780632c111c06116103a9578063349d40461161038f578063349d40461461061957806335acd6c21461062e57806336cd6b5b14610641575f5ffd5b80632c111c06146105f9578063314eb17b146104f6575f5ffd5b806319451a8f146104f65780631d0b435e14610509578063267822471461054d57806326f9b76d14610592575f5ffd5b80631489e707116104145780631489e7071461049d578063153c3b7f146104a557806315981b29146104b8578063188d9180146104c0575f5ffd5b806301fcf6a014610445578063052358be1461046b57806306e7666514610480578063107bf28c14610488575b5f5ffd5b610458610453366004613285565b505f90565b6040519081526020015b60405180910390f35b61047e6104793660046132ea565b610b55565b005b61047e610c1d565b610490610fb6565b60405161046291906133a2565b61047e611042565b61047e6104b33660046133f5565b6110c5565b61047e611224565b603e546104e6907501000000000000000000000000000000000000000000900460ff1681565b6040519015158152602001610462565b61047e610504366004613455565b611042565b61051c6105173660046134ac565b6112fb565b6040517fffffffff000000000000000000000000000000000000000000000000000000009091168152602001610462565b60015461056d9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610462565b6105c86105a0366004613285565b60101b7fffff0000000000000000000000000000000000000000000000000000000000001690565b6040517fffff0000000000000000000000000000000000000000000000000000000000009091168152602001610462565b60085461056d9073ffffffffffffffffffffffffffffffffffffffff1681565b61062161134b565b60405161046291906134dd565b61056d61063c36600461358b565b6115f1565b61049061064f3660046135ce565b611626565b60465461056d9073ffffffffffffffffffffffffffffffffffffffff1681565b60095461056d9073ffffffffffffffffffffffffffffffffffffffff1681565b6009546106bc9074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff9091168152602001610462565b6106d961163e565b60405161046291906135e9565b61045860445481565b60075461070f9068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff9091168152602001610462565b61056d7f000000000000000000000000000000000000000000000000000000000000000081565b61045860455481565b60405461056d9073ffffffffffffffffffffffffffffffffffffffff1681565b61049061177d565b60408051808201909152600681527f76312e302e3000000000000000000000000000000000000000000000000000006020820152610490565b6104906107c7366004613784565b61178a565b603d5461056d9073ffffffffffffffffffffffffffffffffffffffff1681565b61047e6107fa3660046138f3565b6117ae565b61045861080d3660046139ce565b611a5a565b610458610820366004613a13565b60066020525f908152604090205481565b61045860055481565b6105c85f81565b61047e61084f3660046135ce565b611ad4565b61047e610862366004613a3a565b611b9d565b603f5461056d9073ffffffffffffffffffffffffffffffffffffffff1681565b603e5461056d9073ffffffffffffffffffffffffffffffffffffffff1681565b61045860ff81565b6104906040518060400160405280600681526020017f76312e302e30000000000000000000000000000000000000000000000000000081525081565b6104e66108f93660046135ce565b611bcf565b61047e611cdc565b61047e610914366004613aee565b611daf565b61056d7f000000000000000000000000000000000000000000000000000000000000000081565b61047e61094e3660046135ce565b611e82565b61056d7f000000000000000000000000000000000000000000000000000000000000000081565b61047e6109883660046135ce565b611f59565b61047e61099b3660046135ce565b612022565b61047e6109ae3660046135ce565b6121a8565b61047e6122c2565b60075461070f90700100000000000000000000000000000000900467ffffffffffffffff1681565b61047e6109f1366004613784565b6123d3565b610458612465565b61045861251f565b6106bc600181565b60025461056d9073ffffffffffffffffffffffffffffffffffffffff1681565b61056d7f000000000000000000000000000000000000000000000000000000000000000081565b610a68610a633660046139ce565b6125ee565b60408051928352602083019190915201610462565b61056d7f000000000000000000000000000000000000000000000000000000000000000081565b610458612634565b60075461070f9067ffffffffffffffff1681565b6105c8610ace366004613285565b90565b610458610adf366004613285565b60416020525f908152604090205481565b61047e610afe366004613b2d565b6126ca565b5f5461056d9062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b603e546104e69074010000000000000000000000000000000000000000900460ff1681565b61047e61272b565b60465473ffffffffffffffffffffffffffffffffffffffff163314610ba6576040517fd0c34d9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c1784848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525050604080516020601f880181900481028201810190925286815292508691508590819084018382808284375f9201919091525061282392505050565b50505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314610c8c576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f805460ff16907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00815c168217905d505f54600290610100900460ff16158015610cdc57505f5460ff8083169116105b610d6d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff80841691909117610100178255905c16600114610ddd576040517fadc06ae700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f54603f80546201000090920473ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff000000000000000000000000000000000000000090921691909117905560038054610e3790613c20565b90505f03610e9b5760025460408051808201909152600681527f4e4f5f55524c00000000000000000000000000000000000000000000000000006020820152610e969173ffffffffffffffffffffffffffffffffffffffff16906128a0565b610f48565b60025460038054610f489273ffffffffffffffffffffffffffffffffffffffff169190610ec790613c20565b80601f0160208091040260200160405190810160405280929190818152602001828054610ef390613c20565b8015610f3e5780601f10610f1557610100808354040283529160200191610f3e565b820191905f5260205f20905b815481529060010190602001808311610f2157829003601f168201915b50505050506128a0565b6001604455610f556129ec565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a150565b60048054610fc390613c20565b80601f0160208091040260200160405190810160405280929190818152602001828054610fef90613c20565b801561103a5780601f106110115761010080835404028352916020019161103a565b820191905f5260205f20905b81548152906001019060200180831161101d57829003601f168201915b505050505081565b603f5473ffffffffffffffffffffffffffffffffffffffff163314611093576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517fd37a223a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60465473ffffffffffffffffffffffffffffffffffffffff163314611116576040517fd0c34d9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82818114611150576040517f059d0ac900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b8181101561121c5761121486868381811061116f5761116f613c6b565b90506020028101906111819190613c98565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508892508791508590508181106111c9576111c9613c6b565b90506020028101906111db9190613c98565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061282392505050565b600101611152565b505050505050565b60405473ffffffffffffffffffffffffffffffffffffffff163314611275576040517f3ac87ac900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603f80546040805473ffffffffffffffffffffffffffffffffffffffff8082167fffffffffffffffffffffffff000000000000000000000000000000000000000080861682179096559490911682558151921680835260208301939093527f67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f9101610fab565b7fffff00000000000000000000000000000000000000000000000000000000000082167dffff00000000000000000000000000000000000000000000000000000000601083901c16175b92915050565b603e546060907501000000000000000000000000000000000000000000900460ff1615611424577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663349d40466040518163ffffffff1660e01b81526004015f60405180830381865afa1580156113da573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820160405261141f9190810190613cf9565b905090565b6042545f9067ffffffffffffffff81111561144157611441613641565b60405190808252806020026020018201604052801561148657816020015b604080518082019091525f81526060602082015281526020019060019003908161145f5790505b5090505f5b6042548110156115eb576040518060400160405280604283815481106114b3576114b3613c6b565b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160435f6042858154811061150c5761150c613c6b565b5f91825260208083209091015473ffffffffffffffffffffffffffffffffffffffff1683528201929092526040019020805461154790613c20565b80601f016020809104026020016040519081016040528092919081815260200182805461157390613c20565b80156115be5780601f10611595576101008083540402835291602001916115be565b820191905f5260205f20905b8154815290600101906020018083116115a157829003601f168201915b50505050508152508282815181106115d8576115d8613c6b565b602090810291909101015260010161148b565b50919050565b60428181548110611600575f80fd5b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b60436020525f908152604090208054610fc390613c20565b603e546060907501000000000000000000000000000000000000000000900460ff1615611712577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633e1e01216040518163ffffffff1660e01b81526004015f60405180830381865afa1580156116cd573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820160405261141f9190810190613e59565b604280548060200260200160405190810160405280929190818152602001828054801561177357602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611748575b5050505050905090565b60038054610fc390613c20565b805160208183018101805160478252928201919093012091528054610fc390613c20565b603f5473ffffffffffffffffffffffffffffffffffffffff1633146117ff576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f805460ff16907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00815c168217905d505f54600290610100900460ff1615801561184f57505f5460ff8083169116105b6118db576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610d64565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff80841691909117610100178255905c1615611949576040517fadc06ae700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61195a89898989895f8a8180612a70565b83156119ac578251151580611970575060445415155b156119a7576040517f996c343000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119f2565b604080515f808252602082019092526119f2916119ea565b604080518082019091525f80825260208201528152602001906001900390816119c45790505b508484612b98565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050505050505050565b5f5f611a6461251f565b90505f5f611a71856125ee565b6040517c010000000000000000000000000000000000000000000000000000000060208201526024810183905260448101829052606481018690529193509150608401604051602081830303815290604052805190602001209350505050919050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314611b2a576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc090602001610fab565b6040517ff57ac68300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e545f907501000000000000000000000000000000000000000000900460ff1615611ca3576040517f7df73e2700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301527f00000000000000000000000000000000000000000000000000000000000000001690637df73e2790602401602060405180830381865afa158015611c7f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113459190613ef3565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526043602052604081208054611cd290613c20565b9050119050919050565b60015473ffffffffffffffffffffffffffffffffffffffff163314611d2d576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001545f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff1673ffffffffffffffffffffffffffffffffffffffff9092166201000081029290921790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e906020015b60405180910390a1565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314611e1e576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8015611e56576040517f3063965400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517fc118263690a4306c74bd1bc80b55962addc2d9e61619ac0b2c2883badbbd01d8905f90a15050565b603f5473ffffffffffffffffffffffffffffffffffffffff163314611ed3576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6046805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f82ae2ec69f24a6de4517a5a45d4983651b578b3d8dc9262af5e352572fc64373910160405180910390a15050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314611faf576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001610fab565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314612091576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603f5473ffffffffffffffffffffffffffffffffffffffff16156120e1576040517f257bb0bf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff811661212e576040517fd6bdac3f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8316908117909155604080515f815260208101929092527f67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f9101610fab565b603f5473ffffffffffffffffffffffffffffffffffffffff1633146121f9576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116612246576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169182178355603f5483519116815260208101919091527fa3d8e5d045432398be30f83ce7c35a7bfc220c1b66cc5bf3f4dd4d539d93fab69101610fab565b603f5473ffffffffffffffffffffffffffffffffffffffff163314612313576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e547501000000000000000000000000000000000000000000900460ff1615612369576040517f278d998800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e80547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff1675010000000000000000000000000000000000000000001790556040517f67ec953bdc8546ede08f8ee91e5205a1d1814e126cb8f5d00a918ddb1eaa292b905f90a1565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314612429576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60036124358282613f59565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b2081604051610fab91906133a2565b603e545f907501000000000000000000000000000000000000000000900460ff1615612518577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663ca69e7dc6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156124f4573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061141f9190614070565b5060425490565b603e545f907501000000000000000000000000000000000000000000900460ff16156125ae577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663cce7d0df6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156124f4573d5f5f3e3d5ffd5b6045546125e7576040517fdd41f1ef00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5060455490565b5f5f82515f1461262a576040517f3063965400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505f928392509050565b603e545f907501000000000000000000000000000000000000000000900460ff16156126c3577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e75235b86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156124f4573d5f5f3e3d5ffd5b5060445490565b603f5473ffffffffffffffffffffffffffffffffffffffff16331461271b576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612726838383612b98565b505050565b603f5473ffffffffffffffffffffffffffffffffffffffff16331461277c576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e547501000000000000000000000000000000000000000000900460ff166127d1576040517f5aa930a000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e80547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff1690556040517f4c75580a56c734245a7418eb07d8a311e1bff79f982fed747da3589630e414be905f90a1565b806047836040516128349190614087565b9081526020016040518091039020908161284e9190613f59565b508160405161285d9190614087565b60405180910390207f2779f9edd5ec4e0a99bffdea4008c8b979200959062a2bf00142acb939ca1b648260405161289491906133a2565b60405180910390a25050565b73ffffffffffffffffffffffffffffffffffffffff82166128ed576040517f7b3a0df600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80515f03612927576040517f8715f5fb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61293082611bcf565b15612967576040517f38615ecc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60428054600181019091557f38dfe4635b27babeca8be38d3b448cb5161a639b899a14825ba9c8d7892eb8c30180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84169081179091555f9081526043602052604090206127268282613f59565b6044546042604051602001612a0292919061409d565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052805160209091012060458190556044547f66d7b0647fdd512b69cbf4f8e1ce8068bfe0b236168e2704ba13b07425eaa74392611da592604292916140f0565b5f54610100900460ff16612b06576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610d64565b73ffffffffffffffffffffffffffffffffffffffff89161580612b3d575073ffffffffffffffffffffffffffffffffffffffff8816155b15612b74576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612b818989898989612d89565b612b8d84848484612eb7565b505050505050505050565b600183511115612c40575f5b60018451612bb29190614184565b811015612c3e5783612bc5826001614197565b81518110612bd557612bd5613c6b565b602002602001015160200151848281518110612bf357612bf3613c6b565b60200260200101516020015111612c36576040517fb9a11d3100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600101612ba4565b505b5f5b8351811015612c9657612c8e848281518110612c6057612c60613c6b565b60200260200101515f0151858381518110612c7d57612c7d613c6b565b60200260200101516020015161300d565b600101612c42565b505f5b8251811015612ced57612ce5838281518110612cb757612cb7613c6b565b60200260200101515f0151848381518110612cd457612cd4613c6b565b6020026020010151602001516128a0565b600101612c99565b5060425460ff1015612d2b576040517f5a7f382c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604254811180612d45575060425415801590612d45575080155b15612d7c576040517faabd5a0900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60448190556127266129ec565b73ffffffffffffffffffffffffffffffffffffffff8516612dd6576040517fe6cd565400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff166201000073ffffffffffffffffffffffffffffffffffffffff8881169190910291909117909155600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000169186169190911790556003612e5d8382613f59565b506004612e6a8282613f59565b5050600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9390931692909217909155505050565b5f54610100900460ff16612f4d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610d64565b603e80547fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff1674010000000000000000000000000000000000000000951515959095027fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff1694909417750100000000000000000000000000000000000000000093151593909302929092179092557fffffffff00000000000000000000000000000000000000000000000000000000165f90815260416020526040902055565b604254808210613049576040517fd244b30700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff166042838154811061307357613073613c6b565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff16146130cb576040517fd244b30700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83165f9081526043602052604081206130f8916131ff565b6042613105600183614184565b8154811061311557613115613c6b565b5f918252602090912001546042805473ffffffffffffffffffffffffffffffffffffffff909216918490811061314d5761314d613c6b565b905f5260205f20015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060428054806131a3576131a36141aa565b5f8281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055505050565b50805461320b90613c20565b5f825580601f1061321a575050565b601f0160209004905f5260205f20908101906132369190613239565b50565b5b8082111561324d575f815560010161323a565b5090565b80357fffffffff0000000000000000000000000000000000000000000000000000000081168114613280575f5ffd5b919050565b5f60208284031215613295575f5ffd5b61329e82613251565b9392505050565b5f5f83601f8401126132b5575f5ffd5b50813567ffffffffffffffff8111156132cc575f5ffd5b6020830191508360208285010111156132e3575f5ffd5b9250929050565b5f5f5f5f604085870312156132fd575f5ffd5b843567ffffffffffffffff811115613313575f5ffd5b61331f878288016132a5565b909550935050602085013567ffffffffffffffff81111561333e575f5ffd5b61334a878288016132a5565b95989497509550505050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f61329e6020830184613356565b5f5f83601f8401126133c4575f5ffd5b50813567ffffffffffffffff8111156133db575f5ffd5b6020830191508360208260051b85010111156132e3575f5ffd5b5f5f5f5f60408587031215613408575f5ffd5b843567ffffffffffffffff81111561341e575f5ffd5b61342a878288016133b4565b909550935050602085013567ffffffffffffffff811115613449575f5ffd5b61334a878288016133b4565b5f5f60408385031215613466575f5ffd5b61346f83613251565b946020939093013593505050565b80357fffff00000000000000000000000000000000000000000000000000000000000081168114613280575f5ffd5b5f5f604083850312156134bd575f5ffd5b6134c68361347d565b91506134d46020840161347d565b90509250929050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561357f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815173ffffffffffffffffffffffffffffffffffffffff815116865260208101519050604060208701526135696040870182613356565b9550506020938401939190910190600101613503565b50929695505050505050565b5f6020828403121561359b575f5ffd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff81168114613236575f5ffd5b8035613280816135a2565b5f602082840312156135de575f5ffd5b813561329e816135a2565b602080825282518282018190525f918401906040840190835b8181101561363657835173ffffffffffffffffffffffffffffffffffffffff16835260209384019390920191600101613602565b509095945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6040805190810167ffffffffffffffff8111828210171561369157613691613641565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156136de576136de613641565b604052919050565b5f67ffffffffffffffff8211156136ff576136ff613641565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b5f61373d613738846136e6565b613697565b9050828152838383011115613750575f5ffd5b828260208301375f602084830101529392505050565b5f82601f830112613775575f5ffd5b61329e8383356020850161372b565b5f60208284031215613794575f5ffd5b813567ffffffffffffffff8111156137aa575f5ffd5b6137b684828501613766565b949350505050565b8015158114613236575f5ffd5b8035613280816137be565b5f67ffffffffffffffff8211156137ef576137ef613641565b5060051b60200190565b5f82601f830112613808575f5ffd5b8135613816613738826137d6565b8082825260208201915060208360051b860101925085831115613837575f5ffd5b602085015b838110156138e957803567ffffffffffffffff81111561385a575f5ffd5b860160408189037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001121561388d575f5ffd5b61389561366e565b60208201356138a3816135a2565b8152604082013567ffffffffffffffff8111156138be575f5ffd5b6138cd8a602083860101613766565b602083015250808552505060208301925060208101905061383c565b5095945050505050565b5f5f5f5f5f5f5f5f610100898b03121561390b575f5ffd5b613914896135c3565b975061392260208a016135c3565b965061393060408a016135c3565b9550606089013567ffffffffffffffff81111561394b575f5ffd5b6139578b828c01613766565b955050608089013567ffffffffffffffff811115613973575f5ffd5b61397f8b828c01613766565b94505061398e60a08a016137cb565b925060c089013567ffffffffffffffff8111156139a9575f5ffd5b6139b58b828c016137f9565b989b979a50959894979396929550929360e00135925050565b5f602082840312156139de575f5ffd5b813567ffffffffffffffff8111156139f4575f5ffd5b8201601f81018413613a04575f5ffd5b6137b68482356020840161372b565b5f60208284031215613a23575f5ffd5b813567ffffffffffffffff8116811461329e575f5ffd5b5f5f5f5f5f5f60c08789031215613a4f575f5ffd5b8635613a5a816135a2565b95506020870135613a6a816135a2565b9450604087013563ffffffff81168114613a82575f5ffd5b93506060870135613a92816135a2565b9250608087013567ffffffffffffffff811115613aad575f5ffd5b613ab989828a01613766565b92505060a087013567ffffffffffffffff811115613ad5575f5ffd5b613ae189828a01613766565b9150509295509295509295565b5f5f60208385031215613aff575f5ffd5b823567ffffffffffffffff811115613b15575f5ffd5b613b21858286016132a5565b90969095509350505050565b5f5f5f60608486031215613b3f575f5ffd5b833567ffffffffffffffff811115613b55575f5ffd5b8401601f81018613613b65575f5ffd5b8035613b73613738826137d6565b8082825260208201915060208360061b850101925088831115613b94575f5ffd5b6020840193505b82841015613be4576040848a031215613bb2575f5ffd5b613bba61366e565b8435613bc5816135a2565b8152602085810135818301529083526040909401939190910190613b9b565b9550505050602084013567ffffffffffffffff811115613c02575f5ffd5b613c0e868287016137f9565b93969395505050506040919091013590565b600181811c90821680613c3457607f821691505b6020821081036115eb577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f5f83357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112613ccb575f5ffd5b83018035915067ffffffffffffffff821115613ce5575f5ffd5b6020019150368190038213156132e3575f5ffd5b5f60208284031215613d09575f5ffd5b815167ffffffffffffffff811115613d1f575f5ffd5b8201601f81018413613d2f575f5ffd5b8051613d3d613738826137d6565b8082825260208201915060208360051b850101925086831115613d5e575f5ffd5b602084015b83811015613e4e57805167ffffffffffffffff811115613d81575f5ffd5b85016040818a037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0011215613db4575f5ffd5b613dbc61366e565b6020820151613dca816135a2565b8152604082015167ffffffffffffffff811115613de5575f5ffd5b60208184010192505089601f830112613dfc575f5ffd5b8151613e0a613738826136e6565b8181528b6020838601011115613e1e575f5ffd5b8160208501602083015e5f6020838301015280602084015250508085525050602083019250602081019050613d63565b509695505050505050565b5f60208284031215613e69575f5ffd5b815167ffffffffffffffff811115613e7f575f5ffd5b8201601f81018413613e8f575f5ffd5b8051613e9d613738826137d6565b8082825260208201915060208360051b850101925086831115613ebe575f5ffd5b6020840193505b82841015613ee9578351613ed8816135a2565b825260209384019390910190613ec5565b9695505050505050565b5f60208284031215613f03575f5ffd5b815161329e816137be565b601f82111561272657805f5260205f20601f840160051c81016020851015613f335750805b601f840160051c820191505b81811015613f52575f8155600101613f3f565b5050505050565b815167ffffffffffffffff811115613f7357613f73613641565b613f8781613f818454613c20565b84613f0e565b6020601f821160018114613fd8575f8315613fa25750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455613f52565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b828110156140255787850151825560209485019460019092019101614005565b508482101561406157868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b5f60208284031215614080575f5ffd5b5051919050565b5f82518060208501845e5f920191825250919050565b8281525f602082018354845f5260205f205f5b828110156140e457815473ffffffffffffffffffffffffffffffffffffffff168452602090930192600191820191016140b0565b50919695505050505050565b606080825284549082018190525f8581526020812090916080840190835b8181101561414257835473ffffffffffffffffffffffffffffffffffffffff1683526001938401936020909301920161410e565b50506020840195909552505060400152919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8181038181111561134557611345614157565b8082018082111561134557611345614157565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea2646970667358221220b084b330a868de1df8eb79a474626b60dab8805a91149c5ff7480390370d0e3a64736f6c634300081c0033",
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

// MAXAGGCHAINSIGNERS is a free data retrieval call binding the contract method 0x750a6e72.
//
// Solidity: function MAX_AGGCHAIN_SIGNERS() view returns(uint256)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) MAXAGGCHAINSIGNERS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "MAX_AGGCHAIN_SIGNERS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXAGGCHAINSIGNERS is a free data retrieval call binding the contract method 0x750a6e72.
//
// Solidity: function MAX_AGGCHAIN_SIGNERS() view returns(uint256)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) MAXAGGCHAINSIGNERS() (*big.Int, error) {
	return _Aggchainecdsamultisig.Contract.MAXAGGCHAINSIGNERS(&_Aggchainecdsamultisig.CallOpts)
}

// MAXAGGCHAINSIGNERS is a free data retrieval call binding the contract method 0x750a6e72.
//
// Solidity: function MAX_AGGCHAIN_SIGNERS() view returns(uint256)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) MAXAGGCHAINSIGNERS() (*big.Int, error) {
	return _Aggchainecdsamultisig.Contract.MAXAGGCHAINSIGNERS(&_Aggchainecdsamultisig.CallOpts)
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

// AddOwnedAggchainVKey is a free data retrieval call binding the contract method 0x19451a8f.
//
// Solidity: function addOwnedAggchainVKey(bytes4 , bytes32 ) view returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) AddOwnedAggchainVKey(opts *bind.CallOpts, arg0 [4]byte, arg1 [32]byte) error {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "addOwnedAggchainVKey", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// AddOwnedAggchainVKey is a free data retrieval call binding the contract method 0x19451a8f.
//
// Solidity: function addOwnedAggchainVKey(bytes4 , bytes32 ) view returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) AddOwnedAggchainVKey(arg0 [4]byte, arg1 [32]byte) error {
	return _Aggchainecdsamultisig.Contract.AddOwnedAggchainVKey(&_Aggchainecdsamultisig.CallOpts, arg0, arg1)
}

// AddOwnedAggchainVKey is a free data retrieval call binding the contract method 0x19451a8f.
//
// Solidity: function addOwnedAggchainVKey(bytes4 , bytes32 ) view returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) AddOwnedAggchainVKey(arg0 [4]byte, arg1 [32]byte) error {
	return _Aggchainecdsamultisig.Contract.AddOwnedAggchainVKey(&_Aggchainecdsamultisig.CallOpts, arg0, arg1)
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

// AggchainMetadata is a free data retrieval call binding the contract method 0x59a03e0f.
//
// Solidity: function aggchainMetadata(string ) view returns(string)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) AggchainMetadata(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "aggchainMetadata", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AggchainMetadata is a free data retrieval call binding the contract method 0x59a03e0f.
//
// Solidity: function aggchainMetadata(string ) view returns(string)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) AggchainMetadata(arg0 string) (string, error) {
	return _Aggchainecdsamultisig.Contract.AggchainMetadata(&_Aggchainecdsamultisig.CallOpts, arg0)
}

// AggchainMetadata is a free data retrieval call binding the contract method 0x59a03e0f.
//
// Solidity: function aggchainMetadata(string ) view returns(string)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) AggchainMetadata(arg0 string) (string, error) {
	return _Aggchainecdsamultisig.Contract.AggchainMetadata(&_Aggchainecdsamultisig.CallOpts, arg0)
}

// AggchainMetadataManager is a free data retrieval call binding the contract method 0x39b7ec16.
//
// Solidity: function aggchainMetadataManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) AggchainMetadataManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "aggchainMetadataManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggchainMetadataManager is a free data retrieval call binding the contract method 0x39b7ec16.
//
// Solidity: function aggchainMetadataManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) AggchainMetadataManager() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.AggchainMetadataManager(&_Aggchainecdsamultisig.CallOpts)
}

// AggchainMetadataManager is a free data retrieval call binding the contract method 0x39b7ec16.
//
// Solidity: function aggchainMetadataManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) AggchainMetadataManager() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.AggchainMetadataManager(&_Aggchainecdsamultisig.CallOpts)
}

// AggchainMultisigHash is a free data retrieval call binding the contract method 0x4a5db0c1.
//
// Solidity: function aggchainMultisigHash() view returns(bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) AggchainMultisigHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "aggchainMultisigHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AggchainMultisigHash is a free data retrieval call binding the contract method 0x4a5db0c1.
//
// Solidity: function aggchainMultisigHash() view returns(bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) AggchainMultisigHash() ([32]byte, error) {
	return _Aggchainecdsamultisig.Contract.AggchainMultisigHash(&_Aggchainecdsamultisig.CallOpts)
}

// AggchainMultisigHash is a free data retrieval call binding the contract method 0x4a5db0c1.
//
// Solidity: function aggchainMultisigHash() view returns(bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) AggchainMultisigHash() ([32]byte, error) {
	return _Aggchainecdsamultisig.Contract.AggchainMultisigHash(&_Aggchainecdsamultisig.CallOpts)
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

// DisableUseDefaultVkeysFlag is a free data retrieval call binding the contract method 0xefe6c9f4.
//
// Solidity: function disableUseDefaultVkeysFlag() view returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) DisableUseDefaultVkeysFlag(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "disableUseDefaultVkeysFlag")

	if err != nil {
		return err
	}

	return err

}

// DisableUseDefaultVkeysFlag is a free data retrieval call binding the contract method 0xefe6c9f4.
//
// Solidity: function disableUseDefaultVkeysFlag() view returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) DisableUseDefaultVkeysFlag() error {
	return _Aggchainecdsamultisig.Contract.DisableUseDefaultVkeysFlag(&_Aggchainecdsamultisig.CallOpts)
}

// DisableUseDefaultVkeysFlag is a free data retrieval call binding the contract method 0xefe6c9f4.
//
// Solidity: function disableUseDefaultVkeysFlag() view returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) DisableUseDefaultVkeysFlag() error {
	return _Aggchainecdsamultisig.Contract.DisableUseDefaultVkeysFlag(&_Aggchainecdsamultisig.CallOpts)
}

// EnableUseDefaultVkeysFlag is a free data retrieval call binding the contract method 0x1489e707.
//
// Solidity: function enableUseDefaultVkeysFlag() view returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) EnableUseDefaultVkeysFlag(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "enableUseDefaultVkeysFlag")

	if err != nil {
		return err
	}

	return err

}

// EnableUseDefaultVkeysFlag is a free data retrieval call binding the contract method 0x1489e707.
//
// Solidity: function enableUseDefaultVkeysFlag() view returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) EnableUseDefaultVkeysFlag() error {
	return _Aggchainecdsamultisig.Contract.EnableUseDefaultVkeysFlag(&_Aggchainecdsamultisig.CallOpts)
}

// EnableUseDefaultVkeysFlag is a free data retrieval call binding the contract method 0x1489e707.
//
// Solidity: function enableUseDefaultVkeysFlag() view returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) EnableUseDefaultVkeysFlag() error {
	return _Aggchainecdsamultisig.Contract.EnableUseDefaultVkeysFlag(&_Aggchainecdsamultisig.CallOpts)
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

// GetAggchainMultisigHash is a free data retrieval call binding the contract method 0xcce7d0df.
//
// Solidity: function getAggchainMultisigHash() view returns(bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) GetAggchainMultisigHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "getAggchainMultisigHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAggchainMultisigHash is a free data retrieval call binding the contract method 0xcce7d0df.
//
// Solidity: function getAggchainMultisigHash() view returns(bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) GetAggchainMultisigHash() ([32]byte, error) {
	return _Aggchainecdsamultisig.Contract.GetAggchainMultisigHash(&_Aggchainecdsamultisig.CallOpts)
}

// GetAggchainMultisigHash is a free data retrieval call binding the contract method 0xcce7d0df.
//
// Solidity: function getAggchainMultisigHash() view returns(bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) GetAggchainMultisigHash() ([32]byte, error) {
	return _Aggchainecdsamultisig.Contract.GetAggchainMultisigHash(&_Aggchainecdsamultisig.CallOpts)
}

// GetAggchainSignerInfos is a free data retrieval call binding the contract method 0x349d4046.
//
// Solidity: function getAggchainSignerInfos() view returns((address,string)[])
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) GetAggchainSignerInfos(opts *bind.CallOpts) ([]IAggchainSignersSignerInfo, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "getAggchainSignerInfos")

	if err != nil {
		return *new([]IAggchainSignersSignerInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAggchainSignersSignerInfo)).(*[]IAggchainSignersSignerInfo)

	return out0, err

}

// GetAggchainSignerInfos is a free data retrieval call binding the contract method 0x349d4046.
//
// Solidity: function getAggchainSignerInfos() view returns((address,string)[])
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) GetAggchainSignerInfos() ([]IAggchainSignersSignerInfo, error) {
	return _Aggchainecdsamultisig.Contract.GetAggchainSignerInfos(&_Aggchainecdsamultisig.CallOpts)
}

// GetAggchainSignerInfos is a free data retrieval call binding the contract method 0x349d4046.
//
// Solidity: function getAggchainSignerInfos() view returns((address,string)[])
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) GetAggchainSignerInfos() ([]IAggchainSignersSignerInfo, error) {
	return _Aggchainecdsamultisig.Contract.GetAggchainSignerInfos(&_Aggchainecdsamultisig.CallOpts)
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
// Solidity: function getAggchainVKey(bytes4 ) pure returns(bytes32 aggchainVKey)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) GetAggchainVKey(opts *bind.CallOpts, arg0 [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "getAggchainVKey", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAggchainVKey is a free data retrieval call binding the contract method 0x01fcf6a0.
//
// Solidity: function getAggchainVKey(bytes4 ) pure returns(bytes32 aggchainVKey)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) GetAggchainVKey(arg0 [4]byte) ([32]byte, error) {
	return _Aggchainecdsamultisig.Contract.GetAggchainVKey(&_Aggchainecdsamultisig.CallOpts, arg0)
}

// GetAggchainVKey is a free data retrieval call binding the contract method 0x01fcf6a0.
//
// Solidity: function getAggchainVKey(bytes4 ) pure returns(bytes32 aggchainVKey)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) GetAggchainVKey(arg0 [4]byte) ([32]byte, error) {
	return _Aggchainecdsamultisig.Contract.GetAggchainVKey(&_Aggchainecdsamultisig.CallOpts, arg0)
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

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) GetThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) GetThreshold() (*big.Int, error) {
	return _Aggchainecdsamultisig.Contract.GetThreshold(&_Aggchainecdsamultisig.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) GetThreshold() (*big.Int, error) {
	return _Aggchainecdsamultisig.Contract.GetThreshold(&_Aggchainecdsamultisig.CallOpts)
}

// GetVKeyAndAggchainParams is a free data retrieval call binding the contract method 0xd9c28539.
//
// Solidity: function getVKeyAndAggchainParams(bytes aggchainData) pure returns(bytes32, bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) GetVKeyAndAggchainParams(opts *bind.CallOpts, aggchainData []byte) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "getVKeyAndAggchainParams", aggchainData)

	if err != nil {
		return *new([32]byte), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// GetVKeyAndAggchainParams is a free data retrieval call binding the contract method 0xd9c28539.
//
// Solidity: function getVKeyAndAggchainParams(bytes aggchainData) pure returns(bytes32, bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) GetVKeyAndAggchainParams(aggchainData []byte) ([32]byte, [32]byte, error) {
	return _Aggchainecdsamultisig.Contract.GetVKeyAndAggchainParams(&_Aggchainecdsamultisig.CallOpts, aggchainData)
}

// GetVKeyAndAggchainParams is a free data retrieval call binding the contract method 0xd9c28539.
//
// Solidity: function getVKeyAndAggchainParams(bytes aggchainData) pure returns(bytes32, bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) GetVKeyAndAggchainParams(aggchainData []byte) ([32]byte, [32]byte, error) {
	return _Aggchainecdsamultisig.Contract.GetVKeyAndAggchainParams(&_Aggchainecdsamultisig.CallOpts, aggchainData)
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

// UpdateOwnedAggchainVKey is a free data retrieval call binding the contract method 0x314eb17b.
//
// Solidity: function updateOwnedAggchainVKey(bytes4 , bytes32 ) view returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) UpdateOwnedAggchainVKey(opts *bind.CallOpts, arg0 [4]byte, arg1 [32]byte) error {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "updateOwnedAggchainVKey", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// UpdateOwnedAggchainVKey is a free data retrieval call binding the contract method 0x314eb17b.
//
// Solidity: function updateOwnedAggchainVKey(bytes4 , bytes32 ) view returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) UpdateOwnedAggchainVKey(arg0 [4]byte, arg1 [32]byte) error {
	return _Aggchainecdsamultisig.Contract.UpdateOwnedAggchainVKey(&_Aggchainecdsamultisig.CallOpts, arg0, arg1)
}

// UpdateOwnedAggchainVKey is a free data retrieval call binding the contract method 0x314eb17b.
//
// Solidity: function updateOwnedAggchainVKey(bytes4 , bytes32 ) view returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) UpdateOwnedAggchainVKey(arg0 [4]byte, arg1 [32]byte) error {
	return _Aggchainecdsamultisig.Contract.UpdateOwnedAggchainVKey(&_Aggchainecdsamultisig.CallOpts, arg0, arg1)
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

// BatchSetAggchainMetadata is a paid mutator transaction binding the contract method 0x153c3b7f.
//
// Solidity: function batchSetAggchainMetadata(string[] keys, string[] values) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) BatchSetAggchainMetadata(opts *bind.TransactOpts, keys []string, values []string) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "batchSetAggchainMetadata", keys, values)
}

// BatchSetAggchainMetadata is a paid mutator transaction binding the contract method 0x153c3b7f.
//
// Solidity: function batchSetAggchainMetadata(string[] keys, string[] values) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) BatchSetAggchainMetadata(keys []string, values []string) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.BatchSetAggchainMetadata(&_Aggchainecdsamultisig.TransactOpts, keys, values)
}

// BatchSetAggchainMetadata is a paid mutator transaction binding the contract method 0x153c3b7f.
//
// Solidity: function batchSetAggchainMetadata(string[] keys, string[] values) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) BatchSetAggchainMetadata(keys []string, values []string) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.BatchSetAggchainMetadata(&_Aggchainecdsamultisig.TransactOpts, keys, values)
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
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, _trustedSequencer common.Address, _gasTokenAddress common.Address, _trustedSequencerURL string, _networkName string, _useDefaultSigners bool, _signersToAdd []IAggchainSignersSignerInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "initialize", _admin, _trustedSequencer, _gasTokenAddress, _trustedSequencerURL, _networkName, _useDefaultSigners, _signersToAdd, _newThreshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x697427f6.
//
// Solidity: function initialize(address _admin, address _trustedSequencer, address _gasTokenAddress, string _trustedSequencerURL, string _networkName, bool _useDefaultSigners, (address,string)[] _signersToAdd, uint256 _newThreshold) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) Initialize(_admin common.Address, _trustedSequencer common.Address, _gasTokenAddress common.Address, _trustedSequencerURL string, _networkName string, _useDefaultSigners bool, _signersToAdd []IAggchainSignersSignerInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.Initialize(&_Aggchainecdsamultisig.TransactOpts, _admin, _trustedSequencer, _gasTokenAddress, _trustedSequencerURL, _networkName, _useDefaultSigners, _signersToAdd, _newThreshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x697427f6.
//
// Solidity: function initialize(address _admin, address _trustedSequencer, address _gasTokenAddress, string _trustedSequencerURL, string _networkName, bool _useDefaultSigners, (address,string)[] _signersToAdd, uint256 _newThreshold) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) Initialize(_admin common.Address, _trustedSequencer common.Address, _gasTokenAddress common.Address, _trustedSequencerURL string, _networkName string, _useDefaultSigners bool, _signersToAdd []IAggchainSignersSignerInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.Initialize(&_Aggchainecdsamultisig.TransactOpts, _admin, _trustedSequencer, _gasTokenAddress, _trustedSequencerURL, _networkName, _useDefaultSigners, _signersToAdd, _newThreshold)
}

// MigrateFromLegacyConsensus is a paid mutator transaction binding the contract method 0x06e76665.
//
// Solidity: function migrateFromLegacyConsensus() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) MigrateFromLegacyConsensus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "migrateFromLegacyConsensus")
}

// MigrateFromLegacyConsensus is a paid mutator transaction binding the contract method 0x06e76665.
//
// Solidity: function migrateFromLegacyConsensus() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) MigrateFromLegacyConsensus() (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.MigrateFromLegacyConsensus(&_Aggchainecdsamultisig.TransactOpts)
}

// MigrateFromLegacyConsensus is a paid mutator transaction binding the contract method 0x06e76665.
//
// Solidity: function migrateFromLegacyConsensus() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) MigrateFromLegacyConsensus() (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.MigrateFromLegacyConsensus(&_Aggchainecdsamultisig.TransactOpts)
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

// SetAggchainMetadata is a paid mutator transaction binding the contract method 0x052358be.
//
// Solidity: function setAggchainMetadata(string key, string value) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) SetAggchainMetadata(opts *bind.TransactOpts, key string, value string) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "setAggchainMetadata", key, value)
}

// SetAggchainMetadata is a paid mutator transaction binding the contract method 0x052358be.
//
// Solidity: function setAggchainMetadata(string key, string value) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) SetAggchainMetadata(key string, value string) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.SetAggchainMetadata(&_Aggchainecdsamultisig.TransactOpts, key, value)
}

// SetAggchainMetadata is a paid mutator transaction binding the contract method 0x052358be.
//
// Solidity: function setAggchainMetadata(string key, string value) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) SetAggchainMetadata(key string, value string) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.SetAggchainMetadata(&_Aggchainecdsamultisig.TransactOpts, key, value)
}

// SetAggchainMetadataManager is a paid mutator transaction binding the contract method 0xa8d31bd9.
//
// Solidity: function setAggchainMetadataManager(address newAggchainMetadataManager) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) SetAggchainMetadataManager(opts *bind.TransactOpts, newAggchainMetadataManager common.Address) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "setAggchainMetadataManager", newAggchainMetadataManager)
}

// SetAggchainMetadataManager is a paid mutator transaction binding the contract method 0xa8d31bd9.
//
// Solidity: function setAggchainMetadataManager(address newAggchainMetadataManager) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) SetAggchainMetadataManager(newAggchainMetadataManager common.Address) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.SetAggchainMetadataManager(&_Aggchainecdsamultisig.TransactOpts, newAggchainMetadataManager)
}

// SetAggchainMetadataManager is a paid mutator transaction binding the contract method 0xa8d31bd9.
//
// Solidity: function setAggchainMetadataManager(address newAggchainMetadataManager) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) SetAggchainMetadataManager(newAggchainMetadataManager common.Address) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.SetAggchainMetadataManager(&_Aggchainecdsamultisig.TransactOpts, newAggchainMetadataManager)
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

// UpdateSignersAndThreshold is a paid mutator transaction binding the contract method 0xf51f563a.
//
// Solidity: function updateSignersAndThreshold((address,uint256)[] _signersToRemove, (address,string)[] _signersToAdd, uint256 _newThreshold) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) UpdateSignersAndThreshold(opts *bind.TransactOpts, _signersToRemove []IAggchainSignersRemoveSignerInfo, _signersToAdd []IAggchainSignersSignerInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "updateSignersAndThreshold", _signersToRemove, _signersToAdd, _newThreshold)
}

// UpdateSignersAndThreshold is a paid mutator transaction binding the contract method 0xf51f563a.
//
// Solidity: function updateSignersAndThreshold((address,uint256)[] _signersToRemove, (address,string)[] _signersToAdd, uint256 _newThreshold) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) UpdateSignersAndThreshold(_signersToRemove []IAggchainSignersRemoveSignerInfo, _signersToAdd []IAggchainSignersSignerInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.UpdateSignersAndThreshold(&_Aggchainecdsamultisig.TransactOpts, _signersToRemove, _signersToAdd, _newThreshold)
}

// UpdateSignersAndThreshold is a paid mutator transaction binding the contract method 0xf51f563a.
//
// Solidity: function updateSignersAndThreshold((address,uint256)[] _signersToRemove, (address,string)[] _signersToAdd, uint256 _newThreshold) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) UpdateSignersAndThreshold(_signersToRemove []IAggchainSignersRemoveSignerInfo, _signersToAdd []IAggchainSignersSignerInfo, _newThreshold *big.Int) (*types.Transaction, error) {
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

// AggchainecdsamultisigAggchainMetadataSetIterator is returned from FilterAggchainMetadataSet and is used to iterate over the raw logs and unpacked data for AggchainMetadataSet events raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigAggchainMetadataSetIterator struct {
	Event *AggchainecdsamultisigAggchainMetadataSet // Event containing the contract specifics and raw log

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
func (it *AggchainecdsamultisigAggchainMetadataSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsamultisigAggchainMetadataSet)
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
		it.Event = new(AggchainecdsamultisigAggchainMetadataSet)
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
func (it *AggchainecdsamultisigAggchainMetadataSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsamultisigAggchainMetadataSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsamultisigAggchainMetadataSet represents a AggchainMetadataSet event raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigAggchainMetadataSet struct {
	Key   common.Hash
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAggchainMetadataSet is a free log retrieval operation binding the contract event 0x2779f9edd5ec4e0a99bffdea4008c8b979200959062a2bf00142acb939ca1b64.
//
// Solidity: event AggchainMetadataSet(string indexed key, string value)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterAggchainMetadataSet(opts *bind.FilterOpts, key []string) (*AggchainecdsamultisigAggchainMetadataSetIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "AggchainMetadataSet", keyRule)
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigAggchainMetadataSetIterator{contract: _Aggchainecdsamultisig.contract, event: "AggchainMetadataSet", logs: logs, sub: sub}, nil
}

// WatchAggchainMetadataSet is a free log subscription operation binding the contract event 0x2779f9edd5ec4e0a99bffdea4008c8b979200959062a2bf00142acb939ca1b64.
//
// Solidity: event AggchainMetadataSet(string indexed key, string value)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) WatchAggchainMetadataSet(opts *bind.WatchOpts, sink chan<- *AggchainecdsamultisigAggchainMetadataSet, key []string) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Aggchainecdsamultisig.contract.WatchLogs(opts, "AggchainMetadataSet", keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsamultisigAggchainMetadataSet)
				if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "AggchainMetadataSet", log); err != nil {
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

// ParseAggchainMetadataSet is a log parse operation binding the contract event 0x2779f9edd5ec4e0a99bffdea4008c8b979200959062a2bf00142acb939ca1b64.
//
// Solidity: event AggchainMetadataSet(string indexed key, string value)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) ParseAggchainMetadataSet(log types.Log) (*AggchainecdsamultisigAggchainMetadataSet, error) {
	event := new(AggchainecdsamultisigAggchainMetadataSet)
	if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "AggchainMetadataSet", log); err != nil {
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

// AggchainecdsamultisigSetAggchainMetadataManagerIterator is returned from FilterSetAggchainMetadataManager and is used to iterate over the raw logs and unpacked data for SetAggchainMetadataManager events raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigSetAggchainMetadataManagerIterator struct {
	Event *AggchainecdsamultisigSetAggchainMetadataManager // Event containing the contract specifics and raw log

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
func (it *AggchainecdsamultisigSetAggchainMetadataManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsamultisigSetAggchainMetadataManager)
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
		it.Event = new(AggchainecdsamultisigSetAggchainMetadataManager)
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
func (it *AggchainecdsamultisigSetAggchainMetadataManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsamultisigSetAggchainMetadataManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsamultisigSetAggchainMetadataManager represents a SetAggchainMetadataManager event raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigSetAggchainMetadataManager struct {
	OldAggchainMetadataManager common.Address
	NewAggchainMetadataManager common.Address
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterSetAggchainMetadataManager is a free log retrieval operation binding the contract event 0x82ae2ec69f24a6de4517a5a45d4983651b578b3d8dc9262af5e352572fc64373.
//
// Solidity: event SetAggchainMetadataManager(address oldAggchainMetadataManager, address newAggchainMetadataManager)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterSetAggchainMetadataManager(opts *bind.FilterOpts) (*AggchainecdsamultisigSetAggchainMetadataManagerIterator, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "SetAggchainMetadataManager")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigSetAggchainMetadataManagerIterator{contract: _Aggchainecdsamultisig.contract, event: "SetAggchainMetadataManager", logs: logs, sub: sub}, nil
}

// WatchSetAggchainMetadataManager is a free log subscription operation binding the contract event 0x82ae2ec69f24a6de4517a5a45d4983651b578b3d8dc9262af5e352572fc64373.
//
// Solidity: event SetAggchainMetadataManager(address oldAggchainMetadataManager, address newAggchainMetadataManager)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) WatchSetAggchainMetadataManager(opts *bind.WatchOpts, sink chan<- *AggchainecdsamultisigSetAggchainMetadataManager) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.WatchLogs(opts, "SetAggchainMetadataManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsamultisigSetAggchainMetadataManager)
				if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "SetAggchainMetadataManager", log); err != nil {
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

// ParseSetAggchainMetadataManager is a log parse operation binding the contract event 0x82ae2ec69f24a6de4517a5a45d4983651b578b3d8dc9262af5e352572fc64373.
//
// Solidity: event SetAggchainMetadataManager(address oldAggchainMetadataManager, address newAggchainMetadataManager)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) ParseSetAggchainMetadataManager(log types.Log) (*AggchainecdsamultisigSetAggchainMetadataManager, error) {
	event := new(AggchainecdsamultisigSetAggchainMetadataManager)
	if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "SetAggchainMetadataManager", log); err != nil {
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
	AggchainSigners         []common.Address
	NewThreshold            *big.Int
	NewAggchainMultisigHash [32]byte
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSignersAndThresholdUpdated is a free log retrieval operation binding the contract event 0x66d7b0647fdd512b69cbf4f8e1ce8068bfe0b236168e2704ba13b07425eaa743.
//
// Solidity: event SignersAndThresholdUpdated(address[] aggchainSigners, uint256 newThreshold, bytes32 newAggchainMultisigHash)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterSignersAndThresholdUpdated(opts *bind.FilterOpts) (*AggchainecdsamultisigSignersAndThresholdUpdatedIterator, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "SignersAndThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigSignersAndThresholdUpdatedIterator{contract: _Aggchainecdsamultisig.contract, event: "SignersAndThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchSignersAndThresholdUpdated is a free log subscription operation binding the contract event 0x66d7b0647fdd512b69cbf4f8e1ce8068bfe0b236168e2704ba13b07425eaa743.
//
// Solidity: event SignersAndThresholdUpdated(address[] aggchainSigners, uint256 newThreshold, bytes32 newAggchainMultisigHash)
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
// Solidity: event SignersAndThresholdUpdated(address[] aggchainSigners, uint256 newThreshold, bytes32 newAggchainMultisigHash)
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

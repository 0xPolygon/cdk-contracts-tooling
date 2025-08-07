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

// AggchainecdsamultisigMetaData contains all meta data concerning the Aggchainecdsamultisig contract.
var AggchainecdsamultisigMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"_rollupManager\",\"type\":\"address\"},{\"internalType\":\"contractIAggLayerGateway\",\"name\":\"_aggLayerGateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AggchainManagerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptySignersArray\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalAccInputHashDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggLayerGatewayAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggchainDataLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggchainType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAggchainManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAggchainManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingVKeyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyVKeyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdTooHighAfterRemoval\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultGatewayAlreadyDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultGatewayAlreadyEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroValueAggchainVKey\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAggchainManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"AcceptAggchainManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldVKeyManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVKeyManager\",\"type\":\"address\"}],\"name\":\"AcceptVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"AddAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DisableUseDefaultGatewayFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EnableUseDefaultGatewayFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"OnVerifyPessimisticECDSAMultisig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"SignerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"SignerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newSignersHash\",\"type\":\"bytes32\"}],\"name\":\"SignersHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"oldThreshold\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newThreshold\",\"type\":\"uint32\"}],\"name\":\"ThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentAggchainManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAggchainManager\",\"type\":\"address\"}],\"name\":\"TransferAggchainManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentVKeyManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingVKeyManager\",\"type\":\"address\"}],\"name\":\"TransferVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousAggchainVKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"UpdateAggchainVKey\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AGGCHAIN_ECDSA_MULTISIG_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AGGCHAIN_TYPE\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONSENSUS_TYPE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAggchainManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptVKeyManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"}],\"name\":\"addMultiSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"addOwnedAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"addSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggLayerGateway\",\"outputs\":[{\"internalType\":\"contractIAggLayerGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggchainManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableUseDefaultGatewayFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableUseDefaultGatewayFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"getAggchainHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"getAggchainTypeFromSelector\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"getAggchainVKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"aggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes2\",\"name\":\"aggchainVKeyVersion\",\"type\":\"bytes2\"},{\"internalType\":\"bytes2\",\"name\":\"aggchainType\",\"type\":\"bytes2\"}],\"name\":\"getAggchainVKeySelector\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"getAggchainVKeyVersionFromSelector\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSignersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"initAggchainManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initializeBytesAggchain\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isSignerMapping\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"onVerifyPessimistic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"ownedAggchainVKeys\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"ownedAggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAggchainManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingVKeyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_signerIndex\",\"type\":\"uint256\"}],\"name\":\"removeSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"signers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signersHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"transferAggchainManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVKeyManager\",\"type\":\"address\"}],\"name\":\"transferVKeyManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"updatedAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"updateOwnedAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_newThreshold\",\"type\":\"uint32\"}],\"name\":\"updateThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useDefaultGateway\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vKeyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x610120604052348015610010575f5ffd5b50604051613a95380380613a9583398101604081905261002f916101c4565b6001600160a01b0380861660a05280851660805280841660c052821660e05284848484848484848461005f6100f0565b505050506001600160a01b038116158061008057506001600160a01b038516155b8061009257506001600160a01b038416155b806100a457506001600160a01b038316155b806100b657506001600160a01b038216155b156100d45760405163f6b2911f60e01b815260040160405180910390fd5b6001600160a01b03166101005250610235975050505050505050565b5f54610100900460ff161561015b5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff90811610156101ab575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146101c1575f5ffd5b50565b5f5f5f5f5f60a086880312156101d8575f5ffd5b85516101e3816101ad565b60208701519095506101f4816101ad565b6040870151909450610205816101ad565b6060870151909350610216816101ad565b6080870151909250610227816101ad565b809150509295509295909350565b60805160a05160c05160e0516101005161380e6102875f395f818161085b0152610b5a01525f818161063a01528181611b910152611d2f01525f61083401525f61095101525f6109a0015261380e5ff3fe608060405234801561000f575f5ffd5b5060043610610388575f3560e01c80637fa07caf116101df578063c8b39d6211610109578063e631476c116100a9578063eb12d61e11610079578063eb12d61e146109f8578063effb847914610a0b578063f851a44014610a2a578063ff90407914610a4f575f5ffd5b8063e631476c146109c2578063e7a7ed02146109ca578063e810f2a2146109de578063e90a3409146109e7575f5ffd5b8063d02103ca116100e4578063d02103ca1461094c578063dc8c424914610973578063e279984e1461097b578063e46761c41461099b575f5ffd5b8063c8b39d6214610911578063cea5a4c014610924578063cfa8ed471461092c575f5ffd5b8063a3c573eb1161017f578063bdfbed7e1161014f578063bdfbed7e146108a3578063bfb193b6146108b6578063c754c7ed146108d6578063c89e42df146108fe575f5ffd5b8063a3c573eb1461082f578063ab0475cf14610856578063ada8f9191461087d578063b3a326f714610890575f5ffd5b80638c3d7301116101ba5780638c3d7301146107f757806394cf795e146107ff5780639ee4afa314610814578063a0c1deb414610827575f5ffd5b80637fa07caf146107af57806385018182146107d15780638ae20ebf146107e4575f5ffd5b806342cde4e8116102c05780636b8616ce11610260578063712570221161023057806371257022146106f85780637388c4361461070b578063759664d41461072b5780637df73e2714610767575f5ffd5b80636b8616ce146106975780636e05d2cd146106b65780636e7fbce9146106bf5780636ff512cc146106e5575f5ffd5b806349b7b8021161029b57806349b7b80214610635578063527570f11461065c578063542028d51461067c5780636a55f66c14610684575f5ffd5b806342cde4e8146105d9578063439fab91146105e957806345605267146105fc575f5ffd5b806326f9b76d1161032b578063368c822c11610306578063368c822c146105615780633bad5426146105695780633c351e101461057c5780633cbc795b1461059c575f5ffd5b806326f9b76d146104c75780632c111c061461052e578063314eb17b1461054e575f5ffd5b806319451a8f1161036657806319451a8f146103d15780631d0b435e146103e45780632079fb9a1461046f57806326782247146104a7575f5ffd5b806301fcf6a01461038c578063107bf28c146103b257806315981b29146103c7575b5f5ffd5b61039f61039a366004612d44565b610a74565b6040519081526020015b60405180910390f35b6103ba610bde565b6040516103a99190612d66565b6103cf610c6a565b005b6103cf6103df366004612db9565b610d48565b61043e6103f2366004612e12565b60101c7dffff00000000000000000000000000000000000000000000000000000000167fffff000000000000000000000000000000000000000000000000000000000000919091161790565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020016103a9565b61048261047d366004612e43565b610ea9565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016103a9565b6001546104829073ffffffffffffffffffffffffffffffffffffffff1681565b6104fd6104d5366004612d44565b60101b7fffff0000000000000000000000000000000000000000000000000000000000001690565b6040517fffff00000000000000000000000000000000000000000000000000000000000090911681526020016103a9565b6008546104829073ffffffffffffffffffffffffffffffffffffffff1681565b6103cf61055c366004612db9565b610ede565b6103cf611014565b6103cf610577366004612e7b565b6110f0565b6009546104829073ffffffffffffffffffffffffffffffffffffffff1681565b6009546105c49074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff90911681526020016103a9565b6076546105c49063ffffffff1681565b6103cf6105f7366004612f93565b6111e2565b60075461061c9068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020016103a9565b6104827f000000000000000000000000000000000000000000000000000000000000000081565b6040546104829073ffffffffffffffffffffffffffffffffffffffff1681565b6103ba6115bf565b61039f610692366004612f93565b6115cc565b61039f6106a5366004612fe0565b60066020525f908152604090205481565b61039f60055481565b6104fd7e0200000000000000000000000000000000000000000000000000000000000081565b6103cf6106f3366004613007565b611771565b6103cf610706366004613051565b61183a565b603f546104829073ffffffffffffffffffffffffffffffffffffffff1681565b6103ba6040518060400160405280600681526020017f76312e302e30000000000000000000000000000000000000000000000000000081525081565b61079f610775366004613007565b73ffffffffffffffffffffffffffffffffffffffff165f9081526075602052604090205460ff1690565b60405190151581526020016103a9565b61079f6107bd366004613007565b60756020525f908152604090205460ff1681565b6103cf6107df366004613007565b61186c565b6103cf6107f23660046130fd565b61193e565b6103cf611a4f565b610807611b22565b6040516103a99190613118565b6103cf610822366004613170565b611b8f565b60745461039f565b6104827f000000000000000000000000000000000000000000000000000000000000000081565b6104827f000000000000000000000000000000000000000000000000000000000000000081565b6103cf61088b366004613007565b611c64565b6103cf61089e366004613007565b611d2d565b6103cf6108b1366004613007565b611e63565b603e546104829073ffffffffffffffffffffffffffffffffffffffff1681565b60075461061c90700100000000000000000000000000000000900467ffffffffffffffff1681565b6103cf61090c3660046131de565b611f7d565b6103cf61091f366004613210565b61200f565b6105c4600181565b6002546104829073ffffffffffffffffffffffffffffffffffffffff1681565b6104827f000000000000000000000000000000000000000000000000000000000000000081565b6103cf6120e4565b603d546104829073ffffffffffffffffffffffffffffffffffffffff1681565b6104827f000000000000000000000000000000000000000000000000000000000000000081565b6103cf6121db565b60075461061c9067ffffffffffffffff1681565b61039f60775481565b6104fd6109f5366004612d44565b90565b6103cf610a06366004613007565b6122ea565b61039f610a19366004612d44565b60416020525f908152604090205481565b5f546104829062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b603e5461079f9074010000000000000000000000000000000000000000900460ff1681565b603e545f9074010000000000000000000000000000000000000000900460ff1615158103610b0a57507fffffffff0000000000000000000000000000000000000000000000000000000081165f9081526041602052604090205480610b05576040517f925e5a3a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b919050565b6040517f6cabfdab0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000831660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690636cabfdab90602401602060405180830381865afa158015610bb4573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bd89190613271565b92915050565b60048054610beb90613288565b80601f0160208091040260200160405190810160405280929190818152602001828054610c1790613288565b8015610c625780601f10610c3957610100808354040283529160200191610c62565b820191905f5260205f20905b815481529060010190602001808311610c4557829003601f168201915b505050505081565b60405473ffffffffffffffffffffffffffffffffffffffff163314610cbb576040517f3ac87ac900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603f80546040805473ffffffffffffffffffffffffffffffffffffffff8082167fffffffffffffffffffffffff000000000000000000000000000000000000000080861682179096559490911682558151921680835260208301939093527f67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f91015b60405180910390a150565b603d5473ffffffffffffffffffffffffffffffffffffffff163314610d99576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80610dd0576040517fe1dbcf2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f9081526041602052604090205415610e38576040517fe3cc761000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f81815260416020908152604091829020849055815192835282018390527f6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a7991015b60405180910390a15050565b60748181548110610eb8575f80fd5b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b603d5473ffffffffffffffffffffffffffffffffffffffff163314610f2f576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f90815260416020526040902054610f96576040517ff360deaf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f818152604160209081526040918290208054908590558251938452908301819052908201839052907f0aa5f73c189fb0b0a7cc98ae5fa89dfc16595480396208483518178435ed5b4f9060600160405180910390a1505050565b603e5473ffffffffffffffffffffffffffffffffffffffff163314611065576040517f05882cf000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d8054603e805473ffffffffffffffffffffffffffffffffffffffff8082167fffffffffffffffffffffffff0000000000000000000000000000000000000000808616821790965594909116909155604080519190921680825260208201939093527fbb3b066505f14a628f4ba4187a046abd4dd17e96591d7a9ed31c91c79322ffe29101610d3d565b603f5473ffffffffffffffffffffffffffffffffffffffff163314611141576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6074545f0361117c576040517fd244b30700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60745461118b906001906132d9565b60765463ffffffff1611156111cc576040517f0fe639ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111d6828261234f565b6111de6125fe565b5050565b603f5473ffffffffffffffffffffffffffffffffffffffff163314611233576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f805460ff16907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00815c168217905d505f54600290610100900460ff1615801561128357505f5460ff8083169116105b611314576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff80841691909117610100178255815c16900361145f575f5f5f5f5f5f5f5f5f5f5f8c8060200190518101906113709190613417565b9a509a509a509a509a509a509a509a509a509a509a50600260f01b7dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166113db887fffff00000000000000000000000000000000000000000000000000000000000060109190911b1690565b7fffff0000000000000000000000000000000000000000000000000000000000001614611434576040517f4570795000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61143e8b8b612676565b61144f85858585858e8e8e8e61276e565b5050505050505050505050611565565b60ff5f5c16600103611533575f5f5f5f5f5f878060200190518101906114859190613524565b949a509298509096509450925090507fffff000000000000000000000000000000000000000000000000000000000000601083901b167e0200000000000000000000000000000000000000000000000000000000000014611512576040517f4570795000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61151c8686612676565b61152884848484612963565b505050505050611565565b6040517fadc06ae700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001610e9d565b60038054610beb90613288565b5f8151602014611608576040517f3063965400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8280602001905181019061161d91906135b1565b90507fffff000000000000000000000000000000000000000000000000000000000000601082901b167e020000000000000000000000000000000000000000000000000000000000001461169d576040517f4570795000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60016116a882610a74565b6077546076546040516116f3929163ffffffff169060200191825260e01b7fffffffff0000000000000000000000000000000000000000000000000000000016602082015260240190565b604051602081830303815290604052805190602001206040516020016117539392919060e09390931b7fffffffff000000000000000000000000000000000000000000000000000000001683526004830191909152602482015260440190565b60405160208183030381529060405280519060200120915050919050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff1633146117c7576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc090602001610d3d565b6040517ff57ac68300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d5473ffffffffffffffffffffffffffffffffffffffff1633146118bd576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217909255603d546040805191909316815260208101919091527fc54ae01017d0b80bd8af833f66387d6eb547dc16c8206faf13d0b72764aab8b29101610d3d565b603f5473ffffffffffffffffffffffffffffffffffffffff16331461198f576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff811615806119a9575060745463ffffffff8216115b156119e0576040517faabd5a0900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6076805463ffffffff8381167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000083168117909355604080519190921680825260208201939093527f92950fd771d49788c1d64aa1d59e19056f9fe09256db0f449205fe98a2aecb839101610e9d565b60015473ffffffffffffffffffffffffffffffffffffffff163314611aa0576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001545f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff1673ffffffffffffffffffffffffffffffffffffffff9092166201000081029290921790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e906020015b60405180910390a1565b60606074805480602002602001604051908101604052809291908181526020018280548015611b8557602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611b5a575b5050505050905090565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314611bfe576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208114611c38576040517f3063965400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517fc118263690a4306c74bd1bc80b55962addc2d9e61619ac0b2c2883badbbd01d8905f90a15050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314611cba576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001610d3d565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314611d9c576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116611de9576040517fd6bdac3f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8316908117909155604080515f815260208101929092527f67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f9101610d3d565b603f5473ffffffffffffffffffffffffffffffffffffffff163314611eb4576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116611f01576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169182178355603f5483519116815260208101919091527fa3d8e5d045432398be30f83ce7c35a7bfc220c1b66cc5bf3f4dd4d539d93fab69101610d3d565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314611fd3576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6003611fdf8282613618565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b2081604051610d3d9190612d66565b603f5473ffffffffffffffffffffffffffffffffffffffff163314612060576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f81900361209a576040517f7a67bdeb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b818110156120db576120d38383838181106120b9576120b961372f565b90506020020160208101906120ce9190613007565b612ab9565b60010161209c565b506111de6125fe565b603d5473ffffffffffffffffffffffffffffffffffffffff163314612135576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e5474010000000000000000000000000000000000000000900460ff16612189576040517f62de044500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690556040517f67dd1717a1952df380cb73eecb312e949df6d6a086bd7f88669005341972528e905f90a1565b603d5473ffffffffffffffffffffffffffffffffffffffff16331461222c576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e5474010000000000000000000000000000000000000000900460ff1615612281576040517f93be805100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790556040517fb6563aed80fde357e737eb0d19f246a58cb6bfd469933d05701ecbad0f2dca84905f90a1565b603f5473ffffffffffffffffffffffffffffffffffffffff16331461233b576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61234481612ab9565b61234c6125fe565b50565b607454811061238a576040517fd244b30700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff16607482815481106123b4576123b461372f565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff161461240c576040517fd244b30700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82165f9081526075602052604090205460ff1661246a576040517fd244b30700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82165f90815260756020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055607480546124c4906001906132d9565b815481106124d4576124d461372f565b5f918252602090912001546074805473ffffffffffffffffffffffffffffffffffffffff909216918390811061250c5761250c61372f565b905f5260205f20015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060748054806125625761256261375c565b5f8281526020812082017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590910190915560405173ffffffffffffffffffffffffffffffffffffffff8416917f3525e22824a8a7df2c9a6029941c824cf95b6447f1e13d5128fd3826d35afe8b91a25050565b60746040516020016126109190613789565b604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018152908290528051602091820120607781905582527f75fc775cec2bbccee8d459709737ff8264a0b0d5808e53c121f64514b69865a89101611b18565b81515f036126b0576040517f7a67bdeb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff811615806126c9575081518163ffffffff16115b15612700576040517faabd5a0900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b82518110156127355761272d8382815181106127205761272061372f565b6020026020010151612ab9565b600101612702565b50607680547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff83161790556111de6125fe565b5f54610100900460ff16612804576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161130b565b73ffffffffffffffffffffffffffffffffffffffff8916158061283b575073ffffffffffffffffffffffffffffffffffffffff8816155b8061285a575073ffffffffffffffffffffffffffffffffffffffff8116155b15612891576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61289e8989898989612c36565b603e80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000095151595909502949094179093557fffffffff00000000000000000000000000000000000000000000000000000000165f90815260416020526040902055603d80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9092169190911790555050505050565b5f54610100900460ff166129f9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161130b565b603e80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000095151595909502949094179093557fffffffff00000000000000000000000000000000000000000000000000000000165f90815260416020526040902055603d80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909216919091179055565b73ffffffffffffffffffffffffffffffffffffffff8116612b06576040517f7b3a0df600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81165f9081526075602052604090205460ff1615612b65576040517f38615ecc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6074805460018082019092557f19a0b39aa25ac793b5f6e9a0534364cc0b3fd1ea9b651e79c7f50a59d48ef8130180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84169081179091555f8181526075602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016909417909355915190917f47d1c22a25bb3a5d4e481b9b1e6944c2eade3181a0a20b495ed61d35b5323f2491a250565b5f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff166201000073ffffffffffffffffffffffffffffffffffffffff8881169190910291909117909155600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000169186169190911790556003612cbd8382613618565b506004612cca8282613618565b5050600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9390931692909217909155505050565b7fffffffff000000000000000000000000000000000000000000000000000000008116811461234c575f5ffd5b5f60208284031215612d54575f5ffd5b8135612d5f81612d17565b9392505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b5f5f60408385031215612dca575f5ffd5b8235612dd581612d17565b946020939093013593505050565b80357fffff00000000000000000000000000000000000000000000000000000000000081168114610b05575f5ffd5b5f5f60408385031215612e23575f5ffd5b612e2c83612de3565b9150612e3a60208401612de3565b90509250929050565b5f60208284031215612e53575f5ffd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff8116811461234c575f5ffd5b5f5f60408385031215612e8c575f5ffd5b8235612dd581612e5a565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715612f0b57612f0b612e97565b604052919050565b5f67ffffffffffffffff821115612f2c57612f2c612e97565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b5f612f6a612f6584612f13565b612ec4565b9050828152838383011115612f7d575f5ffd5b828260208301375f602084830101529392505050565b5f60208284031215612fa3575f5ffd5b813567ffffffffffffffff811115612fb9575f5ffd5b8201601f81018413612fc9575f5ffd5b612fd884823560208401612f58565b949350505050565b5f60208284031215612ff0575f5ffd5b813567ffffffffffffffff81168114612d5f575f5ffd5b5f60208284031215613017575f5ffd5b8135612d5f81612e5a565b63ffffffff8116811461234c575f5ffd5b5f82601f830112613042575f5ffd5b612d5f83833560208501612f58565b5f5f5f5f5f5f60c08789031215613066575f5ffd5b863561307181612e5a565b9550602087013561308181612e5a565b9450604087013561309181613022565b935060608701356130a181612e5a565b9250608087013567ffffffffffffffff8111156130bc575f5ffd5b6130c889828a01613033565b92505060a087013567ffffffffffffffff8111156130e4575f5ffd5b6130f089828a01613033565b9150509295509295509295565b5f6020828403121561310d575f5ffd5b8135612d5f81613022565b602080825282518282018190525f918401906040840190835b8181101561316557835173ffffffffffffffffffffffffffffffffffffffff16835260209384019390920191600101613131565b509095945050505050565b5f5f60208385031215613181575f5ffd5b823567ffffffffffffffff811115613197575f5ffd5b8301601f810185136131a7575f5ffd5b803567ffffffffffffffff8111156131bd575f5ffd5b8560208284010111156131ce575f5ffd5b6020919091019590945092505050565b5f602082840312156131ee575f5ffd5b813567ffffffffffffffff811115613204575f5ffd5b612fd884828501613033565b5f5f60208385031215613221575f5ffd5b823567ffffffffffffffff811115613237575f5ffd5b8301601f81018513613247575f5ffd5b803567ffffffffffffffff81111561325d575f5ffd5b8560208260051b84010111156131ce575f5ffd5b5f60208284031215613281575f5ffd5b5051919050565b600181811c9082168061329c57607f821691505b6020821081036132d3577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b81810381811115610bd8577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f82601f830112613320575f5ffd5b815167ffffffffffffffff81111561333a5761333a612e97565b8060051b61334a60208201612ec4565b91825260208185018101929081019086841115613365575f5ffd5b6020860192505b8383101561339057825161337f81612e5a565b82526020928301929091019061336c565b9695505050505050565b8051610b0581613022565b80518015158114610b05575f5ffd5b8051610b0581612d17565b8051610b0581612e5a565b5f82601f8301126133d9575f5ffd5b81516133e7612f6582612f13565b8181528460208386010111156133fb575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f5f5f5f5f5f5f5f5f5f5f6101608c8e031215613432575f5ffd5b8b5167ffffffffffffffff811115613448575f5ffd5b6134548e828f01613311565b9b505061346360208d0161339a565b995061347160408d016133a5565b60608d0151909950975061348760808d016133b4565b965061349560a08d016133bf565b95506134a360c08d016133bf565b94506134b160e08d016133bf565b93506134c06101008d016133bf565b92506101208c015167ffffffffffffffff8111156134dc575f5ffd5b6134e88e828f016133ca565b9250506101408c015167ffffffffffffffff811115613505575f5ffd5b6135118e828f016133ca565b9150509295989b509295989b9093969950565b5f5f5f5f5f5f60c08789031215613539575f5ffd5b865167ffffffffffffffff81111561354f575f5ffd5b61355b89828a01613311565b965050602087015161356c81613022565b945061357a604088016133a5565b60608801516080890151919550935061359281612d17565b60a08801519092506135a381612e5a565b809150509295509295509295565b5f602082840312156135c1575f5ffd5b8151612d5f81612d17565b601f82111561361357805f5260205f20601f840160051c810160208510156135f15750805b601f840160051c820191505b81811015613610575f81556001016135fd565b50505b505050565b815167ffffffffffffffff81111561363257613632612e97565b613646816136408454613288565b846135cc565b6020601f821160018114613697575f83156136615750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455613610565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b828110156136e457878501518255602094850194600190920191016136c4565b508482101561372057868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b5f818354839150845f5260205f205f5b828110156137cd57815473ffffffffffffffffffffffffffffffffffffffff16845260209093019260019182019101613799565b50919594505050505056fea2646970667358221220d439d4d96427d63061ea4261e73f3790f3eccbefe9cd31737fccf2ceff8d5c4a64736f6c634300081c0033",
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

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) GetSigners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "getSigners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) GetSigners() ([]common.Address, error) {
	return _Aggchainecdsamultisig.Contract.GetSigners(&_Aggchainecdsamultisig.CallOpts)
}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) GetSigners() ([]common.Address, error) {
	return _Aggchainecdsamultisig.Contract.GetSigners(&_Aggchainecdsamultisig.CallOpts)
}

// GetSignersCount is a free data retrieval call binding the contract method 0xa0c1deb4.
//
// Solidity: function getSignersCount() view returns(uint256)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) GetSignersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "getSignersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSignersCount is a free data retrieval call binding the contract method 0xa0c1deb4.
//
// Solidity: function getSignersCount() view returns(uint256)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) GetSignersCount() (*big.Int, error) {
	return _Aggchainecdsamultisig.Contract.GetSignersCount(&_Aggchainecdsamultisig.CallOpts)
}

// GetSignersCount is a free data retrieval call binding the contract method 0xa0c1deb4.
//
// Solidity: function getSignersCount() view returns(uint256)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) GetSignersCount() (*big.Int, error) {
	return _Aggchainecdsamultisig.Contract.GetSignersCount(&_Aggchainecdsamultisig.CallOpts)
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

// IsSignerMapping is a free data retrieval call binding the contract method 0x7fa07caf.
//
// Solidity: function isSignerMapping(address ) view returns(bool)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) IsSignerMapping(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "isSignerMapping", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSignerMapping is a free data retrieval call binding the contract method 0x7fa07caf.
//
// Solidity: function isSignerMapping(address ) view returns(bool)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) IsSignerMapping(arg0 common.Address) (bool, error) {
	return _Aggchainecdsamultisig.Contract.IsSignerMapping(&_Aggchainecdsamultisig.CallOpts, arg0)
}

// IsSignerMapping is a free data retrieval call binding the contract method 0x7fa07caf.
//
// Solidity: function isSignerMapping(address ) view returns(bool)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) IsSignerMapping(arg0 common.Address) (bool, error) {
	return _Aggchainecdsamultisig.Contract.IsSignerMapping(&_Aggchainecdsamultisig.CallOpts, arg0)
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

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) Signers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "signers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) Signers(arg0 *big.Int) (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.Signers(&_Aggchainecdsamultisig.CallOpts, arg0)
}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) Signers(arg0 *big.Int) (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.Signers(&_Aggchainecdsamultisig.CallOpts, arg0)
}

// SignersHash is a free data retrieval call binding the contract method 0xe810f2a2.
//
// Solidity: function signersHash() view returns(bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) SignersHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "signersHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SignersHash is a free data retrieval call binding the contract method 0xe810f2a2.
//
// Solidity: function signersHash() view returns(bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) SignersHash() ([32]byte, error) {
	return _Aggchainecdsamultisig.Contract.SignersHash(&_Aggchainecdsamultisig.CallOpts)
}

// SignersHash is a free data retrieval call binding the contract method 0xe810f2a2.
//
// Solidity: function signersHash() view returns(bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) SignersHash() ([32]byte, error) {
	return _Aggchainecdsamultisig.Contract.SignersHash(&_Aggchainecdsamultisig.CallOpts)
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

// AddMultiSigners is a paid mutator transaction binding the contract method 0xc8b39d62.
//
// Solidity: function addMultiSigners(address[] _signers) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) AddMultiSigners(opts *bind.TransactOpts, _signers []common.Address) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "addMultiSigners", _signers)
}

// AddMultiSigners is a paid mutator transaction binding the contract method 0xc8b39d62.
//
// Solidity: function addMultiSigners(address[] _signers) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) AddMultiSigners(_signers []common.Address) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.AddMultiSigners(&_Aggchainecdsamultisig.TransactOpts, _signers)
}

// AddMultiSigners is a paid mutator transaction binding the contract method 0xc8b39d62.
//
// Solidity: function addMultiSigners(address[] _signers) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) AddMultiSigners(_signers []common.Address) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.AddMultiSigners(&_Aggchainecdsamultisig.TransactOpts, _signers)
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

// AddSigner is a paid mutator transaction binding the contract method 0xeb12d61e.
//
// Solidity: function addSigner(address _signer) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) AddSigner(opts *bind.TransactOpts, _signer common.Address) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "addSigner", _signer)
}

// AddSigner is a paid mutator transaction binding the contract method 0xeb12d61e.
//
// Solidity: function addSigner(address _signer) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) AddSigner(_signer common.Address) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.AddSigner(&_Aggchainecdsamultisig.TransactOpts, _signer)
}

// AddSigner is a paid mutator transaction binding the contract method 0xeb12d61e.
//
// Solidity: function addSigner(address _signer) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) AddSigner(_signer common.Address) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.AddSigner(&_Aggchainecdsamultisig.TransactOpts, _signer)
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

// RemoveSigner is a paid mutator transaction binding the contract method 0x3bad5426.
//
// Solidity: function removeSigner(address _signer, uint256 _signerIndex) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) RemoveSigner(opts *bind.TransactOpts, _signer common.Address, _signerIndex *big.Int) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "removeSigner", _signer, _signerIndex)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x3bad5426.
//
// Solidity: function removeSigner(address _signer, uint256 _signerIndex) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) RemoveSigner(_signer common.Address, _signerIndex *big.Int) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.RemoveSigner(&_Aggchainecdsamultisig.TransactOpts, _signer, _signerIndex)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x3bad5426.
//
// Solidity: function removeSigner(address _signer, uint256 _signerIndex) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) RemoveSigner(_signer common.Address, _signerIndex *big.Int) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.RemoveSigner(&_Aggchainecdsamultisig.TransactOpts, _signer, _signerIndex)
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

// UpdateThreshold is a paid mutator transaction binding the contract method 0x8ae20ebf.
//
// Solidity: function updateThreshold(uint32 _newThreshold) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) UpdateThreshold(opts *bind.TransactOpts, _newThreshold uint32) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "updateThreshold", _newThreshold)
}

// UpdateThreshold is a paid mutator transaction binding the contract method 0x8ae20ebf.
//
// Solidity: function updateThreshold(uint32 _newThreshold) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) UpdateThreshold(_newThreshold uint32) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.UpdateThreshold(&_Aggchainecdsamultisig.TransactOpts, _newThreshold)
}

// UpdateThreshold is a paid mutator transaction binding the contract method 0x8ae20ebf.
//
// Solidity: function updateThreshold(uint32 _newThreshold) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) UpdateThreshold(_newThreshold uint32) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.UpdateThreshold(&_Aggchainecdsamultisig.TransactOpts, _newThreshold)
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

// AggchainecdsamultisigSignerAddedIterator is returned from FilterSignerAdded and is used to iterate over the raw logs and unpacked data for SignerAdded events raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigSignerAddedIterator struct {
	Event *AggchainecdsamultisigSignerAdded // Event containing the contract specifics and raw log

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
func (it *AggchainecdsamultisigSignerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsamultisigSignerAdded)
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
		it.Event = new(AggchainecdsamultisigSignerAdded)
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
func (it *AggchainecdsamultisigSignerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsamultisigSignerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsamultisigSignerAdded represents a SignerAdded event raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigSignerAdded struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignerAdded is a free log retrieval operation binding the contract event 0x47d1c22a25bb3a5d4e481b9b1e6944c2eade3181a0a20b495ed61d35b5323f24.
//
// Solidity: event SignerAdded(address indexed signer)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterSignerAdded(opts *bind.FilterOpts, signer []common.Address) (*AggchainecdsamultisigSignerAddedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "SignerAdded", signerRule)
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigSignerAddedIterator{contract: _Aggchainecdsamultisig.contract, event: "SignerAdded", logs: logs, sub: sub}, nil
}

// WatchSignerAdded is a free log subscription operation binding the contract event 0x47d1c22a25bb3a5d4e481b9b1e6944c2eade3181a0a20b495ed61d35b5323f24.
//
// Solidity: event SignerAdded(address indexed signer)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) WatchSignerAdded(opts *bind.WatchOpts, sink chan<- *AggchainecdsamultisigSignerAdded, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Aggchainecdsamultisig.contract.WatchLogs(opts, "SignerAdded", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsamultisigSignerAdded)
				if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "SignerAdded", log); err != nil {
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

// ParseSignerAdded is a log parse operation binding the contract event 0x47d1c22a25bb3a5d4e481b9b1e6944c2eade3181a0a20b495ed61d35b5323f24.
//
// Solidity: event SignerAdded(address indexed signer)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) ParseSignerAdded(log types.Log) (*AggchainecdsamultisigSignerAdded, error) {
	event := new(AggchainecdsamultisigSignerAdded)
	if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "SignerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsamultisigSignerRemovedIterator is returned from FilterSignerRemoved and is used to iterate over the raw logs and unpacked data for SignerRemoved events raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigSignerRemovedIterator struct {
	Event *AggchainecdsamultisigSignerRemoved // Event containing the contract specifics and raw log

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
func (it *AggchainecdsamultisigSignerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsamultisigSignerRemoved)
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
		it.Event = new(AggchainecdsamultisigSignerRemoved)
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
func (it *AggchainecdsamultisigSignerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsamultisigSignerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsamultisigSignerRemoved represents a SignerRemoved event raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigSignerRemoved struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignerRemoved is a free log retrieval operation binding the contract event 0x3525e22824a8a7df2c9a6029941c824cf95b6447f1e13d5128fd3826d35afe8b.
//
// Solidity: event SignerRemoved(address indexed signer)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterSignerRemoved(opts *bind.FilterOpts, signer []common.Address) (*AggchainecdsamultisigSignerRemovedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "SignerRemoved", signerRule)
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigSignerRemovedIterator{contract: _Aggchainecdsamultisig.contract, event: "SignerRemoved", logs: logs, sub: sub}, nil
}

// WatchSignerRemoved is a free log subscription operation binding the contract event 0x3525e22824a8a7df2c9a6029941c824cf95b6447f1e13d5128fd3826d35afe8b.
//
// Solidity: event SignerRemoved(address indexed signer)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) WatchSignerRemoved(opts *bind.WatchOpts, sink chan<- *AggchainecdsamultisigSignerRemoved, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Aggchainecdsamultisig.contract.WatchLogs(opts, "SignerRemoved", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsamultisigSignerRemoved)
				if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "SignerRemoved", log); err != nil {
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

// ParseSignerRemoved is a log parse operation binding the contract event 0x3525e22824a8a7df2c9a6029941c824cf95b6447f1e13d5128fd3826d35afe8b.
//
// Solidity: event SignerRemoved(address indexed signer)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) ParseSignerRemoved(log types.Log) (*AggchainecdsamultisigSignerRemoved, error) {
	event := new(AggchainecdsamultisigSignerRemoved)
	if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "SignerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsamultisigSignersHashUpdatedIterator is returned from FilterSignersHashUpdated and is used to iterate over the raw logs and unpacked data for SignersHashUpdated events raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigSignersHashUpdatedIterator struct {
	Event *AggchainecdsamultisigSignersHashUpdated // Event containing the contract specifics and raw log

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
func (it *AggchainecdsamultisigSignersHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsamultisigSignersHashUpdated)
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
		it.Event = new(AggchainecdsamultisigSignersHashUpdated)
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
func (it *AggchainecdsamultisigSignersHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsamultisigSignersHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsamultisigSignersHashUpdated represents a SignersHashUpdated event raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigSignersHashUpdated struct {
	NewSignersHash [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSignersHashUpdated is a free log retrieval operation binding the contract event 0x75fc775cec2bbccee8d459709737ff8264a0b0d5808e53c121f64514b69865a8.
//
// Solidity: event SignersHashUpdated(bytes32 newSignersHash)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterSignersHashUpdated(opts *bind.FilterOpts) (*AggchainecdsamultisigSignersHashUpdatedIterator, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "SignersHashUpdated")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigSignersHashUpdatedIterator{contract: _Aggchainecdsamultisig.contract, event: "SignersHashUpdated", logs: logs, sub: sub}, nil
}

// WatchSignersHashUpdated is a free log subscription operation binding the contract event 0x75fc775cec2bbccee8d459709737ff8264a0b0d5808e53c121f64514b69865a8.
//
// Solidity: event SignersHashUpdated(bytes32 newSignersHash)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) WatchSignersHashUpdated(opts *bind.WatchOpts, sink chan<- *AggchainecdsamultisigSignersHashUpdated) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.WatchLogs(opts, "SignersHashUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsamultisigSignersHashUpdated)
				if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "SignersHashUpdated", log); err != nil {
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

// ParseSignersHashUpdated is a log parse operation binding the contract event 0x75fc775cec2bbccee8d459709737ff8264a0b0d5808e53c121f64514b69865a8.
//
// Solidity: event SignersHashUpdated(bytes32 newSignersHash)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) ParseSignersHashUpdated(log types.Log) (*AggchainecdsamultisigSignersHashUpdated, error) {
	event := new(AggchainecdsamultisigSignersHashUpdated)
	if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "SignersHashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsamultisigThresholdUpdatedIterator is returned from FilterThresholdUpdated and is used to iterate over the raw logs and unpacked data for ThresholdUpdated events raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigThresholdUpdatedIterator struct {
	Event *AggchainecdsamultisigThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *AggchainecdsamultisigThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsamultisigThresholdUpdated)
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
		it.Event = new(AggchainecdsamultisigThresholdUpdated)
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
func (it *AggchainecdsamultisigThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsamultisigThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsamultisigThresholdUpdated represents a ThresholdUpdated event raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigThresholdUpdated struct {
	OldThreshold uint32
	NewThreshold uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterThresholdUpdated is a free log retrieval operation binding the contract event 0x92950fd771d49788c1d64aa1d59e19056f9fe09256db0f449205fe98a2aecb83.
//
// Solidity: event ThresholdUpdated(uint32 oldThreshold, uint32 newThreshold)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterThresholdUpdated(opts *bind.FilterOpts) (*AggchainecdsamultisigThresholdUpdatedIterator, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "ThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigThresholdUpdatedIterator{contract: _Aggchainecdsamultisig.contract, event: "ThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchThresholdUpdated is a free log subscription operation binding the contract event 0x92950fd771d49788c1d64aa1d59e19056f9fe09256db0f449205fe98a2aecb83.
//
// Solidity: event ThresholdUpdated(uint32 oldThreshold, uint32 newThreshold)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) WatchThresholdUpdated(opts *bind.WatchOpts, sink chan<- *AggchainecdsamultisigThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.WatchLogs(opts, "ThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsamultisigThresholdUpdated)
				if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "ThresholdUpdated", log); err != nil {
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

// ParseThresholdUpdated is a log parse operation binding the contract event 0x92950fd771d49788c1d64aa1d59e19056f9fe09256db0f449205fe98a2aecb83.
//
// Solidity: event ThresholdUpdated(uint32 oldThreshold, uint32 newThreshold)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) ParseThresholdUpdated(log types.Log) (*AggchainecdsamultisigThresholdUpdated, error) {
	event := new(AggchainecdsamultisigThresholdUpdated)
	if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "ThresholdUpdated", log); err != nil {
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

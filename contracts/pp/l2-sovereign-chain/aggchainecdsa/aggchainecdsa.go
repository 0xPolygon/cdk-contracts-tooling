// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aggchainecdsa

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

// AggchainecdsaMetaData contains all meta data concerning the Aggchainecdsa contract.
var AggchainecdsaMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"_rollupManager\",\"type\":\"address\"},{\"internalType\":\"contractIAggLayerGateway\",\"name\":\"_aggLayerGateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalAccInputHashDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggLayerGatewayAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingVKeyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyVKeyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultGatewayAlreadyDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultGatewayAlreadyEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroValueAggchainVKey\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldVKeyManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVKeyManager\",\"type\":\"address\"}],\"name\":\"AcceptVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"AddAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DisableUseDefaultGatewayFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EnableUseDefaultGatewayFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"OnVerifyPessimisticECDSA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentVKeyManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingVKeyManager\",\"type\":\"address\"}],\"name\":\"TransferVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousAggchainVKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"UpdateAggchainVKey\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AGGCHAIN_TYPE\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONSENSUS_TYPE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptVKeyManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"addOwnedAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggLayerGateway\",\"outputs\":[{\"internalType\":\"contractIAggLayerGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableUseDefaultGatewayFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableUseDefaultGatewayFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"getAggchainHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"getAggchainVKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"aggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes2\",\"name\":\"aggchainVKeyVersion\",\"type\":\"bytes2\"},{\"internalType\":\"bytes2\",\"name\":\"aggchainType\",\"type\":\"bytes2\"}],\"name\":\"getAggchainVKeySelector\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initializeBytesAggchain\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"onVerifyPessimistic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"ownedAggchainVKeys\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"ownedAggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingVKeyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVKeyManager\",\"type\":\"address\"}],\"name\":\"transferVKeyManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"updatedAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"updateOwnedAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useDefaultGateway\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vKeyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x610120604052348015610010575f5ffd5b5060405161261638038061261683398101604081905261002f9161017b565b6001600160a01b0380861660a05280851660805280841660c052821660e05284848484848484848461005f6100a7565b5050506001600160a01b038216905061008b5760405163ae7e6e0f60e01b815260040160405180910390fd5b6001600160a01b031661010052506101ec975050505050505050565b5f54610100900460ff16156101125760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161015610162575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114610178575f5ffd5b50565b5f5f5f5f5f60a0868803121561018f575f5ffd5b855161019a81610164565b60208701519095506101ab81610164565b60408701519094506101bc81610164565b60608701519093506101cd81610164565b60808701519092506101de81610164565b809150509295509295909350565b60805160a05160c05160e051610100516123d861023e5f395f81816105e1015261088a01525f81816104c401528181610d17015261136f01525f6105ba01525f61069e01525f6106ed01526123d85ff3fe608060405234801561000f575f5ffd5b50600436106102b7575f3560e01c80638501818211610171578063cfa8ed47116100d2578063e631476c11610088578063effb84791161006e578063effb84791461072b578063f851a4401461074a578063ff9040791461076f575f5ffd5b8063e631476c1461070f578063e7a7ed0214610717575f5ffd5b8063dc8c4249116100b8578063dc8c4249146106c0578063e279984e146106c8578063e46761c4146106e8575f5ffd5b8063cfa8ed4714610679578063d02103ca14610699575f5ffd5b8063ada8f91911610127578063c754c7ed1161010d578063c754c7ed14610636578063c89e42df1461065e578063cea5a4c014610671575f5ffd5b8063ada8f91914610603578063bfb193b614610616575f5ffd5b80639ee4afa3116101575780639ee4afa3146105a2578063a3c573eb146105b5578063ab0475cf146105dc575f5ffd5b806385018182146105875780638c3d73011461059a575f5ffd5b8063439fab911161021b5780636b8616ce116101d15780636e7fbce9116101b75780636e7fbce9146105295780636ff512cc146105615780637125702214610574575f5ffd5b80636b8616ce146105015780636e05d2cd14610520575f5ffd5b806349b7b8021161020157806349b7b802146104bf578063542028d5146104e65780636a55f66c146104ee575f5ffd5b8063439fab91146104735780634560526714610486575f5ffd5b80632c111c0611610270578063368c822c11610256578063368c822c1461040e5780633c351e10146104165780633cbc795b14610436575f5ffd5b80632c111c06146103db578063314eb17b146103fb575f5ffd5b806319451a8f116102a057806319451a8f146102f65780631d0b435e1461030b5780632678224714610396575f5ffd5b806301fcf6a0146102bb578063107bf28c146102e1575b5f5ffd5b6102ce6102c9366004611be6565b6107a4565b6040519081526020015b60405180910390f35b6102e961090e565b6040516102d89190611c06565b610309610304366004611c59565b61099a565b005b610365610319366004611cb1565b60101c7dffff00000000000000000000000000000000000000000000000000000000167fffff000000000000000000000000000000000000000000000000000000000000919091161790565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020016102d8565b6001546103b69073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102d8565b6008546103b69073ffffffffffffffffffffffffffffffffffffffff1681565b610309610409366004611c59565b610afb565b610309610c32565b6009546103b69073ffffffffffffffffffffffffffffffffffffffff1681565b60095461045e9074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff90911681526020016102d8565b610309610481366004611de4565b610d15565b6007546104a69068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020016102d8565b6103b67f000000000000000000000000000000000000000000000000000000000000000081565b6102e9610fc5565b6102ce6104fc366004611de4565b610fd2565b6102ce61050f366004611e31565b60066020525f908152604090205481565b6102ce60055481565b6105305f81565b6040517fffff00000000000000000000000000000000000000000000000000000000000090911681526020016102d8565b61030961056f366004611e79565b6110ce565b610309610582366004611eb2565b611197565b610309610595366004611e79565b6111c9565b61030961129b565b6103096105b0366004611f66565b61136d565b6103b67f000000000000000000000000000000000000000000000000000000000000000081565b6103b67f000000000000000000000000000000000000000000000000000000000000000081565b610309610611366004611e79565b61141d565b603d546103b69073ffffffffffffffffffffffffffffffffffffffff1681565b6007546104a690700100000000000000000000000000000000900467ffffffffffffffff1681565b61030961066c366004611fd4565b6114e6565b61045e600181565b6002546103b69073ffffffffffffffffffffffffffffffffffffffff1681565b6103b67f000000000000000000000000000000000000000000000000000000000000000081565b610309611578565b603c546103b69073ffffffffffffffffffffffffffffffffffffffff1681565b6103b67f000000000000000000000000000000000000000000000000000000000000000081565b61030961166f565b6007546104a69067ffffffffffffffff1681565b6102ce610739366004611be6565b603e6020525f908152604090205481565b5f546103b69062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b603d546107949074010000000000000000000000000000000000000000900460ff1681565b60405190151581526020016102d8565b603d545f9074010000000000000000000000000000000000000000900460ff161515810361083a57507fffffffff0000000000000000000000000000000000000000000000000000000081165f908152603e602052604090205480610835576040517f925e5a3a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b919050565b6040517f6cabfdab0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000831660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690636cabfdab90602401602060405180830381865afa1580156108e4573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109089190612006565b92915050565b6004805461091b9061201d565b80601f01602080910402602001604051908101604052809291908181526020018280546109479061201d565b80156109925780601f1061096957610100808354040283529160200191610992565b820191905f5260205f20905b81548152906001019060200180831161097557829003601f168201915b505050505081565b603c5473ffffffffffffffffffffffffffffffffffffffff1633146109eb576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80610a22576040517fe1dbcf2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f908152603e602052604090205415610a8a576040517fe3cc761000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f818152603e6020908152604091829020849055815192835282018390527f6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a7991015b60405180910390a15050565b603c5473ffffffffffffffffffffffffffffffffffffffff163314610b4c576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f908152603e6020526040902054610bb3576040517ff360deaf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f818152603e60209081526040918290208054908590558251938452908301819052908201839052907f0aa5f73c189fb0b0a7cc98ae5fa89dfc16595480396208483518178435ed5b4f906060015b60405180910390a1505050565b603d5473ffffffffffffffffffffffffffffffffffffffff163314610c83576040517f05882cf000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603c8054603d805473ffffffffffffffffffffffffffffffffffffffff8082167fffffffffffffffffffffffff0000000000000000000000000000000000000000808616821790965594909116909155604080519190921680825260208201939093527fbb3b066505f14a628f4ba4187a046abd4dd17e96591d7a9ed31c91c79322ffe291015b60405180910390a150565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314610d84576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f805460ff16907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00815c168217905d505f54600290610100900460ff16158015610dd457505f5460ff8083169116105b610e65576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff80841691909117610100178255815c169003610ef4575f5f5f5f5f5f5f5f5f8a806020019051810190610ebf91906120d5565b985098509850985098509850985098509850610ee685858585858e8e8e8e5f60f01b61177e565b505050505050505050610f6b565b60ff5f5c16600103610f39575f5f5f5f85806020019051810190610f1891906121a7565b9350935093509350610f30848484845f60f01b611933565b50505050610f6b565b6040517fadc06ae700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001610aef565b6003805461091b9061201d565b5f5f82806020019051810190610fe891906121f7565b5090507fffff00000000000000000000000000000000000000000000000000000000000081166001611019826107a4565b60025460405160609190911b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166020820152603401604051602081830303815290604052805190602001206040516020016110af9392919060e09390931b7fffffffff000000000000000000000000000000000000000000000000000000001683526004830191909152602482015260440190565b6040516020818303038152906040528051906020012092505050919050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314611124576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc090602001610d0a565b6040517ff57ac68300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603c5473ffffffffffffffffffffffffffffffffffffffff16331461121a576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217909255603c546040805191909316815260208101919091527fc54ae01017d0b80bd8af833f66387d6eb547dc16c8206faf13d0b72764aab8b29101610d0a565b60015473ffffffffffffffffffffffffffffffffffffffff1633146112ec576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001545f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff1673ffffffffffffffffffffffffffffffffffffffff9092166201000081029290921790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e9060200160405180910390a1565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1633146113dc576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6113e982840184612223565b9150507fffc88430d30048515bd904f24ae1b72af4b78151ac346c6de9271ff5ab7899b481604051610c2591815260200190565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314611473576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001610d0a565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff16331461153c576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6003611548828261228b565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b2081604051610d0a9190611c06565b603c5473ffffffffffffffffffffffffffffffffffffffff1633146115c9576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d5474010000000000000000000000000000000000000000900460ff1661161d576040517f62de044500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690556040517f67dd1717a1952df380cb73eecb312e949df6d6a086bd7f88669005341972528e905f90a1565b603c5473ffffffffffffffffffffffffffffffffffffffff1633146116c0576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d5474010000000000000000000000000000000000000000900460ff1615611715576040517f93be805100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790556040517fb6563aed80fde357e737eb0d19f246a58cb6bfd469933d05701ecbad0f2dca84905f90a1565b5f54610100900460ff16611814576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610e5c565b6118218a8a8a8a8a611ad6565b603d80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000008715150217905583603e5f601084901c7dffff00000000000000000000000000000000000000000000000000000000167fffff0000000000000000000000000000000000000000000000000000000000008716177fffffffff0000000000000000000000000000000000000000000000000000000016815260208101919091526040015f205550603c80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff929092169190911790555050505050505050565b5f54610100900460ff166119c9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610e5c565b603d80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000008715150217905583603e5f601084901c7dffff00000000000000000000000000000000000000000000000000000000167fffff0000000000000000000000000000000000000000000000000000000000008716177fffffffff0000000000000000000000000000000000000000000000000000000016815260208101919091526040015f205550603c80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055505050565b5f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff166201000073ffffffffffffffffffffffffffffffffffffffff8881169190910291909117909155600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000169186169190911790556003611b5d838261228b565b506004611b6a828261228b565b5050600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9390931692909217909155505050565b80357fffffffff0000000000000000000000000000000000000000000000000000000081168114610835575f5ffd5b5f60208284031215611bf6575f5ffd5b611bff82611bb7565b9392505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b5f5f60408385031215611c6a575f5ffd5b611c7383611bb7565b946020939093013593505050565b7fffff00000000000000000000000000000000000000000000000000000000000081168114611cae575f5ffd5b50565b5f5f60408385031215611cc2575f5ffd5b8235611ccd81611c81565b91506020830135611cdd81611c81565b809150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611d5c57611d5c611ce8565b604052919050565b5f67ffffffffffffffff821115611d7d57611d7d611ce8565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b5f611dbb611db684611d64565b611d15565b9050828152838383011115611dce575f5ffd5b828260208301375f602084830101529392505050565b5f60208284031215611df4575f5ffd5b813567ffffffffffffffff811115611e0a575f5ffd5b8201601f81018413611e1a575f5ffd5b611e2984823560208401611da9565b949350505050565b5f60208284031215611e41575f5ffd5b813567ffffffffffffffff81168114611bff575f5ffd5b73ffffffffffffffffffffffffffffffffffffffff81168114611cae575f5ffd5b5f60208284031215611e89575f5ffd5b8135611bff81611e58565b5f82601f830112611ea3575f5ffd5b611bff83833560208501611da9565b5f5f5f5f5f5f60c08789031215611ec7575f5ffd5b8635611ed281611e58565b95506020870135611ee281611e58565b9450604087013563ffffffff81168114611efa575f5ffd5b93506060870135611f0a81611e58565b9250608087013567ffffffffffffffff811115611f25575f5ffd5b611f3189828a01611e94565b92505060a087013567ffffffffffffffff811115611f4d575f5ffd5b611f5989828a01611e94565b9150509295509295509295565b5f5f60208385031215611f77575f5ffd5b823567ffffffffffffffff811115611f8d575f5ffd5b8301601f81018513611f9d575f5ffd5b803567ffffffffffffffff811115611fb3575f5ffd5b856020828401011115611fc4575f5ffd5b6020919091019590945092505050565b5f60208284031215611fe4575f5ffd5b813567ffffffffffffffff811115611ffa575f5ffd5b611e2984828501611e94565b5f60208284031215612016575f5ffd5b5051919050565b600181811c9082168061203157607f821691505b602082108103612068577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b80518015158114610835575f5ffd5b805161083581611e58565b5f82601f830112612097575f5ffd5b81516120a5611db682611d64565b8181528460208386010111156120b9575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f5f5f5f5f5f5f5f5f6101208a8c0312156120ee575f5ffd5b6120f78a61206e565b60208b015160408c0151919a50985061210f81611c81565b965061211d60608b0161207d565b955061212b60808b0161207d565b945061213960a08b0161207d565b935061214760c08b0161207d565b925060e08a015167ffffffffffffffff811115612162575f5ffd5b61216e8c828d01612088565b9250506101008a015167ffffffffffffffff81111561218b575f5ffd5b6121978c828d01612088565b9150509295985092959850929598565b5f5f5f5f608085870312156121ba575f5ffd5b6121c38561206e565b6020860151604087015191955093506121db81611c81565b60608601519092506121ec81611e58565b939692955090935050565b5f5f60408385031215612208575f5ffd5b825161221381611c81565b6020939093015192949293505050565b5f5f60408385031215612234575f5ffd5b8235611c7381611c81565b601f82111561228657805f5260205f20601f840160051c810160208510156122645750805b601f840160051c820191505b81811015612283575f8155600101612270565b50505b505050565b815167ffffffffffffffff8111156122a5576122a5611ce8565b6122b9816122b3845461201d565b8461223f565b6020601f82116001811461230a575f83156122d45750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455612283565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b828110156123575787850151825560209485019460019092019101612337565b508482101561239357868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b0190555056fea2646970667358221220e5ba790f2e551f4e6160eff327d016f9edc431acc5c7d9fe89c4c13612f6e70764736f6c634300081c0033",
}

// AggchainecdsaABI is the input ABI used to generate the binding from.
// Deprecated: Use AggchainecdsaMetaData.ABI instead.
var AggchainecdsaABI = AggchainecdsaMetaData.ABI

// AggchainecdsaBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AggchainecdsaMetaData.Bin instead.
var AggchainecdsaBin = AggchainecdsaMetaData.Bin

// DeployAggchainecdsa deploys a new Ethereum contract, binding an instance of Aggchainecdsa to it.
func DeployAggchainecdsa(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address, _rollupManager common.Address, _aggLayerGateway common.Address) (common.Address, *types.Transaction, *Aggchainecdsa, error) {
	parsed, err := AggchainecdsaMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AggchainecdsaBin), backend, _globalExitRootManager, _pol, _bridgeAddress, _rollupManager, _aggLayerGateway)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Aggchainecdsa{AggchainecdsaCaller: AggchainecdsaCaller{contract: contract}, AggchainecdsaTransactor: AggchainecdsaTransactor{contract: contract}, AggchainecdsaFilterer: AggchainecdsaFilterer{contract: contract}}, nil
}

// Aggchainecdsa is an auto generated Go binding around an Ethereum contract.
type Aggchainecdsa struct {
	AggchainecdsaCaller     // Read-only binding to the contract
	AggchainecdsaTransactor // Write-only binding to the contract
	AggchainecdsaFilterer   // Log filterer for contract events
}

// AggchainecdsaCaller is an auto generated read-only Go binding around an Ethereum contract.
type AggchainecdsaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggchainecdsaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AggchainecdsaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggchainecdsaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AggchainecdsaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggchainecdsaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AggchainecdsaSession struct {
	Contract     *Aggchainecdsa    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AggchainecdsaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AggchainecdsaCallerSession struct {
	Contract *AggchainecdsaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AggchainecdsaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AggchainecdsaTransactorSession struct {
	Contract     *AggchainecdsaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AggchainecdsaRaw is an auto generated low-level Go binding around an Ethereum contract.
type AggchainecdsaRaw struct {
	Contract *Aggchainecdsa // Generic contract binding to access the raw methods on
}

// AggchainecdsaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AggchainecdsaCallerRaw struct {
	Contract *AggchainecdsaCaller // Generic read-only contract binding to access the raw methods on
}

// AggchainecdsaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AggchainecdsaTransactorRaw struct {
	Contract *AggchainecdsaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAggchainecdsa creates a new instance of Aggchainecdsa, bound to a specific deployed contract.
func NewAggchainecdsa(address common.Address, backend bind.ContractBackend) (*Aggchainecdsa, error) {
	contract, err := bindAggchainecdsa(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aggchainecdsa{AggchainecdsaCaller: AggchainecdsaCaller{contract: contract}, AggchainecdsaTransactor: AggchainecdsaTransactor{contract: contract}, AggchainecdsaFilterer: AggchainecdsaFilterer{contract: contract}}, nil
}

// NewAggchainecdsaCaller creates a new read-only instance of Aggchainecdsa, bound to a specific deployed contract.
func NewAggchainecdsaCaller(address common.Address, caller bind.ContractCaller) (*AggchainecdsaCaller, error) {
	contract, err := bindAggchainecdsa(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggchainecdsaCaller{contract: contract}, nil
}

// NewAggchainecdsaTransactor creates a new write-only instance of Aggchainecdsa, bound to a specific deployed contract.
func NewAggchainecdsaTransactor(address common.Address, transactor bind.ContractTransactor) (*AggchainecdsaTransactor, error) {
	contract, err := bindAggchainecdsa(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggchainecdsaTransactor{contract: contract}, nil
}

// NewAggchainecdsaFilterer creates a new log filterer instance of Aggchainecdsa, bound to a specific deployed contract.
func NewAggchainecdsaFilterer(address common.Address, filterer bind.ContractFilterer) (*AggchainecdsaFilterer, error) {
	contract, err := bindAggchainecdsa(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggchainecdsaFilterer{contract: contract}, nil
}

// bindAggchainecdsa binds a generic wrapper to an already deployed contract.
func bindAggchainecdsa(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AggchainecdsaMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggchainecdsa *AggchainecdsaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aggchainecdsa.Contract.AggchainecdsaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggchainecdsa *AggchainecdsaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.AggchainecdsaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggchainecdsa *AggchainecdsaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.AggchainecdsaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggchainecdsa *AggchainecdsaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aggchainecdsa.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggchainecdsa *AggchainecdsaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggchainecdsa *AggchainecdsaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.contract.Transact(opts, method, params...)
}

// AGGCHAINTYPE is a free data retrieval call binding the contract method 0x6e7fbce9.
//
// Solidity: function AGGCHAIN_TYPE() view returns(bytes2)
func (_Aggchainecdsa *AggchainecdsaCaller) AGGCHAINTYPE(opts *bind.CallOpts) ([2]byte, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "AGGCHAIN_TYPE")

	if err != nil {
		return *new([2]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([2]byte)).(*[2]byte)

	return out0, err

}

// AGGCHAINTYPE is a free data retrieval call binding the contract method 0x6e7fbce9.
//
// Solidity: function AGGCHAIN_TYPE() view returns(bytes2)
func (_Aggchainecdsa *AggchainecdsaSession) AGGCHAINTYPE() ([2]byte, error) {
	return _Aggchainecdsa.Contract.AGGCHAINTYPE(&_Aggchainecdsa.CallOpts)
}

// AGGCHAINTYPE is a free data retrieval call binding the contract method 0x6e7fbce9.
//
// Solidity: function AGGCHAIN_TYPE() view returns(bytes2)
func (_Aggchainecdsa *AggchainecdsaCallerSession) AGGCHAINTYPE() ([2]byte, error) {
	return _Aggchainecdsa.Contract.AGGCHAINTYPE(&_Aggchainecdsa.CallOpts)
}

// CONSENSUSTYPE is a free data retrieval call binding the contract method 0xcea5a4c0.
//
// Solidity: function CONSENSUS_TYPE() view returns(uint32)
func (_Aggchainecdsa *AggchainecdsaCaller) CONSENSUSTYPE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "CONSENSUS_TYPE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CONSENSUSTYPE is a free data retrieval call binding the contract method 0xcea5a4c0.
//
// Solidity: function CONSENSUS_TYPE() view returns(uint32)
func (_Aggchainecdsa *AggchainecdsaSession) CONSENSUSTYPE() (uint32, error) {
	return _Aggchainecdsa.Contract.CONSENSUSTYPE(&_Aggchainecdsa.CallOpts)
}

// CONSENSUSTYPE is a free data retrieval call binding the contract method 0xcea5a4c0.
//
// Solidity: function CONSENSUS_TYPE() view returns(uint32)
func (_Aggchainecdsa *AggchainecdsaCallerSession) CONSENSUSTYPE() (uint32, error) {
	return _Aggchainecdsa.Contract.CONSENSUSTYPE(&_Aggchainecdsa.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Aggchainecdsa *AggchainecdsaCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Aggchainecdsa *AggchainecdsaSession) Admin() (common.Address, error) {
	return _Aggchainecdsa.Contract.Admin(&_Aggchainecdsa.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Aggchainecdsa *AggchainecdsaCallerSession) Admin() (common.Address, error) {
	return _Aggchainecdsa.Contract.Admin(&_Aggchainecdsa.CallOpts)
}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Aggchainecdsa *AggchainecdsaCaller) AggLayerGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "aggLayerGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Aggchainecdsa *AggchainecdsaSession) AggLayerGateway() (common.Address, error) {
	return _Aggchainecdsa.Contract.AggLayerGateway(&_Aggchainecdsa.CallOpts)
}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Aggchainecdsa *AggchainecdsaCallerSession) AggLayerGateway() (common.Address, error) {
	return _Aggchainecdsa.Contract.AggLayerGateway(&_Aggchainecdsa.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Aggchainecdsa *AggchainecdsaCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Aggchainecdsa *AggchainecdsaSession) BridgeAddress() (common.Address, error) {
	return _Aggchainecdsa.Contract.BridgeAddress(&_Aggchainecdsa.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Aggchainecdsa *AggchainecdsaCallerSession) BridgeAddress() (common.Address, error) {
	return _Aggchainecdsa.Contract.BridgeAddress(&_Aggchainecdsa.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Aggchainecdsa *AggchainecdsaCaller) ForceBatchAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "forceBatchAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Aggchainecdsa *AggchainecdsaSession) ForceBatchAddress() (common.Address, error) {
	return _Aggchainecdsa.Contract.ForceBatchAddress(&_Aggchainecdsa.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Aggchainecdsa *AggchainecdsaCallerSession) ForceBatchAddress() (common.Address, error) {
	return _Aggchainecdsa.Contract.ForceBatchAddress(&_Aggchainecdsa.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Aggchainecdsa *AggchainecdsaCaller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Aggchainecdsa *AggchainecdsaSession) ForceBatchTimeout() (uint64, error) {
	return _Aggchainecdsa.Contract.ForceBatchTimeout(&_Aggchainecdsa.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Aggchainecdsa *AggchainecdsaCallerSession) ForceBatchTimeout() (uint64, error) {
	return _Aggchainecdsa.Contract.ForceBatchTimeout(&_Aggchainecdsa.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Aggchainecdsa *AggchainecdsaCaller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Aggchainecdsa *AggchainecdsaSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Aggchainecdsa.Contract.ForcedBatches(&_Aggchainecdsa.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Aggchainecdsa *AggchainecdsaCallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Aggchainecdsa.Contract.ForcedBatches(&_Aggchainecdsa.CallOpts, arg0)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Aggchainecdsa *AggchainecdsaCaller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Aggchainecdsa *AggchainecdsaSession) GasTokenAddress() (common.Address, error) {
	return _Aggchainecdsa.Contract.GasTokenAddress(&_Aggchainecdsa.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Aggchainecdsa *AggchainecdsaCallerSession) GasTokenAddress() (common.Address, error) {
	return _Aggchainecdsa.Contract.GasTokenAddress(&_Aggchainecdsa.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Aggchainecdsa *AggchainecdsaCaller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Aggchainecdsa *AggchainecdsaSession) GasTokenNetwork() (uint32, error) {
	return _Aggchainecdsa.Contract.GasTokenNetwork(&_Aggchainecdsa.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Aggchainecdsa *AggchainecdsaCallerSession) GasTokenNetwork() (uint32, error) {
	return _Aggchainecdsa.Contract.GasTokenNetwork(&_Aggchainecdsa.CallOpts)
}

// GetAggchainHash is a free data retrieval call binding the contract method 0x6a55f66c.
//
// Solidity: function getAggchainHash(bytes aggchainData) view returns(bytes32)
func (_Aggchainecdsa *AggchainecdsaCaller) GetAggchainHash(opts *bind.CallOpts, aggchainData []byte) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "getAggchainHash", aggchainData)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAggchainHash is a free data retrieval call binding the contract method 0x6a55f66c.
//
// Solidity: function getAggchainHash(bytes aggchainData) view returns(bytes32)
func (_Aggchainecdsa *AggchainecdsaSession) GetAggchainHash(aggchainData []byte) ([32]byte, error) {
	return _Aggchainecdsa.Contract.GetAggchainHash(&_Aggchainecdsa.CallOpts, aggchainData)
}

// GetAggchainHash is a free data retrieval call binding the contract method 0x6a55f66c.
//
// Solidity: function getAggchainHash(bytes aggchainData) view returns(bytes32)
func (_Aggchainecdsa *AggchainecdsaCallerSession) GetAggchainHash(aggchainData []byte) ([32]byte, error) {
	return _Aggchainecdsa.Contract.GetAggchainHash(&_Aggchainecdsa.CallOpts, aggchainData)
}

// GetAggchainVKey is a free data retrieval call binding the contract method 0x01fcf6a0.
//
// Solidity: function getAggchainVKey(bytes4 aggchainVKeySelector) view returns(bytes32 aggchainVKey)
func (_Aggchainecdsa *AggchainecdsaCaller) GetAggchainVKey(opts *bind.CallOpts, aggchainVKeySelector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "getAggchainVKey", aggchainVKeySelector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAggchainVKey is a free data retrieval call binding the contract method 0x01fcf6a0.
//
// Solidity: function getAggchainVKey(bytes4 aggchainVKeySelector) view returns(bytes32 aggchainVKey)
func (_Aggchainecdsa *AggchainecdsaSession) GetAggchainVKey(aggchainVKeySelector [4]byte) ([32]byte, error) {
	return _Aggchainecdsa.Contract.GetAggchainVKey(&_Aggchainecdsa.CallOpts, aggchainVKeySelector)
}

// GetAggchainVKey is a free data retrieval call binding the contract method 0x01fcf6a0.
//
// Solidity: function getAggchainVKey(bytes4 aggchainVKeySelector) view returns(bytes32 aggchainVKey)
func (_Aggchainecdsa *AggchainecdsaCallerSession) GetAggchainVKey(aggchainVKeySelector [4]byte) ([32]byte, error) {
	return _Aggchainecdsa.Contract.GetAggchainVKey(&_Aggchainecdsa.CallOpts, aggchainVKeySelector)
}

// GetAggchainVKeySelector is a free data retrieval call binding the contract method 0x1d0b435e.
//
// Solidity: function getAggchainVKeySelector(bytes2 aggchainVKeyVersion, bytes2 aggchainType) pure returns(bytes4)
func (_Aggchainecdsa *AggchainecdsaCaller) GetAggchainVKeySelector(opts *bind.CallOpts, aggchainVKeyVersion [2]byte, aggchainType [2]byte) ([4]byte, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "getAggchainVKeySelector", aggchainVKeyVersion, aggchainType)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// GetAggchainVKeySelector is a free data retrieval call binding the contract method 0x1d0b435e.
//
// Solidity: function getAggchainVKeySelector(bytes2 aggchainVKeyVersion, bytes2 aggchainType) pure returns(bytes4)
func (_Aggchainecdsa *AggchainecdsaSession) GetAggchainVKeySelector(aggchainVKeyVersion [2]byte, aggchainType [2]byte) ([4]byte, error) {
	return _Aggchainecdsa.Contract.GetAggchainVKeySelector(&_Aggchainecdsa.CallOpts, aggchainVKeyVersion, aggchainType)
}

// GetAggchainVKeySelector is a free data retrieval call binding the contract method 0x1d0b435e.
//
// Solidity: function getAggchainVKeySelector(bytes2 aggchainVKeyVersion, bytes2 aggchainType) pure returns(bytes4)
func (_Aggchainecdsa *AggchainecdsaCallerSession) GetAggchainVKeySelector(aggchainVKeyVersion [2]byte, aggchainType [2]byte) ([4]byte, error) {
	return _Aggchainecdsa.Contract.GetAggchainVKeySelector(&_Aggchainecdsa.CallOpts, aggchainVKeyVersion, aggchainType)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Aggchainecdsa *AggchainecdsaCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Aggchainecdsa *AggchainecdsaSession) GlobalExitRootManager() (common.Address, error) {
	return _Aggchainecdsa.Contract.GlobalExitRootManager(&_Aggchainecdsa.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Aggchainecdsa *AggchainecdsaCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Aggchainecdsa.Contract.GlobalExitRootManager(&_Aggchainecdsa.CallOpts)
}

// Initialize0 is a free data retrieval call binding the contract method 0x71257022.
//
// Solidity: function initialize(address , address , uint32 , address , string , string ) pure returns()
func (_Aggchainecdsa *AggchainecdsaCaller) Initialize0(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 string, arg5 string) error {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "initialize0", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return err
	}

	return err

}

// Initialize0 is a free data retrieval call binding the contract method 0x71257022.
//
// Solidity: function initialize(address , address , uint32 , address , string , string ) pure returns()
func (_Aggchainecdsa *AggchainecdsaSession) Initialize0(arg0 common.Address, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 string, arg5 string) error {
	return _Aggchainecdsa.Contract.Initialize0(&_Aggchainecdsa.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// Initialize0 is a free data retrieval call binding the contract method 0x71257022.
//
// Solidity: function initialize(address , address , uint32 , address , string , string ) pure returns()
func (_Aggchainecdsa *AggchainecdsaCallerSession) Initialize0(arg0 common.Address, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 string, arg5 string) error {
	return _Aggchainecdsa.Contract.Initialize0(&_Aggchainecdsa.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Aggchainecdsa *AggchainecdsaCaller) LastAccInputHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "lastAccInputHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Aggchainecdsa *AggchainecdsaSession) LastAccInputHash() ([32]byte, error) {
	return _Aggchainecdsa.Contract.LastAccInputHash(&_Aggchainecdsa.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Aggchainecdsa *AggchainecdsaCallerSession) LastAccInputHash() ([32]byte, error) {
	return _Aggchainecdsa.Contract.LastAccInputHash(&_Aggchainecdsa.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Aggchainecdsa *AggchainecdsaCaller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Aggchainecdsa *AggchainecdsaSession) LastForceBatch() (uint64, error) {
	return _Aggchainecdsa.Contract.LastForceBatch(&_Aggchainecdsa.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Aggchainecdsa *AggchainecdsaCallerSession) LastForceBatch() (uint64, error) {
	return _Aggchainecdsa.Contract.LastForceBatch(&_Aggchainecdsa.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Aggchainecdsa *AggchainecdsaCaller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Aggchainecdsa *AggchainecdsaSession) LastForceBatchSequenced() (uint64, error) {
	return _Aggchainecdsa.Contract.LastForceBatchSequenced(&_Aggchainecdsa.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Aggchainecdsa *AggchainecdsaCallerSession) LastForceBatchSequenced() (uint64, error) {
	return _Aggchainecdsa.Contract.LastForceBatchSequenced(&_Aggchainecdsa.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Aggchainecdsa *AggchainecdsaCaller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Aggchainecdsa *AggchainecdsaSession) NetworkName() (string, error) {
	return _Aggchainecdsa.Contract.NetworkName(&_Aggchainecdsa.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Aggchainecdsa *AggchainecdsaCallerSession) NetworkName() (string, error) {
	return _Aggchainecdsa.Contract.NetworkName(&_Aggchainecdsa.CallOpts)
}

// OwnedAggchainVKeys is a free data retrieval call binding the contract method 0xeffb8479.
//
// Solidity: function ownedAggchainVKeys(bytes4 aggchainVKeySelector) view returns(bytes32 ownedAggchainVKey)
func (_Aggchainecdsa *AggchainecdsaCaller) OwnedAggchainVKeys(opts *bind.CallOpts, aggchainVKeySelector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "ownedAggchainVKeys", aggchainVKeySelector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OwnedAggchainVKeys is a free data retrieval call binding the contract method 0xeffb8479.
//
// Solidity: function ownedAggchainVKeys(bytes4 aggchainVKeySelector) view returns(bytes32 ownedAggchainVKey)
func (_Aggchainecdsa *AggchainecdsaSession) OwnedAggchainVKeys(aggchainVKeySelector [4]byte) ([32]byte, error) {
	return _Aggchainecdsa.Contract.OwnedAggchainVKeys(&_Aggchainecdsa.CallOpts, aggchainVKeySelector)
}

// OwnedAggchainVKeys is a free data retrieval call binding the contract method 0xeffb8479.
//
// Solidity: function ownedAggchainVKeys(bytes4 aggchainVKeySelector) view returns(bytes32 ownedAggchainVKey)
func (_Aggchainecdsa *AggchainecdsaCallerSession) OwnedAggchainVKeys(aggchainVKeySelector [4]byte) ([32]byte, error) {
	return _Aggchainecdsa.Contract.OwnedAggchainVKeys(&_Aggchainecdsa.CallOpts, aggchainVKeySelector)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Aggchainecdsa *AggchainecdsaCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Aggchainecdsa *AggchainecdsaSession) PendingAdmin() (common.Address, error) {
	return _Aggchainecdsa.Contract.PendingAdmin(&_Aggchainecdsa.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Aggchainecdsa *AggchainecdsaCallerSession) PendingAdmin() (common.Address, error) {
	return _Aggchainecdsa.Contract.PendingAdmin(&_Aggchainecdsa.CallOpts)
}

// PendingVKeyManager is a free data retrieval call binding the contract method 0xbfb193b6.
//
// Solidity: function pendingVKeyManager() view returns(address)
func (_Aggchainecdsa *AggchainecdsaCaller) PendingVKeyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "pendingVKeyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingVKeyManager is a free data retrieval call binding the contract method 0xbfb193b6.
//
// Solidity: function pendingVKeyManager() view returns(address)
func (_Aggchainecdsa *AggchainecdsaSession) PendingVKeyManager() (common.Address, error) {
	return _Aggchainecdsa.Contract.PendingVKeyManager(&_Aggchainecdsa.CallOpts)
}

// PendingVKeyManager is a free data retrieval call binding the contract method 0xbfb193b6.
//
// Solidity: function pendingVKeyManager() view returns(address)
func (_Aggchainecdsa *AggchainecdsaCallerSession) PendingVKeyManager() (common.Address, error) {
	return _Aggchainecdsa.Contract.PendingVKeyManager(&_Aggchainecdsa.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Aggchainecdsa *AggchainecdsaCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Aggchainecdsa *AggchainecdsaSession) Pol() (common.Address, error) {
	return _Aggchainecdsa.Contract.Pol(&_Aggchainecdsa.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Aggchainecdsa *AggchainecdsaCallerSession) Pol() (common.Address, error) {
	return _Aggchainecdsa.Contract.Pol(&_Aggchainecdsa.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Aggchainecdsa *AggchainecdsaCaller) RollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "rollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Aggchainecdsa *AggchainecdsaSession) RollupManager() (common.Address, error) {
	return _Aggchainecdsa.Contract.RollupManager(&_Aggchainecdsa.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Aggchainecdsa *AggchainecdsaCallerSession) RollupManager() (common.Address, error) {
	return _Aggchainecdsa.Contract.RollupManager(&_Aggchainecdsa.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Aggchainecdsa *AggchainecdsaCaller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Aggchainecdsa *AggchainecdsaSession) TrustedSequencer() (common.Address, error) {
	return _Aggchainecdsa.Contract.TrustedSequencer(&_Aggchainecdsa.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Aggchainecdsa *AggchainecdsaCallerSession) TrustedSequencer() (common.Address, error) {
	return _Aggchainecdsa.Contract.TrustedSequencer(&_Aggchainecdsa.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Aggchainecdsa *AggchainecdsaCaller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Aggchainecdsa *AggchainecdsaSession) TrustedSequencerURL() (string, error) {
	return _Aggchainecdsa.Contract.TrustedSequencerURL(&_Aggchainecdsa.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Aggchainecdsa *AggchainecdsaCallerSession) TrustedSequencerURL() (string, error) {
	return _Aggchainecdsa.Contract.TrustedSequencerURL(&_Aggchainecdsa.CallOpts)
}

// UseDefaultGateway is a free data retrieval call binding the contract method 0xff904079.
//
// Solidity: function useDefaultGateway() view returns(bool)
func (_Aggchainecdsa *AggchainecdsaCaller) UseDefaultGateway(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "useDefaultGateway")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseDefaultGateway is a free data retrieval call binding the contract method 0xff904079.
//
// Solidity: function useDefaultGateway() view returns(bool)
func (_Aggchainecdsa *AggchainecdsaSession) UseDefaultGateway() (bool, error) {
	return _Aggchainecdsa.Contract.UseDefaultGateway(&_Aggchainecdsa.CallOpts)
}

// UseDefaultGateway is a free data retrieval call binding the contract method 0xff904079.
//
// Solidity: function useDefaultGateway() view returns(bool)
func (_Aggchainecdsa *AggchainecdsaCallerSession) UseDefaultGateway() (bool, error) {
	return _Aggchainecdsa.Contract.UseDefaultGateway(&_Aggchainecdsa.CallOpts)
}

// VKeyManager is a free data retrieval call binding the contract method 0xe279984e.
//
// Solidity: function vKeyManager() view returns(address)
func (_Aggchainecdsa *AggchainecdsaCaller) VKeyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "vKeyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VKeyManager is a free data retrieval call binding the contract method 0xe279984e.
//
// Solidity: function vKeyManager() view returns(address)
func (_Aggchainecdsa *AggchainecdsaSession) VKeyManager() (common.Address, error) {
	return _Aggchainecdsa.Contract.VKeyManager(&_Aggchainecdsa.CallOpts)
}

// VKeyManager is a free data retrieval call binding the contract method 0xe279984e.
//
// Solidity: function vKeyManager() view returns(address)
func (_Aggchainecdsa *AggchainecdsaCallerSession) VKeyManager() (common.Address, error) {
	return _Aggchainecdsa.Contract.VKeyManager(&_Aggchainecdsa.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Aggchainecdsa *AggchainecdsaTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainecdsa.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Aggchainecdsa *AggchainecdsaSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.AcceptAdminRole(&_Aggchainecdsa.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Aggchainecdsa *AggchainecdsaTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.AcceptAdminRole(&_Aggchainecdsa.TransactOpts)
}

// AcceptVKeyManagerRole is a paid mutator transaction binding the contract method 0x368c822c.
//
// Solidity: function acceptVKeyManagerRole() returns()
func (_Aggchainecdsa *AggchainecdsaTransactor) AcceptVKeyManagerRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainecdsa.contract.Transact(opts, "acceptVKeyManagerRole")
}

// AcceptVKeyManagerRole is a paid mutator transaction binding the contract method 0x368c822c.
//
// Solidity: function acceptVKeyManagerRole() returns()
func (_Aggchainecdsa *AggchainecdsaSession) AcceptVKeyManagerRole() (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.AcceptVKeyManagerRole(&_Aggchainecdsa.TransactOpts)
}

// AcceptVKeyManagerRole is a paid mutator transaction binding the contract method 0x368c822c.
//
// Solidity: function acceptVKeyManagerRole() returns()
func (_Aggchainecdsa *AggchainecdsaTransactorSession) AcceptVKeyManagerRole() (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.AcceptVKeyManagerRole(&_Aggchainecdsa.TransactOpts)
}

// AddOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x19451a8f.
//
// Solidity: function addOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 newAggchainVKey) returns()
func (_Aggchainecdsa *AggchainecdsaTransactor) AddOwnedAggchainVKey(opts *bind.TransactOpts, aggchainVKeySelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainecdsa.contract.Transact(opts, "addOwnedAggchainVKey", aggchainVKeySelector, newAggchainVKey)
}

// AddOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x19451a8f.
//
// Solidity: function addOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 newAggchainVKey) returns()
func (_Aggchainecdsa *AggchainecdsaSession) AddOwnedAggchainVKey(aggchainVKeySelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.AddOwnedAggchainVKey(&_Aggchainecdsa.TransactOpts, aggchainVKeySelector, newAggchainVKey)
}

// AddOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x19451a8f.
//
// Solidity: function addOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 newAggchainVKey) returns()
func (_Aggchainecdsa *AggchainecdsaTransactorSession) AddOwnedAggchainVKey(aggchainVKeySelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.AddOwnedAggchainVKey(&_Aggchainecdsa.TransactOpts, aggchainVKeySelector, newAggchainVKey)
}

// DisableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xdc8c4249.
//
// Solidity: function disableUseDefaultGatewayFlag() returns()
func (_Aggchainecdsa *AggchainecdsaTransactor) DisableUseDefaultGatewayFlag(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainecdsa.contract.Transact(opts, "disableUseDefaultGatewayFlag")
}

// DisableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xdc8c4249.
//
// Solidity: function disableUseDefaultGatewayFlag() returns()
func (_Aggchainecdsa *AggchainecdsaSession) DisableUseDefaultGatewayFlag() (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.DisableUseDefaultGatewayFlag(&_Aggchainecdsa.TransactOpts)
}

// DisableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xdc8c4249.
//
// Solidity: function disableUseDefaultGatewayFlag() returns()
func (_Aggchainecdsa *AggchainecdsaTransactorSession) DisableUseDefaultGatewayFlag() (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.DisableUseDefaultGatewayFlag(&_Aggchainecdsa.TransactOpts)
}

// EnableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xe631476c.
//
// Solidity: function enableUseDefaultGatewayFlag() returns()
func (_Aggchainecdsa *AggchainecdsaTransactor) EnableUseDefaultGatewayFlag(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainecdsa.contract.Transact(opts, "enableUseDefaultGatewayFlag")
}

// EnableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xe631476c.
//
// Solidity: function enableUseDefaultGatewayFlag() returns()
func (_Aggchainecdsa *AggchainecdsaSession) EnableUseDefaultGatewayFlag() (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.EnableUseDefaultGatewayFlag(&_Aggchainecdsa.TransactOpts)
}

// EnableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xe631476c.
//
// Solidity: function enableUseDefaultGatewayFlag() returns()
func (_Aggchainecdsa *AggchainecdsaTransactorSession) EnableUseDefaultGatewayFlag() (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.EnableUseDefaultGatewayFlag(&_Aggchainecdsa.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializeBytesAggchain) returns()
func (_Aggchainecdsa *AggchainecdsaTransactor) Initialize(opts *bind.TransactOpts, initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Aggchainecdsa.contract.Transact(opts, "initialize", initializeBytesAggchain)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializeBytesAggchain) returns()
func (_Aggchainecdsa *AggchainecdsaSession) Initialize(initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.Initialize(&_Aggchainecdsa.TransactOpts, initializeBytesAggchain)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializeBytesAggchain) returns()
func (_Aggchainecdsa *AggchainecdsaTransactorSession) Initialize(initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.Initialize(&_Aggchainecdsa.TransactOpts, initializeBytesAggchain)
}

// OnVerifyPessimistic is a paid mutator transaction binding the contract method 0x9ee4afa3.
//
// Solidity: function onVerifyPessimistic(bytes aggchainData) returns()
func (_Aggchainecdsa *AggchainecdsaTransactor) OnVerifyPessimistic(opts *bind.TransactOpts, aggchainData []byte) (*types.Transaction, error) {
	return _Aggchainecdsa.contract.Transact(opts, "onVerifyPessimistic", aggchainData)
}

// OnVerifyPessimistic is a paid mutator transaction binding the contract method 0x9ee4afa3.
//
// Solidity: function onVerifyPessimistic(bytes aggchainData) returns()
func (_Aggchainecdsa *AggchainecdsaSession) OnVerifyPessimistic(aggchainData []byte) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.OnVerifyPessimistic(&_Aggchainecdsa.TransactOpts, aggchainData)
}

// OnVerifyPessimistic is a paid mutator transaction binding the contract method 0x9ee4afa3.
//
// Solidity: function onVerifyPessimistic(bytes aggchainData) returns()
func (_Aggchainecdsa *AggchainecdsaTransactorSession) OnVerifyPessimistic(aggchainData []byte) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.OnVerifyPessimistic(&_Aggchainecdsa.TransactOpts, aggchainData)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Aggchainecdsa *AggchainecdsaTransactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Aggchainecdsa.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Aggchainecdsa *AggchainecdsaSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.SetTrustedSequencer(&_Aggchainecdsa.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Aggchainecdsa *AggchainecdsaTransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.SetTrustedSequencer(&_Aggchainecdsa.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Aggchainecdsa *AggchainecdsaTransactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Aggchainecdsa.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Aggchainecdsa *AggchainecdsaSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.SetTrustedSequencerURL(&_Aggchainecdsa.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Aggchainecdsa *AggchainecdsaTransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.SetTrustedSequencerURL(&_Aggchainecdsa.TransactOpts, newTrustedSequencerURL)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Aggchainecdsa *AggchainecdsaTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Aggchainecdsa.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Aggchainecdsa *AggchainecdsaSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.TransferAdminRole(&_Aggchainecdsa.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Aggchainecdsa *AggchainecdsaTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.TransferAdminRole(&_Aggchainecdsa.TransactOpts, newPendingAdmin)
}

// TransferVKeyManagerRole is a paid mutator transaction binding the contract method 0x85018182.
//
// Solidity: function transferVKeyManagerRole(address newVKeyManager) returns()
func (_Aggchainecdsa *AggchainecdsaTransactor) TransferVKeyManagerRole(opts *bind.TransactOpts, newVKeyManager common.Address) (*types.Transaction, error) {
	return _Aggchainecdsa.contract.Transact(opts, "transferVKeyManagerRole", newVKeyManager)
}

// TransferVKeyManagerRole is a paid mutator transaction binding the contract method 0x85018182.
//
// Solidity: function transferVKeyManagerRole(address newVKeyManager) returns()
func (_Aggchainecdsa *AggchainecdsaSession) TransferVKeyManagerRole(newVKeyManager common.Address) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.TransferVKeyManagerRole(&_Aggchainecdsa.TransactOpts, newVKeyManager)
}

// TransferVKeyManagerRole is a paid mutator transaction binding the contract method 0x85018182.
//
// Solidity: function transferVKeyManagerRole(address newVKeyManager) returns()
func (_Aggchainecdsa *AggchainecdsaTransactorSession) TransferVKeyManagerRole(newVKeyManager common.Address) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.TransferVKeyManagerRole(&_Aggchainecdsa.TransactOpts, newVKeyManager)
}

// UpdateOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x314eb17b.
//
// Solidity: function updateOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 updatedAggchainVKey) returns()
func (_Aggchainecdsa *AggchainecdsaTransactor) UpdateOwnedAggchainVKey(opts *bind.TransactOpts, aggchainVKeySelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainecdsa.contract.Transact(opts, "updateOwnedAggchainVKey", aggchainVKeySelector, updatedAggchainVKey)
}

// UpdateOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x314eb17b.
//
// Solidity: function updateOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 updatedAggchainVKey) returns()
func (_Aggchainecdsa *AggchainecdsaSession) UpdateOwnedAggchainVKey(aggchainVKeySelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.UpdateOwnedAggchainVKey(&_Aggchainecdsa.TransactOpts, aggchainVKeySelector, updatedAggchainVKey)
}

// UpdateOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x314eb17b.
//
// Solidity: function updateOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 updatedAggchainVKey) returns()
func (_Aggchainecdsa *AggchainecdsaTransactorSession) UpdateOwnedAggchainVKey(aggchainVKeySelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.UpdateOwnedAggchainVKey(&_Aggchainecdsa.TransactOpts, aggchainVKeySelector, updatedAggchainVKey)
}

// AggchainecdsaAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Aggchainecdsa contract.
type AggchainecdsaAcceptAdminRoleIterator struct {
	Event *AggchainecdsaAcceptAdminRole // Event containing the contract specifics and raw log

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
func (it *AggchainecdsaAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsaAcceptAdminRole)
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
		it.Event = new(AggchainecdsaAcceptAdminRole)
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
func (it *AggchainecdsaAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsaAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsaAcceptAdminRole represents a AcceptAdminRole event raised by the Aggchainecdsa contract.
type AggchainecdsaAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Aggchainecdsa *AggchainecdsaFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*AggchainecdsaAcceptAdminRoleIterator, error) {

	logs, sub, err := _Aggchainecdsa.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsaAcceptAdminRoleIterator{contract: _Aggchainecdsa.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Aggchainecdsa *AggchainecdsaFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *AggchainecdsaAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsa.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsaAcceptAdminRole)
				if err := _Aggchainecdsa.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
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
func (_Aggchainecdsa *AggchainecdsaFilterer) ParseAcceptAdminRole(log types.Log) (*AggchainecdsaAcceptAdminRole, error) {
	event := new(AggchainecdsaAcceptAdminRole)
	if err := _Aggchainecdsa.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsaAcceptVKeyManagerRoleIterator is returned from FilterAcceptVKeyManagerRole and is used to iterate over the raw logs and unpacked data for AcceptVKeyManagerRole events raised by the Aggchainecdsa contract.
type AggchainecdsaAcceptVKeyManagerRoleIterator struct {
	Event *AggchainecdsaAcceptVKeyManagerRole // Event containing the contract specifics and raw log

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
func (it *AggchainecdsaAcceptVKeyManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsaAcceptVKeyManagerRole)
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
		it.Event = new(AggchainecdsaAcceptVKeyManagerRole)
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
func (it *AggchainecdsaAcceptVKeyManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsaAcceptVKeyManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsaAcceptVKeyManagerRole represents a AcceptVKeyManagerRole event raised by the Aggchainecdsa contract.
type AggchainecdsaAcceptVKeyManagerRole struct {
	OldVKeyManager common.Address
	NewVKeyManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAcceptVKeyManagerRole is a free log retrieval operation binding the contract event 0xbb3b066505f14a628f4ba4187a046abd4dd17e96591d7a9ed31c91c79322ffe2.
//
// Solidity: event AcceptVKeyManagerRole(address oldVKeyManager, address newVKeyManager)
func (_Aggchainecdsa *AggchainecdsaFilterer) FilterAcceptVKeyManagerRole(opts *bind.FilterOpts) (*AggchainecdsaAcceptVKeyManagerRoleIterator, error) {

	logs, sub, err := _Aggchainecdsa.contract.FilterLogs(opts, "AcceptVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsaAcceptVKeyManagerRoleIterator{contract: _Aggchainecdsa.contract, event: "AcceptVKeyManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptVKeyManagerRole is a free log subscription operation binding the contract event 0xbb3b066505f14a628f4ba4187a046abd4dd17e96591d7a9ed31c91c79322ffe2.
//
// Solidity: event AcceptVKeyManagerRole(address oldVKeyManager, address newVKeyManager)
func (_Aggchainecdsa *AggchainecdsaFilterer) WatchAcceptVKeyManagerRole(opts *bind.WatchOpts, sink chan<- *AggchainecdsaAcceptVKeyManagerRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsa.contract.WatchLogs(opts, "AcceptVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsaAcceptVKeyManagerRole)
				if err := _Aggchainecdsa.contract.UnpackLog(event, "AcceptVKeyManagerRole", log); err != nil {
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
func (_Aggchainecdsa *AggchainecdsaFilterer) ParseAcceptVKeyManagerRole(log types.Log) (*AggchainecdsaAcceptVKeyManagerRole, error) {
	event := new(AggchainecdsaAcceptVKeyManagerRole)
	if err := _Aggchainecdsa.contract.UnpackLog(event, "AcceptVKeyManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsaAddAggchainVKeyIterator is returned from FilterAddAggchainVKey and is used to iterate over the raw logs and unpacked data for AddAggchainVKey events raised by the Aggchainecdsa contract.
type AggchainecdsaAddAggchainVKeyIterator struct {
	Event *AggchainecdsaAddAggchainVKey // Event containing the contract specifics and raw log

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
func (it *AggchainecdsaAddAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsaAddAggchainVKey)
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
		it.Event = new(AggchainecdsaAddAggchainVKey)
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
func (it *AggchainecdsaAddAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsaAddAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsaAddAggchainVKey represents a AddAggchainVKey event raised by the Aggchainecdsa contract.
type AggchainecdsaAddAggchainVKey struct {
	Selector        [4]byte
	NewAggchainVKey [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddAggchainVKey is a free log retrieval operation binding the contract event 0x6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a79.
//
// Solidity: event AddAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Aggchainecdsa *AggchainecdsaFilterer) FilterAddAggchainVKey(opts *bind.FilterOpts) (*AggchainecdsaAddAggchainVKeyIterator, error) {

	logs, sub, err := _Aggchainecdsa.contract.FilterLogs(opts, "AddAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsaAddAggchainVKeyIterator{contract: _Aggchainecdsa.contract, event: "AddAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchAddAggchainVKey is a free log subscription operation binding the contract event 0x6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a79.
//
// Solidity: event AddAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Aggchainecdsa *AggchainecdsaFilterer) WatchAddAggchainVKey(opts *bind.WatchOpts, sink chan<- *AggchainecdsaAddAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsa.contract.WatchLogs(opts, "AddAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsaAddAggchainVKey)
				if err := _Aggchainecdsa.contract.UnpackLog(event, "AddAggchainVKey", log); err != nil {
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
func (_Aggchainecdsa *AggchainecdsaFilterer) ParseAddAggchainVKey(log types.Log) (*AggchainecdsaAddAggchainVKey, error) {
	event := new(AggchainecdsaAddAggchainVKey)
	if err := _Aggchainecdsa.contract.UnpackLog(event, "AddAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsaDisableUseDefaultGatewayFlagIterator is returned from FilterDisableUseDefaultGatewayFlag and is used to iterate over the raw logs and unpacked data for DisableUseDefaultGatewayFlag events raised by the Aggchainecdsa contract.
type AggchainecdsaDisableUseDefaultGatewayFlagIterator struct {
	Event *AggchainecdsaDisableUseDefaultGatewayFlag // Event containing the contract specifics and raw log

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
func (it *AggchainecdsaDisableUseDefaultGatewayFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsaDisableUseDefaultGatewayFlag)
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
		it.Event = new(AggchainecdsaDisableUseDefaultGatewayFlag)
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
func (it *AggchainecdsaDisableUseDefaultGatewayFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsaDisableUseDefaultGatewayFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsaDisableUseDefaultGatewayFlag represents a DisableUseDefaultGatewayFlag event raised by the Aggchainecdsa contract.
type AggchainecdsaDisableUseDefaultGatewayFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisableUseDefaultGatewayFlag is a free log retrieval operation binding the contract event 0x67dd1717a1952df380cb73eecb312e949df6d6a086bd7f88669005341972528e.
//
// Solidity: event DisableUseDefaultGatewayFlag()
func (_Aggchainecdsa *AggchainecdsaFilterer) FilterDisableUseDefaultGatewayFlag(opts *bind.FilterOpts) (*AggchainecdsaDisableUseDefaultGatewayFlagIterator, error) {

	logs, sub, err := _Aggchainecdsa.contract.FilterLogs(opts, "DisableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsaDisableUseDefaultGatewayFlagIterator{contract: _Aggchainecdsa.contract, event: "DisableUseDefaultGatewayFlag", logs: logs, sub: sub}, nil
}

// WatchDisableUseDefaultGatewayFlag is a free log subscription operation binding the contract event 0x67dd1717a1952df380cb73eecb312e949df6d6a086bd7f88669005341972528e.
//
// Solidity: event DisableUseDefaultGatewayFlag()
func (_Aggchainecdsa *AggchainecdsaFilterer) WatchDisableUseDefaultGatewayFlag(opts *bind.WatchOpts, sink chan<- *AggchainecdsaDisableUseDefaultGatewayFlag) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsa.contract.WatchLogs(opts, "DisableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsaDisableUseDefaultGatewayFlag)
				if err := _Aggchainecdsa.contract.UnpackLog(event, "DisableUseDefaultGatewayFlag", log); err != nil {
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
func (_Aggchainecdsa *AggchainecdsaFilterer) ParseDisableUseDefaultGatewayFlag(log types.Log) (*AggchainecdsaDisableUseDefaultGatewayFlag, error) {
	event := new(AggchainecdsaDisableUseDefaultGatewayFlag)
	if err := _Aggchainecdsa.contract.UnpackLog(event, "DisableUseDefaultGatewayFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsaEnableUseDefaultGatewayFlagIterator is returned from FilterEnableUseDefaultGatewayFlag and is used to iterate over the raw logs and unpacked data for EnableUseDefaultGatewayFlag events raised by the Aggchainecdsa contract.
type AggchainecdsaEnableUseDefaultGatewayFlagIterator struct {
	Event *AggchainecdsaEnableUseDefaultGatewayFlag // Event containing the contract specifics and raw log

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
func (it *AggchainecdsaEnableUseDefaultGatewayFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsaEnableUseDefaultGatewayFlag)
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
		it.Event = new(AggchainecdsaEnableUseDefaultGatewayFlag)
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
func (it *AggchainecdsaEnableUseDefaultGatewayFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsaEnableUseDefaultGatewayFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsaEnableUseDefaultGatewayFlag represents a EnableUseDefaultGatewayFlag event raised by the Aggchainecdsa contract.
type AggchainecdsaEnableUseDefaultGatewayFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEnableUseDefaultGatewayFlag is a free log retrieval operation binding the contract event 0xb6563aed80fde357e737eb0d19f246a58cb6bfd469933d05701ecbad0f2dca84.
//
// Solidity: event EnableUseDefaultGatewayFlag()
func (_Aggchainecdsa *AggchainecdsaFilterer) FilterEnableUseDefaultGatewayFlag(opts *bind.FilterOpts) (*AggchainecdsaEnableUseDefaultGatewayFlagIterator, error) {

	logs, sub, err := _Aggchainecdsa.contract.FilterLogs(opts, "EnableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsaEnableUseDefaultGatewayFlagIterator{contract: _Aggchainecdsa.contract, event: "EnableUseDefaultGatewayFlag", logs: logs, sub: sub}, nil
}

// WatchEnableUseDefaultGatewayFlag is a free log subscription operation binding the contract event 0xb6563aed80fde357e737eb0d19f246a58cb6bfd469933d05701ecbad0f2dca84.
//
// Solidity: event EnableUseDefaultGatewayFlag()
func (_Aggchainecdsa *AggchainecdsaFilterer) WatchEnableUseDefaultGatewayFlag(opts *bind.WatchOpts, sink chan<- *AggchainecdsaEnableUseDefaultGatewayFlag) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsa.contract.WatchLogs(opts, "EnableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsaEnableUseDefaultGatewayFlag)
				if err := _Aggchainecdsa.contract.UnpackLog(event, "EnableUseDefaultGatewayFlag", log); err != nil {
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
func (_Aggchainecdsa *AggchainecdsaFilterer) ParseEnableUseDefaultGatewayFlag(log types.Log) (*AggchainecdsaEnableUseDefaultGatewayFlag, error) {
	event := new(AggchainecdsaEnableUseDefaultGatewayFlag)
	if err := _Aggchainecdsa.contract.UnpackLog(event, "EnableUseDefaultGatewayFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsaInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Aggchainecdsa contract.
type AggchainecdsaInitializedIterator struct {
	Event *AggchainecdsaInitialized // Event containing the contract specifics and raw log

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
func (it *AggchainecdsaInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsaInitialized)
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
		it.Event = new(AggchainecdsaInitialized)
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
func (it *AggchainecdsaInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsaInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsaInitialized represents a Initialized event raised by the Aggchainecdsa contract.
type AggchainecdsaInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Aggchainecdsa *AggchainecdsaFilterer) FilterInitialized(opts *bind.FilterOpts) (*AggchainecdsaInitializedIterator, error) {

	logs, sub, err := _Aggchainecdsa.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsaInitializedIterator{contract: _Aggchainecdsa.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Aggchainecdsa *AggchainecdsaFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AggchainecdsaInitialized) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsa.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsaInitialized)
				if err := _Aggchainecdsa.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Aggchainecdsa *AggchainecdsaFilterer) ParseInitialized(log types.Log) (*AggchainecdsaInitialized, error) {
	event := new(AggchainecdsaInitialized)
	if err := _Aggchainecdsa.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsaOnVerifyPessimisticECDSAIterator is returned from FilterOnVerifyPessimisticECDSA and is used to iterate over the raw logs and unpacked data for OnVerifyPessimisticECDSA events raised by the Aggchainecdsa contract.
type AggchainecdsaOnVerifyPessimisticECDSAIterator struct {
	Event *AggchainecdsaOnVerifyPessimisticECDSA // Event containing the contract specifics and raw log

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
func (it *AggchainecdsaOnVerifyPessimisticECDSAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsaOnVerifyPessimisticECDSA)
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
		it.Event = new(AggchainecdsaOnVerifyPessimisticECDSA)
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
func (it *AggchainecdsaOnVerifyPessimisticECDSAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsaOnVerifyPessimisticECDSAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsaOnVerifyPessimisticECDSA represents a OnVerifyPessimisticECDSA event raised by the Aggchainecdsa contract.
type AggchainecdsaOnVerifyPessimisticECDSA struct {
	NewStateRoot [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOnVerifyPessimisticECDSA is a free log retrieval operation binding the contract event 0xffc88430d30048515bd904f24ae1b72af4b78151ac346c6de9271ff5ab7899b4.
//
// Solidity: event OnVerifyPessimisticECDSA(bytes32 newStateRoot)
func (_Aggchainecdsa *AggchainecdsaFilterer) FilterOnVerifyPessimisticECDSA(opts *bind.FilterOpts) (*AggchainecdsaOnVerifyPessimisticECDSAIterator, error) {

	logs, sub, err := _Aggchainecdsa.contract.FilterLogs(opts, "OnVerifyPessimisticECDSA")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsaOnVerifyPessimisticECDSAIterator{contract: _Aggchainecdsa.contract, event: "OnVerifyPessimisticECDSA", logs: logs, sub: sub}, nil
}

// WatchOnVerifyPessimisticECDSA is a free log subscription operation binding the contract event 0xffc88430d30048515bd904f24ae1b72af4b78151ac346c6de9271ff5ab7899b4.
//
// Solidity: event OnVerifyPessimisticECDSA(bytes32 newStateRoot)
func (_Aggchainecdsa *AggchainecdsaFilterer) WatchOnVerifyPessimisticECDSA(opts *bind.WatchOpts, sink chan<- *AggchainecdsaOnVerifyPessimisticECDSA) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsa.contract.WatchLogs(opts, "OnVerifyPessimisticECDSA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsaOnVerifyPessimisticECDSA)
				if err := _Aggchainecdsa.contract.UnpackLog(event, "OnVerifyPessimisticECDSA", log); err != nil {
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

// ParseOnVerifyPessimisticECDSA is a log parse operation binding the contract event 0xffc88430d30048515bd904f24ae1b72af4b78151ac346c6de9271ff5ab7899b4.
//
// Solidity: event OnVerifyPessimisticECDSA(bytes32 newStateRoot)
func (_Aggchainecdsa *AggchainecdsaFilterer) ParseOnVerifyPessimisticECDSA(log types.Log) (*AggchainecdsaOnVerifyPessimisticECDSA, error) {
	event := new(AggchainecdsaOnVerifyPessimisticECDSA)
	if err := _Aggchainecdsa.contract.UnpackLog(event, "OnVerifyPessimisticECDSA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsaSetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the Aggchainecdsa contract.
type AggchainecdsaSetTrustedSequencerIterator struct {
	Event *AggchainecdsaSetTrustedSequencer // Event containing the contract specifics and raw log

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
func (it *AggchainecdsaSetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsaSetTrustedSequencer)
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
		it.Event = new(AggchainecdsaSetTrustedSequencer)
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
func (it *AggchainecdsaSetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsaSetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsaSetTrustedSequencer represents a SetTrustedSequencer event raised by the Aggchainecdsa contract.
type AggchainecdsaSetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Aggchainecdsa *AggchainecdsaFilterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*AggchainecdsaSetTrustedSequencerIterator, error) {

	logs, sub, err := _Aggchainecdsa.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsaSetTrustedSequencerIterator{contract: _Aggchainecdsa.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Aggchainecdsa *AggchainecdsaFilterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *AggchainecdsaSetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsa.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsaSetTrustedSequencer)
				if err := _Aggchainecdsa.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
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
func (_Aggchainecdsa *AggchainecdsaFilterer) ParseSetTrustedSequencer(log types.Log) (*AggchainecdsaSetTrustedSequencer, error) {
	event := new(AggchainecdsaSetTrustedSequencer)
	if err := _Aggchainecdsa.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsaSetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the Aggchainecdsa contract.
type AggchainecdsaSetTrustedSequencerURLIterator struct {
	Event *AggchainecdsaSetTrustedSequencerURL // Event containing the contract specifics and raw log

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
func (it *AggchainecdsaSetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsaSetTrustedSequencerURL)
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
		it.Event = new(AggchainecdsaSetTrustedSequencerURL)
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
func (it *AggchainecdsaSetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsaSetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsaSetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the Aggchainecdsa contract.
type AggchainecdsaSetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Aggchainecdsa *AggchainecdsaFilterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*AggchainecdsaSetTrustedSequencerURLIterator, error) {

	logs, sub, err := _Aggchainecdsa.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsaSetTrustedSequencerURLIterator{contract: _Aggchainecdsa.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Aggchainecdsa *AggchainecdsaFilterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *AggchainecdsaSetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsa.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsaSetTrustedSequencerURL)
				if err := _Aggchainecdsa.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
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
func (_Aggchainecdsa *AggchainecdsaFilterer) ParseSetTrustedSequencerURL(log types.Log) (*AggchainecdsaSetTrustedSequencerURL, error) {
	event := new(AggchainecdsaSetTrustedSequencerURL)
	if err := _Aggchainecdsa.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsaTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Aggchainecdsa contract.
type AggchainecdsaTransferAdminRoleIterator struct {
	Event *AggchainecdsaTransferAdminRole // Event containing the contract specifics and raw log

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
func (it *AggchainecdsaTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsaTransferAdminRole)
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
		it.Event = new(AggchainecdsaTransferAdminRole)
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
func (it *AggchainecdsaTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsaTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsaTransferAdminRole represents a TransferAdminRole event raised by the Aggchainecdsa contract.
type AggchainecdsaTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Aggchainecdsa *AggchainecdsaFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*AggchainecdsaTransferAdminRoleIterator, error) {

	logs, sub, err := _Aggchainecdsa.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsaTransferAdminRoleIterator{contract: _Aggchainecdsa.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Aggchainecdsa *AggchainecdsaFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *AggchainecdsaTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsa.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsaTransferAdminRole)
				if err := _Aggchainecdsa.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
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
func (_Aggchainecdsa *AggchainecdsaFilterer) ParseTransferAdminRole(log types.Log) (*AggchainecdsaTransferAdminRole, error) {
	event := new(AggchainecdsaTransferAdminRole)
	if err := _Aggchainecdsa.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsaTransferVKeyManagerRoleIterator is returned from FilterTransferVKeyManagerRole and is used to iterate over the raw logs and unpacked data for TransferVKeyManagerRole events raised by the Aggchainecdsa contract.
type AggchainecdsaTransferVKeyManagerRoleIterator struct {
	Event *AggchainecdsaTransferVKeyManagerRole // Event containing the contract specifics and raw log

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
func (it *AggchainecdsaTransferVKeyManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsaTransferVKeyManagerRole)
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
		it.Event = new(AggchainecdsaTransferVKeyManagerRole)
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
func (it *AggchainecdsaTransferVKeyManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsaTransferVKeyManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsaTransferVKeyManagerRole represents a TransferVKeyManagerRole event raised by the Aggchainecdsa contract.
type AggchainecdsaTransferVKeyManagerRole struct {
	CurrentVKeyManager    common.Address
	NewPendingVKeyManager common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterTransferVKeyManagerRole is a free log retrieval operation binding the contract event 0xc54ae01017d0b80bd8af833f66387d6eb547dc16c8206faf13d0b72764aab8b2.
//
// Solidity: event TransferVKeyManagerRole(address currentVKeyManager, address newPendingVKeyManager)
func (_Aggchainecdsa *AggchainecdsaFilterer) FilterTransferVKeyManagerRole(opts *bind.FilterOpts) (*AggchainecdsaTransferVKeyManagerRoleIterator, error) {

	logs, sub, err := _Aggchainecdsa.contract.FilterLogs(opts, "TransferVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsaTransferVKeyManagerRoleIterator{contract: _Aggchainecdsa.contract, event: "TransferVKeyManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferVKeyManagerRole is a free log subscription operation binding the contract event 0xc54ae01017d0b80bd8af833f66387d6eb547dc16c8206faf13d0b72764aab8b2.
//
// Solidity: event TransferVKeyManagerRole(address currentVKeyManager, address newPendingVKeyManager)
func (_Aggchainecdsa *AggchainecdsaFilterer) WatchTransferVKeyManagerRole(opts *bind.WatchOpts, sink chan<- *AggchainecdsaTransferVKeyManagerRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsa.contract.WatchLogs(opts, "TransferVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsaTransferVKeyManagerRole)
				if err := _Aggchainecdsa.contract.UnpackLog(event, "TransferVKeyManagerRole", log); err != nil {
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
func (_Aggchainecdsa *AggchainecdsaFilterer) ParseTransferVKeyManagerRole(log types.Log) (*AggchainecdsaTransferVKeyManagerRole, error) {
	event := new(AggchainecdsaTransferVKeyManagerRole)
	if err := _Aggchainecdsa.contract.UnpackLog(event, "TransferVKeyManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsaUpdateAggchainVKeyIterator is returned from FilterUpdateAggchainVKey and is used to iterate over the raw logs and unpacked data for UpdateAggchainVKey events raised by the Aggchainecdsa contract.
type AggchainecdsaUpdateAggchainVKeyIterator struct {
	Event *AggchainecdsaUpdateAggchainVKey // Event containing the contract specifics and raw log

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
func (it *AggchainecdsaUpdateAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsaUpdateAggchainVKey)
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
		it.Event = new(AggchainecdsaUpdateAggchainVKey)
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
func (it *AggchainecdsaUpdateAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsaUpdateAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsaUpdateAggchainVKey represents a UpdateAggchainVKey event raised by the Aggchainecdsa contract.
type AggchainecdsaUpdateAggchainVKey struct {
	Selector             [4]byte
	PreviousAggchainVKey [32]byte
	NewAggchainVKey      [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateAggchainVKey is a free log retrieval operation binding the contract event 0x0aa5f73c189fb0b0a7cc98ae5fa89dfc16595480396208483518178435ed5b4f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 previousAggchainVKey, bytes32 newAggchainVKey)
func (_Aggchainecdsa *AggchainecdsaFilterer) FilterUpdateAggchainVKey(opts *bind.FilterOpts) (*AggchainecdsaUpdateAggchainVKeyIterator, error) {

	logs, sub, err := _Aggchainecdsa.contract.FilterLogs(opts, "UpdateAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsaUpdateAggchainVKeyIterator{contract: _Aggchainecdsa.contract, event: "UpdateAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchUpdateAggchainVKey is a free log subscription operation binding the contract event 0x0aa5f73c189fb0b0a7cc98ae5fa89dfc16595480396208483518178435ed5b4f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 previousAggchainVKey, bytes32 newAggchainVKey)
func (_Aggchainecdsa *AggchainecdsaFilterer) WatchUpdateAggchainVKey(opts *bind.WatchOpts, sink chan<- *AggchainecdsaUpdateAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsa.contract.WatchLogs(opts, "UpdateAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsaUpdateAggchainVKey)
				if err := _Aggchainecdsa.contract.UnpackLog(event, "UpdateAggchainVKey", log); err != nil {
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
func (_Aggchainecdsa *AggchainecdsaFilterer) ParseUpdateAggchainVKey(log types.Log) (*AggchainecdsaUpdateAggchainVKey, error) {
	event := new(AggchainecdsaUpdateAggchainVKey)
	if err := _Aggchainecdsa.contract.UnpackLog(event, "UpdateAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"_rollupManager\",\"type\":\"address\"},{\"internalType\":\"contractIAggLayerGateway\",\"name\":\"_aggLayerGateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalAccInputHashDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggchainVKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingVKeyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyVKeyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultGatewayAlreadySet\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVKeyManager\",\"type\":\"address\"}],\"name\":\"AcceptVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"AddAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"OnVerifyPessimistic\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVKeyManager\",\"type\":\"address\"}],\"name\":\"TransferVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"UpdateAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"useDefaultGateway\",\"type\":\"bool\"}],\"name\":\"UpdateUseDefaultGatewayFlag\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AGGCHAIN_TYPE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AGGCHAIN_TYPE_SELECTOR\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptVKeyManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainSelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"addOwnedAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggLayerGateway\",\"outputs\":[{\"internalType\":\"contractIAggLayerGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableUseDefaultGatewayFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableUseDefaultGatewayFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"getAggchainHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainSelector\",\"type\":\"bytes4\"}],\"name\":\"getAggchainVKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"aggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initializeBytesCustomChain\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"onVerifyPessimistic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"ownedAggchainVKeys\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"ownedAggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingVKeyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVKeyManager\",\"type\":\"address\"}],\"name\":\"transferVKeyManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainSelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"updatedAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"updateOwnedAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useDefaultGateway\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vKeyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x610120604052348015610010575f5ffd5b5060405161249138038061249183398101604081905261002f91610153565b6001600160a01b0380861660a05280851660805280841660c052821660e05284848484848484848461005f61007f565b505050506001600160a01b031661010052506101c4975050505050505050565b5f54610100900460ff16156100ea5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116101561013a575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114610150575f5ffd5b50565b5f5f5f5f5f60a08688031215610167575f5ffd5b85516101728161013c565b60208701519095506101838161013c565b60408701519094506101948161013c565b60608701519093506101a58161013c565b60808701519092506101b68161013c565b809150509295509295909350565b60805160a05160c05160e0516101005161227b6102165f395f818161054401526107e501525f818161045701528181610c46015261129101525f61051d01525f6105f901525f610648015261227b5ff3fe608060405234801561000f575f5ffd5b506004361061029d575f3560e01c80637125702211610171578063cfa8ed47116100d2578063e631476c11610088578063effb84791161006e578063effb847914610686578063f851a440146106a5578063ff904079146106ca575f5ffd5b8063e631476c1461066a578063e7a7ed0214610672575f5ffd5b8063dc8c4249116100b8578063dc8c42491461061b578063e279984e14610623578063e46761c414610643575f5ffd5b8063cfa8ed47146105d4578063d02103ca146105f4575f5ffd5b8063ab0475cf11610127578063bfb193b61161010d578063bfb193b614610579578063c754c7ed14610599578063c89e42df146105c1575f5ffd5b8063ab0475cf1461053f578063ada8f91914610566575f5ffd5b80638c3d7301116101575780638c3d7301146104fd5780639ee4afa314610505578063a3c573eb14610518575f5ffd5b806371257022146104d757806385018182146104ea575f5ffd5b80633f17f31b1161021b5780636a55f66c116101d15780636e05d2cd116101b75780636e05d2cd146104b35780636e7fbce9146104bc5780636ff512cc146104c4575f5ffd5b80636a55f66c146104815780636b8616ce14610494575f5ffd5b80634560526711610201578063456052671461041957806349b7b80214610452578063542028d514610479575f5ffd5b80633f17f31b146103ce578063439fab9114610406575f5ffd5b80632c111c0611610270578063368c822c11610256578063368c822c146103695780633c351e10146103715780633cbc795b14610391575f5ffd5b80632c111c0614610336578063314eb17b14610356575f5ffd5b806301fcf6a0146102a1578063107bf28c146102c757806319451a8f146102dc57806326782247146102f1575b5f5ffd5b6102b46102af36600461193a565b6106ff565b6040519081526020015b60405180910390f35b6102cf610869565b6040516102be919061195c565b6102ef6102ea3660046119af565b6108f5565b005b6001546103119073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102be565b6008546103119073ffffffffffffffffffffffffffffffffffffffff1681565b6102ef6103643660046119af565b610a56565b6102ef610b77565b6009546103119073ffffffffffffffffffffffffffffffffffffffff1681565b6009546103b99074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff90911681526020016102be565b6103d55f81565b6040517fffff00000000000000000000000000000000000000000000000000000000000090911681526020016102be565b6102ef610414366004611ad5565b610c44565b6007546104399068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020016102be565b6103117f000000000000000000000000000000000000000000000000000000000000000081565b6102cf610ef3565b6102b461048f366004611ad5565b610f00565b6102b46104a2366004611b22565b60066020525f908152604090205481565b6102b460055481565b6103b9600181565b6102ef6104d2366004611b6a565b610ffc565b6102ef6104e5366004611ba3565b6110cc565b6102ef6104f8366004611b6a565b6110fe565b6102ef6111c2565b6102ef610513366004611c57565b61128f565b6103117f000000000000000000000000000000000000000000000000000000000000000081565b6103117f000000000000000000000000000000000000000000000000000000000000000081565b6102ef610574366004611b6a565b61134c565b603d546103119073ffffffffffffffffffffffffffffffffffffffff1681565b60075461043990700100000000000000000000000000000000900467ffffffffffffffff1681565b6102ef6105cf366004611cc5565b611415565b6002546103119073ffffffffffffffffffffffffffffffffffffffff1681565b6103117f000000000000000000000000000000000000000000000000000000000000000081565b6102ef6114a7565b603c546103119073ffffffffffffffffffffffffffffffffffffffff1681565b6103117f000000000000000000000000000000000000000000000000000000000000000081565b6102ef6115c5565b6007546104399067ffffffffffffffff1681565b6102b461069436600461193a565b603e6020525f908152604090205481565b5f546103119062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b603d546106ef9074010000000000000000000000000000000000000000900460ff1681565b60405190151581526020016102be565b603d545f9074010000000000000000000000000000000000000000900460ff161515810361079557507fffffffff0000000000000000000000000000000000000000000000000000000081165f908152603e602052604090205480610790576040517f925e5a3a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b919050565b6040517f6cabfdab0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000831660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690636cabfdab90602401602060405180830381865afa15801561083f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108639190611cf7565b92915050565b6004805461087690611d0e565b80601f01602080910402602001604051908101604052809291908181526020018280546108a290611d0e565b80156108ed5780601f106108c4576101008083540402835291602001916108ed565b820191905f5260205f20905b8154815290600101906020018083116108d057829003601f168201915b505050505081565b603c5473ffffffffffffffffffffffffffffffffffffffff163314610946576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8061097d576040517f4aac8b8800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f908152603e6020526040902054156109e5576040517fe3cc761000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f818152603e6020908152604091829020849055815192835282018390527f6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a7991015b60405180910390a15050565b603c5473ffffffffffffffffffffffffffffffffffffffff163314610aa7576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f908152603e6020526040902054610b0e576040517ff360deaf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f818152603e6020908152604091829020849055815192835282018390527f2ae81c794f274593f3cd22db9ad3af161ff9d93d15922b35629a6ffd9be7243f9101610a4a565b603d5473ffffffffffffffffffffffffffffffffffffffff163314610bc8576040517f05882cf000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d54603c80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691821790556040519081527f75ccbacededf0b9fc9371bb4cf651be4f96efab2c519a529cfd90e14c9d49ad2906020015b60405180910390a1565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314610cb3576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f805460ff16907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00815c168217905d505f54600290610100900460ff16158015610d0357505f5460ff8083169116105b610d93576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840160405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff80841691909117610100178255815c169003610e26575f5f5f5f5f5f5f5f5f8a806020019051810190610ded9190611eb2565b985098509850985098509850985098509850610e0b898989896116e6565b610e188585858585611829565b505050505050505050610e99565b60ff5f5c16600103610e67575f5f5f5f85806020019051810190610e4a9190611fbc565b9350935093509350610e5e848484846116e6565b50505050610e99565b6040517fadc06ae700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001610a4a565b6003805461087690611d0e565b5f5f82806020019051810190610f169190612070565b5090507fffff00000000000000000000000000000000000000000000000000000000000081166001610f47826106ff565b60025460405160609190911b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602082015260340160405160208183030381529060405280519060200120604051602001610fdd9392919060e09390931b7fffffffff000000000000000000000000000000000000000000000000000000001683526004830191909152602482015260440190565b6040516020818303038152906040528051906020012092505050919050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314611052576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0906020015b60405180910390a150565b6040517ff57ac68300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603c5473ffffffffffffffffffffffffffffffffffffffff16331461114f576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527ffa4ea207951b027c94aa15f12ca5d0f4ca0543beb8426209ac79b49d2fc46b10906020016110c1565b60015473ffffffffffffffffffffffffffffffffffffffff163314611213576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001545f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff1673ffffffffffffffffffffffffffffffffffffffff9092166201000081029290921790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e90602001610c3a565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1633146112fe576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61130b8284018461209c565b9150507f18eea7739a19b6e71e562030f68e6850aa7f1408136e6795b2690501810c41088160405161133f91815260200190565b60405180910390a1505050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff1633146113a2576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6906020016110c1565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff16331461146b576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60036114778282612101565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20816040516110c1919061195c565b603c5473ffffffffffffffffffffffffffffffffffffffff1633146114f8576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d5474010000000000000000000000000000000000000000900460ff1661154c576040517f6f318e4c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16908190556040517fd4305f3587ac0e4b2b8b4e5b2072599efd931f81aba626cdbef4b040078b507091610c3a917401000000000000000000000000000000000000000090910460ff161515815260200190565b603c5473ffffffffffffffffffffffffffffffffffffffff163314611616576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d5474010000000000000000000000000000000000000000900460ff161561166b576040517f6f318e4c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1674010000000000000000000000000000000000000000908117918290556040517fd4305f3587ac0e4b2b8b4e5b2072599efd931f81aba626cdbef4b040078b507092610c3a92900460ff161515815260200190565b603d80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000086151502179055603c80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff831617905581518351146117a4576040517f6896511300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b8351811015611822578381815181106117c1576117c1612218565b6020026020010151603e5f8584815181106117de576117de612218565b6020908102919091018101517fffffffff000000000000000000000000000000000000000000000000000000001682528101919091526040015f20556001016117a6565b5050505050565b5f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff166201000073ffffffffffffffffffffffffffffffffffffffff8881169190910291909117909155600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001691861691909117905560036118b08382612101565b5060046118bd8282612101565b5050600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9390931692909217909155505050565b7fffffffff0000000000000000000000000000000000000000000000000000000081168114611937575f5ffd5b50565b5f6020828403121561194a575f5ffd5b81356119558161190a565b9392505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b5f5f604083850312156119c0575f5ffd5b82356119cb8161190a565b946020939093013593505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611a4d57611a4d6119d9565b604052919050565b5f67ffffffffffffffff821115611a6e57611a6e6119d9565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b5f611aac611aa784611a55565b611a06565b9050828152838383011115611abf575f5ffd5b828260208301375f602084830101529392505050565b5f60208284031215611ae5575f5ffd5b813567ffffffffffffffff811115611afb575f5ffd5b8201601f81018413611b0b575f5ffd5b611b1a84823560208401611a9a565b949350505050565b5f60208284031215611b32575f5ffd5b813567ffffffffffffffff81168114611955575f5ffd5b73ffffffffffffffffffffffffffffffffffffffff81168114611937575f5ffd5b5f60208284031215611b7a575f5ffd5b813561195581611b49565b5f82601f830112611b94575f5ffd5b61195583833560208501611a9a565b5f5f5f5f5f5f60c08789031215611bb8575f5ffd5b8635611bc381611b49565b95506020870135611bd381611b49565b9450604087013563ffffffff81168114611beb575f5ffd5b93506060870135611bfb81611b49565b9250608087013567ffffffffffffffff811115611c16575f5ffd5b611c2289828a01611b85565b92505060a087013567ffffffffffffffff811115611c3e575f5ffd5b611c4a89828a01611b85565b9150509295509295509295565b5f5f60208385031215611c68575f5ffd5b823567ffffffffffffffff811115611c7e575f5ffd5b8301601f81018513611c8e575f5ffd5b803567ffffffffffffffff811115611ca4575f5ffd5b856020828401011115611cb5575f5ffd5b6020919091019590945092505050565b5f60208284031215611cd5575f5ffd5b813567ffffffffffffffff811115611ceb575f5ffd5b611b1a84828501611b85565b5f60208284031215611d07575f5ffd5b5051919050565b600181811c90821680611d2257607f821691505b602082108103611d59577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b80518015158114610790575f5ffd5b5f67ffffffffffffffff821115611d8757611d876119d9565b5060051b60200190565b5f82601f830112611da0575f5ffd5b8151611dae611aa782611d6e565b8082825260208201915060208360051b860101925085831115611dcf575f5ffd5b602085015b83811015611dec578051835260209283019201611dd4565b5095945050505050565b5f82601f830112611e05575f5ffd5b8151611e13611aa782611d6e565b8082825260208201915060208360051b860101925085831115611e34575f5ffd5b602085015b83811015611dec578051611e4c8161190a565b835260209283019201611e39565b805161079081611b49565b5f82601f830112611e74575f5ffd5b8151611e82611aa782611a55565b818152846020838601011115611e96575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f5f5f5f5f5f5f5f5f6101208a8c031215611ecb575f5ffd5b611ed48a611d5f565b985060208a015167ffffffffffffffff811115611eef575f5ffd5b611efb8c828d01611d91565b98505060408a015167ffffffffffffffff811115611f17575f5ffd5b611f238c828d01611df6565b975050611f3260608b01611e5a565b9550611f4060808b01611e5a565b9450611f4e60a08b01611e5a565b9350611f5c60c08b01611e5a565b925060e08a015167ffffffffffffffff811115611f77575f5ffd5b611f838c828d01611e65565b9250506101008a015167ffffffffffffffff811115611fa0575f5ffd5b611fac8c828d01611e65565b9150509295985092959850929598565b5f5f5f5f60808587031215611fcf575f5ffd5b611fd885611d5f565b9350602085015167ffffffffffffffff811115611ff3575f5ffd5b611fff87828801611d91565b935050604085015167ffffffffffffffff81111561201b575f5ffd5b61202787828801611df6565b925050606085015161203881611b49565b939692955090935050565b7fffff00000000000000000000000000000000000000000000000000000000000081168114611937575f5ffd5b5f5f60408385031215612081575f5ffd5b825161208c81612043565b6020939093015192949293505050565b5f5f604083850312156120ad575f5ffd5b82356119cb81612043565b601f8211156120fc57805f5260205f20601f840160051c810160208510156120dd5750805b601f840160051c820191505b81811015611822575f81556001016120e9565b505050565b815167ffffffffffffffff81111561211b5761211b6119d9565b61212f816121298454611d0e565b846120b8565b6020601f821160018114612180575f831561214a5750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455611822565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b828110156121cd57878501518255602094850194600190920191016121ad565b508482101561220957868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea264697066735822122071cdfb81cfdcd7a85bf79c5806df90ea49ce19526a2f95500b0f264be1e009e364736f6c634300081c0033",
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
// Solidity: function AGGCHAIN_TYPE() view returns(uint32)
func (_Aggchainecdsa *AggchainecdsaCaller) AGGCHAINTYPE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "AGGCHAIN_TYPE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AGGCHAINTYPE is a free data retrieval call binding the contract method 0x6e7fbce9.
//
// Solidity: function AGGCHAIN_TYPE() view returns(uint32)
func (_Aggchainecdsa *AggchainecdsaSession) AGGCHAINTYPE() (uint32, error) {
	return _Aggchainecdsa.Contract.AGGCHAINTYPE(&_Aggchainecdsa.CallOpts)
}

// AGGCHAINTYPE is a free data retrieval call binding the contract method 0x6e7fbce9.
//
// Solidity: function AGGCHAIN_TYPE() view returns(uint32)
func (_Aggchainecdsa *AggchainecdsaCallerSession) AGGCHAINTYPE() (uint32, error) {
	return _Aggchainecdsa.Contract.AGGCHAINTYPE(&_Aggchainecdsa.CallOpts)
}

// AGGCHAINTYPESELECTOR is a free data retrieval call binding the contract method 0x3f17f31b.
//
// Solidity: function AGGCHAIN_TYPE_SELECTOR() view returns(bytes2)
func (_Aggchainecdsa *AggchainecdsaCaller) AGGCHAINTYPESELECTOR(opts *bind.CallOpts) ([2]byte, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "AGGCHAIN_TYPE_SELECTOR")

	if err != nil {
		return *new([2]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([2]byte)).(*[2]byte)

	return out0, err

}

// AGGCHAINTYPESELECTOR is a free data retrieval call binding the contract method 0x3f17f31b.
//
// Solidity: function AGGCHAIN_TYPE_SELECTOR() view returns(bytes2)
func (_Aggchainecdsa *AggchainecdsaSession) AGGCHAINTYPESELECTOR() ([2]byte, error) {
	return _Aggchainecdsa.Contract.AGGCHAINTYPESELECTOR(&_Aggchainecdsa.CallOpts)
}

// AGGCHAINTYPESELECTOR is a free data retrieval call binding the contract method 0x3f17f31b.
//
// Solidity: function AGGCHAIN_TYPE_SELECTOR() view returns(bytes2)
func (_Aggchainecdsa *AggchainecdsaCallerSession) AGGCHAINTYPESELECTOR() ([2]byte, error) {
	return _Aggchainecdsa.Contract.AGGCHAINTYPESELECTOR(&_Aggchainecdsa.CallOpts)
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
// Solidity: function getAggchainVKey(bytes4 aggchainSelector) view returns(bytes32 aggchainVKey)
func (_Aggchainecdsa *AggchainecdsaCaller) GetAggchainVKey(opts *bind.CallOpts, aggchainSelector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "getAggchainVKey", aggchainSelector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAggchainVKey is a free data retrieval call binding the contract method 0x01fcf6a0.
//
// Solidity: function getAggchainVKey(bytes4 aggchainSelector) view returns(bytes32 aggchainVKey)
func (_Aggchainecdsa *AggchainecdsaSession) GetAggchainVKey(aggchainSelector [4]byte) ([32]byte, error) {
	return _Aggchainecdsa.Contract.GetAggchainVKey(&_Aggchainecdsa.CallOpts, aggchainSelector)
}

// GetAggchainVKey is a free data retrieval call binding the contract method 0x01fcf6a0.
//
// Solidity: function getAggchainVKey(bytes4 aggchainSelector) view returns(bytes32 aggchainVKey)
func (_Aggchainecdsa *AggchainecdsaCallerSession) GetAggchainVKey(aggchainSelector [4]byte) ([32]byte, error) {
	return _Aggchainecdsa.Contract.GetAggchainVKey(&_Aggchainecdsa.CallOpts, aggchainSelector)
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
// Solidity: function addOwnedAggchainVKey(bytes4 aggchainSelector, bytes32 newAggchainVKey) returns()
func (_Aggchainecdsa *AggchainecdsaTransactor) AddOwnedAggchainVKey(opts *bind.TransactOpts, aggchainSelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainecdsa.contract.Transact(opts, "addOwnedAggchainVKey", aggchainSelector, newAggchainVKey)
}

// AddOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x19451a8f.
//
// Solidity: function addOwnedAggchainVKey(bytes4 aggchainSelector, bytes32 newAggchainVKey) returns()
func (_Aggchainecdsa *AggchainecdsaSession) AddOwnedAggchainVKey(aggchainSelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.AddOwnedAggchainVKey(&_Aggchainecdsa.TransactOpts, aggchainSelector, newAggchainVKey)
}

// AddOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x19451a8f.
//
// Solidity: function addOwnedAggchainVKey(bytes4 aggchainSelector, bytes32 newAggchainVKey) returns()
func (_Aggchainecdsa *AggchainecdsaTransactorSession) AddOwnedAggchainVKey(aggchainSelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.AddOwnedAggchainVKey(&_Aggchainecdsa.TransactOpts, aggchainSelector, newAggchainVKey)
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
// Solidity: function initialize(bytes initializeBytesCustomChain) returns()
func (_Aggchainecdsa *AggchainecdsaTransactor) Initialize(opts *bind.TransactOpts, initializeBytesCustomChain []byte) (*types.Transaction, error) {
	return _Aggchainecdsa.contract.Transact(opts, "initialize", initializeBytesCustomChain)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializeBytesCustomChain) returns()
func (_Aggchainecdsa *AggchainecdsaSession) Initialize(initializeBytesCustomChain []byte) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.Initialize(&_Aggchainecdsa.TransactOpts, initializeBytesCustomChain)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializeBytesCustomChain) returns()
func (_Aggchainecdsa *AggchainecdsaTransactorSession) Initialize(initializeBytesCustomChain []byte) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.Initialize(&_Aggchainecdsa.TransactOpts, initializeBytesCustomChain)
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
// Solidity: function updateOwnedAggchainVKey(bytes4 aggchainSelector, bytes32 updatedAggchainVKey) returns()
func (_Aggchainecdsa *AggchainecdsaTransactor) UpdateOwnedAggchainVKey(opts *bind.TransactOpts, aggchainSelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainecdsa.contract.Transact(opts, "updateOwnedAggchainVKey", aggchainSelector, updatedAggchainVKey)
}

// UpdateOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x314eb17b.
//
// Solidity: function updateOwnedAggchainVKey(bytes4 aggchainSelector, bytes32 updatedAggchainVKey) returns()
func (_Aggchainecdsa *AggchainecdsaSession) UpdateOwnedAggchainVKey(aggchainSelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.UpdateOwnedAggchainVKey(&_Aggchainecdsa.TransactOpts, aggchainSelector, updatedAggchainVKey)
}

// UpdateOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x314eb17b.
//
// Solidity: function updateOwnedAggchainVKey(bytes4 aggchainSelector, bytes32 updatedAggchainVKey) returns()
func (_Aggchainecdsa *AggchainecdsaTransactorSession) UpdateOwnedAggchainVKey(aggchainSelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.UpdateOwnedAggchainVKey(&_Aggchainecdsa.TransactOpts, aggchainSelector, updatedAggchainVKey)
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
	NewVKeyManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAcceptVKeyManagerRole is a free log retrieval operation binding the contract event 0x75ccbacededf0b9fc9371bb4cf651be4f96efab2c519a529cfd90e14c9d49ad2.
//
// Solidity: event AcceptVKeyManagerRole(address newVKeyManager)
func (_Aggchainecdsa *AggchainecdsaFilterer) FilterAcceptVKeyManagerRole(opts *bind.FilterOpts) (*AggchainecdsaAcceptVKeyManagerRoleIterator, error) {

	logs, sub, err := _Aggchainecdsa.contract.FilterLogs(opts, "AcceptVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsaAcceptVKeyManagerRoleIterator{contract: _Aggchainecdsa.contract, event: "AcceptVKeyManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptVKeyManagerRole is a free log subscription operation binding the contract event 0x75ccbacededf0b9fc9371bb4cf651be4f96efab2c519a529cfd90e14c9d49ad2.
//
// Solidity: event AcceptVKeyManagerRole(address newVKeyManager)
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

// ParseAcceptVKeyManagerRole is a log parse operation binding the contract event 0x75ccbacededf0b9fc9371bb4cf651be4f96efab2c519a529cfd90e14c9d49ad2.
//
// Solidity: event AcceptVKeyManagerRole(address newVKeyManager)
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

// AggchainecdsaOnVerifyPessimisticIterator is returned from FilterOnVerifyPessimistic and is used to iterate over the raw logs and unpacked data for OnVerifyPessimistic events raised by the Aggchainecdsa contract.
type AggchainecdsaOnVerifyPessimisticIterator struct {
	Event *AggchainecdsaOnVerifyPessimistic // Event containing the contract specifics and raw log

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
func (it *AggchainecdsaOnVerifyPessimisticIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsaOnVerifyPessimistic)
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
		it.Event = new(AggchainecdsaOnVerifyPessimistic)
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
func (it *AggchainecdsaOnVerifyPessimisticIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsaOnVerifyPessimisticIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsaOnVerifyPessimistic represents a OnVerifyPessimistic event raised by the Aggchainecdsa contract.
type AggchainecdsaOnVerifyPessimistic struct {
	NewStateRoot [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOnVerifyPessimistic is a free log retrieval operation binding the contract event 0x18eea7739a19b6e71e562030f68e6850aa7f1408136e6795b2690501810c4108.
//
// Solidity: event OnVerifyPessimistic(bytes32 newStateRoot)
func (_Aggchainecdsa *AggchainecdsaFilterer) FilterOnVerifyPessimistic(opts *bind.FilterOpts) (*AggchainecdsaOnVerifyPessimisticIterator, error) {

	logs, sub, err := _Aggchainecdsa.contract.FilterLogs(opts, "OnVerifyPessimistic")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsaOnVerifyPessimisticIterator{contract: _Aggchainecdsa.contract, event: "OnVerifyPessimistic", logs: logs, sub: sub}, nil
}

// WatchOnVerifyPessimistic is a free log subscription operation binding the contract event 0x18eea7739a19b6e71e562030f68e6850aa7f1408136e6795b2690501810c4108.
//
// Solidity: event OnVerifyPessimistic(bytes32 newStateRoot)
func (_Aggchainecdsa *AggchainecdsaFilterer) WatchOnVerifyPessimistic(opts *bind.WatchOpts, sink chan<- *AggchainecdsaOnVerifyPessimistic) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsa.contract.WatchLogs(opts, "OnVerifyPessimistic")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsaOnVerifyPessimistic)
				if err := _Aggchainecdsa.contract.UnpackLog(event, "OnVerifyPessimistic", log); err != nil {
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

// ParseOnVerifyPessimistic is a log parse operation binding the contract event 0x18eea7739a19b6e71e562030f68e6850aa7f1408136e6795b2690501810c4108.
//
// Solidity: event OnVerifyPessimistic(bytes32 newStateRoot)
func (_Aggchainecdsa *AggchainecdsaFilterer) ParseOnVerifyPessimistic(log types.Log) (*AggchainecdsaOnVerifyPessimistic, error) {
	event := new(AggchainecdsaOnVerifyPessimistic)
	if err := _Aggchainecdsa.contract.UnpackLog(event, "OnVerifyPessimistic", log); err != nil {
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
	NewVKeyManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTransferVKeyManagerRole is a free log retrieval operation binding the contract event 0xfa4ea207951b027c94aa15f12ca5d0f4ca0543beb8426209ac79b49d2fc46b10.
//
// Solidity: event TransferVKeyManagerRole(address newVKeyManager)
func (_Aggchainecdsa *AggchainecdsaFilterer) FilterTransferVKeyManagerRole(opts *bind.FilterOpts) (*AggchainecdsaTransferVKeyManagerRoleIterator, error) {

	logs, sub, err := _Aggchainecdsa.contract.FilterLogs(opts, "TransferVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsaTransferVKeyManagerRoleIterator{contract: _Aggchainecdsa.contract, event: "TransferVKeyManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferVKeyManagerRole is a free log subscription operation binding the contract event 0xfa4ea207951b027c94aa15f12ca5d0f4ca0543beb8426209ac79b49d2fc46b10.
//
// Solidity: event TransferVKeyManagerRole(address newVKeyManager)
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

// ParseTransferVKeyManagerRole is a log parse operation binding the contract event 0xfa4ea207951b027c94aa15f12ca5d0f4ca0543beb8426209ac79b49d2fc46b10.
//
// Solidity: event TransferVKeyManagerRole(address newVKeyManager)
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
	Selector        [4]byte
	NewAggchainVKey [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateAggchainVKey is a free log retrieval operation binding the contract event 0x2ae81c794f274593f3cd22db9ad3af161ff9d93d15922b35629a6ffd9be7243f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Aggchainecdsa *AggchainecdsaFilterer) FilterUpdateAggchainVKey(opts *bind.FilterOpts) (*AggchainecdsaUpdateAggchainVKeyIterator, error) {

	logs, sub, err := _Aggchainecdsa.contract.FilterLogs(opts, "UpdateAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsaUpdateAggchainVKeyIterator{contract: _Aggchainecdsa.contract, event: "UpdateAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchUpdateAggchainVKey is a free log subscription operation binding the contract event 0x2ae81c794f274593f3cd22db9ad3af161ff9d93d15922b35629a6ffd9be7243f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
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

// ParseUpdateAggchainVKey is a log parse operation binding the contract event 0x2ae81c794f274593f3cd22db9ad3af161ff9d93d15922b35629a6ffd9be7243f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Aggchainecdsa *AggchainecdsaFilterer) ParseUpdateAggchainVKey(log types.Log) (*AggchainecdsaUpdateAggchainVKey, error) {
	event := new(AggchainecdsaUpdateAggchainVKey)
	if err := _Aggchainecdsa.contract.UnpackLog(event, "UpdateAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsaUpdateUseDefaultGatewayFlagIterator is returned from FilterUpdateUseDefaultGatewayFlag and is used to iterate over the raw logs and unpacked data for UpdateUseDefaultGatewayFlag events raised by the Aggchainecdsa contract.
type AggchainecdsaUpdateUseDefaultGatewayFlagIterator struct {
	Event *AggchainecdsaUpdateUseDefaultGatewayFlag // Event containing the contract specifics and raw log

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
func (it *AggchainecdsaUpdateUseDefaultGatewayFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsaUpdateUseDefaultGatewayFlag)
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
		it.Event = new(AggchainecdsaUpdateUseDefaultGatewayFlag)
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
func (it *AggchainecdsaUpdateUseDefaultGatewayFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsaUpdateUseDefaultGatewayFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsaUpdateUseDefaultGatewayFlag represents a UpdateUseDefaultGatewayFlag event raised by the Aggchainecdsa contract.
type AggchainecdsaUpdateUseDefaultGatewayFlag struct {
	UseDefaultGateway bool
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdateUseDefaultGatewayFlag is a free log retrieval operation binding the contract event 0xd4305f3587ac0e4b2b8b4e5b2072599efd931f81aba626cdbef4b040078b5070.
//
// Solidity: event UpdateUseDefaultGatewayFlag(bool useDefaultGateway)
func (_Aggchainecdsa *AggchainecdsaFilterer) FilterUpdateUseDefaultGatewayFlag(opts *bind.FilterOpts) (*AggchainecdsaUpdateUseDefaultGatewayFlagIterator, error) {

	logs, sub, err := _Aggchainecdsa.contract.FilterLogs(opts, "UpdateUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsaUpdateUseDefaultGatewayFlagIterator{contract: _Aggchainecdsa.contract, event: "UpdateUseDefaultGatewayFlag", logs: logs, sub: sub}, nil
}

// WatchUpdateUseDefaultGatewayFlag is a free log subscription operation binding the contract event 0xd4305f3587ac0e4b2b8b4e5b2072599efd931f81aba626cdbef4b040078b5070.
//
// Solidity: event UpdateUseDefaultGatewayFlag(bool useDefaultGateway)
func (_Aggchainecdsa *AggchainecdsaFilterer) WatchUpdateUseDefaultGatewayFlag(opts *bind.WatchOpts, sink chan<- *AggchainecdsaUpdateUseDefaultGatewayFlag) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsa.contract.WatchLogs(opts, "UpdateUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsaUpdateUseDefaultGatewayFlag)
				if err := _Aggchainecdsa.contract.UnpackLog(event, "UpdateUseDefaultGatewayFlag", log); err != nil {
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

// ParseUpdateUseDefaultGatewayFlag is a log parse operation binding the contract event 0xd4305f3587ac0e4b2b8b4e5b2072599efd931f81aba626cdbef4b040078b5070.
//
// Solidity: event UpdateUseDefaultGatewayFlag(bool useDefaultGateway)
func (_Aggchainecdsa *AggchainecdsaFilterer) ParseUpdateUseDefaultGatewayFlag(log types.Log) (*AggchainecdsaUpdateUseDefaultGatewayFlag, error) {
	event := new(AggchainecdsaUpdateUseDefaultGatewayFlag)
	if err := _Aggchainecdsa.contract.UnpackLog(event, "UpdateUseDefaultGatewayFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

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

// AggchainfepMetaData contains all meta data concerning the Aggchainfep contract.
var AggchainfepMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"_rollupManager\",\"type\":\"address\"},{\"internalType\":\"contractAggLayerGateway\",\"name\":\"_aggLayerGateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalAccInputHashDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggchainVKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingVKeyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyVKeyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultGatewayAlreadySet\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVKeyManager\",\"type\":\"address\"}],\"name\":\"AcceptVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"AddAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"initStateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"initTimestamp\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"initL2BlockNumber\",\"type\":\"uint128\"}],\"name\":\"OnVerifyPessimistic\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVKeyManager\",\"type\":\"address\"}],\"name\":\"TransferVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"UpdateAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"useDefaultGateway\",\"type\":\"bool\"}],\"name\":\"UpdateUseDefaultGatewayFlag\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AGGCHAIN_TYPE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AGGCHAIN_TYPE_SELECTOR\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptVKeyManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainSelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"addOwnedAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggLayerGateway\",\"outputs\":[{\"internalType\":\"contractIAggLayerGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggregationVkey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainConfigHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"chainData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"lastStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableUseDefaultGatewayFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableUseDefaultGatewayFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"customChainData\",\"type\":\"bytes\"}],\"name\":\"getAggchainHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainSelector\",\"type\":\"bytes4\"}],\"name\":\"getAggchainVKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"aggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initializeBytesCustomChain\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"customData\",\"type\":\"bytes\"}],\"name\":\"onVerifyPessimistic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"ownedAggchainVKeys\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"ownedAggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingVKeyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rangeVkeyCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVKeyManager\",\"type\":\"address\"}],\"name\":\"transferVKeyManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"customInitializeData\",\"type\":\"bytes\"}],\"name\":\"updateCustomInitializeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainSelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"updatedAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"updateOwnedAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useDefaultGateway\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vKeyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x610120604052348015610010575f5ffd5b506040516126ea3803806126ea83398101604081905261002f91610153565b6001600160a01b0380861660a05280851660805280841660c052821660e05284848484848484848461005f61007f565b505050506001600160a01b031661010052506101c4975050505050505050565b5f54610100900460ff16156100ea5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116101561013a575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114610150575f5ffd5b50565b5f5f5f5f5f60a08688031215610167575f5ffd5b85516101728161013c565b60208701519095506101838161013c565b60408701519094506101948161013c565b60608701519093506101a58161013c565b60808701519092506101b68161013c565b809150509295509295909350565b60805160a05160c05160e051610100516124d46102165f395f81816105c501526108b901525f81816104c501528181610d1a015261171901525f61059e01525f61068c01525f6106db01526124d45ff3fe608060405234801561000f575f5ffd5b50600436106102e3575f3560e01c80638501818211610187578063cfa8ed47116100dd578063e59f69c211610093578063effb84791161006e578063effb84791461075a578063f851a44014610779578063ff9040791461079e575f5ffd5b8063e59f69c2146106fd578063e631476c1461073e578063e7a7ed0214610746575f5ffd5b8063dc8c4249116100c3578063dc8c4249146106ae578063e279984e146106b6578063e46761c4146106d6575f5ffd5b8063cfa8ed4714610667578063d02103ca14610687575f5ffd5b8063ada8f9191161013d578063c32e4e3e11610118578063c32e4e3e14610623578063c754c7ed1461062c578063c89e42df14610654575f5ffd5b8063ada8f919146105e7578063b7fd13ce146105fa578063bfb193b614610603575f5ffd5b80639ee4afa31161016d5780639ee4afa314610586578063a3c573eb14610599578063ab0475cf146105c0575f5ffd5b8063850181821461056b5780638c3d73011461057e575f5ffd5b8063439fab911161023c5780636b8616ce116101f25780636ff512cc116101cd5780636ff512cc1461053257806371257022146105455780637e005eb214610558575f5ffd5b80636b8616ce146105025780636e05d2cd146105215780636e7fbce91461052a575f5ffd5b806349b7b8021161022257806349b7b802146104c0578063542028d5146104e75780636a55f66c146104ef575f5ffd5b8063439fab91146104745780634560526714610487575f5ffd5b80632c111c061161029c5780633c351e10116102775780633c351e10146103c05780633cbc795b146103e05780633f17f31b1461041d575f5ffd5b80632c111c0614610385578063314eb17b146103a5578063368c822c146103b8575f5ffd5b806319451a8f116102cc57806319451a8f1461032257806326782247146103375780632b31841e1461037c575f5ffd5b806301fcf6a0146102e7578063107bf28c1461030d575b5f5ffd5b6102fa6102f5366004611cd6565b6107d3565b6040519081526020015b60405180910390f35b61031561093d565b6040516103049190611cf6565b610335610330366004611d49565b6109c9565b005b6001546103579073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610304565b6102fa60735481565b6008546103579073ffffffffffffffffffffffffffffffffffffffff1681565b6103356103b3366004611d49565b610b2a565b610335610c4b565b6009546103579073ffffffffffffffffffffffffffffffffffffffff1681565b6009546104089074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff9091168152602001610304565b6104437e0100000000000000000000000000000000000000000000000000000000000081565b6040517fffff0000000000000000000000000000000000000000000000000000000000009091168152602001610304565b610335610482366004611d71565b610d18565b6007546104a79068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff9091168152602001610304565b6103577f000000000000000000000000000000000000000000000000000000000000000081565b6103156112ce565b6102fa6104fd366004611ea1565b6112db565b6102fa610510366004611eee565b60066020525f908152604090205481565b6102fa60055481565b610408600181565b610335610540366004611f44565b61140c565b610335610553366004611f7d565b6114dc565b610335610566366004611d71565b61150e565b610335610579366004611f44565b611586565b61033561164a565b610335610594366004611ea1565b611717565b6103577f000000000000000000000000000000000000000000000000000000000000000081565b6103577f000000000000000000000000000000000000000000000000000000000000000081565b6103356105f5366004611f44565b6118b7565b6102fa60725481565b603d546103579073ffffffffffffffffffffffffffffffffffffffff1681565b6102fa60715481565b6007546104a790700100000000000000000000000000000000900467ffffffffffffffff1681565b610335610662366004612031565b611980565b6002546103579073ffffffffffffffffffffffffffffffffffffffff1681565b6103577f000000000000000000000000000000000000000000000000000000000000000081565b610335611a12565b603c546103579073ffffffffffffffffffffffffffffffffffffffff1681565b6103577f000000000000000000000000000000000000000000000000000000000000000081565b61071061070b366004612063565b611b30565b604080519384526fffffffffffffffffffffffffffffffff9283166020850152911690820152606001610304565b610335611b86565b6007546104a79067ffffffffffffffff1681565b6102fa610768366004611cd6565b603e6020525f908152604090205481565b5f546103579062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b603d546107c39074010000000000000000000000000000000000000000900460ff1681565b6040519015158152602001610304565b603d545f9074010000000000000000000000000000000000000000900460ff161515810361086957507fffffffff0000000000000000000000000000000000000000000000000000000081165f908152603e602052604090205480610864576040517f925e5a3a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b919050565b6040517f6cabfdab0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000831660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690636cabfdab90602401602060405180830381865afa158015610913573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610937919061207a565b92915050565b6004805461094a90612091565b80601f016020809104026020016040519081016040528092919081815260200182805461097690612091565b80156109c15780601f10610998576101008083540402835291602001916109c1565b820191905f5260205f20905b8154815290600101906020018083116109a457829003601f168201915b505050505081565b603c5473ffffffffffffffffffffffffffffffffffffffff163314610a1a576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80610a51576040517f4aac8b8800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f908152603e602052604090205415610ab9576040517fe3cc761000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f818152603e6020908152604091829020849055815192835282018390527f6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a7991015b60405180910390a15050565b603c5473ffffffffffffffffffffffffffffffffffffffff163314610b7b576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f908152603e6020526040902054610be2576040517ff360deaf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f818152603e6020908152604091829020849055815192835282018390527f2ae81c794f274593f3cd22db9ad3af161ff9d93d15922b35629a6ffd9be7243f9101610b1e565b603d5473ffffffffffffffffffffffffffffffffffffffff163314610c9c576040517f05882cf000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d54603c80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691821790556040519081527f75ccbacededf0b9fc9371bb4cf651be4f96efab2c519a529cfd90e14c9d49ad2906020015b60405180910390a1565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314610d87576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f805460ff16907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00815c168217905d505f54600290610100900460ff16158015610dd757505f5460ff8083169116105b610e67576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840160405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff80841691909117610100178255815c1690036110c7575f5f5f5f5f5f5f5f5f5f5f5f8e8e810190610ebf919061210a565b9b509b509b509b509b509b509b509b509b509b509b509b508b5f60026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508a60025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508960095f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508860039081610fa69190612247565b506004610fb38982612247565b50607196909655607294909455607392909255603c805473ffffffffffffffffffffffffffffffffffffffff9095167fffffffffffffffffffffffff000000000000000000000000000000000000000090951694909417909355604080516060810182529384526fffffffffffffffffffffffffffffffff91821660208501908152928216908401908152607480546001810182555f91909152935160029094027f19a0b39aa25ac793b5f6e9a0534364cc0b3fd1ea9b651e79c7f50a59d48ef813810194909455915191518116700100000000000000000000000000000000029116177f19a0b39aa25ac793b5f6e9a0534364cc0b3fd1ea9b651e79c7f50a59d48ef814909101555061122d9350505050565b60ff5f5c166001036111fb575f8080808080806110e6898b018b61235e565b607196909655607294909455607392909255603c805473ffffffffffffffffffffffffffffffffffffffff9095167fffffffffffffffffffffffff000000000000000000000000000000000000000090951694909417909355604080516060810182529384526fffffffffffffffffffffffffffffffff91821660208501908152928216908401908152607480546001810182555f91909152935160029094027f19a0b39aa25ac793b5f6e9a0534364cc0b3fd1ea9b651e79c7f50a59d48ef813810194909455915191518116700100000000000000000000000000000000029116177f19a0b39aa25ac793b5f6e9a0534364cc0b3fd1ea9b651e79c7f50a59d48ef814909101555061122d95505050505050565b6040517fadc06ae700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790555f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050565b6003805461094a90612091565b5f5f5f5f5f5f868060200190518101906112f591906123cb565b6072546073546071546040805160208082018a9052818301899052606082018890526080820187905260a082019590955260c081019390935260e080840192909252805180840390920182526101009092019091528051910120949950929750909550935091507c01000000000000000000000000000000000000000000000000000000007fffff00000000000000000000000000000000000000000000000000000000000087161760016113a9826107d3565b60405160e09290921b7fffffffff0000000000000000000000000000000000000000000000000000000016602083015260248201526044810183905260640160405160208183030381529060405280519060200120975050505050505050919050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314611462576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0906020015b60405180910390a150565b6040517ff57ac68300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314611564576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f808061157384860186612434565b6071929092556072556073555050505050565b603c5473ffffffffffffffffffffffffffffffffffffffff1633146115d7576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527ffa4ea207951b027c94aa15f12ca5d0f4ca0543beb8426209ac79b49d2fc46b10906020016114d1565b60015473ffffffffffffffffffffffffffffffffffffffff16331461169b576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001545f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff1673ffffffffffffffffffffffffffffffffffffffff9092166201000081029290921790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e90602001610d0e565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314611786576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f5f8380602001905181019061179d919061245d565b604080516060810182528481526fffffffffffffffffffffffffffffffff80851660208301908152818516838501908152607480546001810182555f91909152935160029094027f19a0b39aa25ac793b5f6e9a0534364cc0b3fd1ea9b651e79c7f50a59d48ef813810194909455905190518216700100000000000000000000000000000000029116177f19a0b39aa25ac793b5f6e9a0534364cc0b3fd1ea9b651e79c7f50a59d48ef814909101555192955090935091507fee142fa601a9012bb8f08c7f3c83d2b120f9f912cd921fe80291200cbb9e46b0906118a9908590859085909283526fffffffffffffffffffffffffffffffff918216602084015216604082015260600190565b60405180910390a150505050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff16331461190d576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6906020016114d1565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff1633146119d6576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60036119e28282612247565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20816040516114d19190611cf6565b603c5473ffffffffffffffffffffffffffffffffffffffff163314611a63576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d5474010000000000000000000000000000000000000000900460ff16611ab7576040517f6f318e4c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16908190556040517fd4305f3587ac0e4b2b8b4e5b2072599efd931f81aba626cdbef4b040078b507091610d0e917401000000000000000000000000000000000000000090910460ff161515815260200190565b60748181548110611b3f575f80fd5b5f918252602090912060029091020180546001909101549091506fffffffffffffffffffffffffffffffff8082169170010000000000000000000000000000000090041683565b603c5473ffffffffffffffffffffffffffffffffffffffff163314611bd7576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d5474010000000000000000000000000000000000000000900460ff1615611c2c576040517f6f318e4c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1674010000000000000000000000000000000000000000908117918290556040517fd4305f3587ac0e4b2b8b4e5b2072599efd931f81aba626cdbef4b040078b507092610d0e92900460ff161515815260200190565b80357fffffffff0000000000000000000000000000000000000000000000000000000081168114610864575f5ffd5b5f60208284031215611ce6575f5ffd5b611cef82611ca7565b9392505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b5f5f60408385031215611d5a575f5ffd5b611d6383611ca7565b946020939093013593505050565b5f5f60208385031215611d82575f5ffd5b823567ffffffffffffffff811115611d98575f5ffd5b8301601f81018513611da8575f5ffd5b803567ffffffffffffffff811115611dbe575f5ffd5b856020828401011115611dcf575f5ffd5b6020919091019590945092505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f5f67ffffffffffffffff841115611e2657611e26611ddf565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff82111715611e7357611e73611ddf565b604052838152905080828401851015611e8a575f5ffd5b838360208301375f60208583010152509392505050565b5f60208284031215611eb1575f5ffd5b813567ffffffffffffffff811115611ec7575f5ffd5b8201601f81018413611ed7575f5ffd5b611ee684823560208401611e0c565b949350505050565b5f60208284031215611efe575f5ffd5b813567ffffffffffffffff81168114611cef575f5ffd5b73ffffffffffffffffffffffffffffffffffffffff81168114611f36575f5ffd5b50565b803561086481611f15565b5f60208284031215611f54575f5ffd5b8135611cef81611f15565b5f82601f830112611f6e575f5ffd5b611cef83833560208501611e0c565b5f5f5f5f5f5f60c08789031215611f92575f5ffd5b8635611f9d81611f15565b95506020870135611fad81611f15565b9450604087013563ffffffff81168114611fc5575f5ffd5b93506060870135611fd581611f15565b9250608087013567ffffffffffffffff811115611ff0575f5ffd5b611ffc89828a01611f5f565b92505060a087013567ffffffffffffffff811115612018575f5ffd5b61202489828a01611f5f565b9150509295509295509295565b5f60208284031215612041575f5ffd5b813567ffffffffffffffff811115612057575f5ffd5b611ee684828501611f5f565b5f60208284031215612073575f5ffd5b5035919050565b5f6020828403121561208a575f5ffd5b5051919050565b600181811c908216806120a557607f821691505b6020821081036120dc577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b6fffffffffffffffffffffffffffffffff81168114611f36575f5ffd5b8035610864816120e2565b5f5f5f5f5f5f5f5f5f5f5f5f6101808d8f031215612126575f5ffd5b61212f8d611f39565b9b5061213d60208e01611f39565b9a5061214b60408e01611f39565b995067ffffffffffffffff60608e01351115612165575f5ffd5b6121758e60608f01358f01611f5f565b985067ffffffffffffffff60808e0135111561218f575f5ffd5b61219f8e60808f01358f01611f5f565b975060a08d0135965060c08d0135955060e08d013594506101008d013593506121cb6101208e016120ff565b92506121da6101408e016120ff565b91506121e96101608e01611f39565b90509295989b509295989b509295989b565b601f82111561224257805f5260205f20601f840160051c810160208510156122205750805b601f840160051c820191505b8181101561223f575f815560010161222c565b50505b505050565b815167ffffffffffffffff81111561226157612261611ddf565b6122758161226f8454612091565b846121fb565b6020601f8211600181146122c6575f83156122905750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b17845561223f565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b8281101561231357878501518255602094850194600190920191016122f3565b508482101561234f57868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b5f5f5f5f5f5f5f60e0888a031215612374575f5ffd5b87359650602088013595506040880135945060608801359350608088013561239b816120e2565b925060a08801356123ab816120e2565b915060c08801356123bb81611f15565b8091505092959891949750929550565b5f5f5f5f5f60a086880312156123df575f5ffd5b85517fffff0000000000000000000000000000000000000000000000000000000000008116811461240e575f5ffd5b602087015160408801516060890151608090990151929a91995097965090945092505050565b5f5f5f60608486031215612446575f5ffd5b505081359360208301359350604090920135919050565b5f5f5f6060848603121561246f575f5ffd5b83516020850151909350612482816120e2565b6040850151909250612493816120e2565b80915050925092509256fea26469706673582212202a37b0342161603036ebf8f47bd31800aa464750fe419cbbb00855d4a8ce8b2964736f6c634300081c0033",
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
// Solidity: function AGGCHAIN_TYPE() view returns(uint32)
func (_Aggchainfep *AggchainfepCaller) AGGCHAINTYPE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "AGGCHAIN_TYPE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AGGCHAINTYPE is a free data retrieval call binding the contract method 0x6e7fbce9.
//
// Solidity: function AGGCHAIN_TYPE() view returns(uint32)
func (_Aggchainfep *AggchainfepSession) AGGCHAINTYPE() (uint32, error) {
	return _Aggchainfep.Contract.AGGCHAINTYPE(&_Aggchainfep.CallOpts)
}

// AGGCHAINTYPE is a free data retrieval call binding the contract method 0x6e7fbce9.
//
// Solidity: function AGGCHAIN_TYPE() view returns(uint32)
func (_Aggchainfep *AggchainfepCallerSession) AGGCHAINTYPE() (uint32, error) {
	return _Aggchainfep.Contract.AGGCHAINTYPE(&_Aggchainfep.CallOpts)
}

// AGGCHAINTYPESELECTOR is a free data retrieval call binding the contract method 0x3f17f31b.
//
// Solidity: function AGGCHAIN_TYPE_SELECTOR() view returns(bytes2)
func (_Aggchainfep *AggchainfepCaller) AGGCHAINTYPESELECTOR(opts *bind.CallOpts) ([2]byte, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "AGGCHAIN_TYPE_SELECTOR")

	if err != nil {
		return *new([2]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([2]byte)).(*[2]byte)

	return out0, err

}

// AGGCHAINTYPESELECTOR is a free data retrieval call binding the contract method 0x3f17f31b.
//
// Solidity: function AGGCHAIN_TYPE_SELECTOR() view returns(bytes2)
func (_Aggchainfep *AggchainfepSession) AGGCHAINTYPESELECTOR() ([2]byte, error) {
	return _Aggchainfep.Contract.AGGCHAINTYPESELECTOR(&_Aggchainfep.CallOpts)
}

// AGGCHAINTYPESELECTOR is a free data retrieval call binding the contract method 0x3f17f31b.
//
// Solidity: function AGGCHAIN_TYPE_SELECTOR() view returns(bytes2)
func (_Aggchainfep *AggchainfepCallerSession) AGGCHAINTYPESELECTOR() ([2]byte, error) {
	return _Aggchainfep.Contract.AGGCHAINTYPESELECTOR(&_Aggchainfep.CallOpts)
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

// ChainConfigHash is a free data retrieval call binding the contract method 0xb7fd13ce.
//
// Solidity: function chainConfigHash() view returns(bytes32)
func (_Aggchainfep *AggchainfepCaller) ChainConfigHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "chainConfigHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ChainConfigHash is a free data retrieval call binding the contract method 0xb7fd13ce.
//
// Solidity: function chainConfigHash() view returns(bytes32)
func (_Aggchainfep *AggchainfepSession) ChainConfigHash() ([32]byte, error) {
	return _Aggchainfep.Contract.ChainConfigHash(&_Aggchainfep.CallOpts)
}

// ChainConfigHash is a free data retrieval call binding the contract method 0xb7fd13ce.
//
// Solidity: function chainConfigHash() view returns(bytes32)
func (_Aggchainfep *AggchainfepCallerSession) ChainConfigHash() ([32]byte, error) {
	return _Aggchainfep.Contract.ChainConfigHash(&_Aggchainfep.CallOpts)
}

// ChainData is a free data retrieval call binding the contract method 0xe59f69c2.
//
// Solidity: function chainData(uint256 ) view returns(bytes32 lastStateRoot, uint128 timestamp, uint128 l2BlockNumber)
func (_Aggchainfep *AggchainfepCaller) ChainData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LastStateRoot [32]byte
	Timestamp     *big.Int
	L2BlockNumber *big.Int
}, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "chainData", arg0)

	outstruct := new(struct {
		LastStateRoot [32]byte
		Timestamp     *big.Int
		L2BlockNumber *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LastStateRoot = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.L2BlockNumber = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ChainData is a free data retrieval call binding the contract method 0xe59f69c2.
//
// Solidity: function chainData(uint256 ) view returns(bytes32 lastStateRoot, uint128 timestamp, uint128 l2BlockNumber)
func (_Aggchainfep *AggchainfepSession) ChainData(arg0 *big.Int) (struct {
	LastStateRoot [32]byte
	Timestamp     *big.Int
	L2BlockNumber *big.Int
}, error) {
	return _Aggchainfep.Contract.ChainData(&_Aggchainfep.CallOpts, arg0)
}

// ChainData is a free data retrieval call binding the contract method 0xe59f69c2.
//
// Solidity: function chainData(uint256 ) view returns(bytes32 lastStateRoot, uint128 timestamp, uint128 l2BlockNumber)
func (_Aggchainfep *AggchainfepCallerSession) ChainData(arg0 *big.Int) (struct {
	LastStateRoot [32]byte
	Timestamp     *big.Int
	L2BlockNumber *big.Int
}, error) {
	return _Aggchainfep.Contract.ChainData(&_Aggchainfep.CallOpts, arg0)
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
// Solidity: function getAggchainHash(bytes customChainData) view returns(bytes32)
func (_Aggchainfep *AggchainfepCaller) GetAggchainHash(opts *bind.CallOpts, customChainData []byte) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "getAggchainHash", customChainData)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAggchainHash is a free data retrieval call binding the contract method 0x6a55f66c.
//
// Solidity: function getAggchainHash(bytes customChainData) view returns(bytes32)
func (_Aggchainfep *AggchainfepSession) GetAggchainHash(customChainData []byte) ([32]byte, error) {
	return _Aggchainfep.Contract.GetAggchainHash(&_Aggchainfep.CallOpts, customChainData)
}

// GetAggchainHash is a free data retrieval call binding the contract method 0x6a55f66c.
//
// Solidity: function getAggchainHash(bytes customChainData) view returns(bytes32)
func (_Aggchainfep *AggchainfepCallerSession) GetAggchainHash(customChainData []byte) ([32]byte, error) {
	return _Aggchainfep.Contract.GetAggchainHash(&_Aggchainfep.CallOpts, customChainData)
}

// GetAggchainVKey is a free data retrieval call binding the contract method 0x01fcf6a0.
//
// Solidity: function getAggchainVKey(bytes4 aggchainSelector) view returns(bytes32 aggchainVKey)
func (_Aggchainfep *AggchainfepCaller) GetAggchainVKey(opts *bind.CallOpts, aggchainSelector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "getAggchainVKey", aggchainSelector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAggchainVKey is a free data retrieval call binding the contract method 0x01fcf6a0.
//
// Solidity: function getAggchainVKey(bytes4 aggchainSelector) view returns(bytes32 aggchainVKey)
func (_Aggchainfep *AggchainfepSession) GetAggchainVKey(aggchainSelector [4]byte) ([32]byte, error) {
	return _Aggchainfep.Contract.GetAggchainVKey(&_Aggchainfep.CallOpts, aggchainSelector)
}

// GetAggchainVKey is a free data retrieval call binding the contract method 0x01fcf6a0.
//
// Solidity: function getAggchainVKey(bytes4 aggchainSelector) view returns(bytes32 aggchainVKey)
func (_Aggchainfep *AggchainfepCallerSession) GetAggchainVKey(aggchainSelector [4]byte) ([32]byte, error) {
	return _Aggchainfep.Contract.GetAggchainVKey(&_Aggchainfep.CallOpts, aggchainSelector)
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
// Solidity: function addOwnedAggchainVKey(bytes4 aggchainSelector, bytes32 newAggchainVKey) returns()
func (_Aggchainfep *AggchainfepTransactor) AddOwnedAggchainVKey(opts *bind.TransactOpts, aggchainSelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "addOwnedAggchainVKey", aggchainSelector, newAggchainVKey)
}

// AddOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x19451a8f.
//
// Solidity: function addOwnedAggchainVKey(bytes4 aggchainSelector, bytes32 newAggchainVKey) returns()
func (_Aggchainfep *AggchainfepSession) AddOwnedAggchainVKey(aggchainSelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainfep.Contract.AddOwnedAggchainVKey(&_Aggchainfep.TransactOpts, aggchainSelector, newAggchainVKey)
}

// AddOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x19451a8f.
//
// Solidity: function addOwnedAggchainVKey(bytes4 aggchainSelector, bytes32 newAggchainVKey) returns()
func (_Aggchainfep *AggchainfepTransactorSession) AddOwnedAggchainVKey(aggchainSelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainfep.Contract.AddOwnedAggchainVKey(&_Aggchainfep.TransactOpts, aggchainSelector, newAggchainVKey)
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
// Solidity: function initialize(bytes initializeBytesCustomChain) returns()
func (_Aggchainfep *AggchainfepTransactor) Initialize(opts *bind.TransactOpts, initializeBytesCustomChain []byte) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "initialize", initializeBytesCustomChain)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializeBytesCustomChain) returns()
func (_Aggchainfep *AggchainfepSession) Initialize(initializeBytesCustomChain []byte) (*types.Transaction, error) {
	return _Aggchainfep.Contract.Initialize(&_Aggchainfep.TransactOpts, initializeBytesCustomChain)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializeBytesCustomChain) returns()
func (_Aggchainfep *AggchainfepTransactorSession) Initialize(initializeBytesCustomChain []byte) (*types.Transaction, error) {
	return _Aggchainfep.Contract.Initialize(&_Aggchainfep.TransactOpts, initializeBytesCustomChain)
}

// OnVerifyPessimistic is a paid mutator transaction binding the contract method 0x9ee4afa3.
//
// Solidity: function onVerifyPessimistic(bytes customData) returns()
func (_Aggchainfep *AggchainfepTransactor) OnVerifyPessimistic(opts *bind.TransactOpts, customData []byte) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "onVerifyPessimistic", customData)
}

// OnVerifyPessimistic is a paid mutator transaction binding the contract method 0x9ee4afa3.
//
// Solidity: function onVerifyPessimistic(bytes customData) returns()
func (_Aggchainfep *AggchainfepSession) OnVerifyPessimistic(customData []byte) (*types.Transaction, error) {
	return _Aggchainfep.Contract.OnVerifyPessimistic(&_Aggchainfep.TransactOpts, customData)
}

// OnVerifyPessimistic is a paid mutator transaction binding the contract method 0x9ee4afa3.
//
// Solidity: function onVerifyPessimistic(bytes customData) returns()
func (_Aggchainfep *AggchainfepTransactorSession) OnVerifyPessimistic(customData []byte) (*types.Transaction, error) {
	return _Aggchainfep.Contract.OnVerifyPessimistic(&_Aggchainfep.TransactOpts, customData)
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

// UpdateCustomInitializeData is a paid mutator transaction binding the contract method 0x7e005eb2.
//
// Solidity: function updateCustomInitializeData(bytes customInitializeData) returns()
func (_Aggchainfep *AggchainfepTransactor) UpdateCustomInitializeData(opts *bind.TransactOpts, customInitializeData []byte) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "updateCustomInitializeData", customInitializeData)
}

// UpdateCustomInitializeData is a paid mutator transaction binding the contract method 0x7e005eb2.
//
// Solidity: function updateCustomInitializeData(bytes customInitializeData) returns()
func (_Aggchainfep *AggchainfepSession) UpdateCustomInitializeData(customInitializeData []byte) (*types.Transaction, error) {
	return _Aggchainfep.Contract.UpdateCustomInitializeData(&_Aggchainfep.TransactOpts, customInitializeData)
}

// UpdateCustomInitializeData is a paid mutator transaction binding the contract method 0x7e005eb2.
//
// Solidity: function updateCustomInitializeData(bytes customInitializeData) returns()
func (_Aggchainfep *AggchainfepTransactorSession) UpdateCustomInitializeData(customInitializeData []byte) (*types.Transaction, error) {
	return _Aggchainfep.Contract.UpdateCustomInitializeData(&_Aggchainfep.TransactOpts, customInitializeData)
}

// UpdateOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x314eb17b.
//
// Solidity: function updateOwnedAggchainVKey(bytes4 aggchainSelector, bytes32 updatedAggchainVKey) returns()
func (_Aggchainfep *AggchainfepTransactor) UpdateOwnedAggchainVKey(opts *bind.TransactOpts, aggchainSelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "updateOwnedAggchainVKey", aggchainSelector, updatedAggchainVKey)
}

// UpdateOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x314eb17b.
//
// Solidity: function updateOwnedAggchainVKey(bytes4 aggchainSelector, bytes32 updatedAggchainVKey) returns()
func (_Aggchainfep *AggchainfepSession) UpdateOwnedAggchainVKey(aggchainSelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainfep.Contract.UpdateOwnedAggchainVKey(&_Aggchainfep.TransactOpts, aggchainSelector, updatedAggchainVKey)
}

// UpdateOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x314eb17b.
//
// Solidity: function updateOwnedAggchainVKey(bytes4 aggchainSelector, bytes32 updatedAggchainVKey) returns()
func (_Aggchainfep *AggchainfepTransactorSession) UpdateOwnedAggchainVKey(aggchainSelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainfep.Contract.UpdateOwnedAggchainVKey(&_Aggchainfep.TransactOpts, aggchainSelector, updatedAggchainVKey)
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
	NewVKeyManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAcceptVKeyManagerRole is a free log retrieval operation binding the contract event 0x75ccbacededf0b9fc9371bb4cf651be4f96efab2c519a529cfd90e14c9d49ad2.
//
// Solidity: event AcceptVKeyManagerRole(address newVKeyManager)
func (_Aggchainfep *AggchainfepFilterer) FilterAcceptVKeyManagerRole(opts *bind.FilterOpts) (*AggchainfepAcceptVKeyManagerRoleIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "AcceptVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainfepAcceptVKeyManagerRoleIterator{contract: _Aggchainfep.contract, event: "AcceptVKeyManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptVKeyManagerRole is a free log subscription operation binding the contract event 0x75ccbacededf0b9fc9371bb4cf651be4f96efab2c519a529cfd90e14c9d49ad2.
//
// Solidity: event AcceptVKeyManagerRole(address newVKeyManager)
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

// ParseAcceptVKeyManagerRole is a log parse operation binding the contract event 0x75ccbacededf0b9fc9371bb4cf651be4f96efab2c519a529cfd90e14c9d49ad2.
//
// Solidity: event AcceptVKeyManagerRole(address newVKeyManager)
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

// AggchainfepOnVerifyPessimisticIterator is returned from FilterOnVerifyPessimistic and is used to iterate over the raw logs and unpacked data for OnVerifyPessimistic events raised by the Aggchainfep contract.
type AggchainfepOnVerifyPessimisticIterator struct {
	Event *AggchainfepOnVerifyPessimistic // Event containing the contract specifics and raw log

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
func (it *AggchainfepOnVerifyPessimisticIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepOnVerifyPessimistic)
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
		it.Event = new(AggchainfepOnVerifyPessimistic)
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
func (it *AggchainfepOnVerifyPessimisticIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepOnVerifyPessimisticIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepOnVerifyPessimistic represents a OnVerifyPessimistic event raised by the Aggchainfep contract.
type AggchainfepOnVerifyPessimistic struct {
	InitStateRoot     [32]byte
	InitTimestamp     *big.Int
	InitL2BlockNumber *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOnVerifyPessimistic is a free log retrieval operation binding the contract event 0xee142fa601a9012bb8f08c7f3c83d2b120f9f912cd921fe80291200cbb9e46b0.
//
// Solidity: event OnVerifyPessimistic(bytes32 initStateRoot, uint128 initTimestamp, uint128 initL2BlockNumber)
func (_Aggchainfep *AggchainfepFilterer) FilterOnVerifyPessimistic(opts *bind.FilterOpts) (*AggchainfepOnVerifyPessimisticIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "OnVerifyPessimistic")
	if err != nil {
		return nil, err
	}
	return &AggchainfepOnVerifyPessimisticIterator{contract: _Aggchainfep.contract, event: "OnVerifyPessimistic", logs: logs, sub: sub}, nil
}

// WatchOnVerifyPessimistic is a free log subscription operation binding the contract event 0xee142fa601a9012bb8f08c7f3c83d2b120f9f912cd921fe80291200cbb9e46b0.
//
// Solidity: event OnVerifyPessimistic(bytes32 initStateRoot, uint128 initTimestamp, uint128 initL2BlockNumber)
func (_Aggchainfep *AggchainfepFilterer) WatchOnVerifyPessimistic(opts *bind.WatchOpts, sink chan<- *AggchainfepOnVerifyPessimistic) (event.Subscription, error) {

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "OnVerifyPessimistic")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepOnVerifyPessimistic)
				if err := _Aggchainfep.contract.UnpackLog(event, "OnVerifyPessimistic", log); err != nil {
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

// ParseOnVerifyPessimistic is a log parse operation binding the contract event 0xee142fa601a9012bb8f08c7f3c83d2b120f9f912cd921fe80291200cbb9e46b0.
//
// Solidity: event OnVerifyPessimistic(bytes32 initStateRoot, uint128 initTimestamp, uint128 initL2BlockNumber)
func (_Aggchainfep *AggchainfepFilterer) ParseOnVerifyPessimistic(log types.Log) (*AggchainfepOnVerifyPessimistic, error) {
	event := new(AggchainfepOnVerifyPessimistic)
	if err := _Aggchainfep.contract.UnpackLog(event, "OnVerifyPessimistic", log); err != nil {
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
	NewVKeyManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTransferVKeyManagerRole is a free log retrieval operation binding the contract event 0xfa4ea207951b027c94aa15f12ca5d0f4ca0543beb8426209ac79b49d2fc46b10.
//
// Solidity: event TransferVKeyManagerRole(address newVKeyManager)
func (_Aggchainfep *AggchainfepFilterer) FilterTransferVKeyManagerRole(opts *bind.FilterOpts) (*AggchainfepTransferVKeyManagerRoleIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "TransferVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainfepTransferVKeyManagerRoleIterator{contract: _Aggchainfep.contract, event: "TransferVKeyManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferVKeyManagerRole is a free log subscription operation binding the contract event 0xfa4ea207951b027c94aa15f12ca5d0f4ca0543beb8426209ac79b49d2fc46b10.
//
// Solidity: event TransferVKeyManagerRole(address newVKeyManager)
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

// ParseTransferVKeyManagerRole is a log parse operation binding the contract event 0xfa4ea207951b027c94aa15f12ca5d0f4ca0543beb8426209ac79b49d2fc46b10.
//
// Solidity: event TransferVKeyManagerRole(address newVKeyManager)
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
	Selector        [4]byte
	NewAggchainVKey [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateAggchainVKey is a free log retrieval operation binding the contract event 0x2ae81c794f274593f3cd22db9ad3af161ff9d93d15922b35629a6ffd9be7243f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Aggchainfep *AggchainfepFilterer) FilterUpdateAggchainVKey(opts *bind.FilterOpts) (*AggchainfepUpdateAggchainVKeyIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "UpdateAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &AggchainfepUpdateAggchainVKeyIterator{contract: _Aggchainfep.contract, event: "UpdateAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchUpdateAggchainVKey is a free log subscription operation binding the contract event 0x2ae81c794f274593f3cd22db9ad3af161ff9d93d15922b35629a6ffd9be7243f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
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

// ParseUpdateAggchainVKey is a log parse operation binding the contract event 0x2ae81c794f274593f3cd22db9ad3af161ff9d93d15922b35629a6ffd9be7243f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Aggchainfep *AggchainfepFilterer) ParseUpdateAggchainVKey(log types.Log) (*AggchainfepUpdateAggchainVKey, error) {
	event := new(AggchainfepUpdateAggchainVKey)
	if err := _Aggchainfep.contract.UnpackLog(event, "UpdateAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepUpdateUseDefaultGatewayFlagIterator is returned from FilterUpdateUseDefaultGatewayFlag and is used to iterate over the raw logs and unpacked data for UpdateUseDefaultGatewayFlag events raised by the Aggchainfep contract.
type AggchainfepUpdateUseDefaultGatewayFlagIterator struct {
	Event *AggchainfepUpdateUseDefaultGatewayFlag // Event containing the contract specifics and raw log

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
func (it *AggchainfepUpdateUseDefaultGatewayFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepUpdateUseDefaultGatewayFlag)
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
		it.Event = new(AggchainfepUpdateUseDefaultGatewayFlag)
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
func (it *AggchainfepUpdateUseDefaultGatewayFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepUpdateUseDefaultGatewayFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepUpdateUseDefaultGatewayFlag represents a UpdateUseDefaultGatewayFlag event raised by the Aggchainfep contract.
type AggchainfepUpdateUseDefaultGatewayFlag struct {
	UseDefaultGateway bool
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdateUseDefaultGatewayFlag is a free log retrieval operation binding the contract event 0xd4305f3587ac0e4b2b8b4e5b2072599efd931f81aba626cdbef4b040078b5070.
//
// Solidity: event UpdateUseDefaultGatewayFlag(bool useDefaultGateway)
func (_Aggchainfep *AggchainfepFilterer) FilterUpdateUseDefaultGatewayFlag(opts *bind.FilterOpts) (*AggchainfepUpdateUseDefaultGatewayFlagIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "UpdateUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return &AggchainfepUpdateUseDefaultGatewayFlagIterator{contract: _Aggchainfep.contract, event: "UpdateUseDefaultGatewayFlag", logs: logs, sub: sub}, nil
}

// WatchUpdateUseDefaultGatewayFlag is a free log subscription operation binding the contract event 0xd4305f3587ac0e4b2b8b4e5b2072599efd931f81aba626cdbef4b040078b5070.
//
// Solidity: event UpdateUseDefaultGatewayFlag(bool useDefaultGateway)
func (_Aggchainfep *AggchainfepFilterer) WatchUpdateUseDefaultGatewayFlag(opts *bind.WatchOpts, sink chan<- *AggchainfepUpdateUseDefaultGatewayFlag) (event.Subscription, error) {

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "UpdateUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepUpdateUseDefaultGatewayFlag)
				if err := _Aggchainfep.contract.UnpackLog(event, "UpdateUseDefaultGatewayFlag", log); err != nil {
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
func (_Aggchainfep *AggchainfepFilterer) ParseUpdateUseDefaultGatewayFlag(log types.Log) (*AggchainfepUpdateUseDefaultGatewayFlag, error) {
	event := new(AggchainfepUpdateUseDefaultGatewayFlag)
	if err := _Aggchainfep.contract.UnpackLog(event, "UpdateUseDefaultGatewayFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

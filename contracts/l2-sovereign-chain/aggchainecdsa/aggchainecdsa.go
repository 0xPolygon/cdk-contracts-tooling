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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"_rollupManager\",\"type\":\"address\"},{\"internalType\":\"contractIAggLayerGateway\",\"name\":\"_aggLayerGateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AggchainRouteAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainRouteNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalAccInputHashDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggchainVKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultGatewayAlreadySet\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"AddAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"OnVerifyPessimistic\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"UpdateAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"useDefaultGateway\",\"type\":\"bool\"}],\"name\":\"UpdateUseDefaultGatewayFlag\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AGGCHAIN_TYPE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainSelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"aggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"addAggchainRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggLayerGateway\",\"outputs\":[{\"internalType\":\"contractIAggLayerGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"aggchainVKeyRoutes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"aggchainVKey\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableUseDefaultGatewayFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableUseDefaultGatewayFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"customChainData\",\"type\":\"bytes\"}],\"name\":\"getAggchainHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainSelector\",\"type\":\"bytes4\"}],\"name\":\"getAggchainVKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"aggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initializeBytesCustomChain\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"customChainData\",\"type\":\"bytes\"}],\"name\":\"onVerifyPessimistic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainSelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"updatedAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"updateAggchainRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useDefaultGateway\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x610120604052603c805460ff1916600117905534801561001d575f5ffd5b50604051611ce5380380611ce583398101604081905261003c91610160565b6001600160a01b0380861660a05280851660805280841660c052821660e05284848484848484848461006c61008c565b505050506001600160a01b031661010052506101d1975050505050505050565b5f54610100900460ff16156100f75760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161015610147575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b038116811461015d575f5ffd5b50565b5f5f5f5f5f60a08688031215610174575f5ffd5b855161017f81610149565b602087015190955061019081610149565b60408701519094506101a181610149565b60608701519093506101b281610149565b60808701519092506101c381610149565b809150509295509295909350565b60805160a05160c05160e05161010051611ac26102235f395f8181610487015261068101525f818161036f015281816107950152610d3b01525f61046001525f61052f01525f6105710152611ac25ff3fe608060405234801561000f575f5ffd5b506004361061021b575f3560e01c80638c3d730111610123578063cfa8ed47116100b8578063e46761c411610088578063e7a7ed021161006e578063e7a7ed021461059b578063f851a440146105af578063ff904079146105d4575f5ffd5b8063e46761c41461056c578063e631476c14610593575f5ffd5b8063cfa8ed471461050a578063d02103ca1461052a578063dc8c424914610551578063df4970a814610559575f5ffd5b8063ab27f141116100f3578063ab27f141146104a9578063ada8f919146104bc578063c754c7ed146104cf578063c89e42df146104f7575f5ffd5b80638c3d7301146104405780639ee4afa314610448578063a3c573eb1461045b578063ab0475cf14610482575f5ffd5b806349b7b802116101b35780636b8616ce116101835780636e7fbce9116101695780636e7fbce9146104125780636ff512cc1461041a578063712570221461042d575f5ffd5b80636b8616ce146103ea5780636e05d2cd14610409575f5ffd5b806349b7b8021461036a578063542028d51461039157806369d8f277146103995780636a55f66c146103d7575f5ffd5b80633c351e10116101ee5780633c351e10146102bf5780633cbc795b146102df578063439fab911461031c5780634560526714610331575f5ffd5b806301fcf6a01461021f578063107bf28c14610245578063267822471461025a5780632c111c061461029f575b5f5ffd5b61023261022d3660046113c4565b6105f1565b6040519081526020015b60405180910390f35b61024d610707565b60405161023c91906113e4565b60015461027a9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161023c565b60085461027a9073ffffffffffffffffffffffffffffffffffffffff1681565b60095461027a9073ffffffffffffffffffffffffffffffffffffffff1681565b6009546103079074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff909116815260200161023c565b61032f61032a366004611533565b610793565b005b6007546103519068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff909116815260200161023c565b61027a7f000000000000000000000000000000000000000000000000000000000000000081565b61024d610a5a565b6103c26103a73660046113c4565b603d6020525f90815260409020805460019091015460ff1682565b6040805192835290151560208301520161023c565b6102326103e5366004611533565b610a67565b6102326103f8366004611580565b60066020525f908152604090205481565b61023260055481565b610307600181565b61032f6104283660046115cb565b610b64565b61032f61043b366004611604565b610c34565b61032f610c66565b61032f6104563660046116b8565b610d39565b61027a7f000000000000000000000000000000000000000000000000000000000000000081565b61027a7f000000000000000000000000000000000000000000000000000000000000000081565b61032f6104b7366004611726565b610df5565b61032f6104ca3660046115cb565b610f0e565b60075461035190700100000000000000000000000000000000900467ffffffffffffffff1681565b61032f61050536600461174e565b610fd7565b60025461027a9073ffffffffffffffffffffffffffffffffffffffff1681565b61027a7f000000000000000000000000000000000000000000000000000000000000000081565b61032f611069565b61032f610567366004611726565b61115a565b61027a7f000000000000000000000000000000000000000000000000000000000000000081565b61032f6112ab565b6007546103519067ffffffffffffffff1681565b5f5461027a9062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b603c546105e19060ff1681565b604051901515815260200161023c565b603c545f9060ff1661063157507fffffffff00000000000000000000000000000000000000000000000000000000165f908152603d602052604090205490565b6040517f6cabfdab0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000831660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690636cabfdab90602401602060405180830381865afa1580156106db573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106ff9190611780565b90505b919050565b6004805461071490611797565b80601f016020809104026020016040519081016040528092919081815260200182805461074090611797565b801561078b5780601f106107625761010080835404028352916020019161078b565b820191905f5260205f20905b81548152906001019060200180831161076e57829003601f168201915b505050505081565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314610802576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f54610100900460ff161580801561082057505f54600160ff909116105b806108395750303b15801561083957505f5460ff166001145b6108c9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840160405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610925575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b5f5f5f5f5f8680602001905181019061093e9190611835565b5f805473ffffffffffffffffffffffffffffffffffffffff80881662010000027fffffffffffffffffffff0000000000000000000000000000000000000000ffff90921691909117909155600980548286167fffffffffffffffffffffffff000000000000000000000000000000000000000091821617909155600280549287169290911691909117905593985091965094509250905060036109e1838261191f565b5060046109ee828261191f565b5050505050508015610a56575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b6003805461071490611797565b5f5f82806020019051810190610a7d9190611a36565b9050601081901c7dffff00000000000000000000000000000000000000000000000000000000166001610aaf826105f1565b60025460405160609190911b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602082015260340160405160208183030381529060405280519060200120604051602001610b459392919060e09390931b7fffffffff000000000000000000000000000000000000000000000000000000001683526004830191909152602482015260440190565b6040516020818303038152906040528051906020012092505050919050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314610bba576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0906020015b60405180910390a150565b6040517ff57ac68300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60015473ffffffffffffffffffffffffffffffffffffffff163314610cb7576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001545f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff1673ffffffffffffffffffffffffffffffffffffffff9092166201000081029290921790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e906020015b60405180910390a1565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314610da8576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f610db582840184611a75565b90507f18eea7739a19b6e71e562030f68e6850aa7f1408136e6795b2690501810c410881604051610de891815260200190565b60405180910390a1505050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314610e4b576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f908152603d602052604090208054610eb3576040517f6310104a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b818155604080517fffffffff0000000000000000000000000000000000000000000000000000000085168152602081018490527f2ae81c794f274593f3cd22db9ad3af161ff9d93d15922b35629a6ffd9be7243f9101610de8565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314610f64576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001610c29565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff16331461102d576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6003611039828261191f565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b2081604051610c2991906113e4565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff1633146110bf576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603c5460ff16156110fc576040517f6f318e4c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603c80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660019081179091556040519081527fd4305f3587ac0e4b2b8b4e5b2072599efd931f81aba626cdbef4b040078b507090602001610d2f565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff1633146111b0576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806111e7576040517f4aac8b8800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f908152603d60205260409020805415611250576040517f02eb6b2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b818155604080517fffffffff0000000000000000000000000000000000000000000000000000000085168152602081018490527f6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a799101610de8565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314611301576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603c5460ff1661133d576040517f6f318e4c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603c80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556040515f81527fd4305f3587ac0e4b2b8b4e5b2072599efd931f81aba626cdbef4b040078b507090602001610d2f565b80357fffffffff0000000000000000000000000000000000000000000000000000000081168114610702575f5ffd5b5f602082840312156113d4575f5ffd5b6113dd82611395565b9392505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156114ab576114ab611437565b604052919050565b5f67ffffffffffffffff8211156114cc576114cc611437565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b5f61150a611505846114b3565b611464565b905082815283838301111561151d575f5ffd5b828260208301375f602084830101529392505050565b5f60208284031215611543575f5ffd5b813567ffffffffffffffff811115611559575f5ffd5b8201601f81018413611569575f5ffd5b611578848235602084016114f8565b949350505050565b5f60208284031215611590575f5ffd5b813567ffffffffffffffff811681146113dd575f5ffd5b73ffffffffffffffffffffffffffffffffffffffff811681146115c8575f5ffd5b50565b5f602082840312156115db575f5ffd5b81356113dd816115a7565b5f82601f8301126115f5575f5ffd5b6113dd838335602085016114f8565b5f5f5f5f5f5f60c08789031215611619575f5ffd5b8635611624816115a7565b95506020870135611634816115a7565b9450604087013563ffffffff8116811461164c575f5ffd5b9350606087013561165c816115a7565b9250608087013567ffffffffffffffff811115611677575f5ffd5b61168389828a016115e6565b92505060a087013567ffffffffffffffff81111561169f575f5ffd5b6116ab89828a016115e6565b9150509295509295509295565b5f5f602083850312156116c9575f5ffd5b823567ffffffffffffffff8111156116df575f5ffd5b8301601f810185136116ef575f5ffd5b803567ffffffffffffffff811115611705575f5ffd5b856020828401011115611716575f5ffd5b6020919091019590945092505050565b5f5f60408385031215611737575f5ffd5b61174083611395565b946020939093013593505050565b5f6020828403121561175e575f5ffd5b813567ffffffffffffffff811115611774575f5ffd5b611578848285016115e6565b5f60208284031215611790575f5ffd5b5051919050565b600181811c908216806117ab57607f821691505b6020821081036117e2577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b5f82601f8301126117f7575f5ffd5b8151611805611505826114b3565b818152846020838601011115611819575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f5f5f5f5f60a08688031215611849575f5ffd5b8551611854816115a7565b6020870151909550611865816115a7565b6040870151909450611876816115a7565b606087015190935067ffffffffffffffff811115611892575f5ffd5b61189e888289016117e8565b925050608086015167ffffffffffffffff8111156118ba575f5ffd5b6118c6888289016117e8565b9150509295509295909350565b601f82111561191a57805f5260205f20601f840160051c810160208510156118f85750805b601f840160051c820191505b81811015611917575f8155600101611904565b50505b505050565b815167ffffffffffffffff81111561193957611939611437565b61194d816119478454611797565b846118d3565b6020601f82116001811461199e575f83156119685750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455611917565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b828110156119eb57878501518255602094850194600190920191016119cb565b5084821015611a2757868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b5f60208284031215611a46575f5ffd5b81517fffff000000000000000000000000000000000000000000000000000000000000811681146113dd575f5ffd5b5f60208284031215611a85575f5ffd5b503591905056fea264697066735822122039c908cf67352b33e4657b3ac2941d5b29a562fdf3cad5810158d3227e01715664736f6c634300081c0033",
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

// AggchainVKeyRoutes is a free data retrieval call binding the contract method 0x69d8f277.
//
// Solidity: function aggchainVKeyRoutes(bytes4 aggchainVKeySelector) view returns(bytes32 aggchainVKey, bool frozen)
func (_Aggchainecdsa *AggchainecdsaCaller) AggchainVKeyRoutes(opts *bind.CallOpts, aggchainVKeySelector [4]byte) (struct {
	AggchainVKey [32]byte
	Frozen       bool
}, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "aggchainVKeyRoutes", aggchainVKeySelector)

	outstruct := new(struct {
		AggchainVKey [32]byte
		Frozen       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AggchainVKey = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Frozen = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// AggchainVKeyRoutes is a free data retrieval call binding the contract method 0x69d8f277.
//
// Solidity: function aggchainVKeyRoutes(bytes4 aggchainVKeySelector) view returns(bytes32 aggchainVKey, bool frozen)
func (_Aggchainecdsa *AggchainecdsaSession) AggchainVKeyRoutes(aggchainVKeySelector [4]byte) (struct {
	AggchainVKey [32]byte
	Frozen       bool
}, error) {
	return _Aggchainecdsa.Contract.AggchainVKeyRoutes(&_Aggchainecdsa.CallOpts, aggchainVKeySelector)
}

// AggchainVKeyRoutes is a free data retrieval call binding the contract method 0x69d8f277.
//
// Solidity: function aggchainVKeyRoutes(bytes4 aggchainVKeySelector) view returns(bytes32 aggchainVKey, bool frozen)
func (_Aggchainecdsa *AggchainecdsaCallerSession) AggchainVKeyRoutes(aggchainVKeySelector [4]byte) (struct {
	AggchainVKey [32]byte
	Frozen       bool
}, error) {
	return _Aggchainecdsa.Contract.AggchainVKeyRoutes(&_Aggchainecdsa.CallOpts, aggchainVKeySelector)
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
// Solidity: function getAggchainHash(bytes customChainData) view returns(bytes32)
func (_Aggchainecdsa *AggchainecdsaCaller) GetAggchainHash(opts *bind.CallOpts, customChainData []byte) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainecdsa.contract.Call(opts, &out, "getAggchainHash", customChainData)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAggchainHash is a free data retrieval call binding the contract method 0x6a55f66c.
//
// Solidity: function getAggchainHash(bytes customChainData) view returns(bytes32)
func (_Aggchainecdsa *AggchainecdsaSession) GetAggchainHash(customChainData []byte) ([32]byte, error) {
	return _Aggchainecdsa.Contract.GetAggchainHash(&_Aggchainecdsa.CallOpts, customChainData)
}

// GetAggchainHash is a free data retrieval call binding the contract method 0x6a55f66c.
//
// Solidity: function getAggchainHash(bytes customChainData) view returns(bytes32)
func (_Aggchainecdsa *AggchainecdsaCallerSession) GetAggchainHash(customChainData []byte) ([32]byte, error) {
	return _Aggchainecdsa.Contract.GetAggchainHash(&_Aggchainecdsa.CallOpts, customChainData)
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

// AddAggchainRoute is a paid mutator transaction binding the contract method 0xdf4970a8.
//
// Solidity: function addAggchainRoute(bytes4 aggchainSelector, bytes32 aggchainVKey) returns()
func (_Aggchainecdsa *AggchainecdsaTransactor) AddAggchainRoute(opts *bind.TransactOpts, aggchainSelector [4]byte, aggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainecdsa.contract.Transact(opts, "addAggchainRoute", aggchainSelector, aggchainVKey)
}

// AddAggchainRoute is a paid mutator transaction binding the contract method 0xdf4970a8.
//
// Solidity: function addAggchainRoute(bytes4 aggchainSelector, bytes32 aggchainVKey) returns()
func (_Aggchainecdsa *AggchainecdsaSession) AddAggchainRoute(aggchainSelector [4]byte, aggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.AddAggchainRoute(&_Aggchainecdsa.TransactOpts, aggchainSelector, aggchainVKey)
}

// AddAggchainRoute is a paid mutator transaction binding the contract method 0xdf4970a8.
//
// Solidity: function addAggchainRoute(bytes4 aggchainSelector, bytes32 aggchainVKey) returns()
func (_Aggchainecdsa *AggchainecdsaTransactorSession) AddAggchainRoute(aggchainSelector [4]byte, aggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.AddAggchainRoute(&_Aggchainecdsa.TransactOpts, aggchainSelector, aggchainVKey)
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
// Solidity: function onVerifyPessimistic(bytes customChainData) returns()
func (_Aggchainecdsa *AggchainecdsaTransactor) OnVerifyPessimistic(opts *bind.TransactOpts, customChainData []byte) (*types.Transaction, error) {
	return _Aggchainecdsa.contract.Transact(opts, "onVerifyPessimistic", customChainData)
}

// OnVerifyPessimistic is a paid mutator transaction binding the contract method 0x9ee4afa3.
//
// Solidity: function onVerifyPessimistic(bytes customChainData) returns()
func (_Aggchainecdsa *AggchainecdsaSession) OnVerifyPessimistic(customChainData []byte) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.OnVerifyPessimistic(&_Aggchainecdsa.TransactOpts, customChainData)
}

// OnVerifyPessimistic is a paid mutator transaction binding the contract method 0x9ee4afa3.
//
// Solidity: function onVerifyPessimistic(bytes customChainData) returns()
func (_Aggchainecdsa *AggchainecdsaTransactorSession) OnVerifyPessimistic(customChainData []byte) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.OnVerifyPessimistic(&_Aggchainecdsa.TransactOpts, customChainData)
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

// UpdateAggchainRoute is a paid mutator transaction binding the contract method 0xab27f141.
//
// Solidity: function updateAggchainRoute(bytes4 aggchainSelector, bytes32 updatedAggchainVKey) returns()
func (_Aggchainecdsa *AggchainecdsaTransactor) UpdateAggchainRoute(opts *bind.TransactOpts, aggchainSelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainecdsa.contract.Transact(opts, "updateAggchainRoute", aggchainSelector, updatedAggchainVKey)
}

// UpdateAggchainRoute is a paid mutator transaction binding the contract method 0xab27f141.
//
// Solidity: function updateAggchainRoute(bytes4 aggchainSelector, bytes32 updatedAggchainVKey) returns()
func (_Aggchainecdsa *AggchainecdsaSession) UpdateAggchainRoute(aggchainSelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.UpdateAggchainRoute(&_Aggchainecdsa.TransactOpts, aggchainSelector, updatedAggchainVKey)
}

// UpdateAggchainRoute is a paid mutator transaction binding the contract method 0xab27f141.
//
// Solidity: function updateAggchainRoute(bytes4 aggchainSelector, bytes32 updatedAggchainVKey) returns()
func (_Aggchainecdsa *AggchainecdsaTransactorSession) UpdateAggchainRoute(aggchainSelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainecdsa.Contract.UpdateAggchainRoute(&_Aggchainecdsa.TransactOpts, aggchainSelector, updatedAggchainVKey)
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

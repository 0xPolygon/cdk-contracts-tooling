// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonrollupbaseetrogprevious

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

// PolygonRollupBaseEtrogPreviousBatchData is an auto generated low-level Go binding around an user-defined struct.
type PolygonRollupBaseEtrogPreviousBatchData struct {
	Transactions         []byte
	ForcedGlobalExitRoot [32]byte
	ForcedTimestamp      uint64
	ForcedBlockHashL1    [32]byte
}

// PolygonrollupbaseetrogpreviousMetaData contains all meta data concerning the Polygonrollupbaseetrogprevious contract.
var PolygonrollupbaseetrogpreviousMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"_rollupManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"forceBatchNum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"}],\"name\":\"ForceBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"}],\"name\":\"InitialSequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"l1InfoRoot\",\"type\":\"bytes32\"}],\"name\":\"SequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newForceBatchAddress\",\"type\":\"address\"}],\"name\":\"SetForceBatchAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"SetForceBatchTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatches\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GLOBAL_EXIT_ROOT_MANAGER_L2\",\"outputs\":[{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_LIST_LEN_LEN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_DATA_LEN_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_EFFECTIVE_PERCENTAGE\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_R\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_S\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_V\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculatePolPerForceBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"polAmount\",\"type\":\"uint256\"}],\"name\":\"forceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"}],\"name\":\"generateInitializeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"onVerifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrogPrevious.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"}],\"name\":\"sequenceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrogPrevious.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"}],\"name\":\"sequenceForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newForceBatchAddress\",\"type\":\"address\"}],\"name\":\"setForceBatchAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"setForceBatchTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x61010060405234801562000011575f80fd5b50604051620042443803806200424483398101604081905262000034916200006f565b6001600160a01b0393841660a052918316608052821660c0521660e052620000d4565b6001600160a01b03811681146200006c575f80fd5b50565b5f805f806080858703121562000083575f80fd5b8451620000908162000057565b6020860151909450620000a38162000057565b6040860151909350620000b68162000057565b6060860151909250620000c98162000057565b939692955090935050565b60805160a05160c05160e05161408d620001b75f395f81816104de0152818161094201528181610aae01528181610cf6015281816112c80152818161176301528181611bb201528181611ca701528181612323015281816123e801528181612cf201528181612d6b01528181612d8d0152612ea201525f818161063701528181610eea01528181610fbf01528181611e7101528181611f79015261283d01525f81816106f30152818161113f0152818161253001526128b901525f8181610725015281816107f501528181611bfb015281816125050152612e39015261408d5ff3fe608060405234801561000f575f80fd5b50600436106102cc575f3560e01c80637a5460c51161017c578063c7fffd4b116100dd578063e46761c411610093578063ecef3f991161006e578063ecef3f991461076e578063f35dda4714610781578063f851a44014610789575f80fd5b8063e46761c414610720578063e7a7ed0214610747578063eaeb077b1461075b575f80fd5b8063cfa8ed47116100c3578063cfa8ed47146106ce578063d02103ca146106ee578063d7bc90ff14610715575f80fd5b8063c7fffd4b146106b3578063c89e42df146106bb575f80fd5b8063a3c573eb11610132578063ada8f91911610118578063ada8f9191461066c578063b0afe1541461067f578063c754c7ed1461068b575f80fd5b8063a3c573eb14610632578063a652f26c14610659575f80fd5b806391cafe321161016257806391cafe32146105f15780639e001877146106045780639f26f8401461061f575f80fd5b80637a5460c5146105ad5780638c3d7301146105e9575f80fd5b806340b5de6c11610231578063542028d5116101e75780636e05d2cd116101c25780636e05d2cd1461057e5780636ff512cc14610587578063712570221461059a575f80fd5b8063542028d51461054f578063676870d2146105575780636b8616ce1461055f575f80fd5b806349b7b8021161021757806349b7b802146104d95780634e4877061461050057806352bdeb6d14610513575f80fd5b806340b5de6c1461044857806345605267146104a0575f80fd5b8063267822471161028657806332c2d1531161026c57806332c2d153146103d65780633c351e10146103eb5780633cbc795b1461040b575f80fd5b806326782247146103715780632c111c06146103b6575f80fd5b806305835f37116102b657806305835f3714610306578063107bf28c1461034f57806311e892d414610357575f80fd5b8062d0295d146102d057806303508963146102eb575b5f80fd5b6102d86107ae565b6040519081526020015b60405180910390f35b6102f3602081565b60405161ffff90911681526020016102e2565b6103426040518060400160405280600881526020017f80808401c9c3809400000000000000000000000000000000000000000000000081525081565b6040516102e29190613392565b6103426108b4565b61035f60f981565b60405160ff90911681526020016102e2565b6001546103919073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102e2565b6008546103919073ffffffffffffffffffffffffffffffffffffffff1681565b6103e96103e43660046133e4565b610940565b005b6009546103919073ffffffffffffffffffffffffffffffffffffffff1681565b6009546104339074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff90911681526020016102e2565b61046f7fff0000000000000000000000000000000000000000000000000000000000000081565b6040517fff0000000000000000000000000000000000000000000000000000000000000090911681526020016102e2565b6007546104c09068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020016102e2565b6103917f000000000000000000000000000000000000000000000000000000000000000081565b6103e961050e366004613423565b610a0f565b6103426040518060400160405280600281526020017f80b800000000000000000000000000000000000000000000000000000000000081525081565b610342610c1e565b6102f3601f81565b6102d861056d366004613423565b60066020525f908152604090205481565b6102d860055481565b6103e961059536600461343e565b610c2b565b6103e96105a836600461357d565b610cf4565b6103426040518060400160405280600281526020017f80b900000000000000000000000000000000000000000000000000000000000081525081565b6103e96114fd565b6103e96105ff36600461343e565b6115cf565b61039173a40d5f56745a118d0906a34e69aec8c0db1cb8fa81565b6103e961062d36600461366c565b6116e7565b6103917f000000000000000000000000000000000000000000000000000000000000000081565b6103426106673660046136ab565b611d73565b6103e961067a36600461343e565b612151565b6102d86405ca1ab1e081565b6007546104c090700100000000000000000000000000000000900467ffffffffffffffff1681565b61035f60e481565b6103e96106c936600461371c565b61221a565b6002546103919073ffffffffffffffffffffffffffffffffffffffff1681565b6103917f000000000000000000000000000000000000000000000000000000000000000081565b6102d8635ca1ab1e81565b6103917f000000000000000000000000000000000000000000000000000000000000000081565b6007546104c09067ffffffffffffffff1681565b6103e961076936600461374e565b6122ac565b6103e961077c3660046137bf565b612773565b61035f601b81565b5f546103919062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f90819073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa15801561083a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061085e9190613807565b6007549091505f906108889067ffffffffffffffff6801000000000000000082048116911661384b565b67ffffffffffffffff169050805f036108a3575f9250505090565b6108ad8183613873565b9250505090565b600480546108c1906138ab565b80601f01602080910402602001604051908101604052809291908181526020018280546108ed906138ab565b80156109385780601f1061090f57610100808354040283529160200191610938565b820191905f5260205f20905b81548152906001019060200180831161091b57829003601f168201915b505050505081565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1633146109af576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff168367ffffffffffffffff167f9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f596684604051610a0291815260200190565b60405180910390a3505050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314610a65576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff82161115610aac576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166315064c966040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b15573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b3991906138fc565b610b9a5760075467ffffffffffffffff700100000000000000000000000000000000909104811690821610610b9a576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600780547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527fa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b906020015b60405180910390a150565b600380546108c1906138ab565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314610c81576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc090602001610c13565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314610d63576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f54610100900460ff1615808015610d8157505f54600160ff909116105b80610d9a5750303b158015610d9a57505f5460ff166001145b610e2b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610e87575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b606073ffffffffffffffffffffffffffffffffffffffff8516156110e4576040517fc00f14ab00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063c00f14ab906024015f60405180830381865afa158015610f2e573d5f803e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610f73919081019061391b565b6040517f318aee3d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87811660048301529192505f9182917f00000000000000000000000000000000000000000000000000000000000000009091169063318aee3d906024016040805180830381865afa158015611005573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611029919061398d565b915091508163ffffffff165f146110a0576009805463ffffffff841674010000000000000000000000000000000000000000027fffffffffffffffff00000000000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff8416171790556110e1565b600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff89161790555b50505b6009545f9061112b90889073ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900463ffffffff1685611d73565b90505f818051906020012090505f4290505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156111a6573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111ca9190613807565b90505f808483858f6111dd6001436139c5565b60408051602081019790975286019490945260608086019390935260c09190911b7fffffffffffffffff000000000000000000000000000000000000000000000000166080850152901b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016608883015240609c82015260bc01604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001815290829052805160209091012060058190557f9a908e73000000000000000000000000000000000000000000000000000000008252600160048301526024820181905291507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303815f875af1158015611323573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061134791906139de565b508c5f60026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508b60025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555088600390816113d79190613a3e565b5060046113e48982613a3e565b508c60085f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062069780600760106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507f060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f85838e60405161148493929190613b56565b60405180910390a150505050505080156114f4575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b60015473ffffffffffffffffffffffffffffffffffffffff16331461154e576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001545f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff1673ffffffffffffffffffffffffffffffffffffffff9092166201000081029290921790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e9060200160405180910390a1565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314611625576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60085473ffffffffffffffffffffffffffffffffffffffff16611674576040517fc89374d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb90602001610c13565b60085473ffffffffffffffffffffffffffffffffffffffff168015801590611725575073ffffffffffffffffffffffffffffffffffffffff81163314155b1561175c576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b4262093a807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166330c27dde6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117ca573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117ee91906139de565b6117f89190613b94565b67ffffffffffffffff16111561183a576040517f3d49ed4c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815f819003611875576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e88111156118b1576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075467ffffffffffffffff808216916118d991849168010000000000000000900416613bb5565b1115611911576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6007546005546801000000000000000090910467ffffffffffffffff16905f5b83811015611bac575f87878381811061194c5761194c613bc8565b905060200281019061195e9190613bf5565b61196790613c31565b90508361197381613cb7565b825180516020918201208185015160408087015160608801519151959a509295505f946119df948794929101938452602084019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166040830152604882015260680190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff89165f90815260069093529120549091508114611a67576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff86165f90815260066020526040812055611a8b6001886139c5565b8403611afa5742600760109054906101000a900467ffffffffffffffff168460400151611ab89190613b94565b67ffffffffffffffff161115611afa576040517fc44a082100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208381015160408086015160608088015183519586018b90529285018790528481019390935260c01b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808401523390911b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088830152609c82015260bc016040516020818303038152906040528051906020012094505050508080611ba490613cdd565b915050611931565b50611c227f000000000000000000000000000000000000000000000000000000000000000084611bda6107ae565b611be49190613d14565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169190612f74565b60058190556007805467ffffffffffffffff841668010000000000000000027fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff9091161790556040517f9a908e730000000000000000000000000000000000000000000000000000000081525f9073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690639a908e7390611cf3908790869060040167ffffffffffffffff929092168252602082015260400190565b6020604051808303815f875af1158015611d0f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611d3391906139de565b60405190915067ffffffffffffffff8216907f648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4905f90a250505050505050565b60605f85858573a40d5f56745a118d0906a34e69aec8c0db1cb8fa5f87604051602401611da596959493929190613d2b565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff811bff70000000000000000000000000000000000000000000000000000000017905283519091506060905f03611ef55760f9601f8351611e399190613d8d565b6040518060400160405280600881526020017f80808401c9c380940000000000000000000000000000000000000000000000008152507f00000000000000000000000000000000000000000000000000000000000000006040518060400160405280600281526020017f80b800000000000000000000000000000000000000000000000000000000000081525060e487604051602001611edf9796959493929190613da8565b6040516020818303038152906040529050611ff9565b815161ffff1015611f32576040517f248b8f8200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815160f9611f41602083613d8d565b6040518060400160405280600881526020017f80808401c9c380940000000000000000000000000000000000000000000000008152507f00000000000000000000000000000000000000000000000000000000000000006040518060400160405280600281526020017f80b90000000000000000000000000000000000000000000000000000000000008152508588604051602001611fe69796959493929190613e8a565b6040516020818303038152906040529150505b8051602080830191909120604080515f80825293810180835292909252601b908201526405ca1ab1e06060820152635ca1ab1e608082015260019060a0016020604051602081039080840390855afa158015612057573d5f803e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81166120cf576040517fcd16196600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040515f906121149084906405ca1ab1e090635ca1ab1e90601b907fff0000000000000000000000000000000000000000000000000000000000000090602001613f6c565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529450505050505b949350505050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff1633146121a7576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001610c13565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314612270576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600361227c8282613a3e565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b2081604051610c139190613392565b60085473ffffffffffffffffffffffffffffffffffffffff1680158015906122ea575073ffffffffffffffffffffffffffffffffffffffff81163314155b15612321576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166315064c966040518163ffffffff1660e01b8152600401602060405180830381865afa15801561238a573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906123ae91906138fc565b156123e5576040517f39258d1800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663604691696040518163ffffffff1660e01b8152600401602060405180830381865afa15801561244f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906124739190613807565b9050828111156124af576040517f2354600f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113888411156124eb576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61252d73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001633308461304d565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612597573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906125bb9190613807565b6007805491925067ffffffffffffffff909116905f6125d983613cb7565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550508585604051612610929190613fc7565b60405190819003902081426126266001436139c5565b60408051602081019590955284019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166060830152406068820152608801604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012060075467ffffffffffffffff165f908152600690935291205532330361271c576007546040805183815233602082015260608183018190525f90820152905167ffffffffffffffff909216917ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319181900360800190a261276b565b60075460405167ffffffffffffffff909116907ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319061276290849033908b908b90613fd6565b60405180910390a25b505050505050565b60025473ffffffffffffffffffffffffffffffffffffffff1633146127c4576040517f11e7be1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815f8190036127ff576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e881111561283b576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379e2cf976040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156128a0575f80fd5b505af11580156128b2573d5f803e3d5ffd5b505050505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635ca1e1656040518163ffffffff1660e01b8152600401602060405180830381865afa158015612920573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906129449190613807565b60075460055491925042916801000000000000000090910467ffffffffffffffff1690815f5b86811015612c65575f8a8a8381811061298557612985613bc8565b90506020028101906129979190613bf5565b6129a090613c31565b8051805160209091012060408201519192509067ffffffffffffffff1615612b8057856129cc81613cb7565b9650505f81836020015184604001518560600151604051602001612a2e9493929190938452602084019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166040830152604882015260680190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff8a165f90815260069093529120549091508114612ab6576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208381015160408086015160608088015183519586018c90529285018790528481019390935260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166080840152908d901b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088830152609c82015260bc0160405160208183030381529060405280519060200120955060065f8867ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f905550612c50565b8151516201d4c01015612bbf576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080516020810187905290810182905260608082018a905260c089901b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808301528b901b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660888201525f609c82015260bc016040516020818303038152906040528051906020012094505b50508080612c5d90613cdd565b91505061296a565b5060075467ffffffffffffffff9081169084161115612cb0576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60058290558567ffffffffffffffff84811690831614612d65575f612cd5838661384b565b9050612ceb67ffffffffffffffff8216836139c5565b9150612d247f00000000000000000000000000000000000000000000000000000000000000008267ffffffffffffffff16611bda6107ae565b50600780547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8716021790555b612e61337f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663477fa2706040518163ffffffff1660e01b8152600401602060405180830381865afa158015612df4573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612e189190613807565b612e229190613d14565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692919061304d565b6040517f9a908e7300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff88166004820152602481018490525f907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303815f875af1158015612efd573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612f2191906139de565b90508067ffffffffffffffff167f3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e76688604051612f5f91815260200190565b60405180910390a25050505050505050505050565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526130489084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526130b1565b505050565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526130ab9085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401612fc6565b50505050565b5f613112826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166131bc9092919063ffffffff16565b805190915015613048578080602001905181019061313091906138fc565b613048576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610e22565b606061214984845f85855f808673ffffffffffffffffffffffffffffffffffffffff1685876040516131ee9190614046565b5f6040518083038185875af1925050503d805f8114613228576040519150601f19603f3d011682016040523d82523d5f602084013e61322d565b606091505b509150915061323e87838387613249565b979650505050505050565b606083156132de5782515f036132d75773ffffffffffffffffffffffffffffffffffffffff85163b6132d7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610e22565b5081612149565b61214983838151156132f35781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e229190613392565b5f5b83811015613341578181015183820152602001613329565b50505f910152565b5f8151808452613360816020860160208601613327565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081525f6133a46020830184613349565b9392505050565b67ffffffffffffffff811681146133c0575f80fd5b50565b73ffffffffffffffffffffffffffffffffffffffff811681146133c0575f80fd5b5f805f606084860312156133f6575f80fd5b8335613401816133ab565b9250602084013591506040840135613418816133c3565b809150509250925092565b5f60208284031215613433575f80fd5b81356133a4816133ab565b5f6020828403121561344e575f80fd5b81356133a4816133c3565b63ffffffff811681146133c0575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156134de576134de61346a565b604052919050565b5f67ffffffffffffffff8211156134ff576134ff61346a565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b5f82601f83011261353a575f80fd5b813561354d613548826134e6565b613497565b818152846020838601011115613561575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f805f8060c08789031215613592575f80fd5b863561359d816133c3565b955060208701356135ad816133c3565b945060408701356135bd81613459565b935060608701356135cd816133c3565b9250608087013567ffffffffffffffff808211156135e9575f80fd5b6135f58a838b0161352b565b935060a089013591508082111561360a575f80fd5b5061361789828a0161352b565b9150509295509295509295565b5f8083601f840112613634575f80fd5b50813567ffffffffffffffff81111561364b575f80fd5b6020830191508360208260051b8501011115613665575f80fd5b9250929050565b5f806020838503121561367d575f80fd5b823567ffffffffffffffff811115613693575f80fd5b61369f85828601613624565b90969095509350505050565b5f805f80608085870312156136be575f80fd5b84356136c981613459565b935060208501356136d9816133c3565b925060408501356136e981613459565b9150606085013567ffffffffffffffff811115613704575f80fd5b6137108782880161352b565b91505092959194509250565b5f6020828403121561372c575f80fd5b813567ffffffffffffffff811115613742575f80fd5b6121498482850161352b565b5f805f60408486031215613760575f80fd5b833567ffffffffffffffff80821115613777575f80fd5b818601915086601f83011261378a575f80fd5b813581811115613798575f80fd5b8760208285010111156137a9575f80fd5b6020928301989097509590910135949350505050565b5f805f604084860312156137d1575f80fd5b833567ffffffffffffffff8111156137e7575f80fd5b6137f386828701613624565b9094509250506020840135613418816133c3565b5f60208284031215613817575f80fd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b67ffffffffffffffff82811682821603908082111561386c5761386c61381e565b5092915050565b5f826138a6577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b600181811c908216806138bf57607f821691505b6020821081036138f6577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b5f6020828403121561390c575f80fd5b815180151581146133a4575f80fd5b5f6020828403121561392b575f80fd5b815167ffffffffffffffff811115613941575f80fd5b8201601f81018413613951575f80fd5b805161395f613548826134e6565b818152856020838501011115613973575f80fd5b613984826020830160208601613327565b95945050505050565b5f806040838503121561399e575f80fd5b82516139a981613459565b60208401519092506139ba816133c3565b809150509250929050565b818103818111156139d8576139d861381e565b92915050565b5f602082840312156139ee575f80fd5b81516133a4816133ab565b601f821115613048575f81815260208120601f850160051c81016020861015613a1f5750805b601f850160051c820191505b8181101561276b57828155600101613a2b565b815167ffffffffffffffff811115613a5857613a5861346a565b613a6c81613a6684546138ab565b846139f9565b602080601f831160018114613abe575f8415613a885750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b17855561276b565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015613b0a57888601518255948401946001909101908401613aeb565b5085821015613b4657878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b606081525f613b686060830186613349565b905083602083015273ffffffffffffffffffffffffffffffffffffffff83166040830152949350505050565b67ffffffffffffffff81811683821601908082111561386c5761386c61381e565b808201808211156139d8576139d861381e565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f82357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81833603018112613c27575f80fd5b9190910192915050565b5f60808236031215613c41575f80fd5b6040516080810167ffffffffffffffff8282108183111715613c6557613c6561346a565b816040528435915080821115613c79575f80fd5b50613c863682860161352b565b825250602083013560208201526040830135613ca1816133ab565b6040820152606092830135928101929092525090565b5f67ffffffffffffffff808316818103613cd357613cd361381e565b6001019392505050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203613d0d57613d0d61381e565b5060010190565b80820281158282048414176139d8576139d861381e565b5f63ffffffff808916835273ffffffffffffffffffffffffffffffffffffffff8089166020850152818816604085015280871660608501528086166080850152505060c060a0830152613d8160c0830184613349565b98975050505050505050565b61ffff81811683821601908082111561386c5761386c61381e565b5f7fff00000000000000000000000000000000000000000000000000000000000000808a60f81b1683527fffff0000000000000000000000000000000000000000000000000000000000008960f01b1660018401528751613e10816003860160208c01613327565b80840190507fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008860601b1660038201528651613e53816017840160208b01613327565b808201915050818660f81b16601782015284519150613e79826018830160208801613327565b016018019998505050505050505050565b7fff000000000000000000000000000000000000000000000000000000000000008860f81b1681525f7fffff000000000000000000000000000000000000000000000000000000000000808960f01b1660018401528751613ef2816003860160208c01613327565b80840190507fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008860601b1660038201528651613f35816017840160208b01613327565b808201915050818660f01b16601782015284519150613f5b826019830160208801613327565b016019019998505050505050505050565b5f8651613f7d818460208b01613327565b9190910194855250602084019290925260f81b7fff000000000000000000000000000000000000000000000000000000000000009081166040840152166041820152604201919050565b818382375f9101908152919050565b84815273ffffffffffffffffffffffffffffffffffffffff8416602082015260606040820152816060820152818360808301375f818301608090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01601019392505050565b5f8251613c2781846020870161332756fea26469706673582212207601678cf856e87bac98cb0a8212b6f0358eca4ca427667fa9cee163b3f6709664736f6c63430008140033",
}

// PolygonrollupbaseetrogpreviousABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonrollupbaseetrogpreviousMetaData.ABI instead.
var PolygonrollupbaseetrogpreviousABI = PolygonrollupbaseetrogpreviousMetaData.ABI

// PolygonrollupbaseetrogpreviousBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygonrollupbaseetrogpreviousMetaData.Bin instead.
var PolygonrollupbaseetrogpreviousBin = PolygonrollupbaseetrogpreviousMetaData.Bin

// DeployPolygonrollupbaseetrogprevious deploys a new Ethereum contract, binding an instance of Polygonrollupbaseetrogprevious to it.
func DeployPolygonrollupbaseetrogprevious(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address, _rollupManager common.Address) (common.Address, *types.Transaction, *Polygonrollupbaseetrogprevious, error) {
	parsed, err := PolygonrollupbaseetrogpreviousMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygonrollupbaseetrogpreviousBin), backend, _globalExitRootManager, _pol, _bridgeAddress, _rollupManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonrollupbaseetrogprevious{PolygonrollupbaseetrogpreviousCaller: PolygonrollupbaseetrogpreviousCaller{contract: contract}, PolygonrollupbaseetrogpreviousTransactor: PolygonrollupbaseetrogpreviousTransactor{contract: contract}, PolygonrollupbaseetrogpreviousFilterer: PolygonrollupbaseetrogpreviousFilterer{contract: contract}}, nil
}

// Polygonrollupbaseetrogprevious is an auto generated Go binding around an Ethereum contract.
type Polygonrollupbaseetrogprevious struct {
	PolygonrollupbaseetrogpreviousCaller     // Read-only binding to the contract
	PolygonrollupbaseetrogpreviousTransactor // Write-only binding to the contract
	PolygonrollupbaseetrogpreviousFilterer   // Log filterer for contract events
}

// PolygonrollupbaseetrogpreviousCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonrollupbaseetrogpreviousCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonrollupbaseetrogpreviousTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonrollupbaseetrogpreviousTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonrollupbaseetrogpreviousFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonrollupbaseetrogpreviousFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonrollupbaseetrogpreviousSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonrollupbaseetrogpreviousSession struct {
	Contract     *Polygonrollupbaseetrogprevious // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// PolygonrollupbaseetrogpreviousCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonrollupbaseetrogpreviousCallerSession struct {
	Contract *PolygonrollupbaseetrogpreviousCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// PolygonrollupbaseetrogpreviousTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonrollupbaseetrogpreviousTransactorSession struct {
	Contract     *PolygonrollupbaseetrogpreviousTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// PolygonrollupbaseetrogpreviousRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonrollupbaseetrogpreviousRaw struct {
	Contract *Polygonrollupbaseetrogprevious // Generic contract binding to access the raw methods on
}

// PolygonrollupbaseetrogpreviousCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonrollupbaseetrogpreviousCallerRaw struct {
	Contract *PolygonrollupbaseetrogpreviousCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonrollupbaseetrogpreviousTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonrollupbaseetrogpreviousTransactorRaw struct {
	Contract *PolygonrollupbaseetrogpreviousTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonrollupbaseetrogprevious creates a new instance of Polygonrollupbaseetrogprevious, bound to a specific deployed contract.
func NewPolygonrollupbaseetrogprevious(address common.Address, backend bind.ContractBackend) (*Polygonrollupbaseetrogprevious, error) {
	contract, err := bindPolygonrollupbaseetrogprevious(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonrollupbaseetrogprevious{PolygonrollupbaseetrogpreviousCaller: PolygonrollupbaseetrogpreviousCaller{contract: contract}, PolygonrollupbaseetrogpreviousTransactor: PolygonrollupbaseetrogpreviousTransactor{contract: contract}, PolygonrollupbaseetrogpreviousFilterer: PolygonrollupbaseetrogpreviousFilterer{contract: contract}}, nil
}

// NewPolygonrollupbaseetrogpreviousCaller creates a new read-only instance of Polygonrollupbaseetrogprevious, bound to a specific deployed contract.
func NewPolygonrollupbaseetrogpreviousCaller(address common.Address, caller bind.ContractCaller) (*PolygonrollupbaseetrogpreviousCaller, error) {
	contract, err := bindPolygonrollupbaseetrogprevious(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpreviousCaller{contract: contract}, nil
}

// NewPolygonrollupbaseetrogpreviousTransactor creates a new write-only instance of Polygonrollupbaseetrogprevious, bound to a specific deployed contract.
func NewPolygonrollupbaseetrogpreviousTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonrollupbaseetrogpreviousTransactor, error) {
	contract, err := bindPolygonrollupbaseetrogprevious(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpreviousTransactor{contract: contract}, nil
}

// NewPolygonrollupbaseetrogpreviousFilterer creates a new log filterer instance of Polygonrollupbaseetrogprevious, bound to a specific deployed contract.
func NewPolygonrollupbaseetrogpreviousFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonrollupbaseetrogpreviousFilterer, error) {
	contract, err := bindPolygonrollupbaseetrogprevious(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpreviousFilterer{contract: contract}, nil
}

// bindPolygonrollupbaseetrogprevious binds a generic wrapper to an already deployed contract.
func bindPolygonrollupbaseetrogprevious(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonrollupbaseetrogpreviousMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonrollupbaseetrogprevious.Contract.PolygonrollupbaseetrogpreviousCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.Contract.PolygonrollupbaseetrogpreviousTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.Contract.PolygonrollupbaseetrogpreviousTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonrollupbaseetrogprevious.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.Contract.contract.Transact(opts, method, params...)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) GLOBALEXITROOTMANAGERL2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "GLOBAL_EXIT_ROOT_MANAGER_L2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _Polygonrollupbaseetrogprevious.Contract.GLOBALEXITROOTMANAGERL2(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _Polygonrollupbaseetrogprevious.Contract.GLOBALEXITROOTMANAGERL2(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) INITIALIZETXBRIDGELISTLENLEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_LIST_LEN_LEN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _Polygonrollupbaseetrogprevious.Contract.INITIALIZETXBRIDGELISTLENLEN(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _Polygonrollupbaseetrogprevious.Contract.INITIALIZETXBRIDGELISTLENLEN(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) INITIALIZETXBRIDGEPARAMS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _Polygonrollupbaseetrogprevious.Contract.INITIALIZETXBRIDGEPARAMS(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _Polygonrollupbaseetrogprevious.Contract.INITIALIZETXBRIDGEPARAMS(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _Polygonrollupbaseetrogprevious.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _Polygonrollupbaseetrogprevious.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _Polygonrollupbaseetrogprevious.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _Polygonrollupbaseetrogprevious.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) INITIALIZETXCONSTANTBYTES(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _Polygonrollupbaseetrogprevious.Contract.INITIALIZETXCONSTANTBYTES(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _Polygonrollupbaseetrogprevious.Contract.INITIALIZETXCONSTANTBYTES(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) INITIALIZETXCONSTANTBYTESEMPTYMETADATA(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _Polygonrollupbaseetrogprevious.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _Polygonrollupbaseetrogprevious.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) INITIALIZETXDATALENEMPTYMETADATA(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "INITIALIZE_TX_DATA_LEN_EMPTY_METADATA")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _Polygonrollupbaseetrogprevious.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _Polygonrollupbaseetrogprevious.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) INITIALIZETXEFFECTIVEPERCENTAGE(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "INITIALIZE_TX_EFFECTIVE_PERCENTAGE")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _Polygonrollupbaseetrogprevious.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _Polygonrollupbaseetrogprevious.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) SIGNATUREINITIALIZETXR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_R")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _Polygonrollupbaseetrogprevious.Contract.SIGNATUREINITIALIZETXR(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _Polygonrollupbaseetrogprevious.Contract.SIGNATUREINITIALIZETXR(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) SIGNATUREINITIALIZETXS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_S")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _Polygonrollupbaseetrogprevious.Contract.SIGNATUREINITIALIZETXS(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _Polygonrollupbaseetrogprevious.Contract.SIGNATUREINITIALIZETXS(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) SIGNATUREINITIALIZETXV(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_V")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _Polygonrollupbaseetrogprevious.Contract.SIGNATUREINITIALIZETXV(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _Polygonrollupbaseetrogprevious.Contract.SIGNATUREINITIALIZETXV(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) Admin() (common.Address, error) {
	return _Polygonrollupbaseetrogprevious.Contract.Admin(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) Admin() (common.Address, error) {
	return _Polygonrollupbaseetrogprevious.Contract.Admin(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) BridgeAddress() (common.Address, error) {
	return _Polygonrollupbaseetrogprevious.Contract.BridgeAddress(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) BridgeAddress() (common.Address, error) {
	return _Polygonrollupbaseetrogprevious.Contract.BridgeAddress(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) CalculatePolPerForceBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "calculatePolPerForceBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) CalculatePolPerForceBatch() (*big.Int, error) {
	return _Polygonrollupbaseetrogprevious.Contract.CalculatePolPerForceBatch(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) CalculatePolPerForceBatch() (*big.Int, error) {
	return _Polygonrollupbaseetrogprevious.Contract.CalculatePolPerForceBatch(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) ForceBatchAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "forceBatchAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) ForceBatchAddress() (common.Address, error) {
	return _Polygonrollupbaseetrogprevious.Contract.ForceBatchAddress(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) ForceBatchAddress() (common.Address, error) {
	return _Polygonrollupbaseetrogprevious.Contract.ForceBatchAddress(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonrollupbaseetrogprevious.Contract.ForceBatchTimeout(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonrollupbaseetrogprevious.Contract.ForceBatchTimeout(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonrollupbaseetrogprevious.Contract.ForcedBatches(&_Polygonrollupbaseetrogprevious.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonrollupbaseetrogprevious.Contract.ForcedBatches(&_Polygonrollupbaseetrogprevious.CallOpts, arg0)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) GasTokenAddress() (common.Address, error) {
	return _Polygonrollupbaseetrogprevious.Contract.GasTokenAddress(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) GasTokenAddress() (common.Address, error) {
	return _Polygonrollupbaseetrogprevious.Contract.GasTokenAddress(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) GasTokenNetwork() (uint32, error) {
	return _Polygonrollupbaseetrogprevious.Contract.GasTokenNetwork(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) GasTokenNetwork() (uint32, error) {
	return _Polygonrollupbaseetrogprevious.Contract.GasTokenNetwork(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) GenerateInitializeTransaction(opts *bind.CallOpts, networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "generateInitializeTransaction", networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) GenerateInitializeTransaction(networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _Polygonrollupbaseetrogprevious.Contract.GenerateInitializeTransaction(&_Polygonrollupbaseetrogprevious.CallOpts, networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) GenerateInitializeTransaction(networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _Polygonrollupbaseetrogprevious.Contract.GenerateInitializeTransaction(&_Polygonrollupbaseetrogprevious.CallOpts, networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonrollupbaseetrogprevious.Contract.GlobalExitRootManager(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonrollupbaseetrogprevious.Contract.GlobalExitRootManager(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) LastAccInputHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "lastAccInputHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) LastAccInputHash() ([32]byte, error) {
	return _Polygonrollupbaseetrogprevious.Contract.LastAccInputHash(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) LastAccInputHash() ([32]byte, error) {
	return _Polygonrollupbaseetrogprevious.Contract.LastAccInputHash(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) LastForceBatch() (uint64, error) {
	return _Polygonrollupbaseetrogprevious.Contract.LastForceBatch(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) LastForceBatch() (uint64, error) {
	return _Polygonrollupbaseetrogprevious.Contract.LastForceBatch(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonrollupbaseetrogprevious.Contract.LastForceBatchSequenced(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonrollupbaseetrogprevious.Contract.LastForceBatchSequenced(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) NetworkName() (string, error) {
	return _Polygonrollupbaseetrogprevious.Contract.NetworkName(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) NetworkName() (string, error) {
	return _Polygonrollupbaseetrogprevious.Contract.NetworkName(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) PendingAdmin() (common.Address, error) {
	return _Polygonrollupbaseetrogprevious.Contract.PendingAdmin(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) PendingAdmin() (common.Address, error) {
	return _Polygonrollupbaseetrogprevious.Contract.PendingAdmin(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) Pol() (common.Address, error) {
	return _Polygonrollupbaseetrogprevious.Contract.Pol(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) Pol() (common.Address, error) {
	return _Polygonrollupbaseetrogprevious.Contract.Pol(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) RollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "rollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) RollupManager() (common.Address, error) {
	return _Polygonrollupbaseetrogprevious.Contract.RollupManager(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) RollupManager() (common.Address, error) {
	return _Polygonrollupbaseetrogprevious.Contract.RollupManager(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) TrustedSequencer() (common.Address, error) {
	return _Polygonrollupbaseetrogprevious.Contract.TrustedSequencer(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) TrustedSequencer() (common.Address, error) {
	return _Polygonrollupbaseetrogprevious.Contract.TrustedSequencer(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCaller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogprevious.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) TrustedSequencerURL() (string, error) {
	return _Polygonrollupbaseetrogprevious.Contract.TrustedSequencerURL(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousCallerSession) TrustedSequencerURL() (string, error) {
	return _Polygonrollupbaseetrogprevious.Contract.TrustedSequencerURL(&_Polygonrollupbaseetrogprevious.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.Contract.AcceptAdminRole(&_Polygonrollupbaseetrogprevious.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.Contract.AcceptAdminRole(&_Polygonrollupbaseetrogprevious.TransactOpts)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousTransactor) ForceBatch(opts *bind.TransactOpts, transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.contract.Transact(opts, "forceBatch", transactions, polAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) ForceBatch(transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.Contract.ForceBatch(&_Polygonrollupbaseetrogprevious.TransactOpts, transactions, polAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousTransactorSession) ForceBatch(transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.Contract.ForceBatch(&_Polygonrollupbaseetrogprevious.TransactOpts, transactions, polAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.contract.Transact(opts, "initialize", _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.Contract.Initialize(&_Polygonrollupbaseetrogprevious.TransactOpts, _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousTransactorSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.Contract.Initialize(&_Polygonrollupbaseetrogprevious.TransactOpts, _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousTransactor) OnVerifyBatches(opts *bind.TransactOpts, lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.contract.Transact(opts, "onVerifyBatches", lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.Contract.OnVerifyBatches(&_Polygonrollupbaseetrogprevious.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousTransactorSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.Contract.OnVerifyBatches(&_Polygonrollupbaseetrogprevious.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xecef3f99.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, address l2Coinbase) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousTransactor) SequenceBatches(opts *bind.TransactOpts, batches []PolygonRollupBaseEtrogPreviousBatchData, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.contract.Transact(opts, "sequenceBatches", batches, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xecef3f99.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, address l2Coinbase) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) SequenceBatches(batches []PolygonRollupBaseEtrogPreviousBatchData, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.Contract.SequenceBatches(&_Polygonrollupbaseetrogprevious.TransactOpts, batches, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xecef3f99.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, address l2Coinbase) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousTransactorSession) SequenceBatches(batches []PolygonRollupBaseEtrogPreviousBatchData, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.Contract.SequenceBatches(&_Polygonrollupbaseetrogprevious.TransactOpts, batches, l2Coinbase)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousTransactor) SequenceForceBatches(opts *bind.TransactOpts, batches []PolygonRollupBaseEtrogPreviousBatchData) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.contract.Transact(opts, "sequenceForceBatches", batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) SequenceForceBatches(batches []PolygonRollupBaseEtrogPreviousBatchData) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.Contract.SequenceForceBatches(&_Polygonrollupbaseetrogprevious.TransactOpts, batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousTransactorSession) SequenceForceBatches(batches []PolygonRollupBaseEtrogPreviousBatchData) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.Contract.SequenceForceBatches(&_Polygonrollupbaseetrogprevious.TransactOpts, batches)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousTransactor) SetForceBatchAddress(opts *bind.TransactOpts, newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.contract.Transact(opts, "setForceBatchAddress", newForceBatchAddress)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) SetForceBatchAddress(newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.Contract.SetForceBatchAddress(&_Polygonrollupbaseetrogprevious.TransactOpts, newForceBatchAddress)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousTransactorSession) SetForceBatchAddress(newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.Contract.SetForceBatchAddress(&_Polygonrollupbaseetrogprevious.TransactOpts, newForceBatchAddress)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousTransactor) SetForceBatchTimeout(opts *bind.TransactOpts, newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.contract.Transact(opts, "setForceBatchTimeout", newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.Contract.SetForceBatchTimeout(&_Polygonrollupbaseetrogprevious.TransactOpts, newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousTransactorSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.Contract.SetForceBatchTimeout(&_Polygonrollupbaseetrogprevious.TransactOpts, newforceBatchTimeout)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousTransactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.Contract.SetTrustedSequencer(&_Polygonrollupbaseetrogprevious.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousTransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.Contract.SetTrustedSequencer(&_Polygonrollupbaseetrogprevious.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousTransactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.Contract.SetTrustedSequencerURL(&_Polygonrollupbaseetrogprevious.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousTransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.Contract.SetTrustedSequencerURL(&_Polygonrollupbaseetrogprevious.TransactOpts, newTrustedSequencerURL)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.Contract.TransferAdminRole(&_Polygonrollupbaseetrogprevious.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogprevious.Contract.TransferAdminRole(&_Polygonrollupbaseetrogprevious.TransactOpts, newPendingAdmin)
}

// PolygonrollupbaseetrogpreviousAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Polygonrollupbaseetrogprevious contract.
type PolygonrollupbaseetrogpreviousAcceptAdminRoleIterator struct {
	Event *PolygonrollupbaseetrogpreviousAcceptAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonrollupbaseetrogpreviousAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupbaseetrogpreviousAcceptAdminRole)
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
		it.Event = new(PolygonrollupbaseetrogpreviousAcceptAdminRole)
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
func (it *PolygonrollupbaseetrogpreviousAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupbaseetrogpreviousAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupbaseetrogpreviousAcceptAdminRole represents a AcceptAdminRole event raised by the Polygonrollupbaseetrogprevious contract.
type PolygonrollupbaseetrogpreviousAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*PolygonrollupbaseetrogpreviousAcceptAdminRoleIterator, error) {

	logs, sub, err := _Polygonrollupbaseetrogprevious.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpreviousAcceptAdminRoleIterator{contract: _Polygonrollupbaseetrogprevious.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonrollupbaseetrogpreviousAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupbaseetrogprevious.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupbaseetrogpreviousAcceptAdminRole)
				if err := _Polygonrollupbaseetrogprevious.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
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
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) ParseAcceptAdminRole(log types.Log) (*PolygonrollupbaseetrogpreviousAcceptAdminRole, error) {
	event := new(PolygonrollupbaseetrogpreviousAcceptAdminRole)
	if err := _Polygonrollupbaseetrogprevious.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupbaseetrogpreviousForceBatchIterator is returned from FilterForceBatch and is used to iterate over the raw logs and unpacked data for ForceBatch events raised by the Polygonrollupbaseetrogprevious contract.
type PolygonrollupbaseetrogpreviousForceBatchIterator struct {
	Event *PolygonrollupbaseetrogpreviousForceBatch // Event containing the contract specifics and raw log

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
func (it *PolygonrollupbaseetrogpreviousForceBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupbaseetrogpreviousForceBatch)
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
		it.Event = new(PolygonrollupbaseetrogpreviousForceBatch)
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
func (it *PolygonrollupbaseetrogpreviousForceBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupbaseetrogpreviousForceBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupbaseetrogpreviousForceBatch represents a ForceBatch event raised by the Polygonrollupbaseetrogprevious contract.
type PolygonrollupbaseetrogpreviousForceBatch struct {
	ForceBatchNum      uint64
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Transactions       []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterForceBatch is a free log retrieval operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) FilterForceBatch(opts *bind.FilterOpts, forceBatchNum []uint64) (*PolygonrollupbaseetrogpreviousForceBatchIterator, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Polygonrollupbaseetrogprevious.contract.FilterLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpreviousForceBatchIterator{contract: _Polygonrollupbaseetrogprevious.contract, event: "ForceBatch", logs: logs, sub: sub}, nil
}

// WatchForceBatch is a free log subscription operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) WatchForceBatch(opts *bind.WatchOpts, sink chan<- *PolygonrollupbaseetrogpreviousForceBatch, forceBatchNum []uint64) (event.Subscription, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Polygonrollupbaseetrogprevious.contract.WatchLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupbaseetrogpreviousForceBatch)
				if err := _Polygonrollupbaseetrogprevious.contract.UnpackLog(event, "ForceBatch", log); err != nil {
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

// ParseForceBatch is a log parse operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) ParseForceBatch(log types.Log) (*PolygonrollupbaseetrogpreviousForceBatch, error) {
	event := new(PolygonrollupbaseetrogpreviousForceBatch)
	if err := _Polygonrollupbaseetrogprevious.contract.UnpackLog(event, "ForceBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupbaseetrogpreviousInitialSequenceBatchesIterator is returned from FilterInitialSequenceBatches and is used to iterate over the raw logs and unpacked data for InitialSequenceBatches events raised by the Polygonrollupbaseetrogprevious contract.
type PolygonrollupbaseetrogpreviousInitialSequenceBatchesIterator struct {
	Event *PolygonrollupbaseetrogpreviousInitialSequenceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonrollupbaseetrogpreviousInitialSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupbaseetrogpreviousInitialSequenceBatches)
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
		it.Event = new(PolygonrollupbaseetrogpreviousInitialSequenceBatches)
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
func (it *PolygonrollupbaseetrogpreviousInitialSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupbaseetrogpreviousInitialSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupbaseetrogpreviousInitialSequenceBatches represents a InitialSequenceBatches event raised by the Polygonrollupbaseetrogprevious contract.
type PolygonrollupbaseetrogpreviousInitialSequenceBatches struct {
	Transactions       []byte
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterInitialSequenceBatches is a free log retrieval operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) FilterInitialSequenceBatches(opts *bind.FilterOpts) (*PolygonrollupbaseetrogpreviousInitialSequenceBatchesIterator, error) {

	logs, sub, err := _Polygonrollupbaseetrogprevious.contract.FilterLogs(opts, "InitialSequenceBatches")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpreviousInitialSequenceBatchesIterator{contract: _Polygonrollupbaseetrogprevious.contract, event: "InitialSequenceBatches", logs: logs, sub: sub}, nil
}

// WatchInitialSequenceBatches is a free log subscription operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) WatchInitialSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonrollupbaseetrogpreviousInitialSequenceBatches) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupbaseetrogprevious.contract.WatchLogs(opts, "InitialSequenceBatches")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupbaseetrogpreviousInitialSequenceBatches)
				if err := _Polygonrollupbaseetrogprevious.contract.UnpackLog(event, "InitialSequenceBatches", log); err != nil {
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

// ParseInitialSequenceBatches is a log parse operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) ParseInitialSequenceBatches(log types.Log) (*PolygonrollupbaseetrogpreviousInitialSequenceBatches, error) {
	event := new(PolygonrollupbaseetrogpreviousInitialSequenceBatches)
	if err := _Polygonrollupbaseetrogprevious.contract.UnpackLog(event, "InitialSequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupbaseetrogpreviousInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonrollupbaseetrogprevious contract.
type PolygonrollupbaseetrogpreviousInitializedIterator struct {
	Event *PolygonrollupbaseetrogpreviousInitialized // Event containing the contract specifics and raw log

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
func (it *PolygonrollupbaseetrogpreviousInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupbaseetrogpreviousInitialized)
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
		it.Event = new(PolygonrollupbaseetrogpreviousInitialized)
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
func (it *PolygonrollupbaseetrogpreviousInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupbaseetrogpreviousInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupbaseetrogpreviousInitialized represents a Initialized event raised by the Polygonrollupbaseetrogprevious contract.
type PolygonrollupbaseetrogpreviousInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) FilterInitialized(opts *bind.FilterOpts) (*PolygonrollupbaseetrogpreviousInitializedIterator, error) {

	logs, sub, err := _Polygonrollupbaseetrogprevious.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpreviousInitializedIterator{contract: _Polygonrollupbaseetrogprevious.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PolygonrollupbaseetrogpreviousInitialized) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupbaseetrogprevious.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupbaseetrogpreviousInitialized)
				if err := _Polygonrollupbaseetrogprevious.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) ParseInitialized(log types.Log) (*PolygonrollupbaseetrogpreviousInitialized, error) {
	event := new(PolygonrollupbaseetrogpreviousInitialized)
	if err := _Polygonrollupbaseetrogprevious.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupbaseetrogpreviousSequenceBatchesIterator is returned from FilterSequenceBatches and is used to iterate over the raw logs and unpacked data for SequenceBatches events raised by the Polygonrollupbaseetrogprevious contract.
type PolygonrollupbaseetrogpreviousSequenceBatchesIterator struct {
	Event *PolygonrollupbaseetrogpreviousSequenceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonrollupbaseetrogpreviousSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupbaseetrogpreviousSequenceBatches)
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
		it.Event = new(PolygonrollupbaseetrogpreviousSequenceBatches)
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
func (it *PolygonrollupbaseetrogpreviousSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupbaseetrogpreviousSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupbaseetrogpreviousSequenceBatches represents a SequenceBatches event raised by the Polygonrollupbaseetrogprevious contract.
type PolygonrollupbaseetrogpreviousSequenceBatches struct {
	NumBatch   uint64
	L1InfoRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSequenceBatches is a free log retrieval operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) FilterSequenceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonrollupbaseetrogpreviousSequenceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonrollupbaseetrogprevious.contract.FilterLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpreviousSequenceBatchesIterator{contract: _Polygonrollupbaseetrogprevious.contract, event: "SequenceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceBatches is a free log subscription operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) WatchSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonrollupbaseetrogpreviousSequenceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonrollupbaseetrogprevious.contract.WatchLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupbaseetrogpreviousSequenceBatches)
				if err := _Polygonrollupbaseetrogprevious.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
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

// ParseSequenceBatches is a log parse operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) ParseSequenceBatches(log types.Log) (*PolygonrollupbaseetrogpreviousSequenceBatches, error) {
	event := new(PolygonrollupbaseetrogpreviousSequenceBatches)
	if err := _Polygonrollupbaseetrogprevious.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupbaseetrogpreviousSequenceForceBatchesIterator is returned from FilterSequenceForceBatches and is used to iterate over the raw logs and unpacked data for SequenceForceBatches events raised by the Polygonrollupbaseetrogprevious contract.
type PolygonrollupbaseetrogpreviousSequenceForceBatchesIterator struct {
	Event *PolygonrollupbaseetrogpreviousSequenceForceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonrollupbaseetrogpreviousSequenceForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupbaseetrogpreviousSequenceForceBatches)
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
		it.Event = new(PolygonrollupbaseetrogpreviousSequenceForceBatches)
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
func (it *PolygonrollupbaseetrogpreviousSequenceForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupbaseetrogpreviousSequenceForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupbaseetrogpreviousSequenceForceBatches represents a SequenceForceBatches event raised by the Polygonrollupbaseetrogprevious contract.
type PolygonrollupbaseetrogpreviousSequenceForceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceForceBatches is a free log retrieval operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) FilterSequenceForceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonrollupbaseetrogpreviousSequenceForceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonrollupbaseetrogprevious.contract.FilterLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpreviousSequenceForceBatchesIterator{contract: _Polygonrollupbaseetrogprevious.contract, event: "SequenceForceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceForceBatches is a free log subscription operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) WatchSequenceForceBatches(opts *bind.WatchOpts, sink chan<- *PolygonrollupbaseetrogpreviousSequenceForceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonrollupbaseetrogprevious.contract.WatchLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupbaseetrogpreviousSequenceForceBatches)
				if err := _Polygonrollupbaseetrogprevious.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
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

// ParseSequenceForceBatches is a log parse operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) ParseSequenceForceBatches(log types.Log) (*PolygonrollupbaseetrogpreviousSequenceForceBatches, error) {
	event := new(PolygonrollupbaseetrogpreviousSequenceForceBatches)
	if err := _Polygonrollupbaseetrogprevious.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupbaseetrogpreviousSetForceBatchAddressIterator is returned from FilterSetForceBatchAddress and is used to iterate over the raw logs and unpacked data for SetForceBatchAddress events raised by the Polygonrollupbaseetrogprevious contract.
type PolygonrollupbaseetrogpreviousSetForceBatchAddressIterator struct {
	Event *PolygonrollupbaseetrogpreviousSetForceBatchAddress // Event containing the contract specifics and raw log

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
func (it *PolygonrollupbaseetrogpreviousSetForceBatchAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupbaseetrogpreviousSetForceBatchAddress)
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
		it.Event = new(PolygonrollupbaseetrogpreviousSetForceBatchAddress)
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
func (it *PolygonrollupbaseetrogpreviousSetForceBatchAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupbaseetrogpreviousSetForceBatchAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupbaseetrogpreviousSetForceBatchAddress represents a SetForceBatchAddress event raised by the Polygonrollupbaseetrogprevious contract.
type PolygonrollupbaseetrogpreviousSetForceBatchAddress struct {
	NewForceBatchAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchAddress is a free log retrieval operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) FilterSetForceBatchAddress(opts *bind.FilterOpts) (*PolygonrollupbaseetrogpreviousSetForceBatchAddressIterator, error) {

	logs, sub, err := _Polygonrollupbaseetrogprevious.contract.FilterLogs(opts, "SetForceBatchAddress")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpreviousSetForceBatchAddressIterator{contract: _Polygonrollupbaseetrogprevious.contract, event: "SetForceBatchAddress", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchAddress is a free log subscription operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) WatchSetForceBatchAddress(opts *bind.WatchOpts, sink chan<- *PolygonrollupbaseetrogpreviousSetForceBatchAddress) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupbaseetrogprevious.contract.WatchLogs(opts, "SetForceBatchAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupbaseetrogpreviousSetForceBatchAddress)
				if err := _Polygonrollupbaseetrogprevious.contract.UnpackLog(event, "SetForceBatchAddress", log); err != nil {
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

// ParseSetForceBatchAddress is a log parse operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) ParseSetForceBatchAddress(log types.Log) (*PolygonrollupbaseetrogpreviousSetForceBatchAddress, error) {
	event := new(PolygonrollupbaseetrogpreviousSetForceBatchAddress)
	if err := _Polygonrollupbaseetrogprevious.contract.UnpackLog(event, "SetForceBatchAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupbaseetrogpreviousSetForceBatchTimeoutIterator is returned from FilterSetForceBatchTimeout and is used to iterate over the raw logs and unpacked data for SetForceBatchTimeout events raised by the Polygonrollupbaseetrogprevious contract.
type PolygonrollupbaseetrogpreviousSetForceBatchTimeoutIterator struct {
	Event *PolygonrollupbaseetrogpreviousSetForceBatchTimeout // Event containing the contract specifics and raw log

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
func (it *PolygonrollupbaseetrogpreviousSetForceBatchTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupbaseetrogpreviousSetForceBatchTimeout)
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
		it.Event = new(PolygonrollupbaseetrogpreviousSetForceBatchTimeout)
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
func (it *PolygonrollupbaseetrogpreviousSetForceBatchTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupbaseetrogpreviousSetForceBatchTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupbaseetrogpreviousSetForceBatchTimeout represents a SetForceBatchTimeout event raised by the Polygonrollupbaseetrogprevious contract.
type PolygonrollupbaseetrogpreviousSetForceBatchTimeout struct {
	NewforceBatchTimeout uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchTimeout is a free log retrieval operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) FilterSetForceBatchTimeout(opts *bind.FilterOpts) (*PolygonrollupbaseetrogpreviousSetForceBatchTimeoutIterator, error) {

	logs, sub, err := _Polygonrollupbaseetrogprevious.contract.FilterLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpreviousSetForceBatchTimeoutIterator{contract: _Polygonrollupbaseetrogprevious.contract, event: "SetForceBatchTimeout", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchTimeout is a free log subscription operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) WatchSetForceBatchTimeout(opts *bind.WatchOpts, sink chan<- *PolygonrollupbaseetrogpreviousSetForceBatchTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupbaseetrogprevious.contract.WatchLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupbaseetrogpreviousSetForceBatchTimeout)
				if err := _Polygonrollupbaseetrogprevious.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
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

// ParseSetForceBatchTimeout is a log parse operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) ParseSetForceBatchTimeout(log types.Log) (*PolygonrollupbaseetrogpreviousSetForceBatchTimeout, error) {
	event := new(PolygonrollupbaseetrogpreviousSetForceBatchTimeout)
	if err := _Polygonrollupbaseetrogprevious.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupbaseetrogpreviousSetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the Polygonrollupbaseetrogprevious contract.
type PolygonrollupbaseetrogpreviousSetTrustedSequencerIterator struct {
	Event *PolygonrollupbaseetrogpreviousSetTrustedSequencer // Event containing the contract specifics and raw log

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
func (it *PolygonrollupbaseetrogpreviousSetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupbaseetrogpreviousSetTrustedSequencer)
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
		it.Event = new(PolygonrollupbaseetrogpreviousSetTrustedSequencer)
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
func (it *PolygonrollupbaseetrogpreviousSetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupbaseetrogpreviousSetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupbaseetrogpreviousSetTrustedSequencer represents a SetTrustedSequencer event raised by the Polygonrollupbaseetrogprevious contract.
type PolygonrollupbaseetrogpreviousSetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*PolygonrollupbaseetrogpreviousSetTrustedSequencerIterator, error) {

	logs, sub, err := _Polygonrollupbaseetrogprevious.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpreviousSetTrustedSequencerIterator{contract: _Polygonrollupbaseetrogprevious.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *PolygonrollupbaseetrogpreviousSetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupbaseetrogprevious.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupbaseetrogpreviousSetTrustedSequencer)
				if err := _Polygonrollupbaseetrogprevious.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
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
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) ParseSetTrustedSequencer(log types.Log) (*PolygonrollupbaseetrogpreviousSetTrustedSequencer, error) {
	event := new(PolygonrollupbaseetrogpreviousSetTrustedSequencer)
	if err := _Polygonrollupbaseetrogprevious.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupbaseetrogpreviousSetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the Polygonrollupbaseetrogprevious contract.
type PolygonrollupbaseetrogpreviousSetTrustedSequencerURLIterator struct {
	Event *PolygonrollupbaseetrogpreviousSetTrustedSequencerURL // Event containing the contract specifics and raw log

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
func (it *PolygonrollupbaseetrogpreviousSetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupbaseetrogpreviousSetTrustedSequencerURL)
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
		it.Event = new(PolygonrollupbaseetrogpreviousSetTrustedSequencerURL)
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
func (it *PolygonrollupbaseetrogpreviousSetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupbaseetrogpreviousSetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupbaseetrogpreviousSetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the Polygonrollupbaseetrogprevious contract.
type PolygonrollupbaseetrogpreviousSetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*PolygonrollupbaseetrogpreviousSetTrustedSequencerURLIterator, error) {

	logs, sub, err := _Polygonrollupbaseetrogprevious.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpreviousSetTrustedSequencerURLIterator{contract: _Polygonrollupbaseetrogprevious.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *PolygonrollupbaseetrogpreviousSetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupbaseetrogprevious.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupbaseetrogpreviousSetTrustedSequencerURL)
				if err := _Polygonrollupbaseetrogprevious.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
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
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) ParseSetTrustedSequencerURL(log types.Log) (*PolygonrollupbaseetrogpreviousSetTrustedSequencerURL, error) {
	event := new(PolygonrollupbaseetrogpreviousSetTrustedSequencerURL)
	if err := _Polygonrollupbaseetrogprevious.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupbaseetrogpreviousTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Polygonrollupbaseetrogprevious contract.
type PolygonrollupbaseetrogpreviousTransferAdminRoleIterator struct {
	Event *PolygonrollupbaseetrogpreviousTransferAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonrollupbaseetrogpreviousTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupbaseetrogpreviousTransferAdminRole)
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
		it.Event = new(PolygonrollupbaseetrogpreviousTransferAdminRole)
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
func (it *PolygonrollupbaseetrogpreviousTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupbaseetrogpreviousTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupbaseetrogpreviousTransferAdminRole represents a TransferAdminRole event raised by the Polygonrollupbaseetrogprevious contract.
type PolygonrollupbaseetrogpreviousTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*PolygonrollupbaseetrogpreviousTransferAdminRoleIterator, error) {

	logs, sub, err := _Polygonrollupbaseetrogprevious.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpreviousTransferAdminRoleIterator{contract: _Polygonrollupbaseetrogprevious.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonrollupbaseetrogpreviousTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupbaseetrogprevious.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupbaseetrogpreviousTransferAdminRole)
				if err := _Polygonrollupbaseetrogprevious.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
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
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) ParseTransferAdminRole(log types.Log) (*PolygonrollupbaseetrogpreviousTransferAdminRole, error) {
	event := new(PolygonrollupbaseetrogpreviousTransferAdminRole)
	if err := _Polygonrollupbaseetrogprevious.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupbaseetrogpreviousVerifyBatchesIterator is returned from FilterVerifyBatches and is used to iterate over the raw logs and unpacked data for VerifyBatches events raised by the Polygonrollupbaseetrogprevious contract.
type PolygonrollupbaseetrogpreviousVerifyBatchesIterator struct {
	Event *PolygonrollupbaseetrogpreviousVerifyBatches // Event containing the contract specifics and raw log

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
func (it *PolygonrollupbaseetrogpreviousVerifyBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupbaseetrogpreviousVerifyBatches)
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
		it.Event = new(PolygonrollupbaseetrogpreviousVerifyBatches)
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
func (it *PolygonrollupbaseetrogpreviousVerifyBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupbaseetrogpreviousVerifyBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupbaseetrogpreviousVerifyBatches represents a VerifyBatches event raised by the Polygonrollupbaseetrogprevious contract.
type PolygonrollupbaseetrogpreviousVerifyBatches struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatches is a free log retrieval operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) FilterVerifyBatches(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*PolygonrollupbaseetrogpreviousVerifyBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonrollupbaseetrogprevious.contract.FilterLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpreviousVerifyBatchesIterator{contract: _Polygonrollupbaseetrogprevious.contract, event: "VerifyBatches", logs: logs, sub: sub}, nil
}

// WatchVerifyBatches is a free log subscription operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) WatchVerifyBatches(opts *bind.WatchOpts, sink chan<- *PolygonrollupbaseetrogpreviousVerifyBatches, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonrollupbaseetrogprevious.contract.WatchLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupbaseetrogpreviousVerifyBatches)
				if err := _Polygonrollupbaseetrogprevious.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
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

// ParseVerifyBatches is a log parse operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonrollupbaseetrogprevious *PolygonrollupbaseetrogpreviousFilterer) ParseVerifyBatches(log types.Log) (*PolygonrollupbaseetrogpreviousVerifyBatches, error) {
	event := new(PolygonrollupbaseetrogpreviousVerifyBatches)
	if err := _Polygonrollupbaseetrogprevious.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

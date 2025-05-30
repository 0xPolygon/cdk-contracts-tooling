// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonzkevmetrog

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

// PolygonRollupBaseEtrogBatchData is an auto generated low-level Go binding around an user-defined struct.
type PolygonRollupBaseEtrogBatchData struct {
	Transactions         []byte
	ForcedGlobalExitRoot [32]byte
	ForcedTimestamp      uint64
	ForcedBlockHashL1    [32]byte
}

// PolygonzkevmetrogMetaData contains all meta data concerning the Polygonzkevmetrog contract.
var PolygonzkevmetrogMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"_rollupManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalAccInputHashDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"forceBatchNum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"}],\"name\":\"ForceBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"}],\"name\":\"InitialSequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"targetBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accInputHashToRollback\",\"type\":\"bytes32\"}],\"name\":\"RollbackBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"l1InfoRoot\",\"type\":\"bytes32\"}],\"name\":\"SequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newForceBatchAddress\",\"type\":\"address\"}],\"name\":\"SetForceBatchAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"SetForceBatchTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatches\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GLOBAL_EXIT_ROOT_MANAGER_L2\",\"outputs\":[{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_LIST_LEN_LEN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_DATA_LEN_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_EFFECTIVE_PERCENTAGE\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_R\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_S\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_V\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TIMESTAMP_RANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculatePolPerForceBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"polAmount\",\"type\":\"uint256\"}],\"name\":\"forceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"}],\"name\":\"generateInitializeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"onVerifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"targetBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"accInputHashToRollback\",\"type\":\"bytes32\"}],\"name\":\"rollbackBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrog.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"uint32\",\"name\":\"l1InfoTreeLeafCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"maxSequenceTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"expectedFinalAccInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"}],\"name\":\"sequenceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrog.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"}],\"name\":\"sequenceForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newForceBatchAddress\",\"type\":\"address\"}],\"name\":\"setForceBatchAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"setForceBatchTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x610100604052348015610010575f5ffd5b5060405161451038038061451083398101604081905261002f91610143565b6001600160a01b0380851660a05280841660805280831660c052811660e052838383838383838361005e61006f565b50505050505050505050505061019f565b5f54610100900460ff16156100da5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116101561012a575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114610140575f5ffd5b50565b5f5f5f5f60808587031215610156575f5ffd5b84516101618161012c565b60208601519094506101728161012c565b60408601519093506101838161012c565b60608601519092506101948161012c565b939692955090935050565b60805160a05160c05160e0516142886102885f395f81816104fc0152818161097301528181610adf01528181610c5e01528181610dde01528181611160015281816115fb01528181611a4401528181611b39015281816126c801528181612741015281816127630152818161287801528181612a8e0152612b5301525f818161066801528181611d0301528181611e0b015281816121c401528181612f41015261301601525f818161073701528181610fd7015281816122720152612c9b01525f81816107690152818161082601528181611a8d0152818161280f0152612c7001526142885ff3fe608060405234801561000f575f5ffd5b50600436106102e2575f3560e01c80637125702211610187578063c754c7ed116100dd578063d7bc90ff11610093578063eaeb077b1161006e578063eaeb077b1461079f578063f35dda47146107b2578063f851a440146107ba575f5ffd5b8063d7bc90ff14610759578063e46761c414610764578063e7a7ed021461078b575f5ffd5b8063c89e42df116100c3578063c89e42df146106ff578063cfa8ed4714610712578063d02103ca14610732575f5ffd5b8063c754c7ed146106cf578063c7fffd4b146106f7575f5ffd5b80639f26f8401161013d578063ada8f91911610118578063ada8f9191461069d578063b0afe154146106b0578063b910e0f9146106bc575f5ffd5b80639f26f84014610650578063a3c573eb14610663578063a652f26c1461068a575f5ffd5b80638c3d73011161016d5780638c3d73011461061a57806391cafe32146106225780639e00187714610635575f5ffd5b806371257022146105cb5780637a5460c5146105de575f5ffd5b806342308fab1161023c578063542028d5116101f25780636b8616ce116101cd5780636b8616ce146105905780636e05d2cd146105af5780636ff512cc146105b8575f5ffd5b8063542028d51461056d578063669adece14610575578063676870d214610588575f5ffd5b806349b7b8021161022257806349b7b802146104f75780634e4877061461051e57806352bdeb6d14610531575f5ffd5b806342308fab146104b657806345605267146104be575f5ffd5b8063267822471161029c5780633c351e10116102775780633c351e10146104015780633cbc795b1461042157806340b5de6c1461045e575f5ffd5b806326782247146103875780632c111c06146103cc57806332c2d153146103ec575f5ffd5b806305835f37116102cc57806305835f371461031c578063107bf28c1461036557806311e892d41461036d575f5ffd5b8062d0295d146102e65780630350896314610301575b5f5ffd5b6102ee6107df565b6040519081526020015b60405180910390f35b610309602081565b60405161ffff90911681526020016102f8565b6103586040518060400160405280600881526020017f80808401c9c3809400000000000000000000000000000000000000000000000081525081565b6040516102f8919061353f565b6103586108e5565b61037560f981565b60405160ff90911681526020016102f8565b6001546103a79073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102f8565b6008546103a79073ffffffffffffffffffffffffffffffffffffffff1681565b6103ff6103fa366004613591565b610971565b005b6009546103a79073ffffffffffffffffffffffffffffffffffffffff1681565b6009546104499074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff90911681526020016102f8565b6104857fff0000000000000000000000000000000000000000000000000000000000000081565b6040517fff0000000000000000000000000000000000000000000000000000000000000090911681526020016102f8565b6102ee602481565b6007546104de9068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020016102f8565b6103a77f000000000000000000000000000000000000000000000000000000000000000081565b6103ff61052c3660046135d0565b610a40565b6103586040518060400160405280600281526020017f80b800000000000000000000000000000000000000000000000000000000000081525081565b610358610c4f565b6103ff6105833660046135eb565b610c5c565b610309601f81565b6102ee61059e3660046135d0565b60066020525f908152604090205481565b6102ee60055481565b6103ff6105c6366004613615565b610d13565b6103ff6105d9366004613756565b610ddc565b6103586040518060400160405280600281526020017f80b900000000000000000000000000000000000000000000000000000000000081525081565b6103ff611395565b6103ff610630366004613615565b611467565b6103a773a40d5f56745a118d0906a34e69aec8c0db1cb8fa81565b6103ff61065e36600461384a565b61157f565b6103a77f000000000000000000000000000000000000000000000000000000000000000081565b610358610698366004613889565b611c05565b6103ff6106ab366004613615565b611fe3565b6102ee6405ca1ab1e081565b6103ff6106ca3660046138fa565b6120ac565b6007546104de90700100000000000000000000000000000000900467ffffffffffffffff1681565b61037560e481565b6103ff61070d36600461397a565b612985565b6002546103a79073ffffffffffffffffffffffffffffffffffffffff1681565b6103a77f000000000000000000000000000000000000000000000000000000000000000081565b6102ee635ca1ab1e81565b6103a77f000000000000000000000000000000000000000000000000000000000000000081565b6007546104de9067ffffffffffffffff1681565b6103ff6107ad3660046139ac565b612a17565b610375601b81565b5f546103a79062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f90819073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa15801561086b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061088f9190613a1e565b6007549091505f906108b99067ffffffffffffffff68010000000000000000820481169116613a62565b67ffffffffffffffff169050805f036108d4575f9250505090565b6108de8183613a88565b9250505090565b600480546108f290613ac0565b80601f016020809104026020016040519081016040528092919081815260200182805461091e90613ac0565b80156109695780601f1061094057610100808354040283529160200191610969565b820191905f5260205f20905b81548152906001019060200180831161094c57829003601f168201915b505050505081565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1633146109e0576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff168367ffffffffffffffff167f9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f596684604051610a3391815260200190565b60405180910390a3505050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314610a96576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff82161115610add576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166315064c966040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b46573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b6a9190613b11565b610bcb5760075467ffffffffffffffff700100000000000000000000000000000000909104811690821610610bcb576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600780547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527fa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b906020015b60405180910390a150565b600380546108f290613ac0565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314610ccb576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600581905560405181815267ffffffffffffffff8316907f1125aaf62d132d8e2d02005114f8fc360ff204c3105e4f1a700a1340dc55d5b19060200160405180910390a25050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314610d69576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc090602001610c44565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314610e4b576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f54610100900460ff1615808015610e6957505f54600160ff909116105b80610e825750303b158015610e8257505f5460ff166001145b610f13576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610f6f575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b5f610f7985612ede565b6009549091505f90610fc390889073ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900463ffffffff1685611c05565b90505f818051906020012090505f4290505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561103e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110629190613a1e565b90505f808483858f611075600143613b30565b60408051602081019790975286019490945260608086019390935260c09190911b7fffffffffffffffff000000000000000000000000000000000000000000000000166080850152901b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016608883015240609c82015260bc01604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001815290829052805160209091012060058190557f9a908e73000000000000000000000000000000000000000000000000000000008252600160048301526024820181905291507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303815f875af11580156111bb573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111df9190613b43565b508c5f60026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508b60025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550886003908161126f9190613ba9565b50600461127c8982613ba9565b508c60085f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062069780600760106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507f060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f85838e60405161131c93929190613cc0565b60405180910390a1505050505050801561138c575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b60015473ffffffffffffffffffffffffffffffffffffffff1633146113e6576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001545f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff1673ffffffffffffffffffffffffffffffffffffffff9092166201000081029290921790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e9060200160405180910390a1565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff1633146114bd576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60085473ffffffffffffffffffffffffffffffffffffffff1661150c576040517fc89374d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb90602001610c44565b60085473ffffffffffffffffffffffffffffffffffffffff1680158015906115bd575073ffffffffffffffffffffffffffffffffffffffff81163314155b156115f4576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b4262093a807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166330c27dde6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611662573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116869190613b43565b6116909190613cfe565b67ffffffffffffffff1611156116d2576040517f3d49ed4c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815f81900361170d576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8811115611749576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075467ffffffffffffffff8082169161177191849168010000000000000000900416613d1e565b11156117a9576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6007546005546801000000000000000090910467ffffffffffffffff16905f5b83811015611a3e575f8787838181106117e4576117e4613d31565b90506020028101906117f69190613d5e565b6117ff90613d9a565b90508361180b81613e23565b825180516020918201208185015160408087015160608801519151959a509295505f94611877948794929101938452602084019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166040830152604882015260680190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff89165f908152600690935291205490915081146118ff576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff86165f90815260066020526040812055611923600188613b30565b84036119925742600760109054906101000a900467ffffffffffffffff1684604001516119509190613cfe565b67ffffffffffffffff161115611992576040517fc44a082100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208381015160408086015160608088015183519586018b90529285018790528481019390935260c01b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808401523390911b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088830152609c82015260bc0160405160208183030381529060405280519060200120945050505080806001019150506117c9565b50611ab47f000000000000000000000000000000000000000000000000000000000000000084611a6c6107df565b611a769190613e4f565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169190613140565b60058190556007805467ffffffffffffffff841668010000000000000000027fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff9091161790556040517f9a908e730000000000000000000000000000000000000000000000000000000081525f9073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690639a908e7390611b85908790869060040167ffffffffffffffff929092168252602082015260400190565b6020604051808303815f875af1158015611ba1573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611bc59190613b43565b60405190915067ffffffffffffffff8216907f648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4905f90a250505050505050565b60605f85858573a40d5f56745a118d0906a34e69aec8c0db1cb8fa5f87604051602401611c3796959493929190613e66565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff811bff70000000000000000000000000000000000000000000000000000000017905283519091506060905f03611d875760f9601f8351611ccb9190613ef0565b6040518060400160405280600881526020017f80808401c9c380940000000000000000000000000000000000000000000000008152507f00000000000000000000000000000000000000000000000000000000000000006040518060400160405280600281526020017f80b800000000000000000000000000000000000000000000000000000000000081525060e487604051602001611d719796959493929190613f21565b6040516020818303038152906040529050611e8b565b815161ffff1015611dc4576040517f248b8f8200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815160f9611dd3602083613ef0565b6040518060400160405280600881526020017f80808401c9c380940000000000000000000000000000000000000000000000008152507f00000000000000000000000000000000000000000000000000000000000000006040518060400160405280600281526020017f80b90000000000000000000000000000000000000000000000000000000000008152508588604051602001611e789796959493929190613ffc565b6040516020818303038152906040529150505b8051602080830191909120604080515f80825293810180835292909252601b908201526405ca1ab1e06060820152635ca1ab1e608082015260019060a0016020604051602081039080840390855afa158015611ee9573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116611f61576040517fcd16196600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040515f90611fa69084906405ca1ab1e090635ca1ab1e90601b907fff00000000000000000000000000000000000000000000000000000000000000906020016140c9565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529450505050505b949350505050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314612039576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001610c44565b60025473ffffffffffffffffffffffffffffffffffffffff1633146120fd576040517f11e7be1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b845f819003612138576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8811115612174576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61217f602442613d1e565b8467ffffffffffffffff1611156121c2576040517f0a00feb300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379e2cf976040518163ffffffff1660e01b81526004015f604051808303815f87803b158015612227575f5ffd5b505af1158015612239573d5f5f3e3d5ffd5b50506040517fef4eeb3500000000000000000000000000000000000000000000000000000000815263ffffffff881660048201525f92507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16915063ef4eeb3590602401602060405180830381865afa1580156122cd573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906122f19190613a1e565b90508061232a576040517fa60721e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6007546005546801000000000000000090910467ffffffffffffffff1690815f5b8581101561263b575f8c8c8381811061236657612366613d31565b90506020028101906123789190613d5e565b61238190613d9a565b8051805160209091012060408201519192509067ffffffffffffffff161561256157856123ad81613e23565b9650505f8183602001518460400151856060015160405160200161240f9493929190938452602084019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166040830152604882015260680190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff8a165f90815260069093529120549091508114612497576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208381015160408086015160608088015183519586018c90529285018790528481019390935260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166080840152908c901b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088830152609c82015260bc0160405160208183030381529060405280519060200120955060065f8867ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f905550612631565b8151516201d4c010156125a0576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160208101879052908101829052606080820189905260c08d901b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808301528a901b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660888201525f609c82015260bc016040516020818303038152906040528051906020012094505b505060010161234b565b5060075467ffffffffffffffff9081169084161115612686576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60058290558467ffffffffffffffff8481169083161461273b575f6126ab8386613a62565b90506126c167ffffffffffffffff821683613b30565b91506126fa7f00000000000000000000000000000000000000000000000000000000000000008267ffffffffffffffff16611a6c6107df565b50600780547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8716021790555b612837337f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663477fa2706040518163ffffffff1660e01b8152600401602060405180830381865afa1580156127ca573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906127ee9190613a1e565b6127f89190613e4f565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016929190613219565b6040517f9a908e7300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff87166004820152602481018490525f907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303815f875af11580156128d3573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906128f79190613b43565b9050888414612932576040517fda5bceb900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff167f3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e7668760405161296e91815260200190565b60405180910390a250505050505050505050505050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff1633146129db576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60036129e78282613ba9565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b2081604051610c44919061353f565b60085473ffffffffffffffffffffffffffffffffffffffff168015801590612a55575073ffffffffffffffffffffffffffffffffffffffff81163314155b15612a8c576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166315064c966040518163ffffffff1660e01b8152600401602060405180830381865afa158015612af5573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612b199190613b11565b15612b50576040517f39258d1800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663604691696040518163ffffffff1660e01b8152600401602060405180830381865afa158015612bba573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612bde9190613a1e565b905082811115612c1a576040517f2354600f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611388841115612c56576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612c9873ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016333084613219565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612d02573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612d269190613a1e565b6007805491925067ffffffffffffffff909116905f612d4483613e23565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550508585604051612d7b92919061411b565b6040519081900390208142612d91600143613b30565b60408051602081019590955284019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166060830152406068820152608801604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012060075467ffffffffffffffff165f9081526006909352912055323303612e87576007546040805183815233602082015260608183018190525f90820152905167ffffffffffffffff909216917ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319181900360800190a2612ed6565b60075460405167ffffffffffffffff909116907ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc93190612ecd90849033908b908b9061412a565b60405180910390a25b505050505050565b606073ffffffffffffffffffffffffffffffffffffffff82161561313b576040517fc00f14ab00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063c00f14ab906024015f60405180830381865afa158015612f85573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052612fca919081019061419a565b6040517f318aee3d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301529192505f9182917f00000000000000000000000000000000000000000000000000000000000000009091169063318aee3d906024016040805180830381865afa15801561305c573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613080919061420f565b915091508163ffffffff165f146130f7576009805463ffffffff841674010000000000000000000000000000000000000000027fffffffffffffffff00000000000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff841617179055613138565b600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff86161790555b50505b919050565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526132149084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261327d565b505050565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526132779085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401613192565b50505050565b5f6132de826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166133889092919063ffffffff16565b80519091501561321457808060200190518101906132fc9190613b11565b613214576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610f0a565b6060611fdb84845f85855f5f8673ffffffffffffffffffffffffffffffffffffffff1685876040516133ba9190614247565b5f6040518083038185875af1925050503d805f81146133f4576040519150601f19603f3d011682016040523d82523d5f602084013e6133f9565b606091505b509150915061340a87838387613415565b979650505050505050565b606083156134aa5782515f036134a35773ffffffffffffffffffffffffffffffffffffffff85163b6134a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610f0a565b5081611fdb565b611fdb83838151156134bf5781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f0a919061353f565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f61355160208301846134f3565b9392505050565b67ffffffffffffffff8116811461356d575f5ffd5b50565b73ffffffffffffffffffffffffffffffffffffffff8116811461356d575f5ffd5b5f5f5f606084860312156135a3575f5ffd5b83356135ae81613558565b92506020840135915060408401356135c581613570565b809150509250925092565b5f602082840312156135e0575f5ffd5b813561355181613558565b5f5f604083850312156135fc575f5ffd5b823561360781613558565b946020939093013593505050565b5f60208284031215613625575f5ffd5b813561355181613570565b63ffffffff8116811461356d575f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156136b5576136b5613641565b604052919050565b5f67ffffffffffffffff8211156136d6576136d6613641565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b5f82601f830112613711575f5ffd5b8135602083015f613729613724846136bd565b61366e565b905082815285838301111561373c575f5ffd5b828260208301375f92810160200192909252509392505050565b5f5f5f5f5f5f60c0878903121561376b575f5ffd5b863561377681613570565b9550602087013561378681613570565b9450604087013561379681613630565b935060608701356137a681613570565b9250608087013567ffffffffffffffff8111156137c1575f5ffd5b6137cd89828a01613702565b92505060a087013567ffffffffffffffff8111156137e9575f5ffd5b6137f589828a01613702565b9150509295509295509295565b5f5f83601f840112613812575f5ffd5b50813567ffffffffffffffff811115613829575f5ffd5b6020830191508360208260051b8501011115613843575f5ffd5b9250929050565b5f5f6020838503121561385b575f5ffd5b823567ffffffffffffffff811115613871575f5ffd5b61387d85828601613802565b90969095509350505050565b5f5f5f5f6080858703121561389c575f5ffd5b84356138a781613630565b935060208501356138b781613570565b925060408501356138c781613630565b9150606085013567ffffffffffffffff8111156138e2575f5ffd5b6138ee87828801613702565b91505092959194509250565b5f5f5f5f5f5f60a0878903121561390f575f5ffd5b863567ffffffffffffffff811115613925575f5ffd5b61393189828a01613802565b909750955050602087013561394581613630565b9350604087013561395581613558565b925060608701359150608087013561396c81613570565b809150509295509295509295565b5f6020828403121561398a575f5ffd5b813567ffffffffffffffff8111156139a0575f5ffd5b611fdb84828501613702565b5f5f5f604084860312156139be575f5ffd5b833567ffffffffffffffff8111156139d4575f5ffd5b8401601f810186136139e4575f5ffd5b803567ffffffffffffffff8111156139fa575f5ffd5b866020828401011115613a0b575f5ffd5b6020918201979096509401359392505050565b5f60208284031215613a2e575f5ffd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b67ffffffffffffffff8281168282160390811115613a8257613a82613a35565b92915050565b5f82613abb577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b600181811c90821680613ad457607f821691505b602082108103613b0b577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b5f60208284031215613b21575f5ffd5b81518015158114613551575f5ffd5b81810381811115613a8257613a82613a35565b5f60208284031215613b53575f5ffd5b815161355181613558565b601f82111561321457805f5260205f20601f840160051c81016020851015613b835750805b601f840160051c820191505b81811015613ba2575f8155600101613b8f565b5050505050565b815167ffffffffffffffff811115613bc357613bc3613641565b613bd781613bd18454613ac0565b84613b5e565b6020601f821160018114613c28575f8315613bf25750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455613ba2565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015613c755787850151825560209485019460019092019101613c55565b5084821015613cb157868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b606081525f613cd260608301866134f3565b905083602083015273ffffffffffffffffffffffffffffffffffffffff83166040830152949350505050565b67ffffffffffffffff8181168382160190811115613a8257613a82613a35565b80820180821115613a8257613a82613a35565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f82357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81833603018112613d90575f5ffd5b9190910192915050565b5f60808236031215613daa575f5ffd5b6040516080810167ffffffffffffffff81118282101715613dcd57613dcd613641565b604052823567ffffffffffffffff811115613de6575f5ffd5b613df236828601613702565b825250602083810135908201526040830135613e0d81613558565b6040820152606092830135928101929092525090565b5f67ffffffffffffffff821667ffffffffffffffff8103613e4657613e46613a35565b60010192915050565b8082028115828204841417613a8257613a82613a35565b63ffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015263ffffffff8516604082015273ffffffffffffffffffffffffffffffffffffffff8416606082015273ffffffffffffffffffffffffffffffffffffffff8316608082015260c060a08201525f613ee460c08301846134f3565b98975050505050505050565b61ffff8181168382160190811115613a8257613a82613a35565b5f81518060208401855e5f93019283525090919050565b7fff000000000000000000000000000000000000000000000000000000000000008860f81b1681527fffff0000000000000000000000000000000000000000000000000000000000008760f01b1660018201525f613f826003830188613f0a565b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008760601b168152613fb76014820187613f0a565b90507fff000000000000000000000000000000000000000000000000000000000000008560f81b168152613fee6001820185613f0a565b9a9950505050505050505050565b7fff000000000000000000000000000000000000000000000000000000000000008860f81b1681527fffff0000000000000000000000000000000000000000000000000000000000008760f01b1660018201525f61405d6003830188613f0a565b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008760601b1681526140926014820187613f0a565b90507fffff0000000000000000000000000000000000000000000000000000000000008560f01b168152613fee6002820185613f0a565b5f6140d48288613f0a565b9586525050602084019290925260f81b7fff000000000000000000000000000000000000000000000000000000000000009081166040840152166041820152604201919050565b818382375f9101908152919050565b84815273ffffffffffffffffffffffffffffffffffffffff8416602082015260606040820152816060820152818360808301375f818301608090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01601019392505050565b5f602082840312156141aa575f5ffd5b815167ffffffffffffffff8111156141c0575f5ffd5b8201601f810184136141d0575f5ffd5b80516141de613724826136bd565b8181528560208385010111156141f2575f5ffd5b8160208401602083015e5f91810160200191909152949350505050565b5f5f60408385031215614220575f5ffd5b825161422b81613630565b602084015190925061423c81613570565b809150509250929050565b5f6135518284613f0a56fea264697066735822122019be8f7dd8219c78d3e7f131ab3866c48d967e7be2b6353626bfc9b90c1ecf0864736f6c634300081c0033",
}

// PolygonzkevmetrogABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonzkevmetrogMetaData.ABI instead.
var PolygonzkevmetrogABI = PolygonzkevmetrogMetaData.ABI

// PolygonzkevmetrogBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygonzkevmetrogMetaData.Bin instead.
var PolygonzkevmetrogBin = PolygonzkevmetrogMetaData.Bin

// DeployPolygonzkevmetrog deploys a new Ethereum contract, binding an instance of Polygonzkevmetrog to it.
func DeployPolygonzkevmetrog(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address, _rollupManager common.Address) (common.Address, *types.Transaction, *Polygonzkevmetrog, error) {
	parsed, err := PolygonzkevmetrogMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygonzkevmetrogBin), backend, _globalExitRootManager, _pol, _bridgeAddress, _rollupManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonzkevmetrog{PolygonzkevmetrogCaller: PolygonzkevmetrogCaller{contract: contract}, PolygonzkevmetrogTransactor: PolygonzkevmetrogTransactor{contract: contract}, PolygonzkevmetrogFilterer: PolygonzkevmetrogFilterer{contract: contract}}, nil
}

// Polygonzkevmetrog is an auto generated Go binding around an Ethereum contract.
type Polygonzkevmetrog struct {
	PolygonzkevmetrogCaller     // Read-only binding to the contract
	PolygonzkevmetrogTransactor // Write-only binding to the contract
	PolygonzkevmetrogFilterer   // Log filterer for contract events
}

// PolygonzkevmetrogCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonzkevmetrogCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmetrogTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonzkevmetrogTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmetrogFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonzkevmetrogFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmetrogSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonzkevmetrogSession struct {
	Contract     *Polygonzkevmetrog // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PolygonzkevmetrogCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonzkevmetrogCallerSession struct {
	Contract *PolygonzkevmetrogCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PolygonzkevmetrogTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonzkevmetrogTransactorSession struct {
	Contract     *PolygonzkevmetrogTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PolygonzkevmetrogRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonzkevmetrogRaw struct {
	Contract *Polygonzkevmetrog // Generic contract binding to access the raw methods on
}

// PolygonzkevmetrogCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonzkevmetrogCallerRaw struct {
	Contract *PolygonzkevmetrogCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonzkevmetrogTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonzkevmetrogTransactorRaw struct {
	Contract *PolygonzkevmetrogTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonzkevmetrog creates a new instance of Polygonzkevmetrog, bound to a specific deployed contract.
func NewPolygonzkevmetrog(address common.Address, backend bind.ContractBackend) (*Polygonzkevmetrog, error) {
	contract, err := bindPolygonzkevmetrog(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmetrog{PolygonzkevmetrogCaller: PolygonzkevmetrogCaller{contract: contract}, PolygonzkevmetrogTransactor: PolygonzkevmetrogTransactor{contract: contract}, PolygonzkevmetrogFilterer: PolygonzkevmetrogFilterer{contract: contract}}, nil
}

// NewPolygonzkevmetrogCaller creates a new read-only instance of Polygonzkevmetrog, bound to a specific deployed contract.
func NewPolygonzkevmetrogCaller(address common.Address, caller bind.ContractCaller) (*PolygonzkevmetrogCaller, error) {
	contract, err := bindPolygonzkevmetrog(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmetrogCaller{contract: contract}, nil
}

// NewPolygonzkevmetrogTransactor creates a new write-only instance of Polygonzkevmetrog, bound to a specific deployed contract.
func NewPolygonzkevmetrogTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonzkevmetrogTransactor, error) {
	contract, err := bindPolygonzkevmetrog(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmetrogTransactor{contract: contract}, nil
}

// NewPolygonzkevmetrogFilterer creates a new log filterer instance of Polygonzkevmetrog, bound to a specific deployed contract.
func NewPolygonzkevmetrogFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonzkevmetrogFilterer, error) {
	contract, err := bindPolygonzkevmetrog(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmetrogFilterer{contract: contract}, nil
}

// bindPolygonzkevmetrog binds a generic wrapper to an already deployed contract.
func bindPolygonzkevmetrog(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonzkevmetrogMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmetrog *PolygonzkevmetrogRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmetrog.Contract.PolygonzkevmetrogCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmetrog *PolygonzkevmetrogRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.PolygonzkevmetrogTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmetrog *PolygonzkevmetrogRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.PolygonzkevmetrogTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmetrog.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmetrog *PolygonzkevmetrogTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmetrog *PolygonzkevmetrogTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.contract.Transact(opts, method, params...)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) GLOBALEXITROOTMANAGERL2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "GLOBAL_EXIT_ROOT_MANAGER_L2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _Polygonzkevmetrog.Contract.GLOBALEXITROOTMANAGERL2(&_Polygonzkevmetrog.CallOpts)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _Polygonzkevmetrog.Contract.GLOBALEXITROOTMANAGERL2(&_Polygonzkevmetrog.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) INITIALIZETXBRIDGELISTLENLEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_LIST_LEN_LEN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _Polygonzkevmetrog.Contract.INITIALIZETXBRIDGELISTLENLEN(&_Polygonzkevmetrog.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _Polygonzkevmetrog.Contract.INITIALIZETXBRIDGELISTLENLEN(&_Polygonzkevmetrog.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) INITIALIZETXBRIDGEPARAMS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _Polygonzkevmetrog.Contract.INITIALIZETXBRIDGEPARAMS(&_Polygonzkevmetrog.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _Polygonzkevmetrog.Contract.INITIALIZETXBRIDGEPARAMS(&_Polygonzkevmetrog.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _Polygonzkevmetrog.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_Polygonzkevmetrog.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _Polygonzkevmetrog.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_Polygonzkevmetrog.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _Polygonzkevmetrog.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_Polygonzkevmetrog.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _Polygonzkevmetrog.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_Polygonzkevmetrog.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) INITIALIZETXCONSTANTBYTES(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _Polygonzkevmetrog.Contract.INITIALIZETXCONSTANTBYTES(&_Polygonzkevmetrog.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _Polygonzkevmetrog.Contract.INITIALIZETXCONSTANTBYTES(&_Polygonzkevmetrog.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) INITIALIZETXCONSTANTBYTESEMPTYMETADATA(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _Polygonzkevmetrog.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_Polygonzkevmetrog.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _Polygonzkevmetrog.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_Polygonzkevmetrog.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) INITIALIZETXDATALENEMPTYMETADATA(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "INITIALIZE_TX_DATA_LEN_EMPTY_METADATA")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _Polygonzkevmetrog.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_Polygonzkevmetrog.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _Polygonzkevmetrog.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_Polygonzkevmetrog.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) INITIALIZETXEFFECTIVEPERCENTAGE(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "INITIALIZE_TX_EFFECTIVE_PERCENTAGE")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _Polygonzkevmetrog.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_Polygonzkevmetrog.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _Polygonzkevmetrog.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_Polygonzkevmetrog.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) SIGNATUREINITIALIZETXR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_R")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _Polygonzkevmetrog.Contract.SIGNATUREINITIALIZETXR(&_Polygonzkevmetrog.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _Polygonzkevmetrog.Contract.SIGNATUREINITIALIZETXR(&_Polygonzkevmetrog.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) SIGNATUREINITIALIZETXS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_S")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _Polygonzkevmetrog.Contract.SIGNATUREINITIALIZETXS(&_Polygonzkevmetrog.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _Polygonzkevmetrog.Contract.SIGNATUREINITIALIZETXS(&_Polygonzkevmetrog.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) SIGNATUREINITIALIZETXV(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_V")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _Polygonzkevmetrog.Contract.SIGNATUREINITIALIZETXV(&_Polygonzkevmetrog.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _Polygonzkevmetrog.Contract.SIGNATUREINITIALIZETXV(&_Polygonzkevmetrog.CallOpts)
}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) TIMESTAMPRANGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "TIMESTAMP_RANGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) TIMESTAMPRANGE() (*big.Int, error) {
	return _Polygonzkevmetrog.Contract.TIMESTAMPRANGE(&_Polygonzkevmetrog.CallOpts)
}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) TIMESTAMPRANGE() (*big.Int, error) {
	return _Polygonzkevmetrog.Contract.TIMESTAMPRANGE(&_Polygonzkevmetrog.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) Admin() (common.Address, error) {
	return _Polygonzkevmetrog.Contract.Admin(&_Polygonzkevmetrog.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) Admin() (common.Address, error) {
	return _Polygonzkevmetrog.Contract.Admin(&_Polygonzkevmetrog.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) BridgeAddress() (common.Address, error) {
	return _Polygonzkevmetrog.Contract.BridgeAddress(&_Polygonzkevmetrog.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) BridgeAddress() (common.Address, error) {
	return _Polygonzkevmetrog.Contract.BridgeAddress(&_Polygonzkevmetrog.CallOpts)
}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) CalculatePolPerForceBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "calculatePolPerForceBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) CalculatePolPerForceBatch() (*big.Int, error) {
	return _Polygonzkevmetrog.Contract.CalculatePolPerForceBatch(&_Polygonzkevmetrog.CallOpts)
}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) CalculatePolPerForceBatch() (*big.Int, error) {
	return _Polygonzkevmetrog.Contract.CalculatePolPerForceBatch(&_Polygonzkevmetrog.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) ForceBatchAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "forceBatchAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) ForceBatchAddress() (common.Address, error) {
	return _Polygonzkevmetrog.Contract.ForceBatchAddress(&_Polygonzkevmetrog.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) ForceBatchAddress() (common.Address, error) {
	return _Polygonzkevmetrog.Contract.ForceBatchAddress(&_Polygonzkevmetrog.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonzkevmetrog.Contract.ForceBatchTimeout(&_Polygonzkevmetrog.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonzkevmetrog.Contract.ForceBatchTimeout(&_Polygonzkevmetrog.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonzkevmetrog.Contract.ForcedBatches(&_Polygonzkevmetrog.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonzkevmetrog.Contract.ForcedBatches(&_Polygonzkevmetrog.CallOpts, arg0)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) GasTokenAddress() (common.Address, error) {
	return _Polygonzkevmetrog.Contract.GasTokenAddress(&_Polygonzkevmetrog.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) GasTokenAddress() (common.Address, error) {
	return _Polygonzkevmetrog.Contract.GasTokenAddress(&_Polygonzkevmetrog.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) GasTokenNetwork() (uint32, error) {
	return _Polygonzkevmetrog.Contract.GasTokenNetwork(&_Polygonzkevmetrog.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) GasTokenNetwork() (uint32, error) {
	return _Polygonzkevmetrog.Contract.GasTokenNetwork(&_Polygonzkevmetrog.CallOpts)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) GenerateInitializeTransaction(opts *bind.CallOpts, networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "generateInitializeTransaction", networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) GenerateInitializeTransaction(networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _Polygonzkevmetrog.Contract.GenerateInitializeTransaction(&_Polygonzkevmetrog.CallOpts, networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) GenerateInitializeTransaction(networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _Polygonzkevmetrog.Contract.GenerateInitializeTransaction(&_Polygonzkevmetrog.CallOpts, networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonzkevmetrog.Contract.GlobalExitRootManager(&_Polygonzkevmetrog.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonzkevmetrog.Contract.GlobalExitRootManager(&_Polygonzkevmetrog.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) LastAccInputHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "lastAccInputHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) LastAccInputHash() ([32]byte, error) {
	return _Polygonzkevmetrog.Contract.LastAccInputHash(&_Polygonzkevmetrog.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) LastAccInputHash() ([32]byte, error) {
	return _Polygonzkevmetrog.Contract.LastAccInputHash(&_Polygonzkevmetrog.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) LastForceBatch() (uint64, error) {
	return _Polygonzkevmetrog.Contract.LastForceBatch(&_Polygonzkevmetrog.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) LastForceBatch() (uint64, error) {
	return _Polygonzkevmetrog.Contract.LastForceBatch(&_Polygonzkevmetrog.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonzkevmetrog.Contract.LastForceBatchSequenced(&_Polygonzkevmetrog.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonzkevmetrog.Contract.LastForceBatchSequenced(&_Polygonzkevmetrog.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) NetworkName() (string, error) {
	return _Polygonzkevmetrog.Contract.NetworkName(&_Polygonzkevmetrog.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) NetworkName() (string, error) {
	return _Polygonzkevmetrog.Contract.NetworkName(&_Polygonzkevmetrog.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) PendingAdmin() (common.Address, error) {
	return _Polygonzkevmetrog.Contract.PendingAdmin(&_Polygonzkevmetrog.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) PendingAdmin() (common.Address, error) {
	return _Polygonzkevmetrog.Contract.PendingAdmin(&_Polygonzkevmetrog.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) Pol() (common.Address, error) {
	return _Polygonzkevmetrog.Contract.Pol(&_Polygonzkevmetrog.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) Pol() (common.Address, error) {
	return _Polygonzkevmetrog.Contract.Pol(&_Polygonzkevmetrog.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) RollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "rollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) RollupManager() (common.Address, error) {
	return _Polygonzkevmetrog.Contract.RollupManager(&_Polygonzkevmetrog.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) RollupManager() (common.Address, error) {
	return _Polygonzkevmetrog.Contract.RollupManager(&_Polygonzkevmetrog.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) TrustedSequencer() (common.Address, error) {
	return _Polygonzkevmetrog.Contract.TrustedSequencer(&_Polygonzkevmetrog.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) TrustedSequencer() (common.Address, error) {
	return _Polygonzkevmetrog.Contract.TrustedSequencer(&_Polygonzkevmetrog.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonzkevmetrog *PolygonzkevmetrogCaller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonzkevmetrog.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) TrustedSequencerURL() (string, error) {
	return _Polygonzkevmetrog.Contract.TrustedSequencerURL(&_Polygonzkevmetrog.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonzkevmetrog *PolygonzkevmetrogCallerSession) TrustedSequencerURL() (string, error) {
	return _Polygonzkevmetrog.Contract.TrustedSequencerURL(&_Polygonzkevmetrog.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmetrog.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.AcceptAdminRole(&_Polygonzkevmetrog.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.AcceptAdminRole(&_Polygonzkevmetrog.TransactOpts)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogTransactor) ForceBatch(opts *bind.TransactOpts, transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmetrog.contract.Transact(opts, "forceBatch", transactions, polAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) ForceBatch(transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.ForceBatch(&_Polygonzkevmetrog.TransactOpts, transactions, polAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogTransactorSession) ForceBatch(transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.ForceBatch(&_Polygonzkevmetrog.TransactOpts, transactions, polAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonzkevmetrog.contract.Transact(opts, "initialize", _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.Initialize(&_Polygonzkevmetrog.TransactOpts, _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogTransactorSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.Initialize(&_Polygonzkevmetrog.TransactOpts, _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogTransactor) OnVerifyBatches(opts *bind.TransactOpts, lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonzkevmetrog.contract.Transact(opts, "onVerifyBatches", lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.OnVerifyBatches(&_Polygonzkevmetrog.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogTransactorSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.OnVerifyBatches(&_Polygonzkevmetrog.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x669adece.
//
// Solidity: function rollbackBatches(uint64 targetBatch, bytes32 accInputHashToRollback) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogTransactor) RollbackBatches(opts *bind.TransactOpts, targetBatch uint64, accInputHashToRollback [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmetrog.contract.Transact(opts, "rollbackBatches", targetBatch, accInputHashToRollback)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x669adece.
//
// Solidity: function rollbackBatches(uint64 targetBatch, bytes32 accInputHashToRollback) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) RollbackBatches(targetBatch uint64, accInputHashToRollback [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.RollbackBatches(&_Polygonzkevmetrog.TransactOpts, targetBatch, accInputHashToRollback)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x669adece.
//
// Solidity: function rollbackBatches(uint64 targetBatch, bytes32 accInputHashToRollback) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogTransactorSession) RollbackBatches(targetBatch uint64, accInputHashToRollback [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.RollbackBatches(&_Polygonzkevmetrog.TransactOpts, targetBatch, accInputHashToRollback)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xb910e0f9.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint32 l1InfoTreeLeafCount, uint64 maxSequenceTimestamp, bytes32 expectedFinalAccInputHash, address l2Coinbase) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogTransactor) SequenceBatches(opts *bind.TransactOpts, batches []PolygonRollupBaseEtrogBatchData, l1InfoTreeLeafCount uint32, maxSequenceTimestamp uint64, expectedFinalAccInputHash [32]byte, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonzkevmetrog.contract.Transact(opts, "sequenceBatches", batches, l1InfoTreeLeafCount, maxSequenceTimestamp, expectedFinalAccInputHash, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xb910e0f9.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint32 l1InfoTreeLeafCount, uint64 maxSequenceTimestamp, bytes32 expectedFinalAccInputHash, address l2Coinbase) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) SequenceBatches(batches []PolygonRollupBaseEtrogBatchData, l1InfoTreeLeafCount uint32, maxSequenceTimestamp uint64, expectedFinalAccInputHash [32]byte, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.SequenceBatches(&_Polygonzkevmetrog.TransactOpts, batches, l1InfoTreeLeafCount, maxSequenceTimestamp, expectedFinalAccInputHash, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xb910e0f9.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint32 l1InfoTreeLeafCount, uint64 maxSequenceTimestamp, bytes32 expectedFinalAccInputHash, address l2Coinbase) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogTransactorSession) SequenceBatches(batches []PolygonRollupBaseEtrogBatchData, l1InfoTreeLeafCount uint32, maxSequenceTimestamp uint64, expectedFinalAccInputHash [32]byte, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.SequenceBatches(&_Polygonzkevmetrog.TransactOpts, batches, l1InfoTreeLeafCount, maxSequenceTimestamp, expectedFinalAccInputHash, l2Coinbase)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogTransactor) SequenceForceBatches(opts *bind.TransactOpts, batches []PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _Polygonzkevmetrog.contract.Transact(opts, "sequenceForceBatches", batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) SequenceForceBatches(batches []PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.SequenceForceBatches(&_Polygonzkevmetrog.TransactOpts, batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogTransactorSession) SequenceForceBatches(batches []PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.SequenceForceBatches(&_Polygonzkevmetrog.TransactOpts, batches)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogTransactor) SetForceBatchAddress(opts *bind.TransactOpts, newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonzkevmetrog.contract.Transact(opts, "setForceBatchAddress", newForceBatchAddress)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) SetForceBatchAddress(newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.SetForceBatchAddress(&_Polygonzkevmetrog.TransactOpts, newForceBatchAddress)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogTransactorSession) SetForceBatchAddress(newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.SetForceBatchAddress(&_Polygonzkevmetrog.TransactOpts, newForceBatchAddress)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogTransactor) SetForceBatchTimeout(opts *bind.TransactOpts, newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmetrog.contract.Transact(opts, "setForceBatchTimeout", newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.SetForceBatchTimeout(&_Polygonzkevmetrog.TransactOpts, newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogTransactorSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.SetForceBatchTimeout(&_Polygonzkevmetrog.TransactOpts, newforceBatchTimeout)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogTransactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonzkevmetrog.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.SetTrustedSequencer(&_Polygonzkevmetrog.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogTransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.SetTrustedSequencer(&_Polygonzkevmetrog.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogTransactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonzkevmetrog.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.SetTrustedSequencerURL(&_Polygonzkevmetrog.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogTransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.SetTrustedSequencerURL(&_Polygonzkevmetrog.TransactOpts, newTrustedSequencerURL)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonzkevmetrog.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.TransferAdminRole(&_Polygonzkevmetrog.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonzkevmetrog *PolygonzkevmetrogTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonzkevmetrog.Contract.TransferAdminRole(&_Polygonzkevmetrog.TransactOpts, newPendingAdmin)
}

// PolygonzkevmetrogAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Polygonzkevmetrog contract.
type PolygonzkevmetrogAcceptAdminRoleIterator struct {
	Event *PolygonzkevmetrogAcceptAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmetrogAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmetrogAcceptAdminRole)
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
		it.Event = new(PolygonzkevmetrogAcceptAdminRole)
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
func (it *PolygonzkevmetrogAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmetrogAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmetrogAcceptAdminRole represents a AcceptAdminRole event raised by the Polygonzkevmetrog contract.
type PolygonzkevmetrogAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*PolygonzkevmetrogAcceptAdminRoleIterator, error) {

	logs, sub, err := _Polygonzkevmetrog.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmetrogAcceptAdminRoleIterator{contract: _Polygonzkevmetrog.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonzkevmetrogAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmetrog.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmetrogAcceptAdminRole)
				if err := _Polygonzkevmetrog.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
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
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) ParseAcceptAdminRole(log types.Log) (*PolygonzkevmetrogAcceptAdminRole, error) {
	event := new(PolygonzkevmetrogAcceptAdminRole)
	if err := _Polygonzkevmetrog.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmetrogForceBatchIterator is returned from FilterForceBatch and is used to iterate over the raw logs and unpacked data for ForceBatch events raised by the Polygonzkevmetrog contract.
type PolygonzkevmetrogForceBatchIterator struct {
	Event *PolygonzkevmetrogForceBatch // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmetrogForceBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmetrogForceBatch)
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
		it.Event = new(PolygonzkevmetrogForceBatch)
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
func (it *PolygonzkevmetrogForceBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmetrogForceBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmetrogForceBatch represents a ForceBatch event raised by the Polygonzkevmetrog contract.
type PolygonzkevmetrogForceBatch struct {
	ForceBatchNum      uint64
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Transactions       []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterForceBatch is a free log retrieval operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) FilterForceBatch(opts *bind.FilterOpts, forceBatchNum []uint64) (*PolygonzkevmetrogForceBatchIterator, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Polygonzkevmetrog.contract.FilterLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmetrogForceBatchIterator{contract: _Polygonzkevmetrog.contract, event: "ForceBatch", logs: logs, sub: sub}, nil
}

// WatchForceBatch is a free log subscription operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) WatchForceBatch(opts *bind.WatchOpts, sink chan<- *PolygonzkevmetrogForceBatch, forceBatchNum []uint64) (event.Subscription, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Polygonzkevmetrog.contract.WatchLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmetrogForceBatch)
				if err := _Polygonzkevmetrog.contract.UnpackLog(event, "ForceBatch", log); err != nil {
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
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) ParseForceBatch(log types.Log) (*PolygonzkevmetrogForceBatch, error) {
	event := new(PolygonzkevmetrogForceBatch)
	if err := _Polygonzkevmetrog.contract.UnpackLog(event, "ForceBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmetrogInitialSequenceBatchesIterator is returned from FilterInitialSequenceBatches and is used to iterate over the raw logs and unpacked data for InitialSequenceBatches events raised by the Polygonzkevmetrog contract.
type PolygonzkevmetrogInitialSequenceBatchesIterator struct {
	Event *PolygonzkevmetrogInitialSequenceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmetrogInitialSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmetrogInitialSequenceBatches)
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
		it.Event = new(PolygonzkevmetrogInitialSequenceBatches)
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
func (it *PolygonzkevmetrogInitialSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmetrogInitialSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmetrogInitialSequenceBatches represents a InitialSequenceBatches event raised by the Polygonzkevmetrog contract.
type PolygonzkevmetrogInitialSequenceBatches struct {
	Transactions       []byte
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterInitialSequenceBatches is a free log retrieval operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) FilterInitialSequenceBatches(opts *bind.FilterOpts) (*PolygonzkevmetrogInitialSequenceBatchesIterator, error) {

	logs, sub, err := _Polygonzkevmetrog.contract.FilterLogs(opts, "InitialSequenceBatches")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmetrogInitialSequenceBatchesIterator{contract: _Polygonzkevmetrog.contract, event: "InitialSequenceBatches", logs: logs, sub: sub}, nil
}

// WatchInitialSequenceBatches is a free log subscription operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) WatchInitialSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmetrogInitialSequenceBatches) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmetrog.contract.WatchLogs(opts, "InitialSequenceBatches")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmetrogInitialSequenceBatches)
				if err := _Polygonzkevmetrog.contract.UnpackLog(event, "InitialSequenceBatches", log); err != nil {
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
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) ParseInitialSequenceBatches(log types.Log) (*PolygonzkevmetrogInitialSequenceBatches, error) {
	event := new(PolygonzkevmetrogInitialSequenceBatches)
	if err := _Polygonzkevmetrog.contract.UnpackLog(event, "InitialSequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmetrogInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonzkevmetrog contract.
type PolygonzkevmetrogInitializedIterator struct {
	Event *PolygonzkevmetrogInitialized // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmetrogInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmetrogInitialized)
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
		it.Event = new(PolygonzkevmetrogInitialized)
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
func (it *PolygonzkevmetrogInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmetrogInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmetrogInitialized represents a Initialized event raised by the Polygonzkevmetrog contract.
type PolygonzkevmetrogInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) FilterInitialized(opts *bind.FilterOpts) (*PolygonzkevmetrogInitializedIterator, error) {

	logs, sub, err := _Polygonzkevmetrog.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmetrogInitializedIterator{contract: _Polygonzkevmetrog.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PolygonzkevmetrogInitialized) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmetrog.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmetrogInitialized)
				if err := _Polygonzkevmetrog.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) ParseInitialized(log types.Log) (*PolygonzkevmetrogInitialized, error) {
	event := new(PolygonzkevmetrogInitialized)
	if err := _Polygonzkevmetrog.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmetrogRollbackBatchesIterator is returned from FilterRollbackBatches and is used to iterate over the raw logs and unpacked data for RollbackBatches events raised by the Polygonzkevmetrog contract.
type PolygonzkevmetrogRollbackBatchesIterator struct {
	Event *PolygonzkevmetrogRollbackBatches // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmetrogRollbackBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmetrogRollbackBatches)
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
		it.Event = new(PolygonzkevmetrogRollbackBatches)
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
func (it *PolygonzkevmetrogRollbackBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmetrogRollbackBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmetrogRollbackBatches represents a RollbackBatches event raised by the Polygonzkevmetrog contract.
type PolygonzkevmetrogRollbackBatches struct {
	TargetBatch            uint64
	AccInputHashToRollback [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRollbackBatches is a free log retrieval operation binding the contract event 0x1125aaf62d132d8e2d02005114f8fc360ff204c3105e4f1a700a1340dc55d5b1.
//
// Solidity: event RollbackBatches(uint64 indexed targetBatch, bytes32 accInputHashToRollback)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) FilterRollbackBatches(opts *bind.FilterOpts, targetBatch []uint64) (*PolygonzkevmetrogRollbackBatchesIterator, error) {

	var targetBatchRule []interface{}
	for _, targetBatchItem := range targetBatch {
		targetBatchRule = append(targetBatchRule, targetBatchItem)
	}

	logs, sub, err := _Polygonzkevmetrog.contract.FilterLogs(opts, "RollbackBatches", targetBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmetrogRollbackBatchesIterator{contract: _Polygonzkevmetrog.contract, event: "RollbackBatches", logs: logs, sub: sub}, nil
}

// WatchRollbackBatches is a free log subscription operation binding the contract event 0x1125aaf62d132d8e2d02005114f8fc360ff204c3105e4f1a700a1340dc55d5b1.
//
// Solidity: event RollbackBatches(uint64 indexed targetBatch, bytes32 accInputHashToRollback)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) WatchRollbackBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmetrogRollbackBatches, targetBatch []uint64) (event.Subscription, error) {

	var targetBatchRule []interface{}
	for _, targetBatchItem := range targetBatch {
		targetBatchRule = append(targetBatchRule, targetBatchItem)
	}

	logs, sub, err := _Polygonzkevmetrog.contract.WatchLogs(opts, "RollbackBatches", targetBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmetrogRollbackBatches)
				if err := _Polygonzkevmetrog.contract.UnpackLog(event, "RollbackBatches", log); err != nil {
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

// ParseRollbackBatches is a log parse operation binding the contract event 0x1125aaf62d132d8e2d02005114f8fc360ff204c3105e4f1a700a1340dc55d5b1.
//
// Solidity: event RollbackBatches(uint64 indexed targetBatch, bytes32 accInputHashToRollback)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) ParseRollbackBatches(log types.Log) (*PolygonzkevmetrogRollbackBatches, error) {
	event := new(PolygonzkevmetrogRollbackBatches)
	if err := _Polygonzkevmetrog.contract.UnpackLog(event, "RollbackBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmetrogSequenceBatchesIterator is returned from FilterSequenceBatches and is used to iterate over the raw logs and unpacked data for SequenceBatches events raised by the Polygonzkevmetrog contract.
type PolygonzkevmetrogSequenceBatchesIterator struct {
	Event *PolygonzkevmetrogSequenceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmetrogSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmetrogSequenceBatches)
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
		it.Event = new(PolygonzkevmetrogSequenceBatches)
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
func (it *PolygonzkevmetrogSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmetrogSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmetrogSequenceBatches represents a SequenceBatches event raised by the Polygonzkevmetrog contract.
type PolygonzkevmetrogSequenceBatches struct {
	NumBatch   uint64
	L1InfoRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSequenceBatches is a free log retrieval operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) FilterSequenceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonzkevmetrogSequenceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevmetrog.contract.FilterLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmetrogSequenceBatchesIterator{contract: _Polygonzkevmetrog.contract, event: "SequenceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceBatches is a free log subscription operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) WatchSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmetrogSequenceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevmetrog.contract.WatchLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmetrogSequenceBatches)
				if err := _Polygonzkevmetrog.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
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
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) ParseSequenceBatches(log types.Log) (*PolygonzkevmetrogSequenceBatches, error) {
	event := new(PolygonzkevmetrogSequenceBatches)
	if err := _Polygonzkevmetrog.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmetrogSequenceForceBatchesIterator is returned from FilterSequenceForceBatches and is used to iterate over the raw logs and unpacked data for SequenceForceBatches events raised by the Polygonzkevmetrog contract.
type PolygonzkevmetrogSequenceForceBatchesIterator struct {
	Event *PolygonzkevmetrogSequenceForceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmetrogSequenceForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmetrogSequenceForceBatches)
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
		it.Event = new(PolygonzkevmetrogSequenceForceBatches)
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
func (it *PolygonzkevmetrogSequenceForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmetrogSequenceForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmetrogSequenceForceBatches represents a SequenceForceBatches event raised by the Polygonzkevmetrog contract.
type PolygonzkevmetrogSequenceForceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceForceBatches is a free log retrieval operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) FilterSequenceForceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonzkevmetrogSequenceForceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevmetrog.contract.FilterLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmetrogSequenceForceBatchesIterator{contract: _Polygonzkevmetrog.contract, event: "SequenceForceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceForceBatches is a free log subscription operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) WatchSequenceForceBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmetrogSequenceForceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevmetrog.contract.WatchLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmetrogSequenceForceBatches)
				if err := _Polygonzkevmetrog.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
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
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) ParseSequenceForceBatches(log types.Log) (*PolygonzkevmetrogSequenceForceBatches, error) {
	event := new(PolygonzkevmetrogSequenceForceBatches)
	if err := _Polygonzkevmetrog.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmetrogSetForceBatchAddressIterator is returned from FilterSetForceBatchAddress and is used to iterate over the raw logs and unpacked data for SetForceBatchAddress events raised by the Polygonzkevmetrog contract.
type PolygonzkevmetrogSetForceBatchAddressIterator struct {
	Event *PolygonzkevmetrogSetForceBatchAddress // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmetrogSetForceBatchAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmetrogSetForceBatchAddress)
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
		it.Event = new(PolygonzkevmetrogSetForceBatchAddress)
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
func (it *PolygonzkevmetrogSetForceBatchAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmetrogSetForceBatchAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmetrogSetForceBatchAddress represents a SetForceBatchAddress event raised by the Polygonzkevmetrog contract.
type PolygonzkevmetrogSetForceBatchAddress struct {
	NewForceBatchAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchAddress is a free log retrieval operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) FilterSetForceBatchAddress(opts *bind.FilterOpts) (*PolygonzkevmetrogSetForceBatchAddressIterator, error) {

	logs, sub, err := _Polygonzkevmetrog.contract.FilterLogs(opts, "SetForceBatchAddress")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmetrogSetForceBatchAddressIterator{contract: _Polygonzkevmetrog.contract, event: "SetForceBatchAddress", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchAddress is a free log subscription operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) WatchSetForceBatchAddress(opts *bind.WatchOpts, sink chan<- *PolygonzkevmetrogSetForceBatchAddress) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmetrog.contract.WatchLogs(opts, "SetForceBatchAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmetrogSetForceBatchAddress)
				if err := _Polygonzkevmetrog.contract.UnpackLog(event, "SetForceBatchAddress", log); err != nil {
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
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) ParseSetForceBatchAddress(log types.Log) (*PolygonzkevmetrogSetForceBatchAddress, error) {
	event := new(PolygonzkevmetrogSetForceBatchAddress)
	if err := _Polygonzkevmetrog.contract.UnpackLog(event, "SetForceBatchAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmetrogSetForceBatchTimeoutIterator is returned from FilterSetForceBatchTimeout and is used to iterate over the raw logs and unpacked data for SetForceBatchTimeout events raised by the Polygonzkevmetrog contract.
type PolygonzkevmetrogSetForceBatchTimeoutIterator struct {
	Event *PolygonzkevmetrogSetForceBatchTimeout // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmetrogSetForceBatchTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmetrogSetForceBatchTimeout)
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
		it.Event = new(PolygonzkevmetrogSetForceBatchTimeout)
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
func (it *PolygonzkevmetrogSetForceBatchTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmetrogSetForceBatchTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmetrogSetForceBatchTimeout represents a SetForceBatchTimeout event raised by the Polygonzkevmetrog contract.
type PolygonzkevmetrogSetForceBatchTimeout struct {
	NewforceBatchTimeout uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchTimeout is a free log retrieval operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) FilterSetForceBatchTimeout(opts *bind.FilterOpts) (*PolygonzkevmetrogSetForceBatchTimeoutIterator, error) {

	logs, sub, err := _Polygonzkevmetrog.contract.FilterLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmetrogSetForceBatchTimeoutIterator{contract: _Polygonzkevmetrog.contract, event: "SetForceBatchTimeout", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchTimeout is a free log subscription operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) WatchSetForceBatchTimeout(opts *bind.WatchOpts, sink chan<- *PolygonzkevmetrogSetForceBatchTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmetrog.contract.WatchLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmetrogSetForceBatchTimeout)
				if err := _Polygonzkevmetrog.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
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
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) ParseSetForceBatchTimeout(log types.Log) (*PolygonzkevmetrogSetForceBatchTimeout, error) {
	event := new(PolygonzkevmetrogSetForceBatchTimeout)
	if err := _Polygonzkevmetrog.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmetrogSetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the Polygonzkevmetrog contract.
type PolygonzkevmetrogSetTrustedSequencerIterator struct {
	Event *PolygonzkevmetrogSetTrustedSequencer // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmetrogSetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmetrogSetTrustedSequencer)
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
		it.Event = new(PolygonzkevmetrogSetTrustedSequencer)
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
func (it *PolygonzkevmetrogSetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmetrogSetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmetrogSetTrustedSequencer represents a SetTrustedSequencer event raised by the Polygonzkevmetrog contract.
type PolygonzkevmetrogSetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*PolygonzkevmetrogSetTrustedSequencerIterator, error) {

	logs, sub, err := _Polygonzkevmetrog.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmetrogSetTrustedSequencerIterator{contract: _Polygonzkevmetrog.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *PolygonzkevmetrogSetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmetrog.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmetrogSetTrustedSequencer)
				if err := _Polygonzkevmetrog.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
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
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) ParseSetTrustedSequencer(log types.Log) (*PolygonzkevmetrogSetTrustedSequencer, error) {
	event := new(PolygonzkevmetrogSetTrustedSequencer)
	if err := _Polygonzkevmetrog.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmetrogSetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the Polygonzkevmetrog contract.
type PolygonzkevmetrogSetTrustedSequencerURLIterator struct {
	Event *PolygonzkevmetrogSetTrustedSequencerURL // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmetrogSetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmetrogSetTrustedSequencerURL)
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
		it.Event = new(PolygonzkevmetrogSetTrustedSequencerURL)
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
func (it *PolygonzkevmetrogSetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmetrogSetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmetrogSetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the Polygonzkevmetrog contract.
type PolygonzkevmetrogSetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*PolygonzkevmetrogSetTrustedSequencerURLIterator, error) {

	logs, sub, err := _Polygonzkevmetrog.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmetrogSetTrustedSequencerURLIterator{contract: _Polygonzkevmetrog.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *PolygonzkevmetrogSetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmetrog.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmetrogSetTrustedSequencerURL)
				if err := _Polygonzkevmetrog.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
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
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) ParseSetTrustedSequencerURL(log types.Log) (*PolygonzkevmetrogSetTrustedSequencerURL, error) {
	event := new(PolygonzkevmetrogSetTrustedSequencerURL)
	if err := _Polygonzkevmetrog.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmetrogTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Polygonzkevmetrog contract.
type PolygonzkevmetrogTransferAdminRoleIterator struct {
	Event *PolygonzkevmetrogTransferAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmetrogTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmetrogTransferAdminRole)
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
		it.Event = new(PolygonzkevmetrogTransferAdminRole)
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
func (it *PolygonzkevmetrogTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmetrogTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmetrogTransferAdminRole represents a TransferAdminRole event raised by the Polygonzkevmetrog contract.
type PolygonzkevmetrogTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*PolygonzkevmetrogTransferAdminRoleIterator, error) {

	logs, sub, err := _Polygonzkevmetrog.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmetrogTransferAdminRoleIterator{contract: _Polygonzkevmetrog.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonzkevmetrogTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmetrog.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmetrogTransferAdminRole)
				if err := _Polygonzkevmetrog.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
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
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) ParseTransferAdminRole(log types.Log) (*PolygonzkevmetrogTransferAdminRole, error) {
	event := new(PolygonzkevmetrogTransferAdminRole)
	if err := _Polygonzkevmetrog.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmetrogVerifyBatchesIterator is returned from FilterVerifyBatches and is used to iterate over the raw logs and unpacked data for VerifyBatches events raised by the Polygonzkevmetrog contract.
type PolygonzkevmetrogVerifyBatchesIterator struct {
	Event *PolygonzkevmetrogVerifyBatches // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmetrogVerifyBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmetrogVerifyBatches)
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
		it.Event = new(PolygonzkevmetrogVerifyBatches)
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
func (it *PolygonzkevmetrogVerifyBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmetrogVerifyBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmetrogVerifyBatches represents a VerifyBatches event raised by the Polygonzkevmetrog contract.
type PolygonzkevmetrogVerifyBatches struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatches is a free log retrieval operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) FilterVerifyBatches(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*PolygonzkevmetrogVerifyBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonzkevmetrog.contract.FilterLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmetrogVerifyBatchesIterator{contract: _Polygonzkevmetrog.contract, event: "VerifyBatches", logs: logs, sub: sub}, nil
}

// WatchVerifyBatches is a free log subscription operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) WatchVerifyBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmetrogVerifyBatches, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonzkevmetrog.contract.WatchLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmetrogVerifyBatches)
				if err := _Polygonzkevmetrog.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
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
func (_Polygonzkevmetrog *PolygonzkevmetrogFilterer) ParseVerifyBatches(log types.Log) (*PolygonzkevmetrogVerifyBatches, error) {
	event := new(PolygonzkevmetrogVerifyBatches)
	if err := _Polygonzkevmetrog.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonzkevmexistentetrogpessimistic

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

// PolygonRollupBaseEtrogPessimisticBatchData is an auto generated low-level Go binding around an user-defined struct.
type PolygonRollupBaseEtrogPessimisticBatchData struct {
	Transactions         []byte
	ForcedGlobalExitRoot [32]byte
	ForcedTimestamp      uint64
	ForcedBlockHashL1    [32]byte
}

// PolygonzkevmexistentetrogpessimisticMetaData contains all meta data concerning the Polygonzkevmexistentetrogpessimistic contract.
var PolygonzkevmexistentetrogpessimisticMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractPolygonRollupManagerPessimistic\",\"name\":\"_rollupManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalAccInputHashDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"forceBatchNum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"}],\"name\":\"ForceBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"}],\"name\":\"InitialSequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"targetBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accInputHashToRollback\",\"type\":\"bytes32\"}],\"name\":\"RollbackBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"l1InfoRoot\",\"type\":\"bytes32\"}],\"name\":\"SequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newForceBatchAddress\",\"type\":\"address\"}],\"name\":\"SetForceBatchAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"SetForceBatchTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"}],\"name\":\"UpdateEtrogSequence\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatches\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GLOBAL_EXIT_ROOT_MANAGER_L2\",\"outputs\":[{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_LIST_LEN_LEN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_DATA_LEN_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_EFFECTIVE_PERCENTAGE\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SET_UP_ETROG_TX\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_R\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_S\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_V\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TIMESTAMP_RANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculatePolPerForceBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"polAmount\",\"type\":\"uint256\"}],\"name\":\"forceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"}],\"name\":\"generateInitializeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trustedSequencer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_trustedSequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_lastAccInputHash\",\"type\":\"bytes32\"}],\"name\":\"initializeUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"onVerifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"targetBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"accInputHashToRollback\",\"type\":\"bytes32\"}],\"name\":\"rollbackBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManagerPessimistic\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrogPessimistic.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"uint32\",\"name\":\"l1InfoTreeLeafCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"maxSequenceTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"expectedFinalAccInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"}],\"name\":\"sequenceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrogPessimistic.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"}],\"name\":\"sequenceForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newForceBatchAddress\",\"type\":\"address\"}],\"name\":\"setForceBatchAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"setForceBatchTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x61010060405234801562000011575f80fd5b5060405162004c8438038062004c8483398101604081905262000034916200014e565b6001600160a01b0380851660a05280841660805280831660c052811660e05283838383838383836200006562000077565b505050505050505050505050620001b3565b5f54610100900460ff1615620000e35760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116101562000134575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200014b575f80fd5b50565b5f805f806080858703121562000162575f80fd5b84516200016f8162000136565b6020860151909450620001828162000136565b6040860151909350620001958162000136565b6060860151909250620001a88162000136565b939692955090935050565b60805160a05160c05160e0516149d2620002b25f395f8181610512015281816109a401528181610b1001528181610c8f01528181610fd901528181611212015281816113920152818161170f01528181611baa01528181611ff9015281816120ee01528181612ca401528181612d1d01528181612d3f01528181612e540152818161306a015261312f01525f8181610691015281816122b8015281816123c0015281816127950152818161351c01526135f101525f818161076801528181610e4e0152818161158601528181612843015261327701525f818161079a015281816108570152818161204201528181612deb015261324c01526149d25ff3fe608060405234801561000f575f80fd5b50600436106102f8575f3560e01c80637125702211610192578063b910e0f9116100e8578063d7bc90ff11610093578063eaeb077b1161006e578063eaeb077b146107d0578063f35dda47146107e3578063f851a440146107eb575f80fd5b8063d7bc90ff1461078a578063e46761c414610795578063e7a7ed02146107bc575f80fd5b8063c89e42df116100c3578063c89e42df14610730578063cfa8ed4714610743578063d02103ca14610763575f80fd5b8063b910e0f9146106ed578063c754c7ed14610700578063c7fffd4b14610728575f80fd5b80639f26f84011610148578063ada8f91911610123578063ada8f919146106c6578063af7f3e02146106d9578063b0afe154146106e1575f80fd5b80639f26f84014610679578063a3c573eb1461068c578063a652f26c146106b3575f80fd5b80638c3d7301116101785780638c3d73011461064357806391cafe321461064b5780639e0018771461065e575f80fd5b806371257022146105f45780637a5460c514610607575f80fd5b806342308fab116102525780635d6717a5116101fd5780636b8616ce116101d85780636b8616ce146105b95780636e05d2cd146105d85780636ff512cc146105e1575f80fd5b80635d6717a51461058b578063669adece1461059e578063676870d2146105b1575f80fd5b80634e4877061161022d5780634e4877061461053457806352bdeb6d14610547578063542028d514610583575f80fd5b806342308fab146104cc57806345605267146104d457806349b7b8021461050d575f80fd5b806326782247116102b25780633c351e101161028d5780633c351e10146104175780633cbc795b1461043757806340b5de6c14610474575f80fd5b8063267822471461039d5780632c111c06146103e257806332c2d15314610402575f80fd5b806305835f37116102e257806305835f3714610332578063107bf28c1461037b57806311e892d414610383575f80fd5b8062d0295d146102fc5780630350896314610317575b5f80fd5b610304610810565b6040519081526020015b60405180910390f35b61031f602081565b60405161ffff909116815260200161030e565b61036e6040518060400160405280600881526020017f80808401c9c3809400000000000000000000000000000000000000000000000081525081565b60405161030e9190613b39565b61036e610916565b61038b60f981565b60405160ff909116815260200161030e565b6001546103bd9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161030e565b6008546103bd9073ffffffffffffffffffffffffffffffffffffffff1681565b610415610410366004613b8b565b6109a2565b005b6009546103bd9073ffffffffffffffffffffffffffffffffffffffff1681565b60095461045f9074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff909116815260200161030e565b61049b7fff0000000000000000000000000000000000000000000000000000000000000081565b6040517fff00000000000000000000000000000000000000000000000000000000000000909116815260200161030e565b610304602481565b6007546104f49068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff909116815260200161030e565b6103bd7f000000000000000000000000000000000000000000000000000000000000000081565b610415610542366004613bca565b610a71565b61036e6040518060400160405280600281526020017f80b800000000000000000000000000000000000000000000000000000000000081525081565b61036e610c80565b610415610599366004613cf8565b610c8d565b6104156105ac366004613d83565b611210565b61031f601f81565b6103046105c7366004613bca565b60066020525f908152604090205481565b61030460055481565b6104156105ef366004613dad565b6112c7565b610415610602366004613dd9565b611390565b61036e6040518060400160405280600281526020017f80b900000000000000000000000000000000000000000000000000000000000081525081565b610415611944565b610415610659366004613dad565b611a16565b6103bd73a40d5f56745a118d0906a34e69aec8c0db1cb8fa81565b610415610687366004613ec8565b611b2e565b6103bd7f000000000000000000000000000000000000000000000000000000000000000081565b61036e6106c1366004613f07565b6121ba565b6104156106d4366004613dad565b612598565b61036e612661565b6103046405ca1ab1e081565b6104156106fb366004613f78565b61267d565b6007546104f490700100000000000000000000000000000000900467ffffffffffffffff1681565b61038b60e481565b61041561073e366004613ff8565b612f61565b6002546103bd9073ffffffffffffffffffffffffffffffffffffffff1681565b6103bd7f000000000000000000000000000000000000000000000000000000000000000081565b610304635ca1ab1e81565b6103bd7f000000000000000000000000000000000000000000000000000000000000000081565b6007546104f49067ffffffffffffffff1681565b6104156107de36600461402a565b612ff3565b61038b601b81565b5f546103bd9062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f90819073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa15801561089c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108c0919061409b565b6007549091505f906108ea9067ffffffffffffffff680100000000000000008204811691166140df565b67ffffffffffffffff169050805f03610905575f9250505090565b61090f8183614107565b9250505090565b600480546109239061413f565b80601f016020809104026020016040519081016040528092919081815260200182805461094f9061413f565b801561099a5780601f106109715761010080835404028352916020019161099a565b820191905f5260205f20905b81548152906001019060200180831161097d57829003601f168201915b505050505081565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314610a11576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff168367ffffffffffffffff167f9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f596684604051610a6491815260200190565b60405180910390a3505050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314610ac7576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff82161115610b0e576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166315064c966040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b77573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b9b9190614190565b610bfc5760075467ffffffffffffffff700100000000000000000000000000000000909104811690821610610bfc576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600780547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527fa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b906020015b60405180910390a150565b600380546109239061413f565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314610cfc576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f54610100900460ff1615808015610d1a57505f54600160ff909116105b80610d335750303b158015610d3357505f5460ff166001145b610dc4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610e20575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b5f6040518060a001604052806062815260200161493b6062913990505f818051906020012090505f4290505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610eb5573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ed9919061409b565b90505f868483858d610eec6001436141af565b60408051602081019790975286019490945260608086019390935260c09190911b7fffffffffffffffff000000000000000000000000000000000000000000000000166080850152901b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016608883015240609c82015260bc01604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001815290829052805160209091012060058190557f9a908e73000000000000000000000000000000000000000000000000000000008252600160048301526024820181905291505f907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303815f875af1158015611034573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061105891906141c8565b90508b5f60026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508a60025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555089600390816110e99190614228565b5060046110f68a82614228565b508b60085f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062069780600760106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fd2c80353fc15ef62c6affc7cd6b7ab5b42c43290c50be3372e55ae552cecd19c8187858e6040516111989493929190614340565b60405180910390a15050505050508015611208575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16331461127f576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600581905560405181815267ffffffffffffffff8316907f1125aaf62d132d8e2d02005114f8fc360ff204c3105e4f1a700a1340dc55d5b19060200160405180910390a25050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff16331461131d576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc090602001610c75565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1633146113ff576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f54610100900460ff161580801561141d57505f54600160ff909116105b806114365750303b15801561143657505f5460ff166001145b6114c2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610dbb565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561151e575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b5f611528856134b9565b6009549091505f9061157290889073ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900463ffffffff16856121ba565b90505f818051906020012090505f4290505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115ed573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611611919061409b565b90505f808483858f6116246001436141af565b60408051602081019790975286019490945260608086019390935260c09190911b7fffffffffffffffff000000000000000000000000000000000000000000000000166080850152901b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016608883015240609c82015260bc01604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001815290829052805160209091012060058190557f9a908e73000000000000000000000000000000000000000000000000000000008252600160048301526024820181905291507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303815f875af115801561176a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061178e91906141c8565b508c5f60026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508b60025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550886003908161181e9190614228565b50600461182b8982614228565b508c60085f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062069780600760106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507f060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f85838e6040516118cb9392919061438f565b60405180910390a1505050505050801561193b575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314611995576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001545f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff1673ffffffffffffffffffffffffffffffffffffffff9092166201000081029290921790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e9060200160405180910390a1565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314611a6c576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60085473ffffffffffffffffffffffffffffffffffffffff16611abb576040517fc89374d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb90602001610c75565b60085473ffffffffffffffffffffffffffffffffffffffff168015801590611b6c575073ffffffffffffffffffffffffffffffffffffffff81163314155b15611ba3576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b4262093a807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166330c27dde6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611c11573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611c3591906141c8565b611c3f91906143cd565b67ffffffffffffffff161115611c81576040517f3d49ed4c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815f819003611cbc576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8811115611cf8576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075467ffffffffffffffff80821691611d20918491680100000000000000009004166143ee565b1115611d58576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6007546005546801000000000000000090910467ffffffffffffffff16905f5b83811015611ff3575f878783818110611d9357611d93614401565b9050602002810190611da5919061442e565b611dae9061446a565b905083611dba816144f0565b825180516020918201208185015160408087015160608801519151959a509295505f94611e26948794929101938452602084019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166040830152604882015260680190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff89165f90815260069093529120549091508114611eae576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff86165f90815260066020526040812055611ed26001886141af565b8403611f415742600760109054906101000a900467ffffffffffffffff168460400151611eff91906143cd565b67ffffffffffffffff161115611f41576040517fc44a082100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208381015160408086015160608088015183519586018b90529285018790528481019390935260c01b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808401523390911b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088830152609c82015260bc016040516020818303038152906040528051906020012094505050508080611feb90614516565b915050611d78565b506120697f000000000000000000000000000000000000000000000000000000000000000084612021610810565b61202b919061454d565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016919061371b565b60058190556007805467ffffffffffffffff841668010000000000000000027fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff9091161790556040517f9a908e730000000000000000000000000000000000000000000000000000000081525f9073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690639a908e739061213a908790869060040167ffffffffffffffff929092168252602082015260400190565b6020604051808303815f875af1158015612156573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061217a91906141c8565b60405190915067ffffffffffffffff8216907f648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4905f90a250505050505050565b60605f85858573a40d5f56745a118d0906a34e69aec8c0db1cb8fa5f876040516024016121ec96959493929190614564565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff811bff70000000000000000000000000000000000000000000000000000000017905283519091506060905f0361233c5760f9601f835161228091906145c6565b6040518060400160405280600881526020017f80808401c9c380940000000000000000000000000000000000000000000000008152507f00000000000000000000000000000000000000000000000000000000000000006040518060400160405280600281526020017f80b800000000000000000000000000000000000000000000000000000000000081525060e48760405160200161232697969594939291906145e1565b6040516020818303038152906040529050612440565b815161ffff1015612379576040517f248b8f8200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815160f96123886020836145c6565b6040518060400160405280600881526020017f80808401c9c380940000000000000000000000000000000000000000000000008152507f00000000000000000000000000000000000000000000000000000000000000006040518060400160405280600281526020017f80b9000000000000000000000000000000000000000000000000000000000000815250858860405160200161242d97969594939291906146c3565b6040516020818303038152906040529150505b8051602080830191909120604080515f80825293810180835292909252601b908201526405ca1ab1e06060820152635ca1ab1e608082015260019060a0016020604051602081039080840390855afa15801561249e573d5f803e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116612516576040517fcd16196600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040515f9061255b9084906405ca1ab1e090635ca1ab1e90601b907fff00000000000000000000000000000000000000000000000000000000000000906020016147a5565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529450505050505b949350505050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff1633146125ee576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001610c75565b6040518060a001604052806062815260200161493b6062913981565b60025473ffffffffffffffffffffffffffffffffffffffff1633146126ce576040517f11e7be1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b845f819003612709576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8811115612745576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6127506024426143ee565b8467ffffffffffffffff161115612793576040517f0a00feb300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379e2cf976040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156127f8575f80fd5b505af115801561280a573d5f803e3d5ffd5b50506040517fef4eeb3500000000000000000000000000000000000000000000000000000000815263ffffffff881660048201525f92507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16915063ef4eeb3590602401602060405180830381865afa15801561289e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906128c2919061409b565b9050806128fb576040517fa60721e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6007546005546801000000000000000090910467ffffffffffffffff1690815f5b85811015612c17575f8c8c8381811061293757612937614401565b9050602002810190612949919061442e565b6129529061446a565b8051805160209091012060408201519192509067ffffffffffffffff1615612b32578561297e816144f0565b9650505f818360200151846040015185606001516040516020016129e09493929190938452602084019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166040830152604882015260680190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff8a165f90815260069093529120549091508114612a68576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208381015160408086015160608088015183519586018c90529285018790528481019390935260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166080840152908c901b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088830152609c82015260bc0160405160208183030381529060405280519060200120955060065f8867ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f905550612c02565b8151516201d4c01015612b71576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160208101879052908101829052606080820189905260c08d901b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808301528a901b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660888201525f609c82015260bc016040516020818303038152906040528051906020012094505b50508080612c0f90614516565b91505061291c565b5060075467ffffffffffffffff9081169084161115612c62576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60058290558467ffffffffffffffff84811690831614612d17575f612c8783866140df565b9050612c9d67ffffffffffffffff8216836141af565b9150612cd67f00000000000000000000000000000000000000000000000000000000000000008267ffffffffffffffff16612021610810565b50600780547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8716021790555b612e13337f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663477fa2706040518163ffffffff1660e01b8152600401602060405180830381865afa158015612da6573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612dca919061409b565b612dd4919061454d565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169291906137f4565b6040517f9a908e7300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff87166004820152602481018490525f907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303815f875af1158015612eaf573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612ed391906141c8565b9050888414612f0e576040517fda5bceb900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff167f3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e76687604051612f4a91815260200190565b60405180910390a250505050505050505050505050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314612fb7576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6003612fc38282614228565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b2081604051610c759190613b39565b60085473ffffffffffffffffffffffffffffffffffffffff168015801590613031575073ffffffffffffffffffffffffffffffffffffffff81163314155b15613068576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166315064c966040518163ffffffff1660e01b8152600401602060405180830381865afa1580156130d1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130f59190614190565b1561312c576040517f39258d1800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663604691696040518163ffffffff1660e01b8152600401602060405180830381865afa158015613196573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906131ba919061409b565b9050828111156131f6576040517f2354600f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611388841115613232576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61327473ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163330846137f4565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156132de573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613302919061409b565b6007805491925067ffffffffffffffff909116905f613320836144f0565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550508585604051613357929190614800565b604051908190039020814261336d6001436141af565b60408051602081019590955284019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166060830152406068820152608801604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012060075467ffffffffffffffff165f9081526006909352912055323303613463576007546040805183815233602082015260608183018190525f90820152905167ffffffffffffffff909216917ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319181900360800190a2611208565b60075460405167ffffffffffffffff909116907ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931906134a990849033908b908b9061480f565b60405180910390a2505050505050565b606073ffffffffffffffffffffffffffffffffffffffff821615613716576040517fc00f14ab00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063c00f14ab906024015f60405180830381865afa158015613560573d5f803e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526135a5919081019061487f565b6040517f318aee3d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301529192505f9182917f00000000000000000000000000000000000000000000000000000000000000009091169063318aee3d906024016040805180830381865afa158015613637573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061365b91906148f1565b915091508163ffffffff165f146136d2576009805463ffffffff841674010000000000000000000000000000000000000000027fffffffffffffffff00000000000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff841617179055613713565b600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff86161790555b50505b919050565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526137ef9084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152613858565b505050565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526138529085907f23b872dd000000000000000000000000000000000000000000000000000000009060840161376d565b50505050565b5f6138b9826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166139639092919063ffffffff16565b8051909150156137ef57808060200190518101906138d79190614190565b6137ef576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610dbb565b606061259084845f85855f808673ffffffffffffffffffffffffffffffffffffffff1685876040516139959190614929565b5f6040518083038185875af1925050503d805f81146139cf576040519150601f19603f3d011682016040523d82523d5f602084013e6139d4565b606091505b50915091506139e5878383876139f0565b979650505050505050565b60608315613a855782515f03613a7e5773ffffffffffffffffffffffffffffffffffffffff85163b613a7e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610dbb565b5081612590565b6125908383815115613a9a5781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dbb9190613b39565b5f5b83811015613ae8578181015183820152602001613ad0565b50505f910152565b5f8151808452613b07816020860160208601613ace565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081525f613b4b6020830184613af0565b9392505050565b67ffffffffffffffff81168114613b67575f80fd5b50565b73ffffffffffffffffffffffffffffffffffffffff81168114613b67575f80fd5b5f805f60608486031215613b9d575f80fd5b8335613ba881613b52565b9250602084013591506040840135613bbf81613b6a565b809150509250925092565b5f60208284031215613bda575f80fd5b8135613b4b81613b52565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613c5957613c59613be5565b604052919050565b5f67ffffffffffffffff821115613c7a57613c7a613be5565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b5f82601f830112613cb5575f80fd5b8135613cc8613cc382613c61565b613c12565b818152846020838601011115613cdc575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f805f60a08688031215613d0c575f80fd5b8535613d1781613b6a565b94506020860135613d2781613b6a565b9350604086013567ffffffffffffffff80821115613d43575f80fd5b613d4f89838a01613ca6565b94506060880135915080821115613d64575f80fd5b50613d7188828901613ca6565b95989497509295608001359392505050565b5f8060408385031215613d94575f80fd5b8235613d9f81613b52565b946020939093013593505050565b5f60208284031215613dbd575f80fd5b8135613b4b81613b6a565b63ffffffff81168114613b67575f80fd5b5f805f805f8060c08789031215613dee575f80fd5b8635613df981613b6a565b95506020870135613e0981613b6a565b94506040870135613e1981613dc8565b93506060870135613e2981613b6a565b9250608087013567ffffffffffffffff80821115613e45575f80fd5b613e518a838b01613ca6565b935060a0890135915080821115613e66575f80fd5b50613e7389828a01613ca6565b9150509295509295509295565b5f8083601f840112613e90575f80fd5b50813567ffffffffffffffff811115613ea7575f80fd5b6020830191508360208260051b8501011115613ec1575f80fd5b9250929050565b5f8060208385031215613ed9575f80fd5b823567ffffffffffffffff811115613eef575f80fd5b613efb85828601613e80565b90969095509350505050565b5f805f8060808587031215613f1a575f80fd5b8435613f2581613dc8565b93506020850135613f3581613b6a565b92506040850135613f4581613dc8565b9150606085013567ffffffffffffffff811115613f60575f80fd5b613f6c87828801613ca6565b91505092959194509250565b5f805f805f8060a08789031215613f8d575f80fd5b863567ffffffffffffffff811115613fa3575f80fd5b613faf89828a01613e80565b9097509550506020870135613fc381613dc8565b93506040870135613fd381613b52565b9250606087013591506080870135613fea81613b6a565b809150509295509295509295565b5f60208284031215614008575f80fd5b813567ffffffffffffffff81111561401e575f80fd5b61259084828501613ca6565b5f805f6040848603121561403c575f80fd5b833567ffffffffffffffff80821115614053575f80fd5b818601915086601f830112614066575f80fd5b813581811115614074575f80fd5b876020828501011115614085575f80fd5b6020928301989097509590910135949350505050565b5f602082840312156140ab575f80fd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b67ffffffffffffffff828116828216039080821115614100576141006140b2565b5092915050565b5f8261413a577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b600181811c9082168061415357607f821691505b60208210810361418a577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b5f602082840312156141a0575f80fd5b81518015158114613b4b575f80fd5b818103818111156141c2576141c26140b2565b92915050565b5f602082840312156141d8575f80fd5b8151613b4b81613b52565b601f8211156137ef575f81815260208120601f850160051c810160208610156142095750805b601f850160051c820191505b8181101561120857828155600101614215565b815167ffffffffffffffff81111561424257614242613be5565b61425681614250845461413f565b846141e3565b602080601f8311600181146142a8575f84156142725750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555611208565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b828110156142f4578886015182559484019460019091019084016142d5565b508582101561433057878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b67ffffffffffffffff85168152608060208201525f6143626080830186613af0565b905083604083015273ffffffffffffffffffffffffffffffffffffffff8316606083015295945050505050565b606081525f6143a16060830186613af0565b905083602083015273ffffffffffffffffffffffffffffffffffffffff83166040830152949350505050565b67ffffffffffffffff818116838216019080821115614100576141006140b2565b808201808211156141c2576141c26140b2565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f82357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81833603018112614460575f80fd5b9190910192915050565b5f6080823603121561447a575f80fd5b6040516080810167ffffffffffffffff828210818311171561449e5761449e613be5565b8160405284359150808211156144b2575f80fd5b506144bf36828601613ca6565b8252506020830135602082015260408301356144da81613b52565b6040820152606092830135928101929092525090565b5f67ffffffffffffffff80831681810361450c5761450c6140b2565b6001019392505050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614546576145466140b2565b5060010190565b80820281158282048414176141c2576141c26140b2565b5f63ffffffff808916835273ffffffffffffffffffffffffffffffffffffffff8089166020850152818816604085015280871660608501528086166080850152505060c060a08301526145ba60c0830184613af0565b98975050505050505050565b61ffff818116838216019080821115614100576141006140b2565b5f7fff00000000000000000000000000000000000000000000000000000000000000808a60f81b1683527fffff0000000000000000000000000000000000000000000000000000000000008960f01b1660018401528751614649816003860160208c01613ace565b80840190507fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008860601b166003820152865161468c816017840160208b01613ace565b808201915050818660f81b166017820152845191506146b2826018830160208801613ace565b016018019998505050505050505050565b7fff000000000000000000000000000000000000000000000000000000000000008860f81b1681525f7fffff000000000000000000000000000000000000000000000000000000000000808960f01b166001840152875161472b816003860160208c01613ace565b80840190507fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008860601b166003820152865161476e816017840160208b01613ace565b808201915050818660f01b16601782015284519150614794826019830160208801613ace565b016019019998505050505050505050565b5f86516147b6818460208b01613ace565b9190910194855250602084019290925260f81b7fff000000000000000000000000000000000000000000000000000000000000009081166040840152166041820152604201919050565b818382375f9101908152919050565b84815273ffffffffffffffffffffffffffffffffffffffff8416602082015260606040820152816060820152818360808301375f818301608090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01601019392505050565b5f6020828403121561488f575f80fd5b815167ffffffffffffffff8111156148a5575f80fd5b8201601f810184136148b5575f80fd5b80516148c3613cc382613c61565b8181528560208385010111156148d7575f80fd5b6148e8826020830160208601613ace565b95945050505050565b5f8060408385031215614902575f80fd5b825161490d81613dc8565b602084015190925061491e81613b6a565b809150509250929050565b5f8251614460818460208701613ace56fedf2a8080944d5cf5032b2a844602278b01199ed191a86c93ff8080821092808000000000000000000000000000000000000000000000000000000005ca1ab1e000000000000000000000000000000000000000000000000000000005ca1ab1e01bffa2646970667358221220a56ec966497516d11dcc18b0d9cd5461b524bff43a08bbe52084b505e51ea46564736f6c63430008140033",
}

// PolygonzkevmexistentetrogpessimisticABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonzkevmexistentetrogpessimisticMetaData.ABI instead.
var PolygonzkevmexistentetrogpessimisticABI = PolygonzkevmexistentetrogpessimisticMetaData.ABI

// PolygonzkevmexistentetrogpessimisticBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygonzkevmexistentetrogpessimisticMetaData.Bin instead.
var PolygonzkevmexistentetrogpessimisticBin = PolygonzkevmexistentetrogpessimisticMetaData.Bin

// DeployPolygonzkevmexistentetrogpessimistic deploys a new Ethereum contract, binding an instance of Polygonzkevmexistentetrogpessimistic to it.
func DeployPolygonzkevmexistentetrogpessimistic(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address, _rollupManager common.Address) (common.Address, *types.Transaction, *Polygonzkevmexistentetrogpessimistic, error) {
	parsed, err := PolygonzkevmexistentetrogpessimisticMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygonzkevmexistentetrogpessimisticBin), backend, _globalExitRootManager, _pol, _bridgeAddress, _rollupManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonzkevmexistentetrogpessimistic{PolygonzkevmexistentetrogpessimisticCaller: PolygonzkevmexistentetrogpessimisticCaller{contract: contract}, PolygonzkevmexistentetrogpessimisticTransactor: PolygonzkevmexistentetrogpessimisticTransactor{contract: contract}, PolygonzkevmexistentetrogpessimisticFilterer: PolygonzkevmexistentetrogpessimisticFilterer{contract: contract}}, nil
}

// Polygonzkevmexistentetrogpessimistic is an auto generated Go binding around an Ethereum contract.
type Polygonzkevmexistentetrogpessimistic struct {
	PolygonzkevmexistentetrogpessimisticCaller     // Read-only binding to the contract
	PolygonzkevmexistentetrogpessimisticTransactor // Write-only binding to the contract
	PolygonzkevmexistentetrogpessimisticFilterer   // Log filterer for contract events
}

// PolygonzkevmexistentetrogpessimisticCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonzkevmexistentetrogpessimisticCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmexistentetrogpessimisticTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonzkevmexistentetrogpessimisticTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmexistentetrogpessimisticFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonzkevmexistentetrogpessimisticFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmexistentetrogpessimisticSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonzkevmexistentetrogpessimisticSession struct {
	Contract     *Polygonzkevmexistentetrogpessimistic // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                         // Call options to use throughout this session
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// PolygonzkevmexistentetrogpessimisticCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonzkevmexistentetrogpessimisticCallerSession struct {
	Contract *PolygonzkevmexistentetrogpessimisticCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                               // Call options to use throughout this session
}

// PolygonzkevmexistentetrogpessimisticTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonzkevmexistentetrogpessimisticTransactorSession struct {
	Contract     *PolygonzkevmexistentetrogpessimisticTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                               // Transaction auth options to use throughout this session
}

// PolygonzkevmexistentetrogpessimisticRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonzkevmexistentetrogpessimisticRaw struct {
	Contract *Polygonzkevmexistentetrogpessimistic // Generic contract binding to access the raw methods on
}

// PolygonzkevmexistentetrogpessimisticCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonzkevmexistentetrogpessimisticCallerRaw struct {
	Contract *PolygonzkevmexistentetrogpessimisticCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonzkevmexistentetrogpessimisticTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonzkevmexistentetrogpessimisticTransactorRaw struct {
	Contract *PolygonzkevmexistentetrogpessimisticTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonzkevmexistentetrogpessimistic creates a new instance of Polygonzkevmexistentetrogpessimistic, bound to a specific deployed contract.
func NewPolygonzkevmexistentetrogpessimistic(address common.Address, backend bind.ContractBackend) (*Polygonzkevmexistentetrogpessimistic, error) {
	contract, err := bindPolygonzkevmexistentetrogpessimistic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmexistentetrogpessimistic{PolygonzkevmexistentetrogpessimisticCaller: PolygonzkevmexistentetrogpessimisticCaller{contract: contract}, PolygonzkevmexistentetrogpessimisticTransactor: PolygonzkevmexistentetrogpessimisticTransactor{contract: contract}, PolygonzkevmexistentetrogpessimisticFilterer: PolygonzkevmexistentetrogpessimisticFilterer{contract: contract}}, nil
}

// NewPolygonzkevmexistentetrogpessimisticCaller creates a new read-only instance of Polygonzkevmexistentetrogpessimistic, bound to a specific deployed contract.
func NewPolygonzkevmexistentetrogpessimisticCaller(address common.Address, caller bind.ContractCaller) (*PolygonzkevmexistentetrogpessimisticCaller, error) {
	contract, err := bindPolygonzkevmexistentetrogpessimistic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogpessimisticCaller{contract: contract}, nil
}

// NewPolygonzkevmexistentetrogpessimisticTransactor creates a new write-only instance of Polygonzkevmexistentetrogpessimistic, bound to a specific deployed contract.
func NewPolygonzkevmexistentetrogpessimisticTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonzkevmexistentetrogpessimisticTransactor, error) {
	contract, err := bindPolygonzkevmexistentetrogpessimistic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogpessimisticTransactor{contract: contract}, nil
}

// NewPolygonzkevmexistentetrogpessimisticFilterer creates a new log filterer instance of Polygonzkevmexistentetrogpessimistic, bound to a specific deployed contract.
func NewPolygonzkevmexistentetrogpessimisticFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonzkevmexistentetrogpessimisticFilterer, error) {
	contract, err := bindPolygonzkevmexistentetrogpessimistic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogpessimisticFilterer{contract: contract}, nil
}

// bindPolygonzkevmexistentetrogpessimistic binds a generic wrapper to an already deployed contract.
func bindPolygonzkevmexistentetrogpessimistic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonzkevmexistentetrogpessimisticMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmexistentetrogpessimistic.Contract.PolygonzkevmexistentetrogpessimisticCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.PolygonzkevmexistentetrogpessimisticTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.PolygonzkevmexistentetrogpessimisticTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmexistentetrogpessimistic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.contract.Transact(opts, method, params...)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) GLOBALEXITROOTMANAGERL2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "GLOBAL_EXIT_ROOT_MANAGER_L2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.GLOBALEXITROOTMANAGERL2(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.GLOBALEXITROOTMANAGERL2(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) INITIALIZETXBRIDGELISTLENLEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_LIST_LEN_LEN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.INITIALIZETXBRIDGELISTLENLEN(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.INITIALIZETXBRIDGELISTLENLEN(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) INITIALIZETXBRIDGEPARAMS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.INITIALIZETXBRIDGEPARAMS(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.INITIALIZETXBRIDGEPARAMS(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) INITIALIZETXCONSTANTBYTES(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.INITIALIZETXCONSTANTBYTES(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.INITIALIZETXCONSTANTBYTES(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) INITIALIZETXCONSTANTBYTESEMPTYMETADATA(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) INITIALIZETXDATALENEMPTYMETADATA(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "INITIALIZE_TX_DATA_LEN_EMPTY_METADATA")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) INITIALIZETXEFFECTIVEPERCENTAGE(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "INITIALIZE_TX_EFFECTIVE_PERCENTAGE")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// SETUPETROGTX is a free data retrieval call binding the contract method 0xaf7f3e02.
//
// Solidity: function SET_UP_ETROG_TX() view returns(bytes)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) SETUPETROGTX(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "SET_UP_ETROG_TX")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SETUPETROGTX is a free data retrieval call binding the contract method 0xaf7f3e02.
//
// Solidity: function SET_UP_ETROG_TX() view returns(bytes)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) SETUPETROGTX() ([]byte, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.SETUPETROGTX(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// SETUPETROGTX is a free data retrieval call binding the contract method 0xaf7f3e02.
//
// Solidity: function SET_UP_ETROG_TX() view returns(bytes)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) SETUPETROGTX() ([]byte, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.SETUPETROGTX(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) SIGNATUREINITIALIZETXR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_R")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.SIGNATUREINITIALIZETXR(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.SIGNATUREINITIALIZETXR(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) SIGNATUREINITIALIZETXS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_S")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.SIGNATUREINITIALIZETXS(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.SIGNATUREINITIALIZETXS(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) SIGNATUREINITIALIZETXV(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_V")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.SIGNATUREINITIALIZETXV(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.SIGNATUREINITIALIZETXV(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) TIMESTAMPRANGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "TIMESTAMP_RANGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) TIMESTAMPRANGE() (*big.Int, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.TIMESTAMPRANGE(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) TIMESTAMPRANGE() (*big.Int, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.TIMESTAMPRANGE(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) Admin() (common.Address, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.Admin(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) Admin() (common.Address, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.Admin(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) BridgeAddress() (common.Address, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.BridgeAddress(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) BridgeAddress() (common.Address, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.BridgeAddress(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) CalculatePolPerForceBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "calculatePolPerForceBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) CalculatePolPerForceBatch() (*big.Int, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.CalculatePolPerForceBatch(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) CalculatePolPerForceBatch() (*big.Int, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.CalculatePolPerForceBatch(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) ForceBatchAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "forceBatchAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) ForceBatchAddress() (common.Address, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.ForceBatchAddress(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) ForceBatchAddress() (common.Address, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.ForceBatchAddress(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.ForceBatchTimeout(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.ForceBatchTimeout(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.ForcedBatches(&_Polygonzkevmexistentetrogpessimistic.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.ForcedBatches(&_Polygonzkevmexistentetrogpessimistic.CallOpts, arg0)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) GasTokenAddress() (common.Address, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.GasTokenAddress(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) GasTokenAddress() (common.Address, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.GasTokenAddress(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) GasTokenNetwork() (uint32, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.GasTokenNetwork(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) GasTokenNetwork() (uint32, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.GasTokenNetwork(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) GenerateInitializeTransaction(opts *bind.CallOpts, networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "generateInitializeTransaction", networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) GenerateInitializeTransaction(networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.GenerateInitializeTransaction(&_Polygonzkevmexistentetrogpessimistic.CallOpts, networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) GenerateInitializeTransaction(networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.GenerateInitializeTransaction(&_Polygonzkevmexistentetrogpessimistic.CallOpts, networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.GlobalExitRootManager(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.GlobalExitRootManager(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) LastAccInputHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "lastAccInputHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) LastAccInputHash() ([32]byte, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.LastAccInputHash(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) LastAccInputHash() ([32]byte, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.LastAccInputHash(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) LastForceBatch() (uint64, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.LastForceBatch(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) LastForceBatch() (uint64, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.LastForceBatch(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.LastForceBatchSequenced(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.LastForceBatchSequenced(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) NetworkName() (string, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.NetworkName(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) NetworkName() (string, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.NetworkName(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) PendingAdmin() (common.Address, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.PendingAdmin(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) PendingAdmin() (common.Address, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.PendingAdmin(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) Pol() (common.Address, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.Pol(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) Pol() (common.Address, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.Pol(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) RollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "rollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) RollupManager() (common.Address, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.RollupManager(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) RollupManager() (common.Address, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.RollupManager(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) TrustedSequencer() (common.Address, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.TrustedSequencer(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) TrustedSequencer() (common.Address, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.TrustedSequencer(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCaller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrogpessimistic.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) TrustedSequencerURL() (string, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.TrustedSequencerURL(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticCallerSession) TrustedSequencerURL() (string, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.TrustedSequencerURL(&_Polygonzkevmexistentetrogpessimistic.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.AcceptAdminRole(&_Polygonzkevmexistentetrogpessimistic.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.AcceptAdminRole(&_Polygonzkevmexistentetrogpessimistic.TransactOpts)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactor) ForceBatch(opts *bind.TransactOpts, transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.contract.Transact(opts, "forceBatch", transactions, polAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) ForceBatch(transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.ForceBatch(&_Polygonzkevmexistentetrogpessimistic.TransactOpts, transactions, polAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactorSession) ForceBatch(transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.ForceBatch(&_Polygonzkevmexistentetrogpessimistic.TransactOpts, transactions, polAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.contract.Transact(opts, "initialize", _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.Initialize(&_Polygonzkevmexistentetrogpessimistic.TransactOpts, _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactorSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.Initialize(&_Polygonzkevmexistentetrogpessimistic.TransactOpts, _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// InitializeUpgrade is a paid mutator transaction binding the contract method 0x5d6717a5.
//
// Solidity: function initializeUpgrade(address _admin, address _trustedSequencer, string _trustedSequencerURL, string _networkName, bytes32 _lastAccInputHash) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactor) InitializeUpgrade(opts *bind.TransactOpts, _admin common.Address, _trustedSequencer common.Address, _trustedSequencerURL string, _networkName string, _lastAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.contract.Transact(opts, "initializeUpgrade", _admin, _trustedSequencer, _trustedSequencerURL, _networkName, _lastAccInputHash)
}

// InitializeUpgrade is a paid mutator transaction binding the contract method 0x5d6717a5.
//
// Solidity: function initializeUpgrade(address _admin, address _trustedSequencer, string _trustedSequencerURL, string _networkName, bytes32 _lastAccInputHash) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) InitializeUpgrade(_admin common.Address, _trustedSequencer common.Address, _trustedSequencerURL string, _networkName string, _lastAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.InitializeUpgrade(&_Polygonzkevmexistentetrogpessimistic.TransactOpts, _admin, _trustedSequencer, _trustedSequencerURL, _networkName, _lastAccInputHash)
}

// InitializeUpgrade is a paid mutator transaction binding the contract method 0x5d6717a5.
//
// Solidity: function initializeUpgrade(address _admin, address _trustedSequencer, string _trustedSequencerURL, string _networkName, bytes32 _lastAccInputHash) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactorSession) InitializeUpgrade(_admin common.Address, _trustedSequencer common.Address, _trustedSequencerURL string, _networkName string, _lastAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.InitializeUpgrade(&_Polygonzkevmexistentetrogpessimistic.TransactOpts, _admin, _trustedSequencer, _trustedSequencerURL, _networkName, _lastAccInputHash)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactor) OnVerifyBatches(opts *bind.TransactOpts, lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.contract.Transact(opts, "onVerifyBatches", lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.OnVerifyBatches(&_Polygonzkevmexistentetrogpessimistic.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactorSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.OnVerifyBatches(&_Polygonzkevmexistentetrogpessimistic.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x669adece.
//
// Solidity: function rollbackBatches(uint64 targetBatch, bytes32 accInputHashToRollback) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactor) RollbackBatches(opts *bind.TransactOpts, targetBatch uint64, accInputHashToRollback [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.contract.Transact(opts, "rollbackBatches", targetBatch, accInputHashToRollback)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x669adece.
//
// Solidity: function rollbackBatches(uint64 targetBatch, bytes32 accInputHashToRollback) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) RollbackBatches(targetBatch uint64, accInputHashToRollback [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.RollbackBatches(&_Polygonzkevmexistentetrogpessimistic.TransactOpts, targetBatch, accInputHashToRollback)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x669adece.
//
// Solidity: function rollbackBatches(uint64 targetBatch, bytes32 accInputHashToRollback) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactorSession) RollbackBatches(targetBatch uint64, accInputHashToRollback [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.RollbackBatches(&_Polygonzkevmexistentetrogpessimistic.TransactOpts, targetBatch, accInputHashToRollback)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xb910e0f9.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint32 l1InfoTreeLeafCount, uint64 maxSequenceTimestamp, bytes32 expectedFinalAccInputHash, address l2Coinbase) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactor) SequenceBatches(opts *bind.TransactOpts, batches []PolygonRollupBaseEtrogPessimisticBatchData, l1InfoTreeLeafCount uint32, maxSequenceTimestamp uint64, expectedFinalAccInputHash [32]byte, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.contract.Transact(opts, "sequenceBatches", batches, l1InfoTreeLeafCount, maxSequenceTimestamp, expectedFinalAccInputHash, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xb910e0f9.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint32 l1InfoTreeLeafCount, uint64 maxSequenceTimestamp, bytes32 expectedFinalAccInputHash, address l2Coinbase) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) SequenceBatches(batches []PolygonRollupBaseEtrogPessimisticBatchData, l1InfoTreeLeafCount uint32, maxSequenceTimestamp uint64, expectedFinalAccInputHash [32]byte, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.SequenceBatches(&_Polygonzkevmexistentetrogpessimistic.TransactOpts, batches, l1InfoTreeLeafCount, maxSequenceTimestamp, expectedFinalAccInputHash, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xb910e0f9.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint32 l1InfoTreeLeafCount, uint64 maxSequenceTimestamp, bytes32 expectedFinalAccInputHash, address l2Coinbase) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactorSession) SequenceBatches(batches []PolygonRollupBaseEtrogPessimisticBatchData, l1InfoTreeLeafCount uint32, maxSequenceTimestamp uint64, expectedFinalAccInputHash [32]byte, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.SequenceBatches(&_Polygonzkevmexistentetrogpessimistic.TransactOpts, batches, l1InfoTreeLeafCount, maxSequenceTimestamp, expectedFinalAccInputHash, l2Coinbase)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactor) SequenceForceBatches(opts *bind.TransactOpts, batches []PolygonRollupBaseEtrogPessimisticBatchData) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.contract.Transact(opts, "sequenceForceBatches", batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) SequenceForceBatches(batches []PolygonRollupBaseEtrogPessimisticBatchData) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.SequenceForceBatches(&_Polygonzkevmexistentetrogpessimistic.TransactOpts, batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactorSession) SequenceForceBatches(batches []PolygonRollupBaseEtrogPessimisticBatchData) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.SequenceForceBatches(&_Polygonzkevmexistentetrogpessimistic.TransactOpts, batches)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactor) SetForceBatchAddress(opts *bind.TransactOpts, newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.contract.Transact(opts, "setForceBatchAddress", newForceBatchAddress)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) SetForceBatchAddress(newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.SetForceBatchAddress(&_Polygonzkevmexistentetrogpessimistic.TransactOpts, newForceBatchAddress)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactorSession) SetForceBatchAddress(newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.SetForceBatchAddress(&_Polygonzkevmexistentetrogpessimistic.TransactOpts, newForceBatchAddress)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactor) SetForceBatchTimeout(opts *bind.TransactOpts, newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.contract.Transact(opts, "setForceBatchTimeout", newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.SetForceBatchTimeout(&_Polygonzkevmexistentetrogpessimistic.TransactOpts, newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactorSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.SetForceBatchTimeout(&_Polygonzkevmexistentetrogpessimistic.TransactOpts, newforceBatchTimeout)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.SetTrustedSequencer(&_Polygonzkevmexistentetrogpessimistic.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.SetTrustedSequencer(&_Polygonzkevmexistentetrogpessimistic.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.SetTrustedSequencerURL(&_Polygonzkevmexistentetrogpessimistic.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.SetTrustedSequencerURL(&_Polygonzkevmexistentetrogpessimistic.TransactOpts, newTrustedSequencerURL)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.TransferAdminRole(&_Polygonzkevmexistentetrogpessimistic.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrogpessimistic.Contract.TransferAdminRole(&_Polygonzkevmexistentetrogpessimistic.TransactOpts, newPendingAdmin)
}

// PolygonzkevmexistentetrogpessimisticAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticAcceptAdminRoleIterator struct {
	Event *PolygonzkevmexistentetrogpessimisticAcceptAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogpessimisticAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogpessimisticAcceptAdminRole)
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
		it.Event = new(PolygonzkevmexistentetrogpessimisticAcceptAdminRole)
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
func (it *PolygonzkevmexistentetrogpessimisticAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogpessimisticAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogpessimisticAcceptAdminRole represents a AcceptAdminRole event raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*PolygonzkevmexistentetrogpessimisticAcceptAdminRoleIterator, error) {

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogpessimisticAcceptAdminRoleIterator{contract: _Polygonzkevmexistentetrogpessimistic.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogpessimisticAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogpessimisticAcceptAdminRole)
				if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
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
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) ParseAcceptAdminRole(log types.Log) (*PolygonzkevmexistentetrogpessimisticAcceptAdminRole, error) {
	event := new(PolygonzkevmexistentetrogpessimisticAcceptAdminRole)
	if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmexistentetrogpessimisticForceBatchIterator is returned from FilterForceBatch and is used to iterate over the raw logs and unpacked data for ForceBatch events raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticForceBatchIterator struct {
	Event *PolygonzkevmexistentetrogpessimisticForceBatch // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogpessimisticForceBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogpessimisticForceBatch)
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
		it.Event = new(PolygonzkevmexistentetrogpessimisticForceBatch)
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
func (it *PolygonzkevmexistentetrogpessimisticForceBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogpessimisticForceBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogpessimisticForceBatch represents a ForceBatch event raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticForceBatch struct {
	ForceBatchNum      uint64
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Transactions       []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterForceBatch is a free log retrieval operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) FilterForceBatch(opts *bind.FilterOpts, forceBatchNum []uint64) (*PolygonzkevmexistentetrogpessimisticForceBatchIterator, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.FilterLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogpessimisticForceBatchIterator{contract: _Polygonzkevmexistentetrogpessimistic.contract, event: "ForceBatch", logs: logs, sub: sub}, nil
}

// WatchForceBatch is a free log subscription operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) WatchForceBatch(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogpessimisticForceBatch, forceBatchNum []uint64) (event.Subscription, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.WatchLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogpessimisticForceBatch)
				if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "ForceBatch", log); err != nil {
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
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) ParseForceBatch(log types.Log) (*PolygonzkevmexistentetrogpessimisticForceBatch, error) {
	event := new(PolygonzkevmexistentetrogpessimisticForceBatch)
	if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "ForceBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmexistentetrogpessimisticInitialSequenceBatchesIterator is returned from FilterInitialSequenceBatches and is used to iterate over the raw logs and unpacked data for InitialSequenceBatches events raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticInitialSequenceBatchesIterator struct {
	Event *PolygonzkevmexistentetrogpessimisticInitialSequenceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogpessimisticInitialSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogpessimisticInitialSequenceBatches)
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
		it.Event = new(PolygonzkevmexistentetrogpessimisticInitialSequenceBatches)
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
func (it *PolygonzkevmexistentetrogpessimisticInitialSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogpessimisticInitialSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogpessimisticInitialSequenceBatches represents a InitialSequenceBatches event raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticInitialSequenceBatches struct {
	Transactions       []byte
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterInitialSequenceBatches is a free log retrieval operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) FilterInitialSequenceBatches(opts *bind.FilterOpts) (*PolygonzkevmexistentetrogpessimisticInitialSequenceBatchesIterator, error) {

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.FilterLogs(opts, "InitialSequenceBatches")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogpessimisticInitialSequenceBatchesIterator{contract: _Polygonzkevmexistentetrogpessimistic.contract, event: "InitialSequenceBatches", logs: logs, sub: sub}, nil
}

// WatchInitialSequenceBatches is a free log subscription operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) WatchInitialSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogpessimisticInitialSequenceBatches) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.WatchLogs(opts, "InitialSequenceBatches")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogpessimisticInitialSequenceBatches)
				if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "InitialSequenceBatches", log); err != nil {
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
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) ParseInitialSequenceBatches(log types.Log) (*PolygonzkevmexistentetrogpessimisticInitialSequenceBatches, error) {
	event := new(PolygonzkevmexistentetrogpessimisticInitialSequenceBatches)
	if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "InitialSequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmexistentetrogpessimisticInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticInitializedIterator struct {
	Event *PolygonzkevmexistentetrogpessimisticInitialized // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogpessimisticInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogpessimisticInitialized)
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
		it.Event = new(PolygonzkevmexistentetrogpessimisticInitialized)
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
func (it *PolygonzkevmexistentetrogpessimisticInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogpessimisticInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogpessimisticInitialized represents a Initialized event raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) FilterInitialized(opts *bind.FilterOpts) (*PolygonzkevmexistentetrogpessimisticInitializedIterator, error) {

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogpessimisticInitializedIterator{contract: _Polygonzkevmexistentetrogpessimistic.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogpessimisticInitialized) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogpessimisticInitialized)
				if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) ParseInitialized(log types.Log) (*PolygonzkevmexistentetrogpessimisticInitialized, error) {
	event := new(PolygonzkevmexistentetrogpessimisticInitialized)
	if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmexistentetrogpessimisticRollbackBatchesIterator is returned from FilterRollbackBatches and is used to iterate over the raw logs and unpacked data for RollbackBatches events raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticRollbackBatchesIterator struct {
	Event *PolygonzkevmexistentetrogpessimisticRollbackBatches // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogpessimisticRollbackBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogpessimisticRollbackBatches)
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
		it.Event = new(PolygonzkevmexistentetrogpessimisticRollbackBatches)
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
func (it *PolygonzkevmexistentetrogpessimisticRollbackBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogpessimisticRollbackBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogpessimisticRollbackBatches represents a RollbackBatches event raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticRollbackBatches struct {
	TargetBatch            uint64
	AccInputHashToRollback [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRollbackBatches is a free log retrieval operation binding the contract event 0x1125aaf62d132d8e2d02005114f8fc360ff204c3105e4f1a700a1340dc55d5b1.
//
// Solidity: event RollbackBatches(uint64 indexed targetBatch, bytes32 accInputHashToRollback)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) FilterRollbackBatches(opts *bind.FilterOpts, targetBatch []uint64) (*PolygonzkevmexistentetrogpessimisticRollbackBatchesIterator, error) {

	var targetBatchRule []interface{}
	for _, targetBatchItem := range targetBatch {
		targetBatchRule = append(targetBatchRule, targetBatchItem)
	}

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.FilterLogs(opts, "RollbackBatches", targetBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogpessimisticRollbackBatchesIterator{contract: _Polygonzkevmexistentetrogpessimistic.contract, event: "RollbackBatches", logs: logs, sub: sub}, nil
}

// WatchRollbackBatches is a free log subscription operation binding the contract event 0x1125aaf62d132d8e2d02005114f8fc360ff204c3105e4f1a700a1340dc55d5b1.
//
// Solidity: event RollbackBatches(uint64 indexed targetBatch, bytes32 accInputHashToRollback)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) WatchRollbackBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogpessimisticRollbackBatches, targetBatch []uint64) (event.Subscription, error) {

	var targetBatchRule []interface{}
	for _, targetBatchItem := range targetBatch {
		targetBatchRule = append(targetBatchRule, targetBatchItem)
	}

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.WatchLogs(opts, "RollbackBatches", targetBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogpessimisticRollbackBatches)
				if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "RollbackBatches", log); err != nil {
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
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) ParseRollbackBatches(log types.Log) (*PolygonzkevmexistentetrogpessimisticRollbackBatches, error) {
	event := new(PolygonzkevmexistentetrogpessimisticRollbackBatches)
	if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "RollbackBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmexistentetrogpessimisticSequenceBatchesIterator is returned from FilterSequenceBatches and is used to iterate over the raw logs and unpacked data for SequenceBatches events raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticSequenceBatchesIterator struct {
	Event *PolygonzkevmexistentetrogpessimisticSequenceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogpessimisticSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogpessimisticSequenceBatches)
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
		it.Event = new(PolygonzkevmexistentetrogpessimisticSequenceBatches)
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
func (it *PolygonzkevmexistentetrogpessimisticSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogpessimisticSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogpessimisticSequenceBatches represents a SequenceBatches event raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticSequenceBatches struct {
	NumBatch   uint64
	L1InfoRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSequenceBatches is a free log retrieval operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) FilterSequenceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonzkevmexistentetrogpessimisticSequenceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.FilterLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogpessimisticSequenceBatchesIterator{contract: _Polygonzkevmexistentetrogpessimistic.contract, event: "SequenceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceBatches is a free log subscription operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) WatchSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogpessimisticSequenceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.WatchLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogpessimisticSequenceBatches)
				if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
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
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) ParseSequenceBatches(log types.Log) (*PolygonzkevmexistentetrogpessimisticSequenceBatches, error) {
	event := new(PolygonzkevmexistentetrogpessimisticSequenceBatches)
	if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmexistentetrogpessimisticSequenceForceBatchesIterator is returned from FilterSequenceForceBatches and is used to iterate over the raw logs and unpacked data for SequenceForceBatches events raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticSequenceForceBatchesIterator struct {
	Event *PolygonzkevmexistentetrogpessimisticSequenceForceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogpessimisticSequenceForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogpessimisticSequenceForceBatches)
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
		it.Event = new(PolygonzkevmexistentetrogpessimisticSequenceForceBatches)
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
func (it *PolygonzkevmexistentetrogpessimisticSequenceForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogpessimisticSequenceForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogpessimisticSequenceForceBatches represents a SequenceForceBatches event raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticSequenceForceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceForceBatches is a free log retrieval operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) FilterSequenceForceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonzkevmexistentetrogpessimisticSequenceForceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.FilterLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogpessimisticSequenceForceBatchesIterator{contract: _Polygonzkevmexistentetrogpessimistic.contract, event: "SequenceForceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceForceBatches is a free log subscription operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) WatchSequenceForceBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogpessimisticSequenceForceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.WatchLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogpessimisticSequenceForceBatches)
				if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
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
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) ParseSequenceForceBatches(log types.Log) (*PolygonzkevmexistentetrogpessimisticSequenceForceBatches, error) {
	event := new(PolygonzkevmexistentetrogpessimisticSequenceForceBatches)
	if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmexistentetrogpessimisticSetForceBatchAddressIterator is returned from FilterSetForceBatchAddress and is used to iterate over the raw logs and unpacked data for SetForceBatchAddress events raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticSetForceBatchAddressIterator struct {
	Event *PolygonzkevmexistentetrogpessimisticSetForceBatchAddress // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogpessimisticSetForceBatchAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogpessimisticSetForceBatchAddress)
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
		it.Event = new(PolygonzkevmexistentetrogpessimisticSetForceBatchAddress)
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
func (it *PolygonzkevmexistentetrogpessimisticSetForceBatchAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogpessimisticSetForceBatchAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogpessimisticSetForceBatchAddress represents a SetForceBatchAddress event raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticSetForceBatchAddress struct {
	NewForceBatchAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchAddress is a free log retrieval operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) FilterSetForceBatchAddress(opts *bind.FilterOpts) (*PolygonzkevmexistentetrogpessimisticSetForceBatchAddressIterator, error) {

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.FilterLogs(opts, "SetForceBatchAddress")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogpessimisticSetForceBatchAddressIterator{contract: _Polygonzkevmexistentetrogpessimistic.contract, event: "SetForceBatchAddress", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchAddress is a free log subscription operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) WatchSetForceBatchAddress(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogpessimisticSetForceBatchAddress) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.WatchLogs(opts, "SetForceBatchAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogpessimisticSetForceBatchAddress)
				if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "SetForceBatchAddress", log); err != nil {
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
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) ParseSetForceBatchAddress(log types.Log) (*PolygonzkevmexistentetrogpessimisticSetForceBatchAddress, error) {
	event := new(PolygonzkevmexistentetrogpessimisticSetForceBatchAddress)
	if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "SetForceBatchAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmexistentetrogpessimisticSetForceBatchTimeoutIterator is returned from FilterSetForceBatchTimeout and is used to iterate over the raw logs and unpacked data for SetForceBatchTimeout events raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticSetForceBatchTimeoutIterator struct {
	Event *PolygonzkevmexistentetrogpessimisticSetForceBatchTimeout // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogpessimisticSetForceBatchTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogpessimisticSetForceBatchTimeout)
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
		it.Event = new(PolygonzkevmexistentetrogpessimisticSetForceBatchTimeout)
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
func (it *PolygonzkevmexistentetrogpessimisticSetForceBatchTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogpessimisticSetForceBatchTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogpessimisticSetForceBatchTimeout represents a SetForceBatchTimeout event raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticSetForceBatchTimeout struct {
	NewforceBatchTimeout uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchTimeout is a free log retrieval operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) FilterSetForceBatchTimeout(opts *bind.FilterOpts) (*PolygonzkevmexistentetrogpessimisticSetForceBatchTimeoutIterator, error) {

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.FilterLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogpessimisticSetForceBatchTimeoutIterator{contract: _Polygonzkevmexistentetrogpessimistic.contract, event: "SetForceBatchTimeout", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchTimeout is a free log subscription operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) WatchSetForceBatchTimeout(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogpessimisticSetForceBatchTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.WatchLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogpessimisticSetForceBatchTimeout)
				if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
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
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) ParseSetForceBatchTimeout(log types.Log) (*PolygonzkevmexistentetrogpessimisticSetForceBatchTimeout, error) {
	event := new(PolygonzkevmexistentetrogpessimisticSetForceBatchTimeout)
	if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmexistentetrogpessimisticSetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticSetTrustedSequencerIterator struct {
	Event *PolygonzkevmexistentetrogpessimisticSetTrustedSequencer // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogpessimisticSetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogpessimisticSetTrustedSequencer)
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
		it.Event = new(PolygonzkevmexistentetrogpessimisticSetTrustedSequencer)
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
func (it *PolygonzkevmexistentetrogpessimisticSetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogpessimisticSetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogpessimisticSetTrustedSequencer represents a SetTrustedSequencer event raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticSetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*PolygonzkevmexistentetrogpessimisticSetTrustedSequencerIterator, error) {

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogpessimisticSetTrustedSequencerIterator{contract: _Polygonzkevmexistentetrogpessimistic.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogpessimisticSetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogpessimisticSetTrustedSequencer)
				if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
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
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) ParseSetTrustedSequencer(log types.Log) (*PolygonzkevmexistentetrogpessimisticSetTrustedSequencer, error) {
	event := new(PolygonzkevmexistentetrogpessimisticSetTrustedSequencer)
	if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmexistentetrogpessimisticSetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticSetTrustedSequencerURLIterator struct {
	Event *PolygonzkevmexistentetrogpessimisticSetTrustedSequencerURL // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogpessimisticSetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogpessimisticSetTrustedSequencerURL)
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
		it.Event = new(PolygonzkevmexistentetrogpessimisticSetTrustedSequencerURL)
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
func (it *PolygonzkevmexistentetrogpessimisticSetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogpessimisticSetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogpessimisticSetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticSetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*PolygonzkevmexistentetrogpessimisticSetTrustedSequencerURLIterator, error) {

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogpessimisticSetTrustedSequencerURLIterator{contract: _Polygonzkevmexistentetrogpessimistic.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogpessimisticSetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogpessimisticSetTrustedSequencerURL)
				if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
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
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) ParseSetTrustedSequencerURL(log types.Log) (*PolygonzkevmexistentetrogpessimisticSetTrustedSequencerURL, error) {
	event := new(PolygonzkevmexistentetrogpessimisticSetTrustedSequencerURL)
	if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmexistentetrogpessimisticTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticTransferAdminRoleIterator struct {
	Event *PolygonzkevmexistentetrogpessimisticTransferAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogpessimisticTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogpessimisticTransferAdminRole)
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
		it.Event = new(PolygonzkevmexistentetrogpessimisticTransferAdminRole)
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
func (it *PolygonzkevmexistentetrogpessimisticTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogpessimisticTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogpessimisticTransferAdminRole represents a TransferAdminRole event raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*PolygonzkevmexistentetrogpessimisticTransferAdminRoleIterator, error) {

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogpessimisticTransferAdminRoleIterator{contract: _Polygonzkevmexistentetrogpessimistic.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogpessimisticTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogpessimisticTransferAdminRole)
				if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
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
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) ParseTransferAdminRole(log types.Log) (*PolygonzkevmexistentetrogpessimisticTransferAdminRole, error) {
	event := new(PolygonzkevmexistentetrogpessimisticTransferAdminRole)
	if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmexistentetrogpessimisticUpdateEtrogSequenceIterator is returned from FilterUpdateEtrogSequence and is used to iterate over the raw logs and unpacked data for UpdateEtrogSequence events raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticUpdateEtrogSequenceIterator struct {
	Event *PolygonzkevmexistentetrogpessimisticUpdateEtrogSequence // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogpessimisticUpdateEtrogSequenceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogpessimisticUpdateEtrogSequence)
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
		it.Event = new(PolygonzkevmexistentetrogpessimisticUpdateEtrogSequence)
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
func (it *PolygonzkevmexistentetrogpessimisticUpdateEtrogSequenceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogpessimisticUpdateEtrogSequenceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogpessimisticUpdateEtrogSequence represents a UpdateEtrogSequence event raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticUpdateEtrogSequence struct {
	NumBatch           uint64
	Transactions       []byte
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdateEtrogSequence is a free log retrieval operation binding the contract event 0xd2c80353fc15ef62c6affc7cd6b7ab5b42c43290c50be3372e55ae552cecd19c.
//
// Solidity: event UpdateEtrogSequence(uint64 numBatch, bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) FilterUpdateEtrogSequence(opts *bind.FilterOpts) (*PolygonzkevmexistentetrogpessimisticUpdateEtrogSequenceIterator, error) {

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.FilterLogs(opts, "UpdateEtrogSequence")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogpessimisticUpdateEtrogSequenceIterator{contract: _Polygonzkevmexistentetrogpessimistic.contract, event: "UpdateEtrogSequence", logs: logs, sub: sub}, nil
}

// WatchUpdateEtrogSequence is a free log subscription operation binding the contract event 0xd2c80353fc15ef62c6affc7cd6b7ab5b42c43290c50be3372e55ae552cecd19c.
//
// Solidity: event UpdateEtrogSequence(uint64 numBatch, bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) WatchUpdateEtrogSequence(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogpessimisticUpdateEtrogSequence) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.WatchLogs(opts, "UpdateEtrogSequence")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogpessimisticUpdateEtrogSequence)
				if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "UpdateEtrogSequence", log); err != nil {
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

// ParseUpdateEtrogSequence is a log parse operation binding the contract event 0xd2c80353fc15ef62c6affc7cd6b7ab5b42c43290c50be3372e55ae552cecd19c.
//
// Solidity: event UpdateEtrogSequence(uint64 numBatch, bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) ParseUpdateEtrogSequence(log types.Log) (*PolygonzkevmexistentetrogpessimisticUpdateEtrogSequence, error) {
	event := new(PolygonzkevmexistentetrogpessimisticUpdateEtrogSequence)
	if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "UpdateEtrogSequence", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmexistentetrogpessimisticVerifyBatchesIterator is returned from FilterVerifyBatches and is used to iterate over the raw logs and unpacked data for VerifyBatches events raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticVerifyBatchesIterator struct {
	Event *PolygonzkevmexistentetrogpessimisticVerifyBatches // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogpessimisticVerifyBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogpessimisticVerifyBatches)
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
		it.Event = new(PolygonzkevmexistentetrogpessimisticVerifyBatches)
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
func (it *PolygonzkevmexistentetrogpessimisticVerifyBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogpessimisticVerifyBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogpessimisticVerifyBatches represents a VerifyBatches event raised by the Polygonzkevmexistentetrogpessimistic contract.
type PolygonzkevmexistentetrogpessimisticVerifyBatches struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatches is a free log retrieval operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) FilterVerifyBatches(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*PolygonzkevmexistentetrogpessimisticVerifyBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.FilterLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogpessimisticVerifyBatchesIterator{contract: _Polygonzkevmexistentetrogpessimistic.contract, event: "VerifyBatches", logs: logs, sub: sub}, nil
}

// WatchVerifyBatches is a free log subscription operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) WatchVerifyBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogpessimisticVerifyBatches, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonzkevmexistentetrogpessimistic.contract.WatchLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogpessimisticVerifyBatches)
				if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
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
func (_Polygonzkevmexistentetrogpessimistic *PolygonzkevmexistentetrogpessimisticFilterer) ParseVerifyBatches(log types.Log) (*PolygonzkevmexistentetrogpessimisticVerifyBatches, error) {
	event := new(PolygonzkevmexistentetrogpessimisticVerifyBatches)
	if err := _Polygonzkevmexistentetrogpessimistic.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

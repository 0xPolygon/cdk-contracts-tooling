// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonvalidiumetrog

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

// PolygonValidiumEtrogValidiumBatchData is an auto generated low-level Go binding around an user-defined struct.
type PolygonValidiumEtrogValidiumBatchData struct {
	TransactionsHash     [32]byte
	ForcedGlobalExitRoot [32]byte
	ForcedTimestamp      uint64
	ForcedBlockHashL1    [32]byte
}

// PolygonvalidiumetrogMetaData contains all meta data concerning the Polygonvalidiumetrog contract.
var PolygonvalidiumetrogMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"_rollupManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceWithDataAvailabilityNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwitchToSameValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"forceBatchNum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"}],\"name\":\"ForceBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"}],\"name\":\"InitialSequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"l1InfoRoot\",\"type\":\"bytes32\"}],\"name\":\"SequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newDataAvailabilityProtocol\",\"type\":\"address\"}],\"name\":\"SetDataAvailabilityProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newForceBatchAddress\",\"type\":\"address\"}],\"name\":\"SetForceBatchAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"SetForceBatchTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SwitchSequenceWithDataAvailability\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatches\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GLOBAL_EXIT_ROOT_MANAGER_L2\",\"outputs\":[{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_LIST_LEN_LEN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_DATA_LEN_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_EFFECTIVE_PERCENTAGE\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_R\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_S\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_V\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TIMESTAMP_RANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculatePolPerForceBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataAvailabilityProtocol\",\"outputs\":[{\"internalType\":\"contractIDataAvailabilityProtocol\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"polAmount\",\"type\":\"uint256\"}],\"name\":\"forceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"}],\"name\":\"generateInitializeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSequenceWithDataAvailabilityAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"onVerifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrog.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"maxSequenceTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initSequencedBatch\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"}],\"name\":\"sequenceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonValidiumEtrog.ValidiumBatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"maxSequenceTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initSequencedBatch\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataAvailabilityMessage\",\"type\":\"bytes\"}],\"name\":\"sequenceBatchesValidium\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrog.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"}],\"name\":\"sequenceForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIDataAvailabilityProtocol\",\"name\":\"newDataAvailabilityProtocol\",\"type\":\"address\"}],\"name\":\"setDataAvailabilityProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newForceBatchAddress\",\"type\":\"address\"}],\"name\":\"setForceBatchAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"setForceBatchTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"newIsSequenceWithDataAvailabilityAllowed\",\"type\":\"bool\"}],\"name\":\"switchSequenceWithDataAvailability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b506040516200520038038062005200833981016040819052620000359162000071565b6001600160a01b0393841660a052918316608052821660c0521660e052620000d9565b6001600160a01b03811681146200006e57600080fd5b50565b600080600080608085870312156200008857600080fd5b8451620000958162000058565b6020860151909450620000a88162000058565b6040860151909350620000bb8162000058565b6060860151909250620000ce8162000058565b939692955090935050565b60805160a05160c05160e051615014620001ec6000396000818161054d01528181610b5601528181610cc301528181610f0e015281816114f401528181611a6201528181611eb901528181611faf01528181612bb301528181612c3201528181612c5401528181612df701528181613005015281816130cd01528181613b7901528181613bf201528181613c140152613cbf0152600081816106ef01528181611107015281816111e10152818161218001528181612288015281816126da01526136b60152600081816107ab015281816113680152818161275c015281816132180152613738015260008181610803015281816108e201528181611f0201528181612d0201526131ec01526150146000f3fe608060405234801561001057600080fd5b506004361061031f5760003560e01c80637a5460c5116101a7578063c89e42df116100ee578063e46761c411610097578063eaeb077b11610071578063eaeb077b14610859578063f35dda471461086c578063f851a4401461087457600080fd5b8063e46761c4146107fe578063e57a0b4c14610825578063e7a7ed021461084557600080fd5b8063d7bc90ff116100c8578063d7bc90ff146107cd578063db5b0ed7146107d8578063def57e54146107eb57600080fd5b8063c89e42df14610773578063cfa8ed4714610786578063d02103ca146107a657600080fd5b8063a3c573eb11610150578063b0afe1541161012a578063b0afe15414610737578063c754c7ed14610743578063c7fffd4b1461076b57600080fd5b8063a3c573eb146106ea578063a652f26c14610711578063ada8f9191461072457600080fd5b806391cafe321161018157806391cafe32146106a95780639e001877146106bc5780639f26f840146106d757600080fd5b80637a5460c5146106525780637cd76b8b1461068e5780638c3d7301146106a157600080fd5b806342308fab1161026b578063542028d5116102145780636e05d2cd116101ee5780636e05d2cd146106235780636ff512cc1461062c578063712570221461063f57600080fd5b8063542028d5146105f3578063676870d2146105fb5780636b8616ce1461060357600080fd5b80634c21fef3116102455780634c21fef31461056f5780634e487706146105a457806352bdeb6d146105b757600080fd5b806342308fab14610507578063456052671461050f57806349b7b8021461054857600080fd5b80632acdc2b6116102cd5780633c351e10116102a75780633c351e10146104525780633cbc795b1461047257806340b5de6c146104af57600080fd5b80632acdc2b61461040a5780632c111c061461041f57806332c2d1531461043f57600080fd5b8063107bf28c116102fe578063107bf28c146103a357806311e892d4146103ab57806326782247146103c557600080fd5b8062d0295d14610324578063035089631461033f57806305835f371461035a575b600080fd5b61032c61089a565b6040519081526020015b60405180910390f35b610347602081565b60405161ffff9091168152602001610336565b6103966040518060400160405280600881526020017f80808401c9c3809400000000000000000000000000000000000000000000000081525081565b60405161033691906140d6565b6103966109a6565b6103b360f981565b60405160ff9091168152602001610336565b6001546103e59073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610336565b61041d610418366004614101565b610a34565b005b6008546103e59073ffffffffffffffffffffffffffffffffffffffff1681565b61041d61044d366004614171565b610b54565b6009546103e59073ffffffffffffffffffffffffffffffffffffffff1681565b60095461049a9074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff9091168152602001610336565b6104d67fff0000000000000000000000000000000000000000000000000000000000000081565b6040517fff000000000000000000000000000000000000000000000000000000000000009091168152602001610336565b61032c602481565b60075461052f9068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff9091168152602001610336565b6103e57f000000000000000000000000000000000000000000000000000000000000000081565b603c546105949074010000000000000000000000000000000000000000900460ff1681565b6040519015158152602001610336565b61041d6105b23660046141b3565b610c23565b6103966040518060400160405280600281526020017f80b800000000000000000000000000000000000000000000000000000000000081525081565b610396610e35565b610347601f81565b61032c6106113660046141b3565b60066020526000908152604090205481565b61032c60055481565b61041d61063a3660046141d0565b610e42565b61041d61064d366004614342565b610f0c565b6103966040518060400160405280600281526020017f80b900000000000000000000000000000000000000000000000000000000000081525081565b61041d61069c3660046141d0565b611730565b61041d6117fa565b61041d6106b73660046141d0565b6118cd565b6103e573a40d5f56745a118d0906a34e69aec8c0db1cb8fa81565b61041d6106e536600461443b565b6119e6565b6103e57f000000000000000000000000000000000000000000000000000000000000000081565b61039661071f36600461447d565b61207f565b61041d6107323660046141d0565b612464565b61032c6405ca1ab1e081565b60075461052f90700100000000000000000000000000000000900467ffffffffffffffff1681565b6103b360e481565b61041d6107813660046144f2565b61252e565b6002546103e59073ffffffffffffffffffffffffffffffffffffffff1681565b6103e57f000000000000000000000000000000000000000000000000000000000000000081565b61032c635ca1ab1e81565b61041d6107e6366004614569565b6125c1565b61041d6107f936600461463c565b612f26565b6103e57f000000000000000000000000000000000000000000000000000000000000000081565b603c546103e59073ffffffffffffffffffffffffffffffffffffffff1681565b60075461052f9067ffffffffffffffff1681565b61041d6108673660046146b9565b612f8e565b6103b3601b81565b6000546103e59062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b6040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152600090819073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015610929573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061094d9190614705565b6007549091506000906109789067ffffffffffffffff6801000000000000000082048116911661474d565b67ffffffffffffffff169050806000036109955760009250505090565b61099f8183614775565b9250505090565b600480546109b3906147b0565b80601f01602080910402602001604051908101604052809291908181526020018280546109df906147b0565b8015610a2c5780601f10610a0157610100808354040283529160200191610a2c565b820191906000526020600020905b815481529060010190602001808311610a0f57829003601f168201915b505050505081565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314610a8b576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603c5474010000000000000000000000000000000000000000900460ff16151581151503610ae5576040517f5f0e7abe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603c80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1674010000000000000000000000000000000000000000831515021790556040517ff32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f4190600090a150565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314610bc3576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff168367ffffffffffffffff167f9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f596684604051610c1691815260200190565b60405180910390a3505050565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314610c7a576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff82161115610cc1576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166315064c966040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d2c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d509190614803565b610db15760075467ffffffffffffffff700100000000000000000000000000000000909104811690821610610db1576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600780547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527fa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b906020015b60405180910390a150565b600380546109b3906147b0565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314610e99576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc090602001610e2a565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314610f7b576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600054610100900460ff1615808015610f9b5750600054600160ff909116105b80610fb55750303b158015610fb5575060005460ff166001145b611046576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156110a457600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b606073ffffffffffffffffffffffffffffffffffffffff851615611309576040517fc00f14ab00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063c00f14ab90602401600060405180830381865afa15801561114e573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526111949190810190614820565b6040517f318aee3d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff878116600483015291925060009182917f00000000000000000000000000000000000000000000000000000000000000009091169063318aee3d906024016040805180830381865afa158015611229573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061124d9190614897565b915091508163ffffffff166000146112c5576009805463ffffffff841674010000000000000000000000000000000000000000027fffffffffffffffff00000000000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff841617179055611306565b600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff89161790555b50505b60095460009061135190889073ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900463ffffffff168561207f565b9050600081805190602001209050600042905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156113d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113f59190614705565b90506000808483858f6114096001436148d1565b60408051602081019790975286019490945260608086019390935260c09190911b7fffffffffffffffff000000000000000000000000000000000000000000000000166080850152901b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016608883015240609c82015260bc01604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001815290829052805160209091012060058190557f9a908e73000000000000000000000000000000000000000000000000000000008252600160048301526024820181905291507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303816000875af1158015611552573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061157691906148ea565b508c600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508b600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508860039081611608919061494d565b506004611615898261494d565b508c600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062069780600760106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507f060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f85838e6040516116b693929190614a67565b60405180910390a1505050505050801561172757600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314611787576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603c80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a90602001610e2a565b60015473ffffffffffffffffffffffffffffffffffffffff16331461184b576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600154600080547fffffffffffffffffffff0000000000000000000000000000000000000000ffff1673ffffffffffffffffffffffffffffffffffffffff9092166201000081029290921790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e9060200160405180910390a1565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314611924576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60085473ffffffffffffffffffffffffffffffffffffffff16611973576040517fc89374d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb90602001610e2a565b60085473ffffffffffffffffffffffffffffffffffffffff168015801590611a24575073ffffffffffffffffffffffffffffffffffffffff81163314155b15611a5b576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b4262093a807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166330c27dde6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611acb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611aef91906148ea565b611af99190614aa6565b67ffffffffffffffff161115611b3b576040517f3d49ed4c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b816000819003611b77576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8811115611bb3576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075467ffffffffffffffff80821691611bdb91849168010000000000000000900416614ac7565b1115611c13576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6007546005546801000000000000000090910467ffffffffffffffff169060005b83811015611eb3576000878783818110611c5057611c50614ada565b9050602002810190611c629190614b09565b611c6b90614b47565b905083611c7781614bb5565b825180516020918201208185015160408087015160608801519151959a50929550600094611ce4948794929101938452602084019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166040830152604882015260680190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff8916600090815260069093529120549091508114611d6d576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8616600090815260066020526040812055611d926001886148d1565b8403611e015742600760109054906101000a900467ffffffffffffffff168460400151611dbf9190614aa6565b67ffffffffffffffff161115611e01576040517fc44a082100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208381015160408086015160608088015183519586018b90529285018790528481019390935260c01b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808401523390911b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088830152609c82015260bc016040516020818303038152906040528051906020012094505050508080611eab90614bdc565b915050611c34565b50611f297f000000000000000000000000000000000000000000000000000000000000000084611ee161089a565b611eeb9190614c14565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169190613460565b60058190556007805467ffffffffffffffff841668010000000000000000027fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff9091161790556040517f9a908e7300000000000000000000000000000000000000000000000000000000815260009073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690639a908e7390611ffb908790869060040167ffffffffffffffff929092168252602082015260400190565b6020604051808303816000875af115801561201a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061203e91906148ea565b60405190915067ffffffffffffffff8216907f648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a490600090a250505050505050565b6060600085858573a40d5f56745a118d0906a34e69aec8c0db1cb8fa6000876040516024016120b396959493929190614c2b565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff811bff70000000000000000000000000000000000000000000000000000000017905283519091506060906000036122045760f9601f83516121489190614c8e565b6040518060400160405280600881526020017f80808401c9c380940000000000000000000000000000000000000000000000008152507f00000000000000000000000000000000000000000000000000000000000000006040518060400160405280600281526020017f80b800000000000000000000000000000000000000000000000000000000000081525060e4876040516020016121ee9796959493929190614ca9565b6040516020818303038152906040529050612308565b815161ffff1015612241576040517f248b8f8200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815160f9612250602083614c8e565b6040518060400160405280600881526020017f80808401c9c380940000000000000000000000000000000000000000000000008152507f00000000000000000000000000000000000000000000000000000000000000006040518060400160405280600281526020017f80b900000000000000000000000000000000000000000000000000000000000081525085886040516020016122f59796959493929190614d8c565b6040516020818303038152906040529150505b805160208083019190912060408051600080825293810180835292909252601b908201526405ca1ab1e06060820152635ca1ab1e608082015260019060a0016020604051602081039080840390855afa158015612369573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81166123e1576040517fcd16196600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040516000906124279084906405ca1ab1e090635ca1ab1e90601b907fff0000000000000000000000000000000000000000000000000000000000000090602001614e6f565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529450505050505b949350505050565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff1633146124bb576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001610e2a565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314612585576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6003612591828261494d565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b2081604051610e2a91906140d6565b60025473ffffffffffffffffffffffffffffffffffffffff163314612612576040517f11e7be1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b85600081900361264e576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e881111561268a576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612695602442614ac7565b8667ffffffffffffffff1611156126d8576040517f0a00feb300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379e2cf976040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561274057600080fd5b505af1158015612754573d6000803e3d6000fd5b5050505060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635ca1e1656040518163ffffffff1660e01b8152600401602060405180830381865afa1580156127c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127e99190614705565b60075460055491925068010000000000000000900467ffffffffffffffff1690816000805b86811015612b255760008e8e8381811061282a5761282a614ada565b9050608002018036038101906128409190614ecb565b604081015190915067ffffffffffffffff1615612a34578561286181614bb5565b965050600081600001518260200151836040015184606001516040516020016128c89493929190938452602084019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166040830152604882015260680190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff8a16600090815260069093529120549091508114612951576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b858260000151836020015184604001518f86606001516040516020016129eb969594939291909586526020860194909452604085019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166060808501919091521b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166068830152607c820152609c0190565b604051602081830303815290604052805190602001209550600660008867ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206000905550612b12565b8051604051612a50918591602001918252602082015260400190565b604051602081830303815290604052805190602001209250848160000151888f8e6000801b604051602001612af9969594939291909586526020860194909452604085019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166060808501919091521b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166068830152607c820152609c0190565b6040516020818303038152906040528051906020012094505b5080612b1d81614bdc565b91505061280e565b5060075467ffffffffffffffff9081169085161115612b70576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60058390558567ffffffffffffffff85811690841614612c26576000612b96848761474d565b9050612bac67ffffffffffffffff8216836148d1565b9150612be57f00000000000000000000000000000000000000000000000000000000000000008267ffffffffffffffff16611ee161089a565b50600780547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8816021790555b8015612db557612d2a337f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663477fa2706040518163ffffffff1660e01b8152600401602060405180830381865afa158015612cbd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ce19190614705565b612ceb9190614c14565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016929190613539565b603c546040517f3b51be4b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690633b51be4b90612d849085908d908d90600401614f62565b60006040518083038186803b158015612d9c57600080fd5b505afa158015612db0573d6000803e3d6000fd5b505050505b6040517f9a908e7300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff88166004820152602481018590526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303816000875af1158015612e55573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e7991906148ea565b9050612e85888261474d565b67ffffffffffffffff168c67ffffffffffffffff1614612ed1576040517f1a070d9a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff167f3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e76688604051612f0d91815260200190565b60405180910390a2505050505050505050505050505050565b603c5474010000000000000000000000000000000000000000900460ff16612f7a576040517f821935b400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612f87858585858561359d565b5050505050565b60085473ffffffffffffffffffffffffffffffffffffffff168015801590612fcc575073ffffffffffffffffffffffffffffffffffffffff81163314155b15613003576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166315064c966040518163ffffffff1660e01b8152600401602060405180830381865afa15801561306e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130929190614803565b156130c9576040517f39258d1800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663604691696040518163ffffffff1660e01b8152600401602060405180830381865afa158015613136573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061315a9190614705565b905082811115613196576040517f2354600f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113888411156131d2576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61321473ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016333084613539565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613281573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132a59190614705565b6007805491925067ffffffffffffffff9091169060006132c483614bb5565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505085856040516132fb929190614f7c565b60405190819003902081426133116001436148d1565b60408051602081019590955284019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166060830152406068820152608801604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012060075467ffffffffffffffff166000908152600690935291205532330361340957600754604080518381523360208201526060818301819052600090820152905167ffffffffffffffff909216917ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319181900360800190a2613458565b60075460405167ffffffffffffffff909116907ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319061344f90849033908b908b90614f8c565b60405180910390a25b505050505050565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526135349084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152613deb565b505050565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526135979085907f23b872dd00000000000000000000000000000000000000000000000000000000906084016134b2565b50505050565b60025473ffffffffffffffffffffffffffffffffffffffff1633146135ee576040517f11e7be1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83600081900361362a576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8811115613666576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613671602442614ac7565b8467ffffffffffffffff1611156136b4576040517f0a00feb300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379e2cf976040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561371c57600080fd5b505af1158015613730573d6000803e3d6000fd5b5050505060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635ca1e1656040518163ffffffff1660e01b8152600401602060405180830381865afa1580156137a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137c59190614705565b60075460055491925068010000000000000000900467ffffffffffffffff16908160005b85811015613aeb5760008b8b8381811061380557613805614ada565b90506020028101906138179190614b09565b61382090614b47565b8051805160209091012060408201519192509067ffffffffffffffff1615613a05578561384c81614bb5565b9650506000818360200151846040015185606001516040516020016138af9493929190938452602084019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166040830152604882015260680190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff8a16600090815260069093529120549091508114613938576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208381015160408086015160608088015183519586018c90529285018790528481019390935260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166080840152908c901b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088830152609c82015260bc01604051602081830303815290604052805190602001209550600660008867ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206000905550613ad6565b8151516201d4c01015613a44576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160208101879052908101829052606080820189905260c08d901b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808301528a901b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660888201526000609c82015260bc016040516020818303038152906040528051906020012094505b50508080613ae390614bdc565b9150506137e9565b5060075467ffffffffffffffff9081169084161115613b36576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60058290558467ffffffffffffffff84811690831614613bec576000613b5c838661474d565b9050613b7267ffffffffffffffff8216836148d1565b9150613bab7f00000000000000000000000000000000000000000000000000000000000000008267ffffffffffffffff16611ee161089a565b50600780547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8716021790555b613c7d337f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663477fa2706040518163ffffffff1660e01b8152600401602060405180830381865afa158015612cbd573d6000803e3d6000fd5b6040517f9a908e7300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff87166004820152602481018490526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303816000875af1158015613d1d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613d4191906148ea565b9050613d4d878261474d565b67ffffffffffffffff168967ffffffffffffffff1614613d99576040517f1a070d9a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff167f3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e76687604051613dd591815260200190565b60405180910390a2505050505050505050505050565b6000613e4d826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16613ef79092919063ffffffff16565b8051909150156135345780806020019051810190613e6b9190614803565b613534576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161103d565b606061245c8484600085856000808673ffffffffffffffffffffffffffffffffffffffff168587604051613f2b9190614fcc565b60006040518083038185875af1925050503d8060008114613f68576040519150601f19603f3d011682016040523d82523d6000602084013e613f6d565b606091505b5091509150613f7e87838387613f89565b979650505050505050565b6060831561401f5782516000036140185773ffffffffffffffffffffffffffffffffffffffff85163b614018576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161103d565b508161245c565b61245c83838151156140345781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161103d91906140d6565b60005b8381101561408357818101518382015260200161406b565b50506000910152565b600081518084526140a4816020860160208601614068565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006140e9602083018461408c565b9392505050565b80151581146140fe57600080fd5b50565b60006020828403121561411357600080fd5b81356140e9816140f0565b67ffffffffffffffff811681146140fe57600080fd5b803561413f8161411e565b919050565b73ffffffffffffffffffffffffffffffffffffffff811681146140fe57600080fd5b803561413f81614144565b60008060006060848603121561418657600080fd5b83356141918161411e565b92506020840135915060408401356141a881614144565b809150509250925092565b6000602082840312156141c557600080fd5b81356140e98161411e565b6000602082840312156141e257600080fd5b81356140e981614144565b63ffffffff811681146140fe57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516080810167ffffffffffffffff81118282101715614251576142516141ff565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561429e5761429e6141ff565b604052919050565b600067ffffffffffffffff8211156142c0576142c06141ff565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f8301126142fd57600080fd5b813561431061430b826142a6565b614257565b81815284602083860101111561432557600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060008060c0878903121561435b57600080fd5b863561436681614144565b9550602087013561437681614144565b94506040870135614386816141ed565b9350606087013561439681614144565b9250608087013567ffffffffffffffff808211156143b357600080fd5b6143bf8a838b016142ec565b935060a08901359150808211156143d557600080fd5b506143e289828a016142ec565b9150509295509295509295565b60008083601f84011261440157600080fd5b50813567ffffffffffffffff81111561441957600080fd5b6020830191508360208260051b850101111561443457600080fd5b9250929050565b6000806020838503121561444e57600080fd5b823567ffffffffffffffff81111561446557600080fd5b614471858286016143ef565b90969095509350505050565b6000806000806080858703121561449357600080fd5b843561449e816141ed565b935060208501356144ae81614144565b925060408501356144be816141ed565b9150606085013567ffffffffffffffff8111156144da57600080fd5b6144e6878288016142ec565b91505092959194509250565b60006020828403121561450457600080fd5b813567ffffffffffffffff81111561451b57600080fd5b61245c848285016142ec565b60008083601f84011261453957600080fd5b50813567ffffffffffffffff81111561455157600080fd5b60208301915083602082850101111561443457600080fd5b600080600080600080600060a0888a03121561458457600080fd5b873567ffffffffffffffff8082111561459c57600080fd5b818a0191508a601f8301126145b057600080fd5b8135818111156145bf57600080fd5b8b60208260071b85010111156145d457600080fd5b602083019950809850506145ea60208b01614134565b96506145f860408b01614134565b955061460660608b01614166565b945060808a013591508082111561461c57600080fd5b506146298a828b01614527565b989b979a50959850939692959293505050565b60008060008060006080868803121561465457600080fd5b853567ffffffffffffffff81111561466b57600080fd5b614677888289016143ef565b909650945050602086013561468b8161411e565b9250604086013561469b8161411e565b915060608601356146ab81614144565b809150509295509295909350565b6000806000604084860312156146ce57600080fd5b833567ffffffffffffffff8111156146e557600080fd5b6146f186828701614527565b909790965060209590950135949350505050565b60006020828403121561471757600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b67ffffffffffffffff82811682821603908082111561476e5761476e61471e565b5092915050565b6000826147ab577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b600181811c908216806147c457607f821691505b6020821081036147fd577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b60006020828403121561481557600080fd5b81516140e9816140f0565b60006020828403121561483257600080fd5b815167ffffffffffffffff81111561484957600080fd5b8201601f8101841361485a57600080fd5b805161486861430b826142a6565b81815285602083850101111561487d57600080fd5b61488e826020830160208601614068565b95945050505050565b600080604083850312156148aa57600080fd5b82516148b5816141ed565b60208401519092506148c681614144565b809150509250929050565b818103818111156148e4576148e461471e565b92915050565b6000602082840312156148fc57600080fd5b81516140e98161411e565b601f82111561353457600081815260208120601f850160051c8101602086101561492e5750805b601f850160051c820191505b818110156134585782815560010161493a565b815167ffffffffffffffff811115614967576149676141ff565b61497b8161497584546147b0565b84614907565b602080601f8311600181146149ce57600084156149985750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555613458565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015614a1b578886015182559484019460019091019084016149fc565b5085821015614a5757878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b606081526000614a7a606083018661408c565b905083602083015273ffffffffffffffffffffffffffffffffffffffff83166040830152949350505050565b67ffffffffffffffff81811683821601908082111561476e5761476e61471e565b808201808211156148e4576148e461471e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81833603018112614b3d57600080fd5b9190910192915050565b600060808236031215614b5957600080fd5b614b6161422e565b823567ffffffffffffffff811115614b7857600080fd5b614b84368286016142ec565b825250602083013560208201526040830135614b9f8161411e565b6040820152606092830135928101929092525090565b600067ffffffffffffffff808316818103614bd257614bd261471e565b6001019392505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614c0d57614c0d61471e565b5060010190565b80820281158282048414176148e4576148e461471e565b600063ffffffff808916835273ffffffffffffffffffffffffffffffffffffffff8089166020850152818816604085015280871660608501528086166080850152505060c060a0830152614c8260c083018461408c565b98975050505050505050565b61ffff81811683821601908082111561476e5761476e61471e565b60007fff00000000000000000000000000000000000000000000000000000000000000808a60f81b1683527fffff0000000000000000000000000000000000000000000000000000000000008960f01b1660018401528751614d12816003860160208c01614068565b80840190507fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008860601b1660038201528651614d55816017840160208b01614068565b808201915050818660f81b16601782015284519150614d7b826018830160208801614068565b016018019998505050505050505050565b7fff000000000000000000000000000000000000000000000000000000000000008860f81b16815260007fffff000000000000000000000000000000000000000000000000000000000000808960f01b1660018401528751614df5816003860160208c01614068565b80840190507fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008860601b1660038201528651614e38816017840160208b01614068565b808201915050818660f01b16601782015284519150614e5e826019830160208801614068565b016019019998505050505050505050565b60008651614e81818460208b01614068565b9190910194855250602084019290925260f81b7fff000000000000000000000000000000000000000000000000000000000000009081166040840152166041820152604201919050565b600060808284031215614edd57600080fd5b614ee561422e565b82358152602083013560208201526040830135614f018161411e565b60408201526060928301359281019290925250919050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b83815260406020820152600061488e604083018486614f19565b8183823760009101908152919050565b84815273ffffffffffffffffffffffffffffffffffffffff84166020820152606060408201526000614fc2606083018486614f19565b9695505050505050565b60008251614b3d81846020870161406856fea264697066735822122012c6e0437be7f44797e64b6e94b288e36d741e2455c4f5eb852f0e1f6ab0cc5164736f6c63430008140033",
}

// PolygonvalidiumetrogABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonvalidiumetrogMetaData.ABI instead.
var PolygonvalidiumetrogABI = PolygonvalidiumetrogMetaData.ABI

// PolygonvalidiumetrogBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygonvalidiumetrogMetaData.Bin instead.
var PolygonvalidiumetrogBin = PolygonvalidiumetrogMetaData.Bin

// DeployPolygonvalidiumetrog deploys a new Ethereum contract, binding an instance of Polygonvalidiumetrog to it.
func DeployPolygonvalidiumetrog(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address, _rollupManager common.Address) (common.Address, *types.Transaction, *Polygonvalidiumetrog, error) {
	parsed, err := PolygonvalidiumetrogMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygonvalidiumetrogBin), backend, _globalExitRootManager, _pol, _bridgeAddress, _rollupManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonvalidiumetrog{PolygonvalidiumetrogCaller: PolygonvalidiumetrogCaller{contract: contract}, PolygonvalidiumetrogTransactor: PolygonvalidiumetrogTransactor{contract: contract}, PolygonvalidiumetrogFilterer: PolygonvalidiumetrogFilterer{contract: contract}}, nil
}

// Polygonvalidiumetrog is an auto generated Go binding around an Ethereum contract.
type Polygonvalidiumetrog struct {
	PolygonvalidiumetrogCaller     // Read-only binding to the contract
	PolygonvalidiumetrogTransactor // Write-only binding to the contract
	PolygonvalidiumetrogFilterer   // Log filterer for contract events
}

// PolygonvalidiumetrogCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonvalidiumetrogCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonvalidiumetrogTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonvalidiumetrogTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonvalidiumetrogFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonvalidiumetrogFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonvalidiumetrogSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonvalidiumetrogSession struct {
	Contract     *Polygonvalidiumetrog // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PolygonvalidiumetrogCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonvalidiumetrogCallerSession struct {
	Contract *PolygonvalidiumetrogCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// PolygonvalidiumetrogTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonvalidiumetrogTransactorSession struct {
	Contract     *PolygonvalidiumetrogTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// PolygonvalidiumetrogRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonvalidiumetrogRaw struct {
	Contract *Polygonvalidiumetrog // Generic contract binding to access the raw methods on
}

// PolygonvalidiumetrogCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonvalidiumetrogCallerRaw struct {
	Contract *PolygonvalidiumetrogCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonvalidiumetrogTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonvalidiumetrogTransactorRaw struct {
	Contract *PolygonvalidiumetrogTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonvalidiumetrog creates a new instance of Polygonvalidiumetrog, bound to a specific deployed contract.
func NewPolygonvalidiumetrog(address common.Address, backend bind.ContractBackend) (*Polygonvalidiumetrog, error) {
	contract, err := bindPolygonvalidiumetrog(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonvalidiumetrog{PolygonvalidiumetrogCaller: PolygonvalidiumetrogCaller{contract: contract}, PolygonvalidiumetrogTransactor: PolygonvalidiumetrogTransactor{contract: contract}, PolygonvalidiumetrogFilterer: PolygonvalidiumetrogFilterer{contract: contract}}, nil
}

// NewPolygonvalidiumetrogCaller creates a new read-only instance of Polygonvalidiumetrog, bound to a specific deployed contract.
func NewPolygonvalidiumetrogCaller(address common.Address, caller bind.ContractCaller) (*PolygonvalidiumetrogCaller, error) {
	contract, err := bindPolygonvalidiumetrog(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogCaller{contract: contract}, nil
}

// NewPolygonvalidiumetrogTransactor creates a new write-only instance of Polygonvalidiumetrog, bound to a specific deployed contract.
func NewPolygonvalidiumetrogTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonvalidiumetrogTransactor, error) {
	contract, err := bindPolygonvalidiumetrog(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogTransactor{contract: contract}, nil
}

// NewPolygonvalidiumetrogFilterer creates a new log filterer instance of Polygonvalidiumetrog, bound to a specific deployed contract.
func NewPolygonvalidiumetrogFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonvalidiumetrogFilterer, error) {
	contract, err := bindPolygonvalidiumetrog(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogFilterer{contract: contract}, nil
}

// bindPolygonvalidiumetrog binds a generic wrapper to an already deployed contract.
func bindPolygonvalidiumetrog(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonvalidiumetrogMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonvalidiumetrog *PolygonvalidiumetrogRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonvalidiumetrog.Contract.PolygonvalidiumetrogCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonvalidiumetrog *PolygonvalidiumetrogRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.PolygonvalidiumetrogTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonvalidiumetrog *PolygonvalidiumetrogRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.PolygonvalidiumetrogTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonvalidiumetrog.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.contract.Transact(opts, method, params...)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) GLOBALEXITROOTMANAGERL2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "GLOBAL_EXIT_ROOT_MANAGER_L2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _Polygonvalidiumetrog.Contract.GLOBALEXITROOTMANAGERL2(&_Polygonvalidiumetrog.CallOpts)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _Polygonvalidiumetrog.Contract.GLOBALEXITROOTMANAGERL2(&_Polygonvalidiumetrog.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) INITIALIZETXBRIDGELISTLENLEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_LIST_LEN_LEN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _Polygonvalidiumetrog.Contract.INITIALIZETXBRIDGELISTLENLEN(&_Polygonvalidiumetrog.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _Polygonvalidiumetrog.Contract.INITIALIZETXBRIDGELISTLENLEN(&_Polygonvalidiumetrog.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) INITIALIZETXBRIDGEPARAMS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _Polygonvalidiumetrog.Contract.INITIALIZETXBRIDGEPARAMS(&_Polygonvalidiumetrog.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _Polygonvalidiumetrog.Contract.INITIALIZETXBRIDGEPARAMS(&_Polygonvalidiumetrog.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _Polygonvalidiumetrog.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_Polygonvalidiumetrog.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _Polygonvalidiumetrog.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_Polygonvalidiumetrog.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _Polygonvalidiumetrog.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_Polygonvalidiumetrog.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _Polygonvalidiumetrog.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_Polygonvalidiumetrog.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) INITIALIZETXCONSTANTBYTES(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _Polygonvalidiumetrog.Contract.INITIALIZETXCONSTANTBYTES(&_Polygonvalidiumetrog.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _Polygonvalidiumetrog.Contract.INITIALIZETXCONSTANTBYTES(&_Polygonvalidiumetrog.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) INITIALIZETXCONSTANTBYTESEMPTYMETADATA(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _Polygonvalidiumetrog.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_Polygonvalidiumetrog.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _Polygonvalidiumetrog.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_Polygonvalidiumetrog.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) INITIALIZETXDATALENEMPTYMETADATA(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "INITIALIZE_TX_DATA_LEN_EMPTY_METADATA")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _Polygonvalidiumetrog.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_Polygonvalidiumetrog.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _Polygonvalidiumetrog.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_Polygonvalidiumetrog.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) INITIALIZETXEFFECTIVEPERCENTAGE(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "INITIALIZE_TX_EFFECTIVE_PERCENTAGE")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _Polygonvalidiumetrog.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_Polygonvalidiumetrog.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _Polygonvalidiumetrog.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_Polygonvalidiumetrog.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) SIGNATUREINITIALIZETXR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_R")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _Polygonvalidiumetrog.Contract.SIGNATUREINITIALIZETXR(&_Polygonvalidiumetrog.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _Polygonvalidiumetrog.Contract.SIGNATUREINITIALIZETXR(&_Polygonvalidiumetrog.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) SIGNATUREINITIALIZETXS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_S")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _Polygonvalidiumetrog.Contract.SIGNATUREINITIALIZETXS(&_Polygonvalidiumetrog.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _Polygonvalidiumetrog.Contract.SIGNATUREINITIALIZETXS(&_Polygonvalidiumetrog.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) SIGNATUREINITIALIZETXV(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_V")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _Polygonvalidiumetrog.Contract.SIGNATUREINITIALIZETXV(&_Polygonvalidiumetrog.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _Polygonvalidiumetrog.Contract.SIGNATUREINITIALIZETXV(&_Polygonvalidiumetrog.CallOpts)
}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) TIMESTAMPRANGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "TIMESTAMP_RANGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) TIMESTAMPRANGE() (*big.Int, error) {
	return _Polygonvalidiumetrog.Contract.TIMESTAMPRANGE(&_Polygonvalidiumetrog.CallOpts)
}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) TIMESTAMPRANGE() (*big.Int, error) {
	return _Polygonvalidiumetrog.Contract.TIMESTAMPRANGE(&_Polygonvalidiumetrog.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) Admin() (common.Address, error) {
	return _Polygonvalidiumetrog.Contract.Admin(&_Polygonvalidiumetrog.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) Admin() (common.Address, error) {
	return _Polygonvalidiumetrog.Contract.Admin(&_Polygonvalidiumetrog.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) BridgeAddress() (common.Address, error) {
	return _Polygonvalidiumetrog.Contract.BridgeAddress(&_Polygonvalidiumetrog.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) BridgeAddress() (common.Address, error) {
	return _Polygonvalidiumetrog.Contract.BridgeAddress(&_Polygonvalidiumetrog.CallOpts)
}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) CalculatePolPerForceBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "calculatePolPerForceBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) CalculatePolPerForceBatch() (*big.Int, error) {
	return _Polygonvalidiumetrog.Contract.CalculatePolPerForceBatch(&_Polygonvalidiumetrog.CallOpts)
}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) CalculatePolPerForceBatch() (*big.Int, error) {
	return _Polygonvalidiumetrog.Contract.CalculatePolPerForceBatch(&_Polygonvalidiumetrog.CallOpts)
}

// DataAvailabilityProtocol is a free data retrieval call binding the contract method 0xe57a0b4c.
//
// Solidity: function dataAvailabilityProtocol() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) DataAvailabilityProtocol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "dataAvailabilityProtocol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataAvailabilityProtocol is a free data retrieval call binding the contract method 0xe57a0b4c.
//
// Solidity: function dataAvailabilityProtocol() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) DataAvailabilityProtocol() (common.Address, error) {
	return _Polygonvalidiumetrog.Contract.DataAvailabilityProtocol(&_Polygonvalidiumetrog.CallOpts)
}

// DataAvailabilityProtocol is a free data retrieval call binding the contract method 0xe57a0b4c.
//
// Solidity: function dataAvailabilityProtocol() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) DataAvailabilityProtocol() (common.Address, error) {
	return _Polygonvalidiumetrog.Contract.DataAvailabilityProtocol(&_Polygonvalidiumetrog.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) ForceBatchAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "forceBatchAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) ForceBatchAddress() (common.Address, error) {
	return _Polygonvalidiumetrog.Contract.ForceBatchAddress(&_Polygonvalidiumetrog.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) ForceBatchAddress() (common.Address, error) {
	return _Polygonvalidiumetrog.Contract.ForceBatchAddress(&_Polygonvalidiumetrog.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonvalidiumetrog.Contract.ForceBatchTimeout(&_Polygonvalidiumetrog.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonvalidiumetrog.Contract.ForceBatchTimeout(&_Polygonvalidiumetrog.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonvalidiumetrog.Contract.ForcedBatches(&_Polygonvalidiumetrog.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonvalidiumetrog.Contract.ForcedBatches(&_Polygonvalidiumetrog.CallOpts, arg0)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) GasTokenAddress() (common.Address, error) {
	return _Polygonvalidiumetrog.Contract.GasTokenAddress(&_Polygonvalidiumetrog.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) GasTokenAddress() (common.Address, error) {
	return _Polygonvalidiumetrog.Contract.GasTokenAddress(&_Polygonvalidiumetrog.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) GasTokenNetwork() (uint32, error) {
	return _Polygonvalidiumetrog.Contract.GasTokenNetwork(&_Polygonvalidiumetrog.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) GasTokenNetwork() (uint32, error) {
	return _Polygonvalidiumetrog.Contract.GasTokenNetwork(&_Polygonvalidiumetrog.CallOpts)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) GenerateInitializeTransaction(opts *bind.CallOpts, networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "generateInitializeTransaction", networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) GenerateInitializeTransaction(networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _Polygonvalidiumetrog.Contract.GenerateInitializeTransaction(&_Polygonvalidiumetrog.CallOpts, networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) GenerateInitializeTransaction(networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _Polygonvalidiumetrog.Contract.GenerateInitializeTransaction(&_Polygonvalidiumetrog.CallOpts, networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonvalidiumetrog.Contract.GlobalExitRootManager(&_Polygonvalidiumetrog.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonvalidiumetrog.Contract.GlobalExitRootManager(&_Polygonvalidiumetrog.CallOpts)
}

// IsSequenceWithDataAvailabilityAllowed is a free data retrieval call binding the contract method 0x4c21fef3.
//
// Solidity: function isSequenceWithDataAvailabilityAllowed() view returns(bool)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) IsSequenceWithDataAvailabilityAllowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "isSequenceWithDataAvailabilityAllowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSequenceWithDataAvailabilityAllowed is a free data retrieval call binding the contract method 0x4c21fef3.
//
// Solidity: function isSequenceWithDataAvailabilityAllowed() view returns(bool)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) IsSequenceWithDataAvailabilityAllowed() (bool, error) {
	return _Polygonvalidiumetrog.Contract.IsSequenceWithDataAvailabilityAllowed(&_Polygonvalidiumetrog.CallOpts)
}

// IsSequenceWithDataAvailabilityAllowed is a free data retrieval call binding the contract method 0x4c21fef3.
//
// Solidity: function isSequenceWithDataAvailabilityAllowed() view returns(bool)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) IsSequenceWithDataAvailabilityAllowed() (bool, error) {
	return _Polygonvalidiumetrog.Contract.IsSequenceWithDataAvailabilityAllowed(&_Polygonvalidiumetrog.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) LastAccInputHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "lastAccInputHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) LastAccInputHash() ([32]byte, error) {
	return _Polygonvalidiumetrog.Contract.LastAccInputHash(&_Polygonvalidiumetrog.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) LastAccInputHash() ([32]byte, error) {
	return _Polygonvalidiumetrog.Contract.LastAccInputHash(&_Polygonvalidiumetrog.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) LastForceBatch() (uint64, error) {
	return _Polygonvalidiumetrog.Contract.LastForceBatch(&_Polygonvalidiumetrog.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) LastForceBatch() (uint64, error) {
	return _Polygonvalidiumetrog.Contract.LastForceBatch(&_Polygonvalidiumetrog.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonvalidiumetrog.Contract.LastForceBatchSequenced(&_Polygonvalidiumetrog.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonvalidiumetrog.Contract.LastForceBatchSequenced(&_Polygonvalidiumetrog.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) NetworkName() (string, error) {
	return _Polygonvalidiumetrog.Contract.NetworkName(&_Polygonvalidiumetrog.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) NetworkName() (string, error) {
	return _Polygonvalidiumetrog.Contract.NetworkName(&_Polygonvalidiumetrog.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) PendingAdmin() (common.Address, error) {
	return _Polygonvalidiumetrog.Contract.PendingAdmin(&_Polygonvalidiumetrog.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) PendingAdmin() (common.Address, error) {
	return _Polygonvalidiumetrog.Contract.PendingAdmin(&_Polygonvalidiumetrog.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) Pol() (common.Address, error) {
	return _Polygonvalidiumetrog.Contract.Pol(&_Polygonvalidiumetrog.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) Pol() (common.Address, error) {
	return _Polygonvalidiumetrog.Contract.Pol(&_Polygonvalidiumetrog.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) RollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "rollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) RollupManager() (common.Address, error) {
	return _Polygonvalidiumetrog.Contract.RollupManager(&_Polygonvalidiumetrog.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) RollupManager() (common.Address, error) {
	return _Polygonvalidiumetrog.Contract.RollupManager(&_Polygonvalidiumetrog.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) TrustedSequencer() (common.Address, error) {
	return _Polygonvalidiumetrog.Contract.TrustedSequencer(&_Polygonvalidiumetrog.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) TrustedSequencer() (common.Address, error) {
	return _Polygonvalidiumetrog.Contract.TrustedSequencer(&_Polygonvalidiumetrog.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCaller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonvalidiumetrog.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) TrustedSequencerURL() (string, error) {
	return _Polygonvalidiumetrog.Contract.TrustedSequencerURL(&_Polygonvalidiumetrog.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogCallerSession) TrustedSequencerURL() (string, error) {
	return _Polygonvalidiumetrog.Contract.TrustedSequencerURL(&_Polygonvalidiumetrog.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.AcceptAdminRole(&_Polygonvalidiumetrog.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.AcceptAdminRole(&_Polygonvalidiumetrog.TransactOpts)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactor) ForceBatch(opts *bind.TransactOpts, transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.contract.Transact(opts, "forceBatch", transactions, polAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) ForceBatch(transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.ForceBatch(&_Polygonvalidiumetrog.TransactOpts, transactions, polAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactorSession) ForceBatch(transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.ForceBatch(&_Polygonvalidiumetrog.TransactOpts, transactions, polAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.contract.Transact(opts, "initialize", _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.Initialize(&_Polygonvalidiumetrog.TransactOpts, _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactorSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.Initialize(&_Polygonvalidiumetrog.TransactOpts, _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactor) OnVerifyBatches(opts *bind.TransactOpts, lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.contract.Transact(opts, "onVerifyBatches", lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.OnVerifyBatches(&_Polygonvalidiumetrog.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactorSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.OnVerifyBatches(&_Polygonvalidiumetrog.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xdef57e54.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactor) SequenceBatches(opts *bind.TransactOpts, batches []PolygonRollupBaseEtrogBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.contract.Transact(opts, "sequenceBatches", batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xdef57e54.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) SequenceBatches(batches []PolygonRollupBaseEtrogBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.SequenceBatches(&_Polygonvalidiumetrog.TransactOpts, batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xdef57e54.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactorSession) SequenceBatches(batches []PolygonRollupBaseEtrogBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.SequenceBatches(&_Polygonvalidiumetrog.TransactOpts, batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase)
}

// SequenceBatchesValidium is a paid mutator transaction binding the contract method 0xdb5b0ed7.
//
// Solidity: function sequenceBatchesValidium((bytes32,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactor) SequenceBatchesValidium(opts *bind.TransactOpts, batches []PolygonValidiumEtrogValidiumBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.contract.Transact(opts, "sequenceBatchesValidium", batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase, dataAvailabilityMessage)
}

// SequenceBatchesValidium is a paid mutator transaction binding the contract method 0xdb5b0ed7.
//
// Solidity: function sequenceBatchesValidium((bytes32,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) SequenceBatchesValidium(batches []PolygonValidiumEtrogValidiumBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.SequenceBatchesValidium(&_Polygonvalidiumetrog.TransactOpts, batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase, dataAvailabilityMessage)
}

// SequenceBatchesValidium is a paid mutator transaction binding the contract method 0xdb5b0ed7.
//
// Solidity: function sequenceBatchesValidium((bytes32,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactorSession) SequenceBatchesValidium(batches []PolygonValidiumEtrogValidiumBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.SequenceBatchesValidium(&_Polygonvalidiumetrog.TransactOpts, batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase, dataAvailabilityMessage)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactor) SequenceForceBatches(opts *bind.TransactOpts, batches []PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.contract.Transact(opts, "sequenceForceBatches", batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) SequenceForceBatches(batches []PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.SequenceForceBatches(&_Polygonvalidiumetrog.TransactOpts, batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactorSession) SequenceForceBatches(batches []PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.SequenceForceBatches(&_Polygonvalidiumetrog.TransactOpts, batches)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactor) SetDataAvailabilityProtocol(opts *bind.TransactOpts, newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.contract.Transact(opts, "setDataAvailabilityProtocol", newDataAvailabilityProtocol)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) SetDataAvailabilityProtocol(newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.SetDataAvailabilityProtocol(&_Polygonvalidiumetrog.TransactOpts, newDataAvailabilityProtocol)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactorSession) SetDataAvailabilityProtocol(newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.SetDataAvailabilityProtocol(&_Polygonvalidiumetrog.TransactOpts, newDataAvailabilityProtocol)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactor) SetForceBatchAddress(opts *bind.TransactOpts, newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.contract.Transact(opts, "setForceBatchAddress", newForceBatchAddress)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) SetForceBatchAddress(newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.SetForceBatchAddress(&_Polygonvalidiumetrog.TransactOpts, newForceBatchAddress)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactorSession) SetForceBatchAddress(newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.SetForceBatchAddress(&_Polygonvalidiumetrog.TransactOpts, newForceBatchAddress)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactor) SetForceBatchTimeout(opts *bind.TransactOpts, newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.contract.Transact(opts, "setForceBatchTimeout", newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.SetForceBatchTimeout(&_Polygonvalidiumetrog.TransactOpts, newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactorSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.SetForceBatchTimeout(&_Polygonvalidiumetrog.TransactOpts, newforceBatchTimeout)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.SetTrustedSequencer(&_Polygonvalidiumetrog.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.SetTrustedSequencer(&_Polygonvalidiumetrog.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.SetTrustedSequencerURL(&_Polygonvalidiumetrog.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.SetTrustedSequencerURL(&_Polygonvalidiumetrog.TransactOpts, newTrustedSequencerURL)
}

// SwitchSequenceWithDataAvailability is a paid mutator transaction binding the contract method 0x2acdc2b6.
//
// Solidity: function switchSequenceWithDataAvailability(bool newIsSequenceWithDataAvailabilityAllowed) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactor) SwitchSequenceWithDataAvailability(opts *bind.TransactOpts, newIsSequenceWithDataAvailabilityAllowed bool) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.contract.Transact(opts, "switchSequenceWithDataAvailability", newIsSequenceWithDataAvailabilityAllowed)
}

// SwitchSequenceWithDataAvailability is a paid mutator transaction binding the contract method 0x2acdc2b6.
//
// Solidity: function switchSequenceWithDataAvailability(bool newIsSequenceWithDataAvailabilityAllowed) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) SwitchSequenceWithDataAvailability(newIsSequenceWithDataAvailabilityAllowed bool) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.SwitchSequenceWithDataAvailability(&_Polygonvalidiumetrog.TransactOpts, newIsSequenceWithDataAvailabilityAllowed)
}

// SwitchSequenceWithDataAvailability is a paid mutator transaction binding the contract method 0x2acdc2b6.
//
// Solidity: function switchSequenceWithDataAvailability(bool newIsSequenceWithDataAvailabilityAllowed) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactorSession) SwitchSequenceWithDataAvailability(newIsSequenceWithDataAvailabilityAllowed bool) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.SwitchSequenceWithDataAvailability(&_Polygonvalidiumetrog.TransactOpts, newIsSequenceWithDataAvailabilityAllowed)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.TransferAdminRole(&_Polygonvalidiumetrog.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrog.Contract.TransferAdminRole(&_Polygonvalidiumetrog.TransactOpts, newPendingAdmin)
}

// PolygonvalidiumetrogAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogAcceptAdminRoleIterator struct {
	Event *PolygonvalidiumetrogAcceptAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogAcceptAdminRole)
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
		it.Event = new(PolygonvalidiumetrogAcceptAdminRole)
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
func (it *PolygonvalidiumetrogAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogAcceptAdminRole represents a AcceptAdminRole event raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*PolygonvalidiumetrogAcceptAdminRoleIterator, error) {

	logs, sub, err := _Polygonvalidiumetrog.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogAcceptAdminRoleIterator{contract: _Polygonvalidiumetrog.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumetrog.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogAcceptAdminRole)
				if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
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
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) ParseAcceptAdminRole(log types.Log) (*PolygonvalidiumetrogAcceptAdminRole, error) {
	event := new(PolygonvalidiumetrogAcceptAdminRole)
	if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumetrogForceBatchIterator is returned from FilterForceBatch and is used to iterate over the raw logs and unpacked data for ForceBatch events raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogForceBatchIterator struct {
	Event *PolygonvalidiumetrogForceBatch // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogForceBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogForceBatch)
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
		it.Event = new(PolygonvalidiumetrogForceBatch)
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
func (it *PolygonvalidiumetrogForceBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogForceBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogForceBatch represents a ForceBatch event raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogForceBatch struct {
	ForceBatchNum      uint64
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Transactions       []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterForceBatch is a free log retrieval operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) FilterForceBatch(opts *bind.FilterOpts, forceBatchNum []uint64) (*PolygonvalidiumetrogForceBatchIterator, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Polygonvalidiumetrog.contract.FilterLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogForceBatchIterator{contract: _Polygonvalidiumetrog.contract, event: "ForceBatch", logs: logs, sub: sub}, nil
}

// WatchForceBatch is a free log subscription operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) WatchForceBatch(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogForceBatch, forceBatchNum []uint64) (event.Subscription, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Polygonvalidiumetrog.contract.WatchLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogForceBatch)
				if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "ForceBatch", log); err != nil {
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
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) ParseForceBatch(log types.Log) (*PolygonvalidiumetrogForceBatch, error) {
	event := new(PolygonvalidiumetrogForceBatch)
	if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "ForceBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumetrogInitialSequenceBatchesIterator is returned from FilterInitialSequenceBatches and is used to iterate over the raw logs and unpacked data for InitialSequenceBatches events raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogInitialSequenceBatchesIterator struct {
	Event *PolygonvalidiumetrogInitialSequenceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogInitialSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogInitialSequenceBatches)
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
		it.Event = new(PolygonvalidiumetrogInitialSequenceBatches)
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
func (it *PolygonvalidiumetrogInitialSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogInitialSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogInitialSequenceBatches represents a InitialSequenceBatches event raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogInitialSequenceBatches struct {
	Transactions       []byte
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterInitialSequenceBatches is a free log retrieval operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) FilterInitialSequenceBatches(opts *bind.FilterOpts) (*PolygonvalidiumetrogInitialSequenceBatchesIterator, error) {

	logs, sub, err := _Polygonvalidiumetrog.contract.FilterLogs(opts, "InitialSequenceBatches")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogInitialSequenceBatchesIterator{contract: _Polygonvalidiumetrog.contract, event: "InitialSequenceBatches", logs: logs, sub: sub}, nil
}

// WatchInitialSequenceBatches is a free log subscription operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) WatchInitialSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogInitialSequenceBatches) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumetrog.contract.WatchLogs(opts, "InitialSequenceBatches")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogInitialSequenceBatches)
				if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "InitialSequenceBatches", log); err != nil {
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
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) ParseInitialSequenceBatches(log types.Log) (*PolygonvalidiumetrogInitialSequenceBatches, error) {
	event := new(PolygonvalidiumetrogInitialSequenceBatches)
	if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "InitialSequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumetrogInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogInitializedIterator struct {
	Event *PolygonvalidiumetrogInitialized // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogInitialized)
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
		it.Event = new(PolygonvalidiumetrogInitialized)
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
func (it *PolygonvalidiumetrogInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogInitialized represents a Initialized event raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) FilterInitialized(opts *bind.FilterOpts) (*PolygonvalidiumetrogInitializedIterator, error) {

	logs, sub, err := _Polygonvalidiumetrog.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogInitializedIterator{contract: _Polygonvalidiumetrog.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogInitialized) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumetrog.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogInitialized)
				if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) ParseInitialized(log types.Log) (*PolygonvalidiumetrogInitialized, error) {
	event := new(PolygonvalidiumetrogInitialized)
	if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumetrogSequenceBatchesIterator is returned from FilterSequenceBatches and is used to iterate over the raw logs and unpacked data for SequenceBatches events raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogSequenceBatchesIterator struct {
	Event *PolygonvalidiumetrogSequenceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogSequenceBatches)
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
		it.Event = new(PolygonvalidiumetrogSequenceBatches)
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
func (it *PolygonvalidiumetrogSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogSequenceBatches represents a SequenceBatches event raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogSequenceBatches struct {
	NumBatch   uint64
	L1InfoRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSequenceBatches is a free log retrieval operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) FilterSequenceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonvalidiumetrogSequenceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonvalidiumetrog.contract.FilterLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogSequenceBatchesIterator{contract: _Polygonvalidiumetrog.contract, event: "SequenceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceBatches is a free log subscription operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) WatchSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogSequenceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonvalidiumetrog.contract.WatchLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogSequenceBatches)
				if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
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
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) ParseSequenceBatches(log types.Log) (*PolygonvalidiumetrogSequenceBatches, error) {
	event := new(PolygonvalidiumetrogSequenceBatches)
	if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumetrogSequenceForceBatchesIterator is returned from FilterSequenceForceBatches and is used to iterate over the raw logs and unpacked data for SequenceForceBatches events raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogSequenceForceBatchesIterator struct {
	Event *PolygonvalidiumetrogSequenceForceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogSequenceForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogSequenceForceBatches)
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
		it.Event = new(PolygonvalidiumetrogSequenceForceBatches)
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
func (it *PolygonvalidiumetrogSequenceForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogSequenceForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogSequenceForceBatches represents a SequenceForceBatches event raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogSequenceForceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceForceBatches is a free log retrieval operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) FilterSequenceForceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonvalidiumetrogSequenceForceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonvalidiumetrog.contract.FilterLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogSequenceForceBatchesIterator{contract: _Polygonvalidiumetrog.contract, event: "SequenceForceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceForceBatches is a free log subscription operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) WatchSequenceForceBatches(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogSequenceForceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonvalidiumetrog.contract.WatchLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogSequenceForceBatches)
				if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
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
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) ParseSequenceForceBatches(log types.Log) (*PolygonvalidiumetrogSequenceForceBatches, error) {
	event := new(PolygonvalidiumetrogSequenceForceBatches)
	if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumetrogSetDataAvailabilityProtocolIterator is returned from FilterSetDataAvailabilityProtocol and is used to iterate over the raw logs and unpacked data for SetDataAvailabilityProtocol events raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogSetDataAvailabilityProtocolIterator struct {
	Event *PolygonvalidiumetrogSetDataAvailabilityProtocol // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogSetDataAvailabilityProtocolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogSetDataAvailabilityProtocol)
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
		it.Event = new(PolygonvalidiumetrogSetDataAvailabilityProtocol)
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
func (it *PolygonvalidiumetrogSetDataAvailabilityProtocolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogSetDataAvailabilityProtocolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogSetDataAvailabilityProtocol represents a SetDataAvailabilityProtocol event raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogSetDataAvailabilityProtocol struct {
	NewDataAvailabilityProtocol common.Address
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterSetDataAvailabilityProtocol is a free log retrieval operation binding the contract event 0xd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a.
//
// Solidity: event SetDataAvailabilityProtocol(address newDataAvailabilityProtocol)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) FilterSetDataAvailabilityProtocol(opts *bind.FilterOpts) (*PolygonvalidiumetrogSetDataAvailabilityProtocolIterator, error) {

	logs, sub, err := _Polygonvalidiumetrog.contract.FilterLogs(opts, "SetDataAvailabilityProtocol")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogSetDataAvailabilityProtocolIterator{contract: _Polygonvalidiumetrog.contract, event: "SetDataAvailabilityProtocol", logs: logs, sub: sub}, nil
}

// WatchSetDataAvailabilityProtocol is a free log subscription operation binding the contract event 0xd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a.
//
// Solidity: event SetDataAvailabilityProtocol(address newDataAvailabilityProtocol)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) WatchSetDataAvailabilityProtocol(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogSetDataAvailabilityProtocol) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumetrog.contract.WatchLogs(opts, "SetDataAvailabilityProtocol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogSetDataAvailabilityProtocol)
				if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "SetDataAvailabilityProtocol", log); err != nil {
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

// ParseSetDataAvailabilityProtocol is a log parse operation binding the contract event 0xd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a.
//
// Solidity: event SetDataAvailabilityProtocol(address newDataAvailabilityProtocol)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) ParseSetDataAvailabilityProtocol(log types.Log) (*PolygonvalidiumetrogSetDataAvailabilityProtocol, error) {
	event := new(PolygonvalidiumetrogSetDataAvailabilityProtocol)
	if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "SetDataAvailabilityProtocol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumetrogSetForceBatchAddressIterator is returned from FilterSetForceBatchAddress and is used to iterate over the raw logs and unpacked data for SetForceBatchAddress events raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogSetForceBatchAddressIterator struct {
	Event *PolygonvalidiumetrogSetForceBatchAddress // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogSetForceBatchAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogSetForceBatchAddress)
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
		it.Event = new(PolygonvalidiumetrogSetForceBatchAddress)
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
func (it *PolygonvalidiumetrogSetForceBatchAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogSetForceBatchAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogSetForceBatchAddress represents a SetForceBatchAddress event raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogSetForceBatchAddress struct {
	NewForceBatchAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchAddress is a free log retrieval operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) FilterSetForceBatchAddress(opts *bind.FilterOpts) (*PolygonvalidiumetrogSetForceBatchAddressIterator, error) {

	logs, sub, err := _Polygonvalidiumetrog.contract.FilterLogs(opts, "SetForceBatchAddress")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogSetForceBatchAddressIterator{contract: _Polygonvalidiumetrog.contract, event: "SetForceBatchAddress", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchAddress is a free log subscription operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) WatchSetForceBatchAddress(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogSetForceBatchAddress) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumetrog.contract.WatchLogs(opts, "SetForceBatchAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogSetForceBatchAddress)
				if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "SetForceBatchAddress", log); err != nil {
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
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) ParseSetForceBatchAddress(log types.Log) (*PolygonvalidiumetrogSetForceBatchAddress, error) {
	event := new(PolygonvalidiumetrogSetForceBatchAddress)
	if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "SetForceBatchAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumetrogSetForceBatchTimeoutIterator is returned from FilterSetForceBatchTimeout and is used to iterate over the raw logs and unpacked data for SetForceBatchTimeout events raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogSetForceBatchTimeoutIterator struct {
	Event *PolygonvalidiumetrogSetForceBatchTimeout // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogSetForceBatchTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogSetForceBatchTimeout)
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
		it.Event = new(PolygonvalidiumetrogSetForceBatchTimeout)
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
func (it *PolygonvalidiumetrogSetForceBatchTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogSetForceBatchTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogSetForceBatchTimeout represents a SetForceBatchTimeout event raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogSetForceBatchTimeout struct {
	NewforceBatchTimeout uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchTimeout is a free log retrieval operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) FilterSetForceBatchTimeout(opts *bind.FilterOpts) (*PolygonvalidiumetrogSetForceBatchTimeoutIterator, error) {

	logs, sub, err := _Polygonvalidiumetrog.contract.FilterLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogSetForceBatchTimeoutIterator{contract: _Polygonvalidiumetrog.contract, event: "SetForceBatchTimeout", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchTimeout is a free log subscription operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) WatchSetForceBatchTimeout(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogSetForceBatchTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumetrog.contract.WatchLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogSetForceBatchTimeout)
				if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
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
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) ParseSetForceBatchTimeout(log types.Log) (*PolygonvalidiumetrogSetForceBatchTimeout, error) {
	event := new(PolygonvalidiumetrogSetForceBatchTimeout)
	if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumetrogSetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogSetTrustedSequencerIterator struct {
	Event *PolygonvalidiumetrogSetTrustedSequencer // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogSetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogSetTrustedSequencer)
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
		it.Event = new(PolygonvalidiumetrogSetTrustedSequencer)
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
func (it *PolygonvalidiumetrogSetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogSetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogSetTrustedSequencer represents a SetTrustedSequencer event raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogSetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*PolygonvalidiumetrogSetTrustedSequencerIterator, error) {

	logs, sub, err := _Polygonvalidiumetrog.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogSetTrustedSequencerIterator{contract: _Polygonvalidiumetrog.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogSetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumetrog.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogSetTrustedSequencer)
				if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
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
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) ParseSetTrustedSequencer(log types.Log) (*PolygonvalidiumetrogSetTrustedSequencer, error) {
	event := new(PolygonvalidiumetrogSetTrustedSequencer)
	if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumetrogSetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogSetTrustedSequencerURLIterator struct {
	Event *PolygonvalidiumetrogSetTrustedSequencerURL // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogSetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogSetTrustedSequencerURL)
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
		it.Event = new(PolygonvalidiumetrogSetTrustedSequencerURL)
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
func (it *PolygonvalidiumetrogSetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogSetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogSetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogSetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*PolygonvalidiumetrogSetTrustedSequencerURLIterator, error) {

	logs, sub, err := _Polygonvalidiumetrog.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogSetTrustedSequencerURLIterator{contract: _Polygonvalidiumetrog.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogSetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumetrog.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogSetTrustedSequencerURL)
				if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
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
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) ParseSetTrustedSequencerURL(log types.Log) (*PolygonvalidiumetrogSetTrustedSequencerURL, error) {
	event := new(PolygonvalidiumetrogSetTrustedSequencerURL)
	if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumetrogSwitchSequenceWithDataAvailabilityIterator is returned from FilterSwitchSequenceWithDataAvailability and is used to iterate over the raw logs and unpacked data for SwitchSequenceWithDataAvailability events raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogSwitchSequenceWithDataAvailabilityIterator struct {
	Event *PolygonvalidiumetrogSwitchSequenceWithDataAvailability // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogSwitchSequenceWithDataAvailabilityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogSwitchSequenceWithDataAvailability)
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
		it.Event = new(PolygonvalidiumetrogSwitchSequenceWithDataAvailability)
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
func (it *PolygonvalidiumetrogSwitchSequenceWithDataAvailabilityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogSwitchSequenceWithDataAvailabilityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogSwitchSequenceWithDataAvailability represents a SwitchSequenceWithDataAvailability event raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogSwitchSequenceWithDataAvailability struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSwitchSequenceWithDataAvailability is a free log retrieval operation binding the contract event 0xf32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41.
//
// Solidity: event SwitchSequenceWithDataAvailability()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) FilterSwitchSequenceWithDataAvailability(opts *bind.FilterOpts) (*PolygonvalidiumetrogSwitchSequenceWithDataAvailabilityIterator, error) {

	logs, sub, err := _Polygonvalidiumetrog.contract.FilterLogs(opts, "SwitchSequenceWithDataAvailability")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogSwitchSequenceWithDataAvailabilityIterator{contract: _Polygonvalidiumetrog.contract, event: "SwitchSequenceWithDataAvailability", logs: logs, sub: sub}, nil
}

// WatchSwitchSequenceWithDataAvailability is a free log subscription operation binding the contract event 0xf32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41.
//
// Solidity: event SwitchSequenceWithDataAvailability()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) WatchSwitchSequenceWithDataAvailability(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogSwitchSequenceWithDataAvailability) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumetrog.contract.WatchLogs(opts, "SwitchSequenceWithDataAvailability")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogSwitchSequenceWithDataAvailability)
				if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "SwitchSequenceWithDataAvailability", log); err != nil {
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

// ParseSwitchSequenceWithDataAvailability is a log parse operation binding the contract event 0xf32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41.
//
// Solidity: event SwitchSequenceWithDataAvailability()
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) ParseSwitchSequenceWithDataAvailability(log types.Log) (*PolygonvalidiumetrogSwitchSequenceWithDataAvailability, error) {
	event := new(PolygonvalidiumetrogSwitchSequenceWithDataAvailability)
	if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "SwitchSequenceWithDataAvailability", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumetrogTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogTransferAdminRoleIterator struct {
	Event *PolygonvalidiumetrogTransferAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogTransferAdminRole)
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
		it.Event = new(PolygonvalidiumetrogTransferAdminRole)
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
func (it *PolygonvalidiumetrogTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogTransferAdminRole represents a TransferAdminRole event raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*PolygonvalidiumetrogTransferAdminRoleIterator, error) {

	logs, sub, err := _Polygonvalidiumetrog.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogTransferAdminRoleIterator{contract: _Polygonvalidiumetrog.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumetrog.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogTransferAdminRole)
				if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
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
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) ParseTransferAdminRole(log types.Log) (*PolygonvalidiumetrogTransferAdminRole, error) {
	event := new(PolygonvalidiumetrogTransferAdminRole)
	if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumetrogVerifyBatchesIterator is returned from FilterVerifyBatches and is used to iterate over the raw logs and unpacked data for VerifyBatches events raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogVerifyBatchesIterator struct {
	Event *PolygonvalidiumetrogVerifyBatches // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogVerifyBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogVerifyBatches)
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
		it.Event = new(PolygonvalidiumetrogVerifyBatches)
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
func (it *PolygonvalidiumetrogVerifyBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogVerifyBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogVerifyBatches represents a VerifyBatches event raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumetrogVerifyBatches struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatches is a free log retrieval operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) FilterVerifyBatches(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*PolygonvalidiumetrogVerifyBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonvalidiumetrog.contract.FilterLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogVerifyBatchesIterator{contract: _Polygonvalidiumetrog.contract, event: "VerifyBatches", logs: logs, sub: sub}, nil
}

// WatchVerifyBatches is a free log subscription operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) WatchVerifyBatches(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogVerifyBatches, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonvalidiumetrog.contract.WatchLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogVerifyBatches)
				if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
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
func (_Polygonvalidiumetrog *PolygonvalidiumetrogFilterer) ParseVerifyBatches(log types.Log) (*PolygonvalidiumetrogVerifyBatches, error) {
	event := new(PolygonvalidiumetrogVerifyBatches)
	if err := _Polygonvalidiumetrog.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonvalidiumstoragemigration

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

// PolygonRollupBaseEtrogNoGapBatchData is an auto generated low-level Go binding around an user-defined struct.
type PolygonRollupBaseEtrogNoGapBatchData struct {
	Transactions         []byte
	ForcedGlobalExitRoot [32]byte
	ForcedTimestamp      uint64
	ForcedBlockHashL1    [32]byte
}

// PolygonValidiumStorageMigrationValidiumBatchData is an auto generated low-level Go binding around an user-defined struct.
type PolygonValidiumStorageMigrationValidiumBatchData struct {
	TransactionsHash     [32]byte
	ForcedGlobalExitRoot [32]byte
	ForcedTimestamp      uint64
	ForcedBlockHashL1    [32]byte
}

// PolygonvalidiumstoragemigrationMetaData contains all meta data concerning the Polygonvalidiumstoragemigration contract.
var PolygonvalidiumstoragemigrationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"_rollupManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceWithDataAvailabilityNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwitchToSameValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"forceBatchNum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"}],\"name\":\"ForceBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"}],\"name\":\"InitialSequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"l1InfoRoot\",\"type\":\"bytes32\"}],\"name\":\"SequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newDataAvailabilityProtocol\",\"type\":\"address\"}],\"name\":\"SetDataAvailabilityProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newForceBatchAddress\",\"type\":\"address\"}],\"name\":\"SetForceBatchAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"SetForceBatchTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SwitchSequenceWithDataAvailability\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatches\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GLOBAL_EXIT_ROOT_MANAGER_L2\",\"outputs\":[{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_LIST_LEN_LEN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_DATA_LEN_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_EFFECTIVE_PERCENTAGE\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_R\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_S\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_V\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TIMESTAMP_RANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculatePolPerForceBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataAvailabilityProtocol\",\"outputs\":[{\"internalType\":\"contractIDataAvailabilityProtocol\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"polAmount\",\"type\":\"uint256\"}],\"name\":\"forceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"}],\"name\":\"generateInitializeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeMigration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSequenceWithDataAvailabilityAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"onVerifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrogNoGap.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"maxSequenceTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initSequencedBatch\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"}],\"name\":\"sequenceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonValidiumStorageMigration.ValidiumBatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"maxSequenceTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initSequencedBatch\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataAvailabilityMessage\",\"type\":\"bytes\"}],\"name\":\"sequenceBatchesValidium\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrogNoGap.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"}],\"name\":\"sequenceForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIDataAvailabilityProtocol\",\"name\":\"newDataAvailabilityProtocol\",\"type\":\"address\"}],\"name\":\"setDataAvailabilityProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newForceBatchAddress\",\"type\":\"address\"}],\"name\":\"setForceBatchAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"setForceBatchTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"newIsSequenceWithDataAvailabilityAllowed\",\"type\":\"bool\"}],\"name\":\"switchSequenceWithDataAvailability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b506040516200542f3803806200542f833981016040819052620000359162000071565b6001600160a01b0393841660a052918316608052821660c0521660e052620000d9565b6001600160a01b03811681146200006e57600080fd5b50565b600080600080608085870312156200008857600080fd5b8451620000958162000058565b6020860151909450620000a88162000058565b6040860151909350620000bb8162000058565b6060860151909250620000ce8162000058565b939692955090935050565b60805160a05160c05160e05161523c620001f36000396000818161056001528181610a4901528181610d8a01528181610ef70152818161113b0152818161171c01528181611c8a015281816120e1015281816121d701528181612ddb01528181612e5a01528181612e7c0152818161301f0152818161322d015281816132f501528181613da101528181613e1a01528181613e3c0152613ee70152600081816107020152818161132f01528181611409015281816123a8015281816124b00152818161290201526138de0152600081816107be0152818161159001528181612984015281816134400152613960015260008181610816015281816108f50152818161212a01528181612f2a0152613414015261523c6000f3fe608060405234801561001057600080fd5b506004361061032a5760003560e01c806371257022116101b2578063c7fffd4b116100f9578063def57e54116100a2578063e7a7ed021161007c578063e7a7ed0214610858578063eaeb077b1461086c578063f35dda471461087f578063f851a4401461088757600080fd5b8063def57e54146107fe578063e46761c414610811578063e57a0b4c1461083857600080fd5b8063d02103ca116100d3578063d02103ca146107b9578063d7bc90ff146107e0578063db5b0ed7146107eb57600080fd5b8063c7fffd4b1461077e578063c89e42df14610786578063cfa8ed471461079957600080fd5b80639f26f8401161015b578063ada8f91911610135578063ada8f91914610737578063b0afe1541461074a578063c754c7ed1461075657600080fd5b80639f26f840146106ea578063a3c573eb146106fd578063a652f26c1461072457600080fd5b80638c3d73011161018c5780638c3d7301146106b457806391cafe32146106bc5780639e001877146106cf57600080fd5b806371257022146106525780637a5460c5146106655780637cd76b8b146106a157600080fd5b806340b5de6c1161027657806352bdeb6d1161021f5780636b8616ce116101f95780636b8616ce146106165780636e05d2cd146106365780636ff512cc1461063f57600080fd5b806352bdeb6d146105ca578063542028d514610606578063676870d21461060e57600080fd5b806349b7b8021161025057806349b7b8021461055b5780634c21fef3146105825780634e487706146105b757600080fd5b806340b5de6c146104c257806342308fab1461051a578063456052671461052257600080fd5b806326782247116102d857806332c2d153116102b257806332c2d153146104525780633c351e10146104655780633cbc795b1461048557600080fd5b806326782247146103da5780632acdc2b61461041f5780632c111c061461043257600080fd5b8063107bf28c11610309578063107bf28c146103ae57806311e892d4146103b65780631c8b9370146103d057600080fd5b8062d0295d1461032f578063035089631461034a57806305835f3714610365575b600080fd5b6103376108ad565b6040519081526020015b60405180910390f35b610352602081565b60405161ffff9091168152602001610341565b6103a16040518060400160405280600881526020017f80808401c9c3809400000000000000000000000000000000000000000000000081525081565b60405161034191906142fe565b6103a16109b9565b6103be60f981565b60405160ff9091168152602001610341565b6103d8610a47565b005b6001546103fa9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610341565b6103d861042d366004614329565b610c68565b6008546103fa9073ffffffffffffffffffffffffffffffffffffffff1681565b6103d8610460366004614399565b610d88565b6009546103fa9073ffffffffffffffffffffffffffffffffffffffff1681565b6009546104ad9074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff9091168152602001610341565b6104e97fff0000000000000000000000000000000000000000000000000000000000000081565b6040517fff000000000000000000000000000000000000000000000000000000000000009091168152602001610341565b610337602481565b6007546105429068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff9091168152602001610341565b6103fa7f000000000000000000000000000000000000000000000000000000000000000081565b603c546105a79074010000000000000000000000000000000000000000900460ff1681565b6040519015158152602001610341565b6103d86105c53660046143db565b610e57565b6103a16040518060400160405280600281526020017f80b800000000000000000000000000000000000000000000000000000000000081525081565b6103a1611062565b610352601f81565b6103376106243660046143db565b60066020526000908152604090205481565b61033760055481565b6103d861064d3660046143f8565b61106f565b6103d861066036600461456a565b611139565b6103a16040518060400160405280600281526020017f80b900000000000000000000000000000000000000000000000000000000000081525081565b6103d86106af3660046143f8565b611958565b6103d8611a22565b6103d86106ca3660046143f8565b611af5565b6103fa73a40d5f56745a118d0906a34e69aec8c0db1cb8fa81565b6103d86106f8366004614663565b611c0e565b6103fa7f000000000000000000000000000000000000000000000000000000000000000081565b6103a16107323660046146a5565b6122a7565b6103d86107453660046143f8565b61268c565b6103376405ca1ab1e081565b60075461054290700100000000000000000000000000000000900467ffffffffffffffff1681565b6103be60e481565b6103d861079436600461471a565b612756565b6002546103fa9073ffffffffffffffffffffffffffffffffffffffff1681565b6103fa7f000000000000000000000000000000000000000000000000000000000000000081565b610337635ca1ab1e81565b6103d86107f9366004614791565b6127e9565b6103d861080c366004614864565b61314e565b6103fa7f000000000000000000000000000000000000000000000000000000000000000081565b603c546103fa9073ffffffffffffffffffffffffffffffffffffffff1681565b6007546105429067ffffffffffffffff1681565b6103d861087a3660046148e1565b6131b6565b6103be601b81565b6000546103fa9062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b6040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152600090819073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa15801561093c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610960919061492d565b60075490915060009061098b9067ffffffffffffffff68010000000000000000820481169116614975565b67ffffffffffffffff169050806000036109a85760009250505090565b6109b2818361499d565b9250505090565b600480546109c6906149d8565b80601f01602080910402602001604051908101604052809291908181526020018280546109f2906149d8565b8015610a3f5780601f10610a1457610100808354040283529160200191610a3f565b820191906000526020600020905b815481529060010190602001808311610a2257829003601f168201915b505050505081565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314610ab6576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600054600290610100900460ff16158015610ad8575060005460ff8083169116105b610b69576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b60008054600a8054603c80547401000000000000000000000000000000000000000080840460ff90811615159091027fffffffffffffffffffffff00000000000000000000000000000000000000000092831673ffffffffffffffffffffffffffffffffffffffff861617179092559091169091557fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000909116908316908117610100177fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a150565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314610cbf576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603c5474010000000000000000000000000000000000000000900460ff16151581151503610d19576040517f5f0e7abe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603c80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1674010000000000000000000000000000000000000000831515021790556040517ff32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f4190600090a150565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314610df7576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff168367ffffffffffffffff167f9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f596684604051610e4a91815260200190565b60405180910390a3505050565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314610eae576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff82161115610ef5576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166315064c966040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f60573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f849190614a2b565b610fe55760075467ffffffffffffffff700100000000000000000000000000000000909104811690821610610fe5576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600780547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527fa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b90602001610c5d565b600380546109c6906149d8565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff1633146110c6576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc090602001610c5d565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1633146111a8576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600054610100900460ff16158080156111c85750600054600160ff909116105b806111e25750303b1580156111e2575060005460ff166001145b61126e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610b60565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156112cc57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b606073ffffffffffffffffffffffffffffffffffffffff851615611531576040517fc00f14ab00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063c00f14ab90602401600060405180830381865afa158015611376573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526113bc9190810190614a48565b6040517f318aee3d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff878116600483015291925060009182917f00000000000000000000000000000000000000000000000000000000000000009091169063318aee3d906024016040805180830381865afa158015611451573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114759190614abf565b915091508163ffffffff166000146114ed576009805463ffffffff841674010000000000000000000000000000000000000000027fffffffffffffffff00000000000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff84161717905561152e565b600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff89161790555b50505b60095460009061157990889073ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900463ffffffff16856122a7565b9050600081805190602001209050600042905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061161d919061492d565b90506000808483858f611631600143614af9565b60408051602081019790975286019490945260608086019390935260c09190911b7fffffffffffffffff000000000000000000000000000000000000000000000000166080850152901b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016608883015240609c82015260bc01604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001815290829052805160209091012060058190557f9a908e73000000000000000000000000000000000000000000000000000000008252600160048301526024820181905291507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303816000875af115801561177a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061179e9190614b12565b508c600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508b600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555088600390816118309190614b75565b50600461183d8982614b75565b508c600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062069780600760106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507f060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f85838e6040516118de93929190614c8f565b60405180910390a1505050505050801561194f57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff1633146119af576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603c80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a90602001610c5d565b60015473ffffffffffffffffffffffffffffffffffffffff163314611a73576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600154600080547fffffffffffffffffffff0000000000000000000000000000000000000000ffff1673ffffffffffffffffffffffffffffffffffffffff9092166201000081029290921790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e9060200160405180910390a1565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314611b4c576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60085473ffffffffffffffffffffffffffffffffffffffff16611b9b576040517fc89374d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb90602001610c5d565b60085473ffffffffffffffffffffffffffffffffffffffff168015801590611c4c575073ffffffffffffffffffffffffffffffffffffffff81163314155b15611c83576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b4262093a807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166330c27dde6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611cf3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d179190614b12565b611d219190614cce565b67ffffffffffffffff161115611d63576040517f3d49ed4c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b816000819003611d9f576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8811115611ddb576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075467ffffffffffffffff80821691611e0391849168010000000000000000900416614cef565b1115611e3b576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6007546005546801000000000000000090910467ffffffffffffffff169060005b838110156120db576000878783818110611e7857611e78614d02565b9050602002810190611e8a9190614d31565b611e9390614d6f565b905083611e9f81614ddd565b825180516020918201208185015160408087015160608801519151959a50929550600094611f0c948794929101938452602084019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166040830152604882015260680190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff8916600090815260069093529120549091508114611f95576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8616600090815260066020526040812055611fba600188614af9565b84036120295742600760109054906101000a900467ffffffffffffffff168460400151611fe79190614cce565b67ffffffffffffffff161115612029576040517fc44a082100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208381015160408086015160608088015183519586018b90529285018790528481019390935260c01b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808401523390911b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088830152609c82015260bc0160405160208183030381529060405280519060200120945050505080806120d390614e04565b915050611e5c565b506121517f0000000000000000000000000000000000000000000000000000000000000000846121096108ad565b6121139190614e3c565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169190613688565b60058190556007805467ffffffffffffffff841668010000000000000000027fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff9091161790556040517f9a908e7300000000000000000000000000000000000000000000000000000000815260009073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690639a908e7390612223908790869060040167ffffffffffffffff929092168252602082015260400190565b6020604051808303816000875af1158015612242573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122669190614b12565b60405190915067ffffffffffffffff8216907f648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a490600090a250505050505050565b6060600085858573a40d5f56745a118d0906a34e69aec8c0db1cb8fa6000876040516024016122db96959493929190614e53565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff811bff700000000000000000000000000000000000000000000000000000000179052835190915060609060000361242c5760f9601f83516123709190614eb6565b6040518060400160405280600881526020017f80808401c9c380940000000000000000000000000000000000000000000000008152507f00000000000000000000000000000000000000000000000000000000000000006040518060400160405280600281526020017f80b800000000000000000000000000000000000000000000000000000000000081525060e4876040516020016124169796959493929190614ed1565b6040516020818303038152906040529050612530565b815161ffff1015612469576040517f248b8f8200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815160f9612478602083614eb6565b6040518060400160405280600881526020017f80808401c9c380940000000000000000000000000000000000000000000000008152507f00000000000000000000000000000000000000000000000000000000000000006040518060400160405280600281526020017f80b9000000000000000000000000000000000000000000000000000000000000815250858860405160200161251d9796959493929190614fb4565b6040516020818303038152906040529150505b805160208083019190912060408051600080825293810180835292909252601b908201526405ca1ab1e06060820152635ca1ab1e608082015260019060a0016020604051602081039080840390855afa158015612591573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116612609576040517fcd16196600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60405160009061264f9084906405ca1ab1e090635ca1ab1e90601b907fff0000000000000000000000000000000000000000000000000000000000000090602001615097565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529450505050505b949350505050565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff1633146126e3576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001610c5d565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff1633146127ad576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60036127b98282614b75565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b2081604051610c5d91906142fe565b60025473ffffffffffffffffffffffffffffffffffffffff16331461283a576040517f11e7be1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b856000819003612876576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e88111156128b2576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6128bd602442614cef565b8667ffffffffffffffff161115612900576040517f0a00feb300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379e2cf976040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561296857600080fd5b505af115801561297c573d6000803e3d6000fd5b5050505060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635ca1e1656040518163ffffffff1660e01b8152600401602060405180830381865afa1580156129ed573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a11919061492d565b60075460055491925068010000000000000000900467ffffffffffffffff1690816000805b86811015612d4d5760008e8e83818110612a5257612a52614d02565b905060800201803603810190612a6891906150f3565b604081015190915067ffffffffffffffff1615612c5c5785612a8981614ddd565b96505060008160000151826020015183604001518460600151604051602001612af09493929190938452602084019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166040830152604882015260680190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff8a16600090815260069093529120549091508114612b79576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b858260000151836020015184604001518f8660600151604051602001612c13969594939291909586526020860194909452604085019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166060808501919091521b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166068830152607c820152609c0190565b604051602081830303815290604052805190602001209550600660008867ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206000905550612d3a565b8051604051612c78918591602001918252602082015260400190565b604051602081830303815290604052805190602001209250848160000151888f8e6000801b604051602001612d21969594939291909586526020860194909452604085019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166060808501919091521b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166068830152607c820152609c0190565b6040516020818303038152906040528051906020012094505b5080612d4581614e04565b915050612a36565b5060075467ffffffffffffffff9081169085161115612d98576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60058390558567ffffffffffffffff85811690841614612e4e576000612dbe8487614975565b9050612dd467ffffffffffffffff821683614af9565b9150612e0d7f00000000000000000000000000000000000000000000000000000000000000008267ffffffffffffffff166121096108ad565b50600780547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8816021790555b8015612fdd57612f52337f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663477fa2706040518163ffffffff1660e01b8152600401602060405180830381865afa158015612ee5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f09919061492d565b612f139190614e3c565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016929190613761565b603c546040517f3b51be4b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690633b51be4b90612fac9085908d908d9060040161518a565b60006040518083038186803b158015612fc457600080fd5b505afa158015612fd8573d6000803e3d6000fd5b505050505b6040517f9a908e7300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff88166004820152602481018590526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303816000875af115801561307d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130a19190614b12565b90506130ad8882614975565b67ffffffffffffffff168c67ffffffffffffffff16146130f9576040517f1a070d9a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff167f3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e7668860405161313591815260200190565b60405180910390a2505050505050505050505050505050565b603c5474010000000000000000000000000000000000000000900460ff166131a2576040517f821935b400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6131af85858585856137c5565b5050505050565b60085473ffffffffffffffffffffffffffffffffffffffff1680158015906131f4575073ffffffffffffffffffffffffffffffffffffffff81163314155b1561322b576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166315064c966040518163ffffffff1660e01b8152600401602060405180830381865afa158015613296573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132ba9190614a2b565b156132f1576040517f39258d1800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663604691696040518163ffffffff1660e01b8152600401602060405180830381865afa15801561335e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613382919061492d565b9050828111156133be576040517f2354600f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113888411156133fa576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61343c73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016333084613761565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156134a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134cd919061492d565b6007805491925067ffffffffffffffff9091169060006134ec83614ddd565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505085856040516135239291906151a4565b6040519081900390208142613539600143614af9565b60408051602081019590955284019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166060830152406068820152608801604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012060075467ffffffffffffffff166000908152600690935291205532330361363157600754604080518381523360208201526060818301819052600090820152905167ffffffffffffffff909216917ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319181900360800190a2613680565b60075460405167ffffffffffffffff909116907ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319061367790849033908b908b906151b4565b60405180910390a25b505050505050565b60405173ffffffffffffffffffffffffffffffffffffffff831660248201526044810182905261375c9084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152614013565b505050565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526137bf9085907f23b872dd00000000000000000000000000000000000000000000000000000000906084016136da565b50505050565b60025473ffffffffffffffffffffffffffffffffffffffff163314613816576040517f11e7be1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b836000819003613852576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e881111561388e576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613899602442614cef565b8467ffffffffffffffff1611156138dc576040517f0a00feb300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379e2cf976040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561394457600080fd5b505af1158015613958573d6000803e3d6000fd5b5050505060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635ca1e1656040518163ffffffff1660e01b8152600401602060405180830381865afa1580156139c9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139ed919061492d565b60075460055491925068010000000000000000900467ffffffffffffffff16908160005b85811015613d135760008b8b83818110613a2d57613a2d614d02565b9050602002810190613a3f9190614d31565b613a4890614d6f565b8051805160209091012060408201519192509067ffffffffffffffff1615613c2d5785613a7481614ddd565b965050600081836020015184604001518560600151604051602001613ad79493929190938452602084019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166040830152604882015260680190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff8a16600090815260069093529120549091508114613b60576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208381015160408086015160608088015183519586018c90529285018790528481019390935260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166080840152908c901b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088830152609c82015260bc01604051602081830303815290604052805190602001209550600660008867ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206000905550613cfe565b8151516201d4c01015613c6c576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160208101879052908101829052606080820189905260c08d901b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808301528a901b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660888201526000609c82015260bc016040516020818303038152906040528051906020012094505b50508080613d0b90614e04565b915050613a11565b5060075467ffffffffffffffff9081169084161115613d5e576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60058290558467ffffffffffffffff84811690831614613e14576000613d848386614975565b9050613d9a67ffffffffffffffff821683614af9565b9150613dd37f00000000000000000000000000000000000000000000000000000000000000008267ffffffffffffffff166121096108ad565b50600780547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8716021790555b613ea5337f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663477fa2706040518163ffffffff1660e01b8152600401602060405180830381865afa158015612ee5573d6000803e3d6000fd5b6040517f9a908e7300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff87166004820152602481018490526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303816000875af1158015613f45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613f699190614b12565b9050613f758782614975565b67ffffffffffffffff168967ffffffffffffffff1614613fc1576040517f1a070d9a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff167f3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e76687604051613ffd91815260200190565b60405180910390a2505050505050505050505050565b6000614075826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661411f9092919063ffffffff16565b80519091501561375c57808060200190518101906140939190614a2b565b61375c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610b60565b60606126848484600085856000808673ffffffffffffffffffffffffffffffffffffffff16858760405161415391906151f4565b60006040518083038185875af1925050503d8060008114614190576040519150601f19603f3d011682016040523d82523d6000602084013e614195565b606091505b50915091506141a6878383876141b1565b979650505050505050565b606083156142475782516000036142405773ffffffffffffffffffffffffffffffffffffffff85163b614240576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610b60565b5081612684565b612684838381511561425c5781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b6091906142fe565b60005b838110156142ab578181015183820152602001614293565b50506000910152565b600081518084526142cc816020860160208601614290565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061431160208301846142b4565b9392505050565b801515811461432657600080fd5b50565b60006020828403121561433b57600080fd5b813561431181614318565b67ffffffffffffffff8116811461432657600080fd5b803561436781614346565b919050565b73ffffffffffffffffffffffffffffffffffffffff8116811461432657600080fd5b80356143678161436c565b6000806000606084860312156143ae57600080fd5b83356143b981614346565b92506020840135915060408401356143d08161436c565b809150509250925092565b6000602082840312156143ed57600080fd5b813561431181614346565b60006020828403121561440a57600080fd5b81356143118161436c565b63ffffffff8116811461432657600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516080810167ffffffffffffffff8111828210171561447957614479614427565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156144c6576144c6614427565b604052919050565b600067ffffffffffffffff8211156144e8576144e8614427565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f83011261452557600080fd5b8135614538614533826144ce565b61447f565b81815284602083860101111561454d57600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060008060c0878903121561458357600080fd5b863561458e8161436c565b9550602087013561459e8161436c565b945060408701356145ae81614415565b935060608701356145be8161436c565b9250608087013567ffffffffffffffff808211156145db57600080fd5b6145e78a838b01614514565b935060a08901359150808211156145fd57600080fd5b5061460a89828a01614514565b9150509295509295509295565b60008083601f84011261462957600080fd5b50813567ffffffffffffffff81111561464157600080fd5b6020830191508360208260051b850101111561465c57600080fd5b9250929050565b6000806020838503121561467657600080fd5b823567ffffffffffffffff81111561468d57600080fd5b61469985828601614617565b90969095509350505050565b600080600080608085870312156146bb57600080fd5b84356146c681614415565b935060208501356146d68161436c565b925060408501356146e681614415565b9150606085013567ffffffffffffffff81111561470257600080fd5b61470e87828801614514565b91505092959194509250565b60006020828403121561472c57600080fd5b813567ffffffffffffffff81111561474357600080fd5b61268484828501614514565b60008083601f84011261476157600080fd5b50813567ffffffffffffffff81111561477957600080fd5b60208301915083602082850101111561465c57600080fd5b600080600080600080600060a0888a0312156147ac57600080fd5b873567ffffffffffffffff808211156147c457600080fd5b818a0191508a601f8301126147d857600080fd5b8135818111156147e757600080fd5b8b60208260071b85010111156147fc57600080fd5b6020830199508098505061481260208b0161435c565b965061482060408b0161435c565b955061482e60608b0161438e565b945060808a013591508082111561484457600080fd5b506148518a828b0161474f565b989b979a50959850939692959293505050565b60008060008060006080868803121561487c57600080fd5b853567ffffffffffffffff81111561489357600080fd5b61489f88828901614617565b90965094505060208601356148b381614346565b925060408601356148c381614346565b915060608601356148d38161436c565b809150509295509295909350565b6000806000604084860312156148f657600080fd5b833567ffffffffffffffff81111561490d57600080fd5b6149198682870161474f565b909790965060209590950135949350505050565b60006020828403121561493f57600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b67ffffffffffffffff82811682821603908082111561499657614996614946565b5092915050565b6000826149d3577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b600181811c908216806149ec57607f821691505b602082108103614a25577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b600060208284031215614a3d57600080fd5b815161431181614318565b600060208284031215614a5a57600080fd5b815167ffffffffffffffff811115614a7157600080fd5b8201601f81018413614a8257600080fd5b8051614a90614533826144ce565b818152856020838501011115614aa557600080fd5b614ab6826020830160208601614290565b95945050505050565b60008060408385031215614ad257600080fd5b8251614add81614415565b6020840151909250614aee8161436c565b809150509250929050565b81810381811115614b0c57614b0c614946565b92915050565b600060208284031215614b2457600080fd5b815161431181614346565b601f82111561375c57600081815260208120601f850160051c81016020861015614b565750805b601f850160051c820191505b8181101561368057828155600101614b62565b815167ffffffffffffffff811115614b8f57614b8f614427565b614ba381614b9d84546149d8565b84614b2f565b602080601f831160018114614bf65760008415614bc05750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555613680565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015614c4357888601518255948401946001909101908401614c24565b5085821015614c7f57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b606081526000614ca260608301866142b4565b905083602083015273ffffffffffffffffffffffffffffffffffffffff83166040830152949350505050565b67ffffffffffffffff81811683821601908082111561499657614996614946565b80820180821115614b0c57614b0c614946565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81833603018112614d6557600080fd5b9190910192915050565b600060808236031215614d8157600080fd5b614d89614456565b823567ffffffffffffffff811115614da057600080fd5b614dac36828601614514565b825250602083013560208201526040830135614dc781614346565b6040820152606092830135928101929092525090565b600067ffffffffffffffff808316818103614dfa57614dfa614946565b6001019392505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614e3557614e35614946565b5060010190565b8082028115828204841417614b0c57614b0c614946565b600063ffffffff808916835273ffffffffffffffffffffffffffffffffffffffff8089166020850152818816604085015280871660608501528086166080850152505060c060a0830152614eaa60c08301846142b4565b98975050505050505050565b61ffff81811683821601908082111561499657614996614946565b60007fff00000000000000000000000000000000000000000000000000000000000000808a60f81b1683527fffff0000000000000000000000000000000000000000000000000000000000008960f01b1660018401528751614f3a816003860160208c01614290565b80840190507fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008860601b1660038201528651614f7d816017840160208b01614290565b808201915050818660f81b16601782015284519150614fa3826018830160208801614290565b016018019998505050505050505050565b7fff000000000000000000000000000000000000000000000000000000000000008860f81b16815260007fffff000000000000000000000000000000000000000000000000000000000000808960f01b166001840152875161501d816003860160208c01614290565b80840190507fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008860601b1660038201528651615060816017840160208b01614290565b808201915050818660f01b16601782015284519150615086826019830160208801614290565b016019019998505050505050505050565b600086516150a9818460208b01614290565b9190910194855250602084019290925260f81b7fff000000000000000000000000000000000000000000000000000000000000009081166040840152166041820152604201919050565b60006080828403121561510557600080fd5b61510d614456565b8235815260208301356020820152604083013561512981614346565b60408201526060928301359281019290925250919050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b838152604060208201526000614ab6604083018486615141565b8183823760009101908152919050565b84815273ffffffffffffffffffffffffffffffffffffffff841660208201526060604082015260006151ea606083018486615141565b9695505050505050565b60008251614d6581846020870161429056fea264697066735822122003471585eb7894d571ed4992aac40bbb3809b50a024d5167301eccc4adba852964736f6c63430008140033",
}

// PolygonvalidiumstoragemigrationABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonvalidiumstoragemigrationMetaData.ABI instead.
var PolygonvalidiumstoragemigrationABI = PolygonvalidiumstoragemigrationMetaData.ABI

// PolygonvalidiumstoragemigrationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygonvalidiumstoragemigrationMetaData.Bin instead.
var PolygonvalidiumstoragemigrationBin = PolygonvalidiumstoragemigrationMetaData.Bin

// DeployPolygonvalidiumstoragemigration deploys a new Ethereum contract, binding an instance of Polygonvalidiumstoragemigration to it.
func DeployPolygonvalidiumstoragemigration(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address, _rollupManager common.Address) (common.Address, *types.Transaction, *Polygonvalidiumstoragemigration, error) {
	parsed, err := PolygonvalidiumstoragemigrationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygonvalidiumstoragemigrationBin), backend, _globalExitRootManager, _pol, _bridgeAddress, _rollupManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonvalidiumstoragemigration{PolygonvalidiumstoragemigrationCaller: PolygonvalidiumstoragemigrationCaller{contract: contract}, PolygonvalidiumstoragemigrationTransactor: PolygonvalidiumstoragemigrationTransactor{contract: contract}, PolygonvalidiumstoragemigrationFilterer: PolygonvalidiumstoragemigrationFilterer{contract: contract}}, nil
}

// Polygonvalidiumstoragemigration is an auto generated Go binding around an Ethereum contract.
type Polygonvalidiumstoragemigration struct {
	PolygonvalidiumstoragemigrationCaller     // Read-only binding to the contract
	PolygonvalidiumstoragemigrationTransactor // Write-only binding to the contract
	PolygonvalidiumstoragemigrationFilterer   // Log filterer for contract events
}

// PolygonvalidiumstoragemigrationCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonvalidiumstoragemigrationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonvalidiumstoragemigrationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonvalidiumstoragemigrationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonvalidiumstoragemigrationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonvalidiumstoragemigrationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonvalidiumstoragemigrationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonvalidiumstoragemigrationSession struct {
	Contract     *Polygonvalidiumstoragemigration // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                    // Call options to use throughout this session
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// PolygonvalidiumstoragemigrationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonvalidiumstoragemigrationCallerSession struct {
	Contract *PolygonvalidiumstoragemigrationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                          // Call options to use throughout this session
}

// PolygonvalidiumstoragemigrationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonvalidiumstoragemigrationTransactorSession struct {
	Contract     *PolygonvalidiumstoragemigrationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                          // Transaction auth options to use throughout this session
}

// PolygonvalidiumstoragemigrationRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonvalidiumstoragemigrationRaw struct {
	Contract *Polygonvalidiumstoragemigration // Generic contract binding to access the raw methods on
}

// PolygonvalidiumstoragemigrationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonvalidiumstoragemigrationCallerRaw struct {
	Contract *PolygonvalidiumstoragemigrationCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonvalidiumstoragemigrationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonvalidiumstoragemigrationTransactorRaw struct {
	Contract *PolygonvalidiumstoragemigrationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonvalidiumstoragemigration creates a new instance of Polygonvalidiumstoragemigration, bound to a specific deployed contract.
func NewPolygonvalidiumstoragemigration(address common.Address, backend bind.ContractBackend) (*Polygonvalidiumstoragemigration, error) {
	contract, err := bindPolygonvalidiumstoragemigration(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonvalidiumstoragemigration{PolygonvalidiumstoragemigrationCaller: PolygonvalidiumstoragemigrationCaller{contract: contract}, PolygonvalidiumstoragemigrationTransactor: PolygonvalidiumstoragemigrationTransactor{contract: contract}, PolygonvalidiumstoragemigrationFilterer: PolygonvalidiumstoragemigrationFilterer{contract: contract}}, nil
}

// NewPolygonvalidiumstoragemigrationCaller creates a new read-only instance of Polygonvalidiumstoragemigration, bound to a specific deployed contract.
func NewPolygonvalidiumstoragemigrationCaller(address common.Address, caller bind.ContractCaller) (*PolygonvalidiumstoragemigrationCaller, error) {
	contract, err := bindPolygonvalidiumstoragemigration(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumstoragemigrationCaller{contract: contract}, nil
}

// NewPolygonvalidiumstoragemigrationTransactor creates a new write-only instance of Polygonvalidiumstoragemigration, bound to a specific deployed contract.
func NewPolygonvalidiumstoragemigrationTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonvalidiumstoragemigrationTransactor, error) {
	contract, err := bindPolygonvalidiumstoragemigration(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumstoragemigrationTransactor{contract: contract}, nil
}

// NewPolygonvalidiumstoragemigrationFilterer creates a new log filterer instance of Polygonvalidiumstoragemigration, bound to a specific deployed contract.
func NewPolygonvalidiumstoragemigrationFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonvalidiumstoragemigrationFilterer, error) {
	contract, err := bindPolygonvalidiumstoragemigration(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumstoragemigrationFilterer{contract: contract}, nil
}

// bindPolygonvalidiumstoragemigration binds a generic wrapper to an already deployed contract.
func bindPolygonvalidiumstoragemigration(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonvalidiumstoragemigrationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonvalidiumstoragemigration.Contract.PolygonvalidiumstoragemigrationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.PolygonvalidiumstoragemigrationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.PolygonvalidiumstoragemigrationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonvalidiumstoragemigration.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.contract.Transact(opts, method, params...)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) GLOBALEXITROOTMANAGERL2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "GLOBAL_EXIT_ROOT_MANAGER_L2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _Polygonvalidiumstoragemigration.Contract.GLOBALEXITROOTMANAGERL2(&_Polygonvalidiumstoragemigration.CallOpts)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _Polygonvalidiumstoragemigration.Contract.GLOBALEXITROOTMANAGERL2(&_Polygonvalidiumstoragemigration.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) INITIALIZETXBRIDGELISTLENLEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_LIST_LEN_LEN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _Polygonvalidiumstoragemigration.Contract.INITIALIZETXBRIDGELISTLENLEN(&_Polygonvalidiumstoragemigration.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _Polygonvalidiumstoragemigration.Contract.INITIALIZETXBRIDGELISTLENLEN(&_Polygonvalidiumstoragemigration.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) INITIALIZETXBRIDGEPARAMS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _Polygonvalidiumstoragemigration.Contract.INITIALIZETXBRIDGEPARAMS(&_Polygonvalidiumstoragemigration.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _Polygonvalidiumstoragemigration.Contract.INITIALIZETXBRIDGEPARAMS(&_Polygonvalidiumstoragemigration.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _Polygonvalidiumstoragemigration.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_Polygonvalidiumstoragemigration.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _Polygonvalidiumstoragemigration.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_Polygonvalidiumstoragemigration.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _Polygonvalidiumstoragemigration.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_Polygonvalidiumstoragemigration.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _Polygonvalidiumstoragemigration.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_Polygonvalidiumstoragemigration.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) INITIALIZETXCONSTANTBYTES(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _Polygonvalidiumstoragemigration.Contract.INITIALIZETXCONSTANTBYTES(&_Polygonvalidiumstoragemigration.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _Polygonvalidiumstoragemigration.Contract.INITIALIZETXCONSTANTBYTES(&_Polygonvalidiumstoragemigration.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) INITIALIZETXCONSTANTBYTESEMPTYMETADATA(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _Polygonvalidiumstoragemigration.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_Polygonvalidiumstoragemigration.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _Polygonvalidiumstoragemigration.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_Polygonvalidiumstoragemigration.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) INITIALIZETXDATALENEMPTYMETADATA(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "INITIALIZE_TX_DATA_LEN_EMPTY_METADATA")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _Polygonvalidiumstoragemigration.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_Polygonvalidiumstoragemigration.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _Polygonvalidiumstoragemigration.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_Polygonvalidiumstoragemigration.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) INITIALIZETXEFFECTIVEPERCENTAGE(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "INITIALIZE_TX_EFFECTIVE_PERCENTAGE")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _Polygonvalidiumstoragemigration.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_Polygonvalidiumstoragemigration.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _Polygonvalidiumstoragemigration.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_Polygonvalidiumstoragemigration.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) SIGNATUREINITIALIZETXR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_R")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _Polygonvalidiumstoragemigration.Contract.SIGNATUREINITIALIZETXR(&_Polygonvalidiumstoragemigration.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _Polygonvalidiumstoragemigration.Contract.SIGNATUREINITIALIZETXR(&_Polygonvalidiumstoragemigration.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) SIGNATUREINITIALIZETXS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_S")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _Polygonvalidiumstoragemigration.Contract.SIGNATUREINITIALIZETXS(&_Polygonvalidiumstoragemigration.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _Polygonvalidiumstoragemigration.Contract.SIGNATUREINITIALIZETXS(&_Polygonvalidiumstoragemigration.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) SIGNATUREINITIALIZETXV(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_V")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _Polygonvalidiumstoragemigration.Contract.SIGNATUREINITIALIZETXV(&_Polygonvalidiumstoragemigration.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _Polygonvalidiumstoragemigration.Contract.SIGNATUREINITIALIZETXV(&_Polygonvalidiumstoragemigration.CallOpts)
}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) TIMESTAMPRANGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "TIMESTAMP_RANGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) TIMESTAMPRANGE() (*big.Int, error) {
	return _Polygonvalidiumstoragemigration.Contract.TIMESTAMPRANGE(&_Polygonvalidiumstoragemigration.CallOpts)
}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) TIMESTAMPRANGE() (*big.Int, error) {
	return _Polygonvalidiumstoragemigration.Contract.TIMESTAMPRANGE(&_Polygonvalidiumstoragemigration.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) Admin() (common.Address, error) {
	return _Polygonvalidiumstoragemigration.Contract.Admin(&_Polygonvalidiumstoragemigration.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) Admin() (common.Address, error) {
	return _Polygonvalidiumstoragemigration.Contract.Admin(&_Polygonvalidiumstoragemigration.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) BridgeAddress() (common.Address, error) {
	return _Polygonvalidiumstoragemigration.Contract.BridgeAddress(&_Polygonvalidiumstoragemigration.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) BridgeAddress() (common.Address, error) {
	return _Polygonvalidiumstoragemigration.Contract.BridgeAddress(&_Polygonvalidiumstoragemigration.CallOpts)
}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) CalculatePolPerForceBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "calculatePolPerForceBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) CalculatePolPerForceBatch() (*big.Int, error) {
	return _Polygonvalidiumstoragemigration.Contract.CalculatePolPerForceBatch(&_Polygonvalidiumstoragemigration.CallOpts)
}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) CalculatePolPerForceBatch() (*big.Int, error) {
	return _Polygonvalidiumstoragemigration.Contract.CalculatePolPerForceBatch(&_Polygonvalidiumstoragemigration.CallOpts)
}

// DataAvailabilityProtocol is a free data retrieval call binding the contract method 0xe57a0b4c.
//
// Solidity: function dataAvailabilityProtocol() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) DataAvailabilityProtocol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "dataAvailabilityProtocol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataAvailabilityProtocol is a free data retrieval call binding the contract method 0xe57a0b4c.
//
// Solidity: function dataAvailabilityProtocol() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) DataAvailabilityProtocol() (common.Address, error) {
	return _Polygonvalidiumstoragemigration.Contract.DataAvailabilityProtocol(&_Polygonvalidiumstoragemigration.CallOpts)
}

// DataAvailabilityProtocol is a free data retrieval call binding the contract method 0xe57a0b4c.
//
// Solidity: function dataAvailabilityProtocol() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) DataAvailabilityProtocol() (common.Address, error) {
	return _Polygonvalidiumstoragemigration.Contract.DataAvailabilityProtocol(&_Polygonvalidiumstoragemigration.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) ForceBatchAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "forceBatchAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) ForceBatchAddress() (common.Address, error) {
	return _Polygonvalidiumstoragemigration.Contract.ForceBatchAddress(&_Polygonvalidiumstoragemigration.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) ForceBatchAddress() (common.Address, error) {
	return _Polygonvalidiumstoragemigration.Contract.ForceBatchAddress(&_Polygonvalidiumstoragemigration.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonvalidiumstoragemigration.Contract.ForceBatchTimeout(&_Polygonvalidiumstoragemigration.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonvalidiumstoragemigration.Contract.ForceBatchTimeout(&_Polygonvalidiumstoragemigration.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonvalidiumstoragemigration.Contract.ForcedBatches(&_Polygonvalidiumstoragemigration.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonvalidiumstoragemigration.Contract.ForcedBatches(&_Polygonvalidiumstoragemigration.CallOpts, arg0)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) GasTokenAddress() (common.Address, error) {
	return _Polygonvalidiumstoragemigration.Contract.GasTokenAddress(&_Polygonvalidiumstoragemigration.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) GasTokenAddress() (common.Address, error) {
	return _Polygonvalidiumstoragemigration.Contract.GasTokenAddress(&_Polygonvalidiumstoragemigration.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) GasTokenNetwork() (uint32, error) {
	return _Polygonvalidiumstoragemigration.Contract.GasTokenNetwork(&_Polygonvalidiumstoragemigration.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) GasTokenNetwork() (uint32, error) {
	return _Polygonvalidiumstoragemigration.Contract.GasTokenNetwork(&_Polygonvalidiumstoragemigration.CallOpts)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) GenerateInitializeTransaction(opts *bind.CallOpts, networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "generateInitializeTransaction", networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) GenerateInitializeTransaction(networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _Polygonvalidiumstoragemigration.Contract.GenerateInitializeTransaction(&_Polygonvalidiumstoragemigration.CallOpts, networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) GenerateInitializeTransaction(networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _Polygonvalidiumstoragemigration.Contract.GenerateInitializeTransaction(&_Polygonvalidiumstoragemigration.CallOpts, networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonvalidiumstoragemigration.Contract.GlobalExitRootManager(&_Polygonvalidiumstoragemigration.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonvalidiumstoragemigration.Contract.GlobalExitRootManager(&_Polygonvalidiumstoragemigration.CallOpts)
}

// IsSequenceWithDataAvailabilityAllowed is a free data retrieval call binding the contract method 0x4c21fef3.
//
// Solidity: function isSequenceWithDataAvailabilityAllowed() view returns(bool)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) IsSequenceWithDataAvailabilityAllowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "isSequenceWithDataAvailabilityAllowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSequenceWithDataAvailabilityAllowed is a free data retrieval call binding the contract method 0x4c21fef3.
//
// Solidity: function isSequenceWithDataAvailabilityAllowed() view returns(bool)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) IsSequenceWithDataAvailabilityAllowed() (bool, error) {
	return _Polygonvalidiumstoragemigration.Contract.IsSequenceWithDataAvailabilityAllowed(&_Polygonvalidiumstoragemigration.CallOpts)
}

// IsSequenceWithDataAvailabilityAllowed is a free data retrieval call binding the contract method 0x4c21fef3.
//
// Solidity: function isSequenceWithDataAvailabilityAllowed() view returns(bool)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) IsSequenceWithDataAvailabilityAllowed() (bool, error) {
	return _Polygonvalidiumstoragemigration.Contract.IsSequenceWithDataAvailabilityAllowed(&_Polygonvalidiumstoragemigration.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) LastAccInputHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "lastAccInputHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) LastAccInputHash() ([32]byte, error) {
	return _Polygonvalidiumstoragemigration.Contract.LastAccInputHash(&_Polygonvalidiumstoragemigration.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) LastAccInputHash() ([32]byte, error) {
	return _Polygonvalidiumstoragemigration.Contract.LastAccInputHash(&_Polygonvalidiumstoragemigration.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) LastForceBatch() (uint64, error) {
	return _Polygonvalidiumstoragemigration.Contract.LastForceBatch(&_Polygonvalidiumstoragemigration.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) LastForceBatch() (uint64, error) {
	return _Polygonvalidiumstoragemigration.Contract.LastForceBatch(&_Polygonvalidiumstoragemigration.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonvalidiumstoragemigration.Contract.LastForceBatchSequenced(&_Polygonvalidiumstoragemigration.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonvalidiumstoragemigration.Contract.LastForceBatchSequenced(&_Polygonvalidiumstoragemigration.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) NetworkName() (string, error) {
	return _Polygonvalidiumstoragemigration.Contract.NetworkName(&_Polygonvalidiumstoragemigration.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) NetworkName() (string, error) {
	return _Polygonvalidiumstoragemigration.Contract.NetworkName(&_Polygonvalidiumstoragemigration.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) PendingAdmin() (common.Address, error) {
	return _Polygonvalidiumstoragemigration.Contract.PendingAdmin(&_Polygonvalidiumstoragemigration.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) PendingAdmin() (common.Address, error) {
	return _Polygonvalidiumstoragemigration.Contract.PendingAdmin(&_Polygonvalidiumstoragemigration.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) Pol() (common.Address, error) {
	return _Polygonvalidiumstoragemigration.Contract.Pol(&_Polygonvalidiumstoragemigration.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) Pol() (common.Address, error) {
	return _Polygonvalidiumstoragemigration.Contract.Pol(&_Polygonvalidiumstoragemigration.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) RollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "rollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) RollupManager() (common.Address, error) {
	return _Polygonvalidiumstoragemigration.Contract.RollupManager(&_Polygonvalidiumstoragemigration.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) RollupManager() (common.Address, error) {
	return _Polygonvalidiumstoragemigration.Contract.RollupManager(&_Polygonvalidiumstoragemigration.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) TrustedSequencer() (common.Address, error) {
	return _Polygonvalidiumstoragemigration.Contract.TrustedSequencer(&_Polygonvalidiumstoragemigration.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) TrustedSequencer() (common.Address, error) {
	return _Polygonvalidiumstoragemigration.Contract.TrustedSequencer(&_Polygonvalidiumstoragemigration.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCaller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonvalidiumstoragemigration.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) TrustedSequencerURL() (string, error) {
	return _Polygonvalidiumstoragemigration.Contract.TrustedSequencerURL(&_Polygonvalidiumstoragemigration.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationCallerSession) TrustedSequencerURL() (string, error) {
	return _Polygonvalidiumstoragemigration.Contract.TrustedSequencerURL(&_Polygonvalidiumstoragemigration.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.AcceptAdminRole(&_Polygonvalidiumstoragemigration.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.AcceptAdminRole(&_Polygonvalidiumstoragemigration.TransactOpts)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactor) ForceBatch(opts *bind.TransactOpts, transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.contract.Transact(opts, "forceBatch", transactions, polAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) ForceBatch(transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.ForceBatch(&_Polygonvalidiumstoragemigration.TransactOpts, transactions, polAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactorSession) ForceBatch(transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.ForceBatch(&_Polygonvalidiumstoragemigration.TransactOpts, transactions, polAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.contract.Transact(opts, "initialize", _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.Initialize(&_Polygonvalidiumstoragemigration.TransactOpts, _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactorSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.Initialize(&_Polygonvalidiumstoragemigration.TransactOpts, _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// InitializeMigration is a paid mutator transaction binding the contract method 0x1c8b9370.
//
// Solidity: function initializeMigration() returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactor) InitializeMigration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.contract.Transact(opts, "initializeMigration")
}

// InitializeMigration is a paid mutator transaction binding the contract method 0x1c8b9370.
//
// Solidity: function initializeMigration() returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) InitializeMigration() (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.InitializeMigration(&_Polygonvalidiumstoragemigration.TransactOpts)
}

// InitializeMigration is a paid mutator transaction binding the contract method 0x1c8b9370.
//
// Solidity: function initializeMigration() returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactorSession) InitializeMigration() (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.InitializeMigration(&_Polygonvalidiumstoragemigration.TransactOpts)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactor) OnVerifyBatches(opts *bind.TransactOpts, lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.contract.Transact(opts, "onVerifyBatches", lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.OnVerifyBatches(&_Polygonvalidiumstoragemigration.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactorSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.OnVerifyBatches(&_Polygonvalidiumstoragemigration.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xdef57e54.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactor) SequenceBatches(opts *bind.TransactOpts, batches []PolygonRollupBaseEtrogNoGapBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.contract.Transact(opts, "sequenceBatches", batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xdef57e54.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) SequenceBatches(batches []PolygonRollupBaseEtrogNoGapBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.SequenceBatches(&_Polygonvalidiumstoragemigration.TransactOpts, batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xdef57e54.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactorSession) SequenceBatches(batches []PolygonRollupBaseEtrogNoGapBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.SequenceBatches(&_Polygonvalidiumstoragemigration.TransactOpts, batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase)
}

// SequenceBatchesValidium is a paid mutator transaction binding the contract method 0xdb5b0ed7.
//
// Solidity: function sequenceBatchesValidium((bytes32,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactor) SequenceBatchesValidium(opts *bind.TransactOpts, batches []PolygonValidiumStorageMigrationValidiumBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.contract.Transact(opts, "sequenceBatchesValidium", batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase, dataAvailabilityMessage)
}

// SequenceBatchesValidium is a paid mutator transaction binding the contract method 0xdb5b0ed7.
//
// Solidity: function sequenceBatchesValidium((bytes32,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) SequenceBatchesValidium(batches []PolygonValidiumStorageMigrationValidiumBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.SequenceBatchesValidium(&_Polygonvalidiumstoragemigration.TransactOpts, batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase, dataAvailabilityMessage)
}

// SequenceBatchesValidium is a paid mutator transaction binding the contract method 0xdb5b0ed7.
//
// Solidity: function sequenceBatchesValidium((bytes32,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactorSession) SequenceBatchesValidium(batches []PolygonValidiumStorageMigrationValidiumBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.SequenceBatchesValidium(&_Polygonvalidiumstoragemigration.TransactOpts, batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase, dataAvailabilityMessage)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactor) SequenceForceBatches(opts *bind.TransactOpts, batches []PolygonRollupBaseEtrogNoGapBatchData) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.contract.Transact(opts, "sequenceForceBatches", batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) SequenceForceBatches(batches []PolygonRollupBaseEtrogNoGapBatchData) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.SequenceForceBatches(&_Polygonvalidiumstoragemigration.TransactOpts, batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactorSession) SequenceForceBatches(batches []PolygonRollupBaseEtrogNoGapBatchData) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.SequenceForceBatches(&_Polygonvalidiumstoragemigration.TransactOpts, batches)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactor) SetDataAvailabilityProtocol(opts *bind.TransactOpts, newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.contract.Transact(opts, "setDataAvailabilityProtocol", newDataAvailabilityProtocol)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) SetDataAvailabilityProtocol(newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.SetDataAvailabilityProtocol(&_Polygonvalidiumstoragemigration.TransactOpts, newDataAvailabilityProtocol)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactorSession) SetDataAvailabilityProtocol(newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.SetDataAvailabilityProtocol(&_Polygonvalidiumstoragemigration.TransactOpts, newDataAvailabilityProtocol)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactor) SetForceBatchAddress(opts *bind.TransactOpts, newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.contract.Transact(opts, "setForceBatchAddress", newForceBatchAddress)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) SetForceBatchAddress(newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.SetForceBatchAddress(&_Polygonvalidiumstoragemigration.TransactOpts, newForceBatchAddress)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactorSession) SetForceBatchAddress(newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.SetForceBatchAddress(&_Polygonvalidiumstoragemigration.TransactOpts, newForceBatchAddress)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactor) SetForceBatchTimeout(opts *bind.TransactOpts, newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.contract.Transact(opts, "setForceBatchTimeout", newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.SetForceBatchTimeout(&_Polygonvalidiumstoragemigration.TransactOpts, newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactorSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.SetForceBatchTimeout(&_Polygonvalidiumstoragemigration.TransactOpts, newforceBatchTimeout)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.SetTrustedSequencer(&_Polygonvalidiumstoragemigration.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.SetTrustedSequencer(&_Polygonvalidiumstoragemigration.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.SetTrustedSequencerURL(&_Polygonvalidiumstoragemigration.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.SetTrustedSequencerURL(&_Polygonvalidiumstoragemigration.TransactOpts, newTrustedSequencerURL)
}

// SwitchSequenceWithDataAvailability is a paid mutator transaction binding the contract method 0x2acdc2b6.
//
// Solidity: function switchSequenceWithDataAvailability(bool newIsSequenceWithDataAvailabilityAllowed) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactor) SwitchSequenceWithDataAvailability(opts *bind.TransactOpts, newIsSequenceWithDataAvailabilityAllowed bool) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.contract.Transact(opts, "switchSequenceWithDataAvailability", newIsSequenceWithDataAvailabilityAllowed)
}

// SwitchSequenceWithDataAvailability is a paid mutator transaction binding the contract method 0x2acdc2b6.
//
// Solidity: function switchSequenceWithDataAvailability(bool newIsSequenceWithDataAvailabilityAllowed) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) SwitchSequenceWithDataAvailability(newIsSequenceWithDataAvailabilityAllowed bool) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.SwitchSequenceWithDataAvailability(&_Polygonvalidiumstoragemigration.TransactOpts, newIsSequenceWithDataAvailabilityAllowed)
}

// SwitchSequenceWithDataAvailability is a paid mutator transaction binding the contract method 0x2acdc2b6.
//
// Solidity: function switchSequenceWithDataAvailability(bool newIsSequenceWithDataAvailabilityAllowed) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactorSession) SwitchSequenceWithDataAvailability(newIsSequenceWithDataAvailabilityAllowed bool) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.SwitchSequenceWithDataAvailability(&_Polygonvalidiumstoragemigration.TransactOpts, newIsSequenceWithDataAvailabilityAllowed)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.TransferAdminRole(&_Polygonvalidiumstoragemigration.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumstoragemigration.Contract.TransferAdminRole(&_Polygonvalidiumstoragemigration.TransactOpts, newPendingAdmin)
}

// PolygonvalidiumstoragemigrationAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationAcceptAdminRoleIterator struct {
	Event *PolygonvalidiumstoragemigrationAcceptAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumstoragemigrationAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumstoragemigrationAcceptAdminRole)
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
		it.Event = new(PolygonvalidiumstoragemigrationAcceptAdminRole)
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
func (it *PolygonvalidiumstoragemigrationAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumstoragemigrationAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumstoragemigrationAcceptAdminRole represents a AcceptAdminRole event raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*PolygonvalidiumstoragemigrationAcceptAdminRoleIterator, error) {

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumstoragemigrationAcceptAdminRoleIterator{contract: _Polygonvalidiumstoragemigration.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumstoragemigrationAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumstoragemigrationAcceptAdminRole)
				if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
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
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) ParseAcceptAdminRole(log types.Log) (*PolygonvalidiumstoragemigrationAcceptAdminRole, error) {
	event := new(PolygonvalidiumstoragemigrationAcceptAdminRole)
	if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumstoragemigrationForceBatchIterator is returned from FilterForceBatch and is used to iterate over the raw logs and unpacked data for ForceBatch events raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationForceBatchIterator struct {
	Event *PolygonvalidiumstoragemigrationForceBatch // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumstoragemigrationForceBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumstoragemigrationForceBatch)
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
		it.Event = new(PolygonvalidiumstoragemigrationForceBatch)
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
func (it *PolygonvalidiumstoragemigrationForceBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumstoragemigrationForceBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumstoragemigrationForceBatch represents a ForceBatch event raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationForceBatch struct {
	ForceBatchNum      uint64
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Transactions       []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterForceBatch is a free log retrieval operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) FilterForceBatch(opts *bind.FilterOpts, forceBatchNum []uint64) (*PolygonvalidiumstoragemigrationForceBatchIterator, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.FilterLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumstoragemigrationForceBatchIterator{contract: _Polygonvalidiumstoragemigration.contract, event: "ForceBatch", logs: logs, sub: sub}, nil
}

// WatchForceBatch is a free log subscription operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) WatchForceBatch(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumstoragemigrationForceBatch, forceBatchNum []uint64) (event.Subscription, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.WatchLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumstoragemigrationForceBatch)
				if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "ForceBatch", log); err != nil {
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
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) ParseForceBatch(log types.Log) (*PolygonvalidiumstoragemigrationForceBatch, error) {
	event := new(PolygonvalidiumstoragemigrationForceBatch)
	if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "ForceBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumstoragemigrationInitialSequenceBatchesIterator is returned from FilterInitialSequenceBatches and is used to iterate over the raw logs and unpacked data for InitialSequenceBatches events raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationInitialSequenceBatchesIterator struct {
	Event *PolygonvalidiumstoragemigrationInitialSequenceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumstoragemigrationInitialSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumstoragemigrationInitialSequenceBatches)
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
		it.Event = new(PolygonvalidiumstoragemigrationInitialSequenceBatches)
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
func (it *PolygonvalidiumstoragemigrationInitialSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumstoragemigrationInitialSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumstoragemigrationInitialSequenceBatches represents a InitialSequenceBatches event raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationInitialSequenceBatches struct {
	Transactions       []byte
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterInitialSequenceBatches is a free log retrieval operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) FilterInitialSequenceBatches(opts *bind.FilterOpts) (*PolygonvalidiumstoragemigrationInitialSequenceBatchesIterator, error) {

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.FilterLogs(opts, "InitialSequenceBatches")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumstoragemigrationInitialSequenceBatchesIterator{contract: _Polygonvalidiumstoragemigration.contract, event: "InitialSequenceBatches", logs: logs, sub: sub}, nil
}

// WatchInitialSequenceBatches is a free log subscription operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) WatchInitialSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumstoragemigrationInitialSequenceBatches) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.WatchLogs(opts, "InitialSequenceBatches")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumstoragemigrationInitialSequenceBatches)
				if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "InitialSequenceBatches", log); err != nil {
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
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) ParseInitialSequenceBatches(log types.Log) (*PolygonvalidiumstoragemigrationInitialSequenceBatches, error) {
	event := new(PolygonvalidiumstoragemigrationInitialSequenceBatches)
	if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "InitialSequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumstoragemigrationInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationInitializedIterator struct {
	Event *PolygonvalidiumstoragemigrationInitialized // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumstoragemigrationInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumstoragemigrationInitialized)
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
		it.Event = new(PolygonvalidiumstoragemigrationInitialized)
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
func (it *PolygonvalidiumstoragemigrationInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumstoragemigrationInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumstoragemigrationInitialized represents a Initialized event raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) FilterInitialized(opts *bind.FilterOpts) (*PolygonvalidiumstoragemigrationInitializedIterator, error) {

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumstoragemigrationInitializedIterator{contract: _Polygonvalidiumstoragemigration.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumstoragemigrationInitialized) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumstoragemigrationInitialized)
				if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) ParseInitialized(log types.Log) (*PolygonvalidiumstoragemigrationInitialized, error) {
	event := new(PolygonvalidiumstoragemigrationInitialized)
	if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumstoragemigrationSequenceBatchesIterator is returned from FilterSequenceBatches and is used to iterate over the raw logs and unpacked data for SequenceBatches events raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationSequenceBatchesIterator struct {
	Event *PolygonvalidiumstoragemigrationSequenceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumstoragemigrationSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumstoragemigrationSequenceBatches)
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
		it.Event = new(PolygonvalidiumstoragemigrationSequenceBatches)
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
func (it *PolygonvalidiumstoragemigrationSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumstoragemigrationSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumstoragemigrationSequenceBatches represents a SequenceBatches event raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationSequenceBatches struct {
	NumBatch   uint64
	L1InfoRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSequenceBatches is a free log retrieval operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) FilterSequenceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonvalidiumstoragemigrationSequenceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.FilterLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumstoragemigrationSequenceBatchesIterator{contract: _Polygonvalidiumstoragemigration.contract, event: "SequenceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceBatches is a free log subscription operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) WatchSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumstoragemigrationSequenceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.WatchLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumstoragemigrationSequenceBatches)
				if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
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
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) ParseSequenceBatches(log types.Log) (*PolygonvalidiumstoragemigrationSequenceBatches, error) {
	event := new(PolygonvalidiumstoragemigrationSequenceBatches)
	if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumstoragemigrationSequenceForceBatchesIterator is returned from FilterSequenceForceBatches and is used to iterate over the raw logs and unpacked data for SequenceForceBatches events raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationSequenceForceBatchesIterator struct {
	Event *PolygonvalidiumstoragemigrationSequenceForceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumstoragemigrationSequenceForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumstoragemigrationSequenceForceBatches)
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
		it.Event = new(PolygonvalidiumstoragemigrationSequenceForceBatches)
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
func (it *PolygonvalidiumstoragemigrationSequenceForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumstoragemigrationSequenceForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumstoragemigrationSequenceForceBatches represents a SequenceForceBatches event raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationSequenceForceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceForceBatches is a free log retrieval operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) FilterSequenceForceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonvalidiumstoragemigrationSequenceForceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.FilterLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumstoragemigrationSequenceForceBatchesIterator{contract: _Polygonvalidiumstoragemigration.contract, event: "SequenceForceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceForceBatches is a free log subscription operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) WatchSequenceForceBatches(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumstoragemigrationSequenceForceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.WatchLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumstoragemigrationSequenceForceBatches)
				if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
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
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) ParseSequenceForceBatches(log types.Log) (*PolygonvalidiumstoragemigrationSequenceForceBatches, error) {
	event := new(PolygonvalidiumstoragemigrationSequenceForceBatches)
	if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumstoragemigrationSetDataAvailabilityProtocolIterator is returned from FilterSetDataAvailabilityProtocol and is used to iterate over the raw logs and unpacked data for SetDataAvailabilityProtocol events raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationSetDataAvailabilityProtocolIterator struct {
	Event *PolygonvalidiumstoragemigrationSetDataAvailabilityProtocol // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumstoragemigrationSetDataAvailabilityProtocolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumstoragemigrationSetDataAvailabilityProtocol)
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
		it.Event = new(PolygonvalidiumstoragemigrationSetDataAvailabilityProtocol)
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
func (it *PolygonvalidiumstoragemigrationSetDataAvailabilityProtocolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumstoragemigrationSetDataAvailabilityProtocolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumstoragemigrationSetDataAvailabilityProtocol represents a SetDataAvailabilityProtocol event raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationSetDataAvailabilityProtocol struct {
	NewDataAvailabilityProtocol common.Address
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterSetDataAvailabilityProtocol is a free log retrieval operation binding the contract event 0xd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a.
//
// Solidity: event SetDataAvailabilityProtocol(address newDataAvailabilityProtocol)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) FilterSetDataAvailabilityProtocol(opts *bind.FilterOpts) (*PolygonvalidiumstoragemigrationSetDataAvailabilityProtocolIterator, error) {

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.FilterLogs(opts, "SetDataAvailabilityProtocol")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumstoragemigrationSetDataAvailabilityProtocolIterator{contract: _Polygonvalidiumstoragemigration.contract, event: "SetDataAvailabilityProtocol", logs: logs, sub: sub}, nil
}

// WatchSetDataAvailabilityProtocol is a free log subscription operation binding the contract event 0xd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a.
//
// Solidity: event SetDataAvailabilityProtocol(address newDataAvailabilityProtocol)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) WatchSetDataAvailabilityProtocol(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumstoragemigrationSetDataAvailabilityProtocol) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.WatchLogs(opts, "SetDataAvailabilityProtocol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumstoragemigrationSetDataAvailabilityProtocol)
				if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "SetDataAvailabilityProtocol", log); err != nil {
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
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) ParseSetDataAvailabilityProtocol(log types.Log) (*PolygonvalidiumstoragemigrationSetDataAvailabilityProtocol, error) {
	event := new(PolygonvalidiumstoragemigrationSetDataAvailabilityProtocol)
	if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "SetDataAvailabilityProtocol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumstoragemigrationSetForceBatchAddressIterator is returned from FilterSetForceBatchAddress and is used to iterate over the raw logs and unpacked data for SetForceBatchAddress events raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationSetForceBatchAddressIterator struct {
	Event *PolygonvalidiumstoragemigrationSetForceBatchAddress // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumstoragemigrationSetForceBatchAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumstoragemigrationSetForceBatchAddress)
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
		it.Event = new(PolygonvalidiumstoragemigrationSetForceBatchAddress)
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
func (it *PolygonvalidiumstoragemigrationSetForceBatchAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumstoragemigrationSetForceBatchAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumstoragemigrationSetForceBatchAddress represents a SetForceBatchAddress event raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationSetForceBatchAddress struct {
	NewForceBatchAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchAddress is a free log retrieval operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) FilterSetForceBatchAddress(opts *bind.FilterOpts) (*PolygonvalidiumstoragemigrationSetForceBatchAddressIterator, error) {

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.FilterLogs(opts, "SetForceBatchAddress")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumstoragemigrationSetForceBatchAddressIterator{contract: _Polygonvalidiumstoragemigration.contract, event: "SetForceBatchAddress", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchAddress is a free log subscription operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) WatchSetForceBatchAddress(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumstoragemigrationSetForceBatchAddress) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.WatchLogs(opts, "SetForceBatchAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumstoragemigrationSetForceBatchAddress)
				if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "SetForceBatchAddress", log); err != nil {
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
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) ParseSetForceBatchAddress(log types.Log) (*PolygonvalidiumstoragemigrationSetForceBatchAddress, error) {
	event := new(PolygonvalidiumstoragemigrationSetForceBatchAddress)
	if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "SetForceBatchAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumstoragemigrationSetForceBatchTimeoutIterator is returned from FilterSetForceBatchTimeout and is used to iterate over the raw logs and unpacked data for SetForceBatchTimeout events raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationSetForceBatchTimeoutIterator struct {
	Event *PolygonvalidiumstoragemigrationSetForceBatchTimeout // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumstoragemigrationSetForceBatchTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumstoragemigrationSetForceBatchTimeout)
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
		it.Event = new(PolygonvalidiumstoragemigrationSetForceBatchTimeout)
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
func (it *PolygonvalidiumstoragemigrationSetForceBatchTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumstoragemigrationSetForceBatchTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumstoragemigrationSetForceBatchTimeout represents a SetForceBatchTimeout event raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationSetForceBatchTimeout struct {
	NewforceBatchTimeout uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchTimeout is a free log retrieval operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) FilterSetForceBatchTimeout(opts *bind.FilterOpts) (*PolygonvalidiumstoragemigrationSetForceBatchTimeoutIterator, error) {

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.FilterLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumstoragemigrationSetForceBatchTimeoutIterator{contract: _Polygonvalidiumstoragemigration.contract, event: "SetForceBatchTimeout", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchTimeout is a free log subscription operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) WatchSetForceBatchTimeout(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumstoragemigrationSetForceBatchTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.WatchLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumstoragemigrationSetForceBatchTimeout)
				if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
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
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) ParseSetForceBatchTimeout(log types.Log) (*PolygonvalidiumstoragemigrationSetForceBatchTimeout, error) {
	event := new(PolygonvalidiumstoragemigrationSetForceBatchTimeout)
	if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumstoragemigrationSetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationSetTrustedSequencerIterator struct {
	Event *PolygonvalidiumstoragemigrationSetTrustedSequencer // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumstoragemigrationSetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumstoragemigrationSetTrustedSequencer)
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
		it.Event = new(PolygonvalidiumstoragemigrationSetTrustedSequencer)
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
func (it *PolygonvalidiumstoragemigrationSetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumstoragemigrationSetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumstoragemigrationSetTrustedSequencer represents a SetTrustedSequencer event raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationSetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*PolygonvalidiumstoragemigrationSetTrustedSequencerIterator, error) {

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumstoragemigrationSetTrustedSequencerIterator{contract: _Polygonvalidiumstoragemigration.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumstoragemigrationSetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumstoragemigrationSetTrustedSequencer)
				if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
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
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) ParseSetTrustedSequencer(log types.Log) (*PolygonvalidiumstoragemigrationSetTrustedSequencer, error) {
	event := new(PolygonvalidiumstoragemigrationSetTrustedSequencer)
	if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumstoragemigrationSetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationSetTrustedSequencerURLIterator struct {
	Event *PolygonvalidiumstoragemigrationSetTrustedSequencerURL // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumstoragemigrationSetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumstoragemigrationSetTrustedSequencerURL)
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
		it.Event = new(PolygonvalidiumstoragemigrationSetTrustedSequencerURL)
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
func (it *PolygonvalidiumstoragemigrationSetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumstoragemigrationSetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumstoragemigrationSetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationSetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*PolygonvalidiumstoragemigrationSetTrustedSequencerURLIterator, error) {

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumstoragemigrationSetTrustedSequencerURLIterator{contract: _Polygonvalidiumstoragemigration.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumstoragemigrationSetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumstoragemigrationSetTrustedSequencerURL)
				if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
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
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) ParseSetTrustedSequencerURL(log types.Log) (*PolygonvalidiumstoragemigrationSetTrustedSequencerURL, error) {
	event := new(PolygonvalidiumstoragemigrationSetTrustedSequencerURL)
	if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumstoragemigrationSwitchSequenceWithDataAvailabilityIterator is returned from FilterSwitchSequenceWithDataAvailability and is used to iterate over the raw logs and unpacked data for SwitchSequenceWithDataAvailability events raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationSwitchSequenceWithDataAvailabilityIterator struct {
	Event *PolygonvalidiumstoragemigrationSwitchSequenceWithDataAvailability // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumstoragemigrationSwitchSequenceWithDataAvailabilityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumstoragemigrationSwitchSequenceWithDataAvailability)
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
		it.Event = new(PolygonvalidiumstoragemigrationSwitchSequenceWithDataAvailability)
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
func (it *PolygonvalidiumstoragemigrationSwitchSequenceWithDataAvailabilityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumstoragemigrationSwitchSequenceWithDataAvailabilityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumstoragemigrationSwitchSequenceWithDataAvailability represents a SwitchSequenceWithDataAvailability event raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationSwitchSequenceWithDataAvailability struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSwitchSequenceWithDataAvailability is a free log retrieval operation binding the contract event 0xf32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41.
//
// Solidity: event SwitchSequenceWithDataAvailability()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) FilterSwitchSequenceWithDataAvailability(opts *bind.FilterOpts) (*PolygonvalidiumstoragemigrationSwitchSequenceWithDataAvailabilityIterator, error) {

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.FilterLogs(opts, "SwitchSequenceWithDataAvailability")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumstoragemigrationSwitchSequenceWithDataAvailabilityIterator{contract: _Polygonvalidiumstoragemigration.contract, event: "SwitchSequenceWithDataAvailability", logs: logs, sub: sub}, nil
}

// WatchSwitchSequenceWithDataAvailability is a free log subscription operation binding the contract event 0xf32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41.
//
// Solidity: event SwitchSequenceWithDataAvailability()
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) WatchSwitchSequenceWithDataAvailability(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumstoragemigrationSwitchSequenceWithDataAvailability) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.WatchLogs(opts, "SwitchSequenceWithDataAvailability")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumstoragemigrationSwitchSequenceWithDataAvailability)
				if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "SwitchSequenceWithDataAvailability", log); err != nil {
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
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) ParseSwitchSequenceWithDataAvailability(log types.Log) (*PolygonvalidiumstoragemigrationSwitchSequenceWithDataAvailability, error) {
	event := new(PolygonvalidiumstoragemigrationSwitchSequenceWithDataAvailability)
	if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "SwitchSequenceWithDataAvailability", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumstoragemigrationTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationTransferAdminRoleIterator struct {
	Event *PolygonvalidiumstoragemigrationTransferAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumstoragemigrationTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumstoragemigrationTransferAdminRole)
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
		it.Event = new(PolygonvalidiumstoragemigrationTransferAdminRole)
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
func (it *PolygonvalidiumstoragemigrationTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumstoragemigrationTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumstoragemigrationTransferAdminRole represents a TransferAdminRole event raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*PolygonvalidiumstoragemigrationTransferAdminRoleIterator, error) {

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumstoragemigrationTransferAdminRoleIterator{contract: _Polygonvalidiumstoragemigration.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumstoragemigrationTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumstoragemigrationTransferAdminRole)
				if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
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
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) ParseTransferAdminRole(log types.Log) (*PolygonvalidiumstoragemigrationTransferAdminRole, error) {
	event := new(PolygonvalidiumstoragemigrationTransferAdminRole)
	if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumstoragemigrationVerifyBatchesIterator is returned from FilterVerifyBatches and is used to iterate over the raw logs and unpacked data for VerifyBatches events raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationVerifyBatchesIterator struct {
	Event *PolygonvalidiumstoragemigrationVerifyBatches // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumstoragemigrationVerifyBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumstoragemigrationVerifyBatches)
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
		it.Event = new(PolygonvalidiumstoragemigrationVerifyBatches)
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
func (it *PolygonvalidiumstoragemigrationVerifyBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumstoragemigrationVerifyBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumstoragemigrationVerifyBatches represents a VerifyBatches event raised by the Polygonvalidiumstoragemigration contract.
type PolygonvalidiumstoragemigrationVerifyBatches struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatches is a free log retrieval operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) FilterVerifyBatches(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*PolygonvalidiumstoragemigrationVerifyBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.FilterLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumstoragemigrationVerifyBatchesIterator{contract: _Polygonvalidiumstoragemigration.contract, event: "VerifyBatches", logs: logs, sub: sub}, nil
}

// WatchVerifyBatches is a free log subscription operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) WatchVerifyBatches(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumstoragemigrationVerifyBatches, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonvalidiumstoragemigration.contract.WatchLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumstoragemigrationVerifyBatches)
				if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
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
func (_Polygonvalidiumstoragemigration *PolygonvalidiumstoragemigrationFilterer) ParseVerifyBatches(log types.Log) (*PolygonvalidiumstoragemigrationVerifyBatches, error) {
	event := new(PolygonvalidiumstoragemigrationVerifyBatches)
	if err := _Polygonvalidiumstoragemigration.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

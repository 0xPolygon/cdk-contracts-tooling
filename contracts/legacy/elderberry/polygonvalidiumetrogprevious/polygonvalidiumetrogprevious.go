// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonvalidiumetrogprevious

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

// PolygonValidiumEtrogPreviousValidiumBatchData is an auto generated low-level Go binding around an user-defined struct.
type PolygonValidiumEtrogPreviousValidiumBatchData struct {
	TransactionsHash     [32]byte
	ForcedGlobalExitRoot [32]byte
	ForcedTimestamp      uint64
	ForcedBlockHashL1    [32]byte
}

// PolygonvalidiumetrogpreviousMetaData contains all meta data concerning the Polygonvalidiumetrogprevious contract.
var PolygonvalidiumetrogpreviousMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"_rollupManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceWithDataAvailabilityNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwitchToSameValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"forceBatchNum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"}],\"name\":\"ForceBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"}],\"name\":\"InitialSequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"l1InfoRoot\",\"type\":\"bytes32\"}],\"name\":\"SequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newDataAvailabilityProtocol\",\"type\":\"address\"}],\"name\":\"SetDataAvailabilityProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newForceBatchAddress\",\"type\":\"address\"}],\"name\":\"SetForceBatchAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"SetForceBatchTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SwitchSequenceWithDataAvailability\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatches\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GLOBAL_EXIT_ROOT_MANAGER_L2\",\"outputs\":[{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_LIST_LEN_LEN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_DATA_LEN_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_EFFECTIVE_PERCENTAGE\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_R\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_S\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_V\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculatePolPerForceBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataAvailabilityProtocol\",\"outputs\":[{\"internalType\":\"contractIDataAvailabilityProtocol\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"polAmount\",\"type\":\"uint256\"}],\"name\":\"forceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"}],\"name\":\"generateInitializeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSequenceWithDataAvailabilityAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"onVerifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrogPrevious.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"}],\"name\":\"sequenceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonValidiumEtrogPrevious.ValidiumBatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataAvailabilityMessage\",\"type\":\"bytes\"}],\"name\":\"sequenceBatchesValidium\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrogPrevious.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"}],\"name\":\"sequenceForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIDataAvailabilityProtocol\",\"name\":\"newDataAvailabilityProtocol\",\"type\":\"address\"}],\"name\":\"setDataAvailabilityProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newForceBatchAddress\",\"type\":\"address\"}],\"name\":\"setForceBatchAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"setForceBatchTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"newIsSequenceWithDataAvailabilityAllowed\",\"type\":\"bool\"}],\"name\":\"switchSequenceWithDataAvailability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x61010060405234801562000011575f80fd5b5060405162004f2a38038062004f2a83398101604081905262000034916200006f565b6001600160a01b0393841660a052918316608052821660c0521660e052620000d4565b6001600160a01b03811681146200006c575f80fd5b50565b5f805f806080858703121562000083575f80fd5b8451620000908162000057565b6020860151909450620000a38162000057565b6040860151909350620000b68162000057565b6060860151909250620000c98162000057565b939692955090935050565b60805160a05160c05160e051614d49620001e15f395f818161053b015281816110b50152818161117c0152818161119e0152818161133a015281816114100152818161157c015281816117c401528181611d96015281816122fa01528181612749015281816127f601528181612e7201528181612f37015281816139dd01528181613a5601528181613a780152613b2001525f81816106dc01528181610bed015281816119b801528181611a8d015281816129c001528181612ac8015261352801525f818161079801528181610c6901528181611c0d0152818161307f01526135a401525f81816107ca015281816108ba015281816111080152818161124a01526130540152614d495ff3fe608060405234801561000f575f80fd5b5060043610610303575f3560e01c8063712570221161019d578063c7fffd4b116100e8578063e57a0b4c11610093578063ecef3f991161006e578063ecef3f9914610833578063f35dda4714610846578063f851a4401461084e575f80fd5b8063e57a0b4c146107ec578063e7a7ed021461080c578063eaeb077b14610820575f80fd5b8063d02103ca116100c3578063d02103ca14610793578063d7bc90ff146107ba578063e46761c4146107c5575f80fd5b8063c7fffd4b14610758578063c89e42df14610760578063cfa8ed4714610773575f80fd5b80639f26f84011610148578063ada8f91911610123578063ada8f91914610711578063b0afe15414610724578063c754c7ed14610730575f80fd5b80639f26f840146106c4578063a3c573eb146106d7578063a652f26c146106fe575f80fd5b80638c3d7301116101785780638c3d73011461068e57806391cafe32146106965780639e001877146106a9575f80fd5b8063712570221461062c5780637a5460c51461063f5780637cd76b8b1461067b575f80fd5b80633cbc795b1161025d57806352bdeb6d116102085780636b8616ce116101e35780636b8616ce146105f15780636e05d2cd146106105780636ff512cc14610619575f80fd5b806352bdeb6d146105a5578063542028d5146105e1578063676870d2146105e9575f80fd5b806349b7b8021161023857806349b7b802146105365780634c21fef31461055d5780634e48770614610592575f80fd5b80633cbc795b1461046857806340b5de6c146104a557806345605267146104fd575f80fd5b806326782247116102bd5780632d72c248116102985780632d72c2481461042257806332c2d153146104355780633c351e1014610448575f80fd5b806326782247146103a85780632acdc2b6146103ed5780632c111c0614610402575f80fd5b806305835f37116102ed57806305835f371461033d578063107bf28c1461038657806311e892d41461038e575f80fd5b8062d0295d146103075780630350896314610322575b5f80fd5b61030f610873565b6040519081526020015b60405180910390f35b61032a602081565b60405161ffff9091168152602001610319565b6103796040518060400160405280600881526020017f80808401c9c3809400000000000000000000000000000000000000000000000081525081565b6040516103199190613ed3565b610379610979565b61039660f981565b60405160ff9091168152602001610319565b6001546103c89073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610319565b6104006103fb366004613efc565b610a05565b005b6008546103c89073ffffffffffffffffffffffffffffffffffffffff1681565b610400610430366004613f8d565b610b23565b61040061044336600461404b565b61140e565b6009546103c89073ffffffffffffffffffffffffffffffffffffffff1681565b6009546104909074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff9091168152602001610319565b6104cc7fff0000000000000000000000000000000000000000000000000000000000000081565b6040517fff000000000000000000000000000000000000000000000000000000000000009091168152602001610319565b60075461051d9068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff9091168152602001610319565b6103c87f000000000000000000000000000000000000000000000000000000000000000081565b600a546105829074010000000000000000000000000000000000000000900460ff1681565b6040519015158152602001610319565b6104006105a036600461408a565b6114dd565b6103796040518060400160405280600281526020017f80b800000000000000000000000000000000000000000000000000000000000081525081565b6103796116ec565b61032a601f81565b61030f6105ff36600461408a565b60066020525f908152604090205481565b61030f60055481565b6104006106273660046140a5565b6116f9565b61040061063a36600461420d565b6117c2565b6103796040518060400160405280600281526020017f80b900000000000000000000000000000000000000000000000000000000000081525081565b6104006106893660046140a5565b611fcb565b610400612094565b6104006106a43660046140a5565b612166565b6103c873a40d5f56745a118d0906a34e69aec8c0db1cb8fa81565b6104006106d23660046142f5565b61227e565b6103c87f000000000000000000000000000000000000000000000000000000000000000081565b61037961070c366004614334565b6128c2565b61040061071f3660046140a5565b612ca0565b61030f6405ca1ab1e081565b60075461051d90700100000000000000000000000000000000900467ffffffffffffffff1681565b61039660e481565b61040061076e3660046143a5565b612d69565b6002546103c89073ffffffffffffffffffffffffffffffffffffffff1681565b6103c87f000000000000000000000000000000000000000000000000000000000000000081565b61030f635ca1ab1e81565b6103c87f000000000000000000000000000000000000000000000000000000000000000081565b600a546103c89073ffffffffffffffffffffffffffffffffffffffff1681565b60075461051d9067ffffffffffffffff1681565b61040061082e3660046143d7565b612dfb565b61040061084136600461441f565b6132c2565b610396601b81565b5f546103c89062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f90819073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa1580156108ff573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109239190614467565b6007549091505f9061094d9067ffffffffffffffff680100000000000000008204811691166144ab565b67ffffffffffffffff169050805f03610968575f9250505090565b61097281836144d3565b9250505090565b600480546109869061450b565b80601f01602080910402602001604051908101604052809291908181526020018280546109b29061450b565b80156109fd5780601f106109d4576101008083540402835291602001916109fd565b820191905f5260205f20905b8154815290600101906020018083116109e057829003601f168201915b505050505081565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314610a5b576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600a5474010000000000000000000000000000000000000000900460ff16151581151503610ab5576040517f5f0e7abe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600a80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1674010000000000000000000000000000000000000000831515021790556040517ff32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41905f90a150565b60025473ffffffffffffffffffffffffffffffffffffffff163314610b74576040517f11e7be1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b835f819003610baf576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8811115610beb576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379e2cf976040518163ffffffff1660e01b81526004015f604051808303815f87803b158015610c50575f80fd5b505af1158015610c62573d5f803e3d5ffd5b505050505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635ca1e1656040518163ffffffff1660e01b8152600401602060405180830381865afa158015610cd0573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cf49190614467565b60075460055491925068010000000000000000900467ffffffffffffffff1690815f805b86811015611028575f8c8c83818110610d3357610d3361455c565b905060800201803603810190610d499190614589565b604081015190915067ffffffffffffffff1615610f365785610d6a816145d5565b9650505f815f0151826020015183604001518460600151604051602001610dcf9493929190938452602084019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166040830152604882015260680190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff8a165f90815260069093529120549091508114610e57576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b85825f0151836020015184604001518f8660600151604051602001610ef0969594939291909586526020860194909452604085019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166060808501919091521b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166068830152607c820152609c0190565b60405160208183030381529060405280519060200120955060065f8867ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f905550611015565b8051604051610f52918591602001918252602082015260400190565b604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018152828252805160209182012084519184018990529183015260608083018a90524260c01b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808401528d901b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660888301525f609c830152935060bc016040516020818303038152906040528051906020012094505b5080611020816145fb565b915050610d18565b5060075467ffffffffffffffff9081169085161115611073576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60058390558567ffffffffffffffff85811690841614611170575f61109884876144ab565b90506110ae67ffffffffffffffff821683614632565b915061112f7f00000000000000000000000000000000000000000000000000000000000000008267ffffffffffffffff166110e7610873565b6110f1919061464b565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169190613326565b50600780547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8816021790555b80156112f957611272337f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663477fa2706040518163ffffffff1660e01b8152600401602060405180830381865afa158015611205573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112299190614467565b611233919061464b565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169291906133fa565b600a546040517f3b51be4b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690633b51be4b906112cc9085908d908d906004016146a9565b5f6040518083038186803b1580156112e2575f80fd5b505afa1580156112f4573d5f803e3d5ffd5b505050505b6040517f9a908e7300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff88166004820152602481018590525f907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303815f875af1158015611395573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113b991906146cb565b90508067ffffffffffffffff167f3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766886040516113f791815260200190565b60405180910390a250505050505050505050505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16331461147d576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff168367ffffffffffffffff167f9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966846040516114d091815260200190565b60405180910390a3505050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314611533576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff8216111561157a576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166315064c966040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115e3573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061160791906146e6565b6116685760075467ffffffffffffffff700100000000000000000000000000000000909104811690821610611668576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600780547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527fa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b906020015b60405180910390a150565b600380546109869061450b565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff16331461174f576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0906020016116e1565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314611831576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f54610100900460ff161580801561184f57505f54600160ff909116105b806118685750303b15801561186857505f5460ff166001145b6118f9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015611955575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b606073ffffffffffffffffffffffffffffffffffffffff851615611bb2576040517fc00f14ab00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063c00f14ab906024015f60405180830381865afa1580156119fc573d5f803e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052611a419190810190614701565b6040517f318aee3d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87811660048301529192505f9182917f00000000000000000000000000000000000000000000000000000000000000009091169063318aee3d906024016040805180830381865afa158015611ad3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611af7919061476a565b915091508163ffffffff165f14611b6e576009805463ffffffff841674010000000000000000000000000000000000000000027fffffffffffffffff00000000000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff841617179055611baf565b600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff89161790555b50505b6009545f90611bf990889073ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900463ffffffff16856128c2565b90505f818051906020012090505f4290505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611c74573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611c989190614467565b90505f808483858f611cab600143614632565b60408051602081019790975286019490945260608086019390935260c09190911b7fffffffffffffffff000000000000000000000000000000000000000000000000166080850152901b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016608883015240609c82015260bc01604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001815290829052805160209091012060058190557f9a908e73000000000000000000000000000000000000000000000000000000008252600160048301526024820181905291507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303815f875af1158015611df1573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e1591906146cb565b508c5f60026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508b60025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508860039081611ea591906147e7565b506004611eb289826147e7565b508c60085f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062069780600760106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507f060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f85838e604051611f52939291906148ff565b60405180910390a15050505050508015611fc2575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314612021576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600a80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a906020016116e1565b60015473ffffffffffffffffffffffffffffffffffffffff1633146120e5576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001545f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff1673ffffffffffffffffffffffffffffffffffffffff9092166201000081029290921790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e9060200160405180910390a1565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff1633146121bc576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60085473ffffffffffffffffffffffffffffffffffffffff1661220b576040517fc89374d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb906020016116e1565b60085473ffffffffffffffffffffffffffffffffffffffff1680158015906122bc575073ffffffffffffffffffffffffffffffffffffffff81163314155b156122f3576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b4262093a807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166330c27dde6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612361573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061238591906146cb565b61238f919061493d565b67ffffffffffffffff1611156123d1576040517f3d49ed4c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815f81900361240c576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8811115612448576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075467ffffffffffffffff808216916124709184916801000000000000000090041661495e565b11156124a8576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6007546005546801000000000000000090910467ffffffffffffffff16905f5b83811015612743575f8787838181106124e3576124e361455c565b90506020028101906124f59190614971565b6124fe906149ad565b90508361250a816145d5565b825180516020918201208185015160408087015160608801519151959a509295505f94612576948794929101938452602084019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166040830152604882015260680190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff89165f908152600690935291205490915081146125fe576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff86165f90815260066020526040812055612622600188614632565b84036126915742600760109054906101000a900467ffffffffffffffff16846040015161264f919061493d565b67ffffffffffffffff161115612691576040517fc44a082100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208381015160408086015160608088015183519586018b90529285018790528481019390935260c01b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808401523390911b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088830152609c82015260bc01604051602081830303815290604052805190602001209450505050808061273b906145fb565b9150506124c8565b506127717f0000000000000000000000000000000000000000000000000000000000000000846110e7610873565b60058190556007805467ffffffffffffffff841668010000000000000000027fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff9091161790556040517f9a908e730000000000000000000000000000000000000000000000000000000081525f9073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690639a908e7390612842908790869060040167ffffffffffffffff929092168252602082015260400190565b6020604051808303815f875af115801561285e573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061288291906146cb565b60405190915067ffffffffffffffff8216907f648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4905f90a250505050505050565b60605f85858573a40d5f56745a118d0906a34e69aec8c0db1cb8fa5f876040516024016128f496959493929190614a18565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff811bff70000000000000000000000000000000000000000000000000000000017905283519091506060905f03612a445760f9601f83516129889190614a7a565b6040518060400160405280600881526020017f80808401c9c380940000000000000000000000000000000000000000000000008152507f00000000000000000000000000000000000000000000000000000000000000006040518060400160405280600281526020017f80b800000000000000000000000000000000000000000000000000000000000081525060e487604051602001612a2e9796959493929190614a95565b6040516020818303038152906040529050612b48565b815161ffff1015612a81576040517f248b8f8200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815160f9612a90602083614a7a565b6040518060400160405280600881526020017f80808401c9c380940000000000000000000000000000000000000000000000008152507f00000000000000000000000000000000000000000000000000000000000000006040518060400160405280600281526020017f80b90000000000000000000000000000000000000000000000000000000000008152508588604051602001612b359796959493929190614b77565b6040516020818303038152906040529150505b8051602080830191909120604080515f80825293810180835292909252601b908201526405ca1ab1e06060820152635ca1ab1e608082015260019060a0016020604051602081039080840390855afa158015612ba6573d5f803e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116612c1e576040517fcd16196600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040515f90612c639084906405ca1ab1e090635ca1ab1e90601b907fff0000000000000000000000000000000000000000000000000000000000000090602001614c59565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529450505050505b949350505050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314612cf6576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6906020016116e1565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314612dbf576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6003612dcb82826147e7565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20816040516116e19190613ed3565b60085473ffffffffffffffffffffffffffffffffffffffff168015801590612e39575073ffffffffffffffffffffffffffffffffffffffff81163314155b15612e70576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166315064c966040518163ffffffff1660e01b8152600401602060405180830381865afa158015612ed9573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612efd91906146e6565b15612f34576040517f39258d1800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663604691696040518163ffffffff1660e01b8152600401602060405180830381865afa158015612f9e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612fc29190614467565b905082811115612ffe576040517f2354600f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61138884111561303a576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61307c73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163330846133fa565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156130e6573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061310a9190614467565b6007805491925067ffffffffffffffff909116905f613128836145d5565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050858560405161315f929190614cb4565b6040519081900390208142613175600143614632565b60408051602081019590955284019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166060830152406068820152608801604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012060075467ffffffffffffffff165f908152600690935291205532330361326b576007546040805183815233602082015260608183018190525f90820152905167ffffffffffffffff909216917ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319181900360800190a26132ba565b60075460405167ffffffffffffffff909116907ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931906132b190849033908b908b90614cc3565b60405180910390a25b505050505050565b600a5474010000000000000000000000000000000000000000900460ff16613316576040517f821935b400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61332183838361345e565b505050565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526133219084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152613bf2565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526134589085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401613378565b50505050565b60025473ffffffffffffffffffffffffffffffffffffffff1633146134af576040517f11e7be1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815f8190036134ea576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8811115613526576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379e2cf976040518163ffffffff1660e01b81526004015f604051808303815f87803b15801561358b575f80fd5b505af115801561359d573d5f803e3d5ffd5b505050505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635ca1e1656040518163ffffffff1660e01b8152600401602060405180830381865afa15801561360b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061362f9190614467565b60075460055491925042916801000000000000000090910467ffffffffffffffff1690815f5b86811015613950575f8a8a838181106136705761367061455c565b90506020028101906136829190614971565b61368b906149ad565b8051805160209091012060408201519192509067ffffffffffffffff161561386b57856136b7816145d5565b9650505f818360200151846040015185606001516040516020016137199493929190938452602084019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166040830152604882015260680190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff8a165f908152600690935291205490915081146137a1576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208381015160408086015160608088015183519586018c90529285018790528481019390935260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166080840152908d901b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088830152609c82015260bc0160405160208183030381529060405280519060200120955060065f8867ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f90555061393b565b8151516201d4c010156138aa576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080516020810187905290810182905260608082018a905260c089901b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808301528b901b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660888201525f609c82015260bc016040516020818303038152906040528051906020012094505b50508080613948906145fb565b915050613655565b5060075467ffffffffffffffff908116908416111561399b576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60058290558567ffffffffffffffff84811690831614613a50575f6139c083866144ab565b90506139d667ffffffffffffffff821683614632565b9150613a0f7f00000000000000000000000000000000000000000000000000000000000000008267ffffffffffffffff166110e7610873565b50600780547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8716021790555b613adf337f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663477fa2706040518163ffffffff1660e01b8152600401602060405180830381865afa158015611205573d5f803e3d5ffd5b6040517f9a908e7300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff88166004820152602481018490525f907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303815f875af1158015613b7b573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613b9f91906146cb565b90508067ffffffffffffffff167f3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e76688604051613bdd91815260200190565b60405180910390a25050505050505050505050565b5f613c53826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16613cfd9092919063ffffffff16565b8051909150156133215780806020019051810190613c7191906146e6565b613321576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016118f0565b6060612c9884845f85855f808673ffffffffffffffffffffffffffffffffffffffff168587604051613d2f9190614d02565b5f6040518083038185875af1925050503d805f8114613d69576040519150601f19603f3d011682016040523d82523d5f602084013e613d6e565b606091505b5091509150613d7f87838387613d8a565b979650505050505050565b60608315613e1f5782515f03613e185773ffffffffffffffffffffffffffffffffffffffff85163b613e18576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016118f0565b5081612c98565b612c988383815115613e345781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118f09190613ed3565b5f5b83811015613e82578181015183820152602001613e6a565b50505f910152565b5f8151808452613ea1816020860160208601613e68565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081525f613ee56020830184613e8a565b9392505050565b8015158114613ef9575f80fd5b50565b5f60208284031215613f0c575f80fd5b8135613ee581613eec565b73ffffffffffffffffffffffffffffffffffffffff81168114613ef9575f80fd5b8035613f4381613f17565b919050565b5f8083601f840112613f58575f80fd5b50813567ffffffffffffffff811115613f6f575f80fd5b602083019150836020828501011115613f86575f80fd5b9250929050565b5f805f805f60608688031215613fa1575f80fd5b853567ffffffffffffffff80821115613fb8575f80fd5b818801915088601f830112613fcb575f80fd5b813581811115613fd9575f80fd5b8960208260071b8501011115613fed575f80fd5b6020830197508096505061400360208901613f38565b94506040880135915080821115614018575f80fd5b5061402588828901613f48565b969995985093965092949392505050565b67ffffffffffffffff81168114613ef9575f80fd5b5f805f6060848603121561405d575f80fd5b833561406881614036565b925060208401359150604084013561407f81613f17565b809150509250925092565b5f6020828403121561409a575f80fd5b8135613ee581614036565b5f602082840312156140b5575f80fd5b8135613ee581613f17565b63ffffffff81168114613ef9575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6040516080810167ffffffffffffffff81118282101715614121576141216140d1565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561416e5761416e6140d1565b604052919050565b5f67ffffffffffffffff82111561418f5761418f6140d1565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b5f82601f8301126141ca575f80fd5b81356141dd6141d882614176565b614127565b8181528460208386010111156141f1575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f805f8060c08789031215614222575f80fd5b863561422d81613f17565b9550602087013561423d81613f17565b9450604087013561424d816140c0565b9350606087013561425d81613f17565b9250608087013567ffffffffffffffff80821115614279575f80fd5b6142858a838b016141bb565b935060a089013591508082111561429a575f80fd5b506142a789828a016141bb565b9150509295509295509295565b5f8083601f8401126142c4575f80fd5b50813567ffffffffffffffff8111156142db575f80fd5b6020830191508360208260051b8501011115613f86575f80fd5b5f8060208385031215614306575f80fd5b823567ffffffffffffffff81111561431c575f80fd5b614328858286016142b4565b90969095509350505050565b5f805f8060808587031215614347575f80fd5b8435614352816140c0565b9350602085013561436281613f17565b92506040850135614372816140c0565b9150606085013567ffffffffffffffff81111561438d575f80fd5b614399878288016141bb565b91505092959194509250565b5f602082840312156143b5575f80fd5b813567ffffffffffffffff8111156143cb575f80fd5b612c98848285016141bb565b5f805f604084860312156143e9575f80fd5b833567ffffffffffffffff8111156143ff575f80fd5b61440b86828701613f48565b909790965060209590950135949350505050565b5f805f60408486031215614431575f80fd5b833567ffffffffffffffff811115614447575f80fd5b614453868287016142b4565b909450925050602084013561407f81613f17565b5f60208284031215614477575f80fd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b67ffffffffffffffff8281168282160390808211156144cc576144cc61447e565b5092915050565b5f82614506577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b600181811c9082168061451f57607f821691505b602082108103614556577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f60808284031215614599575f80fd5b6145a16140fe565b823581526020830135602082015260408301356145bd81614036565b60408201526060928301359281019290925250919050565b5f67ffffffffffffffff8083168181036145f1576145f161447e565b6001019392505050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361462b5761462b61447e565b5060010190565b818103818111156146455761464561447e565b92915050565b80820281158282048414176146455761464561447e565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b838152604060208201525f6146c2604083018486614662565b95945050505050565b5f602082840312156146db575f80fd5b8151613ee581614036565b5f602082840312156146f6575f80fd5b8151613ee581613eec565b5f60208284031215614711575f80fd5b815167ffffffffffffffff811115614727575f80fd5b8201601f81018413614737575f80fd5b80516147456141d882614176565b818152856020838501011115614759575f80fd5b6146c2826020830160208601613e68565b5f806040838503121561477b575f80fd5b8251614786816140c0565b602084015190925061479781613f17565b809150509250929050565b601f821115613321575f81815260208120601f850160051c810160208610156147c85750805b601f850160051c820191505b818110156132ba578281556001016147d4565b815167ffffffffffffffff811115614801576148016140d1565b6148158161480f845461450b565b846147a2565b602080601f831160018114614867575f84156148315750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b1785556132ba565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b828110156148b357888601518255948401946001909101908401614894565b50858210156148ef57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b606081525f6149116060830186613e8a565b905083602083015273ffffffffffffffffffffffffffffffffffffffff83166040830152949350505050565b67ffffffffffffffff8181168382160190808211156144cc576144cc61447e565b808201808211156146455761464561447e565b5f82357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff818336030181126149a3575f80fd5b9190910192915050565b5f608082360312156149bd575f80fd5b6149c56140fe565b823567ffffffffffffffff8111156149db575f80fd5b6149e7368286016141bb565b825250602083013560208201526040830135614a0281614036565b6040820152606092830135928101929092525090565b5f63ffffffff808916835273ffffffffffffffffffffffffffffffffffffffff8089166020850152818816604085015280871660608501528086166080850152505060c060a0830152614a6e60c0830184613e8a565b98975050505050505050565b61ffff8181168382160190808211156144cc576144cc61447e565b5f7fff00000000000000000000000000000000000000000000000000000000000000808a60f81b1683527fffff0000000000000000000000000000000000000000000000000000000000008960f01b1660018401528751614afd816003860160208c01613e68565b80840190507fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008860601b1660038201528651614b40816017840160208b01613e68565b808201915050818660f81b16601782015284519150614b66826018830160208801613e68565b016018019998505050505050505050565b7fff000000000000000000000000000000000000000000000000000000000000008860f81b1681525f7fffff000000000000000000000000000000000000000000000000000000000000808960f01b1660018401528751614bdf816003860160208c01613e68565b80840190507fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008860601b1660038201528651614c22816017840160208b01613e68565b808201915050818660f01b16601782015284519150614c48826019830160208801613e68565b016019019998505050505050505050565b5f8651614c6a818460208b01613e68565b9190910194855250602084019290925260f81b7fff000000000000000000000000000000000000000000000000000000000000009081166040840152166041820152604201919050565b818382375f9101908152919050565b84815273ffffffffffffffffffffffffffffffffffffffff84166020820152606060408201525f614cf8606083018486614662565b9695505050505050565b5f82516149a3818460208701613e6856fea2646970667358221220ef9322cb79865871b6601d2405fd457dd4ecbf5ae96146f22f35277060e55e2264736f6c63430008140033",
}

// PolygonvalidiumetrogpreviousABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonvalidiumetrogpreviousMetaData.ABI instead.
var PolygonvalidiumetrogpreviousABI = PolygonvalidiumetrogpreviousMetaData.ABI

// PolygonvalidiumetrogpreviousBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygonvalidiumetrogpreviousMetaData.Bin instead.
var PolygonvalidiumetrogpreviousBin = PolygonvalidiumetrogpreviousMetaData.Bin

// DeployPolygonvalidiumetrogprevious deploys a new Ethereum contract, binding an instance of Polygonvalidiumetrogprevious to it.
func DeployPolygonvalidiumetrogprevious(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address, _rollupManager common.Address) (common.Address, *types.Transaction, *Polygonvalidiumetrogprevious, error) {
	parsed, err := PolygonvalidiumetrogpreviousMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygonvalidiumetrogpreviousBin), backend, _globalExitRootManager, _pol, _bridgeAddress, _rollupManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonvalidiumetrogprevious{PolygonvalidiumetrogpreviousCaller: PolygonvalidiumetrogpreviousCaller{contract: contract}, PolygonvalidiumetrogpreviousTransactor: PolygonvalidiumetrogpreviousTransactor{contract: contract}, PolygonvalidiumetrogpreviousFilterer: PolygonvalidiumetrogpreviousFilterer{contract: contract}}, nil
}

// Polygonvalidiumetrogprevious is an auto generated Go binding around an Ethereum contract.
type Polygonvalidiumetrogprevious struct {
	PolygonvalidiumetrogpreviousCaller     // Read-only binding to the contract
	PolygonvalidiumetrogpreviousTransactor // Write-only binding to the contract
	PolygonvalidiumetrogpreviousFilterer   // Log filterer for contract events
}

// PolygonvalidiumetrogpreviousCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonvalidiumetrogpreviousCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonvalidiumetrogpreviousTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonvalidiumetrogpreviousTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonvalidiumetrogpreviousFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonvalidiumetrogpreviousFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonvalidiumetrogpreviousSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonvalidiumetrogpreviousSession struct {
	Contract     *Polygonvalidiumetrogprevious // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// PolygonvalidiumetrogpreviousCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonvalidiumetrogpreviousCallerSession struct {
	Contract *PolygonvalidiumetrogpreviousCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// PolygonvalidiumetrogpreviousTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonvalidiumetrogpreviousTransactorSession struct {
	Contract     *PolygonvalidiumetrogpreviousTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// PolygonvalidiumetrogpreviousRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonvalidiumetrogpreviousRaw struct {
	Contract *Polygonvalidiumetrogprevious // Generic contract binding to access the raw methods on
}

// PolygonvalidiumetrogpreviousCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonvalidiumetrogpreviousCallerRaw struct {
	Contract *PolygonvalidiumetrogpreviousCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonvalidiumetrogpreviousTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonvalidiumetrogpreviousTransactorRaw struct {
	Contract *PolygonvalidiumetrogpreviousTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonvalidiumetrogprevious creates a new instance of Polygonvalidiumetrogprevious, bound to a specific deployed contract.
func NewPolygonvalidiumetrogprevious(address common.Address, backend bind.ContractBackend) (*Polygonvalidiumetrogprevious, error) {
	contract, err := bindPolygonvalidiumetrogprevious(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonvalidiumetrogprevious{PolygonvalidiumetrogpreviousCaller: PolygonvalidiumetrogpreviousCaller{contract: contract}, PolygonvalidiumetrogpreviousTransactor: PolygonvalidiumetrogpreviousTransactor{contract: contract}, PolygonvalidiumetrogpreviousFilterer: PolygonvalidiumetrogpreviousFilterer{contract: contract}}, nil
}

// NewPolygonvalidiumetrogpreviousCaller creates a new read-only instance of Polygonvalidiumetrogprevious, bound to a specific deployed contract.
func NewPolygonvalidiumetrogpreviousCaller(address common.Address, caller bind.ContractCaller) (*PolygonvalidiumetrogpreviousCaller, error) {
	contract, err := bindPolygonvalidiumetrogprevious(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogpreviousCaller{contract: contract}, nil
}

// NewPolygonvalidiumetrogpreviousTransactor creates a new write-only instance of Polygonvalidiumetrogprevious, bound to a specific deployed contract.
func NewPolygonvalidiumetrogpreviousTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonvalidiumetrogpreviousTransactor, error) {
	contract, err := bindPolygonvalidiumetrogprevious(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogpreviousTransactor{contract: contract}, nil
}

// NewPolygonvalidiumetrogpreviousFilterer creates a new log filterer instance of Polygonvalidiumetrogprevious, bound to a specific deployed contract.
func NewPolygonvalidiumetrogpreviousFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonvalidiumetrogpreviousFilterer, error) {
	contract, err := bindPolygonvalidiumetrogprevious(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogpreviousFilterer{contract: contract}, nil
}

// bindPolygonvalidiumetrogprevious binds a generic wrapper to an already deployed contract.
func bindPolygonvalidiumetrogprevious(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonvalidiumetrogpreviousMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonvalidiumetrogprevious.Contract.PolygonvalidiumetrogpreviousCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.PolygonvalidiumetrogpreviousTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.PolygonvalidiumetrogpreviousTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonvalidiumetrogprevious.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.contract.Transact(opts, method, params...)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) GLOBALEXITROOTMANAGERL2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "GLOBAL_EXIT_ROOT_MANAGER_L2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _Polygonvalidiumetrogprevious.Contract.GLOBALEXITROOTMANAGERL2(&_Polygonvalidiumetrogprevious.CallOpts)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _Polygonvalidiumetrogprevious.Contract.GLOBALEXITROOTMANAGERL2(&_Polygonvalidiumetrogprevious.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) INITIALIZETXBRIDGELISTLENLEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_LIST_LEN_LEN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _Polygonvalidiumetrogprevious.Contract.INITIALIZETXBRIDGELISTLENLEN(&_Polygonvalidiumetrogprevious.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _Polygonvalidiumetrogprevious.Contract.INITIALIZETXBRIDGELISTLENLEN(&_Polygonvalidiumetrogprevious.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) INITIALIZETXBRIDGEPARAMS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _Polygonvalidiumetrogprevious.Contract.INITIALIZETXBRIDGEPARAMS(&_Polygonvalidiumetrogprevious.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _Polygonvalidiumetrogprevious.Contract.INITIALIZETXBRIDGEPARAMS(&_Polygonvalidiumetrogprevious.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _Polygonvalidiumetrogprevious.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_Polygonvalidiumetrogprevious.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _Polygonvalidiumetrogprevious.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_Polygonvalidiumetrogprevious.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _Polygonvalidiumetrogprevious.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_Polygonvalidiumetrogprevious.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _Polygonvalidiumetrogprevious.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_Polygonvalidiumetrogprevious.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) INITIALIZETXCONSTANTBYTES(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _Polygonvalidiumetrogprevious.Contract.INITIALIZETXCONSTANTBYTES(&_Polygonvalidiumetrogprevious.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _Polygonvalidiumetrogprevious.Contract.INITIALIZETXCONSTANTBYTES(&_Polygonvalidiumetrogprevious.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) INITIALIZETXCONSTANTBYTESEMPTYMETADATA(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _Polygonvalidiumetrogprevious.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_Polygonvalidiumetrogprevious.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _Polygonvalidiumetrogprevious.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_Polygonvalidiumetrogprevious.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) INITIALIZETXDATALENEMPTYMETADATA(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "INITIALIZE_TX_DATA_LEN_EMPTY_METADATA")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _Polygonvalidiumetrogprevious.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_Polygonvalidiumetrogprevious.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _Polygonvalidiumetrogprevious.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_Polygonvalidiumetrogprevious.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) INITIALIZETXEFFECTIVEPERCENTAGE(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "INITIALIZE_TX_EFFECTIVE_PERCENTAGE")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _Polygonvalidiumetrogprevious.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_Polygonvalidiumetrogprevious.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _Polygonvalidiumetrogprevious.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_Polygonvalidiumetrogprevious.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) SIGNATUREINITIALIZETXR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_R")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _Polygonvalidiumetrogprevious.Contract.SIGNATUREINITIALIZETXR(&_Polygonvalidiumetrogprevious.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _Polygonvalidiumetrogprevious.Contract.SIGNATUREINITIALIZETXR(&_Polygonvalidiumetrogprevious.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) SIGNATUREINITIALIZETXS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_S")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _Polygonvalidiumetrogprevious.Contract.SIGNATUREINITIALIZETXS(&_Polygonvalidiumetrogprevious.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _Polygonvalidiumetrogprevious.Contract.SIGNATUREINITIALIZETXS(&_Polygonvalidiumetrogprevious.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) SIGNATUREINITIALIZETXV(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_V")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _Polygonvalidiumetrogprevious.Contract.SIGNATUREINITIALIZETXV(&_Polygonvalidiumetrogprevious.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _Polygonvalidiumetrogprevious.Contract.SIGNATUREINITIALIZETXV(&_Polygonvalidiumetrogprevious.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) Admin() (common.Address, error) {
	return _Polygonvalidiumetrogprevious.Contract.Admin(&_Polygonvalidiumetrogprevious.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) Admin() (common.Address, error) {
	return _Polygonvalidiumetrogprevious.Contract.Admin(&_Polygonvalidiumetrogprevious.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) BridgeAddress() (common.Address, error) {
	return _Polygonvalidiumetrogprevious.Contract.BridgeAddress(&_Polygonvalidiumetrogprevious.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) BridgeAddress() (common.Address, error) {
	return _Polygonvalidiumetrogprevious.Contract.BridgeAddress(&_Polygonvalidiumetrogprevious.CallOpts)
}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) CalculatePolPerForceBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "calculatePolPerForceBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) CalculatePolPerForceBatch() (*big.Int, error) {
	return _Polygonvalidiumetrogprevious.Contract.CalculatePolPerForceBatch(&_Polygonvalidiumetrogprevious.CallOpts)
}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) CalculatePolPerForceBatch() (*big.Int, error) {
	return _Polygonvalidiumetrogprevious.Contract.CalculatePolPerForceBatch(&_Polygonvalidiumetrogprevious.CallOpts)
}

// DataAvailabilityProtocol is a free data retrieval call binding the contract method 0xe57a0b4c.
//
// Solidity: function dataAvailabilityProtocol() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) DataAvailabilityProtocol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "dataAvailabilityProtocol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataAvailabilityProtocol is a free data retrieval call binding the contract method 0xe57a0b4c.
//
// Solidity: function dataAvailabilityProtocol() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) DataAvailabilityProtocol() (common.Address, error) {
	return _Polygonvalidiumetrogprevious.Contract.DataAvailabilityProtocol(&_Polygonvalidiumetrogprevious.CallOpts)
}

// DataAvailabilityProtocol is a free data retrieval call binding the contract method 0xe57a0b4c.
//
// Solidity: function dataAvailabilityProtocol() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) DataAvailabilityProtocol() (common.Address, error) {
	return _Polygonvalidiumetrogprevious.Contract.DataAvailabilityProtocol(&_Polygonvalidiumetrogprevious.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) ForceBatchAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "forceBatchAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) ForceBatchAddress() (common.Address, error) {
	return _Polygonvalidiumetrogprevious.Contract.ForceBatchAddress(&_Polygonvalidiumetrogprevious.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) ForceBatchAddress() (common.Address, error) {
	return _Polygonvalidiumetrogprevious.Contract.ForceBatchAddress(&_Polygonvalidiumetrogprevious.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonvalidiumetrogprevious.Contract.ForceBatchTimeout(&_Polygonvalidiumetrogprevious.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonvalidiumetrogprevious.Contract.ForceBatchTimeout(&_Polygonvalidiumetrogprevious.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonvalidiumetrogprevious.Contract.ForcedBatches(&_Polygonvalidiumetrogprevious.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonvalidiumetrogprevious.Contract.ForcedBatches(&_Polygonvalidiumetrogprevious.CallOpts, arg0)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) GasTokenAddress() (common.Address, error) {
	return _Polygonvalidiumetrogprevious.Contract.GasTokenAddress(&_Polygonvalidiumetrogprevious.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) GasTokenAddress() (common.Address, error) {
	return _Polygonvalidiumetrogprevious.Contract.GasTokenAddress(&_Polygonvalidiumetrogprevious.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) GasTokenNetwork() (uint32, error) {
	return _Polygonvalidiumetrogprevious.Contract.GasTokenNetwork(&_Polygonvalidiumetrogprevious.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) GasTokenNetwork() (uint32, error) {
	return _Polygonvalidiumetrogprevious.Contract.GasTokenNetwork(&_Polygonvalidiumetrogprevious.CallOpts)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) GenerateInitializeTransaction(opts *bind.CallOpts, networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "generateInitializeTransaction", networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) GenerateInitializeTransaction(networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _Polygonvalidiumetrogprevious.Contract.GenerateInitializeTransaction(&_Polygonvalidiumetrogprevious.CallOpts, networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) GenerateInitializeTransaction(networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _Polygonvalidiumetrogprevious.Contract.GenerateInitializeTransaction(&_Polygonvalidiumetrogprevious.CallOpts, networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonvalidiumetrogprevious.Contract.GlobalExitRootManager(&_Polygonvalidiumetrogprevious.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonvalidiumetrogprevious.Contract.GlobalExitRootManager(&_Polygonvalidiumetrogprevious.CallOpts)
}

// IsSequenceWithDataAvailabilityAllowed is a free data retrieval call binding the contract method 0x4c21fef3.
//
// Solidity: function isSequenceWithDataAvailabilityAllowed() view returns(bool)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) IsSequenceWithDataAvailabilityAllowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "isSequenceWithDataAvailabilityAllowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSequenceWithDataAvailabilityAllowed is a free data retrieval call binding the contract method 0x4c21fef3.
//
// Solidity: function isSequenceWithDataAvailabilityAllowed() view returns(bool)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) IsSequenceWithDataAvailabilityAllowed() (bool, error) {
	return _Polygonvalidiumetrogprevious.Contract.IsSequenceWithDataAvailabilityAllowed(&_Polygonvalidiumetrogprevious.CallOpts)
}

// IsSequenceWithDataAvailabilityAllowed is a free data retrieval call binding the contract method 0x4c21fef3.
//
// Solidity: function isSequenceWithDataAvailabilityAllowed() view returns(bool)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) IsSequenceWithDataAvailabilityAllowed() (bool, error) {
	return _Polygonvalidiumetrogprevious.Contract.IsSequenceWithDataAvailabilityAllowed(&_Polygonvalidiumetrogprevious.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) LastAccInputHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "lastAccInputHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) LastAccInputHash() ([32]byte, error) {
	return _Polygonvalidiumetrogprevious.Contract.LastAccInputHash(&_Polygonvalidiumetrogprevious.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) LastAccInputHash() ([32]byte, error) {
	return _Polygonvalidiumetrogprevious.Contract.LastAccInputHash(&_Polygonvalidiumetrogprevious.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) LastForceBatch() (uint64, error) {
	return _Polygonvalidiumetrogprevious.Contract.LastForceBatch(&_Polygonvalidiumetrogprevious.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) LastForceBatch() (uint64, error) {
	return _Polygonvalidiumetrogprevious.Contract.LastForceBatch(&_Polygonvalidiumetrogprevious.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonvalidiumetrogprevious.Contract.LastForceBatchSequenced(&_Polygonvalidiumetrogprevious.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonvalidiumetrogprevious.Contract.LastForceBatchSequenced(&_Polygonvalidiumetrogprevious.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) NetworkName() (string, error) {
	return _Polygonvalidiumetrogprevious.Contract.NetworkName(&_Polygonvalidiumetrogprevious.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) NetworkName() (string, error) {
	return _Polygonvalidiumetrogprevious.Contract.NetworkName(&_Polygonvalidiumetrogprevious.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) PendingAdmin() (common.Address, error) {
	return _Polygonvalidiumetrogprevious.Contract.PendingAdmin(&_Polygonvalidiumetrogprevious.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) PendingAdmin() (common.Address, error) {
	return _Polygonvalidiumetrogprevious.Contract.PendingAdmin(&_Polygonvalidiumetrogprevious.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) Pol() (common.Address, error) {
	return _Polygonvalidiumetrogprevious.Contract.Pol(&_Polygonvalidiumetrogprevious.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) Pol() (common.Address, error) {
	return _Polygonvalidiumetrogprevious.Contract.Pol(&_Polygonvalidiumetrogprevious.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) RollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "rollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) RollupManager() (common.Address, error) {
	return _Polygonvalidiumetrogprevious.Contract.RollupManager(&_Polygonvalidiumetrogprevious.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) RollupManager() (common.Address, error) {
	return _Polygonvalidiumetrogprevious.Contract.RollupManager(&_Polygonvalidiumetrogprevious.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) TrustedSequencer() (common.Address, error) {
	return _Polygonvalidiumetrogprevious.Contract.TrustedSequencer(&_Polygonvalidiumetrogprevious.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) TrustedSequencer() (common.Address, error) {
	return _Polygonvalidiumetrogprevious.Contract.TrustedSequencer(&_Polygonvalidiumetrogprevious.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCaller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonvalidiumetrogprevious.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) TrustedSequencerURL() (string, error) {
	return _Polygonvalidiumetrogprevious.Contract.TrustedSequencerURL(&_Polygonvalidiumetrogprevious.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousCallerSession) TrustedSequencerURL() (string, error) {
	return _Polygonvalidiumetrogprevious.Contract.TrustedSequencerURL(&_Polygonvalidiumetrogprevious.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.AcceptAdminRole(&_Polygonvalidiumetrogprevious.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.AcceptAdminRole(&_Polygonvalidiumetrogprevious.TransactOpts)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactor) ForceBatch(opts *bind.TransactOpts, transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.contract.Transact(opts, "forceBatch", transactions, polAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) ForceBatch(transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.ForceBatch(&_Polygonvalidiumetrogprevious.TransactOpts, transactions, polAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactorSession) ForceBatch(transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.ForceBatch(&_Polygonvalidiumetrogprevious.TransactOpts, transactions, polAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.contract.Transact(opts, "initialize", _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.Initialize(&_Polygonvalidiumetrogprevious.TransactOpts, _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactorSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.Initialize(&_Polygonvalidiumetrogprevious.TransactOpts, _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactor) OnVerifyBatches(opts *bind.TransactOpts, lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.contract.Transact(opts, "onVerifyBatches", lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.OnVerifyBatches(&_Polygonvalidiumetrogprevious.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactorSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.OnVerifyBatches(&_Polygonvalidiumetrogprevious.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xecef3f99.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, address l2Coinbase) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactor) SequenceBatches(opts *bind.TransactOpts, batches []PolygonRollupBaseEtrogPreviousBatchData, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.contract.Transact(opts, "sequenceBatches", batches, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xecef3f99.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, address l2Coinbase) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) SequenceBatches(batches []PolygonRollupBaseEtrogPreviousBatchData, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.SequenceBatches(&_Polygonvalidiumetrogprevious.TransactOpts, batches, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xecef3f99.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, address l2Coinbase) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactorSession) SequenceBatches(batches []PolygonRollupBaseEtrogPreviousBatchData, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.SequenceBatches(&_Polygonvalidiumetrogprevious.TransactOpts, batches, l2Coinbase)
}

// SequenceBatchesValidium is a paid mutator transaction binding the contract method 0x2d72c248.
//
// Solidity: function sequenceBatchesValidium((bytes32,bytes32,uint64,bytes32)[] batches, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactor) SequenceBatchesValidium(opts *bind.TransactOpts, batches []PolygonValidiumEtrogPreviousValidiumBatchData, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.contract.Transact(opts, "sequenceBatchesValidium", batches, l2Coinbase, dataAvailabilityMessage)
}

// SequenceBatchesValidium is a paid mutator transaction binding the contract method 0x2d72c248.
//
// Solidity: function sequenceBatchesValidium((bytes32,bytes32,uint64,bytes32)[] batches, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) SequenceBatchesValidium(batches []PolygonValidiumEtrogPreviousValidiumBatchData, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.SequenceBatchesValidium(&_Polygonvalidiumetrogprevious.TransactOpts, batches, l2Coinbase, dataAvailabilityMessage)
}

// SequenceBatchesValidium is a paid mutator transaction binding the contract method 0x2d72c248.
//
// Solidity: function sequenceBatchesValidium((bytes32,bytes32,uint64,bytes32)[] batches, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactorSession) SequenceBatchesValidium(batches []PolygonValidiumEtrogPreviousValidiumBatchData, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.SequenceBatchesValidium(&_Polygonvalidiumetrogprevious.TransactOpts, batches, l2Coinbase, dataAvailabilityMessage)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactor) SequenceForceBatches(opts *bind.TransactOpts, batches []PolygonRollupBaseEtrogPreviousBatchData) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.contract.Transact(opts, "sequenceForceBatches", batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) SequenceForceBatches(batches []PolygonRollupBaseEtrogPreviousBatchData) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.SequenceForceBatches(&_Polygonvalidiumetrogprevious.TransactOpts, batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactorSession) SequenceForceBatches(batches []PolygonRollupBaseEtrogPreviousBatchData) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.SequenceForceBatches(&_Polygonvalidiumetrogprevious.TransactOpts, batches)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactor) SetDataAvailabilityProtocol(opts *bind.TransactOpts, newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.contract.Transact(opts, "setDataAvailabilityProtocol", newDataAvailabilityProtocol)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) SetDataAvailabilityProtocol(newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.SetDataAvailabilityProtocol(&_Polygonvalidiumetrogprevious.TransactOpts, newDataAvailabilityProtocol)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactorSession) SetDataAvailabilityProtocol(newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.SetDataAvailabilityProtocol(&_Polygonvalidiumetrogprevious.TransactOpts, newDataAvailabilityProtocol)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactor) SetForceBatchAddress(opts *bind.TransactOpts, newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.contract.Transact(opts, "setForceBatchAddress", newForceBatchAddress)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) SetForceBatchAddress(newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.SetForceBatchAddress(&_Polygonvalidiumetrogprevious.TransactOpts, newForceBatchAddress)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactorSession) SetForceBatchAddress(newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.SetForceBatchAddress(&_Polygonvalidiumetrogprevious.TransactOpts, newForceBatchAddress)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactor) SetForceBatchTimeout(opts *bind.TransactOpts, newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.contract.Transact(opts, "setForceBatchTimeout", newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.SetForceBatchTimeout(&_Polygonvalidiumetrogprevious.TransactOpts, newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactorSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.SetForceBatchTimeout(&_Polygonvalidiumetrogprevious.TransactOpts, newforceBatchTimeout)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.SetTrustedSequencer(&_Polygonvalidiumetrogprevious.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.SetTrustedSequencer(&_Polygonvalidiumetrogprevious.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.SetTrustedSequencerURL(&_Polygonvalidiumetrogprevious.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.SetTrustedSequencerURL(&_Polygonvalidiumetrogprevious.TransactOpts, newTrustedSequencerURL)
}

// SwitchSequenceWithDataAvailability is a paid mutator transaction binding the contract method 0x2acdc2b6.
//
// Solidity: function switchSequenceWithDataAvailability(bool newIsSequenceWithDataAvailabilityAllowed) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactor) SwitchSequenceWithDataAvailability(opts *bind.TransactOpts, newIsSequenceWithDataAvailabilityAllowed bool) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.contract.Transact(opts, "switchSequenceWithDataAvailability", newIsSequenceWithDataAvailabilityAllowed)
}

// SwitchSequenceWithDataAvailability is a paid mutator transaction binding the contract method 0x2acdc2b6.
//
// Solidity: function switchSequenceWithDataAvailability(bool newIsSequenceWithDataAvailabilityAllowed) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) SwitchSequenceWithDataAvailability(newIsSequenceWithDataAvailabilityAllowed bool) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.SwitchSequenceWithDataAvailability(&_Polygonvalidiumetrogprevious.TransactOpts, newIsSequenceWithDataAvailabilityAllowed)
}

// SwitchSequenceWithDataAvailability is a paid mutator transaction binding the contract method 0x2acdc2b6.
//
// Solidity: function switchSequenceWithDataAvailability(bool newIsSequenceWithDataAvailabilityAllowed) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactorSession) SwitchSequenceWithDataAvailability(newIsSequenceWithDataAvailabilityAllowed bool) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.SwitchSequenceWithDataAvailability(&_Polygonvalidiumetrogprevious.TransactOpts, newIsSequenceWithDataAvailabilityAllowed)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.TransferAdminRole(&_Polygonvalidiumetrogprevious.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonvalidiumetrogprevious.Contract.TransferAdminRole(&_Polygonvalidiumetrogprevious.TransactOpts, newPendingAdmin)
}

// PolygonvalidiumetrogpreviousAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousAcceptAdminRoleIterator struct {
	Event *PolygonvalidiumetrogpreviousAcceptAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogpreviousAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogpreviousAcceptAdminRole)
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
		it.Event = new(PolygonvalidiumetrogpreviousAcceptAdminRole)
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
func (it *PolygonvalidiumetrogpreviousAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogpreviousAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogpreviousAcceptAdminRole represents a AcceptAdminRole event raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*PolygonvalidiumetrogpreviousAcceptAdminRoleIterator, error) {

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogpreviousAcceptAdminRoleIterator{contract: _Polygonvalidiumetrogprevious.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogpreviousAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogpreviousAcceptAdminRole)
				if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
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
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) ParseAcceptAdminRole(log types.Log) (*PolygonvalidiumetrogpreviousAcceptAdminRole, error) {
	event := new(PolygonvalidiumetrogpreviousAcceptAdminRole)
	if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumetrogpreviousForceBatchIterator is returned from FilterForceBatch and is used to iterate over the raw logs and unpacked data for ForceBatch events raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousForceBatchIterator struct {
	Event *PolygonvalidiumetrogpreviousForceBatch // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogpreviousForceBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogpreviousForceBatch)
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
		it.Event = new(PolygonvalidiumetrogpreviousForceBatch)
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
func (it *PolygonvalidiumetrogpreviousForceBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogpreviousForceBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogpreviousForceBatch represents a ForceBatch event raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousForceBatch struct {
	ForceBatchNum      uint64
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Transactions       []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterForceBatch is a free log retrieval operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) FilterForceBatch(opts *bind.FilterOpts, forceBatchNum []uint64) (*PolygonvalidiumetrogpreviousForceBatchIterator, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.FilterLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogpreviousForceBatchIterator{contract: _Polygonvalidiumetrogprevious.contract, event: "ForceBatch", logs: logs, sub: sub}, nil
}

// WatchForceBatch is a free log subscription operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) WatchForceBatch(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogpreviousForceBatch, forceBatchNum []uint64) (event.Subscription, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.WatchLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogpreviousForceBatch)
				if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "ForceBatch", log); err != nil {
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
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) ParseForceBatch(log types.Log) (*PolygonvalidiumetrogpreviousForceBatch, error) {
	event := new(PolygonvalidiumetrogpreviousForceBatch)
	if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "ForceBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumetrogpreviousInitialSequenceBatchesIterator is returned from FilterInitialSequenceBatches and is used to iterate over the raw logs and unpacked data for InitialSequenceBatches events raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousInitialSequenceBatchesIterator struct {
	Event *PolygonvalidiumetrogpreviousInitialSequenceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogpreviousInitialSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogpreviousInitialSequenceBatches)
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
		it.Event = new(PolygonvalidiumetrogpreviousInitialSequenceBatches)
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
func (it *PolygonvalidiumetrogpreviousInitialSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogpreviousInitialSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogpreviousInitialSequenceBatches represents a InitialSequenceBatches event raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousInitialSequenceBatches struct {
	Transactions       []byte
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterInitialSequenceBatches is a free log retrieval operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) FilterInitialSequenceBatches(opts *bind.FilterOpts) (*PolygonvalidiumetrogpreviousInitialSequenceBatchesIterator, error) {

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.FilterLogs(opts, "InitialSequenceBatches")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogpreviousInitialSequenceBatchesIterator{contract: _Polygonvalidiumetrogprevious.contract, event: "InitialSequenceBatches", logs: logs, sub: sub}, nil
}

// WatchInitialSequenceBatches is a free log subscription operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) WatchInitialSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogpreviousInitialSequenceBatches) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.WatchLogs(opts, "InitialSequenceBatches")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogpreviousInitialSequenceBatches)
				if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "InitialSequenceBatches", log); err != nil {
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
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) ParseInitialSequenceBatches(log types.Log) (*PolygonvalidiumetrogpreviousInitialSequenceBatches, error) {
	event := new(PolygonvalidiumetrogpreviousInitialSequenceBatches)
	if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "InitialSequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumetrogpreviousInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousInitializedIterator struct {
	Event *PolygonvalidiumetrogpreviousInitialized // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogpreviousInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogpreviousInitialized)
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
		it.Event = new(PolygonvalidiumetrogpreviousInitialized)
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
func (it *PolygonvalidiumetrogpreviousInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogpreviousInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogpreviousInitialized represents a Initialized event raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) FilterInitialized(opts *bind.FilterOpts) (*PolygonvalidiumetrogpreviousInitializedIterator, error) {

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogpreviousInitializedIterator{contract: _Polygonvalidiumetrogprevious.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogpreviousInitialized) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogpreviousInitialized)
				if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) ParseInitialized(log types.Log) (*PolygonvalidiumetrogpreviousInitialized, error) {
	event := new(PolygonvalidiumetrogpreviousInitialized)
	if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumetrogpreviousSequenceBatchesIterator is returned from FilterSequenceBatches and is used to iterate over the raw logs and unpacked data for SequenceBatches events raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousSequenceBatchesIterator struct {
	Event *PolygonvalidiumetrogpreviousSequenceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogpreviousSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogpreviousSequenceBatches)
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
		it.Event = new(PolygonvalidiumetrogpreviousSequenceBatches)
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
func (it *PolygonvalidiumetrogpreviousSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogpreviousSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogpreviousSequenceBatches represents a SequenceBatches event raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousSequenceBatches struct {
	NumBatch   uint64
	L1InfoRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSequenceBatches is a free log retrieval operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) FilterSequenceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonvalidiumetrogpreviousSequenceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.FilterLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogpreviousSequenceBatchesIterator{contract: _Polygonvalidiumetrogprevious.contract, event: "SequenceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceBatches is a free log subscription operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) WatchSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogpreviousSequenceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.WatchLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogpreviousSequenceBatches)
				if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
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
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) ParseSequenceBatches(log types.Log) (*PolygonvalidiumetrogpreviousSequenceBatches, error) {
	event := new(PolygonvalidiumetrogpreviousSequenceBatches)
	if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumetrogpreviousSequenceForceBatchesIterator is returned from FilterSequenceForceBatches and is used to iterate over the raw logs and unpacked data for SequenceForceBatches events raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousSequenceForceBatchesIterator struct {
	Event *PolygonvalidiumetrogpreviousSequenceForceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogpreviousSequenceForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogpreviousSequenceForceBatches)
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
		it.Event = new(PolygonvalidiumetrogpreviousSequenceForceBatches)
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
func (it *PolygonvalidiumetrogpreviousSequenceForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogpreviousSequenceForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogpreviousSequenceForceBatches represents a SequenceForceBatches event raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousSequenceForceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceForceBatches is a free log retrieval operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) FilterSequenceForceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonvalidiumetrogpreviousSequenceForceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.FilterLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogpreviousSequenceForceBatchesIterator{contract: _Polygonvalidiumetrogprevious.contract, event: "SequenceForceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceForceBatches is a free log subscription operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) WatchSequenceForceBatches(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogpreviousSequenceForceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.WatchLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogpreviousSequenceForceBatches)
				if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
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
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) ParseSequenceForceBatches(log types.Log) (*PolygonvalidiumetrogpreviousSequenceForceBatches, error) {
	event := new(PolygonvalidiumetrogpreviousSequenceForceBatches)
	if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumetrogpreviousSetDataAvailabilityProtocolIterator is returned from FilterSetDataAvailabilityProtocol and is used to iterate over the raw logs and unpacked data for SetDataAvailabilityProtocol events raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousSetDataAvailabilityProtocolIterator struct {
	Event *PolygonvalidiumetrogpreviousSetDataAvailabilityProtocol // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogpreviousSetDataAvailabilityProtocolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogpreviousSetDataAvailabilityProtocol)
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
		it.Event = new(PolygonvalidiumetrogpreviousSetDataAvailabilityProtocol)
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
func (it *PolygonvalidiumetrogpreviousSetDataAvailabilityProtocolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogpreviousSetDataAvailabilityProtocolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogpreviousSetDataAvailabilityProtocol represents a SetDataAvailabilityProtocol event raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousSetDataAvailabilityProtocol struct {
	NewDataAvailabilityProtocol common.Address
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterSetDataAvailabilityProtocol is a free log retrieval operation binding the contract event 0xd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a.
//
// Solidity: event SetDataAvailabilityProtocol(address newDataAvailabilityProtocol)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) FilterSetDataAvailabilityProtocol(opts *bind.FilterOpts) (*PolygonvalidiumetrogpreviousSetDataAvailabilityProtocolIterator, error) {

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.FilterLogs(opts, "SetDataAvailabilityProtocol")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogpreviousSetDataAvailabilityProtocolIterator{contract: _Polygonvalidiumetrogprevious.contract, event: "SetDataAvailabilityProtocol", logs: logs, sub: sub}, nil
}

// WatchSetDataAvailabilityProtocol is a free log subscription operation binding the contract event 0xd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a.
//
// Solidity: event SetDataAvailabilityProtocol(address newDataAvailabilityProtocol)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) WatchSetDataAvailabilityProtocol(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogpreviousSetDataAvailabilityProtocol) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.WatchLogs(opts, "SetDataAvailabilityProtocol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogpreviousSetDataAvailabilityProtocol)
				if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "SetDataAvailabilityProtocol", log); err != nil {
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
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) ParseSetDataAvailabilityProtocol(log types.Log) (*PolygonvalidiumetrogpreviousSetDataAvailabilityProtocol, error) {
	event := new(PolygonvalidiumetrogpreviousSetDataAvailabilityProtocol)
	if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "SetDataAvailabilityProtocol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumetrogpreviousSetForceBatchAddressIterator is returned from FilterSetForceBatchAddress and is used to iterate over the raw logs and unpacked data for SetForceBatchAddress events raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousSetForceBatchAddressIterator struct {
	Event *PolygonvalidiumetrogpreviousSetForceBatchAddress // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogpreviousSetForceBatchAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogpreviousSetForceBatchAddress)
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
		it.Event = new(PolygonvalidiumetrogpreviousSetForceBatchAddress)
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
func (it *PolygonvalidiumetrogpreviousSetForceBatchAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogpreviousSetForceBatchAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogpreviousSetForceBatchAddress represents a SetForceBatchAddress event raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousSetForceBatchAddress struct {
	NewForceBatchAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchAddress is a free log retrieval operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) FilterSetForceBatchAddress(opts *bind.FilterOpts) (*PolygonvalidiumetrogpreviousSetForceBatchAddressIterator, error) {

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.FilterLogs(opts, "SetForceBatchAddress")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogpreviousSetForceBatchAddressIterator{contract: _Polygonvalidiumetrogprevious.contract, event: "SetForceBatchAddress", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchAddress is a free log subscription operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) WatchSetForceBatchAddress(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogpreviousSetForceBatchAddress) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.WatchLogs(opts, "SetForceBatchAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogpreviousSetForceBatchAddress)
				if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "SetForceBatchAddress", log); err != nil {
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
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) ParseSetForceBatchAddress(log types.Log) (*PolygonvalidiumetrogpreviousSetForceBatchAddress, error) {
	event := new(PolygonvalidiumetrogpreviousSetForceBatchAddress)
	if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "SetForceBatchAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumetrogpreviousSetForceBatchTimeoutIterator is returned from FilterSetForceBatchTimeout and is used to iterate over the raw logs and unpacked data for SetForceBatchTimeout events raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousSetForceBatchTimeoutIterator struct {
	Event *PolygonvalidiumetrogpreviousSetForceBatchTimeout // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogpreviousSetForceBatchTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogpreviousSetForceBatchTimeout)
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
		it.Event = new(PolygonvalidiumetrogpreviousSetForceBatchTimeout)
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
func (it *PolygonvalidiumetrogpreviousSetForceBatchTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogpreviousSetForceBatchTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogpreviousSetForceBatchTimeout represents a SetForceBatchTimeout event raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousSetForceBatchTimeout struct {
	NewforceBatchTimeout uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchTimeout is a free log retrieval operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) FilterSetForceBatchTimeout(opts *bind.FilterOpts) (*PolygonvalidiumetrogpreviousSetForceBatchTimeoutIterator, error) {

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.FilterLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogpreviousSetForceBatchTimeoutIterator{contract: _Polygonvalidiumetrogprevious.contract, event: "SetForceBatchTimeout", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchTimeout is a free log subscription operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) WatchSetForceBatchTimeout(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogpreviousSetForceBatchTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.WatchLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogpreviousSetForceBatchTimeout)
				if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
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
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) ParseSetForceBatchTimeout(log types.Log) (*PolygonvalidiumetrogpreviousSetForceBatchTimeout, error) {
	event := new(PolygonvalidiumetrogpreviousSetForceBatchTimeout)
	if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumetrogpreviousSetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousSetTrustedSequencerIterator struct {
	Event *PolygonvalidiumetrogpreviousSetTrustedSequencer // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogpreviousSetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogpreviousSetTrustedSequencer)
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
		it.Event = new(PolygonvalidiumetrogpreviousSetTrustedSequencer)
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
func (it *PolygonvalidiumetrogpreviousSetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogpreviousSetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogpreviousSetTrustedSequencer represents a SetTrustedSequencer event raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousSetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*PolygonvalidiumetrogpreviousSetTrustedSequencerIterator, error) {

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogpreviousSetTrustedSequencerIterator{contract: _Polygonvalidiumetrogprevious.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogpreviousSetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogpreviousSetTrustedSequencer)
				if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
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
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) ParseSetTrustedSequencer(log types.Log) (*PolygonvalidiumetrogpreviousSetTrustedSequencer, error) {
	event := new(PolygonvalidiumetrogpreviousSetTrustedSequencer)
	if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumetrogpreviousSetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousSetTrustedSequencerURLIterator struct {
	Event *PolygonvalidiumetrogpreviousSetTrustedSequencerURL // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogpreviousSetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogpreviousSetTrustedSequencerURL)
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
		it.Event = new(PolygonvalidiumetrogpreviousSetTrustedSequencerURL)
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
func (it *PolygonvalidiumetrogpreviousSetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogpreviousSetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogpreviousSetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousSetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*PolygonvalidiumetrogpreviousSetTrustedSequencerURLIterator, error) {

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogpreviousSetTrustedSequencerURLIterator{contract: _Polygonvalidiumetrogprevious.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogpreviousSetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogpreviousSetTrustedSequencerURL)
				if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
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
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) ParseSetTrustedSequencerURL(log types.Log) (*PolygonvalidiumetrogpreviousSetTrustedSequencerURL, error) {
	event := new(PolygonvalidiumetrogpreviousSetTrustedSequencerURL)
	if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumetrogpreviousSwitchSequenceWithDataAvailabilityIterator is returned from FilterSwitchSequenceWithDataAvailability and is used to iterate over the raw logs and unpacked data for SwitchSequenceWithDataAvailability events raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousSwitchSequenceWithDataAvailabilityIterator struct {
	Event *PolygonvalidiumetrogpreviousSwitchSequenceWithDataAvailability // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogpreviousSwitchSequenceWithDataAvailabilityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogpreviousSwitchSequenceWithDataAvailability)
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
		it.Event = new(PolygonvalidiumetrogpreviousSwitchSequenceWithDataAvailability)
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
func (it *PolygonvalidiumetrogpreviousSwitchSequenceWithDataAvailabilityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogpreviousSwitchSequenceWithDataAvailabilityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogpreviousSwitchSequenceWithDataAvailability represents a SwitchSequenceWithDataAvailability event raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousSwitchSequenceWithDataAvailability struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSwitchSequenceWithDataAvailability is a free log retrieval operation binding the contract event 0xf32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41.
//
// Solidity: event SwitchSequenceWithDataAvailability()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) FilterSwitchSequenceWithDataAvailability(opts *bind.FilterOpts) (*PolygonvalidiumetrogpreviousSwitchSequenceWithDataAvailabilityIterator, error) {

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.FilterLogs(opts, "SwitchSequenceWithDataAvailability")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogpreviousSwitchSequenceWithDataAvailabilityIterator{contract: _Polygonvalidiumetrogprevious.contract, event: "SwitchSequenceWithDataAvailability", logs: logs, sub: sub}, nil
}

// WatchSwitchSequenceWithDataAvailability is a free log subscription operation binding the contract event 0xf32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41.
//
// Solidity: event SwitchSequenceWithDataAvailability()
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) WatchSwitchSequenceWithDataAvailability(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogpreviousSwitchSequenceWithDataAvailability) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.WatchLogs(opts, "SwitchSequenceWithDataAvailability")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogpreviousSwitchSequenceWithDataAvailability)
				if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "SwitchSequenceWithDataAvailability", log); err != nil {
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
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) ParseSwitchSequenceWithDataAvailability(log types.Log) (*PolygonvalidiumetrogpreviousSwitchSequenceWithDataAvailability, error) {
	event := new(PolygonvalidiumetrogpreviousSwitchSequenceWithDataAvailability)
	if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "SwitchSequenceWithDataAvailability", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumetrogpreviousTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousTransferAdminRoleIterator struct {
	Event *PolygonvalidiumetrogpreviousTransferAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogpreviousTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogpreviousTransferAdminRole)
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
		it.Event = new(PolygonvalidiumetrogpreviousTransferAdminRole)
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
func (it *PolygonvalidiumetrogpreviousTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogpreviousTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogpreviousTransferAdminRole represents a TransferAdminRole event raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*PolygonvalidiumetrogpreviousTransferAdminRoleIterator, error) {

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogpreviousTransferAdminRoleIterator{contract: _Polygonvalidiumetrogprevious.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogpreviousTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogpreviousTransferAdminRole)
				if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
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
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) ParseTransferAdminRole(log types.Log) (*PolygonvalidiumetrogpreviousTransferAdminRole, error) {
	event := new(PolygonvalidiumetrogpreviousTransferAdminRole)
	if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumetrogpreviousVerifyBatchesIterator is returned from FilterVerifyBatches and is used to iterate over the raw logs and unpacked data for VerifyBatches events raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousVerifyBatchesIterator struct {
	Event *PolygonvalidiumetrogpreviousVerifyBatches // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumetrogpreviousVerifyBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumetrogpreviousVerifyBatches)
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
		it.Event = new(PolygonvalidiumetrogpreviousVerifyBatches)
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
func (it *PolygonvalidiumetrogpreviousVerifyBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumetrogpreviousVerifyBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumetrogpreviousVerifyBatches represents a VerifyBatches event raised by the Polygonvalidiumetrogprevious contract.
type PolygonvalidiumetrogpreviousVerifyBatches struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatches is a free log retrieval operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) FilterVerifyBatches(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*PolygonvalidiumetrogpreviousVerifyBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.FilterLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumetrogpreviousVerifyBatchesIterator{contract: _Polygonvalidiumetrogprevious.contract, event: "VerifyBatches", logs: logs, sub: sub}, nil
}

// WatchVerifyBatches is a free log subscription operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) WatchVerifyBatches(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumetrogpreviousVerifyBatches, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonvalidiumetrogprevious.contract.WatchLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumetrogpreviousVerifyBatches)
				if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
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
func (_Polygonvalidiumetrogprevious *PolygonvalidiumetrogpreviousFilterer) ParseVerifyBatches(log types.Log) (*PolygonvalidiumetrogpreviousVerifyBatches, error) {
	event := new(PolygonvalidiumetrogpreviousVerifyBatches)
	if err := _Polygonvalidiumetrogprevious.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

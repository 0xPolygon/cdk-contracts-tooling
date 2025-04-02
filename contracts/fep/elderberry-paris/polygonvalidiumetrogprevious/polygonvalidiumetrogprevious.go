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
	Bin: "0x6101006040523480156200001257600080fd5b506040516200504138038062005041833981016040819052620000359162000071565b6001600160a01b0393841660a052918316608052821660c0521660e052620000d9565b6001600160a01b03811681146200006e57600080fd5b50565b600080600080608085870312156200008857600080fd5b8451620000958162000058565b6020860151909450620000a88162000058565b6040860151909350620000bb8162000058565b6060860151909250620000ce8162000058565b939692955090935050565b60805160a05160c05160e051614e55620001ec6000396000818161054d015281816110e7015281816111ae015281816111d0015281816113730152818161144c015281816115b90152818161180401528181611dea01528181612358015281816127af0152818161285d01528181612ee601528181612fae01528181613a6e01528181613ae701528181613b090152613bb40152600081816106ef01528181610c0c015281816119fd01528181611ad701528181612a2e01528181612b3601526135a80152600081816107ab01528181610c8e01528181611c5e015281816130f9015261362a0152600081816107dd015281816108cf0152818161113a0152818161127e01526130cd0152614e556000f3fe608060405234801561001057600080fd5b50600436106103145760003560e01c806371257022116101a7578063c7fffd4b116100ee578063e57a0b4c11610097578063ecef3f9911610071578063ecef3f9914610846578063f35dda4714610859578063f851a4401461086157600080fd5b8063e57a0b4c146107ff578063e7a7ed021461081f578063eaeb077b1461083357600080fd5b8063d02103ca116100c8578063d02103ca146107a6578063d7bc90ff146107cd578063e46761c4146107d857600080fd5b8063c7fffd4b1461076b578063c89e42df14610773578063cfa8ed471461078657600080fd5b80639f26f84011610150578063ada8f9191161012a578063ada8f91914610724578063b0afe15414610737578063c754c7ed1461074357600080fd5b80639f26f840146106d7578063a3c573eb146106ea578063a652f26c1461071157600080fd5b80638c3d7301116101815780638c3d7301146106a157806391cafe32146106a95780639e001877146106bc57600080fd5b8063712570221461063f5780637a5460c5146106525780637cd76b8b1461068e57600080fd5b80633cbc795b1161026b57806352bdeb6d116102145780636b8616ce116101ee5780636b8616ce146106035780636e05d2cd146106235780636ff512cc1461062c57600080fd5b806352bdeb6d146105b7578063542028d5146105f3578063676870d2146105fb57600080fd5b806349b7b8021161024557806349b7b802146105485780634c21fef31461056f5780634e487706146105a457600080fd5b80633cbc795b1461047a57806340b5de6c146104b7578063456052671461050f57600080fd5b806326782247116102cd5780632d72c248116102a75780632d72c2481461043457806332c2d153146104475780633c351e101461045a57600080fd5b806326782247146103ba5780632acdc2b6146103ff5780632c111c061461041457600080fd5b806305835f37116102fe57806305835f371461034f578063107bf28c1461039857806311e892d4146103a057600080fd5b8062d0295d146103195780630350896314610334575b600080fd5b610321610887565b6040519081526020015b60405180910390f35b61033c602081565b60405161ffff909116815260200161032b565b61038b6040518060400160405280600881526020017f80808401c9c3809400000000000000000000000000000000000000000000000081525081565b60405161032b9190613f74565b61038b610993565b6103a860f981565b60405160ff909116815260200161032b565b6001546103da9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161032b565b61041261040d366004613f9f565b610a21565b005b6008546103da9073ffffffffffffffffffffffffffffffffffffffff1681565b610412610442366004614037565b610b41565b6104126104553660046140ff565b61144a565b6009546103da9073ffffffffffffffffffffffffffffffffffffffff1681565b6009546104a29074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff909116815260200161032b565b6104de7fff0000000000000000000000000000000000000000000000000000000000000081565b6040517fff00000000000000000000000000000000000000000000000000000000000000909116815260200161032b565b60075461052f9068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff909116815260200161032b565b6103da7f000000000000000000000000000000000000000000000000000000000000000081565b600a546105949074010000000000000000000000000000000000000000900460ff1681565b604051901515815260200161032b565b6104126105b2366004614141565b611519565b61038b6040518060400160405280600281526020017f80b800000000000000000000000000000000000000000000000000000000000081525081565b61038b61172b565b61033c601f81565b610321610611366004614141565b60066020526000908152604090205481565b61032160055481565b61041261063a36600461415e565b611738565b61041261064d3660046142d0565b611802565b61038b6040518060400160405280600281526020017f80b900000000000000000000000000000000000000000000000000000000000081525081565b61041261069c36600461415e565b612026565b6104126120f0565b6104126106b736600461415e565b6121c3565b6103da73a40d5f56745a118d0906a34e69aec8c0db1cb8fa81565b6104126106e53660046143c2565b6122dc565b6103da7f000000000000000000000000000000000000000000000000000000000000000081565b61038b61071f366004614404565b61292d565b61041261073236600461415e565b612d12565b6103216405ca1ab1e081565b60075461052f90700100000000000000000000000000000000900467ffffffffffffffff1681565b6103a860e481565b610412610781366004614479565b612ddc565b6002546103da9073ffffffffffffffffffffffffffffffffffffffff1681565b6103da7f000000000000000000000000000000000000000000000000000000000000000081565b610321635ca1ab1e81565b6103da7f000000000000000000000000000000000000000000000000000000000000000081565b600a546103da9073ffffffffffffffffffffffffffffffffffffffff1681565b60075461052f9067ffffffffffffffff1681565b6104126108413660046144ae565b612e6f565b6104126108543660046144fa565b613341565b6103a8601b81565b6000546103da9062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b6040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152600090819073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015610916573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061093a9190614546565b6007549091506000906109659067ffffffffffffffff6801000000000000000082048116911661458e565b67ffffffffffffffff169050806000036109825760009250505090565b61098c81836145b6565b9250505090565b600480546109a0906145f1565b80601f01602080910402602001604051908101604052809291908181526020018280546109cc906145f1565b8015610a195780601f106109ee57610100808354040283529160200191610a19565b820191906000526020600020905b8154815290600101906020018083116109fc57829003601f168201915b505050505081565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314610a78576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600a5474010000000000000000000000000000000000000000900460ff16151581151503610ad2576040517f5f0e7abe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600a80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1674010000000000000000000000000000000000000000831515021790556040517ff32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f4190600090a150565b60025473ffffffffffffffffffffffffffffffffffffffff163314610b92576040517f11e7be1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b836000819003610bce576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8811115610c0a576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379e2cf976040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610c7257600080fd5b505af1158015610c86573d6000803e3d6000fd5b5050505060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635ca1e1656040518163ffffffff1660e01b8152600401602060405180830381865afa158015610cf7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d1b9190614546565b60075460055491925068010000000000000000900467ffffffffffffffff1690816000805b868110156110595760008c8c83818110610d5c57610d5c614644565b905060800201803603810190610d729190614673565b604081015190915067ffffffffffffffff1615610f665785610d93816146c1565b96505060008160000151826020015183604001518460600151604051602001610dfa9493929190938452602084019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166040830152604882015260680190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff8a16600090815260069093529120549091508114610e83576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b858260000151836020015184604001518f8660600151604051602001610f1d969594939291909586526020860194909452604085019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166060808501919091521b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166068830152607c820152609c0190565b604051602081830303815290604052805190602001209550600660008867ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206000905550611046565b8051604051610f82918591602001918252602082015260400190565b604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018152828252805160209182012084519184018990529183015260608083018a90524260c01b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808401528d901b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660888301526000609c830152935060bc016040516020818303038152906040528051906020012094505b5080611051816146e8565b915050610d40565b5060075467ffffffffffffffff90811690851611156110a4576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60058390558567ffffffffffffffff858116908416146111a25760006110ca848761458e565b90506110e067ffffffffffffffff821683614720565b91506111617f00000000000000000000000000000000000000000000000000000000000000008267ffffffffffffffff16611119610887565b6111239190614739565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691906133a5565b50600780547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8816021790555b8015611331576112a6337f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663477fa2706040518163ffffffff1660e01b8152600401602060405180830381865afa158015611239573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061125d9190614546565b6112679190614739565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016929190613479565b600a546040517f3b51be4b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690633b51be4b906113009085908d908d90600401614799565b60006040518083038186803b15801561131857600080fd5b505afa15801561132c573d6000803e3d6000fd5b505050505b6040517f9a908e7300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff88166004820152602481018590526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303816000875af11580156113d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113f591906147bc565b90508067ffffffffffffffff167f3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e7668860405161143391815260200190565b60405180910390a250505050505050505050505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1633146114b9576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff168367ffffffffffffffff167f9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f59668460405161150c91815260200190565b60405180910390a3505050565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314611570576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff821611156115b7576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166315064c966040518163ffffffff1660e01b8152600401602060405180830381865afa158015611622573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061164691906147d9565b6116a75760075467ffffffffffffffff7001000000000000000000000000000000009091048116908216106116a7576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600780547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527fa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b906020015b60405180910390a150565b600380546109a0906145f1565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff16331461178f576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc090602001611720565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314611871576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600054610100900460ff16158080156118915750600054600160ff909116105b806118ab5750303b1580156118ab575060005460ff166001145b61193c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561199a57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b606073ffffffffffffffffffffffffffffffffffffffff851615611bff576040517fc00f14ab00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063c00f14ab90602401600060405180830381865afa158015611a44573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052611a8a91908101906147f6565b6040517f318aee3d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff878116600483015291925060009182917f00000000000000000000000000000000000000000000000000000000000000009091169063318aee3d906024016040805180830381865afa158015611b1f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b439190614864565b915091508163ffffffff16600014611bbb576009805463ffffffff841674010000000000000000000000000000000000000000027fffffffffffffffff00000000000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff841617179055611bfc565b600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff89161790555b50505b600954600090611c4790889073ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900463ffffffff168561292d565b9050600081805190602001209050600042905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611cc7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ceb9190614546565b90506000808483858f611cff600143614720565b60408051602081019790975286019490945260608086019390935260c09190911b7fffffffffffffffff000000000000000000000000000000000000000000000000166080850152901b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016608883015240609c82015260bc01604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001815290829052805160209091012060058190557f9a908e73000000000000000000000000000000000000000000000000000000008252600160048301526024820181905291507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303816000875af1158015611e48573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e6c91906147bc565b508c600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508b600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508860039081611efe91906148e4565b506004611f0b89826148e4565b508c600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062069780600760106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507f060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f85838e604051611fac939291906149fe565b60405180910390a1505050505050801561201d57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff16331461207d576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600a80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a90602001611720565b60015473ffffffffffffffffffffffffffffffffffffffff163314612141576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600154600080547fffffffffffffffffffff0000000000000000000000000000000000000000ffff1673ffffffffffffffffffffffffffffffffffffffff9092166201000081029290921790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e9060200160405180910390a1565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff16331461221a576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60085473ffffffffffffffffffffffffffffffffffffffff16612269576040517fc89374d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb90602001611720565b60085473ffffffffffffffffffffffffffffffffffffffff16801580159061231a575073ffffffffffffffffffffffffffffffffffffffff81163314155b15612351576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b4262093a807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166330c27dde6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156123c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123e591906147bc565b6123ef9190614a3d565b67ffffffffffffffff161115612431576040517f3d49ed4c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81600081900361246d576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e88111156124a9576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075467ffffffffffffffff808216916124d191849168010000000000000000900416614a5e565b1115612509576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6007546005546801000000000000000090910467ffffffffffffffff169060005b838110156127a957600087878381811061254657612546614644565b90506020028101906125589190614a71565b61256190614aaf565b90508361256d816146c1565b825180516020918201208185015160408087015160608801519151959a509295506000946125da948794929101938452602084019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166040830152604882015260680190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff8916600090815260069093529120549091508114612663576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8616600090815260066020526040812055612688600188614720565b84036126f75742600760109054906101000a900467ffffffffffffffff1684604001516126b59190614a3d565b67ffffffffffffffff1611156126f7576040517fc44a082100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208381015160408086015160608088015183519586018b90529285018790528481019390935260c01b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808401523390911b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088830152609c82015260bc0160405160208183030381529060405280519060200120945050505080806127a1906146e8565b91505061252a565b506127d77f000000000000000000000000000000000000000000000000000000000000000084611119610887565b60058190556007805467ffffffffffffffff841668010000000000000000027fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff9091161790556040517f9a908e7300000000000000000000000000000000000000000000000000000000815260009073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690639a908e73906128a9908790869060040167ffffffffffffffff929092168252602082015260400190565b6020604051808303816000875af11580156128c8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128ec91906147bc565b60405190915067ffffffffffffffff8216907f648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a490600090a250505050505050565b6060600085858573a40d5f56745a118d0906a34e69aec8c0db1cb8fa60008760405160240161296196959493929190614b1d565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff811bff7000000000000000000000000000000000000000000000000000000001790528351909150606090600003612ab25760f9601f83516129f69190614b80565b6040518060400160405280600881526020017f80808401c9c380940000000000000000000000000000000000000000000000008152507f00000000000000000000000000000000000000000000000000000000000000006040518060400160405280600281526020017f80b800000000000000000000000000000000000000000000000000000000000081525060e487604051602001612a9c9796959493929190614b9b565b6040516020818303038152906040529050612bb6565b815161ffff1015612aef576040517f248b8f8200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815160f9612afe602083614b80565b6040518060400160405280600881526020017f80808401c9c380940000000000000000000000000000000000000000000000008152507f00000000000000000000000000000000000000000000000000000000000000006040518060400160405280600281526020017f80b90000000000000000000000000000000000000000000000000000000000008152508588604051602001612ba39796959493929190614c7e565b6040516020818303038152906040529150505b805160208083019190912060408051600080825293810180835292909252601b908201526405ca1ab1e06060820152635ca1ab1e608082015260019060a0016020604051602081039080840390855afa158015612c17573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116612c8f576040517fcd16196600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604051600090612cd59084906405ca1ab1e090635ca1ab1e90601b907fff0000000000000000000000000000000000000000000000000000000000000090602001614d61565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529450505050505b949350505050565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314612d69576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001611720565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314612e33576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6003612e3f82826148e4565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20816040516117209190613f74565b60085473ffffffffffffffffffffffffffffffffffffffff168015801590612ead575073ffffffffffffffffffffffffffffffffffffffff81163314155b15612ee4576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166315064c966040518163ffffffff1660e01b8152600401602060405180830381865afa158015612f4f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f7391906147d9565b15612faa576040517f39258d1800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663604691696040518163ffffffff1660e01b8152600401602060405180830381865afa158015613017573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061303b9190614546565b905082811115613077576040517f2354600f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113888411156130b3576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6130f573ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016333084613479565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613162573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131869190614546565b6007805491925067ffffffffffffffff9091169060006131a5836146c1565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505085856040516131dc929190614dbd565b60405190819003902081426131f2600143614720565b60408051602081019590955284019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166060830152406068820152608801604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012060075467ffffffffffffffff16600090815260069093529120553233036132ea57600754604080518381523360208201526060818301819052600090820152905167ffffffffffffffff909216917ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319181900360800190a2613339565b60075460405167ffffffffffffffff909116907ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319061333090849033908b908b90614dcd565b60405180910390a25b505050505050565b600a5474010000000000000000000000000000000000000000900460ff16613395576040517f821935b400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6133a08383836134dd565b505050565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526133a09084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152613c89565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526134d79085907f23b872dd00000000000000000000000000000000000000000000000000000000906084016133f7565b50505050565b60025473ffffffffffffffffffffffffffffffffffffffff16331461352e576040517f11e7be1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81600081900361356a576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e88111156135a6576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379e2cf976040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561360e57600080fd5b505af1158015613622573d6000803e3d6000fd5b5050505060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635ca1e1656040518163ffffffff1660e01b8152600401602060405180830381865afa158015613693573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136b79190614546565b60075460055491925042916801000000000000000090910467ffffffffffffffff16908160005b868110156139e05760008a8a838181106136fa576136fa614644565b905060200281019061370c9190614a71565b61371590614aaf565b8051805160209091012060408201519192509067ffffffffffffffff16156138fa5785613741816146c1565b9650506000818360200151846040015185606001516040516020016137a49493929190938452602084019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166040830152604882015260680190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff8a1660009081526006909352912054909150811461382d576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208381015160408086015160608088015183519586018c90529285018790528481019390935260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166080840152908d901b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088830152609c82015260bc01604051602081830303815290604052805190602001209550600660008867ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060009055506139cb565b8151516201d4c01015613939576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080516020810187905290810182905260608082018a905260c089901b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808301528b901b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660888201526000609c82015260bc016040516020818303038152906040528051906020012094505b505080806139d8906146e8565b9150506136de565b5060075467ffffffffffffffff9081169084161115613a2b576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60058290558567ffffffffffffffff84811690831614613ae1576000613a51838661458e565b9050613a6767ffffffffffffffff821683614720565b9150613aa07f00000000000000000000000000000000000000000000000000000000000000008267ffffffffffffffff16611119610887565b50600780547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8716021790555b613b72337f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663477fa2706040518163ffffffff1660e01b8152600401602060405180830381865afa158015611239573d6000803e3d6000fd5b6040517f9a908e7300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff88166004820152602481018490526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303816000875af1158015613c12573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613c3691906147bc565b90508067ffffffffffffffff167f3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e76688604051613c7491815260200190565b60405180910390a25050505050505050505050565b6000613ceb826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16613d959092919063ffffffff16565b8051909150156133a05780806020019051810190613d0991906147d9565b6133a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401611933565b6060612d0a8484600085856000808673ffffffffffffffffffffffffffffffffffffffff168587604051613dc99190614e0d565b60006040518083038185875af1925050503d8060008114613e06576040519150601f19603f3d011682016040523d82523d6000602084013e613e0b565b606091505b5091509150613e1c87838387613e27565b979650505050505050565b60608315613ebd578251600003613eb65773ffffffffffffffffffffffffffffffffffffffff85163b613eb6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401611933565b5081612d0a565b612d0a8383815115613ed25781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119339190613f74565b60005b83811015613f21578181015183820152602001613f09565b50506000910152565b60008151808452613f42816020860160208601613f06565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000613f876020830184613f2a565b9392505050565b8015158114613f9c57600080fd5b50565b600060208284031215613fb157600080fd5b8135613f8781613f8e565b73ffffffffffffffffffffffffffffffffffffffff81168114613f9c57600080fd5b8035613fe981613fbc565b919050565b60008083601f84011261400057600080fd5b50813567ffffffffffffffff81111561401857600080fd5b60208301915083602082850101111561403057600080fd5b9250929050565b60008060008060006060868803121561404f57600080fd5b853567ffffffffffffffff8082111561406757600080fd5b818801915088601f83011261407b57600080fd5b81358181111561408a57600080fd5b8960208260071b850101111561409f57600080fd5b602083019750809650506140b560208901613fde565b945060408801359150808211156140cb57600080fd5b506140d888828901613fee565b969995985093965092949392505050565b67ffffffffffffffff81168114613f9c57600080fd5b60008060006060848603121561411457600080fd5b833561411f816140e9565b925060208401359150604084013561413681613fbc565b809150509250925092565b60006020828403121561415357600080fd5b8135613f87816140e9565b60006020828403121561417057600080fd5b8135613f8781613fbc565b63ffffffff81168114613f9c57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516080810167ffffffffffffffff811182821017156141df576141df61418d565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561422c5761422c61418d565b604052919050565b600067ffffffffffffffff82111561424e5761424e61418d565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f83011261428b57600080fd5b813561429e61429982614234565b6141e5565b8181528460208386010111156142b357600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060008060c087890312156142e957600080fd5b86356142f481613fbc565b9550602087013561430481613fbc565b945060408701356143148161417b565b9350606087013561432481613fbc565b9250608087013567ffffffffffffffff8082111561434157600080fd5b61434d8a838b0161427a565b935060a089013591508082111561436357600080fd5b5061437089828a0161427a565b9150509295509295509295565b60008083601f84011261438f57600080fd5b50813567ffffffffffffffff8111156143a757600080fd5b6020830191508360208260051b850101111561403057600080fd5b600080602083850312156143d557600080fd5b823567ffffffffffffffff8111156143ec57600080fd5b6143f88582860161437d565b90969095509350505050565b6000806000806080858703121561441a57600080fd5b84356144258161417b565b9350602085013561443581613fbc565b925060408501356144458161417b565b9150606085013567ffffffffffffffff81111561446157600080fd5b61446d8782880161427a565b91505092959194509250565b60006020828403121561448b57600080fd5b813567ffffffffffffffff8111156144a257600080fd5b612d0a8482850161427a565b6000806000604084860312156144c357600080fd5b833567ffffffffffffffff8111156144da57600080fd5b6144e686828701613fee565b909790965060209590950135949350505050565b60008060006040848603121561450f57600080fd5b833567ffffffffffffffff81111561452657600080fd5b6145328682870161437d565b909450925050602084013561413681613fbc565b60006020828403121561455857600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b67ffffffffffffffff8281168282160390808211156145af576145af61455f565b5092915050565b6000826145ec577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b600181811c9082168061460557607f821691505b60208210810361463e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006080828403121561468557600080fd5b61468d6141bc565b823581526020830135602082015260408301356146a9816140e9565b60408201526060928301359281019290925250919050565b600067ffffffffffffffff8083168181036146de576146de61455f565b6001019392505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036147195761471961455f565b5060010190565b818103818111156147335761473361455f565b92915050565b80820281158282048414176147335761473361455f565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b8381526040602082015260006147b3604083018486614750565b95945050505050565b6000602082840312156147ce57600080fd5b8151613f87816140e9565b6000602082840312156147eb57600080fd5b8151613f8781613f8e565b60006020828403121561480857600080fd5b815167ffffffffffffffff81111561481f57600080fd5b8201601f8101841361483057600080fd5b805161483e61429982614234565b81815285602083850101111561485357600080fd5b6147b3826020830160208601613f06565b6000806040838503121561487757600080fd5b82516148828161417b565b602084015190925061489381613fbc565b809150509250929050565b601f8211156133a057600081815260208120601f850160051c810160208610156148c55750805b601f850160051c820191505b81811015613339578281556001016148d1565b815167ffffffffffffffff8111156148fe576148fe61418d565b6149128161490c84546145f1565b8461489e565b602080601f831160018114614965576000841561492f5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555613339565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b828110156149b257888601518255948401946001909101908401614993565b50858210156149ee57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b606081526000614a116060830186613f2a565b905083602083015273ffffffffffffffffffffffffffffffffffffffff83166040830152949350505050565b67ffffffffffffffff8181168382160190808211156145af576145af61455f565b808201808211156147335761473361455f565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81833603018112614aa557600080fd5b9190910192915050565b600060808236031215614ac157600080fd5b614ac96141bc565b823567ffffffffffffffff811115614ae057600080fd5b614aec3682860161427a565b825250602083013560208201526040830135614b07816140e9565b6040820152606092830135928101929092525090565b600063ffffffff808916835273ffffffffffffffffffffffffffffffffffffffff8089166020850152818816604085015280871660608501528086166080850152505060c060a0830152614b7460c0830184613f2a565b98975050505050505050565b61ffff8181168382160190808211156145af576145af61455f565b60007fff00000000000000000000000000000000000000000000000000000000000000808a60f81b1683527fffff0000000000000000000000000000000000000000000000000000000000008960f01b1660018401528751614c04816003860160208c01613f06565b80840190507fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008860601b1660038201528651614c47816017840160208b01613f06565b808201915050818660f81b16601782015284519150614c6d826018830160208801613f06565b016018019998505050505050505050565b7fff000000000000000000000000000000000000000000000000000000000000008860f81b16815260007fffff000000000000000000000000000000000000000000000000000000000000808960f01b1660018401528751614ce7816003860160208c01613f06565b80840190507fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008860601b1660038201528651614d2a816017840160208b01613f06565b808201915050818660f01b16601782015284519150614d50826019830160208801613f06565b016019019998505050505050505050565b60008651614d73818460208b01613f06565b9190910194855250602084019290925260f81b7fff000000000000000000000000000000000000000000000000000000000000009081166040840152166041820152604201919050565b8183823760009101908152919050565b84815273ffffffffffffffffffffffffffffffffffffffff84166020820152606060408201526000614e03606083018486614750565b9695505050505050565b60008251614aa5818460208701613f0656fea264697066735822122055e4c528e076e3ab4e58001a686a1b351e9ffb9ffd3bb8a50c3b72f4665206e664736f6c63430008140033",
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

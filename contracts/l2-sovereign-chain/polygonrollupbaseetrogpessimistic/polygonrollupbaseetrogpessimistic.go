// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonrollupbaseetrogpessimistic

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

// PolygonrollupbaseetrogpessimisticMetaData contains all meta data concerning the Polygonrollupbaseetrogpessimistic contract.
var PolygonrollupbaseetrogpessimisticMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalAccInputHashDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"forceBatchNum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"}],\"name\":\"ForceBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"}],\"name\":\"InitialSequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"targetBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accInputHashToRollback\",\"type\":\"bytes32\"}],\"name\":\"RollbackBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"l1InfoRoot\",\"type\":\"bytes32\"}],\"name\":\"SequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newForceBatchAddress\",\"type\":\"address\"}],\"name\":\"SetForceBatchAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"SetForceBatchTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatches\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GLOBAL_EXIT_ROOT_MANAGER_L2\",\"outputs\":[{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_LIST_LEN_LEN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_DATA_LEN_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_EFFECTIVE_PERCENTAGE\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_R\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_S\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_V\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TIMESTAMP_RANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculatePolPerForceBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"polAmount\",\"type\":\"uint256\"}],\"name\":\"forceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"}],\"name\":\"generateInitializeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"onVerifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"targetBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"accInputHashToRollback\",\"type\":\"bytes32\"}],\"name\":\"rollbackBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManagerPessimistic\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrogPessimistic.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"uint32\",\"name\":\"l1InfoTreeLeafCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"maxSequenceTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"expectedFinalAccInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"}],\"name\":\"sequenceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrogPessimistic.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"}],\"name\":\"sequenceForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newForceBatchAddress\",\"type\":\"address\"}],\"name\":\"setForceBatchAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"setForceBatchTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PolygonrollupbaseetrogpessimisticABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonrollupbaseetrogpessimisticMetaData.ABI instead.
var PolygonrollupbaseetrogpessimisticABI = PolygonrollupbaseetrogpessimisticMetaData.ABI

// Polygonrollupbaseetrogpessimistic is an auto generated Go binding around an Ethereum contract.
type Polygonrollupbaseetrogpessimistic struct {
	PolygonrollupbaseetrogpessimisticCaller     // Read-only binding to the contract
	PolygonrollupbaseetrogpessimisticTransactor // Write-only binding to the contract
	PolygonrollupbaseetrogpessimisticFilterer   // Log filterer for contract events
}

// PolygonrollupbaseetrogpessimisticCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonrollupbaseetrogpessimisticCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonrollupbaseetrogpessimisticTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonrollupbaseetrogpessimisticTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonrollupbaseetrogpessimisticFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonrollupbaseetrogpessimisticFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonrollupbaseetrogpessimisticSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonrollupbaseetrogpessimisticSession struct {
	Contract     *Polygonrollupbaseetrogpessimistic // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                      // Call options to use throughout this session
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// PolygonrollupbaseetrogpessimisticCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonrollupbaseetrogpessimisticCallerSession struct {
	Contract *PolygonrollupbaseetrogpessimisticCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                            // Call options to use throughout this session
}

// PolygonrollupbaseetrogpessimisticTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonrollupbaseetrogpessimisticTransactorSession struct {
	Contract     *PolygonrollupbaseetrogpessimisticTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                            // Transaction auth options to use throughout this session
}

// PolygonrollupbaseetrogpessimisticRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonrollupbaseetrogpessimisticRaw struct {
	Contract *Polygonrollupbaseetrogpessimistic // Generic contract binding to access the raw methods on
}

// PolygonrollupbaseetrogpessimisticCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonrollupbaseetrogpessimisticCallerRaw struct {
	Contract *PolygonrollupbaseetrogpessimisticCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonrollupbaseetrogpessimisticTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonrollupbaseetrogpessimisticTransactorRaw struct {
	Contract *PolygonrollupbaseetrogpessimisticTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonrollupbaseetrogpessimistic creates a new instance of Polygonrollupbaseetrogpessimistic, bound to a specific deployed contract.
func NewPolygonrollupbaseetrogpessimistic(address common.Address, backend bind.ContractBackend) (*Polygonrollupbaseetrogpessimistic, error) {
	contract, err := bindPolygonrollupbaseetrogpessimistic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonrollupbaseetrogpessimistic{PolygonrollupbaseetrogpessimisticCaller: PolygonrollupbaseetrogpessimisticCaller{contract: contract}, PolygonrollupbaseetrogpessimisticTransactor: PolygonrollupbaseetrogpessimisticTransactor{contract: contract}, PolygonrollupbaseetrogpessimisticFilterer: PolygonrollupbaseetrogpessimisticFilterer{contract: contract}}, nil
}

// NewPolygonrollupbaseetrogpessimisticCaller creates a new read-only instance of Polygonrollupbaseetrogpessimistic, bound to a specific deployed contract.
func NewPolygonrollupbaseetrogpessimisticCaller(address common.Address, caller bind.ContractCaller) (*PolygonrollupbaseetrogpessimisticCaller, error) {
	contract, err := bindPolygonrollupbaseetrogpessimistic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpessimisticCaller{contract: contract}, nil
}

// NewPolygonrollupbaseetrogpessimisticTransactor creates a new write-only instance of Polygonrollupbaseetrogpessimistic, bound to a specific deployed contract.
func NewPolygonrollupbaseetrogpessimisticTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonrollupbaseetrogpessimisticTransactor, error) {
	contract, err := bindPolygonrollupbaseetrogpessimistic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpessimisticTransactor{contract: contract}, nil
}

// NewPolygonrollupbaseetrogpessimisticFilterer creates a new log filterer instance of Polygonrollupbaseetrogpessimistic, bound to a specific deployed contract.
func NewPolygonrollupbaseetrogpessimisticFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonrollupbaseetrogpessimisticFilterer, error) {
	contract, err := bindPolygonrollupbaseetrogpessimistic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpessimisticFilterer{contract: contract}, nil
}

// bindPolygonrollupbaseetrogpessimistic binds a generic wrapper to an already deployed contract.
func bindPolygonrollupbaseetrogpessimistic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonrollupbaseetrogpessimisticMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonrollupbaseetrogpessimistic.Contract.PolygonrollupbaseetrogpessimisticCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.PolygonrollupbaseetrogpessimisticTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.PolygonrollupbaseetrogpessimisticTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonrollupbaseetrogpessimistic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.contract.Transact(opts, method, params...)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) GLOBALEXITROOTMANAGERL2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "GLOBAL_EXIT_ROOT_MANAGER_L2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.GLOBALEXITROOTMANAGERL2(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.GLOBALEXITROOTMANAGERL2(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) INITIALIZETXBRIDGELISTLENLEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_LIST_LEN_LEN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.INITIALIZETXBRIDGELISTLENLEN(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.INITIALIZETXBRIDGELISTLENLEN(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) INITIALIZETXBRIDGEPARAMS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.INITIALIZETXBRIDGEPARAMS(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.INITIALIZETXBRIDGEPARAMS(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) INITIALIZETXCONSTANTBYTES(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.INITIALIZETXCONSTANTBYTES(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.INITIALIZETXCONSTANTBYTES(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) INITIALIZETXCONSTANTBYTESEMPTYMETADATA(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) INITIALIZETXDATALENEMPTYMETADATA(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "INITIALIZE_TX_DATA_LEN_EMPTY_METADATA")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) INITIALIZETXEFFECTIVEPERCENTAGE(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "INITIALIZE_TX_EFFECTIVE_PERCENTAGE")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) SIGNATUREINITIALIZETXR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_R")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.SIGNATUREINITIALIZETXR(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.SIGNATUREINITIALIZETXR(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) SIGNATUREINITIALIZETXS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_S")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.SIGNATUREINITIALIZETXS(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.SIGNATUREINITIALIZETXS(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) SIGNATUREINITIALIZETXV(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_V")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.SIGNATUREINITIALIZETXV(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.SIGNATUREINITIALIZETXV(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) TIMESTAMPRANGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "TIMESTAMP_RANGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) TIMESTAMPRANGE() (*big.Int, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.TIMESTAMPRANGE(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) TIMESTAMPRANGE() (*big.Int, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.TIMESTAMPRANGE(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) Admin() (common.Address, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.Admin(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) Admin() (common.Address, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.Admin(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) BridgeAddress() (common.Address, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.BridgeAddress(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) BridgeAddress() (common.Address, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.BridgeAddress(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) CalculatePolPerForceBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "calculatePolPerForceBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) CalculatePolPerForceBatch() (*big.Int, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.CalculatePolPerForceBatch(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) CalculatePolPerForceBatch() (*big.Int, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.CalculatePolPerForceBatch(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) ForceBatchAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "forceBatchAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) ForceBatchAddress() (common.Address, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.ForceBatchAddress(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) ForceBatchAddress() (common.Address, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.ForceBatchAddress(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.ForceBatchTimeout(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.ForceBatchTimeout(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.ForcedBatches(&_Polygonrollupbaseetrogpessimistic.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.ForcedBatches(&_Polygonrollupbaseetrogpessimistic.CallOpts, arg0)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) GasTokenAddress() (common.Address, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.GasTokenAddress(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) GasTokenAddress() (common.Address, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.GasTokenAddress(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) GasTokenNetwork() (uint32, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.GasTokenNetwork(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) GasTokenNetwork() (uint32, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.GasTokenNetwork(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) GenerateInitializeTransaction(opts *bind.CallOpts, networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "generateInitializeTransaction", networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) GenerateInitializeTransaction(networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.GenerateInitializeTransaction(&_Polygonrollupbaseetrogpessimistic.CallOpts, networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) GenerateInitializeTransaction(networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.GenerateInitializeTransaction(&_Polygonrollupbaseetrogpessimistic.CallOpts, networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.GlobalExitRootManager(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.GlobalExitRootManager(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) LastAccInputHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "lastAccInputHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) LastAccInputHash() ([32]byte, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.LastAccInputHash(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) LastAccInputHash() ([32]byte, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.LastAccInputHash(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) LastForceBatch() (uint64, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.LastForceBatch(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) LastForceBatch() (uint64, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.LastForceBatch(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.LastForceBatchSequenced(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.LastForceBatchSequenced(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) NetworkName() (string, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.NetworkName(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) NetworkName() (string, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.NetworkName(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) PendingAdmin() (common.Address, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.PendingAdmin(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) PendingAdmin() (common.Address, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.PendingAdmin(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) Pol() (common.Address, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.Pol(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) Pol() (common.Address, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.Pol(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) RollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "rollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) RollupManager() (common.Address, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.RollupManager(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) RollupManager() (common.Address, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.RollupManager(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) TrustedSequencer() (common.Address, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.TrustedSequencer(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) TrustedSequencer() (common.Address, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.TrustedSequencer(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCaller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonrollupbaseetrogpessimistic.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) TrustedSequencerURL() (string, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.TrustedSequencerURL(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticCallerSession) TrustedSequencerURL() (string, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.TrustedSequencerURL(&_Polygonrollupbaseetrogpessimistic.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.AcceptAdminRole(&_Polygonrollupbaseetrogpessimistic.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.AcceptAdminRole(&_Polygonrollupbaseetrogpessimistic.TransactOpts)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticTransactor) ForceBatch(opts *bind.TransactOpts, transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.contract.Transact(opts, "forceBatch", transactions, polAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) ForceBatch(transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.ForceBatch(&_Polygonrollupbaseetrogpessimistic.TransactOpts, transactions, polAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticTransactorSession) ForceBatch(transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.ForceBatch(&_Polygonrollupbaseetrogpessimistic.TransactOpts, transactions, polAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.contract.Transact(opts, "initialize", _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.Initialize(&_Polygonrollupbaseetrogpessimistic.TransactOpts, _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticTransactorSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.Initialize(&_Polygonrollupbaseetrogpessimistic.TransactOpts, _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticTransactor) OnVerifyBatches(opts *bind.TransactOpts, lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.contract.Transact(opts, "onVerifyBatches", lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.OnVerifyBatches(&_Polygonrollupbaseetrogpessimistic.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticTransactorSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.OnVerifyBatches(&_Polygonrollupbaseetrogpessimistic.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x669adece.
//
// Solidity: function rollbackBatches(uint64 targetBatch, bytes32 accInputHashToRollback) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticTransactor) RollbackBatches(opts *bind.TransactOpts, targetBatch uint64, accInputHashToRollback [32]byte) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.contract.Transact(opts, "rollbackBatches", targetBatch, accInputHashToRollback)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x669adece.
//
// Solidity: function rollbackBatches(uint64 targetBatch, bytes32 accInputHashToRollback) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) RollbackBatches(targetBatch uint64, accInputHashToRollback [32]byte) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.RollbackBatches(&_Polygonrollupbaseetrogpessimistic.TransactOpts, targetBatch, accInputHashToRollback)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x669adece.
//
// Solidity: function rollbackBatches(uint64 targetBatch, bytes32 accInputHashToRollback) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticTransactorSession) RollbackBatches(targetBatch uint64, accInputHashToRollback [32]byte) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.RollbackBatches(&_Polygonrollupbaseetrogpessimistic.TransactOpts, targetBatch, accInputHashToRollback)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xb910e0f9.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint32 l1InfoTreeLeafCount, uint64 maxSequenceTimestamp, bytes32 expectedFinalAccInputHash, address l2Coinbase) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticTransactor) SequenceBatches(opts *bind.TransactOpts, batches []PolygonRollupBaseEtrogPessimisticBatchData, l1InfoTreeLeafCount uint32, maxSequenceTimestamp uint64, expectedFinalAccInputHash [32]byte, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.contract.Transact(opts, "sequenceBatches", batches, l1InfoTreeLeafCount, maxSequenceTimestamp, expectedFinalAccInputHash, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xb910e0f9.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint32 l1InfoTreeLeafCount, uint64 maxSequenceTimestamp, bytes32 expectedFinalAccInputHash, address l2Coinbase) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) SequenceBatches(batches []PolygonRollupBaseEtrogPessimisticBatchData, l1InfoTreeLeafCount uint32, maxSequenceTimestamp uint64, expectedFinalAccInputHash [32]byte, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.SequenceBatches(&_Polygonrollupbaseetrogpessimistic.TransactOpts, batches, l1InfoTreeLeafCount, maxSequenceTimestamp, expectedFinalAccInputHash, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xb910e0f9.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint32 l1InfoTreeLeafCount, uint64 maxSequenceTimestamp, bytes32 expectedFinalAccInputHash, address l2Coinbase) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticTransactorSession) SequenceBatches(batches []PolygonRollupBaseEtrogPessimisticBatchData, l1InfoTreeLeafCount uint32, maxSequenceTimestamp uint64, expectedFinalAccInputHash [32]byte, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.SequenceBatches(&_Polygonrollupbaseetrogpessimistic.TransactOpts, batches, l1InfoTreeLeafCount, maxSequenceTimestamp, expectedFinalAccInputHash, l2Coinbase)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticTransactor) SequenceForceBatches(opts *bind.TransactOpts, batches []PolygonRollupBaseEtrogPessimisticBatchData) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.contract.Transact(opts, "sequenceForceBatches", batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) SequenceForceBatches(batches []PolygonRollupBaseEtrogPessimisticBatchData) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.SequenceForceBatches(&_Polygonrollupbaseetrogpessimistic.TransactOpts, batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticTransactorSession) SequenceForceBatches(batches []PolygonRollupBaseEtrogPessimisticBatchData) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.SequenceForceBatches(&_Polygonrollupbaseetrogpessimistic.TransactOpts, batches)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticTransactor) SetForceBatchAddress(opts *bind.TransactOpts, newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.contract.Transact(opts, "setForceBatchAddress", newForceBatchAddress)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) SetForceBatchAddress(newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.SetForceBatchAddress(&_Polygonrollupbaseetrogpessimistic.TransactOpts, newForceBatchAddress)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticTransactorSession) SetForceBatchAddress(newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.SetForceBatchAddress(&_Polygonrollupbaseetrogpessimistic.TransactOpts, newForceBatchAddress)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticTransactor) SetForceBatchTimeout(opts *bind.TransactOpts, newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.contract.Transact(opts, "setForceBatchTimeout", newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.SetForceBatchTimeout(&_Polygonrollupbaseetrogpessimistic.TransactOpts, newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticTransactorSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.SetForceBatchTimeout(&_Polygonrollupbaseetrogpessimistic.TransactOpts, newforceBatchTimeout)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticTransactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.SetTrustedSequencer(&_Polygonrollupbaseetrogpessimistic.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticTransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.SetTrustedSequencer(&_Polygonrollupbaseetrogpessimistic.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticTransactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.SetTrustedSequencerURL(&_Polygonrollupbaseetrogpessimistic.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticTransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.SetTrustedSequencerURL(&_Polygonrollupbaseetrogpessimistic.TransactOpts, newTrustedSequencerURL)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.TransferAdminRole(&_Polygonrollupbaseetrogpessimistic.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonrollupbaseetrogpessimistic.Contract.TransferAdminRole(&_Polygonrollupbaseetrogpessimistic.TransactOpts, newPendingAdmin)
}

// PolygonrollupbaseetrogpessimisticAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Polygonrollupbaseetrogpessimistic contract.
type PolygonrollupbaseetrogpessimisticAcceptAdminRoleIterator struct {
	Event *PolygonrollupbaseetrogpessimisticAcceptAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonrollupbaseetrogpessimisticAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupbaseetrogpessimisticAcceptAdminRole)
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
		it.Event = new(PolygonrollupbaseetrogpessimisticAcceptAdminRole)
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
func (it *PolygonrollupbaseetrogpessimisticAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupbaseetrogpessimisticAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupbaseetrogpessimisticAcceptAdminRole represents a AcceptAdminRole event raised by the Polygonrollupbaseetrogpessimistic contract.
type PolygonrollupbaseetrogpessimisticAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*PolygonrollupbaseetrogpessimisticAcceptAdminRoleIterator, error) {

	logs, sub, err := _Polygonrollupbaseetrogpessimistic.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpessimisticAcceptAdminRoleIterator{contract: _Polygonrollupbaseetrogpessimistic.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonrollupbaseetrogpessimisticAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupbaseetrogpessimistic.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupbaseetrogpessimisticAcceptAdminRole)
				if err := _Polygonrollupbaseetrogpessimistic.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
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
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) ParseAcceptAdminRole(log types.Log) (*PolygonrollupbaseetrogpessimisticAcceptAdminRole, error) {
	event := new(PolygonrollupbaseetrogpessimisticAcceptAdminRole)
	if err := _Polygonrollupbaseetrogpessimistic.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupbaseetrogpessimisticForceBatchIterator is returned from FilterForceBatch and is used to iterate over the raw logs and unpacked data for ForceBatch events raised by the Polygonrollupbaseetrogpessimistic contract.
type PolygonrollupbaseetrogpessimisticForceBatchIterator struct {
	Event *PolygonrollupbaseetrogpessimisticForceBatch // Event containing the contract specifics and raw log

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
func (it *PolygonrollupbaseetrogpessimisticForceBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupbaseetrogpessimisticForceBatch)
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
		it.Event = new(PolygonrollupbaseetrogpessimisticForceBatch)
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
func (it *PolygonrollupbaseetrogpessimisticForceBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupbaseetrogpessimisticForceBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupbaseetrogpessimisticForceBatch represents a ForceBatch event raised by the Polygonrollupbaseetrogpessimistic contract.
type PolygonrollupbaseetrogpessimisticForceBatch struct {
	ForceBatchNum      uint64
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Transactions       []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterForceBatch is a free log retrieval operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) FilterForceBatch(opts *bind.FilterOpts, forceBatchNum []uint64) (*PolygonrollupbaseetrogpessimisticForceBatchIterator, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Polygonrollupbaseetrogpessimistic.contract.FilterLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpessimisticForceBatchIterator{contract: _Polygonrollupbaseetrogpessimistic.contract, event: "ForceBatch", logs: logs, sub: sub}, nil
}

// WatchForceBatch is a free log subscription operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) WatchForceBatch(opts *bind.WatchOpts, sink chan<- *PolygonrollupbaseetrogpessimisticForceBatch, forceBatchNum []uint64) (event.Subscription, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Polygonrollupbaseetrogpessimistic.contract.WatchLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupbaseetrogpessimisticForceBatch)
				if err := _Polygonrollupbaseetrogpessimistic.contract.UnpackLog(event, "ForceBatch", log); err != nil {
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
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) ParseForceBatch(log types.Log) (*PolygonrollupbaseetrogpessimisticForceBatch, error) {
	event := new(PolygonrollupbaseetrogpessimisticForceBatch)
	if err := _Polygonrollupbaseetrogpessimistic.contract.UnpackLog(event, "ForceBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupbaseetrogpessimisticInitialSequenceBatchesIterator is returned from FilterInitialSequenceBatches and is used to iterate over the raw logs and unpacked data for InitialSequenceBatches events raised by the Polygonrollupbaseetrogpessimistic contract.
type PolygonrollupbaseetrogpessimisticInitialSequenceBatchesIterator struct {
	Event *PolygonrollupbaseetrogpessimisticInitialSequenceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonrollupbaseetrogpessimisticInitialSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupbaseetrogpessimisticInitialSequenceBatches)
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
		it.Event = new(PolygonrollupbaseetrogpessimisticInitialSequenceBatches)
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
func (it *PolygonrollupbaseetrogpessimisticInitialSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupbaseetrogpessimisticInitialSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupbaseetrogpessimisticInitialSequenceBatches represents a InitialSequenceBatches event raised by the Polygonrollupbaseetrogpessimistic contract.
type PolygonrollupbaseetrogpessimisticInitialSequenceBatches struct {
	Transactions       []byte
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterInitialSequenceBatches is a free log retrieval operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) FilterInitialSequenceBatches(opts *bind.FilterOpts) (*PolygonrollupbaseetrogpessimisticInitialSequenceBatchesIterator, error) {

	logs, sub, err := _Polygonrollupbaseetrogpessimistic.contract.FilterLogs(opts, "InitialSequenceBatches")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpessimisticInitialSequenceBatchesIterator{contract: _Polygonrollupbaseetrogpessimistic.contract, event: "InitialSequenceBatches", logs: logs, sub: sub}, nil
}

// WatchInitialSequenceBatches is a free log subscription operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) WatchInitialSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonrollupbaseetrogpessimisticInitialSequenceBatches) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupbaseetrogpessimistic.contract.WatchLogs(opts, "InitialSequenceBatches")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupbaseetrogpessimisticInitialSequenceBatches)
				if err := _Polygonrollupbaseetrogpessimistic.contract.UnpackLog(event, "InitialSequenceBatches", log); err != nil {
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
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) ParseInitialSequenceBatches(log types.Log) (*PolygonrollupbaseetrogpessimisticInitialSequenceBatches, error) {
	event := new(PolygonrollupbaseetrogpessimisticInitialSequenceBatches)
	if err := _Polygonrollupbaseetrogpessimistic.contract.UnpackLog(event, "InitialSequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupbaseetrogpessimisticInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonrollupbaseetrogpessimistic contract.
type PolygonrollupbaseetrogpessimisticInitializedIterator struct {
	Event *PolygonrollupbaseetrogpessimisticInitialized // Event containing the contract specifics and raw log

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
func (it *PolygonrollupbaseetrogpessimisticInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupbaseetrogpessimisticInitialized)
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
		it.Event = new(PolygonrollupbaseetrogpessimisticInitialized)
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
func (it *PolygonrollupbaseetrogpessimisticInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupbaseetrogpessimisticInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupbaseetrogpessimisticInitialized represents a Initialized event raised by the Polygonrollupbaseetrogpessimistic contract.
type PolygonrollupbaseetrogpessimisticInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) FilterInitialized(opts *bind.FilterOpts) (*PolygonrollupbaseetrogpessimisticInitializedIterator, error) {

	logs, sub, err := _Polygonrollupbaseetrogpessimistic.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpessimisticInitializedIterator{contract: _Polygonrollupbaseetrogpessimistic.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PolygonrollupbaseetrogpessimisticInitialized) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupbaseetrogpessimistic.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupbaseetrogpessimisticInitialized)
				if err := _Polygonrollupbaseetrogpessimistic.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) ParseInitialized(log types.Log) (*PolygonrollupbaseetrogpessimisticInitialized, error) {
	event := new(PolygonrollupbaseetrogpessimisticInitialized)
	if err := _Polygonrollupbaseetrogpessimistic.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupbaseetrogpessimisticRollbackBatchesIterator is returned from FilterRollbackBatches and is used to iterate over the raw logs and unpacked data for RollbackBatches events raised by the Polygonrollupbaseetrogpessimistic contract.
type PolygonrollupbaseetrogpessimisticRollbackBatchesIterator struct {
	Event *PolygonrollupbaseetrogpessimisticRollbackBatches // Event containing the contract specifics and raw log

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
func (it *PolygonrollupbaseetrogpessimisticRollbackBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupbaseetrogpessimisticRollbackBatches)
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
		it.Event = new(PolygonrollupbaseetrogpessimisticRollbackBatches)
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
func (it *PolygonrollupbaseetrogpessimisticRollbackBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupbaseetrogpessimisticRollbackBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupbaseetrogpessimisticRollbackBatches represents a RollbackBatches event raised by the Polygonrollupbaseetrogpessimistic contract.
type PolygonrollupbaseetrogpessimisticRollbackBatches struct {
	TargetBatch            uint64
	AccInputHashToRollback [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRollbackBatches is a free log retrieval operation binding the contract event 0x1125aaf62d132d8e2d02005114f8fc360ff204c3105e4f1a700a1340dc55d5b1.
//
// Solidity: event RollbackBatches(uint64 indexed targetBatch, bytes32 accInputHashToRollback)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) FilterRollbackBatches(opts *bind.FilterOpts, targetBatch []uint64) (*PolygonrollupbaseetrogpessimisticRollbackBatchesIterator, error) {

	var targetBatchRule []interface{}
	for _, targetBatchItem := range targetBatch {
		targetBatchRule = append(targetBatchRule, targetBatchItem)
	}

	logs, sub, err := _Polygonrollupbaseetrogpessimistic.contract.FilterLogs(opts, "RollbackBatches", targetBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpessimisticRollbackBatchesIterator{contract: _Polygonrollupbaseetrogpessimistic.contract, event: "RollbackBatches", logs: logs, sub: sub}, nil
}

// WatchRollbackBatches is a free log subscription operation binding the contract event 0x1125aaf62d132d8e2d02005114f8fc360ff204c3105e4f1a700a1340dc55d5b1.
//
// Solidity: event RollbackBatches(uint64 indexed targetBatch, bytes32 accInputHashToRollback)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) WatchRollbackBatches(opts *bind.WatchOpts, sink chan<- *PolygonrollupbaseetrogpessimisticRollbackBatches, targetBatch []uint64) (event.Subscription, error) {

	var targetBatchRule []interface{}
	for _, targetBatchItem := range targetBatch {
		targetBatchRule = append(targetBatchRule, targetBatchItem)
	}

	logs, sub, err := _Polygonrollupbaseetrogpessimistic.contract.WatchLogs(opts, "RollbackBatches", targetBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupbaseetrogpessimisticRollbackBatches)
				if err := _Polygonrollupbaseetrogpessimistic.contract.UnpackLog(event, "RollbackBatches", log); err != nil {
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
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) ParseRollbackBatches(log types.Log) (*PolygonrollupbaseetrogpessimisticRollbackBatches, error) {
	event := new(PolygonrollupbaseetrogpessimisticRollbackBatches)
	if err := _Polygonrollupbaseetrogpessimistic.contract.UnpackLog(event, "RollbackBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupbaseetrogpessimisticSequenceBatchesIterator is returned from FilterSequenceBatches and is used to iterate over the raw logs and unpacked data for SequenceBatches events raised by the Polygonrollupbaseetrogpessimistic contract.
type PolygonrollupbaseetrogpessimisticSequenceBatchesIterator struct {
	Event *PolygonrollupbaseetrogpessimisticSequenceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonrollupbaseetrogpessimisticSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupbaseetrogpessimisticSequenceBatches)
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
		it.Event = new(PolygonrollupbaseetrogpessimisticSequenceBatches)
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
func (it *PolygonrollupbaseetrogpessimisticSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupbaseetrogpessimisticSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupbaseetrogpessimisticSequenceBatches represents a SequenceBatches event raised by the Polygonrollupbaseetrogpessimistic contract.
type PolygonrollupbaseetrogpessimisticSequenceBatches struct {
	NumBatch   uint64
	L1InfoRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSequenceBatches is a free log retrieval operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) FilterSequenceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonrollupbaseetrogpessimisticSequenceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonrollupbaseetrogpessimistic.contract.FilterLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpessimisticSequenceBatchesIterator{contract: _Polygonrollupbaseetrogpessimistic.contract, event: "SequenceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceBatches is a free log subscription operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) WatchSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonrollupbaseetrogpessimisticSequenceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonrollupbaseetrogpessimistic.contract.WatchLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupbaseetrogpessimisticSequenceBatches)
				if err := _Polygonrollupbaseetrogpessimistic.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
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
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) ParseSequenceBatches(log types.Log) (*PolygonrollupbaseetrogpessimisticSequenceBatches, error) {
	event := new(PolygonrollupbaseetrogpessimisticSequenceBatches)
	if err := _Polygonrollupbaseetrogpessimistic.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupbaseetrogpessimisticSequenceForceBatchesIterator is returned from FilterSequenceForceBatches and is used to iterate over the raw logs and unpacked data for SequenceForceBatches events raised by the Polygonrollupbaseetrogpessimistic contract.
type PolygonrollupbaseetrogpessimisticSequenceForceBatchesIterator struct {
	Event *PolygonrollupbaseetrogpessimisticSequenceForceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonrollupbaseetrogpessimisticSequenceForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupbaseetrogpessimisticSequenceForceBatches)
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
		it.Event = new(PolygonrollupbaseetrogpessimisticSequenceForceBatches)
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
func (it *PolygonrollupbaseetrogpessimisticSequenceForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupbaseetrogpessimisticSequenceForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupbaseetrogpessimisticSequenceForceBatches represents a SequenceForceBatches event raised by the Polygonrollupbaseetrogpessimistic contract.
type PolygonrollupbaseetrogpessimisticSequenceForceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceForceBatches is a free log retrieval operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) FilterSequenceForceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonrollupbaseetrogpessimisticSequenceForceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonrollupbaseetrogpessimistic.contract.FilterLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpessimisticSequenceForceBatchesIterator{contract: _Polygonrollupbaseetrogpessimistic.contract, event: "SequenceForceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceForceBatches is a free log subscription operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) WatchSequenceForceBatches(opts *bind.WatchOpts, sink chan<- *PolygonrollupbaseetrogpessimisticSequenceForceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonrollupbaseetrogpessimistic.contract.WatchLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupbaseetrogpessimisticSequenceForceBatches)
				if err := _Polygonrollupbaseetrogpessimistic.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
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
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) ParseSequenceForceBatches(log types.Log) (*PolygonrollupbaseetrogpessimisticSequenceForceBatches, error) {
	event := new(PolygonrollupbaseetrogpessimisticSequenceForceBatches)
	if err := _Polygonrollupbaseetrogpessimistic.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupbaseetrogpessimisticSetForceBatchAddressIterator is returned from FilterSetForceBatchAddress and is used to iterate over the raw logs and unpacked data for SetForceBatchAddress events raised by the Polygonrollupbaseetrogpessimistic contract.
type PolygonrollupbaseetrogpessimisticSetForceBatchAddressIterator struct {
	Event *PolygonrollupbaseetrogpessimisticSetForceBatchAddress // Event containing the contract specifics and raw log

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
func (it *PolygonrollupbaseetrogpessimisticSetForceBatchAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupbaseetrogpessimisticSetForceBatchAddress)
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
		it.Event = new(PolygonrollupbaseetrogpessimisticSetForceBatchAddress)
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
func (it *PolygonrollupbaseetrogpessimisticSetForceBatchAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupbaseetrogpessimisticSetForceBatchAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupbaseetrogpessimisticSetForceBatchAddress represents a SetForceBatchAddress event raised by the Polygonrollupbaseetrogpessimistic contract.
type PolygonrollupbaseetrogpessimisticSetForceBatchAddress struct {
	NewForceBatchAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchAddress is a free log retrieval operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) FilterSetForceBatchAddress(opts *bind.FilterOpts) (*PolygonrollupbaseetrogpessimisticSetForceBatchAddressIterator, error) {

	logs, sub, err := _Polygonrollupbaseetrogpessimistic.contract.FilterLogs(opts, "SetForceBatchAddress")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpessimisticSetForceBatchAddressIterator{contract: _Polygonrollupbaseetrogpessimistic.contract, event: "SetForceBatchAddress", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchAddress is a free log subscription operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) WatchSetForceBatchAddress(opts *bind.WatchOpts, sink chan<- *PolygonrollupbaseetrogpessimisticSetForceBatchAddress) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupbaseetrogpessimistic.contract.WatchLogs(opts, "SetForceBatchAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupbaseetrogpessimisticSetForceBatchAddress)
				if err := _Polygonrollupbaseetrogpessimistic.contract.UnpackLog(event, "SetForceBatchAddress", log); err != nil {
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
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) ParseSetForceBatchAddress(log types.Log) (*PolygonrollupbaseetrogpessimisticSetForceBatchAddress, error) {
	event := new(PolygonrollupbaseetrogpessimisticSetForceBatchAddress)
	if err := _Polygonrollupbaseetrogpessimistic.contract.UnpackLog(event, "SetForceBatchAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupbaseetrogpessimisticSetForceBatchTimeoutIterator is returned from FilterSetForceBatchTimeout and is used to iterate over the raw logs and unpacked data for SetForceBatchTimeout events raised by the Polygonrollupbaseetrogpessimistic contract.
type PolygonrollupbaseetrogpessimisticSetForceBatchTimeoutIterator struct {
	Event *PolygonrollupbaseetrogpessimisticSetForceBatchTimeout // Event containing the contract specifics and raw log

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
func (it *PolygonrollupbaseetrogpessimisticSetForceBatchTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupbaseetrogpessimisticSetForceBatchTimeout)
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
		it.Event = new(PolygonrollupbaseetrogpessimisticSetForceBatchTimeout)
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
func (it *PolygonrollupbaseetrogpessimisticSetForceBatchTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupbaseetrogpessimisticSetForceBatchTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupbaseetrogpessimisticSetForceBatchTimeout represents a SetForceBatchTimeout event raised by the Polygonrollupbaseetrogpessimistic contract.
type PolygonrollupbaseetrogpessimisticSetForceBatchTimeout struct {
	NewforceBatchTimeout uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchTimeout is a free log retrieval operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) FilterSetForceBatchTimeout(opts *bind.FilterOpts) (*PolygonrollupbaseetrogpessimisticSetForceBatchTimeoutIterator, error) {

	logs, sub, err := _Polygonrollupbaseetrogpessimistic.contract.FilterLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpessimisticSetForceBatchTimeoutIterator{contract: _Polygonrollupbaseetrogpessimistic.contract, event: "SetForceBatchTimeout", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchTimeout is a free log subscription operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) WatchSetForceBatchTimeout(opts *bind.WatchOpts, sink chan<- *PolygonrollupbaseetrogpessimisticSetForceBatchTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupbaseetrogpessimistic.contract.WatchLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupbaseetrogpessimisticSetForceBatchTimeout)
				if err := _Polygonrollupbaseetrogpessimistic.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
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
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) ParseSetForceBatchTimeout(log types.Log) (*PolygonrollupbaseetrogpessimisticSetForceBatchTimeout, error) {
	event := new(PolygonrollupbaseetrogpessimisticSetForceBatchTimeout)
	if err := _Polygonrollupbaseetrogpessimistic.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupbaseetrogpessimisticSetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the Polygonrollupbaseetrogpessimistic contract.
type PolygonrollupbaseetrogpessimisticSetTrustedSequencerIterator struct {
	Event *PolygonrollupbaseetrogpessimisticSetTrustedSequencer // Event containing the contract specifics and raw log

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
func (it *PolygonrollupbaseetrogpessimisticSetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupbaseetrogpessimisticSetTrustedSequencer)
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
		it.Event = new(PolygonrollupbaseetrogpessimisticSetTrustedSequencer)
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
func (it *PolygonrollupbaseetrogpessimisticSetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupbaseetrogpessimisticSetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupbaseetrogpessimisticSetTrustedSequencer represents a SetTrustedSequencer event raised by the Polygonrollupbaseetrogpessimistic contract.
type PolygonrollupbaseetrogpessimisticSetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*PolygonrollupbaseetrogpessimisticSetTrustedSequencerIterator, error) {

	logs, sub, err := _Polygonrollupbaseetrogpessimistic.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpessimisticSetTrustedSequencerIterator{contract: _Polygonrollupbaseetrogpessimistic.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *PolygonrollupbaseetrogpessimisticSetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupbaseetrogpessimistic.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupbaseetrogpessimisticSetTrustedSequencer)
				if err := _Polygonrollupbaseetrogpessimistic.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
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
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) ParseSetTrustedSequencer(log types.Log) (*PolygonrollupbaseetrogpessimisticSetTrustedSequencer, error) {
	event := new(PolygonrollupbaseetrogpessimisticSetTrustedSequencer)
	if err := _Polygonrollupbaseetrogpessimistic.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupbaseetrogpessimisticSetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the Polygonrollupbaseetrogpessimistic contract.
type PolygonrollupbaseetrogpessimisticSetTrustedSequencerURLIterator struct {
	Event *PolygonrollupbaseetrogpessimisticSetTrustedSequencerURL // Event containing the contract specifics and raw log

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
func (it *PolygonrollupbaseetrogpessimisticSetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupbaseetrogpessimisticSetTrustedSequencerURL)
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
		it.Event = new(PolygonrollupbaseetrogpessimisticSetTrustedSequencerURL)
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
func (it *PolygonrollupbaseetrogpessimisticSetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupbaseetrogpessimisticSetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupbaseetrogpessimisticSetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the Polygonrollupbaseetrogpessimistic contract.
type PolygonrollupbaseetrogpessimisticSetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*PolygonrollupbaseetrogpessimisticSetTrustedSequencerURLIterator, error) {

	logs, sub, err := _Polygonrollupbaseetrogpessimistic.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpessimisticSetTrustedSequencerURLIterator{contract: _Polygonrollupbaseetrogpessimistic.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *PolygonrollupbaseetrogpessimisticSetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupbaseetrogpessimistic.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupbaseetrogpessimisticSetTrustedSequencerURL)
				if err := _Polygonrollupbaseetrogpessimistic.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
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
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) ParseSetTrustedSequencerURL(log types.Log) (*PolygonrollupbaseetrogpessimisticSetTrustedSequencerURL, error) {
	event := new(PolygonrollupbaseetrogpessimisticSetTrustedSequencerURL)
	if err := _Polygonrollupbaseetrogpessimistic.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupbaseetrogpessimisticTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Polygonrollupbaseetrogpessimistic contract.
type PolygonrollupbaseetrogpessimisticTransferAdminRoleIterator struct {
	Event *PolygonrollupbaseetrogpessimisticTransferAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonrollupbaseetrogpessimisticTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupbaseetrogpessimisticTransferAdminRole)
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
		it.Event = new(PolygonrollupbaseetrogpessimisticTransferAdminRole)
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
func (it *PolygonrollupbaseetrogpessimisticTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupbaseetrogpessimisticTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupbaseetrogpessimisticTransferAdminRole represents a TransferAdminRole event raised by the Polygonrollupbaseetrogpessimistic contract.
type PolygonrollupbaseetrogpessimisticTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*PolygonrollupbaseetrogpessimisticTransferAdminRoleIterator, error) {

	logs, sub, err := _Polygonrollupbaseetrogpessimistic.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpessimisticTransferAdminRoleIterator{contract: _Polygonrollupbaseetrogpessimistic.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonrollupbaseetrogpessimisticTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupbaseetrogpessimistic.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupbaseetrogpessimisticTransferAdminRole)
				if err := _Polygonrollupbaseetrogpessimistic.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
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
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) ParseTransferAdminRole(log types.Log) (*PolygonrollupbaseetrogpessimisticTransferAdminRole, error) {
	event := new(PolygonrollupbaseetrogpessimisticTransferAdminRole)
	if err := _Polygonrollupbaseetrogpessimistic.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupbaseetrogpessimisticVerifyBatchesIterator is returned from FilterVerifyBatches and is used to iterate over the raw logs and unpacked data for VerifyBatches events raised by the Polygonrollupbaseetrogpessimistic contract.
type PolygonrollupbaseetrogpessimisticVerifyBatchesIterator struct {
	Event *PolygonrollupbaseetrogpessimisticVerifyBatches // Event containing the contract specifics and raw log

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
func (it *PolygonrollupbaseetrogpessimisticVerifyBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupbaseetrogpessimisticVerifyBatches)
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
		it.Event = new(PolygonrollupbaseetrogpessimisticVerifyBatches)
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
func (it *PolygonrollupbaseetrogpessimisticVerifyBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupbaseetrogpessimisticVerifyBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupbaseetrogpessimisticVerifyBatches represents a VerifyBatches event raised by the Polygonrollupbaseetrogpessimistic contract.
type PolygonrollupbaseetrogpessimisticVerifyBatches struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatches is a free log retrieval operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) FilterVerifyBatches(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*PolygonrollupbaseetrogpessimisticVerifyBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonrollupbaseetrogpessimistic.contract.FilterLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupbaseetrogpessimisticVerifyBatchesIterator{contract: _Polygonrollupbaseetrogpessimistic.contract, event: "VerifyBatches", logs: logs, sub: sub}, nil
}

// WatchVerifyBatches is a free log subscription operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) WatchVerifyBatches(opts *bind.WatchOpts, sink chan<- *PolygonrollupbaseetrogpessimisticVerifyBatches, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonrollupbaseetrogpessimistic.contract.WatchLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupbaseetrogpessimisticVerifyBatches)
				if err := _Polygonrollupbaseetrogpessimistic.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
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
func (_Polygonrollupbaseetrogpessimistic *PolygonrollupbaseetrogpessimisticFilterer) ParseVerifyBatches(log types.Log) (*PolygonrollupbaseetrogpessimisticVerifyBatches, error) {
	event := new(PolygonrollupbaseetrogpessimisticVerifyBatches)
	if err := _Polygonrollupbaseetrogpessimistic.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

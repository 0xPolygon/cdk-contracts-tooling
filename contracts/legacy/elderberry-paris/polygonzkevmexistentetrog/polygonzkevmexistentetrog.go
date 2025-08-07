// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonzkevmexistentetrog

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

// PolygonzkevmexistentetrogMetaData contains all meta data concerning the Polygonzkevmexistentetrog contract.
var PolygonzkevmexistentetrogMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"_rollupManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"forceBatchNum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"}],\"name\":\"ForceBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"}],\"name\":\"InitialSequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"l1InfoRoot\",\"type\":\"bytes32\"}],\"name\":\"SequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newForceBatchAddress\",\"type\":\"address\"}],\"name\":\"SetForceBatchAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"SetForceBatchTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"}],\"name\":\"UpdateEtrogSequence\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatches\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GLOBAL_EXIT_ROOT_MANAGER_L2\",\"outputs\":[{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_LIST_LEN_LEN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_DATA_LEN_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_EFFECTIVE_PERCENTAGE\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SET_UP_ETROG_TX\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_R\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_S\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_V\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TIMESTAMP_RANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculatePolPerForceBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"polAmount\",\"type\":\"uint256\"}],\"name\":\"forceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"}],\"name\":\"generateInitializeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trustedSequencer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_trustedSequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_lastAccInputHash\",\"type\":\"bytes32\"}],\"name\":\"initializeUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"onVerifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrog.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"maxSequenceTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initSequencedBatch\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"}],\"name\":\"sequenceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrog.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"}],\"name\":\"sequenceForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newForceBatchAddress\",\"type\":\"address\"}],\"name\":\"setForceBatchAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"setForceBatchTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b5060405162004b4538038062004b45833981016040819052620000359162000071565b6001600160a01b0393841660a052918316608052821660c0521660e052620000d9565b6001600160a01b03811681146200006e57600080fd5b50565b600080600080608085870312156200008857600080fd5b8451620000958162000058565b6020860151909450620000a88162000058565b6040860151909350620000bb8162000058565b6060860151909250620000ce8162000058565b939692955090935050565b60805160a05160c05160e05161496e620001d760003960008181610519015281816109a201528181610b0f01528181610c9001528181610fe7015281816112f1015281816118d201528181611d76015281816121cd015281816122c301528181612ecd01528181612f4601528181612f68015281816130800152818161322301526132eb015260008181610686015281816114e5015281816115bf015281816124940152818161259c0152612a0a01526000818161074a01528181610e580152818161174601528181612a8c015261343601526000818161078f0152818161084e0152818161221601528181613016015261340a015261496e6000f3fe608060405234801561001057600080fd5b50600436106102fe5760003560e01c8063712570221161019c578063c754c7ed116100ee578063def57e5411610097578063eaeb077b11610071578063eaeb077b146107c5578063f35dda47146107d8578063f851a440146107e057600080fd5b8063def57e5414610777578063e46761c41461078a578063e7a7ed02146107b157600080fd5b8063cfa8ed47116100c8578063cfa8ed4714610725578063d02103ca14610745578063d7bc90ff1461076c57600080fd5b8063c754c7ed146106e2578063c7fffd4b1461070a578063c89e42df1461071257600080fd5b80639f26f84011610150578063ada8f9191161012a578063ada8f919146106bb578063af7f3e02146106ce578063b0afe154146106d657600080fd5b80639f26f8401461066e578063a3c573eb14610681578063a652f26c146106a857600080fd5b80638c3d7301116101815780638c3d73011461063857806391cafe32146106405780639e0018771461065357600080fd5b806371257022146105e95780637a5460c5146105fc57600080fd5b806342308fab11610255578063542028d5116102095780636b8616ce116101e35780636b8616ce146105ad5780636e05d2cd146105cd5780636ff512cc146105d657600080fd5b8063542028d51461058a5780635d6717a514610592578063676870d2146105a557600080fd5b806349b7b8021161023a57806349b7b802146105145780634e4877061461053b57806352bdeb6d1461054e57600080fd5b806342308fab146104d357806345605267146104db57600080fd5b806326782247116102b75780633c351e10116102915780633c351e101461041e5780633cbc795b1461043e57806340b5de6c1461047b57600080fd5b806326782247146103a45780632c111c06146103e957806332c2d1531461040957600080fd5b806305835f37116102e857806305835f3714610339578063107bf28c1461038257806311e892d41461038a57600080fd5b8062d0295d14610303578063035089631461031e575b600080fd5b61030b610806565b6040519081526020015b60405180910390f35b610326602081565b60405161ffff9091168152602001610315565b6103756040518060400160405280600881526020017f80808401c9c3809400000000000000000000000000000000000000000000000081525081565b6040516103159190613aa5565b610375610912565b61039260f981565b60405160ff9091168152602001610315565b6001546103c49073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610315565b6008546103c49073ffffffffffffffffffffffffffffffffffffffff1681565b61041c610417366004613afa565b6109a0565b005b6009546103c49073ffffffffffffffffffffffffffffffffffffffff1681565b6009546104669074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff9091168152602001610315565b6104a27fff0000000000000000000000000000000000000000000000000000000000000081565b6040517fff000000000000000000000000000000000000000000000000000000000000009091168152602001610315565b61030b602481565b6007546104fb9068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff9091168152602001610315565b6103c47f000000000000000000000000000000000000000000000000000000000000000081565b61041c610549366004613b3c565b610a6f565b6103756040518060400160405280600281526020017f80b800000000000000000000000000000000000000000000000000000000000081525081565b610375610c81565b61041c6105a0366004613c73565b610c8e565b610326601f81565b61030b6105bb366004613b3c565b60066020526000908152604090205481565b61030b60055481565b61041c6105e4366004613d04565b611225565b61041c6105f7366004613d33565b6112ef565b6103756040518060400160405280600281526020017f80b900000000000000000000000000000000000000000000000000000000000081525081565b61041c611b0e565b61041c61064e366004613d04565b611be1565b6103c473a40d5f56745a118d0906a34e69aec8c0db1cb8fa81565b61041c61067c366004613e2c565b611cfa565b6103c47f000000000000000000000000000000000000000000000000000000000000000081565b6103756106b6366004613e6e565b612393565b61041c6106c9366004613d04565b612778565b610375612842565b61030b6405ca1ab1e081565b6007546104fb90700100000000000000000000000000000000900467ffffffffffffffff1681565b61039260e481565b61041c610720366004613ee3565b61285e565b6002546103c49073ffffffffffffffffffffffffffffffffffffffff1681565b6103c47f000000000000000000000000000000000000000000000000000000000000000081565b61030b635ca1ab1e81565b61041c610785366004613f18565b6128f1565b6103c47f000000000000000000000000000000000000000000000000000000000000000081565b6007546104fb9067ffffffffffffffff1681565b61041c6107d3366004613f95565b6131ac565b610392601b81565b6000546103c49062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b6040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152600090819073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015610895573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108b9919061400d565b6007549091506000906108e49067ffffffffffffffff68010000000000000000820481169116614055565b67ffffffffffffffff169050806000036109015760009250505090565b61090b818361407d565b9250505090565b6004805461091f906140b8565b80601f016020809104026020016040519081016040528092919081815260200182805461094b906140b8565b80156109985780601f1061096d57610100808354040283529160200191610998565b820191906000526020600020905b81548152906001019060200180831161097b57829003601f168201915b505050505081565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314610a0f576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff168367ffffffffffffffff167f9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f596684604051610a6291815260200190565b60405180910390a3505050565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314610ac6576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff82161115610b0d576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166315064c966040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b78573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b9c919061410b565b610bfd5760075467ffffffffffffffff700100000000000000000000000000000000909104811690821610610bfd576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600780547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527fa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b906020015b60405180910390a150565b6003805461091f906140b8565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314610cfd576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600054610100900460ff1615808015610d1d5750600054600160ff909116105b80610d375750303b158015610d37575060005460ff166001145b610dc8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610e2657600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b60006040518060a00160405280606281526020016148d7606291399050600081805190602001209050600042905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ec1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ee5919061400d565b90506000868483858d610ef960014361412d565b60408051602081019790975286019490945260608086019390935260c09190911b7fffffffffffffffff000000000000000000000000000000000000000000000000166080850152901b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016608883015240609c82015260bc01604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001815290829052805160209091012060058190557f9a908e73000000000000000000000000000000000000000000000000000000008252600160048301526024820181905291506000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303816000875af1158015611045573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110699190614146565b90508b600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508a600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555089600390816110fc91906141a9565b5060046111098a826141a9565b508b600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062069780600760106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fd2c80353fc15ef62c6affc7cd6b7ab5b42c43290c50be3372e55ae552cecd19c8187858e6040516111ac94939291906142c3565b60405180910390a1505050505050801561121d57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff16331461127c576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc090602001610c76565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16331461135e576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600054610100900460ff161580801561137e5750600054600160ff909116105b806113985750303b158015611398575060005460ff166001145b611424576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610dbf565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561148257600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b606073ffffffffffffffffffffffffffffffffffffffff8516156116e7576040517fc00f14ab00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063c00f14ab90602401600060405180830381865afa15801561152c573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526115729190810190614313565b6040517f318aee3d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff878116600483015291925060009182917f00000000000000000000000000000000000000000000000000000000000000009091169063318aee3d906024016040805180830381865afa158015611607573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061162b919061438a565b915091508163ffffffff166000146116a3576009805463ffffffff841674010000000000000000000000000000000000000000027fffffffffffffffff00000000000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff8416171790556116e4565b600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff89161790555b50505b60095460009061172f90889073ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900463ffffffff1685612393565b9050600081805190602001209050600042905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117af573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117d3919061400d565b90506000808483858f6117e760014361412d565b60408051602081019790975286019490945260608086019390935260c09190911b7fffffffffffffffff000000000000000000000000000000000000000000000000166080850152901b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016608883015240609c82015260bc01604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001815290829052805160209091012060058190557f9a908e73000000000000000000000000000000000000000000000000000000008252600160048301526024820181905291507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303816000875af1158015611930573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119549190614146565b508c600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508b600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555088600390816119e691906141a9565b5060046119f389826141a9565b508c600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062069780600760106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507f060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f85838e604051611a94939291906143c4565b60405180910390a15050505050508015611b0557600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314611b5f576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600154600080547fffffffffffffffffffff0000000000000000000000000000000000000000ffff1673ffffffffffffffffffffffffffffffffffffffff9092166201000081029290921790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e9060200160405180910390a1565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314611c38576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60085473ffffffffffffffffffffffffffffffffffffffff16611c87576040517fc89374d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb90602001610c76565b60085473ffffffffffffffffffffffffffffffffffffffff168015801590611d38575073ffffffffffffffffffffffffffffffffffffffff81163314155b15611d6f576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b4262093a807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166330c27dde6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611ddf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e039190614146565b611e0d9190614403565b67ffffffffffffffff161115611e4f576040517f3d49ed4c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b816000819003611e8b576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8811115611ec7576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075467ffffffffffffffff80821691611eef91849168010000000000000000900416614424565b1115611f27576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6007546005546801000000000000000090910467ffffffffffffffff169060005b838110156121c7576000878783818110611f6457611f64614437565b9050602002810190611f769190614466565b611f7f906144a4565b905083611f8b8161452d565b825180516020918201208185015160408087015160608801519151959a50929550600094611ff8948794929101938452602084019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166040830152604882015260680190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff8916600090815260069093529120549091508114612081576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff86166000908152600660205260408120556120a660018861412d565b84036121155742600760109054906101000a900467ffffffffffffffff1684604001516120d39190614403565b67ffffffffffffffff161115612115576040517fc44a082100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208381015160408086015160608088015183519586018b90529285018790528481019390935260c01b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808401523390911b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088830152609c82015260bc0160405160208183030381529060405280519060200120945050505080806121bf90614554565b915050611f48565b5061223d7f0000000000000000000000000000000000000000000000000000000000000000846121f5610806565b6121ff919061458c565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016919061367d565b60058190556007805467ffffffffffffffff841668010000000000000000027fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff9091161790556040517f9a908e7300000000000000000000000000000000000000000000000000000000815260009073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690639a908e739061230f908790869060040167ffffffffffffffff929092168252602082015260400190565b6020604051808303816000875af115801561232e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123529190614146565b60405190915067ffffffffffffffff8216907f648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a490600090a250505050505050565b6060600085858573a40d5f56745a118d0906a34e69aec8c0db1cb8fa6000876040516024016123c7969594939291906145a3565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff811bff70000000000000000000000000000000000000000000000000000000017905283519091506060906000036125185760f9601f835161245c9190614606565b6040518060400160405280600881526020017f80808401c9c380940000000000000000000000000000000000000000000000008152507f00000000000000000000000000000000000000000000000000000000000000006040518060400160405280600281526020017f80b800000000000000000000000000000000000000000000000000000000000081525060e4876040516020016125029796959493929190614621565b604051602081830303815290604052905061261c565b815161ffff1015612555576040517f248b8f8200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815160f9612564602083614606565b6040518060400160405280600881526020017f80808401c9c380940000000000000000000000000000000000000000000000008152507f00000000000000000000000000000000000000000000000000000000000000006040518060400160405280600281526020017f80b900000000000000000000000000000000000000000000000000000000000081525085886040516020016126099796959493929190614704565b6040516020818303038152906040529150505b805160208083019190912060408051600080825293810180835292909252601b908201526405ca1ab1e06060820152635ca1ab1e608082015260019060a0016020604051602081039080840390855afa15801561267d573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81166126f5576040517fcd16196600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60405160009061273b9084906405ca1ab1e090635ca1ab1e90601b907fff00000000000000000000000000000000000000000000000000000000000000906020016147e7565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529450505050505b949350505050565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff1633146127cf576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001610c76565b6040518060a00160405280606281526020016148d76062913981565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff1633146128b5576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60036128c182826141a9565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b2081604051610c769190613aa5565b60025473ffffffffffffffffffffffffffffffffffffffff163314612942576040517f11e7be1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83600081900361297e576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e88111156129ba576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6129c5602442614424565b8467ffffffffffffffff161115612a08576040517f0a00feb300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379e2cf976040518163ffffffff1660e01b8152600401600060405180830381600087803b158015612a7057600080fd5b505af1158015612a84573d6000803e3d6000fd5b5050505060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635ca1e1656040518163ffffffff1660e01b8152600401602060405180830381865afa158015612af5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b19919061400d565b60075460055491925068010000000000000000900467ffffffffffffffff16908160005b85811015612e3f5760008b8b83818110612b5957612b59614437565b9050602002810190612b6b9190614466565b612b74906144a4565b8051805160209091012060408201519192509067ffffffffffffffff1615612d595785612ba08161452d565b965050600081836020015184604001518560600151604051602001612c039493929190938452602084019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166040830152604882015260680190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff8a16600090815260069093529120549091508114612c8c576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208381015160408086015160608088015183519586018c90529285018790528481019390935260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166080840152908c901b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088830152609c82015260bc01604051602081830303815290604052805190602001209550600660008867ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206000905550612e2a565b8151516201d4c01015612d98576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160208101879052908101829052606080820189905260c08d901b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808301528a901b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660888201526000609c82015260bc016040516020818303038152906040528051906020012094505b50508080612e3790614554565b915050612b3d565b5060075467ffffffffffffffff9081169084161115612e8a576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60058290558467ffffffffffffffff84811690831614612f40576000612eb08386614055565b9050612ec667ffffffffffffffff82168361412d565b9150612eff7f00000000000000000000000000000000000000000000000000000000000000008267ffffffffffffffff166121f5610806565b50600780547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8716021790555b61303e337f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663477fa2706040518163ffffffff1660e01b8152600401602060405180830381865afa158015612fd1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ff5919061400d565b612fff919061458c565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016929190613756565b6040517f9a908e7300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff87166004820152602481018490526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303816000875af11580156130de573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131029190614146565b905061310e8782614055565b67ffffffffffffffff168967ffffffffffffffff161461315a576040517f1a070d9a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff167f3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e7668760405161319691815260200190565b60405180910390a2505050505050505050505050565b60085473ffffffffffffffffffffffffffffffffffffffff1680158015906131ea575073ffffffffffffffffffffffffffffffffffffffff81163314155b15613221576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166315064c966040518163ffffffff1660e01b8152600401602060405180830381865afa15801561328c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132b0919061410b565b156132e7576040517f39258d1800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663604691696040518163ffffffff1660e01b8152600401602060405180830381865afa158015613354573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613378919061400d565b9050828111156133b4576040517f2354600f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113888411156133f0576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61343273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016333084613756565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561349f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134c3919061400d565b6007805491925067ffffffffffffffff9091169060006134e28361452d565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550508585604051613519929190614843565b604051908190039020814261352f60014361412d565b60408051602081019590955284019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166060830152406068820152608801604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012060075467ffffffffffffffff166000908152600690935291205532330361362757600754604080518381523360208201526060818301819052600090820152905167ffffffffffffffff909216917ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319181900360800190a261121d565b60075460405167ffffffffffffffff909116907ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319061366d90849033908b908b90614853565b60405180910390a2505050505050565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526137519084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526137ba565b505050565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526137b49085907f23b872dd00000000000000000000000000000000000000000000000000000000906084016136cf565b50505050565b600061381c826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166138c69092919063ffffffff16565b805190915015613751578080602001905181019061383a919061410b565b613751576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610dbf565b60606127708484600085856000808673ffffffffffffffffffffffffffffffffffffffff1685876040516138fa91906148c4565b60006040518083038185875af1925050503d8060008114613937576040519150601f19603f3d011682016040523d82523d6000602084013e61393c565b606091505b509150915061394d87838387613958565b979650505050505050565b606083156139ee5782516000036139e75773ffffffffffffffffffffffffffffffffffffffff85163b6139e7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610dbf565b5081612770565b6127708383815115613a035781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dbf9190613aa5565b60005b83811015613a52578181015183820152602001613a3a565b50506000910152565b60008151808452613a73816020860160208601613a37565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000613ab86020830184613a5b565b9392505050565b67ffffffffffffffff81168114613ad557600080fd5b50565b73ffffffffffffffffffffffffffffffffffffffff81168114613ad557600080fd5b600080600060608486031215613b0f57600080fd5b8335613b1a81613abf565b9250602084013591506040840135613b3181613ad8565b809150509250925092565b600060208284031215613b4e57600080fd5b8135613ab881613abf565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613bcf57613bcf613b59565b604052919050565b600067ffffffffffffffff821115613bf157613bf1613b59565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f830112613c2e57600080fd5b8135613c41613c3c82613bd7565b613b88565b818152846020838601011115613c5657600080fd5b816020850160208301376000918101602001919091529392505050565b600080600080600060a08688031215613c8b57600080fd5b8535613c9681613ad8565b94506020860135613ca681613ad8565b9350604086013567ffffffffffffffff80821115613cc357600080fd5b613ccf89838a01613c1d565b94506060880135915080821115613ce557600080fd5b50613cf288828901613c1d565b95989497509295608001359392505050565b600060208284031215613d1657600080fd5b8135613ab881613ad8565b63ffffffff81168114613ad557600080fd5b60008060008060008060c08789031215613d4c57600080fd5b8635613d5781613ad8565b95506020870135613d6781613ad8565b94506040870135613d7781613d21565b93506060870135613d8781613ad8565b9250608087013567ffffffffffffffff80821115613da457600080fd5b613db08a838b01613c1d565b935060a0890135915080821115613dc657600080fd5b50613dd389828a01613c1d565b9150509295509295509295565b60008083601f840112613df257600080fd5b50813567ffffffffffffffff811115613e0a57600080fd5b6020830191508360208260051b8501011115613e2557600080fd5b9250929050565b60008060208385031215613e3f57600080fd5b823567ffffffffffffffff811115613e5657600080fd5b613e6285828601613de0565b90969095509350505050565b60008060008060808587031215613e8457600080fd5b8435613e8f81613d21565b93506020850135613e9f81613ad8565b92506040850135613eaf81613d21565b9150606085013567ffffffffffffffff811115613ecb57600080fd5b613ed787828801613c1d565b91505092959194509250565b600060208284031215613ef557600080fd5b813567ffffffffffffffff811115613f0c57600080fd5b61277084828501613c1d565b600080600080600060808688031215613f3057600080fd5b853567ffffffffffffffff811115613f4757600080fd5b613f5388828901613de0565b9096509450506020860135613f6781613abf565b92506040860135613f7781613abf565b91506060860135613f8781613ad8565b809150509295509295909350565b600080600060408486031215613faa57600080fd5b833567ffffffffffffffff80821115613fc257600080fd5b818601915086601f830112613fd657600080fd5b813581811115613fe557600080fd5b876020828501011115613ff757600080fd5b6020928301989097509590910135949350505050565b60006020828403121561401f57600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b67ffffffffffffffff82811682821603908082111561407657614076614026565b5092915050565b6000826140b3577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b600181811c908216806140cc57607f821691505b602082108103614105577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b60006020828403121561411d57600080fd5b81518015158114613ab857600080fd5b8181038181111561414057614140614026565b92915050565b60006020828403121561415857600080fd5b8151613ab881613abf565b601f82111561375157600081815260208120601f850160051c8101602086101561418a5750805b601f850160051c820191505b8181101561121d57828155600101614196565b815167ffffffffffffffff8111156141c3576141c3613b59565b6141d7816141d184546140b8565b84614163565b602080601f83116001811461422a57600084156141f45750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b17855561121d565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b8281101561427757888601518255948401946001909101908401614258565b50858210156142b357878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b67ffffffffffffffff851681526080602082015260006142e66080830186613a5b565b905083604083015273ffffffffffffffffffffffffffffffffffffffff8316606083015295945050505050565b60006020828403121561432557600080fd5b815167ffffffffffffffff81111561433c57600080fd5b8201601f8101841361434d57600080fd5b805161435b613c3c82613bd7565b81815285602083850101111561437057600080fd5b614381826020830160208601613a37565b95945050505050565b6000806040838503121561439d57600080fd5b82516143a881613d21565b60208401519092506143b981613ad8565b809150509250929050565b6060815260006143d76060830186613a5b565b905083602083015273ffffffffffffffffffffffffffffffffffffffff83166040830152949350505050565b67ffffffffffffffff81811683821601908082111561407657614076614026565b8082018082111561414057614140614026565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8183360301811261449a57600080fd5b9190910192915050565b6000608082360312156144b657600080fd5b6040516080810167ffffffffffffffff82821081831117156144da576144da613b59565b8160405284359150808211156144ef57600080fd5b506144fc36828601613c1d565b82525060208301356020820152604083013561451781613abf565b6040820152606092830135928101929092525090565b600067ffffffffffffffff80831681810361454a5761454a614026565b6001019392505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361458557614585614026565b5060010190565b808202811582820484141761414057614140614026565b600063ffffffff808916835273ffffffffffffffffffffffffffffffffffffffff8089166020850152818816604085015280871660608501528086166080850152505060c060a08301526145fa60c0830184613a5b565b98975050505050505050565b61ffff81811683821601908082111561407657614076614026565b60007fff00000000000000000000000000000000000000000000000000000000000000808a60f81b1683527fffff0000000000000000000000000000000000000000000000000000000000008960f01b166001840152875161468a816003860160208c01613a37565b80840190507fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008860601b16600382015286516146cd816017840160208b01613a37565b808201915050818660f81b166017820152845191506146f3826018830160208801613a37565b016018019998505050505050505050565b7fff000000000000000000000000000000000000000000000000000000000000008860f81b16815260007fffff000000000000000000000000000000000000000000000000000000000000808960f01b166001840152875161476d816003860160208c01613a37565b80840190507fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008860601b16600382015286516147b0816017840160208b01613a37565b808201915050818660f01b166017820152845191506147d6826019830160208801613a37565b016019019998505050505050505050565b600086516147f9818460208b01613a37565b9190910194855250602084019290925260f81b7fff000000000000000000000000000000000000000000000000000000000000009081166040840152166041820152604201919050565b8183823760009101908152919050565b84815273ffffffffffffffffffffffffffffffffffffffff8416602082015260606040820152816060820152818360808301376000818301608090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01601019392505050565b6000825161449a818460208701613a3756fedf2a8080944d5cf5032b2a844602278b01199ed191a86c93ff8080821092808000000000000000000000000000000000000000000000000000000005ca1ab1e000000000000000000000000000000000000000000000000000000005ca1ab1e01bffa26469706673582212200373944e7c0220cf416ef0ded0f2ced76f28d9e68ccf98092c7452d6cd61826f64736f6c63430008140033",
}

// PolygonzkevmexistentetrogABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonzkevmexistentetrogMetaData.ABI instead.
var PolygonzkevmexistentetrogABI = PolygonzkevmexistentetrogMetaData.ABI

// PolygonzkevmexistentetrogBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygonzkevmexistentetrogMetaData.Bin instead.
var PolygonzkevmexistentetrogBin = PolygonzkevmexistentetrogMetaData.Bin

// DeployPolygonzkevmexistentetrog deploys a new Ethereum contract, binding an instance of Polygonzkevmexistentetrog to it.
func DeployPolygonzkevmexistentetrog(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address, _rollupManager common.Address) (common.Address, *types.Transaction, *Polygonzkevmexistentetrog, error) {
	parsed, err := PolygonzkevmexistentetrogMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygonzkevmexistentetrogBin), backend, _globalExitRootManager, _pol, _bridgeAddress, _rollupManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonzkevmexistentetrog{PolygonzkevmexistentetrogCaller: PolygonzkevmexistentetrogCaller{contract: contract}, PolygonzkevmexistentetrogTransactor: PolygonzkevmexistentetrogTransactor{contract: contract}, PolygonzkevmexistentetrogFilterer: PolygonzkevmexistentetrogFilterer{contract: contract}}, nil
}

// Polygonzkevmexistentetrog is an auto generated Go binding around an Ethereum contract.
type Polygonzkevmexistentetrog struct {
	PolygonzkevmexistentetrogCaller     // Read-only binding to the contract
	PolygonzkevmexistentetrogTransactor // Write-only binding to the contract
	PolygonzkevmexistentetrogFilterer   // Log filterer for contract events
}

// PolygonzkevmexistentetrogCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonzkevmexistentetrogCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmexistentetrogTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonzkevmexistentetrogTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmexistentetrogFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonzkevmexistentetrogFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmexistentetrogSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonzkevmexistentetrogSession struct {
	Contract     *Polygonzkevmexistentetrog // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PolygonzkevmexistentetrogCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonzkevmexistentetrogCallerSession struct {
	Contract *PolygonzkevmexistentetrogCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// PolygonzkevmexistentetrogTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonzkevmexistentetrogTransactorSession struct {
	Contract     *PolygonzkevmexistentetrogTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// PolygonzkevmexistentetrogRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonzkevmexistentetrogRaw struct {
	Contract *Polygonzkevmexistentetrog // Generic contract binding to access the raw methods on
}

// PolygonzkevmexistentetrogCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonzkevmexistentetrogCallerRaw struct {
	Contract *PolygonzkevmexistentetrogCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonzkevmexistentetrogTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonzkevmexistentetrogTransactorRaw struct {
	Contract *PolygonzkevmexistentetrogTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonzkevmexistentetrog creates a new instance of Polygonzkevmexistentetrog, bound to a specific deployed contract.
func NewPolygonzkevmexistentetrog(address common.Address, backend bind.ContractBackend) (*Polygonzkevmexistentetrog, error) {
	contract, err := bindPolygonzkevmexistentetrog(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmexistentetrog{PolygonzkevmexistentetrogCaller: PolygonzkevmexistentetrogCaller{contract: contract}, PolygonzkevmexistentetrogTransactor: PolygonzkevmexistentetrogTransactor{contract: contract}, PolygonzkevmexistentetrogFilterer: PolygonzkevmexistentetrogFilterer{contract: contract}}, nil
}

// NewPolygonzkevmexistentetrogCaller creates a new read-only instance of Polygonzkevmexistentetrog, bound to a specific deployed contract.
func NewPolygonzkevmexistentetrogCaller(address common.Address, caller bind.ContractCaller) (*PolygonzkevmexistentetrogCaller, error) {
	contract, err := bindPolygonzkevmexistentetrog(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogCaller{contract: contract}, nil
}

// NewPolygonzkevmexistentetrogTransactor creates a new write-only instance of Polygonzkevmexistentetrog, bound to a specific deployed contract.
func NewPolygonzkevmexistentetrogTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonzkevmexistentetrogTransactor, error) {
	contract, err := bindPolygonzkevmexistentetrog(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogTransactor{contract: contract}, nil
}

// NewPolygonzkevmexistentetrogFilterer creates a new log filterer instance of Polygonzkevmexistentetrog, bound to a specific deployed contract.
func NewPolygonzkevmexistentetrogFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonzkevmexistentetrogFilterer, error) {
	contract, err := bindPolygonzkevmexistentetrog(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogFilterer{contract: contract}, nil
}

// bindPolygonzkevmexistentetrog binds a generic wrapper to an already deployed contract.
func bindPolygonzkevmexistentetrog(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonzkevmexistentetrogMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmexistentetrog.Contract.PolygonzkevmexistentetrogCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.PolygonzkevmexistentetrogTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.PolygonzkevmexistentetrogTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmexistentetrog.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.contract.Transact(opts, method, params...)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) GLOBALEXITROOTMANAGERL2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "GLOBAL_EXIT_ROOT_MANAGER_L2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _Polygonzkevmexistentetrog.Contract.GLOBALEXITROOTMANAGERL2(&_Polygonzkevmexistentetrog.CallOpts)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _Polygonzkevmexistentetrog.Contract.GLOBALEXITROOTMANAGERL2(&_Polygonzkevmexistentetrog.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) INITIALIZETXBRIDGELISTLENLEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_LIST_LEN_LEN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _Polygonzkevmexistentetrog.Contract.INITIALIZETXBRIDGELISTLENLEN(&_Polygonzkevmexistentetrog.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _Polygonzkevmexistentetrog.Contract.INITIALIZETXBRIDGELISTLENLEN(&_Polygonzkevmexistentetrog.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) INITIALIZETXBRIDGEPARAMS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _Polygonzkevmexistentetrog.Contract.INITIALIZETXBRIDGEPARAMS(&_Polygonzkevmexistentetrog.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _Polygonzkevmexistentetrog.Contract.INITIALIZETXBRIDGEPARAMS(&_Polygonzkevmexistentetrog.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _Polygonzkevmexistentetrog.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_Polygonzkevmexistentetrog.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _Polygonzkevmexistentetrog.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_Polygonzkevmexistentetrog.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _Polygonzkevmexistentetrog.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_Polygonzkevmexistentetrog.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _Polygonzkevmexistentetrog.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_Polygonzkevmexistentetrog.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) INITIALIZETXCONSTANTBYTES(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _Polygonzkevmexistentetrog.Contract.INITIALIZETXCONSTANTBYTES(&_Polygonzkevmexistentetrog.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _Polygonzkevmexistentetrog.Contract.INITIALIZETXCONSTANTBYTES(&_Polygonzkevmexistentetrog.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) INITIALIZETXCONSTANTBYTESEMPTYMETADATA(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _Polygonzkevmexistentetrog.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_Polygonzkevmexistentetrog.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _Polygonzkevmexistentetrog.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_Polygonzkevmexistentetrog.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) INITIALIZETXDATALENEMPTYMETADATA(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "INITIALIZE_TX_DATA_LEN_EMPTY_METADATA")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _Polygonzkevmexistentetrog.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_Polygonzkevmexistentetrog.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _Polygonzkevmexistentetrog.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_Polygonzkevmexistentetrog.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) INITIALIZETXEFFECTIVEPERCENTAGE(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "INITIALIZE_TX_EFFECTIVE_PERCENTAGE")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _Polygonzkevmexistentetrog.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_Polygonzkevmexistentetrog.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _Polygonzkevmexistentetrog.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_Polygonzkevmexistentetrog.CallOpts)
}

// SETUPETROGTX is a free data retrieval call binding the contract method 0xaf7f3e02.
//
// Solidity: function SET_UP_ETROG_TX() view returns(bytes)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) SETUPETROGTX(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "SET_UP_ETROG_TX")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SETUPETROGTX is a free data retrieval call binding the contract method 0xaf7f3e02.
//
// Solidity: function SET_UP_ETROG_TX() view returns(bytes)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) SETUPETROGTX() ([]byte, error) {
	return _Polygonzkevmexistentetrog.Contract.SETUPETROGTX(&_Polygonzkevmexistentetrog.CallOpts)
}

// SETUPETROGTX is a free data retrieval call binding the contract method 0xaf7f3e02.
//
// Solidity: function SET_UP_ETROG_TX() view returns(bytes)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) SETUPETROGTX() ([]byte, error) {
	return _Polygonzkevmexistentetrog.Contract.SETUPETROGTX(&_Polygonzkevmexistentetrog.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) SIGNATUREINITIALIZETXR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_R")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _Polygonzkevmexistentetrog.Contract.SIGNATUREINITIALIZETXR(&_Polygonzkevmexistentetrog.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _Polygonzkevmexistentetrog.Contract.SIGNATUREINITIALIZETXR(&_Polygonzkevmexistentetrog.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) SIGNATUREINITIALIZETXS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_S")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _Polygonzkevmexistentetrog.Contract.SIGNATUREINITIALIZETXS(&_Polygonzkevmexistentetrog.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _Polygonzkevmexistentetrog.Contract.SIGNATUREINITIALIZETXS(&_Polygonzkevmexistentetrog.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) SIGNATUREINITIALIZETXV(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_V")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _Polygonzkevmexistentetrog.Contract.SIGNATUREINITIALIZETXV(&_Polygonzkevmexistentetrog.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _Polygonzkevmexistentetrog.Contract.SIGNATUREINITIALIZETXV(&_Polygonzkevmexistentetrog.CallOpts)
}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) TIMESTAMPRANGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "TIMESTAMP_RANGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) TIMESTAMPRANGE() (*big.Int, error) {
	return _Polygonzkevmexistentetrog.Contract.TIMESTAMPRANGE(&_Polygonzkevmexistentetrog.CallOpts)
}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) TIMESTAMPRANGE() (*big.Int, error) {
	return _Polygonzkevmexistentetrog.Contract.TIMESTAMPRANGE(&_Polygonzkevmexistentetrog.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) Admin() (common.Address, error) {
	return _Polygonzkevmexistentetrog.Contract.Admin(&_Polygonzkevmexistentetrog.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) Admin() (common.Address, error) {
	return _Polygonzkevmexistentetrog.Contract.Admin(&_Polygonzkevmexistentetrog.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) BridgeAddress() (common.Address, error) {
	return _Polygonzkevmexistentetrog.Contract.BridgeAddress(&_Polygonzkevmexistentetrog.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) BridgeAddress() (common.Address, error) {
	return _Polygonzkevmexistentetrog.Contract.BridgeAddress(&_Polygonzkevmexistentetrog.CallOpts)
}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) CalculatePolPerForceBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "calculatePolPerForceBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) CalculatePolPerForceBatch() (*big.Int, error) {
	return _Polygonzkevmexistentetrog.Contract.CalculatePolPerForceBatch(&_Polygonzkevmexistentetrog.CallOpts)
}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) CalculatePolPerForceBatch() (*big.Int, error) {
	return _Polygonzkevmexistentetrog.Contract.CalculatePolPerForceBatch(&_Polygonzkevmexistentetrog.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) ForceBatchAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "forceBatchAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) ForceBatchAddress() (common.Address, error) {
	return _Polygonzkevmexistentetrog.Contract.ForceBatchAddress(&_Polygonzkevmexistentetrog.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) ForceBatchAddress() (common.Address, error) {
	return _Polygonzkevmexistentetrog.Contract.ForceBatchAddress(&_Polygonzkevmexistentetrog.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonzkevmexistentetrog.Contract.ForceBatchTimeout(&_Polygonzkevmexistentetrog.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonzkevmexistentetrog.Contract.ForceBatchTimeout(&_Polygonzkevmexistentetrog.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonzkevmexistentetrog.Contract.ForcedBatches(&_Polygonzkevmexistentetrog.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonzkevmexistentetrog.Contract.ForcedBatches(&_Polygonzkevmexistentetrog.CallOpts, arg0)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) GasTokenAddress() (common.Address, error) {
	return _Polygonzkevmexistentetrog.Contract.GasTokenAddress(&_Polygonzkevmexistentetrog.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) GasTokenAddress() (common.Address, error) {
	return _Polygonzkevmexistentetrog.Contract.GasTokenAddress(&_Polygonzkevmexistentetrog.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) GasTokenNetwork() (uint32, error) {
	return _Polygonzkevmexistentetrog.Contract.GasTokenNetwork(&_Polygonzkevmexistentetrog.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) GasTokenNetwork() (uint32, error) {
	return _Polygonzkevmexistentetrog.Contract.GasTokenNetwork(&_Polygonzkevmexistentetrog.CallOpts)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) GenerateInitializeTransaction(opts *bind.CallOpts, networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "generateInitializeTransaction", networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) GenerateInitializeTransaction(networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _Polygonzkevmexistentetrog.Contract.GenerateInitializeTransaction(&_Polygonzkevmexistentetrog.CallOpts, networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) GenerateInitializeTransaction(networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _Polygonzkevmexistentetrog.Contract.GenerateInitializeTransaction(&_Polygonzkevmexistentetrog.CallOpts, networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonzkevmexistentetrog.Contract.GlobalExitRootManager(&_Polygonzkevmexistentetrog.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonzkevmexistentetrog.Contract.GlobalExitRootManager(&_Polygonzkevmexistentetrog.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) LastAccInputHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "lastAccInputHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) LastAccInputHash() ([32]byte, error) {
	return _Polygonzkevmexistentetrog.Contract.LastAccInputHash(&_Polygonzkevmexistentetrog.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) LastAccInputHash() ([32]byte, error) {
	return _Polygonzkevmexistentetrog.Contract.LastAccInputHash(&_Polygonzkevmexistentetrog.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) LastForceBatch() (uint64, error) {
	return _Polygonzkevmexistentetrog.Contract.LastForceBatch(&_Polygonzkevmexistentetrog.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) LastForceBatch() (uint64, error) {
	return _Polygonzkevmexistentetrog.Contract.LastForceBatch(&_Polygonzkevmexistentetrog.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonzkevmexistentetrog.Contract.LastForceBatchSequenced(&_Polygonzkevmexistentetrog.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonzkevmexistentetrog.Contract.LastForceBatchSequenced(&_Polygonzkevmexistentetrog.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) NetworkName() (string, error) {
	return _Polygonzkevmexistentetrog.Contract.NetworkName(&_Polygonzkevmexistentetrog.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) NetworkName() (string, error) {
	return _Polygonzkevmexistentetrog.Contract.NetworkName(&_Polygonzkevmexistentetrog.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) PendingAdmin() (common.Address, error) {
	return _Polygonzkevmexistentetrog.Contract.PendingAdmin(&_Polygonzkevmexistentetrog.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) PendingAdmin() (common.Address, error) {
	return _Polygonzkevmexistentetrog.Contract.PendingAdmin(&_Polygonzkevmexistentetrog.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) Pol() (common.Address, error) {
	return _Polygonzkevmexistentetrog.Contract.Pol(&_Polygonzkevmexistentetrog.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) Pol() (common.Address, error) {
	return _Polygonzkevmexistentetrog.Contract.Pol(&_Polygonzkevmexistentetrog.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) RollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "rollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) RollupManager() (common.Address, error) {
	return _Polygonzkevmexistentetrog.Contract.RollupManager(&_Polygonzkevmexistentetrog.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) RollupManager() (common.Address, error) {
	return _Polygonzkevmexistentetrog.Contract.RollupManager(&_Polygonzkevmexistentetrog.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) TrustedSequencer() (common.Address, error) {
	return _Polygonzkevmexistentetrog.Contract.TrustedSequencer(&_Polygonzkevmexistentetrog.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) TrustedSequencer() (common.Address, error) {
	return _Polygonzkevmexistentetrog.Contract.TrustedSequencer(&_Polygonzkevmexistentetrog.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCaller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonzkevmexistentetrog.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) TrustedSequencerURL() (string, error) {
	return _Polygonzkevmexistentetrog.Contract.TrustedSequencerURL(&_Polygonzkevmexistentetrog.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogCallerSession) TrustedSequencerURL() (string, error) {
	return _Polygonzkevmexistentetrog.Contract.TrustedSequencerURL(&_Polygonzkevmexistentetrog.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.AcceptAdminRole(&_Polygonzkevmexistentetrog.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.AcceptAdminRole(&_Polygonzkevmexistentetrog.TransactOpts)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogTransactor) ForceBatch(opts *bind.TransactOpts, transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.contract.Transact(opts, "forceBatch", transactions, polAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) ForceBatch(transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.ForceBatch(&_Polygonzkevmexistentetrog.TransactOpts, transactions, polAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogTransactorSession) ForceBatch(transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.ForceBatch(&_Polygonzkevmexistentetrog.TransactOpts, transactions, polAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.contract.Transact(opts, "initialize", _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.Initialize(&_Polygonzkevmexistentetrog.TransactOpts, _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogTransactorSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.Initialize(&_Polygonzkevmexistentetrog.TransactOpts, _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// InitializeUpgrade is a paid mutator transaction binding the contract method 0x5d6717a5.
//
// Solidity: function initializeUpgrade(address _admin, address _trustedSequencer, string _trustedSequencerURL, string _networkName, bytes32 _lastAccInputHash) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogTransactor) InitializeUpgrade(opts *bind.TransactOpts, _admin common.Address, _trustedSequencer common.Address, _trustedSequencerURL string, _networkName string, _lastAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.contract.Transact(opts, "initializeUpgrade", _admin, _trustedSequencer, _trustedSequencerURL, _networkName, _lastAccInputHash)
}

// InitializeUpgrade is a paid mutator transaction binding the contract method 0x5d6717a5.
//
// Solidity: function initializeUpgrade(address _admin, address _trustedSequencer, string _trustedSequencerURL, string _networkName, bytes32 _lastAccInputHash) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) InitializeUpgrade(_admin common.Address, _trustedSequencer common.Address, _trustedSequencerURL string, _networkName string, _lastAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.InitializeUpgrade(&_Polygonzkevmexistentetrog.TransactOpts, _admin, _trustedSequencer, _trustedSequencerURL, _networkName, _lastAccInputHash)
}

// InitializeUpgrade is a paid mutator transaction binding the contract method 0x5d6717a5.
//
// Solidity: function initializeUpgrade(address _admin, address _trustedSequencer, string _trustedSequencerURL, string _networkName, bytes32 _lastAccInputHash) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogTransactorSession) InitializeUpgrade(_admin common.Address, _trustedSequencer common.Address, _trustedSequencerURL string, _networkName string, _lastAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.InitializeUpgrade(&_Polygonzkevmexistentetrog.TransactOpts, _admin, _trustedSequencer, _trustedSequencerURL, _networkName, _lastAccInputHash)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogTransactor) OnVerifyBatches(opts *bind.TransactOpts, lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.contract.Transact(opts, "onVerifyBatches", lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.OnVerifyBatches(&_Polygonzkevmexistentetrog.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogTransactorSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.OnVerifyBatches(&_Polygonzkevmexistentetrog.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xdef57e54.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogTransactor) SequenceBatches(opts *bind.TransactOpts, batches []PolygonRollupBaseEtrogBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.contract.Transact(opts, "sequenceBatches", batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xdef57e54.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) SequenceBatches(batches []PolygonRollupBaseEtrogBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.SequenceBatches(&_Polygonzkevmexistentetrog.TransactOpts, batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xdef57e54.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogTransactorSession) SequenceBatches(batches []PolygonRollupBaseEtrogBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.SequenceBatches(&_Polygonzkevmexistentetrog.TransactOpts, batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogTransactor) SequenceForceBatches(opts *bind.TransactOpts, batches []PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.contract.Transact(opts, "sequenceForceBatches", batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) SequenceForceBatches(batches []PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.SequenceForceBatches(&_Polygonzkevmexistentetrog.TransactOpts, batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogTransactorSession) SequenceForceBatches(batches []PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.SequenceForceBatches(&_Polygonzkevmexistentetrog.TransactOpts, batches)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogTransactor) SetForceBatchAddress(opts *bind.TransactOpts, newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.contract.Transact(opts, "setForceBatchAddress", newForceBatchAddress)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) SetForceBatchAddress(newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.SetForceBatchAddress(&_Polygonzkevmexistentetrog.TransactOpts, newForceBatchAddress)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogTransactorSession) SetForceBatchAddress(newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.SetForceBatchAddress(&_Polygonzkevmexistentetrog.TransactOpts, newForceBatchAddress)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogTransactor) SetForceBatchTimeout(opts *bind.TransactOpts, newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.contract.Transact(opts, "setForceBatchTimeout", newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.SetForceBatchTimeout(&_Polygonzkevmexistentetrog.TransactOpts, newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogTransactorSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.SetForceBatchTimeout(&_Polygonzkevmexistentetrog.TransactOpts, newforceBatchTimeout)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogTransactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.SetTrustedSequencer(&_Polygonzkevmexistentetrog.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogTransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.SetTrustedSequencer(&_Polygonzkevmexistentetrog.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogTransactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.SetTrustedSequencerURL(&_Polygonzkevmexistentetrog.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogTransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.SetTrustedSequencerURL(&_Polygonzkevmexistentetrog.TransactOpts, newTrustedSequencerURL)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.TransferAdminRole(&_Polygonzkevmexistentetrog.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonzkevmexistentetrog.Contract.TransferAdminRole(&_Polygonzkevmexistentetrog.TransactOpts, newPendingAdmin)
}

// PolygonzkevmexistentetrogAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Polygonzkevmexistentetrog contract.
type PolygonzkevmexistentetrogAcceptAdminRoleIterator struct {
	Event *PolygonzkevmexistentetrogAcceptAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogAcceptAdminRole)
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
		it.Event = new(PolygonzkevmexistentetrogAcceptAdminRole)
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
func (it *PolygonzkevmexistentetrogAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogAcceptAdminRole represents a AcceptAdminRole event raised by the Polygonzkevmexistentetrog contract.
type PolygonzkevmexistentetrogAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*PolygonzkevmexistentetrogAcceptAdminRoleIterator, error) {

	logs, sub, err := _Polygonzkevmexistentetrog.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogAcceptAdminRoleIterator{contract: _Polygonzkevmexistentetrog.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmexistentetrog.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogAcceptAdminRole)
				if err := _Polygonzkevmexistentetrog.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
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
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) ParseAcceptAdminRole(log types.Log) (*PolygonzkevmexistentetrogAcceptAdminRole, error) {
	event := new(PolygonzkevmexistentetrogAcceptAdminRole)
	if err := _Polygonzkevmexistentetrog.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmexistentetrogForceBatchIterator is returned from FilterForceBatch and is used to iterate over the raw logs and unpacked data for ForceBatch events raised by the Polygonzkevmexistentetrog contract.
type PolygonzkevmexistentetrogForceBatchIterator struct {
	Event *PolygonzkevmexistentetrogForceBatch // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogForceBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogForceBatch)
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
		it.Event = new(PolygonzkevmexistentetrogForceBatch)
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
func (it *PolygonzkevmexistentetrogForceBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogForceBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogForceBatch represents a ForceBatch event raised by the Polygonzkevmexistentetrog contract.
type PolygonzkevmexistentetrogForceBatch struct {
	ForceBatchNum      uint64
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Transactions       []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterForceBatch is a free log retrieval operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) FilterForceBatch(opts *bind.FilterOpts, forceBatchNum []uint64) (*PolygonzkevmexistentetrogForceBatchIterator, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Polygonzkevmexistentetrog.contract.FilterLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogForceBatchIterator{contract: _Polygonzkevmexistentetrog.contract, event: "ForceBatch", logs: logs, sub: sub}, nil
}

// WatchForceBatch is a free log subscription operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) WatchForceBatch(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogForceBatch, forceBatchNum []uint64) (event.Subscription, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Polygonzkevmexistentetrog.contract.WatchLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogForceBatch)
				if err := _Polygonzkevmexistentetrog.contract.UnpackLog(event, "ForceBatch", log); err != nil {
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
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) ParseForceBatch(log types.Log) (*PolygonzkevmexistentetrogForceBatch, error) {
	event := new(PolygonzkevmexistentetrogForceBatch)
	if err := _Polygonzkevmexistentetrog.contract.UnpackLog(event, "ForceBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmexistentetrogInitialSequenceBatchesIterator is returned from FilterInitialSequenceBatches and is used to iterate over the raw logs and unpacked data for InitialSequenceBatches events raised by the Polygonzkevmexistentetrog contract.
type PolygonzkevmexistentetrogInitialSequenceBatchesIterator struct {
	Event *PolygonzkevmexistentetrogInitialSequenceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogInitialSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogInitialSequenceBatches)
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
		it.Event = new(PolygonzkevmexistentetrogInitialSequenceBatches)
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
func (it *PolygonzkevmexistentetrogInitialSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogInitialSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogInitialSequenceBatches represents a InitialSequenceBatches event raised by the Polygonzkevmexistentetrog contract.
type PolygonzkevmexistentetrogInitialSequenceBatches struct {
	Transactions       []byte
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterInitialSequenceBatches is a free log retrieval operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) FilterInitialSequenceBatches(opts *bind.FilterOpts) (*PolygonzkevmexistentetrogInitialSequenceBatchesIterator, error) {

	logs, sub, err := _Polygonzkevmexistentetrog.contract.FilterLogs(opts, "InitialSequenceBatches")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogInitialSequenceBatchesIterator{contract: _Polygonzkevmexistentetrog.contract, event: "InitialSequenceBatches", logs: logs, sub: sub}, nil
}

// WatchInitialSequenceBatches is a free log subscription operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) WatchInitialSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogInitialSequenceBatches) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmexistentetrog.contract.WatchLogs(opts, "InitialSequenceBatches")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogInitialSequenceBatches)
				if err := _Polygonzkevmexistentetrog.contract.UnpackLog(event, "InitialSequenceBatches", log); err != nil {
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
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) ParseInitialSequenceBatches(log types.Log) (*PolygonzkevmexistentetrogInitialSequenceBatches, error) {
	event := new(PolygonzkevmexistentetrogInitialSequenceBatches)
	if err := _Polygonzkevmexistentetrog.contract.UnpackLog(event, "InitialSequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmexistentetrogInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonzkevmexistentetrog contract.
type PolygonzkevmexistentetrogInitializedIterator struct {
	Event *PolygonzkevmexistentetrogInitialized // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogInitialized)
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
		it.Event = new(PolygonzkevmexistentetrogInitialized)
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
func (it *PolygonzkevmexistentetrogInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogInitialized represents a Initialized event raised by the Polygonzkevmexistentetrog contract.
type PolygonzkevmexistentetrogInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) FilterInitialized(opts *bind.FilterOpts) (*PolygonzkevmexistentetrogInitializedIterator, error) {

	logs, sub, err := _Polygonzkevmexistentetrog.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogInitializedIterator{contract: _Polygonzkevmexistentetrog.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogInitialized) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmexistentetrog.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogInitialized)
				if err := _Polygonzkevmexistentetrog.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) ParseInitialized(log types.Log) (*PolygonzkevmexistentetrogInitialized, error) {
	event := new(PolygonzkevmexistentetrogInitialized)
	if err := _Polygonzkevmexistentetrog.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmexistentetrogSequenceBatchesIterator is returned from FilterSequenceBatches and is used to iterate over the raw logs and unpacked data for SequenceBatches events raised by the Polygonzkevmexistentetrog contract.
type PolygonzkevmexistentetrogSequenceBatchesIterator struct {
	Event *PolygonzkevmexistentetrogSequenceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogSequenceBatches)
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
		it.Event = new(PolygonzkevmexistentetrogSequenceBatches)
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
func (it *PolygonzkevmexistentetrogSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogSequenceBatches represents a SequenceBatches event raised by the Polygonzkevmexistentetrog contract.
type PolygonzkevmexistentetrogSequenceBatches struct {
	NumBatch   uint64
	L1InfoRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSequenceBatches is a free log retrieval operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) FilterSequenceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonzkevmexistentetrogSequenceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevmexistentetrog.contract.FilterLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogSequenceBatchesIterator{contract: _Polygonzkevmexistentetrog.contract, event: "SequenceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceBatches is a free log subscription operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) WatchSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogSequenceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevmexistentetrog.contract.WatchLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogSequenceBatches)
				if err := _Polygonzkevmexistentetrog.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
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
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) ParseSequenceBatches(log types.Log) (*PolygonzkevmexistentetrogSequenceBatches, error) {
	event := new(PolygonzkevmexistentetrogSequenceBatches)
	if err := _Polygonzkevmexistentetrog.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmexistentetrogSequenceForceBatchesIterator is returned from FilterSequenceForceBatches and is used to iterate over the raw logs and unpacked data for SequenceForceBatches events raised by the Polygonzkevmexistentetrog contract.
type PolygonzkevmexistentetrogSequenceForceBatchesIterator struct {
	Event *PolygonzkevmexistentetrogSequenceForceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogSequenceForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogSequenceForceBatches)
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
		it.Event = new(PolygonzkevmexistentetrogSequenceForceBatches)
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
func (it *PolygonzkevmexistentetrogSequenceForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogSequenceForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogSequenceForceBatches represents a SequenceForceBatches event raised by the Polygonzkevmexistentetrog contract.
type PolygonzkevmexistentetrogSequenceForceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceForceBatches is a free log retrieval operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) FilterSequenceForceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonzkevmexistentetrogSequenceForceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevmexistentetrog.contract.FilterLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogSequenceForceBatchesIterator{contract: _Polygonzkevmexistentetrog.contract, event: "SequenceForceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceForceBatches is a free log subscription operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) WatchSequenceForceBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogSequenceForceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevmexistentetrog.contract.WatchLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogSequenceForceBatches)
				if err := _Polygonzkevmexistentetrog.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
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
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) ParseSequenceForceBatches(log types.Log) (*PolygonzkevmexistentetrogSequenceForceBatches, error) {
	event := new(PolygonzkevmexistentetrogSequenceForceBatches)
	if err := _Polygonzkevmexistentetrog.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmexistentetrogSetForceBatchAddressIterator is returned from FilterSetForceBatchAddress and is used to iterate over the raw logs and unpacked data for SetForceBatchAddress events raised by the Polygonzkevmexistentetrog contract.
type PolygonzkevmexistentetrogSetForceBatchAddressIterator struct {
	Event *PolygonzkevmexistentetrogSetForceBatchAddress // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogSetForceBatchAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogSetForceBatchAddress)
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
		it.Event = new(PolygonzkevmexistentetrogSetForceBatchAddress)
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
func (it *PolygonzkevmexistentetrogSetForceBatchAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogSetForceBatchAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogSetForceBatchAddress represents a SetForceBatchAddress event raised by the Polygonzkevmexistentetrog contract.
type PolygonzkevmexistentetrogSetForceBatchAddress struct {
	NewForceBatchAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchAddress is a free log retrieval operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) FilterSetForceBatchAddress(opts *bind.FilterOpts) (*PolygonzkevmexistentetrogSetForceBatchAddressIterator, error) {

	logs, sub, err := _Polygonzkevmexistentetrog.contract.FilterLogs(opts, "SetForceBatchAddress")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogSetForceBatchAddressIterator{contract: _Polygonzkevmexistentetrog.contract, event: "SetForceBatchAddress", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchAddress is a free log subscription operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) WatchSetForceBatchAddress(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogSetForceBatchAddress) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmexistentetrog.contract.WatchLogs(opts, "SetForceBatchAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogSetForceBatchAddress)
				if err := _Polygonzkevmexistentetrog.contract.UnpackLog(event, "SetForceBatchAddress", log); err != nil {
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
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) ParseSetForceBatchAddress(log types.Log) (*PolygonzkevmexistentetrogSetForceBatchAddress, error) {
	event := new(PolygonzkevmexistentetrogSetForceBatchAddress)
	if err := _Polygonzkevmexistentetrog.contract.UnpackLog(event, "SetForceBatchAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmexistentetrogSetForceBatchTimeoutIterator is returned from FilterSetForceBatchTimeout and is used to iterate over the raw logs and unpacked data for SetForceBatchTimeout events raised by the Polygonzkevmexistentetrog contract.
type PolygonzkevmexistentetrogSetForceBatchTimeoutIterator struct {
	Event *PolygonzkevmexistentetrogSetForceBatchTimeout // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogSetForceBatchTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogSetForceBatchTimeout)
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
		it.Event = new(PolygonzkevmexistentetrogSetForceBatchTimeout)
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
func (it *PolygonzkevmexistentetrogSetForceBatchTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogSetForceBatchTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogSetForceBatchTimeout represents a SetForceBatchTimeout event raised by the Polygonzkevmexistentetrog contract.
type PolygonzkevmexistentetrogSetForceBatchTimeout struct {
	NewforceBatchTimeout uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchTimeout is a free log retrieval operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) FilterSetForceBatchTimeout(opts *bind.FilterOpts) (*PolygonzkevmexistentetrogSetForceBatchTimeoutIterator, error) {

	logs, sub, err := _Polygonzkevmexistentetrog.contract.FilterLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogSetForceBatchTimeoutIterator{contract: _Polygonzkevmexistentetrog.contract, event: "SetForceBatchTimeout", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchTimeout is a free log subscription operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) WatchSetForceBatchTimeout(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogSetForceBatchTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmexistentetrog.contract.WatchLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogSetForceBatchTimeout)
				if err := _Polygonzkevmexistentetrog.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
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
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) ParseSetForceBatchTimeout(log types.Log) (*PolygonzkevmexistentetrogSetForceBatchTimeout, error) {
	event := new(PolygonzkevmexistentetrogSetForceBatchTimeout)
	if err := _Polygonzkevmexistentetrog.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmexistentetrogSetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the Polygonzkevmexistentetrog contract.
type PolygonzkevmexistentetrogSetTrustedSequencerIterator struct {
	Event *PolygonzkevmexistentetrogSetTrustedSequencer // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogSetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogSetTrustedSequencer)
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
		it.Event = new(PolygonzkevmexistentetrogSetTrustedSequencer)
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
func (it *PolygonzkevmexistentetrogSetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogSetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogSetTrustedSequencer represents a SetTrustedSequencer event raised by the Polygonzkevmexistentetrog contract.
type PolygonzkevmexistentetrogSetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*PolygonzkevmexistentetrogSetTrustedSequencerIterator, error) {

	logs, sub, err := _Polygonzkevmexistentetrog.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogSetTrustedSequencerIterator{contract: _Polygonzkevmexistentetrog.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogSetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmexistentetrog.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogSetTrustedSequencer)
				if err := _Polygonzkevmexistentetrog.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
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
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) ParseSetTrustedSequencer(log types.Log) (*PolygonzkevmexistentetrogSetTrustedSequencer, error) {
	event := new(PolygonzkevmexistentetrogSetTrustedSequencer)
	if err := _Polygonzkevmexistentetrog.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmexistentetrogSetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the Polygonzkevmexistentetrog contract.
type PolygonzkevmexistentetrogSetTrustedSequencerURLIterator struct {
	Event *PolygonzkevmexistentetrogSetTrustedSequencerURL // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogSetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogSetTrustedSequencerURL)
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
		it.Event = new(PolygonzkevmexistentetrogSetTrustedSequencerURL)
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
func (it *PolygonzkevmexistentetrogSetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogSetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogSetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the Polygonzkevmexistentetrog contract.
type PolygonzkevmexistentetrogSetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*PolygonzkevmexistentetrogSetTrustedSequencerURLIterator, error) {

	logs, sub, err := _Polygonzkevmexistentetrog.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogSetTrustedSequencerURLIterator{contract: _Polygonzkevmexistentetrog.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogSetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmexistentetrog.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogSetTrustedSequencerURL)
				if err := _Polygonzkevmexistentetrog.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
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
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) ParseSetTrustedSequencerURL(log types.Log) (*PolygonzkevmexistentetrogSetTrustedSequencerURL, error) {
	event := new(PolygonzkevmexistentetrogSetTrustedSequencerURL)
	if err := _Polygonzkevmexistentetrog.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmexistentetrogTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Polygonzkevmexistentetrog contract.
type PolygonzkevmexistentetrogTransferAdminRoleIterator struct {
	Event *PolygonzkevmexistentetrogTransferAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogTransferAdminRole)
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
		it.Event = new(PolygonzkevmexistentetrogTransferAdminRole)
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
func (it *PolygonzkevmexistentetrogTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogTransferAdminRole represents a TransferAdminRole event raised by the Polygonzkevmexistentetrog contract.
type PolygonzkevmexistentetrogTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*PolygonzkevmexistentetrogTransferAdminRoleIterator, error) {

	logs, sub, err := _Polygonzkevmexistentetrog.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogTransferAdminRoleIterator{contract: _Polygonzkevmexistentetrog.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmexistentetrog.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogTransferAdminRole)
				if err := _Polygonzkevmexistentetrog.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
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
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) ParseTransferAdminRole(log types.Log) (*PolygonzkevmexistentetrogTransferAdminRole, error) {
	event := new(PolygonzkevmexistentetrogTransferAdminRole)
	if err := _Polygonzkevmexistentetrog.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmexistentetrogUpdateEtrogSequenceIterator is returned from FilterUpdateEtrogSequence and is used to iterate over the raw logs and unpacked data for UpdateEtrogSequence events raised by the Polygonzkevmexistentetrog contract.
type PolygonzkevmexistentetrogUpdateEtrogSequenceIterator struct {
	Event *PolygonzkevmexistentetrogUpdateEtrogSequence // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogUpdateEtrogSequenceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogUpdateEtrogSequence)
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
		it.Event = new(PolygonzkevmexistentetrogUpdateEtrogSequence)
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
func (it *PolygonzkevmexistentetrogUpdateEtrogSequenceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogUpdateEtrogSequenceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogUpdateEtrogSequence represents a UpdateEtrogSequence event raised by the Polygonzkevmexistentetrog contract.
type PolygonzkevmexistentetrogUpdateEtrogSequence struct {
	NumBatch           uint64
	Transactions       []byte
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdateEtrogSequence is a free log retrieval operation binding the contract event 0xd2c80353fc15ef62c6affc7cd6b7ab5b42c43290c50be3372e55ae552cecd19c.
//
// Solidity: event UpdateEtrogSequence(uint64 numBatch, bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) FilterUpdateEtrogSequence(opts *bind.FilterOpts) (*PolygonzkevmexistentetrogUpdateEtrogSequenceIterator, error) {

	logs, sub, err := _Polygonzkevmexistentetrog.contract.FilterLogs(opts, "UpdateEtrogSequence")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogUpdateEtrogSequenceIterator{contract: _Polygonzkevmexistentetrog.contract, event: "UpdateEtrogSequence", logs: logs, sub: sub}, nil
}

// WatchUpdateEtrogSequence is a free log subscription operation binding the contract event 0xd2c80353fc15ef62c6affc7cd6b7ab5b42c43290c50be3372e55ae552cecd19c.
//
// Solidity: event UpdateEtrogSequence(uint64 numBatch, bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) WatchUpdateEtrogSequence(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogUpdateEtrogSequence) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmexistentetrog.contract.WatchLogs(opts, "UpdateEtrogSequence")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogUpdateEtrogSequence)
				if err := _Polygonzkevmexistentetrog.contract.UnpackLog(event, "UpdateEtrogSequence", log); err != nil {
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
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) ParseUpdateEtrogSequence(log types.Log) (*PolygonzkevmexistentetrogUpdateEtrogSequence, error) {
	event := new(PolygonzkevmexistentetrogUpdateEtrogSequence)
	if err := _Polygonzkevmexistentetrog.contract.UnpackLog(event, "UpdateEtrogSequence", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmexistentetrogVerifyBatchesIterator is returned from FilterVerifyBatches and is used to iterate over the raw logs and unpacked data for VerifyBatches events raised by the Polygonzkevmexistentetrog contract.
type PolygonzkevmexistentetrogVerifyBatchesIterator struct {
	Event *PolygonzkevmexistentetrogVerifyBatches // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmexistentetrogVerifyBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmexistentetrogVerifyBatches)
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
		it.Event = new(PolygonzkevmexistentetrogVerifyBatches)
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
func (it *PolygonzkevmexistentetrogVerifyBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmexistentetrogVerifyBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmexistentetrogVerifyBatches represents a VerifyBatches event raised by the Polygonzkevmexistentetrog contract.
type PolygonzkevmexistentetrogVerifyBatches struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatches is a free log retrieval operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) FilterVerifyBatches(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*PolygonzkevmexistentetrogVerifyBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonzkevmexistentetrog.contract.FilterLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmexistentetrogVerifyBatchesIterator{contract: _Polygonzkevmexistentetrog.contract, event: "VerifyBatches", logs: logs, sub: sub}, nil
}

// WatchVerifyBatches is a free log subscription operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) WatchVerifyBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmexistentetrogVerifyBatches, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonzkevmexistentetrog.contract.WatchLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmexistentetrogVerifyBatches)
				if err := _Polygonzkevmexistentetrog.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
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
func (_Polygonzkevmexistentetrog *PolygonzkevmexistentetrogFilterer) ParseVerifyBatches(log types.Log) (*PolygonzkevmexistentetrogVerifyBatches, error) {
	event := new(PolygonzkevmexistentetrogVerifyBatches)
	if err := _Polygonzkevmexistentetrog.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

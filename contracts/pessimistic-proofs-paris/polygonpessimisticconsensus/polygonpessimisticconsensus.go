// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonpessimisticconsensus

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

// PolygonpessimisticconsensusMetaData contains all meta data concerning the Polygonpessimisticconsensus contract.
var PolygonpessimisticconsensusMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"_rollupManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalAccInputHashDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CONSENSUS_TYPE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsensusHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b50604051620011cd380380620011cd83398101604081905262000035916200014b565b6001600160a01b0380851660a05280841660805280831660c052811660e052838383836200006262000070565b5050505050505050620001b3565b600054610100900460ff1615620000dd5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000130576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200014857600080fd5b50565b600080600080608085870312156200016257600080fd5b84516200016f8162000132565b6020860151909450620001828162000132565b6040860151909350620001958162000132565b6060860151909250620001a88162000132565b939692955090935050565b60805160a05160c05160e051610fd9620001f4600039600081816102bb015261063c015260006103510152600061044b015260006104720152610fd96000f3fe608060405234801561001057600080fd5b50600436106101985760003560e01c80638c3d7301116100e3578063cea5a4c01161008c578063e46761c411610066578063e46761c41461046d578063e7a7ed0214610494578063f851a440146104a857600080fd5b8063cea5a4c01461041e578063cfa8ed4714610426578063d02103ca1461044657600080fd5b8063ada8f919116100bd578063ada8f919146103d0578063c754c7ed146103e3578063c89e42df1461040b57600080fd5b80638c3d730114610344578063a3c573eb1461034c578063ad1edf341461037357600080fd5b806349b7b802116101455780636e05d2cd1161011f5780636e05d2cd146103135780636ff512cc1461031c578063712570221461033157600080fd5b806349b7b802146102b6578063542028d5146102dd5780636b8616ce146102e557600080fd5b80633c351e10116101765780633c351e10146102205780633cbc795b14610240578063456052671461027d57600080fd5b8063107bf28c1461019d57806326782247146101bb5780632c111c0614610200575b600080fd5b6101a56104ce565b6040516101b29190610b3f565b60405180910390f35b6001546101db9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101b2565b6008546101db9073ffffffffffffffffffffffffffffffffffffffff1681565b6009546101db9073ffffffffffffffffffffffffffffffffffffffff1681565b6009546102689074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff90911681526020016101b2565b60075461029d9068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020016101b2565b6101db7f000000000000000000000000000000000000000000000000000000000000000081565b6101a561055c565b6103056102f3366004610bab565b60066020526000908152604090205481565b6040519081526020016101b2565b61030560055481565b61032f61032a366004610c05565b610569565b005b61032f61033f366004610cfa565b61063a565b61032f610913565b6101db7f000000000000000000000000000000000000000000000000000000000000000081565b61030560025460405160006020820181905260609290921b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602482015260380160405160208183030381529060405280519060200120905090565b61032f6103de366004610c05565b6109e6565b60075461029d90700100000000000000000000000000000000900467ffffffffffffffff1681565b61032f610419366004610daa565b610ab0565b610268600081565b6002546101db9073ffffffffffffffffffffffffffffffffffffffff1681565b6101db7f000000000000000000000000000000000000000000000000000000000000000081565b6101db7f000000000000000000000000000000000000000000000000000000000000000081565b60075461029d9067ffffffffffffffff1681565b6000546101db9062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b600480546104db90610de7565b80601f016020809104026020016040519081016040528092919081815260200182805461050790610de7565b80156105545780601f1061052957610100808354040283529160200191610554565b820191906000526020600020905b81548152906001019060200180831161053757829003601f168201915b505050505081565b600380546104db90610de7565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff1633146105c0576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0906020015b60405180910390a150565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1633146106a9576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600054610100900460ff16158080156106c95750600054600160ff909116105b806106e35750303b1580156106e3575060005460ff166001145b610773576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840160405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156107d157600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b600080547fffffffffffffffffffff0000000000000000000000000000000000000000ffff166201000073ffffffffffffffffffffffffffffffffffffffff8a81169190910291909117909155600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001691881691909117905560036108598482610e89565b5060046108668382610e89565b50600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8616179055801561090a57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610964576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600154600080547fffffffffffffffffffff0000000000000000000000000000000000000000ffff1673ffffffffffffffffffffffffffffffffffffffff9092166201000081029290921790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e9060200160405180910390a1565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314610a3d576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce69060200161062f565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314610b07576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6003610b138282610e89565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b208160405161062f91905b600060208083528351808285015260005b81811015610b6c57858101830151858201604001528201610b50565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b600060208284031215610bbd57600080fd5b813567ffffffffffffffff81168114610bd557600080fd5b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610c0057600080fd5b919050565b600060208284031215610c1757600080fd5b610bd582610bdc565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f830112610c6057600080fd5b813567ffffffffffffffff80821115610c7b57610c7b610c20565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715610cc157610cc1610c20565b81604052838152866020858801011115610cda57600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060008060008060c08789031215610d1357600080fd5b610d1c87610bdc565b9550610d2a60208801610bdc565b9450604087013563ffffffff81168114610d4357600080fd5b9350610d5160608801610bdc565b9250608087013567ffffffffffffffff80821115610d6e57600080fd5b610d7a8a838b01610c4f565b935060a0890135915080821115610d9057600080fd5b50610d9d89828a01610c4f565b9150509295509295509295565b600060208284031215610dbc57600080fd5b813567ffffffffffffffff811115610dd357600080fd5b610ddf84828501610c4f565b949350505050565b600181811c90821680610dfb57607f821691505b602082108103610e34577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b601f821115610e8457600081815260208120601f850160051c81016020861015610e615750805b601f850160051c820191505b81811015610e8057828155600101610e6d565b5050505b505050565b815167ffffffffffffffff811115610ea357610ea3610c20565b610eb781610eb18454610de7565b84610e3a565b602080601f831160018114610f0a5760008415610ed45750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555610e80565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015610f5757888601518255948401946001909101908401610f38565b5085821015610f9357878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b0190555056fea2646970667358221220447ada8f38d3a295ec6d1ea0a3557384a568970e1dcae7c151708c49c5b84c5764736f6c63430008140033",
}

// PolygonpessimisticconsensusABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonpessimisticconsensusMetaData.ABI instead.
var PolygonpessimisticconsensusABI = PolygonpessimisticconsensusMetaData.ABI

// PolygonpessimisticconsensusBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygonpessimisticconsensusMetaData.Bin instead.
var PolygonpessimisticconsensusBin = PolygonpessimisticconsensusMetaData.Bin

// DeployPolygonpessimisticconsensus deploys a new Ethereum contract, binding an instance of Polygonpessimisticconsensus to it.
func DeployPolygonpessimisticconsensus(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address, _rollupManager common.Address) (common.Address, *types.Transaction, *Polygonpessimisticconsensus, error) {
	parsed, err := PolygonpessimisticconsensusMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygonpessimisticconsensusBin), backend, _globalExitRootManager, _pol, _bridgeAddress, _rollupManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonpessimisticconsensus{PolygonpessimisticconsensusCaller: PolygonpessimisticconsensusCaller{contract: contract}, PolygonpessimisticconsensusTransactor: PolygonpessimisticconsensusTransactor{contract: contract}, PolygonpessimisticconsensusFilterer: PolygonpessimisticconsensusFilterer{contract: contract}}, nil
}

// Polygonpessimisticconsensus is an auto generated Go binding around an Ethereum contract.
type Polygonpessimisticconsensus struct {
	PolygonpessimisticconsensusCaller     // Read-only binding to the contract
	PolygonpessimisticconsensusTransactor // Write-only binding to the contract
	PolygonpessimisticconsensusFilterer   // Log filterer for contract events
}

// PolygonpessimisticconsensusCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonpessimisticconsensusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonpessimisticconsensusTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonpessimisticconsensusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonpessimisticconsensusFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonpessimisticconsensusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonpessimisticconsensusSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonpessimisticconsensusSession struct {
	Contract     *Polygonpessimisticconsensus // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PolygonpessimisticconsensusCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonpessimisticconsensusCallerSession struct {
	Contract *PolygonpessimisticconsensusCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// PolygonpessimisticconsensusTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonpessimisticconsensusTransactorSession struct {
	Contract     *PolygonpessimisticconsensusTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// PolygonpessimisticconsensusRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonpessimisticconsensusRaw struct {
	Contract *Polygonpessimisticconsensus // Generic contract binding to access the raw methods on
}

// PolygonpessimisticconsensusCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonpessimisticconsensusCallerRaw struct {
	Contract *PolygonpessimisticconsensusCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonpessimisticconsensusTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonpessimisticconsensusTransactorRaw struct {
	Contract *PolygonpessimisticconsensusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonpessimisticconsensus creates a new instance of Polygonpessimisticconsensus, bound to a specific deployed contract.
func NewPolygonpessimisticconsensus(address common.Address, backend bind.ContractBackend) (*Polygonpessimisticconsensus, error) {
	contract, err := bindPolygonpessimisticconsensus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonpessimisticconsensus{PolygonpessimisticconsensusCaller: PolygonpessimisticconsensusCaller{contract: contract}, PolygonpessimisticconsensusTransactor: PolygonpessimisticconsensusTransactor{contract: contract}, PolygonpessimisticconsensusFilterer: PolygonpessimisticconsensusFilterer{contract: contract}}, nil
}

// NewPolygonpessimisticconsensusCaller creates a new read-only instance of Polygonpessimisticconsensus, bound to a specific deployed contract.
func NewPolygonpessimisticconsensusCaller(address common.Address, caller bind.ContractCaller) (*PolygonpessimisticconsensusCaller, error) {
	contract, err := bindPolygonpessimisticconsensus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonpessimisticconsensusCaller{contract: contract}, nil
}

// NewPolygonpessimisticconsensusTransactor creates a new write-only instance of Polygonpessimisticconsensus, bound to a specific deployed contract.
func NewPolygonpessimisticconsensusTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonpessimisticconsensusTransactor, error) {
	contract, err := bindPolygonpessimisticconsensus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonpessimisticconsensusTransactor{contract: contract}, nil
}

// NewPolygonpessimisticconsensusFilterer creates a new log filterer instance of Polygonpessimisticconsensus, bound to a specific deployed contract.
func NewPolygonpessimisticconsensusFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonpessimisticconsensusFilterer, error) {
	contract, err := bindPolygonpessimisticconsensus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonpessimisticconsensusFilterer{contract: contract}, nil
}

// bindPolygonpessimisticconsensus binds a generic wrapper to an already deployed contract.
func bindPolygonpessimisticconsensus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonpessimisticconsensusMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonpessimisticconsensus.Contract.PolygonpessimisticconsensusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonpessimisticconsensus.Contract.PolygonpessimisticconsensusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonpessimisticconsensus.Contract.PolygonpessimisticconsensusTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonpessimisticconsensus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonpessimisticconsensus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonpessimisticconsensus.Contract.contract.Transact(opts, method, params...)
}

// CONSENSUSTYPE is a free data retrieval call binding the contract method 0xcea5a4c0.
//
// Solidity: function CONSENSUS_TYPE() view returns(uint32)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCaller) CONSENSUSTYPE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonpessimisticconsensus.contract.Call(opts, &out, "CONSENSUS_TYPE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CONSENSUSTYPE is a free data retrieval call binding the contract method 0xcea5a4c0.
//
// Solidity: function CONSENSUS_TYPE() view returns(uint32)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusSession) CONSENSUSTYPE() (uint32, error) {
	return _Polygonpessimisticconsensus.Contract.CONSENSUSTYPE(&_Polygonpessimisticconsensus.CallOpts)
}

// CONSENSUSTYPE is a free data retrieval call binding the contract method 0xcea5a4c0.
//
// Solidity: function CONSENSUS_TYPE() view returns(uint32)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCallerSession) CONSENSUSTYPE() (uint32, error) {
	return _Polygonpessimisticconsensus.Contract.CONSENSUSTYPE(&_Polygonpessimisticconsensus.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonpessimisticconsensus.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusSession) Admin() (common.Address, error) {
	return _Polygonpessimisticconsensus.Contract.Admin(&_Polygonpessimisticconsensus.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCallerSession) Admin() (common.Address, error) {
	return _Polygonpessimisticconsensus.Contract.Admin(&_Polygonpessimisticconsensus.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonpessimisticconsensus.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusSession) BridgeAddress() (common.Address, error) {
	return _Polygonpessimisticconsensus.Contract.BridgeAddress(&_Polygonpessimisticconsensus.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCallerSession) BridgeAddress() (common.Address, error) {
	return _Polygonpessimisticconsensus.Contract.BridgeAddress(&_Polygonpessimisticconsensus.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCaller) ForceBatchAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonpessimisticconsensus.contract.Call(opts, &out, "forceBatchAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusSession) ForceBatchAddress() (common.Address, error) {
	return _Polygonpessimisticconsensus.Contract.ForceBatchAddress(&_Polygonpessimisticconsensus.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCallerSession) ForceBatchAddress() (common.Address, error) {
	return _Polygonpessimisticconsensus.Contract.ForceBatchAddress(&_Polygonpessimisticconsensus.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCaller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonpessimisticconsensus.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonpessimisticconsensus.Contract.ForceBatchTimeout(&_Polygonpessimisticconsensus.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCallerSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonpessimisticconsensus.Contract.ForceBatchTimeout(&_Polygonpessimisticconsensus.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCaller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Polygonpessimisticconsensus.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonpessimisticconsensus.Contract.ForcedBatches(&_Polygonpessimisticconsensus.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonpessimisticconsensus.Contract.ForcedBatches(&_Polygonpessimisticconsensus.CallOpts, arg0)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCaller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonpessimisticconsensus.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusSession) GasTokenAddress() (common.Address, error) {
	return _Polygonpessimisticconsensus.Contract.GasTokenAddress(&_Polygonpessimisticconsensus.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCallerSession) GasTokenAddress() (common.Address, error) {
	return _Polygonpessimisticconsensus.Contract.GasTokenAddress(&_Polygonpessimisticconsensus.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCaller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonpessimisticconsensus.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusSession) GasTokenNetwork() (uint32, error) {
	return _Polygonpessimisticconsensus.Contract.GasTokenNetwork(&_Polygonpessimisticconsensus.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCallerSession) GasTokenNetwork() (uint32, error) {
	return _Polygonpessimisticconsensus.Contract.GasTokenNetwork(&_Polygonpessimisticconsensus.CallOpts)
}

// GetConsensusHash is a free data retrieval call binding the contract method 0xad1edf34.
//
// Solidity: function getConsensusHash() view returns(bytes32)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCaller) GetConsensusHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonpessimisticconsensus.contract.Call(opts, &out, "getConsensusHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetConsensusHash is a free data retrieval call binding the contract method 0xad1edf34.
//
// Solidity: function getConsensusHash() view returns(bytes32)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusSession) GetConsensusHash() ([32]byte, error) {
	return _Polygonpessimisticconsensus.Contract.GetConsensusHash(&_Polygonpessimisticconsensus.CallOpts)
}

// GetConsensusHash is a free data retrieval call binding the contract method 0xad1edf34.
//
// Solidity: function getConsensusHash() view returns(bytes32)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCallerSession) GetConsensusHash() ([32]byte, error) {
	return _Polygonpessimisticconsensus.Contract.GetConsensusHash(&_Polygonpessimisticconsensus.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonpessimisticconsensus.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonpessimisticconsensus.Contract.GlobalExitRootManager(&_Polygonpessimisticconsensus.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonpessimisticconsensus.Contract.GlobalExitRootManager(&_Polygonpessimisticconsensus.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCaller) LastAccInputHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonpessimisticconsensus.contract.Call(opts, &out, "lastAccInputHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusSession) LastAccInputHash() ([32]byte, error) {
	return _Polygonpessimisticconsensus.Contract.LastAccInputHash(&_Polygonpessimisticconsensus.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCallerSession) LastAccInputHash() ([32]byte, error) {
	return _Polygonpessimisticconsensus.Contract.LastAccInputHash(&_Polygonpessimisticconsensus.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCaller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonpessimisticconsensus.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusSession) LastForceBatch() (uint64, error) {
	return _Polygonpessimisticconsensus.Contract.LastForceBatch(&_Polygonpessimisticconsensus.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCallerSession) LastForceBatch() (uint64, error) {
	return _Polygonpessimisticconsensus.Contract.LastForceBatch(&_Polygonpessimisticconsensus.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCaller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonpessimisticconsensus.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonpessimisticconsensus.Contract.LastForceBatchSequenced(&_Polygonpessimisticconsensus.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCallerSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonpessimisticconsensus.Contract.LastForceBatchSequenced(&_Polygonpessimisticconsensus.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCaller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonpessimisticconsensus.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusSession) NetworkName() (string, error) {
	return _Polygonpessimisticconsensus.Contract.NetworkName(&_Polygonpessimisticconsensus.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCallerSession) NetworkName() (string, error) {
	return _Polygonpessimisticconsensus.Contract.NetworkName(&_Polygonpessimisticconsensus.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonpessimisticconsensus.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusSession) PendingAdmin() (common.Address, error) {
	return _Polygonpessimisticconsensus.Contract.PendingAdmin(&_Polygonpessimisticconsensus.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCallerSession) PendingAdmin() (common.Address, error) {
	return _Polygonpessimisticconsensus.Contract.PendingAdmin(&_Polygonpessimisticconsensus.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonpessimisticconsensus.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusSession) Pol() (common.Address, error) {
	return _Polygonpessimisticconsensus.Contract.Pol(&_Polygonpessimisticconsensus.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCallerSession) Pol() (common.Address, error) {
	return _Polygonpessimisticconsensus.Contract.Pol(&_Polygonpessimisticconsensus.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCaller) RollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonpessimisticconsensus.contract.Call(opts, &out, "rollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusSession) RollupManager() (common.Address, error) {
	return _Polygonpessimisticconsensus.Contract.RollupManager(&_Polygonpessimisticconsensus.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCallerSession) RollupManager() (common.Address, error) {
	return _Polygonpessimisticconsensus.Contract.RollupManager(&_Polygonpessimisticconsensus.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCaller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonpessimisticconsensus.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusSession) TrustedSequencer() (common.Address, error) {
	return _Polygonpessimisticconsensus.Contract.TrustedSequencer(&_Polygonpessimisticconsensus.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCallerSession) TrustedSequencer() (common.Address, error) {
	return _Polygonpessimisticconsensus.Contract.TrustedSequencer(&_Polygonpessimisticconsensus.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCaller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonpessimisticconsensus.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusSession) TrustedSequencerURL() (string, error) {
	return _Polygonpessimisticconsensus.Contract.TrustedSequencerURL(&_Polygonpessimisticconsensus.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusCallerSession) TrustedSequencerURL() (string, error) {
	return _Polygonpessimisticconsensus.Contract.TrustedSequencerURL(&_Polygonpessimisticconsensus.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonpessimisticconsensus.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonpessimisticconsensus.Contract.AcceptAdminRole(&_Polygonpessimisticconsensus.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonpessimisticconsensus.Contract.AcceptAdminRole(&_Polygonpessimisticconsensus.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 , address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, sequencer common.Address, arg2 uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonpessimisticconsensus.contract.Transact(opts, "initialize", _admin, sequencer, arg2, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 , address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusSession) Initialize(_admin common.Address, sequencer common.Address, arg2 uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonpessimisticconsensus.Contract.Initialize(&_Polygonpessimisticconsensus.TransactOpts, _admin, sequencer, arg2, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 , address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusTransactorSession) Initialize(_admin common.Address, sequencer common.Address, arg2 uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonpessimisticconsensus.Contract.Initialize(&_Polygonpessimisticconsensus.TransactOpts, _admin, sequencer, arg2, _gasTokenAddress, sequencerURL, _networkName)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusTransactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonpessimisticconsensus.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonpessimisticconsensus.Contract.SetTrustedSequencer(&_Polygonpessimisticconsensus.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusTransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonpessimisticconsensus.Contract.SetTrustedSequencer(&_Polygonpessimisticconsensus.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusTransactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonpessimisticconsensus.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonpessimisticconsensus.Contract.SetTrustedSequencerURL(&_Polygonpessimisticconsensus.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusTransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonpessimisticconsensus.Contract.SetTrustedSequencerURL(&_Polygonpessimisticconsensus.TransactOpts, newTrustedSequencerURL)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonpessimisticconsensus.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonpessimisticconsensus.Contract.TransferAdminRole(&_Polygonpessimisticconsensus.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonpessimisticconsensus.Contract.TransferAdminRole(&_Polygonpessimisticconsensus.TransactOpts, newPendingAdmin)
}

// PolygonpessimisticconsensusAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Polygonpessimisticconsensus contract.
type PolygonpessimisticconsensusAcceptAdminRoleIterator struct {
	Event *PolygonpessimisticconsensusAcceptAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonpessimisticconsensusAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonpessimisticconsensusAcceptAdminRole)
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
		it.Event = new(PolygonpessimisticconsensusAcceptAdminRole)
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
func (it *PolygonpessimisticconsensusAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonpessimisticconsensusAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonpessimisticconsensusAcceptAdminRole represents a AcceptAdminRole event raised by the Polygonpessimisticconsensus contract.
type PolygonpessimisticconsensusAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*PolygonpessimisticconsensusAcceptAdminRoleIterator, error) {

	logs, sub, err := _Polygonpessimisticconsensus.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonpessimisticconsensusAcceptAdminRoleIterator{contract: _Polygonpessimisticconsensus.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonpessimisticconsensusAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonpessimisticconsensus.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonpessimisticconsensusAcceptAdminRole)
				if err := _Polygonpessimisticconsensus.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
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
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusFilterer) ParseAcceptAdminRole(log types.Log) (*PolygonpessimisticconsensusAcceptAdminRole, error) {
	event := new(PolygonpessimisticconsensusAcceptAdminRole)
	if err := _Polygonpessimisticconsensus.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonpessimisticconsensusInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonpessimisticconsensus contract.
type PolygonpessimisticconsensusInitializedIterator struct {
	Event *PolygonpessimisticconsensusInitialized // Event containing the contract specifics and raw log

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
func (it *PolygonpessimisticconsensusInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonpessimisticconsensusInitialized)
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
		it.Event = new(PolygonpessimisticconsensusInitialized)
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
func (it *PolygonpessimisticconsensusInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonpessimisticconsensusInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonpessimisticconsensusInitialized represents a Initialized event raised by the Polygonpessimisticconsensus contract.
type PolygonpessimisticconsensusInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusFilterer) FilterInitialized(opts *bind.FilterOpts) (*PolygonpessimisticconsensusInitializedIterator, error) {

	logs, sub, err := _Polygonpessimisticconsensus.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PolygonpessimisticconsensusInitializedIterator{contract: _Polygonpessimisticconsensus.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PolygonpessimisticconsensusInitialized) (event.Subscription, error) {

	logs, sub, err := _Polygonpessimisticconsensus.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonpessimisticconsensusInitialized)
				if err := _Polygonpessimisticconsensus.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusFilterer) ParseInitialized(log types.Log) (*PolygonpessimisticconsensusInitialized, error) {
	event := new(PolygonpessimisticconsensusInitialized)
	if err := _Polygonpessimisticconsensus.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonpessimisticconsensusSetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the Polygonpessimisticconsensus contract.
type PolygonpessimisticconsensusSetTrustedSequencerIterator struct {
	Event *PolygonpessimisticconsensusSetTrustedSequencer // Event containing the contract specifics and raw log

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
func (it *PolygonpessimisticconsensusSetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonpessimisticconsensusSetTrustedSequencer)
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
		it.Event = new(PolygonpessimisticconsensusSetTrustedSequencer)
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
func (it *PolygonpessimisticconsensusSetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonpessimisticconsensusSetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonpessimisticconsensusSetTrustedSequencer represents a SetTrustedSequencer event raised by the Polygonpessimisticconsensus contract.
type PolygonpessimisticconsensusSetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusFilterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*PolygonpessimisticconsensusSetTrustedSequencerIterator, error) {

	logs, sub, err := _Polygonpessimisticconsensus.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &PolygonpessimisticconsensusSetTrustedSequencerIterator{contract: _Polygonpessimisticconsensus.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusFilterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *PolygonpessimisticconsensusSetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _Polygonpessimisticconsensus.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonpessimisticconsensusSetTrustedSequencer)
				if err := _Polygonpessimisticconsensus.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
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
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusFilterer) ParseSetTrustedSequencer(log types.Log) (*PolygonpessimisticconsensusSetTrustedSequencer, error) {
	event := new(PolygonpessimisticconsensusSetTrustedSequencer)
	if err := _Polygonpessimisticconsensus.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonpessimisticconsensusSetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the Polygonpessimisticconsensus contract.
type PolygonpessimisticconsensusSetTrustedSequencerURLIterator struct {
	Event *PolygonpessimisticconsensusSetTrustedSequencerURL // Event containing the contract specifics and raw log

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
func (it *PolygonpessimisticconsensusSetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonpessimisticconsensusSetTrustedSequencerURL)
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
		it.Event = new(PolygonpessimisticconsensusSetTrustedSequencerURL)
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
func (it *PolygonpessimisticconsensusSetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonpessimisticconsensusSetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonpessimisticconsensusSetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the Polygonpessimisticconsensus contract.
type PolygonpessimisticconsensusSetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusFilterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*PolygonpessimisticconsensusSetTrustedSequencerURLIterator, error) {

	logs, sub, err := _Polygonpessimisticconsensus.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &PolygonpessimisticconsensusSetTrustedSequencerURLIterator{contract: _Polygonpessimisticconsensus.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusFilterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *PolygonpessimisticconsensusSetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _Polygonpessimisticconsensus.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonpessimisticconsensusSetTrustedSequencerURL)
				if err := _Polygonpessimisticconsensus.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
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
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusFilterer) ParseSetTrustedSequencerURL(log types.Log) (*PolygonpessimisticconsensusSetTrustedSequencerURL, error) {
	event := new(PolygonpessimisticconsensusSetTrustedSequencerURL)
	if err := _Polygonpessimisticconsensus.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonpessimisticconsensusTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Polygonpessimisticconsensus contract.
type PolygonpessimisticconsensusTransferAdminRoleIterator struct {
	Event *PolygonpessimisticconsensusTransferAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonpessimisticconsensusTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonpessimisticconsensusTransferAdminRole)
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
		it.Event = new(PolygonpessimisticconsensusTransferAdminRole)
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
func (it *PolygonpessimisticconsensusTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonpessimisticconsensusTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonpessimisticconsensusTransferAdminRole represents a TransferAdminRole event raised by the Polygonpessimisticconsensus contract.
type PolygonpessimisticconsensusTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*PolygonpessimisticconsensusTransferAdminRoleIterator, error) {

	logs, sub, err := _Polygonpessimisticconsensus.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonpessimisticconsensusTransferAdminRoleIterator{contract: _Polygonpessimisticconsensus.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonpessimisticconsensusTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonpessimisticconsensus.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonpessimisticconsensusTransferAdminRole)
				if err := _Polygonpessimisticconsensus.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
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
func (_Polygonpessimisticconsensus *PolygonpessimisticconsensusFilterer) ParseTransferAdminRole(log types.Log) (*PolygonpessimisticconsensusTransferAdminRole, error) {
	event := new(PolygonpessimisticconsensusTransferAdminRole)
	if err := _Polygonpessimisticconsensus.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonzkevmglobalexitrootv2mock

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

// Polygonzkevmglobalexitrootv2mockMetaData contains all meta data concerning the Polygonzkevmglobalexitrootv2mock contract.
var Polygonzkevmglobalexitrootv2mockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollupManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"GlobalExitRootAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughGlobalExitRootsInserted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotLastInsertedGlobalExitRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootRemover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootUpdater\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"leafCount\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"currentL1InfoRoot\",\"type\":\"bytes32\"}],\"name\":\"InitL1InfoRootMap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"}],\"name\":\"UpdateL1InfoTree\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"currentL1InfoRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"leafCount\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockhash\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"minTimestamp\",\"type\":\"uint64\"}],\"name\":\"UpdateL1InfoTreeV2\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"calculateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastGlobalExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"lastBlockHash\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"getLeafValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"depositCount\",\"type\":\"uint32\"}],\"name\":\"injectGER\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"leafCount\",\"type\":\"uint32\"}],\"name\":\"l1InfoRootMap\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"l1InfoRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastMainnetExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"verifyMerkleProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561000f575f80fd5b50604051610f09380380610f0983398101604081905261002e9161012f565b6001600160a01b0380831660a0528116608052818161004b610054565b50505050610160565b602e54610100900460ff16156100c05760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b602e5460ff908116101561011257602e805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b80516001600160a01b038116811461012a575f80fd5b919050565b5f8060408385031215610140575f80fd5b61014983610114565b915061015760208401610114565b90509250929050565b60805160a051610d7a61018f5f395f8181610162015261033501525f818161023901526102e90152610d7a5ff3fe608060405234801561000f575f80fd5b50600436106100f0575f3560e01c80635ca1e16511610093578063a3c573eb11610063578063a3c573eb14610234578063a43ca6a61461025b578063ef4eeb351461028e578063fb570834146102ad575f80fd5b80635ca1e165146101a95780635d810501146101b15780638129fc1c1461021957806383f2440314610221575f80fd5b8063319cf735116100ce578063319cf7351461013757806333d6247d146101405780633ed691ef1461015557806349b7b8021461015d575f80fd5b806301fd9044146100f4578063257b36321461010f5780632dfdf0b51461012e575b5f80fd5b6100fc5f5481565b6040519081526020015b60405180910390f35b6100fc61011d3660046109fd565b60026020525f908152604090205481565b6100fc60235481565b6100fc60015481565b61015361014e3660046109fd565b6102d0565b005b6100fc6104f6565b6101847f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610106565b6100fc610509565b6100fc6101bf366004610a14565b604080516020808201959095528082019390935260c09190911b7fffffffffffffffff0000000000000000000000000000000000000000000000001660608301528051604881840301815260689092019052805191012090565b610153610512565b6100fc61022f366004610a7f565b610706565b6101847f000000000000000000000000000000000000000000000000000000000000000081565b610153610269366004610abb565b5f82815260026020908152604080832042905563ffffffff9093168252602f90522055565b6100fc61029c366004610ae5565b602f6020525f908152604090205481565b6102c06102bb366004610b05565b6107db565b6040519015158152602001610106565b5f8073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016330361031e57505060018190555f548161039d565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016330361036b5750505f819055600154819061039d565b6040517fb49365dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6103a882846107f2565b5f81815260026020526040812054919250036104f057425f6103cb600143610b77565b5f8481526002602090815260409182902092409283905581518082018790528083018490527fffffffffffffffff00000000000000000000000000000000000000000000000060c087901b166060820152825180820360480181526068909101909252815191012090915061043f90610821565b5f610448610509565b60235463ffffffff165f908152602f602052604080822083905551919250879187917fda61aa7823fcd807e37b95aabcbe17f03a6f3efd514176444dae191d27fd66b391a360235463ffffffff167faf6c6cd7790e0180a4d22eb8ed846e55846f54ed10e5946db19972b5a0813a598284866040516104e493929190928352602083019190915267ffffffffffffffff16604082015260600190565b60405180910390a25050505b50505050565b5f6105046001545f546107f2565b905090565b5f610504610921565b602e54610100900460ff16158080156105325750602e54600160ff909116105b8061054c5750303b15801561054c5750602e5460ff166001145b6105dc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840160405180910390fd5b602e80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561063a57602e80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b5f610643610509565b6023805463ffffffff9081165f908152602f602090815260409182902085905592548151921682529181018390529192507f11f50c71891002839c2637ce302087160298255a87f1ea60d40e8db081383fad910160405180910390a150801561070357602e80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b5f83815b60208110156107d257600163ffffffff8516821c811690036107755784816020811061073857610738610b8a565b602002013582604051602001610758929190918252602082015260400190565b6040516020818303038152906040528051906020012091506107c0565b8185826020811061078857610788610b8a565b60200201356040516020016107a7929190918252602082015260400190565b6040516020818303038152906040528051906020012091505b806107ca81610bb7565b91505061070a565b50949350505050565b5f816107e8868686610706565b1495945050505050565b604080516020808201859052818301849052825180830384018152606090920190925280519101205b92915050565b80600161083060206002610d0c565b61083a9190610b77565b60235410610874576040517fef5ccf6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60235f815461088390610bb7565b918290555090505f5b6020811015610913578082901c6001166001036108bf5782600382602081106108b7576108b7610b8a565b015550505050565b600381602081106108d2576108d2610b8a565b01546040805160208101929092528101849052606001604051602081830303815290604052805190602001209250808061090b90610bb7565b91505061088c565b5061091c610d17565b505050565b6023545f90819081805b60208110156109f4578083901c600116600103610988576003816020811061095557610955610b8a565b015460408051602081019290925281018590526060016040516020818303038152906040528051906020012093506109b5565b60408051602081018690529081018390526060016040516020818303038152906040528051906020012093505b604080516020810184905290810183905260600160405160208183030381529060405280519060200120915080806109ec90610bb7565b91505061092b565b50919392505050565b5f60208284031215610a0d575f80fd5b5035919050565b5f805f60608486031215610a26575f80fd5b8335925060208401359150604084013567ffffffffffffffff81168114610a4b575f80fd5b809150509250925092565b80610400810183101561081b575f80fd5b803563ffffffff81168114610a7a575f80fd5b919050565b5f805f6104408486031215610a92575f80fd5b83359250610aa38560208601610a56565b9150610ab26104208501610a67565b90509250925092565b5f8060408385031215610acc575f80fd5b82359150610adc60208401610a67565b90509250929050565b5f60208284031215610af5575f80fd5b610afe82610a67565b9392505050565b5f805f806104608587031215610b19575f80fd5b84359350610b2a8660208701610a56565b9250610b396104208601610a67565b939692955092936104400135925050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8181038181111561081b5761081b610b4a565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610be757610be7610b4a565b5060010190565b600181815b80851115610c4757817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610c2d57610c2d610b4a565b80851615610c3a57918102915b93841c9390800290610bf3565b509250929050565b5f82610c5d5750600161081b565b81610c6957505f61081b565b8160018114610c7f5760028114610c8957610ca5565b600191505061081b565b60ff841115610c9a57610c9a610b4a565b50506001821b61081b565b5060208310610133831016604e8410600b8410161715610cc8575081810a61081b565b610cd28383610bee565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610d0457610d04610b4a565b029392505050565b5f610afe8383610c4f565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52600160045260245ffdfea26469706673582212207745abd6881f35b929d4fe9acf88eb1d05ef205354a666c0d098996526a9179e64736f6c63430008140033",
}

// Polygonzkevmglobalexitrootv2mockABI is the input ABI used to generate the binding from.
// Deprecated: Use Polygonzkevmglobalexitrootv2mockMetaData.ABI instead.
var Polygonzkevmglobalexitrootv2mockABI = Polygonzkevmglobalexitrootv2mockMetaData.ABI

// Polygonzkevmglobalexitrootv2mockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Polygonzkevmglobalexitrootv2mockMetaData.Bin instead.
var Polygonzkevmglobalexitrootv2mockBin = Polygonzkevmglobalexitrootv2mockMetaData.Bin

// DeployPolygonzkevmglobalexitrootv2mock deploys a new Ethereum contract, binding an instance of Polygonzkevmglobalexitrootv2mock to it.
func DeployPolygonzkevmglobalexitrootv2mock(auth *bind.TransactOpts, backend bind.ContractBackend, _rollupManager common.Address, _bridgeAddress common.Address) (common.Address, *types.Transaction, *Polygonzkevmglobalexitrootv2mock, error) {
	parsed, err := Polygonzkevmglobalexitrootv2mockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Polygonzkevmglobalexitrootv2mockBin), backend, _rollupManager, _bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonzkevmglobalexitrootv2mock{Polygonzkevmglobalexitrootv2mockCaller: Polygonzkevmglobalexitrootv2mockCaller{contract: contract}, Polygonzkevmglobalexitrootv2mockTransactor: Polygonzkevmglobalexitrootv2mockTransactor{contract: contract}, Polygonzkevmglobalexitrootv2mockFilterer: Polygonzkevmglobalexitrootv2mockFilterer{contract: contract}}, nil
}

// Polygonzkevmglobalexitrootv2mock is an auto generated Go binding around an Ethereum contract.
type Polygonzkevmglobalexitrootv2mock struct {
	Polygonzkevmglobalexitrootv2mockCaller     // Read-only binding to the contract
	Polygonzkevmglobalexitrootv2mockTransactor // Write-only binding to the contract
	Polygonzkevmglobalexitrootv2mockFilterer   // Log filterer for contract events
}

// Polygonzkevmglobalexitrootv2mockCaller is an auto generated read-only Go binding around an Ethereum contract.
type Polygonzkevmglobalexitrootv2mockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Polygonzkevmglobalexitrootv2mockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Polygonzkevmglobalexitrootv2mockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Polygonzkevmglobalexitrootv2mockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Polygonzkevmglobalexitrootv2mockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Polygonzkevmglobalexitrootv2mockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Polygonzkevmglobalexitrootv2mockSession struct {
	Contract     *Polygonzkevmglobalexitrootv2mock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                     // Call options to use throughout this session
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// Polygonzkevmglobalexitrootv2mockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Polygonzkevmglobalexitrootv2mockCallerSession struct {
	Contract *Polygonzkevmglobalexitrootv2mockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                           // Call options to use throughout this session
}

// Polygonzkevmglobalexitrootv2mockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Polygonzkevmglobalexitrootv2mockTransactorSession struct {
	Contract     *Polygonzkevmglobalexitrootv2mockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                           // Transaction auth options to use throughout this session
}

// Polygonzkevmglobalexitrootv2mockRaw is an auto generated low-level Go binding around an Ethereum contract.
type Polygonzkevmglobalexitrootv2mockRaw struct {
	Contract *Polygonzkevmglobalexitrootv2mock // Generic contract binding to access the raw methods on
}

// Polygonzkevmglobalexitrootv2mockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Polygonzkevmglobalexitrootv2mockCallerRaw struct {
	Contract *Polygonzkevmglobalexitrootv2mockCaller // Generic read-only contract binding to access the raw methods on
}

// Polygonzkevmglobalexitrootv2mockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Polygonzkevmglobalexitrootv2mockTransactorRaw struct {
	Contract *Polygonzkevmglobalexitrootv2mockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonzkevmglobalexitrootv2mock creates a new instance of Polygonzkevmglobalexitrootv2mock, bound to a specific deployed contract.
func NewPolygonzkevmglobalexitrootv2mock(address common.Address, backend bind.ContractBackend) (*Polygonzkevmglobalexitrootv2mock, error) {
	contract, err := bindPolygonzkevmglobalexitrootv2mock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootv2mock{Polygonzkevmglobalexitrootv2mockCaller: Polygonzkevmglobalexitrootv2mockCaller{contract: contract}, Polygonzkevmglobalexitrootv2mockTransactor: Polygonzkevmglobalexitrootv2mockTransactor{contract: contract}, Polygonzkevmglobalexitrootv2mockFilterer: Polygonzkevmglobalexitrootv2mockFilterer{contract: contract}}, nil
}

// NewPolygonzkevmglobalexitrootv2mockCaller creates a new read-only instance of Polygonzkevmglobalexitrootv2mock, bound to a specific deployed contract.
func NewPolygonzkevmglobalexitrootv2mockCaller(address common.Address, caller bind.ContractCaller) (*Polygonzkevmglobalexitrootv2mockCaller, error) {
	contract, err := bindPolygonzkevmglobalexitrootv2mock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootv2mockCaller{contract: contract}, nil
}

// NewPolygonzkevmglobalexitrootv2mockTransactor creates a new write-only instance of Polygonzkevmglobalexitrootv2mock, bound to a specific deployed contract.
func NewPolygonzkevmglobalexitrootv2mockTransactor(address common.Address, transactor bind.ContractTransactor) (*Polygonzkevmglobalexitrootv2mockTransactor, error) {
	contract, err := bindPolygonzkevmglobalexitrootv2mock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootv2mockTransactor{contract: contract}, nil
}

// NewPolygonzkevmglobalexitrootv2mockFilterer creates a new log filterer instance of Polygonzkevmglobalexitrootv2mock, bound to a specific deployed contract.
func NewPolygonzkevmglobalexitrootv2mockFilterer(address common.Address, filterer bind.ContractFilterer) (*Polygonzkevmglobalexitrootv2mockFilterer, error) {
	contract, err := bindPolygonzkevmglobalexitrootv2mock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootv2mockFilterer{contract: contract}, nil
}

// bindPolygonzkevmglobalexitrootv2mock binds a generic wrapper to an already deployed contract.
func bindPolygonzkevmglobalexitrootv2mock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Polygonzkevmglobalexitrootv2mockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmglobalexitrootv2mock.Contract.Polygonzkevmglobalexitrootv2mockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.Polygonzkevmglobalexitrootv2mockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.Polygonzkevmglobalexitrootv2mockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmglobalexitrootv2mock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootv2mock.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockSession) BridgeAddress() (common.Address, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.BridgeAddress(&_Polygonzkevmglobalexitrootv2mock.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockCallerSession) BridgeAddress() (common.Address, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.BridgeAddress(&_Polygonzkevmglobalexitrootv2mock.CallOpts)
}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockCaller) CalculateRoot(opts *bind.CallOpts, leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootv2mock.contract.Call(opts, &out, "calculateRoot", leafHash, smtProof, index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockSession) CalculateRoot(leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.CalculateRoot(&_Polygonzkevmglobalexitrootv2mock.CallOpts, leafHash, smtProof, index)
}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockCallerSession) CalculateRoot(leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.CalculateRoot(&_Polygonzkevmglobalexitrootv2mock.CallOpts, leafHash, smtProof, index)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockCaller) DepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootv2mock.contract.Call(opts, &out, "depositCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockSession) DepositCount() (*big.Int, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.DepositCount(&_Polygonzkevmglobalexitrootv2mock.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockCallerSession) DepositCount() (*big.Int, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.DepositCount(&_Polygonzkevmglobalexitrootv2mock.CallOpts)
}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockCaller) GetLastGlobalExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootv2mock.contract.Call(opts, &out, "getLastGlobalExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockSession) GetLastGlobalExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.GetLastGlobalExitRoot(&_Polygonzkevmglobalexitrootv2mock.CallOpts)
}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockCallerSession) GetLastGlobalExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.GetLastGlobalExitRoot(&_Polygonzkevmglobalexitrootv2mock.CallOpts)
}

// GetLeafValue is a free data retrieval call binding the contract method 0x5d810501.
//
// Solidity: function getLeafValue(bytes32 newGlobalExitRoot, uint256 lastBlockHash, uint64 timestamp) pure returns(bytes32)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockCaller) GetLeafValue(opts *bind.CallOpts, newGlobalExitRoot [32]byte, lastBlockHash *big.Int, timestamp uint64) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootv2mock.contract.Call(opts, &out, "getLeafValue", newGlobalExitRoot, lastBlockHash, timestamp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLeafValue is a free data retrieval call binding the contract method 0x5d810501.
//
// Solidity: function getLeafValue(bytes32 newGlobalExitRoot, uint256 lastBlockHash, uint64 timestamp) pure returns(bytes32)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockSession) GetLeafValue(newGlobalExitRoot [32]byte, lastBlockHash *big.Int, timestamp uint64) ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.GetLeafValue(&_Polygonzkevmglobalexitrootv2mock.CallOpts, newGlobalExitRoot, lastBlockHash, timestamp)
}

// GetLeafValue is a free data retrieval call binding the contract method 0x5d810501.
//
// Solidity: function getLeafValue(bytes32 newGlobalExitRoot, uint256 lastBlockHash, uint64 timestamp) pure returns(bytes32)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockCallerSession) GetLeafValue(newGlobalExitRoot [32]byte, lastBlockHash *big.Int, timestamp uint64) ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.GetLeafValue(&_Polygonzkevmglobalexitrootv2mock.CallOpts, newGlobalExitRoot, lastBlockHash, timestamp)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockCaller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootv2mock.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockSession) GetRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.GetRoot(&_Polygonzkevmglobalexitrootv2mock.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockCallerSession) GetRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.GetRoot(&_Polygonzkevmglobalexitrootv2mock.CallOpts)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockCaller) GlobalExitRootMap(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootv2mock.contract.Call(opts, &out, "globalExitRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.GlobalExitRootMap(&_Polygonzkevmglobalexitrootv2mock.CallOpts, arg0)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockCallerSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.GlobalExitRootMap(&_Polygonzkevmglobalexitrootv2mock.CallOpts, arg0)
}

// L1InfoRootMap is a free data retrieval call binding the contract method 0xef4eeb35.
//
// Solidity: function l1InfoRootMap(uint32 leafCount) view returns(bytes32 l1InfoRoot)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockCaller) L1InfoRootMap(opts *bind.CallOpts, leafCount uint32) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootv2mock.contract.Call(opts, &out, "l1InfoRootMap", leafCount)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L1InfoRootMap is a free data retrieval call binding the contract method 0xef4eeb35.
//
// Solidity: function l1InfoRootMap(uint32 leafCount) view returns(bytes32 l1InfoRoot)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockSession) L1InfoRootMap(leafCount uint32) ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.L1InfoRootMap(&_Polygonzkevmglobalexitrootv2mock.CallOpts, leafCount)
}

// L1InfoRootMap is a free data retrieval call binding the contract method 0xef4eeb35.
//
// Solidity: function l1InfoRootMap(uint32 leafCount) view returns(bytes32 l1InfoRoot)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockCallerSession) L1InfoRootMap(leafCount uint32) ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.L1InfoRootMap(&_Polygonzkevmglobalexitrootv2mock.CallOpts, leafCount)
}

// LastMainnetExitRoot is a free data retrieval call binding the contract method 0x319cf735.
//
// Solidity: function lastMainnetExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockCaller) LastMainnetExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootv2mock.contract.Call(opts, &out, "lastMainnetExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastMainnetExitRoot is a free data retrieval call binding the contract method 0x319cf735.
//
// Solidity: function lastMainnetExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockSession) LastMainnetExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.LastMainnetExitRoot(&_Polygonzkevmglobalexitrootv2mock.CallOpts)
}

// LastMainnetExitRoot is a free data retrieval call binding the contract method 0x319cf735.
//
// Solidity: function lastMainnetExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockCallerSession) LastMainnetExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.LastMainnetExitRoot(&_Polygonzkevmglobalexitrootv2mock.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockCaller) LastRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootv2mock.contract.Call(opts, &out, "lastRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockSession) LastRollupExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.LastRollupExitRoot(&_Polygonzkevmglobalexitrootv2mock.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockCallerSession) LastRollupExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.LastRollupExitRoot(&_Polygonzkevmglobalexitrootv2mock.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockCaller) RollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootv2mock.contract.Call(opts, &out, "rollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockSession) RollupManager() (common.Address, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.RollupManager(&_Polygonzkevmglobalexitrootv2mock.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockCallerSession) RollupManager() (common.Address, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.RollupManager(&_Polygonzkevmglobalexitrootv2mock.CallOpts)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockCaller) VerifyMerkleProof(opts *bind.CallOpts, leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootv2mock.contract.Call(opts, &out, "verifyMerkleProof", leafHash, smtProof, index, root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockSession) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.VerifyMerkleProof(&_Polygonzkevmglobalexitrootv2mock.CallOpts, leafHash, smtProof, index, root)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockCallerSession) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.VerifyMerkleProof(&_Polygonzkevmglobalexitrootv2mock.CallOpts, leafHash, smtProof, index, root)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootv2mock.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockSession) Initialize() (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.Initialize(&_Polygonzkevmglobalexitrootv2mock.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockTransactorSession) Initialize() (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.Initialize(&_Polygonzkevmglobalexitrootv2mock.TransactOpts)
}

// InjectGER is a paid mutator transaction binding the contract method 0xa43ca6a6.
//
// Solidity: function injectGER(bytes32 _root, uint32 depositCount) returns()
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockTransactor) InjectGER(opts *bind.TransactOpts, _root [32]byte, depositCount uint32) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootv2mock.contract.Transact(opts, "injectGER", _root, depositCount)
}

// InjectGER is a paid mutator transaction binding the contract method 0xa43ca6a6.
//
// Solidity: function injectGER(bytes32 _root, uint32 depositCount) returns()
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockSession) InjectGER(_root [32]byte, depositCount uint32) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.InjectGER(&_Polygonzkevmglobalexitrootv2mock.TransactOpts, _root, depositCount)
}

// InjectGER is a paid mutator transaction binding the contract method 0xa43ca6a6.
//
// Solidity: function injectGER(bytes32 _root, uint32 depositCount) returns()
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockTransactorSession) InjectGER(_root [32]byte, depositCount uint32) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.InjectGER(&_Polygonzkevmglobalexitrootv2mock.TransactOpts, _root, depositCount)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockTransactor) UpdateExitRoot(opts *bind.TransactOpts, newRoot [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootv2mock.contract.Transact(opts, "updateExitRoot", newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockSession) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.UpdateExitRoot(&_Polygonzkevmglobalexitrootv2mock.TransactOpts, newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockTransactorSession) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootv2mock.Contract.UpdateExitRoot(&_Polygonzkevmglobalexitrootv2mock.TransactOpts, newRoot)
}

// Polygonzkevmglobalexitrootv2mockInitL1InfoRootMapIterator is returned from FilterInitL1InfoRootMap and is used to iterate over the raw logs and unpacked data for InitL1InfoRootMap events raised by the Polygonzkevmglobalexitrootv2mock contract.
type Polygonzkevmglobalexitrootv2mockInitL1InfoRootMapIterator struct {
	Event *Polygonzkevmglobalexitrootv2mockInitL1InfoRootMap // Event containing the contract specifics and raw log

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
func (it *Polygonzkevmglobalexitrootv2mockInitL1InfoRootMapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonzkevmglobalexitrootv2mockInitL1InfoRootMap)
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
		it.Event = new(Polygonzkevmglobalexitrootv2mockInitL1InfoRootMap)
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
func (it *Polygonzkevmglobalexitrootv2mockInitL1InfoRootMapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonzkevmglobalexitrootv2mockInitL1InfoRootMapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonzkevmglobalexitrootv2mockInitL1InfoRootMap represents a InitL1InfoRootMap event raised by the Polygonzkevmglobalexitrootv2mock contract.
type Polygonzkevmglobalexitrootv2mockInitL1InfoRootMap struct {
	LeafCount         uint32
	CurrentL1InfoRoot [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterInitL1InfoRootMap is a free log retrieval operation binding the contract event 0x11f50c71891002839c2637ce302087160298255a87f1ea60d40e8db081383fad.
//
// Solidity: event InitL1InfoRootMap(uint32 leafCount, bytes32 currentL1InfoRoot)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockFilterer) FilterInitL1InfoRootMap(opts *bind.FilterOpts) (*Polygonzkevmglobalexitrootv2mockInitL1InfoRootMapIterator, error) {

	logs, sub, err := _Polygonzkevmglobalexitrootv2mock.contract.FilterLogs(opts, "InitL1InfoRootMap")
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootv2mockInitL1InfoRootMapIterator{contract: _Polygonzkevmglobalexitrootv2mock.contract, event: "InitL1InfoRootMap", logs: logs, sub: sub}, nil
}

// WatchInitL1InfoRootMap is a free log subscription operation binding the contract event 0x11f50c71891002839c2637ce302087160298255a87f1ea60d40e8db081383fad.
//
// Solidity: event InitL1InfoRootMap(uint32 leafCount, bytes32 currentL1InfoRoot)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockFilterer) WatchInitL1InfoRootMap(opts *bind.WatchOpts, sink chan<- *Polygonzkevmglobalexitrootv2mockInitL1InfoRootMap) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmglobalexitrootv2mock.contract.WatchLogs(opts, "InitL1InfoRootMap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonzkevmglobalexitrootv2mockInitL1InfoRootMap)
				if err := _Polygonzkevmglobalexitrootv2mock.contract.UnpackLog(event, "InitL1InfoRootMap", log); err != nil {
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

// ParseInitL1InfoRootMap is a log parse operation binding the contract event 0x11f50c71891002839c2637ce302087160298255a87f1ea60d40e8db081383fad.
//
// Solidity: event InitL1InfoRootMap(uint32 leafCount, bytes32 currentL1InfoRoot)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockFilterer) ParseInitL1InfoRootMap(log types.Log) (*Polygonzkevmglobalexitrootv2mockInitL1InfoRootMap, error) {
	event := new(Polygonzkevmglobalexitrootv2mockInitL1InfoRootMap)
	if err := _Polygonzkevmglobalexitrootv2mock.contract.UnpackLog(event, "InitL1InfoRootMap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonzkevmglobalexitrootv2mockInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonzkevmglobalexitrootv2mock contract.
type Polygonzkevmglobalexitrootv2mockInitializedIterator struct {
	Event *Polygonzkevmglobalexitrootv2mockInitialized // Event containing the contract specifics and raw log

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
func (it *Polygonzkevmglobalexitrootv2mockInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonzkevmglobalexitrootv2mockInitialized)
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
		it.Event = new(Polygonzkevmglobalexitrootv2mockInitialized)
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
func (it *Polygonzkevmglobalexitrootv2mockInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonzkevmglobalexitrootv2mockInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonzkevmglobalexitrootv2mockInitialized represents a Initialized event raised by the Polygonzkevmglobalexitrootv2mock contract.
type Polygonzkevmglobalexitrootv2mockInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockFilterer) FilterInitialized(opts *bind.FilterOpts) (*Polygonzkevmglobalexitrootv2mockInitializedIterator, error) {

	logs, sub, err := _Polygonzkevmglobalexitrootv2mock.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootv2mockInitializedIterator{contract: _Polygonzkevmglobalexitrootv2mock.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Polygonzkevmglobalexitrootv2mockInitialized) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmglobalexitrootv2mock.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonzkevmglobalexitrootv2mockInitialized)
				if err := _Polygonzkevmglobalexitrootv2mock.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockFilterer) ParseInitialized(log types.Log) (*Polygonzkevmglobalexitrootv2mockInitialized, error) {
	event := new(Polygonzkevmglobalexitrootv2mockInitialized)
	if err := _Polygonzkevmglobalexitrootv2mock.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonzkevmglobalexitrootv2mockUpdateL1InfoTreeIterator is returned from FilterUpdateL1InfoTree and is used to iterate over the raw logs and unpacked data for UpdateL1InfoTree events raised by the Polygonzkevmglobalexitrootv2mock contract.
type Polygonzkevmglobalexitrootv2mockUpdateL1InfoTreeIterator struct {
	Event *Polygonzkevmglobalexitrootv2mockUpdateL1InfoTree // Event containing the contract specifics and raw log

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
func (it *Polygonzkevmglobalexitrootv2mockUpdateL1InfoTreeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonzkevmglobalexitrootv2mockUpdateL1InfoTree)
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
		it.Event = new(Polygonzkevmglobalexitrootv2mockUpdateL1InfoTree)
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
func (it *Polygonzkevmglobalexitrootv2mockUpdateL1InfoTreeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonzkevmglobalexitrootv2mockUpdateL1InfoTreeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonzkevmglobalexitrootv2mockUpdateL1InfoTree represents a UpdateL1InfoTree event raised by the Polygonzkevmglobalexitrootv2mock contract.
type Polygonzkevmglobalexitrootv2mockUpdateL1InfoTree struct {
	MainnetExitRoot [32]byte
	RollupExitRoot  [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateL1InfoTree is a free log retrieval operation binding the contract event 0xda61aa7823fcd807e37b95aabcbe17f03a6f3efd514176444dae191d27fd66b3.
//
// Solidity: event UpdateL1InfoTree(bytes32 indexed mainnetExitRoot, bytes32 indexed rollupExitRoot)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockFilterer) FilterUpdateL1InfoTree(opts *bind.FilterOpts, mainnetExitRoot [][32]byte, rollupExitRoot [][32]byte) (*Polygonzkevmglobalexitrootv2mockUpdateL1InfoTreeIterator, error) {

	var mainnetExitRootRule []interface{}
	for _, mainnetExitRootItem := range mainnetExitRoot {
		mainnetExitRootRule = append(mainnetExitRootRule, mainnetExitRootItem)
	}
	var rollupExitRootRule []interface{}
	for _, rollupExitRootItem := range rollupExitRoot {
		rollupExitRootRule = append(rollupExitRootRule, rollupExitRootItem)
	}

	logs, sub, err := _Polygonzkevmglobalexitrootv2mock.contract.FilterLogs(opts, "UpdateL1InfoTree", mainnetExitRootRule, rollupExitRootRule)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootv2mockUpdateL1InfoTreeIterator{contract: _Polygonzkevmglobalexitrootv2mock.contract, event: "UpdateL1InfoTree", logs: logs, sub: sub}, nil
}

// WatchUpdateL1InfoTree is a free log subscription operation binding the contract event 0xda61aa7823fcd807e37b95aabcbe17f03a6f3efd514176444dae191d27fd66b3.
//
// Solidity: event UpdateL1InfoTree(bytes32 indexed mainnetExitRoot, bytes32 indexed rollupExitRoot)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockFilterer) WatchUpdateL1InfoTree(opts *bind.WatchOpts, sink chan<- *Polygonzkevmglobalexitrootv2mockUpdateL1InfoTree, mainnetExitRoot [][32]byte, rollupExitRoot [][32]byte) (event.Subscription, error) {

	var mainnetExitRootRule []interface{}
	for _, mainnetExitRootItem := range mainnetExitRoot {
		mainnetExitRootRule = append(mainnetExitRootRule, mainnetExitRootItem)
	}
	var rollupExitRootRule []interface{}
	for _, rollupExitRootItem := range rollupExitRoot {
		rollupExitRootRule = append(rollupExitRootRule, rollupExitRootItem)
	}

	logs, sub, err := _Polygonzkevmglobalexitrootv2mock.contract.WatchLogs(opts, "UpdateL1InfoTree", mainnetExitRootRule, rollupExitRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonzkevmglobalexitrootv2mockUpdateL1InfoTree)
				if err := _Polygonzkevmglobalexitrootv2mock.contract.UnpackLog(event, "UpdateL1InfoTree", log); err != nil {
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

// ParseUpdateL1InfoTree is a log parse operation binding the contract event 0xda61aa7823fcd807e37b95aabcbe17f03a6f3efd514176444dae191d27fd66b3.
//
// Solidity: event UpdateL1InfoTree(bytes32 indexed mainnetExitRoot, bytes32 indexed rollupExitRoot)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockFilterer) ParseUpdateL1InfoTree(log types.Log) (*Polygonzkevmglobalexitrootv2mockUpdateL1InfoTree, error) {
	event := new(Polygonzkevmglobalexitrootv2mockUpdateL1InfoTree)
	if err := _Polygonzkevmglobalexitrootv2mock.contract.UnpackLog(event, "UpdateL1InfoTree", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonzkevmglobalexitrootv2mockUpdateL1InfoTreeV2Iterator is returned from FilterUpdateL1InfoTreeV2 and is used to iterate over the raw logs and unpacked data for UpdateL1InfoTreeV2 events raised by the Polygonzkevmglobalexitrootv2mock contract.
type Polygonzkevmglobalexitrootv2mockUpdateL1InfoTreeV2Iterator struct {
	Event *Polygonzkevmglobalexitrootv2mockUpdateL1InfoTreeV2 // Event containing the contract specifics and raw log

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
func (it *Polygonzkevmglobalexitrootv2mockUpdateL1InfoTreeV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonzkevmglobalexitrootv2mockUpdateL1InfoTreeV2)
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
		it.Event = new(Polygonzkevmglobalexitrootv2mockUpdateL1InfoTreeV2)
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
func (it *Polygonzkevmglobalexitrootv2mockUpdateL1InfoTreeV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonzkevmglobalexitrootv2mockUpdateL1InfoTreeV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonzkevmglobalexitrootv2mockUpdateL1InfoTreeV2 represents a UpdateL1InfoTreeV2 event raised by the Polygonzkevmglobalexitrootv2mock contract.
type Polygonzkevmglobalexitrootv2mockUpdateL1InfoTreeV2 struct {
	CurrentL1InfoRoot [32]byte
	LeafCount         uint32
	Blockhash         *big.Int
	MinTimestamp      uint64
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdateL1InfoTreeV2 is a free log retrieval operation binding the contract event 0xaf6c6cd7790e0180a4d22eb8ed846e55846f54ed10e5946db19972b5a0813a59.
//
// Solidity: event UpdateL1InfoTreeV2(bytes32 currentL1InfoRoot, uint32 indexed leafCount, uint256 blockhash, uint64 minTimestamp)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockFilterer) FilterUpdateL1InfoTreeV2(opts *bind.FilterOpts, leafCount []uint32) (*Polygonzkevmglobalexitrootv2mockUpdateL1InfoTreeV2Iterator, error) {

	var leafCountRule []interface{}
	for _, leafCountItem := range leafCount {
		leafCountRule = append(leafCountRule, leafCountItem)
	}

	logs, sub, err := _Polygonzkevmglobalexitrootv2mock.contract.FilterLogs(opts, "UpdateL1InfoTreeV2", leafCountRule)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootv2mockUpdateL1InfoTreeV2Iterator{contract: _Polygonzkevmglobalexitrootv2mock.contract, event: "UpdateL1InfoTreeV2", logs: logs, sub: sub}, nil
}

// WatchUpdateL1InfoTreeV2 is a free log subscription operation binding the contract event 0xaf6c6cd7790e0180a4d22eb8ed846e55846f54ed10e5946db19972b5a0813a59.
//
// Solidity: event UpdateL1InfoTreeV2(bytes32 currentL1InfoRoot, uint32 indexed leafCount, uint256 blockhash, uint64 minTimestamp)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockFilterer) WatchUpdateL1InfoTreeV2(opts *bind.WatchOpts, sink chan<- *Polygonzkevmglobalexitrootv2mockUpdateL1InfoTreeV2, leafCount []uint32) (event.Subscription, error) {

	var leafCountRule []interface{}
	for _, leafCountItem := range leafCount {
		leafCountRule = append(leafCountRule, leafCountItem)
	}

	logs, sub, err := _Polygonzkevmglobalexitrootv2mock.contract.WatchLogs(opts, "UpdateL1InfoTreeV2", leafCountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonzkevmglobalexitrootv2mockUpdateL1InfoTreeV2)
				if err := _Polygonzkevmglobalexitrootv2mock.contract.UnpackLog(event, "UpdateL1InfoTreeV2", log); err != nil {
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

// ParseUpdateL1InfoTreeV2 is a log parse operation binding the contract event 0xaf6c6cd7790e0180a4d22eb8ed846e55846f54ed10e5946db19972b5a0813a59.
//
// Solidity: event UpdateL1InfoTreeV2(bytes32 currentL1InfoRoot, uint32 indexed leafCount, uint256 blockhash, uint64 minTimestamp)
func (_Polygonzkevmglobalexitrootv2mock *Polygonzkevmglobalexitrootv2mockFilterer) ParseUpdateL1InfoTreeV2(log types.Log) (*Polygonzkevmglobalexitrootv2mockUpdateL1InfoTreeV2, error) {
	event := new(Polygonzkevmglobalexitrootv2mockUpdateL1InfoTreeV2)
	if err := _Polygonzkevmglobalexitrootv2mock.contract.UnpackLog(event, "UpdateL1InfoTreeV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

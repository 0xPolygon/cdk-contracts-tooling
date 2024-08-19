// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonzkevmglobalexitrootv2

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

// Polygonzkevmglobalexitrootv2MetaData contains all meta data concerning the Polygonzkevmglobalexitrootv2 contract.
var Polygonzkevmglobalexitrootv2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollupManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"leafCount\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"currentL1InfoRoot\",\"type\":\"bytes32\"}],\"name\":\"InitL1InfoRootMap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"}],\"name\":\"UpdateL1InfoTree\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"currentL1InfoRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"leafCount\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockhash\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"minTimestamp\",\"type\":\"uint64\"}],\"name\":\"UpdateL1InfoTreeV2\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"calculateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastGlobalExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"lastBlockHash\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"getLeafValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"leafCount\",\"type\":\"uint32\"}],\"name\":\"l1InfoRootMap\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"l1InfoRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastMainnetExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"verifyMerkleProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561000f575f80fd5b50604051610e9d380380610e9d83398101604081905261002e9161012b565b6001600160a01b0380831660a0528116608052610049610050565b505061015c565b602e54610100900460ff16156100bc5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b602e5460ff908116101561010e57602e805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b80516001600160a01b0381168114610126575f80fd5b919050565b5f806040838503121561013c575f80fd5b61014583610110565b915061015360208401610110565b90509250929050565b60805160a051610d1261018b5f395f818161015701526102f701525f818161022e01526102ab0152610d125ff3fe608060405234801561000f575f80fd5b50600436106100e5575f3560e01c80635ca1e1651161008857806383f244031161006357806383f2440314610216578063a3c573eb14610229578063ef4eeb3514610250578063fb5708341461026f575f80fd5b80635ca1e1651461019e5780635d810501146101a65780638129fc1c1461020e575f80fd5b8063319cf735116100c3578063319cf7351461012c57806333d6247d146101355780633ed691ef1461014a57806349b7b80214610152575f80fd5b806301fd9044146100e9578063257b3632146101045780632dfdf0b514610123575b5f80fd5b6100f15f5481565b6040519081526020015b60405180910390f35b6100f16101123660046109bf565b60026020525f908152604090205481565b6100f160235481565b6100f160015481565b6101486101433660046109bf565b610292565b005b6100f16104b8565b6101797f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100fb565b6100f16104cb565b6100f16101b43660046109d6565b604080516020808201959095528082019390935260c09190911b7fffffffffffffffff0000000000000000000000000000000000000000000000001660608301528051604881840301815260689092019052805191012090565b6101486104d4565b6100f1610224366004610a41565b6106c8565b6101797f000000000000000000000000000000000000000000000000000000000000000081565b6100f161025e366004610a7d565b602f6020525f908152604090205481565b61028261027d366004610a9d565b61079d565b60405190151581526020016100fb565b5f8073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001633036102e057505060018190555f548161035f565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016330361032d5750505f819055600154819061035f565b6040517fb49365dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61036a82846107b4565b5f81815260026020526040812054919250036104b257425f61038d600143610b0f565b5f8481526002602090815260409182902092409283905581518082018790528083018490527fffffffffffffffff00000000000000000000000000000000000000000000000060c087901b1660608201528251808203604801815260689091019092528151910120909150610401906107e3565b5f61040a6104cb565b60235463ffffffff165f908152602f602052604080822083905551919250879187917fda61aa7823fcd807e37b95aabcbe17f03a6f3efd514176444dae191d27fd66b391a360235463ffffffff167faf6c6cd7790e0180a4d22eb8ed846e55846f54ed10e5946db19972b5a0813a598284866040516104a693929190928352602083019190915267ffffffffffffffff16604082015260600190565b60405180910390a25050505b50505050565b5f6104c66001545f546107b4565b905090565b5f6104c66108e3565b602e54610100900460ff16158080156104f45750602e54600160ff909116105b8061050e5750303b15801561050e5750602e5460ff166001145b61059e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840160405180910390fd5b602e80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156105fc57602e80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b5f6106056104cb565b6023805463ffffffff9081165f908152602f602090815260409182902085905592548151921682529181018390529192507f11f50c71891002839c2637ce302087160298255a87f1ea60d40e8db081383fad910160405180910390a15080156106c557602e80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b5f83815b602081101561079457600163ffffffff8516821c81169003610737578481602081106106fa576106fa610b22565b60200201358260405160200161071a929190918252602082015260400190565b604051602081830303815290604052805190602001209150610782565b8185826020811061074a5761074a610b22565b6020020135604051602001610769929190918252602082015260400190565b6040516020818303038152906040528051906020012091505b8061078c81610b4f565b9150506106cc565b50949350505050565b5f816107aa8686866106c8565b1495945050505050565b604080516020808201859052818301849052825180830384018152606090920190925280519101205b92915050565b8060016107f260206002610ca4565b6107fc9190610b0f565b60235410610836576040517fef5ccf6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60235f815461084590610b4f565b918290555090505f5b60208110156108d5578082901c60011660010361088157826003826020811061087957610879610b22565b015550505050565b6003816020811061089457610894610b22565b0154604080516020810192909252810184905260600160405160208183030381529060405280519060200120925080806108cd90610b4f565b91505061084e565b506108de610caf565b505050565b6023545f90819081805b60208110156109b6578083901c60011660010361094a576003816020811061091757610917610b22565b01546040805160208101929092528101859052606001604051602081830303815290604052805190602001209350610977565b60408051602081018690529081018390526060016040516020818303038152906040528051906020012093505b604080516020810184905290810183905260600160405160208183030381529060405280519060200120915080806109ae90610b4f565b9150506108ed565b50919392505050565b5f602082840312156109cf575f80fd5b5035919050565b5f805f606084860312156109e8575f80fd5b8335925060208401359150604084013567ffffffffffffffff81168114610a0d575f80fd5b809150509250925092565b8061040081018310156107dd575f80fd5b803563ffffffff81168114610a3c575f80fd5b919050565b5f805f6104408486031215610a54575f80fd5b83359250610a658560208601610a18565b9150610a746104208501610a29565b90509250925092565b5f60208284031215610a8d575f80fd5b610a9682610a29565b9392505050565b5f805f806104608587031215610ab1575f80fd5b84359350610ac28660208701610a18565b9250610ad16104208601610a29565b939692955092936104400135925050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b818103818111156107dd576107dd610ae2565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610b7f57610b7f610ae2565b5060010190565b600181815b80851115610bdf57817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610bc557610bc5610ae2565b80851615610bd257918102915b93841c9390800290610b8b565b509250929050565b5f82610bf5575060016107dd565b81610c0157505f6107dd565b8160018114610c175760028114610c2157610c3d565b60019150506107dd565b60ff841115610c3257610c32610ae2565b50506001821b6107dd565b5060208310610133831016604e8410600b8410161715610c60575081810a6107dd565b610c6a8383610b86565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610c9c57610c9c610ae2565b029392505050565b5f610a968383610be7565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52600160045260245ffdfea2646970667358221220266cf00c331b0b224bcd6976734955c98ca9e47455a941626d1788f29cf7b3bd64736f6c63430008140033",
}

// Polygonzkevmglobalexitrootv2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Polygonzkevmglobalexitrootv2MetaData.ABI instead.
var Polygonzkevmglobalexitrootv2ABI = Polygonzkevmglobalexitrootv2MetaData.ABI

// Polygonzkevmglobalexitrootv2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Polygonzkevmglobalexitrootv2MetaData.Bin instead.
var Polygonzkevmglobalexitrootv2Bin = Polygonzkevmglobalexitrootv2MetaData.Bin

// DeployPolygonzkevmglobalexitrootv2 deploys a new Ethereum contract, binding an instance of Polygonzkevmglobalexitrootv2 to it.
func DeployPolygonzkevmglobalexitrootv2(auth *bind.TransactOpts, backend bind.ContractBackend, _rollupManager common.Address, _bridgeAddress common.Address) (common.Address, *types.Transaction, *Polygonzkevmglobalexitrootv2, error) {
	parsed, err := Polygonzkevmglobalexitrootv2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Polygonzkevmglobalexitrootv2Bin), backend, _rollupManager, _bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonzkevmglobalexitrootv2{Polygonzkevmglobalexitrootv2Caller: Polygonzkevmglobalexitrootv2Caller{contract: contract}, Polygonzkevmglobalexitrootv2Transactor: Polygonzkevmglobalexitrootv2Transactor{contract: contract}, Polygonzkevmglobalexitrootv2Filterer: Polygonzkevmglobalexitrootv2Filterer{contract: contract}}, nil
}

// Polygonzkevmglobalexitrootv2 is an auto generated Go binding around an Ethereum contract.
type Polygonzkevmglobalexitrootv2 struct {
	Polygonzkevmglobalexitrootv2Caller     // Read-only binding to the contract
	Polygonzkevmglobalexitrootv2Transactor // Write-only binding to the contract
	Polygonzkevmglobalexitrootv2Filterer   // Log filterer for contract events
}

// Polygonzkevmglobalexitrootv2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Polygonzkevmglobalexitrootv2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Polygonzkevmglobalexitrootv2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Polygonzkevmglobalexitrootv2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Polygonzkevmglobalexitrootv2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Polygonzkevmglobalexitrootv2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Polygonzkevmglobalexitrootv2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Polygonzkevmglobalexitrootv2Session struct {
	Contract     *Polygonzkevmglobalexitrootv2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// Polygonzkevmglobalexitrootv2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Polygonzkevmglobalexitrootv2CallerSession struct {
	Contract *Polygonzkevmglobalexitrootv2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// Polygonzkevmglobalexitrootv2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Polygonzkevmglobalexitrootv2TransactorSession struct {
	Contract     *Polygonzkevmglobalexitrootv2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// Polygonzkevmglobalexitrootv2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Polygonzkevmglobalexitrootv2Raw struct {
	Contract *Polygonzkevmglobalexitrootv2 // Generic contract binding to access the raw methods on
}

// Polygonzkevmglobalexitrootv2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Polygonzkevmglobalexitrootv2CallerRaw struct {
	Contract *Polygonzkevmglobalexitrootv2Caller // Generic read-only contract binding to access the raw methods on
}

// Polygonzkevmglobalexitrootv2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Polygonzkevmglobalexitrootv2TransactorRaw struct {
	Contract *Polygonzkevmglobalexitrootv2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonzkevmglobalexitrootv2 creates a new instance of Polygonzkevmglobalexitrootv2, bound to a specific deployed contract.
func NewPolygonzkevmglobalexitrootv2(address common.Address, backend bind.ContractBackend) (*Polygonzkevmglobalexitrootv2, error) {
	contract, err := bindPolygonzkevmglobalexitrootv2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootv2{Polygonzkevmglobalexitrootv2Caller: Polygonzkevmglobalexitrootv2Caller{contract: contract}, Polygonzkevmglobalexitrootv2Transactor: Polygonzkevmglobalexitrootv2Transactor{contract: contract}, Polygonzkevmglobalexitrootv2Filterer: Polygonzkevmglobalexitrootv2Filterer{contract: contract}}, nil
}

// NewPolygonzkevmglobalexitrootv2Caller creates a new read-only instance of Polygonzkevmglobalexitrootv2, bound to a specific deployed contract.
func NewPolygonzkevmglobalexitrootv2Caller(address common.Address, caller bind.ContractCaller) (*Polygonzkevmglobalexitrootv2Caller, error) {
	contract, err := bindPolygonzkevmglobalexitrootv2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootv2Caller{contract: contract}, nil
}

// NewPolygonzkevmglobalexitrootv2Transactor creates a new write-only instance of Polygonzkevmglobalexitrootv2, bound to a specific deployed contract.
func NewPolygonzkevmglobalexitrootv2Transactor(address common.Address, transactor bind.ContractTransactor) (*Polygonzkevmglobalexitrootv2Transactor, error) {
	contract, err := bindPolygonzkevmglobalexitrootv2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootv2Transactor{contract: contract}, nil
}

// NewPolygonzkevmglobalexitrootv2Filterer creates a new log filterer instance of Polygonzkevmglobalexitrootv2, bound to a specific deployed contract.
func NewPolygonzkevmglobalexitrootv2Filterer(address common.Address, filterer bind.ContractFilterer) (*Polygonzkevmglobalexitrootv2Filterer, error) {
	contract, err := bindPolygonzkevmglobalexitrootv2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootv2Filterer{contract: contract}, nil
}

// bindPolygonzkevmglobalexitrootv2 binds a generic wrapper to an already deployed contract.
func bindPolygonzkevmglobalexitrootv2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Polygonzkevmglobalexitrootv2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmglobalexitrootv2.Contract.Polygonzkevmglobalexitrootv2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.Polygonzkevmglobalexitrootv2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.Polygonzkevmglobalexitrootv2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmglobalexitrootv2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Caller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootv2.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Session) BridgeAddress() (common.Address, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.BridgeAddress(&_Polygonzkevmglobalexitrootv2.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2CallerSession) BridgeAddress() (common.Address, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.BridgeAddress(&_Polygonzkevmglobalexitrootv2.CallOpts)
}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Caller) CalculateRoot(opts *bind.CallOpts, leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootv2.contract.Call(opts, &out, "calculateRoot", leafHash, smtProof, index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Session) CalculateRoot(leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.CalculateRoot(&_Polygonzkevmglobalexitrootv2.CallOpts, leafHash, smtProof, index)
}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2CallerSession) CalculateRoot(leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.CalculateRoot(&_Polygonzkevmglobalexitrootv2.CallOpts, leafHash, smtProof, index)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Caller) DepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootv2.contract.Call(opts, &out, "depositCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Session) DepositCount() (*big.Int, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.DepositCount(&_Polygonzkevmglobalexitrootv2.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2CallerSession) DepositCount() (*big.Int, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.DepositCount(&_Polygonzkevmglobalexitrootv2.CallOpts)
}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Caller) GetLastGlobalExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootv2.contract.Call(opts, &out, "getLastGlobalExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Session) GetLastGlobalExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.GetLastGlobalExitRoot(&_Polygonzkevmglobalexitrootv2.CallOpts)
}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2CallerSession) GetLastGlobalExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.GetLastGlobalExitRoot(&_Polygonzkevmglobalexitrootv2.CallOpts)
}

// GetLeafValue is a free data retrieval call binding the contract method 0x5d810501.
//
// Solidity: function getLeafValue(bytes32 newGlobalExitRoot, uint256 lastBlockHash, uint64 timestamp) pure returns(bytes32)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Caller) GetLeafValue(opts *bind.CallOpts, newGlobalExitRoot [32]byte, lastBlockHash *big.Int, timestamp uint64) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootv2.contract.Call(opts, &out, "getLeafValue", newGlobalExitRoot, lastBlockHash, timestamp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLeafValue is a free data retrieval call binding the contract method 0x5d810501.
//
// Solidity: function getLeafValue(bytes32 newGlobalExitRoot, uint256 lastBlockHash, uint64 timestamp) pure returns(bytes32)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Session) GetLeafValue(newGlobalExitRoot [32]byte, lastBlockHash *big.Int, timestamp uint64) ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.GetLeafValue(&_Polygonzkevmglobalexitrootv2.CallOpts, newGlobalExitRoot, lastBlockHash, timestamp)
}

// GetLeafValue is a free data retrieval call binding the contract method 0x5d810501.
//
// Solidity: function getLeafValue(bytes32 newGlobalExitRoot, uint256 lastBlockHash, uint64 timestamp) pure returns(bytes32)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2CallerSession) GetLeafValue(newGlobalExitRoot [32]byte, lastBlockHash *big.Int, timestamp uint64) ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.GetLeafValue(&_Polygonzkevmglobalexitrootv2.CallOpts, newGlobalExitRoot, lastBlockHash, timestamp)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Caller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootv2.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Session) GetRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.GetRoot(&_Polygonzkevmglobalexitrootv2.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2CallerSession) GetRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.GetRoot(&_Polygonzkevmglobalexitrootv2.CallOpts)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Caller) GlobalExitRootMap(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootv2.contract.Call(opts, &out, "globalExitRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Session) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.GlobalExitRootMap(&_Polygonzkevmglobalexitrootv2.CallOpts, arg0)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2CallerSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.GlobalExitRootMap(&_Polygonzkevmglobalexitrootv2.CallOpts, arg0)
}

// L1InfoRootMap is a free data retrieval call binding the contract method 0xef4eeb35.
//
// Solidity: function l1InfoRootMap(uint32 leafCount) view returns(bytes32 l1InfoRoot)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Caller) L1InfoRootMap(opts *bind.CallOpts, leafCount uint32) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootv2.contract.Call(opts, &out, "l1InfoRootMap", leafCount)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L1InfoRootMap is a free data retrieval call binding the contract method 0xef4eeb35.
//
// Solidity: function l1InfoRootMap(uint32 leafCount) view returns(bytes32 l1InfoRoot)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Session) L1InfoRootMap(leafCount uint32) ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.L1InfoRootMap(&_Polygonzkevmglobalexitrootv2.CallOpts, leafCount)
}

// L1InfoRootMap is a free data retrieval call binding the contract method 0xef4eeb35.
//
// Solidity: function l1InfoRootMap(uint32 leafCount) view returns(bytes32 l1InfoRoot)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2CallerSession) L1InfoRootMap(leafCount uint32) ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.L1InfoRootMap(&_Polygonzkevmglobalexitrootv2.CallOpts, leafCount)
}

// LastMainnetExitRoot is a free data retrieval call binding the contract method 0x319cf735.
//
// Solidity: function lastMainnetExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Caller) LastMainnetExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootv2.contract.Call(opts, &out, "lastMainnetExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastMainnetExitRoot is a free data retrieval call binding the contract method 0x319cf735.
//
// Solidity: function lastMainnetExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Session) LastMainnetExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.LastMainnetExitRoot(&_Polygonzkevmglobalexitrootv2.CallOpts)
}

// LastMainnetExitRoot is a free data retrieval call binding the contract method 0x319cf735.
//
// Solidity: function lastMainnetExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2CallerSession) LastMainnetExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.LastMainnetExitRoot(&_Polygonzkevmglobalexitrootv2.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Caller) LastRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootv2.contract.Call(opts, &out, "lastRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Session) LastRollupExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.LastRollupExitRoot(&_Polygonzkevmglobalexitrootv2.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2CallerSession) LastRollupExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.LastRollupExitRoot(&_Polygonzkevmglobalexitrootv2.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Caller) RollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootv2.contract.Call(opts, &out, "rollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Session) RollupManager() (common.Address, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.RollupManager(&_Polygonzkevmglobalexitrootv2.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2CallerSession) RollupManager() (common.Address, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.RollupManager(&_Polygonzkevmglobalexitrootv2.CallOpts)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Caller) VerifyMerkleProof(opts *bind.CallOpts, leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootv2.contract.Call(opts, &out, "verifyMerkleProof", leafHash, smtProof, index, root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Session) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.VerifyMerkleProof(&_Polygonzkevmglobalexitrootv2.CallOpts, leafHash, smtProof, index, root)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2CallerSession) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.VerifyMerkleProof(&_Polygonzkevmglobalexitrootv2.CallOpts, leafHash, smtProof, index, root)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Transactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootv2.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Session) Initialize() (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.Initialize(&_Polygonzkevmglobalexitrootv2.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2TransactorSession) Initialize() (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.Initialize(&_Polygonzkevmglobalexitrootv2.TransactOpts)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Transactor) UpdateExitRoot(opts *bind.TransactOpts, newRoot [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootv2.contract.Transact(opts, "updateExitRoot", newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Session) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.UpdateExitRoot(&_Polygonzkevmglobalexitrootv2.TransactOpts, newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2TransactorSession) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.UpdateExitRoot(&_Polygonzkevmglobalexitrootv2.TransactOpts, newRoot)
}

// Polygonzkevmglobalexitrootv2InitL1InfoRootMapIterator is returned from FilterInitL1InfoRootMap and is used to iterate over the raw logs and unpacked data for InitL1InfoRootMap events raised by the Polygonzkevmglobalexitrootv2 contract.
type Polygonzkevmglobalexitrootv2InitL1InfoRootMapIterator struct {
	Event *Polygonzkevmglobalexitrootv2InitL1InfoRootMap // Event containing the contract specifics and raw log

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
func (it *Polygonzkevmglobalexitrootv2InitL1InfoRootMapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonzkevmglobalexitrootv2InitL1InfoRootMap)
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
		it.Event = new(Polygonzkevmglobalexitrootv2InitL1InfoRootMap)
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
func (it *Polygonzkevmglobalexitrootv2InitL1InfoRootMapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonzkevmglobalexitrootv2InitL1InfoRootMapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonzkevmglobalexitrootv2InitL1InfoRootMap represents a InitL1InfoRootMap event raised by the Polygonzkevmglobalexitrootv2 contract.
type Polygonzkevmglobalexitrootv2InitL1InfoRootMap struct {
	LeafCount         uint32
	CurrentL1InfoRoot [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterInitL1InfoRootMap is a free log retrieval operation binding the contract event 0x11f50c71891002839c2637ce302087160298255a87f1ea60d40e8db081383fad.
//
// Solidity: event InitL1InfoRootMap(uint32 leafCount, bytes32 currentL1InfoRoot)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Filterer) FilterInitL1InfoRootMap(opts *bind.FilterOpts) (*Polygonzkevmglobalexitrootv2InitL1InfoRootMapIterator, error) {

	logs, sub, err := _Polygonzkevmglobalexitrootv2.contract.FilterLogs(opts, "InitL1InfoRootMap")
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootv2InitL1InfoRootMapIterator{contract: _Polygonzkevmglobalexitrootv2.contract, event: "InitL1InfoRootMap", logs: logs, sub: sub}, nil
}

// WatchInitL1InfoRootMap is a free log subscription operation binding the contract event 0x11f50c71891002839c2637ce302087160298255a87f1ea60d40e8db081383fad.
//
// Solidity: event InitL1InfoRootMap(uint32 leafCount, bytes32 currentL1InfoRoot)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Filterer) WatchInitL1InfoRootMap(opts *bind.WatchOpts, sink chan<- *Polygonzkevmglobalexitrootv2InitL1InfoRootMap) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmglobalexitrootv2.contract.WatchLogs(opts, "InitL1InfoRootMap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonzkevmglobalexitrootv2InitL1InfoRootMap)
				if err := _Polygonzkevmglobalexitrootv2.contract.UnpackLog(event, "InitL1InfoRootMap", log); err != nil {
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
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Filterer) ParseInitL1InfoRootMap(log types.Log) (*Polygonzkevmglobalexitrootv2InitL1InfoRootMap, error) {
	event := new(Polygonzkevmglobalexitrootv2InitL1InfoRootMap)
	if err := _Polygonzkevmglobalexitrootv2.contract.UnpackLog(event, "InitL1InfoRootMap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonzkevmglobalexitrootv2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonzkevmglobalexitrootv2 contract.
type Polygonzkevmglobalexitrootv2InitializedIterator struct {
	Event *Polygonzkevmglobalexitrootv2Initialized // Event containing the contract specifics and raw log

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
func (it *Polygonzkevmglobalexitrootv2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonzkevmglobalexitrootv2Initialized)
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
		it.Event = new(Polygonzkevmglobalexitrootv2Initialized)
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
func (it *Polygonzkevmglobalexitrootv2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonzkevmglobalexitrootv2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonzkevmglobalexitrootv2Initialized represents a Initialized event raised by the Polygonzkevmglobalexitrootv2 contract.
type Polygonzkevmglobalexitrootv2Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Filterer) FilterInitialized(opts *bind.FilterOpts) (*Polygonzkevmglobalexitrootv2InitializedIterator, error) {

	logs, sub, err := _Polygonzkevmglobalexitrootv2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootv2InitializedIterator{contract: _Polygonzkevmglobalexitrootv2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Polygonzkevmglobalexitrootv2Initialized) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmglobalexitrootv2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonzkevmglobalexitrootv2Initialized)
				if err := _Polygonzkevmglobalexitrootv2.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Filterer) ParseInitialized(log types.Log) (*Polygonzkevmglobalexitrootv2Initialized, error) {
	event := new(Polygonzkevmglobalexitrootv2Initialized)
	if err := _Polygonzkevmglobalexitrootv2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonzkevmglobalexitrootv2UpdateL1InfoTreeIterator is returned from FilterUpdateL1InfoTree and is used to iterate over the raw logs and unpacked data for UpdateL1InfoTree events raised by the Polygonzkevmglobalexitrootv2 contract.
type Polygonzkevmglobalexitrootv2UpdateL1InfoTreeIterator struct {
	Event *Polygonzkevmglobalexitrootv2UpdateL1InfoTree // Event containing the contract specifics and raw log

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
func (it *Polygonzkevmglobalexitrootv2UpdateL1InfoTreeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonzkevmglobalexitrootv2UpdateL1InfoTree)
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
		it.Event = new(Polygonzkevmglobalexitrootv2UpdateL1InfoTree)
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
func (it *Polygonzkevmglobalexitrootv2UpdateL1InfoTreeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonzkevmglobalexitrootv2UpdateL1InfoTreeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonzkevmglobalexitrootv2UpdateL1InfoTree represents a UpdateL1InfoTree event raised by the Polygonzkevmglobalexitrootv2 contract.
type Polygonzkevmglobalexitrootv2UpdateL1InfoTree struct {
	MainnetExitRoot [32]byte
	RollupExitRoot  [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateL1InfoTree is a free log retrieval operation binding the contract event 0xda61aa7823fcd807e37b95aabcbe17f03a6f3efd514176444dae191d27fd66b3.
//
// Solidity: event UpdateL1InfoTree(bytes32 indexed mainnetExitRoot, bytes32 indexed rollupExitRoot)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Filterer) FilterUpdateL1InfoTree(opts *bind.FilterOpts, mainnetExitRoot [][32]byte, rollupExitRoot [][32]byte) (*Polygonzkevmglobalexitrootv2UpdateL1InfoTreeIterator, error) {

	var mainnetExitRootRule []interface{}
	for _, mainnetExitRootItem := range mainnetExitRoot {
		mainnetExitRootRule = append(mainnetExitRootRule, mainnetExitRootItem)
	}
	var rollupExitRootRule []interface{}
	for _, rollupExitRootItem := range rollupExitRoot {
		rollupExitRootRule = append(rollupExitRootRule, rollupExitRootItem)
	}

	logs, sub, err := _Polygonzkevmglobalexitrootv2.contract.FilterLogs(opts, "UpdateL1InfoTree", mainnetExitRootRule, rollupExitRootRule)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootv2UpdateL1InfoTreeIterator{contract: _Polygonzkevmglobalexitrootv2.contract, event: "UpdateL1InfoTree", logs: logs, sub: sub}, nil
}

// WatchUpdateL1InfoTree is a free log subscription operation binding the contract event 0xda61aa7823fcd807e37b95aabcbe17f03a6f3efd514176444dae191d27fd66b3.
//
// Solidity: event UpdateL1InfoTree(bytes32 indexed mainnetExitRoot, bytes32 indexed rollupExitRoot)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Filterer) WatchUpdateL1InfoTree(opts *bind.WatchOpts, sink chan<- *Polygonzkevmglobalexitrootv2UpdateL1InfoTree, mainnetExitRoot [][32]byte, rollupExitRoot [][32]byte) (event.Subscription, error) {

	var mainnetExitRootRule []interface{}
	for _, mainnetExitRootItem := range mainnetExitRoot {
		mainnetExitRootRule = append(mainnetExitRootRule, mainnetExitRootItem)
	}
	var rollupExitRootRule []interface{}
	for _, rollupExitRootItem := range rollupExitRoot {
		rollupExitRootRule = append(rollupExitRootRule, rollupExitRootItem)
	}

	logs, sub, err := _Polygonzkevmglobalexitrootv2.contract.WatchLogs(opts, "UpdateL1InfoTree", mainnetExitRootRule, rollupExitRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonzkevmglobalexitrootv2UpdateL1InfoTree)
				if err := _Polygonzkevmglobalexitrootv2.contract.UnpackLog(event, "UpdateL1InfoTree", log); err != nil {
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
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Filterer) ParseUpdateL1InfoTree(log types.Log) (*Polygonzkevmglobalexitrootv2UpdateL1InfoTree, error) {
	event := new(Polygonzkevmglobalexitrootv2UpdateL1InfoTree)
	if err := _Polygonzkevmglobalexitrootv2.contract.UnpackLog(event, "UpdateL1InfoTree", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonzkevmglobalexitrootv2UpdateL1InfoTreeV2Iterator is returned from FilterUpdateL1InfoTreeV2 and is used to iterate over the raw logs and unpacked data for UpdateL1InfoTreeV2 events raised by the Polygonzkevmglobalexitrootv2 contract.
type Polygonzkevmglobalexitrootv2UpdateL1InfoTreeV2Iterator struct {
	Event *Polygonzkevmglobalexitrootv2UpdateL1InfoTreeV2 // Event containing the contract specifics and raw log

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
func (it *Polygonzkevmglobalexitrootv2UpdateL1InfoTreeV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonzkevmglobalexitrootv2UpdateL1InfoTreeV2)
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
		it.Event = new(Polygonzkevmglobalexitrootv2UpdateL1InfoTreeV2)
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
func (it *Polygonzkevmglobalexitrootv2UpdateL1InfoTreeV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonzkevmglobalexitrootv2UpdateL1InfoTreeV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonzkevmglobalexitrootv2UpdateL1InfoTreeV2 represents a UpdateL1InfoTreeV2 event raised by the Polygonzkevmglobalexitrootv2 contract.
type Polygonzkevmglobalexitrootv2UpdateL1InfoTreeV2 struct {
	CurrentL1InfoRoot [32]byte
	LeafCount         uint32
	Blockhash         *big.Int
	MinTimestamp      uint64
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdateL1InfoTreeV2 is a free log retrieval operation binding the contract event 0xaf6c6cd7790e0180a4d22eb8ed846e55846f54ed10e5946db19972b5a0813a59.
//
// Solidity: event UpdateL1InfoTreeV2(bytes32 currentL1InfoRoot, uint32 indexed leafCount, uint256 blockhash, uint64 minTimestamp)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Filterer) FilterUpdateL1InfoTreeV2(opts *bind.FilterOpts, leafCount []uint32) (*Polygonzkevmglobalexitrootv2UpdateL1InfoTreeV2Iterator, error) {

	var leafCountRule []interface{}
	for _, leafCountItem := range leafCount {
		leafCountRule = append(leafCountRule, leafCountItem)
	}

	logs, sub, err := _Polygonzkevmglobalexitrootv2.contract.FilterLogs(opts, "UpdateL1InfoTreeV2", leafCountRule)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootv2UpdateL1InfoTreeV2Iterator{contract: _Polygonzkevmglobalexitrootv2.contract, event: "UpdateL1InfoTreeV2", logs: logs, sub: sub}, nil
}

// WatchUpdateL1InfoTreeV2 is a free log subscription operation binding the contract event 0xaf6c6cd7790e0180a4d22eb8ed846e55846f54ed10e5946db19972b5a0813a59.
//
// Solidity: event UpdateL1InfoTreeV2(bytes32 currentL1InfoRoot, uint32 indexed leafCount, uint256 blockhash, uint64 minTimestamp)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Filterer) WatchUpdateL1InfoTreeV2(opts *bind.WatchOpts, sink chan<- *Polygonzkevmglobalexitrootv2UpdateL1InfoTreeV2, leafCount []uint32) (event.Subscription, error) {

	var leafCountRule []interface{}
	for _, leafCountItem := range leafCount {
		leafCountRule = append(leafCountRule, leafCountItem)
	}

	logs, sub, err := _Polygonzkevmglobalexitrootv2.contract.WatchLogs(opts, "UpdateL1InfoTreeV2", leafCountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonzkevmglobalexitrootv2UpdateL1InfoTreeV2)
				if err := _Polygonzkevmglobalexitrootv2.contract.UnpackLog(event, "UpdateL1InfoTreeV2", log); err != nil {
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
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Filterer) ParseUpdateL1InfoTreeV2(log types.Log) (*Polygonzkevmglobalexitrootv2UpdateL1InfoTreeV2, error) {
	event := new(Polygonzkevmglobalexitrootv2UpdateL1InfoTreeV2)
	if err := _Polygonzkevmglobalexitrootv2.contract.UnpackLog(event, "UpdateL1InfoTreeV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

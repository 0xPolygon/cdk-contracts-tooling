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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollupManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"leafIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"currentL1InfoRoot\",\"type\":\"bytes32\"}],\"name\":\"InitL1InfoRootMap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"currentL1InfoRoot\",\"type\":\"bytes32\"}],\"name\":\"UpdateL1InfoTree\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"calculateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastGlobalExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"lastBlockHash\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"getLeafValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"leafIndex\",\"type\":\"uint32\"}],\"name\":\"l1InfoRootMap\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"l1InfoRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastMainnetExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"verifyMerkleProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561001057600080fd5b50604051610e8f380380610e8f83398101604081905261002f9161012d565b6001600160a01b0380831660a052811660805261004a610051565b5050610160565b602e54610100900460ff16156100bd5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b602e5460ff908116101561010f57602e805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b80516001600160a01b038116811461012857600080fd5b919050565b6000806040838503121561014057600080fd5b61014983610111565b915061015760208401610111565b90509250929050565b60805160a051610cfc6101936000396000818161015f015261030201526000818161023601526102b50152610cfc6000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80635ca1e1651161008c57806383f244031161006657806383f244031461021e578063a3c573eb14610231578063ef4eeb3514610258578063fb5708341461027857600080fd5b80635ca1e165146101a65780635d810501146101ae5780638129fc1c1461021657600080fd5b8063319cf735116100c8578063319cf7351461013457806333d6247d1461013d5780633ed691ef1461015257806349b7b8021461015a57600080fd5b806301fd9044146100ef578063257b36321461010b5780632dfdf0b51461012b575b600080fd5b6100f860005481565b6040519081526020015b60405180910390f35b6100f861011936600461098f565b60026020526000908152604090205481565b6100f860235481565b6100f860015481565b61015061014b36600461098f565b61029b565b005b6100f861047d565b6101817f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610102565b6100f8610492565b6100f86101bc3660046109a8565b604080516020808201959095528082019390935260c09190911b7fffffffffffffffff0000000000000000000000000000000000000000000000001660608301528051604881840301815260689092019052805191012090565b61015061049c565b6100f861022c366004610a19565b610692565b6101817f000000000000000000000000000000000000000000000000000000000000000081565b6100f8610266366004610a58565b602f6020526000908152604090205481565b61028b610286366004610a7a565b610768565b6040519015158152602001610102565b60008073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001633036102eb57505060018190556000548161036b565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001633036103395750506000819055600154819061036b565b6040517fb49365dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006103778284610780565b6000818152600260205260408120549192500361047757600061039b600143610af1565b60008381526002602090815260409182902092409283905581518082018690528083018490527fffffffffffffffff0000000000000000000000000000000000000000000000004260c01b166060820152825180820360480181526068909101909252815191012090915061040f906107af565b6000610419610492565b60235463ffffffff166000908152602f60205260409081902082905551909150859085907ffe83b6f5b84bf01e2e7e3c0efaf3626b47dd84ade929d7b8002b237fe8d5427f9061046c9085815260200190565b60405180910390a350505b50505050565b600061048d600154600054610780565b905090565b600061048d6108b2565b602e54610100900460ff16158080156104bc5750602e54600160ff909116105b806104d65750303b1580156104d65750602e5460ff166001145b610566576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840160405180910390fd5b602e80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156105c457602e80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b60006105ce610492565b6023805463ffffffff9081166000908152602f602090815260409182902085905592548151921682529181018390529192507f11f50c71891002839c2637ce302087160298255a87f1ea60d40e8db081383fad910160405180910390a150801561068f57602e80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b600083815b602081101561075f57600163ffffffff8516821c81169003610702578481602081106106c5576106c5610b04565b6020020135826040516020016106e5929190918252602082015260400190565b60405160208183030381529060405280519060200120915061074d565b8185826020811061071557610715610b04565b6020020135604051602001610734929190918252602082015260400190565b6040516020818303038152906040528051906020012091505b8061075781610b33565b915050610697565b50949350505050565b600081610776868686610692565b1495945050505050565b604080516020808201859052818301849052825180830384018152606090920190925280519101205b92915050565b8060016107be60206002610c8b565b6107c89190610af1565b60235410610802576040517fef5ccf6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060236000815461081390610b33565b9182905550905060005b60208110156108a4578082901c60011660010361085057826003826020811061084857610848610b04565b015550505050565b6003816020811061086357610863610b04565b01546040805160208101929092528101849052606001604051602081830303815290604052805190602001209250808061089c90610b33565b91505061081d565b506108ad610c97565b505050565b602354600090819081805b6020811015610986578083901c60011660010361091a57600381602081106108e7576108e7610b04565b01546040805160208101929092528101859052606001604051602081830303815290604052805190602001209350610947565b60408051602081018690529081018390526060016040516020818303038152906040528051906020012093505b6040805160208101849052908101839052606001604051602081830303815290604052805190602001209150808061097e90610b33565b9150506108bd565b50919392505050565b6000602082840312156109a157600080fd5b5035919050565b6000806000606084860312156109bd57600080fd5b8335925060208401359150604084013567ffffffffffffffff811681146109e357600080fd5b809150509250925092565b8061040081018310156107a957600080fd5b803563ffffffff81168114610a1457600080fd5b919050565b60008060006104408486031215610a2f57600080fd5b83359250610a4085602086016109ee565b9150610a4f6104208501610a00565b90509250925092565b600060208284031215610a6a57600080fd5b610a7382610a00565b9392505050565b6000806000806104608587031215610a9157600080fd5b84359350610aa286602087016109ee565b9250610ab16104208601610a00565b939692955092936104400135925050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818103818111156107a9576107a9610ac2565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610b6457610b64610ac2565b5060010190565b600181815b80851115610bc457817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610baa57610baa610ac2565b80851615610bb757918102915b93841c9390800290610b70565b509250929050565b600082610bdb575060016107a9565b81610be8575060006107a9565b8160018114610bfe5760028114610c0857610c24565b60019150506107a9565b60ff841115610c1957610c19610ac2565b50506001821b6107a9565b5060208310610133831016604e8410600b8410161715610c47575081810a6107a9565b610c518383610b6b565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610c8357610c83610ac2565b029392505050565b6000610a738383610bcc565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fdfea2646970667358221220f17a69b2e768e85b4758bf42c4f597eaa08622df67fedb7a70362f4cbe21a29d64736f6c63430008140033",
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
// Solidity: function l1InfoRootMap(uint32 leafIndex) view returns(bytes32 l1InfoRoot)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Caller) L1InfoRootMap(opts *bind.CallOpts, leafIndex uint32) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootv2.contract.Call(opts, &out, "l1InfoRootMap", leafIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L1InfoRootMap is a free data retrieval call binding the contract method 0xef4eeb35.
//
// Solidity: function l1InfoRootMap(uint32 leafIndex) view returns(bytes32 l1InfoRoot)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Session) L1InfoRootMap(leafIndex uint32) ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.L1InfoRootMap(&_Polygonzkevmglobalexitrootv2.CallOpts, leafIndex)
}

// L1InfoRootMap is a free data retrieval call binding the contract method 0xef4eeb35.
//
// Solidity: function l1InfoRootMap(uint32 leafIndex) view returns(bytes32 l1InfoRoot)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2CallerSession) L1InfoRootMap(leafIndex uint32) ([32]byte, error) {
	return _Polygonzkevmglobalexitrootv2.Contract.L1InfoRootMap(&_Polygonzkevmglobalexitrootv2.CallOpts, leafIndex)
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
	LeafIndex         uint32
	CurrentL1InfoRoot [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterInitL1InfoRootMap is a free log retrieval operation binding the contract event 0x11f50c71891002839c2637ce302087160298255a87f1ea60d40e8db081383fad.
//
// Solidity: event InitL1InfoRootMap(uint32 leafIndex, bytes32 currentL1InfoRoot)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Filterer) FilterInitL1InfoRootMap(opts *bind.FilterOpts) (*Polygonzkevmglobalexitrootv2InitL1InfoRootMapIterator, error) {

	logs, sub, err := _Polygonzkevmglobalexitrootv2.contract.FilterLogs(opts, "InitL1InfoRootMap")
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootv2InitL1InfoRootMapIterator{contract: _Polygonzkevmglobalexitrootv2.contract, event: "InitL1InfoRootMap", logs: logs, sub: sub}, nil
}

// WatchInitL1InfoRootMap is a free log subscription operation binding the contract event 0x11f50c71891002839c2637ce302087160298255a87f1ea60d40e8db081383fad.
//
// Solidity: event InitL1InfoRootMap(uint32 leafIndex, bytes32 currentL1InfoRoot)
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
// Solidity: event InitL1InfoRootMap(uint32 leafIndex, bytes32 currentL1InfoRoot)
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
	MainnetExitRoot   [32]byte
	RollupExitRoot    [32]byte
	CurrentL1InfoRoot [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdateL1InfoTree is a free log retrieval operation binding the contract event 0xfe83b6f5b84bf01e2e7e3c0efaf3626b47dd84ade929d7b8002b237fe8d5427f.
//
// Solidity: event UpdateL1InfoTree(bytes32 indexed mainnetExitRoot, bytes32 indexed rollupExitRoot, bytes32 currentL1InfoRoot)
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

// WatchUpdateL1InfoTree is a free log subscription operation binding the contract event 0xfe83b6f5b84bf01e2e7e3c0efaf3626b47dd84ade929d7b8002b237fe8d5427f.
//
// Solidity: event UpdateL1InfoTree(bytes32 indexed mainnetExitRoot, bytes32 indexed rollupExitRoot, bytes32 currentL1InfoRoot)
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

// ParseUpdateL1InfoTree is a log parse operation binding the contract event 0xfe83b6f5b84bf01e2e7e3c0efaf3626b47dd84ade929d7b8002b237fe8d5427f.
//
// Solidity: event UpdateL1InfoTree(bytes32 indexed mainnetExitRoot, bytes32 indexed rollupExitRoot, bytes32 currentL1InfoRoot)
func (_Polygonzkevmglobalexitrootv2 *Polygonzkevmglobalexitrootv2Filterer) ParseUpdateL1InfoTree(log types.Log) (*Polygonzkevmglobalexitrootv2UpdateL1InfoTree, error) {
	event := new(Polygonzkevmglobalexitrootv2UpdateL1InfoTree)
	if err := _Polygonzkevmglobalexitrootv2.contract.UnpackLog(event, "UpdateL1InfoTree", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

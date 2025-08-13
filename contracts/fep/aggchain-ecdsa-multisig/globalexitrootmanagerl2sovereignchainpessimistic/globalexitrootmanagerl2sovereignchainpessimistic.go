// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package globalexitrootmanagerl2sovereignchainpessimistic

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

// Globalexitrootmanagerl2sovereignchainpessimisticMetaData contains all meta data concerning the Globalexitrootmanagerl2sovereignchainpessimistic contract.
var Globalexitrootmanagerl2sovereignchainpessimisticMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"GlobalExitRootAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughGlobalExitRootsInserted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotLastInsertedGlobalExitRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootRemover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootUpdater\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newGlobalExitRoot\",\"type\":\"bytes32\"}],\"name\":\"InsertGlobalExitRoot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"removedGlobalExitRoot\",\"type\":\"bytes32\"}],\"name\":\"RemoveLastGlobalExitRoot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGlobalExitRootRemover\",\"type\":\"address\"}],\"name\":\"SetGlobalExitRootRemover\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGlobalExitRootUpdater\",\"type\":\"address\"}],\"name\":\"SetGlobalExitRootUpdater\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootRemover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_globalExitRootUpdater\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_globalExitRootRemover\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_newRoot\",\"type\":\"bytes32\"}],\"name\":\"insertGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"insertedGERCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"gersToRemove\",\"type\":\"bytes32[]\"}],\"name\":\"removeLastGlobalExitRoots\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_globalExitRootRemover\",\"type\":\"address\"}],\"name\":\"setGlobalExitRootRemover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_globalExitRootUpdater\",\"type\":\"address\"}],\"name\":\"setGlobalExitRootUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b50604051610c89380380610c8983398101604081905261002e91610109565b6001600160a01b038116608052610043610049565b50610136565b603454610100900460ff16156100b55760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60345460ff9081161015610107576034805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b5f60208284031215610119575f80fd5b81516001600160a01b038116811461012f575f80fd5b9392505050565b608051610b346101555f395f81816101e801526103820152610b345ff3fe608060405234801561000f575f80fd5b50600436106100cf575f3560e01c80636da0e4ab1161007d57806391eb796d1161005857806391eb796d146101c3578063a3c573eb146101e3578063d0267f391461020a575f80fd5b80636da0e4ab1461015c5780637c314ce31461016f5780638bd0eb1c146101ba575f80fd5b806333d6247d116100ad57806333d6247d14610123578063485cc9551461013657806357dfb57214610149575f80fd5b806301fd9044146100d357806312da06b2146100ef578063257b363214610104575b5f80fd5b6100dc60015481565b6040519081526020015b60405180910390f35b6101026100fd36600461093a565b61021d565b005b6100dc61011236600461093a565b5f6020819052908152604090205481565b61010261013136600461093a565b61036a565b610102610144366004610979565b6103de565b6101026101573660046109aa565b6105ea565b61010261016a366004610a19565b61074e565b6034546101959062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100e6565b6100dc60365481565b6035546101959073ffffffffffffffffffffffffffffffffffffffff1681565b6101957f000000000000000000000000000000000000000000000000000000000000000081565b610102610218366004610a19565b61087c565b60345462010000900473ffffffffffffffffffffffffffffffffffffffff1661027e57413314610279576040517fc758fc1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6102d5565b60345462010000900473ffffffffffffffffffffffffffffffffffffffff1633146102d5576040517fc758fc1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8181526020819052604081205490036103385760365f81546102f790610a66565b91829055505f8281526020819052604080822092909255905182917fb1b866fe5fac68e8f1a4ab2520c7a6b493a954934bbd0f054bd91d6674a4c0d591a250565b6040517f1f97a58200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146103d9576040517fb49365dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600155565b603454610100900460ff16158080156103fe5750603454600160ff909116105b806104185750303b158015610418575060345460ff166001145b6104a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840160405180910390fd5b603480547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561050657603480547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b603480547fffffffffffffffffffff0000000000000000000000000000000000000000ffff166201000073ffffffffffffffffffffffffffffffffffffffff8681169190910291909117909155603580547fffffffffffffffffffffffff00000000000000000000000000000000000000001691841691909117905580156105e557603480547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b60355473ffffffffffffffffffffffffffffffffffffffff16331461063b576040517fa34ddeb100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60365480821115610678576040517f56feb4f500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b82811015610746575f84848381811061069557610695610a9d565b9050602002013590505f805f8381526020019081526020015f205490508381146106eb576040517fae765ff600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f828152602081905260408120558361070381610aca565b6040519095508391507f605764d0b65b62ecf05dc90f674a00a2e2531fabaf120fdde65790e407fcb7a2905f90a25050808061073e90610a66565b91505061067a565b506036555050565b60345462010000900473ffffffffffffffffffffffffffffffffffffffff166107af574133146107aa576040517fc758fc1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610806565b60345462010000900473ffffffffffffffffffffffffffffffffffffffff163314610806576040517fc758fc1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603480547fffffffffffffffffffff0000000000000000000000000000000000000000ffff166201000073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040517f992b80814dbc3fba903486d81daddb07d1d5b20483742458c8b0540e3a37e37c905f90a250565b60355473ffffffffffffffffffffffffffffffffffffffff1633146108cd576040517fa34ddeb100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517eb4672b6135d1dfbd4e9520e01abb14ea5eac645990b0d24dfda00ae999b758905f90a250565b5f6020828403121561094a575f80fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610974575f80fd5b919050565b5f806040838503121561098a575f80fd5b61099383610951565b91506109a160208401610951565b90509250929050565b5f80602083850312156109bb575f80fd5b823567ffffffffffffffff808211156109d2575f80fd5b818501915085601f8301126109e5575f80fd5b8135818111156109f3575f80fd5b8660208260051b8501011115610a07575f80fd5b60209290920196919550909350505050565b5f60208284031215610a29575f80fd5b610a3282610951565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610a9657610a96610a39565b5060010190565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81610ad857610ad8610a39565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019056fea2646970667358221220ba4814d4ad63912bd705f8743f900d79aebfc4295139d4b3fdc9395b63238e6864736f6c63430008140033",
}

// Globalexitrootmanagerl2sovereignchainpessimisticABI is the input ABI used to generate the binding from.
// Deprecated: Use Globalexitrootmanagerl2sovereignchainpessimisticMetaData.ABI instead.
var Globalexitrootmanagerl2sovereignchainpessimisticABI = Globalexitrootmanagerl2sovereignchainpessimisticMetaData.ABI

// Globalexitrootmanagerl2sovereignchainpessimisticBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Globalexitrootmanagerl2sovereignchainpessimisticMetaData.Bin instead.
var Globalexitrootmanagerl2sovereignchainpessimisticBin = Globalexitrootmanagerl2sovereignchainpessimisticMetaData.Bin

// DeployGlobalexitrootmanagerl2sovereignchainpessimistic deploys a new Ethereum contract, binding an instance of Globalexitrootmanagerl2sovereignchainpessimistic to it.
func DeployGlobalexitrootmanagerl2sovereignchainpessimistic(auth *bind.TransactOpts, backend bind.ContractBackend, _bridgeAddress common.Address) (common.Address, *types.Transaction, *Globalexitrootmanagerl2sovereignchainpessimistic, error) {
	parsed, err := Globalexitrootmanagerl2sovereignchainpessimisticMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Globalexitrootmanagerl2sovereignchainpessimisticBin), backend, _bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Globalexitrootmanagerl2sovereignchainpessimistic{Globalexitrootmanagerl2sovereignchainpessimisticCaller: Globalexitrootmanagerl2sovereignchainpessimisticCaller{contract: contract}, Globalexitrootmanagerl2sovereignchainpessimisticTransactor: Globalexitrootmanagerl2sovereignchainpessimisticTransactor{contract: contract}, Globalexitrootmanagerl2sovereignchainpessimisticFilterer: Globalexitrootmanagerl2sovereignchainpessimisticFilterer{contract: contract}}, nil
}

// Globalexitrootmanagerl2sovereignchainpessimistic is an auto generated Go binding around an Ethereum contract.
type Globalexitrootmanagerl2sovereignchainpessimistic struct {
	Globalexitrootmanagerl2sovereignchainpessimisticCaller     // Read-only binding to the contract
	Globalexitrootmanagerl2sovereignchainpessimisticTransactor // Write-only binding to the contract
	Globalexitrootmanagerl2sovereignchainpessimisticFilterer   // Log filterer for contract events
}

// Globalexitrootmanagerl2sovereignchainpessimisticCaller is an auto generated read-only Go binding around an Ethereum contract.
type Globalexitrootmanagerl2sovereignchainpessimisticCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Globalexitrootmanagerl2sovereignchainpessimisticTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Globalexitrootmanagerl2sovereignchainpessimisticTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Globalexitrootmanagerl2sovereignchainpessimisticFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Globalexitrootmanagerl2sovereignchainpessimisticFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Globalexitrootmanagerl2sovereignchainpessimisticSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Globalexitrootmanagerl2sovereignchainpessimisticSession struct {
	Contract     *Globalexitrootmanagerl2sovereignchainpessimistic // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                                     // Call options to use throughout this session
	TransactOpts bind.TransactOpts                                 // Transaction auth options to use throughout this session
}

// Globalexitrootmanagerl2sovereignchainpessimisticCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Globalexitrootmanagerl2sovereignchainpessimisticCallerSession struct {
	Contract *Globalexitrootmanagerl2sovereignchainpessimisticCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                           // Call options to use throughout this session
}

// Globalexitrootmanagerl2sovereignchainpessimisticTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Globalexitrootmanagerl2sovereignchainpessimisticTransactorSession struct {
	Contract     *Globalexitrootmanagerl2sovereignchainpessimisticTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                           // Transaction auth options to use throughout this session
}

// Globalexitrootmanagerl2sovereignchainpessimisticRaw is an auto generated low-level Go binding around an Ethereum contract.
type Globalexitrootmanagerl2sovereignchainpessimisticRaw struct {
	Contract *Globalexitrootmanagerl2sovereignchainpessimistic // Generic contract binding to access the raw methods on
}

// Globalexitrootmanagerl2sovereignchainpessimisticCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Globalexitrootmanagerl2sovereignchainpessimisticCallerRaw struct {
	Contract *Globalexitrootmanagerl2sovereignchainpessimisticCaller // Generic read-only contract binding to access the raw methods on
}

// Globalexitrootmanagerl2sovereignchainpessimisticTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Globalexitrootmanagerl2sovereignchainpessimisticTransactorRaw struct {
	Contract *Globalexitrootmanagerl2sovereignchainpessimisticTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlobalexitrootmanagerl2sovereignchainpessimistic creates a new instance of Globalexitrootmanagerl2sovereignchainpessimistic, bound to a specific deployed contract.
func NewGlobalexitrootmanagerl2sovereignchainpessimistic(address common.Address, backend bind.ContractBackend) (*Globalexitrootmanagerl2sovereignchainpessimistic, error) {
	contract, err := bindGlobalexitrootmanagerl2sovereignchainpessimistic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainpessimistic{Globalexitrootmanagerl2sovereignchainpessimisticCaller: Globalexitrootmanagerl2sovereignchainpessimisticCaller{contract: contract}, Globalexitrootmanagerl2sovereignchainpessimisticTransactor: Globalexitrootmanagerl2sovereignchainpessimisticTransactor{contract: contract}, Globalexitrootmanagerl2sovereignchainpessimisticFilterer: Globalexitrootmanagerl2sovereignchainpessimisticFilterer{contract: contract}}, nil
}

// NewGlobalexitrootmanagerl2sovereignchainpessimisticCaller creates a new read-only instance of Globalexitrootmanagerl2sovereignchainpessimistic, bound to a specific deployed contract.
func NewGlobalexitrootmanagerl2sovereignchainpessimisticCaller(address common.Address, caller bind.ContractCaller) (*Globalexitrootmanagerl2sovereignchainpessimisticCaller, error) {
	contract, err := bindGlobalexitrootmanagerl2sovereignchainpessimistic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainpessimisticCaller{contract: contract}, nil
}

// NewGlobalexitrootmanagerl2sovereignchainpessimisticTransactor creates a new write-only instance of Globalexitrootmanagerl2sovereignchainpessimistic, bound to a specific deployed contract.
func NewGlobalexitrootmanagerl2sovereignchainpessimisticTransactor(address common.Address, transactor bind.ContractTransactor) (*Globalexitrootmanagerl2sovereignchainpessimisticTransactor, error) {
	contract, err := bindGlobalexitrootmanagerl2sovereignchainpessimistic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainpessimisticTransactor{contract: contract}, nil
}

// NewGlobalexitrootmanagerl2sovereignchainpessimisticFilterer creates a new log filterer instance of Globalexitrootmanagerl2sovereignchainpessimistic, bound to a specific deployed contract.
func NewGlobalexitrootmanagerl2sovereignchainpessimisticFilterer(address common.Address, filterer bind.ContractFilterer) (*Globalexitrootmanagerl2sovereignchainpessimisticFilterer, error) {
	contract, err := bindGlobalexitrootmanagerl2sovereignchainpessimistic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainpessimisticFilterer{contract: contract}, nil
}

// bindGlobalexitrootmanagerl2sovereignchainpessimistic binds a generic wrapper to an already deployed contract.
func bindGlobalexitrootmanagerl2sovereignchainpessimistic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Globalexitrootmanagerl2sovereignchainpessimisticMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.Globalexitrootmanagerl2sovereignchainpessimisticCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.Globalexitrootmanagerl2sovereignchainpessimisticTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.Globalexitrootmanagerl2sovereignchainpessimisticTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Globalexitrootmanagerl2sovereignchainpessimistic.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticSession) BridgeAddress() (common.Address, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.BridgeAddress(&_Globalexitrootmanagerl2sovereignchainpessimistic.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticCallerSession) BridgeAddress() (common.Address, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.BridgeAddress(&_Globalexitrootmanagerl2sovereignchainpessimistic.CallOpts)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticCaller) GlobalExitRootMap(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Globalexitrootmanagerl2sovereignchainpessimistic.contract.Call(opts, &out, "globalExitRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.GlobalExitRootMap(&_Globalexitrootmanagerl2sovereignchainpessimistic.CallOpts, arg0)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticCallerSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.GlobalExitRootMap(&_Globalexitrootmanagerl2sovereignchainpessimistic.CallOpts, arg0)
}

// GlobalExitRootRemover is a free data retrieval call binding the contract method 0x91eb796d.
//
// Solidity: function globalExitRootRemover() view returns(address)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticCaller) GlobalExitRootRemover(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Globalexitrootmanagerl2sovereignchainpessimistic.contract.Call(opts, &out, "globalExitRootRemover")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootRemover is a free data retrieval call binding the contract method 0x91eb796d.
//
// Solidity: function globalExitRootRemover() view returns(address)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticSession) GlobalExitRootRemover() (common.Address, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.GlobalExitRootRemover(&_Globalexitrootmanagerl2sovereignchainpessimistic.CallOpts)
}

// GlobalExitRootRemover is a free data retrieval call binding the contract method 0x91eb796d.
//
// Solidity: function globalExitRootRemover() view returns(address)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticCallerSession) GlobalExitRootRemover() (common.Address, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.GlobalExitRootRemover(&_Globalexitrootmanagerl2sovereignchainpessimistic.CallOpts)
}

// GlobalExitRootUpdater is a free data retrieval call binding the contract method 0x7c314ce3.
//
// Solidity: function globalExitRootUpdater() view returns(address)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticCaller) GlobalExitRootUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Globalexitrootmanagerl2sovereignchainpessimistic.contract.Call(opts, &out, "globalExitRootUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootUpdater is a free data retrieval call binding the contract method 0x7c314ce3.
//
// Solidity: function globalExitRootUpdater() view returns(address)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticSession) GlobalExitRootUpdater() (common.Address, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.GlobalExitRootUpdater(&_Globalexitrootmanagerl2sovereignchainpessimistic.CallOpts)
}

// GlobalExitRootUpdater is a free data retrieval call binding the contract method 0x7c314ce3.
//
// Solidity: function globalExitRootUpdater() view returns(address)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticCallerSession) GlobalExitRootUpdater() (common.Address, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.GlobalExitRootUpdater(&_Globalexitrootmanagerl2sovereignchainpessimistic.CallOpts)
}

// InsertedGERCount is a free data retrieval call binding the contract method 0x8bd0eb1c.
//
// Solidity: function insertedGERCount() view returns(uint256)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticCaller) InsertedGERCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Globalexitrootmanagerl2sovereignchainpessimistic.contract.Call(opts, &out, "insertedGERCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InsertedGERCount is a free data retrieval call binding the contract method 0x8bd0eb1c.
//
// Solidity: function insertedGERCount() view returns(uint256)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticSession) InsertedGERCount() (*big.Int, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.InsertedGERCount(&_Globalexitrootmanagerl2sovereignchainpessimistic.CallOpts)
}

// InsertedGERCount is a free data retrieval call binding the contract method 0x8bd0eb1c.
//
// Solidity: function insertedGERCount() view returns(uint256)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticCallerSession) InsertedGERCount() (*big.Int, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.InsertedGERCount(&_Globalexitrootmanagerl2sovereignchainpessimistic.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticCaller) LastRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Globalexitrootmanagerl2sovereignchainpessimistic.contract.Call(opts, &out, "lastRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticSession) LastRollupExitRoot() ([32]byte, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.LastRollupExitRoot(&_Globalexitrootmanagerl2sovereignchainpessimistic.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticCallerSession) LastRollupExitRoot() ([32]byte, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.LastRollupExitRoot(&_Globalexitrootmanagerl2sovereignchainpessimistic.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _globalExitRootUpdater, address _globalExitRootRemover) returns()
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticTransactor) Initialize(opts *bind.TransactOpts, _globalExitRootUpdater common.Address, _globalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.contract.Transact(opts, "initialize", _globalExitRootUpdater, _globalExitRootRemover)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _globalExitRootUpdater, address _globalExitRootRemover) returns()
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticSession) Initialize(_globalExitRootUpdater common.Address, _globalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.Initialize(&_Globalexitrootmanagerl2sovereignchainpessimistic.TransactOpts, _globalExitRootUpdater, _globalExitRootRemover)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _globalExitRootUpdater, address _globalExitRootRemover) returns()
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticTransactorSession) Initialize(_globalExitRootUpdater common.Address, _globalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.Initialize(&_Globalexitrootmanagerl2sovereignchainpessimistic.TransactOpts, _globalExitRootUpdater, _globalExitRootRemover)
}

// InsertGlobalExitRoot is a paid mutator transaction binding the contract method 0x12da06b2.
//
// Solidity: function insertGlobalExitRoot(bytes32 _newRoot) returns()
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticTransactor) InsertGlobalExitRoot(opts *bind.TransactOpts, _newRoot [32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.contract.Transact(opts, "insertGlobalExitRoot", _newRoot)
}

// InsertGlobalExitRoot is a paid mutator transaction binding the contract method 0x12da06b2.
//
// Solidity: function insertGlobalExitRoot(bytes32 _newRoot) returns()
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticSession) InsertGlobalExitRoot(_newRoot [32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.InsertGlobalExitRoot(&_Globalexitrootmanagerl2sovereignchainpessimistic.TransactOpts, _newRoot)
}

// InsertGlobalExitRoot is a paid mutator transaction binding the contract method 0x12da06b2.
//
// Solidity: function insertGlobalExitRoot(bytes32 _newRoot) returns()
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticTransactorSession) InsertGlobalExitRoot(_newRoot [32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.InsertGlobalExitRoot(&_Globalexitrootmanagerl2sovereignchainpessimistic.TransactOpts, _newRoot)
}

// RemoveLastGlobalExitRoots is a paid mutator transaction binding the contract method 0x57dfb572.
//
// Solidity: function removeLastGlobalExitRoots(bytes32[] gersToRemove) returns()
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticTransactor) RemoveLastGlobalExitRoots(opts *bind.TransactOpts, gersToRemove [][32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.contract.Transact(opts, "removeLastGlobalExitRoots", gersToRemove)
}

// RemoveLastGlobalExitRoots is a paid mutator transaction binding the contract method 0x57dfb572.
//
// Solidity: function removeLastGlobalExitRoots(bytes32[] gersToRemove) returns()
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticSession) RemoveLastGlobalExitRoots(gersToRemove [][32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.RemoveLastGlobalExitRoots(&_Globalexitrootmanagerl2sovereignchainpessimistic.TransactOpts, gersToRemove)
}

// RemoveLastGlobalExitRoots is a paid mutator transaction binding the contract method 0x57dfb572.
//
// Solidity: function removeLastGlobalExitRoots(bytes32[] gersToRemove) returns()
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticTransactorSession) RemoveLastGlobalExitRoots(gersToRemove [][32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.RemoveLastGlobalExitRoots(&_Globalexitrootmanagerl2sovereignchainpessimistic.TransactOpts, gersToRemove)
}

// SetGlobalExitRootRemover is a paid mutator transaction binding the contract method 0xd0267f39.
//
// Solidity: function setGlobalExitRootRemover(address _globalExitRootRemover) returns()
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticTransactor) SetGlobalExitRootRemover(opts *bind.TransactOpts, _globalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.contract.Transact(opts, "setGlobalExitRootRemover", _globalExitRootRemover)
}

// SetGlobalExitRootRemover is a paid mutator transaction binding the contract method 0xd0267f39.
//
// Solidity: function setGlobalExitRootRemover(address _globalExitRootRemover) returns()
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticSession) SetGlobalExitRootRemover(_globalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.SetGlobalExitRootRemover(&_Globalexitrootmanagerl2sovereignchainpessimistic.TransactOpts, _globalExitRootRemover)
}

// SetGlobalExitRootRemover is a paid mutator transaction binding the contract method 0xd0267f39.
//
// Solidity: function setGlobalExitRootRemover(address _globalExitRootRemover) returns()
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticTransactorSession) SetGlobalExitRootRemover(_globalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.SetGlobalExitRootRemover(&_Globalexitrootmanagerl2sovereignchainpessimistic.TransactOpts, _globalExitRootRemover)
}

// SetGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x6da0e4ab.
//
// Solidity: function setGlobalExitRootUpdater(address _globalExitRootUpdater) returns()
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticTransactor) SetGlobalExitRootUpdater(opts *bind.TransactOpts, _globalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.contract.Transact(opts, "setGlobalExitRootUpdater", _globalExitRootUpdater)
}

// SetGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x6da0e4ab.
//
// Solidity: function setGlobalExitRootUpdater(address _globalExitRootUpdater) returns()
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticSession) SetGlobalExitRootUpdater(_globalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.SetGlobalExitRootUpdater(&_Globalexitrootmanagerl2sovereignchainpessimistic.TransactOpts, _globalExitRootUpdater)
}

// SetGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x6da0e4ab.
//
// Solidity: function setGlobalExitRootUpdater(address _globalExitRootUpdater) returns()
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticTransactorSession) SetGlobalExitRootUpdater(_globalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.SetGlobalExitRootUpdater(&_Globalexitrootmanagerl2sovereignchainpessimistic.TransactOpts, _globalExitRootUpdater)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticTransactor) UpdateExitRoot(opts *bind.TransactOpts, newRoot [32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.contract.Transact(opts, "updateExitRoot", newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticSession) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.UpdateExitRoot(&_Globalexitrootmanagerl2sovereignchainpessimistic.TransactOpts, newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticTransactorSession) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchainpessimistic.Contract.UpdateExitRoot(&_Globalexitrootmanagerl2sovereignchainpessimistic.TransactOpts, newRoot)
}

// Globalexitrootmanagerl2sovereignchainpessimisticInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Globalexitrootmanagerl2sovereignchainpessimistic contract.
type Globalexitrootmanagerl2sovereignchainpessimisticInitializedIterator struct {
	Event *Globalexitrootmanagerl2sovereignchainpessimisticInitialized // Event containing the contract specifics and raw log

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
func (it *Globalexitrootmanagerl2sovereignchainpessimisticInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Globalexitrootmanagerl2sovereignchainpessimisticInitialized)
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
		it.Event = new(Globalexitrootmanagerl2sovereignchainpessimisticInitialized)
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
func (it *Globalexitrootmanagerl2sovereignchainpessimisticInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Globalexitrootmanagerl2sovereignchainpessimisticInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Globalexitrootmanagerl2sovereignchainpessimisticInitialized represents a Initialized event raised by the Globalexitrootmanagerl2sovereignchainpessimistic contract.
type Globalexitrootmanagerl2sovereignchainpessimisticInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticFilterer) FilterInitialized(opts *bind.FilterOpts) (*Globalexitrootmanagerl2sovereignchainpessimisticInitializedIterator, error) {

	logs, sub, err := _Globalexitrootmanagerl2sovereignchainpessimistic.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainpessimisticInitializedIterator{contract: _Globalexitrootmanagerl2sovereignchainpessimistic.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Globalexitrootmanagerl2sovereignchainpessimisticInitialized) (event.Subscription, error) {

	logs, sub, err := _Globalexitrootmanagerl2sovereignchainpessimistic.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Globalexitrootmanagerl2sovereignchainpessimisticInitialized)
				if err := _Globalexitrootmanagerl2sovereignchainpessimistic.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticFilterer) ParseInitialized(log types.Log) (*Globalexitrootmanagerl2sovereignchainpessimisticInitialized, error) {
	event := new(Globalexitrootmanagerl2sovereignchainpessimisticInitialized)
	if err := _Globalexitrootmanagerl2sovereignchainpessimistic.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Globalexitrootmanagerl2sovereignchainpessimisticInsertGlobalExitRootIterator is returned from FilterInsertGlobalExitRoot and is used to iterate over the raw logs and unpacked data for InsertGlobalExitRoot events raised by the Globalexitrootmanagerl2sovereignchainpessimistic contract.
type Globalexitrootmanagerl2sovereignchainpessimisticInsertGlobalExitRootIterator struct {
	Event *Globalexitrootmanagerl2sovereignchainpessimisticInsertGlobalExitRoot // Event containing the contract specifics and raw log

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
func (it *Globalexitrootmanagerl2sovereignchainpessimisticInsertGlobalExitRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Globalexitrootmanagerl2sovereignchainpessimisticInsertGlobalExitRoot)
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
		it.Event = new(Globalexitrootmanagerl2sovereignchainpessimisticInsertGlobalExitRoot)
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
func (it *Globalexitrootmanagerl2sovereignchainpessimisticInsertGlobalExitRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Globalexitrootmanagerl2sovereignchainpessimisticInsertGlobalExitRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Globalexitrootmanagerl2sovereignchainpessimisticInsertGlobalExitRoot represents a InsertGlobalExitRoot event raised by the Globalexitrootmanagerl2sovereignchainpessimistic contract.
type Globalexitrootmanagerl2sovereignchainpessimisticInsertGlobalExitRoot struct {
	NewGlobalExitRoot [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterInsertGlobalExitRoot is a free log retrieval operation binding the contract event 0xb1b866fe5fac68e8f1a4ab2520c7a6b493a954934bbd0f054bd91d6674a4c0d5.
//
// Solidity: event InsertGlobalExitRoot(bytes32 indexed newGlobalExitRoot)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticFilterer) FilterInsertGlobalExitRoot(opts *bind.FilterOpts, newGlobalExitRoot [][32]byte) (*Globalexitrootmanagerl2sovereignchainpessimisticInsertGlobalExitRootIterator, error) {

	var newGlobalExitRootRule []interface{}
	for _, newGlobalExitRootItem := range newGlobalExitRoot {
		newGlobalExitRootRule = append(newGlobalExitRootRule, newGlobalExitRootItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchainpessimistic.contract.FilterLogs(opts, "InsertGlobalExitRoot", newGlobalExitRootRule)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainpessimisticInsertGlobalExitRootIterator{contract: _Globalexitrootmanagerl2sovereignchainpessimistic.contract, event: "InsertGlobalExitRoot", logs: logs, sub: sub}, nil
}

// WatchInsertGlobalExitRoot is a free log subscription operation binding the contract event 0xb1b866fe5fac68e8f1a4ab2520c7a6b493a954934bbd0f054bd91d6674a4c0d5.
//
// Solidity: event InsertGlobalExitRoot(bytes32 indexed newGlobalExitRoot)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticFilterer) WatchInsertGlobalExitRoot(opts *bind.WatchOpts, sink chan<- *Globalexitrootmanagerl2sovereignchainpessimisticInsertGlobalExitRoot, newGlobalExitRoot [][32]byte) (event.Subscription, error) {

	var newGlobalExitRootRule []interface{}
	for _, newGlobalExitRootItem := range newGlobalExitRoot {
		newGlobalExitRootRule = append(newGlobalExitRootRule, newGlobalExitRootItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchainpessimistic.contract.WatchLogs(opts, "InsertGlobalExitRoot", newGlobalExitRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Globalexitrootmanagerl2sovereignchainpessimisticInsertGlobalExitRoot)
				if err := _Globalexitrootmanagerl2sovereignchainpessimistic.contract.UnpackLog(event, "InsertGlobalExitRoot", log); err != nil {
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

// ParseInsertGlobalExitRoot is a log parse operation binding the contract event 0xb1b866fe5fac68e8f1a4ab2520c7a6b493a954934bbd0f054bd91d6674a4c0d5.
//
// Solidity: event InsertGlobalExitRoot(bytes32 indexed newGlobalExitRoot)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticFilterer) ParseInsertGlobalExitRoot(log types.Log) (*Globalexitrootmanagerl2sovereignchainpessimisticInsertGlobalExitRoot, error) {
	event := new(Globalexitrootmanagerl2sovereignchainpessimisticInsertGlobalExitRoot)
	if err := _Globalexitrootmanagerl2sovereignchainpessimistic.contract.UnpackLog(event, "InsertGlobalExitRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Globalexitrootmanagerl2sovereignchainpessimisticRemoveLastGlobalExitRootIterator is returned from FilterRemoveLastGlobalExitRoot and is used to iterate over the raw logs and unpacked data for RemoveLastGlobalExitRoot events raised by the Globalexitrootmanagerl2sovereignchainpessimistic contract.
type Globalexitrootmanagerl2sovereignchainpessimisticRemoveLastGlobalExitRootIterator struct {
	Event *Globalexitrootmanagerl2sovereignchainpessimisticRemoveLastGlobalExitRoot // Event containing the contract specifics and raw log

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
func (it *Globalexitrootmanagerl2sovereignchainpessimisticRemoveLastGlobalExitRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Globalexitrootmanagerl2sovereignchainpessimisticRemoveLastGlobalExitRoot)
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
		it.Event = new(Globalexitrootmanagerl2sovereignchainpessimisticRemoveLastGlobalExitRoot)
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
func (it *Globalexitrootmanagerl2sovereignchainpessimisticRemoveLastGlobalExitRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Globalexitrootmanagerl2sovereignchainpessimisticRemoveLastGlobalExitRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Globalexitrootmanagerl2sovereignchainpessimisticRemoveLastGlobalExitRoot represents a RemoveLastGlobalExitRoot event raised by the Globalexitrootmanagerl2sovereignchainpessimistic contract.
type Globalexitrootmanagerl2sovereignchainpessimisticRemoveLastGlobalExitRoot struct {
	RemovedGlobalExitRoot [32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRemoveLastGlobalExitRoot is a free log retrieval operation binding the contract event 0x605764d0b65b62ecf05dc90f674a00a2e2531fabaf120fdde65790e407fcb7a2.
//
// Solidity: event RemoveLastGlobalExitRoot(bytes32 indexed removedGlobalExitRoot)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticFilterer) FilterRemoveLastGlobalExitRoot(opts *bind.FilterOpts, removedGlobalExitRoot [][32]byte) (*Globalexitrootmanagerl2sovereignchainpessimisticRemoveLastGlobalExitRootIterator, error) {

	var removedGlobalExitRootRule []interface{}
	for _, removedGlobalExitRootItem := range removedGlobalExitRoot {
		removedGlobalExitRootRule = append(removedGlobalExitRootRule, removedGlobalExitRootItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchainpessimistic.contract.FilterLogs(opts, "RemoveLastGlobalExitRoot", removedGlobalExitRootRule)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainpessimisticRemoveLastGlobalExitRootIterator{contract: _Globalexitrootmanagerl2sovereignchainpessimistic.contract, event: "RemoveLastGlobalExitRoot", logs: logs, sub: sub}, nil
}

// WatchRemoveLastGlobalExitRoot is a free log subscription operation binding the contract event 0x605764d0b65b62ecf05dc90f674a00a2e2531fabaf120fdde65790e407fcb7a2.
//
// Solidity: event RemoveLastGlobalExitRoot(bytes32 indexed removedGlobalExitRoot)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticFilterer) WatchRemoveLastGlobalExitRoot(opts *bind.WatchOpts, sink chan<- *Globalexitrootmanagerl2sovereignchainpessimisticRemoveLastGlobalExitRoot, removedGlobalExitRoot [][32]byte) (event.Subscription, error) {

	var removedGlobalExitRootRule []interface{}
	for _, removedGlobalExitRootItem := range removedGlobalExitRoot {
		removedGlobalExitRootRule = append(removedGlobalExitRootRule, removedGlobalExitRootItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchainpessimistic.contract.WatchLogs(opts, "RemoveLastGlobalExitRoot", removedGlobalExitRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Globalexitrootmanagerl2sovereignchainpessimisticRemoveLastGlobalExitRoot)
				if err := _Globalexitrootmanagerl2sovereignchainpessimistic.contract.UnpackLog(event, "RemoveLastGlobalExitRoot", log); err != nil {
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

// ParseRemoveLastGlobalExitRoot is a log parse operation binding the contract event 0x605764d0b65b62ecf05dc90f674a00a2e2531fabaf120fdde65790e407fcb7a2.
//
// Solidity: event RemoveLastGlobalExitRoot(bytes32 indexed removedGlobalExitRoot)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticFilterer) ParseRemoveLastGlobalExitRoot(log types.Log) (*Globalexitrootmanagerl2sovereignchainpessimisticRemoveLastGlobalExitRoot, error) {
	event := new(Globalexitrootmanagerl2sovereignchainpessimisticRemoveLastGlobalExitRoot)
	if err := _Globalexitrootmanagerl2sovereignchainpessimistic.contract.UnpackLog(event, "RemoveLastGlobalExitRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootRemoverIterator is returned from FilterSetGlobalExitRootRemover and is used to iterate over the raw logs and unpacked data for SetGlobalExitRootRemover events raised by the Globalexitrootmanagerl2sovereignchainpessimistic contract.
type Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootRemoverIterator struct {
	Event *Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootRemover // Event containing the contract specifics and raw log

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
func (it *Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootRemoverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootRemover)
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
		it.Event = new(Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootRemover)
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
func (it *Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootRemoverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootRemoverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootRemover represents a SetGlobalExitRootRemover event raised by the Globalexitrootmanagerl2sovereignchainpessimistic contract.
type Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootRemover struct {
	NewGlobalExitRootRemover common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetGlobalExitRootRemover is a free log retrieval operation binding the contract event 0x00b4672b6135d1dfbd4e9520e01abb14ea5eac645990b0d24dfda00ae999b758.
//
// Solidity: event SetGlobalExitRootRemover(address indexed newGlobalExitRootRemover)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticFilterer) FilterSetGlobalExitRootRemover(opts *bind.FilterOpts, newGlobalExitRootRemover []common.Address) (*Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootRemoverIterator, error) {

	var newGlobalExitRootRemoverRule []interface{}
	for _, newGlobalExitRootRemoverItem := range newGlobalExitRootRemover {
		newGlobalExitRootRemoverRule = append(newGlobalExitRootRemoverRule, newGlobalExitRootRemoverItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchainpessimistic.contract.FilterLogs(opts, "SetGlobalExitRootRemover", newGlobalExitRootRemoverRule)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootRemoverIterator{contract: _Globalexitrootmanagerl2sovereignchainpessimistic.contract, event: "SetGlobalExitRootRemover", logs: logs, sub: sub}, nil
}

// WatchSetGlobalExitRootRemover is a free log subscription operation binding the contract event 0x00b4672b6135d1dfbd4e9520e01abb14ea5eac645990b0d24dfda00ae999b758.
//
// Solidity: event SetGlobalExitRootRemover(address indexed newGlobalExitRootRemover)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticFilterer) WatchSetGlobalExitRootRemover(opts *bind.WatchOpts, sink chan<- *Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootRemover, newGlobalExitRootRemover []common.Address) (event.Subscription, error) {

	var newGlobalExitRootRemoverRule []interface{}
	for _, newGlobalExitRootRemoverItem := range newGlobalExitRootRemover {
		newGlobalExitRootRemoverRule = append(newGlobalExitRootRemoverRule, newGlobalExitRootRemoverItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchainpessimistic.contract.WatchLogs(opts, "SetGlobalExitRootRemover", newGlobalExitRootRemoverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootRemover)
				if err := _Globalexitrootmanagerl2sovereignchainpessimistic.contract.UnpackLog(event, "SetGlobalExitRootRemover", log); err != nil {
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

// ParseSetGlobalExitRootRemover is a log parse operation binding the contract event 0x00b4672b6135d1dfbd4e9520e01abb14ea5eac645990b0d24dfda00ae999b758.
//
// Solidity: event SetGlobalExitRootRemover(address indexed newGlobalExitRootRemover)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticFilterer) ParseSetGlobalExitRootRemover(log types.Log) (*Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootRemover, error) {
	event := new(Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootRemover)
	if err := _Globalexitrootmanagerl2sovereignchainpessimistic.contract.UnpackLog(event, "SetGlobalExitRootRemover", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootUpdaterIterator is returned from FilterSetGlobalExitRootUpdater and is used to iterate over the raw logs and unpacked data for SetGlobalExitRootUpdater events raised by the Globalexitrootmanagerl2sovereignchainpessimistic contract.
type Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootUpdaterIterator struct {
	Event *Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootUpdater // Event containing the contract specifics and raw log

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
func (it *Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootUpdaterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootUpdater)
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
		it.Event = new(Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootUpdater)
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
func (it *Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootUpdaterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootUpdaterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootUpdater represents a SetGlobalExitRootUpdater event raised by the Globalexitrootmanagerl2sovereignchainpessimistic contract.
type Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootUpdater struct {
	NewGlobalExitRootUpdater common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetGlobalExitRootUpdater is a free log retrieval operation binding the contract event 0x992b80814dbc3fba903486d81daddb07d1d5b20483742458c8b0540e3a37e37c.
//
// Solidity: event SetGlobalExitRootUpdater(address indexed newGlobalExitRootUpdater)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticFilterer) FilterSetGlobalExitRootUpdater(opts *bind.FilterOpts, newGlobalExitRootUpdater []common.Address) (*Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootUpdaterIterator, error) {

	var newGlobalExitRootUpdaterRule []interface{}
	for _, newGlobalExitRootUpdaterItem := range newGlobalExitRootUpdater {
		newGlobalExitRootUpdaterRule = append(newGlobalExitRootUpdaterRule, newGlobalExitRootUpdaterItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchainpessimistic.contract.FilterLogs(opts, "SetGlobalExitRootUpdater", newGlobalExitRootUpdaterRule)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootUpdaterIterator{contract: _Globalexitrootmanagerl2sovereignchainpessimistic.contract, event: "SetGlobalExitRootUpdater", logs: logs, sub: sub}, nil
}

// WatchSetGlobalExitRootUpdater is a free log subscription operation binding the contract event 0x992b80814dbc3fba903486d81daddb07d1d5b20483742458c8b0540e3a37e37c.
//
// Solidity: event SetGlobalExitRootUpdater(address indexed newGlobalExitRootUpdater)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticFilterer) WatchSetGlobalExitRootUpdater(opts *bind.WatchOpts, sink chan<- *Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootUpdater, newGlobalExitRootUpdater []common.Address) (event.Subscription, error) {

	var newGlobalExitRootUpdaterRule []interface{}
	for _, newGlobalExitRootUpdaterItem := range newGlobalExitRootUpdater {
		newGlobalExitRootUpdaterRule = append(newGlobalExitRootUpdaterRule, newGlobalExitRootUpdaterItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchainpessimistic.contract.WatchLogs(opts, "SetGlobalExitRootUpdater", newGlobalExitRootUpdaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootUpdater)
				if err := _Globalexitrootmanagerl2sovereignchainpessimistic.contract.UnpackLog(event, "SetGlobalExitRootUpdater", log); err != nil {
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

// ParseSetGlobalExitRootUpdater is a log parse operation binding the contract event 0x992b80814dbc3fba903486d81daddb07d1d5b20483742458c8b0540e3a37e37c.
//
// Solidity: event SetGlobalExitRootUpdater(address indexed newGlobalExitRootUpdater)
func (_Globalexitrootmanagerl2sovereignchainpessimistic *Globalexitrootmanagerl2sovereignchainpessimisticFilterer) ParseSetGlobalExitRootUpdater(log types.Log) (*Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootUpdater, error) {
	event := new(Globalexitrootmanagerl2sovereignchainpessimisticSetGlobalExitRootUpdater)
	if err := _Globalexitrootmanagerl2sovereignchainpessimistic.contract.UnpackLog(event, "SetGlobalExitRootUpdater", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

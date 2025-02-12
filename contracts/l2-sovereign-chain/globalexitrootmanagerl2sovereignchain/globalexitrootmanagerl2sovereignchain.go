// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package globalexitrootmanagerl2sovereignchain

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

// Globalexitrootmanagerl2sovereignchainMetaData contains all meta data concerning the Globalexitrootmanagerl2sovereignchain contract.
var Globalexitrootmanagerl2sovereignchainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"GlobalExitRootAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootRemover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootUpdater\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newHashChainValue\",\"type\":\"bytes32\"}],\"name\":\"InsertGlobalExitRoot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"removedGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newRemovalHashChainValue\",\"type\":\"bytes32\"}],\"name\":\"RemoveGlobalExitRoot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGlobalExitRootRemover\",\"type\":\"address\"}],\"name\":\"SetGlobalExitRootRemover\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGlobalExitRootUpdater\",\"type\":\"address\"}],\"name\":\"SetGlobalExitRootUpdater\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentHashChainValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRemovalHashChainValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootRemover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_globalExitRootUpdater\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_globalExitRootRemover\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_newRoot\",\"type\":\"bytes32\"}],\"name\":\"insertGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"gersToRemove\",\"type\":\"bytes32[]\"}],\"name\":\"removeGlobalExitRoots\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_globalExitRootRemover\",\"type\":\"address\"}],\"name\":\"setGlobalExitRootRemover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_globalExitRootUpdater\",\"type\":\"address\"}],\"name\":\"setGlobalExitRootUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f5ffd5b50604051610c35380380610c3583398101604081905261002e91610109565b6001600160a01b038116608052610043610049565b50610136565b603454610100900460ff16156100b55760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60345460ff9081161015610107576034805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b5f60208284031215610119575f5ffd5b81516001600160a01b038116811461012f575f5ffd5b9392505050565b608051610ae06101555f395f81816101fc01526103d20152610ae05ff3fe608060405234801561000f575f5ffd5b50600436106100da575f3560e01c8063485cc955116100885780637c314ce3116100635780637c314ce31461018c57806391eb796d146101d7578063a3c573eb146101f7578063d0267f391461021e575f5ffd5b8063485cc9551461015357806365f0e347146101665780636da0e4ab14610179575f5ffd5b80631f71e0da116100b85780631f71e0da14610118578063257b36321461012157806333d6247d14610140575f5ffd5b806301fd9044146100de57806312da06b2146100fa57806312e8abbf1461010f575b5f5ffd5b6100e760015481565b6040519081526020015b60405180910390f35b61010d61010836600461097c565b610231565b005b6100e760375481565b6100e760365481565b6100e761012f36600461097c565b5f6020819052908152604090205481565b61010d61014e36600461097c565b6103ba565b61010d6101613660046109bb565b61042e565b61010d6101743660046109ec565b61063a565b61010d610187366004610a5d565b610790565b6034546101b29062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100f1565b6035546101b29073ffffffffffffffffffffffffffffffffffffffff1681565b6101b27f000000000000000000000000000000000000000000000000000000000000000081565b61010d61022c366004610a5d565b6108be565b60345462010000900473ffffffffffffffffffffffffffffffffffffffff166102925741331461028d576040517fc758fc1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6102e9565b60345462010000900473ffffffffffffffffffffffffffffffffffffffff1633146102e9576040517fc758fc1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f818152602081905260408120549003610388575f8181526020818152604091829020429055603654825191820152908101829052606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052805160209091012060368190559082907f4818797817f28eec29f9612bc1a89632e565ec870d55c593b51905c6f8250490905f90a350565b6040517f1f97a58200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610429576040517fb49365dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600155565b603454610100900460ff161580801561044e5750603454600160ff909116105b806104685750303b158015610468575060345460ff166001145b6104f8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840160405180910390fd5b603480547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561055657603480547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b603480547fffffffffffffffffffff0000000000000000000000000000000000000000ffff166201000073ffffffffffffffffffffffffffffffffffffffff8681169190910291909117909155603580547fffffffffffffffffffffffff000000000000000000000000000000000000000016918416919091179055801561063557603480547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b60355473ffffffffffffffffffffffffffffffffffffffff16331461068b576040517fa34ddeb100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6037545f5b82811015610788575f8484838181106106ab576106ab610a7d565b9050602002013590505f5f8281526020019081526020015f20545f036106fd576040517ff4a66f9d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160208101859052908101829052606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815282825280516020918201205f85815291829052918120819055909450849183917f4349c9bb198dfc0cd4f9774f09f0fad9c69a642e73ab02de85e28ce1ec6ab74091a350600101610690565b506037555050565b60345462010000900473ffffffffffffffffffffffffffffffffffffffff166107f1574133146107ec576040517fc758fc1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610848565b60345462010000900473ffffffffffffffffffffffffffffffffffffffff163314610848576040517fc758fc1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603480547fffffffffffffffffffff0000000000000000000000000000000000000000ffff166201000073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040517f992b80814dbc3fba903486d81daddb07d1d5b20483742458c8b0540e3a37e37c905f90a250565b60355473ffffffffffffffffffffffffffffffffffffffff16331461090f576040517fa34ddeb100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517eb4672b6135d1dfbd4e9520e01abb14ea5eac645990b0d24dfda00ae999b758905f90a250565b5f6020828403121561098c575f5ffd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146109b6575f5ffd5b919050565b5f5f604083850312156109cc575f5ffd5b6109d583610993565b91506109e360208401610993565b90509250929050565b5f5f602083850312156109fd575f5ffd5b823567ffffffffffffffff811115610a13575f5ffd5b8301601f81018513610a23575f5ffd5b803567ffffffffffffffff811115610a39575f5ffd5b8560208260051b8401011115610a4d575f5ffd5b6020919091019590945092505050565b5f60208284031215610a6d575f5ffd5b610a7682610993565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea264697066735822122053e8da55c999b5d386c5ba04914376582cac88659456c7fc6b74fb1febcce82964736f6c634300081c0033",
}

// Globalexitrootmanagerl2sovereignchainABI is the input ABI used to generate the binding from.
// Deprecated: Use Globalexitrootmanagerl2sovereignchainMetaData.ABI instead.
var Globalexitrootmanagerl2sovereignchainABI = Globalexitrootmanagerl2sovereignchainMetaData.ABI

// Globalexitrootmanagerl2sovereignchainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Globalexitrootmanagerl2sovereignchainMetaData.Bin instead.
var Globalexitrootmanagerl2sovereignchainBin = Globalexitrootmanagerl2sovereignchainMetaData.Bin

// DeployGlobalexitrootmanagerl2sovereignchain deploys a new Ethereum contract, binding an instance of Globalexitrootmanagerl2sovereignchain to it.
func DeployGlobalexitrootmanagerl2sovereignchain(auth *bind.TransactOpts, backend bind.ContractBackend, _bridgeAddress common.Address) (common.Address, *types.Transaction, *Globalexitrootmanagerl2sovereignchain, error) {
	parsed, err := Globalexitrootmanagerl2sovereignchainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Globalexitrootmanagerl2sovereignchainBin), backend, _bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Globalexitrootmanagerl2sovereignchain{Globalexitrootmanagerl2sovereignchainCaller: Globalexitrootmanagerl2sovereignchainCaller{contract: contract}, Globalexitrootmanagerl2sovereignchainTransactor: Globalexitrootmanagerl2sovereignchainTransactor{contract: contract}, Globalexitrootmanagerl2sovereignchainFilterer: Globalexitrootmanagerl2sovereignchainFilterer{contract: contract}}, nil
}

// Globalexitrootmanagerl2sovereignchain is an auto generated Go binding around an Ethereum contract.
type Globalexitrootmanagerl2sovereignchain struct {
	Globalexitrootmanagerl2sovereignchainCaller     // Read-only binding to the contract
	Globalexitrootmanagerl2sovereignchainTransactor // Write-only binding to the contract
	Globalexitrootmanagerl2sovereignchainFilterer   // Log filterer for contract events
}

// Globalexitrootmanagerl2sovereignchainCaller is an auto generated read-only Go binding around an Ethereum contract.
type Globalexitrootmanagerl2sovereignchainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Globalexitrootmanagerl2sovereignchainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Globalexitrootmanagerl2sovereignchainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Globalexitrootmanagerl2sovereignchainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Globalexitrootmanagerl2sovereignchainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Globalexitrootmanagerl2sovereignchainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Globalexitrootmanagerl2sovereignchainSession struct {
	Contract     *Globalexitrootmanagerl2sovereignchain // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                          // Call options to use throughout this session
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// Globalexitrootmanagerl2sovereignchainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Globalexitrootmanagerl2sovereignchainCallerSession struct {
	Contract *Globalexitrootmanagerl2sovereignchainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                // Call options to use throughout this session
}

// Globalexitrootmanagerl2sovereignchainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Globalexitrootmanagerl2sovereignchainTransactorSession struct {
	Contract     *Globalexitrootmanagerl2sovereignchainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                // Transaction auth options to use throughout this session
}

// Globalexitrootmanagerl2sovereignchainRaw is an auto generated low-level Go binding around an Ethereum contract.
type Globalexitrootmanagerl2sovereignchainRaw struct {
	Contract *Globalexitrootmanagerl2sovereignchain // Generic contract binding to access the raw methods on
}

// Globalexitrootmanagerl2sovereignchainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Globalexitrootmanagerl2sovereignchainCallerRaw struct {
	Contract *Globalexitrootmanagerl2sovereignchainCaller // Generic read-only contract binding to access the raw methods on
}

// Globalexitrootmanagerl2sovereignchainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Globalexitrootmanagerl2sovereignchainTransactorRaw struct {
	Contract *Globalexitrootmanagerl2sovereignchainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlobalexitrootmanagerl2sovereignchain creates a new instance of Globalexitrootmanagerl2sovereignchain, bound to a specific deployed contract.
func NewGlobalexitrootmanagerl2sovereignchain(address common.Address, backend bind.ContractBackend) (*Globalexitrootmanagerl2sovereignchain, error) {
	contract, err := bindGlobalexitrootmanagerl2sovereignchain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchain{Globalexitrootmanagerl2sovereignchainCaller: Globalexitrootmanagerl2sovereignchainCaller{contract: contract}, Globalexitrootmanagerl2sovereignchainTransactor: Globalexitrootmanagerl2sovereignchainTransactor{contract: contract}, Globalexitrootmanagerl2sovereignchainFilterer: Globalexitrootmanagerl2sovereignchainFilterer{contract: contract}}, nil
}

// NewGlobalexitrootmanagerl2sovereignchainCaller creates a new read-only instance of Globalexitrootmanagerl2sovereignchain, bound to a specific deployed contract.
func NewGlobalexitrootmanagerl2sovereignchainCaller(address common.Address, caller bind.ContractCaller) (*Globalexitrootmanagerl2sovereignchainCaller, error) {
	contract, err := bindGlobalexitrootmanagerl2sovereignchain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainCaller{contract: contract}, nil
}

// NewGlobalexitrootmanagerl2sovereignchainTransactor creates a new write-only instance of Globalexitrootmanagerl2sovereignchain, bound to a specific deployed contract.
func NewGlobalexitrootmanagerl2sovereignchainTransactor(address common.Address, transactor bind.ContractTransactor) (*Globalexitrootmanagerl2sovereignchainTransactor, error) {
	contract, err := bindGlobalexitrootmanagerl2sovereignchain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainTransactor{contract: contract}, nil
}

// NewGlobalexitrootmanagerl2sovereignchainFilterer creates a new log filterer instance of Globalexitrootmanagerl2sovereignchain, bound to a specific deployed contract.
func NewGlobalexitrootmanagerl2sovereignchainFilterer(address common.Address, filterer bind.ContractFilterer) (*Globalexitrootmanagerl2sovereignchainFilterer, error) {
	contract, err := bindGlobalexitrootmanagerl2sovereignchain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainFilterer{contract: contract}, nil
}

// bindGlobalexitrootmanagerl2sovereignchain binds a generic wrapper to an already deployed contract.
func bindGlobalexitrootmanagerl2sovereignchain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Globalexitrootmanagerl2sovereignchainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Globalexitrootmanagerl2sovereignchain.Contract.Globalexitrootmanagerl2sovereignchainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.Globalexitrootmanagerl2sovereignchainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.Globalexitrootmanagerl2sovereignchainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Globalexitrootmanagerl2sovereignchain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Globalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) BridgeAddress() (common.Address, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.BridgeAddress(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCallerSession) BridgeAddress() (common.Address, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.BridgeAddress(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// CurrentHashChainValue is a free data retrieval call binding the contract method 0x1f71e0da.
//
// Solidity: function currentHashChainValue() view returns(bytes32)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCaller) CurrentHashChainValue(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Globalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "currentHashChainValue")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentHashChainValue is a free data retrieval call binding the contract method 0x1f71e0da.
//
// Solidity: function currentHashChainValue() view returns(bytes32)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) CurrentHashChainValue() ([32]byte, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.CurrentHashChainValue(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// CurrentHashChainValue is a free data retrieval call binding the contract method 0x1f71e0da.
//
// Solidity: function currentHashChainValue() view returns(bytes32)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCallerSession) CurrentHashChainValue() ([32]byte, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.CurrentHashChainValue(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// CurrentRemovalHashChainValue is a free data retrieval call binding the contract method 0x12e8abbf.
//
// Solidity: function currentRemovalHashChainValue() view returns(bytes32)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCaller) CurrentRemovalHashChainValue(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Globalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "currentRemovalHashChainValue")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentRemovalHashChainValue is a free data retrieval call binding the contract method 0x12e8abbf.
//
// Solidity: function currentRemovalHashChainValue() view returns(bytes32)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) CurrentRemovalHashChainValue() ([32]byte, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.CurrentRemovalHashChainValue(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// CurrentRemovalHashChainValue is a free data retrieval call binding the contract method 0x12e8abbf.
//
// Solidity: function currentRemovalHashChainValue() view returns(bytes32)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCallerSession) CurrentRemovalHashChainValue() ([32]byte, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.CurrentRemovalHashChainValue(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCaller) GlobalExitRootMap(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Globalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "globalExitRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.GlobalExitRootMap(&_Globalexitrootmanagerl2sovereignchain.CallOpts, arg0)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCallerSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.GlobalExitRootMap(&_Globalexitrootmanagerl2sovereignchain.CallOpts, arg0)
}

// GlobalExitRootRemover is a free data retrieval call binding the contract method 0x91eb796d.
//
// Solidity: function globalExitRootRemover() view returns(address)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCaller) GlobalExitRootRemover(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Globalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "globalExitRootRemover")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootRemover is a free data retrieval call binding the contract method 0x91eb796d.
//
// Solidity: function globalExitRootRemover() view returns(address)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) GlobalExitRootRemover() (common.Address, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.GlobalExitRootRemover(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// GlobalExitRootRemover is a free data retrieval call binding the contract method 0x91eb796d.
//
// Solidity: function globalExitRootRemover() view returns(address)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCallerSession) GlobalExitRootRemover() (common.Address, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.GlobalExitRootRemover(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// GlobalExitRootUpdater is a free data retrieval call binding the contract method 0x7c314ce3.
//
// Solidity: function globalExitRootUpdater() view returns(address)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCaller) GlobalExitRootUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Globalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "globalExitRootUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootUpdater is a free data retrieval call binding the contract method 0x7c314ce3.
//
// Solidity: function globalExitRootUpdater() view returns(address)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) GlobalExitRootUpdater() (common.Address, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.GlobalExitRootUpdater(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// GlobalExitRootUpdater is a free data retrieval call binding the contract method 0x7c314ce3.
//
// Solidity: function globalExitRootUpdater() view returns(address)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCallerSession) GlobalExitRootUpdater() (common.Address, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.GlobalExitRootUpdater(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCaller) LastRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Globalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "lastRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) LastRollupExitRoot() ([32]byte, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.LastRollupExitRoot(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCallerSession) LastRollupExitRoot() ([32]byte, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.LastRollupExitRoot(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _globalExitRootUpdater, address _globalExitRootRemover) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactor) Initialize(opts *bind.TransactOpts, _globalExitRootUpdater common.Address, _globalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.contract.Transact(opts, "initialize", _globalExitRootUpdater, _globalExitRootRemover)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _globalExitRootUpdater, address _globalExitRootRemover) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) Initialize(_globalExitRootUpdater common.Address, _globalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.Initialize(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, _globalExitRootUpdater, _globalExitRootRemover)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _globalExitRootUpdater, address _globalExitRootRemover) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactorSession) Initialize(_globalExitRootUpdater common.Address, _globalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.Initialize(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, _globalExitRootUpdater, _globalExitRootRemover)
}

// InsertGlobalExitRoot is a paid mutator transaction binding the contract method 0x12da06b2.
//
// Solidity: function insertGlobalExitRoot(bytes32 _newRoot) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactor) InsertGlobalExitRoot(opts *bind.TransactOpts, _newRoot [32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.contract.Transact(opts, "insertGlobalExitRoot", _newRoot)
}

// InsertGlobalExitRoot is a paid mutator transaction binding the contract method 0x12da06b2.
//
// Solidity: function insertGlobalExitRoot(bytes32 _newRoot) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) InsertGlobalExitRoot(_newRoot [32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.InsertGlobalExitRoot(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, _newRoot)
}

// InsertGlobalExitRoot is a paid mutator transaction binding the contract method 0x12da06b2.
//
// Solidity: function insertGlobalExitRoot(bytes32 _newRoot) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactorSession) InsertGlobalExitRoot(_newRoot [32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.InsertGlobalExitRoot(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, _newRoot)
}

// RemoveGlobalExitRoots is a paid mutator transaction binding the contract method 0x65f0e347.
//
// Solidity: function removeGlobalExitRoots(bytes32[] gersToRemove) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactor) RemoveGlobalExitRoots(opts *bind.TransactOpts, gersToRemove [][32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.contract.Transact(opts, "removeGlobalExitRoots", gersToRemove)
}

// RemoveGlobalExitRoots is a paid mutator transaction binding the contract method 0x65f0e347.
//
// Solidity: function removeGlobalExitRoots(bytes32[] gersToRemove) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) RemoveGlobalExitRoots(gersToRemove [][32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.RemoveGlobalExitRoots(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, gersToRemove)
}

// RemoveGlobalExitRoots is a paid mutator transaction binding the contract method 0x65f0e347.
//
// Solidity: function removeGlobalExitRoots(bytes32[] gersToRemove) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactorSession) RemoveGlobalExitRoots(gersToRemove [][32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.RemoveGlobalExitRoots(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, gersToRemove)
}

// SetGlobalExitRootRemover is a paid mutator transaction binding the contract method 0xd0267f39.
//
// Solidity: function setGlobalExitRootRemover(address _globalExitRootRemover) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactor) SetGlobalExitRootRemover(opts *bind.TransactOpts, _globalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.contract.Transact(opts, "setGlobalExitRootRemover", _globalExitRootRemover)
}

// SetGlobalExitRootRemover is a paid mutator transaction binding the contract method 0xd0267f39.
//
// Solidity: function setGlobalExitRootRemover(address _globalExitRootRemover) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) SetGlobalExitRootRemover(_globalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.SetGlobalExitRootRemover(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, _globalExitRootRemover)
}

// SetGlobalExitRootRemover is a paid mutator transaction binding the contract method 0xd0267f39.
//
// Solidity: function setGlobalExitRootRemover(address _globalExitRootRemover) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactorSession) SetGlobalExitRootRemover(_globalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.SetGlobalExitRootRemover(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, _globalExitRootRemover)
}

// SetGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x6da0e4ab.
//
// Solidity: function setGlobalExitRootUpdater(address _globalExitRootUpdater) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactor) SetGlobalExitRootUpdater(opts *bind.TransactOpts, _globalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.contract.Transact(opts, "setGlobalExitRootUpdater", _globalExitRootUpdater)
}

// SetGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x6da0e4ab.
//
// Solidity: function setGlobalExitRootUpdater(address _globalExitRootUpdater) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) SetGlobalExitRootUpdater(_globalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.SetGlobalExitRootUpdater(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, _globalExitRootUpdater)
}

// SetGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x6da0e4ab.
//
// Solidity: function setGlobalExitRootUpdater(address _globalExitRootUpdater) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactorSession) SetGlobalExitRootUpdater(_globalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.SetGlobalExitRootUpdater(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, _globalExitRootUpdater)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactor) UpdateExitRoot(opts *bind.TransactOpts, newRoot [32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.contract.Transact(opts, "updateExitRoot", newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.UpdateExitRoot(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactorSession) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.UpdateExitRoot(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, newRoot)
}

// Globalexitrootmanagerl2sovereignchainInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainInitializedIterator struct {
	Event *Globalexitrootmanagerl2sovereignchainInitialized // Event containing the contract specifics and raw log

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
func (it *Globalexitrootmanagerl2sovereignchainInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Globalexitrootmanagerl2sovereignchainInitialized)
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
		it.Event = new(Globalexitrootmanagerl2sovereignchainInitialized)
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
func (it *Globalexitrootmanagerl2sovereignchainInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Globalexitrootmanagerl2sovereignchainInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Globalexitrootmanagerl2sovereignchainInitialized represents a Initialized event raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) FilterInitialized(opts *bind.FilterOpts) (*Globalexitrootmanagerl2sovereignchainInitializedIterator, error) {

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainInitializedIterator{contract: _Globalexitrootmanagerl2sovereignchain.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Globalexitrootmanagerl2sovereignchainInitialized) (event.Subscription, error) {

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Globalexitrootmanagerl2sovereignchainInitialized)
				if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) ParseInitialized(log types.Log) (*Globalexitrootmanagerl2sovereignchainInitialized, error) {
	event := new(Globalexitrootmanagerl2sovereignchainInitialized)
	if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Globalexitrootmanagerl2sovereignchainInsertGlobalExitRootIterator is returned from FilterInsertGlobalExitRoot and is used to iterate over the raw logs and unpacked data for InsertGlobalExitRoot events raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainInsertGlobalExitRootIterator struct {
	Event *Globalexitrootmanagerl2sovereignchainInsertGlobalExitRoot // Event containing the contract specifics and raw log

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
func (it *Globalexitrootmanagerl2sovereignchainInsertGlobalExitRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Globalexitrootmanagerl2sovereignchainInsertGlobalExitRoot)
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
		it.Event = new(Globalexitrootmanagerl2sovereignchainInsertGlobalExitRoot)
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
func (it *Globalexitrootmanagerl2sovereignchainInsertGlobalExitRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Globalexitrootmanagerl2sovereignchainInsertGlobalExitRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Globalexitrootmanagerl2sovereignchainInsertGlobalExitRoot represents a InsertGlobalExitRoot event raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainInsertGlobalExitRoot struct {
	NewGlobalExitRoot [32]byte
	NewHashChainValue [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterInsertGlobalExitRoot is a free log retrieval operation binding the contract event 0x4818797817f28eec29f9612bc1a89632e565ec870d55c593b51905c6f8250490.
//
// Solidity: event InsertGlobalExitRoot(bytes32 indexed newGlobalExitRoot, bytes32 indexed newHashChainValue)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) FilterInsertGlobalExitRoot(opts *bind.FilterOpts, newGlobalExitRoot [][32]byte, newHashChainValue [][32]byte) (*Globalexitrootmanagerl2sovereignchainInsertGlobalExitRootIterator, error) {

	var newGlobalExitRootRule []interface{}
	for _, newGlobalExitRootItem := range newGlobalExitRoot {
		newGlobalExitRootRule = append(newGlobalExitRootRule, newGlobalExitRootItem)
	}
	var newHashChainValueRule []interface{}
	for _, newHashChainValueItem := range newHashChainValue {
		newHashChainValueRule = append(newHashChainValueRule, newHashChainValueItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.FilterLogs(opts, "InsertGlobalExitRoot", newGlobalExitRootRule, newHashChainValueRule)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainInsertGlobalExitRootIterator{contract: _Globalexitrootmanagerl2sovereignchain.contract, event: "InsertGlobalExitRoot", logs: logs, sub: sub}, nil
}

// WatchInsertGlobalExitRoot is a free log subscription operation binding the contract event 0x4818797817f28eec29f9612bc1a89632e565ec870d55c593b51905c6f8250490.
//
// Solidity: event InsertGlobalExitRoot(bytes32 indexed newGlobalExitRoot, bytes32 indexed newHashChainValue)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) WatchInsertGlobalExitRoot(opts *bind.WatchOpts, sink chan<- *Globalexitrootmanagerl2sovereignchainInsertGlobalExitRoot, newGlobalExitRoot [][32]byte, newHashChainValue [][32]byte) (event.Subscription, error) {

	var newGlobalExitRootRule []interface{}
	for _, newGlobalExitRootItem := range newGlobalExitRoot {
		newGlobalExitRootRule = append(newGlobalExitRootRule, newGlobalExitRootItem)
	}
	var newHashChainValueRule []interface{}
	for _, newHashChainValueItem := range newHashChainValue {
		newHashChainValueRule = append(newHashChainValueRule, newHashChainValueItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.WatchLogs(opts, "InsertGlobalExitRoot", newGlobalExitRootRule, newHashChainValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Globalexitrootmanagerl2sovereignchainInsertGlobalExitRoot)
				if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "InsertGlobalExitRoot", log); err != nil {
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

// ParseInsertGlobalExitRoot is a log parse operation binding the contract event 0x4818797817f28eec29f9612bc1a89632e565ec870d55c593b51905c6f8250490.
//
// Solidity: event InsertGlobalExitRoot(bytes32 indexed newGlobalExitRoot, bytes32 indexed newHashChainValue)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) ParseInsertGlobalExitRoot(log types.Log) (*Globalexitrootmanagerl2sovereignchainInsertGlobalExitRoot, error) {
	event := new(Globalexitrootmanagerl2sovereignchainInsertGlobalExitRoot)
	if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "InsertGlobalExitRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Globalexitrootmanagerl2sovereignchainRemoveGlobalExitRootIterator is returned from FilterRemoveGlobalExitRoot and is used to iterate over the raw logs and unpacked data for RemoveGlobalExitRoot events raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainRemoveGlobalExitRootIterator struct {
	Event *Globalexitrootmanagerl2sovereignchainRemoveGlobalExitRoot // Event containing the contract specifics and raw log

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
func (it *Globalexitrootmanagerl2sovereignchainRemoveGlobalExitRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Globalexitrootmanagerl2sovereignchainRemoveGlobalExitRoot)
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
		it.Event = new(Globalexitrootmanagerl2sovereignchainRemoveGlobalExitRoot)
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
func (it *Globalexitrootmanagerl2sovereignchainRemoveGlobalExitRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Globalexitrootmanagerl2sovereignchainRemoveGlobalExitRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Globalexitrootmanagerl2sovereignchainRemoveGlobalExitRoot represents a RemoveGlobalExitRoot event raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainRemoveGlobalExitRoot struct {
	RemovedGlobalExitRoot    [32]byte
	NewRemovalHashChainValue [32]byte
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterRemoveGlobalExitRoot is a free log retrieval operation binding the contract event 0x4349c9bb198dfc0cd4f9774f09f0fad9c69a642e73ab02de85e28ce1ec6ab740.
//
// Solidity: event RemoveGlobalExitRoot(bytes32 indexed removedGlobalExitRoot, bytes32 indexed newRemovalHashChainValue)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) FilterRemoveGlobalExitRoot(opts *bind.FilterOpts, removedGlobalExitRoot [][32]byte, newRemovalHashChainValue [][32]byte) (*Globalexitrootmanagerl2sovereignchainRemoveGlobalExitRootIterator, error) {

	var removedGlobalExitRootRule []interface{}
	for _, removedGlobalExitRootItem := range removedGlobalExitRoot {
		removedGlobalExitRootRule = append(removedGlobalExitRootRule, removedGlobalExitRootItem)
	}
	var newRemovalHashChainValueRule []interface{}
	for _, newRemovalHashChainValueItem := range newRemovalHashChainValue {
		newRemovalHashChainValueRule = append(newRemovalHashChainValueRule, newRemovalHashChainValueItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.FilterLogs(opts, "RemoveGlobalExitRoot", removedGlobalExitRootRule, newRemovalHashChainValueRule)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainRemoveGlobalExitRootIterator{contract: _Globalexitrootmanagerl2sovereignchain.contract, event: "RemoveGlobalExitRoot", logs: logs, sub: sub}, nil
}

// WatchRemoveGlobalExitRoot is a free log subscription operation binding the contract event 0x4349c9bb198dfc0cd4f9774f09f0fad9c69a642e73ab02de85e28ce1ec6ab740.
//
// Solidity: event RemoveGlobalExitRoot(bytes32 indexed removedGlobalExitRoot, bytes32 indexed newRemovalHashChainValue)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) WatchRemoveGlobalExitRoot(opts *bind.WatchOpts, sink chan<- *Globalexitrootmanagerl2sovereignchainRemoveGlobalExitRoot, removedGlobalExitRoot [][32]byte, newRemovalHashChainValue [][32]byte) (event.Subscription, error) {

	var removedGlobalExitRootRule []interface{}
	for _, removedGlobalExitRootItem := range removedGlobalExitRoot {
		removedGlobalExitRootRule = append(removedGlobalExitRootRule, removedGlobalExitRootItem)
	}
	var newRemovalHashChainValueRule []interface{}
	for _, newRemovalHashChainValueItem := range newRemovalHashChainValue {
		newRemovalHashChainValueRule = append(newRemovalHashChainValueRule, newRemovalHashChainValueItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.WatchLogs(opts, "RemoveGlobalExitRoot", removedGlobalExitRootRule, newRemovalHashChainValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Globalexitrootmanagerl2sovereignchainRemoveGlobalExitRoot)
				if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "RemoveGlobalExitRoot", log); err != nil {
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

// ParseRemoveGlobalExitRoot is a log parse operation binding the contract event 0x4349c9bb198dfc0cd4f9774f09f0fad9c69a642e73ab02de85e28ce1ec6ab740.
//
// Solidity: event RemoveGlobalExitRoot(bytes32 indexed removedGlobalExitRoot, bytes32 indexed newRemovalHashChainValue)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) ParseRemoveGlobalExitRoot(log types.Log) (*Globalexitrootmanagerl2sovereignchainRemoveGlobalExitRoot, error) {
	event := new(Globalexitrootmanagerl2sovereignchainRemoveGlobalExitRoot)
	if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "RemoveGlobalExitRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemoverIterator is returned from FilterSetGlobalExitRootRemover and is used to iterate over the raw logs and unpacked data for SetGlobalExitRootRemover events raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemoverIterator struct {
	Event *Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover // Event containing the contract specifics and raw log

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
func (it *Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemoverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover)
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
		it.Event = new(Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover)
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
func (it *Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemoverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemoverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover represents a SetGlobalExitRootRemover event raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover struct {
	NewGlobalExitRootRemover common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetGlobalExitRootRemover is a free log retrieval operation binding the contract event 0x00b4672b6135d1dfbd4e9520e01abb14ea5eac645990b0d24dfda00ae999b758.
//
// Solidity: event SetGlobalExitRootRemover(address indexed newGlobalExitRootRemover)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) FilterSetGlobalExitRootRemover(opts *bind.FilterOpts, newGlobalExitRootRemover []common.Address) (*Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemoverIterator, error) {

	var newGlobalExitRootRemoverRule []interface{}
	for _, newGlobalExitRootRemoverItem := range newGlobalExitRootRemover {
		newGlobalExitRootRemoverRule = append(newGlobalExitRootRemoverRule, newGlobalExitRootRemoverItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.FilterLogs(opts, "SetGlobalExitRootRemover", newGlobalExitRootRemoverRule)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemoverIterator{contract: _Globalexitrootmanagerl2sovereignchain.contract, event: "SetGlobalExitRootRemover", logs: logs, sub: sub}, nil
}

// WatchSetGlobalExitRootRemover is a free log subscription operation binding the contract event 0x00b4672b6135d1dfbd4e9520e01abb14ea5eac645990b0d24dfda00ae999b758.
//
// Solidity: event SetGlobalExitRootRemover(address indexed newGlobalExitRootRemover)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) WatchSetGlobalExitRootRemover(opts *bind.WatchOpts, sink chan<- *Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover, newGlobalExitRootRemover []common.Address) (event.Subscription, error) {

	var newGlobalExitRootRemoverRule []interface{}
	for _, newGlobalExitRootRemoverItem := range newGlobalExitRootRemover {
		newGlobalExitRootRemoverRule = append(newGlobalExitRootRemoverRule, newGlobalExitRootRemoverItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.WatchLogs(opts, "SetGlobalExitRootRemover", newGlobalExitRootRemoverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover)
				if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "SetGlobalExitRootRemover", log); err != nil {
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
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) ParseSetGlobalExitRootRemover(log types.Log) (*Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover, error) {
	event := new(Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover)
	if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "SetGlobalExitRootRemover", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdaterIterator is returned from FilterSetGlobalExitRootUpdater and is used to iterate over the raw logs and unpacked data for SetGlobalExitRootUpdater events raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdaterIterator struct {
	Event *Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater // Event containing the contract specifics and raw log

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
func (it *Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdaterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater)
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
		it.Event = new(Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater)
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
func (it *Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdaterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdaterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater represents a SetGlobalExitRootUpdater event raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater struct {
	NewGlobalExitRootUpdater common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetGlobalExitRootUpdater is a free log retrieval operation binding the contract event 0x992b80814dbc3fba903486d81daddb07d1d5b20483742458c8b0540e3a37e37c.
//
// Solidity: event SetGlobalExitRootUpdater(address indexed newGlobalExitRootUpdater)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) FilterSetGlobalExitRootUpdater(opts *bind.FilterOpts, newGlobalExitRootUpdater []common.Address) (*Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdaterIterator, error) {

	var newGlobalExitRootUpdaterRule []interface{}
	for _, newGlobalExitRootUpdaterItem := range newGlobalExitRootUpdater {
		newGlobalExitRootUpdaterRule = append(newGlobalExitRootUpdaterRule, newGlobalExitRootUpdaterItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.FilterLogs(opts, "SetGlobalExitRootUpdater", newGlobalExitRootUpdaterRule)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdaterIterator{contract: _Globalexitrootmanagerl2sovereignchain.contract, event: "SetGlobalExitRootUpdater", logs: logs, sub: sub}, nil
}

// WatchSetGlobalExitRootUpdater is a free log subscription operation binding the contract event 0x992b80814dbc3fba903486d81daddb07d1d5b20483742458c8b0540e3a37e37c.
//
// Solidity: event SetGlobalExitRootUpdater(address indexed newGlobalExitRootUpdater)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) WatchSetGlobalExitRootUpdater(opts *bind.WatchOpts, sink chan<- *Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater, newGlobalExitRootUpdater []common.Address) (event.Subscription, error) {

	var newGlobalExitRootUpdaterRule []interface{}
	for _, newGlobalExitRootUpdaterItem := range newGlobalExitRootUpdater {
		newGlobalExitRootUpdaterRule = append(newGlobalExitRootUpdaterRule, newGlobalExitRootUpdaterItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.WatchLogs(opts, "SetGlobalExitRootUpdater", newGlobalExitRootUpdaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater)
				if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "SetGlobalExitRootUpdater", log); err != nil {
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
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) ParseSetGlobalExitRootUpdater(log types.Log) (*Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater, error) {
	event := new(Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater)
	if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "SetGlobalExitRootUpdater", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

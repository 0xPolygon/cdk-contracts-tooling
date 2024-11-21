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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"GlobalExitRootAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughGlobalExitRootsInserted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotLastInsertedGlobalExitRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootUpdater\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newGlobalExitRoot\",\"type\":\"bytes32\"}],\"name\":\"InsertGlobalExitRoot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"removedGlobalExitRoot\",\"type\":\"bytes32\"}],\"name\":\"RemoveLastGlobalExitRoot\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_globalExitRootUpdater\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_newRoot\",\"type\":\"bytes32\"}],\"name\":\"insertGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"insertedGERCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"gersToRemove\",\"type\":\"bytes32[]\"}],\"name\":\"removeLastGlobalExitRoots\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b50604051610a1c380380610a1c83398101604081905261002e91610109565b6001600160a01b038116608052610043610049565b50610136565b603454610100900460ff16156100b55760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60345460ff9081161015610107576034805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b5f60208284031215610119575f80fd5b81516001600160a01b038116811461012f575f80fd5b9392505050565b6080516108c76101555f395f8181610172015261030c01526108c75ff3fe608060405234801561000f575f80fd5b506004361061009f575f3560e01c806357dfb572116100725780638bd0eb1c116100585780638bd0eb1c14610164578063a3c573eb1461016d578063c4d66de814610194575f80fd5b806357dfb572146101065780637c314ce314610119575f80fd5b806301fd9044146100a357806312da06b2146100bf578063257b3632146100d457806333d6247d146100f3575b5f80fd5b6100ac60015481565b6040519081526020015b60405180910390f35b6100d26100cd36600461070c565b6101a7565b005b6100ac6100e236600461070c565b5f6020819052908152604090205481565b6100d261010136600461070c565b6102f4565b6100d2610114366004610723565b610368565b60345461013f9062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100b6565b6100ac60355481565b61013f7f000000000000000000000000000000000000000000000000000000000000000081565b6100d26101a2366004610792565b610538565b60345462010000900473ffffffffffffffffffffffffffffffffffffffff1661020857413314610203576040517fc758fc1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61025f565b60345462010000900473ffffffffffffffffffffffffffffffffffffffff16331461025f576040517fc758fc1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8181526020819052604081205490036102c25760355f8154610281906107f9565b91829055505f8281526020819052604080822092909255905182917fb1b866fe5fac68e8f1a4ab2520c7a6b493a954934bbd0f054bd91d6674a4c0d591a250565b6040517f1f97a58200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610363576040517fb49365dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600155565b60345462010000900473ffffffffffffffffffffffffffffffffffffffff166103c9574133146103c4576040517fc758fc1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610420565b60345462010000900473ffffffffffffffffffffffffffffffffffffffff163314610420576040517fc758fc1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6035548082111561045d576040517f56feb4f500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b82811015610532575f84848381811061047a5761047a610830565b9050602002013590505f805f8381526020019081526020015f205490508381146104d0576040517fae765ff600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f82815260208190526040812081905560358054916104ee8361085d565b909155505060405182907f605764d0b65b62ecf05dc90f674a00a2e2531fabaf120fdde65790e407fcb7a2905f90a25050808061052a906107f9565b91505061045f565b50505050565b603454610100900460ff16158080156105585750603454600160ff909116105b806105725750303b158015610572575060345460ff166001145b610602576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840160405180910390fd5b603480547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561066057603480547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b603480547fffffffffffffffffffff0000000000000000000000000000000000000000ffff166201000073ffffffffffffffffffffffffffffffffffffffff851602179055801561070857603480547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b5f6020828403121561071c575f80fd5b5035919050565b5f8060208385031215610734575f80fd5b823567ffffffffffffffff8082111561074b575f80fd5b818501915085601f83011261075e575f80fd5b81358181111561076c575f80fd5b8660208260051b8501011115610780575f80fd5b60209290920196919550909350505050565b5f602082840312156107a2575f80fd5b813573ffffffffffffffffffffffffffffffffffffffff811681146107c5575f80fd5b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610829576108296107cc565b5060010190565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f8161086b5761086b6107cc565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019056fea26469706673582212208aa57b6d01fb3f5a888f9860a686b1b8ce93fe905c3eda5b60a6ea5984b0dd8164736f6c63430008140033",
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

// InsertedGERCount is a free data retrieval call binding the contract method 0x8bd0eb1c.
//
// Solidity: function insertedGERCount() view returns(uint256)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCaller) InsertedGERCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Globalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "insertedGERCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InsertedGERCount is a free data retrieval call binding the contract method 0x8bd0eb1c.
//
// Solidity: function insertedGERCount() view returns(uint256)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) InsertedGERCount() (*big.Int, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.InsertedGERCount(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// InsertedGERCount is a free data retrieval call binding the contract method 0x8bd0eb1c.
//
// Solidity: function insertedGERCount() view returns(uint256)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCallerSession) InsertedGERCount() (*big.Int, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.InsertedGERCount(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _globalExitRootUpdater) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactor) Initialize(opts *bind.TransactOpts, _globalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.contract.Transact(opts, "initialize", _globalExitRootUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _globalExitRootUpdater) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) Initialize(_globalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.Initialize(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, _globalExitRootUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _globalExitRootUpdater) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactorSession) Initialize(_globalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.Initialize(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, _globalExitRootUpdater)
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

// RemoveLastGlobalExitRoots is a paid mutator transaction binding the contract method 0x57dfb572.
//
// Solidity: function removeLastGlobalExitRoots(bytes32[] gersToRemove) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactor) RemoveLastGlobalExitRoots(opts *bind.TransactOpts, gersToRemove [][32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.contract.Transact(opts, "removeLastGlobalExitRoots", gersToRemove)
}

// RemoveLastGlobalExitRoots is a paid mutator transaction binding the contract method 0x57dfb572.
//
// Solidity: function removeLastGlobalExitRoots(bytes32[] gersToRemove) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) RemoveLastGlobalExitRoots(gersToRemove [][32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.RemoveLastGlobalExitRoots(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, gersToRemove)
}

// RemoveLastGlobalExitRoots is a paid mutator transaction binding the contract method 0x57dfb572.
//
// Solidity: function removeLastGlobalExitRoots(bytes32[] gersToRemove) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactorSession) RemoveLastGlobalExitRoots(gersToRemove [][32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.RemoveLastGlobalExitRoots(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, gersToRemove)
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
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterInsertGlobalExitRoot is a free log retrieval operation binding the contract event 0xb1b866fe5fac68e8f1a4ab2520c7a6b493a954934bbd0f054bd91d6674a4c0d5.
//
// Solidity: event InsertGlobalExitRoot(bytes32 indexed newGlobalExitRoot)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) FilterInsertGlobalExitRoot(opts *bind.FilterOpts, newGlobalExitRoot [][32]byte) (*Globalexitrootmanagerl2sovereignchainInsertGlobalExitRootIterator, error) {

	var newGlobalExitRootRule []interface{}
	for _, newGlobalExitRootItem := range newGlobalExitRoot {
		newGlobalExitRootRule = append(newGlobalExitRootRule, newGlobalExitRootItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.FilterLogs(opts, "InsertGlobalExitRoot", newGlobalExitRootRule)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainInsertGlobalExitRootIterator{contract: _Globalexitrootmanagerl2sovereignchain.contract, event: "InsertGlobalExitRoot", logs: logs, sub: sub}, nil
}

// WatchInsertGlobalExitRoot is a free log subscription operation binding the contract event 0xb1b866fe5fac68e8f1a4ab2520c7a6b493a954934bbd0f054bd91d6674a4c0d5.
//
// Solidity: event InsertGlobalExitRoot(bytes32 indexed newGlobalExitRoot)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) WatchInsertGlobalExitRoot(opts *bind.WatchOpts, sink chan<- *Globalexitrootmanagerl2sovereignchainInsertGlobalExitRoot, newGlobalExitRoot [][32]byte) (event.Subscription, error) {

	var newGlobalExitRootRule []interface{}
	for _, newGlobalExitRootItem := range newGlobalExitRoot {
		newGlobalExitRootRule = append(newGlobalExitRootRule, newGlobalExitRootItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.WatchLogs(opts, "InsertGlobalExitRoot", newGlobalExitRootRule)
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

// ParseInsertGlobalExitRoot is a log parse operation binding the contract event 0xb1b866fe5fac68e8f1a4ab2520c7a6b493a954934bbd0f054bd91d6674a4c0d5.
//
// Solidity: event InsertGlobalExitRoot(bytes32 indexed newGlobalExitRoot)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) ParseInsertGlobalExitRoot(log types.Log) (*Globalexitrootmanagerl2sovereignchainInsertGlobalExitRoot, error) {
	event := new(Globalexitrootmanagerl2sovereignchainInsertGlobalExitRoot)
	if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "InsertGlobalExitRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRootIterator is returned from FilterRemoveLastGlobalExitRoot and is used to iterate over the raw logs and unpacked data for RemoveLastGlobalExitRoot events raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRootIterator struct {
	Event *Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot // Event containing the contract specifics and raw log

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
func (it *Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot)
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
		it.Event = new(Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot)
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
func (it *Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot represents a RemoveLastGlobalExitRoot event raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot struct {
	RemovedGlobalExitRoot [32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRemoveLastGlobalExitRoot is a free log retrieval operation binding the contract event 0x605764d0b65b62ecf05dc90f674a00a2e2531fabaf120fdde65790e407fcb7a2.
//
// Solidity: event RemoveLastGlobalExitRoot(bytes32 indexed removedGlobalExitRoot)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) FilterRemoveLastGlobalExitRoot(opts *bind.FilterOpts, removedGlobalExitRoot [][32]byte) (*Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRootIterator, error) {

	var removedGlobalExitRootRule []interface{}
	for _, removedGlobalExitRootItem := range removedGlobalExitRoot {
		removedGlobalExitRootRule = append(removedGlobalExitRootRule, removedGlobalExitRootItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.FilterLogs(opts, "RemoveLastGlobalExitRoot", removedGlobalExitRootRule)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRootIterator{contract: _Globalexitrootmanagerl2sovereignchain.contract, event: "RemoveLastGlobalExitRoot", logs: logs, sub: sub}, nil
}

// WatchRemoveLastGlobalExitRoot is a free log subscription operation binding the contract event 0x605764d0b65b62ecf05dc90f674a00a2e2531fabaf120fdde65790e407fcb7a2.
//
// Solidity: event RemoveLastGlobalExitRoot(bytes32 indexed removedGlobalExitRoot)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) WatchRemoveLastGlobalExitRoot(opts *bind.WatchOpts, sink chan<- *Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot, removedGlobalExitRoot [][32]byte) (event.Subscription, error) {

	var removedGlobalExitRootRule []interface{}
	for _, removedGlobalExitRootItem := range removedGlobalExitRoot {
		removedGlobalExitRootRule = append(removedGlobalExitRootRule, removedGlobalExitRootItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.WatchLogs(opts, "RemoveLastGlobalExitRoot", removedGlobalExitRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot)
				if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "RemoveLastGlobalExitRoot", log); err != nil {
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
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) ParseRemoveLastGlobalExitRoot(log types.Log) (*Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot, error) {
	event := new(Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot)
	if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "RemoveLastGlobalExitRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

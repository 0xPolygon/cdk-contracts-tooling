// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package batchl2datacreatedrollup

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

// Batchl2datacreatedrollupMetaData contains all meta data concerning the Batchl2datacreatedrollup contract.
var Batchl2datacreatedrollupMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"GLOBAL_EXIT_ROOT_MANAGER_L2\",\"outputs\":[{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_LIST_LEN_LEN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_DATA_LEN_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_EFFECTIVE_PERCENTAGE\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_R\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_S\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_V\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"}],\"name\":\"generateInitializeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50610ae38061001c5f395ff3fe608060405234801561000f575f5ffd5b50600436106100da575f3560e01c80637a5460c511610088578063c23c541111610063578063c23c541114610293578063c7fffd4b146102a6578063d7bc90ff146102ae578063f35dda47146102b9575f5ffd5b80637a5460c5146101fd5780639e00187714610239578063b0afe15414610279575f5ffd5b806340b5de6c116100b857806340b5de6c1461016157806352bdeb6d146101b9578063676870d2146101f5575f5ffd5b806303508963146100de57806305835f37146100fe57806311e892d414610147575b5f5ffd5b6100e6602081565b60405161ffff90911681526020015b60405180910390f35b61013a6040518060400160405280600881526020017f80808401c9c3809400000000000000000000000000000000000000000000000081525081565b6040516100f5919061061b565b61014f60f981565b60405160ff90911681526020016100f5565b6101887fff0000000000000000000000000000000000000000000000000000000000000081565b6040517fff0000000000000000000000000000000000000000000000000000000000000090911681526020016100f5565b61013a6040518060400160405280600281526020017f80b800000000000000000000000000000000000000000000000000000000000081525081565b6100e6601f81565b61013a6040518060400160405280600281526020017f80b900000000000000000000000000000000000000000000000000000000000081525081565b61025473a40d5f56745a118d0906a34e69aec8c0db1cb8fa81565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100f5565b6102856405ca1ab1e081565b6040519081526020016100f5565b61013a6102a136600461069c565b6102c1565b61014f60e481565b610285635ca1ab1e81565b61014f601b81565b60605f86858573a40d5f56745a118d0906a34e69aec8c0db1cb8fa5f876040516024016102f3969594939291906107cd565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff811bff700000000000000000000000000000000000000000000000000000000179052805184519192506060915f036104225760f9610386601f83610857565b6040518060400160405280600881526020017f80808401c9c380940000000000000000000000000000000000000000000000008152508a6040518060400160405280600281526020017f80b800000000000000000000000000000000000000000000000000000000000081525060e48860405160200161040c97969594939291906108b3565b60405160208183030381529060405291506104c6565b60f961042f602083610857565b6040518060400160405280600881526020017f80808401c9c380940000000000000000000000000000000000000000000000008152508a6040518060400160405280600281526020017f80b900000000000000000000000000000000000000000000000000000000000081525085886040516020016104b4979695949392919061098e565b60405160208183030381529060405291505b8151602080840191909120604080515f80825293810180835292909252601b908201526405ca1ab1e06060820152635ca1ab1e608082015260019060a0016020604051602081039080840390855afa158015610524573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015192505f91506105929085906405ca1ab1e090635ca1ab1e90601b907fff0000000000000000000000000000000000000000000000000000000000000090602001610a5b565b604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe00181529190529b9a5050505050505050505050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f61062d60208301846105cf565b9392505050565b803563ffffffff81168114610647575f5ffd5b919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610647575f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f5f5f5f5f60a086880312156106b0575f5ffd5b6106b986610634565b94506106c76020870161064c565b93506106d56040870161064c565b92506106e360608701610634565b9150608086013567ffffffffffffffff8111156106fe575f5ffd5b8601601f8101881361070e575f5ffd5b803567ffffffffffffffff8111156107285761072861066f565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff821117156107945761079461066f565b6040528181528282016020018a10156107ab575f5ffd5b816020840160208301375f602083830101528093505050509295509295909350565b63ffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015263ffffffff8516604082015273ffffffffffffffffffffffffffffffffffffffff8416606082015273ffffffffffffffffffffffffffffffffffffffff8316608082015260c060a08201525f61084b60c08301846105cf565b98975050505050505050565b61ffff8181168382160190811115610896577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b92915050565b5f81518060208401855e5f93019283525090919050565b7fff000000000000000000000000000000000000000000000000000000000000008860f81b1681527fffff0000000000000000000000000000000000000000000000000000000000008760f01b1660018201525f610914600383018861089c565b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008760601b168152610949601482018761089c565b90507fff000000000000000000000000000000000000000000000000000000000000008560f81b168152610980600182018561089c565b9a9950505050505050505050565b7fff000000000000000000000000000000000000000000000000000000000000008860f81b1681527fffff0000000000000000000000000000000000000000000000000000000000008760f01b1660018201525f6109ef600383018861089c565b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008760601b168152610a24601482018761089c565b90507fffff0000000000000000000000000000000000000000000000000000000000008560f01b168152610980600282018561089c565b5f610a66828861089c565b9586525050602084019290925260f81b7fff00000000000000000000000000000000000000000000000000000000000000908116604084015216604182015260420191905056fea2646970667358221220140475e10834da495b509eef7bd788d0ad3ad7dcb8a44a610210fb7a1e6316d964736f6c634300081c0033",
}

// Batchl2datacreatedrollupABI is the input ABI used to generate the binding from.
// Deprecated: Use Batchl2datacreatedrollupMetaData.ABI instead.
var Batchl2datacreatedrollupABI = Batchl2datacreatedrollupMetaData.ABI

// Batchl2datacreatedrollupBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Batchl2datacreatedrollupMetaData.Bin instead.
var Batchl2datacreatedrollupBin = Batchl2datacreatedrollupMetaData.Bin

// DeployBatchl2datacreatedrollup deploys a new Ethereum contract, binding an instance of Batchl2datacreatedrollup to it.
func DeployBatchl2datacreatedrollup(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Batchl2datacreatedrollup, error) {
	parsed, err := Batchl2datacreatedrollupMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Batchl2datacreatedrollupBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Batchl2datacreatedrollup{Batchl2datacreatedrollupCaller: Batchl2datacreatedrollupCaller{contract: contract}, Batchl2datacreatedrollupTransactor: Batchl2datacreatedrollupTransactor{contract: contract}, Batchl2datacreatedrollupFilterer: Batchl2datacreatedrollupFilterer{contract: contract}}, nil
}

// Batchl2datacreatedrollup is an auto generated Go binding around an Ethereum contract.
type Batchl2datacreatedrollup struct {
	Batchl2datacreatedrollupCaller     // Read-only binding to the contract
	Batchl2datacreatedrollupTransactor // Write-only binding to the contract
	Batchl2datacreatedrollupFilterer   // Log filterer for contract events
}

// Batchl2datacreatedrollupCaller is an auto generated read-only Go binding around an Ethereum contract.
type Batchl2datacreatedrollupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Batchl2datacreatedrollupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Batchl2datacreatedrollupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Batchl2datacreatedrollupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Batchl2datacreatedrollupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Batchl2datacreatedrollupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Batchl2datacreatedrollupSession struct {
	Contract     *Batchl2datacreatedrollup // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Batchl2datacreatedrollupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Batchl2datacreatedrollupCallerSession struct {
	Contract *Batchl2datacreatedrollupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// Batchl2datacreatedrollupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Batchl2datacreatedrollupTransactorSession struct {
	Contract     *Batchl2datacreatedrollupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// Batchl2datacreatedrollupRaw is an auto generated low-level Go binding around an Ethereum contract.
type Batchl2datacreatedrollupRaw struct {
	Contract *Batchl2datacreatedrollup // Generic contract binding to access the raw methods on
}

// Batchl2datacreatedrollupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Batchl2datacreatedrollupCallerRaw struct {
	Contract *Batchl2datacreatedrollupCaller // Generic read-only contract binding to access the raw methods on
}

// Batchl2datacreatedrollupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Batchl2datacreatedrollupTransactorRaw struct {
	Contract *Batchl2datacreatedrollupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBatchl2datacreatedrollup creates a new instance of Batchl2datacreatedrollup, bound to a specific deployed contract.
func NewBatchl2datacreatedrollup(address common.Address, backend bind.ContractBackend) (*Batchl2datacreatedrollup, error) {
	contract, err := bindBatchl2datacreatedrollup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Batchl2datacreatedrollup{Batchl2datacreatedrollupCaller: Batchl2datacreatedrollupCaller{contract: contract}, Batchl2datacreatedrollupTransactor: Batchl2datacreatedrollupTransactor{contract: contract}, Batchl2datacreatedrollupFilterer: Batchl2datacreatedrollupFilterer{contract: contract}}, nil
}

// NewBatchl2datacreatedrollupCaller creates a new read-only instance of Batchl2datacreatedrollup, bound to a specific deployed contract.
func NewBatchl2datacreatedrollupCaller(address common.Address, caller bind.ContractCaller) (*Batchl2datacreatedrollupCaller, error) {
	contract, err := bindBatchl2datacreatedrollup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Batchl2datacreatedrollupCaller{contract: contract}, nil
}

// NewBatchl2datacreatedrollupTransactor creates a new write-only instance of Batchl2datacreatedrollup, bound to a specific deployed contract.
func NewBatchl2datacreatedrollupTransactor(address common.Address, transactor bind.ContractTransactor) (*Batchl2datacreatedrollupTransactor, error) {
	contract, err := bindBatchl2datacreatedrollup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Batchl2datacreatedrollupTransactor{contract: contract}, nil
}

// NewBatchl2datacreatedrollupFilterer creates a new log filterer instance of Batchl2datacreatedrollup, bound to a specific deployed contract.
func NewBatchl2datacreatedrollupFilterer(address common.Address, filterer bind.ContractFilterer) (*Batchl2datacreatedrollupFilterer, error) {
	contract, err := bindBatchl2datacreatedrollup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Batchl2datacreatedrollupFilterer{contract: contract}, nil
}

// bindBatchl2datacreatedrollup binds a generic wrapper to an already deployed contract.
func bindBatchl2datacreatedrollup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Batchl2datacreatedrollupMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Batchl2datacreatedrollup.Contract.Batchl2datacreatedrollupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Batchl2datacreatedrollup.Contract.Batchl2datacreatedrollupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Batchl2datacreatedrollup.Contract.Batchl2datacreatedrollupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Batchl2datacreatedrollup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Batchl2datacreatedrollup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Batchl2datacreatedrollup.Contract.contract.Transact(opts, method, params...)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCaller) GLOBALEXITROOTMANAGERL2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Batchl2datacreatedrollup.contract.Call(opts, &out, "GLOBAL_EXIT_ROOT_MANAGER_L2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _Batchl2datacreatedrollup.Contract.GLOBALEXITROOTMANAGERL2(&_Batchl2datacreatedrollup.CallOpts)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCallerSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _Batchl2datacreatedrollup.Contract.GLOBALEXITROOTMANAGERL2(&_Batchl2datacreatedrollup.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCaller) INITIALIZETXBRIDGELISTLENLEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Batchl2datacreatedrollup.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_LIST_LEN_LEN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _Batchl2datacreatedrollup.Contract.INITIALIZETXBRIDGELISTLENLEN(&_Batchl2datacreatedrollup.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCallerSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _Batchl2datacreatedrollup.Contract.INITIALIZETXBRIDGELISTLENLEN(&_Batchl2datacreatedrollup.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCaller) INITIALIZETXBRIDGEPARAMS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Batchl2datacreatedrollup.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _Batchl2datacreatedrollup.Contract.INITIALIZETXBRIDGEPARAMS(&_Batchl2datacreatedrollup.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCallerSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _Batchl2datacreatedrollup.Contract.INITIALIZETXBRIDGEPARAMS(&_Batchl2datacreatedrollup.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Batchl2datacreatedrollup.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _Batchl2datacreatedrollup.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_Batchl2datacreatedrollup.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _Batchl2datacreatedrollup.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_Batchl2datacreatedrollup.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Batchl2datacreatedrollup.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _Batchl2datacreatedrollup.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_Batchl2datacreatedrollup.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _Batchl2datacreatedrollup.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_Batchl2datacreatedrollup.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCaller) INITIALIZETXCONSTANTBYTES(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Batchl2datacreatedrollup.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _Batchl2datacreatedrollup.Contract.INITIALIZETXCONSTANTBYTES(&_Batchl2datacreatedrollup.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCallerSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _Batchl2datacreatedrollup.Contract.INITIALIZETXCONSTANTBYTES(&_Batchl2datacreatedrollup.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCaller) INITIALIZETXCONSTANTBYTESEMPTYMETADATA(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Batchl2datacreatedrollup.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _Batchl2datacreatedrollup.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_Batchl2datacreatedrollup.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCallerSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _Batchl2datacreatedrollup.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_Batchl2datacreatedrollup.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCaller) INITIALIZETXDATALENEMPTYMETADATA(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Batchl2datacreatedrollup.contract.Call(opts, &out, "INITIALIZE_TX_DATA_LEN_EMPTY_METADATA")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _Batchl2datacreatedrollup.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_Batchl2datacreatedrollup.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCallerSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _Batchl2datacreatedrollup.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_Batchl2datacreatedrollup.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCaller) INITIALIZETXEFFECTIVEPERCENTAGE(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _Batchl2datacreatedrollup.contract.Call(opts, &out, "INITIALIZE_TX_EFFECTIVE_PERCENTAGE")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _Batchl2datacreatedrollup.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_Batchl2datacreatedrollup.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCallerSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _Batchl2datacreatedrollup.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_Batchl2datacreatedrollup.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCaller) SIGNATUREINITIALIZETXR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Batchl2datacreatedrollup.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_R")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _Batchl2datacreatedrollup.Contract.SIGNATUREINITIALIZETXR(&_Batchl2datacreatedrollup.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCallerSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _Batchl2datacreatedrollup.Contract.SIGNATUREINITIALIZETXR(&_Batchl2datacreatedrollup.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCaller) SIGNATUREINITIALIZETXS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Batchl2datacreatedrollup.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_S")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _Batchl2datacreatedrollup.Contract.SIGNATUREINITIALIZETXS(&_Batchl2datacreatedrollup.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCallerSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _Batchl2datacreatedrollup.Contract.SIGNATUREINITIALIZETXS(&_Batchl2datacreatedrollup.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCaller) SIGNATUREINITIALIZETXV(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Batchl2datacreatedrollup.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_V")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _Batchl2datacreatedrollup.Contract.SIGNATUREINITIALIZETXV(&_Batchl2datacreatedrollup.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCallerSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _Batchl2datacreatedrollup.Contract.SIGNATUREINITIALIZETXV(&_Batchl2datacreatedrollup.CallOpts)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xc23c5411.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address bridgeAddress, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCaller) GenerateInitializeTransaction(opts *bind.CallOpts, networkID uint32, bridgeAddress common.Address, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	var out []interface{}
	err := _Batchl2datacreatedrollup.contract.Call(opts, &out, "generateInitializeTransaction", networkID, bridgeAddress, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xc23c5411.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address bridgeAddress, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupSession) GenerateInitializeTransaction(networkID uint32, bridgeAddress common.Address, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _Batchl2datacreatedrollup.Contract.GenerateInitializeTransaction(&_Batchl2datacreatedrollup.CallOpts, networkID, bridgeAddress, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xc23c5411.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address bridgeAddress, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Batchl2datacreatedrollup *Batchl2datacreatedrollupCallerSession) GenerateInitializeTransaction(networkID uint32, bridgeAddress common.Address, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _Batchl2datacreatedrollup.Contract.GenerateInitializeTransaction(&_Batchl2datacreatedrollup.CallOpts, networkID, bridgeAddress, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

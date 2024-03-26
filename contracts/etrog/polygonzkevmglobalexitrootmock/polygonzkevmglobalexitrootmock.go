// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonzkevmglobalexitrootmock

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

// PolygonzkevmglobalexitrootmockMetaData contains all meta data concerning the Polygonzkevmglobalexitrootmock contract.
var PolygonzkevmglobalexitrootmockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollupAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"}],\"name\":\"UpdateGlobalExitRoot\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastGlobalExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastMainnetExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"setGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"setLastGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561000f575f80fd5b5060405161047838038061047883398101604081905261002e91610060565b6001600160a01b0391821660a05216608052610091565b80516001600160a01b038116811461005b575f80fd5b919050565b5f8060408385031215610071575f80fd5b61007a83610045565b915061008860208401610045565b90509250929050565b60805160a0516103b86100c05f395f818161013a015261022c01525f818161018601526101e301526103b85ff3fe608060405234801561000f575f80fd5b506004361061009f575f3560e01c806333d6247d116100725780635bcef673116100585780635bcef673146101165780635ec6a8df14610135578063a3c573eb14610181575f80fd5b806333d6247d146100fb5780633ed691ef1461010e575f80fd5b806301fd9044146100a3578063051a9e28146100be578063257b3632146100d3578063319cf735146100f2575b5f80fd5b6100ab5f5481565b6040519081526020015b60405180910390f35b6100d16100cc36600461034b565b6101a8565b005b6100ab6100e136600461034b565b60026020525f908152604090205481565b6100ab60015481565b6100d161010936600461034b565b6101c7565b6100ab610311565b6100d1610124366004610362565b5f9182526002602052604090912055565b61015c7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100b5565b61015c7f000000000000000000000000000000000000000000000000000000000000000081565b8060025f6101b4610311565b815260208101919091526040015f205550565b5f5460015473ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001633036102155750600182905581610290565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016330361025e575f839055829150610290565b6040517fb49365dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080516020808201849052818301859052825180830384018152606090920190925280519101205f905f818152600260205260408120549192500361030b575f8181526002602052604080822042905551849184917f61014378f82a0d809aefaf87a8ac9505b89c321808287a6e7810f29304c1fce39190a35b50505050565b5f6103466001545f54604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b905090565b5f6020828403121561035b575f80fd5b5035919050565b5f8060408385031215610373575f80fd5b5050803592602090910135915056fea26469706673582212200799b5f60aa8eaf6d49205025cc41bfba78ce4f6ad614164e5da3c0bdc94810464736f6c63430008140033",
}

// PolygonzkevmglobalexitrootmockABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonzkevmglobalexitrootmockMetaData.ABI instead.
var PolygonzkevmglobalexitrootmockABI = PolygonzkevmglobalexitrootmockMetaData.ABI

// PolygonzkevmglobalexitrootmockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygonzkevmglobalexitrootmockMetaData.Bin instead.
var PolygonzkevmglobalexitrootmockBin = PolygonzkevmglobalexitrootmockMetaData.Bin

// DeployPolygonzkevmglobalexitrootmock deploys a new Ethereum contract, binding an instance of Polygonzkevmglobalexitrootmock to it.
func DeployPolygonzkevmglobalexitrootmock(auth *bind.TransactOpts, backend bind.ContractBackend, _rollupAddress common.Address, _bridgeAddress common.Address) (common.Address, *types.Transaction, *Polygonzkevmglobalexitrootmock, error) {
	parsed, err := PolygonzkevmglobalexitrootmockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygonzkevmglobalexitrootmockBin), backend, _rollupAddress, _bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonzkevmglobalexitrootmock{PolygonzkevmglobalexitrootmockCaller: PolygonzkevmglobalexitrootmockCaller{contract: contract}, PolygonzkevmglobalexitrootmockTransactor: PolygonzkevmglobalexitrootmockTransactor{contract: contract}, PolygonzkevmglobalexitrootmockFilterer: PolygonzkevmglobalexitrootmockFilterer{contract: contract}}, nil
}

// Polygonzkevmglobalexitrootmock is an auto generated Go binding around an Ethereum contract.
type Polygonzkevmglobalexitrootmock struct {
	PolygonzkevmglobalexitrootmockCaller     // Read-only binding to the contract
	PolygonzkevmglobalexitrootmockTransactor // Write-only binding to the contract
	PolygonzkevmglobalexitrootmockFilterer   // Log filterer for contract events
}

// PolygonzkevmglobalexitrootmockCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonzkevmglobalexitrootmockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmglobalexitrootmockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonzkevmglobalexitrootmockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmglobalexitrootmockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonzkevmglobalexitrootmockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmglobalexitrootmockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonzkevmglobalexitrootmockSession struct {
	Contract     *Polygonzkevmglobalexitrootmock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// PolygonzkevmglobalexitrootmockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonzkevmglobalexitrootmockCallerSession struct {
	Contract *PolygonzkevmglobalexitrootmockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// PolygonzkevmglobalexitrootmockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonzkevmglobalexitrootmockTransactorSession struct {
	Contract     *PolygonzkevmglobalexitrootmockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// PolygonzkevmglobalexitrootmockRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonzkevmglobalexitrootmockRaw struct {
	Contract *Polygonzkevmglobalexitrootmock // Generic contract binding to access the raw methods on
}

// PolygonzkevmglobalexitrootmockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonzkevmglobalexitrootmockCallerRaw struct {
	Contract *PolygonzkevmglobalexitrootmockCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonzkevmglobalexitrootmockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonzkevmglobalexitrootmockTransactorRaw struct {
	Contract *PolygonzkevmglobalexitrootmockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonzkevmglobalexitrootmock creates a new instance of Polygonzkevmglobalexitrootmock, bound to a specific deployed contract.
func NewPolygonzkevmglobalexitrootmock(address common.Address, backend bind.ContractBackend) (*Polygonzkevmglobalexitrootmock, error) {
	contract, err := bindPolygonzkevmglobalexitrootmock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootmock{PolygonzkevmglobalexitrootmockCaller: PolygonzkevmglobalexitrootmockCaller{contract: contract}, PolygonzkevmglobalexitrootmockTransactor: PolygonzkevmglobalexitrootmockTransactor{contract: contract}, PolygonzkevmglobalexitrootmockFilterer: PolygonzkevmglobalexitrootmockFilterer{contract: contract}}, nil
}

// NewPolygonzkevmglobalexitrootmockCaller creates a new read-only instance of Polygonzkevmglobalexitrootmock, bound to a specific deployed contract.
func NewPolygonzkevmglobalexitrootmockCaller(address common.Address, caller bind.ContractCaller) (*PolygonzkevmglobalexitrootmockCaller, error) {
	contract, err := bindPolygonzkevmglobalexitrootmock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmglobalexitrootmockCaller{contract: contract}, nil
}

// NewPolygonzkevmglobalexitrootmockTransactor creates a new write-only instance of Polygonzkevmglobalexitrootmock, bound to a specific deployed contract.
func NewPolygonzkevmglobalexitrootmockTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonzkevmglobalexitrootmockTransactor, error) {
	contract, err := bindPolygonzkevmglobalexitrootmock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmglobalexitrootmockTransactor{contract: contract}, nil
}

// NewPolygonzkevmglobalexitrootmockFilterer creates a new log filterer instance of Polygonzkevmglobalexitrootmock, bound to a specific deployed contract.
func NewPolygonzkevmglobalexitrootmockFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonzkevmglobalexitrootmockFilterer, error) {
	contract, err := bindPolygonzkevmglobalexitrootmock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmglobalexitrootmockFilterer{contract: contract}, nil
}

// bindPolygonzkevmglobalexitrootmock binds a generic wrapper to an already deployed contract.
func bindPolygonzkevmglobalexitrootmock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonzkevmglobalexitrootmockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmglobalexitrootmock.Contract.PolygonzkevmglobalexitrootmockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootmock.Contract.PolygonzkevmglobalexitrootmockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootmock.Contract.PolygonzkevmglobalexitrootmockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmglobalexitrootmock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootmock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootmock.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootmock.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockSession) BridgeAddress() (common.Address, error) {
	return _Polygonzkevmglobalexitrootmock.Contract.BridgeAddress(&_Polygonzkevmglobalexitrootmock.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockCallerSession) BridgeAddress() (common.Address, error) {
	return _Polygonzkevmglobalexitrootmock.Contract.BridgeAddress(&_Polygonzkevmglobalexitrootmock.CallOpts)
}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockCaller) GetLastGlobalExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootmock.contract.Call(opts, &out, "getLastGlobalExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockSession) GetLastGlobalExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootmock.Contract.GetLastGlobalExitRoot(&_Polygonzkevmglobalexitrootmock.CallOpts)
}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockCallerSession) GetLastGlobalExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootmock.Contract.GetLastGlobalExitRoot(&_Polygonzkevmglobalexitrootmock.CallOpts)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockCaller) GlobalExitRootMap(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootmock.contract.Call(opts, &out, "globalExitRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Polygonzkevmglobalexitrootmock.Contract.GlobalExitRootMap(&_Polygonzkevmglobalexitrootmock.CallOpts, arg0)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockCallerSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Polygonzkevmglobalexitrootmock.Contract.GlobalExitRootMap(&_Polygonzkevmglobalexitrootmock.CallOpts, arg0)
}

// LastMainnetExitRoot is a free data retrieval call binding the contract method 0x319cf735.
//
// Solidity: function lastMainnetExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockCaller) LastMainnetExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootmock.contract.Call(opts, &out, "lastMainnetExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastMainnetExitRoot is a free data retrieval call binding the contract method 0x319cf735.
//
// Solidity: function lastMainnetExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockSession) LastMainnetExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootmock.Contract.LastMainnetExitRoot(&_Polygonzkevmglobalexitrootmock.CallOpts)
}

// LastMainnetExitRoot is a free data retrieval call binding the contract method 0x319cf735.
//
// Solidity: function lastMainnetExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockCallerSession) LastMainnetExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootmock.Contract.LastMainnetExitRoot(&_Polygonzkevmglobalexitrootmock.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockCaller) LastRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootmock.contract.Call(opts, &out, "lastRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockSession) LastRollupExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootmock.Contract.LastRollupExitRoot(&_Polygonzkevmglobalexitrootmock.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockCallerSession) LastRollupExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootmock.Contract.LastRollupExitRoot(&_Polygonzkevmglobalexitrootmock.CallOpts)
}

// RollupAddress is a free data retrieval call binding the contract method 0x5ec6a8df.
//
// Solidity: function rollupAddress() view returns(address)
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockCaller) RollupAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootmock.contract.Call(opts, &out, "rollupAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupAddress is a free data retrieval call binding the contract method 0x5ec6a8df.
//
// Solidity: function rollupAddress() view returns(address)
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockSession) RollupAddress() (common.Address, error) {
	return _Polygonzkevmglobalexitrootmock.Contract.RollupAddress(&_Polygonzkevmglobalexitrootmock.CallOpts)
}

// RollupAddress is a free data retrieval call binding the contract method 0x5ec6a8df.
//
// Solidity: function rollupAddress() view returns(address)
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockCallerSession) RollupAddress() (common.Address, error) {
	return _Polygonzkevmglobalexitrootmock.Contract.RollupAddress(&_Polygonzkevmglobalexitrootmock.CallOpts)
}

// SetGlobalExitRoot is a paid mutator transaction binding the contract method 0x5bcef673.
//
// Solidity: function setGlobalExitRoot(bytes32 globalExitRoot, uint256 timestamp) returns()
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockTransactor) SetGlobalExitRoot(opts *bind.TransactOpts, globalExitRoot [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootmock.contract.Transact(opts, "setGlobalExitRoot", globalExitRoot, timestamp)
}

// SetGlobalExitRoot is a paid mutator transaction binding the contract method 0x5bcef673.
//
// Solidity: function setGlobalExitRoot(bytes32 globalExitRoot, uint256 timestamp) returns()
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockSession) SetGlobalExitRoot(globalExitRoot [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootmock.Contract.SetGlobalExitRoot(&_Polygonzkevmglobalexitrootmock.TransactOpts, globalExitRoot, timestamp)
}

// SetGlobalExitRoot is a paid mutator transaction binding the contract method 0x5bcef673.
//
// Solidity: function setGlobalExitRoot(bytes32 globalExitRoot, uint256 timestamp) returns()
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockTransactorSession) SetGlobalExitRoot(globalExitRoot [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootmock.Contract.SetGlobalExitRoot(&_Polygonzkevmglobalexitrootmock.TransactOpts, globalExitRoot, timestamp)
}

// SetLastGlobalExitRoot is a paid mutator transaction binding the contract method 0x051a9e28.
//
// Solidity: function setLastGlobalExitRoot(uint256 timestamp) returns()
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockTransactor) SetLastGlobalExitRoot(opts *bind.TransactOpts, timestamp *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootmock.contract.Transact(opts, "setLastGlobalExitRoot", timestamp)
}

// SetLastGlobalExitRoot is a paid mutator transaction binding the contract method 0x051a9e28.
//
// Solidity: function setLastGlobalExitRoot(uint256 timestamp) returns()
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockSession) SetLastGlobalExitRoot(timestamp *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootmock.Contract.SetLastGlobalExitRoot(&_Polygonzkevmglobalexitrootmock.TransactOpts, timestamp)
}

// SetLastGlobalExitRoot is a paid mutator transaction binding the contract method 0x051a9e28.
//
// Solidity: function setLastGlobalExitRoot(uint256 timestamp) returns()
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockTransactorSession) SetLastGlobalExitRoot(timestamp *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootmock.Contract.SetLastGlobalExitRoot(&_Polygonzkevmglobalexitrootmock.TransactOpts, timestamp)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockTransactor) UpdateExitRoot(opts *bind.TransactOpts, newRoot [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootmock.contract.Transact(opts, "updateExitRoot", newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockSession) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootmock.Contract.UpdateExitRoot(&_Polygonzkevmglobalexitrootmock.TransactOpts, newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockTransactorSession) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootmock.Contract.UpdateExitRoot(&_Polygonzkevmglobalexitrootmock.TransactOpts, newRoot)
}

// PolygonzkevmglobalexitrootmockUpdateGlobalExitRootIterator is returned from FilterUpdateGlobalExitRoot and is used to iterate over the raw logs and unpacked data for UpdateGlobalExitRoot events raised by the Polygonzkevmglobalexitrootmock contract.
type PolygonzkevmglobalexitrootmockUpdateGlobalExitRootIterator struct {
	Event *PolygonzkevmglobalexitrootmockUpdateGlobalExitRoot // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmglobalexitrootmockUpdateGlobalExitRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmglobalexitrootmockUpdateGlobalExitRoot)
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
		it.Event = new(PolygonzkevmglobalexitrootmockUpdateGlobalExitRoot)
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
func (it *PolygonzkevmglobalexitrootmockUpdateGlobalExitRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmglobalexitrootmockUpdateGlobalExitRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmglobalexitrootmockUpdateGlobalExitRoot represents a UpdateGlobalExitRoot event raised by the Polygonzkevmglobalexitrootmock contract.
type PolygonzkevmglobalexitrootmockUpdateGlobalExitRoot struct {
	MainnetExitRoot [32]byte
	RollupExitRoot  [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateGlobalExitRoot is a free log retrieval operation binding the contract event 0x61014378f82a0d809aefaf87a8ac9505b89c321808287a6e7810f29304c1fce3.
//
// Solidity: event UpdateGlobalExitRoot(bytes32 indexed mainnetExitRoot, bytes32 indexed rollupExitRoot)
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockFilterer) FilterUpdateGlobalExitRoot(opts *bind.FilterOpts, mainnetExitRoot [][32]byte, rollupExitRoot [][32]byte) (*PolygonzkevmglobalexitrootmockUpdateGlobalExitRootIterator, error) {

	var mainnetExitRootRule []interface{}
	for _, mainnetExitRootItem := range mainnetExitRoot {
		mainnetExitRootRule = append(mainnetExitRootRule, mainnetExitRootItem)
	}
	var rollupExitRootRule []interface{}
	for _, rollupExitRootItem := range rollupExitRoot {
		rollupExitRootRule = append(rollupExitRootRule, rollupExitRootItem)
	}

	logs, sub, err := _Polygonzkevmglobalexitrootmock.contract.FilterLogs(opts, "UpdateGlobalExitRoot", mainnetExitRootRule, rollupExitRootRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmglobalexitrootmockUpdateGlobalExitRootIterator{contract: _Polygonzkevmglobalexitrootmock.contract, event: "UpdateGlobalExitRoot", logs: logs, sub: sub}, nil
}

// WatchUpdateGlobalExitRoot is a free log subscription operation binding the contract event 0x61014378f82a0d809aefaf87a8ac9505b89c321808287a6e7810f29304c1fce3.
//
// Solidity: event UpdateGlobalExitRoot(bytes32 indexed mainnetExitRoot, bytes32 indexed rollupExitRoot)
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockFilterer) WatchUpdateGlobalExitRoot(opts *bind.WatchOpts, sink chan<- *PolygonzkevmglobalexitrootmockUpdateGlobalExitRoot, mainnetExitRoot [][32]byte, rollupExitRoot [][32]byte) (event.Subscription, error) {

	var mainnetExitRootRule []interface{}
	for _, mainnetExitRootItem := range mainnetExitRoot {
		mainnetExitRootRule = append(mainnetExitRootRule, mainnetExitRootItem)
	}
	var rollupExitRootRule []interface{}
	for _, rollupExitRootItem := range rollupExitRoot {
		rollupExitRootRule = append(rollupExitRootRule, rollupExitRootItem)
	}

	logs, sub, err := _Polygonzkevmglobalexitrootmock.contract.WatchLogs(opts, "UpdateGlobalExitRoot", mainnetExitRootRule, rollupExitRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmglobalexitrootmockUpdateGlobalExitRoot)
				if err := _Polygonzkevmglobalexitrootmock.contract.UnpackLog(event, "UpdateGlobalExitRoot", log); err != nil {
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

// ParseUpdateGlobalExitRoot is a log parse operation binding the contract event 0x61014378f82a0d809aefaf87a8ac9505b89c321808287a6e7810f29304c1fce3.
//
// Solidity: event UpdateGlobalExitRoot(bytes32 indexed mainnetExitRoot, bytes32 indexed rollupExitRoot)
func (_Polygonzkevmglobalexitrootmock *PolygonzkevmglobalexitrootmockFilterer) ParseUpdateGlobalExitRoot(log types.Log) (*PolygonzkevmglobalexitrootmockUpdateGlobalExitRoot, error) {
	event := new(PolygonzkevmglobalexitrootmockUpdateGlobalExitRoot)
	if err := _Polygonzkevmglobalexitrootmock.contract.UnpackLog(event, "UpdateGlobalExitRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

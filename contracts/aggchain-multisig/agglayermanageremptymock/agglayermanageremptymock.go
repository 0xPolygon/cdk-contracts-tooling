// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package agglayermanageremptymock

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

// AgglayermanageremptymockMetaData contains all meta data concerning the Agglayermanageremptymock contract.
var AgglayermanageremptymockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getForcedBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastDeactivatedEmergencyStateTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newSequencedBatches\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newAccInputHash\",\"type\":\"bytes32\"}],\"name\":\"onSequenceBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"contractIPolygonRollupBase\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"onVerifyBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"newAcceptSequenceBatches\",\"type\":\"bool\"}],\"name\":\"setAcceptSequenceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052600c805460ff19166001179055348015601b575f5ffd5b506104a2806100295f395ff3fe608060405234801561000f575f5ffd5b506004361061009f575f3560e01c8063477fa270116100725780637e4da1bc116100585780637e4da1bc146101195780639a908e7314610158578063dbc169761461016b575f5ffd5b8063477fa2701461010b5780636046916914610112575f5ffd5b806315064c96146100a35780632072f6c5146100c557806330c27dde146100cf57806332c2d153146100df575b5f5ffd5b600a546100b09060ff1681565b60405190151581526020015b60405180910390f35b6100cd610173565b005b5f5b6040519081526020016100bc565b6100f26100ed366004610391565b61017d565b60405167ffffffffffffffff90911681526020016100bc565b60016100d1565b600a6100d1565b6100cd6101273660046103e6565b600c80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055565b6100f261016636600461040c565b610217565b6100cd61024d565b61017b610255565b565b6040517f32c2d15300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84166004820152602481018390523360448201525f9073ffffffffffffffffffffffffffffffffffffffff8316906332c2d153906064015f604051808303815f87803b1580156101fa575f5ffd5b505af115801561020c573d5f5f3e3d5ffd5b505050509392505050565b600c545f9060ff16610227575f5ffd5b8267ffffffffffffffff16600b5461023f9190610434565b600b81905590505b92915050565b61017b6102e7565b600a5460ff1615610292576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600a80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556040517f2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497905f90a1565b600a5460ff16610323576040517f5386698100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600a80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556040517f1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3905f90a1565b803567ffffffffffffffff8116811461038c575f5ffd5b919050565b5f5f5f606084860312156103a3575f5ffd5b6103ac84610375565b925060208401359150604084013573ffffffffffffffffffffffffffffffffffffffff811681146103db575f5ffd5b809150509250925092565b5f602082840312156103f6575f5ffd5b81358015158114610405575f5ffd5b9392505050565b5f5f6040838503121561041d575f5ffd5b61042683610375565b946020939093013593505050565b80820180821115610247577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffdfea264697066735822122070b936a4313d37691e745a385a7c19f6dec5dd65d640d0771282b5007669d00164736f6c634300081c0033",
}

// AgglayermanageremptymockABI is the input ABI used to generate the binding from.
// Deprecated: Use AgglayermanageremptymockMetaData.ABI instead.
var AgglayermanageremptymockABI = AgglayermanageremptymockMetaData.ABI

// AgglayermanageremptymockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AgglayermanageremptymockMetaData.Bin instead.
var AgglayermanageremptymockBin = AgglayermanageremptymockMetaData.Bin

// DeployAgglayermanageremptymock deploys a new Ethereum contract, binding an instance of Agglayermanageremptymock to it.
func DeployAgglayermanageremptymock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Agglayermanageremptymock, error) {
	parsed, err := AgglayermanageremptymockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AgglayermanageremptymockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Agglayermanageremptymock{AgglayermanageremptymockCaller: AgglayermanageremptymockCaller{contract: contract}, AgglayermanageremptymockTransactor: AgglayermanageremptymockTransactor{contract: contract}, AgglayermanageremptymockFilterer: AgglayermanageremptymockFilterer{contract: contract}}, nil
}

// Agglayermanageremptymock is an auto generated Go binding around an Ethereum contract.
type Agglayermanageremptymock struct {
	AgglayermanageremptymockCaller     // Read-only binding to the contract
	AgglayermanageremptymockTransactor // Write-only binding to the contract
	AgglayermanageremptymockFilterer   // Log filterer for contract events
}

// AgglayermanageremptymockCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgglayermanageremptymockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgglayermanageremptymockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgglayermanageremptymockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgglayermanageremptymockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgglayermanageremptymockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgglayermanageremptymockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgglayermanageremptymockSession struct {
	Contract     *Agglayermanageremptymock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AgglayermanageremptymockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgglayermanageremptymockCallerSession struct {
	Contract *AgglayermanageremptymockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// AgglayermanageremptymockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgglayermanageremptymockTransactorSession struct {
	Contract     *AgglayermanageremptymockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// AgglayermanageremptymockRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgglayermanageremptymockRaw struct {
	Contract *Agglayermanageremptymock // Generic contract binding to access the raw methods on
}

// AgglayermanageremptymockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgglayermanageremptymockCallerRaw struct {
	Contract *AgglayermanageremptymockCaller // Generic read-only contract binding to access the raw methods on
}

// AgglayermanageremptymockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgglayermanageremptymockTransactorRaw struct {
	Contract *AgglayermanageremptymockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgglayermanageremptymock creates a new instance of Agglayermanageremptymock, bound to a specific deployed contract.
func NewAgglayermanageremptymock(address common.Address, backend bind.ContractBackend) (*Agglayermanageremptymock, error) {
	contract, err := bindAgglayermanageremptymock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Agglayermanageremptymock{AgglayermanageremptymockCaller: AgglayermanageremptymockCaller{contract: contract}, AgglayermanageremptymockTransactor: AgglayermanageremptymockTransactor{contract: contract}, AgglayermanageremptymockFilterer: AgglayermanageremptymockFilterer{contract: contract}}, nil
}

// NewAgglayermanageremptymockCaller creates a new read-only instance of Agglayermanageremptymock, bound to a specific deployed contract.
func NewAgglayermanageremptymockCaller(address common.Address, caller bind.ContractCaller) (*AgglayermanageremptymockCaller, error) {
	contract, err := bindAgglayermanageremptymock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgglayermanageremptymockCaller{contract: contract}, nil
}

// NewAgglayermanageremptymockTransactor creates a new write-only instance of Agglayermanageremptymock, bound to a specific deployed contract.
func NewAgglayermanageremptymockTransactor(address common.Address, transactor bind.ContractTransactor) (*AgglayermanageremptymockTransactor, error) {
	contract, err := bindAgglayermanageremptymock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgglayermanageremptymockTransactor{contract: contract}, nil
}

// NewAgglayermanageremptymockFilterer creates a new log filterer instance of Agglayermanageremptymock, bound to a specific deployed contract.
func NewAgglayermanageremptymockFilterer(address common.Address, filterer bind.ContractFilterer) (*AgglayermanageremptymockFilterer, error) {
	contract, err := bindAgglayermanageremptymock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgglayermanageremptymockFilterer{contract: contract}, nil
}

// bindAgglayermanageremptymock binds a generic wrapper to an already deployed contract.
func bindAgglayermanageremptymock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AgglayermanageremptymockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Agglayermanageremptymock *AgglayermanageremptymockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Agglayermanageremptymock.Contract.AgglayermanageremptymockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Agglayermanageremptymock *AgglayermanageremptymockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayermanageremptymock.Contract.AgglayermanageremptymockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Agglayermanageremptymock *AgglayermanageremptymockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Agglayermanageremptymock.Contract.AgglayermanageremptymockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Agglayermanageremptymock *AgglayermanageremptymockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Agglayermanageremptymock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Agglayermanageremptymock *AgglayermanageremptymockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayermanageremptymock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Agglayermanageremptymock *AgglayermanageremptymockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Agglayermanageremptymock.Contract.contract.Transact(opts, method, params...)
}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Agglayermanageremptymock *AgglayermanageremptymockCaller) GetBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Agglayermanageremptymock.contract.Call(opts, &out, "getBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Agglayermanageremptymock *AgglayermanageremptymockSession) GetBatchFee() (*big.Int, error) {
	return _Agglayermanageremptymock.Contract.GetBatchFee(&_Agglayermanageremptymock.CallOpts)
}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Agglayermanageremptymock *AgglayermanageremptymockCallerSession) GetBatchFee() (*big.Int, error) {
	return _Agglayermanageremptymock.Contract.GetBatchFee(&_Agglayermanageremptymock.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Agglayermanageremptymock *AgglayermanageremptymockCaller) GetForcedBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Agglayermanageremptymock.contract.Call(opts, &out, "getForcedBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Agglayermanageremptymock *AgglayermanageremptymockSession) GetForcedBatchFee() (*big.Int, error) {
	return _Agglayermanageremptymock.Contract.GetForcedBatchFee(&_Agglayermanageremptymock.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Agglayermanageremptymock *AgglayermanageremptymockCallerSession) GetForcedBatchFee() (*big.Int, error) {
	return _Agglayermanageremptymock.Contract.GetForcedBatchFee(&_Agglayermanageremptymock.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Agglayermanageremptymock *AgglayermanageremptymockCaller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Agglayermanageremptymock.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Agglayermanageremptymock *AgglayermanageremptymockSession) IsEmergencyState() (bool, error) {
	return _Agglayermanageremptymock.Contract.IsEmergencyState(&_Agglayermanageremptymock.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Agglayermanageremptymock *AgglayermanageremptymockCallerSession) IsEmergencyState() (bool, error) {
	return _Agglayermanageremptymock.Contract.IsEmergencyState(&_Agglayermanageremptymock.CallOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Agglayermanageremptymock *AgglayermanageremptymockTransactor) ActivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayermanageremptymock.contract.Transact(opts, "activateEmergencyState")
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Agglayermanageremptymock *AgglayermanageremptymockSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Agglayermanageremptymock.Contract.ActivateEmergencyState(&_Agglayermanageremptymock.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Agglayermanageremptymock *AgglayermanageremptymockTransactorSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Agglayermanageremptymock.Contract.ActivateEmergencyState(&_Agglayermanageremptymock.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Agglayermanageremptymock *AgglayermanageremptymockTransactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayermanageremptymock.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Agglayermanageremptymock *AgglayermanageremptymockSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Agglayermanageremptymock.Contract.DeactivateEmergencyState(&_Agglayermanageremptymock.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Agglayermanageremptymock *AgglayermanageremptymockTransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Agglayermanageremptymock.Contract.DeactivateEmergencyState(&_Agglayermanageremptymock.TransactOpts)
}

// LastDeactivatedEmergencyStateTimestamp is a paid mutator transaction binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() returns(uint256)
func (_Agglayermanageremptymock *AgglayermanageremptymockTransactor) LastDeactivatedEmergencyStateTimestamp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayermanageremptymock.contract.Transact(opts, "lastDeactivatedEmergencyStateTimestamp")
}

// LastDeactivatedEmergencyStateTimestamp is a paid mutator transaction binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() returns(uint256)
func (_Agglayermanageremptymock *AgglayermanageremptymockSession) LastDeactivatedEmergencyStateTimestamp() (*types.Transaction, error) {
	return _Agglayermanageremptymock.Contract.LastDeactivatedEmergencyStateTimestamp(&_Agglayermanageremptymock.TransactOpts)
}

// LastDeactivatedEmergencyStateTimestamp is a paid mutator transaction binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() returns(uint256)
func (_Agglayermanageremptymock *AgglayermanageremptymockTransactorSession) LastDeactivatedEmergencyStateTimestamp() (*types.Transaction, error) {
	return _Agglayermanageremptymock.Contract.LastDeactivatedEmergencyStateTimestamp(&_Agglayermanageremptymock.TransactOpts)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Agglayermanageremptymock *AgglayermanageremptymockTransactor) OnSequenceBatches(opts *bind.TransactOpts, newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Agglayermanageremptymock.contract.Transact(opts, "onSequenceBatches", newSequencedBatches, newAccInputHash)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Agglayermanageremptymock *AgglayermanageremptymockSession) OnSequenceBatches(newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Agglayermanageremptymock.Contract.OnSequenceBatches(&_Agglayermanageremptymock.TransactOpts, newSequencedBatches, newAccInputHash)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Agglayermanageremptymock *AgglayermanageremptymockTransactorSession) OnSequenceBatches(newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Agglayermanageremptymock.Contract.OnSequenceBatches(&_Agglayermanageremptymock.TransactOpts, newSequencedBatches, newAccInputHash)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 finalNewBatch, bytes32 newStateRoot, address rollup) returns(uint64)
func (_Agglayermanageremptymock *AgglayermanageremptymockTransactor) OnVerifyBatches(opts *bind.TransactOpts, finalNewBatch uint64, newStateRoot [32]byte, rollup common.Address) (*types.Transaction, error) {
	return _Agglayermanageremptymock.contract.Transact(opts, "onVerifyBatches", finalNewBatch, newStateRoot, rollup)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 finalNewBatch, bytes32 newStateRoot, address rollup) returns(uint64)
func (_Agglayermanageremptymock *AgglayermanageremptymockSession) OnVerifyBatches(finalNewBatch uint64, newStateRoot [32]byte, rollup common.Address) (*types.Transaction, error) {
	return _Agglayermanageremptymock.Contract.OnVerifyBatches(&_Agglayermanageremptymock.TransactOpts, finalNewBatch, newStateRoot, rollup)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 finalNewBatch, bytes32 newStateRoot, address rollup) returns(uint64)
func (_Agglayermanageremptymock *AgglayermanageremptymockTransactorSession) OnVerifyBatches(finalNewBatch uint64, newStateRoot [32]byte, rollup common.Address) (*types.Transaction, error) {
	return _Agglayermanageremptymock.Contract.OnVerifyBatches(&_Agglayermanageremptymock.TransactOpts, finalNewBatch, newStateRoot, rollup)
}

// SetAcceptSequenceBatches is a paid mutator transaction binding the contract method 0x7e4da1bc.
//
// Solidity: function setAcceptSequenceBatches(bool newAcceptSequenceBatches) returns()
func (_Agglayermanageremptymock *AgglayermanageremptymockTransactor) SetAcceptSequenceBatches(opts *bind.TransactOpts, newAcceptSequenceBatches bool) (*types.Transaction, error) {
	return _Agglayermanageremptymock.contract.Transact(opts, "setAcceptSequenceBatches", newAcceptSequenceBatches)
}

// SetAcceptSequenceBatches is a paid mutator transaction binding the contract method 0x7e4da1bc.
//
// Solidity: function setAcceptSequenceBatches(bool newAcceptSequenceBatches) returns()
func (_Agglayermanageremptymock *AgglayermanageremptymockSession) SetAcceptSequenceBatches(newAcceptSequenceBatches bool) (*types.Transaction, error) {
	return _Agglayermanageremptymock.Contract.SetAcceptSequenceBatches(&_Agglayermanageremptymock.TransactOpts, newAcceptSequenceBatches)
}

// SetAcceptSequenceBatches is a paid mutator transaction binding the contract method 0x7e4da1bc.
//
// Solidity: function setAcceptSequenceBatches(bool newAcceptSequenceBatches) returns()
func (_Agglayermanageremptymock *AgglayermanageremptymockTransactorSession) SetAcceptSequenceBatches(newAcceptSequenceBatches bool) (*types.Transaction, error) {
	return _Agglayermanageremptymock.Contract.SetAcceptSequenceBatches(&_Agglayermanageremptymock.TransactOpts, newAcceptSequenceBatches)
}

// AgglayermanageremptymockEmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the Agglayermanageremptymock contract.
type AgglayermanageremptymockEmergencyStateActivatedIterator struct {
	Event *AgglayermanageremptymockEmergencyStateActivated // Event containing the contract specifics and raw log

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
func (it *AgglayermanageremptymockEmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanageremptymockEmergencyStateActivated)
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
		it.Event = new(AgglayermanageremptymockEmergencyStateActivated)
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
func (it *AgglayermanageremptymockEmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanageremptymockEmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanageremptymockEmergencyStateActivated represents a EmergencyStateActivated event raised by the Agglayermanageremptymock contract.
type AgglayermanageremptymockEmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Agglayermanageremptymock *AgglayermanageremptymockFilterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*AgglayermanageremptymockEmergencyStateActivatedIterator, error) {

	logs, sub, err := _Agglayermanageremptymock.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &AgglayermanageremptymockEmergencyStateActivatedIterator{contract: _Agglayermanageremptymock.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Agglayermanageremptymock *AgglayermanageremptymockFilterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *AgglayermanageremptymockEmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _Agglayermanageremptymock.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanageremptymockEmergencyStateActivated)
				if err := _Agglayermanageremptymock.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
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

// ParseEmergencyStateActivated is a log parse operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Agglayermanageremptymock *AgglayermanageremptymockFilterer) ParseEmergencyStateActivated(log types.Log) (*AgglayermanageremptymockEmergencyStateActivated, error) {
	event := new(AgglayermanageremptymockEmergencyStateActivated)
	if err := _Agglayermanageremptymock.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanageremptymockEmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the Agglayermanageremptymock contract.
type AgglayermanageremptymockEmergencyStateDeactivatedIterator struct {
	Event *AgglayermanageremptymockEmergencyStateDeactivated // Event containing the contract specifics and raw log

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
func (it *AgglayermanageremptymockEmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanageremptymockEmergencyStateDeactivated)
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
		it.Event = new(AgglayermanageremptymockEmergencyStateDeactivated)
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
func (it *AgglayermanageremptymockEmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanageremptymockEmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanageremptymockEmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the Agglayermanageremptymock contract.
type AgglayermanageremptymockEmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Agglayermanageremptymock *AgglayermanageremptymockFilterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*AgglayermanageremptymockEmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _Agglayermanageremptymock.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &AgglayermanageremptymockEmergencyStateDeactivatedIterator{contract: _Agglayermanageremptymock.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Agglayermanageremptymock *AgglayermanageremptymockFilterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *AgglayermanageremptymockEmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _Agglayermanageremptymock.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanageremptymockEmergencyStateDeactivated)
				if err := _Agglayermanageremptymock.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
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

// ParseEmergencyStateDeactivated is a log parse operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Agglayermanageremptymock *AgglayermanageremptymockFilterer) ParseEmergencyStateDeactivated(log types.Log) (*AgglayermanageremptymockEmergencyStateDeactivated, error) {
	event := new(AgglayermanageremptymockEmergencyStateDeactivated)
	if err := _Agglayermanageremptymock.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

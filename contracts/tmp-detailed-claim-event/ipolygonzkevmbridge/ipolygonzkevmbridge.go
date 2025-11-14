// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipolygonzkevmbridge

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

// IpolygonzkevmbridgeMetaData contains all meta data concerning the Ipolygonzkevmbridge contract.
var IpolygonzkevmbridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountDoesNotMatchMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DestinationNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EtherTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSmtProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MsgValueNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPolygonZkEVM\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"permitData\",\"type\":\"bytes\"}],\"name\":\"bridgeAsset\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IpolygonzkevmbridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use IpolygonzkevmbridgeMetaData.ABI instead.
var IpolygonzkevmbridgeABI = IpolygonzkevmbridgeMetaData.ABI

// Ipolygonzkevmbridge is an auto generated Go binding around an Ethereum contract.
type Ipolygonzkevmbridge struct {
	IpolygonzkevmbridgeCaller     // Read-only binding to the contract
	IpolygonzkevmbridgeTransactor // Write-only binding to the contract
	IpolygonzkevmbridgeFilterer   // Log filterer for contract events
}

// IpolygonzkevmbridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IpolygonzkevmbridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonzkevmbridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IpolygonzkevmbridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonzkevmbridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IpolygonzkevmbridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonzkevmbridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IpolygonzkevmbridgeSession struct {
	Contract     *Ipolygonzkevmbridge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IpolygonzkevmbridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IpolygonzkevmbridgeCallerSession struct {
	Contract *IpolygonzkevmbridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IpolygonzkevmbridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IpolygonzkevmbridgeTransactorSession struct {
	Contract     *IpolygonzkevmbridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IpolygonzkevmbridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IpolygonzkevmbridgeRaw struct {
	Contract *Ipolygonzkevmbridge // Generic contract binding to access the raw methods on
}

// IpolygonzkevmbridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IpolygonzkevmbridgeCallerRaw struct {
	Contract *IpolygonzkevmbridgeCaller // Generic read-only contract binding to access the raw methods on
}

// IpolygonzkevmbridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IpolygonzkevmbridgeTransactorRaw struct {
	Contract *IpolygonzkevmbridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIpolygonzkevmbridge creates a new instance of Ipolygonzkevmbridge, bound to a specific deployed contract.
func NewIpolygonzkevmbridge(address common.Address, backend bind.ContractBackend) (*Ipolygonzkevmbridge, error) {
	contract, err := bindIpolygonzkevmbridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ipolygonzkevmbridge{IpolygonzkevmbridgeCaller: IpolygonzkevmbridgeCaller{contract: contract}, IpolygonzkevmbridgeTransactor: IpolygonzkevmbridgeTransactor{contract: contract}, IpolygonzkevmbridgeFilterer: IpolygonzkevmbridgeFilterer{contract: contract}}, nil
}

// NewIpolygonzkevmbridgeCaller creates a new read-only instance of Ipolygonzkevmbridge, bound to a specific deployed contract.
func NewIpolygonzkevmbridgeCaller(address common.Address, caller bind.ContractCaller) (*IpolygonzkevmbridgeCaller, error) {
	contract, err := bindIpolygonzkevmbridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygonzkevmbridgeCaller{contract: contract}, nil
}

// NewIpolygonzkevmbridgeTransactor creates a new write-only instance of Ipolygonzkevmbridge, bound to a specific deployed contract.
func NewIpolygonzkevmbridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*IpolygonzkevmbridgeTransactor, error) {
	contract, err := bindIpolygonzkevmbridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygonzkevmbridgeTransactor{contract: contract}, nil
}

// NewIpolygonzkevmbridgeFilterer creates a new log filterer instance of Ipolygonzkevmbridge, bound to a specific deployed contract.
func NewIpolygonzkevmbridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*IpolygonzkevmbridgeFilterer, error) {
	contract, err := bindIpolygonzkevmbridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IpolygonzkevmbridgeFilterer{contract: contract}, nil
}

// bindIpolygonzkevmbridge binds a generic wrapper to an already deployed contract.
func bindIpolygonzkevmbridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IpolygonzkevmbridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonzkevmbridge.Contract.IpolygonzkevmbridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonzkevmbridge.Contract.IpolygonzkevmbridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonzkevmbridge.Contract.IpolygonzkevmbridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonzkevmbridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonzkevmbridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonzkevmbridge.Contract.contract.Transact(opts, method, params...)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeTransactor) ActivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonzkevmbridge.contract.Transact(opts, "activateEmergencyState")
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Ipolygonzkevmbridge.Contract.ActivateEmergencyState(&_Ipolygonzkevmbridge.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeTransactorSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Ipolygonzkevmbridge.Contract.ActivateEmergencyState(&_Ipolygonzkevmbridge.TransactOpts)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeTransactor) BridgeAsset(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridge.contract.Transact(opts, "bridgeAsset", destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeSession) BridgeAsset(destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridge.Contract.BridgeAsset(&_Ipolygonzkevmbridge.TransactOpts, destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeTransactorSession) BridgeAsset(destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridge.Contract.BridgeAsset(&_Ipolygonzkevmbridge.TransactOpts, destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeTransactor) BridgeMessage(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridge.contract.Transact(opts, "bridgeMessage", destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeSession) BridgeMessage(destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridge.Contract.BridgeMessage(&_Ipolygonzkevmbridge.TransactOpts, destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeTransactorSession) BridgeMessage(destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridge.Contract.BridgeMessage(&_Ipolygonzkevmbridge.TransactOpts, destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0x2cffd02e.
//
// Solidity: function claimAsset(bytes32[32] smtProof, uint32 index, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeTransactor) ClaimAsset(opts *bind.TransactOpts, smtProof [32][32]byte, index uint32, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridge.contract.Transact(opts, "claimAsset", smtProof, index, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0x2cffd02e.
//
// Solidity: function claimAsset(bytes32[32] smtProof, uint32 index, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeSession) ClaimAsset(smtProof [32][32]byte, index uint32, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridge.Contract.ClaimAsset(&_Ipolygonzkevmbridge.TransactOpts, smtProof, index, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0x2cffd02e.
//
// Solidity: function claimAsset(bytes32[32] smtProof, uint32 index, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeTransactorSession) ClaimAsset(smtProof [32][32]byte, index uint32, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridge.Contract.ClaimAsset(&_Ipolygonzkevmbridge.TransactOpts, smtProof, index, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x2d2c9d94.
//
// Solidity: function claimMessage(bytes32[32] smtProof, uint32 index, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeTransactor) ClaimMessage(opts *bind.TransactOpts, smtProof [32][32]byte, index uint32, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridge.contract.Transact(opts, "claimMessage", smtProof, index, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x2d2c9d94.
//
// Solidity: function claimMessage(bytes32[32] smtProof, uint32 index, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeSession) ClaimMessage(smtProof [32][32]byte, index uint32, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridge.Contract.ClaimMessage(&_Ipolygonzkevmbridge.TransactOpts, smtProof, index, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x2d2c9d94.
//
// Solidity: function claimMessage(bytes32[32] smtProof, uint32 index, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeTransactorSession) ClaimMessage(smtProof [32][32]byte, index uint32, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridge.Contract.ClaimMessage(&_Ipolygonzkevmbridge.TransactOpts, smtProof, index, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeTransactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonzkevmbridge.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Ipolygonzkevmbridge.Contract.DeactivateEmergencyState(&_Ipolygonzkevmbridge.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeTransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Ipolygonzkevmbridge.Contract.DeactivateEmergencyState(&_Ipolygonzkevmbridge.TransactOpts)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeTransactor) UpdateGlobalExitRoot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonzkevmbridge.contract.Transact(opts, "updateGlobalExitRoot")
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeSession) UpdateGlobalExitRoot() (*types.Transaction, error) {
	return _Ipolygonzkevmbridge.Contract.UpdateGlobalExitRoot(&_Ipolygonzkevmbridge.TransactOpts)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Ipolygonzkevmbridge *IpolygonzkevmbridgeTransactorSession) UpdateGlobalExitRoot() (*types.Transaction, error) {
	return _Ipolygonzkevmbridge.Contract.UpdateGlobalExitRoot(&_Ipolygonzkevmbridge.TransactOpts)
}

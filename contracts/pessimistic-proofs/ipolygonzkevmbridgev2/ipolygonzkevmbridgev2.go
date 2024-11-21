// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipolygonzkevmbridgev2

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

// Ipolygonzkevmbridgev2MetaData contains all meta data concerning the Ipolygonzkevmbridgev2 contract.
var Ipolygonzkevmbridgev2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountDoesNotMatchMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DestinationNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EtherTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTokenWrappedDeployment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSmtProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MsgValueNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeTokenIsEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoValueInMessagesOnGasTokenNetworks\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"permitData\",\"type\":\"bytes\"}],\"name\":\"bridgeAsset\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountWETH\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessageWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_polygonRollupManager\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"}],\"name\":\"wrappedTokenToTokenInfo\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Ipolygonzkevmbridgev2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Ipolygonzkevmbridgev2MetaData.ABI instead.
var Ipolygonzkevmbridgev2ABI = Ipolygonzkevmbridgev2MetaData.ABI

// Ipolygonzkevmbridgev2 is an auto generated Go binding around an Ethereum contract.
type Ipolygonzkevmbridgev2 struct {
	Ipolygonzkevmbridgev2Caller     // Read-only binding to the contract
	Ipolygonzkevmbridgev2Transactor // Write-only binding to the contract
	Ipolygonzkevmbridgev2Filterer   // Log filterer for contract events
}

// Ipolygonzkevmbridgev2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Ipolygonzkevmbridgev2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ipolygonzkevmbridgev2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Ipolygonzkevmbridgev2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ipolygonzkevmbridgev2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ipolygonzkevmbridgev2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ipolygonzkevmbridgev2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ipolygonzkevmbridgev2Session struct {
	Contract     *Ipolygonzkevmbridgev2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// Ipolygonzkevmbridgev2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ipolygonzkevmbridgev2CallerSession struct {
	Contract *Ipolygonzkevmbridgev2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// Ipolygonzkevmbridgev2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ipolygonzkevmbridgev2TransactorSession struct {
	Contract     *Ipolygonzkevmbridgev2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// Ipolygonzkevmbridgev2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Ipolygonzkevmbridgev2Raw struct {
	Contract *Ipolygonzkevmbridgev2 // Generic contract binding to access the raw methods on
}

// Ipolygonzkevmbridgev2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ipolygonzkevmbridgev2CallerRaw struct {
	Contract *Ipolygonzkevmbridgev2Caller // Generic read-only contract binding to access the raw methods on
}

// Ipolygonzkevmbridgev2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ipolygonzkevmbridgev2TransactorRaw struct {
	Contract *Ipolygonzkevmbridgev2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIpolygonzkevmbridgev2 creates a new instance of Ipolygonzkevmbridgev2, bound to a specific deployed contract.
func NewIpolygonzkevmbridgev2(address common.Address, backend bind.ContractBackend) (*Ipolygonzkevmbridgev2, error) {
	contract, err := bindIpolygonzkevmbridgev2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ipolygonzkevmbridgev2{Ipolygonzkevmbridgev2Caller: Ipolygonzkevmbridgev2Caller{contract: contract}, Ipolygonzkevmbridgev2Transactor: Ipolygonzkevmbridgev2Transactor{contract: contract}, Ipolygonzkevmbridgev2Filterer: Ipolygonzkevmbridgev2Filterer{contract: contract}}, nil
}

// NewIpolygonzkevmbridgev2Caller creates a new read-only instance of Ipolygonzkevmbridgev2, bound to a specific deployed contract.
func NewIpolygonzkevmbridgev2Caller(address common.Address, caller bind.ContractCaller) (*Ipolygonzkevmbridgev2Caller, error) {
	contract, err := bindIpolygonzkevmbridgev2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ipolygonzkevmbridgev2Caller{contract: contract}, nil
}

// NewIpolygonzkevmbridgev2Transactor creates a new write-only instance of Ipolygonzkevmbridgev2, bound to a specific deployed contract.
func NewIpolygonzkevmbridgev2Transactor(address common.Address, transactor bind.ContractTransactor) (*Ipolygonzkevmbridgev2Transactor, error) {
	contract, err := bindIpolygonzkevmbridgev2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ipolygonzkevmbridgev2Transactor{contract: contract}, nil
}

// NewIpolygonzkevmbridgev2Filterer creates a new log filterer instance of Ipolygonzkevmbridgev2, bound to a specific deployed contract.
func NewIpolygonzkevmbridgev2Filterer(address common.Address, filterer bind.ContractFilterer) (*Ipolygonzkevmbridgev2Filterer, error) {
	contract, err := bindIpolygonzkevmbridgev2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ipolygonzkevmbridgev2Filterer{contract: contract}, nil
}

// bindIpolygonzkevmbridgev2 binds a generic wrapper to an already deployed contract.
func bindIpolygonzkevmbridgev2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Ipolygonzkevmbridgev2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonzkevmbridgev2.Contract.Ipolygonzkevmbridgev2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.Contract.Ipolygonzkevmbridgev2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.Contract.Ipolygonzkevmbridgev2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonzkevmbridgev2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.Contract.contract.Transact(opts, method, params...)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2Caller) GetTokenMetadata(opts *bind.CallOpts, token common.Address) ([]byte, error) {
	var out []interface{}
	err := _Ipolygonzkevmbridgev2.contract.Call(opts, &out, "getTokenMetadata", token)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2Session) GetTokenMetadata(token common.Address) ([]byte, error) {
	return _Ipolygonzkevmbridgev2.Contract.GetTokenMetadata(&_Ipolygonzkevmbridgev2.CallOpts, token)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2CallerSession) GetTokenMetadata(token common.Address) ([]byte, error) {
	return _Ipolygonzkevmbridgev2.Contract.GetTokenMetadata(&_Ipolygonzkevmbridgev2.CallOpts, token)
}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address destinationAddress) view returns(uint32, address)
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2Caller) WrappedTokenToTokenInfo(opts *bind.CallOpts, destinationAddress common.Address) (uint32, common.Address, error) {
	var out []interface{}
	err := _Ipolygonzkevmbridgev2.contract.Call(opts, &out, "wrappedTokenToTokenInfo", destinationAddress)

	if err != nil {
		return *new(uint32), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address destinationAddress) view returns(uint32, address)
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2Session) WrappedTokenToTokenInfo(destinationAddress common.Address) (uint32, common.Address, error) {
	return _Ipolygonzkevmbridgev2.Contract.WrappedTokenToTokenInfo(&_Ipolygonzkevmbridgev2.CallOpts, destinationAddress)
}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address destinationAddress) view returns(uint32, address)
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2CallerSession) WrappedTokenToTokenInfo(destinationAddress common.Address) (uint32, common.Address, error) {
	return _Ipolygonzkevmbridgev2.Contract.WrappedTokenToTokenInfo(&_Ipolygonzkevmbridgev2.CallOpts, destinationAddress)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2Transactor) ActivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.contract.Transact(opts, "activateEmergencyState")
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2Session) ActivateEmergencyState() (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.Contract.ActivateEmergencyState(&_Ipolygonzkevmbridgev2.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2TransactorSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.Contract.ActivateEmergencyState(&_Ipolygonzkevmbridgev2.TransactOpts)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2Transactor) BridgeAsset(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.contract.Transact(opts, "bridgeAsset", destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2Session) BridgeAsset(destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.Contract.BridgeAsset(&_Ipolygonzkevmbridgev2.TransactOpts, destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2TransactorSession) BridgeAsset(destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.Contract.BridgeAsset(&_Ipolygonzkevmbridgev2.TransactOpts, destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2Transactor) BridgeMessage(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.contract.Transact(opts, "bridgeMessage", destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2Session) BridgeMessage(destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.Contract.BridgeMessage(&_Ipolygonzkevmbridgev2.TransactOpts, destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2TransactorSession) BridgeMessage(destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.Contract.BridgeMessage(&_Ipolygonzkevmbridgev2.TransactOpts, destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2Transactor) BridgeMessageWETH(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.contract.Transact(opts, "bridgeMessageWETH", destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2Session) BridgeMessageWETH(destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.Contract.BridgeMessageWETH(&_Ipolygonzkevmbridgev2.TransactOpts, destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2TransactorSession) BridgeMessageWETH(destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.Contract.BridgeMessageWETH(&_Ipolygonzkevmbridgev2.TransactOpts, destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2Transactor) ClaimAsset(opts *bind.TransactOpts, smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.contract.Transact(opts, "claimAsset", smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2Session) ClaimAsset(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.Contract.ClaimAsset(&_Ipolygonzkevmbridgev2.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2TransactorSession) ClaimAsset(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.Contract.ClaimAsset(&_Ipolygonzkevmbridgev2.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2Transactor) ClaimMessage(opts *bind.TransactOpts, smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.contract.Transact(opts, "claimMessage", smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2Session) ClaimMessage(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.Contract.ClaimMessage(&_Ipolygonzkevmbridgev2.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2TransactorSession) ClaimMessage(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.Contract.ClaimMessage(&_Ipolygonzkevmbridgev2.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2Transactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2Session) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.Contract.DeactivateEmergencyState(&_Ipolygonzkevmbridgev2.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2TransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.Contract.DeactivateEmergencyState(&_Ipolygonzkevmbridgev2.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata) returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2Transactor) Initialize(opts *bind.TransactOpts, _networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.contract.Transact(opts, "initialize", _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata)
}

// Initialize is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata) returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2Session) Initialize(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.Contract.Initialize(&_Ipolygonzkevmbridgev2.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata)
}

// Initialize is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata) returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2TransactorSession) Initialize(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte) (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.Contract.Initialize(&_Ipolygonzkevmbridgev2.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2Transactor) UpdateGlobalExitRoot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.contract.Transact(opts, "updateGlobalExitRoot")
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2Session) UpdateGlobalExitRoot() (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.Contract.UpdateGlobalExitRoot(&_Ipolygonzkevmbridgev2.TransactOpts)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Ipolygonzkevmbridgev2 *Ipolygonzkevmbridgev2TransactorSession) UpdateGlobalExitRoot() (*types.Transaction, error) {
	return _Ipolygonzkevmbridgev2.Contract.UpdateGlobalExitRoot(&_Ipolygonzkevmbridgev2.TransactOpts)
}

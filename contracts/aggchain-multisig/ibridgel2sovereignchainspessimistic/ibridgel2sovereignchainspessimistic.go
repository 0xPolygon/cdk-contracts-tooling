// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibridgel2sovereignchainspessimistic

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

// Ibridgel2sovereignchainspessimisticMetaData contains all meta data concerning the Ibridgel2sovereignchainspessimistic contract.
var Ibridgel2sovereignchainspessimisticMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountDoesNotMatchMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DestinationNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmergencyStateNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EtherTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTokenWrappedDeployment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputArraysLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSmtProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSovereignWETHAddressParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MsgValueNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeTokenIsEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoValueInMessagesOnGasTokenNetworks\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyBridgeManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OriginNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyMapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyUpdated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotMapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotRemapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WETHRemappingNotSupportedOnGasTokenNetworks\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"permitData\",\"type\":\"bytes\"}],\"name\":\"bridgeAsset\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountWETH\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessageWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_polygonRollupManager\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_bridgeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sovereignWETHAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_sovereignWETHAddressIsNotMintable\",\"type\":\"bool\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_polygonRollupManager\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"}],\"name\":\"wrappedTokenToTokenInfo\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Ibridgel2sovereignchainspessimisticABI is the input ABI used to generate the binding from.
// Deprecated: Use Ibridgel2sovereignchainspessimisticMetaData.ABI instead.
var Ibridgel2sovereignchainspessimisticABI = Ibridgel2sovereignchainspessimisticMetaData.ABI

// Ibridgel2sovereignchainspessimistic is an auto generated Go binding around an Ethereum contract.
type Ibridgel2sovereignchainspessimistic struct {
	Ibridgel2sovereignchainspessimisticCaller     // Read-only binding to the contract
	Ibridgel2sovereignchainspessimisticTransactor // Write-only binding to the contract
	Ibridgel2sovereignchainspessimisticFilterer   // Log filterer for contract events
}

// Ibridgel2sovereignchainspessimisticCaller is an auto generated read-only Go binding around an Ethereum contract.
type Ibridgel2sovereignchainspessimisticCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ibridgel2sovereignchainspessimisticTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Ibridgel2sovereignchainspessimisticTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ibridgel2sovereignchainspessimisticFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ibridgel2sovereignchainspessimisticFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ibridgel2sovereignchainspessimisticSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ibridgel2sovereignchainspessimisticSession struct {
	Contract     *Ibridgel2sovereignchainspessimistic // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                        // Call options to use throughout this session
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// Ibridgel2sovereignchainspessimisticCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ibridgel2sovereignchainspessimisticCallerSession struct {
	Contract *Ibridgel2sovereignchainspessimisticCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                              // Call options to use throughout this session
}

// Ibridgel2sovereignchainspessimisticTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ibridgel2sovereignchainspessimisticTransactorSession struct {
	Contract     *Ibridgel2sovereignchainspessimisticTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                              // Transaction auth options to use throughout this session
}

// Ibridgel2sovereignchainspessimisticRaw is an auto generated low-level Go binding around an Ethereum contract.
type Ibridgel2sovereignchainspessimisticRaw struct {
	Contract *Ibridgel2sovereignchainspessimistic // Generic contract binding to access the raw methods on
}

// Ibridgel2sovereignchainspessimisticCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ibridgel2sovereignchainspessimisticCallerRaw struct {
	Contract *Ibridgel2sovereignchainspessimisticCaller // Generic read-only contract binding to access the raw methods on
}

// Ibridgel2sovereignchainspessimisticTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ibridgel2sovereignchainspessimisticTransactorRaw struct {
	Contract *Ibridgel2sovereignchainspessimisticTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbridgel2sovereignchainspessimistic creates a new instance of Ibridgel2sovereignchainspessimistic, bound to a specific deployed contract.
func NewIbridgel2sovereignchainspessimistic(address common.Address, backend bind.ContractBackend) (*Ibridgel2sovereignchainspessimistic, error) {
	contract, err := bindIbridgel2sovereignchainspessimistic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ibridgel2sovereignchainspessimistic{Ibridgel2sovereignchainspessimisticCaller: Ibridgel2sovereignchainspessimisticCaller{contract: contract}, Ibridgel2sovereignchainspessimisticTransactor: Ibridgel2sovereignchainspessimisticTransactor{contract: contract}, Ibridgel2sovereignchainspessimisticFilterer: Ibridgel2sovereignchainspessimisticFilterer{contract: contract}}, nil
}

// NewIbridgel2sovereignchainspessimisticCaller creates a new read-only instance of Ibridgel2sovereignchainspessimistic, bound to a specific deployed contract.
func NewIbridgel2sovereignchainspessimisticCaller(address common.Address, caller bind.ContractCaller) (*Ibridgel2sovereignchainspessimisticCaller, error) {
	contract, err := bindIbridgel2sovereignchainspessimistic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ibridgel2sovereignchainspessimisticCaller{contract: contract}, nil
}

// NewIbridgel2sovereignchainspessimisticTransactor creates a new write-only instance of Ibridgel2sovereignchainspessimistic, bound to a specific deployed contract.
func NewIbridgel2sovereignchainspessimisticTransactor(address common.Address, transactor bind.ContractTransactor) (*Ibridgel2sovereignchainspessimisticTransactor, error) {
	contract, err := bindIbridgel2sovereignchainspessimistic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ibridgel2sovereignchainspessimisticTransactor{contract: contract}, nil
}

// NewIbridgel2sovereignchainspessimisticFilterer creates a new log filterer instance of Ibridgel2sovereignchainspessimistic, bound to a specific deployed contract.
func NewIbridgel2sovereignchainspessimisticFilterer(address common.Address, filterer bind.ContractFilterer) (*Ibridgel2sovereignchainspessimisticFilterer, error) {
	contract, err := bindIbridgel2sovereignchainspessimistic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ibridgel2sovereignchainspessimisticFilterer{contract: contract}, nil
}

// bindIbridgel2sovereignchainspessimistic binds a generic wrapper to an already deployed contract.
func bindIbridgel2sovereignchainspessimistic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Ibridgel2sovereignchainspessimisticMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibridgel2sovereignchainspessimistic.Contract.Ibridgel2sovereignchainspessimisticCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.Ibridgel2sovereignchainspessimisticTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.Ibridgel2sovereignchainspessimisticTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibridgel2sovereignchainspessimistic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.contract.Transact(opts, method, params...)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticCaller) GetTokenMetadata(opts *bind.CallOpts, token common.Address) ([]byte, error) {
	var out []interface{}
	err := _Ibridgel2sovereignchainspessimistic.contract.Call(opts, &out, "getTokenMetadata", token)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticSession) GetTokenMetadata(token common.Address) ([]byte, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.GetTokenMetadata(&_Ibridgel2sovereignchainspessimistic.CallOpts, token)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticCallerSession) GetTokenMetadata(token common.Address) ([]byte, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.GetTokenMetadata(&_Ibridgel2sovereignchainspessimistic.CallOpts, token)
}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address destinationAddress) view returns(uint32, address)
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticCaller) WrappedTokenToTokenInfo(opts *bind.CallOpts, destinationAddress common.Address) (uint32, common.Address, error) {
	var out []interface{}
	err := _Ibridgel2sovereignchainspessimistic.contract.Call(opts, &out, "wrappedTokenToTokenInfo", destinationAddress)

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
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticSession) WrappedTokenToTokenInfo(destinationAddress common.Address) (uint32, common.Address, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.WrappedTokenToTokenInfo(&_Ibridgel2sovereignchainspessimistic.CallOpts, destinationAddress)
}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address destinationAddress) view returns(uint32, address)
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticCallerSession) WrappedTokenToTokenInfo(destinationAddress common.Address) (uint32, common.Address, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.WrappedTokenToTokenInfo(&_Ibridgel2sovereignchainspessimistic.CallOpts, destinationAddress)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticTransactor) ActivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.contract.Transact(opts, "activateEmergencyState")
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.ActivateEmergencyState(&_Ibridgel2sovereignchainspessimistic.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticTransactorSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.ActivateEmergencyState(&_Ibridgel2sovereignchainspessimistic.TransactOpts)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticTransactor) BridgeAsset(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.contract.Transact(opts, "bridgeAsset", destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticSession) BridgeAsset(destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.BridgeAsset(&_Ibridgel2sovereignchainspessimistic.TransactOpts, destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticTransactorSession) BridgeAsset(destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.BridgeAsset(&_Ibridgel2sovereignchainspessimistic.TransactOpts, destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticTransactor) BridgeMessage(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.contract.Transact(opts, "bridgeMessage", destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticSession) BridgeMessage(destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.BridgeMessage(&_Ibridgel2sovereignchainspessimistic.TransactOpts, destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticTransactorSession) BridgeMessage(destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.BridgeMessage(&_Ibridgel2sovereignchainspessimistic.TransactOpts, destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticTransactor) BridgeMessageWETH(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.contract.Transact(opts, "bridgeMessageWETH", destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticSession) BridgeMessageWETH(destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.BridgeMessageWETH(&_Ibridgel2sovereignchainspessimistic.TransactOpts, destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticTransactorSession) BridgeMessageWETH(destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.BridgeMessageWETH(&_Ibridgel2sovereignchainspessimistic.TransactOpts, destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticTransactor) ClaimAsset(opts *bind.TransactOpts, smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.contract.Transact(opts, "claimAsset", smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticSession) ClaimAsset(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.ClaimAsset(&_Ibridgel2sovereignchainspessimistic.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticTransactorSession) ClaimAsset(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.ClaimAsset(&_Ibridgel2sovereignchainspessimistic.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticTransactor) ClaimMessage(opts *bind.TransactOpts, smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.contract.Transact(opts, "claimMessage", smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticSession) ClaimMessage(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.ClaimMessage(&_Ibridgel2sovereignchainspessimistic.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticTransactorSession) ClaimMessage(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.ClaimMessage(&_Ibridgel2sovereignchainspessimistic.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticTransactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.DeactivateEmergencyState(&_Ibridgel2sovereignchainspessimistic.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticTransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.DeactivateEmergencyState(&_Ibridgel2sovereignchainspessimistic.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8c0dd470.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable) returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticTransactor) Initialize(opts *bind.TransactOpts, _networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.contract.Transact(opts, "initialize", _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, sovereignWETHAddress, _sovereignWETHAddressIsNotMintable)
}

// Initialize is a paid mutator transaction binding the contract method 0x8c0dd470.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable) returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticSession) Initialize(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.Initialize(&_Ibridgel2sovereignchainspessimistic.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, sovereignWETHAddress, _sovereignWETHAddressIsNotMintable)
}

// Initialize is a paid mutator transaction binding the contract method 0x8c0dd470.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable) returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticTransactorSession) Initialize(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.Initialize(&_Ibridgel2sovereignchainspessimistic.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, sovereignWETHAddress, _sovereignWETHAddressIsNotMintable)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata) returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticTransactor) Initialize0(opts *bind.TransactOpts, _networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.contract.Transact(opts, "initialize0", _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata) returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticSession) Initialize0(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.Initialize0(&_Ibridgel2sovereignchainspessimistic.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata) returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticTransactorSession) Initialize0(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.Initialize0(&_Ibridgel2sovereignchainspessimistic.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticTransactor) UpdateGlobalExitRoot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.contract.Transact(opts, "updateGlobalExitRoot")
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticSession) UpdateGlobalExitRoot() (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.UpdateGlobalExitRoot(&_Ibridgel2sovereignchainspessimistic.TransactOpts)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Ibridgel2sovereignchainspessimistic *Ibridgel2sovereignchainspessimisticTransactorSession) UpdateGlobalExitRoot() (*types.Transaction, error) {
	return _Ibridgel2sovereignchainspessimistic.Contract.UpdateGlobalExitRoot(&_Ibridgel2sovereignchainspessimistic.TransactOpts)
}

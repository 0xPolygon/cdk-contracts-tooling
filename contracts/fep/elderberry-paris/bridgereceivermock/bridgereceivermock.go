// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridgereceivermock

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

// BridgereceivermockMetaData contains all meta data concerning the Bridgereceivermock contract.
var BridgereceivermockMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"indexed\":false,\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"ClaimAsset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"indexed\":false,\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"ClaimMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"calldataBytes\",\"type\":\"bytes\"}],\"name\":\"FallbackEvent\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506103c8806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063ccaa2d1114610072578063f5efcd7914610087575b7f1806efb26d87b13a8bdf58cc21831eb46678a980183a229cad8afbefd74028c7600036604051610068929190610183565b60405180910390a1005b6100856100803660046101f4565b61009a565b005b6100856100953660046101f4565b6100f5565b7ffa1453e856bf3cf1590dbcac05a3ca95a9dc5a8f419d6e1ee7e4e95edb5005b88c8c8c8c8c8c8c8c8c8c8c8c6040516100df9c9b9a999897969594939291906102fc565b60405180910390a1505050505050505050505050565b7f3f39965d0f29412dc6d768d60dd156238c12e8d6c8a15b45c337d074412a4cea8c8c8c8c8c8c8c8c8c8c8c8c6040516100df9c9b9a999897969594939291906102fc565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60208152600061019760208301848661013a565b949350505050565b8061040081018310156101b157600080fd5b92915050565b803563ffffffff811681146101cb57600080fd5b919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146101cb57600080fd5b6000806000806000806000806000806000806109208d8f03121561021757600080fd5b6102218e8e61019f565b9b506102318e6104008f0161019f565b9a506108008d013599506108208d013598506108408d013597506102586108608e016101b7565b96506102676108808e016101d0565b95506102766108a08e016101b7565b94506102856108c08e016101d0565b93506108e08d013592506109008d013567ffffffffffffffff808211156102ab57600080fd5b818f0191508f601f8301126102bf57600080fd5b80823511156102cd57600080fd5b508e6020823583010111156102e157600080fd5b60208101925080359150509295989b509295989b509295989b565b6000610400808f8437808e82850137508b6108008301528a6108208301528961084083015263ffffffff808a1661086084015273ffffffffffffffffffffffffffffffffffffffff808a166108808501528189166108a08501528088166108c08501525050846108e08301526109206109008301526103806109208301848661013a565b9e9d505050505050505050505050505056fea26469706673582212204863614ccab25b4514e5427f8923e2d1ecc271e8cadcff55cfcf77816a661fd864736f6c63430008140033",
}

// BridgereceivermockABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgereceivermockMetaData.ABI instead.
var BridgereceivermockABI = BridgereceivermockMetaData.ABI

// BridgereceivermockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgereceivermockMetaData.Bin instead.
var BridgereceivermockBin = BridgereceivermockMetaData.Bin

// DeployBridgereceivermock deploys a new Ethereum contract, binding an instance of Bridgereceivermock to it.
func DeployBridgereceivermock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bridgereceivermock, error) {
	parsed, err := BridgereceivermockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgereceivermockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridgereceivermock{BridgereceivermockCaller: BridgereceivermockCaller{contract: contract}, BridgereceivermockTransactor: BridgereceivermockTransactor{contract: contract}, BridgereceivermockFilterer: BridgereceivermockFilterer{contract: contract}}, nil
}

// Bridgereceivermock is an auto generated Go binding around an Ethereum contract.
type Bridgereceivermock struct {
	BridgereceivermockCaller     // Read-only binding to the contract
	BridgereceivermockTransactor // Write-only binding to the contract
	BridgereceivermockFilterer   // Log filterer for contract events
}

// BridgereceivermockCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgereceivermockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgereceivermockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgereceivermockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgereceivermockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgereceivermockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgereceivermockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgereceivermockSession struct {
	Contract     *Bridgereceivermock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BridgereceivermockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgereceivermockCallerSession struct {
	Contract *BridgereceivermockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// BridgereceivermockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgereceivermockTransactorSession struct {
	Contract     *BridgereceivermockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BridgereceivermockRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgereceivermockRaw struct {
	Contract *Bridgereceivermock // Generic contract binding to access the raw methods on
}

// BridgereceivermockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgereceivermockCallerRaw struct {
	Contract *BridgereceivermockCaller // Generic read-only contract binding to access the raw methods on
}

// BridgereceivermockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgereceivermockTransactorRaw struct {
	Contract *BridgereceivermockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgereceivermock creates a new instance of Bridgereceivermock, bound to a specific deployed contract.
func NewBridgereceivermock(address common.Address, backend bind.ContractBackend) (*Bridgereceivermock, error) {
	contract, err := bindBridgereceivermock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridgereceivermock{BridgereceivermockCaller: BridgereceivermockCaller{contract: contract}, BridgereceivermockTransactor: BridgereceivermockTransactor{contract: contract}, BridgereceivermockFilterer: BridgereceivermockFilterer{contract: contract}}, nil
}

// NewBridgereceivermockCaller creates a new read-only instance of Bridgereceivermock, bound to a specific deployed contract.
func NewBridgereceivermockCaller(address common.Address, caller bind.ContractCaller) (*BridgereceivermockCaller, error) {
	contract, err := bindBridgereceivermock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgereceivermockCaller{contract: contract}, nil
}

// NewBridgereceivermockTransactor creates a new write-only instance of Bridgereceivermock, bound to a specific deployed contract.
func NewBridgereceivermockTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgereceivermockTransactor, error) {
	contract, err := bindBridgereceivermock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgereceivermockTransactor{contract: contract}, nil
}

// NewBridgereceivermockFilterer creates a new log filterer instance of Bridgereceivermock, bound to a specific deployed contract.
func NewBridgereceivermockFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgereceivermockFilterer, error) {
	contract, err := bindBridgereceivermock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgereceivermockFilterer{contract: contract}, nil
}

// bindBridgereceivermock binds a generic wrapper to an already deployed contract.
func bindBridgereceivermock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgereceivermockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridgereceivermock *BridgereceivermockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridgereceivermock.Contract.BridgereceivermockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridgereceivermock *BridgereceivermockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgereceivermock.Contract.BridgereceivermockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridgereceivermock *BridgereceivermockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridgereceivermock.Contract.BridgereceivermockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridgereceivermock *BridgereceivermockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridgereceivermock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridgereceivermock *BridgereceivermockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgereceivermock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridgereceivermock *BridgereceivermockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridgereceivermock.Contract.contract.Transact(opts, method, params...)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgereceivermock *BridgereceivermockTransactor) ClaimAsset(opts *bind.TransactOpts, smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgereceivermock.contract.Transact(opts, "claimAsset", smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgereceivermock *BridgereceivermockSession) ClaimAsset(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgereceivermock.Contract.ClaimAsset(&_Bridgereceivermock.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgereceivermock *BridgereceivermockTransactorSession) ClaimAsset(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgereceivermock.Contract.ClaimAsset(&_Bridgereceivermock.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgereceivermock *BridgereceivermockTransactor) ClaimMessage(opts *bind.TransactOpts, smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgereceivermock.contract.Transact(opts, "claimMessage", smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgereceivermock *BridgereceivermockSession) ClaimMessage(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgereceivermock.Contract.ClaimMessage(&_Bridgereceivermock.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgereceivermock *BridgereceivermockTransactorSession) ClaimMessage(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgereceivermock.Contract.ClaimMessage(&_Bridgereceivermock.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Bridgereceivermock *BridgereceivermockTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Bridgereceivermock.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Bridgereceivermock *BridgereceivermockSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Bridgereceivermock.Contract.Fallback(&_Bridgereceivermock.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Bridgereceivermock *BridgereceivermockTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Bridgereceivermock.Contract.Fallback(&_Bridgereceivermock.TransactOpts, calldata)
}

// BridgereceivermockClaimAssetIterator is returned from FilterClaimAsset and is used to iterate over the raw logs and unpacked data for ClaimAsset events raised by the Bridgereceivermock contract.
type BridgereceivermockClaimAssetIterator struct {
	Event *BridgereceivermockClaimAsset // Event containing the contract specifics and raw log

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
func (it *BridgereceivermockClaimAssetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgereceivermockClaimAsset)
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
		it.Event = new(BridgereceivermockClaimAsset)
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
func (it *BridgereceivermockClaimAssetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgereceivermockClaimAssetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgereceivermockClaimAsset represents a ClaimAsset event raised by the Bridgereceivermock contract.
type BridgereceivermockClaimAsset struct {
	SmtProofLocalExitRoot  [32][32]byte
	SmtProofRollupExitRoot [32][32]byte
	GlobalIndex            *big.Int
	MainnetExitRoot        [32]byte
	RollupExitRoot         [32]byte
	OriginNetwork          uint32
	OriginTokenAddress     common.Address
	DestinationNetwork     uint32
	DestinationAddress     common.Address
	Amount                 *big.Int
	Metadata               []byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterClaimAsset is a free log retrieval operation binding the contract event 0xfa1453e856bf3cf1590dbcac05a3ca95a9dc5a8f419d6e1ee7e4e95edb5005b8.
//
// Solidity: event ClaimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata)
func (_Bridgereceivermock *BridgereceivermockFilterer) FilterClaimAsset(opts *bind.FilterOpts) (*BridgereceivermockClaimAssetIterator, error) {

	logs, sub, err := _Bridgereceivermock.contract.FilterLogs(opts, "ClaimAsset")
	if err != nil {
		return nil, err
	}
	return &BridgereceivermockClaimAssetIterator{contract: _Bridgereceivermock.contract, event: "ClaimAsset", logs: logs, sub: sub}, nil
}

// WatchClaimAsset is a free log subscription operation binding the contract event 0xfa1453e856bf3cf1590dbcac05a3ca95a9dc5a8f419d6e1ee7e4e95edb5005b8.
//
// Solidity: event ClaimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata)
func (_Bridgereceivermock *BridgereceivermockFilterer) WatchClaimAsset(opts *bind.WatchOpts, sink chan<- *BridgereceivermockClaimAsset) (event.Subscription, error) {

	logs, sub, err := _Bridgereceivermock.contract.WatchLogs(opts, "ClaimAsset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgereceivermockClaimAsset)
				if err := _Bridgereceivermock.contract.UnpackLog(event, "ClaimAsset", log); err != nil {
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

// ParseClaimAsset is a log parse operation binding the contract event 0xfa1453e856bf3cf1590dbcac05a3ca95a9dc5a8f419d6e1ee7e4e95edb5005b8.
//
// Solidity: event ClaimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata)
func (_Bridgereceivermock *BridgereceivermockFilterer) ParseClaimAsset(log types.Log) (*BridgereceivermockClaimAsset, error) {
	event := new(BridgereceivermockClaimAsset)
	if err := _Bridgereceivermock.contract.UnpackLog(event, "ClaimAsset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgereceivermockClaimMessageIterator is returned from FilterClaimMessage and is used to iterate over the raw logs and unpacked data for ClaimMessage events raised by the Bridgereceivermock contract.
type BridgereceivermockClaimMessageIterator struct {
	Event *BridgereceivermockClaimMessage // Event containing the contract specifics and raw log

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
func (it *BridgereceivermockClaimMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgereceivermockClaimMessage)
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
		it.Event = new(BridgereceivermockClaimMessage)
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
func (it *BridgereceivermockClaimMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgereceivermockClaimMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgereceivermockClaimMessage represents a ClaimMessage event raised by the Bridgereceivermock contract.
type BridgereceivermockClaimMessage struct {
	SmtProofLocalExitRoot  [32][32]byte
	SmtProofRollupExitRoot [32][32]byte
	GlobalIndex            *big.Int
	MainnetExitRoot        [32]byte
	RollupExitRoot         [32]byte
	OriginNetwork          uint32
	OriginTokenAddress     common.Address
	DestinationNetwork     uint32
	DestinationAddress     common.Address
	Amount                 *big.Int
	Metadata               []byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterClaimMessage is a free log retrieval operation binding the contract event 0x3f39965d0f29412dc6d768d60dd156238c12e8d6c8a15b45c337d074412a4cea.
//
// Solidity: event ClaimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata)
func (_Bridgereceivermock *BridgereceivermockFilterer) FilterClaimMessage(opts *bind.FilterOpts) (*BridgereceivermockClaimMessageIterator, error) {

	logs, sub, err := _Bridgereceivermock.contract.FilterLogs(opts, "ClaimMessage")
	if err != nil {
		return nil, err
	}
	return &BridgereceivermockClaimMessageIterator{contract: _Bridgereceivermock.contract, event: "ClaimMessage", logs: logs, sub: sub}, nil
}

// WatchClaimMessage is a free log subscription operation binding the contract event 0x3f39965d0f29412dc6d768d60dd156238c12e8d6c8a15b45c337d074412a4cea.
//
// Solidity: event ClaimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata)
func (_Bridgereceivermock *BridgereceivermockFilterer) WatchClaimMessage(opts *bind.WatchOpts, sink chan<- *BridgereceivermockClaimMessage) (event.Subscription, error) {

	logs, sub, err := _Bridgereceivermock.contract.WatchLogs(opts, "ClaimMessage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgereceivermockClaimMessage)
				if err := _Bridgereceivermock.contract.UnpackLog(event, "ClaimMessage", log); err != nil {
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

// ParseClaimMessage is a log parse operation binding the contract event 0x3f39965d0f29412dc6d768d60dd156238c12e8d6c8a15b45c337d074412a4cea.
//
// Solidity: event ClaimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata)
func (_Bridgereceivermock *BridgereceivermockFilterer) ParseClaimMessage(log types.Log) (*BridgereceivermockClaimMessage, error) {
	event := new(BridgereceivermockClaimMessage)
	if err := _Bridgereceivermock.contract.UnpackLog(event, "ClaimMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgereceivermockFallbackEventIterator is returned from FilterFallbackEvent and is used to iterate over the raw logs and unpacked data for FallbackEvent events raised by the Bridgereceivermock contract.
type BridgereceivermockFallbackEventIterator struct {
	Event *BridgereceivermockFallbackEvent // Event containing the contract specifics and raw log

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
func (it *BridgereceivermockFallbackEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgereceivermockFallbackEvent)
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
		it.Event = new(BridgereceivermockFallbackEvent)
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
func (it *BridgereceivermockFallbackEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgereceivermockFallbackEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgereceivermockFallbackEvent represents a FallbackEvent event raised by the Bridgereceivermock contract.
type BridgereceivermockFallbackEvent struct {
	CalldataBytes []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFallbackEvent is a free log retrieval operation binding the contract event 0x1806efb26d87b13a8bdf58cc21831eb46678a980183a229cad8afbefd74028c7.
//
// Solidity: event FallbackEvent(bytes calldataBytes)
func (_Bridgereceivermock *BridgereceivermockFilterer) FilterFallbackEvent(opts *bind.FilterOpts) (*BridgereceivermockFallbackEventIterator, error) {

	logs, sub, err := _Bridgereceivermock.contract.FilterLogs(opts, "FallbackEvent")
	if err != nil {
		return nil, err
	}
	return &BridgereceivermockFallbackEventIterator{contract: _Bridgereceivermock.contract, event: "FallbackEvent", logs: logs, sub: sub}, nil
}

// WatchFallbackEvent is a free log subscription operation binding the contract event 0x1806efb26d87b13a8bdf58cc21831eb46678a980183a229cad8afbefd74028c7.
//
// Solidity: event FallbackEvent(bytes calldataBytes)
func (_Bridgereceivermock *BridgereceivermockFilterer) WatchFallbackEvent(opts *bind.WatchOpts, sink chan<- *BridgereceivermockFallbackEvent) (event.Subscription, error) {

	logs, sub, err := _Bridgereceivermock.contract.WatchLogs(opts, "FallbackEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgereceivermockFallbackEvent)
				if err := _Bridgereceivermock.contract.UnpackLog(event, "FallbackEvent", log); err != nil {
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

// ParseFallbackEvent is a log parse operation binding the contract event 0x1806efb26d87b13a8bdf58cc21831eb46678a980183a229cad8afbefd74028c7.
//
// Solidity: event FallbackEvent(bytes calldataBytes)
func (_Bridgereceivermock *BridgereceivermockFilterer) ParseFallbackEvent(log types.Log) (*BridgereceivermockFallbackEvent, error) {
	event := new(BridgereceivermockFallbackEvent)
	if err := _Bridgereceivermock.contract.UnpackLog(event, "FallbackEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

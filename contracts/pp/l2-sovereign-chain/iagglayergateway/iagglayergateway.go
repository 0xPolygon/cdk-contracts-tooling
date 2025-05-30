// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iagglayergateway

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

// IagglayergatewayMetaData contains all meta data concerning the Iagglayergateway contract.
var IagglayergatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AggchainVKeyAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProofBytesLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAggLayerAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAggLayerAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PPSelectorCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"RouteAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"RouteIsAlreadyFrozen\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"RouteIsFrozen\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"RouteNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VKeyCannotBeZero\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newVKey\",\"type\":\"bytes32\"}],\"name\":\"AddDefaultAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pessimisticVKey\",\"type\":\"bytes32\"}],\"name\":\"RouteAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pessimisticVKey\",\"type\":\"bytes32\"}],\"name\":\"RouteFrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousVKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newVKey\",\"type\":\"bytes32\"}],\"name\":\"UpdateDefaultAggchainVKey\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"pessimisticVKeySelector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"pessimisticVKey\",\"type\":\"bytes32\"}],\"name\":\"addPessimisticVKeyRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"pessimisticVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"freezePessimisticVKeyRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"defaultAggchainSelector\",\"type\":\"bytes4\"}],\"name\":\"getDefaultAggchainVKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicValues\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofBytes\",\"type\":\"bytes\"}],\"name\":\"verifyPessimisticProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IagglayergatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use IagglayergatewayMetaData.ABI instead.
var IagglayergatewayABI = IagglayergatewayMetaData.ABI

// Iagglayergateway is an auto generated Go binding around an Ethereum contract.
type Iagglayergateway struct {
	IagglayergatewayCaller     // Read-only binding to the contract
	IagglayergatewayTransactor // Write-only binding to the contract
	IagglayergatewayFilterer   // Log filterer for contract events
}

// IagglayergatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type IagglayergatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IagglayergatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IagglayergatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IagglayergatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IagglayergatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IagglayergatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IagglayergatewaySession struct {
	Contract     *Iagglayergateway // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IagglayergatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IagglayergatewayCallerSession struct {
	Contract *IagglayergatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IagglayergatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IagglayergatewayTransactorSession struct {
	Contract     *IagglayergatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IagglayergatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type IagglayergatewayRaw struct {
	Contract *Iagglayergateway // Generic contract binding to access the raw methods on
}

// IagglayergatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IagglayergatewayCallerRaw struct {
	Contract *IagglayergatewayCaller // Generic read-only contract binding to access the raw methods on
}

// IagglayergatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IagglayergatewayTransactorRaw struct {
	Contract *IagglayergatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIagglayergateway creates a new instance of Iagglayergateway, bound to a specific deployed contract.
func NewIagglayergateway(address common.Address, backend bind.ContractBackend) (*Iagglayergateway, error) {
	contract, err := bindIagglayergateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iagglayergateway{IagglayergatewayCaller: IagglayergatewayCaller{contract: contract}, IagglayergatewayTransactor: IagglayergatewayTransactor{contract: contract}, IagglayergatewayFilterer: IagglayergatewayFilterer{contract: contract}}, nil
}

// NewIagglayergatewayCaller creates a new read-only instance of Iagglayergateway, bound to a specific deployed contract.
func NewIagglayergatewayCaller(address common.Address, caller bind.ContractCaller) (*IagglayergatewayCaller, error) {
	contract, err := bindIagglayergateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayCaller{contract: contract}, nil
}

// NewIagglayergatewayTransactor creates a new write-only instance of Iagglayergateway, bound to a specific deployed contract.
func NewIagglayergatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*IagglayergatewayTransactor, error) {
	contract, err := bindIagglayergateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayTransactor{contract: contract}, nil
}

// NewIagglayergatewayFilterer creates a new log filterer instance of Iagglayergateway, bound to a specific deployed contract.
func NewIagglayergatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*IagglayergatewayFilterer, error) {
	contract, err := bindIagglayergateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayFilterer{contract: contract}, nil
}

// bindIagglayergateway binds a generic wrapper to an already deployed contract.
func bindIagglayergateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IagglayergatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iagglayergateway *IagglayergatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iagglayergateway.Contract.IagglayergatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iagglayergateway *IagglayergatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iagglayergateway.Contract.IagglayergatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iagglayergateway *IagglayergatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iagglayergateway.Contract.IagglayergatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iagglayergateway *IagglayergatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iagglayergateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iagglayergateway *IagglayergatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iagglayergateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iagglayergateway *IagglayergatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iagglayergateway.Contract.contract.Transact(opts, method, params...)
}

// GetDefaultAggchainVKey is a free data retrieval call binding the contract method 0x6cabfdab.
//
// Solidity: function getDefaultAggchainVKey(bytes4 defaultAggchainSelector) view returns(bytes32)
func (_Iagglayergateway *IagglayergatewayCaller) GetDefaultAggchainVKey(opts *bind.CallOpts, defaultAggchainSelector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _Iagglayergateway.contract.Call(opts, &out, "getDefaultAggchainVKey", defaultAggchainSelector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDefaultAggchainVKey is a free data retrieval call binding the contract method 0x6cabfdab.
//
// Solidity: function getDefaultAggchainVKey(bytes4 defaultAggchainSelector) view returns(bytes32)
func (_Iagglayergateway *IagglayergatewaySession) GetDefaultAggchainVKey(defaultAggchainSelector [4]byte) ([32]byte, error) {
	return _Iagglayergateway.Contract.GetDefaultAggchainVKey(&_Iagglayergateway.CallOpts, defaultAggchainSelector)
}

// GetDefaultAggchainVKey is a free data retrieval call binding the contract method 0x6cabfdab.
//
// Solidity: function getDefaultAggchainVKey(bytes4 defaultAggchainSelector) view returns(bytes32)
func (_Iagglayergateway *IagglayergatewayCallerSession) GetDefaultAggchainVKey(defaultAggchainSelector [4]byte) ([32]byte, error) {
	return _Iagglayergateway.Contract.GetDefaultAggchainVKey(&_Iagglayergateway.CallOpts, defaultAggchainSelector)
}

// VerifyPessimisticProof is a free data retrieval call binding the contract method 0xa48fd377.
//
// Solidity: function verifyPessimisticProof(bytes publicValues, bytes proofBytes) view returns()
func (_Iagglayergateway *IagglayergatewayCaller) VerifyPessimisticProof(opts *bind.CallOpts, publicValues []byte, proofBytes []byte) error {
	var out []interface{}
	err := _Iagglayergateway.contract.Call(opts, &out, "verifyPessimisticProof", publicValues, proofBytes)

	if err != nil {
		return err
	}

	return err

}

// VerifyPessimisticProof is a free data retrieval call binding the contract method 0xa48fd377.
//
// Solidity: function verifyPessimisticProof(bytes publicValues, bytes proofBytes) view returns()
func (_Iagglayergateway *IagglayergatewaySession) VerifyPessimisticProof(publicValues []byte, proofBytes []byte) error {
	return _Iagglayergateway.Contract.VerifyPessimisticProof(&_Iagglayergateway.CallOpts, publicValues, proofBytes)
}

// VerifyPessimisticProof is a free data retrieval call binding the contract method 0xa48fd377.
//
// Solidity: function verifyPessimisticProof(bytes publicValues, bytes proofBytes) view returns()
func (_Iagglayergateway *IagglayergatewayCallerSession) VerifyPessimisticProof(publicValues []byte, proofBytes []byte) error {
	return _Iagglayergateway.Contract.VerifyPessimisticProof(&_Iagglayergateway.CallOpts, publicValues, proofBytes)
}

// AddPessimisticVKeyRoute is a paid mutator transaction binding the contract method 0x9584a516.
//
// Solidity: function addPessimisticVKeyRoute(bytes4 pessimisticVKeySelector, address verifier, bytes32 pessimisticVKey) returns()
func (_Iagglayergateway *IagglayergatewayTransactor) AddPessimisticVKeyRoute(opts *bind.TransactOpts, pessimisticVKeySelector [4]byte, verifier common.Address, pessimisticVKey [32]byte) (*types.Transaction, error) {
	return _Iagglayergateway.contract.Transact(opts, "addPessimisticVKeyRoute", pessimisticVKeySelector, verifier, pessimisticVKey)
}

// AddPessimisticVKeyRoute is a paid mutator transaction binding the contract method 0x9584a516.
//
// Solidity: function addPessimisticVKeyRoute(bytes4 pessimisticVKeySelector, address verifier, bytes32 pessimisticVKey) returns()
func (_Iagglayergateway *IagglayergatewaySession) AddPessimisticVKeyRoute(pessimisticVKeySelector [4]byte, verifier common.Address, pessimisticVKey [32]byte) (*types.Transaction, error) {
	return _Iagglayergateway.Contract.AddPessimisticVKeyRoute(&_Iagglayergateway.TransactOpts, pessimisticVKeySelector, verifier, pessimisticVKey)
}

// AddPessimisticVKeyRoute is a paid mutator transaction binding the contract method 0x9584a516.
//
// Solidity: function addPessimisticVKeyRoute(bytes4 pessimisticVKeySelector, address verifier, bytes32 pessimisticVKey) returns()
func (_Iagglayergateway *IagglayergatewayTransactorSession) AddPessimisticVKeyRoute(pessimisticVKeySelector [4]byte, verifier common.Address, pessimisticVKey [32]byte) (*types.Transaction, error) {
	return _Iagglayergateway.Contract.AddPessimisticVKeyRoute(&_Iagglayergateway.TransactOpts, pessimisticVKeySelector, verifier, pessimisticVKey)
}

// FreezePessimisticVKeyRoute is a paid mutator transaction binding the contract method 0x95d2a391.
//
// Solidity: function freezePessimisticVKeyRoute(bytes4 pessimisticVKeySelector) returns()
func (_Iagglayergateway *IagglayergatewayTransactor) FreezePessimisticVKeyRoute(opts *bind.TransactOpts, pessimisticVKeySelector [4]byte) (*types.Transaction, error) {
	return _Iagglayergateway.contract.Transact(opts, "freezePessimisticVKeyRoute", pessimisticVKeySelector)
}

// FreezePessimisticVKeyRoute is a paid mutator transaction binding the contract method 0x95d2a391.
//
// Solidity: function freezePessimisticVKeyRoute(bytes4 pessimisticVKeySelector) returns()
func (_Iagglayergateway *IagglayergatewaySession) FreezePessimisticVKeyRoute(pessimisticVKeySelector [4]byte) (*types.Transaction, error) {
	return _Iagglayergateway.Contract.FreezePessimisticVKeyRoute(&_Iagglayergateway.TransactOpts, pessimisticVKeySelector)
}

// FreezePessimisticVKeyRoute is a paid mutator transaction binding the contract method 0x95d2a391.
//
// Solidity: function freezePessimisticVKeyRoute(bytes4 pessimisticVKeySelector) returns()
func (_Iagglayergateway *IagglayergatewayTransactorSession) FreezePessimisticVKeyRoute(pessimisticVKeySelector [4]byte) (*types.Transaction, error) {
	return _Iagglayergateway.Contract.FreezePessimisticVKeyRoute(&_Iagglayergateway.TransactOpts, pessimisticVKeySelector)
}

// IagglayergatewayAddDefaultAggchainVKeyIterator is returned from FilterAddDefaultAggchainVKey and is used to iterate over the raw logs and unpacked data for AddDefaultAggchainVKey events raised by the Iagglayergateway contract.
type IagglayergatewayAddDefaultAggchainVKeyIterator struct {
	Event *IagglayergatewayAddDefaultAggchainVKey // Event containing the contract specifics and raw log

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
func (it *IagglayergatewayAddDefaultAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IagglayergatewayAddDefaultAggchainVKey)
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
		it.Event = new(IagglayergatewayAddDefaultAggchainVKey)
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
func (it *IagglayergatewayAddDefaultAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IagglayergatewayAddDefaultAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IagglayergatewayAddDefaultAggchainVKey represents a AddDefaultAggchainVKey event raised by the Iagglayergateway contract.
type IagglayergatewayAddDefaultAggchainVKey struct {
	Selector [4]byte
	NewVKey  [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddDefaultAggchainVKey is a free log retrieval operation binding the contract event 0x64b3528974574a41ddc830adc96440dccb99e64ede6b863194b951bfbd17e1bc.
//
// Solidity: event AddDefaultAggchainVKey(bytes4 selector, bytes32 newVKey)
func (_Iagglayergateway *IagglayergatewayFilterer) FilterAddDefaultAggchainVKey(opts *bind.FilterOpts) (*IagglayergatewayAddDefaultAggchainVKeyIterator, error) {

	logs, sub, err := _Iagglayergateway.contract.FilterLogs(opts, "AddDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayAddDefaultAggchainVKeyIterator{contract: _Iagglayergateway.contract, event: "AddDefaultAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchAddDefaultAggchainVKey is a free log subscription operation binding the contract event 0x64b3528974574a41ddc830adc96440dccb99e64ede6b863194b951bfbd17e1bc.
//
// Solidity: event AddDefaultAggchainVKey(bytes4 selector, bytes32 newVKey)
func (_Iagglayergateway *IagglayergatewayFilterer) WatchAddDefaultAggchainVKey(opts *bind.WatchOpts, sink chan<- *IagglayergatewayAddDefaultAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Iagglayergateway.contract.WatchLogs(opts, "AddDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IagglayergatewayAddDefaultAggchainVKey)
				if err := _Iagglayergateway.contract.UnpackLog(event, "AddDefaultAggchainVKey", log); err != nil {
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

// ParseAddDefaultAggchainVKey is a log parse operation binding the contract event 0x64b3528974574a41ddc830adc96440dccb99e64ede6b863194b951bfbd17e1bc.
//
// Solidity: event AddDefaultAggchainVKey(bytes4 selector, bytes32 newVKey)
func (_Iagglayergateway *IagglayergatewayFilterer) ParseAddDefaultAggchainVKey(log types.Log) (*IagglayergatewayAddDefaultAggchainVKey, error) {
	event := new(IagglayergatewayAddDefaultAggchainVKey)
	if err := _Iagglayergateway.contract.UnpackLog(event, "AddDefaultAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IagglayergatewayRouteAddedIterator is returned from FilterRouteAdded and is used to iterate over the raw logs and unpacked data for RouteAdded events raised by the Iagglayergateway contract.
type IagglayergatewayRouteAddedIterator struct {
	Event *IagglayergatewayRouteAdded // Event containing the contract specifics and raw log

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
func (it *IagglayergatewayRouteAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IagglayergatewayRouteAdded)
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
		it.Event = new(IagglayergatewayRouteAdded)
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
func (it *IagglayergatewayRouteAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IagglayergatewayRouteAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IagglayergatewayRouteAdded represents a RouteAdded event raised by the Iagglayergateway contract.
type IagglayergatewayRouteAdded struct {
	Selector        [4]byte
	Verifier        common.Address
	PessimisticVKey [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRouteAdded is a free log retrieval operation binding the contract event 0x4c4455a30528c319643ae6fc35f3a0fcabde6c015f2a1a5270382c4a190b0fa3.
//
// Solidity: event RouteAdded(bytes4 selector, address verifier, bytes32 pessimisticVKey)
func (_Iagglayergateway *IagglayergatewayFilterer) FilterRouteAdded(opts *bind.FilterOpts) (*IagglayergatewayRouteAddedIterator, error) {

	logs, sub, err := _Iagglayergateway.contract.FilterLogs(opts, "RouteAdded")
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayRouteAddedIterator{contract: _Iagglayergateway.contract, event: "RouteAdded", logs: logs, sub: sub}, nil
}

// WatchRouteAdded is a free log subscription operation binding the contract event 0x4c4455a30528c319643ae6fc35f3a0fcabde6c015f2a1a5270382c4a190b0fa3.
//
// Solidity: event RouteAdded(bytes4 selector, address verifier, bytes32 pessimisticVKey)
func (_Iagglayergateway *IagglayergatewayFilterer) WatchRouteAdded(opts *bind.WatchOpts, sink chan<- *IagglayergatewayRouteAdded) (event.Subscription, error) {

	logs, sub, err := _Iagglayergateway.contract.WatchLogs(opts, "RouteAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IagglayergatewayRouteAdded)
				if err := _Iagglayergateway.contract.UnpackLog(event, "RouteAdded", log); err != nil {
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

// ParseRouteAdded is a log parse operation binding the contract event 0x4c4455a30528c319643ae6fc35f3a0fcabde6c015f2a1a5270382c4a190b0fa3.
//
// Solidity: event RouteAdded(bytes4 selector, address verifier, bytes32 pessimisticVKey)
func (_Iagglayergateway *IagglayergatewayFilterer) ParseRouteAdded(log types.Log) (*IagglayergatewayRouteAdded, error) {
	event := new(IagglayergatewayRouteAdded)
	if err := _Iagglayergateway.contract.UnpackLog(event, "RouteAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IagglayergatewayRouteFrozenIterator is returned from FilterRouteFrozen and is used to iterate over the raw logs and unpacked data for RouteFrozen events raised by the Iagglayergateway contract.
type IagglayergatewayRouteFrozenIterator struct {
	Event *IagglayergatewayRouteFrozen // Event containing the contract specifics and raw log

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
func (it *IagglayergatewayRouteFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IagglayergatewayRouteFrozen)
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
		it.Event = new(IagglayergatewayRouteFrozen)
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
func (it *IagglayergatewayRouteFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IagglayergatewayRouteFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IagglayergatewayRouteFrozen represents a RouteFrozen event raised by the Iagglayergateway contract.
type IagglayergatewayRouteFrozen struct {
	Selector        [4]byte
	Verifier        common.Address
	PessimisticVKey [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRouteFrozen is a free log retrieval operation binding the contract event 0x6efcfd4090b998a9d58c7f753bdda3291e9d781c02226b9f480b18ef07e0012b.
//
// Solidity: event RouteFrozen(bytes4 selector, address verifier, bytes32 pessimisticVKey)
func (_Iagglayergateway *IagglayergatewayFilterer) FilterRouteFrozen(opts *bind.FilterOpts) (*IagglayergatewayRouteFrozenIterator, error) {

	logs, sub, err := _Iagglayergateway.contract.FilterLogs(opts, "RouteFrozen")
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayRouteFrozenIterator{contract: _Iagglayergateway.contract, event: "RouteFrozen", logs: logs, sub: sub}, nil
}

// WatchRouteFrozen is a free log subscription operation binding the contract event 0x6efcfd4090b998a9d58c7f753bdda3291e9d781c02226b9f480b18ef07e0012b.
//
// Solidity: event RouteFrozen(bytes4 selector, address verifier, bytes32 pessimisticVKey)
func (_Iagglayergateway *IagglayergatewayFilterer) WatchRouteFrozen(opts *bind.WatchOpts, sink chan<- *IagglayergatewayRouteFrozen) (event.Subscription, error) {

	logs, sub, err := _Iagglayergateway.contract.WatchLogs(opts, "RouteFrozen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IagglayergatewayRouteFrozen)
				if err := _Iagglayergateway.contract.UnpackLog(event, "RouteFrozen", log); err != nil {
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

// ParseRouteFrozen is a log parse operation binding the contract event 0x6efcfd4090b998a9d58c7f753bdda3291e9d781c02226b9f480b18ef07e0012b.
//
// Solidity: event RouteFrozen(bytes4 selector, address verifier, bytes32 pessimisticVKey)
func (_Iagglayergateway *IagglayergatewayFilterer) ParseRouteFrozen(log types.Log) (*IagglayergatewayRouteFrozen, error) {
	event := new(IagglayergatewayRouteFrozen)
	if err := _Iagglayergateway.contract.UnpackLog(event, "RouteFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IagglayergatewayUpdateDefaultAggchainVKeyIterator is returned from FilterUpdateDefaultAggchainVKey and is used to iterate over the raw logs and unpacked data for UpdateDefaultAggchainVKey events raised by the Iagglayergateway contract.
type IagglayergatewayUpdateDefaultAggchainVKeyIterator struct {
	Event *IagglayergatewayUpdateDefaultAggchainVKey // Event containing the contract specifics and raw log

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
func (it *IagglayergatewayUpdateDefaultAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IagglayergatewayUpdateDefaultAggchainVKey)
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
		it.Event = new(IagglayergatewayUpdateDefaultAggchainVKey)
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
func (it *IagglayergatewayUpdateDefaultAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IagglayergatewayUpdateDefaultAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IagglayergatewayUpdateDefaultAggchainVKey represents a UpdateDefaultAggchainVKey event raised by the Iagglayergateway contract.
type IagglayergatewayUpdateDefaultAggchainVKey struct {
	Selector     [4]byte
	PreviousVKey [32]byte
	NewVKey      [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateDefaultAggchainVKey is a free log retrieval operation binding the contract event 0x477ac863279848383fdb5d6bb8957cc8f6f10c0200379e075df3f789d14248e6.
//
// Solidity: event UpdateDefaultAggchainVKey(bytes4 selector, bytes32 previousVKey, bytes32 newVKey)
func (_Iagglayergateway *IagglayergatewayFilterer) FilterUpdateDefaultAggchainVKey(opts *bind.FilterOpts) (*IagglayergatewayUpdateDefaultAggchainVKeyIterator, error) {

	logs, sub, err := _Iagglayergateway.contract.FilterLogs(opts, "UpdateDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayUpdateDefaultAggchainVKeyIterator{contract: _Iagglayergateway.contract, event: "UpdateDefaultAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchUpdateDefaultAggchainVKey is a free log subscription operation binding the contract event 0x477ac863279848383fdb5d6bb8957cc8f6f10c0200379e075df3f789d14248e6.
//
// Solidity: event UpdateDefaultAggchainVKey(bytes4 selector, bytes32 previousVKey, bytes32 newVKey)
func (_Iagglayergateway *IagglayergatewayFilterer) WatchUpdateDefaultAggchainVKey(opts *bind.WatchOpts, sink chan<- *IagglayergatewayUpdateDefaultAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Iagglayergateway.contract.WatchLogs(opts, "UpdateDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IagglayergatewayUpdateDefaultAggchainVKey)
				if err := _Iagglayergateway.contract.UnpackLog(event, "UpdateDefaultAggchainVKey", log); err != nil {
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

// ParseUpdateDefaultAggchainVKey is a log parse operation binding the contract event 0x477ac863279848383fdb5d6bb8957cc8f6f10c0200379e075df3f789d14248e6.
//
// Solidity: event UpdateDefaultAggchainVKey(bytes4 selector, bytes32 previousVKey, bytes32 newVKey)
func (_Iagglayergateway *IagglayergatewayFilterer) ParseUpdateDefaultAggchainVKey(log types.Log) (*IagglayergatewayUpdateDefaultAggchainVKey, error) {
	event := new(IagglayergatewayUpdateDefaultAggchainVKey)
	if err := _Iagglayergateway.contract.UnpackLog(event, "UpdateDefaultAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

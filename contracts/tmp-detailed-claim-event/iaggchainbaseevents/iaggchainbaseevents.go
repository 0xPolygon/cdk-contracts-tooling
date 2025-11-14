// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iaggchainbaseevents

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

// IaggchainbaseeventsMetaData contains all meta data concerning the Iaggchainbaseevents contract.
var IaggchainbaseeventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAggchainManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"AcceptAggchainManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"AddAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"AggchainMetadataSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DisableUseDefaultSignersFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DisableUseDefaultVkeysFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EnableUseDefaultSignersFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EnableUseDefaultVkeysFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAggchainMetadataManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAggchainMetadataManager\",\"type\":\"address\"}],\"name\":\"SetAggchainMetadataManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentAggchainManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAggchainManager\",\"type\":\"address\"}],\"name\":\"TransferAggchainManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousAggchainVKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"UpdateAggchainVKey\",\"type\":\"event\"}]",
}

// IaggchainbaseeventsABI is the input ABI used to generate the binding from.
// Deprecated: Use IaggchainbaseeventsMetaData.ABI instead.
var IaggchainbaseeventsABI = IaggchainbaseeventsMetaData.ABI

// Iaggchainbaseevents is an auto generated Go binding around an Ethereum contract.
type Iaggchainbaseevents struct {
	IaggchainbaseeventsCaller     // Read-only binding to the contract
	IaggchainbaseeventsTransactor // Write-only binding to the contract
	IaggchainbaseeventsFilterer   // Log filterer for contract events
}

// IaggchainbaseeventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IaggchainbaseeventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggchainbaseeventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IaggchainbaseeventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggchainbaseeventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IaggchainbaseeventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggchainbaseeventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IaggchainbaseeventsSession struct {
	Contract     *Iaggchainbaseevents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IaggchainbaseeventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IaggchainbaseeventsCallerSession struct {
	Contract *IaggchainbaseeventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IaggchainbaseeventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IaggchainbaseeventsTransactorSession struct {
	Contract     *IaggchainbaseeventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IaggchainbaseeventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IaggchainbaseeventsRaw struct {
	Contract *Iaggchainbaseevents // Generic contract binding to access the raw methods on
}

// IaggchainbaseeventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IaggchainbaseeventsCallerRaw struct {
	Contract *IaggchainbaseeventsCaller // Generic read-only contract binding to access the raw methods on
}

// IaggchainbaseeventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IaggchainbaseeventsTransactorRaw struct {
	Contract *IaggchainbaseeventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIaggchainbaseevents creates a new instance of Iaggchainbaseevents, bound to a specific deployed contract.
func NewIaggchainbaseevents(address common.Address, backend bind.ContractBackend) (*Iaggchainbaseevents, error) {
	contract, err := bindIaggchainbaseevents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iaggchainbaseevents{IaggchainbaseeventsCaller: IaggchainbaseeventsCaller{contract: contract}, IaggchainbaseeventsTransactor: IaggchainbaseeventsTransactor{contract: contract}, IaggchainbaseeventsFilterer: IaggchainbaseeventsFilterer{contract: contract}}, nil
}

// NewIaggchainbaseeventsCaller creates a new read-only instance of Iaggchainbaseevents, bound to a specific deployed contract.
func NewIaggchainbaseeventsCaller(address common.Address, caller bind.ContractCaller) (*IaggchainbaseeventsCaller, error) {
	contract, err := bindIaggchainbaseevents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventsCaller{contract: contract}, nil
}

// NewIaggchainbaseeventsTransactor creates a new write-only instance of Iaggchainbaseevents, bound to a specific deployed contract.
func NewIaggchainbaseeventsTransactor(address common.Address, transactor bind.ContractTransactor) (*IaggchainbaseeventsTransactor, error) {
	contract, err := bindIaggchainbaseevents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventsTransactor{contract: contract}, nil
}

// NewIaggchainbaseeventsFilterer creates a new log filterer instance of Iaggchainbaseevents, bound to a specific deployed contract.
func NewIaggchainbaseeventsFilterer(address common.Address, filterer bind.ContractFilterer) (*IaggchainbaseeventsFilterer, error) {
	contract, err := bindIaggchainbaseevents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventsFilterer{contract: contract}, nil
}

// bindIaggchainbaseevents binds a generic wrapper to an already deployed contract.
func bindIaggchainbaseevents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IaggchainbaseeventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iaggchainbaseevents *IaggchainbaseeventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iaggchainbaseevents.Contract.IaggchainbaseeventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iaggchainbaseevents *IaggchainbaseeventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iaggchainbaseevents.Contract.IaggchainbaseeventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iaggchainbaseevents *IaggchainbaseeventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iaggchainbaseevents.Contract.IaggchainbaseeventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iaggchainbaseevents *IaggchainbaseeventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iaggchainbaseevents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iaggchainbaseevents *IaggchainbaseeventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iaggchainbaseevents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iaggchainbaseevents *IaggchainbaseeventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iaggchainbaseevents.Contract.contract.Transact(opts, method, params...)
}

// IaggchainbaseeventsAcceptAggchainManagerRoleIterator is returned from FilterAcceptAggchainManagerRole and is used to iterate over the raw logs and unpacked data for AcceptAggchainManagerRole events raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsAcceptAggchainManagerRoleIterator struct {
	Event *IaggchainbaseeventsAcceptAggchainManagerRole // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseeventsAcceptAggchainManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseeventsAcceptAggchainManagerRole)
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
		it.Event = new(IaggchainbaseeventsAcceptAggchainManagerRole)
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
func (it *IaggchainbaseeventsAcceptAggchainManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseeventsAcceptAggchainManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseeventsAcceptAggchainManagerRole represents a AcceptAggchainManagerRole event raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsAcceptAggchainManagerRole struct {
	OldAggchainManager common.Address
	NewAggchainManager common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAcceptAggchainManagerRole is a free log retrieval operation binding the contract event 0x67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f.
//
// Solidity: event AcceptAggchainManagerRole(address oldAggchainManager, address newAggchainManager)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) FilterAcceptAggchainManagerRole(opts *bind.FilterOpts) (*IaggchainbaseeventsAcceptAggchainManagerRoleIterator, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.FilterLogs(opts, "AcceptAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventsAcceptAggchainManagerRoleIterator{contract: _Iaggchainbaseevents.contract, event: "AcceptAggchainManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAggchainManagerRole is a free log subscription operation binding the contract event 0x67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f.
//
// Solidity: event AcceptAggchainManagerRole(address oldAggchainManager, address newAggchainManager)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) WatchAcceptAggchainManagerRole(opts *bind.WatchOpts, sink chan<- *IaggchainbaseeventsAcceptAggchainManagerRole) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.WatchLogs(opts, "AcceptAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseeventsAcceptAggchainManagerRole)
				if err := _Iaggchainbaseevents.contract.UnpackLog(event, "AcceptAggchainManagerRole", log); err != nil {
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

// ParseAcceptAggchainManagerRole is a log parse operation binding the contract event 0x67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f.
//
// Solidity: event AcceptAggchainManagerRole(address oldAggchainManager, address newAggchainManager)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) ParseAcceptAggchainManagerRole(log types.Log) (*IaggchainbaseeventsAcceptAggchainManagerRole, error) {
	event := new(IaggchainbaseeventsAcceptAggchainManagerRole)
	if err := _Iaggchainbaseevents.contract.UnpackLog(event, "AcceptAggchainManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseeventsAddAggchainVKeyIterator is returned from FilterAddAggchainVKey and is used to iterate over the raw logs and unpacked data for AddAggchainVKey events raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsAddAggchainVKeyIterator struct {
	Event *IaggchainbaseeventsAddAggchainVKey // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseeventsAddAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseeventsAddAggchainVKey)
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
		it.Event = new(IaggchainbaseeventsAddAggchainVKey)
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
func (it *IaggchainbaseeventsAddAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseeventsAddAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseeventsAddAggchainVKey represents a AddAggchainVKey event raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsAddAggchainVKey struct {
	Selector        [4]byte
	NewAggchainVKey [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddAggchainVKey is a free log retrieval operation binding the contract event 0x6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a79.
//
// Solidity: event AddAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) FilterAddAggchainVKey(opts *bind.FilterOpts) (*IaggchainbaseeventsAddAggchainVKeyIterator, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.FilterLogs(opts, "AddAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventsAddAggchainVKeyIterator{contract: _Iaggchainbaseevents.contract, event: "AddAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchAddAggchainVKey is a free log subscription operation binding the contract event 0x6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a79.
//
// Solidity: event AddAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) WatchAddAggchainVKey(opts *bind.WatchOpts, sink chan<- *IaggchainbaseeventsAddAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.WatchLogs(opts, "AddAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseeventsAddAggchainVKey)
				if err := _Iaggchainbaseevents.contract.UnpackLog(event, "AddAggchainVKey", log); err != nil {
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

// ParseAddAggchainVKey is a log parse operation binding the contract event 0x6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a79.
//
// Solidity: event AddAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) ParseAddAggchainVKey(log types.Log) (*IaggchainbaseeventsAddAggchainVKey, error) {
	event := new(IaggchainbaseeventsAddAggchainVKey)
	if err := _Iaggchainbaseevents.contract.UnpackLog(event, "AddAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseeventsAggchainMetadataSetIterator is returned from FilterAggchainMetadataSet and is used to iterate over the raw logs and unpacked data for AggchainMetadataSet events raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsAggchainMetadataSetIterator struct {
	Event *IaggchainbaseeventsAggchainMetadataSet // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseeventsAggchainMetadataSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseeventsAggchainMetadataSet)
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
		it.Event = new(IaggchainbaseeventsAggchainMetadataSet)
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
func (it *IaggchainbaseeventsAggchainMetadataSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseeventsAggchainMetadataSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseeventsAggchainMetadataSet represents a AggchainMetadataSet event raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsAggchainMetadataSet struct {
	Key   common.Hash
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAggchainMetadataSet is a free log retrieval operation binding the contract event 0x2779f9edd5ec4e0a99bffdea4008c8b979200959062a2bf00142acb939ca1b64.
//
// Solidity: event AggchainMetadataSet(string indexed key, string value)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) FilterAggchainMetadataSet(opts *bind.FilterOpts, key []string) (*IaggchainbaseeventsAggchainMetadataSetIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Iaggchainbaseevents.contract.FilterLogs(opts, "AggchainMetadataSet", keyRule)
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventsAggchainMetadataSetIterator{contract: _Iaggchainbaseevents.contract, event: "AggchainMetadataSet", logs: logs, sub: sub}, nil
}

// WatchAggchainMetadataSet is a free log subscription operation binding the contract event 0x2779f9edd5ec4e0a99bffdea4008c8b979200959062a2bf00142acb939ca1b64.
//
// Solidity: event AggchainMetadataSet(string indexed key, string value)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) WatchAggchainMetadataSet(opts *bind.WatchOpts, sink chan<- *IaggchainbaseeventsAggchainMetadataSet, key []string) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Iaggchainbaseevents.contract.WatchLogs(opts, "AggchainMetadataSet", keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseeventsAggchainMetadataSet)
				if err := _Iaggchainbaseevents.contract.UnpackLog(event, "AggchainMetadataSet", log); err != nil {
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

// ParseAggchainMetadataSet is a log parse operation binding the contract event 0x2779f9edd5ec4e0a99bffdea4008c8b979200959062a2bf00142acb939ca1b64.
//
// Solidity: event AggchainMetadataSet(string indexed key, string value)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) ParseAggchainMetadataSet(log types.Log) (*IaggchainbaseeventsAggchainMetadataSet, error) {
	event := new(IaggchainbaseeventsAggchainMetadataSet)
	if err := _Iaggchainbaseevents.contract.UnpackLog(event, "AggchainMetadataSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseeventsDisableUseDefaultSignersFlagIterator is returned from FilterDisableUseDefaultSignersFlag and is used to iterate over the raw logs and unpacked data for DisableUseDefaultSignersFlag events raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsDisableUseDefaultSignersFlagIterator struct {
	Event *IaggchainbaseeventsDisableUseDefaultSignersFlag // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseeventsDisableUseDefaultSignersFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseeventsDisableUseDefaultSignersFlag)
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
		it.Event = new(IaggchainbaseeventsDisableUseDefaultSignersFlag)
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
func (it *IaggchainbaseeventsDisableUseDefaultSignersFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseeventsDisableUseDefaultSignersFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseeventsDisableUseDefaultSignersFlag represents a DisableUseDefaultSignersFlag event raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsDisableUseDefaultSignersFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisableUseDefaultSignersFlag is a free log retrieval operation binding the contract event 0x4c75580a56c734245a7418eb07d8a311e1bff79f982fed747da3589630e414be.
//
// Solidity: event DisableUseDefaultSignersFlag()
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) FilterDisableUseDefaultSignersFlag(opts *bind.FilterOpts) (*IaggchainbaseeventsDisableUseDefaultSignersFlagIterator, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.FilterLogs(opts, "DisableUseDefaultSignersFlag")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventsDisableUseDefaultSignersFlagIterator{contract: _Iaggchainbaseevents.contract, event: "DisableUseDefaultSignersFlag", logs: logs, sub: sub}, nil
}

// WatchDisableUseDefaultSignersFlag is a free log subscription operation binding the contract event 0x4c75580a56c734245a7418eb07d8a311e1bff79f982fed747da3589630e414be.
//
// Solidity: event DisableUseDefaultSignersFlag()
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) WatchDisableUseDefaultSignersFlag(opts *bind.WatchOpts, sink chan<- *IaggchainbaseeventsDisableUseDefaultSignersFlag) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.WatchLogs(opts, "DisableUseDefaultSignersFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseeventsDisableUseDefaultSignersFlag)
				if err := _Iaggchainbaseevents.contract.UnpackLog(event, "DisableUseDefaultSignersFlag", log); err != nil {
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

// ParseDisableUseDefaultSignersFlag is a log parse operation binding the contract event 0x4c75580a56c734245a7418eb07d8a311e1bff79f982fed747da3589630e414be.
//
// Solidity: event DisableUseDefaultSignersFlag()
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) ParseDisableUseDefaultSignersFlag(log types.Log) (*IaggchainbaseeventsDisableUseDefaultSignersFlag, error) {
	event := new(IaggchainbaseeventsDisableUseDefaultSignersFlag)
	if err := _Iaggchainbaseevents.contract.UnpackLog(event, "DisableUseDefaultSignersFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseeventsDisableUseDefaultVkeysFlagIterator is returned from FilterDisableUseDefaultVkeysFlag and is used to iterate over the raw logs and unpacked data for DisableUseDefaultVkeysFlag events raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsDisableUseDefaultVkeysFlagIterator struct {
	Event *IaggchainbaseeventsDisableUseDefaultVkeysFlag // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseeventsDisableUseDefaultVkeysFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseeventsDisableUseDefaultVkeysFlag)
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
		it.Event = new(IaggchainbaseeventsDisableUseDefaultVkeysFlag)
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
func (it *IaggchainbaseeventsDisableUseDefaultVkeysFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseeventsDisableUseDefaultVkeysFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseeventsDisableUseDefaultVkeysFlag represents a DisableUseDefaultVkeysFlag event raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsDisableUseDefaultVkeysFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisableUseDefaultVkeysFlag is a free log retrieval operation binding the contract event 0x922aeecd3505b46415820aae489ed9dac9e250e74d497b14c33e8360b581ac07.
//
// Solidity: event DisableUseDefaultVkeysFlag()
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) FilterDisableUseDefaultVkeysFlag(opts *bind.FilterOpts) (*IaggchainbaseeventsDisableUseDefaultVkeysFlagIterator, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.FilterLogs(opts, "DisableUseDefaultVkeysFlag")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventsDisableUseDefaultVkeysFlagIterator{contract: _Iaggchainbaseevents.contract, event: "DisableUseDefaultVkeysFlag", logs: logs, sub: sub}, nil
}

// WatchDisableUseDefaultVkeysFlag is a free log subscription operation binding the contract event 0x922aeecd3505b46415820aae489ed9dac9e250e74d497b14c33e8360b581ac07.
//
// Solidity: event DisableUseDefaultVkeysFlag()
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) WatchDisableUseDefaultVkeysFlag(opts *bind.WatchOpts, sink chan<- *IaggchainbaseeventsDisableUseDefaultVkeysFlag) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.WatchLogs(opts, "DisableUseDefaultVkeysFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseeventsDisableUseDefaultVkeysFlag)
				if err := _Iaggchainbaseevents.contract.UnpackLog(event, "DisableUseDefaultVkeysFlag", log); err != nil {
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

// ParseDisableUseDefaultVkeysFlag is a log parse operation binding the contract event 0x922aeecd3505b46415820aae489ed9dac9e250e74d497b14c33e8360b581ac07.
//
// Solidity: event DisableUseDefaultVkeysFlag()
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) ParseDisableUseDefaultVkeysFlag(log types.Log) (*IaggchainbaseeventsDisableUseDefaultVkeysFlag, error) {
	event := new(IaggchainbaseeventsDisableUseDefaultVkeysFlag)
	if err := _Iaggchainbaseevents.contract.UnpackLog(event, "DisableUseDefaultVkeysFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseeventsEnableUseDefaultSignersFlagIterator is returned from FilterEnableUseDefaultSignersFlag and is used to iterate over the raw logs and unpacked data for EnableUseDefaultSignersFlag events raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsEnableUseDefaultSignersFlagIterator struct {
	Event *IaggchainbaseeventsEnableUseDefaultSignersFlag // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseeventsEnableUseDefaultSignersFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseeventsEnableUseDefaultSignersFlag)
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
		it.Event = new(IaggchainbaseeventsEnableUseDefaultSignersFlag)
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
func (it *IaggchainbaseeventsEnableUseDefaultSignersFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseeventsEnableUseDefaultSignersFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseeventsEnableUseDefaultSignersFlag represents a EnableUseDefaultSignersFlag event raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsEnableUseDefaultSignersFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEnableUseDefaultSignersFlag is a free log retrieval operation binding the contract event 0x67ec953bdc8546ede08f8ee91e5205a1d1814e126cb8f5d00a918ddb1eaa292b.
//
// Solidity: event EnableUseDefaultSignersFlag()
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) FilterEnableUseDefaultSignersFlag(opts *bind.FilterOpts) (*IaggchainbaseeventsEnableUseDefaultSignersFlagIterator, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.FilterLogs(opts, "EnableUseDefaultSignersFlag")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventsEnableUseDefaultSignersFlagIterator{contract: _Iaggchainbaseevents.contract, event: "EnableUseDefaultSignersFlag", logs: logs, sub: sub}, nil
}

// WatchEnableUseDefaultSignersFlag is a free log subscription operation binding the contract event 0x67ec953bdc8546ede08f8ee91e5205a1d1814e126cb8f5d00a918ddb1eaa292b.
//
// Solidity: event EnableUseDefaultSignersFlag()
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) WatchEnableUseDefaultSignersFlag(opts *bind.WatchOpts, sink chan<- *IaggchainbaseeventsEnableUseDefaultSignersFlag) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.WatchLogs(opts, "EnableUseDefaultSignersFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseeventsEnableUseDefaultSignersFlag)
				if err := _Iaggchainbaseevents.contract.UnpackLog(event, "EnableUseDefaultSignersFlag", log); err != nil {
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

// ParseEnableUseDefaultSignersFlag is a log parse operation binding the contract event 0x67ec953bdc8546ede08f8ee91e5205a1d1814e126cb8f5d00a918ddb1eaa292b.
//
// Solidity: event EnableUseDefaultSignersFlag()
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) ParseEnableUseDefaultSignersFlag(log types.Log) (*IaggchainbaseeventsEnableUseDefaultSignersFlag, error) {
	event := new(IaggchainbaseeventsEnableUseDefaultSignersFlag)
	if err := _Iaggchainbaseevents.contract.UnpackLog(event, "EnableUseDefaultSignersFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseeventsEnableUseDefaultVkeysFlagIterator is returned from FilterEnableUseDefaultVkeysFlag and is used to iterate over the raw logs and unpacked data for EnableUseDefaultVkeysFlag events raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsEnableUseDefaultVkeysFlagIterator struct {
	Event *IaggchainbaseeventsEnableUseDefaultVkeysFlag // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseeventsEnableUseDefaultVkeysFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseeventsEnableUseDefaultVkeysFlag)
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
		it.Event = new(IaggchainbaseeventsEnableUseDefaultVkeysFlag)
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
func (it *IaggchainbaseeventsEnableUseDefaultVkeysFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseeventsEnableUseDefaultVkeysFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseeventsEnableUseDefaultVkeysFlag represents a EnableUseDefaultVkeysFlag event raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsEnableUseDefaultVkeysFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEnableUseDefaultVkeysFlag is a free log retrieval operation binding the contract event 0xaacf3fb6dd8daa3bebb71f5548f782413c3f3531625c6f9057c0f3d751b83829.
//
// Solidity: event EnableUseDefaultVkeysFlag()
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) FilterEnableUseDefaultVkeysFlag(opts *bind.FilterOpts) (*IaggchainbaseeventsEnableUseDefaultVkeysFlagIterator, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.FilterLogs(opts, "EnableUseDefaultVkeysFlag")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventsEnableUseDefaultVkeysFlagIterator{contract: _Iaggchainbaseevents.contract, event: "EnableUseDefaultVkeysFlag", logs: logs, sub: sub}, nil
}

// WatchEnableUseDefaultVkeysFlag is a free log subscription operation binding the contract event 0xaacf3fb6dd8daa3bebb71f5548f782413c3f3531625c6f9057c0f3d751b83829.
//
// Solidity: event EnableUseDefaultVkeysFlag()
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) WatchEnableUseDefaultVkeysFlag(opts *bind.WatchOpts, sink chan<- *IaggchainbaseeventsEnableUseDefaultVkeysFlag) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.WatchLogs(opts, "EnableUseDefaultVkeysFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseeventsEnableUseDefaultVkeysFlag)
				if err := _Iaggchainbaseevents.contract.UnpackLog(event, "EnableUseDefaultVkeysFlag", log); err != nil {
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

// ParseEnableUseDefaultVkeysFlag is a log parse operation binding the contract event 0xaacf3fb6dd8daa3bebb71f5548f782413c3f3531625c6f9057c0f3d751b83829.
//
// Solidity: event EnableUseDefaultVkeysFlag()
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) ParseEnableUseDefaultVkeysFlag(log types.Log) (*IaggchainbaseeventsEnableUseDefaultVkeysFlag, error) {
	event := new(IaggchainbaseeventsEnableUseDefaultVkeysFlag)
	if err := _Iaggchainbaseevents.contract.UnpackLog(event, "EnableUseDefaultVkeysFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseeventsSetAggchainMetadataManagerIterator is returned from FilterSetAggchainMetadataManager and is used to iterate over the raw logs and unpacked data for SetAggchainMetadataManager events raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsSetAggchainMetadataManagerIterator struct {
	Event *IaggchainbaseeventsSetAggchainMetadataManager // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseeventsSetAggchainMetadataManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseeventsSetAggchainMetadataManager)
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
		it.Event = new(IaggchainbaseeventsSetAggchainMetadataManager)
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
func (it *IaggchainbaseeventsSetAggchainMetadataManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseeventsSetAggchainMetadataManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseeventsSetAggchainMetadataManager represents a SetAggchainMetadataManager event raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsSetAggchainMetadataManager struct {
	OldAggchainMetadataManager common.Address
	NewAggchainMetadataManager common.Address
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterSetAggchainMetadataManager is a free log retrieval operation binding the contract event 0x82ae2ec69f24a6de4517a5a45d4983651b578b3d8dc9262af5e352572fc64373.
//
// Solidity: event SetAggchainMetadataManager(address oldAggchainMetadataManager, address newAggchainMetadataManager)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) FilterSetAggchainMetadataManager(opts *bind.FilterOpts) (*IaggchainbaseeventsSetAggchainMetadataManagerIterator, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.FilterLogs(opts, "SetAggchainMetadataManager")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventsSetAggchainMetadataManagerIterator{contract: _Iaggchainbaseevents.contract, event: "SetAggchainMetadataManager", logs: logs, sub: sub}, nil
}

// WatchSetAggchainMetadataManager is a free log subscription operation binding the contract event 0x82ae2ec69f24a6de4517a5a45d4983651b578b3d8dc9262af5e352572fc64373.
//
// Solidity: event SetAggchainMetadataManager(address oldAggchainMetadataManager, address newAggchainMetadataManager)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) WatchSetAggchainMetadataManager(opts *bind.WatchOpts, sink chan<- *IaggchainbaseeventsSetAggchainMetadataManager) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.WatchLogs(opts, "SetAggchainMetadataManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseeventsSetAggchainMetadataManager)
				if err := _Iaggchainbaseevents.contract.UnpackLog(event, "SetAggchainMetadataManager", log); err != nil {
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

// ParseSetAggchainMetadataManager is a log parse operation binding the contract event 0x82ae2ec69f24a6de4517a5a45d4983651b578b3d8dc9262af5e352572fc64373.
//
// Solidity: event SetAggchainMetadataManager(address oldAggchainMetadataManager, address newAggchainMetadataManager)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) ParseSetAggchainMetadataManager(log types.Log) (*IaggchainbaseeventsSetAggchainMetadataManager, error) {
	event := new(IaggchainbaseeventsSetAggchainMetadataManager)
	if err := _Iaggchainbaseevents.contract.UnpackLog(event, "SetAggchainMetadataManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseeventsTransferAggchainManagerRoleIterator is returned from FilterTransferAggchainManagerRole and is used to iterate over the raw logs and unpacked data for TransferAggchainManagerRole events raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsTransferAggchainManagerRoleIterator struct {
	Event *IaggchainbaseeventsTransferAggchainManagerRole // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseeventsTransferAggchainManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseeventsTransferAggchainManagerRole)
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
		it.Event = new(IaggchainbaseeventsTransferAggchainManagerRole)
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
func (it *IaggchainbaseeventsTransferAggchainManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseeventsTransferAggchainManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseeventsTransferAggchainManagerRole represents a TransferAggchainManagerRole event raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsTransferAggchainManagerRole struct {
	CurrentAggchainManager    common.Address
	NewPendingAggchainManager common.Address
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterTransferAggchainManagerRole is a free log retrieval operation binding the contract event 0xa3d8e5d045432398be30f83ce7c35a7bfc220c1b66cc5bf3f4dd4d539d93fab6.
//
// Solidity: event TransferAggchainManagerRole(address currentAggchainManager, address newPendingAggchainManager)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) FilterTransferAggchainManagerRole(opts *bind.FilterOpts) (*IaggchainbaseeventsTransferAggchainManagerRoleIterator, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.FilterLogs(opts, "TransferAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventsTransferAggchainManagerRoleIterator{contract: _Iaggchainbaseevents.contract, event: "TransferAggchainManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferAggchainManagerRole is a free log subscription operation binding the contract event 0xa3d8e5d045432398be30f83ce7c35a7bfc220c1b66cc5bf3f4dd4d539d93fab6.
//
// Solidity: event TransferAggchainManagerRole(address currentAggchainManager, address newPendingAggchainManager)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) WatchTransferAggchainManagerRole(opts *bind.WatchOpts, sink chan<- *IaggchainbaseeventsTransferAggchainManagerRole) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.WatchLogs(opts, "TransferAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseeventsTransferAggchainManagerRole)
				if err := _Iaggchainbaseevents.contract.UnpackLog(event, "TransferAggchainManagerRole", log); err != nil {
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

// ParseTransferAggchainManagerRole is a log parse operation binding the contract event 0xa3d8e5d045432398be30f83ce7c35a7bfc220c1b66cc5bf3f4dd4d539d93fab6.
//
// Solidity: event TransferAggchainManagerRole(address currentAggchainManager, address newPendingAggchainManager)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) ParseTransferAggchainManagerRole(log types.Log) (*IaggchainbaseeventsTransferAggchainManagerRole, error) {
	event := new(IaggchainbaseeventsTransferAggchainManagerRole)
	if err := _Iaggchainbaseevents.contract.UnpackLog(event, "TransferAggchainManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseeventsUpdateAggchainVKeyIterator is returned from FilterUpdateAggchainVKey and is used to iterate over the raw logs and unpacked data for UpdateAggchainVKey events raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsUpdateAggchainVKeyIterator struct {
	Event *IaggchainbaseeventsUpdateAggchainVKey // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseeventsUpdateAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseeventsUpdateAggchainVKey)
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
		it.Event = new(IaggchainbaseeventsUpdateAggchainVKey)
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
func (it *IaggchainbaseeventsUpdateAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseeventsUpdateAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseeventsUpdateAggchainVKey represents a UpdateAggchainVKey event raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsUpdateAggchainVKey struct {
	Selector             [4]byte
	PreviousAggchainVKey [32]byte
	NewAggchainVKey      [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateAggchainVKey is a free log retrieval operation binding the contract event 0x0aa5f73c189fb0b0a7cc98ae5fa89dfc16595480396208483518178435ed5b4f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 previousAggchainVKey, bytes32 newAggchainVKey)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) FilterUpdateAggchainVKey(opts *bind.FilterOpts) (*IaggchainbaseeventsUpdateAggchainVKeyIterator, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.FilterLogs(opts, "UpdateAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventsUpdateAggchainVKeyIterator{contract: _Iaggchainbaseevents.contract, event: "UpdateAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchUpdateAggchainVKey is a free log subscription operation binding the contract event 0x0aa5f73c189fb0b0a7cc98ae5fa89dfc16595480396208483518178435ed5b4f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 previousAggchainVKey, bytes32 newAggchainVKey)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) WatchUpdateAggchainVKey(opts *bind.WatchOpts, sink chan<- *IaggchainbaseeventsUpdateAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.WatchLogs(opts, "UpdateAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseeventsUpdateAggchainVKey)
				if err := _Iaggchainbaseevents.contract.UnpackLog(event, "UpdateAggchainVKey", log); err != nil {
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

// ParseUpdateAggchainVKey is a log parse operation binding the contract event 0x0aa5f73c189fb0b0a7cc98ae5fa89dfc16595480396208483518178435ed5b4f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 previousAggchainVKey, bytes32 newAggchainVKey)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) ParseUpdateAggchainVKey(log types.Log) (*IaggchainbaseeventsUpdateAggchainVKey, error) {
	event := new(IaggchainbaseeventsUpdateAggchainVKey)
	if err := _Iaggchainbaseevents.contract.UnpackLog(event, "UpdateAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

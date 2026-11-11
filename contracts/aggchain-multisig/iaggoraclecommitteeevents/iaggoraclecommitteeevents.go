// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iaggoraclecommitteeevents

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

// IaggoraclecommitteeeventsMetaData contains all meta data concerning the Iaggoraclecommitteeevents contract.
var IaggoraclecommitteeeventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"}],\"name\":\"GERInjected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"}],\"name\":\"GERMarkedAsInjected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newProposer\",\"type\":\"address\"}],\"name\":\"UpdateAggOracleProposer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ValidatorsAndThresholdUpdated\",\"type\":\"event\"}]",
}

// IaggoraclecommitteeeventsABI is the input ABI used to generate the binding from.
// Deprecated: Use IaggoraclecommitteeeventsMetaData.ABI instead.
var IaggoraclecommitteeeventsABI = IaggoraclecommitteeeventsMetaData.ABI

// Iaggoraclecommitteeevents is an auto generated Go binding around an Ethereum contract.
type Iaggoraclecommitteeevents struct {
	IaggoraclecommitteeeventsCaller     // Read-only binding to the contract
	IaggoraclecommitteeeventsTransactor // Write-only binding to the contract
	IaggoraclecommitteeeventsFilterer   // Log filterer for contract events
}

// IaggoraclecommitteeeventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IaggoraclecommitteeeventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggoraclecommitteeeventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IaggoraclecommitteeeventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggoraclecommitteeeventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IaggoraclecommitteeeventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggoraclecommitteeeventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IaggoraclecommitteeeventsSession struct {
	Contract     *Iaggoraclecommitteeevents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IaggoraclecommitteeeventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IaggoraclecommitteeeventsCallerSession struct {
	Contract *IaggoraclecommitteeeventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// IaggoraclecommitteeeventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IaggoraclecommitteeeventsTransactorSession struct {
	Contract     *IaggoraclecommitteeeventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// IaggoraclecommitteeeventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IaggoraclecommitteeeventsRaw struct {
	Contract *Iaggoraclecommitteeevents // Generic contract binding to access the raw methods on
}

// IaggoraclecommitteeeventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IaggoraclecommitteeeventsCallerRaw struct {
	Contract *IaggoraclecommitteeeventsCaller // Generic read-only contract binding to access the raw methods on
}

// IaggoraclecommitteeeventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IaggoraclecommitteeeventsTransactorRaw struct {
	Contract *IaggoraclecommitteeeventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIaggoraclecommitteeevents creates a new instance of Iaggoraclecommitteeevents, bound to a specific deployed contract.
func NewIaggoraclecommitteeevents(address common.Address, backend bind.ContractBackend) (*Iaggoraclecommitteeevents, error) {
	contract, err := bindIaggoraclecommitteeevents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iaggoraclecommitteeevents{IaggoraclecommitteeeventsCaller: IaggoraclecommitteeeventsCaller{contract: contract}, IaggoraclecommitteeeventsTransactor: IaggoraclecommitteeeventsTransactor{contract: contract}, IaggoraclecommitteeeventsFilterer: IaggoraclecommitteeeventsFilterer{contract: contract}}, nil
}

// NewIaggoraclecommitteeeventsCaller creates a new read-only instance of Iaggoraclecommitteeevents, bound to a specific deployed contract.
func NewIaggoraclecommitteeeventsCaller(address common.Address, caller bind.ContractCaller) (*IaggoraclecommitteeeventsCaller, error) {
	contract, err := bindIaggoraclecommitteeevents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IaggoraclecommitteeeventsCaller{contract: contract}, nil
}

// NewIaggoraclecommitteeeventsTransactor creates a new write-only instance of Iaggoraclecommitteeevents, bound to a specific deployed contract.
func NewIaggoraclecommitteeeventsTransactor(address common.Address, transactor bind.ContractTransactor) (*IaggoraclecommitteeeventsTransactor, error) {
	contract, err := bindIaggoraclecommitteeevents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IaggoraclecommitteeeventsTransactor{contract: contract}, nil
}

// NewIaggoraclecommitteeeventsFilterer creates a new log filterer instance of Iaggoraclecommitteeevents, bound to a specific deployed contract.
func NewIaggoraclecommitteeeventsFilterer(address common.Address, filterer bind.ContractFilterer) (*IaggoraclecommitteeeventsFilterer, error) {
	contract, err := bindIaggoraclecommitteeevents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IaggoraclecommitteeeventsFilterer{contract: contract}, nil
}

// bindIaggoraclecommitteeevents binds a generic wrapper to an already deployed contract.
func bindIaggoraclecommitteeevents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IaggoraclecommitteeeventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iaggoraclecommitteeevents *IaggoraclecommitteeeventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iaggoraclecommitteeevents.Contract.IaggoraclecommitteeeventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iaggoraclecommitteeevents *IaggoraclecommitteeeventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iaggoraclecommitteeevents.Contract.IaggoraclecommitteeeventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iaggoraclecommitteeevents *IaggoraclecommitteeeventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iaggoraclecommitteeevents.Contract.IaggoraclecommitteeeventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iaggoraclecommitteeevents *IaggoraclecommitteeeventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iaggoraclecommitteeevents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iaggoraclecommitteeevents *IaggoraclecommitteeeventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iaggoraclecommitteeevents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iaggoraclecommitteeevents *IaggoraclecommitteeeventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iaggoraclecommitteeevents.Contract.contract.Transact(opts, method, params...)
}

// IaggoraclecommitteeeventsGERInjectedIterator is returned from FilterGERInjected and is used to iterate over the raw logs and unpacked data for GERInjected events raised by the Iaggoraclecommitteeevents contract.
type IaggoraclecommitteeeventsGERInjectedIterator struct {
	Event *IaggoraclecommitteeeventsGERInjected // Event containing the contract specifics and raw log

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
func (it *IaggoraclecommitteeeventsGERInjectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggoraclecommitteeeventsGERInjected)
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
		it.Event = new(IaggoraclecommitteeeventsGERInjected)
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
func (it *IaggoraclecommitteeeventsGERInjectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggoraclecommitteeeventsGERInjectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggoraclecommitteeeventsGERInjected represents a GERInjected event raised by the Iaggoraclecommitteeevents contract.
type IaggoraclecommitteeeventsGERInjected struct {
	GlobalExitRoot [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGERInjected is a free log retrieval operation binding the contract event 0x45b91a7a09575a05b9dd504dff707cdeb7eda5f9d514491caf5951551b020e1e.
//
// Solidity: event GERInjected(bytes32 indexed globalExitRoot)
func (_Iaggoraclecommitteeevents *IaggoraclecommitteeeventsFilterer) FilterGERInjected(opts *bind.FilterOpts, globalExitRoot [][32]byte) (*IaggoraclecommitteeeventsGERInjectedIterator, error) {

	var globalExitRootRule []interface{}
	for _, globalExitRootItem := range globalExitRoot {
		globalExitRootRule = append(globalExitRootRule, globalExitRootItem)
	}

	logs, sub, err := _Iaggoraclecommitteeevents.contract.FilterLogs(opts, "GERInjected", globalExitRootRule)
	if err != nil {
		return nil, err
	}
	return &IaggoraclecommitteeeventsGERInjectedIterator{contract: _Iaggoraclecommitteeevents.contract, event: "GERInjected", logs: logs, sub: sub}, nil
}

// WatchGERInjected is a free log subscription operation binding the contract event 0x45b91a7a09575a05b9dd504dff707cdeb7eda5f9d514491caf5951551b020e1e.
//
// Solidity: event GERInjected(bytes32 indexed globalExitRoot)
func (_Iaggoraclecommitteeevents *IaggoraclecommitteeeventsFilterer) WatchGERInjected(opts *bind.WatchOpts, sink chan<- *IaggoraclecommitteeeventsGERInjected, globalExitRoot [][32]byte) (event.Subscription, error) {

	var globalExitRootRule []interface{}
	for _, globalExitRootItem := range globalExitRoot {
		globalExitRootRule = append(globalExitRootRule, globalExitRootItem)
	}

	logs, sub, err := _Iaggoraclecommitteeevents.contract.WatchLogs(opts, "GERInjected", globalExitRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggoraclecommitteeeventsGERInjected)
				if err := _Iaggoraclecommitteeevents.contract.UnpackLog(event, "GERInjected", log); err != nil {
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

// ParseGERInjected is a log parse operation binding the contract event 0x45b91a7a09575a05b9dd504dff707cdeb7eda5f9d514491caf5951551b020e1e.
//
// Solidity: event GERInjected(bytes32 indexed globalExitRoot)
func (_Iaggoraclecommitteeevents *IaggoraclecommitteeeventsFilterer) ParseGERInjected(log types.Log) (*IaggoraclecommitteeeventsGERInjected, error) {
	event := new(IaggoraclecommitteeeventsGERInjected)
	if err := _Iaggoraclecommitteeevents.contract.UnpackLog(event, "GERInjected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggoraclecommitteeeventsGERMarkedAsInjectedIterator is returned from FilterGERMarkedAsInjected and is used to iterate over the raw logs and unpacked data for GERMarkedAsInjected events raised by the Iaggoraclecommitteeevents contract.
type IaggoraclecommitteeeventsGERMarkedAsInjectedIterator struct {
	Event *IaggoraclecommitteeeventsGERMarkedAsInjected // Event containing the contract specifics and raw log

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
func (it *IaggoraclecommitteeeventsGERMarkedAsInjectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggoraclecommitteeeventsGERMarkedAsInjected)
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
		it.Event = new(IaggoraclecommitteeeventsGERMarkedAsInjected)
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
func (it *IaggoraclecommitteeeventsGERMarkedAsInjectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggoraclecommitteeeventsGERMarkedAsInjectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggoraclecommitteeeventsGERMarkedAsInjected represents a GERMarkedAsInjected event raised by the Iaggoraclecommitteeevents contract.
type IaggoraclecommitteeeventsGERMarkedAsInjected struct {
	GlobalExitRoot [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGERMarkedAsInjected is a free log retrieval operation binding the contract event 0xcafdbc8c7e71f2990afc9e787ff729a197666534c9a49959c1373c819b708a8f.
//
// Solidity: event GERMarkedAsInjected(bytes32 indexed globalExitRoot)
func (_Iaggoraclecommitteeevents *IaggoraclecommitteeeventsFilterer) FilterGERMarkedAsInjected(opts *bind.FilterOpts, globalExitRoot [][32]byte) (*IaggoraclecommitteeeventsGERMarkedAsInjectedIterator, error) {

	var globalExitRootRule []interface{}
	for _, globalExitRootItem := range globalExitRoot {
		globalExitRootRule = append(globalExitRootRule, globalExitRootItem)
	}

	logs, sub, err := _Iaggoraclecommitteeevents.contract.FilterLogs(opts, "GERMarkedAsInjected", globalExitRootRule)
	if err != nil {
		return nil, err
	}
	return &IaggoraclecommitteeeventsGERMarkedAsInjectedIterator{contract: _Iaggoraclecommitteeevents.contract, event: "GERMarkedAsInjected", logs: logs, sub: sub}, nil
}

// WatchGERMarkedAsInjected is a free log subscription operation binding the contract event 0xcafdbc8c7e71f2990afc9e787ff729a197666534c9a49959c1373c819b708a8f.
//
// Solidity: event GERMarkedAsInjected(bytes32 indexed globalExitRoot)
func (_Iaggoraclecommitteeevents *IaggoraclecommitteeeventsFilterer) WatchGERMarkedAsInjected(opts *bind.WatchOpts, sink chan<- *IaggoraclecommitteeeventsGERMarkedAsInjected, globalExitRoot [][32]byte) (event.Subscription, error) {

	var globalExitRootRule []interface{}
	for _, globalExitRootItem := range globalExitRoot {
		globalExitRootRule = append(globalExitRootRule, globalExitRootItem)
	}

	logs, sub, err := _Iaggoraclecommitteeevents.contract.WatchLogs(opts, "GERMarkedAsInjected", globalExitRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggoraclecommitteeeventsGERMarkedAsInjected)
				if err := _Iaggoraclecommitteeevents.contract.UnpackLog(event, "GERMarkedAsInjected", log); err != nil {
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

// ParseGERMarkedAsInjected is a log parse operation binding the contract event 0xcafdbc8c7e71f2990afc9e787ff729a197666534c9a49959c1373c819b708a8f.
//
// Solidity: event GERMarkedAsInjected(bytes32 indexed globalExitRoot)
func (_Iaggoraclecommitteeevents *IaggoraclecommitteeeventsFilterer) ParseGERMarkedAsInjected(log types.Log) (*IaggoraclecommitteeeventsGERMarkedAsInjected, error) {
	event := new(IaggoraclecommitteeeventsGERMarkedAsInjected)
	if err := _Iaggoraclecommitteeevents.contract.UnpackLog(event, "GERMarkedAsInjected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggoraclecommitteeeventsUpdateAggOracleProposerIterator is returned from FilterUpdateAggOracleProposer and is used to iterate over the raw logs and unpacked data for UpdateAggOracleProposer events raised by the Iaggoraclecommitteeevents contract.
type IaggoraclecommitteeeventsUpdateAggOracleProposerIterator struct {
	Event *IaggoraclecommitteeeventsUpdateAggOracleProposer // Event containing the contract specifics and raw log

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
func (it *IaggoraclecommitteeeventsUpdateAggOracleProposerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggoraclecommitteeeventsUpdateAggOracleProposer)
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
		it.Event = new(IaggoraclecommitteeeventsUpdateAggOracleProposer)
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
func (it *IaggoraclecommitteeeventsUpdateAggOracleProposerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggoraclecommitteeeventsUpdateAggOracleProposerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggoraclecommitteeeventsUpdateAggOracleProposer represents a UpdateAggOracleProposer event raised by the Iaggoraclecommitteeevents contract.
type IaggoraclecommitteeeventsUpdateAggOracleProposer struct {
	NewProposer common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateAggOracleProposer is a free log retrieval operation binding the contract event 0x93c536fad4968838f70b6a5841cc12ffbd1a657048ef449f715f0d9ebdbeffca.
//
// Solidity: event UpdateAggOracleProposer(address indexed newProposer)
func (_Iaggoraclecommitteeevents *IaggoraclecommitteeeventsFilterer) FilterUpdateAggOracleProposer(opts *bind.FilterOpts, newProposer []common.Address) (*IaggoraclecommitteeeventsUpdateAggOracleProposerIterator, error) {

	var newProposerRule []interface{}
	for _, newProposerItem := range newProposer {
		newProposerRule = append(newProposerRule, newProposerItem)
	}

	logs, sub, err := _Iaggoraclecommitteeevents.contract.FilterLogs(opts, "UpdateAggOracleProposer", newProposerRule)
	if err != nil {
		return nil, err
	}
	return &IaggoraclecommitteeeventsUpdateAggOracleProposerIterator{contract: _Iaggoraclecommitteeevents.contract, event: "UpdateAggOracleProposer", logs: logs, sub: sub}, nil
}

// WatchUpdateAggOracleProposer is a free log subscription operation binding the contract event 0x93c536fad4968838f70b6a5841cc12ffbd1a657048ef449f715f0d9ebdbeffca.
//
// Solidity: event UpdateAggOracleProposer(address indexed newProposer)
func (_Iaggoraclecommitteeevents *IaggoraclecommitteeeventsFilterer) WatchUpdateAggOracleProposer(opts *bind.WatchOpts, sink chan<- *IaggoraclecommitteeeventsUpdateAggOracleProposer, newProposer []common.Address) (event.Subscription, error) {

	var newProposerRule []interface{}
	for _, newProposerItem := range newProposer {
		newProposerRule = append(newProposerRule, newProposerItem)
	}

	logs, sub, err := _Iaggoraclecommitteeevents.contract.WatchLogs(opts, "UpdateAggOracleProposer", newProposerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggoraclecommitteeeventsUpdateAggOracleProposer)
				if err := _Iaggoraclecommitteeevents.contract.UnpackLog(event, "UpdateAggOracleProposer", log); err != nil {
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

// ParseUpdateAggOracleProposer is a log parse operation binding the contract event 0x93c536fad4968838f70b6a5841cc12ffbd1a657048ef449f715f0d9ebdbeffca.
//
// Solidity: event UpdateAggOracleProposer(address indexed newProposer)
func (_Iaggoraclecommitteeevents *IaggoraclecommitteeeventsFilterer) ParseUpdateAggOracleProposer(log types.Log) (*IaggoraclecommitteeeventsUpdateAggOracleProposer, error) {
	event := new(IaggoraclecommitteeeventsUpdateAggOracleProposer)
	if err := _Iaggoraclecommitteeevents.contract.UnpackLog(event, "UpdateAggOracleProposer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggoraclecommitteeeventsValidatorsAndThresholdUpdatedIterator is returned from FilterValidatorsAndThresholdUpdated and is used to iterate over the raw logs and unpacked data for ValidatorsAndThresholdUpdated events raised by the Iaggoraclecommitteeevents contract.
type IaggoraclecommitteeeventsValidatorsAndThresholdUpdatedIterator struct {
	Event *IaggoraclecommitteeeventsValidatorsAndThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *IaggoraclecommitteeeventsValidatorsAndThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggoraclecommitteeeventsValidatorsAndThresholdUpdated)
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
		it.Event = new(IaggoraclecommitteeeventsValidatorsAndThresholdUpdated)
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
func (it *IaggoraclecommitteeeventsValidatorsAndThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggoraclecommitteeeventsValidatorsAndThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggoraclecommitteeeventsValidatorsAndThresholdUpdated represents a ValidatorsAndThresholdUpdated event raised by the Iaggoraclecommitteeevents contract.
type IaggoraclecommitteeeventsValidatorsAndThresholdUpdated struct {
	Validators []common.Address
	Threshold  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorsAndThresholdUpdated is a free log retrieval operation binding the contract event 0x50051c796f4289791fedb6e1f365497b6527df6f6f19527d566e3c9791af5460.
//
// Solidity: event ValidatorsAndThresholdUpdated(address[] validators, uint256 threshold)
func (_Iaggoraclecommitteeevents *IaggoraclecommitteeeventsFilterer) FilterValidatorsAndThresholdUpdated(opts *bind.FilterOpts) (*IaggoraclecommitteeeventsValidatorsAndThresholdUpdatedIterator, error) {

	logs, sub, err := _Iaggoraclecommitteeevents.contract.FilterLogs(opts, "ValidatorsAndThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &IaggoraclecommitteeeventsValidatorsAndThresholdUpdatedIterator{contract: _Iaggoraclecommitteeevents.contract, event: "ValidatorsAndThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorsAndThresholdUpdated is a free log subscription operation binding the contract event 0x50051c796f4289791fedb6e1f365497b6527df6f6f19527d566e3c9791af5460.
//
// Solidity: event ValidatorsAndThresholdUpdated(address[] validators, uint256 threshold)
func (_Iaggoraclecommitteeevents *IaggoraclecommitteeeventsFilterer) WatchValidatorsAndThresholdUpdated(opts *bind.WatchOpts, sink chan<- *IaggoraclecommitteeeventsValidatorsAndThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _Iaggoraclecommitteeevents.contract.WatchLogs(opts, "ValidatorsAndThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggoraclecommitteeeventsValidatorsAndThresholdUpdated)
				if err := _Iaggoraclecommitteeevents.contract.UnpackLog(event, "ValidatorsAndThresholdUpdated", log); err != nil {
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

// ParseValidatorsAndThresholdUpdated is a log parse operation binding the contract event 0x50051c796f4289791fedb6e1f365497b6527df6f6f19527d566e3c9791af5460.
//
// Solidity: event ValidatorsAndThresholdUpdated(address[] validators, uint256 threshold)
func (_Iaggoraclecommitteeevents *IaggoraclecommitteeeventsFilterer) ParseValidatorsAndThresholdUpdated(log types.Log) (*IaggoraclecommitteeeventsValidatorsAndThresholdUpdated, error) {
	event := new(IaggoraclecommitteeeventsValidatorsAndThresholdUpdated)
	if err := _Iaggoraclecommitteeevents.contract.UnpackLog(event, "ValidatorsAndThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

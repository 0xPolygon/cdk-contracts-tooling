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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAggchainManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"AcceptAggchainManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldVKeyManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVKeyManager\",\"type\":\"address\"}],\"name\":\"AcceptVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"AddAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainSignersHash\",\"type\":\"bytes32\"}],\"name\":\"AggchainSignersHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DisableUseDefaultGatewayFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EnableUseDefaultGatewayFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"aggchainSigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newThreshold\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainSignersHash\",\"type\":\"bytes32\"}],\"name\":\"SignersAndThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentAggchainManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAggchainManager\",\"type\":\"address\"}],\"name\":\"TransferAggchainManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentVKeyManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingVKeyManager\",\"type\":\"address\"}],\"name\":\"TransferVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousAggchainVKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"UpdateAggchainVKey\",\"type\":\"event\"}]",
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

// IaggchainbaseeventsAcceptVKeyManagerRoleIterator is returned from FilterAcceptVKeyManagerRole and is used to iterate over the raw logs and unpacked data for AcceptVKeyManagerRole events raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsAcceptVKeyManagerRoleIterator struct {
	Event *IaggchainbaseeventsAcceptVKeyManagerRole // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseeventsAcceptVKeyManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseeventsAcceptVKeyManagerRole)
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
		it.Event = new(IaggchainbaseeventsAcceptVKeyManagerRole)
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
func (it *IaggchainbaseeventsAcceptVKeyManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseeventsAcceptVKeyManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseeventsAcceptVKeyManagerRole represents a AcceptVKeyManagerRole event raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsAcceptVKeyManagerRole struct {
	OldVKeyManager common.Address
	NewVKeyManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAcceptVKeyManagerRole is a free log retrieval operation binding the contract event 0xbb3b066505f14a628f4ba4187a046abd4dd17e96591d7a9ed31c91c79322ffe2.
//
// Solidity: event AcceptVKeyManagerRole(address oldVKeyManager, address newVKeyManager)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) FilterAcceptVKeyManagerRole(opts *bind.FilterOpts) (*IaggchainbaseeventsAcceptVKeyManagerRoleIterator, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.FilterLogs(opts, "AcceptVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventsAcceptVKeyManagerRoleIterator{contract: _Iaggchainbaseevents.contract, event: "AcceptVKeyManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptVKeyManagerRole is a free log subscription operation binding the contract event 0xbb3b066505f14a628f4ba4187a046abd4dd17e96591d7a9ed31c91c79322ffe2.
//
// Solidity: event AcceptVKeyManagerRole(address oldVKeyManager, address newVKeyManager)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) WatchAcceptVKeyManagerRole(opts *bind.WatchOpts, sink chan<- *IaggchainbaseeventsAcceptVKeyManagerRole) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.WatchLogs(opts, "AcceptVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseeventsAcceptVKeyManagerRole)
				if err := _Iaggchainbaseevents.contract.UnpackLog(event, "AcceptVKeyManagerRole", log); err != nil {
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

// ParseAcceptVKeyManagerRole is a log parse operation binding the contract event 0xbb3b066505f14a628f4ba4187a046abd4dd17e96591d7a9ed31c91c79322ffe2.
//
// Solidity: event AcceptVKeyManagerRole(address oldVKeyManager, address newVKeyManager)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) ParseAcceptVKeyManagerRole(log types.Log) (*IaggchainbaseeventsAcceptVKeyManagerRole, error) {
	event := new(IaggchainbaseeventsAcceptVKeyManagerRole)
	if err := _Iaggchainbaseevents.contract.UnpackLog(event, "AcceptVKeyManagerRole", log); err != nil {
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

// IaggchainbaseeventsAggchainSignersHashUpdatedIterator is returned from FilterAggchainSignersHashUpdated and is used to iterate over the raw logs and unpacked data for AggchainSignersHashUpdated events raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsAggchainSignersHashUpdatedIterator struct {
	Event *IaggchainbaseeventsAggchainSignersHashUpdated // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseeventsAggchainSignersHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseeventsAggchainSignersHashUpdated)
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
		it.Event = new(IaggchainbaseeventsAggchainSignersHashUpdated)
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
func (it *IaggchainbaseeventsAggchainSignersHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseeventsAggchainSignersHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseeventsAggchainSignersHashUpdated represents a AggchainSignersHashUpdated event raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsAggchainSignersHashUpdated struct {
	NewAggchainSignersHash [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterAggchainSignersHashUpdated is a free log retrieval operation binding the contract event 0x43a1dd43d2705e069faf1da3afc2772475650d6642c2d71bd620e0fd2f7c3ee8.
//
// Solidity: event AggchainSignersHashUpdated(bytes32 newAggchainSignersHash)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) FilterAggchainSignersHashUpdated(opts *bind.FilterOpts) (*IaggchainbaseeventsAggchainSignersHashUpdatedIterator, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.FilterLogs(opts, "AggchainSignersHashUpdated")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventsAggchainSignersHashUpdatedIterator{contract: _Iaggchainbaseevents.contract, event: "AggchainSignersHashUpdated", logs: logs, sub: sub}, nil
}

// WatchAggchainSignersHashUpdated is a free log subscription operation binding the contract event 0x43a1dd43d2705e069faf1da3afc2772475650d6642c2d71bd620e0fd2f7c3ee8.
//
// Solidity: event AggchainSignersHashUpdated(bytes32 newAggchainSignersHash)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) WatchAggchainSignersHashUpdated(opts *bind.WatchOpts, sink chan<- *IaggchainbaseeventsAggchainSignersHashUpdated) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.WatchLogs(opts, "AggchainSignersHashUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseeventsAggchainSignersHashUpdated)
				if err := _Iaggchainbaseevents.contract.UnpackLog(event, "AggchainSignersHashUpdated", log); err != nil {
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

// ParseAggchainSignersHashUpdated is a log parse operation binding the contract event 0x43a1dd43d2705e069faf1da3afc2772475650d6642c2d71bd620e0fd2f7c3ee8.
//
// Solidity: event AggchainSignersHashUpdated(bytes32 newAggchainSignersHash)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) ParseAggchainSignersHashUpdated(log types.Log) (*IaggchainbaseeventsAggchainSignersHashUpdated, error) {
	event := new(IaggchainbaseeventsAggchainSignersHashUpdated)
	if err := _Iaggchainbaseevents.contract.UnpackLog(event, "AggchainSignersHashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseeventsDisableUseDefaultGatewayFlagIterator is returned from FilterDisableUseDefaultGatewayFlag and is used to iterate over the raw logs and unpacked data for DisableUseDefaultGatewayFlag events raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsDisableUseDefaultGatewayFlagIterator struct {
	Event *IaggchainbaseeventsDisableUseDefaultGatewayFlag // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseeventsDisableUseDefaultGatewayFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseeventsDisableUseDefaultGatewayFlag)
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
		it.Event = new(IaggchainbaseeventsDisableUseDefaultGatewayFlag)
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
func (it *IaggchainbaseeventsDisableUseDefaultGatewayFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseeventsDisableUseDefaultGatewayFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseeventsDisableUseDefaultGatewayFlag represents a DisableUseDefaultGatewayFlag event raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsDisableUseDefaultGatewayFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisableUseDefaultGatewayFlag is a free log retrieval operation binding the contract event 0x67dd1717a1952df380cb73eecb312e949df6d6a086bd7f88669005341972528e.
//
// Solidity: event DisableUseDefaultGatewayFlag()
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) FilterDisableUseDefaultGatewayFlag(opts *bind.FilterOpts) (*IaggchainbaseeventsDisableUseDefaultGatewayFlagIterator, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.FilterLogs(opts, "DisableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventsDisableUseDefaultGatewayFlagIterator{contract: _Iaggchainbaseevents.contract, event: "DisableUseDefaultGatewayFlag", logs: logs, sub: sub}, nil
}

// WatchDisableUseDefaultGatewayFlag is a free log subscription operation binding the contract event 0x67dd1717a1952df380cb73eecb312e949df6d6a086bd7f88669005341972528e.
//
// Solidity: event DisableUseDefaultGatewayFlag()
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) WatchDisableUseDefaultGatewayFlag(opts *bind.WatchOpts, sink chan<- *IaggchainbaseeventsDisableUseDefaultGatewayFlag) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.WatchLogs(opts, "DisableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseeventsDisableUseDefaultGatewayFlag)
				if err := _Iaggchainbaseevents.contract.UnpackLog(event, "DisableUseDefaultGatewayFlag", log); err != nil {
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

// ParseDisableUseDefaultGatewayFlag is a log parse operation binding the contract event 0x67dd1717a1952df380cb73eecb312e949df6d6a086bd7f88669005341972528e.
//
// Solidity: event DisableUseDefaultGatewayFlag()
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) ParseDisableUseDefaultGatewayFlag(log types.Log) (*IaggchainbaseeventsDisableUseDefaultGatewayFlag, error) {
	event := new(IaggchainbaseeventsDisableUseDefaultGatewayFlag)
	if err := _Iaggchainbaseevents.contract.UnpackLog(event, "DisableUseDefaultGatewayFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseeventsEnableUseDefaultGatewayFlagIterator is returned from FilterEnableUseDefaultGatewayFlag and is used to iterate over the raw logs and unpacked data for EnableUseDefaultGatewayFlag events raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsEnableUseDefaultGatewayFlagIterator struct {
	Event *IaggchainbaseeventsEnableUseDefaultGatewayFlag // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseeventsEnableUseDefaultGatewayFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseeventsEnableUseDefaultGatewayFlag)
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
		it.Event = new(IaggchainbaseeventsEnableUseDefaultGatewayFlag)
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
func (it *IaggchainbaseeventsEnableUseDefaultGatewayFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseeventsEnableUseDefaultGatewayFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseeventsEnableUseDefaultGatewayFlag represents a EnableUseDefaultGatewayFlag event raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsEnableUseDefaultGatewayFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEnableUseDefaultGatewayFlag is a free log retrieval operation binding the contract event 0xb6563aed80fde357e737eb0d19f246a58cb6bfd469933d05701ecbad0f2dca84.
//
// Solidity: event EnableUseDefaultGatewayFlag()
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) FilterEnableUseDefaultGatewayFlag(opts *bind.FilterOpts) (*IaggchainbaseeventsEnableUseDefaultGatewayFlagIterator, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.FilterLogs(opts, "EnableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventsEnableUseDefaultGatewayFlagIterator{contract: _Iaggchainbaseevents.contract, event: "EnableUseDefaultGatewayFlag", logs: logs, sub: sub}, nil
}

// WatchEnableUseDefaultGatewayFlag is a free log subscription operation binding the contract event 0xb6563aed80fde357e737eb0d19f246a58cb6bfd469933d05701ecbad0f2dca84.
//
// Solidity: event EnableUseDefaultGatewayFlag()
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) WatchEnableUseDefaultGatewayFlag(opts *bind.WatchOpts, sink chan<- *IaggchainbaseeventsEnableUseDefaultGatewayFlag) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.WatchLogs(opts, "EnableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseeventsEnableUseDefaultGatewayFlag)
				if err := _Iaggchainbaseevents.contract.UnpackLog(event, "EnableUseDefaultGatewayFlag", log); err != nil {
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

// ParseEnableUseDefaultGatewayFlag is a log parse operation binding the contract event 0xb6563aed80fde357e737eb0d19f246a58cb6bfd469933d05701ecbad0f2dca84.
//
// Solidity: event EnableUseDefaultGatewayFlag()
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) ParseEnableUseDefaultGatewayFlag(log types.Log) (*IaggchainbaseeventsEnableUseDefaultGatewayFlag, error) {
	event := new(IaggchainbaseeventsEnableUseDefaultGatewayFlag)
	if err := _Iaggchainbaseevents.contract.UnpackLog(event, "EnableUseDefaultGatewayFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseeventsSignersAndThresholdUpdatedIterator is returned from FilterSignersAndThresholdUpdated and is used to iterate over the raw logs and unpacked data for SignersAndThresholdUpdated events raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsSignersAndThresholdUpdatedIterator struct {
	Event *IaggchainbaseeventsSignersAndThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseeventsSignersAndThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseeventsSignersAndThresholdUpdated)
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
		it.Event = new(IaggchainbaseeventsSignersAndThresholdUpdated)
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
func (it *IaggchainbaseeventsSignersAndThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseeventsSignersAndThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseeventsSignersAndThresholdUpdated represents a SignersAndThresholdUpdated event raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsSignersAndThresholdUpdated struct {
	AggchainSigners        []common.Address
	NewThreshold           uint32
	NewAggchainSignersHash [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSignersAndThresholdUpdated is a free log retrieval operation binding the contract event 0xa9031663ce9b6bcb10b14a169f269c3a2ab3685eba97d673e3af691473bf59f9.
//
// Solidity: event SignersAndThresholdUpdated(address[] aggchainSigners, uint32 newThreshold, bytes32 newAggchainSignersHash)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) FilterSignersAndThresholdUpdated(opts *bind.FilterOpts) (*IaggchainbaseeventsSignersAndThresholdUpdatedIterator, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.FilterLogs(opts, "SignersAndThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventsSignersAndThresholdUpdatedIterator{contract: _Iaggchainbaseevents.contract, event: "SignersAndThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchSignersAndThresholdUpdated is a free log subscription operation binding the contract event 0xa9031663ce9b6bcb10b14a169f269c3a2ab3685eba97d673e3af691473bf59f9.
//
// Solidity: event SignersAndThresholdUpdated(address[] aggchainSigners, uint32 newThreshold, bytes32 newAggchainSignersHash)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) WatchSignersAndThresholdUpdated(opts *bind.WatchOpts, sink chan<- *IaggchainbaseeventsSignersAndThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.WatchLogs(opts, "SignersAndThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseeventsSignersAndThresholdUpdated)
				if err := _Iaggchainbaseevents.contract.UnpackLog(event, "SignersAndThresholdUpdated", log); err != nil {
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

// ParseSignersAndThresholdUpdated is a log parse operation binding the contract event 0xa9031663ce9b6bcb10b14a169f269c3a2ab3685eba97d673e3af691473bf59f9.
//
// Solidity: event SignersAndThresholdUpdated(address[] aggchainSigners, uint32 newThreshold, bytes32 newAggchainSignersHash)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) ParseSignersAndThresholdUpdated(log types.Log) (*IaggchainbaseeventsSignersAndThresholdUpdated, error) {
	event := new(IaggchainbaseeventsSignersAndThresholdUpdated)
	if err := _Iaggchainbaseevents.contract.UnpackLog(event, "SignersAndThresholdUpdated", log); err != nil {
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

// IaggchainbaseeventsTransferVKeyManagerRoleIterator is returned from FilterTransferVKeyManagerRole and is used to iterate over the raw logs and unpacked data for TransferVKeyManagerRole events raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsTransferVKeyManagerRoleIterator struct {
	Event *IaggchainbaseeventsTransferVKeyManagerRole // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseeventsTransferVKeyManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseeventsTransferVKeyManagerRole)
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
		it.Event = new(IaggchainbaseeventsTransferVKeyManagerRole)
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
func (it *IaggchainbaseeventsTransferVKeyManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseeventsTransferVKeyManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseeventsTransferVKeyManagerRole represents a TransferVKeyManagerRole event raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsTransferVKeyManagerRole struct {
	CurrentVKeyManager    common.Address
	NewPendingVKeyManager common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterTransferVKeyManagerRole is a free log retrieval operation binding the contract event 0xc54ae01017d0b80bd8af833f66387d6eb547dc16c8206faf13d0b72764aab8b2.
//
// Solidity: event TransferVKeyManagerRole(address currentVKeyManager, address newPendingVKeyManager)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) FilterTransferVKeyManagerRole(opts *bind.FilterOpts) (*IaggchainbaseeventsTransferVKeyManagerRoleIterator, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.FilterLogs(opts, "TransferVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventsTransferVKeyManagerRoleIterator{contract: _Iaggchainbaseevents.contract, event: "TransferVKeyManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferVKeyManagerRole is a free log subscription operation binding the contract event 0xc54ae01017d0b80bd8af833f66387d6eb547dc16c8206faf13d0b72764aab8b2.
//
// Solidity: event TransferVKeyManagerRole(address currentVKeyManager, address newPendingVKeyManager)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) WatchTransferVKeyManagerRole(opts *bind.WatchOpts, sink chan<- *IaggchainbaseeventsTransferVKeyManagerRole) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.WatchLogs(opts, "TransferVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseeventsTransferVKeyManagerRole)
				if err := _Iaggchainbaseevents.contract.UnpackLog(event, "TransferVKeyManagerRole", log); err != nil {
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

// ParseTransferVKeyManagerRole is a log parse operation binding the contract event 0xc54ae01017d0b80bd8af833f66387d6eb547dc16c8206faf13d0b72764aab8b2.
//
// Solidity: event TransferVKeyManagerRole(address currentVKeyManager, address newPendingVKeyManager)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) ParseTransferVKeyManagerRole(log types.Log) (*IaggchainbaseeventsTransferVKeyManagerRole, error) {
	event := new(IaggchainbaseeventsTransferVKeyManagerRole)
	if err := _Iaggchainbaseevents.contract.UnpackLog(event, "TransferVKeyManagerRole", log); err != nil {
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

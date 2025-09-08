// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iaggchainbaseeventsprevious

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

// IaggchainbaseeventspreviousMetaData contains all meta data concerning the Iaggchainbaseeventsprevious contract.
var IaggchainbaseeventspreviousMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAggchainManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"AcceptAggchainManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldVKeyManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVKeyManager\",\"type\":\"address\"}],\"name\":\"AcceptVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"AddAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DisableUseDefaultGatewayFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EnableUseDefaultGatewayFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentAggchainManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAggchainManager\",\"type\":\"address\"}],\"name\":\"TransferAggchainManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentVKeyManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingVKeyManager\",\"type\":\"address\"}],\"name\":\"TransferVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousAggchainVKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"UpdateAggchainVKey\",\"type\":\"event\"}]",
}

// IaggchainbaseeventspreviousABI is the input ABI used to generate the binding from.
// Deprecated: Use IaggchainbaseeventspreviousMetaData.ABI instead.
var IaggchainbaseeventspreviousABI = IaggchainbaseeventspreviousMetaData.ABI

// Iaggchainbaseeventsprevious is an auto generated Go binding around an Ethereum contract.
type Iaggchainbaseeventsprevious struct {
	IaggchainbaseeventspreviousCaller     // Read-only binding to the contract
	IaggchainbaseeventspreviousTransactor // Write-only binding to the contract
	IaggchainbaseeventspreviousFilterer   // Log filterer for contract events
}

// IaggchainbaseeventspreviousCaller is an auto generated read-only Go binding around an Ethereum contract.
type IaggchainbaseeventspreviousCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggchainbaseeventspreviousTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IaggchainbaseeventspreviousTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggchainbaseeventspreviousFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IaggchainbaseeventspreviousFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggchainbaseeventspreviousSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IaggchainbaseeventspreviousSession struct {
	Contract     *Iaggchainbaseeventsprevious // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IaggchainbaseeventspreviousCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IaggchainbaseeventspreviousCallerSession struct {
	Contract *IaggchainbaseeventspreviousCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// IaggchainbaseeventspreviousTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IaggchainbaseeventspreviousTransactorSession struct {
	Contract     *IaggchainbaseeventspreviousTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// IaggchainbaseeventspreviousRaw is an auto generated low-level Go binding around an Ethereum contract.
type IaggchainbaseeventspreviousRaw struct {
	Contract *Iaggchainbaseeventsprevious // Generic contract binding to access the raw methods on
}

// IaggchainbaseeventspreviousCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IaggchainbaseeventspreviousCallerRaw struct {
	Contract *IaggchainbaseeventspreviousCaller // Generic read-only contract binding to access the raw methods on
}

// IaggchainbaseeventspreviousTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IaggchainbaseeventspreviousTransactorRaw struct {
	Contract *IaggchainbaseeventspreviousTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIaggchainbaseeventsprevious creates a new instance of Iaggchainbaseeventsprevious, bound to a specific deployed contract.
func NewIaggchainbaseeventsprevious(address common.Address, backend bind.ContractBackend) (*Iaggchainbaseeventsprevious, error) {
	contract, err := bindIaggchainbaseeventsprevious(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iaggchainbaseeventsprevious{IaggchainbaseeventspreviousCaller: IaggchainbaseeventspreviousCaller{contract: contract}, IaggchainbaseeventspreviousTransactor: IaggchainbaseeventspreviousTransactor{contract: contract}, IaggchainbaseeventspreviousFilterer: IaggchainbaseeventspreviousFilterer{contract: contract}}, nil
}

// NewIaggchainbaseeventspreviousCaller creates a new read-only instance of Iaggchainbaseeventsprevious, bound to a specific deployed contract.
func NewIaggchainbaseeventspreviousCaller(address common.Address, caller bind.ContractCaller) (*IaggchainbaseeventspreviousCaller, error) {
	contract, err := bindIaggchainbaseeventsprevious(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventspreviousCaller{contract: contract}, nil
}

// NewIaggchainbaseeventspreviousTransactor creates a new write-only instance of Iaggchainbaseeventsprevious, bound to a specific deployed contract.
func NewIaggchainbaseeventspreviousTransactor(address common.Address, transactor bind.ContractTransactor) (*IaggchainbaseeventspreviousTransactor, error) {
	contract, err := bindIaggchainbaseeventsprevious(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventspreviousTransactor{contract: contract}, nil
}

// NewIaggchainbaseeventspreviousFilterer creates a new log filterer instance of Iaggchainbaseeventsprevious, bound to a specific deployed contract.
func NewIaggchainbaseeventspreviousFilterer(address common.Address, filterer bind.ContractFilterer) (*IaggchainbaseeventspreviousFilterer, error) {
	contract, err := bindIaggchainbaseeventsprevious(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventspreviousFilterer{contract: contract}, nil
}

// bindIaggchainbaseeventsprevious binds a generic wrapper to an already deployed contract.
func bindIaggchainbaseeventsprevious(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IaggchainbaseeventspreviousMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iaggchainbaseeventsprevious.Contract.IaggchainbaseeventspreviousCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iaggchainbaseeventsprevious.Contract.IaggchainbaseeventspreviousTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iaggchainbaseeventsprevious.Contract.IaggchainbaseeventspreviousTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iaggchainbaseeventsprevious.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iaggchainbaseeventsprevious.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iaggchainbaseeventsprevious.Contract.contract.Transact(opts, method, params...)
}

// IaggchainbaseeventspreviousAcceptAggchainManagerRoleIterator is returned from FilterAcceptAggchainManagerRole and is used to iterate over the raw logs and unpacked data for AcceptAggchainManagerRole events raised by the Iaggchainbaseeventsprevious contract.
type IaggchainbaseeventspreviousAcceptAggchainManagerRoleIterator struct {
	Event *IaggchainbaseeventspreviousAcceptAggchainManagerRole // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseeventspreviousAcceptAggchainManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseeventspreviousAcceptAggchainManagerRole)
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
		it.Event = new(IaggchainbaseeventspreviousAcceptAggchainManagerRole)
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
func (it *IaggchainbaseeventspreviousAcceptAggchainManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseeventspreviousAcceptAggchainManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseeventspreviousAcceptAggchainManagerRole represents a AcceptAggchainManagerRole event raised by the Iaggchainbaseeventsprevious contract.
type IaggchainbaseeventspreviousAcceptAggchainManagerRole struct {
	OldAggchainManager common.Address
	NewAggchainManager common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAcceptAggchainManagerRole is a free log retrieval operation binding the contract event 0x67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f.
//
// Solidity: event AcceptAggchainManagerRole(address oldAggchainManager, address newAggchainManager)
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousFilterer) FilterAcceptAggchainManagerRole(opts *bind.FilterOpts) (*IaggchainbaseeventspreviousAcceptAggchainManagerRoleIterator, error) {

	logs, sub, err := _Iaggchainbaseeventsprevious.contract.FilterLogs(opts, "AcceptAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventspreviousAcceptAggchainManagerRoleIterator{contract: _Iaggchainbaseeventsprevious.contract, event: "AcceptAggchainManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAggchainManagerRole is a free log subscription operation binding the contract event 0x67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f.
//
// Solidity: event AcceptAggchainManagerRole(address oldAggchainManager, address newAggchainManager)
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousFilterer) WatchAcceptAggchainManagerRole(opts *bind.WatchOpts, sink chan<- *IaggchainbaseeventspreviousAcceptAggchainManagerRole) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbaseeventsprevious.contract.WatchLogs(opts, "AcceptAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseeventspreviousAcceptAggchainManagerRole)
				if err := _Iaggchainbaseeventsprevious.contract.UnpackLog(event, "AcceptAggchainManagerRole", log); err != nil {
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
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousFilterer) ParseAcceptAggchainManagerRole(log types.Log) (*IaggchainbaseeventspreviousAcceptAggchainManagerRole, error) {
	event := new(IaggchainbaseeventspreviousAcceptAggchainManagerRole)
	if err := _Iaggchainbaseeventsprevious.contract.UnpackLog(event, "AcceptAggchainManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseeventspreviousAcceptVKeyManagerRoleIterator is returned from FilterAcceptVKeyManagerRole and is used to iterate over the raw logs and unpacked data for AcceptVKeyManagerRole events raised by the Iaggchainbaseeventsprevious contract.
type IaggchainbaseeventspreviousAcceptVKeyManagerRoleIterator struct {
	Event *IaggchainbaseeventspreviousAcceptVKeyManagerRole // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseeventspreviousAcceptVKeyManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseeventspreviousAcceptVKeyManagerRole)
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
		it.Event = new(IaggchainbaseeventspreviousAcceptVKeyManagerRole)
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
func (it *IaggchainbaseeventspreviousAcceptVKeyManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseeventspreviousAcceptVKeyManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseeventspreviousAcceptVKeyManagerRole represents a AcceptVKeyManagerRole event raised by the Iaggchainbaseeventsprevious contract.
type IaggchainbaseeventspreviousAcceptVKeyManagerRole struct {
	OldVKeyManager common.Address
	NewVKeyManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAcceptVKeyManagerRole is a free log retrieval operation binding the contract event 0xbb3b066505f14a628f4ba4187a046abd4dd17e96591d7a9ed31c91c79322ffe2.
//
// Solidity: event AcceptVKeyManagerRole(address oldVKeyManager, address newVKeyManager)
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousFilterer) FilterAcceptVKeyManagerRole(opts *bind.FilterOpts) (*IaggchainbaseeventspreviousAcceptVKeyManagerRoleIterator, error) {

	logs, sub, err := _Iaggchainbaseeventsprevious.contract.FilterLogs(opts, "AcceptVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventspreviousAcceptVKeyManagerRoleIterator{contract: _Iaggchainbaseeventsprevious.contract, event: "AcceptVKeyManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptVKeyManagerRole is a free log subscription operation binding the contract event 0xbb3b066505f14a628f4ba4187a046abd4dd17e96591d7a9ed31c91c79322ffe2.
//
// Solidity: event AcceptVKeyManagerRole(address oldVKeyManager, address newVKeyManager)
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousFilterer) WatchAcceptVKeyManagerRole(opts *bind.WatchOpts, sink chan<- *IaggchainbaseeventspreviousAcceptVKeyManagerRole) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbaseeventsprevious.contract.WatchLogs(opts, "AcceptVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseeventspreviousAcceptVKeyManagerRole)
				if err := _Iaggchainbaseeventsprevious.contract.UnpackLog(event, "AcceptVKeyManagerRole", log); err != nil {
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
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousFilterer) ParseAcceptVKeyManagerRole(log types.Log) (*IaggchainbaseeventspreviousAcceptVKeyManagerRole, error) {
	event := new(IaggchainbaseeventspreviousAcceptVKeyManagerRole)
	if err := _Iaggchainbaseeventsprevious.contract.UnpackLog(event, "AcceptVKeyManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseeventspreviousAddAggchainVKeyIterator is returned from FilterAddAggchainVKey and is used to iterate over the raw logs and unpacked data for AddAggchainVKey events raised by the Iaggchainbaseeventsprevious contract.
type IaggchainbaseeventspreviousAddAggchainVKeyIterator struct {
	Event *IaggchainbaseeventspreviousAddAggchainVKey // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseeventspreviousAddAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseeventspreviousAddAggchainVKey)
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
		it.Event = new(IaggchainbaseeventspreviousAddAggchainVKey)
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
func (it *IaggchainbaseeventspreviousAddAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseeventspreviousAddAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseeventspreviousAddAggchainVKey represents a AddAggchainVKey event raised by the Iaggchainbaseeventsprevious contract.
type IaggchainbaseeventspreviousAddAggchainVKey struct {
	Selector        [4]byte
	NewAggchainVKey [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddAggchainVKey is a free log retrieval operation binding the contract event 0x6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a79.
//
// Solidity: event AddAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousFilterer) FilterAddAggchainVKey(opts *bind.FilterOpts) (*IaggchainbaseeventspreviousAddAggchainVKeyIterator, error) {

	logs, sub, err := _Iaggchainbaseeventsprevious.contract.FilterLogs(opts, "AddAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventspreviousAddAggchainVKeyIterator{contract: _Iaggchainbaseeventsprevious.contract, event: "AddAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchAddAggchainVKey is a free log subscription operation binding the contract event 0x6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a79.
//
// Solidity: event AddAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousFilterer) WatchAddAggchainVKey(opts *bind.WatchOpts, sink chan<- *IaggchainbaseeventspreviousAddAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbaseeventsprevious.contract.WatchLogs(opts, "AddAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseeventspreviousAddAggchainVKey)
				if err := _Iaggchainbaseeventsprevious.contract.UnpackLog(event, "AddAggchainVKey", log); err != nil {
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
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousFilterer) ParseAddAggchainVKey(log types.Log) (*IaggchainbaseeventspreviousAddAggchainVKey, error) {
	event := new(IaggchainbaseeventspreviousAddAggchainVKey)
	if err := _Iaggchainbaseeventsprevious.contract.UnpackLog(event, "AddAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseeventspreviousDisableUseDefaultGatewayFlagIterator is returned from FilterDisableUseDefaultGatewayFlag and is used to iterate over the raw logs and unpacked data for DisableUseDefaultGatewayFlag events raised by the Iaggchainbaseeventsprevious contract.
type IaggchainbaseeventspreviousDisableUseDefaultGatewayFlagIterator struct {
	Event *IaggchainbaseeventspreviousDisableUseDefaultGatewayFlag // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseeventspreviousDisableUseDefaultGatewayFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseeventspreviousDisableUseDefaultGatewayFlag)
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
		it.Event = new(IaggchainbaseeventspreviousDisableUseDefaultGatewayFlag)
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
func (it *IaggchainbaseeventspreviousDisableUseDefaultGatewayFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseeventspreviousDisableUseDefaultGatewayFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseeventspreviousDisableUseDefaultGatewayFlag represents a DisableUseDefaultGatewayFlag event raised by the Iaggchainbaseeventsprevious contract.
type IaggchainbaseeventspreviousDisableUseDefaultGatewayFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisableUseDefaultGatewayFlag is a free log retrieval operation binding the contract event 0x67dd1717a1952df380cb73eecb312e949df6d6a086bd7f88669005341972528e.
//
// Solidity: event DisableUseDefaultGatewayFlag()
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousFilterer) FilterDisableUseDefaultGatewayFlag(opts *bind.FilterOpts) (*IaggchainbaseeventspreviousDisableUseDefaultGatewayFlagIterator, error) {

	logs, sub, err := _Iaggchainbaseeventsprevious.contract.FilterLogs(opts, "DisableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventspreviousDisableUseDefaultGatewayFlagIterator{contract: _Iaggchainbaseeventsprevious.contract, event: "DisableUseDefaultGatewayFlag", logs: logs, sub: sub}, nil
}

// WatchDisableUseDefaultGatewayFlag is a free log subscription operation binding the contract event 0x67dd1717a1952df380cb73eecb312e949df6d6a086bd7f88669005341972528e.
//
// Solidity: event DisableUseDefaultGatewayFlag()
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousFilterer) WatchDisableUseDefaultGatewayFlag(opts *bind.WatchOpts, sink chan<- *IaggchainbaseeventspreviousDisableUseDefaultGatewayFlag) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbaseeventsprevious.contract.WatchLogs(opts, "DisableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseeventspreviousDisableUseDefaultGatewayFlag)
				if err := _Iaggchainbaseeventsprevious.contract.UnpackLog(event, "DisableUseDefaultGatewayFlag", log); err != nil {
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
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousFilterer) ParseDisableUseDefaultGatewayFlag(log types.Log) (*IaggchainbaseeventspreviousDisableUseDefaultGatewayFlag, error) {
	event := new(IaggchainbaseeventspreviousDisableUseDefaultGatewayFlag)
	if err := _Iaggchainbaseeventsprevious.contract.UnpackLog(event, "DisableUseDefaultGatewayFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseeventspreviousEnableUseDefaultGatewayFlagIterator is returned from FilterEnableUseDefaultGatewayFlag and is used to iterate over the raw logs and unpacked data for EnableUseDefaultGatewayFlag events raised by the Iaggchainbaseeventsprevious contract.
type IaggchainbaseeventspreviousEnableUseDefaultGatewayFlagIterator struct {
	Event *IaggchainbaseeventspreviousEnableUseDefaultGatewayFlag // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseeventspreviousEnableUseDefaultGatewayFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseeventspreviousEnableUseDefaultGatewayFlag)
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
		it.Event = new(IaggchainbaseeventspreviousEnableUseDefaultGatewayFlag)
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
func (it *IaggchainbaseeventspreviousEnableUseDefaultGatewayFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseeventspreviousEnableUseDefaultGatewayFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseeventspreviousEnableUseDefaultGatewayFlag represents a EnableUseDefaultGatewayFlag event raised by the Iaggchainbaseeventsprevious contract.
type IaggchainbaseeventspreviousEnableUseDefaultGatewayFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEnableUseDefaultGatewayFlag is a free log retrieval operation binding the contract event 0xb6563aed80fde357e737eb0d19f246a58cb6bfd469933d05701ecbad0f2dca84.
//
// Solidity: event EnableUseDefaultGatewayFlag()
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousFilterer) FilterEnableUseDefaultGatewayFlag(opts *bind.FilterOpts) (*IaggchainbaseeventspreviousEnableUseDefaultGatewayFlagIterator, error) {

	logs, sub, err := _Iaggchainbaseeventsprevious.contract.FilterLogs(opts, "EnableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventspreviousEnableUseDefaultGatewayFlagIterator{contract: _Iaggchainbaseeventsprevious.contract, event: "EnableUseDefaultGatewayFlag", logs: logs, sub: sub}, nil
}

// WatchEnableUseDefaultGatewayFlag is a free log subscription operation binding the contract event 0xb6563aed80fde357e737eb0d19f246a58cb6bfd469933d05701ecbad0f2dca84.
//
// Solidity: event EnableUseDefaultGatewayFlag()
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousFilterer) WatchEnableUseDefaultGatewayFlag(opts *bind.WatchOpts, sink chan<- *IaggchainbaseeventspreviousEnableUseDefaultGatewayFlag) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbaseeventsprevious.contract.WatchLogs(opts, "EnableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseeventspreviousEnableUseDefaultGatewayFlag)
				if err := _Iaggchainbaseeventsprevious.contract.UnpackLog(event, "EnableUseDefaultGatewayFlag", log); err != nil {
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
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousFilterer) ParseEnableUseDefaultGatewayFlag(log types.Log) (*IaggchainbaseeventspreviousEnableUseDefaultGatewayFlag, error) {
	event := new(IaggchainbaseeventspreviousEnableUseDefaultGatewayFlag)
	if err := _Iaggchainbaseeventsprevious.contract.UnpackLog(event, "EnableUseDefaultGatewayFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseeventspreviousTransferAggchainManagerRoleIterator is returned from FilterTransferAggchainManagerRole and is used to iterate over the raw logs and unpacked data for TransferAggchainManagerRole events raised by the Iaggchainbaseeventsprevious contract.
type IaggchainbaseeventspreviousTransferAggchainManagerRoleIterator struct {
	Event *IaggchainbaseeventspreviousTransferAggchainManagerRole // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseeventspreviousTransferAggchainManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseeventspreviousTransferAggchainManagerRole)
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
		it.Event = new(IaggchainbaseeventspreviousTransferAggchainManagerRole)
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
func (it *IaggchainbaseeventspreviousTransferAggchainManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseeventspreviousTransferAggchainManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseeventspreviousTransferAggchainManagerRole represents a TransferAggchainManagerRole event raised by the Iaggchainbaseeventsprevious contract.
type IaggchainbaseeventspreviousTransferAggchainManagerRole struct {
	CurrentAggchainManager    common.Address
	NewPendingAggchainManager common.Address
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterTransferAggchainManagerRole is a free log retrieval operation binding the contract event 0xa3d8e5d045432398be30f83ce7c35a7bfc220c1b66cc5bf3f4dd4d539d93fab6.
//
// Solidity: event TransferAggchainManagerRole(address currentAggchainManager, address newPendingAggchainManager)
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousFilterer) FilterTransferAggchainManagerRole(opts *bind.FilterOpts) (*IaggchainbaseeventspreviousTransferAggchainManagerRoleIterator, error) {

	logs, sub, err := _Iaggchainbaseeventsprevious.contract.FilterLogs(opts, "TransferAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventspreviousTransferAggchainManagerRoleIterator{contract: _Iaggchainbaseeventsprevious.contract, event: "TransferAggchainManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferAggchainManagerRole is a free log subscription operation binding the contract event 0xa3d8e5d045432398be30f83ce7c35a7bfc220c1b66cc5bf3f4dd4d539d93fab6.
//
// Solidity: event TransferAggchainManagerRole(address currentAggchainManager, address newPendingAggchainManager)
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousFilterer) WatchTransferAggchainManagerRole(opts *bind.WatchOpts, sink chan<- *IaggchainbaseeventspreviousTransferAggchainManagerRole) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbaseeventsprevious.contract.WatchLogs(opts, "TransferAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseeventspreviousTransferAggchainManagerRole)
				if err := _Iaggchainbaseeventsprevious.contract.UnpackLog(event, "TransferAggchainManagerRole", log); err != nil {
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
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousFilterer) ParseTransferAggchainManagerRole(log types.Log) (*IaggchainbaseeventspreviousTransferAggchainManagerRole, error) {
	event := new(IaggchainbaseeventspreviousTransferAggchainManagerRole)
	if err := _Iaggchainbaseeventsprevious.contract.UnpackLog(event, "TransferAggchainManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseeventspreviousTransferVKeyManagerRoleIterator is returned from FilterTransferVKeyManagerRole and is used to iterate over the raw logs and unpacked data for TransferVKeyManagerRole events raised by the Iaggchainbaseeventsprevious contract.
type IaggchainbaseeventspreviousTransferVKeyManagerRoleIterator struct {
	Event *IaggchainbaseeventspreviousTransferVKeyManagerRole // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseeventspreviousTransferVKeyManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseeventspreviousTransferVKeyManagerRole)
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
		it.Event = new(IaggchainbaseeventspreviousTransferVKeyManagerRole)
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
func (it *IaggchainbaseeventspreviousTransferVKeyManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseeventspreviousTransferVKeyManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseeventspreviousTransferVKeyManagerRole represents a TransferVKeyManagerRole event raised by the Iaggchainbaseeventsprevious contract.
type IaggchainbaseeventspreviousTransferVKeyManagerRole struct {
	CurrentVKeyManager    common.Address
	NewPendingVKeyManager common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterTransferVKeyManagerRole is a free log retrieval operation binding the contract event 0xc54ae01017d0b80bd8af833f66387d6eb547dc16c8206faf13d0b72764aab8b2.
//
// Solidity: event TransferVKeyManagerRole(address currentVKeyManager, address newPendingVKeyManager)
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousFilterer) FilterTransferVKeyManagerRole(opts *bind.FilterOpts) (*IaggchainbaseeventspreviousTransferVKeyManagerRoleIterator, error) {

	logs, sub, err := _Iaggchainbaseeventsprevious.contract.FilterLogs(opts, "TransferVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventspreviousTransferVKeyManagerRoleIterator{contract: _Iaggchainbaseeventsprevious.contract, event: "TransferVKeyManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferVKeyManagerRole is a free log subscription operation binding the contract event 0xc54ae01017d0b80bd8af833f66387d6eb547dc16c8206faf13d0b72764aab8b2.
//
// Solidity: event TransferVKeyManagerRole(address currentVKeyManager, address newPendingVKeyManager)
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousFilterer) WatchTransferVKeyManagerRole(opts *bind.WatchOpts, sink chan<- *IaggchainbaseeventspreviousTransferVKeyManagerRole) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbaseeventsprevious.contract.WatchLogs(opts, "TransferVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseeventspreviousTransferVKeyManagerRole)
				if err := _Iaggchainbaseeventsprevious.contract.UnpackLog(event, "TransferVKeyManagerRole", log); err != nil {
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
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousFilterer) ParseTransferVKeyManagerRole(log types.Log) (*IaggchainbaseeventspreviousTransferVKeyManagerRole, error) {
	event := new(IaggchainbaseeventspreviousTransferVKeyManagerRole)
	if err := _Iaggchainbaseeventsprevious.contract.UnpackLog(event, "TransferVKeyManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseeventspreviousUpdateAggchainVKeyIterator is returned from FilterUpdateAggchainVKey and is used to iterate over the raw logs and unpacked data for UpdateAggchainVKey events raised by the Iaggchainbaseeventsprevious contract.
type IaggchainbaseeventspreviousUpdateAggchainVKeyIterator struct {
	Event *IaggchainbaseeventspreviousUpdateAggchainVKey // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseeventspreviousUpdateAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseeventspreviousUpdateAggchainVKey)
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
		it.Event = new(IaggchainbaseeventspreviousUpdateAggchainVKey)
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
func (it *IaggchainbaseeventspreviousUpdateAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseeventspreviousUpdateAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseeventspreviousUpdateAggchainVKey represents a UpdateAggchainVKey event raised by the Iaggchainbaseeventsprevious contract.
type IaggchainbaseeventspreviousUpdateAggchainVKey struct {
	Selector             [4]byte
	PreviousAggchainVKey [32]byte
	NewAggchainVKey      [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateAggchainVKey is a free log retrieval operation binding the contract event 0x0aa5f73c189fb0b0a7cc98ae5fa89dfc16595480396208483518178435ed5b4f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 previousAggchainVKey, bytes32 newAggchainVKey)
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousFilterer) FilterUpdateAggchainVKey(opts *bind.FilterOpts) (*IaggchainbaseeventspreviousUpdateAggchainVKeyIterator, error) {

	logs, sub, err := _Iaggchainbaseeventsprevious.contract.FilterLogs(opts, "UpdateAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventspreviousUpdateAggchainVKeyIterator{contract: _Iaggchainbaseeventsprevious.contract, event: "UpdateAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchUpdateAggchainVKey is a free log subscription operation binding the contract event 0x0aa5f73c189fb0b0a7cc98ae5fa89dfc16595480396208483518178435ed5b4f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 previousAggchainVKey, bytes32 newAggchainVKey)
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousFilterer) WatchUpdateAggchainVKey(opts *bind.WatchOpts, sink chan<- *IaggchainbaseeventspreviousUpdateAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbaseeventsprevious.contract.WatchLogs(opts, "UpdateAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseeventspreviousUpdateAggchainVKey)
				if err := _Iaggchainbaseeventsprevious.contract.UnpackLog(event, "UpdateAggchainVKey", log); err != nil {
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
func (_Iaggchainbaseeventsprevious *IaggchainbaseeventspreviousFilterer) ParseUpdateAggchainVKey(log types.Log) (*IaggchainbaseeventspreviousUpdateAggchainVKey, error) {
	event := new(IaggchainbaseeventspreviousUpdateAggchainVKey)
	if err := _Iaggchainbaseeventsprevious.contract.UnpackLog(event, "UpdateAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

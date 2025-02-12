// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iagglayergatewayevents

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

// IagglayergatewayeventsMetaData contains all meta data concerning the Iagglayergatewayevents contract.
var IagglayergatewayeventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newVKey\",\"type\":\"bytes32\"}],\"name\":\"AddDefaultAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pessimisticVKey\",\"type\":\"bytes32\"}],\"name\":\"RouteAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"RouteFrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newVKey\",\"type\":\"bytes32\"}],\"name\":\"UpdateDefaultAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newPessimisticVKey\",\"type\":\"bytes32\"}],\"name\":\"UpdatePessimisticVKey\",\"type\":\"event\"}]",
}

// IagglayergatewayeventsABI is the input ABI used to generate the binding from.
// Deprecated: Use IagglayergatewayeventsMetaData.ABI instead.
var IagglayergatewayeventsABI = IagglayergatewayeventsMetaData.ABI

// Iagglayergatewayevents is an auto generated Go binding around an Ethereum contract.
type Iagglayergatewayevents struct {
	IagglayergatewayeventsCaller     // Read-only binding to the contract
	IagglayergatewayeventsTransactor // Write-only binding to the contract
	IagglayergatewayeventsFilterer   // Log filterer for contract events
}

// IagglayergatewayeventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IagglayergatewayeventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IagglayergatewayeventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IagglayergatewayeventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IagglayergatewayeventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IagglayergatewayeventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IagglayergatewayeventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IagglayergatewayeventsSession struct {
	Contract     *Iagglayergatewayevents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IagglayergatewayeventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IagglayergatewayeventsCallerSession struct {
	Contract *IagglayergatewayeventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// IagglayergatewayeventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IagglayergatewayeventsTransactorSession struct {
	Contract     *IagglayergatewayeventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// IagglayergatewayeventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IagglayergatewayeventsRaw struct {
	Contract *Iagglayergatewayevents // Generic contract binding to access the raw methods on
}

// IagglayergatewayeventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IagglayergatewayeventsCallerRaw struct {
	Contract *IagglayergatewayeventsCaller // Generic read-only contract binding to access the raw methods on
}

// IagglayergatewayeventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IagglayergatewayeventsTransactorRaw struct {
	Contract *IagglayergatewayeventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIagglayergatewayevents creates a new instance of Iagglayergatewayevents, bound to a specific deployed contract.
func NewIagglayergatewayevents(address common.Address, backend bind.ContractBackend) (*Iagglayergatewayevents, error) {
	contract, err := bindIagglayergatewayevents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iagglayergatewayevents{IagglayergatewayeventsCaller: IagglayergatewayeventsCaller{contract: contract}, IagglayergatewayeventsTransactor: IagglayergatewayeventsTransactor{contract: contract}, IagglayergatewayeventsFilterer: IagglayergatewayeventsFilterer{contract: contract}}, nil
}

// NewIagglayergatewayeventsCaller creates a new read-only instance of Iagglayergatewayevents, bound to a specific deployed contract.
func NewIagglayergatewayeventsCaller(address common.Address, caller bind.ContractCaller) (*IagglayergatewayeventsCaller, error) {
	contract, err := bindIagglayergatewayevents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayeventsCaller{contract: contract}, nil
}

// NewIagglayergatewayeventsTransactor creates a new write-only instance of Iagglayergatewayevents, bound to a specific deployed contract.
func NewIagglayergatewayeventsTransactor(address common.Address, transactor bind.ContractTransactor) (*IagglayergatewayeventsTransactor, error) {
	contract, err := bindIagglayergatewayevents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayeventsTransactor{contract: contract}, nil
}

// NewIagglayergatewayeventsFilterer creates a new log filterer instance of Iagglayergatewayevents, bound to a specific deployed contract.
func NewIagglayergatewayeventsFilterer(address common.Address, filterer bind.ContractFilterer) (*IagglayergatewayeventsFilterer, error) {
	contract, err := bindIagglayergatewayevents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayeventsFilterer{contract: contract}, nil
}

// bindIagglayergatewayevents binds a generic wrapper to an already deployed contract.
func bindIagglayergatewayevents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IagglayergatewayeventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iagglayergatewayevents *IagglayergatewayeventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iagglayergatewayevents.Contract.IagglayergatewayeventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iagglayergatewayevents *IagglayergatewayeventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iagglayergatewayevents.Contract.IagglayergatewayeventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iagglayergatewayevents *IagglayergatewayeventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iagglayergatewayevents.Contract.IagglayergatewayeventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iagglayergatewayevents *IagglayergatewayeventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iagglayergatewayevents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iagglayergatewayevents *IagglayergatewayeventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iagglayergatewayevents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iagglayergatewayevents *IagglayergatewayeventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iagglayergatewayevents.Contract.contract.Transact(opts, method, params...)
}

// IagglayergatewayeventsAddDefaultAggchainVKeyIterator is returned from FilterAddDefaultAggchainVKey and is used to iterate over the raw logs and unpacked data for AddDefaultAggchainVKey events raised by the Iagglayergatewayevents contract.
type IagglayergatewayeventsAddDefaultAggchainVKeyIterator struct {
	Event *IagglayergatewayeventsAddDefaultAggchainVKey // Event containing the contract specifics and raw log

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
func (it *IagglayergatewayeventsAddDefaultAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IagglayergatewayeventsAddDefaultAggchainVKey)
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
		it.Event = new(IagglayergatewayeventsAddDefaultAggchainVKey)
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
func (it *IagglayergatewayeventsAddDefaultAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IagglayergatewayeventsAddDefaultAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IagglayergatewayeventsAddDefaultAggchainVKey represents a AddDefaultAggchainVKey event raised by the Iagglayergatewayevents contract.
type IagglayergatewayeventsAddDefaultAggchainVKey struct {
	Selector [4]byte
	NewVKey  [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddDefaultAggchainVKey is a free log retrieval operation binding the contract event 0x64b3528974574a41ddc830adc96440dccb99e64ede6b863194b951bfbd17e1bc.
//
// Solidity: event AddDefaultAggchainVKey(bytes4 selector, bytes32 newVKey)
func (_Iagglayergatewayevents *IagglayergatewayeventsFilterer) FilterAddDefaultAggchainVKey(opts *bind.FilterOpts) (*IagglayergatewayeventsAddDefaultAggchainVKeyIterator, error) {

	logs, sub, err := _Iagglayergatewayevents.contract.FilterLogs(opts, "AddDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayeventsAddDefaultAggchainVKeyIterator{contract: _Iagglayergatewayevents.contract, event: "AddDefaultAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchAddDefaultAggchainVKey is a free log subscription operation binding the contract event 0x64b3528974574a41ddc830adc96440dccb99e64ede6b863194b951bfbd17e1bc.
//
// Solidity: event AddDefaultAggchainVKey(bytes4 selector, bytes32 newVKey)
func (_Iagglayergatewayevents *IagglayergatewayeventsFilterer) WatchAddDefaultAggchainVKey(opts *bind.WatchOpts, sink chan<- *IagglayergatewayeventsAddDefaultAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Iagglayergatewayevents.contract.WatchLogs(opts, "AddDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IagglayergatewayeventsAddDefaultAggchainVKey)
				if err := _Iagglayergatewayevents.contract.UnpackLog(event, "AddDefaultAggchainVKey", log); err != nil {
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
func (_Iagglayergatewayevents *IagglayergatewayeventsFilterer) ParseAddDefaultAggchainVKey(log types.Log) (*IagglayergatewayeventsAddDefaultAggchainVKey, error) {
	event := new(IagglayergatewayeventsAddDefaultAggchainVKey)
	if err := _Iagglayergatewayevents.contract.UnpackLog(event, "AddDefaultAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IagglayergatewayeventsRouteAddedIterator is returned from FilterRouteAdded and is used to iterate over the raw logs and unpacked data for RouteAdded events raised by the Iagglayergatewayevents contract.
type IagglayergatewayeventsRouteAddedIterator struct {
	Event *IagglayergatewayeventsRouteAdded // Event containing the contract specifics and raw log

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
func (it *IagglayergatewayeventsRouteAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IagglayergatewayeventsRouteAdded)
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
		it.Event = new(IagglayergatewayeventsRouteAdded)
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
func (it *IagglayergatewayeventsRouteAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IagglayergatewayeventsRouteAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IagglayergatewayeventsRouteAdded represents a RouteAdded event raised by the Iagglayergatewayevents contract.
type IagglayergatewayeventsRouteAdded struct {
	Selector        [4]byte
	Verifier        common.Address
	PessimisticVKey [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRouteAdded is a free log retrieval operation binding the contract event 0x4c4455a30528c319643ae6fc35f3a0fcabde6c015f2a1a5270382c4a190b0fa3.
//
// Solidity: event RouteAdded(bytes4 selector, address verifier, bytes32 pessimisticVKey)
func (_Iagglayergatewayevents *IagglayergatewayeventsFilterer) FilterRouteAdded(opts *bind.FilterOpts) (*IagglayergatewayeventsRouteAddedIterator, error) {

	logs, sub, err := _Iagglayergatewayevents.contract.FilterLogs(opts, "RouteAdded")
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayeventsRouteAddedIterator{contract: _Iagglayergatewayevents.contract, event: "RouteAdded", logs: logs, sub: sub}, nil
}

// WatchRouteAdded is a free log subscription operation binding the contract event 0x4c4455a30528c319643ae6fc35f3a0fcabde6c015f2a1a5270382c4a190b0fa3.
//
// Solidity: event RouteAdded(bytes4 selector, address verifier, bytes32 pessimisticVKey)
func (_Iagglayergatewayevents *IagglayergatewayeventsFilterer) WatchRouteAdded(opts *bind.WatchOpts, sink chan<- *IagglayergatewayeventsRouteAdded) (event.Subscription, error) {

	logs, sub, err := _Iagglayergatewayevents.contract.WatchLogs(opts, "RouteAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IagglayergatewayeventsRouteAdded)
				if err := _Iagglayergatewayevents.contract.UnpackLog(event, "RouteAdded", log); err != nil {
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
func (_Iagglayergatewayevents *IagglayergatewayeventsFilterer) ParseRouteAdded(log types.Log) (*IagglayergatewayeventsRouteAdded, error) {
	event := new(IagglayergatewayeventsRouteAdded)
	if err := _Iagglayergatewayevents.contract.UnpackLog(event, "RouteAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IagglayergatewayeventsRouteFrozenIterator is returned from FilterRouteFrozen and is used to iterate over the raw logs and unpacked data for RouteFrozen events raised by the Iagglayergatewayevents contract.
type IagglayergatewayeventsRouteFrozenIterator struct {
	Event *IagglayergatewayeventsRouteFrozen // Event containing the contract specifics and raw log

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
func (it *IagglayergatewayeventsRouteFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IagglayergatewayeventsRouteFrozen)
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
		it.Event = new(IagglayergatewayeventsRouteFrozen)
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
func (it *IagglayergatewayeventsRouteFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IagglayergatewayeventsRouteFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IagglayergatewayeventsRouteFrozen represents a RouteFrozen event raised by the Iagglayergatewayevents contract.
type IagglayergatewayeventsRouteFrozen struct {
	Selector [4]byte
	Verifier common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRouteFrozen is a free log retrieval operation binding the contract event 0x63ad2363b183cb8bb562b9590c5b4428e2a566260df053db156576d3d171438d.
//
// Solidity: event RouteFrozen(bytes4 selector, address verifier)
func (_Iagglayergatewayevents *IagglayergatewayeventsFilterer) FilterRouteFrozen(opts *bind.FilterOpts) (*IagglayergatewayeventsRouteFrozenIterator, error) {

	logs, sub, err := _Iagglayergatewayevents.contract.FilterLogs(opts, "RouteFrozen")
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayeventsRouteFrozenIterator{contract: _Iagglayergatewayevents.contract, event: "RouteFrozen", logs: logs, sub: sub}, nil
}

// WatchRouteFrozen is a free log subscription operation binding the contract event 0x63ad2363b183cb8bb562b9590c5b4428e2a566260df053db156576d3d171438d.
//
// Solidity: event RouteFrozen(bytes4 selector, address verifier)
func (_Iagglayergatewayevents *IagglayergatewayeventsFilterer) WatchRouteFrozen(opts *bind.WatchOpts, sink chan<- *IagglayergatewayeventsRouteFrozen) (event.Subscription, error) {

	logs, sub, err := _Iagglayergatewayevents.contract.WatchLogs(opts, "RouteFrozen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IagglayergatewayeventsRouteFrozen)
				if err := _Iagglayergatewayevents.contract.UnpackLog(event, "RouteFrozen", log); err != nil {
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

// ParseRouteFrozen is a log parse operation binding the contract event 0x63ad2363b183cb8bb562b9590c5b4428e2a566260df053db156576d3d171438d.
//
// Solidity: event RouteFrozen(bytes4 selector, address verifier)
func (_Iagglayergatewayevents *IagglayergatewayeventsFilterer) ParseRouteFrozen(log types.Log) (*IagglayergatewayeventsRouteFrozen, error) {
	event := new(IagglayergatewayeventsRouteFrozen)
	if err := _Iagglayergatewayevents.contract.UnpackLog(event, "RouteFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IagglayergatewayeventsUpdateDefaultAggchainVKeyIterator is returned from FilterUpdateDefaultAggchainVKey and is used to iterate over the raw logs and unpacked data for UpdateDefaultAggchainVKey events raised by the Iagglayergatewayevents contract.
type IagglayergatewayeventsUpdateDefaultAggchainVKeyIterator struct {
	Event *IagglayergatewayeventsUpdateDefaultAggchainVKey // Event containing the contract specifics and raw log

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
func (it *IagglayergatewayeventsUpdateDefaultAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IagglayergatewayeventsUpdateDefaultAggchainVKey)
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
		it.Event = new(IagglayergatewayeventsUpdateDefaultAggchainVKey)
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
func (it *IagglayergatewayeventsUpdateDefaultAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IagglayergatewayeventsUpdateDefaultAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IagglayergatewayeventsUpdateDefaultAggchainVKey represents a UpdateDefaultAggchainVKey event raised by the Iagglayergatewayevents contract.
type IagglayergatewayeventsUpdateDefaultAggchainVKey struct {
	Selector [4]byte
	NewVKey  [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateDefaultAggchainVKey is a free log retrieval operation binding the contract event 0x3d81518c9943e29af3aa7c037594823b2dee8d8a8136d0eb25721ced48eb74e6.
//
// Solidity: event UpdateDefaultAggchainVKey(bytes4 selector, bytes32 newVKey)
func (_Iagglayergatewayevents *IagglayergatewayeventsFilterer) FilterUpdateDefaultAggchainVKey(opts *bind.FilterOpts) (*IagglayergatewayeventsUpdateDefaultAggchainVKeyIterator, error) {

	logs, sub, err := _Iagglayergatewayevents.contract.FilterLogs(opts, "UpdateDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayeventsUpdateDefaultAggchainVKeyIterator{contract: _Iagglayergatewayevents.contract, event: "UpdateDefaultAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchUpdateDefaultAggchainVKey is a free log subscription operation binding the contract event 0x3d81518c9943e29af3aa7c037594823b2dee8d8a8136d0eb25721ced48eb74e6.
//
// Solidity: event UpdateDefaultAggchainVKey(bytes4 selector, bytes32 newVKey)
func (_Iagglayergatewayevents *IagglayergatewayeventsFilterer) WatchUpdateDefaultAggchainVKey(opts *bind.WatchOpts, sink chan<- *IagglayergatewayeventsUpdateDefaultAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Iagglayergatewayevents.contract.WatchLogs(opts, "UpdateDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IagglayergatewayeventsUpdateDefaultAggchainVKey)
				if err := _Iagglayergatewayevents.contract.UnpackLog(event, "UpdateDefaultAggchainVKey", log); err != nil {
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

// ParseUpdateDefaultAggchainVKey is a log parse operation binding the contract event 0x3d81518c9943e29af3aa7c037594823b2dee8d8a8136d0eb25721ced48eb74e6.
//
// Solidity: event UpdateDefaultAggchainVKey(bytes4 selector, bytes32 newVKey)
func (_Iagglayergatewayevents *IagglayergatewayeventsFilterer) ParseUpdateDefaultAggchainVKey(log types.Log) (*IagglayergatewayeventsUpdateDefaultAggchainVKey, error) {
	event := new(IagglayergatewayeventsUpdateDefaultAggchainVKey)
	if err := _Iagglayergatewayevents.contract.UnpackLog(event, "UpdateDefaultAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IagglayergatewayeventsUpdatePessimisticVKeyIterator is returned from FilterUpdatePessimisticVKey and is used to iterate over the raw logs and unpacked data for UpdatePessimisticVKey events raised by the Iagglayergatewayevents contract.
type IagglayergatewayeventsUpdatePessimisticVKeyIterator struct {
	Event *IagglayergatewayeventsUpdatePessimisticVKey // Event containing the contract specifics and raw log

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
func (it *IagglayergatewayeventsUpdatePessimisticVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IagglayergatewayeventsUpdatePessimisticVKey)
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
		it.Event = new(IagglayergatewayeventsUpdatePessimisticVKey)
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
func (it *IagglayergatewayeventsUpdatePessimisticVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IagglayergatewayeventsUpdatePessimisticVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IagglayergatewayeventsUpdatePessimisticVKey represents a UpdatePessimisticVKey event raised by the Iagglayergatewayevents contract.
type IagglayergatewayeventsUpdatePessimisticVKey struct {
	Selector           [4]byte
	Verifier           common.Address
	NewPessimisticVKey [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdatePessimisticVKey is a free log retrieval operation binding the contract event 0x1efbcfeb43381d70549495ea01d14d747d8fe4c19e19a430eb36ba478065806a.
//
// Solidity: event UpdatePessimisticVKey(bytes4 selector, address verifier, bytes32 newPessimisticVKey)
func (_Iagglayergatewayevents *IagglayergatewayeventsFilterer) FilterUpdatePessimisticVKey(opts *bind.FilterOpts) (*IagglayergatewayeventsUpdatePessimisticVKeyIterator, error) {

	logs, sub, err := _Iagglayergatewayevents.contract.FilterLogs(opts, "UpdatePessimisticVKey")
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayeventsUpdatePessimisticVKeyIterator{contract: _Iagglayergatewayevents.contract, event: "UpdatePessimisticVKey", logs: logs, sub: sub}, nil
}

// WatchUpdatePessimisticVKey is a free log subscription operation binding the contract event 0x1efbcfeb43381d70549495ea01d14d747d8fe4c19e19a430eb36ba478065806a.
//
// Solidity: event UpdatePessimisticVKey(bytes4 selector, address verifier, bytes32 newPessimisticVKey)
func (_Iagglayergatewayevents *IagglayergatewayeventsFilterer) WatchUpdatePessimisticVKey(opts *bind.WatchOpts, sink chan<- *IagglayergatewayeventsUpdatePessimisticVKey) (event.Subscription, error) {

	logs, sub, err := _Iagglayergatewayevents.contract.WatchLogs(opts, "UpdatePessimisticVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IagglayergatewayeventsUpdatePessimisticVKey)
				if err := _Iagglayergatewayevents.contract.UnpackLog(event, "UpdatePessimisticVKey", log); err != nil {
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

// ParseUpdatePessimisticVKey is a log parse operation binding the contract event 0x1efbcfeb43381d70549495ea01d14d747d8fe4c19e19a430eb36ba478065806a.
//
// Solidity: event UpdatePessimisticVKey(bytes4 selector, address verifier, bytes32 newPessimisticVKey)
func (_Iagglayergatewayevents *IagglayergatewayeventsFilterer) ParseUpdatePessimisticVKey(log types.Log) (*IagglayergatewayeventsUpdatePessimisticVKey, error) {
	event := new(IagglayergatewayeventsUpdatePessimisticVKey)
	if err := _Iagglayergatewayevents.contract.UnpackLog(event, "UpdatePessimisticVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

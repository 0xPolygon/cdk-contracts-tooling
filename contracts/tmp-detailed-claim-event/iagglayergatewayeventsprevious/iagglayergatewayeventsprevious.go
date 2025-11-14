// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iagglayergatewayeventsprevious

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

// IagglayergatewayeventspreviousMetaData contains all meta data concerning the Iagglayergatewayeventsprevious contract.
var IagglayergatewayeventspreviousMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newVKey\",\"type\":\"bytes32\"}],\"name\":\"AddDefaultAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pessimisticVKey\",\"type\":\"bytes32\"}],\"name\":\"RouteAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pessimisticVKey\",\"type\":\"bytes32\"}],\"name\":\"RouteFrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"UnsetDefaultAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousVKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newVKey\",\"type\":\"bytes32\"}],\"name\":\"UpdateDefaultAggchainVKey\",\"type\":\"event\"}]",
}

// IagglayergatewayeventspreviousABI is the input ABI used to generate the binding from.
// Deprecated: Use IagglayergatewayeventspreviousMetaData.ABI instead.
var IagglayergatewayeventspreviousABI = IagglayergatewayeventspreviousMetaData.ABI

// Iagglayergatewayeventsprevious is an auto generated Go binding around an Ethereum contract.
type Iagglayergatewayeventsprevious struct {
	IagglayergatewayeventspreviousCaller     // Read-only binding to the contract
	IagglayergatewayeventspreviousTransactor // Write-only binding to the contract
	IagglayergatewayeventspreviousFilterer   // Log filterer for contract events
}

// IagglayergatewayeventspreviousCaller is an auto generated read-only Go binding around an Ethereum contract.
type IagglayergatewayeventspreviousCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IagglayergatewayeventspreviousTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IagglayergatewayeventspreviousTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IagglayergatewayeventspreviousFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IagglayergatewayeventspreviousFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IagglayergatewayeventspreviousSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IagglayergatewayeventspreviousSession struct {
	Contract     *Iagglayergatewayeventsprevious // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IagglayergatewayeventspreviousCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IagglayergatewayeventspreviousCallerSession struct {
	Contract *IagglayergatewayeventspreviousCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// IagglayergatewayeventspreviousTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IagglayergatewayeventspreviousTransactorSession struct {
	Contract     *IagglayergatewayeventspreviousTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// IagglayergatewayeventspreviousRaw is an auto generated low-level Go binding around an Ethereum contract.
type IagglayergatewayeventspreviousRaw struct {
	Contract *Iagglayergatewayeventsprevious // Generic contract binding to access the raw methods on
}

// IagglayergatewayeventspreviousCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IagglayergatewayeventspreviousCallerRaw struct {
	Contract *IagglayergatewayeventspreviousCaller // Generic read-only contract binding to access the raw methods on
}

// IagglayergatewayeventspreviousTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IagglayergatewayeventspreviousTransactorRaw struct {
	Contract *IagglayergatewayeventspreviousTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIagglayergatewayeventsprevious creates a new instance of Iagglayergatewayeventsprevious, bound to a specific deployed contract.
func NewIagglayergatewayeventsprevious(address common.Address, backend bind.ContractBackend) (*Iagglayergatewayeventsprevious, error) {
	contract, err := bindIagglayergatewayeventsprevious(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iagglayergatewayeventsprevious{IagglayergatewayeventspreviousCaller: IagglayergatewayeventspreviousCaller{contract: contract}, IagglayergatewayeventspreviousTransactor: IagglayergatewayeventspreviousTransactor{contract: contract}, IagglayergatewayeventspreviousFilterer: IagglayergatewayeventspreviousFilterer{contract: contract}}, nil
}

// NewIagglayergatewayeventspreviousCaller creates a new read-only instance of Iagglayergatewayeventsprevious, bound to a specific deployed contract.
func NewIagglayergatewayeventspreviousCaller(address common.Address, caller bind.ContractCaller) (*IagglayergatewayeventspreviousCaller, error) {
	contract, err := bindIagglayergatewayeventsprevious(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayeventspreviousCaller{contract: contract}, nil
}

// NewIagglayergatewayeventspreviousTransactor creates a new write-only instance of Iagglayergatewayeventsprevious, bound to a specific deployed contract.
func NewIagglayergatewayeventspreviousTransactor(address common.Address, transactor bind.ContractTransactor) (*IagglayergatewayeventspreviousTransactor, error) {
	contract, err := bindIagglayergatewayeventsprevious(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayeventspreviousTransactor{contract: contract}, nil
}

// NewIagglayergatewayeventspreviousFilterer creates a new log filterer instance of Iagglayergatewayeventsprevious, bound to a specific deployed contract.
func NewIagglayergatewayeventspreviousFilterer(address common.Address, filterer bind.ContractFilterer) (*IagglayergatewayeventspreviousFilterer, error) {
	contract, err := bindIagglayergatewayeventsprevious(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayeventspreviousFilterer{contract: contract}, nil
}

// bindIagglayergatewayeventsprevious binds a generic wrapper to an already deployed contract.
func bindIagglayergatewayeventsprevious(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IagglayergatewayeventspreviousMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iagglayergatewayeventsprevious *IagglayergatewayeventspreviousRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iagglayergatewayeventsprevious.Contract.IagglayergatewayeventspreviousCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iagglayergatewayeventsprevious *IagglayergatewayeventspreviousRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iagglayergatewayeventsprevious.Contract.IagglayergatewayeventspreviousTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iagglayergatewayeventsprevious *IagglayergatewayeventspreviousRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iagglayergatewayeventsprevious.Contract.IagglayergatewayeventspreviousTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iagglayergatewayeventsprevious *IagglayergatewayeventspreviousCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iagglayergatewayeventsprevious.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iagglayergatewayeventsprevious *IagglayergatewayeventspreviousTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iagglayergatewayeventsprevious.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iagglayergatewayeventsprevious *IagglayergatewayeventspreviousTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iagglayergatewayeventsprevious.Contract.contract.Transact(opts, method, params...)
}

// IagglayergatewayeventspreviousAddDefaultAggchainVKeyIterator is returned from FilterAddDefaultAggchainVKey and is used to iterate over the raw logs and unpacked data for AddDefaultAggchainVKey events raised by the Iagglayergatewayeventsprevious contract.
type IagglayergatewayeventspreviousAddDefaultAggchainVKeyIterator struct {
	Event *IagglayergatewayeventspreviousAddDefaultAggchainVKey // Event containing the contract specifics and raw log

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
func (it *IagglayergatewayeventspreviousAddDefaultAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IagglayergatewayeventspreviousAddDefaultAggchainVKey)
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
		it.Event = new(IagglayergatewayeventspreviousAddDefaultAggchainVKey)
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
func (it *IagglayergatewayeventspreviousAddDefaultAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IagglayergatewayeventspreviousAddDefaultAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IagglayergatewayeventspreviousAddDefaultAggchainVKey represents a AddDefaultAggchainVKey event raised by the Iagglayergatewayeventsprevious contract.
type IagglayergatewayeventspreviousAddDefaultAggchainVKey struct {
	Selector [4]byte
	NewVKey  [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddDefaultAggchainVKey is a free log retrieval operation binding the contract event 0x64b3528974574a41ddc830adc96440dccb99e64ede6b863194b951bfbd17e1bc.
//
// Solidity: event AddDefaultAggchainVKey(bytes4 selector, bytes32 newVKey)
func (_Iagglayergatewayeventsprevious *IagglayergatewayeventspreviousFilterer) FilterAddDefaultAggchainVKey(opts *bind.FilterOpts) (*IagglayergatewayeventspreviousAddDefaultAggchainVKeyIterator, error) {

	logs, sub, err := _Iagglayergatewayeventsprevious.contract.FilterLogs(opts, "AddDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayeventspreviousAddDefaultAggchainVKeyIterator{contract: _Iagglayergatewayeventsprevious.contract, event: "AddDefaultAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchAddDefaultAggchainVKey is a free log subscription operation binding the contract event 0x64b3528974574a41ddc830adc96440dccb99e64ede6b863194b951bfbd17e1bc.
//
// Solidity: event AddDefaultAggchainVKey(bytes4 selector, bytes32 newVKey)
func (_Iagglayergatewayeventsprevious *IagglayergatewayeventspreviousFilterer) WatchAddDefaultAggchainVKey(opts *bind.WatchOpts, sink chan<- *IagglayergatewayeventspreviousAddDefaultAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Iagglayergatewayeventsprevious.contract.WatchLogs(opts, "AddDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IagglayergatewayeventspreviousAddDefaultAggchainVKey)
				if err := _Iagglayergatewayeventsprevious.contract.UnpackLog(event, "AddDefaultAggchainVKey", log); err != nil {
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
func (_Iagglayergatewayeventsprevious *IagglayergatewayeventspreviousFilterer) ParseAddDefaultAggchainVKey(log types.Log) (*IagglayergatewayeventspreviousAddDefaultAggchainVKey, error) {
	event := new(IagglayergatewayeventspreviousAddDefaultAggchainVKey)
	if err := _Iagglayergatewayeventsprevious.contract.UnpackLog(event, "AddDefaultAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IagglayergatewayeventspreviousRouteAddedIterator is returned from FilterRouteAdded and is used to iterate over the raw logs and unpacked data for RouteAdded events raised by the Iagglayergatewayeventsprevious contract.
type IagglayergatewayeventspreviousRouteAddedIterator struct {
	Event *IagglayergatewayeventspreviousRouteAdded // Event containing the contract specifics and raw log

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
func (it *IagglayergatewayeventspreviousRouteAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IagglayergatewayeventspreviousRouteAdded)
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
		it.Event = new(IagglayergatewayeventspreviousRouteAdded)
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
func (it *IagglayergatewayeventspreviousRouteAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IagglayergatewayeventspreviousRouteAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IagglayergatewayeventspreviousRouteAdded represents a RouteAdded event raised by the Iagglayergatewayeventsprevious contract.
type IagglayergatewayeventspreviousRouteAdded struct {
	Selector        [4]byte
	Verifier        common.Address
	PessimisticVKey [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRouteAdded is a free log retrieval operation binding the contract event 0x4c4455a30528c319643ae6fc35f3a0fcabde6c015f2a1a5270382c4a190b0fa3.
//
// Solidity: event RouteAdded(bytes4 selector, address verifier, bytes32 pessimisticVKey)
func (_Iagglayergatewayeventsprevious *IagglayergatewayeventspreviousFilterer) FilterRouteAdded(opts *bind.FilterOpts) (*IagglayergatewayeventspreviousRouteAddedIterator, error) {

	logs, sub, err := _Iagglayergatewayeventsprevious.contract.FilterLogs(opts, "RouteAdded")
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayeventspreviousRouteAddedIterator{contract: _Iagglayergatewayeventsprevious.contract, event: "RouteAdded", logs: logs, sub: sub}, nil
}

// WatchRouteAdded is a free log subscription operation binding the contract event 0x4c4455a30528c319643ae6fc35f3a0fcabde6c015f2a1a5270382c4a190b0fa3.
//
// Solidity: event RouteAdded(bytes4 selector, address verifier, bytes32 pessimisticVKey)
func (_Iagglayergatewayeventsprevious *IagglayergatewayeventspreviousFilterer) WatchRouteAdded(opts *bind.WatchOpts, sink chan<- *IagglayergatewayeventspreviousRouteAdded) (event.Subscription, error) {

	logs, sub, err := _Iagglayergatewayeventsprevious.contract.WatchLogs(opts, "RouteAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IagglayergatewayeventspreviousRouteAdded)
				if err := _Iagglayergatewayeventsprevious.contract.UnpackLog(event, "RouteAdded", log); err != nil {
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
func (_Iagglayergatewayeventsprevious *IagglayergatewayeventspreviousFilterer) ParseRouteAdded(log types.Log) (*IagglayergatewayeventspreviousRouteAdded, error) {
	event := new(IagglayergatewayeventspreviousRouteAdded)
	if err := _Iagglayergatewayeventsprevious.contract.UnpackLog(event, "RouteAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IagglayergatewayeventspreviousRouteFrozenIterator is returned from FilterRouteFrozen and is used to iterate over the raw logs and unpacked data for RouteFrozen events raised by the Iagglayergatewayeventsprevious contract.
type IagglayergatewayeventspreviousRouteFrozenIterator struct {
	Event *IagglayergatewayeventspreviousRouteFrozen // Event containing the contract specifics and raw log

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
func (it *IagglayergatewayeventspreviousRouteFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IagglayergatewayeventspreviousRouteFrozen)
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
		it.Event = new(IagglayergatewayeventspreviousRouteFrozen)
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
func (it *IagglayergatewayeventspreviousRouteFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IagglayergatewayeventspreviousRouteFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IagglayergatewayeventspreviousRouteFrozen represents a RouteFrozen event raised by the Iagglayergatewayeventsprevious contract.
type IagglayergatewayeventspreviousRouteFrozen struct {
	Selector        [4]byte
	Verifier        common.Address
	PessimisticVKey [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRouteFrozen is a free log retrieval operation binding the contract event 0x6efcfd4090b998a9d58c7f753bdda3291e9d781c02226b9f480b18ef07e0012b.
//
// Solidity: event RouteFrozen(bytes4 selector, address verifier, bytes32 pessimisticVKey)
func (_Iagglayergatewayeventsprevious *IagglayergatewayeventspreviousFilterer) FilterRouteFrozen(opts *bind.FilterOpts) (*IagglayergatewayeventspreviousRouteFrozenIterator, error) {

	logs, sub, err := _Iagglayergatewayeventsprevious.contract.FilterLogs(opts, "RouteFrozen")
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayeventspreviousRouteFrozenIterator{contract: _Iagglayergatewayeventsprevious.contract, event: "RouteFrozen", logs: logs, sub: sub}, nil
}

// WatchRouteFrozen is a free log subscription operation binding the contract event 0x6efcfd4090b998a9d58c7f753bdda3291e9d781c02226b9f480b18ef07e0012b.
//
// Solidity: event RouteFrozen(bytes4 selector, address verifier, bytes32 pessimisticVKey)
func (_Iagglayergatewayeventsprevious *IagglayergatewayeventspreviousFilterer) WatchRouteFrozen(opts *bind.WatchOpts, sink chan<- *IagglayergatewayeventspreviousRouteFrozen) (event.Subscription, error) {

	logs, sub, err := _Iagglayergatewayeventsprevious.contract.WatchLogs(opts, "RouteFrozen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IagglayergatewayeventspreviousRouteFrozen)
				if err := _Iagglayergatewayeventsprevious.contract.UnpackLog(event, "RouteFrozen", log); err != nil {
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
func (_Iagglayergatewayeventsprevious *IagglayergatewayeventspreviousFilterer) ParseRouteFrozen(log types.Log) (*IagglayergatewayeventspreviousRouteFrozen, error) {
	event := new(IagglayergatewayeventspreviousRouteFrozen)
	if err := _Iagglayergatewayeventsprevious.contract.UnpackLog(event, "RouteFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IagglayergatewayeventspreviousUnsetDefaultAggchainVKeyIterator is returned from FilterUnsetDefaultAggchainVKey and is used to iterate over the raw logs and unpacked data for UnsetDefaultAggchainVKey events raised by the Iagglayergatewayeventsprevious contract.
type IagglayergatewayeventspreviousUnsetDefaultAggchainVKeyIterator struct {
	Event *IagglayergatewayeventspreviousUnsetDefaultAggchainVKey // Event containing the contract specifics and raw log

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
func (it *IagglayergatewayeventspreviousUnsetDefaultAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IagglayergatewayeventspreviousUnsetDefaultAggchainVKey)
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
		it.Event = new(IagglayergatewayeventspreviousUnsetDefaultAggchainVKey)
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
func (it *IagglayergatewayeventspreviousUnsetDefaultAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IagglayergatewayeventspreviousUnsetDefaultAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IagglayergatewayeventspreviousUnsetDefaultAggchainVKey represents a UnsetDefaultAggchainVKey event raised by the Iagglayergatewayeventsprevious contract.
type IagglayergatewayeventspreviousUnsetDefaultAggchainVKey struct {
	Selector [4]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnsetDefaultAggchainVKey is a free log retrieval operation binding the contract event 0x43b46d79ee2fbaa8a981b5fd78aaf6fb9e38b18d02e1af17564ce37c67078cab.
//
// Solidity: event UnsetDefaultAggchainVKey(bytes4 selector)
func (_Iagglayergatewayeventsprevious *IagglayergatewayeventspreviousFilterer) FilterUnsetDefaultAggchainVKey(opts *bind.FilterOpts) (*IagglayergatewayeventspreviousUnsetDefaultAggchainVKeyIterator, error) {

	logs, sub, err := _Iagglayergatewayeventsprevious.contract.FilterLogs(opts, "UnsetDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayeventspreviousUnsetDefaultAggchainVKeyIterator{contract: _Iagglayergatewayeventsprevious.contract, event: "UnsetDefaultAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchUnsetDefaultAggchainVKey is a free log subscription operation binding the contract event 0x43b46d79ee2fbaa8a981b5fd78aaf6fb9e38b18d02e1af17564ce37c67078cab.
//
// Solidity: event UnsetDefaultAggchainVKey(bytes4 selector)
func (_Iagglayergatewayeventsprevious *IagglayergatewayeventspreviousFilterer) WatchUnsetDefaultAggchainVKey(opts *bind.WatchOpts, sink chan<- *IagglayergatewayeventspreviousUnsetDefaultAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Iagglayergatewayeventsprevious.contract.WatchLogs(opts, "UnsetDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IagglayergatewayeventspreviousUnsetDefaultAggchainVKey)
				if err := _Iagglayergatewayeventsprevious.contract.UnpackLog(event, "UnsetDefaultAggchainVKey", log); err != nil {
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

// ParseUnsetDefaultAggchainVKey is a log parse operation binding the contract event 0x43b46d79ee2fbaa8a981b5fd78aaf6fb9e38b18d02e1af17564ce37c67078cab.
//
// Solidity: event UnsetDefaultAggchainVKey(bytes4 selector)
func (_Iagglayergatewayeventsprevious *IagglayergatewayeventspreviousFilterer) ParseUnsetDefaultAggchainVKey(log types.Log) (*IagglayergatewayeventspreviousUnsetDefaultAggchainVKey, error) {
	event := new(IagglayergatewayeventspreviousUnsetDefaultAggchainVKey)
	if err := _Iagglayergatewayeventsprevious.contract.UnpackLog(event, "UnsetDefaultAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IagglayergatewayeventspreviousUpdateDefaultAggchainVKeyIterator is returned from FilterUpdateDefaultAggchainVKey and is used to iterate over the raw logs and unpacked data for UpdateDefaultAggchainVKey events raised by the Iagglayergatewayeventsprevious contract.
type IagglayergatewayeventspreviousUpdateDefaultAggchainVKeyIterator struct {
	Event *IagglayergatewayeventspreviousUpdateDefaultAggchainVKey // Event containing the contract specifics and raw log

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
func (it *IagglayergatewayeventspreviousUpdateDefaultAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IagglayergatewayeventspreviousUpdateDefaultAggchainVKey)
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
		it.Event = new(IagglayergatewayeventspreviousUpdateDefaultAggchainVKey)
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
func (it *IagglayergatewayeventspreviousUpdateDefaultAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IagglayergatewayeventspreviousUpdateDefaultAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IagglayergatewayeventspreviousUpdateDefaultAggchainVKey represents a UpdateDefaultAggchainVKey event raised by the Iagglayergatewayeventsprevious contract.
type IagglayergatewayeventspreviousUpdateDefaultAggchainVKey struct {
	Selector     [4]byte
	PreviousVKey [32]byte
	NewVKey      [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateDefaultAggchainVKey is a free log retrieval operation binding the contract event 0x477ac863279848383fdb5d6bb8957cc8f6f10c0200379e075df3f789d14248e6.
//
// Solidity: event UpdateDefaultAggchainVKey(bytes4 selector, bytes32 previousVKey, bytes32 newVKey)
func (_Iagglayergatewayeventsprevious *IagglayergatewayeventspreviousFilterer) FilterUpdateDefaultAggchainVKey(opts *bind.FilterOpts) (*IagglayergatewayeventspreviousUpdateDefaultAggchainVKeyIterator, error) {

	logs, sub, err := _Iagglayergatewayeventsprevious.contract.FilterLogs(opts, "UpdateDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayeventspreviousUpdateDefaultAggchainVKeyIterator{contract: _Iagglayergatewayeventsprevious.contract, event: "UpdateDefaultAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchUpdateDefaultAggchainVKey is a free log subscription operation binding the contract event 0x477ac863279848383fdb5d6bb8957cc8f6f10c0200379e075df3f789d14248e6.
//
// Solidity: event UpdateDefaultAggchainVKey(bytes4 selector, bytes32 previousVKey, bytes32 newVKey)
func (_Iagglayergatewayeventsprevious *IagglayergatewayeventspreviousFilterer) WatchUpdateDefaultAggchainVKey(opts *bind.WatchOpts, sink chan<- *IagglayergatewayeventspreviousUpdateDefaultAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Iagglayergatewayeventsprevious.contract.WatchLogs(opts, "UpdateDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IagglayergatewayeventspreviousUpdateDefaultAggchainVKey)
				if err := _Iagglayergatewayeventsprevious.contract.UnpackLog(event, "UpdateDefaultAggchainVKey", log); err != nil {
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
func (_Iagglayergatewayeventsprevious *IagglayergatewayeventspreviousFilterer) ParseUpdateDefaultAggchainVKey(log types.Log) (*IagglayergatewayeventspreviousUpdateDefaultAggchainVKey, error) {
	event := new(IagglayergatewayeventspreviousUpdateDefaultAggchainVKey)
	if err := _Iagglayergatewayeventsprevious.contract.UnpackLog(event, "UpdateDefaultAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

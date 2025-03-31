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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVKeyManager\",\"type\":\"address\"}],\"name\":\"AcceptVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"AddAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVKeyManager\",\"type\":\"address\"}],\"name\":\"TransferVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"UpdateAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"useDefaultGateway\",\"type\":\"bool\"}],\"name\":\"UpdateUseDefaultGatewayFlag\",\"type\":\"event\"}]",
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
	NewVKeyManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAcceptVKeyManagerRole is a free log retrieval operation binding the contract event 0x75ccbacededf0b9fc9371bb4cf651be4f96efab2c519a529cfd90e14c9d49ad2.
//
// Solidity: event AcceptVKeyManagerRole(address newVKeyManager)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) FilterAcceptVKeyManagerRole(opts *bind.FilterOpts) (*IaggchainbaseeventsAcceptVKeyManagerRoleIterator, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.FilterLogs(opts, "AcceptVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventsAcceptVKeyManagerRoleIterator{contract: _Iaggchainbaseevents.contract, event: "AcceptVKeyManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptVKeyManagerRole is a free log subscription operation binding the contract event 0x75ccbacededf0b9fc9371bb4cf651be4f96efab2c519a529cfd90e14c9d49ad2.
//
// Solidity: event AcceptVKeyManagerRole(address newVKeyManager)
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

// ParseAcceptVKeyManagerRole is a log parse operation binding the contract event 0x75ccbacededf0b9fc9371bb4cf651be4f96efab2c519a529cfd90e14c9d49ad2.
//
// Solidity: event AcceptVKeyManagerRole(address newVKeyManager)
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
	NewVKeyManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTransferVKeyManagerRole is a free log retrieval operation binding the contract event 0xfa4ea207951b027c94aa15f12ca5d0f4ca0543beb8426209ac79b49d2fc46b10.
//
// Solidity: event TransferVKeyManagerRole(address newVKeyManager)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) FilterTransferVKeyManagerRole(opts *bind.FilterOpts) (*IaggchainbaseeventsTransferVKeyManagerRoleIterator, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.FilterLogs(opts, "TransferVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventsTransferVKeyManagerRoleIterator{contract: _Iaggchainbaseevents.contract, event: "TransferVKeyManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferVKeyManagerRole is a free log subscription operation binding the contract event 0xfa4ea207951b027c94aa15f12ca5d0f4ca0543beb8426209ac79b49d2fc46b10.
//
// Solidity: event TransferVKeyManagerRole(address newVKeyManager)
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

// ParseTransferVKeyManagerRole is a log parse operation binding the contract event 0xfa4ea207951b027c94aa15f12ca5d0f4ca0543beb8426209ac79b49d2fc46b10.
//
// Solidity: event TransferVKeyManagerRole(address newVKeyManager)
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
	Selector        [4]byte
	NewAggchainVKey [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateAggchainVKey is a free log retrieval operation binding the contract event 0x2ae81c794f274593f3cd22db9ad3af161ff9d93d15922b35629a6ffd9be7243f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) FilterUpdateAggchainVKey(opts *bind.FilterOpts) (*IaggchainbaseeventsUpdateAggchainVKeyIterator, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.FilterLogs(opts, "UpdateAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventsUpdateAggchainVKeyIterator{contract: _Iaggchainbaseevents.contract, event: "UpdateAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchUpdateAggchainVKey is a free log subscription operation binding the contract event 0x2ae81c794f274593f3cd22db9ad3af161ff9d93d15922b35629a6ffd9be7243f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
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

// ParseUpdateAggchainVKey is a log parse operation binding the contract event 0x2ae81c794f274593f3cd22db9ad3af161ff9d93d15922b35629a6ffd9be7243f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) ParseUpdateAggchainVKey(log types.Log) (*IaggchainbaseeventsUpdateAggchainVKey, error) {
	event := new(IaggchainbaseeventsUpdateAggchainVKey)
	if err := _Iaggchainbaseevents.contract.UnpackLog(event, "UpdateAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseeventsUpdateUseDefaultGatewayFlagIterator is returned from FilterUpdateUseDefaultGatewayFlag and is used to iterate over the raw logs and unpacked data for UpdateUseDefaultGatewayFlag events raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsUpdateUseDefaultGatewayFlagIterator struct {
	Event *IaggchainbaseeventsUpdateUseDefaultGatewayFlag // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseeventsUpdateUseDefaultGatewayFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseeventsUpdateUseDefaultGatewayFlag)
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
		it.Event = new(IaggchainbaseeventsUpdateUseDefaultGatewayFlag)
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
func (it *IaggchainbaseeventsUpdateUseDefaultGatewayFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseeventsUpdateUseDefaultGatewayFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseeventsUpdateUseDefaultGatewayFlag represents a UpdateUseDefaultGatewayFlag event raised by the Iaggchainbaseevents contract.
type IaggchainbaseeventsUpdateUseDefaultGatewayFlag struct {
	UseDefaultGateway bool
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdateUseDefaultGatewayFlag is a free log retrieval operation binding the contract event 0xd4305f3587ac0e4b2b8b4e5b2072599efd931f81aba626cdbef4b040078b5070.
//
// Solidity: event UpdateUseDefaultGatewayFlag(bool useDefaultGateway)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) FilterUpdateUseDefaultGatewayFlag(opts *bind.FilterOpts) (*IaggchainbaseeventsUpdateUseDefaultGatewayFlagIterator, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.FilterLogs(opts, "UpdateUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseeventsUpdateUseDefaultGatewayFlagIterator{contract: _Iaggchainbaseevents.contract, event: "UpdateUseDefaultGatewayFlag", logs: logs, sub: sub}, nil
}

// WatchUpdateUseDefaultGatewayFlag is a free log subscription operation binding the contract event 0xd4305f3587ac0e4b2b8b4e5b2072599efd931f81aba626cdbef4b040078b5070.
//
// Solidity: event UpdateUseDefaultGatewayFlag(bool useDefaultGateway)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) WatchUpdateUseDefaultGatewayFlag(opts *bind.WatchOpts, sink chan<- *IaggchainbaseeventsUpdateUseDefaultGatewayFlag) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbaseevents.contract.WatchLogs(opts, "UpdateUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseeventsUpdateUseDefaultGatewayFlag)
				if err := _Iaggchainbaseevents.contract.UnpackLog(event, "UpdateUseDefaultGatewayFlag", log); err != nil {
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

// ParseUpdateUseDefaultGatewayFlag is a log parse operation binding the contract event 0xd4305f3587ac0e4b2b8b4e5b2072599efd931f81aba626cdbef4b040078b5070.
//
// Solidity: event UpdateUseDefaultGatewayFlag(bool useDefaultGateway)
func (_Iaggchainbaseevents *IaggchainbaseeventsFilterer) ParseUpdateUseDefaultGatewayFlag(log types.Log) (*IaggchainbaseeventsUpdateUseDefaultGatewayFlag, error) {
	event := new(IaggchainbaseeventsUpdateUseDefaultGatewayFlag)
	if err := _Iaggchainbaseevents.contract.UnpackLog(event, "UpdateUseDefaultGatewayFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iemergencymanager

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

// IemergencymanagerMetaData contains all meta data concerning the Iemergencymanager contract.
var IemergencymanagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IemergencymanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IemergencymanagerMetaData.ABI instead.
var IemergencymanagerABI = IemergencymanagerMetaData.ABI

// Iemergencymanager is an auto generated Go binding around an Ethereum contract.
type Iemergencymanager struct {
	IemergencymanagerCaller     // Read-only binding to the contract
	IemergencymanagerTransactor // Write-only binding to the contract
	IemergencymanagerFilterer   // Log filterer for contract events
}

// IemergencymanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IemergencymanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IemergencymanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IemergencymanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IemergencymanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IemergencymanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IemergencymanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IemergencymanagerSession struct {
	Contract     *Iemergencymanager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IemergencymanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IemergencymanagerCallerSession struct {
	Contract *IemergencymanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IemergencymanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IemergencymanagerTransactorSession struct {
	Contract     *IemergencymanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IemergencymanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IemergencymanagerRaw struct {
	Contract *Iemergencymanager // Generic contract binding to access the raw methods on
}

// IemergencymanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IemergencymanagerCallerRaw struct {
	Contract *IemergencymanagerCaller // Generic read-only contract binding to access the raw methods on
}

// IemergencymanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IemergencymanagerTransactorRaw struct {
	Contract *IemergencymanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIemergencymanager creates a new instance of Iemergencymanager, bound to a specific deployed contract.
func NewIemergencymanager(address common.Address, backend bind.ContractBackend) (*Iemergencymanager, error) {
	contract, err := bindIemergencymanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iemergencymanager{IemergencymanagerCaller: IemergencymanagerCaller{contract: contract}, IemergencymanagerTransactor: IemergencymanagerTransactor{contract: contract}, IemergencymanagerFilterer: IemergencymanagerFilterer{contract: contract}}, nil
}

// NewIemergencymanagerCaller creates a new read-only instance of Iemergencymanager, bound to a specific deployed contract.
func NewIemergencymanagerCaller(address common.Address, caller bind.ContractCaller) (*IemergencymanagerCaller, error) {
	contract, err := bindIemergencymanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IemergencymanagerCaller{contract: contract}, nil
}

// NewIemergencymanagerTransactor creates a new write-only instance of Iemergencymanager, bound to a specific deployed contract.
func NewIemergencymanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IemergencymanagerTransactor, error) {
	contract, err := bindIemergencymanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IemergencymanagerTransactor{contract: contract}, nil
}

// NewIemergencymanagerFilterer creates a new log filterer instance of Iemergencymanager, bound to a specific deployed contract.
func NewIemergencymanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IemergencymanagerFilterer, error) {
	contract, err := bindIemergencymanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IemergencymanagerFilterer{contract: contract}, nil
}

// bindIemergencymanager binds a generic wrapper to an already deployed contract.
func bindIemergencymanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IemergencymanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iemergencymanager *IemergencymanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iemergencymanager.Contract.IemergencymanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iemergencymanager *IemergencymanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iemergencymanager.Contract.IemergencymanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iemergencymanager *IemergencymanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iemergencymanager.Contract.IemergencymanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iemergencymanager *IemergencymanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iemergencymanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iemergencymanager *IemergencymanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iemergencymanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iemergencymanager *IemergencymanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iemergencymanager.Contract.contract.Transact(opts, method, params...)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Iemergencymanager *IemergencymanagerCaller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Iemergencymanager.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Iemergencymanager *IemergencymanagerSession) IsEmergencyState() (bool, error) {
	return _Iemergencymanager.Contract.IsEmergencyState(&_Iemergencymanager.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Iemergencymanager *IemergencymanagerCallerSession) IsEmergencyState() (bool, error) {
	return _Iemergencymanager.Contract.IsEmergencyState(&_Iemergencymanager.CallOpts)
}

// IemergencymanagerEmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the Iemergencymanager contract.
type IemergencymanagerEmergencyStateActivatedIterator struct {
	Event *IemergencymanagerEmergencyStateActivated // Event containing the contract specifics and raw log

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
func (it *IemergencymanagerEmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IemergencymanagerEmergencyStateActivated)
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
		it.Event = new(IemergencymanagerEmergencyStateActivated)
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
func (it *IemergencymanagerEmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IemergencymanagerEmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IemergencymanagerEmergencyStateActivated represents a EmergencyStateActivated event raised by the Iemergencymanager contract.
type IemergencymanagerEmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Iemergencymanager *IemergencymanagerFilterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*IemergencymanagerEmergencyStateActivatedIterator, error) {

	logs, sub, err := _Iemergencymanager.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &IemergencymanagerEmergencyStateActivatedIterator{contract: _Iemergencymanager.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Iemergencymanager *IemergencymanagerFilterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *IemergencymanagerEmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _Iemergencymanager.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IemergencymanagerEmergencyStateActivated)
				if err := _Iemergencymanager.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
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

// ParseEmergencyStateActivated is a log parse operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Iemergencymanager *IemergencymanagerFilterer) ParseEmergencyStateActivated(log types.Log) (*IemergencymanagerEmergencyStateActivated, error) {
	event := new(IemergencymanagerEmergencyStateActivated)
	if err := _Iemergencymanager.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IemergencymanagerEmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the Iemergencymanager contract.
type IemergencymanagerEmergencyStateDeactivatedIterator struct {
	Event *IemergencymanagerEmergencyStateDeactivated // Event containing the contract specifics and raw log

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
func (it *IemergencymanagerEmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IemergencymanagerEmergencyStateDeactivated)
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
		it.Event = new(IemergencymanagerEmergencyStateDeactivated)
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
func (it *IemergencymanagerEmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IemergencymanagerEmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IemergencymanagerEmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the Iemergencymanager contract.
type IemergencymanagerEmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Iemergencymanager *IemergencymanagerFilterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*IemergencymanagerEmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _Iemergencymanager.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &IemergencymanagerEmergencyStateDeactivatedIterator{contract: _Iemergencymanager.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Iemergencymanager *IemergencymanagerFilterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *IemergencymanagerEmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _Iemergencymanager.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IemergencymanagerEmergencyStateDeactivated)
				if err := _Iemergencymanager.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
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

// ParseEmergencyStateDeactivated is a log parse operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Iemergencymanager *IemergencymanagerFilterer) ParseEmergencyStateDeactivated(log types.Log) (*IemergencymanagerEmergencyStateDeactivated, error) {
	event := new(IemergencymanagerEmergencyStateDeactivated)
	if err := _Iemergencymanager.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

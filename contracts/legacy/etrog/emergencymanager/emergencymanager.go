// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package emergencymanager

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

// EmergencymanagerMetaData contains all meta data concerning the Emergencymanager contract.
var EmergencymanagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b5060808061001b5f395ff3fe6080604052348015600e575f80fd5b50600436106026575f3560e01c806315064c9614602a575b5f80fd5b600a5460369060ff1681565b604051901515815260200160405180910390f3fea26469706673582212204734603d06470aa99bf4792361223d1649c5eead548fda9bef178c61a7eee0ba64736f6c63430008140033",
}

// EmergencymanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use EmergencymanagerMetaData.ABI instead.
var EmergencymanagerABI = EmergencymanagerMetaData.ABI

// EmergencymanagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EmergencymanagerMetaData.Bin instead.
var EmergencymanagerBin = EmergencymanagerMetaData.Bin

// DeployEmergencymanager deploys a new Ethereum contract, binding an instance of Emergencymanager to it.
func DeployEmergencymanager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Emergencymanager, error) {
	parsed, err := EmergencymanagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EmergencymanagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Emergencymanager{EmergencymanagerCaller: EmergencymanagerCaller{contract: contract}, EmergencymanagerTransactor: EmergencymanagerTransactor{contract: contract}, EmergencymanagerFilterer: EmergencymanagerFilterer{contract: contract}}, nil
}

// Emergencymanager is an auto generated Go binding around an Ethereum contract.
type Emergencymanager struct {
	EmergencymanagerCaller     // Read-only binding to the contract
	EmergencymanagerTransactor // Write-only binding to the contract
	EmergencymanagerFilterer   // Log filterer for contract events
}

// EmergencymanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type EmergencymanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmergencymanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EmergencymanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmergencymanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EmergencymanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmergencymanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EmergencymanagerSession struct {
	Contract     *Emergencymanager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EmergencymanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EmergencymanagerCallerSession struct {
	Contract *EmergencymanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// EmergencymanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EmergencymanagerTransactorSession struct {
	Contract     *EmergencymanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// EmergencymanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type EmergencymanagerRaw struct {
	Contract *Emergencymanager // Generic contract binding to access the raw methods on
}

// EmergencymanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EmergencymanagerCallerRaw struct {
	Contract *EmergencymanagerCaller // Generic read-only contract binding to access the raw methods on
}

// EmergencymanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EmergencymanagerTransactorRaw struct {
	Contract *EmergencymanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEmergencymanager creates a new instance of Emergencymanager, bound to a specific deployed contract.
func NewEmergencymanager(address common.Address, backend bind.ContractBackend) (*Emergencymanager, error) {
	contract, err := bindEmergencymanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Emergencymanager{EmergencymanagerCaller: EmergencymanagerCaller{contract: contract}, EmergencymanagerTransactor: EmergencymanagerTransactor{contract: contract}, EmergencymanagerFilterer: EmergencymanagerFilterer{contract: contract}}, nil
}

// NewEmergencymanagerCaller creates a new read-only instance of Emergencymanager, bound to a specific deployed contract.
func NewEmergencymanagerCaller(address common.Address, caller bind.ContractCaller) (*EmergencymanagerCaller, error) {
	contract, err := bindEmergencymanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EmergencymanagerCaller{contract: contract}, nil
}

// NewEmergencymanagerTransactor creates a new write-only instance of Emergencymanager, bound to a specific deployed contract.
func NewEmergencymanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*EmergencymanagerTransactor, error) {
	contract, err := bindEmergencymanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EmergencymanagerTransactor{contract: contract}, nil
}

// NewEmergencymanagerFilterer creates a new log filterer instance of Emergencymanager, bound to a specific deployed contract.
func NewEmergencymanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*EmergencymanagerFilterer, error) {
	contract, err := bindEmergencymanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EmergencymanagerFilterer{contract: contract}, nil
}

// bindEmergencymanager binds a generic wrapper to an already deployed contract.
func bindEmergencymanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EmergencymanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Emergencymanager *EmergencymanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Emergencymanager.Contract.EmergencymanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Emergencymanager *EmergencymanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Emergencymanager.Contract.EmergencymanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Emergencymanager *EmergencymanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Emergencymanager.Contract.EmergencymanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Emergencymanager *EmergencymanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Emergencymanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Emergencymanager *EmergencymanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Emergencymanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Emergencymanager *EmergencymanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Emergencymanager.Contract.contract.Transact(opts, method, params...)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Emergencymanager *EmergencymanagerCaller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Emergencymanager.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Emergencymanager *EmergencymanagerSession) IsEmergencyState() (bool, error) {
	return _Emergencymanager.Contract.IsEmergencyState(&_Emergencymanager.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Emergencymanager *EmergencymanagerCallerSession) IsEmergencyState() (bool, error) {
	return _Emergencymanager.Contract.IsEmergencyState(&_Emergencymanager.CallOpts)
}

// EmergencymanagerEmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the Emergencymanager contract.
type EmergencymanagerEmergencyStateActivatedIterator struct {
	Event *EmergencymanagerEmergencyStateActivated // Event containing the contract specifics and raw log

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
func (it *EmergencymanagerEmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmergencymanagerEmergencyStateActivated)
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
		it.Event = new(EmergencymanagerEmergencyStateActivated)
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
func (it *EmergencymanagerEmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmergencymanagerEmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmergencymanagerEmergencyStateActivated represents a EmergencyStateActivated event raised by the Emergencymanager contract.
type EmergencymanagerEmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Emergencymanager *EmergencymanagerFilterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*EmergencymanagerEmergencyStateActivatedIterator, error) {

	logs, sub, err := _Emergencymanager.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &EmergencymanagerEmergencyStateActivatedIterator{contract: _Emergencymanager.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Emergencymanager *EmergencymanagerFilterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *EmergencymanagerEmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _Emergencymanager.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmergencymanagerEmergencyStateActivated)
				if err := _Emergencymanager.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
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
func (_Emergencymanager *EmergencymanagerFilterer) ParseEmergencyStateActivated(log types.Log) (*EmergencymanagerEmergencyStateActivated, error) {
	event := new(EmergencymanagerEmergencyStateActivated)
	if err := _Emergencymanager.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmergencymanagerEmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the Emergencymanager contract.
type EmergencymanagerEmergencyStateDeactivatedIterator struct {
	Event *EmergencymanagerEmergencyStateDeactivated // Event containing the contract specifics and raw log

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
func (it *EmergencymanagerEmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmergencymanagerEmergencyStateDeactivated)
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
		it.Event = new(EmergencymanagerEmergencyStateDeactivated)
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
func (it *EmergencymanagerEmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmergencymanagerEmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmergencymanagerEmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the Emergencymanager contract.
type EmergencymanagerEmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Emergencymanager *EmergencymanagerFilterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*EmergencymanagerEmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _Emergencymanager.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &EmergencymanagerEmergencyStateDeactivatedIterator{contract: _Emergencymanager.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Emergencymanager *EmergencymanagerFilterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *EmergencymanagerEmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _Emergencymanager.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmergencymanagerEmergencyStateDeactivated)
				if err := _Emergencymanager.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
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
func (_Emergencymanager *EmergencymanagerFilterer) ParseEmergencyStateDeactivated(log types.Log) (*EmergencymanagerEmergencyStateDeactivated, error) {
	event := new(EmergencymanagerEmergencyStateDeactivated)
	if err := _Emergencymanager.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

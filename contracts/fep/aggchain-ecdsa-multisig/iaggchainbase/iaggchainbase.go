// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iaggchainbase

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

// IaggchainbaseMetaData contains all meta data concerning the Iaggchainbase contract.
var IaggchainbaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AggchainManagerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainSignersHashNotInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainSignersTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyAggchainSignersArray\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndicesNotInDescendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggLayerGatewayAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggchainDataLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggchainType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAggchainManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAggchainManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingVKeyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyVKeyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerURLCannotBeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdTooHighAfterRemoval\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultSignersAlreadyDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultSignersAlreadyEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultVkeysAlreadyDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultVkeysAlreadyEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroValueAggchainVKey\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAggchainManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"AcceptAggchainManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldVKeyManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVKeyManager\",\"type\":\"address\"}],\"name\":\"AcceptVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"AddAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DisableUseDefaultSignersFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DisableUseDefaultVkeysFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EnableUseDefaultSignersFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EnableUseDefaultVkeysFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"aggchainSigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainSignersHash\",\"type\":\"bytes32\"}],\"name\":\"SignersAndThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentAggchainManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAggchainManager\",\"type\":\"address\"}],\"name\":\"TransferAggchainManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentVKeyManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingVKeyManager\",\"type\":\"address\"}],\"name\":\"TransferVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousAggchainVKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"UpdateAggchainVKey\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AGGCHAIN_TYPE\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"getAggchainHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"initAggchainManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"onVerifyPessimistic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IaggchainbaseABI is the input ABI used to generate the binding from.
// Deprecated: Use IaggchainbaseMetaData.ABI instead.
var IaggchainbaseABI = IaggchainbaseMetaData.ABI

// Iaggchainbase is an auto generated Go binding around an Ethereum contract.
type Iaggchainbase struct {
	IaggchainbaseCaller     // Read-only binding to the contract
	IaggchainbaseTransactor // Write-only binding to the contract
	IaggchainbaseFilterer   // Log filterer for contract events
}

// IaggchainbaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type IaggchainbaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggchainbaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IaggchainbaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggchainbaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IaggchainbaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggchainbaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IaggchainbaseSession struct {
	Contract     *Iaggchainbase    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IaggchainbaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IaggchainbaseCallerSession struct {
	Contract *IaggchainbaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IaggchainbaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IaggchainbaseTransactorSession struct {
	Contract     *IaggchainbaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IaggchainbaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type IaggchainbaseRaw struct {
	Contract *Iaggchainbase // Generic contract binding to access the raw methods on
}

// IaggchainbaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IaggchainbaseCallerRaw struct {
	Contract *IaggchainbaseCaller // Generic read-only contract binding to access the raw methods on
}

// IaggchainbaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IaggchainbaseTransactorRaw struct {
	Contract *IaggchainbaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIaggchainbase creates a new instance of Iaggchainbase, bound to a specific deployed contract.
func NewIaggchainbase(address common.Address, backend bind.ContractBackend) (*Iaggchainbase, error) {
	contract, err := bindIaggchainbase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iaggchainbase{IaggchainbaseCaller: IaggchainbaseCaller{contract: contract}, IaggchainbaseTransactor: IaggchainbaseTransactor{contract: contract}, IaggchainbaseFilterer: IaggchainbaseFilterer{contract: contract}}, nil
}

// NewIaggchainbaseCaller creates a new read-only instance of Iaggchainbase, bound to a specific deployed contract.
func NewIaggchainbaseCaller(address common.Address, caller bind.ContractCaller) (*IaggchainbaseCaller, error) {
	contract, err := bindIaggchainbase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseCaller{contract: contract}, nil
}

// NewIaggchainbaseTransactor creates a new write-only instance of Iaggchainbase, bound to a specific deployed contract.
func NewIaggchainbaseTransactor(address common.Address, transactor bind.ContractTransactor) (*IaggchainbaseTransactor, error) {
	contract, err := bindIaggchainbase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseTransactor{contract: contract}, nil
}

// NewIaggchainbaseFilterer creates a new log filterer instance of Iaggchainbase, bound to a specific deployed contract.
func NewIaggchainbaseFilterer(address common.Address, filterer bind.ContractFilterer) (*IaggchainbaseFilterer, error) {
	contract, err := bindIaggchainbase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseFilterer{contract: contract}, nil
}

// bindIaggchainbase binds a generic wrapper to an already deployed contract.
func bindIaggchainbase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IaggchainbaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iaggchainbase *IaggchainbaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iaggchainbase.Contract.IaggchainbaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iaggchainbase *IaggchainbaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iaggchainbase.Contract.IaggchainbaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iaggchainbase *IaggchainbaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iaggchainbase.Contract.IaggchainbaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iaggchainbase *IaggchainbaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iaggchainbase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iaggchainbase *IaggchainbaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iaggchainbase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iaggchainbase *IaggchainbaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iaggchainbase.Contract.contract.Transact(opts, method, params...)
}

// AGGCHAINTYPE is a free data retrieval call binding the contract method 0x6e7fbce9.
//
// Solidity: function AGGCHAIN_TYPE() view returns(bytes2)
func (_Iaggchainbase *IaggchainbaseCaller) AGGCHAINTYPE(opts *bind.CallOpts) ([2]byte, error) {
	var out []interface{}
	err := _Iaggchainbase.contract.Call(opts, &out, "AGGCHAIN_TYPE")

	if err != nil {
		return *new([2]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([2]byte)).(*[2]byte)

	return out0, err

}

// AGGCHAINTYPE is a free data retrieval call binding the contract method 0x6e7fbce9.
//
// Solidity: function AGGCHAIN_TYPE() view returns(bytes2)
func (_Iaggchainbase *IaggchainbaseSession) AGGCHAINTYPE() ([2]byte, error) {
	return _Iaggchainbase.Contract.AGGCHAINTYPE(&_Iaggchainbase.CallOpts)
}

// AGGCHAINTYPE is a free data retrieval call binding the contract method 0x6e7fbce9.
//
// Solidity: function AGGCHAIN_TYPE() view returns(bytes2)
func (_Iaggchainbase *IaggchainbaseCallerSession) AGGCHAINTYPE() ([2]byte, error) {
	return _Iaggchainbase.Contract.AGGCHAINTYPE(&_Iaggchainbase.CallOpts)
}

// GetAggchainHash is a free data retrieval call binding the contract method 0x6a55f66c.
//
// Solidity: function getAggchainHash(bytes aggchainData) view returns(bytes32)
func (_Iaggchainbase *IaggchainbaseCaller) GetAggchainHash(opts *bind.CallOpts, aggchainData []byte) ([32]byte, error) {
	var out []interface{}
	err := _Iaggchainbase.contract.Call(opts, &out, "getAggchainHash", aggchainData)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAggchainHash is a free data retrieval call binding the contract method 0x6a55f66c.
//
// Solidity: function getAggchainHash(bytes aggchainData) view returns(bytes32)
func (_Iaggchainbase *IaggchainbaseSession) GetAggchainHash(aggchainData []byte) ([32]byte, error) {
	return _Iaggchainbase.Contract.GetAggchainHash(&_Iaggchainbase.CallOpts, aggchainData)
}

// GetAggchainHash is a free data retrieval call binding the contract method 0x6a55f66c.
//
// Solidity: function getAggchainHash(bytes aggchainData) view returns(bytes32)
func (_Iaggchainbase *IaggchainbaseCallerSession) GetAggchainHash(aggchainData []byte) ([32]byte, error) {
	return _Iaggchainbase.Contract.GetAggchainHash(&_Iaggchainbase.CallOpts, aggchainData)
}

// InitAggchainManager is a paid mutator transaction binding the contract method 0xb3a326f7.
//
// Solidity: function initAggchainManager(address newAggchainManager) returns()
func (_Iaggchainbase *IaggchainbaseTransactor) InitAggchainManager(opts *bind.TransactOpts, newAggchainManager common.Address) (*types.Transaction, error) {
	return _Iaggchainbase.contract.Transact(opts, "initAggchainManager", newAggchainManager)
}

// InitAggchainManager is a paid mutator transaction binding the contract method 0xb3a326f7.
//
// Solidity: function initAggchainManager(address newAggchainManager) returns()
func (_Iaggchainbase *IaggchainbaseSession) InitAggchainManager(newAggchainManager common.Address) (*types.Transaction, error) {
	return _Iaggchainbase.Contract.InitAggchainManager(&_Iaggchainbase.TransactOpts, newAggchainManager)
}

// InitAggchainManager is a paid mutator transaction binding the contract method 0xb3a326f7.
//
// Solidity: function initAggchainManager(address newAggchainManager) returns()
func (_Iaggchainbase *IaggchainbaseTransactorSession) InitAggchainManager(newAggchainManager common.Address) (*types.Transaction, error) {
	return _Iaggchainbase.Contract.InitAggchainManager(&_Iaggchainbase.TransactOpts, newAggchainManager)
}

// OnVerifyPessimistic is a paid mutator transaction binding the contract method 0x9ee4afa3.
//
// Solidity: function onVerifyPessimistic(bytes aggchainData) returns()
func (_Iaggchainbase *IaggchainbaseTransactor) OnVerifyPessimistic(opts *bind.TransactOpts, aggchainData []byte) (*types.Transaction, error) {
	return _Iaggchainbase.contract.Transact(opts, "onVerifyPessimistic", aggchainData)
}

// OnVerifyPessimistic is a paid mutator transaction binding the contract method 0x9ee4afa3.
//
// Solidity: function onVerifyPessimistic(bytes aggchainData) returns()
func (_Iaggchainbase *IaggchainbaseSession) OnVerifyPessimistic(aggchainData []byte) (*types.Transaction, error) {
	return _Iaggchainbase.Contract.OnVerifyPessimistic(&_Iaggchainbase.TransactOpts, aggchainData)
}

// OnVerifyPessimistic is a paid mutator transaction binding the contract method 0x9ee4afa3.
//
// Solidity: function onVerifyPessimistic(bytes aggchainData) returns()
func (_Iaggchainbase *IaggchainbaseTransactorSession) OnVerifyPessimistic(aggchainData []byte) (*types.Transaction, error) {
	return _Iaggchainbase.Contract.OnVerifyPessimistic(&_Iaggchainbase.TransactOpts, aggchainData)
}

// IaggchainbaseAcceptAggchainManagerRoleIterator is returned from FilterAcceptAggchainManagerRole and is used to iterate over the raw logs and unpacked data for AcceptAggchainManagerRole events raised by the Iaggchainbase contract.
type IaggchainbaseAcceptAggchainManagerRoleIterator struct {
	Event *IaggchainbaseAcceptAggchainManagerRole // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseAcceptAggchainManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseAcceptAggchainManagerRole)
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
		it.Event = new(IaggchainbaseAcceptAggchainManagerRole)
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
func (it *IaggchainbaseAcceptAggchainManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseAcceptAggchainManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseAcceptAggchainManagerRole represents a AcceptAggchainManagerRole event raised by the Iaggchainbase contract.
type IaggchainbaseAcceptAggchainManagerRole struct {
	OldAggchainManager common.Address
	NewAggchainManager common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAcceptAggchainManagerRole is a free log retrieval operation binding the contract event 0x67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f.
//
// Solidity: event AcceptAggchainManagerRole(address oldAggchainManager, address newAggchainManager)
func (_Iaggchainbase *IaggchainbaseFilterer) FilterAcceptAggchainManagerRole(opts *bind.FilterOpts) (*IaggchainbaseAcceptAggchainManagerRoleIterator, error) {

	logs, sub, err := _Iaggchainbase.contract.FilterLogs(opts, "AcceptAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseAcceptAggchainManagerRoleIterator{contract: _Iaggchainbase.contract, event: "AcceptAggchainManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAggchainManagerRole is a free log subscription operation binding the contract event 0x67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f.
//
// Solidity: event AcceptAggchainManagerRole(address oldAggchainManager, address newAggchainManager)
func (_Iaggchainbase *IaggchainbaseFilterer) WatchAcceptAggchainManagerRole(opts *bind.WatchOpts, sink chan<- *IaggchainbaseAcceptAggchainManagerRole) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbase.contract.WatchLogs(opts, "AcceptAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseAcceptAggchainManagerRole)
				if err := _Iaggchainbase.contract.UnpackLog(event, "AcceptAggchainManagerRole", log); err != nil {
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
func (_Iaggchainbase *IaggchainbaseFilterer) ParseAcceptAggchainManagerRole(log types.Log) (*IaggchainbaseAcceptAggchainManagerRole, error) {
	event := new(IaggchainbaseAcceptAggchainManagerRole)
	if err := _Iaggchainbase.contract.UnpackLog(event, "AcceptAggchainManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseAcceptVKeyManagerRoleIterator is returned from FilterAcceptVKeyManagerRole and is used to iterate over the raw logs and unpacked data for AcceptVKeyManagerRole events raised by the Iaggchainbase contract.
type IaggchainbaseAcceptVKeyManagerRoleIterator struct {
	Event *IaggchainbaseAcceptVKeyManagerRole // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseAcceptVKeyManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseAcceptVKeyManagerRole)
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
		it.Event = new(IaggchainbaseAcceptVKeyManagerRole)
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
func (it *IaggchainbaseAcceptVKeyManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseAcceptVKeyManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseAcceptVKeyManagerRole represents a AcceptVKeyManagerRole event raised by the Iaggchainbase contract.
type IaggchainbaseAcceptVKeyManagerRole struct {
	OldVKeyManager common.Address
	NewVKeyManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAcceptVKeyManagerRole is a free log retrieval operation binding the contract event 0xbb3b066505f14a628f4ba4187a046abd4dd17e96591d7a9ed31c91c79322ffe2.
//
// Solidity: event AcceptVKeyManagerRole(address oldVKeyManager, address newVKeyManager)
func (_Iaggchainbase *IaggchainbaseFilterer) FilterAcceptVKeyManagerRole(opts *bind.FilterOpts) (*IaggchainbaseAcceptVKeyManagerRoleIterator, error) {

	logs, sub, err := _Iaggchainbase.contract.FilterLogs(opts, "AcceptVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseAcceptVKeyManagerRoleIterator{contract: _Iaggchainbase.contract, event: "AcceptVKeyManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptVKeyManagerRole is a free log subscription operation binding the contract event 0xbb3b066505f14a628f4ba4187a046abd4dd17e96591d7a9ed31c91c79322ffe2.
//
// Solidity: event AcceptVKeyManagerRole(address oldVKeyManager, address newVKeyManager)
func (_Iaggchainbase *IaggchainbaseFilterer) WatchAcceptVKeyManagerRole(opts *bind.WatchOpts, sink chan<- *IaggchainbaseAcceptVKeyManagerRole) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbase.contract.WatchLogs(opts, "AcceptVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseAcceptVKeyManagerRole)
				if err := _Iaggchainbase.contract.UnpackLog(event, "AcceptVKeyManagerRole", log); err != nil {
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
func (_Iaggchainbase *IaggchainbaseFilterer) ParseAcceptVKeyManagerRole(log types.Log) (*IaggchainbaseAcceptVKeyManagerRole, error) {
	event := new(IaggchainbaseAcceptVKeyManagerRole)
	if err := _Iaggchainbase.contract.UnpackLog(event, "AcceptVKeyManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseAddAggchainVKeyIterator is returned from FilterAddAggchainVKey and is used to iterate over the raw logs and unpacked data for AddAggchainVKey events raised by the Iaggchainbase contract.
type IaggchainbaseAddAggchainVKeyIterator struct {
	Event *IaggchainbaseAddAggchainVKey // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseAddAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseAddAggchainVKey)
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
		it.Event = new(IaggchainbaseAddAggchainVKey)
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
func (it *IaggchainbaseAddAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseAddAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseAddAggchainVKey represents a AddAggchainVKey event raised by the Iaggchainbase contract.
type IaggchainbaseAddAggchainVKey struct {
	Selector        [4]byte
	NewAggchainVKey [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddAggchainVKey is a free log retrieval operation binding the contract event 0x6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a79.
//
// Solidity: event AddAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Iaggchainbase *IaggchainbaseFilterer) FilterAddAggchainVKey(opts *bind.FilterOpts) (*IaggchainbaseAddAggchainVKeyIterator, error) {

	logs, sub, err := _Iaggchainbase.contract.FilterLogs(opts, "AddAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseAddAggchainVKeyIterator{contract: _Iaggchainbase.contract, event: "AddAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchAddAggchainVKey is a free log subscription operation binding the contract event 0x6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a79.
//
// Solidity: event AddAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Iaggchainbase *IaggchainbaseFilterer) WatchAddAggchainVKey(opts *bind.WatchOpts, sink chan<- *IaggchainbaseAddAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbase.contract.WatchLogs(opts, "AddAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseAddAggchainVKey)
				if err := _Iaggchainbase.contract.UnpackLog(event, "AddAggchainVKey", log); err != nil {
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
func (_Iaggchainbase *IaggchainbaseFilterer) ParseAddAggchainVKey(log types.Log) (*IaggchainbaseAddAggchainVKey, error) {
	event := new(IaggchainbaseAddAggchainVKey)
	if err := _Iaggchainbase.contract.UnpackLog(event, "AddAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseDisableUseDefaultSignersFlagIterator is returned from FilterDisableUseDefaultSignersFlag and is used to iterate over the raw logs and unpacked data for DisableUseDefaultSignersFlag events raised by the Iaggchainbase contract.
type IaggchainbaseDisableUseDefaultSignersFlagIterator struct {
	Event *IaggchainbaseDisableUseDefaultSignersFlag // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseDisableUseDefaultSignersFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseDisableUseDefaultSignersFlag)
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
		it.Event = new(IaggchainbaseDisableUseDefaultSignersFlag)
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
func (it *IaggchainbaseDisableUseDefaultSignersFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseDisableUseDefaultSignersFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseDisableUseDefaultSignersFlag represents a DisableUseDefaultSignersFlag event raised by the Iaggchainbase contract.
type IaggchainbaseDisableUseDefaultSignersFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisableUseDefaultSignersFlag is a free log retrieval operation binding the contract event 0x4c75580a56c734245a7418eb07d8a311e1bff79f982fed747da3589630e414be.
//
// Solidity: event DisableUseDefaultSignersFlag()
func (_Iaggchainbase *IaggchainbaseFilterer) FilterDisableUseDefaultSignersFlag(opts *bind.FilterOpts) (*IaggchainbaseDisableUseDefaultSignersFlagIterator, error) {

	logs, sub, err := _Iaggchainbase.contract.FilterLogs(opts, "DisableUseDefaultSignersFlag")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseDisableUseDefaultSignersFlagIterator{contract: _Iaggchainbase.contract, event: "DisableUseDefaultSignersFlag", logs: logs, sub: sub}, nil
}

// WatchDisableUseDefaultSignersFlag is a free log subscription operation binding the contract event 0x4c75580a56c734245a7418eb07d8a311e1bff79f982fed747da3589630e414be.
//
// Solidity: event DisableUseDefaultSignersFlag()
func (_Iaggchainbase *IaggchainbaseFilterer) WatchDisableUseDefaultSignersFlag(opts *bind.WatchOpts, sink chan<- *IaggchainbaseDisableUseDefaultSignersFlag) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbase.contract.WatchLogs(opts, "DisableUseDefaultSignersFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseDisableUseDefaultSignersFlag)
				if err := _Iaggchainbase.contract.UnpackLog(event, "DisableUseDefaultSignersFlag", log); err != nil {
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
func (_Iaggchainbase *IaggchainbaseFilterer) ParseDisableUseDefaultSignersFlag(log types.Log) (*IaggchainbaseDisableUseDefaultSignersFlag, error) {
	event := new(IaggchainbaseDisableUseDefaultSignersFlag)
	if err := _Iaggchainbase.contract.UnpackLog(event, "DisableUseDefaultSignersFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseDisableUseDefaultVkeysFlagIterator is returned from FilterDisableUseDefaultVkeysFlag and is used to iterate over the raw logs and unpacked data for DisableUseDefaultVkeysFlag events raised by the Iaggchainbase contract.
type IaggchainbaseDisableUseDefaultVkeysFlagIterator struct {
	Event *IaggchainbaseDisableUseDefaultVkeysFlag // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseDisableUseDefaultVkeysFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseDisableUseDefaultVkeysFlag)
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
		it.Event = new(IaggchainbaseDisableUseDefaultVkeysFlag)
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
func (it *IaggchainbaseDisableUseDefaultVkeysFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseDisableUseDefaultVkeysFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseDisableUseDefaultVkeysFlag represents a DisableUseDefaultVkeysFlag event raised by the Iaggchainbase contract.
type IaggchainbaseDisableUseDefaultVkeysFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisableUseDefaultVkeysFlag is a free log retrieval operation binding the contract event 0x922aeecd3505b46415820aae489ed9dac9e250e74d497b14c33e8360b581ac07.
//
// Solidity: event DisableUseDefaultVkeysFlag()
func (_Iaggchainbase *IaggchainbaseFilterer) FilterDisableUseDefaultVkeysFlag(opts *bind.FilterOpts) (*IaggchainbaseDisableUseDefaultVkeysFlagIterator, error) {

	logs, sub, err := _Iaggchainbase.contract.FilterLogs(opts, "DisableUseDefaultVkeysFlag")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseDisableUseDefaultVkeysFlagIterator{contract: _Iaggchainbase.contract, event: "DisableUseDefaultVkeysFlag", logs: logs, sub: sub}, nil
}

// WatchDisableUseDefaultVkeysFlag is a free log subscription operation binding the contract event 0x922aeecd3505b46415820aae489ed9dac9e250e74d497b14c33e8360b581ac07.
//
// Solidity: event DisableUseDefaultVkeysFlag()
func (_Iaggchainbase *IaggchainbaseFilterer) WatchDisableUseDefaultVkeysFlag(opts *bind.WatchOpts, sink chan<- *IaggchainbaseDisableUseDefaultVkeysFlag) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbase.contract.WatchLogs(opts, "DisableUseDefaultVkeysFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseDisableUseDefaultVkeysFlag)
				if err := _Iaggchainbase.contract.UnpackLog(event, "DisableUseDefaultVkeysFlag", log); err != nil {
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
func (_Iaggchainbase *IaggchainbaseFilterer) ParseDisableUseDefaultVkeysFlag(log types.Log) (*IaggchainbaseDisableUseDefaultVkeysFlag, error) {
	event := new(IaggchainbaseDisableUseDefaultVkeysFlag)
	if err := _Iaggchainbase.contract.UnpackLog(event, "DisableUseDefaultVkeysFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseEnableUseDefaultSignersFlagIterator is returned from FilterEnableUseDefaultSignersFlag and is used to iterate over the raw logs and unpacked data for EnableUseDefaultSignersFlag events raised by the Iaggchainbase contract.
type IaggchainbaseEnableUseDefaultSignersFlagIterator struct {
	Event *IaggchainbaseEnableUseDefaultSignersFlag // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseEnableUseDefaultSignersFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseEnableUseDefaultSignersFlag)
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
		it.Event = new(IaggchainbaseEnableUseDefaultSignersFlag)
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
func (it *IaggchainbaseEnableUseDefaultSignersFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseEnableUseDefaultSignersFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseEnableUseDefaultSignersFlag represents a EnableUseDefaultSignersFlag event raised by the Iaggchainbase contract.
type IaggchainbaseEnableUseDefaultSignersFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEnableUseDefaultSignersFlag is a free log retrieval operation binding the contract event 0x67ec953bdc8546ede08f8ee91e5205a1d1814e126cb8f5d00a918ddb1eaa292b.
//
// Solidity: event EnableUseDefaultSignersFlag()
func (_Iaggchainbase *IaggchainbaseFilterer) FilterEnableUseDefaultSignersFlag(opts *bind.FilterOpts) (*IaggchainbaseEnableUseDefaultSignersFlagIterator, error) {

	logs, sub, err := _Iaggchainbase.contract.FilterLogs(opts, "EnableUseDefaultSignersFlag")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseEnableUseDefaultSignersFlagIterator{contract: _Iaggchainbase.contract, event: "EnableUseDefaultSignersFlag", logs: logs, sub: sub}, nil
}

// WatchEnableUseDefaultSignersFlag is a free log subscription operation binding the contract event 0x67ec953bdc8546ede08f8ee91e5205a1d1814e126cb8f5d00a918ddb1eaa292b.
//
// Solidity: event EnableUseDefaultSignersFlag()
func (_Iaggchainbase *IaggchainbaseFilterer) WatchEnableUseDefaultSignersFlag(opts *bind.WatchOpts, sink chan<- *IaggchainbaseEnableUseDefaultSignersFlag) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbase.contract.WatchLogs(opts, "EnableUseDefaultSignersFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseEnableUseDefaultSignersFlag)
				if err := _Iaggchainbase.contract.UnpackLog(event, "EnableUseDefaultSignersFlag", log); err != nil {
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
func (_Iaggchainbase *IaggchainbaseFilterer) ParseEnableUseDefaultSignersFlag(log types.Log) (*IaggchainbaseEnableUseDefaultSignersFlag, error) {
	event := new(IaggchainbaseEnableUseDefaultSignersFlag)
	if err := _Iaggchainbase.contract.UnpackLog(event, "EnableUseDefaultSignersFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseEnableUseDefaultVkeysFlagIterator is returned from FilterEnableUseDefaultVkeysFlag and is used to iterate over the raw logs and unpacked data for EnableUseDefaultVkeysFlag events raised by the Iaggchainbase contract.
type IaggchainbaseEnableUseDefaultVkeysFlagIterator struct {
	Event *IaggchainbaseEnableUseDefaultVkeysFlag // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseEnableUseDefaultVkeysFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseEnableUseDefaultVkeysFlag)
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
		it.Event = new(IaggchainbaseEnableUseDefaultVkeysFlag)
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
func (it *IaggchainbaseEnableUseDefaultVkeysFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseEnableUseDefaultVkeysFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseEnableUseDefaultVkeysFlag represents a EnableUseDefaultVkeysFlag event raised by the Iaggchainbase contract.
type IaggchainbaseEnableUseDefaultVkeysFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEnableUseDefaultVkeysFlag is a free log retrieval operation binding the contract event 0xaacf3fb6dd8daa3bebb71f5548f782413c3f3531625c6f9057c0f3d751b83829.
//
// Solidity: event EnableUseDefaultVkeysFlag()
func (_Iaggchainbase *IaggchainbaseFilterer) FilterEnableUseDefaultVkeysFlag(opts *bind.FilterOpts) (*IaggchainbaseEnableUseDefaultVkeysFlagIterator, error) {

	logs, sub, err := _Iaggchainbase.contract.FilterLogs(opts, "EnableUseDefaultVkeysFlag")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseEnableUseDefaultVkeysFlagIterator{contract: _Iaggchainbase.contract, event: "EnableUseDefaultVkeysFlag", logs: logs, sub: sub}, nil
}

// WatchEnableUseDefaultVkeysFlag is a free log subscription operation binding the contract event 0xaacf3fb6dd8daa3bebb71f5548f782413c3f3531625c6f9057c0f3d751b83829.
//
// Solidity: event EnableUseDefaultVkeysFlag()
func (_Iaggchainbase *IaggchainbaseFilterer) WatchEnableUseDefaultVkeysFlag(opts *bind.WatchOpts, sink chan<- *IaggchainbaseEnableUseDefaultVkeysFlag) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbase.contract.WatchLogs(opts, "EnableUseDefaultVkeysFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseEnableUseDefaultVkeysFlag)
				if err := _Iaggchainbase.contract.UnpackLog(event, "EnableUseDefaultVkeysFlag", log); err != nil {
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
func (_Iaggchainbase *IaggchainbaseFilterer) ParseEnableUseDefaultVkeysFlag(log types.Log) (*IaggchainbaseEnableUseDefaultVkeysFlag, error) {
	event := new(IaggchainbaseEnableUseDefaultVkeysFlag)
	if err := _Iaggchainbase.contract.UnpackLog(event, "EnableUseDefaultVkeysFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseSignersAndThresholdUpdatedIterator is returned from FilterSignersAndThresholdUpdated and is used to iterate over the raw logs and unpacked data for SignersAndThresholdUpdated events raised by the Iaggchainbase contract.
type IaggchainbaseSignersAndThresholdUpdatedIterator struct {
	Event *IaggchainbaseSignersAndThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseSignersAndThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseSignersAndThresholdUpdated)
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
		it.Event = new(IaggchainbaseSignersAndThresholdUpdated)
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
func (it *IaggchainbaseSignersAndThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseSignersAndThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseSignersAndThresholdUpdated represents a SignersAndThresholdUpdated event raised by the Iaggchainbase contract.
type IaggchainbaseSignersAndThresholdUpdated struct {
	AggchainSigners        []common.Address
	NewThreshold           *big.Int
	NewAggchainSignersHash [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSignersAndThresholdUpdated is a free log retrieval operation binding the contract event 0x66d7b0647fdd512b69cbf4f8e1ce8068bfe0b236168e2704ba13b07425eaa743.
//
// Solidity: event SignersAndThresholdUpdated(address[] aggchainSigners, uint256 newThreshold, bytes32 newAggchainSignersHash)
func (_Iaggchainbase *IaggchainbaseFilterer) FilterSignersAndThresholdUpdated(opts *bind.FilterOpts) (*IaggchainbaseSignersAndThresholdUpdatedIterator, error) {

	logs, sub, err := _Iaggchainbase.contract.FilterLogs(opts, "SignersAndThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseSignersAndThresholdUpdatedIterator{contract: _Iaggchainbase.contract, event: "SignersAndThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchSignersAndThresholdUpdated is a free log subscription operation binding the contract event 0x66d7b0647fdd512b69cbf4f8e1ce8068bfe0b236168e2704ba13b07425eaa743.
//
// Solidity: event SignersAndThresholdUpdated(address[] aggchainSigners, uint256 newThreshold, bytes32 newAggchainSignersHash)
func (_Iaggchainbase *IaggchainbaseFilterer) WatchSignersAndThresholdUpdated(opts *bind.WatchOpts, sink chan<- *IaggchainbaseSignersAndThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbase.contract.WatchLogs(opts, "SignersAndThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseSignersAndThresholdUpdated)
				if err := _Iaggchainbase.contract.UnpackLog(event, "SignersAndThresholdUpdated", log); err != nil {
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

// ParseSignersAndThresholdUpdated is a log parse operation binding the contract event 0x66d7b0647fdd512b69cbf4f8e1ce8068bfe0b236168e2704ba13b07425eaa743.
//
// Solidity: event SignersAndThresholdUpdated(address[] aggchainSigners, uint256 newThreshold, bytes32 newAggchainSignersHash)
func (_Iaggchainbase *IaggchainbaseFilterer) ParseSignersAndThresholdUpdated(log types.Log) (*IaggchainbaseSignersAndThresholdUpdated, error) {
	event := new(IaggchainbaseSignersAndThresholdUpdated)
	if err := _Iaggchainbase.contract.UnpackLog(event, "SignersAndThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseTransferAggchainManagerRoleIterator is returned from FilterTransferAggchainManagerRole and is used to iterate over the raw logs and unpacked data for TransferAggchainManagerRole events raised by the Iaggchainbase contract.
type IaggchainbaseTransferAggchainManagerRoleIterator struct {
	Event *IaggchainbaseTransferAggchainManagerRole // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseTransferAggchainManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseTransferAggchainManagerRole)
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
		it.Event = new(IaggchainbaseTransferAggchainManagerRole)
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
func (it *IaggchainbaseTransferAggchainManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseTransferAggchainManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseTransferAggchainManagerRole represents a TransferAggchainManagerRole event raised by the Iaggchainbase contract.
type IaggchainbaseTransferAggchainManagerRole struct {
	CurrentAggchainManager    common.Address
	NewPendingAggchainManager common.Address
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterTransferAggchainManagerRole is a free log retrieval operation binding the contract event 0xa3d8e5d045432398be30f83ce7c35a7bfc220c1b66cc5bf3f4dd4d539d93fab6.
//
// Solidity: event TransferAggchainManagerRole(address currentAggchainManager, address newPendingAggchainManager)
func (_Iaggchainbase *IaggchainbaseFilterer) FilterTransferAggchainManagerRole(opts *bind.FilterOpts) (*IaggchainbaseTransferAggchainManagerRoleIterator, error) {

	logs, sub, err := _Iaggchainbase.contract.FilterLogs(opts, "TransferAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseTransferAggchainManagerRoleIterator{contract: _Iaggchainbase.contract, event: "TransferAggchainManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferAggchainManagerRole is a free log subscription operation binding the contract event 0xa3d8e5d045432398be30f83ce7c35a7bfc220c1b66cc5bf3f4dd4d539d93fab6.
//
// Solidity: event TransferAggchainManagerRole(address currentAggchainManager, address newPendingAggchainManager)
func (_Iaggchainbase *IaggchainbaseFilterer) WatchTransferAggchainManagerRole(opts *bind.WatchOpts, sink chan<- *IaggchainbaseTransferAggchainManagerRole) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbase.contract.WatchLogs(opts, "TransferAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseTransferAggchainManagerRole)
				if err := _Iaggchainbase.contract.UnpackLog(event, "TransferAggchainManagerRole", log); err != nil {
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
func (_Iaggchainbase *IaggchainbaseFilterer) ParseTransferAggchainManagerRole(log types.Log) (*IaggchainbaseTransferAggchainManagerRole, error) {
	event := new(IaggchainbaseTransferAggchainManagerRole)
	if err := _Iaggchainbase.contract.UnpackLog(event, "TransferAggchainManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseTransferVKeyManagerRoleIterator is returned from FilterTransferVKeyManagerRole and is used to iterate over the raw logs and unpacked data for TransferVKeyManagerRole events raised by the Iaggchainbase contract.
type IaggchainbaseTransferVKeyManagerRoleIterator struct {
	Event *IaggchainbaseTransferVKeyManagerRole // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseTransferVKeyManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseTransferVKeyManagerRole)
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
		it.Event = new(IaggchainbaseTransferVKeyManagerRole)
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
func (it *IaggchainbaseTransferVKeyManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseTransferVKeyManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseTransferVKeyManagerRole represents a TransferVKeyManagerRole event raised by the Iaggchainbase contract.
type IaggchainbaseTransferVKeyManagerRole struct {
	CurrentVKeyManager    common.Address
	NewPendingVKeyManager common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterTransferVKeyManagerRole is a free log retrieval operation binding the contract event 0xc54ae01017d0b80bd8af833f66387d6eb547dc16c8206faf13d0b72764aab8b2.
//
// Solidity: event TransferVKeyManagerRole(address currentVKeyManager, address newPendingVKeyManager)
func (_Iaggchainbase *IaggchainbaseFilterer) FilterTransferVKeyManagerRole(opts *bind.FilterOpts) (*IaggchainbaseTransferVKeyManagerRoleIterator, error) {

	logs, sub, err := _Iaggchainbase.contract.FilterLogs(opts, "TransferVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseTransferVKeyManagerRoleIterator{contract: _Iaggchainbase.contract, event: "TransferVKeyManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferVKeyManagerRole is a free log subscription operation binding the contract event 0xc54ae01017d0b80bd8af833f66387d6eb547dc16c8206faf13d0b72764aab8b2.
//
// Solidity: event TransferVKeyManagerRole(address currentVKeyManager, address newPendingVKeyManager)
func (_Iaggchainbase *IaggchainbaseFilterer) WatchTransferVKeyManagerRole(opts *bind.WatchOpts, sink chan<- *IaggchainbaseTransferVKeyManagerRole) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbase.contract.WatchLogs(opts, "TransferVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseTransferVKeyManagerRole)
				if err := _Iaggchainbase.contract.UnpackLog(event, "TransferVKeyManagerRole", log); err != nil {
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
func (_Iaggchainbase *IaggchainbaseFilterer) ParseTransferVKeyManagerRole(log types.Log) (*IaggchainbaseTransferVKeyManagerRole, error) {
	event := new(IaggchainbaseTransferVKeyManagerRole)
	if err := _Iaggchainbase.contract.UnpackLog(event, "TransferVKeyManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggchainbaseUpdateAggchainVKeyIterator is returned from FilterUpdateAggchainVKey and is used to iterate over the raw logs and unpacked data for UpdateAggchainVKey events raised by the Iaggchainbase contract.
type IaggchainbaseUpdateAggchainVKeyIterator struct {
	Event *IaggchainbaseUpdateAggchainVKey // Event containing the contract specifics and raw log

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
func (it *IaggchainbaseUpdateAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainbaseUpdateAggchainVKey)
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
		it.Event = new(IaggchainbaseUpdateAggchainVKey)
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
func (it *IaggchainbaseUpdateAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainbaseUpdateAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainbaseUpdateAggchainVKey represents a UpdateAggchainVKey event raised by the Iaggchainbase contract.
type IaggchainbaseUpdateAggchainVKey struct {
	Selector             [4]byte
	PreviousAggchainVKey [32]byte
	NewAggchainVKey      [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateAggchainVKey is a free log retrieval operation binding the contract event 0x0aa5f73c189fb0b0a7cc98ae5fa89dfc16595480396208483518178435ed5b4f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 previousAggchainVKey, bytes32 newAggchainVKey)
func (_Iaggchainbase *IaggchainbaseFilterer) FilterUpdateAggchainVKey(opts *bind.FilterOpts) (*IaggchainbaseUpdateAggchainVKeyIterator, error) {

	logs, sub, err := _Iaggchainbase.contract.FilterLogs(opts, "UpdateAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseUpdateAggchainVKeyIterator{contract: _Iaggchainbase.contract, event: "UpdateAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchUpdateAggchainVKey is a free log subscription operation binding the contract event 0x0aa5f73c189fb0b0a7cc98ae5fa89dfc16595480396208483518178435ed5b4f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 previousAggchainVKey, bytes32 newAggchainVKey)
func (_Iaggchainbase *IaggchainbaseFilterer) WatchUpdateAggchainVKey(opts *bind.WatchOpts, sink chan<- *IaggchainbaseUpdateAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Iaggchainbase.contract.WatchLogs(opts, "UpdateAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainbaseUpdateAggchainVKey)
				if err := _Iaggchainbase.contract.UnpackLog(event, "UpdateAggchainVKey", log); err != nil {
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
func (_Iaggchainbase *IaggchainbaseFilterer) ParseUpdateAggchainVKey(log types.Log) (*IaggchainbaseUpdateAggchainVKey, error) {
	event := new(IaggchainbaseUpdateAggchainVKey)
	if err := _Iaggchainbase.contract.UnpackLog(event, "UpdateAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

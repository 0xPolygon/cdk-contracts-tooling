// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iagglayergatewayprevious

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

// IagglayergatewaypreviousMetaData contains all meta data concerning the Iagglayergatewayprevious contract.
var IagglayergatewaypreviousMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AggchainVKeyAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProofBytesLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAggLayerAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAggLayerAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PPSelectorCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"RouteAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"RouteIsAlreadyFrozen\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"RouteIsFrozen\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"RouteNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VKeyCannotBeZero\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newVKey\",\"type\":\"bytes32\"}],\"name\":\"AddDefaultAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pessimisticVKey\",\"type\":\"bytes32\"}],\"name\":\"RouteAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pessimisticVKey\",\"type\":\"bytes32\"}],\"name\":\"RouteFrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"UnsetDefaultAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousVKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newVKey\",\"type\":\"bytes32\"}],\"name\":\"UpdateDefaultAggchainVKey\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"pessimisticVKeySelector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"pessimisticVKey\",\"type\":\"bytes32\"}],\"name\":\"addPessimisticVKeyRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"pessimisticVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"freezePessimisticVKeyRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"defaultAggchainSelector\",\"type\":\"bytes4\"}],\"name\":\"getDefaultAggchainVKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicValues\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofBytes\",\"type\":\"bytes\"}],\"name\":\"verifyPessimisticProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IagglayergatewaypreviousABI is the input ABI used to generate the binding from.
// Deprecated: Use IagglayergatewaypreviousMetaData.ABI instead.
var IagglayergatewaypreviousABI = IagglayergatewaypreviousMetaData.ABI

// Iagglayergatewayprevious is an auto generated Go binding around an Ethereum contract.
type Iagglayergatewayprevious struct {
	IagglayergatewaypreviousCaller     // Read-only binding to the contract
	IagglayergatewaypreviousTransactor // Write-only binding to the contract
	IagglayergatewaypreviousFilterer   // Log filterer for contract events
}

// IagglayergatewaypreviousCaller is an auto generated read-only Go binding around an Ethereum contract.
type IagglayergatewaypreviousCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IagglayergatewaypreviousTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IagglayergatewaypreviousTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IagglayergatewaypreviousFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IagglayergatewaypreviousFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IagglayergatewaypreviousSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IagglayergatewaypreviousSession struct {
	Contract     *Iagglayergatewayprevious // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IagglayergatewaypreviousCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IagglayergatewaypreviousCallerSession struct {
	Contract *IagglayergatewaypreviousCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// IagglayergatewaypreviousTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IagglayergatewaypreviousTransactorSession struct {
	Contract     *IagglayergatewaypreviousTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// IagglayergatewaypreviousRaw is an auto generated low-level Go binding around an Ethereum contract.
type IagglayergatewaypreviousRaw struct {
	Contract *Iagglayergatewayprevious // Generic contract binding to access the raw methods on
}

// IagglayergatewaypreviousCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IagglayergatewaypreviousCallerRaw struct {
	Contract *IagglayergatewaypreviousCaller // Generic read-only contract binding to access the raw methods on
}

// IagglayergatewaypreviousTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IagglayergatewaypreviousTransactorRaw struct {
	Contract *IagglayergatewaypreviousTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIagglayergatewayprevious creates a new instance of Iagglayergatewayprevious, bound to a specific deployed contract.
func NewIagglayergatewayprevious(address common.Address, backend bind.ContractBackend) (*Iagglayergatewayprevious, error) {
	contract, err := bindIagglayergatewayprevious(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iagglayergatewayprevious{IagglayergatewaypreviousCaller: IagglayergatewaypreviousCaller{contract: contract}, IagglayergatewaypreviousTransactor: IagglayergatewaypreviousTransactor{contract: contract}, IagglayergatewaypreviousFilterer: IagglayergatewaypreviousFilterer{contract: contract}}, nil
}

// NewIagglayergatewaypreviousCaller creates a new read-only instance of Iagglayergatewayprevious, bound to a specific deployed contract.
func NewIagglayergatewaypreviousCaller(address common.Address, caller bind.ContractCaller) (*IagglayergatewaypreviousCaller, error) {
	contract, err := bindIagglayergatewayprevious(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IagglayergatewaypreviousCaller{contract: contract}, nil
}

// NewIagglayergatewaypreviousTransactor creates a new write-only instance of Iagglayergatewayprevious, bound to a specific deployed contract.
func NewIagglayergatewaypreviousTransactor(address common.Address, transactor bind.ContractTransactor) (*IagglayergatewaypreviousTransactor, error) {
	contract, err := bindIagglayergatewayprevious(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IagglayergatewaypreviousTransactor{contract: contract}, nil
}

// NewIagglayergatewaypreviousFilterer creates a new log filterer instance of Iagglayergatewayprevious, bound to a specific deployed contract.
func NewIagglayergatewaypreviousFilterer(address common.Address, filterer bind.ContractFilterer) (*IagglayergatewaypreviousFilterer, error) {
	contract, err := bindIagglayergatewayprevious(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IagglayergatewaypreviousFilterer{contract: contract}, nil
}

// bindIagglayergatewayprevious binds a generic wrapper to an already deployed contract.
func bindIagglayergatewayprevious(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IagglayergatewaypreviousMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iagglayergatewayprevious *IagglayergatewaypreviousRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iagglayergatewayprevious.Contract.IagglayergatewaypreviousCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iagglayergatewayprevious *IagglayergatewaypreviousRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iagglayergatewayprevious.Contract.IagglayergatewaypreviousTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iagglayergatewayprevious *IagglayergatewaypreviousRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iagglayergatewayprevious.Contract.IagglayergatewaypreviousTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iagglayergatewayprevious *IagglayergatewaypreviousCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iagglayergatewayprevious.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iagglayergatewayprevious *IagglayergatewaypreviousTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iagglayergatewayprevious.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iagglayergatewayprevious *IagglayergatewaypreviousTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iagglayergatewayprevious.Contract.contract.Transact(opts, method, params...)
}

// GetDefaultAggchainVKey is a free data retrieval call binding the contract method 0x6cabfdab.
//
// Solidity: function getDefaultAggchainVKey(bytes4 defaultAggchainSelector) view returns(bytes32)
func (_Iagglayergatewayprevious *IagglayergatewaypreviousCaller) GetDefaultAggchainVKey(opts *bind.CallOpts, defaultAggchainSelector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _Iagglayergatewayprevious.contract.Call(opts, &out, "getDefaultAggchainVKey", defaultAggchainSelector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDefaultAggchainVKey is a free data retrieval call binding the contract method 0x6cabfdab.
//
// Solidity: function getDefaultAggchainVKey(bytes4 defaultAggchainSelector) view returns(bytes32)
func (_Iagglayergatewayprevious *IagglayergatewaypreviousSession) GetDefaultAggchainVKey(defaultAggchainSelector [4]byte) ([32]byte, error) {
	return _Iagglayergatewayprevious.Contract.GetDefaultAggchainVKey(&_Iagglayergatewayprevious.CallOpts, defaultAggchainSelector)
}

// GetDefaultAggchainVKey is a free data retrieval call binding the contract method 0x6cabfdab.
//
// Solidity: function getDefaultAggchainVKey(bytes4 defaultAggchainSelector) view returns(bytes32)
func (_Iagglayergatewayprevious *IagglayergatewaypreviousCallerSession) GetDefaultAggchainVKey(defaultAggchainSelector [4]byte) ([32]byte, error) {
	return _Iagglayergatewayprevious.Contract.GetDefaultAggchainVKey(&_Iagglayergatewayprevious.CallOpts, defaultAggchainSelector)
}

// VerifyPessimisticProof is a free data retrieval call binding the contract method 0xa48fd377.
//
// Solidity: function verifyPessimisticProof(bytes publicValues, bytes proofBytes) view returns()
func (_Iagglayergatewayprevious *IagglayergatewaypreviousCaller) VerifyPessimisticProof(opts *bind.CallOpts, publicValues []byte, proofBytes []byte) error {
	var out []interface{}
	err := _Iagglayergatewayprevious.contract.Call(opts, &out, "verifyPessimisticProof", publicValues, proofBytes)

	if err != nil {
		return err
	}

	return err

}

// VerifyPessimisticProof is a free data retrieval call binding the contract method 0xa48fd377.
//
// Solidity: function verifyPessimisticProof(bytes publicValues, bytes proofBytes) view returns()
func (_Iagglayergatewayprevious *IagglayergatewaypreviousSession) VerifyPessimisticProof(publicValues []byte, proofBytes []byte) error {
	return _Iagglayergatewayprevious.Contract.VerifyPessimisticProof(&_Iagglayergatewayprevious.CallOpts, publicValues, proofBytes)
}

// VerifyPessimisticProof is a free data retrieval call binding the contract method 0xa48fd377.
//
// Solidity: function verifyPessimisticProof(bytes publicValues, bytes proofBytes) view returns()
func (_Iagglayergatewayprevious *IagglayergatewaypreviousCallerSession) VerifyPessimisticProof(publicValues []byte, proofBytes []byte) error {
	return _Iagglayergatewayprevious.Contract.VerifyPessimisticProof(&_Iagglayergatewayprevious.CallOpts, publicValues, proofBytes)
}

// AddPessimisticVKeyRoute is a paid mutator transaction binding the contract method 0x9584a516.
//
// Solidity: function addPessimisticVKeyRoute(bytes4 pessimisticVKeySelector, address verifier, bytes32 pessimisticVKey) returns()
func (_Iagglayergatewayprevious *IagglayergatewaypreviousTransactor) AddPessimisticVKeyRoute(opts *bind.TransactOpts, pessimisticVKeySelector [4]byte, verifier common.Address, pessimisticVKey [32]byte) (*types.Transaction, error) {
	return _Iagglayergatewayprevious.contract.Transact(opts, "addPessimisticVKeyRoute", pessimisticVKeySelector, verifier, pessimisticVKey)
}

// AddPessimisticVKeyRoute is a paid mutator transaction binding the contract method 0x9584a516.
//
// Solidity: function addPessimisticVKeyRoute(bytes4 pessimisticVKeySelector, address verifier, bytes32 pessimisticVKey) returns()
func (_Iagglayergatewayprevious *IagglayergatewaypreviousSession) AddPessimisticVKeyRoute(pessimisticVKeySelector [4]byte, verifier common.Address, pessimisticVKey [32]byte) (*types.Transaction, error) {
	return _Iagglayergatewayprevious.Contract.AddPessimisticVKeyRoute(&_Iagglayergatewayprevious.TransactOpts, pessimisticVKeySelector, verifier, pessimisticVKey)
}

// AddPessimisticVKeyRoute is a paid mutator transaction binding the contract method 0x9584a516.
//
// Solidity: function addPessimisticVKeyRoute(bytes4 pessimisticVKeySelector, address verifier, bytes32 pessimisticVKey) returns()
func (_Iagglayergatewayprevious *IagglayergatewaypreviousTransactorSession) AddPessimisticVKeyRoute(pessimisticVKeySelector [4]byte, verifier common.Address, pessimisticVKey [32]byte) (*types.Transaction, error) {
	return _Iagglayergatewayprevious.Contract.AddPessimisticVKeyRoute(&_Iagglayergatewayprevious.TransactOpts, pessimisticVKeySelector, verifier, pessimisticVKey)
}

// FreezePessimisticVKeyRoute is a paid mutator transaction binding the contract method 0x95d2a391.
//
// Solidity: function freezePessimisticVKeyRoute(bytes4 pessimisticVKeySelector) returns()
func (_Iagglayergatewayprevious *IagglayergatewaypreviousTransactor) FreezePessimisticVKeyRoute(opts *bind.TransactOpts, pessimisticVKeySelector [4]byte) (*types.Transaction, error) {
	return _Iagglayergatewayprevious.contract.Transact(opts, "freezePessimisticVKeyRoute", pessimisticVKeySelector)
}

// FreezePessimisticVKeyRoute is a paid mutator transaction binding the contract method 0x95d2a391.
//
// Solidity: function freezePessimisticVKeyRoute(bytes4 pessimisticVKeySelector) returns()
func (_Iagglayergatewayprevious *IagglayergatewaypreviousSession) FreezePessimisticVKeyRoute(pessimisticVKeySelector [4]byte) (*types.Transaction, error) {
	return _Iagglayergatewayprevious.Contract.FreezePessimisticVKeyRoute(&_Iagglayergatewayprevious.TransactOpts, pessimisticVKeySelector)
}

// FreezePessimisticVKeyRoute is a paid mutator transaction binding the contract method 0x95d2a391.
//
// Solidity: function freezePessimisticVKeyRoute(bytes4 pessimisticVKeySelector) returns()
func (_Iagglayergatewayprevious *IagglayergatewaypreviousTransactorSession) FreezePessimisticVKeyRoute(pessimisticVKeySelector [4]byte) (*types.Transaction, error) {
	return _Iagglayergatewayprevious.Contract.FreezePessimisticVKeyRoute(&_Iagglayergatewayprevious.TransactOpts, pessimisticVKeySelector)
}

// IagglayergatewaypreviousAddDefaultAggchainVKeyIterator is returned from FilterAddDefaultAggchainVKey and is used to iterate over the raw logs and unpacked data for AddDefaultAggchainVKey events raised by the Iagglayergatewayprevious contract.
type IagglayergatewaypreviousAddDefaultAggchainVKeyIterator struct {
	Event *IagglayergatewaypreviousAddDefaultAggchainVKey // Event containing the contract specifics and raw log

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
func (it *IagglayergatewaypreviousAddDefaultAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IagglayergatewaypreviousAddDefaultAggchainVKey)
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
		it.Event = new(IagglayergatewaypreviousAddDefaultAggchainVKey)
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
func (it *IagglayergatewaypreviousAddDefaultAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IagglayergatewaypreviousAddDefaultAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IagglayergatewaypreviousAddDefaultAggchainVKey represents a AddDefaultAggchainVKey event raised by the Iagglayergatewayprevious contract.
type IagglayergatewaypreviousAddDefaultAggchainVKey struct {
	Selector [4]byte
	NewVKey  [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddDefaultAggchainVKey is a free log retrieval operation binding the contract event 0x64b3528974574a41ddc830adc96440dccb99e64ede6b863194b951bfbd17e1bc.
//
// Solidity: event AddDefaultAggchainVKey(bytes4 selector, bytes32 newVKey)
func (_Iagglayergatewayprevious *IagglayergatewaypreviousFilterer) FilterAddDefaultAggchainVKey(opts *bind.FilterOpts) (*IagglayergatewaypreviousAddDefaultAggchainVKeyIterator, error) {

	logs, sub, err := _Iagglayergatewayprevious.contract.FilterLogs(opts, "AddDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &IagglayergatewaypreviousAddDefaultAggchainVKeyIterator{contract: _Iagglayergatewayprevious.contract, event: "AddDefaultAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchAddDefaultAggchainVKey is a free log subscription operation binding the contract event 0x64b3528974574a41ddc830adc96440dccb99e64ede6b863194b951bfbd17e1bc.
//
// Solidity: event AddDefaultAggchainVKey(bytes4 selector, bytes32 newVKey)
func (_Iagglayergatewayprevious *IagglayergatewaypreviousFilterer) WatchAddDefaultAggchainVKey(opts *bind.WatchOpts, sink chan<- *IagglayergatewaypreviousAddDefaultAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Iagglayergatewayprevious.contract.WatchLogs(opts, "AddDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IagglayergatewaypreviousAddDefaultAggchainVKey)
				if err := _Iagglayergatewayprevious.contract.UnpackLog(event, "AddDefaultAggchainVKey", log); err != nil {
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
func (_Iagglayergatewayprevious *IagglayergatewaypreviousFilterer) ParseAddDefaultAggchainVKey(log types.Log) (*IagglayergatewaypreviousAddDefaultAggchainVKey, error) {
	event := new(IagglayergatewaypreviousAddDefaultAggchainVKey)
	if err := _Iagglayergatewayprevious.contract.UnpackLog(event, "AddDefaultAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IagglayergatewaypreviousRouteAddedIterator is returned from FilterRouteAdded and is used to iterate over the raw logs and unpacked data for RouteAdded events raised by the Iagglayergatewayprevious contract.
type IagglayergatewaypreviousRouteAddedIterator struct {
	Event *IagglayergatewaypreviousRouteAdded // Event containing the contract specifics and raw log

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
func (it *IagglayergatewaypreviousRouteAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IagglayergatewaypreviousRouteAdded)
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
		it.Event = new(IagglayergatewaypreviousRouteAdded)
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
func (it *IagglayergatewaypreviousRouteAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IagglayergatewaypreviousRouteAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IagglayergatewaypreviousRouteAdded represents a RouteAdded event raised by the Iagglayergatewayprevious contract.
type IagglayergatewaypreviousRouteAdded struct {
	Selector        [4]byte
	Verifier        common.Address
	PessimisticVKey [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRouteAdded is a free log retrieval operation binding the contract event 0x4c4455a30528c319643ae6fc35f3a0fcabde6c015f2a1a5270382c4a190b0fa3.
//
// Solidity: event RouteAdded(bytes4 selector, address verifier, bytes32 pessimisticVKey)
func (_Iagglayergatewayprevious *IagglayergatewaypreviousFilterer) FilterRouteAdded(opts *bind.FilterOpts) (*IagglayergatewaypreviousRouteAddedIterator, error) {

	logs, sub, err := _Iagglayergatewayprevious.contract.FilterLogs(opts, "RouteAdded")
	if err != nil {
		return nil, err
	}
	return &IagglayergatewaypreviousRouteAddedIterator{contract: _Iagglayergatewayprevious.contract, event: "RouteAdded", logs: logs, sub: sub}, nil
}

// WatchRouteAdded is a free log subscription operation binding the contract event 0x4c4455a30528c319643ae6fc35f3a0fcabde6c015f2a1a5270382c4a190b0fa3.
//
// Solidity: event RouteAdded(bytes4 selector, address verifier, bytes32 pessimisticVKey)
func (_Iagglayergatewayprevious *IagglayergatewaypreviousFilterer) WatchRouteAdded(opts *bind.WatchOpts, sink chan<- *IagglayergatewaypreviousRouteAdded) (event.Subscription, error) {

	logs, sub, err := _Iagglayergatewayprevious.contract.WatchLogs(opts, "RouteAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IagglayergatewaypreviousRouteAdded)
				if err := _Iagglayergatewayprevious.contract.UnpackLog(event, "RouteAdded", log); err != nil {
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
func (_Iagglayergatewayprevious *IagglayergatewaypreviousFilterer) ParseRouteAdded(log types.Log) (*IagglayergatewaypreviousRouteAdded, error) {
	event := new(IagglayergatewaypreviousRouteAdded)
	if err := _Iagglayergatewayprevious.contract.UnpackLog(event, "RouteAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IagglayergatewaypreviousRouteFrozenIterator is returned from FilterRouteFrozen and is used to iterate over the raw logs and unpacked data for RouteFrozen events raised by the Iagglayergatewayprevious contract.
type IagglayergatewaypreviousRouteFrozenIterator struct {
	Event *IagglayergatewaypreviousRouteFrozen // Event containing the contract specifics and raw log

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
func (it *IagglayergatewaypreviousRouteFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IagglayergatewaypreviousRouteFrozen)
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
		it.Event = new(IagglayergatewaypreviousRouteFrozen)
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
func (it *IagglayergatewaypreviousRouteFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IagglayergatewaypreviousRouteFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IagglayergatewaypreviousRouteFrozen represents a RouteFrozen event raised by the Iagglayergatewayprevious contract.
type IagglayergatewaypreviousRouteFrozen struct {
	Selector        [4]byte
	Verifier        common.Address
	PessimisticVKey [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRouteFrozen is a free log retrieval operation binding the contract event 0x6efcfd4090b998a9d58c7f753bdda3291e9d781c02226b9f480b18ef07e0012b.
//
// Solidity: event RouteFrozen(bytes4 selector, address verifier, bytes32 pessimisticVKey)
func (_Iagglayergatewayprevious *IagglayergatewaypreviousFilterer) FilterRouteFrozen(opts *bind.FilterOpts) (*IagglayergatewaypreviousRouteFrozenIterator, error) {

	logs, sub, err := _Iagglayergatewayprevious.contract.FilterLogs(opts, "RouteFrozen")
	if err != nil {
		return nil, err
	}
	return &IagglayergatewaypreviousRouteFrozenIterator{contract: _Iagglayergatewayprevious.contract, event: "RouteFrozen", logs: logs, sub: sub}, nil
}

// WatchRouteFrozen is a free log subscription operation binding the contract event 0x6efcfd4090b998a9d58c7f753bdda3291e9d781c02226b9f480b18ef07e0012b.
//
// Solidity: event RouteFrozen(bytes4 selector, address verifier, bytes32 pessimisticVKey)
func (_Iagglayergatewayprevious *IagglayergatewaypreviousFilterer) WatchRouteFrozen(opts *bind.WatchOpts, sink chan<- *IagglayergatewaypreviousRouteFrozen) (event.Subscription, error) {

	logs, sub, err := _Iagglayergatewayprevious.contract.WatchLogs(opts, "RouteFrozen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IagglayergatewaypreviousRouteFrozen)
				if err := _Iagglayergatewayprevious.contract.UnpackLog(event, "RouteFrozen", log); err != nil {
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
func (_Iagglayergatewayprevious *IagglayergatewaypreviousFilterer) ParseRouteFrozen(log types.Log) (*IagglayergatewaypreviousRouteFrozen, error) {
	event := new(IagglayergatewaypreviousRouteFrozen)
	if err := _Iagglayergatewayprevious.contract.UnpackLog(event, "RouteFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IagglayergatewaypreviousUnsetDefaultAggchainVKeyIterator is returned from FilterUnsetDefaultAggchainVKey and is used to iterate over the raw logs and unpacked data for UnsetDefaultAggchainVKey events raised by the Iagglayergatewayprevious contract.
type IagglayergatewaypreviousUnsetDefaultAggchainVKeyIterator struct {
	Event *IagglayergatewaypreviousUnsetDefaultAggchainVKey // Event containing the contract specifics and raw log

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
func (it *IagglayergatewaypreviousUnsetDefaultAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IagglayergatewaypreviousUnsetDefaultAggchainVKey)
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
		it.Event = new(IagglayergatewaypreviousUnsetDefaultAggchainVKey)
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
func (it *IagglayergatewaypreviousUnsetDefaultAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IagglayergatewaypreviousUnsetDefaultAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IagglayergatewaypreviousUnsetDefaultAggchainVKey represents a UnsetDefaultAggchainVKey event raised by the Iagglayergatewayprevious contract.
type IagglayergatewaypreviousUnsetDefaultAggchainVKey struct {
	Selector [4]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnsetDefaultAggchainVKey is a free log retrieval operation binding the contract event 0x43b46d79ee2fbaa8a981b5fd78aaf6fb9e38b18d02e1af17564ce37c67078cab.
//
// Solidity: event UnsetDefaultAggchainVKey(bytes4 selector)
func (_Iagglayergatewayprevious *IagglayergatewaypreviousFilterer) FilterUnsetDefaultAggchainVKey(opts *bind.FilterOpts) (*IagglayergatewaypreviousUnsetDefaultAggchainVKeyIterator, error) {

	logs, sub, err := _Iagglayergatewayprevious.contract.FilterLogs(opts, "UnsetDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &IagglayergatewaypreviousUnsetDefaultAggchainVKeyIterator{contract: _Iagglayergatewayprevious.contract, event: "UnsetDefaultAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchUnsetDefaultAggchainVKey is a free log subscription operation binding the contract event 0x43b46d79ee2fbaa8a981b5fd78aaf6fb9e38b18d02e1af17564ce37c67078cab.
//
// Solidity: event UnsetDefaultAggchainVKey(bytes4 selector)
func (_Iagglayergatewayprevious *IagglayergatewaypreviousFilterer) WatchUnsetDefaultAggchainVKey(opts *bind.WatchOpts, sink chan<- *IagglayergatewaypreviousUnsetDefaultAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Iagglayergatewayprevious.contract.WatchLogs(opts, "UnsetDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IagglayergatewaypreviousUnsetDefaultAggchainVKey)
				if err := _Iagglayergatewayprevious.contract.UnpackLog(event, "UnsetDefaultAggchainVKey", log); err != nil {
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
func (_Iagglayergatewayprevious *IagglayergatewaypreviousFilterer) ParseUnsetDefaultAggchainVKey(log types.Log) (*IagglayergatewaypreviousUnsetDefaultAggchainVKey, error) {
	event := new(IagglayergatewaypreviousUnsetDefaultAggchainVKey)
	if err := _Iagglayergatewayprevious.contract.UnpackLog(event, "UnsetDefaultAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IagglayergatewaypreviousUpdateDefaultAggchainVKeyIterator is returned from FilterUpdateDefaultAggchainVKey and is used to iterate over the raw logs and unpacked data for UpdateDefaultAggchainVKey events raised by the Iagglayergatewayprevious contract.
type IagglayergatewaypreviousUpdateDefaultAggchainVKeyIterator struct {
	Event *IagglayergatewaypreviousUpdateDefaultAggchainVKey // Event containing the contract specifics and raw log

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
func (it *IagglayergatewaypreviousUpdateDefaultAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IagglayergatewaypreviousUpdateDefaultAggchainVKey)
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
		it.Event = new(IagglayergatewaypreviousUpdateDefaultAggchainVKey)
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
func (it *IagglayergatewaypreviousUpdateDefaultAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IagglayergatewaypreviousUpdateDefaultAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IagglayergatewaypreviousUpdateDefaultAggchainVKey represents a UpdateDefaultAggchainVKey event raised by the Iagglayergatewayprevious contract.
type IagglayergatewaypreviousUpdateDefaultAggchainVKey struct {
	Selector     [4]byte
	PreviousVKey [32]byte
	NewVKey      [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateDefaultAggchainVKey is a free log retrieval operation binding the contract event 0x477ac863279848383fdb5d6bb8957cc8f6f10c0200379e075df3f789d14248e6.
//
// Solidity: event UpdateDefaultAggchainVKey(bytes4 selector, bytes32 previousVKey, bytes32 newVKey)
func (_Iagglayergatewayprevious *IagglayergatewaypreviousFilterer) FilterUpdateDefaultAggchainVKey(opts *bind.FilterOpts) (*IagglayergatewaypreviousUpdateDefaultAggchainVKeyIterator, error) {

	logs, sub, err := _Iagglayergatewayprevious.contract.FilterLogs(opts, "UpdateDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &IagglayergatewaypreviousUpdateDefaultAggchainVKeyIterator{contract: _Iagglayergatewayprevious.contract, event: "UpdateDefaultAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchUpdateDefaultAggchainVKey is a free log subscription operation binding the contract event 0x477ac863279848383fdb5d6bb8957cc8f6f10c0200379e075df3f789d14248e6.
//
// Solidity: event UpdateDefaultAggchainVKey(bytes4 selector, bytes32 previousVKey, bytes32 newVKey)
func (_Iagglayergatewayprevious *IagglayergatewaypreviousFilterer) WatchUpdateDefaultAggchainVKey(opts *bind.WatchOpts, sink chan<- *IagglayergatewaypreviousUpdateDefaultAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Iagglayergatewayprevious.contract.WatchLogs(opts, "UpdateDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IagglayergatewaypreviousUpdateDefaultAggchainVKey)
				if err := _Iagglayergatewayprevious.contract.UnpackLog(event, "UpdateDefaultAggchainVKey", log); err != nil {
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
func (_Iagglayergatewayprevious *IagglayergatewaypreviousFilterer) ParseUpdateDefaultAggchainVKey(log types.Log) (*IagglayergatewaypreviousUpdateDefaultAggchainVKey, error) {
	event := new(IagglayergatewaypreviousUpdateDefaultAggchainVKey)
	if err := _Iagglayergatewayprevious.contract.UnpackLog(event, "UpdateDefaultAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

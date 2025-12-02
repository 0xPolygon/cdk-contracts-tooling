// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package depositcontractv2

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

// Depositcontractv2MetaData contains all meta data concerning the Depositcontractv2 contract.
var Depositcontractv2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewDepositCountExceedsMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonZeroValueForUnusedFrontier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubtreeFrontierMismatch\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061012a8061001c5f395ff3fe6080604052348015600e575f5ffd5b50600436106030575f3560e01c80632dfdf0b51460345780635ca1e16514604e575b5f5ffd5b603c60535481565b60405190815260200160405180910390f35b603c6053545f90819081805b602081101560be578083901c600116600103609857609260338260208110608157608160c7565b0154855f9182526020526040902090565b935060a8565b5f84815260208390526040902093505b5f8281526020839052604090209150600101605a565b50919392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea264697066735822122044ddb562dfccbe39855ed5f72e5a00a169f5a28dca755556cc6efec88f9680f164736f6c634300081c0033",
}

// Depositcontractv2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Depositcontractv2MetaData.ABI instead.
var Depositcontractv2ABI = Depositcontractv2MetaData.ABI

// Depositcontractv2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Depositcontractv2MetaData.Bin instead.
var Depositcontractv2Bin = Depositcontractv2MetaData.Bin

// DeployDepositcontractv2 deploys a new Ethereum contract, binding an instance of Depositcontractv2 to it.
func DeployDepositcontractv2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Depositcontractv2, error) {
	parsed, err := Depositcontractv2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Depositcontractv2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Depositcontractv2{Depositcontractv2Caller: Depositcontractv2Caller{contract: contract}, Depositcontractv2Transactor: Depositcontractv2Transactor{contract: contract}, Depositcontractv2Filterer: Depositcontractv2Filterer{contract: contract}}, nil
}

// Depositcontractv2 is an auto generated Go binding around an Ethereum contract.
type Depositcontractv2 struct {
	Depositcontractv2Caller     // Read-only binding to the contract
	Depositcontractv2Transactor // Write-only binding to the contract
	Depositcontractv2Filterer   // Log filterer for contract events
}

// Depositcontractv2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Depositcontractv2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Depositcontractv2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Depositcontractv2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Depositcontractv2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Depositcontractv2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Depositcontractv2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Depositcontractv2Session struct {
	Contract     *Depositcontractv2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Depositcontractv2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Depositcontractv2CallerSession struct {
	Contract *Depositcontractv2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// Depositcontractv2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Depositcontractv2TransactorSession struct {
	Contract     *Depositcontractv2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// Depositcontractv2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Depositcontractv2Raw struct {
	Contract *Depositcontractv2 // Generic contract binding to access the raw methods on
}

// Depositcontractv2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Depositcontractv2CallerRaw struct {
	Contract *Depositcontractv2Caller // Generic read-only contract binding to access the raw methods on
}

// Depositcontractv2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Depositcontractv2TransactorRaw struct {
	Contract *Depositcontractv2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDepositcontractv2 creates a new instance of Depositcontractv2, bound to a specific deployed contract.
func NewDepositcontractv2(address common.Address, backend bind.ContractBackend) (*Depositcontractv2, error) {
	contract, err := bindDepositcontractv2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Depositcontractv2{Depositcontractv2Caller: Depositcontractv2Caller{contract: contract}, Depositcontractv2Transactor: Depositcontractv2Transactor{contract: contract}, Depositcontractv2Filterer: Depositcontractv2Filterer{contract: contract}}, nil
}

// NewDepositcontractv2Caller creates a new read-only instance of Depositcontractv2, bound to a specific deployed contract.
func NewDepositcontractv2Caller(address common.Address, caller bind.ContractCaller) (*Depositcontractv2Caller, error) {
	contract, err := bindDepositcontractv2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Depositcontractv2Caller{contract: contract}, nil
}

// NewDepositcontractv2Transactor creates a new write-only instance of Depositcontractv2, bound to a specific deployed contract.
func NewDepositcontractv2Transactor(address common.Address, transactor bind.ContractTransactor) (*Depositcontractv2Transactor, error) {
	contract, err := bindDepositcontractv2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Depositcontractv2Transactor{contract: contract}, nil
}

// NewDepositcontractv2Filterer creates a new log filterer instance of Depositcontractv2, bound to a specific deployed contract.
func NewDepositcontractv2Filterer(address common.Address, filterer bind.ContractFilterer) (*Depositcontractv2Filterer, error) {
	contract, err := bindDepositcontractv2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Depositcontractv2Filterer{contract: contract}, nil
}

// bindDepositcontractv2 binds a generic wrapper to an already deployed contract.
func bindDepositcontractv2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Depositcontractv2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Depositcontractv2 *Depositcontractv2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Depositcontractv2.Contract.Depositcontractv2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Depositcontractv2 *Depositcontractv2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositcontractv2.Contract.Depositcontractv2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Depositcontractv2 *Depositcontractv2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Depositcontractv2.Contract.Depositcontractv2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Depositcontractv2 *Depositcontractv2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Depositcontractv2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Depositcontractv2 *Depositcontractv2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositcontractv2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Depositcontractv2 *Depositcontractv2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Depositcontractv2.Contract.contract.Transact(opts, method, params...)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Depositcontractv2 *Depositcontractv2Caller) DepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Depositcontractv2.contract.Call(opts, &out, "depositCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Depositcontractv2 *Depositcontractv2Session) DepositCount() (*big.Int, error) {
	return _Depositcontractv2.Contract.DepositCount(&_Depositcontractv2.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Depositcontractv2 *Depositcontractv2CallerSession) DepositCount() (*big.Int, error) {
	return _Depositcontractv2.Contract.DepositCount(&_Depositcontractv2.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Depositcontractv2 *Depositcontractv2Caller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Depositcontractv2.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Depositcontractv2 *Depositcontractv2Session) GetRoot() ([32]byte, error) {
	return _Depositcontractv2.Contract.GetRoot(&_Depositcontractv2.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Depositcontractv2 *Depositcontractv2CallerSession) GetRoot() ([32]byte, error) {
	return _Depositcontractv2.Contract.GetRoot(&_Depositcontractv2.CallOpts)
}

// Depositcontractv2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Depositcontractv2 contract.
type Depositcontractv2InitializedIterator struct {
	Event *Depositcontractv2Initialized // Event containing the contract specifics and raw log

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
func (it *Depositcontractv2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Depositcontractv2Initialized)
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
		it.Event = new(Depositcontractv2Initialized)
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
func (it *Depositcontractv2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Depositcontractv2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Depositcontractv2Initialized represents a Initialized event raised by the Depositcontractv2 contract.
type Depositcontractv2Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Depositcontractv2 *Depositcontractv2Filterer) FilterInitialized(opts *bind.FilterOpts) (*Depositcontractv2InitializedIterator, error) {

	logs, sub, err := _Depositcontractv2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Depositcontractv2InitializedIterator{contract: _Depositcontractv2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Depositcontractv2 *Depositcontractv2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Depositcontractv2Initialized) (event.Subscription, error) {

	logs, sub, err := _Depositcontractv2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Depositcontractv2Initialized)
				if err := _Depositcontractv2.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Depositcontractv2 *Depositcontractv2Filterer) ParseInitialized(log types.Log) (*Depositcontractv2Initialized, error) {
	event := new(Depositcontractv2Initialized)
	if err := _Depositcontractv2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

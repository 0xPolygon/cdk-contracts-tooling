// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonaccesscontrolupgradeable

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

// PolygonaccesscontrolupgradeableMetaData contains all meta data concerning the Polygonaccesscontrolupgradeable contract.
var PolygonaccesscontrolupgradeableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlOnlyCanRenounceRolesForSelf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressDoNotHaveRequiredRole\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PolygonaccesscontrolupgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonaccesscontrolupgradeableMetaData.ABI instead.
var PolygonaccesscontrolupgradeableABI = PolygonaccesscontrolupgradeableMetaData.ABI

// Polygonaccesscontrolupgradeable is an auto generated Go binding around an Ethereum contract.
type Polygonaccesscontrolupgradeable struct {
	PolygonaccesscontrolupgradeableCaller     // Read-only binding to the contract
	PolygonaccesscontrolupgradeableTransactor // Write-only binding to the contract
	PolygonaccesscontrolupgradeableFilterer   // Log filterer for contract events
}

// PolygonaccesscontrolupgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonaccesscontrolupgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonaccesscontrolupgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonaccesscontrolupgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonaccesscontrolupgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonaccesscontrolupgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonaccesscontrolupgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonaccesscontrolupgradeableSession struct {
	Contract     *Polygonaccesscontrolupgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                    // Call options to use throughout this session
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// PolygonaccesscontrolupgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonaccesscontrolupgradeableCallerSession struct {
	Contract *PolygonaccesscontrolupgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                          // Call options to use throughout this session
}

// PolygonaccesscontrolupgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonaccesscontrolupgradeableTransactorSession struct {
	Contract     *PolygonaccesscontrolupgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                          // Transaction auth options to use throughout this session
}

// PolygonaccesscontrolupgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonaccesscontrolupgradeableRaw struct {
	Contract *Polygonaccesscontrolupgradeable // Generic contract binding to access the raw methods on
}

// PolygonaccesscontrolupgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonaccesscontrolupgradeableCallerRaw struct {
	Contract *PolygonaccesscontrolupgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonaccesscontrolupgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonaccesscontrolupgradeableTransactorRaw struct {
	Contract *PolygonaccesscontrolupgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonaccesscontrolupgradeable creates a new instance of Polygonaccesscontrolupgradeable, bound to a specific deployed contract.
func NewPolygonaccesscontrolupgradeable(address common.Address, backend bind.ContractBackend) (*Polygonaccesscontrolupgradeable, error) {
	contract, err := bindPolygonaccesscontrolupgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonaccesscontrolupgradeable{PolygonaccesscontrolupgradeableCaller: PolygonaccesscontrolupgradeableCaller{contract: contract}, PolygonaccesscontrolupgradeableTransactor: PolygonaccesscontrolupgradeableTransactor{contract: contract}, PolygonaccesscontrolupgradeableFilterer: PolygonaccesscontrolupgradeableFilterer{contract: contract}}, nil
}

// NewPolygonaccesscontrolupgradeableCaller creates a new read-only instance of Polygonaccesscontrolupgradeable, bound to a specific deployed contract.
func NewPolygonaccesscontrolupgradeableCaller(address common.Address, caller bind.ContractCaller) (*PolygonaccesscontrolupgradeableCaller, error) {
	contract, err := bindPolygonaccesscontrolupgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonaccesscontrolupgradeableCaller{contract: contract}, nil
}

// NewPolygonaccesscontrolupgradeableTransactor creates a new write-only instance of Polygonaccesscontrolupgradeable, bound to a specific deployed contract.
func NewPolygonaccesscontrolupgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonaccesscontrolupgradeableTransactor, error) {
	contract, err := bindPolygonaccesscontrolupgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonaccesscontrolupgradeableTransactor{contract: contract}, nil
}

// NewPolygonaccesscontrolupgradeableFilterer creates a new log filterer instance of Polygonaccesscontrolupgradeable, bound to a specific deployed contract.
func NewPolygonaccesscontrolupgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonaccesscontrolupgradeableFilterer, error) {
	contract, err := bindPolygonaccesscontrolupgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonaccesscontrolupgradeableFilterer{contract: contract}, nil
}

// bindPolygonaccesscontrolupgradeable binds a generic wrapper to an already deployed contract.
func bindPolygonaccesscontrolupgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonaccesscontrolupgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonaccesscontrolupgradeable.Contract.PolygonaccesscontrolupgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonaccesscontrolupgradeable.Contract.PolygonaccesscontrolupgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonaccesscontrolupgradeable.Contract.PolygonaccesscontrolupgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonaccesscontrolupgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonaccesscontrolupgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonaccesscontrolupgradeable.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonaccesscontrolupgradeable.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Polygonaccesscontrolupgradeable.Contract.DEFAULTADMINROLE(&_Polygonaccesscontrolupgradeable.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Polygonaccesscontrolupgradeable.Contract.DEFAULTADMINROLE(&_Polygonaccesscontrolupgradeable.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Polygonaccesscontrolupgradeable.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Polygonaccesscontrolupgradeable.Contract.GetRoleAdmin(&_Polygonaccesscontrolupgradeable.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Polygonaccesscontrolupgradeable.Contract.GetRoleAdmin(&_Polygonaccesscontrolupgradeable.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Polygonaccesscontrolupgradeable.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Polygonaccesscontrolupgradeable.Contract.HasRole(&_Polygonaccesscontrolupgradeable.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Polygonaccesscontrolupgradeable.Contract.HasRole(&_Polygonaccesscontrolupgradeable.CallOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonaccesscontrolupgradeable.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonaccesscontrolupgradeable.Contract.GrantRole(&_Polygonaccesscontrolupgradeable.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonaccesscontrolupgradeable.Contract.GrantRole(&_Polygonaccesscontrolupgradeable.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonaccesscontrolupgradeable.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonaccesscontrolupgradeable.Contract.RenounceRole(&_Polygonaccesscontrolupgradeable.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonaccesscontrolupgradeable.Contract.RenounceRole(&_Polygonaccesscontrolupgradeable.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonaccesscontrolupgradeable.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonaccesscontrolupgradeable.Contract.RevokeRole(&_Polygonaccesscontrolupgradeable.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonaccesscontrolupgradeable.Contract.RevokeRole(&_Polygonaccesscontrolupgradeable.TransactOpts, role, account)
}

// PolygonaccesscontrolupgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonaccesscontrolupgradeable contract.
type PolygonaccesscontrolupgradeableInitializedIterator struct {
	Event *PolygonaccesscontrolupgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *PolygonaccesscontrolupgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonaccesscontrolupgradeableInitialized)
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
		it.Event = new(PolygonaccesscontrolupgradeableInitialized)
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
func (it *PolygonaccesscontrolupgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonaccesscontrolupgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonaccesscontrolupgradeableInitialized represents a Initialized event raised by the Polygonaccesscontrolupgradeable contract.
type PolygonaccesscontrolupgradeableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*PolygonaccesscontrolupgradeableInitializedIterator, error) {

	logs, sub, err := _Polygonaccesscontrolupgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PolygonaccesscontrolupgradeableInitializedIterator{contract: _Polygonaccesscontrolupgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PolygonaccesscontrolupgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _Polygonaccesscontrolupgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonaccesscontrolupgradeableInitialized)
				if err := _Polygonaccesscontrolupgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableFilterer) ParseInitialized(log types.Log) (*PolygonaccesscontrolupgradeableInitialized, error) {
	event := new(PolygonaccesscontrolupgradeableInitialized)
	if err := _Polygonaccesscontrolupgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonaccesscontrolupgradeableRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Polygonaccesscontrolupgradeable contract.
type PolygonaccesscontrolupgradeableRoleAdminChangedIterator struct {
	Event *PolygonaccesscontrolupgradeableRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PolygonaccesscontrolupgradeableRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonaccesscontrolupgradeableRoleAdminChanged)
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
		it.Event = new(PolygonaccesscontrolupgradeableRoleAdminChanged)
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
func (it *PolygonaccesscontrolupgradeableRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonaccesscontrolupgradeableRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonaccesscontrolupgradeableRoleAdminChanged represents a RoleAdminChanged event raised by the Polygonaccesscontrolupgradeable contract.
type PolygonaccesscontrolupgradeableRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PolygonaccesscontrolupgradeableRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Polygonaccesscontrolupgradeable.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PolygonaccesscontrolupgradeableRoleAdminChangedIterator{contract: _Polygonaccesscontrolupgradeable.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PolygonaccesscontrolupgradeableRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Polygonaccesscontrolupgradeable.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonaccesscontrolupgradeableRoleAdminChanged)
				if err := _Polygonaccesscontrolupgradeable.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableFilterer) ParseRoleAdminChanged(log types.Log) (*PolygonaccesscontrolupgradeableRoleAdminChanged, error) {
	event := new(PolygonaccesscontrolupgradeableRoleAdminChanged)
	if err := _Polygonaccesscontrolupgradeable.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonaccesscontrolupgradeableRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Polygonaccesscontrolupgradeable contract.
type PolygonaccesscontrolupgradeableRoleGrantedIterator struct {
	Event *PolygonaccesscontrolupgradeableRoleGranted // Event containing the contract specifics and raw log

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
func (it *PolygonaccesscontrolupgradeableRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonaccesscontrolupgradeableRoleGranted)
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
		it.Event = new(PolygonaccesscontrolupgradeableRoleGranted)
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
func (it *PolygonaccesscontrolupgradeableRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonaccesscontrolupgradeableRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonaccesscontrolupgradeableRoleGranted represents a RoleGranted event raised by the Polygonaccesscontrolupgradeable contract.
type PolygonaccesscontrolupgradeableRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PolygonaccesscontrolupgradeableRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Polygonaccesscontrolupgradeable.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PolygonaccesscontrolupgradeableRoleGrantedIterator{contract: _Polygonaccesscontrolupgradeable.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PolygonaccesscontrolupgradeableRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Polygonaccesscontrolupgradeable.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonaccesscontrolupgradeableRoleGranted)
				if err := _Polygonaccesscontrolupgradeable.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableFilterer) ParseRoleGranted(log types.Log) (*PolygonaccesscontrolupgradeableRoleGranted, error) {
	event := new(PolygonaccesscontrolupgradeableRoleGranted)
	if err := _Polygonaccesscontrolupgradeable.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonaccesscontrolupgradeableRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Polygonaccesscontrolupgradeable contract.
type PolygonaccesscontrolupgradeableRoleRevokedIterator struct {
	Event *PolygonaccesscontrolupgradeableRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PolygonaccesscontrolupgradeableRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonaccesscontrolupgradeableRoleRevoked)
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
		it.Event = new(PolygonaccesscontrolupgradeableRoleRevoked)
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
func (it *PolygonaccesscontrolupgradeableRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonaccesscontrolupgradeableRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonaccesscontrolupgradeableRoleRevoked represents a RoleRevoked event raised by the Polygonaccesscontrolupgradeable contract.
type PolygonaccesscontrolupgradeableRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PolygonaccesscontrolupgradeableRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Polygonaccesscontrolupgradeable.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PolygonaccesscontrolupgradeableRoleRevokedIterator{contract: _Polygonaccesscontrolupgradeable.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PolygonaccesscontrolupgradeableRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Polygonaccesscontrolupgradeable.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonaccesscontrolupgradeableRoleRevoked)
				if err := _Polygonaccesscontrolupgradeable.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonaccesscontrolupgradeable *PolygonaccesscontrolupgradeableFilterer) ParseRoleRevoked(log types.Log) (*PolygonaccesscontrolupgradeableRoleRevoked, error) {
	event := new(PolygonaccesscontrolupgradeableRoleRevoked)
	if err := _Polygonaccesscontrolupgradeable.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

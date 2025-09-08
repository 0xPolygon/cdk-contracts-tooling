// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iaggchainsigners

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

// IAggchainSignersSignerInfo is an auto generated low-level Go binding around an user-defined struct.
type IAggchainSignersSignerInfo struct {
	Addr common.Address
	Url  string
}

// IaggchainsignersMetaData contains all meta data concerning the Iaggchainsigners contract.
var IaggchainsignersMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"aggchainSigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainMultisigHash\",\"type\":\"bytes32\"}],\"name\":\"SignersAndThresholdUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getAggchainMultisigHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggchainSignerInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structIAggchainSigners.SignerInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggchainSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggchainSignersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IaggchainsignersABI is the input ABI used to generate the binding from.
// Deprecated: Use IaggchainsignersMetaData.ABI instead.
var IaggchainsignersABI = IaggchainsignersMetaData.ABI

// Iaggchainsigners is an auto generated Go binding around an Ethereum contract.
type Iaggchainsigners struct {
	IaggchainsignersCaller     // Read-only binding to the contract
	IaggchainsignersTransactor // Write-only binding to the contract
	IaggchainsignersFilterer   // Log filterer for contract events
}

// IaggchainsignersCaller is an auto generated read-only Go binding around an Ethereum contract.
type IaggchainsignersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggchainsignersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IaggchainsignersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggchainsignersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IaggchainsignersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggchainsignersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IaggchainsignersSession struct {
	Contract     *Iaggchainsigners // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IaggchainsignersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IaggchainsignersCallerSession struct {
	Contract *IaggchainsignersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IaggchainsignersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IaggchainsignersTransactorSession struct {
	Contract     *IaggchainsignersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IaggchainsignersRaw is an auto generated low-level Go binding around an Ethereum contract.
type IaggchainsignersRaw struct {
	Contract *Iaggchainsigners // Generic contract binding to access the raw methods on
}

// IaggchainsignersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IaggchainsignersCallerRaw struct {
	Contract *IaggchainsignersCaller // Generic read-only contract binding to access the raw methods on
}

// IaggchainsignersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IaggchainsignersTransactorRaw struct {
	Contract *IaggchainsignersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIaggchainsigners creates a new instance of Iaggchainsigners, bound to a specific deployed contract.
func NewIaggchainsigners(address common.Address, backend bind.ContractBackend) (*Iaggchainsigners, error) {
	contract, err := bindIaggchainsigners(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iaggchainsigners{IaggchainsignersCaller: IaggchainsignersCaller{contract: contract}, IaggchainsignersTransactor: IaggchainsignersTransactor{contract: contract}, IaggchainsignersFilterer: IaggchainsignersFilterer{contract: contract}}, nil
}

// NewIaggchainsignersCaller creates a new read-only instance of Iaggchainsigners, bound to a specific deployed contract.
func NewIaggchainsignersCaller(address common.Address, caller bind.ContractCaller) (*IaggchainsignersCaller, error) {
	contract, err := bindIaggchainsigners(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IaggchainsignersCaller{contract: contract}, nil
}

// NewIaggchainsignersTransactor creates a new write-only instance of Iaggchainsigners, bound to a specific deployed contract.
func NewIaggchainsignersTransactor(address common.Address, transactor bind.ContractTransactor) (*IaggchainsignersTransactor, error) {
	contract, err := bindIaggchainsigners(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IaggchainsignersTransactor{contract: contract}, nil
}

// NewIaggchainsignersFilterer creates a new log filterer instance of Iaggchainsigners, bound to a specific deployed contract.
func NewIaggchainsignersFilterer(address common.Address, filterer bind.ContractFilterer) (*IaggchainsignersFilterer, error) {
	contract, err := bindIaggchainsigners(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IaggchainsignersFilterer{contract: contract}, nil
}

// bindIaggchainsigners binds a generic wrapper to an already deployed contract.
func bindIaggchainsigners(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IaggchainsignersMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iaggchainsigners *IaggchainsignersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iaggchainsigners.Contract.IaggchainsignersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iaggchainsigners *IaggchainsignersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iaggchainsigners.Contract.IaggchainsignersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iaggchainsigners *IaggchainsignersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iaggchainsigners.Contract.IaggchainsignersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iaggchainsigners *IaggchainsignersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iaggchainsigners.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iaggchainsigners *IaggchainsignersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iaggchainsigners.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iaggchainsigners *IaggchainsignersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iaggchainsigners.Contract.contract.Transact(opts, method, params...)
}

// GetAggchainMultisigHash is a free data retrieval call binding the contract method 0xcce7d0df.
//
// Solidity: function getAggchainMultisigHash() view returns(bytes32)
func (_Iaggchainsigners *IaggchainsignersCaller) GetAggchainMultisigHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Iaggchainsigners.contract.Call(opts, &out, "getAggchainMultisigHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAggchainMultisigHash is a free data retrieval call binding the contract method 0xcce7d0df.
//
// Solidity: function getAggchainMultisigHash() view returns(bytes32)
func (_Iaggchainsigners *IaggchainsignersSession) GetAggchainMultisigHash() ([32]byte, error) {
	return _Iaggchainsigners.Contract.GetAggchainMultisigHash(&_Iaggchainsigners.CallOpts)
}

// GetAggchainMultisigHash is a free data retrieval call binding the contract method 0xcce7d0df.
//
// Solidity: function getAggchainMultisigHash() view returns(bytes32)
func (_Iaggchainsigners *IaggchainsignersCallerSession) GetAggchainMultisigHash() ([32]byte, error) {
	return _Iaggchainsigners.Contract.GetAggchainMultisigHash(&_Iaggchainsigners.CallOpts)
}

// GetAggchainSignerInfos is a free data retrieval call binding the contract method 0x349d4046.
//
// Solidity: function getAggchainSignerInfos() view returns((address,string)[])
func (_Iaggchainsigners *IaggchainsignersCaller) GetAggchainSignerInfos(opts *bind.CallOpts) ([]IAggchainSignersSignerInfo, error) {
	var out []interface{}
	err := _Iaggchainsigners.contract.Call(opts, &out, "getAggchainSignerInfos")

	if err != nil {
		return *new([]IAggchainSignersSignerInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAggchainSignersSignerInfo)).(*[]IAggchainSignersSignerInfo)

	return out0, err

}

// GetAggchainSignerInfos is a free data retrieval call binding the contract method 0x349d4046.
//
// Solidity: function getAggchainSignerInfos() view returns((address,string)[])
func (_Iaggchainsigners *IaggchainsignersSession) GetAggchainSignerInfos() ([]IAggchainSignersSignerInfo, error) {
	return _Iaggchainsigners.Contract.GetAggchainSignerInfos(&_Iaggchainsigners.CallOpts)
}

// GetAggchainSignerInfos is a free data retrieval call binding the contract method 0x349d4046.
//
// Solidity: function getAggchainSignerInfos() view returns((address,string)[])
func (_Iaggchainsigners *IaggchainsignersCallerSession) GetAggchainSignerInfos() ([]IAggchainSignersSignerInfo, error) {
	return _Iaggchainsigners.Contract.GetAggchainSignerInfos(&_Iaggchainsigners.CallOpts)
}

// GetAggchainSigners is a free data retrieval call binding the contract method 0x3e1e0121.
//
// Solidity: function getAggchainSigners() view returns(address[])
func (_Iaggchainsigners *IaggchainsignersCaller) GetAggchainSigners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Iaggchainsigners.contract.Call(opts, &out, "getAggchainSigners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAggchainSigners is a free data retrieval call binding the contract method 0x3e1e0121.
//
// Solidity: function getAggchainSigners() view returns(address[])
func (_Iaggchainsigners *IaggchainsignersSession) GetAggchainSigners() ([]common.Address, error) {
	return _Iaggchainsigners.Contract.GetAggchainSigners(&_Iaggchainsigners.CallOpts)
}

// GetAggchainSigners is a free data retrieval call binding the contract method 0x3e1e0121.
//
// Solidity: function getAggchainSigners() view returns(address[])
func (_Iaggchainsigners *IaggchainsignersCallerSession) GetAggchainSigners() ([]common.Address, error) {
	return _Iaggchainsigners.Contract.GetAggchainSigners(&_Iaggchainsigners.CallOpts)
}

// GetAggchainSignersCount is a free data retrieval call binding the contract method 0xca69e7dc.
//
// Solidity: function getAggchainSignersCount() view returns(uint256)
func (_Iaggchainsigners *IaggchainsignersCaller) GetAggchainSignersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Iaggchainsigners.contract.Call(opts, &out, "getAggchainSignersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAggchainSignersCount is a free data retrieval call binding the contract method 0xca69e7dc.
//
// Solidity: function getAggchainSignersCount() view returns(uint256)
func (_Iaggchainsigners *IaggchainsignersSession) GetAggchainSignersCount() (*big.Int, error) {
	return _Iaggchainsigners.Contract.GetAggchainSignersCount(&_Iaggchainsigners.CallOpts)
}

// GetAggchainSignersCount is a free data retrieval call binding the contract method 0xca69e7dc.
//
// Solidity: function getAggchainSignersCount() view returns(uint256)
func (_Iaggchainsigners *IaggchainsignersCallerSession) GetAggchainSignersCount() (*big.Int, error) {
	return _Iaggchainsigners.Contract.GetAggchainSignersCount(&_Iaggchainsigners.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_Iaggchainsigners *IaggchainsignersCaller) GetThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Iaggchainsigners.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_Iaggchainsigners *IaggchainsignersSession) GetThreshold() (*big.Int, error) {
	return _Iaggchainsigners.Contract.GetThreshold(&_Iaggchainsigners.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_Iaggchainsigners *IaggchainsignersCallerSession) GetThreshold() (*big.Int, error) {
	return _Iaggchainsigners.Contract.GetThreshold(&_Iaggchainsigners.CallOpts)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address _signer) view returns(bool)
func (_Iaggchainsigners *IaggchainsignersCaller) IsSigner(opts *bind.CallOpts, _signer common.Address) (bool, error) {
	var out []interface{}
	err := _Iaggchainsigners.contract.Call(opts, &out, "isSigner", _signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address _signer) view returns(bool)
func (_Iaggchainsigners *IaggchainsignersSession) IsSigner(_signer common.Address) (bool, error) {
	return _Iaggchainsigners.Contract.IsSigner(&_Iaggchainsigners.CallOpts, _signer)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address _signer) view returns(bool)
func (_Iaggchainsigners *IaggchainsignersCallerSession) IsSigner(_signer common.Address) (bool, error) {
	return _Iaggchainsigners.Contract.IsSigner(&_Iaggchainsigners.CallOpts, _signer)
}

// IaggchainsignersSignersAndThresholdUpdatedIterator is returned from FilterSignersAndThresholdUpdated and is used to iterate over the raw logs and unpacked data for SignersAndThresholdUpdated events raised by the Iaggchainsigners contract.
type IaggchainsignersSignersAndThresholdUpdatedIterator struct {
	Event *IaggchainsignersSignersAndThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *IaggchainsignersSignersAndThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggchainsignersSignersAndThresholdUpdated)
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
		it.Event = new(IaggchainsignersSignersAndThresholdUpdated)
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
func (it *IaggchainsignersSignersAndThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggchainsignersSignersAndThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggchainsignersSignersAndThresholdUpdated represents a SignersAndThresholdUpdated event raised by the Iaggchainsigners contract.
type IaggchainsignersSignersAndThresholdUpdated struct {
	AggchainSigners         []common.Address
	NewThreshold            *big.Int
	NewAggchainMultisigHash [32]byte
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSignersAndThresholdUpdated is a free log retrieval operation binding the contract event 0x66d7b0647fdd512b69cbf4f8e1ce8068bfe0b236168e2704ba13b07425eaa743.
//
// Solidity: event SignersAndThresholdUpdated(address[] aggchainSigners, uint256 newThreshold, bytes32 newAggchainMultisigHash)
func (_Iaggchainsigners *IaggchainsignersFilterer) FilterSignersAndThresholdUpdated(opts *bind.FilterOpts) (*IaggchainsignersSignersAndThresholdUpdatedIterator, error) {

	logs, sub, err := _Iaggchainsigners.contract.FilterLogs(opts, "SignersAndThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &IaggchainsignersSignersAndThresholdUpdatedIterator{contract: _Iaggchainsigners.contract, event: "SignersAndThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchSignersAndThresholdUpdated is a free log subscription operation binding the contract event 0x66d7b0647fdd512b69cbf4f8e1ce8068bfe0b236168e2704ba13b07425eaa743.
//
// Solidity: event SignersAndThresholdUpdated(address[] aggchainSigners, uint256 newThreshold, bytes32 newAggchainMultisigHash)
func (_Iaggchainsigners *IaggchainsignersFilterer) WatchSignersAndThresholdUpdated(opts *bind.WatchOpts, sink chan<- *IaggchainsignersSignersAndThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _Iaggchainsigners.contract.WatchLogs(opts, "SignersAndThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggchainsignersSignersAndThresholdUpdated)
				if err := _Iaggchainsigners.contract.UnpackLog(event, "SignersAndThresholdUpdated", log); err != nil {
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
// Solidity: event SignersAndThresholdUpdated(address[] aggchainSigners, uint256 newThreshold, bytes32 newAggchainMultisigHash)
func (_Iaggchainsigners *IaggchainsignersFilterer) ParseSignersAndThresholdUpdated(log types.Log) (*IaggchainsignersSignersAndThresholdUpdated, error) {
	event := new(IaggchainsignersSignersAndThresholdUpdated)
	if err := _Iaggchainsigners.contract.UnpackLog(event, "SignersAndThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

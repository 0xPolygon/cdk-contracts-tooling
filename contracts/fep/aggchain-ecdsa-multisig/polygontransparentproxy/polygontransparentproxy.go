// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygontransparentproxy

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

// PolygontransparentproxyMetaData contains all meta data concerning the Polygontransparentproxy contract.
var PolygontransparentproxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"}]",
	Bin: "0x60a060405260405161088b38038061088b83398101604081905261002291610327565b828161002e8282610056565b50506001600160a01b03821660805261004e61004960805190565b6100b4565b50505061040e565b61005f82610121565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156100a8576100a3828261019f565b505050565b6100b0610212565b5050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6100f35f51602061086b5f395f51905f52546001600160a01b031690565b604080516001600160a01b03928316815291841660208301520160405180910390a161011e81610233565b50565b806001600160a01b03163b5f0361015b57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5b80546001600160a01b0319166001600160a01b039290921691909117905550565b60605f5f846001600160a01b0316846040516101bb91906103f8565b5f60405180830381855af49150503d805f81146101f3576040519150601f19603f3d011682016040523d82523d5f602084013e6101f8565b606091505b509092509050610209858383610270565b95945050505050565b34156102315760405163b398979f60e01b815260040160405180910390fd5b565b6001600160a01b03811661025c57604051633173bdd160e11b81525f6004820152602401610152565b805f51602061086b5f395f51905f5261017e565b60608261028557610280826102cf565b6102c8565b815115801561029c57506001600160a01b0384163b155b156102c557604051639996b31560e01b81526001600160a01b0385166004820152602401610152565b50805b9392505050565b8051156102df5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b80516001600160a01b038116811461030e575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f60608486031215610339575f5ffd5b610342846102f8565b9250610350602085016102f8565b60408501519092506001600160401b0381111561036b575f5ffd5b8401601f8101861361037b575f5ffd5b80516001600160401b0381111561039457610394610313565b604051601f8201601f19908116603f011681016001600160401b03811182821017156103c2576103c2610313565b6040528181528282016020018810156103d9575f5ffd5b8160208401602083015e5f602083830101528093505050509250925092565b5f82518060208501845e5f920191825250919050565b6080516104466104255f395f601001526104465ff3fe608060405261000c61000e565b005b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163303610069575f356001600160e01b03191663278f794360e11b146100615761005f61006d565b565b61005f61007d565b61005f5b61005f6100786100ab565b6100e2565b5f8061008c36600481846102ee565b8101906100999190610329565b915091506100a78282610100565b5050565b5f6100dd7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b905090565b365f5f375f5f365f845af43d5f5f3e8080156100fc573d5ff35b3d5ffd5b6101098261015a565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156101525761014d82826101d5565b505050565b6100a7610247565b806001600160a01b03163b5f0361019457604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b0316846040516101f191906103fa565b5f60405180830381855af49150503d805f8114610229576040519150601f19603f3d011682016040523d82523d5f602084013e61022e565b606091505b509150915061023e858383610266565b95945050505050565b341561005f5760405163b398979f60e01b815260040160405180910390fd5b60608261027b57610276826102c5565b6102be565b815115801561029257506001600160a01b0384163b155b156102bb57604051639996b31560e01b81526001600160a01b038516600482015260240161018b565b50805b9392505050565b8051156102d55780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b5f5f858511156102fc575f5ffd5b83861115610308575f5ffd5b5050820193919092039150565b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561033a575f5ffd5b82356001600160a01b0381168114610350575f5ffd5b9150602083013567ffffffffffffffff81111561036b575f5ffd5b8301601f8101851361037b575f5ffd5b803567ffffffffffffffff81111561039557610395610315565b604051601f8201601f19908116603f0116810167ffffffffffffffff811182821017156103c4576103c4610315565b6040528181528282016020018710156103db575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f82518060208501845e5f92019182525091905056fea2646970667358221220feaf4b46b8ab9968f0c64da39e84fea5f82745ccc735417771c87125c48cd4a364736f6c634300081c0033b53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103",
}

// PolygontransparentproxyABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygontransparentproxyMetaData.ABI instead.
var PolygontransparentproxyABI = PolygontransparentproxyMetaData.ABI

// PolygontransparentproxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygontransparentproxyMetaData.Bin instead.
var PolygontransparentproxyBin = PolygontransparentproxyMetaData.Bin

// DeployPolygontransparentproxy deploys a new Ethereum contract, binding an instance of Polygontransparentproxy to it.
func DeployPolygontransparentproxy(auth *bind.TransactOpts, backend bind.ContractBackend, _logic common.Address, admin common.Address, _data []byte) (common.Address, *types.Transaction, *Polygontransparentproxy, error) {
	parsed, err := PolygontransparentproxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygontransparentproxyBin), backend, _logic, admin, _data)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygontransparentproxy{PolygontransparentproxyCaller: PolygontransparentproxyCaller{contract: contract}, PolygontransparentproxyTransactor: PolygontransparentproxyTransactor{contract: contract}, PolygontransparentproxyFilterer: PolygontransparentproxyFilterer{contract: contract}}, nil
}

// Polygontransparentproxy is an auto generated Go binding around an Ethereum contract.
type Polygontransparentproxy struct {
	PolygontransparentproxyCaller     // Read-only binding to the contract
	PolygontransparentproxyTransactor // Write-only binding to the contract
	PolygontransparentproxyFilterer   // Log filterer for contract events
}

// PolygontransparentproxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygontransparentproxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygontransparentproxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygontransparentproxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygontransparentproxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygontransparentproxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygontransparentproxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygontransparentproxySession struct {
	Contract     *Polygontransparentproxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PolygontransparentproxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygontransparentproxyCallerSession struct {
	Contract *PolygontransparentproxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// PolygontransparentproxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygontransparentproxyTransactorSession struct {
	Contract     *PolygontransparentproxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// PolygontransparentproxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygontransparentproxyRaw struct {
	Contract *Polygontransparentproxy // Generic contract binding to access the raw methods on
}

// PolygontransparentproxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygontransparentproxyCallerRaw struct {
	Contract *PolygontransparentproxyCaller // Generic read-only contract binding to access the raw methods on
}

// PolygontransparentproxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygontransparentproxyTransactorRaw struct {
	Contract *PolygontransparentproxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygontransparentproxy creates a new instance of Polygontransparentproxy, bound to a specific deployed contract.
func NewPolygontransparentproxy(address common.Address, backend bind.ContractBackend) (*Polygontransparentproxy, error) {
	contract, err := bindPolygontransparentproxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygontransparentproxy{PolygontransparentproxyCaller: PolygontransparentproxyCaller{contract: contract}, PolygontransparentproxyTransactor: PolygontransparentproxyTransactor{contract: contract}, PolygontransparentproxyFilterer: PolygontransparentproxyFilterer{contract: contract}}, nil
}

// NewPolygontransparentproxyCaller creates a new read-only instance of Polygontransparentproxy, bound to a specific deployed contract.
func NewPolygontransparentproxyCaller(address common.Address, caller bind.ContractCaller) (*PolygontransparentproxyCaller, error) {
	contract, err := bindPolygontransparentproxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygontransparentproxyCaller{contract: contract}, nil
}

// NewPolygontransparentproxyTransactor creates a new write-only instance of Polygontransparentproxy, bound to a specific deployed contract.
func NewPolygontransparentproxyTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygontransparentproxyTransactor, error) {
	contract, err := bindPolygontransparentproxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygontransparentproxyTransactor{contract: contract}, nil
}

// NewPolygontransparentproxyFilterer creates a new log filterer instance of Polygontransparentproxy, bound to a specific deployed contract.
func NewPolygontransparentproxyFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygontransparentproxyFilterer, error) {
	contract, err := bindPolygontransparentproxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygontransparentproxyFilterer{contract: contract}, nil
}

// bindPolygontransparentproxy binds a generic wrapper to an already deployed contract.
func bindPolygontransparentproxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygontransparentproxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygontransparentproxy *PolygontransparentproxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygontransparentproxy.Contract.PolygontransparentproxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygontransparentproxy *PolygontransparentproxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygontransparentproxy.Contract.PolygontransparentproxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygontransparentproxy *PolygontransparentproxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygontransparentproxy.Contract.PolygontransparentproxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygontransparentproxy *PolygontransparentproxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygontransparentproxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygontransparentproxy *PolygontransparentproxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygontransparentproxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygontransparentproxy *PolygontransparentproxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygontransparentproxy.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Polygontransparentproxy *PolygontransparentproxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Polygontransparentproxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Polygontransparentproxy *PolygontransparentproxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Polygontransparentproxy.Contract.Fallback(&_Polygontransparentproxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Polygontransparentproxy *PolygontransparentproxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Polygontransparentproxy.Contract.Fallback(&_Polygontransparentproxy.TransactOpts, calldata)
}

// PolygontransparentproxyAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Polygontransparentproxy contract.
type PolygontransparentproxyAdminChangedIterator struct {
	Event *PolygontransparentproxyAdminChanged // Event containing the contract specifics and raw log

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
func (it *PolygontransparentproxyAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygontransparentproxyAdminChanged)
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
		it.Event = new(PolygontransparentproxyAdminChanged)
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
func (it *PolygontransparentproxyAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygontransparentproxyAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygontransparentproxyAdminChanged represents a AdminChanged event raised by the Polygontransparentproxy contract.
type PolygontransparentproxyAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Polygontransparentproxy *PolygontransparentproxyFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*PolygontransparentproxyAdminChangedIterator, error) {

	logs, sub, err := _Polygontransparentproxy.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &PolygontransparentproxyAdminChangedIterator{contract: _Polygontransparentproxy.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Polygontransparentproxy *PolygontransparentproxyFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *PolygontransparentproxyAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Polygontransparentproxy.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygontransparentproxyAdminChanged)
				if err := _Polygontransparentproxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Polygontransparentproxy *PolygontransparentproxyFilterer) ParseAdminChanged(log types.Log) (*PolygontransparentproxyAdminChanged, error) {
	event := new(PolygontransparentproxyAdminChanged)
	if err := _Polygontransparentproxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygontransparentproxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Polygontransparentproxy contract.
type PolygontransparentproxyUpgradedIterator struct {
	Event *PolygontransparentproxyUpgraded // Event containing the contract specifics and raw log

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
func (it *PolygontransparentproxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygontransparentproxyUpgraded)
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
		it.Event = new(PolygontransparentproxyUpgraded)
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
func (it *PolygontransparentproxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygontransparentproxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygontransparentproxyUpgraded represents a Upgraded event raised by the Polygontransparentproxy contract.
type PolygontransparentproxyUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Polygontransparentproxy *PolygontransparentproxyFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*PolygontransparentproxyUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Polygontransparentproxy.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &PolygontransparentproxyUpgradedIterator{contract: _Polygontransparentproxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Polygontransparentproxy *PolygontransparentproxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PolygontransparentproxyUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Polygontransparentproxy.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygontransparentproxyUpgraded)
				if err := _Polygontransparentproxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Polygontransparentproxy *PolygontransparentproxyFilterer) ParseUpgraded(log types.Log) (*PolygontransparentproxyUpgraded, error) {
	event := new(PolygontransparentproxyUpgraded)
	if err := _Polygontransparentproxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

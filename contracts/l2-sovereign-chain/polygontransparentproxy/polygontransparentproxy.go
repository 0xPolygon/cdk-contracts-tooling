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
	Bin: "0x60a06040526040516108ec3803806108ec83398101604081905261002291610349565b828161002e8282610056565b50506001600160a01b03821660805261004e61004960805190565b6100b4565b505050610434565b61005f82610121565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156100a8576100a3828261019f565b505050565b6100b0610212565b5050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6100f35f5160206108cc5f395f51905f52546001600160a01b031690565b604080516001600160a01b03928316815291841660208301520160405180910390a161011e81610233565b50565b806001600160a01b03163b5f0361015b57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5b80546001600160a01b0319166001600160a01b039290921691909117905550565b60605f5f846001600160a01b0316846040516101bb9190610419565b5f60405180830381855af49150503d805f81146101f3576040519150601f19603f3d011682016040523d82523d5f602084013e6101f8565b606091505b509092509050610209858383610270565b95945050505050565b34156102315760405163b398979f60e01b815260040160405180910390fd5b565b6001600160a01b03811661025c57604051633173bdd160e11b81525f6004820152602401610152565b805f5160206108cc5f395f51905f5261017e565b60608261028557610280826102cf565b6102c8565b815115801561029c57506001600160a01b0384163b155b156102c557604051639996b31560e01b81526001600160a01b0385166004820152602401610152565b50805b9392505050565b8051156102df5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b80516001600160a01b038116811461030e575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f5b83811015610341578181015183820152602001610329565b50505f910152565b5f5f5f6060848603121561035b575f5ffd5b610364846102f8565b9250610372602085016102f8565b60408501519092506001600160401b0381111561038d575f5ffd5b8401601f8101861361039d575f5ffd5b80516001600160401b038111156103b6576103b6610313565b604051601f8201601f19908116603f011681016001600160401b03811182821017156103e4576103e4610313565b6040528181528282016020018810156103fb575f5ffd5b61040c826020830160208601610327565b8093505050509250925092565b5f825161042a818460208701610327565b9190910192915050565b60805161048161044b5f395f601001526104815ff3fe608060405261000c61000e565b005b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163303610081575f357fffffffff000000000000000000000000000000000000000000000000000000001663278f794360e11b1461007957610077610085565b565b610077610095565b6100775b6100776100906100c3565b6100fa565b5f806100a43660048184610313565b8101906100b1919061034e565b915091506100bf8282610118565b5050565b5f6100f57f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b905090565b365f5f375f5f365f845af43d5f5f3e808015610114573d5ff35b3d5ffd5b61012182610172565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561016a5761016582826101fa565b505050565b6100bf61026c565b806001600160a01b03163b5f036101ac57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051610216919061041f565b5f60405180830381855af49150503d805f811461024e576040519150601f19603f3d011682016040523d82523d5f602084013e610253565b606091505b509150915061026385838361028b565b95945050505050565b34156100775760405163b398979f60e01b815260040160405180910390fd5b6060826102a05761029b826102ea565b6102e3565b81511580156102b757506001600160a01b0384163b155b156102e057604051639996b31560e01b81526001600160a01b03851660048201526024016101a3565b50805b9392505050565b8051156102fa5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b5f5f85851115610321575f5ffd5b8386111561032d575f5ffd5b5050820193919092039150565b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561035f575f5ffd5b82356001600160a01b0381168114610375575f5ffd5b9150602083013567ffffffffffffffff811115610390575f5ffd5b8301601f810185136103a0575f5ffd5b803567ffffffffffffffff8111156103ba576103ba61033a565b604051601f8201601f19908116603f0116810167ffffffffffffffff811182821017156103e9576103e961033a565b604052818152828201602001871015610400575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f82515f5b8181101561043e5760208186018101518583015201610424565b505f92019182525091905056fea2646970667358221220efd9e2b9a4d2427a35e1c3911d4a0a175cf3cfcf1a3a79870ec7a7d7ed34292f64736f6c634300081c0033b53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103",
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

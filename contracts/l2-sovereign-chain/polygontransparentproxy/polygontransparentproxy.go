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
	Bin: "0x60a06040526040516200091d3803806200091d833981016040819052620000269162000375565b828162000034828262000060565b50506001600160a01b038216608052620000576200005160805190565b620000c5565b5050506200046c565b6200006b8262000136565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115620000b757620000b28282620001b5565b505050565b620000c16200022e565b5050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f620001065f80516020620008fd833981519152546001600160a01b031690565b604080516001600160a01b03928316815291841660208301520160405180910390a1620001338162000250565b50565b806001600160a01b03163b5f036200017157604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5b80546001600160a01b0319166001600160a01b039290921691909117905550565b60605f80846001600160a01b031684604051620001d391906200044f565b5f60405180830381855af49150503d805f81146200020d576040519150601f19603f3d011682016040523d82523d5f602084013e62000212565b606091505b5090925090506200022585838362000291565b95945050505050565b34156200024e5760405163b398979f60e01b815260040160405180910390fd5b565b6001600160a01b0381166200027b57604051633173bdd160e11b81525f600482015260240162000168565b805f80516020620008fd83398151915262000194565b606082620002aa57620002a482620002f7565b620002f0565b8151158015620002c257506001600160a01b0384163b155b15620002ed57604051639996b31560e01b81526001600160a01b038516600482015260240162000168565b50805b9392505050565b805115620003085780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b80516001600160a01b038116811462000338575f80fd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f5b838110156200036d57818101518382015260200162000353565b50505f910152565b5f805f6060848603121562000388575f80fd5b620003938462000321565b9250620003a36020850162000321565b60408501519092506001600160401b0380821115620003c0575f80fd5b818601915086601f830112620003d4575f80fd5b815181811115620003e957620003e96200033d565b604051601f8201601f19908116603f011681019083821181831017156200041457620004146200033d565b816040528281528960208487010111156200042d575f80fd5b6200044083602083016020880162000351565b80955050505050509250925092565b5f82516200046281846020870162000351565b9190910192915050565b608051610479620004845f395f601001526104795ff3fe608060405261000c61000e565b005b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163303610081575f357fffffffff000000000000000000000000000000000000000000000000000000001663278f794360e11b1461007957610077610085565b565b610077610095565b6100775b6100776100906100c3565b6100fa565b5f806100a43660048184610313565b8101906100b1919061034e565b915091506100bf8282610118565b5050565b5f6100f57f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b905090565b365f80375f80365f845af43d5f803e808015610114573d5ff35b3d5ffd5b61012182610172565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561016a5761016582826101fa565b505050565b6100bf61026c565b806001600160a01b03163b5f036101ac57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b60605f80846001600160a01b0316846040516102169190610417565b5f60405180830381855af49150503d805f811461024e576040519150601f19603f3d011682016040523d82523d5f602084013e610253565b606091505b509150915061026385838361028b565b95945050505050565b34156100775760405163b398979f60e01b815260040160405180910390fd5b6060826102a05761029b826102ea565b6102e3565b81511580156102b757506001600160a01b0384163b155b156102e057604051639996b31560e01b81526001600160a01b03851660048201526024016101a3565b50805b9392505050565b8051156102fa5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b5f8085851115610321575f80fd5b8386111561032d575f80fd5b5050820193919092039150565b634e487b7160e01b5f52604160045260245ffd5b5f806040838503121561035f575f80fd5b82356001600160a01b0381168114610375575f80fd5b9150602083013567ffffffffffffffff80821115610391575f80fd5b818501915085601f8301126103a4575f80fd5b8135818111156103b6576103b661033a565b604051601f8201601f19908116603f011681019083821181831017156103de576103de61033a565b816040528281528860208487010111156103f6575f80fd5b826020860160208301375f6020848301015280955050505050509250929050565b5f82515f5b81811015610436576020818601810151858301520161041c565b505f92019182525091905056fea264697066735822122021e23b4641727aa4aa8fd2d3ef7960a43cab420d4c8632503bedc3eb2d30690764736f6c63430008140033b53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103",
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

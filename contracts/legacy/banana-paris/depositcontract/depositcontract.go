// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package depositcontract

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

// DepositcontractMetaData contains all meta data concerning the Depositcontract contract.
var DepositcontractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDepositRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"name\":\"getLeafValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"verifyMerkleProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610500806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80632dfdf0b5146100515780633ae050471461006d5780633e19704314610075578063fb57083414610158575b600080fd5b61005a60535481565b6040519081526020015b60405180910390f35b61005a61017b565b61005a61008336600461036e565b6040517fff0000000000000000000000000000000000000000000000000000000000000060f889901b1660208201527fffffffff0000000000000000000000000000000000000000000000000000000060e088811b821660218401527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606089811b821660258601529188901b909216603984015285901b16603d8201526051810183905260718101829052600090609101604051602081830303815290604052805190602001209050979650505050505050565b61016b6101663660046103ee565b610258565b6040519015158152602001610064565b605354600090819081805b602081101561024f578083901c6001166001036101e357603381602081106101b0576101b061043c565b01546040805160208101929092528101859052606001604051602081830303815290604052805190602001209350610210565b60408051602081018690529081018390526060016040516020818303038152906040528051906020012093505b604080516020810184905290810183905260600160405160208183030381529060405280519060200120915080806102479061046b565b915050610186565b50919392505050565b600084815b602081101561032557600163ffffffff8616821c811690036102c85785816020811061028b5761028b61043c565b6020020135826040516020016102ab929190918252602082015260400190565b604051602081830303815290604052805190602001209150610313565b818682602081106102db576102db61043c565b60200201356040516020016102fa929190918252602082015260400190565b6040516020818303038152906040528051906020012091505b8061031d8161046b565b91505061025d565b50909114949350505050565b803563ffffffff8116811461034557600080fd5b919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461034557600080fd5b600080600080600080600060e0888a03121561038957600080fd5b873560ff8116811461039a57600080fd5b96506103a860208901610331565b95506103b66040890161034a565b94506103c460608901610331565b93506103d26080890161034a565b925060a0880135915060c0880135905092959891949750929550565b600080600080610460858703121561040557600080fd5b8435935061042085018681111561041b57600080fd5b60208601935061042a81610331565b94979396509394610440013593505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036104c3577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b506001019056fea2646970667358221220532c5ab1291ad947b8bd52fb2b9454c6df16b34daf39f766d8aca6af4256acec64736f6c63430008140033",
}

// DepositcontractABI is the input ABI used to generate the binding from.
// Deprecated: Use DepositcontractMetaData.ABI instead.
var DepositcontractABI = DepositcontractMetaData.ABI

// DepositcontractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DepositcontractMetaData.Bin instead.
var DepositcontractBin = DepositcontractMetaData.Bin

// DeployDepositcontract deploys a new Ethereum contract, binding an instance of Depositcontract to it.
func DeployDepositcontract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Depositcontract, error) {
	parsed, err := DepositcontractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DepositcontractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Depositcontract{DepositcontractCaller: DepositcontractCaller{contract: contract}, DepositcontractTransactor: DepositcontractTransactor{contract: contract}, DepositcontractFilterer: DepositcontractFilterer{contract: contract}}, nil
}

// Depositcontract is an auto generated Go binding around an Ethereum contract.
type Depositcontract struct {
	DepositcontractCaller     // Read-only binding to the contract
	DepositcontractTransactor // Write-only binding to the contract
	DepositcontractFilterer   // Log filterer for contract events
}

// DepositcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositcontractSession struct {
	Contract     *Depositcontract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DepositcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositcontractCallerSession struct {
	Contract *DepositcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DepositcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositcontractTransactorSession struct {
	Contract     *DepositcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DepositcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositcontractRaw struct {
	Contract *Depositcontract // Generic contract binding to access the raw methods on
}

// DepositcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositcontractCallerRaw struct {
	Contract *DepositcontractCaller // Generic read-only contract binding to access the raw methods on
}

// DepositcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositcontractTransactorRaw struct {
	Contract *DepositcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDepositcontract creates a new instance of Depositcontract, bound to a specific deployed contract.
func NewDepositcontract(address common.Address, backend bind.ContractBackend) (*Depositcontract, error) {
	contract, err := bindDepositcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Depositcontract{DepositcontractCaller: DepositcontractCaller{contract: contract}, DepositcontractTransactor: DepositcontractTransactor{contract: contract}, DepositcontractFilterer: DepositcontractFilterer{contract: contract}}, nil
}

// NewDepositcontractCaller creates a new read-only instance of Depositcontract, bound to a specific deployed contract.
func NewDepositcontractCaller(address common.Address, caller bind.ContractCaller) (*DepositcontractCaller, error) {
	contract, err := bindDepositcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositcontractCaller{contract: contract}, nil
}

// NewDepositcontractTransactor creates a new write-only instance of Depositcontract, bound to a specific deployed contract.
func NewDepositcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositcontractTransactor, error) {
	contract, err := bindDepositcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositcontractTransactor{contract: contract}, nil
}

// NewDepositcontractFilterer creates a new log filterer instance of Depositcontract, bound to a specific deployed contract.
func NewDepositcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositcontractFilterer, error) {
	contract, err := bindDepositcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositcontractFilterer{contract: contract}, nil
}

// bindDepositcontract binds a generic wrapper to an already deployed contract.
func bindDepositcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DepositcontractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Depositcontract *DepositcontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Depositcontract.Contract.DepositcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Depositcontract *DepositcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositcontract.Contract.DepositcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Depositcontract *DepositcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Depositcontract.Contract.DepositcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Depositcontract *DepositcontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Depositcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Depositcontract *DepositcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Depositcontract *DepositcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Depositcontract.Contract.contract.Transact(opts, method, params...)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Depositcontract *DepositcontractCaller) DepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Depositcontract.contract.Call(opts, &out, "depositCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Depositcontract *DepositcontractSession) DepositCount() (*big.Int, error) {
	return _Depositcontract.Contract.DepositCount(&_Depositcontract.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Depositcontract *DepositcontractCallerSession) DepositCount() (*big.Int, error) {
	return _Depositcontract.Contract.DepositCount(&_Depositcontract.CallOpts)
}

// GetDepositRoot is a free data retrieval call binding the contract method 0x3ae05047.
//
// Solidity: function getDepositRoot() view returns(bytes32)
func (_Depositcontract *DepositcontractCaller) GetDepositRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Depositcontract.contract.Call(opts, &out, "getDepositRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDepositRoot is a free data retrieval call binding the contract method 0x3ae05047.
//
// Solidity: function getDepositRoot() view returns(bytes32)
func (_Depositcontract *DepositcontractSession) GetDepositRoot() ([32]byte, error) {
	return _Depositcontract.Contract.GetDepositRoot(&_Depositcontract.CallOpts)
}

// GetDepositRoot is a free data retrieval call binding the contract method 0x3ae05047.
//
// Solidity: function getDepositRoot() view returns(bytes32)
func (_Depositcontract *DepositcontractCallerSession) GetDepositRoot() ([32]byte, error) {
	return _Depositcontract.Contract.GetDepositRoot(&_Depositcontract.CallOpts)
}

// GetLeafValue is a free data retrieval call binding the contract method 0x3e197043.
//
// Solidity: function getLeafValue(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes32 metadataHash) pure returns(bytes32)
func (_Depositcontract *DepositcontractCaller) GetLeafValue(opts *bind.CallOpts, leafType uint8, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadataHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Depositcontract.contract.Call(opts, &out, "getLeafValue", leafType, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadataHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLeafValue is a free data retrieval call binding the contract method 0x3e197043.
//
// Solidity: function getLeafValue(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes32 metadataHash) pure returns(bytes32)
func (_Depositcontract *DepositcontractSession) GetLeafValue(leafType uint8, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadataHash [32]byte) ([32]byte, error) {
	return _Depositcontract.Contract.GetLeafValue(&_Depositcontract.CallOpts, leafType, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadataHash)
}

// GetLeafValue is a free data retrieval call binding the contract method 0x3e197043.
//
// Solidity: function getLeafValue(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes32 metadataHash) pure returns(bytes32)
func (_Depositcontract *DepositcontractCallerSession) GetLeafValue(leafType uint8, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadataHash [32]byte) ([32]byte, error) {
	return _Depositcontract.Contract.GetLeafValue(&_Depositcontract.CallOpts, leafType, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadataHash)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Depositcontract *DepositcontractCaller) VerifyMerkleProof(opts *bind.CallOpts, leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	var out []interface{}
	err := _Depositcontract.contract.Call(opts, &out, "verifyMerkleProof", leafHash, smtProof, index, root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Depositcontract *DepositcontractSession) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	return _Depositcontract.Contract.VerifyMerkleProof(&_Depositcontract.CallOpts, leafHash, smtProof, index, root)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Depositcontract *DepositcontractCallerSession) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	return _Depositcontract.Contract.VerifyMerkleProof(&_Depositcontract.CallOpts, leafHash, smtProof, index, root)
}

// DepositcontractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Depositcontract contract.
type DepositcontractInitializedIterator struct {
	Event *DepositcontractInitialized // Event containing the contract specifics and raw log

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
func (it *DepositcontractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositcontractInitialized)
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
		it.Event = new(DepositcontractInitialized)
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
func (it *DepositcontractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositcontractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositcontractInitialized represents a Initialized event raised by the Depositcontract contract.
type DepositcontractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Depositcontract *DepositcontractFilterer) FilterInitialized(opts *bind.FilterOpts) (*DepositcontractInitializedIterator, error) {

	logs, sub, err := _Depositcontract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DepositcontractInitializedIterator{contract: _Depositcontract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Depositcontract *DepositcontractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DepositcontractInitialized) (event.Subscription, error) {

	logs, sub, err := _Depositcontract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositcontractInitialized)
				if err := _Depositcontract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Depositcontract *DepositcontractFilterer) ParseInitialized(log types.Log) (*DepositcontractInitialized, error) {
	event := new(DepositcontractInitialized)
	if err := _Depositcontract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

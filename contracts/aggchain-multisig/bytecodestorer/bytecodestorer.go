// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bytecodestorer

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

// BytecodestorerMetaData contains all meta data concerning the Bytecodestorer contract.
var BytecodestorerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"INIT_BYTECODE_TRANSPARENT_PROXY\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50610fb38061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610029575f3560e01c8063c514f24e1461002d575b5f5ffd5b61003561004b565b604051610042919061006a565b60405180910390f35b60405180610f000160405280610ec881526020016100b6610ec8913981565b602081525f82518060208401525f5b818110156100965760208186018101516040868401015201610079565b505f604082850101526040601f19601f8301168401019150509291505056fe60806040819052631d97f74d60e11b81523390633b2fee9a90608490602090600481865afa158015610033573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100579190610433565b604080515f80825260208201909252905061007382825f6100e2565b50506100dd336001600160a01b03166338b8fbbb6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100b4573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100d89190610433565b61010d565b6104c8565b6100eb8361017a565b5f825111806100f75750805b156101085761010683836101b9565b505b505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61014c5f516020610e815f395f51905f52546001600160a01b031690565b604080516001600160a01b03928316815291841660208301520160405180910390a1610177816101e5565b50565b61018381610280565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606101de8383604051806060016040528060278152602001610ea160279139610314565b9392505050565b6001600160a01b03811661024f5760405162461bcd60e51b815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b805f516020610e815f395f51905f525b80546001600160a01b0319166001600160a01b039290921691909117905550565b6001600160a01b0381163b6102ed5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610246565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc61025f565b60605f5f856001600160a01b031685604051610330919061047b565b5f60405180830381855af49150503d805f8114610368576040519150601f19603f3d011682016040523d82523d5f602084013e61036d565b606091505b50909250905061037f86838387610389565b9695505050505050565b606083156103f75782515f036103f0576001600160a01b0385163b6103f05760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610246565b5081610401565b6104018383610409565b949350505050565b8151156104195781518083602001fd5b8060405162461bcd60e51b81526004016102469190610496565b5f60208284031215610443575f5ffd5b81516001600160a01b03811681146101de575f5ffd5b5f5b8381101561047357818101518382015260200161045b565b50505f910152565b5f825161048c818460208701610459565b9190910192915050565b602081525f82518060208401526104b4816040850160208701610459565b601f01601f19169190910160400192915050565b6109ac806104d55f395ff3fe60806040526004361061005d575f3560e01c80635c60da1b116100425780635c60da1b146100a65780638f283970146100e3578063f851a440146101025761006c565b80633659cfe6146100745780634f1ef286146100935761006c565b3661006c5761006a610116565b005b61006a610116565b34801561007f575f5ffd5b5061006a61008e366004610854565b610130565b61006a6100a136600461086d565b610178565b3480156100b1575f5ffd5b506100ba6101eb565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b3480156100ee575f5ffd5b5061006a6100fd366004610854565b610228565b34801561010d575f5ffd5b506100ba610255565b61011e610282565b61012e610129610359565b610362565b565b610138610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101705761016d8160405180602001604052805f8152505f6103bf565b50565b61016d610116565b610180610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101e3576101de8383838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250600192506103bf915050565b505050565b6101de610116565b5f6101f4610380565b73ffffffffffffffffffffffffffffffffffffffff16330361021d57610218610359565b905090565b610225610116565b90565b610230610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101705761016d816103e9565b5f61025e610380565b73ffffffffffffffffffffffffffffffffffffffff16330361021d57610218610380565b61028a610380565b73ffffffffffffffffffffffffffffffffffffffff16330361012e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604260248201527f5472616e73706172656e745570677261646561626c6550726f78793a2061646d60448201527f696e2063616e6e6f742066616c6c6261636b20746f2070726f7879207461726760648201527f6574000000000000000000000000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b5f61021861044a565b365f5f375f5f365f845af43d5f5f3e80801561037c573d5ff35b3d5ffd5b5f7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b5473ffffffffffffffffffffffffffffffffffffffff16919050565b6103c883610471565b5f825111806103d45750805b156101de576103e383836104bd565b50505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f610412610380565b6040805173ffffffffffffffffffffffffffffffffffffffff928316815291841660208301520160405180910390a161016d816104e9565b5f7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6103a3565b61047a816105f5565b60405173ffffffffffffffffffffffffffffffffffffffff8216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606104e28383604051806060016040528060278152602001610979602791396106c0565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff811661058c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610350565b807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b73ffffffffffffffffffffffffffffffffffffffff81163b610699576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201527f6f74206120636f6e7472616374000000000000000000000000000000000000006064820152608401610350565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6105af565b60605f5f8573ffffffffffffffffffffffffffffffffffffffff16856040516106e9919061090d565b5f60405180830381855af49150503d805f8114610721576040519150601f19603f3d011682016040523d82523d5f602084013e610726565b606091505b509150915061073786838387610741565b9695505050505050565b606083156107d65782515f036107cf5773ffffffffffffffffffffffffffffffffffffffff85163b6107cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610350565b50816107e0565b6107e083836107e8565b949350505050565b8151156107f85781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103509190610928565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f575f5ffd5b919050565b5f60208284031215610864575f5ffd5b6104e28261082c565b5f5f5f6040848603121561087f575f5ffd5b6108888461082c565b9250602084013567ffffffffffffffff8111156108a3575f5ffd5b8401601f810186136108b3575f5ffd5b803567ffffffffffffffff8111156108c9575f5ffd5b8660208284010111156108da575f5ffd5b939660209190910195509293505050565b5f5b838110156109055781810151838201526020016108ed565b50505f910152565b5f825161091e8184602087016108eb565b9190910192915050565b602081525f82518060208401526109468160408501602087016108eb565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a164736f6c634300081c000ab53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220959dc41fc91b71aeb8e8530841d9366109c0c669513f9b502f94e79b14287a3164736f6c634300081c0033",
}

// BytecodestorerABI is the input ABI used to generate the binding from.
// Deprecated: Use BytecodestorerMetaData.ABI instead.
var BytecodestorerABI = BytecodestorerMetaData.ABI

// BytecodestorerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BytecodestorerMetaData.Bin instead.
var BytecodestorerBin = BytecodestorerMetaData.Bin

// DeployBytecodestorer deploys a new Ethereum contract, binding an instance of Bytecodestorer to it.
func DeployBytecodestorer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bytecodestorer, error) {
	parsed, err := BytecodestorerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BytecodestorerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bytecodestorer{BytecodestorerCaller: BytecodestorerCaller{contract: contract}, BytecodestorerTransactor: BytecodestorerTransactor{contract: contract}, BytecodestorerFilterer: BytecodestorerFilterer{contract: contract}}, nil
}

// Bytecodestorer is an auto generated Go binding around an Ethereum contract.
type Bytecodestorer struct {
	BytecodestorerCaller     // Read-only binding to the contract
	BytecodestorerTransactor // Write-only binding to the contract
	BytecodestorerFilterer   // Log filterer for contract events
}

// BytecodestorerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BytecodestorerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytecodestorerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BytecodestorerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytecodestorerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BytecodestorerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytecodestorerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BytecodestorerSession struct {
	Contract     *Bytecodestorer   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytecodestorerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BytecodestorerCallerSession struct {
	Contract *BytecodestorerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BytecodestorerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BytecodestorerTransactorSession struct {
	Contract     *BytecodestorerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BytecodestorerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BytecodestorerRaw struct {
	Contract *Bytecodestorer // Generic contract binding to access the raw methods on
}

// BytecodestorerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BytecodestorerCallerRaw struct {
	Contract *BytecodestorerCaller // Generic read-only contract binding to access the raw methods on
}

// BytecodestorerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BytecodestorerTransactorRaw struct {
	Contract *BytecodestorerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBytecodestorer creates a new instance of Bytecodestorer, bound to a specific deployed contract.
func NewBytecodestorer(address common.Address, backend bind.ContractBackend) (*Bytecodestorer, error) {
	contract, err := bindBytecodestorer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bytecodestorer{BytecodestorerCaller: BytecodestorerCaller{contract: contract}, BytecodestorerTransactor: BytecodestorerTransactor{contract: contract}, BytecodestorerFilterer: BytecodestorerFilterer{contract: contract}}, nil
}

// NewBytecodestorerCaller creates a new read-only instance of Bytecodestorer, bound to a specific deployed contract.
func NewBytecodestorerCaller(address common.Address, caller bind.ContractCaller) (*BytecodestorerCaller, error) {
	contract, err := bindBytecodestorer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BytecodestorerCaller{contract: contract}, nil
}

// NewBytecodestorerTransactor creates a new write-only instance of Bytecodestorer, bound to a specific deployed contract.
func NewBytecodestorerTransactor(address common.Address, transactor bind.ContractTransactor) (*BytecodestorerTransactor, error) {
	contract, err := bindBytecodestorer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BytecodestorerTransactor{contract: contract}, nil
}

// NewBytecodestorerFilterer creates a new log filterer instance of Bytecodestorer, bound to a specific deployed contract.
func NewBytecodestorerFilterer(address common.Address, filterer bind.ContractFilterer) (*BytecodestorerFilterer, error) {
	contract, err := bindBytecodestorer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BytecodestorerFilterer{contract: contract}, nil
}

// bindBytecodestorer binds a generic wrapper to an already deployed contract.
func bindBytecodestorer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BytecodestorerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bytecodestorer *BytecodestorerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bytecodestorer.Contract.BytecodestorerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bytecodestorer *BytecodestorerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bytecodestorer.Contract.BytecodestorerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bytecodestorer *BytecodestorerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bytecodestorer.Contract.BytecodestorerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bytecodestorer *BytecodestorerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bytecodestorer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bytecodestorer *BytecodestorerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bytecodestorer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bytecodestorer *BytecodestorerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bytecodestorer.Contract.contract.Transact(opts, method, params...)
}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() view returns(bytes)
func (_Bytecodestorer *BytecodestorerCaller) INITBYTECODETRANSPARENTPROXY(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Bytecodestorer.contract.Call(opts, &out, "INIT_BYTECODE_TRANSPARENT_PROXY")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() view returns(bytes)
func (_Bytecodestorer *BytecodestorerSession) INITBYTECODETRANSPARENTPROXY() ([]byte, error) {
	return _Bytecodestorer.Contract.INITBYTECODETRANSPARENTPROXY(&_Bytecodestorer.CallOpts)
}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() view returns(bytes)
func (_Bytecodestorer *BytecodestorerCallerSession) INITBYTECODETRANSPARENTPROXY() ([]byte, error) {
	return _Bytecodestorer.Contract.INITBYTECODETRANSPARENTPROXY(&_Bytecodestorer.CallOpts)
}

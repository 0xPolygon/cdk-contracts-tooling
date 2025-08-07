// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonzkevmdeployer

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

// PolygonzkevmdeployerMetaData contains all meta data concerning the Polygonzkevmdeployer contract.
var PolygonzkevmdeployerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"FunctionCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newContractAddress\",\"type\":\"address\"}],\"name\":\"NewDeterministicDeployment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"initBytecode\",\"type\":\"bytes\"}],\"name\":\"deployDeterministic\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"initBytecode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"dataCall\",\"type\":\"bytes\"}],\"name\":\"deployDeterministicAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataCall\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"functionCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bytecodeHash\",\"type\":\"bytes32\"}],\"name\":\"predictDeterministicAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610c71380380610c7183398101604081905261002f91610097565b61003833610047565b61004181610047565b506100c7565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100a957600080fd5b81516001600160a01b03811681146100c057600080fd5b9392505050565b610b9b806100d66000396000f3fe6080604052600436106100705760003560e01c8063715018a61161004e578063715018a6146100e65780638da5cb5b146100fb578063e11ae6cb14610126578063f2fde38b1461013957600080fd5b80632b79805a146100755780634a94d4871461008a5780636d07dbf81461009d575b600080fd5b610088610083366004610927565b610159565b005b6100886100983660046109c7565b6101cb565b3480156100a957600080fd5b506100bd6100b8366004610a1e565b61020d565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b3480156100f257600080fd5b50610088610220565b34801561010757600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff166100bd565b610088610134366004610a40565b610234565b34801561014557600080fd5b50610088610154366004610a90565b61029b565b610161610357565b600061016e8585856103d8565b905061017a8183610537565b5060405173ffffffffffffffffffffffffffffffffffffffff821681527fba82f25fed02cd2a23d9f5d11c2ef588d22af5437cbf23bfe61d87257c480e4c9060200160405180910390a15050505050565b6101d3610357565b6101de83838361057b565b506040517f25adb19089b6a549831a273acdf7908cff8b7ee5f551f8d1d37996cf01c5df5b90600090a1505050565b600061021983836105a9565b9392505050565b610228610357565b61023260006105b6565b565b61023c610357565b60006102498484846103d8565b60405173ffffffffffffffffffffffffffffffffffffffff821681529091507fba82f25fed02cd2a23d9f5d11c2ef588d22af5437cbf23bfe61d87257c480e4c9060200160405180910390a150505050565b6102a3610357565b73ffffffffffffffffffffffffffffffffffffffff811661034b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b610354816105b6565b50565b60005473ffffffffffffffffffffffffffffffffffffffff163314610232576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610342565b600083471015610444576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f437265617465323a20696e73756666696369656e742062616c616e63650000006044820152606401610342565b81516000036104af576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f437265617465323a2062797465636f6465206c656e677468206973207a65726f6044820152606401610342565b8282516020840186f5905073ffffffffffffffffffffffffffffffffffffffff8116610219576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f437265617465323a204661696c6564206f6e206465706c6f79000000000000006044820152606401610342565b6060610219838360006040518060400160405280601e81526020017f416464726573733a206c6f772d6c6576656c2063616c6c206661696c6564000081525061062b565b60606105a1848484604051806060016040528060298152602001610b3d6029913961062b565b949350505050565b6000610219838330610744565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6060824710156106bd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610342565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516106e69190610acf565b60006040518083038185875af1925050503d8060008114610723576040519150601f19603f3d011682016040523d82523d6000602084013e610728565b606091505b50915091506107398783838761076e565b979650505050505050565b6000604051836040820152846020820152828152600b8101905060ff815360559020949350505050565b606083156108045782516000036107fd5773ffffffffffffffffffffffffffffffffffffffff85163b6107fd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610342565b50816105a1565b6105a183838151156108195781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103429190610aeb565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261088d57600080fd5b813567ffffffffffffffff808211156108a8576108a861084d565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019082821181831017156108ee576108ee61084d565b8160405283815286602085880101111561090757600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000806080858703121561093d57600080fd5b8435935060208501359250604085013567ffffffffffffffff8082111561096357600080fd5b61096f8883890161087c565b9350606087013591508082111561098557600080fd5b506109928782880161087c565b91505092959194509250565b803573ffffffffffffffffffffffffffffffffffffffff811681146109c257600080fd5b919050565b6000806000606084860312156109dc57600080fd5b6109e58461099e565b9250602084013567ffffffffffffffff811115610a0157600080fd5b610a0d8682870161087c565b925050604084013590509250925092565b60008060408385031215610a3157600080fd5b50508035926020909101359150565b600080600060608486031215610a5557600080fd5b8335925060208401359150604084013567ffffffffffffffff811115610a7a57600080fd5b610a868682870161087c565b9150509250925092565b600060208284031215610aa257600080fd5b6102198261099e565b60005b83811015610ac6578181015183820152602001610aae565b50506000910152565b60008251610ae1818460208701610aab565b9190910192915050565b6020815260008251806020840152610b0a816040850160208701610aab565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016040019291505056fe416464726573733a206c6f772d6c6576656c2063616c6c20776974682076616c7565206661696c6564a2646970667358221220964619cee0e0baf94c6f8763f013be157da5d54c89e5cff4a8caf4266e13f13a64736f6c63430008140033",
}

// PolygonzkevmdeployerABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonzkevmdeployerMetaData.ABI instead.
var PolygonzkevmdeployerABI = PolygonzkevmdeployerMetaData.ABI

// PolygonzkevmdeployerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygonzkevmdeployerMetaData.Bin instead.
var PolygonzkevmdeployerBin = PolygonzkevmdeployerMetaData.Bin

// DeployPolygonzkevmdeployer deploys a new Ethereum contract, binding an instance of Polygonzkevmdeployer to it.
func DeployPolygonzkevmdeployer(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *Polygonzkevmdeployer, error) {
	parsed, err := PolygonzkevmdeployerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygonzkevmdeployerBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonzkevmdeployer{PolygonzkevmdeployerCaller: PolygonzkevmdeployerCaller{contract: contract}, PolygonzkevmdeployerTransactor: PolygonzkevmdeployerTransactor{contract: contract}, PolygonzkevmdeployerFilterer: PolygonzkevmdeployerFilterer{contract: contract}}, nil
}

// Polygonzkevmdeployer is an auto generated Go binding around an Ethereum contract.
type Polygonzkevmdeployer struct {
	PolygonzkevmdeployerCaller     // Read-only binding to the contract
	PolygonzkevmdeployerTransactor // Write-only binding to the contract
	PolygonzkevmdeployerFilterer   // Log filterer for contract events
}

// PolygonzkevmdeployerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonzkevmdeployerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmdeployerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonzkevmdeployerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmdeployerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonzkevmdeployerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmdeployerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonzkevmdeployerSession struct {
	Contract     *Polygonzkevmdeployer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PolygonzkevmdeployerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonzkevmdeployerCallerSession struct {
	Contract *PolygonzkevmdeployerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// PolygonzkevmdeployerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonzkevmdeployerTransactorSession struct {
	Contract     *PolygonzkevmdeployerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// PolygonzkevmdeployerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonzkevmdeployerRaw struct {
	Contract *Polygonzkevmdeployer // Generic contract binding to access the raw methods on
}

// PolygonzkevmdeployerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonzkevmdeployerCallerRaw struct {
	Contract *PolygonzkevmdeployerCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonzkevmdeployerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonzkevmdeployerTransactorRaw struct {
	Contract *PolygonzkevmdeployerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonzkevmdeployer creates a new instance of Polygonzkevmdeployer, bound to a specific deployed contract.
func NewPolygonzkevmdeployer(address common.Address, backend bind.ContractBackend) (*Polygonzkevmdeployer, error) {
	contract, err := bindPolygonzkevmdeployer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmdeployer{PolygonzkevmdeployerCaller: PolygonzkevmdeployerCaller{contract: contract}, PolygonzkevmdeployerTransactor: PolygonzkevmdeployerTransactor{contract: contract}, PolygonzkevmdeployerFilterer: PolygonzkevmdeployerFilterer{contract: contract}}, nil
}

// NewPolygonzkevmdeployerCaller creates a new read-only instance of Polygonzkevmdeployer, bound to a specific deployed contract.
func NewPolygonzkevmdeployerCaller(address common.Address, caller bind.ContractCaller) (*PolygonzkevmdeployerCaller, error) {
	contract, err := bindPolygonzkevmdeployer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmdeployerCaller{contract: contract}, nil
}

// NewPolygonzkevmdeployerTransactor creates a new write-only instance of Polygonzkevmdeployer, bound to a specific deployed contract.
func NewPolygonzkevmdeployerTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonzkevmdeployerTransactor, error) {
	contract, err := bindPolygonzkevmdeployer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmdeployerTransactor{contract: contract}, nil
}

// NewPolygonzkevmdeployerFilterer creates a new log filterer instance of Polygonzkevmdeployer, bound to a specific deployed contract.
func NewPolygonzkevmdeployerFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonzkevmdeployerFilterer, error) {
	contract, err := bindPolygonzkevmdeployer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmdeployerFilterer{contract: contract}, nil
}

// bindPolygonzkevmdeployer binds a generic wrapper to an already deployed contract.
func bindPolygonzkevmdeployer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonzkevmdeployerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmdeployer *PolygonzkevmdeployerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmdeployer.Contract.PolygonzkevmdeployerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmdeployer *PolygonzkevmdeployerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmdeployer.Contract.PolygonzkevmdeployerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmdeployer *PolygonzkevmdeployerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmdeployer.Contract.PolygonzkevmdeployerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmdeployer *PolygonzkevmdeployerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmdeployer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmdeployer *PolygonzkevmdeployerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmdeployer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmdeployer *PolygonzkevmdeployerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmdeployer.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Polygonzkevmdeployer *PolygonzkevmdeployerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmdeployer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Polygonzkevmdeployer *PolygonzkevmdeployerSession) Owner() (common.Address, error) {
	return _Polygonzkevmdeployer.Contract.Owner(&_Polygonzkevmdeployer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Polygonzkevmdeployer *PolygonzkevmdeployerCallerSession) Owner() (common.Address, error) {
	return _Polygonzkevmdeployer.Contract.Owner(&_Polygonzkevmdeployer.CallOpts)
}

// PredictDeterministicAddress is a free data retrieval call binding the contract method 0x6d07dbf8.
//
// Solidity: function predictDeterministicAddress(bytes32 salt, bytes32 bytecodeHash) view returns(address)
func (_Polygonzkevmdeployer *PolygonzkevmdeployerCaller) PredictDeterministicAddress(opts *bind.CallOpts, salt [32]byte, bytecodeHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmdeployer.contract.Call(opts, &out, "predictDeterministicAddress", salt, bytecodeHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PredictDeterministicAddress is a free data retrieval call binding the contract method 0x6d07dbf8.
//
// Solidity: function predictDeterministicAddress(bytes32 salt, bytes32 bytecodeHash) view returns(address)
func (_Polygonzkevmdeployer *PolygonzkevmdeployerSession) PredictDeterministicAddress(salt [32]byte, bytecodeHash [32]byte) (common.Address, error) {
	return _Polygonzkevmdeployer.Contract.PredictDeterministicAddress(&_Polygonzkevmdeployer.CallOpts, salt, bytecodeHash)
}

// PredictDeterministicAddress is a free data retrieval call binding the contract method 0x6d07dbf8.
//
// Solidity: function predictDeterministicAddress(bytes32 salt, bytes32 bytecodeHash) view returns(address)
func (_Polygonzkevmdeployer *PolygonzkevmdeployerCallerSession) PredictDeterministicAddress(salt [32]byte, bytecodeHash [32]byte) (common.Address, error) {
	return _Polygonzkevmdeployer.Contract.PredictDeterministicAddress(&_Polygonzkevmdeployer.CallOpts, salt, bytecodeHash)
}

// DeployDeterministic is a paid mutator transaction binding the contract method 0xe11ae6cb.
//
// Solidity: function deployDeterministic(uint256 amount, bytes32 salt, bytes initBytecode) payable returns()
func (_Polygonzkevmdeployer *PolygonzkevmdeployerTransactor) DeployDeterministic(opts *bind.TransactOpts, amount *big.Int, salt [32]byte, initBytecode []byte) (*types.Transaction, error) {
	return _Polygonzkevmdeployer.contract.Transact(opts, "deployDeterministic", amount, salt, initBytecode)
}

// DeployDeterministic is a paid mutator transaction binding the contract method 0xe11ae6cb.
//
// Solidity: function deployDeterministic(uint256 amount, bytes32 salt, bytes initBytecode) payable returns()
func (_Polygonzkevmdeployer *PolygonzkevmdeployerSession) DeployDeterministic(amount *big.Int, salt [32]byte, initBytecode []byte) (*types.Transaction, error) {
	return _Polygonzkevmdeployer.Contract.DeployDeterministic(&_Polygonzkevmdeployer.TransactOpts, amount, salt, initBytecode)
}

// DeployDeterministic is a paid mutator transaction binding the contract method 0xe11ae6cb.
//
// Solidity: function deployDeterministic(uint256 amount, bytes32 salt, bytes initBytecode) payable returns()
func (_Polygonzkevmdeployer *PolygonzkevmdeployerTransactorSession) DeployDeterministic(amount *big.Int, salt [32]byte, initBytecode []byte) (*types.Transaction, error) {
	return _Polygonzkevmdeployer.Contract.DeployDeterministic(&_Polygonzkevmdeployer.TransactOpts, amount, salt, initBytecode)
}

// DeployDeterministicAndCall is a paid mutator transaction binding the contract method 0x2b79805a.
//
// Solidity: function deployDeterministicAndCall(uint256 amount, bytes32 salt, bytes initBytecode, bytes dataCall) payable returns()
func (_Polygonzkevmdeployer *PolygonzkevmdeployerTransactor) DeployDeterministicAndCall(opts *bind.TransactOpts, amount *big.Int, salt [32]byte, initBytecode []byte, dataCall []byte) (*types.Transaction, error) {
	return _Polygonzkevmdeployer.contract.Transact(opts, "deployDeterministicAndCall", amount, salt, initBytecode, dataCall)
}

// DeployDeterministicAndCall is a paid mutator transaction binding the contract method 0x2b79805a.
//
// Solidity: function deployDeterministicAndCall(uint256 amount, bytes32 salt, bytes initBytecode, bytes dataCall) payable returns()
func (_Polygonzkevmdeployer *PolygonzkevmdeployerSession) DeployDeterministicAndCall(amount *big.Int, salt [32]byte, initBytecode []byte, dataCall []byte) (*types.Transaction, error) {
	return _Polygonzkevmdeployer.Contract.DeployDeterministicAndCall(&_Polygonzkevmdeployer.TransactOpts, amount, salt, initBytecode, dataCall)
}

// DeployDeterministicAndCall is a paid mutator transaction binding the contract method 0x2b79805a.
//
// Solidity: function deployDeterministicAndCall(uint256 amount, bytes32 salt, bytes initBytecode, bytes dataCall) payable returns()
func (_Polygonzkevmdeployer *PolygonzkevmdeployerTransactorSession) DeployDeterministicAndCall(amount *big.Int, salt [32]byte, initBytecode []byte, dataCall []byte) (*types.Transaction, error) {
	return _Polygonzkevmdeployer.Contract.DeployDeterministicAndCall(&_Polygonzkevmdeployer.TransactOpts, amount, salt, initBytecode, dataCall)
}

// FunctionCall is a paid mutator transaction binding the contract method 0x4a94d487.
//
// Solidity: function functionCall(address targetAddress, bytes dataCall, uint256 amount) payable returns()
func (_Polygonzkevmdeployer *PolygonzkevmdeployerTransactor) FunctionCall(opts *bind.TransactOpts, targetAddress common.Address, dataCall []byte, amount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmdeployer.contract.Transact(opts, "functionCall", targetAddress, dataCall, amount)
}

// FunctionCall is a paid mutator transaction binding the contract method 0x4a94d487.
//
// Solidity: function functionCall(address targetAddress, bytes dataCall, uint256 amount) payable returns()
func (_Polygonzkevmdeployer *PolygonzkevmdeployerSession) FunctionCall(targetAddress common.Address, dataCall []byte, amount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmdeployer.Contract.FunctionCall(&_Polygonzkevmdeployer.TransactOpts, targetAddress, dataCall, amount)
}

// FunctionCall is a paid mutator transaction binding the contract method 0x4a94d487.
//
// Solidity: function functionCall(address targetAddress, bytes dataCall, uint256 amount) payable returns()
func (_Polygonzkevmdeployer *PolygonzkevmdeployerTransactorSession) FunctionCall(targetAddress common.Address, dataCall []byte, amount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmdeployer.Contract.FunctionCall(&_Polygonzkevmdeployer.TransactOpts, targetAddress, dataCall, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Polygonzkevmdeployer *PolygonzkevmdeployerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmdeployer.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Polygonzkevmdeployer *PolygonzkevmdeployerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Polygonzkevmdeployer.Contract.RenounceOwnership(&_Polygonzkevmdeployer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Polygonzkevmdeployer *PolygonzkevmdeployerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Polygonzkevmdeployer.Contract.RenounceOwnership(&_Polygonzkevmdeployer.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Polygonzkevmdeployer *PolygonzkevmdeployerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Polygonzkevmdeployer.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Polygonzkevmdeployer *PolygonzkevmdeployerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Polygonzkevmdeployer.Contract.TransferOwnership(&_Polygonzkevmdeployer.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Polygonzkevmdeployer *PolygonzkevmdeployerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Polygonzkevmdeployer.Contract.TransferOwnership(&_Polygonzkevmdeployer.TransactOpts, newOwner)
}

// PolygonzkevmdeployerFunctionCallIterator is returned from FilterFunctionCall and is used to iterate over the raw logs and unpacked data for FunctionCall events raised by the Polygonzkevmdeployer contract.
type PolygonzkevmdeployerFunctionCallIterator struct {
	Event *PolygonzkevmdeployerFunctionCall // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmdeployerFunctionCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmdeployerFunctionCall)
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
		it.Event = new(PolygonzkevmdeployerFunctionCall)
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
func (it *PolygonzkevmdeployerFunctionCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmdeployerFunctionCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmdeployerFunctionCall represents a FunctionCall event raised by the Polygonzkevmdeployer contract.
type PolygonzkevmdeployerFunctionCall struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFunctionCall is a free log retrieval operation binding the contract event 0x25adb19089b6a549831a273acdf7908cff8b7ee5f551f8d1d37996cf01c5df5b.
//
// Solidity: event FunctionCall()
func (_Polygonzkevmdeployer *PolygonzkevmdeployerFilterer) FilterFunctionCall(opts *bind.FilterOpts) (*PolygonzkevmdeployerFunctionCallIterator, error) {

	logs, sub, err := _Polygonzkevmdeployer.contract.FilterLogs(opts, "FunctionCall")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmdeployerFunctionCallIterator{contract: _Polygonzkevmdeployer.contract, event: "FunctionCall", logs: logs, sub: sub}, nil
}

// WatchFunctionCall is a free log subscription operation binding the contract event 0x25adb19089b6a549831a273acdf7908cff8b7ee5f551f8d1d37996cf01c5df5b.
//
// Solidity: event FunctionCall()
func (_Polygonzkevmdeployer *PolygonzkevmdeployerFilterer) WatchFunctionCall(opts *bind.WatchOpts, sink chan<- *PolygonzkevmdeployerFunctionCall) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmdeployer.contract.WatchLogs(opts, "FunctionCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmdeployerFunctionCall)
				if err := _Polygonzkevmdeployer.contract.UnpackLog(event, "FunctionCall", log); err != nil {
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

// ParseFunctionCall is a log parse operation binding the contract event 0x25adb19089b6a549831a273acdf7908cff8b7ee5f551f8d1d37996cf01c5df5b.
//
// Solidity: event FunctionCall()
func (_Polygonzkevmdeployer *PolygonzkevmdeployerFilterer) ParseFunctionCall(log types.Log) (*PolygonzkevmdeployerFunctionCall, error) {
	event := new(PolygonzkevmdeployerFunctionCall)
	if err := _Polygonzkevmdeployer.contract.UnpackLog(event, "FunctionCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmdeployerNewDeterministicDeploymentIterator is returned from FilterNewDeterministicDeployment and is used to iterate over the raw logs and unpacked data for NewDeterministicDeployment events raised by the Polygonzkevmdeployer contract.
type PolygonzkevmdeployerNewDeterministicDeploymentIterator struct {
	Event *PolygonzkevmdeployerNewDeterministicDeployment // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmdeployerNewDeterministicDeploymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmdeployerNewDeterministicDeployment)
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
		it.Event = new(PolygonzkevmdeployerNewDeterministicDeployment)
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
func (it *PolygonzkevmdeployerNewDeterministicDeploymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmdeployerNewDeterministicDeploymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmdeployerNewDeterministicDeployment represents a NewDeterministicDeployment event raised by the Polygonzkevmdeployer contract.
type PolygonzkevmdeployerNewDeterministicDeployment struct {
	NewContractAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewDeterministicDeployment is a free log retrieval operation binding the contract event 0xba82f25fed02cd2a23d9f5d11c2ef588d22af5437cbf23bfe61d87257c480e4c.
//
// Solidity: event NewDeterministicDeployment(address newContractAddress)
func (_Polygonzkevmdeployer *PolygonzkevmdeployerFilterer) FilterNewDeterministicDeployment(opts *bind.FilterOpts) (*PolygonzkevmdeployerNewDeterministicDeploymentIterator, error) {

	logs, sub, err := _Polygonzkevmdeployer.contract.FilterLogs(opts, "NewDeterministicDeployment")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmdeployerNewDeterministicDeploymentIterator{contract: _Polygonzkevmdeployer.contract, event: "NewDeterministicDeployment", logs: logs, sub: sub}, nil
}

// WatchNewDeterministicDeployment is a free log subscription operation binding the contract event 0xba82f25fed02cd2a23d9f5d11c2ef588d22af5437cbf23bfe61d87257c480e4c.
//
// Solidity: event NewDeterministicDeployment(address newContractAddress)
func (_Polygonzkevmdeployer *PolygonzkevmdeployerFilterer) WatchNewDeterministicDeployment(opts *bind.WatchOpts, sink chan<- *PolygonzkevmdeployerNewDeterministicDeployment) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmdeployer.contract.WatchLogs(opts, "NewDeterministicDeployment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmdeployerNewDeterministicDeployment)
				if err := _Polygonzkevmdeployer.contract.UnpackLog(event, "NewDeterministicDeployment", log); err != nil {
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

// ParseNewDeterministicDeployment is a log parse operation binding the contract event 0xba82f25fed02cd2a23d9f5d11c2ef588d22af5437cbf23bfe61d87257c480e4c.
//
// Solidity: event NewDeterministicDeployment(address newContractAddress)
func (_Polygonzkevmdeployer *PolygonzkevmdeployerFilterer) ParseNewDeterministicDeployment(log types.Log) (*PolygonzkevmdeployerNewDeterministicDeployment, error) {
	event := new(PolygonzkevmdeployerNewDeterministicDeployment)
	if err := _Polygonzkevmdeployer.contract.UnpackLog(event, "NewDeterministicDeployment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmdeployerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Polygonzkevmdeployer contract.
type PolygonzkevmdeployerOwnershipTransferredIterator struct {
	Event *PolygonzkevmdeployerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmdeployerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmdeployerOwnershipTransferred)
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
		it.Event = new(PolygonzkevmdeployerOwnershipTransferred)
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
func (it *PolygonzkevmdeployerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmdeployerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmdeployerOwnershipTransferred represents a OwnershipTransferred event raised by the Polygonzkevmdeployer contract.
type PolygonzkevmdeployerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Polygonzkevmdeployer *PolygonzkevmdeployerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PolygonzkevmdeployerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Polygonzkevmdeployer.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmdeployerOwnershipTransferredIterator{contract: _Polygonzkevmdeployer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Polygonzkevmdeployer *PolygonzkevmdeployerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PolygonzkevmdeployerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Polygonzkevmdeployer.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmdeployerOwnershipTransferred)
				if err := _Polygonzkevmdeployer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Polygonzkevmdeployer *PolygonzkevmdeployerFilterer) ParseOwnershipTransferred(log types.Log) (*PolygonzkevmdeployerOwnershipTransferred, error) {
	event := new(PolygonzkevmdeployerOwnershipTransferred)
	if err := _Polygonzkevmdeployer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

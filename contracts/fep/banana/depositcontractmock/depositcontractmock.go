// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package depositcontractmock

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

// DepositcontractmockMetaData contains all meta data concerning the Depositcontractmock contract.
var DepositcontractmockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDepositRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"name\":\"getLeafValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"verifyMerkleProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061001861001d565b610124565b5f54610100900460ff161580801561003b57505f54600160ff909116105b806100545750303b15801561005457505f5460ff166001145b6100bb5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b5f805460ff1916600117905580156100dc575f805461ff0019166101001790555b8015610121575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b610a08806101315f395ff3fe608060405234801561000f575f80fd5b506004361061006f575f3560e01c80638129fc1c1161004d5780638129fc1c1461017957806382c0190214610183578063fb57083414610196575f80fd5b80632dfdf0b5146100735780633ae050471461008f5780633e19704314610097575b5f80fd5b61007c60535481565b6040519081526020015b60405180910390f35b61007c6101b9565b61007c6100a536600461070b565b6040517fff0000000000000000000000000000000000000000000000000000000000000060f889901b1660208201527fffffffff0000000000000000000000000000000000000000000000000000000060e088811b821660218401527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606089811b821660258601529188901b909216603984015285901b16603d82015260518101839052607181018290525f90609101604051602081830303815290604052805190602001209050979650505050505050565b610181610295565b005b61018161019136600461070b565b61041d565b6101a96101a4366004610785565b6104f8565b6040519015158152602001610086565b6053545f90819081805b602081101561028c578083901c60011660010361022057603381602081106101ed576101ed6107cf565b0154604080516020810192909252810185905260600160405160208183030381529060405280519060200120935061024d565b60408051602081018690529081018390526060016040516020818303038152906040528051906020012093505b6040805160208101849052908101839052606001604051602081830303815290604052805190602001209150808061028490610829565b9150506101c3565b50919392505050565b5f54610100900460ff16158080156102b357505f54600160ff909116105b806102cc5750303b1580156102cc57505f5460ff166001145b61035c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840160405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156103b8575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b801561041a575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b604080517fff0000000000000000000000000000000000000000000000000000000000000060f88a901b166020808301919091527fffffffff0000000000000000000000000000000000000000000000000000000060e08a811b821660218501527fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060608b811b82166025870152918a901b909216603985015287901b16603d83015260518201859052607180830185905283518084039091018152609190920190925280519101206104ef906105d0565b50505050505050565b5f84815b60208110156105c457600163ffffffff8616821c811690036105675785816020811061052a5761052a6107cf565b60200201358260405160200161054a929190918252602082015260400190565b6040516020818303038152906040528051906020012091506105b2565b8186826020811061057a5761057a6107cf565b6020020135604051602001610599929190918252602082015260400190565b6040516020818303038152906040528051906020012091505b806105bc81610829565b9150506104fc565b50909114949350505050565b8060016105df60206002610980565b6105e99190610992565b60535410610623576040517fef5ccf6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60535f815461063290610829565b918290555090505f5b60208110156106c2578082901c60011660010361066e578260338260208110610666576106666107cf565b015550505050565b60338160208110610681576106816107cf565b0154604080516020810192909252810184905260600160405160208183030381529060405280519060200120925080806106ba90610829565b91505061063b565b506106cb6109a5565b505050565b803563ffffffff811681146106e3575f80fd5b919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146106e3575f80fd5b5f805f805f805f60e0888a031215610721575f80fd5b873560ff81168114610731575f80fd5b965061073f602089016106d0565b955061074d604089016106e8565b945061075b606089016106d0565b9350610769608089016106e8565b925060a0880135915060c0880135905092959891949750929550565b5f805f806104608587031215610799575f80fd5b843593506104208501868111156107ae575f80fd5b6020860193506107bd816106d0565b94979396509394610440013593505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610859576108596107fc565b5060010190565b600181815b808511156108b957817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0482111561089f5761089f6107fc565b808516156108ac57918102915b93841c9390800290610865565b509250929050565b5f826108cf5750600161097a565b816108db57505f61097a565b81600181146108f157600281146108fb57610917565b600191505061097a565b60ff84111561090c5761090c6107fc565b50506001821b61097a565b5060208310610133831016604e8410600b841016171561093a575081810a61097a565b6109448383610860565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610976576109766107fc565b0290505b92915050565b5f61098b83836108c1565b9392505050565b8181038181111561097a5761097a6107fc565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52600160045260245ffdfea2646970667358221220dfa64240ca3a53d05f0a998c9969b898e97f500d3bcd8d4a14e25b4b6bdbd62064736f6c63430008140033",
}

// DepositcontractmockABI is the input ABI used to generate the binding from.
// Deprecated: Use DepositcontractmockMetaData.ABI instead.
var DepositcontractmockABI = DepositcontractmockMetaData.ABI

// DepositcontractmockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DepositcontractmockMetaData.Bin instead.
var DepositcontractmockBin = DepositcontractmockMetaData.Bin

// DeployDepositcontractmock deploys a new Ethereum contract, binding an instance of Depositcontractmock to it.
func DeployDepositcontractmock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Depositcontractmock, error) {
	parsed, err := DepositcontractmockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DepositcontractmockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Depositcontractmock{DepositcontractmockCaller: DepositcontractmockCaller{contract: contract}, DepositcontractmockTransactor: DepositcontractmockTransactor{contract: contract}, DepositcontractmockFilterer: DepositcontractmockFilterer{contract: contract}}, nil
}

// Depositcontractmock is an auto generated Go binding around an Ethereum contract.
type Depositcontractmock struct {
	DepositcontractmockCaller     // Read-only binding to the contract
	DepositcontractmockTransactor // Write-only binding to the contract
	DepositcontractmockFilterer   // Log filterer for contract events
}

// DepositcontractmockCaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositcontractmockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositcontractmockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositcontractmockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositcontractmockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositcontractmockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositcontractmockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositcontractmockSession struct {
	Contract     *Depositcontractmock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DepositcontractmockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositcontractmockCallerSession struct {
	Contract *DepositcontractmockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// DepositcontractmockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositcontractmockTransactorSession struct {
	Contract     *DepositcontractmockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// DepositcontractmockRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositcontractmockRaw struct {
	Contract *Depositcontractmock // Generic contract binding to access the raw methods on
}

// DepositcontractmockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositcontractmockCallerRaw struct {
	Contract *DepositcontractmockCaller // Generic read-only contract binding to access the raw methods on
}

// DepositcontractmockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositcontractmockTransactorRaw struct {
	Contract *DepositcontractmockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDepositcontractmock creates a new instance of Depositcontractmock, bound to a specific deployed contract.
func NewDepositcontractmock(address common.Address, backend bind.ContractBackend) (*Depositcontractmock, error) {
	contract, err := bindDepositcontractmock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Depositcontractmock{DepositcontractmockCaller: DepositcontractmockCaller{contract: contract}, DepositcontractmockTransactor: DepositcontractmockTransactor{contract: contract}, DepositcontractmockFilterer: DepositcontractmockFilterer{contract: contract}}, nil
}

// NewDepositcontractmockCaller creates a new read-only instance of Depositcontractmock, bound to a specific deployed contract.
func NewDepositcontractmockCaller(address common.Address, caller bind.ContractCaller) (*DepositcontractmockCaller, error) {
	contract, err := bindDepositcontractmock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositcontractmockCaller{contract: contract}, nil
}

// NewDepositcontractmockTransactor creates a new write-only instance of Depositcontractmock, bound to a specific deployed contract.
func NewDepositcontractmockTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositcontractmockTransactor, error) {
	contract, err := bindDepositcontractmock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositcontractmockTransactor{contract: contract}, nil
}

// NewDepositcontractmockFilterer creates a new log filterer instance of Depositcontractmock, bound to a specific deployed contract.
func NewDepositcontractmockFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositcontractmockFilterer, error) {
	contract, err := bindDepositcontractmock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositcontractmockFilterer{contract: contract}, nil
}

// bindDepositcontractmock binds a generic wrapper to an already deployed contract.
func bindDepositcontractmock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DepositcontractmockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Depositcontractmock *DepositcontractmockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Depositcontractmock.Contract.DepositcontractmockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Depositcontractmock *DepositcontractmockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositcontractmock.Contract.DepositcontractmockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Depositcontractmock *DepositcontractmockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Depositcontractmock.Contract.DepositcontractmockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Depositcontractmock *DepositcontractmockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Depositcontractmock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Depositcontractmock *DepositcontractmockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositcontractmock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Depositcontractmock *DepositcontractmockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Depositcontractmock.Contract.contract.Transact(opts, method, params...)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Depositcontractmock *DepositcontractmockCaller) DepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Depositcontractmock.contract.Call(opts, &out, "depositCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Depositcontractmock *DepositcontractmockSession) DepositCount() (*big.Int, error) {
	return _Depositcontractmock.Contract.DepositCount(&_Depositcontractmock.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Depositcontractmock *DepositcontractmockCallerSession) DepositCount() (*big.Int, error) {
	return _Depositcontractmock.Contract.DepositCount(&_Depositcontractmock.CallOpts)
}

// GetDepositRoot is a free data retrieval call binding the contract method 0x3ae05047.
//
// Solidity: function getDepositRoot() view returns(bytes32)
func (_Depositcontractmock *DepositcontractmockCaller) GetDepositRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Depositcontractmock.contract.Call(opts, &out, "getDepositRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDepositRoot is a free data retrieval call binding the contract method 0x3ae05047.
//
// Solidity: function getDepositRoot() view returns(bytes32)
func (_Depositcontractmock *DepositcontractmockSession) GetDepositRoot() ([32]byte, error) {
	return _Depositcontractmock.Contract.GetDepositRoot(&_Depositcontractmock.CallOpts)
}

// GetDepositRoot is a free data retrieval call binding the contract method 0x3ae05047.
//
// Solidity: function getDepositRoot() view returns(bytes32)
func (_Depositcontractmock *DepositcontractmockCallerSession) GetDepositRoot() ([32]byte, error) {
	return _Depositcontractmock.Contract.GetDepositRoot(&_Depositcontractmock.CallOpts)
}

// GetLeafValue is a free data retrieval call binding the contract method 0x3e197043.
//
// Solidity: function getLeafValue(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes32 metadataHash) pure returns(bytes32)
func (_Depositcontractmock *DepositcontractmockCaller) GetLeafValue(opts *bind.CallOpts, leafType uint8, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadataHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Depositcontractmock.contract.Call(opts, &out, "getLeafValue", leafType, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadataHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLeafValue is a free data retrieval call binding the contract method 0x3e197043.
//
// Solidity: function getLeafValue(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes32 metadataHash) pure returns(bytes32)
func (_Depositcontractmock *DepositcontractmockSession) GetLeafValue(leafType uint8, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadataHash [32]byte) ([32]byte, error) {
	return _Depositcontractmock.Contract.GetLeafValue(&_Depositcontractmock.CallOpts, leafType, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadataHash)
}

// GetLeafValue is a free data retrieval call binding the contract method 0x3e197043.
//
// Solidity: function getLeafValue(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes32 metadataHash) pure returns(bytes32)
func (_Depositcontractmock *DepositcontractmockCallerSession) GetLeafValue(leafType uint8, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadataHash [32]byte) ([32]byte, error) {
	return _Depositcontractmock.Contract.GetLeafValue(&_Depositcontractmock.CallOpts, leafType, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadataHash)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Depositcontractmock *DepositcontractmockCaller) VerifyMerkleProof(opts *bind.CallOpts, leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	var out []interface{}
	err := _Depositcontractmock.contract.Call(opts, &out, "verifyMerkleProof", leafHash, smtProof, index, root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Depositcontractmock *DepositcontractmockSession) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	return _Depositcontractmock.Contract.VerifyMerkleProof(&_Depositcontractmock.CallOpts, leafHash, smtProof, index, root)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Depositcontractmock *DepositcontractmockCallerSession) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	return _Depositcontractmock.Contract.VerifyMerkleProof(&_Depositcontractmock.CallOpts, leafHash, smtProof, index, root)
}

// Deposit is a paid mutator transaction binding the contract method 0x82c01902.
//
// Solidity: function deposit(uint8 leafType, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes32 metadataHash) returns()
func (_Depositcontractmock *DepositcontractmockTransactor) Deposit(opts *bind.TransactOpts, leafType uint8, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadataHash [32]byte) (*types.Transaction, error) {
	return _Depositcontractmock.contract.Transact(opts, "deposit", leafType, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadataHash)
}

// Deposit is a paid mutator transaction binding the contract method 0x82c01902.
//
// Solidity: function deposit(uint8 leafType, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes32 metadataHash) returns()
func (_Depositcontractmock *DepositcontractmockSession) Deposit(leafType uint8, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadataHash [32]byte) (*types.Transaction, error) {
	return _Depositcontractmock.Contract.Deposit(&_Depositcontractmock.TransactOpts, leafType, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadataHash)
}

// Deposit is a paid mutator transaction binding the contract method 0x82c01902.
//
// Solidity: function deposit(uint8 leafType, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes32 metadataHash) returns()
func (_Depositcontractmock *DepositcontractmockTransactorSession) Deposit(leafType uint8, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadataHash [32]byte) (*types.Transaction, error) {
	return _Depositcontractmock.Contract.Deposit(&_Depositcontractmock.TransactOpts, leafType, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadataHash)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Depositcontractmock *DepositcontractmockTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositcontractmock.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Depositcontractmock *DepositcontractmockSession) Initialize() (*types.Transaction, error) {
	return _Depositcontractmock.Contract.Initialize(&_Depositcontractmock.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Depositcontractmock *DepositcontractmockTransactorSession) Initialize() (*types.Transaction, error) {
	return _Depositcontractmock.Contract.Initialize(&_Depositcontractmock.TransactOpts)
}

// DepositcontractmockInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Depositcontractmock contract.
type DepositcontractmockInitializedIterator struct {
	Event *DepositcontractmockInitialized // Event containing the contract specifics and raw log

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
func (it *DepositcontractmockInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositcontractmockInitialized)
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
		it.Event = new(DepositcontractmockInitialized)
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
func (it *DepositcontractmockInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositcontractmockInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositcontractmockInitialized represents a Initialized event raised by the Depositcontractmock contract.
type DepositcontractmockInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Depositcontractmock *DepositcontractmockFilterer) FilterInitialized(opts *bind.FilterOpts) (*DepositcontractmockInitializedIterator, error) {

	logs, sub, err := _Depositcontractmock.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DepositcontractmockInitializedIterator{contract: _Depositcontractmock.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Depositcontractmock *DepositcontractmockFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DepositcontractmockInitialized) (event.Subscription, error) {

	logs, sub, err := _Depositcontractmock.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositcontractmockInitialized)
				if err := _Depositcontractmock.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Depositcontractmock *DepositcontractmockFilterer) ParseInitialized(log types.Log) (*DepositcontractmockInitialized, error) {
	event := new(DepositcontractmockInitialized)
	if err := _Depositcontractmock.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package depositcontractbase

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

// DepositcontractbaseMetaData contains all meta data concerning the Depositcontractbase contract.
var DepositcontractbaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"calculateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"verifyMerkleProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506103de8061001d5f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c80632dfdf0b51461004e5780635ca1e1651461006a57806383f2440314610072578063fb57083414610085575b5f80fd5b61005760205481565b6040519081526020015b60405180910390f35b6100576100a8565b61005761008036600461029e565b610183565b6100986100933660046102da565b610258565b6040519015158152602001610061565b6020545f90819081805b602081101561017a578083901c60011660010361010e575f81602081106100db576100db61031f565b0154604080516020810192909252810185905260600160405160208183030381529060405280519060200120935061013b565b60408051602081018690529081018390526060016040516020818303038152906040528051906020012093505b604080516020810184905290810183905260600160405160208183030381529060405280519060200120915080806101729061034c565b9150506100b2565b50919392505050565b5f83815b602081101561024f57600163ffffffff8516821c811690036101f2578481602081106101b5576101b561031f565b6020020135826040516020016101d5929190918252602082015260400190565b60405160208183030381529060405280519060200120915061023d565b818582602081106102055761020561031f565b6020020135604051602001610224929190918252602082015260400190565b6040516020818303038152906040528051906020012091505b806102478161034c565b915050610187565b50949350505050565b5f81610265868686610183565b1495945050505050565b806104008101831015610280575f80fd5b92915050565b803563ffffffff81168114610299575f80fd5b919050565b5f805f61044084860312156102b1575f80fd5b833592506102c2856020860161026f565b91506102d16104208501610286565b90509250925092565b5f805f8061046085870312156102ee575f80fd5b843593506102ff866020870161026f565b925061030e6104208601610286565b939692955092936104400135925050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036103a1577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b506001019056fea264697066735822122011a6793f8bf3ba032bf8b7b939110b18c43994d8b34ea3e08724cf956e1aebbe64736f6c63430008140033",
}

// DepositcontractbaseABI is the input ABI used to generate the binding from.
// Deprecated: Use DepositcontractbaseMetaData.ABI instead.
var DepositcontractbaseABI = DepositcontractbaseMetaData.ABI

// DepositcontractbaseBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DepositcontractbaseMetaData.Bin instead.
var DepositcontractbaseBin = DepositcontractbaseMetaData.Bin

// DeployDepositcontractbase deploys a new Ethereum contract, binding an instance of Depositcontractbase to it.
func DeployDepositcontractbase(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Depositcontractbase, error) {
	parsed, err := DepositcontractbaseMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DepositcontractbaseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Depositcontractbase{DepositcontractbaseCaller: DepositcontractbaseCaller{contract: contract}, DepositcontractbaseTransactor: DepositcontractbaseTransactor{contract: contract}, DepositcontractbaseFilterer: DepositcontractbaseFilterer{contract: contract}}, nil
}

// Depositcontractbase is an auto generated Go binding around an Ethereum contract.
type Depositcontractbase struct {
	DepositcontractbaseCaller     // Read-only binding to the contract
	DepositcontractbaseTransactor // Write-only binding to the contract
	DepositcontractbaseFilterer   // Log filterer for contract events
}

// DepositcontractbaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositcontractbaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositcontractbaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositcontractbaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositcontractbaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositcontractbaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositcontractbaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositcontractbaseSession struct {
	Contract     *Depositcontractbase // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DepositcontractbaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositcontractbaseCallerSession struct {
	Contract *DepositcontractbaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// DepositcontractbaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositcontractbaseTransactorSession struct {
	Contract     *DepositcontractbaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// DepositcontractbaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositcontractbaseRaw struct {
	Contract *Depositcontractbase // Generic contract binding to access the raw methods on
}

// DepositcontractbaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositcontractbaseCallerRaw struct {
	Contract *DepositcontractbaseCaller // Generic read-only contract binding to access the raw methods on
}

// DepositcontractbaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositcontractbaseTransactorRaw struct {
	Contract *DepositcontractbaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDepositcontractbase creates a new instance of Depositcontractbase, bound to a specific deployed contract.
func NewDepositcontractbase(address common.Address, backend bind.ContractBackend) (*Depositcontractbase, error) {
	contract, err := bindDepositcontractbase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Depositcontractbase{DepositcontractbaseCaller: DepositcontractbaseCaller{contract: contract}, DepositcontractbaseTransactor: DepositcontractbaseTransactor{contract: contract}, DepositcontractbaseFilterer: DepositcontractbaseFilterer{contract: contract}}, nil
}

// NewDepositcontractbaseCaller creates a new read-only instance of Depositcontractbase, bound to a specific deployed contract.
func NewDepositcontractbaseCaller(address common.Address, caller bind.ContractCaller) (*DepositcontractbaseCaller, error) {
	contract, err := bindDepositcontractbase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositcontractbaseCaller{contract: contract}, nil
}

// NewDepositcontractbaseTransactor creates a new write-only instance of Depositcontractbase, bound to a specific deployed contract.
func NewDepositcontractbaseTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositcontractbaseTransactor, error) {
	contract, err := bindDepositcontractbase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositcontractbaseTransactor{contract: contract}, nil
}

// NewDepositcontractbaseFilterer creates a new log filterer instance of Depositcontractbase, bound to a specific deployed contract.
func NewDepositcontractbaseFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositcontractbaseFilterer, error) {
	contract, err := bindDepositcontractbase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositcontractbaseFilterer{contract: contract}, nil
}

// bindDepositcontractbase binds a generic wrapper to an already deployed contract.
func bindDepositcontractbase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DepositcontractbaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Depositcontractbase *DepositcontractbaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Depositcontractbase.Contract.DepositcontractbaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Depositcontractbase *DepositcontractbaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositcontractbase.Contract.DepositcontractbaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Depositcontractbase *DepositcontractbaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Depositcontractbase.Contract.DepositcontractbaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Depositcontractbase *DepositcontractbaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Depositcontractbase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Depositcontractbase *DepositcontractbaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositcontractbase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Depositcontractbase *DepositcontractbaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Depositcontractbase.Contract.contract.Transact(opts, method, params...)
}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_Depositcontractbase *DepositcontractbaseCaller) CalculateRoot(opts *bind.CallOpts, leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	var out []interface{}
	err := _Depositcontractbase.contract.Call(opts, &out, "calculateRoot", leafHash, smtProof, index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_Depositcontractbase *DepositcontractbaseSession) CalculateRoot(leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	return _Depositcontractbase.Contract.CalculateRoot(&_Depositcontractbase.CallOpts, leafHash, smtProof, index)
}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_Depositcontractbase *DepositcontractbaseCallerSession) CalculateRoot(leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	return _Depositcontractbase.Contract.CalculateRoot(&_Depositcontractbase.CallOpts, leafHash, smtProof, index)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Depositcontractbase *DepositcontractbaseCaller) DepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Depositcontractbase.contract.Call(opts, &out, "depositCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Depositcontractbase *DepositcontractbaseSession) DepositCount() (*big.Int, error) {
	return _Depositcontractbase.Contract.DepositCount(&_Depositcontractbase.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Depositcontractbase *DepositcontractbaseCallerSession) DepositCount() (*big.Int, error) {
	return _Depositcontractbase.Contract.DepositCount(&_Depositcontractbase.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Depositcontractbase *DepositcontractbaseCaller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Depositcontractbase.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Depositcontractbase *DepositcontractbaseSession) GetRoot() ([32]byte, error) {
	return _Depositcontractbase.Contract.GetRoot(&_Depositcontractbase.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Depositcontractbase *DepositcontractbaseCallerSession) GetRoot() ([32]byte, error) {
	return _Depositcontractbase.Contract.GetRoot(&_Depositcontractbase.CallOpts)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Depositcontractbase *DepositcontractbaseCaller) VerifyMerkleProof(opts *bind.CallOpts, leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	var out []interface{}
	err := _Depositcontractbase.contract.Call(opts, &out, "verifyMerkleProof", leafHash, smtProof, index, root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Depositcontractbase *DepositcontractbaseSession) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	return _Depositcontractbase.Contract.VerifyMerkleProof(&_Depositcontractbase.CallOpts, leafHash, smtProof, index, root)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Depositcontractbase *DepositcontractbaseCallerSession) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	return _Depositcontractbase.Contract.VerifyMerkleProof(&_Depositcontractbase.CallOpts, leafHash, smtProof, index, root)
}

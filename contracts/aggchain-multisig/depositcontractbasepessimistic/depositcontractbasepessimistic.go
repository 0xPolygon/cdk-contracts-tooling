// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package depositcontractbasepessimistic

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

// DepositcontractbasepessimisticMetaData contains all meta data concerning the Depositcontractbasepessimistic contract.
var DepositcontractbasepessimisticMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"calculateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"verifyMerkleProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061038d8061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c80632dfdf0b51461004e5780635ca1e1651461006a57806383f2440314610072578063fb57083414610085575b5f5ffd5b61005760205481565b6040519081526020015b60405180910390f35b6100576100a8565b6100576100803660046102a9565b610198565b6100986100933660046102e5565b610263565b6040519015158152602001610061565b6020545f90819081805b602081101561018f578083901c60011660010361010e575f81602081106100db576100db61032a565b0154604080516020810192909252810185905260600160405160208183030381529060405280519060200120935061013b565b60408051602081018690529081018390526060016040516020818303038152906040528051906020012093505b6040805160208101849052908101839052606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012091506001016100b2565b50919392505050565b5f83815b602081101561025a57600163ffffffff8516821c81169003610207578481602081106101ca576101ca61032a565b6020020135826040516020016101ea929190918252602082015260400190565b604051602081830303815290604052805190602001209150610252565b8185826020811061021a5761021a61032a565b6020020135604051602001610239929190918252602082015260400190565b6040516020818303038152906040528051906020012091505b60010161019c565b50949350505050565b5f81610270868686610198565b1495945050505050565b80610400810183101561028b575f5ffd5b92915050565b803563ffffffff811681146102a4575f5ffd5b919050565b5f5f5f61044084860312156102bc575f5ffd5b833592506102cd856020860161027a565b91506102dc6104208501610291565b90509250925092565b5f5f5f5f61046085870312156102f9575f5ffd5b8435935061030a866020870161027a565b92506103196104208601610291565b939692955092936104400135925050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea26469706673582212208b612c397b43dc942e46ea157b2716a427350e7b1ddbd6606b3dee29d2bf6e7864736f6c634300081c0033",
}

// DepositcontractbasepessimisticABI is the input ABI used to generate the binding from.
// Deprecated: Use DepositcontractbasepessimisticMetaData.ABI instead.
var DepositcontractbasepessimisticABI = DepositcontractbasepessimisticMetaData.ABI

// DepositcontractbasepessimisticBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DepositcontractbasepessimisticMetaData.Bin instead.
var DepositcontractbasepessimisticBin = DepositcontractbasepessimisticMetaData.Bin

// DeployDepositcontractbasepessimistic deploys a new Ethereum contract, binding an instance of Depositcontractbasepessimistic to it.
func DeployDepositcontractbasepessimistic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Depositcontractbasepessimistic, error) {
	parsed, err := DepositcontractbasepessimisticMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DepositcontractbasepessimisticBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Depositcontractbasepessimistic{DepositcontractbasepessimisticCaller: DepositcontractbasepessimisticCaller{contract: contract}, DepositcontractbasepessimisticTransactor: DepositcontractbasepessimisticTransactor{contract: contract}, DepositcontractbasepessimisticFilterer: DepositcontractbasepessimisticFilterer{contract: contract}}, nil
}

// Depositcontractbasepessimistic is an auto generated Go binding around an Ethereum contract.
type Depositcontractbasepessimistic struct {
	DepositcontractbasepessimisticCaller     // Read-only binding to the contract
	DepositcontractbasepessimisticTransactor // Write-only binding to the contract
	DepositcontractbasepessimisticFilterer   // Log filterer for contract events
}

// DepositcontractbasepessimisticCaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositcontractbasepessimisticCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositcontractbasepessimisticTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositcontractbasepessimisticTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositcontractbasepessimisticFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositcontractbasepessimisticFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositcontractbasepessimisticSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositcontractbasepessimisticSession struct {
	Contract     *Depositcontractbasepessimistic // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// DepositcontractbasepessimisticCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositcontractbasepessimisticCallerSession struct {
	Contract *DepositcontractbasepessimisticCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// DepositcontractbasepessimisticTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositcontractbasepessimisticTransactorSession struct {
	Contract     *DepositcontractbasepessimisticTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// DepositcontractbasepessimisticRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositcontractbasepessimisticRaw struct {
	Contract *Depositcontractbasepessimistic // Generic contract binding to access the raw methods on
}

// DepositcontractbasepessimisticCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositcontractbasepessimisticCallerRaw struct {
	Contract *DepositcontractbasepessimisticCaller // Generic read-only contract binding to access the raw methods on
}

// DepositcontractbasepessimisticTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositcontractbasepessimisticTransactorRaw struct {
	Contract *DepositcontractbasepessimisticTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDepositcontractbasepessimistic creates a new instance of Depositcontractbasepessimistic, bound to a specific deployed contract.
func NewDepositcontractbasepessimistic(address common.Address, backend bind.ContractBackend) (*Depositcontractbasepessimistic, error) {
	contract, err := bindDepositcontractbasepessimistic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Depositcontractbasepessimistic{DepositcontractbasepessimisticCaller: DepositcontractbasepessimisticCaller{contract: contract}, DepositcontractbasepessimisticTransactor: DepositcontractbasepessimisticTransactor{contract: contract}, DepositcontractbasepessimisticFilterer: DepositcontractbasepessimisticFilterer{contract: contract}}, nil
}

// NewDepositcontractbasepessimisticCaller creates a new read-only instance of Depositcontractbasepessimistic, bound to a specific deployed contract.
func NewDepositcontractbasepessimisticCaller(address common.Address, caller bind.ContractCaller) (*DepositcontractbasepessimisticCaller, error) {
	contract, err := bindDepositcontractbasepessimistic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositcontractbasepessimisticCaller{contract: contract}, nil
}

// NewDepositcontractbasepessimisticTransactor creates a new write-only instance of Depositcontractbasepessimistic, bound to a specific deployed contract.
func NewDepositcontractbasepessimisticTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositcontractbasepessimisticTransactor, error) {
	contract, err := bindDepositcontractbasepessimistic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositcontractbasepessimisticTransactor{contract: contract}, nil
}

// NewDepositcontractbasepessimisticFilterer creates a new log filterer instance of Depositcontractbasepessimistic, bound to a specific deployed contract.
func NewDepositcontractbasepessimisticFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositcontractbasepessimisticFilterer, error) {
	contract, err := bindDepositcontractbasepessimistic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositcontractbasepessimisticFilterer{contract: contract}, nil
}

// bindDepositcontractbasepessimistic binds a generic wrapper to an already deployed contract.
func bindDepositcontractbasepessimistic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DepositcontractbasepessimisticMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Depositcontractbasepessimistic *DepositcontractbasepessimisticRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Depositcontractbasepessimistic.Contract.DepositcontractbasepessimisticCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Depositcontractbasepessimistic *DepositcontractbasepessimisticRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositcontractbasepessimistic.Contract.DepositcontractbasepessimisticTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Depositcontractbasepessimistic *DepositcontractbasepessimisticRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Depositcontractbasepessimistic.Contract.DepositcontractbasepessimisticTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Depositcontractbasepessimistic *DepositcontractbasepessimisticCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Depositcontractbasepessimistic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Depositcontractbasepessimistic *DepositcontractbasepessimisticTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositcontractbasepessimistic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Depositcontractbasepessimistic *DepositcontractbasepessimisticTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Depositcontractbasepessimistic.Contract.contract.Transact(opts, method, params...)
}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_Depositcontractbasepessimistic *DepositcontractbasepessimisticCaller) CalculateRoot(opts *bind.CallOpts, leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	var out []interface{}
	err := _Depositcontractbasepessimistic.contract.Call(opts, &out, "calculateRoot", leafHash, smtProof, index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_Depositcontractbasepessimistic *DepositcontractbasepessimisticSession) CalculateRoot(leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	return _Depositcontractbasepessimistic.Contract.CalculateRoot(&_Depositcontractbasepessimistic.CallOpts, leafHash, smtProof, index)
}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_Depositcontractbasepessimistic *DepositcontractbasepessimisticCallerSession) CalculateRoot(leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	return _Depositcontractbasepessimistic.Contract.CalculateRoot(&_Depositcontractbasepessimistic.CallOpts, leafHash, smtProof, index)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Depositcontractbasepessimistic *DepositcontractbasepessimisticCaller) DepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Depositcontractbasepessimistic.contract.Call(opts, &out, "depositCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Depositcontractbasepessimistic *DepositcontractbasepessimisticSession) DepositCount() (*big.Int, error) {
	return _Depositcontractbasepessimistic.Contract.DepositCount(&_Depositcontractbasepessimistic.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Depositcontractbasepessimistic *DepositcontractbasepessimisticCallerSession) DepositCount() (*big.Int, error) {
	return _Depositcontractbasepessimistic.Contract.DepositCount(&_Depositcontractbasepessimistic.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Depositcontractbasepessimistic *DepositcontractbasepessimisticCaller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Depositcontractbasepessimistic.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Depositcontractbasepessimistic *DepositcontractbasepessimisticSession) GetRoot() ([32]byte, error) {
	return _Depositcontractbasepessimistic.Contract.GetRoot(&_Depositcontractbasepessimistic.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Depositcontractbasepessimistic *DepositcontractbasepessimisticCallerSession) GetRoot() ([32]byte, error) {
	return _Depositcontractbasepessimistic.Contract.GetRoot(&_Depositcontractbasepessimistic.CallOpts)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Depositcontractbasepessimistic *DepositcontractbasepessimisticCaller) VerifyMerkleProof(opts *bind.CallOpts, leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	var out []interface{}
	err := _Depositcontractbasepessimistic.contract.Call(opts, &out, "verifyMerkleProof", leafHash, smtProof, index, root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Depositcontractbasepessimistic *DepositcontractbasepessimisticSession) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	return _Depositcontractbasepessimistic.Contract.VerifyMerkleProof(&_Depositcontractbasepessimistic.CallOpts, leafHash, smtProof, index, root)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Depositcontractbasepessimistic *DepositcontractbasepessimisticCallerSession) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	return _Depositcontractbasepessimistic.Contract.VerifyMerkleProof(&_Depositcontractbasepessimistic.CallOpts, leafHash, smtProof, index, root)
}

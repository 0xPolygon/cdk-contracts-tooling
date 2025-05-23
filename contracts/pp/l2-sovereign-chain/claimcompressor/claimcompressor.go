// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package claimcompressor

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

// ClaimCompressorCompressClaimCallData is an auto generated low-level Go binding around an user-defined struct.
type ClaimCompressorCompressClaimCallData struct {
	SmtProofLocalExitRoot  [32][32]byte
	SmtProofRollupExitRoot [32][32]byte
	GlobalIndex            *big.Int
	OriginNetwork          uint32
	OriginAddress          common.Address
	DestinationAddress     common.Address
	Amount                 *big.Int
	Metadata               []byte
	IsMessage              bool
}

// ClaimcompressorMetaData contains all meta data concerning the Claimcompressor contract.
var ClaimcompressorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"__networkID\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isMessage\",\"type\":\"bool\"}],\"internalType\":\"structClaimCompressor.CompressClaimCallData[]\",\"name\":\"compressClaimCalldata\",\"type\":\"tuple[]\"}],\"name\":\"compressClaimCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"compressedClaimCalls\",\"type\":\"bytes\"}],\"name\":\"sendCompressedClaims\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c0604052348015600e575f5ffd5b50604051610d83380380610d83833981016040819052602b916046565b6001600160a01b0390911660805263ffffffff1660a052608e565b5f5f604083850312156056575f5ffd5b82516001600160a01b0381168114606b575f5ffd5b602084015190925063ffffffff811681146083575f5ffd5b809150509250929050565b60805160a051610cd46100af5f395f61046301525f6104850152610cd45ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c806304e5557b1461003857806397b1410f14610061575b5f5ffd5b61004b61004636600461062d565b610076565b60405161005891906106ac565b60405180910390f35b61007461006f3660046106ff565b61045c565b005b60605f8585604051602001610095929190918252602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905290505f805b84811015610450575f8686838181106100e3576100e361076d565b90506020028101906100f5919061079a565b6100fe9061098d565b90505f82810361011057506020610196565b5f5b6020811015610194578888610128600187610a83565b8181106101375761013761076d565b9050602002810190610149919061079a565b816020811061015a5761015a61076d565b6020020135835f015182602081106101745761017461076d565b60200201511461018c57610189816001610a9c565b91505b600101610112565b505b6040805160f883901b7fff000000000000000000000000000000000000000000000000000000000000001660208201528151600181830301815260219091019091525f5b82811015610245578351829082602081106101f7576101f761076d565b602002015160405160200161020d929190610ac6565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905291506001016101da565b5060208301516103e001515f92501561030157831580610263575084155b1561027157602091506102fc565b5f5b60208110156102fa578989610289600188610a83565b8181106102985761029861076d565b90506020028101906102aa919061079a565b6104000181602081106102bf576102bf61076d565b6020020135846020015182602081106102da576102da61076d565b6020020151146102f2576102ef816001610a9c565b92505b600101610273565b505b600194505b8082604051602001610314929190610ade565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905290505f5b828110156103b35781846020015182602081106103655761036561076d565b602002015160405160200161037b929190610ac6565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529150600101610346565b505f8361010001516103c5575f6103c8565b60015b6040858101516060870151608088015160a089015160c08a015160e08b01518051875161040a99988c989081901c979096909590949093909291602001610b1d565b6040516020818303038152906040529050868160405160200161042e929190610c82565b60405160208183030381529060405296505050505080806001019150506100c8565b50909695505050505050565b63ffffffff7f0000000000000000000000000000000000000000000000000000000000000000167f000000000000000000000000000000000000000000000000000000000000000063ccaa2d1163f5efcd798585602082610824376020828101610844376108a486905261092061090452604082015b81830181101561062257803560f81c80156104f457600181146105115761052a565b856003538560081c6002538560101c6001538560181c5f5361052a565b846003538460081c6002538460101c6001538460181c5f535b5060028101906004906001013560f81c60200280838337909101600181019161040001906020903560f81c02808383379190910190610400810190823560f81c9061041701536001820191506008826018830137600882019150606081019050600482601c830137600482019150602081019050601482600c83013760149182019160408201918390604c01376014820191506020810190506020828237602082013560e01c60408201819052602490920191606090910190808383375f9181019190915290810190601f808216602003160161094401622625a05a1015610610575f5ffd5b5f5f825f5f8b621e8480f150506104d2565b505050505050505050565b5f5f5f5f60608587031215610640575f5ffd5b8435935060208501359250604085013567ffffffffffffffff811115610664575f5ffd5b8501601f81018713610674575f5ffd5b803567ffffffffffffffff81111561068a575f5ffd5b8760208260051b840101111561069e575f5ffd5b949793965060200194505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b5f5f60208385031215610710575f5ffd5b823567ffffffffffffffff811115610726575f5ffd5b8301601f81018513610736575f5ffd5b803567ffffffffffffffff81111561074c575f5ffd5b85602082840101111561075d575f5ffd5b6020919091019590945092505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f82357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7218336030181126107cc575f5ffd5b9190910192915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051610120810167ffffffffffffffff81118282101715610827576108276107d6565b60405290565b5f82601f83011261083c575f5ffd5b604051610400810167ffffffffffffffff81118282101715610860576108606107d6565b60405280610400840185811115610875575f5ffd5b845b8181101561088f578035835260209283019201610877565b509195945050505050565b803563ffffffff811681146108ad575f5ffd5b919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146108ad575f5ffd5b5f82601f8301126108e4575f5ffd5b813567ffffffffffffffff8111156108fe576108fe6107d6565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810167ffffffffffffffff8111828210171561094b5761094b6107d6565b604052818152838201602001851015610962575f5ffd5b816020850160208301375f918101602001919091529392505050565b803580151581146108ad575f5ffd5b5f6108e0823603121561099e575f5ffd5b6109a6610803565b6109b0368461082d565b81526109c036610400850161082d565b602082015261080083013560408201526109dd610820840161089a565b60608201526109ef61084084016108b2565b6080820152610a0161086084016108b2565b60a082015261088083013560c08201526108a083013567ffffffffffffffff811115610a2b575f5ffd5b610a37368286016108d5565b60e083015250610a4a6108c0840161097e565b61010082015292915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b81810381811115610a9657610a96610a56565b92915050565b80820180821115610a9657610a96610a56565b5f81518060208401855e5f93019283525090919050565b5f610ad18285610aaf565b9283525050602001919050565b5f610ae98285610aaf565b60f89390931b7fff000000000000000000000000000000000000000000000000000000000000001683525050600101919050565b7fff000000000000000000000000000000000000000000000000000000000000008b60f81b1681525f610b53600183018c610aaf565b7fff000000000000000000000000000000000000000000000000000000000000008b60f81b1681527fffffffffffffffff0000000000000000000000000000000000000000000000008a60c01b1660018201527fffffffff000000000000000000000000000000000000000000000000000000008960e01b1660098201527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008860601b16600d820152610c2d602182018860601b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000169052565b856035820152610c64605582018660e01b7fffffffff00000000000000000000000000000000000000000000000000000000169052565b610c716059820185610aaf565b9d9c50505050505050505050505050565b5f610c96610c908386610aaf565b84610aaf565b94935050505056fea2646970667358221220ab1ab0f53955b110d4a441de94e11ae5ee8708657fed7a3cc5f83add6a48d7fb64736f6c634300081c0033",
}

// ClaimcompressorABI is the input ABI used to generate the binding from.
// Deprecated: Use ClaimcompressorMetaData.ABI instead.
var ClaimcompressorABI = ClaimcompressorMetaData.ABI

// ClaimcompressorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ClaimcompressorMetaData.Bin instead.
var ClaimcompressorBin = ClaimcompressorMetaData.Bin

// DeployClaimcompressor deploys a new Ethereum contract, binding an instance of Claimcompressor to it.
func DeployClaimcompressor(auth *bind.TransactOpts, backend bind.ContractBackend, __bridgeAddress common.Address, __networkID uint32) (common.Address, *types.Transaction, *Claimcompressor, error) {
	parsed, err := ClaimcompressorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ClaimcompressorBin), backend, __bridgeAddress, __networkID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Claimcompressor{ClaimcompressorCaller: ClaimcompressorCaller{contract: contract}, ClaimcompressorTransactor: ClaimcompressorTransactor{contract: contract}, ClaimcompressorFilterer: ClaimcompressorFilterer{contract: contract}}, nil
}

// Claimcompressor is an auto generated Go binding around an Ethereum contract.
type Claimcompressor struct {
	ClaimcompressorCaller     // Read-only binding to the contract
	ClaimcompressorTransactor // Write-only binding to the contract
	ClaimcompressorFilterer   // Log filterer for contract events
}

// ClaimcompressorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClaimcompressorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimcompressorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClaimcompressorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimcompressorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClaimcompressorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimcompressorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClaimcompressorSession struct {
	Contract     *Claimcompressor  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClaimcompressorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClaimcompressorCallerSession struct {
	Contract *ClaimcompressorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ClaimcompressorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClaimcompressorTransactorSession struct {
	Contract     *ClaimcompressorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ClaimcompressorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClaimcompressorRaw struct {
	Contract *Claimcompressor // Generic contract binding to access the raw methods on
}

// ClaimcompressorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClaimcompressorCallerRaw struct {
	Contract *ClaimcompressorCaller // Generic read-only contract binding to access the raw methods on
}

// ClaimcompressorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClaimcompressorTransactorRaw struct {
	Contract *ClaimcompressorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClaimcompressor creates a new instance of Claimcompressor, bound to a specific deployed contract.
func NewClaimcompressor(address common.Address, backend bind.ContractBackend) (*Claimcompressor, error) {
	contract, err := bindClaimcompressor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Claimcompressor{ClaimcompressorCaller: ClaimcompressorCaller{contract: contract}, ClaimcompressorTransactor: ClaimcompressorTransactor{contract: contract}, ClaimcompressorFilterer: ClaimcompressorFilterer{contract: contract}}, nil
}

// NewClaimcompressorCaller creates a new read-only instance of Claimcompressor, bound to a specific deployed contract.
func NewClaimcompressorCaller(address common.Address, caller bind.ContractCaller) (*ClaimcompressorCaller, error) {
	contract, err := bindClaimcompressor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimcompressorCaller{contract: contract}, nil
}

// NewClaimcompressorTransactor creates a new write-only instance of Claimcompressor, bound to a specific deployed contract.
func NewClaimcompressorTransactor(address common.Address, transactor bind.ContractTransactor) (*ClaimcompressorTransactor, error) {
	contract, err := bindClaimcompressor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimcompressorTransactor{contract: contract}, nil
}

// NewClaimcompressorFilterer creates a new log filterer instance of Claimcompressor, bound to a specific deployed contract.
func NewClaimcompressorFilterer(address common.Address, filterer bind.ContractFilterer) (*ClaimcompressorFilterer, error) {
	contract, err := bindClaimcompressor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClaimcompressorFilterer{contract: contract}, nil
}

// bindClaimcompressor binds a generic wrapper to an already deployed contract.
func bindClaimcompressor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ClaimcompressorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Claimcompressor *ClaimcompressorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Claimcompressor.Contract.ClaimcompressorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Claimcompressor *ClaimcompressorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Claimcompressor.Contract.ClaimcompressorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Claimcompressor *ClaimcompressorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Claimcompressor.Contract.ClaimcompressorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Claimcompressor *ClaimcompressorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Claimcompressor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Claimcompressor *ClaimcompressorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Claimcompressor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Claimcompressor *ClaimcompressorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Claimcompressor.Contract.contract.Transact(opts, method, params...)
}

// CompressClaimCall is a free data retrieval call binding the contract method 0x04e5557b.
//
// Solidity: function compressClaimCall(bytes32 mainnetExitRoot, bytes32 rollupExitRoot, (bytes32[32],bytes32[32],uint256,uint32,address,address,uint256,bytes,bool)[] compressClaimCalldata) pure returns(bytes)
func (_Claimcompressor *ClaimcompressorCaller) CompressClaimCall(opts *bind.CallOpts, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, compressClaimCalldata []ClaimCompressorCompressClaimCallData) ([]byte, error) {
	var out []interface{}
	err := _Claimcompressor.contract.Call(opts, &out, "compressClaimCall", mainnetExitRoot, rollupExitRoot, compressClaimCalldata)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CompressClaimCall is a free data retrieval call binding the contract method 0x04e5557b.
//
// Solidity: function compressClaimCall(bytes32 mainnetExitRoot, bytes32 rollupExitRoot, (bytes32[32],bytes32[32],uint256,uint32,address,address,uint256,bytes,bool)[] compressClaimCalldata) pure returns(bytes)
func (_Claimcompressor *ClaimcompressorSession) CompressClaimCall(mainnetExitRoot [32]byte, rollupExitRoot [32]byte, compressClaimCalldata []ClaimCompressorCompressClaimCallData) ([]byte, error) {
	return _Claimcompressor.Contract.CompressClaimCall(&_Claimcompressor.CallOpts, mainnetExitRoot, rollupExitRoot, compressClaimCalldata)
}

// CompressClaimCall is a free data retrieval call binding the contract method 0x04e5557b.
//
// Solidity: function compressClaimCall(bytes32 mainnetExitRoot, bytes32 rollupExitRoot, (bytes32[32],bytes32[32],uint256,uint32,address,address,uint256,bytes,bool)[] compressClaimCalldata) pure returns(bytes)
func (_Claimcompressor *ClaimcompressorCallerSession) CompressClaimCall(mainnetExitRoot [32]byte, rollupExitRoot [32]byte, compressClaimCalldata []ClaimCompressorCompressClaimCallData) ([]byte, error) {
	return _Claimcompressor.Contract.CompressClaimCall(&_Claimcompressor.CallOpts, mainnetExitRoot, rollupExitRoot, compressClaimCalldata)
}

// SendCompressedClaims is a paid mutator transaction binding the contract method 0x97b1410f.
//
// Solidity: function sendCompressedClaims(bytes compressedClaimCalls) returns()
func (_Claimcompressor *ClaimcompressorTransactor) SendCompressedClaims(opts *bind.TransactOpts, compressedClaimCalls []byte) (*types.Transaction, error) {
	return _Claimcompressor.contract.Transact(opts, "sendCompressedClaims", compressedClaimCalls)
}

// SendCompressedClaims is a paid mutator transaction binding the contract method 0x97b1410f.
//
// Solidity: function sendCompressedClaims(bytes compressedClaimCalls) returns()
func (_Claimcompressor *ClaimcompressorSession) SendCompressedClaims(compressedClaimCalls []byte) (*types.Transaction, error) {
	return _Claimcompressor.Contract.SendCompressedClaims(&_Claimcompressor.TransactOpts, compressedClaimCalls)
}

// SendCompressedClaims is a paid mutator transaction binding the contract method 0x97b1410f.
//
// Solidity: function sendCompressedClaims(bytes compressedClaimCalls) returns()
func (_Claimcompressor *ClaimcompressorTransactorSession) SendCompressedClaims(compressedClaimCalls []byte) (*types.Transaction, error) {
	return _Claimcompressor.Contract.SendCompressedClaims(&_Claimcompressor.TransactOpts, compressedClaimCalls)
}

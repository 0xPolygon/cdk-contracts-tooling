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
	Bin: "0x60c060405234801561000f575f80fd5b50604051610de8380380610de883398101604081905261002e9161004a565b6001600160a01b0390911660805263ffffffff1660a052610095565b5f806040838503121561005b575f80fd5b82516001600160a01b0381168114610071575f80fd5b602084015190925063ffffffff8116811461008a575f80fd5b809150509250929050565b60805160a051610d326100b65f395f61043201525f6104540152610d325ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c806304e5557b1461003857806397b1410f14610061575b5f80fd5b61004b6100463660046105fc565b610076565b604051610058919061069a565b60405180910390f35b61007461006f3660046106ea565b61042b565b005b60605f8585604051602001610095929190918252602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905290505f805b8481101561041f575f8686838181106100e3576100e3610756565b90506020028101906100f59190610783565b6100fe90610972565b90505f825f03610110575060206101a0565b5f5b602081101561019e578888610128600187610a68565b81811061013757610137610756565b90506020028101906101499190610783565b816020811061015a5761015a610756565b6020020135835f0151826020811061017457610174610756565b60200201511461018c57610189816001610a81565b91505b8061019681610a94565b915050610112565b505b6040805160f883901b7fff000000000000000000000000000000000000000000000000000000000000001660208201528151600181830301815260219091019091525f5b8281101561023b5783518290826020811061020157610201610756565b6020020151604051602001610217929190610ae6565b6040516020818303038152906040529150808061023390610a94565b9150506101e4565b506020830151515f9250156102fd57831580610255575084155b1561026357602091506102f8565b5f5b60208110156102f657898961027b600188610a68565b81811061028a5761028a610756565b905060200281019061029c9190610783565b6104000181602081106102b1576102b1610756565b6020020135846020015182602081106102cc576102cc610756565b6020020151146102e4576102e1816001610a81565b92505b806102ee81610a94565b915050610265565b505b600194505b8082604051602001610310929190610b07565b60405160208183030381529060405290505f5b8281101561037c57818460200151826020811061034257610342610756565b6020020151604051602001610358929190610ae6565b6040516020818303038152906040529150808061037490610a94565b915050610323565b505f83610100015161038e575f610391565b60015b6040858101516060870151608088015160a089015160c08a015160e08b0151805187516103d399988c989081901c979096909590949093909291602001610b50565b604051602081830303815290604052905086816040516020016103f7929190610cce565b604051602081830303815290604052965050505050808061041790610a94565b9150506100c8565b50909695505050505050565b63ffffffff7f0000000000000000000000000000000000000000000000000000000000000000167f000000000000000000000000000000000000000000000000000000000000000063ccaa2d1163f5efcd798585602082610824376020828101610844376108a486905261092061090452604082015b8183018110156105f157803560f81c80156104c357600181146104e0576104f9565b856003538560081c6002538560101c6001538560181c5f536104f9565b846003538460081c6002538460101c6001538460181c5f535b5060028101906004906001013560f81c60200280838337909101600181019161040001906020903560f81c02808383379190910190610400810190823560f81c9061041701536001820191506008826018830137600882019150606081019050600482601c830137600482019150602081019050601482600c83013760149182019160408201918390604c01376014820191506020810190506020828237602082013560e01c60408201819052602490920191606090910190808383375f9181019190915290810190601f808216602003160161094401622625a05a10156105df575f80fd5b5f80825f808b621e8480f150506104a1565b505050505050505050565b5f805f806060858703121561060f575f80fd5b8435935060208501359250604085013567ffffffffffffffff80821115610634575f80fd5b818701915087601f830112610647575f80fd5b813581811115610655575f80fd5b8860208260051b8501011115610669575f80fd5b95989497505060200194505050565b5f5b8381101561069257818101518382015260200161067a565b50505f910152565b602081525f82518060208401526106b8816040850160208701610678565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b5f80602083850312156106fb575f80fd5b823567ffffffffffffffff80821115610712575f80fd5b818501915085601f830112610725575f80fd5b813581811115610733575f80fd5b866020828501011115610744575f80fd5b60209290920196919550909350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f82357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7218336030181126107b5575f80fd5b9190910192915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051610120810167ffffffffffffffff81118282101715610810576108106107bf565b60405290565b5f82601f830112610825575f80fd5b60405161040080820182811067ffffffffffffffff8211171561084a5761084a6107bf565b6040528301818582111561085c575f80fd5b845b8281101561087657803582526020918201910161085e565b509195945050505050565b803563ffffffff81168114610894575f80fd5b919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610894575f80fd5b5f82601f8301126108cb575f80fd5b813567ffffffffffffffff808211156108e6576108e66107bf565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561092c5761092c6107bf565b81604052838152866020858801011115610944575f80fd5b836020870160208301375f602085830101528094505050505092915050565b80358015158114610894575f80fd5b5f6108e08236031215610983575f80fd5b61098b6107ec565b6109953684610816565b81526109a5366104008501610816565b602082015261080083013560408201526109c26108208401610881565b60608201526109d46108408401610899565b60808201526109e66108608401610899565b60a082015261088083013560c08201526108a083013567ffffffffffffffff811115610a10575f80fd5b610a1c368286016108bc565b60e083015250610a2f6108c08401610963565b61010082015292915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b81810381811115610a7b57610a7b610a3b565b92915050565b80820180821115610a7b57610a7b610a3b565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610ac457610ac4610a3b565b5060010190565b5f8151610adc818560208601610678565b9290920192915050565b5f8351610af7818460208801610678565b9190910191825250602001919050565b5f8351610b18818460208801610678565b60f89390931b7fff00000000000000000000000000000000000000000000000000000000000000169190920190815260010192915050565b7fff000000000000000000000000000000000000000000000000000000000000008b60f81b1681525f8a51610b8c816001850160208f01610678565b7fff000000000000000000000000000000000000000000000000000000000000008b60f81b16600182850101527fffffffffffffffff0000000000000000000000000000000000000000000000008a60c01b16600282850101527fffffffff000000000000000000000000000000000000000000000000000000008960e01b16600a82850101527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008860601b16600e8285010152610c736022828501018860601b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000169052565b8560368285010152610cae6056828501018660e01b7fffffffff00000000000000000000000000000000000000000000000000000000169052565b610cbd605a8285010185610acb565b9d9c50505050505050505050505050565b5f8351610cdf818460208801610678565b835190830190610cf3818360208801610678565b0194935050505056fea2646970667358221220e4d58efa3ec5d8956396945aa120792a5a8040c43e9137f90e0c7143e51882e164736f6c63430008140033",
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

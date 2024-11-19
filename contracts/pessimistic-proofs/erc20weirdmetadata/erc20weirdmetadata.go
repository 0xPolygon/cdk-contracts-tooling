// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20weirdmetadata

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

// Erc20weirdmetadataMetaData contains all meta data concerning the Erc20weirdmetadata contract.
var Erc20weirdmetadataMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"symbol_\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"decimals_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isRevert\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDecimals\",\"type\":\"uint256\"}],\"name\":\"setDecimals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleIsRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b5060405162000f7738038062000f77833981016040819052620000339162000067565b60038390556004620000468382620001d4565b50600555506200029c9050565b634e487b7160e01b5f52604160045260245ffd5b5f805f606084860312156200007a575f80fd5b8351602080860151919450906001600160401b03808211156200009b575f80fd5b818701915087601f830112620000af575f80fd5b815181811115620000c457620000c462000053565b604051601f8201601f19908116603f01168101908382118183101715620000ef57620000ef62000053565b816040528281528a8684870101111562000107575f80fd5b5f93505b828410156200012a57848401860151818501870152928501926200010b565b5f868483010152809750505050505050604084015190509250925092565b600181811c908216806200015d57607f821691505b6020821081036200017c57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115620001cf575f81815260208120601f850160051c81016020861015620001aa5750805b601f850160051c820191505b81811015620001cb57828155600101620001b6565b5050505b505050565b81516001600160401b03811115620001f057620001f062000053565b620002088162000201845462000148565b8462000182565b602080601f8311600181146200023e575f8415620002265750858301515b5f19600386901b1c1916600185901b178555620001cb565b5f85815260208120601f198616915b828110156200026e578886015182559484019460019091019084016200024d565b50858210156200028c57878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b610ccd80620002aa5f395ff3fe608060405234801561000f575f80fd5b50600436106100f0575f3560e01c80633ccafd111161009357806395d89b411161006357806395d89b4114610207578063a457c2d71461021c578063a9059cbb1461022f578063dd62ed3e14610242575f80fd5b80633ccafd111461019f57806340c10f19146101ac57806370a08231146101bf5780638c8885c8146101f4575f80fd5b806318160ddd116100ce57806318160ddd1461016957806323b872dd14610171578063313ce56714610184578063395093511461018c575f80fd5b806306fdde03146100f4578063095ea7b31461010f57806315a87bee14610132575b5f80fd5b6100fc610287565b6040519081526020015b60405180910390f35b61012261011d366004610adc565b61029f565b6040519015158152602001610106565b610167600680547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00811660ff90911615179055565b005b6002546100fc565b61012261017f366004610b04565b6102b8565b6100fc6102db565b61012261019a366004610adc565b6102f3565b6006546101229060ff1681565b6101676101ba366004610adc565b61033e565b6100fc6101cd366004610b3d565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205490565b610167610202366004610b5d565b600555565b61020f61034c565b6040516101069190610b74565b61012261022a366004610adc565b6103ec565b61012261023d366004610adc565b6104c1565b6100fc610250366004610bdd565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260016020908152604080832093909416825291909152205490565b6006545f9060ff1615610298575f80fd5b5060035490565b5f336102ac8185856104ce565b60019150505b92915050565b5f336102c5858285610680565b6102d0858585610756565b506001949350505050565b6006545f9060ff16156102ec575f80fd5b5060055490565b335f81815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091906102ac9082908690610339908790610c0e565b6104ce565b61034882826109c3565b5050565b60065460609060ff161561035e575f80fd5b6004805461036b90610c46565b80601f016020809104026020016040519081016040528092919081815260200182805461039790610c46565b80156103e25780601f106103b9576101008083540402835291602001916103e2565b820191905f5260205f20905b8154815290600101906020018083116103c557829003601f168201915b5050505050905090565b335f81815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152812054909190838110156104b4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f00000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6102d082868684036104ce565b5f336102ac818585610756565b73ffffffffffffffffffffffffffffffffffffffff8316610570576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f726573730000000000000000000000000000000000000000000000000000000060648201526084016104ab565b73ffffffffffffffffffffffffffffffffffffffff8216610613576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f737300000000000000000000000000000000000000000000000000000000000060648201526084016104ab565b73ffffffffffffffffffffffffffffffffffffffff8381165f8181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146107505781811015610743576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016104ab565b61075084848484036104ce565b50505050565b73ffffffffffffffffffffffffffffffffffffffff83166107f9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f647265737300000000000000000000000000000000000000000000000000000060648201526084016104ab565b73ffffffffffffffffffffffffffffffffffffffff821661089c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f657373000000000000000000000000000000000000000000000000000000000060648201526084016104ab565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526020819052604090205481811015610951576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e6365000000000000000000000000000000000000000000000000000060648201526084016104ab565b73ffffffffffffffffffffffffffffffffffffffff8481165f81815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3610750565b73ffffffffffffffffffffffffffffffffffffffff8216610a40576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016104ab565b8060025f828254610a519190610c0e565b909155505073ffffffffffffffffffffffffffffffffffffffff82165f81815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610ad7575f80fd5b919050565b5f8060408385031215610aed575f80fd5b610af683610ab4565b946020939093013593505050565b5f805f60608486031215610b16575f80fd5b610b1f84610ab4565b9250610b2d60208501610ab4565b9150604084013590509250925092565b5f60208284031215610b4d575f80fd5b610b5682610ab4565b9392505050565b5f60208284031215610b6d575f80fd5b5035919050565b5f6020808352835180828501525f5b81811015610b9f57858101830151858201604001528201610b83565b505f6040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b5f8060408385031215610bee575f80fd5b610bf783610ab4565b9150610c0560208401610ab4565b90509250929050565b808201808211156102b2577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b600181811c90821680610c5a57607f821691505b602082108103610c91577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5091905056fea2646970667358221220a4c8ff07772525257fa614b3e6a0f3f87ccb3bb53c9540dc842cd99ccc40216b64736f6c63430008140033",
}

// Erc20weirdmetadataABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20weirdmetadataMetaData.ABI instead.
var Erc20weirdmetadataABI = Erc20weirdmetadataMetaData.ABI

// Erc20weirdmetadataBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Erc20weirdmetadataMetaData.Bin instead.
var Erc20weirdmetadataBin = Erc20weirdmetadataMetaData.Bin

// DeployErc20weirdmetadata deploys a new Ethereum contract, binding an instance of Erc20weirdmetadata to it.
func DeployErc20weirdmetadata(auth *bind.TransactOpts, backend bind.ContractBackend, name_ [32]byte, symbol_ []byte, decimals_ *big.Int) (common.Address, *types.Transaction, *Erc20weirdmetadata, error) {
	parsed, err := Erc20weirdmetadataMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Erc20weirdmetadataBin), backend, name_, symbol_, decimals_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc20weirdmetadata{Erc20weirdmetadataCaller: Erc20weirdmetadataCaller{contract: contract}, Erc20weirdmetadataTransactor: Erc20weirdmetadataTransactor{contract: contract}, Erc20weirdmetadataFilterer: Erc20weirdmetadataFilterer{contract: contract}}, nil
}

// Erc20weirdmetadata is an auto generated Go binding around an Ethereum contract.
type Erc20weirdmetadata struct {
	Erc20weirdmetadataCaller     // Read-only binding to the contract
	Erc20weirdmetadataTransactor // Write-only binding to the contract
	Erc20weirdmetadataFilterer   // Log filterer for contract events
}

// Erc20weirdmetadataCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20weirdmetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20weirdmetadataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20weirdmetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20weirdmetadataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20weirdmetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20weirdmetadataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20weirdmetadataSession struct {
	Contract     *Erc20weirdmetadata // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Erc20weirdmetadataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20weirdmetadataCallerSession struct {
	Contract *Erc20weirdmetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// Erc20weirdmetadataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20weirdmetadataTransactorSession struct {
	Contract     *Erc20weirdmetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// Erc20weirdmetadataRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20weirdmetadataRaw struct {
	Contract *Erc20weirdmetadata // Generic contract binding to access the raw methods on
}

// Erc20weirdmetadataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20weirdmetadataCallerRaw struct {
	Contract *Erc20weirdmetadataCaller // Generic read-only contract binding to access the raw methods on
}

// Erc20weirdmetadataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20weirdmetadataTransactorRaw struct {
	Contract *Erc20weirdmetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20weirdmetadata creates a new instance of Erc20weirdmetadata, bound to a specific deployed contract.
func NewErc20weirdmetadata(address common.Address, backend bind.ContractBackend) (*Erc20weirdmetadata, error) {
	contract, err := bindErc20weirdmetadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20weirdmetadata{Erc20weirdmetadataCaller: Erc20weirdmetadataCaller{contract: contract}, Erc20weirdmetadataTransactor: Erc20weirdmetadataTransactor{contract: contract}, Erc20weirdmetadataFilterer: Erc20weirdmetadataFilterer{contract: contract}}, nil
}

// NewErc20weirdmetadataCaller creates a new read-only instance of Erc20weirdmetadata, bound to a specific deployed contract.
func NewErc20weirdmetadataCaller(address common.Address, caller bind.ContractCaller) (*Erc20weirdmetadataCaller, error) {
	contract, err := bindErc20weirdmetadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20weirdmetadataCaller{contract: contract}, nil
}

// NewErc20weirdmetadataTransactor creates a new write-only instance of Erc20weirdmetadata, bound to a specific deployed contract.
func NewErc20weirdmetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc20weirdmetadataTransactor, error) {
	contract, err := bindErc20weirdmetadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20weirdmetadataTransactor{contract: contract}, nil
}

// NewErc20weirdmetadataFilterer creates a new log filterer instance of Erc20weirdmetadata, bound to a specific deployed contract.
func NewErc20weirdmetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc20weirdmetadataFilterer, error) {
	contract, err := bindErc20weirdmetadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20weirdmetadataFilterer{contract: contract}, nil
}

// bindErc20weirdmetadata binds a generic wrapper to an already deployed contract.
func bindErc20weirdmetadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc20weirdmetadataMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20weirdmetadata *Erc20weirdmetadataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20weirdmetadata.Contract.Erc20weirdmetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20weirdmetadata *Erc20weirdmetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20weirdmetadata.Contract.Erc20weirdmetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20weirdmetadata *Erc20weirdmetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20weirdmetadata.Contract.Erc20weirdmetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20weirdmetadata *Erc20weirdmetadataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20weirdmetadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20weirdmetadata *Erc20weirdmetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20weirdmetadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20weirdmetadata *Erc20weirdmetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20weirdmetadata.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20weirdmetadata *Erc20weirdmetadataCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20weirdmetadata.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20weirdmetadata *Erc20weirdmetadataSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc20weirdmetadata.Contract.Allowance(&_Erc20weirdmetadata.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20weirdmetadata *Erc20weirdmetadataCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc20weirdmetadata.Contract.Allowance(&_Erc20weirdmetadata.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20weirdmetadata *Erc20weirdmetadataCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20weirdmetadata.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20weirdmetadata *Erc20weirdmetadataSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Erc20weirdmetadata.Contract.BalanceOf(&_Erc20weirdmetadata.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20weirdmetadata *Erc20weirdmetadataCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Erc20weirdmetadata.Contract.BalanceOf(&_Erc20weirdmetadata.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Erc20weirdmetadata *Erc20weirdmetadataCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20weirdmetadata.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Erc20weirdmetadata *Erc20weirdmetadataSession) Decimals() (*big.Int, error) {
	return _Erc20weirdmetadata.Contract.Decimals(&_Erc20weirdmetadata.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Erc20weirdmetadata *Erc20weirdmetadataCallerSession) Decimals() (*big.Int, error) {
	return _Erc20weirdmetadata.Contract.Decimals(&_Erc20weirdmetadata.CallOpts)
}

// IsRevert is a free data retrieval call binding the contract method 0x3ccafd11.
//
// Solidity: function isRevert() view returns(bool)
func (_Erc20weirdmetadata *Erc20weirdmetadataCaller) IsRevert(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Erc20weirdmetadata.contract.Call(opts, &out, "isRevert")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRevert is a free data retrieval call binding the contract method 0x3ccafd11.
//
// Solidity: function isRevert() view returns(bool)
func (_Erc20weirdmetadata *Erc20weirdmetadataSession) IsRevert() (bool, error) {
	return _Erc20weirdmetadata.Contract.IsRevert(&_Erc20weirdmetadata.CallOpts)
}

// IsRevert is a free data retrieval call binding the contract method 0x3ccafd11.
//
// Solidity: function isRevert() view returns(bool)
func (_Erc20weirdmetadata *Erc20weirdmetadataCallerSession) IsRevert() (bool, error) {
	return _Erc20weirdmetadata.Contract.IsRevert(&_Erc20weirdmetadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(bytes32)
func (_Erc20weirdmetadata *Erc20weirdmetadataCaller) Name(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Erc20weirdmetadata.contract.Call(opts, &out, "name")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(bytes32)
func (_Erc20weirdmetadata *Erc20weirdmetadataSession) Name() ([32]byte, error) {
	return _Erc20weirdmetadata.Contract.Name(&_Erc20weirdmetadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(bytes32)
func (_Erc20weirdmetadata *Erc20weirdmetadataCallerSession) Name() ([32]byte, error) {
	return _Erc20weirdmetadata.Contract.Name(&_Erc20weirdmetadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(bytes)
func (_Erc20weirdmetadata *Erc20weirdmetadataCaller) Symbol(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Erc20weirdmetadata.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(bytes)
func (_Erc20weirdmetadata *Erc20weirdmetadataSession) Symbol() ([]byte, error) {
	return _Erc20weirdmetadata.Contract.Symbol(&_Erc20weirdmetadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(bytes)
func (_Erc20weirdmetadata *Erc20weirdmetadataCallerSession) Symbol() ([]byte, error) {
	return _Erc20weirdmetadata.Contract.Symbol(&_Erc20weirdmetadata.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20weirdmetadata *Erc20weirdmetadataCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20weirdmetadata.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20weirdmetadata *Erc20weirdmetadataSession) TotalSupply() (*big.Int, error) {
	return _Erc20weirdmetadata.Contract.TotalSupply(&_Erc20weirdmetadata.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20weirdmetadata *Erc20weirdmetadataCallerSession) TotalSupply() (*big.Int, error) {
	return _Erc20weirdmetadata.Contract.TotalSupply(&_Erc20weirdmetadata.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20weirdmetadata *Erc20weirdmetadataTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20weirdmetadata.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20weirdmetadata *Erc20weirdmetadataSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20weirdmetadata.Contract.Approve(&_Erc20weirdmetadata.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20weirdmetadata *Erc20weirdmetadataTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20weirdmetadata.Contract.Approve(&_Erc20weirdmetadata.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Erc20weirdmetadata *Erc20weirdmetadataTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Erc20weirdmetadata.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Erc20weirdmetadata *Erc20weirdmetadataSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Erc20weirdmetadata.Contract.DecreaseAllowance(&_Erc20weirdmetadata.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Erc20weirdmetadata *Erc20weirdmetadataTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Erc20weirdmetadata.Contract.DecreaseAllowance(&_Erc20weirdmetadata.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Erc20weirdmetadata *Erc20weirdmetadataTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Erc20weirdmetadata.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Erc20weirdmetadata *Erc20weirdmetadataSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Erc20weirdmetadata.Contract.IncreaseAllowance(&_Erc20weirdmetadata.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Erc20weirdmetadata *Erc20weirdmetadataTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Erc20weirdmetadata.Contract.IncreaseAllowance(&_Erc20weirdmetadata.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Erc20weirdmetadata *Erc20weirdmetadataTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20weirdmetadata.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Erc20weirdmetadata *Erc20weirdmetadataSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20weirdmetadata.Contract.Mint(&_Erc20weirdmetadata.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Erc20weirdmetadata *Erc20weirdmetadataTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20weirdmetadata.Contract.Mint(&_Erc20weirdmetadata.TransactOpts, account, amount)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x8c8885c8.
//
// Solidity: function setDecimals(uint256 newDecimals) returns()
func (_Erc20weirdmetadata *Erc20weirdmetadataTransactor) SetDecimals(opts *bind.TransactOpts, newDecimals *big.Int) (*types.Transaction, error) {
	return _Erc20weirdmetadata.contract.Transact(opts, "setDecimals", newDecimals)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x8c8885c8.
//
// Solidity: function setDecimals(uint256 newDecimals) returns()
func (_Erc20weirdmetadata *Erc20weirdmetadataSession) SetDecimals(newDecimals *big.Int) (*types.Transaction, error) {
	return _Erc20weirdmetadata.Contract.SetDecimals(&_Erc20weirdmetadata.TransactOpts, newDecimals)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x8c8885c8.
//
// Solidity: function setDecimals(uint256 newDecimals) returns()
func (_Erc20weirdmetadata *Erc20weirdmetadataTransactorSession) SetDecimals(newDecimals *big.Int) (*types.Transaction, error) {
	return _Erc20weirdmetadata.Contract.SetDecimals(&_Erc20weirdmetadata.TransactOpts, newDecimals)
}

// ToggleIsRevert is a paid mutator transaction binding the contract method 0x15a87bee.
//
// Solidity: function toggleIsRevert() returns()
func (_Erc20weirdmetadata *Erc20weirdmetadataTransactor) ToggleIsRevert(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20weirdmetadata.contract.Transact(opts, "toggleIsRevert")
}

// ToggleIsRevert is a paid mutator transaction binding the contract method 0x15a87bee.
//
// Solidity: function toggleIsRevert() returns()
func (_Erc20weirdmetadata *Erc20weirdmetadataSession) ToggleIsRevert() (*types.Transaction, error) {
	return _Erc20weirdmetadata.Contract.ToggleIsRevert(&_Erc20weirdmetadata.TransactOpts)
}

// ToggleIsRevert is a paid mutator transaction binding the contract method 0x15a87bee.
//
// Solidity: function toggleIsRevert() returns()
func (_Erc20weirdmetadata *Erc20weirdmetadataTransactorSession) ToggleIsRevert() (*types.Transaction, error) {
	return _Erc20weirdmetadata.Contract.ToggleIsRevert(&_Erc20weirdmetadata.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Erc20weirdmetadata *Erc20weirdmetadataTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20weirdmetadata.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Erc20weirdmetadata *Erc20weirdmetadataSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20weirdmetadata.Contract.Transfer(&_Erc20weirdmetadata.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Erc20weirdmetadata *Erc20weirdmetadataTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20weirdmetadata.Contract.Transfer(&_Erc20weirdmetadata.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Erc20weirdmetadata *Erc20weirdmetadataTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20weirdmetadata.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Erc20weirdmetadata *Erc20weirdmetadataSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20weirdmetadata.Contract.TransferFrom(&_Erc20weirdmetadata.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Erc20weirdmetadata *Erc20weirdmetadataTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20weirdmetadata.Contract.TransferFrom(&_Erc20weirdmetadata.TransactOpts, from, to, amount)
}

// Erc20weirdmetadataApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc20weirdmetadata contract.
type Erc20weirdmetadataApprovalIterator struct {
	Event *Erc20weirdmetadataApproval // Event containing the contract specifics and raw log

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
func (it *Erc20weirdmetadataApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20weirdmetadataApproval)
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
		it.Event = new(Erc20weirdmetadataApproval)
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
func (it *Erc20weirdmetadataApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20weirdmetadataApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20weirdmetadataApproval represents a Approval event raised by the Erc20weirdmetadata contract.
type Erc20weirdmetadataApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20weirdmetadata *Erc20weirdmetadataFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Erc20weirdmetadataApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20weirdmetadata.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Erc20weirdmetadataApprovalIterator{contract: _Erc20weirdmetadata.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20weirdmetadata *Erc20weirdmetadataFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc20weirdmetadataApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20weirdmetadata.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20weirdmetadataApproval)
				if err := _Erc20weirdmetadata.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20weirdmetadata *Erc20weirdmetadataFilterer) ParseApproval(log types.Log) (*Erc20weirdmetadataApproval, error) {
	event := new(Erc20weirdmetadataApproval)
	if err := _Erc20weirdmetadata.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20weirdmetadataTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc20weirdmetadata contract.
type Erc20weirdmetadataTransferIterator struct {
	Event *Erc20weirdmetadataTransfer // Event containing the contract specifics and raw log

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
func (it *Erc20weirdmetadataTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20weirdmetadataTransfer)
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
		it.Event = new(Erc20weirdmetadataTransfer)
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
func (it *Erc20weirdmetadataTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20weirdmetadataTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20weirdmetadataTransfer represents a Transfer event raised by the Erc20weirdmetadata contract.
type Erc20weirdmetadataTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20weirdmetadata *Erc20weirdmetadataFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Erc20weirdmetadataTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20weirdmetadata.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc20weirdmetadataTransferIterator{contract: _Erc20weirdmetadata.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20weirdmetadata *Erc20weirdmetadataFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc20weirdmetadataTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20weirdmetadata.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20weirdmetadataTransfer)
				if err := _Erc20weirdmetadata.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20weirdmetadata *Erc20weirdmetadataFilterer) ParseTransfer(log types.Log) (*Erc20weirdmetadataTransfer, error) {
	event := new(Erc20weirdmetadataTransfer)
	if err := _Erc20weirdmetadata.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20invalidmetadata

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

// Erc20invalidmetadataMetaData contains all meta data concerning the Erc20invalidmetadata contract.
var Erc20invalidmetadataMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"symbol_\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"decimals_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isRevert\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDecimals\",\"type\":\"uint256\"}],\"name\":\"setDecimals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleIsRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000eab38038062000eab83398101604081905262000034916200006a565b60038390556004620000478382620001e2565b5060055550620002ae9050565b634e487b7160e01b600052604160045260246000fd5b6000806000606084860312156200008057600080fd5b8351602080860151919450906001600160401b0380821115620000a257600080fd5b818701915087601f830112620000b757600080fd5b815181811115620000cc57620000cc62000054565b604051601f8201601f19908116603f01168101908382118183101715620000f757620000f762000054565b816040528281528a868487010111156200011057600080fd5b600093505b8284101562000134578484018601518185018701529285019262000115565b6000868483010152809750505050505050604084015190509250925092565b600181811c908216806200016857607f821691505b6020821081036200018957634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620001dd57600081815260208120601f850160051c81016020861015620001b85750805b601f850160051c820191505b81811015620001d957828155600101620001c4565b5050505b505050565b81516001600160401b03811115620001fe57620001fe62000054565b62000216816200020f845462000153565b846200018f565b602080601f8311600181146200024e5760008415620002355750858301515b600019600386901b1c1916600185901b178555620001d9565b600085815260208120601f198616915b828110156200027f578886015182559484019460019091019084016200025e565b50858210156200029e5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b610bed80620002be6000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80633ccafd111161009757806395d89b411161006657806395d89b411461021b578063a457c2d714610220578063a9059cbb14610233578063dd62ed3e1461024657600080fd5b80633ccafd11146101b257806340c10f19146101bf57806370a08231146101d25780638c8885c81461020857600080fd5b806318160ddd116100d357806318160ddd1461017257806323b872dd14610184578063313ce56714610197578063395093511461019f57600080fd5b806306fdde03146100fa578063095ea7b31461011857806315a87bee1461013b575b600080fd5b61010261028c565b60405161010f9190610a14565b60405180910390f35b61012b610126366004610aa9565b610292565b604051901515815260200161010f565b610170600680547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00811660ff90911615179055565b005b6002545b60405190815260200161010f565b61012b610192366004610ad3565b6102ac565b6101766102d0565b61012b6101ad366004610aa9565b6102ea565b60065461012b9060ff1681565b6101706101cd366004610aa9565b610336565b6101766101e0366004610b0f565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b610170610216366004610b31565b600555565b610102005b61012b61022e366004610aa9565b610344565b61012b610241366004610aa9565b61041a565b610176610254366004610b4a565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b600a6000f35b6000336102a0818585610428565b60019150505b92915050565b6000336102ba8582856105db565b6102c58585856106b2565b506001949350505050565b60065460009060ff16156102e357600080fd5b5060055490565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091906102a09082908690610331908790610b7d565b610428565b6103408282610921565b5050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091908381101561040d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f00000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6102c58286868403610428565b6000336102a08185856106b2565b73ffffffffffffffffffffffffffffffffffffffff83166104ca576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610404565b73ffffffffffffffffffffffffffffffffffffffff821661056d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152608401610404565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146106ac578181101561069f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610404565b6106ac8484848403610428565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610755576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152608401610404565b73ffffffffffffffffffffffffffffffffffffffff82166107f8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152608401610404565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040902054818110156108ae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152608401610404565b73ffffffffffffffffffffffffffffffffffffffff848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a36106ac565b73ffffffffffffffffffffffffffffffffffffffff821661099e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610404565b80600260008282546109b09190610b7d565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b600060208083528351808285015260005b81811015610a4157858101830151858201604001528201610a25565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610aa457600080fd5b919050565b60008060408385031215610abc57600080fd5b610ac583610a80565b946020939093013593505050565b600080600060608486031215610ae857600080fd5b610af184610a80565b9250610aff60208501610a80565b9150604084013590509250925092565b600060208284031215610b2157600080fd5b610b2a82610a80565b9392505050565b600060208284031215610b4357600080fd5b5035919050565b60008060408385031215610b5d57600080fd5b610b6683610a80565b9150610b7460208401610a80565b90509250929050565b808201808211156102a6577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220943c63d74c498258be591aa5555259c45bbb236a3bbbe2f1f45b84a3598b760b64736f6c63430008140033",
}

// Erc20invalidmetadataABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20invalidmetadataMetaData.ABI instead.
var Erc20invalidmetadataABI = Erc20invalidmetadataMetaData.ABI

// Erc20invalidmetadataBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Erc20invalidmetadataMetaData.Bin instead.
var Erc20invalidmetadataBin = Erc20invalidmetadataMetaData.Bin

// DeployErc20invalidmetadata deploys a new Ethereum contract, binding an instance of Erc20invalidmetadata to it.
func DeployErc20invalidmetadata(auth *bind.TransactOpts, backend bind.ContractBackend, name_ [32]byte, symbol_ []byte, decimals_ *big.Int) (common.Address, *types.Transaction, *Erc20invalidmetadata, error) {
	parsed, err := Erc20invalidmetadataMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Erc20invalidmetadataBin), backend, name_, symbol_, decimals_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc20invalidmetadata{Erc20invalidmetadataCaller: Erc20invalidmetadataCaller{contract: contract}, Erc20invalidmetadataTransactor: Erc20invalidmetadataTransactor{contract: contract}, Erc20invalidmetadataFilterer: Erc20invalidmetadataFilterer{contract: contract}}, nil
}

// Erc20invalidmetadata is an auto generated Go binding around an Ethereum contract.
type Erc20invalidmetadata struct {
	Erc20invalidmetadataCaller     // Read-only binding to the contract
	Erc20invalidmetadataTransactor // Write-only binding to the contract
	Erc20invalidmetadataFilterer   // Log filterer for contract events
}

// Erc20invalidmetadataCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20invalidmetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20invalidmetadataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20invalidmetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20invalidmetadataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20invalidmetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20invalidmetadataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20invalidmetadataSession struct {
	Contract     *Erc20invalidmetadata // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Erc20invalidmetadataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20invalidmetadataCallerSession struct {
	Contract *Erc20invalidmetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// Erc20invalidmetadataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20invalidmetadataTransactorSession struct {
	Contract     *Erc20invalidmetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// Erc20invalidmetadataRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20invalidmetadataRaw struct {
	Contract *Erc20invalidmetadata // Generic contract binding to access the raw methods on
}

// Erc20invalidmetadataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20invalidmetadataCallerRaw struct {
	Contract *Erc20invalidmetadataCaller // Generic read-only contract binding to access the raw methods on
}

// Erc20invalidmetadataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20invalidmetadataTransactorRaw struct {
	Contract *Erc20invalidmetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20invalidmetadata creates a new instance of Erc20invalidmetadata, bound to a specific deployed contract.
func NewErc20invalidmetadata(address common.Address, backend bind.ContractBackend) (*Erc20invalidmetadata, error) {
	contract, err := bindErc20invalidmetadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20invalidmetadata{Erc20invalidmetadataCaller: Erc20invalidmetadataCaller{contract: contract}, Erc20invalidmetadataTransactor: Erc20invalidmetadataTransactor{contract: contract}, Erc20invalidmetadataFilterer: Erc20invalidmetadataFilterer{contract: contract}}, nil
}

// NewErc20invalidmetadataCaller creates a new read-only instance of Erc20invalidmetadata, bound to a specific deployed contract.
func NewErc20invalidmetadataCaller(address common.Address, caller bind.ContractCaller) (*Erc20invalidmetadataCaller, error) {
	contract, err := bindErc20invalidmetadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20invalidmetadataCaller{contract: contract}, nil
}

// NewErc20invalidmetadataTransactor creates a new write-only instance of Erc20invalidmetadata, bound to a specific deployed contract.
func NewErc20invalidmetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc20invalidmetadataTransactor, error) {
	contract, err := bindErc20invalidmetadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20invalidmetadataTransactor{contract: contract}, nil
}

// NewErc20invalidmetadataFilterer creates a new log filterer instance of Erc20invalidmetadata, bound to a specific deployed contract.
func NewErc20invalidmetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc20invalidmetadataFilterer, error) {
	contract, err := bindErc20invalidmetadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20invalidmetadataFilterer{contract: contract}, nil
}

// bindErc20invalidmetadata binds a generic wrapper to an already deployed contract.
func bindErc20invalidmetadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc20invalidmetadataMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20invalidmetadata *Erc20invalidmetadataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20invalidmetadata.Contract.Erc20invalidmetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20invalidmetadata *Erc20invalidmetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20invalidmetadata.Contract.Erc20invalidmetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20invalidmetadata *Erc20invalidmetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20invalidmetadata.Contract.Erc20invalidmetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20invalidmetadata *Erc20invalidmetadataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20invalidmetadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20invalidmetadata *Erc20invalidmetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20invalidmetadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20invalidmetadata *Erc20invalidmetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20invalidmetadata.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20invalidmetadata *Erc20invalidmetadataCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20invalidmetadata.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20invalidmetadata *Erc20invalidmetadataSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc20invalidmetadata.Contract.Allowance(&_Erc20invalidmetadata.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20invalidmetadata *Erc20invalidmetadataCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc20invalidmetadata.Contract.Allowance(&_Erc20invalidmetadata.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20invalidmetadata *Erc20invalidmetadataCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20invalidmetadata.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20invalidmetadata *Erc20invalidmetadataSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Erc20invalidmetadata.Contract.BalanceOf(&_Erc20invalidmetadata.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20invalidmetadata *Erc20invalidmetadataCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Erc20invalidmetadata.Contract.BalanceOf(&_Erc20invalidmetadata.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Erc20invalidmetadata *Erc20invalidmetadataCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20invalidmetadata.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Erc20invalidmetadata *Erc20invalidmetadataSession) Decimals() (*big.Int, error) {
	return _Erc20invalidmetadata.Contract.Decimals(&_Erc20invalidmetadata.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Erc20invalidmetadata *Erc20invalidmetadataCallerSession) Decimals() (*big.Int, error) {
	return _Erc20invalidmetadata.Contract.Decimals(&_Erc20invalidmetadata.CallOpts)
}

// IsRevert is a free data retrieval call binding the contract method 0x3ccafd11.
//
// Solidity: function isRevert() view returns(bool)
func (_Erc20invalidmetadata *Erc20invalidmetadataCaller) IsRevert(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Erc20invalidmetadata.contract.Call(opts, &out, "isRevert")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRevert is a free data retrieval call binding the contract method 0x3ccafd11.
//
// Solidity: function isRevert() view returns(bool)
func (_Erc20invalidmetadata *Erc20invalidmetadataSession) IsRevert() (bool, error) {
	return _Erc20invalidmetadata.Contract.IsRevert(&_Erc20invalidmetadata.CallOpts)
}

// IsRevert is a free data retrieval call binding the contract method 0x3ccafd11.
//
// Solidity: function isRevert() view returns(bool)
func (_Erc20invalidmetadata *Erc20invalidmetadataCallerSession) IsRevert() (bool, error) {
	return _Erc20invalidmetadata.Contract.IsRevert(&_Erc20invalidmetadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(bytes)
func (_Erc20invalidmetadata *Erc20invalidmetadataCaller) Name(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Erc20invalidmetadata.contract.Call(opts, &out, "name")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(bytes)
func (_Erc20invalidmetadata *Erc20invalidmetadataSession) Name() ([]byte, error) {
	return _Erc20invalidmetadata.Contract.Name(&_Erc20invalidmetadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(bytes)
func (_Erc20invalidmetadata *Erc20invalidmetadataCallerSession) Name() ([]byte, error) {
	return _Erc20invalidmetadata.Contract.Name(&_Erc20invalidmetadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(bytes)
func (_Erc20invalidmetadata *Erc20invalidmetadataCaller) Symbol(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Erc20invalidmetadata.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(bytes)
func (_Erc20invalidmetadata *Erc20invalidmetadataSession) Symbol() ([]byte, error) {
	return _Erc20invalidmetadata.Contract.Symbol(&_Erc20invalidmetadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(bytes)
func (_Erc20invalidmetadata *Erc20invalidmetadataCallerSession) Symbol() ([]byte, error) {
	return _Erc20invalidmetadata.Contract.Symbol(&_Erc20invalidmetadata.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20invalidmetadata *Erc20invalidmetadataCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20invalidmetadata.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20invalidmetadata *Erc20invalidmetadataSession) TotalSupply() (*big.Int, error) {
	return _Erc20invalidmetadata.Contract.TotalSupply(&_Erc20invalidmetadata.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20invalidmetadata *Erc20invalidmetadataCallerSession) TotalSupply() (*big.Int, error) {
	return _Erc20invalidmetadata.Contract.TotalSupply(&_Erc20invalidmetadata.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20invalidmetadata *Erc20invalidmetadataTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20invalidmetadata.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20invalidmetadata *Erc20invalidmetadataSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20invalidmetadata.Contract.Approve(&_Erc20invalidmetadata.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20invalidmetadata *Erc20invalidmetadataTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20invalidmetadata.Contract.Approve(&_Erc20invalidmetadata.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Erc20invalidmetadata *Erc20invalidmetadataTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Erc20invalidmetadata.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Erc20invalidmetadata *Erc20invalidmetadataSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Erc20invalidmetadata.Contract.DecreaseAllowance(&_Erc20invalidmetadata.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Erc20invalidmetadata *Erc20invalidmetadataTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Erc20invalidmetadata.Contract.DecreaseAllowance(&_Erc20invalidmetadata.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Erc20invalidmetadata *Erc20invalidmetadataTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Erc20invalidmetadata.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Erc20invalidmetadata *Erc20invalidmetadataSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Erc20invalidmetadata.Contract.IncreaseAllowance(&_Erc20invalidmetadata.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Erc20invalidmetadata *Erc20invalidmetadataTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Erc20invalidmetadata.Contract.IncreaseAllowance(&_Erc20invalidmetadata.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Erc20invalidmetadata *Erc20invalidmetadataTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20invalidmetadata.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Erc20invalidmetadata *Erc20invalidmetadataSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20invalidmetadata.Contract.Mint(&_Erc20invalidmetadata.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Erc20invalidmetadata *Erc20invalidmetadataTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20invalidmetadata.Contract.Mint(&_Erc20invalidmetadata.TransactOpts, account, amount)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x8c8885c8.
//
// Solidity: function setDecimals(uint256 newDecimals) returns()
func (_Erc20invalidmetadata *Erc20invalidmetadataTransactor) SetDecimals(opts *bind.TransactOpts, newDecimals *big.Int) (*types.Transaction, error) {
	return _Erc20invalidmetadata.contract.Transact(opts, "setDecimals", newDecimals)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x8c8885c8.
//
// Solidity: function setDecimals(uint256 newDecimals) returns()
func (_Erc20invalidmetadata *Erc20invalidmetadataSession) SetDecimals(newDecimals *big.Int) (*types.Transaction, error) {
	return _Erc20invalidmetadata.Contract.SetDecimals(&_Erc20invalidmetadata.TransactOpts, newDecimals)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x8c8885c8.
//
// Solidity: function setDecimals(uint256 newDecimals) returns()
func (_Erc20invalidmetadata *Erc20invalidmetadataTransactorSession) SetDecimals(newDecimals *big.Int) (*types.Transaction, error) {
	return _Erc20invalidmetadata.Contract.SetDecimals(&_Erc20invalidmetadata.TransactOpts, newDecimals)
}

// ToggleIsRevert is a paid mutator transaction binding the contract method 0x15a87bee.
//
// Solidity: function toggleIsRevert() returns()
func (_Erc20invalidmetadata *Erc20invalidmetadataTransactor) ToggleIsRevert(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20invalidmetadata.contract.Transact(opts, "toggleIsRevert")
}

// ToggleIsRevert is a paid mutator transaction binding the contract method 0x15a87bee.
//
// Solidity: function toggleIsRevert() returns()
func (_Erc20invalidmetadata *Erc20invalidmetadataSession) ToggleIsRevert() (*types.Transaction, error) {
	return _Erc20invalidmetadata.Contract.ToggleIsRevert(&_Erc20invalidmetadata.TransactOpts)
}

// ToggleIsRevert is a paid mutator transaction binding the contract method 0x15a87bee.
//
// Solidity: function toggleIsRevert() returns()
func (_Erc20invalidmetadata *Erc20invalidmetadataTransactorSession) ToggleIsRevert() (*types.Transaction, error) {
	return _Erc20invalidmetadata.Contract.ToggleIsRevert(&_Erc20invalidmetadata.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Erc20invalidmetadata *Erc20invalidmetadataTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20invalidmetadata.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Erc20invalidmetadata *Erc20invalidmetadataSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20invalidmetadata.Contract.Transfer(&_Erc20invalidmetadata.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Erc20invalidmetadata *Erc20invalidmetadataTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20invalidmetadata.Contract.Transfer(&_Erc20invalidmetadata.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Erc20invalidmetadata *Erc20invalidmetadataTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20invalidmetadata.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Erc20invalidmetadata *Erc20invalidmetadataSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20invalidmetadata.Contract.TransferFrom(&_Erc20invalidmetadata.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Erc20invalidmetadata *Erc20invalidmetadataTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20invalidmetadata.Contract.TransferFrom(&_Erc20invalidmetadata.TransactOpts, from, to, amount)
}

// Erc20invalidmetadataApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc20invalidmetadata contract.
type Erc20invalidmetadataApprovalIterator struct {
	Event *Erc20invalidmetadataApproval // Event containing the contract specifics and raw log

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
func (it *Erc20invalidmetadataApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20invalidmetadataApproval)
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
		it.Event = new(Erc20invalidmetadataApproval)
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
func (it *Erc20invalidmetadataApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20invalidmetadataApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20invalidmetadataApproval represents a Approval event raised by the Erc20invalidmetadata contract.
type Erc20invalidmetadataApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20invalidmetadata *Erc20invalidmetadataFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Erc20invalidmetadataApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20invalidmetadata.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Erc20invalidmetadataApprovalIterator{contract: _Erc20invalidmetadata.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20invalidmetadata *Erc20invalidmetadataFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc20invalidmetadataApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20invalidmetadata.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20invalidmetadataApproval)
				if err := _Erc20invalidmetadata.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Erc20invalidmetadata *Erc20invalidmetadataFilterer) ParseApproval(log types.Log) (*Erc20invalidmetadataApproval, error) {
	event := new(Erc20invalidmetadataApproval)
	if err := _Erc20invalidmetadata.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20invalidmetadataTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc20invalidmetadata contract.
type Erc20invalidmetadataTransferIterator struct {
	Event *Erc20invalidmetadataTransfer // Event containing the contract specifics and raw log

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
func (it *Erc20invalidmetadataTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20invalidmetadataTransfer)
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
		it.Event = new(Erc20invalidmetadataTransfer)
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
func (it *Erc20invalidmetadataTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20invalidmetadataTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20invalidmetadataTransfer represents a Transfer event raised by the Erc20invalidmetadata contract.
type Erc20invalidmetadataTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20invalidmetadata *Erc20invalidmetadataFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Erc20invalidmetadataTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20invalidmetadata.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc20invalidmetadataTransferIterator{contract: _Erc20invalidmetadata.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20invalidmetadata *Erc20invalidmetadataFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc20invalidmetadataTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20invalidmetadata.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20invalidmetadataTransfer)
				if err := _Erc20invalidmetadata.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Erc20invalidmetadata *Erc20invalidmetadataFilterer) ParseTransfer(log types.Log) (*Erc20invalidmetadataTransfer, error) {
	event := new(Erc20invalidmetadataTransfer)
	if err := _Erc20invalidmetadata.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

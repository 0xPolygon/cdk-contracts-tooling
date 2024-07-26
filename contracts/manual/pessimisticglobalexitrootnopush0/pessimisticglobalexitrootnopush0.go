// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pessimisticglobalexitrootnopush0

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

// Pessimisticglobalexitrootnopush0MetaData contains all meta data concerning the Pessimisticglobalexitrootnopush0 contract.
var Pessimisticglobalexitrootnopush0MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlOnlyCanRenounceRolesForSelf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressDoNotHaveRequiredRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_newRoot\",\"type\":\"bytes32\"}],\"name\":\"updateGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b506040516108e63803806108e683398101604081905261002f91610243565b6001600160a01b0381166080526100747f7b95520991dfda409891be0afa2635b63540f92ee996fda0bf695a166e5c51766000805160206108c68339815191526100b2565b61008c6000805160206108c6833981519152806100b2565b6100a46000805160206108c6833981519152336100fd565b6100ac610183565b50610273565b600082815260346020526040808220600101805490849055905190918391839186917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a4505050565b60008281526034602090815260408083206001600160a01b038516845290915290205460ff1661017f5760008281526034602090815260408083206001600160a01b0385168085529252808320805460ff1916600117905551339285917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45b5050565b600054610100900460ff16156101ef5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161015610241576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b60006020828403121561025557600080fd5b81516001600160a01b038116811461026c57600080fd5b9392505050565b608051610631610295600039600081816101de015261027a01526106316000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c806336568abe11610081578063a217fddf1161005b578063a217fddf146101d1578063a3c573eb146101d9578063d547741f1461022557600080fd5b806336568abe1461015557806348da4bdd1461016857806391d148541461017b57600080fd5b8063257b3632116100b2578063257b36321461010d5780632f2ff15d1461012d57806333d6247d1461014257600080fd5b806301fd9044146100ce578063248a9ca3146100ea575b600080fd5b6100d760665481565b6040519081526020015b60405180910390f35b6100d76100f8366004610599565b60009081526034602052604090206001015490565b6100d761011b366004610599565b60656020526000908152604090205481565b61014061013b3660046105b2565b610238565b005b610140610150366004610599565b610262565b6101406101633660046105b2565b6102d6565b610140610176366004610599565b610333565b6101c16101893660046105b2565b600091825260346020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b60405190151581526020016100e1565b6100d7600081565b6102007f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100e1565b6101406102333660046105b2565b610386565b600082815260346020526040902060010154610253816103ab565b61025d83836103b8565b505050565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146102d1576040517fb49365dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606655565b73ffffffffffffffffffffffffffffffffffffffff81163314610325576040517f5a568e6800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61032f8282610475565b5050565b7f7b95520991dfda409891be0afa2635b63540f92ee996fda0bf695a166e5c517661035d816103ab565b600082815260656020526040812054900361032f57506000908152606560205260409020429055565b6000828152603460205260409020600101546103a1816103ab565b61025d8383610475565b6103b58133610530565b50565b600082815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1661032f57600082815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905551339285917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45050565b600082815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff161561032f57600082815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b600082815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1661032f576040517fec2b7c3e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000602082840312156105ab57600080fd5b5035919050565b600080604083850312156105c557600080fd5b82359150602083013573ffffffffffffffffffffffffffffffffffffffff811681146105f057600080fd5b80915050925092905056fea26469706673582212202d04d937f4aab58050111c2239443a8b7008cb880dd00c4fa495c5bac0ca556564736f6c634300081400335ee7b1f4df16da4d137fdcc53b4b785d709ccbf7d129b16d96c5c3bedc99aef7",
}

// Pessimisticglobalexitrootnopush0ABI is the input ABI used to generate the binding from.
// Deprecated: Use Pessimisticglobalexitrootnopush0MetaData.ABI instead.
var Pessimisticglobalexitrootnopush0ABI = Pessimisticglobalexitrootnopush0MetaData.ABI

// Pessimisticglobalexitrootnopush0Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Pessimisticglobalexitrootnopush0MetaData.Bin instead.
var Pessimisticglobalexitrootnopush0Bin = Pessimisticglobalexitrootnopush0MetaData.Bin

// DeployPessimisticglobalexitrootnopush0 deploys a new Ethereum contract, binding an instance of Pessimisticglobalexitrootnopush0 to it.
func DeployPessimisticglobalexitrootnopush0(auth *bind.TransactOpts, backend bind.ContractBackend, _bridgeAddress common.Address) (common.Address, *types.Transaction, *Pessimisticglobalexitrootnopush0, error) {
	parsed, err := Pessimisticglobalexitrootnopush0MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Pessimisticglobalexitrootnopush0Bin), backend, _bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pessimisticglobalexitrootnopush0{Pessimisticglobalexitrootnopush0Caller: Pessimisticglobalexitrootnopush0Caller{contract: contract}, Pessimisticglobalexitrootnopush0Transactor: Pessimisticglobalexitrootnopush0Transactor{contract: contract}, Pessimisticglobalexitrootnopush0Filterer: Pessimisticglobalexitrootnopush0Filterer{contract: contract}}, nil
}

// Pessimisticglobalexitrootnopush0 is an auto generated Go binding around an Ethereum contract.
type Pessimisticglobalexitrootnopush0 struct {
	Pessimisticglobalexitrootnopush0Caller     // Read-only binding to the contract
	Pessimisticglobalexitrootnopush0Transactor // Write-only binding to the contract
	Pessimisticglobalexitrootnopush0Filterer   // Log filterer for contract events
}

// Pessimisticglobalexitrootnopush0Caller is an auto generated read-only Go binding around an Ethereum contract.
type Pessimisticglobalexitrootnopush0Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Pessimisticglobalexitrootnopush0Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Pessimisticglobalexitrootnopush0Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Pessimisticglobalexitrootnopush0Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Pessimisticglobalexitrootnopush0Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Pessimisticglobalexitrootnopush0Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Pessimisticglobalexitrootnopush0Session struct {
	Contract     *Pessimisticglobalexitrootnopush0 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                     // Call options to use throughout this session
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// Pessimisticglobalexitrootnopush0CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Pessimisticglobalexitrootnopush0CallerSession struct {
	Contract *Pessimisticglobalexitrootnopush0Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                           // Call options to use throughout this session
}

// Pessimisticglobalexitrootnopush0TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Pessimisticglobalexitrootnopush0TransactorSession struct {
	Contract     *Pessimisticglobalexitrootnopush0Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                           // Transaction auth options to use throughout this session
}

// Pessimisticglobalexitrootnopush0Raw is an auto generated low-level Go binding around an Ethereum contract.
type Pessimisticglobalexitrootnopush0Raw struct {
	Contract *Pessimisticglobalexitrootnopush0 // Generic contract binding to access the raw methods on
}

// Pessimisticglobalexitrootnopush0CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Pessimisticglobalexitrootnopush0CallerRaw struct {
	Contract *Pessimisticglobalexitrootnopush0Caller // Generic read-only contract binding to access the raw methods on
}

// Pessimisticglobalexitrootnopush0TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Pessimisticglobalexitrootnopush0TransactorRaw struct {
	Contract *Pessimisticglobalexitrootnopush0Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPessimisticglobalexitrootnopush0 creates a new instance of Pessimisticglobalexitrootnopush0, bound to a specific deployed contract.
func NewPessimisticglobalexitrootnopush0(address common.Address, backend bind.ContractBackend) (*Pessimisticglobalexitrootnopush0, error) {
	contract, err := bindPessimisticglobalexitrootnopush0(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pessimisticglobalexitrootnopush0{Pessimisticglobalexitrootnopush0Caller: Pessimisticglobalexitrootnopush0Caller{contract: contract}, Pessimisticglobalexitrootnopush0Transactor: Pessimisticglobalexitrootnopush0Transactor{contract: contract}, Pessimisticglobalexitrootnopush0Filterer: Pessimisticglobalexitrootnopush0Filterer{contract: contract}}, nil
}

// NewPessimisticglobalexitrootnopush0Caller creates a new read-only instance of Pessimisticglobalexitrootnopush0, bound to a specific deployed contract.
func NewPessimisticglobalexitrootnopush0Caller(address common.Address, caller bind.ContractCaller) (*Pessimisticglobalexitrootnopush0Caller, error) {
	contract, err := bindPessimisticglobalexitrootnopush0(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Pessimisticglobalexitrootnopush0Caller{contract: contract}, nil
}

// NewPessimisticglobalexitrootnopush0Transactor creates a new write-only instance of Pessimisticglobalexitrootnopush0, bound to a specific deployed contract.
func NewPessimisticglobalexitrootnopush0Transactor(address common.Address, transactor bind.ContractTransactor) (*Pessimisticglobalexitrootnopush0Transactor, error) {
	contract, err := bindPessimisticglobalexitrootnopush0(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Pessimisticglobalexitrootnopush0Transactor{contract: contract}, nil
}

// NewPessimisticglobalexitrootnopush0Filterer creates a new log filterer instance of Pessimisticglobalexitrootnopush0, bound to a specific deployed contract.
func NewPessimisticglobalexitrootnopush0Filterer(address common.Address, filterer bind.ContractFilterer) (*Pessimisticglobalexitrootnopush0Filterer, error) {
	contract, err := bindPessimisticglobalexitrootnopush0(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Pessimisticglobalexitrootnopush0Filterer{contract: contract}, nil
}

// bindPessimisticglobalexitrootnopush0 binds a generic wrapper to an already deployed contract.
func bindPessimisticglobalexitrootnopush0(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Pessimisticglobalexitrootnopush0MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pessimisticglobalexitrootnopush0.Contract.Pessimisticglobalexitrootnopush0Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pessimisticglobalexitrootnopush0.Contract.Pessimisticglobalexitrootnopush0Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pessimisticglobalexitrootnopush0.Contract.Pessimisticglobalexitrootnopush0Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pessimisticglobalexitrootnopush0.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pessimisticglobalexitrootnopush0.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pessimisticglobalexitrootnopush0.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pessimisticglobalexitrootnopush0.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _Pessimisticglobalexitrootnopush0.Contract.DEFAULTADMINROLE(&_Pessimisticglobalexitrootnopush0.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Pessimisticglobalexitrootnopush0.Contract.DEFAULTADMINROLE(&_Pessimisticglobalexitrootnopush0.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Caller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pessimisticglobalexitrootnopush0.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Session) BridgeAddress() (common.Address, error) {
	return _Pessimisticglobalexitrootnopush0.Contract.BridgeAddress(&_Pessimisticglobalexitrootnopush0.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0CallerSession) BridgeAddress() (common.Address, error) {
	return _Pessimisticglobalexitrootnopush0.Contract.BridgeAddress(&_Pessimisticglobalexitrootnopush0.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Pessimisticglobalexitrootnopush0.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Pessimisticglobalexitrootnopush0.Contract.GetRoleAdmin(&_Pessimisticglobalexitrootnopush0.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Pessimisticglobalexitrootnopush0.Contract.GetRoleAdmin(&_Pessimisticglobalexitrootnopush0.CallOpts, role)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Caller) GlobalExitRootMap(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Pessimisticglobalexitrootnopush0.contract.Call(opts, &out, "globalExitRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Session) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Pessimisticglobalexitrootnopush0.Contract.GlobalExitRootMap(&_Pessimisticglobalexitrootnopush0.CallOpts, arg0)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0CallerSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Pessimisticglobalexitrootnopush0.Contract.GlobalExitRootMap(&_Pessimisticglobalexitrootnopush0.CallOpts, arg0)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Pessimisticglobalexitrootnopush0.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Pessimisticglobalexitrootnopush0.Contract.HasRole(&_Pessimisticglobalexitrootnopush0.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Pessimisticglobalexitrootnopush0.Contract.HasRole(&_Pessimisticglobalexitrootnopush0.CallOpts, role, account)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Caller) LastRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pessimisticglobalexitrootnopush0.contract.Call(opts, &out, "lastRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Session) LastRollupExitRoot() ([32]byte, error) {
	return _Pessimisticglobalexitrootnopush0.Contract.LastRollupExitRoot(&_Pessimisticglobalexitrootnopush0.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0CallerSession) LastRollupExitRoot() ([32]byte, error) {
	return _Pessimisticglobalexitrootnopush0.Contract.LastRollupExitRoot(&_Pessimisticglobalexitrootnopush0.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pessimisticglobalexitrootnopush0.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pessimisticglobalexitrootnopush0.Contract.GrantRole(&_Pessimisticglobalexitrootnopush0.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pessimisticglobalexitrootnopush0.Contract.GrantRole(&_Pessimisticglobalexitrootnopush0.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pessimisticglobalexitrootnopush0.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pessimisticglobalexitrootnopush0.Contract.RenounceRole(&_Pessimisticglobalexitrootnopush0.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pessimisticglobalexitrootnopush0.Contract.RenounceRole(&_Pessimisticglobalexitrootnopush0.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pessimisticglobalexitrootnopush0.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pessimisticglobalexitrootnopush0.Contract.RevokeRole(&_Pessimisticglobalexitrootnopush0.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pessimisticglobalexitrootnopush0.Contract.RevokeRole(&_Pessimisticglobalexitrootnopush0.TransactOpts, role, account)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Transactor) UpdateExitRoot(opts *bind.TransactOpts, newRoot [32]byte) (*types.Transaction, error) {
	return _Pessimisticglobalexitrootnopush0.contract.Transact(opts, "updateExitRoot", newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Session) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Pessimisticglobalexitrootnopush0.Contract.UpdateExitRoot(&_Pessimisticglobalexitrootnopush0.TransactOpts, newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0TransactorSession) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Pessimisticglobalexitrootnopush0.Contract.UpdateExitRoot(&_Pessimisticglobalexitrootnopush0.TransactOpts, newRoot)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x48da4bdd.
//
// Solidity: function updateGlobalExitRoot(bytes32 _newRoot) returns()
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Transactor) UpdateGlobalExitRoot(opts *bind.TransactOpts, _newRoot [32]byte) (*types.Transaction, error) {
	return _Pessimisticglobalexitrootnopush0.contract.Transact(opts, "updateGlobalExitRoot", _newRoot)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x48da4bdd.
//
// Solidity: function updateGlobalExitRoot(bytes32 _newRoot) returns()
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Session) UpdateGlobalExitRoot(_newRoot [32]byte) (*types.Transaction, error) {
	return _Pessimisticglobalexitrootnopush0.Contract.UpdateGlobalExitRoot(&_Pessimisticglobalexitrootnopush0.TransactOpts, _newRoot)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x48da4bdd.
//
// Solidity: function updateGlobalExitRoot(bytes32 _newRoot) returns()
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0TransactorSession) UpdateGlobalExitRoot(_newRoot [32]byte) (*types.Transaction, error) {
	return _Pessimisticglobalexitrootnopush0.Contract.UpdateGlobalExitRoot(&_Pessimisticglobalexitrootnopush0.TransactOpts, _newRoot)
}

// Pessimisticglobalexitrootnopush0InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Pessimisticglobalexitrootnopush0 contract.
type Pessimisticglobalexitrootnopush0InitializedIterator struct {
	Event *Pessimisticglobalexitrootnopush0Initialized // Event containing the contract specifics and raw log

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
func (it *Pessimisticglobalexitrootnopush0InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pessimisticglobalexitrootnopush0Initialized)
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
		it.Event = new(Pessimisticglobalexitrootnopush0Initialized)
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
func (it *Pessimisticglobalexitrootnopush0InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pessimisticglobalexitrootnopush0InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pessimisticglobalexitrootnopush0Initialized represents a Initialized event raised by the Pessimisticglobalexitrootnopush0 contract.
type Pessimisticglobalexitrootnopush0Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Filterer) FilterInitialized(opts *bind.FilterOpts) (*Pessimisticglobalexitrootnopush0InitializedIterator, error) {

	logs, sub, err := _Pessimisticglobalexitrootnopush0.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Pessimisticglobalexitrootnopush0InitializedIterator{contract: _Pessimisticglobalexitrootnopush0.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Pessimisticglobalexitrootnopush0Initialized) (event.Subscription, error) {

	logs, sub, err := _Pessimisticglobalexitrootnopush0.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pessimisticglobalexitrootnopush0Initialized)
				if err := _Pessimisticglobalexitrootnopush0.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Filterer) ParseInitialized(log types.Log) (*Pessimisticglobalexitrootnopush0Initialized, error) {
	event := new(Pessimisticglobalexitrootnopush0Initialized)
	if err := _Pessimisticglobalexitrootnopush0.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Pessimisticglobalexitrootnopush0RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Pessimisticglobalexitrootnopush0 contract.
type Pessimisticglobalexitrootnopush0RoleAdminChangedIterator struct {
	Event *Pessimisticglobalexitrootnopush0RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *Pessimisticglobalexitrootnopush0RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pessimisticglobalexitrootnopush0RoleAdminChanged)
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
		it.Event = new(Pessimisticglobalexitrootnopush0RoleAdminChanged)
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
func (it *Pessimisticglobalexitrootnopush0RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pessimisticglobalexitrootnopush0RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pessimisticglobalexitrootnopush0RoleAdminChanged represents a RoleAdminChanged event raised by the Pessimisticglobalexitrootnopush0 contract.
type Pessimisticglobalexitrootnopush0RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*Pessimisticglobalexitrootnopush0RoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Pessimisticglobalexitrootnopush0.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &Pessimisticglobalexitrootnopush0RoleAdminChangedIterator{contract: _Pessimisticglobalexitrootnopush0.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *Pessimisticglobalexitrootnopush0RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Pessimisticglobalexitrootnopush0.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pessimisticglobalexitrootnopush0RoleAdminChanged)
				if err := _Pessimisticglobalexitrootnopush0.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Filterer) ParseRoleAdminChanged(log types.Log) (*Pessimisticglobalexitrootnopush0RoleAdminChanged, error) {
	event := new(Pessimisticglobalexitrootnopush0RoleAdminChanged)
	if err := _Pessimisticglobalexitrootnopush0.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Pessimisticglobalexitrootnopush0RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Pessimisticglobalexitrootnopush0 contract.
type Pessimisticglobalexitrootnopush0RoleGrantedIterator struct {
	Event *Pessimisticglobalexitrootnopush0RoleGranted // Event containing the contract specifics and raw log

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
func (it *Pessimisticglobalexitrootnopush0RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pessimisticglobalexitrootnopush0RoleGranted)
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
		it.Event = new(Pessimisticglobalexitrootnopush0RoleGranted)
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
func (it *Pessimisticglobalexitrootnopush0RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pessimisticglobalexitrootnopush0RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pessimisticglobalexitrootnopush0RoleGranted represents a RoleGranted event raised by the Pessimisticglobalexitrootnopush0 contract.
type Pessimisticglobalexitrootnopush0RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Pessimisticglobalexitrootnopush0RoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pessimisticglobalexitrootnopush0.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Pessimisticglobalexitrootnopush0RoleGrantedIterator{contract: _Pessimisticglobalexitrootnopush0.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *Pessimisticglobalexitrootnopush0RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pessimisticglobalexitrootnopush0.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pessimisticglobalexitrootnopush0RoleGranted)
				if err := _Pessimisticglobalexitrootnopush0.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Filterer) ParseRoleGranted(log types.Log) (*Pessimisticglobalexitrootnopush0RoleGranted, error) {
	event := new(Pessimisticglobalexitrootnopush0RoleGranted)
	if err := _Pessimisticglobalexitrootnopush0.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Pessimisticglobalexitrootnopush0RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Pessimisticglobalexitrootnopush0 contract.
type Pessimisticglobalexitrootnopush0RoleRevokedIterator struct {
	Event *Pessimisticglobalexitrootnopush0RoleRevoked // Event containing the contract specifics and raw log

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
func (it *Pessimisticglobalexitrootnopush0RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pessimisticglobalexitrootnopush0RoleRevoked)
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
		it.Event = new(Pessimisticglobalexitrootnopush0RoleRevoked)
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
func (it *Pessimisticglobalexitrootnopush0RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pessimisticglobalexitrootnopush0RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pessimisticglobalexitrootnopush0RoleRevoked represents a RoleRevoked event raised by the Pessimisticglobalexitrootnopush0 contract.
type Pessimisticglobalexitrootnopush0RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Pessimisticglobalexitrootnopush0RoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pessimisticglobalexitrootnopush0.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Pessimisticglobalexitrootnopush0RoleRevokedIterator{contract: _Pessimisticglobalexitrootnopush0.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *Pessimisticglobalexitrootnopush0RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pessimisticglobalexitrootnopush0.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pessimisticglobalexitrootnopush0RoleRevoked)
				if err := _Pessimisticglobalexitrootnopush0.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pessimisticglobalexitrootnopush0 *Pessimisticglobalexitrootnopush0Filterer) ParseRoleRevoked(log types.Log) (*Pessimisticglobalexitrootnopush0RoleRevoked, error) {
	event := new(Pessimisticglobalexitrootnopush0RoleRevoked)
	if err := _Pessimisticglobalexitrootnopush0.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

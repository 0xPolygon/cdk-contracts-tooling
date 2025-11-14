// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package agglayerger

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

// AgglayergerMetaData contains all meta data concerning the Agglayerger contract.
var AgglayergerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollupManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"GlobalExitRootAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewDepositCountExceedsMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonZeroValueForUnusedFrontier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootRemover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootUpdater\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingGlobalExitRootRemover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingGlobalExitRootUpdater\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubtreeFrontierMismatch\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"leafCount\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"currentL1InfoRoot\",\"type\":\"bytes32\"}],\"name\":\"InitL1InfoRootMap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"}],\"name\":\"UpdateL1InfoTree\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"currentL1InfoRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"leafCount\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockhash\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"minTimestamp\",\"type\":\"uint64\"}],\"name\":\"UpdateL1InfoTreeV2\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GER_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastGlobalExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"lastBlockHash\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"getLeafValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"leafCount\",\"type\":\"uint32\"}],\"name\":\"l1InfoRootMap\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"l1InfoRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastMainnetExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561000f575f5ffd5b50604051610d1b380380610d1b83398101604081905261002e9161012b565b6001600160a01b0380831660a0528116608052610049610050565b505061015c565b602e54610100900460ff16156100bc5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b602e5460ff908116101561010e57602e805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b80516001600160a01b0381168114610126575f5ffd5b919050565b5f5f6040838503121561013c575f5ffd5b61014583610110565b915061015360208401610110565b90509250929050565b60805160a051610b9061018b5f395f81816101a0015261034301525f818161029d01526102f70152610b905ff3fe608060405234801561000f575f5ffd5b50600436106100e5575f3560e01c806349b7b802116100885780635d810501116100635780635d810501146102285780638129fc1c14610290578063a3c573eb14610298578063ef4eeb35146102bf575f5ffd5b806349b7b8021461019b57806354fd4d50146101e75780635ca1e16514610220575f5ffd5b80632dfdf0b5116100c35780632dfdf0b51461016c578063319cf7351461017557806333d6247d1461017e5780633ed691ef14610193575f5ffd5b806301fd9044146100e95780630f14261914610104578063257b36321461014d575b5f5ffd5b6100f15f5481565b6040519081526020015b60405180910390f35b6101406040518060400160405280600681526020017f76312e302e30000000000000000000000000000000000000000000000000000081525081565b6040516100fb9190610887565b6100f161015b3660046108da565b60026020525f908152604090205481565b6100f160235481565b6100f160015481565b61019161018c3660046108da565b6102de565b005b6100f1610504565b6101c27f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100fb565b60408051808201909152600681527f76312e302e3000000000000000000000000000000000000000000000000000006020820152610140565b6100f1610517565b6100f16102363660046108f1565b604080516020808201959095528082019390935260c09190911b7fffffffffffffffff0000000000000000000000000000000000000000000000001660608301528051604881840301815260689092019052805191012090565b610191610520565b6101c27f000000000000000000000000000000000000000000000000000000000000000081565b6100f16102cd366004610933565b602f6020525f908152604090205481565b5f8073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016330361032c57505060018190555f54816103ab565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001633036103795750505f81905560015481906103ab565b6040517fb49365dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6103b68284610714565b5f81815260026020526040812054919250036104fe57425f6103d960014361098a565b5f8481526002602090815260409182902092409283905581518082018790528083018490527fffffffffffffffff00000000000000000000000000000000000000000000000060c087901b166060820152825180820360480181526068909101909252815191012090915061044d9061072b565b5f610456610517565b60235463ffffffff165f908152602f602052604080822083905551919250879187917fda61aa7823fcd807e37b95aabcbe17f03a6f3efd514176444dae191d27fd66b391a360235463ffffffff167faf6c6cd7790e0180a4d22eb8ed846e55846f54ed10e5946db19972b5a0813a598284866040516104f293929190928352602083019190915267ffffffffffffffff16604082015260600190565b60405180910390a25050505b50505050565b5f6105126001545f54610714565b905090565b5f610512610808565b602e54610100900460ff16158080156105405750602e54600160ff909116105b8061055a5750303b15801561055a5750602e5460ff166001145b6105ea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840160405180910390fd5b602e80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561064857602e80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b5f610651610517565b6023805463ffffffff9081165f908152602f602090815260409182902085905592548151921682529181018390529192507f11f50c71891002839c2637ce302087160298255a87f1ea60d40e8db081383fad910160405180910390a150801561071157602e80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b5f8281526020829052604081205b90505b92915050565b80600161073a60206002610abe565b610744919061098a565b6023541061077e576040517fef5ccf6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60235f815461078d90610ac9565b918290555090505f5b60208110156107fa578082901c6001166001036107c95782600382602081106107c1576107c1610b00565b015550505050565b6107f0600382602081106107df576107df610b00565b0154845f9182526020526040902090565b9250600101610796565b50610803610b2d565b505050565b6023545f90819081805b602081101561087e578083901c600116600103610857576108506003826020811061083f5761083f610b00565b0154855f9182526020526040902090565b9350610867565b5f84815260208390526040902093505b5f8281526020839052604090209150600101610812565b50919392505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b5f602082840312156108ea575f5ffd5b5035919050565b5f5f5f60608486031215610903575f5ffd5b8335925060208401359150604084013567ffffffffffffffff81168114610928575f5ffd5b809150509250925092565b5f60208284031215610943575f5ffd5b813563ffffffff81168114610956575f5ffd5b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b818103818111156107255761072561095d565b6001815b60018411156109d8578085048111156109bc576109bc61095d565b60018416156109ca57908102905b60019390931c9280026109a1565b935093915050565b5f826109ee57506001610725565b816109fa57505f610725565b8160018114610a105760028114610a1a57610a36565b6001915050610725565b60ff841115610a2b57610a2b61095d565b50506001821b610725565b5060208310610133831016604e8410600b8410161715610a59575081810a610725565b610a847fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff848461099d565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610ab657610ab661095d565b029392505050565b5f61072283836109e0565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610af957610af961095d565b5060010190565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52600160045260245ffdfea26469706673582212207efa18d38f9c894a5afaf7589a2afc39c8ce406e44414e5485dc456a7064db8d64736f6c634300081c0033",
}

// AgglayergerABI is the input ABI used to generate the binding from.
// Deprecated: Use AgglayergerMetaData.ABI instead.
var AgglayergerABI = AgglayergerMetaData.ABI

// AgglayergerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AgglayergerMetaData.Bin instead.
var AgglayergerBin = AgglayergerMetaData.Bin

// DeployAgglayerger deploys a new Ethereum contract, binding an instance of Agglayerger to it.
func DeployAgglayerger(auth *bind.TransactOpts, backend bind.ContractBackend, _rollupManager common.Address, _bridgeAddress common.Address) (common.Address, *types.Transaction, *Agglayerger, error) {
	parsed, err := AgglayergerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AgglayergerBin), backend, _rollupManager, _bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Agglayerger{AgglayergerCaller: AgglayergerCaller{contract: contract}, AgglayergerTransactor: AgglayergerTransactor{contract: contract}, AgglayergerFilterer: AgglayergerFilterer{contract: contract}}, nil
}

// Agglayerger is an auto generated Go binding around an Ethereum contract.
type Agglayerger struct {
	AgglayergerCaller     // Read-only binding to the contract
	AgglayergerTransactor // Write-only binding to the contract
	AgglayergerFilterer   // Log filterer for contract events
}

// AgglayergerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgglayergerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgglayergerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgglayergerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgglayergerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgglayergerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgglayergerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgglayergerSession struct {
	Contract     *Agglayerger      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgglayergerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgglayergerCallerSession struct {
	Contract *AgglayergerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AgglayergerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgglayergerTransactorSession struct {
	Contract     *AgglayergerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AgglayergerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgglayergerRaw struct {
	Contract *Agglayerger // Generic contract binding to access the raw methods on
}

// AgglayergerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgglayergerCallerRaw struct {
	Contract *AgglayergerCaller // Generic read-only contract binding to access the raw methods on
}

// AgglayergerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgglayergerTransactorRaw struct {
	Contract *AgglayergerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgglayerger creates a new instance of Agglayerger, bound to a specific deployed contract.
func NewAgglayerger(address common.Address, backend bind.ContractBackend) (*Agglayerger, error) {
	contract, err := bindAgglayerger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Agglayerger{AgglayergerCaller: AgglayergerCaller{contract: contract}, AgglayergerTransactor: AgglayergerTransactor{contract: contract}, AgglayergerFilterer: AgglayergerFilterer{contract: contract}}, nil
}

// NewAgglayergerCaller creates a new read-only instance of Agglayerger, bound to a specific deployed contract.
func NewAgglayergerCaller(address common.Address, caller bind.ContractCaller) (*AgglayergerCaller, error) {
	contract, err := bindAgglayerger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgglayergerCaller{contract: contract}, nil
}

// NewAgglayergerTransactor creates a new write-only instance of Agglayerger, bound to a specific deployed contract.
func NewAgglayergerTransactor(address common.Address, transactor bind.ContractTransactor) (*AgglayergerTransactor, error) {
	contract, err := bindAgglayerger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgglayergerTransactor{contract: contract}, nil
}

// NewAgglayergerFilterer creates a new log filterer instance of Agglayerger, bound to a specific deployed contract.
func NewAgglayergerFilterer(address common.Address, filterer bind.ContractFilterer) (*AgglayergerFilterer, error) {
	contract, err := bindAgglayerger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgglayergerFilterer{contract: contract}, nil
}

// bindAgglayerger binds a generic wrapper to an already deployed contract.
func bindAgglayerger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AgglayergerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Agglayerger *AgglayergerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Agglayerger.Contract.AgglayergerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Agglayerger *AgglayergerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayerger.Contract.AgglayergerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Agglayerger *AgglayergerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Agglayerger.Contract.AgglayergerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Agglayerger *AgglayergerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Agglayerger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Agglayerger *AgglayergerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayerger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Agglayerger *AgglayergerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Agglayerger.Contract.contract.Transact(opts, method, params...)
}

// GERVERSION is a free data retrieval call binding the contract method 0x0f142619.
//
// Solidity: function GER_VERSION() view returns(string)
func (_Agglayerger *AgglayergerCaller) GERVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Agglayerger.contract.Call(opts, &out, "GER_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GERVERSION is a free data retrieval call binding the contract method 0x0f142619.
//
// Solidity: function GER_VERSION() view returns(string)
func (_Agglayerger *AgglayergerSession) GERVERSION() (string, error) {
	return _Agglayerger.Contract.GERVERSION(&_Agglayerger.CallOpts)
}

// GERVERSION is a free data retrieval call binding the contract method 0x0f142619.
//
// Solidity: function GER_VERSION() view returns(string)
func (_Agglayerger *AgglayergerCallerSession) GERVERSION() (string, error) {
	return _Agglayerger.Contract.GERVERSION(&_Agglayerger.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Agglayerger *AgglayergerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayerger.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Agglayerger *AgglayergerSession) BridgeAddress() (common.Address, error) {
	return _Agglayerger.Contract.BridgeAddress(&_Agglayerger.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Agglayerger *AgglayergerCallerSession) BridgeAddress() (common.Address, error) {
	return _Agglayerger.Contract.BridgeAddress(&_Agglayerger.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Agglayerger *AgglayergerCaller) DepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Agglayerger.contract.Call(opts, &out, "depositCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Agglayerger *AgglayergerSession) DepositCount() (*big.Int, error) {
	return _Agglayerger.Contract.DepositCount(&_Agglayerger.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Agglayerger *AgglayergerCallerSession) DepositCount() (*big.Int, error) {
	return _Agglayerger.Contract.DepositCount(&_Agglayerger.CallOpts)
}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Agglayerger *AgglayergerCaller) GetLastGlobalExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Agglayerger.contract.Call(opts, &out, "getLastGlobalExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Agglayerger *AgglayergerSession) GetLastGlobalExitRoot() ([32]byte, error) {
	return _Agglayerger.Contract.GetLastGlobalExitRoot(&_Agglayerger.CallOpts)
}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Agglayerger *AgglayergerCallerSession) GetLastGlobalExitRoot() ([32]byte, error) {
	return _Agglayerger.Contract.GetLastGlobalExitRoot(&_Agglayerger.CallOpts)
}

// GetLeafValue is a free data retrieval call binding the contract method 0x5d810501.
//
// Solidity: function getLeafValue(bytes32 newGlobalExitRoot, uint256 lastBlockHash, uint64 timestamp) pure returns(bytes32)
func (_Agglayerger *AgglayergerCaller) GetLeafValue(opts *bind.CallOpts, newGlobalExitRoot [32]byte, lastBlockHash *big.Int, timestamp uint64) ([32]byte, error) {
	var out []interface{}
	err := _Agglayerger.contract.Call(opts, &out, "getLeafValue", newGlobalExitRoot, lastBlockHash, timestamp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLeafValue is a free data retrieval call binding the contract method 0x5d810501.
//
// Solidity: function getLeafValue(bytes32 newGlobalExitRoot, uint256 lastBlockHash, uint64 timestamp) pure returns(bytes32)
func (_Agglayerger *AgglayergerSession) GetLeafValue(newGlobalExitRoot [32]byte, lastBlockHash *big.Int, timestamp uint64) ([32]byte, error) {
	return _Agglayerger.Contract.GetLeafValue(&_Agglayerger.CallOpts, newGlobalExitRoot, lastBlockHash, timestamp)
}

// GetLeafValue is a free data retrieval call binding the contract method 0x5d810501.
//
// Solidity: function getLeafValue(bytes32 newGlobalExitRoot, uint256 lastBlockHash, uint64 timestamp) pure returns(bytes32)
func (_Agglayerger *AgglayergerCallerSession) GetLeafValue(newGlobalExitRoot [32]byte, lastBlockHash *big.Int, timestamp uint64) ([32]byte, error) {
	return _Agglayerger.Contract.GetLeafValue(&_Agglayerger.CallOpts, newGlobalExitRoot, lastBlockHash, timestamp)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Agglayerger *AgglayergerCaller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Agglayerger.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Agglayerger *AgglayergerSession) GetRoot() ([32]byte, error) {
	return _Agglayerger.Contract.GetRoot(&_Agglayerger.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Agglayerger *AgglayergerCallerSession) GetRoot() ([32]byte, error) {
	return _Agglayerger.Contract.GetRoot(&_Agglayerger.CallOpts)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Agglayerger *AgglayergerCaller) GlobalExitRootMap(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Agglayerger.contract.Call(opts, &out, "globalExitRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Agglayerger *AgglayergerSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Agglayerger.Contract.GlobalExitRootMap(&_Agglayerger.CallOpts, arg0)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Agglayerger *AgglayergerCallerSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Agglayerger.Contract.GlobalExitRootMap(&_Agglayerger.CallOpts, arg0)
}

// L1InfoRootMap is a free data retrieval call binding the contract method 0xef4eeb35.
//
// Solidity: function l1InfoRootMap(uint32 leafCount) view returns(bytes32 l1InfoRoot)
func (_Agglayerger *AgglayergerCaller) L1InfoRootMap(opts *bind.CallOpts, leafCount uint32) ([32]byte, error) {
	var out []interface{}
	err := _Agglayerger.contract.Call(opts, &out, "l1InfoRootMap", leafCount)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L1InfoRootMap is a free data retrieval call binding the contract method 0xef4eeb35.
//
// Solidity: function l1InfoRootMap(uint32 leafCount) view returns(bytes32 l1InfoRoot)
func (_Agglayerger *AgglayergerSession) L1InfoRootMap(leafCount uint32) ([32]byte, error) {
	return _Agglayerger.Contract.L1InfoRootMap(&_Agglayerger.CallOpts, leafCount)
}

// L1InfoRootMap is a free data retrieval call binding the contract method 0xef4eeb35.
//
// Solidity: function l1InfoRootMap(uint32 leafCount) view returns(bytes32 l1InfoRoot)
func (_Agglayerger *AgglayergerCallerSession) L1InfoRootMap(leafCount uint32) ([32]byte, error) {
	return _Agglayerger.Contract.L1InfoRootMap(&_Agglayerger.CallOpts, leafCount)
}

// LastMainnetExitRoot is a free data retrieval call binding the contract method 0x319cf735.
//
// Solidity: function lastMainnetExitRoot() view returns(bytes32)
func (_Agglayerger *AgglayergerCaller) LastMainnetExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Agglayerger.contract.Call(opts, &out, "lastMainnetExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastMainnetExitRoot is a free data retrieval call binding the contract method 0x319cf735.
//
// Solidity: function lastMainnetExitRoot() view returns(bytes32)
func (_Agglayerger *AgglayergerSession) LastMainnetExitRoot() ([32]byte, error) {
	return _Agglayerger.Contract.LastMainnetExitRoot(&_Agglayerger.CallOpts)
}

// LastMainnetExitRoot is a free data retrieval call binding the contract method 0x319cf735.
//
// Solidity: function lastMainnetExitRoot() view returns(bytes32)
func (_Agglayerger *AgglayergerCallerSession) LastMainnetExitRoot() ([32]byte, error) {
	return _Agglayerger.Contract.LastMainnetExitRoot(&_Agglayerger.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Agglayerger *AgglayergerCaller) LastRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Agglayerger.contract.Call(opts, &out, "lastRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Agglayerger *AgglayergerSession) LastRollupExitRoot() ([32]byte, error) {
	return _Agglayerger.Contract.LastRollupExitRoot(&_Agglayerger.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Agglayerger *AgglayergerCallerSession) LastRollupExitRoot() ([32]byte, error) {
	return _Agglayerger.Contract.LastRollupExitRoot(&_Agglayerger.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Agglayerger *AgglayergerCaller) RollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayerger.contract.Call(opts, &out, "rollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Agglayerger *AgglayergerSession) RollupManager() (common.Address, error) {
	return _Agglayerger.Contract.RollupManager(&_Agglayerger.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Agglayerger *AgglayergerCallerSession) RollupManager() (common.Address, error) {
	return _Agglayerger.Contract.RollupManager(&_Agglayerger.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Agglayerger *AgglayergerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Agglayerger.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Agglayerger *AgglayergerSession) Version() (string, error) {
	return _Agglayerger.Contract.Version(&_Agglayerger.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Agglayerger *AgglayergerCallerSession) Version() (string, error) {
	return _Agglayerger.Contract.Version(&_Agglayerger.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Agglayerger *AgglayergerTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayerger.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Agglayerger *AgglayergerSession) Initialize() (*types.Transaction, error) {
	return _Agglayerger.Contract.Initialize(&_Agglayerger.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Agglayerger *AgglayergerTransactorSession) Initialize() (*types.Transaction, error) {
	return _Agglayerger.Contract.Initialize(&_Agglayerger.TransactOpts)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Agglayerger *AgglayergerTransactor) UpdateExitRoot(opts *bind.TransactOpts, newRoot [32]byte) (*types.Transaction, error) {
	return _Agglayerger.contract.Transact(opts, "updateExitRoot", newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Agglayerger *AgglayergerSession) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Agglayerger.Contract.UpdateExitRoot(&_Agglayerger.TransactOpts, newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Agglayerger *AgglayergerTransactorSession) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Agglayerger.Contract.UpdateExitRoot(&_Agglayerger.TransactOpts, newRoot)
}

// AgglayergerInitL1InfoRootMapIterator is returned from FilterInitL1InfoRootMap and is used to iterate over the raw logs and unpacked data for InitL1InfoRootMap events raised by the Agglayerger contract.
type AgglayergerInitL1InfoRootMapIterator struct {
	Event *AgglayergerInitL1InfoRootMap // Event containing the contract specifics and raw log

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
func (it *AgglayergerInitL1InfoRootMapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayergerInitL1InfoRootMap)
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
		it.Event = new(AgglayergerInitL1InfoRootMap)
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
func (it *AgglayergerInitL1InfoRootMapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayergerInitL1InfoRootMapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayergerInitL1InfoRootMap represents a InitL1InfoRootMap event raised by the Agglayerger contract.
type AgglayergerInitL1InfoRootMap struct {
	LeafCount         uint32
	CurrentL1InfoRoot [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterInitL1InfoRootMap is a free log retrieval operation binding the contract event 0x11f50c71891002839c2637ce302087160298255a87f1ea60d40e8db081383fad.
//
// Solidity: event InitL1InfoRootMap(uint32 leafCount, bytes32 currentL1InfoRoot)
func (_Agglayerger *AgglayergerFilterer) FilterInitL1InfoRootMap(opts *bind.FilterOpts) (*AgglayergerInitL1InfoRootMapIterator, error) {

	logs, sub, err := _Agglayerger.contract.FilterLogs(opts, "InitL1InfoRootMap")
	if err != nil {
		return nil, err
	}
	return &AgglayergerInitL1InfoRootMapIterator{contract: _Agglayerger.contract, event: "InitL1InfoRootMap", logs: logs, sub: sub}, nil
}

// WatchInitL1InfoRootMap is a free log subscription operation binding the contract event 0x11f50c71891002839c2637ce302087160298255a87f1ea60d40e8db081383fad.
//
// Solidity: event InitL1InfoRootMap(uint32 leafCount, bytes32 currentL1InfoRoot)
func (_Agglayerger *AgglayergerFilterer) WatchInitL1InfoRootMap(opts *bind.WatchOpts, sink chan<- *AgglayergerInitL1InfoRootMap) (event.Subscription, error) {

	logs, sub, err := _Agglayerger.contract.WatchLogs(opts, "InitL1InfoRootMap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayergerInitL1InfoRootMap)
				if err := _Agglayerger.contract.UnpackLog(event, "InitL1InfoRootMap", log); err != nil {
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

// ParseInitL1InfoRootMap is a log parse operation binding the contract event 0x11f50c71891002839c2637ce302087160298255a87f1ea60d40e8db081383fad.
//
// Solidity: event InitL1InfoRootMap(uint32 leafCount, bytes32 currentL1InfoRoot)
func (_Agglayerger *AgglayergerFilterer) ParseInitL1InfoRootMap(log types.Log) (*AgglayergerInitL1InfoRootMap, error) {
	event := new(AgglayergerInitL1InfoRootMap)
	if err := _Agglayerger.contract.UnpackLog(event, "InitL1InfoRootMap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayergerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Agglayerger contract.
type AgglayergerInitializedIterator struct {
	Event *AgglayergerInitialized // Event containing the contract specifics and raw log

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
func (it *AgglayergerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayergerInitialized)
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
		it.Event = new(AgglayergerInitialized)
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
func (it *AgglayergerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayergerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayergerInitialized represents a Initialized event raised by the Agglayerger contract.
type AgglayergerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Agglayerger *AgglayergerFilterer) FilterInitialized(opts *bind.FilterOpts) (*AgglayergerInitializedIterator, error) {

	logs, sub, err := _Agglayerger.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AgglayergerInitializedIterator{contract: _Agglayerger.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Agglayerger *AgglayergerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AgglayergerInitialized) (event.Subscription, error) {

	logs, sub, err := _Agglayerger.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayergerInitialized)
				if err := _Agglayerger.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Agglayerger *AgglayergerFilterer) ParseInitialized(log types.Log) (*AgglayergerInitialized, error) {
	event := new(AgglayergerInitialized)
	if err := _Agglayerger.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayergerUpdateL1InfoTreeIterator is returned from FilterUpdateL1InfoTree and is used to iterate over the raw logs and unpacked data for UpdateL1InfoTree events raised by the Agglayerger contract.
type AgglayergerUpdateL1InfoTreeIterator struct {
	Event *AgglayergerUpdateL1InfoTree // Event containing the contract specifics and raw log

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
func (it *AgglayergerUpdateL1InfoTreeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayergerUpdateL1InfoTree)
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
		it.Event = new(AgglayergerUpdateL1InfoTree)
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
func (it *AgglayergerUpdateL1InfoTreeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayergerUpdateL1InfoTreeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayergerUpdateL1InfoTree represents a UpdateL1InfoTree event raised by the Agglayerger contract.
type AgglayergerUpdateL1InfoTree struct {
	MainnetExitRoot [32]byte
	RollupExitRoot  [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateL1InfoTree is a free log retrieval operation binding the contract event 0xda61aa7823fcd807e37b95aabcbe17f03a6f3efd514176444dae191d27fd66b3.
//
// Solidity: event UpdateL1InfoTree(bytes32 indexed mainnetExitRoot, bytes32 indexed rollupExitRoot)
func (_Agglayerger *AgglayergerFilterer) FilterUpdateL1InfoTree(opts *bind.FilterOpts, mainnetExitRoot [][32]byte, rollupExitRoot [][32]byte) (*AgglayergerUpdateL1InfoTreeIterator, error) {

	var mainnetExitRootRule []interface{}
	for _, mainnetExitRootItem := range mainnetExitRoot {
		mainnetExitRootRule = append(mainnetExitRootRule, mainnetExitRootItem)
	}
	var rollupExitRootRule []interface{}
	for _, rollupExitRootItem := range rollupExitRoot {
		rollupExitRootRule = append(rollupExitRootRule, rollupExitRootItem)
	}

	logs, sub, err := _Agglayerger.contract.FilterLogs(opts, "UpdateL1InfoTree", mainnetExitRootRule, rollupExitRootRule)
	if err != nil {
		return nil, err
	}
	return &AgglayergerUpdateL1InfoTreeIterator{contract: _Agglayerger.contract, event: "UpdateL1InfoTree", logs: logs, sub: sub}, nil
}

// WatchUpdateL1InfoTree is a free log subscription operation binding the contract event 0xda61aa7823fcd807e37b95aabcbe17f03a6f3efd514176444dae191d27fd66b3.
//
// Solidity: event UpdateL1InfoTree(bytes32 indexed mainnetExitRoot, bytes32 indexed rollupExitRoot)
func (_Agglayerger *AgglayergerFilterer) WatchUpdateL1InfoTree(opts *bind.WatchOpts, sink chan<- *AgglayergerUpdateL1InfoTree, mainnetExitRoot [][32]byte, rollupExitRoot [][32]byte) (event.Subscription, error) {

	var mainnetExitRootRule []interface{}
	for _, mainnetExitRootItem := range mainnetExitRoot {
		mainnetExitRootRule = append(mainnetExitRootRule, mainnetExitRootItem)
	}
	var rollupExitRootRule []interface{}
	for _, rollupExitRootItem := range rollupExitRoot {
		rollupExitRootRule = append(rollupExitRootRule, rollupExitRootItem)
	}

	logs, sub, err := _Agglayerger.contract.WatchLogs(opts, "UpdateL1InfoTree", mainnetExitRootRule, rollupExitRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayergerUpdateL1InfoTree)
				if err := _Agglayerger.contract.UnpackLog(event, "UpdateL1InfoTree", log); err != nil {
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

// ParseUpdateL1InfoTree is a log parse operation binding the contract event 0xda61aa7823fcd807e37b95aabcbe17f03a6f3efd514176444dae191d27fd66b3.
//
// Solidity: event UpdateL1InfoTree(bytes32 indexed mainnetExitRoot, bytes32 indexed rollupExitRoot)
func (_Agglayerger *AgglayergerFilterer) ParseUpdateL1InfoTree(log types.Log) (*AgglayergerUpdateL1InfoTree, error) {
	event := new(AgglayergerUpdateL1InfoTree)
	if err := _Agglayerger.contract.UnpackLog(event, "UpdateL1InfoTree", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayergerUpdateL1InfoTreeV2Iterator is returned from FilterUpdateL1InfoTreeV2 and is used to iterate over the raw logs and unpacked data for UpdateL1InfoTreeV2 events raised by the Agglayerger contract.
type AgglayergerUpdateL1InfoTreeV2Iterator struct {
	Event *AgglayergerUpdateL1InfoTreeV2 // Event containing the contract specifics and raw log

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
func (it *AgglayergerUpdateL1InfoTreeV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayergerUpdateL1InfoTreeV2)
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
		it.Event = new(AgglayergerUpdateL1InfoTreeV2)
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
func (it *AgglayergerUpdateL1InfoTreeV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayergerUpdateL1InfoTreeV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayergerUpdateL1InfoTreeV2 represents a UpdateL1InfoTreeV2 event raised by the Agglayerger contract.
type AgglayergerUpdateL1InfoTreeV2 struct {
	CurrentL1InfoRoot [32]byte
	LeafCount         uint32
	Blockhash         *big.Int
	MinTimestamp      uint64
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdateL1InfoTreeV2 is a free log retrieval operation binding the contract event 0xaf6c6cd7790e0180a4d22eb8ed846e55846f54ed10e5946db19972b5a0813a59.
//
// Solidity: event UpdateL1InfoTreeV2(bytes32 currentL1InfoRoot, uint32 indexed leafCount, uint256 blockhash, uint64 minTimestamp)
func (_Agglayerger *AgglayergerFilterer) FilterUpdateL1InfoTreeV2(opts *bind.FilterOpts, leafCount []uint32) (*AgglayergerUpdateL1InfoTreeV2Iterator, error) {

	var leafCountRule []interface{}
	for _, leafCountItem := range leafCount {
		leafCountRule = append(leafCountRule, leafCountItem)
	}

	logs, sub, err := _Agglayerger.contract.FilterLogs(opts, "UpdateL1InfoTreeV2", leafCountRule)
	if err != nil {
		return nil, err
	}
	return &AgglayergerUpdateL1InfoTreeV2Iterator{contract: _Agglayerger.contract, event: "UpdateL1InfoTreeV2", logs: logs, sub: sub}, nil
}

// WatchUpdateL1InfoTreeV2 is a free log subscription operation binding the contract event 0xaf6c6cd7790e0180a4d22eb8ed846e55846f54ed10e5946db19972b5a0813a59.
//
// Solidity: event UpdateL1InfoTreeV2(bytes32 currentL1InfoRoot, uint32 indexed leafCount, uint256 blockhash, uint64 minTimestamp)
func (_Agglayerger *AgglayergerFilterer) WatchUpdateL1InfoTreeV2(opts *bind.WatchOpts, sink chan<- *AgglayergerUpdateL1InfoTreeV2, leafCount []uint32) (event.Subscription, error) {

	var leafCountRule []interface{}
	for _, leafCountItem := range leafCount {
		leafCountRule = append(leafCountRule, leafCountItem)
	}

	logs, sub, err := _Agglayerger.contract.WatchLogs(opts, "UpdateL1InfoTreeV2", leafCountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayergerUpdateL1InfoTreeV2)
				if err := _Agglayerger.contract.UnpackLog(event, "UpdateL1InfoTreeV2", log); err != nil {
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

// ParseUpdateL1InfoTreeV2 is a log parse operation binding the contract event 0xaf6c6cd7790e0180a4d22eb8ed846e55846f54ed10e5946db19972b5a0813a59.
//
// Solidity: event UpdateL1InfoTreeV2(bytes32 currentL1InfoRoot, uint32 indexed leafCount, uint256 blockhash, uint64 minTimestamp)
func (_Agglayerger *AgglayergerFilterer) ParseUpdateL1InfoTreeV2(log types.Log) (*AgglayergerUpdateL1InfoTreeV2, error) {
	event := new(AgglayergerUpdateL1InfoTreeV2)
	if err := _Agglayerger.contract.UnpackLog(event, "UpdateL1InfoTreeV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

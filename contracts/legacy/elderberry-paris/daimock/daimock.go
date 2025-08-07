// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package daimock

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

// DaimockMetaData contains all meta data concerning the Daimock contract.
var DaimockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"move\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"pull\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"push\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161163b38038061163b8339818101604052602081101561003357600080fd5b5051336000908152602081905260409081902060019055518060526115e98239604080519182900360520182208282018252600e83527f44616920537461626c65636f696e00000000000000000000000000000000000060209384015281518083018352600181527f3100000000000000000000000000000000000000000000000000000000000000908401528151808401919091527f0b1461ddc0c1d5ded79a1db0f74dae949050a7c0b28728c724b24958c27a328b818301527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6606082015260808101949094523060a0808601919091528151808603909101815260c09094019052825192019190912060055550611497806101526000396000f3fe608060405234801561001057600080fd5b50600436106101825760003560e01c80637ecebe00116100d8578063a9059cbb1161008c578063bf353dbb11610066578063bf353dbb14610548578063dd62ed3e1461057b578063f2d5d56b146105b657610182565b8063a9059cbb14610493578063b753a98c146104cc578063bb35783b1461050557610182565b806395d89b41116100bd57806395d89b411461041f5780639c52a7f1146104275780639dc29fac1461045a57610182565b80637ecebe00146103855780638fcbaf0c146103b857610182565b8063313ce5671161013a57806354fd4d501161011457806354fd4d501461031757806365fae35e1461031f57806370a082311461035257610182565b8063313ce567146102b65780633644e515146102d457806340c10f19146102dc57610182565b806318160ddd1161016b57806318160ddd1461025157806323b872dd1461026b57806330adf81f146102ae57610182565b806306fdde0314610187578063095ea7b314610204575b600080fd5b61018f6105ef565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101c95781810151838201526020016101b1565b50505050905090810190601f1680156101f65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61023d6004803603604081101561021a57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135610628565b604080519115158252519081900360200190f35b61025961069c565b60408051918252519081900360200190f35b61023d6004803603606081101561028157600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135811691602081013590911690604001356106a2565b610259610997565b6102be6109bb565b6040805160ff9092168252519081900360200190f35b6102596109c0565b610315600480360360408110156102f257600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356109c6565b005b61018f610afc565b6103156004803603602081101561033557600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610b35565b6102596004803603602081101561036857600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610bdd565b6102596004803603602081101561039b57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610bef565b61031560048036036101008110156103cf57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135811691602081013590911690604081013590606081013590608081013515159060ff60a0820135169060c08101359060e00135610c01565b61018f611022565b6103156004803603602081101561043d57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661105b565b6103156004803603604081101561047057600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135611100565b61023d600480360360408110156104a957600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356113d3565b610315600480360360408110156104e257600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356113e7565b6103156004803603606081101561051b57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135811691602081013590911690604001356113f7565b6102596004803603602081101561055e57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16611408565b6102596004803603604081101561059157600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135811691602001351661141a565b610315600480360360408110156105cc57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135611437565b6040518060400160405280600e81526020017f44616920537461626c65636f696e00000000000000000000000000000000000081525081565b33600081815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a35060015b92915050565b60015481565b73ffffffffffffffffffffffffffffffffffffffff831660009081526002602052604081205482111561073657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4461692f696e73756666696369656e742d62616c616e63650000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff841633148015906107ac575073ffffffffffffffffffffffffffffffffffffffff841660009081526003602090815260408083203384529091529020547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff14155b156108bd5773ffffffffffffffffffffffffffffffffffffffff8416600090815260036020908152604080832033845290915290205482111561085057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f4461692f696e73756666696369656e742d616c6c6f77616e6365000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff8416600090815260036020908152604080832033845290915290205461088b9083611442565b73ffffffffffffffffffffffffffffffffffffffff851660009081526003602090815260408083203384529091529020555b73ffffffffffffffffffffffffffffffffffffffff84166000908152600260205260409020546108ed9083611442565b73ffffffffffffffffffffffffffffffffffffffff80861660009081526002602052604080822093909355908516815220546109299083611452565b73ffffffffffffffffffffffffffffffffffffffff80851660008181526002602090815260409182902094909455805186815290519193928816927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a35060019392505050565b7fea2aa0a1be11a07ed86d755c93467f4f82362b452371d1ba94d1715123511acb81565b601281565b60055481565b33600090815260208190526040902054600114610a4457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f4461692f6e6f742d617574686f72697a65640000000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff8216600090815260026020526040902054610a749082611452565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260026020526040902055600154610aa79082611452565b60015560408051828152905173ffffffffffffffffffffffffffffffffffffffff8416916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b6040518060400160405280600181526020017f310000000000000000000000000000000000000000000000000000000000000081525081565b33600090815260208190526040902054600114610bb357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f4461692f6e6f742d617574686f72697a65640000000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff16600090815260208190526040902060019055565b60026020526000908152604090205481565b60046020526000908152604090205481565b600554604080517fea2aa0a1be11a07ed86d755c93467f4f82362b452371d1ba94d1715123511acb60208083019190915273ffffffffffffffffffffffffffffffffffffffff808d16838501819052908c166060840152608083018b905260a083018a905288151560c0808501919091528451808503909101815260e0840185528051908301207f190100000000000000000000000000000000000000000000000000000000000061010085015261010284019590955261012280840195909552835180840390950185526101429092019092528251929091019190912090610d4b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f4461692f696e76616c69642d616464726573732d300000000000000000000000604482015290519081900360640190fd5b6040805160008152602080820180845284905260ff8716828401526060820186905260808201859052915160019260a08084019391927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081019281900390910190855afa158015610dc0573d6000803e3d6000fd5b5050506020604051035173ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff1614610e6357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f4461692f696e76616c69642d7065726d69740000000000000000000000000000604482015290519081900360640190fd5b851580610e705750854211155b610edb57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f4461692f7065726d69742d657870697265640000000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff891660009081526004602052604090208054600181019091558714610f7657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f4461692f696e76616c69642d6e6f6e6365000000000000000000000000000000604482015290519081900360640190fd5b600085610f84576000610fa6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b73ffffffffffffffffffffffffffffffffffffffff808c166000818152600360209081526040808320948f168084529482529182902085905581518581529151949550929391927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92592918290030190a350505050505050505050565b6040518060400160405280600381526020017f444149000000000000000000000000000000000000000000000000000000000081525081565b336000908152602081905260409020546001146110d957604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f4461692f6e6f742d617574686f72697a65640000000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff16600090815260208190526040812055565b73ffffffffffffffffffffffffffffffffffffffff821660009081526002602052604090205481111561119457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4461692f696e73756666696369656e742d62616c616e63650000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff8216331480159061120a575073ffffffffffffffffffffffffffffffffffffffff821660009081526003602090815260408083203384529091529020547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff14155b1561131b5773ffffffffffffffffffffffffffffffffffffffff821660009081526003602090815260408083203384529091529020548111156112ae57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f4461692f696e73756666696369656e742d616c6c6f77616e6365000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff821660009081526003602090815260408083203384529091529020546112e99082611442565b73ffffffffffffffffffffffffffffffffffffffff831660009081526003602090815260408083203384529091529020555b73ffffffffffffffffffffffffffffffffffffffff821660009081526002602052604090205461134b9082611442565b73ffffffffffffffffffffffffffffffffffffffff831660009081526002602052604090205560015461137e9082611442565b60015560408051828152905160009173ffffffffffffffffffffffffffffffffffffffff8516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b60006113e03384846106a2565b9392505050565b6113f23383836106a2565b505050565b6114028383836106a2565b50505050565b60006020819052908152604090205481565b600360209081526000928352604080842090915290825290205481565b6113f28233836106a2565b8082038281111561069657600080fd5b8082018281101561069657600080fdfea265627a7a7231582006574b5f65913aeb884ff4ee56f988723c87221632066861b27206ca66b50d4464736f6c634300050c0032454950373132446f6d61696e28737472696e67206e616d652c737472696e672076657273696f6e2c75696e7432353620636861696e49642c6164647265737320766572696679696e67436f6e747261637429",
}

// DaimockABI is the input ABI used to generate the binding from.
// Deprecated: Use DaimockMetaData.ABI instead.
var DaimockABI = DaimockMetaData.ABI

// DaimockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DaimockMetaData.Bin instead.
var DaimockBin = DaimockMetaData.Bin

// DeployDaimock deploys a new Ethereum contract, binding an instance of Daimock to it.
func DeployDaimock(auth *bind.TransactOpts, backend bind.ContractBackend, chainId_ *big.Int) (common.Address, *types.Transaction, *Daimock, error) {
	parsed, err := DaimockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DaimockBin), backend, chainId_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Daimock{DaimockCaller: DaimockCaller{contract: contract}, DaimockTransactor: DaimockTransactor{contract: contract}, DaimockFilterer: DaimockFilterer{contract: contract}}, nil
}

// Daimock is an auto generated Go binding around an Ethereum contract.
type Daimock struct {
	DaimockCaller     // Read-only binding to the contract
	DaimockTransactor // Write-only binding to the contract
	DaimockFilterer   // Log filterer for contract events
}

// DaimockCaller is an auto generated read-only Go binding around an Ethereum contract.
type DaimockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaimockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DaimockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaimockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DaimockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaimockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DaimockSession struct {
	Contract     *Daimock          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DaimockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DaimockCallerSession struct {
	Contract *DaimockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DaimockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DaimockTransactorSession struct {
	Contract     *DaimockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DaimockRaw is an auto generated low-level Go binding around an Ethereum contract.
type DaimockRaw struct {
	Contract *Daimock // Generic contract binding to access the raw methods on
}

// DaimockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DaimockCallerRaw struct {
	Contract *DaimockCaller // Generic read-only contract binding to access the raw methods on
}

// DaimockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DaimockTransactorRaw struct {
	Contract *DaimockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDaimock creates a new instance of Daimock, bound to a specific deployed contract.
func NewDaimock(address common.Address, backend bind.ContractBackend) (*Daimock, error) {
	contract, err := bindDaimock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Daimock{DaimockCaller: DaimockCaller{contract: contract}, DaimockTransactor: DaimockTransactor{contract: contract}, DaimockFilterer: DaimockFilterer{contract: contract}}, nil
}

// NewDaimockCaller creates a new read-only instance of Daimock, bound to a specific deployed contract.
func NewDaimockCaller(address common.Address, caller bind.ContractCaller) (*DaimockCaller, error) {
	contract, err := bindDaimock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DaimockCaller{contract: contract}, nil
}

// NewDaimockTransactor creates a new write-only instance of Daimock, bound to a specific deployed contract.
func NewDaimockTransactor(address common.Address, transactor bind.ContractTransactor) (*DaimockTransactor, error) {
	contract, err := bindDaimock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DaimockTransactor{contract: contract}, nil
}

// NewDaimockFilterer creates a new log filterer instance of Daimock, bound to a specific deployed contract.
func NewDaimockFilterer(address common.Address, filterer bind.ContractFilterer) (*DaimockFilterer, error) {
	contract, err := bindDaimock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DaimockFilterer{contract: contract}, nil
}

// bindDaimock binds a generic wrapper to an already deployed contract.
func bindDaimock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DaimockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Daimock *DaimockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Daimock.Contract.DaimockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Daimock *DaimockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Daimock.Contract.DaimockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Daimock *DaimockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Daimock.Contract.DaimockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Daimock *DaimockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Daimock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Daimock *DaimockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Daimock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Daimock *DaimockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Daimock.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Daimock *DaimockCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Daimock.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Daimock *DaimockSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Daimock.Contract.DOMAINSEPARATOR(&_Daimock.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Daimock *DaimockCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Daimock.Contract.DOMAINSEPARATOR(&_Daimock.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Daimock *DaimockCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Daimock.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Daimock *DaimockSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Daimock.Contract.PERMITTYPEHASH(&_Daimock.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Daimock *DaimockCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Daimock.Contract.PERMITTYPEHASH(&_Daimock.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Daimock *DaimockCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Daimock.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Daimock *DaimockSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Daimock.Contract.Allowance(&_Daimock.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Daimock *DaimockCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Daimock.Contract.Allowance(&_Daimock.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Daimock *DaimockCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Daimock.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Daimock *DaimockSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Daimock.Contract.BalanceOf(&_Daimock.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Daimock *DaimockCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Daimock.Contract.BalanceOf(&_Daimock.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Daimock *DaimockCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Daimock.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Daimock *DaimockSession) Decimals() (uint8, error) {
	return _Daimock.Contract.Decimals(&_Daimock.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Daimock *DaimockCallerSession) Decimals() (uint8, error) {
	return _Daimock.Contract.Decimals(&_Daimock.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Daimock *DaimockCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Daimock.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Daimock *DaimockSession) Name() (string, error) {
	return _Daimock.Contract.Name(&_Daimock.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Daimock *DaimockCallerSession) Name() (string, error) {
	return _Daimock.Contract.Name(&_Daimock.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Daimock *DaimockCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Daimock.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Daimock *DaimockSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Daimock.Contract.Nonces(&_Daimock.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Daimock *DaimockCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Daimock.Contract.Nonces(&_Daimock.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Daimock *DaimockCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Daimock.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Daimock *DaimockSession) Symbol() (string, error) {
	return _Daimock.Contract.Symbol(&_Daimock.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Daimock *DaimockCallerSession) Symbol() (string, error) {
	return _Daimock.Contract.Symbol(&_Daimock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Daimock *DaimockCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Daimock.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Daimock *DaimockSession) TotalSupply() (*big.Int, error) {
	return _Daimock.Contract.TotalSupply(&_Daimock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Daimock *DaimockCallerSession) TotalSupply() (*big.Int, error) {
	return _Daimock.Contract.TotalSupply(&_Daimock.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Daimock *DaimockCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Daimock.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Daimock *DaimockSession) Version() (string, error) {
	return _Daimock.Contract.Version(&_Daimock.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Daimock *DaimockCallerSession) Version() (string, error) {
	return _Daimock.Contract.Version(&_Daimock.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Daimock *DaimockCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Daimock.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Daimock *DaimockSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Daimock.Contract.Wards(&_Daimock.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Daimock *DaimockCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Daimock.Contract.Wards(&_Daimock.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address usr, uint256 wad) returns(bool)
func (_Daimock *DaimockTransactor) Approve(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Daimock.contract.Transact(opts, "approve", usr, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address usr, uint256 wad) returns(bool)
func (_Daimock *DaimockSession) Approve(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Daimock.Contract.Approve(&_Daimock.TransactOpts, usr, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address usr, uint256 wad) returns(bool)
func (_Daimock *DaimockTransactorSession) Approve(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Daimock.Contract.Approve(&_Daimock.TransactOpts, usr, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address usr, uint256 wad) returns()
func (_Daimock *DaimockTransactor) Burn(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Daimock.contract.Transact(opts, "burn", usr, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address usr, uint256 wad) returns()
func (_Daimock *DaimockSession) Burn(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Daimock.Contract.Burn(&_Daimock.TransactOpts, usr, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address usr, uint256 wad) returns()
func (_Daimock *DaimockTransactorSession) Burn(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Daimock.Contract.Burn(&_Daimock.TransactOpts, usr, wad)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_Daimock *DaimockTransactor) Deny(opts *bind.TransactOpts, guy common.Address) (*types.Transaction, error) {
	return _Daimock.contract.Transact(opts, "deny", guy)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_Daimock *DaimockSession) Deny(guy common.Address) (*types.Transaction, error) {
	return _Daimock.Contract.Deny(&_Daimock.TransactOpts, guy)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_Daimock *DaimockTransactorSession) Deny(guy common.Address) (*types.Transaction, error) {
	return _Daimock.Contract.Deny(&_Daimock.TransactOpts, guy)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address usr, uint256 wad) returns()
func (_Daimock *DaimockTransactor) Mint(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Daimock.contract.Transact(opts, "mint", usr, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address usr, uint256 wad) returns()
func (_Daimock *DaimockSession) Mint(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Daimock.Contract.Mint(&_Daimock.TransactOpts, usr, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address usr, uint256 wad) returns()
func (_Daimock *DaimockTransactorSession) Mint(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Daimock.Contract.Mint(&_Daimock.TransactOpts, usr, wad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 wad) returns()
func (_Daimock *DaimockTransactor) Move(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Daimock.contract.Transact(opts, "move", src, dst, wad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 wad) returns()
func (_Daimock *DaimockSession) Move(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Daimock.Contract.Move(&_Daimock.TransactOpts, src, dst, wad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 wad) returns()
func (_Daimock *DaimockTransactorSession) Move(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Daimock.Contract.Move(&_Daimock.TransactOpts, src, dst, wad)
}

// Permit is a paid mutator transaction binding the contract method 0x8fcbaf0c.
//
// Solidity: function permit(address holder, address spender, uint256 nonce, uint256 expiry, bool allowed, uint8 v, bytes32 r, bytes32 s) returns()
func (_Daimock *DaimockTransactor) Permit(opts *bind.TransactOpts, holder common.Address, spender common.Address, nonce *big.Int, expiry *big.Int, allowed bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Daimock.contract.Transact(opts, "permit", holder, spender, nonce, expiry, allowed, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0x8fcbaf0c.
//
// Solidity: function permit(address holder, address spender, uint256 nonce, uint256 expiry, bool allowed, uint8 v, bytes32 r, bytes32 s) returns()
func (_Daimock *DaimockSession) Permit(holder common.Address, spender common.Address, nonce *big.Int, expiry *big.Int, allowed bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Daimock.Contract.Permit(&_Daimock.TransactOpts, holder, spender, nonce, expiry, allowed, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0x8fcbaf0c.
//
// Solidity: function permit(address holder, address spender, uint256 nonce, uint256 expiry, bool allowed, uint8 v, bytes32 r, bytes32 s) returns()
func (_Daimock *DaimockTransactorSession) Permit(holder common.Address, spender common.Address, nonce *big.Int, expiry *big.Int, allowed bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Daimock.Contract.Permit(&_Daimock.TransactOpts, holder, spender, nonce, expiry, allowed, v, r, s)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address usr, uint256 wad) returns()
func (_Daimock *DaimockTransactor) Pull(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Daimock.contract.Transact(opts, "pull", usr, wad)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address usr, uint256 wad) returns()
func (_Daimock *DaimockSession) Pull(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Daimock.Contract.Pull(&_Daimock.TransactOpts, usr, wad)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address usr, uint256 wad) returns()
func (_Daimock *DaimockTransactorSession) Pull(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Daimock.Contract.Pull(&_Daimock.TransactOpts, usr, wad)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(address usr, uint256 wad) returns()
func (_Daimock *DaimockTransactor) Push(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Daimock.contract.Transact(opts, "push", usr, wad)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(address usr, uint256 wad) returns()
func (_Daimock *DaimockSession) Push(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Daimock.Contract.Push(&_Daimock.TransactOpts, usr, wad)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(address usr, uint256 wad) returns()
func (_Daimock *DaimockTransactorSession) Push(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Daimock.Contract.Push(&_Daimock.TransactOpts, usr, wad)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_Daimock *DaimockTransactor) Rely(opts *bind.TransactOpts, guy common.Address) (*types.Transaction, error) {
	return _Daimock.contract.Transact(opts, "rely", guy)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_Daimock *DaimockSession) Rely(guy common.Address) (*types.Transaction, error) {
	return _Daimock.Contract.Rely(&_Daimock.TransactOpts, guy)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_Daimock *DaimockTransactorSession) Rely(guy common.Address) (*types.Transaction, error) {
	return _Daimock.Contract.Rely(&_Daimock.TransactOpts, guy)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Daimock *DaimockTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Daimock.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Daimock *DaimockSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Daimock.Contract.Transfer(&_Daimock.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Daimock *DaimockTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Daimock.Contract.Transfer(&_Daimock.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Daimock *DaimockTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Daimock.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Daimock *DaimockSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Daimock.Contract.TransferFrom(&_Daimock.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Daimock *DaimockTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Daimock.Contract.TransferFrom(&_Daimock.TransactOpts, src, dst, wad)
}

// DaimockApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Daimock contract.
type DaimockApprovalIterator struct {
	Event *DaimockApproval // Event containing the contract specifics and raw log

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
func (it *DaimockApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaimockApproval)
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
		it.Event = new(DaimockApproval)
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
func (it *DaimockApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaimockApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaimockApproval represents a Approval event raised by the Daimock contract.
type DaimockApproval struct {
	Src common.Address
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_Daimock *DaimockFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, guy []common.Address) (*DaimockApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Daimock.contract.FilterLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return &DaimockApprovalIterator{contract: _Daimock.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_Daimock *DaimockFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DaimockApproval, src []common.Address, guy []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Daimock.contract.WatchLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaimockApproval)
				if err := _Daimock.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_Daimock *DaimockFilterer) ParseApproval(log types.Log) (*DaimockApproval, error) {
	event := new(DaimockApproval)
	if err := _Daimock.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaimockTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Daimock contract.
type DaimockTransferIterator struct {
	Event *DaimockTransfer // Event containing the contract specifics and raw log

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
func (it *DaimockTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaimockTransfer)
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
		it.Event = new(DaimockTransfer)
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
func (it *DaimockTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaimockTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaimockTransfer represents a Transfer event raised by the Daimock contract.
type DaimockTransfer struct {
	Src common.Address
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_Daimock *DaimockFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*DaimockTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Daimock.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &DaimockTransferIterator{contract: _Daimock.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_Daimock *DaimockFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DaimockTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Daimock.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaimockTransfer)
				if err := _Daimock.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_Daimock *DaimockFilterer) ParseTransfer(log types.Log) (*DaimockTransfer, error) {
	event := new(DaimockTransfer)
	if err := _Daimock.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

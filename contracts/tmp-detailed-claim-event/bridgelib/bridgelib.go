// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridgelib

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

// BridgelibMetaData contains all meta data concerning the Bridgelib contract.
var BridgelibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"NotValidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INIT_BYTECODE_TRANSPARENT_PROXY\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"safeDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"safeName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"safeSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"permitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"expectedOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expectedSpender\",\"type\":\"address\"}],\"name\":\"validateAndProcessPermit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50611bbc8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610060575f3560e01c8063737ce34114610064578063a28fa4a31461008d578063be3dcf62146100b0578063c00f14ab146100d5578063c514f24e146100e8578063cf825e55146100f0575b5f5ffd5b6100776100723660046108de565b610103565b604051610084919061094d565b60405180910390f35b6100a061009b36600461095f565b6101c8565b6040519015158152602001610084565b6100c36100be3660046108de565b61056c565b60405160ff9091168152602001610084565b6100776100e33660046108de565b610620565b610077610665565b6100776100fe3660046108de565b610684565b60408051600481526024810182526020810180516001600160e01b03166395d89b4160e01b17905290516060915f9182916001600160a01b038616916101499190610a02565b5f60405180830381855afa9150503d805f8114610181576040519150601f19603f3d011682016040523d82523d5f602084013e610186565b606091505b5091509150816101b757604051806040016040528060098152602001681393d7d4d6535093d360ba1b8152506101c0565b6101c081610736565b949350505050565b5f806101d76004828789610a1d565b6101e091610a44565b9050632afa533160e01b6001600160e01b031982160161038f575f5f5f5f5f5f5f8c8c600490809261021493929190610a1d565b8101906102219190610a8a565b96509650965096509650965096508a6001600160a01b0316876001600160a01b0316146102615760405163912ecce760e01b815260040160405180910390fd5b896001600160a01b0316866001600160a01b0316146102935760405163750643af60e01b815260040160405180910390fd5b5f8e6001600160a01b031663d505accf60e01b898989898989896040516024016102ff97969594939291906001600160a01b0397881681529590961660208601526040850193909352606084019190915260ff16608083015260a082015260c081019190915260e00190565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b031990941693909317909252905161033d9190610a02565b5f604051808303815f865af19150503d805f8114610376576040519150601f19603f3d011682016040523d82523d5f602084013e61037b565b606091505b50909a506105639950505050505050505050565b631c0d143d60e21b6001600160e01b031982160161054a575f5f5f5f5f5f5f5f8d8d60049080926103c293929190610a1d565b8101906103cf9190610af6565b975097509750975097509750975097508b6001600160a01b0316886001600160a01b0316146104115760405163912ecce760e01b815260040160405180910390fd5b8a6001600160a01b0316876001600160a01b0316146104435760405163750643af60e01b815260040160405180910390fd5b5f8f6001600160a01b0316638fcbaf0c60e01b8a8a8a8a8a8a8a8a6040516024016104b99897969594939291906001600160a01b039889168152969097166020870152604086019490945260608501929092521515608084015260ff1660a083015260c082015260e08101919091526101000190565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516104f79190610a02565b5f604051808303815f865af19150503d805f8114610530576040519150601f19603f3d011682016040523d82523d5f602084013e610535565b606091505b50909b506105639a5050505050505050505050565b604051637141605d60e11b815260040160405180910390fd5b95945050505050565b60408051600481526024810182526020810180516001600160e01b031663313ce56760e01b17905290515f91829182916001600160a01b038616916105b19190610a02565b5f60405180830381855afa9150503d805f81146105e9576040519150601f19603f3d011682016040523d82523d5f602084013e6105ee565b606091505b5091509150818015610601575080516020145b61060c5760126101c0565b808060200190518101906101c09190610b78565b606061062b82610684565b61063483610103565b61063d8461056c565b60405160200161064f93929190610b93565b6040516020818303038152906040529050919050565b60405180610f000160405280610ec88152602001610cbf610ec8913981565b60408051600481526024810182526020810180516001600160e01b03166306fdde0360e01b17905290516060915f9182916001600160a01b038616916106ca9190610a02565b5f60405180830381855afa9150503d805f8114610702576040519150601f19603f3d011682016040523d82523d5f602084013e610707565b606091505b5091509150816101b757604051806040016040528060078152602001664e4f5f4e414d4560c81b8152506101c0565b6060604082511061075b57818060200190518101906107559190610bdf565b92915050565b8151602003610889575f5b602081108015610795575082818151811061078357610783610c86565b01602001516001600160f81b03191615155b156107ac57806107a481610c9a565b915050610766565b805f036107e35750506040805180820190915260128152714e4f545f56414c49445f454e434f44494e4760701b6020820152919050565b5f8167ffffffffffffffff8111156107fd576107fd610bcb565b6040519080825280601f01601f191660200182016040528015610827576020820181803683370190505b5090505f5b828110156108815784818151811061084657610846610c86565b602001015160f81c60f81b82828151811061086357610863610c86565b60200101906001600160f81b03191690815f1a90535060010161082c565b509392505050565b50506040805180820190915260128152714e4f545f56414c49445f454e434f44494e4760701b602082015290565b919050565b6001600160a01b03811681146108d0575f5ffd5b50565b80356108b7816108bc565b5f602082840312156108ee575f5ffd5b81356108f9816108bc565b9392505050565b5f5b8381101561091a578181015183820152602001610902565b50505f910152565b5f8151808452610939816020860160208601610900565b601f01601f19169290920160200192915050565b602081525f6108f96020830184610922565b5f5f5f5f5f60808688031215610973575f5ffd5b853561097e816108bc565b9450602086013567ffffffffffffffff811115610999575f5ffd5b8601601f810188136109a9575f5ffd5b803567ffffffffffffffff8111156109bf575f5ffd5b8860208284010111156109d0575f5ffd5b6020919091019450925060408601356109e8816108bc565b91506109f6606087016108d3565b90509295509295909350565b5f8251610a13818460208701610900565b9190910192915050565b5f5f85851115610a2b575f5ffd5b83861115610a37575f5ffd5b5050820193919092039150565b80356001600160e01b03198116906004841015610a75576001600160e01b0319600485900360031b81901b82161691505b5092915050565b60ff811681146108d0575f5ffd5b5f5f5f5f5f5f5f60e0888a031215610aa0575f5ffd5b8735610aab816108bc565b96506020880135610abb816108bc565b955060408801359450606088013593506080880135610ad981610a7c565b9699959850939692959460a0840135945060c09093013592915050565b5f5f5f5f5f5f5f5f610100898b031215610b0e575f5ffd5b8835610b19816108bc565b97506020890135610b29816108bc565b9650604089013595506060890135945060808901358015158114610b4b575f5ffd5b935060a0890135610b5b81610a7c565b979a969950949793969295929450505060c08201359160e0013590565b5f60208284031215610b88575f5ffd5b81516108f981610a7c565b606081525f610ba56060830186610922565b8281036020840152610bb78186610922565b91505060ff83166040830152949350505050565b634e487b7160e01b5f52604160045260245ffd5b5f60208284031215610bef575f5ffd5b815167ffffffffffffffff811115610c05575f5ffd5b8201601f81018413610c15575f5ffd5b805167ffffffffffffffff811115610c2f57610c2f610bcb565b604051601f8201601f19908116603f0116810167ffffffffffffffff81118282101715610c5e57610c5e610bcb565b604052818152828201602001861015610c75575f5ffd5b610563826020830160208601610900565b634e487b7160e01b5f52603260045260245ffd5b5f60018201610cb757634e487b7160e01b5f52601160045260245ffd5b506001019056fe60806040819052631d97f74d60e11b81523390633b2fee9a90608490602090600481865afa158015610033573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100579190610433565b604080515f80825260208201909252905061007382825f6100e2565b50506100dd336001600160a01b03166338b8fbbb6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100b4573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100d89190610433565b61010d565b6104c8565b6100eb8361017a565b5f825111806100f75750805b156101085761010683836101b9565b505b505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61014c5f516020610e815f395f51905f52546001600160a01b031690565b604080516001600160a01b03928316815291841660208301520160405180910390a1610177816101e5565b50565b61018381610280565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606101de8383604051806060016040528060278152602001610ea160279139610314565b9392505050565b6001600160a01b03811661024f5760405162461bcd60e51b815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b805f516020610e815f395f51905f525b80546001600160a01b0319166001600160a01b039290921691909117905550565b6001600160a01b0381163b6102ed5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610246565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc61025f565b60605f5f856001600160a01b031685604051610330919061047b565b5f60405180830381855af49150503d805f8114610368576040519150601f19603f3d011682016040523d82523d5f602084013e61036d565b606091505b50909250905061037f86838387610389565b9695505050505050565b606083156103f75782515f036103f0576001600160a01b0385163b6103f05760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610246565b5081610401565b6104018383610409565b949350505050565b8151156104195781518083602001fd5b8060405162461bcd60e51b81526004016102469190610496565b5f60208284031215610443575f5ffd5b81516001600160a01b03811681146101de575f5ffd5b5f5b8381101561047357818101518382015260200161045b565b50505f910152565b5f825161048c818460208701610459565b9190910192915050565b602081525f82518060208401526104b4816040850160208701610459565b601f01601f19169190910160400192915050565b6109ac806104d55f395ff3fe60806040526004361061005d575f3560e01c80635c60da1b116100425780635c60da1b146100a65780638f283970146100e3578063f851a440146101025761006c565b80633659cfe6146100745780634f1ef286146100935761006c565b3661006c5761006a610116565b005b61006a610116565b34801561007f575f5ffd5b5061006a61008e366004610854565b610130565b61006a6100a136600461086d565b610178565b3480156100b1575f5ffd5b506100ba6101eb565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b3480156100ee575f5ffd5b5061006a6100fd366004610854565b610228565b34801561010d575f5ffd5b506100ba610255565b61011e610282565b61012e610129610359565b610362565b565b610138610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101705761016d8160405180602001604052805f8152505f6103bf565b50565b61016d610116565b610180610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101e3576101de8383838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250600192506103bf915050565b505050565b6101de610116565b5f6101f4610380565b73ffffffffffffffffffffffffffffffffffffffff16330361021d57610218610359565b905090565b610225610116565b90565b610230610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101705761016d816103e9565b5f61025e610380565b73ffffffffffffffffffffffffffffffffffffffff16330361021d57610218610380565b61028a610380565b73ffffffffffffffffffffffffffffffffffffffff16330361012e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604260248201527f5472616e73706172656e745570677261646561626c6550726f78793a2061646d60448201527f696e2063616e6e6f742066616c6c6261636b20746f2070726f7879207461726760648201527f6574000000000000000000000000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b5f61021861044a565b365f5f375f5f365f845af43d5f5f3e80801561037c573d5ff35b3d5ffd5b5f7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b5473ffffffffffffffffffffffffffffffffffffffff16919050565b6103c883610471565b5f825111806103d45750805b156101de576103e383836104bd565b50505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f610412610380565b6040805173ffffffffffffffffffffffffffffffffffffffff928316815291841660208301520160405180910390a161016d816104e9565b5f7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6103a3565b61047a816105f5565b60405173ffffffffffffffffffffffffffffffffffffffff8216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606104e28383604051806060016040528060278152602001610979602791396106c0565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff811661058c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610350565b807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b73ffffffffffffffffffffffffffffffffffffffff81163b610699576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201527f6f74206120636f6e7472616374000000000000000000000000000000000000006064820152608401610350565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6105af565b60605f5f8573ffffffffffffffffffffffffffffffffffffffff16856040516106e9919061090d565b5f60405180830381855af49150503d805f8114610721576040519150601f19603f3d011682016040523d82523d5f602084013e610726565b606091505b509150915061073786838387610741565b9695505050505050565b606083156107d65782515f036107cf5773ffffffffffffffffffffffffffffffffffffffff85163b6107cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610350565b50816107e0565b6107e083836107e8565b949350505050565b8151156107f85781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103509190610928565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f575f5ffd5b919050565b5f60208284031215610864575f5ffd5b6104e28261082c565b5f5f5f6040848603121561087f575f5ffd5b6108888461082c565b9250602084013567ffffffffffffffff8111156108a3575f5ffd5b8401601f810186136108b3575f5ffd5b803567ffffffffffffffff8111156108c9575f5ffd5b8660208284010111156108da575f5ffd5b939660209190910195509293505050565b5f5b838110156109055781810151838201526020016108ed565b50505f910152565b5f825161091e8184602087016108eb565b9190910192915050565b602081525f82518060208401526109468160408501602087016108eb565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a164736f6c634300081c000ab53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220e756f10dcbbd3136d1778842f98489f6878e3593f87ade778daf1023ff70da6764736f6c634300081c0033",
}

// BridgelibABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgelibMetaData.ABI instead.
var BridgelibABI = BridgelibMetaData.ABI

// BridgelibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgelibMetaData.Bin instead.
var BridgelibBin = BridgelibMetaData.Bin

// DeployBridgelib deploys a new Ethereum contract, binding an instance of Bridgelib to it.
func DeployBridgelib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bridgelib, error) {
	parsed, err := BridgelibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgelibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridgelib{BridgelibCaller: BridgelibCaller{contract: contract}, BridgelibTransactor: BridgelibTransactor{contract: contract}, BridgelibFilterer: BridgelibFilterer{contract: contract}}, nil
}

// Bridgelib is an auto generated Go binding around an Ethereum contract.
type Bridgelib struct {
	BridgelibCaller     // Read-only binding to the contract
	BridgelibTransactor // Write-only binding to the contract
	BridgelibFilterer   // Log filterer for contract events
}

// BridgelibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgelibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgelibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgelibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgelibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgelibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgelibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgelibSession struct {
	Contract     *Bridgelib        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgelibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgelibCallerSession struct {
	Contract *BridgelibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BridgelibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgelibTransactorSession struct {
	Contract     *BridgelibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BridgelibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgelibRaw struct {
	Contract *Bridgelib // Generic contract binding to access the raw methods on
}

// BridgelibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgelibCallerRaw struct {
	Contract *BridgelibCaller // Generic read-only contract binding to access the raw methods on
}

// BridgelibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgelibTransactorRaw struct {
	Contract *BridgelibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgelib creates a new instance of Bridgelib, bound to a specific deployed contract.
func NewBridgelib(address common.Address, backend bind.ContractBackend) (*Bridgelib, error) {
	contract, err := bindBridgelib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridgelib{BridgelibCaller: BridgelibCaller{contract: contract}, BridgelibTransactor: BridgelibTransactor{contract: contract}, BridgelibFilterer: BridgelibFilterer{contract: contract}}, nil
}

// NewBridgelibCaller creates a new read-only instance of Bridgelib, bound to a specific deployed contract.
func NewBridgelibCaller(address common.Address, caller bind.ContractCaller) (*BridgelibCaller, error) {
	contract, err := bindBridgelib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgelibCaller{contract: contract}, nil
}

// NewBridgelibTransactor creates a new write-only instance of Bridgelib, bound to a specific deployed contract.
func NewBridgelibTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgelibTransactor, error) {
	contract, err := bindBridgelib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgelibTransactor{contract: contract}, nil
}

// NewBridgelibFilterer creates a new log filterer instance of Bridgelib, bound to a specific deployed contract.
func NewBridgelibFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgelibFilterer, error) {
	contract, err := bindBridgelib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgelibFilterer{contract: contract}, nil
}

// bindBridgelib binds a generic wrapper to an already deployed contract.
func bindBridgelib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgelibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridgelib *BridgelibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridgelib.Contract.BridgelibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridgelib *BridgelibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgelib.Contract.BridgelibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridgelib *BridgelibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridgelib.Contract.BridgelibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridgelib *BridgelibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridgelib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridgelib *BridgelibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgelib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridgelib *BridgelibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridgelib.Contract.contract.Transact(opts, method, params...)
}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() view returns(bytes)
func (_Bridgelib *BridgelibCaller) INITBYTECODETRANSPARENTPROXY(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Bridgelib.contract.Call(opts, &out, "INIT_BYTECODE_TRANSPARENT_PROXY")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() view returns(bytes)
func (_Bridgelib *BridgelibSession) INITBYTECODETRANSPARENTPROXY() ([]byte, error) {
	return _Bridgelib.Contract.INITBYTECODETRANSPARENTPROXY(&_Bridgelib.CallOpts)
}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() view returns(bytes)
func (_Bridgelib *BridgelibCallerSession) INITBYTECODETRANSPARENTPROXY() ([]byte, error) {
	return _Bridgelib.Contract.INITBYTECODETRANSPARENTPROXY(&_Bridgelib.CallOpts)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Bridgelib *BridgelibCaller) GetTokenMetadata(opts *bind.CallOpts, token common.Address) ([]byte, error) {
	var out []interface{}
	err := _Bridgelib.contract.Call(opts, &out, "getTokenMetadata", token)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Bridgelib *BridgelibSession) GetTokenMetadata(token common.Address) ([]byte, error) {
	return _Bridgelib.Contract.GetTokenMetadata(&_Bridgelib.CallOpts, token)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Bridgelib *BridgelibCallerSession) GetTokenMetadata(token common.Address) ([]byte, error) {
	return _Bridgelib.Contract.GetTokenMetadata(&_Bridgelib.CallOpts, token)
}

// SafeDecimals is a free data retrieval call binding the contract method 0xbe3dcf62.
//
// Solidity: function safeDecimals(address token) view returns(uint8)
func (_Bridgelib *BridgelibCaller) SafeDecimals(opts *bind.CallOpts, token common.Address) (uint8, error) {
	var out []interface{}
	err := _Bridgelib.contract.Call(opts, &out, "safeDecimals", token)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SafeDecimals is a free data retrieval call binding the contract method 0xbe3dcf62.
//
// Solidity: function safeDecimals(address token) view returns(uint8)
func (_Bridgelib *BridgelibSession) SafeDecimals(token common.Address) (uint8, error) {
	return _Bridgelib.Contract.SafeDecimals(&_Bridgelib.CallOpts, token)
}

// SafeDecimals is a free data retrieval call binding the contract method 0xbe3dcf62.
//
// Solidity: function safeDecimals(address token) view returns(uint8)
func (_Bridgelib *BridgelibCallerSession) SafeDecimals(token common.Address) (uint8, error) {
	return _Bridgelib.Contract.SafeDecimals(&_Bridgelib.CallOpts, token)
}

// SafeName is a free data retrieval call binding the contract method 0xcf825e55.
//
// Solidity: function safeName(address token) view returns(string)
func (_Bridgelib *BridgelibCaller) SafeName(opts *bind.CallOpts, token common.Address) (string, error) {
	var out []interface{}
	err := _Bridgelib.contract.Call(opts, &out, "safeName", token)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SafeName is a free data retrieval call binding the contract method 0xcf825e55.
//
// Solidity: function safeName(address token) view returns(string)
func (_Bridgelib *BridgelibSession) SafeName(token common.Address) (string, error) {
	return _Bridgelib.Contract.SafeName(&_Bridgelib.CallOpts, token)
}

// SafeName is a free data retrieval call binding the contract method 0xcf825e55.
//
// Solidity: function safeName(address token) view returns(string)
func (_Bridgelib *BridgelibCallerSession) SafeName(token common.Address) (string, error) {
	return _Bridgelib.Contract.SafeName(&_Bridgelib.CallOpts, token)
}

// SafeSymbol is a free data retrieval call binding the contract method 0x737ce341.
//
// Solidity: function safeSymbol(address token) view returns(string)
func (_Bridgelib *BridgelibCaller) SafeSymbol(opts *bind.CallOpts, token common.Address) (string, error) {
	var out []interface{}
	err := _Bridgelib.contract.Call(opts, &out, "safeSymbol", token)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SafeSymbol is a free data retrieval call binding the contract method 0x737ce341.
//
// Solidity: function safeSymbol(address token) view returns(string)
func (_Bridgelib *BridgelibSession) SafeSymbol(token common.Address) (string, error) {
	return _Bridgelib.Contract.SafeSymbol(&_Bridgelib.CallOpts, token)
}

// SafeSymbol is a free data retrieval call binding the contract method 0x737ce341.
//
// Solidity: function safeSymbol(address token) view returns(string)
func (_Bridgelib *BridgelibCallerSession) SafeSymbol(token common.Address) (string, error) {
	return _Bridgelib.Contract.SafeSymbol(&_Bridgelib.CallOpts, token)
}

// ValidateAndProcessPermit is a paid mutator transaction binding the contract method 0xa28fa4a3.
//
// Solidity: function validateAndProcessPermit(address token, bytes permitData, address expectedOwner, address expectedSpender) returns(bool success)
func (_Bridgelib *BridgelibTransactor) ValidateAndProcessPermit(opts *bind.TransactOpts, token common.Address, permitData []byte, expectedOwner common.Address, expectedSpender common.Address) (*types.Transaction, error) {
	return _Bridgelib.contract.Transact(opts, "validateAndProcessPermit", token, permitData, expectedOwner, expectedSpender)
}

// ValidateAndProcessPermit is a paid mutator transaction binding the contract method 0xa28fa4a3.
//
// Solidity: function validateAndProcessPermit(address token, bytes permitData, address expectedOwner, address expectedSpender) returns(bool success)
func (_Bridgelib *BridgelibSession) ValidateAndProcessPermit(token common.Address, permitData []byte, expectedOwner common.Address, expectedSpender common.Address) (*types.Transaction, error) {
	return _Bridgelib.Contract.ValidateAndProcessPermit(&_Bridgelib.TransactOpts, token, permitData, expectedOwner, expectedSpender)
}

// ValidateAndProcessPermit is a paid mutator transaction binding the contract method 0xa28fa4a3.
//
// Solidity: function validateAndProcessPermit(address token, bytes permitData, address expectedOwner, address expectedSpender) returns(bool success)
func (_Bridgelib *BridgelibTransactorSession) ValidateAndProcessPermit(token common.Address, permitData []byte, expectedOwner common.Address, expectedSpender common.Address) (*types.Transaction, error) {
	return _Bridgelib.Contract.ValidateAndProcessPermit(&_Bridgelib.TransactOpts, token, permitData, expectedOwner, expectedSpender)
}

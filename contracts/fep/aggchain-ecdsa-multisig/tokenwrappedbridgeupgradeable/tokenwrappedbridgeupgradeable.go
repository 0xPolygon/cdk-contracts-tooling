// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenwrappedbridgeupgradeable

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

// TokenwrappedbridgeupgradeableMetaData contains all meta data concerning the Tokenwrappedbridgeupgradeable contract.
var TokenwrappedbridgeupgradeableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyBridge\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"__decimals\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060156019565b60c9565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560685760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c65780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6116fb806100d65f395ff3fe608060405234801561000f575f5ffd5b5060043610610115575f3560e01c806370a08231116100ad5780639dc29fac1161007d578063a9059cbb11610063578063a9059cbb146102a6578063d505accf146102b9578063dd62ed3e146102cc575f5ffd5b80639dc29fac1461024b578063a3c573eb1461025e575f5ffd5b806370a08231146102025780637ecebe001461021557806384b0196e1461022857806395d89b4114610243575f5ffd5b806323b872dd116100e857806323b872dd146101a0578063313ce567146101b35780633644e515146101e757806340c10f19146101ef575f5ffd5b806306fdde0314610119578063095ea7b3146101375780631624f6c61461015a57806318160ddd1461016f575b5f5ffd5b610121610323565b60405161012e919061125c565b60405180910390f35b61014a610145366004611290565b6103db565b604051901515815260200161012e565b61016d610168366004611367565b6103f4565b005b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02545b60405190815260200161012e565b61014a6101ae3660046113db565b610576565b7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d005460405160ff909116815260200161012e565b610192610599565b61016d6101fd366004611290565b6105a7565b610192610210366004611415565b61060a565b610192610223366004611415565b61064d565b610230610657565b60405161012e979695949392919061142e565b610121610720565b61016d610259366004611290565b610771565b7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d005461010090046001600160a01b03166040516001600160a01b03909116815260200161012e565b61014a6102b4366004611290565b6107cf565b61016d6102c73660046114c4565b6107dc565b6101926102da36600461152a565b6001600160a01b039182165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b60605f7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b90508060030180546103599061155b565b80601f01602080910402602001604051908101604052809291908181526020018280546103859061155b565b80156103d05780601f106103a7576101008083540402835291602001916103d0565b820191905f5260205f20905b8154815290600101906020018083116103b357829003601f168201915b505050505091505090565b5f336103e8818585610931565b60019150505b92915050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f8115801561043e5750825b90505f8267ffffffffffffffff16600114801561045a5750303b155b905081158015610468575080155b156104865760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156104ba57845468ff00000000000000001916680100000000000000001785555b6104c4888861093e565b6104cd88610954565b7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d00805460ff88167fffffffffffffffffffffff000000000000000000000000000000000000000000909116176101003302179055831561056c57845468ff000000000000000019168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b5f33610583858285610982565b61058e858585610a1c565b506001949350505050565b5f6105a2610a79565b905090565b5f7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d00805490915061010090046001600160a01b031633146105fb576040516338da3b1560e01b815260040160405180910390fd5b6106058383610a82565b505050565b5f807f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b6001600160a01b039093165f9081526020939093525050604090205490565b5f6103ee82610ab6565b5f60608082808083817fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100805490915015801561069557506001810154155b6106e65760405162461bcd60e51b815260206004820152601560248201527f4549503731323a20556e696e697469616c697a6564000000000000000000000060448201526064015b60405180910390fd5b6106ee610ac0565b6106f6610b11565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060917f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00916103599061155b565b5f7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d00805490915061010090046001600160a01b031633146107c5576040516338da3b1560e01b815260040160405180910390fd5b6106058383610b3a565b5f336103e8818585610a1c565b834211156108005760405163313c898160e11b8152600481018590526024016106dd565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c988888861086a8c6001600160a01b03165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f6108c482610b6e565b90505f6108d382878787610b9a565b9050896001600160a01b0316816001600160a01b03161461091a576040516325c0072360e11b81526001600160a01b0380831660048301528b1660248201526044016106dd565b6109258a8a8a610931565b50505050505050505050565b6106058383836001610bc6565b610946610cbd565b6109508282610d0d565b5050565b61095c610cbd565b61097f81604051806040016040528060018152602001603160f81b815250610d70565b50565b6001600160a01b038381165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0160209081526040808320938616835292905220545f198114610a165781811015610a0857604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064016106dd565b610a1684848484035f610bc6565b50505050565b6001600160a01b038316610a4557604051634b637e8f60e11b81525f60048201526024016106dd565b6001600160a01b038216610a6e5760405163ec442f0560e01b81525f60048201526024016106dd565b610605838383610de2565b5f6105a2610f2e565b6001600160a01b038216610aab5760405163ec442f0560e01b81525f60048201526024016106dd565b6109505f8383610de2565b5f6103ee82610fa1565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060917fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100916103599061155b565b60605f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100610348565b6001600160a01b038216610b6357604051634b637e8f60e11b81525f60048201526024016106dd565b610950825f83610de2565b5f6103ee610b7a610a79565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f610baa88888888610fc9565b925092509250610bba8282611091565b50909695505050505050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace006001600160a01b038516610c105760405163e602df0560e01b81525f60048201526024016106dd565b6001600160a01b038416610c3957604051634a1406b160e11b81525f60048201526024016106dd565b6001600160a01b038086165f90815260018301602090815260408083209388168352929052208390558115610cb657836001600160a01b0316856001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92585604051610cad91815260200190565b60405180910390a35b5050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16610d0b57604051631afcd79f60e31b815260040160405180910390fd5b565b610d15610cbd565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace007f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace03610d6184826115d7565b5060048101610a1683826115d7565b610d78610cbd565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1007fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102610dc484826115d7565b5060038101610dd383826115d7565b505f8082556001909101555050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace006001600160a01b038416610e2f5781816002015f828254610e249190611692565b90915550610e9f9050565b6001600160a01b0384165f9081526020829052604090205482811015610e815760405163391434e360e21b81526001600160a01b038616600482015260248101829052604481018490526064016106dd565b6001600160a01b0385165f9081526020839052604090209083900390555b6001600160a01b038316610ebd576002810180548390039055610edb565b6001600160a01b0383165f9081526020829052604090208054830190555b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610f2091815260200190565b60405180910390a350505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f610f58611149565b610f606111c4565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f807f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0061062e565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561100257505f91506003905082611087565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015611053573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b03811661107e57505f925060019150829050611087565b92505f91508190505b9450945094915050565b5f8260038111156110a4576110a46116b1565b036110ad575050565b60018260038111156110c1576110c16116b1565b036110df5760405163f645eedf60e01b815260040160405180910390fd5b60028260038111156110f3576110f36116b1565b036111145760405163fce698f760e01b8152600481018290526024016106dd565b6003826003811115611128576111286116b1565b03610950576040516335e2f38360e21b8152600481018290526024016106dd565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10081611174610ac0565b80519091501561118c57805160209091012092915050565b8154801561119b579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100816111ef610b11565b80519091501561120757805160209091012092915050565b6001820154801561119b579392505050565b5f81518084525f5b8181101561123d57602081850181015186830182015201611221565b505f602082860101526020601f19601f83011685010191505092915050565b602081525f61126e6020830184611219565b9392505050565b80356001600160a01b038116811461128b575f5ffd5b919050565b5f5f604083850312156112a1575f5ffd5b6112aa83611275565b946020939093013593505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f8301126112db575f5ffd5b813567ffffffffffffffff8111156112f5576112f56112b8565b604051601f8201601f19908116603f0116810167ffffffffffffffff81118282101715611324576113246112b8565b60405281815283820160200185101561133b575f5ffd5b816020850160208301375f918101602001919091529392505050565b803560ff8116811461128b575f5ffd5b5f5f5f60608486031215611379575f5ffd5b833567ffffffffffffffff81111561138f575f5ffd5b61139b868287016112cc565b935050602084013567ffffffffffffffff8111156113b7575f5ffd5b6113c3868287016112cc565b9250506113d260408501611357565b90509250925092565b5f5f5f606084860312156113ed575f5ffd5b6113f684611275565b925061140460208501611275565b929592945050506040919091013590565b5f60208284031215611425575f5ffd5b61126e82611275565b60ff60f81b8816815260e060208201525f61144c60e0830189611219565b828103604084015261145e8189611219565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b818110156114b3578351835260209384019390920191600101611495565b50909b9a5050505050505050505050565b5f5f5f5f5f5f5f60e0888a0312156114da575f5ffd5b6114e388611275565b96506114f160208901611275565b9550604088013594506060880135935061150d60808901611357565b9699959850939692959460a0840135945060c09093013592915050565b5f5f6040838503121561153b575f5ffd5b61154483611275565b915061155260208401611275565b90509250929050565b600181811c9082168061156f57607f821691505b60208210810361158d57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561060557805f5260205f20601f840160051c810160208510156115b85750805b601f840160051c820191505b81811015610cb6575f81556001016115c4565b815167ffffffffffffffff8111156115f1576115f16112b8565b611605816115ff845461155b565b84611593565b6020601f821160018114611637575f83156116205750848201515b5f19600385901b1c1916600184901b178455610cb6565b5f84815260208120601f198516915b828110156116665787850151825560209485019460019092019101611646565b508482101561168357868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b808201808211156103ee57634e487b7160e01b5f52601160045260245ffd5b634e487b7160e01b5f52602160045260245ffdfea2646970667358221220a561c2f5d53554a9eb5f9ce0194d9c8147ec03d3ba1e77539eb6b379627db98d64736f6c634300081c0033",
}

// TokenwrappedbridgeupgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenwrappedbridgeupgradeableMetaData.ABI instead.
var TokenwrappedbridgeupgradeableABI = TokenwrappedbridgeupgradeableMetaData.ABI

// TokenwrappedbridgeupgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenwrappedbridgeupgradeableMetaData.Bin instead.
var TokenwrappedbridgeupgradeableBin = TokenwrappedbridgeupgradeableMetaData.Bin

// DeployTokenwrappedbridgeupgradeable deploys a new Ethereum contract, binding an instance of Tokenwrappedbridgeupgradeable to it.
func DeployTokenwrappedbridgeupgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tokenwrappedbridgeupgradeable, error) {
	parsed, err := TokenwrappedbridgeupgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenwrappedbridgeupgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tokenwrappedbridgeupgradeable{TokenwrappedbridgeupgradeableCaller: TokenwrappedbridgeupgradeableCaller{contract: contract}, TokenwrappedbridgeupgradeableTransactor: TokenwrappedbridgeupgradeableTransactor{contract: contract}, TokenwrappedbridgeupgradeableFilterer: TokenwrappedbridgeupgradeableFilterer{contract: contract}}, nil
}

// Tokenwrappedbridgeupgradeable is an auto generated Go binding around an Ethereum contract.
type Tokenwrappedbridgeupgradeable struct {
	TokenwrappedbridgeupgradeableCaller     // Read-only binding to the contract
	TokenwrappedbridgeupgradeableTransactor // Write-only binding to the contract
	TokenwrappedbridgeupgradeableFilterer   // Log filterer for contract events
}

// TokenwrappedbridgeupgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenwrappedbridgeupgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenwrappedbridgeupgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenwrappedbridgeupgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenwrappedbridgeupgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenwrappedbridgeupgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenwrappedbridgeupgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenwrappedbridgeupgradeableSession struct {
	Contract     *Tokenwrappedbridgeupgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TokenwrappedbridgeupgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenwrappedbridgeupgradeableCallerSession struct {
	Contract *TokenwrappedbridgeupgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// TokenwrappedbridgeupgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenwrappedbridgeupgradeableTransactorSession struct {
	Contract     *TokenwrappedbridgeupgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// TokenwrappedbridgeupgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenwrappedbridgeupgradeableRaw struct {
	Contract *Tokenwrappedbridgeupgradeable // Generic contract binding to access the raw methods on
}

// TokenwrappedbridgeupgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenwrappedbridgeupgradeableCallerRaw struct {
	Contract *TokenwrappedbridgeupgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// TokenwrappedbridgeupgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenwrappedbridgeupgradeableTransactorRaw struct {
	Contract *TokenwrappedbridgeupgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenwrappedbridgeupgradeable creates a new instance of Tokenwrappedbridgeupgradeable, bound to a specific deployed contract.
func NewTokenwrappedbridgeupgradeable(address common.Address, backend bind.ContractBackend) (*Tokenwrappedbridgeupgradeable, error) {
	contract, err := bindTokenwrappedbridgeupgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokenwrappedbridgeupgradeable{TokenwrappedbridgeupgradeableCaller: TokenwrappedbridgeupgradeableCaller{contract: contract}, TokenwrappedbridgeupgradeableTransactor: TokenwrappedbridgeupgradeableTransactor{contract: contract}, TokenwrappedbridgeupgradeableFilterer: TokenwrappedbridgeupgradeableFilterer{contract: contract}}, nil
}

// NewTokenwrappedbridgeupgradeableCaller creates a new read-only instance of Tokenwrappedbridgeupgradeable, bound to a specific deployed contract.
func NewTokenwrappedbridgeupgradeableCaller(address common.Address, caller bind.ContractCaller) (*TokenwrappedbridgeupgradeableCaller, error) {
	contract, err := bindTokenwrappedbridgeupgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenwrappedbridgeupgradeableCaller{contract: contract}, nil
}

// NewTokenwrappedbridgeupgradeableTransactor creates a new write-only instance of Tokenwrappedbridgeupgradeable, bound to a specific deployed contract.
func NewTokenwrappedbridgeupgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenwrappedbridgeupgradeableTransactor, error) {
	contract, err := bindTokenwrappedbridgeupgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenwrappedbridgeupgradeableTransactor{contract: contract}, nil
}

// NewTokenwrappedbridgeupgradeableFilterer creates a new log filterer instance of Tokenwrappedbridgeupgradeable, bound to a specific deployed contract.
func NewTokenwrappedbridgeupgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenwrappedbridgeupgradeableFilterer, error) {
	contract, err := bindTokenwrappedbridgeupgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenwrappedbridgeupgradeableFilterer{contract: contract}, nil
}

// bindTokenwrappedbridgeupgradeable binds a generic wrapper to an already deployed contract.
func bindTokenwrappedbridgeupgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenwrappedbridgeupgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenwrappedbridgeupgradeable.Contract.TokenwrappedbridgeupgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.TokenwrappedbridgeupgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.TokenwrappedbridgeupgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenwrappedbridgeupgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tokenwrappedbridgeupgradeable.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.DOMAINSEPARATOR(&_Tokenwrappedbridgeupgradeable.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.DOMAINSEPARATOR(&_Tokenwrappedbridgeupgradeable.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tokenwrappedbridgeupgradeable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.Allowance(&_Tokenwrappedbridgeupgradeable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.Allowance(&_Tokenwrappedbridgeupgradeable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tokenwrappedbridgeupgradeable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.BalanceOf(&_Tokenwrappedbridgeupgradeable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.BalanceOf(&_Tokenwrappedbridgeupgradeable.CallOpts, account)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenwrappedbridgeupgradeable.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableSession) BridgeAddress() (common.Address, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.BridgeAddress(&_Tokenwrappedbridgeupgradeable.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableCallerSession) BridgeAddress() (common.Address, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.BridgeAddress(&_Tokenwrappedbridgeupgradeable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenwrappedbridgeupgradeable.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableSession) Decimals() (uint8, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.Decimals(&_Tokenwrappedbridgeupgradeable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableCallerSession) Decimals() (uint8, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.Decimals(&_Tokenwrappedbridgeupgradeable.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Tokenwrappedbridgeupgradeable.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.Eip712Domain(&_Tokenwrappedbridgeupgradeable.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.Eip712Domain(&_Tokenwrappedbridgeupgradeable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tokenwrappedbridgeupgradeable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableSession) Name() (string, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.Name(&_Tokenwrappedbridgeupgradeable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableCallerSession) Name() (string, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.Name(&_Tokenwrappedbridgeupgradeable.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tokenwrappedbridgeupgradeable.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.Nonces(&_Tokenwrappedbridgeupgradeable.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.Nonces(&_Tokenwrappedbridgeupgradeable.CallOpts, owner)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tokenwrappedbridgeupgradeable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableSession) Symbol() (string, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.Symbol(&_Tokenwrappedbridgeupgradeable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableCallerSession) Symbol() (string, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.Symbol(&_Tokenwrappedbridgeupgradeable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokenwrappedbridgeupgradeable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableSession) TotalSupply() (*big.Int, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.TotalSupply(&_Tokenwrappedbridgeupgradeable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableCallerSession) TotalSupply() (*big.Int, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.TotalSupply(&_Tokenwrappedbridgeupgradeable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Tokenwrappedbridgeupgradeable.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.Approve(&_Tokenwrappedbridgeupgradeable.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.Approve(&_Tokenwrappedbridgeupgradeable.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 value) returns()
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableTransactor) Burn(opts *bind.TransactOpts, account common.Address, value *big.Int) (*types.Transaction, error) {
	return _Tokenwrappedbridgeupgradeable.contract.Transact(opts, "burn", account, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 value) returns()
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableSession) Burn(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.Burn(&_Tokenwrappedbridgeupgradeable.TransactOpts, account, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 value) returns()
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableTransactorSession) Burn(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.Burn(&_Tokenwrappedbridgeupgradeable.TransactOpts, account, value)
}

// Initialize is a paid mutator transaction binding the contract method 0x1624f6c6.
//
// Solidity: function initialize(string name, string symbol, uint8 __decimals) returns()
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, __decimals uint8) (*types.Transaction, error) {
	return _Tokenwrappedbridgeupgradeable.contract.Transact(opts, "initialize", name, symbol, __decimals)
}

// Initialize is a paid mutator transaction binding the contract method 0x1624f6c6.
//
// Solidity: function initialize(string name, string symbol, uint8 __decimals) returns()
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableSession) Initialize(name string, symbol string, __decimals uint8) (*types.Transaction, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.Initialize(&_Tokenwrappedbridgeupgradeable.TransactOpts, name, symbol, __decimals)
}

// Initialize is a paid mutator transaction binding the contract method 0x1624f6c6.
//
// Solidity: function initialize(string name, string symbol, uint8 __decimals) returns()
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableTransactorSession) Initialize(name string, symbol string, __decimals uint8) (*types.Transaction, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.Initialize(&_Tokenwrappedbridgeupgradeable.TransactOpts, name, symbol, __decimals)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns()
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableTransactor) Mint(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Tokenwrappedbridgeupgradeable.contract.Transact(opts, "mint", to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns()
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.Mint(&_Tokenwrappedbridgeupgradeable.TransactOpts, to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns()
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableTransactorSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.Mint(&_Tokenwrappedbridgeupgradeable.TransactOpts, to, value)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Tokenwrappedbridgeupgradeable.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.Permit(&_Tokenwrappedbridgeupgradeable.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.Permit(&_Tokenwrappedbridgeupgradeable.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Tokenwrappedbridgeupgradeable.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.Transfer(&_Tokenwrappedbridgeupgradeable.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.Transfer(&_Tokenwrappedbridgeupgradeable.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Tokenwrappedbridgeupgradeable.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.TransferFrom(&_Tokenwrappedbridgeupgradeable.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Tokenwrappedbridgeupgradeable.Contract.TransferFrom(&_Tokenwrappedbridgeupgradeable.TransactOpts, from, to, value)
}

// TokenwrappedbridgeupgradeableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Tokenwrappedbridgeupgradeable contract.
type TokenwrappedbridgeupgradeableApprovalIterator struct {
	Event *TokenwrappedbridgeupgradeableApproval // Event containing the contract specifics and raw log

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
func (it *TokenwrappedbridgeupgradeableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenwrappedbridgeupgradeableApproval)
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
		it.Event = new(TokenwrappedbridgeupgradeableApproval)
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
func (it *TokenwrappedbridgeupgradeableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenwrappedbridgeupgradeableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenwrappedbridgeupgradeableApproval represents a Approval event raised by the Tokenwrappedbridgeupgradeable contract.
type TokenwrappedbridgeupgradeableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TokenwrappedbridgeupgradeableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Tokenwrappedbridgeupgradeable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenwrappedbridgeupgradeableApprovalIterator{contract: _Tokenwrappedbridgeupgradeable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenwrappedbridgeupgradeableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Tokenwrappedbridgeupgradeable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenwrappedbridgeupgradeableApproval)
				if err := _Tokenwrappedbridgeupgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableFilterer) ParseApproval(log types.Log) (*TokenwrappedbridgeupgradeableApproval, error) {
	event := new(TokenwrappedbridgeupgradeableApproval)
	if err := _Tokenwrappedbridgeupgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenwrappedbridgeupgradeableEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Tokenwrappedbridgeupgradeable contract.
type TokenwrappedbridgeupgradeableEIP712DomainChangedIterator struct {
	Event *TokenwrappedbridgeupgradeableEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *TokenwrappedbridgeupgradeableEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenwrappedbridgeupgradeableEIP712DomainChanged)
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
		it.Event = new(TokenwrappedbridgeupgradeableEIP712DomainChanged)
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
func (it *TokenwrappedbridgeupgradeableEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenwrappedbridgeupgradeableEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenwrappedbridgeupgradeableEIP712DomainChanged represents a EIP712DomainChanged event raised by the Tokenwrappedbridgeupgradeable contract.
type TokenwrappedbridgeupgradeableEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*TokenwrappedbridgeupgradeableEIP712DomainChangedIterator, error) {

	logs, sub, err := _Tokenwrappedbridgeupgradeable.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &TokenwrappedbridgeupgradeableEIP712DomainChangedIterator{contract: _Tokenwrappedbridgeupgradeable.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *TokenwrappedbridgeupgradeableEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Tokenwrappedbridgeupgradeable.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenwrappedbridgeupgradeableEIP712DomainChanged)
				if err := _Tokenwrappedbridgeupgradeable.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableFilterer) ParseEIP712DomainChanged(log types.Log) (*TokenwrappedbridgeupgradeableEIP712DomainChanged, error) {
	event := new(TokenwrappedbridgeupgradeableEIP712DomainChanged)
	if err := _Tokenwrappedbridgeupgradeable.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenwrappedbridgeupgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Tokenwrappedbridgeupgradeable contract.
type TokenwrappedbridgeupgradeableInitializedIterator struct {
	Event *TokenwrappedbridgeupgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *TokenwrappedbridgeupgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenwrappedbridgeupgradeableInitialized)
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
		it.Event = new(TokenwrappedbridgeupgradeableInitialized)
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
func (it *TokenwrappedbridgeupgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenwrappedbridgeupgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenwrappedbridgeupgradeableInitialized represents a Initialized event raised by the Tokenwrappedbridgeupgradeable contract.
type TokenwrappedbridgeupgradeableInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*TokenwrappedbridgeupgradeableInitializedIterator, error) {

	logs, sub, err := _Tokenwrappedbridgeupgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TokenwrappedbridgeupgradeableInitializedIterator{contract: _Tokenwrappedbridgeupgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TokenwrappedbridgeupgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _Tokenwrappedbridgeupgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenwrappedbridgeupgradeableInitialized)
				if err := _Tokenwrappedbridgeupgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableFilterer) ParseInitialized(log types.Log) (*TokenwrappedbridgeupgradeableInitialized, error) {
	event := new(TokenwrappedbridgeupgradeableInitialized)
	if err := _Tokenwrappedbridgeupgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenwrappedbridgeupgradeableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Tokenwrappedbridgeupgradeable contract.
type TokenwrappedbridgeupgradeableTransferIterator struct {
	Event *TokenwrappedbridgeupgradeableTransfer // Event containing the contract specifics and raw log

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
func (it *TokenwrappedbridgeupgradeableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenwrappedbridgeupgradeableTransfer)
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
		it.Event = new(TokenwrappedbridgeupgradeableTransfer)
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
func (it *TokenwrappedbridgeupgradeableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenwrappedbridgeupgradeableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenwrappedbridgeupgradeableTransfer represents a Transfer event raised by the Tokenwrappedbridgeupgradeable contract.
type TokenwrappedbridgeupgradeableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenwrappedbridgeupgradeableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Tokenwrappedbridgeupgradeable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenwrappedbridgeupgradeableTransferIterator{contract: _Tokenwrappedbridgeupgradeable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenwrappedbridgeupgradeableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Tokenwrappedbridgeupgradeable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenwrappedbridgeupgradeableTransfer)
				if err := _Tokenwrappedbridgeupgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Tokenwrappedbridgeupgradeable *TokenwrappedbridgeupgradeableFilterer) ParseTransfer(log types.Log) (*TokenwrappedbridgeupgradeableTransfer, error) {
	event := new(TokenwrappedbridgeupgradeableTransfer)
	if err := _Tokenwrappedbridgeupgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

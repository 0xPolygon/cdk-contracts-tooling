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
	Bin: "0x6080604052348015600e575f5ffd5b5060156019565b60c9565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560685760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c65780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b61163e806100d65f395ff3fe608060405234801561000f575f5ffd5b5060043610610106575f3560e01c806370a082311161009e5780639dc29fac1161006e5780639dc29fac1461023c578063a3c573eb1461024f578063a9059cbb14610297578063d505accf146102aa578063dd62ed3e146102bd575f5ffd5b806370a08231146101f35780637ecebe001461020657806384b0196e1461021957806395d89b4114610234575f5ffd5b806323b872dd116100d957806323b872dd14610191578063313ce567146101a45780633644e515146101d857806340c10f19146101e0575f5ffd5b806306fdde031461010a578063095ea7b3146101285780631624f6c61461014b57806318160ddd14610160575b5f5ffd5b610112610314565b60405161011f919061115f565b60405180910390f35b61013b610136366004611193565b6103b9565b604051901515815260200161011f565b61015e61015936600461126a565b6103d2565b005b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02545b60405190815260200161011f565b61013b61019f3660046112de565b61054a565b7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d005460405160ff909116815260200161011f565b61018361056d565b61015e6101ee366004611193565b61057b565b610183610201366004611318565b6105de565b610183610214366004611318565b61060e565b610221610618565b60405161011f9796959493929190611331565b6101126106ce565b61015e61024a366004611193565b61070c565b7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d005461010090046001600160a01b03166040516001600160a01b03909116815260200161011f565b61013b6102a5366004611193565b61076a565b61015e6102b83660046113c7565b610777565b6101836102cb36600461142d565b6001600160a01b039182165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b60605f5f5160206115c95f395f51905f525b90508060030180546103379061145e565b80601f01602080910402602001604051908101604052809291908181526020018280546103639061145e565b80156103ae5780601f10610385576101008083540402835291602001916103ae565b820191905f5260205f20905b81548152906001019060200180831161039157829003601f168201915b505050505091505090565b5f336103c68185856108cc565b60019150505b92915050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f8115801561041c5750825b90505f8267ffffffffffffffff1660011480156104385750303b155b905081158015610446575080155b156104645760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561049857845468ff00000000000000001916680100000000000000001785555b6104a288886108d9565b6104ab886108ef565b7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d00805460ff881674ffffffffffffffffffffffffffffffffffffffffff19909116176101003302179055831561054057845468ff000000000000000019168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b5f3361055785828561091d565b6105628585856109b7565b506001949350505050565b5f610576610a14565b905090565b5f7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d00805490915061010090046001600160a01b031633146105cf576040516338da3b1560e01b815260040160405180910390fd5b6105d98383610a1d565b505050565b5f805f5160206115c95f395f51905f525b6001600160a01b039093165f9081526020939093525050604090205490565b5f6103cc82610a51565b5f60608082808083815f5160206115e95f395f51905f52805490915015801561064357506001810154155b6106945760405162461bcd60e51b815260206004820152601560248201527f4549503731323a20556e696e697469616c697a6564000000000000000000000060448201526064015b60405180910390fd5b61069c610a5b565b6106a4610a99565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060915f5160206115c95f395f51905f52916103379061145e565b5f7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d00805490915061010090046001600160a01b03163314610760576040516338da3b1560e01b815260040160405180910390fd5b6105d98383610aaf565b5f336103c68185856109b7565b8342111561079b5760405163313c898160e11b81526004810185905260240161068b565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98888886108058c6001600160a01b03165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f61085f82610ae3565b90505f61086e82878787610b0f565b9050896001600160a01b0316816001600160a01b0316146108b5576040516325c0072360e11b81526001600160a01b0380831660048301528b16602482015260440161068b565b6108c08a8a8a6108cc565b50505050505050505050565b6105d98383836001610b3b565b6108e1610c1f565b6108eb8282610c6f565b5050565b6108f7610c1f565b61091a81604051806040016040528060018152602001603160f81b815250610cbf565b50565b6001600160a01b038381165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0160209081526040808320938616835292905220545f1981146109b157818110156109a357604051637dc7a0d960e11b81526001600160a01b0384166004820152602481018290526044810183905260640161068b565b6109b184848484035f610b3b565b50505050565b6001600160a01b0383166109e057604051634b637e8f60e11b81525f600482015260240161068b565b6001600160a01b038216610a095760405163ec442f0560e01b81525f600482015260240161068b565b6105d9838383610d1e565b5f610576610e57565b6001600160a01b038216610a465760405163ec442f0560e01b81525f600482015260240161068b565b6108eb5f8383610d1e565b5f6103cc82610eca565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f5160206115e95f395f51905f52916103379061145e565b60605f5f5160206115e95f395f51905f52610326565b6001600160a01b038216610ad857604051634b637e8f60e11b81525f600482015260240161068b565b6108eb825f83610d1e565b5f6103cc610aef610a14565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f610b1f88888888610ef2565b925092509250610b2f8282610fba565b50909695505050505050565b5f5160206115c95f395f51905f526001600160a01b038516610b725760405163e602df0560e01b81525f600482015260240161068b565b6001600160a01b038416610b9b57604051634a1406b160e11b81525f600482015260240161068b565b6001600160a01b038086165f90815260018301602090815260408083209388168352929052208390558115610c1857836001600160a01b0316856001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92585604051610c0f91815260200190565b60405180910390a35b5050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16610c6d57604051631afcd79f60e31b815260040160405180910390fd5b565b610c77610c1f565b5f5160206115c95f395f51905f527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace03610cb084826114da565b50600481016109b183826114da565b610cc7610c1f565b5f5160206115e95f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102610d0084826114da565b5060038101610d0f83826114da565b505f8082556001909101555050565b5f5160206115c95f395f51905f526001600160a01b038416610d585781816002015f828254610d4d9190611595565b90915550610dc89050565b6001600160a01b0384165f9081526020829052604090205482811015610daa5760405163391434e360e21b81526001600160a01b0386166004820152602481018290526044810184905260640161068b565b6001600160a01b0385165f9081526020839052604090209083900390555b6001600160a01b038316610de6576002810180548390039055610e04565b6001600160a01b0383165f9081526020829052604090208054830190555b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610e4991815260200190565b60405180910390a350505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f610e81611072565b610e896110da565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f807f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006105ef565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115610f2b57505f91506003905082610fb0565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015610f7c573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116610fa757505f925060019150829050610fb0565b92505f91508190505b9450945094915050565b5f826003811115610fcd57610fcd6115b4565b03610fd6575050565b6001826003811115610fea57610fea6115b4565b036110085760405163f645eedf60e01b815260040160405180910390fd5b600282600381111561101c5761101c6115b4565b0361103d5760405163fce698f760e01b81526004810182905260240161068b565b6003826003811115611051576110516115b4565b036108eb576040516335e2f38360e21b81526004810182905260240161068b565b5f5f5160206115e95f395f51905f528161108a610a5b565b8051909150156110a257805160209091012092915050565b815480156110b1579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f5160206115e95f395f51905f52816110f2610a99565b80519091501561110a57805160209091012092915050565b600182015480156110b1579392505050565b5f81518084525f5b8181101561114057602081850181015186830182015201611124565b505f602082860101526020601f19601f83011685010191505092915050565b602081525f611171602083018461111c565b9392505050565b80356001600160a01b038116811461118e575f5ffd5b919050565b5f5f604083850312156111a4575f5ffd5b6111ad83611178565b946020939093013593505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f8301126111de575f5ffd5b813567ffffffffffffffff8111156111f8576111f86111bb565b604051601f8201601f19908116603f0116810167ffffffffffffffff81118282101715611227576112276111bb565b60405281815283820160200185101561123e575f5ffd5b816020850160208301375f918101602001919091529392505050565b803560ff8116811461118e575f5ffd5b5f5f5f6060848603121561127c575f5ffd5b833567ffffffffffffffff811115611292575f5ffd5b61129e868287016111cf565b935050602084013567ffffffffffffffff8111156112ba575f5ffd5b6112c6868287016111cf565b9250506112d56040850161125a565b90509250925092565b5f5f5f606084860312156112f0575f5ffd5b6112f984611178565b925061130760208501611178565b929592945050506040919091013590565b5f60208284031215611328575f5ffd5b61117182611178565b60ff60f81b8816815260e060208201525f61134f60e083018961111c565b8281036040840152611361818961111c565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b818110156113b6578351835260209384019390920191600101611398565b50909b9a5050505050505050505050565b5f5f5f5f5f5f5f60e0888a0312156113dd575f5ffd5b6113e688611178565b96506113f460208901611178565b955060408801359450606088013593506114106080890161125a565b9699959850939692959460a0840135945060c09093013592915050565b5f5f6040838503121561143e575f5ffd5b61144783611178565b915061145560208401611178565b90509250929050565b600181811c9082168061147257607f821691505b60208210810361149057634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156105d957805f5260205f20601f840160051c810160208510156114bb5750805b601f840160051c820191505b81811015610c18575f81556001016114c7565b815167ffffffffffffffff8111156114f4576114f46111bb565b61150881611502845461145e565b84611496565b6020601f82116001811461153a575f83156115235750848201515b5f19600385901b1c1916600184901b178455610c18565b5f84815260208120601f198516915b828110156115695787850151825560209485019460019092019101611549565b508482101561158657868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b808201808211156103cc57634e487b7160e01b5f52601160045260245ffd5b634e487b7160e01b5f52602160045260245ffdfe52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100a2646970667358221220f0e100f498b6f41fbf597025b81d217623a1e4b7d8ee5bf92572b87e86172a6264736f6c634300081c0033",
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

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20decimals

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

// Erc20decimalsMetaData contains all meta data concerning the Erc20decimals contract.
var Erc20decimalsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"initialAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712DOMAIN_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approveInternal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferInternal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052604051620018c7380380620018c783398101604081905262000026916200020c565b8484600362000036838262000341565b50600462000045828262000341565b5050506200005a83836200008560201b60201c565b845160209095019490942060075550506005805460ff191660ff90931692909217909155506200042f565b6001600160a01b038216620000e05760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640160405180910390fd5b8060025f828254620000f3919062000409565b90915550506001600160a01b0382165f81815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f83011262000172575f80fd5b81516001600160401b03808211156200018f576200018f6200014e565b604051601f8301601f19908116603f01168101908282118183101715620001ba57620001ba6200014e565b81604052838152602092508683858801011115620001d6575f80fd5b5f91505b83821015620001f95785820183015181830184015290820190620001da565b5f93810190920192909252949350505050565b5f805f805f60a0868803121562000221575f80fd5b85516001600160401b038082111562000238575f80fd5b6200024689838a0162000162565b965060208801519150808211156200025c575f80fd5b506200026b8882890162000162565b604088015190955090506001600160a01b03811681146200028a575f80fd5b60608701516080880151919450925060ff81168114620002a8575f80fd5b809150509295509295909350565b600181811c90821680620002cb57607f821691505b602082108103620002ea57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111562000149575f81815260208120601f850160051c81016020861015620003185750805b601f850160051c820191505b81811015620003395782815560010162000324565b505050505050565b81516001600160401b038111156200035d576200035d6200014e565b62000375816200036e8454620002b6565b84620002f0565b602080601f831160018114620003ab575f8415620003935750858301515b5f19600386901b1c1916600185901b17855562000339565b5f85815260208120601f198616915b82811015620003db57888601518255948401946001909101908401620003ba565b5085821015620003f957878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b808201808211156200042957634e487b7160e01b5f52601160045260245ffd5b92915050565b61148a806200043d5f395ff3fe608060405234801561000f575f80fd5b5060043610610184575f3560e01c806340c10f19116100dd5780639e4e731811610088578063c473af3311610063578063c473af33146103cf578063d505accf146103f6578063dd62ed3e14610409575f80fd5b80639e4e731814610382578063a457c2d7146103a9578063a9059cbb146103bc575f80fd5b806370a08231116100b857806370a08231146103265780637ecebe001461035b57806395d89b411461037a575f80fd5b806340c10f19146102ed57806342966c681461030057806356189cb414610313575f80fd5b806323b872dd1161013d5780633408e470116101185780633408e470146102485780633644e5151461024e57806339509351146102da575f80fd5b806323b872dd146101f957806330adf81f1461020c578063313ce56714610233575f80fd5b8063095ea7b31161016d578063095ea7b3146101b957806318160ddd146101dc578063222f5be0146101e4575f80fd5b806304622c2e1461018857806306fdde03146101a4575b5f80fd5b61019160075481565b6040519081526020015b60405180910390f35b6101ac61044e565b60405161019b91906111c5565b6101cc6101c7366004611256565b6104de565b604051901515815260200161019b565b600254610191565b6101f76101f236600461127e565b6104f7565b005b6101cc61020736600461127e565b610507565b6101917f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c981565b60055460405160ff909116815260200161019b565b46610191565b6101916007545f907f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f907fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc646604080516020810195909552840192909252606083015260808201523060a082015260c00160405160208183030381529060405280519060200120905090565b6101cc6102e8366004611256565b61052a565b6101f76102fb366004611256565b610575565b6101f761030e3660046112b7565b610583565b6101f761032136600461127e565b610590565b6101916103343660046112ce565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205490565b6101916103693660046112ce565b60066020525f908152604090205481565b6101ac61059b565b6101917fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc681565b6101cc6103b7366004611256565b6105aa565b6101cc6103ca366004611256565b61067f565b6101917f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f81565b6101f76104043660046112ee565b61068c565b61019161041736600461135b565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260016020908152604080832093909416825291909152205490565b60606003805461045d9061138c565b80601f01602080910402602001604051908101604052809291908181526020018280546104899061138c565b80156104d45780601f106104ab576101008083540402835291602001916104d4565b820191905f5260205f20905b8154815290600101906020018083116104b757829003601f168201915b5050505050905090565b5f336104eb8185856107d2565b60019150505b92915050565b610502838383610984565b505050565b5f33610514858285610bf3565b61051f858585610984565b506001949350505050565b335f81815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091906104eb908290869061057090879061140a565b6107d2565b61057f8282610cc3565b5050565b61058d3382610db4565b50565b6105028383836107d2565b60606004805461045d9061138c565b335f81815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919083811015610672576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f00000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b61051f82868684036107d2565b5f336104eb818585610984565b428410156106f6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f48455a3a3a7065726d69743a20415554485f45585049524544000000000000006044820152606401610669565b73ffffffffffffffffffffffffffffffffffffffff87165f90815260066020526040812080547f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9918a918a918a91908661074f8361141d565b9091555060408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090506107bd8882868686610f76565b6107c88888886107d2565b5050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8316610874576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff8216610917576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff8381165f8181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8316610a27576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff8216610aca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526020819052604090205481811015610b7f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff8481165f81815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35b50505050565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610bed5781811015610cb6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610669565b610bed84848484036107d2565b73ffffffffffffffffffffffffffffffffffffffff8216610d40576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610669565b8060025f828254610d51919061140a565b909155505073ffffffffffffffffffffffffffffffffffffffff82165f81815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b73ffffffffffffffffffffffffffffffffffffffff8216610e57576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f73000000000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526020819052604090205481811015610f0c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f63650000000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff83165f818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3505050565b600754604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f602080830191909152818301939093527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201524660808201523060a0808301919091528251808303909101815260c08201909252815191909201207f190100000000000000000000000000000000000000000000000000000000000060e083015260e282018190526101028201869052905f9061012201604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815282825280516020918201205f80855291840180845281905260ff89169284019290925260608301879052608083018690529092509060019060a0016020604051602081039080840390855afa1580156110be573d5f803e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81161580159061113957508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b6107c8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f48455a3a3a5f76616c69646174655369676e6564446174613a20494e56414c4960448201527f445f5349474e41545552450000000000000000000000000000000000000000006064820152608401610669565b5f6020808352835180828501525f5b818110156111f0578581018301518582016040015282016111d4565b505f6040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611251575f80fd5b919050565b5f8060408385031215611267575f80fd5b6112708361122e565b946020939093013593505050565b5f805f60608486031215611290575f80fd5b6112998461122e565b92506112a76020850161122e565b9150604084013590509250925092565b5f602082840312156112c7575f80fd5b5035919050565b5f602082840312156112de575f80fd5b6112e78261122e565b9392505050565b5f805f805f805f60e0888a031215611304575f80fd5b61130d8861122e565b965061131b6020890161122e565b95506040880135945060608801359350608088013560ff8116811461133e575f80fd5b9699959850939692959460a0840135945060c09093013592915050565b5f806040838503121561136c575f80fd5b6113758361122e565b91506113836020840161122e565b90509250929050565b600181811c908216806113a057607f821691505b6020821081036113d7577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b808201808211156104f1576104f16113dd565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361144d5761144d6113dd565b506001019056fea2646970667358221220e3a84b2fdb714d5930e07e9e0ed4e27a45873b4a8eb379915031a66640c192fb64736f6c63430008140033",
}

// Erc20decimalsABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20decimalsMetaData.ABI instead.
var Erc20decimalsABI = Erc20decimalsMetaData.ABI

// Erc20decimalsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Erc20decimalsMetaData.Bin instead.
var Erc20decimalsBin = Erc20decimalsMetaData.Bin

// DeployErc20decimals deploys a new Ethereum contract, binding an instance of Erc20decimals to it.
func DeployErc20decimals(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, initialAccount common.Address, initialBalance *big.Int, decimals_ uint8) (common.Address, *types.Transaction, *Erc20decimals, error) {
	parsed, err := Erc20decimalsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Erc20decimalsBin), backend, name, symbol, initialAccount, initialBalance, decimals_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc20decimals{Erc20decimalsCaller: Erc20decimalsCaller{contract: contract}, Erc20decimalsTransactor: Erc20decimalsTransactor{contract: contract}, Erc20decimalsFilterer: Erc20decimalsFilterer{contract: contract}}, nil
}

// Erc20decimals is an auto generated Go binding around an Ethereum contract.
type Erc20decimals struct {
	Erc20decimalsCaller     // Read-only binding to the contract
	Erc20decimalsTransactor // Write-only binding to the contract
	Erc20decimalsFilterer   // Log filterer for contract events
}

// Erc20decimalsCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20decimalsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20decimalsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20decimalsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20decimalsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20decimalsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20decimalsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20decimalsSession struct {
	Contract     *Erc20decimals    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20decimalsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20decimalsCallerSession struct {
	Contract *Erc20decimalsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// Erc20decimalsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20decimalsTransactorSession struct {
	Contract     *Erc20decimalsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// Erc20decimalsRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20decimalsRaw struct {
	Contract *Erc20decimals // Generic contract binding to access the raw methods on
}

// Erc20decimalsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20decimalsCallerRaw struct {
	Contract *Erc20decimalsCaller // Generic read-only contract binding to access the raw methods on
}

// Erc20decimalsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20decimalsTransactorRaw struct {
	Contract *Erc20decimalsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20decimals creates a new instance of Erc20decimals, bound to a specific deployed contract.
func NewErc20decimals(address common.Address, backend bind.ContractBackend) (*Erc20decimals, error) {
	contract, err := bindErc20decimals(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20decimals{Erc20decimalsCaller: Erc20decimalsCaller{contract: contract}, Erc20decimalsTransactor: Erc20decimalsTransactor{contract: contract}, Erc20decimalsFilterer: Erc20decimalsFilterer{contract: contract}}, nil
}

// NewErc20decimalsCaller creates a new read-only instance of Erc20decimals, bound to a specific deployed contract.
func NewErc20decimalsCaller(address common.Address, caller bind.ContractCaller) (*Erc20decimalsCaller, error) {
	contract, err := bindErc20decimals(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20decimalsCaller{contract: contract}, nil
}

// NewErc20decimalsTransactor creates a new write-only instance of Erc20decimals, bound to a specific deployed contract.
func NewErc20decimalsTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc20decimalsTransactor, error) {
	contract, err := bindErc20decimals(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20decimalsTransactor{contract: contract}, nil
}

// NewErc20decimalsFilterer creates a new log filterer instance of Erc20decimals, bound to a specific deployed contract.
func NewErc20decimalsFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc20decimalsFilterer, error) {
	contract, err := bindErc20decimals(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20decimalsFilterer{contract: contract}, nil
}

// bindErc20decimals binds a generic wrapper to an already deployed contract.
func bindErc20decimals(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc20decimalsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20decimals *Erc20decimalsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20decimals.Contract.Erc20decimalsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20decimals *Erc20decimalsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20decimals.Contract.Erc20decimalsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20decimals *Erc20decimalsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20decimals.Contract.Erc20decimalsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20decimals *Erc20decimalsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20decimals.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20decimals *Erc20decimalsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20decimals.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20decimals *Erc20decimalsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20decimals.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Erc20decimals *Erc20decimalsCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Erc20decimals.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Erc20decimals *Erc20decimalsSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Erc20decimals.Contract.DOMAINSEPARATOR(&_Erc20decimals.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Erc20decimals *Erc20decimalsCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Erc20decimals.Contract.DOMAINSEPARATOR(&_Erc20decimals.CallOpts)
}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xc473af33.
//
// Solidity: function EIP712DOMAIN_HASH() view returns(bytes32)
func (_Erc20decimals *Erc20decimalsCaller) EIP712DOMAINHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Erc20decimals.contract.Call(opts, &out, "EIP712DOMAIN_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xc473af33.
//
// Solidity: function EIP712DOMAIN_HASH() view returns(bytes32)
func (_Erc20decimals *Erc20decimalsSession) EIP712DOMAINHASH() ([32]byte, error) {
	return _Erc20decimals.Contract.EIP712DOMAINHASH(&_Erc20decimals.CallOpts)
}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xc473af33.
//
// Solidity: function EIP712DOMAIN_HASH() view returns(bytes32)
func (_Erc20decimals *Erc20decimalsCallerSession) EIP712DOMAINHASH() ([32]byte, error) {
	return _Erc20decimals.Contract.EIP712DOMAINHASH(&_Erc20decimals.CallOpts)
}

// NAMEHASH is a free data retrieval call binding the contract method 0x04622c2e.
//
// Solidity: function NAME_HASH() view returns(bytes32)
func (_Erc20decimals *Erc20decimalsCaller) NAMEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Erc20decimals.contract.Call(opts, &out, "NAME_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NAMEHASH is a free data retrieval call binding the contract method 0x04622c2e.
//
// Solidity: function NAME_HASH() view returns(bytes32)
func (_Erc20decimals *Erc20decimalsSession) NAMEHASH() ([32]byte, error) {
	return _Erc20decimals.Contract.NAMEHASH(&_Erc20decimals.CallOpts)
}

// NAMEHASH is a free data retrieval call binding the contract method 0x04622c2e.
//
// Solidity: function NAME_HASH() view returns(bytes32)
func (_Erc20decimals *Erc20decimalsCallerSession) NAMEHASH() ([32]byte, error) {
	return _Erc20decimals.Contract.NAMEHASH(&_Erc20decimals.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Erc20decimals *Erc20decimalsCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Erc20decimals.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Erc20decimals *Erc20decimalsSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Erc20decimals.Contract.PERMITTYPEHASH(&_Erc20decimals.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Erc20decimals *Erc20decimalsCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Erc20decimals.Contract.PERMITTYPEHASH(&_Erc20decimals.CallOpts)
}

// VERSIONHASH is a free data retrieval call binding the contract method 0x9e4e7318.
//
// Solidity: function VERSION_HASH() view returns(bytes32)
func (_Erc20decimals *Erc20decimalsCaller) VERSIONHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Erc20decimals.contract.Call(opts, &out, "VERSION_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VERSIONHASH is a free data retrieval call binding the contract method 0x9e4e7318.
//
// Solidity: function VERSION_HASH() view returns(bytes32)
func (_Erc20decimals *Erc20decimalsSession) VERSIONHASH() ([32]byte, error) {
	return _Erc20decimals.Contract.VERSIONHASH(&_Erc20decimals.CallOpts)
}

// VERSIONHASH is a free data retrieval call binding the contract method 0x9e4e7318.
//
// Solidity: function VERSION_HASH() view returns(bytes32)
func (_Erc20decimals *Erc20decimalsCallerSession) VERSIONHASH() ([32]byte, error) {
	return _Erc20decimals.Contract.VERSIONHASH(&_Erc20decimals.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20decimals *Erc20decimalsCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20decimals.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20decimals *Erc20decimalsSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc20decimals.Contract.Allowance(&_Erc20decimals.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20decimals *Erc20decimalsCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc20decimals.Contract.Allowance(&_Erc20decimals.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20decimals *Erc20decimalsCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20decimals.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20decimals *Erc20decimalsSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Erc20decimals.Contract.BalanceOf(&_Erc20decimals.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20decimals *Erc20decimalsCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Erc20decimals.Contract.BalanceOf(&_Erc20decimals.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20decimals *Erc20decimalsCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Erc20decimals.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20decimals *Erc20decimalsSession) Decimals() (uint8, error) {
	return _Erc20decimals.Contract.Decimals(&_Erc20decimals.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20decimals *Erc20decimalsCallerSession) Decimals() (uint8, error) {
	return _Erc20decimals.Contract.Decimals(&_Erc20decimals.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainId)
func (_Erc20decimals *Erc20decimalsCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20decimals.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainId)
func (_Erc20decimals *Erc20decimalsSession) GetChainId() (*big.Int, error) {
	return _Erc20decimals.Contract.GetChainId(&_Erc20decimals.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainId)
func (_Erc20decimals *Erc20decimalsCallerSession) GetChainId() (*big.Int, error) {
	return _Erc20decimals.Contract.GetChainId(&_Erc20decimals.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20decimals *Erc20decimalsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc20decimals.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20decimals *Erc20decimalsSession) Name() (string, error) {
	return _Erc20decimals.Contract.Name(&_Erc20decimals.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20decimals *Erc20decimalsCallerSession) Name() (string, error) {
	return _Erc20decimals.Contract.Name(&_Erc20decimals.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Erc20decimals *Erc20decimalsCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20decimals.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Erc20decimals *Erc20decimalsSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Erc20decimals.Contract.Nonces(&_Erc20decimals.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Erc20decimals *Erc20decimalsCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Erc20decimals.Contract.Nonces(&_Erc20decimals.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20decimals *Erc20decimalsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc20decimals.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20decimals *Erc20decimalsSession) Symbol() (string, error) {
	return _Erc20decimals.Contract.Symbol(&_Erc20decimals.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20decimals *Erc20decimalsCallerSession) Symbol() (string, error) {
	return _Erc20decimals.Contract.Symbol(&_Erc20decimals.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20decimals *Erc20decimalsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20decimals.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20decimals *Erc20decimalsSession) TotalSupply() (*big.Int, error) {
	return _Erc20decimals.Contract.TotalSupply(&_Erc20decimals.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20decimals *Erc20decimalsCallerSession) TotalSupply() (*big.Int, error) {
	return _Erc20decimals.Contract.TotalSupply(&_Erc20decimals.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20decimals *Erc20decimalsTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20decimals *Erc20decimalsSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.Contract.Approve(&_Erc20decimals.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20decimals *Erc20decimalsTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.Contract.Approve(&_Erc20decimals.TransactOpts, spender, amount)
}

// ApproveInternal is a paid mutator transaction binding the contract method 0x56189cb4.
//
// Solidity: function approveInternal(address owner, address spender, uint256 value) returns()
func (_Erc20decimals *Erc20decimalsTransactor) ApproveInternal(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.contract.Transact(opts, "approveInternal", owner, spender, value)
}

// ApproveInternal is a paid mutator transaction binding the contract method 0x56189cb4.
//
// Solidity: function approveInternal(address owner, address spender, uint256 value) returns()
func (_Erc20decimals *Erc20decimalsSession) ApproveInternal(owner common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.Contract.ApproveInternal(&_Erc20decimals.TransactOpts, owner, spender, value)
}

// ApproveInternal is a paid mutator transaction binding the contract method 0x56189cb4.
//
// Solidity: function approveInternal(address owner, address spender, uint256 value) returns()
func (_Erc20decimals *Erc20decimalsTransactorSession) ApproveInternal(owner common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.Contract.ApproveInternal(&_Erc20decimals.TransactOpts, owner, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Erc20decimals *Erc20decimalsTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Erc20decimals *Erc20decimalsSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.Contract.Burn(&_Erc20decimals.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Erc20decimals *Erc20decimalsTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.Contract.Burn(&_Erc20decimals.TransactOpts, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Erc20decimals *Erc20decimalsTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Erc20decimals *Erc20decimalsSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.Contract.DecreaseAllowance(&_Erc20decimals.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Erc20decimals *Erc20decimalsTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.Contract.DecreaseAllowance(&_Erc20decimals.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Erc20decimals *Erc20decimalsTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Erc20decimals *Erc20decimalsSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.Contract.IncreaseAllowance(&_Erc20decimals.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Erc20decimals *Erc20decimalsTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.Contract.IncreaseAllowance(&_Erc20decimals.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Erc20decimals *Erc20decimalsTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Erc20decimals *Erc20decimalsSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.Contract.Mint(&_Erc20decimals.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Erc20decimals *Erc20decimalsTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.Contract.Mint(&_Erc20decimals.TransactOpts, account, amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Erc20decimals *Erc20decimalsTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Erc20decimals.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Erc20decimals *Erc20decimalsSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Erc20decimals.Contract.Permit(&_Erc20decimals.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Erc20decimals *Erc20decimalsTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Erc20decimals.Contract.Permit(&_Erc20decimals.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Erc20decimals *Erc20decimalsTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Erc20decimals *Erc20decimalsSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.Contract.Transfer(&_Erc20decimals.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Erc20decimals *Erc20decimalsTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.Contract.Transfer(&_Erc20decimals.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Erc20decimals *Erc20decimalsTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Erc20decimals *Erc20decimalsSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.Contract.TransferFrom(&_Erc20decimals.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Erc20decimals *Erc20decimalsTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.Contract.TransferFrom(&_Erc20decimals.TransactOpts, from, to, amount)
}

// TransferInternal is a paid mutator transaction binding the contract method 0x222f5be0.
//
// Solidity: function transferInternal(address from, address to, uint256 value) returns()
func (_Erc20decimals *Erc20decimalsTransactor) TransferInternal(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.contract.Transact(opts, "transferInternal", from, to, value)
}

// TransferInternal is a paid mutator transaction binding the contract method 0x222f5be0.
//
// Solidity: function transferInternal(address from, address to, uint256 value) returns()
func (_Erc20decimals *Erc20decimalsSession) TransferInternal(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.Contract.TransferInternal(&_Erc20decimals.TransactOpts, from, to, value)
}

// TransferInternal is a paid mutator transaction binding the contract method 0x222f5be0.
//
// Solidity: function transferInternal(address from, address to, uint256 value) returns()
func (_Erc20decimals *Erc20decimalsTransactorSession) TransferInternal(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20decimals.Contract.TransferInternal(&_Erc20decimals.TransactOpts, from, to, value)
}

// Erc20decimalsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc20decimals contract.
type Erc20decimalsApprovalIterator struct {
	Event *Erc20decimalsApproval // Event containing the contract specifics and raw log

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
func (it *Erc20decimalsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20decimalsApproval)
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
		it.Event = new(Erc20decimalsApproval)
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
func (it *Erc20decimalsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20decimalsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20decimalsApproval represents a Approval event raised by the Erc20decimals contract.
type Erc20decimalsApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20decimals *Erc20decimalsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Erc20decimalsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20decimals.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Erc20decimalsApprovalIterator{contract: _Erc20decimals.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20decimals *Erc20decimalsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc20decimalsApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20decimals.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20decimalsApproval)
				if err := _Erc20decimals.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Erc20decimals *Erc20decimalsFilterer) ParseApproval(log types.Log) (*Erc20decimalsApproval, error) {
	event := new(Erc20decimalsApproval)
	if err := _Erc20decimals.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20decimalsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc20decimals contract.
type Erc20decimalsTransferIterator struct {
	Event *Erc20decimalsTransfer // Event containing the contract specifics and raw log

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
func (it *Erc20decimalsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20decimalsTransfer)
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
		it.Event = new(Erc20decimalsTransfer)
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
func (it *Erc20decimalsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20decimalsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20decimalsTransfer represents a Transfer event raised by the Erc20decimals contract.
type Erc20decimalsTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20decimals *Erc20decimalsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Erc20decimalsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20decimals.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc20decimalsTransferIterator{contract: _Erc20decimals.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20decimals *Erc20decimalsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc20decimalsTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20decimals.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20decimalsTransfer)
				if err := _Erc20decimals.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Erc20decimals *Erc20decimalsFilterer) ParseTransfer(log types.Log) (*Erc20decimalsTransfer, error) {
	event := new(Erc20decimalsTransfer)
	if err := _Erc20decimals.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

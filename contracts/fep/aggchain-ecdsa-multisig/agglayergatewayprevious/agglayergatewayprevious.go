// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package agglayergatewayprevious

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

// AgglayergatewaypreviousMetaData contains all meta data concerning the Agglayergatewayprevious contract.
var AgglayergatewaypreviousMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainVKeyAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProofBytesLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAggLayerAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAggLayerAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PPSelectorCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"RouteAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"RouteIsAlreadyFrozen\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"RouteIsFrozen\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"RouteNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VKeyCannotBeZero\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newVKey\",\"type\":\"bytes32\"}],\"name\":\"AddDefaultAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pessimisticVKey\",\"type\":\"bytes32\"}],\"name\":\"RouteAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pessimisticVKey\",\"type\":\"bytes32\"}],\"name\":\"RouteFrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"UnsetDefaultAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousVKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newVKey\",\"type\":\"bytes32\"}],\"name\":\"UpdateDefaultAggchainVKey\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AGGLAYER_GATEWAY_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"defaultAggchainSelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"addDefaultAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"pessimisticVKeySelector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"pessimisticVKey\",\"type\":\"bytes32\"}],\"name\":\"addPessimisticVKeyRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"defaultAggchainSelector\",\"type\":\"bytes4\"}],\"name\":\"defaultAggchainVKeys\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"defaultAggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"pessimisticVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"freezePessimisticVKeyRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"defaultAggchainSelector\",\"type\":\"bytes4\"}],\"name\":\"getDefaultAggchainVKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"aggchainDefaultVKeyRole\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addRouteRole\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"freezeRouteRole\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"pessimisticVKeySelector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"pessimisticVKey\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"pessimisticVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"pessimisticVKeyRoutes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"pessimisticVKey\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"defaultAggchainSelector\",\"type\":\"bytes4\"}],\"name\":\"unsetDefaultAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"defaultAggchainSelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"newDefaultAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"updateDefaultAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicValues\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofBytes\",\"type\":\"bytes\"}],\"name\":\"verifyPessimisticProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060156019565b60c9565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560685760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c65780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b611967806100d65f395ff3fe608060405234801561000f575f5ffd5b5060043610610149575f3560e01c8063818c8c21116100c7578063a217fddf1161007d578063c49196ec11610063578063c49196ec146103ba578063d547741f146103f6578063f4c4468114610409575f5ffd5b8063a217fddf146103a0578063a48fd377146103a7575f5ffd5b806391d14854116100ad57806391d14854146103165780639584a5161461037a57806395d2a3911461038d575f5ffd5b8063818c8c21146102e457806382bfaea1146102f7575f5ffd5b806347c63d361161011c5780636cabfdab116101025780636cabfdab14610241578063716be075146102545780637b552bea146102d1575f5ffd5b806347c63d36146101ec57806354fd4d50146101ff575f5ffd5b806301ffc9a71461014d578063248a9ca3146101755780632f2ff15d146101c457806336568abe146101d9575b5f5ffd5b61016061015b3660046115bf565b61041c565b60405190151581526020015b60405180910390f35b6101b66101833660046115df565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b60405190815260200161016c565b6101d76101d2366004611619565b6104b4565b005b6101d76101e7366004611619565b6104fd565b6101d76101fa3660046115bf565b61055b565b60408051808201909152600681527f76312e302e30000000000000000000000000000000000000000000000000000060208201525b60405161016c9190611643565b6101b661024f3660046115bf565b610655565b61029d6102623660046115bf565b600160208190525f918252604090912080549181015460029091015473ffffffffffffffffffffffffffffffffffffffff9092169160ff1683565b6040805173ffffffffffffffffffffffffffffffffffffffff9094168452602084019290925215159082015260600161016c565b6101d76102df366004611696565b6106f0565b6101d76102f2366004611713565b6109ae565b6101b66103053660046115bf565b5f6020819052908152604090205481565b610160610324366004611619565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b6101d761038836600461173b565b610ae7565b6101d761039b3660046115bf565b610b1c565b6101b65f81565b6101d76103b53660046117ba565b610cf9565b6102346040518060400160405280600681526020017f76312e302e30000000000000000000000000000000000000000000000000000081525081565b6101d7610404366004611619565b610f04565b6101d7610417366004611713565b610f47565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806104ae57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546104ed8161108d565b6104f7838361109a565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8116331461054c576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61055682826111bf565b505050565b7f131410eab1236cee2db19035b0e825c94e5ab705dffe23321dd53856da5316176105858161108d565b7fffffffff0000000000000000000000000000000000000000000000000000000082165f908152602081905260409020546105ec576040517f925e5a3a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f818152602081815260408083209290925590519182527f43b46d79ee2fbaa8a981b5fd78aaf6fb9e38b18d02e1af17564ce37c67078cab910160405180910390a15050565b7fffffffff0000000000000000000000000000000000000000000000000000000081165f908152602081905260408120546106bc576040517f925e5a3a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b507fffffffff00000000000000000000000000000000000000000000000000000000165f9081526020819052604090205490565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f8115801561073a5750825b90505f8267ffffffffffffffff1660011480156107565750303b155b905081158015610764575080155b1561079b576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156107fc5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b73ffffffffffffffffffffffffffffffffffffffff8c161580610833575073ffffffffffffffffffffffffffffffffffffffff8b16155b80610852575073ffffffffffffffffffffffffffffffffffffffff8a16155b80610871575073ffffffffffffffffffffffffffffffffffffffff8916155b156108a8576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108b25f8d61109a565b506108dd7f131410eab1236cee2db19035b0e825c94e5ab705dffe23321dd53856da5316178c61109a565b506109087f0fdc2a718b96bc741c7544001e3dd7c26730802c54781668fa78a120e622629b8b61109a565b506109337fca75ae4228cde6195f9fa3dbde8dc352fb30aa63780717a378ccfc50274355dd8a61109a565b5061093f88888861129b565b83156109a05784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050505050565b7f131410eab1236cee2db19035b0e825c94e5ab705dffe23321dd53856da5316176109d88161108d565b7fffffffff0000000000000000000000000000000000000000000000000000000083165f9081526020819052604090205415610a40576040517f22a1bdc400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81610a77576040517f6745305e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000083165f8181526020818152604091829020859055815192835282018490527f64b3528974574a41ddc830adc96440dccb99e64ede6b863194b951bfbd17e1bc91015b60405180910390a1505050565b7f0fdc2a718b96bc741c7544001e3dd7c26730802c54781668fa78a120e622629b610b118161108d565b6104f784848461129b565b7fca75ae4228cde6195f9fa3dbde8dc352fb30aa63780717a378ccfc50274355dd610b468161108d565b7fffffffff0000000000000000000000000000000000000000000000000000000082165f908152600160205260409020805473ffffffffffffffffffffffffffffffffffffffff16610bed576040517ff208777e0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000841660048201526024015b60405180910390fd5b600281015460ff1615610c50576040517fc7b4a1f20000000000000000000000000000000000000000000000000000000081527fffffffff0000000000000000000000000000000000000000000000000000000084166004820152602401610be4565b6002810180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001908117909155815490820154604080517fffffffff000000000000000000000000000000000000000000000000000000008716815273ffffffffffffffffffffffffffffffffffffffff90931660208401528201527f6efcfd4090b998a9d58c7f753bdda3291e9d781c02226b9f480b18ef07e0012b90606001610ada565b6004811015610d34576040517fae80851c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f610d426004828486611826565b610d4b9161184d565b7fffffffff0000000000000000000000000000000000000000000000000000000081165f908152600160208181526040928390208351606081018552815473ffffffffffffffffffffffffffffffffffffffff1680825293820154928101929092526002015460ff1615159281019290925291925090610e1b576040517ff208777e0000000000000000000000000000000000000000000000000000000081527fffffffff0000000000000000000000000000000000000000000000000000000083166004820152602401610be4565b806040015115610e7b576040517febf108230000000000000000000000000000000000000000000000000000000081527fffffffff0000000000000000000000000000000000000000000000000000000083166004820152602401610be4565b8051602082015173ffffffffffffffffffffffffffffffffffffffff909116906341493c60908888610eb0886004818c611826565b6040518663ffffffff1660e01b8152600401610ed09594939291906118f9565b5f6040518083038186803b158015610ee6575f5ffd5b505afa158015610ef8573d5f5f3e3d5ffd5b50505050505050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610f3d8161108d565b6104f783836111bf565b7f131410eab1236cee2db19035b0e825c94e5ab705dffe23321dd53856da531617610f718161108d565b7fffffffff0000000000000000000000000000000000000000000000000000000083165f90815260208190526040902054610fd8576040517f925e5a3a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8161100f576040517f6745305e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000083165f81815260208181526040918290208054908690558251938452908301819052908201849052907f477ac863279848383fdb5d6bb8957cc8f6f10c0200379e075df3f789d14248e6906060015b60405180910390a150505050565b61109781336114e1565b50565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff166111af575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905561114b3390565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506104ae565b5f9150506104ae565b5092915050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff16156111af575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506104ae565b73ffffffffffffffffffffffffffffffffffffffff82166112e8576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff000000000000000000000000000000000000000000000000000000008316611341576040517f7500f20900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80611378576040517f6745305e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000083165f908152600160205260409020805473ffffffffffffffffffffffffffffffffffffffff161561143a5780546040517f4bc7c2780000000000000000000000000000000000000000000000000000000081527fffffffff000000000000000000000000000000000000000000000000000000008616600482015273ffffffffffffffffffffffffffffffffffffffff9091166024820152604401610be4565b80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8416908117825560018201839055604080517fffffffff0000000000000000000000000000000000000000000000000000000087168152602081019290925281018390527f4c4455a30528c319643ae6fc35f3a0fcabde6c015f2a1a5270382c4a190b0fa39060600161107f565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16611587576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610be4565b5050565b80357fffffffff00000000000000000000000000000000000000000000000000000000811681146115ba575f5ffd5b919050565b5f602082840312156115cf575f5ffd5b6115d88261158b565b9392505050565b5f602082840312156115ef575f5ffd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146115ba575f5ffd5b5f5f6040838503121561162a575f5ffd5b8235915061163a602084016115f6565b90509250929050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b5f5f5f5f5f5f5f60e0888a0312156116ac575f5ffd5b6116b5886115f6565b96506116c3602089016115f6565b95506116d1604089016115f6565b94506116df606089016115f6565b93506116ed6080890161158b565b92506116fb60a089016115f6565b96999598509396929591949193505060c09091013590565b5f5f60408385031215611724575f5ffd5b61172d8361158b565b946020939093013593505050565b5f5f5f6060848603121561174d575f5ffd5b6117568461158b565b9250611764602085016115f6565b929592945050506040919091013590565b5f5f83601f840112611785575f5ffd5b50813567ffffffffffffffff81111561179c575f5ffd5b6020830191508360208285010111156117b3575f5ffd5b9250929050565b5f5f5f5f604085870312156117cd575f5ffd5b843567ffffffffffffffff8111156117e3575f5ffd5b6117ef87828801611775565b909550935050602085013567ffffffffffffffff81111561180e575f5ffd5b61181a87828801611775565b95989497509550505050565b5f5f85851115611834575f5ffd5b83861115611840575f5ffd5b5050820193919092039150565b80357fffffffff0000000000000000000000000000000000000000000000000000000081169060048410156111b8577fffffffff00000000000000000000000000000000000000000000000000000000808560040360031b1b82161691505092915050565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b858152606060208201525f6119126060830186886118b2565b82810360408401526119258185876118b2565b9897505050505050505056fea2646970667358221220dab8622e091836a1a767e41ccb40e7f9872daa1aef8d80d1b25c4c6a6d26d51764736f6c634300081c0033",
}

// AgglayergatewaypreviousABI is the input ABI used to generate the binding from.
// Deprecated: Use AgglayergatewaypreviousMetaData.ABI instead.
var AgglayergatewaypreviousABI = AgglayergatewaypreviousMetaData.ABI

// AgglayergatewaypreviousBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AgglayergatewaypreviousMetaData.Bin instead.
var AgglayergatewaypreviousBin = AgglayergatewaypreviousMetaData.Bin

// DeployAgglayergatewayprevious deploys a new Ethereum contract, binding an instance of Agglayergatewayprevious to it.
func DeployAgglayergatewayprevious(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Agglayergatewayprevious, error) {
	parsed, err := AgglayergatewaypreviousMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AgglayergatewaypreviousBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Agglayergatewayprevious{AgglayergatewaypreviousCaller: AgglayergatewaypreviousCaller{contract: contract}, AgglayergatewaypreviousTransactor: AgglayergatewaypreviousTransactor{contract: contract}, AgglayergatewaypreviousFilterer: AgglayergatewaypreviousFilterer{contract: contract}}, nil
}

// Agglayergatewayprevious is an auto generated Go binding around an Ethereum contract.
type Agglayergatewayprevious struct {
	AgglayergatewaypreviousCaller     // Read-only binding to the contract
	AgglayergatewaypreviousTransactor // Write-only binding to the contract
	AgglayergatewaypreviousFilterer   // Log filterer for contract events
}

// AgglayergatewaypreviousCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgglayergatewaypreviousCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgglayergatewaypreviousTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgglayergatewaypreviousTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgglayergatewaypreviousFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgglayergatewaypreviousFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgglayergatewaypreviousSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgglayergatewaypreviousSession struct {
	Contract     *Agglayergatewayprevious // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AgglayergatewaypreviousCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgglayergatewaypreviousCallerSession struct {
	Contract *AgglayergatewaypreviousCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// AgglayergatewaypreviousTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgglayergatewaypreviousTransactorSession struct {
	Contract     *AgglayergatewaypreviousTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// AgglayergatewaypreviousRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgglayergatewaypreviousRaw struct {
	Contract *Agglayergatewayprevious // Generic contract binding to access the raw methods on
}

// AgglayergatewaypreviousCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgglayergatewaypreviousCallerRaw struct {
	Contract *AgglayergatewaypreviousCaller // Generic read-only contract binding to access the raw methods on
}

// AgglayergatewaypreviousTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgglayergatewaypreviousTransactorRaw struct {
	Contract *AgglayergatewaypreviousTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgglayergatewayprevious creates a new instance of Agglayergatewayprevious, bound to a specific deployed contract.
func NewAgglayergatewayprevious(address common.Address, backend bind.ContractBackend) (*Agglayergatewayprevious, error) {
	contract, err := bindAgglayergatewayprevious(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Agglayergatewayprevious{AgglayergatewaypreviousCaller: AgglayergatewaypreviousCaller{contract: contract}, AgglayergatewaypreviousTransactor: AgglayergatewaypreviousTransactor{contract: contract}, AgglayergatewaypreviousFilterer: AgglayergatewaypreviousFilterer{contract: contract}}, nil
}

// NewAgglayergatewaypreviousCaller creates a new read-only instance of Agglayergatewayprevious, bound to a specific deployed contract.
func NewAgglayergatewaypreviousCaller(address common.Address, caller bind.ContractCaller) (*AgglayergatewaypreviousCaller, error) {
	contract, err := bindAgglayergatewayprevious(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgglayergatewaypreviousCaller{contract: contract}, nil
}

// NewAgglayergatewaypreviousTransactor creates a new write-only instance of Agglayergatewayprevious, bound to a specific deployed contract.
func NewAgglayergatewaypreviousTransactor(address common.Address, transactor bind.ContractTransactor) (*AgglayergatewaypreviousTransactor, error) {
	contract, err := bindAgglayergatewayprevious(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgglayergatewaypreviousTransactor{contract: contract}, nil
}

// NewAgglayergatewaypreviousFilterer creates a new log filterer instance of Agglayergatewayprevious, bound to a specific deployed contract.
func NewAgglayergatewaypreviousFilterer(address common.Address, filterer bind.ContractFilterer) (*AgglayergatewaypreviousFilterer, error) {
	contract, err := bindAgglayergatewayprevious(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgglayergatewaypreviousFilterer{contract: contract}, nil
}

// bindAgglayergatewayprevious binds a generic wrapper to an already deployed contract.
func bindAgglayergatewayprevious(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AgglayergatewaypreviousMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Agglayergatewayprevious *AgglayergatewaypreviousRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Agglayergatewayprevious.Contract.AgglayergatewaypreviousCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Agglayergatewayprevious *AgglayergatewaypreviousRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayergatewayprevious.Contract.AgglayergatewaypreviousTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Agglayergatewayprevious *AgglayergatewaypreviousRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Agglayergatewayprevious.Contract.AgglayergatewaypreviousTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Agglayergatewayprevious *AgglayergatewaypreviousCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Agglayergatewayprevious.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Agglayergatewayprevious *AgglayergatewaypreviousTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayergatewayprevious.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Agglayergatewayprevious *AgglayergatewaypreviousTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Agglayergatewayprevious.Contract.contract.Transact(opts, method, params...)
}

// AGGLAYERGATEWAYVERSION is a free data retrieval call binding the contract method 0xc49196ec.
//
// Solidity: function AGGLAYER_GATEWAY_VERSION() view returns(string)
func (_Agglayergatewayprevious *AgglayergatewaypreviousCaller) AGGLAYERGATEWAYVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Agglayergatewayprevious.contract.Call(opts, &out, "AGGLAYER_GATEWAY_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AGGLAYERGATEWAYVERSION is a free data retrieval call binding the contract method 0xc49196ec.
//
// Solidity: function AGGLAYER_GATEWAY_VERSION() view returns(string)
func (_Agglayergatewayprevious *AgglayergatewaypreviousSession) AGGLAYERGATEWAYVERSION() (string, error) {
	return _Agglayergatewayprevious.Contract.AGGLAYERGATEWAYVERSION(&_Agglayergatewayprevious.CallOpts)
}

// AGGLAYERGATEWAYVERSION is a free data retrieval call binding the contract method 0xc49196ec.
//
// Solidity: function AGGLAYER_GATEWAY_VERSION() view returns(string)
func (_Agglayergatewayprevious *AgglayergatewaypreviousCallerSession) AGGLAYERGATEWAYVERSION() (string, error) {
	return _Agglayergatewayprevious.Contract.AGGLAYERGATEWAYVERSION(&_Agglayergatewayprevious.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Agglayergatewayprevious *AgglayergatewaypreviousCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Agglayergatewayprevious.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Agglayergatewayprevious *AgglayergatewaypreviousSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Agglayergatewayprevious.Contract.DEFAULTADMINROLE(&_Agglayergatewayprevious.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Agglayergatewayprevious *AgglayergatewaypreviousCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Agglayergatewayprevious.Contract.DEFAULTADMINROLE(&_Agglayergatewayprevious.CallOpts)
}

// DefaultAggchainVKeys is a free data retrieval call binding the contract method 0x82bfaea1.
//
// Solidity: function defaultAggchainVKeys(bytes4 defaultAggchainSelector) view returns(bytes32 defaultAggchainVKey)
func (_Agglayergatewayprevious *AgglayergatewaypreviousCaller) DefaultAggchainVKeys(opts *bind.CallOpts, defaultAggchainSelector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _Agglayergatewayprevious.contract.Call(opts, &out, "defaultAggchainVKeys", defaultAggchainSelector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DefaultAggchainVKeys is a free data retrieval call binding the contract method 0x82bfaea1.
//
// Solidity: function defaultAggchainVKeys(bytes4 defaultAggchainSelector) view returns(bytes32 defaultAggchainVKey)
func (_Agglayergatewayprevious *AgglayergatewaypreviousSession) DefaultAggchainVKeys(defaultAggchainSelector [4]byte) ([32]byte, error) {
	return _Agglayergatewayprevious.Contract.DefaultAggchainVKeys(&_Agglayergatewayprevious.CallOpts, defaultAggchainSelector)
}

// DefaultAggchainVKeys is a free data retrieval call binding the contract method 0x82bfaea1.
//
// Solidity: function defaultAggchainVKeys(bytes4 defaultAggchainSelector) view returns(bytes32 defaultAggchainVKey)
func (_Agglayergatewayprevious *AgglayergatewaypreviousCallerSession) DefaultAggchainVKeys(defaultAggchainSelector [4]byte) ([32]byte, error) {
	return _Agglayergatewayprevious.Contract.DefaultAggchainVKeys(&_Agglayergatewayprevious.CallOpts, defaultAggchainSelector)
}

// GetDefaultAggchainVKey is a free data retrieval call binding the contract method 0x6cabfdab.
//
// Solidity: function getDefaultAggchainVKey(bytes4 defaultAggchainSelector) view returns(bytes32)
func (_Agglayergatewayprevious *AgglayergatewaypreviousCaller) GetDefaultAggchainVKey(opts *bind.CallOpts, defaultAggchainSelector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _Agglayergatewayprevious.contract.Call(opts, &out, "getDefaultAggchainVKey", defaultAggchainSelector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDefaultAggchainVKey is a free data retrieval call binding the contract method 0x6cabfdab.
//
// Solidity: function getDefaultAggchainVKey(bytes4 defaultAggchainSelector) view returns(bytes32)
func (_Agglayergatewayprevious *AgglayergatewaypreviousSession) GetDefaultAggchainVKey(defaultAggchainSelector [4]byte) ([32]byte, error) {
	return _Agglayergatewayprevious.Contract.GetDefaultAggchainVKey(&_Agglayergatewayprevious.CallOpts, defaultAggchainSelector)
}

// GetDefaultAggchainVKey is a free data retrieval call binding the contract method 0x6cabfdab.
//
// Solidity: function getDefaultAggchainVKey(bytes4 defaultAggchainSelector) view returns(bytes32)
func (_Agglayergatewayprevious *AgglayergatewaypreviousCallerSession) GetDefaultAggchainVKey(defaultAggchainSelector [4]byte) ([32]byte, error) {
	return _Agglayergatewayprevious.Contract.GetDefaultAggchainVKey(&_Agglayergatewayprevious.CallOpts, defaultAggchainSelector)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Agglayergatewayprevious *AgglayergatewaypreviousCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Agglayergatewayprevious.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Agglayergatewayprevious *AgglayergatewaypreviousSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Agglayergatewayprevious.Contract.GetRoleAdmin(&_Agglayergatewayprevious.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Agglayergatewayprevious *AgglayergatewaypreviousCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Agglayergatewayprevious.Contract.GetRoleAdmin(&_Agglayergatewayprevious.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Agglayergatewayprevious *AgglayergatewaypreviousCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Agglayergatewayprevious.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Agglayergatewayprevious *AgglayergatewaypreviousSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Agglayergatewayprevious.Contract.HasRole(&_Agglayergatewayprevious.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Agglayergatewayprevious *AgglayergatewaypreviousCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Agglayergatewayprevious.Contract.HasRole(&_Agglayergatewayprevious.CallOpts, role, account)
}

// PessimisticVKeyRoutes is a free data retrieval call binding the contract method 0x716be075.
//
// Solidity: function pessimisticVKeyRoutes(bytes4 pessimisticVKeySelector) view returns(address verifier, bytes32 pessimisticVKey, bool frozen)
func (_Agglayergatewayprevious *AgglayergatewaypreviousCaller) PessimisticVKeyRoutes(opts *bind.CallOpts, pessimisticVKeySelector [4]byte) (struct {
	Verifier        common.Address
	PessimisticVKey [32]byte
	Frozen          bool
}, error) {
	var out []interface{}
	err := _Agglayergatewayprevious.contract.Call(opts, &out, "pessimisticVKeyRoutes", pessimisticVKeySelector)

	outstruct := new(struct {
		Verifier        common.Address
		PessimisticVKey [32]byte
		Frozen          bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Verifier = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PessimisticVKey = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Frozen = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// PessimisticVKeyRoutes is a free data retrieval call binding the contract method 0x716be075.
//
// Solidity: function pessimisticVKeyRoutes(bytes4 pessimisticVKeySelector) view returns(address verifier, bytes32 pessimisticVKey, bool frozen)
func (_Agglayergatewayprevious *AgglayergatewaypreviousSession) PessimisticVKeyRoutes(pessimisticVKeySelector [4]byte) (struct {
	Verifier        common.Address
	PessimisticVKey [32]byte
	Frozen          bool
}, error) {
	return _Agglayergatewayprevious.Contract.PessimisticVKeyRoutes(&_Agglayergatewayprevious.CallOpts, pessimisticVKeySelector)
}

// PessimisticVKeyRoutes is a free data retrieval call binding the contract method 0x716be075.
//
// Solidity: function pessimisticVKeyRoutes(bytes4 pessimisticVKeySelector) view returns(address verifier, bytes32 pessimisticVKey, bool frozen)
func (_Agglayergatewayprevious *AgglayergatewaypreviousCallerSession) PessimisticVKeyRoutes(pessimisticVKeySelector [4]byte) (struct {
	Verifier        common.Address
	PessimisticVKey [32]byte
	Frozen          bool
}, error) {
	return _Agglayergatewayprevious.Contract.PessimisticVKeyRoutes(&_Agglayergatewayprevious.CallOpts, pessimisticVKeySelector)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Agglayergatewayprevious *AgglayergatewaypreviousCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Agglayergatewayprevious.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Agglayergatewayprevious *AgglayergatewaypreviousSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Agglayergatewayprevious.Contract.SupportsInterface(&_Agglayergatewayprevious.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Agglayergatewayprevious *AgglayergatewaypreviousCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Agglayergatewayprevious.Contract.SupportsInterface(&_Agglayergatewayprevious.CallOpts, interfaceId)
}

// VerifyPessimisticProof is a free data retrieval call binding the contract method 0xa48fd377.
//
// Solidity: function verifyPessimisticProof(bytes publicValues, bytes proofBytes) view returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousCaller) VerifyPessimisticProof(opts *bind.CallOpts, publicValues []byte, proofBytes []byte) error {
	var out []interface{}
	err := _Agglayergatewayprevious.contract.Call(opts, &out, "verifyPessimisticProof", publicValues, proofBytes)

	if err != nil {
		return err
	}

	return err

}

// VerifyPessimisticProof is a free data retrieval call binding the contract method 0xa48fd377.
//
// Solidity: function verifyPessimisticProof(bytes publicValues, bytes proofBytes) view returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousSession) VerifyPessimisticProof(publicValues []byte, proofBytes []byte) error {
	return _Agglayergatewayprevious.Contract.VerifyPessimisticProof(&_Agglayergatewayprevious.CallOpts, publicValues, proofBytes)
}

// VerifyPessimisticProof is a free data retrieval call binding the contract method 0xa48fd377.
//
// Solidity: function verifyPessimisticProof(bytes publicValues, bytes proofBytes) view returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousCallerSession) VerifyPessimisticProof(publicValues []byte, proofBytes []byte) error {
	return _Agglayergatewayprevious.Contract.VerifyPessimisticProof(&_Agglayergatewayprevious.CallOpts, publicValues, proofBytes)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Agglayergatewayprevious *AgglayergatewaypreviousCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Agglayergatewayprevious.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Agglayergatewayprevious *AgglayergatewaypreviousSession) Version() (string, error) {
	return _Agglayergatewayprevious.Contract.Version(&_Agglayergatewayprevious.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Agglayergatewayprevious *AgglayergatewaypreviousCallerSession) Version() (string, error) {
	return _Agglayergatewayprevious.Contract.Version(&_Agglayergatewayprevious.CallOpts)
}

// AddDefaultAggchainVKey is a paid mutator transaction binding the contract method 0x818c8c21.
//
// Solidity: function addDefaultAggchainVKey(bytes4 defaultAggchainSelector, bytes32 newAggchainVKey) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousTransactor) AddDefaultAggchainVKey(opts *bind.TransactOpts, defaultAggchainSelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Agglayergatewayprevious.contract.Transact(opts, "addDefaultAggchainVKey", defaultAggchainSelector, newAggchainVKey)
}

// AddDefaultAggchainVKey is a paid mutator transaction binding the contract method 0x818c8c21.
//
// Solidity: function addDefaultAggchainVKey(bytes4 defaultAggchainSelector, bytes32 newAggchainVKey) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousSession) AddDefaultAggchainVKey(defaultAggchainSelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Agglayergatewayprevious.Contract.AddDefaultAggchainVKey(&_Agglayergatewayprevious.TransactOpts, defaultAggchainSelector, newAggchainVKey)
}

// AddDefaultAggchainVKey is a paid mutator transaction binding the contract method 0x818c8c21.
//
// Solidity: function addDefaultAggchainVKey(bytes4 defaultAggchainSelector, bytes32 newAggchainVKey) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousTransactorSession) AddDefaultAggchainVKey(defaultAggchainSelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Agglayergatewayprevious.Contract.AddDefaultAggchainVKey(&_Agglayergatewayprevious.TransactOpts, defaultAggchainSelector, newAggchainVKey)
}

// AddPessimisticVKeyRoute is a paid mutator transaction binding the contract method 0x9584a516.
//
// Solidity: function addPessimisticVKeyRoute(bytes4 pessimisticVKeySelector, address verifier, bytes32 pessimisticVKey) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousTransactor) AddPessimisticVKeyRoute(opts *bind.TransactOpts, pessimisticVKeySelector [4]byte, verifier common.Address, pessimisticVKey [32]byte) (*types.Transaction, error) {
	return _Agglayergatewayprevious.contract.Transact(opts, "addPessimisticVKeyRoute", pessimisticVKeySelector, verifier, pessimisticVKey)
}

// AddPessimisticVKeyRoute is a paid mutator transaction binding the contract method 0x9584a516.
//
// Solidity: function addPessimisticVKeyRoute(bytes4 pessimisticVKeySelector, address verifier, bytes32 pessimisticVKey) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousSession) AddPessimisticVKeyRoute(pessimisticVKeySelector [4]byte, verifier common.Address, pessimisticVKey [32]byte) (*types.Transaction, error) {
	return _Agglayergatewayprevious.Contract.AddPessimisticVKeyRoute(&_Agglayergatewayprevious.TransactOpts, pessimisticVKeySelector, verifier, pessimisticVKey)
}

// AddPessimisticVKeyRoute is a paid mutator transaction binding the contract method 0x9584a516.
//
// Solidity: function addPessimisticVKeyRoute(bytes4 pessimisticVKeySelector, address verifier, bytes32 pessimisticVKey) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousTransactorSession) AddPessimisticVKeyRoute(pessimisticVKeySelector [4]byte, verifier common.Address, pessimisticVKey [32]byte) (*types.Transaction, error) {
	return _Agglayergatewayprevious.Contract.AddPessimisticVKeyRoute(&_Agglayergatewayprevious.TransactOpts, pessimisticVKeySelector, verifier, pessimisticVKey)
}

// FreezePessimisticVKeyRoute is a paid mutator transaction binding the contract method 0x95d2a391.
//
// Solidity: function freezePessimisticVKeyRoute(bytes4 pessimisticVKeySelector) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousTransactor) FreezePessimisticVKeyRoute(opts *bind.TransactOpts, pessimisticVKeySelector [4]byte) (*types.Transaction, error) {
	return _Agglayergatewayprevious.contract.Transact(opts, "freezePessimisticVKeyRoute", pessimisticVKeySelector)
}

// FreezePessimisticVKeyRoute is a paid mutator transaction binding the contract method 0x95d2a391.
//
// Solidity: function freezePessimisticVKeyRoute(bytes4 pessimisticVKeySelector) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousSession) FreezePessimisticVKeyRoute(pessimisticVKeySelector [4]byte) (*types.Transaction, error) {
	return _Agglayergatewayprevious.Contract.FreezePessimisticVKeyRoute(&_Agglayergatewayprevious.TransactOpts, pessimisticVKeySelector)
}

// FreezePessimisticVKeyRoute is a paid mutator transaction binding the contract method 0x95d2a391.
//
// Solidity: function freezePessimisticVKeyRoute(bytes4 pessimisticVKeySelector) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousTransactorSession) FreezePessimisticVKeyRoute(pessimisticVKeySelector [4]byte) (*types.Transaction, error) {
	return _Agglayergatewayprevious.Contract.FreezePessimisticVKeyRoute(&_Agglayergatewayprevious.TransactOpts, pessimisticVKeySelector)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayergatewayprevious.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayergatewayprevious.Contract.GrantRole(&_Agglayergatewayprevious.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayergatewayprevious.Contract.GrantRole(&_Agglayergatewayprevious.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x7b552bea.
//
// Solidity: function initialize(address defaultAdmin, address aggchainDefaultVKeyRole, address addRouteRole, address freezeRouteRole, bytes4 pessimisticVKeySelector, address verifier, bytes32 pessimisticVKey) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousTransactor) Initialize(opts *bind.TransactOpts, defaultAdmin common.Address, aggchainDefaultVKeyRole common.Address, addRouteRole common.Address, freezeRouteRole common.Address, pessimisticVKeySelector [4]byte, verifier common.Address, pessimisticVKey [32]byte) (*types.Transaction, error) {
	return _Agglayergatewayprevious.contract.Transact(opts, "initialize", defaultAdmin, aggchainDefaultVKeyRole, addRouteRole, freezeRouteRole, pessimisticVKeySelector, verifier, pessimisticVKey)
}

// Initialize is a paid mutator transaction binding the contract method 0x7b552bea.
//
// Solidity: function initialize(address defaultAdmin, address aggchainDefaultVKeyRole, address addRouteRole, address freezeRouteRole, bytes4 pessimisticVKeySelector, address verifier, bytes32 pessimisticVKey) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousSession) Initialize(defaultAdmin common.Address, aggchainDefaultVKeyRole common.Address, addRouteRole common.Address, freezeRouteRole common.Address, pessimisticVKeySelector [4]byte, verifier common.Address, pessimisticVKey [32]byte) (*types.Transaction, error) {
	return _Agglayergatewayprevious.Contract.Initialize(&_Agglayergatewayprevious.TransactOpts, defaultAdmin, aggchainDefaultVKeyRole, addRouteRole, freezeRouteRole, pessimisticVKeySelector, verifier, pessimisticVKey)
}

// Initialize is a paid mutator transaction binding the contract method 0x7b552bea.
//
// Solidity: function initialize(address defaultAdmin, address aggchainDefaultVKeyRole, address addRouteRole, address freezeRouteRole, bytes4 pessimisticVKeySelector, address verifier, bytes32 pessimisticVKey) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousTransactorSession) Initialize(defaultAdmin common.Address, aggchainDefaultVKeyRole common.Address, addRouteRole common.Address, freezeRouteRole common.Address, pessimisticVKeySelector [4]byte, verifier common.Address, pessimisticVKey [32]byte) (*types.Transaction, error) {
	return _Agglayergatewayprevious.Contract.Initialize(&_Agglayergatewayprevious.TransactOpts, defaultAdmin, aggchainDefaultVKeyRole, addRouteRole, freezeRouteRole, pessimisticVKeySelector, verifier, pessimisticVKey)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Agglayergatewayprevious.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Agglayergatewayprevious.Contract.RenounceRole(&_Agglayergatewayprevious.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Agglayergatewayprevious.Contract.RenounceRole(&_Agglayergatewayprevious.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayergatewayprevious.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayergatewayprevious.Contract.RevokeRole(&_Agglayergatewayprevious.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayergatewayprevious.Contract.RevokeRole(&_Agglayergatewayprevious.TransactOpts, role, account)
}

// UnsetDefaultAggchainVKey is a paid mutator transaction binding the contract method 0x47c63d36.
//
// Solidity: function unsetDefaultAggchainVKey(bytes4 defaultAggchainSelector) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousTransactor) UnsetDefaultAggchainVKey(opts *bind.TransactOpts, defaultAggchainSelector [4]byte) (*types.Transaction, error) {
	return _Agglayergatewayprevious.contract.Transact(opts, "unsetDefaultAggchainVKey", defaultAggchainSelector)
}

// UnsetDefaultAggchainVKey is a paid mutator transaction binding the contract method 0x47c63d36.
//
// Solidity: function unsetDefaultAggchainVKey(bytes4 defaultAggchainSelector) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousSession) UnsetDefaultAggchainVKey(defaultAggchainSelector [4]byte) (*types.Transaction, error) {
	return _Agglayergatewayprevious.Contract.UnsetDefaultAggchainVKey(&_Agglayergatewayprevious.TransactOpts, defaultAggchainSelector)
}

// UnsetDefaultAggchainVKey is a paid mutator transaction binding the contract method 0x47c63d36.
//
// Solidity: function unsetDefaultAggchainVKey(bytes4 defaultAggchainSelector) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousTransactorSession) UnsetDefaultAggchainVKey(defaultAggchainSelector [4]byte) (*types.Transaction, error) {
	return _Agglayergatewayprevious.Contract.UnsetDefaultAggchainVKey(&_Agglayergatewayprevious.TransactOpts, defaultAggchainSelector)
}

// UpdateDefaultAggchainVKey is a paid mutator transaction binding the contract method 0xf4c44681.
//
// Solidity: function updateDefaultAggchainVKey(bytes4 defaultAggchainSelector, bytes32 newDefaultAggchainVKey) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousTransactor) UpdateDefaultAggchainVKey(opts *bind.TransactOpts, defaultAggchainSelector [4]byte, newDefaultAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Agglayergatewayprevious.contract.Transact(opts, "updateDefaultAggchainVKey", defaultAggchainSelector, newDefaultAggchainVKey)
}

// UpdateDefaultAggchainVKey is a paid mutator transaction binding the contract method 0xf4c44681.
//
// Solidity: function updateDefaultAggchainVKey(bytes4 defaultAggchainSelector, bytes32 newDefaultAggchainVKey) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousSession) UpdateDefaultAggchainVKey(defaultAggchainSelector [4]byte, newDefaultAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Agglayergatewayprevious.Contract.UpdateDefaultAggchainVKey(&_Agglayergatewayprevious.TransactOpts, defaultAggchainSelector, newDefaultAggchainVKey)
}

// UpdateDefaultAggchainVKey is a paid mutator transaction binding the contract method 0xf4c44681.
//
// Solidity: function updateDefaultAggchainVKey(bytes4 defaultAggchainSelector, bytes32 newDefaultAggchainVKey) returns()
func (_Agglayergatewayprevious *AgglayergatewaypreviousTransactorSession) UpdateDefaultAggchainVKey(defaultAggchainSelector [4]byte, newDefaultAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Agglayergatewayprevious.Contract.UpdateDefaultAggchainVKey(&_Agglayergatewayprevious.TransactOpts, defaultAggchainSelector, newDefaultAggchainVKey)
}

// AgglayergatewaypreviousAddDefaultAggchainVKeyIterator is returned from FilterAddDefaultAggchainVKey and is used to iterate over the raw logs and unpacked data for AddDefaultAggchainVKey events raised by the Agglayergatewayprevious contract.
type AgglayergatewaypreviousAddDefaultAggchainVKeyIterator struct {
	Event *AgglayergatewaypreviousAddDefaultAggchainVKey // Event containing the contract specifics and raw log

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
func (it *AgglayergatewaypreviousAddDefaultAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayergatewaypreviousAddDefaultAggchainVKey)
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
		it.Event = new(AgglayergatewaypreviousAddDefaultAggchainVKey)
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
func (it *AgglayergatewaypreviousAddDefaultAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayergatewaypreviousAddDefaultAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayergatewaypreviousAddDefaultAggchainVKey represents a AddDefaultAggchainVKey event raised by the Agglayergatewayprevious contract.
type AgglayergatewaypreviousAddDefaultAggchainVKey struct {
	Selector [4]byte
	NewVKey  [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddDefaultAggchainVKey is a free log retrieval operation binding the contract event 0x64b3528974574a41ddc830adc96440dccb99e64ede6b863194b951bfbd17e1bc.
//
// Solidity: event AddDefaultAggchainVKey(bytes4 selector, bytes32 newVKey)
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) FilterAddDefaultAggchainVKey(opts *bind.FilterOpts) (*AgglayergatewaypreviousAddDefaultAggchainVKeyIterator, error) {

	logs, sub, err := _Agglayergatewayprevious.contract.FilterLogs(opts, "AddDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &AgglayergatewaypreviousAddDefaultAggchainVKeyIterator{contract: _Agglayergatewayprevious.contract, event: "AddDefaultAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchAddDefaultAggchainVKey is a free log subscription operation binding the contract event 0x64b3528974574a41ddc830adc96440dccb99e64ede6b863194b951bfbd17e1bc.
//
// Solidity: event AddDefaultAggchainVKey(bytes4 selector, bytes32 newVKey)
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) WatchAddDefaultAggchainVKey(opts *bind.WatchOpts, sink chan<- *AgglayergatewaypreviousAddDefaultAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Agglayergatewayprevious.contract.WatchLogs(opts, "AddDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayergatewaypreviousAddDefaultAggchainVKey)
				if err := _Agglayergatewayprevious.contract.UnpackLog(event, "AddDefaultAggchainVKey", log); err != nil {
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

// ParseAddDefaultAggchainVKey is a log parse operation binding the contract event 0x64b3528974574a41ddc830adc96440dccb99e64ede6b863194b951bfbd17e1bc.
//
// Solidity: event AddDefaultAggchainVKey(bytes4 selector, bytes32 newVKey)
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) ParseAddDefaultAggchainVKey(log types.Log) (*AgglayergatewaypreviousAddDefaultAggchainVKey, error) {
	event := new(AgglayergatewaypreviousAddDefaultAggchainVKey)
	if err := _Agglayergatewayprevious.contract.UnpackLog(event, "AddDefaultAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayergatewaypreviousInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Agglayergatewayprevious contract.
type AgglayergatewaypreviousInitializedIterator struct {
	Event *AgglayergatewaypreviousInitialized // Event containing the contract specifics and raw log

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
func (it *AgglayergatewaypreviousInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayergatewaypreviousInitialized)
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
		it.Event = new(AgglayergatewaypreviousInitialized)
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
func (it *AgglayergatewaypreviousInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayergatewaypreviousInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayergatewaypreviousInitialized represents a Initialized event raised by the Agglayergatewayprevious contract.
type AgglayergatewaypreviousInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) FilterInitialized(opts *bind.FilterOpts) (*AgglayergatewaypreviousInitializedIterator, error) {

	logs, sub, err := _Agglayergatewayprevious.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AgglayergatewaypreviousInitializedIterator{contract: _Agglayergatewayprevious.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AgglayergatewaypreviousInitialized) (event.Subscription, error) {

	logs, sub, err := _Agglayergatewayprevious.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayergatewaypreviousInitialized)
				if err := _Agglayergatewayprevious.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) ParseInitialized(log types.Log) (*AgglayergatewaypreviousInitialized, error) {
	event := new(AgglayergatewaypreviousInitialized)
	if err := _Agglayergatewayprevious.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayergatewaypreviousRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Agglayergatewayprevious contract.
type AgglayergatewaypreviousRoleAdminChangedIterator struct {
	Event *AgglayergatewaypreviousRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AgglayergatewaypreviousRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayergatewaypreviousRoleAdminChanged)
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
		it.Event = new(AgglayergatewaypreviousRoleAdminChanged)
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
func (it *AgglayergatewaypreviousRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayergatewaypreviousRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayergatewaypreviousRoleAdminChanged represents a RoleAdminChanged event raised by the Agglayergatewayprevious contract.
type AgglayergatewaypreviousRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AgglayergatewaypreviousRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Agglayergatewayprevious.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AgglayergatewaypreviousRoleAdminChangedIterator{contract: _Agglayergatewayprevious.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AgglayergatewaypreviousRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Agglayergatewayprevious.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayergatewaypreviousRoleAdminChanged)
				if err := _Agglayergatewayprevious.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) ParseRoleAdminChanged(log types.Log) (*AgglayergatewaypreviousRoleAdminChanged, error) {
	event := new(AgglayergatewaypreviousRoleAdminChanged)
	if err := _Agglayergatewayprevious.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayergatewaypreviousRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Agglayergatewayprevious contract.
type AgglayergatewaypreviousRoleGrantedIterator struct {
	Event *AgglayergatewaypreviousRoleGranted // Event containing the contract specifics and raw log

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
func (it *AgglayergatewaypreviousRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayergatewaypreviousRoleGranted)
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
		it.Event = new(AgglayergatewaypreviousRoleGranted)
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
func (it *AgglayergatewaypreviousRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayergatewaypreviousRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayergatewaypreviousRoleGranted represents a RoleGranted event raised by the Agglayergatewayprevious contract.
type AgglayergatewaypreviousRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AgglayergatewaypreviousRoleGrantedIterator, error) {

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

	logs, sub, err := _Agglayergatewayprevious.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AgglayergatewaypreviousRoleGrantedIterator{contract: _Agglayergatewayprevious.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AgglayergatewaypreviousRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Agglayergatewayprevious.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayergatewaypreviousRoleGranted)
				if err := _Agglayergatewayprevious.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) ParseRoleGranted(log types.Log) (*AgglayergatewaypreviousRoleGranted, error) {
	event := new(AgglayergatewaypreviousRoleGranted)
	if err := _Agglayergatewayprevious.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayergatewaypreviousRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Agglayergatewayprevious contract.
type AgglayergatewaypreviousRoleRevokedIterator struct {
	Event *AgglayergatewaypreviousRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AgglayergatewaypreviousRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayergatewaypreviousRoleRevoked)
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
		it.Event = new(AgglayergatewaypreviousRoleRevoked)
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
func (it *AgglayergatewaypreviousRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayergatewaypreviousRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayergatewaypreviousRoleRevoked represents a RoleRevoked event raised by the Agglayergatewayprevious contract.
type AgglayergatewaypreviousRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AgglayergatewaypreviousRoleRevokedIterator, error) {

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

	logs, sub, err := _Agglayergatewayprevious.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AgglayergatewaypreviousRoleRevokedIterator{contract: _Agglayergatewayprevious.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AgglayergatewaypreviousRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Agglayergatewayprevious.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayergatewaypreviousRoleRevoked)
				if err := _Agglayergatewayprevious.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) ParseRoleRevoked(log types.Log) (*AgglayergatewaypreviousRoleRevoked, error) {
	event := new(AgglayergatewaypreviousRoleRevoked)
	if err := _Agglayergatewayprevious.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayergatewaypreviousRouteAddedIterator is returned from FilterRouteAdded and is used to iterate over the raw logs and unpacked data for RouteAdded events raised by the Agglayergatewayprevious contract.
type AgglayergatewaypreviousRouteAddedIterator struct {
	Event *AgglayergatewaypreviousRouteAdded // Event containing the contract specifics and raw log

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
func (it *AgglayergatewaypreviousRouteAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayergatewaypreviousRouteAdded)
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
		it.Event = new(AgglayergatewaypreviousRouteAdded)
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
func (it *AgglayergatewaypreviousRouteAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayergatewaypreviousRouteAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayergatewaypreviousRouteAdded represents a RouteAdded event raised by the Agglayergatewayprevious contract.
type AgglayergatewaypreviousRouteAdded struct {
	Selector        [4]byte
	Verifier        common.Address
	PessimisticVKey [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRouteAdded is a free log retrieval operation binding the contract event 0x4c4455a30528c319643ae6fc35f3a0fcabde6c015f2a1a5270382c4a190b0fa3.
//
// Solidity: event RouteAdded(bytes4 selector, address verifier, bytes32 pessimisticVKey)
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) FilterRouteAdded(opts *bind.FilterOpts) (*AgglayergatewaypreviousRouteAddedIterator, error) {

	logs, sub, err := _Agglayergatewayprevious.contract.FilterLogs(opts, "RouteAdded")
	if err != nil {
		return nil, err
	}
	return &AgglayergatewaypreviousRouteAddedIterator{contract: _Agglayergatewayprevious.contract, event: "RouteAdded", logs: logs, sub: sub}, nil
}

// WatchRouteAdded is a free log subscription operation binding the contract event 0x4c4455a30528c319643ae6fc35f3a0fcabde6c015f2a1a5270382c4a190b0fa3.
//
// Solidity: event RouteAdded(bytes4 selector, address verifier, bytes32 pessimisticVKey)
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) WatchRouteAdded(opts *bind.WatchOpts, sink chan<- *AgglayergatewaypreviousRouteAdded) (event.Subscription, error) {

	logs, sub, err := _Agglayergatewayprevious.contract.WatchLogs(opts, "RouteAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayergatewaypreviousRouteAdded)
				if err := _Agglayergatewayprevious.contract.UnpackLog(event, "RouteAdded", log); err != nil {
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

// ParseRouteAdded is a log parse operation binding the contract event 0x4c4455a30528c319643ae6fc35f3a0fcabde6c015f2a1a5270382c4a190b0fa3.
//
// Solidity: event RouteAdded(bytes4 selector, address verifier, bytes32 pessimisticVKey)
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) ParseRouteAdded(log types.Log) (*AgglayergatewaypreviousRouteAdded, error) {
	event := new(AgglayergatewaypreviousRouteAdded)
	if err := _Agglayergatewayprevious.contract.UnpackLog(event, "RouteAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayergatewaypreviousRouteFrozenIterator is returned from FilterRouteFrozen and is used to iterate over the raw logs and unpacked data for RouteFrozen events raised by the Agglayergatewayprevious contract.
type AgglayergatewaypreviousRouteFrozenIterator struct {
	Event *AgglayergatewaypreviousRouteFrozen // Event containing the contract specifics and raw log

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
func (it *AgglayergatewaypreviousRouteFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayergatewaypreviousRouteFrozen)
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
		it.Event = new(AgglayergatewaypreviousRouteFrozen)
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
func (it *AgglayergatewaypreviousRouteFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayergatewaypreviousRouteFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayergatewaypreviousRouteFrozen represents a RouteFrozen event raised by the Agglayergatewayprevious contract.
type AgglayergatewaypreviousRouteFrozen struct {
	Selector        [4]byte
	Verifier        common.Address
	PessimisticVKey [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRouteFrozen is a free log retrieval operation binding the contract event 0x6efcfd4090b998a9d58c7f753bdda3291e9d781c02226b9f480b18ef07e0012b.
//
// Solidity: event RouteFrozen(bytes4 selector, address verifier, bytes32 pessimisticVKey)
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) FilterRouteFrozen(opts *bind.FilterOpts) (*AgglayergatewaypreviousRouteFrozenIterator, error) {

	logs, sub, err := _Agglayergatewayprevious.contract.FilterLogs(opts, "RouteFrozen")
	if err != nil {
		return nil, err
	}
	return &AgglayergatewaypreviousRouteFrozenIterator{contract: _Agglayergatewayprevious.contract, event: "RouteFrozen", logs: logs, sub: sub}, nil
}

// WatchRouteFrozen is a free log subscription operation binding the contract event 0x6efcfd4090b998a9d58c7f753bdda3291e9d781c02226b9f480b18ef07e0012b.
//
// Solidity: event RouteFrozen(bytes4 selector, address verifier, bytes32 pessimisticVKey)
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) WatchRouteFrozen(opts *bind.WatchOpts, sink chan<- *AgglayergatewaypreviousRouteFrozen) (event.Subscription, error) {

	logs, sub, err := _Agglayergatewayprevious.contract.WatchLogs(opts, "RouteFrozen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayergatewaypreviousRouteFrozen)
				if err := _Agglayergatewayprevious.contract.UnpackLog(event, "RouteFrozen", log); err != nil {
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

// ParseRouteFrozen is a log parse operation binding the contract event 0x6efcfd4090b998a9d58c7f753bdda3291e9d781c02226b9f480b18ef07e0012b.
//
// Solidity: event RouteFrozen(bytes4 selector, address verifier, bytes32 pessimisticVKey)
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) ParseRouteFrozen(log types.Log) (*AgglayergatewaypreviousRouteFrozen, error) {
	event := new(AgglayergatewaypreviousRouteFrozen)
	if err := _Agglayergatewayprevious.contract.UnpackLog(event, "RouteFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayergatewaypreviousUnsetDefaultAggchainVKeyIterator is returned from FilterUnsetDefaultAggchainVKey and is used to iterate over the raw logs and unpacked data for UnsetDefaultAggchainVKey events raised by the Agglayergatewayprevious contract.
type AgglayergatewaypreviousUnsetDefaultAggchainVKeyIterator struct {
	Event *AgglayergatewaypreviousUnsetDefaultAggchainVKey // Event containing the contract specifics and raw log

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
func (it *AgglayergatewaypreviousUnsetDefaultAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayergatewaypreviousUnsetDefaultAggchainVKey)
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
		it.Event = new(AgglayergatewaypreviousUnsetDefaultAggchainVKey)
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
func (it *AgglayergatewaypreviousUnsetDefaultAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayergatewaypreviousUnsetDefaultAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayergatewaypreviousUnsetDefaultAggchainVKey represents a UnsetDefaultAggchainVKey event raised by the Agglayergatewayprevious contract.
type AgglayergatewaypreviousUnsetDefaultAggchainVKey struct {
	Selector [4]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnsetDefaultAggchainVKey is a free log retrieval operation binding the contract event 0x43b46d79ee2fbaa8a981b5fd78aaf6fb9e38b18d02e1af17564ce37c67078cab.
//
// Solidity: event UnsetDefaultAggchainVKey(bytes4 selector)
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) FilterUnsetDefaultAggchainVKey(opts *bind.FilterOpts) (*AgglayergatewaypreviousUnsetDefaultAggchainVKeyIterator, error) {

	logs, sub, err := _Agglayergatewayprevious.contract.FilterLogs(opts, "UnsetDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &AgglayergatewaypreviousUnsetDefaultAggchainVKeyIterator{contract: _Agglayergatewayprevious.contract, event: "UnsetDefaultAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchUnsetDefaultAggchainVKey is a free log subscription operation binding the contract event 0x43b46d79ee2fbaa8a981b5fd78aaf6fb9e38b18d02e1af17564ce37c67078cab.
//
// Solidity: event UnsetDefaultAggchainVKey(bytes4 selector)
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) WatchUnsetDefaultAggchainVKey(opts *bind.WatchOpts, sink chan<- *AgglayergatewaypreviousUnsetDefaultAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Agglayergatewayprevious.contract.WatchLogs(opts, "UnsetDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayergatewaypreviousUnsetDefaultAggchainVKey)
				if err := _Agglayergatewayprevious.contract.UnpackLog(event, "UnsetDefaultAggchainVKey", log); err != nil {
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

// ParseUnsetDefaultAggchainVKey is a log parse operation binding the contract event 0x43b46d79ee2fbaa8a981b5fd78aaf6fb9e38b18d02e1af17564ce37c67078cab.
//
// Solidity: event UnsetDefaultAggchainVKey(bytes4 selector)
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) ParseUnsetDefaultAggchainVKey(log types.Log) (*AgglayergatewaypreviousUnsetDefaultAggchainVKey, error) {
	event := new(AgglayergatewaypreviousUnsetDefaultAggchainVKey)
	if err := _Agglayergatewayprevious.contract.UnpackLog(event, "UnsetDefaultAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayergatewaypreviousUpdateDefaultAggchainVKeyIterator is returned from FilterUpdateDefaultAggchainVKey and is used to iterate over the raw logs and unpacked data for UpdateDefaultAggchainVKey events raised by the Agglayergatewayprevious contract.
type AgglayergatewaypreviousUpdateDefaultAggchainVKeyIterator struct {
	Event *AgglayergatewaypreviousUpdateDefaultAggchainVKey // Event containing the contract specifics and raw log

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
func (it *AgglayergatewaypreviousUpdateDefaultAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayergatewaypreviousUpdateDefaultAggchainVKey)
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
		it.Event = new(AgglayergatewaypreviousUpdateDefaultAggchainVKey)
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
func (it *AgglayergatewaypreviousUpdateDefaultAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayergatewaypreviousUpdateDefaultAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayergatewaypreviousUpdateDefaultAggchainVKey represents a UpdateDefaultAggchainVKey event raised by the Agglayergatewayprevious contract.
type AgglayergatewaypreviousUpdateDefaultAggchainVKey struct {
	Selector     [4]byte
	PreviousVKey [32]byte
	NewVKey      [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateDefaultAggchainVKey is a free log retrieval operation binding the contract event 0x477ac863279848383fdb5d6bb8957cc8f6f10c0200379e075df3f789d14248e6.
//
// Solidity: event UpdateDefaultAggchainVKey(bytes4 selector, bytes32 previousVKey, bytes32 newVKey)
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) FilterUpdateDefaultAggchainVKey(opts *bind.FilterOpts) (*AgglayergatewaypreviousUpdateDefaultAggchainVKeyIterator, error) {

	logs, sub, err := _Agglayergatewayprevious.contract.FilterLogs(opts, "UpdateDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &AgglayergatewaypreviousUpdateDefaultAggchainVKeyIterator{contract: _Agglayergatewayprevious.contract, event: "UpdateDefaultAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchUpdateDefaultAggchainVKey is a free log subscription operation binding the contract event 0x477ac863279848383fdb5d6bb8957cc8f6f10c0200379e075df3f789d14248e6.
//
// Solidity: event UpdateDefaultAggchainVKey(bytes4 selector, bytes32 previousVKey, bytes32 newVKey)
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) WatchUpdateDefaultAggchainVKey(opts *bind.WatchOpts, sink chan<- *AgglayergatewaypreviousUpdateDefaultAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Agglayergatewayprevious.contract.WatchLogs(opts, "UpdateDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayergatewaypreviousUpdateDefaultAggchainVKey)
				if err := _Agglayergatewayprevious.contract.UnpackLog(event, "UpdateDefaultAggchainVKey", log); err != nil {
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

// ParseUpdateDefaultAggchainVKey is a log parse operation binding the contract event 0x477ac863279848383fdb5d6bb8957cc8f6f10c0200379e075df3f789d14248e6.
//
// Solidity: event UpdateDefaultAggchainVKey(bytes4 selector, bytes32 previousVKey, bytes32 newVKey)
func (_Agglayergatewayprevious *AgglayergatewaypreviousFilterer) ParseUpdateDefaultAggchainVKey(log types.Log) (*AgglayergatewaypreviousUpdateDefaultAggchainVKey, error) {
	event := new(AgglayergatewaypreviousUpdateDefaultAggchainVKey)
	if err := _Agglayergatewayprevious.contract.UnpackLog(event, "UpdateDefaultAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

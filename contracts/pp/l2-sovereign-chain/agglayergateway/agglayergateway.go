// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package agglayergateway

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

// AgglayergatewayMetaData contains all meta data concerning the Agglayergateway contract.
var AgglayergatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainVKeyAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAggLayerAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAggLayerAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PPSelectorCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"RouteAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"RouteIsAlreadyFrozen\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"RouteIsFrozen\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"RouteNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VKeyCannotBeZero\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newVKey\",\"type\":\"bytes32\"}],\"name\":\"AddDefaultAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pessimisticVKey\",\"type\":\"bytes32\"}],\"name\":\"RouteAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pessimisticVKey\",\"type\":\"bytes32\"}],\"name\":\"RouteFrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousVKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newVKey\",\"type\":\"bytes32\"}],\"name\":\"UpdateDefaultAggchainVKey\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"defaultAggchainSelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"addDefaultAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"pessimisticVKeySelector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"pessimisticVKey\",\"type\":\"bytes32\"}],\"name\":\"addPessimisticVKeyRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"defaultAggchainSelector\",\"type\":\"bytes4\"}],\"name\":\"defaultAggchainVKeys\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"defaultAggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"pessimisticVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"freezePessimisticVKeyRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"defaultAggchainSelector\",\"type\":\"bytes4\"}],\"name\":\"getDefaultAggchainVKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"aggchainDefaultVKeyRole\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addRouteRole\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"freezeRouteRole\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"pessimisticVKeySelector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"pessimisticVKey\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"pessimisticVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"pessimisticVKeyRoutes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"pessimisticVKey\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"defaultAggchainSelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"newDefaultAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"updateDefaultAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicValues\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofBytes\",\"type\":\"bytes\"}],\"name\":\"verifyPessimisticProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060156019565b60c9565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560685760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c65780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6114e6806100d65f395ff3fe608060405234801561000f575f5ffd5b50600436106100fb575f3560e01c806382bfaea111610093578063a217fddf11610063578063a217fddf146102bd578063a48fd377146102c4578063d547741f146102d7578063f4c44681146102ea575f5ffd5b806382bfaea11461023557806391d14854146102545780639584a5161461029757806395d2a391146102aa575f5ffd5b80636cabfdab116100ce5780636cabfdab1461017f578063716be075146101925780637b552bea1461020f578063818c8c2114610222575f5ffd5b806301ffc9a7146100ff578063248a9ca3146101275780632f2ff15d1461015757806336568abe1461016c575b5f5ffd5b61011261010d366004611190565b6102fd565b60405190151581526020015b60405180910390f35b6101496101353660046111b0565b5f9081526020819052604090206001015490565b60405190815260200161011e565b61016a6101653660046111ea565b610395565b005b61016a61017a3660046111ea565b6103bf565b61014961018d366004611190565b61041d565b6101db6101a0366004611190565b600260208190525f918252604090912080546001820154919092015473ffffffffffffffffffffffffffffffffffffffff9092169160ff1683565b6040805173ffffffffffffffffffffffffffffffffffffffff9094168452602084019290925215159082015260600161011e565b61016a61021d366004611214565b6104b8565b61016a610230366004611291565b6106ca565b610149610243366004611190565b60016020525f908152604090205481565b6101126102623660046111ea565b5f9182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b61016a6102a53660046112b9565b610805565b61016a6102b8366004611190565b61083a565b6101495f81565b61016a6102d2366004611338565b610a17565b61016a6102e53660046111ea565b610be6565b61016a6102f8366004611291565b610c0a565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061038f57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b5f828152602081905260409020600101546103af81610d1b565b6103b98383610d28565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8116331461040e576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6104188282610e21565b505050565b7fffffffff0000000000000000000000000000000000000000000000000000000081165f90815260016020526040812054610484576040517f925e5a3a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b507fffffffff00000000000000000000000000000000000000000000000000000000165f9081526001602052604090205490565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f811580156105025750825b90505f8267ffffffffffffffff16600114801561051e5750303b155b90508115801561052c575080155b15610563576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156105c45784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6105ce5f8d610d28565b506105f97f131410eab1236cee2db19035b0e825c94e5ab705dffe23321dd53856da5316178c610d28565b506106247f0fdc2a718b96bc741c7544001e3dd7c26730802c54781668fa78a120e622629b8b610d28565b5061064f7fca75ae4228cde6195f9fa3dbde8dc352fb30aa63780717a378ccfc50274355dd8a610d28565b5061065b888888610eda565b83156106bc5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050505050565b7f131410eab1236cee2db19035b0e825c94e5ab705dffe23321dd53856da5316176106f481610d1b565b7fffffffff0000000000000000000000000000000000000000000000000000000083165f908152600160205260409020541561075c576040517f22a1bdc400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81610793576040517f6745305e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000083165f81815260016020908152604091829020859055815192835282018490527f64b3528974574a41ddc830adc96440dccb99e64ede6b863194b951bfbd17e1bc91015b60405180910390a1505050565b7f0fdc2a718b96bc741c7544001e3dd7c26730802c54781668fa78a120e622629b61082f81610d1b565b6103b9848484610eda565b7fca75ae4228cde6195f9fa3dbde8dc352fb30aa63780717a378ccfc50274355dd61086481610d1b565b7fffffffff0000000000000000000000000000000000000000000000000000000082165f908152600260205260409020805473ffffffffffffffffffffffffffffffffffffffff1661090b576040517ff208777e0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000841660048201526024015b60405180910390fd5b600281015460ff161561096e576040517fc7b4a1f20000000000000000000000000000000000000000000000000000000081527fffffffff0000000000000000000000000000000000000000000000000000000084166004820152602401610902565b6002810180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001908117909155815490820154604080517fffffffff000000000000000000000000000000000000000000000000000000008716815273ffffffffffffffffffffffffffffffffffffffff90931660208401528201527f6efcfd4090b998a9d58c7f753bdda3291e9d781c02226b9f480b18ef07e0012b906060016107f8565b5f610a2560048284866113a4565b610a2e916113cb565b7fffffffff0000000000000000000000000000000000000000000000000000000081165f908152600260208181526040928390208351606081018552815473ffffffffffffffffffffffffffffffffffffffff1680825260018301549382019390935292015460ff16151592820192909252919250610afd576040517ff208777e0000000000000000000000000000000000000000000000000000000081527fffffffff0000000000000000000000000000000000000000000000000000000083166004820152602401610902565b806040015115610b5d576040517febf108230000000000000000000000000000000000000000000000000000000081527fffffffff0000000000000000000000000000000000000000000000000000000083166004820152602401610902565b8051602082015173ffffffffffffffffffffffffffffffffffffffff909116906341493c60908888610b92886004818c6113a4565b6040518663ffffffff1660e01b8152600401610bb2959493929190611478565b5f6040518083038186803b158015610bc8575f5ffd5b505afa158015610bda573d5f5f3e3d5ffd5b50505050505050505050565b5f82815260208190526040902060010154610c0081610d1b565b6103b98383610e21565b7f131410eab1236cee2db19035b0e825c94e5ab705dffe23321dd53856da531617610c3481610d1b565b7fffffffff0000000000000000000000000000000000000000000000000000000083165f90815260016020526040902054610c9b576040517f925e5a3a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000083165f818152600160209081526040918290208054908690558251938452908301819052908201849052907f477ac863279848383fdb5d6bb8957cc8f6f10c0200379e075df3f789d14248e6906060015b60405180910390a150505050565b610d2581336110d3565b50565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16610e1a575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610db83390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600161038f565b505f61038f565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615610e1a575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a450600161038f565b7fffffffff000000000000000000000000000000000000000000000000000000008316610f33576040517f7500f20900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80610f6a576040517f6745305e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000083165f908152600260205260409020805473ffffffffffffffffffffffffffffffffffffffff161561102c5780546040517f4bc7c2780000000000000000000000000000000000000000000000000000000081527fffffffff000000000000000000000000000000000000000000000000000000008616600482015273ffffffffffffffffffffffffffffffffffffffff9091166024820152604401610902565b80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8416908117825560018201839055604080517fffffffff0000000000000000000000000000000000000000000000000000000087168152602081019290925281018390527f4c4455a30528c319643ae6fc35f3a0fcabde6c015f2a1a5270382c4a190b0fa390606001610d0d565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16611158576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610902565b5050565b80357fffffffff000000000000000000000000000000000000000000000000000000008116811461118b575f5ffd5b919050565b5f602082840312156111a0575f5ffd5b6111a98261115c565b9392505050565b5f602082840312156111c0575f5ffd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461118b575f5ffd5b5f5f604083850312156111fb575f5ffd5b8235915061120b602084016111c7565b90509250929050565b5f5f5f5f5f5f5f60e0888a03121561122a575f5ffd5b611233886111c7565b9650611241602089016111c7565b955061124f604089016111c7565b945061125d606089016111c7565b935061126b6080890161115c565b925061127960a089016111c7565b96999598509396929591949193505060c09091013590565b5f5f604083850312156112a2575f5ffd5b6112ab8361115c565b946020939093013593505050565b5f5f5f606084860312156112cb575f5ffd5b6112d48461115c565b92506112e2602085016111c7565b929592945050506040919091013590565b5f5f83601f840112611303575f5ffd5b50813567ffffffffffffffff81111561131a575f5ffd5b602083019150836020828501011115611331575f5ffd5b9250929050565b5f5f5f5f6040858703121561134b575f5ffd5b843567ffffffffffffffff811115611361575f5ffd5b61136d878288016112f3565b909550935050602085013567ffffffffffffffff81111561138c575f5ffd5b611398878288016112f3565b95989497509550505050565b5f5f858511156113b2575f5ffd5b838611156113be575f5ffd5b5050820193919092039150565b80357fffffffff00000000000000000000000000000000000000000000000000000000811690600484101561142a577fffffffff00000000000000000000000000000000000000000000000000000000808560040360031b1b82161691505b5092915050565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b858152606060208201525f611491606083018688611431565b82810360408401526114a4818587611431565b9897505050505050505056fea2646970667358221220dfadf1677e474cc686b98544e61522fcac2245f0ed590f9149a21f8cce9d305b64736f6c634300081c0033",
}

// AgglayergatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use AgglayergatewayMetaData.ABI instead.
var AgglayergatewayABI = AgglayergatewayMetaData.ABI

// AgglayergatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AgglayergatewayMetaData.Bin instead.
var AgglayergatewayBin = AgglayergatewayMetaData.Bin

// DeployAgglayergateway deploys a new Ethereum contract, binding an instance of Agglayergateway to it.
func DeployAgglayergateway(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Agglayergateway, error) {
	parsed, err := AgglayergatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AgglayergatewayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Agglayergateway{AgglayergatewayCaller: AgglayergatewayCaller{contract: contract}, AgglayergatewayTransactor: AgglayergatewayTransactor{contract: contract}, AgglayergatewayFilterer: AgglayergatewayFilterer{contract: contract}}, nil
}

// Agglayergateway is an auto generated Go binding around an Ethereum contract.
type Agglayergateway struct {
	AgglayergatewayCaller     // Read-only binding to the contract
	AgglayergatewayTransactor // Write-only binding to the contract
	AgglayergatewayFilterer   // Log filterer for contract events
}

// AgglayergatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgglayergatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgglayergatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgglayergatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgglayergatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgglayergatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgglayergatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgglayergatewaySession struct {
	Contract     *Agglayergateway  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgglayergatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgglayergatewayCallerSession struct {
	Contract *AgglayergatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AgglayergatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgglayergatewayTransactorSession struct {
	Contract     *AgglayergatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AgglayergatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgglayergatewayRaw struct {
	Contract *Agglayergateway // Generic contract binding to access the raw methods on
}

// AgglayergatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgglayergatewayCallerRaw struct {
	Contract *AgglayergatewayCaller // Generic read-only contract binding to access the raw methods on
}

// AgglayergatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgglayergatewayTransactorRaw struct {
	Contract *AgglayergatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgglayergateway creates a new instance of Agglayergateway, bound to a specific deployed contract.
func NewAgglayergateway(address common.Address, backend bind.ContractBackend) (*Agglayergateway, error) {
	contract, err := bindAgglayergateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Agglayergateway{AgglayergatewayCaller: AgglayergatewayCaller{contract: contract}, AgglayergatewayTransactor: AgglayergatewayTransactor{contract: contract}, AgglayergatewayFilterer: AgglayergatewayFilterer{contract: contract}}, nil
}

// NewAgglayergatewayCaller creates a new read-only instance of Agglayergateway, bound to a specific deployed contract.
func NewAgglayergatewayCaller(address common.Address, caller bind.ContractCaller) (*AgglayergatewayCaller, error) {
	contract, err := bindAgglayergateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgglayergatewayCaller{contract: contract}, nil
}

// NewAgglayergatewayTransactor creates a new write-only instance of Agglayergateway, bound to a specific deployed contract.
func NewAgglayergatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*AgglayergatewayTransactor, error) {
	contract, err := bindAgglayergateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgglayergatewayTransactor{contract: contract}, nil
}

// NewAgglayergatewayFilterer creates a new log filterer instance of Agglayergateway, bound to a specific deployed contract.
func NewAgglayergatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*AgglayergatewayFilterer, error) {
	contract, err := bindAgglayergateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgglayergatewayFilterer{contract: contract}, nil
}

// bindAgglayergateway binds a generic wrapper to an already deployed contract.
func bindAgglayergateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AgglayergatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Agglayergateway *AgglayergatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Agglayergateway.Contract.AgglayergatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Agglayergateway *AgglayergatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayergateway.Contract.AgglayergatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Agglayergateway *AgglayergatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Agglayergateway.Contract.AgglayergatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Agglayergateway *AgglayergatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Agglayergateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Agglayergateway *AgglayergatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayergateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Agglayergateway *AgglayergatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Agglayergateway.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Agglayergateway *AgglayergatewayCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Agglayergateway.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Agglayergateway *AgglayergatewaySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Agglayergateway.Contract.DEFAULTADMINROLE(&_Agglayergateway.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Agglayergateway *AgglayergatewayCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Agglayergateway.Contract.DEFAULTADMINROLE(&_Agglayergateway.CallOpts)
}

// DefaultAggchainVKeys is a free data retrieval call binding the contract method 0x82bfaea1.
//
// Solidity: function defaultAggchainVKeys(bytes4 defaultAggchainSelector) view returns(bytes32 defaultAggchainVKey)
func (_Agglayergateway *AgglayergatewayCaller) DefaultAggchainVKeys(opts *bind.CallOpts, defaultAggchainSelector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _Agglayergateway.contract.Call(opts, &out, "defaultAggchainVKeys", defaultAggchainSelector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DefaultAggchainVKeys is a free data retrieval call binding the contract method 0x82bfaea1.
//
// Solidity: function defaultAggchainVKeys(bytes4 defaultAggchainSelector) view returns(bytes32 defaultAggchainVKey)
func (_Agglayergateway *AgglayergatewaySession) DefaultAggchainVKeys(defaultAggchainSelector [4]byte) ([32]byte, error) {
	return _Agglayergateway.Contract.DefaultAggchainVKeys(&_Agglayergateway.CallOpts, defaultAggchainSelector)
}

// DefaultAggchainVKeys is a free data retrieval call binding the contract method 0x82bfaea1.
//
// Solidity: function defaultAggchainVKeys(bytes4 defaultAggchainSelector) view returns(bytes32 defaultAggchainVKey)
func (_Agglayergateway *AgglayergatewayCallerSession) DefaultAggchainVKeys(defaultAggchainSelector [4]byte) ([32]byte, error) {
	return _Agglayergateway.Contract.DefaultAggchainVKeys(&_Agglayergateway.CallOpts, defaultAggchainSelector)
}

// GetDefaultAggchainVKey is a free data retrieval call binding the contract method 0x6cabfdab.
//
// Solidity: function getDefaultAggchainVKey(bytes4 defaultAggchainSelector) view returns(bytes32)
func (_Agglayergateway *AgglayergatewayCaller) GetDefaultAggchainVKey(opts *bind.CallOpts, defaultAggchainSelector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _Agglayergateway.contract.Call(opts, &out, "getDefaultAggchainVKey", defaultAggchainSelector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDefaultAggchainVKey is a free data retrieval call binding the contract method 0x6cabfdab.
//
// Solidity: function getDefaultAggchainVKey(bytes4 defaultAggchainSelector) view returns(bytes32)
func (_Agglayergateway *AgglayergatewaySession) GetDefaultAggchainVKey(defaultAggchainSelector [4]byte) ([32]byte, error) {
	return _Agglayergateway.Contract.GetDefaultAggchainVKey(&_Agglayergateway.CallOpts, defaultAggchainSelector)
}

// GetDefaultAggchainVKey is a free data retrieval call binding the contract method 0x6cabfdab.
//
// Solidity: function getDefaultAggchainVKey(bytes4 defaultAggchainSelector) view returns(bytes32)
func (_Agglayergateway *AgglayergatewayCallerSession) GetDefaultAggchainVKey(defaultAggchainSelector [4]byte) ([32]byte, error) {
	return _Agglayergateway.Contract.GetDefaultAggchainVKey(&_Agglayergateway.CallOpts, defaultAggchainSelector)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Agglayergateway *AgglayergatewayCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Agglayergateway.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Agglayergateway *AgglayergatewaySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Agglayergateway.Contract.GetRoleAdmin(&_Agglayergateway.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Agglayergateway *AgglayergatewayCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Agglayergateway.Contract.GetRoleAdmin(&_Agglayergateway.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Agglayergateway *AgglayergatewayCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Agglayergateway.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Agglayergateway *AgglayergatewaySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Agglayergateway.Contract.HasRole(&_Agglayergateway.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Agglayergateway *AgglayergatewayCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Agglayergateway.Contract.HasRole(&_Agglayergateway.CallOpts, role, account)
}

// PessimisticVKeyRoutes is a free data retrieval call binding the contract method 0x716be075.
//
// Solidity: function pessimisticVKeyRoutes(bytes4 pessimisticVKeySelector) view returns(address verifier, bytes32 pessimisticVKey, bool frozen)
func (_Agglayergateway *AgglayergatewayCaller) PessimisticVKeyRoutes(opts *bind.CallOpts, pessimisticVKeySelector [4]byte) (struct {
	Verifier        common.Address
	PessimisticVKey [32]byte
	Frozen          bool
}, error) {
	var out []interface{}
	err := _Agglayergateway.contract.Call(opts, &out, "pessimisticVKeyRoutes", pessimisticVKeySelector)

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
func (_Agglayergateway *AgglayergatewaySession) PessimisticVKeyRoutes(pessimisticVKeySelector [4]byte) (struct {
	Verifier        common.Address
	PessimisticVKey [32]byte
	Frozen          bool
}, error) {
	return _Agglayergateway.Contract.PessimisticVKeyRoutes(&_Agglayergateway.CallOpts, pessimisticVKeySelector)
}

// PessimisticVKeyRoutes is a free data retrieval call binding the contract method 0x716be075.
//
// Solidity: function pessimisticVKeyRoutes(bytes4 pessimisticVKeySelector) view returns(address verifier, bytes32 pessimisticVKey, bool frozen)
func (_Agglayergateway *AgglayergatewayCallerSession) PessimisticVKeyRoutes(pessimisticVKeySelector [4]byte) (struct {
	Verifier        common.Address
	PessimisticVKey [32]byte
	Frozen          bool
}, error) {
	return _Agglayergateway.Contract.PessimisticVKeyRoutes(&_Agglayergateway.CallOpts, pessimisticVKeySelector)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Agglayergateway *AgglayergatewayCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Agglayergateway.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Agglayergateway *AgglayergatewaySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Agglayergateway.Contract.SupportsInterface(&_Agglayergateway.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Agglayergateway *AgglayergatewayCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Agglayergateway.Contract.SupportsInterface(&_Agglayergateway.CallOpts, interfaceId)
}

// VerifyPessimisticProof is a free data retrieval call binding the contract method 0xa48fd377.
//
// Solidity: function verifyPessimisticProof(bytes publicValues, bytes proofBytes) view returns()
func (_Agglayergateway *AgglayergatewayCaller) VerifyPessimisticProof(opts *bind.CallOpts, publicValues []byte, proofBytes []byte) error {
	var out []interface{}
	err := _Agglayergateway.contract.Call(opts, &out, "verifyPessimisticProof", publicValues, proofBytes)

	if err != nil {
		return err
	}

	return err

}

// VerifyPessimisticProof is a free data retrieval call binding the contract method 0xa48fd377.
//
// Solidity: function verifyPessimisticProof(bytes publicValues, bytes proofBytes) view returns()
func (_Agglayergateway *AgglayergatewaySession) VerifyPessimisticProof(publicValues []byte, proofBytes []byte) error {
	return _Agglayergateway.Contract.VerifyPessimisticProof(&_Agglayergateway.CallOpts, publicValues, proofBytes)
}

// VerifyPessimisticProof is a free data retrieval call binding the contract method 0xa48fd377.
//
// Solidity: function verifyPessimisticProof(bytes publicValues, bytes proofBytes) view returns()
func (_Agglayergateway *AgglayergatewayCallerSession) VerifyPessimisticProof(publicValues []byte, proofBytes []byte) error {
	return _Agglayergateway.Contract.VerifyPessimisticProof(&_Agglayergateway.CallOpts, publicValues, proofBytes)
}

// AddDefaultAggchainVKey is a paid mutator transaction binding the contract method 0x818c8c21.
//
// Solidity: function addDefaultAggchainVKey(bytes4 defaultAggchainSelector, bytes32 newAggchainVKey) returns()
func (_Agglayergateway *AgglayergatewayTransactor) AddDefaultAggchainVKey(opts *bind.TransactOpts, defaultAggchainSelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Agglayergateway.contract.Transact(opts, "addDefaultAggchainVKey", defaultAggchainSelector, newAggchainVKey)
}

// AddDefaultAggchainVKey is a paid mutator transaction binding the contract method 0x818c8c21.
//
// Solidity: function addDefaultAggchainVKey(bytes4 defaultAggchainSelector, bytes32 newAggchainVKey) returns()
func (_Agglayergateway *AgglayergatewaySession) AddDefaultAggchainVKey(defaultAggchainSelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Agglayergateway.Contract.AddDefaultAggchainVKey(&_Agglayergateway.TransactOpts, defaultAggchainSelector, newAggchainVKey)
}

// AddDefaultAggchainVKey is a paid mutator transaction binding the contract method 0x818c8c21.
//
// Solidity: function addDefaultAggchainVKey(bytes4 defaultAggchainSelector, bytes32 newAggchainVKey) returns()
func (_Agglayergateway *AgglayergatewayTransactorSession) AddDefaultAggchainVKey(defaultAggchainSelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Agglayergateway.Contract.AddDefaultAggchainVKey(&_Agglayergateway.TransactOpts, defaultAggchainSelector, newAggchainVKey)
}

// AddPessimisticVKeyRoute is a paid mutator transaction binding the contract method 0x9584a516.
//
// Solidity: function addPessimisticVKeyRoute(bytes4 pessimisticVKeySelector, address verifier, bytes32 pessimisticVKey) returns()
func (_Agglayergateway *AgglayergatewayTransactor) AddPessimisticVKeyRoute(opts *bind.TransactOpts, pessimisticVKeySelector [4]byte, verifier common.Address, pessimisticVKey [32]byte) (*types.Transaction, error) {
	return _Agglayergateway.contract.Transact(opts, "addPessimisticVKeyRoute", pessimisticVKeySelector, verifier, pessimisticVKey)
}

// AddPessimisticVKeyRoute is a paid mutator transaction binding the contract method 0x9584a516.
//
// Solidity: function addPessimisticVKeyRoute(bytes4 pessimisticVKeySelector, address verifier, bytes32 pessimisticVKey) returns()
func (_Agglayergateway *AgglayergatewaySession) AddPessimisticVKeyRoute(pessimisticVKeySelector [4]byte, verifier common.Address, pessimisticVKey [32]byte) (*types.Transaction, error) {
	return _Agglayergateway.Contract.AddPessimisticVKeyRoute(&_Agglayergateway.TransactOpts, pessimisticVKeySelector, verifier, pessimisticVKey)
}

// AddPessimisticVKeyRoute is a paid mutator transaction binding the contract method 0x9584a516.
//
// Solidity: function addPessimisticVKeyRoute(bytes4 pessimisticVKeySelector, address verifier, bytes32 pessimisticVKey) returns()
func (_Agglayergateway *AgglayergatewayTransactorSession) AddPessimisticVKeyRoute(pessimisticVKeySelector [4]byte, verifier common.Address, pessimisticVKey [32]byte) (*types.Transaction, error) {
	return _Agglayergateway.Contract.AddPessimisticVKeyRoute(&_Agglayergateway.TransactOpts, pessimisticVKeySelector, verifier, pessimisticVKey)
}

// FreezePessimisticVKeyRoute is a paid mutator transaction binding the contract method 0x95d2a391.
//
// Solidity: function freezePessimisticVKeyRoute(bytes4 pessimisticVKeySelector) returns()
func (_Agglayergateway *AgglayergatewayTransactor) FreezePessimisticVKeyRoute(opts *bind.TransactOpts, pessimisticVKeySelector [4]byte) (*types.Transaction, error) {
	return _Agglayergateway.contract.Transact(opts, "freezePessimisticVKeyRoute", pessimisticVKeySelector)
}

// FreezePessimisticVKeyRoute is a paid mutator transaction binding the contract method 0x95d2a391.
//
// Solidity: function freezePessimisticVKeyRoute(bytes4 pessimisticVKeySelector) returns()
func (_Agglayergateway *AgglayergatewaySession) FreezePessimisticVKeyRoute(pessimisticVKeySelector [4]byte) (*types.Transaction, error) {
	return _Agglayergateway.Contract.FreezePessimisticVKeyRoute(&_Agglayergateway.TransactOpts, pessimisticVKeySelector)
}

// FreezePessimisticVKeyRoute is a paid mutator transaction binding the contract method 0x95d2a391.
//
// Solidity: function freezePessimisticVKeyRoute(bytes4 pessimisticVKeySelector) returns()
func (_Agglayergateway *AgglayergatewayTransactorSession) FreezePessimisticVKeyRoute(pessimisticVKeySelector [4]byte) (*types.Transaction, error) {
	return _Agglayergateway.Contract.FreezePessimisticVKeyRoute(&_Agglayergateway.TransactOpts, pessimisticVKeySelector)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Agglayergateway *AgglayergatewayTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayergateway.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Agglayergateway *AgglayergatewaySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayergateway.Contract.GrantRole(&_Agglayergateway.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Agglayergateway *AgglayergatewayTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayergateway.Contract.GrantRole(&_Agglayergateway.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x7b552bea.
//
// Solidity: function initialize(address defaultAdmin, address aggchainDefaultVKeyRole, address addRouteRole, address freezeRouteRole, bytes4 pessimisticVKeySelector, address verifier, bytes32 pessimisticVKey) returns()
func (_Agglayergateway *AgglayergatewayTransactor) Initialize(opts *bind.TransactOpts, defaultAdmin common.Address, aggchainDefaultVKeyRole common.Address, addRouteRole common.Address, freezeRouteRole common.Address, pessimisticVKeySelector [4]byte, verifier common.Address, pessimisticVKey [32]byte) (*types.Transaction, error) {
	return _Agglayergateway.contract.Transact(opts, "initialize", defaultAdmin, aggchainDefaultVKeyRole, addRouteRole, freezeRouteRole, pessimisticVKeySelector, verifier, pessimisticVKey)
}

// Initialize is a paid mutator transaction binding the contract method 0x7b552bea.
//
// Solidity: function initialize(address defaultAdmin, address aggchainDefaultVKeyRole, address addRouteRole, address freezeRouteRole, bytes4 pessimisticVKeySelector, address verifier, bytes32 pessimisticVKey) returns()
func (_Agglayergateway *AgglayergatewaySession) Initialize(defaultAdmin common.Address, aggchainDefaultVKeyRole common.Address, addRouteRole common.Address, freezeRouteRole common.Address, pessimisticVKeySelector [4]byte, verifier common.Address, pessimisticVKey [32]byte) (*types.Transaction, error) {
	return _Agglayergateway.Contract.Initialize(&_Agglayergateway.TransactOpts, defaultAdmin, aggchainDefaultVKeyRole, addRouteRole, freezeRouteRole, pessimisticVKeySelector, verifier, pessimisticVKey)
}

// Initialize is a paid mutator transaction binding the contract method 0x7b552bea.
//
// Solidity: function initialize(address defaultAdmin, address aggchainDefaultVKeyRole, address addRouteRole, address freezeRouteRole, bytes4 pessimisticVKeySelector, address verifier, bytes32 pessimisticVKey) returns()
func (_Agglayergateway *AgglayergatewayTransactorSession) Initialize(defaultAdmin common.Address, aggchainDefaultVKeyRole common.Address, addRouteRole common.Address, freezeRouteRole common.Address, pessimisticVKeySelector [4]byte, verifier common.Address, pessimisticVKey [32]byte) (*types.Transaction, error) {
	return _Agglayergateway.Contract.Initialize(&_Agglayergateway.TransactOpts, defaultAdmin, aggchainDefaultVKeyRole, addRouteRole, freezeRouteRole, pessimisticVKeySelector, verifier, pessimisticVKey)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Agglayergateway *AgglayergatewayTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Agglayergateway.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Agglayergateway *AgglayergatewaySession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Agglayergateway.Contract.RenounceRole(&_Agglayergateway.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Agglayergateway *AgglayergatewayTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Agglayergateway.Contract.RenounceRole(&_Agglayergateway.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Agglayergateway *AgglayergatewayTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayergateway.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Agglayergateway *AgglayergatewaySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayergateway.Contract.RevokeRole(&_Agglayergateway.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Agglayergateway *AgglayergatewayTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayergateway.Contract.RevokeRole(&_Agglayergateway.TransactOpts, role, account)
}

// UpdateDefaultAggchainVKey is a paid mutator transaction binding the contract method 0xf4c44681.
//
// Solidity: function updateDefaultAggchainVKey(bytes4 defaultAggchainSelector, bytes32 newDefaultAggchainVKey) returns()
func (_Agglayergateway *AgglayergatewayTransactor) UpdateDefaultAggchainVKey(opts *bind.TransactOpts, defaultAggchainSelector [4]byte, newDefaultAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Agglayergateway.contract.Transact(opts, "updateDefaultAggchainVKey", defaultAggchainSelector, newDefaultAggchainVKey)
}

// UpdateDefaultAggchainVKey is a paid mutator transaction binding the contract method 0xf4c44681.
//
// Solidity: function updateDefaultAggchainVKey(bytes4 defaultAggchainSelector, bytes32 newDefaultAggchainVKey) returns()
func (_Agglayergateway *AgglayergatewaySession) UpdateDefaultAggchainVKey(defaultAggchainSelector [4]byte, newDefaultAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Agglayergateway.Contract.UpdateDefaultAggchainVKey(&_Agglayergateway.TransactOpts, defaultAggchainSelector, newDefaultAggchainVKey)
}

// UpdateDefaultAggchainVKey is a paid mutator transaction binding the contract method 0xf4c44681.
//
// Solidity: function updateDefaultAggchainVKey(bytes4 defaultAggchainSelector, bytes32 newDefaultAggchainVKey) returns()
func (_Agglayergateway *AgglayergatewayTransactorSession) UpdateDefaultAggchainVKey(defaultAggchainSelector [4]byte, newDefaultAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Agglayergateway.Contract.UpdateDefaultAggchainVKey(&_Agglayergateway.TransactOpts, defaultAggchainSelector, newDefaultAggchainVKey)
}

// AgglayergatewayAddDefaultAggchainVKeyIterator is returned from FilterAddDefaultAggchainVKey and is used to iterate over the raw logs and unpacked data for AddDefaultAggchainVKey events raised by the Agglayergateway contract.
type AgglayergatewayAddDefaultAggchainVKeyIterator struct {
	Event *AgglayergatewayAddDefaultAggchainVKey // Event containing the contract specifics and raw log

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
func (it *AgglayergatewayAddDefaultAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayergatewayAddDefaultAggchainVKey)
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
		it.Event = new(AgglayergatewayAddDefaultAggchainVKey)
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
func (it *AgglayergatewayAddDefaultAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayergatewayAddDefaultAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayergatewayAddDefaultAggchainVKey represents a AddDefaultAggchainVKey event raised by the Agglayergateway contract.
type AgglayergatewayAddDefaultAggchainVKey struct {
	Selector [4]byte
	NewVKey  [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddDefaultAggchainVKey is a free log retrieval operation binding the contract event 0x64b3528974574a41ddc830adc96440dccb99e64ede6b863194b951bfbd17e1bc.
//
// Solidity: event AddDefaultAggchainVKey(bytes4 selector, bytes32 newVKey)
func (_Agglayergateway *AgglayergatewayFilterer) FilterAddDefaultAggchainVKey(opts *bind.FilterOpts) (*AgglayergatewayAddDefaultAggchainVKeyIterator, error) {

	logs, sub, err := _Agglayergateway.contract.FilterLogs(opts, "AddDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &AgglayergatewayAddDefaultAggchainVKeyIterator{contract: _Agglayergateway.contract, event: "AddDefaultAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchAddDefaultAggchainVKey is a free log subscription operation binding the contract event 0x64b3528974574a41ddc830adc96440dccb99e64ede6b863194b951bfbd17e1bc.
//
// Solidity: event AddDefaultAggchainVKey(bytes4 selector, bytes32 newVKey)
func (_Agglayergateway *AgglayergatewayFilterer) WatchAddDefaultAggchainVKey(opts *bind.WatchOpts, sink chan<- *AgglayergatewayAddDefaultAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Agglayergateway.contract.WatchLogs(opts, "AddDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayergatewayAddDefaultAggchainVKey)
				if err := _Agglayergateway.contract.UnpackLog(event, "AddDefaultAggchainVKey", log); err != nil {
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
func (_Agglayergateway *AgglayergatewayFilterer) ParseAddDefaultAggchainVKey(log types.Log) (*AgglayergatewayAddDefaultAggchainVKey, error) {
	event := new(AgglayergatewayAddDefaultAggchainVKey)
	if err := _Agglayergateway.contract.UnpackLog(event, "AddDefaultAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayergatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Agglayergateway contract.
type AgglayergatewayInitializedIterator struct {
	Event *AgglayergatewayInitialized // Event containing the contract specifics and raw log

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
func (it *AgglayergatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayergatewayInitialized)
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
		it.Event = new(AgglayergatewayInitialized)
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
func (it *AgglayergatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayergatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayergatewayInitialized represents a Initialized event raised by the Agglayergateway contract.
type AgglayergatewayInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Agglayergateway *AgglayergatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*AgglayergatewayInitializedIterator, error) {

	logs, sub, err := _Agglayergateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AgglayergatewayInitializedIterator{contract: _Agglayergateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Agglayergateway *AgglayergatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AgglayergatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _Agglayergateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayergatewayInitialized)
				if err := _Agglayergateway.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Agglayergateway *AgglayergatewayFilterer) ParseInitialized(log types.Log) (*AgglayergatewayInitialized, error) {
	event := new(AgglayergatewayInitialized)
	if err := _Agglayergateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayergatewayRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Agglayergateway contract.
type AgglayergatewayRoleAdminChangedIterator struct {
	Event *AgglayergatewayRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AgglayergatewayRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayergatewayRoleAdminChanged)
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
		it.Event = new(AgglayergatewayRoleAdminChanged)
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
func (it *AgglayergatewayRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayergatewayRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayergatewayRoleAdminChanged represents a RoleAdminChanged event raised by the Agglayergateway contract.
type AgglayergatewayRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Agglayergateway *AgglayergatewayFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AgglayergatewayRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Agglayergateway.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AgglayergatewayRoleAdminChangedIterator{contract: _Agglayergateway.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Agglayergateway *AgglayergatewayFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AgglayergatewayRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Agglayergateway.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayergatewayRoleAdminChanged)
				if err := _Agglayergateway.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Agglayergateway *AgglayergatewayFilterer) ParseRoleAdminChanged(log types.Log) (*AgglayergatewayRoleAdminChanged, error) {
	event := new(AgglayergatewayRoleAdminChanged)
	if err := _Agglayergateway.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayergatewayRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Agglayergateway contract.
type AgglayergatewayRoleGrantedIterator struct {
	Event *AgglayergatewayRoleGranted // Event containing the contract specifics and raw log

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
func (it *AgglayergatewayRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayergatewayRoleGranted)
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
		it.Event = new(AgglayergatewayRoleGranted)
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
func (it *AgglayergatewayRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayergatewayRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayergatewayRoleGranted represents a RoleGranted event raised by the Agglayergateway contract.
type AgglayergatewayRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Agglayergateway *AgglayergatewayFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AgglayergatewayRoleGrantedIterator, error) {

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

	logs, sub, err := _Agglayergateway.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AgglayergatewayRoleGrantedIterator{contract: _Agglayergateway.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Agglayergateway *AgglayergatewayFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AgglayergatewayRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Agglayergateway.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayergatewayRoleGranted)
				if err := _Agglayergateway.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Agglayergateway *AgglayergatewayFilterer) ParseRoleGranted(log types.Log) (*AgglayergatewayRoleGranted, error) {
	event := new(AgglayergatewayRoleGranted)
	if err := _Agglayergateway.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayergatewayRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Agglayergateway contract.
type AgglayergatewayRoleRevokedIterator struct {
	Event *AgglayergatewayRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AgglayergatewayRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayergatewayRoleRevoked)
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
		it.Event = new(AgglayergatewayRoleRevoked)
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
func (it *AgglayergatewayRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayergatewayRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayergatewayRoleRevoked represents a RoleRevoked event raised by the Agglayergateway contract.
type AgglayergatewayRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Agglayergateway *AgglayergatewayFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AgglayergatewayRoleRevokedIterator, error) {

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

	logs, sub, err := _Agglayergateway.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AgglayergatewayRoleRevokedIterator{contract: _Agglayergateway.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Agglayergateway *AgglayergatewayFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AgglayergatewayRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Agglayergateway.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayergatewayRoleRevoked)
				if err := _Agglayergateway.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Agglayergateway *AgglayergatewayFilterer) ParseRoleRevoked(log types.Log) (*AgglayergatewayRoleRevoked, error) {
	event := new(AgglayergatewayRoleRevoked)
	if err := _Agglayergateway.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayergatewayRouteAddedIterator is returned from FilterRouteAdded and is used to iterate over the raw logs and unpacked data for RouteAdded events raised by the Agglayergateway contract.
type AgglayergatewayRouteAddedIterator struct {
	Event *AgglayergatewayRouteAdded // Event containing the contract specifics and raw log

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
func (it *AgglayergatewayRouteAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayergatewayRouteAdded)
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
		it.Event = new(AgglayergatewayRouteAdded)
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
func (it *AgglayergatewayRouteAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayergatewayRouteAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayergatewayRouteAdded represents a RouteAdded event raised by the Agglayergateway contract.
type AgglayergatewayRouteAdded struct {
	Selector        [4]byte
	Verifier        common.Address
	PessimisticVKey [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRouteAdded is a free log retrieval operation binding the contract event 0x4c4455a30528c319643ae6fc35f3a0fcabde6c015f2a1a5270382c4a190b0fa3.
//
// Solidity: event RouteAdded(bytes4 selector, address verifier, bytes32 pessimisticVKey)
func (_Agglayergateway *AgglayergatewayFilterer) FilterRouteAdded(opts *bind.FilterOpts) (*AgglayergatewayRouteAddedIterator, error) {

	logs, sub, err := _Agglayergateway.contract.FilterLogs(opts, "RouteAdded")
	if err != nil {
		return nil, err
	}
	return &AgglayergatewayRouteAddedIterator{contract: _Agglayergateway.contract, event: "RouteAdded", logs: logs, sub: sub}, nil
}

// WatchRouteAdded is a free log subscription operation binding the contract event 0x4c4455a30528c319643ae6fc35f3a0fcabde6c015f2a1a5270382c4a190b0fa3.
//
// Solidity: event RouteAdded(bytes4 selector, address verifier, bytes32 pessimisticVKey)
func (_Agglayergateway *AgglayergatewayFilterer) WatchRouteAdded(opts *bind.WatchOpts, sink chan<- *AgglayergatewayRouteAdded) (event.Subscription, error) {

	logs, sub, err := _Agglayergateway.contract.WatchLogs(opts, "RouteAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayergatewayRouteAdded)
				if err := _Agglayergateway.contract.UnpackLog(event, "RouteAdded", log); err != nil {
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
func (_Agglayergateway *AgglayergatewayFilterer) ParseRouteAdded(log types.Log) (*AgglayergatewayRouteAdded, error) {
	event := new(AgglayergatewayRouteAdded)
	if err := _Agglayergateway.contract.UnpackLog(event, "RouteAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayergatewayRouteFrozenIterator is returned from FilterRouteFrozen and is used to iterate over the raw logs and unpacked data for RouteFrozen events raised by the Agglayergateway contract.
type AgglayergatewayRouteFrozenIterator struct {
	Event *AgglayergatewayRouteFrozen // Event containing the contract specifics and raw log

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
func (it *AgglayergatewayRouteFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayergatewayRouteFrozen)
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
		it.Event = new(AgglayergatewayRouteFrozen)
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
func (it *AgglayergatewayRouteFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayergatewayRouteFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayergatewayRouteFrozen represents a RouteFrozen event raised by the Agglayergateway contract.
type AgglayergatewayRouteFrozen struct {
	Selector        [4]byte
	Verifier        common.Address
	PessimisticVKey [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRouteFrozen is a free log retrieval operation binding the contract event 0x6efcfd4090b998a9d58c7f753bdda3291e9d781c02226b9f480b18ef07e0012b.
//
// Solidity: event RouteFrozen(bytes4 selector, address verifier, bytes32 pessimisticVKey)
func (_Agglayergateway *AgglayergatewayFilterer) FilterRouteFrozen(opts *bind.FilterOpts) (*AgglayergatewayRouteFrozenIterator, error) {

	logs, sub, err := _Agglayergateway.contract.FilterLogs(opts, "RouteFrozen")
	if err != nil {
		return nil, err
	}
	return &AgglayergatewayRouteFrozenIterator{contract: _Agglayergateway.contract, event: "RouteFrozen", logs: logs, sub: sub}, nil
}

// WatchRouteFrozen is a free log subscription operation binding the contract event 0x6efcfd4090b998a9d58c7f753bdda3291e9d781c02226b9f480b18ef07e0012b.
//
// Solidity: event RouteFrozen(bytes4 selector, address verifier, bytes32 pessimisticVKey)
func (_Agglayergateway *AgglayergatewayFilterer) WatchRouteFrozen(opts *bind.WatchOpts, sink chan<- *AgglayergatewayRouteFrozen) (event.Subscription, error) {

	logs, sub, err := _Agglayergateway.contract.WatchLogs(opts, "RouteFrozen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayergatewayRouteFrozen)
				if err := _Agglayergateway.contract.UnpackLog(event, "RouteFrozen", log); err != nil {
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
func (_Agglayergateway *AgglayergatewayFilterer) ParseRouteFrozen(log types.Log) (*AgglayergatewayRouteFrozen, error) {
	event := new(AgglayergatewayRouteFrozen)
	if err := _Agglayergateway.contract.UnpackLog(event, "RouteFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayergatewayUpdateDefaultAggchainVKeyIterator is returned from FilterUpdateDefaultAggchainVKey and is used to iterate over the raw logs and unpacked data for UpdateDefaultAggchainVKey events raised by the Agglayergateway contract.
type AgglayergatewayUpdateDefaultAggchainVKeyIterator struct {
	Event *AgglayergatewayUpdateDefaultAggchainVKey // Event containing the contract specifics and raw log

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
func (it *AgglayergatewayUpdateDefaultAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayergatewayUpdateDefaultAggchainVKey)
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
		it.Event = new(AgglayergatewayUpdateDefaultAggchainVKey)
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
func (it *AgglayergatewayUpdateDefaultAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayergatewayUpdateDefaultAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayergatewayUpdateDefaultAggchainVKey represents a UpdateDefaultAggchainVKey event raised by the Agglayergateway contract.
type AgglayergatewayUpdateDefaultAggchainVKey struct {
	Selector     [4]byte
	PreviousVKey [32]byte
	NewVKey      [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateDefaultAggchainVKey is a free log retrieval operation binding the contract event 0x477ac863279848383fdb5d6bb8957cc8f6f10c0200379e075df3f789d14248e6.
//
// Solidity: event UpdateDefaultAggchainVKey(bytes4 selector, bytes32 previousVKey, bytes32 newVKey)
func (_Agglayergateway *AgglayergatewayFilterer) FilterUpdateDefaultAggchainVKey(opts *bind.FilterOpts) (*AgglayergatewayUpdateDefaultAggchainVKeyIterator, error) {

	logs, sub, err := _Agglayergateway.contract.FilterLogs(opts, "UpdateDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &AgglayergatewayUpdateDefaultAggchainVKeyIterator{contract: _Agglayergateway.contract, event: "UpdateDefaultAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchUpdateDefaultAggchainVKey is a free log subscription operation binding the contract event 0x477ac863279848383fdb5d6bb8957cc8f6f10c0200379e075df3f789d14248e6.
//
// Solidity: event UpdateDefaultAggchainVKey(bytes4 selector, bytes32 previousVKey, bytes32 newVKey)
func (_Agglayergateway *AgglayergatewayFilterer) WatchUpdateDefaultAggchainVKey(opts *bind.WatchOpts, sink chan<- *AgglayergatewayUpdateDefaultAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Agglayergateway.contract.WatchLogs(opts, "UpdateDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayergatewayUpdateDefaultAggchainVKey)
				if err := _Agglayergateway.contract.UnpackLog(event, "UpdateDefaultAggchainVKey", log); err != nil {
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
func (_Agglayergateway *AgglayergatewayFilterer) ParseUpdateDefaultAggchainVKey(log types.Log) (*AgglayergatewayUpdateDefaultAggchainVKey, error) {
	event := new(AgglayergatewayUpdateDefaultAggchainVKey)
	if err := _Agglayergateway.contract.UnpackLog(event, "UpdateDefaultAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

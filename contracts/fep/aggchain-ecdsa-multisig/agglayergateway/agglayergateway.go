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

// IAggLayerGatewayRemoveSignerInfo is an auto generated low-level Go binding around an user-defined struct.
type IAggLayerGatewayRemoveSignerInfo struct {
	Addr  common.Address
	Index *big.Int
}

// IAggLayerGatewaySignerInfo is an auto generated low-level Go binding around an user-defined struct.
type IAggLayerGatewaySignerInfo struct {
	Addr common.Address
	Url  string
}

// AgglayergatewayMetaData contains all meta data concerning the Agglayergateway contract.
var AgglayergatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainSignersHashNotInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainSignersTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainVKeyAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndicesNotInDescendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProofBytesLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAggLayerAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAggLayerAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PPSelectorCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"RouteAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"RouteIsAlreadyFrozen\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"RouteIsFrozen\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"RouteNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerURLCannotBeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VKeyCannotBeZero\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newVKey\",\"type\":\"bytes32\"}],\"name\":\"AddDefaultAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pessimisticVKey\",\"type\":\"bytes32\"}],\"name\":\"RouteAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pessimisticVKey\",\"type\":\"bytes32\"}],\"name\":\"RouteFrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signersHash\",\"type\":\"bytes32\"}],\"name\":\"SignersAndThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"UnsetDefaultAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousVKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newVKey\",\"type\":\"bytes32\"}],\"name\":\"UpdateDefaultAggchainVKey\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AGGLAYER_GATEWAY_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"defaultAggchainSelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"addDefaultAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"pessimisticVKeySelector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"pessimisticVKey\",\"type\":\"bytes32\"}],\"name\":\"addPessimisticVKeyRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"aggchainSigners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggchainSignersHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"defaultAggchainSelector\",\"type\":\"bytes4\"}],\"name\":\"defaultAggchainVKeys\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"defaultAggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"pessimisticVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"freezePessimisticVKeyRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggchainSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggchainSignersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggchainSignersHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"defaultAggchainSelector\",\"type\":\"bytes4\"}],\"name\":\"getDefaultAggchainVKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"multisigRole\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structIAggLayerGateway.SignerInfo[]\",\"name\":\"signersToAdd\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"aggchainDefaultVKeyRole\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addRouteRole\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"freezeRouteRole\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"pessimisticVKeySelector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"pessimisticVKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"multisigRole\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structIAggLayerGateway.SignerInfo[]\",\"name\":\"signersToAdd\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"pessimisticVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"pessimisticVKeyRoutes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"pessimisticVKey\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signerToURLs\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"defaultAggchainSelector\",\"type\":\"bytes4\"}],\"name\":\"unsetDefaultAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"defaultAggchainSelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"newDefaultAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"updateDefaultAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structIAggLayerGateway.RemoveSignerInfo[]\",\"name\":\"_signersToRemove\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structIAggLayerGateway.SignerInfo[]\",\"name\":\"_signersToAdd\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_newThreshold\",\"type\":\"uint256\"}],\"name\":\"updateSignersAndThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicValues\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofBytes\",\"type\":\"bytes\"}],\"name\":\"verifyPessimisticProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060156019565b60c9565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560685760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c65780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b612cff806100d65f395ff3fe608060405234801561000f575f5ffd5b50600436106101c6575f3560e01c8063716be075116100fe578063a217fddf1161009e578063d356e5771161006e578063d356e5771461050c578063d547741f1461051f578063f4c4468114610532578063f51f563a14610545575f5ffd5b8063a217fddf146104ae578063a48fd377146104b5578063c49196ec146104c8578063ca69e7dc14610504575f5ffd5b806382bfaea1116100d957806382bfaea11461040557806391d14854146104245780639584a5161461048857806395d2a3911461049b575f5ffd5b8063716be075146103625780637df73e27146103df578063818c8c21146103f2575f5ffd5b806336568abe1161016957806342cde4e81161014457806342cde4e8146102fa57806347c63d361461030357806354fd4d50146103165780636cabfdab1461034f575f5ffd5b806336568abe146102b257806336cd6b5b146102c55780633e1e0121146102e5575f5ffd5b806317b7a9f0116101a457806317b7a9f01461021d578063248a9ca3146102265780632f2ff15d1461026757806335acd6c21461027a575f5ffd5b806301ffc9a7146101ca5780630c7dd8b3146101f257806312b941ca14610207575b5f5ffd5b6101dd6101d836600461221d565b610558565b60405190151581526020015b60405180910390f35b6102056102003660046124a0565b6105f0565b005b61020f61087b565b6040519081526020016101e9565b61020f60055481565b61020f6102343660046124f4565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b61020561027536600461250b565b6108bd565b61028d6102883660046124f4565b610906565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101e9565b6102056102c036600461250b565b61093b565b6102d86102d3366004612535565b610999565b6040516101e9919061254e565b6102ed610a30565b6040516101e991906125a1565b61020f60045481565b61020561031136600461221d565b610a9d565b60408051808201909152600681527f76312e302e30000000000000000000000000000000000000000000000000000060208201526102d8565b61020f61035d36600461221d565b610b97565b6103ab61037036600461221d565b600160208190525f918252604090912080549181015460029091015473ffffffffffffffffffffffffffffffffffffffff9092169160ff1683565b6040805173ffffffffffffffffffffffffffffffffffffffff909416845260208401929092521515908201526060016101e9565b6101dd6103ed366004612535565b610c32565b6102056104003660046125f9565b610c6e565b61020f61041336600461221d565b5f6020819052908152604090205481565b6101dd61043236600461250b565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b610205610496366004612621565b610da7565b6102056104a936600461221d565b610ddc565b61020f5f81565b6102056104c33660046126a0565b610fb9565b6102d86040518060400160405280600681526020017f76312e302e30000000000000000000000000000000000000000000000000000081525081565b60025461020f565b61020561051a36600461270c565b6111c4565b61020561052d36600461250b565b61153b565b6102056105403660046125f9565b61157e565b6102056105533660046127d1565b6116c4565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806105ea57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005467ffffffffffffffff165f805c7fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff831617905d507ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0080546002919068010000000000000000900460ff168061069f5750805467ffffffffffffffff808416911610155b156106d6576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80547fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000001667ffffffffffffffff80841691909117680100000000000000001782555f5c16600114610753576040517fadc06ae700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff85166107a0576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107ca7f0c3038a1ecdf82843b70709289ff1703351ad391e3e27df7f6fa7d913601e15e866116f9565b50604080515f8082526020820190925261081191610809565b604080518082019091525f80825260208201528152602001906001900390816107e35790505b50858561181e565b80547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16815560405167ffffffffffffffff831681527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15050505050565b6005545f906108b6576040517fdd41f1ef00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5060055490565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546108f6816119fa565b61090083836116f9565b50505050565b60028181548110610915575f80fd5b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b73ffffffffffffffffffffffffffffffffffffffff8116331461098a576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109948282611a07565b505050565b60036020525f9081526040902080546109b1906128a4565b80601f01602080910402602001604051908101604052809291908181526020018280546109dd906128a4565b8015610a285780601f106109ff57610100808354040283529160200191610a28565b820191905f5260205f20905b815481529060010190602001808311610a0b57829003601f168201915b505050505081565b60606002805480602002602001604051908101604052809291908181526020018280548015610a9357602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610a68575b5050505050905090565b7f131410eab1236cee2db19035b0e825c94e5ab705dffe23321dd53856da531617610ac7816119fa565b7fffffffff0000000000000000000000000000000000000000000000000000000082165f90815260208190526040902054610b2e576040517f925e5a3a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f818152602081815260408083209290925590519182527f43b46d79ee2fbaa8a981b5fd78aaf6fb9e38b18d02e1af17564ce37c67078cab910160405180910390a15050565b7fffffffff0000000000000000000000000000000000000000000000000000000081165f90815260208190526040812054610bfe576040517f925e5a3a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b507fffffffff00000000000000000000000000000000000000000000000000000000165f9081526020819052604090205490565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526003602052604081208054829190610c64906128a4565b9050119050919050565b7f131410eab1236cee2db19035b0e825c94e5ab705dffe23321dd53856da531617610c98816119fa565b7fffffffff0000000000000000000000000000000000000000000000000000000083165f9081526020819052604090205415610d00576040517f22a1bdc400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81610d37576040517f6745305e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000083165f8181526020818152604091829020859055815192835282018490527f64b3528974574a41ddc830adc96440dccb99e64ede6b863194b951bfbd17e1bc91015b60405180910390a1505050565b7f0fdc2a718b96bc741c7544001e3dd7c26730802c54781668fa78a120e622629b610dd1816119fa565b610900848484611ae3565b7fca75ae4228cde6195f9fa3dbde8dc352fb30aa63780717a378ccfc50274355dd610e06816119fa565b7fffffffff0000000000000000000000000000000000000000000000000000000082165f908152600160205260409020805473ffffffffffffffffffffffffffffffffffffffff16610ead576040517ff208777e0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000841660048201526024015b60405180910390fd5b600281015460ff1615610f10576040517fc7b4a1f20000000000000000000000000000000000000000000000000000000081527fffffffff0000000000000000000000000000000000000000000000000000000084166004820152602401610ea4565b6002810180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001908117909155815490820154604080517fffffffff000000000000000000000000000000000000000000000000000000008716815273ffffffffffffffffffffffffffffffffffffffff90931660208401528201527f6efcfd4090b998a9d58c7f753bdda3291e9d781c02226b9f480b18ef07e0012b90606001610d9a565b6004811015610ff4576040517fae80851c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61100260048284866128f5565b61100b9161291c565b7fffffffff0000000000000000000000000000000000000000000000000000000081165f908152600160208181526040928390208351606081018552815473ffffffffffffffffffffffffffffffffffffffff1680825293820154928101929092526002015460ff16151592810192909252919250906110db576040517ff208777e0000000000000000000000000000000000000000000000000000000081527fffffffff0000000000000000000000000000000000000000000000000000000083166004820152602401610ea4565b80604001511561113b576040517febf108230000000000000000000000000000000000000000000000000000000081527fffffffff0000000000000000000000000000000000000000000000000000000083166004820152602401610ea4565b8051602082015173ffffffffffffffffffffffffffffffffffffffff909116906341493c60908888611170886004818c6128f5565b6040518663ffffffff1660e01b81526004016111909594939291906129c8565b5f6040518083038186803b1580156111a6575f5ffd5b505afa1580156111b8573d5f5f3e3d5ffd5b50505050505050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005467ffffffffffffffff165f805c7fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff831617905d507ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0080546002919068010000000000000000900460ff16806112735750805467ffffffffffffffff808416911610155b156112aa576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80547fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000001667ffffffffffffffff80841691909117680100000000000000001782555f5c1615611325576040517fadc06ae700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8516158061135c575073ffffffffffffffffffffffffffffffffffffffff8c16155b8061137b575073ffffffffffffffffffffffffffffffffffffffff8b16155b8061139a575073ffffffffffffffffffffffffffffffffffffffff8a16155b806113b9575073ffffffffffffffffffffffffffffffffffffffff8916155b156113f0576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113fa5f8d6116f9565b506114257f131410eab1236cee2db19035b0e825c94e5ab705dffe23321dd53856da5316178c6116f9565b506114507f0fdc2a718b96bc741c7544001e3dd7c26730802c54781668fa78a120e622629b8b6116f9565b5061147b7fca75ae4228cde6195f9fa3dbde8dc352fb30aa63780717a378ccfc50274355dd8a6116f9565b506114a67f0c3038a1ecdf82843b70709289ff1703351ad391e3e27df7f6fa7d913601e15e866116f9565b506114b2888888611ae3565b604080515f808252602082019092526114ca91610809565b80547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16815560405167ffffffffffffffff831681527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a1505050505050505050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154611574816119fa565b6109008383611a07565b7f131410eab1236cee2db19035b0e825c94e5ab705dffe23321dd53856da5316176115a8816119fa565b7fffffffff0000000000000000000000000000000000000000000000000000000083165f9081526020819052604090205461160f576040517f925e5a3a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81611646576040517f6745305e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000083165f81815260208181526040918290208054908690558251938452908301819052908201849052907f477ac863279848383fdb5d6bb8957cc8f6f10c0200379e075df3f789d14248e6906060015b60405180910390a150505050565b7f0c3038a1ecdf82843b70709289ff1703351ad391e3e27df7f6fa7d913601e15e6116ee816119fa565b61090084848461181e565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff1661180e575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556117aa3390565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506105ea565b5f9150506105ea565b5092915050565b6001835111156118c6575f5b600184516118389190612a2d565b8110156118c4578361184b826001612a40565b8151811061185b5761185b612a53565b60200260200101516020015184828151811061187957611879612a53565b602002602001015160200151116118bc576040517fb9a11d3100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60010161182a565b505b5f5b835181101561191c576119148482815181106118e6576118e6612a53565b60200260200101515f015185838151811061190357611903612a53565b602002602001015160200151611d29565b6001016118c8565b505f5b82518110156119735761196b83828151811061193d5761193d612a53565b60200260200101515f015184838151811061195a5761195a612a53565b602002602001015160200151611f1b565b60010161191f565b5060025460ff10156119b1576040517f5a7f382c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002548111156119ed576040517faabd5a0900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6004819055610994612067565b611a0481336120f5565b50565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff161561180e575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506105ea565b73ffffffffffffffffffffffffffffffffffffffff8216611b30576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff000000000000000000000000000000000000000000000000000000008316611b89576040517f7500f20900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80611bc0576040517f6745305e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000083165f908152600160205260409020805473ffffffffffffffffffffffffffffffffffffffff1615611c825780546040517f4bc7c2780000000000000000000000000000000000000000000000000000000081527fffffffff000000000000000000000000000000000000000000000000000000008616600482015273ffffffffffffffffffffffffffffffffffffffff9091166024820152604401610ea4565b80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8416908117825560018201839055604080517fffffffff0000000000000000000000000000000000000000000000000000000087168152602081019290925281018390527f4c4455a30528c319643ae6fc35f3a0fcabde6c015f2a1a5270382c4a190b0fa3906060016116b6565b600254808210611d65576040517fd244b30700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff1660028381548110611d8f57611d8f612a53565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff1614611de7576040517fd244b30700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83165f908152600360205260408120611e149161219f565b6002611e21600183612a2d565b81548110611e3157611e31612a53565b5f918252602090912001546002805473ffffffffffffffffffffffffffffffffffffffff9092169184908110611e6957611e69612a53565b905f5260205f20015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506002805480611ebf57611ebf612a80565b5f8281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055505050565b73ffffffffffffffffffffffffffffffffffffffff8216611f68576040517f7b3a0df600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80515f03611fa2576040517f8715f5fb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611fab82610c32565b15611fe2576040517f38615ecc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028054600181019091557f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84169081179091555f9081526003602052604090206109948282612af8565b600454600260405160200161207d929190612c0f565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052805160209091012060058190556004547f66d7b0647fdd512b69cbf4f8e1ce8068bfe0b236168e2704ba13b07425eaa743926120eb9260029291612c62565b60405180910390a1565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1661219b576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610ea4565b5050565b5080546121ab906128a4565b5f825580601f106121ba575050565b601f0160209004905f5260205f2090810190611a0491905b808211156121e5575f81556001016121d2565b5090565b80357fffffffff0000000000000000000000000000000000000000000000000000000081168114612218575f5ffd5b919050565b5f6020828403121561222d575f5ffd5b612236826121e9565b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114612218575f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6040805190810167ffffffffffffffff811182821017156122b0576122b0612260565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156122fd576122fd612260565b604052919050565b5f67ffffffffffffffff82111561231e5761231e612260565b5060051b60200190565b5f82601f830112612337575f5ffd5b813561234a61234582612305565b6122b6565b8082825260208201915060208360051b86010192508583111561236b575f5ffd5b602085015b8381101561249657803567ffffffffffffffff81111561238e575f5ffd5b860160408189037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe00112156123c1575f5ffd5b6123c961228d565b6123d56020830161223d565b8152604082013567ffffffffffffffff8111156123f0575f5ffd5b60208184010192505088601f830112612407575f5ffd5b813567ffffffffffffffff81111561242157612421612260565b61245260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016122b6565b8181528a6020838601011115612466575f5ffd5b816020850160208301375f6020838301015280602084015250508085525050602083019250602081019050612370565b5095945050505050565b5f5f5f606084860312156124b2575f5ffd5b6124bb8461223d565b9250602084013567ffffffffffffffff8111156124d6575f5ffd5b6124e286828701612328565b93969395505050506040919091013590565b5f60208284031215612504575f5ffd5b5035919050565b5f5f6040838503121561251c575f5ffd5b8235915061252c6020840161223d565b90509250929050565b5f60208284031215612545575f5ffd5b6122368261223d565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b602080825282518282018190525f918401906040840190835b818110156125ee57835173ffffffffffffffffffffffffffffffffffffffff168352602093840193909201916001016125ba565b509095945050505050565b5f5f6040838503121561260a575f5ffd5b612613836121e9565b946020939093013593505050565b5f5f5f60608486031215612633575f5ffd5b61263c846121e9565b925061264a6020850161223d565b929592945050506040919091013590565b5f5f83601f84011261266b575f5ffd5b50813567ffffffffffffffff811115612682575f5ffd5b602083019150836020828501011115612699575f5ffd5b9250929050565b5f5f5f5f604085870312156126b3575f5ffd5b843567ffffffffffffffff8111156126c9575f5ffd5b6126d58782880161265b565b909550935050602085013567ffffffffffffffff8111156126f4575f5ffd5b6127008782880161265b565b95989497509550505050565b5f5f5f5f5f5f5f5f5f5f6101408b8d031215612726575f5ffd5b61272f8b61223d565b995061273d60208c0161223d565b985061274b60408c0161223d565b975061275960608c0161223d565b965061276760808c016121e9565b955061277560a08c0161223d565b945060c08b0135935061278a60e08c0161223d565b92506101008b013567ffffffffffffffff8111156127a6575f5ffd5b6127b28d828e01612328565b9250505f6101208c01359050809150509295989b9194979a5092959850565b5f5f5f606084860312156127e3575f5ffd5b833567ffffffffffffffff8111156127f9575f5ffd5b8401601f81018613612809575f5ffd5b803561281761234582612305565b8082825260208201915060208360061b850101925088831115612838575f5ffd5b6020840193505b82841015612886576040848a031215612856575f5ffd5b61285e61228d565b6128678561223d565b815260208581013581830152908352604090940193919091019061283f565b9550505050602084013567ffffffffffffffff8111156124d6575f5ffd5b600181811c908216806128b857607f821691505b6020821081036128ef577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b5f5f85851115612903575f5ffd5b8386111561290f575f5ffd5b5050820193919092039150565b80357fffffffff000000000000000000000000000000000000000000000000000000008116906004841015611817577fffffffff00000000000000000000000000000000000000000000000000000000808560040360031b1b82161691505092915050565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b858152606060208201525f6129e1606083018688612981565b82810360408401526129f4818587612981565b98975050505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b818103818111156105ea576105ea612a00565b808201808211156105ea576105ea612a00565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b601f82111561099457805f5260205f20601f840160051c81016020851015612ad25750805b601f840160051c820191505b81811015612af1575f8155600101612ade565b5050505050565b815167ffffffffffffffff811115612b1257612b12612260565b612b2681612b2084546128a4565b84612aad565b6020601f821160018114612b77575f8315612b415750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455612af1565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015612bc45787850151825560209485019460019092019101612ba4565b5084821015612c0057868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b8281525f602082018354845f5260205f205f5b82811015612c5657815473ffffffffffffffffffffffffffffffffffffffff16845260209093019260019182019101612c22565b50919695505050505050565b606080825284549082018190525f8581526020812090916080840190835b81811015612cb457835473ffffffffffffffffffffffffffffffffffffffff16835260019384019360209093019201612c80565b5050602084019590955250506040015291905056fea2646970667358221220dee7a73b681379c4e1a9633ebaac5994eb58ac33f50f345735adccec2ab7ff2a64736f6c634300081c0033",
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

// AGGLAYERGATEWAYVERSION is a free data retrieval call binding the contract method 0xc49196ec.
//
// Solidity: function AGGLAYER_GATEWAY_VERSION() view returns(string)
func (_Agglayergateway *AgglayergatewayCaller) AGGLAYERGATEWAYVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Agglayergateway.contract.Call(opts, &out, "AGGLAYER_GATEWAY_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AGGLAYERGATEWAYVERSION is a free data retrieval call binding the contract method 0xc49196ec.
//
// Solidity: function AGGLAYER_GATEWAY_VERSION() view returns(string)
func (_Agglayergateway *AgglayergatewaySession) AGGLAYERGATEWAYVERSION() (string, error) {
	return _Agglayergateway.Contract.AGGLAYERGATEWAYVERSION(&_Agglayergateway.CallOpts)
}

// AGGLAYERGATEWAYVERSION is a free data retrieval call binding the contract method 0xc49196ec.
//
// Solidity: function AGGLAYER_GATEWAY_VERSION() view returns(string)
func (_Agglayergateway *AgglayergatewayCallerSession) AGGLAYERGATEWAYVERSION() (string, error) {
	return _Agglayergateway.Contract.AGGLAYERGATEWAYVERSION(&_Agglayergateway.CallOpts)
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

// AggchainSigners is a free data retrieval call binding the contract method 0x35acd6c2.
//
// Solidity: function aggchainSigners(uint256 ) view returns(address)
func (_Agglayergateway *AgglayergatewayCaller) AggchainSigners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Agglayergateway.contract.Call(opts, &out, "aggchainSigners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggchainSigners is a free data retrieval call binding the contract method 0x35acd6c2.
//
// Solidity: function aggchainSigners(uint256 ) view returns(address)
func (_Agglayergateway *AgglayergatewaySession) AggchainSigners(arg0 *big.Int) (common.Address, error) {
	return _Agglayergateway.Contract.AggchainSigners(&_Agglayergateway.CallOpts, arg0)
}

// AggchainSigners is a free data retrieval call binding the contract method 0x35acd6c2.
//
// Solidity: function aggchainSigners(uint256 ) view returns(address)
func (_Agglayergateway *AgglayergatewayCallerSession) AggchainSigners(arg0 *big.Int) (common.Address, error) {
	return _Agglayergateway.Contract.AggchainSigners(&_Agglayergateway.CallOpts, arg0)
}

// AggchainSignersHash is a free data retrieval call binding the contract method 0x17b7a9f0.
//
// Solidity: function aggchainSignersHash() view returns(bytes32)
func (_Agglayergateway *AgglayergatewayCaller) AggchainSignersHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Agglayergateway.contract.Call(opts, &out, "aggchainSignersHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AggchainSignersHash is a free data retrieval call binding the contract method 0x17b7a9f0.
//
// Solidity: function aggchainSignersHash() view returns(bytes32)
func (_Agglayergateway *AgglayergatewaySession) AggchainSignersHash() ([32]byte, error) {
	return _Agglayergateway.Contract.AggchainSignersHash(&_Agglayergateway.CallOpts)
}

// AggchainSignersHash is a free data retrieval call binding the contract method 0x17b7a9f0.
//
// Solidity: function aggchainSignersHash() view returns(bytes32)
func (_Agglayergateway *AgglayergatewayCallerSession) AggchainSignersHash() ([32]byte, error) {
	return _Agglayergateway.Contract.AggchainSignersHash(&_Agglayergateway.CallOpts)
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

// GetAggchainSigners is a free data retrieval call binding the contract method 0x3e1e0121.
//
// Solidity: function getAggchainSigners() view returns(address[])
func (_Agglayergateway *AgglayergatewayCaller) GetAggchainSigners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Agglayergateway.contract.Call(opts, &out, "getAggchainSigners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAggchainSigners is a free data retrieval call binding the contract method 0x3e1e0121.
//
// Solidity: function getAggchainSigners() view returns(address[])
func (_Agglayergateway *AgglayergatewaySession) GetAggchainSigners() ([]common.Address, error) {
	return _Agglayergateway.Contract.GetAggchainSigners(&_Agglayergateway.CallOpts)
}

// GetAggchainSigners is a free data retrieval call binding the contract method 0x3e1e0121.
//
// Solidity: function getAggchainSigners() view returns(address[])
func (_Agglayergateway *AgglayergatewayCallerSession) GetAggchainSigners() ([]common.Address, error) {
	return _Agglayergateway.Contract.GetAggchainSigners(&_Agglayergateway.CallOpts)
}

// GetAggchainSignersCount is a free data retrieval call binding the contract method 0xca69e7dc.
//
// Solidity: function getAggchainSignersCount() view returns(uint256)
func (_Agglayergateway *AgglayergatewayCaller) GetAggchainSignersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Agglayergateway.contract.Call(opts, &out, "getAggchainSignersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAggchainSignersCount is a free data retrieval call binding the contract method 0xca69e7dc.
//
// Solidity: function getAggchainSignersCount() view returns(uint256)
func (_Agglayergateway *AgglayergatewaySession) GetAggchainSignersCount() (*big.Int, error) {
	return _Agglayergateway.Contract.GetAggchainSignersCount(&_Agglayergateway.CallOpts)
}

// GetAggchainSignersCount is a free data retrieval call binding the contract method 0xca69e7dc.
//
// Solidity: function getAggchainSignersCount() view returns(uint256)
func (_Agglayergateway *AgglayergatewayCallerSession) GetAggchainSignersCount() (*big.Int, error) {
	return _Agglayergateway.Contract.GetAggchainSignersCount(&_Agglayergateway.CallOpts)
}

// GetAggchainSignersHash is a free data retrieval call binding the contract method 0x12b941ca.
//
// Solidity: function getAggchainSignersHash() view returns(bytes32)
func (_Agglayergateway *AgglayergatewayCaller) GetAggchainSignersHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Agglayergateway.contract.Call(opts, &out, "getAggchainSignersHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAggchainSignersHash is a free data retrieval call binding the contract method 0x12b941ca.
//
// Solidity: function getAggchainSignersHash() view returns(bytes32)
func (_Agglayergateway *AgglayergatewaySession) GetAggchainSignersHash() ([32]byte, error) {
	return _Agglayergateway.Contract.GetAggchainSignersHash(&_Agglayergateway.CallOpts)
}

// GetAggchainSignersHash is a free data retrieval call binding the contract method 0x12b941ca.
//
// Solidity: function getAggchainSignersHash() view returns(bytes32)
func (_Agglayergateway *AgglayergatewayCallerSession) GetAggchainSignersHash() ([32]byte, error) {
	return _Agglayergateway.Contract.GetAggchainSignersHash(&_Agglayergateway.CallOpts)
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

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address _signer) view returns(bool)
func (_Agglayergateway *AgglayergatewayCaller) IsSigner(opts *bind.CallOpts, _signer common.Address) (bool, error) {
	var out []interface{}
	err := _Agglayergateway.contract.Call(opts, &out, "isSigner", _signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address _signer) view returns(bool)
func (_Agglayergateway *AgglayergatewaySession) IsSigner(_signer common.Address) (bool, error) {
	return _Agglayergateway.Contract.IsSigner(&_Agglayergateway.CallOpts, _signer)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address _signer) view returns(bool)
func (_Agglayergateway *AgglayergatewayCallerSession) IsSigner(_signer common.Address) (bool, error) {
	return _Agglayergateway.Contract.IsSigner(&_Agglayergateway.CallOpts, _signer)
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

// SignerToURLs is a free data retrieval call binding the contract method 0x36cd6b5b.
//
// Solidity: function signerToURLs(address ) view returns(string)
func (_Agglayergateway *AgglayergatewayCaller) SignerToURLs(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _Agglayergateway.contract.Call(opts, &out, "signerToURLs", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SignerToURLs is a free data retrieval call binding the contract method 0x36cd6b5b.
//
// Solidity: function signerToURLs(address ) view returns(string)
func (_Agglayergateway *AgglayergatewaySession) SignerToURLs(arg0 common.Address) (string, error) {
	return _Agglayergateway.Contract.SignerToURLs(&_Agglayergateway.CallOpts, arg0)
}

// SignerToURLs is a free data retrieval call binding the contract method 0x36cd6b5b.
//
// Solidity: function signerToURLs(address ) view returns(string)
func (_Agglayergateway *AgglayergatewayCallerSession) SignerToURLs(arg0 common.Address) (string, error) {
	return _Agglayergateway.Contract.SignerToURLs(&_Agglayergateway.CallOpts, arg0)
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

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_Agglayergateway *AgglayergatewayCaller) Threshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Agglayergateway.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_Agglayergateway *AgglayergatewaySession) Threshold() (*big.Int, error) {
	return _Agglayergateway.Contract.Threshold(&_Agglayergateway.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_Agglayergateway *AgglayergatewayCallerSession) Threshold() (*big.Int, error) {
	return _Agglayergateway.Contract.Threshold(&_Agglayergateway.CallOpts)
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

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Agglayergateway *AgglayergatewayCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Agglayergateway.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Agglayergateway *AgglayergatewaySession) Version() (string, error) {
	return _Agglayergateway.Contract.Version(&_Agglayergateway.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Agglayergateway *AgglayergatewayCallerSession) Version() (string, error) {
	return _Agglayergateway.Contract.Version(&_Agglayergateway.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0x0c7dd8b3.
//
// Solidity: function initialize(address multisigRole, (address,string)[] signersToAdd, uint256 newThreshold) returns()
func (_Agglayergateway *AgglayergatewayTransactor) Initialize(opts *bind.TransactOpts, multisigRole common.Address, signersToAdd []IAggLayerGatewaySignerInfo, newThreshold *big.Int) (*types.Transaction, error) {
	return _Agglayergateway.contract.Transact(opts, "initialize", multisigRole, signersToAdd, newThreshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x0c7dd8b3.
//
// Solidity: function initialize(address multisigRole, (address,string)[] signersToAdd, uint256 newThreshold) returns()
func (_Agglayergateway *AgglayergatewaySession) Initialize(multisigRole common.Address, signersToAdd []IAggLayerGatewaySignerInfo, newThreshold *big.Int) (*types.Transaction, error) {
	return _Agglayergateway.Contract.Initialize(&_Agglayergateway.TransactOpts, multisigRole, signersToAdd, newThreshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x0c7dd8b3.
//
// Solidity: function initialize(address multisigRole, (address,string)[] signersToAdd, uint256 newThreshold) returns()
func (_Agglayergateway *AgglayergatewayTransactorSession) Initialize(multisigRole common.Address, signersToAdd []IAggLayerGatewaySignerInfo, newThreshold *big.Int) (*types.Transaction, error) {
	return _Agglayergateway.Contract.Initialize(&_Agglayergateway.TransactOpts, multisigRole, signersToAdd, newThreshold)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xd356e577.
//
// Solidity: function initialize(address defaultAdmin, address aggchainDefaultVKeyRole, address addRouteRole, address freezeRouteRole, bytes4 pessimisticVKeySelector, address verifier, bytes32 pessimisticVKey, address multisigRole, (address,string)[] signersToAdd, uint256 newThreshold) returns()
func (_Agglayergateway *AgglayergatewayTransactor) Initialize0(opts *bind.TransactOpts, defaultAdmin common.Address, aggchainDefaultVKeyRole common.Address, addRouteRole common.Address, freezeRouteRole common.Address, pessimisticVKeySelector [4]byte, verifier common.Address, pessimisticVKey [32]byte, multisigRole common.Address, signersToAdd []IAggLayerGatewaySignerInfo, newThreshold *big.Int) (*types.Transaction, error) {
	return _Agglayergateway.contract.Transact(opts, "initialize0", defaultAdmin, aggchainDefaultVKeyRole, addRouteRole, freezeRouteRole, pessimisticVKeySelector, verifier, pessimisticVKey, multisigRole, signersToAdd, newThreshold)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xd356e577.
//
// Solidity: function initialize(address defaultAdmin, address aggchainDefaultVKeyRole, address addRouteRole, address freezeRouteRole, bytes4 pessimisticVKeySelector, address verifier, bytes32 pessimisticVKey, address multisigRole, (address,string)[] signersToAdd, uint256 newThreshold) returns()
func (_Agglayergateway *AgglayergatewaySession) Initialize0(defaultAdmin common.Address, aggchainDefaultVKeyRole common.Address, addRouteRole common.Address, freezeRouteRole common.Address, pessimisticVKeySelector [4]byte, verifier common.Address, pessimisticVKey [32]byte, multisigRole common.Address, signersToAdd []IAggLayerGatewaySignerInfo, newThreshold *big.Int) (*types.Transaction, error) {
	return _Agglayergateway.Contract.Initialize0(&_Agglayergateway.TransactOpts, defaultAdmin, aggchainDefaultVKeyRole, addRouteRole, freezeRouteRole, pessimisticVKeySelector, verifier, pessimisticVKey, multisigRole, signersToAdd, newThreshold)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xd356e577.
//
// Solidity: function initialize(address defaultAdmin, address aggchainDefaultVKeyRole, address addRouteRole, address freezeRouteRole, bytes4 pessimisticVKeySelector, address verifier, bytes32 pessimisticVKey, address multisigRole, (address,string)[] signersToAdd, uint256 newThreshold) returns()
func (_Agglayergateway *AgglayergatewayTransactorSession) Initialize0(defaultAdmin common.Address, aggchainDefaultVKeyRole common.Address, addRouteRole common.Address, freezeRouteRole common.Address, pessimisticVKeySelector [4]byte, verifier common.Address, pessimisticVKey [32]byte, multisigRole common.Address, signersToAdd []IAggLayerGatewaySignerInfo, newThreshold *big.Int) (*types.Transaction, error) {
	return _Agglayergateway.Contract.Initialize0(&_Agglayergateway.TransactOpts, defaultAdmin, aggchainDefaultVKeyRole, addRouteRole, freezeRouteRole, pessimisticVKeySelector, verifier, pessimisticVKey, multisigRole, signersToAdd, newThreshold)
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

// UnsetDefaultAggchainVKey is a paid mutator transaction binding the contract method 0x47c63d36.
//
// Solidity: function unsetDefaultAggchainVKey(bytes4 defaultAggchainSelector) returns()
func (_Agglayergateway *AgglayergatewayTransactor) UnsetDefaultAggchainVKey(opts *bind.TransactOpts, defaultAggchainSelector [4]byte) (*types.Transaction, error) {
	return _Agglayergateway.contract.Transact(opts, "unsetDefaultAggchainVKey", defaultAggchainSelector)
}

// UnsetDefaultAggchainVKey is a paid mutator transaction binding the contract method 0x47c63d36.
//
// Solidity: function unsetDefaultAggchainVKey(bytes4 defaultAggchainSelector) returns()
func (_Agglayergateway *AgglayergatewaySession) UnsetDefaultAggchainVKey(defaultAggchainSelector [4]byte) (*types.Transaction, error) {
	return _Agglayergateway.Contract.UnsetDefaultAggchainVKey(&_Agglayergateway.TransactOpts, defaultAggchainSelector)
}

// UnsetDefaultAggchainVKey is a paid mutator transaction binding the contract method 0x47c63d36.
//
// Solidity: function unsetDefaultAggchainVKey(bytes4 defaultAggchainSelector) returns()
func (_Agglayergateway *AgglayergatewayTransactorSession) UnsetDefaultAggchainVKey(defaultAggchainSelector [4]byte) (*types.Transaction, error) {
	return _Agglayergateway.Contract.UnsetDefaultAggchainVKey(&_Agglayergateway.TransactOpts, defaultAggchainSelector)
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

// UpdateSignersAndThreshold is a paid mutator transaction binding the contract method 0xf51f563a.
//
// Solidity: function updateSignersAndThreshold((address,uint256)[] _signersToRemove, (address,string)[] _signersToAdd, uint256 _newThreshold) returns()
func (_Agglayergateway *AgglayergatewayTransactor) UpdateSignersAndThreshold(opts *bind.TransactOpts, _signersToRemove []IAggLayerGatewayRemoveSignerInfo, _signersToAdd []IAggLayerGatewaySignerInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Agglayergateway.contract.Transact(opts, "updateSignersAndThreshold", _signersToRemove, _signersToAdd, _newThreshold)
}

// UpdateSignersAndThreshold is a paid mutator transaction binding the contract method 0xf51f563a.
//
// Solidity: function updateSignersAndThreshold((address,uint256)[] _signersToRemove, (address,string)[] _signersToAdd, uint256 _newThreshold) returns()
func (_Agglayergateway *AgglayergatewaySession) UpdateSignersAndThreshold(_signersToRemove []IAggLayerGatewayRemoveSignerInfo, _signersToAdd []IAggLayerGatewaySignerInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Agglayergateway.Contract.UpdateSignersAndThreshold(&_Agglayergateway.TransactOpts, _signersToRemove, _signersToAdd, _newThreshold)
}

// UpdateSignersAndThreshold is a paid mutator transaction binding the contract method 0xf51f563a.
//
// Solidity: function updateSignersAndThreshold((address,uint256)[] _signersToRemove, (address,string)[] _signersToAdd, uint256 _newThreshold) returns()
func (_Agglayergateway *AgglayergatewayTransactorSession) UpdateSignersAndThreshold(_signersToRemove []IAggLayerGatewayRemoveSignerInfo, _signersToAdd []IAggLayerGatewaySignerInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Agglayergateway.Contract.UpdateSignersAndThreshold(&_Agglayergateway.TransactOpts, _signersToRemove, _signersToAdd, _newThreshold)
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

// AgglayergatewaySignersAndThresholdUpdatedIterator is returned from FilterSignersAndThresholdUpdated and is used to iterate over the raw logs and unpacked data for SignersAndThresholdUpdated events raised by the Agglayergateway contract.
type AgglayergatewaySignersAndThresholdUpdatedIterator struct {
	Event *AgglayergatewaySignersAndThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *AgglayergatewaySignersAndThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayergatewaySignersAndThresholdUpdated)
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
		it.Event = new(AgglayergatewaySignersAndThresholdUpdated)
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
func (it *AgglayergatewaySignersAndThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayergatewaySignersAndThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayergatewaySignersAndThresholdUpdated represents a SignersAndThresholdUpdated event raised by the Agglayergateway contract.
type AgglayergatewaySignersAndThresholdUpdated struct {
	Signers     []common.Address
	Threshold   *big.Int
	SignersHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSignersAndThresholdUpdated is a free log retrieval operation binding the contract event 0x66d7b0647fdd512b69cbf4f8e1ce8068bfe0b236168e2704ba13b07425eaa743.
//
// Solidity: event SignersAndThresholdUpdated(address[] signers, uint256 threshold, bytes32 signersHash)
func (_Agglayergateway *AgglayergatewayFilterer) FilterSignersAndThresholdUpdated(opts *bind.FilterOpts) (*AgglayergatewaySignersAndThresholdUpdatedIterator, error) {

	logs, sub, err := _Agglayergateway.contract.FilterLogs(opts, "SignersAndThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &AgglayergatewaySignersAndThresholdUpdatedIterator{contract: _Agglayergateway.contract, event: "SignersAndThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchSignersAndThresholdUpdated is a free log subscription operation binding the contract event 0x66d7b0647fdd512b69cbf4f8e1ce8068bfe0b236168e2704ba13b07425eaa743.
//
// Solidity: event SignersAndThresholdUpdated(address[] signers, uint256 threshold, bytes32 signersHash)
func (_Agglayergateway *AgglayergatewayFilterer) WatchSignersAndThresholdUpdated(opts *bind.WatchOpts, sink chan<- *AgglayergatewaySignersAndThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _Agglayergateway.contract.WatchLogs(opts, "SignersAndThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayergatewaySignersAndThresholdUpdated)
				if err := _Agglayergateway.contract.UnpackLog(event, "SignersAndThresholdUpdated", log); err != nil {
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

// ParseSignersAndThresholdUpdated is a log parse operation binding the contract event 0x66d7b0647fdd512b69cbf4f8e1ce8068bfe0b236168e2704ba13b07425eaa743.
//
// Solidity: event SignersAndThresholdUpdated(address[] signers, uint256 threshold, bytes32 signersHash)
func (_Agglayergateway *AgglayergatewayFilterer) ParseSignersAndThresholdUpdated(log types.Log) (*AgglayergatewaySignersAndThresholdUpdated, error) {
	event := new(AgglayergatewaySignersAndThresholdUpdated)
	if err := _Agglayergateway.contract.UnpackLog(event, "SignersAndThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayergatewayUnsetDefaultAggchainVKeyIterator is returned from FilterUnsetDefaultAggchainVKey and is used to iterate over the raw logs and unpacked data for UnsetDefaultAggchainVKey events raised by the Agglayergateway contract.
type AgglayergatewayUnsetDefaultAggchainVKeyIterator struct {
	Event *AgglayergatewayUnsetDefaultAggchainVKey // Event containing the contract specifics and raw log

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
func (it *AgglayergatewayUnsetDefaultAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayergatewayUnsetDefaultAggchainVKey)
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
		it.Event = new(AgglayergatewayUnsetDefaultAggchainVKey)
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
func (it *AgglayergatewayUnsetDefaultAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayergatewayUnsetDefaultAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayergatewayUnsetDefaultAggchainVKey represents a UnsetDefaultAggchainVKey event raised by the Agglayergateway contract.
type AgglayergatewayUnsetDefaultAggchainVKey struct {
	Selector [4]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnsetDefaultAggchainVKey is a free log retrieval operation binding the contract event 0x43b46d79ee2fbaa8a981b5fd78aaf6fb9e38b18d02e1af17564ce37c67078cab.
//
// Solidity: event UnsetDefaultAggchainVKey(bytes4 selector)
func (_Agglayergateway *AgglayergatewayFilterer) FilterUnsetDefaultAggchainVKey(opts *bind.FilterOpts) (*AgglayergatewayUnsetDefaultAggchainVKeyIterator, error) {

	logs, sub, err := _Agglayergateway.contract.FilterLogs(opts, "UnsetDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &AgglayergatewayUnsetDefaultAggchainVKeyIterator{contract: _Agglayergateway.contract, event: "UnsetDefaultAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchUnsetDefaultAggchainVKey is a free log subscription operation binding the contract event 0x43b46d79ee2fbaa8a981b5fd78aaf6fb9e38b18d02e1af17564ce37c67078cab.
//
// Solidity: event UnsetDefaultAggchainVKey(bytes4 selector)
func (_Agglayergateway *AgglayergatewayFilterer) WatchUnsetDefaultAggchainVKey(opts *bind.WatchOpts, sink chan<- *AgglayergatewayUnsetDefaultAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Agglayergateway.contract.WatchLogs(opts, "UnsetDefaultAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayergatewayUnsetDefaultAggchainVKey)
				if err := _Agglayergateway.contract.UnpackLog(event, "UnsetDefaultAggchainVKey", log); err != nil {
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
func (_Agglayergateway *AgglayergatewayFilterer) ParseUnsetDefaultAggchainVKey(log types.Log) (*AgglayergatewayUnsetDefaultAggchainVKey, error) {
	event := new(AgglayergatewayUnsetDefaultAggchainVKey)
	if err := _Agglayergateway.contract.UnpackLog(event, "UnsetDefaultAggchainVKey", log); err != nil {
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

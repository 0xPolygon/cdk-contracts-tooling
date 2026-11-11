// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aggoraclecommittee

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

// IAggOracleCommitteeRemoveValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IAggOracleCommitteeRemoveValidatorInfo struct {
	Addr  common.Address
	Index *big.Int
}

// IAggOracleCommitteeValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IAggOracleCommitteeValidatorInfo struct {
	Addr common.Address
	Url  string
}

// AggoraclecommitteeMetaData contains all meta data concerning the Aggoraclecommittee contract.
var AggoraclecommitteeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIAgglayerGERL2\",\"name\":\"_agglayerGERL2\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AgglayerGERL2CannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GERAlreadyInjected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GERCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndicesNotInDescendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignaturesLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAggOracleProposerCanCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProposerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RenounceOwnershipDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdCannotBeGreaterThanValidators\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyValidators\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorCannotBeThisContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorIndexMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorIndexOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorsNotOrdered\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"}],\"name\":\"GERInjected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"}],\"name\":\"GERMarkedAsInjected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newProposer\",\"type\":\"address\"}],\"name\":\"UpdateAggOracleProposer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ValidatorsAndThresholdUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_VALIDATORS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION_AGG_ORACLE_COMMITTEE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptGlobalExitRootUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agglayerGERL2\",\"outputs\":[{\"internalType\":\"contractIAgglayerGERL2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getValidatorIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorsInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structIAggOracleCommittee.ValidatorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newProposer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structIAggOracleCommittee.ValidatorInfo[]\",\"name\":\"_newValidators\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_newThreshold\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"injectGER\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"injectedGERs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGlobalExitRootUpdater\",\"type\":\"address\"}],\"name\":\"transferGlobalExitRootUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newProposer\",\"type\":\"address\"}],\"name\":\"updateAggOracleProposer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structIAggOracleCommittee.RemoveValidatorInfo[]\",\"name\":\"_validatorsToRemove\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structIAggOracleCommittee.ValidatorInfo[]\",\"name\":\"_validatorsToAdd\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_newThreshold\",\"type\":\"uint256\"}],\"name\":\"updateValidatorsAndThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validatorInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidator\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"verifyGER\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f5ffd5b506040516127a73803806127a783398101604081905261002e91610122565b6001600160a01b03811661005557604051632fc46f4360e01b815260040160405180910390fd5b6001600160a01b03811660805261006a610070565b5061014f565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100c05760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161461011f5780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b5f60208284031215610132575f5ffd5b81516001600160a01b0381168114610148575f5ffd5b9392505050565b6080516126246101835f395f81816102890152818161047501528181610bfe01528181610d020152610d7e01526126245ff3fe608060405234801561000f575f5ffd5b506004361061019a575f3560e01c80638452230e116100e8578063c380353111610093578063e905f1421161006e578063e905f142146103a8578063f2fde38b146103ca578063facd743b146103dd578063fedaa2a714610415575f5ffd5b8063c380353114610351578063d6ad3f8f1461038d578063e30c3978146103a0575f5ffd5b8063a31dee4a116100c3578063a31dee4a14610316578063a8e4fb9014610329578063c053902a14610349575f5ffd5b80638452230e146102e65780638da5cb5b146102fb57806398e04bf814610303575f5ffd5b806354fd4d5011610148578063714897df11610123578063714897df146102ce578063715018a6146102d657806379ba5097146102de575f5ffd5b806354fd4d501461024257806356eb763714610284578063677f3950146102ab575f5ffd5b806335aa2e441161017857806335aa2e44146101e057806342cde4e8146102185780634f1811dd14610221575f5ffd5b80630e1bbf9f1461019e5780631fcc449e146101b357806327498240146101d9575b5f5ffd5b6101b16101ac366004611d3f565b610428565b005b6101c66101c1366004611d3f565b6104cf565b6040519081526020015b60405180910390f35b5f546101c6565b6101f36101ee366004611d5f565b61056d565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101d0565b6101c660015481565b61023461022f366004611d3f565b6105a1565b6040516101d0929190611dc2565b60408051808201909152600681527f76322e302e30000000000000000000000000000000000000000000000000000060208201525b6040516101d09190611de4565b6101f37f000000000000000000000000000000000000000000000000000000000000000081565b6102be6102b9366004611df6565b610647565b60405190151581526020016101d0565b6101c660ff81565b6101b161065d565b6101b161068f565b6102ee61070f565b6040516101d09190611e6d565b6101f36108c6565b6101b161031136600461215b565b610907565b6101b1610324366004611df6565b610b2c565b6003546101f39073ffffffffffffffffffffffffffffffffffffffff1681565b6101b1610d74565b6102776040518060400160405280600681526020017f76322e302e30000000000000000000000000000000000000000000000000000081525081565b6101b161039b366004611d3f565b610df9565b6101f3610e0a565b6102be6103b6366004611d5f565b60046020525f908152604090205460ff1681565b6101b16103d8366004611d3f565b610e32565b6102be6103eb366004611d3f565b73ffffffffffffffffffffffffffffffffffffffff165f9081526002602052604090205460ff1690565b6101b16104233660046121bd565b610ee9565b610430610efc565b6040517f0e1bbf9f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301527f00000000000000000000000000000000000000000000000000000000000000001690630e1bbf9f906024015f604051808303815f87803b1580156104b6575f5ffd5b505af11580156104c8573d5f5f3e3d5ffd5b5050505050565b5f8054815b8181101561053a578373ffffffffffffffffffffffffffffffffffffffff165f8281548110610505576105056122ae565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff1603610532579392505050565b6001016104d4565b506040517f580e542f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f818154811061057b575f80fd5b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b60026020525f90815260409020805460018201805460ff90921692916105c6906122db565b80601f01602080910402602001604051908101604052809291908181526020018280546105f2906122db565b801561063d5780601f106106145761010080835404028352916020019161063d565b820191905f5260205f20905b81548152906001019060200180831161062057829003601f168201915b5050505050905082565b5f610653848484610f56565b5060019392505050565b6040517fffc0fd9300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3380610699610e0a565b73ffffffffffffffffffffffffffffffffffffffff1614610703576040517f118cdaa700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024015b60405180910390fd5b61070c81610fec565b50565b5f80546060918167ffffffffffffffff81111561072e5761072e611f1b565b60405190808252806020026020018201604052801561077357816020015b604080518082019091525f81526060602082015281526020019060019003908161074c5790505b5090505f5b828110156108bf5760405180604001604052805f838154811061079d5761079d6122ae565b5f91825260208083209091015473ffffffffffffffffffffffffffffffffffffffff1683528154920191600291908190869081106107dd576107dd6122ae565b5f91825260208083209091015473ffffffffffffffffffffffffffffffffffffffff1683528201929092526040019020600101805461081b906122db565b80601f0160208091040260200160405190810160405280929190818152602001828054610847906122db565b80156108925780601f1061086957610100808354040283529160200191610892565b820191905f5260205f20905b81548152906001019060200180831161087557829003601f168201915b50505050508152508282815181106108ac576108ac6122ae565b6020908102919091010152600101610778565b5092915050565b5f807f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005b5473ffffffffffffffffffffffffffffffffffffffff1692915050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f811580156109515750825b90505f8267ffffffffffffffff16600114801561096d5750303b155b90508115801561097b575080155b156109b2576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610a135784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b73ffffffffffffffffffffffffffffffffffffffff8916610a60576040517f9b15e16f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a6989611040565b610a71611051565b610a7a88611059565b604080515f80825260208201909252610ac091610ab8565b604080518082019091525f8082526020820152815260200190600190039081610a925790505b508888611114565b8315610b215784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b60035473ffffffffffffffffffffffffffffffffffffffff163314610b7d576040517f85d8f64f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b88838383610f56565b5f8381526004602081905260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905590517f257b363200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169163257b363291610c359187910190815260200190565b6020604051808303815f875af1158015610c51573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c75919061232c565b15610ca95760405183907fcafdbc8c7e71f2990afc9e787ff729a197666534c9a49959c1373c819b708a8f905f90a2505050565b60405183907f45b91a7a09575a05b9dd504dff707cdeb7eda5f9d514491caf5951551b020e1e905f90a26040517f12da06b2000000000000000000000000000000000000000000000000000000008152600481018490527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906312da06b2906024015f604051808303815f87803b158015610d58575f5ffd5b505af1158015610d6a573d5f5f3e3d5ffd5b505050505b505050565b610d7c610efc565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c053902a6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015610de1575f5ffd5b505af1158015610df3573d5f5f3e3d5ffd5b50505050565b610e01610efc565b61070c81611059565b5f807f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c006108ea565b610e3a610efc565b7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081178255610ea36108c6565b73ffffffffffffffffffffffffffffffffffffffff167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a35050565b610ef1610efc565b610d6f838383611114565b33610f056108c6565b73ffffffffffffffffffffffffffffffffffffffff1614610f54576040517f118cdaa70000000000000000000000000000000000000000000000000000000081523360048201526024016106fa565b565b82610f8d576040517fa3e6df8400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8381526004602052604090205460ff1615610fd5576040517f7bde264700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f610fdf846112ef565b9050610df3818484611320565b7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0080547fffffffffffffffffffffffff000000000000000000000000000000000000000016815561103c826114b3565b5050565b611048611548565b61070c816115af565b610f54611548565b73ffffffffffffffffffffffffffffffffffffffff81166110a6576040517f18cba8a700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517f93c536fad4968838f70b6a5841cc12ffbd1a657048ef449f715f0d9ebdbeffca905f90a250565b6001835111156111bc575f5b6001845161112e9190612370565b8110156111ba5783611141826001612383565b81518110611151576111516122ae565b60200260200101516020015184828151811061116f5761116f6122ae565b602002602001015160200151116111b2576040517fb9a11d3100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600101611120565b505b5f5b83518110156112125761120a8482815181106111dc576111dc6122ae565b60200260200101515f01518583815181106111f9576111f96122ae565b602002602001015160200151611606565b6001016111be565b505f5b825181101561126957611261838281518110611233576112336122ae565b60200260200101515f0151848381518110611250576112506122ae565b602002602001015160200151611827565b600101611215565b505f5460ff10156112a6576040517f2bb8575c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112af816119eb565b7f50051c796f4289791fedb6e1f365497b6527df6f6f19527d566e3c9791af54605f6001546040516112e2929190612396565b60405180910390a1505050565b5f8160405160200161130391815260200190565b604051602081830303815290604052805190602001209050919050565b60015461132e8160416123fb565b8214611366576040517ffb6c730400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80805b83811015610d6a576113e08787876113838560416123fb565b9061138f866001612383565b61139a9060416123fb565b926113a793929190612412565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250611a6492505050565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526002602052604090205490925060ff16611441576040517fe51315d200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16116114a6576040517fe211dd1e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b909150819060010161136a565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080547fffffffffffffffffffffffff0000000000000000000000000000000000000000811673ffffffffffffffffffffffffffffffffffffffff848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16610f54576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6115b7611548565b73ffffffffffffffffffffffffffffffffffffffff8116610703576040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081525f60048201526024016106fa565b5f54808210611641576040517f99e56ef700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff165f838154811061166a5761166a6122ae565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff16146116c2576040517fd29e638b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83165f90815260026020526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681559061171a6001830182611ccd565b505f9050611729600183612370565b81548110611739576117396122ae565b5f918252602082200154815473ffffffffffffffffffffffffffffffffffffffff90911691908490811061176f5761176f6122ae565b5f918252602082200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff93909316929092179091558054806117cb576117cb612439565b5f8281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055505050565b73ffffffffffffffffffffffffffffffffffffffff8216611874576040517f5c074f7000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3073ffffffffffffffffffffffffffffffffffffffff8316036118c3576040517f7f32165400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82165f9081526002602052604090205460ff1615611922576040517f8d09a73e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8054600180820183557f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56390910180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff86169081179091556040805180820182528381526020808201878152938652600290529320835181547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016901515178155905190918201906104c890826124aa565b805f03611a24576040517ff412416600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f54811115611a5f576040517f270e262e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600155565b5f5f5f5f611a728686611a8e565b925092509250611a828282611ad7565b50909150505b92915050565b5f5f5f8351604103611ac5576020840151604085015160608601515f1a611ab788828585611bda565b955095509550505050611ad0565b505081515f91506002905b9250925092565b5f826003811115611aea57611aea6125c1565b03611af3575050565b6001826003811115611b0757611b076125c1565b03611b3e576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002826003811115611b5257611b526125c1565b03611b8c576040517ffce698f7000000000000000000000000000000000000000000000000000000008152600481018290526024016106fa565b6003826003811115611ba057611ba06125c1565b0361103c576040517fd78bce0c000000000000000000000000000000000000000000000000000000008152600481018290526024016106fa565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115611c1357505f91506003905082611cc3565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015611c64573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116611cba57505f925060019150829050611cc3565b92505f91508190505b9450945094915050565b508054611cd9906122db565b5f825580601f10611ce8575050565b601f0160209004905f5260205f209081019061070c91905b80821115611d13575f8155600101611d00565b5090565b803573ffffffffffffffffffffffffffffffffffffffff81168114611d3a575f5ffd5b919050565b5f60208284031215611d4f575f5ffd5b611d5882611d17565b9392505050565b5f60208284031215611d6f575f5ffd5b5035919050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b8215158152604060208201525f611ddc6040830184611d76565b949350505050565b602081525f611d586020830184611d76565b5f5f5f60408486031215611e08575f5ffd5b83359250602084013567ffffffffffffffff811115611e25575f5ffd5b8401601f81018613611e35575f5ffd5b803567ffffffffffffffff811115611e4b575f5ffd5b866020828401011115611e5c575f5ffd5b939660209190910195509293505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b82811015611f0f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815173ffffffffffffffffffffffffffffffffffffffff81511686526020810151905060406020870152611ef96040870182611d76565b9550506020938401939190910190600101611e93565b50929695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6040805190810167ffffffffffffffff81118282101715611f6b57611f6b611f1b565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611fb857611fb8611f1b565b604052919050565b5f67ffffffffffffffff821115611fd957611fd9611f1b565b5060051b60200190565b5f82601f830112611ff2575f5ffd5b813561200561200082611fc0565b611f71565b8082825260208201915060208360051b860101925085831115612026575f5ffd5b602085015b8381101561215157803567ffffffffffffffff811115612049575f5ffd5b860160408189037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001121561207c575f5ffd5b612084611f48565b61209060208301611d17565b8152604082013567ffffffffffffffff8111156120ab575f5ffd5b60208184010192505088601f8301126120c2575f5ffd5b813567ffffffffffffffff8111156120dc576120dc611f1b565b61210d60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611f71565b8181528a6020838601011115612121575f5ffd5b816020850160208301375f602083830101528060208401525050808552505060208301925060208101905061202b565b5095945050505050565b5f5f5f5f6080858703121561216e575f5ffd5b61217785611d17565b935061218560208601611d17565b9250604085013567ffffffffffffffff8111156121a0575f5ffd5b6121ac87828801611fe3565b949793965093946060013593505050565b5f5f5f606084860312156121cf575f5ffd5b833567ffffffffffffffff8111156121e5575f5ffd5b8401601f810186136121f5575f5ffd5b803561220361200082611fc0565b8082825260208201915060208360061b850101925088831115612224575f5ffd5b6020840193505b82841015612272576040848a031215612242575f5ffd5b61224a611f48565b61225385611d17565b815260208581013581830152908352604090940193919091019061222b565b9550505050602084013567ffffffffffffffff811115612290575f5ffd5b61229c86828701611fe3565b93969395505050506040919091013590565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b600181811c908216806122ef57607f821691505b602082108103612326577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b5f6020828403121561233c575f5ffd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b81810381811115611a8857611a88612343565b80820180821115611a8857611a88612343565b604080825283549082018190525f8481526020812090916060840190835b818110156123e857835473ffffffffffffffffffffffffffffffffffffffff168352600193840193602090930192016123b4565b5050602093909301939093525092915050565b8082028115828204841417611a8857611a88612343565b5f5f85851115612420575f5ffd5b8386111561242c575f5ffd5b5050820193919092039150565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b601f821115610d6f57805f5260205f20601f840160051c8101602085101561248b5750805b601f840160051c820191505b818110156104c8575f8155600101612497565b815167ffffffffffffffff8111156124c4576124c4611f1b565b6124d8816124d284546122db565b84612466565b6020601f821160018114612529575f83156124f35750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b1784556104c8565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b828110156125765787850151825560209485019460019092019101612556565b50848210156125b257868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffdfea2646970667358221220d0e57f7961ae323ccc8cadafcf2253c81a11824b093d4ba4f4bc71b42024942164736f6c634300081c0033",
}

// AggoraclecommitteeABI is the input ABI used to generate the binding from.
// Deprecated: Use AggoraclecommitteeMetaData.ABI instead.
var AggoraclecommitteeABI = AggoraclecommitteeMetaData.ABI

// AggoraclecommitteeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AggoraclecommitteeMetaData.Bin instead.
var AggoraclecommitteeBin = AggoraclecommitteeMetaData.Bin

// DeployAggoraclecommittee deploys a new Ethereum contract, binding an instance of Aggoraclecommittee to it.
func DeployAggoraclecommittee(auth *bind.TransactOpts, backend bind.ContractBackend, _agglayerGERL2 common.Address) (common.Address, *types.Transaction, *Aggoraclecommittee, error) {
	parsed, err := AggoraclecommitteeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AggoraclecommitteeBin), backend, _agglayerGERL2)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Aggoraclecommittee{AggoraclecommitteeCaller: AggoraclecommitteeCaller{contract: contract}, AggoraclecommitteeTransactor: AggoraclecommitteeTransactor{contract: contract}, AggoraclecommitteeFilterer: AggoraclecommitteeFilterer{contract: contract}}, nil
}

// Aggoraclecommittee is an auto generated Go binding around an Ethereum contract.
type Aggoraclecommittee struct {
	AggoraclecommitteeCaller     // Read-only binding to the contract
	AggoraclecommitteeTransactor // Write-only binding to the contract
	AggoraclecommitteeFilterer   // Log filterer for contract events
}

// AggoraclecommitteeCaller is an auto generated read-only Go binding around an Ethereum contract.
type AggoraclecommitteeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggoraclecommitteeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AggoraclecommitteeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggoraclecommitteeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AggoraclecommitteeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggoraclecommitteeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AggoraclecommitteeSession struct {
	Contract     *Aggoraclecommittee // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AggoraclecommitteeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AggoraclecommitteeCallerSession struct {
	Contract *AggoraclecommitteeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AggoraclecommitteeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AggoraclecommitteeTransactorSession struct {
	Contract     *AggoraclecommitteeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AggoraclecommitteeRaw is an auto generated low-level Go binding around an Ethereum contract.
type AggoraclecommitteeRaw struct {
	Contract *Aggoraclecommittee // Generic contract binding to access the raw methods on
}

// AggoraclecommitteeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AggoraclecommitteeCallerRaw struct {
	Contract *AggoraclecommitteeCaller // Generic read-only contract binding to access the raw methods on
}

// AggoraclecommitteeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AggoraclecommitteeTransactorRaw struct {
	Contract *AggoraclecommitteeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAggoraclecommittee creates a new instance of Aggoraclecommittee, bound to a specific deployed contract.
func NewAggoraclecommittee(address common.Address, backend bind.ContractBackend) (*Aggoraclecommittee, error) {
	contract, err := bindAggoraclecommittee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aggoraclecommittee{AggoraclecommitteeCaller: AggoraclecommitteeCaller{contract: contract}, AggoraclecommitteeTransactor: AggoraclecommitteeTransactor{contract: contract}, AggoraclecommitteeFilterer: AggoraclecommitteeFilterer{contract: contract}}, nil
}

// NewAggoraclecommitteeCaller creates a new read-only instance of Aggoraclecommittee, bound to a specific deployed contract.
func NewAggoraclecommitteeCaller(address common.Address, caller bind.ContractCaller) (*AggoraclecommitteeCaller, error) {
	contract, err := bindAggoraclecommittee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggoraclecommitteeCaller{contract: contract}, nil
}

// NewAggoraclecommitteeTransactor creates a new write-only instance of Aggoraclecommittee, bound to a specific deployed contract.
func NewAggoraclecommitteeTransactor(address common.Address, transactor bind.ContractTransactor) (*AggoraclecommitteeTransactor, error) {
	contract, err := bindAggoraclecommittee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggoraclecommitteeTransactor{contract: contract}, nil
}

// NewAggoraclecommitteeFilterer creates a new log filterer instance of Aggoraclecommittee, bound to a specific deployed contract.
func NewAggoraclecommitteeFilterer(address common.Address, filterer bind.ContractFilterer) (*AggoraclecommitteeFilterer, error) {
	contract, err := bindAggoraclecommittee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggoraclecommitteeFilterer{contract: contract}, nil
}

// bindAggoraclecommittee binds a generic wrapper to an already deployed contract.
func bindAggoraclecommittee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AggoraclecommitteeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggoraclecommittee *AggoraclecommitteeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aggoraclecommittee.Contract.AggoraclecommitteeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggoraclecommittee *AggoraclecommitteeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.AggoraclecommitteeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggoraclecommittee *AggoraclecommitteeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.AggoraclecommitteeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggoraclecommittee *AggoraclecommitteeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aggoraclecommittee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggoraclecommittee *AggoraclecommitteeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggoraclecommittee *AggoraclecommitteeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.contract.Transact(opts, method, params...)
}

// MAXVALIDATORS is a free data retrieval call binding the contract method 0x714897df.
//
// Solidity: function MAX_VALIDATORS() view returns(uint256)
func (_Aggoraclecommittee *AggoraclecommitteeCaller) MAXVALIDATORS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "MAX_VALIDATORS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXVALIDATORS is a free data retrieval call binding the contract method 0x714897df.
//
// Solidity: function MAX_VALIDATORS() view returns(uint256)
func (_Aggoraclecommittee *AggoraclecommitteeSession) MAXVALIDATORS() (*big.Int, error) {
	return _Aggoraclecommittee.Contract.MAXVALIDATORS(&_Aggoraclecommittee.CallOpts)
}

// MAXVALIDATORS is a free data retrieval call binding the contract method 0x714897df.
//
// Solidity: function MAX_VALIDATORS() view returns(uint256)
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) MAXVALIDATORS() (*big.Int, error) {
	return _Aggoraclecommittee.Contract.MAXVALIDATORS(&_Aggoraclecommittee.CallOpts)
}

// VERSIONAGGORACLECOMMITTEE is a free data retrieval call binding the contract method 0xc3803531.
//
// Solidity: function VERSION_AGG_ORACLE_COMMITTEE() view returns(string)
func (_Aggoraclecommittee *AggoraclecommitteeCaller) VERSIONAGGORACLECOMMITTEE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "VERSION_AGG_ORACLE_COMMITTEE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSIONAGGORACLECOMMITTEE is a free data retrieval call binding the contract method 0xc3803531.
//
// Solidity: function VERSION_AGG_ORACLE_COMMITTEE() view returns(string)
func (_Aggoraclecommittee *AggoraclecommitteeSession) VERSIONAGGORACLECOMMITTEE() (string, error) {
	return _Aggoraclecommittee.Contract.VERSIONAGGORACLECOMMITTEE(&_Aggoraclecommittee.CallOpts)
}

// VERSIONAGGORACLECOMMITTEE is a free data retrieval call binding the contract method 0xc3803531.
//
// Solidity: function VERSION_AGG_ORACLE_COMMITTEE() view returns(string)
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) VERSIONAGGORACLECOMMITTEE() (string, error) {
	return _Aggoraclecommittee.Contract.VERSIONAGGORACLECOMMITTEE(&_Aggoraclecommittee.CallOpts)
}

// AgglayerGERL2 is a free data retrieval call binding the contract method 0x56eb7637.
//
// Solidity: function agglayerGERL2() view returns(address)
func (_Aggoraclecommittee *AggoraclecommitteeCaller) AgglayerGERL2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "agglayerGERL2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgglayerGERL2 is a free data retrieval call binding the contract method 0x56eb7637.
//
// Solidity: function agglayerGERL2() view returns(address)
func (_Aggoraclecommittee *AggoraclecommitteeSession) AgglayerGERL2() (common.Address, error) {
	return _Aggoraclecommittee.Contract.AgglayerGERL2(&_Aggoraclecommittee.CallOpts)
}

// AgglayerGERL2 is a free data retrieval call binding the contract method 0x56eb7637.
//
// Solidity: function agglayerGERL2() view returns(address)
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) AgglayerGERL2() (common.Address, error) {
	return _Aggoraclecommittee.Contract.AgglayerGERL2(&_Aggoraclecommittee.CallOpts)
}

// GetValidatorIndex is a free data retrieval call binding the contract method 0x1fcc449e.
//
// Solidity: function getValidatorIndex(address _validator) view returns(uint256)
func (_Aggoraclecommittee *AggoraclecommitteeCaller) GetValidatorIndex(opts *bind.CallOpts, _validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "getValidatorIndex", _validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorIndex is a free data retrieval call binding the contract method 0x1fcc449e.
//
// Solidity: function getValidatorIndex(address _validator) view returns(uint256)
func (_Aggoraclecommittee *AggoraclecommitteeSession) GetValidatorIndex(_validator common.Address) (*big.Int, error) {
	return _Aggoraclecommittee.Contract.GetValidatorIndex(&_Aggoraclecommittee.CallOpts, _validator)
}

// GetValidatorIndex is a free data retrieval call binding the contract method 0x1fcc449e.
//
// Solidity: function getValidatorIndex(address _validator) view returns(uint256)
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) GetValidatorIndex(_validator common.Address) (*big.Int, error) {
	return _Aggoraclecommittee.Contract.GetValidatorIndex(&_Aggoraclecommittee.CallOpts, _validator)
}

// GetValidatorsCount is a free data retrieval call binding the contract method 0x27498240.
//
// Solidity: function getValidatorsCount() view returns(uint256)
func (_Aggoraclecommittee *AggoraclecommitteeCaller) GetValidatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "getValidatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorsCount is a free data retrieval call binding the contract method 0x27498240.
//
// Solidity: function getValidatorsCount() view returns(uint256)
func (_Aggoraclecommittee *AggoraclecommitteeSession) GetValidatorsCount() (*big.Int, error) {
	return _Aggoraclecommittee.Contract.GetValidatorsCount(&_Aggoraclecommittee.CallOpts)
}

// GetValidatorsCount is a free data retrieval call binding the contract method 0x27498240.
//
// Solidity: function getValidatorsCount() view returns(uint256)
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) GetValidatorsCount() (*big.Int, error) {
	return _Aggoraclecommittee.Contract.GetValidatorsCount(&_Aggoraclecommittee.CallOpts)
}

// GetValidatorsInfo is a free data retrieval call binding the contract method 0x8452230e.
//
// Solidity: function getValidatorsInfo() view returns((address,string)[])
func (_Aggoraclecommittee *AggoraclecommitteeCaller) GetValidatorsInfo(opts *bind.CallOpts) ([]IAggOracleCommitteeValidatorInfo, error) {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "getValidatorsInfo")

	if err != nil {
		return *new([]IAggOracleCommitteeValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAggOracleCommitteeValidatorInfo)).(*[]IAggOracleCommitteeValidatorInfo)

	return out0, err

}

// GetValidatorsInfo is a free data retrieval call binding the contract method 0x8452230e.
//
// Solidity: function getValidatorsInfo() view returns((address,string)[])
func (_Aggoraclecommittee *AggoraclecommitteeSession) GetValidatorsInfo() ([]IAggOracleCommitteeValidatorInfo, error) {
	return _Aggoraclecommittee.Contract.GetValidatorsInfo(&_Aggoraclecommittee.CallOpts)
}

// GetValidatorsInfo is a free data retrieval call binding the contract method 0x8452230e.
//
// Solidity: function getValidatorsInfo() view returns((address,string)[])
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) GetValidatorsInfo() ([]IAggOracleCommitteeValidatorInfo, error) {
	return _Aggoraclecommittee.Contract.GetValidatorsInfo(&_Aggoraclecommittee.CallOpts)
}

// InjectedGERs is a free data retrieval call binding the contract method 0xe905f142.
//
// Solidity: function injectedGERs(bytes32 ) view returns(bool)
func (_Aggoraclecommittee *AggoraclecommitteeCaller) InjectedGERs(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "injectedGERs", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InjectedGERs is a free data retrieval call binding the contract method 0xe905f142.
//
// Solidity: function injectedGERs(bytes32 ) view returns(bool)
func (_Aggoraclecommittee *AggoraclecommitteeSession) InjectedGERs(arg0 [32]byte) (bool, error) {
	return _Aggoraclecommittee.Contract.InjectedGERs(&_Aggoraclecommittee.CallOpts, arg0)
}

// InjectedGERs is a free data retrieval call binding the contract method 0xe905f142.
//
// Solidity: function injectedGERs(bytes32 ) view returns(bool)
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) InjectedGERs(arg0 [32]byte) (bool, error) {
	return _Aggoraclecommittee.Contract.InjectedGERs(&_Aggoraclecommittee.CallOpts, arg0)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _validator) view returns(bool)
func (_Aggoraclecommittee *AggoraclecommitteeCaller) IsValidator(opts *bind.CallOpts, _validator common.Address) (bool, error) {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "isValidator", _validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _validator) view returns(bool)
func (_Aggoraclecommittee *AggoraclecommitteeSession) IsValidator(_validator common.Address) (bool, error) {
	return _Aggoraclecommittee.Contract.IsValidator(&_Aggoraclecommittee.CallOpts, _validator)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _validator) view returns(bool)
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) IsValidator(_validator common.Address) (bool, error) {
	return _Aggoraclecommittee.Contract.IsValidator(&_Aggoraclecommittee.CallOpts, _validator)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Aggoraclecommittee *AggoraclecommitteeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Aggoraclecommittee *AggoraclecommitteeSession) Owner() (common.Address, error) {
	return _Aggoraclecommittee.Contract.Owner(&_Aggoraclecommittee.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) Owner() (common.Address, error) {
	return _Aggoraclecommittee.Contract.Owner(&_Aggoraclecommittee.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Aggoraclecommittee *AggoraclecommitteeCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Aggoraclecommittee *AggoraclecommitteeSession) PendingOwner() (common.Address, error) {
	return _Aggoraclecommittee.Contract.PendingOwner(&_Aggoraclecommittee.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) PendingOwner() (common.Address, error) {
	return _Aggoraclecommittee.Contract.PendingOwner(&_Aggoraclecommittee.CallOpts)
}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_Aggoraclecommittee *AggoraclecommitteeCaller) Proposer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "proposer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_Aggoraclecommittee *AggoraclecommitteeSession) Proposer() (common.Address, error) {
	return _Aggoraclecommittee.Contract.Proposer(&_Aggoraclecommittee.CallOpts)
}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) Proposer() (common.Address, error) {
	return _Aggoraclecommittee.Contract.Proposer(&_Aggoraclecommittee.CallOpts)
}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() pure returns()
func (_Aggoraclecommittee *AggoraclecommitteeCaller) RenounceOwnership(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "renounceOwnership")

	if err != nil {
		return err
	}

	return err

}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() pure returns()
func (_Aggoraclecommittee *AggoraclecommitteeSession) RenounceOwnership() error {
	return _Aggoraclecommittee.Contract.RenounceOwnership(&_Aggoraclecommittee.CallOpts)
}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() pure returns()
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) RenounceOwnership() error {
	return _Aggoraclecommittee.Contract.RenounceOwnership(&_Aggoraclecommittee.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_Aggoraclecommittee *AggoraclecommitteeCaller) Threshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_Aggoraclecommittee *AggoraclecommitteeSession) Threshold() (*big.Int, error) {
	return _Aggoraclecommittee.Contract.Threshold(&_Aggoraclecommittee.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) Threshold() (*big.Int, error) {
	return _Aggoraclecommittee.Contract.Threshold(&_Aggoraclecommittee.CallOpts)
}

// ValidatorInfo is a free data retrieval call binding the contract method 0x4f1811dd.
//
// Solidity: function validatorInfo(address ) view returns(bool isValidator, string url)
func (_Aggoraclecommittee *AggoraclecommitteeCaller) ValidatorInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	IsValidator bool
	Url         string
}, error) {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "validatorInfo", arg0)

	outstruct := new(struct {
		IsValidator bool
		Url         string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsValidator = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Url = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// ValidatorInfo is a free data retrieval call binding the contract method 0x4f1811dd.
//
// Solidity: function validatorInfo(address ) view returns(bool isValidator, string url)
func (_Aggoraclecommittee *AggoraclecommitteeSession) ValidatorInfo(arg0 common.Address) (struct {
	IsValidator bool
	Url         string
}, error) {
	return _Aggoraclecommittee.Contract.ValidatorInfo(&_Aggoraclecommittee.CallOpts, arg0)
}

// ValidatorInfo is a free data retrieval call binding the contract method 0x4f1811dd.
//
// Solidity: function validatorInfo(address ) view returns(bool isValidator, string url)
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) ValidatorInfo(arg0 common.Address) (struct {
	IsValidator bool
	Url         string
}, error) {
	return _Aggoraclecommittee.Contract.ValidatorInfo(&_Aggoraclecommittee.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_Aggoraclecommittee *AggoraclecommitteeCaller) Validators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "validators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_Aggoraclecommittee *AggoraclecommitteeSession) Validators(arg0 *big.Int) (common.Address, error) {
	return _Aggoraclecommittee.Contract.Validators(&_Aggoraclecommittee.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) Validators(arg0 *big.Int) (common.Address, error) {
	return _Aggoraclecommittee.Contract.Validators(&_Aggoraclecommittee.CallOpts, arg0)
}

// VerifyGER is a free data retrieval call binding the contract method 0x677f3950.
//
// Solidity: function verifyGER(bytes32 _globalExitRoot, bytes _signatures) view returns(bool)
func (_Aggoraclecommittee *AggoraclecommitteeCaller) VerifyGER(opts *bind.CallOpts, _globalExitRoot [32]byte, _signatures []byte) (bool, error) {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "verifyGER", _globalExitRoot, _signatures)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyGER is a free data retrieval call binding the contract method 0x677f3950.
//
// Solidity: function verifyGER(bytes32 _globalExitRoot, bytes _signatures) view returns(bool)
func (_Aggoraclecommittee *AggoraclecommitteeSession) VerifyGER(_globalExitRoot [32]byte, _signatures []byte) (bool, error) {
	return _Aggoraclecommittee.Contract.VerifyGER(&_Aggoraclecommittee.CallOpts, _globalExitRoot, _signatures)
}

// VerifyGER is a free data retrieval call binding the contract method 0x677f3950.
//
// Solidity: function verifyGER(bytes32 _globalExitRoot, bytes _signatures) view returns(bool)
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) VerifyGER(_globalExitRoot [32]byte, _signatures []byte) (bool, error) {
	return _Aggoraclecommittee.Contract.VerifyGER(&_Aggoraclecommittee.CallOpts, _globalExitRoot, _signatures)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Aggoraclecommittee *AggoraclecommitteeCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Aggoraclecommittee *AggoraclecommitteeSession) Version() (string, error) {
	return _Aggoraclecommittee.Contract.Version(&_Aggoraclecommittee.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) Version() (string, error) {
	return _Aggoraclecommittee.Contract.Version(&_Aggoraclecommittee.CallOpts)
}

// AcceptGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0xc053902a.
//
// Solidity: function acceptGlobalExitRootUpdater() returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactor) AcceptGlobalExitRootUpdater(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggoraclecommittee.contract.Transact(opts, "acceptGlobalExitRootUpdater")
}

// AcceptGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0xc053902a.
//
// Solidity: function acceptGlobalExitRootUpdater() returns()
func (_Aggoraclecommittee *AggoraclecommitteeSession) AcceptGlobalExitRootUpdater() (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.AcceptGlobalExitRootUpdater(&_Aggoraclecommittee.TransactOpts)
}

// AcceptGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0xc053902a.
//
// Solidity: function acceptGlobalExitRootUpdater() returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactorSession) AcceptGlobalExitRootUpdater() (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.AcceptGlobalExitRootUpdater(&_Aggoraclecommittee.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggoraclecommittee.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Aggoraclecommittee *AggoraclecommitteeSession) AcceptOwnership() (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.AcceptOwnership(&_Aggoraclecommittee.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.AcceptOwnership(&_Aggoraclecommittee.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x98e04bf8.
//
// Solidity: function initialize(address _owner, address _newProposer, (address,string)[] _newValidators, uint256 _newThreshold) returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _newProposer common.Address, _newValidators []IAggOracleCommitteeValidatorInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Aggoraclecommittee.contract.Transact(opts, "initialize", _owner, _newProposer, _newValidators, _newThreshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x98e04bf8.
//
// Solidity: function initialize(address _owner, address _newProposer, (address,string)[] _newValidators, uint256 _newThreshold) returns()
func (_Aggoraclecommittee *AggoraclecommitteeSession) Initialize(_owner common.Address, _newProposer common.Address, _newValidators []IAggOracleCommitteeValidatorInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.Initialize(&_Aggoraclecommittee.TransactOpts, _owner, _newProposer, _newValidators, _newThreshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x98e04bf8.
//
// Solidity: function initialize(address _owner, address _newProposer, (address,string)[] _newValidators, uint256 _newThreshold) returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactorSession) Initialize(_owner common.Address, _newProposer common.Address, _newValidators []IAggOracleCommitteeValidatorInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.Initialize(&_Aggoraclecommittee.TransactOpts, _owner, _newProposer, _newValidators, _newThreshold)
}

// InjectGER is a paid mutator transaction binding the contract method 0xa31dee4a.
//
// Solidity: function injectGER(bytes32 _globalExitRoot, bytes _signatures) returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactor) InjectGER(opts *bind.TransactOpts, _globalExitRoot [32]byte, _signatures []byte) (*types.Transaction, error) {
	return _Aggoraclecommittee.contract.Transact(opts, "injectGER", _globalExitRoot, _signatures)
}

// InjectGER is a paid mutator transaction binding the contract method 0xa31dee4a.
//
// Solidity: function injectGER(bytes32 _globalExitRoot, bytes _signatures) returns()
func (_Aggoraclecommittee *AggoraclecommitteeSession) InjectGER(_globalExitRoot [32]byte, _signatures []byte) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.InjectGER(&_Aggoraclecommittee.TransactOpts, _globalExitRoot, _signatures)
}

// InjectGER is a paid mutator transaction binding the contract method 0xa31dee4a.
//
// Solidity: function injectGER(bytes32 _globalExitRoot, bytes _signatures) returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactorSession) InjectGER(_globalExitRoot [32]byte, _signatures []byte) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.InjectGER(&_Aggoraclecommittee.TransactOpts, _globalExitRoot, _signatures)
}

// TransferGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x0e1bbf9f.
//
// Solidity: function transferGlobalExitRootUpdater(address _newGlobalExitRootUpdater) returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactor) TransferGlobalExitRootUpdater(opts *bind.TransactOpts, _newGlobalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Aggoraclecommittee.contract.Transact(opts, "transferGlobalExitRootUpdater", _newGlobalExitRootUpdater)
}

// TransferGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x0e1bbf9f.
//
// Solidity: function transferGlobalExitRootUpdater(address _newGlobalExitRootUpdater) returns()
func (_Aggoraclecommittee *AggoraclecommitteeSession) TransferGlobalExitRootUpdater(_newGlobalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.TransferGlobalExitRootUpdater(&_Aggoraclecommittee.TransactOpts, _newGlobalExitRootUpdater)
}

// TransferGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x0e1bbf9f.
//
// Solidity: function transferGlobalExitRootUpdater(address _newGlobalExitRootUpdater) returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactorSession) TransferGlobalExitRootUpdater(_newGlobalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.TransferGlobalExitRootUpdater(&_Aggoraclecommittee.TransactOpts, _newGlobalExitRootUpdater)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Aggoraclecommittee.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Aggoraclecommittee *AggoraclecommitteeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.TransferOwnership(&_Aggoraclecommittee.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.TransferOwnership(&_Aggoraclecommittee.TransactOpts, newOwner)
}

// UpdateAggOracleProposer is a paid mutator transaction binding the contract method 0xd6ad3f8f.
//
// Solidity: function updateAggOracleProposer(address _newProposer) returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactor) UpdateAggOracleProposer(opts *bind.TransactOpts, _newProposer common.Address) (*types.Transaction, error) {
	return _Aggoraclecommittee.contract.Transact(opts, "updateAggOracleProposer", _newProposer)
}

// UpdateAggOracleProposer is a paid mutator transaction binding the contract method 0xd6ad3f8f.
//
// Solidity: function updateAggOracleProposer(address _newProposer) returns()
func (_Aggoraclecommittee *AggoraclecommitteeSession) UpdateAggOracleProposer(_newProposer common.Address) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.UpdateAggOracleProposer(&_Aggoraclecommittee.TransactOpts, _newProposer)
}

// UpdateAggOracleProposer is a paid mutator transaction binding the contract method 0xd6ad3f8f.
//
// Solidity: function updateAggOracleProposer(address _newProposer) returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactorSession) UpdateAggOracleProposer(_newProposer common.Address) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.UpdateAggOracleProposer(&_Aggoraclecommittee.TransactOpts, _newProposer)
}

// UpdateValidatorsAndThreshold is a paid mutator transaction binding the contract method 0xfedaa2a7.
//
// Solidity: function updateValidatorsAndThreshold((address,uint256)[] _validatorsToRemove, (address,string)[] _validatorsToAdd, uint256 _newThreshold) returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactor) UpdateValidatorsAndThreshold(opts *bind.TransactOpts, _validatorsToRemove []IAggOracleCommitteeRemoveValidatorInfo, _validatorsToAdd []IAggOracleCommitteeValidatorInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Aggoraclecommittee.contract.Transact(opts, "updateValidatorsAndThreshold", _validatorsToRemove, _validatorsToAdd, _newThreshold)
}

// UpdateValidatorsAndThreshold is a paid mutator transaction binding the contract method 0xfedaa2a7.
//
// Solidity: function updateValidatorsAndThreshold((address,uint256)[] _validatorsToRemove, (address,string)[] _validatorsToAdd, uint256 _newThreshold) returns()
func (_Aggoraclecommittee *AggoraclecommitteeSession) UpdateValidatorsAndThreshold(_validatorsToRemove []IAggOracleCommitteeRemoveValidatorInfo, _validatorsToAdd []IAggOracleCommitteeValidatorInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.UpdateValidatorsAndThreshold(&_Aggoraclecommittee.TransactOpts, _validatorsToRemove, _validatorsToAdd, _newThreshold)
}

// UpdateValidatorsAndThreshold is a paid mutator transaction binding the contract method 0xfedaa2a7.
//
// Solidity: function updateValidatorsAndThreshold((address,uint256)[] _validatorsToRemove, (address,string)[] _validatorsToAdd, uint256 _newThreshold) returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactorSession) UpdateValidatorsAndThreshold(_validatorsToRemove []IAggOracleCommitteeRemoveValidatorInfo, _validatorsToAdd []IAggOracleCommitteeValidatorInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.UpdateValidatorsAndThreshold(&_Aggoraclecommittee.TransactOpts, _validatorsToRemove, _validatorsToAdd, _newThreshold)
}

// AggoraclecommitteeGERInjectedIterator is returned from FilterGERInjected and is used to iterate over the raw logs and unpacked data for GERInjected events raised by the Aggoraclecommittee contract.
type AggoraclecommitteeGERInjectedIterator struct {
	Event *AggoraclecommitteeGERInjected // Event containing the contract specifics and raw log

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
func (it *AggoraclecommitteeGERInjectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggoraclecommitteeGERInjected)
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
		it.Event = new(AggoraclecommitteeGERInjected)
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
func (it *AggoraclecommitteeGERInjectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggoraclecommitteeGERInjectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggoraclecommitteeGERInjected represents a GERInjected event raised by the Aggoraclecommittee contract.
type AggoraclecommitteeGERInjected struct {
	GlobalExitRoot [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGERInjected is a free log retrieval operation binding the contract event 0x45b91a7a09575a05b9dd504dff707cdeb7eda5f9d514491caf5951551b020e1e.
//
// Solidity: event GERInjected(bytes32 indexed globalExitRoot)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) FilterGERInjected(opts *bind.FilterOpts, globalExitRoot [][32]byte) (*AggoraclecommitteeGERInjectedIterator, error) {

	var globalExitRootRule []interface{}
	for _, globalExitRootItem := range globalExitRoot {
		globalExitRootRule = append(globalExitRootRule, globalExitRootItem)
	}

	logs, sub, err := _Aggoraclecommittee.contract.FilterLogs(opts, "GERInjected", globalExitRootRule)
	if err != nil {
		return nil, err
	}
	return &AggoraclecommitteeGERInjectedIterator{contract: _Aggoraclecommittee.contract, event: "GERInjected", logs: logs, sub: sub}, nil
}

// WatchGERInjected is a free log subscription operation binding the contract event 0x45b91a7a09575a05b9dd504dff707cdeb7eda5f9d514491caf5951551b020e1e.
//
// Solidity: event GERInjected(bytes32 indexed globalExitRoot)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) WatchGERInjected(opts *bind.WatchOpts, sink chan<- *AggoraclecommitteeGERInjected, globalExitRoot [][32]byte) (event.Subscription, error) {

	var globalExitRootRule []interface{}
	for _, globalExitRootItem := range globalExitRoot {
		globalExitRootRule = append(globalExitRootRule, globalExitRootItem)
	}

	logs, sub, err := _Aggoraclecommittee.contract.WatchLogs(opts, "GERInjected", globalExitRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggoraclecommitteeGERInjected)
				if err := _Aggoraclecommittee.contract.UnpackLog(event, "GERInjected", log); err != nil {
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

// ParseGERInjected is a log parse operation binding the contract event 0x45b91a7a09575a05b9dd504dff707cdeb7eda5f9d514491caf5951551b020e1e.
//
// Solidity: event GERInjected(bytes32 indexed globalExitRoot)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) ParseGERInjected(log types.Log) (*AggoraclecommitteeGERInjected, error) {
	event := new(AggoraclecommitteeGERInjected)
	if err := _Aggoraclecommittee.contract.UnpackLog(event, "GERInjected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggoraclecommitteeGERMarkedAsInjectedIterator is returned from FilterGERMarkedAsInjected and is used to iterate over the raw logs and unpacked data for GERMarkedAsInjected events raised by the Aggoraclecommittee contract.
type AggoraclecommitteeGERMarkedAsInjectedIterator struct {
	Event *AggoraclecommitteeGERMarkedAsInjected // Event containing the contract specifics and raw log

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
func (it *AggoraclecommitteeGERMarkedAsInjectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggoraclecommitteeGERMarkedAsInjected)
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
		it.Event = new(AggoraclecommitteeGERMarkedAsInjected)
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
func (it *AggoraclecommitteeGERMarkedAsInjectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggoraclecommitteeGERMarkedAsInjectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggoraclecommitteeGERMarkedAsInjected represents a GERMarkedAsInjected event raised by the Aggoraclecommittee contract.
type AggoraclecommitteeGERMarkedAsInjected struct {
	GlobalExitRoot [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGERMarkedAsInjected is a free log retrieval operation binding the contract event 0xcafdbc8c7e71f2990afc9e787ff729a197666534c9a49959c1373c819b708a8f.
//
// Solidity: event GERMarkedAsInjected(bytes32 indexed globalExitRoot)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) FilterGERMarkedAsInjected(opts *bind.FilterOpts, globalExitRoot [][32]byte) (*AggoraclecommitteeGERMarkedAsInjectedIterator, error) {

	var globalExitRootRule []interface{}
	for _, globalExitRootItem := range globalExitRoot {
		globalExitRootRule = append(globalExitRootRule, globalExitRootItem)
	}

	logs, sub, err := _Aggoraclecommittee.contract.FilterLogs(opts, "GERMarkedAsInjected", globalExitRootRule)
	if err != nil {
		return nil, err
	}
	return &AggoraclecommitteeGERMarkedAsInjectedIterator{contract: _Aggoraclecommittee.contract, event: "GERMarkedAsInjected", logs: logs, sub: sub}, nil
}

// WatchGERMarkedAsInjected is a free log subscription operation binding the contract event 0xcafdbc8c7e71f2990afc9e787ff729a197666534c9a49959c1373c819b708a8f.
//
// Solidity: event GERMarkedAsInjected(bytes32 indexed globalExitRoot)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) WatchGERMarkedAsInjected(opts *bind.WatchOpts, sink chan<- *AggoraclecommitteeGERMarkedAsInjected, globalExitRoot [][32]byte) (event.Subscription, error) {

	var globalExitRootRule []interface{}
	for _, globalExitRootItem := range globalExitRoot {
		globalExitRootRule = append(globalExitRootRule, globalExitRootItem)
	}

	logs, sub, err := _Aggoraclecommittee.contract.WatchLogs(opts, "GERMarkedAsInjected", globalExitRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggoraclecommitteeGERMarkedAsInjected)
				if err := _Aggoraclecommittee.contract.UnpackLog(event, "GERMarkedAsInjected", log); err != nil {
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

// ParseGERMarkedAsInjected is a log parse operation binding the contract event 0xcafdbc8c7e71f2990afc9e787ff729a197666534c9a49959c1373c819b708a8f.
//
// Solidity: event GERMarkedAsInjected(bytes32 indexed globalExitRoot)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) ParseGERMarkedAsInjected(log types.Log) (*AggoraclecommitteeGERMarkedAsInjected, error) {
	event := new(AggoraclecommitteeGERMarkedAsInjected)
	if err := _Aggoraclecommittee.contract.UnpackLog(event, "GERMarkedAsInjected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggoraclecommitteeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Aggoraclecommittee contract.
type AggoraclecommitteeInitializedIterator struct {
	Event *AggoraclecommitteeInitialized // Event containing the contract specifics and raw log

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
func (it *AggoraclecommitteeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggoraclecommitteeInitialized)
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
		it.Event = new(AggoraclecommitteeInitialized)
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
func (it *AggoraclecommitteeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggoraclecommitteeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggoraclecommitteeInitialized represents a Initialized event raised by the Aggoraclecommittee contract.
type AggoraclecommitteeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) FilterInitialized(opts *bind.FilterOpts) (*AggoraclecommitteeInitializedIterator, error) {

	logs, sub, err := _Aggoraclecommittee.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AggoraclecommitteeInitializedIterator{contract: _Aggoraclecommittee.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AggoraclecommitteeInitialized) (event.Subscription, error) {

	logs, sub, err := _Aggoraclecommittee.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggoraclecommitteeInitialized)
				if err := _Aggoraclecommittee.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) ParseInitialized(log types.Log) (*AggoraclecommitteeInitialized, error) {
	event := new(AggoraclecommitteeInitialized)
	if err := _Aggoraclecommittee.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggoraclecommitteeOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Aggoraclecommittee contract.
type AggoraclecommitteeOwnershipTransferStartedIterator struct {
	Event *AggoraclecommitteeOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *AggoraclecommitteeOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggoraclecommitteeOwnershipTransferStarted)
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
		it.Event = new(AggoraclecommitteeOwnershipTransferStarted)
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
func (it *AggoraclecommitteeOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggoraclecommitteeOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggoraclecommitteeOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Aggoraclecommittee contract.
type AggoraclecommitteeOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AggoraclecommitteeOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Aggoraclecommittee.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AggoraclecommitteeOwnershipTransferStartedIterator{contract: _Aggoraclecommittee.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *AggoraclecommitteeOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Aggoraclecommittee.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggoraclecommitteeOwnershipTransferStarted)
				if err := _Aggoraclecommittee.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) ParseOwnershipTransferStarted(log types.Log) (*AggoraclecommitteeOwnershipTransferStarted, error) {
	event := new(AggoraclecommitteeOwnershipTransferStarted)
	if err := _Aggoraclecommittee.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggoraclecommitteeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Aggoraclecommittee contract.
type AggoraclecommitteeOwnershipTransferredIterator struct {
	Event *AggoraclecommitteeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AggoraclecommitteeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggoraclecommitteeOwnershipTransferred)
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
		it.Event = new(AggoraclecommitteeOwnershipTransferred)
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
func (it *AggoraclecommitteeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggoraclecommitteeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggoraclecommitteeOwnershipTransferred represents a OwnershipTransferred event raised by the Aggoraclecommittee contract.
type AggoraclecommitteeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AggoraclecommitteeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Aggoraclecommittee.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AggoraclecommitteeOwnershipTransferredIterator{contract: _Aggoraclecommittee.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AggoraclecommitteeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Aggoraclecommittee.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggoraclecommitteeOwnershipTransferred)
				if err := _Aggoraclecommittee.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) ParseOwnershipTransferred(log types.Log) (*AggoraclecommitteeOwnershipTransferred, error) {
	event := new(AggoraclecommitteeOwnershipTransferred)
	if err := _Aggoraclecommittee.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggoraclecommitteeUpdateAggOracleProposerIterator is returned from FilterUpdateAggOracleProposer and is used to iterate over the raw logs and unpacked data for UpdateAggOracleProposer events raised by the Aggoraclecommittee contract.
type AggoraclecommitteeUpdateAggOracleProposerIterator struct {
	Event *AggoraclecommitteeUpdateAggOracleProposer // Event containing the contract specifics and raw log

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
func (it *AggoraclecommitteeUpdateAggOracleProposerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggoraclecommitteeUpdateAggOracleProposer)
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
		it.Event = new(AggoraclecommitteeUpdateAggOracleProposer)
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
func (it *AggoraclecommitteeUpdateAggOracleProposerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggoraclecommitteeUpdateAggOracleProposerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggoraclecommitteeUpdateAggOracleProposer represents a UpdateAggOracleProposer event raised by the Aggoraclecommittee contract.
type AggoraclecommitteeUpdateAggOracleProposer struct {
	NewProposer common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateAggOracleProposer is a free log retrieval operation binding the contract event 0x93c536fad4968838f70b6a5841cc12ffbd1a657048ef449f715f0d9ebdbeffca.
//
// Solidity: event UpdateAggOracleProposer(address indexed newProposer)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) FilterUpdateAggOracleProposer(opts *bind.FilterOpts, newProposer []common.Address) (*AggoraclecommitteeUpdateAggOracleProposerIterator, error) {

	var newProposerRule []interface{}
	for _, newProposerItem := range newProposer {
		newProposerRule = append(newProposerRule, newProposerItem)
	}

	logs, sub, err := _Aggoraclecommittee.contract.FilterLogs(opts, "UpdateAggOracleProposer", newProposerRule)
	if err != nil {
		return nil, err
	}
	return &AggoraclecommitteeUpdateAggOracleProposerIterator{contract: _Aggoraclecommittee.contract, event: "UpdateAggOracleProposer", logs: logs, sub: sub}, nil
}

// WatchUpdateAggOracleProposer is a free log subscription operation binding the contract event 0x93c536fad4968838f70b6a5841cc12ffbd1a657048ef449f715f0d9ebdbeffca.
//
// Solidity: event UpdateAggOracleProposer(address indexed newProposer)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) WatchUpdateAggOracleProposer(opts *bind.WatchOpts, sink chan<- *AggoraclecommitteeUpdateAggOracleProposer, newProposer []common.Address) (event.Subscription, error) {

	var newProposerRule []interface{}
	for _, newProposerItem := range newProposer {
		newProposerRule = append(newProposerRule, newProposerItem)
	}

	logs, sub, err := _Aggoraclecommittee.contract.WatchLogs(opts, "UpdateAggOracleProposer", newProposerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggoraclecommitteeUpdateAggOracleProposer)
				if err := _Aggoraclecommittee.contract.UnpackLog(event, "UpdateAggOracleProposer", log); err != nil {
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

// ParseUpdateAggOracleProposer is a log parse operation binding the contract event 0x93c536fad4968838f70b6a5841cc12ffbd1a657048ef449f715f0d9ebdbeffca.
//
// Solidity: event UpdateAggOracleProposer(address indexed newProposer)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) ParseUpdateAggOracleProposer(log types.Log) (*AggoraclecommitteeUpdateAggOracleProposer, error) {
	event := new(AggoraclecommitteeUpdateAggOracleProposer)
	if err := _Aggoraclecommittee.contract.UnpackLog(event, "UpdateAggOracleProposer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggoraclecommitteeValidatorsAndThresholdUpdatedIterator is returned from FilterValidatorsAndThresholdUpdated and is used to iterate over the raw logs and unpacked data for ValidatorsAndThresholdUpdated events raised by the Aggoraclecommittee contract.
type AggoraclecommitteeValidatorsAndThresholdUpdatedIterator struct {
	Event *AggoraclecommitteeValidatorsAndThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *AggoraclecommitteeValidatorsAndThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggoraclecommitteeValidatorsAndThresholdUpdated)
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
		it.Event = new(AggoraclecommitteeValidatorsAndThresholdUpdated)
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
func (it *AggoraclecommitteeValidatorsAndThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggoraclecommitteeValidatorsAndThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggoraclecommitteeValidatorsAndThresholdUpdated represents a ValidatorsAndThresholdUpdated event raised by the Aggoraclecommittee contract.
type AggoraclecommitteeValidatorsAndThresholdUpdated struct {
	Validators []common.Address
	Threshold  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorsAndThresholdUpdated is a free log retrieval operation binding the contract event 0x50051c796f4289791fedb6e1f365497b6527df6f6f19527d566e3c9791af5460.
//
// Solidity: event ValidatorsAndThresholdUpdated(address[] validators, uint256 threshold)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) FilterValidatorsAndThresholdUpdated(opts *bind.FilterOpts) (*AggoraclecommitteeValidatorsAndThresholdUpdatedIterator, error) {

	logs, sub, err := _Aggoraclecommittee.contract.FilterLogs(opts, "ValidatorsAndThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &AggoraclecommitteeValidatorsAndThresholdUpdatedIterator{contract: _Aggoraclecommittee.contract, event: "ValidatorsAndThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorsAndThresholdUpdated is a free log subscription operation binding the contract event 0x50051c796f4289791fedb6e1f365497b6527df6f6f19527d566e3c9791af5460.
//
// Solidity: event ValidatorsAndThresholdUpdated(address[] validators, uint256 threshold)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) WatchValidatorsAndThresholdUpdated(opts *bind.WatchOpts, sink chan<- *AggoraclecommitteeValidatorsAndThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _Aggoraclecommittee.contract.WatchLogs(opts, "ValidatorsAndThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggoraclecommitteeValidatorsAndThresholdUpdated)
				if err := _Aggoraclecommittee.contract.UnpackLog(event, "ValidatorsAndThresholdUpdated", log); err != nil {
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

// ParseValidatorsAndThresholdUpdated is a log parse operation binding the contract event 0x50051c796f4289791fedb6e1f365497b6527df6f6f19527d566e3c9791af5460.
//
// Solidity: event ValidatorsAndThresholdUpdated(address[] validators, uint256 threshold)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) ParseValidatorsAndThresholdUpdated(log types.Log) (*AggoraclecommitteeValidatorsAndThresholdUpdated, error) {
	event := new(AggoraclecommitteeValidatorsAndThresholdUpdated)
	if err := _Aggoraclecommittee.contract.UnpackLog(event, "ValidatorsAndThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

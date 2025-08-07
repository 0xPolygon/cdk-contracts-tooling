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

// AggoraclecommitteeMetaData contains all meta data concerning the Aggoraclecommittee contract.
var AggoraclecommitteeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIGlobalExitRootManagerL2SovereignChain\",\"name\":\"globalExitRootManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyOracleMember\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyVotedForThisGER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProposedGER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOracleMember\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleMemberCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleMemberIndexMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleMemberNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QuorumCannotBeGreaterThanAggOracleMembers\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QuorumCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QuorumNotReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WasNotOracleMember\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOracleMember\",\"type\":\"address\"}],\"name\":\"AddAggOracleMember\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"consolidatedGlobalExitRoot\",\"type\":\"bytes32\"}],\"name\":\"ConsolidatedGlobalExitRoot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"proposedGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"}],\"name\":\"ProposedGlobalExitRoot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oracleMemberRemoved\",\"type\":\"address\"}],\"name\":\"RemoveAggOracleMember\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newQuorum\",\"type\":\"uint64\"}],\"name\":\"UpdateQuorum\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INITIAL_PROPOSED_GER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptGlobalExitRootUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOracleMember\",\"type\":\"address\"}],\"name\":\"addOracleMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addressToLastProposedGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"aggOracleMembers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"}],\"name\":\"consolidateGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleMember\",\"type\":\"address\"}],\"name\":\"getAggOracleMemberIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggOracleMembersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllAggOracleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManagerL2Sovereign\",\"outputs\":[{\"internalType\":\"contractIGlobalExitRootManagerL2SovereignChain\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_aggOracleMembers\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_quorum\",\"type\":\"uint64\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposedGlobalExitRoot\",\"type\":\"bytes32\"}],\"name\":\"proposeGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"proposedGERToReport\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"votes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorum\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleMemberAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"oracleMemberIndex\",\"type\":\"uint256\"}],\"name\":\"removeOracleMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGlobalExitRootUpdater\",\"type\":\"address\"}],\"name\":\"transferGlobalExitRootUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newQuorum\",\"type\":\"uint64\"}],\"name\":\"updateQuorum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f5ffd5b5060405161189238038061189283398101604081905261002e916100fb565b6001600160a01b038116608052610043610049565b50610128565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100995760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100f85780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b5f6020828403121561010b575f5ffd5b81516001600160a01b0381168114610121575f5ffd5b9392505050565b60805161173d6101555f395f81816103540152818161043a01528181610f6c01526112a6015261173d5ff3fe608060405234801561000f575f5ffd5b506004361061016e575f3560e01c8063542b56b3116100d25780639915c0d211610088578063c053902a11610063578063c053902a14610389578063f2fde38b14610391578063ffa1ad74146103a4575f5ffd5b80639915c0d21461033c578063a7a90e2a1461034f578063b164e43714610376575f5ffd5b8063715018a6116100b8578063715018a6146102e457806372312af8146102ec5780638da5cb5b146102ff575f5ffd5b8063542b56b3146102bc578063703d8d42146102cf575f5ffd5b806329218b611161012757806351146e101161010d57806351146e1014610247578063512bab0d146102a157806353985e5a146102a9575f5ffd5b806329218b611461021557806338fd8a0314610228575f5ffd5b8063145d1ed611610157578063145d1ed61461019d5780631703a018146101b05780631ffbaa74146101dd575f5ffd5b8063017e37d6146101725780630e1bbf9f14610188575b5f5ffd5b5f545b6040519081526020015b60405180910390f35b61019b61019636600461146f565b6103ed565b005b61019b6101ab3660046114a6565b610494565b6001546101c49067ffffffffffffffff1681565b60405167ffffffffffffffff909116815260200161017f565b6101f06101eb366004611538565b61074f565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161017f565b61019b61022336600461154f565b610783565b61017561023636600461146f565b60026020525f908152604090205481565b610280610255366004611538565b60036020525f908152604090205467ffffffffffffffff808216916801000000000000000090041682565b6040805167ffffffffffffffff93841681529290911660208301520161017f565b610175600181565b61019b6102b7366004611568565b61083c565b61019b6102ca366004611538565b610b22565b6102d7610da9565b60405161017f9190611590565b61019b610e15565b6101756102fa36600461146f565b610e28565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005473ffffffffffffffffffffffffffffffffffffffff166101f0565b61019b61034a366004611538565b610ec4565b6101f07f000000000000000000000000000000000000000000000000000000000000000081565b61019b61038436600461146f565b610f4e565b61019b610f62565b61019b61039f36600461146f565b610fe7565b6103e06040518060400160405280600681526020017f76312e302e30000000000000000000000000000000000000000000000000000081525081565b60405161017f91906115e8565b6103f561104c565b6040517f0e1bbf9f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301527f00000000000000000000000000000000000000000000000000000000000000001690630e1bbf9f906024015f604051808303815f87803b15801561047b575f5ffd5b505af115801561048d573d5f5f3e3d5ffd5b5050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f811580156104de5750825b90505f8267ffffffffffffffff1660011480156104fa5750303b155b905081158015610508575080155b1561053f576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156105a05784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b8567ffffffffffffffff165f036105e3576040517fb7c8d1cb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8616871015610627576040517ffeca976200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff88161790555f5b8781101561069c5761069489898381811061067a5761067a61163b565b905060200201602081019061068f919061146f565b6110da565b60010161065d565b506106a689611233565b60405167ffffffffffffffff871681527fb600f3cf7f38a4b49bb0c75f722ef69f7e3e39ef3bb4aa8207fd86e724a232499060200160405180910390a183156107445784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b5f818154811061075d575f80fd5b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b61078b61104c565b8067ffffffffffffffff165f036107ce576040517fb7c8d1cb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff83169081179091556040519081527fb600f3cf7f38a4b49bb0c75f722ef69f7e3e39ef3bb4aa8207fd86e724a23249906020015b60405180910390a150565b61084461104c565b73ffffffffffffffffffffffffffffffffffffffff82165f90815260026020526040902054806108a0576040517f9463fb0000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff165f83815481106108c9576108c961163b565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff1614610921576040517f9ddd5e0300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600181146109a3575f818152600360205260409020805467ffffffffffffffff16156109a15780547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000811667ffffffffffffffff9182167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019091161781555b505b73ffffffffffffffffffffffffffffffffffffffff83165f90815260026020526040812081905580546109d890600190611695565b815481106109e8576109e861163b565b5f918252602082200154815473ffffffffffffffffffffffffffffffffffffffff909116919084908110610a1e57610a1e61163b565b5f918252602082200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9390931692909217909155805480610a7a57610a7a6116ae565b5f8281526020908190207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908301810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590910190915560405173ffffffffffffffffffffffffffffffffffffffff851681527f396186e0feff55f3e52260aee4ec024318786c7f5b8b01639b1d42c24b59eafc910160405180910390a1505050565b60018114801590610b3257508015155b610b68576040517fa3c967af00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b335f9081526002602052604090205480610bae576040517f9cf1889e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b818103610be7576040517f5dc4100700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018114610c69575f818152600360205260409020805467ffffffffffffffff1615610c675780547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000811667ffffffffffffffff9182167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019091161781555b505b5f82815260036020908152604080832081518083019092525467ffffffffffffffff808216835268010000000000000000909104169181018290529103610cc35767ffffffffffffffff4216602082015260018152610cdd565b805181610ccf826116db565b67ffffffffffffffff169052505b604080518481523360208201527f9f298266555f7b22d45b8ea34a9c99cba101d1ccbc3ff6dc4c392b7c915dc3b5910160405180910390a1600154815167ffffffffffffffff918216911610610d3b57610d3683611244565b505050565b5f838152600360209081526040808320845181548487015167ffffffffffffffff90811668010000000000000000027fffffffffffffffffffffffffffffffff0000000000000000000000000000000090921692169190911717905533835260029091529020839055505050565b60605f805480602002602001604051908101604052809291908181526020018280548015610e0b57602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610de0575b5050505050905090565b610e1d61104c565b610e265f611343565b565b5f805b5f54811015610e91578273ffffffffffffffffffffffffffffffffffffffff165f8281548110610e5d57610e5d61163b565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff1603610e895792915050565b600101610e2b565b506040517f1cec48b600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8181526003602090815260409182902082518084019093525467ffffffffffffffff808216808552680100000000000000009092048116928401929092526001549091161115610f41576040517faa26a69300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f4a82611244565b5050565b610f5661104c565b610f5f816110da565b50565b610f6a61104c565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c053902a6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015610fcf575f5ffd5b505af1158015610fe1573d5f5f3e3d5ffd5b50505050565b610fef61104c565b73ffffffffffffffffffffffffffffffffffffffff8116611043576040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081525f60048201526024015b60405180910390fd5b610f5f81611343565b3361108b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614610e26576040517f118cdaa700000000000000000000000000000000000000000000000000000000815233600482015260240161103a565b73ffffffffffffffffffffffffffffffffffffffff8116611127576040517f3846812a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81165f9081526002602052604090205415611183576040517fb078ca8800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81165f818152600260209081526040808320600190819055835490810184559280527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56390920180547fffffffffffffffffffffffff0000000000000000000000000000000000000000168417905590519182527f59f8c2dd0bdc5787ef574566b9edcd7324a62b79ac952930f02679a1b25bbf6f9101610831565b61123b6113d8565b610f5f8161143f565b5f818152600360205260409081902080547fffffffffffffffffffffffffffffffff00000000000000000000000000000000169055517f12da06b2000000000000000000000000000000000000000000000000000000008152600481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906312da06b2906024015f604051808303815f87803b1580156112fc575f5ffd5b505af115801561130e573d5f5f3e3d5ffd5b505050507f1676a1944900a9899dac554c2da10ba62637fe8d08375ffc3fb8494c92918d988160405161083191815260200190565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080547fffffffffffffffffffffffff0000000000000000000000000000000000000000811673ffffffffffffffffffffffffffffffffffffffff848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16610e26576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610fef6113d8565b803573ffffffffffffffffffffffffffffffffffffffff8116811461146a575f5ffd5b919050565b5f6020828403121561147f575f5ffd5b61148882611447565b9392505050565b803567ffffffffffffffff8116811461146a575f5ffd5b5f5f5f5f606085870312156114b9575f5ffd5b6114c285611447565b9350602085013567ffffffffffffffff8111156114dd575f5ffd5b8501601f810187136114ed575f5ffd5b803567ffffffffffffffff811115611503575f5ffd5b8760208260051b8401011115611517575f5ffd5b6020919091019350915061152d6040860161148f565b905092959194509250565b5f60208284031215611548575f5ffd5b5035919050565b5f6020828403121561155f575f5ffd5b6114888261148f565b5f5f60408385031215611579575f5ffd5b61158283611447565b946020939093013593505050565b602080825282518282018190525f918401906040840190835b818110156115dd57835173ffffffffffffffffffffffffffffffffffffffff168352602093840193909201916001016115a9565b509095945050505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b818103818111156116a8576116a8611668565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b5f67ffffffffffffffff821667ffffffffffffffff81036116fe576116fe611668565b6001019291505056fea264697066735822122036d9b1f6c3c7177d25b8c375b3eeb8bc0ae9925a4d6d12d4c89e9fc162c84a8964736f6c634300081c0033",
}

// AggoraclecommitteeABI is the input ABI used to generate the binding from.
// Deprecated: Use AggoraclecommitteeMetaData.ABI instead.
var AggoraclecommitteeABI = AggoraclecommitteeMetaData.ABI

// AggoraclecommitteeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AggoraclecommitteeMetaData.Bin instead.
var AggoraclecommitteeBin = AggoraclecommitteeMetaData.Bin

// DeployAggoraclecommittee deploys a new Ethereum contract, binding an instance of Aggoraclecommittee to it.
func DeployAggoraclecommittee(auth *bind.TransactOpts, backend bind.ContractBackend, globalExitRootManager common.Address) (common.Address, *types.Transaction, *Aggoraclecommittee, error) {
	parsed, err := AggoraclecommitteeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AggoraclecommitteeBin), backend, globalExitRootManager)
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

// INITIALPROPOSEDGER is a free data retrieval call binding the contract method 0x512bab0d.
//
// Solidity: function INITIAL_PROPOSED_GER() view returns(bytes32)
func (_Aggoraclecommittee *AggoraclecommitteeCaller) INITIALPROPOSEDGER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "INITIAL_PROPOSED_GER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// INITIALPROPOSEDGER is a free data retrieval call binding the contract method 0x512bab0d.
//
// Solidity: function INITIAL_PROPOSED_GER() view returns(bytes32)
func (_Aggoraclecommittee *AggoraclecommitteeSession) INITIALPROPOSEDGER() ([32]byte, error) {
	return _Aggoraclecommittee.Contract.INITIALPROPOSEDGER(&_Aggoraclecommittee.CallOpts)
}

// INITIALPROPOSEDGER is a free data retrieval call binding the contract method 0x512bab0d.
//
// Solidity: function INITIAL_PROPOSED_GER() view returns(bytes32)
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) INITIALPROPOSEDGER() ([32]byte, error) {
	return _Aggoraclecommittee.Contract.INITIALPROPOSEDGER(&_Aggoraclecommittee.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Aggoraclecommittee *AggoraclecommitteeCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Aggoraclecommittee *AggoraclecommitteeSession) VERSION() (string, error) {
	return _Aggoraclecommittee.Contract.VERSION(&_Aggoraclecommittee.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) VERSION() (string, error) {
	return _Aggoraclecommittee.Contract.VERSION(&_Aggoraclecommittee.CallOpts)
}

// AddressToLastProposedGER is a free data retrieval call binding the contract method 0x38fd8a03.
//
// Solidity: function addressToLastProposedGER(address ) view returns(bytes32)
func (_Aggoraclecommittee *AggoraclecommitteeCaller) AddressToLastProposedGER(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "addressToLastProposedGER", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AddressToLastProposedGER is a free data retrieval call binding the contract method 0x38fd8a03.
//
// Solidity: function addressToLastProposedGER(address ) view returns(bytes32)
func (_Aggoraclecommittee *AggoraclecommitteeSession) AddressToLastProposedGER(arg0 common.Address) ([32]byte, error) {
	return _Aggoraclecommittee.Contract.AddressToLastProposedGER(&_Aggoraclecommittee.CallOpts, arg0)
}

// AddressToLastProposedGER is a free data retrieval call binding the contract method 0x38fd8a03.
//
// Solidity: function addressToLastProposedGER(address ) view returns(bytes32)
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) AddressToLastProposedGER(arg0 common.Address) ([32]byte, error) {
	return _Aggoraclecommittee.Contract.AddressToLastProposedGER(&_Aggoraclecommittee.CallOpts, arg0)
}

// AggOracleMembers is a free data retrieval call binding the contract method 0x1ffbaa74.
//
// Solidity: function aggOracleMembers(uint256 ) view returns(address)
func (_Aggoraclecommittee *AggoraclecommitteeCaller) AggOracleMembers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "aggOracleMembers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggOracleMembers is a free data retrieval call binding the contract method 0x1ffbaa74.
//
// Solidity: function aggOracleMembers(uint256 ) view returns(address)
func (_Aggoraclecommittee *AggoraclecommitteeSession) AggOracleMembers(arg0 *big.Int) (common.Address, error) {
	return _Aggoraclecommittee.Contract.AggOracleMembers(&_Aggoraclecommittee.CallOpts, arg0)
}

// AggOracleMembers is a free data retrieval call binding the contract method 0x1ffbaa74.
//
// Solidity: function aggOracleMembers(uint256 ) view returns(address)
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) AggOracleMembers(arg0 *big.Int) (common.Address, error) {
	return _Aggoraclecommittee.Contract.AggOracleMembers(&_Aggoraclecommittee.CallOpts, arg0)
}

// GetAggOracleMemberIndex is a free data retrieval call binding the contract method 0x72312af8.
//
// Solidity: function getAggOracleMemberIndex(address oracleMember) view returns(uint256)
func (_Aggoraclecommittee *AggoraclecommitteeCaller) GetAggOracleMemberIndex(opts *bind.CallOpts, oracleMember common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "getAggOracleMemberIndex", oracleMember)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAggOracleMemberIndex is a free data retrieval call binding the contract method 0x72312af8.
//
// Solidity: function getAggOracleMemberIndex(address oracleMember) view returns(uint256)
func (_Aggoraclecommittee *AggoraclecommitteeSession) GetAggOracleMemberIndex(oracleMember common.Address) (*big.Int, error) {
	return _Aggoraclecommittee.Contract.GetAggOracleMemberIndex(&_Aggoraclecommittee.CallOpts, oracleMember)
}

// GetAggOracleMemberIndex is a free data retrieval call binding the contract method 0x72312af8.
//
// Solidity: function getAggOracleMemberIndex(address oracleMember) view returns(uint256)
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) GetAggOracleMemberIndex(oracleMember common.Address) (*big.Int, error) {
	return _Aggoraclecommittee.Contract.GetAggOracleMemberIndex(&_Aggoraclecommittee.CallOpts, oracleMember)
}

// GetAggOracleMembersCount is a free data retrieval call binding the contract method 0x017e37d6.
//
// Solidity: function getAggOracleMembersCount() view returns(uint256)
func (_Aggoraclecommittee *AggoraclecommitteeCaller) GetAggOracleMembersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "getAggOracleMembersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAggOracleMembersCount is a free data retrieval call binding the contract method 0x017e37d6.
//
// Solidity: function getAggOracleMembersCount() view returns(uint256)
func (_Aggoraclecommittee *AggoraclecommitteeSession) GetAggOracleMembersCount() (*big.Int, error) {
	return _Aggoraclecommittee.Contract.GetAggOracleMembersCount(&_Aggoraclecommittee.CallOpts)
}

// GetAggOracleMembersCount is a free data retrieval call binding the contract method 0x017e37d6.
//
// Solidity: function getAggOracleMembersCount() view returns(uint256)
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) GetAggOracleMembersCount() (*big.Int, error) {
	return _Aggoraclecommittee.Contract.GetAggOracleMembersCount(&_Aggoraclecommittee.CallOpts)
}

// GetAllAggOracleMembers is a free data retrieval call binding the contract method 0x703d8d42.
//
// Solidity: function getAllAggOracleMembers() view returns(address[])
func (_Aggoraclecommittee *AggoraclecommitteeCaller) GetAllAggOracleMembers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "getAllAggOracleMembers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllAggOracleMembers is a free data retrieval call binding the contract method 0x703d8d42.
//
// Solidity: function getAllAggOracleMembers() view returns(address[])
func (_Aggoraclecommittee *AggoraclecommitteeSession) GetAllAggOracleMembers() ([]common.Address, error) {
	return _Aggoraclecommittee.Contract.GetAllAggOracleMembers(&_Aggoraclecommittee.CallOpts)
}

// GetAllAggOracleMembers is a free data retrieval call binding the contract method 0x703d8d42.
//
// Solidity: function getAllAggOracleMembers() view returns(address[])
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) GetAllAggOracleMembers() ([]common.Address, error) {
	return _Aggoraclecommittee.Contract.GetAllAggOracleMembers(&_Aggoraclecommittee.CallOpts)
}

// GlobalExitRootManagerL2Sovereign is a free data retrieval call binding the contract method 0xa7a90e2a.
//
// Solidity: function globalExitRootManagerL2Sovereign() view returns(address)
func (_Aggoraclecommittee *AggoraclecommitteeCaller) GlobalExitRootManagerL2Sovereign(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "globalExitRootManagerL2Sovereign")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManagerL2Sovereign is a free data retrieval call binding the contract method 0xa7a90e2a.
//
// Solidity: function globalExitRootManagerL2Sovereign() view returns(address)
func (_Aggoraclecommittee *AggoraclecommitteeSession) GlobalExitRootManagerL2Sovereign() (common.Address, error) {
	return _Aggoraclecommittee.Contract.GlobalExitRootManagerL2Sovereign(&_Aggoraclecommittee.CallOpts)
}

// GlobalExitRootManagerL2Sovereign is a free data retrieval call binding the contract method 0xa7a90e2a.
//
// Solidity: function globalExitRootManagerL2Sovereign() view returns(address)
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) GlobalExitRootManagerL2Sovereign() (common.Address, error) {
	return _Aggoraclecommittee.Contract.GlobalExitRootManagerL2Sovereign(&_Aggoraclecommittee.CallOpts)
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

// ProposedGERToReport is a free data retrieval call binding the contract method 0x51146e10.
//
// Solidity: function proposedGERToReport(bytes32 ) view returns(uint64 votes, uint64 timestamp)
func (_Aggoraclecommittee *AggoraclecommitteeCaller) ProposedGERToReport(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Votes     uint64
	Timestamp uint64
}, error) {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "proposedGERToReport", arg0)

	outstruct := new(struct {
		Votes     uint64
		Timestamp uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Votes = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// ProposedGERToReport is a free data retrieval call binding the contract method 0x51146e10.
//
// Solidity: function proposedGERToReport(bytes32 ) view returns(uint64 votes, uint64 timestamp)
func (_Aggoraclecommittee *AggoraclecommitteeSession) ProposedGERToReport(arg0 [32]byte) (struct {
	Votes     uint64
	Timestamp uint64
}, error) {
	return _Aggoraclecommittee.Contract.ProposedGERToReport(&_Aggoraclecommittee.CallOpts, arg0)
}

// ProposedGERToReport is a free data retrieval call binding the contract method 0x51146e10.
//
// Solidity: function proposedGERToReport(bytes32 ) view returns(uint64 votes, uint64 timestamp)
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) ProposedGERToReport(arg0 [32]byte) (struct {
	Votes     uint64
	Timestamp uint64
}, error) {
	return _Aggoraclecommittee.Contract.ProposedGERToReport(&_Aggoraclecommittee.CallOpts, arg0)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint64)
func (_Aggoraclecommittee *AggoraclecommitteeCaller) Quorum(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Aggoraclecommittee.contract.Call(opts, &out, "quorum")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint64)
func (_Aggoraclecommittee *AggoraclecommitteeSession) Quorum() (uint64, error) {
	return _Aggoraclecommittee.Contract.Quorum(&_Aggoraclecommittee.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint64)
func (_Aggoraclecommittee *AggoraclecommitteeCallerSession) Quorum() (uint64, error) {
	return _Aggoraclecommittee.Contract.Quorum(&_Aggoraclecommittee.CallOpts)
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

// AddOracleMember is a paid mutator transaction binding the contract method 0xb164e437.
//
// Solidity: function addOracleMember(address newOracleMember) returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactor) AddOracleMember(opts *bind.TransactOpts, newOracleMember common.Address) (*types.Transaction, error) {
	return _Aggoraclecommittee.contract.Transact(opts, "addOracleMember", newOracleMember)
}

// AddOracleMember is a paid mutator transaction binding the contract method 0xb164e437.
//
// Solidity: function addOracleMember(address newOracleMember) returns()
func (_Aggoraclecommittee *AggoraclecommitteeSession) AddOracleMember(newOracleMember common.Address) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.AddOracleMember(&_Aggoraclecommittee.TransactOpts, newOracleMember)
}

// AddOracleMember is a paid mutator transaction binding the contract method 0xb164e437.
//
// Solidity: function addOracleMember(address newOracleMember) returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactorSession) AddOracleMember(newOracleMember common.Address) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.AddOracleMember(&_Aggoraclecommittee.TransactOpts, newOracleMember)
}

// ConsolidateGlobalExitRoot is a paid mutator transaction binding the contract method 0x9915c0d2.
//
// Solidity: function consolidateGlobalExitRoot(bytes32 globalExitRoot) returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactor) ConsolidateGlobalExitRoot(opts *bind.TransactOpts, globalExitRoot [32]byte) (*types.Transaction, error) {
	return _Aggoraclecommittee.contract.Transact(opts, "consolidateGlobalExitRoot", globalExitRoot)
}

// ConsolidateGlobalExitRoot is a paid mutator transaction binding the contract method 0x9915c0d2.
//
// Solidity: function consolidateGlobalExitRoot(bytes32 globalExitRoot) returns()
func (_Aggoraclecommittee *AggoraclecommitteeSession) ConsolidateGlobalExitRoot(globalExitRoot [32]byte) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.ConsolidateGlobalExitRoot(&_Aggoraclecommittee.TransactOpts, globalExitRoot)
}

// ConsolidateGlobalExitRoot is a paid mutator transaction binding the contract method 0x9915c0d2.
//
// Solidity: function consolidateGlobalExitRoot(bytes32 globalExitRoot) returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactorSession) ConsolidateGlobalExitRoot(globalExitRoot [32]byte) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.ConsolidateGlobalExitRoot(&_Aggoraclecommittee.TransactOpts, globalExitRoot)
}

// Initialize is a paid mutator transaction binding the contract method 0x145d1ed6.
//
// Solidity: function initialize(address _owner, address[] _aggOracleMembers, uint64 _quorum) returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _aggOracleMembers []common.Address, _quorum uint64) (*types.Transaction, error) {
	return _Aggoraclecommittee.contract.Transact(opts, "initialize", _owner, _aggOracleMembers, _quorum)
}

// Initialize is a paid mutator transaction binding the contract method 0x145d1ed6.
//
// Solidity: function initialize(address _owner, address[] _aggOracleMembers, uint64 _quorum) returns()
func (_Aggoraclecommittee *AggoraclecommitteeSession) Initialize(_owner common.Address, _aggOracleMembers []common.Address, _quorum uint64) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.Initialize(&_Aggoraclecommittee.TransactOpts, _owner, _aggOracleMembers, _quorum)
}

// Initialize is a paid mutator transaction binding the contract method 0x145d1ed6.
//
// Solidity: function initialize(address _owner, address[] _aggOracleMembers, uint64 _quorum) returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactorSession) Initialize(_owner common.Address, _aggOracleMembers []common.Address, _quorum uint64) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.Initialize(&_Aggoraclecommittee.TransactOpts, _owner, _aggOracleMembers, _quorum)
}

// ProposeGlobalExitRoot is a paid mutator transaction binding the contract method 0x542b56b3.
//
// Solidity: function proposeGlobalExitRoot(bytes32 proposedGlobalExitRoot) returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactor) ProposeGlobalExitRoot(opts *bind.TransactOpts, proposedGlobalExitRoot [32]byte) (*types.Transaction, error) {
	return _Aggoraclecommittee.contract.Transact(opts, "proposeGlobalExitRoot", proposedGlobalExitRoot)
}

// ProposeGlobalExitRoot is a paid mutator transaction binding the contract method 0x542b56b3.
//
// Solidity: function proposeGlobalExitRoot(bytes32 proposedGlobalExitRoot) returns()
func (_Aggoraclecommittee *AggoraclecommitteeSession) ProposeGlobalExitRoot(proposedGlobalExitRoot [32]byte) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.ProposeGlobalExitRoot(&_Aggoraclecommittee.TransactOpts, proposedGlobalExitRoot)
}

// ProposeGlobalExitRoot is a paid mutator transaction binding the contract method 0x542b56b3.
//
// Solidity: function proposeGlobalExitRoot(bytes32 proposedGlobalExitRoot) returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactorSession) ProposeGlobalExitRoot(proposedGlobalExitRoot [32]byte) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.ProposeGlobalExitRoot(&_Aggoraclecommittee.TransactOpts, proposedGlobalExitRoot)
}

// RemoveOracleMember is a paid mutator transaction binding the contract method 0x53985e5a.
//
// Solidity: function removeOracleMember(address oracleMemberAddress, uint256 oracleMemberIndex) returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactor) RemoveOracleMember(opts *bind.TransactOpts, oracleMemberAddress common.Address, oracleMemberIndex *big.Int) (*types.Transaction, error) {
	return _Aggoraclecommittee.contract.Transact(opts, "removeOracleMember", oracleMemberAddress, oracleMemberIndex)
}

// RemoveOracleMember is a paid mutator transaction binding the contract method 0x53985e5a.
//
// Solidity: function removeOracleMember(address oracleMemberAddress, uint256 oracleMemberIndex) returns()
func (_Aggoraclecommittee *AggoraclecommitteeSession) RemoveOracleMember(oracleMemberAddress common.Address, oracleMemberIndex *big.Int) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.RemoveOracleMember(&_Aggoraclecommittee.TransactOpts, oracleMemberAddress, oracleMemberIndex)
}

// RemoveOracleMember is a paid mutator transaction binding the contract method 0x53985e5a.
//
// Solidity: function removeOracleMember(address oracleMemberAddress, uint256 oracleMemberIndex) returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactorSession) RemoveOracleMember(oracleMemberAddress common.Address, oracleMemberIndex *big.Int) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.RemoveOracleMember(&_Aggoraclecommittee.TransactOpts, oracleMemberAddress, oracleMemberIndex)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggoraclecommittee.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aggoraclecommittee *AggoraclecommitteeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.RenounceOwnership(&_Aggoraclecommittee.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.RenounceOwnership(&_Aggoraclecommittee.TransactOpts)
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

// UpdateQuorum is a paid mutator transaction binding the contract method 0x29218b61.
//
// Solidity: function updateQuorum(uint64 newQuorum) returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactor) UpdateQuorum(opts *bind.TransactOpts, newQuorum uint64) (*types.Transaction, error) {
	return _Aggoraclecommittee.contract.Transact(opts, "updateQuorum", newQuorum)
}

// UpdateQuorum is a paid mutator transaction binding the contract method 0x29218b61.
//
// Solidity: function updateQuorum(uint64 newQuorum) returns()
func (_Aggoraclecommittee *AggoraclecommitteeSession) UpdateQuorum(newQuorum uint64) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.UpdateQuorum(&_Aggoraclecommittee.TransactOpts, newQuorum)
}

// UpdateQuorum is a paid mutator transaction binding the contract method 0x29218b61.
//
// Solidity: function updateQuorum(uint64 newQuorum) returns()
func (_Aggoraclecommittee *AggoraclecommitteeTransactorSession) UpdateQuorum(newQuorum uint64) (*types.Transaction, error) {
	return _Aggoraclecommittee.Contract.UpdateQuorum(&_Aggoraclecommittee.TransactOpts, newQuorum)
}

// AggoraclecommitteeAddAggOracleMemberIterator is returned from FilterAddAggOracleMember and is used to iterate over the raw logs and unpacked data for AddAggOracleMember events raised by the Aggoraclecommittee contract.
type AggoraclecommitteeAddAggOracleMemberIterator struct {
	Event *AggoraclecommitteeAddAggOracleMember // Event containing the contract specifics and raw log

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
func (it *AggoraclecommitteeAddAggOracleMemberIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggoraclecommitteeAddAggOracleMember)
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
		it.Event = new(AggoraclecommitteeAddAggOracleMember)
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
func (it *AggoraclecommitteeAddAggOracleMemberIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggoraclecommitteeAddAggOracleMemberIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggoraclecommitteeAddAggOracleMember represents a AddAggOracleMember event raised by the Aggoraclecommittee contract.
type AggoraclecommitteeAddAggOracleMember struct {
	NewOracleMember common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddAggOracleMember is a free log retrieval operation binding the contract event 0x59f8c2dd0bdc5787ef574566b9edcd7324a62b79ac952930f02679a1b25bbf6f.
//
// Solidity: event AddAggOracleMember(address newOracleMember)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) FilterAddAggOracleMember(opts *bind.FilterOpts) (*AggoraclecommitteeAddAggOracleMemberIterator, error) {

	logs, sub, err := _Aggoraclecommittee.contract.FilterLogs(opts, "AddAggOracleMember")
	if err != nil {
		return nil, err
	}
	return &AggoraclecommitteeAddAggOracleMemberIterator{contract: _Aggoraclecommittee.contract, event: "AddAggOracleMember", logs: logs, sub: sub}, nil
}

// WatchAddAggOracleMember is a free log subscription operation binding the contract event 0x59f8c2dd0bdc5787ef574566b9edcd7324a62b79ac952930f02679a1b25bbf6f.
//
// Solidity: event AddAggOracleMember(address newOracleMember)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) WatchAddAggOracleMember(opts *bind.WatchOpts, sink chan<- *AggoraclecommitteeAddAggOracleMember) (event.Subscription, error) {

	logs, sub, err := _Aggoraclecommittee.contract.WatchLogs(opts, "AddAggOracleMember")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggoraclecommitteeAddAggOracleMember)
				if err := _Aggoraclecommittee.contract.UnpackLog(event, "AddAggOracleMember", log); err != nil {
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

// ParseAddAggOracleMember is a log parse operation binding the contract event 0x59f8c2dd0bdc5787ef574566b9edcd7324a62b79ac952930f02679a1b25bbf6f.
//
// Solidity: event AddAggOracleMember(address newOracleMember)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) ParseAddAggOracleMember(log types.Log) (*AggoraclecommitteeAddAggOracleMember, error) {
	event := new(AggoraclecommitteeAddAggOracleMember)
	if err := _Aggoraclecommittee.contract.UnpackLog(event, "AddAggOracleMember", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggoraclecommitteeConsolidatedGlobalExitRootIterator is returned from FilterConsolidatedGlobalExitRoot and is used to iterate over the raw logs and unpacked data for ConsolidatedGlobalExitRoot events raised by the Aggoraclecommittee contract.
type AggoraclecommitteeConsolidatedGlobalExitRootIterator struct {
	Event *AggoraclecommitteeConsolidatedGlobalExitRoot // Event containing the contract specifics and raw log

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
func (it *AggoraclecommitteeConsolidatedGlobalExitRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggoraclecommitteeConsolidatedGlobalExitRoot)
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
		it.Event = new(AggoraclecommitteeConsolidatedGlobalExitRoot)
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
func (it *AggoraclecommitteeConsolidatedGlobalExitRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggoraclecommitteeConsolidatedGlobalExitRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggoraclecommitteeConsolidatedGlobalExitRoot represents a ConsolidatedGlobalExitRoot event raised by the Aggoraclecommittee contract.
type AggoraclecommitteeConsolidatedGlobalExitRoot struct {
	ConsolidatedGlobalExitRoot [32]byte
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterConsolidatedGlobalExitRoot is a free log retrieval operation binding the contract event 0x1676a1944900a9899dac554c2da10ba62637fe8d08375ffc3fb8494c92918d98.
//
// Solidity: event ConsolidatedGlobalExitRoot(bytes32 consolidatedGlobalExitRoot)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) FilterConsolidatedGlobalExitRoot(opts *bind.FilterOpts) (*AggoraclecommitteeConsolidatedGlobalExitRootIterator, error) {

	logs, sub, err := _Aggoraclecommittee.contract.FilterLogs(opts, "ConsolidatedGlobalExitRoot")
	if err != nil {
		return nil, err
	}
	return &AggoraclecommitteeConsolidatedGlobalExitRootIterator{contract: _Aggoraclecommittee.contract, event: "ConsolidatedGlobalExitRoot", logs: logs, sub: sub}, nil
}

// WatchConsolidatedGlobalExitRoot is a free log subscription operation binding the contract event 0x1676a1944900a9899dac554c2da10ba62637fe8d08375ffc3fb8494c92918d98.
//
// Solidity: event ConsolidatedGlobalExitRoot(bytes32 consolidatedGlobalExitRoot)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) WatchConsolidatedGlobalExitRoot(opts *bind.WatchOpts, sink chan<- *AggoraclecommitteeConsolidatedGlobalExitRoot) (event.Subscription, error) {

	logs, sub, err := _Aggoraclecommittee.contract.WatchLogs(opts, "ConsolidatedGlobalExitRoot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggoraclecommitteeConsolidatedGlobalExitRoot)
				if err := _Aggoraclecommittee.contract.UnpackLog(event, "ConsolidatedGlobalExitRoot", log); err != nil {
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

// ParseConsolidatedGlobalExitRoot is a log parse operation binding the contract event 0x1676a1944900a9899dac554c2da10ba62637fe8d08375ffc3fb8494c92918d98.
//
// Solidity: event ConsolidatedGlobalExitRoot(bytes32 consolidatedGlobalExitRoot)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) ParseConsolidatedGlobalExitRoot(log types.Log) (*AggoraclecommitteeConsolidatedGlobalExitRoot, error) {
	event := new(AggoraclecommitteeConsolidatedGlobalExitRoot)
	if err := _Aggoraclecommittee.contract.UnpackLog(event, "ConsolidatedGlobalExitRoot", log); err != nil {
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

// AggoraclecommitteeProposedGlobalExitRootIterator is returned from FilterProposedGlobalExitRoot and is used to iterate over the raw logs and unpacked data for ProposedGlobalExitRoot events raised by the Aggoraclecommittee contract.
type AggoraclecommitteeProposedGlobalExitRootIterator struct {
	Event *AggoraclecommitteeProposedGlobalExitRoot // Event containing the contract specifics and raw log

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
func (it *AggoraclecommitteeProposedGlobalExitRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggoraclecommitteeProposedGlobalExitRoot)
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
		it.Event = new(AggoraclecommitteeProposedGlobalExitRoot)
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
func (it *AggoraclecommitteeProposedGlobalExitRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggoraclecommitteeProposedGlobalExitRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggoraclecommitteeProposedGlobalExitRoot represents a ProposedGlobalExitRoot event raised by the Aggoraclecommittee contract.
type AggoraclecommitteeProposedGlobalExitRoot struct {
	ProposedGlobalExitRoot [32]byte
	Proposer               common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterProposedGlobalExitRoot is a free log retrieval operation binding the contract event 0x9f298266555f7b22d45b8ea34a9c99cba101d1ccbc3ff6dc4c392b7c915dc3b5.
//
// Solidity: event ProposedGlobalExitRoot(bytes32 proposedGlobalExitRoot, address proposer)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) FilterProposedGlobalExitRoot(opts *bind.FilterOpts) (*AggoraclecommitteeProposedGlobalExitRootIterator, error) {

	logs, sub, err := _Aggoraclecommittee.contract.FilterLogs(opts, "ProposedGlobalExitRoot")
	if err != nil {
		return nil, err
	}
	return &AggoraclecommitteeProposedGlobalExitRootIterator{contract: _Aggoraclecommittee.contract, event: "ProposedGlobalExitRoot", logs: logs, sub: sub}, nil
}

// WatchProposedGlobalExitRoot is a free log subscription operation binding the contract event 0x9f298266555f7b22d45b8ea34a9c99cba101d1ccbc3ff6dc4c392b7c915dc3b5.
//
// Solidity: event ProposedGlobalExitRoot(bytes32 proposedGlobalExitRoot, address proposer)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) WatchProposedGlobalExitRoot(opts *bind.WatchOpts, sink chan<- *AggoraclecommitteeProposedGlobalExitRoot) (event.Subscription, error) {

	logs, sub, err := _Aggoraclecommittee.contract.WatchLogs(opts, "ProposedGlobalExitRoot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggoraclecommitteeProposedGlobalExitRoot)
				if err := _Aggoraclecommittee.contract.UnpackLog(event, "ProposedGlobalExitRoot", log); err != nil {
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

// ParseProposedGlobalExitRoot is a log parse operation binding the contract event 0x9f298266555f7b22d45b8ea34a9c99cba101d1ccbc3ff6dc4c392b7c915dc3b5.
//
// Solidity: event ProposedGlobalExitRoot(bytes32 proposedGlobalExitRoot, address proposer)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) ParseProposedGlobalExitRoot(log types.Log) (*AggoraclecommitteeProposedGlobalExitRoot, error) {
	event := new(AggoraclecommitteeProposedGlobalExitRoot)
	if err := _Aggoraclecommittee.contract.UnpackLog(event, "ProposedGlobalExitRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggoraclecommitteeRemoveAggOracleMemberIterator is returned from FilterRemoveAggOracleMember and is used to iterate over the raw logs and unpacked data for RemoveAggOracleMember events raised by the Aggoraclecommittee contract.
type AggoraclecommitteeRemoveAggOracleMemberIterator struct {
	Event *AggoraclecommitteeRemoveAggOracleMember // Event containing the contract specifics and raw log

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
func (it *AggoraclecommitteeRemoveAggOracleMemberIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggoraclecommitteeRemoveAggOracleMember)
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
		it.Event = new(AggoraclecommitteeRemoveAggOracleMember)
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
func (it *AggoraclecommitteeRemoveAggOracleMemberIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggoraclecommitteeRemoveAggOracleMemberIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggoraclecommitteeRemoveAggOracleMember represents a RemoveAggOracleMember event raised by the Aggoraclecommittee contract.
type AggoraclecommitteeRemoveAggOracleMember struct {
	OracleMemberRemoved common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRemoveAggOracleMember is a free log retrieval operation binding the contract event 0x396186e0feff55f3e52260aee4ec024318786c7f5b8b01639b1d42c24b59eafc.
//
// Solidity: event RemoveAggOracleMember(address oracleMemberRemoved)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) FilterRemoveAggOracleMember(opts *bind.FilterOpts) (*AggoraclecommitteeRemoveAggOracleMemberIterator, error) {

	logs, sub, err := _Aggoraclecommittee.contract.FilterLogs(opts, "RemoveAggOracleMember")
	if err != nil {
		return nil, err
	}
	return &AggoraclecommitteeRemoveAggOracleMemberIterator{contract: _Aggoraclecommittee.contract, event: "RemoveAggOracleMember", logs: logs, sub: sub}, nil
}

// WatchRemoveAggOracleMember is a free log subscription operation binding the contract event 0x396186e0feff55f3e52260aee4ec024318786c7f5b8b01639b1d42c24b59eafc.
//
// Solidity: event RemoveAggOracleMember(address oracleMemberRemoved)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) WatchRemoveAggOracleMember(opts *bind.WatchOpts, sink chan<- *AggoraclecommitteeRemoveAggOracleMember) (event.Subscription, error) {

	logs, sub, err := _Aggoraclecommittee.contract.WatchLogs(opts, "RemoveAggOracleMember")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggoraclecommitteeRemoveAggOracleMember)
				if err := _Aggoraclecommittee.contract.UnpackLog(event, "RemoveAggOracleMember", log); err != nil {
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

// ParseRemoveAggOracleMember is a log parse operation binding the contract event 0x396186e0feff55f3e52260aee4ec024318786c7f5b8b01639b1d42c24b59eafc.
//
// Solidity: event RemoveAggOracleMember(address oracleMemberRemoved)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) ParseRemoveAggOracleMember(log types.Log) (*AggoraclecommitteeRemoveAggOracleMember, error) {
	event := new(AggoraclecommitteeRemoveAggOracleMember)
	if err := _Aggoraclecommittee.contract.UnpackLog(event, "RemoveAggOracleMember", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggoraclecommitteeUpdateQuorumIterator is returned from FilterUpdateQuorum and is used to iterate over the raw logs and unpacked data for UpdateQuorum events raised by the Aggoraclecommittee contract.
type AggoraclecommitteeUpdateQuorumIterator struct {
	Event *AggoraclecommitteeUpdateQuorum // Event containing the contract specifics and raw log

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
func (it *AggoraclecommitteeUpdateQuorumIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggoraclecommitteeUpdateQuorum)
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
		it.Event = new(AggoraclecommitteeUpdateQuorum)
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
func (it *AggoraclecommitteeUpdateQuorumIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggoraclecommitteeUpdateQuorumIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggoraclecommitteeUpdateQuorum represents a UpdateQuorum event raised by the Aggoraclecommittee contract.
type AggoraclecommitteeUpdateQuorum struct {
	NewQuorum uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateQuorum is a free log retrieval operation binding the contract event 0xb600f3cf7f38a4b49bb0c75f722ef69f7e3e39ef3bb4aa8207fd86e724a23249.
//
// Solidity: event UpdateQuorum(uint64 newQuorum)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) FilterUpdateQuorum(opts *bind.FilterOpts) (*AggoraclecommitteeUpdateQuorumIterator, error) {

	logs, sub, err := _Aggoraclecommittee.contract.FilterLogs(opts, "UpdateQuorum")
	if err != nil {
		return nil, err
	}
	return &AggoraclecommitteeUpdateQuorumIterator{contract: _Aggoraclecommittee.contract, event: "UpdateQuorum", logs: logs, sub: sub}, nil
}

// WatchUpdateQuorum is a free log subscription operation binding the contract event 0xb600f3cf7f38a4b49bb0c75f722ef69f7e3e39ef3bb4aa8207fd86e724a23249.
//
// Solidity: event UpdateQuorum(uint64 newQuorum)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) WatchUpdateQuorum(opts *bind.WatchOpts, sink chan<- *AggoraclecommitteeUpdateQuorum) (event.Subscription, error) {

	logs, sub, err := _Aggoraclecommittee.contract.WatchLogs(opts, "UpdateQuorum")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggoraclecommitteeUpdateQuorum)
				if err := _Aggoraclecommittee.contract.UnpackLog(event, "UpdateQuorum", log); err != nil {
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

// ParseUpdateQuorum is a log parse operation binding the contract event 0xb600f3cf7f38a4b49bb0c75f722ef69f7e3e39ef3bb4aa8207fd86e724a23249.
//
// Solidity: event UpdateQuorum(uint64 newQuorum)
func (_Aggoraclecommittee *AggoraclecommitteeFilterer) ParseUpdateQuorum(log types.Log) (*AggoraclecommitteeUpdateQuorum, error) {
	event := new(AggoraclecommitteeUpdateQuorum)
	if err := _Aggoraclecommittee.contract.UnpackLog(event, "UpdateQuorum", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

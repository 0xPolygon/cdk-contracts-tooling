// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aggoraclemanager

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

// AggoraclemanagerMetaData contains all meta data concerning the Aggoraclemanager contract.
var AggoraclemanagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractGlobalExitRootManagerL2SovereignChain\",\"name\":\"globalExitRootManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyOracleMember\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProposedGER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOracleMember\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleMemberCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleMemberIndexMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleMemberNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QuorumCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WasNotOracleMember\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOracleMember\",\"type\":\"address\"}],\"name\":\"AddAggOracleMember\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"consolidatedGlobalExitRoot\",\"type\":\"bytes32\"}],\"name\":\"ConsolidatedGlobalExitRoot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"proposedGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"}],\"name\":\"ProposedGlobalExitRoot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oracleMemberRemoved\",\"type\":\"address\"}],\"name\":\"RemoveAggOracleMember\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newQuorum\",\"type\":\"uint64\"}],\"name\":\"UpdateQuorum\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INITIAL_PROPOSED_GER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptGlobalExitRootUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOracleMember\",\"type\":\"address\"}],\"name\":\"addOracleMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addressToLastProposedGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"aggOracleMembers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleMember\",\"type\":\"address\"}],\"name\":\"getAggOracleMemberIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggOracleMembersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllAggOracleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManagerL2Sovereign\",\"outputs\":[{\"internalType\":\"contractGlobalExitRootManagerL2SovereignChain\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_aggOracleMembers\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_quorum\",\"type\":\"uint64\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposedGlobalExitRoot\",\"type\":\"bytes32\"}],\"name\":\"proposeGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"proposedGERToReport\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"votes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorum\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleMemberAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"oracleMemberIndex\",\"type\":\"uint256\"}],\"name\":\"removeOracleMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGlobalExitRootUpdater\",\"type\":\"address\"}],\"name\":\"transferGlobalExitRootUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newQuorum\",\"type\":\"uint64\"}],\"name\":\"updateQuorum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f5ffd5b506040516116aa3803806116aa83398101604081905261002e916100fb565b6001600160a01b038116608052610043610049565b50610128565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100995760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100f85780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b5f6020828403121561010b575f5ffd5b81516001600160a01b0381168114610121575f5ffd5b9392505050565b6080516115556101555f395f818161031c015281816103b901528181610c920152610ed601526115555ff3fe608060405234801561000f575f5ffd5b5060043610610149575f3560e01c806353985e5a116100c75780638da5cb5b1161007d578063b164e43711610063578063b164e4371461033e578063c053902a14610351578063f2fde38b14610359575f5ffd5b80638da5cb5b146102da578063a7a90e2a14610317575f5ffd5b8063703d8d42116100ad578063703d8d42146102aa578063715018a6146102bf57806372312af8146102c7575f5ffd5b806353985e5a14610284578063542b56b314610297575f5ffd5b80631ffbaa741161011c57806338fd8a031161010257806338fd8a031461020357806351146e1014610222578063512bab0d1461027c575f5ffd5b80631ffbaa74146101b857806329218b61146101f0575f5ffd5b8063017e37d61461014d5780630e1bbf9f14610163578063145d1ed6146101785780631703a0181461018b575b5f5ffd5b5f545b6040519081526020015b60405180910390f35b6101766101713660046112da565b61036c565b005b610176610186366004611311565b610413565b60015461019f9067ffffffffffffffff1681565b60405167ffffffffffffffff909116815260200161015a565b6101cb6101c63660046113a3565b61068a565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161015a565b6101766101fe3660046113ba565b6106be565b6101506102113660046112da565b60026020525f908152604090205481565b61025b6102303660046113a3565b60036020525f908152604090205467ffffffffffffffff808216916801000000000000000090041682565b6040805167ffffffffffffffff93841681529290911660208301520161015a565b610150600181565b6101766102923660046113d3565b610777565b6101766102a53660046113a3565b610a5e565b6102b2610d9d565b60405161015a91906113fb565b610176610e09565b6101506102d53660046112da565b610e1c565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005473ffffffffffffffffffffffffffffffffffffffff166101cb565b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b61017661034c3660046112da565b610eb8565b610176610ecc565b6101766103673660046112da565b610f51565b610374610fb6565b6040517f0e1bbf9f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301527f00000000000000000000000000000000000000000000000000000000000000001690630e1bbf9f906024015f604051808303815f87803b1580156103fa575f5ffd5b505af115801561040c573d5f5f3e3d5ffd5b5050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f8115801561045d5750825b90505f8267ffffffffffffffff1660011480156104795750303b155b905081158015610487575080155b156104be576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000166001178555831561051f5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b8567ffffffffffffffff165f03610562576040517fb7c8d1cb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff88161790555f5b878110156105d7576105cf8989838181106105b5576105b5611453565b90506020020160208101906105ca91906112da565b611044565b600101610598565b506105e18961119d565b60405167ffffffffffffffff871681527fb600f3cf7f38a4b49bb0c75f722ef69f7e3e39ef3bb4aa8207fd86e724a232499060200160405180910390a1831561067f5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b5f8181548110610698575f80fd5b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b6106c6610fb6565b8067ffffffffffffffff165f03610709576040517fb7c8d1cb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff83169081179091556040519081527fb600f3cf7f38a4b49bb0c75f722ef69f7e3e39ef3bb4aa8207fd86e724a23249906020015b60405180910390a150565b61077f610fb6565b73ffffffffffffffffffffffffffffffffffffffff82165f90815260026020526040902054806107db576040517f9463fb0000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff165f838154811061080457610804611453565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff161461085c576040517f9ddd5e0300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600181146108de575f818152600360205260409020805467ffffffffffffffff16156108dc5780547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000811667ffffffffffffffff9182167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019091161781555b505b73ffffffffffffffffffffffffffffffffffffffff83165f9081526002602052604081208190558054610913906001906114ad565b8154811061092357610923611453565b5f918252602082200154815473ffffffffffffffffffffffffffffffffffffffff90911691908490811061095957610959611453565b5f918252602082200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff93909316929092179091558054806109b5576109b56114c6565b5f8281526020908190207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908301810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590910190915560405173ffffffffffffffffffffffffffffffffffffffff851681527f396186e0feff55f3e52260aee4ec024318786c7f5b8b01639b1d42c24b59eafc91015b60405180910390a1505050565b60018114801590610a6e57508015155b610aa4576040517fa3c967af00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b335f9081526002602052604090205480610aea576040517f9cf1889e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018114610b6c575f818152600360205260409020805467ffffffffffffffff1615610b6a5780547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000811667ffffffffffffffff9182167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019091161781555b505b5f82815260036020908152604080832081518083019092525467ffffffffffffffff808216835268010000000000000000909104169181018290529103610bc65767ffffffffffffffff4216602082015260018152610be0565b805181610bd2826114f3565b67ffffffffffffffff169052505b604080518481523360208201527f9f298266555f7b22d45b8ea34a9c99cba101d1ccbc3ff6dc4c392b7c915dc3b5910160405180910390a1600154815167ffffffffffffffff918216911610610d2f575f838152600360205260409081902080547fffffffffffffffffffffffffffffffff00000000000000000000000000000000169055517f12da06b2000000000000000000000000000000000000000000000000000000008152600481018490527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906312da06b2906024015f604051808303815f87803b158015610ce8575f5ffd5b505af1158015610cfa573d5f5f3e3d5ffd5b505050507f1676a1944900a9899dac554c2da10ba62637fe8d08375ffc3fb8494c92918d9883604051610a5191815260200190565b5f838152600360209081526040808320845181548487015167ffffffffffffffff90811668010000000000000000027fffffffffffffffffffffffffffffffff0000000000000000000000000000000090921692169190911717905533835260029091529020839055505050565b60605f805480602002602001604051908101604052809291908181526020018280548015610dff57602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610dd4575b5050505050905090565b610e11610fb6565b610e1a5f6111ae565b565b5f805b5f54811015610e85578273ffffffffffffffffffffffffffffffffffffffff165f8281548110610e5157610e51611453565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff1603610e7d5792915050565b600101610e1f565b506040517f1cec48b600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ec0610fb6565b610ec981611044565b50565b610ed4610fb6565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c053902a6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015610f39575f5ffd5b505af1158015610f4b573d5f5f3e3d5ffd5b50505050565b610f59610fb6565b73ffffffffffffffffffffffffffffffffffffffff8116610fad576040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081525f60048201526024015b60405180910390fd5b610ec9816111ae565b33610ff57f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614610e1a576040517f118cdaa7000000000000000000000000000000000000000000000000000000008152336004820152602401610fa4565b73ffffffffffffffffffffffffffffffffffffffff8116611091576040517f3846812a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81165f90815260026020526040902054156110ed576040517fb078ca8800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81165f818152600260209081526040808320600190819055835490810184559280527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56390920180547fffffffffffffffffffffffff0000000000000000000000000000000000000000168417905590519182527f59f8c2dd0bdc5787ef574566b9edcd7324a62b79ac952930f02679a1b25bbf6f910161076c565b6111a5611243565b610ec9816112aa565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080547fffffffffffffffffffffffff0000000000000000000000000000000000000000811673ffffffffffffffffffffffffffffffffffffffff848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16610e1a576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f59611243565b803573ffffffffffffffffffffffffffffffffffffffff811681146112d5575f5ffd5b919050565b5f602082840312156112ea575f5ffd5b6112f3826112b2565b9392505050565b803567ffffffffffffffff811681146112d5575f5ffd5b5f5f5f5f60608587031215611324575f5ffd5b61132d856112b2565b9350602085013567ffffffffffffffff811115611348575f5ffd5b8501601f81018713611358575f5ffd5b803567ffffffffffffffff81111561136e575f5ffd5b8760208260051b8401011115611382575f5ffd5b60209190910193509150611398604086016112fa565b905092959194509250565b5f602082840312156113b3575f5ffd5b5035919050565b5f602082840312156113ca575f5ffd5b6112f3826112fa565b5f5f604083850312156113e4575f5ffd5b6113ed836112b2565b946020939093013593505050565b602080825282518282018190525f918401906040840190835b8181101561144857835173ffffffffffffffffffffffffffffffffffffffff16835260209384019390920191600101611414565b509095945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b818103818111156114c0576114c0611480565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b5f67ffffffffffffffff821667ffffffffffffffff810361151657611516611480565b6001019291505056fea264697066735822122066607f20983ae5a6bcefc7cb66c059e2e3c119bfbb3f71eec0363c4c3f9efdd864736f6c634300081c0033",
}

// AggoraclemanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use AggoraclemanagerMetaData.ABI instead.
var AggoraclemanagerABI = AggoraclemanagerMetaData.ABI

// AggoraclemanagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AggoraclemanagerMetaData.Bin instead.
var AggoraclemanagerBin = AggoraclemanagerMetaData.Bin

// DeployAggoraclemanager deploys a new Ethereum contract, binding an instance of Aggoraclemanager to it.
func DeployAggoraclemanager(auth *bind.TransactOpts, backend bind.ContractBackend, globalExitRootManager common.Address) (common.Address, *types.Transaction, *Aggoraclemanager, error) {
	parsed, err := AggoraclemanagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AggoraclemanagerBin), backend, globalExitRootManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Aggoraclemanager{AggoraclemanagerCaller: AggoraclemanagerCaller{contract: contract}, AggoraclemanagerTransactor: AggoraclemanagerTransactor{contract: contract}, AggoraclemanagerFilterer: AggoraclemanagerFilterer{contract: contract}}, nil
}

// Aggoraclemanager is an auto generated Go binding around an Ethereum contract.
type Aggoraclemanager struct {
	AggoraclemanagerCaller     // Read-only binding to the contract
	AggoraclemanagerTransactor // Write-only binding to the contract
	AggoraclemanagerFilterer   // Log filterer for contract events
}

// AggoraclemanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AggoraclemanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggoraclemanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AggoraclemanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggoraclemanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AggoraclemanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggoraclemanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AggoraclemanagerSession struct {
	Contract     *Aggoraclemanager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AggoraclemanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AggoraclemanagerCallerSession struct {
	Contract *AggoraclemanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// AggoraclemanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AggoraclemanagerTransactorSession struct {
	Contract     *AggoraclemanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// AggoraclemanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AggoraclemanagerRaw struct {
	Contract *Aggoraclemanager // Generic contract binding to access the raw methods on
}

// AggoraclemanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AggoraclemanagerCallerRaw struct {
	Contract *AggoraclemanagerCaller // Generic read-only contract binding to access the raw methods on
}

// AggoraclemanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AggoraclemanagerTransactorRaw struct {
	Contract *AggoraclemanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAggoraclemanager creates a new instance of Aggoraclemanager, bound to a specific deployed contract.
func NewAggoraclemanager(address common.Address, backend bind.ContractBackend) (*Aggoraclemanager, error) {
	contract, err := bindAggoraclemanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aggoraclemanager{AggoraclemanagerCaller: AggoraclemanagerCaller{contract: contract}, AggoraclemanagerTransactor: AggoraclemanagerTransactor{contract: contract}, AggoraclemanagerFilterer: AggoraclemanagerFilterer{contract: contract}}, nil
}

// NewAggoraclemanagerCaller creates a new read-only instance of Aggoraclemanager, bound to a specific deployed contract.
func NewAggoraclemanagerCaller(address common.Address, caller bind.ContractCaller) (*AggoraclemanagerCaller, error) {
	contract, err := bindAggoraclemanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggoraclemanagerCaller{contract: contract}, nil
}

// NewAggoraclemanagerTransactor creates a new write-only instance of Aggoraclemanager, bound to a specific deployed contract.
func NewAggoraclemanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*AggoraclemanagerTransactor, error) {
	contract, err := bindAggoraclemanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggoraclemanagerTransactor{contract: contract}, nil
}

// NewAggoraclemanagerFilterer creates a new log filterer instance of Aggoraclemanager, bound to a specific deployed contract.
func NewAggoraclemanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*AggoraclemanagerFilterer, error) {
	contract, err := bindAggoraclemanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggoraclemanagerFilterer{contract: contract}, nil
}

// bindAggoraclemanager binds a generic wrapper to an already deployed contract.
func bindAggoraclemanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AggoraclemanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggoraclemanager *AggoraclemanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aggoraclemanager.Contract.AggoraclemanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggoraclemanager *AggoraclemanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggoraclemanager.Contract.AggoraclemanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggoraclemanager *AggoraclemanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggoraclemanager.Contract.AggoraclemanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggoraclemanager *AggoraclemanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aggoraclemanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggoraclemanager *AggoraclemanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggoraclemanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggoraclemanager *AggoraclemanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggoraclemanager.Contract.contract.Transact(opts, method, params...)
}

// INITIALPROPOSEDGER is a free data retrieval call binding the contract method 0x512bab0d.
//
// Solidity: function INITIAL_PROPOSED_GER() view returns(bytes32)
func (_Aggoraclemanager *AggoraclemanagerCaller) INITIALPROPOSEDGER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aggoraclemanager.contract.Call(opts, &out, "INITIAL_PROPOSED_GER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// INITIALPROPOSEDGER is a free data retrieval call binding the contract method 0x512bab0d.
//
// Solidity: function INITIAL_PROPOSED_GER() view returns(bytes32)
func (_Aggoraclemanager *AggoraclemanagerSession) INITIALPROPOSEDGER() ([32]byte, error) {
	return _Aggoraclemanager.Contract.INITIALPROPOSEDGER(&_Aggoraclemanager.CallOpts)
}

// INITIALPROPOSEDGER is a free data retrieval call binding the contract method 0x512bab0d.
//
// Solidity: function INITIAL_PROPOSED_GER() view returns(bytes32)
func (_Aggoraclemanager *AggoraclemanagerCallerSession) INITIALPROPOSEDGER() ([32]byte, error) {
	return _Aggoraclemanager.Contract.INITIALPROPOSEDGER(&_Aggoraclemanager.CallOpts)
}

// AddressToLastProposedGER is a free data retrieval call binding the contract method 0x38fd8a03.
//
// Solidity: function addressToLastProposedGER(address ) view returns(bytes32)
func (_Aggoraclemanager *AggoraclemanagerCaller) AddressToLastProposedGER(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Aggoraclemanager.contract.Call(opts, &out, "addressToLastProposedGER", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AddressToLastProposedGER is a free data retrieval call binding the contract method 0x38fd8a03.
//
// Solidity: function addressToLastProposedGER(address ) view returns(bytes32)
func (_Aggoraclemanager *AggoraclemanagerSession) AddressToLastProposedGER(arg0 common.Address) ([32]byte, error) {
	return _Aggoraclemanager.Contract.AddressToLastProposedGER(&_Aggoraclemanager.CallOpts, arg0)
}

// AddressToLastProposedGER is a free data retrieval call binding the contract method 0x38fd8a03.
//
// Solidity: function addressToLastProposedGER(address ) view returns(bytes32)
func (_Aggoraclemanager *AggoraclemanagerCallerSession) AddressToLastProposedGER(arg0 common.Address) ([32]byte, error) {
	return _Aggoraclemanager.Contract.AddressToLastProposedGER(&_Aggoraclemanager.CallOpts, arg0)
}

// AggOracleMembers is a free data retrieval call binding the contract method 0x1ffbaa74.
//
// Solidity: function aggOracleMembers(uint256 ) view returns(address)
func (_Aggoraclemanager *AggoraclemanagerCaller) AggOracleMembers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Aggoraclemanager.contract.Call(opts, &out, "aggOracleMembers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggOracleMembers is a free data retrieval call binding the contract method 0x1ffbaa74.
//
// Solidity: function aggOracleMembers(uint256 ) view returns(address)
func (_Aggoraclemanager *AggoraclemanagerSession) AggOracleMembers(arg0 *big.Int) (common.Address, error) {
	return _Aggoraclemanager.Contract.AggOracleMembers(&_Aggoraclemanager.CallOpts, arg0)
}

// AggOracleMembers is a free data retrieval call binding the contract method 0x1ffbaa74.
//
// Solidity: function aggOracleMembers(uint256 ) view returns(address)
func (_Aggoraclemanager *AggoraclemanagerCallerSession) AggOracleMembers(arg0 *big.Int) (common.Address, error) {
	return _Aggoraclemanager.Contract.AggOracleMembers(&_Aggoraclemanager.CallOpts, arg0)
}

// GetAggOracleMemberIndex is a free data retrieval call binding the contract method 0x72312af8.
//
// Solidity: function getAggOracleMemberIndex(address oracleMember) view returns(uint256)
func (_Aggoraclemanager *AggoraclemanagerCaller) GetAggOracleMemberIndex(opts *bind.CallOpts, oracleMember common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aggoraclemanager.contract.Call(opts, &out, "getAggOracleMemberIndex", oracleMember)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAggOracleMemberIndex is a free data retrieval call binding the contract method 0x72312af8.
//
// Solidity: function getAggOracleMemberIndex(address oracleMember) view returns(uint256)
func (_Aggoraclemanager *AggoraclemanagerSession) GetAggOracleMemberIndex(oracleMember common.Address) (*big.Int, error) {
	return _Aggoraclemanager.Contract.GetAggOracleMemberIndex(&_Aggoraclemanager.CallOpts, oracleMember)
}

// GetAggOracleMemberIndex is a free data retrieval call binding the contract method 0x72312af8.
//
// Solidity: function getAggOracleMemberIndex(address oracleMember) view returns(uint256)
func (_Aggoraclemanager *AggoraclemanagerCallerSession) GetAggOracleMemberIndex(oracleMember common.Address) (*big.Int, error) {
	return _Aggoraclemanager.Contract.GetAggOracleMemberIndex(&_Aggoraclemanager.CallOpts, oracleMember)
}

// GetAggOracleMembersCount is a free data retrieval call binding the contract method 0x017e37d6.
//
// Solidity: function getAggOracleMembersCount() view returns(uint256)
func (_Aggoraclemanager *AggoraclemanagerCaller) GetAggOracleMembersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggoraclemanager.contract.Call(opts, &out, "getAggOracleMembersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAggOracleMembersCount is a free data retrieval call binding the contract method 0x017e37d6.
//
// Solidity: function getAggOracleMembersCount() view returns(uint256)
func (_Aggoraclemanager *AggoraclemanagerSession) GetAggOracleMembersCount() (*big.Int, error) {
	return _Aggoraclemanager.Contract.GetAggOracleMembersCount(&_Aggoraclemanager.CallOpts)
}

// GetAggOracleMembersCount is a free data retrieval call binding the contract method 0x017e37d6.
//
// Solidity: function getAggOracleMembersCount() view returns(uint256)
func (_Aggoraclemanager *AggoraclemanagerCallerSession) GetAggOracleMembersCount() (*big.Int, error) {
	return _Aggoraclemanager.Contract.GetAggOracleMembersCount(&_Aggoraclemanager.CallOpts)
}

// GetAllAggOracleMembers is a free data retrieval call binding the contract method 0x703d8d42.
//
// Solidity: function getAllAggOracleMembers() view returns(address[])
func (_Aggoraclemanager *AggoraclemanagerCaller) GetAllAggOracleMembers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Aggoraclemanager.contract.Call(opts, &out, "getAllAggOracleMembers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllAggOracleMembers is a free data retrieval call binding the contract method 0x703d8d42.
//
// Solidity: function getAllAggOracleMembers() view returns(address[])
func (_Aggoraclemanager *AggoraclemanagerSession) GetAllAggOracleMembers() ([]common.Address, error) {
	return _Aggoraclemanager.Contract.GetAllAggOracleMembers(&_Aggoraclemanager.CallOpts)
}

// GetAllAggOracleMembers is a free data retrieval call binding the contract method 0x703d8d42.
//
// Solidity: function getAllAggOracleMembers() view returns(address[])
func (_Aggoraclemanager *AggoraclemanagerCallerSession) GetAllAggOracleMembers() ([]common.Address, error) {
	return _Aggoraclemanager.Contract.GetAllAggOracleMembers(&_Aggoraclemanager.CallOpts)
}

// GlobalExitRootManagerL2Sovereign is a free data retrieval call binding the contract method 0xa7a90e2a.
//
// Solidity: function globalExitRootManagerL2Sovereign() view returns(address)
func (_Aggoraclemanager *AggoraclemanagerCaller) GlobalExitRootManagerL2Sovereign(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggoraclemanager.contract.Call(opts, &out, "globalExitRootManagerL2Sovereign")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManagerL2Sovereign is a free data retrieval call binding the contract method 0xa7a90e2a.
//
// Solidity: function globalExitRootManagerL2Sovereign() view returns(address)
func (_Aggoraclemanager *AggoraclemanagerSession) GlobalExitRootManagerL2Sovereign() (common.Address, error) {
	return _Aggoraclemanager.Contract.GlobalExitRootManagerL2Sovereign(&_Aggoraclemanager.CallOpts)
}

// GlobalExitRootManagerL2Sovereign is a free data retrieval call binding the contract method 0xa7a90e2a.
//
// Solidity: function globalExitRootManagerL2Sovereign() view returns(address)
func (_Aggoraclemanager *AggoraclemanagerCallerSession) GlobalExitRootManagerL2Sovereign() (common.Address, error) {
	return _Aggoraclemanager.Contract.GlobalExitRootManagerL2Sovereign(&_Aggoraclemanager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Aggoraclemanager *AggoraclemanagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggoraclemanager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Aggoraclemanager *AggoraclemanagerSession) Owner() (common.Address, error) {
	return _Aggoraclemanager.Contract.Owner(&_Aggoraclemanager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Aggoraclemanager *AggoraclemanagerCallerSession) Owner() (common.Address, error) {
	return _Aggoraclemanager.Contract.Owner(&_Aggoraclemanager.CallOpts)
}

// ProposedGERToReport is a free data retrieval call binding the contract method 0x51146e10.
//
// Solidity: function proposedGERToReport(bytes32 ) view returns(uint64 votes, uint64 timestamp)
func (_Aggoraclemanager *AggoraclemanagerCaller) ProposedGERToReport(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Votes     uint64
	Timestamp uint64
}, error) {
	var out []interface{}
	err := _Aggoraclemanager.contract.Call(opts, &out, "proposedGERToReport", arg0)

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
func (_Aggoraclemanager *AggoraclemanagerSession) ProposedGERToReport(arg0 [32]byte) (struct {
	Votes     uint64
	Timestamp uint64
}, error) {
	return _Aggoraclemanager.Contract.ProposedGERToReport(&_Aggoraclemanager.CallOpts, arg0)
}

// ProposedGERToReport is a free data retrieval call binding the contract method 0x51146e10.
//
// Solidity: function proposedGERToReport(bytes32 ) view returns(uint64 votes, uint64 timestamp)
func (_Aggoraclemanager *AggoraclemanagerCallerSession) ProposedGERToReport(arg0 [32]byte) (struct {
	Votes     uint64
	Timestamp uint64
}, error) {
	return _Aggoraclemanager.Contract.ProposedGERToReport(&_Aggoraclemanager.CallOpts, arg0)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint64)
func (_Aggoraclemanager *AggoraclemanagerCaller) Quorum(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Aggoraclemanager.contract.Call(opts, &out, "quorum")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint64)
func (_Aggoraclemanager *AggoraclemanagerSession) Quorum() (uint64, error) {
	return _Aggoraclemanager.Contract.Quorum(&_Aggoraclemanager.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint64)
func (_Aggoraclemanager *AggoraclemanagerCallerSession) Quorum() (uint64, error) {
	return _Aggoraclemanager.Contract.Quorum(&_Aggoraclemanager.CallOpts)
}

// AcceptGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0xc053902a.
//
// Solidity: function acceptGlobalExitRootUpdater() returns()
func (_Aggoraclemanager *AggoraclemanagerTransactor) AcceptGlobalExitRootUpdater(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggoraclemanager.contract.Transact(opts, "acceptGlobalExitRootUpdater")
}

// AcceptGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0xc053902a.
//
// Solidity: function acceptGlobalExitRootUpdater() returns()
func (_Aggoraclemanager *AggoraclemanagerSession) AcceptGlobalExitRootUpdater() (*types.Transaction, error) {
	return _Aggoraclemanager.Contract.AcceptGlobalExitRootUpdater(&_Aggoraclemanager.TransactOpts)
}

// AcceptGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0xc053902a.
//
// Solidity: function acceptGlobalExitRootUpdater() returns()
func (_Aggoraclemanager *AggoraclemanagerTransactorSession) AcceptGlobalExitRootUpdater() (*types.Transaction, error) {
	return _Aggoraclemanager.Contract.AcceptGlobalExitRootUpdater(&_Aggoraclemanager.TransactOpts)
}

// AddOracleMember is a paid mutator transaction binding the contract method 0xb164e437.
//
// Solidity: function addOracleMember(address newOracleMember) returns()
func (_Aggoraclemanager *AggoraclemanagerTransactor) AddOracleMember(opts *bind.TransactOpts, newOracleMember common.Address) (*types.Transaction, error) {
	return _Aggoraclemanager.contract.Transact(opts, "addOracleMember", newOracleMember)
}

// AddOracleMember is a paid mutator transaction binding the contract method 0xb164e437.
//
// Solidity: function addOracleMember(address newOracleMember) returns()
func (_Aggoraclemanager *AggoraclemanagerSession) AddOracleMember(newOracleMember common.Address) (*types.Transaction, error) {
	return _Aggoraclemanager.Contract.AddOracleMember(&_Aggoraclemanager.TransactOpts, newOracleMember)
}

// AddOracleMember is a paid mutator transaction binding the contract method 0xb164e437.
//
// Solidity: function addOracleMember(address newOracleMember) returns()
func (_Aggoraclemanager *AggoraclemanagerTransactorSession) AddOracleMember(newOracleMember common.Address) (*types.Transaction, error) {
	return _Aggoraclemanager.Contract.AddOracleMember(&_Aggoraclemanager.TransactOpts, newOracleMember)
}

// Initialize is a paid mutator transaction binding the contract method 0x145d1ed6.
//
// Solidity: function initialize(address _owner, address[] _aggOracleMembers, uint64 _quorum) returns()
func (_Aggoraclemanager *AggoraclemanagerTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _aggOracleMembers []common.Address, _quorum uint64) (*types.Transaction, error) {
	return _Aggoraclemanager.contract.Transact(opts, "initialize", _owner, _aggOracleMembers, _quorum)
}

// Initialize is a paid mutator transaction binding the contract method 0x145d1ed6.
//
// Solidity: function initialize(address _owner, address[] _aggOracleMembers, uint64 _quorum) returns()
func (_Aggoraclemanager *AggoraclemanagerSession) Initialize(_owner common.Address, _aggOracleMembers []common.Address, _quorum uint64) (*types.Transaction, error) {
	return _Aggoraclemanager.Contract.Initialize(&_Aggoraclemanager.TransactOpts, _owner, _aggOracleMembers, _quorum)
}

// Initialize is a paid mutator transaction binding the contract method 0x145d1ed6.
//
// Solidity: function initialize(address _owner, address[] _aggOracleMembers, uint64 _quorum) returns()
func (_Aggoraclemanager *AggoraclemanagerTransactorSession) Initialize(_owner common.Address, _aggOracleMembers []common.Address, _quorum uint64) (*types.Transaction, error) {
	return _Aggoraclemanager.Contract.Initialize(&_Aggoraclemanager.TransactOpts, _owner, _aggOracleMembers, _quorum)
}

// ProposeGlobalExitRoot is a paid mutator transaction binding the contract method 0x542b56b3.
//
// Solidity: function proposeGlobalExitRoot(bytes32 proposedGlobalExitRoot) returns()
func (_Aggoraclemanager *AggoraclemanagerTransactor) ProposeGlobalExitRoot(opts *bind.TransactOpts, proposedGlobalExitRoot [32]byte) (*types.Transaction, error) {
	return _Aggoraclemanager.contract.Transact(opts, "proposeGlobalExitRoot", proposedGlobalExitRoot)
}

// ProposeGlobalExitRoot is a paid mutator transaction binding the contract method 0x542b56b3.
//
// Solidity: function proposeGlobalExitRoot(bytes32 proposedGlobalExitRoot) returns()
func (_Aggoraclemanager *AggoraclemanagerSession) ProposeGlobalExitRoot(proposedGlobalExitRoot [32]byte) (*types.Transaction, error) {
	return _Aggoraclemanager.Contract.ProposeGlobalExitRoot(&_Aggoraclemanager.TransactOpts, proposedGlobalExitRoot)
}

// ProposeGlobalExitRoot is a paid mutator transaction binding the contract method 0x542b56b3.
//
// Solidity: function proposeGlobalExitRoot(bytes32 proposedGlobalExitRoot) returns()
func (_Aggoraclemanager *AggoraclemanagerTransactorSession) ProposeGlobalExitRoot(proposedGlobalExitRoot [32]byte) (*types.Transaction, error) {
	return _Aggoraclemanager.Contract.ProposeGlobalExitRoot(&_Aggoraclemanager.TransactOpts, proposedGlobalExitRoot)
}

// RemoveOracleMember is a paid mutator transaction binding the contract method 0x53985e5a.
//
// Solidity: function removeOracleMember(address oracleMemberAddress, uint256 oracleMemberIndex) returns()
func (_Aggoraclemanager *AggoraclemanagerTransactor) RemoveOracleMember(opts *bind.TransactOpts, oracleMemberAddress common.Address, oracleMemberIndex *big.Int) (*types.Transaction, error) {
	return _Aggoraclemanager.contract.Transact(opts, "removeOracleMember", oracleMemberAddress, oracleMemberIndex)
}

// RemoveOracleMember is a paid mutator transaction binding the contract method 0x53985e5a.
//
// Solidity: function removeOracleMember(address oracleMemberAddress, uint256 oracleMemberIndex) returns()
func (_Aggoraclemanager *AggoraclemanagerSession) RemoveOracleMember(oracleMemberAddress common.Address, oracleMemberIndex *big.Int) (*types.Transaction, error) {
	return _Aggoraclemanager.Contract.RemoveOracleMember(&_Aggoraclemanager.TransactOpts, oracleMemberAddress, oracleMemberIndex)
}

// RemoveOracleMember is a paid mutator transaction binding the contract method 0x53985e5a.
//
// Solidity: function removeOracleMember(address oracleMemberAddress, uint256 oracleMemberIndex) returns()
func (_Aggoraclemanager *AggoraclemanagerTransactorSession) RemoveOracleMember(oracleMemberAddress common.Address, oracleMemberIndex *big.Int) (*types.Transaction, error) {
	return _Aggoraclemanager.Contract.RemoveOracleMember(&_Aggoraclemanager.TransactOpts, oracleMemberAddress, oracleMemberIndex)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aggoraclemanager *AggoraclemanagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggoraclemanager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aggoraclemanager *AggoraclemanagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Aggoraclemanager.Contract.RenounceOwnership(&_Aggoraclemanager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aggoraclemanager *AggoraclemanagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Aggoraclemanager.Contract.RenounceOwnership(&_Aggoraclemanager.TransactOpts)
}

// TransferGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x0e1bbf9f.
//
// Solidity: function transferGlobalExitRootUpdater(address _newGlobalExitRootUpdater) returns()
func (_Aggoraclemanager *AggoraclemanagerTransactor) TransferGlobalExitRootUpdater(opts *bind.TransactOpts, _newGlobalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Aggoraclemanager.contract.Transact(opts, "transferGlobalExitRootUpdater", _newGlobalExitRootUpdater)
}

// TransferGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x0e1bbf9f.
//
// Solidity: function transferGlobalExitRootUpdater(address _newGlobalExitRootUpdater) returns()
func (_Aggoraclemanager *AggoraclemanagerSession) TransferGlobalExitRootUpdater(_newGlobalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Aggoraclemanager.Contract.TransferGlobalExitRootUpdater(&_Aggoraclemanager.TransactOpts, _newGlobalExitRootUpdater)
}

// TransferGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x0e1bbf9f.
//
// Solidity: function transferGlobalExitRootUpdater(address _newGlobalExitRootUpdater) returns()
func (_Aggoraclemanager *AggoraclemanagerTransactorSession) TransferGlobalExitRootUpdater(_newGlobalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Aggoraclemanager.Contract.TransferGlobalExitRootUpdater(&_Aggoraclemanager.TransactOpts, _newGlobalExitRootUpdater)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Aggoraclemanager *AggoraclemanagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Aggoraclemanager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Aggoraclemanager *AggoraclemanagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Aggoraclemanager.Contract.TransferOwnership(&_Aggoraclemanager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Aggoraclemanager *AggoraclemanagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Aggoraclemanager.Contract.TransferOwnership(&_Aggoraclemanager.TransactOpts, newOwner)
}

// UpdateQuorum is a paid mutator transaction binding the contract method 0x29218b61.
//
// Solidity: function updateQuorum(uint64 newQuorum) returns()
func (_Aggoraclemanager *AggoraclemanagerTransactor) UpdateQuorum(opts *bind.TransactOpts, newQuorum uint64) (*types.Transaction, error) {
	return _Aggoraclemanager.contract.Transact(opts, "updateQuorum", newQuorum)
}

// UpdateQuorum is a paid mutator transaction binding the contract method 0x29218b61.
//
// Solidity: function updateQuorum(uint64 newQuorum) returns()
func (_Aggoraclemanager *AggoraclemanagerSession) UpdateQuorum(newQuorum uint64) (*types.Transaction, error) {
	return _Aggoraclemanager.Contract.UpdateQuorum(&_Aggoraclemanager.TransactOpts, newQuorum)
}

// UpdateQuorum is a paid mutator transaction binding the contract method 0x29218b61.
//
// Solidity: function updateQuorum(uint64 newQuorum) returns()
func (_Aggoraclemanager *AggoraclemanagerTransactorSession) UpdateQuorum(newQuorum uint64) (*types.Transaction, error) {
	return _Aggoraclemanager.Contract.UpdateQuorum(&_Aggoraclemanager.TransactOpts, newQuorum)
}

// AggoraclemanagerAddAggOracleMemberIterator is returned from FilterAddAggOracleMember and is used to iterate over the raw logs and unpacked data for AddAggOracleMember events raised by the Aggoraclemanager contract.
type AggoraclemanagerAddAggOracleMemberIterator struct {
	Event *AggoraclemanagerAddAggOracleMember // Event containing the contract specifics and raw log

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
func (it *AggoraclemanagerAddAggOracleMemberIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggoraclemanagerAddAggOracleMember)
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
		it.Event = new(AggoraclemanagerAddAggOracleMember)
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
func (it *AggoraclemanagerAddAggOracleMemberIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggoraclemanagerAddAggOracleMemberIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggoraclemanagerAddAggOracleMember represents a AddAggOracleMember event raised by the Aggoraclemanager contract.
type AggoraclemanagerAddAggOracleMember struct {
	NewOracleMember common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddAggOracleMember is a free log retrieval operation binding the contract event 0x59f8c2dd0bdc5787ef574566b9edcd7324a62b79ac952930f02679a1b25bbf6f.
//
// Solidity: event AddAggOracleMember(address newOracleMember)
func (_Aggoraclemanager *AggoraclemanagerFilterer) FilterAddAggOracleMember(opts *bind.FilterOpts) (*AggoraclemanagerAddAggOracleMemberIterator, error) {

	logs, sub, err := _Aggoraclemanager.contract.FilterLogs(opts, "AddAggOracleMember")
	if err != nil {
		return nil, err
	}
	return &AggoraclemanagerAddAggOracleMemberIterator{contract: _Aggoraclemanager.contract, event: "AddAggOracleMember", logs: logs, sub: sub}, nil
}

// WatchAddAggOracleMember is a free log subscription operation binding the contract event 0x59f8c2dd0bdc5787ef574566b9edcd7324a62b79ac952930f02679a1b25bbf6f.
//
// Solidity: event AddAggOracleMember(address newOracleMember)
func (_Aggoraclemanager *AggoraclemanagerFilterer) WatchAddAggOracleMember(opts *bind.WatchOpts, sink chan<- *AggoraclemanagerAddAggOracleMember) (event.Subscription, error) {

	logs, sub, err := _Aggoraclemanager.contract.WatchLogs(opts, "AddAggOracleMember")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggoraclemanagerAddAggOracleMember)
				if err := _Aggoraclemanager.contract.UnpackLog(event, "AddAggOracleMember", log); err != nil {
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
func (_Aggoraclemanager *AggoraclemanagerFilterer) ParseAddAggOracleMember(log types.Log) (*AggoraclemanagerAddAggOracleMember, error) {
	event := new(AggoraclemanagerAddAggOracleMember)
	if err := _Aggoraclemanager.contract.UnpackLog(event, "AddAggOracleMember", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggoraclemanagerConsolidatedGlobalExitRootIterator is returned from FilterConsolidatedGlobalExitRoot and is used to iterate over the raw logs and unpacked data for ConsolidatedGlobalExitRoot events raised by the Aggoraclemanager contract.
type AggoraclemanagerConsolidatedGlobalExitRootIterator struct {
	Event *AggoraclemanagerConsolidatedGlobalExitRoot // Event containing the contract specifics and raw log

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
func (it *AggoraclemanagerConsolidatedGlobalExitRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggoraclemanagerConsolidatedGlobalExitRoot)
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
		it.Event = new(AggoraclemanagerConsolidatedGlobalExitRoot)
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
func (it *AggoraclemanagerConsolidatedGlobalExitRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggoraclemanagerConsolidatedGlobalExitRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggoraclemanagerConsolidatedGlobalExitRoot represents a ConsolidatedGlobalExitRoot event raised by the Aggoraclemanager contract.
type AggoraclemanagerConsolidatedGlobalExitRoot struct {
	ConsolidatedGlobalExitRoot [32]byte
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterConsolidatedGlobalExitRoot is a free log retrieval operation binding the contract event 0x1676a1944900a9899dac554c2da10ba62637fe8d08375ffc3fb8494c92918d98.
//
// Solidity: event ConsolidatedGlobalExitRoot(bytes32 consolidatedGlobalExitRoot)
func (_Aggoraclemanager *AggoraclemanagerFilterer) FilterConsolidatedGlobalExitRoot(opts *bind.FilterOpts) (*AggoraclemanagerConsolidatedGlobalExitRootIterator, error) {

	logs, sub, err := _Aggoraclemanager.contract.FilterLogs(opts, "ConsolidatedGlobalExitRoot")
	if err != nil {
		return nil, err
	}
	return &AggoraclemanagerConsolidatedGlobalExitRootIterator{contract: _Aggoraclemanager.contract, event: "ConsolidatedGlobalExitRoot", logs: logs, sub: sub}, nil
}

// WatchConsolidatedGlobalExitRoot is a free log subscription operation binding the contract event 0x1676a1944900a9899dac554c2da10ba62637fe8d08375ffc3fb8494c92918d98.
//
// Solidity: event ConsolidatedGlobalExitRoot(bytes32 consolidatedGlobalExitRoot)
func (_Aggoraclemanager *AggoraclemanagerFilterer) WatchConsolidatedGlobalExitRoot(opts *bind.WatchOpts, sink chan<- *AggoraclemanagerConsolidatedGlobalExitRoot) (event.Subscription, error) {

	logs, sub, err := _Aggoraclemanager.contract.WatchLogs(opts, "ConsolidatedGlobalExitRoot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggoraclemanagerConsolidatedGlobalExitRoot)
				if err := _Aggoraclemanager.contract.UnpackLog(event, "ConsolidatedGlobalExitRoot", log); err != nil {
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
func (_Aggoraclemanager *AggoraclemanagerFilterer) ParseConsolidatedGlobalExitRoot(log types.Log) (*AggoraclemanagerConsolidatedGlobalExitRoot, error) {
	event := new(AggoraclemanagerConsolidatedGlobalExitRoot)
	if err := _Aggoraclemanager.contract.UnpackLog(event, "ConsolidatedGlobalExitRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggoraclemanagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Aggoraclemanager contract.
type AggoraclemanagerInitializedIterator struct {
	Event *AggoraclemanagerInitialized // Event containing the contract specifics and raw log

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
func (it *AggoraclemanagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggoraclemanagerInitialized)
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
		it.Event = new(AggoraclemanagerInitialized)
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
func (it *AggoraclemanagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggoraclemanagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggoraclemanagerInitialized represents a Initialized event raised by the Aggoraclemanager contract.
type AggoraclemanagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Aggoraclemanager *AggoraclemanagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*AggoraclemanagerInitializedIterator, error) {

	logs, sub, err := _Aggoraclemanager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AggoraclemanagerInitializedIterator{contract: _Aggoraclemanager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Aggoraclemanager *AggoraclemanagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AggoraclemanagerInitialized) (event.Subscription, error) {

	logs, sub, err := _Aggoraclemanager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggoraclemanagerInitialized)
				if err := _Aggoraclemanager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Aggoraclemanager *AggoraclemanagerFilterer) ParseInitialized(log types.Log) (*AggoraclemanagerInitialized, error) {
	event := new(AggoraclemanagerInitialized)
	if err := _Aggoraclemanager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggoraclemanagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Aggoraclemanager contract.
type AggoraclemanagerOwnershipTransferredIterator struct {
	Event *AggoraclemanagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AggoraclemanagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggoraclemanagerOwnershipTransferred)
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
		it.Event = new(AggoraclemanagerOwnershipTransferred)
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
func (it *AggoraclemanagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggoraclemanagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggoraclemanagerOwnershipTransferred represents a OwnershipTransferred event raised by the Aggoraclemanager contract.
type AggoraclemanagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Aggoraclemanager *AggoraclemanagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AggoraclemanagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Aggoraclemanager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AggoraclemanagerOwnershipTransferredIterator{contract: _Aggoraclemanager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Aggoraclemanager *AggoraclemanagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AggoraclemanagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Aggoraclemanager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggoraclemanagerOwnershipTransferred)
				if err := _Aggoraclemanager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Aggoraclemanager *AggoraclemanagerFilterer) ParseOwnershipTransferred(log types.Log) (*AggoraclemanagerOwnershipTransferred, error) {
	event := new(AggoraclemanagerOwnershipTransferred)
	if err := _Aggoraclemanager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggoraclemanagerProposedGlobalExitRootIterator is returned from FilterProposedGlobalExitRoot and is used to iterate over the raw logs and unpacked data for ProposedGlobalExitRoot events raised by the Aggoraclemanager contract.
type AggoraclemanagerProposedGlobalExitRootIterator struct {
	Event *AggoraclemanagerProposedGlobalExitRoot // Event containing the contract specifics and raw log

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
func (it *AggoraclemanagerProposedGlobalExitRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggoraclemanagerProposedGlobalExitRoot)
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
		it.Event = new(AggoraclemanagerProposedGlobalExitRoot)
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
func (it *AggoraclemanagerProposedGlobalExitRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggoraclemanagerProposedGlobalExitRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggoraclemanagerProposedGlobalExitRoot represents a ProposedGlobalExitRoot event raised by the Aggoraclemanager contract.
type AggoraclemanagerProposedGlobalExitRoot struct {
	ProposedGlobalExitRoot [32]byte
	Proposer               common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterProposedGlobalExitRoot is a free log retrieval operation binding the contract event 0x9f298266555f7b22d45b8ea34a9c99cba101d1ccbc3ff6dc4c392b7c915dc3b5.
//
// Solidity: event ProposedGlobalExitRoot(bytes32 proposedGlobalExitRoot, address proposer)
func (_Aggoraclemanager *AggoraclemanagerFilterer) FilterProposedGlobalExitRoot(opts *bind.FilterOpts) (*AggoraclemanagerProposedGlobalExitRootIterator, error) {

	logs, sub, err := _Aggoraclemanager.contract.FilterLogs(opts, "ProposedGlobalExitRoot")
	if err != nil {
		return nil, err
	}
	return &AggoraclemanagerProposedGlobalExitRootIterator{contract: _Aggoraclemanager.contract, event: "ProposedGlobalExitRoot", logs: logs, sub: sub}, nil
}

// WatchProposedGlobalExitRoot is a free log subscription operation binding the contract event 0x9f298266555f7b22d45b8ea34a9c99cba101d1ccbc3ff6dc4c392b7c915dc3b5.
//
// Solidity: event ProposedGlobalExitRoot(bytes32 proposedGlobalExitRoot, address proposer)
func (_Aggoraclemanager *AggoraclemanagerFilterer) WatchProposedGlobalExitRoot(opts *bind.WatchOpts, sink chan<- *AggoraclemanagerProposedGlobalExitRoot) (event.Subscription, error) {

	logs, sub, err := _Aggoraclemanager.contract.WatchLogs(opts, "ProposedGlobalExitRoot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggoraclemanagerProposedGlobalExitRoot)
				if err := _Aggoraclemanager.contract.UnpackLog(event, "ProposedGlobalExitRoot", log); err != nil {
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
func (_Aggoraclemanager *AggoraclemanagerFilterer) ParseProposedGlobalExitRoot(log types.Log) (*AggoraclemanagerProposedGlobalExitRoot, error) {
	event := new(AggoraclemanagerProposedGlobalExitRoot)
	if err := _Aggoraclemanager.contract.UnpackLog(event, "ProposedGlobalExitRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggoraclemanagerRemoveAggOracleMemberIterator is returned from FilterRemoveAggOracleMember and is used to iterate over the raw logs and unpacked data for RemoveAggOracleMember events raised by the Aggoraclemanager contract.
type AggoraclemanagerRemoveAggOracleMemberIterator struct {
	Event *AggoraclemanagerRemoveAggOracleMember // Event containing the contract specifics and raw log

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
func (it *AggoraclemanagerRemoveAggOracleMemberIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggoraclemanagerRemoveAggOracleMember)
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
		it.Event = new(AggoraclemanagerRemoveAggOracleMember)
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
func (it *AggoraclemanagerRemoveAggOracleMemberIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggoraclemanagerRemoveAggOracleMemberIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggoraclemanagerRemoveAggOracleMember represents a RemoveAggOracleMember event raised by the Aggoraclemanager contract.
type AggoraclemanagerRemoveAggOracleMember struct {
	OracleMemberRemoved common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRemoveAggOracleMember is a free log retrieval operation binding the contract event 0x396186e0feff55f3e52260aee4ec024318786c7f5b8b01639b1d42c24b59eafc.
//
// Solidity: event RemoveAggOracleMember(address oracleMemberRemoved)
func (_Aggoraclemanager *AggoraclemanagerFilterer) FilterRemoveAggOracleMember(opts *bind.FilterOpts) (*AggoraclemanagerRemoveAggOracleMemberIterator, error) {

	logs, sub, err := _Aggoraclemanager.contract.FilterLogs(opts, "RemoveAggOracleMember")
	if err != nil {
		return nil, err
	}
	return &AggoraclemanagerRemoveAggOracleMemberIterator{contract: _Aggoraclemanager.contract, event: "RemoveAggOracleMember", logs: logs, sub: sub}, nil
}

// WatchRemoveAggOracleMember is a free log subscription operation binding the contract event 0x396186e0feff55f3e52260aee4ec024318786c7f5b8b01639b1d42c24b59eafc.
//
// Solidity: event RemoveAggOracleMember(address oracleMemberRemoved)
func (_Aggoraclemanager *AggoraclemanagerFilterer) WatchRemoveAggOracleMember(opts *bind.WatchOpts, sink chan<- *AggoraclemanagerRemoveAggOracleMember) (event.Subscription, error) {

	logs, sub, err := _Aggoraclemanager.contract.WatchLogs(opts, "RemoveAggOracleMember")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggoraclemanagerRemoveAggOracleMember)
				if err := _Aggoraclemanager.contract.UnpackLog(event, "RemoveAggOracleMember", log); err != nil {
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
func (_Aggoraclemanager *AggoraclemanagerFilterer) ParseRemoveAggOracleMember(log types.Log) (*AggoraclemanagerRemoveAggOracleMember, error) {
	event := new(AggoraclemanagerRemoveAggOracleMember)
	if err := _Aggoraclemanager.contract.UnpackLog(event, "RemoveAggOracleMember", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggoraclemanagerUpdateQuorumIterator is returned from FilterUpdateQuorum and is used to iterate over the raw logs and unpacked data for UpdateQuorum events raised by the Aggoraclemanager contract.
type AggoraclemanagerUpdateQuorumIterator struct {
	Event *AggoraclemanagerUpdateQuorum // Event containing the contract specifics and raw log

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
func (it *AggoraclemanagerUpdateQuorumIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggoraclemanagerUpdateQuorum)
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
		it.Event = new(AggoraclemanagerUpdateQuorum)
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
func (it *AggoraclemanagerUpdateQuorumIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggoraclemanagerUpdateQuorumIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggoraclemanagerUpdateQuorum represents a UpdateQuorum event raised by the Aggoraclemanager contract.
type AggoraclemanagerUpdateQuorum struct {
	NewQuorum uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateQuorum is a free log retrieval operation binding the contract event 0xb600f3cf7f38a4b49bb0c75f722ef69f7e3e39ef3bb4aa8207fd86e724a23249.
//
// Solidity: event UpdateQuorum(uint64 newQuorum)
func (_Aggoraclemanager *AggoraclemanagerFilterer) FilterUpdateQuorum(opts *bind.FilterOpts) (*AggoraclemanagerUpdateQuorumIterator, error) {

	logs, sub, err := _Aggoraclemanager.contract.FilterLogs(opts, "UpdateQuorum")
	if err != nil {
		return nil, err
	}
	return &AggoraclemanagerUpdateQuorumIterator{contract: _Aggoraclemanager.contract, event: "UpdateQuorum", logs: logs, sub: sub}, nil
}

// WatchUpdateQuorum is a free log subscription operation binding the contract event 0xb600f3cf7f38a4b49bb0c75f722ef69f7e3e39ef3bb4aa8207fd86e724a23249.
//
// Solidity: event UpdateQuorum(uint64 newQuorum)
func (_Aggoraclemanager *AggoraclemanagerFilterer) WatchUpdateQuorum(opts *bind.WatchOpts, sink chan<- *AggoraclemanagerUpdateQuorum) (event.Subscription, error) {

	logs, sub, err := _Aggoraclemanager.contract.WatchLogs(opts, "UpdateQuorum")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggoraclemanagerUpdateQuorum)
				if err := _Aggoraclemanager.contract.UnpackLog(event, "UpdateQuorum", log); err != nil {
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
func (_Aggoraclemanager *AggoraclemanagerFilterer) ParseUpdateQuorum(log types.Log) (*AggoraclemanagerUpdateQuorum, error) {
	event := new(AggoraclemanagerUpdateQuorum)
	if err := _Aggoraclemanager.contract.UnpackLog(event, "UpdateQuorum", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

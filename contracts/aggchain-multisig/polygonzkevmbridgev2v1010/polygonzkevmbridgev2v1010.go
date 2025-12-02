// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonzkevmbridgev2v1010

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

// Polygonzkevmbridgev2v1010MetaData contains all meta data concerning the Polygonzkevmbridgev2v1010 contract.
var Polygonzkevmbridgev2v1010MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountDoesNotMatchMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeAddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DestinationNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EtherTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedProxyDeployment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGlobalIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxyAdmin\",\"type\":\"address\"}],\"name\":\"InvalidProxyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSmtProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxyAdmin\",\"type\":\"address\"}],\"name\":\"InvalidZeroProxyAdminOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MsgValueNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeTokenIsEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewDepositCountExceedsMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoValueInMessagesOnGasTokenNetworks\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonZeroValueForUnusedFrontier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingProxiedTokensManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProxiedTokensManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubtreeFrontierMismatch\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldProxiedTokensManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newProxiedTokensManager\",\"type\":\"address\"}],\"name\":\"AcceptProxiedTokensManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"depositCount\",\"type\":\"uint32\"}],\"name\":\"BridgeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wrappedTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"NewWrappedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentProxiedTokensManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newProxiedTokensManager\",\"type\":\"address\"}],\"name\":\"TransferProxiedTokensManagerRole\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BRIDGE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_BYTECODE_TRANSPARENT_PROXY\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETHToken\",\"outputs\":[{\"internalType\":\"contractITokenWrappedBridgeUpgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptProxiedTokensManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"permitData\",\"type\":\"bytes\"}],\"name\":\"bridgeAsset\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountWETH\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessageWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimedBitMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"name\":\"computeTokenProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProxiedTokensManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenWrappedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWrappedTokenBridgeImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIBaseLegacyAgglayerGER\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"contractIBaseLegacyAgglayerGER\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_polygonRollupManager\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"leafIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"sourceBridgeNetwork\",\"type\":\"uint32\"}],\"name\":\"isClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdatedDepositCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingProxiedTokensManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"polygonRollupManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiedTokensManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenInfoToWrappedToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newProxiedTokensManager\",\"type\":\"address\"}],\"name\":\"transferProxiedTokensManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedTokenBytecodeStorer\",\"outputs\":[{\"internalType\":\"contractIBytecodeStorer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wrappedTokenToTokenInfo\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561000f575f5ffd5b5060405161001c9061013e565b604051809103905ff080158015610035573d5f5f3e3d5ffd5b506001600160a01b031660805260405161004e9061014b565b604051809103905ff080158015610067573d5f5f3e3d5ffd5b506001600160a01b031660a05261007c610081565b610158565b5f54610100900460ff16156100ec5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116101561013c575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b610fd7806150ae83390190565b611d368061608583390190565b60805160a051614f2e6101805f395f61042d01525f81816103d301526110040152614f2e5ff3fe60806040526004361061025d575f3560e01c806381b1c1741161014b578063cc461632116100c6578063ece93c6f1161007c578063f214e16111610062578063f214e161146107df578063f5efcd79146107fe578063f811bff71461081d575f5ffd5b8063ece93c6f14610788578063ee25560b146107b4575f5ffd5b8063cd586579116100ac578063cd5865791461072c578063d02103ca1461073f578063dbc1697614610774575f5ffd5b8063cc461632146106ee578063ccaa2d111461070d575f5ffd5b8063b8b284d01161011b578063be5831c711610101578063be5831c714610682578063c00f14ab146106bb578063c514f24e146106da575f5ffd5b8063b8b284d014610642578063bab161bf14610661575f5ffd5b806381b1c174146105a25780638bd309c3146105e35780638c668f1c146106025780638ed7e3f214610616575f5ffd5b80633b2fee9a116101db5780635ca1e165116101ab5780636e4ecfed116101915780636e4ecfed1461054e57806379e2cf971461057a5780638129fc1c1461058e575f5ffd5b80635ca1e165146104f257806365d6f65414610506575f5ffd5b80633b2fee9a1461041f5780633c351e10146104515780633cbc795b1461047d5780634b2f336d146104c6575f5ffd5b806327aef4e811610230578063318aee3d11610216578063318aee3d14610340578063381fef6d146103c257806338b8fbbb146103f5575f5ffd5b806327aef4e8146102fc5780632dfdf0b51461031d575f5ffd5b806315064c96146102615780632072f6c51461028f57806322e95f2c146102a5578063240ff378146102e9575b5f5ffd5b34801561026c575f5ffd5b5060685461027a9060ff1681565b60405190151581526020015b60405180910390f35b34801561029a575f5ffd5b506102a361083c565b005b3480156102b0575f5ffd5b506102c46102bf36600461402e565b610897565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610286565b6102a36102f73660046140b5565b610939565b348015610307575f5ffd5b506103106109e8565b6040516102869190614176565b348015610328575f5ffd5b5061033260535481565b604051908152602001610286565b34801561034b575f5ffd5b5061039161035a366004614188565b606b6020525f908152604090205463ffffffff811690640100000000900473ffffffffffffffffffffffffffffffffffffffff1682565b6040805163ffffffff909316835273ffffffffffffffffffffffffffffffffffffffff909116602083015201610286565b3480156103cd575f5ffd5b506102c47f000000000000000000000000000000000000000000000000000000000000000081565b348015610400575f5ffd5b5060705473ffffffffffffffffffffffffffffffffffffffff166102c4565b34801561042a575f5ffd5b507f00000000000000000000000000000000000000000000000000000000000000006102c4565b34801561045c575f5ffd5b50606d546102c49073ffffffffffffffffffffffffffffffffffffffff1681565b348015610488575f5ffd5b50606d546104b19074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff9091168152602001610286565b3480156104d1575f5ffd5b50606f546102c49073ffffffffffffffffffffffffffffffffffffffff1681565b3480156104fd575f5ffd5b50610332610a74565b348015610511575f5ffd5b506103106040518060400160405280600681526020017f76312e302e30000000000000000000000000000000000000000000000000000081525081565b348015610559575f5ffd5b506070546102c49073ffffffffffffffffffffffffffffffffffffffff1681565b348015610585575f5ffd5b506102a3610af3565b348015610599575f5ffd5b506102a3610b2a565b3480156105ad575f5ffd5b506102c46105bc3660046141a3565b606a6020525f908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b3480156105ee575f5ffd5b506102a36105fd366004614188565b610d3c565b34801561060d575f5ffd5b506102a3610e15565b348015610621575f5ffd5b50606c546102c49073ffffffffffffffffffffffffffffffffffffffff1681565b34801561064d575f5ffd5b506102a361065c3660046141ba565b610ef1565b34801561066c575f5ffd5b506068546104b190610100900463ffffffff1681565b34801561068d575f5ffd5b506068546104b190790100000000000000000000000000000000000000000000000000900463ffffffff1681565b3480156106c6575f5ffd5b506103106106d5366004614188565b610fbb565b3480156106e5575f5ffd5b50610310611000565b3480156106f9575f5ffd5b5061027a610708366004614238565b6110b4565b348015610718575f5ffd5b506102a361072736600461427a565b61113c565b6102a361073a366004614359565b6116ee565b34801561074a575f5ffd5b506068546102c49065010000000000900473ffffffffffffffffffffffffffffffffffffffff1681565b34801561077f575f5ffd5b506102a3611b8c565b348015610793575f5ffd5b506071546102c49073ffffffffffffffffffffffffffffffffffffffff1681565b3480156107bf575f5ffd5b506103326107ce3660046141a3565b60696020525f908152604090205481565b3480156107ea575f5ffd5b506102c46107f936600461402e565b611be5565b348015610809575f5ffd5b506102a361081836600461427a565b611d46565b348015610828575f5ffd5b506102a36108373660046144aa565b6120ca565b606c5473ffffffffffffffffffffffffffffffffffffffff16331461088d576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61089561252a565b565b6040805160e084901b7fffffffff0000000000000000000000000000000000000000000000000000000016602080830191909152606084901b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602483015282516018818403018152603890920183528151918101919091205f908152606a909152205473ffffffffffffffffffffffffffffffffffffffff165b92915050565b60685460ff1615610976576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b341580159061099c5750606f5473ffffffffffffffffffffffffffffffffffffffff1615155b156109d3576040517f6f625c4000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109e18585348686866125bc565b5050505050565b606e80546109f59061457a565b80601f0160208091040260200160405190810160405280929190818152602001828054610a219061457a565b8015610a6c5780601f10610a4357610100808354040283529160200191610a6c565b820191905f5260205f20905b815481529060010190602001808311610a4f57829003601f168201915b505050505081565b6053545f90819081805b6020811015610aea578083901c600116600103610ac357610abc60338260208110610aab57610aab6145cb565b0154855f9182526020526040902090565b9350610ad3565b5f84815260208390526040902093505b5f8281526020839052604090209150600101610a7e565b50919392505050565b605354606854790100000000000000000000000000000000000000000000000000900463ffffffff161015610895576108956126b1565b5f5460ff16607180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000060ff938416021790555f5460029161010090910416158015610b9357505f5460ff8083169116105b610c24576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff80841691909117610100178255607154740100000000000000000000000000000000000000009004169003610cac576040517ff57ac68300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610cb4612782565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150607180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff169055565b60705473ffffffffffffffffffffffffffffffffffffffff163314610d8d576040517f0866750300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169182179092556070546040805191909316815260208101919091527f0a34baa3feb299aef9c05cb59c6e0c8e7c0bcc65cbf0a647e7a7c8a2411591e291015b60405180910390a150565b60715473ffffffffffffffffffffffffffffffffffffffff163314610e66576040517f2d67bc9c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607080546071805473ffffffffffffffffffffffffffffffffffffffff8082167fffffffffffffffffffffffff0000000000000000000000000000000000000000808616821790965594909116909155604080519190921680825260208201939093527fa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f29101610e0a565b60685460ff1615610f2e576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5473ffffffffffffffffffffffffffffffffffffffff16610f7d576040517fdde3cda700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f545f90610fa29073ffffffffffffffffffffffffffffffffffffffff1686612915565b9050610fb28787838787876125bc565b50505050505050565b6060610fc68261299f565b610fcf83612ab3565b610fd884612bb6565b604051602001610fea939291906145f8565b6040516020818303038152906040529050919050565b60607f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c514f24e6040518163ffffffff1660e01b81526004015f60405180830381865afa15801561106a573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526110af9190810190614666565b905090565b6068545f908190610100900463ffffffff161580156110d9575063ffffffff83166001145b156110eb575063ffffffff8316611113565b61110064010000000063ffffffff85166146d8565b6111109063ffffffff86166146ef565b90505b600881901c5f90815260696020526040902054600160ff9092169190911b908116149392505050565b60685460ff1615611179576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611181612ca5565b60685463ffffffff86811661010090920416146111ca576040517f0595ea2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111f58c8c8c8c8c5f8d8d8d8d8d8d8d6040516111e8929190614702565b6040518091039020612d18565b604080518b815263ffffffff8916602082015273ffffffffffffffffffffffffffffffffffffffff88811682840152861660608201526080810185905290517f1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d9181900360a00190a173ffffffffffffffffffffffffffffffffffffffff8616158015611286575063ffffffff8716155b156113a457606f5473ffffffffffffffffffffffffffffffffffffffff1661137b575f73ffffffffffffffffffffffffffffffffffffffff851684825b6040519080825280601f01601f1916602001820160405280156112ed576020820181803683370190505b506040516112fb9190614711565b5f6040518083038185875af1925050503d805f8114611335576040519150601f19603f3d011682016040523d82523d5f602084013e61133a565b606091505b5050905080611375576040517f6747a28800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506116d7565b606f5461139f9073ffffffffffffffffffffffffffffffffffffffff168585612df4565b6116d7565b606d5473ffffffffffffffffffffffffffffffffffffffff87811691161480156113f05750606d5463ffffffff8881167401000000000000000000000000000000000000000090920416145b15611414575f73ffffffffffffffffffffffffffffffffffffffff851684826112c3565b60685463ffffffff61010090910481169088160361144d5761139f73ffffffffffffffffffffffffffffffffffffffff87168585612e73565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e089901b1660208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606088901b1660248201525f90603801604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301205f818152606a90935291205490915073ffffffffffffffffffffffffffffffffffffffff16806116c9575f61154d8386868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250612ef992505050565b905061155a818888612df4565b80606a5f8581526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060405180604001604052808b63ffffffff1681526020018a73ffffffffffffffffffffffffffffffffffffffff16815250606b5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f015f6101000a81548163ffffffff021916908363ffffffff1602179055506020820151815f0160046101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050507f490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a398a8a8388886040516116bb95949392919061476e565b60405180910390a1506116d4565b6116d4818787612df4565b50505b6116e060018055565b505050505050505050505050565b60685460ff161561172b576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611733612ca5565b60685463ffffffff61010090910481169088160361177d576040517f0595ea2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8060608773ffffffffffffffffffffffffffffffffffffffff88166118a4578834146117d6576040517fb89240f500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606d54606e805473ffffffffffffffffffffffffffffffffffffffff831696507401000000000000000000000000000000000000000090920463ffffffff169450906118219061457a565b80601f016020809104026020016040519081016040528092919081815260200182805461184d9061457a565b80156118985780601f1061186f57610100808354040283529160200191611898565b820191905f5260205f20905b81548152906001019060200180831161187b57829003601f168201915b50505050509150611b14565b34156118dc576040517f798ee6f100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84156118ed576118ed888787613013565b606f5473ffffffffffffffffffffffffffffffffffffffff908116908916036119215761191a888a612915565b9050611b14565b73ffffffffffffffffffffffffffffffffffffffff8089165f908152606b602090815260409182902082518084019093525463ffffffff811683526401000000009004909216918101829052901515806119815750805163ffffffff1615155b156119a357611990898b612915565b6020820151825190965094509150611b07565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f9073ffffffffffffffffffffffffffffffffffffffff8b16906370a0823190602401602060405180830381865afa158015611a0d573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611a3191906147d0565b9050611a5573ffffffffffffffffffffffffffffffffffffffff8b1633308e6134d0565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f9073ffffffffffffffffffffffffffffffffffffffff8c16906370a0823190602401602060405180830381865afa158015611abf573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611ae391906147d0565b9050611aef82826147e7565b6068548c9850610100900463ffffffff169650935050505b611b1089610fbb565b9250505b7f501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b5f84868e8e8688605354604051611b539897969594939291906147fa565b60405180910390a1611b715f84868e8e868880519060200120613516565b8615611b7f57611b7f6126b1565b50505050610fb260018055565b606c5473ffffffffffffffffffffffffffffffffffffffff163314611bdd576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108956135e8565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1660208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606083901b1660248201525f9081906038016040516020818303038152906040528051906020012090505f60ff60f81b3083611c6f611000565b604051602001611c7f9190614711565b60405160208183030381529060405280519060200120604051602001611d0794939291907fff0000000000000000000000000000000000000000000000000000000000000094909416845260609290921b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660018401526015830152603582015260550190565b604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018152919052805160209091012095945050505050565b60685460ff1615611d83576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60685463ffffffff8681166101009092041614611dcc576040517f0595ea2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611deb8c8c8c8c8c60018d8d8d8d8d8d8d6040516111e8929190614702565b604080518b815263ffffffff8916602082015273ffffffffffffffffffffffffffffffffffffffff88811682840152861660608201526080810185905290517f1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d9181900360a00190a1606f545f9073ffffffffffffffffffffffffffffffffffffffff16611f6d578473ffffffffffffffffffffffffffffffffffffffff1684888a8686604051602401611ea29493929190614888565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f1806b5f20000000000000000000000000000000000000000000000000000000017905251611f239190614711565b5f6040518083038185875af1925050503d805f8114611f5d576040519150601f19603f3d011682016040523d82523d5f602084013e611f62565b606091505b505080915050612084565b606f54611f919073ffffffffffffffffffffffffffffffffffffffff168686612df4565b8473ffffffffffffffffffffffffffffffffffffffff1687898585604051602401611fbf9493929190614888565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f1806b5f200000000000000000000000000000000000000000000000000000000179052516120409190614711565b5f604051808303815f865af19150503d805f8114612079576040519150601f19603f3d011682016040523d82523d5f602084013e61207e565b606091505b50909150505b806120bb576040517f37e391c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50505050505050505050505050565b5f5460ff16607180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000060ff938416021790555f546002916101009091041615801561213357505f5460ff8083169116105b6121bf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610c1b565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff80841691909117610100179091556071547401000000000000000000000000000000000000000090041615612247576040517ff57ac68300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61224f612782565b606880547fffffffffffffff000000000000000000000000000000000000000000000000ff1661010063ffffffff8a16027fffffffffffffff0000000000000000000000000000000000000000ffffffffff16176501000000000073ffffffffffffffffffffffffffffffffffffffff8781169190910291909117909155606c80547fffffffffffffffffffffffff00000000000000000000000000000000000000001685831617905586166123425763ffffffff85161561233d576040517f1a874c1200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612494565b606d805463ffffffff871674010000000000000000000000000000000000000000027fffffffffffffffff00000000000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff891617179055606e6123af8382614907565b5061244e5f5f1b601260405160200161243a91906060808252600d908201527f5772617070656420457468657200000000000000000000000000000000000000608082015260a0602082018190526004908201527f574554480000000000000000000000000000000000000000000000000000000060c082015260ff91909116604082015260e00190565b604051602081830303815290604052612ef9565b606f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff929092169190911790555b61249c613676565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050607180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690555050505050565b60685460ff1615612567576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606880547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556040517f2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497905f90a1565b60685463ffffffff610100909104811690871603612606576040517f0595ea2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b6001606860019054906101000a900463ffffffff1633898989888860535460405161265a99989796959493929190614a1e565b60405180910390a161269b6001606860019054906101000a900463ffffffff1633898989888860405161268e929190614702565b6040518091039020613516565b82156126a9576126a96126b1565b505050505050565b6053546068805463ffffffff909216790100000000000000000000000000000000000000000000000000027fffffff00000000ffffffffffffffffffffffffffffffffffffffffffffffffff909216919091179081905573ffffffffffffffffffffffffffffffffffffffff65010000000000909104166333d6247d612735610a74565b6040518263ffffffff1660e01b815260040161275391815260200190565b5f604051808303815f87803b15801561276a575f5ffd5b505af115801561277c573d5f5f3e3d5ffd5b50505050565b5f6127c17fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035473ffffffffffffffffffffffffffffffffffffffff1690565b90508073ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561280c573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906128309190614aae565b607080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216918217905581906128c5576040517f54a0d80a00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610c1b565b50607054604080515f815273ffffffffffffffffffffffffffffffffffffffff90921660208301527fa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f29101610e0a565b6040517f9dc29fac000000000000000000000000000000000000000000000000000000008152336004820152602481018290525f9073ffffffffffffffffffffffffffffffffffffffff841690639dc29fac906044015f604051808303815f87803b158015612982575f5ffd5b505af1158015612994573d5f5f3e3d5ffd5b509395945050505050565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f06fdde030000000000000000000000000000000000000000000000000000000017905290516060915f91829173ffffffffffffffffffffffffffffffffffffffff861691612a209190614711565b5f60405180830381855afa9150503d805f8114612a58576040519150601f19603f3d011682016040523d82523d5f602084013e612a5d565b606091505b509150915081612aa2576040518060400160405280600781526020017f4e4f5f4e414d4500000000000000000000000000000000000000000000000000815250612aab565b612aab81613714565b949350505050565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f95d89b410000000000000000000000000000000000000000000000000000000017905290516060915f91829173ffffffffffffffffffffffffffffffffffffffff861691612b349190614711565b5f60405180830381855afa9150503d805f8114612b6c576040519150601f19603f3d011682016040523d82523d5f602084013e612b71565b606091505b509150915081612aa2576040518060400160405280600981526020017f4e4f5f53594d424f4c0000000000000000000000000000000000000000000000815250612aab565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f313ce5670000000000000000000000000000000000000000000000000000000017905290515f918291829173ffffffffffffffffffffffffffffffffffffffff861691612c369190614711565b5f60405180830381855afa9150503d805f8114612c6e576040519150601f19603f3d011682016040523d82523d5f602084013e612c73565b606091505b5091509150818015612c86575080516020145b612c91576012612aab565b80806020019051810190612aab9190614ad7565b600260015403612d11576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610c1b565b6002600155565b604080517fff0000000000000000000000000000000000000000000000000000000000000060f88a901b166020808301919091527fffffffff0000000000000000000000000000000000000000000000000000000060e08a811b821660218501527fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060608b811b82166025870152918a901b909216603985015287901b16603d83015260518201859052607180830185905283518084039091018152609190920190925280519101206116e0908d908d908d908d908d906138da565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8381166004830152602482018390528416906340c10f19906044015f604051808303815f87803b158015612e61575f5ffd5b505af1158015610fb2573d5f5f3e3d5ffd5b60405173ffffffffffffffffffffffffffffffffffffffff838116602483015260448201839052612ef491859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050613b21565b505050565b5f5f612f03611000565b604051602001612f139190614711565b6040516020818303038152906040529050838151602083015ff5915073ffffffffffffffffffffffffffffffffffffffff8216612f7c576040517f62d05d1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f5f85806020019051810190612f939190614b10565b9250925092508473ffffffffffffffffffffffffffffffffffffffff16631624f6c68484846040518463ffffffff1660e01b8152600401612fd6939291906145f8565b5f604051808303815f87803b158015612fed575f5ffd5b505af1158015612fff573d5f5f3e3d5ffd5b505050505050505092915050565b60018055565b5f6130216004828486614b88565b61302a91614baf565b90507f2afa5331000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000082160161325e575f808080808080613089896004818d614b88565b8101906130969190614c15565b96509650965096509650965096503373ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff1614613109576040517f912ecce700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff86163014613158576040517f750643af00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805173ffffffffffffffffffffffffffffffffffffffff89811660248301528881166044830152606482018890526084820187905260ff861660a483015260c4820185905260e48083018590528351808403909101815261010490920183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd505accf000000000000000000000000000000000000000000000000000000001790529151918d16916132119190614711565b5f604051808303815f865af19150503d805f811461324a576040519150601f19603f3d011682016040523d82523d5f602084013e61324f565b606091505b5050505050505050505061277c565b7fffffffff0000000000000000000000000000000000000000000000000000000081167f8fcbaf0c00000000000000000000000000000000000000000000000000000000146132d9576040517fe282c0ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f808080808080806132ee8a6004818e614b88565b8101906132fb9190614c81565b975097509750975097509750975097503373ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff1614613370576040517f912ecce700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff871630146133bf576040517f750643af00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805173ffffffffffffffffffffffffffffffffffffffff8a811660248301528981166044830152606482018990526084820188905286151560a483015260ff861660c483015260e482018590526101048083018590528351808403909101815261012490920183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f8fcbaf0c000000000000000000000000000000000000000000000000000000001790529151918e16916134819190614711565b5f604051808303815f865af19150503d805f81146134ba576040519150601f19603f3d011682016040523d82523d5f602084013e6134bf565b606091505b505050505050505050505050505050565b60405173ffffffffffffffffffffffffffffffffffffffff848116602483015283811660448301526064820183905261277c9186918216906323b872dd90608401612ead565b604080517fff0000000000000000000000000000000000000000000000000000000000000060f88a901b166020808301919091527fffffffff0000000000000000000000000000000000000000000000000000000060e08a811b821660218501527fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060608b811b82166025870152918a901b909216603985015287901b16603d8301526051820185905260718083018590528351808403909101815260919092019092528051910120610fb290613bb5565b60685460ff16613624576040517f5386698100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606880547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556040517f1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3905f90a1565b5f54610100900460ff1661370c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610c1b565b610895613c8d565b6060604082511061373357818060200190518101906109339190614cff565b815160200361389c575f5b602081108015613785575082818151811061375b5761375b6145cb565b01602001517fff000000000000000000000000000000000000000000000000000000000000001615155b1561379c578061379481614d31565b91505061373e565b805f036137de57505060408051808201909152601281527f4e4f545f56414c49445f454e434f44494e4700000000000000000000000000006020820152919050565b5f8167ffffffffffffffff8111156137f8576137f86143e9565b6040519080825280601f01601f191660200182016040528015613822576020820181803683370190505b5090505f5b8281101561389457848181518110613841576138416145cb565b602001015160f81c60f81b82828151811061385e5761385e6145cb565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600101613827565b509392505050565b505060408051808201909152601281527f4e4f545f56414c49445f454e434f44494e470000000000000000000000000000602082015290565b919050565b6068545f9065010000000000900473ffffffffffffffffffffffffffffffffffffffff1663257b363261390d8686613d23565b6040518263ffffffff1660e01b815260040161392b91815260200190565b6020604051808303815f875af1158015613947573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061396b91906147d0565b9050805f036139a5576040517e2f6fad00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8068010000000000000000871615613a5357869150816139d563ffffffff8216680100000000000000006146ef565b14613a0c576040517f071389e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613a18848a8489613d38565b613a4e576040517fe0417cec00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613b0c565b602087901c613a63816001614d68565b889350915082613a8763ffffffff821667ffffffff00000000602085901b166146ef565b14613abe576040517f071389e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613ad4613acc868c86613d4f565b8a8389613d38565b613b0a576040517fe0417cec00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b613b168282613ddc565b505050505050505050565b5f613b4273ffffffffffffffffffffffffffffffffffffffff841683613e9c565b905080515f14158015613b66575080806020019051810190613b649190614d84565b155b15612ef4576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610c1b565b806001613bc460206002614ec0565b613bce91906147e7565b60535410613c08576040517fef5ccf6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60535f8154613c1790614d31565b918290555090505f5b6020811015613c84578082901c600116600103613c53578260338260208110613c4b57613c4b6145cb565b015550505050565b613c7a60338260208110613c6957613c696145cb565b0154845f9182526020526040902090565b9250600101613c20565b50612ef4614ecb565b5f54610100900460ff1661300d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610c1b565b5f8281526020829052604081205b9392505050565b5f81613d45868686613d4f565b1495945050505050565b5f83815b6020811015613dd357600163ffffffff8516821c81169003613d9f57613d98858260208110613d8457613d846145cb565b6020020135835f9182526020526040902090565b9150613dcb565b613dc882868360208110613db557613db56145cb565b60200201355f9182526020526040902090565b91505b600101613d53565b50949350505050565b6068545f90610100900463ffffffff16158015613dff575063ffffffff82166001145b15613e11575063ffffffff8216613e39565b613e2664010000000063ffffffff84166146d8565b613e369063ffffffff85166146ef565b90505b600881901c5f8181526069602052604081208054600160ff861690811b91821892839055929091908183169003610fb2576040517f646cf55800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060613d3183835f845f5f8573ffffffffffffffffffffffffffffffffffffffff168486604051613ecd9190614711565b5f6040518083038185875af1925050503d805f8114613f07576040519150601f19603f3d011682016040523d82523d5f602084013e613f0c565b606091505b5091509150613f1c868383613f26565b9695505050505050565b606082613f3b57613f3682613fb5565b613d31565b8151158015613f5f575073ffffffffffffffffffffffffffffffffffffffff84163b155b15613fae576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610c1b565b5080613d31565b805115613fc55780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b803563ffffffff811681146138d5575f5ffd5b73ffffffffffffffffffffffffffffffffffffffff81168114613ff7575f5ffd5b5f5f6040838503121561403f575f5ffd5b61404883613ffa565b915060208301356140588161400d565b809150509250929050565b8015158114613ff7575f5ffd5b5f5f83601f840112614080575f5ffd5b50813567ffffffffffffffff811115614097575f5ffd5b6020830191508360208285010111156140ae575f5ffd5b9250929050565b5f5f5f5f5f608086880312156140c9575f5ffd5b6140d286613ffa565b945060208601356140e28161400d565b935060408601356140f281614063565b9250606086013567ffffffffffffffff81111561410d575f5ffd5b61411988828901614070565b969995985093965092949392505050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f613d31602083018461412a565b5f60208284031215614198575f5ffd5b8135613d318161400d565b5f602082840312156141b3575f5ffd5b5035919050565b5f5f5f5f5f5f60a087890312156141cf575f5ffd5b6141d887613ffa565b955060208701356141e88161400d565b94506040870135935060608701356141ff81614063565b9250608087013567ffffffffffffffff81111561421a575f5ffd5b61422689828a01614070565b979a9699509497509295939492505050565b5f5f60408385031215614249575f5ffd5b61425283613ffa565b915061426060208401613ffa565b90509250929050565b806104008101831015610933575f5ffd5b5f5f5f5f5f5f5f5f5f5f5f5f6109208d8f031215614296575f5ffd5b6142a08e8e614269565b9b506142b08e6104008f01614269565b9a506108008d013599506108208d013598506108408d013597506142d76108608e01613ffa565b96506142e76108808e013561400d565b6108808d013595506142fc6108a08e01613ffa565b94506108c08d013561430d8161400d565b93506108e08d0135925067ffffffffffffffff6109008e01351115614330575f5ffd5b6143418e6109008f01358f01614070565b81935080925050509295989b509295989b509295989b565b5f5f5f5f5f5f5f60c0888a03121561436f575f5ffd5b61437888613ffa565b965060208801356143888161400d565b955060408801359450606088013561439f8161400d565b935060808801356143af81614063565b925060a088013567ffffffffffffffff8111156143ca575f5ffd5b6143d68a828b01614070565b989b979a50959850939692959293505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561445d5761445d6143e9565b604052919050565b5f67ffffffffffffffff82111561447e5761447e6143e9565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b5f5f5f5f5f5f60c087890312156144bf575f5ffd5b6144c887613ffa565b955060208701356144d88161400d565b94506144e660408801613ffa565b935060608701356144f68161400d565b925060808701356145068161400d565b915060a087013567ffffffffffffffff811115614521575f5ffd5b8701601f81018913614531575f5ffd5b803561454461453f82614465565b614416565b8181528a6020838501011115614558575f5ffd5b816020840160208301375f602083830101528093505050509295509295509295565b600181811c9082168061458e57607f821691505b6020821081036145c5577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b606081525f61460a606083018661412a565b828103602084015261461c818661412a565b91505060ff83166040830152949350505050565b5f61463d61453f84614465565b9050828152838383011115614650575f5ffd5b8282602083015e5f602084830101529392505050565b5f60208284031215614676575f5ffd5b815167ffffffffffffffff81111561468c575f5ffd5b8201601f8101841361469c575f5ffd5b612aab84825160208401614630565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8082028115828204841417610933576109336146ab565b80820180821115610933576109336146ab565b818382375f9101908152919050565b5f82518060208501845e5f920191825250919050565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b63ffffffff8616815273ffffffffffffffffffffffffffffffffffffffff8516602082015273ffffffffffffffffffffffffffffffffffffffff84166040820152608060608201525f6147c5608083018486614727565b979650505050505050565b5f602082840312156147e0575f5ffd5b5051919050565b81810381811115610933576109336146ab565b60ff8916815263ffffffff8816602082015273ffffffffffffffffffffffffffffffffffffffff8716604082015263ffffffff8616606082015273ffffffffffffffffffffffffffffffffffffffff851660808201528360a082015261010060c08201525f61486d61010083018561412a565b905063ffffffff831660e08301529998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8516815263ffffffff84166020820152606060408201525f613f1c606083018486614727565b601f821115612ef457805f5260205f20601f840160051c810160208510156148e85750805b601f840160051c820191505b818110156109e1575f81556001016148f4565b815167ffffffffffffffff811115614921576149216143e9565b6149358161492f845461457a565b846148c3565b6020601f821160018114614986575f83156149505750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b1784556109e1565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b828110156149d357878501518255602094850194600190920191016149b3565b5084821015614a0f57868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b60ff8a16815263ffffffff8916602082015273ffffffffffffffffffffffffffffffffffffffff8816604082015263ffffffff8716606082015273ffffffffffffffffffffffffffffffffffffffff861660808201528460a082015261010060c08201525f614a9261010083018587614727565b905063ffffffff831660e08301529a9950505050505050505050565b5f60208284031215614abe575f5ffd5b8151613d318161400d565b60ff81168114613ff7575f5ffd5b5f60208284031215614ae7575f5ffd5b8151613d3181614ac9565b5f82601f830112614b01575f5ffd5b613d3183835160208501614630565b5f5f5f60608486031215614b22575f5ffd5b835167ffffffffffffffff811115614b38575f5ffd5b614b4486828701614af2565b935050602084015167ffffffffffffffff811115614b60575f5ffd5b614b6c86828701614af2565b9250506040840151614b7d81614ac9565b809150509250925092565b5f5f85851115614b96575f5ffd5b83861115614ba2575f5ffd5b5050820193919092039150565b80357fffffffff000000000000000000000000000000000000000000000000000000008116906004841015614c0e577fffffffff00000000000000000000000000000000000000000000000000000000808560040360031b1b82161691505b5092915050565b5f5f5f5f5f5f5f60e0888a031215614c2b575f5ffd5b8735614c368161400d565b96506020880135614c468161400d565b955060408801359450606088013593506080880135614c6481614ac9565b9699959850939692959460a0840135945060c09093013592915050565b5f5f5f5f5f5f5f5f610100898b031215614c99575f5ffd5b8835614ca48161400d565b97506020890135614cb48161400d565b965060408901359550606089013594506080890135614cd281614063565b935060a0890135614ce281614ac9565b979a969950949793969295929450505060c08201359160e0013590565b5f60208284031215614d0f575f5ffd5b815167ffffffffffffffff811115614d25575f5ffd5b612aab84828501614af2565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614d6157614d616146ab565b5060010190565b63ffffffff8181168382160190811115610933576109336146ab565b5f60208284031215614d94575f5ffd5b8151613d3181614063565b6001815b6001841115614dda57808504811115614dbe57614dbe6146ab565b6001841615614dcc57908102905b60019390931c928002614da3565b935093915050565b5f82614df057506001610933565b81614dfc57505f610933565b8160018114614e125760028114614e1c57614e38565b6001915050610933565b60ff841115614e2d57614e2d6146ab565b50506001821b610933565b5060208310610133831016604e8410600b8410161715614e5b575081810a610933565b614e867fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484614d9f565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115614eb857614eb86146ab565b029392505050565b5f613d318383614de2565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52600160045260245ffdfea2646970667358221220a68de3821248d5916ae5df5864d0a8156ec41767286506e8a97fef1137c1704364736f6c634300081c00336080604052348015600e575f5ffd5b50610fbb8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610029575f3560e01c8063c514f24e1461002d575b5f5ffd5b61003561004b565b604051610042919061006a565b60405180910390f35b60405180610f000160405280610ec881526020016100be610ec8913981565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168401019150509291505056fe60806040819052631d97f74d60e11b81523390633b2fee9a90608490602090600481865afa158015610033573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100579190610433565b604080515f80825260208201909252905061007382825f6100e2565b50506100dd336001600160a01b03166338b8fbbb6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100b4573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100d89190610433565b61010d565b6104c8565b6100eb8361017a565b5f825111806100f75750805b156101085761010683836101b9565b505b505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61014c5f516020610e815f395f51905f52546001600160a01b031690565b604080516001600160a01b03928316815291841660208301520160405180910390a1610177816101e5565b50565b61018381610280565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606101de8383604051806060016040528060278152602001610ea160279139610314565b9392505050565b6001600160a01b03811661024f5760405162461bcd60e51b815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b805f516020610e815f395f51905f525b80546001600160a01b0319166001600160a01b039290921691909117905550565b6001600160a01b0381163b6102ed5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610246565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc61025f565b60605f5f856001600160a01b031685604051610330919061047b565b5f60405180830381855af49150503d805f8114610368576040519150601f19603f3d011682016040523d82523d5f602084013e61036d565b606091505b50909250905061037f86838387610389565b9695505050505050565b606083156103f75782515f036103f0576001600160a01b0385163b6103f05760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610246565b5081610401565b6104018383610409565b949350505050565b8151156104195781518083602001fd5b8060405162461bcd60e51b81526004016102469190610496565b5f60208284031215610443575f5ffd5b81516001600160a01b03811681146101de575f5ffd5b5f5b8381101561047357818101518382015260200161045b565b50505f910152565b5f825161048c818460208701610459565b9190910192915050565b602081525f82518060208401526104b4816040850160208701610459565b601f01601f19169190910160400192915050565b6109ac806104d55f395ff3fe60806040526004361061005d575f3560e01c80635c60da1b116100425780635c60da1b146100a65780638f283970146100e3578063f851a440146101025761006c565b80633659cfe6146100745780634f1ef286146100935761006c565b3661006c5761006a610116565b005b61006a610116565b34801561007f575f5ffd5b5061006a61008e366004610854565b610130565b61006a6100a136600461086d565b610178565b3480156100b1575f5ffd5b506100ba6101eb565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b3480156100ee575f5ffd5b5061006a6100fd366004610854565b610228565b34801561010d575f5ffd5b506100ba610255565b61011e610282565b61012e610129610359565b610362565b565b610138610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101705761016d8160405180602001604052805f8152505f6103bf565b50565b61016d610116565b610180610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101e3576101de8383838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250600192506103bf915050565b505050565b6101de610116565b5f6101f4610380565b73ffffffffffffffffffffffffffffffffffffffff16330361021d57610218610359565b905090565b610225610116565b90565b610230610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101705761016d816103e9565b5f61025e610380565b73ffffffffffffffffffffffffffffffffffffffff16330361021d57610218610380565b61028a610380565b73ffffffffffffffffffffffffffffffffffffffff16330361012e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604260248201527f5472616e73706172656e745570677261646561626c6550726f78793a2061646d60448201527f696e2063616e6e6f742066616c6c6261636b20746f2070726f7879207461726760648201527f6574000000000000000000000000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b5f61021861044a565b365f5f375f5f365f845af43d5f5f3e80801561037c573d5ff35b3d5ffd5b5f7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b5473ffffffffffffffffffffffffffffffffffffffff16919050565b6103c883610471565b5f825111806103d45750805b156101de576103e383836104bd565b50505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f610412610380565b6040805173ffffffffffffffffffffffffffffffffffffffff928316815291841660208301520160405180910390a161016d816104e9565b5f7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6103a3565b61047a816105f5565b60405173ffffffffffffffffffffffffffffffffffffffff8216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606104e28383604051806060016040528060278152602001610979602791396106c0565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff811661058c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610350565b807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b73ffffffffffffffffffffffffffffffffffffffff81163b610699576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201527f6f74206120636f6e7472616374000000000000000000000000000000000000006064820152608401610350565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6105af565b60605f5f8573ffffffffffffffffffffffffffffffffffffffff16856040516106e9919061090d565b5f60405180830381855af49150503d805f8114610721576040519150601f19603f3d011682016040523d82523d5f602084013e610726565b606091505b509150915061073786838387610741565b9695505050505050565b606083156107d65782515f036107cf5773ffffffffffffffffffffffffffffffffffffffff85163b6107cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610350565b50816107e0565b6107e083836107e8565b949350505050565b8151156107f85781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103509190610928565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f575f5ffd5b919050565b5f60208284031215610864575f5ffd5b6104e28261082c565b5f5f5f6040848603121561087f575f5ffd5b6108888461082c565b9250602084013567ffffffffffffffff8111156108a3575f5ffd5b8401601f810186136108b3575f5ffd5b803567ffffffffffffffff8111156108c9575f5ffd5b8660208284010111156108da575f5ffd5b939660209190910195509293505050565b5f5b838110156109055781810151838201526020016108ed565b50505f910152565b5f825161091e8184602087016108eb565b9190910192915050565b602081525f82518060208401526109468160408501602087016108eb565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a164736f6c634300081c000ab53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a264697066735822122022d6b4c88ed40072a29d174af4106e1d1402760c56f0574d6d384cdc677052c164736f6c634300081c00336080604052348015600e575f5ffd5b5060156019565b60c9565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560685760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c65780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b611c60806100d65f395ff3fe608060405234801561000f575f5ffd5b5060043610610115575f3560e01c806370a08231116100ad5780639dc29fac1161007d578063a9059cbb11610063578063a9059cbb146102c0578063d505accf146102d3578063dd62ed3e146102e6575f5ffd5b80639dc29fac1461024b578063a3c573eb1461025e575f5ffd5b806370a08231146102025780637ecebe001461021557806384b0196e1461022857806395d89b4114610243575f5ffd5b806323b872dd116100e857806323b872dd146101a0578063313ce567146101b35780633644e515146101e757806340c10f19146101ef575f5ffd5b806306fdde0314610119578063095ea7b3146101375780631624f6c61461015a57806318160ddd1461016f575b5f5ffd5b61012161034a565b60405161012e919061168e565b60405180910390f35b61014a6101453660046116cf565b610402565b604051901515815260200161012e565b61016d6101683660046117fc565b61041b565b005b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02545b60405190815260200161012e565b61014a6101ae366004611870565b6105f9565b7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d005460405160ff909116815260200161012e565b61019261061c565b61016d6101fd3660046116cf565b61062a565b6101926102103660046118aa565b6106b3565b6101926102233660046118aa565b610703565b61023061070d565b60405161012e97969594939291906118c3565b61012161080c565b61016d6102593660046116cf565b61085d565b7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d0054610100900473ffffffffffffffffffffffffffffffffffffffff1660405173ffffffffffffffffffffffffffffffffffffffff909116815260200161012e565b61014a6102ce3660046116cf565b6108e1565b61016d6102e1366004611982565b6108ee565b6101926102f43660046119e8565b73ffffffffffffffffffffffffffffffffffffffff9182165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b60605f7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b905080600301805461038090611a19565b80601f01602080910402602001604051908101604052809291908181526020018280546103ac90611a19565b80156103f75780601f106103ce576101008083540402835291602001916103f7565b820191905f5260205f20905b8154815290600101906020018083116103da57829003601f168201915b505050505091505090565b5f3361040f818585610ab6565b60019150505b92915050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f811580156104655750825b90505f8267ffffffffffffffff1660011480156104815750303b155b90508115801561048f575080155b156104c6576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156105275784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6105318888610ac3565b61053a88610ad9565b7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d00805460ff88167fffffffffffffffffffffff00000000000000000000000000000000000000000090911617610100330217905583156105ef5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b5f33610606858285610b23565b610611858585610c0f565b506001949350505050565b5f610625610cb8565b905090565b5f7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d008054909150610100900473ffffffffffffffffffffffffffffffffffffffff1633146106a4576040517f38da3b1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106ae8383610cc1565b505050565b5f807f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b73ffffffffffffffffffffffffffffffffffffffff9093165f9081526020939093525050604090205490565b5f61041582610d1b565b5f60608082808083817fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100805490915015801561074b57506001810154155b6107b6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f4549503731323a20556e696e697469616c697a6564000000000000000000000060448201526064015b60405180910390fd5b6107be610d25565b6107c6610d76565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009c939b5091995046985030975095509350915050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060917f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace009161038090611a19565b5f7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d008054909150610100900473ffffffffffffffffffffffffffffffffffffffff1633146108d7576040517f38da3b1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106ae8383610d9f565b5f3361040f818585610c0f565b8342111561092b576040517f62791302000000000000000000000000000000000000000000000000000000008152600481018590526024016107ad565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98888886109a28c73ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f610a0982610df9565b90505f610a1882878787610e40565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610a9f576040517f4b800e4600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301528b1660248201526044016107ad565b610aaa8a8a8a610ab6565b50505050505050505050565b6106ae8383836001610e6c565b610acb610fd6565b610ad5828261103f565b5050565b610ae1610fd6565b610b20816040518060400160405280600181526020017f31000000000000000000000000000000000000000000000000000000000000008152506110a2565b50565b73ffffffffffffffffffffffffffffffffffffffff8381165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610c095781811015610bfb576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064016107ad565b610c0984848484035f610e6c565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610c5e576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f60048201526024016107ad565b73ffffffffffffffffffffffffffffffffffffffff8216610cad576040517fec442f050000000000000000000000000000000000000000000000000000000081525f60048201526024016107ad565b6106ae838383611114565b5f6106256112e1565b73ffffffffffffffffffffffffffffffffffffffff8216610d10576040517fec442f050000000000000000000000000000000000000000000000000000000081525f60048201526024016107ad565b610ad55f8383611114565b5f61041582611354565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060917fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1009161038090611a19565b60605f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10061036f565b73ffffffffffffffffffffffffffffffffffffffff8216610dee576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f60048201526024016107ad565b610ad5825f83611114565b5f610415610e05610cb8565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b5f5f5f5f610e508888888861137c565b925092509250610e60828261146f565b50909695505050505050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff8516610edc576040517fe602df050000000000000000000000000000000000000000000000000000000081525f60048201526024016107ad565b73ffffffffffffffffffffffffffffffffffffffff8416610f2b576040517f94280d620000000000000000000000000000000000000000000000000000000081525f60048201526024016107ad565b73ffffffffffffffffffffffffffffffffffffffff8086165f90815260018301602090815260408083209388168352929052208390558115610fcf578373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92585604051610fc691815260200190565b60405180910390a35b5050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff1661103d576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b611047610fd6565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace007f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace036110938482611aae565b5060048101610c098382611aae565b6110aa610fd6565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1007fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1026110f68482611aae565b50600381016111058382611aae565b505f8082556001909101555050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff841661116e5781816002015f8282546111639190611bc5565b9091555061121e9050565b73ffffffffffffffffffffffffffffffffffffffff84165f90815260208290526040902054828110156111f3576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8616600482015260248101829052604481018490526064016107ad565b73ffffffffffffffffffffffffffffffffffffffff85165f9081526020839052604090209083900390555b73ffffffffffffffffffffffffffffffffffffffff8316611249576002810180548390039055611274565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526020829052604090208054830190555b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516112d391815260200190565b60405180910390a350505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f61130b611572565b6113136115ed565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f807f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006106d7565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156113b557505f91506003905082611465565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015611406573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811661145c57505f925060019150829050611465565b92505f91508190505b9450945094915050565b5f82600381111561148257611482611bfd565b0361148b575050565b600182600381111561149f5761149f611bfd565b036114d6576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028260038111156114ea576114ea611bfd565b03611524576040517ffce698f7000000000000000000000000000000000000000000000000000000008152600481018290526024016107ad565b600382600381111561153857611538611bfd565b03610ad5576040517fd78bce0c000000000000000000000000000000000000000000000000000000008152600481018290526024016107ad565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1008161159d610d25565b8051909150156115b557805160209091012092915050565b815480156115c4579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10081611618610d76565b80519091501561163057805160209091012092915050565b600182015480156115c4579392505050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f6116a06020830184611642565b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146116ca575f5ffd5b919050565b5f5f604083850312156116e0575f5ffd5b6116e9836116a7565b946020939093013593505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112611733575f5ffd5b813567ffffffffffffffff81111561174d5761174d6116f7565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff821117156117b9576117b96116f7565b6040528181528382016020018510156117d0575f5ffd5b816020850160208301375f918101602001919091529392505050565b803560ff811681146116ca575f5ffd5b5f5f5f6060848603121561180e575f5ffd5b833567ffffffffffffffff811115611824575f5ffd5b61183086828701611724565b935050602084013567ffffffffffffffff81111561184c575f5ffd5b61185886828701611724565b925050611867604085016117ec565b90509250925092565b5f5f5f60608486031215611882575f5ffd5b61188b846116a7565b9250611899602085016116a7565b929592945050506040919091013590565b5f602082840312156118ba575f5ffd5b6116a0826116a7565b7fff000000000000000000000000000000000000000000000000000000000000008816815260e060208201525f6118fd60e0830189611642565b828103604084015261190f8189611642565b6060840188905273ffffffffffffffffffffffffffffffffffffffff8716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015611971578351835260209384019390920191600101611953565b50909b9a5050505050505050505050565b5f5f5f5f5f5f5f60e0888a031215611998575f5ffd5b6119a1886116a7565b96506119af602089016116a7565b955060408801359450606088013593506119cb608089016117ec565b9699959850939692959460a0840135945060c09093013592915050565b5f5f604083850312156119f9575f5ffd5b611a02836116a7565b9150611a10602084016116a7565b90509250929050565b600181811c90821680611a2d57607f821691505b602082108103611a64577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b601f8211156106ae57805f5260205f20601f840160051c81016020851015611a8f5750805b601f840160051c820191505b81811015610fcf575f8155600101611a9b565b815167ffffffffffffffff811115611ac857611ac86116f7565b611adc81611ad68454611a19565b84611a6a565b6020601f821160018114611b2d575f8315611af75750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455610fcf565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015611b7a5787850151825560209485019460019092019101611b5a565b5084821015611bb657868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b80820180821115610415577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffdfea2646970667358221220cab1bda26a25f523842457c2ae87e0281ec40e493b88229fc95b7a72c369c0de64736f6c634300081c0033",
}

// Polygonzkevmbridgev2v1010ABI is the input ABI used to generate the binding from.
// Deprecated: Use Polygonzkevmbridgev2v1010MetaData.ABI instead.
var Polygonzkevmbridgev2v1010ABI = Polygonzkevmbridgev2v1010MetaData.ABI

// Polygonzkevmbridgev2v1010Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Polygonzkevmbridgev2v1010MetaData.Bin instead.
var Polygonzkevmbridgev2v1010Bin = Polygonzkevmbridgev2v1010MetaData.Bin

// DeployPolygonzkevmbridgev2v1010 deploys a new Ethereum contract, binding an instance of Polygonzkevmbridgev2v1010 to it.
func DeployPolygonzkevmbridgev2v1010(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Polygonzkevmbridgev2v1010, error) {
	parsed, err := Polygonzkevmbridgev2v1010MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Polygonzkevmbridgev2v1010Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonzkevmbridgev2v1010{Polygonzkevmbridgev2v1010Caller: Polygonzkevmbridgev2v1010Caller{contract: contract}, Polygonzkevmbridgev2v1010Transactor: Polygonzkevmbridgev2v1010Transactor{contract: contract}, Polygonzkevmbridgev2v1010Filterer: Polygonzkevmbridgev2v1010Filterer{contract: contract}}, nil
}

// Polygonzkevmbridgev2v1010 is an auto generated Go binding around an Ethereum contract.
type Polygonzkevmbridgev2v1010 struct {
	Polygonzkevmbridgev2v1010Caller     // Read-only binding to the contract
	Polygonzkevmbridgev2v1010Transactor // Write-only binding to the contract
	Polygonzkevmbridgev2v1010Filterer   // Log filterer for contract events
}

// Polygonzkevmbridgev2v1010Caller is an auto generated read-only Go binding around an Ethereum contract.
type Polygonzkevmbridgev2v1010Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Polygonzkevmbridgev2v1010Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Polygonzkevmbridgev2v1010Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Polygonzkevmbridgev2v1010Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Polygonzkevmbridgev2v1010Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Polygonzkevmbridgev2v1010Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Polygonzkevmbridgev2v1010Session struct {
	Contract     *Polygonzkevmbridgev2v1010 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// Polygonzkevmbridgev2v1010CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Polygonzkevmbridgev2v1010CallerSession struct {
	Contract *Polygonzkevmbridgev2v1010Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// Polygonzkevmbridgev2v1010TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Polygonzkevmbridgev2v1010TransactorSession struct {
	Contract     *Polygonzkevmbridgev2v1010Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// Polygonzkevmbridgev2v1010Raw is an auto generated low-level Go binding around an Ethereum contract.
type Polygonzkevmbridgev2v1010Raw struct {
	Contract *Polygonzkevmbridgev2v1010 // Generic contract binding to access the raw methods on
}

// Polygonzkevmbridgev2v1010CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Polygonzkevmbridgev2v1010CallerRaw struct {
	Contract *Polygonzkevmbridgev2v1010Caller // Generic read-only contract binding to access the raw methods on
}

// Polygonzkevmbridgev2v1010TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Polygonzkevmbridgev2v1010TransactorRaw struct {
	Contract *Polygonzkevmbridgev2v1010Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonzkevmbridgev2v1010 creates a new instance of Polygonzkevmbridgev2v1010, bound to a specific deployed contract.
func NewPolygonzkevmbridgev2v1010(address common.Address, backend bind.ContractBackend) (*Polygonzkevmbridgev2v1010, error) {
	contract, err := bindPolygonzkevmbridgev2v1010(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmbridgev2v1010{Polygonzkevmbridgev2v1010Caller: Polygonzkevmbridgev2v1010Caller{contract: contract}, Polygonzkevmbridgev2v1010Transactor: Polygonzkevmbridgev2v1010Transactor{contract: contract}, Polygonzkevmbridgev2v1010Filterer: Polygonzkevmbridgev2v1010Filterer{contract: contract}}, nil
}

// NewPolygonzkevmbridgev2v1010Caller creates a new read-only instance of Polygonzkevmbridgev2v1010, bound to a specific deployed contract.
func NewPolygonzkevmbridgev2v1010Caller(address common.Address, caller bind.ContractCaller) (*Polygonzkevmbridgev2v1010Caller, error) {
	contract, err := bindPolygonzkevmbridgev2v1010(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmbridgev2v1010Caller{contract: contract}, nil
}

// NewPolygonzkevmbridgev2v1010Transactor creates a new write-only instance of Polygonzkevmbridgev2v1010, bound to a specific deployed contract.
func NewPolygonzkevmbridgev2v1010Transactor(address common.Address, transactor bind.ContractTransactor) (*Polygonzkevmbridgev2v1010Transactor, error) {
	contract, err := bindPolygonzkevmbridgev2v1010(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmbridgev2v1010Transactor{contract: contract}, nil
}

// NewPolygonzkevmbridgev2v1010Filterer creates a new log filterer instance of Polygonzkevmbridgev2v1010, bound to a specific deployed contract.
func NewPolygonzkevmbridgev2v1010Filterer(address common.Address, filterer bind.ContractFilterer) (*Polygonzkevmbridgev2v1010Filterer, error) {
	contract, err := bindPolygonzkevmbridgev2v1010(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmbridgev2v1010Filterer{contract: contract}, nil
}

// bindPolygonzkevmbridgev2v1010 binds a generic wrapper to an already deployed contract.
func bindPolygonzkevmbridgev2v1010(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Polygonzkevmbridgev2v1010MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmbridgev2v1010.Contract.Polygonzkevmbridgev2v1010Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.Polygonzkevmbridgev2v1010Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.Polygonzkevmbridgev2v1010Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmbridgev2v1010.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.contract.Transact(opts, method, params...)
}

// BRIDGEVERSION is a free data retrieval call binding the contract method 0x65d6f654.
//
// Solidity: function BRIDGE_VERSION() view returns(string)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Caller) BRIDGEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonzkevmbridgev2v1010.contract.Call(opts, &out, "BRIDGE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BRIDGEVERSION is a free data retrieval call binding the contract method 0x65d6f654.
//
// Solidity: function BRIDGE_VERSION() view returns(string)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) BRIDGEVERSION() (string, error) {
	return _Polygonzkevmbridgev2v1010.Contract.BRIDGEVERSION(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// BRIDGEVERSION is a free data retrieval call binding the contract method 0x65d6f654.
//
// Solidity: function BRIDGE_VERSION() view returns(string)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010CallerSession) BRIDGEVERSION() (string, error) {
	return _Polygonzkevmbridgev2v1010.Contract.BRIDGEVERSION(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() view returns(bytes)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Caller) INITBYTECODETRANSPARENTPROXY(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevmbridgev2v1010.contract.Call(opts, &out, "INIT_BYTECODE_TRANSPARENT_PROXY")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() view returns(bytes)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) INITBYTECODETRANSPARENTPROXY() ([]byte, error) {
	return _Polygonzkevmbridgev2v1010.Contract.INITBYTECODETRANSPARENTPROXY(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() view returns(bytes)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010CallerSession) INITBYTECODETRANSPARENTPROXY() ([]byte, error) {
	return _Polygonzkevmbridgev2v1010.Contract.INITBYTECODETRANSPARENTPROXY(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Caller) WETHToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmbridgev2v1010.contract.Call(opts, &out, "WETHToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) WETHToken() (common.Address, error) {
	return _Polygonzkevmbridgev2v1010.Contract.WETHToken(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010CallerSession) WETHToken() (common.Address, error) {
	return _Polygonzkevmbridgev2v1010.Contract.WETHToken(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Caller) ClaimedBitMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmbridgev2v1010.contract.Call(opts, &out, "claimedBitMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) ClaimedBitMap(arg0 *big.Int) (*big.Int, error) {
	return _Polygonzkevmbridgev2v1010.Contract.ClaimedBitMap(&_Polygonzkevmbridgev2v1010.CallOpts, arg0)
}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010CallerSession) ClaimedBitMap(arg0 *big.Int) (*big.Int, error) {
	return _Polygonzkevmbridgev2v1010.Contract.ClaimedBitMap(&_Polygonzkevmbridgev2v1010.CallOpts, arg0)
}

// ComputeTokenProxyAddress is a free data retrieval call binding the contract method 0xf214e161.
//
// Solidity: function computeTokenProxyAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Caller) ComputeTokenProxyAddress(opts *bind.CallOpts, originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmbridgev2v1010.contract.Call(opts, &out, "computeTokenProxyAddress", originNetwork, originTokenAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeTokenProxyAddress is a free data retrieval call binding the contract method 0xf214e161.
//
// Solidity: function computeTokenProxyAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) ComputeTokenProxyAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _Polygonzkevmbridgev2v1010.Contract.ComputeTokenProxyAddress(&_Polygonzkevmbridgev2v1010.CallOpts, originNetwork, originTokenAddress)
}

// ComputeTokenProxyAddress is a free data retrieval call binding the contract method 0xf214e161.
//
// Solidity: function computeTokenProxyAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010CallerSession) ComputeTokenProxyAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _Polygonzkevmbridgev2v1010.Contract.ComputeTokenProxyAddress(&_Polygonzkevmbridgev2v1010.CallOpts, originNetwork, originTokenAddress)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Caller) DepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmbridgev2v1010.contract.Call(opts, &out, "depositCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) DepositCount() (*big.Int, error) {
	return _Polygonzkevmbridgev2v1010.Contract.DepositCount(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010CallerSession) DepositCount() (*big.Int, error) {
	return _Polygonzkevmbridgev2v1010.Contract.DepositCount(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Caller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmbridgev2v1010.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) GasTokenAddress() (common.Address, error) {
	return _Polygonzkevmbridgev2v1010.Contract.GasTokenAddress(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010CallerSession) GasTokenAddress() (common.Address, error) {
	return _Polygonzkevmbridgev2v1010.Contract.GasTokenAddress(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// GasTokenMetadata is a free data retrieval call binding the contract method 0x27aef4e8.
//
// Solidity: function gasTokenMetadata() view returns(bytes)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Caller) GasTokenMetadata(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevmbridgev2v1010.contract.Call(opts, &out, "gasTokenMetadata")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GasTokenMetadata is a free data retrieval call binding the contract method 0x27aef4e8.
//
// Solidity: function gasTokenMetadata() view returns(bytes)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) GasTokenMetadata() ([]byte, error) {
	return _Polygonzkevmbridgev2v1010.Contract.GasTokenMetadata(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// GasTokenMetadata is a free data retrieval call binding the contract method 0x27aef4e8.
//
// Solidity: function gasTokenMetadata() view returns(bytes)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010CallerSession) GasTokenMetadata() ([]byte, error) {
	return _Polygonzkevmbridgev2v1010.Contract.GasTokenMetadata(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Caller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonzkevmbridgev2v1010.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) GasTokenNetwork() (uint32, error) {
	return _Polygonzkevmbridgev2v1010.Contract.GasTokenNetwork(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010CallerSession) GasTokenNetwork() (uint32, error) {
	return _Polygonzkevmbridgev2v1010.Contract.GasTokenNetwork(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// GetProxiedTokensManager is a free data retrieval call binding the contract method 0x38b8fbbb.
//
// Solidity: function getProxiedTokensManager() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Caller) GetProxiedTokensManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmbridgev2v1010.contract.Call(opts, &out, "getProxiedTokensManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxiedTokensManager is a free data retrieval call binding the contract method 0x38b8fbbb.
//
// Solidity: function getProxiedTokensManager() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) GetProxiedTokensManager() (common.Address, error) {
	return _Polygonzkevmbridgev2v1010.Contract.GetProxiedTokensManager(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// GetProxiedTokensManager is a free data retrieval call binding the contract method 0x38b8fbbb.
//
// Solidity: function getProxiedTokensManager() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010CallerSession) GetProxiedTokensManager() (common.Address, error) {
	return _Polygonzkevmbridgev2v1010.Contract.GetProxiedTokensManager(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Caller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmbridgev2v1010.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) GetRoot() ([32]byte, error) {
	return _Polygonzkevmbridgev2v1010.Contract.GetRoot(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010CallerSession) GetRoot() ([32]byte, error) {
	return _Polygonzkevmbridgev2v1010.Contract.GetRoot(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Caller) GetTokenMetadata(opts *bind.CallOpts, token common.Address) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevmbridgev2v1010.contract.Call(opts, &out, "getTokenMetadata", token)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) GetTokenMetadata(token common.Address) ([]byte, error) {
	return _Polygonzkevmbridgev2v1010.Contract.GetTokenMetadata(&_Polygonzkevmbridgev2v1010.CallOpts, token)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010CallerSession) GetTokenMetadata(token common.Address) ([]byte, error) {
	return _Polygonzkevmbridgev2v1010.Contract.GetTokenMetadata(&_Polygonzkevmbridgev2v1010.CallOpts, token)
}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Caller) GetTokenWrappedAddress(opts *bind.CallOpts, originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmbridgev2v1010.contract.Call(opts, &out, "getTokenWrappedAddress", originNetwork, originTokenAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) GetTokenWrappedAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _Polygonzkevmbridgev2v1010.Contract.GetTokenWrappedAddress(&_Polygonzkevmbridgev2v1010.CallOpts, originNetwork, originTokenAddress)
}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010CallerSession) GetTokenWrappedAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _Polygonzkevmbridgev2v1010.Contract.GetTokenWrappedAddress(&_Polygonzkevmbridgev2v1010.CallOpts, originNetwork, originTokenAddress)
}

// GetWrappedTokenBridgeImplementation is a free data retrieval call binding the contract method 0x3b2fee9a.
//
// Solidity: function getWrappedTokenBridgeImplementation() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Caller) GetWrappedTokenBridgeImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmbridgev2v1010.contract.Call(opts, &out, "getWrappedTokenBridgeImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWrappedTokenBridgeImplementation is a free data retrieval call binding the contract method 0x3b2fee9a.
//
// Solidity: function getWrappedTokenBridgeImplementation() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) GetWrappedTokenBridgeImplementation() (common.Address, error) {
	return _Polygonzkevmbridgev2v1010.Contract.GetWrappedTokenBridgeImplementation(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// GetWrappedTokenBridgeImplementation is a free data retrieval call binding the contract method 0x3b2fee9a.
//
// Solidity: function getWrappedTokenBridgeImplementation() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010CallerSession) GetWrappedTokenBridgeImplementation() (common.Address, error) {
	return _Polygonzkevmbridgev2v1010.Contract.GetWrappedTokenBridgeImplementation(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Caller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmbridgev2v1010.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) GlobalExitRootManager() (common.Address, error) {
	return _Polygonzkevmbridgev2v1010.Contract.GlobalExitRootManager(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010CallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonzkevmbridgev2v1010.Contract.GlobalExitRootManager(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// IsClaimed is a free data retrieval call binding the contract method 0xcc461632.
//
// Solidity: function isClaimed(uint32 leafIndex, uint32 sourceBridgeNetwork) view returns(bool)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Caller) IsClaimed(opts *bind.CallOpts, leafIndex uint32, sourceBridgeNetwork uint32) (bool, error) {
	var out []interface{}
	err := _Polygonzkevmbridgev2v1010.contract.Call(opts, &out, "isClaimed", leafIndex, sourceBridgeNetwork)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClaimed is a free data retrieval call binding the contract method 0xcc461632.
//
// Solidity: function isClaimed(uint32 leafIndex, uint32 sourceBridgeNetwork) view returns(bool)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) IsClaimed(leafIndex uint32, sourceBridgeNetwork uint32) (bool, error) {
	return _Polygonzkevmbridgev2v1010.Contract.IsClaimed(&_Polygonzkevmbridgev2v1010.CallOpts, leafIndex, sourceBridgeNetwork)
}

// IsClaimed is a free data retrieval call binding the contract method 0xcc461632.
//
// Solidity: function isClaimed(uint32 leafIndex, uint32 sourceBridgeNetwork) view returns(bool)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010CallerSession) IsClaimed(leafIndex uint32, sourceBridgeNetwork uint32) (bool, error) {
	return _Polygonzkevmbridgev2v1010.Contract.IsClaimed(&_Polygonzkevmbridgev2v1010.CallOpts, leafIndex, sourceBridgeNetwork)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Caller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Polygonzkevmbridgev2v1010.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) IsEmergencyState() (bool, error) {
	return _Polygonzkevmbridgev2v1010.Contract.IsEmergencyState(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010CallerSession) IsEmergencyState() (bool, error) {
	return _Polygonzkevmbridgev2v1010.Contract.IsEmergencyState(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Caller) LastUpdatedDepositCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonzkevmbridgev2v1010.contract.Call(opts, &out, "lastUpdatedDepositCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) LastUpdatedDepositCount() (uint32, error) {
	return _Polygonzkevmbridgev2v1010.Contract.LastUpdatedDepositCount(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010CallerSession) LastUpdatedDepositCount() (uint32, error) {
	return _Polygonzkevmbridgev2v1010.Contract.LastUpdatedDepositCount(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Caller) NetworkID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonzkevmbridgev2v1010.contract.Call(opts, &out, "networkID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) NetworkID() (uint32, error) {
	return _Polygonzkevmbridgev2v1010.Contract.NetworkID(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010CallerSession) NetworkID() (uint32, error) {
	return _Polygonzkevmbridgev2v1010.Contract.NetworkID(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// PendingProxiedTokensManager is a free data retrieval call binding the contract method 0xece93c6f.
//
// Solidity: function pendingProxiedTokensManager() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Caller) PendingProxiedTokensManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmbridgev2v1010.contract.Call(opts, &out, "pendingProxiedTokensManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingProxiedTokensManager is a free data retrieval call binding the contract method 0xece93c6f.
//
// Solidity: function pendingProxiedTokensManager() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) PendingProxiedTokensManager() (common.Address, error) {
	return _Polygonzkevmbridgev2v1010.Contract.PendingProxiedTokensManager(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// PendingProxiedTokensManager is a free data retrieval call binding the contract method 0xece93c6f.
//
// Solidity: function pendingProxiedTokensManager() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010CallerSession) PendingProxiedTokensManager() (common.Address, error) {
	return _Polygonzkevmbridgev2v1010.Contract.PendingProxiedTokensManager(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// PolygonRollupManager is a free data retrieval call binding the contract method 0x8ed7e3f2.
//
// Solidity: function polygonRollupManager() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Caller) PolygonRollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmbridgev2v1010.contract.Call(opts, &out, "polygonRollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PolygonRollupManager is a free data retrieval call binding the contract method 0x8ed7e3f2.
//
// Solidity: function polygonRollupManager() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) PolygonRollupManager() (common.Address, error) {
	return _Polygonzkevmbridgev2v1010.Contract.PolygonRollupManager(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// PolygonRollupManager is a free data retrieval call binding the contract method 0x8ed7e3f2.
//
// Solidity: function polygonRollupManager() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010CallerSession) PolygonRollupManager() (common.Address, error) {
	return _Polygonzkevmbridgev2v1010.Contract.PolygonRollupManager(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// ProxiedTokensManager is a free data retrieval call binding the contract method 0x6e4ecfed.
//
// Solidity: function proxiedTokensManager() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Caller) ProxiedTokensManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmbridgev2v1010.contract.Call(opts, &out, "proxiedTokensManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxiedTokensManager is a free data retrieval call binding the contract method 0x6e4ecfed.
//
// Solidity: function proxiedTokensManager() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) ProxiedTokensManager() (common.Address, error) {
	return _Polygonzkevmbridgev2v1010.Contract.ProxiedTokensManager(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// ProxiedTokensManager is a free data retrieval call binding the contract method 0x6e4ecfed.
//
// Solidity: function proxiedTokensManager() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010CallerSession) ProxiedTokensManager() (common.Address, error) {
	return _Polygonzkevmbridgev2v1010.Contract.ProxiedTokensManager(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Caller) TokenInfoToWrappedToken(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmbridgev2v1010.contract.Call(opts, &out, "tokenInfoToWrappedToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) TokenInfoToWrappedToken(arg0 [32]byte) (common.Address, error) {
	return _Polygonzkevmbridgev2v1010.Contract.TokenInfoToWrappedToken(&_Polygonzkevmbridgev2v1010.CallOpts, arg0)
}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010CallerSession) TokenInfoToWrappedToken(arg0 [32]byte) (common.Address, error) {
	return _Polygonzkevmbridgev2v1010.Contract.TokenInfoToWrappedToken(&_Polygonzkevmbridgev2v1010.CallOpts, arg0)
}

// WrappedTokenBytecodeStorer is a free data retrieval call binding the contract method 0x381fef6d.
//
// Solidity: function wrappedTokenBytecodeStorer() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Caller) WrappedTokenBytecodeStorer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmbridgev2v1010.contract.Call(opts, &out, "wrappedTokenBytecodeStorer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedTokenBytecodeStorer is a free data retrieval call binding the contract method 0x381fef6d.
//
// Solidity: function wrappedTokenBytecodeStorer() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) WrappedTokenBytecodeStorer() (common.Address, error) {
	return _Polygonzkevmbridgev2v1010.Contract.WrappedTokenBytecodeStorer(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// WrappedTokenBytecodeStorer is a free data retrieval call binding the contract method 0x381fef6d.
//
// Solidity: function wrappedTokenBytecodeStorer() view returns(address)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010CallerSession) WrappedTokenBytecodeStorer() (common.Address, error) {
	return _Polygonzkevmbridgev2v1010.Contract.WrappedTokenBytecodeStorer(&_Polygonzkevmbridgev2v1010.CallOpts)
}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address ) view returns(uint32 originNetwork, address originTokenAddress)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Caller) WrappedTokenToTokenInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	var out []interface{}
	err := _Polygonzkevmbridgev2v1010.contract.Call(opts, &out, "wrappedTokenToTokenInfo", arg0)

	outstruct := new(struct {
		OriginNetwork      uint32
		OriginTokenAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OriginNetwork = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.OriginTokenAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address ) view returns(uint32 originNetwork, address originTokenAddress)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) WrappedTokenToTokenInfo(arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	return _Polygonzkevmbridgev2v1010.Contract.WrappedTokenToTokenInfo(&_Polygonzkevmbridgev2v1010.CallOpts, arg0)
}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address ) view returns(uint32 originNetwork, address originTokenAddress)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010CallerSession) WrappedTokenToTokenInfo(arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	return _Polygonzkevmbridgev2v1010.Contract.WrappedTokenToTokenInfo(&_Polygonzkevmbridgev2v1010.CallOpts, arg0)
}

// AcceptProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8c668f1c.
//
// Solidity: function acceptProxiedTokensManagerRole() returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Transactor) AcceptProxiedTokensManagerRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.contract.Transact(opts, "acceptProxiedTokensManagerRole")
}

// AcceptProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8c668f1c.
//
// Solidity: function acceptProxiedTokensManagerRole() returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) AcceptProxiedTokensManagerRole() (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.AcceptProxiedTokensManagerRole(&_Polygonzkevmbridgev2v1010.TransactOpts)
}

// AcceptProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8c668f1c.
//
// Solidity: function acceptProxiedTokensManagerRole() returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010TransactorSession) AcceptProxiedTokensManagerRole() (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.AcceptProxiedTokensManagerRole(&_Polygonzkevmbridgev2v1010.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Transactor) ActivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.contract.Transact(opts, "activateEmergencyState")
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) ActivateEmergencyState() (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.ActivateEmergencyState(&_Polygonzkevmbridgev2v1010.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010TransactorSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.ActivateEmergencyState(&_Polygonzkevmbridgev2v1010.TransactOpts)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Transactor) BridgeAsset(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.contract.Transact(opts, "bridgeAsset", destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) BridgeAsset(destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.BridgeAsset(&_Polygonzkevmbridgev2v1010.TransactOpts, destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010TransactorSession) BridgeAsset(destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.BridgeAsset(&_Polygonzkevmbridgev2v1010.TransactOpts, destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Transactor) BridgeMessage(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.contract.Transact(opts, "bridgeMessage", destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) BridgeMessage(destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.BridgeMessage(&_Polygonzkevmbridgev2v1010.TransactOpts, destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010TransactorSession) BridgeMessage(destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.BridgeMessage(&_Polygonzkevmbridgev2v1010.TransactOpts, destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Transactor) BridgeMessageWETH(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.contract.Transact(opts, "bridgeMessageWETH", destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) BridgeMessageWETH(destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.BridgeMessageWETH(&_Polygonzkevmbridgev2v1010.TransactOpts, destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010TransactorSession) BridgeMessageWETH(destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.BridgeMessageWETH(&_Polygonzkevmbridgev2v1010.TransactOpts, destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Transactor) ClaimAsset(opts *bind.TransactOpts, smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.contract.Transact(opts, "claimAsset", smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) ClaimAsset(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.ClaimAsset(&_Polygonzkevmbridgev2v1010.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010TransactorSession) ClaimAsset(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.ClaimAsset(&_Polygonzkevmbridgev2v1010.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Transactor) ClaimMessage(opts *bind.TransactOpts, smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.contract.Transact(opts, "claimMessage", smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) ClaimMessage(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.ClaimMessage(&_Polygonzkevmbridgev2v1010.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010TransactorSession) ClaimMessage(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.ClaimMessage(&_Polygonzkevmbridgev2v1010.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Transactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.DeactivateEmergencyState(&_Polygonzkevmbridgev2v1010.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010TransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.DeactivateEmergencyState(&_Polygonzkevmbridgev2v1010.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Transactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) Initialize() (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.Initialize(&_Polygonzkevmbridgev2v1010.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010TransactorSession) Initialize() (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.Initialize(&_Polygonzkevmbridgev2v1010.TransactOpts)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata) returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Transactor) Initialize0(opts *bind.TransactOpts, _networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.contract.Transact(opts, "initialize0", _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata) returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) Initialize0(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.Initialize0(&_Polygonzkevmbridgev2v1010.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata) returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010TransactorSession) Initialize0(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.Initialize0(&_Polygonzkevmbridgev2v1010.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata)
}

// TransferProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8bd309c3.
//
// Solidity: function transferProxiedTokensManagerRole(address newProxiedTokensManager) returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Transactor) TransferProxiedTokensManagerRole(opts *bind.TransactOpts, newProxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.contract.Transact(opts, "transferProxiedTokensManagerRole", newProxiedTokensManager)
}

// TransferProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8bd309c3.
//
// Solidity: function transferProxiedTokensManagerRole(address newProxiedTokensManager) returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) TransferProxiedTokensManagerRole(newProxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.TransferProxiedTokensManagerRole(&_Polygonzkevmbridgev2v1010.TransactOpts, newProxiedTokensManager)
}

// TransferProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8bd309c3.
//
// Solidity: function transferProxiedTokensManagerRole(address newProxiedTokensManager) returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010TransactorSession) TransferProxiedTokensManagerRole(newProxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.TransferProxiedTokensManagerRole(&_Polygonzkevmbridgev2v1010.TransactOpts, newProxiedTokensManager)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Transactor) UpdateGlobalExitRoot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.contract.Transact(opts, "updateGlobalExitRoot")
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Session) UpdateGlobalExitRoot() (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.UpdateGlobalExitRoot(&_Polygonzkevmbridgev2v1010.TransactOpts)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010TransactorSession) UpdateGlobalExitRoot() (*types.Transaction, error) {
	return _Polygonzkevmbridgev2v1010.Contract.UpdateGlobalExitRoot(&_Polygonzkevmbridgev2v1010.TransactOpts)
}

// Polygonzkevmbridgev2v1010AcceptProxiedTokensManagerRoleIterator is returned from FilterAcceptProxiedTokensManagerRole and is used to iterate over the raw logs and unpacked data for AcceptProxiedTokensManagerRole events raised by the Polygonzkevmbridgev2v1010 contract.
type Polygonzkevmbridgev2v1010AcceptProxiedTokensManagerRoleIterator struct {
	Event *Polygonzkevmbridgev2v1010AcceptProxiedTokensManagerRole // Event containing the contract specifics and raw log

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
func (it *Polygonzkevmbridgev2v1010AcceptProxiedTokensManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonzkevmbridgev2v1010AcceptProxiedTokensManagerRole)
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
		it.Event = new(Polygonzkevmbridgev2v1010AcceptProxiedTokensManagerRole)
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
func (it *Polygonzkevmbridgev2v1010AcceptProxiedTokensManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonzkevmbridgev2v1010AcceptProxiedTokensManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonzkevmbridgev2v1010AcceptProxiedTokensManagerRole represents a AcceptProxiedTokensManagerRole event raised by the Polygonzkevmbridgev2v1010 contract.
type Polygonzkevmbridgev2v1010AcceptProxiedTokensManagerRole struct {
	OldProxiedTokensManager common.Address
	NewProxiedTokensManager common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterAcceptProxiedTokensManagerRole is a free log retrieval operation binding the contract event 0xa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f2.
//
// Solidity: event AcceptProxiedTokensManagerRole(address oldProxiedTokensManager, address newProxiedTokensManager)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Filterer) FilterAcceptProxiedTokensManagerRole(opts *bind.FilterOpts) (*Polygonzkevmbridgev2v1010AcceptProxiedTokensManagerRoleIterator, error) {

	logs, sub, err := _Polygonzkevmbridgev2v1010.contract.FilterLogs(opts, "AcceptProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmbridgev2v1010AcceptProxiedTokensManagerRoleIterator{contract: _Polygonzkevmbridgev2v1010.contract, event: "AcceptProxiedTokensManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptProxiedTokensManagerRole is a free log subscription operation binding the contract event 0xa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f2.
//
// Solidity: event AcceptProxiedTokensManagerRole(address oldProxiedTokensManager, address newProxiedTokensManager)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Filterer) WatchAcceptProxiedTokensManagerRole(opts *bind.WatchOpts, sink chan<- *Polygonzkevmbridgev2v1010AcceptProxiedTokensManagerRole) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmbridgev2v1010.contract.WatchLogs(opts, "AcceptProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonzkevmbridgev2v1010AcceptProxiedTokensManagerRole)
				if err := _Polygonzkevmbridgev2v1010.contract.UnpackLog(event, "AcceptProxiedTokensManagerRole", log); err != nil {
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

// ParseAcceptProxiedTokensManagerRole is a log parse operation binding the contract event 0xa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f2.
//
// Solidity: event AcceptProxiedTokensManagerRole(address oldProxiedTokensManager, address newProxiedTokensManager)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Filterer) ParseAcceptProxiedTokensManagerRole(log types.Log) (*Polygonzkevmbridgev2v1010AcceptProxiedTokensManagerRole, error) {
	event := new(Polygonzkevmbridgev2v1010AcceptProxiedTokensManagerRole)
	if err := _Polygonzkevmbridgev2v1010.contract.UnpackLog(event, "AcceptProxiedTokensManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonzkevmbridgev2v1010BridgeEventIterator is returned from FilterBridgeEvent and is used to iterate over the raw logs and unpacked data for BridgeEvent events raised by the Polygonzkevmbridgev2v1010 contract.
type Polygonzkevmbridgev2v1010BridgeEventIterator struct {
	Event *Polygonzkevmbridgev2v1010BridgeEvent // Event containing the contract specifics and raw log

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
func (it *Polygonzkevmbridgev2v1010BridgeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonzkevmbridgev2v1010BridgeEvent)
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
		it.Event = new(Polygonzkevmbridgev2v1010BridgeEvent)
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
func (it *Polygonzkevmbridgev2v1010BridgeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonzkevmbridgev2v1010BridgeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonzkevmbridgev2v1010BridgeEvent represents a BridgeEvent event raised by the Polygonzkevmbridgev2v1010 contract.
type Polygonzkevmbridgev2v1010BridgeEvent struct {
	LeafType           uint8
	OriginNetwork      uint32
	OriginAddress      common.Address
	DestinationNetwork uint32
	DestinationAddress common.Address
	Amount             *big.Int
	Metadata           []byte
	DepositCount       uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBridgeEvent is a free log retrieval operation binding the contract event 0x501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b.
//
// Solidity: event BridgeEvent(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata, uint32 depositCount)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Filterer) FilterBridgeEvent(opts *bind.FilterOpts) (*Polygonzkevmbridgev2v1010BridgeEventIterator, error) {

	logs, sub, err := _Polygonzkevmbridgev2v1010.contract.FilterLogs(opts, "BridgeEvent")
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmbridgev2v1010BridgeEventIterator{contract: _Polygonzkevmbridgev2v1010.contract, event: "BridgeEvent", logs: logs, sub: sub}, nil
}

// WatchBridgeEvent is a free log subscription operation binding the contract event 0x501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b.
//
// Solidity: event BridgeEvent(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata, uint32 depositCount)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Filterer) WatchBridgeEvent(opts *bind.WatchOpts, sink chan<- *Polygonzkevmbridgev2v1010BridgeEvent) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmbridgev2v1010.contract.WatchLogs(opts, "BridgeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonzkevmbridgev2v1010BridgeEvent)
				if err := _Polygonzkevmbridgev2v1010.contract.UnpackLog(event, "BridgeEvent", log); err != nil {
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

// ParseBridgeEvent is a log parse operation binding the contract event 0x501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b.
//
// Solidity: event BridgeEvent(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata, uint32 depositCount)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Filterer) ParseBridgeEvent(log types.Log) (*Polygonzkevmbridgev2v1010BridgeEvent, error) {
	event := new(Polygonzkevmbridgev2v1010BridgeEvent)
	if err := _Polygonzkevmbridgev2v1010.contract.UnpackLog(event, "BridgeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonzkevmbridgev2v1010ClaimEventIterator is returned from FilterClaimEvent and is used to iterate over the raw logs and unpacked data for ClaimEvent events raised by the Polygonzkevmbridgev2v1010 contract.
type Polygonzkevmbridgev2v1010ClaimEventIterator struct {
	Event *Polygonzkevmbridgev2v1010ClaimEvent // Event containing the contract specifics and raw log

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
func (it *Polygonzkevmbridgev2v1010ClaimEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonzkevmbridgev2v1010ClaimEvent)
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
		it.Event = new(Polygonzkevmbridgev2v1010ClaimEvent)
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
func (it *Polygonzkevmbridgev2v1010ClaimEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonzkevmbridgev2v1010ClaimEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonzkevmbridgev2v1010ClaimEvent represents a ClaimEvent event raised by the Polygonzkevmbridgev2v1010 contract.
type Polygonzkevmbridgev2v1010ClaimEvent struct {
	GlobalIndex        *big.Int
	OriginNetwork      uint32
	OriginAddress      common.Address
	DestinationAddress common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterClaimEvent is a free log retrieval operation binding the contract event 0x1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d.
//
// Solidity: event ClaimEvent(uint256 globalIndex, uint32 originNetwork, address originAddress, address destinationAddress, uint256 amount)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Filterer) FilterClaimEvent(opts *bind.FilterOpts) (*Polygonzkevmbridgev2v1010ClaimEventIterator, error) {

	logs, sub, err := _Polygonzkevmbridgev2v1010.contract.FilterLogs(opts, "ClaimEvent")
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmbridgev2v1010ClaimEventIterator{contract: _Polygonzkevmbridgev2v1010.contract, event: "ClaimEvent", logs: logs, sub: sub}, nil
}

// WatchClaimEvent is a free log subscription operation binding the contract event 0x1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d.
//
// Solidity: event ClaimEvent(uint256 globalIndex, uint32 originNetwork, address originAddress, address destinationAddress, uint256 amount)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Filterer) WatchClaimEvent(opts *bind.WatchOpts, sink chan<- *Polygonzkevmbridgev2v1010ClaimEvent) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmbridgev2v1010.contract.WatchLogs(opts, "ClaimEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonzkevmbridgev2v1010ClaimEvent)
				if err := _Polygonzkevmbridgev2v1010.contract.UnpackLog(event, "ClaimEvent", log); err != nil {
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

// ParseClaimEvent is a log parse operation binding the contract event 0x1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d.
//
// Solidity: event ClaimEvent(uint256 globalIndex, uint32 originNetwork, address originAddress, address destinationAddress, uint256 amount)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Filterer) ParseClaimEvent(log types.Log) (*Polygonzkevmbridgev2v1010ClaimEvent, error) {
	event := new(Polygonzkevmbridgev2v1010ClaimEvent)
	if err := _Polygonzkevmbridgev2v1010.contract.UnpackLog(event, "ClaimEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonzkevmbridgev2v1010EmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the Polygonzkevmbridgev2v1010 contract.
type Polygonzkevmbridgev2v1010EmergencyStateActivatedIterator struct {
	Event *Polygonzkevmbridgev2v1010EmergencyStateActivated // Event containing the contract specifics and raw log

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
func (it *Polygonzkevmbridgev2v1010EmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonzkevmbridgev2v1010EmergencyStateActivated)
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
		it.Event = new(Polygonzkevmbridgev2v1010EmergencyStateActivated)
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
func (it *Polygonzkevmbridgev2v1010EmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonzkevmbridgev2v1010EmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonzkevmbridgev2v1010EmergencyStateActivated represents a EmergencyStateActivated event raised by the Polygonzkevmbridgev2v1010 contract.
type Polygonzkevmbridgev2v1010EmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Filterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*Polygonzkevmbridgev2v1010EmergencyStateActivatedIterator, error) {

	logs, sub, err := _Polygonzkevmbridgev2v1010.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmbridgev2v1010EmergencyStateActivatedIterator{contract: _Polygonzkevmbridgev2v1010.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Filterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *Polygonzkevmbridgev2v1010EmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmbridgev2v1010.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonzkevmbridgev2v1010EmergencyStateActivated)
				if err := _Polygonzkevmbridgev2v1010.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
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

// ParseEmergencyStateActivated is a log parse operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Filterer) ParseEmergencyStateActivated(log types.Log) (*Polygonzkevmbridgev2v1010EmergencyStateActivated, error) {
	event := new(Polygonzkevmbridgev2v1010EmergencyStateActivated)
	if err := _Polygonzkevmbridgev2v1010.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonzkevmbridgev2v1010EmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the Polygonzkevmbridgev2v1010 contract.
type Polygonzkevmbridgev2v1010EmergencyStateDeactivatedIterator struct {
	Event *Polygonzkevmbridgev2v1010EmergencyStateDeactivated // Event containing the contract specifics and raw log

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
func (it *Polygonzkevmbridgev2v1010EmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonzkevmbridgev2v1010EmergencyStateDeactivated)
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
		it.Event = new(Polygonzkevmbridgev2v1010EmergencyStateDeactivated)
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
func (it *Polygonzkevmbridgev2v1010EmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonzkevmbridgev2v1010EmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonzkevmbridgev2v1010EmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the Polygonzkevmbridgev2v1010 contract.
type Polygonzkevmbridgev2v1010EmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Filterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*Polygonzkevmbridgev2v1010EmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _Polygonzkevmbridgev2v1010.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmbridgev2v1010EmergencyStateDeactivatedIterator{contract: _Polygonzkevmbridgev2v1010.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Filterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *Polygonzkevmbridgev2v1010EmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmbridgev2v1010.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonzkevmbridgev2v1010EmergencyStateDeactivated)
				if err := _Polygonzkevmbridgev2v1010.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
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

// ParseEmergencyStateDeactivated is a log parse operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Filterer) ParseEmergencyStateDeactivated(log types.Log) (*Polygonzkevmbridgev2v1010EmergencyStateDeactivated, error) {
	event := new(Polygonzkevmbridgev2v1010EmergencyStateDeactivated)
	if err := _Polygonzkevmbridgev2v1010.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonzkevmbridgev2v1010InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonzkevmbridgev2v1010 contract.
type Polygonzkevmbridgev2v1010InitializedIterator struct {
	Event *Polygonzkevmbridgev2v1010Initialized // Event containing the contract specifics and raw log

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
func (it *Polygonzkevmbridgev2v1010InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonzkevmbridgev2v1010Initialized)
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
		it.Event = new(Polygonzkevmbridgev2v1010Initialized)
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
func (it *Polygonzkevmbridgev2v1010InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonzkevmbridgev2v1010InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonzkevmbridgev2v1010Initialized represents a Initialized event raised by the Polygonzkevmbridgev2v1010 contract.
type Polygonzkevmbridgev2v1010Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Filterer) FilterInitialized(opts *bind.FilterOpts) (*Polygonzkevmbridgev2v1010InitializedIterator, error) {

	logs, sub, err := _Polygonzkevmbridgev2v1010.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmbridgev2v1010InitializedIterator{contract: _Polygonzkevmbridgev2v1010.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Polygonzkevmbridgev2v1010Initialized) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmbridgev2v1010.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonzkevmbridgev2v1010Initialized)
				if err := _Polygonzkevmbridgev2v1010.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Filterer) ParseInitialized(log types.Log) (*Polygonzkevmbridgev2v1010Initialized, error) {
	event := new(Polygonzkevmbridgev2v1010Initialized)
	if err := _Polygonzkevmbridgev2v1010.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonzkevmbridgev2v1010NewWrappedTokenIterator is returned from FilterNewWrappedToken and is used to iterate over the raw logs and unpacked data for NewWrappedToken events raised by the Polygonzkevmbridgev2v1010 contract.
type Polygonzkevmbridgev2v1010NewWrappedTokenIterator struct {
	Event *Polygonzkevmbridgev2v1010NewWrappedToken // Event containing the contract specifics and raw log

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
func (it *Polygonzkevmbridgev2v1010NewWrappedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonzkevmbridgev2v1010NewWrappedToken)
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
		it.Event = new(Polygonzkevmbridgev2v1010NewWrappedToken)
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
func (it *Polygonzkevmbridgev2v1010NewWrappedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonzkevmbridgev2v1010NewWrappedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonzkevmbridgev2v1010NewWrappedToken represents a NewWrappedToken event raised by the Polygonzkevmbridgev2v1010 contract.
type Polygonzkevmbridgev2v1010NewWrappedToken struct {
	OriginNetwork       uint32
	OriginTokenAddress  common.Address
	WrappedTokenAddress common.Address
	Metadata            []byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewWrappedToken is a free log retrieval operation binding the contract event 0x490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a39.
//
// Solidity: event NewWrappedToken(uint32 originNetwork, address originTokenAddress, address wrappedTokenAddress, bytes metadata)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Filterer) FilterNewWrappedToken(opts *bind.FilterOpts) (*Polygonzkevmbridgev2v1010NewWrappedTokenIterator, error) {

	logs, sub, err := _Polygonzkevmbridgev2v1010.contract.FilterLogs(opts, "NewWrappedToken")
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmbridgev2v1010NewWrappedTokenIterator{contract: _Polygonzkevmbridgev2v1010.contract, event: "NewWrappedToken", logs: logs, sub: sub}, nil
}

// WatchNewWrappedToken is a free log subscription operation binding the contract event 0x490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a39.
//
// Solidity: event NewWrappedToken(uint32 originNetwork, address originTokenAddress, address wrappedTokenAddress, bytes metadata)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Filterer) WatchNewWrappedToken(opts *bind.WatchOpts, sink chan<- *Polygonzkevmbridgev2v1010NewWrappedToken) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmbridgev2v1010.contract.WatchLogs(opts, "NewWrappedToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonzkevmbridgev2v1010NewWrappedToken)
				if err := _Polygonzkevmbridgev2v1010.contract.UnpackLog(event, "NewWrappedToken", log); err != nil {
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

// ParseNewWrappedToken is a log parse operation binding the contract event 0x490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a39.
//
// Solidity: event NewWrappedToken(uint32 originNetwork, address originTokenAddress, address wrappedTokenAddress, bytes metadata)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Filterer) ParseNewWrappedToken(log types.Log) (*Polygonzkevmbridgev2v1010NewWrappedToken, error) {
	event := new(Polygonzkevmbridgev2v1010NewWrappedToken)
	if err := _Polygonzkevmbridgev2v1010.contract.UnpackLog(event, "NewWrappedToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonzkevmbridgev2v1010TransferProxiedTokensManagerRoleIterator is returned from FilterTransferProxiedTokensManagerRole and is used to iterate over the raw logs and unpacked data for TransferProxiedTokensManagerRole events raised by the Polygonzkevmbridgev2v1010 contract.
type Polygonzkevmbridgev2v1010TransferProxiedTokensManagerRoleIterator struct {
	Event *Polygonzkevmbridgev2v1010TransferProxiedTokensManagerRole // Event containing the contract specifics and raw log

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
func (it *Polygonzkevmbridgev2v1010TransferProxiedTokensManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonzkevmbridgev2v1010TransferProxiedTokensManagerRole)
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
		it.Event = new(Polygonzkevmbridgev2v1010TransferProxiedTokensManagerRole)
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
func (it *Polygonzkevmbridgev2v1010TransferProxiedTokensManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonzkevmbridgev2v1010TransferProxiedTokensManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonzkevmbridgev2v1010TransferProxiedTokensManagerRole represents a TransferProxiedTokensManagerRole event raised by the Polygonzkevmbridgev2v1010 contract.
type Polygonzkevmbridgev2v1010TransferProxiedTokensManagerRole struct {
	CurrentProxiedTokensManager common.Address
	NewProxiedTokensManager     common.Address
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterTransferProxiedTokensManagerRole is a free log retrieval operation binding the contract event 0x0a34baa3feb299aef9c05cb59c6e0c8e7c0bcc65cbf0a647e7a7c8a2411591e2.
//
// Solidity: event TransferProxiedTokensManagerRole(address currentProxiedTokensManager, address newProxiedTokensManager)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Filterer) FilterTransferProxiedTokensManagerRole(opts *bind.FilterOpts) (*Polygonzkevmbridgev2v1010TransferProxiedTokensManagerRoleIterator, error) {

	logs, sub, err := _Polygonzkevmbridgev2v1010.contract.FilterLogs(opts, "TransferProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmbridgev2v1010TransferProxiedTokensManagerRoleIterator{contract: _Polygonzkevmbridgev2v1010.contract, event: "TransferProxiedTokensManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferProxiedTokensManagerRole is a free log subscription operation binding the contract event 0x0a34baa3feb299aef9c05cb59c6e0c8e7c0bcc65cbf0a647e7a7c8a2411591e2.
//
// Solidity: event TransferProxiedTokensManagerRole(address currentProxiedTokensManager, address newProxiedTokensManager)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Filterer) WatchTransferProxiedTokensManagerRole(opts *bind.WatchOpts, sink chan<- *Polygonzkevmbridgev2v1010TransferProxiedTokensManagerRole) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmbridgev2v1010.contract.WatchLogs(opts, "TransferProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonzkevmbridgev2v1010TransferProxiedTokensManagerRole)
				if err := _Polygonzkevmbridgev2v1010.contract.UnpackLog(event, "TransferProxiedTokensManagerRole", log); err != nil {
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

// ParseTransferProxiedTokensManagerRole is a log parse operation binding the contract event 0x0a34baa3feb299aef9c05cb59c6e0c8e7c0bcc65cbf0a647e7a7c8a2411591e2.
//
// Solidity: event TransferProxiedTokensManagerRole(address currentProxiedTokensManager, address newProxiedTokensManager)
func (_Polygonzkevmbridgev2v1010 *Polygonzkevmbridgev2v1010Filterer) ParseTransferProxiedTokensManagerRole(log types.Log) (*Polygonzkevmbridgev2v1010TransferProxiedTokensManagerRole, error) {
	event := new(Polygonzkevmbridgev2v1010TransferProxiedTokensManagerRole)
	if err := _Polygonzkevmbridgev2v1010.contract.UnpackLog(event, "TransferProxiedTokensManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

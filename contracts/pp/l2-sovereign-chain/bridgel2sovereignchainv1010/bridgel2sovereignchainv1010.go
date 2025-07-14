// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridgel2sovereignchainv1010

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

// Bridgel2sovereignchainv1010MetaData contains all meta data concerning the Bridgel2sovereignchainv1010 contract.
var Bridgel2sovereignchainv1010MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountDoesNotMatchMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeAddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DestinationNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmergencyStateNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EtherTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedProxyDeployment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputArraysLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGlobalIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxyAdmin\",\"type\":\"address\"}],\"name\":\"InvalidProxyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSmtProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSovereignWETHAddressParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroNetworkID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxyAdmin\",\"type\":\"address\"}],\"name\":\"InvalidZeroProxyAdminOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"localBalanceTreeAmount\",\"type\":\"uint256\"}],\"name\":\"LocalBalanceTreeOverflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"localBalanceTreeAmount\",\"type\":\"uint256\"}],\"name\":\"LocalBalanceTreeUnderflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MsgValueNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeTokenIsEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoValueInMessagesOnGasTokenNetworks\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyBridgeManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyBridgePauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyBridgeUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootRemover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingEmergencyBridgePauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingEmergencyBridgeUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingProxiedTokensManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProxiedTokensManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OriginNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyMapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyUpdated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotMapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotRemapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WETHRemappingNotSupportedOnGasTokenNetworks\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldEmergencyBridgePauser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newEmergencyBridgePauser\",\"type\":\"address\"}],\"name\":\"AcceptEmergencyBridgePauserRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldEmergencyBridgeUnpauser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newEmergencyBridgeUnpauser\",\"type\":\"address\"}],\"name\":\"AcceptEmergencyBridgeUnpauserRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldProxiedTokensManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newProxiedTokensManager\",\"type\":\"address\"}],\"name\":\"AcceptProxiedTokensManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"depositCount\",\"type\":\"uint32\"}],\"name\":\"BridgeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"legacyTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"updatedTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MigrateLegacyToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wrappedTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"NewWrappedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignTokenAddress\",\"type\":\"address\"}],\"name\":\"RemoveLegacySovereignTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridgeManager\",\"type\":\"address\"}],\"name\":\"SetBridgeManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tokenInfoHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SetInitialLocalBalanceTreeAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"SetSovereignTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignWETHTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"SetSovereignWETHAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentEmergencyBridgePauser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newEmergencyBridgePauser\",\"type\":\"address\"}],\"name\":\"TransferEmergencyBridgePauserRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentEmergencyBridgeUnpauser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newEmergencyBridgeUnpauser\",\"type\":\"address\"}],\"name\":\"TransferEmergencyBridgeUnpauserRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentProxiedTokensManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newProxiedTokensManager\",\"type\":\"address\"}],\"name\":\"TransferProxiedTokensManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"claimedGlobalIndex\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newClaimedGlobalIndexHashChain\",\"type\":\"bytes32\"}],\"name\":\"UpdatedClaimedGlobalIndexHashChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"unsetGlobalIndex\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newUnsetGlobalIndexHashChain\",\"type\":\"bytes32\"}],\"name\":\"UpdatedUnsetGlobalIndexHashChain\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BRIDGE_SOVEREIGN_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BRIDGE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_BYTECODE_TRANSPARENT_PROXY\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETHToken\",\"outputs\":[{\"internalType\":\"contractITokenWrappedBridgeUpgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptEmergencyBridgePauserRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptEmergencyBridgeUnpauserRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptProxiedTokensManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"permitData\",\"type\":\"bytes\"}],\"name\":\"bridgeAsset\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountWETH\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessageWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"calculateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimedBitMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimedGlobalIndexHashChain\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"name\":\"computeTokenProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"deployWrappedTokenAndRemap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyBridgePauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyBridgeUnpauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"name\":\"getLeafValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProxiedTokensManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenWrappedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWrappedTokenBridgeImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_polygonRollupManager\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_bridgeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sovereignWETHAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_sovereignWETHAddressIsNotMintable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_emergencyBridgePauser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_emergencyBridgeUnpauser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proxiedTokensManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"tokenInfoHash\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_emergencyBridgeUnpauser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proxiedTokensManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"leafIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"sourceBridgeNetwork\",\"type\":\"uint32\"}],\"name\":\"isClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdatedDepositCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tokenInfoHash\",\"type\":\"bytes32\"}],\"name\":\"localBalanceTree\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"legacyTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permitData\",\"type\":\"bytes\"}],\"name\":\"migrateLegacyToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingEmergencyBridgePauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingEmergencyBridgeUnpauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingProxiedTokensManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"polygonRollupManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiedTokensManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"legacySovereignTokenAddress\",\"type\":\"address\"}],\"name\":\"removeLegacySovereignTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeManager\",\"type\":\"address\"}],\"name\":\"setBridgeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"originNetworks\",\"type\":\"uint32[]\"},{\"internalType\":\"address[]\",\"name\":\"originTokenAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"sovereignTokenAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"isNotMintable\",\"type\":\"bool[]\"}],\"name\":\"setMultipleSovereignTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sovereignWETHTokenAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"setSovereignWETHAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenInfoToWrappedToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newEmergencyBridgePauser\",\"type\":\"address\"}],\"name\":\"transferEmergencyBridgePauserRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newEmergencyBridgeUnpauser\",\"type\":\"address\"}],\"name\":\"transferEmergencyBridgeUnpauserRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newProxiedTokensManager\",\"type\":\"address\"}],\"name\":\"transferProxiedTokensManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unsetGlobalIndexHashChain\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"globalIndexes\",\"type\":\"uint256[]\"}],\"name\":\"unsetMultipleClaims\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"verifyMerkleProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrappedAddress\",\"type\":\"address\"}],\"name\":\"wrappedAddressIsNotMintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedTokenBytecodeStorer\",\"outputs\":[{\"internalType\":\"contractIBytecodeStorer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wrappedTokenToTokenInfo\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561000f575f5ffd5b5060405161001c90610146565b604051809103905ff080158015610035573d5f5f3e3d5ffd5b506001600160a01b031660805260405161004e90610153565b604051809103905ff080158015610067573d5f5f3e3d5ffd5b506001600160a01b031660a05261007c610089565b610084610089565b610160565b5f54610100900460ff16156100f45760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161015610144575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b610fd780617eba83390190565b611d3680618e9183390190565b60805160a051617d326101885f395f61065301525f81816105f90152612c5c0152617d325ff3fe6080604052600436106103a7575f3560e01c806383f24403116101e9578063cc46163211610108578063ece93c6f1161009d578063f5efcd791161006d578063f5efcd7914610d36578063f67566e414610d55578063f811bff714610d9d578063fb57083414610dbc575f5ffd5b8063ece93c6f14610ca1578063ee25560b14610ccd578063f0a3d95514610cf8578063f214e16114610d17575f5ffd5b8063d9cb3aec116100d8578063d9cb3aec14610c2f578063dbc1697614610c5a578063e88f043614610c6e578063eabd372a14610c82575f5ffd5b8063cc46163214610ba9578063ccaa2d1114610bc8578063cd58657914610be7578063d02103ca14610bfa575f5ffd5b8063b45869621161017e578063bf130d7f1161014e578063bf130d7f14610b29578063c00f14ab14610b48578063c0f4916314610b67578063c514f24e14610b95575f5ffd5b8063b458696214610a91578063b8b284d014610ab0578063bab161bf14610acf578063be5831c714610af0575f5ffd5b80638d942096116101b95780638d942096146109fb5780638ed7e3f214610a1a578063ae24490a14610a46578063b0b3792014610a72575f5ffd5b806383f24403146109955780638b37b873146109b45780638bd309c3146109c85780638c668f1c146109e7575f5ffd5b80633b2fee9a116102d557806365d6f6541161026a5780636ee84b231161023a5780636ee84b231461091757806379e2cf971461092c5780638129fc1c1461094057806381b1c17414610954575f5ffd5b806365d6f6541461086557806369e3ab12146108ad5780636e4ecfed146108cc5780636e974cd4146108f8575f5ffd5b80634b2f336d116102a55780634b2f336d146107da57806357cfbee3146108065780635ca1e16514610825578063606617ff14610839575f5ffd5b80633b2fee9a146106455780633c351e10146106775780633cbc795b146106a35780633e197043146106ec575f5ffd5b806322e95f2c1161034b5780632f84c6901161031b5780632f84c6901461053a578063318aee3d14610566578063381fef6d146105e857806338b8fbbb1461061b575f5ffd5b806322e95f2c146104d2578063240ff378146104f157806327aef4e8146105045780632dfdf0b514610525575f5ffd5b806314cc01a01161038657806314cc01a01461044657806315064c96146104725780631d081d8c1461049b5780632072f6c5146104be575f5ffd5b80626ee171146103ab57806303e6e116146103cc578063136a2c6014610427575b5f5ffd5b3480156103b6575f5ffd5b506103ca6103c536600461694a565b610ddb565b005b3480156103d7575f5ffd5b5060a8546103fd90610100900473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b348015610432575f5ffd5b506103ca610441366004616a5e565b611697565b348015610451575f5ffd5b5060a3546103fd9073ffffffffffffffffffffffffffffffffffffffff1681565b34801561047d575f5ffd5b5060685461048b9060ff1681565b604051901515815260200161041e565b3480156104a6575f5ffd5b506104b060a55481565b60405190815260200161041e565b3480156104c9575f5ffd5b506103ca611841565b3480156104dd575f5ffd5b506103fd6104ec366004616ae5565b61189c565b6103ca6104ff366004616b5f565b61193e565b34801561050f575f5ffd5b506105186119ed565b60405161041e9190616c20565b348015610530575f5ffd5b506104b060535481565b348015610545575f5ffd5b5060a4546103fd9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610571575f5ffd5b506105b7610580366004616c32565b606b6020525f908152604090205463ffffffff811690640100000000900473ffffffffffffffffffffffffffffffffffffffff1682565b6040805163ffffffff909316835273ffffffffffffffffffffffffffffffffffffffff90911660208301520161041e565b3480156105f3575f5ffd5b506103fd7f000000000000000000000000000000000000000000000000000000000000000081565b348015610626575f5ffd5b5060705473ffffffffffffffffffffffffffffffffffffffff166103fd565b348015610650575f5ffd5b507f00000000000000000000000000000000000000000000000000000000000000006103fd565b348015610682575f5ffd5b50606d546103fd9073ffffffffffffffffffffffffffffffffffffffff1681565b3480156106ae575f5ffd5b50606d546106d79074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff909116815260200161041e565b3480156106f7575f5ffd5b506104b0610706366004616c5b565b6040517fff0000000000000000000000000000000000000000000000000000000000000060f889901b1660208201527fffffffff0000000000000000000000000000000000000000000000000000000060e088811b821660218401527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606089811b821660258601529188901b909216603984015285901b16603d82015260518101839052607181018290525f90609101604051602081830303815290604052805190602001209050979650505050505050565b3480156107e5575f5ffd5b50606f546103fd9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610811575f5ffd5b506103ca610820366004616da7565b611a79565b348015610830575f5ffd5b506104b0611ba4565b348015610844575f5ffd5b5060aa546103fd9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610870575f5ffd5b506105186040518060400160405280600981526020017f616c2d76302e332e31000000000000000000000000000000000000000000000081525081565b3480156108b8575f5ffd5b506103ca6108c7366004616c32565b611c23565b3480156108d7575f5ffd5b506070546103fd9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610903575f5ffd5b506103ca610912366004616ebe565b611d03565b348015610922575f5ffd5b506104b060a65481565b348015610937575f5ffd5b506103ca612207565b34801561094b575f5ffd5b506103ca61223e565b34801561095f575f5ffd5b506103fd61096e366004616f04565b606a6020525f908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b3480156109a0575f5ffd5b506104b06109af366004616f2c565b612270565b3480156109bf575f5ffd5b506103ca6122ff565b3480156109d3575f5ffd5b506103ca6109e2366004616c32565b6123db565b3480156109f2575f5ffd5b506103ca6124ad565b348015610a06575f5ffd5b506103ca610a15366004616c32565b612589565b348015610a25575f5ffd5b50606c546103fd9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610a51575f5ffd5b5060a9546103fd9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610a7d575f5ffd5b506103ca610a8c366004616f68565b61265b565b348015610a9c575f5ffd5b506103ca610aab366004616c32565b612877565b348015610abb575f5ffd5b506103ca610aca366004616fc0565b612aee565b348015610ada575f5ffd5b506068546106d790610100900463ffffffff1681565b348015610afb575f5ffd5b506068546106d790790100000000000000000000000000000000000000000000000000900463ffffffff1681565b348015610b34575f5ffd5b506103ca610b4336600461703e565b612bb8565b348015610b53575f5ffd5b50610518610b62366004616c32565b612c13565b348015610b72575f5ffd5b5061048b610b81366004616c32565b60a26020525f908152604090205460ff1681565b348015610ba0575f5ffd5b50610518612c58565b348015610bb4575f5ffd5b5061048b610bc336600461706a565b612d0c565b348015610bd3575f5ffd5b506103ca610be236600461709b565b612d5d565b6103ca610bf536600461717a565b61330f565b348015610c05575f5ffd5b506068546103fd9065010000000000900473ffffffffffffffffffffffffffffffffffffffff1681565b348015610c3a575f5ffd5b506104b0610c49366004616f04565b60a76020525f908152604090205481565b348015610c65575f5ffd5b506103ca6137ad565b348015610c79575f5ffd5b506103ca613806565b348015610c8d575f5ffd5b506103ca610c9c366004616c32565b61390a565b348015610cac575f5ffd5b506071546103fd9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610cd8575f5ffd5b506104b0610ce7366004616f04565b60696020525f908152604090205481565b348015610d03575f5ffd5b506103ca610d1236600461724b565b613a1b565b348015610d22575f5ffd5b506103fd610d31366004616ae5565b613e42565b348015610d41575f5ffd5b506103ca610d5036600461709b565b613fa3565b348015610d60575f5ffd5b506105186040518060400160405280600781526020017f7631302e312e300000000000000000000000000000000000000000000000000081525081565b348015610da8575f5ffd5b506103ca610db73660046172df565b614327565b348015610dc7575f5ffd5b5061048b610dd636600461736f565b614477565b5f5460ff16607180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000060ff938416021790555f5460039161010090910416158015610e4457505f5460ff8083169116105b610ed5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff80841691909117610100179091556071547401000000000000000000000000000000000000000090041615610f5d576040517ff57ac68300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8a16610faa576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8c63ffffffff165f03610fe9576040517f4e702fa500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8c606860016101000a81548163ffffffff021916908363ffffffff16021790555089606860056101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555088606c5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508660a35f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508360a45f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f85d2bdfbe58cd81abf8199c13ce2509204be4aba8603b9d29f52c4e13e7bb7935f60a45f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660405161117f92919073ffffffffffffffffffffffffffffffffffffffff92831681529116602082015260400190565b60405180910390a160a980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8516908117909155604080515f815260208101929092527f24cc8295aa5110cc216695db944ad2458c7795c6404449be980c3ce14aed752d910160405180910390a13073ffffffffffffffffffffffffffffffffffffffff831603611253576040517f1ae0e03300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166112a0576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8416908117909155604080515f815260208101929092527fa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f2910160405180910390a173ffffffffffffffffffffffffffffffffffffffff8c166113d35763ffffffff8b1615611376576040517f1a874c1200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff86161515806113975750845b156113ce576040517f1cdc46ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6115fb565b606d805463ffffffff8d1674010000000000000000000000000000000000000000027fffffffffffffffff00000000000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff8f1617179055606e6114408982617449565b5073ffffffffffffffffffffffffffffffffffffffff861661158057841515600103611498576040517f1cdc46ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6115365f5f1b601260405160200161152291906060808252600d908201527f5772617070656420457468657200000000000000000000000000000000000000608082015260a0602082018190526004908201527f574554480000000000000000000000000000000000000000000000000000000060c082015260ff91909116604082015260e00190565b60405160208183030381529060405261448e565b606f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff929092169190911790556115fb565b606f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff88169081179091555f90815260a26020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168615151790555b6116036145a2565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050607180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690555050505050505050505050565b606854604080517f91eb796d0000000000000000000000000000000000000000000000000000000081529051339265010000000000900473ffffffffffffffffffffffffffffffffffffffff16916391eb796d9160048083019260209291908290030181865afa15801561170d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117319190617560565b73ffffffffffffffffffffffffffffffffffffffff161461177e576040517fa34ddeb100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b815181101561183d575f82828151811061179c5761179c61757b565b602002602001015190505f5f6801000000000000000083165f146117c2578291506117d9565b602083901c6117d28160016175d5565b9150839250505b6117e38282614640565b60a6545f90815260208490526040902060a68190556040805185815260208101929092527fc80e0aca446a59735359a7ae46124b57c47b892827642779bc6dafc84ba90b03910160405180910390a1505050600101611780565b5050565b60a45473ffffffffffffffffffffffffffffffffffffffff163314611892576040517f26898bbe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61189a6146cb565b565b6040805160e084901b7fffffffff0000000000000000000000000000000000000000000000000000000016602080830191909152606084901b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602483015282516018818403018152603890920183528151918101919091205f908152606a909152205473ffffffffffffffffffffffffffffffffffffffff165b92915050565b60685460ff161561197b576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b34158015906119a15750606f5473ffffffffffffffffffffffffffffffffffffffff1615155b156119d8576040517f6f625c4000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119e685853486868661475d565b5050505050565b606e80546119fa906173b4565b80601f0160208091040260200160405190810160405280929190818152602001828054611a26906173b4565b8015611a715780601f10611a4857610100808354040283529160200191611a71565b820191905f5260205f20905b815481529060010190602001808311611a5457829003601f168201915b505050505081565b60a35473ffffffffffffffffffffffffffffffffffffffff163314611aca576040517faf6e71a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82518451141580611add57508151845114155b80611aea57508051845114155b15611b21576040517f869e93ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b82518110156119e657611b9c858281518110611b4157611b4161757b565b6020026020010151858381518110611b5b57611b5b61757b565b6020026020010151858481518110611b7557611b7561757b565b6020026020010151858581518110611b8f57611b8f61757b565b602002602001015161484a565b600101611b23565b6053545f90819081805b6020811015611c1a578083901c600116600103611bf357611bec60338260208110611bdb57611bdb61757b565b0154855f9182526020526040902090565b9350611c03565b5f84815260208390526040902093505b5f8281526020839052604090209150600101611bae565b50919392505050565b60a45473ffffffffffffffffffffffffffffffffffffffff163314611c74576040517f26898bbe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60a880547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1661010073ffffffffffffffffffffffffffffffffffffffff8481169182029290921790925560a4546040805191909216815260208101929092527fb27de219766f47b82684842855ba6130b6dbf288ac66d1c3509e7bf17f4e925a91015b60405180910390a150565b60a35473ffffffffffffffffffffffffffffffffffffffff163314611d54576040517faf6e71a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216158015611d7c575063ffffffff8316155b15611f96575f611f845f5f1b606f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166306fdde036040518163ffffffff1660e01b81526004015f60405180830381865afa158015611df1573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052611e369190810190617645565b606f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166395d89b416040518163ffffffff1660e01b81526004015f60405180830381865afa158015611e9f573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052611ee49190810190617645565b606f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015611f4e573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611f729190617677565b60405160200161152293929190617692565b9050611f908183614b33565b50505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1660208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084901b1660248201525f90603801604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301205f818152606a90935291205490915073ffffffffffffffffffffffffffffffffffffffff1680612088576040517f828d566300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6121f0838373ffffffffffffffffffffffffffffffffffffffff166306fdde036040518163ffffffff1660e01b81526004015f60405180830381865afa1580156120d5573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820160405261211a9190810190617645565b8473ffffffffffffffffffffffffffffffffffffffff166395d89b416040518163ffffffff1660e01b81526004015f60405180830381865afa158015612162573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526121a79190810190617645565b8573ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015611f4e573d5f5f3e3d5ffd5b90506121fe8686838761484a565b5050505b505050565b605354606854790100000000000000000000000000000000000000000000000000900463ffffffff16101561189a5761189a614c3c565b6040517ff57ac68300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f83815b60208110156122f457600163ffffffff8516821c811690036122c0576122b98582602081106122a5576122a561757b565b6020020135835f9182526020526040902090565b91506122ec565b6122e9828683602081106122d6576122d661757b565b60200201355f9182526020526040902090565b91505b600101612274565b5090505b9392505050565b60aa5473ffffffffffffffffffffffffffffffffffffffff163314612350576040517fd491f0c100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60a9805460aa805473ffffffffffffffffffffffffffffffffffffffff8082167fffffffffffffffffffffffff0000000000000000000000000000000000000000808616821790965594909116909155604080519190921680825260208201939093527f85d2bdfbe58cd81abf8199c13ce2509204be4aba8603b9d29f52c4e13e7bb7939101611cf8565b60705473ffffffffffffffffffffffffffffffffffffffff16331461242c576040517f0866750300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169182179092556070546040805191909316815260208101919091527f0a34baa3feb299aef9c05cb59c6e0c8e7c0bcc65cbf0a647e7a7c8a2411591e29101611cf8565b60715473ffffffffffffffffffffffffffffffffffffffff1633146124fe576040517f2d67bc9c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607080546071805473ffffffffffffffffffffffffffffffffffffffff8082167fffffffffffffffffffffffff0000000000000000000000000000000000000000808616821790965594909116909155604080519190921680825260208201939093527fa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f29101611cf8565b60a95473ffffffffffffffffffffffffffffffffffffffff1633146125da576040517f8e9d821f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60aa80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560a9546040805191909316815260208101919091527ff01a62a06940517bbc898dec8c75794b9feabcd2d263c8de823b36dbbeb8779b9101611cf8565b801561266c5761266c848383614d07565b73ffffffffffffffffffffffffffffffffffffffff8085165f908152606b602090815260409182902082518084019093525463ffffffff811683526401000000009004909216918101829052906126ef576040517f828d566300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f606a5f835f0151846020015160405160200161276292919060e09290921b7fffffffff0000000000000000000000000000000000000000000000000000000016825260601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016600482015260180190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181528151602092830120835290820192909252015f205473ffffffffffffffffffffffffffffffffffffffff9081169150861681036127f9576040517fe273c4a100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61280487876151c4565b90506128118233836153cf565b6040805133815273ffffffffffffffffffffffffffffffffffffffff89811660208301528416818301526060810183905290517fb7f8fd4d1faf9b2929dc269f59c53e3a2bccc44e9950f33a568fcbcb37eb69a99181900360800190a150505050505050565b60a35473ffffffffffffffffffffffffffffffffffffffff1633146128c8576040517faf6e71a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8082165f908152606b6020908152604080832081518083018352905463ffffffff81168083526401000000009091049095168184018190529151909461297b939092910160e09290921b7fffffffff0000000000000000000000000000000000000000000000000000000016825260601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016600482015260180190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301205f818152606a90935291205490915073ffffffffffffffffffffffffffffffffffffffff161580612a0757505f818152606a602052604090205473ffffffffffffffffffffffffffffffffffffffff8481169116145b15612a3e576040517fe0c897a700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83165f818152606b6020908152604080832080547fffffffffffffffff00000000000000000000000000000000000000000000000016905560a282529182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905590519182527fc2ae0bd0ec0fd0352bfe5bacac49637af342c1e40f1b80a7f74440dc7fe3f063910160405180910390a1505050565b60685460ff1615612b2b576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5473ffffffffffffffffffffffffffffffffffffffff16612b7a576040517fdde3cda700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f545f90612b9f9073ffffffffffffffffffffffffffffffffffffffff16866151c4565b9050612baf87878387878761475d565b50505050505050565b60a35473ffffffffffffffffffffffffffffffffffffffff163314612c09576040517faf6e71a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61183d8282614b33565b6060612c1e826154e9565b612c27836155fd565b612c3084615700565b604051602001612c4293929190617692565b6040516020818303038152906040529050919050565b60607f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c514f24e6040518163ffffffff1660e01b81526004015f60405180830381865afa158015612cc2573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052612d0791908101906176ca565b905090565b5f80612d2364010000000063ffffffff851661770f565b612d339063ffffffff8616617726565b600881901c5f90815260696020526040902054600160ff9092169190911b90811614949350505050565b60685460ff1615612d9a576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612da26157ef565b60685463ffffffff8681166101009092041614612deb576040517f0595ea2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612e168c8c8c8c8c5f8d8d8d8d8d8d8d604051612e09929190617739565b6040518091039020615862565b604080518b815263ffffffff8916602082015273ffffffffffffffffffffffffffffffffffffffff88811682840152861660608201526080810185905290517f1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d9181900360a00190a173ffffffffffffffffffffffffffffffffffffffff8616158015612ea7575063ffffffff8716155b15612fc557606f5473ffffffffffffffffffffffffffffffffffffffff16612f9c575f73ffffffffffffffffffffffffffffffffffffffff851684825b6040519080825280601f01601f191660200182016040528015612f0e576020820181803683370190505b50604051612f1c9190617748565b5f6040518083038185875af1925050503d805f8114612f56576040519150601f19603f3d011682016040523d82523d5f602084013e612f5b565b606091505b5050905080612f96576040517f6747a28800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506132f8565b606f54612fc09073ffffffffffffffffffffffffffffffffffffffff1685856153cf565b6132f8565b606d5473ffffffffffffffffffffffffffffffffffffffff87811691161480156130115750606d5463ffffffff8881167401000000000000000000000000000000000000000090920416145b15613035575f73ffffffffffffffffffffffffffffffffffffffff85168482612ee4565b60685463ffffffff61010090910481169088160361306e57612fc073ffffffffffffffffffffffffffffffffffffffff871685856159e8565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e089901b1660208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606088901b1660248201525f90603801604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301205f818152606a90935291205490915073ffffffffffffffffffffffffffffffffffffffff16806132ea575f61316e8386868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061448e92505050565b905061317b8188886153cf565b80606a5f8581526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060405180604001604052808b63ffffffff1681526020018a73ffffffffffffffffffffffffffffffffffffffff16815250606b5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f015f6101000a81548163ffffffff021916908363ffffffff1602179055506020820151815f0160046101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050507f490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a398a8a8388886040516132dc9594939291906177a5565b60405180910390a1506132f5565b6132f58187876153cf565b50505b61330160018055565b505050505050505050505050565b60685460ff161561334c576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6133546157ef565b60685463ffffffff61010090910481169088160361339e576040517f0595ea2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8060608773ffffffffffffffffffffffffffffffffffffffff88166134c5578834146133f7576040517fb89240f500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606d54606e805473ffffffffffffffffffffffffffffffffffffffff831696507401000000000000000000000000000000000000000090920463ffffffff16945090613442906173b4565b80601f016020809104026020016040519081016040528092919081815260200182805461346e906173b4565b80156134b95780601f10613490576101008083540402835291602001916134b9565b820191905f5260205f20905b81548152906001019060200180831161349c57829003601f168201915b50505050509150613735565b34156134fd576040517f798ee6f100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b841561350e5761350e888787614d07565b606f5473ffffffffffffffffffffffffffffffffffffffff908116908916036135425761353b888a6151c4565b9050613735565b73ffffffffffffffffffffffffffffffffffffffff8089165f908152606b602090815260409182902082518084019093525463ffffffff811683526401000000009004909216918101829052901515806135a25750805163ffffffff1615155b156135c4576135b1898b6151c4565b6020820151825190965094509150613728565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f9073ffffffffffffffffffffffffffffffffffffffff8b16906370a0823190602401602060405180830381865afa15801561362e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906136529190617807565b905061367673ffffffffffffffffffffffffffffffffffffffff8b1633308e615a6f565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f9073ffffffffffffffffffffffffffffffffffffffff8c16906370a0823190602401602060405180830381865afa1580156136e0573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906137049190617807565b9050613710828261781e565b6068548c9850610100900463ffffffff169650935050505b61373189612c13565b9250505b7f501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b5f84868e8e8688605354604051613774989796959493929190617831565b60405180910390a16137925f84868e8e868880519060200120615ab5565b86156137a0576137a0614c3c565b50505050612baf60018055565b60a95473ffffffffffffffffffffffffffffffffffffffff1633146137fe576040517f8e9d821f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61189a615b0c565b60a854610100900473ffffffffffffffffffffffffffffffffffffffff16331461385c576040517f7bb0100f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60a4805460a8805473ffffffffffffffffffffffffffffffffffffffff610100820481167fffffffffffffffffffffffff0000000000000000000000000000000000000000851681179095557fffffffffffffffffffffff0000000000000000000000000000000000000000ff909116909155604080519190921680825260208201939093527f85d2bdfbe58cd81abf8199c13ce2509204be4aba8603b9d29f52c4e13e7bb7939101611cf8565b60a35473ffffffffffffffffffffffffffffffffffffffff16331461395b576040517faf6e71a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81166139a8576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60a380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f32cf74f8a6d5f88593984d2cd52be5592bfa6884f5896175801a5069ef09cd6790602001611cf8565b5f5460ff16607180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000060ff938416021790555f5460039161010090910416158015613a8457505f5460ff8083169116105b613b10576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610ecc565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff80841691909117610100178255607154740100000000000000000000000000000000000000009004169003613b98576040517ff57ac68300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b858414613bd1576040517f869e93ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b86811015613c1d57613c15888883818110613bf057613bf061757b565b90506020020135878784818110613c0957613c0961757b565b90506020020135615b9a565b600101613bd3565b5060a980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8516908117909155604080515f815260208101929092527f24cc8295aa5110cc216695db944ad2458c7795c6404449be980c3ce14aed752d910160405180910390a13073ffffffffffffffffffffffffffffffffffffffff831603613cea576040517f1ae0e03300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216613d37576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8416908117909155604080515f815260208101929092527fa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f2910160405180910390a15f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050607180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690555050505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1660208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606083901b1660248201525f9081906038016040516020818303038152906040528051906020012090505f60ff60f81b3083613ecc612c58565b604051602001613edc9190617748565b60405160208183030381529060405280519060200120604051602001613f6494939291907fff0000000000000000000000000000000000000000000000000000000000000094909416845260609290921b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660018401526015830152603582015260550190565b604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018152919052805160209091012095945050505050565b60685460ff1615613fe0576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60685463ffffffff8681166101009092041614614029576040517f0595ea2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6140488c8c8c8c8c60018d8d8d8d8d8d8d604051612e09929190617739565b604080518b815263ffffffff8916602082015273ffffffffffffffffffffffffffffffffffffffff88811682840152861660608201526080810185905290517f1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d9181900360a00190a1606f545f9073ffffffffffffffffffffffffffffffffffffffff166141ca578473ffffffffffffffffffffffffffffffffffffffff1684888a86866040516024016140ff94939291906178bf565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f1806b5f200000000000000000000000000000000000000000000000000000000179052516141809190617748565b5f6040518083038185875af1925050503d805f81146141ba576040519150601f19603f3d011682016040523d82523d5f602084013e6141bf565b606091505b5050809150506142e1565b606f546141ee9073ffffffffffffffffffffffffffffffffffffffff1686866153cf565b8473ffffffffffffffffffffffffffffffffffffffff168789858560405160240161421c94939291906178bf565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f1806b5f2000000000000000000000000000000000000000000000000000000001790525161429d9190617748565b5f604051808303815f865af19150503d805f81146142d6576040519150601f19603f3d011682016040523d82523d5f602084013e6142db565b606091505b50909150505b80614318576040517f37e391c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50505050505050505050505050565b5f54610100900460ff161580801561434557505f54600160ff909116105b8061435e5750303b15801561435e57505f5460ff166001145b6143ea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610ecc565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561223e575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790556040517ff57ac68300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f81614484868686612270565b1495945050505050565b5f5f614498612c58565b6040516020016144a89190617748565b6040516020818303038152906040529050838151602083015ff5915073ffffffffffffffffffffffffffffffffffffffff8216614511576040517f62d05d1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f5f8580602001905181019061452891906178fa565b9250925092508473ffffffffffffffffffffffffffffffffffffffff16631624f6c68484846040518463ffffffff1660e01b815260040161456b93929190617692565b5f604051808303815f87803b158015614582575f5ffd5b505af1158015614594573d5f5f3e3d5ffd5b505050505050505092915050565b5f54610100900460ff16614638576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610ecc565b61189a615be1565b5f61465664010000000063ffffffff841661770f565b6146669063ffffffff8516617726565b600881901c5f8181526069602052604090208054600160ff851690811b918218928390559394509192919080821615612baf576040517f318dafb800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60685460ff1615614708576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606880547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556040517f2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497905f90a1565b60685463ffffffff6101009091048116908716036147a7576040517f0595ea2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b6001606860019054906101000a900463ffffffff163389898988886053546040516147fb99989796959493929190617967565b60405180910390a161483c6001606860019054906101000a900463ffffffff1633898989888860405161482f929190617739565b6040518091039020615ab5565b82156121fe576121fe614c3c565b73ffffffffffffffffffffffffffffffffffffffff83161580614881575073ffffffffffffffffffffffffffffffffffffffff8216155b156148b8576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60685463ffffffff610100909104811690851603614902576040517f658b23ad00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8281165f908152606b602052604090205464010000000090041615614968576040517f5eaf7bac00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b1660208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606085901b1660248201525f90603801604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001815282825280516020918201205f818152606a835283812080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8a8116918217909255868601865263ffffffff8c81168089528c8416878a01818152848752606b89528987209a518b54915194167fffffffffffffffff0000000000000000000000000000000000000000000000009091161764010000000093909516929092029390931790975560a285529185902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001689151590811790915585519182529381019590955292840192909252606083015291507fdbe8a5da6a7a916d9adfda9160167a0f8a3da415ee6610e810e753853597fce79060800160405180910390a15050505050565b606d5473ffffffffffffffffffffffffffffffffffffffff16614b82576040517f9968e22600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84169081179091555f81815260a2602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168515159081179091558251938452908301527fc7318b7ed6ba4f2908a3de396d8ab49b1dadb55db5b55123247a401f29ff8d8291015b60405180910390a15050565b6053546068805463ffffffff909216790100000000000000000000000000000000000000000000000000027fffffff00000000ffffffffffffffffffffffffffffffffffffffffffffffffff909216919091179081905573ffffffffffffffffffffffffffffffffffffffff65010000000000909104166333d6247d614cc0611ba4565b6040518263ffffffff1660e01b8152600401614cde91815260200190565b5f604051808303815f87803b158015614cf5575f5ffd5b505af1158015611f90573d5f5f3e3d5ffd5b5f614d1560048284866179f7565b614d1e91617a1e565b90507f2afa5331000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000821601614f52575f808080808080614d7d896004818d6179f7565b810190614d8a9190617a84565b96509650965096509650965096503373ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff1614614dfd576040517f912ecce700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff86163014614e4c576040517f750643af00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805173ffffffffffffffffffffffffffffffffffffffff89811660248301528881166044830152606482018890526084820187905260ff861660a483015260c4820185905260e48083018590528351808403909101815261010490920183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd505accf000000000000000000000000000000000000000000000000000000001790529151918d1691614f059190617748565b5f604051808303815f865af19150503d805f8114614f3e576040519150601f19603f3d011682016040523d82523d5f602084013e614f43565b606091505b50505050505050505050611f90565b7fffffffff0000000000000000000000000000000000000000000000000000000081167f8fcbaf0c0000000000000000000000000000000000000000000000000000000014614fcd576040517fe282c0ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80808080808080614fe28a6004818e6179f7565b810190614fef9190617ad3565b975097509750975097509750975097503373ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff1614615064576040517f912ecce700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff871630146150b3576040517f750643af00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805173ffffffffffffffffffffffffffffffffffffffff8a811660248301528981166044830152606482018990526084820188905286151560a483015260ff861660c483015260e482018590526101048083018590528351808403909101815261012490920183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f8fcbaf0c000000000000000000000000000000000000000000000000000000001790529151918e16916151759190617748565b5f604051808303815f865af19150503d805f81146151ae576040519150601f19603f3d011682016040523d82523d5f602084013e6151b3565b606091505b505050505050505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff82165f90815260a2602052604081205460ff1615615346576040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f9073ffffffffffffffffffffffffffffffffffffffff8516906370a0823190602401602060405180830381865afa15801561525b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061527f9190617807565b90506152a373ffffffffffffffffffffffffffffffffffffffff8516333086615a6f565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f9073ffffffffffffffffffffffffffffffffffffffff8616906370a0823190602401602060405180830381865afa15801561530d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906153319190617807565b905061533d828261781e565b92505050611938565b6040517f9dc29fac0000000000000000000000000000000000000000000000000000000081523360048201526024810183905273ffffffffffffffffffffffffffffffffffffffff841690639dc29fac906044015f604051808303815f87803b1580156153b1575f5ffd5b505af11580156153c3573d5f5f3e3d5ffd5b50505050819050611938565b73ffffffffffffffffffffffffffffffffffffffff821661541c576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83165f90815260a2602052604090205460ff161561546a5761220273ffffffffffffffffffffffffffffffffffffffff841683836159e8565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8381166004830152602482018390528416906340c10f19906044015f604051808303815f87803b1580156154d7575f5ffd5b505af1158015612baf573d5f5f3e3d5ffd5b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f06fdde030000000000000000000000000000000000000000000000000000000017905290516060915f91829173ffffffffffffffffffffffffffffffffffffffff86169161556a9190617748565b5f60405180830381855afa9150503d805f81146155a2576040519150601f19603f3d011682016040523d82523d5f602084013e6155a7565b606091505b5091509150816155ec576040518060400160405280600781526020017f4e4f5f4e414d45000000000000000000000000000000000000000000000000008152506155f5565b6155f581615c77565b949350505050565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f95d89b410000000000000000000000000000000000000000000000000000000017905290516060915f91829173ffffffffffffffffffffffffffffffffffffffff86169161567e9190617748565b5f60405180830381855afa9150503d805f81146156b6576040519150601f19603f3d011682016040523d82523d5f602084013e6156bb565b606091505b5091509150816155ec576040518060400160405280600981526020017f4e4f5f53594d424f4c00000000000000000000000000000000000000000000008152506155f5565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f313ce5670000000000000000000000000000000000000000000000000000000017905290515f918291829173ffffffffffffffffffffffffffffffffffffffff8616916157809190617748565b5f60405180830381855afa9150503d805f81146157b8576040519150601f19603f3d011682016040523d82523d5f602084013e6157bd565b606091505b50915091508180156157d0575080516020145b6157db5760126155f5565b808060200190518101906155f59190617677565b60026001540361585b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610ecc565b6002600155565b604080517fff0000000000000000000000000000000000000000000000000000000000000060f88a901b166020808301919091527fffffffff0000000000000000000000000000000000000000000000000000000060e08a811b821660218501527fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060608b811b82166025870152918a901b909216603985015287901b16603d83015260518201859052607180830185905283518084039091018152609190920190925280519101206159398d8d8d8d8d86615e3d565b60a554615961906159538d845f9182526020526040902090565b5f9182526020526040902090565b60a5819055604080518d815260208101929092527f3e5936f910a78eb5181813a939c8d4c3e4d85f87943f659380d82ac6221b0e92910160405180910390a160ff88166159b3576159b387878561608d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60ff891601614318576143185f5f8561608d565b60405173ffffffffffffffffffffffffffffffffffffffff83811660248301526044820183905261220291859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061621c565b60018055565b60405173ffffffffffffffffffffffffffffffffffffffff8481166024830152838116604483015260648201839052611f909186918216906323b872dd90608401615a22565b615ac4878787878787876162b0565b60ff8716615ad757615ad7868684616382565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60ff881601612baf57612baf5f5f84616382565b60685460ff16615b48576040517f5386698100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606880547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556040517f1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3905f90a1565b5f82815260a7602090815260409182902083905581518481529081018390527f2277ec68451dc01bd131765a9858d6de94d7e11220704d8ac1718fdb8de07cb29101614c30565b5f54610100900460ff16615a69576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610ecc565b60606040825110615c9657818060200190518101906119389190617645565b8151602003615dff575f5b602081108015615ce85750828181518110615cbe57615cbe61757b565b01602001517fff000000000000000000000000000000000000000000000000000000000000001615155b15615cff5780615cf781617b51565b915050615ca1565b805f03615d4157505060408051808201909152601281527f4e4f545f56414c49445f454e434f44494e4700000000000000000000000000006020820152919050565b5f8167ffffffffffffffff811115615d5b57615d5b61681f565b6040519080825280601f01601f191660200182016040528015615d85576020820181803683370190505b5090505f5b82811015615df757848181518110615da457615da461757b565b602001015160f81c60f81b828281518110615dc157615dc161757b565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600101615d8a565b509392505050565b505060408051808201909152601281527f4e4f545f56414c49445f454e434f44494e470000000000000000000000000000602082015290565b919050565b6068545f9065010000000000900473ffffffffffffffffffffffffffffffffffffffff1663257b3632615e7986865f9182526020526040902090565b6040518263ffffffff1660e01b8152600401615e9791815260200190565b6020604051808303815f875af1158015615eb3573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190615ed79190617807565b9050805f03615f11576040517e2f6fad00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8068010000000000000000871615615fbf5786915081615f4163ffffffff821668010000000000000000617726565b14615f78576040517f071389e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b615f84848a8489614477565b615fba576040517fe0417cec00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b616078565b602087901c615fcf8160016175d5565b889350915082615ff363ffffffff821667ffffffff00000000602085901b16617726565b1461602a576040517f071389e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b616040616038868c86612270565b8a8389614477565b616076576040517fe0417cec00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b61608282826164dc565b505050505050505050565b60685463ffffffff6101009091048116908416036160aa57505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1660208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084901b1660248201525f90603801604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301205f81815260a7909352912054909150616179907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61781e565b8211156161f4575f81815260a76020526040908190205490517f23d7213300000000000000000000000000000000000000000000000000000000815263ffffffff8616600482015273ffffffffffffffffffffffffffffffffffffffff85166024820152604481018490526064810191909152608401610ecc565b5f81815260a7602052604081208054849290616211908490617726565b909155505050505050565b5f61623d73ffffffffffffffffffffffffffffffffffffffff841683616568565b905080515f1415801561626157508080602001905181019061625f9190617b88565b155b15612202576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610ecc565b604080517fff0000000000000000000000000000000000000000000000000000000000000060f88a901b166020808301919091527fffffffff0000000000000000000000000000000000000000000000000000000060e08a811b821660218501527fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060608b811b82166025870152918a901b909216603985015287901b16603d8301526051820185905260718083018590528351808403909101815260919092019092528051910120612baf90616575565b60685463ffffffff61010090910481169084160361639f57505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1660208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084901b1660248201525f90603801604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301205f81815260a79093529120549091508211156164bf575f81815260a76020526040908190205490517f14603c0100000000000000000000000000000000000000000000000000000000815263ffffffff8616600482015273ffffffffffffffffffffffffffffffffffffffff85166024820152604481018490526064810191909152608401610ecc565b5f81815260a760205260408120805484929061621190849061781e565b5f6164f264010000000063ffffffff841661770f565b6165029063ffffffff8516617726565b600881901c5f8181526069602052604081208054600160ff861690811b91821892839055949550929392918183169003612baf576040517f646cf55800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60606122f883835f61664d565b80600161658460206002617cc4565b61658e919061781e565b605354106165c8576040517fef5ccf6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60535f81546165d790617b51565b918290555090505f5b6020811015616644578082901c60011660010361661357826033826020811061660b5761660b61757b565b015550505050565b61663a603382602081106166295761662961757b565b0154845f9182526020526040902090565b92506001016165e0565b50612202617ccf565b60608147101561668b576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401610ecc565b5f5f8573ffffffffffffffffffffffffffffffffffffffff1684866040516166b39190617748565b5f6040518083038185875af1925050503d805f81146166ed576040519150601f19603f3d011682016040523d82523d5f602084013e6166f2565b606091505b509150915061670286838361670c565b9695505050505050565b6060826167215761671c8261679b565b6122f8565b8151158015616745575073ffffffffffffffffffffffffffffffffffffffff84163b155b15616794576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610ecc565b50806122f8565b8051156167ab5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b803563ffffffff81168114615e38575f5ffd5b73ffffffffffffffffffffffffffffffffffffffff811681146167dd575f5ffd5b8035615e38816167f3565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156168935761689361681f565b604052919050565b5f67ffffffffffffffff8211156168b4576168b461681f565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b5f82601f8301126168ef575f5ffd5b81356169026168fd8261689b565b61684c565b818152846020838601011115616916575f5ffd5b816020850160208301375f918101602001919091529392505050565b80151581146167dd575f5ffd5b8035615e3881616932565b5f5f5f5f5f5f5f5f5f5f5f5f6101808d8f031215616966575f5ffd5b61696f8d6167e0565b9b5061697d60208e01616814565b9a5061698b60408e016167e0565b995061699960608e01616814565b98506169a760808e01616814565b975067ffffffffffffffff60a08e013511156169c1575f5ffd5b6169d18e60a08f01358f016168e0565b96506169df60c08e01616814565b95506169ed60e08e01616814565b94506169fc6101008e0161693f565b9350616a0b6101208e01616814565b9250616a1a6101408e01616814565b9150616a296101608e01616814565b90509295989b509295989b509295989b565b5f67ffffffffffffffff821115616a5457616a5461681f565b5060051b60200190565b5f60208284031215616a6e575f5ffd5b813567ffffffffffffffff811115616a84575f5ffd5b8201601f81018413616a94575f5ffd5b8035616aa26168fd82616a3b565b8082825260208201915060208360051b850101925086831115616ac3575f5ffd5b6020840193505b82841015616702578335825260209384019390910190616aca565b5f5f60408385031215616af6575f5ffd5b616aff836167e0565b91506020830135616b0f816167f3565b809150509250929050565b5f5f83601f840112616b2a575f5ffd5b50813567ffffffffffffffff811115616b41575f5ffd5b602083019150836020828501011115616b58575f5ffd5b9250929050565b5f5f5f5f5f60808688031215616b73575f5ffd5b616b7c866167e0565b94506020860135616b8c816167f3565b93506040860135616b9c81616932565b9250606086013567ffffffffffffffff811115616bb7575f5ffd5b616bc388828901616b1a565b969995985093965092949392505050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f6122f86020830184616bd4565b5f60208284031215616c42575f5ffd5b81356122f8816167f3565b60ff811681146167dd575f5ffd5b5f5f5f5f5f5f5f60e0888a031215616c71575f5ffd5b8735616c7c81616c4d565b9650616c8a602089016167e0565b95506040880135616c9a816167f3565b9450616ca8606089016167e0565b93506080880135616cb8816167f3565b9699959850939692959460a0840135945060c09093013592915050565b5f82601f830112616ce4575f5ffd5b8135616cf26168fd82616a3b565b8082825260208201915060208360051b860101925085831115616d13575f5ffd5b602085015b83811015616d39578035616d2b816167f3565b835260209283019201616d18565b5095945050505050565b5f82601f830112616d52575f5ffd5b8135616d606168fd82616a3b565b8082825260208201915060208360051b860101925085831115616d81575f5ffd5b602085015b83811015616d39578035616d9981616932565b835260209283019201616d86565b5f5f5f5f60808587031215616dba575f5ffd5b843567ffffffffffffffff811115616dd0575f5ffd5b8501601f81018713616de0575f5ffd5b8035616dee6168fd82616a3b565b8082825260208201915060208360051b850101925089831115616e0f575f5ffd5b6020840193505b82841015616e3857616e27846167e0565b825260209384019390910190616e16565b9650505050602085013567ffffffffffffffff811115616e56575f5ffd5b616e6287828801616cd5565b935050604085013567ffffffffffffffff811115616e7e575f5ffd5b616e8a87828801616cd5565b925050606085013567ffffffffffffffff811115616ea6575f5ffd5b616eb287828801616d43565b91505092959194509250565b5f5f5f60608486031215616ed0575f5ffd5b616ed9846167e0565b92506020840135616ee9816167f3565b91506040840135616ef981616932565b809150509250925092565b5f60208284031215616f14575f5ffd5b5035919050565b806104008101831015611938575f5ffd5b5f5f5f6104408486031215616f3f575f5ffd5b83359250616f508560208601616f1b565b9150616f5f61042085016167e0565b90509250925092565b5f5f5f5f60608587031215616f7b575f5ffd5b8435616f86816167f3565b935060208501359250604085013567ffffffffffffffff811115616fa8575f5ffd5b616fb487828801616b1a565b95989497509550505050565b5f5f5f5f5f5f60a08789031215616fd5575f5ffd5b616fde876167e0565b95506020870135616fee816167f3565b945060408701359350606087013561700581616932565b9250608087013567ffffffffffffffff811115617020575f5ffd5b61702c89828a01616b1a565b979a9699509497509295939492505050565b5f5f6040838503121561704f575f5ffd5b823561705a816167f3565b91506020830135616b0f81616932565b5f5f6040838503121561707b575f5ffd5b617084836167e0565b9150617092602084016167e0565b90509250929050565b5f5f5f5f5f5f5f5f5f5f5f5f6109208d8f0312156170b7575f5ffd5b6170c18e8e616f1b565b9b506170d18e6104008f01616f1b565b9a506108008d013599506108208d013598506108408d013597506170f86108608e016167e0565b96506171086108808e01356167f3565b6108808d0135955061711d6108a08e016167e0565b94506108c08d013561712e816167f3565b93506108e08d0135925067ffffffffffffffff6109008e01351115617151575f5ffd5b6171628e6109008f01358f01616b1a565b81935080925050509295989b509295989b509295989b565b5f5f5f5f5f5f5f60c0888a031215617190575f5ffd5b617199886167e0565b965060208801356171a9816167f3565b95506040880135945060608801356171c0816167f3565b935060808801356171d081616932565b925060a088013567ffffffffffffffff8111156171eb575f5ffd5b6171f78a828b01616b1a565b989b979a50959850939692959293505050565b5f5f83601f84011261721a575f5ffd5b50813567ffffffffffffffff811115617231575f5ffd5b6020830191508360208260051b8501011115616b58575f5ffd5b5f5f5f5f5f5f60808789031215617260575f5ffd5b863567ffffffffffffffff811115617276575f5ffd5b61728289828a0161720a565b909750955050602087013567ffffffffffffffff8111156172a1575f5ffd5b6172ad89828a0161720a565b90955093505060408701356172c1816167f3565b915060608701356172d1816167f3565b809150509295509295509295565b5f5f5f5f5f5f60c087890312156172f4575f5ffd5b6172fd876167e0565b9550602087013561730d816167f3565b945061731b604088016167e0565b9350606087013561732b816167f3565b9250608087013561733b816167f3565b915060a087013567ffffffffffffffff811115617356575f5ffd5b61736289828a016168e0565b9150509295509295509295565b5f5f5f5f6104608587031215617383575f5ffd5b843593506173948660208701616f1b565b92506173a361042086016167e0565b939692955092936104400135925050565b600181811c908216806173c857607f821691505b6020821081036173ff577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b601f82111561220257805f5260205f20601f840160051c8101602085101561742a5750805b601f840160051c820191505b818110156119e6575f8155600101617436565b815167ffffffffffffffff8111156174635761746361681f565b6174778161747184546173b4565b84617405565b6020601f8211600181146174c8575f83156174925750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b1784556119e6565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b8281101561751557878501518255602094850194600190920191016174f5565b508482101561755157868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b5f60208284031215617570575f5ffd5b81516122f8816167f3565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b63ffffffff8181168382160190811115611938576119386175a8565b5f6175fe6168fd8461689b565b9050828152838383011115617611575f5ffd5b8282602083015e5f602084830101529392505050565b5f82601f830112617636575f5ffd5b6122f8838351602085016175f1565b5f60208284031215617655575f5ffd5b815167ffffffffffffffff81111561766b575f5ffd5b6155f584828501617627565b5f60208284031215617687575f5ffd5b81516122f881616c4d565b606081525f6176a46060830186616bd4565b82810360208401526176b68186616bd4565b91505060ff83166040830152949350505050565b5f602082840312156176da575f5ffd5b815167ffffffffffffffff8111156176f0575f5ffd5b8201601f81018413617700575f5ffd5b6155f5848251602084016175f1565b8082028115828204841417611938576119386175a8565b80820180821115611938576119386175a8565b818382375f9101908152919050565b5f82518060208501845e5f920191825250919050565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b63ffffffff8616815273ffffffffffffffffffffffffffffffffffffffff8516602082015273ffffffffffffffffffffffffffffffffffffffff84166040820152608060608201525f6177fc60808301848661775e565b979650505050505050565b5f60208284031215617817575f5ffd5b5051919050565b81810381811115611938576119386175a8565b60ff8916815263ffffffff8816602082015273ffffffffffffffffffffffffffffffffffffffff8716604082015263ffffffff8616606082015273ffffffffffffffffffffffffffffffffffffffff851660808201528360a082015261010060c08201525f6178a4610100830185616bd4565b905063ffffffff831660e08301529998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8516815263ffffffff84166020820152606060408201525f61670260608301848661775e565b5f5f5f6060848603121561790c575f5ffd5b835167ffffffffffffffff811115617922575f5ffd5b61792e86828701617627565b935050602084015167ffffffffffffffff81111561794a575f5ffd5b61795686828701617627565b9250506040840151616ef981616c4d565b60ff8a16815263ffffffff8916602082015273ffffffffffffffffffffffffffffffffffffffff8816604082015263ffffffff8716606082015273ffffffffffffffffffffffffffffffffffffffff861660808201528460a082015261010060c08201525f6179db6101008301858761775e565b905063ffffffff831660e08301529a9950505050505050505050565b5f5f85851115617a05575f5ffd5b83861115617a11575f5ffd5b5050820193919092039150565b80357fffffffff000000000000000000000000000000000000000000000000000000008116906004841015617a7d577fffffffff00000000000000000000000000000000000000000000000000000000808560040360031b1b82161691505b5092915050565b5f5f5f5f5f5f5f60e0888a031215617a9a575f5ffd5b8735617aa5816167f3565b96506020880135617ab5816167f3565b955060408801359450606088013593506080880135616cb881616c4d565b5f5f5f5f5f5f5f5f610100898b031215617aeb575f5ffd5b8835617af6816167f3565b97506020890135617b06816167f3565b965060408901359550606089013594506080890135617b2481616932565b935060a0890135617b3481616c4d565b979a969950949793969295929450505060c08201359160e0013590565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203617b8157617b816175a8565b5060010190565b5f60208284031215617b98575f5ffd5b81516122f881616932565b6001815b6001841115617bde57808504811115617bc257617bc26175a8565b6001841615617bd057908102905b60019390931c928002617ba7565b935093915050565b5f82617bf457506001611938565b81617c0057505f611938565b8160018114617c165760028114617c2057617c3c565b6001915050611938565b60ff841115617c3157617c316175a8565b50506001821b611938565b5060208310610133831016604e8410600b8410161715617c5f575081810a611938565b617c8a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484617ba3565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115617cbc57617cbc6175a8565b029392505050565b5f6122f88383617be6565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52600160045260245ffdfea26469706673582212200756c05cb9a8369a9f3b85bbb1fabaa09d3f280f6876557e3254b0e2f3873ab964736f6c634300081c00336080604052348015600e575f5ffd5b50610fbb8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610029575f3560e01c8063c514f24e1461002d575b5f5ffd5b61003561004b565b604051610042919061006a565b60405180910390f35b60405180610f000160405280610ec881526020016100be610ec8913981565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168401019150509291505056fe60806040819052631d97f74d60e11b81523390633b2fee9a90608490602090600481865afa158015610033573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100579190610433565b604080515f80825260208201909252905061007382825f6100e2565b50506100dd336001600160a01b03166338b8fbbb6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100b4573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100d89190610433565b61010d565b6104c8565b6100eb8361017a565b5f825111806100f75750805b156101085761010683836101b9565b505b505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61014c5f516020610e815f395f51905f52546001600160a01b031690565b604080516001600160a01b03928316815291841660208301520160405180910390a1610177816101e5565b50565b61018381610280565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606101de8383604051806060016040528060278152602001610ea160279139610314565b9392505050565b6001600160a01b03811661024f5760405162461bcd60e51b815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b805f516020610e815f395f51905f525b80546001600160a01b0319166001600160a01b039290921691909117905550565b6001600160a01b0381163b6102ed5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610246565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc61025f565b60605f5f856001600160a01b031685604051610330919061047b565b5f60405180830381855af49150503d805f8114610368576040519150601f19603f3d011682016040523d82523d5f602084013e61036d565b606091505b50909250905061037f86838387610389565b9695505050505050565b606083156103f75782515f036103f0576001600160a01b0385163b6103f05760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610246565b5081610401565b6104018383610409565b949350505050565b8151156104195781518083602001fd5b8060405162461bcd60e51b81526004016102469190610496565b5f60208284031215610443575f5ffd5b81516001600160a01b03811681146101de575f5ffd5b5f5b8381101561047357818101518382015260200161045b565b50505f910152565b5f825161048c818460208701610459565b9190910192915050565b602081525f82518060208401526104b4816040850160208701610459565b601f01601f19169190910160400192915050565b6109ac806104d55f395ff3fe60806040526004361061005d575f3560e01c80635c60da1b116100425780635c60da1b146100a65780638f283970146100e3578063f851a440146101025761006c565b80633659cfe6146100745780634f1ef286146100935761006c565b3661006c5761006a610116565b005b61006a610116565b34801561007f575f5ffd5b5061006a61008e366004610854565b610130565b61006a6100a136600461086d565b610178565b3480156100b1575f5ffd5b506100ba6101eb565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b3480156100ee575f5ffd5b5061006a6100fd366004610854565b610228565b34801561010d575f5ffd5b506100ba610255565b61011e610282565b61012e610129610359565b610362565b565b610138610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101705761016d8160405180602001604052805f8152505f6103bf565b50565b61016d610116565b610180610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101e3576101de8383838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250600192506103bf915050565b505050565b6101de610116565b5f6101f4610380565b73ffffffffffffffffffffffffffffffffffffffff16330361021d57610218610359565b905090565b610225610116565b90565b610230610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101705761016d816103e9565b5f61025e610380565b73ffffffffffffffffffffffffffffffffffffffff16330361021d57610218610380565b61028a610380565b73ffffffffffffffffffffffffffffffffffffffff16330361012e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604260248201527f5472616e73706172656e745570677261646561626c6550726f78793a2061646d60448201527f696e2063616e6e6f742066616c6c6261636b20746f2070726f7879207461726760648201527f6574000000000000000000000000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b5f61021861044a565b365f5f375f5f365f845af43d5f5f3e80801561037c573d5ff35b3d5ffd5b5f7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b5473ffffffffffffffffffffffffffffffffffffffff16919050565b6103c883610471565b5f825111806103d45750805b156101de576103e383836104bd565b50505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f610412610380565b6040805173ffffffffffffffffffffffffffffffffffffffff928316815291841660208301520160405180910390a161016d816104e9565b5f7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6103a3565b61047a816105f5565b60405173ffffffffffffffffffffffffffffffffffffffff8216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606104e28383604051806060016040528060278152602001610979602791396106c0565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff811661058c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610350565b807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b73ffffffffffffffffffffffffffffffffffffffff81163b610699576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201527f6f74206120636f6e7472616374000000000000000000000000000000000000006064820152608401610350565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6105af565b60605f5f8573ffffffffffffffffffffffffffffffffffffffff16856040516106e9919061090d565b5f60405180830381855af49150503d805f8114610721576040519150601f19603f3d011682016040523d82523d5f602084013e610726565b606091505b509150915061073786838387610741565b9695505050505050565b606083156107d65782515f036107cf5773ffffffffffffffffffffffffffffffffffffffff85163b6107cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610350565b50816107e0565b6107e083836107e8565b949350505050565b8151156107f85781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103509190610928565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f575f5ffd5b919050565b5f60208284031215610864575f5ffd5b6104e28261082c565b5f5f5f6040848603121561087f575f5ffd5b6108888461082c565b9250602084013567ffffffffffffffff8111156108a3575f5ffd5b8401601f810186136108b3575f5ffd5b803567ffffffffffffffff8111156108c9575f5ffd5b8660208284010111156108da575f5ffd5b939660209190910195509293505050565b5f5b838110156109055781810151838201526020016108ed565b50505f910152565b5f825161091e8184602087016108eb565b9190910192915050565b602081525f82518060208401526109468160408501602087016108eb565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a164736f6c634300081c000ab53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220becfdf760b178b7f8b65bce7cca5f71bb6fa431d13ab71af3c00556d932d4de364736f6c634300081c00336080604052348015600e575f5ffd5b5060156019565b60c9565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560685760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c65780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b611c60806100d65f395ff3fe608060405234801561000f575f5ffd5b5060043610610115575f3560e01c806370a08231116100ad5780639dc29fac1161007d578063a9059cbb11610063578063a9059cbb146102c0578063d505accf146102d3578063dd62ed3e146102e6575f5ffd5b80639dc29fac1461024b578063a3c573eb1461025e575f5ffd5b806370a08231146102025780637ecebe001461021557806384b0196e1461022857806395d89b4114610243575f5ffd5b806323b872dd116100e857806323b872dd146101a0578063313ce567146101b35780633644e515146101e757806340c10f19146101ef575f5ffd5b806306fdde0314610119578063095ea7b3146101375780631624f6c61461015a57806318160ddd1461016f575b5f5ffd5b61012161034a565b60405161012e919061168e565b60405180910390f35b61014a6101453660046116cf565b610402565b604051901515815260200161012e565b61016d6101683660046117fc565b61041b565b005b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02545b60405190815260200161012e565b61014a6101ae366004611870565b6105f9565b7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d005460405160ff909116815260200161012e565b61019261061c565b61016d6101fd3660046116cf565b61062a565b6101926102103660046118aa565b6106b3565b6101926102233660046118aa565b610703565b61023061070d565b60405161012e97969594939291906118c3565b61012161080c565b61016d6102593660046116cf565b61085d565b7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d0054610100900473ffffffffffffffffffffffffffffffffffffffff1660405173ffffffffffffffffffffffffffffffffffffffff909116815260200161012e565b61014a6102ce3660046116cf565b6108e1565b61016d6102e1366004611982565b6108ee565b6101926102f43660046119e8565b73ffffffffffffffffffffffffffffffffffffffff9182165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b60605f7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b905080600301805461038090611a19565b80601f01602080910402602001604051908101604052809291908181526020018280546103ac90611a19565b80156103f75780601f106103ce576101008083540402835291602001916103f7565b820191905f5260205f20905b8154815290600101906020018083116103da57829003601f168201915b505050505091505090565b5f3361040f818585610ab6565b60019150505b92915050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f811580156104655750825b90505f8267ffffffffffffffff1660011480156104815750303b155b90508115801561048f575080155b156104c6576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156105275784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6105318888610ac3565b61053a88610ad9565b7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d00805460ff88167fffffffffffffffffffffff00000000000000000000000000000000000000000090911617610100330217905583156105ef5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b5f33610606858285610b23565b610611858585610c0f565b506001949350505050565b5f610625610cb8565b905090565b5f7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d008054909150610100900473ffffffffffffffffffffffffffffffffffffffff1633146106a4576040517f38da3b1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106ae8383610cc1565b505050565b5f807f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b73ffffffffffffffffffffffffffffffffffffffff9093165f9081526020939093525050604090205490565b5f61041582610d1b565b5f60608082808083817fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100805490915015801561074b57506001810154155b6107b6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f4549503731323a20556e696e697469616c697a6564000000000000000000000060448201526064015b60405180910390fd5b6107be610d25565b6107c6610d76565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009c939b5091995046985030975095509350915050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060917f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace009161038090611a19565b5f7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d008054909150610100900473ffffffffffffffffffffffffffffffffffffffff1633146108d7576040517f38da3b1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106ae8383610d9f565b5f3361040f818585610c0f565b8342111561092b576040517f62791302000000000000000000000000000000000000000000000000000000008152600481018590526024016107ad565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98888886109a28c73ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f610a0982610df9565b90505f610a1882878787610e40565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610a9f576040517f4b800e4600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301528b1660248201526044016107ad565b610aaa8a8a8a610ab6565b50505050505050505050565b6106ae8383836001610e6c565b610acb610fd6565b610ad5828261103f565b5050565b610ae1610fd6565b610b20816040518060400160405280600181526020017f31000000000000000000000000000000000000000000000000000000000000008152506110a2565b50565b73ffffffffffffffffffffffffffffffffffffffff8381165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610c095781811015610bfb576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064016107ad565b610c0984848484035f610e6c565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610c5e576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f60048201526024016107ad565b73ffffffffffffffffffffffffffffffffffffffff8216610cad576040517fec442f050000000000000000000000000000000000000000000000000000000081525f60048201526024016107ad565b6106ae838383611114565b5f6106256112e1565b73ffffffffffffffffffffffffffffffffffffffff8216610d10576040517fec442f050000000000000000000000000000000000000000000000000000000081525f60048201526024016107ad565b610ad55f8383611114565b5f61041582611354565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060917fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1009161038090611a19565b60605f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10061036f565b73ffffffffffffffffffffffffffffffffffffffff8216610dee576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f60048201526024016107ad565b610ad5825f83611114565b5f610415610e05610cb8565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b5f5f5f5f610e508888888861137c565b925092509250610e60828261146f565b50909695505050505050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff8516610edc576040517fe602df050000000000000000000000000000000000000000000000000000000081525f60048201526024016107ad565b73ffffffffffffffffffffffffffffffffffffffff8416610f2b576040517f94280d620000000000000000000000000000000000000000000000000000000081525f60048201526024016107ad565b73ffffffffffffffffffffffffffffffffffffffff8086165f90815260018301602090815260408083209388168352929052208390558115610fcf578373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92585604051610fc691815260200190565b60405180910390a35b5050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff1661103d576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b611047610fd6565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace007f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace036110938482611aae565b5060048101610c098382611aae565b6110aa610fd6565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1007fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1026110f68482611aae565b50600381016111058382611aae565b505f8082556001909101555050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff841661116e5781816002015f8282546111639190611bc5565b9091555061121e9050565b73ffffffffffffffffffffffffffffffffffffffff84165f90815260208290526040902054828110156111f3576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8616600482015260248101829052604481018490526064016107ad565b73ffffffffffffffffffffffffffffffffffffffff85165f9081526020839052604090209083900390555b73ffffffffffffffffffffffffffffffffffffffff8316611249576002810180548390039055611274565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526020829052604090208054830190555b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516112d391815260200190565b60405180910390a350505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f61130b611572565b6113136115ed565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f807f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006106d7565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156113b557505f91506003905082611465565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015611406573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811661145c57505f925060019150829050611465565b92505f91508190505b9450945094915050565b5f82600381111561148257611482611bfd565b0361148b575050565b600182600381111561149f5761149f611bfd565b036114d6576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028260038111156114ea576114ea611bfd565b03611524576040517ffce698f7000000000000000000000000000000000000000000000000000000008152600481018290526024016107ad565b600382600381111561153857611538611bfd565b03610ad5576040517fd78bce0c000000000000000000000000000000000000000000000000000000008152600481018290526024016107ad565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1008161159d610d25565b8051909150156115b557805160209091012092915050565b815480156115c4579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10081611618610d76565b80519091501561163057805160209091012092915050565b600182015480156115c4579392505050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f6116a06020830184611642565b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146116ca575f5ffd5b919050565b5f5f604083850312156116e0575f5ffd5b6116e9836116a7565b946020939093013593505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112611733575f5ffd5b813567ffffffffffffffff81111561174d5761174d6116f7565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff821117156117b9576117b96116f7565b6040528181528382016020018510156117d0575f5ffd5b816020850160208301375f918101602001919091529392505050565b803560ff811681146116ca575f5ffd5b5f5f5f6060848603121561180e575f5ffd5b833567ffffffffffffffff811115611824575f5ffd5b61183086828701611724565b935050602084013567ffffffffffffffff81111561184c575f5ffd5b61185886828701611724565b925050611867604085016117ec565b90509250925092565b5f5f5f60608486031215611882575f5ffd5b61188b846116a7565b9250611899602085016116a7565b929592945050506040919091013590565b5f602082840312156118ba575f5ffd5b6116a0826116a7565b7fff000000000000000000000000000000000000000000000000000000000000008816815260e060208201525f6118fd60e0830189611642565b828103604084015261190f8189611642565b6060840188905273ffffffffffffffffffffffffffffffffffffffff8716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015611971578351835260209384019390920191600101611953565b50909b9a5050505050505050505050565b5f5f5f5f5f5f5f60e0888a031215611998575f5ffd5b6119a1886116a7565b96506119af602089016116a7565b955060408801359450606088013593506119cb608089016117ec565b9699959850939692959460a0840135945060c09093013592915050565b5f5f604083850312156119f9575f5ffd5b611a02836116a7565b9150611a10602084016116a7565b90509250929050565b600181811c90821680611a2d57607f821691505b602082108103611a64577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b601f8211156106ae57805f5260205f20601f840160051c81016020851015611a8f5750805b601f840160051c820191505b81811015610fcf575f8155600101611a9b565b815167ffffffffffffffff811115611ac857611ac86116f7565b611adc81611ad68454611a19565b84611a6a565b6020601f821160018114611b2d575f8315611af75750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455610fcf565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015611b7a5787850151825560209485019460019092019101611b5a565b5084821015611bb657868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b80820180821115610415577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffdfea264697066735822122091c6c245e1c344b558d8103a850c099a6d0b612735d3d26cebd98dd6e4cfdeb964736f6c634300081c0033",
}

// Bridgel2sovereignchainv1010ABI is the input ABI used to generate the binding from.
// Deprecated: Use Bridgel2sovereignchainv1010MetaData.ABI instead.
var Bridgel2sovereignchainv1010ABI = Bridgel2sovereignchainv1010MetaData.ABI

// Bridgel2sovereignchainv1010Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Bridgel2sovereignchainv1010MetaData.Bin instead.
var Bridgel2sovereignchainv1010Bin = Bridgel2sovereignchainv1010MetaData.Bin

// DeployBridgel2sovereignchainv1010 deploys a new Ethereum contract, binding an instance of Bridgel2sovereignchainv1010 to it.
func DeployBridgel2sovereignchainv1010(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bridgel2sovereignchainv1010, error) {
	parsed, err := Bridgel2sovereignchainv1010MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Bridgel2sovereignchainv1010Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridgel2sovereignchainv1010{Bridgel2sovereignchainv1010Caller: Bridgel2sovereignchainv1010Caller{contract: contract}, Bridgel2sovereignchainv1010Transactor: Bridgel2sovereignchainv1010Transactor{contract: contract}, Bridgel2sovereignchainv1010Filterer: Bridgel2sovereignchainv1010Filterer{contract: contract}}, nil
}

// Bridgel2sovereignchainv1010 is an auto generated Go binding around an Ethereum contract.
type Bridgel2sovereignchainv1010 struct {
	Bridgel2sovereignchainv1010Caller     // Read-only binding to the contract
	Bridgel2sovereignchainv1010Transactor // Write-only binding to the contract
	Bridgel2sovereignchainv1010Filterer   // Log filterer for contract events
}

// Bridgel2sovereignchainv1010Caller is an auto generated read-only Go binding around an Ethereum contract.
type Bridgel2sovereignchainv1010Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bridgel2sovereignchainv1010Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Bridgel2sovereignchainv1010Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bridgel2sovereignchainv1010Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Bridgel2sovereignchainv1010Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bridgel2sovereignchainv1010Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Bridgel2sovereignchainv1010Session struct {
	Contract     *Bridgel2sovereignchainv1010 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// Bridgel2sovereignchainv1010CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Bridgel2sovereignchainv1010CallerSession struct {
	Contract *Bridgel2sovereignchainv1010Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// Bridgel2sovereignchainv1010TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Bridgel2sovereignchainv1010TransactorSession struct {
	Contract     *Bridgel2sovereignchainv1010Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// Bridgel2sovereignchainv1010Raw is an auto generated low-level Go binding around an Ethereum contract.
type Bridgel2sovereignchainv1010Raw struct {
	Contract *Bridgel2sovereignchainv1010 // Generic contract binding to access the raw methods on
}

// Bridgel2sovereignchainv1010CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Bridgel2sovereignchainv1010CallerRaw struct {
	Contract *Bridgel2sovereignchainv1010Caller // Generic read-only contract binding to access the raw methods on
}

// Bridgel2sovereignchainv1010TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Bridgel2sovereignchainv1010TransactorRaw struct {
	Contract *Bridgel2sovereignchainv1010Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgel2sovereignchainv1010 creates a new instance of Bridgel2sovereignchainv1010, bound to a specific deployed contract.
func NewBridgel2sovereignchainv1010(address common.Address, backend bind.ContractBackend) (*Bridgel2sovereignchainv1010, error) {
	contract, err := bindBridgel2sovereignchainv1010(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainv1010{Bridgel2sovereignchainv1010Caller: Bridgel2sovereignchainv1010Caller{contract: contract}, Bridgel2sovereignchainv1010Transactor: Bridgel2sovereignchainv1010Transactor{contract: contract}, Bridgel2sovereignchainv1010Filterer: Bridgel2sovereignchainv1010Filterer{contract: contract}}, nil
}

// NewBridgel2sovereignchainv1010Caller creates a new read-only instance of Bridgel2sovereignchainv1010, bound to a specific deployed contract.
func NewBridgel2sovereignchainv1010Caller(address common.Address, caller bind.ContractCaller) (*Bridgel2sovereignchainv1010Caller, error) {
	contract, err := bindBridgel2sovereignchainv1010(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainv1010Caller{contract: contract}, nil
}

// NewBridgel2sovereignchainv1010Transactor creates a new write-only instance of Bridgel2sovereignchainv1010, bound to a specific deployed contract.
func NewBridgel2sovereignchainv1010Transactor(address common.Address, transactor bind.ContractTransactor) (*Bridgel2sovereignchainv1010Transactor, error) {
	contract, err := bindBridgel2sovereignchainv1010(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainv1010Transactor{contract: contract}, nil
}

// NewBridgel2sovereignchainv1010Filterer creates a new log filterer instance of Bridgel2sovereignchainv1010, bound to a specific deployed contract.
func NewBridgel2sovereignchainv1010Filterer(address common.Address, filterer bind.ContractFilterer) (*Bridgel2sovereignchainv1010Filterer, error) {
	contract, err := bindBridgel2sovereignchainv1010(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainv1010Filterer{contract: contract}, nil
}

// bindBridgel2sovereignchainv1010 binds a generic wrapper to an already deployed contract.
func bindBridgel2sovereignchainv1010(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Bridgel2sovereignchainv1010MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridgel2sovereignchainv1010.Contract.Bridgel2sovereignchainv1010Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.Bridgel2sovereignchainv1010Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.Bridgel2sovereignchainv1010Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridgel2sovereignchainv1010.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.contract.Transact(opts, method, params...)
}

// BRIDGESOVEREIGNVERSION is a free data retrieval call binding the contract method 0xf67566e4.
//
// Solidity: function BRIDGE_SOVEREIGN_VERSION() view returns(string)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) BRIDGESOVEREIGNVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "BRIDGE_SOVEREIGN_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BRIDGESOVEREIGNVERSION is a free data retrieval call binding the contract method 0xf67566e4.
//
// Solidity: function BRIDGE_SOVEREIGN_VERSION() view returns(string)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) BRIDGESOVEREIGNVERSION() (string, error) {
	return _Bridgel2sovereignchainv1010.Contract.BRIDGESOVEREIGNVERSION(&_Bridgel2sovereignchainv1010.CallOpts)
}

// BRIDGESOVEREIGNVERSION is a free data retrieval call binding the contract method 0xf67566e4.
//
// Solidity: function BRIDGE_SOVEREIGN_VERSION() view returns(string)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) BRIDGESOVEREIGNVERSION() (string, error) {
	return _Bridgel2sovereignchainv1010.Contract.BRIDGESOVEREIGNVERSION(&_Bridgel2sovereignchainv1010.CallOpts)
}

// BRIDGEVERSION is a free data retrieval call binding the contract method 0x65d6f654.
//
// Solidity: function BRIDGE_VERSION() view returns(string)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) BRIDGEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "BRIDGE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BRIDGEVERSION is a free data retrieval call binding the contract method 0x65d6f654.
//
// Solidity: function BRIDGE_VERSION() view returns(string)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) BRIDGEVERSION() (string, error) {
	return _Bridgel2sovereignchainv1010.Contract.BRIDGEVERSION(&_Bridgel2sovereignchainv1010.CallOpts)
}

// BRIDGEVERSION is a free data retrieval call binding the contract method 0x65d6f654.
//
// Solidity: function BRIDGE_VERSION() view returns(string)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) BRIDGEVERSION() (string, error) {
	return _Bridgel2sovereignchainv1010.Contract.BRIDGEVERSION(&_Bridgel2sovereignchainv1010.CallOpts)
}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() view returns(bytes)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) INITBYTECODETRANSPARENTPROXY(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "INIT_BYTECODE_TRANSPARENT_PROXY")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() view returns(bytes)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) INITBYTECODETRANSPARENTPROXY() ([]byte, error) {
	return _Bridgel2sovereignchainv1010.Contract.INITBYTECODETRANSPARENTPROXY(&_Bridgel2sovereignchainv1010.CallOpts)
}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() view returns(bytes)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) INITBYTECODETRANSPARENTPROXY() ([]byte, error) {
	return _Bridgel2sovereignchainv1010.Contract.INITBYTECODETRANSPARENTPROXY(&_Bridgel2sovereignchainv1010.CallOpts)
}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) WETHToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "WETHToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) WETHToken() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.WETHToken(&_Bridgel2sovereignchainv1010.CallOpts)
}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) WETHToken() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.WETHToken(&_Bridgel2sovereignchainv1010.CallOpts)
}

// BridgeManager is a free data retrieval call binding the contract method 0x14cc01a0.
//
// Solidity: function bridgeManager() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) BridgeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "bridgeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeManager is a free data retrieval call binding the contract method 0x14cc01a0.
//
// Solidity: function bridgeManager() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) BridgeManager() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.BridgeManager(&_Bridgel2sovereignchainv1010.CallOpts)
}

// BridgeManager is a free data retrieval call binding the contract method 0x14cc01a0.
//
// Solidity: function bridgeManager() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) BridgeManager() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.BridgeManager(&_Bridgel2sovereignchainv1010.CallOpts)
}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) CalculateRoot(opts *bind.CallOpts, leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "calculateRoot", leafHash, smtProof, index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) CalculateRoot(leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	return _Bridgel2sovereignchainv1010.Contract.CalculateRoot(&_Bridgel2sovereignchainv1010.CallOpts, leafHash, smtProof, index)
}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) CalculateRoot(leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	return _Bridgel2sovereignchainv1010.Contract.CalculateRoot(&_Bridgel2sovereignchainv1010.CallOpts, leafHash, smtProof, index)
}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) ClaimedBitMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "claimedBitMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) ClaimedBitMap(arg0 *big.Int) (*big.Int, error) {
	return _Bridgel2sovereignchainv1010.Contract.ClaimedBitMap(&_Bridgel2sovereignchainv1010.CallOpts, arg0)
}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) ClaimedBitMap(arg0 *big.Int) (*big.Int, error) {
	return _Bridgel2sovereignchainv1010.Contract.ClaimedBitMap(&_Bridgel2sovereignchainv1010.CallOpts, arg0)
}

// ClaimedGlobalIndexHashChain is a free data retrieval call binding the contract method 0x1d081d8c.
//
// Solidity: function claimedGlobalIndexHashChain() view returns(bytes32)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) ClaimedGlobalIndexHashChain(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "claimedGlobalIndexHashChain")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ClaimedGlobalIndexHashChain is a free data retrieval call binding the contract method 0x1d081d8c.
//
// Solidity: function claimedGlobalIndexHashChain() view returns(bytes32)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) ClaimedGlobalIndexHashChain() ([32]byte, error) {
	return _Bridgel2sovereignchainv1010.Contract.ClaimedGlobalIndexHashChain(&_Bridgel2sovereignchainv1010.CallOpts)
}

// ClaimedGlobalIndexHashChain is a free data retrieval call binding the contract method 0x1d081d8c.
//
// Solidity: function claimedGlobalIndexHashChain() view returns(bytes32)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) ClaimedGlobalIndexHashChain() ([32]byte, error) {
	return _Bridgel2sovereignchainv1010.Contract.ClaimedGlobalIndexHashChain(&_Bridgel2sovereignchainv1010.CallOpts)
}

// ComputeTokenProxyAddress is a free data retrieval call binding the contract method 0xf214e161.
//
// Solidity: function computeTokenProxyAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) ComputeTokenProxyAddress(opts *bind.CallOpts, originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "computeTokenProxyAddress", originNetwork, originTokenAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeTokenProxyAddress is a free data retrieval call binding the contract method 0xf214e161.
//
// Solidity: function computeTokenProxyAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) ComputeTokenProxyAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.ComputeTokenProxyAddress(&_Bridgel2sovereignchainv1010.CallOpts, originNetwork, originTokenAddress)
}

// ComputeTokenProxyAddress is a free data retrieval call binding the contract method 0xf214e161.
//
// Solidity: function computeTokenProxyAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) ComputeTokenProxyAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.ComputeTokenProxyAddress(&_Bridgel2sovereignchainv1010.CallOpts, originNetwork, originTokenAddress)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) DepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "depositCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) DepositCount() (*big.Int, error) {
	return _Bridgel2sovereignchainv1010.Contract.DepositCount(&_Bridgel2sovereignchainv1010.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) DepositCount() (*big.Int, error) {
	return _Bridgel2sovereignchainv1010.Contract.DepositCount(&_Bridgel2sovereignchainv1010.CallOpts)
}

// EmergencyBridgePauser is a free data retrieval call binding the contract method 0x2f84c690.
//
// Solidity: function emergencyBridgePauser() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) EmergencyBridgePauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "emergencyBridgePauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmergencyBridgePauser is a free data retrieval call binding the contract method 0x2f84c690.
//
// Solidity: function emergencyBridgePauser() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) EmergencyBridgePauser() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.EmergencyBridgePauser(&_Bridgel2sovereignchainv1010.CallOpts)
}

// EmergencyBridgePauser is a free data retrieval call binding the contract method 0x2f84c690.
//
// Solidity: function emergencyBridgePauser() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) EmergencyBridgePauser() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.EmergencyBridgePauser(&_Bridgel2sovereignchainv1010.CallOpts)
}

// EmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0xae24490a.
//
// Solidity: function emergencyBridgeUnpauser() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) EmergencyBridgeUnpauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "emergencyBridgeUnpauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0xae24490a.
//
// Solidity: function emergencyBridgeUnpauser() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) EmergencyBridgeUnpauser() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.EmergencyBridgeUnpauser(&_Bridgel2sovereignchainv1010.CallOpts)
}

// EmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0xae24490a.
//
// Solidity: function emergencyBridgeUnpauser() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) EmergencyBridgeUnpauser() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.EmergencyBridgeUnpauser(&_Bridgel2sovereignchainv1010.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) GasTokenAddress() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.GasTokenAddress(&_Bridgel2sovereignchainv1010.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) GasTokenAddress() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.GasTokenAddress(&_Bridgel2sovereignchainv1010.CallOpts)
}

// GasTokenMetadata is a free data retrieval call binding the contract method 0x27aef4e8.
//
// Solidity: function gasTokenMetadata() view returns(bytes)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) GasTokenMetadata(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "gasTokenMetadata")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GasTokenMetadata is a free data retrieval call binding the contract method 0x27aef4e8.
//
// Solidity: function gasTokenMetadata() view returns(bytes)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) GasTokenMetadata() ([]byte, error) {
	return _Bridgel2sovereignchainv1010.Contract.GasTokenMetadata(&_Bridgel2sovereignchainv1010.CallOpts)
}

// GasTokenMetadata is a free data retrieval call binding the contract method 0x27aef4e8.
//
// Solidity: function gasTokenMetadata() view returns(bytes)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) GasTokenMetadata() ([]byte, error) {
	return _Bridgel2sovereignchainv1010.Contract.GasTokenMetadata(&_Bridgel2sovereignchainv1010.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) GasTokenNetwork() (uint32, error) {
	return _Bridgel2sovereignchainv1010.Contract.GasTokenNetwork(&_Bridgel2sovereignchainv1010.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) GasTokenNetwork() (uint32, error) {
	return _Bridgel2sovereignchainv1010.Contract.GasTokenNetwork(&_Bridgel2sovereignchainv1010.CallOpts)
}

// GetLeafValue is a free data retrieval call binding the contract method 0x3e197043.
//
// Solidity: function getLeafValue(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes32 metadataHash) pure returns(bytes32)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) GetLeafValue(opts *bind.CallOpts, leafType uint8, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadataHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "getLeafValue", leafType, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadataHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLeafValue is a free data retrieval call binding the contract method 0x3e197043.
//
// Solidity: function getLeafValue(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes32 metadataHash) pure returns(bytes32)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) GetLeafValue(leafType uint8, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadataHash [32]byte) ([32]byte, error) {
	return _Bridgel2sovereignchainv1010.Contract.GetLeafValue(&_Bridgel2sovereignchainv1010.CallOpts, leafType, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadataHash)
}

// GetLeafValue is a free data retrieval call binding the contract method 0x3e197043.
//
// Solidity: function getLeafValue(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes32 metadataHash) pure returns(bytes32)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) GetLeafValue(leafType uint8, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadataHash [32]byte) ([32]byte, error) {
	return _Bridgel2sovereignchainv1010.Contract.GetLeafValue(&_Bridgel2sovereignchainv1010.CallOpts, leafType, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadataHash)
}

// GetProxiedTokensManager is a free data retrieval call binding the contract method 0x38b8fbbb.
//
// Solidity: function getProxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) GetProxiedTokensManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "getProxiedTokensManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxiedTokensManager is a free data retrieval call binding the contract method 0x38b8fbbb.
//
// Solidity: function getProxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) GetProxiedTokensManager() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.GetProxiedTokensManager(&_Bridgel2sovereignchainv1010.CallOpts)
}

// GetProxiedTokensManager is a free data retrieval call binding the contract method 0x38b8fbbb.
//
// Solidity: function getProxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) GetProxiedTokensManager() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.GetProxiedTokensManager(&_Bridgel2sovereignchainv1010.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) GetRoot() ([32]byte, error) {
	return _Bridgel2sovereignchainv1010.Contract.GetRoot(&_Bridgel2sovereignchainv1010.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) GetRoot() ([32]byte, error) {
	return _Bridgel2sovereignchainv1010.Contract.GetRoot(&_Bridgel2sovereignchainv1010.CallOpts)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) GetTokenMetadata(opts *bind.CallOpts, token common.Address) ([]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "getTokenMetadata", token)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) GetTokenMetadata(token common.Address) ([]byte, error) {
	return _Bridgel2sovereignchainv1010.Contract.GetTokenMetadata(&_Bridgel2sovereignchainv1010.CallOpts, token)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) GetTokenMetadata(token common.Address) ([]byte, error) {
	return _Bridgel2sovereignchainv1010.Contract.GetTokenMetadata(&_Bridgel2sovereignchainv1010.CallOpts, token)
}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) GetTokenWrappedAddress(opts *bind.CallOpts, originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "getTokenWrappedAddress", originNetwork, originTokenAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) GetTokenWrappedAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.GetTokenWrappedAddress(&_Bridgel2sovereignchainv1010.CallOpts, originNetwork, originTokenAddress)
}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) GetTokenWrappedAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.GetTokenWrappedAddress(&_Bridgel2sovereignchainv1010.CallOpts, originNetwork, originTokenAddress)
}

// GetWrappedTokenBridgeImplementation is a free data retrieval call binding the contract method 0x3b2fee9a.
//
// Solidity: function getWrappedTokenBridgeImplementation() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) GetWrappedTokenBridgeImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "getWrappedTokenBridgeImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWrappedTokenBridgeImplementation is a free data retrieval call binding the contract method 0x3b2fee9a.
//
// Solidity: function getWrappedTokenBridgeImplementation() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) GetWrappedTokenBridgeImplementation() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.GetWrappedTokenBridgeImplementation(&_Bridgel2sovereignchainv1010.CallOpts)
}

// GetWrappedTokenBridgeImplementation is a free data retrieval call binding the contract method 0x3b2fee9a.
//
// Solidity: function getWrappedTokenBridgeImplementation() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) GetWrappedTokenBridgeImplementation() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.GetWrappedTokenBridgeImplementation(&_Bridgel2sovereignchainv1010.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) GlobalExitRootManager() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.GlobalExitRootManager(&_Bridgel2sovereignchainv1010.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.GlobalExitRootManager(&_Bridgel2sovereignchainv1010.CallOpts)
}

// Initialize0 is a free data retrieval call binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() pure returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) Initialize0(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "initialize0")

	if err != nil {
		return err
	}

	return err

}

// Initialize0 is a free data retrieval call binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() pure returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) Initialize0() error {
	return _Bridgel2sovereignchainv1010.Contract.Initialize0(&_Bridgel2sovereignchainv1010.CallOpts)
}

// Initialize0 is a free data retrieval call binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() pure returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) Initialize0() error {
	return _Bridgel2sovereignchainv1010.Contract.Initialize0(&_Bridgel2sovereignchainv1010.CallOpts)
}

// IsClaimed is a free data retrieval call binding the contract method 0xcc461632.
//
// Solidity: function isClaimed(uint32 leafIndex, uint32 sourceBridgeNetwork) view returns(bool)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) IsClaimed(opts *bind.CallOpts, leafIndex uint32, sourceBridgeNetwork uint32) (bool, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "isClaimed", leafIndex, sourceBridgeNetwork)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClaimed is a free data retrieval call binding the contract method 0xcc461632.
//
// Solidity: function isClaimed(uint32 leafIndex, uint32 sourceBridgeNetwork) view returns(bool)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) IsClaimed(leafIndex uint32, sourceBridgeNetwork uint32) (bool, error) {
	return _Bridgel2sovereignchainv1010.Contract.IsClaimed(&_Bridgel2sovereignchainv1010.CallOpts, leafIndex, sourceBridgeNetwork)
}

// IsClaimed is a free data retrieval call binding the contract method 0xcc461632.
//
// Solidity: function isClaimed(uint32 leafIndex, uint32 sourceBridgeNetwork) view returns(bool)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) IsClaimed(leafIndex uint32, sourceBridgeNetwork uint32) (bool, error) {
	return _Bridgel2sovereignchainv1010.Contract.IsClaimed(&_Bridgel2sovereignchainv1010.CallOpts, leafIndex, sourceBridgeNetwork)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) IsEmergencyState() (bool, error) {
	return _Bridgel2sovereignchainv1010.Contract.IsEmergencyState(&_Bridgel2sovereignchainv1010.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) IsEmergencyState() (bool, error) {
	return _Bridgel2sovereignchainv1010.Contract.IsEmergencyState(&_Bridgel2sovereignchainv1010.CallOpts)
}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) LastUpdatedDepositCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "lastUpdatedDepositCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) LastUpdatedDepositCount() (uint32, error) {
	return _Bridgel2sovereignchainv1010.Contract.LastUpdatedDepositCount(&_Bridgel2sovereignchainv1010.CallOpts)
}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) LastUpdatedDepositCount() (uint32, error) {
	return _Bridgel2sovereignchainv1010.Contract.LastUpdatedDepositCount(&_Bridgel2sovereignchainv1010.CallOpts)
}

// LocalBalanceTree is a free data retrieval call binding the contract method 0xd9cb3aec.
//
// Solidity: function localBalanceTree(bytes32 tokenInfoHash) view returns(uint256 amount)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) LocalBalanceTree(opts *bind.CallOpts, tokenInfoHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "localBalanceTree", tokenInfoHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LocalBalanceTree is a free data retrieval call binding the contract method 0xd9cb3aec.
//
// Solidity: function localBalanceTree(bytes32 tokenInfoHash) view returns(uint256 amount)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) LocalBalanceTree(tokenInfoHash [32]byte) (*big.Int, error) {
	return _Bridgel2sovereignchainv1010.Contract.LocalBalanceTree(&_Bridgel2sovereignchainv1010.CallOpts, tokenInfoHash)
}

// LocalBalanceTree is a free data retrieval call binding the contract method 0xd9cb3aec.
//
// Solidity: function localBalanceTree(bytes32 tokenInfoHash) view returns(uint256 amount)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) LocalBalanceTree(tokenInfoHash [32]byte) (*big.Int, error) {
	return _Bridgel2sovereignchainv1010.Contract.LocalBalanceTree(&_Bridgel2sovereignchainv1010.CallOpts, tokenInfoHash)
}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) NetworkID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "networkID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) NetworkID() (uint32, error) {
	return _Bridgel2sovereignchainv1010.Contract.NetworkID(&_Bridgel2sovereignchainv1010.CallOpts)
}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) NetworkID() (uint32, error) {
	return _Bridgel2sovereignchainv1010.Contract.NetworkID(&_Bridgel2sovereignchainv1010.CallOpts)
}

// PendingEmergencyBridgePauser is a free data retrieval call binding the contract method 0x03e6e116.
//
// Solidity: function pendingEmergencyBridgePauser() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) PendingEmergencyBridgePauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "pendingEmergencyBridgePauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingEmergencyBridgePauser is a free data retrieval call binding the contract method 0x03e6e116.
//
// Solidity: function pendingEmergencyBridgePauser() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) PendingEmergencyBridgePauser() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.PendingEmergencyBridgePauser(&_Bridgel2sovereignchainv1010.CallOpts)
}

// PendingEmergencyBridgePauser is a free data retrieval call binding the contract method 0x03e6e116.
//
// Solidity: function pendingEmergencyBridgePauser() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) PendingEmergencyBridgePauser() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.PendingEmergencyBridgePauser(&_Bridgel2sovereignchainv1010.CallOpts)
}

// PendingEmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0x606617ff.
//
// Solidity: function pendingEmergencyBridgeUnpauser() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) PendingEmergencyBridgeUnpauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "pendingEmergencyBridgeUnpauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingEmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0x606617ff.
//
// Solidity: function pendingEmergencyBridgeUnpauser() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) PendingEmergencyBridgeUnpauser() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.PendingEmergencyBridgeUnpauser(&_Bridgel2sovereignchainv1010.CallOpts)
}

// PendingEmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0x606617ff.
//
// Solidity: function pendingEmergencyBridgeUnpauser() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) PendingEmergencyBridgeUnpauser() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.PendingEmergencyBridgeUnpauser(&_Bridgel2sovereignchainv1010.CallOpts)
}

// PendingProxiedTokensManager is a free data retrieval call binding the contract method 0xece93c6f.
//
// Solidity: function pendingProxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) PendingProxiedTokensManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "pendingProxiedTokensManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingProxiedTokensManager is a free data retrieval call binding the contract method 0xece93c6f.
//
// Solidity: function pendingProxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) PendingProxiedTokensManager() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.PendingProxiedTokensManager(&_Bridgel2sovereignchainv1010.CallOpts)
}

// PendingProxiedTokensManager is a free data retrieval call binding the contract method 0xece93c6f.
//
// Solidity: function pendingProxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) PendingProxiedTokensManager() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.PendingProxiedTokensManager(&_Bridgel2sovereignchainv1010.CallOpts)
}

// PolygonRollupManager is a free data retrieval call binding the contract method 0x8ed7e3f2.
//
// Solidity: function polygonRollupManager() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) PolygonRollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "polygonRollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PolygonRollupManager is a free data retrieval call binding the contract method 0x8ed7e3f2.
//
// Solidity: function polygonRollupManager() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) PolygonRollupManager() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.PolygonRollupManager(&_Bridgel2sovereignchainv1010.CallOpts)
}

// PolygonRollupManager is a free data retrieval call binding the contract method 0x8ed7e3f2.
//
// Solidity: function polygonRollupManager() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) PolygonRollupManager() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.PolygonRollupManager(&_Bridgel2sovereignchainv1010.CallOpts)
}

// ProxiedTokensManager is a free data retrieval call binding the contract method 0x6e4ecfed.
//
// Solidity: function proxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) ProxiedTokensManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "proxiedTokensManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxiedTokensManager is a free data retrieval call binding the contract method 0x6e4ecfed.
//
// Solidity: function proxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) ProxiedTokensManager() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.ProxiedTokensManager(&_Bridgel2sovereignchainv1010.CallOpts)
}

// ProxiedTokensManager is a free data retrieval call binding the contract method 0x6e4ecfed.
//
// Solidity: function proxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) ProxiedTokensManager() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.ProxiedTokensManager(&_Bridgel2sovereignchainv1010.CallOpts)
}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) TokenInfoToWrappedToken(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "tokenInfoToWrappedToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) TokenInfoToWrappedToken(arg0 [32]byte) (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.TokenInfoToWrappedToken(&_Bridgel2sovereignchainv1010.CallOpts, arg0)
}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) TokenInfoToWrappedToken(arg0 [32]byte) (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.TokenInfoToWrappedToken(&_Bridgel2sovereignchainv1010.CallOpts, arg0)
}

// UnsetGlobalIndexHashChain is a free data retrieval call binding the contract method 0x6ee84b23.
//
// Solidity: function unsetGlobalIndexHashChain() view returns(bytes32)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) UnsetGlobalIndexHashChain(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "unsetGlobalIndexHashChain")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UnsetGlobalIndexHashChain is a free data retrieval call binding the contract method 0x6ee84b23.
//
// Solidity: function unsetGlobalIndexHashChain() view returns(bytes32)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) UnsetGlobalIndexHashChain() ([32]byte, error) {
	return _Bridgel2sovereignchainv1010.Contract.UnsetGlobalIndexHashChain(&_Bridgel2sovereignchainv1010.CallOpts)
}

// UnsetGlobalIndexHashChain is a free data retrieval call binding the contract method 0x6ee84b23.
//
// Solidity: function unsetGlobalIndexHashChain() view returns(bytes32)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) UnsetGlobalIndexHashChain() ([32]byte, error) {
	return _Bridgel2sovereignchainv1010.Contract.UnsetGlobalIndexHashChain(&_Bridgel2sovereignchainv1010.CallOpts)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) VerifyMerkleProof(opts *bind.CallOpts, leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "verifyMerkleProof", leafHash, smtProof, index, root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	return _Bridgel2sovereignchainv1010.Contract.VerifyMerkleProof(&_Bridgel2sovereignchainv1010.CallOpts, leafHash, smtProof, index, root)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	return _Bridgel2sovereignchainv1010.Contract.VerifyMerkleProof(&_Bridgel2sovereignchainv1010.CallOpts, leafHash, smtProof, index, root)
}

// WrappedAddressIsNotMintable is a free data retrieval call binding the contract method 0xc0f49163.
//
// Solidity: function wrappedAddressIsNotMintable(address wrappedAddress) view returns(bool isNotMintable)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) WrappedAddressIsNotMintable(opts *bind.CallOpts, wrappedAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "wrappedAddressIsNotMintable", wrappedAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WrappedAddressIsNotMintable is a free data retrieval call binding the contract method 0xc0f49163.
//
// Solidity: function wrappedAddressIsNotMintable(address wrappedAddress) view returns(bool isNotMintable)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) WrappedAddressIsNotMintable(wrappedAddress common.Address) (bool, error) {
	return _Bridgel2sovereignchainv1010.Contract.WrappedAddressIsNotMintable(&_Bridgel2sovereignchainv1010.CallOpts, wrappedAddress)
}

// WrappedAddressIsNotMintable is a free data retrieval call binding the contract method 0xc0f49163.
//
// Solidity: function wrappedAddressIsNotMintable(address wrappedAddress) view returns(bool isNotMintable)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) WrappedAddressIsNotMintable(wrappedAddress common.Address) (bool, error) {
	return _Bridgel2sovereignchainv1010.Contract.WrappedAddressIsNotMintable(&_Bridgel2sovereignchainv1010.CallOpts, wrappedAddress)
}

// WrappedTokenBytecodeStorer is a free data retrieval call binding the contract method 0x381fef6d.
//
// Solidity: function wrappedTokenBytecodeStorer() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) WrappedTokenBytecodeStorer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "wrappedTokenBytecodeStorer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedTokenBytecodeStorer is a free data retrieval call binding the contract method 0x381fef6d.
//
// Solidity: function wrappedTokenBytecodeStorer() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) WrappedTokenBytecodeStorer() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.WrappedTokenBytecodeStorer(&_Bridgel2sovereignchainv1010.CallOpts)
}

// WrappedTokenBytecodeStorer is a free data retrieval call binding the contract method 0x381fef6d.
//
// Solidity: function wrappedTokenBytecodeStorer() view returns(address)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) WrappedTokenBytecodeStorer() (common.Address, error) {
	return _Bridgel2sovereignchainv1010.Contract.WrappedTokenBytecodeStorer(&_Bridgel2sovereignchainv1010.CallOpts)
}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address ) view returns(uint32 originNetwork, address originTokenAddress)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Caller) WrappedTokenToTokenInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainv1010.contract.Call(opts, &out, "wrappedTokenToTokenInfo", arg0)

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
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) WrappedTokenToTokenInfo(arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	return _Bridgel2sovereignchainv1010.Contract.WrappedTokenToTokenInfo(&_Bridgel2sovereignchainv1010.CallOpts, arg0)
}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address ) view returns(uint32 originNetwork, address originTokenAddress)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010CallerSession) WrappedTokenToTokenInfo(arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	return _Bridgel2sovereignchainv1010.Contract.WrappedTokenToTokenInfo(&_Bridgel2sovereignchainv1010.CallOpts, arg0)
}

// AcceptEmergencyBridgePauserRole is a paid mutator transaction binding the contract method 0xe88f0436.
//
// Solidity: function acceptEmergencyBridgePauserRole() returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Transactor) AcceptEmergencyBridgePauserRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.contract.Transact(opts, "acceptEmergencyBridgePauserRole")
}

// AcceptEmergencyBridgePauserRole is a paid mutator transaction binding the contract method 0xe88f0436.
//
// Solidity: function acceptEmergencyBridgePauserRole() returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) AcceptEmergencyBridgePauserRole() (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.AcceptEmergencyBridgePauserRole(&_Bridgel2sovereignchainv1010.TransactOpts)
}

// AcceptEmergencyBridgePauserRole is a paid mutator transaction binding the contract method 0xe88f0436.
//
// Solidity: function acceptEmergencyBridgePauserRole() returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010TransactorSession) AcceptEmergencyBridgePauserRole() (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.AcceptEmergencyBridgePauserRole(&_Bridgel2sovereignchainv1010.TransactOpts)
}

// AcceptEmergencyBridgeUnpauserRole is a paid mutator transaction binding the contract method 0x8b37b873.
//
// Solidity: function acceptEmergencyBridgeUnpauserRole() returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Transactor) AcceptEmergencyBridgeUnpauserRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.contract.Transact(opts, "acceptEmergencyBridgeUnpauserRole")
}

// AcceptEmergencyBridgeUnpauserRole is a paid mutator transaction binding the contract method 0x8b37b873.
//
// Solidity: function acceptEmergencyBridgeUnpauserRole() returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) AcceptEmergencyBridgeUnpauserRole() (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.AcceptEmergencyBridgeUnpauserRole(&_Bridgel2sovereignchainv1010.TransactOpts)
}

// AcceptEmergencyBridgeUnpauserRole is a paid mutator transaction binding the contract method 0x8b37b873.
//
// Solidity: function acceptEmergencyBridgeUnpauserRole() returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010TransactorSession) AcceptEmergencyBridgeUnpauserRole() (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.AcceptEmergencyBridgeUnpauserRole(&_Bridgel2sovereignchainv1010.TransactOpts)
}

// AcceptProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8c668f1c.
//
// Solidity: function acceptProxiedTokensManagerRole() returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Transactor) AcceptProxiedTokensManagerRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.contract.Transact(opts, "acceptProxiedTokensManagerRole")
}

// AcceptProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8c668f1c.
//
// Solidity: function acceptProxiedTokensManagerRole() returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) AcceptProxiedTokensManagerRole() (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.AcceptProxiedTokensManagerRole(&_Bridgel2sovereignchainv1010.TransactOpts)
}

// AcceptProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8c668f1c.
//
// Solidity: function acceptProxiedTokensManagerRole() returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010TransactorSession) AcceptProxiedTokensManagerRole() (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.AcceptProxiedTokensManagerRole(&_Bridgel2sovereignchainv1010.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Transactor) ActivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.contract.Transact(opts, "activateEmergencyState")
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) ActivateEmergencyState() (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.ActivateEmergencyState(&_Bridgel2sovereignchainv1010.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010TransactorSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.ActivateEmergencyState(&_Bridgel2sovereignchainv1010.TransactOpts)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Transactor) BridgeAsset(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.contract.Transact(opts, "bridgeAsset", destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) BridgeAsset(destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.BridgeAsset(&_Bridgel2sovereignchainv1010.TransactOpts, destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010TransactorSession) BridgeAsset(destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.BridgeAsset(&_Bridgel2sovereignchainv1010.TransactOpts, destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Transactor) BridgeMessage(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.contract.Transact(opts, "bridgeMessage", destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) BridgeMessage(destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.BridgeMessage(&_Bridgel2sovereignchainv1010.TransactOpts, destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010TransactorSession) BridgeMessage(destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.BridgeMessage(&_Bridgel2sovereignchainv1010.TransactOpts, destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Transactor) BridgeMessageWETH(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.contract.Transact(opts, "bridgeMessageWETH", destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) BridgeMessageWETH(destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.BridgeMessageWETH(&_Bridgel2sovereignchainv1010.TransactOpts, destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010TransactorSession) BridgeMessageWETH(destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.BridgeMessageWETH(&_Bridgel2sovereignchainv1010.TransactOpts, destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Transactor) ClaimAsset(opts *bind.TransactOpts, smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.contract.Transact(opts, "claimAsset", smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) ClaimAsset(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.ClaimAsset(&_Bridgel2sovereignchainv1010.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010TransactorSession) ClaimAsset(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.ClaimAsset(&_Bridgel2sovereignchainv1010.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Transactor) ClaimMessage(opts *bind.TransactOpts, smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.contract.Transact(opts, "claimMessage", smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) ClaimMessage(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.ClaimMessage(&_Bridgel2sovereignchainv1010.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010TransactorSession) ClaimMessage(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.ClaimMessage(&_Bridgel2sovereignchainv1010.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Transactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.DeactivateEmergencyState(&_Bridgel2sovereignchainv1010.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010TransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.DeactivateEmergencyState(&_Bridgel2sovereignchainv1010.TransactOpts)
}

// DeployWrappedTokenAndRemap is a paid mutator transaction binding the contract method 0x6e974cd4.
//
// Solidity: function deployWrappedTokenAndRemap(uint32 originNetwork, address originTokenAddress, bool isNotMintable) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Transactor) DeployWrappedTokenAndRemap(opts *bind.TransactOpts, originNetwork uint32, originTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.contract.Transact(opts, "deployWrappedTokenAndRemap", originNetwork, originTokenAddress, isNotMintable)
}

// DeployWrappedTokenAndRemap is a paid mutator transaction binding the contract method 0x6e974cd4.
//
// Solidity: function deployWrappedTokenAndRemap(uint32 originNetwork, address originTokenAddress, bool isNotMintable) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) DeployWrappedTokenAndRemap(originNetwork uint32, originTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.DeployWrappedTokenAndRemap(&_Bridgel2sovereignchainv1010.TransactOpts, originNetwork, originTokenAddress, isNotMintable)
}

// DeployWrappedTokenAndRemap is a paid mutator transaction binding the contract method 0x6e974cd4.
//
// Solidity: function deployWrappedTokenAndRemap(uint32 originNetwork, address originTokenAddress, bool isNotMintable) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010TransactorSession) DeployWrappedTokenAndRemap(originNetwork uint32, originTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.DeployWrappedTokenAndRemap(&_Bridgel2sovereignchainv1010.TransactOpts, originNetwork, originTokenAddress, isNotMintable)
}

// Initialize is a paid mutator transaction binding the contract method 0x006ee171.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address _sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable, address _emergencyBridgePauser, address _emergencyBridgeUnpauser, address _proxiedTokensManager) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Transactor) Initialize(opts *bind.TransactOpts, _networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, _sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool, _emergencyBridgePauser common.Address, _emergencyBridgeUnpauser common.Address, _proxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.contract.Transact(opts, "initialize", _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, _sovereignWETHAddress, _sovereignWETHAddressIsNotMintable, _emergencyBridgePauser, _emergencyBridgeUnpauser, _proxiedTokensManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x006ee171.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address _sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable, address _emergencyBridgePauser, address _emergencyBridgeUnpauser, address _proxiedTokensManager) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) Initialize(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, _sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool, _emergencyBridgePauser common.Address, _emergencyBridgeUnpauser common.Address, _proxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.Initialize(&_Bridgel2sovereignchainv1010.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, _sovereignWETHAddress, _sovereignWETHAddressIsNotMintable, _emergencyBridgePauser, _emergencyBridgeUnpauser, _proxiedTokensManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x006ee171.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address _sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable, address _emergencyBridgePauser, address _emergencyBridgeUnpauser, address _proxiedTokensManager) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010TransactorSession) Initialize(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, _sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool, _emergencyBridgePauser common.Address, _emergencyBridgeUnpauser common.Address, _proxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.Initialize(&_Bridgel2sovereignchainv1010.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, _sovereignWETHAddress, _sovereignWETHAddressIsNotMintable, _emergencyBridgePauser, _emergencyBridgeUnpauser, _proxiedTokensManager)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xf0a3d955.
//
// Solidity: function initialize(bytes32[] tokenInfoHash, uint256[] amount, address _emergencyBridgeUnpauser, address _proxiedTokensManager) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Transactor) Initialize1(opts *bind.TransactOpts, tokenInfoHash [][32]byte, amount []*big.Int, _emergencyBridgeUnpauser common.Address, _proxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.contract.Transact(opts, "initialize1", tokenInfoHash, amount, _emergencyBridgeUnpauser, _proxiedTokensManager)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xf0a3d955.
//
// Solidity: function initialize(bytes32[] tokenInfoHash, uint256[] amount, address _emergencyBridgeUnpauser, address _proxiedTokensManager) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) Initialize1(tokenInfoHash [][32]byte, amount []*big.Int, _emergencyBridgeUnpauser common.Address, _proxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.Initialize1(&_Bridgel2sovereignchainv1010.TransactOpts, tokenInfoHash, amount, _emergencyBridgeUnpauser, _proxiedTokensManager)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xf0a3d955.
//
// Solidity: function initialize(bytes32[] tokenInfoHash, uint256[] amount, address _emergencyBridgeUnpauser, address _proxiedTokensManager) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010TransactorSession) Initialize1(tokenInfoHash [][32]byte, amount []*big.Int, _emergencyBridgeUnpauser common.Address, _proxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.Initialize1(&_Bridgel2sovereignchainv1010.TransactOpts, tokenInfoHash, amount, _emergencyBridgeUnpauser, _proxiedTokensManager)
}

// Initialize2 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes ) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Transactor) Initialize2(opts *bind.TransactOpts, arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.contract.Transact(opts, "initialize2", arg0, arg1, arg2, arg3, arg4, arg5)
}

// Initialize2 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes ) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) Initialize2(arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.Initialize2(&_Bridgel2sovereignchainv1010.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// Initialize2 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes ) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010TransactorSession) Initialize2(arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.Initialize2(&_Bridgel2sovereignchainv1010.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// MigrateLegacyToken is a paid mutator transaction binding the contract method 0xb0b37920.
//
// Solidity: function migrateLegacyToken(address legacyTokenAddress, uint256 amount, bytes permitData) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Transactor) MigrateLegacyToken(opts *bind.TransactOpts, legacyTokenAddress common.Address, amount *big.Int, permitData []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.contract.Transact(opts, "migrateLegacyToken", legacyTokenAddress, amount, permitData)
}

// MigrateLegacyToken is a paid mutator transaction binding the contract method 0xb0b37920.
//
// Solidity: function migrateLegacyToken(address legacyTokenAddress, uint256 amount, bytes permitData) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) MigrateLegacyToken(legacyTokenAddress common.Address, amount *big.Int, permitData []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.MigrateLegacyToken(&_Bridgel2sovereignchainv1010.TransactOpts, legacyTokenAddress, amount, permitData)
}

// MigrateLegacyToken is a paid mutator transaction binding the contract method 0xb0b37920.
//
// Solidity: function migrateLegacyToken(address legacyTokenAddress, uint256 amount, bytes permitData) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010TransactorSession) MigrateLegacyToken(legacyTokenAddress common.Address, amount *big.Int, permitData []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.MigrateLegacyToken(&_Bridgel2sovereignchainv1010.TransactOpts, legacyTokenAddress, amount, permitData)
}

// RemoveLegacySovereignTokenAddress is a paid mutator transaction binding the contract method 0xb4586962.
//
// Solidity: function removeLegacySovereignTokenAddress(address legacySovereignTokenAddress) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Transactor) RemoveLegacySovereignTokenAddress(opts *bind.TransactOpts, legacySovereignTokenAddress common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.contract.Transact(opts, "removeLegacySovereignTokenAddress", legacySovereignTokenAddress)
}

// RemoveLegacySovereignTokenAddress is a paid mutator transaction binding the contract method 0xb4586962.
//
// Solidity: function removeLegacySovereignTokenAddress(address legacySovereignTokenAddress) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) RemoveLegacySovereignTokenAddress(legacySovereignTokenAddress common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.RemoveLegacySovereignTokenAddress(&_Bridgel2sovereignchainv1010.TransactOpts, legacySovereignTokenAddress)
}

// RemoveLegacySovereignTokenAddress is a paid mutator transaction binding the contract method 0xb4586962.
//
// Solidity: function removeLegacySovereignTokenAddress(address legacySovereignTokenAddress) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010TransactorSession) RemoveLegacySovereignTokenAddress(legacySovereignTokenAddress common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.RemoveLegacySovereignTokenAddress(&_Bridgel2sovereignchainv1010.TransactOpts, legacySovereignTokenAddress)
}

// SetBridgeManager is a paid mutator transaction binding the contract method 0xeabd372a.
//
// Solidity: function setBridgeManager(address _bridgeManager) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Transactor) SetBridgeManager(opts *bind.TransactOpts, _bridgeManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.contract.Transact(opts, "setBridgeManager", _bridgeManager)
}

// SetBridgeManager is a paid mutator transaction binding the contract method 0xeabd372a.
//
// Solidity: function setBridgeManager(address _bridgeManager) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) SetBridgeManager(_bridgeManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.SetBridgeManager(&_Bridgel2sovereignchainv1010.TransactOpts, _bridgeManager)
}

// SetBridgeManager is a paid mutator transaction binding the contract method 0xeabd372a.
//
// Solidity: function setBridgeManager(address _bridgeManager) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010TransactorSession) SetBridgeManager(_bridgeManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.SetBridgeManager(&_Bridgel2sovereignchainv1010.TransactOpts, _bridgeManager)
}

// SetMultipleSovereignTokenAddress is a paid mutator transaction binding the contract method 0x57cfbee3.
//
// Solidity: function setMultipleSovereignTokenAddress(uint32[] originNetworks, address[] originTokenAddresses, address[] sovereignTokenAddresses, bool[] isNotMintable) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Transactor) SetMultipleSovereignTokenAddress(opts *bind.TransactOpts, originNetworks []uint32, originTokenAddresses []common.Address, sovereignTokenAddresses []common.Address, isNotMintable []bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.contract.Transact(opts, "setMultipleSovereignTokenAddress", originNetworks, originTokenAddresses, sovereignTokenAddresses, isNotMintable)
}

// SetMultipleSovereignTokenAddress is a paid mutator transaction binding the contract method 0x57cfbee3.
//
// Solidity: function setMultipleSovereignTokenAddress(uint32[] originNetworks, address[] originTokenAddresses, address[] sovereignTokenAddresses, bool[] isNotMintable) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) SetMultipleSovereignTokenAddress(originNetworks []uint32, originTokenAddresses []common.Address, sovereignTokenAddresses []common.Address, isNotMintable []bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.SetMultipleSovereignTokenAddress(&_Bridgel2sovereignchainv1010.TransactOpts, originNetworks, originTokenAddresses, sovereignTokenAddresses, isNotMintable)
}

// SetMultipleSovereignTokenAddress is a paid mutator transaction binding the contract method 0x57cfbee3.
//
// Solidity: function setMultipleSovereignTokenAddress(uint32[] originNetworks, address[] originTokenAddresses, address[] sovereignTokenAddresses, bool[] isNotMintable) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010TransactorSession) SetMultipleSovereignTokenAddress(originNetworks []uint32, originTokenAddresses []common.Address, sovereignTokenAddresses []common.Address, isNotMintable []bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.SetMultipleSovereignTokenAddress(&_Bridgel2sovereignchainv1010.TransactOpts, originNetworks, originTokenAddresses, sovereignTokenAddresses, isNotMintable)
}

// SetSovereignWETHAddress is a paid mutator transaction binding the contract method 0xbf130d7f.
//
// Solidity: function setSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Transactor) SetSovereignWETHAddress(opts *bind.TransactOpts, sovereignWETHTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.contract.Transact(opts, "setSovereignWETHAddress", sovereignWETHTokenAddress, isNotMintable)
}

// SetSovereignWETHAddress is a paid mutator transaction binding the contract method 0xbf130d7f.
//
// Solidity: function setSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) SetSovereignWETHAddress(sovereignWETHTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.SetSovereignWETHAddress(&_Bridgel2sovereignchainv1010.TransactOpts, sovereignWETHTokenAddress, isNotMintable)
}

// SetSovereignWETHAddress is a paid mutator transaction binding the contract method 0xbf130d7f.
//
// Solidity: function setSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010TransactorSession) SetSovereignWETHAddress(sovereignWETHTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.SetSovereignWETHAddress(&_Bridgel2sovereignchainv1010.TransactOpts, sovereignWETHTokenAddress, isNotMintable)
}

// TransferEmergencyBridgePauserRole is a paid mutator transaction binding the contract method 0x69e3ab12.
//
// Solidity: function transferEmergencyBridgePauserRole(address newEmergencyBridgePauser) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Transactor) TransferEmergencyBridgePauserRole(opts *bind.TransactOpts, newEmergencyBridgePauser common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.contract.Transact(opts, "transferEmergencyBridgePauserRole", newEmergencyBridgePauser)
}

// TransferEmergencyBridgePauserRole is a paid mutator transaction binding the contract method 0x69e3ab12.
//
// Solidity: function transferEmergencyBridgePauserRole(address newEmergencyBridgePauser) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) TransferEmergencyBridgePauserRole(newEmergencyBridgePauser common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.TransferEmergencyBridgePauserRole(&_Bridgel2sovereignchainv1010.TransactOpts, newEmergencyBridgePauser)
}

// TransferEmergencyBridgePauserRole is a paid mutator transaction binding the contract method 0x69e3ab12.
//
// Solidity: function transferEmergencyBridgePauserRole(address newEmergencyBridgePauser) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010TransactorSession) TransferEmergencyBridgePauserRole(newEmergencyBridgePauser common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.TransferEmergencyBridgePauserRole(&_Bridgel2sovereignchainv1010.TransactOpts, newEmergencyBridgePauser)
}

// TransferEmergencyBridgeUnpauserRole is a paid mutator transaction binding the contract method 0x8d942096.
//
// Solidity: function transferEmergencyBridgeUnpauserRole(address newEmergencyBridgeUnpauser) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Transactor) TransferEmergencyBridgeUnpauserRole(opts *bind.TransactOpts, newEmergencyBridgeUnpauser common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.contract.Transact(opts, "transferEmergencyBridgeUnpauserRole", newEmergencyBridgeUnpauser)
}

// TransferEmergencyBridgeUnpauserRole is a paid mutator transaction binding the contract method 0x8d942096.
//
// Solidity: function transferEmergencyBridgeUnpauserRole(address newEmergencyBridgeUnpauser) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) TransferEmergencyBridgeUnpauserRole(newEmergencyBridgeUnpauser common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.TransferEmergencyBridgeUnpauserRole(&_Bridgel2sovereignchainv1010.TransactOpts, newEmergencyBridgeUnpauser)
}

// TransferEmergencyBridgeUnpauserRole is a paid mutator transaction binding the contract method 0x8d942096.
//
// Solidity: function transferEmergencyBridgeUnpauserRole(address newEmergencyBridgeUnpauser) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010TransactorSession) TransferEmergencyBridgeUnpauserRole(newEmergencyBridgeUnpauser common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.TransferEmergencyBridgeUnpauserRole(&_Bridgel2sovereignchainv1010.TransactOpts, newEmergencyBridgeUnpauser)
}

// TransferProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8bd309c3.
//
// Solidity: function transferProxiedTokensManagerRole(address newProxiedTokensManager) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Transactor) TransferProxiedTokensManagerRole(opts *bind.TransactOpts, newProxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.contract.Transact(opts, "transferProxiedTokensManagerRole", newProxiedTokensManager)
}

// TransferProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8bd309c3.
//
// Solidity: function transferProxiedTokensManagerRole(address newProxiedTokensManager) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) TransferProxiedTokensManagerRole(newProxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.TransferProxiedTokensManagerRole(&_Bridgel2sovereignchainv1010.TransactOpts, newProxiedTokensManager)
}

// TransferProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8bd309c3.
//
// Solidity: function transferProxiedTokensManagerRole(address newProxiedTokensManager) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010TransactorSession) TransferProxiedTokensManagerRole(newProxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.TransferProxiedTokensManagerRole(&_Bridgel2sovereignchainv1010.TransactOpts, newProxiedTokensManager)
}

// UnsetMultipleClaims is a paid mutator transaction binding the contract method 0x136a2c60.
//
// Solidity: function unsetMultipleClaims(uint256[] globalIndexes) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Transactor) UnsetMultipleClaims(opts *bind.TransactOpts, globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.contract.Transact(opts, "unsetMultipleClaims", globalIndexes)
}

// UnsetMultipleClaims is a paid mutator transaction binding the contract method 0x136a2c60.
//
// Solidity: function unsetMultipleClaims(uint256[] globalIndexes) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) UnsetMultipleClaims(globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.UnsetMultipleClaims(&_Bridgel2sovereignchainv1010.TransactOpts, globalIndexes)
}

// UnsetMultipleClaims is a paid mutator transaction binding the contract method 0x136a2c60.
//
// Solidity: function unsetMultipleClaims(uint256[] globalIndexes) returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010TransactorSession) UnsetMultipleClaims(globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.UnsetMultipleClaims(&_Bridgel2sovereignchainv1010.TransactOpts, globalIndexes)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Transactor) UpdateGlobalExitRoot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.contract.Transact(opts, "updateGlobalExitRoot")
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Session) UpdateGlobalExitRoot() (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.UpdateGlobalExitRoot(&_Bridgel2sovereignchainv1010.TransactOpts)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010TransactorSession) UpdateGlobalExitRoot() (*types.Transaction, error) {
	return _Bridgel2sovereignchainv1010.Contract.UpdateGlobalExitRoot(&_Bridgel2sovereignchainv1010.TransactOpts)
}

// Bridgel2sovereignchainv1010AcceptEmergencyBridgePauserRoleIterator is returned from FilterAcceptEmergencyBridgePauserRole and is used to iterate over the raw logs and unpacked data for AcceptEmergencyBridgePauserRole events raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010AcceptEmergencyBridgePauserRoleIterator struct {
	Event *Bridgel2sovereignchainv1010AcceptEmergencyBridgePauserRole // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainv1010AcceptEmergencyBridgePauserRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainv1010AcceptEmergencyBridgePauserRole)
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
		it.Event = new(Bridgel2sovereignchainv1010AcceptEmergencyBridgePauserRole)
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
func (it *Bridgel2sovereignchainv1010AcceptEmergencyBridgePauserRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainv1010AcceptEmergencyBridgePauserRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainv1010AcceptEmergencyBridgePauserRole represents a AcceptEmergencyBridgePauserRole event raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010AcceptEmergencyBridgePauserRole struct {
	OldEmergencyBridgePauser common.Address
	NewEmergencyBridgePauser common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterAcceptEmergencyBridgePauserRole is a free log retrieval operation binding the contract event 0x85d2bdfbe58cd81abf8199c13ce2509204be4aba8603b9d29f52c4e13e7bb793.
//
// Solidity: event AcceptEmergencyBridgePauserRole(address oldEmergencyBridgePauser, address newEmergencyBridgePauser)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) FilterAcceptEmergencyBridgePauserRole(opts *bind.FilterOpts) (*Bridgel2sovereignchainv1010AcceptEmergencyBridgePauserRoleIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.FilterLogs(opts, "AcceptEmergencyBridgePauserRole")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainv1010AcceptEmergencyBridgePauserRoleIterator{contract: _Bridgel2sovereignchainv1010.contract, event: "AcceptEmergencyBridgePauserRole", logs: logs, sub: sub}, nil
}

// WatchAcceptEmergencyBridgePauserRole is a free log subscription operation binding the contract event 0x85d2bdfbe58cd81abf8199c13ce2509204be4aba8603b9d29f52c4e13e7bb793.
//
// Solidity: event AcceptEmergencyBridgePauserRole(address oldEmergencyBridgePauser, address newEmergencyBridgePauser)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) WatchAcceptEmergencyBridgePauserRole(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainv1010AcceptEmergencyBridgePauserRole) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.WatchLogs(opts, "AcceptEmergencyBridgePauserRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainv1010AcceptEmergencyBridgePauserRole)
				if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "AcceptEmergencyBridgePauserRole", log); err != nil {
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

// ParseAcceptEmergencyBridgePauserRole is a log parse operation binding the contract event 0x85d2bdfbe58cd81abf8199c13ce2509204be4aba8603b9d29f52c4e13e7bb793.
//
// Solidity: event AcceptEmergencyBridgePauserRole(address oldEmergencyBridgePauser, address newEmergencyBridgePauser)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) ParseAcceptEmergencyBridgePauserRole(log types.Log) (*Bridgel2sovereignchainv1010AcceptEmergencyBridgePauserRole, error) {
	event := new(Bridgel2sovereignchainv1010AcceptEmergencyBridgePauserRole)
	if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "AcceptEmergencyBridgePauserRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainv1010AcceptEmergencyBridgeUnpauserRoleIterator is returned from FilterAcceptEmergencyBridgeUnpauserRole and is used to iterate over the raw logs and unpacked data for AcceptEmergencyBridgeUnpauserRole events raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010AcceptEmergencyBridgeUnpauserRoleIterator struct {
	Event *Bridgel2sovereignchainv1010AcceptEmergencyBridgeUnpauserRole // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainv1010AcceptEmergencyBridgeUnpauserRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainv1010AcceptEmergencyBridgeUnpauserRole)
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
		it.Event = new(Bridgel2sovereignchainv1010AcceptEmergencyBridgeUnpauserRole)
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
func (it *Bridgel2sovereignchainv1010AcceptEmergencyBridgeUnpauserRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainv1010AcceptEmergencyBridgeUnpauserRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainv1010AcceptEmergencyBridgeUnpauserRole represents a AcceptEmergencyBridgeUnpauserRole event raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010AcceptEmergencyBridgeUnpauserRole struct {
	OldEmergencyBridgeUnpauser common.Address
	NewEmergencyBridgeUnpauser common.Address
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterAcceptEmergencyBridgeUnpauserRole is a free log retrieval operation binding the contract event 0x24cc8295aa5110cc216695db944ad2458c7795c6404449be980c3ce14aed752d.
//
// Solidity: event AcceptEmergencyBridgeUnpauserRole(address oldEmergencyBridgeUnpauser, address newEmergencyBridgeUnpauser)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) FilterAcceptEmergencyBridgeUnpauserRole(opts *bind.FilterOpts) (*Bridgel2sovereignchainv1010AcceptEmergencyBridgeUnpauserRoleIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.FilterLogs(opts, "AcceptEmergencyBridgeUnpauserRole")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainv1010AcceptEmergencyBridgeUnpauserRoleIterator{contract: _Bridgel2sovereignchainv1010.contract, event: "AcceptEmergencyBridgeUnpauserRole", logs: logs, sub: sub}, nil
}

// WatchAcceptEmergencyBridgeUnpauserRole is a free log subscription operation binding the contract event 0x24cc8295aa5110cc216695db944ad2458c7795c6404449be980c3ce14aed752d.
//
// Solidity: event AcceptEmergencyBridgeUnpauserRole(address oldEmergencyBridgeUnpauser, address newEmergencyBridgeUnpauser)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) WatchAcceptEmergencyBridgeUnpauserRole(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainv1010AcceptEmergencyBridgeUnpauserRole) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.WatchLogs(opts, "AcceptEmergencyBridgeUnpauserRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainv1010AcceptEmergencyBridgeUnpauserRole)
				if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "AcceptEmergencyBridgeUnpauserRole", log); err != nil {
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

// ParseAcceptEmergencyBridgeUnpauserRole is a log parse operation binding the contract event 0x24cc8295aa5110cc216695db944ad2458c7795c6404449be980c3ce14aed752d.
//
// Solidity: event AcceptEmergencyBridgeUnpauserRole(address oldEmergencyBridgeUnpauser, address newEmergencyBridgeUnpauser)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) ParseAcceptEmergencyBridgeUnpauserRole(log types.Log) (*Bridgel2sovereignchainv1010AcceptEmergencyBridgeUnpauserRole, error) {
	event := new(Bridgel2sovereignchainv1010AcceptEmergencyBridgeUnpauserRole)
	if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "AcceptEmergencyBridgeUnpauserRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainv1010AcceptProxiedTokensManagerRoleIterator is returned from FilterAcceptProxiedTokensManagerRole and is used to iterate over the raw logs and unpacked data for AcceptProxiedTokensManagerRole events raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010AcceptProxiedTokensManagerRoleIterator struct {
	Event *Bridgel2sovereignchainv1010AcceptProxiedTokensManagerRole // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainv1010AcceptProxiedTokensManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainv1010AcceptProxiedTokensManagerRole)
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
		it.Event = new(Bridgel2sovereignchainv1010AcceptProxiedTokensManagerRole)
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
func (it *Bridgel2sovereignchainv1010AcceptProxiedTokensManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainv1010AcceptProxiedTokensManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainv1010AcceptProxiedTokensManagerRole represents a AcceptProxiedTokensManagerRole event raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010AcceptProxiedTokensManagerRole struct {
	OldProxiedTokensManager common.Address
	NewProxiedTokensManager common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterAcceptProxiedTokensManagerRole is a free log retrieval operation binding the contract event 0xa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f2.
//
// Solidity: event AcceptProxiedTokensManagerRole(address oldProxiedTokensManager, address newProxiedTokensManager)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) FilterAcceptProxiedTokensManagerRole(opts *bind.FilterOpts) (*Bridgel2sovereignchainv1010AcceptProxiedTokensManagerRoleIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.FilterLogs(opts, "AcceptProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainv1010AcceptProxiedTokensManagerRoleIterator{contract: _Bridgel2sovereignchainv1010.contract, event: "AcceptProxiedTokensManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptProxiedTokensManagerRole is a free log subscription operation binding the contract event 0xa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f2.
//
// Solidity: event AcceptProxiedTokensManagerRole(address oldProxiedTokensManager, address newProxiedTokensManager)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) WatchAcceptProxiedTokensManagerRole(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainv1010AcceptProxiedTokensManagerRole) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.WatchLogs(opts, "AcceptProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainv1010AcceptProxiedTokensManagerRole)
				if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "AcceptProxiedTokensManagerRole", log); err != nil {
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
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) ParseAcceptProxiedTokensManagerRole(log types.Log) (*Bridgel2sovereignchainv1010AcceptProxiedTokensManagerRole, error) {
	event := new(Bridgel2sovereignchainv1010AcceptProxiedTokensManagerRole)
	if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "AcceptProxiedTokensManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainv1010BridgeEventIterator is returned from FilterBridgeEvent and is used to iterate over the raw logs and unpacked data for BridgeEvent events raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010BridgeEventIterator struct {
	Event *Bridgel2sovereignchainv1010BridgeEvent // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainv1010BridgeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainv1010BridgeEvent)
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
		it.Event = new(Bridgel2sovereignchainv1010BridgeEvent)
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
func (it *Bridgel2sovereignchainv1010BridgeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainv1010BridgeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainv1010BridgeEvent represents a BridgeEvent event raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010BridgeEvent struct {
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
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) FilterBridgeEvent(opts *bind.FilterOpts) (*Bridgel2sovereignchainv1010BridgeEventIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.FilterLogs(opts, "BridgeEvent")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainv1010BridgeEventIterator{contract: _Bridgel2sovereignchainv1010.contract, event: "BridgeEvent", logs: logs, sub: sub}, nil
}

// WatchBridgeEvent is a free log subscription operation binding the contract event 0x501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b.
//
// Solidity: event BridgeEvent(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata, uint32 depositCount)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) WatchBridgeEvent(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainv1010BridgeEvent) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.WatchLogs(opts, "BridgeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainv1010BridgeEvent)
				if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "BridgeEvent", log); err != nil {
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
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) ParseBridgeEvent(log types.Log) (*Bridgel2sovereignchainv1010BridgeEvent, error) {
	event := new(Bridgel2sovereignchainv1010BridgeEvent)
	if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "BridgeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainv1010ClaimEventIterator is returned from FilterClaimEvent and is used to iterate over the raw logs and unpacked data for ClaimEvent events raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010ClaimEventIterator struct {
	Event *Bridgel2sovereignchainv1010ClaimEvent // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainv1010ClaimEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainv1010ClaimEvent)
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
		it.Event = new(Bridgel2sovereignchainv1010ClaimEvent)
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
func (it *Bridgel2sovereignchainv1010ClaimEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainv1010ClaimEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainv1010ClaimEvent represents a ClaimEvent event raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010ClaimEvent struct {
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
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) FilterClaimEvent(opts *bind.FilterOpts) (*Bridgel2sovereignchainv1010ClaimEventIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.FilterLogs(opts, "ClaimEvent")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainv1010ClaimEventIterator{contract: _Bridgel2sovereignchainv1010.contract, event: "ClaimEvent", logs: logs, sub: sub}, nil
}

// WatchClaimEvent is a free log subscription operation binding the contract event 0x1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d.
//
// Solidity: event ClaimEvent(uint256 globalIndex, uint32 originNetwork, address originAddress, address destinationAddress, uint256 amount)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) WatchClaimEvent(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainv1010ClaimEvent) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.WatchLogs(opts, "ClaimEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainv1010ClaimEvent)
				if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "ClaimEvent", log); err != nil {
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
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) ParseClaimEvent(log types.Log) (*Bridgel2sovereignchainv1010ClaimEvent, error) {
	event := new(Bridgel2sovereignchainv1010ClaimEvent)
	if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "ClaimEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainv1010EmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010EmergencyStateActivatedIterator struct {
	Event *Bridgel2sovereignchainv1010EmergencyStateActivated // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainv1010EmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainv1010EmergencyStateActivated)
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
		it.Event = new(Bridgel2sovereignchainv1010EmergencyStateActivated)
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
func (it *Bridgel2sovereignchainv1010EmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainv1010EmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainv1010EmergencyStateActivated represents a EmergencyStateActivated event raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010EmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*Bridgel2sovereignchainv1010EmergencyStateActivatedIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainv1010EmergencyStateActivatedIterator{contract: _Bridgel2sovereignchainv1010.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainv1010EmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainv1010EmergencyStateActivated)
				if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
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
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) ParseEmergencyStateActivated(log types.Log) (*Bridgel2sovereignchainv1010EmergencyStateActivated, error) {
	event := new(Bridgel2sovereignchainv1010EmergencyStateActivated)
	if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainv1010EmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010EmergencyStateDeactivatedIterator struct {
	Event *Bridgel2sovereignchainv1010EmergencyStateDeactivated // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainv1010EmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainv1010EmergencyStateDeactivated)
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
		it.Event = new(Bridgel2sovereignchainv1010EmergencyStateDeactivated)
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
func (it *Bridgel2sovereignchainv1010EmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainv1010EmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainv1010EmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010EmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*Bridgel2sovereignchainv1010EmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainv1010EmergencyStateDeactivatedIterator{contract: _Bridgel2sovereignchainv1010.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainv1010EmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainv1010EmergencyStateDeactivated)
				if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
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
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) ParseEmergencyStateDeactivated(log types.Log) (*Bridgel2sovereignchainv1010EmergencyStateDeactivated, error) {
	event := new(Bridgel2sovereignchainv1010EmergencyStateDeactivated)
	if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainv1010InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010InitializedIterator struct {
	Event *Bridgel2sovereignchainv1010Initialized // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainv1010InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainv1010Initialized)
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
		it.Event = new(Bridgel2sovereignchainv1010Initialized)
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
func (it *Bridgel2sovereignchainv1010InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainv1010InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainv1010Initialized represents a Initialized event raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) FilterInitialized(opts *bind.FilterOpts) (*Bridgel2sovereignchainv1010InitializedIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainv1010InitializedIterator{contract: _Bridgel2sovereignchainv1010.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainv1010Initialized) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainv1010Initialized)
				if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) ParseInitialized(log types.Log) (*Bridgel2sovereignchainv1010Initialized, error) {
	event := new(Bridgel2sovereignchainv1010Initialized)
	if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainv1010MigrateLegacyTokenIterator is returned from FilterMigrateLegacyToken and is used to iterate over the raw logs and unpacked data for MigrateLegacyToken events raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010MigrateLegacyTokenIterator struct {
	Event *Bridgel2sovereignchainv1010MigrateLegacyToken // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainv1010MigrateLegacyTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainv1010MigrateLegacyToken)
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
		it.Event = new(Bridgel2sovereignchainv1010MigrateLegacyToken)
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
func (it *Bridgel2sovereignchainv1010MigrateLegacyTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainv1010MigrateLegacyTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainv1010MigrateLegacyToken represents a MigrateLegacyToken event raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010MigrateLegacyToken struct {
	Sender              common.Address
	LegacyTokenAddress  common.Address
	UpdatedTokenAddress common.Address
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMigrateLegacyToken is a free log retrieval operation binding the contract event 0xb7f8fd4d1faf9b2929dc269f59c53e3a2bccc44e9950f33a568fcbcb37eb69a9.
//
// Solidity: event MigrateLegacyToken(address sender, address legacyTokenAddress, address updatedTokenAddress, uint256 amount)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) FilterMigrateLegacyToken(opts *bind.FilterOpts) (*Bridgel2sovereignchainv1010MigrateLegacyTokenIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.FilterLogs(opts, "MigrateLegacyToken")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainv1010MigrateLegacyTokenIterator{contract: _Bridgel2sovereignchainv1010.contract, event: "MigrateLegacyToken", logs: logs, sub: sub}, nil
}

// WatchMigrateLegacyToken is a free log subscription operation binding the contract event 0xb7f8fd4d1faf9b2929dc269f59c53e3a2bccc44e9950f33a568fcbcb37eb69a9.
//
// Solidity: event MigrateLegacyToken(address sender, address legacyTokenAddress, address updatedTokenAddress, uint256 amount)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) WatchMigrateLegacyToken(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainv1010MigrateLegacyToken) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.WatchLogs(opts, "MigrateLegacyToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainv1010MigrateLegacyToken)
				if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "MigrateLegacyToken", log); err != nil {
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

// ParseMigrateLegacyToken is a log parse operation binding the contract event 0xb7f8fd4d1faf9b2929dc269f59c53e3a2bccc44e9950f33a568fcbcb37eb69a9.
//
// Solidity: event MigrateLegacyToken(address sender, address legacyTokenAddress, address updatedTokenAddress, uint256 amount)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) ParseMigrateLegacyToken(log types.Log) (*Bridgel2sovereignchainv1010MigrateLegacyToken, error) {
	event := new(Bridgel2sovereignchainv1010MigrateLegacyToken)
	if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "MigrateLegacyToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainv1010NewWrappedTokenIterator is returned from FilterNewWrappedToken and is used to iterate over the raw logs and unpacked data for NewWrappedToken events raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010NewWrappedTokenIterator struct {
	Event *Bridgel2sovereignchainv1010NewWrappedToken // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainv1010NewWrappedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainv1010NewWrappedToken)
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
		it.Event = new(Bridgel2sovereignchainv1010NewWrappedToken)
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
func (it *Bridgel2sovereignchainv1010NewWrappedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainv1010NewWrappedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainv1010NewWrappedToken represents a NewWrappedToken event raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010NewWrappedToken struct {
	OriginNetwork       uint32
	OriginTokenAddress  common.Address
	WrappedTokenAddress common.Address
	Metadata            []byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewWrappedToken is a free log retrieval operation binding the contract event 0x490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a39.
//
// Solidity: event NewWrappedToken(uint32 originNetwork, address originTokenAddress, address wrappedTokenAddress, bytes metadata)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) FilterNewWrappedToken(opts *bind.FilterOpts) (*Bridgel2sovereignchainv1010NewWrappedTokenIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.FilterLogs(opts, "NewWrappedToken")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainv1010NewWrappedTokenIterator{contract: _Bridgel2sovereignchainv1010.contract, event: "NewWrappedToken", logs: logs, sub: sub}, nil
}

// WatchNewWrappedToken is a free log subscription operation binding the contract event 0x490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a39.
//
// Solidity: event NewWrappedToken(uint32 originNetwork, address originTokenAddress, address wrappedTokenAddress, bytes metadata)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) WatchNewWrappedToken(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainv1010NewWrappedToken) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.WatchLogs(opts, "NewWrappedToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainv1010NewWrappedToken)
				if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "NewWrappedToken", log); err != nil {
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
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) ParseNewWrappedToken(log types.Log) (*Bridgel2sovereignchainv1010NewWrappedToken, error) {
	event := new(Bridgel2sovereignchainv1010NewWrappedToken)
	if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "NewWrappedToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainv1010RemoveLegacySovereignTokenAddressIterator is returned from FilterRemoveLegacySovereignTokenAddress and is used to iterate over the raw logs and unpacked data for RemoveLegacySovereignTokenAddress events raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010RemoveLegacySovereignTokenAddressIterator struct {
	Event *Bridgel2sovereignchainv1010RemoveLegacySovereignTokenAddress // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainv1010RemoveLegacySovereignTokenAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainv1010RemoveLegacySovereignTokenAddress)
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
		it.Event = new(Bridgel2sovereignchainv1010RemoveLegacySovereignTokenAddress)
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
func (it *Bridgel2sovereignchainv1010RemoveLegacySovereignTokenAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainv1010RemoveLegacySovereignTokenAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainv1010RemoveLegacySovereignTokenAddress represents a RemoveLegacySovereignTokenAddress event raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010RemoveLegacySovereignTokenAddress struct {
	SovereignTokenAddress common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRemoveLegacySovereignTokenAddress is a free log retrieval operation binding the contract event 0xc2ae0bd0ec0fd0352bfe5bacac49637af342c1e40f1b80a7f74440dc7fe3f063.
//
// Solidity: event RemoveLegacySovereignTokenAddress(address sovereignTokenAddress)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) FilterRemoveLegacySovereignTokenAddress(opts *bind.FilterOpts) (*Bridgel2sovereignchainv1010RemoveLegacySovereignTokenAddressIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.FilterLogs(opts, "RemoveLegacySovereignTokenAddress")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainv1010RemoveLegacySovereignTokenAddressIterator{contract: _Bridgel2sovereignchainv1010.contract, event: "RemoveLegacySovereignTokenAddress", logs: logs, sub: sub}, nil
}

// WatchRemoveLegacySovereignTokenAddress is a free log subscription operation binding the contract event 0xc2ae0bd0ec0fd0352bfe5bacac49637af342c1e40f1b80a7f74440dc7fe3f063.
//
// Solidity: event RemoveLegacySovereignTokenAddress(address sovereignTokenAddress)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) WatchRemoveLegacySovereignTokenAddress(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainv1010RemoveLegacySovereignTokenAddress) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.WatchLogs(opts, "RemoveLegacySovereignTokenAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainv1010RemoveLegacySovereignTokenAddress)
				if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "RemoveLegacySovereignTokenAddress", log); err != nil {
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

// ParseRemoveLegacySovereignTokenAddress is a log parse operation binding the contract event 0xc2ae0bd0ec0fd0352bfe5bacac49637af342c1e40f1b80a7f74440dc7fe3f063.
//
// Solidity: event RemoveLegacySovereignTokenAddress(address sovereignTokenAddress)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) ParseRemoveLegacySovereignTokenAddress(log types.Log) (*Bridgel2sovereignchainv1010RemoveLegacySovereignTokenAddress, error) {
	event := new(Bridgel2sovereignchainv1010RemoveLegacySovereignTokenAddress)
	if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "RemoveLegacySovereignTokenAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainv1010SetBridgeManagerIterator is returned from FilterSetBridgeManager and is used to iterate over the raw logs and unpacked data for SetBridgeManager events raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010SetBridgeManagerIterator struct {
	Event *Bridgel2sovereignchainv1010SetBridgeManager // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainv1010SetBridgeManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainv1010SetBridgeManager)
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
		it.Event = new(Bridgel2sovereignchainv1010SetBridgeManager)
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
func (it *Bridgel2sovereignchainv1010SetBridgeManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainv1010SetBridgeManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainv1010SetBridgeManager represents a SetBridgeManager event raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010SetBridgeManager struct {
	BridgeManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetBridgeManager is a free log retrieval operation binding the contract event 0x32cf74f8a6d5f88593984d2cd52be5592bfa6884f5896175801a5069ef09cd67.
//
// Solidity: event SetBridgeManager(address bridgeManager)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) FilterSetBridgeManager(opts *bind.FilterOpts) (*Bridgel2sovereignchainv1010SetBridgeManagerIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.FilterLogs(opts, "SetBridgeManager")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainv1010SetBridgeManagerIterator{contract: _Bridgel2sovereignchainv1010.contract, event: "SetBridgeManager", logs: logs, sub: sub}, nil
}

// WatchSetBridgeManager is a free log subscription operation binding the contract event 0x32cf74f8a6d5f88593984d2cd52be5592bfa6884f5896175801a5069ef09cd67.
//
// Solidity: event SetBridgeManager(address bridgeManager)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) WatchSetBridgeManager(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainv1010SetBridgeManager) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.WatchLogs(opts, "SetBridgeManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainv1010SetBridgeManager)
				if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "SetBridgeManager", log); err != nil {
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

// ParseSetBridgeManager is a log parse operation binding the contract event 0x32cf74f8a6d5f88593984d2cd52be5592bfa6884f5896175801a5069ef09cd67.
//
// Solidity: event SetBridgeManager(address bridgeManager)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) ParseSetBridgeManager(log types.Log) (*Bridgel2sovereignchainv1010SetBridgeManager, error) {
	event := new(Bridgel2sovereignchainv1010SetBridgeManager)
	if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "SetBridgeManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainv1010SetInitialLocalBalanceTreeAmountIterator is returned from FilterSetInitialLocalBalanceTreeAmount and is used to iterate over the raw logs and unpacked data for SetInitialLocalBalanceTreeAmount events raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010SetInitialLocalBalanceTreeAmountIterator struct {
	Event *Bridgel2sovereignchainv1010SetInitialLocalBalanceTreeAmount // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainv1010SetInitialLocalBalanceTreeAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainv1010SetInitialLocalBalanceTreeAmount)
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
		it.Event = new(Bridgel2sovereignchainv1010SetInitialLocalBalanceTreeAmount)
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
func (it *Bridgel2sovereignchainv1010SetInitialLocalBalanceTreeAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainv1010SetInitialLocalBalanceTreeAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainv1010SetInitialLocalBalanceTreeAmount represents a SetInitialLocalBalanceTreeAmount event raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010SetInitialLocalBalanceTreeAmount struct {
	TokenInfoHash [32]byte
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetInitialLocalBalanceTreeAmount is a free log retrieval operation binding the contract event 0x2277ec68451dc01bd131765a9858d6de94d7e11220704d8ac1718fdb8de07cb2.
//
// Solidity: event SetInitialLocalBalanceTreeAmount(bytes32 tokenInfoHash, uint256 amount)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) FilterSetInitialLocalBalanceTreeAmount(opts *bind.FilterOpts) (*Bridgel2sovereignchainv1010SetInitialLocalBalanceTreeAmountIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.FilterLogs(opts, "SetInitialLocalBalanceTreeAmount")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainv1010SetInitialLocalBalanceTreeAmountIterator{contract: _Bridgel2sovereignchainv1010.contract, event: "SetInitialLocalBalanceTreeAmount", logs: logs, sub: sub}, nil
}

// WatchSetInitialLocalBalanceTreeAmount is a free log subscription operation binding the contract event 0x2277ec68451dc01bd131765a9858d6de94d7e11220704d8ac1718fdb8de07cb2.
//
// Solidity: event SetInitialLocalBalanceTreeAmount(bytes32 tokenInfoHash, uint256 amount)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) WatchSetInitialLocalBalanceTreeAmount(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainv1010SetInitialLocalBalanceTreeAmount) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.WatchLogs(opts, "SetInitialLocalBalanceTreeAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainv1010SetInitialLocalBalanceTreeAmount)
				if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "SetInitialLocalBalanceTreeAmount", log); err != nil {
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

// ParseSetInitialLocalBalanceTreeAmount is a log parse operation binding the contract event 0x2277ec68451dc01bd131765a9858d6de94d7e11220704d8ac1718fdb8de07cb2.
//
// Solidity: event SetInitialLocalBalanceTreeAmount(bytes32 tokenInfoHash, uint256 amount)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) ParseSetInitialLocalBalanceTreeAmount(log types.Log) (*Bridgel2sovereignchainv1010SetInitialLocalBalanceTreeAmount, error) {
	event := new(Bridgel2sovereignchainv1010SetInitialLocalBalanceTreeAmount)
	if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "SetInitialLocalBalanceTreeAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainv1010SetSovereignTokenAddressIterator is returned from FilterSetSovereignTokenAddress and is used to iterate over the raw logs and unpacked data for SetSovereignTokenAddress events raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010SetSovereignTokenAddressIterator struct {
	Event *Bridgel2sovereignchainv1010SetSovereignTokenAddress // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainv1010SetSovereignTokenAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainv1010SetSovereignTokenAddress)
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
		it.Event = new(Bridgel2sovereignchainv1010SetSovereignTokenAddress)
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
func (it *Bridgel2sovereignchainv1010SetSovereignTokenAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainv1010SetSovereignTokenAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainv1010SetSovereignTokenAddress represents a SetSovereignTokenAddress event raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010SetSovereignTokenAddress struct {
	OriginNetwork         uint32
	OriginTokenAddress    common.Address
	SovereignTokenAddress common.Address
	IsNotMintable         bool
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetSovereignTokenAddress is a free log retrieval operation binding the contract event 0xdbe8a5da6a7a916d9adfda9160167a0f8a3da415ee6610e810e753853597fce7.
//
// Solidity: event SetSovereignTokenAddress(uint32 originNetwork, address originTokenAddress, address sovereignTokenAddress, bool isNotMintable)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) FilterSetSovereignTokenAddress(opts *bind.FilterOpts) (*Bridgel2sovereignchainv1010SetSovereignTokenAddressIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.FilterLogs(opts, "SetSovereignTokenAddress")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainv1010SetSovereignTokenAddressIterator{contract: _Bridgel2sovereignchainv1010.contract, event: "SetSovereignTokenAddress", logs: logs, sub: sub}, nil
}

// WatchSetSovereignTokenAddress is a free log subscription operation binding the contract event 0xdbe8a5da6a7a916d9adfda9160167a0f8a3da415ee6610e810e753853597fce7.
//
// Solidity: event SetSovereignTokenAddress(uint32 originNetwork, address originTokenAddress, address sovereignTokenAddress, bool isNotMintable)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) WatchSetSovereignTokenAddress(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainv1010SetSovereignTokenAddress) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.WatchLogs(opts, "SetSovereignTokenAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainv1010SetSovereignTokenAddress)
				if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "SetSovereignTokenAddress", log); err != nil {
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

// ParseSetSovereignTokenAddress is a log parse operation binding the contract event 0xdbe8a5da6a7a916d9adfda9160167a0f8a3da415ee6610e810e753853597fce7.
//
// Solidity: event SetSovereignTokenAddress(uint32 originNetwork, address originTokenAddress, address sovereignTokenAddress, bool isNotMintable)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) ParseSetSovereignTokenAddress(log types.Log) (*Bridgel2sovereignchainv1010SetSovereignTokenAddress, error) {
	event := new(Bridgel2sovereignchainv1010SetSovereignTokenAddress)
	if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "SetSovereignTokenAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainv1010SetSovereignWETHAddressIterator is returned from FilterSetSovereignWETHAddress and is used to iterate over the raw logs and unpacked data for SetSovereignWETHAddress events raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010SetSovereignWETHAddressIterator struct {
	Event *Bridgel2sovereignchainv1010SetSovereignWETHAddress // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainv1010SetSovereignWETHAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainv1010SetSovereignWETHAddress)
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
		it.Event = new(Bridgel2sovereignchainv1010SetSovereignWETHAddress)
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
func (it *Bridgel2sovereignchainv1010SetSovereignWETHAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainv1010SetSovereignWETHAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainv1010SetSovereignWETHAddress represents a SetSovereignWETHAddress event raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010SetSovereignWETHAddress struct {
	SovereignWETHTokenAddress common.Address
	IsNotMintable             bool
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterSetSovereignWETHAddress is a free log retrieval operation binding the contract event 0xc7318b7ed6ba4f2908a3de396d8ab49b1dadb55db5b55123247a401f29ff8d82.
//
// Solidity: event SetSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) FilterSetSovereignWETHAddress(opts *bind.FilterOpts) (*Bridgel2sovereignchainv1010SetSovereignWETHAddressIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.FilterLogs(opts, "SetSovereignWETHAddress")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainv1010SetSovereignWETHAddressIterator{contract: _Bridgel2sovereignchainv1010.contract, event: "SetSovereignWETHAddress", logs: logs, sub: sub}, nil
}

// WatchSetSovereignWETHAddress is a free log subscription operation binding the contract event 0xc7318b7ed6ba4f2908a3de396d8ab49b1dadb55db5b55123247a401f29ff8d82.
//
// Solidity: event SetSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) WatchSetSovereignWETHAddress(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainv1010SetSovereignWETHAddress) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.WatchLogs(opts, "SetSovereignWETHAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainv1010SetSovereignWETHAddress)
				if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "SetSovereignWETHAddress", log); err != nil {
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

// ParseSetSovereignWETHAddress is a log parse operation binding the contract event 0xc7318b7ed6ba4f2908a3de396d8ab49b1dadb55db5b55123247a401f29ff8d82.
//
// Solidity: event SetSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) ParseSetSovereignWETHAddress(log types.Log) (*Bridgel2sovereignchainv1010SetSovereignWETHAddress, error) {
	event := new(Bridgel2sovereignchainv1010SetSovereignWETHAddress)
	if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "SetSovereignWETHAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainv1010TransferEmergencyBridgePauserRoleIterator is returned from FilterTransferEmergencyBridgePauserRole and is used to iterate over the raw logs and unpacked data for TransferEmergencyBridgePauserRole events raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010TransferEmergencyBridgePauserRoleIterator struct {
	Event *Bridgel2sovereignchainv1010TransferEmergencyBridgePauserRole // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainv1010TransferEmergencyBridgePauserRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainv1010TransferEmergencyBridgePauserRole)
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
		it.Event = new(Bridgel2sovereignchainv1010TransferEmergencyBridgePauserRole)
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
func (it *Bridgel2sovereignchainv1010TransferEmergencyBridgePauserRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainv1010TransferEmergencyBridgePauserRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainv1010TransferEmergencyBridgePauserRole represents a TransferEmergencyBridgePauserRole event raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010TransferEmergencyBridgePauserRole struct {
	CurrentEmergencyBridgePauser common.Address
	NewEmergencyBridgePauser     common.Address
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterTransferEmergencyBridgePauserRole is a free log retrieval operation binding the contract event 0xb27de219766f47b82684842855ba6130b6dbf288ac66d1c3509e7bf17f4e925a.
//
// Solidity: event TransferEmergencyBridgePauserRole(address currentEmergencyBridgePauser, address newEmergencyBridgePauser)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) FilterTransferEmergencyBridgePauserRole(opts *bind.FilterOpts) (*Bridgel2sovereignchainv1010TransferEmergencyBridgePauserRoleIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.FilterLogs(opts, "TransferEmergencyBridgePauserRole")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainv1010TransferEmergencyBridgePauserRoleIterator{contract: _Bridgel2sovereignchainv1010.contract, event: "TransferEmergencyBridgePauserRole", logs: logs, sub: sub}, nil
}

// WatchTransferEmergencyBridgePauserRole is a free log subscription operation binding the contract event 0xb27de219766f47b82684842855ba6130b6dbf288ac66d1c3509e7bf17f4e925a.
//
// Solidity: event TransferEmergencyBridgePauserRole(address currentEmergencyBridgePauser, address newEmergencyBridgePauser)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) WatchTransferEmergencyBridgePauserRole(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainv1010TransferEmergencyBridgePauserRole) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.WatchLogs(opts, "TransferEmergencyBridgePauserRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainv1010TransferEmergencyBridgePauserRole)
				if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "TransferEmergencyBridgePauserRole", log); err != nil {
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

// ParseTransferEmergencyBridgePauserRole is a log parse operation binding the contract event 0xb27de219766f47b82684842855ba6130b6dbf288ac66d1c3509e7bf17f4e925a.
//
// Solidity: event TransferEmergencyBridgePauserRole(address currentEmergencyBridgePauser, address newEmergencyBridgePauser)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) ParseTransferEmergencyBridgePauserRole(log types.Log) (*Bridgel2sovereignchainv1010TransferEmergencyBridgePauserRole, error) {
	event := new(Bridgel2sovereignchainv1010TransferEmergencyBridgePauserRole)
	if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "TransferEmergencyBridgePauserRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainv1010TransferEmergencyBridgeUnpauserRoleIterator is returned from FilterTransferEmergencyBridgeUnpauserRole and is used to iterate over the raw logs and unpacked data for TransferEmergencyBridgeUnpauserRole events raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010TransferEmergencyBridgeUnpauserRoleIterator struct {
	Event *Bridgel2sovereignchainv1010TransferEmergencyBridgeUnpauserRole // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainv1010TransferEmergencyBridgeUnpauserRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainv1010TransferEmergencyBridgeUnpauserRole)
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
		it.Event = new(Bridgel2sovereignchainv1010TransferEmergencyBridgeUnpauserRole)
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
func (it *Bridgel2sovereignchainv1010TransferEmergencyBridgeUnpauserRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainv1010TransferEmergencyBridgeUnpauserRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainv1010TransferEmergencyBridgeUnpauserRole represents a TransferEmergencyBridgeUnpauserRole event raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010TransferEmergencyBridgeUnpauserRole struct {
	CurrentEmergencyBridgeUnpauser common.Address
	NewEmergencyBridgeUnpauser     common.Address
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterTransferEmergencyBridgeUnpauserRole is a free log retrieval operation binding the contract event 0xf01a62a06940517bbc898dec8c75794b9feabcd2d263c8de823b36dbbeb8779b.
//
// Solidity: event TransferEmergencyBridgeUnpauserRole(address currentEmergencyBridgeUnpauser, address newEmergencyBridgeUnpauser)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) FilterTransferEmergencyBridgeUnpauserRole(opts *bind.FilterOpts) (*Bridgel2sovereignchainv1010TransferEmergencyBridgeUnpauserRoleIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.FilterLogs(opts, "TransferEmergencyBridgeUnpauserRole")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainv1010TransferEmergencyBridgeUnpauserRoleIterator{contract: _Bridgel2sovereignchainv1010.contract, event: "TransferEmergencyBridgeUnpauserRole", logs: logs, sub: sub}, nil
}

// WatchTransferEmergencyBridgeUnpauserRole is a free log subscription operation binding the contract event 0xf01a62a06940517bbc898dec8c75794b9feabcd2d263c8de823b36dbbeb8779b.
//
// Solidity: event TransferEmergencyBridgeUnpauserRole(address currentEmergencyBridgeUnpauser, address newEmergencyBridgeUnpauser)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) WatchTransferEmergencyBridgeUnpauserRole(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainv1010TransferEmergencyBridgeUnpauserRole) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.WatchLogs(opts, "TransferEmergencyBridgeUnpauserRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainv1010TransferEmergencyBridgeUnpauserRole)
				if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "TransferEmergencyBridgeUnpauserRole", log); err != nil {
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

// ParseTransferEmergencyBridgeUnpauserRole is a log parse operation binding the contract event 0xf01a62a06940517bbc898dec8c75794b9feabcd2d263c8de823b36dbbeb8779b.
//
// Solidity: event TransferEmergencyBridgeUnpauserRole(address currentEmergencyBridgeUnpauser, address newEmergencyBridgeUnpauser)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) ParseTransferEmergencyBridgeUnpauserRole(log types.Log) (*Bridgel2sovereignchainv1010TransferEmergencyBridgeUnpauserRole, error) {
	event := new(Bridgel2sovereignchainv1010TransferEmergencyBridgeUnpauserRole)
	if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "TransferEmergencyBridgeUnpauserRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainv1010TransferProxiedTokensManagerRoleIterator is returned from FilterTransferProxiedTokensManagerRole and is used to iterate over the raw logs and unpacked data for TransferProxiedTokensManagerRole events raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010TransferProxiedTokensManagerRoleIterator struct {
	Event *Bridgel2sovereignchainv1010TransferProxiedTokensManagerRole // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainv1010TransferProxiedTokensManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainv1010TransferProxiedTokensManagerRole)
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
		it.Event = new(Bridgel2sovereignchainv1010TransferProxiedTokensManagerRole)
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
func (it *Bridgel2sovereignchainv1010TransferProxiedTokensManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainv1010TransferProxiedTokensManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainv1010TransferProxiedTokensManagerRole represents a TransferProxiedTokensManagerRole event raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010TransferProxiedTokensManagerRole struct {
	CurrentProxiedTokensManager common.Address
	NewProxiedTokensManager     common.Address
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterTransferProxiedTokensManagerRole is a free log retrieval operation binding the contract event 0x0a34baa3feb299aef9c05cb59c6e0c8e7c0bcc65cbf0a647e7a7c8a2411591e2.
//
// Solidity: event TransferProxiedTokensManagerRole(address currentProxiedTokensManager, address newProxiedTokensManager)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) FilterTransferProxiedTokensManagerRole(opts *bind.FilterOpts) (*Bridgel2sovereignchainv1010TransferProxiedTokensManagerRoleIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.FilterLogs(opts, "TransferProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainv1010TransferProxiedTokensManagerRoleIterator{contract: _Bridgel2sovereignchainv1010.contract, event: "TransferProxiedTokensManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferProxiedTokensManagerRole is a free log subscription operation binding the contract event 0x0a34baa3feb299aef9c05cb59c6e0c8e7c0bcc65cbf0a647e7a7c8a2411591e2.
//
// Solidity: event TransferProxiedTokensManagerRole(address currentProxiedTokensManager, address newProxiedTokensManager)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) WatchTransferProxiedTokensManagerRole(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainv1010TransferProxiedTokensManagerRole) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.WatchLogs(opts, "TransferProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainv1010TransferProxiedTokensManagerRole)
				if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "TransferProxiedTokensManagerRole", log); err != nil {
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
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) ParseTransferProxiedTokensManagerRole(log types.Log) (*Bridgel2sovereignchainv1010TransferProxiedTokensManagerRole, error) {
	event := new(Bridgel2sovereignchainv1010TransferProxiedTokensManagerRole)
	if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "TransferProxiedTokensManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainv1010UpdatedClaimedGlobalIndexHashChainIterator is returned from FilterUpdatedClaimedGlobalIndexHashChain and is used to iterate over the raw logs and unpacked data for UpdatedClaimedGlobalIndexHashChain events raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010UpdatedClaimedGlobalIndexHashChainIterator struct {
	Event *Bridgel2sovereignchainv1010UpdatedClaimedGlobalIndexHashChain // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainv1010UpdatedClaimedGlobalIndexHashChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainv1010UpdatedClaimedGlobalIndexHashChain)
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
		it.Event = new(Bridgel2sovereignchainv1010UpdatedClaimedGlobalIndexHashChain)
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
func (it *Bridgel2sovereignchainv1010UpdatedClaimedGlobalIndexHashChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainv1010UpdatedClaimedGlobalIndexHashChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainv1010UpdatedClaimedGlobalIndexHashChain represents a UpdatedClaimedGlobalIndexHashChain event raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010UpdatedClaimedGlobalIndexHashChain struct {
	ClaimedGlobalIndex             [32]byte
	NewClaimedGlobalIndexHashChain [32]byte
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterUpdatedClaimedGlobalIndexHashChain is a free log retrieval operation binding the contract event 0x3e5936f910a78eb5181813a939c8d4c3e4d85f87943f659380d82ac6221b0e92.
//
// Solidity: event UpdatedClaimedGlobalIndexHashChain(bytes32 claimedGlobalIndex, bytes32 newClaimedGlobalIndexHashChain)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) FilterUpdatedClaimedGlobalIndexHashChain(opts *bind.FilterOpts) (*Bridgel2sovereignchainv1010UpdatedClaimedGlobalIndexHashChainIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.FilterLogs(opts, "UpdatedClaimedGlobalIndexHashChain")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainv1010UpdatedClaimedGlobalIndexHashChainIterator{contract: _Bridgel2sovereignchainv1010.contract, event: "UpdatedClaimedGlobalIndexHashChain", logs: logs, sub: sub}, nil
}

// WatchUpdatedClaimedGlobalIndexHashChain is a free log subscription operation binding the contract event 0x3e5936f910a78eb5181813a939c8d4c3e4d85f87943f659380d82ac6221b0e92.
//
// Solidity: event UpdatedClaimedGlobalIndexHashChain(bytes32 claimedGlobalIndex, bytes32 newClaimedGlobalIndexHashChain)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) WatchUpdatedClaimedGlobalIndexHashChain(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainv1010UpdatedClaimedGlobalIndexHashChain) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.WatchLogs(opts, "UpdatedClaimedGlobalIndexHashChain")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainv1010UpdatedClaimedGlobalIndexHashChain)
				if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "UpdatedClaimedGlobalIndexHashChain", log); err != nil {
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

// ParseUpdatedClaimedGlobalIndexHashChain is a log parse operation binding the contract event 0x3e5936f910a78eb5181813a939c8d4c3e4d85f87943f659380d82ac6221b0e92.
//
// Solidity: event UpdatedClaimedGlobalIndexHashChain(bytes32 claimedGlobalIndex, bytes32 newClaimedGlobalIndexHashChain)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) ParseUpdatedClaimedGlobalIndexHashChain(log types.Log) (*Bridgel2sovereignchainv1010UpdatedClaimedGlobalIndexHashChain, error) {
	event := new(Bridgel2sovereignchainv1010UpdatedClaimedGlobalIndexHashChain)
	if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "UpdatedClaimedGlobalIndexHashChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainv1010UpdatedUnsetGlobalIndexHashChainIterator is returned from FilterUpdatedUnsetGlobalIndexHashChain and is used to iterate over the raw logs and unpacked data for UpdatedUnsetGlobalIndexHashChain events raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010UpdatedUnsetGlobalIndexHashChainIterator struct {
	Event *Bridgel2sovereignchainv1010UpdatedUnsetGlobalIndexHashChain // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainv1010UpdatedUnsetGlobalIndexHashChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainv1010UpdatedUnsetGlobalIndexHashChain)
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
		it.Event = new(Bridgel2sovereignchainv1010UpdatedUnsetGlobalIndexHashChain)
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
func (it *Bridgel2sovereignchainv1010UpdatedUnsetGlobalIndexHashChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainv1010UpdatedUnsetGlobalIndexHashChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainv1010UpdatedUnsetGlobalIndexHashChain represents a UpdatedUnsetGlobalIndexHashChain event raised by the Bridgel2sovereignchainv1010 contract.
type Bridgel2sovereignchainv1010UpdatedUnsetGlobalIndexHashChain struct {
	UnsetGlobalIndex             [32]byte
	NewUnsetGlobalIndexHashChain [32]byte
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedUnsetGlobalIndexHashChain is a free log retrieval operation binding the contract event 0xc80e0aca446a59735359a7ae46124b57c47b892827642779bc6dafc84ba90b03.
//
// Solidity: event UpdatedUnsetGlobalIndexHashChain(bytes32 unsetGlobalIndex, bytes32 newUnsetGlobalIndexHashChain)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) FilterUpdatedUnsetGlobalIndexHashChain(opts *bind.FilterOpts) (*Bridgel2sovereignchainv1010UpdatedUnsetGlobalIndexHashChainIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.FilterLogs(opts, "UpdatedUnsetGlobalIndexHashChain")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainv1010UpdatedUnsetGlobalIndexHashChainIterator{contract: _Bridgel2sovereignchainv1010.contract, event: "UpdatedUnsetGlobalIndexHashChain", logs: logs, sub: sub}, nil
}

// WatchUpdatedUnsetGlobalIndexHashChain is a free log subscription operation binding the contract event 0xc80e0aca446a59735359a7ae46124b57c47b892827642779bc6dafc84ba90b03.
//
// Solidity: event UpdatedUnsetGlobalIndexHashChain(bytes32 unsetGlobalIndex, bytes32 newUnsetGlobalIndexHashChain)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) WatchUpdatedUnsetGlobalIndexHashChain(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainv1010UpdatedUnsetGlobalIndexHashChain) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainv1010.contract.WatchLogs(opts, "UpdatedUnsetGlobalIndexHashChain")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainv1010UpdatedUnsetGlobalIndexHashChain)
				if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "UpdatedUnsetGlobalIndexHashChain", log); err != nil {
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

// ParseUpdatedUnsetGlobalIndexHashChain is a log parse operation binding the contract event 0xc80e0aca446a59735359a7ae46124b57c47b892827642779bc6dafc84ba90b03.
//
// Solidity: event UpdatedUnsetGlobalIndexHashChain(bytes32 unsetGlobalIndex, bytes32 newUnsetGlobalIndexHashChain)
func (_Bridgel2sovereignchainv1010 *Bridgel2sovereignchainv1010Filterer) ParseUpdatedUnsetGlobalIndexHashChain(log types.Log) (*Bridgel2sovereignchainv1010UpdatedUnsetGlobalIndexHashChain, error) {
	event := new(Bridgel2sovereignchainv1010UpdatedUnsetGlobalIndexHashChain)
	if err := _Bridgel2sovereignchainv1010.contract.UnpackLog(event, "UpdatedUnsetGlobalIndexHashChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

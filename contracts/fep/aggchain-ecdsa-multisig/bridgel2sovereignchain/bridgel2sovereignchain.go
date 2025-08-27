// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridgel2sovereignchain

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

// BridgeL2SovereignChainLeafData is an auto generated low-level Go binding around an user-defined struct.
type BridgeL2SovereignChainLeafData struct {
	LeafType           uint8
	OriginNetwork      uint32
	OriginAddress      common.Address
	DestinationNetwork uint32
	DestinationAddress common.Address
	Amount             *big.Int
	MetadataHash       [32]byte
}

// Bridgel2sovereignchainMetaData contains all meta data concerning the Bridgel2sovereignchain contract.
var Bridgel2sovereignchainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountDoesNotMatchMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeAddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DestinationNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmergencyStateNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EtherTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedProxyDeployment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputArraysLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDepositCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExpectedRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGlobalIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLBTLeaf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLeavesLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxyAdmin\",\"type\":\"address\"}],\"name\":\"InvalidProxyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSmtProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSovereignWETHAddressParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSubtreeFrontier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroNetworkID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxyAdmin\",\"type\":\"address\"}],\"name\":\"InvalidZeroProxyAdminOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"localBalanceTreeAmount\",\"type\":\"uint256\"}],\"name\":\"LocalBalanceTreeOverflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"localBalanceTreeAmount\",\"type\":\"uint256\"}],\"name\":\"LocalBalanceTreeUnderflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MsgValueNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeTokenIsEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewDepositCountExceedsMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoValueInMessagesOnGasTokenNetworks\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyBridgeManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyDeployer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyBridgePauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyBridgeUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootRemover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingEmergencyBridgePauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingEmergencyBridgeUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingProxiedTokensManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProxiedTokensManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OriginNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyMapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyUpdated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotMapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotRemapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WETHRemappingNotSupportedOnGasTokenNetworks\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldEmergencyBridgePauser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newEmergencyBridgePauser\",\"type\":\"address\"}],\"name\":\"AcceptEmergencyBridgePauserRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldEmergencyBridgeUnpauser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newEmergencyBridgeUnpauser\",\"type\":\"address\"}],\"name\":\"AcceptEmergencyBridgeUnpauserRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldProxiedTokensManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newProxiedTokensManager\",\"type\":\"address\"}],\"name\":\"AcceptProxiedTokensManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDepositCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"BackwardLET\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"depositCount\",\"type\":\"uint32\"}],\"name\":\"BridgeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDepositCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"ForwardLET\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"legacyTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"updatedTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MigrateLegacyToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wrappedTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"NewWrappedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignTokenAddress\",\"type\":\"address\"}],\"name\":\"RemoveLegacySovereignTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridgeManager\",\"type\":\"address\"}],\"name\":\"SetBridgeManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"leafIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"sourceNetwork\",\"type\":\"uint32\"}],\"name\":\"SetClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"SetLocalBalanceTree\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"SetSovereignTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignWETHTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"SetSovereignWETHAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentEmergencyBridgePauser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newEmergencyBridgePauser\",\"type\":\"address\"}],\"name\":\"TransferEmergencyBridgePauserRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentEmergencyBridgeUnpauser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newEmergencyBridgeUnpauser\",\"type\":\"address\"}],\"name\":\"TransferEmergencyBridgeUnpauserRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentProxiedTokensManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newProxiedTokensManager\",\"type\":\"address\"}],\"name\":\"TransferProxiedTokensManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"claimedGlobalIndex\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newClaimedGlobalIndexHashChain\",\"type\":\"bytes32\"}],\"name\":\"UpdatedClaimedGlobalIndexHashChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"unsetGlobalIndex\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newUnsetGlobalIndexHashChain\",\"type\":\"bytes32\"}],\"name\":\"UpdatedUnsetGlobalIndexHashChain\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BRIDGE_SOVEREIGN_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BRIDGE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_BYTECODE_TRANSPARENT_PROXY\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETHToken\",\"outputs\":[{\"internalType\":\"contractITokenWrappedBridgeUpgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptEmergencyBridgePauserRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptEmergencyBridgeUnpauserRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptProxiedTokensManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDepositCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[32]\",\"name\":\"newFrontier\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32\",\"name\":\"nextLeaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"proof\",\"type\":\"bytes32[32]\"}],\"name\":\"backwardLET\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"permitData\",\"type\":\"bytes\"}],\"name\":\"bridgeAsset\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeLib\",\"outputs\":[{\"internalType\":\"contractBridgeLib\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountWETH\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessageWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"calculateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimedBitMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimedGlobalIndexHashChain\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"name\":\"computeTokenProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"deployWrappedTokenAndRemap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyBridgePauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyBridgeUnpauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structBridgeL2SovereignChain.LeafData[]\",\"name\":\"newLeaves\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"expectedStateRoot\",\"type\":\"bytes32\"}],\"name\":\"forwardLET\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"name\":\"getLeafValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProxiedTokensManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenWrappedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWrappedTokenBridgeImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_polygonRollupManager\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_bridgeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sovereignWETHAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_sovereignWETHAddressIsNotMintable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_emergencyBridgePauser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_emergencyBridgeUnpauser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proxiedTokensManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"leafIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"sourceBridgeNetwork\",\"type\":\"uint32\"}],\"name\":\"isClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdatedDepositCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tokenInfoHash\",\"type\":\"bytes32\"}],\"name\":\"localBalanceTree\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"legacyTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permitData\",\"type\":\"bytes\"}],\"name\":\"migrateLegacyToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingEmergencyBridgePauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingEmergencyBridgeUnpauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingProxiedTokensManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"polygonRollupManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiedTokensManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"legacySovereignTokenAddress\",\"type\":\"address\"}],\"name\":\"removeLegacySovereignTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeManager\",\"type\":\"address\"}],\"name\":\"setBridgeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"originNetwork\",\"type\":\"uint32[]\"},{\"internalType\":\"address[]\",\"name\":\"originTokenAddress\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"name\":\"setLocalBalanceTree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"globalIndexes\",\"type\":\"uint256[]\"}],\"name\":\"setMultipleClaims\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"originNetworks\",\"type\":\"uint32[]\"},{\"internalType\":\"address[]\",\"name\":\"originTokenAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"sovereignTokenAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"isNotMintable\",\"type\":\"bool[]\"}],\"name\":\"setMultipleSovereignTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sovereignWETHTokenAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"setSovereignWETHAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenInfoToWrappedToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newEmergencyBridgePauser\",\"type\":\"address\"}],\"name\":\"transferEmergencyBridgePauserRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newEmergencyBridgeUnpauser\",\"type\":\"address\"}],\"name\":\"transferEmergencyBridgeUnpauserRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newProxiedTokensManager\",\"type\":\"address\"}],\"name\":\"transferProxiedTokensManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unsetGlobalIndexHashChain\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"globalIndexes\",\"type\":\"uint256[]\"}],\"name\":\"unsetMultipleClaims\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"verifyMerkleProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrappedAddress\",\"type\":\"address\"}],\"name\":\"wrappedAddressIsNotMintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedTokenBytecodeStorer\",\"outputs\":[{\"internalType\":\"contractIBytecodeStorer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wrappedTokenToTokenInfo\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x610100604052348015610010575f5ffd5b5060405161001d9061017d565b604051809103905ff080158015610036573d5f5f3e3d5ffd5b506001600160a01b031660805260405161004f9061018a565b604051809103905ff080158015610068573d5f5f3e3d5ffd5b506001600160a01b031660c05260405161008190610197565b604051809103905ff08015801561009a573d5f5f3e3d5ffd5b506001600160a01b031660a0526100af6100c0565b3360e0526100bb6100c0565b6101a4565b5f54610100900460ff161561012b5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116101561017b575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b610fcf806161d583390190565b6117d1806171a483390190565b610d5d8061897583390190565b60805160a05160c05160e051615fe26101f35f395f610dd501525f61068a01525f818161088e015281816128ec015281816131a1015261411301525f818161063d015261295d0152615fe25ff3fe60806040526004361061041a575f3560e01c806381b1c1741161021d578063c514f24e11610122578063eabd372a116100b7578063f5efcd7911610087578063f811bff71161006d578063f811bff714610cd2578063fb57083414610cf1578063fd7640e814610d10575f5ffd5b8063f5efcd7914610c82578063f67566e414610ca1575f5ffd5b8063eabd372a14610bfa578063ece93c6f14610c19578063ee25560b14610c38578063f214e16114610c63575f5ffd5b8063d02103ca116100f2578063d02103ca14610b7f578063d9cb3aec14610ba7578063dbc1697614610bd2578063e88f043614610be6575f5ffd5b8063c514f24e14610b1a578063cc46163214610b2e578063ccaa2d1114610b4d578063cd58657914610b6c575f5ffd5b8063ae24490a116101b2578063bab161bf11610182578063bf130d7f11610168578063bf130d7f14610aae578063c00f14ab14610acd578063c0f4916314610aec575f5ffd5b8063bab161bf14610a6a578063be5831c714610a8b575f5ffd5b8063ae24490a146109ee578063b0b3792014610a0d578063b458696214610a2c578063b8b284d014610a4b575f5ffd5b80638c668f1c116101ed5780638c668f1c1461097d5780638d942096146109915780638ed7e3f2146109b05780638f9720bb146109cf575f5ffd5b806381b1c174146108f757806383f244031461092b5780638b37b8731461094a5780638bd309c31461095e575f5ffd5b80633c351e101161032357806365d6f654116102b85780636ee84b231161028857806370ecac1b1161026e57806370ecac1b146108b057806379e2cf97146108cf5780638129fc1c146108e3575f5ffd5b80636ee84b23146108685780636f0bc3da1461087d575f5ffd5b806365d6f654146107c357806369e3ab121461080b5780636e4ecfed1461082a5780636e974cd414610849575f5ffd5b806354fd4d50116102f357806354fd4d501461074357806357cfbee3146107715780635ca1e16514610790578063606617ff146107a4575f5ffd5b80633c351e10146106ae5780633cbc795b146106cd5780633e197043146107055780634b2f336d14610724575f5ffd5b806322e95f2c116103b35780632f84c69011610383578063381fef6d11610369578063381fef6d1461062c57806338b8fbbb1461065f5780633b2fee9a1461067c575f5ffd5b80632f84c690146105a5578063318aee3d146105c4575f5ffd5b806322e95f2c1461053d578063240ff3781461055c57806327aef4e81461056f5780632dfdf0b514610590575f5ffd5b806315064c96116103ee57806315064c96146104be5780631c208229146104e75780631d081d8c146105065780632072f6c514610529575f5ffd5b80626ee1711461041e57806303e6e1161461043f578063136a2c601461048057806314cc01a01461049f575b5f5ffd5b348015610429575f5ffd5b5061043d610438366004614d99565b610d2f565b005b34801561044a575f5ffd5b5060a8546104639061010090046001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561048b575f5ffd5b5061043d61049a366004614f12565b611288565b3480156104aa575f5ffd5b5060a354610463906001600160a01b031681565b3480156104c9575f5ffd5b506068546104d79060ff1681565b6040519015158152602001610477565b3480156104f2575f5ffd5b5061043d610501366004614f12565b61144d565b348015610511575f5ffd5b5061051b60a55481565b604051908152602001610477565b348015610534575f5ffd5b5061043d6115ff565b348015610548575f5ffd5b50610463610557366004614f4c565b611634565b61043d61056a366004614fc6565b611699565b34801561057a575f5ffd5b50610583611709565b6040516104779190615088565b34801561059b575f5ffd5b5061051b60535481565b3480156105b0575f5ffd5b5060a454610463906001600160a01b031681565b3480156105cf575f5ffd5b506106086105de36600461509a565b606b6020525f908152604090205463ffffffff81169064010000000090046001600160a01b031682565b6040805163ffffffff90931683526001600160a01b03909116602083015201610477565b348015610637575f5ffd5b506104637f000000000000000000000000000000000000000000000000000000000000000081565b34801561066a575f5ffd5b506070546001600160a01b0316610463565b348015610687575f5ffd5b507f0000000000000000000000000000000000000000000000000000000000000000610463565b3480156106b9575f5ffd5b50606d54610463906001600160a01b031681565b3480156106d8575f5ffd5b50606d546106f090600160a01b900463ffffffff1681565b60405163ffffffff9091168152602001610477565b348015610710575f5ffd5b5061051b61071f3660046150c3565b611795565b34801561072f575f5ffd5b50606f54610463906001600160a01b031681565b34801561074e575f5ffd5b50604080518082019091526006815265076322e302e360d41b6020820152610583565b34801561077c575f5ffd5b5061043d61078b366004615203565b611839565b34801561079b575f5ffd5b5061051b611925565b3480156107af575f5ffd5b5060aa54610463906001600160a01b031681565b3480156107ce575f5ffd5b506105836040518060400160405280600681526020017f76312e312e30000000000000000000000000000000000000000000000000000081525081565b348015610816575f5ffd5b5061043d61082536600461509a565b6119a4565b348015610835575f5ffd5b50607054610463906001600160a01b031681565b348015610854575f5ffd5b5061043d61086336600461531c565b611a47565b348015610873575f5ffd5b5061051b60a65481565b348015610888575f5ffd5b506104637f000000000000000000000000000000000000000000000000000000000000000081565b3480156108bb575f5ffd5b5061043d6108ca366004615362565b611db7565b3480156108da575f5ffd5b5061043d611fab565b3480156108ee575f5ffd5b5061043d611fcc565b348015610902575f5ffd5b506104636109113660046153d7565b606a6020525f90815260409020546001600160a01b031681565b348015610936575f5ffd5b5061051b6109453660046153ff565b611fe5565b348015610955575f5ffd5b5061043d612074565b348015610969575f5ffd5b5061043d61097836600461509a565b612105565b348015610988575f5ffd5b5061043d61218c565b34801561099c575f5ffd5b5061043d6109ab36600461509a565b61221d565b3480156109bb575f5ffd5b50606c54610463906001600160a01b031681565b3480156109da575f5ffd5b5061043d6109e936600461543b565b6122a4565b3480156109f9575f5ffd5b5060a954610463906001600160a01b031681565b348015610a18575f5ffd5b5061043d610a273660046154c9565b612522565b348015610a37575f5ffd5b5061043d610a4636600461509a565b612697565b348015610a56575f5ffd5b5061043d610a65366004615521565b612817565b348015610a75575f5ffd5b506068546106f090610100900463ffffffff1681565b348015610a96575f5ffd5b506068546106f090600160c81b900463ffffffff1681565b348015610ab9575f5ffd5b5061043d610ac836600461559f565b612895565b348015610ad8575f5ffd5b50610583610ae736600461509a565b6128ca565b348015610af7575f5ffd5b506104d7610b0636600461509a565b60a26020525f908152604090205460ff1681565b348015610b25575f5ffd5b50610583612959565b348015610b39575f5ffd5b506104d7610b483660046155cb565b6129e2565b348015610b58575f5ffd5b5061043d610b673660046155fc565b612a33565b61043d610b7a3660046156db565b612e6b565b348015610b8a575f5ffd5b50606854610463906501000000000090046001600160a01b031681565b348015610bb2575f5ffd5b5061051b610bc13660046153d7565b60a76020525f908152604090205481565b348015610bdd575f5ffd5b5061043d613288565b348015610bf1575f5ffd5b5061043d6132bb565b348015610c05575f5ffd5b5061043d610c1436600461509a565b61336a565b348015610c24575f5ffd5b50607154610463906001600160a01b031681565b348015610c43575f5ffd5b5061051b610c523660046153d7565b60696020525f908152604090205481565b348015610c6e575f5ffd5b50610463610c7d366004614f4c565b61340a565b348015610c8d575f5ffd5b5061043d610c9c3660046155fc565b613505565b348015610cac575f5ffd5b5061058360405180604001604052806006815260200165076322e302e360d41b81525081565b348015610cdd575f5ffd5b5061043d610cec36600461576b565b613765565b348015610cfc575f5ffd5b506104d7610d0b3660046157fb565b613838565b348015610d1b575f5ffd5b5061043d610d2a366004615840565b61384f565b5f54600390610100900460ff16158015610d4f57505f5460ff8083169116105b610db75760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805461ffff191660ff831617610100179055336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610e135760405163618bbdd560e01b815260040160405180910390fd5b6001600160a01b038a16610e3a5760405163f6b2911f60e01b815260040160405180910390fd5b8c63ffffffff165f03610e6057604051634e702fa560e01b815260040160405180910390fd5b8c606860016101000a81548163ffffffff021916908363ffffffff16021790555089606860056101000a8154816001600160a01b0302191690836001600160a01b0316021790555088606c5f6101000a8154816001600160a01b0302191690836001600160a01b031602179055508660a35f6101000a8154816001600160a01b0302191690836001600160a01b031602179055508360a45f6101000a8154816001600160a01b0302191690836001600160a01b031602179055507f85d2bdfbe58cd81abf8199c13ce2509204be4aba8603b9d29f52c4e13e7bb7935f60a45f9054906101000a90046001600160a01b0316604051610f749291906001600160a01b0392831681529116602082015260400190565b60405180910390a160a980546001600160a01b0319166001600160a01b038516908117909155604080515f815260208101929092527f24cc8295aa5110cc216695db944ad2458c7795c6404449be980c3ce14aed752d910160405180910390a1306001600160a01b03831603610ffd57604051631ae0e03360e01b815260040160405180910390fd5b6001600160a01b0382166110245760405163f6b2911f60e01b815260040160405180910390fd5b607080546001600160a01b0319166001600160a01b038416908117909155604080515f815260208101929092527fa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f2910160405180910390a16001600160a01b038c166110e65763ffffffff8b16156110af57604051630d43a60960e11b815260040160405180910390fd5b6001600160a01b0386161515806110c35750845b156110e157604051630e6e237560e11b815260040160405180910390fd5b611231565b606d805463ffffffff8d16600160a01b026001600160c01b03199091166001600160a01b038f1617179055606e61111d8982615904565b506001600160a01b0386166111f95784151560010361114f57604051630e6e237560e11b815260040160405180910390fd5b6111d45f5f1b60126040516020016111c091906060808252600d908201527f5772617070656420457468657200000000000000000000000000000000000000608082015260a060208201819052600490820152630ae8aa8960e31b60c082015260ff91909116604082015260e00190565b604051602081830303815290604052613a52565b606f80546001600160a01b0319166001600160a01b0392909216919091179055611231565b606f80546001600160a01b0319166001600160a01b0388169081179091555f90815260a260205260409020805460ff19168615151790555b611239613b33565b5f805461ff001916905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050505050505050505050565b606854604080516391eb796d60e01b8152905133926501000000000090046001600160a01b0316916391eb796d9160048083019260209291908290030181865afa1580156112d8573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112fc91906159bf565b6001600160a01b0316146113235760405163a34ddeb160e01b815260040160405180910390fd5b5f5b8151811015611449575f828281518110611341576113416159da565b60209081029190910101519050805f600160401b821681036113b257602083901c61136d816001615a02565b91508361138e63ffffffff851667ffffffff00000000602085901b16615a1e565b146113ac5760405163071389e960e01b815260040160405180910390fd5b506113e5565b826113c763ffffffff8416600160401b615a1e565b146113e55760405163071389e960e01b815260040160405180910390fd5b6113ef8282613ba5565b60a6545f90815260208490526040902060a68190556040805185815260208101929092527fc80e0aca446a59735359a7ae46124b57c47b892827642779bc6dafc84ba90b03910160405180910390a1505050600101611325565b5050565b606854604080516391eb796d60e01b8152905133926501000000000090046001600160a01b0316916391eb796d9160048083019260209291908290030181865afa15801561149d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114c191906159bf565b6001600160a01b0316146114e85760405163a34ddeb160e01b815260040160405180910390fd5b5f5b8151811015611449575f828281518110611506576115066159da565b60209081029190910101519050805f600160401b8216810361157757602083901c611532816001615a02565b91508361155363ffffffff851667ffffffff00000000602085901b16615a1e565b146115715760405163071389e960e01b815260040160405180910390fd5b506115aa565b8261158c63ffffffff8416600160401b615a1e565b146115aa5760405163071389e960e01b815260040160405180910390fd5b6115b48282613c17565b6040805163ffffffff8085168252831660208201527fab48452e173be3f004b3208cc344b14cc7f8d977fc1ca28171286a7abb5d7570910160405180910390a15050506001016114ea565b60a4546001600160a01b0316331461162a57604051631344c5df60e11b815260040160405180910390fd5b611632613c8a565b565b6040805160e084901b6001600160e01b031916602080830191909152606084901b6001600160601b031916602483015282516018818403018152603890920183528151918101919091205f908152606a90915220546001600160a01b03165b92915050565b60685460ff16156116bd57604051630bc011ff60e21b815260040160405180910390fd5b34158015906116d65750606f546001600160a01b031615155b156116f4576040516301bd897160e61b815260040160405180910390fd5b611702858534868686613ce5565b5050505050565b606e805461171690615888565b80601f016020809104026020016040519081016040528092919081815260200182805461174290615888565b801561178d5780601f106117645761010080835404028352916020019161178d565b820191905f5260205f20905b81548152906001019060200180831161177057829003601f168201915b505050505081565b6040517fff0000000000000000000000000000000000000000000000000000000000000060f889901b1660208201526001600160e01b031960e088811b821660218401526001600160601b0319606089811b821660258601529188901b909216603984015285901b16603d82015260518101839052607181018290525f90609101604051602081830303815290604052805190602001209050979650505050505050565b60a3546001600160a01b03163314611864576040516357b738d160e11b815260040160405180910390fd5b8251845114158061187757508151845114155b8061188457508051845114155b156118a25760405163434f49f560e11b815260040160405180910390fd5b5f5b82518110156117025761191d8582815181106118c2576118c26159da565b60200260200101518583815181106118dc576118dc6159da565b60200260200101518584815181106118f6576118f66159da565b6020026020010151858581518110611910576119106159da565b6020026020010151613db9565b6001016118a4565b6053545f90819081805b602081101561199b578083901c6001166001036119745761196d6033826020811061195c5761195c6159da565b0154855f9182526020526040902090565b9350611984565b5f84815260208390526040902093505b5f828152602083905260409020915060010161192f565b50919392505050565b60a4546001600160a01b031633146119cf57604051631344c5df60e11b815260040160405180910390fd5b60a8805474ffffffffffffffffffffffffffffffffffffffff0019166101006001600160a01b038481169182029290921790925560a4546040805191909216815260208101929092527fb27de219766f47b82684842855ba6130b6dbf288ac66d1c3509e7bf17f4e925a91015b60405180910390a150565b60a3546001600160a01b03163314611a72576040516357b738d160e11b815260040160405180910390fd5b6001600160a01b038216158015611a8d575063ffffffff8316155b15611c1d575f611c0b5f5f1b606f5f9054906101000a90046001600160a01b03166001600160a01b03166306fdde036040518163ffffffff1660e01b81526004015f60405180830381865afa158015611ae8573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611b0f9190810190615a7d565b606f5f9054906101000a90046001600160a01b03166001600160a01b03166395d89b416040518163ffffffff1660e01b81526004015f60405180830381865afa158015611b5e573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611b859190810190615a7d565b606f5f9054906101000a90046001600160a01b03166001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015611bd5573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611bf99190615aaf565b6040516020016111c093929190615aca565b9050611c178183613f7d565b50505050565b6040516001600160e01b031960e085901b1660208201526001600160601b0319606084901b1660248201525f9060380160408051601f1981840301815291815281516020928301205f818152606a9093529120549091506001600160a01b031680611c9b5760405163828d566360e01b815260040160405180910390fd5b5f611da083836001600160a01b03166306fdde036040518163ffffffff1660e01b81526004015f60405180830381865afa158015611cdb573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611d029190810190615a7d565b846001600160a01b03166395d89b416040518163ffffffff1660e01b81526004015f60405180830381865afa158015611d3d573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611d649190810190615a7d565b856001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015611bd5573d5f5f3e3d5ffd5b9050611dae86868387613db9565b5050505b505050565b606854604080516391eb796d60e01b8152905133926501000000000090046001600160a01b0316916391eb796d9160048083019260209291908290030181865afa158015611e07573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e2b91906159bf565b6001600160a01b031614611e525760405163a34ddeb160e01b815260040160405180910390fd5b5f829003611e7357604051638488056760e01b815260040160405180910390fd5b5f5b82811015611edb575f848483818110611e9057611e906159da565b905060e00201803603810190611ea69190615b02565b9050611ed2815f015182602001518360400151846060015185608001518660a001518760c0015161401c565b50600101611e75565b505f611ee5611925565b9050818114611f075760405163f9c04eb160e01b815260040160405180910390fd5b6068546040516333d6247d60e01b815260048101839052650100000000009091046001600160a01b0316906333d6247d906024015f604051808303815f87803b158015611f52575f5ffd5b505af1158015611f64573d5f5f3e3d5ffd5b505060535460408051918252602082018590527ff833886dec6b10040de3e31a518be28ced3c9689a75d17741c23bcfb41975670935001905060405180910390a150505050565b605354606854600160c81b900463ffffffff16101561163257611632614054565b60405163f57ac68360e01b815260040160405180910390fd5b5f83815b602081101561206957600163ffffffff8516821c811690036120355761202e85826020811061201a5761201a6159da565b6020020135835f9182526020526040902090565b9150612061565b61205e8286836020811061204b5761204b6159da565b60200201355f9182526020526040902090565b91505b600101611fe9565b5090505b9392505050565b60aa546001600160a01b0316331461209f5760405163d491f0c160e01b815260040160405180910390fd5b60a9805460aa80546001600160a01b038082166001600160a01b0319808616821790965594909116909155604080519190921680825260208201939093527f85d2bdfbe58cd81abf8199c13ce2509204be4aba8603b9d29f52c4e13e7bb7939101611a3c565b6070546001600160a01b0316331461213057604051630866750360e01b815260040160405180910390fd5b607180546001600160a01b0319166001600160a01b038381169182179092556070546040805191909316815260208101919091527f0a34baa3feb299aef9c05cb59c6e0c8e7c0bcc65cbf0a647e7a7c8a2411591e29101611a3c565b6071546001600160a01b031633146121b757604051630b59ef2760e21b815260040160405180910390fd5b60708054607180546001600160a01b038082166001600160a01b0319808616821790965594909116909155604080519190921680825260208201939093527fa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f29101611a3c565b60a9546001600160a01b0316331461224857604051638e9d821f60e01b815260040160405180910390fd5b60aa80546001600160a01b0319166001600160a01b0383811691821790925560a9546040805191909316815260208101919091527ff01a62a06940517bbc898dec8c75794b9feabcd2d263c8de823b36dbbeb8779b9101611a3c565b606854604080516391eb796d60e01b8152905133926501000000000090046001600160a01b0316916391eb796d9160048083019260209291908290030181865afa1580156122f4573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061231891906159bf565b6001600160a01b03161461233f5760405163a34ddeb160e01b815260040160405180910390fd5b8151835114158061235257508051835114155b156123705760405163434f49f560e11b815260040160405180910390fd5b5f5b8351811015611c1757606860019054906101000a900463ffffffff1663ffffffff168482815181106123a6576123a66159da565b602002602001015163ffffffff16036123d25760405163b869a63f60e01b815260040160405180910390fd5b5f8482815181106123e5576123e56159da565b60200260200101518483815181106123ff576123ff6159da565b602002602001015160405160200161243d92919060e09290921b6001600160e01b031916825260601b6001600160601b031916600482015260180190565b604051602081830303815290604052805190602001209050828281518110612467576124676159da565b602002602001015160a75f8381526020019081526020015f2081905550838281518110612496576124966159da565b60200260200101516001600160a01b03168583815181106124b9576124b96159da565b602002602001015163ffffffff167f73a03a6d9efc14b9f22a3af967e98580549eb76b1113b6a09a57ce1dae36ecd28585815181106124fa576124fa6159da565b602002602001015160405161251191815260200190565b60405180910390a350600101612372565b8015612533576125338483836140fc565b6001600160a01b038085165f908152606b602090815260409182902082518084019093525463ffffffff811683526401000000009004909216918101829052906125905760405163828d566360e01b815260040160405180910390fd5b5f606a5f835f015184602001516040516020016125d392919060e09290921b6001600160e01b031916825260601b6001600160601b031916600482015260180190565b60408051601f198184030181529181528151602092830120835290820192909252015f20546001600160a01b039081169150861681036126265760405163e273c4a160e01b815260040160405180910390fd5b5f6126318787614190565b905061263e82338361430f565b604080513381526001600160a01b0389811660208301528416818301526060810183905290517fb7f8fd4d1faf9b2929dc269f59c53e3a2bccc44e9950f33a568fcbcb37eb69a99181900360800190a150505050505050565b60a3546001600160a01b031633146126c2576040516357b738d160e11b815260040160405180910390fd5b6001600160a01b038082165f908152606b6020908152604080832081518083018352905463ffffffff811680835264010000000090910490951681840181905291519094612738939092910160e09290921b6001600160e01b031916825260601b6001600160601b031916600482015260180190565b60408051601f1981840301815291815281516020928301205f818152606a9093529120549091506001600160a01b0316158061278c57505f818152606a60205260409020546001600160a01b038481169116145b156127aa5760405163e0c897a760e01b815260040160405180910390fd5b6001600160a01b0383165f818152606b6020908152604080832080546001600160c01b031916905560a2825291829020805460ff1916905590519182527fc2ae0bd0ec0fd0352bfe5bacac49637af342c1e40f1b80a7f74440dc7fe3f063910160405180910390a1505050565b60685460ff161561283b57604051630bc011ff60e21b815260040160405180910390fd5b606f546001600160a01b03166128645760405163dde3cda760e01b815260040160405180910390fd5b606f545f9061287c906001600160a01b031686614190565b905061288c878783878787613ce5565b50505050505050565b60a3546001600160a01b031633146128c0576040516357b738d160e11b815260040160405180910390fd5b6114498282613f7d565b60405163c00f14ab60e01b81526001600160a01b0382811660048301526060917f00000000000000000000000000000000000000000000000000000000000000009091169063c00f14ab906024015f60405180830381865afa158015612932573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526116939190810190615b8e565b60607f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c514f24e6040518163ffffffff1660e01b81526004015f60405180830381865afa1580156129b6573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526129dd9190810190615b8e565b905090565b5f806129f964010000000063ffffffff8516615bd3565b612a099063ffffffff8616615a1e565b600881901c5f90815260696020526040902054600160ff9092169190911b90811614949350505050565b60685460ff1615612a5757604051630bc011ff60e21b815260040160405180910390fd5b612a5f61439c565b60685463ffffffff8681166101009092041614612a8f576040516302caf51760e11b815260040160405180910390fd5b612aba8c8c8c8c8c5f8d8d8d8d8d8d8d604051612aad929190615bea565b60405180910390206143f5565b604080518b815263ffffffff891660208201526001600160a01b0388811682840152861660608201526080810185905290517f1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d9181900360a00190a16001600160a01b038616158015612b31575063ffffffff8716155b15612c0f57606f546001600160a01b0316612bf3575f6001600160a01b03851684825b6040519080825280601f01601f191660200182016040528015612b7e576020820181803683370190505b50604051612b8c9190615bf9565b5f6040518083038185875af1925050503d805f8114612bc6576040519150601f19603f3d011682016040523d82523d5f602084013e612bcb565b606091505b5050905080612bed57604051630ce8f45160e31b815260040160405180910390fd5b50612e54565b606f54612c0a906001600160a01b0316858561430f565b612e54565b606d546001600160a01b038781169116148015612c3d5750606d5463ffffffff888116600160a01b90920416145b15612c54575f6001600160a01b0385168482612b54565b60685463ffffffff610100909104811690881603612c8057612c0a6001600160a01b03871685856144a5565b6040516001600160e01b031960e089901b1660208201526001600160601b0319606088901b1660248201525f9060380160408051601f1981840301815291815281516020928301205f818152606a9093529120549091506001600160a01b031680612e46575f612d258386868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250613a5292505050565b9050612d3281888861430f565b80606a5f8581526020019081526020015f205f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555060405180604001604052808b63ffffffff1681526020018a6001600160a01b0316815250606b5f836001600160a01b03166001600160a01b031681526020019081526020015f205f820151815f015f6101000a81548163ffffffff021916908363ffffffff1602179055506020820151815f0160046101000a8154816001600160a01b0302191690836001600160a01b031602179055509050507f490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a398a8a838888604051612e38959493929190615c3c565b60405180910390a150612e51565b612e5181878761430f565b50505b612e5d60018055565b505050505050505050505050565b60685460ff1615612e8f57604051630bc011ff60e21b815260040160405180910390fd5b612e9761439c565b60685463ffffffff610100909104811690881603612ec8576040516302caf51760e11b815260040160405180910390fd5b5f806060876001600160a01b038816612fab57883414612efb5760405163b89240f560e01b815260040160405180910390fd5b606d54606e80546001600160a01b0383169650600160a01b90920463ffffffff16945090612f2890615888565b80601f0160208091040260200160405190810160405280929190818152602001828054612f5490615888565b8015612f9f5780601f10612f7657610100808354040283529160200191612f9f565b820191905f5260205f20905b815481529060010190602001808311612f8257829003601f168201915b50505050509150613210565b3415612fca5760405163798ee6f160e01b815260040160405180910390fd5b8415612fdb57612fdb8887876140fc565b606f546001600160a01b039081169089160361300257612ffb888a614190565b9050613210565b6001600160a01b038089165f908152606b602090815260409182902082518084019093525463ffffffff811683526401000000009004909216918101829052901515806130555750805163ffffffff1615155b1561307757613064898b614190565b6020820151825190965094509150613182565b6040516370a0823160e01b81523060048201525f906001600160a01b038b16906370a0823190602401602060405180830381865afa1580156130bb573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130df9190615c84565b90506130f66001600160a01b038b1633308e61450a565b6040516370a0823160e01b81523060048201525f906001600160a01b038c16906370a0823190602401602060405180830381865afa15801561313a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061315e9190615c84565b905061316a8282615c9b565b6068548c9850610100900463ffffffff169650935050505b60405163c00f14ab60e01b81526001600160a01b038a811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063c00f14ab906024015f60405180830381865afa1580156131e5573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261320c9190810190615b8e565b9250505b7f501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b5f84868e8e868860535460405161324f989796959493929190615cae565b60405180910390a161326d5f84868e8e86888051906020012061401c565b861561327b5761327b614054565b5050505061288c60018055565b60a9546001600160a01b031633146132b357604051638e9d821f60e01b815260040160405180910390fd5b611632614543565b60a85461010090046001600160a01b031633146132eb57604051637bb0100f60e01b815260040160405180910390fd5b60a4805460a880546001600160a01b03610100820481166001600160a01b03198516811790955574ffffffffffffffffffffffffffffffffffffffff0019909116909155604080519190921680825260208201939093527f85d2bdfbe58cd81abf8199c13ce2509204be4aba8603b9d29f52c4e13e7bb7939101611a3c565b60a3546001600160a01b03163314613395576040516357b738d160e11b815260040160405180910390fd5b6001600160a01b0381166133bc5760405163f6b2911f60e01b815260040160405180910390fd5b60a380546001600160a01b0319166001600160a01b0383169081179091556040519081527f32cf74f8a6d5f88593984d2cd52be5592bfa6884f5896175801a5069ef09cd6790602001611a3c565b6040516001600160e01b031960e084901b1660208201526001600160601b0319606083901b1660248201525f9081906038016040516020818303038152906040528051906020012090505f60ff60f81b3083613464612959565b6040516020016134749190615bf9565b604051602081830303815290604052805190602001206040516020016134e494939291907fff0000000000000000000000000000000000000000000000000000000000000094909416845260609290921b6001600160601b03191660018401526015830152603582015260550190565b60408051808303601f19018152919052805160209091012095945050505050565b60685460ff161561352957604051630bc011ff60e21b815260040160405180910390fd5b60685463ffffffff8681166101009092041614613559576040516302caf51760e11b815260040160405180910390fd5b6135788c8c8c8c8c60018d8d8d8d8d8d8d604051612aad929190615bea565b604080518b815263ffffffff891660208201526001600160a01b0388811682840152861660608201526080810185905290517f1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d9181900360a00190a1606f545f906001600160a01b031661368757846001600160a01b031684888a86866040516024016136089493929190615d22565b60408051601f198184030181529181526020820180516001600160e01b0316630c035af960e11b1790525161363d9190615bf9565b5f6040518083038185875af1925050503d805f8114613677576040519150601f19603f3d011682016040523d82523d5f602084013e61367c565b606091505b505080915050613738565b606f5461369e906001600160a01b0316868661430f565b846001600160a01b0316878985856040516024016136bf9493929190615d22565b60408051601f198184030181529181526020820180516001600160e01b0316630c035af960e11b179052516136f49190615bf9565b5f604051808303815f865af19150503d805f811461372d576040519150601f19603f3d011682016040523d82523d5f602084013e613732565b606091505b50909150505b80613756576040516337e391c360e01b815260040160405180910390fd5b50505050505050505050505050565b5f54610100900460ff161580801561378357505f54600160ff909116105b8061379c5750303b15801561379c57505f5460ff166001145b6137ff5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610dae565b5f805460ff191660011790558015611fcc575f805461ff00191661010017905560405163f57ac68360e01b815260040160405180910390fd5b5f81613845868686611fe5565b1495945050505050565b606854604080516391eb796d60e01b8152905133926501000000000090046001600160a01b0316916391eb796d9160048083019260209291908290030181865afa15801561389f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906138c391906159bf565b6001600160a01b0316146138ea5760405163a34ddeb160e01b815260040160405180910390fd5b605354841061390c57604051631abd431560e31b815260040160405180910390fd5b61391a828286610d0b611925565b613937576040516338105f3b60e21b815260040160405180910390fd5b61394284848361459a565b61395f576040516315ea0fc560e31b815260040160405180910390fd5b5f5b602081101561399e5783816020811061397c5761397c6159da565b602002013560338260208110613994576139946159da565b0155600101613961565b5060538490555f6139ad611925565b6068546040516333d6247d60e01b8152600481018390529192506501000000000090046001600160a01b0316906333d6247d906024015f604051808303815f87803b1580156139fa575f5ffd5b505af1158015613a0c573d5f5f3e3d5ffd5b505060408051888152602081018590527f90799a4e436a5b97adca76982886b3c894dc74b50169ab082b9a069f99317e6193500190505b60405180910390a15050505050565b5f5f613a5c612959565b604051602001613a6c9190615bf9565b6040516020818303038152906040529050838151602083015ff591506001600160a01b038216613aaf576040516331682e8d60e11b815260040160405180910390fd5b5f5f5f85806020019051810190613ac69190615d50565b925092509250846001600160a01b0316631624f6c68484846040518463ffffffff1660e01b8152600401613afc93929190615aca565b5f604051808303815f87803b158015613b13575f5ffd5b505af1158015613b25573d5f5f3e3d5ffd5b505050505050505092915050565b5f54610100900460ff16613b9d5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608401610dae565b61163261461f565b5f613bbb64010000000063ffffffff8416615bd3565b613bcb9063ffffffff8516615a1e565b600881901c5f8181526069602052604090208054600160ff851690811b91821892839055939450919291908082161561288c57604051630631b5f760e31b815260040160405180910390fd5b5f613c2d64010000000063ffffffff8416615bd3565b613c3d9063ffffffff8516615a1e565b600881901c5f8181526069602052604081208054600160ff861690811b9182189283905594955092939291818316900361288c57604051630c8d9eab60e31b815260040160405180910390fd5b60685460ff1615613cae57604051630bc011ff60e21b815260040160405180910390fd5b6068805460ff191660011790556040517f2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497905f90a1565b60685463ffffffff610100909104811690871603613d16576040516302caf51760e11b815260040160405180910390fd5b7f501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b6001606860019054906101000a900463ffffffff16338989898888605354604051613d6a99989796959493929190615dbd565b60405180910390a1613dab6001606860019054906101000a900463ffffffff16338989898888604051613d9e929190615bea565b604051809103902061401c565b8215611dae57611dae614054565b6001600160a01b0383161580613dd657506001600160a01b038216155b15613df45760405163f6b2911f60e01b815260040160405180910390fd5b60685463ffffffff610100909104811690851603613e255760405163658b23ad60e01b815260040160405180910390fd5b6001600160a01b038281165f908152606b602052604090205464010000000090041615613e65576040516317abdeeb60e21b815260040160405180910390fd5b6040516001600160e01b031960e086901b1660208201526001600160601b0319606085901b1660248201525f9060380160408051808303601f1901815282825280516020918201205f818152606a835283812080546001600160a01b0319166001600160a01b038a8116918217909255868601865263ffffffff8c81168089528c8416878a01818152848752606b89528987209a518b54915194166001600160c01b03199091161764010000000093909516929092029390931790975560a2855291859020805460ff191689151590811790915585519182529381019590955292840192909252606083015291507fdbe8a5da6a7a916d9adfda9160167a0f8a3da415ee6610e810e753853597fce790608001613a43565b606d546001600160a01b0316613fa657604051634cb4711360e11b815260040160405180910390fd5b606f80546001600160a01b0319166001600160a01b0384169081179091555f81815260a26020908152604091829020805460ff19168515159081179091558251938452908301527fc7318b7ed6ba4f2908a3de396d8ab49b1dadb55db5b55123247a401f29ff8d82910160405180910390a15050565b61402b87878787878787614689565b60ff871661403e5761403e8686846146a0565b5f1960ff88160161288c5761288c5f5f846146a0565b6053546068805463ffffffff909216600160c81b027fffffff00000000ffffffffffffffffffffffffffffffffffffffffffffffffff90921691909117908190556001600160a01b0365010000000000909104166333d6247d6140b5611925565b6040518263ffffffff1660e01b81526004016140d391815260200190565b5f604051808303815f87803b1580156140ea575f5ffd5b505af1158015611c17573d5f5f3e3d5ffd5b60405163a28fa4a360e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a28fa4a3906141509086908690869033903090600401615e33565b6020604051808303815f875af115801561416c573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611c179190615e77565b6001600160a01b0382165f90815260a2602052604081205460ff16156142ac576040516370a0823160e01b81523060048201525f906001600160a01b038516906370a0823190602401602060405180830381865afa1580156141f4573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906142189190615c84565b905061422f6001600160a01b03851633308661450a565b6040516370a0823160e01b81523060048201525f906001600160a01b038616906370a0823190602401602060405180830381865afa158015614273573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906142979190615c84565b90506142a38282615c9b565b92505050611693565b604051632770a7eb60e21b8152336004820152602481018390526001600160a01b03841690639dc29fac906044015f604051808303815f87803b1580156142f1575f5ffd5b505af1158015614303573d5f5f3e3d5ffd5b50505050819050611693565b6001600160a01b0383165f90815260a2602052604090205460ff161561434357611db26001600160a01b03841683836144a5565b6040516340c10f1960e01b81526001600160a01b038381166004830152602482018390528416906340c10f19906044015f604051808303815f87803b15801561438a575f5ffd5b505af115801561288c573d5f5f3e3d5ffd5b6002600154036143ee5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610dae565b6002600155565b5f61440588888888888888611795565b90506144158d8d8d8d8d86614791565b60a55461443d9061442f8d845f9182526020526040902090565b5f9182526020526040902090565b60a5819055604080518d815260208101929092527f3e5936f910a78eb5181813a939c8d4c3e4d85f87943f659380d82ac6221b0e92910160405180910390a160ff881661448f5761448f87878561494d565b5f1960ff891601613756576137565f5f8561494d565b6040516001600160a01b03838116602483015260448201839052611db291859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050614a3e565b60018055565b6040516001600160a01b038481166024830152838116604483015260648201839052611c179186918216906323b872dd906084016144d2565b60685460ff1661456657604051635386698160e01b815260040160405180910390fd5b6068805460ff191690556040517f1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3905f90a1565b5f83815b81158015906145ad5750602081105b1561461357816001166001036145fa578381602081106145cf576145cf6159da565b60200201358582602081106145e6576145e66159da565b6020020135146145fa575f9250505061206d565b8061460481615e92565b915050600182901c915061459e565b50600195945050505050565b5f54610100900460ff166145045760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608401610dae565b61288c61469b88888888888888611795565b614a9f565b60685463ffffffff6101009091048116908416036146bd57505050565b6040516001600160e01b031960e085901b1660208201526001600160601b0319606084901b1660248201525f9060380160408051601f1981840301815291815281516020928301205f81815260a7909352912054909150821115614769575f81815260a76020526040908190205490516314603c0160e01b815263ffffffff861660048201526001600160a01b0385166024820152604481018490526064810191909152608401610dae565b5f81815260a7602052604081208054849290614786908490615c9b565b909155505050505050565b6068545f906501000000000090046001600160a01b031663257b36326147c086865f9182526020526040902090565b6040518263ffffffff1660e01b81526004016147de91815260200190565b6020604051808303815f875af11580156147fa573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061481e9190615c84565b9050805f0361483f57604051622f6fad60e01b815260040160405180910390fd5b5f80600160401b8716156148b1578691508161486563ffffffff8216600160401b615a1e565b146148835760405163071389e960e01b815260040160405180910390fd5b61488f848a8489613838565b6148ac576040516338105f3b60e21b815260040160405180910390fd5b614938565b602087901c6148c1816001615a02565b8893509150826148e563ffffffff821667ffffffff00000000602085901b16615a1e565b146149035760405163071389e960e01b815260040160405180910390fd5b614919614911868c86611fe5565b8a8389613838565b614936576040516338105f3b60e21b815260040160405180910390fd5b505b6149428282613c17565b505050505050505050565b60685463ffffffff61010090910481169084160361496a57505050565b6040516001600160e01b031960e085901b1660208201526001600160601b0319606084901b1660248201525f9060380160408051601f1981840301815291815281516020928301205f81815260a79093529120549091506149cc905f19615c9b565b821115614a21575f81815260a76020526040908190205490516323d7213360e01b815263ffffffff861660048201526001600160a01b0385166024820152604481018490526064810191909152608401610dae565b5f81815260a7602052604081208054849290614786908490615a1e565b5f614a526001600160a01b03841683614b5e565b905080515f14158015614a76575080806020019051810190614a749190615e77565b155b15611db257604051635274afe760e01b81526001600160a01b0384166004820152602401610dae565b806001614aae60206002615f8d565b614ab89190615c9b565b60535410614ad9576040516377ae67b360e11b815260040160405180910390fd5b5f60535f8154614ae890615e92565b918290555090505f5b6020811015614b55578082901c600116600103614b24578260338260208110614b1c57614b1c6159da565b015550505050565b614b4b60338260208110614b3a57614b3a6159da565b0154845f9182526020526040902090565b9250600101614af1565b50611db2615f98565b606061206d83835f845f5f856001600160a01b03168486604051614b829190615bf9565b5f6040518083038185875af1925050503d805f8114614bbc576040519150601f19603f3d011682016040523d82523d5f602084013e614bc1565b606091505b5091509150614bd1868383614bdb565b9695505050505050565b606082614bf057614beb82614c37565b61206d565b8151158015614c0757506001600160a01b0384163b155b15614c3057604051639996b31560e01b81526001600160a01b0385166004820152602401610dae565b508061206d565b805115614c475780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b50565b803563ffffffff81168114614c76575f5ffd5b919050565b6001600160a01b0381168114614c60575f5ffd5b8035614c7681614c7b565b634e487b7160e01b5f52604160045260245ffd5b60405160e0810167ffffffffffffffff81118282101715614cd157614cd1614c9a565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715614d0057614d00614c9a565b604052919050565b5f67ffffffffffffffff821115614d2157614d21614c9a565b50601f01601f191660200190565b5f82601f830112614d3e575f5ffd5b8135614d51614d4c82614d08565b614cd7565b818152846020838601011115614d65575f5ffd5b816020850160208301375f918101602001919091529392505050565b8015158114614c60575f5ffd5b8035614c7681614d81565b5f5f5f5f5f5f5f5f5f5f5f5f6101808d8f031215614db5575f5ffd5b614dbe8d614c63565b9b50614dcc60208e01614c8f565b9a50614dda60408e01614c63565b9950614de860608e01614c8f565b9850614df660808e01614c8f565b975067ffffffffffffffff60a08e01351115614e10575f5ffd5b614e208e60a08f01358f01614d2f565b9650614e2e60c08e01614c8f565b9550614e3c60e08e01614c8f565b9450614e4b6101008e01614d8e565b9350614e5a6101208e01614c8f565b9250614e696101408e01614c8f565b9150614e786101608e01614c8f565b90509295989b509295989b509295989b565b5f67ffffffffffffffff821115614ea357614ea3614c9a565b5060051b60200190565b5f82601f830112614ebc575f5ffd5b8135614eca614d4c82614e8a565b8082825260208201915060208360051b860101925085831115614eeb575f5ffd5b602085015b83811015614f08578035835260209283019201614ef0565b5095945050505050565b5f60208284031215614f22575f5ffd5b813567ffffffffffffffff811115614f38575f5ffd5b614f4484828501614ead565b949350505050565b5f5f60408385031215614f5d575f5ffd5b614f6683614c63565b91506020830135614f7681614c7b565b809150509250929050565b5f5f83601f840112614f91575f5ffd5b50813567ffffffffffffffff811115614fa8575f5ffd5b602083019150836020828501011115614fbf575f5ffd5b9250929050565b5f5f5f5f5f60808688031215614fda575f5ffd5b614fe386614c63565b94506020860135614ff381614c7b565b9350604086013561500381614d81565b9250606086013567ffffffffffffffff81111561501e575f5ffd5b61502a88828901614f81565b969995985093965092949392505050565b5f5b8381101561505557818101518382015260200161503d565b50505f910152565b5f815180845261507481602086016020860161503b565b601f01601f19169290920160200192915050565b602081525f61206d602083018461505d565b5f602082840312156150aa575f5ffd5b813561206d81614c7b565b60ff81168114614c60575f5ffd5b5f5f5f5f5f5f5f60e0888a0312156150d9575f5ffd5b87356150e4816150b5565b96506150f260208901614c63565b9550604088013561510281614c7b565b945061511060608901614c63565b9350608088013561512081614c7b565b9699959850939692959460a0840135945060c09093013592915050565b5f82601f83011261514c575f5ffd5b813561515a614d4c82614e8a565b8082825260208201915060208360051b86010192508583111561517b575f5ffd5b602085015b83811015614f085761519181614c63565b835260209283019201615180565b5f82601f8301126151ae575f5ffd5b81356151bc614d4c82614e8a565b8082825260208201915060208360051b8601019250858311156151dd575f5ffd5b602085015b83811015614f085780356151f581614c7b565b8352602092830192016151e2565b5f5f5f5f60808587031215615216575f5ffd5b843567ffffffffffffffff81111561522c575f5ffd5b6152388782880161513d565b945050602085013567ffffffffffffffff811115615254575f5ffd5b6152608782880161519f565b935050604085013567ffffffffffffffff81111561527c575f5ffd5b6152888782880161519f565b925050606085013567ffffffffffffffff8111156152a4575f5ffd5b8501601f810187136152b4575f5ffd5b80356152c2614d4c82614e8a565b8082825260208201915060208360051b8501019250898311156152e3575f5ffd5b6020840193505b8284101561530e5783356152fd81614d81565b8252602093840193909101906152ea565b969995985093965050505050565b5f5f5f6060848603121561532e575f5ffd5b61533784614c63565b9250602084013561534781614c7b565b9150604084013561535781614d81565b809150509250925092565b5f5f5f60408486031215615374575f5ffd5b833567ffffffffffffffff81111561538a575f5ffd5b8401601f8101861361539a575f5ffd5b803567ffffffffffffffff8111156153b0575f5ffd5b86602060e0830284010111156153c4575f5ffd5b6020918201979096509401359392505050565b5f602082840312156153e7575f5ffd5b5035919050565b806104008101831015611693575f5ffd5b5f5f5f6104408486031215615412575f5ffd5b8335925061542385602086016153ee565b91506154326104208501614c63565b90509250925092565b5f5f5f6060848603121561544d575f5ffd5b833567ffffffffffffffff811115615463575f5ffd5b61546f8682870161513d565b935050602084013567ffffffffffffffff81111561548b575f5ffd5b6154978682870161519f565b925050604084013567ffffffffffffffff8111156154b3575f5ffd5b6154bf86828701614ead565b9150509250925092565b5f5f5f5f606085870312156154dc575f5ffd5b84356154e781614c7b565b935060208501359250604085013567ffffffffffffffff811115615509575f5ffd5b61551587828801614f81565b95989497509550505050565b5f5f5f5f5f5f60a08789031215615536575f5ffd5b61553f87614c63565b9550602087013561554f81614c7b565b945060408701359350606087013561556681614d81565b9250608087013567ffffffffffffffff811115615581575f5ffd5b61558d89828a01614f81565b979a9699509497509295939492505050565b5f5f604083850312156155b0575f5ffd5b82356155bb81614c7b565b91506020830135614f7681614d81565b5f5f604083850312156155dc575f5ffd5b6155e583614c63565b91506155f360208401614c63565b90509250929050565b5f5f5f5f5f5f5f5f5f5f5f5f6109208d8f031215615618575f5ffd5b6156228e8e6153ee565b9b506156328e6104008f016153ee565b9a506108008d013599506108208d013598506108408d013597506156596108608e01614c63565b96506156696108808e0135614c7b565b6108808d0135955061567e6108a08e01614c63565b94506108c08d013561568f81614c7b565b93506108e08d0135925067ffffffffffffffff6109008e013511156156b2575f5ffd5b6156c38e6109008f01358f01614f81565b81935080925050509295989b509295989b509295989b565b5f5f5f5f5f5f5f60c0888a0312156156f1575f5ffd5b6156fa88614c63565b9650602088013561570a81614c7b565b955060408801359450606088013561572181614c7b565b9350608088013561573181614d81565b925060a088013567ffffffffffffffff81111561574c575f5ffd5b6157588a828b01614f81565b989b979a50959850939692959293505050565b5f5f5f5f5f5f60c08789031215615780575f5ffd5b61578987614c63565b9550602087013561579981614c7b565b94506157a760408801614c63565b935060608701356157b781614c7b565b925060808701356157c781614c7b565b915060a087013567ffffffffffffffff8111156157e2575f5ffd5b6157ee89828a01614d2f565b9150509295509295509295565b5f5f5f5f610460858703121561580f575f5ffd5b8435935061582086602087016153ee565b925061582f6104208601614c63565b939692955092936104400135925050565b5f5f5f5f6108408587031215615854575f5ffd5b8435935061586586602087016153ee565b9250610420850135915061587d8661044087016153ee565b905092959194509250565b600181811c9082168061589c57607f821691505b6020821081036158ba57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115611db257805f5260205f20601f840160051c810160208510156158e55750805b601f840160051c820191505b81811015611702575f81556001016158f1565b815167ffffffffffffffff81111561591e5761591e614c9a565b6159328161592c8454615888565b846158c0565b6020601f821160018114615964575f831561594d5750848201515b5f19600385901b1c1916600184901b178455611702565b5f84815260208120601f198516915b828110156159935787850151825560209485019460019092019101615973565b50848210156159b057868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f602082840312156159cf575f5ffd5b815161206d81614c7b565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b63ffffffff8181168382160190811115611693576116936159ee565b80820180821115611693576116936159ee565b5f615a3e614d4c84614d08565b9050828152838383011115615a51575f5ffd5b61206d83602083018461503b565b5f82601f830112615a6e575f5ffd5b61206d83835160208501615a31565b5f60208284031215615a8d575f5ffd5b815167ffffffffffffffff811115615aa3575f5ffd5b614f4484828501615a5f565b5f60208284031215615abf575f5ffd5b815161206d816150b5565b606081525f615adc606083018661505d565b8281036020840152615aee818661505d565b91505060ff83166040830152949350505050565b5f60e0828403128015615b13575f5ffd5b50615b1c614cae565b8235615b27816150b5565b8152615b3560208401614c63565b60208201526040830135615b4881614c7b565b6040820152615b5960608401614c63565b60608201526080830135615b6c81614c7b565b608082015260a0838101359082015260c0928301359281019290925250919050565b5f60208284031215615b9e575f5ffd5b815167ffffffffffffffff811115615bb4575f5ffd5b8201601f81018413615bc4575f5ffd5b614f4484825160208401615a31565b8082028115828204841417611693576116936159ee565b818382375f9101908152919050565b5f8251615c0a81846020870161503b565b9190910192915050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b63ffffffff861681526001600160a01b03851660208201526001600160a01b0384166040820152608060608201525f615c79608083018486615c14565b979650505050505050565b5f60208284031215615c94575f5ffd5b5051919050565b81810381811115611693576116936159ee565b60ff8916815263ffffffff881660208201526001600160a01b038716604082015263ffffffff861660608201526001600160a01b03851660808201528360a082015261010060c08201525f615d0761010083018561505d565b905063ffffffff831660e08301529998505050505050505050565b6001600160a01b038516815263ffffffff84166020820152606060408201525f614bd1606083018486615c14565b5f5f5f60608486031215615d62575f5ffd5b835167ffffffffffffffff811115615d78575f5ffd5b615d8486828701615a5f565b935050602084015167ffffffffffffffff811115615da0575f5ffd5b615dac86828701615a5f565b9250506040840151615357816150b5565b60ff8a16815263ffffffff891660208201526001600160a01b038816604082015263ffffffff871660608201526001600160a01b03861660808201528460a082015261010060c08201525f615e1761010083018587615c14565b905063ffffffff831660e08301529a9950505050505050505050565b6001600160a01b0386168152608060208201525f615e55608083018688615c14565b6001600160a01b03948516604084015292909316606090910152949350505050565b5f60208284031215615e87575f5ffd5b815161206d81614d81565b5f60018201615ea357615ea36159ee565b5060010190565b6001815b6001841115615ee557808504811115615ec957615ec96159ee565b6001841615615ed757908102905b60019390931c928002615eae565b935093915050565b5f82615efb57506001611693565b81615f0757505f611693565b8160018114615f1d5760028114615f2757615f43565b6001915050611693565b60ff841115615f3857615f386159ee565b50506001821b611693565b5060208310610133831016604e8410600b8410161715615f66575081810a611693565b615f725f198484615eaa565b805f1904821115615f8557615f856159ee565b029392505050565b5f61206d8383615eed565b634e487b7160e01b5f52600160045260245ffdfea26469706673582212207243a40aa417ed0244c6a94dec7180c76afaf6b59a83039e261a0e7c6db751c864736f6c634300081c00336080604052348015600e575f5ffd5b50610fb38061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610029575f3560e01c8063c514f24e1461002d575b5f5ffd5b61003561004b565b604051610042919061006a565b60405180910390f35b60405180610f000160405280610ec881526020016100b6610ec8913981565b602081525f82518060208401525f5b818110156100965760208186018101516040868401015201610079565b505f604082850101526040601f19601f8301168401019150509291505056fe60806040819052631d97f74d60e11b81523390633b2fee9a90608490602090600481865afa158015610033573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100579190610433565b604080515f80825260208201909252905061007382825f6100e2565b50506100dd336001600160a01b03166338b8fbbb6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100b4573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100d89190610433565b61010d565b6104c8565b6100eb8361017a565b5f825111806100f75750805b156101085761010683836101b9565b505b505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61014c5f516020610e815f395f51905f52546001600160a01b031690565b604080516001600160a01b03928316815291841660208301520160405180910390a1610177816101e5565b50565b61018381610280565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606101de8383604051806060016040528060278152602001610ea160279139610314565b9392505050565b6001600160a01b03811661024f5760405162461bcd60e51b815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b805f516020610e815f395f51905f525b80546001600160a01b0319166001600160a01b039290921691909117905550565b6001600160a01b0381163b6102ed5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610246565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc61025f565b60605f5f856001600160a01b031685604051610330919061047b565b5f60405180830381855af49150503d805f8114610368576040519150601f19603f3d011682016040523d82523d5f602084013e61036d565b606091505b50909250905061037f86838387610389565b9695505050505050565b606083156103f75782515f036103f0576001600160a01b0385163b6103f05760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610246565b5081610401565b6104018383610409565b949350505050565b8151156104195781518083602001fd5b8060405162461bcd60e51b81526004016102469190610496565b5f60208284031215610443575f5ffd5b81516001600160a01b03811681146101de575f5ffd5b5f5b8381101561047357818101518382015260200161045b565b50505f910152565b5f825161048c818460208701610459565b9190910192915050565b602081525f82518060208401526104b4816040850160208701610459565b601f01601f19169190910160400192915050565b6109ac806104d55f395ff3fe60806040526004361061005d575f3560e01c80635c60da1b116100425780635c60da1b146100a65780638f283970146100e3578063f851a440146101025761006c565b80633659cfe6146100745780634f1ef286146100935761006c565b3661006c5761006a610116565b005b61006a610116565b34801561007f575f5ffd5b5061006a61008e366004610854565b610130565b61006a6100a136600461086d565b610178565b3480156100b1575f5ffd5b506100ba6101eb565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b3480156100ee575f5ffd5b5061006a6100fd366004610854565b610228565b34801561010d575f5ffd5b506100ba610255565b61011e610282565b61012e610129610359565b610362565b565b610138610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101705761016d8160405180602001604052805f8152505f6103bf565b50565b61016d610116565b610180610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101e3576101de8383838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250600192506103bf915050565b505050565b6101de610116565b5f6101f4610380565b73ffffffffffffffffffffffffffffffffffffffff16330361021d57610218610359565b905090565b610225610116565b90565b610230610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101705761016d816103e9565b5f61025e610380565b73ffffffffffffffffffffffffffffffffffffffff16330361021d57610218610380565b61028a610380565b73ffffffffffffffffffffffffffffffffffffffff16330361012e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604260248201527f5472616e73706172656e745570677261646561626c6550726f78793a2061646d60448201527f696e2063616e6e6f742066616c6c6261636b20746f2070726f7879207461726760648201527f6574000000000000000000000000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b5f61021861044a565b365f5f375f5f365f845af43d5f5f3e80801561037c573d5ff35b3d5ffd5b5f7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b5473ffffffffffffffffffffffffffffffffffffffff16919050565b6103c883610471565b5f825111806103d45750805b156101de576103e383836104bd565b50505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f610412610380565b6040805173ffffffffffffffffffffffffffffffffffffffff928316815291841660208301520160405180910390a161016d816104e9565b5f7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6103a3565b61047a816105f5565b60405173ffffffffffffffffffffffffffffffffffffffff8216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606104e28383604051806060016040528060278152602001610979602791396106c0565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff811661058c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610350565b807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b73ffffffffffffffffffffffffffffffffffffffff81163b610699576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201527f6f74206120636f6e7472616374000000000000000000000000000000000000006064820152608401610350565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6105af565b60605f5f8573ffffffffffffffffffffffffffffffffffffffff16856040516106e9919061090d565b5f60405180830381855af49150503d805f8114610721576040519150601f19603f3d011682016040523d82523d5f602084013e610726565b606091505b509150915061073786838387610741565b9695505050505050565b606083156107d65782515f036107cf5773ffffffffffffffffffffffffffffffffffffffff85163b6107cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610350565b50816107e0565b6107e083836107e8565b949350505050565b8151156107f85781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103509190610928565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f575f5ffd5b919050565b5f60208284031215610864575f5ffd5b6104e28261082c565b5f5f5f6040848603121561087f575f5ffd5b6108888461082c565b9250602084013567ffffffffffffffff8111156108a3575f5ffd5b8401601f810186136108b3575f5ffd5b803567ffffffffffffffff8111156108c9575f5ffd5b8660208284010111156108da575f5ffd5b939660209190910195509293505050565b5f5b838110156109055781810151838201526020016108ed565b50505f910152565b5f825161091e8184602087016108eb565b9190910192915050565b602081525f82518060208401526109468160408501602087016108eb565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a164736f6c634300081c000ab53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220b0d63e282eb51974e150b528a8c2f62ea4af5ce8758d6b8a36f9a4fadec6221664736f6c634300081c00336080604052348015600e575f5ffd5b5060156019565b60c9565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560685760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c65780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6116fb806100d65f395ff3fe608060405234801561000f575f5ffd5b5060043610610115575f3560e01c806370a08231116100ad5780639dc29fac1161007d578063a9059cbb11610063578063a9059cbb146102a6578063d505accf146102b9578063dd62ed3e146102cc575f5ffd5b80639dc29fac1461024b578063a3c573eb1461025e575f5ffd5b806370a08231146102025780637ecebe001461021557806384b0196e1461022857806395d89b4114610243575f5ffd5b806323b872dd116100e857806323b872dd146101a0578063313ce567146101b35780633644e515146101e757806340c10f19146101ef575f5ffd5b806306fdde0314610119578063095ea7b3146101375780631624f6c61461015a57806318160ddd1461016f575b5f5ffd5b610121610323565b60405161012e919061125c565b60405180910390f35b61014a610145366004611290565b6103db565b604051901515815260200161012e565b61016d610168366004611367565b6103f4565b005b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02545b60405190815260200161012e565b61014a6101ae3660046113db565b610576565b7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d005460405160ff909116815260200161012e565b610192610599565b61016d6101fd366004611290565b6105a7565b610192610210366004611415565b61060a565b610192610223366004611415565b61064d565b610230610657565b60405161012e979695949392919061142e565b610121610720565b61016d610259366004611290565b610771565b7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d005461010090046001600160a01b03166040516001600160a01b03909116815260200161012e565b61014a6102b4366004611290565b6107cf565b61016d6102c73660046114c4565b6107dc565b6101926102da36600461152a565b6001600160a01b039182165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b60605f7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b90508060030180546103599061155b565b80601f01602080910402602001604051908101604052809291908181526020018280546103859061155b565b80156103d05780601f106103a7576101008083540402835291602001916103d0565b820191905f5260205f20905b8154815290600101906020018083116103b357829003601f168201915b505050505091505090565b5f336103e8818585610931565b60019150505b92915050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f8115801561043e5750825b90505f8267ffffffffffffffff16600114801561045a5750303b155b905081158015610468575080155b156104865760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156104ba57845468ff00000000000000001916680100000000000000001785555b6104c4888861093e565b6104cd88610954565b7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d00805460ff88167fffffffffffffffffffffff000000000000000000000000000000000000000000909116176101003302179055831561056c57845468ff000000000000000019168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b5f33610583858285610982565b61058e858585610a1c565b506001949350505050565b5f6105a2610a79565b905090565b5f7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d00805490915061010090046001600160a01b031633146105fb576040516338da3b1560e01b815260040160405180910390fd5b6106058383610a82565b505050565b5f807f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b6001600160a01b039093165f9081526020939093525050604090205490565b5f6103ee82610ab6565b5f60608082808083817fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100805490915015801561069557506001810154155b6106e65760405162461bcd60e51b815260206004820152601560248201527f4549503731323a20556e696e697469616c697a6564000000000000000000000060448201526064015b60405180910390fd5b6106ee610ac0565b6106f6610b11565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060917f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00916103599061155b565b5f7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d00805490915061010090046001600160a01b031633146107c5576040516338da3b1560e01b815260040160405180910390fd5b6106058383610b3a565b5f336103e8818585610a1c565b834211156108005760405163313c898160e11b8152600481018590526024016106dd565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c988888861086a8c6001600160a01b03165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f6108c482610b6e565b90505f6108d382878787610b9a565b9050896001600160a01b0316816001600160a01b03161461091a576040516325c0072360e11b81526001600160a01b0380831660048301528b1660248201526044016106dd565b6109258a8a8a610931565b50505050505050505050565b6106058383836001610bc6565b610946610cbd565b6109508282610d0d565b5050565b61095c610cbd565b61097f81604051806040016040528060018152602001603160f81b815250610d70565b50565b6001600160a01b038381165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0160209081526040808320938616835292905220545f198114610a165781811015610a0857604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064016106dd565b610a1684848484035f610bc6565b50505050565b6001600160a01b038316610a4557604051634b637e8f60e11b81525f60048201526024016106dd565b6001600160a01b038216610a6e5760405163ec442f0560e01b81525f60048201526024016106dd565b610605838383610de2565b5f6105a2610f2e565b6001600160a01b038216610aab5760405163ec442f0560e01b81525f60048201526024016106dd565b6109505f8383610de2565b5f6103ee82610fa1565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060917fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100916103599061155b565b60605f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100610348565b6001600160a01b038216610b6357604051634b637e8f60e11b81525f60048201526024016106dd565b610950825f83610de2565b5f6103ee610b7a610a79565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f610baa88888888610fc9565b925092509250610bba8282611091565b50909695505050505050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace006001600160a01b038516610c105760405163e602df0560e01b81525f60048201526024016106dd565b6001600160a01b038416610c3957604051634a1406b160e11b81525f60048201526024016106dd565b6001600160a01b038086165f90815260018301602090815260408083209388168352929052208390558115610cb657836001600160a01b0316856001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92585604051610cad91815260200190565b60405180910390a35b5050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16610d0b57604051631afcd79f60e31b815260040160405180910390fd5b565b610d15610cbd565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace007f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace03610d6184826115d7565b5060048101610a1683826115d7565b610d78610cbd565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1007fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102610dc484826115d7565b5060038101610dd383826115d7565b505f8082556001909101555050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace006001600160a01b038416610e2f5781816002015f828254610e249190611692565b90915550610e9f9050565b6001600160a01b0384165f9081526020829052604090205482811015610e815760405163391434e360e21b81526001600160a01b038616600482015260248101829052604481018490526064016106dd565b6001600160a01b0385165f9081526020839052604090209083900390555b6001600160a01b038316610ebd576002810180548390039055610edb565b6001600160a01b0383165f9081526020829052604090208054830190555b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610f2091815260200190565b60405180910390a350505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f610f58611149565b610f606111c4565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f807f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0061062e565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561100257505f91506003905082611087565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015611053573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b03811661107e57505f925060019150829050611087565b92505f91508190505b9450945094915050565b5f8260038111156110a4576110a46116b1565b036110ad575050565b60018260038111156110c1576110c16116b1565b036110df5760405163f645eedf60e01b815260040160405180910390fd5b60028260038111156110f3576110f36116b1565b036111145760405163fce698f760e01b8152600481018290526024016106dd565b6003826003811115611128576111286116b1565b03610950576040516335e2f38360e21b8152600481018290526024016106dd565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10081611174610ac0565b80519091501561118c57805160209091012092915050565b8154801561119b579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100816111ef610b11565b80519091501561120757805160209091012092915050565b6001820154801561119b579392505050565b5f81518084525f5b8181101561123d57602081850181015186830182015201611221565b505f602082860101526020601f19601f83011685010191505092915050565b602081525f61126e6020830184611219565b9392505050565b80356001600160a01b038116811461128b575f5ffd5b919050565b5f5f604083850312156112a1575f5ffd5b6112aa83611275565b946020939093013593505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f8301126112db575f5ffd5b813567ffffffffffffffff8111156112f5576112f56112b8565b604051601f8201601f19908116603f0116810167ffffffffffffffff81118282101715611324576113246112b8565b60405281815283820160200185101561133b575f5ffd5b816020850160208301375f918101602001919091529392505050565b803560ff8116811461128b575f5ffd5b5f5f5f60608486031215611379575f5ffd5b833567ffffffffffffffff81111561138f575f5ffd5b61139b868287016112cc565b935050602084013567ffffffffffffffff8111156113b7575f5ffd5b6113c3868287016112cc565b9250506113d260408501611357565b90509250925092565b5f5f5f606084860312156113ed575f5ffd5b6113f684611275565b925061140460208501611275565b929592945050506040919091013590565b5f60208284031215611425575f5ffd5b61126e82611275565b60ff60f81b8816815260e060208201525f61144c60e0830189611219565b828103604084015261145e8189611219565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b818110156114b3578351835260209384019390920191600101611495565b50909b9a5050505050505050505050565b5f5f5f5f5f5f5f60e0888a0312156114da575f5ffd5b6114e388611275565b96506114f160208901611275565b9550604088013594506060880135935061150d60808901611357565b9699959850939692959460a0840135945060c09093013592915050565b5f5f6040838503121561153b575f5ffd5b61154483611275565b915061155260208401611275565b90509250929050565b600181811c9082168061156f57607f821691505b60208210810361158d57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561060557805f5260205f20601f840160051c810160208510156115b85750805b601f840160051c820191505b81811015610cb6575f81556001016115c4565b815167ffffffffffffffff8111156115f1576115f16112b8565b611605816115ff845461155b565b84611593565b6020601f821160018114611637575f83156116205750848201515b5f19600385901b1c1916600184901b178455610cb6565b5f84815260208120601f198516915b828110156116665787850151825560209485019460019092019101611646565b508482101561168357868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b808201808211156103ee57634e487b7160e01b5f52601160045260245ffd5b634e487b7160e01b5f52602160045260245ffdfea2646970667358221220a561c2f5d53554a9eb5f9ce0194d9c8147ec03d3ba1e77539eb6b379627db98d64736f6c634300081c00336080604052348015600e575f5ffd5b50610d418061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610064575f3560e01c8063be3dcf621161004d578063be3dcf62146100b4578063c00f14ab146100d9578063cf825e55146100ec575f5ffd5b8063737ce34114610068578063a28fa4a314610091575b5f5ffd5b61007b61007636600461092b565b6100ff565b604051610088919061099a565b60405180910390f35b6100a461009f3660046109ac565b6101d8565b6040519015158152602001610088565b6100c76100c236600461092b565b61057c565b60405160ff9091168152602001610088565b61007b6100e736600461092b565b610630565b61007b6100fa36600461092b565b610675565b60408051600481526024810182526020810180516001600160e01b03166395d89b4160e01b17905290516060915f9182916001600160a01b038616916101459190610a4f565b5f60405180830381855afa9150503d805f811461017d576040519150601f19603f3d011682016040523d82523d5f602084013e610182565b606091505b5091509150816101c7576040518060400160405280600981526020017f4e4f5f53594d424f4c00000000000000000000000000000000000000000000008152506101d0565b6101d08161073d565b949350505050565b5f806101e76004828789610a6a565b6101f091610a91565b9050632afa533160e01b6001600160e01b031982160161039f575f5f5f5f5f5f5f8c8c600490809261022493929190610a6a565b8101906102319190610ad7565b96509650965096509650965096508a6001600160a01b0316876001600160a01b0316146102715760405163912ecce760e01b815260040160405180910390fd5b896001600160a01b0316866001600160a01b0316146102a35760405163750643af60e01b815260040160405180910390fd5b5f8e6001600160a01b031663d505accf60e01b8989898989898960405160240161030f97969594939291906001600160a01b0397881681529590961660208601526040850193909352606084019190915260ff16608083015260a082015260c081019190915260e00190565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b031990941693909317909252905161034d9190610a4f565b5f604051808303815f865af19150503d805f8114610386576040519150601f19603f3d011682016040523d82523d5f602084013e61038b565b606091505b50909a506105739950505050505050505050565b631c0d143d60e21b6001600160e01b031982160161055a575f5f5f5f5f5f5f5f8d8d60049080926103d293929190610a6a565b8101906103df9190610b43565b975097509750975097509750975097508b6001600160a01b0316886001600160a01b0316146104215760405163912ecce760e01b815260040160405180910390fd5b8a6001600160a01b0316876001600160a01b0316146104535760405163750643af60e01b815260040160405180910390fd5b5f8f6001600160a01b0316638fcbaf0c60e01b8a8a8a8a8a8a8a8a6040516024016104c99897969594939291906001600160a01b039889168152969097166020870152604086019490945260608501929092521515608084015260ff1660a083015260c082015260e08101919091526101000190565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516105079190610a4f565b5f604051808303815f865af19150503d805f8114610540576040519150601f19603f3d011682016040523d82523d5f602084013e610545565b606091505b50909b506105739a5050505050505050505050565b604051637141605d60e11b815260040160405180910390fd5b95945050505050565b60408051600481526024810182526020810180516001600160e01b031663313ce56760e01b17905290515f91829182916001600160a01b038616916105c19190610a4f565b5f60405180830381855afa9150503d805f81146105f9576040519150601f19603f3d011682016040523d82523d5f602084013e6105fe565b606091505b5091509150818015610611575080516020145b61061c5760126101d0565b808060200190518101906101d09190610bc5565b606061063b82610675565b610644836100ff565b61064d8461057c565b60405160200161065f93929190610be0565b6040516020818303038152906040529050919050565b60408051600481526024810182526020810180516001600160e01b03166306fdde0360e01b17905290516060915f9182916001600160a01b038616916106bb9190610a4f565b5f60405180830381855afa9150503d805f81146106f3576040519150601f19603f3d011682016040523d82523d5f602084013e6106f8565b606091505b5091509150816101c7576040518060400160405280600781526020017f4e4f5f4e414d45000000000000000000000000000000000000000000000000008152506101d0565b60606040825110610762578180602001905181019061075c9190610c2c565b92915050565b81516020036108cb575f5b6020811080156107b4575082818151811061078a5761078a610cd3565b01602001517fff000000000000000000000000000000000000000000000000000000000000001615155b156107cb57806107c381610ce7565b91505061076d565b805f0361080d57505060408051808201909152601281527f4e4f545f56414c49445f454e434f44494e4700000000000000000000000000006020820152919050565b5f8167ffffffffffffffff81111561082757610827610c18565b6040519080825280601f01601f191660200182016040528015610851576020820181803683370190505b5090505f5b828110156108c35784818151811061087057610870610cd3565b602001015160f81c60f81b82828151811061088d5761088d610cd3565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600101610856565b509392505050565b505060408051808201909152601281527f4e4f545f56414c49445f454e434f44494e470000000000000000000000000000602082015290565b919050565b6001600160a01b038116811461091d575f5ffd5b50565b803561090481610909565b5f6020828403121561093b575f5ffd5b813561094681610909565b9392505050565b5f5b8381101561096757818101518382015260200161094f565b50505f910152565b5f815180845261098681602086016020860161094d565b601f01601f19169290920160200192915050565b602081525f610946602083018461096f565b5f5f5f5f5f608086880312156109c0575f5ffd5b85356109cb81610909565b9450602086013567ffffffffffffffff8111156109e6575f5ffd5b8601601f810188136109f6575f5ffd5b803567ffffffffffffffff811115610a0c575f5ffd5b886020828401011115610a1d575f5ffd5b602091909101945092506040860135610a3581610909565b9150610a4360608701610920565b90509295509295909350565b5f8251610a6081846020870161094d565b9190910192915050565b5f5f85851115610a78575f5ffd5b83861115610a84575f5ffd5b5050820193919092039150565b80356001600160e01b03198116906004841015610ac2576001600160e01b0319600485900360031b81901b82161691505b5092915050565b60ff8116811461091d575f5ffd5b5f5f5f5f5f5f5f60e0888a031215610aed575f5ffd5b8735610af881610909565b96506020880135610b0881610909565b955060408801359450606088013593506080880135610b2681610ac9565b9699959850939692959460a0840135945060c09093013592915050565b5f5f5f5f5f5f5f5f610100898b031215610b5b575f5ffd5b8835610b6681610909565b97506020890135610b7681610909565b9650604089013595506060890135945060808901358015158114610b98575f5ffd5b935060a0890135610ba881610ac9565b979a969950949793969295929450505060c08201359160e0013590565b5f60208284031215610bd5575f5ffd5b815161094681610ac9565b606081525f610bf2606083018661096f565b8281036020840152610c04818661096f565b91505060ff83166040830152949350505050565b634e487b7160e01b5f52604160045260245ffd5b5f60208284031215610c3c575f5ffd5b815167ffffffffffffffff811115610c52575f5ffd5b8201601f81018413610c62575f5ffd5b805167ffffffffffffffff811115610c7c57610c7c610c18565b604051601f8201601f19908116603f0116810167ffffffffffffffff81118282101715610cab57610cab610c18565b604052818152828201602001861015610cc2575f5ffd5b61057382602083016020860161094d565b634e487b7160e01b5f52603260045260245ffd5b5f60018201610d0457634e487b7160e01b5f52601160045260245ffd5b506001019056fea2646970667358221220d40b6abc187e86e6ec18c2c69585102da3f507e6f6b6827a03363037a0d1de8164736f6c634300081c0033",
}

// Bridgel2sovereignchainABI is the input ABI used to generate the binding from.
// Deprecated: Use Bridgel2sovereignchainMetaData.ABI instead.
var Bridgel2sovereignchainABI = Bridgel2sovereignchainMetaData.ABI

// Bridgel2sovereignchainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Bridgel2sovereignchainMetaData.Bin instead.
var Bridgel2sovereignchainBin = Bridgel2sovereignchainMetaData.Bin

// DeployBridgel2sovereignchain deploys a new Ethereum contract, binding an instance of Bridgel2sovereignchain to it.
func DeployBridgel2sovereignchain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bridgel2sovereignchain, error) {
	parsed, err := Bridgel2sovereignchainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Bridgel2sovereignchainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridgel2sovereignchain{Bridgel2sovereignchainCaller: Bridgel2sovereignchainCaller{contract: contract}, Bridgel2sovereignchainTransactor: Bridgel2sovereignchainTransactor{contract: contract}, Bridgel2sovereignchainFilterer: Bridgel2sovereignchainFilterer{contract: contract}}, nil
}

// Bridgel2sovereignchain is an auto generated Go binding around an Ethereum contract.
type Bridgel2sovereignchain struct {
	Bridgel2sovereignchainCaller     // Read-only binding to the contract
	Bridgel2sovereignchainTransactor // Write-only binding to the contract
	Bridgel2sovereignchainFilterer   // Log filterer for contract events
}

// Bridgel2sovereignchainCaller is an auto generated read-only Go binding around an Ethereum contract.
type Bridgel2sovereignchainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bridgel2sovereignchainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Bridgel2sovereignchainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bridgel2sovereignchainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Bridgel2sovereignchainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bridgel2sovereignchainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Bridgel2sovereignchainSession struct {
	Contract     *Bridgel2sovereignchain // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Bridgel2sovereignchainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Bridgel2sovereignchainCallerSession struct {
	Contract *Bridgel2sovereignchainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// Bridgel2sovereignchainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Bridgel2sovereignchainTransactorSession struct {
	Contract     *Bridgel2sovereignchainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// Bridgel2sovereignchainRaw is an auto generated low-level Go binding around an Ethereum contract.
type Bridgel2sovereignchainRaw struct {
	Contract *Bridgel2sovereignchain // Generic contract binding to access the raw methods on
}

// Bridgel2sovereignchainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Bridgel2sovereignchainCallerRaw struct {
	Contract *Bridgel2sovereignchainCaller // Generic read-only contract binding to access the raw methods on
}

// Bridgel2sovereignchainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Bridgel2sovereignchainTransactorRaw struct {
	Contract *Bridgel2sovereignchainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgel2sovereignchain creates a new instance of Bridgel2sovereignchain, bound to a specific deployed contract.
func NewBridgel2sovereignchain(address common.Address, backend bind.ContractBackend) (*Bridgel2sovereignchain, error) {
	contract, err := bindBridgel2sovereignchain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchain{Bridgel2sovereignchainCaller: Bridgel2sovereignchainCaller{contract: contract}, Bridgel2sovereignchainTransactor: Bridgel2sovereignchainTransactor{contract: contract}, Bridgel2sovereignchainFilterer: Bridgel2sovereignchainFilterer{contract: contract}}, nil
}

// NewBridgel2sovereignchainCaller creates a new read-only instance of Bridgel2sovereignchain, bound to a specific deployed contract.
func NewBridgel2sovereignchainCaller(address common.Address, caller bind.ContractCaller) (*Bridgel2sovereignchainCaller, error) {
	contract, err := bindBridgel2sovereignchain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainCaller{contract: contract}, nil
}

// NewBridgel2sovereignchainTransactor creates a new write-only instance of Bridgel2sovereignchain, bound to a specific deployed contract.
func NewBridgel2sovereignchainTransactor(address common.Address, transactor bind.ContractTransactor) (*Bridgel2sovereignchainTransactor, error) {
	contract, err := bindBridgel2sovereignchain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainTransactor{contract: contract}, nil
}

// NewBridgel2sovereignchainFilterer creates a new log filterer instance of Bridgel2sovereignchain, bound to a specific deployed contract.
func NewBridgel2sovereignchainFilterer(address common.Address, filterer bind.ContractFilterer) (*Bridgel2sovereignchainFilterer, error) {
	contract, err := bindBridgel2sovereignchain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainFilterer{contract: contract}, nil
}

// bindBridgel2sovereignchain binds a generic wrapper to an already deployed contract.
func bindBridgel2sovereignchain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Bridgel2sovereignchainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridgel2sovereignchain *Bridgel2sovereignchainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridgel2sovereignchain.Contract.Bridgel2sovereignchainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridgel2sovereignchain *Bridgel2sovereignchainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.Bridgel2sovereignchainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridgel2sovereignchain *Bridgel2sovereignchainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.Bridgel2sovereignchainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridgel2sovereignchain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.contract.Transact(opts, method, params...)
}

// BRIDGESOVEREIGNVERSION is a free data retrieval call binding the contract method 0xf67566e4.
//
// Solidity: function BRIDGE_SOVEREIGN_VERSION() view returns(string)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) BRIDGESOVEREIGNVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "BRIDGE_SOVEREIGN_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BRIDGESOVEREIGNVERSION is a free data retrieval call binding the contract method 0xf67566e4.
//
// Solidity: function BRIDGE_SOVEREIGN_VERSION() view returns(string)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) BRIDGESOVEREIGNVERSION() (string, error) {
	return _Bridgel2sovereignchain.Contract.BRIDGESOVEREIGNVERSION(&_Bridgel2sovereignchain.CallOpts)
}

// BRIDGESOVEREIGNVERSION is a free data retrieval call binding the contract method 0xf67566e4.
//
// Solidity: function BRIDGE_SOVEREIGN_VERSION() view returns(string)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) BRIDGESOVEREIGNVERSION() (string, error) {
	return _Bridgel2sovereignchain.Contract.BRIDGESOVEREIGNVERSION(&_Bridgel2sovereignchain.CallOpts)
}

// BRIDGEVERSION is a free data retrieval call binding the contract method 0x65d6f654.
//
// Solidity: function BRIDGE_VERSION() view returns(string)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) BRIDGEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "BRIDGE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BRIDGEVERSION is a free data retrieval call binding the contract method 0x65d6f654.
//
// Solidity: function BRIDGE_VERSION() view returns(string)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) BRIDGEVERSION() (string, error) {
	return _Bridgel2sovereignchain.Contract.BRIDGEVERSION(&_Bridgel2sovereignchain.CallOpts)
}

// BRIDGEVERSION is a free data retrieval call binding the contract method 0x65d6f654.
//
// Solidity: function BRIDGE_VERSION() view returns(string)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) BRIDGEVERSION() (string, error) {
	return _Bridgel2sovereignchain.Contract.BRIDGEVERSION(&_Bridgel2sovereignchain.CallOpts)
}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() view returns(bytes)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) INITBYTECODETRANSPARENTPROXY(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "INIT_BYTECODE_TRANSPARENT_PROXY")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() view returns(bytes)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) INITBYTECODETRANSPARENTPROXY() ([]byte, error) {
	return _Bridgel2sovereignchain.Contract.INITBYTECODETRANSPARENTPROXY(&_Bridgel2sovereignchain.CallOpts)
}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() view returns(bytes)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) INITBYTECODETRANSPARENTPROXY() ([]byte, error) {
	return _Bridgel2sovereignchain.Contract.INITBYTECODETRANSPARENTPROXY(&_Bridgel2sovereignchain.CallOpts)
}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) WETHToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "WETHToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) WETHToken() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.WETHToken(&_Bridgel2sovereignchain.CallOpts)
}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) WETHToken() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.WETHToken(&_Bridgel2sovereignchain.CallOpts)
}

// BridgeLib is a free data retrieval call binding the contract method 0x6f0bc3da.
//
// Solidity: function bridgeLib() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) BridgeLib(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "bridgeLib")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeLib is a free data retrieval call binding the contract method 0x6f0bc3da.
//
// Solidity: function bridgeLib() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) BridgeLib() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.BridgeLib(&_Bridgel2sovereignchain.CallOpts)
}

// BridgeLib is a free data retrieval call binding the contract method 0x6f0bc3da.
//
// Solidity: function bridgeLib() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) BridgeLib() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.BridgeLib(&_Bridgel2sovereignchain.CallOpts)
}

// BridgeManager is a free data retrieval call binding the contract method 0x14cc01a0.
//
// Solidity: function bridgeManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) BridgeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "bridgeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeManager is a free data retrieval call binding the contract method 0x14cc01a0.
//
// Solidity: function bridgeManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) BridgeManager() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.BridgeManager(&_Bridgel2sovereignchain.CallOpts)
}

// BridgeManager is a free data retrieval call binding the contract method 0x14cc01a0.
//
// Solidity: function bridgeManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) BridgeManager() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.BridgeManager(&_Bridgel2sovereignchain.CallOpts)
}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) CalculateRoot(opts *bind.CallOpts, leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "calculateRoot", leafHash, smtProof, index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) CalculateRoot(leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	return _Bridgel2sovereignchain.Contract.CalculateRoot(&_Bridgel2sovereignchain.CallOpts, leafHash, smtProof, index)
}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) CalculateRoot(leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	return _Bridgel2sovereignchain.Contract.CalculateRoot(&_Bridgel2sovereignchain.CallOpts, leafHash, smtProof, index)
}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) ClaimedBitMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "claimedBitMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) ClaimedBitMap(arg0 *big.Int) (*big.Int, error) {
	return _Bridgel2sovereignchain.Contract.ClaimedBitMap(&_Bridgel2sovereignchain.CallOpts, arg0)
}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) ClaimedBitMap(arg0 *big.Int) (*big.Int, error) {
	return _Bridgel2sovereignchain.Contract.ClaimedBitMap(&_Bridgel2sovereignchain.CallOpts, arg0)
}

// ClaimedGlobalIndexHashChain is a free data retrieval call binding the contract method 0x1d081d8c.
//
// Solidity: function claimedGlobalIndexHashChain() view returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) ClaimedGlobalIndexHashChain(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "claimedGlobalIndexHashChain")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ClaimedGlobalIndexHashChain is a free data retrieval call binding the contract method 0x1d081d8c.
//
// Solidity: function claimedGlobalIndexHashChain() view returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) ClaimedGlobalIndexHashChain() ([32]byte, error) {
	return _Bridgel2sovereignchain.Contract.ClaimedGlobalIndexHashChain(&_Bridgel2sovereignchain.CallOpts)
}

// ClaimedGlobalIndexHashChain is a free data retrieval call binding the contract method 0x1d081d8c.
//
// Solidity: function claimedGlobalIndexHashChain() view returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) ClaimedGlobalIndexHashChain() ([32]byte, error) {
	return _Bridgel2sovereignchain.Contract.ClaimedGlobalIndexHashChain(&_Bridgel2sovereignchain.CallOpts)
}

// ComputeTokenProxyAddress is a free data retrieval call binding the contract method 0xf214e161.
//
// Solidity: function computeTokenProxyAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) ComputeTokenProxyAddress(opts *bind.CallOpts, originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "computeTokenProxyAddress", originNetwork, originTokenAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeTokenProxyAddress is a free data retrieval call binding the contract method 0xf214e161.
//
// Solidity: function computeTokenProxyAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) ComputeTokenProxyAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.ComputeTokenProxyAddress(&_Bridgel2sovereignchain.CallOpts, originNetwork, originTokenAddress)
}

// ComputeTokenProxyAddress is a free data retrieval call binding the contract method 0xf214e161.
//
// Solidity: function computeTokenProxyAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) ComputeTokenProxyAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.ComputeTokenProxyAddress(&_Bridgel2sovereignchain.CallOpts, originNetwork, originTokenAddress)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) DepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "depositCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) DepositCount() (*big.Int, error) {
	return _Bridgel2sovereignchain.Contract.DepositCount(&_Bridgel2sovereignchain.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) DepositCount() (*big.Int, error) {
	return _Bridgel2sovereignchain.Contract.DepositCount(&_Bridgel2sovereignchain.CallOpts)
}

// EmergencyBridgePauser is a free data retrieval call binding the contract method 0x2f84c690.
//
// Solidity: function emergencyBridgePauser() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) EmergencyBridgePauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "emergencyBridgePauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmergencyBridgePauser is a free data retrieval call binding the contract method 0x2f84c690.
//
// Solidity: function emergencyBridgePauser() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) EmergencyBridgePauser() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.EmergencyBridgePauser(&_Bridgel2sovereignchain.CallOpts)
}

// EmergencyBridgePauser is a free data retrieval call binding the contract method 0x2f84c690.
//
// Solidity: function emergencyBridgePauser() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) EmergencyBridgePauser() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.EmergencyBridgePauser(&_Bridgel2sovereignchain.CallOpts)
}

// EmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0xae24490a.
//
// Solidity: function emergencyBridgeUnpauser() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) EmergencyBridgeUnpauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "emergencyBridgeUnpauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0xae24490a.
//
// Solidity: function emergencyBridgeUnpauser() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) EmergencyBridgeUnpauser() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.EmergencyBridgeUnpauser(&_Bridgel2sovereignchain.CallOpts)
}

// EmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0xae24490a.
//
// Solidity: function emergencyBridgeUnpauser() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) EmergencyBridgeUnpauser() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.EmergencyBridgeUnpauser(&_Bridgel2sovereignchain.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) GasTokenAddress() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.GasTokenAddress(&_Bridgel2sovereignchain.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) GasTokenAddress() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.GasTokenAddress(&_Bridgel2sovereignchain.CallOpts)
}

// GasTokenMetadata is a free data retrieval call binding the contract method 0x27aef4e8.
//
// Solidity: function gasTokenMetadata() view returns(bytes)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) GasTokenMetadata(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "gasTokenMetadata")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GasTokenMetadata is a free data retrieval call binding the contract method 0x27aef4e8.
//
// Solidity: function gasTokenMetadata() view returns(bytes)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) GasTokenMetadata() ([]byte, error) {
	return _Bridgel2sovereignchain.Contract.GasTokenMetadata(&_Bridgel2sovereignchain.CallOpts)
}

// GasTokenMetadata is a free data retrieval call binding the contract method 0x27aef4e8.
//
// Solidity: function gasTokenMetadata() view returns(bytes)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) GasTokenMetadata() ([]byte, error) {
	return _Bridgel2sovereignchain.Contract.GasTokenMetadata(&_Bridgel2sovereignchain.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) GasTokenNetwork() (uint32, error) {
	return _Bridgel2sovereignchain.Contract.GasTokenNetwork(&_Bridgel2sovereignchain.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) GasTokenNetwork() (uint32, error) {
	return _Bridgel2sovereignchain.Contract.GasTokenNetwork(&_Bridgel2sovereignchain.CallOpts)
}

// GetLeafValue is a free data retrieval call binding the contract method 0x3e197043.
//
// Solidity: function getLeafValue(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes32 metadataHash) pure returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) GetLeafValue(opts *bind.CallOpts, leafType uint8, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadataHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "getLeafValue", leafType, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadataHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLeafValue is a free data retrieval call binding the contract method 0x3e197043.
//
// Solidity: function getLeafValue(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes32 metadataHash) pure returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) GetLeafValue(leafType uint8, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadataHash [32]byte) ([32]byte, error) {
	return _Bridgel2sovereignchain.Contract.GetLeafValue(&_Bridgel2sovereignchain.CallOpts, leafType, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadataHash)
}

// GetLeafValue is a free data retrieval call binding the contract method 0x3e197043.
//
// Solidity: function getLeafValue(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes32 metadataHash) pure returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) GetLeafValue(leafType uint8, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadataHash [32]byte) ([32]byte, error) {
	return _Bridgel2sovereignchain.Contract.GetLeafValue(&_Bridgel2sovereignchain.CallOpts, leafType, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadataHash)
}

// GetProxiedTokensManager is a free data retrieval call binding the contract method 0x38b8fbbb.
//
// Solidity: function getProxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) GetProxiedTokensManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "getProxiedTokensManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxiedTokensManager is a free data retrieval call binding the contract method 0x38b8fbbb.
//
// Solidity: function getProxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) GetProxiedTokensManager() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.GetProxiedTokensManager(&_Bridgel2sovereignchain.CallOpts)
}

// GetProxiedTokensManager is a free data retrieval call binding the contract method 0x38b8fbbb.
//
// Solidity: function getProxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) GetProxiedTokensManager() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.GetProxiedTokensManager(&_Bridgel2sovereignchain.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) GetRoot() ([32]byte, error) {
	return _Bridgel2sovereignchain.Contract.GetRoot(&_Bridgel2sovereignchain.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) GetRoot() ([32]byte, error) {
	return _Bridgel2sovereignchain.Contract.GetRoot(&_Bridgel2sovereignchain.CallOpts)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) GetTokenMetadata(opts *bind.CallOpts, token common.Address) ([]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "getTokenMetadata", token)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) GetTokenMetadata(token common.Address) ([]byte, error) {
	return _Bridgel2sovereignchain.Contract.GetTokenMetadata(&_Bridgel2sovereignchain.CallOpts, token)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) GetTokenMetadata(token common.Address) ([]byte, error) {
	return _Bridgel2sovereignchain.Contract.GetTokenMetadata(&_Bridgel2sovereignchain.CallOpts, token)
}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) GetTokenWrappedAddress(opts *bind.CallOpts, originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "getTokenWrappedAddress", originNetwork, originTokenAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) GetTokenWrappedAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.GetTokenWrappedAddress(&_Bridgel2sovereignchain.CallOpts, originNetwork, originTokenAddress)
}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) GetTokenWrappedAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.GetTokenWrappedAddress(&_Bridgel2sovereignchain.CallOpts, originNetwork, originTokenAddress)
}

// GetWrappedTokenBridgeImplementation is a free data retrieval call binding the contract method 0x3b2fee9a.
//
// Solidity: function getWrappedTokenBridgeImplementation() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) GetWrappedTokenBridgeImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "getWrappedTokenBridgeImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWrappedTokenBridgeImplementation is a free data retrieval call binding the contract method 0x3b2fee9a.
//
// Solidity: function getWrappedTokenBridgeImplementation() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) GetWrappedTokenBridgeImplementation() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.GetWrappedTokenBridgeImplementation(&_Bridgel2sovereignchain.CallOpts)
}

// GetWrappedTokenBridgeImplementation is a free data retrieval call binding the contract method 0x3b2fee9a.
//
// Solidity: function getWrappedTokenBridgeImplementation() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) GetWrappedTokenBridgeImplementation() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.GetWrappedTokenBridgeImplementation(&_Bridgel2sovereignchain.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) GlobalExitRootManager() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.GlobalExitRootManager(&_Bridgel2sovereignchain.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.GlobalExitRootManager(&_Bridgel2sovereignchain.CallOpts)
}

// Initialize0 is a free data retrieval call binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() pure returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) Initialize0(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "initialize0")

	if err != nil {
		return err
	}

	return err

}

// Initialize0 is a free data retrieval call binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() pure returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) Initialize0() error {
	return _Bridgel2sovereignchain.Contract.Initialize0(&_Bridgel2sovereignchain.CallOpts)
}

// Initialize0 is a free data retrieval call binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() pure returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) Initialize0() error {
	return _Bridgel2sovereignchain.Contract.Initialize0(&_Bridgel2sovereignchain.CallOpts)
}

// IsClaimed is a free data retrieval call binding the contract method 0xcc461632.
//
// Solidity: function isClaimed(uint32 leafIndex, uint32 sourceBridgeNetwork) view returns(bool)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) IsClaimed(opts *bind.CallOpts, leafIndex uint32, sourceBridgeNetwork uint32) (bool, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "isClaimed", leafIndex, sourceBridgeNetwork)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClaimed is a free data retrieval call binding the contract method 0xcc461632.
//
// Solidity: function isClaimed(uint32 leafIndex, uint32 sourceBridgeNetwork) view returns(bool)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) IsClaimed(leafIndex uint32, sourceBridgeNetwork uint32) (bool, error) {
	return _Bridgel2sovereignchain.Contract.IsClaimed(&_Bridgel2sovereignchain.CallOpts, leafIndex, sourceBridgeNetwork)
}

// IsClaimed is a free data retrieval call binding the contract method 0xcc461632.
//
// Solidity: function isClaimed(uint32 leafIndex, uint32 sourceBridgeNetwork) view returns(bool)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) IsClaimed(leafIndex uint32, sourceBridgeNetwork uint32) (bool, error) {
	return _Bridgel2sovereignchain.Contract.IsClaimed(&_Bridgel2sovereignchain.CallOpts, leafIndex, sourceBridgeNetwork)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) IsEmergencyState() (bool, error) {
	return _Bridgel2sovereignchain.Contract.IsEmergencyState(&_Bridgel2sovereignchain.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) IsEmergencyState() (bool, error) {
	return _Bridgel2sovereignchain.Contract.IsEmergencyState(&_Bridgel2sovereignchain.CallOpts)
}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) LastUpdatedDepositCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "lastUpdatedDepositCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) LastUpdatedDepositCount() (uint32, error) {
	return _Bridgel2sovereignchain.Contract.LastUpdatedDepositCount(&_Bridgel2sovereignchain.CallOpts)
}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) LastUpdatedDepositCount() (uint32, error) {
	return _Bridgel2sovereignchain.Contract.LastUpdatedDepositCount(&_Bridgel2sovereignchain.CallOpts)
}

// LocalBalanceTree is a free data retrieval call binding the contract method 0xd9cb3aec.
//
// Solidity: function localBalanceTree(bytes32 tokenInfoHash) view returns(uint256 amount)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) LocalBalanceTree(opts *bind.CallOpts, tokenInfoHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "localBalanceTree", tokenInfoHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LocalBalanceTree is a free data retrieval call binding the contract method 0xd9cb3aec.
//
// Solidity: function localBalanceTree(bytes32 tokenInfoHash) view returns(uint256 amount)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) LocalBalanceTree(tokenInfoHash [32]byte) (*big.Int, error) {
	return _Bridgel2sovereignchain.Contract.LocalBalanceTree(&_Bridgel2sovereignchain.CallOpts, tokenInfoHash)
}

// LocalBalanceTree is a free data retrieval call binding the contract method 0xd9cb3aec.
//
// Solidity: function localBalanceTree(bytes32 tokenInfoHash) view returns(uint256 amount)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) LocalBalanceTree(tokenInfoHash [32]byte) (*big.Int, error) {
	return _Bridgel2sovereignchain.Contract.LocalBalanceTree(&_Bridgel2sovereignchain.CallOpts, tokenInfoHash)
}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) NetworkID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "networkID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) NetworkID() (uint32, error) {
	return _Bridgel2sovereignchain.Contract.NetworkID(&_Bridgel2sovereignchain.CallOpts)
}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) NetworkID() (uint32, error) {
	return _Bridgel2sovereignchain.Contract.NetworkID(&_Bridgel2sovereignchain.CallOpts)
}

// PendingEmergencyBridgePauser is a free data retrieval call binding the contract method 0x03e6e116.
//
// Solidity: function pendingEmergencyBridgePauser() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) PendingEmergencyBridgePauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "pendingEmergencyBridgePauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingEmergencyBridgePauser is a free data retrieval call binding the contract method 0x03e6e116.
//
// Solidity: function pendingEmergencyBridgePauser() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) PendingEmergencyBridgePauser() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.PendingEmergencyBridgePauser(&_Bridgel2sovereignchain.CallOpts)
}

// PendingEmergencyBridgePauser is a free data retrieval call binding the contract method 0x03e6e116.
//
// Solidity: function pendingEmergencyBridgePauser() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) PendingEmergencyBridgePauser() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.PendingEmergencyBridgePauser(&_Bridgel2sovereignchain.CallOpts)
}

// PendingEmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0x606617ff.
//
// Solidity: function pendingEmergencyBridgeUnpauser() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) PendingEmergencyBridgeUnpauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "pendingEmergencyBridgeUnpauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingEmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0x606617ff.
//
// Solidity: function pendingEmergencyBridgeUnpauser() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) PendingEmergencyBridgeUnpauser() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.PendingEmergencyBridgeUnpauser(&_Bridgel2sovereignchain.CallOpts)
}

// PendingEmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0x606617ff.
//
// Solidity: function pendingEmergencyBridgeUnpauser() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) PendingEmergencyBridgeUnpauser() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.PendingEmergencyBridgeUnpauser(&_Bridgel2sovereignchain.CallOpts)
}

// PendingProxiedTokensManager is a free data retrieval call binding the contract method 0xece93c6f.
//
// Solidity: function pendingProxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) PendingProxiedTokensManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "pendingProxiedTokensManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingProxiedTokensManager is a free data retrieval call binding the contract method 0xece93c6f.
//
// Solidity: function pendingProxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) PendingProxiedTokensManager() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.PendingProxiedTokensManager(&_Bridgel2sovereignchain.CallOpts)
}

// PendingProxiedTokensManager is a free data retrieval call binding the contract method 0xece93c6f.
//
// Solidity: function pendingProxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) PendingProxiedTokensManager() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.PendingProxiedTokensManager(&_Bridgel2sovereignchain.CallOpts)
}

// PolygonRollupManager is a free data retrieval call binding the contract method 0x8ed7e3f2.
//
// Solidity: function polygonRollupManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) PolygonRollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "polygonRollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PolygonRollupManager is a free data retrieval call binding the contract method 0x8ed7e3f2.
//
// Solidity: function polygonRollupManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) PolygonRollupManager() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.PolygonRollupManager(&_Bridgel2sovereignchain.CallOpts)
}

// PolygonRollupManager is a free data retrieval call binding the contract method 0x8ed7e3f2.
//
// Solidity: function polygonRollupManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) PolygonRollupManager() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.PolygonRollupManager(&_Bridgel2sovereignchain.CallOpts)
}

// ProxiedTokensManager is a free data retrieval call binding the contract method 0x6e4ecfed.
//
// Solidity: function proxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) ProxiedTokensManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "proxiedTokensManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxiedTokensManager is a free data retrieval call binding the contract method 0x6e4ecfed.
//
// Solidity: function proxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) ProxiedTokensManager() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.ProxiedTokensManager(&_Bridgel2sovereignchain.CallOpts)
}

// ProxiedTokensManager is a free data retrieval call binding the contract method 0x6e4ecfed.
//
// Solidity: function proxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) ProxiedTokensManager() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.ProxiedTokensManager(&_Bridgel2sovereignchain.CallOpts)
}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) TokenInfoToWrappedToken(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "tokenInfoToWrappedToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) TokenInfoToWrappedToken(arg0 [32]byte) (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.TokenInfoToWrappedToken(&_Bridgel2sovereignchain.CallOpts, arg0)
}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) TokenInfoToWrappedToken(arg0 [32]byte) (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.TokenInfoToWrappedToken(&_Bridgel2sovereignchain.CallOpts, arg0)
}

// UnsetGlobalIndexHashChain is a free data retrieval call binding the contract method 0x6ee84b23.
//
// Solidity: function unsetGlobalIndexHashChain() view returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) UnsetGlobalIndexHashChain(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "unsetGlobalIndexHashChain")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UnsetGlobalIndexHashChain is a free data retrieval call binding the contract method 0x6ee84b23.
//
// Solidity: function unsetGlobalIndexHashChain() view returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) UnsetGlobalIndexHashChain() ([32]byte, error) {
	return _Bridgel2sovereignchain.Contract.UnsetGlobalIndexHashChain(&_Bridgel2sovereignchain.CallOpts)
}

// UnsetGlobalIndexHashChain is a free data retrieval call binding the contract method 0x6ee84b23.
//
// Solidity: function unsetGlobalIndexHashChain() view returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) UnsetGlobalIndexHashChain() ([32]byte, error) {
	return _Bridgel2sovereignchain.Contract.UnsetGlobalIndexHashChain(&_Bridgel2sovereignchain.CallOpts)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) VerifyMerkleProof(opts *bind.CallOpts, leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "verifyMerkleProof", leafHash, smtProof, index, root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	return _Bridgel2sovereignchain.Contract.VerifyMerkleProof(&_Bridgel2sovereignchain.CallOpts, leafHash, smtProof, index, root)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	return _Bridgel2sovereignchain.Contract.VerifyMerkleProof(&_Bridgel2sovereignchain.CallOpts, leafHash, smtProof, index, root)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) Version() (string, error) {
	return _Bridgel2sovereignchain.Contract.Version(&_Bridgel2sovereignchain.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) Version() (string, error) {
	return _Bridgel2sovereignchain.Contract.Version(&_Bridgel2sovereignchain.CallOpts)
}

// WrappedAddressIsNotMintable is a free data retrieval call binding the contract method 0xc0f49163.
//
// Solidity: function wrappedAddressIsNotMintable(address wrappedAddress) view returns(bool isNotMintable)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) WrappedAddressIsNotMintable(opts *bind.CallOpts, wrappedAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "wrappedAddressIsNotMintable", wrappedAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WrappedAddressIsNotMintable is a free data retrieval call binding the contract method 0xc0f49163.
//
// Solidity: function wrappedAddressIsNotMintable(address wrappedAddress) view returns(bool isNotMintable)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) WrappedAddressIsNotMintable(wrappedAddress common.Address) (bool, error) {
	return _Bridgel2sovereignchain.Contract.WrappedAddressIsNotMintable(&_Bridgel2sovereignchain.CallOpts, wrappedAddress)
}

// WrappedAddressIsNotMintable is a free data retrieval call binding the contract method 0xc0f49163.
//
// Solidity: function wrappedAddressIsNotMintable(address wrappedAddress) view returns(bool isNotMintable)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) WrappedAddressIsNotMintable(wrappedAddress common.Address) (bool, error) {
	return _Bridgel2sovereignchain.Contract.WrappedAddressIsNotMintable(&_Bridgel2sovereignchain.CallOpts, wrappedAddress)
}

// WrappedTokenBytecodeStorer is a free data retrieval call binding the contract method 0x381fef6d.
//
// Solidity: function wrappedTokenBytecodeStorer() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) WrappedTokenBytecodeStorer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "wrappedTokenBytecodeStorer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedTokenBytecodeStorer is a free data retrieval call binding the contract method 0x381fef6d.
//
// Solidity: function wrappedTokenBytecodeStorer() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) WrappedTokenBytecodeStorer() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.WrappedTokenBytecodeStorer(&_Bridgel2sovereignchain.CallOpts)
}

// WrappedTokenBytecodeStorer is a free data retrieval call binding the contract method 0x381fef6d.
//
// Solidity: function wrappedTokenBytecodeStorer() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) WrappedTokenBytecodeStorer() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.WrappedTokenBytecodeStorer(&_Bridgel2sovereignchain.CallOpts)
}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address ) view returns(uint32 originNetwork, address originTokenAddress)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) WrappedTokenToTokenInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "wrappedTokenToTokenInfo", arg0)

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
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) WrappedTokenToTokenInfo(arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	return _Bridgel2sovereignchain.Contract.WrappedTokenToTokenInfo(&_Bridgel2sovereignchain.CallOpts, arg0)
}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address ) view returns(uint32 originNetwork, address originTokenAddress)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) WrappedTokenToTokenInfo(arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	return _Bridgel2sovereignchain.Contract.WrappedTokenToTokenInfo(&_Bridgel2sovereignchain.CallOpts, arg0)
}

// AcceptEmergencyBridgePauserRole is a paid mutator transaction binding the contract method 0xe88f0436.
//
// Solidity: function acceptEmergencyBridgePauserRole() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) AcceptEmergencyBridgePauserRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "acceptEmergencyBridgePauserRole")
}

// AcceptEmergencyBridgePauserRole is a paid mutator transaction binding the contract method 0xe88f0436.
//
// Solidity: function acceptEmergencyBridgePauserRole() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) AcceptEmergencyBridgePauserRole() (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.AcceptEmergencyBridgePauserRole(&_Bridgel2sovereignchain.TransactOpts)
}

// AcceptEmergencyBridgePauserRole is a paid mutator transaction binding the contract method 0xe88f0436.
//
// Solidity: function acceptEmergencyBridgePauserRole() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) AcceptEmergencyBridgePauserRole() (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.AcceptEmergencyBridgePauserRole(&_Bridgel2sovereignchain.TransactOpts)
}

// AcceptEmergencyBridgeUnpauserRole is a paid mutator transaction binding the contract method 0x8b37b873.
//
// Solidity: function acceptEmergencyBridgeUnpauserRole() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) AcceptEmergencyBridgeUnpauserRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "acceptEmergencyBridgeUnpauserRole")
}

// AcceptEmergencyBridgeUnpauserRole is a paid mutator transaction binding the contract method 0x8b37b873.
//
// Solidity: function acceptEmergencyBridgeUnpauserRole() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) AcceptEmergencyBridgeUnpauserRole() (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.AcceptEmergencyBridgeUnpauserRole(&_Bridgel2sovereignchain.TransactOpts)
}

// AcceptEmergencyBridgeUnpauserRole is a paid mutator transaction binding the contract method 0x8b37b873.
//
// Solidity: function acceptEmergencyBridgeUnpauserRole() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) AcceptEmergencyBridgeUnpauserRole() (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.AcceptEmergencyBridgeUnpauserRole(&_Bridgel2sovereignchain.TransactOpts)
}

// AcceptProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8c668f1c.
//
// Solidity: function acceptProxiedTokensManagerRole() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) AcceptProxiedTokensManagerRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "acceptProxiedTokensManagerRole")
}

// AcceptProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8c668f1c.
//
// Solidity: function acceptProxiedTokensManagerRole() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) AcceptProxiedTokensManagerRole() (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.AcceptProxiedTokensManagerRole(&_Bridgel2sovereignchain.TransactOpts)
}

// AcceptProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8c668f1c.
//
// Solidity: function acceptProxiedTokensManagerRole() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) AcceptProxiedTokensManagerRole() (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.AcceptProxiedTokensManagerRole(&_Bridgel2sovereignchain.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) ActivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "activateEmergencyState")
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.ActivateEmergencyState(&_Bridgel2sovereignchain.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.ActivateEmergencyState(&_Bridgel2sovereignchain.TransactOpts)
}

// BackwardLET is a paid mutator transaction binding the contract method 0xfd7640e8.
//
// Solidity: function backwardLET(uint256 newDepositCount, bytes32[32] newFrontier, bytes32 nextLeaf, bytes32[32] proof) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) BackwardLET(opts *bind.TransactOpts, newDepositCount *big.Int, newFrontier [32][32]byte, nextLeaf [32]byte, proof [32][32]byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "backwardLET", newDepositCount, newFrontier, nextLeaf, proof)
}

// BackwardLET is a paid mutator transaction binding the contract method 0xfd7640e8.
//
// Solidity: function backwardLET(uint256 newDepositCount, bytes32[32] newFrontier, bytes32 nextLeaf, bytes32[32] proof) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) BackwardLET(newDepositCount *big.Int, newFrontier [32][32]byte, nextLeaf [32]byte, proof [32][32]byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.BackwardLET(&_Bridgel2sovereignchain.TransactOpts, newDepositCount, newFrontier, nextLeaf, proof)
}

// BackwardLET is a paid mutator transaction binding the contract method 0xfd7640e8.
//
// Solidity: function backwardLET(uint256 newDepositCount, bytes32[32] newFrontier, bytes32 nextLeaf, bytes32[32] proof) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) BackwardLET(newDepositCount *big.Int, newFrontier [32][32]byte, nextLeaf [32]byte, proof [32][32]byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.BackwardLET(&_Bridgel2sovereignchain.TransactOpts, newDepositCount, newFrontier, nextLeaf, proof)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) BridgeAsset(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "bridgeAsset", destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) BridgeAsset(destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.BridgeAsset(&_Bridgel2sovereignchain.TransactOpts, destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) BridgeAsset(destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.BridgeAsset(&_Bridgel2sovereignchain.TransactOpts, destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) BridgeMessage(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "bridgeMessage", destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) BridgeMessage(destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.BridgeMessage(&_Bridgel2sovereignchain.TransactOpts, destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) BridgeMessage(destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.BridgeMessage(&_Bridgel2sovereignchain.TransactOpts, destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) BridgeMessageWETH(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "bridgeMessageWETH", destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) BridgeMessageWETH(destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.BridgeMessageWETH(&_Bridgel2sovereignchain.TransactOpts, destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) BridgeMessageWETH(destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.BridgeMessageWETH(&_Bridgel2sovereignchain.TransactOpts, destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) ClaimAsset(opts *bind.TransactOpts, smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "claimAsset", smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) ClaimAsset(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.ClaimAsset(&_Bridgel2sovereignchain.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) ClaimAsset(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.ClaimAsset(&_Bridgel2sovereignchain.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) ClaimMessage(opts *bind.TransactOpts, smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "claimMessage", smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) ClaimMessage(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.ClaimMessage(&_Bridgel2sovereignchain.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) ClaimMessage(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.ClaimMessage(&_Bridgel2sovereignchain.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.DeactivateEmergencyState(&_Bridgel2sovereignchain.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.DeactivateEmergencyState(&_Bridgel2sovereignchain.TransactOpts)
}

// DeployWrappedTokenAndRemap is a paid mutator transaction binding the contract method 0x6e974cd4.
//
// Solidity: function deployWrappedTokenAndRemap(uint32 originNetwork, address originTokenAddress, bool isNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) DeployWrappedTokenAndRemap(opts *bind.TransactOpts, originNetwork uint32, originTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "deployWrappedTokenAndRemap", originNetwork, originTokenAddress, isNotMintable)
}

// DeployWrappedTokenAndRemap is a paid mutator transaction binding the contract method 0x6e974cd4.
//
// Solidity: function deployWrappedTokenAndRemap(uint32 originNetwork, address originTokenAddress, bool isNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) DeployWrappedTokenAndRemap(originNetwork uint32, originTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.DeployWrappedTokenAndRemap(&_Bridgel2sovereignchain.TransactOpts, originNetwork, originTokenAddress, isNotMintable)
}

// DeployWrappedTokenAndRemap is a paid mutator transaction binding the contract method 0x6e974cd4.
//
// Solidity: function deployWrappedTokenAndRemap(uint32 originNetwork, address originTokenAddress, bool isNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) DeployWrappedTokenAndRemap(originNetwork uint32, originTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.DeployWrappedTokenAndRemap(&_Bridgel2sovereignchain.TransactOpts, originNetwork, originTokenAddress, isNotMintable)
}

// ForwardLET is a paid mutator transaction binding the contract method 0x70ecac1b.
//
// Solidity: function forwardLET((uint8,uint32,address,uint32,address,uint256,bytes32)[] newLeaves, bytes32 expectedStateRoot) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) ForwardLET(opts *bind.TransactOpts, newLeaves []BridgeL2SovereignChainLeafData, expectedStateRoot [32]byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "forwardLET", newLeaves, expectedStateRoot)
}

// ForwardLET is a paid mutator transaction binding the contract method 0x70ecac1b.
//
// Solidity: function forwardLET((uint8,uint32,address,uint32,address,uint256,bytes32)[] newLeaves, bytes32 expectedStateRoot) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) ForwardLET(newLeaves []BridgeL2SovereignChainLeafData, expectedStateRoot [32]byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.ForwardLET(&_Bridgel2sovereignchain.TransactOpts, newLeaves, expectedStateRoot)
}

// ForwardLET is a paid mutator transaction binding the contract method 0x70ecac1b.
//
// Solidity: function forwardLET((uint8,uint32,address,uint32,address,uint256,bytes32)[] newLeaves, bytes32 expectedStateRoot) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) ForwardLET(newLeaves []BridgeL2SovereignChainLeafData, expectedStateRoot [32]byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.ForwardLET(&_Bridgel2sovereignchain.TransactOpts, newLeaves, expectedStateRoot)
}

// Initialize is a paid mutator transaction binding the contract method 0x006ee171.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address _sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable, address _emergencyBridgePauser, address _emergencyBridgeUnpauser, address _proxiedTokensManager) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) Initialize(opts *bind.TransactOpts, _networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, _sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool, _emergencyBridgePauser common.Address, _emergencyBridgeUnpauser common.Address, _proxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "initialize", _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, _sovereignWETHAddress, _sovereignWETHAddressIsNotMintable, _emergencyBridgePauser, _emergencyBridgeUnpauser, _proxiedTokensManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x006ee171.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address _sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable, address _emergencyBridgePauser, address _emergencyBridgeUnpauser, address _proxiedTokensManager) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) Initialize(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, _sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool, _emergencyBridgePauser common.Address, _emergencyBridgeUnpauser common.Address, _proxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.Initialize(&_Bridgel2sovereignchain.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, _sovereignWETHAddress, _sovereignWETHAddressIsNotMintable, _emergencyBridgePauser, _emergencyBridgeUnpauser, _proxiedTokensManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x006ee171.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address _sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable, address _emergencyBridgePauser, address _emergencyBridgeUnpauser, address _proxiedTokensManager) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) Initialize(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, _sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool, _emergencyBridgePauser common.Address, _emergencyBridgeUnpauser common.Address, _proxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.Initialize(&_Bridgel2sovereignchain.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, _sovereignWETHAddress, _sovereignWETHAddressIsNotMintable, _emergencyBridgePauser, _emergencyBridgeUnpauser, _proxiedTokensManager)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes ) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) Initialize1(opts *bind.TransactOpts, arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "initialize1", arg0, arg1, arg2, arg3, arg4, arg5)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes ) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) Initialize1(arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.Initialize1(&_Bridgel2sovereignchain.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes ) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) Initialize1(arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.Initialize1(&_Bridgel2sovereignchain.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// MigrateLegacyToken is a paid mutator transaction binding the contract method 0xb0b37920.
//
// Solidity: function migrateLegacyToken(address legacyTokenAddress, uint256 amount, bytes permitData) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) MigrateLegacyToken(opts *bind.TransactOpts, legacyTokenAddress common.Address, amount *big.Int, permitData []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "migrateLegacyToken", legacyTokenAddress, amount, permitData)
}

// MigrateLegacyToken is a paid mutator transaction binding the contract method 0xb0b37920.
//
// Solidity: function migrateLegacyToken(address legacyTokenAddress, uint256 amount, bytes permitData) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) MigrateLegacyToken(legacyTokenAddress common.Address, amount *big.Int, permitData []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.MigrateLegacyToken(&_Bridgel2sovereignchain.TransactOpts, legacyTokenAddress, amount, permitData)
}

// MigrateLegacyToken is a paid mutator transaction binding the contract method 0xb0b37920.
//
// Solidity: function migrateLegacyToken(address legacyTokenAddress, uint256 amount, bytes permitData) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) MigrateLegacyToken(legacyTokenAddress common.Address, amount *big.Int, permitData []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.MigrateLegacyToken(&_Bridgel2sovereignchain.TransactOpts, legacyTokenAddress, amount, permitData)
}

// RemoveLegacySovereignTokenAddress is a paid mutator transaction binding the contract method 0xb4586962.
//
// Solidity: function removeLegacySovereignTokenAddress(address legacySovereignTokenAddress) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) RemoveLegacySovereignTokenAddress(opts *bind.TransactOpts, legacySovereignTokenAddress common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "removeLegacySovereignTokenAddress", legacySovereignTokenAddress)
}

// RemoveLegacySovereignTokenAddress is a paid mutator transaction binding the contract method 0xb4586962.
//
// Solidity: function removeLegacySovereignTokenAddress(address legacySovereignTokenAddress) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) RemoveLegacySovereignTokenAddress(legacySovereignTokenAddress common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.RemoveLegacySovereignTokenAddress(&_Bridgel2sovereignchain.TransactOpts, legacySovereignTokenAddress)
}

// RemoveLegacySovereignTokenAddress is a paid mutator transaction binding the contract method 0xb4586962.
//
// Solidity: function removeLegacySovereignTokenAddress(address legacySovereignTokenAddress) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) RemoveLegacySovereignTokenAddress(legacySovereignTokenAddress common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.RemoveLegacySovereignTokenAddress(&_Bridgel2sovereignchain.TransactOpts, legacySovereignTokenAddress)
}

// SetBridgeManager is a paid mutator transaction binding the contract method 0xeabd372a.
//
// Solidity: function setBridgeManager(address _bridgeManager) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) SetBridgeManager(opts *bind.TransactOpts, _bridgeManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "setBridgeManager", _bridgeManager)
}

// SetBridgeManager is a paid mutator transaction binding the contract method 0xeabd372a.
//
// Solidity: function setBridgeManager(address _bridgeManager) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) SetBridgeManager(_bridgeManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.SetBridgeManager(&_Bridgel2sovereignchain.TransactOpts, _bridgeManager)
}

// SetBridgeManager is a paid mutator transaction binding the contract method 0xeabd372a.
//
// Solidity: function setBridgeManager(address _bridgeManager) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) SetBridgeManager(_bridgeManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.SetBridgeManager(&_Bridgel2sovereignchain.TransactOpts, _bridgeManager)
}

// SetLocalBalanceTree is a paid mutator transaction binding the contract method 0x8f9720bb.
//
// Solidity: function setLocalBalanceTree(uint32[] originNetwork, address[] originTokenAddress, uint256[] amount) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) SetLocalBalanceTree(opts *bind.TransactOpts, originNetwork []uint32, originTokenAddress []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "setLocalBalanceTree", originNetwork, originTokenAddress, amount)
}

// SetLocalBalanceTree is a paid mutator transaction binding the contract method 0x8f9720bb.
//
// Solidity: function setLocalBalanceTree(uint32[] originNetwork, address[] originTokenAddress, uint256[] amount) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) SetLocalBalanceTree(originNetwork []uint32, originTokenAddress []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.SetLocalBalanceTree(&_Bridgel2sovereignchain.TransactOpts, originNetwork, originTokenAddress, amount)
}

// SetLocalBalanceTree is a paid mutator transaction binding the contract method 0x8f9720bb.
//
// Solidity: function setLocalBalanceTree(uint32[] originNetwork, address[] originTokenAddress, uint256[] amount) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) SetLocalBalanceTree(originNetwork []uint32, originTokenAddress []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.SetLocalBalanceTree(&_Bridgel2sovereignchain.TransactOpts, originNetwork, originTokenAddress, amount)
}

// SetMultipleClaims is a paid mutator transaction binding the contract method 0x1c208229.
//
// Solidity: function setMultipleClaims(uint256[] globalIndexes) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) SetMultipleClaims(opts *bind.TransactOpts, globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "setMultipleClaims", globalIndexes)
}

// SetMultipleClaims is a paid mutator transaction binding the contract method 0x1c208229.
//
// Solidity: function setMultipleClaims(uint256[] globalIndexes) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) SetMultipleClaims(globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.SetMultipleClaims(&_Bridgel2sovereignchain.TransactOpts, globalIndexes)
}

// SetMultipleClaims is a paid mutator transaction binding the contract method 0x1c208229.
//
// Solidity: function setMultipleClaims(uint256[] globalIndexes) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) SetMultipleClaims(globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.SetMultipleClaims(&_Bridgel2sovereignchain.TransactOpts, globalIndexes)
}

// SetMultipleSovereignTokenAddress is a paid mutator transaction binding the contract method 0x57cfbee3.
//
// Solidity: function setMultipleSovereignTokenAddress(uint32[] originNetworks, address[] originTokenAddresses, address[] sovereignTokenAddresses, bool[] isNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) SetMultipleSovereignTokenAddress(opts *bind.TransactOpts, originNetworks []uint32, originTokenAddresses []common.Address, sovereignTokenAddresses []common.Address, isNotMintable []bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "setMultipleSovereignTokenAddress", originNetworks, originTokenAddresses, sovereignTokenAddresses, isNotMintable)
}

// SetMultipleSovereignTokenAddress is a paid mutator transaction binding the contract method 0x57cfbee3.
//
// Solidity: function setMultipleSovereignTokenAddress(uint32[] originNetworks, address[] originTokenAddresses, address[] sovereignTokenAddresses, bool[] isNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) SetMultipleSovereignTokenAddress(originNetworks []uint32, originTokenAddresses []common.Address, sovereignTokenAddresses []common.Address, isNotMintable []bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.SetMultipleSovereignTokenAddress(&_Bridgel2sovereignchain.TransactOpts, originNetworks, originTokenAddresses, sovereignTokenAddresses, isNotMintable)
}

// SetMultipleSovereignTokenAddress is a paid mutator transaction binding the contract method 0x57cfbee3.
//
// Solidity: function setMultipleSovereignTokenAddress(uint32[] originNetworks, address[] originTokenAddresses, address[] sovereignTokenAddresses, bool[] isNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) SetMultipleSovereignTokenAddress(originNetworks []uint32, originTokenAddresses []common.Address, sovereignTokenAddresses []common.Address, isNotMintable []bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.SetMultipleSovereignTokenAddress(&_Bridgel2sovereignchain.TransactOpts, originNetworks, originTokenAddresses, sovereignTokenAddresses, isNotMintable)
}

// SetSovereignWETHAddress is a paid mutator transaction binding the contract method 0xbf130d7f.
//
// Solidity: function setSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) SetSovereignWETHAddress(opts *bind.TransactOpts, sovereignWETHTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "setSovereignWETHAddress", sovereignWETHTokenAddress, isNotMintable)
}

// SetSovereignWETHAddress is a paid mutator transaction binding the contract method 0xbf130d7f.
//
// Solidity: function setSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) SetSovereignWETHAddress(sovereignWETHTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.SetSovereignWETHAddress(&_Bridgel2sovereignchain.TransactOpts, sovereignWETHTokenAddress, isNotMintable)
}

// SetSovereignWETHAddress is a paid mutator transaction binding the contract method 0xbf130d7f.
//
// Solidity: function setSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) SetSovereignWETHAddress(sovereignWETHTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.SetSovereignWETHAddress(&_Bridgel2sovereignchain.TransactOpts, sovereignWETHTokenAddress, isNotMintable)
}

// TransferEmergencyBridgePauserRole is a paid mutator transaction binding the contract method 0x69e3ab12.
//
// Solidity: function transferEmergencyBridgePauserRole(address newEmergencyBridgePauser) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) TransferEmergencyBridgePauserRole(opts *bind.TransactOpts, newEmergencyBridgePauser common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "transferEmergencyBridgePauserRole", newEmergencyBridgePauser)
}

// TransferEmergencyBridgePauserRole is a paid mutator transaction binding the contract method 0x69e3ab12.
//
// Solidity: function transferEmergencyBridgePauserRole(address newEmergencyBridgePauser) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) TransferEmergencyBridgePauserRole(newEmergencyBridgePauser common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.TransferEmergencyBridgePauserRole(&_Bridgel2sovereignchain.TransactOpts, newEmergencyBridgePauser)
}

// TransferEmergencyBridgePauserRole is a paid mutator transaction binding the contract method 0x69e3ab12.
//
// Solidity: function transferEmergencyBridgePauserRole(address newEmergencyBridgePauser) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) TransferEmergencyBridgePauserRole(newEmergencyBridgePauser common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.TransferEmergencyBridgePauserRole(&_Bridgel2sovereignchain.TransactOpts, newEmergencyBridgePauser)
}

// TransferEmergencyBridgeUnpauserRole is a paid mutator transaction binding the contract method 0x8d942096.
//
// Solidity: function transferEmergencyBridgeUnpauserRole(address newEmergencyBridgeUnpauser) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) TransferEmergencyBridgeUnpauserRole(opts *bind.TransactOpts, newEmergencyBridgeUnpauser common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "transferEmergencyBridgeUnpauserRole", newEmergencyBridgeUnpauser)
}

// TransferEmergencyBridgeUnpauserRole is a paid mutator transaction binding the contract method 0x8d942096.
//
// Solidity: function transferEmergencyBridgeUnpauserRole(address newEmergencyBridgeUnpauser) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) TransferEmergencyBridgeUnpauserRole(newEmergencyBridgeUnpauser common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.TransferEmergencyBridgeUnpauserRole(&_Bridgel2sovereignchain.TransactOpts, newEmergencyBridgeUnpauser)
}

// TransferEmergencyBridgeUnpauserRole is a paid mutator transaction binding the contract method 0x8d942096.
//
// Solidity: function transferEmergencyBridgeUnpauserRole(address newEmergencyBridgeUnpauser) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) TransferEmergencyBridgeUnpauserRole(newEmergencyBridgeUnpauser common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.TransferEmergencyBridgeUnpauserRole(&_Bridgel2sovereignchain.TransactOpts, newEmergencyBridgeUnpauser)
}

// TransferProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8bd309c3.
//
// Solidity: function transferProxiedTokensManagerRole(address newProxiedTokensManager) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) TransferProxiedTokensManagerRole(opts *bind.TransactOpts, newProxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "transferProxiedTokensManagerRole", newProxiedTokensManager)
}

// TransferProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8bd309c3.
//
// Solidity: function transferProxiedTokensManagerRole(address newProxiedTokensManager) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) TransferProxiedTokensManagerRole(newProxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.TransferProxiedTokensManagerRole(&_Bridgel2sovereignchain.TransactOpts, newProxiedTokensManager)
}

// TransferProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8bd309c3.
//
// Solidity: function transferProxiedTokensManagerRole(address newProxiedTokensManager) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) TransferProxiedTokensManagerRole(newProxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.TransferProxiedTokensManagerRole(&_Bridgel2sovereignchain.TransactOpts, newProxiedTokensManager)
}

// UnsetMultipleClaims is a paid mutator transaction binding the contract method 0x136a2c60.
//
// Solidity: function unsetMultipleClaims(uint256[] globalIndexes) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) UnsetMultipleClaims(opts *bind.TransactOpts, globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "unsetMultipleClaims", globalIndexes)
}

// UnsetMultipleClaims is a paid mutator transaction binding the contract method 0x136a2c60.
//
// Solidity: function unsetMultipleClaims(uint256[] globalIndexes) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) UnsetMultipleClaims(globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.UnsetMultipleClaims(&_Bridgel2sovereignchain.TransactOpts, globalIndexes)
}

// UnsetMultipleClaims is a paid mutator transaction binding the contract method 0x136a2c60.
//
// Solidity: function unsetMultipleClaims(uint256[] globalIndexes) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) UnsetMultipleClaims(globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.UnsetMultipleClaims(&_Bridgel2sovereignchain.TransactOpts, globalIndexes)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) UpdateGlobalExitRoot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "updateGlobalExitRoot")
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) UpdateGlobalExitRoot() (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.UpdateGlobalExitRoot(&_Bridgel2sovereignchain.TransactOpts)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) UpdateGlobalExitRoot() (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.UpdateGlobalExitRoot(&_Bridgel2sovereignchain.TransactOpts)
}

// Bridgel2sovereignchainAcceptEmergencyBridgePauserRoleIterator is returned from FilterAcceptEmergencyBridgePauserRole and is used to iterate over the raw logs and unpacked data for AcceptEmergencyBridgePauserRole events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainAcceptEmergencyBridgePauserRoleIterator struct {
	Event *Bridgel2sovereignchainAcceptEmergencyBridgePauserRole // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainAcceptEmergencyBridgePauserRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainAcceptEmergencyBridgePauserRole)
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
		it.Event = new(Bridgel2sovereignchainAcceptEmergencyBridgePauserRole)
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
func (it *Bridgel2sovereignchainAcceptEmergencyBridgePauserRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainAcceptEmergencyBridgePauserRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainAcceptEmergencyBridgePauserRole represents a AcceptEmergencyBridgePauserRole event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainAcceptEmergencyBridgePauserRole struct {
	OldEmergencyBridgePauser common.Address
	NewEmergencyBridgePauser common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterAcceptEmergencyBridgePauserRole is a free log retrieval operation binding the contract event 0x85d2bdfbe58cd81abf8199c13ce2509204be4aba8603b9d29f52c4e13e7bb793.
//
// Solidity: event AcceptEmergencyBridgePauserRole(address oldEmergencyBridgePauser, address newEmergencyBridgePauser)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterAcceptEmergencyBridgePauserRole(opts *bind.FilterOpts) (*Bridgel2sovereignchainAcceptEmergencyBridgePauserRoleIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "AcceptEmergencyBridgePauserRole")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainAcceptEmergencyBridgePauserRoleIterator{contract: _Bridgel2sovereignchain.contract, event: "AcceptEmergencyBridgePauserRole", logs: logs, sub: sub}, nil
}

// WatchAcceptEmergencyBridgePauserRole is a free log subscription operation binding the contract event 0x85d2bdfbe58cd81abf8199c13ce2509204be4aba8603b9d29f52c4e13e7bb793.
//
// Solidity: event AcceptEmergencyBridgePauserRole(address oldEmergencyBridgePauser, address newEmergencyBridgePauser)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchAcceptEmergencyBridgePauserRole(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainAcceptEmergencyBridgePauserRole) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "AcceptEmergencyBridgePauserRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainAcceptEmergencyBridgePauserRole)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "AcceptEmergencyBridgePauserRole", log); err != nil {
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
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseAcceptEmergencyBridgePauserRole(log types.Log) (*Bridgel2sovereignchainAcceptEmergencyBridgePauserRole, error) {
	event := new(Bridgel2sovereignchainAcceptEmergencyBridgePauserRole)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "AcceptEmergencyBridgePauserRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRoleIterator is returned from FilterAcceptEmergencyBridgeUnpauserRole and is used to iterate over the raw logs and unpacked data for AcceptEmergencyBridgeUnpauserRole events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRoleIterator struct {
	Event *Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRole // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRole)
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
		it.Event = new(Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRole)
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
func (it *Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRole represents a AcceptEmergencyBridgeUnpauserRole event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRole struct {
	OldEmergencyBridgeUnpauser common.Address
	NewEmergencyBridgeUnpauser common.Address
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterAcceptEmergencyBridgeUnpauserRole is a free log retrieval operation binding the contract event 0x24cc8295aa5110cc216695db944ad2458c7795c6404449be980c3ce14aed752d.
//
// Solidity: event AcceptEmergencyBridgeUnpauserRole(address oldEmergencyBridgeUnpauser, address newEmergencyBridgeUnpauser)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterAcceptEmergencyBridgeUnpauserRole(opts *bind.FilterOpts) (*Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRoleIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "AcceptEmergencyBridgeUnpauserRole")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRoleIterator{contract: _Bridgel2sovereignchain.contract, event: "AcceptEmergencyBridgeUnpauserRole", logs: logs, sub: sub}, nil
}

// WatchAcceptEmergencyBridgeUnpauserRole is a free log subscription operation binding the contract event 0x24cc8295aa5110cc216695db944ad2458c7795c6404449be980c3ce14aed752d.
//
// Solidity: event AcceptEmergencyBridgeUnpauserRole(address oldEmergencyBridgeUnpauser, address newEmergencyBridgeUnpauser)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchAcceptEmergencyBridgeUnpauserRole(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRole) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "AcceptEmergencyBridgeUnpauserRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRole)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "AcceptEmergencyBridgeUnpauserRole", log); err != nil {
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
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseAcceptEmergencyBridgeUnpauserRole(log types.Log) (*Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRole, error) {
	event := new(Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRole)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "AcceptEmergencyBridgeUnpauserRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainAcceptProxiedTokensManagerRoleIterator is returned from FilterAcceptProxiedTokensManagerRole and is used to iterate over the raw logs and unpacked data for AcceptProxiedTokensManagerRole events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainAcceptProxiedTokensManagerRoleIterator struct {
	Event *Bridgel2sovereignchainAcceptProxiedTokensManagerRole // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainAcceptProxiedTokensManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainAcceptProxiedTokensManagerRole)
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
		it.Event = new(Bridgel2sovereignchainAcceptProxiedTokensManagerRole)
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
func (it *Bridgel2sovereignchainAcceptProxiedTokensManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainAcceptProxiedTokensManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainAcceptProxiedTokensManagerRole represents a AcceptProxiedTokensManagerRole event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainAcceptProxiedTokensManagerRole struct {
	OldProxiedTokensManager common.Address
	NewProxiedTokensManager common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterAcceptProxiedTokensManagerRole is a free log retrieval operation binding the contract event 0xa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f2.
//
// Solidity: event AcceptProxiedTokensManagerRole(address oldProxiedTokensManager, address newProxiedTokensManager)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterAcceptProxiedTokensManagerRole(opts *bind.FilterOpts) (*Bridgel2sovereignchainAcceptProxiedTokensManagerRoleIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "AcceptProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainAcceptProxiedTokensManagerRoleIterator{contract: _Bridgel2sovereignchain.contract, event: "AcceptProxiedTokensManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptProxiedTokensManagerRole is a free log subscription operation binding the contract event 0xa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f2.
//
// Solidity: event AcceptProxiedTokensManagerRole(address oldProxiedTokensManager, address newProxiedTokensManager)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchAcceptProxiedTokensManagerRole(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainAcceptProxiedTokensManagerRole) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "AcceptProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainAcceptProxiedTokensManagerRole)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "AcceptProxiedTokensManagerRole", log); err != nil {
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
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseAcceptProxiedTokensManagerRole(log types.Log) (*Bridgel2sovereignchainAcceptProxiedTokensManagerRole, error) {
	event := new(Bridgel2sovereignchainAcceptProxiedTokensManagerRole)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "AcceptProxiedTokensManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainBackwardLETIterator is returned from FilterBackwardLET and is used to iterate over the raw logs and unpacked data for BackwardLET events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainBackwardLETIterator struct {
	Event *Bridgel2sovereignchainBackwardLET // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainBackwardLETIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainBackwardLET)
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
		it.Event = new(Bridgel2sovereignchainBackwardLET)
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
func (it *Bridgel2sovereignchainBackwardLETIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainBackwardLETIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainBackwardLET represents a BackwardLET event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainBackwardLET struct {
	NewDepositCount *big.Int
	NewRoot         [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBackwardLET is a free log retrieval operation binding the contract event 0x90799a4e436a5b97adca76982886b3c894dc74b50169ab082b9a069f99317e61.
//
// Solidity: event BackwardLET(uint256 newDepositCount, bytes32 newRoot)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterBackwardLET(opts *bind.FilterOpts) (*Bridgel2sovereignchainBackwardLETIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "BackwardLET")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainBackwardLETIterator{contract: _Bridgel2sovereignchain.contract, event: "BackwardLET", logs: logs, sub: sub}, nil
}

// WatchBackwardLET is a free log subscription operation binding the contract event 0x90799a4e436a5b97adca76982886b3c894dc74b50169ab082b9a069f99317e61.
//
// Solidity: event BackwardLET(uint256 newDepositCount, bytes32 newRoot)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchBackwardLET(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainBackwardLET) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "BackwardLET")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainBackwardLET)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "BackwardLET", log); err != nil {
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

// ParseBackwardLET is a log parse operation binding the contract event 0x90799a4e436a5b97adca76982886b3c894dc74b50169ab082b9a069f99317e61.
//
// Solidity: event BackwardLET(uint256 newDepositCount, bytes32 newRoot)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseBackwardLET(log types.Log) (*Bridgel2sovereignchainBackwardLET, error) {
	event := new(Bridgel2sovereignchainBackwardLET)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "BackwardLET", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainBridgeEventIterator is returned from FilterBridgeEvent and is used to iterate over the raw logs and unpacked data for BridgeEvent events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainBridgeEventIterator struct {
	Event *Bridgel2sovereignchainBridgeEvent // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainBridgeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainBridgeEvent)
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
		it.Event = new(Bridgel2sovereignchainBridgeEvent)
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
func (it *Bridgel2sovereignchainBridgeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainBridgeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainBridgeEvent represents a BridgeEvent event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainBridgeEvent struct {
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
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterBridgeEvent(opts *bind.FilterOpts) (*Bridgel2sovereignchainBridgeEventIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "BridgeEvent")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainBridgeEventIterator{contract: _Bridgel2sovereignchain.contract, event: "BridgeEvent", logs: logs, sub: sub}, nil
}

// WatchBridgeEvent is a free log subscription operation binding the contract event 0x501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b.
//
// Solidity: event BridgeEvent(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata, uint32 depositCount)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchBridgeEvent(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainBridgeEvent) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "BridgeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainBridgeEvent)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "BridgeEvent", log); err != nil {
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
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseBridgeEvent(log types.Log) (*Bridgel2sovereignchainBridgeEvent, error) {
	event := new(Bridgel2sovereignchainBridgeEvent)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "BridgeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainClaimEventIterator is returned from FilterClaimEvent and is used to iterate over the raw logs and unpacked data for ClaimEvent events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainClaimEventIterator struct {
	Event *Bridgel2sovereignchainClaimEvent // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainClaimEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainClaimEvent)
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
		it.Event = new(Bridgel2sovereignchainClaimEvent)
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
func (it *Bridgel2sovereignchainClaimEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainClaimEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainClaimEvent represents a ClaimEvent event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainClaimEvent struct {
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
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterClaimEvent(opts *bind.FilterOpts) (*Bridgel2sovereignchainClaimEventIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "ClaimEvent")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainClaimEventIterator{contract: _Bridgel2sovereignchain.contract, event: "ClaimEvent", logs: logs, sub: sub}, nil
}

// WatchClaimEvent is a free log subscription operation binding the contract event 0x1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d.
//
// Solidity: event ClaimEvent(uint256 globalIndex, uint32 originNetwork, address originAddress, address destinationAddress, uint256 amount)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchClaimEvent(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainClaimEvent) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "ClaimEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainClaimEvent)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "ClaimEvent", log); err != nil {
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
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseClaimEvent(log types.Log) (*Bridgel2sovereignchainClaimEvent, error) {
	event := new(Bridgel2sovereignchainClaimEvent)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "ClaimEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainEmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainEmergencyStateActivatedIterator struct {
	Event *Bridgel2sovereignchainEmergencyStateActivated // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainEmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainEmergencyStateActivated)
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
		it.Event = new(Bridgel2sovereignchainEmergencyStateActivated)
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
func (it *Bridgel2sovereignchainEmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainEmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainEmergencyStateActivated represents a EmergencyStateActivated event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainEmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*Bridgel2sovereignchainEmergencyStateActivatedIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainEmergencyStateActivatedIterator{contract: _Bridgel2sovereignchain.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainEmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainEmergencyStateActivated)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
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
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseEmergencyStateActivated(log types.Log) (*Bridgel2sovereignchainEmergencyStateActivated, error) {
	event := new(Bridgel2sovereignchainEmergencyStateActivated)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainEmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainEmergencyStateDeactivatedIterator struct {
	Event *Bridgel2sovereignchainEmergencyStateDeactivated // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainEmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainEmergencyStateDeactivated)
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
		it.Event = new(Bridgel2sovereignchainEmergencyStateDeactivated)
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
func (it *Bridgel2sovereignchainEmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainEmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainEmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainEmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*Bridgel2sovereignchainEmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainEmergencyStateDeactivatedIterator{contract: _Bridgel2sovereignchain.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainEmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainEmergencyStateDeactivated)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
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
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseEmergencyStateDeactivated(log types.Log) (*Bridgel2sovereignchainEmergencyStateDeactivated, error) {
	event := new(Bridgel2sovereignchainEmergencyStateDeactivated)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainForwardLETIterator is returned from FilterForwardLET and is used to iterate over the raw logs and unpacked data for ForwardLET events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainForwardLETIterator struct {
	Event *Bridgel2sovereignchainForwardLET // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainForwardLETIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainForwardLET)
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
		it.Event = new(Bridgel2sovereignchainForwardLET)
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
func (it *Bridgel2sovereignchainForwardLETIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainForwardLETIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainForwardLET represents a ForwardLET event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainForwardLET struct {
	NewDepositCount *big.Int
	NewRoot         [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterForwardLET is a free log retrieval operation binding the contract event 0xf833886dec6b10040de3e31a518be28ced3c9689a75d17741c23bcfb41975670.
//
// Solidity: event ForwardLET(uint256 newDepositCount, bytes32 newRoot)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterForwardLET(opts *bind.FilterOpts) (*Bridgel2sovereignchainForwardLETIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "ForwardLET")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainForwardLETIterator{contract: _Bridgel2sovereignchain.contract, event: "ForwardLET", logs: logs, sub: sub}, nil
}

// WatchForwardLET is a free log subscription operation binding the contract event 0xf833886dec6b10040de3e31a518be28ced3c9689a75d17741c23bcfb41975670.
//
// Solidity: event ForwardLET(uint256 newDepositCount, bytes32 newRoot)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchForwardLET(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainForwardLET) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "ForwardLET")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainForwardLET)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "ForwardLET", log); err != nil {
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

// ParseForwardLET is a log parse operation binding the contract event 0xf833886dec6b10040de3e31a518be28ced3c9689a75d17741c23bcfb41975670.
//
// Solidity: event ForwardLET(uint256 newDepositCount, bytes32 newRoot)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseForwardLET(log types.Log) (*Bridgel2sovereignchainForwardLET, error) {
	event := new(Bridgel2sovereignchainForwardLET)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "ForwardLET", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainInitializedIterator struct {
	Event *Bridgel2sovereignchainInitialized // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainInitialized)
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
		it.Event = new(Bridgel2sovereignchainInitialized)
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
func (it *Bridgel2sovereignchainInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainInitialized represents a Initialized event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterInitialized(opts *bind.FilterOpts) (*Bridgel2sovereignchainInitializedIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainInitializedIterator{contract: _Bridgel2sovereignchain.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainInitialized) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainInitialized)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseInitialized(log types.Log) (*Bridgel2sovereignchainInitialized, error) {
	event := new(Bridgel2sovereignchainInitialized)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainMigrateLegacyTokenIterator is returned from FilterMigrateLegacyToken and is used to iterate over the raw logs and unpacked data for MigrateLegacyToken events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainMigrateLegacyTokenIterator struct {
	Event *Bridgel2sovereignchainMigrateLegacyToken // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainMigrateLegacyTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainMigrateLegacyToken)
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
		it.Event = new(Bridgel2sovereignchainMigrateLegacyToken)
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
func (it *Bridgel2sovereignchainMigrateLegacyTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainMigrateLegacyTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainMigrateLegacyToken represents a MigrateLegacyToken event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainMigrateLegacyToken struct {
	Sender              common.Address
	LegacyTokenAddress  common.Address
	UpdatedTokenAddress common.Address
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMigrateLegacyToken is a free log retrieval operation binding the contract event 0xb7f8fd4d1faf9b2929dc269f59c53e3a2bccc44e9950f33a568fcbcb37eb69a9.
//
// Solidity: event MigrateLegacyToken(address sender, address legacyTokenAddress, address updatedTokenAddress, uint256 amount)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterMigrateLegacyToken(opts *bind.FilterOpts) (*Bridgel2sovereignchainMigrateLegacyTokenIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "MigrateLegacyToken")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainMigrateLegacyTokenIterator{contract: _Bridgel2sovereignchain.contract, event: "MigrateLegacyToken", logs: logs, sub: sub}, nil
}

// WatchMigrateLegacyToken is a free log subscription operation binding the contract event 0xb7f8fd4d1faf9b2929dc269f59c53e3a2bccc44e9950f33a568fcbcb37eb69a9.
//
// Solidity: event MigrateLegacyToken(address sender, address legacyTokenAddress, address updatedTokenAddress, uint256 amount)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchMigrateLegacyToken(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainMigrateLegacyToken) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "MigrateLegacyToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainMigrateLegacyToken)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "MigrateLegacyToken", log); err != nil {
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
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseMigrateLegacyToken(log types.Log) (*Bridgel2sovereignchainMigrateLegacyToken, error) {
	event := new(Bridgel2sovereignchainMigrateLegacyToken)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "MigrateLegacyToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainNewWrappedTokenIterator is returned from FilterNewWrappedToken and is used to iterate over the raw logs and unpacked data for NewWrappedToken events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainNewWrappedTokenIterator struct {
	Event *Bridgel2sovereignchainNewWrappedToken // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainNewWrappedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainNewWrappedToken)
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
		it.Event = new(Bridgel2sovereignchainNewWrappedToken)
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
func (it *Bridgel2sovereignchainNewWrappedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainNewWrappedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainNewWrappedToken represents a NewWrappedToken event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainNewWrappedToken struct {
	OriginNetwork       uint32
	OriginTokenAddress  common.Address
	WrappedTokenAddress common.Address
	Metadata            []byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewWrappedToken is a free log retrieval operation binding the contract event 0x490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a39.
//
// Solidity: event NewWrappedToken(uint32 originNetwork, address originTokenAddress, address wrappedTokenAddress, bytes metadata)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterNewWrappedToken(opts *bind.FilterOpts) (*Bridgel2sovereignchainNewWrappedTokenIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "NewWrappedToken")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainNewWrappedTokenIterator{contract: _Bridgel2sovereignchain.contract, event: "NewWrappedToken", logs: logs, sub: sub}, nil
}

// WatchNewWrappedToken is a free log subscription operation binding the contract event 0x490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a39.
//
// Solidity: event NewWrappedToken(uint32 originNetwork, address originTokenAddress, address wrappedTokenAddress, bytes metadata)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchNewWrappedToken(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainNewWrappedToken) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "NewWrappedToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainNewWrappedToken)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "NewWrappedToken", log); err != nil {
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
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseNewWrappedToken(log types.Log) (*Bridgel2sovereignchainNewWrappedToken, error) {
	event := new(Bridgel2sovereignchainNewWrappedToken)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "NewWrappedToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainRemoveLegacySovereignTokenAddressIterator is returned from FilterRemoveLegacySovereignTokenAddress and is used to iterate over the raw logs and unpacked data for RemoveLegacySovereignTokenAddress events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainRemoveLegacySovereignTokenAddressIterator struct {
	Event *Bridgel2sovereignchainRemoveLegacySovereignTokenAddress // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainRemoveLegacySovereignTokenAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainRemoveLegacySovereignTokenAddress)
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
		it.Event = new(Bridgel2sovereignchainRemoveLegacySovereignTokenAddress)
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
func (it *Bridgel2sovereignchainRemoveLegacySovereignTokenAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainRemoveLegacySovereignTokenAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainRemoveLegacySovereignTokenAddress represents a RemoveLegacySovereignTokenAddress event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainRemoveLegacySovereignTokenAddress struct {
	SovereignTokenAddress common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRemoveLegacySovereignTokenAddress is a free log retrieval operation binding the contract event 0xc2ae0bd0ec0fd0352bfe5bacac49637af342c1e40f1b80a7f74440dc7fe3f063.
//
// Solidity: event RemoveLegacySovereignTokenAddress(address sovereignTokenAddress)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterRemoveLegacySovereignTokenAddress(opts *bind.FilterOpts) (*Bridgel2sovereignchainRemoveLegacySovereignTokenAddressIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "RemoveLegacySovereignTokenAddress")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainRemoveLegacySovereignTokenAddressIterator{contract: _Bridgel2sovereignchain.contract, event: "RemoveLegacySovereignTokenAddress", logs: logs, sub: sub}, nil
}

// WatchRemoveLegacySovereignTokenAddress is a free log subscription operation binding the contract event 0xc2ae0bd0ec0fd0352bfe5bacac49637af342c1e40f1b80a7f74440dc7fe3f063.
//
// Solidity: event RemoveLegacySovereignTokenAddress(address sovereignTokenAddress)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchRemoveLegacySovereignTokenAddress(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainRemoveLegacySovereignTokenAddress) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "RemoveLegacySovereignTokenAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainRemoveLegacySovereignTokenAddress)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "RemoveLegacySovereignTokenAddress", log); err != nil {
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
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseRemoveLegacySovereignTokenAddress(log types.Log) (*Bridgel2sovereignchainRemoveLegacySovereignTokenAddress, error) {
	event := new(Bridgel2sovereignchainRemoveLegacySovereignTokenAddress)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "RemoveLegacySovereignTokenAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainSetBridgeManagerIterator is returned from FilterSetBridgeManager and is used to iterate over the raw logs and unpacked data for SetBridgeManager events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainSetBridgeManagerIterator struct {
	Event *Bridgel2sovereignchainSetBridgeManager // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainSetBridgeManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainSetBridgeManager)
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
		it.Event = new(Bridgel2sovereignchainSetBridgeManager)
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
func (it *Bridgel2sovereignchainSetBridgeManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainSetBridgeManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainSetBridgeManager represents a SetBridgeManager event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainSetBridgeManager struct {
	BridgeManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetBridgeManager is a free log retrieval operation binding the contract event 0x32cf74f8a6d5f88593984d2cd52be5592bfa6884f5896175801a5069ef09cd67.
//
// Solidity: event SetBridgeManager(address bridgeManager)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterSetBridgeManager(opts *bind.FilterOpts) (*Bridgel2sovereignchainSetBridgeManagerIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "SetBridgeManager")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainSetBridgeManagerIterator{contract: _Bridgel2sovereignchain.contract, event: "SetBridgeManager", logs: logs, sub: sub}, nil
}

// WatchSetBridgeManager is a free log subscription operation binding the contract event 0x32cf74f8a6d5f88593984d2cd52be5592bfa6884f5896175801a5069ef09cd67.
//
// Solidity: event SetBridgeManager(address bridgeManager)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchSetBridgeManager(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainSetBridgeManager) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "SetBridgeManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainSetBridgeManager)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "SetBridgeManager", log); err != nil {
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
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseSetBridgeManager(log types.Log) (*Bridgel2sovereignchainSetBridgeManager, error) {
	event := new(Bridgel2sovereignchainSetBridgeManager)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "SetBridgeManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainSetClaimIterator is returned from FilterSetClaim and is used to iterate over the raw logs and unpacked data for SetClaim events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainSetClaimIterator struct {
	Event *Bridgel2sovereignchainSetClaim // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainSetClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainSetClaim)
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
		it.Event = new(Bridgel2sovereignchainSetClaim)
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
func (it *Bridgel2sovereignchainSetClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainSetClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainSetClaim represents a SetClaim event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainSetClaim struct {
	LeafIndex     uint32
	SourceNetwork uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetClaim is a free log retrieval operation binding the contract event 0xab48452e173be3f004b3208cc344b14cc7f8d977fc1ca28171286a7abb5d7570.
//
// Solidity: event SetClaim(uint32 leafIndex, uint32 sourceNetwork)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterSetClaim(opts *bind.FilterOpts) (*Bridgel2sovereignchainSetClaimIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "SetClaim")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainSetClaimIterator{contract: _Bridgel2sovereignchain.contract, event: "SetClaim", logs: logs, sub: sub}, nil
}

// WatchSetClaim is a free log subscription operation binding the contract event 0xab48452e173be3f004b3208cc344b14cc7f8d977fc1ca28171286a7abb5d7570.
//
// Solidity: event SetClaim(uint32 leafIndex, uint32 sourceNetwork)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchSetClaim(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainSetClaim) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "SetClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainSetClaim)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "SetClaim", log); err != nil {
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

// ParseSetClaim is a log parse operation binding the contract event 0xab48452e173be3f004b3208cc344b14cc7f8d977fc1ca28171286a7abb5d7570.
//
// Solidity: event SetClaim(uint32 leafIndex, uint32 sourceNetwork)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseSetClaim(log types.Log) (*Bridgel2sovereignchainSetClaim, error) {
	event := new(Bridgel2sovereignchainSetClaim)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "SetClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainSetLocalBalanceTreeIterator is returned from FilterSetLocalBalanceTree and is used to iterate over the raw logs and unpacked data for SetLocalBalanceTree events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainSetLocalBalanceTreeIterator struct {
	Event *Bridgel2sovereignchainSetLocalBalanceTree // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainSetLocalBalanceTreeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainSetLocalBalanceTree)
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
		it.Event = new(Bridgel2sovereignchainSetLocalBalanceTree)
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
func (it *Bridgel2sovereignchainSetLocalBalanceTreeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainSetLocalBalanceTreeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainSetLocalBalanceTree represents a SetLocalBalanceTree event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainSetLocalBalanceTree struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
	NewAmount          *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSetLocalBalanceTree is a free log retrieval operation binding the contract event 0x73a03a6d9efc14b9f22a3af967e98580549eb76b1113b6a09a57ce1dae36ecd2.
//
// Solidity: event SetLocalBalanceTree(uint32 indexed originNetwork, address indexed originTokenAddress, uint256 newAmount)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterSetLocalBalanceTree(opts *bind.FilterOpts, originNetwork []uint32, originTokenAddress []common.Address) (*Bridgel2sovereignchainSetLocalBalanceTreeIterator, error) {

	var originNetworkRule []interface{}
	for _, originNetworkItem := range originNetwork {
		originNetworkRule = append(originNetworkRule, originNetworkItem)
	}
	var originTokenAddressRule []interface{}
	for _, originTokenAddressItem := range originTokenAddress {
		originTokenAddressRule = append(originTokenAddressRule, originTokenAddressItem)
	}

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "SetLocalBalanceTree", originNetworkRule, originTokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainSetLocalBalanceTreeIterator{contract: _Bridgel2sovereignchain.contract, event: "SetLocalBalanceTree", logs: logs, sub: sub}, nil
}

// WatchSetLocalBalanceTree is a free log subscription operation binding the contract event 0x73a03a6d9efc14b9f22a3af967e98580549eb76b1113b6a09a57ce1dae36ecd2.
//
// Solidity: event SetLocalBalanceTree(uint32 indexed originNetwork, address indexed originTokenAddress, uint256 newAmount)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchSetLocalBalanceTree(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainSetLocalBalanceTree, originNetwork []uint32, originTokenAddress []common.Address) (event.Subscription, error) {

	var originNetworkRule []interface{}
	for _, originNetworkItem := range originNetwork {
		originNetworkRule = append(originNetworkRule, originNetworkItem)
	}
	var originTokenAddressRule []interface{}
	for _, originTokenAddressItem := range originTokenAddress {
		originTokenAddressRule = append(originTokenAddressRule, originTokenAddressItem)
	}

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "SetLocalBalanceTree", originNetworkRule, originTokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainSetLocalBalanceTree)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "SetLocalBalanceTree", log); err != nil {
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

// ParseSetLocalBalanceTree is a log parse operation binding the contract event 0x73a03a6d9efc14b9f22a3af967e98580549eb76b1113b6a09a57ce1dae36ecd2.
//
// Solidity: event SetLocalBalanceTree(uint32 indexed originNetwork, address indexed originTokenAddress, uint256 newAmount)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseSetLocalBalanceTree(log types.Log) (*Bridgel2sovereignchainSetLocalBalanceTree, error) {
	event := new(Bridgel2sovereignchainSetLocalBalanceTree)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "SetLocalBalanceTree", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainSetSovereignTokenAddressIterator is returned from FilterSetSovereignTokenAddress and is used to iterate over the raw logs and unpacked data for SetSovereignTokenAddress events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainSetSovereignTokenAddressIterator struct {
	Event *Bridgel2sovereignchainSetSovereignTokenAddress // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainSetSovereignTokenAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainSetSovereignTokenAddress)
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
		it.Event = new(Bridgel2sovereignchainSetSovereignTokenAddress)
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
func (it *Bridgel2sovereignchainSetSovereignTokenAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainSetSovereignTokenAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainSetSovereignTokenAddress represents a SetSovereignTokenAddress event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainSetSovereignTokenAddress struct {
	OriginNetwork         uint32
	OriginTokenAddress    common.Address
	SovereignTokenAddress common.Address
	IsNotMintable         bool
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetSovereignTokenAddress is a free log retrieval operation binding the contract event 0xdbe8a5da6a7a916d9adfda9160167a0f8a3da415ee6610e810e753853597fce7.
//
// Solidity: event SetSovereignTokenAddress(uint32 originNetwork, address originTokenAddress, address sovereignTokenAddress, bool isNotMintable)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterSetSovereignTokenAddress(opts *bind.FilterOpts) (*Bridgel2sovereignchainSetSovereignTokenAddressIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "SetSovereignTokenAddress")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainSetSovereignTokenAddressIterator{contract: _Bridgel2sovereignchain.contract, event: "SetSovereignTokenAddress", logs: logs, sub: sub}, nil
}

// WatchSetSovereignTokenAddress is a free log subscription operation binding the contract event 0xdbe8a5da6a7a916d9adfda9160167a0f8a3da415ee6610e810e753853597fce7.
//
// Solidity: event SetSovereignTokenAddress(uint32 originNetwork, address originTokenAddress, address sovereignTokenAddress, bool isNotMintable)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchSetSovereignTokenAddress(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainSetSovereignTokenAddress) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "SetSovereignTokenAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainSetSovereignTokenAddress)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "SetSovereignTokenAddress", log); err != nil {
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
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseSetSovereignTokenAddress(log types.Log) (*Bridgel2sovereignchainSetSovereignTokenAddress, error) {
	event := new(Bridgel2sovereignchainSetSovereignTokenAddress)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "SetSovereignTokenAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainSetSovereignWETHAddressIterator is returned from FilterSetSovereignWETHAddress and is used to iterate over the raw logs and unpacked data for SetSovereignWETHAddress events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainSetSovereignWETHAddressIterator struct {
	Event *Bridgel2sovereignchainSetSovereignWETHAddress // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainSetSovereignWETHAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainSetSovereignWETHAddress)
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
		it.Event = new(Bridgel2sovereignchainSetSovereignWETHAddress)
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
func (it *Bridgel2sovereignchainSetSovereignWETHAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainSetSovereignWETHAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainSetSovereignWETHAddress represents a SetSovereignWETHAddress event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainSetSovereignWETHAddress struct {
	SovereignWETHTokenAddress common.Address
	IsNotMintable             bool
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterSetSovereignWETHAddress is a free log retrieval operation binding the contract event 0xc7318b7ed6ba4f2908a3de396d8ab49b1dadb55db5b55123247a401f29ff8d82.
//
// Solidity: event SetSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterSetSovereignWETHAddress(opts *bind.FilterOpts) (*Bridgel2sovereignchainSetSovereignWETHAddressIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "SetSovereignWETHAddress")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainSetSovereignWETHAddressIterator{contract: _Bridgel2sovereignchain.contract, event: "SetSovereignWETHAddress", logs: logs, sub: sub}, nil
}

// WatchSetSovereignWETHAddress is a free log subscription operation binding the contract event 0xc7318b7ed6ba4f2908a3de396d8ab49b1dadb55db5b55123247a401f29ff8d82.
//
// Solidity: event SetSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchSetSovereignWETHAddress(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainSetSovereignWETHAddress) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "SetSovereignWETHAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainSetSovereignWETHAddress)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "SetSovereignWETHAddress", log); err != nil {
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
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseSetSovereignWETHAddress(log types.Log) (*Bridgel2sovereignchainSetSovereignWETHAddress, error) {
	event := new(Bridgel2sovereignchainSetSovereignWETHAddress)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "SetSovereignWETHAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainTransferEmergencyBridgePauserRoleIterator is returned from FilterTransferEmergencyBridgePauserRole and is used to iterate over the raw logs and unpacked data for TransferEmergencyBridgePauserRole events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainTransferEmergencyBridgePauserRoleIterator struct {
	Event *Bridgel2sovereignchainTransferEmergencyBridgePauserRole // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainTransferEmergencyBridgePauserRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainTransferEmergencyBridgePauserRole)
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
		it.Event = new(Bridgel2sovereignchainTransferEmergencyBridgePauserRole)
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
func (it *Bridgel2sovereignchainTransferEmergencyBridgePauserRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainTransferEmergencyBridgePauserRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainTransferEmergencyBridgePauserRole represents a TransferEmergencyBridgePauserRole event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainTransferEmergencyBridgePauserRole struct {
	CurrentEmergencyBridgePauser common.Address
	NewEmergencyBridgePauser     common.Address
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterTransferEmergencyBridgePauserRole is a free log retrieval operation binding the contract event 0xb27de219766f47b82684842855ba6130b6dbf288ac66d1c3509e7bf17f4e925a.
//
// Solidity: event TransferEmergencyBridgePauserRole(address currentEmergencyBridgePauser, address newEmergencyBridgePauser)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterTransferEmergencyBridgePauserRole(opts *bind.FilterOpts) (*Bridgel2sovereignchainTransferEmergencyBridgePauserRoleIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "TransferEmergencyBridgePauserRole")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainTransferEmergencyBridgePauserRoleIterator{contract: _Bridgel2sovereignchain.contract, event: "TransferEmergencyBridgePauserRole", logs: logs, sub: sub}, nil
}

// WatchTransferEmergencyBridgePauserRole is a free log subscription operation binding the contract event 0xb27de219766f47b82684842855ba6130b6dbf288ac66d1c3509e7bf17f4e925a.
//
// Solidity: event TransferEmergencyBridgePauserRole(address currentEmergencyBridgePauser, address newEmergencyBridgePauser)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchTransferEmergencyBridgePauserRole(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainTransferEmergencyBridgePauserRole) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "TransferEmergencyBridgePauserRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainTransferEmergencyBridgePauserRole)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "TransferEmergencyBridgePauserRole", log); err != nil {
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
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseTransferEmergencyBridgePauserRole(log types.Log) (*Bridgel2sovereignchainTransferEmergencyBridgePauserRole, error) {
	event := new(Bridgel2sovereignchainTransferEmergencyBridgePauserRole)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "TransferEmergencyBridgePauserRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRoleIterator is returned from FilterTransferEmergencyBridgeUnpauserRole and is used to iterate over the raw logs and unpacked data for TransferEmergencyBridgeUnpauserRole events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRoleIterator struct {
	Event *Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRole // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRole)
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
		it.Event = new(Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRole)
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
func (it *Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRole represents a TransferEmergencyBridgeUnpauserRole event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRole struct {
	CurrentEmergencyBridgeUnpauser common.Address
	NewEmergencyBridgeUnpauser     common.Address
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterTransferEmergencyBridgeUnpauserRole is a free log retrieval operation binding the contract event 0xf01a62a06940517bbc898dec8c75794b9feabcd2d263c8de823b36dbbeb8779b.
//
// Solidity: event TransferEmergencyBridgeUnpauserRole(address currentEmergencyBridgeUnpauser, address newEmergencyBridgeUnpauser)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterTransferEmergencyBridgeUnpauserRole(opts *bind.FilterOpts) (*Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRoleIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "TransferEmergencyBridgeUnpauserRole")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRoleIterator{contract: _Bridgel2sovereignchain.contract, event: "TransferEmergencyBridgeUnpauserRole", logs: logs, sub: sub}, nil
}

// WatchTransferEmergencyBridgeUnpauserRole is a free log subscription operation binding the contract event 0xf01a62a06940517bbc898dec8c75794b9feabcd2d263c8de823b36dbbeb8779b.
//
// Solidity: event TransferEmergencyBridgeUnpauserRole(address currentEmergencyBridgeUnpauser, address newEmergencyBridgeUnpauser)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchTransferEmergencyBridgeUnpauserRole(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRole) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "TransferEmergencyBridgeUnpauserRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRole)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "TransferEmergencyBridgeUnpauserRole", log); err != nil {
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
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseTransferEmergencyBridgeUnpauserRole(log types.Log) (*Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRole, error) {
	event := new(Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRole)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "TransferEmergencyBridgeUnpauserRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainTransferProxiedTokensManagerRoleIterator is returned from FilterTransferProxiedTokensManagerRole and is used to iterate over the raw logs and unpacked data for TransferProxiedTokensManagerRole events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainTransferProxiedTokensManagerRoleIterator struct {
	Event *Bridgel2sovereignchainTransferProxiedTokensManagerRole // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainTransferProxiedTokensManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainTransferProxiedTokensManagerRole)
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
		it.Event = new(Bridgel2sovereignchainTransferProxiedTokensManagerRole)
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
func (it *Bridgel2sovereignchainTransferProxiedTokensManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainTransferProxiedTokensManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainTransferProxiedTokensManagerRole represents a TransferProxiedTokensManagerRole event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainTransferProxiedTokensManagerRole struct {
	CurrentProxiedTokensManager common.Address
	NewProxiedTokensManager     common.Address
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterTransferProxiedTokensManagerRole is a free log retrieval operation binding the contract event 0x0a34baa3feb299aef9c05cb59c6e0c8e7c0bcc65cbf0a647e7a7c8a2411591e2.
//
// Solidity: event TransferProxiedTokensManagerRole(address currentProxiedTokensManager, address newProxiedTokensManager)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterTransferProxiedTokensManagerRole(opts *bind.FilterOpts) (*Bridgel2sovereignchainTransferProxiedTokensManagerRoleIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "TransferProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainTransferProxiedTokensManagerRoleIterator{contract: _Bridgel2sovereignchain.contract, event: "TransferProxiedTokensManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferProxiedTokensManagerRole is a free log subscription operation binding the contract event 0x0a34baa3feb299aef9c05cb59c6e0c8e7c0bcc65cbf0a647e7a7c8a2411591e2.
//
// Solidity: event TransferProxiedTokensManagerRole(address currentProxiedTokensManager, address newProxiedTokensManager)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchTransferProxiedTokensManagerRole(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainTransferProxiedTokensManagerRole) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "TransferProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainTransferProxiedTokensManagerRole)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "TransferProxiedTokensManagerRole", log); err != nil {
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
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseTransferProxiedTokensManagerRole(log types.Log) (*Bridgel2sovereignchainTransferProxiedTokensManagerRole, error) {
	event := new(Bridgel2sovereignchainTransferProxiedTokensManagerRole)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "TransferProxiedTokensManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChainIterator is returned from FilterUpdatedClaimedGlobalIndexHashChain and is used to iterate over the raw logs and unpacked data for UpdatedClaimedGlobalIndexHashChain events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChainIterator struct {
	Event *Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChain // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChain)
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
		it.Event = new(Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChain)
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
func (it *Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChain represents a UpdatedClaimedGlobalIndexHashChain event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChain struct {
	ClaimedGlobalIndex             [32]byte
	NewClaimedGlobalIndexHashChain [32]byte
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterUpdatedClaimedGlobalIndexHashChain is a free log retrieval operation binding the contract event 0x3e5936f910a78eb5181813a939c8d4c3e4d85f87943f659380d82ac6221b0e92.
//
// Solidity: event UpdatedClaimedGlobalIndexHashChain(bytes32 claimedGlobalIndex, bytes32 newClaimedGlobalIndexHashChain)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterUpdatedClaimedGlobalIndexHashChain(opts *bind.FilterOpts) (*Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChainIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "UpdatedClaimedGlobalIndexHashChain")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChainIterator{contract: _Bridgel2sovereignchain.contract, event: "UpdatedClaimedGlobalIndexHashChain", logs: logs, sub: sub}, nil
}

// WatchUpdatedClaimedGlobalIndexHashChain is a free log subscription operation binding the contract event 0x3e5936f910a78eb5181813a939c8d4c3e4d85f87943f659380d82ac6221b0e92.
//
// Solidity: event UpdatedClaimedGlobalIndexHashChain(bytes32 claimedGlobalIndex, bytes32 newClaimedGlobalIndexHashChain)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchUpdatedClaimedGlobalIndexHashChain(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChain) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "UpdatedClaimedGlobalIndexHashChain")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChain)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "UpdatedClaimedGlobalIndexHashChain", log); err != nil {
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
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseUpdatedClaimedGlobalIndexHashChain(log types.Log) (*Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChain, error) {
	event := new(Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChain)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "UpdatedClaimedGlobalIndexHashChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChainIterator is returned from FilterUpdatedUnsetGlobalIndexHashChain and is used to iterate over the raw logs and unpacked data for UpdatedUnsetGlobalIndexHashChain events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChainIterator struct {
	Event *Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChain // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChain)
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
		it.Event = new(Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChain)
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
func (it *Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChain represents a UpdatedUnsetGlobalIndexHashChain event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChain struct {
	UnsetGlobalIndex             [32]byte
	NewUnsetGlobalIndexHashChain [32]byte
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedUnsetGlobalIndexHashChain is a free log retrieval operation binding the contract event 0xc80e0aca446a59735359a7ae46124b57c47b892827642779bc6dafc84ba90b03.
//
// Solidity: event UpdatedUnsetGlobalIndexHashChain(bytes32 unsetGlobalIndex, bytes32 newUnsetGlobalIndexHashChain)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterUpdatedUnsetGlobalIndexHashChain(opts *bind.FilterOpts) (*Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChainIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "UpdatedUnsetGlobalIndexHashChain")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChainIterator{contract: _Bridgel2sovereignchain.contract, event: "UpdatedUnsetGlobalIndexHashChain", logs: logs, sub: sub}, nil
}

// WatchUpdatedUnsetGlobalIndexHashChain is a free log subscription operation binding the contract event 0xc80e0aca446a59735359a7ae46124b57c47b892827642779bc6dafc84ba90b03.
//
// Solidity: event UpdatedUnsetGlobalIndexHashChain(bytes32 unsetGlobalIndex, bytes32 newUnsetGlobalIndexHashChain)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchUpdatedUnsetGlobalIndexHashChain(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChain) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "UpdatedUnsetGlobalIndexHashChain")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChain)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "UpdatedUnsetGlobalIndexHashChain", log); err != nil {
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
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseUpdatedUnsetGlobalIndexHashChain(log types.Log) (*Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChain, error) {
	event := new(Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChain)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "UpdatedUnsetGlobalIndexHashChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

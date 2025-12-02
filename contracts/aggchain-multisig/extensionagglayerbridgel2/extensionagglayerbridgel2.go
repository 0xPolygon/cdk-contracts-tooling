// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package extensionagglayerbridgel2

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

// AgglayerBridgeL2ClaimData is an auto generated low-level Go binding around an user-defined struct.
type AgglayerBridgeL2ClaimData struct {
	SmtProofLocalExitRoot  [32][32]byte
	SmtProofRollupExitRoot [32][32]byte
	GlobalIndex            *big.Int
	MainnetExitRoot        [32]byte
	RollupExitRoot         [32]byte
	LeafType               uint8
	OriginNetwork          uint32
	OriginAddress          common.Address
	DestinationNetwork     uint32
	DestinationAddress     common.Address
	Amount                 *big.Int
	Metadata               []byte
}

// AgglayerBridgeL2LeafData is an auto generated low-level Go binding around an user-defined struct.
type AgglayerBridgeL2LeafData struct {
	LeafType           uint8
	OriginNetwork      uint32
	OriginAddress      common.Address
	DestinationNetwork uint32
	DestinationAddress common.Address
	Amount             *big.Int
	Metadata           []byte
}

// Extensionagglayerbridgel2MetaData contains all meta data concerning the Extensionagglayerbridgel2 contract.
var Extensionagglayerbridgel2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountDoesNotMatchMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeAddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DestinationNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmergencyStateNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EtherTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedProxyDeployment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputArraysLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDepositCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExpectedLER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGlobalIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLBTLeaf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLeafType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLeavesLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxyAdmin\",\"type\":\"address\"}],\"name\":\"InvalidProxyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSmtProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSovereignWETHAddressParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSubtreeFrontier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroNetworkID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxyAdmin\",\"type\":\"address\"}],\"name\":\"InvalidZeroProxyAdminOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"localBalanceTreeAmount\",\"type\":\"uint256\"}],\"name\":\"LocalBalanceTreeOverflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"localBalanceTreeAmount\",\"type\":\"uint256\"}],\"name\":\"LocalBalanceTreeUnderflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MsgValueNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeTokenIsEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewDepositCountExceedsMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoValueInMessagesOnGasTokenNetworks\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonSupportedFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonZeroValueForUnusedFrontier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyBridgeManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyDeployer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyBridgePauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyBridgeUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootRemover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingEmergencyBridgePauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingEmergencyBridgeUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingProxiedTokensManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProxiedTokensManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OriginNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubtreeFrontierMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyMapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyUpdated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotMapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotRemapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WETHRemappingNotSupportedOnGasTokenNetworks\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldEmergencyBridgePauser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newEmergencyBridgePauser\",\"type\":\"address\"}],\"name\":\"AcceptEmergencyBridgePauserRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldEmergencyBridgeUnpauser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newEmergencyBridgeUnpauser\",\"type\":\"address\"}],\"name\":\"AcceptEmergencyBridgeUnpauserRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldProxiedTokensManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newProxiedTokensManager\",\"type\":\"address\"}],\"name\":\"AcceptProxiedTokensManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousDepositCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDepositCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"BackwardLET\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"depositCount\",\"type\":\"uint32\"}],\"name\":\"BridgeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"indexed\":false,\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"DetailedClaimEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousDepositCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDepositCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newLeaves\",\"type\":\"bytes\"}],\"name\":\"ForwardLET\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"legacyTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"updatedTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MigrateLegacyToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wrappedTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"NewWrappedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignTokenAddress\",\"type\":\"address\"}],\"name\":\"RemoveLegacySovereignTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridgeManager\",\"type\":\"address\"}],\"name\":\"SetBridgeManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"globalIndex\",\"type\":\"bytes32\"}],\"name\":\"SetClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"SetLocalBalanceTree\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"SetSovereignTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignWETHTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"SetSovereignWETHAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentEmergencyBridgePauser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newEmergencyBridgePauser\",\"type\":\"address\"}],\"name\":\"TransferEmergencyBridgePauserRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentEmergencyBridgeUnpauser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newEmergencyBridgeUnpauser\",\"type\":\"address\"}],\"name\":\"TransferEmergencyBridgeUnpauserRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentProxiedTokensManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newProxiedTokensManager\",\"type\":\"address\"}],\"name\":\"TransferProxiedTokensManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"claimedGlobalIndex\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newClaimedGlobalIndexHashChain\",\"type\":\"bytes32\"}],\"name\":\"UpdatedClaimedGlobalIndexHashChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"unsetGlobalIndex\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newUnsetGlobalIndexHashChain\",\"type\":\"bytes32\"}],\"name\":\"UpdatedUnsetGlobalIndexHashChain\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INIT_BYTECODE_TRANSPARENT_PROXY\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETHToken\",\"outputs\":[{\"internalType\":\"contractITokenWrappedBridgeUpgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptEmergencyBridgePauserRole\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptEmergencyBridgeUnpauserRole\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptProxiedTokensManagerRole\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[32]\",\"name\":\"\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"\",\"type\":\"bytes32[32]\"}],\"name\":\"backwardLET\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"bridgeAsset\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeLib\",\"outputs\":[{\"internalType\":\"contractBridgeLib\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"bridgeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"bridgeMessageWETH\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"claimAsset\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"claimMessage\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimedBitMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimedGlobalIndexHashChain\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"computeTokenProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"deployWrappedTokenAndRemap\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyBridgePauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyBridgeUnpauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structAgglayerBridgeL2.ClaimData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"name\":\"forceEmitDetailedClaimEvent\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structAgglayerBridgeL2.LeafData[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"forwardLET\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProxiedTokensManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getTokenWrappedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWrappedTokenBridgeImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIBaseLegacyAgglayerGER\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"contractIBaseLegacyAgglayerGER\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_polygonRollupManager\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_bridgeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sovereignWETHAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_sovereignWETHAddressIsNotMintable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_emergencyBridgePauser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_emergencyBridgeUnpauser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proxiedTokensManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"contractIBaseLegacyAgglayerGER\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"isClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdatedDepositCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tokenInfoHash\",\"type\":\"bytes32\"}],\"name\":\"localBalanceTree\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"migrateLegacyToken\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingEmergencyBridgePauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingEmergencyBridgeUnpauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingProxiedTokensManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"polygonRollupManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"removeLegacySovereignTokenAddress\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setBridgeManager\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"\",\"type\":\"uint32[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"name\":\"setLocalBalanceTree\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"name\":\"setMultipleClaims\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"\",\"type\":\"uint32[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"name\":\"setMultipleSovereignTokenAddress\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"setSovereignWETHAddress\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenInfoToWrappedToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferEmergencyBridgePauserRole\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferEmergencyBridgeUnpauserRole\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferProxiedTokensManagerRole\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unsetGlobalIndexHashChain\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"name\":\"unsetMultipleClaims\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrappedAddress\",\"type\":\"address\"}],\"name\":\"wrappedAddressIsNotMintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wrappedTokenToTokenInfo\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x610100604052348015610010575f5ffd5b5060405161001d90610157565b604051809103905ff080158015610036573d5f5f3e3d5ffd5b506001600160a01b031660a05260405161004f90610164565b604051809103905ff080158015610068573d5f5f3e3d5ffd5b506001600160a01b031660805261007d61009a565b3360c05261008961009a565b3360e05261009561009a565b610171565b5f54610100900460ff16156101055760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161015610155575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b611d36806127ad83390190565b612001806144e383390190565b60805160a05160c05160e05161260f61019e5f395f610b7301525f50505f50505f610764015261260f5ff3fe608060405260043610610391575f3560e01c80638b37b873116101de578063c514f24e11610108578063eabd372a1161009d578063f214e1611161006d578063f214e161146104bc578063f5efcd7914610948578063f811bff714610a41578063fd7640e814610a60575f5ffd5b8063eabd372a1461070a578063ece93c6f146109d0578063ee25560b146109fc578063f0e7080814610a27575f5ffd5b8063d02103ca116100d8578063d02103ca14610970578063d9cb3aec146109a5578063dbc16976146104a8578063e88f0436146104a8575f5ffd5b8063c514f24e14610682578063cc4616321461092e578063ccaa2d1114610948578063cd58657914610962575f5ffd5b8063b0b379201161017e578063be5831c71161014e578063be5831c71461088e578063bf130d7f146108c7578063c00f14ab146108e1578063c0f4916314610900575f5ffd5b8063b0b3792014610839578063b45869621461070a578063b8b284d014610853578063bab161bf1461086d575f5ffd5b80638d942096116101b95780638d9420961461070a5780638ed7e3f2146107c75780638f9720bb146107f3578063ae24490a1461080d575f5ffd5b80638b37b873146104a85780638bd309c31461070a5780638c668f1c146104a8575f5ffd5b80633b2fee9a116102bf5780635ca1e1651161025f5780636ee84b231161022f5780636ee84b231461073e5780636f0bc3da1461075357806379e2cf97146104a857806381b1c17414610786575f5ffd5b80635ca1e165146106ca578063606617ff146106de57806369e3ab121461070a5780636e974cd414610724575f5ffd5b80634b2f336d1161029a5780634b2f336d1461065657806354fd4d5014610682578063567f13e41461069657806357cfbee3146106b0575f5ffd5b80633b2fee9a146105cd5780633c351e10146105e15780633cbc795b1461060d575f5ffd5b80632072f6c5116103355780632dfdf0b5116103055780632dfdf0b51461050a5780632f84c6901461051f578063318aee3d1461054b57806338b8fbbb146105cd575f5ffd5b80632072f6c5146104a857806322e95f2c146104bc578063240ff378146104db57806327aef4e8146104e9575f5ffd5b806314cc01a01161037057806314cc01a01461043057806315064c961461045c5780631c208229146104115780631d081d8c14610485575f5ffd5b80626ee1711461039557806303e6e116146103b6578063136a2c6014610411575b5f5ffd5b3480156103a0575f5ffd5b506103b46103af3660046118ec565b610a7a565b005b3480156103c1575f5ffd5b5060a8546103e790610100900473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561041c575f5ffd5b506103b461042b366004611a65565b6112dc565b34801561043b575f5ffd5b5060a3546103e79073ffffffffffffffffffffffffffffffffffffffff1681565b348015610467575f5ffd5b506068546104759060ff1681565b6040519015158152602001610408565b348015610490575f5ffd5b5061049a60a55481565b604051908152602001610408565b3480156104b3575f5ffd5b506103b46112dc565b3480156104c7575f5ffd5b506103e76104d6366004611a9f565b61130e565b6103b461042b366004611b19565b3480156104f4575f5ffd5b506104fd611341565b6040516104089190611bd8565b348015610515575f5ffd5b5061049a60535481565b34801561052a575f5ffd5b5060a4546103e79073ffffffffffffffffffffffffffffffffffffffff1681565b348015610556575f5ffd5b5061059c610565366004611bf1565b606b6020525f908152604090205463ffffffff811690640100000000900473ffffffffffffffffffffffffffffffffffffffff1682565b6040805163ffffffff909316835273ffffffffffffffffffffffffffffffffffffffff909116602083015201610408565b3480156105d8575f5ffd5b506103e761130e565b3480156105ec575f5ffd5b50606d546103e79073ffffffffffffffffffffffffffffffffffffffff1681565b348015610618575f5ffd5b50606d546106419074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff9091168152602001610408565b348015610661575f5ffd5b50606f546103e79073ffffffffffffffffffffffffffffffffffffffff1681565b34801561068d575f5ffd5b506104fd6113cd565b3480156106a1575f5ffd5b506103b461042b366004611c4d565b3480156106bb575f5ffd5b506103b461042b366004611d52565b3480156106d5575f5ffd5b5061049a61130e565b3480156106e9575f5ffd5b5060aa546103e79073ffffffffffffffffffffffffffffffffffffffff1681565b348015610715575f5ffd5b506103b461042b366004611bf1565b34801561072f575f5ffd5b506103b461042b366004611e69565b348015610749575f5ffd5b5061049a60a65481565b34801561075e575f5ffd5b506103e77f000000000000000000000000000000000000000000000000000000000000000081565b348015610791575f5ffd5b506103e76107a0366004611eab565b606a6020525f908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b3480156107d2575f5ffd5b50606c546103e79073ffffffffffffffffffffffffffffffffffffffff1681565b3480156107fe575f5ffd5b506103b461042b366004611ec2565b348015610818575f5ffd5b5060a9546103e79073ffffffffffffffffffffffffffffffffffffffff1681565b348015610844575f5ffd5b506103b461042b366004611f50565b34801561085e575f5ffd5b506103b461042b366004611fa8565b348015610878575f5ffd5b5060685461064190610100900463ffffffff1681565b348015610899575f5ffd5b5060685461064190790100000000000000000000000000000000000000000000000000900463ffffffff1681565b3480156108d2575f5ffd5b506103b461042b366004612024565b3480156108ec575f5ffd5b506104fd6108fb366004611bf1565b6113cd565b34801561090b575f5ffd5b5061047561091a366004611bf1565b60a26020525f908152604090205460ff1681565b348015610939575f5ffd5b506104756104d6366004612057565b348015610953575f5ffd5b506103b461042b366004612096565b6103b461042b366004612175565b34801561097b575f5ffd5b506068546103e79065010000000000900473ffffffffffffffffffffffffffffffffffffffff1681565b3480156109b0575f5ffd5b5061049a6109bf366004611eab565b60a76020525f908152604090205481565b3480156109db575f5ffd5b506071546103e79073ffffffffffffffffffffffffffffffffffffffff1681565b348015610a07575f5ffd5b5061049a610a16366004611eab565b60696020525f908152604090205481565b348015610a32575f5ffd5b506103b461042b366004612203565b348015610a4c575f5ffd5b506103b4610a5b36600461224b565b611401565b348015610a6b575f5ffd5b506103b461042b3660046122db565b5f54600390610100900460ff16158015610a9a57505f5460ff8083169116105b610b2b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff8316176101001790553373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610bca576040517f618bbdd500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8a16610c17576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8c63ffffffff165f03610c56576040517f4e702fa500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8c606860016101000a81548163ffffffff021916908363ffffffff16021790555089606860056101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555088606c5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508660a35f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508360a45f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f85d2bdfbe58cd81abf8199c13ce2509204be4aba8603b9d29f52c4e13e7bb7935f60a45f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051610dec92919073ffffffffffffffffffffffffffffffffffffffff92831681529116602082015260400190565b60405180910390a160a980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8516908117909155604080515f815260208101929092527f24cc8295aa5110cc216695db944ad2458c7795c6404449be980c3ce14aed752d910160405180910390a13073ffffffffffffffffffffffffffffffffffffffff831603610ec0576040517f1ae0e03300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216610f0d576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8416908117909155604080515f815260208101929092527fa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f2910160405180910390a173ffffffffffffffffffffffffffffffffffffffff8c166110405763ffffffff8b1615610fe3576040517f1a874c1200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff86161515806110045750845b1561103b576040517f1cdc46ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611268565b606d805463ffffffff8d1674010000000000000000000000000000000000000000027fffffffffffffffff00000000000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff8f1617179055606e6110ad89826123c0565b5073ffffffffffffffffffffffffffffffffffffffff86166111ed57841515600103611105576040517f1cdc46ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111a35f5f1b601260405160200161118f91906060808252600d908201527f5772617070656420457468657200000000000000000000000000000000000000608082015260a0602082018190526004908201527f574554480000000000000000000000000000000000000000000000000000000060c082015260ff91909116604082015260e00190565b604051602081830303815290604052611552565b606f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055611268565b606f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff88169081179091555f90815260a26020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168615151790555b611270611647565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050505050505050505050565b6040517f31f7db8500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6040517f31f7db8500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606e805461134e90612323565b80601f016020809104026020016040519081016040528092919081815260200182805461137a90612323565b80156113c55780601f1061139c576101008083540402835291602001916113c5565b820191905f5260205f20905b8154815290600101906020018083116113a857829003601f168201915b505050505081565b60606040517f31f7db8500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f54610100900460ff161580801561141f57505f54600160ff909116105b806114385750303b15801561143857505f5460ff166001145b6114c4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610b22565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015611520575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6040517ff57ac68300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f61155c6113cd565b9050838151602083015ff5915073ffffffffffffffffffffffffffffffffffffffff82166115b6576040517f62d05d1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f5f858060200190518101906115cd9190612524565b9250925092508473ffffffffffffffffffffffffffffffffffffffff16631624f6c68484846040518463ffffffff1660e01b8152600401611610939291906125a1565b5f604051808303815f87803b158015611627575f5ffd5b505af1158015611639573d5f5f3e3d5ffd5b505050505050505092915050565b5f54610100900460ff166116dd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610b22565b6116e56116e7565b565b5f54610100900460ff1661177d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610b22565b60018055565b803563ffffffff81168114611796575f5ffd5b919050565b73ffffffffffffffffffffffffffffffffffffffff811681146117bc575f5ffd5b50565b80356117968161179b565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561183e5761183e6117ca565b604052919050565b5f67ffffffffffffffff82111561185f5761185f6117ca565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b5f82601f83011261189a575f5ffd5b81356118ad6118a882611846565b6117f7565b8181528460208386010111156118c1575f5ffd5b816020850160208301375f918101602001919091529392505050565b80358015158114611796575f5ffd5b5f5f5f5f5f5f5f5f5f5f5f5f6101808d8f031215611908575f5ffd5b6119118d611783565b9b5061191f60208e016117bf565b9a5061192d60408e01611783565b995061193b60608e016117bf565b985061194960808e016117bf565b975067ffffffffffffffff60a08e01351115611963575f5ffd5b6119738e60a08f01358f0161188b565b965061198160c08e016117bf565b955061198f60e08e016117bf565b945061199e6101008e016118dd565b93506119ad6101208e016117bf565b92506119bc6101408e016117bf565b91506119cb6101608e016117bf565b90509295989b509295989b509295989b565b5f67ffffffffffffffff8211156119f6576119f66117ca565b5060051b60200190565b5f82601f830112611a0f575f5ffd5b8135611a1d6118a8826119dd565b8082825260208201915060208360051b860101925085831115611a3e575f5ffd5b602085015b83811015611a5b578035835260209283019201611a43565b5095945050505050565b5f60208284031215611a75575f5ffd5b813567ffffffffffffffff811115611a8b575f5ffd5b611a9784828501611a00565b949350505050565b5f5f60408385031215611ab0575f5ffd5b611ab983611783565b91506020830135611ac98161179b565b809150509250929050565b5f5f83601f840112611ae4575f5ffd5b50813567ffffffffffffffff811115611afb575f5ffd5b602083019150836020828501011115611b12575f5ffd5b9250929050565b5f5f5f5f5f60808688031215611b2d575f5ffd5b611b3686611783565b94506020860135611b468161179b565b9350611b54604087016118dd565b9250606086013567ffffffffffffffff811115611b6f575f5ffd5b611b7b88828901611ad4565b969995985093965092949392505050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f611bea6020830184611b8c565b9392505050565b5f60208284031215611c01575f5ffd5b8135611bea8161179b565b5f5f83601f840112611c1c575f5ffd5b50813567ffffffffffffffff811115611c33575f5ffd5b6020830191508360208260051b8501011115611b12575f5ffd5b5f5f60208385031215611c5e575f5ffd5b823567ffffffffffffffff811115611c74575f5ffd5b611c8085828601611c0c565b90969095509350505050565b5f82601f830112611c9b575f5ffd5b8135611ca96118a8826119dd565b8082825260208201915060208360051b860101925085831115611cca575f5ffd5b602085015b83811015611a5b57611ce081611783565b835260209283019201611ccf565b5f82601f830112611cfd575f5ffd5b8135611d0b6118a8826119dd565b8082825260208201915060208360051b860101925085831115611d2c575f5ffd5b602085015b83811015611a5b578035611d448161179b565b835260209283019201611d31565b5f5f5f5f60808587031215611d65575f5ffd5b843567ffffffffffffffff811115611d7b575f5ffd5b611d8787828801611c8c565b945050602085013567ffffffffffffffff811115611da3575f5ffd5b611daf87828801611cee565b935050604085013567ffffffffffffffff811115611dcb575f5ffd5b611dd787828801611cee565b925050606085013567ffffffffffffffff811115611df3575f5ffd5b8501601f81018713611e03575f5ffd5b8035611e116118a8826119dd565b8082825260208201915060208360051b850101925089831115611e32575f5ffd5b6020840193505b82841015611e5b57611e4a846118dd565b825260209384019390910190611e39565b969995985093965050505050565b5f5f5f60608486031215611e7b575f5ffd5b611e8484611783565b92506020840135611e948161179b565b9150611ea2604085016118dd565b90509250925092565b5f60208284031215611ebb575f5ffd5b5035919050565b5f5f5f60608486031215611ed4575f5ffd5b833567ffffffffffffffff811115611eea575f5ffd5b611ef686828701611c8c565b935050602084013567ffffffffffffffff811115611f12575f5ffd5b611f1e86828701611cee565b925050604084013567ffffffffffffffff811115611f3a575f5ffd5b611f4686828701611a00565b9150509250925092565b5f5f5f5f60608587031215611f63575f5ffd5b8435611f6e8161179b565b935060208501359250604085013567ffffffffffffffff811115611f90575f5ffd5b611f9c87828801611ad4565b95989497509550505050565b5f5f5f5f5f5f60a08789031215611fbd575f5ffd5b611fc687611783565b95506020870135611fd68161179b565b945060408701359350611feb606088016118dd565b9250608087013567ffffffffffffffff811115612006575f5ffd5b61201289828a01611ad4565b979a9699509497509295939492505050565b5f5f60408385031215612035575f5ffd5b82356120408161179b565b915061204e602084016118dd565b90509250929050565b5f5f60408385031215612068575f5ffd5b61207183611783565b915061204e60208401611783565b806104008101831015612090575f5ffd5b92915050565b5f5f5f5f5f5f5f5f5f5f5f5f6109208d8f0312156120b2575f5ffd5b6120bc8e8e61207f565b9b506120cc8e6104008f0161207f565b9a506108008d013599506108208d013598506108408d013597506120f36108608e01611783565b96506121036108808e013561179b565b6108808d013595506121186108a08e01611783565b94506108c08d01356121298161179b565b93506108e08d0135925067ffffffffffffffff6109008e0135111561214c575f5ffd5b61215d8e6109008f01358f01611ad4565b81935080925050509295989b509295989b509295989b565b5f5f5f5f5f5f5f60c0888a03121561218b575f5ffd5b61219488611783565b965060208801356121a48161179b565b95506040880135945060608801356121bb8161179b565b93506121c9608089016118dd565b925060a088013567ffffffffffffffff8111156121e4575f5ffd5b6121f08a828b01611ad4565b989b979a50959850939692959293505050565b5f5f5f60408486031215612215575f5ffd5b833567ffffffffffffffff81111561222b575f5ffd5b61223786828701611c0c565b909790965060209590950135949350505050565b5f5f5f5f5f5f60c08789031215612260575f5ffd5b61226987611783565b955060208701356122798161179b565b945061228760408801611783565b935060608701356122978161179b565b925060808701356122a78161179b565b915060a087013567ffffffffffffffff8111156122c2575f5ffd5b6122ce89828a0161188b565b9150509295509295509295565b5f5f5f5f61084085870312156122ef575f5ffd5b84359350612300866020870161207f565b9250610420850135915061231886610440870161207f565b905092959194509250565b600181811c9082168061233757607f821691505b60208210810361236e577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b601f8211156123bb57805f5260205f20601f840160051c810160208510156123995750805b601f840160051c820191505b818110156123b8575f81556001016123a5565b50505b505050565b815167ffffffffffffffff8111156123da576123da6117ca565b6123ee816123e88454612323565b84612374565b6020601f82116001811461243f575f83156124095750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b1784556123b8565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b8281101561248c578785015182556020948501946001909201910161246c565b50848210156124c857868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b5f82601f8301126124e6575f5ffd5b81516124f46118a882611846565b818152846020838601011115612508575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f5f5f60608486031215612536575f5ffd5b835167ffffffffffffffff81111561254c575f5ffd5b612558868287016124d7565b935050602084015167ffffffffffffffff811115612574575f5ffd5b612580868287016124d7565b925050604084015160ff81168114612596575f5ffd5b809150509250925092565b606081525f6125b36060830186611b8c565b82810360208401526125c58186611b8c565b91505060ff8316604083015294935050505056fea2646970667358221220f3b826fc2f4cc895e26e414b9fcfd355c0f55f89addda378d522468dce555a0a64736f6c634300081c00336080604052348015600e575f5ffd5b5060156019565b60c9565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560685760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c65780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b611c60806100d65f395ff3fe608060405234801561000f575f5ffd5b5060043610610115575f3560e01c806370a08231116100ad5780639dc29fac1161007d578063a9059cbb11610063578063a9059cbb146102c0578063d505accf146102d3578063dd62ed3e146102e6575f5ffd5b80639dc29fac1461024b578063a3c573eb1461025e575f5ffd5b806370a08231146102025780637ecebe001461021557806384b0196e1461022857806395d89b4114610243575f5ffd5b806323b872dd116100e857806323b872dd146101a0578063313ce567146101b35780633644e515146101e757806340c10f19146101ef575f5ffd5b806306fdde0314610119578063095ea7b3146101375780631624f6c61461015a57806318160ddd1461016f575b5f5ffd5b61012161034a565b60405161012e919061168e565b60405180910390f35b61014a6101453660046116cf565b610402565b604051901515815260200161012e565b61016d6101683660046117fc565b61041b565b005b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02545b60405190815260200161012e565b61014a6101ae366004611870565b6105f9565b7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d005460405160ff909116815260200161012e565b61019261061c565b61016d6101fd3660046116cf565b61062a565b6101926102103660046118aa565b6106b3565b6101926102233660046118aa565b610703565b61023061070d565b60405161012e97969594939291906118c3565b61012161080c565b61016d6102593660046116cf565b61085d565b7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d0054610100900473ffffffffffffffffffffffffffffffffffffffff1660405173ffffffffffffffffffffffffffffffffffffffff909116815260200161012e565b61014a6102ce3660046116cf565b6108e1565b61016d6102e1366004611982565b6108ee565b6101926102f43660046119e8565b73ffffffffffffffffffffffffffffffffffffffff9182165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b60605f7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b905080600301805461038090611a19565b80601f01602080910402602001604051908101604052809291908181526020018280546103ac90611a19565b80156103f75780601f106103ce576101008083540402835291602001916103f7565b820191905f5260205f20905b8154815290600101906020018083116103da57829003601f168201915b505050505091505090565b5f3361040f818585610ab6565b60019150505b92915050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f811580156104655750825b90505f8267ffffffffffffffff1660011480156104815750303b155b90508115801561048f575080155b156104c6576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156105275784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6105318888610ac3565b61053a88610ad9565b7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d00805460ff88167fffffffffffffffffffffff00000000000000000000000000000000000000000090911617610100330217905583156105ef5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b5f33610606858285610b23565b610611858585610c0f565b506001949350505050565b5f610625610cb8565b905090565b5f7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d008054909150610100900473ffffffffffffffffffffffffffffffffffffffff1633146106a4576040517f38da3b1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106ae8383610cc1565b505050565b5f807f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b73ffffffffffffffffffffffffffffffffffffffff9093165f9081526020939093525050604090205490565b5f61041582610d1b565b5f60608082808083817fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100805490915015801561074b57506001810154155b6107b6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f4549503731323a20556e696e697469616c697a6564000000000000000000000060448201526064015b60405180910390fd5b6107be610d25565b6107c6610d76565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009c939b5091995046985030975095509350915050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060917f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace009161038090611a19565b5f7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d008054909150610100900473ffffffffffffffffffffffffffffffffffffffff1633146108d7576040517f38da3b1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106ae8383610d9f565b5f3361040f818585610c0f565b8342111561092b576040517f62791302000000000000000000000000000000000000000000000000000000008152600481018590526024016107ad565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98888886109a28c73ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f610a0982610df9565b90505f610a1882878787610e40565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610a9f576040517f4b800e4600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301528b1660248201526044016107ad565b610aaa8a8a8a610ab6565b50505050505050505050565b6106ae8383836001610e6c565b610acb610fd6565b610ad5828261103f565b5050565b610ae1610fd6565b610b20816040518060400160405280600181526020017f31000000000000000000000000000000000000000000000000000000000000008152506110a2565b50565b73ffffffffffffffffffffffffffffffffffffffff8381165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610c095781811015610bfb576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064016107ad565b610c0984848484035f610e6c565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610c5e576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f60048201526024016107ad565b73ffffffffffffffffffffffffffffffffffffffff8216610cad576040517fec442f050000000000000000000000000000000000000000000000000000000081525f60048201526024016107ad565b6106ae838383611114565b5f6106256112e1565b73ffffffffffffffffffffffffffffffffffffffff8216610d10576040517fec442f050000000000000000000000000000000000000000000000000000000081525f60048201526024016107ad565b610ad55f8383611114565b5f61041582611354565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060917fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1009161038090611a19565b60605f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10061036f565b73ffffffffffffffffffffffffffffffffffffffff8216610dee576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f60048201526024016107ad565b610ad5825f83611114565b5f610415610e05610cb8565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b5f5f5f5f610e508888888861137c565b925092509250610e60828261146f565b50909695505050505050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff8516610edc576040517fe602df050000000000000000000000000000000000000000000000000000000081525f60048201526024016107ad565b73ffffffffffffffffffffffffffffffffffffffff8416610f2b576040517f94280d620000000000000000000000000000000000000000000000000000000081525f60048201526024016107ad565b73ffffffffffffffffffffffffffffffffffffffff8086165f90815260018301602090815260408083209388168352929052208390558115610fcf578373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92585604051610fc691815260200190565b60405180910390a35b5050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff1661103d576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b611047610fd6565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace007f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace036110938482611aae565b5060048101610c098382611aae565b6110aa610fd6565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1007fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1026110f68482611aae565b50600381016111058382611aae565b505f8082556001909101555050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff841661116e5781816002015f8282546111639190611bc5565b9091555061121e9050565b73ffffffffffffffffffffffffffffffffffffffff84165f90815260208290526040902054828110156111f3576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8616600482015260248101829052604481018490526064016107ad565b73ffffffffffffffffffffffffffffffffffffffff85165f9081526020839052604090209083900390555b73ffffffffffffffffffffffffffffffffffffffff8316611249576002810180548390039055611274565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526020829052604090208054830190555b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516112d391815260200190565b60405180910390a350505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f61130b611572565b6113136115ed565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f807f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006106d7565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156113b557505f91506003905082611465565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015611406573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811661145c57505f925060019150829050611465565b92505f91508190505b9450945094915050565b5f82600381111561148257611482611bfd565b0361148b575050565b600182600381111561149f5761149f611bfd565b036114d6576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028260038111156114ea576114ea611bfd565b03611524576040517ffce698f7000000000000000000000000000000000000000000000000000000008152600481018290526024016107ad565b600382600381111561153857611538611bfd565b03610ad5576040517fd78bce0c000000000000000000000000000000000000000000000000000000008152600481018290526024016107ad565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1008161159d610d25565b8051909150156115b557805160209091012092915050565b815480156115c4579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10081611618610d76565b80519091501561163057805160209091012092915050565b600182015480156115c4579392505050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f6116a06020830184611642565b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146116ca575f5ffd5b919050565b5f5f604083850312156116e0575f5ffd5b6116e9836116a7565b946020939093013593505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112611733575f5ffd5b813567ffffffffffffffff81111561174d5761174d6116f7565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff821117156117b9576117b96116f7565b6040528181528382016020018510156117d0575f5ffd5b816020850160208301375f918101602001919091529392505050565b803560ff811681146116ca575f5ffd5b5f5f5f6060848603121561180e575f5ffd5b833567ffffffffffffffff811115611824575f5ffd5b61183086828701611724565b935050602084013567ffffffffffffffff81111561184c575f5ffd5b61185886828701611724565b925050611867604085016117ec565b90509250925092565b5f5f5f60608486031215611882575f5ffd5b61188b846116a7565b9250611899602085016116a7565b929592945050506040919091013590565b5f602082840312156118ba575f5ffd5b6116a0826116a7565b7fff000000000000000000000000000000000000000000000000000000000000008816815260e060208201525f6118fd60e0830189611642565b828103604084015261190f8189611642565b6060840188905273ffffffffffffffffffffffffffffffffffffffff8716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015611971578351835260209384019390920191600101611953565b50909b9a5050505050505050505050565b5f5f5f5f5f5f5f60e0888a031215611998575f5ffd5b6119a1886116a7565b96506119af602089016116a7565b955060408801359450606088013593506119cb608089016117ec565b9699959850939692959460a0840135945060c09093013592915050565b5f5f604083850312156119f9575f5ffd5b611a02836116a7565b9150611a10602084016116a7565b90509250929050565b600181811c90821680611a2d57607f821691505b602082108103611a64577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b601f8211156106ae57805f5260205f20601f840160051c81016020851015611a8f5750805b601f840160051c820191505b81811015610fcf575f8155600101611a9b565b815167ffffffffffffffff811115611ac857611ac86116f7565b611adc81611ad68454611a19565b84611a6a565b6020601f821160018114611b2d575f8315611af75750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455610fcf565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015611b7a5787850151825560209485019460019092019101611b5a565b5084821015611bb657868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b80820180821115610415577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffdfea2646970667358221220cab1bda26a25f523842457c2ae87e0281ec40e493b88229fc95b7a72c369c0de64736f6c634300081c00336080604052348015600e575f5ffd5b50611fe58061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061006f575f3560e01c8063c00f14ab1161004d578063c00f14ab146100e4578063c514f24e146100f7578063cf825e55146100ff575f5ffd5b8063737ce34114610073578063a28fa4a31461009c578063be3dcf62146100bf575b5f5ffd5b610086610081366004610c2c565b610112565b6040516100939190610c9a565b60405180910390f35b6100af6100aa366004610cac565b610226565b6040519015158152602001610093565b6100d26100cd366004610c2c565b6107db565b60405160ff9091168152602001610093565b6100866100f2366004610c2c565b6108ca565b61008661090f565b61008661010d366004610c2c565b61092e565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f95d89b410000000000000000000000000000000000000000000000000000000017905290516060915f91829173ffffffffffffffffffffffffffffffffffffffff8616916101939190610d4f565b5f60405180830381855afa9150503d805f81146101cb576040519150601f19603f3d011682016040523d82523d5f602084013e6101d0565b606091505b509150915081610215576040518060400160405280600981526020017f4e4f5f53594d424f4c000000000000000000000000000000000000000000000081525061021e565b61021e81610a31565b949350505050565b5f806102356004828789610d65565b61023e91610d8c565b90507f2afa5331000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216016104e9575f5f5f5f5f5f5f8c8c60049080926102a393929190610d65565b8101906102b09190610e00565b96509650965096509650965096508a73ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff1614610323576040517f912ecce700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8973ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614610388576040517f750643af00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8e73ffffffffffffffffffffffffffffffffffffffff1663d505accf60e01b8989898989898960405160240161040e979695949392919073ffffffffffffffffffffffffffffffffffffffff97881681529590961660208601526040850193909352606084019190915260ff16608083015260a082015260c081019190915260e00190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009094169390931790925290516104979190610d4f565b5f604051808303815f865af19150503d805f81146104d0576040519150601f19603f3d011682016040523d82523d5f602084013e6104d5565b606091505b50909a506107d29950505050505050505050565b7f703450f4000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216016107a0575f5f5f5f5f5f5f5f8d8d600490809261054d93929190610d65565b81019061055a9190610e6c565b975097509750975097509750975097508b73ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff16146105cf576040517f912ecce700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8a73ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff1614610634576040517f750643af00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8f73ffffffffffffffffffffffffffffffffffffffff16638fcbaf0c60e01b8a8a8a8a8a8a8a8a6040516024016106c498979695949392919073ffffffffffffffffffffffffffffffffffffffff9889168152969097166020870152604086019490945260608501929092521515608084015260ff1660a083015260c082015260e08101919091526101000190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090941693909317909252905161074d9190610d4f565b5f604051808303815f865af19150503d805f8114610786576040519150601f19603f3d011682016040523d82523d5f602084013e61078b565b606091505b50909b506107d29a5050505050505050505050565b6040517fe282c0ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b95945050505050565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f313ce5670000000000000000000000000000000000000000000000000000000017905290515f918291829173ffffffffffffffffffffffffffffffffffffffff86169161085b9190610d4f565b5f60405180830381855afa9150503d805f8114610893576040519150601f19603f3d011682016040523d82523d5f602084013e610898565b606091505b50915091508180156108ab575080516020145b6108b657601261021e565b8080602001905181019061021e9190610eee565b60606108d58261092e565b6108de83610112565b6108e7846107db565b6040516020016108f993929190610f09565b6040516020818303038152906040529050919050565b60405180610f000160405280610ec881526020016110e8610ec8913981565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f06fdde030000000000000000000000000000000000000000000000000000000017905290516060915f91829173ffffffffffffffffffffffffffffffffffffffff8616916109af9190610d4f565b5f60405180830381855afa9150503d805f81146109e7576040519150601f19603f3d011682016040523d82523d5f602084013e6109ec565b606091505b509150915081610215576040518060400160405280600781526020017f4e4f5f4e414d450000000000000000000000000000000000000000000000000081525061021e565b60606040825110610a565781806020019051810190610a509190610f6e565b92915050565b8151602003610bbf575f5b602081108015610aa85750828181518110610a7e57610a7e61105e565b01602001517fff000000000000000000000000000000000000000000000000000000000000001615155b15610abf5780610ab78161108b565b915050610a61565b805f03610b0157505060408051808201909152601281527f4e4f545f56414c49445f454e434f44494e4700000000000000000000000000006020820152919050565b5f8167ffffffffffffffff811115610b1b57610b1b610f41565b6040519080825280601f01601f191660200182016040528015610b45576020820181803683370190505b5090505f5b82811015610bb757848181518110610b6457610b6461105e565b602001015160f81c60f81b828281518110610b8157610b8161105e565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600101610b4a565b509392505050565b505060408051808201909152601281527f4e4f545f56414c49445f454e434f44494e470000000000000000000000000000602082015290565b919050565b73ffffffffffffffffffffffffffffffffffffffff81168114610c1e575f5ffd5b50565b8035610bf881610bfd565b5f60208284031215610c3c575f5ffd5b8135610c4781610bfd565b9392505050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f610c476020830184610c4e565b5f5f5f5f5f60808688031215610cc0575f5ffd5b8535610ccb81610bfd565b9450602086013567ffffffffffffffff811115610ce6575f5ffd5b8601601f81018813610cf6575f5ffd5b803567ffffffffffffffff811115610d0c575f5ffd5b886020828401011115610d1d575f5ffd5b602091909101945092506040860135610d3581610bfd565b9150610d4360608701610c21565b90509295509295909350565b5f82518060208501845e5f920191825250919050565b5f5f85851115610d73575f5ffd5b83861115610d7f575f5ffd5b5050820193919092039150565b80357fffffffff000000000000000000000000000000000000000000000000000000008116906004841015610deb577fffffffff00000000000000000000000000000000000000000000000000000000808560040360031b1b82161691505b5092915050565b60ff81168114610c1e575f5ffd5b5f5f5f5f5f5f5f60e0888a031215610e16575f5ffd5b8735610e2181610bfd565b96506020880135610e3181610bfd565b955060408801359450606088013593506080880135610e4f81610df2565b9699959850939692959460a0840135945060c09093013592915050565b5f5f5f5f5f5f5f5f610100898b031215610e84575f5ffd5b8835610e8f81610bfd565b97506020890135610e9f81610bfd565b9650604089013595506060890135945060808901358015158114610ec1575f5ffd5b935060a0890135610ed181610df2565b979a969950949793969295929450505060c08201359160e0013590565b5f60208284031215610efe575f5ffd5b8151610c4781610df2565b606081525f610f1b6060830186610c4e565b8281036020840152610f2d8186610c4e565b91505060ff83166040830152949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f60208284031215610f7e575f5ffd5b815167ffffffffffffffff811115610f94575f5ffd5b8201601f81018413610fa4575f5ffd5b805167ffffffffffffffff811115610fbe57610fbe610f41565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561102a5761102a610f41565b604052818152828201602001861015611041575f5ffd5b8160208401602083015e5f91810160200191909152949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036110e0577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b506001019056fe60806040819052631d97f74d60e11b81523390633b2fee9a90608490602090600481865afa158015610033573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100579190610433565b604080515f80825260208201909252905061007382825f6100e2565b50506100dd336001600160a01b03166338b8fbbb6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100b4573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100d89190610433565b61010d565b6104c8565b6100eb8361017a565b5f825111806100f75750805b156101085761010683836101b9565b505b505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61014c5f516020610e815f395f51905f52546001600160a01b031690565b604080516001600160a01b03928316815291841660208301520160405180910390a1610177816101e5565b50565b61018381610280565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606101de8383604051806060016040528060278152602001610ea160279139610314565b9392505050565b6001600160a01b03811661024f5760405162461bcd60e51b815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b805f516020610e815f395f51905f525b80546001600160a01b0319166001600160a01b039290921691909117905550565b6001600160a01b0381163b6102ed5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610246565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc61025f565b60605f5f856001600160a01b031685604051610330919061047b565b5f60405180830381855af49150503d805f8114610368576040519150601f19603f3d011682016040523d82523d5f602084013e61036d565b606091505b50909250905061037f86838387610389565b9695505050505050565b606083156103f75782515f036103f0576001600160a01b0385163b6103f05760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610246565b5081610401565b6104018383610409565b949350505050565b8151156104195781518083602001fd5b8060405162461bcd60e51b81526004016102469190610496565b5f60208284031215610443575f5ffd5b81516001600160a01b03811681146101de575f5ffd5b5f5b8381101561047357818101518382015260200161045b565b50505f910152565b5f825161048c818460208701610459565b9190910192915050565b602081525f82518060208401526104b4816040850160208701610459565b601f01601f19169190910160400192915050565b6109ac806104d55f395ff3fe60806040526004361061005d575f3560e01c80635c60da1b116100425780635c60da1b146100a65780638f283970146100e3578063f851a440146101025761006c565b80633659cfe6146100745780634f1ef286146100935761006c565b3661006c5761006a610116565b005b61006a610116565b34801561007f575f5ffd5b5061006a61008e366004610854565b610130565b61006a6100a136600461086d565b610178565b3480156100b1575f5ffd5b506100ba6101eb565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b3480156100ee575f5ffd5b5061006a6100fd366004610854565b610228565b34801561010d575f5ffd5b506100ba610255565b61011e610282565b61012e610129610359565b610362565b565b610138610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101705761016d8160405180602001604052805f8152505f6103bf565b50565b61016d610116565b610180610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101e3576101de8383838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250600192506103bf915050565b505050565b6101de610116565b5f6101f4610380565b73ffffffffffffffffffffffffffffffffffffffff16330361021d57610218610359565b905090565b610225610116565b90565b610230610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101705761016d816103e9565b5f61025e610380565b73ffffffffffffffffffffffffffffffffffffffff16330361021d57610218610380565b61028a610380565b73ffffffffffffffffffffffffffffffffffffffff16330361012e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604260248201527f5472616e73706172656e745570677261646561626c6550726f78793a2061646d60448201527f696e2063616e6e6f742066616c6c6261636b20746f2070726f7879207461726760648201527f6574000000000000000000000000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b5f61021861044a565b365f5f375f5f365f845af43d5f5f3e80801561037c573d5ff35b3d5ffd5b5f7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b5473ffffffffffffffffffffffffffffffffffffffff16919050565b6103c883610471565b5f825111806103d45750805b156101de576103e383836104bd565b50505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f610412610380565b6040805173ffffffffffffffffffffffffffffffffffffffff928316815291841660208301520160405180910390a161016d816104e9565b5f7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6103a3565b61047a816105f5565b60405173ffffffffffffffffffffffffffffffffffffffff8216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606104e28383604051806060016040528060278152602001610979602791396106c0565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff811661058c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610350565b807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b73ffffffffffffffffffffffffffffffffffffffff81163b610699576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201527f6f74206120636f6e7472616374000000000000000000000000000000000000006064820152608401610350565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6105af565b60605f5f8573ffffffffffffffffffffffffffffffffffffffff16856040516106e9919061090d565b5f60405180830381855af49150503d805f8114610721576040519150601f19603f3d011682016040523d82523d5f602084013e610726565b606091505b509150915061073786838387610741565b9695505050505050565b606083156107d65782515f036107cf5773ffffffffffffffffffffffffffffffffffffffff85163b6107cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610350565b50816107e0565b6107e083836107e8565b949350505050565b8151156107f85781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103509190610928565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f575f5ffd5b919050565b5f60208284031215610864575f5ffd5b6104e28261082c565b5f5f5f6040848603121561087f575f5ffd5b6108888461082c565b9250602084013567ffffffffffffffff8111156108a3575f5ffd5b8401601f810186136108b3575f5ffd5b803567ffffffffffffffff8111156108c9575f5ffd5b8660208284010111156108da575f5ffd5b939660209190910195509293505050565b5f5b838110156109055781810151838201526020016108ed565b50505f910152565b5f825161091e8184602087016108eb565b9190910192915050565b602081525f82518060208401526109468160408501602087016108eb565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a164736f6c634300081c000ab53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220d609d342b5a041d079d116700c7fe69426363ead3fe1c19c15bc7d159c406f0364736f6c634300081c0033",
}

// Extensionagglayerbridgel2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Extensionagglayerbridgel2MetaData.ABI instead.
var Extensionagglayerbridgel2ABI = Extensionagglayerbridgel2MetaData.ABI

// Extensionagglayerbridgel2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Extensionagglayerbridgel2MetaData.Bin instead.
var Extensionagglayerbridgel2Bin = Extensionagglayerbridgel2MetaData.Bin

// DeployExtensionagglayerbridgel2 deploys a new Ethereum contract, binding an instance of Extensionagglayerbridgel2 to it.
func DeployExtensionagglayerbridgel2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Extensionagglayerbridgel2, error) {
	parsed, err := Extensionagglayerbridgel2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Extensionagglayerbridgel2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Extensionagglayerbridgel2{Extensionagglayerbridgel2Caller: Extensionagglayerbridgel2Caller{contract: contract}, Extensionagglayerbridgel2Transactor: Extensionagglayerbridgel2Transactor{contract: contract}, Extensionagglayerbridgel2Filterer: Extensionagglayerbridgel2Filterer{contract: contract}}, nil
}

// Extensionagglayerbridgel2 is an auto generated Go binding around an Ethereum contract.
type Extensionagglayerbridgel2 struct {
	Extensionagglayerbridgel2Caller     // Read-only binding to the contract
	Extensionagglayerbridgel2Transactor // Write-only binding to the contract
	Extensionagglayerbridgel2Filterer   // Log filterer for contract events
}

// Extensionagglayerbridgel2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Extensionagglayerbridgel2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Extensionagglayerbridgel2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Extensionagglayerbridgel2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Extensionagglayerbridgel2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Extensionagglayerbridgel2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Extensionagglayerbridgel2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Extensionagglayerbridgel2Session struct {
	Contract     *Extensionagglayerbridgel2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// Extensionagglayerbridgel2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Extensionagglayerbridgel2CallerSession struct {
	Contract *Extensionagglayerbridgel2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// Extensionagglayerbridgel2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Extensionagglayerbridgel2TransactorSession struct {
	Contract     *Extensionagglayerbridgel2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// Extensionagglayerbridgel2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Extensionagglayerbridgel2Raw struct {
	Contract *Extensionagglayerbridgel2 // Generic contract binding to access the raw methods on
}

// Extensionagglayerbridgel2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Extensionagglayerbridgel2CallerRaw struct {
	Contract *Extensionagglayerbridgel2Caller // Generic read-only contract binding to access the raw methods on
}

// Extensionagglayerbridgel2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Extensionagglayerbridgel2TransactorRaw struct {
	Contract *Extensionagglayerbridgel2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewExtensionagglayerbridgel2 creates a new instance of Extensionagglayerbridgel2, bound to a specific deployed contract.
func NewExtensionagglayerbridgel2(address common.Address, backend bind.ContractBackend) (*Extensionagglayerbridgel2, error) {
	contract, err := bindExtensionagglayerbridgel2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2{Extensionagglayerbridgel2Caller: Extensionagglayerbridgel2Caller{contract: contract}, Extensionagglayerbridgel2Transactor: Extensionagglayerbridgel2Transactor{contract: contract}, Extensionagglayerbridgel2Filterer: Extensionagglayerbridgel2Filterer{contract: contract}}, nil
}

// NewExtensionagglayerbridgel2Caller creates a new read-only instance of Extensionagglayerbridgel2, bound to a specific deployed contract.
func NewExtensionagglayerbridgel2Caller(address common.Address, caller bind.ContractCaller) (*Extensionagglayerbridgel2Caller, error) {
	contract, err := bindExtensionagglayerbridgel2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2Caller{contract: contract}, nil
}

// NewExtensionagglayerbridgel2Transactor creates a new write-only instance of Extensionagglayerbridgel2, bound to a specific deployed contract.
func NewExtensionagglayerbridgel2Transactor(address common.Address, transactor bind.ContractTransactor) (*Extensionagglayerbridgel2Transactor, error) {
	contract, err := bindExtensionagglayerbridgel2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2Transactor{contract: contract}, nil
}

// NewExtensionagglayerbridgel2Filterer creates a new log filterer instance of Extensionagglayerbridgel2, bound to a specific deployed contract.
func NewExtensionagglayerbridgel2Filterer(address common.Address, filterer bind.ContractFilterer) (*Extensionagglayerbridgel2Filterer, error) {
	contract, err := bindExtensionagglayerbridgel2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2Filterer{contract: contract}, nil
}

// bindExtensionagglayerbridgel2 binds a generic wrapper to an already deployed contract.
func bindExtensionagglayerbridgel2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Extensionagglayerbridgel2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Extensionagglayerbridgel2.Contract.Extensionagglayerbridgel2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Extensionagglayerbridgel2.Contract.Extensionagglayerbridgel2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Extensionagglayerbridgel2.Contract.Extensionagglayerbridgel2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Extensionagglayerbridgel2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Extensionagglayerbridgel2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Extensionagglayerbridgel2.Contract.contract.Transact(opts, method, params...)
}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() pure returns(bytes)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) INITBYTECODETRANSPARENTPROXY(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "INIT_BYTECODE_TRANSPARENT_PROXY")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() pure returns(bytes)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) INITBYTECODETRANSPARENTPROXY() ([]byte, error) {
	return _Extensionagglayerbridgel2.Contract.INITBYTECODETRANSPARENTPROXY(&_Extensionagglayerbridgel2.CallOpts)
}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() pure returns(bytes)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) INITBYTECODETRANSPARENTPROXY() ([]byte, error) {
	return _Extensionagglayerbridgel2.Contract.INITBYTECODETRANSPARENTPROXY(&_Extensionagglayerbridgel2.CallOpts)
}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) WETHToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "WETHToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) WETHToken() (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.WETHToken(&_Extensionagglayerbridgel2.CallOpts)
}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) WETHToken() (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.WETHToken(&_Extensionagglayerbridgel2.CallOpts)
}

// AcceptEmergencyBridgePauserRole is a free data retrieval call binding the contract method 0xe88f0436.
//
// Solidity: function acceptEmergencyBridgePauserRole() pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) AcceptEmergencyBridgePauserRole(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "acceptEmergencyBridgePauserRole")

	if err != nil {
		return err
	}

	return err

}

// AcceptEmergencyBridgePauserRole is a free data retrieval call binding the contract method 0xe88f0436.
//
// Solidity: function acceptEmergencyBridgePauserRole() pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) AcceptEmergencyBridgePauserRole() error {
	return _Extensionagglayerbridgel2.Contract.AcceptEmergencyBridgePauserRole(&_Extensionagglayerbridgel2.CallOpts)
}

// AcceptEmergencyBridgePauserRole is a free data retrieval call binding the contract method 0xe88f0436.
//
// Solidity: function acceptEmergencyBridgePauserRole() pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) AcceptEmergencyBridgePauserRole() error {
	return _Extensionagglayerbridgel2.Contract.AcceptEmergencyBridgePauserRole(&_Extensionagglayerbridgel2.CallOpts)
}

// AcceptEmergencyBridgeUnpauserRole is a free data retrieval call binding the contract method 0x8b37b873.
//
// Solidity: function acceptEmergencyBridgeUnpauserRole() pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) AcceptEmergencyBridgeUnpauserRole(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "acceptEmergencyBridgeUnpauserRole")

	if err != nil {
		return err
	}

	return err

}

// AcceptEmergencyBridgeUnpauserRole is a free data retrieval call binding the contract method 0x8b37b873.
//
// Solidity: function acceptEmergencyBridgeUnpauserRole() pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) AcceptEmergencyBridgeUnpauserRole() error {
	return _Extensionagglayerbridgel2.Contract.AcceptEmergencyBridgeUnpauserRole(&_Extensionagglayerbridgel2.CallOpts)
}

// AcceptEmergencyBridgeUnpauserRole is a free data retrieval call binding the contract method 0x8b37b873.
//
// Solidity: function acceptEmergencyBridgeUnpauserRole() pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) AcceptEmergencyBridgeUnpauserRole() error {
	return _Extensionagglayerbridgel2.Contract.AcceptEmergencyBridgeUnpauserRole(&_Extensionagglayerbridgel2.CallOpts)
}

// AcceptProxiedTokensManagerRole is a free data retrieval call binding the contract method 0x8c668f1c.
//
// Solidity: function acceptProxiedTokensManagerRole() pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) AcceptProxiedTokensManagerRole(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "acceptProxiedTokensManagerRole")

	if err != nil {
		return err
	}

	return err

}

// AcceptProxiedTokensManagerRole is a free data retrieval call binding the contract method 0x8c668f1c.
//
// Solidity: function acceptProxiedTokensManagerRole() pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) AcceptProxiedTokensManagerRole() error {
	return _Extensionagglayerbridgel2.Contract.AcceptProxiedTokensManagerRole(&_Extensionagglayerbridgel2.CallOpts)
}

// AcceptProxiedTokensManagerRole is a free data retrieval call binding the contract method 0x8c668f1c.
//
// Solidity: function acceptProxiedTokensManagerRole() pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) AcceptProxiedTokensManagerRole() error {
	return _Extensionagglayerbridgel2.Contract.AcceptProxiedTokensManagerRole(&_Extensionagglayerbridgel2.CallOpts)
}

// ActivateEmergencyState is a free data retrieval call binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) ActivateEmergencyState(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "activateEmergencyState")

	if err != nil {
		return err
	}

	return err

}

// ActivateEmergencyState is a free data retrieval call binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) ActivateEmergencyState() error {
	return _Extensionagglayerbridgel2.Contract.ActivateEmergencyState(&_Extensionagglayerbridgel2.CallOpts)
}

// ActivateEmergencyState is a free data retrieval call binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) ActivateEmergencyState() error {
	return _Extensionagglayerbridgel2.Contract.ActivateEmergencyState(&_Extensionagglayerbridgel2.CallOpts)
}

// BackwardLET is a free data retrieval call binding the contract method 0xfd7640e8.
//
// Solidity: function backwardLET(uint256 , bytes32[32] , bytes32 , bytes32[32] ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) BackwardLET(opts *bind.CallOpts, arg0 *big.Int, arg1 [32][32]byte, arg2 [32]byte, arg3 [32][32]byte) error {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "backwardLET", arg0, arg1, arg2, arg3)

	if err != nil {
		return err
	}

	return err

}

// BackwardLET is a free data retrieval call binding the contract method 0xfd7640e8.
//
// Solidity: function backwardLET(uint256 , bytes32[32] , bytes32 , bytes32[32] ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) BackwardLET(arg0 *big.Int, arg1 [32][32]byte, arg2 [32]byte, arg3 [32][32]byte) error {
	return _Extensionagglayerbridgel2.Contract.BackwardLET(&_Extensionagglayerbridgel2.CallOpts, arg0, arg1, arg2, arg3)
}

// BackwardLET is a free data retrieval call binding the contract method 0xfd7640e8.
//
// Solidity: function backwardLET(uint256 , bytes32[32] , bytes32 , bytes32[32] ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) BackwardLET(arg0 *big.Int, arg1 [32][32]byte, arg2 [32]byte, arg3 [32][32]byte) error {
	return _Extensionagglayerbridgel2.Contract.BackwardLET(&_Extensionagglayerbridgel2.CallOpts, arg0, arg1, arg2, arg3)
}

// BridgeLib is a free data retrieval call binding the contract method 0x6f0bc3da.
//
// Solidity: function bridgeLib() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) BridgeLib(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "bridgeLib")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeLib is a free data retrieval call binding the contract method 0x6f0bc3da.
//
// Solidity: function bridgeLib() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) BridgeLib() (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.BridgeLib(&_Extensionagglayerbridgel2.CallOpts)
}

// BridgeLib is a free data retrieval call binding the contract method 0x6f0bc3da.
//
// Solidity: function bridgeLib() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) BridgeLib() (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.BridgeLib(&_Extensionagglayerbridgel2.CallOpts)
}

// BridgeManager is a free data retrieval call binding the contract method 0x14cc01a0.
//
// Solidity: function bridgeManager() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) BridgeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "bridgeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeManager is a free data retrieval call binding the contract method 0x14cc01a0.
//
// Solidity: function bridgeManager() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) BridgeManager() (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.BridgeManager(&_Extensionagglayerbridgel2.CallOpts)
}

// BridgeManager is a free data retrieval call binding the contract method 0x14cc01a0.
//
// Solidity: function bridgeManager() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) BridgeManager() (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.BridgeManager(&_Extensionagglayerbridgel2.CallOpts)
}

// BridgeMessageWETH is a free data retrieval call binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 , address , uint256 , bool , bytes ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) BridgeMessageWETH(opts *bind.CallOpts, arg0 uint32, arg1 common.Address, arg2 *big.Int, arg3 bool, arg4 []byte) error {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "bridgeMessageWETH", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return err
	}

	return err

}

// BridgeMessageWETH is a free data retrieval call binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 , address , uint256 , bool , bytes ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) BridgeMessageWETH(arg0 uint32, arg1 common.Address, arg2 *big.Int, arg3 bool, arg4 []byte) error {
	return _Extensionagglayerbridgel2.Contract.BridgeMessageWETH(&_Extensionagglayerbridgel2.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// BridgeMessageWETH is a free data retrieval call binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 , address , uint256 , bool , bytes ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) BridgeMessageWETH(arg0 uint32, arg1 common.Address, arg2 *big.Int, arg3 bool, arg4 []byte) error {
	return _Extensionagglayerbridgel2.Contract.BridgeMessageWETH(&_Extensionagglayerbridgel2.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// ClaimAsset is a free data retrieval call binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] , bytes32[32] , uint256 , bytes32 , bytes32 , uint32 , address , uint32 , address , uint256 , bytes ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) ClaimAsset(opts *bind.CallOpts, arg0 [32][32]byte, arg1 [32][32]byte, arg2 *big.Int, arg3 [32]byte, arg4 [32]byte, arg5 uint32, arg6 common.Address, arg7 uint32, arg8 common.Address, arg9 *big.Int, arg10 []byte) error {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "claimAsset", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)

	if err != nil {
		return err
	}

	return err

}

// ClaimAsset is a free data retrieval call binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] , bytes32[32] , uint256 , bytes32 , bytes32 , uint32 , address , uint32 , address , uint256 , bytes ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) ClaimAsset(arg0 [32][32]byte, arg1 [32][32]byte, arg2 *big.Int, arg3 [32]byte, arg4 [32]byte, arg5 uint32, arg6 common.Address, arg7 uint32, arg8 common.Address, arg9 *big.Int, arg10 []byte) error {
	return _Extensionagglayerbridgel2.Contract.ClaimAsset(&_Extensionagglayerbridgel2.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
}

// ClaimAsset is a free data retrieval call binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] , bytes32[32] , uint256 , bytes32 , bytes32 , uint32 , address , uint32 , address , uint256 , bytes ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) ClaimAsset(arg0 [32][32]byte, arg1 [32][32]byte, arg2 *big.Int, arg3 [32]byte, arg4 [32]byte, arg5 uint32, arg6 common.Address, arg7 uint32, arg8 common.Address, arg9 *big.Int, arg10 []byte) error {
	return _Extensionagglayerbridgel2.Contract.ClaimAsset(&_Extensionagglayerbridgel2.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
}

// ClaimMessage is a free data retrieval call binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] , bytes32[32] , uint256 , bytes32 , bytes32 , uint32 , address , uint32 , address , uint256 , bytes ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) ClaimMessage(opts *bind.CallOpts, arg0 [32][32]byte, arg1 [32][32]byte, arg2 *big.Int, arg3 [32]byte, arg4 [32]byte, arg5 uint32, arg6 common.Address, arg7 uint32, arg8 common.Address, arg9 *big.Int, arg10 []byte) error {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "claimMessage", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)

	if err != nil {
		return err
	}

	return err

}

// ClaimMessage is a free data retrieval call binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] , bytes32[32] , uint256 , bytes32 , bytes32 , uint32 , address , uint32 , address , uint256 , bytes ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) ClaimMessage(arg0 [32][32]byte, arg1 [32][32]byte, arg2 *big.Int, arg3 [32]byte, arg4 [32]byte, arg5 uint32, arg6 common.Address, arg7 uint32, arg8 common.Address, arg9 *big.Int, arg10 []byte) error {
	return _Extensionagglayerbridgel2.Contract.ClaimMessage(&_Extensionagglayerbridgel2.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
}

// ClaimMessage is a free data retrieval call binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] , bytes32[32] , uint256 , bytes32 , bytes32 , uint32 , address , uint32 , address , uint256 , bytes ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) ClaimMessage(arg0 [32][32]byte, arg1 [32][32]byte, arg2 *big.Int, arg3 [32]byte, arg4 [32]byte, arg5 uint32, arg6 common.Address, arg7 uint32, arg8 common.Address, arg9 *big.Int, arg10 []byte) error {
	return _Extensionagglayerbridgel2.Contract.ClaimMessage(&_Extensionagglayerbridgel2.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) ClaimedBitMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "claimedBitMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) ClaimedBitMap(arg0 *big.Int) (*big.Int, error) {
	return _Extensionagglayerbridgel2.Contract.ClaimedBitMap(&_Extensionagglayerbridgel2.CallOpts, arg0)
}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) ClaimedBitMap(arg0 *big.Int) (*big.Int, error) {
	return _Extensionagglayerbridgel2.Contract.ClaimedBitMap(&_Extensionagglayerbridgel2.CallOpts, arg0)
}

// ClaimedGlobalIndexHashChain is a free data retrieval call binding the contract method 0x1d081d8c.
//
// Solidity: function claimedGlobalIndexHashChain() view returns(bytes32)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) ClaimedGlobalIndexHashChain(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "claimedGlobalIndexHashChain")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ClaimedGlobalIndexHashChain is a free data retrieval call binding the contract method 0x1d081d8c.
//
// Solidity: function claimedGlobalIndexHashChain() view returns(bytes32)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) ClaimedGlobalIndexHashChain() ([32]byte, error) {
	return _Extensionagglayerbridgel2.Contract.ClaimedGlobalIndexHashChain(&_Extensionagglayerbridgel2.CallOpts)
}

// ClaimedGlobalIndexHashChain is a free data retrieval call binding the contract method 0x1d081d8c.
//
// Solidity: function claimedGlobalIndexHashChain() view returns(bytes32)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) ClaimedGlobalIndexHashChain() ([32]byte, error) {
	return _Extensionagglayerbridgel2.Contract.ClaimedGlobalIndexHashChain(&_Extensionagglayerbridgel2.CallOpts)
}

// ComputeTokenProxyAddress is a free data retrieval call binding the contract method 0xf214e161.
//
// Solidity: function computeTokenProxyAddress(uint32 , address ) pure returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) ComputeTokenProxyAddress(opts *bind.CallOpts, arg0 uint32, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "computeTokenProxyAddress", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeTokenProxyAddress is a free data retrieval call binding the contract method 0xf214e161.
//
// Solidity: function computeTokenProxyAddress(uint32 , address ) pure returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) ComputeTokenProxyAddress(arg0 uint32, arg1 common.Address) (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.ComputeTokenProxyAddress(&_Extensionagglayerbridgel2.CallOpts, arg0, arg1)
}

// ComputeTokenProxyAddress is a free data retrieval call binding the contract method 0xf214e161.
//
// Solidity: function computeTokenProxyAddress(uint32 , address ) pure returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) ComputeTokenProxyAddress(arg0 uint32, arg1 common.Address) (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.ComputeTokenProxyAddress(&_Extensionagglayerbridgel2.CallOpts, arg0, arg1)
}

// DeactivateEmergencyState is a free data retrieval call binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) DeactivateEmergencyState(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "deactivateEmergencyState")

	if err != nil {
		return err
	}

	return err

}

// DeactivateEmergencyState is a free data retrieval call binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) DeactivateEmergencyState() error {
	return _Extensionagglayerbridgel2.Contract.DeactivateEmergencyState(&_Extensionagglayerbridgel2.CallOpts)
}

// DeactivateEmergencyState is a free data retrieval call binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) DeactivateEmergencyState() error {
	return _Extensionagglayerbridgel2.Contract.DeactivateEmergencyState(&_Extensionagglayerbridgel2.CallOpts)
}

// DeployWrappedTokenAndRemap is a free data retrieval call binding the contract method 0x6e974cd4.
//
// Solidity: function deployWrappedTokenAndRemap(uint32 , address , bool ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) DeployWrappedTokenAndRemap(opts *bind.CallOpts, arg0 uint32, arg1 common.Address, arg2 bool) error {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "deployWrappedTokenAndRemap", arg0, arg1, arg2)

	if err != nil {
		return err
	}

	return err

}

// DeployWrappedTokenAndRemap is a free data retrieval call binding the contract method 0x6e974cd4.
//
// Solidity: function deployWrappedTokenAndRemap(uint32 , address , bool ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) DeployWrappedTokenAndRemap(arg0 uint32, arg1 common.Address, arg2 bool) error {
	return _Extensionagglayerbridgel2.Contract.DeployWrappedTokenAndRemap(&_Extensionagglayerbridgel2.CallOpts, arg0, arg1, arg2)
}

// DeployWrappedTokenAndRemap is a free data retrieval call binding the contract method 0x6e974cd4.
//
// Solidity: function deployWrappedTokenAndRemap(uint32 , address , bool ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) DeployWrappedTokenAndRemap(arg0 uint32, arg1 common.Address, arg2 bool) error {
	return _Extensionagglayerbridgel2.Contract.DeployWrappedTokenAndRemap(&_Extensionagglayerbridgel2.CallOpts, arg0, arg1, arg2)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) DepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "depositCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) DepositCount() (*big.Int, error) {
	return _Extensionagglayerbridgel2.Contract.DepositCount(&_Extensionagglayerbridgel2.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) DepositCount() (*big.Int, error) {
	return _Extensionagglayerbridgel2.Contract.DepositCount(&_Extensionagglayerbridgel2.CallOpts)
}

// EmergencyBridgePauser is a free data retrieval call binding the contract method 0x2f84c690.
//
// Solidity: function emergencyBridgePauser() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) EmergencyBridgePauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "emergencyBridgePauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmergencyBridgePauser is a free data retrieval call binding the contract method 0x2f84c690.
//
// Solidity: function emergencyBridgePauser() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) EmergencyBridgePauser() (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.EmergencyBridgePauser(&_Extensionagglayerbridgel2.CallOpts)
}

// EmergencyBridgePauser is a free data retrieval call binding the contract method 0x2f84c690.
//
// Solidity: function emergencyBridgePauser() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) EmergencyBridgePauser() (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.EmergencyBridgePauser(&_Extensionagglayerbridgel2.CallOpts)
}

// EmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0xae24490a.
//
// Solidity: function emergencyBridgeUnpauser() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) EmergencyBridgeUnpauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "emergencyBridgeUnpauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0xae24490a.
//
// Solidity: function emergencyBridgeUnpauser() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) EmergencyBridgeUnpauser() (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.EmergencyBridgeUnpauser(&_Extensionagglayerbridgel2.CallOpts)
}

// EmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0xae24490a.
//
// Solidity: function emergencyBridgeUnpauser() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) EmergencyBridgeUnpauser() (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.EmergencyBridgeUnpauser(&_Extensionagglayerbridgel2.CallOpts)
}

// ForceEmitDetailedClaimEvent is a free data retrieval call binding the contract method 0x567f13e4.
//
// Solidity: function forceEmitDetailedClaimEvent((bytes32[32],bytes32[32],uint256,bytes32,bytes32,uint8,uint32,address,uint32,address,uint256,bytes)[] ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) ForceEmitDetailedClaimEvent(opts *bind.CallOpts, arg0 []AgglayerBridgeL2ClaimData) error {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "forceEmitDetailedClaimEvent", arg0)

	if err != nil {
		return err
	}

	return err

}

// ForceEmitDetailedClaimEvent is a free data retrieval call binding the contract method 0x567f13e4.
//
// Solidity: function forceEmitDetailedClaimEvent((bytes32[32],bytes32[32],uint256,bytes32,bytes32,uint8,uint32,address,uint32,address,uint256,bytes)[] ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) ForceEmitDetailedClaimEvent(arg0 []AgglayerBridgeL2ClaimData) error {
	return _Extensionagglayerbridgel2.Contract.ForceEmitDetailedClaimEvent(&_Extensionagglayerbridgel2.CallOpts, arg0)
}

// ForceEmitDetailedClaimEvent is a free data retrieval call binding the contract method 0x567f13e4.
//
// Solidity: function forceEmitDetailedClaimEvent((bytes32[32],bytes32[32],uint256,bytes32,bytes32,uint8,uint32,address,uint32,address,uint256,bytes)[] ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) ForceEmitDetailedClaimEvent(arg0 []AgglayerBridgeL2ClaimData) error {
	return _Extensionagglayerbridgel2.Contract.ForceEmitDetailedClaimEvent(&_Extensionagglayerbridgel2.CallOpts, arg0)
}

// ForwardLET is a free data retrieval call binding the contract method 0xf0e70808.
//
// Solidity: function forwardLET((uint8,uint32,address,uint32,address,uint256,bytes)[] , bytes32 ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) ForwardLET(opts *bind.CallOpts, arg0 []AgglayerBridgeL2LeafData, arg1 [32]byte) error {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "forwardLET", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// ForwardLET is a free data retrieval call binding the contract method 0xf0e70808.
//
// Solidity: function forwardLET((uint8,uint32,address,uint32,address,uint256,bytes)[] , bytes32 ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) ForwardLET(arg0 []AgglayerBridgeL2LeafData, arg1 [32]byte) error {
	return _Extensionagglayerbridgel2.Contract.ForwardLET(&_Extensionagglayerbridgel2.CallOpts, arg0, arg1)
}

// ForwardLET is a free data retrieval call binding the contract method 0xf0e70808.
//
// Solidity: function forwardLET((uint8,uint32,address,uint32,address,uint256,bytes)[] , bytes32 ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) ForwardLET(arg0 []AgglayerBridgeL2LeafData, arg1 [32]byte) error {
	return _Extensionagglayerbridgel2.Contract.ForwardLET(&_Extensionagglayerbridgel2.CallOpts, arg0, arg1)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) GasTokenAddress() (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.GasTokenAddress(&_Extensionagglayerbridgel2.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) GasTokenAddress() (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.GasTokenAddress(&_Extensionagglayerbridgel2.CallOpts)
}

// GasTokenMetadata is a free data retrieval call binding the contract method 0x27aef4e8.
//
// Solidity: function gasTokenMetadata() view returns(bytes)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) GasTokenMetadata(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "gasTokenMetadata")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GasTokenMetadata is a free data retrieval call binding the contract method 0x27aef4e8.
//
// Solidity: function gasTokenMetadata() view returns(bytes)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) GasTokenMetadata() ([]byte, error) {
	return _Extensionagglayerbridgel2.Contract.GasTokenMetadata(&_Extensionagglayerbridgel2.CallOpts)
}

// GasTokenMetadata is a free data retrieval call binding the contract method 0x27aef4e8.
//
// Solidity: function gasTokenMetadata() view returns(bytes)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) GasTokenMetadata() ([]byte, error) {
	return _Extensionagglayerbridgel2.Contract.GasTokenMetadata(&_Extensionagglayerbridgel2.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) GasTokenNetwork() (uint32, error) {
	return _Extensionagglayerbridgel2.Contract.GasTokenNetwork(&_Extensionagglayerbridgel2.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) GasTokenNetwork() (uint32, error) {
	return _Extensionagglayerbridgel2.Contract.GasTokenNetwork(&_Extensionagglayerbridgel2.CallOpts)
}

// GetProxiedTokensManager is a free data retrieval call binding the contract method 0x38b8fbbb.
//
// Solidity: function getProxiedTokensManager() pure returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) GetProxiedTokensManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "getProxiedTokensManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxiedTokensManager is a free data retrieval call binding the contract method 0x38b8fbbb.
//
// Solidity: function getProxiedTokensManager() pure returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) GetProxiedTokensManager() (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.GetProxiedTokensManager(&_Extensionagglayerbridgel2.CallOpts)
}

// GetProxiedTokensManager is a free data retrieval call binding the contract method 0x38b8fbbb.
//
// Solidity: function getProxiedTokensManager() pure returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) GetProxiedTokensManager() (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.GetProxiedTokensManager(&_Extensionagglayerbridgel2.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() pure returns(bytes32)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() pure returns(bytes32)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) GetRoot() ([32]byte, error) {
	return _Extensionagglayerbridgel2.Contract.GetRoot(&_Extensionagglayerbridgel2.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() pure returns(bytes32)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) GetRoot() ([32]byte, error) {
	return _Extensionagglayerbridgel2.Contract.GetRoot(&_Extensionagglayerbridgel2.CallOpts)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address ) pure returns(bytes)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) GetTokenMetadata(opts *bind.CallOpts, arg0 common.Address) ([]byte, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "getTokenMetadata", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address ) pure returns(bytes)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) GetTokenMetadata(arg0 common.Address) ([]byte, error) {
	return _Extensionagglayerbridgel2.Contract.GetTokenMetadata(&_Extensionagglayerbridgel2.CallOpts, arg0)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address ) pure returns(bytes)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) GetTokenMetadata(arg0 common.Address) ([]byte, error) {
	return _Extensionagglayerbridgel2.Contract.GetTokenMetadata(&_Extensionagglayerbridgel2.CallOpts, arg0)
}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 , address ) pure returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) GetTokenWrappedAddress(opts *bind.CallOpts, arg0 uint32, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "getTokenWrappedAddress", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 , address ) pure returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) GetTokenWrappedAddress(arg0 uint32, arg1 common.Address) (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.GetTokenWrappedAddress(&_Extensionagglayerbridgel2.CallOpts, arg0, arg1)
}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 , address ) pure returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) GetTokenWrappedAddress(arg0 uint32, arg1 common.Address) (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.GetTokenWrappedAddress(&_Extensionagglayerbridgel2.CallOpts, arg0, arg1)
}

// GetWrappedTokenBridgeImplementation is a free data retrieval call binding the contract method 0x3b2fee9a.
//
// Solidity: function getWrappedTokenBridgeImplementation() pure returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) GetWrappedTokenBridgeImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "getWrappedTokenBridgeImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWrappedTokenBridgeImplementation is a free data retrieval call binding the contract method 0x3b2fee9a.
//
// Solidity: function getWrappedTokenBridgeImplementation() pure returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) GetWrappedTokenBridgeImplementation() (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.GetWrappedTokenBridgeImplementation(&_Extensionagglayerbridgel2.CallOpts)
}

// GetWrappedTokenBridgeImplementation is a free data retrieval call binding the contract method 0x3b2fee9a.
//
// Solidity: function getWrappedTokenBridgeImplementation() pure returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) GetWrappedTokenBridgeImplementation() (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.GetWrappedTokenBridgeImplementation(&_Extensionagglayerbridgel2.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) GlobalExitRootManager() (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.GlobalExitRootManager(&_Extensionagglayerbridgel2.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.GlobalExitRootManager(&_Extensionagglayerbridgel2.CallOpts)
}

// IsClaimed is a free data retrieval call binding the contract method 0xcc461632.
//
// Solidity: function isClaimed(uint32 , uint32 ) pure returns(bool)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) IsClaimed(opts *bind.CallOpts, arg0 uint32, arg1 uint32) (bool, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "isClaimed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClaimed is a free data retrieval call binding the contract method 0xcc461632.
//
// Solidity: function isClaimed(uint32 , uint32 ) pure returns(bool)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) IsClaimed(arg0 uint32, arg1 uint32) (bool, error) {
	return _Extensionagglayerbridgel2.Contract.IsClaimed(&_Extensionagglayerbridgel2.CallOpts, arg0, arg1)
}

// IsClaimed is a free data retrieval call binding the contract method 0xcc461632.
//
// Solidity: function isClaimed(uint32 , uint32 ) pure returns(bool)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) IsClaimed(arg0 uint32, arg1 uint32) (bool, error) {
	return _Extensionagglayerbridgel2.Contract.IsClaimed(&_Extensionagglayerbridgel2.CallOpts, arg0, arg1)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) IsEmergencyState() (bool, error) {
	return _Extensionagglayerbridgel2.Contract.IsEmergencyState(&_Extensionagglayerbridgel2.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) IsEmergencyState() (bool, error) {
	return _Extensionagglayerbridgel2.Contract.IsEmergencyState(&_Extensionagglayerbridgel2.CallOpts)
}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) LastUpdatedDepositCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "lastUpdatedDepositCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) LastUpdatedDepositCount() (uint32, error) {
	return _Extensionagglayerbridgel2.Contract.LastUpdatedDepositCount(&_Extensionagglayerbridgel2.CallOpts)
}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) LastUpdatedDepositCount() (uint32, error) {
	return _Extensionagglayerbridgel2.Contract.LastUpdatedDepositCount(&_Extensionagglayerbridgel2.CallOpts)
}

// LocalBalanceTree is a free data retrieval call binding the contract method 0xd9cb3aec.
//
// Solidity: function localBalanceTree(bytes32 tokenInfoHash) view returns(uint256 amount)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) LocalBalanceTree(opts *bind.CallOpts, tokenInfoHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "localBalanceTree", tokenInfoHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LocalBalanceTree is a free data retrieval call binding the contract method 0xd9cb3aec.
//
// Solidity: function localBalanceTree(bytes32 tokenInfoHash) view returns(uint256 amount)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) LocalBalanceTree(tokenInfoHash [32]byte) (*big.Int, error) {
	return _Extensionagglayerbridgel2.Contract.LocalBalanceTree(&_Extensionagglayerbridgel2.CallOpts, tokenInfoHash)
}

// LocalBalanceTree is a free data retrieval call binding the contract method 0xd9cb3aec.
//
// Solidity: function localBalanceTree(bytes32 tokenInfoHash) view returns(uint256 amount)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) LocalBalanceTree(tokenInfoHash [32]byte) (*big.Int, error) {
	return _Extensionagglayerbridgel2.Contract.LocalBalanceTree(&_Extensionagglayerbridgel2.CallOpts, tokenInfoHash)
}

// MigrateLegacyToken is a free data retrieval call binding the contract method 0xb0b37920.
//
// Solidity: function migrateLegacyToken(address , uint256 , bytes ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) MigrateLegacyToken(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 []byte) error {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "migrateLegacyToken", arg0, arg1, arg2)

	if err != nil {
		return err
	}

	return err

}

// MigrateLegacyToken is a free data retrieval call binding the contract method 0xb0b37920.
//
// Solidity: function migrateLegacyToken(address , uint256 , bytes ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) MigrateLegacyToken(arg0 common.Address, arg1 *big.Int, arg2 []byte) error {
	return _Extensionagglayerbridgel2.Contract.MigrateLegacyToken(&_Extensionagglayerbridgel2.CallOpts, arg0, arg1, arg2)
}

// MigrateLegacyToken is a free data retrieval call binding the contract method 0xb0b37920.
//
// Solidity: function migrateLegacyToken(address , uint256 , bytes ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) MigrateLegacyToken(arg0 common.Address, arg1 *big.Int, arg2 []byte) error {
	return _Extensionagglayerbridgel2.Contract.MigrateLegacyToken(&_Extensionagglayerbridgel2.CallOpts, arg0, arg1, arg2)
}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) NetworkID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "networkID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) NetworkID() (uint32, error) {
	return _Extensionagglayerbridgel2.Contract.NetworkID(&_Extensionagglayerbridgel2.CallOpts)
}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) NetworkID() (uint32, error) {
	return _Extensionagglayerbridgel2.Contract.NetworkID(&_Extensionagglayerbridgel2.CallOpts)
}

// PendingEmergencyBridgePauser is a free data retrieval call binding the contract method 0x03e6e116.
//
// Solidity: function pendingEmergencyBridgePauser() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) PendingEmergencyBridgePauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "pendingEmergencyBridgePauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingEmergencyBridgePauser is a free data retrieval call binding the contract method 0x03e6e116.
//
// Solidity: function pendingEmergencyBridgePauser() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) PendingEmergencyBridgePauser() (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.PendingEmergencyBridgePauser(&_Extensionagglayerbridgel2.CallOpts)
}

// PendingEmergencyBridgePauser is a free data retrieval call binding the contract method 0x03e6e116.
//
// Solidity: function pendingEmergencyBridgePauser() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) PendingEmergencyBridgePauser() (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.PendingEmergencyBridgePauser(&_Extensionagglayerbridgel2.CallOpts)
}

// PendingEmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0x606617ff.
//
// Solidity: function pendingEmergencyBridgeUnpauser() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) PendingEmergencyBridgeUnpauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "pendingEmergencyBridgeUnpauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingEmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0x606617ff.
//
// Solidity: function pendingEmergencyBridgeUnpauser() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) PendingEmergencyBridgeUnpauser() (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.PendingEmergencyBridgeUnpauser(&_Extensionagglayerbridgel2.CallOpts)
}

// PendingEmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0x606617ff.
//
// Solidity: function pendingEmergencyBridgeUnpauser() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) PendingEmergencyBridgeUnpauser() (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.PendingEmergencyBridgeUnpauser(&_Extensionagglayerbridgel2.CallOpts)
}

// PendingProxiedTokensManager is a free data retrieval call binding the contract method 0xece93c6f.
//
// Solidity: function pendingProxiedTokensManager() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) PendingProxiedTokensManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "pendingProxiedTokensManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingProxiedTokensManager is a free data retrieval call binding the contract method 0xece93c6f.
//
// Solidity: function pendingProxiedTokensManager() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) PendingProxiedTokensManager() (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.PendingProxiedTokensManager(&_Extensionagglayerbridgel2.CallOpts)
}

// PendingProxiedTokensManager is a free data retrieval call binding the contract method 0xece93c6f.
//
// Solidity: function pendingProxiedTokensManager() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) PendingProxiedTokensManager() (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.PendingProxiedTokensManager(&_Extensionagglayerbridgel2.CallOpts)
}

// PolygonRollupManager is a free data retrieval call binding the contract method 0x8ed7e3f2.
//
// Solidity: function polygonRollupManager() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) PolygonRollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "polygonRollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PolygonRollupManager is a free data retrieval call binding the contract method 0x8ed7e3f2.
//
// Solidity: function polygonRollupManager() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) PolygonRollupManager() (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.PolygonRollupManager(&_Extensionagglayerbridgel2.CallOpts)
}

// PolygonRollupManager is a free data retrieval call binding the contract method 0x8ed7e3f2.
//
// Solidity: function polygonRollupManager() view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) PolygonRollupManager() (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.PolygonRollupManager(&_Extensionagglayerbridgel2.CallOpts)
}

// RemoveLegacySovereignTokenAddress is a free data retrieval call binding the contract method 0xb4586962.
//
// Solidity: function removeLegacySovereignTokenAddress(address ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) RemoveLegacySovereignTokenAddress(opts *bind.CallOpts, arg0 common.Address) error {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "removeLegacySovereignTokenAddress", arg0)

	if err != nil {
		return err
	}

	return err

}

// RemoveLegacySovereignTokenAddress is a free data retrieval call binding the contract method 0xb4586962.
//
// Solidity: function removeLegacySovereignTokenAddress(address ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) RemoveLegacySovereignTokenAddress(arg0 common.Address) error {
	return _Extensionagglayerbridgel2.Contract.RemoveLegacySovereignTokenAddress(&_Extensionagglayerbridgel2.CallOpts, arg0)
}

// RemoveLegacySovereignTokenAddress is a free data retrieval call binding the contract method 0xb4586962.
//
// Solidity: function removeLegacySovereignTokenAddress(address ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) RemoveLegacySovereignTokenAddress(arg0 common.Address) error {
	return _Extensionagglayerbridgel2.Contract.RemoveLegacySovereignTokenAddress(&_Extensionagglayerbridgel2.CallOpts, arg0)
}

// SetBridgeManager is a free data retrieval call binding the contract method 0xeabd372a.
//
// Solidity: function setBridgeManager(address ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) SetBridgeManager(opts *bind.CallOpts, arg0 common.Address) error {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "setBridgeManager", arg0)

	if err != nil {
		return err
	}

	return err

}

// SetBridgeManager is a free data retrieval call binding the contract method 0xeabd372a.
//
// Solidity: function setBridgeManager(address ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) SetBridgeManager(arg0 common.Address) error {
	return _Extensionagglayerbridgel2.Contract.SetBridgeManager(&_Extensionagglayerbridgel2.CallOpts, arg0)
}

// SetBridgeManager is a free data retrieval call binding the contract method 0xeabd372a.
//
// Solidity: function setBridgeManager(address ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) SetBridgeManager(arg0 common.Address) error {
	return _Extensionagglayerbridgel2.Contract.SetBridgeManager(&_Extensionagglayerbridgel2.CallOpts, arg0)
}

// SetLocalBalanceTree is a free data retrieval call binding the contract method 0x8f9720bb.
//
// Solidity: function setLocalBalanceTree(uint32[] , address[] , uint256[] ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) SetLocalBalanceTree(opts *bind.CallOpts, arg0 []uint32, arg1 []common.Address, arg2 []*big.Int) error {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "setLocalBalanceTree", arg0, arg1, arg2)

	if err != nil {
		return err
	}

	return err

}

// SetLocalBalanceTree is a free data retrieval call binding the contract method 0x8f9720bb.
//
// Solidity: function setLocalBalanceTree(uint32[] , address[] , uint256[] ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) SetLocalBalanceTree(arg0 []uint32, arg1 []common.Address, arg2 []*big.Int) error {
	return _Extensionagglayerbridgel2.Contract.SetLocalBalanceTree(&_Extensionagglayerbridgel2.CallOpts, arg0, arg1, arg2)
}

// SetLocalBalanceTree is a free data retrieval call binding the contract method 0x8f9720bb.
//
// Solidity: function setLocalBalanceTree(uint32[] , address[] , uint256[] ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) SetLocalBalanceTree(arg0 []uint32, arg1 []common.Address, arg2 []*big.Int) error {
	return _Extensionagglayerbridgel2.Contract.SetLocalBalanceTree(&_Extensionagglayerbridgel2.CallOpts, arg0, arg1, arg2)
}

// SetMultipleClaims is a free data retrieval call binding the contract method 0x1c208229.
//
// Solidity: function setMultipleClaims(uint256[] ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) SetMultipleClaims(opts *bind.CallOpts, arg0 []*big.Int) error {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "setMultipleClaims", arg0)

	if err != nil {
		return err
	}

	return err

}

// SetMultipleClaims is a free data retrieval call binding the contract method 0x1c208229.
//
// Solidity: function setMultipleClaims(uint256[] ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) SetMultipleClaims(arg0 []*big.Int) error {
	return _Extensionagglayerbridgel2.Contract.SetMultipleClaims(&_Extensionagglayerbridgel2.CallOpts, arg0)
}

// SetMultipleClaims is a free data retrieval call binding the contract method 0x1c208229.
//
// Solidity: function setMultipleClaims(uint256[] ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) SetMultipleClaims(arg0 []*big.Int) error {
	return _Extensionagglayerbridgel2.Contract.SetMultipleClaims(&_Extensionagglayerbridgel2.CallOpts, arg0)
}

// SetMultipleSovereignTokenAddress is a free data retrieval call binding the contract method 0x57cfbee3.
//
// Solidity: function setMultipleSovereignTokenAddress(uint32[] , address[] , address[] , bool[] ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) SetMultipleSovereignTokenAddress(opts *bind.CallOpts, arg0 []uint32, arg1 []common.Address, arg2 []common.Address, arg3 []bool) error {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "setMultipleSovereignTokenAddress", arg0, arg1, arg2, arg3)

	if err != nil {
		return err
	}

	return err

}

// SetMultipleSovereignTokenAddress is a free data retrieval call binding the contract method 0x57cfbee3.
//
// Solidity: function setMultipleSovereignTokenAddress(uint32[] , address[] , address[] , bool[] ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) SetMultipleSovereignTokenAddress(arg0 []uint32, arg1 []common.Address, arg2 []common.Address, arg3 []bool) error {
	return _Extensionagglayerbridgel2.Contract.SetMultipleSovereignTokenAddress(&_Extensionagglayerbridgel2.CallOpts, arg0, arg1, arg2, arg3)
}

// SetMultipleSovereignTokenAddress is a free data retrieval call binding the contract method 0x57cfbee3.
//
// Solidity: function setMultipleSovereignTokenAddress(uint32[] , address[] , address[] , bool[] ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) SetMultipleSovereignTokenAddress(arg0 []uint32, arg1 []common.Address, arg2 []common.Address, arg3 []bool) error {
	return _Extensionagglayerbridgel2.Contract.SetMultipleSovereignTokenAddress(&_Extensionagglayerbridgel2.CallOpts, arg0, arg1, arg2, arg3)
}

// SetSovereignWETHAddress is a free data retrieval call binding the contract method 0xbf130d7f.
//
// Solidity: function setSovereignWETHAddress(address , bool ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) SetSovereignWETHAddress(opts *bind.CallOpts, arg0 common.Address, arg1 bool) error {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "setSovereignWETHAddress", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// SetSovereignWETHAddress is a free data retrieval call binding the contract method 0xbf130d7f.
//
// Solidity: function setSovereignWETHAddress(address , bool ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) SetSovereignWETHAddress(arg0 common.Address, arg1 bool) error {
	return _Extensionagglayerbridgel2.Contract.SetSovereignWETHAddress(&_Extensionagglayerbridgel2.CallOpts, arg0, arg1)
}

// SetSovereignWETHAddress is a free data retrieval call binding the contract method 0xbf130d7f.
//
// Solidity: function setSovereignWETHAddress(address , bool ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) SetSovereignWETHAddress(arg0 common.Address, arg1 bool) error {
	return _Extensionagglayerbridgel2.Contract.SetSovereignWETHAddress(&_Extensionagglayerbridgel2.CallOpts, arg0, arg1)
}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) TokenInfoToWrappedToken(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "tokenInfoToWrappedToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) TokenInfoToWrappedToken(arg0 [32]byte) (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.TokenInfoToWrappedToken(&_Extensionagglayerbridgel2.CallOpts, arg0)
}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) TokenInfoToWrappedToken(arg0 [32]byte) (common.Address, error) {
	return _Extensionagglayerbridgel2.Contract.TokenInfoToWrappedToken(&_Extensionagglayerbridgel2.CallOpts, arg0)
}

// TransferEmergencyBridgePauserRole is a free data retrieval call binding the contract method 0x69e3ab12.
//
// Solidity: function transferEmergencyBridgePauserRole(address ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) TransferEmergencyBridgePauserRole(opts *bind.CallOpts, arg0 common.Address) error {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "transferEmergencyBridgePauserRole", arg0)

	if err != nil {
		return err
	}

	return err

}

// TransferEmergencyBridgePauserRole is a free data retrieval call binding the contract method 0x69e3ab12.
//
// Solidity: function transferEmergencyBridgePauserRole(address ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) TransferEmergencyBridgePauserRole(arg0 common.Address) error {
	return _Extensionagglayerbridgel2.Contract.TransferEmergencyBridgePauserRole(&_Extensionagglayerbridgel2.CallOpts, arg0)
}

// TransferEmergencyBridgePauserRole is a free data retrieval call binding the contract method 0x69e3ab12.
//
// Solidity: function transferEmergencyBridgePauserRole(address ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) TransferEmergencyBridgePauserRole(arg0 common.Address) error {
	return _Extensionagglayerbridgel2.Contract.TransferEmergencyBridgePauserRole(&_Extensionagglayerbridgel2.CallOpts, arg0)
}

// TransferEmergencyBridgeUnpauserRole is a free data retrieval call binding the contract method 0x8d942096.
//
// Solidity: function transferEmergencyBridgeUnpauserRole(address ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) TransferEmergencyBridgeUnpauserRole(opts *bind.CallOpts, arg0 common.Address) error {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "transferEmergencyBridgeUnpauserRole", arg0)

	if err != nil {
		return err
	}

	return err

}

// TransferEmergencyBridgeUnpauserRole is a free data retrieval call binding the contract method 0x8d942096.
//
// Solidity: function transferEmergencyBridgeUnpauserRole(address ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) TransferEmergencyBridgeUnpauserRole(arg0 common.Address) error {
	return _Extensionagglayerbridgel2.Contract.TransferEmergencyBridgeUnpauserRole(&_Extensionagglayerbridgel2.CallOpts, arg0)
}

// TransferEmergencyBridgeUnpauserRole is a free data retrieval call binding the contract method 0x8d942096.
//
// Solidity: function transferEmergencyBridgeUnpauserRole(address ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) TransferEmergencyBridgeUnpauserRole(arg0 common.Address) error {
	return _Extensionagglayerbridgel2.Contract.TransferEmergencyBridgeUnpauserRole(&_Extensionagglayerbridgel2.CallOpts, arg0)
}

// TransferProxiedTokensManagerRole is a free data retrieval call binding the contract method 0x8bd309c3.
//
// Solidity: function transferProxiedTokensManagerRole(address ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) TransferProxiedTokensManagerRole(opts *bind.CallOpts, arg0 common.Address) error {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "transferProxiedTokensManagerRole", arg0)

	if err != nil {
		return err
	}

	return err

}

// TransferProxiedTokensManagerRole is a free data retrieval call binding the contract method 0x8bd309c3.
//
// Solidity: function transferProxiedTokensManagerRole(address ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) TransferProxiedTokensManagerRole(arg0 common.Address) error {
	return _Extensionagglayerbridgel2.Contract.TransferProxiedTokensManagerRole(&_Extensionagglayerbridgel2.CallOpts, arg0)
}

// TransferProxiedTokensManagerRole is a free data retrieval call binding the contract method 0x8bd309c3.
//
// Solidity: function transferProxiedTokensManagerRole(address ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) TransferProxiedTokensManagerRole(arg0 common.Address) error {
	return _Extensionagglayerbridgel2.Contract.TransferProxiedTokensManagerRole(&_Extensionagglayerbridgel2.CallOpts, arg0)
}

// UnsetGlobalIndexHashChain is a free data retrieval call binding the contract method 0x6ee84b23.
//
// Solidity: function unsetGlobalIndexHashChain() view returns(bytes32)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) UnsetGlobalIndexHashChain(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "unsetGlobalIndexHashChain")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UnsetGlobalIndexHashChain is a free data retrieval call binding the contract method 0x6ee84b23.
//
// Solidity: function unsetGlobalIndexHashChain() view returns(bytes32)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) UnsetGlobalIndexHashChain() ([32]byte, error) {
	return _Extensionagglayerbridgel2.Contract.UnsetGlobalIndexHashChain(&_Extensionagglayerbridgel2.CallOpts)
}

// UnsetGlobalIndexHashChain is a free data retrieval call binding the contract method 0x6ee84b23.
//
// Solidity: function unsetGlobalIndexHashChain() view returns(bytes32)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) UnsetGlobalIndexHashChain() ([32]byte, error) {
	return _Extensionagglayerbridgel2.Contract.UnsetGlobalIndexHashChain(&_Extensionagglayerbridgel2.CallOpts)
}

// UnsetMultipleClaims is a free data retrieval call binding the contract method 0x136a2c60.
//
// Solidity: function unsetMultipleClaims(uint256[] ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) UnsetMultipleClaims(opts *bind.CallOpts, arg0 []*big.Int) error {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "unsetMultipleClaims", arg0)

	if err != nil {
		return err
	}

	return err

}

// UnsetMultipleClaims is a free data retrieval call binding the contract method 0x136a2c60.
//
// Solidity: function unsetMultipleClaims(uint256[] ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) UnsetMultipleClaims(arg0 []*big.Int) error {
	return _Extensionagglayerbridgel2.Contract.UnsetMultipleClaims(&_Extensionagglayerbridgel2.CallOpts, arg0)
}

// UnsetMultipleClaims is a free data retrieval call binding the contract method 0x136a2c60.
//
// Solidity: function unsetMultipleClaims(uint256[] ) pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) UnsetMultipleClaims(arg0 []*big.Int) error {
	return _Extensionagglayerbridgel2.Contract.UnsetMultipleClaims(&_Extensionagglayerbridgel2.CallOpts, arg0)
}

// UpdateGlobalExitRoot is a free data retrieval call binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) UpdateGlobalExitRoot(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "updateGlobalExitRoot")

	if err != nil {
		return err
	}

	return err

}

// UpdateGlobalExitRoot is a free data retrieval call binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) UpdateGlobalExitRoot() error {
	return _Extensionagglayerbridgel2.Contract.UpdateGlobalExitRoot(&_Extensionagglayerbridgel2.CallOpts)
}

// UpdateGlobalExitRoot is a free data retrieval call binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() pure returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) UpdateGlobalExitRoot() error {
	return _Extensionagglayerbridgel2.Contract.UpdateGlobalExitRoot(&_Extensionagglayerbridgel2.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) Version() (string, error) {
	return _Extensionagglayerbridgel2.Contract.Version(&_Extensionagglayerbridgel2.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) Version() (string, error) {
	return _Extensionagglayerbridgel2.Contract.Version(&_Extensionagglayerbridgel2.CallOpts)
}

// WrappedAddressIsNotMintable is a free data retrieval call binding the contract method 0xc0f49163.
//
// Solidity: function wrappedAddressIsNotMintable(address wrappedAddress) view returns(bool isNotMintable)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) WrappedAddressIsNotMintable(opts *bind.CallOpts, wrappedAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "wrappedAddressIsNotMintable", wrappedAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WrappedAddressIsNotMintable is a free data retrieval call binding the contract method 0xc0f49163.
//
// Solidity: function wrappedAddressIsNotMintable(address wrappedAddress) view returns(bool isNotMintable)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) WrappedAddressIsNotMintable(wrappedAddress common.Address) (bool, error) {
	return _Extensionagglayerbridgel2.Contract.WrappedAddressIsNotMintable(&_Extensionagglayerbridgel2.CallOpts, wrappedAddress)
}

// WrappedAddressIsNotMintable is a free data retrieval call binding the contract method 0xc0f49163.
//
// Solidity: function wrappedAddressIsNotMintable(address wrappedAddress) view returns(bool isNotMintable)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) WrappedAddressIsNotMintable(wrappedAddress common.Address) (bool, error) {
	return _Extensionagglayerbridgel2.Contract.WrappedAddressIsNotMintable(&_Extensionagglayerbridgel2.CallOpts, wrappedAddress)
}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address ) view returns(uint32 originNetwork, address originTokenAddress)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Caller) WrappedTokenToTokenInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	var out []interface{}
	err := _Extensionagglayerbridgel2.contract.Call(opts, &out, "wrappedTokenToTokenInfo", arg0)

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
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) WrappedTokenToTokenInfo(arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	return _Extensionagglayerbridgel2.Contract.WrappedTokenToTokenInfo(&_Extensionagglayerbridgel2.CallOpts, arg0)
}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address ) view returns(uint32 originNetwork, address originTokenAddress)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2CallerSession) WrappedTokenToTokenInfo(arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	return _Extensionagglayerbridgel2.Contract.WrappedTokenToTokenInfo(&_Extensionagglayerbridgel2.CallOpts, arg0)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 , address , uint256 , address , bool , bytes ) payable returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Transactor) BridgeAsset(opts *bind.TransactOpts, arg0 uint32, arg1 common.Address, arg2 *big.Int, arg3 common.Address, arg4 bool, arg5 []byte) (*types.Transaction, error) {
	return _Extensionagglayerbridgel2.contract.Transact(opts, "bridgeAsset", arg0, arg1, arg2, arg3, arg4, arg5)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 , address , uint256 , address , bool , bytes ) payable returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) BridgeAsset(arg0 uint32, arg1 common.Address, arg2 *big.Int, arg3 common.Address, arg4 bool, arg5 []byte) (*types.Transaction, error) {
	return _Extensionagglayerbridgel2.Contract.BridgeAsset(&_Extensionagglayerbridgel2.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 , address , uint256 , address , bool , bytes ) payable returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2TransactorSession) BridgeAsset(arg0 uint32, arg1 common.Address, arg2 *big.Int, arg3 common.Address, arg4 bool, arg5 []byte) (*types.Transaction, error) {
	return _Extensionagglayerbridgel2.Contract.BridgeAsset(&_Extensionagglayerbridgel2.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 , address , bool , bytes ) payable returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Transactor) BridgeMessage(opts *bind.TransactOpts, arg0 uint32, arg1 common.Address, arg2 bool, arg3 []byte) (*types.Transaction, error) {
	return _Extensionagglayerbridgel2.contract.Transact(opts, "bridgeMessage", arg0, arg1, arg2, arg3)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 , address , bool , bytes ) payable returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) BridgeMessage(arg0 uint32, arg1 common.Address, arg2 bool, arg3 []byte) (*types.Transaction, error) {
	return _Extensionagglayerbridgel2.Contract.BridgeMessage(&_Extensionagglayerbridgel2.TransactOpts, arg0, arg1, arg2, arg3)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 , address , bool , bytes ) payable returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2TransactorSession) BridgeMessage(arg0 uint32, arg1 common.Address, arg2 bool, arg3 []byte) (*types.Transaction, error) {
	return _Extensionagglayerbridgel2.Contract.BridgeMessage(&_Extensionagglayerbridgel2.TransactOpts, arg0, arg1, arg2, arg3)
}

// Initialize is a paid mutator transaction binding the contract method 0x006ee171.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address _sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable, address _emergencyBridgePauser, address _emergencyBridgeUnpauser, address _proxiedTokensManager) returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Transactor) Initialize(opts *bind.TransactOpts, _networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, _sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool, _emergencyBridgePauser common.Address, _emergencyBridgeUnpauser common.Address, _proxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Extensionagglayerbridgel2.contract.Transact(opts, "initialize", _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, _sovereignWETHAddress, _sovereignWETHAddressIsNotMintable, _emergencyBridgePauser, _emergencyBridgeUnpauser, _proxiedTokensManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x006ee171.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address _sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable, address _emergencyBridgePauser, address _emergencyBridgeUnpauser, address _proxiedTokensManager) returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) Initialize(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, _sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool, _emergencyBridgePauser common.Address, _emergencyBridgeUnpauser common.Address, _proxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Extensionagglayerbridgel2.Contract.Initialize(&_Extensionagglayerbridgel2.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, _sovereignWETHAddress, _sovereignWETHAddressIsNotMintable, _emergencyBridgePauser, _emergencyBridgeUnpauser, _proxiedTokensManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x006ee171.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address _sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable, address _emergencyBridgePauser, address _emergencyBridgeUnpauser, address _proxiedTokensManager) returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2TransactorSession) Initialize(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, _sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool, _emergencyBridgePauser common.Address, _emergencyBridgeUnpauser common.Address, _proxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Extensionagglayerbridgel2.Contract.Initialize(&_Extensionagglayerbridgel2.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, _sovereignWETHAddress, _sovereignWETHAddressIsNotMintable, _emergencyBridgePauser, _emergencyBridgeUnpauser, _proxiedTokensManager)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes ) returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Transactor) Initialize0(opts *bind.TransactOpts, arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte) (*types.Transaction, error) {
	return _Extensionagglayerbridgel2.contract.Transact(opts, "initialize0", arg0, arg1, arg2, arg3, arg4, arg5)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes ) returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Session) Initialize0(arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte) (*types.Transaction, error) {
	return _Extensionagglayerbridgel2.Contract.Initialize0(&_Extensionagglayerbridgel2.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes ) returns()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2TransactorSession) Initialize0(arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte) (*types.Transaction, error) {
	return _Extensionagglayerbridgel2.Contract.Initialize0(&_Extensionagglayerbridgel2.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// Extensionagglayerbridgel2AcceptEmergencyBridgePauserRoleIterator is returned from FilterAcceptEmergencyBridgePauserRole and is used to iterate over the raw logs and unpacked data for AcceptEmergencyBridgePauserRole events raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2AcceptEmergencyBridgePauserRoleIterator struct {
	Event *Extensionagglayerbridgel2AcceptEmergencyBridgePauserRole // Event containing the contract specifics and raw log

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
func (it *Extensionagglayerbridgel2AcceptEmergencyBridgePauserRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Extensionagglayerbridgel2AcceptEmergencyBridgePauserRole)
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
		it.Event = new(Extensionagglayerbridgel2AcceptEmergencyBridgePauserRole)
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
func (it *Extensionagglayerbridgel2AcceptEmergencyBridgePauserRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Extensionagglayerbridgel2AcceptEmergencyBridgePauserRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Extensionagglayerbridgel2AcceptEmergencyBridgePauserRole represents a AcceptEmergencyBridgePauserRole event raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2AcceptEmergencyBridgePauserRole struct {
	OldEmergencyBridgePauser common.Address
	NewEmergencyBridgePauser common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterAcceptEmergencyBridgePauserRole is a free log retrieval operation binding the contract event 0x85d2bdfbe58cd81abf8199c13ce2509204be4aba8603b9d29f52c4e13e7bb793.
//
// Solidity: event AcceptEmergencyBridgePauserRole(address oldEmergencyBridgePauser, address newEmergencyBridgePauser)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) FilterAcceptEmergencyBridgePauserRole(opts *bind.FilterOpts) (*Extensionagglayerbridgel2AcceptEmergencyBridgePauserRoleIterator, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.FilterLogs(opts, "AcceptEmergencyBridgePauserRole")
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2AcceptEmergencyBridgePauserRoleIterator{contract: _Extensionagglayerbridgel2.contract, event: "AcceptEmergencyBridgePauserRole", logs: logs, sub: sub}, nil
}

// WatchAcceptEmergencyBridgePauserRole is a free log subscription operation binding the contract event 0x85d2bdfbe58cd81abf8199c13ce2509204be4aba8603b9d29f52c4e13e7bb793.
//
// Solidity: event AcceptEmergencyBridgePauserRole(address oldEmergencyBridgePauser, address newEmergencyBridgePauser)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) WatchAcceptEmergencyBridgePauserRole(opts *bind.WatchOpts, sink chan<- *Extensionagglayerbridgel2AcceptEmergencyBridgePauserRole) (event.Subscription, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.WatchLogs(opts, "AcceptEmergencyBridgePauserRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Extensionagglayerbridgel2AcceptEmergencyBridgePauserRole)
				if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "AcceptEmergencyBridgePauserRole", log); err != nil {
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
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) ParseAcceptEmergencyBridgePauserRole(log types.Log) (*Extensionagglayerbridgel2AcceptEmergencyBridgePauserRole, error) {
	event := new(Extensionagglayerbridgel2AcceptEmergencyBridgePauserRole)
	if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "AcceptEmergencyBridgePauserRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Extensionagglayerbridgel2AcceptEmergencyBridgeUnpauserRoleIterator is returned from FilterAcceptEmergencyBridgeUnpauserRole and is used to iterate over the raw logs and unpacked data for AcceptEmergencyBridgeUnpauserRole events raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2AcceptEmergencyBridgeUnpauserRoleIterator struct {
	Event *Extensionagglayerbridgel2AcceptEmergencyBridgeUnpauserRole // Event containing the contract specifics and raw log

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
func (it *Extensionagglayerbridgel2AcceptEmergencyBridgeUnpauserRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Extensionagglayerbridgel2AcceptEmergencyBridgeUnpauserRole)
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
		it.Event = new(Extensionagglayerbridgel2AcceptEmergencyBridgeUnpauserRole)
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
func (it *Extensionagglayerbridgel2AcceptEmergencyBridgeUnpauserRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Extensionagglayerbridgel2AcceptEmergencyBridgeUnpauserRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Extensionagglayerbridgel2AcceptEmergencyBridgeUnpauserRole represents a AcceptEmergencyBridgeUnpauserRole event raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2AcceptEmergencyBridgeUnpauserRole struct {
	OldEmergencyBridgeUnpauser common.Address
	NewEmergencyBridgeUnpauser common.Address
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterAcceptEmergencyBridgeUnpauserRole is a free log retrieval operation binding the contract event 0x24cc8295aa5110cc216695db944ad2458c7795c6404449be980c3ce14aed752d.
//
// Solidity: event AcceptEmergencyBridgeUnpauserRole(address oldEmergencyBridgeUnpauser, address newEmergencyBridgeUnpauser)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) FilterAcceptEmergencyBridgeUnpauserRole(opts *bind.FilterOpts) (*Extensionagglayerbridgel2AcceptEmergencyBridgeUnpauserRoleIterator, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.FilterLogs(opts, "AcceptEmergencyBridgeUnpauserRole")
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2AcceptEmergencyBridgeUnpauserRoleIterator{contract: _Extensionagglayerbridgel2.contract, event: "AcceptEmergencyBridgeUnpauserRole", logs: logs, sub: sub}, nil
}

// WatchAcceptEmergencyBridgeUnpauserRole is a free log subscription operation binding the contract event 0x24cc8295aa5110cc216695db944ad2458c7795c6404449be980c3ce14aed752d.
//
// Solidity: event AcceptEmergencyBridgeUnpauserRole(address oldEmergencyBridgeUnpauser, address newEmergencyBridgeUnpauser)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) WatchAcceptEmergencyBridgeUnpauserRole(opts *bind.WatchOpts, sink chan<- *Extensionagglayerbridgel2AcceptEmergencyBridgeUnpauserRole) (event.Subscription, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.WatchLogs(opts, "AcceptEmergencyBridgeUnpauserRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Extensionagglayerbridgel2AcceptEmergencyBridgeUnpauserRole)
				if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "AcceptEmergencyBridgeUnpauserRole", log); err != nil {
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
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) ParseAcceptEmergencyBridgeUnpauserRole(log types.Log) (*Extensionagglayerbridgel2AcceptEmergencyBridgeUnpauserRole, error) {
	event := new(Extensionagglayerbridgel2AcceptEmergencyBridgeUnpauserRole)
	if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "AcceptEmergencyBridgeUnpauserRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Extensionagglayerbridgel2AcceptProxiedTokensManagerRoleIterator is returned from FilterAcceptProxiedTokensManagerRole and is used to iterate over the raw logs and unpacked data for AcceptProxiedTokensManagerRole events raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2AcceptProxiedTokensManagerRoleIterator struct {
	Event *Extensionagglayerbridgel2AcceptProxiedTokensManagerRole // Event containing the contract specifics and raw log

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
func (it *Extensionagglayerbridgel2AcceptProxiedTokensManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Extensionagglayerbridgel2AcceptProxiedTokensManagerRole)
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
		it.Event = new(Extensionagglayerbridgel2AcceptProxiedTokensManagerRole)
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
func (it *Extensionagglayerbridgel2AcceptProxiedTokensManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Extensionagglayerbridgel2AcceptProxiedTokensManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Extensionagglayerbridgel2AcceptProxiedTokensManagerRole represents a AcceptProxiedTokensManagerRole event raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2AcceptProxiedTokensManagerRole struct {
	OldProxiedTokensManager common.Address
	NewProxiedTokensManager common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterAcceptProxiedTokensManagerRole is a free log retrieval operation binding the contract event 0xa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f2.
//
// Solidity: event AcceptProxiedTokensManagerRole(address oldProxiedTokensManager, address newProxiedTokensManager)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) FilterAcceptProxiedTokensManagerRole(opts *bind.FilterOpts) (*Extensionagglayerbridgel2AcceptProxiedTokensManagerRoleIterator, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.FilterLogs(opts, "AcceptProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2AcceptProxiedTokensManagerRoleIterator{contract: _Extensionagglayerbridgel2.contract, event: "AcceptProxiedTokensManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptProxiedTokensManagerRole is a free log subscription operation binding the contract event 0xa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f2.
//
// Solidity: event AcceptProxiedTokensManagerRole(address oldProxiedTokensManager, address newProxiedTokensManager)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) WatchAcceptProxiedTokensManagerRole(opts *bind.WatchOpts, sink chan<- *Extensionagglayerbridgel2AcceptProxiedTokensManagerRole) (event.Subscription, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.WatchLogs(opts, "AcceptProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Extensionagglayerbridgel2AcceptProxiedTokensManagerRole)
				if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "AcceptProxiedTokensManagerRole", log); err != nil {
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
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) ParseAcceptProxiedTokensManagerRole(log types.Log) (*Extensionagglayerbridgel2AcceptProxiedTokensManagerRole, error) {
	event := new(Extensionagglayerbridgel2AcceptProxiedTokensManagerRole)
	if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "AcceptProxiedTokensManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Extensionagglayerbridgel2BackwardLETIterator is returned from FilterBackwardLET and is used to iterate over the raw logs and unpacked data for BackwardLET events raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2BackwardLETIterator struct {
	Event *Extensionagglayerbridgel2BackwardLET // Event containing the contract specifics and raw log

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
func (it *Extensionagglayerbridgel2BackwardLETIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Extensionagglayerbridgel2BackwardLET)
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
		it.Event = new(Extensionagglayerbridgel2BackwardLET)
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
func (it *Extensionagglayerbridgel2BackwardLETIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Extensionagglayerbridgel2BackwardLETIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Extensionagglayerbridgel2BackwardLET represents a BackwardLET event raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2BackwardLET struct {
	PreviousDepositCount *big.Int
	PreviousRoot         [32]byte
	NewDepositCount      *big.Int
	NewRoot              [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterBackwardLET is a free log retrieval operation binding the contract event 0x2f8b2f36ca0d63af2efffe9d5e728a61c2bc273a3fc0574d12b558a41f149555.
//
// Solidity: event BackwardLET(uint256 previousDepositCount, bytes32 previousRoot, uint256 newDepositCount, bytes32 newRoot)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) FilterBackwardLET(opts *bind.FilterOpts) (*Extensionagglayerbridgel2BackwardLETIterator, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.FilterLogs(opts, "BackwardLET")
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2BackwardLETIterator{contract: _Extensionagglayerbridgel2.contract, event: "BackwardLET", logs: logs, sub: sub}, nil
}

// WatchBackwardLET is a free log subscription operation binding the contract event 0x2f8b2f36ca0d63af2efffe9d5e728a61c2bc273a3fc0574d12b558a41f149555.
//
// Solidity: event BackwardLET(uint256 previousDepositCount, bytes32 previousRoot, uint256 newDepositCount, bytes32 newRoot)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) WatchBackwardLET(opts *bind.WatchOpts, sink chan<- *Extensionagglayerbridgel2BackwardLET) (event.Subscription, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.WatchLogs(opts, "BackwardLET")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Extensionagglayerbridgel2BackwardLET)
				if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "BackwardLET", log); err != nil {
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

// ParseBackwardLET is a log parse operation binding the contract event 0x2f8b2f36ca0d63af2efffe9d5e728a61c2bc273a3fc0574d12b558a41f149555.
//
// Solidity: event BackwardLET(uint256 previousDepositCount, bytes32 previousRoot, uint256 newDepositCount, bytes32 newRoot)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) ParseBackwardLET(log types.Log) (*Extensionagglayerbridgel2BackwardLET, error) {
	event := new(Extensionagglayerbridgel2BackwardLET)
	if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "BackwardLET", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Extensionagglayerbridgel2BridgeEventIterator is returned from FilterBridgeEvent and is used to iterate over the raw logs and unpacked data for BridgeEvent events raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2BridgeEventIterator struct {
	Event *Extensionagglayerbridgel2BridgeEvent // Event containing the contract specifics and raw log

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
func (it *Extensionagglayerbridgel2BridgeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Extensionagglayerbridgel2BridgeEvent)
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
		it.Event = new(Extensionagglayerbridgel2BridgeEvent)
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
func (it *Extensionagglayerbridgel2BridgeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Extensionagglayerbridgel2BridgeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Extensionagglayerbridgel2BridgeEvent represents a BridgeEvent event raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2BridgeEvent struct {
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
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) FilterBridgeEvent(opts *bind.FilterOpts) (*Extensionagglayerbridgel2BridgeEventIterator, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.FilterLogs(opts, "BridgeEvent")
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2BridgeEventIterator{contract: _Extensionagglayerbridgel2.contract, event: "BridgeEvent", logs: logs, sub: sub}, nil
}

// WatchBridgeEvent is a free log subscription operation binding the contract event 0x501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b.
//
// Solidity: event BridgeEvent(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata, uint32 depositCount)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) WatchBridgeEvent(opts *bind.WatchOpts, sink chan<- *Extensionagglayerbridgel2BridgeEvent) (event.Subscription, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.WatchLogs(opts, "BridgeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Extensionagglayerbridgel2BridgeEvent)
				if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "BridgeEvent", log); err != nil {
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
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) ParseBridgeEvent(log types.Log) (*Extensionagglayerbridgel2BridgeEvent, error) {
	event := new(Extensionagglayerbridgel2BridgeEvent)
	if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "BridgeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Extensionagglayerbridgel2ClaimEventIterator is returned from FilterClaimEvent and is used to iterate over the raw logs and unpacked data for ClaimEvent events raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2ClaimEventIterator struct {
	Event *Extensionagglayerbridgel2ClaimEvent // Event containing the contract specifics and raw log

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
func (it *Extensionagglayerbridgel2ClaimEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Extensionagglayerbridgel2ClaimEvent)
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
		it.Event = new(Extensionagglayerbridgel2ClaimEvent)
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
func (it *Extensionagglayerbridgel2ClaimEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Extensionagglayerbridgel2ClaimEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Extensionagglayerbridgel2ClaimEvent represents a ClaimEvent event raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2ClaimEvent struct {
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
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) FilterClaimEvent(opts *bind.FilterOpts) (*Extensionagglayerbridgel2ClaimEventIterator, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.FilterLogs(opts, "ClaimEvent")
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2ClaimEventIterator{contract: _Extensionagglayerbridgel2.contract, event: "ClaimEvent", logs: logs, sub: sub}, nil
}

// WatchClaimEvent is a free log subscription operation binding the contract event 0x1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d.
//
// Solidity: event ClaimEvent(uint256 globalIndex, uint32 originNetwork, address originAddress, address destinationAddress, uint256 amount)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) WatchClaimEvent(opts *bind.WatchOpts, sink chan<- *Extensionagglayerbridgel2ClaimEvent) (event.Subscription, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.WatchLogs(opts, "ClaimEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Extensionagglayerbridgel2ClaimEvent)
				if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "ClaimEvent", log); err != nil {
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
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) ParseClaimEvent(log types.Log) (*Extensionagglayerbridgel2ClaimEvent, error) {
	event := new(Extensionagglayerbridgel2ClaimEvent)
	if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "ClaimEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Extensionagglayerbridgel2DetailedClaimEventIterator is returned from FilterDetailedClaimEvent and is used to iterate over the raw logs and unpacked data for DetailedClaimEvent events raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2DetailedClaimEventIterator struct {
	Event *Extensionagglayerbridgel2DetailedClaimEvent // Event containing the contract specifics and raw log

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
func (it *Extensionagglayerbridgel2DetailedClaimEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Extensionagglayerbridgel2DetailedClaimEvent)
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
		it.Event = new(Extensionagglayerbridgel2DetailedClaimEvent)
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
func (it *Extensionagglayerbridgel2DetailedClaimEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Extensionagglayerbridgel2DetailedClaimEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Extensionagglayerbridgel2DetailedClaimEvent represents a DetailedClaimEvent event raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2DetailedClaimEvent struct {
	SmtProofLocalExitRoot  [32][32]byte
	SmtProofRollupExitRoot [32][32]byte
	GlobalIndex            *big.Int
	MainnetExitRoot        [32]byte
	RollupExitRoot         [32]byte
	LeafType               uint8
	OriginNetwork          uint32
	OriginTokenAddress     common.Address
	DestinationNetwork     uint32
	DestinationAddress     common.Address
	Amount                 *big.Int
	Metadata               []byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterDetailedClaimEvent is a free log retrieval operation binding the contract event 0xfc09ab352fe828dc1fbb0cad35e812349ba00649f10067f83eafa8d79b161d81.
//
// Solidity: event DetailedClaimEvent(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 indexed globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint8 leafType, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address indexed destinationAddress, uint256 amount, bytes metadata)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) FilterDetailedClaimEvent(opts *bind.FilterOpts, globalIndex []*big.Int, destinationAddress []common.Address) (*Extensionagglayerbridgel2DetailedClaimEventIterator, error) {

	var globalIndexRule []interface{}
	for _, globalIndexItem := range globalIndex {
		globalIndexRule = append(globalIndexRule, globalIndexItem)
	}

	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	logs, sub, err := _Extensionagglayerbridgel2.contract.FilterLogs(opts, "DetailedClaimEvent", globalIndexRule, destinationAddressRule)
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2DetailedClaimEventIterator{contract: _Extensionagglayerbridgel2.contract, event: "DetailedClaimEvent", logs: logs, sub: sub}, nil
}

// WatchDetailedClaimEvent is a free log subscription operation binding the contract event 0xfc09ab352fe828dc1fbb0cad35e812349ba00649f10067f83eafa8d79b161d81.
//
// Solidity: event DetailedClaimEvent(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 indexed globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint8 leafType, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address indexed destinationAddress, uint256 amount, bytes metadata)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) WatchDetailedClaimEvent(opts *bind.WatchOpts, sink chan<- *Extensionagglayerbridgel2DetailedClaimEvent, globalIndex []*big.Int, destinationAddress []common.Address) (event.Subscription, error) {

	var globalIndexRule []interface{}
	for _, globalIndexItem := range globalIndex {
		globalIndexRule = append(globalIndexRule, globalIndexItem)
	}

	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	logs, sub, err := _Extensionagglayerbridgel2.contract.WatchLogs(opts, "DetailedClaimEvent", globalIndexRule, destinationAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Extensionagglayerbridgel2DetailedClaimEvent)
				if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "DetailedClaimEvent", log); err != nil {
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

// ParseDetailedClaimEvent is a log parse operation binding the contract event 0xfc09ab352fe828dc1fbb0cad35e812349ba00649f10067f83eafa8d79b161d81.
//
// Solidity: event DetailedClaimEvent(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 indexed globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint8 leafType, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address indexed destinationAddress, uint256 amount, bytes metadata)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) ParseDetailedClaimEvent(log types.Log) (*Extensionagglayerbridgel2DetailedClaimEvent, error) {
	event := new(Extensionagglayerbridgel2DetailedClaimEvent)
	if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "DetailedClaimEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Extensionagglayerbridgel2EmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2EmergencyStateActivatedIterator struct {
	Event *Extensionagglayerbridgel2EmergencyStateActivated // Event containing the contract specifics and raw log

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
func (it *Extensionagglayerbridgel2EmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Extensionagglayerbridgel2EmergencyStateActivated)
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
		it.Event = new(Extensionagglayerbridgel2EmergencyStateActivated)
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
func (it *Extensionagglayerbridgel2EmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Extensionagglayerbridgel2EmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Extensionagglayerbridgel2EmergencyStateActivated represents a EmergencyStateActivated event raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2EmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*Extensionagglayerbridgel2EmergencyStateActivatedIterator, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2EmergencyStateActivatedIterator{contract: _Extensionagglayerbridgel2.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *Extensionagglayerbridgel2EmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Extensionagglayerbridgel2EmergencyStateActivated)
				if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
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
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) ParseEmergencyStateActivated(log types.Log) (*Extensionagglayerbridgel2EmergencyStateActivated, error) {
	event := new(Extensionagglayerbridgel2EmergencyStateActivated)
	if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Extensionagglayerbridgel2EmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2EmergencyStateDeactivatedIterator struct {
	Event *Extensionagglayerbridgel2EmergencyStateDeactivated // Event containing the contract specifics and raw log

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
func (it *Extensionagglayerbridgel2EmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Extensionagglayerbridgel2EmergencyStateDeactivated)
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
		it.Event = new(Extensionagglayerbridgel2EmergencyStateDeactivated)
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
func (it *Extensionagglayerbridgel2EmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Extensionagglayerbridgel2EmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Extensionagglayerbridgel2EmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2EmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*Extensionagglayerbridgel2EmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2EmergencyStateDeactivatedIterator{contract: _Extensionagglayerbridgel2.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *Extensionagglayerbridgel2EmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Extensionagglayerbridgel2EmergencyStateDeactivated)
				if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
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
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) ParseEmergencyStateDeactivated(log types.Log) (*Extensionagglayerbridgel2EmergencyStateDeactivated, error) {
	event := new(Extensionagglayerbridgel2EmergencyStateDeactivated)
	if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Extensionagglayerbridgel2ForwardLETIterator is returned from FilterForwardLET and is used to iterate over the raw logs and unpacked data for ForwardLET events raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2ForwardLETIterator struct {
	Event *Extensionagglayerbridgel2ForwardLET // Event containing the contract specifics and raw log

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
func (it *Extensionagglayerbridgel2ForwardLETIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Extensionagglayerbridgel2ForwardLET)
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
		it.Event = new(Extensionagglayerbridgel2ForwardLET)
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
func (it *Extensionagglayerbridgel2ForwardLETIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Extensionagglayerbridgel2ForwardLETIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Extensionagglayerbridgel2ForwardLET represents a ForwardLET event raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2ForwardLET struct {
	PreviousDepositCount *big.Int
	PreviousRoot         [32]byte
	NewDepositCount      *big.Int
	NewRoot              [32]byte
	NewLeaves            []byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterForwardLET is a free log retrieval operation binding the contract event 0xe90f3fcbcd86a778012608da0cd7622521c8dac9919dddf25a3fb634adb3d2c8.
//
// Solidity: event ForwardLET(uint256 previousDepositCount, bytes32 previousRoot, uint256 newDepositCount, bytes32 newRoot, bytes newLeaves)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) FilterForwardLET(opts *bind.FilterOpts) (*Extensionagglayerbridgel2ForwardLETIterator, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.FilterLogs(opts, "ForwardLET")
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2ForwardLETIterator{contract: _Extensionagglayerbridgel2.contract, event: "ForwardLET", logs: logs, sub: sub}, nil
}

// WatchForwardLET is a free log subscription operation binding the contract event 0xe90f3fcbcd86a778012608da0cd7622521c8dac9919dddf25a3fb634adb3d2c8.
//
// Solidity: event ForwardLET(uint256 previousDepositCount, bytes32 previousRoot, uint256 newDepositCount, bytes32 newRoot, bytes newLeaves)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) WatchForwardLET(opts *bind.WatchOpts, sink chan<- *Extensionagglayerbridgel2ForwardLET) (event.Subscription, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.WatchLogs(opts, "ForwardLET")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Extensionagglayerbridgel2ForwardLET)
				if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "ForwardLET", log); err != nil {
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

// ParseForwardLET is a log parse operation binding the contract event 0xe90f3fcbcd86a778012608da0cd7622521c8dac9919dddf25a3fb634adb3d2c8.
//
// Solidity: event ForwardLET(uint256 previousDepositCount, bytes32 previousRoot, uint256 newDepositCount, bytes32 newRoot, bytes newLeaves)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) ParseForwardLET(log types.Log) (*Extensionagglayerbridgel2ForwardLET, error) {
	event := new(Extensionagglayerbridgel2ForwardLET)
	if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "ForwardLET", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Extensionagglayerbridgel2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2InitializedIterator struct {
	Event *Extensionagglayerbridgel2Initialized // Event containing the contract specifics and raw log

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
func (it *Extensionagglayerbridgel2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Extensionagglayerbridgel2Initialized)
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
		it.Event = new(Extensionagglayerbridgel2Initialized)
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
func (it *Extensionagglayerbridgel2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Extensionagglayerbridgel2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Extensionagglayerbridgel2Initialized represents a Initialized event raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) FilterInitialized(opts *bind.FilterOpts) (*Extensionagglayerbridgel2InitializedIterator, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2InitializedIterator{contract: _Extensionagglayerbridgel2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Extensionagglayerbridgel2Initialized) (event.Subscription, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Extensionagglayerbridgel2Initialized)
				if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) ParseInitialized(log types.Log) (*Extensionagglayerbridgel2Initialized, error) {
	event := new(Extensionagglayerbridgel2Initialized)
	if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Extensionagglayerbridgel2MigrateLegacyTokenIterator is returned from FilterMigrateLegacyToken and is used to iterate over the raw logs and unpacked data for MigrateLegacyToken events raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2MigrateLegacyTokenIterator struct {
	Event *Extensionagglayerbridgel2MigrateLegacyToken // Event containing the contract specifics and raw log

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
func (it *Extensionagglayerbridgel2MigrateLegacyTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Extensionagglayerbridgel2MigrateLegacyToken)
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
		it.Event = new(Extensionagglayerbridgel2MigrateLegacyToken)
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
func (it *Extensionagglayerbridgel2MigrateLegacyTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Extensionagglayerbridgel2MigrateLegacyTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Extensionagglayerbridgel2MigrateLegacyToken represents a MigrateLegacyToken event raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2MigrateLegacyToken struct {
	Sender              common.Address
	LegacyTokenAddress  common.Address
	UpdatedTokenAddress common.Address
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMigrateLegacyToken is a free log retrieval operation binding the contract event 0xb7f8fd4d1faf9b2929dc269f59c53e3a2bccc44e9950f33a568fcbcb37eb69a9.
//
// Solidity: event MigrateLegacyToken(address sender, address legacyTokenAddress, address updatedTokenAddress, uint256 amount)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) FilterMigrateLegacyToken(opts *bind.FilterOpts) (*Extensionagglayerbridgel2MigrateLegacyTokenIterator, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.FilterLogs(opts, "MigrateLegacyToken")
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2MigrateLegacyTokenIterator{contract: _Extensionagglayerbridgel2.contract, event: "MigrateLegacyToken", logs: logs, sub: sub}, nil
}

// WatchMigrateLegacyToken is a free log subscription operation binding the contract event 0xb7f8fd4d1faf9b2929dc269f59c53e3a2bccc44e9950f33a568fcbcb37eb69a9.
//
// Solidity: event MigrateLegacyToken(address sender, address legacyTokenAddress, address updatedTokenAddress, uint256 amount)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) WatchMigrateLegacyToken(opts *bind.WatchOpts, sink chan<- *Extensionagglayerbridgel2MigrateLegacyToken) (event.Subscription, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.WatchLogs(opts, "MigrateLegacyToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Extensionagglayerbridgel2MigrateLegacyToken)
				if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "MigrateLegacyToken", log); err != nil {
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
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) ParseMigrateLegacyToken(log types.Log) (*Extensionagglayerbridgel2MigrateLegacyToken, error) {
	event := new(Extensionagglayerbridgel2MigrateLegacyToken)
	if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "MigrateLegacyToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Extensionagglayerbridgel2NewWrappedTokenIterator is returned from FilterNewWrappedToken and is used to iterate over the raw logs and unpacked data for NewWrappedToken events raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2NewWrappedTokenIterator struct {
	Event *Extensionagglayerbridgel2NewWrappedToken // Event containing the contract specifics and raw log

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
func (it *Extensionagglayerbridgel2NewWrappedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Extensionagglayerbridgel2NewWrappedToken)
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
		it.Event = new(Extensionagglayerbridgel2NewWrappedToken)
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
func (it *Extensionagglayerbridgel2NewWrappedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Extensionagglayerbridgel2NewWrappedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Extensionagglayerbridgel2NewWrappedToken represents a NewWrappedToken event raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2NewWrappedToken struct {
	OriginNetwork       uint32
	OriginTokenAddress  common.Address
	WrappedTokenAddress common.Address
	Metadata            []byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewWrappedToken is a free log retrieval operation binding the contract event 0x490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a39.
//
// Solidity: event NewWrappedToken(uint32 originNetwork, address originTokenAddress, address wrappedTokenAddress, bytes metadata)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) FilterNewWrappedToken(opts *bind.FilterOpts) (*Extensionagglayerbridgel2NewWrappedTokenIterator, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.FilterLogs(opts, "NewWrappedToken")
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2NewWrappedTokenIterator{contract: _Extensionagglayerbridgel2.contract, event: "NewWrappedToken", logs: logs, sub: sub}, nil
}

// WatchNewWrappedToken is a free log subscription operation binding the contract event 0x490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a39.
//
// Solidity: event NewWrappedToken(uint32 originNetwork, address originTokenAddress, address wrappedTokenAddress, bytes metadata)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) WatchNewWrappedToken(opts *bind.WatchOpts, sink chan<- *Extensionagglayerbridgel2NewWrappedToken) (event.Subscription, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.WatchLogs(opts, "NewWrappedToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Extensionagglayerbridgel2NewWrappedToken)
				if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "NewWrappedToken", log); err != nil {
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
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) ParseNewWrappedToken(log types.Log) (*Extensionagglayerbridgel2NewWrappedToken, error) {
	event := new(Extensionagglayerbridgel2NewWrappedToken)
	if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "NewWrappedToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Extensionagglayerbridgel2RemoveLegacySovereignTokenAddressIterator is returned from FilterRemoveLegacySovereignTokenAddress and is used to iterate over the raw logs and unpacked data for RemoveLegacySovereignTokenAddress events raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2RemoveLegacySovereignTokenAddressIterator struct {
	Event *Extensionagglayerbridgel2RemoveLegacySovereignTokenAddress // Event containing the contract specifics and raw log

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
func (it *Extensionagglayerbridgel2RemoveLegacySovereignTokenAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Extensionagglayerbridgel2RemoveLegacySovereignTokenAddress)
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
		it.Event = new(Extensionagglayerbridgel2RemoveLegacySovereignTokenAddress)
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
func (it *Extensionagglayerbridgel2RemoveLegacySovereignTokenAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Extensionagglayerbridgel2RemoveLegacySovereignTokenAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Extensionagglayerbridgel2RemoveLegacySovereignTokenAddress represents a RemoveLegacySovereignTokenAddress event raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2RemoveLegacySovereignTokenAddress struct {
	SovereignTokenAddress common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRemoveLegacySovereignTokenAddress is a free log retrieval operation binding the contract event 0xc2ae0bd0ec0fd0352bfe5bacac49637af342c1e40f1b80a7f74440dc7fe3f063.
//
// Solidity: event RemoveLegacySovereignTokenAddress(address sovereignTokenAddress)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) FilterRemoveLegacySovereignTokenAddress(opts *bind.FilterOpts) (*Extensionagglayerbridgel2RemoveLegacySovereignTokenAddressIterator, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.FilterLogs(opts, "RemoveLegacySovereignTokenAddress")
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2RemoveLegacySovereignTokenAddressIterator{contract: _Extensionagglayerbridgel2.contract, event: "RemoveLegacySovereignTokenAddress", logs: logs, sub: sub}, nil
}

// WatchRemoveLegacySovereignTokenAddress is a free log subscription operation binding the contract event 0xc2ae0bd0ec0fd0352bfe5bacac49637af342c1e40f1b80a7f74440dc7fe3f063.
//
// Solidity: event RemoveLegacySovereignTokenAddress(address sovereignTokenAddress)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) WatchRemoveLegacySovereignTokenAddress(opts *bind.WatchOpts, sink chan<- *Extensionagglayerbridgel2RemoveLegacySovereignTokenAddress) (event.Subscription, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.WatchLogs(opts, "RemoveLegacySovereignTokenAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Extensionagglayerbridgel2RemoveLegacySovereignTokenAddress)
				if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "RemoveLegacySovereignTokenAddress", log); err != nil {
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
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) ParseRemoveLegacySovereignTokenAddress(log types.Log) (*Extensionagglayerbridgel2RemoveLegacySovereignTokenAddress, error) {
	event := new(Extensionagglayerbridgel2RemoveLegacySovereignTokenAddress)
	if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "RemoveLegacySovereignTokenAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Extensionagglayerbridgel2SetBridgeManagerIterator is returned from FilterSetBridgeManager and is used to iterate over the raw logs and unpacked data for SetBridgeManager events raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2SetBridgeManagerIterator struct {
	Event *Extensionagglayerbridgel2SetBridgeManager // Event containing the contract specifics and raw log

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
func (it *Extensionagglayerbridgel2SetBridgeManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Extensionagglayerbridgel2SetBridgeManager)
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
		it.Event = new(Extensionagglayerbridgel2SetBridgeManager)
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
func (it *Extensionagglayerbridgel2SetBridgeManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Extensionagglayerbridgel2SetBridgeManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Extensionagglayerbridgel2SetBridgeManager represents a SetBridgeManager event raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2SetBridgeManager struct {
	BridgeManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetBridgeManager is a free log retrieval operation binding the contract event 0x32cf74f8a6d5f88593984d2cd52be5592bfa6884f5896175801a5069ef09cd67.
//
// Solidity: event SetBridgeManager(address bridgeManager)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) FilterSetBridgeManager(opts *bind.FilterOpts) (*Extensionagglayerbridgel2SetBridgeManagerIterator, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.FilterLogs(opts, "SetBridgeManager")
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2SetBridgeManagerIterator{contract: _Extensionagglayerbridgel2.contract, event: "SetBridgeManager", logs: logs, sub: sub}, nil
}

// WatchSetBridgeManager is a free log subscription operation binding the contract event 0x32cf74f8a6d5f88593984d2cd52be5592bfa6884f5896175801a5069ef09cd67.
//
// Solidity: event SetBridgeManager(address bridgeManager)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) WatchSetBridgeManager(opts *bind.WatchOpts, sink chan<- *Extensionagglayerbridgel2SetBridgeManager) (event.Subscription, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.WatchLogs(opts, "SetBridgeManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Extensionagglayerbridgel2SetBridgeManager)
				if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "SetBridgeManager", log); err != nil {
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
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) ParseSetBridgeManager(log types.Log) (*Extensionagglayerbridgel2SetBridgeManager, error) {
	event := new(Extensionagglayerbridgel2SetBridgeManager)
	if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "SetBridgeManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Extensionagglayerbridgel2SetClaimIterator is returned from FilterSetClaim and is used to iterate over the raw logs and unpacked data for SetClaim events raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2SetClaimIterator struct {
	Event *Extensionagglayerbridgel2SetClaim // Event containing the contract specifics and raw log

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
func (it *Extensionagglayerbridgel2SetClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Extensionagglayerbridgel2SetClaim)
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
		it.Event = new(Extensionagglayerbridgel2SetClaim)
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
func (it *Extensionagglayerbridgel2SetClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Extensionagglayerbridgel2SetClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Extensionagglayerbridgel2SetClaim represents a SetClaim event raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2SetClaim struct {
	GlobalIndex [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetClaim is a free log retrieval operation binding the contract event 0x8b65fb527138ff0fedc6b8e75e5a12268c71a1231b370b03209f7a286746823c.
//
// Solidity: event SetClaim(bytes32 globalIndex)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) FilterSetClaim(opts *bind.FilterOpts) (*Extensionagglayerbridgel2SetClaimIterator, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.FilterLogs(opts, "SetClaim")
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2SetClaimIterator{contract: _Extensionagglayerbridgel2.contract, event: "SetClaim", logs: logs, sub: sub}, nil
}

// WatchSetClaim is a free log subscription operation binding the contract event 0x8b65fb527138ff0fedc6b8e75e5a12268c71a1231b370b03209f7a286746823c.
//
// Solidity: event SetClaim(bytes32 globalIndex)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) WatchSetClaim(opts *bind.WatchOpts, sink chan<- *Extensionagglayerbridgel2SetClaim) (event.Subscription, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.WatchLogs(opts, "SetClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Extensionagglayerbridgel2SetClaim)
				if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "SetClaim", log); err != nil {
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

// ParseSetClaim is a log parse operation binding the contract event 0x8b65fb527138ff0fedc6b8e75e5a12268c71a1231b370b03209f7a286746823c.
//
// Solidity: event SetClaim(bytes32 globalIndex)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) ParseSetClaim(log types.Log) (*Extensionagglayerbridgel2SetClaim, error) {
	event := new(Extensionagglayerbridgel2SetClaim)
	if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "SetClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Extensionagglayerbridgel2SetLocalBalanceTreeIterator is returned from FilterSetLocalBalanceTree and is used to iterate over the raw logs and unpacked data for SetLocalBalanceTree events raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2SetLocalBalanceTreeIterator struct {
	Event *Extensionagglayerbridgel2SetLocalBalanceTree // Event containing the contract specifics and raw log

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
func (it *Extensionagglayerbridgel2SetLocalBalanceTreeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Extensionagglayerbridgel2SetLocalBalanceTree)
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
		it.Event = new(Extensionagglayerbridgel2SetLocalBalanceTree)
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
func (it *Extensionagglayerbridgel2SetLocalBalanceTreeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Extensionagglayerbridgel2SetLocalBalanceTreeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Extensionagglayerbridgel2SetLocalBalanceTree represents a SetLocalBalanceTree event raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2SetLocalBalanceTree struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
	NewAmount          *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSetLocalBalanceTree is a free log retrieval operation binding the contract event 0x73a03a6d9efc14b9f22a3af967e98580549eb76b1113b6a09a57ce1dae36ecd2.
//
// Solidity: event SetLocalBalanceTree(uint32 indexed originNetwork, address indexed originTokenAddress, uint256 newAmount)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) FilterSetLocalBalanceTree(opts *bind.FilterOpts, originNetwork []uint32, originTokenAddress []common.Address) (*Extensionagglayerbridgel2SetLocalBalanceTreeIterator, error) {

	var originNetworkRule []interface{}
	for _, originNetworkItem := range originNetwork {
		originNetworkRule = append(originNetworkRule, originNetworkItem)
	}
	var originTokenAddressRule []interface{}
	for _, originTokenAddressItem := range originTokenAddress {
		originTokenAddressRule = append(originTokenAddressRule, originTokenAddressItem)
	}

	logs, sub, err := _Extensionagglayerbridgel2.contract.FilterLogs(opts, "SetLocalBalanceTree", originNetworkRule, originTokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2SetLocalBalanceTreeIterator{contract: _Extensionagglayerbridgel2.contract, event: "SetLocalBalanceTree", logs: logs, sub: sub}, nil
}

// WatchSetLocalBalanceTree is a free log subscription operation binding the contract event 0x73a03a6d9efc14b9f22a3af967e98580549eb76b1113b6a09a57ce1dae36ecd2.
//
// Solidity: event SetLocalBalanceTree(uint32 indexed originNetwork, address indexed originTokenAddress, uint256 newAmount)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) WatchSetLocalBalanceTree(opts *bind.WatchOpts, sink chan<- *Extensionagglayerbridgel2SetLocalBalanceTree, originNetwork []uint32, originTokenAddress []common.Address) (event.Subscription, error) {

	var originNetworkRule []interface{}
	for _, originNetworkItem := range originNetwork {
		originNetworkRule = append(originNetworkRule, originNetworkItem)
	}
	var originTokenAddressRule []interface{}
	for _, originTokenAddressItem := range originTokenAddress {
		originTokenAddressRule = append(originTokenAddressRule, originTokenAddressItem)
	}

	logs, sub, err := _Extensionagglayerbridgel2.contract.WatchLogs(opts, "SetLocalBalanceTree", originNetworkRule, originTokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Extensionagglayerbridgel2SetLocalBalanceTree)
				if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "SetLocalBalanceTree", log); err != nil {
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
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) ParseSetLocalBalanceTree(log types.Log) (*Extensionagglayerbridgel2SetLocalBalanceTree, error) {
	event := new(Extensionagglayerbridgel2SetLocalBalanceTree)
	if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "SetLocalBalanceTree", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Extensionagglayerbridgel2SetSovereignTokenAddressIterator is returned from FilterSetSovereignTokenAddress and is used to iterate over the raw logs and unpacked data for SetSovereignTokenAddress events raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2SetSovereignTokenAddressIterator struct {
	Event *Extensionagglayerbridgel2SetSovereignTokenAddress // Event containing the contract specifics and raw log

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
func (it *Extensionagglayerbridgel2SetSovereignTokenAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Extensionagglayerbridgel2SetSovereignTokenAddress)
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
		it.Event = new(Extensionagglayerbridgel2SetSovereignTokenAddress)
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
func (it *Extensionagglayerbridgel2SetSovereignTokenAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Extensionagglayerbridgel2SetSovereignTokenAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Extensionagglayerbridgel2SetSovereignTokenAddress represents a SetSovereignTokenAddress event raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2SetSovereignTokenAddress struct {
	OriginNetwork         uint32
	OriginTokenAddress    common.Address
	SovereignTokenAddress common.Address
	IsNotMintable         bool
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetSovereignTokenAddress is a free log retrieval operation binding the contract event 0xdbe8a5da6a7a916d9adfda9160167a0f8a3da415ee6610e810e753853597fce7.
//
// Solidity: event SetSovereignTokenAddress(uint32 originNetwork, address originTokenAddress, address sovereignTokenAddress, bool isNotMintable)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) FilterSetSovereignTokenAddress(opts *bind.FilterOpts) (*Extensionagglayerbridgel2SetSovereignTokenAddressIterator, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.FilterLogs(opts, "SetSovereignTokenAddress")
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2SetSovereignTokenAddressIterator{contract: _Extensionagglayerbridgel2.contract, event: "SetSovereignTokenAddress", logs: logs, sub: sub}, nil
}

// WatchSetSovereignTokenAddress is a free log subscription operation binding the contract event 0xdbe8a5da6a7a916d9adfda9160167a0f8a3da415ee6610e810e753853597fce7.
//
// Solidity: event SetSovereignTokenAddress(uint32 originNetwork, address originTokenAddress, address sovereignTokenAddress, bool isNotMintable)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) WatchSetSovereignTokenAddress(opts *bind.WatchOpts, sink chan<- *Extensionagglayerbridgel2SetSovereignTokenAddress) (event.Subscription, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.WatchLogs(opts, "SetSovereignTokenAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Extensionagglayerbridgel2SetSovereignTokenAddress)
				if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "SetSovereignTokenAddress", log); err != nil {
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
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) ParseSetSovereignTokenAddress(log types.Log) (*Extensionagglayerbridgel2SetSovereignTokenAddress, error) {
	event := new(Extensionagglayerbridgel2SetSovereignTokenAddress)
	if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "SetSovereignTokenAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Extensionagglayerbridgel2SetSovereignWETHAddressIterator is returned from FilterSetSovereignWETHAddress and is used to iterate over the raw logs and unpacked data for SetSovereignWETHAddress events raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2SetSovereignWETHAddressIterator struct {
	Event *Extensionagglayerbridgel2SetSovereignWETHAddress // Event containing the contract specifics and raw log

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
func (it *Extensionagglayerbridgel2SetSovereignWETHAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Extensionagglayerbridgel2SetSovereignWETHAddress)
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
		it.Event = new(Extensionagglayerbridgel2SetSovereignWETHAddress)
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
func (it *Extensionagglayerbridgel2SetSovereignWETHAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Extensionagglayerbridgel2SetSovereignWETHAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Extensionagglayerbridgel2SetSovereignWETHAddress represents a SetSovereignWETHAddress event raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2SetSovereignWETHAddress struct {
	SovereignWETHTokenAddress common.Address
	IsNotMintable             bool
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterSetSovereignWETHAddress is a free log retrieval operation binding the contract event 0xc7318b7ed6ba4f2908a3de396d8ab49b1dadb55db5b55123247a401f29ff8d82.
//
// Solidity: event SetSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) FilterSetSovereignWETHAddress(opts *bind.FilterOpts) (*Extensionagglayerbridgel2SetSovereignWETHAddressIterator, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.FilterLogs(opts, "SetSovereignWETHAddress")
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2SetSovereignWETHAddressIterator{contract: _Extensionagglayerbridgel2.contract, event: "SetSovereignWETHAddress", logs: logs, sub: sub}, nil
}

// WatchSetSovereignWETHAddress is a free log subscription operation binding the contract event 0xc7318b7ed6ba4f2908a3de396d8ab49b1dadb55db5b55123247a401f29ff8d82.
//
// Solidity: event SetSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) WatchSetSovereignWETHAddress(opts *bind.WatchOpts, sink chan<- *Extensionagglayerbridgel2SetSovereignWETHAddress) (event.Subscription, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.WatchLogs(opts, "SetSovereignWETHAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Extensionagglayerbridgel2SetSovereignWETHAddress)
				if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "SetSovereignWETHAddress", log); err != nil {
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
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) ParseSetSovereignWETHAddress(log types.Log) (*Extensionagglayerbridgel2SetSovereignWETHAddress, error) {
	event := new(Extensionagglayerbridgel2SetSovereignWETHAddress)
	if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "SetSovereignWETHAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Extensionagglayerbridgel2TransferEmergencyBridgePauserRoleIterator is returned from FilterTransferEmergencyBridgePauserRole and is used to iterate over the raw logs and unpacked data for TransferEmergencyBridgePauserRole events raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2TransferEmergencyBridgePauserRoleIterator struct {
	Event *Extensionagglayerbridgel2TransferEmergencyBridgePauserRole // Event containing the contract specifics and raw log

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
func (it *Extensionagglayerbridgel2TransferEmergencyBridgePauserRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Extensionagglayerbridgel2TransferEmergencyBridgePauserRole)
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
		it.Event = new(Extensionagglayerbridgel2TransferEmergencyBridgePauserRole)
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
func (it *Extensionagglayerbridgel2TransferEmergencyBridgePauserRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Extensionagglayerbridgel2TransferEmergencyBridgePauserRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Extensionagglayerbridgel2TransferEmergencyBridgePauserRole represents a TransferEmergencyBridgePauserRole event raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2TransferEmergencyBridgePauserRole struct {
	CurrentEmergencyBridgePauser common.Address
	NewEmergencyBridgePauser     common.Address
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterTransferEmergencyBridgePauserRole is a free log retrieval operation binding the contract event 0xb27de219766f47b82684842855ba6130b6dbf288ac66d1c3509e7bf17f4e925a.
//
// Solidity: event TransferEmergencyBridgePauserRole(address currentEmergencyBridgePauser, address newEmergencyBridgePauser)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) FilterTransferEmergencyBridgePauserRole(opts *bind.FilterOpts) (*Extensionagglayerbridgel2TransferEmergencyBridgePauserRoleIterator, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.FilterLogs(opts, "TransferEmergencyBridgePauserRole")
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2TransferEmergencyBridgePauserRoleIterator{contract: _Extensionagglayerbridgel2.contract, event: "TransferEmergencyBridgePauserRole", logs: logs, sub: sub}, nil
}

// WatchTransferEmergencyBridgePauserRole is a free log subscription operation binding the contract event 0xb27de219766f47b82684842855ba6130b6dbf288ac66d1c3509e7bf17f4e925a.
//
// Solidity: event TransferEmergencyBridgePauserRole(address currentEmergencyBridgePauser, address newEmergencyBridgePauser)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) WatchTransferEmergencyBridgePauserRole(opts *bind.WatchOpts, sink chan<- *Extensionagglayerbridgel2TransferEmergencyBridgePauserRole) (event.Subscription, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.WatchLogs(opts, "TransferEmergencyBridgePauserRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Extensionagglayerbridgel2TransferEmergencyBridgePauserRole)
				if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "TransferEmergencyBridgePauserRole", log); err != nil {
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
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) ParseTransferEmergencyBridgePauserRole(log types.Log) (*Extensionagglayerbridgel2TransferEmergencyBridgePauserRole, error) {
	event := new(Extensionagglayerbridgel2TransferEmergencyBridgePauserRole)
	if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "TransferEmergencyBridgePauserRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Extensionagglayerbridgel2TransferEmergencyBridgeUnpauserRoleIterator is returned from FilterTransferEmergencyBridgeUnpauserRole and is used to iterate over the raw logs and unpacked data for TransferEmergencyBridgeUnpauserRole events raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2TransferEmergencyBridgeUnpauserRoleIterator struct {
	Event *Extensionagglayerbridgel2TransferEmergencyBridgeUnpauserRole // Event containing the contract specifics and raw log

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
func (it *Extensionagglayerbridgel2TransferEmergencyBridgeUnpauserRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Extensionagglayerbridgel2TransferEmergencyBridgeUnpauserRole)
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
		it.Event = new(Extensionagglayerbridgel2TransferEmergencyBridgeUnpauserRole)
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
func (it *Extensionagglayerbridgel2TransferEmergencyBridgeUnpauserRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Extensionagglayerbridgel2TransferEmergencyBridgeUnpauserRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Extensionagglayerbridgel2TransferEmergencyBridgeUnpauserRole represents a TransferEmergencyBridgeUnpauserRole event raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2TransferEmergencyBridgeUnpauserRole struct {
	CurrentEmergencyBridgeUnpauser common.Address
	NewEmergencyBridgeUnpauser     common.Address
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterTransferEmergencyBridgeUnpauserRole is a free log retrieval operation binding the contract event 0xf01a62a06940517bbc898dec8c75794b9feabcd2d263c8de823b36dbbeb8779b.
//
// Solidity: event TransferEmergencyBridgeUnpauserRole(address currentEmergencyBridgeUnpauser, address newEmergencyBridgeUnpauser)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) FilterTransferEmergencyBridgeUnpauserRole(opts *bind.FilterOpts) (*Extensionagglayerbridgel2TransferEmergencyBridgeUnpauserRoleIterator, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.FilterLogs(opts, "TransferEmergencyBridgeUnpauserRole")
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2TransferEmergencyBridgeUnpauserRoleIterator{contract: _Extensionagglayerbridgel2.contract, event: "TransferEmergencyBridgeUnpauserRole", logs: logs, sub: sub}, nil
}

// WatchTransferEmergencyBridgeUnpauserRole is a free log subscription operation binding the contract event 0xf01a62a06940517bbc898dec8c75794b9feabcd2d263c8de823b36dbbeb8779b.
//
// Solidity: event TransferEmergencyBridgeUnpauserRole(address currentEmergencyBridgeUnpauser, address newEmergencyBridgeUnpauser)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) WatchTransferEmergencyBridgeUnpauserRole(opts *bind.WatchOpts, sink chan<- *Extensionagglayerbridgel2TransferEmergencyBridgeUnpauserRole) (event.Subscription, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.WatchLogs(opts, "TransferEmergencyBridgeUnpauserRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Extensionagglayerbridgel2TransferEmergencyBridgeUnpauserRole)
				if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "TransferEmergencyBridgeUnpauserRole", log); err != nil {
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
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) ParseTransferEmergencyBridgeUnpauserRole(log types.Log) (*Extensionagglayerbridgel2TransferEmergencyBridgeUnpauserRole, error) {
	event := new(Extensionagglayerbridgel2TransferEmergencyBridgeUnpauserRole)
	if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "TransferEmergencyBridgeUnpauserRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Extensionagglayerbridgel2TransferProxiedTokensManagerRoleIterator is returned from FilterTransferProxiedTokensManagerRole and is used to iterate over the raw logs and unpacked data for TransferProxiedTokensManagerRole events raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2TransferProxiedTokensManagerRoleIterator struct {
	Event *Extensionagglayerbridgel2TransferProxiedTokensManagerRole // Event containing the contract specifics and raw log

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
func (it *Extensionagglayerbridgel2TransferProxiedTokensManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Extensionagglayerbridgel2TransferProxiedTokensManagerRole)
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
		it.Event = new(Extensionagglayerbridgel2TransferProxiedTokensManagerRole)
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
func (it *Extensionagglayerbridgel2TransferProxiedTokensManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Extensionagglayerbridgel2TransferProxiedTokensManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Extensionagglayerbridgel2TransferProxiedTokensManagerRole represents a TransferProxiedTokensManagerRole event raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2TransferProxiedTokensManagerRole struct {
	CurrentProxiedTokensManager common.Address
	NewProxiedTokensManager     common.Address
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterTransferProxiedTokensManagerRole is a free log retrieval operation binding the contract event 0x0a34baa3feb299aef9c05cb59c6e0c8e7c0bcc65cbf0a647e7a7c8a2411591e2.
//
// Solidity: event TransferProxiedTokensManagerRole(address currentProxiedTokensManager, address newProxiedTokensManager)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) FilterTransferProxiedTokensManagerRole(opts *bind.FilterOpts) (*Extensionagglayerbridgel2TransferProxiedTokensManagerRoleIterator, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.FilterLogs(opts, "TransferProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2TransferProxiedTokensManagerRoleIterator{contract: _Extensionagglayerbridgel2.contract, event: "TransferProxiedTokensManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferProxiedTokensManagerRole is a free log subscription operation binding the contract event 0x0a34baa3feb299aef9c05cb59c6e0c8e7c0bcc65cbf0a647e7a7c8a2411591e2.
//
// Solidity: event TransferProxiedTokensManagerRole(address currentProxiedTokensManager, address newProxiedTokensManager)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) WatchTransferProxiedTokensManagerRole(opts *bind.WatchOpts, sink chan<- *Extensionagglayerbridgel2TransferProxiedTokensManagerRole) (event.Subscription, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.WatchLogs(opts, "TransferProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Extensionagglayerbridgel2TransferProxiedTokensManagerRole)
				if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "TransferProxiedTokensManagerRole", log); err != nil {
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
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) ParseTransferProxiedTokensManagerRole(log types.Log) (*Extensionagglayerbridgel2TransferProxiedTokensManagerRole, error) {
	event := new(Extensionagglayerbridgel2TransferProxiedTokensManagerRole)
	if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "TransferProxiedTokensManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Extensionagglayerbridgel2UpdatedClaimedGlobalIndexHashChainIterator is returned from FilterUpdatedClaimedGlobalIndexHashChain and is used to iterate over the raw logs and unpacked data for UpdatedClaimedGlobalIndexHashChain events raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2UpdatedClaimedGlobalIndexHashChainIterator struct {
	Event *Extensionagglayerbridgel2UpdatedClaimedGlobalIndexHashChain // Event containing the contract specifics and raw log

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
func (it *Extensionagglayerbridgel2UpdatedClaimedGlobalIndexHashChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Extensionagglayerbridgel2UpdatedClaimedGlobalIndexHashChain)
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
		it.Event = new(Extensionagglayerbridgel2UpdatedClaimedGlobalIndexHashChain)
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
func (it *Extensionagglayerbridgel2UpdatedClaimedGlobalIndexHashChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Extensionagglayerbridgel2UpdatedClaimedGlobalIndexHashChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Extensionagglayerbridgel2UpdatedClaimedGlobalIndexHashChain represents a UpdatedClaimedGlobalIndexHashChain event raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2UpdatedClaimedGlobalIndexHashChain struct {
	ClaimedGlobalIndex             [32]byte
	NewClaimedGlobalIndexHashChain [32]byte
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterUpdatedClaimedGlobalIndexHashChain is a free log retrieval operation binding the contract event 0x3e5936f910a78eb5181813a939c8d4c3e4d85f87943f659380d82ac6221b0e92.
//
// Solidity: event UpdatedClaimedGlobalIndexHashChain(bytes32 claimedGlobalIndex, bytes32 newClaimedGlobalIndexHashChain)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) FilterUpdatedClaimedGlobalIndexHashChain(opts *bind.FilterOpts) (*Extensionagglayerbridgel2UpdatedClaimedGlobalIndexHashChainIterator, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.FilterLogs(opts, "UpdatedClaimedGlobalIndexHashChain")
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2UpdatedClaimedGlobalIndexHashChainIterator{contract: _Extensionagglayerbridgel2.contract, event: "UpdatedClaimedGlobalIndexHashChain", logs: logs, sub: sub}, nil
}

// WatchUpdatedClaimedGlobalIndexHashChain is a free log subscription operation binding the contract event 0x3e5936f910a78eb5181813a939c8d4c3e4d85f87943f659380d82ac6221b0e92.
//
// Solidity: event UpdatedClaimedGlobalIndexHashChain(bytes32 claimedGlobalIndex, bytes32 newClaimedGlobalIndexHashChain)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) WatchUpdatedClaimedGlobalIndexHashChain(opts *bind.WatchOpts, sink chan<- *Extensionagglayerbridgel2UpdatedClaimedGlobalIndexHashChain) (event.Subscription, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.WatchLogs(opts, "UpdatedClaimedGlobalIndexHashChain")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Extensionagglayerbridgel2UpdatedClaimedGlobalIndexHashChain)
				if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "UpdatedClaimedGlobalIndexHashChain", log); err != nil {
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
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) ParseUpdatedClaimedGlobalIndexHashChain(log types.Log) (*Extensionagglayerbridgel2UpdatedClaimedGlobalIndexHashChain, error) {
	event := new(Extensionagglayerbridgel2UpdatedClaimedGlobalIndexHashChain)
	if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "UpdatedClaimedGlobalIndexHashChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Extensionagglayerbridgel2UpdatedUnsetGlobalIndexHashChainIterator is returned from FilterUpdatedUnsetGlobalIndexHashChain and is used to iterate over the raw logs and unpacked data for UpdatedUnsetGlobalIndexHashChain events raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2UpdatedUnsetGlobalIndexHashChainIterator struct {
	Event *Extensionagglayerbridgel2UpdatedUnsetGlobalIndexHashChain // Event containing the contract specifics and raw log

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
func (it *Extensionagglayerbridgel2UpdatedUnsetGlobalIndexHashChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Extensionagglayerbridgel2UpdatedUnsetGlobalIndexHashChain)
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
		it.Event = new(Extensionagglayerbridgel2UpdatedUnsetGlobalIndexHashChain)
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
func (it *Extensionagglayerbridgel2UpdatedUnsetGlobalIndexHashChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Extensionagglayerbridgel2UpdatedUnsetGlobalIndexHashChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Extensionagglayerbridgel2UpdatedUnsetGlobalIndexHashChain represents a UpdatedUnsetGlobalIndexHashChain event raised by the Extensionagglayerbridgel2 contract.
type Extensionagglayerbridgel2UpdatedUnsetGlobalIndexHashChain struct {
	UnsetGlobalIndex             [32]byte
	NewUnsetGlobalIndexHashChain [32]byte
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedUnsetGlobalIndexHashChain is a free log retrieval operation binding the contract event 0xc80e0aca446a59735359a7ae46124b57c47b892827642779bc6dafc84ba90b03.
//
// Solidity: event UpdatedUnsetGlobalIndexHashChain(bytes32 unsetGlobalIndex, bytes32 newUnsetGlobalIndexHashChain)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) FilterUpdatedUnsetGlobalIndexHashChain(opts *bind.FilterOpts) (*Extensionagglayerbridgel2UpdatedUnsetGlobalIndexHashChainIterator, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.FilterLogs(opts, "UpdatedUnsetGlobalIndexHashChain")
	if err != nil {
		return nil, err
	}
	return &Extensionagglayerbridgel2UpdatedUnsetGlobalIndexHashChainIterator{contract: _Extensionagglayerbridgel2.contract, event: "UpdatedUnsetGlobalIndexHashChain", logs: logs, sub: sub}, nil
}

// WatchUpdatedUnsetGlobalIndexHashChain is a free log subscription operation binding the contract event 0xc80e0aca446a59735359a7ae46124b57c47b892827642779bc6dafc84ba90b03.
//
// Solidity: event UpdatedUnsetGlobalIndexHashChain(bytes32 unsetGlobalIndex, bytes32 newUnsetGlobalIndexHashChain)
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) WatchUpdatedUnsetGlobalIndexHashChain(opts *bind.WatchOpts, sink chan<- *Extensionagglayerbridgel2UpdatedUnsetGlobalIndexHashChain) (event.Subscription, error) {

	logs, sub, err := _Extensionagglayerbridgel2.contract.WatchLogs(opts, "UpdatedUnsetGlobalIndexHashChain")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Extensionagglayerbridgel2UpdatedUnsetGlobalIndexHashChain)
				if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "UpdatedUnsetGlobalIndexHashChain", log); err != nil {
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
func (_Extensionagglayerbridgel2 *Extensionagglayerbridgel2Filterer) ParseUpdatedUnsetGlobalIndexHashChain(log types.Log) (*Extensionagglayerbridgel2UpdatedUnsetGlobalIndexHashChain, error) {
	event := new(Extensionagglayerbridgel2UpdatedUnsetGlobalIndexHashChain)
	if err := _Extensionagglayerbridgel2.contract.UnpackLog(event, "UpdatedUnsetGlobalIndexHashChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

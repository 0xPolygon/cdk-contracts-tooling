// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package agglayerbridgel2frometrog

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

// Agglayerbridgel2frometrogMetaData contains all meta data concerning the Agglayerbridgel2frometrog contract.
var Agglayerbridgel2frometrogMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountDoesNotMatchMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeAddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DestinationNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmergencyStateNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EtherTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedProxyDeployment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputArraysLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDepositCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExpectedLER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGlobalIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLBTLeaf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLeafType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLeavesLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxyAdmin\",\"type\":\"address\"}],\"name\":\"InvalidProxyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSmtProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSovereignWETHAddressParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSubtreeFrontier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroNetworkID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxyAdmin\",\"type\":\"address\"}],\"name\":\"InvalidZeroProxyAdminOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"localBalanceTreeAmount\",\"type\":\"uint256\"}],\"name\":\"LocalBalanceTreeOverflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"localBalanceTreeAmount\",\"type\":\"uint256\"}],\"name\":\"LocalBalanceTreeUnderflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MsgValueNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeTokenIsEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewDepositCountExceedsMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoValueInMessagesOnGasTokenNetworks\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonZeroValueForUnusedFrontier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyBridgeManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyDeployer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyBridgePauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyBridgeUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootRemover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingEmergencyBridgePauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingEmergencyBridgeUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingProxiedTokensManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProxiedTokensManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OriginNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ParameterAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubtreeFrontierMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyMapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyUpdated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotMapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotRemapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WETHRemappingNotSupportedOnGasTokenNetworks\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldEmergencyBridgePauser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newEmergencyBridgePauser\",\"type\":\"address\"}],\"name\":\"AcceptEmergencyBridgePauserRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldEmergencyBridgeUnpauser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newEmergencyBridgeUnpauser\",\"type\":\"address\"}],\"name\":\"AcceptEmergencyBridgeUnpauserRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldProxiedTokensManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newProxiedTokensManager\",\"type\":\"address\"}],\"name\":\"AcceptProxiedTokensManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousDepositCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDepositCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"BackwardLET\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"depositCount\",\"type\":\"uint32\"}],\"name\":\"BridgeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"indexed\":false,\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"DetailedClaimEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousDepositCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDepositCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newLeaves\",\"type\":\"bytes\"}],\"name\":\"ForwardLET\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"legacyTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"updatedTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MigrateLegacyToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wrappedTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"NewWrappedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignTokenAddress\",\"type\":\"address\"}],\"name\":\"RemoveLegacySovereignTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridgeManager\",\"type\":\"address\"}],\"name\":\"SetBridgeManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"globalIndex\",\"type\":\"bytes32\"}],\"name\":\"SetClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"SetLocalBalanceTree\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"SetSovereignTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignWETHTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"SetSovereignWETHAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentEmergencyBridgePauser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newEmergencyBridgePauser\",\"type\":\"address\"}],\"name\":\"TransferEmergencyBridgePauserRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentEmergencyBridgeUnpauser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newEmergencyBridgeUnpauser\",\"type\":\"address\"}],\"name\":\"TransferEmergencyBridgeUnpauserRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentProxiedTokensManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newProxiedTokensManager\",\"type\":\"address\"}],\"name\":\"TransferProxiedTokensManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"claimedGlobalIndex\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newClaimedGlobalIndexHashChain\",\"type\":\"bytes32\"}],\"name\":\"UpdatedClaimedGlobalIndexHashChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"unsetGlobalIndex\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newUnsetGlobalIndexHashChain\",\"type\":\"bytes32\"}],\"name\":\"UpdatedUnsetGlobalIndexHashChain\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INIT_BYTECODE_TRANSPARENT_PROXY\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETHToken\",\"outputs\":[{\"internalType\":\"contractITokenWrappedBridgeUpgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptEmergencyBridgePauserRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptEmergencyBridgeUnpauserRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptProxiedTokensManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDepositCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[32]\",\"name\":\"newFrontier\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32\",\"name\":\"nextLeaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"proof\",\"type\":\"bytes32[32]\"}],\"name\":\"backwardLET\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"permitData\",\"type\":\"bytes\"}],\"name\":\"bridgeAsset\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeLib\",\"outputs\":[{\"internalType\":\"contractBridgeLib\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountWETH\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessageWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimedBitMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimedGlobalIndexHashChain\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"name\":\"computeTokenProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"deployWrappedTokenAndRemap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyBridgePauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyBridgeUnpauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structAgglayerBridgeL2.ClaimData[]\",\"name\":\"claims\",\"type\":\"tuple[]\"}],\"name\":\"forceEmitDetailedClaimEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structAgglayerBridgeL2.LeafData[]\",\"name\":\"newLeaves\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"expectedLER\",\"type\":\"bytes32\"}],\"name\":\"forwardLET\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProxiedTokensManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenWrappedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWrappedTokenBridgeImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIBaseLegacyAgglayerGER\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"contractIBaseLegacyAgglayerGER\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"contractIBaseLegacyAgglayerGER\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_emergencyBridgePauser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_emergencyBridgeUnpauser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proxiedTokensManager\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"wrappedTokensAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint128\",\"name\":\"initNativeSupply\",\"type\":\"uint128\"}],\"name\":\"initializeFromEtrog\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"leafIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"sourceBridgeNetwork\",\"type\":\"uint32\"}],\"name\":\"isClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdatedDepositCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tokenInfoHash\",\"type\":\"bytes32\"}],\"name\":\"localBalanceTree\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"legacyTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permitData\",\"type\":\"bytes\"}],\"name\":\"migrateLegacyToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingEmergencyBridgePauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingEmergencyBridgeUnpauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingProxiedTokensManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"polygonRollupManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"legacySovereignTokenAddress\",\"type\":\"address\"}],\"name\":\"removeLegacySovereignTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeManager\",\"type\":\"address\"}],\"name\":\"setBridgeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"originNetworkArray\",\"type\":\"uint32[]\"},{\"internalType\":\"address[]\",\"name\":\"originTokenAddressArray\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amountArray\",\"type\":\"uint256[]\"}],\"name\":\"setLocalBalanceTree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"globalIndexes\",\"type\":\"uint256[]\"}],\"name\":\"setMultipleClaims\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"originNetworks\",\"type\":\"uint32[]\"},{\"internalType\":\"address[]\",\"name\":\"originTokenAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"sovereignTokenAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"isNotMintable\",\"type\":\"bool[]\"}],\"name\":\"setMultipleSovereignTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sovereignWETHTokenAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"setSovereignWETHAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenInfoToWrappedToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newEmergencyBridgePauser\",\"type\":\"address\"}],\"name\":\"transferEmergencyBridgePauserRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newEmergencyBridgeUnpauser\",\"type\":\"address\"}],\"name\":\"transferEmergencyBridgeUnpauserRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newProxiedTokensManager\",\"type\":\"address\"}],\"name\":\"transferProxiedTokensManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unsetGlobalIndexHashChain\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"globalIndexes\",\"type\":\"uint256[]\"}],\"name\":\"unsetMultipleClaims\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrappedAddress\",\"type\":\"address\"}],\"name\":\"wrappedAddressIsNotMintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wrappedTokenToTokenInfo\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561000f575f5ffd5b5060405161001c9061014a565b604051809103905ff080158015610035573d5f5f3e3d5ffd5b506001600160a01b031660a05260405161004e90610157565b604051809103905ff080158015610067573d5f5f3e3d5ffd5b506001600160a01b031660805261007c61008d565b3360c05261008861008d565b610164565b5f54610100900460ff16156100f85760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161015610148575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6115818061604d83390190565b611bd2806175ce83390190565b60805160a05160c051615ea66101a75f395f50505f61050f01525f81816106ac01528181611f9e015281816120180152818161282e0152613b210152615ea65ff3fe6080604052600436106102bb575f3560e01c80626ee171146102bf57806303e6e116146102e0578063136a2c601461031a57806314cc01a01461033957806315064c961461035857806315419305146103815780631c208229146103a05780631d081d8c146103bf5780632072f6c5146103e257806322e95f2c146103f6578063240ff3781461041557806327aef4e8146104285780632dfdf0b5146104495780632f84c6901461045e578063318aee3d1461047d57806338b8fbbb146104e45780633b2fee9a146105015780633c351e10146105335780633cbc795b146105525780634b2f336d1461058a57806354fd4d50146105a9578063567f13e4146105d757806357cfbee3146105f65780635ca1e16514610615578063606617ff1461062957806369e3ab12146106485780636e974cd4146106675780636ee84b23146106865780636f0bc3da1461069b57806379e2cf97146106ce57806381b1c174146106e25780638b37b873146107165780638bd309c31461072a5780638c668f1c146107495780638d9420961461075d5780638ed7e3f21461077c5780638f9720bb1461079b578063ae24490a146107ba578063b0b37920146107d9578063b4586962146107f8578063b8b284d014610817578063bab161bf14610836578063be5831c714610857578063bf130d7f1461087a578063c00f14ab14610899578063c0f49163146108b8578063c514f24e146108e6578063cc461632146108fa578063ccaa2d1114610919578063cd58657914610938578063d02103ca1461094b578063d9cb3aec14610971578063dbc169761461099c578063e88f0436146109b0578063eabd372a146109c4578063ece93c6f146109e3578063ee25560b14610a02578063f0e7080814610a2d578063f214e16114610a4c578063f5efcd7914610a6b578063f811bff714610a8a578063fd7640e814610aa4575b5f5ffd5b3480156102ca575f5ffd5b506102de6102d9366004614867565b610ac3565b005b3480156102eb575f5ffd5b5060a8546103049061010090046001600160a01b031681565b6040516103119190614964565b60405180910390f35b348015610325575f5ffd5b506102de6103343660046149ff565b610b68565b348015610344575f5ffd5b5060a354610304906001600160a01b031681565b348015610363575f5ffd5b506068546103719060ff1681565b6040519015158152602001610311565b34801561038c575f5ffd5b506102de61039b366004614a9c565b610c9e565b3480156103ab575f5ffd5b506102de6103ba3660046149ff565b610f38565b3480156103ca575f5ffd5b506103d460a55481565b604051908152602001610311565b3480156103ed575f5ffd5b506102de611051565b348015610401575f5ffd5b50610304610410366004614b3c565b611086565b6102de610423366004614bb5565b6110d4565b348015610433575f5ffd5b5061043c611144565b6040516103119190614c76565b348015610454575f5ffd5b506103d460535481565b348015610469575f5ffd5b5060a454610304906001600160a01b031681565b348015610488575f5ffd5b506104c0610497366004614c88565b606b6020525f908152604090205463ffffffff811690600160201b90046001600160a01b031682565b6040805163ffffffff90931683526001600160a01b03909116602083015201610311565b3480156104ef575f5ffd5b506070546001600160a01b0316610304565b34801561050c575f5ffd5b507f0000000000000000000000000000000000000000000000000000000000000000610304565b34801561053e575f5ffd5b50606d54610304906001600160a01b031681565b34801561055d575f5ffd5b50606d5461057590600160a01b900463ffffffff1681565b60405163ffffffff9091168152602001610311565b348015610595575f5ffd5b50606f54610304906001600160a01b031681565b3480156105b4575f5ffd5b50604080518082019091526006815265076312e322e360d41b602082015261043c565b3480156105e2575f5ffd5b506102de6105f1366004614ce3565b6111d0565b348015610601575f5ffd5b506102de610610366004614d83565b611364565b348015610620575f5ffd5b506103d4611450565b348015610634575f5ffd5b5060aa54610304906001600160a01b031681565b348015610653575f5ffd5b506102de610662366004614c88565b6114c0565b348015610672575f5ffd5b506102de610681366004614e98565b611552565b348015610691575f5ffd5b506103d460a65481565b3480156106a6575f5ffd5b506103047f000000000000000000000000000000000000000000000000000000000000000081565b3480156106d9575f5ffd5b506102de6118b9565b3480156106ed575f5ffd5b506103046106fc366004614ede565b606a6020525f90815260409020546001600160a01b031681565b348015610721575f5ffd5b506102de6118da565b348015610735575f5ffd5b506102de610744366004614c88565b611950565b348015610754575f5ffd5b506102de6119d0565b348015610768575f5ffd5b506102de610777366004614c88565b611a46565b348015610787575f5ffd5b50606c54610304906001600160a01b031681565b3480156107a6575f5ffd5b506102de6107b5366004614ef5565b611ac6565b3480156107c5575f5ffd5b5060a954610304906001600160a01b031681565b3480156107e4575f5ffd5b506102de6107f3366004614f80565b611c1c565b348015610803575f5ffd5b506102de610812366004614c88565b611d70565b348015610822575f5ffd5b506102de610831366004614fd7565b611ed1565b348015610841575f5ffd5b5060685461057590610100900463ffffffff1681565b348015610862575f5ffd5b5060685461057590600160c81b900463ffffffff1681565b348015610885575f5ffd5b506102de610894366004615054565b611f4f565b3480156108a4575f5ffd5b5061043c6108b3366004614c88565b611f84565b3480156108c3575f5ffd5b506103716108d2366004614c88565b60a26020525f908152604090205460ff1681565b3480156108f1575f5ffd5b5061043c612014565b348015610905575f5ffd5b50610371610914366004615080565b61209d565b348015610924575f5ffd5b506102de6109333660046150c2565b6120ed565b6102de6109463660046151a0565b6124f7565b348015610956575f5ffd5b5060685461030490600160281b90046001600160a01b031681565b34801561097c575f5ffd5b506103d461098b366004614ede565b60a76020525f908152604090205481565b3480156109a7575f5ffd5b506102de61290d565b3480156109bb575f5ffd5b506102de612940565b3480156109cf575f5ffd5b506102de6109de366004614c88565b6129c7565b3480156109ee575f5ffd5b50607154610304906001600160a01b031681565b348015610a0d575f5ffd5b506103d4610a1c366004614ede565b60696020525f908152604090205481565b348015610a38575f5ffd5b506102de610a4736600461522f565b612a65565b348015610a57575f5ffd5b50610304610a66366004614b3c565b612d15565b348015610a76575f5ffd5b506102de610a853660046150c2565b612dba565b348015610a95575f5ffd5b506102de6102d9366004615276565b348015610aaf575f5ffd5b506102de610abe366004615305565b613006565b5f54610100900460ff1615808015610ae157505f54600160ff909116105b80610afa5750303b158015610afa57505f5460ff166001145b610b1f5760405162461bcd60e51b8152600401610b169061534d565b60405180910390fd5b5f805460ff191660011790558015610b40575f805461ff0019166101001790555b60405163f57ac68360e01b815260040160405180910390fd5b50505050505050505050505050565b606854604080516391eb796d60e01b815290513392600160281b90046001600160a01b0316916391eb796d9160048083019260209291908290030181865afa158015610bb6573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bda919061539b565b6001600160a01b031614610c015760405163a34ddeb160e01b815260040160405180910390fd5b5f5b8151811015610c9a575f828281518110610c1f57610c1f6153b6565b602002602001015190505f5f610c348361323f565b9250509150610c4382826132e2565b60a654610c509084613353565b60a68190556040805185815260208101929092527fc80e0aca446a59735359a7ae46124b57c47b892827642779bc6dafc84ba90b03910160405180910390a1505050600101610c03565b5050565b5f5460ff166071805460ff60a01b1916600160a01b60ff938416021790555f5460039161010090910416158015610cdb57505f5460ff8083169116105b610cf75760405162461bcd60e51b8152600401610b169061534d565b5f805461ffff191660ff80841691909117610100178255607154600160a01b9004169003610d385760405163f57ac68360e01b815260040160405180910390fd5b60a3546001600160a01b0316151580610d5b575060a4546001600160a01b031615155b80610d70575060a9546001600160a01b031615155b80610d8557506070546001600160a01b031615155b15610da35760405163761d7f9d60e11b815260040160405180910390fd5b60a380546001600160a01b03808a166001600160a01b03199283161790925560a4805492891692909116821790556040515f516020615e115f395f51905f5291610def915f91906153ca565b60405180910390a160a980546001600160a01b0319166001600160a01b0387169081179091556040515f516020615e515f395f51905f5291610e33915f91906153ca565b60405180910390a1306001600160a01b03851603610e6457604051631ae0e03360e01b815260040160405180910390fd5b6001600160a01b038416610e8b5760405163f6b2911f60e01b815260040160405180910390fd5b607080546001600160a01b0319166001600160a01b0386169081179091556040515f516020615df15f395f51905f5291610ec7915f91906153ca565b60405180910390a1610ee283836001600160801b0316613361565b5f805461ff001916905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150506071805460ff60a01b191690555050505050565b606854604080516391eb796d60e01b815290513392600160281b90046001600160a01b0316916391eb796d9160048083019260209291908290030181865afa158015610f86573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610faa919061539b565b6001600160a01b031614610fd15760405163a34ddeb160e01b815260040160405180910390fd5b5f5b8151811015610c9a575f828281518110610fef57610fef6153b6565b602002602001015190505f5f6110048361323f565b92505091506110138282613527565b6040518381527f8b65fb527138ff0fedc6b8e75e5a12268c71a1231b370b03209f7a286746823c9060200160405180910390a1505050600101610fd3565b60a4546001600160a01b0316331461107c57604051631344c5df60e11b815260040160405180910390fd5b611084613599565b565b5f606a5f848460405160200161109d9291906153e4565b60408051601f198184030181529181528151602092830120835290820192909252015f20546001600160a01b031690505b92915050565b60685460ff16156110f857604051630bc011ff60e21b815260040160405180910390fd5b34158015906111115750606f546001600160a01b031615155b1561112f576040516301bd897160e61b815260040160405180910390fd5b61113d8585348686866135f4565b5050505050565b606e80546111519061540e565b80601f016020809104026020016040519081016040528092919081815260200182805461117d9061540e565b80156111c85780601f1061119f576101008083540402835291602001916111c8565b820191905f5260205f20905b8154815290600101906020018083116111ab57829003601f168201915b505050505081565b606854604080516391eb796d60e01b815290513392600160281b90046001600160a01b0316916391eb796d9160048083019260209291908290030181865afa15801561121e573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611242919061539b565b6001600160a01b0316146112695760405163a34ddeb160e01b815260040160405180910390fd5b5f5b8181101561135f5736838383818110611286576112866153b6565b90506020028101906112989190615446565b90506112ac61090082016108e08301614c88565b6001600160a01b03166108008201355f516020615dd15f395f51905f528361040081016108208201356108408301356112ed6108808501610860860161547e565b6112ff6108a08a016108808b01615499565b6113116108c08b016108a08c01614c88565b6113236108e08c016108c08d01615499565b6109008c01356113376109208e018e6154b2565b60405161134e9b9a9998979695949392919061551c565b60405180910390a35060010161126b565b505050565b60a3546001600160a01b0316331461138f576040516357b738d160e11b815260040160405180910390fd5b825184511415806113a257508151845114155b806113af57508051845114155b156113cd5760405163434f49f560e11b815260040160405180910390fd5b5f5b825181101561113d576114488582815181106113ed576113ed6153b6565b6020026020010151858381518110611407576114076153b6565b6020026020010151858481518110611421576114216153b6565b602002602001015185858151811061143b5761143b6153b6565b60200260200101516136b5565b6001016113cf565b6053545f90819081805b60208110156114b7578083901c6001166001036114965761148f60338260208110611487576114876153b6565b015485613353565b93506114a3565b6114a08483613353565b93505b6114ad8283613353565b915060010161145a565b50919392505050565b60a4546001600160a01b031633146114eb57604051631344c5df60e11b815260040160405180910390fd5b60a880546001600160a01b0380841661010002610100600160a81b03199092169190911790915560a4546040517fb27de219766f47b82684842855ba6130b6dbf288ac66d1c3509e7bf17f4e925a9261154792169084906153ca565b60405180910390a150565b60a3546001600160a01b0316331461157d576040516357b738d160e11b815260040160405180910390fd5b6001600160a01b038216158015611598575063ffffffff8316155b1561173c575f61172a5f5f1b606f5f9054906101000a90046001600160a01b03166001600160a01b03166306fdde036040518163ffffffff1660e01b81526004015f60405180830381865afa1580156115f3573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261161a91908101906155ea565b606f5f9054906101000a90046001600160a01b03166001600160a01b03166395d89b416040518163ffffffff1660e01b81526004015f60405180830381865afa158015611669573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261169091908101906155ea565b606f5f9054906101000a90046001600160a01b03166001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa1580156116e0573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611704919061561b565b60405160200161171693929190615636565b604051602081830303815290604052613865565b90506117368183613927565b50505050565b5f83836040516020016117509291906153e4565b60408051601f1981840301815291815281516020928301205f818152606a9093529120549091506001600160a01b03168061179e5760405163828d566360e01b815260040160405180910390fd5b5f6118a383836001600160a01b03166306fdde036040518163ffffffff1660e01b81526004015f60405180830381865afa1580156117de573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261180591908101906155ea565b846001600160a01b03166395d89b416040518163ffffffff1660e01b81526004015f60405180830381865afa158015611840573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261186791908101906155ea565b856001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa1580156116e0573d5f5f3e3d5ffd5b90506118b1868683876136b5565b505050505050565b605354606854600160c81b900463ffffffff161015611084576110846139c6565b60aa546001600160a01b031633146119055760405163d491f0c160e01b815260040160405180910390fd5b60a9805460aa80546001600160a01b038082166001600160a01b03198086168217909655949091169091556040519116915f516020615e515f395f51905f52916115479184916153ca565b6070546001600160a01b0316331461197b57604051630866750360e01b815260040160405180910390fd5b607180546001600160a01b0319166001600160a01b03838116919091179091556070546040517f0a34baa3feb299aef9c05cb59c6e0c8e7c0bcc65cbf0a647e7a7c8a2411591e29261154792169084906153ca565b6071546001600160a01b031633146119fb57604051630b59ef2760e21b815260040160405180910390fd5b60708054607180546001600160a01b038082166001600160a01b03198086168217909655949091169091556040519116915f516020615df15f395f51905f52916115479184916153ca565b60a9546001600160a01b03163314611a7157604051638e9d821f60e01b815260040160405180910390fd5b60aa80546001600160a01b0319166001600160a01b038381169190911790915560a9546040517ff01a62a06940517bbc898dec8c75794b9feabcd2d263c8de823b36dbbeb8779b9261154792169084906153ca565b606854604080516391eb796d60e01b815290513392600160281b90046001600160a01b0316916391eb796d9160048083019260209291908290030181865afa158015611b14573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611b38919061539b565b6001600160a01b031614611b5f5760405163a34ddeb160e01b815260040160405180910390fd5b60685460ff16611b8257604051635386698160e01b815260040160405180910390fd5b81518351141580611b9557508051835114155b15611bb35760405163434f49f560e11b815260040160405180910390fd5b5f5b835181101561173657611c14848281518110611bd357611bd36153b6565b6020026020010151848381518110611bed57611bed6153b6565b6020026020010151848481518110611c0757611c076153b6565b6020026020010151613a54565b600101611bb5565b8015611c2d57611c2d848383613b0a565b6001600160a01b038085165f908152606b602090815260409182902082518084019093525463ffffffff81168352600160201b900490921691810182905290611c895760405163828d566360e01b815260040160405180910390fd5b5f606a5f835f01518460200151604051602001611ca79291906153e4565b60408051601f198184030181529181528151602092830120835290820192909252015f20546001600160a01b03908116915086168103611cfa5760405163e273c4a160e01b815260040160405180910390fd5b5f611d058787613b9e565b9050611d12823383613d27565b604080513381526001600160a01b03808a166020830152841691810191909152606081018290527fb7f8fd4d1faf9b2929dc269f59c53e3a2bccc44e9950f33a568fcbcb37eb69a9906080015b60405180910390a150505050505050565b60a3546001600160a01b03163314611d9b576040516357b738d160e11b815260040160405180910390fd5b6001600160a01b038082165f908152606b6020908152604080832081518083018352905463ffffffff8116808352600160201b90910490951681840181905291519094611deb93909291016153e4565b60408051601f1981840301815291815281516020928301205f818152606a9093529120549091506001600160a01b03161580611e3f57505f818152606a60205260409020546001600160a01b038481169116145b15611e5d5760405163e0c897a760e01b815260040160405180910390fd5b6001600160a01b0383165f908152606b6020908152604080832080546001600160c01b031916905560a290915290819020805460ff19169055517fc2ae0bd0ec0fd0352bfe5bacac49637af342c1e40f1b80a7f74440dc7fe3f06390611ec4908590614964565b60405180910390a1505050565b60685460ff1615611ef557604051630bc011ff60e21b815260040160405180910390fd5b606f546001600160a01b0316611f1e5760405163dde3cda760e01b815260040160405180910390fd5b606f545f90611f36906001600160a01b031686613b9e565b9050611f468787838787876135f4565b50505050505050565b60a3546001600160a01b03163314611f7a576040516357b738d160e11b815260040160405180910390fd5b610c9a8282613927565b60405163c00f14ab60e01b81526060906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063c00f14ab90611fd3908590600401614964565b5f60405180830381865afa158015611fed573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526110ce919081019061566e565b60607f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c514f24e6040518163ffffffff1660e01b81526004015f60405180830381865afa158015612071573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052612098919081019061566e565b905090565b5f806120b3600160201b63ffffffff85166156c6565b6120c39063ffffffff86166156dd565b600881901c5f90815260696020526040902054600160ff9092169190911b90811614949350505050565b60685460ff161561211157604051630bc011ff60e21b815260040160405180910390fd5b612119613db2565b60685463ffffffff8681166101009092041614612149576040516302caf51760e11b815260040160405180910390fd5b6121928c8c8c8c8c5f8d8d8d8d8d8d8d8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250613e0b92505050565b5f516020615e315f395f51905f528a888887876040516121b69594939291906156f0565b60405180910390a16001600160a01b0386161580156121d9575063ffffffff8716155b156122b757606f546001600160a01b031661229b575f6001600160a01b03851684825b6040519080825280601f01601f191660200182016040528015612226576020820181803683370190505b506040516122349190615722565b5f6040518083038185875af1925050503d805f811461226e576040519150601f19603f3d011682016040523d82523d5f602084013e612273565b606091505b505090508061229557604051630ce8f45160e31b815260040160405180910390fd5b506124e0565b606f546122b2906001600160a01b03168585613d27565b6124e0565b606d546001600160a01b0387811691161480156122e55750606d5463ffffffff888116600160a01b90920416145b156122fc575f6001600160a01b03851684826121fc565b60685463ffffffff610100909104811690881603612328576122b26001600160a01b0387168585613f12565b5f878760405160200161233c9291906153e4565b60408051601f1981840301815291815281516020928301205f818152606a9093529120549091506001600160a01b0316806124d2575f6123b18386868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061386592505050565b90506123be818888613d27565b80606a5f8581526020019081526020015f205f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555060405180604001604052808b63ffffffff1681526020018a6001600160a01b0316815250606b5f836001600160a01b03166001600160a01b031681526020019081526020015f205f820151815f015f6101000a81548163ffffffff021916908363ffffffff1602179055506020820151815f0160046101000a8154816001600160a01b0302191690836001600160a01b031602179055509050507f490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a398a8a8388886040516124c4959493929190615733565b60405180910390a1506124dd565b6124dd818787613d27565b50505b6124e960018055565b505050505050505050505050565b60685460ff161561251b57604051630bc011ff60e21b815260040160405180910390fd5b612523613db2565b60685463ffffffff610100909104811690881603612554576040516302caf51760e11b815260040160405180910390fd5b5f806060876001600160a01b038816612637578834146125875760405163b89240f560e01b815260040160405180910390fd5b606d54606e80546001600160a01b0383169650600160a01b90920463ffffffff169450906125b49061540e565b80601f01602080910402602001604051908101604052809291908181526020018280546125e09061540e565b801561262b5780601f106126025761010080835404028352916020019161262b565b820191905f5260205f20905b81548152906001019060200180831161260e57829003601f168201915b505050505091506128a8565b34156126565760405163798ee6f160e01b815260040160405180910390fd5b841561266757612667888787613b0a565b606f546001600160a01b039081169089160361268e57612687888a613b9e565b90506128a8565b6001600160a01b038089165f908152606b602090815260409182902082518084019093525463ffffffff81168352600160201b9004909216918101829052901515806126e05750805163ffffffff1615155b15612702576126ef898b613b9e565b6020820151825190965094509150612817565b6040516370a0823160e01b81525f906001600160a01b038b16906370a0823190612730903090600401614964565b602060405180830381865afa15801561274b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061276f9190615776565b90506127866001600160a01b038b1633308e613f6a565b6040516370a0823160e01b81525f906001600160a01b038c16906370a08231906127b4903090600401614964565b602060405180830381865afa1580156127cf573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906127f39190615776565b90506127ff828261578d565b6068548c9850610100900463ffffffff169650935050505b60405163c00f14ab60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063c00f14ab90612863908c90600401614964565b5f60405180830381865afa15801561287d573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526128a4919081019061566e565b9250505b5f516020615db15f395f51905f525f84868e8e86886053546040516128d49897969594939291906157a0565b60405180910390a16128f25f84868e8e868880519060200120613fa3565b8615612900576129006139c6565b50505050611f4660018055565b60a9546001600160a01b0316331461293857604051638e9d821f60e01b815260040160405180910390fd5b611084613fdb565b60a85461010090046001600160a01b0316331461297057604051637bb0100f60e01b815260040160405180910390fd5b60a4805460a880546001600160a01b03610100820481166001600160a01b031985168117909555610100600160a81b03199091169091556040519116915f516020615e115f395f51905f52916115479184916153ca565b60a3546001600160a01b031633146129f2576040516357b738d160e11b815260040160405180910390fd5b6001600160a01b038116612a195760405163f6b2911f60e01b815260040160405180910390fd5b60a380546001600160a01b0319166001600160a01b0383169081179091556040517f32cf74f8a6d5f88593984d2cd52be5592bfa6884f5896175801a5069ef09cd679161154791614964565b606854604080516391eb796d60e01b815290513392600160281b90046001600160a01b0316916391eb796d9160048083019260209291908290030181865afa158015612ab3573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612ad7919061539b565b6001600160a01b031614612afe5760405163a34ddeb160e01b815260040160405180910390fd5b60685460ff16612b2157604051635386698160e01b815260040160405180910390fd5b5f829003612b4257604051638488056760e01b815260040160405180910390fd5b6053545f612b4e611450565b90505f5b84811015612bfb575f868683818110612b6d57612b6d6153b6565b9050602002810190612b7f919061580c565b612b8890615820565b805190915060ff1615801590612ba35750805160ff16600114155b15612bc157604051632f162e1960e11b815260040160405180910390fd5b612bf2815f015182602001518360400151846060015185608001518660a001518760c0015180519060200120614032565b50600101612b52565b505f612c05611450565b9050838114612c275760405163277ffcf360e01b815260040160405180910390fd5b6053546068805463ffffffff60c81b1916600160c81b63ffffffff9093169290920291909117908190556040516333d6247d60e01b815260048101839052600160281b9091046001600160a01b0316906333d6247d906024015f604051808303815f87803b158015612c97575f5ffd5b505af1158015612ca9573d5f5f3e3d5ffd5b505050507fe90f3fcbcd86a778012608da0cd7622521c8dac9919dddf25a3fb634adb3d2c88383605354848a8a604051602001612ce7929190615902565b60408051601f1981840301815290829052612d059594939291615a10565b60405180910390a1505050505050565b5f5f8383604051602001612d2a9291906153e4565b6040516020818303038152906040528051906020012090505f60ff60f81b3083612d52612014565b8051602091820120604051612d9995949392016001600160f81b031994909416845260609290921b6001600160601b03191660018401526015830152603582015260550190565b60408051808303601f19018152919052805160209091012095945050505050565b60685460ff1615612dde57604051630bc011ff60e21b815260040160405180910390fd5b60685463ffffffff8681166101009092041614612e0e576040516302caf51760e11b815260040160405180910390fd5b612e588c8c8c8c8c60018d8d8d8d8d8d8d8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250613e0b92505050565b5f516020615e315f395f51905f528a88888787604051612e7c9594939291906156f0565b60405180910390a1606f545f906001600160a01b0316612f3757846001600160a01b031684888a8686604051602401612eb89493929190615a3a565b60408051601f198184030181529181526020820180516001600160e01b0316630c035af960e11b17905251612eed9190615722565b5f6040518083038185875af1925050503d805f8114612f27576040519150601f19603f3d011682016040523d82523d5f602084013e612f2c565b606091505b505080915050612fe8565b606f54612f4e906001600160a01b03168686613d27565b846001600160a01b031687898585604051602401612f6f9493929190615a3a565b60408051601f198184030181529181526020820180516001600160e01b0316630c035af960e11b17905251612fa49190615722565b5f604051808303815f865af19150503d805f8114612fdd576040519150601f19603f3d011682016040523d82523d5f602084013e612fe2565b606091505b50909150505b80610b59576040516337e391c360e01b815260040160405180910390fd5b606854604080516391eb796d60e01b815290513392600160281b90046001600160a01b0316916391eb796d9160048083019260209291908290030181865afa158015613054573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613078919061539b565b6001600160a01b03161461309f5760405163a34ddeb160e01b815260040160405180910390fd5b60685460ff166130c257604051635386698160e01b815260040160405180910390fd5b60535484106130e457604051631abd431560e31b815260040160405180910390fd5b6130f78282866130f2611450565b614049565b613114576040516338105f3b60e21b815260040160405180910390fd5b61311f848483614060565b6053545f61312b611450565b90505f5b602081101561316c5785816020811061314a5761314a6153b6565b602002013560338260208110613162576131626153b6565b015560010161312f565b5060538690555f61317b611450565b6068805463ffffffff60c81b1916600160c81b63ffffffff8b160217908190556040516333d6247d60e01b815260048101839052919250600160281b90046001600160a01b0316906333d6247d906024015f604051808303815f87803b1580156131e3575f5ffd5b505af11580156131f5573d5f5f3e3d5ffd5b505060408051868152602081018690529081018a9052606081018490527f2f8b2f36ca0d63af2efffe9d5e728a61c2bc273a3fc0574d12b558a41f14955592506080019050611d5f565b805f80600160401b83161561328b57505f9050808361326863ffffffff8516600160401b6156dd565b146132865760405163071389e960e01b815260040160405180910390fd5b6132db565b602084901c915061329d826001615a6a565b9050836132bd63ffffffff851663ffffffff60201b602086901b166156dd565b146132db5760405163071389e960e01b815260040160405180910390fd5b9193909250565b5f6132f7600160201b63ffffffff84166156c6565b6133079063ffffffff85166156dd565b600881901c5f8181526069602052604090208054600160ff851690811b918218928390559394509192919080821615611f4657604051630631b5f760e31b815260040160405180910390fd5b5f9182526020526040902090565b5f61336c478361578d565b606d549091506133939063ffffffff600160a01b820416906001600160a01b031683613a54565b606f546001600160a01b03161561341e57606f54604080516318160ddd60e01b815290515f926001600160a01b0316916318160ddd9160048083019260209291908290030181865afa1580156133eb573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061340f9190615776565b905061341c5f5f83613a54565b505b5f5b8351811015611736575f84828151811061343c5761343c6153b6565b6020908102919091018101516001600160a01b038082165f908152606b845260409081902081518083019092525463ffffffff81168252600160201b9004909116928101839052909250906134a45760405163828d566360e01b815260040160405180910390fd5b5f826001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156134e1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906135059190615776565b9050613519825f0151836020015183613a54565b505050806001019050613420565b5f61353c600160201b63ffffffff84166156c6565b61354c9063ffffffff85166156dd565b600881901c5f8181526069602052604081208054600160ff861690811b91821892839055949550929392918183169003611f4657604051630c8d9eab60e31b815260040160405180910390fd5b60685460ff16156135bd57604051630bc011ff60e21b815260040160405180910390fd5b6068805460ff191660011790556040517f2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497905f90a1565b60685463ffffffff610100909104811690871603613625576040516302caf51760e11b815260040160405180910390fd5b5f516020615db15f395f51905f526001606860019054906101000a900463ffffffff1633898989888860535460405161366699989796959493929190615a86565b60405180910390a16136a76001606860019054906101000a900463ffffffff1633898989888860405161369a929190615af4565b6040518091039020613fa3565b82156118b1576118b16139c6565b6001600160a01b03831615806136d257506001600160a01b038216155b156136f05760405163f6b2911f60e01b815260040160405180910390fd5b60685463ffffffff6101009091048116908516036137215760405163658b23ad60e01b815260040160405180910390fd5b6001600160a01b038281165f908152606b6020526040902054600160201b90041615613760576040516317abdeeb60e21b815260040160405180910390fd5b5f84846040516020016137749291906153e4565b60408051808303601f1901815282825280516020918201205f818152606a835283812080546001600160a01b0319166001600160a01b038a8116918217909255868601865263ffffffff8c81168089528c8416878a01818152848752606b89528987209a518b54915194166001600160c01b031990911617600160201b93909516929092029390931790975560a2855291859020805460ff191689151590811790915585519182529381019590955292840192909252606083015291507fdbe8a5da6a7a916d9adfda9160167a0f8a3da415ee6610e810e753853597fce79060800160405180910390a15050505050565b5f5f61386f612014565b9050838151602083015ff591506001600160a01b0382166138a3576040516331682e8d60e11b815260040160405180910390fd5b5f5f5f858060200190518101906138ba9190615b03565b925092509250846001600160a01b0316631624f6c68484846040518463ffffffff1660e01b81526004016138f093929190615636565b5f604051808303815f87803b158015613907575f5ffd5b505af1158015613919573d5f5f3e3d5ffd5b505050505050505092915050565b606d546001600160a01b031661395057604051634cb4711360e11b815260040160405180910390fd5b606f80546001600160a01b0319166001600160a01b0384169081179091555f81815260a26020908152604091829020805460ff19168515159081179091558251938452908301527fc7318b7ed6ba4f2908a3de396d8ab49b1dadb55db5b55123247a401f29ff8d82910160405180910390a15050565b6053546068805463ffffffff909216600160c81b0263ffffffff60c81b1990921691909117908190556001600160a01b03600160281b909104166333d6247d613a0d611450565b6040518263ffffffff1660e01b8152600401613a2b91815260200190565b5f604051808303815f87803b158015613a42575f5ffd5b505af1158015611736573d5f5f3e3d5ffd5b60685463ffffffff610100909104811690841603613a855760405163b869a63f60e01b815260040160405180910390fd5b5f8383604051602001613a999291906153e4565b60408051808303601f1901815282825280516020918201205f81815260a78352929092208590558483529092506001600160a01b0385169163ffffffff8716917f73a03a6d9efc14b9f22a3af967e98580549eb76b1113b6a09a57ce1dae36ecd2910160405180910390a350505050565b60405163a28fa4a360e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a28fa4a390613b5e9086908690869033903090600401615b70565b6020604051808303815f875af1158015613b7a573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117369190615bb6565b6001600160a01b0382165f90815260a2602052604081205460ff1615613cc4576040516370a0823160e01b81525f906001600160a01b038516906370a0823190613bec903090600401614964565b602060405180830381865afa158015613c07573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613c2b9190615776565b9050613c426001600160a01b038516333086613f6a565b6040516370a0823160e01b81525f906001600160a01b038616906370a0823190613c70903090600401614964565b602060405180830381865afa158015613c8b573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613caf9190615776565b9050613cbb828261578d565b925050506110ce565b604051632770a7eb60e21b81526001600160a01b03841690639dc29fac90613cf29033908690600401615bd1565b5f604051808303815f87803b158015613d09575f5ffd5b505af1158015613d1b573d5f5f3e3d5ffd5b505050508190506110ce565b6001600160a01b0383165f90815260a2602052604090205460ff1615613d5b5761135f6001600160a01b0384168383613f12565b6040516340c10f1960e01b81526001600160a01b038416906340c10f1990613d899085908590600401615bd1565b5f604051808303815f87803b158015613da0575f5ffd5b505af1158015611f46573d5f5f3e3d5ffd5b600260015403613e045760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610b16565b6002600155565b826001600160a01b03168a5f516020615dd15f395f51905f528e8e8d8d8d8d8d8d8c8c604051613e449a99989796959493929190615bea565b60405180910390a35f613e638888888888888880519060200120614173565b90505f5f613e758f8f8f8f8f886141ff565b91509150613e838282613527565b60a554613e9990613e948f86613353565b613353565b60a5819055604080518f815260208101929092527f3e5936f910a78eb5181813a939c8d4c3e4d85f87943f659380d82ac6221b0e92910160405180910390a160ff8a16613eeb57613eeb898987614338565b5f1960ff8b1601613f0157613f015f5f87614338565b505050505050505050505050505050565b61135f83846001600160a01b031663a9059cbb8585604051602401613f38929190615bd1565b604051602081830303815290604052915060e01b6020820180516001600160e01b0383818316178352505050506143f8565b6040516001600160a01b0384811660248301528381166044830152606482018390526117369186918216906323b872dd90608401613f38565b613fb287878787878787614032565b60ff8716613fc557613fc5868684614450565b5f1960ff881601611f4657611f465f5f84614450565b60685460ff16613ffe57604051635386698160e01b815260040160405180910390fd5b6068805460ff191690556040517f1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3905f90a1565b611f4661404488888888888888614173565b6144fa565b5f816140568686866145b0565b1495945050505050565b825f5b81158015906140725750602081105b1561412257816001166001036140d357828160208110614094576140946153b6565b60200201358482602081106140ab576140ab6153b6565b6020020135146140ce576040516379c13f8b60e01b815260040160405180910390fd5b614109565b5f8482602081106140e6576140e66153b6565b60200201351461410957604051630323f86f60e41b815260040160405180910390fd5b8061411381615c6a565b915050600182901c9150614063565b602081101561113d575f84826020811061413e5761413e6153b6565b60200201351461416157604051630323f86f60e41b815260040160405180910390fd5b8061416b81615c6a565b915050614122565b6040516001600160f81b031960f889901b1660208201526001600160e01b031960e088811b821660218401526001600160601b0319606089811b821660258601529188901b909216603984015285901b16603d82015260518101839052607181018290525f90609101604051602081830303815290604052805190602001209050979650505050505050565b6068545f9081908190600160281b90046001600160a01b031663257b3632614227888861462d565b6040518263ffffffff1660e01b815260040161424591815260200190565b6020604051808303815f875af1158015614261573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906142859190615776565b9050805f036142a657604051622f6fad60e01b815260040160405180910390fd5b5f5f5f6142b28a61323f565b925092509250600160401b8a165f146142f3576142d1878d858c614049565b6142ee576040516338105f3b60e21b815260040160405180910390fd5b614326565b614309614301888e866145b0565b8c848b614049565b614326576040516338105f3b60e21b815260040160405180910390fd5b919b919a509098505050505050505050565b60685463ffffffff61010090910481169084160361435557505050565b5f83836040516020016143699291906153e4565b60408051601f1981840301815291815281516020928301205f81815260a790935291205490915061439b905f1961578d565b8211156143d0575f81815260a76020526040908190205490516323d7213360e01b8152610b1691869186918691600401615c82565b5f81815260a76020526040812080548492906143ed9084906156dd565b909155505050505050565b5f61440c6001600160a01b03841683614638565b905080515f1415801561443057508080602001905181019061442e9190615bb6565b155b1561135f5782604051635274afe760e01b8152600401610b169190614964565b60685463ffffffff61010090910481169084160361446d57505050565b5f83836040516020016144819291906153e4565b60408051601f1981840301815291815281516020928301205f81815260a79093529120549091508211156144dd575f81815260a76020526040908190205490516314603c0160e01b8152610b1691869186918691600401615c82565b5f81815260a76020526040812080548492906143ed90849061578d565b80600161450960206002615d91565b614513919061578d565b60535410614534576040516377ae67b360e11b815260040160405180910390fd5b5f60535f815461454390615c6a565b918290555090505f5b60208110156145a7578082901c60011660010361457f578260338260208110614577576145776153b6565b015550505050565b61459d60338260208110614595576145956153b6565b015484613353565b925060010161454c565b5061135f615d9c565b5f83815b602081101561462257600163ffffffff8516821c811690036145f7576145f08582602081106145e5576145e56153b6565b602002013583613353565b915061461a565b6146178286836020811061460d5761460d6153b6565b6020020135613353565b91505b6001016145b4565b5090505b9392505050565b5f6146268383613353565b606061462683835f845f5f856001600160a01b0316848660405161465c9190615722565b5f6040518083038185875af1925050503d805f8114614696576040519150601f19603f3d011682016040523d82523d5f602084013e61469b565b606091505b50915091506146ab8683836146b5565b9695505050505050565b6060826146ca576146c582614708565b614626565b81511580156146e157506001600160a01b0384163b155b156147015783604051639996b31560e01b8152600401610b169190614964565b5080614626565b8051156147185780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b50565b803563ffffffff81168114614747575f5ffd5b919050565b6001600160a01b0381168114614731575f5ffd5b80356147478161474c565b634e487b7160e01b5f52604160045260245ffd5b60405160e081016001600160401b03811182821017156147a1576147a161476b565b60405290565b604051601f8201601f191681016001600160401b03811182821017156147cf576147cf61476b565b604052919050565b5f6001600160401b038211156147ef576147ef61476b565b50601f01601f191660200190565b5f82601f83011261480c575f5ffd5b813561481f61481a826147d7565b6147a7565b818152846020838601011115614833575f5ffd5b816020850160208301375f918101602001919091529392505050565b8015158114614731575f5ffd5b80356147478161484f565b5f5f5f5f5f5f5f5f5f5f5f5f6101808d8f031215614883575f5ffd5b61488c8d614734565b9b5061489a60208e01614760565b9a506148a860408e01614734565b99506148b660608e01614760565b98506148c460808e01614760565b97506001600160401b0360a08e013511156148dd575f5ffd5b6148ed8e60a08f01358f016147fd565b96506148fb60c08e01614760565b955061490960e08e01614760565b94506149186101008e0161485c565b93506149276101208e01614760565b92506149366101408e01614760565b91506149456101608e01614760565b90509295989b509295989b509295989b565b6001600160a01b03169052565b6001600160a01b0391909116815260200190565b5f6001600160401b038211156149905761499061476b565b5060051b60200190565b5f82601f8301126149a9575f5ffd5b81356149b761481a82614978565b8082825260208201915060208360051b8601019250858311156149d8575f5ffd5b602085015b838110156149f55780358352602092830192016149dd565b5095945050505050565b5f60208284031215614a0f575f5ffd5b81356001600160401b03811115614a24575f5ffd5b614a308482850161499a565b949350505050565b5f82601f830112614a47575f5ffd5b8135614a5561481a82614978565b8082825260208201915060208360051b860101925085831115614a76575f5ffd5b602085015b838110156149f5578035614a8e8161474c565b835260209283019201614a7b565b5f5f5f5f5f5f60c08789031215614ab1575f5ffd5b8635614abc8161474c565b95506020870135614acc8161474c565b94506040870135614adc8161474c565b93506060870135614aec8161474c565b925060808701356001600160401b03811115614b06575f5ffd5b614b1289828a01614a38565b92505060a08701356001600160801b0381168114614b2e575f5ffd5b809150509295509295509295565b5f5f60408385031215614b4d575f5ffd5b614b5683614734565b91506020830135614b668161474c565b809150509250929050565b5f5f83601f840112614b81575f5ffd5b5081356001600160401b03811115614b97575f5ffd5b602083019150836020828501011115614bae575f5ffd5b9250929050565b5f5f5f5f5f60808688031215614bc9575f5ffd5b614bd286614734565b94506020860135614be28161474c565b93506040860135614bf28161484f565b925060608601356001600160401b03811115614c0c575f5ffd5b614c1888828901614b71565b969995985093965092949392505050565b5f5b83811015614c43578181015183820152602001614c2b565b50505f910152565b5f8151808452614c62816020860160208601614c29565b601f01601f19169290920160200192915050565b602081525f6146266020830184614c4b565b5f60208284031215614c98575f5ffd5b81356146268161474c565b5f5f83601f840112614cb3575f5ffd5b5081356001600160401b03811115614cc9575f5ffd5b6020830191508360208260051b8501011115614bae575f5ffd5b5f5f60208385031215614cf4575f5ffd5b82356001600160401b03811115614d09575f5ffd5b614d1585828601614ca3565b90969095509350505050565b5f82601f830112614d30575f5ffd5b8135614d3e61481a82614978565b8082825260208201915060208360051b860101925085831115614d5f575f5ffd5b602085015b838110156149f557614d7581614734565b835260209283019201614d64565b5f5f5f5f60808587031215614d96575f5ffd5b84356001600160401b03811115614dab575f5ffd5b614db787828801614d21565b94505060208501356001600160401b03811115614dd2575f5ffd5b614dde87828801614a38565b93505060408501356001600160401b03811115614df9575f5ffd5b614e0587828801614a38565b92505060608501356001600160401b03811115614e20575f5ffd5b8501601f81018713614e30575f5ffd5b8035614e3e61481a82614978565b8082825260208201915060208360051b850101925089831115614e5f575f5ffd5b6020840193505b82841015614e8a578335614e798161484f565b825260209384019390910190614e66565b969995985093965050505050565b5f5f5f60608486031215614eaa575f5ffd5b614eb384614734565b92506020840135614ec38161474c565b91506040840135614ed38161484f565b809150509250925092565b5f60208284031215614eee575f5ffd5b5035919050565b5f5f5f60608486031215614f07575f5ffd5b83356001600160401b03811115614f1c575f5ffd5b614f2886828701614d21565b93505060208401356001600160401b03811115614f43575f5ffd5b614f4f86828701614a38565b92505060408401356001600160401b03811115614f6a575f5ffd5b614f768682870161499a565b9150509250925092565b5f5f5f5f60608587031215614f93575f5ffd5b8435614f9e8161474c565b93506020850135925060408501356001600160401b03811115614fbf575f5ffd5b614fcb87828801614b71565b95989497509550505050565b5f5f5f5f5f5f60a08789031215614fec575f5ffd5b614ff587614734565b955060208701356150058161474c565b945060408701359350606087013561501c8161484f565b925060808701356001600160401b03811115615036575f5ffd5b61504289828a01614b71565b979a9699509497509295939492505050565b5f5f60408385031215615065575f5ffd5b82356150708161474c565b91506020830135614b668161484f565b5f5f60408385031215615091575f5ffd5b61509a83614734565b91506150a860208401614734565b90509250929050565b8061040081018310156110ce575f5ffd5b5f5f5f5f5f5f5f5f5f5f5f5f6109208d8f0312156150de575f5ffd5b6150e88e8e6150b1565b9b506150f88e6104008f016150b1565b9a506108008d013599506108208d013598506108408d0135975061511f6108608e01614734565b965061512f6108808e013561474c565b6108808d013595506151446108a08e01614734565b94506108c08d01356151558161474c565b93506108e08d013592506001600160401b036109008e01351115615177575f5ffd5b6151888e6109008f01358f01614b71565b81935080925050509295989b509295989b509295989b565b5f5f5f5f5f5f5f60c0888a0312156151b6575f5ffd5b6151bf88614734565b965060208801356151cf8161474c565b95506040880135945060608801356151e68161474c565b935060808801356151f68161484f565b925060a08801356001600160401b03811115615210575f5ffd5b61521c8a828b01614b71565b989b979a50959850939692959293505050565b5f5f5f60408486031215615241575f5ffd5b83356001600160401b03811115615256575f5ffd5b61526286828701614ca3565b909790965060209590950135949350505050565b5f5f5f5f5f5f60c0878903121561528b575f5ffd5b61529487614734565b955060208701356152a48161474c565b94506152b260408801614734565b935060608701356152c28161474c565b925060808701356152d28161474c565b915060a08701356001600160401b038111156152ec575f5ffd5b6152f889828a016147fd565b9150509295509295509295565b5f5f5f5f6108408587031215615319575f5ffd5b8435935061532a86602087016150b1565b925061042085013591506153428661044087016150b1565b905092959194509250565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b5f602082840312156153ab575f5ffd5b81516146268161474c565b634e487b7160e01b5f52603260045260245ffd5b6001600160a01b0392831681529116602082015260400190565b60e09290921b6001600160e01b031916825260601b6001600160601b031916600482015260180190565b600181811c9082168061542257607f821691505b60208210810361544057634e487b7160e01b5f52602260045260245ffd5b50919050565b5f823561093e1983360301811261545b575f5ffd5b9190910192915050565b60ff81168114614731575f5ffd5b803561474781615465565b5f6020828403121561548e575f5ffd5b813561462681615465565b5f602082840312156154a9575f5ffd5b61462682614734565b5f5f8335601e198436030181126154c7575f5ffd5b8301803591506001600160401b038211156154e0575f5ffd5b602001915036819003821315614bae575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b6104008c82376104008b610400830137896108008201528861082082015260ff881661084082015263ffffffff871661086082015260018060a01b03861661088082015263ffffffff85166108a0820152836108c08201526109006108e08201525f61558d610900830184866154f4565b9d9c50505050505050505050505050565b5f6155ab61481a846147d7565b90508281528383830111156155be575f5ffd5b614626836020830184614c29565b5f82601f8301126155db575f5ffd5b6146268383516020850161559e565b5f602082840312156155fa575f5ffd5b81516001600160401b0381111561560f575f5ffd5b614a30848285016155cc565b5f6020828403121561562b575f5ffd5b815161462681615465565b606081525f6156486060830186614c4b565b828103602084015261565a8186614c4b565b91505060ff83166040830152949350505050565b5f6020828403121561567e575f5ffd5b81516001600160401b03811115615693575f5ffd5b8201601f810184136156a3575f5ffd5b614a308482516020840161559e565b634e487b7160e01b5f52601160045260245ffd5b80820281158282048414176110ce576110ce6156b2565b808201808211156110ce576110ce6156b2565b94855263ffffffff9390931660208501526001600160a01b039182166040850152166060830152608082015260a00190565b5f825161545b818460208701614c29565b63ffffffff861681526001600160a01b038581166020830152841660408201526080606082018190525f9061576b90830184866154f4565b979650505050505050565b5f60208284031215615786575f5ffd5b5051919050565b818103818111156110ce576110ce6156b2565b60ff8916815263ffffffff88811660208301526001600160a01b03888116604084015290871660608301528516608082015260a0810184905261010060c082018190525f906157f190830185614c4b565b905063ffffffff831660e08301529998505050505050505050565b5f823560de1983360301811261545b575f5ffd5b5f60e08236031215615830575f5ffd5b61583861477f565b61584183615473565b815261584f60208401614734565b602082015261586060408401614760565b604082015261587160608401614734565b606082015261588260808401614760565b608082015260a0838101359082015260c08301356001600160401b038111156158a9575f5ffd5b6158b5368286016147fd565b60c08301525092915050565b5f5f8335601e198436030181126158d6575f5ffd5b83016020810192503590506001600160401b038111156158f4575f5ffd5b803603821315614bae575f5ffd5b602080825281018290525f6040600584901b83018101908301858360de1936839003015b87821015615a0357868503603f190184528235818112615944575f5ffd5b8901803561595181615465565b60ff16865263ffffffff61596760208301614734565b166020870152604081013561597b8161474c565b6159886040880182614957565b5061599560608201614734565b63ffffffff1660608701526159ac60808201614760565b6159b96080880182614957565b5060a081810135908701526159d160c08201826158c1565b915060e060c08801526159e860e0880183836154f4565b96505050602083019250602084019350600182019150615926565b5092979650505050505050565b85815284602082015283604082015282606082015260a060808201525f61576b60a0830184614c4b565b6001600160a01b038516815263ffffffff841660208201526060604082018190525f906146ab90830184866154f4565b63ffffffff81811683821601908111156110ce576110ce6156b2565b60ff8a16815263ffffffff89811660208301526001600160a01b03898116604084015290881660608301528616608082015260a0810185905261010060c082018190525f90615ad890830185876154f4565b905063ffffffff831660e08301529a9950505050505050505050565b818382375f9101908152919050565b5f5f5f60608486031215615b15575f5ffd5b83516001600160401b03811115615b2a575f5ffd5b615b36868287016155cc565b602086015190945090506001600160401b03811115615b53575f5ffd5b615b5f868287016155cc565b9250506040840151614ed381615465565b6001600160a01b03861681526080602082018190525f90615b9490830186886154f4565b6001600160a01b03948516604084015292909316606090910152949350505050565b5f60208284031215615bc6575f5ffd5b81516146268161484f565b6001600160a01b03929092168252602082015260400190565b6104008b82376104008a610400830137886108008201528761082082015260ff871661084082015263ffffffff861661086082015260018060a01b03851661088082015263ffffffff84166108a0820152826108c08201526109006108e08201525f615c5a610900830184614c4b565b9c9b505050505050505050505050565b5f60018201615c7b57615c7b6156b2565b5060010190565b63ffffffff9490941684526001600160a01b039290921660208401526040830152606082015260800190565b6001815b6001841115615ce957808504811115615ccd57615ccd6156b2565b6001841615615cdb57908102905b60019390931c928002615cb2565b935093915050565b5f82615cff575060016110ce565b81615d0b57505f6110ce565b8160018114615d215760028114615d2b57615d47565b60019150506110ce565b60ff841115615d3c57615d3c6156b2565b50506001821b6110ce565b5060208310610133831016604e8410600b8410161715615d6a575081810a6110ce565b615d765f198484615cae565b805f1904821115615d8957615d896156b2565b029392505050565b5f6146268383615cf1565b634e487b7160e01b5f52600160045260245ffdfe501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9bfc09ab352fe828dc1fbb0cad35e812349ba00649f10067f83eafa8d79b161d81a9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f285d2bdfbe58cd81abf8199c13ce2509204be4aba8603b9d29f52c4e13e7bb7931df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d24cc8295aa5110cc216695db944ad2458c7795c6404449be980c3ce14aed752da264697066735822122030a059911df70236cf22665d852efb9bd37199b5b225d66628f86a21b79b7c5764736f6c634300081c00336080604052348015600e575f5ffd5b5060156019565b60c9565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560685760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c65780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6114ab806100d65f395ff3fe608060405234801561000f575f5ffd5b50600436106100d9575f3560e01c806306fdde03146100dd578063095ea7b3146100fb5780631624f6c61461011e57806318160ddd1461013357806323b872dd14610149578063313ce5671461015c5780633644e5151461017657806340c10f191461017e57806370a08231146101915780637ecebe00146101a457806384b0196e146101b757806395d89b41146101d25780639dc29fac146101da578063a3c573eb146101ed578063a9059cbb14610202578063d505accf14610215578063dd62ed3e14610228575b5f5ffd5b6100e561023b565b6040516100f29190610fbc565b60405180910390f35b61010e610109366004610ff0565b6102d9565b60405190151581526020016100f2565b61013161012c3660046110c5565b6102f2565b005b61013b610419565b6040519081526020016100f2565b61010e610157366004611137565b61042d565b610164610450565b60405160ff90911681526020016100f2565b61013b610464565b61013161018c366004610ff0565b610472565b61013b61019f366004611171565b6104bc565b61013b6101b2366004611171565b6104e5565b6101bf6104ef565b6040516100f2979695949392919061118a565b6100e5610598565b6101316101e8366004610ff0565b6105b4565b6101f56105f9565b6040516100f29190611220565b61010e610210366004610ff0565b610618565b610131610223366004611234565b610625565b61013b61023636600461129a565b610767565b60605f6102466107a1565b9050806003018054610257906112cb565b80601f0160208091040260200160405190810160405280929190818152602001828054610283906112cb565b80156102ce5780601f106102a5576101008083540402835291602001916102ce565b820191905f5260205f20905b8154815290600101906020018083116102b157829003601f168201915b505050505091505090565b5f336102e68185856107c5565b60019150505b92915050565b5f6102fb6107d2565b805490915060ff600160401b82041615906001600160401b03165f811580156103215750825b90505f826001600160401b0316600114801561033c5750303b155b90508115801561034a575080155b156103685760405163f92ee8a960e01b815260040160405180910390fd5b84546001600160401b0319166001178555831561039157845460ff60401b1916600160401b1785555b61039b88886107f6565b6103a48861080c565b5f6103ad61083a565b805433610100026001600160a81b031990911660ff8a161717905550831561040f57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b5f5f6104236107a1565b6002015492915050565b5f3361043a85828561085e565b6104458585856108ae565b506001949350505050565b5f5f61045a61083a565b5460ff1692915050565b5f61046d61090b565b905090565b5f61047b61083a565b805490915061010090046001600160a01b031633146104ad576040516338da3b1560e01b815260040160405180910390fd5b6104b78383610914565b505050565b5f5f6104c66107a1565b6001600160a01b039093165f9081526020939093525050604090205490565b5f6102ec82610948565b5f6060805f5f5f60605f610501610962565b805490915015801561051557506001810154155b61055e5760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b60448201526064015b60405180910390fd5b610566610986565b61056e6109a2565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b60605f6105a36107a1565b9050806004018054610257906112cb565b5f6105bd61083a565b805490915061010090046001600160a01b031633146105ef576040516338da3b1560e01b815260040160405180910390fd5b6104b783836109ad565b5f5f61060361083a565b5461010090046001600160a01b031692915050565b5f336102e68185856108ae565b834211156106495760405163313c898160e11b815260048101859052602401610555565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98888886106a08c6001600160a01b03165f9081525f5160206114565f395f51905f526020526040902080546001810190915590565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f6106fa826109e1565b90505f61070982878787610a0d565b9050896001600160a01b0316816001600160a01b031614610750576040516325c0072360e11b81526001600160a01b0380831660048301528b166024820152604401610555565b61075b8a8a8a6107c5565b50505050505050505050565b5f5f6107716107a1565b6001600160a01b039485165f90815260019190910160209081526040808320959096168252939093525050205490565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0090565b6104b78383836001610a39565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0090565b6107fe610b1a565b6108088282610b41565b5050565b610814610b1a565b61083781604051806040016040528060018152602001603160f81b815250610b71565b50565b7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d0090565b5f6108698484610767565b90505f1981146108a8578181101561089a57828183604051637dc7a0d960e11b815260040161055593929190611303565b6108a884848484035f610a39565b50505050565b6001600160a01b0383166108d7575f604051634b637e8f60e11b81526004016105559190611220565b6001600160a01b038216610900575f60405163ec442f0560e01b81526004016105559190611220565b6104b7838383610bb0565b5f61046d610cd3565b6001600160a01b03821661093d575f60405163ec442f0560e01b81526004016105559190611220565b6108085f8383610bb0565b5f6102ec825f805f5160206114565f395f51905f526104c6565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10090565b60605f610991610962565b9050806002018054610257906112cb565b60605f610246610962565b6001600160a01b0382166109d6575f604051634b637e8f60e11b81526004016105559190611220565b610808825f83610bb0565b5f6102ec6109ed61090b565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f610a1d88888888610d46565b925092509250610a2d8282610e04565b50909695505050505050565b5f610a426107a1565b90506001600160a01b038516610a6d575f60405163e602df0560e01b81526004016105559190611220565b6001600160a01b038416610a96575f604051634a1406b160e11b81526004016105559190611220565b6001600160a01b038086165f90815260018301602090815260408083209388168352929052208390558115610b1357836001600160a01b0316856001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92585604051610b0a91815260200190565b60405180910390a35b5050505050565b610b22610ebc565b610b3f57604051631afcd79f60e31b815260040160405180910390fd5b565b610b49610b1a565b5f610b526107a1565b905060038101610b628482611368565b50600481016108a88382611368565b610b79610b1a565b5f610b82610962565b905060028101610b928482611368565b5060038101610ba18382611368565b505f8082556001909101555050565b5f610bb96107a1565b90506001600160a01b038416610be75781816002015f828254610bdc9190611422565b90915550610c449050565b6001600160a01b0384165f9081526020829052604090205482811015610c265784818460405163391434e360e21b815260040161055593929190611303565b6001600160a01b0385165f9081526020839052604090209083900390555b6001600160a01b038316610c62576002810180548390039055610c80565b6001600160a01b0383165f9081526020829052604090208054830190555b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610cc591815260200190565b60405180910390a350505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f610cfd610ed5565b610d05610f3a565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f80806fa2a8918ca85bafe22016d0b997e4df60600160ff1b03841115610d7557505f91506003905082610dfa565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015610dc6573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116610df157505f925060019150829050610dfa565b92505f91508190505b9450945094915050565b5f826003811115610e1757610e17611441565b03610e20575050565b6001826003811115610e3457610e34611441565b03610e525760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115610e6657610e66611441565b03610e875760405163fce698f760e01b815260048101829052602401610555565b6003826003811115610e9b57610e9b611441565b03610808576040516335e2f38360e21b815260048101829052602401610555565b5f610ec56107d2565b54600160401b900460ff16919050565b5f5f610edf610962565b90505f610eea610986565b805190915015610f0257805160209091012092915050565b81548015610f11579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f610f44610962565b90505f610f4f6109a2565b805190915015610f6757805160209091012092915050565b60018201548015610f11579392505050565b5f81518084525f5b81811015610f9d57602081850181015186830182015201610f81565b505f602082860101526020601f19601f83011685010191505092915050565b602081525f610fce6020830184610f79565b9392505050565b80356001600160a01b0381168114610feb575f5ffd5b919050565b5f5f60408385031215611001575f5ffd5b61100a83610fd5565b946020939093013593505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f83011261103b575f5ffd5b81356001600160401b0381111561105457611054611018565b604051601f8201601f19908116603f011681016001600160401b038111828210171561108257611082611018565b604052818152838201602001851015611099575f5ffd5b816020850160208301375f918101602001919091529392505050565b803560ff81168114610feb575f5ffd5b5f5f5f606084860312156110d7575f5ffd5b83356001600160401b038111156110ec575f5ffd5b6110f88682870161102c565b93505060208401356001600160401b03811115611113575f5ffd5b61111f8682870161102c565b92505061112e604085016110b5565b90509250925092565b5f5f5f60608486031215611149575f5ffd5b61115284610fd5565b925061116060208501610fd5565b929592945050506040919091013590565b5f60208284031215611181575f5ffd5b610fce82610fd5565b60ff60f81b8816815260e060208201525f6111a860e0830189610f79565b82810360408401526111ba8189610f79565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b8181101561120f5783518352602093840193909201916001016111f1565b50909b9a5050505050505050505050565b6001600160a01b0391909116815260200190565b5f5f5f5f5f5f5f60e0888a03121561124a575f5ffd5b61125388610fd5565b965061126160208901610fd5565b9550604088013594506060880135935061127d608089016110b5565b9699959850939692959460a0840135945060c09093013592915050565b5f5f604083850312156112ab575f5ffd5b6112b483610fd5565b91506112c260208401610fd5565b90509250929050565b600181811c908216806112df57607f821691505b6020821081036112fd57634e487b7160e01b5f52602260045260245ffd5b50919050565b6001600160a01b039390931683526020830191909152604082015260600190565b601f8211156104b757805f5260205f20601f840160051c810160208510156113495750805b601f840160051c820191505b81811015610b13575f8155600101611355565b81516001600160401b0381111561138157611381611018565b6113958161138f84546112cb565b84611324565b6020601f8211600181146113c7575f83156113b05750848201515b5f19600385901b1c1916600184901b178455610b13565b5f84815260208120601f198516915b828110156113f657878501518255602094850194600190920191016113d6565b508482101561141357868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b808201808211156102ec57634e487b7160e01b5f52601160045260245ffd5b634e487b7160e01b5f52602160045260245ffdfe5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00a2646970667358221220367da95948f18cc3e9bdcda6e72786d5932d986c0478b78c09457cf82a1b465864736f6c634300081c00336080604052348015600e575f5ffd5b50611bb68061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610060575f3560e01c8063737ce34114610064578063a28fa4a31461008d578063be3dcf62146100b0578063c00f14ab146100d5578063c514f24e146100e8578063cf825e55146100f0575b5f5ffd5b6100776100723660046108dd565b610103565b604051610084919061094c565b60405180910390f35b6100a061009b36600461095e565b6101c8565b6040519015158152602001610084565b6100c36100be3660046108dd565b61056c565b60405160ff9091168152602001610084565b6100776100e33660046108dd565b610620565b610077610665565b6100776100fe3660046108dd565b610684565b60408051600481526024810182526020810180516001600160e01b03166395d89b4160e01b17905290516060915f9182916001600160a01b0386169161014991906109ff565b5f60405180830381855afa9150503d805f8114610181576040519150601f19603f3d011682016040523d82523d5f602084013e610186565b606091505b5091509150816101b757604051806040016040528060098152602001681393d7d4d6535093d360ba1b8152506101c0565b6101c081610736565b949350505050565b5f806101d76004828789610a1a565b6101e091610a41565b9050632afa533160e01b6001600160e01b031982160161038f575f5f5f5f5f5f5f8c8c600490809261021493929190610a1a565b8101906102219190610a87565b96509650965096509650965096508a6001600160a01b0316876001600160a01b0316146102615760405163912ecce760e01b815260040160405180910390fd5b896001600160a01b0316866001600160a01b0316146102935760405163750643af60e01b815260040160405180910390fd5b5f8e6001600160a01b031663d505accf60e01b898989898989896040516024016102ff97969594939291906001600160a01b0397881681529590961660208601526040850193909352606084019190915260ff16608083015260a082015260c081019190915260e00190565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b031990941693909317909252905161033d91906109ff565b5f604051808303815f865af19150503d805f8114610376576040519150601f19603f3d011682016040523d82523d5f602084013e61037b565b606091505b50909a506105639950505050505050505050565b631c0d143d60e21b6001600160e01b031982160161054a575f5f5f5f5f5f5f5f8d8d60049080926103c293929190610a1a565b8101906103cf9190610af3565b975097509750975097509750975097508b6001600160a01b0316886001600160a01b0316146104115760405163912ecce760e01b815260040160405180910390fd5b8a6001600160a01b0316876001600160a01b0316146104435760405163750643af60e01b815260040160405180910390fd5b5f8f6001600160a01b0316638fcbaf0c60e01b8a8a8a8a8a8a8a8a6040516024016104b99897969594939291906001600160a01b039889168152969097166020870152604086019490945260608501929092521515608084015260ff1660a083015260c082015260e08101919091526101000190565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516104f791906109ff565b5f604051808303815f865af19150503d805f8114610530576040519150601f19603f3d011682016040523d82523d5f602084013e610535565b606091505b50909b506105639a5050505050505050505050565b604051637141605d60e11b815260040160405180910390fd5b95945050505050565b60408051600481526024810182526020810180516001600160e01b031663313ce56760e01b17905290515f91829182916001600160a01b038616916105b191906109ff565b5f60405180830381855afa9150503d805f81146105e9576040519150601f19603f3d011682016040523d82523d5f602084013e6105ee565b606091505b5091509150818015610601575080516020145b61060c5760126101c0565b808060200190518101906101c09190610b75565b606061062b82610684565b61063483610103565b61063d8461056c565b60405160200161064f93929190610b90565b6040516020818303038152906040529050919050565b60405180610f000160405280610ec88152602001610cb9610ec8913981565b60408051600481526024810182526020810180516001600160e01b03166306fdde0360e01b17905290516060915f9182916001600160a01b038616916106ca91906109ff565b5f60405180830381855afa9150503d805f8114610702576040519150601f19603f3d011682016040523d82523d5f602084013e610707565b606091505b5091509150816101b757604051806040016040528060078152602001664e4f5f4e414d4560c81b8152506101c0565b6060604082511061075b57818060200190518101906107559190610bdc565b92915050565b8151602003610888575f5b602081108015610795575082818151811061078357610783610c80565b01602001516001600160f81b03191615155b156107ac57806107a481610c94565b915050610766565b805f036107e35750506040805180820190915260128152714e4f545f56414c49445f454e434f44494e4760701b6020820152919050565b5f816001600160401b038111156107fc576107fc610bc8565b6040519080825280601f01601f191660200182016040528015610826576020820181803683370190505b5090505f5b828110156108805784818151811061084557610845610c80565b602001015160f81c60f81b82828151811061086257610862610c80565b60200101906001600160f81b03191690815f1a90535060010161082b565b509392505050565b50506040805180820190915260128152714e4f545f56414c49445f454e434f44494e4760701b602082015290565b919050565b6001600160a01b03811681146108cf575f5ffd5b50565b80356108b6816108bb565b5f602082840312156108ed575f5ffd5b81356108f8816108bb565b9392505050565b5f5b83811015610919578181015183820152602001610901565b50505f910152565b5f81518084526109388160208601602086016108ff565b601f01601f19169290920160200192915050565b602081525f6108f86020830184610921565b5f5f5f5f5f60808688031215610972575f5ffd5b853561097d816108bb565b945060208601356001600160401b03811115610997575f5ffd5b8601601f810188136109a7575f5ffd5b80356001600160401b038111156109bc575f5ffd5b8860208284010111156109cd575f5ffd5b6020919091019450925060408601356109e5816108bb565b91506109f3606087016108d2565b90509295509295909350565b5f8251610a108184602087016108ff565b9190910192915050565b5f5f85851115610a28575f5ffd5b83861115610a34575f5ffd5b5050820193919092039150565b80356001600160e01b03198116906004841015610a72576001600160e01b0319600485900360031b81901b82161691505b5092915050565b60ff811681146108cf575f5ffd5b5f5f5f5f5f5f5f60e0888a031215610a9d575f5ffd5b8735610aa8816108bb565b96506020880135610ab8816108bb565b955060408801359450606088013593506080880135610ad681610a79565b9699959850939692959460a0840135945060c09093013592915050565b5f5f5f5f5f5f5f5f610100898b031215610b0b575f5ffd5b8835610b16816108bb565b97506020890135610b26816108bb565b9650604089013595506060890135945060808901358015158114610b48575f5ffd5b935060a0890135610b5881610a79565b979a969950949793969295929450505060c08201359160e0013590565b5f60208284031215610b85575f5ffd5b81516108f881610a79565b606081525f610ba26060830186610921565b8281036020840152610bb48186610921565b91505060ff83166040830152949350505050565b634e487b7160e01b5f52604160045260245ffd5b5f60208284031215610bec575f5ffd5b81516001600160401b03811115610c01575f5ffd5b8201601f81018413610c11575f5ffd5b80516001600160401b03811115610c2a57610c2a610bc8565b604051601f8201601f19908116603f011681016001600160401b0381118282101715610c5857610c58610bc8565b604052818152828201602001861015610c6f575f5ffd5b6105638260208301602086016108ff565b634e487b7160e01b5f52603260045260245ffd5b5f60018201610cb157634e487b7160e01b5f52601160045260245ffd5b506001019056fe60806040819052631d97f74d60e11b81523390633b2fee9a90608490602090600481865afa158015610033573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100579190610433565b604080515f80825260208201909252905061007382825f6100e2565b50506100dd336001600160a01b03166338b8fbbb6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100b4573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100d89190610433565b61010d565b6104c8565b6100eb8361017a565b5f825111806100f75750805b156101085761010683836101b9565b505b505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61014c5f516020610e815f395f51905f52546001600160a01b031690565b604080516001600160a01b03928316815291841660208301520160405180910390a1610177816101e5565b50565b61018381610280565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606101de8383604051806060016040528060278152602001610ea160279139610314565b9392505050565b6001600160a01b03811661024f5760405162461bcd60e51b815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b805f516020610e815f395f51905f525b80546001600160a01b0319166001600160a01b039290921691909117905550565b6001600160a01b0381163b6102ed5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610246565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc61025f565b60605f5f856001600160a01b031685604051610330919061047b565b5f60405180830381855af49150503d805f8114610368576040519150601f19603f3d011682016040523d82523d5f602084013e61036d565b606091505b50909250905061037f86838387610389565b9695505050505050565b606083156103f75782515f036103f0576001600160a01b0385163b6103f05760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610246565b5081610401565b6104018383610409565b949350505050565b8151156104195781518083602001fd5b8060405162461bcd60e51b81526004016102469190610496565b5f60208284031215610443575f5ffd5b81516001600160a01b03811681146101de575f5ffd5b5f5b8381101561047357818101518382015260200161045b565b50505f910152565b5f825161048c818460208701610459565b9190910192915050565b602081525f82518060208401526104b4816040850160208701610459565b601f01601f19169190910160400192915050565b6109ac806104d55f395ff3fe60806040526004361061005d575f3560e01c80635c60da1b116100425780635c60da1b146100a65780638f283970146100e3578063f851a440146101025761006c565b80633659cfe6146100745780634f1ef286146100935761006c565b3661006c5761006a610116565b005b61006a610116565b34801561007f575f5ffd5b5061006a61008e366004610854565b610130565b61006a6100a136600461086d565b610178565b3480156100b1575f5ffd5b506100ba6101eb565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b3480156100ee575f5ffd5b5061006a6100fd366004610854565b610228565b34801561010d575f5ffd5b506100ba610255565b61011e610282565b61012e610129610359565b610362565b565b610138610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101705761016d8160405180602001604052805f8152505f6103bf565b50565b61016d610116565b610180610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101e3576101de8383838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250600192506103bf915050565b505050565b6101de610116565b5f6101f4610380565b73ffffffffffffffffffffffffffffffffffffffff16330361021d57610218610359565b905090565b610225610116565b90565b610230610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101705761016d816103e9565b5f61025e610380565b73ffffffffffffffffffffffffffffffffffffffff16330361021d57610218610380565b61028a610380565b73ffffffffffffffffffffffffffffffffffffffff16330361012e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604260248201527f5472616e73706172656e745570677261646561626c6550726f78793a2061646d60448201527f696e2063616e6e6f742066616c6c6261636b20746f2070726f7879207461726760648201527f6574000000000000000000000000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b5f61021861044a565b365f5f375f5f365f845af43d5f5f3e80801561037c573d5ff35b3d5ffd5b5f7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b5473ffffffffffffffffffffffffffffffffffffffff16919050565b6103c883610471565b5f825111806103d45750805b156101de576103e383836104bd565b50505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f610412610380565b6040805173ffffffffffffffffffffffffffffffffffffffff928316815291841660208301520160405180910390a161016d816104e9565b5f7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6103a3565b61047a816105f5565b60405173ffffffffffffffffffffffffffffffffffffffff8216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606104e28383604051806060016040528060278152602001610979602791396106c0565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff811661058c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610350565b807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b73ffffffffffffffffffffffffffffffffffffffff81163b610699576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201527f6f74206120636f6e7472616374000000000000000000000000000000000000006064820152608401610350565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6105af565b60605f5f8573ffffffffffffffffffffffffffffffffffffffff16856040516106e9919061090d565b5f60405180830381855af49150503d805f8114610721576040519150601f19603f3d011682016040523d82523d5f602084013e610726565b606091505b509150915061073786838387610741565b9695505050505050565b606083156107d65782515f036107cf5773ffffffffffffffffffffffffffffffffffffffff85163b6107cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610350565b50816107e0565b6107e083836107e8565b949350505050565b8151156107f85781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103509190610928565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f575f5ffd5b919050565b5f60208284031215610864575f5ffd5b6104e28261082c565b5f5f5f6040848603121561087f575f5ffd5b6108888461082c565b9250602084013567ffffffffffffffff8111156108a3575f5ffd5b8401601f810186136108b3575f5ffd5b803567ffffffffffffffff8111156108c9575f5ffd5b8660208284010111156108da575f5ffd5b939660209190910195509293505050565b5f5b838110156109055781810151838201526020016108ed565b50505f910152565b5f825161091e8184602087016108eb565b9190910192915050565b602081525f82518060208401526109468160408501602087016108eb565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a164736f6c634300081c000ab53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220fc98387674097593539459356c434fa5891415aa5e617f46280959010be2031964736f6c634300081c0033",
}

// Agglayerbridgel2frometrogABI is the input ABI used to generate the binding from.
// Deprecated: Use Agglayerbridgel2frometrogMetaData.ABI instead.
var Agglayerbridgel2frometrogABI = Agglayerbridgel2frometrogMetaData.ABI

// Agglayerbridgel2frometrogBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Agglayerbridgel2frometrogMetaData.Bin instead.
var Agglayerbridgel2frometrogBin = Agglayerbridgel2frometrogMetaData.Bin

// DeployAgglayerbridgel2frometrog deploys a new Ethereum contract, binding an instance of Agglayerbridgel2frometrog to it.
func DeployAgglayerbridgel2frometrog(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Agglayerbridgel2frometrog, error) {
	parsed, err := Agglayerbridgel2frometrogMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Agglayerbridgel2frometrogBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Agglayerbridgel2frometrog{Agglayerbridgel2frometrogCaller: Agglayerbridgel2frometrogCaller{contract: contract}, Agglayerbridgel2frometrogTransactor: Agglayerbridgel2frometrogTransactor{contract: contract}, Agglayerbridgel2frometrogFilterer: Agglayerbridgel2frometrogFilterer{contract: contract}}, nil
}

// Agglayerbridgel2frometrog is an auto generated Go binding around an Ethereum contract.
type Agglayerbridgel2frometrog struct {
	Agglayerbridgel2frometrogCaller     // Read-only binding to the contract
	Agglayerbridgel2frometrogTransactor // Write-only binding to the contract
	Agglayerbridgel2frometrogFilterer   // Log filterer for contract events
}

// Agglayerbridgel2frometrogCaller is an auto generated read-only Go binding around an Ethereum contract.
type Agglayerbridgel2frometrogCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Agglayerbridgel2frometrogTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Agglayerbridgel2frometrogTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Agglayerbridgel2frometrogFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Agglayerbridgel2frometrogFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Agglayerbridgel2frometrogSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Agglayerbridgel2frometrogSession struct {
	Contract     *Agglayerbridgel2frometrog // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// Agglayerbridgel2frometrogCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Agglayerbridgel2frometrogCallerSession struct {
	Contract *Agglayerbridgel2frometrogCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// Agglayerbridgel2frometrogTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Agglayerbridgel2frometrogTransactorSession struct {
	Contract     *Agglayerbridgel2frometrogTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// Agglayerbridgel2frometrogRaw is an auto generated low-level Go binding around an Ethereum contract.
type Agglayerbridgel2frometrogRaw struct {
	Contract *Agglayerbridgel2frometrog // Generic contract binding to access the raw methods on
}

// Agglayerbridgel2frometrogCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Agglayerbridgel2frometrogCallerRaw struct {
	Contract *Agglayerbridgel2frometrogCaller // Generic read-only contract binding to access the raw methods on
}

// Agglayerbridgel2frometrogTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Agglayerbridgel2frometrogTransactorRaw struct {
	Contract *Agglayerbridgel2frometrogTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgglayerbridgel2frometrog creates a new instance of Agglayerbridgel2frometrog, bound to a specific deployed contract.
func NewAgglayerbridgel2frometrog(address common.Address, backend bind.ContractBackend) (*Agglayerbridgel2frometrog, error) {
	contract, err := bindAgglayerbridgel2frometrog(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrog{Agglayerbridgel2frometrogCaller: Agglayerbridgel2frometrogCaller{contract: contract}, Agglayerbridgel2frometrogTransactor: Agglayerbridgel2frometrogTransactor{contract: contract}, Agglayerbridgel2frometrogFilterer: Agglayerbridgel2frometrogFilterer{contract: contract}}, nil
}

// NewAgglayerbridgel2frometrogCaller creates a new read-only instance of Agglayerbridgel2frometrog, bound to a specific deployed contract.
func NewAgglayerbridgel2frometrogCaller(address common.Address, caller bind.ContractCaller) (*Agglayerbridgel2frometrogCaller, error) {
	contract, err := bindAgglayerbridgel2frometrog(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogCaller{contract: contract}, nil
}

// NewAgglayerbridgel2frometrogTransactor creates a new write-only instance of Agglayerbridgel2frometrog, bound to a specific deployed contract.
func NewAgglayerbridgel2frometrogTransactor(address common.Address, transactor bind.ContractTransactor) (*Agglayerbridgel2frometrogTransactor, error) {
	contract, err := bindAgglayerbridgel2frometrog(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogTransactor{contract: contract}, nil
}

// NewAgglayerbridgel2frometrogFilterer creates a new log filterer instance of Agglayerbridgel2frometrog, bound to a specific deployed contract.
func NewAgglayerbridgel2frometrogFilterer(address common.Address, filterer bind.ContractFilterer) (*Agglayerbridgel2frometrogFilterer, error) {
	contract, err := bindAgglayerbridgel2frometrog(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogFilterer{contract: contract}, nil
}

// bindAgglayerbridgel2frometrog binds a generic wrapper to an already deployed contract.
func bindAgglayerbridgel2frometrog(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Agglayerbridgel2frometrogMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Agglayerbridgel2frometrog.Contract.Agglayerbridgel2frometrogCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.Agglayerbridgel2frometrogTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.Agglayerbridgel2frometrogTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Agglayerbridgel2frometrog.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.contract.Transact(opts, method, params...)
}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() view returns(bytes)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) INITBYTECODETRANSPARENTPROXY(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "INIT_BYTECODE_TRANSPARENT_PROXY")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() view returns(bytes)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) INITBYTECODETRANSPARENTPROXY() ([]byte, error) {
	return _Agglayerbridgel2frometrog.Contract.INITBYTECODETRANSPARENTPROXY(&_Agglayerbridgel2frometrog.CallOpts)
}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() view returns(bytes)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) INITBYTECODETRANSPARENTPROXY() ([]byte, error) {
	return _Agglayerbridgel2frometrog.Contract.INITBYTECODETRANSPARENTPROXY(&_Agglayerbridgel2frometrog.CallOpts)
}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) WETHToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "WETHToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) WETHToken() (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.WETHToken(&_Agglayerbridgel2frometrog.CallOpts)
}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) WETHToken() (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.WETHToken(&_Agglayerbridgel2frometrog.CallOpts)
}

// BridgeLib is a free data retrieval call binding the contract method 0x6f0bc3da.
//
// Solidity: function bridgeLib() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) BridgeLib(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "bridgeLib")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeLib is a free data retrieval call binding the contract method 0x6f0bc3da.
//
// Solidity: function bridgeLib() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) BridgeLib() (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.BridgeLib(&_Agglayerbridgel2frometrog.CallOpts)
}

// BridgeLib is a free data retrieval call binding the contract method 0x6f0bc3da.
//
// Solidity: function bridgeLib() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) BridgeLib() (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.BridgeLib(&_Agglayerbridgel2frometrog.CallOpts)
}

// BridgeManager is a free data retrieval call binding the contract method 0x14cc01a0.
//
// Solidity: function bridgeManager() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) BridgeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "bridgeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeManager is a free data retrieval call binding the contract method 0x14cc01a0.
//
// Solidity: function bridgeManager() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) BridgeManager() (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.BridgeManager(&_Agglayerbridgel2frometrog.CallOpts)
}

// BridgeManager is a free data retrieval call binding the contract method 0x14cc01a0.
//
// Solidity: function bridgeManager() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) BridgeManager() (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.BridgeManager(&_Agglayerbridgel2frometrog.CallOpts)
}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) ClaimedBitMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "claimedBitMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) ClaimedBitMap(arg0 *big.Int) (*big.Int, error) {
	return _Agglayerbridgel2frometrog.Contract.ClaimedBitMap(&_Agglayerbridgel2frometrog.CallOpts, arg0)
}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) ClaimedBitMap(arg0 *big.Int) (*big.Int, error) {
	return _Agglayerbridgel2frometrog.Contract.ClaimedBitMap(&_Agglayerbridgel2frometrog.CallOpts, arg0)
}

// ClaimedGlobalIndexHashChain is a free data retrieval call binding the contract method 0x1d081d8c.
//
// Solidity: function claimedGlobalIndexHashChain() view returns(bytes32)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) ClaimedGlobalIndexHashChain(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "claimedGlobalIndexHashChain")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ClaimedGlobalIndexHashChain is a free data retrieval call binding the contract method 0x1d081d8c.
//
// Solidity: function claimedGlobalIndexHashChain() view returns(bytes32)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) ClaimedGlobalIndexHashChain() ([32]byte, error) {
	return _Agglayerbridgel2frometrog.Contract.ClaimedGlobalIndexHashChain(&_Agglayerbridgel2frometrog.CallOpts)
}

// ClaimedGlobalIndexHashChain is a free data retrieval call binding the contract method 0x1d081d8c.
//
// Solidity: function claimedGlobalIndexHashChain() view returns(bytes32)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) ClaimedGlobalIndexHashChain() ([32]byte, error) {
	return _Agglayerbridgel2frometrog.Contract.ClaimedGlobalIndexHashChain(&_Agglayerbridgel2frometrog.CallOpts)
}

// ComputeTokenProxyAddress is a free data retrieval call binding the contract method 0xf214e161.
//
// Solidity: function computeTokenProxyAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) ComputeTokenProxyAddress(opts *bind.CallOpts, originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "computeTokenProxyAddress", originNetwork, originTokenAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeTokenProxyAddress is a free data retrieval call binding the contract method 0xf214e161.
//
// Solidity: function computeTokenProxyAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) ComputeTokenProxyAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.ComputeTokenProxyAddress(&_Agglayerbridgel2frometrog.CallOpts, originNetwork, originTokenAddress)
}

// ComputeTokenProxyAddress is a free data retrieval call binding the contract method 0xf214e161.
//
// Solidity: function computeTokenProxyAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) ComputeTokenProxyAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.ComputeTokenProxyAddress(&_Agglayerbridgel2frometrog.CallOpts, originNetwork, originTokenAddress)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) DepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "depositCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) DepositCount() (*big.Int, error) {
	return _Agglayerbridgel2frometrog.Contract.DepositCount(&_Agglayerbridgel2frometrog.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) DepositCount() (*big.Int, error) {
	return _Agglayerbridgel2frometrog.Contract.DepositCount(&_Agglayerbridgel2frometrog.CallOpts)
}

// EmergencyBridgePauser is a free data retrieval call binding the contract method 0x2f84c690.
//
// Solidity: function emergencyBridgePauser() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) EmergencyBridgePauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "emergencyBridgePauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmergencyBridgePauser is a free data retrieval call binding the contract method 0x2f84c690.
//
// Solidity: function emergencyBridgePauser() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) EmergencyBridgePauser() (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.EmergencyBridgePauser(&_Agglayerbridgel2frometrog.CallOpts)
}

// EmergencyBridgePauser is a free data retrieval call binding the contract method 0x2f84c690.
//
// Solidity: function emergencyBridgePauser() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) EmergencyBridgePauser() (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.EmergencyBridgePauser(&_Agglayerbridgel2frometrog.CallOpts)
}

// EmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0xae24490a.
//
// Solidity: function emergencyBridgeUnpauser() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) EmergencyBridgeUnpauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "emergencyBridgeUnpauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0xae24490a.
//
// Solidity: function emergencyBridgeUnpauser() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) EmergencyBridgeUnpauser() (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.EmergencyBridgeUnpauser(&_Agglayerbridgel2frometrog.CallOpts)
}

// EmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0xae24490a.
//
// Solidity: function emergencyBridgeUnpauser() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) EmergencyBridgeUnpauser() (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.EmergencyBridgeUnpauser(&_Agglayerbridgel2frometrog.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) GasTokenAddress() (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.GasTokenAddress(&_Agglayerbridgel2frometrog.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) GasTokenAddress() (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.GasTokenAddress(&_Agglayerbridgel2frometrog.CallOpts)
}

// GasTokenMetadata is a free data retrieval call binding the contract method 0x27aef4e8.
//
// Solidity: function gasTokenMetadata() view returns(bytes)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) GasTokenMetadata(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "gasTokenMetadata")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GasTokenMetadata is a free data retrieval call binding the contract method 0x27aef4e8.
//
// Solidity: function gasTokenMetadata() view returns(bytes)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) GasTokenMetadata() ([]byte, error) {
	return _Agglayerbridgel2frometrog.Contract.GasTokenMetadata(&_Agglayerbridgel2frometrog.CallOpts)
}

// GasTokenMetadata is a free data retrieval call binding the contract method 0x27aef4e8.
//
// Solidity: function gasTokenMetadata() view returns(bytes)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) GasTokenMetadata() ([]byte, error) {
	return _Agglayerbridgel2frometrog.Contract.GasTokenMetadata(&_Agglayerbridgel2frometrog.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) GasTokenNetwork() (uint32, error) {
	return _Agglayerbridgel2frometrog.Contract.GasTokenNetwork(&_Agglayerbridgel2frometrog.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) GasTokenNetwork() (uint32, error) {
	return _Agglayerbridgel2frometrog.Contract.GasTokenNetwork(&_Agglayerbridgel2frometrog.CallOpts)
}

// GetProxiedTokensManager is a free data retrieval call binding the contract method 0x38b8fbbb.
//
// Solidity: function getProxiedTokensManager() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) GetProxiedTokensManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "getProxiedTokensManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxiedTokensManager is a free data retrieval call binding the contract method 0x38b8fbbb.
//
// Solidity: function getProxiedTokensManager() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) GetProxiedTokensManager() (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.GetProxiedTokensManager(&_Agglayerbridgel2frometrog.CallOpts)
}

// GetProxiedTokensManager is a free data retrieval call binding the contract method 0x38b8fbbb.
//
// Solidity: function getProxiedTokensManager() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) GetProxiedTokensManager() (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.GetProxiedTokensManager(&_Agglayerbridgel2frometrog.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) GetRoot() ([32]byte, error) {
	return _Agglayerbridgel2frometrog.Contract.GetRoot(&_Agglayerbridgel2frometrog.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) GetRoot() ([32]byte, error) {
	return _Agglayerbridgel2frometrog.Contract.GetRoot(&_Agglayerbridgel2frometrog.CallOpts)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) GetTokenMetadata(opts *bind.CallOpts, token common.Address) ([]byte, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "getTokenMetadata", token)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) GetTokenMetadata(token common.Address) ([]byte, error) {
	return _Agglayerbridgel2frometrog.Contract.GetTokenMetadata(&_Agglayerbridgel2frometrog.CallOpts, token)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) GetTokenMetadata(token common.Address) ([]byte, error) {
	return _Agglayerbridgel2frometrog.Contract.GetTokenMetadata(&_Agglayerbridgel2frometrog.CallOpts, token)
}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) GetTokenWrappedAddress(opts *bind.CallOpts, originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "getTokenWrappedAddress", originNetwork, originTokenAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) GetTokenWrappedAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.GetTokenWrappedAddress(&_Agglayerbridgel2frometrog.CallOpts, originNetwork, originTokenAddress)
}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) GetTokenWrappedAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.GetTokenWrappedAddress(&_Agglayerbridgel2frometrog.CallOpts, originNetwork, originTokenAddress)
}

// GetWrappedTokenBridgeImplementation is a free data retrieval call binding the contract method 0x3b2fee9a.
//
// Solidity: function getWrappedTokenBridgeImplementation() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) GetWrappedTokenBridgeImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "getWrappedTokenBridgeImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWrappedTokenBridgeImplementation is a free data retrieval call binding the contract method 0x3b2fee9a.
//
// Solidity: function getWrappedTokenBridgeImplementation() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) GetWrappedTokenBridgeImplementation() (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.GetWrappedTokenBridgeImplementation(&_Agglayerbridgel2frometrog.CallOpts)
}

// GetWrappedTokenBridgeImplementation is a free data retrieval call binding the contract method 0x3b2fee9a.
//
// Solidity: function getWrappedTokenBridgeImplementation() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) GetWrappedTokenBridgeImplementation() (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.GetWrappedTokenBridgeImplementation(&_Agglayerbridgel2frometrog.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) GlobalExitRootManager() (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.GlobalExitRootManager(&_Agglayerbridgel2frometrog.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.GlobalExitRootManager(&_Agglayerbridgel2frometrog.CallOpts)
}

// IsClaimed is a free data retrieval call binding the contract method 0xcc461632.
//
// Solidity: function isClaimed(uint32 leafIndex, uint32 sourceBridgeNetwork) view returns(bool)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) IsClaimed(opts *bind.CallOpts, leafIndex uint32, sourceBridgeNetwork uint32) (bool, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "isClaimed", leafIndex, sourceBridgeNetwork)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClaimed is a free data retrieval call binding the contract method 0xcc461632.
//
// Solidity: function isClaimed(uint32 leafIndex, uint32 sourceBridgeNetwork) view returns(bool)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) IsClaimed(leafIndex uint32, sourceBridgeNetwork uint32) (bool, error) {
	return _Agglayerbridgel2frometrog.Contract.IsClaimed(&_Agglayerbridgel2frometrog.CallOpts, leafIndex, sourceBridgeNetwork)
}

// IsClaimed is a free data retrieval call binding the contract method 0xcc461632.
//
// Solidity: function isClaimed(uint32 leafIndex, uint32 sourceBridgeNetwork) view returns(bool)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) IsClaimed(leafIndex uint32, sourceBridgeNetwork uint32) (bool, error) {
	return _Agglayerbridgel2frometrog.Contract.IsClaimed(&_Agglayerbridgel2frometrog.CallOpts, leafIndex, sourceBridgeNetwork)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) IsEmergencyState() (bool, error) {
	return _Agglayerbridgel2frometrog.Contract.IsEmergencyState(&_Agglayerbridgel2frometrog.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) IsEmergencyState() (bool, error) {
	return _Agglayerbridgel2frometrog.Contract.IsEmergencyState(&_Agglayerbridgel2frometrog.CallOpts)
}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) LastUpdatedDepositCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "lastUpdatedDepositCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) LastUpdatedDepositCount() (uint32, error) {
	return _Agglayerbridgel2frometrog.Contract.LastUpdatedDepositCount(&_Agglayerbridgel2frometrog.CallOpts)
}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) LastUpdatedDepositCount() (uint32, error) {
	return _Agglayerbridgel2frometrog.Contract.LastUpdatedDepositCount(&_Agglayerbridgel2frometrog.CallOpts)
}

// LocalBalanceTree is a free data retrieval call binding the contract method 0xd9cb3aec.
//
// Solidity: function localBalanceTree(bytes32 tokenInfoHash) view returns(uint256 amount)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) LocalBalanceTree(opts *bind.CallOpts, tokenInfoHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "localBalanceTree", tokenInfoHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LocalBalanceTree is a free data retrieval call binding the contract method 0xd9cb3aec.
//
// Solidity: function localBalanceTree(bytes32 tokenInfoHash) view returns(uint256 amount)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) LocalBalanceTree(tokenInfoHash [32]byte) (*big.Int, error) {
	return _Agglayerbridgel2frometrog.Contract.LocalBalanceTree(&_Agglayerbridgel2frometrog.CallOpts, tokenInfoHash)
}

// LocalBalanceTree is a free data retrieval call binding the contract method 0xd9cb3aec.
//
// Solidity: function localBalanceTree(bytes32 tokenInfoHash) view returns(uint256 amount)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) LocalBalanceTree(tokenInfoHash [32]byte) (*big.Int, error) {
	return _Agglayerbridgel2frometrog.Contract.LocalBalanceTree(&_Agglayerbridgel2frometrog.CallOpts, tokenInfoHash)
}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) NetworkID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "networkID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) NetworkID() (uint32, error) {
	return _Agglayerbridgel2frometrog.Contract.NetworkID(&_Agglayerbridgel2frometrog.CallOpts)
}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) NetworkID() (uint32, error) {
	return _Agglayerbridgel2frometrog.Contract.NetworkID(&_Agglayerbridgel2frometrog.CallOpts)
}

// PendingEmergencyBridgePauser is a free data retrieval call binding the contract method 0x03e6e116.
//
// Solidity: function pendingEmergencyBridgePauser() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) PendingEmergencyBridgePauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "pendingEmergencyBridgePauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingEmergencyBridgePauser is a free data retrieval call binding the contract method 0x03e6e116.
//
// Solidity: function pendingEmergencyBridgePauser() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) PendingEmergencyBridgePauser() (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.PendingEmergencyBridgePauser(&_Agglayerbridgel2frometrog.CallOpts)
}

// PendingEmergencyBridgePauser is a free data retrieval call binding the contract method 0x03e6e116.
//
// Solidity: function pendingEmergencyBridgePauser() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) PendingEmergencyBridgePauser() (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.PendingEmergencyBridgePauser(&_Agglayerbridgel2frometrog.CallOpts)
}

// PendingEmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0x606617ff.
//
// Solidity: function pendingEmergencyBridgeUnpauser() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) PendingEmergencyBridgeUnpauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "pendingEmergencyBridgeUnpauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingEmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0x606617ff.
//
// Solidity: function pendingEmergencyBridgeUnpauser() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) PendingEmergencyBridgeUnpauser() (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.PendingEmergencyBridgeUnpauser(&_Agglayerbridgel2frometrog.CallOpts)
}

// PendingEmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0x606617ff.
//
// Solidity: function pendingEmergencyBridgeUnpauser() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) PendingEmergencyBridgeUnpauser() (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.PendingEmergencyBridgeUnpauser(&_Agglayerbridgel2frometrog.CallOpts)
}

// PendingProxiedTokensManager is a free data retrieval call binding the contract method 0xece93c6f.
//
// Solidity: function pendingProxiedTokensManager() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) PendingProxiedTokensManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "pendingProxiedTokensManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingProxiedTokensManager is a free data retrieval call binding the contract method 0xece93c6f.
//
// Solidity: function pendingProxiedTokensManager() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) PendingProxiedTokensManager() (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.PendingProxiedTokensManager(&_Agglayerbridgel2frometrog.CallOpts)
}

// PendingProxiedTokensManager is a free data retrieval call binding the contract method 0xece93c6f.
//
// Solidity: function pendingProxiedTokensManager() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) PendingProxiedTokensManager() (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.PendingProxiedTokensManager(&_Agglayerbridgel2frometrog.CallOpts)
}

// PolygonRollupManager is a free data retrieval call binding the contract method 0x8ed7e3f2.
//
// Solidity: function polygonRollupManager() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) PolygonRollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "polygonRollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PolygonRollupManager is a free data retrieval call binding the contract method 0x8ed7e3f2.
//
// Solidity: function polygonRollupManager() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) PolygonRollupManager() (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.PolygonRollupManager(&_Agglayerbridgel2frometrog.CallOpts)
}

// PolygonRollupManager is a free data retrieval call binding the contract method 0x8ed7e3f2.
//
// Solidity: function polygonRollupManager() view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) PolygonRollupManager() (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.PolygonRollupManager(&_Agglayerbridgel2frometrog.CallOpts)
}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) TokenInfoToWrappedToken(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "tokenInfoToWrappedToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) TokenInfoToWrappedToken(arg0 [32]byte) (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.TokenInfoToWrappedToken(&_Agglayerbridgel2frometrog.CallOpts, arg0)
}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) TokenInfoToWrappedToken(arg0 [32]byte) (common.Address, error) {
	return _Agglayerbridgel2frometrog.Contract.TokenInfoToWrappedToken(&_Agglayerbridgel2frometrog.CallOpts, arg0)
}

// UnsetGlobalIndexHashChain is a free data retrieval call binding the contract method 0x6ee84b23.
//
// Solidity: function unsetGlobalIndexHashChain() view returns(bytes32)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) UnsetGlobalIndexHashChain(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "unsetGlobalIndexHashChain")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UnsetGlobalIndexHashChain is a free data retrieval call binding the contract method 0x6ee84b23.
//
// Solidity: function unsetGlobalIndexHashChain() view returns(bytes32)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) UnsetGlobalIndexHashChain() ([32]byte, error) {
	return _Agglayerbridgel2frometrog.Contract.UnsetGlobalIndexHashChain(&_Agglayerbridgel2frometrog.CallOpts)
}

// UnsetGlobalIndexHashChain is a free data retrieval call binding the contract method 0x6ee84b23.
//
// Solidity: function unsetGlobalIndexHashChain() view returns(bytes32)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) UnsetGlobalIndexHashChain() ([32]byte, error) {
	return _Agglayerbridgel2frometrog.Contract.UnsetGlobalIndexHashChain(&_Agglayerbridgel2frometrog.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) Version() (string, error) {
	return _Agglayerbridgel2frometrog.Contract.Version(&_Agglayerbridgel2frometrog.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) Version() (string, error) {
	return _Agglayerbridgel2frometrog.Contract.Version(&_Agglayerbridgel2frometrog.CallOpts)
}

// WrappedAddressIsNotMintable is a free data retrieval call binding the contract method 0xc0f49163.
//
// Solidity: function wrappedAddressIsNotMintable(address wrappedAddress) view returns(bool isNotMintable)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) WrappedAddressIsNotMintable(opts *bind.CallOpts, wrappedAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "wrappedAddressIsNotMintable", wrappedAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WrappedAddressIsNotMintable is a free data retrieval call binding the contract method 0xc0f49163.
//
// Solidity: function wrappedAddressIsNotMintable(address wrappedAddress) view returns(bool isNotMintable)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) WrappedAddressIsNotMintable(wrappedAddress common.Address) (bool, error) {
	return _Agglayerbridgel2frometrog.Contract.WrappedAddressIsNotMintable(&_Agglayerbridgel2frometrog.CallOpts, wrappedAddress)
}

// WrappedAddressIsNotMintable is a free data retrieval call binding the contract method 0xc0f49163.
//
// Solidity: function wrappedAddressIsNotMintable(address wrappedAddress) view returns(bool isNotMintable)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) WrappedAddressIsNotMintable(wrappedAddress common.Address) (bool, error) {
	return _Agglayerbridgel2frometrog.Contract.WrappedAddressIsNotMintable(&_Agglayerbridgel2frometrog.CallOpts, wrappedAddress)
}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address ) view returns(uint32 originNetwork, address originTokenAddress)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCaller) WrappedTokenToTokenInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	var out []interface{}
	err := _Agglayerbridgel2frometrog.contract.Call(opts, &out, "wrappedTokenToTokenInfo", arg0)

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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) WrappedTokenToTokenInfo(arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	return _Agglayerbridgel2frometrog.Contract.WrappedTokenToTokenInfo(&_Agglayerbridgel2frometrog.CallOpts, arg0)
}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address ) view returns(uint32 originNetwork, address originTokenAddress)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogCallerSession) WrappedTokenToTokenInfo(arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	return _Agglayerbridgel2frometrog.Contract.WrappedTokenToTokenInfo(&_Agglayerbridgel2frometrog.CallOpts, arg0)
}

// AcceptEmergencyBridgePauserRole is a paid mutator transaction binding the contract method 0xe88f0436.
//
// Solidity: function acceptEmergencyBridgePauserRole() returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) AcceptEmergencyBridgePauserRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "acceptEmergencyBridgePauserRole")
}

// AcceptEmergencyBridgePauserRole is a paid mutator transaction binding the contract method 0xe88f0436.
//
// Solidity: function acceptEmergencyBridgePauserRole() returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) AcceptEmergencyBridgePauserRole() (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.AcceptEmergencyBridgePauserRole(&_Agglayerbridgel2frometrog.TransactOpts)
}

// AcceptEmergencyBridgePauserRole is a paid mutator transaction binding the contract method 0xe88f0436.
//
// Solidity: function acceptEmergencyBridgePauserRole() returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) AcceptEmergencyBridgePauserRole() (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.AcceptEmergencyBridgePauserRole(&_Agglayerbridgel2frometrog.TransactOpts)
}

// AcceptEmergencyBridgeUnpauserRole is a paid mutator transaction binding the contract method 0x8b37b873.
//
// Solidity: function acceptEmergencyBridgeUnpauserRole() returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) AcceptEmergencyBridgeUnpauserRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "acceptEmergencyBridgeUnpauserRole")
}

// AcceptEmergencyBridgeUnpauserRole is a paid mutator transaction binding the contract method 0x8b37b873.
//
// Solidity: function acceptEmergencyBridgeUnpauserRole() returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) AcceptEmergencyBridgeUnpauserRole() (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.AcceptEmergencyBridgeUnpauserRole(&_Agglayerbridgel2frometrog.TransactOpts)
}

// AcceptEmergencyBridgeUnpauserRole is a paid mutator transaction binding the contract method 0x8b37b873.
//
// Solidity: function acceptEmergencyBridgeUnpauserRole() returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) AcceptEmergencyBridgeUnpauserRole() (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.AcceptEmergencyBridgeUnpauserRole(&_Agglayerbridgel2frometrog.TransactOpts)
}

// AcceptProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8c668f1c.
//
// Solidity: function acceptProxiedTokensManagerRole() returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) AcceptProxiedTokensManagerRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "acceptProxiedTokensManagerRole")
}

// AcceptProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8c668f1c.
//
// Solidity: function acceptProxiedTokensManagerRole() returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) AcceptProxiedTokensManagerRole() (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.AcceptProxiedTokensManagerRole(&_Agglayerbridgel2frometrog.TransactOpts)
}

// AcceptProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8c668f1c.
//
// Solidity: function acceptProxiedTokensManagerRole() returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) AcceptProxiedTokensManagerRole() (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.AcceptProxiedTokensManagerRole(&_Agglayerbridgel2frometrog.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) ActivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "activateEmergencyState")
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.ActivateEmergencyState(&_Agglayerbridgel2frometrog.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.ActivateEmergencyState(&_Agglayerbridgel2frometrog.TransactOpts)
}

// BackwardLET is a paid mutator transaction binding the contract method 0xfd7640e8.
//
// Solidity: function backwardLET(uint256 newDepositCount, bytes32[32] newFrontier, bytes32 nextLeaf, bytes32[32] proof) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) BackwardLET(opts *bind.TransactOpts, newDepositCount *big.Int, newFrontier [32][32]byte, nextLeaf [32]byte, proof [32][32]byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "backwardLET", newDepositCount, newFrontier, nextLeaf, proof)
}

// BackwardLET is a paid mutator transaction binding the contract method 0xfd7640e8.
//
// Solidity: function backwardLET(uint256 newDepositCount, bytes32[32] newFrontier, bytes32 nextLeaf, bytes32[32] proof) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) BackwardLET(newDepositCount *big.Int, newFrontier [32][32]byte, nextLeaf [32]byte, proof [32][32]byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.BackwardLET(&_Agglayerbridgel2frometrog.TransactOpts, newDepositCount, newFrontier, nextLeaf, proof)
}

// BackwardLET is a paid mutator transaction binding the contract method 0xfd7640e8.
//
// Solidity: function backwardLET(uint256 newDepositCount, bytes32[32] newFrontier, bytes32 nextLeaf, bytes32[32] proof) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) BackwardLET(newDepositCount *big.Int, newFrontier [32][32]byte, nextLeaf [32]byte, proof [32][32]byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.BackwardLET(&_Agglayerbridgel2frometrog.TransactOpts, newDepositCount, newFrontier, nextLeaf, proof)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) BridgeAsset(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "bridgeAsset", destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) BridgeAsset(destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.BridgeAsset(&_Agglayerbridgel2frometrog.TransactOpts, destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) BridgeAsset(destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.BridgeAsset(&_Agglayerbridgel2frometrog.TransactOpts, destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) BridgeMessage(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "bridgeMessage", destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) BridgeMessage(destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.BridgeMessage(&_Agglayerbridgel2frometrog.TransactOpts, destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) BridgeMessage(destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.BridgeMessage(&_Agglayerbridgel2frometrog.TransactOpts, destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) BridgeMessageWETH(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "bridgeMessageWETH", destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) BridgeMessageWETH(destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.BridgeMessageWETH(&_Agglayerbridgel2frometrog.TransactOpts, destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) BridgeMessageWETH(destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.BridgeMessageWETH(&_Agglayerbridgel2frometrog.TransactOpts, destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) ClaimAsset(opts *bind.TransactOpts, smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "claimAsset", smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) ClaimAsset(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.ClaimAsset(&_Agglayerbridgel2frometrog.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) ClaimAsset(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.ClaimAsset(&_Agglayerbridgel2frometrog.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) ClaimMessage(opts *bind.TransactOpts, smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "claimMessage", smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) ClaimMessage(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.ClaimMessage(&_Agglayerbridgel2frometrog.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) ClaimMessage(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.ClaimMessage(&_Agglayerbridgel2frometrog.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.DeactivateEmergencyState(&_Agglayerbridgel2frometrog.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.DeactivateEmergencyState(&_Agglayerbridgel2frometrog.TransactOpts)
}

// DeployWrappedTokenAndRemap is a paid mutator transaction binding the contract method 0x6e974cd4.
//
// Solidity: function deployWrappedTokenAndRemap(uint32 originNetwork, address originTokenAddress, bool isNotMintable) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) DeployWrappedTokenAndRemap(opts *bind.TransactOpts, originNetwork uint32, originTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "deployWrappedTokenAndRemap", originNetwork, originTokenAddress, isNotMintable)
}

// DeployWrappedTokenAndRemap is a paid mutator transaction binding the contract method 0x6e974cd4.
//
// Solidity: function deployWrappedTokenAndRemap(uint32 originNetwork, address originTokenAddress, bool isNotMintable) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) DeployWrappedTokenAndRemap(originNetwork uint32, originTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.DeployWrappedTokenAndRemap(&_Agglayerbridgel2frometrog.TransactOpts, originNetwork, originTokenAddress, isNotMintable)
}

// DeployWrappedTokenAndRemap is a paid mutator transaction binding the contract method 0x6e974cd4.
//
// Solidity: function deployWrappedTokenAndRemap(uint32 originNetwork, address originTokenAddress, bool isNotMintable) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) DeployWrappedTokenAndRemap(originNetwork uint32, originTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.DeployWrappedTokenAndRemap(&_Agglayerbridgel2frometrog.TransactOpts, originNetwork, originTokenAddress, isNotMintable)
}

// ForceEmitDetailedClaimEvent is a paid mutator transaction binding the contract method 0x567f13e4.
//
// Solidity: function forceEmitDetailedClaimEvent((bytes32[32],bytes32[32],uint256,bytes32,bytes32,uint8,uint32,address,uint32,address,uint256,bytes)[] claims) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) ForceEmitDetailedClaimEvent(opts *bind.TransactOpts, claims []AgglayerBridgeL2ClaimData) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "forceEmitDetailedClaimEvent", claims)
}

// ForceEmitDetailedClaimEvent is a paid mutator transaction binding the contract method 0x567f13e4.
//
// Solidity: function forceEmitDetailedClaimEvent((bytes32[32],bytes32[32],uint256,bytes32,bytes32,uint8,uint32,address,uint32,address,uint256,bytes)[] claims) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) ForceEmitDetailedClaimEvent(claims []AgglayerBridgeL2ClaimData) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.ForceEmitDetailedClaimEvent(&_Agglayerbridgel2frometrog.TransactOpts, claims)
}

// ForceEmitDetailedClaimEvent is a paid mutator transaction binding the contract method 0x567f13e4.
//
// Solidity: function forceEmitDetailedClaimEvent((bytes32[32],bytes32[32],uint256,bytes32,bytes32,uint8,uint32,address,uint32,address,uint256,bytes)[] claims) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) ForceEmitDetailedClaimEvent(claims []AgglayerBridgeL2ClaimData) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.ForceEmitDetailedClaimEvent(&_Agglayerbridgel2frometrog.TransactOpts, claims)
}

// ForwardLET is a paid mutator transaction binding the contract method 0xf0e70808.
//
// Solidity: function forwardLET((uint8,uint32,address,uint32,address,uint256,bytes)[] newLeaves, bytes32 expectedLER) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) ForwardLET(opts *bind.TransactOpts, newLeaves []AgglayerBridgeL2LeafData, expectedLER [32]byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "forwardLET", newLeaves, expectedLER)
}

// ForwardLET is a paid mutator transaction binding the contract method 0xf0e70808.
//
// Solidity: function forwardLET((uint8,uint32,address,uint32,address,uint256,bytes)[] newLeaves, bytes32 expectedLER) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) ForwardLET(newLeaves []AgglayerBridgeL2LeafData, expectedLER [32]byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.ForwardLET(&_Agglayerbridgel2frometrog.TransactOpts, newLeaves, expectedLER)
}

// ForwardLET is a paid mutator transaction binding the contract method 0xf0e70808.
//
// Solidity: function forwardLET((uint8,uint32,address,uint32,address,uint256,bytes)[] newLeaves, bytes32 expectedLER) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) ForwardLET(newLeaves []AgglayerBridgeL2LeafData, expectedLER [32]byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.ForwardLET(&_Agglayerbridgel2frometrog.TransactOpts, newLeaves, expectedLER)
}

// Initialize is a paid mutator transaction binding the contract method 0x006ee171.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes , address , address , bool , address , address , address ) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) Initialize(opts *bind.TransactOpts, arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte, arg6 common.Address, arg7 common.Address, arg8 bool, arg9 common.Address, arg10 common.Address, arg11 common.Address) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "initialize", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11)
}

// Initialize is a paid mutator transaction binding the contract method 0x006ee171.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes , address , address , bool , address , address , address ) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) Initialize(arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte, arg6 common.Address, arg7 common.Address, arg8 bool, arg9 common.Address, arg10 common.Address, arg11 common.Address) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.Initialize(&_Agglayerbridgel2frometrog.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11)
}

// Initialize is a paid mutator transaction binding the contract method 0x006ee171.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes , address , address , bool , address , address , address ) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) Initialize(arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte, arg6 common.Address, arg7 common.Address, arg8 bool, arg9 common.Address, arg10 common.Address, arg11 common.Address) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.Initialize(&_Agglayerbridgel2frometrog.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes ) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) Initialize0(opts *bind.TransactOpts, arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "initialize0", arg0, arg1, arg2, arg3, arg4, arg5)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes ) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) Initialize0(arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.Initialize0(&_Agglayerbridgel2frometrog.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes ) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) Initialize0(arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.Initialize0(&_Agglayerbridgel2frometrog.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// InitializeFromEtrog is a paid mutator transaction binding the contract method 0x15419305.
//
// Solidity: function initializeFromEtrog(address _bridgeManager, address _emergencyBridgePauser, address _emergencyBridgeUnpauser, address _proxiedTokensManager, address[] wrappedTokensAddresses, uint128 initNativeSupply) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) InitializeFromEtrog(opts *bind.TransactOpts, _bridgeManager common.Address, _emergencyBridgePauser common.Address, _emergencyBridgeUnpauser common.Address, _proxiedTokensManager common.Address, wrappedTokensAddresses []common.Address, initNativeSupply *big.Int) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "initializeFromEtrog", _bridgeManager, _emergencyBridgePauser, _emergencyBridgeUnpauser, _proxiedTokensManager, wrappedTokensAddresses, initNativeSupply)
}

// InitializeFromEtrog is a paid mutator transaction binding the contract method 0x15419305.
//
// Solidity: function initializeFromEtrog(address _bridgeManager, address _emergencyBridgePauser, address _emergencyBridgeUnpauser, address _proxiedTokensManager, address[] wrappedTokensAddresses, uint128 initNativeSupply) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) InitializeFromEtrog(_bridgeManager common.Address, _emergencyBridgePauser common.Address, _emergencyBridgeUnpauser common.Address, _proxiedTokensManager common.Address, wrappedTokensAddresses []common.Address, initNativeSupply *big.Int) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.InitializeFromEtrog(&_Agglayerbridgel2frometrog.TransactOpts, _bridgeManager, _emergencyBridgePauser, _emergencyBridgeUnpauser, _proxiedTokensManager, wrappedTokensAddresses, initNativeSupply)
}

// InitializeFromEtrog is a paid mutator transaction binding the contract method 0x15419305.
//
// Solidity: function initializeFromEtrog(address _bridgeManager, address _emergencyBridgePauser, address _emergencyBridgeUnpauser, address _proxiedTokensManager, address[] wrappedTokensAddresses, uint128 initNativeSupply) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) InitializeFromEtrog(_bridgeManager common.Address, _emergencyBridgePauser common.Address, _emergencyBridgeUnpauser common.Address, _proxiedTokensManager common.Address, wrappedTokensAddresses []common.Address, initNativeSupply *big.Int) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.InitializeFromEtrog(&_Agglayerbridgel2frometrog.TransactOpts, _bridgeManager, _emergencyBridgePauser, _emergencyBridgeUnpauser, _proxiedTokensManager, wrappedTokensAddresses, initNativeSupply)
}

// MigrateLegacyToken is a paid mutator transaction binding the contract method 0xb0b37920.
//
// Solidity: function migrateLegacyToken(address legacyTokenAddress, uint256 amount, bytes permitData) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) MigrateLegacyToken(opts *bind.TransactOpts, legacyTokenAddress common.Address, amount *big.Int, permitData []byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "migrateLegacyToken", legacyTokenAddress, amount, permitData)
}

// MigrateLegacyToken is a paid mutator transaction binding the contract method 0xb0b37920.
//
// Solidity: function migrateLegacyToken(address legacyTokenAddress, uint256 amount, bytes permitData) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) MigrateLegacyToken(legacyTokenAddress common.Address, amount *big.Int, permitData []byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.MigrateLegacyToken(&_Agglayerbridgel2frometrog.TransactOpts, legacyTokenAddress, amount, permitData)
}

// MigrateLegacyToken is a paid mutator transaction binding the contract method 0xb0b37920.
//
// Solidity: function migrateLegacyToken(address legacyTokenAddress, uint256 amount, bytes permitData) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) MigrateLegacyToken(legacyTokenAddress common.Address, amount *big.Int, permitData []byte) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.MigrateLegacyToken(&_Agglayerbridgel2frometrog.TransactOpts, legacyTokenAddress, amount, permitData)
}

// RemoveLegacySovereignTokenAddress is a paid mutator transaction binding the contract method 0xb4586962.
//
// Solidity: function removeLegacySovereignTokenAddress(address legacySovereignTokenAddress) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) RemoveLegacySovereignTokenAddress(opts *bind.TransactOpts, legacySovereignTokenAddress common.Address) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "removeLegacySovereignTokenAddress", legacySovereignTokenAddress)
}

// RemoveLegacySovereignTokenAddress is a paid mutator transaction binding the contract method 0xb4586962.
//
// Solidity: function removeLegacySovereignTokenAddress(address legacySovereignTokenAddress) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) RemoveLegacySovereignTokenAddress(legacySovereignTokenAddress common.Address) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.RemoveLegacySovereignTokenAddress(&_Agglayerbridgel2frometrog.TransactOpts, legacySovereignTokenAddress)
}

// RemoveLegacySovereignTokenAddress is a paid mutator transaction binding the contract method 0xb4586962.
//
// Solidity: function removeLegacySovereignTokenAddress(address legacySovereignTokenAddress) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) RemoveLegacySovereignTokenAddress(legacySovereignTokenAddress common.Address) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.RemoveLegacySovereignTokenAddress(&_Agglayerbridgel2frometrog.TransactOpts, legacySovereignTokenAddress)
}

// SetBridgeManager is a paid mutator transaction binding the contract method 0xeabd372a.
//
// Solidity: function setBridgeManager(address _bridgeManager) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) SetBridgeManager(opts *bind.TransactOpts, _bridgeManager common.Address) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "setBridgeManager", _bridgeManager)
}

// SetBridgeManager is a paid mutator transaction binding the contract method 0xeabd372a.
//
// Solidity: function setBridgeManager(address _bridgeManager) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) SetBridgeManager(_bridgeManager common.Address) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.SetBridgeManager(&_Agglayerbridgel2frometrog.TransactOpts, _bridgeManager)
}

// SetBridgeManager is a paid mutator transaction binding the contract method 0xeabd372a.
//
// Solidity: function setBridgeManager(address _bridgeManager) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) SetBridgeManager(_bridgeManager common.Address) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.SetBridgeManager(&_Agglayerbridgel2frometrog.TransactOpts, _bridgeManager)
}

// SetLocalBalanceTree is a paid mutator transaction binding the contract method 0x8f9720bb.
//
// Solidity: function setLocalBalanceTree(uint32[] originNetworkArray, address[] originTokenAddressArray, uint256[] amountArray) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) SetLocalBalanceTree(opts *bind.TransactOpts, originNetworkArray []uint32, originTokenAddressArray []common.Address, amountArray []*big.Int) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "setLocalBalanceTree", originNetworkArray, originTokenAddressArray, amountArray)
}

// SetLocalBalanceTree is a paid mutator transaction binding the contract method 0x8f9720bb.
//
// Solidity: function setLocalBalanceTree(uint32[] originNetworkArray, address[] originTokenAddressArray, uint256[] amountArray) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) SetLocalBalanceTree(originNetworkArray []uint32, originTokenAddressArray []common.Address, amountArray []*big.Int) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.SetLocalBalanceTree(&_Agglayerbridgel2frometrog.TransactOpts, originNetworkArray, originTokenAddressArray, amountArray)
}

// SetLocalBalanceTree is a paid mutator transaction binding the contract method 0x8f9720bb.
//
// Solidity: function setLocalBalanceTree(uint32[] originNetworkArray, address[] originTokenAddressArray, uint256[] amountArray) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) SetLocalBalanceTree(originNetworkArray []uint32, originTokenAddressArray []common.Address, amountArray []*big.Int) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.SetLocalBalanceTree(&_Agglayerbridgel2frometrog.TransactOpts, originNetworkArray, originTokenAddressArray, amountArray)
}

// SetMultipleClaims is a paid mutator transaction binding the contract method 0x1c208229.
//
// Solidity: function setMultipleClaims(uint256[] globalIndexes) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) SetMultipleClaims(opts *bind.TransactOpts, globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "setMultipleClaims", globalIndexes)
}

// SetMultipleClaims is a paid mutator transaction binding the contract method 0x1c208229.
//
// Solidity: function setMultipleClaims(uint256[] globalIndexes) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) SetMultipleClaims(globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.SetMultipleClaims(&_Agglayerbridgel2frometrog.TransactOpts, globalIndexes)
}

// SetMultipleClaims is a paid mutator transaction binding the contract method 0x1c208229.
//
// Solidity: function setMultipleClaims(uint256[] globalIndexes) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) SetMultipleClaims(globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.SetMultipleClaims(&_Agglayerbridgel2frometrog.TransactOpts, globalIndexes)
}

// SetMultipleSovereignTokenAddress is a paid mutator transaction binding the contract method 0x57cfbee3.
//
// Solidity: function setMultipleSovereignTokenAddress(uint32[] originNetworks, address[] originTokenAddresses, address[] sovereignTokenAddresses, bool[] isNotMintable) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) SetMultipleSovereignTokenAddress(opts *bind.TransactOpts, originNetworks []uint32, originTokenAddresses []common.Address, sovereignTokenAddresses []common.Address, isNotMintable []bool) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "setMultipleSovereignTokenAddress", originNetworks, originTokenAddresses, sovereignTokenAddresses, isNotMintable)
}

// SetMultipleSovereignTokenAddress is a paid mutator transaction binding the contract method 0x57cfbee3.
//
// Solidity: function setMultipleSovereignTokenAddress(uint32[] originNetworks, address[] originTokenAddresses, address[] sovereignTokenAddresses, bool[] isNotMintable) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) SetMultipleSovereignTokenAddress(originNetworks []uint32, originTokenAddresses []common.Address, sovereignTokenAddresses []common.Address, isNotMintable []bool) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.SetMultipleSovereignTokenAddress(&_Agglayerbridgel2frometrog.TransactOpts, originNetworks, originTokenAddresses, sovereignTokenAddresses, isNotMintable)
}

// SetMultipleSovereignTokenAddress is a paid mutator transaction binding the contract method 0x57cfbee3.
//
// Solidity: function setMultipleSovereignTokenAddress(uint32[] originNetworks, address[] originTokenAddresses, address[] sovereignTokenAddresses, bool[] isNotMintable) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) SetMultipleSovereignTokenAddress(originNetworks []uint32, originTokenAddresses []common.Address, sovereignTokenAddresses []common.Address, isNotMintable []bool) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.SetMultipleSovereignTokenAddress(&_Agglayerbridgel2frometrog.TransactOpts, originNetworks, originTokenAddresses, sovereignTokenAddresses, isNotMintable)
}

// SetSovereignWETHAddress is a paid mutator transaction binding the contract method 0xbf130d7f.
//
// Solidity: function setSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) SetSovereignWETHAddress(opts *bind.TransactOpts, sovereignWETHTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "setSovereignWETHAddress", sovereignWETHTokenAddress, isNotMintable)
}

// SetSovereignWETHAddress is a paid mutator transaction binding the contract method 0xbf130d7f.
//
// Solidity: function setSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) SetSovereignWETHAddress(sovereignWETHTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.SetSovereignWETHAddress(&_Agglayerbridgel2frometrog.TransactOpts, sovereignWETHTokenAddress, isNotMintable)
}

// SetSovereignWETHAddress is a paid mutator transaction binding the contract method 0xbf130d7f.
//
// Solidity: function setSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) SetSovereignWETHAddress(sovereignWETHTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.SetSovereignWETHAddress(&_Agglayerbridgel2frometrog.TransactOpts, sovereignWETHTokenAddress, isNotMintable)
}

// TransferEmergencyBridgePauserRole is a paid mutator transaction binding the contract method 0x69e3ab12.
//
// Solidity: function transferEmergencyBridgePauserRole(address newEmergencyBridgePauser) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) TransferEmergencyBridgePauserRole(opts *bind.TransactOpts, newEmergencyBridgePauser common.Address) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "transferEmergencyBridgePauserRole", newEmergencyBridgePauser)
}

// TransferEmergencyBridgePauserRole is a paid mutator transaction binding the contract method 0x69e3ab12.
//
// Solidity: function transferEmergencyBridgePauserRole(address newEmergencyBridgePauser) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) TransferEmergencyBridgePauserRole(newEmergencyBridgePauser common.Address) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.TransferEmergencyBridgePauserRole(&_Agglayerbridgel2frometrog.TransactOpts, newEmergencyBridgePauser)
}

// TransferEmergencyBridgePauserRole is a paid mutator transaction binding the contract method 0x69e3ab12.
//
// Solidity: function transferEmergencyBridgePauserRole(address newEmergencyBridgePauser) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) TransferEmergencyBridgePauserRole(newEmergencyBridgePauser common.Address) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.TransferEmergencyBridgePauserRole(&_Agglayerbridgel2frometrog.TransactOpts, newEmergencyBridgePauser)
}

// TransferEmergencyBridgeUnpauserRole is a paid mutator transaction binding the contract method 0x8d942096.
//
// Solidity: function transferEmergencyBridgeUnpauserRole(address newEmergencyBridgeUnpauser) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) TransferEmergencyBridgeUnpauserRole(opts *bind.TransactOpts, newEmergencyBridgeUnpauser common.Address) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "transferEmergencyBridgeUnpauserRole", newEmergencyBridgeUnpauser)
}

// TransferEmergencyBridgeUnpauserRole is a paid mutator transaction binding the contract method 0x8d942096.
//
// Solidity: function transferEmergencyBridgeUnpauserRole(address newEmergencyBridgeUnpauser) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) TransferEmergencyBridgeUnpauserRole(newEmergencyBridgeUnpauser common.Address) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.TransferEmergencyBridgeUnpauserRole(&_Agglayerbridgel2frometrog.TransactOpts, newEmergencyBridgeUnpauser)
}

// TransferEmergencyBridgeUnpauserRole is a paid mutator transaction binding the contract method 0x8d942096.
//
// Solidity: function transferEmergencyBridgeUnpauserRole(address newEmergencyBridgeUnpauser) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) TransferEmergencyBridgeUnpauserRole(newEmergencyBridgeUnpauser common.Address) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.TransferEmergencyBridgeUnpauserRole(&_Agglayerbridgel2frometrog.TransactOpts, newEmergencyBridgeUnpauser)
}

// TransferProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8bd309c3.
//
// Solidity: function transferProxiedTokensManagerRole(address newProxiedTokensManager) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) TransferProxiedTokensManagerRole(opts *bind.TransactOpts, newProxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "transferProxiedTokensManagerRole", newProxiedTokensManager)
}

// TransferProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8bd309c3.
//
// Solidity: function transferProxiedTokensManagerRole(address newProxiedTokensManager) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) TransferProxiedTokensManagerRole(newProxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.TransferProxiedTokensManagerRole(&_Agglayerbridgel2frometrog.TransactOpts, newProxiedTokensManager)
}

// TransferProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8bd309c3.
//
// Solidity: function transferProxiedTokensManagerRole(address newProxiedTokensManager) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) TransferProxiedTokensManagerRole(newProxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.TransferProxiedTokensManagerRole(&_Agglayerbridgel2frometrog.TransactOpts, newProxiedTokensManager)
}

// UnsetMultipleClaims is a paid mutator transaction binding the contract method 0x136a2c60.
//
// Solidity: function unsetMultipleClaims(uint256[] globalIndexes) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) UnsetMultipleClaims(opts *bind.TransactOpts, globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "unsetMultipleClaims", globalIndexes)
}

// UnsetMultipleClaims is a paid mutator transaction binding the contract method 0x136a2c60.
//
// Solidity: function unsetMultipleClaims(uint256[] globalIndexes) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) UnsetMultipleClaims(globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.UnsetMultipleClaims(&_Agglayerbridgel2frometrog.TransactOpts, globalIndexes)
}

// UnsetMultipleClaims is a paid mutator transaction binding the contract method 0x136a2c60.
//
// Solidity: function unsetMultipleClaims(uint256[] globalIndexes) returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) UnsetMultipleClaims(globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.UnsetMultipleClaims(&_Agglayerbridgel2frometrog.TransactOpts, globalIndexes)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactor) UpdateGlobalExitRoot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.contract.Transact(opts, "updateGlobalExitRoot")
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogSession) UpdateGlobalExitRoot() (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.UpdateGlobalExitRoot(&_Agglayerbridgel2frometrog.TransactOpts)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogTransactorSession) UpdateGlobalExitRoot() (*types.Transaction, error) {
	return _Agglayerbridgel2frometrog.Contract.UpdateGlobalExitRoot(&_Agglayerbridgel2frometrog.TransactOpts)
}

// Agglayerbridgel2frometrogAcceptEmergencyBridgePauserRoleIterator is returned from FilterAcceptEmergencyBridgePauserRole and is used to iterate over the raw logs and unpacked data for AcceptEmergencyBridgePauserRole events raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogAcceptEmergencyBridgePauserRoleIterator struct {
	Event *Agglayerbridgel2frometrogAcceptEmergencyBridgePauserRole // Event containing the contract specifics and raw log

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
func (it *Agglayerbridgel2frometrogAcceptEmergencyBridgePauserRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Agglayerbridgel2frometrogAcceptEmergencyBridgePauserRole)
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
		it.Event = new(Agglayerbridgel2frometrogAcceptEmergencyBridgePauserRole)
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
func (it *Agglayerbridgel2frometrogAcceptEmergencyBridgePauserRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Agglayerbridgel2frometrogAcceptEmergencyBridgePauserRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Agglayerbridgel2frometrogAcceptEmergencyBridgePauserRole represents a AcceptEmergencyBridgePauserRole event raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogAcceptEmergencyBridgePauserRole struct {
	OldEmergencyBridgePauser common.Address
	NewEmergencyBridgePauser common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterAcceptEmergencyBridgePauserRole is a free log retrieval operation binding the contract event 0x85d2bdfbe58cd81abf8199c13ce2509204be4aba8603b9d29f52c4e13e7bb793.
//
// Solidity: event AcceptEmergencyBridgePauserRole(address oldEmergencyBridgePauser, address newEmergencyBridgePauser)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) FilterAcceptEmergencyBridgePauserRole(opts *bind.FilterOpts) (*Agglayerbridgel2frometrogAcceptEmergencyBridgePauserRoleIterator, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.FilterLogs(opts, "AcceptEmergencyBridgePauserRole")
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogAcceptEmergencyBridgePauserRoleIterator{contract: _Agglayerbridgel2frometrog.contract, event: "AcceptEmergencyBridgePauserRole", logs: logs, sub: sub}, nil
}

// WatchAcceptEmergencyBridgePauserRole is a free log subscription operation binding the contract event 0x85d2bdfbe58cd81abf8199c13ce2509204be4aba8603b9d29f52c4e13e7bb793.
//
// Solidity: event AcceptEmergencyBridgePauserRole(address oldEmergencyBridgePauser, address newEmergencyBridgePauser)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) WatchAcceptEmergencyBridgePauserRole(opts *bind.WatchOpts, sink chan<- *Agglayerbridgel2frometrogAcceptEmergencyBridgePauserRole) (event.Subscription, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.WatchLogs(opts, "AcceptEmergencyBridgePauserRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Agglayerbridgel2frometrogAcceptEmergencyBridgePauserRole)
				if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "AcceptEmergencyBridgePauserRole", log); err != nil {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) ParseAcceptEmergencyBridgePauserRole(log types.Log) (*Agglayerbridgel2frometrogAcceptEmergencyBridgePauserRole, error) {
	event := new(Agglayerbridgel2frometrogAcceptEmergencyBridgePauserRole)
	if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "AcceptEmergencyBridgePauserRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Agglayerbridgel2frometrogAcceptEmergencyBridgeUnpauserRoleIterator is returned from FilterAcceptEmergencyBridgeUnpauserRole and is used to iterate over the raw logs and unpacked data for AcceptEmergencyBridgeUnpauserRole events raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogAcceptEmergencyBridgeUnpauserRoleIterator struct {
	Event *Agglayerbridgel2frometrogAcceptEmergencyBridgeUnpauserRole // Event containing the contract specifics and raw log

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
func (it *Agglayerbridgel2frometrogAcceptEmergencyBridgeUnpauserRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Agglayerbridgel2frometrogAcceptEmergencyBridgeUnpauserRole)
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
		it.Event = new(Agglayerbridgel2frometrogAcceptEmergencyBridgeUnpauserRole)
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
func (it *Agglayerbridgel2frometrogAcceptEmergencyBridgeUnpauserRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Agglayerbridgel2frometrogAcceptEmergencyBridgeUnpauserRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Agglayerbridgel2frometrogAcceptEmergencyBridgeUnpauserRole represents a AcceptEmergencyBridgeUnpauserRole event raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogAcceptEmergencyBridgeUnpauserRole struct {
	OldEmergencyBridgeUnpauser common.Address
	NewEmergencyBridgeUnpauser common.Address
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterAcceptEmergencyBridgeUnpauserRole is a free log retrieval operation binding the contract event 0x24cc8295aa5110cc216695db944ad2458c7795c6404449be980c3ce14aed752d.
//
// Solidity: event AcceptEmergencyBridgeUnpauserRole(address oldEmergencyBridgeUnpauser, address newEmergencyBridgeUnpauser)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) FilterAcceptEmergencyBridgeUnpauserRole(opts *bind.FilterOpts) (*Agglayerbridgel2frometrogAcceptEmergencyBridgeUnpauserRoleIterator, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.FilterLogs(opts, "AcceptEmergencyBridgeUnpauserRole")
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogAcceptEmergencyBridgeUnpauserRoleIterator{contract: _Agglayerbridgel2frometrog.contract, event: "AcceptEmergencyBridgeUnpauserRole", logs: logs, sub: sub}, nil
}

// WatchAcceptEmergencyBridgeUnpauserRole is a free log subscription operation binding the contract event 0x24cc8295aa5110cc216695db944ad2458c7795c6404449be980c3ce14aed752d.
//
// Solidity: event AcceptEmergencyBridgeUnpauserRole(address oldEmergencyBridgeUnpauser, address newEmergencyBridgeUnpauser)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) WatchAcceptEmergencyBridgeUnpauserRole(opts *bind.WatchOpts, sink chan<- *Agglayerbridgel2frometrogAcceptEmergencyBridgeUnpauserRole) (event.Subscription, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.WatchLogs(opts, "AcceptEmergencyBridgeUnpauserRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Agglayerbridgel2frometrogAcceptEmergencyBridgeUnpauserRole)
				if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "AcceptEmergencyBridgeUnpauserRole", log); err != nil {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) ParseAcceptEmergencyBridgeUnpauserRole(log types.Log) (*Agglayerbridgel2frometrogAcceptEmergencyBridgeUnpauserRole, error) {
	event := new(Agglayerbridgel2frometrogAcceptEmergencyBridgeUnpauserRole)
	if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "AcceptEmergencyBridgeUnpauserRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Agglayerbridgel2frometrogAcceptProxiedTokensManagerRoleIterator is returned from FilterAcceptProxiedTokensManagerRole and is used to iterate over the raw logs and unpacked data for AcceptProxiedTokensManagerRole events raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogAcceptProxiedTokensManagerRoleIterator struct {
	Event *Agglayerbridgel2frometrogAcceptProxiedTokensManagerRole // Event containing the contract specifics and raw log

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
func (it *Agglayerbridgel2frometrogAcceptProxiedTokensManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Agglayerbridgel2frometrogAcceptProxiedTokensManagerRole)
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
		it.Event = new(Agglayerbridgel2frometrogAcceptProxiedTokensManagerRole)
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
func (it *Agglayerbridgel2frometrogAcceptProxiedTokensManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Agglayerbridgel2frometrogAcceptProxiedTokensManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Agglayerbridgel2frometrogAcceptProxiedTokensManagerRole represents a AcceptProxiedTokensManagerRole event raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogAcceptProxiedTokensManagerRole struct {
	OldProxiedTokensManager common.Address
	NewProxiedTokensManager common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterAcceptProxiedTokensManagerRole is a free log retrieval operation binding the contract event 0xa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f2.
//
// Solidity: event AcceptProxiedTokensManagerRole(address oldProxiedTokensManager, address newProxiedTokensManager)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) FilterAcceptProxiedTokensManagerRole(opts *bind.FilterOpts) (*Agglayerbridgel2frometrogAcceptProxiedTokensManagerRoleIterator, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.FilterLogs(opts, "AcceptProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogAcceptProxiedTokensManagerRoleIterator{contract: _Agglayerbridgel2frometrog.contract, event: "AcceptProxiedTokensManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptProxiedTokensManagerRole is a free log subscription operation binding the contract event 0xa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f2.
//
// Solidity: event AcceptProxiedTokensManagerRole(address oldProxiedTokensManager, address newProxiedTokensManager)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) WatchAcceptProxiedTokensManagerRole(opts *bind.WatchOpts, sink chan<- *Agglayerbridgel2frometrogAcceptProxiedTokensManagerRole) (event.Subscription, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.WatchLogs(opts, "AcceptProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Agglayerbridgel2frometrogAcceptProxiedTokensManagerRole)
				if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "AcceptProxiedTokensManagerRole", log); err != nil {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) ParseAcceptProxiedTokensManagerRole(log types.Log) (*Agglayerbridgel2frometrogAcceptProxiedTokensManagerRole, error) {
	event := new(Agglayerbridgel2frometrogAcceptProxiedTokensManagerRole)
	if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "AcceptProxiedTokensManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Agglayerbridgel2frometrogBackwardLETIterator is returned from FilterBackwardLET and is used to iterate over the raw logs and unpacked data for BackwardLET events raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogBackwardLETIterator struct {
	Event *Agglayerbridgel2frometrogBackwardLET // Event containing the contract specifics and raw log

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
func (it *Agglayerbridgel2frometrogBackwardLETIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Agglayerbridgel2frometrogBackwardLET)
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
		it.Event = new(Agglayerbridgel2frometrogBackwardLET)
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
func (it *Agglayerbridgel2frometrogBackwardLETIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Agglayerbridgel2frometrogBackwardLETIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Agglayerbridgel2frometrogBackwardLET represents a BackwardLET event raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogBackwardLET struct {
	PreviousDepositCount *big.Int
	PreviousRoot         [32]byte
	NewDepositCount      *big.Int
	NewRoot              [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterBackwardLET is a free log retrieval operation binding the contract event 0x2f8b2f36ca0d63af2efffe9d5e728a61c2bc273a3fc0574d12b558a41f149555.
//
// Solidity: event BackwardLET(uint256 previousDepositCount, bytes32 previousRoot, uint256 newDepositCount, bytes32 newRoot)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) FilterBackwardLET(opts *bind.FilterOpts) (*Agglayerbridgel2frometrogBackwardLETIterator, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.FilterLogs(opts, "BackwardLET")
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogBackwardLETIterator{contract: _Agglayerbridgel2frometrog.contract, event: "BackwardLET", logs: logs, sub: sub}, nil
}

// WatchBackwardLET is a free log subscription operation binding the contract event 0x2f8b2f36ca0d63af2efffe9d5e728a61c2bc273a3fc0574d12b558a41f149555.
//
// Solidity: event BackwardLET(uint256 previousDepositCount, bytes32 previousRoot, uint256 newDepositCount, bytes32 newRoot)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) WatchBackwardLET(opts *bind.WatchOpts, sink chan<- *Agglayerbridgel2frometrogBackwardLET) (event.Subscription, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.WatchLogs(opts, "BackwardLET")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Agglayerbridgel2frometrogBackwardLET)
				if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "BackwardLET", log); err != nil {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) ParseBackwardLET(log types.Log) (*Agglayerbridgel2frometrogBackwardLET, error) {
	event := new(Agglayerbridgel2frometrogBackwardLET)
	if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "BackwardLET", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Agglayerbridgel2frometrogBridgeEventIterator is returned from FilterBridgeEvent and is used to iterate over the raw logs and unpacked data for BridgeEvent events raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogBridgeEventIterator struct {
	Event *Agglayerbridgel2frometrogBridgeEvent // Event containing the contract specifics and raw log

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
func (it *Agglayerbridgel2frometrogBridgeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Agglayerbridgel2frometrogBridgeEvent)
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
		it.Event = new(Agglayerbridgel2frometrogBridgeEvent)
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
func (it *Agglayerbridgel2frometrogBridgeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Agglayerbridgel2frometrogBridgeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Agglayerbridgel2frometrogBridgeEvent represents a BridgeEvent event raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogBridgeEvent struct {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) FilterBridgeEvent(opts *bind.FilterOpts) (*Agglayerbridgel2frometrogBridgeEventIterator, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.FilterLogs(opts, "BridgeEvent")
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogBridgeEventIterator{contract: _Agglayerbridgel2frometrog.contract, event: "BridgeEvent", logs: logs, sub: sub}, nil
}

// WatchBridgeEvent is a free log subscription operation binding the contract event 0x501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b.
//
// Solidity: event BridgeEvent(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata, uint32 depositCount)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) WatchBridgeEvent(opts *bind.WatchOpts, sink chan<- *Agglayerbridgel2frometrogBridgeEvent) (event.Subscription, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.WatchLogs(opts, "BridgeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Agglayerbridgel2frometrogBridgeEvent)
				if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "BridgeEvent", log); err != nil {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) ParseBridgeEvent(log types.Log) (*Agglayerbridgel2frometrogBridgeEvent, error) {
	event := new(Agglayerbridgel2frometrogBridgeEvent)
	if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "BridgeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Agglayerbridgel2frometrogClaimEventIterator is returned from FilterClaimEvent and is used to iterate over the raw logs and unpacked data for ClaimEvent events raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogClaimEventIterator struct {
	Event *Agglayerbridgel2frometrogClaimEvent // Event containing the contract specifics and raw log

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
func (it *Agglayerbridgel2frometrogClaimEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Agglayerbridgel2frometrogClaimEvent)
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
		it.Event = new(Agglayerbridgel2frometrogClaimEvent)
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
func (it *Agglayerbridgel2frometrogClaimEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Agglayerbridgel2frometrogClaimEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Agglayerbridgel2frometrogClaimEvent represents a ClaimEvent event raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogClaimEvent struct {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) FilterClaimEvent(opts *bind.FilterOpts) (*Agglayerbridgel2frometrogClaimEventIterator, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.FilterLogs(opts, "ClaimEvent")
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogClaimEventIterator{contract: _Agglayerbridgel2frometrog.contract, event: "ClaimEvent", logs: logs, sub: sub}, nil
}

// WatchClaimEvent is a free log subscription operation binding the contract event 0x1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d.
//
// Solidity: event ClaimEvent(uint256 globalIndex, uint32 originNetwork, address originAddress, address destinationAddress, uint256 amount)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) WatchClaimEvent(opts *bind.WatchOpts, sink chan<- *Agglayerbridgel2frometrogClaimEvent) (event.Subscription, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.WatchLogs(opts, "ClaimEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Agglayerbridgel2frometrogClaimEvent)
				if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "ClaimEvent", log); err != nil {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) ParseClaimEvent(log types.Log) (*Agglayerbridgel2frometrogClaimEvent, error) {
	event := new(Agglayerbridgel2frometrogClaimEvent)
	if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "ClaimEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Agglayerbridgel2frometrogDetailedClaimEventIterator is returned from FilterDetailedClaimEvent and is used to iterate over the raw logs and unpacked data for DetailedClaimEvent events raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogDetailedClaimEventIterator struct {
	Event *Agglayerbridgel2frometrogDetailedClaimEvent // Event containing the contract specifics and raw log

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
func (it *Agglayerbridgel2frometrogDetailedClaimEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Agglayerbridgel2frometrogDetailedClaimEvent)
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
		it.Event = new(Agglayerbridgel2frometrogDetailedClaimEvent)
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
func (it *Agglayerbridgel2frometrogDetailedClaimEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Agglayerbridgel2frometrogDetailedClaimEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Agglayerbridgel2frometrogDetailedClaimEvent represents a DetailedClaimEvent event raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogDetailedClaimEvent struct {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) FilterDetailedClaimEvent(opts *bind.FilterOpts, globalIndex []*big.Int, destinationAddress []common.Address) (*Agglayerbridgel2frometrogDetailedClaimEventIterator, error) {

	var globalIndexRule []interface{}
	for _, globalIndexItem := range globalIndex {
		globalIndexRule = append(globalIndexRule, globalIndexItem)
	}

	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	logs, sub, err := _Agglayerbridgel2frometrog.contract.FilterLogs(opts, "DetailedClaimEvent", globalIndexRule, destinationAddressRule)
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogDetailedClaimEventIterator{contract: _Agglayerbridgel2frometrog.contract, event: "DetailedClaimEvent", logs: logs, sub: sub}, nil
}

// WatchDetailedClaimEvent is a free log subscription operation binding the contract event 0xfc09ab352fe828dc1fbb0cad35e812349ba00649f10067f83eafa8d79b161d81.
//
// Solidity: event DetailedClaimEvent(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 indexed globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint8 leafType, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address indexed destinationAddress, uint256 amount, bytes metadata)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) WatchDetailedClaimEvent(opts *bind.WatchOpts, sink chan<- *Agglayerbridgel2frometrogDetailedClaimEvent, globalIndex []*big.Int, destinationAddress []common.Address) (event.Subscription, error) {

	var globalIndexRule []interface{}
	for _, globalIndexItem := range globalIndex {
		globalIndexRule = append(globalIndexRule, globalIndexItem)
	}

	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	logs, sub, err := _Agglayerbridgel2frometrog.contract.WatchLogs(opts, "DetailedClaimEvent", globalIndexRule, destinationAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Agglayerbridgel2frometrogDetailedClaimEvent)
				if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "DetailedClaimEvent", log); err != nil {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) ParseDetailedClaimEvent(log types.Log) (*Agglayerbridgel2frometrogDetailedClaimEvent, error) {
	event := new(Agglayerbridgel2frometrogDetailedClaimEvent)
	if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "DetailedClaimEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Agglayerbridgel2frometrogEmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogEmergencyStateActivatedIterator struct {
	Event *Agglayerbridgel2frometrogEmergencyStateActivated // Event containing the contract specifics and raw log

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
func (it *Agglayerbridgel2frometrogEmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Agglayerbridgel2frometrogEmergencyStateActivated)
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
		it.Event = new(Agglayerbridgel2frometrogEmergencyStateActivated)
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
func (it *Agglayerbridgel2frometrogEmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Agglayerbridgel2frometrogEmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Agglayerbridgel2frometrogEmergencyStateActivated represents a EmergencyStateActivated event raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogEmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*Agglayerbridgel2frometrogEmergencyStateActivatedIterator, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogEmergencyStateActivatedIterator{contract: _Agglayerbridgel2frometrog.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *Agglayerbridgel2frometrogEmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Agglayerbridgel2frometrogEmergencyStateActivated)
				if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) ParseEmergencyStateActivated(log types.Log) (*Agglayerbridgel2frometrogEmergencyStateActivated, error) {
	event := new(Agglayerbridgel2frometrogEmergencyStateActivated)
	if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Agglayerbridgel2frometrogEmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogEmergencyStateDeactivatedIterator struct {
	Event *Agglayerbridgel2frometrogEmergencyStateDeactivated // Event containing the contract specifics and raw log

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
func (it *Agglayerbridgel2frometrogEmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Agglayerbridgel2frometrogEmergencyStateDeactivated)
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
		it.Event = new(Agglayerbridgel2frometrogEmergencyStateDeactivated)
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
func (it *Agglayerbridgel2frometrogEmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Agglayerbridgel2frometrogEmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Agglayerbridgel2frometrogEmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogEmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*Agglayerbridgel2frometrogEmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogEmergencyStateDeactivatedIterator{contract: _Agglayerbridgel2frometrog.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *Agglayerbridgel2frometrogEmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Agglayerbridgel2frometrogEmergencyStateDeactivated)
				if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) ParseEmergencyStateDeactivated(log types.Log) (*Agglayerbridgel2frometrogEmergencyStateDeactivated, error) {
	event := new(Agglayerbridgel2frometrogEmergencyStateDeactivated)
	if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Agglayerbridgel2frometrogForwardLETIterator is returned from FilterForwardLET and is used to iterate over the raw logs and unpacked data for ForwardLET events raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogForwardLETIterator struct {
	Event *Agglayerbridgel2frometrogForwardLET // Event containing the contract specifics and raw log

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
func (it *Agglayerbridgel2frometrogForwardLETIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Agglayerbridgel2frometrogForwardLET)
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
		it.Event = new(Agglayerbridgel2frometrogForwardLET)
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
func (it *Agglayerbridgel2frometrogForwardLETIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Agglayerbridgel2frometrogForwardLETIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Agglayerbridgel2frometrogForwardLET represents a ForwardLET event raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogForwardLET struct {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) FilterForwardLET(opts *bind.FilterOpts) (*Agglayerbridgel2frometrogForwardLETIterator, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.FilterLogs(opts, "ForwardLET")
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogForwardLETIterator{contract: _Agglayerbridgel2frometrog.contract, event: "ForwardLET", logs: logs, sub: sub}, nil
}

// WatchForwardLET is a free log subscription operation binding the contract event 0xe90f3fcbcd86a778012608da0cd7622521c8dac9919dddf25a3fb634adb3d2c8.
//
// Solidity: event ForwardLET(uint256 previousDepositCount, bytes32 previousRoot, uint256 newDepositCount, bytes32 newRoot, bytes newLeaves)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) WatchForwardLET(opts *bind.WatchOpts, sink chan<- *Agglayerbridgel2frometrogForwardLET) (event.Subscription, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.WatchLogs(opts, "ForwardLET")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Agglayerbridgel2frometrogForwardLET)
				if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "ForwardLET", log); err != nil {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) ParseForwardLET(log types.Log) (*Agglayerbridgel2frometrogForwardLET, error) {
	event := new(Agglayerbridgel2frometrogForwardLET)
	if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "ForwardLET", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Agglayerbridgel2frometrogInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogInitializedIterator struct {
	Event *Agglayerbridgel2frometrogInitialized // Event containing the contract specifics and raw log

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
func (it *Agglayerbridgel2frometrogInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Agglayerbridgel2frometrogInitialized)
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
		it.Event = new(Agglayerbridgel2frometrogInitialized)
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
func (it *Agglayerbridgel2frometrogInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Agglayerbridgel2frometrogInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Agglayerbridgel2frometrogInitialized represents a Initialized event raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) FilterInitialized(opts *bind.FilterOpts) (*Agglayerbridgel2frometrogInitializedIterator, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogInitializedIterator{contract: _Agglayerbridgel2frometrog.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Agglayerbridgel2frometrogInitialized) (event.Subscription, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Agglayerbridgel2frometrogInitialized)
				if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) ParseInitialized(log types.Log) (*Agglayerbridgel2frometrogInitialized, error) {
	event := new(Agglayerbridgel2frometrogInitialized)
	if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Agglayerbridgel2frometrogMigrateLegacyTokenIterator is returned from FilterMigrateLegacyToken and is used to iterate over the raw logs and unpacked data for MigrateLegacyToken events raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogMigrateLegacyTokenIterator struct {
	Event *Agglayerbridgel2frometrogMigrateLegacyToken // Event containing the contract specifics and raw log

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
func (it *Agglayerbridgel2frometrogMigrateLegacyTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Agglayerbridgel2frometrogMigrateLegacyToken)
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
		it.Event = new(Agglayerbridgel2frometrogMigrateLegacyToken)
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
func (it *Agglayerbridgel2frometrogMigrateLegacyTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Agglayerbridgel2frometrogMigrateLegacyTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Agglayerbridgel2frometrogMigrateLegacyToken represents a MigrateLegacyToken event raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogMigrateLegacyToken struct {
	Sender              common.Address
	LegacyTokenAddress  common.Address
	UpdatedTokenAddress common.Address
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMigrateLegacyToken is a free log retrieval operation binding the contract event 0xb7f8fd4d1faf9b2929dc269f59c53e3a2bccc44e9950f33a568fcbcb37eb69a9.
//
// Solidity: event MigrateLegacyToken(address sender, address legacyTokenAddress, address updatedTokenAddress, uint256 amount)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) FilterMigrateLegacyToken(opts *bind.FilterOpts) (*Agglayerbridgel2frometrogMigrateLegacyTokenIterator, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.FilterLogs(opts, "MigrateLegacyToken")
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogMigrateLegacyTokenIterator{contract: _Agglayerbridgel2frometrog.contract, event: "MigrateLegacyToken", logs: logs, sub: sub}, nil
}

// WatchMigrateLegacyToken is a free log subscription operation binding the contract event 0xb7f8fd4d1faf9b2929dc269f59c53e3a2bccc44e9950f33a568fcbcb37eb69a9.
//
// Solidity: event MigrateLegacyToken(address sender, address legacyTokenAddress, address updatedTokenAddress, uint256 amount)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) WatchMigrateLegacyToken(opts *bind.WatchOpts, sink chan<- *Agglayerbridgel2frometrogMigrateLegacyToken) (event.Subscription, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.WatchLogs(opts, "MigrateLegacyToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Agglayerbridgel2frometrogMigrateLegacyToken)
				if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "MigrateLegacyToken", log); err != nil {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) ParseMigrateLegacyToken(log types.Log) (*Agglayerbridgel2frometrogMigrateLegacyToken, error) {
	event := new(Agglayerbridgel2frometrogMigrateLegacyToken)
	if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "MigrateLegacyToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Agglayerbridgel2frometrogNewWrappedTokenIterator is returned from FilterNewWrappedToken and is used to iterate over the raw logs and unpacked data for NewWrappedToken events raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogNewWrappedTokenIterator struct {
	Event *Agglayerbridgel2frometrogNewWrappedToken // Event containing the contract specifics and raw log

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
func (it *Agglayerbridgel2frometrogNewWrappedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Agglayerbridgel2frometrogNewWrappedToken)
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
		it.Event = new(Agglayerbridgel2frometrogNewWrappedToken)
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
func (it *Agglayerbridgel2frometrogNewWrappedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Agglayerbridgel2frometrogNewWrappedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Agglayerbridgel2frometrogNewWrappedToken represents a NewWrappedToken event raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogNewWrappedToken struct {
	OriginNetwork       uint32
	OriginTokenAddress  common.Address
	WrappedTokenAddress common.Address
	Metadata            []byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewWrappedToken is a free log retrieval operation binding the contract event 0x490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a39.
//
// Solidity: event NewWrappedToken(uint32 originNetwork, address originTokenAddress, address wrappedTokenAddress, bytes metadata)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) FilterNewWrappedToken(opts *bind.FilterOpts) (*Agglayerbridgel2frometrogNewWrappedTokenIterator, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.FilterLogs(opts, "NewWrappedToken")
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogNewWrappedTokenIterator{contract: _Agglayerbridgel2frometrog.contract, event: "NewWrappedToken", logs: logs, sub: sub}, nil
}

// WatchNewWrappedToken is a free log subscription operation binding the contract event 0x490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a39.
//
// Solidity: event NewWrappedToken(uint32 originNetwork, address originTokenAddress, address wrappedTokenAddress, bytes metadata)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) WatchNewWrappedToken(opts *bind.WatchOpts, sink chan<- *Agglayerbridgel2frometrogNewWrappedToken) (event.Subscription, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.WatchLogs(opts, "NewWrappedToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Agglayerbridgel2frometrogNewWrappedToken)
				if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "NewWrappedToken", log); err != nil {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) ParseNewWrappedToken(log types.Log) (*Agglayerbridgel2frometrogNewWrappedToken, error) {
	event := new(Agglayerbridgel2frometrogNewWrappedToken)
	if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "NewWrappedToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Agglayerbridgel2frometrogRemoveLegacySovereignTokenAddressIterator is returned from FilterRemoveLegacySovereignTokenAddress and is used to iterate over the raw logs and unpacked data for RemoveLegacySovereignTokenAddress events raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogRemoveLegacySovereignTokenAddressIterator struct {
	Event *Agglayerbridgel2frometrogRemoveLegacySovereignTokenAddress // Event containing the contract specifics and raw log

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
func (it *Agglayerbridgel2frometrogRemoveLegacySovereignTokenAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Agglayerbridgel2frometrogRemoveLegacySovereignTokenAddress)
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
		it.Event = new(Agglayerbridgel2frometrogRemoveLegacySovereignTokenAddress)
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
func (it *Agglayerbridgel2frometrogRemoveLegacySovereignTokenAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Agglayerbridgel2frometrogRemoveLegacySovereignTokenAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Agglayerbridgel2frometrogRemoveLegacySovereignTokenAddress represents a RemoveLegacySovereignTokenAddress event raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogRemoveLegacySovereignTokenAddress struct {
	SovereignTokenAddress common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRemoveLegacySovereignTokenAddress is a free log retrieval operation binding the contract event 0xc2ae0bd0ec0fd0352bfe5bacac49637af342c1e40f1b80a7f74440dc7fe3f063.
//
// Solidity: event RemoveLegacySovereignTokenAddress(address sovereignTokenAddress)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) FilterRemoveLegacySovereignTokenAddress(opts *bind.FilterOpts) (*Agglayerbridgel2frometrogRemoveLegacySovereignTokenAddressIterator, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.FilterLogs(opts, "RemoveLegacySovereignTokenAddress")
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogRemoveLegacySovereignTokenAddressIterator{contract: _Agglayerbridgel2frometrog.contract, event: "RemoveLegacySovereignTokenAddress", logs: logs, sub: sub}, nil
}

// WatchRemoveLegacySovereignTokenAddress is a free log subscription operation binding the contract event 0xc2ae0bd0ec0fd0352bfe5bacac49637af342c1e40f1b80a7f74440dc7fe3f063.
//
// Solidity: event RemoveLegacySovereignTokenAddress(address sovereignTokenAddress)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) WatchRemoveLegacySovereignTokenAddress(opts *bind.WatchOpts, sink chan<- *Agglayerbridgel2frometrogRemoveLegacySovereignTokenAddress) (event.Subscription, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.WatchLogs(opts, "RemoveLegacySovereignTokenAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Agglayerbridgel2frometrogRemoveLegacySovereignTokenAddress)
				if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "RemoveLegacySovereignTokenAddress", log); err != nil {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) ParseRemoveLegacySovereignTokenAddress(log types.Log) (*Agglayerbridgel2frometrogRemoveLegacySovereignTokenAddress, error) {
	event := new(Agglayerbridgel2frometrogRemoveLegacySovereignTokenAddress)
	if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "RemoveLegacySovereignTokenAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Agglayerbridgel2frometrogSetBridgeManagerIterator is returned from FilterSetBridgeManager and is used to iterate over the raw logs and unpacked data for SetBridgeManager events raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogSetBridgeManagerIterator struct {
	Event *Agglayerbridgel2frometrogSetBridgeManager // Event containing the contract specifics and raw log

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
func (it *Agglayerbridgel2frometrogSetBridgeManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Agglayerbridgel2frometrogSetBridgeManager)
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
		it.Event = new(Agglayerbridgel2frometrogSetBridgeManager)
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
func (it *Agglayerbridgel2frometrogSetBridgeManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Agglayerbridgel2frometrogSetBridgeManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Agglayerbridgel2frometrogSetBridgeManager represents a SetBridgeManager event raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogSetBridgeManager struct {
	BridgeManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetBridgeManager is a free log retrieval operation binding the contract event 0x32cf74f8a6d5f88593984d2cd52be5592bfa6884f5896175801a5069ef09cd67.
//
// Solidity: event SetBridgeManager(address bridgeManager)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) FilterSetBridgeManager(opts *bind.FilterOpts) (*Agglayerbridgel2frometrogSetBridgeManagerIterator, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.FilterLogs(opts, "SetBridgeManager")
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogSetBridgeManagerIterator{contract: _Agglayerbridgel2frometrog.contract, event: "SetBridgeManager", logs: logs, sub: sub}, nil
}

// WatchSetBridgeManager is a free log subscription operation binding the contract event 0x32cf74f8a6d5f88593984d2cd52be5592bfa6884f5896175801a5069ef09cd67.
//
// Solidity: event SetBridgeManager(address bridgeManager)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) WatchSetBridgeManager(opts *bind.WatchOpts, sink chan<- *Agglayerbridgel2frometrogSetBridgeManager) (event.Subscription, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.WatchLogs(opts, "SetBridgeManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Agglayerbridgel2frometrogSetBridgeManager)
				if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "SetBridgeManager", log); err != nil {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) ParseSetBridgeManager(log types.Log) (*Agglayerbridgel2frometrogSetBridgeManager, error) {
	event := new(Agglayerbridgel2frometrogSetBridgeManager)
	if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "SetBridgeManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Agglayerbridgel2frometrogSetClaimIterator is returned from FilterSetClaim and is used to iterate over the raw logs and unpacked data for SetClaim events raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogSetClaimIterator struct {
	Event *Agglayerbridgel2frometrogSetClaim // Event containing the contract specifics and raw log

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
func (it *Agglayerbridgel2frometrogSetClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Agglayerbridgel2frometrogSetClaim)
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
		it.Event = new(Agglayerbridgel2frometrogSetClaim)
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
func (it *Agglayerbridgel2frometrogSetClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Agglayerbridgel2frometrogSetClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Agglayerbridgel2frometrogSetClaim represents a SetClaim event raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogSetClaim struct {
	GlobalIndex [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetClaim is a free log retrieval operation binding the contract event 0x8b65fb527138ff0fedc6b8e75e5a12268c71a1231b370b03209f7a286746823c.
//
// Solidity: event SetClaim(bytes32 globalIndex)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) FilterSetClaim(opts *bind.FilterOpts) (*Agglayerbridgel2frometrogSetClaimIterator, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.FilterLogs(opts, "SetClaim")
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogSetClaimIterator{contract: _Agglayerbridgel2frometrog.contract, event: "SetClaim", logs: logs, sub: sub}, nil
}

// WatchSetClaim is a free log subscription operation binding the contract event 0x8b65fb527138ff0fedc6b8e75e5a12268c71a1231b370b03209f7a286746823c.
//
// Solidity: event SetClaim(bytes32 globalIndex)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) WatchSetClaim(opts *bind.WatchOpts, sink chan<- *Agglayerbridgel2frometrogSetClaim) (event.Subscription, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.WatchLogs(opts, "SetClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Agglayerbridgel2frometrogSetClaim)
				if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "SetClaim", log); err != nil {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) ParseSetClaim(log types.Log) (*Agglayerbridgel2frometrogSetClaim, error) {
	event := new(Agglayerbridgel2frometrogSetClaim)
	if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "SetClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Agglayerbridgel2frometrogSetLocalBalanceTreeIterator is returned from FilterSetLocalBalanceTree and is used to iterate over the raw logs and unpacked data for SetLocalBalanceTree events raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogSetLocalBalanceTreeIterator struct {
	Event *Agglayerbridgel2frometrogSetLocalBalanceTree // Event containing the contract specifics and raw log

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
func (it *Agglayerbridgel2frometrogSetLocalBalanceTreeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Agglayerbridgel2frometrogSetLocalBalanceTree)
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
		it.Event = new(Agglayerbridgel2frometrogSetLocalBalanceTree)
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
func (it *Agglayerbridgel2frometrogSetLocalBalanceTreeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Agglayerbridgel2frometrogSetLocalBalanceTreeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Agglayerbridgel2frometrogSetLocalBalanceTree represents a SetLocalBalanceTree event raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogSetLocalBalanceTree struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
	NewAmount          *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSetLocalBalanceTree is a free log retrieval operation binding the contract event 0x73a03a6d9efc14b9f22a3af967e98580549eb76b1113b6a09a57ce1dae36ecd2.
//
// Solidity: event SetLocalBalanceTree(uint32 indexed originNetwork, address indexed originTokenAddress, uint256 newAmount)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) FilterSetLocalBalanceTree(opts *bind.FilterOpts, originNetwork []uint32, originTokenAddress []common.Address) (*Agglayerbridgel2frometrogSetLocalBalanceTreeIterator, error) {

	var originNetworkRule []interface{}
	for _, originNetworkItem := range originNetwork {
		originNetworkRule = append(originNetworkRule, originNetworkItem)
	}
	var originTokenAddressRule []interface{}
	for _, originTokenAddressItem := range originTokenAddress {
		originTokenAddressRule = append(originTokenAddressRule, originTokenAddressItem)
	}

	logs, sub, err := _Agglayerbridgel2frometrog.contract.FilterLogs(opts, "SetLocalBalanceTree", originNetworkRule, originTokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogSetLocalBalanceTreeIterator{contract: _Agglayerbridgel2frometrog.contract, event: "SetLocalBalanceTree", logs: logs, sub: sub}, nil
}

// WatchSetLocalBalanceTree is a free log subscription operation binding the contract event 0x73a03a6d9efc14b9f22a3af967e98580549eb76b1113b6a09a57ce1dae36ecd2.
//
// Solidity: event SetLocalBalanceTree(uint32 indexed originNetwork, address indexed originTokenAddress, uint256 newAmount)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) WatchSetLocalBalanceTree(opts *bind.WatchOpts, sink chan<- *Agglayerbridgel2frometrogSetLocalBalanceTree, originNetwork []uint32, originTokenAddress []common.Address) (event.Subscription, error) {

	var originNetworkRule []interface{}
	for _, originNetworkItem := range originNetwork {
		originNetworkRule = append(originNetworkRule, originNetworkItem)
	}
	var originTokenAddressRule []interface{}
	for _, originTokenAddressItem := range originTokenAddress {
		originTokenAddressRule = append(originTokenAddressRule, originTokenAddressItem)
	}

	logs, sub, err := _Agglayerbridgel2frometrog.contract.WatchLogs(opts, "SetLocalBalanceTree", originNetworkRule, originTokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Agglayerbridgel2frometrogSetLocalBalanceTree)
				if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "SetLocalBalanceTree", log); err != nil {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) ParseSetLocalBalanceTree(log types.Log) (*Agglayerbridgel2frometrogSetLocalBalanceTree, error) {
	event := new(Agglayerbridgel2frometrogSetLocalBalanceTree)
	if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "SetLocalBalanceTree", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Agglayerbridgel2frometrogSetSovereignTokenAddressIterator is returned from FilterSetSovereignTokenAddress and is used to iterate over the raw logs and unpacked data for SetSovereignTokenAddress events raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogSetSovereignTokenAddressIterator struct {
	Event *Agglayerbridgel2frometrogSetSovereignTokenAddress // Event containing the contract specifics and raw log

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
func (it *Agglayerbridgel2frometrogSetSovereignTokenAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Agglayerbridgel2frometrogSetSovereignTokenAddress)
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
		it.Event = new(Agglayerbridgel2frometrogSetSovereignTokenAddress)
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
func (it *Agglayerbridgel2frometrogSetSovereignTokenAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Agglayerbridgel2frometrogSetSovereignTokenAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Agglayerbridgel2frometrogSetSovereignTokenAddress represents a SetSovereignTokenAddress event raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogSetSovereignTokenAddress struct {
	OriginNetwork         uint32
	OriginTokenAddress    common.Address
	SovereignTokenAddress common.Address
	IsNotMintable         bool
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetSovereignTokenAddress is a free log retrieval operation binding the contract event 0xdbe8a5da6a7a916d9adfda9160167a0f8a3da415ee6610e810e753853597fce7.
//
// Solidity: event SetSovereignTokenAddress(uint32 originNetwork, address originTokenAddress, address sovereignTokenAddress, bool isNotMintable)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) FilterSetSovereignTokenAddress(opts *bind.FilterOpts) (*Agglayerbridgel2frometrogSetSovereignTokenAddressIterator, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.FilterLogs(opts, "SetSovereignTokenAddress")
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogSetSovereignTokenAddressIterator{contract: _Agglayerbridgel2frometrog.contract, event: "SetSovereignTokenAddress", logs: logs, sub: sub}, nil
}

// WatchSetSovereignTokenAddress is a free log subscription operation binding the contract event 0xdbe8a5da6a7a916d9adfda9160167a0f8a3da415ee6610e810e753853597fce7.
//
// Solidity: event SetSovereignTokenAddress(uint32 originNetwork, address originTokenAddress, address sovereignTokenAddress, bool isNotMintable)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) WatchSetSovereignTokenAddress(opts *bind.WatchOpts, sink chan<- *Agglayerbridgel2frometrogSetSovereignTokenAddress) (event.Subscription, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.WatchLogs(opts, "SetSovereignTokenAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Agglayerbridgel2frometrogSetSovereignTokenAddress)
				if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "SetSovereignTokenAddress", log); err != nil {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) ParseSetSovereignTokenAddress(log types.Log) (*Agglayerbridgel2frometrogSetSovereignTokenAddress, error) {
	event := new(Agglayerbridgel2frometrogSetSovereignTokenAddress)
	if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "SetSovereignTokenAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Agglayerbridgel2frometrogSetSovereignWETHAddressIterator is returned from FilterSetSovereignWETHAddress and is used to iterate over the raw logs and unpacked data for SetSovereignWETHAddress events raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogSetSovereignWETHAddressIterator struct {
	Event *Agglayerbridgel2frometrogSetSovereignWETHAddress // Event containing the contract specifics and raw log

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
func (it *Agglayerbridgel2frometrogSetSovereignWETHAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Agglayerbridgel2frometrogSetSovereignWETHAddress)
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
		it.Event = new(Agglayerbridgel2frometrogSetSovereignWETHAddress)
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
func (it *Agglayerbridgel2frometrogSetSovereignWETHAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Agglayerbridgel2frometrogSetSovereignWETHAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Agglayerbridgel2frometrogSetSovereignWETHAddress represents a SetSovereignWETHAddress event raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogSetSovereignWETHAddress struct {
	SovereignWETHTokenAddress common.Address
	IsNotMintable             bool
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterSetSovereignWETHAddress is a free log retrieval operation binding the contract event 0xc7318b7ed6ba4f2908a3de396d8ab49b1dadb55db5b55123247a401f29ff8d82.
//
// Solidity: event SetSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) FilterSetSovereignWETHAddress(opts *bind.FilterOpts) (*Agglayerbridgel2frometrogSetSovereignWETHAddressIterator, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.FilterLogs(opts, "SetSovereignWETHAddress")
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogSetSovereignWETHAddressIterator{contract: _Agglayerbridgel2frometrog.contract, event: "SetSovereignWETHAddress", logs: logs, sub: sub}, nil
}

// WatchSetSovereignWETHAddress is a free log subscription operation binding the contract event 0xc7318b7ed6ba4f2908a3de396d8ab49b1dadb55db5b55123247a401f29ff8d82.
//
// Solidity: event SetSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) WatchSetSovereignWETHAddress(opts *bind.WatchOpts, sink chan<- *Agglayerbridgel2frometrogSetSovereignWETHAddress) (event.Subscription, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.WatchLogs(opts, "SetSovereignWETHAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Agglayerbridgel2frometrogSetSovereignWETHAddress)
				if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "SetSovereignWETHAddress", log); err != nil {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) ParseSetSovereignWETHAddress(log types.Log) (*Agglayerbridgel2frometrogSetSovereignWETHAddress, error) {
	event := new(Agglayerbridgel2frometrogSetSovereignWETHAddress)
	if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "SetSovereignWETHAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Agglayerbridgel2frometrogTransferEmergencyBridgePauserRoleIterator is returned from FilterTransferEmergencyBridgePauserRole and is used to iterate over the raw logs and unpacked data for TransferEmergencyBridgePauserRole events raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogTransferEmergencyBridgePauserRoleIterator struct {
	Event *Agglayerbridgel2frometrogTransferEmergencyBridgePauserRole // Event containing the contract specifics and raw log

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
func (it *Agglayerbridgel2frometrogTransferEmergencyBridgePauserRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Agglayerbridgel2frometrogTransferEmergencyBridgePauserRole)
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
		it.Event = new(Agglayerbridgel2frometrogTransferEmergencyBridgePauserRole)
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
func (it *Agglayerbridgel2frometrogTransferEmergencyBridgePauserRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Agglayerbridgel2frometrogTransferEmergencyBridgePauserRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Agglayerbridgel2frometrogTransferEmergencyBridgePauserRole represents a TransferEmergencyBridgePauserRole event raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogTransferEmergencyBridgePauserRole struct {
	CurrentEmergencyBridgePauser common.Address
	NewEmergencyBridgePauser     common.Address
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterTransferEmergencyBridgePauserRole is a free log retrieval operation binding the contract event 0xb27de219766f47b82684842855ba6130b6dbf288ac66d1c3509e7bf17f4e925a.
//
// Solidity: event TransferEmergencyBridgePauserRole(address currentEmergencyBridgePauser, address newEmergencyBridgePauser)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) FilterTransferEmergencyBridgePauserRole(opts *bind.FilterOpts) (*Agglayerbridgel2frometrogTransferEmergencyBridgePauserRoleIterator, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.FilterLogs(opts, "TransferEmergencyBridgePauserRole")
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogTransferEmergencyBridgePauserRoleIterator{contract: _Agglayerbridgel2frometrog.contract, event: "TransferEmergencyBridgePauserRole", logs: logs, sub: sub}, nil
}

// WatchTransferEmergencyBridgePauserRole is a free log subscription operation binding the contract event 0xb27de219766f47b82684842855ba6130b6dbf288ac66d1c3509e7bf17f4e925a.
//
// Solidity: event TransferEmergencyBridgePauserRole(address currentEmergencyBridgePauser, address newEmergencyBridgePauser)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) WatchTransferEmergencyBridgePauserRole(opts *bind.WatchOpts, sink chan<- *Agglayerbridgel2frometrogTransferEmergencyBridgePauserRole) (event.Subscription, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.WatchLogs(opts, "TransferEmergencyBridgePauserRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Agglayerbridgel2frometrogTransferEmergencyBridgePauserRole)
				if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "TransferEmergencyBridgePauserRole", log); err != nil {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) ParseTransferEmergencyBridgePauserRole(log types.Log) (*Agglayerbridgel2frometrogTransferEmergencyBridgePauserRole, error) {
	event := new(Agglayerbridgel2frometrogTransferEmergencyBridgePauserRole)
	if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "TransferEmergencyBridgePauserRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Agglayerbridgel2frometrogTransferEmergencyBridgeUnpauserRoleIterator is returned from FilterTransferEmergencyBridgeUnpauserRole and is used to iterate over the raw logs and unpacked data for TransferEmergencyBridgeUnpauserRole events raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogTransferEmergencyBridgeUnpauserRoleIterator struct {
	Event *Agglayerbridgel2frometrogTransferEmergencyBridgeUnpauserRole // Event containing the contract specifics and raw log

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
func (it *Agglayerbridgel2frometrogTransferEmergencyBridgeUnpauserRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Agglayerbridgel2frometrogTransferEmergencyBridgeUnpauserRole)
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
		it.Event = new(Agglayerbridgel2frometrogTransferEmergencyBridgeUnpauserRole)
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
func (it *Agglayerbridgel2frometrogTransferEmergencyBridgeUnpauserRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Agglayerbridgel2frometrogTransferEmergencyBridgeUnpauserRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Agglayerbridgel2frometrogTransferEmergencyBridgeUnpauserRole represents a TransferEmergencyBridgeUnpauserRole event raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogTransferEmergencyBridgeUnpauserRole struct {
	CurrentEmergencyBridgeUnpauser common.Address
	NewEmergencyBridgeUnpauser     common.Address
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterTransferEmergencyBridgeUnpauserRole is a free log retrieval operation binding the contract event 0xf01a62a06940517bbc898dec8c75794b9feabcd2d263c8de823b36dbbeb8779b.
//
// Solidity: event TransferEmergencyBridgeUnpauserRole(address currentEmergencyBridgeUnpauser, address newEmergencyBridgeUnpauser)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) FilterTransferEmergencyBridgeUnpauserRole(opts *bind.FilterOpts) (*Agglayerbridgel2frometrogTransferEmergencyBridgeUnpauserRoleIterator, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.FilterLogs(opts, "TransferEmergencyBridgeUnpauserRole")
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogTransferEmergencyBridgeUnpauserRoleIterator{contract: _Agglayerbridgel2frometrog.contract, event: "TransferEmergencyBridgeUnpauserRole", logs: logs, sub: sub}, nil
}

// WatchTransferEmergencyBridgeUnpauserRole is a free log subscription operation binding the contract event 0xf01a62a06940517bbc898dec8c75794b9feabcd2d263c8de823b36dbbeb8779b.
//
// Solidity: event TransferEmergencyBridgeUnpauserRole(address currentEmergencyBridgeUnpauser, address newEmergencyBridgeUnpauser)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) WatchTransferEmergencyBridgeUnpauserRole(opts *bind.WatchOpts, sink chan<- *Agglayerbridgel2frometrogTransferEmergencyBridgeUnpauserRole) (event.Subscription, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.WatchLogs(opts, "TransferEmergencyBridgeUnpauserRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Agglayerbridgel2frometrogTransferEmergencyBridgeUnpauserRole)
				if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "TransferEmergencyBridgeUnpauserRole", log); err != nil {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) ParseTransferEmergencyBridgeUnpauserRole(log types.Log) (*Agglayerbridgel2frometrogTransferEmergencyBridgeUnpauserRole, error) {
	event := new(Agglayerbridgel2frometrogTransferEmergencyBridgeUnpauserRole)
	if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "TransferEmergencyBridgeUnpauserRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Agglayerbridgel2frometrogTransferProxiedTokensManagerRoleIterator is returned from FilterTransferProxiedTokensManagerRole and is used to iterate over the raw logs and unpacked data for TransferProxiedTokensManagerRole events raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogTransferProxiedTokensManagerRoleIterator struct {
	Event *Agglayerbridgel2frometrogTransferProxiedTokensManagerRole // Event containing the contract specifics and raw log

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
func (it *Agglayerbridgel2frometrogTransferProxiedTokensManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Agglayerbridgel2frometrogTransferProxiedTokensManagerRole)
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
		it.Event = new(Agglayerbridgel2frometrogTransferProxiedTokensManagerRole)
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
func (it *Agglayerbridgel2frometrogTransferProxiedTokensManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Agglayerbridgel2frometrogTransferProxiedTokensManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Agglayerbridgel2frometrogTransferProxiedTokensManagerRole represents a TransferProxiedTokensManagerRole event raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogTransferProxiedTokensManagerRole struct {
	CurrentProxiedTokensManager common.Address
	NewProxiedTokensManager     common.Address
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterTransferProxiedTokensManagerRole is a free log retrieval operation binding the contract event 0x0a34baa3feb299aef9c05cb59c6e0c8e7c0bcc65cbf0a647e7a7c8a2411591e2.
//
// Solidity: event TransferProxiedTokensManagerRole(address currentProxiedTokensManager, address newProxiedTokensManager)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) FilterTransferProxiedTokensManagerRole(opts *bind.FilterOpts) (*Agglayerbridgel2frometrogTransferProxiedTokensManagerRoleIterator, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.FilterLogs(opts, "TransferProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogTransferProxiedTokensManagerRoleIterator{contract: _Agglayerbridgel2frometrog.contract, event: "TransferProxiedTokensManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferProxiedTokensManagerRole is a free log subscription operation binding the contract event 0x0a34baa3feb299aef9c05cb59c6e0c8e7c0bcc65cbf0a647e7a7c8a2411591e2.
//
// Solidity: event TransferProxiedTokensManagerRole(address currentProxiedTokensManager, address newProxiedTokensManager)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) WatchTransferProxiedTokensManagerRole(opts *bind.WatchOpts, sink chan<- *Agglayerbridgel2frometrogTransferProxiedTokensManagerRole) (event.Subscription, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.WatchLogs(opts, "TransferProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Agglayerbridgel2frometrogTransferProxiedTokensManagerRole)
				if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "TransferProxiedTokensManagerRole", log); err != nil {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) ParseTransferProxiedTokensManagerRole(log types.Log) (*Agglayerbridgel2frometrogTransferProxiedTokensManagerRole, error) {
	event := new(Agglayerbridgel2frometrogTransferProxiedTokensManagerRole)
	if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "TransferProxiedTokensManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Agglayerbridgel2frometrogUpdatedClaimedGlobalIndexHashChainIterator is returned from FilterUpdatedClaimedGlobalIndexHashChain and is used to iterate over the raw logs and unpacked data for UpdatedClaimedGlobalIndexHashChain events raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogUpdatedClaimedGlobalIndexHashChainIterator struct {
	Event *Agglayerbridgel2frometrogUpdatedClaimedGlobalIndexHashChain // Event containing the contract specifics and raw log

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
func (it *Agglayerbridgel2frometrogUpdatedClaimedGlobalIndexHashChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Agglayerbridgel2frometrogUpdatedClaimedGlobalIndexHashChain)
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
		it.Event = new(Agglayerbridgel2frometrogUpdatedClaimedGlobalIndexHashChain)
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
func (it *Agglayerbridgel2frometrogUpdatedClaimedGlobalIndexHashChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Agglayerbridgel2frometrogUpdatedClaimedGlobalIndexHashChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Agglayerbridgel2frometrogUpdatedClaimedGlobalIndexHashChain represents a UpdatedClaimedGlobalIndexHashChain event raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogUpdatedClaimedGlobalIndexHashChain struct {
	ClaimedGlobalIndex             [32]byte
	NewClaimedGlobalIndexHashChain [32]byte
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterUpdatedClaimedGlobalIndexHashChain is a free log retrieval operation binding the contract event 0x3e5936f910a78eb5181813a939c8d4c3e4d85f87943f659380d82ac6221b0e92.
//
// Solidity: event UpdatedClaimedGlobalIndexHashChain(bytes32 claimedGlobalIndex, bytes32 newClaimedGlobalIndexHashChain)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) FilterUpdatedClaimedGlobalIndexHashChain(opts *bind.FilterOpts) (*Agglayerbridgel2frometrogUpdatedClaimedGlobalIndexHashChainIterator, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.FilterLogs(opts, "UpdatedClaimedGlobalIndexHashChain")
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogUpdatedClaimedGlobalIndexHashChainIterator{contract: _Agglayerbridgel2frometrog.contract, event: "UpdatedClaimedGlobalIndexHashChain", logs: logs, sub: sub}, nil
}

// WatchUpdatedClaimedGlobalIndexHashChain is a free log subscription operation binding the contract event 0x3e5936f910a78eb5181813a939c8d4c3e4d85f87943f659380d82ac6221b0e92.
//
// Solidity: event UpdatedClaimedGlobalIndexHashChain(bytes32 claimedGlobalIndex, bytes32 newClaimedGlobalIndexHashChain)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) WatchUpdatedClaimedGlobalIndexHashChain(opts *bind.WatchOpts, sink chan<- *Agglayerbridgel2frometrogUpdatedClaimedGlobalIndexHashChain) (event.Subscription, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.WatchLogs(opts, "UpdatedClaimedGlobalIndexHashChain")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Agglayerbridgel2frometrogUpdatedClaimedGlobalIndexHashChain)
				if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "UpdatedClaimedGlobalIndexHashChain", log); err != nil {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) ParseUpdatedClaimedGlobalIndexHashChain(log types.Log) (*Agglayerbridgel2frometrogUpdatedClaimedGlobalIndexHashChain, error) {
	event := new(Agglayerbridgel2frometrogUpdatedClaimedGlobalIndexHashChain)
	if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "UpdatedClaimedGlobalIndexHashChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Agglayerbridgel2frometrogUpdatedUnsetGlobalIndexHashChainIterator is returned from FilterUpdatedUnsetGlobalIndexHashChain and is used to iterate over the raw logs and unpacked data for UpdatedUnsetGlobalIndexHashChain events raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogUpdatedUnsetGlobalIndexHashChainIterator struct {
	Event *Agglayerbridgel2frometrogUpdatedUnsetGlobalIndexHashChain // Event containing the contract specifics and raw log

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
func (it *Agglayerbridgel2frometrogUpdatedUnsetGlobalIndexHashChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Agglayerbridgel2frometrogUpdatedUnsetGlobalIndexHashChain)
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
		it.Event = new(Agglayerbridgel2frometrogUpdatedUnsetGlobalIndexHashChain)
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
func (it *Agglayerbridgel2frometrogUpdatedUnsetGlobalIndexHashChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Agglayerbridgel2frometrogUpdatedUnsetGlobalIndexHashChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Agglayerbridgel2frometrogUpdatedUnsetGlobalIndexHashChain represents a UpdatedUnsetGlobalIndexHashChain event raised by the Agglayerbridgel2frometrog contract.
type Agglayerbridgel2frometrogUpdatedUnsetGlobalIndexHashChain struct {
	UnsetGlobalIndex             [32]byte
	NewUnsetGlobalIndexHashChain [32]byte
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedUnsetGlobalIndexHashChain is a free log retrieval operation binding the contract event 0xc80e0aca446a59735359a7ae46124b57c47b892827642779bc6dafc84ba90b03.
//
// Solidity: event UpdatedUnsetGlobalIndexHashChain(bytes32 unsetGlobalIndex, bytes32 newUnsetGlobalIndexHashChain)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) FilterUpdatedUnsetGlobalIndexHashChain(opts *bind.FilterOpts) (*Agglayerbridgel2frometrogUpdatedUnsetGlobalIndexHashChainIterator, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.FilterLogs(opts, "UpdatedUnsetGlobalIndexHashChain")
	if err != nil {
		return nil, err
	}
	return &Agglayerbridgel2frometrogUpdatedUnsetGlobalIndexHashChainIterator{contract: _Agglayerbridgel2frometrog.contract, event: "UpdatedUnsetGlobalIndexHashChain", logs: logs, sub: sub}, nil
}

// WatchUpdatedUnsetGlobalIndexHashChain is a free log subscription operation binding the contract event 0xc80e0aca446a59735359a7ae46124b57c47b892827642779bc6dafc84ba90b03.
//
// Solidity: event UpdatedUnsetGlobalIndexHashChain(bytes32 unsetGlobalIndex, bytes32 newUnsetGlobalIndexHashChain)
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) WatchUpdatedUnsetGlobalIndexHashChain(opts *bind.WatchOpts, sink chan<- *Agglayerbridgel2frometrogUpdatedUnsetGlobalIndexHashChain) (event.Subscription, error) {

	logs, sub, err := _Agglayerbridgel2frometrog.contract.WatchLogs(opts, "UpdatedUnsetGlobalIndexHashChain")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Agglayerbridgel2frometrogUpdatedUnsetGlobalIndexHashChain)
				if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "UpdatedUnsetGlobalIndexHashChain", log); err != nil {
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
func (_Agglayerbridgel2frometrog *Agglayerbridgel2frometrogFilterer) ParseUpdatedUnsetGlobalIndexHashChain(log types.Log) (*Agglayerbridgel2frometrogUpdatedUnsetGlobalIndexHashChain, error) {
	event := new(Agglayerbridgel2frometrogUpdatedUnsetGlobalIndexHashChain)
	if err := _Agglayerbridgel2frometrog.contract.UnpackLog(event, "UpdatedUnsetGlobalIndexHashChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

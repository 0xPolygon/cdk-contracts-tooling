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

// Bridgel2sovereignchainMetaData contains all meta data concerning the Bridgel2sovereignchain contract.
var Bridgel2sovereignchainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountDoesNotMatchMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DestinationNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EtherTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTokenWrappedDeployment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputArraysLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSmtProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSovereignWETHAddressParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MsgValueNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeTokenIsEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoValueInMessagesOnGasTokenNetworks\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidBridgeManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyBridgeManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OriginNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyMapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyUpdated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotMapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotRemapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WETHRemappingNotSupportedOnGasTokenNetworks\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"depositCount\",\"type\":\"uint32\"}],\"name\":\"BridgeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"legacyTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"updatedTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MigrateLegacyToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wrappedTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"NewWrappedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignTokenAddress\",\"type\":\"address\"}],\"name\":\"RemoveLegacySovereignTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridgeManager\",\"type\":\"address\"}],\"name\":\"SetBridgeManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"SetSovereignTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignWETHTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"SetSovereignWETHAddress\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASE_INIT_BYTECODE_WRAPPED_TOKEN\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETHToken\",\"outputs\":[{\"internalType\":\"contractTokenWrapped\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"permitData\",\"type\":\"bytes\"}],\"name\":\"bridgeAsset\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountWETH\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessageWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"calculateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"calculateTokenWrapperAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimedBitMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"name\":\"getLeafValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenWrappedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_polygonRollupManager\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_bridgeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sovereignWETHAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_sovereignWETHAddressIsNotMintable\",\"type\":\"bool\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"leafIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"sourceBridgeNetwork\",\"type\":\"uint32\"}],\"name\":\"isClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdatedDepositCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"legacyTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"migrateLegacyToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"polygonRollupManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"precalculatedWrapperAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sovereignTokenAddress\",\"type\":\"address\"}],\"name\":\"removeLegacySovereignTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeManager\",\"type\":\"address\"}],\"name\":\"setBridgeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"originNetworks\",\"type\":\"uint32[]\"},{\"internalType\":\"address[]\",\"name\":\"originTokenAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"sovereignTokenAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"isNotMintable\",\"type\":\"bool[]\"}],\"name\":\"setMultipleSovereignTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sovereignTokenAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"setSovereignTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sovereignWETHTokenAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"setSovereignWETHAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenInfoToWrappedToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"verifyMerkleProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrappedAddress\",\"type\":\"address\"}],\"name\":\"wrappedAddressIsNotMintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wrappedTokenToTokenInfo\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506200001c6200002c565b620000266200002c565b620000ee565b600054610100900460ff1615620000995760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161015620000ec576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b61794280620000fe6000396000f3fe6080604052600436106102c65760003560e01c80638ed7e3f211610179578063c0f49163116100d6578063dbc169761161008a578063f5efcd7911610064578063f5efcd7914610965578063f811bff714610985578063fb570834146109a557600080fd5b8063dbc169761461034c578063eabd372a14610918578063ee25560b1461093857600080fd5b8063ccaa2d11116100bb578063ccaa2d11146108af578063cd586579146108cf578063d02103ca146108e257600080fd5b8063c0f491631461085f578063cc4616321461088f57600080fd5b8063b8b284d01161012d578063be5831c711610112578063be5831c7146107e5578063bf130d7f1461081f578063c00f14ab1461083f57600080fd5b8063b8b284d0146107a3578063bab161bf146107c357600080fd5b8063aaa13cc21161015e578063aaa13cc214610743578063b42f6b3a14610763578063b45869621461078357600080fd5b80638ed7e3f2146106f65780639e76158f1461072357600080fd5b80633e1970431161022757806379e2cf97116101db57806383c43a55116101c057806383c43a55146106a157806383f24403146106b65780638c0dd470146106d657600080fd5b806379e2cf971461064957806381b1c1741461065e57600080fd5b806357cfbee31161020c57806357cfbee3146105f45780635ca1e165146106145780637843298b1461062957600080fd5b80633e197043146104d75780634b2f336d146105c757600080fd5b806327aef4e81161027e578063318aee3d11610263578063318aee3d146103dc5780633c351e10146104605780633cbc795b1461048d57600080fd5b806327aef4e8146103965780632dfdf0b5146103b857600080fd5b80632072f6c5116102af5780632072f6c51461034c57806322e95f2c14610363578063240ff3781461038357600080fd5b806314cc01a0146102cb57806315064c9614610322575b600080fd5b3480156102d757600080fd5b5060a3546102f89073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561032e57600080fd5b5060685461033c9060ff1681565b6040519015158152602001610319565b34801561035857600080fd5b506103616109c5565b005b34801561036f57600080fd5b506102f861037e3660046149fe565b6109f7565b610361610391366004614a8c565b610a9a565b3480156103a257600080fd5b506103ab610b49565b6040516103199190614b74565b3480156103c457600080fd5b506103ce60535481565b604051908152602001610319565b3480156103e857600080fd5b5061042f6103f7366004614b8e565b606b6020526000908152604090205463ffffffff811690640100000000900473ffffffffffffffffffffffffffffffffffffffff1682565b6040805163ffffffff909316835273ffffffffffffffffffffffffffffffffffffffff909116602083015201610319565b34801561046c57600080fd5b50606d546102f89073ffffffffffffffffffffffffffffffffffffffff1681565b34801561049957600080fd5b50606d546104c29074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff9091168152602001610319565b3480156104e357600080fd5b506103ce6104f2366004614bba565b6040517fff0000000000000000000000000000000000000000000000000000000000000060f889901b1660208201527fffffffff0000000000000000000000000000000000000000000000000000000060e088811b821660218401527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606089811b821660258601529188901b909216603984015285901b16603d8201526051810183905260718101829052600090609101604051602081830303815290604052805190602001209050979650505050505050565b3480156105d357600080fd5b50606f546102f89073ffffffffffffffffffffffffffffffffffffffff1681565b34801561060057600080fd5b5061036161060f366004614db3565b610bd7565b34801561062057600080fd5b506103ce610d0d565b34801561063557600080fd5b506102f8610644366004614ebc565b610dea565b34801561065557600080fd5b50610361610e14565b34801561066a57600080fd5b506102f8610679366004614f05565b606a6020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b3480156106ad57600080fd5b506103ab610e4d565b3480156106c257600080fd5b506103ce6106d1366004614f30565b610e6c565b3480156106e257600080fd5b506103616106f1366004615006565b610f42565b34801561070257600080fd5b50606c546102f89073ffffffffffffffffffffffffffffffffffffffff1681565b34801561072f57600080fd5b5061036161073e3660046150d7565b61146c565b34801561074f57600080fd5b506102f861075e366004615103565b611676565b34801561076f57600080fd5b5061036161077e36600461519f565b611831565b34801561078f57600080fd5b5061036161079e366004614b8e565b611894565b3480156107af57600080fd5b506103616107be3660046151f9565b611b0f565b3480156107cf57600080fd5b506068546104c290610100900463ffffffff1681565b3480156107f157600080fd5b506068546104c290790100000000000000000000000000000000000000000000000000900463ffffffff1681565b34801561082b57600080fd5b5061036161083a36600461527c565b611bd4565b34801561084b57600080fd5b506103ab61085a366004614b8e565b611d2e565b34801561086b57600080fd5b5061033c61087a366004614b8e565b60a26020526000908152604090205460ff1681565b34801561089b57600080fd5b5061033c6108aa3660046152aa565b611d73565b3480156108bb57600080fd5b506103616108ca3660046152dd565b611dfd565b6103616108dd3660046153c9565b612477565b3480156108ee57600080fd5b506068546102f89065010000000000900473ffffffffffffffffffffffffffffffffffffffff1681565b34801561092457600080fd5b50610361610933366004614b8e565b6129ec565b34801561094457600080fd5b506103ce610953366004614f05565b60696020526000908152604090205481565b34801561097157600080fd5b506103616109803660046152dd565b612b03565b34801561099157600080fd5b506103616109a036600461545f565b612e92565b3480156109b157600080fd5b5061033c6109c03660046154f4565b612fe8565b6040517f376f0fe600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160e084901b7fffffffff0000000000000000000000000000000000000000000000000000000016602080830191909152606084901b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602483015282516018818403018152603890920183528151918101919091206000908152606a909152205473ffffffffffffffffffffffffffffffffffffffff165b92915050565b60685460ff1615610ad7576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3415801590610afd5750606f5473ffffffffffffffffffffffffffffffffffffffff1615155b15610b34576040517f6f625c4000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b42858534868686613000565b5050505050565b606e8054610b569061553c565b80601f0160208091040260200160405190810160405280929190818152602001828054610b829061553c565b8015610bcf5780601f10610ba457610100808354040283529160200191610bcf565b820191906000526020600020905b815481529060010190602001808311610bb257829003601f168201915b505050505081565b60a35473ffffffffffffffffffffffffffffffffffffffff163314610c28576040517faf6e71a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82518451141580610c3b57508151845114155b80610c4857508051845114155b15610c7f576040517f869e93ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8251811015610b4257610cfb858281518110610ca057610ca061558f565b6020026020010151858381518110610cba57610cba61558f565b6020026020010151858481518110610cd457610cd461558f565b6020026020010151858581518110610cee57610cee61558f565b60200260200101516130e3565b80610d05816155ed565b915050610c82565b605354600090819081805b6020811015610de1578083901c600116600103610d755760338160208110610d4257610d4261558f565b01546040805160208101929092528101859052606001604051602081830303815290604052805190602001209350610da2565b60408051602081018690529081018390526060016040516020818303038152906040528051906020012093505b60408051602081018490529081018390526060016040516020818303038152906040528051906020012091508080610dd9906155ed565b915050610d18565b50919392505050565b6000610e0c8484610dfa856133cf565b610e03866134df565b61075e876135e6565b949350505050565b605354606854790100000000000000000000000000000000000000000000000000900463ffffffff161015610e4b57610e4b6136d9565b565b60405180611ba00160405280611b668152602001615da7611b66913981565b600083815b6020811015610f3957600163ffffffff8516821c81169003610edc57848160208110610e9f57610e9f61558f565b602002013582604051602001610ebf929190918252602082015260400190565b604051602081830303815290604052805190602001209150610f27565b81858260208110610eef57610eef61558f565b6020020135604051602001610f0e929190918252602082015260400190565b6040516020818303038152906040528051906020012091505b80610f31816155ed565b915050610e71565b50949350505050565b600054610100900460ff1615808015610f625750600054600160ff909116105b80610f7c5750303b158015610f7c575060005460ff166001145b61100d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561106b57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b606880547fffffffffffffff000000000000000000000000000000000000000000000000ff1661010063ffffffff8d16027fffffffffffffff0000000000000000000000000000000000000000ffffffffff16176501000000000073ffffffffffffffffffffffffffffffffffffffff8a81169190910291909117909155606c80547fffffffffffffffffffffffff00000000000000000000000000000000000000009081168984161790915560a3805490911686831617905589166111cb5763ffffffff881615611169576040517f1a874c1200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff831615158061118f57506001821515145b156111c6576040517f1cdc46ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113f5565b606d805463ffffffff8a1674010000000000000000000000000000000000000000027fffffffffffffffff00000000000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff8c1617179055606e611238868261566b565b5073ffffffffffffffffffffffffffffffffffffffff831661137957811515600103611290576040517f1cdc46ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61132f6000801b601260405160200161131b91906060808252600d908201527f5772617070656420457468657200000000000000000000000000000000000000608082015260a0602082018190526004908201527f574554480000000000000000000000000000000000000000000000000000000060c082015260ff91909116604082015260e00190565b6040516020818303038152906040526137a9565b606f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff929092169190911790556113f5565b606f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8516908117909155600090815260a26020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168315151790555b6113fd61384b565b801561146057600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8083166000908152606b602090815260409182902082518084019093525463ffffffff811683526401000000009004909216918101829052906114f0576040517f828d566300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000606a60008360000151846020015160405160200161156692919060e09290921b7fffffffff0000000000000000000000000000000000000000000000000000000016825260601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016600482015260180190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181528151602092830120835290820192909252016000205473ffffffffffffffffffffffffffffffffffffffff9081169150841681036115fe576040517fe273c4a100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61160884846138ea565b6116138133856139c0565b6040805133815273ffffffffffffffffffffffffffffffffffffffff86811660208301528316818301526060810185905290517fb7f8fd4d1faf9b2929dc269f59c53e3a2bccc44e9950f33a568fcbcb37eb69a99181900360800190a150505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e087901b1660208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606086901b1660248201526000908190603801604051602081830303815290604052805190602001209050600060ff60f81b308360405180611ba00160405280611b668152602001615da7611b66913989898960405160200161172b93929190615785565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529082905261176792916020016157be565b604051602081830303815290604052805190602001206040516020016117ef94939291907fff0000000000000000000000000000000000000000000000000000000000000094909416845260609290921b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660018401526015830152603582015260550190565b604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018152919052805160209091012098975050505050505050565b60a35473ffffffffffffffffffffffffffffffffffffffff163314611882576040517faf6e71a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61188e848484846130e3565b50505050565b60a35473ffffffffffffffffffffffffffffffffffffffff1633146118e5576040517faf6e71a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8082166000908152606b6020908152604080832081518083018352905463ffffffff811680835264010000000090910490951681840181905291519094611999939092910160e09290921b7fffffffff0000000000000000000000000000000000000000000000000000000016825260601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016600482015260180190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152606a90935291205490915073ffffffffffffffffffffffffffffffffffffffff161580611a2757506000818152606a602052604090205473ffffffffffffffffffffffffffffffffffffffff8481169116145b15611a5e576040517fe0c897a700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83166000818152606b6020908152604080832080547fffffffffffffffff00000000000000000000000000000000000000000000000016905560a282529182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905590519182527fc2ae0bd0ec0fd0352bfe5bacac49637af342c1e40f1b80a7f74440dc7fe3f063910160405180910390a1505050565b60685460ff1615611b4c576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5473ffffffffffffffffffffffffffffffffffffffff16611b9b576040517fdde3cda700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f54611bbe9073ffffffffffffffffffffffffffffffffffffffff16856138ea565b611bcc868686868686613000565b505050505050565b60a35473ffffffffffffffffffffffffffffffffffffffff163314611c25576040517faf6e71a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606d5473ffffffffffffffffffffffffffffffffffffffff16611c74576040517f9968e22600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8416908117909155600081815260a2602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168515159081179091558251938452908301527fc7318b7ed6ba4f2908a3de396d8ab49b1dadb55db5b55123247a401f29ff8d82910160405180910390a15050565b6060611d39826133cf565b611d42836134df565b611d4b846135e6565b604051602001611d5d93929190615785565b6040516020818303038152906040529050919050565b6068546000908190610100900463ffffffff16158015611d99575063ffffffff83166001145b15611dab575063ffffffff8316611dd3565b611dc064010000000063ffffffff85166157ed565b611dd09063ffffffff8616615804565b90505b600881901c600090815260696020526040902054600160ff9092169190911b908116149392505050565b60685460ff1615611e3a576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60685463ffffffff8681166101009092041614611e83576040517f0595ea2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611f878c8c8c8c8c611f8260008e8e8e8e8e8e8e604051611ea5929190615817565b60405180910390206040517fff0000000000000000000000000000000000000000000000000000000000000060f889901b1660208201527fffffffff0000000000000000000000000000000000000000000000000000000060e088811b821660218401527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606089811b821660258601529188901b909216603984015285901b16603d8201526051810183905260718101829052600090609101604051602081830303815290604052805190602001209050979650505050505050565b613a98565b73ffffffffffffffffffffffffffffffffffffffff86166120bf57606f5473ffffffffffffffffffffffffffffffffffffffff1661209657600073ffffffffffffffffffffffffffffffffffffffff851684825b6040519080825280601f01601f191660200182016040528015612005576020820181803683370190505b506040516120139190615827565b60006040518083038185875af1925050503d8060008114612050576040519150601f19603f3d011682016040523d82523d6000602084013e612055565b606091505b5050905080612090576040517f6747a28800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50612400565b606f546120ba9073ffffffffffffffffffffffffffffffffffffffff1685856139c0565b612400565b606d5473ffffffffffffffffffffffffffffffffffffffff878116911614801561210b5750606d5463ffffffff8881167401000000000000000000000000000000000000000090920416145b1561213057600073ffffffffffffffffffffffffffffffffffffffff85168482611fdb565b60685463ffffffff610100909104811690881603612169576120ba73ffffffffffffffffffffffffffffffffffffffff87168585613c6f565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e089901b1660208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606088901b166024820152600090603801604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152606a90935291205490915073ffffffffffffffffffffffffffffffffffffffff16806123f257600061226d8386868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506137a992505050565b905061227a8188886139c0565b80606a600085815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060405180604001604052808b63ffffffff1681526020018a73ffffffffffffffffffffffffffffffffffffffff16815250606b60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548163ffffffff021916908363ffffffff16021790555060208201518160000160046101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050507f490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a398a8a8388886040516123e495949392919061588c565b60405180910390a1506123fd565b6123fd8187876139c0565b50505b604080518b815263ffffffff8916602082015273ffffffffffffffffffffffffffffffffffffffff88811682840152861660608201526080810185905290517f1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d9181900360a00190a1505050505050505050505050565b60685460ff16156124b4576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6124bc613d43565b60685463ffffffff610100909104811690881603612506576040517f0595ea2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008060608773ffffffffffffffffffffffffffffffffffffffff881661263057883414612560576040517fb89240f500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606d54606e805473ffffffffffffffffffffffffffffffffffffffff831696507401000000000000000000000000000000000000000090920463ffffffff169450906125ab9061553c565b80601f01602080910402602001604051908101604052809291908181526020018280546125d79061553c565b80156126245780601f106125f957610100808354040283529160200191612624565b820191906000526020600020905b81548152906001019060200180831161260757829003601f168201915b50505050509150612891565b3415612668576040517f798ee6f100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5473ffffffffffffffffffffffffffffffffffffffff9081169089160361269a57612695888a6138ea565b612891565b73ffffffffffffffffffffffffffffffffffffffff8089166000908152606b602090815260409182902082518084019093525463ffffffff8116835264010000000090049092169181018290529015612708576126f7898b6138ea565b602081015181519095509350612884565b851561271a5761271a898b8989613db6565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260009073ffffffffffffffffffffffffffffffffffffffff8b16906370a0823190602401602060405180830381865afa158015612787573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127ab91906158d2565b90506127cf73ffffffffffffffffffffffffffffffffffffffff8b1633308e6142b8565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260009073ffffffffffffffffffffffffffffffffffffffff8c16906370a0823190602401602060405180830381865afa15801561283c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061286091906158d2565b905061286c82826158eb565b6068548c9850610100900463ffffffff169650935050505b61288d89611d2e565b9250505b7f501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b600084868e8e86886053546040516128d19897969594939291906158fe565b60405180910390a16129c86129c3600085878f8f8789805190602001206040517fff0000000000000000000000000000000000000000000000000000000000000060f889901b1660208201527fffffffff0000000000000000000000000000000000000000000000000000000060e088811b821660218401527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606089811b821660258601529188901b909216603984015285901b16603d8201526051810183905260718101829052600090609101604051602081830303815290604052805190602001209050979650505050505050565b614316565b86156129d6576129d66136d9565b505050506129e360018055565b50505050505050565b60a35473ffffffffffffffffffffffffffffffffffffffff163314612a3d576040517faf6e71a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116612a8a576040517f376f0fe600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60a380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f32cf74f8a6d5f88593984d2cd52be5592bfa6884f5896175801a5069ef09cd679060200160405180910390a150565b60685460ff1615612b40576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60685463ffffffff8681166101009092041614612b89576040517f0595ea2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612bab8c8c8c8c8c611f8260018e8e8e8e8e8e8e604051611ea5929190615817565b606f5460009073ffffffffffffffffffffffffffffffffffffffff16612cc8578473ffffffffffffffffffffffffffffffffffffffff1684888a8686604051602401612bfa9493929190615975565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f1806b5f20000000000000000000000000000000000000000000000000000000017905251612c7b9190615827565b60006040518083038185875af1925050503d8060008114612cb8576040519150601f19603f3d011682016040523d82523d6000602084013e612cbd565b606091505b505080915050612de3565b606f54612cec9073ffffffffffffffffffffffffffffffffffffffff1686866139c0565b8473ffffffffffffffffffffffffffffffffffffffff1687898585604051602401612d1a9493929190615975565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f1806b5f20000000000000000000000000000000000000000000000000000000017905251612d9b9190615827565b6000604051808303816000865af19150503d8060008114612dd8576040519150601f19603f3d011682016040523d82523d6000602084013e612ddd565b606091505b50909150505b80612e1a576040517f37e391c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080518c815263ffffffff8a16602082015273ffffffffffffffffffffffffffffffffffffffff89811682840152871660608201526080810186905290517f1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d9181900360a00190a150505050505050505050505050565b600054610100900460ff1615808015612eb25750600054600160ff909116105b80612ecc5750303b158015612ecc575060005460ff166001145b612f58576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401611004565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015612fb657600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6040517ff57ac68300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600081612ff6868686610e6c565b1495945050505050565b60685463ffffffff61010090910481169087160361304a576040517f0595ea2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b6001606860019054906101000a900463ffffffff1633898989888860535460405161309e999897969594939291906159bb565b60405180910390a16130d56129c36001606860019054906101000a900463ffffffff16338a8a8a8989604051611ea5929190615817565b8215611bcc57611bcc6136d9565b73ffffffffffffffffffffffffffffffffffffffff8316158061311a575073ffffffffffffffffffffffffffffffffffffffff8216155b15613151576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60685463ffffffff61010090910481169085160361319b576040517f658b23ad00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8281166000908152606b602052604090205464010000000090041615613202576040517f5eaf7bac00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b1660208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606085901b166024820152600090603801604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001815282825280516020918201206000818152606a835283812080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8a8116918217909255868601865263ffffffff8c81168089528c8416878a01818152848752606b89528987209a518b54915194167fffffffffffffffff0000000000000000000000000000000000000000000000009091161764010000000093909516929092029390931790975560a285529185902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001689151590811790915585519182529381019590955292840192909252606083015291507fdbe8a5da6a7a916d9adfda9160167a0f8a3da415ee6610e810e753853597fce79060800160405180910390a15050505050565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f06fdde03000000000000000000000000000000000000000000000000000000001790529051606091600091829173ffffffffffffffffffffffffffffffffffffffff8616916134519190615827565b600060405180830381855afa9150503d806000811461348c576040519150601f19603f3d011682016040523d82523d6000602084013e613491565b606091505b5091509150816134d6576040518060400160405280600781526020017f4e4f5f4e414d4500000000000000000000000000000000000000000000000000815250610e0c565b610e0c8161441a565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f95d89b41000000000000000000000000000000000000000000000000000000001790529051606091600091829173ffffffffffffffffffffffffffffffffffffffff8616916135619190615827565b600060405180830381855afa9150503d806000811461359c576040519150601f19603f3d011682016040523d82523d6000602084013e6135a1565b606091505b5091509150816134d6576040518060400160405280600981526020017f4e4f5f53594d424f4c0000000000000000000000000000000000000000000000815250610e0c565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f313ce5670000000000000000000000000000000000000000000000000000000017905290516000918291829173ffffffffffffffffffffffffffffffffffffffff8616916136679190615827565b600060405180830381855afa9150503d80600081146136a2576040519150601f19603f3d011682016040523d82523d6000602084013e6136a7565b606091505b50915091508180156136ba575080516020145b6136c5576012610e0c565b80806020019051810190610e0c9190615a34565b6053546068805463ffffffff909216790100000000000000000000000000000000000000000000000000027fffffff00000000ffffffffffffffffffffffffffffffffffffffffffffffffff909216919091179081905573ffffffffffffffffffffffffffffffffffffffff65010000000000909104166333d6247d61375d610d0d565b6040518263ffffffff1660e01b815260040161377b91815260200190565b600060405180830381600087803b15801561379557600080fd5b505af115801561188e573d6000803e3d6000fd5b60008060405180611ba00160405280611b668152602001615da7611b669139836040516020016137da9291906157be565b6040516020818303038152906040529050838151602083016000f5915073ffffffffffffffffffffffffffffffffffffffff8216613844576040517fbefb092000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5092915050565b600054610100900460ff166138e2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401611004565b610e4b6145ef565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260a2602052604090205460ff161561393e5761393a73ffffffffffffffffffffffffffffffffffffffff83163330846142b8565b5050565b6040517f9dc29fac0000000000000000000000000000000000000000000000000000000081523360048201526024810182905273ffffffffffffffffffffffffffffffffffffffff831690639dc29fac90604401600060405180830381600087803b1580156139ac57600080fd5b505af1158015611bcc573d6000803e3d6000fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260a2602052604090205460ff1615613a1457613a0f73ffffffffffffffffffffffffffffffffffffffff84168383613c6f565b505050565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8381166004830152602482018390528416906340c10f1990604401600060405180830381600087803b158015613a8457600080fd5b505af11580156129e3573d6000803e3d6000fd5b606854604080516020808201879052818301869052825180830384018152606083019384905280519101207f257b363200000000000000000000000000000000000000000000000000000000909252606481019190915260009165010000000000900473ffffffffffffffffffffffffffffffffffffffff169063257b3632906084016020604051808303816000875af1158015613b3a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b5e91906158d2565b905080600003613b99576040517e2f6fad00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008068010000000000000000871615613bf757869150613bbc848a8489612fe8565b613bf2576040517fe0417cec00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613c5a565b602087901c613c07816001615a51565b9150879250613c22613c1a868c86610e6c565b8a8389612fe8565b613c58576040517fe0417cec00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b613c648282614686565b505050505050505050565b60405173ffffffffffffffffffffffffffffffffffffffff8316602482015260448101829052613a0f9084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152614748565b600260015403613daf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401611004565b6002600155565b6000613dc56004828486615a6e565b613dce91615a98565b90507f2afa5331000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000821601614040576000808080808080613e2e896004818d615a6e565b810190613e3b9190615ae0565b96509650965096509650965096503373ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff1614613eae576040517f912ecce700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff86163014613efd576040517f750643af00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8a8514613f36576040517f03fffc4b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805173ffffffffffffffffffffffffffffffffffffffff89811660248301528881166044830152606482018890526084820187905260ff861660a483015260c4820185905260e48083018590528351808403909101815261010490920183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd505accf000000000000000000000000000000000000000000000000000000001790529151918e1691613fef9190615827565b6000604051808303816000865af19150503d806000811461402c576040519150601f19603f3d011682016040523d82523d6000602084013e614031565b606091505b50505050505050505050610b42565b7fffffffff0000000000000000000000000000000000000000000000000000000081167f8fcbaf0c00000000000000000000000000000000000000000000000000000000146140bb576040517fe282c0ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000808080808080806140d18a6004818e615a6e565b8101906140de9190615b34565b975097509750975097509750975097503373ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff1614614153576040517f912ecce700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff871630146141a2576040517f750643af00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805173ffffffffffffffffffffffffffffffffffffffff8a811660248301528981166044830152606482018990526084820188905286151560a483015260ff861660c483015260e482018590526101048083018590528351808403909101815261012490920183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f8fcbaf0c000000000000000000000000000000000000000000000000000000001790529151918f16916142649190615827565b6000604051808303816000865af19150503d80600081146142a1576040519150601f19603f3d011682016040523d82523d6000602084013e6142a6565b606091505b50505050505050505050505050505050565b60405173ffffffffffffffffffffffffffffffffffffffff8085166024830152831660448201526064810182905261188e9085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401613cc1565b80600161432560206002615cd7565b61432f91906158eb565b60535410614369576040517fef5ccf6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060536000815461437a906155ed565b9182905550905060005b602081101561440b578082901c6001166001036143b75782603382602081106143af576143af61558f565b015550505050565b603381602081106143ca576143ca61558f565b015460408051602081019290925281018490526060016040516020818303038152906040528051906020012092508080614403906155ed565b915050614384565b50613a0f615ce3565b60018055565b606060408251106144395781806020019051810190610a949190615d12565b81516020036145b15760005b60208110801561448c57508281815181106144625761446261558f565b01602001517fff000000000000000000000000000000000000000000000000000000000000001615155b156144a3578061449b816155ed565b915050614445565b806000036144e657505060408051808201909152601281527f4e4f545f56414c49445f454e434f44494e4700000000000000000000000000006020820152919050565b60008167ffffffffffffffff81111561450157614501614c39565b6040519080825280601f01601f19166020018201604052801561452b576020820181803683370190505b50905060005b828110156145a95784818151811061454b5761454b61558f565b602001015160f81c60f81b8282815181106145685761456861558f565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350806145a1816155ed565b915050614531565b509392505050565b505060408051808201909152601281527f4e4f545f56414c49445f454e434f44494e470000000000000000000000000000602082015290565b919050565b600054610100900460ff16614414576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401611004565b606854600090610100900463ffffffff161580156146aa575063ffffffff82166001145b156146bc575063ffffffff82166146e4565b6146d164010000000063ffffffff84166157ed565b6146e19063ffffffff8516615804565b90505b600881901c60008181526069602052604081208054600160ff861690811b918218928390559290919081831690036129e3576040517f646cf55800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006147aa826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166148549092919063ffffffff16565b805190915015613a0f57808060200190518101906147c89190615d89565b613a0f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401611004565b6060610e0c8484600085856000808673ffffffffffffffffffffffffffffffffffffffff1685876040516148889190615827565b60006040518083038185875af1925050503d80600081146148c5576040519150601f19603f3d011682016040523d82523d6000602084013e6148ca565b606091505b50915091506148db878383876148e6565b979650505050505050565b6060831561497c5782516000036149755773ffffffffffffffffffffffffffffffffffffffff85163b614975576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401611004565b5081610e0c565b610e0c83838151156149915781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110049190614b74565b803563ffffffff811681146145ea57600080fd5b73ffffffffffffffffffffffffffffffffffffffff811681146149fb57600080fd5b50565b60008060408385031215614a1157600080fd5b614a1a836149c5565b91506020830135614a2a816149d9565b809150509250929050565b80151581146149fb57600080fd5b60008083601f840112614a5557600080fd5b50813567ffffffffffffffff811115614a6d57600080fd5b602083019150836020828501011115614a8557600080fd5b9250929050565b600080600080600060808688031215614aa457600080fd5b614aad866149c5565b94506020860135614abd816149d9565b93506040860135614acd81614a35565b9250606086013567ffffffffffffffff811115614ae957600080fd5b614af588828901614a43565b969995985093965092949392505050565b60005b83811015614b21578181015183820152602001614b09565b50506000910152565b60008151808452614b42816020860160208601614b06565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000614b876020830184614b2a565b9392505050565b600060208284031215614ba057600080fd5b8135614b87816149d9565b60ff811681146149fb57600080fd5b600080600080600080600060e0888a031215614bd557600080fd5b8735614be081614bab565b9650614bee602089016149c5565b95506040880135614bfe816149d9565b9450614c0c606089016149c5565b93506080880135614c1c816149d9565b9699959850939692959460a0840135945060c09093013592915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715614caf57614caf614c39565b604052919050565b600067ffffffffffffffff821115614cd157614cd1614c39565b5060051b60200190565b600082601f830112614cec57600080fd5b81356020614d01614cfc83614cb7565b614c68565b82815260059290921b84018101918181019086841115614d2057600080fd5b8286015b84811015614d44578035614d37816149d9565b8352918301918301614d24565b509695505050505050565b600082601f830112614d6057600080fd5b81356020614d70614cfc83614cb7565b82815260059290921b84018101918181019086841115614d8f57600080fd5b8286015b84811015614d44578035614da681614a35565b8352918301918301614d93565b60008060008060808587031215614dc957600080fd5b843567ffffffffffffffff80821115614de157600080fd5b818701915087601f830112614df557600080fd5b81356020614e05614cfc83614cb7565b82815260059290921b8401810191818101908b841115614e2457600080fd5b948201945b83861015614e4957614e3a866149c5565b82529482019490820190614e29565b98505088013592505080821115614e5f57600080fd5b614e6b88838901614cdb565b94506040870135915080821115614e8157600080fd5b614e8d88838901614cdb565b93506060870135915080821115614ea357600080fd5b50614eb087828801614d4f565b91505092959194509250565b600080600060608486031215614ed157600080fd5b614eda846149c5565b92506020840135614eea816149d9565b91506040840135614efa816149d9565b809150509250925092565b600060208284031215614f1757600080fd5b5035919050565b806104008101831015610a9457600080fd5b60008060006104408486031215614f4657600080fd5b83359250614f578560208601614f1e565b9150614f6661042085016149c5565b90509250925092565b600067ffffffffffffffff821115614f8957614f89614c39565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f830112614fc657600080fd5b8135614fd4614cfc82614f6f565b818152846020838601011115614fe957600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060008060008060006101208a8c03121561502557600080fd5b61502e8a6149c5565b985060208a013561503e816149d9565b975061504c60408b016149c5565b965060608a013561505c816149d9565b955060808a013561506c816149d9565b945060a08a013567ffffffffffffffff81111561508857600080fd5b6150948c828d01614fb5565b94505060c08a01356150a5816149d9565b925060e08a01356150b5816149d9565b91506101008a01356150c681614a35565b809150509295985092959850929598565b600080604083850312156150ea57600080fd5b82356150f5816149d9565b946020939093013593505050565b600080600080600060a0868803121561511b57600080fd5b615124866149c5565b94506020860135615134816149d9565b9350604086013567ffffffffffffffff8082111561515157600080fd5b61515d89838a01614fb5565b9450606088013591508082111561517357600080fd5b5061518088828901614fb5565b925050608086013561519181614bab565b809150509295509295909350565b600080600080608085870312156151b557600080fd5b6151be856149c5565b935060208501356151ce816149d9565b925060408501356151de816149d9565b915060608501356151ee81614a35565b939692955090935050565b60008060008060008060a0878903121561521257600080fd5b61521b876149c5565b9550602087013561522b816149d9565b945060408701359350606087013561524281614a35565b9250608087013567ffffffffffffffff81111561525e57600080fd5b61526a89828a01614a43565b979a9699509497509295939492505050565b6000806040838503121561528f57600080fd5b823561529a816149d9565b91506020830135614a2a81614a35565b600080604083850312156152bd57600080fd5b6152c6836149c5565b91506152d4602084016149c5565b90509250929050565b6000806000806000806000806000806000806109208d8f03121561530057600080fd5b61530a8e8e614f1e565b9b5061531a8e6104008f01614f1e565b9a506108008d013599506108208d013598506108408d013597506153416108608e016149c5565b96506153516108808e01356149d9565b6108808d013595506153666108a08e016149c5565b94506153766108c08e01356149d9565b6108c08d013593506108e08d0135925067ffffffffffffffff6109008e013511156153a057600080fd5b6153b18e6109008f01358f01614a43565b81935080925050509295989b509295989b509295989b565b600080600080600080600060c0888a0312156153e457600080fd5b6153ed886149c5565b965060208801356153fd816149d9565b9550604088013594506060880135615414816149d9565b9350608088013561542481614a35565b925060a088013567ffffffffffffffff81111561544057600080fd5b61544c8a828b01614a43565b989b979a50959850939692959293505050565b60008060008060008060c0878903121561547857600080fd5b615481876149c5565b95506020870135615491816149d9565b945061549f604088016149c5565b935060608701356154af816149d9565b925060808701356154bf816149d9565b915060a087013567ffffffffffffffff8111156154db57600080fd5b6154e789828a01614fb5565b9150509295509295509295565b600080600080610460858703121561550b57600080fd5b8435935061551c8660208701614f1e565b925061552b61042086016149c5565b939692955092936104400135925050565b600181811c9082168061555057607f821691505b602082108103615589577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361561e5761561e6155be565b5060010190565b601f821115613a0f57600081815260208120601f850160051c8101602086101561564c5750805b601f850160051c820191505b81811015611bcc57828155600101615658565b815167ffffffffffffffff81111561568557615685614c39565b61569981615693845461553c565b84615625565b602080601f8311600181146156ec57600084156156b65750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555611bcc565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b828110156157395788860151825594840194600190910190840161571a565b508582101561577557878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b6060815260006157986060830186614b2a565b82810360208401526157aa8186614b2a565b91505060ff83166040830152949350505050565b600083516157d0818460208801614b06565b8351908301906157e4818360208801614b06565b01949350505050565b8082028115828204841417610a9457610a946155be565b80820180821115610a9457610a946155be565b8183823760009101908152919050565b60008251615839818460208701614b06565b9190910192915050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b63ffffffff86168152600073ffffffffffffffffffffffffffffffffffffffff8087166020840152808616604084015250608060608301526148db608083018486615843565b6000602082840312156158e457600080fd5b5051919050565b81810381811115610a9457610a946155be565b600061010060ff8b16835263ffffffff808b16602085015273ffffffffffffffffffffffffffffffffffffffff808b166040860152818a1660608601528089166080860152508660a08501528160c085015261595c82850187614b2a565b925080851660e085015250509998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8516815263ffffffff841660208201526060604082015260006159b1606083018486615843565b9695505050505050565b600061010060ff8c16835263ffffffff808c16602085015273ffffffffffffffffffffffffffffffffffffffff808c166040860152818b166060860152808a166080860152508760a08501528160c0850152615a1a8285018789615843565b925080851660e085015250509a9950505050505050505050565b600060208284031215615a4657600080fd5b8151614b8781614bab565b63ffffffff818116838216019080821115613844576138446155be565b60008085851115615a7e57600080fd5b83861115615a8b57600080fd5b5050820193919092039150565b7fffffffff000000000000000000000000000000000000000000000000000000008135818116916004851015615ad85780818660040360031b1b83161692505b505092915050565b600080600080600080600060e0888a031215615afb57600080fd5b8735615b06816149d9565b96506020880135615b16816149d9565b955060408801359450606088013593506080880135614c1c81614bab565b600080600080600080600080610100898b031215615b5157600080fd5b8835615b5c816149d9565b97506020890135615b6c816149d9565b965060408901359550606089013594506080890135615b8a81614a35565b935060a0890135615b9a81614bab565b979a969950949793969295929450505060c08201359160e0013590565b600181815b80851115615c1057817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115615bf657615bf66155be565b80851615615c0357918102915b93841c9390800290615bbc565b509250929050565b600082615c2757506001610a94565b81615c3457506000610a94565b8160018114615c4a5760028114615c5457615c70565b6001915050610a94565b60ff841115615c6557615c656155be565b50506001821b610a94565b5060208310610133831016604e8410600b8410161715615c93575081810a610a94565b615c9d8383615bb7565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115615ccf57615ccf6155be565b029392505050565b6000614b878383615c18565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600060208284031215615d2457600080fd5b815167ffffffffffffffff811115615d3b57600080fd5b8201601f81018413615d4c57600080fd5b8051615d5a614cfc82614f6f565b818152856020838501011115615d6f57600080fd5b615d80826020830160208601614b06565b95945050505050565b600060208284031215615d9b57600080fd5b8151614b8781614a3556fe6101006040523480156200001257600080fd5b5060405162001b6638038062001b6683398101604081905262000035916200028d565b82826003620000458382620003a1565b506004620000548282620003a1565b50503360c0525060ff811660e052466080819052620000739062000080565b60a052506200046d915050565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f620000ad6200012e565b805160209182012060408051808201825260018152603160f81b90840152805192830193909352918101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc66060820152608081018390523060a082015260c001604051602081830303815290604052805190602001209050919050565b6060600380546200013f9062000312565b80601f01602080910402602001604051908101604052809291908181526020018280546200016d9062000312565b8015620001be5780601f106200019257610100808354040283529160200191620001be565b820191906000526020600020905b815481529060010190602001808311620001a057829003601f168201915b5050505050905090565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620001f057600080fd5b81516001600160401b03808211156200020d576200020d620001c8565b604051601f8301601f19908116603f01168101908282118183101715620002385762000238620001c8565b816040528381526020925086838588010111156200025557600080fd5b600091505b838210156200027957858201830151818301840152908201906200025a565b600093810190920192909252949350505050565b600080600060608486031215620002a357600080fd5b83516001600160401b0380821115620002bb57600080fd5b620002c987838801620001de565b94506020860151915080821115620002e057600080fd5b50620002ef86828701620001de565b925050604084015160ff811681146200030757600080fd5b809150509250925092565b600181811c908216806200032757607f821691505b6020821081036200034857634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200039c57600081815260208120601f850160051c81016020861015620003775750805b601f850160051c820191505b81811015620003985782815560010162000383565b5050505b505050565b81516001600160401b03811115620003bd57620003bd620001c8565b620003d581620003ce845462000312565b846200034e565b602080601f8311600181146200040d5760008415620003f45750858301515b600019600386901b1c1916600185901b17855562000398565b600085815260208120601f198616915b828110156200043e578886015182559484019460019091019084016200041d565b50858210156200045d5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a05160c05160e0516116aa620004bc6000396000610237015260008181610307015281816105c001526106a70152600061053a015260008181610379015261050401526116aa6000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c806370a08231116100d8578063a457c2d71161008c578063d505accf11610066578063d505accf1461039b578063dd62ed3e146103ae578063ffa1ad74146103f457600080fd5b8063a457c2d71461034e578063a9059cbb14610361578063cd0d00961461037457600080fd5b806395d89b41116100bd57806395d89b41146102e75780639dc29fac146102ef578063a3c573eb1461030257600080fd5b806370a08231146102915780637ecebe00146102c757600080fd5b806330adf81f1161012f5780633644e515116101145780633644e51514610261578063395093511461026957806340c10f191461027c57600080fd5b806330adf81f14610209578063313ce5671461023057600080fd5b806318160ddd1161016057806318160ddd146101bd57806320606b70146101cf57806323b872dd146101f657600080fd5b806306fdde031461017c578063095ea7b31461019a575b600080fd5b610184610430565b60405161019191906113e4565b60405180910390f35b6101ad6101a8366004611479565b6104c2565b6040519015158152602001610191565b6002545b604051908152602001610191565b6101c17f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f81565b6101ad6102043660046114a3565b6104dc565b6101c17f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c981565b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610191565b6101c1610500565b6101ad610277366004611479565b61055c565b61028f61028a366004611479565b6105a8565b005b6101c161029f3660046114df565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6101c16102d53660046114df565b60056020526000908152604090205481565b610184610680565b61028f6102fd366004611479565b61068f565b6103297f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610191565b6101ad61035c366004611479565b61075e565b6101ad61036f366004611479565b61082f565b6101c17f000000000000000000000000000000000000000000000000000000000000000081565b61028f6103a9366004611501565b61083d565b6101c16103bc366004611574565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6101846040518060400160405280600181526020017f310000000000000000000000000000000000000000000000000000000000000081525081565b60606003805461043f906115a7565b80601f016020809104026020016040519081016040528092919081815260200182805461046b906115a7565b80156104b85780601f1061048d576101008083540402835291602001916104b8565b820191906000526020600020905b81548152906001019060200180831161049b57829003601f168201915b5050505050905090565b6000336104d0818585610b73565b60019150505b92915050565b6000336104ea858285610d27565b6104f5858585610dfe565b506001949350505050565b60007f00000000000000000000000000000000000000000000000000000000000000004614610537576105324661106d565b905090565b507f000000000000000000000000000000000000000000000000000000000000000090565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091906104d090829086906105a3908790611629565b610b73565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610672576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f546f6b656e577261707065643a3a6f6e6c794272696467653a204e6f7420506f60448201527f6c79676f6e5a6b45564d4272696467650000000000000000000000000000000060648201526084015b60405180910390fd5b61067c8282611135565b5050565b60606004805461043f906115a7565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610754576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f546f6b656e577261707065643a3a6f6e6c794272696467653a204e6f7420506f60448201527f6c79676f6e5a6b45564d427269646765000000000000000000000000000000006064820152608401610669565b61067c8282611228565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919083811015610822576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f0000000000000000000000000000000000000000000000000000006064820152608401610669565b6104f58286868403610b73565b6000336104d0818585610dfe565b834211156108cc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f546f6b656e577261707065643a3a7065726d69743a204578706972656420706560448201527f726d6974000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff8716600090815260056020526040812080547f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9918a918a918a9190866109268361163c565b9091555060408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090506000610991610500565b6040517f19010000000000000000000000000000000000000000000000000000000000006020820152602281019190915260428101839052606201604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600080855291840180845281905260ff89169284019290925260608301879052608083018690529092509060019060a0016020604051602081039080840390855afa158015610a55573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811615801590610ad057508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b610b5c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f546f6b656e577261707065643a3a7065726d69743a20496e76616c696420736960448201527f676e6174757265000000000000000000000000000000000000000000000000006064820152608401610669565b610b678a8a8a610b73565b50505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8316610c15576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff8216610cb8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610df85781811015610deb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610669565b610df88484848403610b73565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610ea1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff8216610f44576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610ffa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3610df8565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f611098610430565b8051602091820120604080518082018252600181527f310000000000000000000000000000000000000000000000000000000000000090840152805192830193909352918101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc66060820152608081018390523060a082015260c001604051602081830303815290604052805190602001209050919050565b73ffffffffffffffffffffffffffffffffffffffff82166111b2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610669565b80600260008282546111c49190611629565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b73ffffffffffffffffffffffffffffffffffffffff82166112cb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f73000000000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090205481811015611381576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f63650000000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff83166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101610d1a565b600060208083528351808285015260005b81811015611411578581018301518582016040015282016113f5565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461147457600080fd5b919050565b6000806040838503121561148c57600080fd5b61149583611450565b946020939093013593505050565b6000806000606084860312156114b857600080fd5b6114c184611450565b92506114cf60208501611450565b9150604084013590509250925092565b6000602082840312156114f157600080fd5b6114fa82611450565b9392505050565b600080600080600080600060e0888a03121561151c57600080fd5b61152588611450565b965061153360208901611450565b95506040880135945060608801359350608088013560ff8116811461155757600080fd5b9699959850939692959460a0840135945060c09093013592915050565b6000806040838503121561158757600080fd5b61159083611450565b915061159e60208401611450565b90509250929050565b600181811c908216806115bb57607f821691505b6020821081036115f4577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156104d6576104d66115fa565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361166d5761166d6115fa565b506001019056fea26469706673582212208d88fee561cff7120d381c345cfc534cef8229a272dc5809d4bbb685ad67141164736f6c63430008110033a2646970667358221220bedaed4c9e67bd0f1f96d4141578969eae27bac5903385f3fd4861a1780d92a964736f6c63430008140033",
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

// BASEINITBYTECODEWRAPPEDTOKEN is a free data retrieval call binding the contract method 0x83c43a55.
//
// Solidity: function BASE_INIT_BYTECODE_WRAPPED_TOKEN() view returns(bytes)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) BASEINITBYTECODEWRAPPEDTOKEN(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "BASE_INIT_BYTECODE_WRAPPED_TOKEN")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// BASEINITBYTECODEWRAPPEDTOKEN is a free data retrieval call binding the contract method 0x83c43a55.
//
// Solidity: function BASE_INIT_BYTECODE_WRAPPED_TOKEN() view returns(bytes)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) BASEINITBYTECODEWRAPPEDTOKEN() ([]byte, error) {
	return _Bridgel2sovereignchain.Contract.BASEINITBYTECODEWRAPPEDTOKEN(&_Bridgel2sovereignchain.CallOpts)
}

// BASEINITBYTECODEWRAPPEDTOKEN is a free data retrieval call binding the contract method 0x83c43a55.
//
// Solidity: function BASE_INIT_BYTECODE_WRAPPED_TOKEN() view returns(bytes)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) BASEINITBYTECODEWRAPPEDTOKEN() ([]byte, error) {
	return _Bridgel2sovereignchain.Contract.BASEINITBYTECODEWRAPPEDTOKEN(&_Bridgel2sovereignchain.CallOpts)
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

// ActivateEmergencyState is a free data retrieval call binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() pure returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) ActivateEmergencyState(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "activateEmergencyState")

	if err != nil {
		return err
	}

	return err

}

// ActivateEmergencyState is a free data retrieval call binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() pure returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) ActivateEmergencyState() error {
	return _Bridgel2sovereignchain.Contract.ActivateEmergencyState(&_Bridgel2sovereignchain.CallOpts)
}

// ActivateEmergencyState is a free data retrieval call binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() pure returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) ActivateEmergencyState() error {
	return _Bridgel2sovereignchain.Contract.ActivateEmergencyState(&_Bridgel2sovereignchain.CallOpts)
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

// CalculateTokenWrapperAddress is a free data retrieval call binding the contract method 0x7843298b.
//
// Solidity: function calculateTokenWrapperAddress(uint32 originNetwork, address originTokenAddress, address token) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) CalculateTokenWrapperAddress(opts *bind.CallOpts, originNetwork uint32, originTokenAddress common.Address, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "calculateTokenWrapperAddress", originNetwork, originTokenAddress, token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CalculateTokenWrapperAddress is a free data retrieval call binding the contract method 0x7843298b.
//
// Solidity: function calculateTokenWrapperAddress(uint32 originNetwork, address originTokenAddress, address token) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) CalculateTokenWrapperAddress(originNetwork uint32, originTokenAddress common.Address, token common.Address) (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.CalculateTokenWrapperAddress(&_Bridgel2sovereignchain.CallOpts, originNetwork, originTokenAddress, token)
}

// CalculateTokenWrapperAddress is a free data retrieval call binding the contract method 0x7843298b.
//
// Solidity: function calculateTokenWrapperAddress(uint32 originNetwork, address originTokenAddress, address token) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) CalculateTokenWrapperAddress(originNetwork uint32, originTokenAddress common.Address, token common.Address) (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.CalculateTokenWrapperAddress(&_Bridgel2sovereignchain.CallOpts, originNetwork, originTokenAddress, token)
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

// DeactivateEmergencyState is a free data retrieval call binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() pure returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) DeactivateEmergencyState(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "deactivateEmergencyState")

	if err != nil {
		return err
	}

	return err

}

// DeactivateEmergencyState is a free data retrieval call binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() pure returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) DeactivateEmergencyState() error {
	return _Bridgel2sovereignchain.Contract.DeactivateEmergencyState(&_Bridgel2sovereignchain.CallOpts)
}

// DeactivateEmergencyState is a free data retrieval call binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() pure returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) DeactivateEmergencyState() error {
	return _Bridgel2sovereignchain.Contract.DeactivateEmergencyState(&_Bridgel2sovereignchain.CallOpts)
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

// PrecalculatedWrapperAddress is a free data retrieval call binding the contract method 0xaaa13cc2.
//
// Solidity: function precalculatedWrapperAddress(uint32 originNetwork, address originTokenAddress, string name, string symbol, uint8 decimals) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) PrecalculatedWrapperAddress(opts *bind.CallOpts, originNetwork uint32, originTokenAddress common.Address, name string, symbol string, decimals uint8) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "precalculatedWrapperAddress", originNetwork, originTokenAddress, name, symbol, decimals)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrecalculatedWrapperAddress is a free data retrieval call binding the contract method 0xaaa13cc2.
//
// Solidity: function precalculatedWrapperAddress(uint32 originNetwork, address originTokenAddress, string name, string symbol, uint8 decimals) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) PrecalculatedWrapperAddress(originNetwork uint32, originTokenAddress common.Address, name string, symbol string, decimals uint8) (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.PrecalculatedWrapperAddress(&_Bridgel2sovereignchain.CallOpts, originNetwork, originTokenAddress, name, symbol, decimals)
}

// PrecalculatedWrapperAddress is a free data retrieval call binding the contract method 0xaaa13cc2.
//
// Solidity: function precalculatedWrapperAddress(uint32 originNetwork, address originTokenAddress, string name, string symbol, uint8 decimals) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) PrecalculatedWrapperAddress(originNetwork uint32, originTokenAddress common.Address, name string, symbol string, decimals uint8) (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.PrecalculatedWrapperAddress(&_Bridgel2sovereignchain.CallOpts, originNetwork, originTokenAddress, name, symbol, decimals)
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

// Initialize is a paid mutator transaction binding the contract method 0x8c0dd470.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address _sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) Initialize(opts *bind.TransactOpts, _networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, _sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "initialize", _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, _sovereignWETHAddress, _sovereignWETHAddressIsNotMintable)
}

// Initialize is a paid mutator transaction binding the contract method 0x8c0dd470.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address _sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) Initialize(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, _sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.Initialize(&_Bridgel2sovereignchain.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, _sovereignWETHAddress, _sovereignWETHAddressIsNotMintable)
}

// Initialize is a paid mutator transaction binding the contract method 0x8c0dd470.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address _sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) Initialize(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, _sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.Initialize(&_Bridgel2sovereignchain.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, _sovereignWETHAddress, _sovereignWETHAddressIsNotMintable)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes ) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) Initialize0(opts *bind.TransactOpts, arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "initialize0", arg0, arg1, arg2, arg3, arg4, arg5)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes ) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) Initialize0(arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.Initialize0(&_Bridgel2sovereignchain.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes ) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) Initialize0(arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.Initialize0(&_Bridgel2sovereignchain.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// MigrateLegacyToken is a paid mutator transaction binding the contract method 0x9e76158f.
//
// Solidity: function migrateLegacyToken(address legacyTokenAddress, uint256 amount) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) MigrateLegacyToken(opts *bind.TransactOpts, legacyTokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "migrateLegacyToken", legacyTokenAddress, amount)
}

// MigrateLegacyToken is a paid mutator transaction binding the contract method 0x9e76158f.
//
// Solidity: function migrateLegacyToken(address legacyTokenAddress, uint256 amount) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) MigrateLegacyToken(legacyTokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.MigrateLegacyToken(&_Bridgel2sovereignchain.TransactOpts, legacyTokenAddress, amount)
}

// MigrateLegacyToken is a paid mutator transaction binding the contract method 0x9e76158f.
//
// Solidity: function migrateLegacyToken(address legacyTokenAddress, uint256 amount) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) MigrateLegacyToken(legacyTokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.MigrateLegacyToken(&_Bridgel2sovereignchain.TransactOpts, legacyTokenAddress, amount)
}

// RemoveLegacySovereignTokenAddress is a paid mutator transaction binding the contract method 0xb4586962.
//
// Solidity: function removeLegacySovereignTokenAddress(address sovereignTokenAddress) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) RemoveLegacySovereignTokenAddress(opts *bind.TransactOpts, sovereignTokenAddress common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "removeLegacySovereignTokenAddress", sovereignTokenAddress)
}

// RemoveLegacySovereignTokenAddress is a paid mutator transaction binding the contract method 0xb4586962.
//
// Solidity: function removeLegacySovereignTokenAddress(address sovereignTokenAddress) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) RemoveLegacySovereignTokenAddress(sovereignTokenAddress common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.RemoveLegacySovereignTokenAddress(&_Bridgel2sovereignchain.TransactOpts, sovereignTokenAddress)
}

// RemoveLegacySovereignTokenAddress is a paid mutator transaction binding the contract method 0xb4586962.
//
// Solidity: function removeLegacySovereignTokenAddress(address sovereignTokenAddress) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) RemoveLegacySovereignTokenAddress(sovereignTokenAddress common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.RemoveLegacySovereignTokenAddress(&_Bridgel2sovereignchain.TransactOpts, sovereignTokenAddress)
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

// SetSovereignTokenAddress is a paid mutator transaction binding the contract method 0xb42f6b3a.
//
// Solidity: function setSovereignTokenAddress(uint32 originNetwork, address originTokenAddress, address sovereignTokenAddress, bool isNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) SetSovereignTokenAddress(opts *bind.TransactOpts, originNetwork uint32, originTokenAddress common.Address, sovereignTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "setSovereignTokenAddress", originNetwork, originTokenAddress, sovereignTokenAddress, isNotMintable)
}

// SetSovereignTokenAddress is a paid mutator transaction binding the contract method 0xb42f6b3a.
//
// Solidity: function setSovereignTokenAddress(uint32 originNetwork, address originTokenAddress, address sovereignTokenAddress, bool isNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) SetSovereignTokenAddress(originNetwork uint32, originTokenAddress common.Address, sovereignTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.SetSovereignTokenAddress(&_Bridgel2sovereignchain.TransactOpts, originNetwork, originTokenAddress, sovereignTokenAddress, isNotMintable)
}

// SetSovereignTokenAddress is a paid mutator transaction binding the contract method 0xb42f6b3a.
//
// Solidity: function setSovereignTokenAddress(uint32 originNetwork, address originTokenAddress, address sovereignTokenAddress, bool isNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) SetSovereignTokenAddress(originNetwork uint32, originTokenAddress common.Address, sovereignTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.SetSovereignTokenAddress(&_Bridgel2sovereignchain.TransactOpts, originNetwork, originTokenAddress, sovereignTokenAddress, isNotMintable)
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
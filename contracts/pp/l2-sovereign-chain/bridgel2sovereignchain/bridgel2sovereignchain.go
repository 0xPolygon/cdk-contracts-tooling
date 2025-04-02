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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountDoesNotMatchMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DestinationNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmergencyStateNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EtherTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTokenWrappedDeployment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputArraysLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSmtProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSovereignWETHAddressParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MsgValueNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeTokenIsEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoValueInMessagesOnGasTokenNetworks\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyBridgeManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OriginNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyMapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyUpdated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotMapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotRemapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WETHRemappingNotSupportedOnGasTokenNetworks\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"depositCount\",\"type\":\"uint32\"}],\"name\":\"BridgeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"legacyTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"updatedTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MigrateLegacyToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wrappedTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"NewWrappedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignTokenAddress\",\"type\":\"address\"}],\"name\":\"RemoveLegacySovereignTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridgeManager\",\"type\":\"address\"}],\"name\":\"SetBridgeManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"SetSovereignTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignWETHTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"SetSovereignWETHAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"claimedGlobalIndex\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newClaimedGlobalIndexHashChain\",\"type\":\"bytes32\"}],\"name\":\"UpdatedClaimedGlobalIndexHashChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"unsetGlobalIndex\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newUnsetGlobalIndexHashChain\",\"type\":\"bytes32\"}],\"name\":\"UpdatedUnsetGlobalIndexHashChain\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASE_INIT_BYTECODE_WRAPPED_TOKEN\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETHToken\",\"outputs\":[{\"internalType\":\"contractTokenWrapped\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"permitData\",\"type\":\"bytes\"}],\"name\":\"bridgeAsset\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountWETH\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessageWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"calculateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"calculateTokenWrapperAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimedBitMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimedGlobalIndexHashChain\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"name\":\"getLeafValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenWrappedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_polygonRollupManager\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_bridgeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sovereignWETHAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_sovereignWETHAddressIsNotMintable\",\"type\":\"bool\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"leafIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"sourceBridgeNetwork\",\"type\":\"uint32\"}],\"name\":\"isClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdatedDepositCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"legacyTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"migrateLegacyToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"polygonRollupManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"precalculatedWrapperAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"legacySovereignTokenAddress\",\"type\":\"address\"}],\"name\":\"removeLegacySovereignTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeManager\",\"type\":\"address\"}],\"name\":\"setBridgeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"originNetworks\",\"type\":\"uint32[]\"},{\"internalType\":\"address[]\",\"name\":\"originTokenAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"sovereignTokenAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"isNotMintable\",\"type\":\"bool[]\"}],\"name\":\"setMultipleSovereignTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sovereignWETHTokenAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"setSovereignWETHAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenInfoToWrappedToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unsetGlobalIndexHashChain\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"globalIndexes\",\"type\":\"uint256[]\"}],\"name\":\"unsetMultipleClaims\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"verifyMerkleProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrappedAddress\",\"type\":\"address\"}],\"name\":\"wrappedAddressIsNotMintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wrappedTokenToTokenInfo\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506015601f565b601b601f565b60da565b5f54610100900460ff161560895760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116101560d8575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b615fd8806100e75f395ff3fe6080604052600436106101fa575f3560e01c806383f244031161011257806383f24403146104db5780638c0dd470146104fa5780638ed7e3f2146105195780639e76158f14610538578063aaa13cc214610557578063b458696214610576578063b8b284d014610595578063bab161bf146105b4578063be5831c7146105d5578063bf130d7f146105f8578063c00f14ab14610617578063c0f4916314610636578063cc46163214610664578063ccaa2d1114610683578063cd586579146106a2578063d02103ca146106b5578063dbc16976146102a0578063eabd372a146106db578063ee25560b146106fa578063f5efcd7914610725578063f811bff714610744578063fb57083414610763575f5ffd5b8063136a2c60146101fe57806314cc01a01461021f57806315064c96146102545780631d081d8c1461027d5780632072f6c5146102a057806322e95f2c146102b4578063240ff378146102d357806327aef4e8146102e65780632dfdf0b514610307578063318aee3d1461031c5780633c351e10146103835780633cbc795b146103a25780633e197043146103da5780634b2f336d146103f957806357cfbee3146104185780635ca1e165146104375780636ee84b231461044b5780637843298b1461046057806379e2cf971461047f57806381b1c1741461049357806383c43a55146104c7575b5f5ffd5b348015610209575f5ffd5b5061021d6102183660046131c1565b610782565b005b34801561022a575f5ffd5b5060a35461023e906001600160a01b031681565b60405161024b9190613256565b60405180910390f35b34801561025f575f5ffd5b5060685461026d9060ff1681565b604051901515815260200161024b565b348015610288575f5ffd5b5061029260a45481565b60405190815260200161024b565b3480156102ab575f5ffd5b5061021d610868565b3480156102bf575f5ffd5b5061023e6102ce366004613294565b610881565b61021d6102e136600461331a565b6108cf565b3480156102f1575f5ffd5b506102fa61093f565b60405161024b91906133db565b348015610312575f5ffd5b5061029260535481565b348015610327575f5ffd5b5061035f6103363660046133ed565b606b6020525f908152604090205463ffffffff811690600160201b90046001600160a01b031682565b6040805163ffffffff90931683526001600160a01b0390911660208301520161024b565b34801561038e575f5ffd5b50606d5461023e906001600160a01b031681565b3480156103ad575f5ffd5b50606d546103c590600160a01b900463ffffffff1681565b60405163ffffffff909116815260200161024b565b3480156103e5575f5ffd5b506102926103f4366004613416565b6109cb565b348015610404575f5ffd5b50606f5461023e906001600160a01b031681565b348015610423575f5ffd5b5061021d610432366004613562565b610a57565b348015610442575f5ffd5b50610292610b43565b348015610456575f5ffd5b5061029260a55481565b34801561046b575f5ffd5b5061023e61047a366004613675565b610bb3565b34801561048a575f5ffd5b5061021d610bdc565b34801561049e575f5ffd5b5061023e6104ad3660046136bb565b606a6020525f90815260409020546001600160a01b031681565b3480156104d2575f5ffd5b506102fa610bff565b3480156104e6575f5ffd5b506102926104f53660046136e3565b610c1e565b348015610505575f5ffd5b5061021d610514366004613794565b610c99565b348015610524575f5ffd5b50606c5461023e906001600160a01b031681565b348015610543575f5ffd5b5061021d61055236600461385d565b610f7a565b348015610562575f5ffd5b5061023e610571366004613887565b6110b2565b348015610581575f5ffd5b5061021d6105903660046133ed565b6111b2565b3480156105a0575f5ffd5b5061021d6105af366004613920565b611313565b3480156105bf575f5ffd5b506068546103c590610100900463ffffffff1681565b3480156105e0575f5ffd5b506068546103c590600160c81b900463ffffffff1681565b348015610603575f5ffd5b5061021d61061236600461399d565b61138c565b348015610622575f5ffd5b506102fa6106313660046133ed565b611456565b348015610641575f5ffd5b5061026d6106503660046133ed565b60a26020525f908152604090205460ff1681565b34801561066f575f5ffd5b5061026d61067e3660046139c9565b61149b565b34801561068e575f5ffd5b5061021d61069d3660046139fa565b6114eb565b61021d6106b0366004613ad8565b6118df565b3480156106c0575f5ffd5b5060685461023e90600160281b90046001600160a01b031681565b3480156106e6575f5ffd5b5061021d6106f53660046133ed565b611c7f565b348015610705575f5ffd5b506102926107143660046136bb565b60696020525f908152604090205481565b348015610730575f5ffd5b5061021d61073f3660046139fa565b611d01565b34801561074f575f5ffd5b5061021d61075e366004613b67565b611f34565b34801561076e575f5ffd5b5061026d61077d366004613bf6565b611fc8565b60a3546001600160a01b031633146107ad576040516357b738d160e11b815260040160405180910390fd5b5f5b8151811015610864575f8282815181106107cb576107cb613c3b565b602002602001015190505f5f600160401b83165f146107ec57829150610803565b602083901c6107fc816001613c63565b9150839250505b61080d8282611fdf565b60a55461081a9084612050565b60a58190556040805185815260208101929092527fc80e0aca446a59735359a7ae46124b57c47b892827642779bc6dafc84ba90b03910160405180910390a15050506001016107af565b5050565b60405163441845b160e01b815260040160405180910390fd5b5f606a5f8484604051602001610898929190613c7f565b60408051601f198184030181529181528151602092830120835290820192909252015f20546001600160a01b031690505b92915050565b60685460ff16156108f357604051630bc011ff60e21b815260040160405180910390fd5b341580159061090c5750606f546001600160a01b031615155b1561092a576040516301bd897160e61b815260040160405180910390fd5b61093885853486868661205e565b5050505050565b606e805461094c90613ca9565b80601f016020809104026020016040519081016040528092919081815260200182805461097890613ca9565b80156109c35780601f1061099a576101008083540402835291602001916109c3565b820191905f5260205f20905b8154815290600101906020018083116109a657829003601f168201915b505050505081565b6040516001600160f81b031960f889901b1660208201526001600160e01b031960e088811b821660218401526001600160601b0319606089811b821660258601529188901b909216603984015285901b16603d82015260518101839052607181018290525f90609101604051602081830303815290604052805190602001209050979650505050505050565b60a3546001600160a01b03163314610a82576040516357b738d160e11b815260040160405180910390fd5b82518451141580610a9557508151845114155b80610aa257508051845114155b15610ac05760405163434f49f560e11b815260040160405180910390fd5b5f5b825181101561093857610b3b858281518110610ae057610ae0613c3b565b6020026020010151858381518110610afa57610afa613c3b565b6020026020010151858481518110610b1457610b14613c3b565b6020026020010151858581518110610b2e57610b2e613c3b565b6020026020010151612115565b600101610ac2565b6053545f90819081805b6020811015610baa578083901c600116600103610b8957610b8260338260208110610b7a57610b7a613c3b565b015485612050565b9350610b96565b610b938483612050565b93505b610ba08283612050565b9150600101610b4d565b50919392505050565b5f610bd48484610bc2856122c5565b610bcb86612380565b61057187612434565b949350505050565b605354606854600160c81b900463ffffffff161015610bfd57610bfd6124e8565b565b60405180611ba00160405280611b6681526020016143fd611b66913981565b5f83815b6020811015610c9057600163ffffffff8516821c81169003610c6557610c5e858260208110610c5357610c53613c3b565b602002013583612050565b9150610c88565b610c8582868360208110610c7b57610c7b613c3b565b6020020135612050565b91505b600101610c22565b50949350505050565b5f54610100900460ff1615808015610cb757505f54600160ff909116105b80610cd75750610cc63061257c565b158015610cd757505f5460ff166001145b610cfc5760405162461bcd60e51b8152600401610cf390613ce1565b60405180910390fd5b5f805460ff191660011790558015610d1d575f805461ff0019166101001790555b60688054610100600160c81b03191661010063ffffffff8d1602600160281b600160c81b03191617600160281b6001600160a01b038a81169190910291909117909155606c80546001600160a01b03199081168984161790915560a380549091168683161790558916610de65763ffffffff881615610daf57604051630d43a60960e11b815260040160405180910390fd5b6001600160a01b038316151580610dc35750815b15610de157604051630e6e237560e11b815260040160405180910390fd5b610f21565b606d805463ffffffff8a16600160a01b026001600160c01b03199091166001600160a01b038c1617179055606e610e1d8682613d73565b506001600160a01b038316610ee957811515600103610e4f57604051630e6e237560e11b815260040160405180910390fd5b610ec45f5f1b6012604051602001610eb091906060808252600d908201526c2bb930b83832b21022ba3432b960991b608082015260a060208201819052600490820152630ae8aa8960e31b60c082015260ff91909116604082015260e00190565b60405160208183030381529060405261258b565b606f80546001600160a01b0319166001600160a01b0392909216919091179055610f21565b606f80546001600160a01b0319166001600160a01b0385169081179091555f90815260a260205260409020805460ff19168315151790555b610f29612605565b8015610f6e575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050505050565b6001600160a01b038083165f908152606b602090815260409182902082518084019093525463ffffffff81168352600160201b900490921691810182905290610fd65760405163828d566360e01b815260040160405180910390fd5b5f606a5f835f01518460200151604051602001610ff4929190613c7f565b60408051601f198184030181529181528151602092830120835290820192909252015f20546001600160a01b039081169150841681036110475760405163e273c4a160e01b815260040160405180910390fd5b6110518484612633565b61105c8133856126bf565b604080513381526001600160a01b0386811660208301528316818301526060810185905290517fb7f8fd4d1faf9b2929dc269f59c53e3a2bccc44e9950f33a568fcbcb37eb69a99181900360800190a150505050565b5f5f86866040516020016110c7929190613c7f565b6040516020818303038152906040528051906020012090505f60ff60f81b308360405180611ba00160405280611b6681526020016143fd611b66913989898960405160200161111893929190613e2d565b60408051601f19818403018152908290526111369291602001613e65565b6040516020818303038152906040528051906020012060405160200161118e94939291906001600160f81b031994909416845260609290921b6001600160601b03191660018401526015830152603582015260550190565b60408051808303601f19018152919052805160209091012098975050505050505050565b60a3546001600160a01b031633146111dd576040516357b738d160e11b815260040160405180910390fd5b6001600160a01b038082165f908152606b6020908152604080832081518083018352905463ffffffff8116808352600160201b9091049095168184018190529151909461122d9390929101613c7f565b60408051601f1981840301815291815281516020928301205f818152606a9093529120549091506001600160a01b0316158061128157505f818152606a60205260409020546001600160a01b038481169116145b1561129f5760405163e0c897a760e01b815260040160405180910390fd5b6001600160a01b0383165f908152606b6020908152604080832080546001600160c01b031916905560a290915290819020805460ff19169055517fc2ae0bd0ec0fd0352bfe5bacac49637af342c1e40f1b80a7f74440dc7fe3f06390611306908590613256565b60405180910390a1505050565b60685460ff161561133757604051630bc011ff60e21b815260040160405180910390fd5b606f546001600160a01b03166113605760405163dde3cda760e01b815260040160405180910390fd5b606f54611376906001600160a01b031685612633565b61138486868686868661205e565b505050505050565b60a3546001600160a01b031633146113b7576040516357b738d160e11b815260040160405180910390fd5b606d546001600160a01b03166113e057604051634cb4711360e11b815260040160405180910390fd5b606f80546001600160a01b0319166001600160a01b0384169081179091555f81815260a26020908152604091829020805460ff19168515159081179091558251938452908301527fc7318b7ed6ba4f2908a3de396d8ab49b1dadb55db5b55123247a401f29ff8d82910160405180910390a15050565b6060611461826122c5565b61146a83612380565b61147384612434565b60405160200161148593929190613e2d565b6040516020818303038152906040529050919050565b5f806114b1600160201b63ffffffff8516613e93565b6114c19063ffffffff8616613eaa565b600881901c5f90815260696020526040902054600160ff9092169190911b90811614949350505050565b60685460ff161561150f57604051630bc011ff60e21b815260040160405180910390fd5b61151761274f565b60685463ffffffff8681166101009092041614611547576040516302caf51760e11b815260040160405180910390fd5b61157a8c8c8c8c8c6115755f8e8e8e8e8e8e8e604051611568929190613ebd565b60405180910390206109cb565b6127a8565b6001600160a01b038616158015611595575063ffffffff8716155b1561167357606f546001600160a01b0316611657575f6001600160a01b03851684825b6040519080825280601f01601f1916602001820160405280156115e2576020820181803683370190505b506040516115f09190613ecc565b5f6040518083038185875af1925050503d805f811461162a576040519150601f19603f3d011682016040523d82523d5f602084013e61162f565b606091505b505090508061165157604051630ce8f45160e31b815260040160405180910390fd5b5061189c565b606f5461166e906001600160a01b031685856126bf565b61189c565b606d546001600160a01b0387811691161480156116a15750606d5463ffffffff888116600160a01b90920416145b156116b8575f6001600160a01b03851684826115b8565b60685463ffffffff6101009091048116908816036116e45761166e6001600160a01b038716858561280a565b5f87876040516020016116f8929190613c7f565b60408051601f1981840301815291815281516020928301205f818152606a9093529120549091506001600160a01b03168061188e575f61176d8386868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061258b92505050565b905061177a8188886126bf565b80606a5f8581526020019081526020015f205f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555060405180604001604052808b63ffffffff1681526020018a6001600160a01b0316815250606b5f836001600160a01b03166001600160a01b031681526020019081526020015f205f820151815f015f6101000a81548163ffffffff021916908363ffffffff1602179055506020820151815f0160046101000a8154816001600160a01b0302191690836001600160a01b031602179055509050507f490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a398a8a838888604051611880959493929190613f0f565b60405180910390a150611899565b6118998187876126bf565b50505b5f516020615f835f395f51905f528a888887876040516118c0959493929190613f47565b60405180910390a16118d160018055565b505050505050505050505050565b60685460ff161561190357604051630bc011ff60e21b815260040160405180910390fd5b61190b61274f565b60685463ffffffff61010090910481169088160361193c576040516302caf51760e11b815260040160405180910390fd5b5f806060876001600160a01b038816611a1f5788341461196f5760405163b89240f560e01b815260040160405180910390fd5b606d54606e80546001600160a01b0383169650600160a01b90920463ffffffff1694509061199c90613ca9565b80601f01602080910402602001604051908101604052809291908181526020018280546119c890613ca9565b8015611a135780601f106119ea57610100808354040283529160200191611a13565b820191905f5260205f20905b8154815290600101906020018083116119f657829003601f168201915b50505050509150611c09565b3415611a3e5760405163798ee6f160e01b815260040160405180910390fd5b606f546001600160a01b0390811690891603611a6357611a5e888a612633565b611c09565b6001600160a01b038089165f908152606b602090815260409182902082518084019093525463ffffffff81168352600160201b900490921691810182905290151580611ab55750805163ffffffff1615155b15611ad557611ac4898b612633565b602081015181519095509350611bfc565b8515611ae757611ae7898b8989612866565b6040516370a0823160e01b81525f906001600160a01b038b16906370a0823190611b15903090600401613256565b602060405180830381865afa158015611b30573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611b549190613f79565b9050611b6b6001600160a01b038b1633308e612b17565b6040516370a0823160e01b81525f906001600160a01b038c16906370a0823190611b99903090600401613256565b602060405180830381865afa158015611bb4573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611bd89190613f79565b9050611be48282613f90565b6068548c9850610100900463ffffffff169650935050505b611c0589611456565b9250505b5f516020615f635f395f51905f525f84868e8e8688605354604051611c35989796959493929190613fa3565b60405180910390a1611c5b611c565f85878f8f8789805190602001206109cb565b612b4f565b8615611c6957611c696124e8565b50505050611c7660018055565b50505050505050565b60a3546001600160a01b03163314611caa576040516357b738d160e11b815260040160405180910390fd5b60a380546001600160a01b0319166001600160a01b0383169081179091556040517f32cf74f8a6d5f88593984d2cd52be5592bfa6884f5896175801a5069ef09cd6791611cf691613256565b60405180910390a150565b60685460ff1615611d2557604051630bc011ff60e21b815260040160405180910390fd5b60685463ffffffff8681166101009092041614611d55576040516302caf51760e11b815260040160405180910390fd5b611d778c8c8c8c8c61157560018e8e8e8e8e8e8e604051611568929190613ebd565b606f545f906001600160a01b0316611e2a57846001600160a01b031684888a8686604051602401611dab949392919061400f565b60408051601f198184030181529181526020820180516001600160e01b0316630c035af960e11b17905251611de09190613ecc565b5f6040518083038185875af1925050503d805f8114611e1a576040519150601f19603f3d011682016040523d82523d5f602084013e611e1f565b606091505b505080915050611edb565b606f54611e41906001600160a01b031686866126bf565b846001600160a01b031687898585604051602401611e62949392919061400f565b60408051601f198184030181529181526020820180516001600160e01b0316630c035af960e11b17905251611e979190613ecc565b5f604051808303815f865af19150503d805f8114611ed0576040519150601f19603f3d011682016040523d82523d5f602084013e611ed5565b606091505b50909150505b80611ef9576040516337e391c360e01b815260040160405180910390fd5b5f516020615f835f395f51905f528b89898888604051611f1d959493929190613f47565b60405180910390a150505050505050505050505050565b5f54610100900460ff1615808015611f5257505f54600160ff909116105b80611f725750611f613061257c565b158015611f7257505f5460ff166001145b611f8e5760405162461bcd60e51b8152600401610cf390613ce1565b5f805460ff191660011790558015611faf575f805461ff0019166101001790555b60405163f57ac68360e01b815260040160405180910390fd5b5f81611fd5868686610c1e565b1495945050505050565b5f611ff4600160201b63ffffffff8416613e93565b6120049063ffffffff8516613eaa565b600881901c5f8181526069602052604090208054600160ff851690811b918218928390559394509192919080821615611c7657604051630631b5f760e31b815260040160405180910390fd5b5f9182526020526040902090565b60685463ffffffff61010090910481169087160361208f576040516302caf51760e11b815260040160405180910390fd5b5f516020615f635f395f51905f526001606860019054906101000a900463ffffffff163389898988886053546040516120d09998979695949392919061403f565b60405180910390a1612107611c566001606860019054906101000a900463ffffffff16338a8a8a8989604051611568929190613ebd565b8215611384576113846124e8565b6001600160a01b038316158061213257506001600160a01b038216155b156121505760405163f6b2911f60e01b815260040160405180910390fd5b60685463ffffffff6101009091048116908516036121815760405163658b23ad60e01b815260040160405180910390fd5b6001600160a01b038281165f908152606b6020526040902054600160201b900416156121c0576040516317abdeeb60e21b815260040160405180910390fd5b5f84846040516020016121d4929190613c7f565b60408051808303601f1901815282825280516020918201205f818152606a835283812080546001600160a01b0319166001600160a01b038a8116918217909255868601865263ffffffff8c81168089528c8416878a01818152848752606b89528987209a518b54915194166001600160c01b031990911617600160201b93909516929092029390931790975560a2855291859020805460ff191689151590811790915585519182529381019590955292840192909252606083015291507fdbe8a5da6a7a916d9adfda9160167a0f8a3da415ee6610e810e753853597fce79060800160405180910390a15050505050565b60408051600481526024810182526020810180516001600160e01b03166306fdde0360e01b17905290516060915f9182916001600160a01b0386169161230b9190613ecc565b5f60405180830381855afa9150503d805f8114612343576040519150601f19603f3d011682016040523d82523d5f602084013e612348565b606091505b50915091508161237757604051806040016040528060078152602001664e4f5f4e414d4560c81b815250610bd4565b610bd481612c05565b60408051600481526024810182526020810180516001600160e01b03166395d89b4160e01b17905290516060915f9182916001600160a01b038616916123c69190613ecc565b5f60405180830381855afa9150503d805f81146123fe576040519150601f19603f3d011682016040523d82523d5f602084013e612403565b606091505b50915091508161237757604051806040016040528060098152602001681393d7d4d6535093d360ba1b815250610bd4565b60408051600481526024810182526020810180516001600160e01b031663313ce56760e01b17905290515f91829182916001600160a01b038616916124799190613ecc565b5f60405180830381855afa9150503d805f81146124b1576040519150601f19603f3d011682016040523d82523d5f602084013e6124b6565b606091505b50915091508180156124c9575080516020145b6124d4576012610bd4565b80806020019051810190610bd491906140ad565b6053546068805463ffffffff909216600160c81b0263ffffffff60c81b1990921691909117908190556001600160a01b03600160281b909104166333d6247d61252f610b43565b6040518263ffffffff1660e01b815260040161254d91815260200190565b5f604051808303815f87803b158015612564575f5ffd5b505af1158015612576573d5f5f3e3d5ffd5b50505050565b6001600160a01b03163b151590565b5f5f60405180611ba00160405280611b6681526020016143fd611b669139836040516020016125bb929190613e65565b6040516020818303038152906040529050838151602083015ff591506001600160a01b0382166125fe576040516305f7d84960e51b815260040160405180910390fd5b5092915050565b5f54610100900460ff1661262b5760405162461bcd60e51b8152600401610cf3906140c8565b610bfd612d84565b6001600160a01b0382165f90815260a2602052604090205460ff1615612668576108646001600160a01b038316333084612b17565b604051632770a7eb60e21b81526001600160a01b03831690639dc29fac906126969033908590600401614113565b5f604051808303815f87803b1580156126ad575f5ffd5b505af1158015611384573d5f5f3e3d5ffd5b6001600160a01b0383165f90815260a2602052604090205460ff16156126f8576126f36001600160a01b038416838361280a565b505050565b6040516340c10f1960e01b81526001600160a01b038416906340c10f19906127269085908590600401614113565b5f604051808303815f87803b15801561273d575f5ffd5b505af1158015611c76573d5f5f3e3d5ffd5b6002600154036127a15760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610cf3565b6002600155565b6127b6868686868686612daa565b60a4546127c39085612050565b60a48190556040805186815260208101929092527f3e5936f910a78eb5181813a939c8d4c3e4d85f87943f659380d82ac6221b0e92910160405180910390a1505050505050565b6126f38363a9059cbb60e01b8484604051602401612829929190614113565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152612eeb565b60018055565b5f612874600482848661412c565b61287d91614153565b9050632afa533160e01b6001600160e01b03198216016129b1575f8080808080806128ab896004818d61412c565b8101906128b89190614189565b96509650965096509650965096508a85146128e6576040516303fffc4b60e01b815260040160405180910390fd5b604080516001600160a01b0389811660248301528881166044830152606482018890526084820187905260ff861660a483015260c4820185905260e48083018590528351808403909101815261010490920183526020820180516001600160e01b031663d505accf60e01b1790529151918e16916129649190613ecc565b5f604051808303815f865af19150503d805f811461299d576040519150601f19603f3d011682016040523d82523d5f602084013e6129a2565b606091505b50505050505050505050610938565b6001600160e01b031981166323f2ebc360e21b146129e257604051637141605d60e11b815260040160405180910390fd5b5f808080808080806129f78a6004818e61412c565b810190612a0491906141d8565b975097509750975097509750975097508c6001600160a01b0316638fcbaf0c60e01b8989898989898989604051602401612a899897969594939291906001600160a01b039889168152969097166020870152604086019490945260608501929092521515608084015260ff1660a083015260c082015260e08101919091526101000190565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051612ac79190613ecc565b5f604051808303815f865af19150503d805f8114612b00576040519150601f19603f3d011682016040523d82523d5f602084013e612b05565b606091505b50505050505050505050505050505050565b6040516001600160a01b03808516602483015283166044820152606481018290526125769085906323b872dd60e01b90608401612829565b806001612b5e60206002614339565b612b689190613f90565b60535410612b89576040516377ae67b360e11b815260040160405180910390fd5b5f60535f8154612b9890614344565b918290555090505f5b6020811015612bfc578082901c600116600103612bd4578260338260208110612bcc57612bcc613c3b565b015550505050565b612bf260338260208110612bea57612bea613c3b565b015484612050565b9250600101612ba1565b506126f361435c565b60606040825110612c2457818060200190518101906108c99190614370565b8151602003612d51575f5b602081108015612c5e5750828181518110612c4c57612c4c613c3b565b01602001516001600160f81b03191615155b15612c755780612c6d81614344565b915050612c2f565b805f03612cac5750506040805180820190915260128152714e4f545f56414c49445f454e434f44494e4760701b6020820152919050565b5f816001600160401b03811115612cc557612cc561315b565b6040519080825280601f01601f191660200182016040528015612cef576020820181803683370190505b5090505f5b82811015612d4957848181518110612d0e57612d0e613c3b565b602001015160f81c60f81b828281518110612d2b57612d2b613c3b565b60200101906001600160f81b03191690815f1a905350600101612cf4565b509392505050565b50506040805180820190915260128152714e4f545f56414c49445f454e434f44494e4760701b602082015290565b919050565b5f54610100900460ff166128605760405162461bcd60e51b8152600401610cf3906140c8565b6068545f90600160281b90046001600160a01b031663257b3632612dce8686612fbc565b6040518263ffffffff1660e01b8152600401612dec91815260200190565b6020604051808303815f875af1158015612e08573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612e2c9190613f79565b9050805f03612e4d57604051622f6fad60e01b815260040160405180910390fd5b5f80600160401b871615612e8c57869150612e6a848a8489611fc8565b612e87576040516338105f3b60e21b815260040160405180910390fd5b612ed6565b602087901c612e9c816001613c63565b9150879250612eb7612eaf868c86610c1e565b8a8389611fc8565b612ed4576040516338105f3b60e21b815260040160405180910390fd5b505b612ee08282612fce565b505050505050505050565b5f612f3f826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166130409092919063ffffffff16565b8051909150156126f35780806020019051810190612f5d91906143e1565b6126f35760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610cf3565b5f612fc78383612050565b9392505050565b5f612fe3600160201b63ffffffff8416613e93565b612ff39063ffffffff8516613eaa565b600881901c5f8181526069602052604081208054600160ff861690811b91821892839055949550929392918183169003611c7657604051630c8d9eab60e31b815260040160405180910390fd5b6060610bd484845f85855f5f866001600160a01b031685876040516130659190613ecc565b5f6040518083038185875af1925050503d805f811461309f576040519150601f19603f3d011682016040523d82523d5f602084013e6130a4565b606091505b50915091506130b5878383876130c0565b979650505050505050565b6060831561312c5782515f03613125576130d98561257c565b6131255760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610cf3565b5081610bd4565b610bd483838151156131415781518083602001fd5b8060405162461bcd60e51b8152600401610cf391906133db565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b03811182821017156131975761319761315b565b604052919050565b5f6001600160401b038211156131b7576131b761315b565b5060051b60200190565b5f602082840312156131d1575f5ffd5b81356001600160401b038111156131e6575f5ffd5b8201601f810184136131f6575f5ffd5b80356132096132048261319f565b61316f565b8082825260208201915060208360051b85010192508683111561322a575f5ffd5b6020840193505b8284101561324c578335825260209384019390910190613231565b9695505050505050565b6001600160a01b0391909116815260200190565b803563ffffffff81168114612d7f575f5ffd5b6001600160a01b0381168114613291575f5ffd5b50565b5f5f604083850312156132a5575f5ffd5b6132ae8361326a565b915060208301356132be8161327d565b809150509250929050565b8015158114613291575f5ffd5b5f5f83601f8401126132e6575f5ffd5b5081356001600160401b038111156132fc575f5ffd5b602083019150836020828501011115613313575f5ffd5b9250929050565b5f5f5f5f5f6080868803121561332e575f5ffd5b6133378661326a565b945060208601356133478161327d565b93506040860135613357816132c9565b925060608601356001600160401b03811115613371575f5ffd5b61337d888289016132d6565b969995985093965092949392505050565b5f5b838110156133a8578181015183820152602001613390565b50505f910152565b5f81518084526133c781602086016020860161338e565b601f01601f19169290920160200192915050565b602081525f612fc760208301846133b0565b5f602082840312156133fd575f5ffd5b8135612fc78161327d565b60ff81168114613291575f5ffd5b5f5f5f5f5f5f5f60e0888a03121561342c575f5ffd5b873561343781613408565b96506134456020890161326a565b955060408801356134558161327d565b94506134636060890161326a565b935060808801356134738161327d565b9699959850939692959460a0840135945060c09093013592915050565b5f82601f83011261349f575f5ffd5b81356134ad6132048261319f565b8082825260208201915060208360051b8601019250858311156134ce575f5ffd5b602085015b838110156134f45780356134e68161327d565b8352602092830192016134d3565b5095945050505050565b5f82601f83011261350d575f5ffd5b813561351b6132048261319f565b8082825260208201915060208360051b86010192508583111561353c575f5ffd5b602085015b838110156134f4578035613554816132c9565b835260209283019201613541565b5f5f5f5f60808587031215613575575f5ffd5b84356001600160401b0381111561358a575f5ffd5b8501601f8101871361359a575f5ffd5b80356135a86132048261319f565b8082825260208201915060208360051b8501019250898311156135c9575f5ffd5b6020840193505b828410156135f2576135e18461326a565b8252602093840193909101906135d0565b965050505060208501356001600160401b0381111561360f575f5ffd5b61361b87828801613490565b93505060408501356001600160401b03811115613636575f5ffd5b61364287828801613490565b92505060608501356001600160401b0381111561365d575f5ffd5b613669878288016134fe565b91505092959194509250565b5f5f5f60608486031215613687575f5ffd5b6136908461326a565b925060208401356136a08161327d565b915060408401356136b08161327d565b809150509250925092565b5f602082840312156136cb575f5ffd5b5035919050565b8061040081018310156108c9575f5ffd5b5f5f5f61044084860312156136f6575f5ffd5b8335925061370785602086016136d2565b9150613716610420850161326a565b90509250925092565b5f6001600160401b038211156137375761373761315b565b50601f01601f191660200190565b5f82601f830112613754575f5ffd5b8135602083015f6137676132048461371f565b905082815285838301111561377a575f5ffd5b828260208301375f92810160200192909252509392505050565b5f5f5f5f5f5f5f5f5f6101208a8c0312156137ad575f5ffd5b6137b68a61326a565b985060208a01356137c68161327d565b97506137d460408b0161326a565b965060608a01356137e48161327d565b955060808a01356137f48161327d565b945060a08a01356001600160401b0381111561380e575f5ffd5b61381a8c828d01613745565b94505060c08a013561382b8161327d565b925060e08a013561383b8161327d565b91506101008a013561384c816132c9565b809150509295985092959850929598565b5f5f6040838503121561386e575f5ffd5b82356138798161327d565b946020939093013593505050565b5f5f5f5f5f60a0868803121561389b575f5ffd5b6138a48661326a565b945060208601356138b48161327d565b935060408601356001600160401b038111156138ce575f5ffd5b6138da88828901613745565b93505060608601356001600160401b038111156138f5575f5ffd5b61390188828901613745565b925050608086013561391281613408565b809150509295509295909350565b5f5f5f5f5f5f60a08789031215613935575f5ffd5b61393e8761326a565b9550602087013561394e8161327d565b9450604087013593506060870135613965816132c9565b925060808701356001600160401b0381111561397f575f5ffd5b61398b89828a016132d6565b979a9699509497509295939492505050565b5f5f604083850312156139ae575f5ffd5b82356139b98161327d565b915060208301356132be816132c9565b5f5f604083850312156139da575f5ffd5b6139e38361326a565b91506139f16020840161326a565b90509250929050565b5f5f5f5f5f5f5f5f5f5f5f5f6109208d8f031215613a16575f5ffd5b613a208e8e6136d2565b9b50613a308e6104008f016136d2565b9a506108008d013599506108208d013598506108408d01359750613a576108608e0161326a565b9650613a676108808e013561327d565b6108808d01359550613a7c6108a08e0161326a565b94506108c08d0135613a8d8161327d565b93506108e08d013592506001600160401b036109008e01351115613aaf575f5ffd5b613ac08e6109008f01358f016132d6565b81935080925050509295989b509295989b509295989b565b5f5f5f5f5f5f5f60c0888a031215613aee575f5ffd5b613af78861326a565b96506020880135613b078161327d565b9550604088013594506060880135613b1e8161327d565b93506080880135613b2e816132c9565b925060a08801356001600160401b03811115613b48575f5ffd5b613b548a828b016132d6565b989b979a50959850939692959293505050565b5f5f5f5f5f5f60c08789031215613b7c575f5ffd5b613b858761326a565b95506020870135613b958161327d565b9450613ba36040880161326a565b93506060870135613bb38161327d565b92506080870135613bc38161327d565b915060a08701356001600160401b03811115613bdd575f5ffd5b613be989828a01613745565b9150509295509295509295565b5f5f5f5f6104608587031215613c0a575f5ffd5b84359350613c1b86602087016136d2565b9250613c2a610420860161326a565b939692955092936104400135925050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b63ffffffff81811683821601908111156108c9576108c9613c4f565b60e09290921b6001600160e01b031916825260601b6001600160601b031916600482015260180190565b600181811c90821680613cbd57607f821691505b602082108103613cdb57634e487b7160e01b5f52602260045260245ffd5b50919050565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b601f8211156126f357805f5260205f20601f840160051c81016020851015613d545750805b601f840160051c820191505b81811015610938575f8155600101613d60565b81516001600160401b03811115613d8c57613d8c61315b565b613da081613d9a8454613ca9565b84613d2f565b6020601f821160018114613dd2575f8315613dbb5750848201515b5f19600385901b1c1916600184901b178455610938565b5f84815260208120601f198516915b82811015613e015787850151825560209485019460019092019101613de1565b5084821015613e1e57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b606081525f613e3f60608301866133b0565b8281036020840152613e5181866133b0565b91505060ff83166040830152949350505050565b5f8351613e7681846020880161338e565b835190830190613e8a81836020880161338e565b01949350505050565b80820281158282048414176108c9576108c9613c4f565b808201808211156108c9576108c9613c4f565b818382375f9101908152919050565b5f8251613edd81846020870161338e565b9190910192915050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b63ffffffff861681526001600160a01b038581166020830152841660408201526080606082018190525f906130b59083018486613ee7565b94855263ffffffff9390931660208501526001600160a01b039182166040850152166060830152608082015260a00190565b5f60208284031215613f89575f5ffd5b5051919050565b818103818111156108c9576108c9613c4f565b60ff8916815263ffffffff88811660208301526001600160a01b03888116604084015290871660608301528516608082015260a0810184905261010060c082018190525f90613ff4908301856133b0565b905063ffffffff831660e08301529998505050505050505050565b6001600160a01b038516815263ffffffff841660208201526060604082018190525f9061324c9083018486613ee7565b60ff8a16815263ffffffff89811660208301526001600160a01b03898116604084015290881660608301528616608082015260a0810185905261010060c082018190525f906140919083018587613ee7565b905063ffffffff831660e08301529a9950505050505050505050565b5f602082840312156140bd575f5ffd5b8151612fc781613408565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b6001600160a01b03929092168252602082015260400190565b5f5f8585111561413a575f5ffd5b83861115614146575f5ffd5b5050820193919092039150565b80356001600160e01b031981169060048410156125fe576001600160e01b031960049490940360031b84901b1690921692915050565b5f5f5f5f5f5f5f60e0888a03121561419f575f5ffd5b87356141aa8161327d565b965060208801356141ba8161327d565b95506040880135945060608801359350608088013561347381613408565b5f5f5f5f5f5f5f5f610100898b0312156141f0575f5ffd5b88356141fb8161327d565b9750602089013561420b8161327d565b965060408901359550606089013594506080890135614229816132c9565b935060a089013561423981613408565b979a969950949793969295929450505060c08201359160e0013590565b6001815b60018411156142915780850481111561427557614275613c4f565b600184161561428357908102905b60019390931c92800261425a565b935093915050565b5f826142a7575060016108c9565b816142b357505f6108c9565b81600181146142c957600281146142d3576142ef565b60019150506108c9565b60ff8411156142e4576142e4613c4f565b50506001821b6108c9565b5060208310610133831016604e8410600b8410161715614312575081810a6108c9565b61431e5f198484614256565b805f190482111561433157614331613c4f565b029392505050565b5f612fc78383614299565b5f6001820161435557614355613c4f565b5060010190565b634e487b7160e01b5f52600160045260245ffd5b5f60208284031215614380575f5ffd5b81516001600160401b03811115614395575f5ffd5b8201601f810184136143a5575f5ffd5b80516143b36132048261371f565b8181528560208385010111156143c7575f5ffd5b6143d882602083016020860161338e565b95945050505050565b5f602082840312156143f1575f5ffd5b8151612fc7816132c956fe6101006040523480156200001257600080fd5b5060405162001b6638038062001b6683398101604081905262000035916200028d565b82826003620000458382620003a1565b506004620000548282620003a1565b50503360c0525060ff811660e052466080819052620000739062000080565b60a052506200046d915050565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f620000ad6200012e565b805160209182012060408051808201825260018152603160f81b90840152805192830193909352918101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc66060820152608081018390523060a082015260c001604051602081830303815290604052805190602001209050919050565b6060600380546200013f9062000312565b80601f01602080910402602001604051908101604052809291908181526020018280546200016d9062000312565b8015620001be5780601f106200019257610100808354040283529160200191620001be565b820191906000526020600020905b815481529060010190602001808311620001a057829003601f168201915b5050505050905090565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620001f057600080fd5b81516001600160401b03808211156200020d576200020d620001c8565b604051601f8301601f19908116603f01168101908282118183101715620002385762000238620001c8565b816040528381526020925086838588010111156200025557600080fd5b600091505b838210156200027957858201830151818301840152908201906200025a565b600093810190920192909252949350505050565b600080600060608486031215620002a357600080fd5b83516001600160401b0380821115620002bb57600080fd5b620002c987838801620001de565b94506020860151915080821115620002e057600080fd5b50620002ef86828701620001de565b925050604084015160ff811681146200030757600080fd5b809150509250925092565b600181811c908216806200032757607f821691505b6020821081036200034857634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200039c57600081815260208120601f850160051c81016020861015620003775750805b601f850160051c820191505b81811015620003985782815560010162000383565b5050505b505050565b81516001600160401b03811115620003bd57620003bd620001c8565b620003d581620003ce845462000312565b846200034e565b602080601f8311600181146200040d5760008415620003f45750858301515b600019600386901b1c1916600185901b17855562000398565b600085815260208120601f198616915b828110156200043e578886015182559484019460019091019084016200041d565b50858210156200045d5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a05160c05160e0516116aa620004bc6000396000610237015260008181610307015281816105c001526106a70152600061053a015260008181610379015261050401526116aa6000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c806370a08231116100d8578063a457c2d71161008c578063d505accf11610066578063d505accf1461039b578063dd62ed3e146103ae578063ffa1ad74146103f457600080fd5b8063a457c2d71461034e578063a9059cbb14610361578063cd0d00961461037457600080fd5b806395d89b41116100bd57806395d89b41146102e75780639dc29fac146102ef578063a3c573eb1461030257600080fd5b806370a08231146102915780637ecebe00146102c757600080fd5b806330adf81f1161012f5780633644e515116101145780633644e51514610261578063395093511461026957806340c10f191461027c57600080fd5b806330adf81f14610209578063313ce5671461023057600080fd5b806318160ddd1161016057806318160ddd146101bd57806320606b70146101cf57806323b872dd146101f657600080fd5b806306fdde031461017c578063095ea7b31461019a575b600080fd5b610184610430565b60405161019191906113e4565b60405180910390f35b6101ad6101a8366004611479565b6104c2565b6040519015158152602001610191565b6002545b604051908152602001610191565b6101c17f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f81565b6101ad6102043660046114a3565b6104dc565b6101c17f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c981565b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610191565b6101c1610500565b6101ad610277366004611479565b61055c565b61028f61028a366004611479565b6105a8565b005b6101c161029f3660046114df565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6101c16102d53660046114df565b60056020526000908152604090205481565b610184610680565b61028f6102fd366004611479565b61068f565b6103297f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610191565b6101ad61035c366004611479565b61075e565b6101ad61036f366004611479565b61082f565b6101c17f000000000000000000000000000000000000000000000000000000000000000081565b61028f6103a9366004611501565b61083d565b6101c16103bc366004611574565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6101846040518060400160405280600181526020017f310000000000000000000000000000000000000000000000000000000000000081525081565b60606003805461043f906115a7565b80601f016020809104026020016040519081016040528092919081815260200182805461046b906115a7565b80156104b85780601f1061048d576101008083540402835291602001916104b8565b820191906000526020600020905b81548152906001019060200180831161049b57829003601f168201915b5050505050905090565b6000336104d0818585610b73565b60019150505b92915050565b6000336104ea858285610d27565b6104f5858585610dfe565b506001949350505050565b60007f00000000000000000000000000000000000000000000000000000000000000004614610537576105324661106d565b905090565b507f000000000000000000000000000000000000000000000000000000000000000090565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091906104d090829086906105a3908790611629565b610b73565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610672576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f546f6b656e577261707065643a3a6f6e6c794272696467653a204e6f7420506f60448201527f6c79676f6e5a6b45564d4272696467650000000000000000000000000000000060648201526084015b60405180910390fd5b61067c8282611135565b5050565b60606004805461043f906115a7565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610754576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f546f6b656e577261707065643a3a6f6e6c794272696467653a204e6f7420506f60448201527f6c79676f6e5a6b45564d427269646765000000000000000000000000000000006064820152608401610669565b61067c8282611228565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919083811015610822576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f0000000000000000000000000000000000000000000000000000006064820152608401610669565b6104f58286868403610b73565b6000336104d0818585610dfe565b834211156108cc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f546f6b656e577261707065643a3a7065726d69743a204578706972656420706560448201527f726d6974000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff8716600090815260056020526040812080547f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9918a918a918a9190866109268361163c565b9091555060408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090506000610991610500565b6040517f19010000000000000000000000000000000000000000000000000000000000006020820152602281019190915260428101839052606201604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600080855291840180845281905260ff89169284019290925260608301879052608083018690529092509060019060a0016020604051602081039080840390855afa158015610a55573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811615801590610ad057508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b610b5c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f546f6b656e577261707065643a3a7065726d69743a20496e76616c696420736960448201527f676e6174757265000000000000000000000000000000000000000000000000006064820152608401610669565b610b678a8a8a610b73565b50505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8316610c15576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff8216610cb8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610df85781811015610deb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610669565b610df88484848403610b73565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610ea1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff8216610f44576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610ffa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3610df8565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f611098610430565b8051602091820120604080518082018252600181527f310000000000000000000000000000000000000000000000000000000000000090840152805192830193909352918101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc66060820152608081018390523060a082015260c001604051602081830303815290604052805190602001209050919050565b73ffffffffffffffffffffffffffffffffffffffff82166111b2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610669565b80600260008282546111c49190611629565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b73ffffffffffffffffffffffffffffffffffffffff82166112cb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f73000000000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090205481811015611381576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f63650000000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff83166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101610d1a565b600060208083528351808285015260005b81811015611411578581018301518582016040015282016113f5565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461147457600080fd5b919050565b6000806040838503121561148c57600080fd5b61149583611450565b946020939093013593505050565b6000806000606084860312156114b857600080fd5b6114c184611450565b92506114cf60208501611450565b9150604084013590509250925092565b6000602082840312156114f157600080fd5b6114fa82611450565b9392505050565b600080600080600080600060e0888a03121561151c57600080fd5b61152588611450565b965061153360208901611450565b95506040880135945060608801359350608088013560ff8116811461155757600080fd5b9699959850939692959460a0840135945060c09093013592915050565b6000806040838503121561158757600080fd5b61159083611450565b915061159e60208401611450565b90509250929050565b600181811c908216806115bb57607f821691505b6020821081036115f4577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156104d6576104d66115fa565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361166d5761166d6115fa565b506001019056fea26469706673582212208d88fee561cff7120d381c345cfc534cef8229a272dc5809d4bbb685ad67141164736f6c63430008110033501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4da264697066735822122006b1493f1f40587ba906bedbfa38d1efc36f3d57b050f08104c3aa5a4fd9dba164736f6c634300081c0033",
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

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridgel2sovereignchainpessimistic

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

// Bridgel2sovereignchainpessimisticMetaData contains all meta data concerning the Bridgel2sovereignchainpessimistic contract.
var Bridgel2sovereignchainpessimisticMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountDoesNotMatchMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DestinationNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmergencyStateNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EtherTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTokenWrappedDeployment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputArraysLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSmtProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSovereignWETHAddressParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MsgValueNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeTokenIsEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoValueInMessagesOnGasTokenNetworks\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyBridgeManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OriginNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyMapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyUpdated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotMapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotRemapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WETHRemappingNotSupportedOnGasTokenNetworks\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"depositCount\",\"type\":\"uint32\"}],\"name\":\"BridgeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"legacyTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"updatedTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MigrateLegacyToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wrappedTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"NewWrappedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignTokenAddress\",\"type\":\"address\"}],\"name\":\"RemoveLegacySovereignTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridgeManager\",\"type\":\"address\"}],\"name\":\"SetBridgeManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"SetSovereignTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignWETHTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"SetSovereignWETHAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"leafIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"sourceBridgeNetwork\",\"type\":\"uint32\"}],\"name\":\"UnsetClaim\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASE_INIT_BYTECODE_WRAPPED_TOKEN\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETHToken\",\"outputs\":[{\"internalType\":\"contractTokenWrapped\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"permitData\",\"type\":\"bytes\"}],\"name\":\"bridgeAsset\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountWETH\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessageWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"calculateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"calculateTokenWrapperAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimedBitMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"name\":\"getLeafValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenWrappedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_polygonRollupManager\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_bridgeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sovereignWETHAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_sovereignWETHAddressIsNotMintable\",\"type\":\"bool\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"leafIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"sourceBridgeNetwork\",\"type\":\"uint32\"}],\"name\":\"isClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdatedDepositCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"legacyTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"migrateLegacyToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"polygonRollupManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"precalculatedWrapperAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"legacySovereignTokenAddress\",\"type\":\"address\"}],\"name\":\"removeLegacySovereignTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeManager\",\"type\":\"address\"}],\"name\":\"setBridgeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"originNetworks\",\"type\":\"uint32[]\"},{\"internalType\":\"address[]\",\"name\":\"originTokenAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"sovereignTokenAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"isNotMintable\",\"type\":\"bool[]\"}],\"name\":\"setMultipleSovereignTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sovereignWETHTokenAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"setSovereignWETHAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenInfoToWrappedToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"leafIndexes\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"sourceBridgeNetworks\",\"type\":\"uint32[]\"}],\"name\":\"unsetMultipleClaimedBitmap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"verifyMerkleProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrappedAddress\",\"type\":\"address\"}],\"name\":\"wrappedAddressIsNotMintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wrappedTokenToTokenInfo\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b506200001b6200002b565b620000256200002b565b620000ea565b5f54610100900460ff1615620000975760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161015620000e8575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6176a480620000f85f395ff3fe6080604052600436106102b6575f3560e01c80638c0dd47011610170578063c0f49163116100d1578063dbc1697611610087578063f5efcd7911610062578063f5efcd791461092c578063f811bff71461094b578063fb5708341461096a575f80fd5b8063dbc1697614610339578063eabd372a146108e2578063ee25560b14610901575f80fd5b8063ccaa2d11116100b7578063ccaa2d111461087b578063cd5865791461089a578063d02103ca146108ad575f80fd5b8063c0f491631461082e578063cc4616321461085c575f80fd5b8063b8b284d011610126578063be5831c71161010c578063be5831c7146107b7578063bf130d7f146107f0578063c00f14ab1461080f575f80fd5b8063b8b284d014610777578063bab161bf14610796575f80fd5b80639e76158f116101565780639e76158f1461071a578063aaa13cc214610739578063b458696214610758575f80fd5b80638c0dd470146106cf5780638ed7e3f2146106ee575f80fd5b80633e1970431161021a57806379e2cf97116101d057806383c43a55116101b657806383c43a551461067d57806383f24403146106915780638781a5c5146106b0575f80fd5b806379e2cf971461062857806381b1c1741461063c575f80fd5b806357cfbee31161020057806357cfbee3146105d65780635ca1e165146105f55780637843298b14610609575f80fd5b80633e197043146104bc5780634b2f336d146105aa575f80fd5b806327aef4e81161026f578063318aee3d11610255578063318aee3d146103c55780633c351e10146104475780633cbc795b14610473575f80fd5b806327aef4e8146103815780632dfdf0b5146103a2575f80fd5b80632072f6c51161029f5780632072f6c51461033957806322e95f2c1461034f578063240ff3781461036e575f80fd5b806314cc01a0146102ba57806315064c9614610310575b5f80fd5b3480156102c5575f80fd5b5060a3546102e69073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561031b575f80fd5b506068546103299060ff1681565b6040519015158152602001610307565b348015610344575f80fd5b5061034d610989565b005b34801561035a575f80fd5b506102e66103693660046147fa565b6109bb565b61034d61037c366004614881565b610a5d565b34801561038c575f80fd5b50610395610b0c565b6040516103079190614961565b3480156103ad575f80fd5b506103b760535481565b604051908152602001610307565b3480156103d0575f80fd5b506104166103df366004614973565b606b6020525f908152604090205463ffffffff811690640100000000900473ffffffffffffffffffffffffffffffffffffffff1682565b6040805163ffffffff909316835273ffffffffffffffffffffffffffffffffffffffff909116602083015201610307565b348015610452575f80fd5b50606d546102e69073ffffffffffffffffffffffffffffffffffffffff1681565b34801561047e575f80fd5b50606d546104a79074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff9091168152602001610307565b3480156104c7575f80fd5b506103b76104d636600461499c565b6040517fff0000000000000000000000000000000000000000000000000000000000000060f889901b1660208201527fffffffff0000000000000000000000000000000000000000000000000000000060e088811b821660218401527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606089811b821660258601529188901b909216603984015285901b16603d82015260518101839052607181018290525f90609101604051602081830303815290604052805190602001209050979650505050505050565b3480156105b5575f80fd5b50606f546102e69073ffffffffffffffffffffffffffffffffffffffff1681565b3480156105e1575f80fd5b5061034d6105f0366004614b85565b610b98565b348015610600575f80fd5b506103b7610ccd565b348015610614575f80fd5b506102e6610623366004614c85565b610d56565b348015610633575f80fd5b5061034d610d7f565b348015610647575f80fd5b506102e6610656366004614ccb565b606a6020525f908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b348015610688575f80fd5b50610395610db8565b34801561069c575f80fd5b506103b76106ab366004614cf3565b610dd7565b3480156106bb575f80fd5b5061034d6106ca366004614d2f565b610e6e565b3480156106da575f80fd5b5061034d6106e9366004614e21565b610f58565b3480156106f9575f80fd5b50606c546102e69073ffffffffffffffffffffffffffffffffffffffff1681565b348015610725575f80fd5b5061034d610734366004614eeb565b611475565b348015610744575f80fd5b506102e6610753366004614f15565b61167a565b348015610763575f80fd5b5061034d610772366004614973565b611833565b348015610782575f80fd5b5061034d610791366004614fab565b611aaa565b3480156107a1575f80fd5b506068546104a790610100900463ffffffff1681565b3480156107c2575f80fd5b506068546104a790790100000000000000000000000000000000000000000000000000900463ffffffff1681565b3480156107fb575f80fd5b5061034d61080a366004615029565b611b6f565b34801561081a575f80fd5b50610395610829366004614973565b611cc8565b348015610839575f80fd5b50610329610848366004614973565b60a26020525f908152604090205460ff1681565b348015610867575f80fd5b50610329610876366004615055565b611d0d565b348015610886575f80fd5b5061034d610895366004615086565b611d5e565b61034d6108a836600461516a565b6123c4565b3480156108b8575f80fd5b506068546102e69065010000000000900473ffffffffffffffffffffffffffffffffffffffff1681565b3480156108ed575f80fd5b5061034d6108fc366004614973565b61292c565b34801561090c575f80fd5b506103b761091b366004614ccb565b60696020525f908152604090205481565b348015610937575f80fd5b5061034d610946366004615086565b6129f6565b348015610956575f80fd5b5061034d6109653660046151fa565b612d7d565b348015610975575f80fd5b5061032961098436600461528a565b612edf565b6040517f441845b100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160e084901b7fffffffff0000000000000000000000000000000000000000000000000000000016602080830191909152606084901b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602483015282516018818403018152603890920183528151918101919091205f908152606a909152205473ffffffffffffffffffffffffffffffffffffffff165b92915050565b60685460ff1615610a9a576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3415801590610ac05750606f5473ffffffffffffffffffffffffffffffffffffffff1615155b15610af7576040517f6f625c4000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b05858534868686612ef6565b5050505050565b606e8054610b19906152cf565b80601f0160208091040260200160405190810160405280929190818152602001828054610b45906152cf565b8015610b905780601f10610b6757610100808354040283529160200191610b90565b820191905f5260205f20905b815481529060010190602001808311610b7357829003601f168201915b505050505081565b60a35473ffffffffffffffffffffffffffffffffffffffff163314610be9576040517faf6e71a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82518451141580610bfc57508151845114155b80610c0957508051845114155b15610c40576040517f869e93ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b8251811015610b0557610cbb858281518110610c6057610c60615320565b6020026020010151858381518110610c7a57610c7a615320565b6020026020010151858481518110610c9457610c94615320565b6020026020010151858581518110610cae57610cae615320565b6020026020010151612fd9565b80610cc58161537a565b915050610c42565b6053545f90819081805b6020811015610d4d578083901c600116600103610d1c57610d1560338260208110610d0457610d04615320565b0154855f9182526020526040902090565b9350610d2c565b5f84815260208390526040902093505b5f828152602083905260409020915080610d458161537a565b915050610cd7565b50919392505050565b5f610d778484610d65856132c2565b610d6e866133ce565b610753876134d1565b949350505050565b605354606854790100000000000000000000000000000000000000000000000000900463ffffffff161015610db657610db66135c0565b565b60405180611ba00160405280611b668152602001615b09611b66913981565b5f83815b6020811015610e6557600163ffffffff8516821c81169003610e2757610e20858260208110610e0c57610e0c615320565b6020020135835f9182526020526040902090565b9150610e53565b610e5082868360208110610e3d57610e3d615320565b60200201355f9182526020526040902090565b91505b80610e5d8161537a565b915050610ddb565b50949350505050565b60a35473ffffffffffffffffffffffffffffffffffffffff163314610ebf576040517faf6e71a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051825114610efa576040517f869e93ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b8251811015610f5357610f41838281518110610f1a57610f1a615320565b6020026020010151838381518110610f3457610f34615320565b6020026020010151613691565b80610f4b8161537a565b915050610efc565b505050565b5f54610100900460ff1615808015610f7657505f54600160ff909116105b80610f8f5750303b158015610f8f57505f5460ff166001145b611020576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561107c575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b606880547fffffffffffffff000000000000000000000000000000000000000000000000ff1661010063ffffffff8d16027fffffffffffffff0000000000000000000000000000000000000000ffffffffff16176501000000000073ffffffffffffffffffffffffffffffffffffffff8a81169190910291909117909155606c80547fffffffffffffffffffffffff00000000000000000000000000000000000000009081168984161790915560a3805490911686831617905589166111d75763ffffffff88161561117a576040517f1a874c1200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff831615158061119b5750815b156111d2576040517f1cdc46ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113ff565b606d805463ffffffff8a1674010000000000000000000000000000000000000000027fffffffffffffffff00000000000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff8c1617179055606e61124486826153f6565b5073ffffffffffffffffffffffffffffffffffffffff83166113845781151560010361129c576040517f1cdc46ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61133a5f801b601260405160200161132691906060808252600d908201527f5772617070656420457468657200000000000000000000000000000000000000608082015260a0602082018190526004908201527f574554480000000000000000000000000000000000000000000000000000000060c082015260ff91909116604082015260e00190565b604051602081830303815290604052613759565b606f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff929092169190911790556113ff565b606f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091555f90815260a26020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168315151790555b6114076137f9565b8015611469575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8083165f908152606b602090815260409182902082518084019093525463ffffffff811683526401000000009004909216918101829052906114f8576040517f828d566300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f606a5f835f0151846020015160405160200161156b92919060e09290921b7fffffffff0000000000000000000000000000000000000000000000000000000016825260601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016600482015260180190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181528151602092830120835290820192909252015f205473ffffffffffffffffffffffffffffffffffffffff908116915084168103611602576040517fe273c4a100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61160c8484613897565b611617813385613967565b6040805133815273ffffffffffffffffffffffffffffffffffffffff86811660208301528316818301526060810185905290517fb7f8fd4d1faf9b2929dc269f59c53e3a2bccc44e9950f33a568fcbcb37eb69a99181900360800190a150505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e087901b1660208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606086901b1660248201525f9081906038016040516020818303038152906040528051906020012090505f60ff60f81b308360405180611ba00160405280611b668152602001615b09611b66913989898960405160200161172d9392919061550e565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526117699291602001615546565b604051602081830303815290604052805190602001206040516020016117f194939291907fff0000000000000000000000000000000000000000000000000000000000000094909416845260609290921b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660018401526015830152603582015260550190565b604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018152919052805160209091012098975050505050505050565b60a35473ffffffffffffffffffffffffffffffffffffffff163314611884576040517faf6e71a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8082165f908152606b6020908152604080832081518083018352905463ffffffff811680835264010000000090910490951681840181905291519094611937939092910160e09290921b7fffffffff0000000000000000000000000000000000000000000000000000000016825260601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016600482015260180190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301205f818152606a90935291205490915073ffffffffffffffffffffffffffffffffffffffff1615806119c357505f818152606a602052604090205473ffffffffffffffffffffffffffffffffffffffff8481169116145b156119fa576040517fe0c897a700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83165f818152606b6020908152604080832080547fffffffffffffffff00000000000000000000000000000000000000000000000016905560a282529182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905590519182527fc2ae0bd0ec0fd0352bfe5bacac49637af342c1e40f1b80a7f74440dc7fe3f063910160405180910390a1505050565b60685460ff1615611ae7576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5473ffffffffffffffffffffffffffffffffffffffff16611b36576040517fdde3cda700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f54611b599073ffffffffffffffffffffffffffffffffffffffff1685613897565b611b67868686868686612ef6565b505050505050565b60a35473ffffffffffffffffffffffffffffffffffffffff163314611bc0576040517faf6e71a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606d5473ffffffffffffffffffffffffffffffffffffffff16611c0f576040517f9968e22600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84169081179091555f81815260a2602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168515159081179091558251938452908301527fc7318b7ed6ba4f2908a3de396d8ab49b1dadb55db5b55123247a401f29ff8d82910160405180910390a15050565b6060611cd3826132c2565b611cdc836133ce565b611ce5846134d1565b604051602001611cf79392919061550e565b6040516020818303038152906040529050919050565b5f80611d2464010000000063ffffffff8516615574565b611d349063ffffffff861661558b565b600881901c5f90815260696020526040902054600160ff9092169190911b90811614949350505050565b60685460ff1615611d9b576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60685463ffffffff8681166101009092041614611de4576040517f0595ea2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611ee68c8c8c8c8c611ee15f8e8e8e8e8e8e8e604051611e0592919061559e565b60405180910390206040517fff0000000000000000000000000000000000000000000000000000000000000060f889901b1660208201527fffffffff0000000000000000000000000000000000000000000000000000000060e088811b821660218401527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606089811b821660258601529188901b909216603984015285901b16603d82015260518101839052607181018290525f90609101604051602081830303815290604052805190602001209050979650505050505050565b613a34565b73ffffffffffffffffffffffffffffffffffffffff861661201a57606f5473ffffffffffffffffffffffffffffffffffffffff16611ff1575f73ffffffffffffffffffffffffffffffffffffffff851684825b6040519080825280601f01601f191660200182016040528015611f63576020820181803683370190505b50604051611f7191906155ad565b5f6040518083038185875af1925050503d805f8114611fab576040519150601f19603f3d011682016040523d82523d5f602084013e611fb0565b606091505b5050905080611feb576040517f6747a28800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5061234d565b606f546120159073ffffffffffffffffffffffffffffffffffffffff168585613967565b61234d565b606d5473ffffffffffffffffffffffffffffffffffffffff87811691161480156120665750606d5463ffffffff8881167401000000000000000000000000000000000000000090920416145b1561208a575f73ffffffffffffffffffffffffffffffffffffffff85168482611f39565b60685463ffffffff6101009091048116908816036120c35761201573ffffffffffffffffffffffffffffffffffffffff87168585613bd4565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e089901b1660208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606088901b1660248201525f90603801604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301205f818152606a90935291205490915073ffffffffffffffffffffffffffffffffffffffff168061233f575f6121c38386868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061375992505050565b90506121d0818888613967565b80606a5f8581526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060405180604001604052808b63ffffffff1681526020018a73ffffffffffffffffffffffffffffffffffffffff16815250606b5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f015f6101000a81548163ffffffff021916908363ffffffff1602179055506020820151815f0160046101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050507f490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a398a8a83888860405161233195949392919061560f565b60405180910390a15061234a565b61234a818787613967565b50505b604080518b815263ffffffff8916602082015273ffffffffffffffffffffffffffffffffffffffff88811682840152861660608201526080810185905290517f1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d9181900360a00190a1505050505050505050505050565b60685460ff1615612401576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612409613ca8565b60685463ffffffff610100909104811690881603612453576040517f0595ea2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8060608773ffffffffffffffffffffffffffffffffffffffff881661257a578834146124ac576040517fb89240f500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606d54606e805473ffffffffffffffffffffffffffffffffffffffff831696507401000000000000000000000000000000000000000090920463ffffffff169450906124f7906152cf565b80601f0160208091040260200160405190810160405280929190818152602001828054612523906152cf565b801561256e5780601f106125455761010080835404028352916020019161256e565b820191905f5260205f20905b81548152906001019060200180831161255157829003601f168201915b505050505091506127d4565b34156125b2576040517f798ee6f100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5473ffffffffffffffffffffffffffffffffffffffff908116908916036125e4576125df888a613897565b6127d4565b73ffffffffffffffffffffffffffffffffffffffff8089165f908152606b602090815260409182902082518084019093525463ffffffff811683526401000000009004909216918101829052901561265157612640898b613897565b6020810151815190955093506127c7565b851561266357612663898b8989613d1b565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f9073ffffffffffffffffffffffffffffffffffffffff8b16906370a0823190602401602060405180830381865afa1580156126cd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906126f19190615654565b905061271573ffffffffffffffffffffffffffffffffffffffff8b1633308e614100565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f9073ffffffffffffffffffffffffffffffffffffffff8c16906370a0823190602401602060405180830381865afa15801561277f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906127a39190615654565b90506127af828261566b565b6068548c9850610100900463ffffffff169650935050505b6127d089611cc8565b9250505b7f501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b5f84868e8e868860535460405161281398979695949392919061567e565b60405180910390a16129086129035f85878f8f8789805190602001206040517fff0000000000000000000000000000000000000000000000000000000000000060f889901b1660208201527fffffffff0000000000000000000000000000000000000000000000000000000060e088811b821660218401527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606089811b821660258601529188901b909216603984015285901b16603d82015260518101839052607181018290525f90609101604051602081830303815290604052805190602001209050979650505050505050565b61415e565b8615612916576129166135c0565b5050505061292360018055565b50505050505050565b60a35473ffffffffffffffffffffffffffffffffffffffff16331461297d576040517faf6e71a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60a380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f32cf74f8a6d5f88593984d2cd52be5592bfa6884f5896175801a5069ef09cd679060200160405180910390a150565b60685460ff1615612a33576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60685463ffffffff8681166101009092041614612a7c576040517f0595ea2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612a9e8c8c8c8c8c611ee160018e8e8e8e8e8e8e604051611e0592919061559e565b606f545f9073ffffffffffffffffffffffffffffffffffffffff16612bb7578473ffffffffffffffffffffffffffffffffffffffff1684888a8686604051602401612aec94939291906156f4565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f1806b5f20000000000000000000000000000000000000000000000000000000017905251612b6d91906155ad565b5f6040518083038185875af1925050503d805f8114612ba7576040519150601f19603f3d011682016040523d82523d5f602084013e612bac565b606091505b505080915050612cce565b606f54612bdb9073ffffffffffffffffffffffffffffffffffffffff168686613967565b8473ffffffffffffffffffffffffffffffffffffffff1687898585604051602401612c0994939291906156f4565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f1806b5f20000000000000000000000000000000000000000000000000000000017905251612c8a91906155ad565b5f604051808303815f865af19150503d805f8114612cc3576040519150601f19603f3d011682016040523d82523d5f602084013e612cc8565b606091505b50909150505b80612d05576040517f37e391c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080518c815263ffffffff8a16602082015273ffffffffffffffffffffffffffffffffffffffff89811682840152871660608201526080810186905290517f1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d9181900360a00190a150505050505050505050505050565b5f54610100900460ff1615808015612d9b57505f54600160ff909116105b80612db45750303b158015612db457505f5460ff166001145b612e40576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401611017565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015612e9c575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6040517ff57ac68300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60405180910390a150505050505050565b5f81612eec868686610dd7565b1495945050505050565b60685463ffffffff610100909104811690871603612f40576040517f0595ea2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b6001606860019054906101000a900463ffffffff16338989898888605354604051612f9499989796959493929190615739565b60405180910390a1612fcb6129036001606860019054906101000a900463ffffffff16338a8a8a8989604051611e0592919061559e565b8215611b6757611b676135c0565b73ffffffffffffffffffffffffffffffffffffffff83161580613010575073ffffffffffffffffffffffffffffffffffffffff8216155b15613047576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60685463ffffffff610100909104811690851603613091576040517f658b23ad00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8281165f908152606b6020526040902054640100000000900416156130f7576040517f5eaf7bac00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b1660208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606085901b1660248201525f90603801604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001815282825280516020918201205f818152606a835283812080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8a8116918217909255868601865263ffffffff8c81168089528c8416878a01818152848752606b89528987209a518b54915194167fffffffffffffffff0000000000000000000000000000000000000000000000009091161764010000000093909516929092029390931790975560a285529185902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001689151590811790915585519182529381019590955292840192909252606083015291507fdbe8a5da6a7a916d9adfda9160167a0f8a3da415ee6610e810e753853597fce79060800160405180910390a15050505050565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f06fdde030000000000000000000000000000000000000000000000000000000017905290516060915f91829173ffffffffffffffffffffffffffffffffffffffff86169161334391906155ad565b5f60405180830381855afa9150503d805f811461337b576040519150601f19603f3d011682016040523d82523d5f602084013e613380565b606091505b5091509150816133c5576040518060400160405280600781526020017f4e4f5f4e414d4500000000000000000000000000000000000000000000000000815250610d77565b610d7781614246565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f95d89b410000000000000000000000000000000000000000000000000000000017905290516060915f91829173ffffffffffffffffffffffffffffffffffffffff86169161344f91906155ad565b5f60405180830381855afa9150503d805f8114613487576040519150601f19603f3d011682016040523d82523d5f602084013e61348c565b606091505b5091509150816133c5576040518060400160405280600981526020017f4e4f5f53594d424f4c0000000000000000000000000000000000000000000000815250610d77565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f313ce5670000000000000000000000000000000000000000000000000000000017905290515f918291829173ffffffffffffffffffffffffffffffffffffffff86169161355191906155ad565b5f60405180830381855afa9150503d805f8114613589576040519150601f19603f3d011682016040523d82523d5f602084013e61358e565b606091505b50915091508180156135a1575080516020145b6135ac576012610d77565b80806020019051810190610d7791906157b1565b6053546068805463ffffffff909216790100000000000000000000000000000000000000000000000000027fffffff00000000ffffffffffffffffffffffffffffffffffffffffffffffffff909216919091179081905573ffffffffffffffffffffffffffffffffffffffff65010000000000909104166333d6247d613644610ccd565b6040518263ffffffff1660e01b815260040161366291815260200190565b5f604051808303815f87803b158015613679575f80fd5b505af115801561368b573d5f803e3d5ffd5b50505050565b5f6136a764010000000063ffffffff8416615574565b6136b79063ffffffff851661558b565b600881901c5f8181526069602052604090208054600160ff851690811b91821892839055939450919291908082161561371c576040517f318dafb800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805163ffffffff808a168252881660208201527f4a402ac607e71571d0543be8fcccc358a4df62dcd019a1e7f7e98ca6175e8f2a9101612ece565b5f8060405180611ba00160405280611b668152602001615b09611b66913983604051602001613789929190615546565b6040516020818303038152906040529050838151602083015ff5915073ffffffffffffffffffffffffffffffffffffffff82166137f2576040517fbefb092000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5092915050565b5f54610100900460ff1661388f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401611017565b610db6614416565b73ffffffffffffffffffffffffffffffffffffffff82165f90815260a2602052604090205460ff16156138ea576138e673ffffffffffffffffffffffffffffffffffffffff8316333084614100565b5050565b6040517f9dc29fac0000000000000000000000000000000000000000000000000000000081523360048201526024810182905273ffffffffffffffffffffffffffffffffffffffff831690639dc29fac906044015f604051808303815f87803b158015613955575f80fd5b505af1158015611b67573d5f803e3d5ffd5b73ffffffffffffffffffffffffffffffffffffffff83165f90815260a2602052604090205460ff16156139b557610f5373ffffffffffffffffffffffffffffffffffffffff84168383613bd4565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8381166004830152602482018390528416906340c10f19906044015f604051808303815f87803b158015613a22575f80fd5b505af1158015612923573d5f803e3d5ffd5b6068545f9065010000000000900473ffffffffffffffffffffffffffffffffffffffff1663257b3632613a6786866144ac565b6040518263ffffffff1660e01b8152600401613a8591815260200190565b6020604051808303815f875af1158015613aa1573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613ac59190615654565b9050805f03613aff576040517e2f6fad00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8068010000000000000000871615613b5c57869150613b21848a8489612edf565b613b57576040517fe0417cec00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613bbf565b602087901c613b6c8160016157cc565b9150879250613b87613b7f868c86610dd7565b8a8389612edf565b613bbd576040517fe0417cec00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b613bc982826144c1565b505050505050505050565b60405173ffffffffffffffffffffffffffffffffffffffff8316602482015260448101829052610f539084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261454d565b600260015403613d14576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401611017565b6002600155565b5f613d2960048284866157e9565b613d3291615810565b90507f2afa5331000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000821601613eeb575f808080808080613d91896004818d6157e9565b810190613d9e9190615858565b96509650965096509650965096508a8514613de5576040517f03fffc4b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805173ffffffffffffffffffffffffffffffffffffffff89811660248301528881166044830152606482018890526084820187905260ff861660a483015260c4820185905260e48083018590528351808403909101815261010490920183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd505accf000000000000000000000000000000000000000000000000000000001790529151918e1691613e9e91906155ad565b5f604051808303815f865af19150503d805f8114613ed7576040519150601f19603f3d011682016040523d82523d5f602084013e613edc565b606091505b50505050505050505050610b05565b7fffffffff0000000000000000000000000000000000000000000000000000000081167f8fcbaf0c0000000000000000000000000000000000000000000000000000000014613f66576040517fe282c0ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80808080808080613f7b8a6004818e6157e9565b810190613f8891906158a7565b975097509750975097509750975097508c73ffffffffffffffffffffffffffffffffffffffff16638fcbaf0c60e01b898989898989898960405160240161402798979695949392919073ffffffffffffffffffffffffffffffffffffffff9889168152969097166020870152604086019490945260608501929092521515608084015260ff1660a083015260c082015260e08101919091526101000190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009094169390931790925290516140b091906155ad565b5f604051808303815f865af19150503d805f81146140e9576040519150601f19603f3d011682016040523d82523d5f602084013e6140ee565b606091505b50505050505050505050505050505050565b60405173ffffffffffffffffffffffffffffffffffffffff8085166024830152831660448201526064810182905261368b9085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401613c26565b80600161416d60206002615a43565b614177919061566b565b605354106141b1576040517fef5ccf6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60535f81546141c09061537a565b918290555090505f5b6020811015614237578082901c6001166001036141fc5782603382602081106141f4576141f4615320565b015550505050565b6142236033826020811061421257614212615320565b0154845f9182526020526040902090565b92508061422f8161537a565b9150506141c9565b50610f53615a4e565b60018055565b606060408251106142655781806020019051810190610a579190615a7b565b81516020036143d8575f5b6020811080156142b7575082818151811061428d5761428d615320565b01602001517fff000000000000000000000000000000000000000000000000000000000000001615155b156142ce57806142c68161537a565b915050614270565b805f0361431057505060408051808201909152601281527f4e4f545f56414c49445f454e434f44494e4700000000000000000000000000006020820152919050565b5f8167ffffffffffffffff81111561432a5761432a614a16565b6040519080825280601f01601f191660200182016040528015614354576020820181803683370190505b5090505f5b828110156143d05784818151811061437357614373615320565b602001015160f81c60f81b82828151811061439057614390615320565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350806143c88161537a565b915050614359565b509392505050565b505060408051808201909152601281527f4e4f545f56414c49445f454e434f44494e470000000000000000000000000000602082015290565b919050565b5f54610100900460ff16614240576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401611017565b5f8281526020829052604081205b9392505050565b5f6144d764010000000063ffffffff8416615574565b6144e79063ffffffff851661558b565b600881901c5f8181526069602052604081208054600160ff861690811b91821892839055949550929392918183169003612923576040517f646cf55800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6145ae826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166146589092919063ffffffff16565b805190915015610f5357808060200190518101906145cc9190615aed565b610f53576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401611017565b6060610d7784845f85855f808673ffffffffffffffffffffffffffffffffffffffff16858760405161468a91906155ad565b5f6040518083038185875af1925050503d805f81146146c4576040519150601f19603f3d011682016040523d82523d5f602084013e6146c9565b606091505b50915091506146da878383876146e5565b979650505050505050565b6060831561477a5782515f036147735773ffffffffffffffffffffffffffffffffffffffff85163b614773576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401611017565b5081610d77565b610d77838381511561478f5781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110179190614961565b803563ffffffff81168114614411575f80fd5b73ffffffffffffffffffffffffffffffffffffffff811681146147f7575f80fd5b50565b5f806040838503121561480b575f80fd5b614814836147c3565b91506020830135614824816147d6565b809150509250929050565b80151581146147f7575f80fd5b5f8083601f84011261484c575f80fd5b50813567ffffffffffffffff811115614863575f80fd5b60208301915083602082850101111561487a575f80fd5b9250929050565b5f805f805f60808688031215614895575f80fd5b61489e866147c3565b945060208601356148ae816147d6565b935060408601356148be8161482f565b9250606086013567ffffffffffffffff8111156148d9575f80fd5b6148e58882890161483c565b969995985093965092949392505050565b5f5b838110156149105781810151838201526020016148f8565b50505f910152565b5f815180845261492f8160208601602086016148f6565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081525f6144ba6020830184614918565b5f60208284031215614983575f80fd5b81356144ba816147d6565b60ff811681146147f7575f80fd5b5f805f805f805f60e0888a0312156149b2575f80fd5b87356149bd8161498e565b96506149cb602089016147c3565b955060408801356149db816147d6565b94506149e9606089016147c3565b935060808801356149f9816147d6565b9699959850939692959460a0840135945060c09093013592915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715614a8a57614a8a614a16565b604052919050565b5f67ffffffffffffffff821115614aab57614aab614a16565b5060051b60200190565b5f82601f830112614ac4575f80fd5b81356020614ad9614ad483614a92565b614a43565b82815260059290921b84018101918181019086841115614af7575f80fd5b8286015b84811015614b1957614b0c816147c3565b8352918301918301614afb565b509695505050505050565b5f82601f830112614b33575f80fd5b81356020614b43614ad483614a92565b82815260059290921b84018101918181019086841115614b61575f80fd5b8286015b84811015614b19578035614b78816147d6565b8352918301918301614b65565b5f805f8060808587031215614b98575f80fd5b843567ffffffffffffffff80821115614baf575f80fd5b614bbb88838901614ab5565b9550602091508187013581811115614bd1575f80fd5b614bdd89828a01614b24565b955050604087013581811115614bf1575f80fd5b614bfd89828a01614b24565b945050606087013581811115614c11575f80fd5b87019050601f81018813614c23575f80fd5b8035614c31614ad482614a92565b81815260059190911b8201830190838101908a831115614c4f575f80fd5b928401925b82841015614c76578335614c678161482f565b82529284019290840190614c54565b979a9699509497505050505050565b5f805f60608486031215614c97575f80fd5b614ca0846147c3565b92506020840135614cb0816147d6565b91506040840135614cc0816147d6565b809150509250925092565b5f60208284031215614cdb575f80fd5b5035919050565b806104008101831015610a57575f80fd5b5f805f6104408486031215614d06575f80fd5b83359250614d178560208601614ce2565b9150614d2661042085016147c3565b90509250925092565b5f8060408385031215614d40575f80fd5b823567ffffffffffffffff80821115614d57575f80fd5b614d6386838701614ab5565b93506020850135915080821115614d78575f80fd5b50614d8585828601614ab5565b9150509250929050565b5f67ffffffffffffffff821115614da857614da8614a16565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b5f82601f830112614de3575f80fd5b8135614df1614ad482614d8f565b818152846020838601011115614e05575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f805f805f805f6101208a8c031215614e3a575f80fd5b614e438a6147c3565b985060208a0135614e53816147d6565b9750614e6160408b016147c3565b965060608a0135614e71816147d6565b955060808a0135614e81816147d6565b945060a08a013567ffffffffffffffff811115614e9c575f80fd5b614ea88c828d01614dd4565b94505060c08a0135614eb9816147d6565b925060e08a0135614ec9816147d6565b91506101008a0135614eda8161482f565b809150509295985092959850929598565b5f8060408385031215614efc575f80fd5b8235614f07816147d6565b946020939093013593505050565b5f805f805f60a08688031215614f29575f80fd5b614f32866147c3565b94506020860135614f42816147d6565b9350604086013567ffffffffffffffff80821115614f5e575f80fd5b614f6a89838a01614dd4565b94506060880135915080821115614f7f575f80fd5b50614f8c88828901614dd4565b9250506080860135614f9d8161498e565b809150509295509295909350565b5f805f805f8060a08789031215614fc0575f80fd5b614fc9876147c3565b95506020870135614fd9816147d6565b9450604087013593506060870135614ff08161482f565b9250608087013567ffffffffffffffff81111561500b575f80fd5b61501789828a0161483c565b979a9699509497509295939492505050565b5f806040838503121561503a575f80fd5b8235615045816147d6565b915060208301356148248161482f565b5f8060408385031215615066575f80fd5b61506f836147c3565b915061507d602084016147c3565b90509250929050565b5f805f805f805f805f805f806109208d8f0312156150a2575f80fd5b6150ac8e8e614ce2565b9b506150bc8e6104008f01614ce2565b9a506108008d013599506108208d013598506108408d013597506150e36108608e016147c3565b96506150f36108808e01356147d6565b6108808d013595506151086108a08e016147c3565b94506151186108c08e01356147d6565b6108c08d013593506108e08d0135925067ffffffffffffffff6109008e01351115615141575f80fd5b6151528e6109008f01358f0161483c565b81935080925050509295989b509295989b509295989b565b5f805f805f805f60c0888a031215615180575f80fd5b615189886147c3565b96506020880135615199816147d6565b95506040880135945060608801356151b0816147d6565b935060808801356151c08161482f565b925060a088013567ffffffffffffffff8111156151db575f80fd5b6151e78a828b0161483c565b989b979a50959850939692959293505050565b5f805f805f8060c0878903121561520f575f80fd5b615218876147c3565b95506020870135615228816147d6565b9450615236604088016147c3565b93506060870135615246816147d6565b92506080870135615256816147d6565b915060a087013567ffffffffffffffff811115615271575f80fd5b61527d89828a01614dd4565b9150509295509295509295565b5f805f80610460858703121561529e575f80fd5b843593506152af8660208701614ce2565b92506152be61042086016147c3565b939692955092936104400135925050565b600181811c908216806152e357607f821691505b60208210810361531a577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036153aa576153aa61534d565b5060010190565b601f821115610f53575f81815260208120601f850160051c810160208610156153d75750805b601f850160051c820191505b81811015611b67578281556001016153e3565b815167ffffffffffffffff81111561541057615410614a16565b6154248161541e84546152cf565b846153b1565b602080601f831160018114615476575f84156154405750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555611b67565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b828110156154c2578886015182559484019460019091019084016154a3565b50858210156154fe57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b606081525f6155206060830186614918565b82810360208401526155328186614918565b91505060ff83166040830152949350505050565b5f83516155578184602088016148f6565b83519083019061556b8183602088016148f6565b01949350505050565b8082028115828204841417610a5757610a5761534d565b80820180821115610a5757610a5761534d565b818382375f9101908152919050565b5f82516155be8184602087016148f6565b9190910192915050565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b63ffffffff861681525f73ffffffffffffffffffffffffffffffffffffffff8087166020840152808616604084015250608060608301526146da6080830184866155c8565b5f60208284031215615664575f80fd5b5051919050565b81810381811115610a5757610a5761534d565b5f61010060ff8b16835263ffffffff808b16602085015273ffffffffffffffffffffffffffffffffffffffff808b166040860152818a1660608601528089166080860152508660a08501528160c08501526156db82850187614918565b925080851660e085015250509998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8516815263ffffffff84166020820152606060408201525f61572f6060830184866155c8565b9695505050505050565b5f61010060ff8c16835263ffffffff808c16602085015273ffffffffffffffffffffffffffffffffffffffff808c166040860152818b166060860152808a166080860152508760a08501528160c085015261579782850187896155c8565b925080851660e085015250509a9950505050505050505050565b5f602082840312156157c1575f80fd5b81516144ba8161498e565b63ffffffff8181168382160190808211156137f2576137f261534d565b5f80858511156157f7575f80fd5b83861115615803575f80fd5b5050820193919092039150565b7fffffffff0000000000000000000000000000000000000000000000000000000081358181169160048510156158505780818660040360031b1b83161692505b505092915050565b5f805f805f805f60e0888a03121561586e575f80fd5b8735615879816147d6565b96506020880135615889816147d6565b9550604088013594506060880135935060808801356149f98161498e565b5f805f805f805f80610100898b0312156158bf575f80fd5b88356158ca816147d6565b975060208901356158da816147d6565b9650604089013595506060890135945060808901356158f88161482f565b935060a08901356159088161498e565b979a969950949793969295929450505060c08201359160e0013590565b600181815b8085111561597e57817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048211156159645761596461534d565b8085161561597157918102915b93841c939080029061592a565b509250929050565b5f8261599457506001610a57565b816159a057505f610a57565b81600181146159b657600281146159c0576159dc565b6001915050610a57565b60ff8411156159d1576159d161534d565b50506001821b610a57565b5060208310610133831016604e8410600b84101617156159ff575081810a610a57565b615a098383615925565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115615a3b57615a3b61534d565b029392505050565b5f6144ba8383615986565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52600160045260245ffd5b5f60208284031215615a8b575f80fd5b815167ffffffffffffffff811115615aa1575f80fd5b8201601f81018413615ab1575f80fd5b8051615abf614ad482614d8f565b818152856020838501011115615ad3575f80fd5b615ae48260208301602086016148f6565b95945050505050565b5f60208284031215615afd575f80fd5b81516144ba8161482f56fe6101006040523480156200001257600080fd5b5060405162001b6638038062001b6683398101604081905262000035916200028d565b82826003620000458382620003a1565b506004620000548282620003a1565b50503360c0525060ff811660e052466080819052620000739062000080565b60a052506200046d915050565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f620000ad6200012e565b805160209182012060408051808201825260018152603160f81b90840152805192830193909352918101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc66060820152608081018390523060a082015260c001604051602081830303815290604052805190602001209050919050565b6060600380546200013f9062000312565b80601f01602080910402602001604051908101604052809291908181526020018280546200016d9062000312565b8015620001be5780601f106200019257610100808354040283529160200191620001be565b820191906000526020600020905b815481529060010190602001808311620001a057829003601f168201915b5050505050905090565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620001f057600080fd5b81516001600160401b03808211156200020d576200020d620001c8565b604051601f8301601f19908116603f01168101908282118183101715620002385762000238620001c8565b816040528381526020925086838588010111156200025557600080fd5b600091505b838210156200027957858201830151818301840152908201906200025a565b600093810190920192909252949350505050565b600080600060608486031215620002a357600080fd5b83516001600160401b0380821115620002bb57600080fd5b620002c987838801620001de565b94506020860151915080821115620002e057600080fd5b50620002ef86828701620001de565b925050604084015160ff811681146200030757600080fd5b809150509250925092565b600181811c908216806200032757607f821691505b6020821081036200034857634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200039c57600081815260208120601f850160051c81016020861015620003775750805b601f850160051c820191505b81811015620003985782815560010162000383565b5050505b505050565b81516001600160401b03811115620003bd57620003bd620001c8565b620003d581620003ce845462000312565b846200034e565b602080601f8311600181146200040d5760008415620003f45750858301515b600019600386901b1c1916600185901b17855562000398565b600085815260208120601f198616915b828110156200043e578886015182559484019460019091019084016200041d565b50858210156200045d5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a05160c05160e0516116aa620004bc6000396000610237015260008181610307015281816105c001526106a70152600061053a015260008181610379015261050401526116aa6000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c806370a08231116100d8578063a457c2d71161008c578063d505accf11610066578063d505accf1461039b578063dd62ed3e146103ae578063ffa1ad74146103f457600080fd5b8063a457c2d71461034e578063a9059cbb14610361578063cd0d00961461037457600080fd5b806395d89b41116100bd57806395d89b41146102e75780639dc29fac146102ef578063a3c573eb1461030257600080fd5b806370a08231146102915780637ecebe00146102c757600080fd5b806330adf81f1161012f5780633644e515116101145780633644e51514610261578063395093511461026957806340c10f191461027c57600080fd5b806330adf81f14610209578063313ce5671461023057600080fd5b806318160ddd1161016057806318160ddd146101bd57806320606b70146101cf57806323b872dd146101f657600080fd5b806306fdde031461017c578063095ea7b31461019a575b600080fd5b610184610430565b60405161019191906113e4565b60405180910390f35b6101ad6101a8366004611479565b6104c2565b6040519015158152602001610191565b6002545b604051908152602001610191565b6101c17f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f81565b6101ad6102043660046114a3565b6104dc565b6101c17f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c981565b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610191565b6101c1610500565b6101ad610277366004611479565b61055c565b61028f61028a366004611479565b6105a8565b005b6101c161029f3660046114df565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6101c16102d53660046114df565b60056020526000908152604090205481565b610184610680565b61028f6102fd366004611479565b61068f565b6103297f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610191565b6101ad61035c366004611479565b61075e565b6101ad61036f366004611479565b61082f565b6101c17f000000000000000000000000000000000000000000000000000000000000000081565b61028f6103a9366004611501565b61083d565b6101c16103bc366004611574565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6101846040518060400160405280600181526020017f310000000000000000000000000000000000000000000000000000000000000081525081565b60606003805461043f906115a7565b80601f016020809104026020016040519081016040528092919081815260200182805461046b906115a7565b80156104b85780601f1061048d576101008083540402835291602001916104b8565b820191906000526020600020905b81548152906001019060200180831161049b57829003601f168201915b5050505050905090565b6000336104d0818585610b73565b60019150505b92915050565b6000336104ea858285610d27565b6104f5858585610dfe565b506001949350505050565b60007f00000000000000000000000000000000000000000000000000000000000000004614610537576105324661106d565b905090565b507f000000000000000000000000000000000000000000000000000000000000000090565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091906104d090829086906105a3908790611629565b610b73565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610672576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f546f6b656e577261707065643a3a6f6e6c794272696467653a204e6f7420506f60448201527f6c79676f6e5a6b45564d4272696467650000000000000000000000000000000060648201526084015b60405180910390fd5b61067c8282611135565b5050565b60606004805461043f906115a7565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610754576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f546f6b656e577261707065643a3a6f6e6c794272696467653a204e6f7420506f60448201527f6c79676f6e5a6b45564d427269646765000000000000000000000000000000006064820152608401610669565b61067c8282611228565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919083811015610822576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f0000000000000000000000000000000000000000000000000000006064820152608401610669565b6104f58286868403610b73565b6000336104d0818585610dfe565b834211156108cc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f546f6b656e577261707065643a3a7065726d69743a204578706972656420706560448201527f726d6974000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff8716600090815260056020526040812080547f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9918a918a918a9190866109268361163c565b9091555060408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090506000610991610500565b6040517f19010000000000000000000000000000000000000000000000000000000000006020820152602281019190915260428101839052606201604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600080855291840180845281905260ff89169284019290925260608301879052608083018690529092509060019060a0016020604051602081039080840390855afa158015610a55573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811615801590610ad057508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b610b5c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f546f6b656e577261707065643a3a7065726d69743a20496e76616c696420736960448201527f676e6174757265000000000000000000000000000000000000000000000000006064820152608401610669565b610b678a8a8a610b73565b50505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8316610c15576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff8216610cb8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610df85781811015610deb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610669565b610df88484848403610b73565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610ea1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff8216610f44576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610ffa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3610df8565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f611098610430565b8051602091820120604080518082018252600181527f310000000000000000000000000000000000000000000000000000000000000090840152805192830193909352918101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc66060820152608081018390523060a082015260c001604051602081830303815290604052805190602001209050919050565b73ffffffffffffffffffffffffffffffffffffffff82166111b2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610669565b80600260008282546111c49190611629565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b73ffffffffffffffffffffffffffffffffffffffff82166112cb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f73000000000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090205481811015611381576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f63650000000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff83166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101610d1a565b600060208083528351808285015260005b81811015611411578581018301518582016040015282016113f5565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461147457600080fd5b919050565b6000806040838503121561148c57600080fd5b61149583611450565b946020939093013593505050565b6000806000606084860312156114b857600080fd5b6114c184611450565b92506114cf60208501611450565b9150604084013590509250925092565b6000602082840312156114f157600080fd5b6114fa82611450565b9392505050565b600080600080600080600060e0888a03121561151c57600080fd5b61152588611450565b965061153360208901611450565b95506040880135945060608801359350608088013560ff8116811461155757600080fd5b9699959850939692959460a0840135945060c09093013592915050565b6000806040838503121561158757600080fd5b61159083611450565b915061159e60208401611450565b90509250929050565b600181811c908216806115bb57607f821691505b6020821081036115f4577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156104d6576104d66115fa565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361166d5761166d6115fa565b506001019056fea26469706673582212208d88fee561cff7120d381c345cfc534cef8229a272dc5809d4bbb685ad67141164736f6c63430008110033a2646970667358221220b0f948a92f8b838ff3e863251878781d63e3d245e9aaea0b865ed81e65bb903b64736f6c63430008140033",
}

// Bridgel2sovereignchainpessimisticABI is the input ABI used to generate the binding from.
// Deprecated: Use Bridgel2sovereignchainpessimisticMetaData.ABI instead.
var Bridgel2sovereignchainpessimisticABI = Bridgel2sovereignchainpessimisticMetaData.ABI

// Bridgel2sovereignchainpessimisticBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Bridgel2sovereignchainpessimisticMetaData.Bin instead.
var Bridgel2sovereignchainpessimisticBin = Bridgel2sovereignchainpessimisticMetaData.Bin

// DeployBridgel2sovereignchainpessimistic deploys a new Ethereum contract, binding an instance of Bridgel2sovereignchainpessimistic to it.
func DeployBridgel2sovereignchainpessimistic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bridgel2sovereignchainpessimistic, error) {
	parsed, err := Bridgel2sovereignchainpessimisticMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Bridgel2sovereignchainpessimisticBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridgel2sovereignchainpessimistic{Bridgel2sovereignchainpessimisticCaller: Bridgel2sovereignchainpessimisticCaller{contract: contract}, Bridgel2sovereignchainpessimisticTransactor: Bridgel2sovereignchainpessimisticTransactor{contract: contract}, Bridgel2sovereignchainpessimisticFilterer: Bridgel2sovereignchainpessimisticFilterer{contract: contract}}, nil
}

// Bridgel2sovereignchainpessimistic is an auto generated Go binding around an Ethereum contract.
type Bridgel2sovereignchainpessimistic struct {
	Bridgel2sovereignchainpessimisticCaller     // Read-only binding to the contract
	Bridgel2sovereignchainpessimisticTransactor // Write-only binding to the contract
	Bridgel2sovereignchainpessimisticFilterer   // Log filterer for contract events
}

// Bridgel2sovereignchainpessimisticCaller is an auto generated read-only Go binding around an Ethereum contract.
type Bridgel2sovereignchainpessimisticCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bridgel2sovereignchainpessimisticTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Bridgel2sovereignchainpessimisticTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bridgel2sovereignchainpessimisticFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Bridgel2sovereignchainpessimisticFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bridgel2sovereignchainpessimisticSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Bridgel2sovereignchainpessimisticSession struct {
	Contract     *Bridgel2sovereignchainpessimistic // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                      // Call options to use throughout this session
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// Bridgel2sovereignchainpessimisticCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Bridgel2sovereignchainpessimisticCallerSession struct {
	Contract *Bridgel2sovereignchainpessimisticCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                            // Call options to use throughout this session
}

// Bridgel2sovereignchainpessimisticTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Bridgel2sovereignchainpessimisticTransactorSession struct {
	Contract     *Bridgel2sovereignchainpessimisticTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                            // Transaction auth options to use throughout this session
}

// Bridgel2sovereignchainpessimisticRaw is an auto generated low-level Go binding around an Ethereum contract.
type Bridgel2sovereignchainpessimisticRaw struct {
	Contract *Bridgel2sovereignchainpessimistic // Generic contract binding to access the raw methods on
}

// Bridgel2sovereignchainpessimisticCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Bridgel2sovereignchainpessimisticCallerRaw struct {
	Contract *Bridgel2sovereignchainpessimisticCaller // Generic read-only contract binding to access the raw methods on
}

// Bridgel2sovereignchainpessimisticTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Bridgel2sovereignchainpessimisticTransactorRaw struct {
	Contract *Bridgel2sovereignchainpessimisticTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgel2sovereignchainpessimistic creates a new instance of Bridgel2sovereignchainpessimistic, bound to a specific deployed contract.
func NewBridgel2sovereignchainpessimistic(address common.Address, backend bind.ContractBackend) (*Bridgel2sovereignchainpessimistic, error) {
	contract, err := bindBridgel2sovereignchainpessimistic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainpessimistic{Bridgel2sovereignchainpessimisticCaller: Bridgel2sovereignchainpessimisticCaller{contract: contract}, Bridgel2sovereignchainpessimisticTransactor: Bridgel2sovereignchainpessimisticTransactor{contract: contract}, Bridgel2sovereignchainpessimisticFilterer: Bridgel2sovereignchainpessimisticFilterer{contract: contract}}, nil
}

// NewBridgel2sovereignchainpessimisticCaller creates a new read-only instance of Bridgel2sovereignchainpessimistic, bound to a specific deployed contract.
func NewBridgel2sovereignchainpessimisticCaller(address common.Address, caller bind.ContractCaller) (*Bridgel2sovereignchainpessimisticCaller, error) {
	contract, err := bindBridgel2sovereignchainpessimistic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainpessimisticCaller{contract: contract}, nil
}

// NewBridgel2sovereignchainpessimisticTransactor creates a new write-only instance of Bridgel2sovereignchainpessimistic, bound to a specific deployed contract.
func NewBridgel2sovereignchainpessimisticTransactor(address common.Address, transactor bind.ContractTransactor) (*Bridgel2sovereignchainpessimisticTransactor, error) {
	contract, err := bindBridgel2sovereignchainpessimistic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainpessimisticTransactor{contract: contract}, nil
}

// NewBridgel2sovereignchainpessimisticFilterer creates a new log filterer instance of Bridgel2sovereignchainpessimistic, bound to a specific deployed contract.
func NewBridgel2sovereignchainpessimisticFilterer(address common.Address, filterer bind.ContractFilterer) (*Bridgel2sovereignchainpessimisticFilterer, error) {
	contract, err := bindBridgel2sovereignchainpessimistic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainpessimisticFilterer{contract: contract}, nil
}

// bindBridgel2sovereignchainpessimistic binds a generic wrapper to an already deployed contract.
func bindBridgel2sovereignchainpessimistic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Bridgel2sovereignchainpessimisticMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridgel2sovereignchainpessimistic.Contract.Bridgel2sovereignchainpessimisticCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.Bridgel2sovereignchainpessimisticTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.Bridgel2sovereignchainpessimisticTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridgel2sovereignchainpessimistic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.contract.Transact(opts, method, params...)
}

// BASEINITBYTECODEWRAPPEDTOKEN is a free data retrieval call binding the contract method 0x83c43a55.
//
// Solidity: function BASE_INIT_BYTECODE_WRAPPED_TOKEN() view returns(bytes)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) BASEINITBYTECODEWRAPPEDTOKEN(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "BASE_INIT_BYTECODE_WRAPPED_TOKEN")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// BASEINITBYTECODEWRAPPEDTOKEN is a free data retrieval call binding the contract method 0x83c43a55.
//
// Solidity: function BASE_INIT_BYTECODE_WRAPPED_TOKEN() view returns(bytes)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) BASEINITBYTECODEWRAPPEDTOKEN() ([]byte, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.BASEINITBYTECODEWRAPPEDTOKEN(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// BASEINITBYTECODEWRAPPEDTOKEN is a free data retrieval call binding the contract method 0x83c43a55.
//
// Solidity: function BASE_INIT_BYTECODE_WRAPPED_TOKEN() view returns(bytes)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) BASEINITBYTECODEWRAPPEDTOKEN() ([]byte, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.BASEINITBYTECODEWRAPPEDTOKEN(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) WETHToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "WETHToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) WETHToken() (common.Address, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.WETHToken(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) WETHToken() (common.Address, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.WETHToken(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// ActivateEmergencyState is a free data retrieval call binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() pure returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) ActivateEmergencyState(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "activateEmergencyState")

	if err != nil {
		return err
	}

	return err

}

// ActivateEmergencyState is a free data retrieval call binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() pure returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) ActivateEmergencyState() error {
	return _Bridgel2sovereignchainpessimistic.Contract.ActivateEmergencyState(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// ActivateEmergencyState is a free data retrieval call binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() pure returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) ActivateEmergencyState() error {
	return _Bridgel2sovereignchainpessimistic.Contract.ActivateEmergencyState(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// BridgeManager is a free data retrieval call binding the contract method 0x14cc01a0.
//
// Solidity: function bridgeManager() view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) BridgeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "bridgeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeManager is a free data retrieval call binding the contract method 0x14cc01a0.
//
// Solidity: function bridgeManager() view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) BridgeManager() (common.Address, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.BridgeManager(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// BridgeManager is a free data retrieval call binding the contract method 0x14cc01a0.
//
// Solidity: function bridgeManager() view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) BridgeManager() (common.Address, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.BridgeManager(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) CalculateRoot(opts *bind.CallOpts, leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "calculateRoot", leafHash, smtProof, index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) CalculateRoot(leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.CalculateRoot(&_Bridgel2sovereignchainpessimistic.CallOpts, leafHash, smtProof, index)
}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) CalculateRoot(leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.CalculateRoot(&_Bridgel2sovereignchainpessimistic.CallOpts, leafHash, smtProof, index)
}

// CalculateTokenWrapperAddress is a free data retrieval call binding the contract method 0x7843298b.
//
// Solidity: function calculateTokenWrapperAddress(uint32 originNetwork, address originTokenAddress, address token) view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) CalculateTokenWrapperAddress(opts *bind.CallOpts, originNetwork uint32, originTokenAddress common.Address, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "calculateTokenWrapperAddress", originNetwork, originTokenAddress, token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CalculateTokenWrapperAddress is a free data retrieval call binding the contract method 0x7843298b.
//
// Solidity: function calculateTokenWrapperAddress(uint32 originNetwork, address originTokenAddress, address token) view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) CalculateTokenWrapperAddress(originNetwork uint32, originTokenAddress common.Address, token common.Address) (common.Address, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.CalculateTokenWrapperAddress(&_Bridgel2sovereignchainpessimistic.CallOpts, originNetwork, originTokenAddress, token)
}

// CalculateTokenWrapperAddress is a free data retrieval call binding the contract method 0x7843298b.
//
// Solidity: function calculateTokenWrapperAddress(uint32 originNetwork, address originTokenAddress, address token) view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) CalculateTokenWrapperAddress(originNetwork uint32, originTokenAddress common.Address, token common.Address) (common.Address, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.CalculateTokenWrapperAddress(&_Bridgel2sovereignchainpessimistic.CallOpts, originNetwork, originTokenAddress, token)
}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) ClaimedBitMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "claimedBitMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) ClaimedBitMap(arg0 *big.Int) (*big.Int, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.ClaimedBitMap(&_Bridgel2sovereignchainpessimistic.CallOpts, arg0)
}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) ClaimedBitMap(arg0 *big.Int) (*big.Int, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.ClaimedBitMap(&_Bridgel2sovereignchainpessimistic.CallOpts, arg0)
}

// DeactivateEmergencyState is a free data retrieval call binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() pure returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) DeactivateEmergencyState(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "deactivateEmergencyState")

	if err != nil {
		return err
	}

	return err

}

// DeactivateEmergencyState is a free data retrieval call binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() pure returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) DeactivateEmergencyState() error {
	return _Bridgel2sovereignchainpessimistic.Contract.DeactivateEmergencyState(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// DeactivateEmergencyState is a free data retrieval call binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() pure returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) DeactivateEmergencyState() error {
	return _Bridgel2sovereignchainpessimistic.Contract.DeactivateEmergencyState(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) DepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "depositCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) DepositCount() (*big.Int, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.DepositCount(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) DepositCount() (*big.Int, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.DepositCount(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) GasTokenAddress() (common.Address, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.GasTokenAddress(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) GasTokenAddress() (common.Address, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.GasTokenAddress(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// GasTokenMetadata is a free data retrieval call binding the contract method 0x27aef4e8.
//
// Solidity: function gasTokenMetadata() view returns(bytes)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) GasTokenMetadata(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "gasTokenMetadata")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GasTokenMetadata is a free data retrieval call binding the contract method 0x27aef4e8.
//
// Solidity: function gasTokenMetadata() view returns(bytes)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) GasTokenMetadata() ([]byte, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.GasTokenMetadata(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// GasTokenMetadata is a free data retrieval call binding the contract method 0x27aef4e8.
//
// Solidity: function gasTokenMetadata() view returns(bytes)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) GasTokenMetadata() ([]byte, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.GasTokenMetadata(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) GasTokenNetwork() (uint32, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.GasTokenNetwork(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) GasTokenNetwork() (uint32, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.GasTokenNetwork(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// GetLeafValue is a free data retrieval call binding the contract method 0x3e197043.
//
// Solidity: function getLeafValue(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes32 metadataHash) pure returns(bytes32)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) GetLeafValue(opts *bind.CallOpts, leafType uint8, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadataHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "getLeafValue", leafType, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadataHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLeafValue is a free data retrieval call binding the contract method 0x3e197043.
//
// Solidity: function getLeafValue(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes32 metadataHash) pure returns(bytes32)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) GetLeafValue(leafType uint8, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadataHash [32]byte) ([32]byte, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.GetLeafValue(&_Bridgel2sovereignchainpessimistic.CallOpts, leafType, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadataHash)
}

// GetLeafValue is a free data retrieval call binding the contract method 0x3e197043.
//
// Solidity: function getLeafValue(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes32 metadataHash) pure returns(bytes32)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) GetLeafValue(leafType uint8, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadataHash [32]byte) ([32]byte, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.GetLeafValue(&_Bridgel2sovereignchainpessimistic.CallOpts, leafType, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadataHash)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) GetRoot() ([32]byte, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.GetRoot(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) GetRoot() ([32]byte, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.GetRoot(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) GetTokenMetadata(opts *bind.CallOpts, token common.Address) ([]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "getTokenMetadata", token)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) GetTokenMetadata(token common.Address) ([]byte, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.GetTokenMetadata(&_Bridgel2sovereignchainpessimistic.CallOpts, token)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) GetTokenMetadata(token common.Address) ([]byte, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.GetTokenMetadata(&_Bridgel2sovereignchainpessimistic.CallOpts, token)
}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) GetTokenWrappedAddress(opts *bind.CallOpts, originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "getTokenWrappedAddress", originNetwork, originTokenAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) GetTokenWrappedAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.GetTokenWrappedAddress(&_Bridgel2sovereignchainpessimistic.CallOpts, originNetwork, originTokenAddress)
}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) GetTokenWrappedAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.GetTokenWrappedAddress(&_Bridgel2sovereignchainpessimistic.CallOpts, originNetwork, originTokenAddress)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) GlobalExitRootManager() (common.Address, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.GlobalExitRootManager(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.GlobalExitRootManager(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// IsClaimed is a free data retrieval call binding the contract method 0xcc461632.
//
// Solidity: function isClaimed(uint32 leafIndex, uint32 sourceBridgeNetwork) view returns(bool)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) IsClaimed(opts *bind.CallOpts, leafIndex uint32, sourceBridgeNetwork uint32) (bool, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "isClaimed", leafIndex, sourceBridgeNetwork)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClaimed is a free data retrieval call binding the contract method 0xcc461632.
//
// Solidity: function isClaimed(uint32 leafIndex, uint32 sourceBridgeNetwork) view returns(bool)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) IsClaimed(leafIndex uint32, sourceBridgeNetwork uint32) (bool, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.IsClaimed(&_Bridgel2sovereignchainpessimistic.CallOpts, leafIndex, sourceBridgeNetwork)
}

// IsClaimed is a free data retrieval call binding the contract method 0xcc461632.
//
// Solidity: function isClaimed(uint32 leafIndex, uint32 sourceBridgeNetwork) view returns(bool)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) IsClaimed(leafIndex uint32, sourceBridgeNetwork uint32) (bool, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.IsClaimed(&_Bridgel2sovereignchainpessimistic.CallOpts, leafIndex, sourceBridgeNetwork)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) IsEmergencyState() (bool, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.IsEmergencyState(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) IsEmergencyState() (bool, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.IsEmergencyState(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) LastUpdatedDepositCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "lastUpdatedDepositCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) LastUpdatedDepositCount() (uint32, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.LastUpdatedDepositCount(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) LastUpdatedDepositCount() (uint32, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.LastUpdatedDepositCount(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) NetworkID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "networkID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) NetworkID() (uint32, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.NetworkID(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) NetworkID() (uint32, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.NetworkID(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// PolygonRollupManager is a free data retrieval call binding the contract method 0x8ed7e3f2.
//
// Solidity: function polygonRollupManager() view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) PolygonRollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "polygonRollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PolygonRollupManager is a free data retrieval call binding the contract method 0x8ed7e3f2.
//
// Solidity: function polygonRollupManager() view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) PolygonRollupManager() (common.Address, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.PolygonRollupManager(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// PolygonRollupManager is a free data retrieval call binding the contract method 0x8ed7e3f2.
//
// Solidity: function polygonRollupManager() view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) PolygonRollupManager() (common.Address, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.PolygonRollupManager(&_Bridgel2sovereignchainpessimistic.CallOpts)
}

// PrecalculatedWrapperAddress is a free data retrieval call binding the contract method 0xaaa13cc2.
//
// Solidity: function precalculatedWrapperAddress(uint32 originNetwork, address originTokenAddress, string name, string symbol, uint8 decimals) view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) PrecalculatedWrapperAddress(opts *bind.CallOpts, originNetwork uint32, originTokenAddress common.Address, name string, symbol string, decimals uint8) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "precalculatedWrapperAddress", originNetwork, originTokenAddress, name, symbol, decimals)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrecalculatedWrapperAddress is a free data retrieval call binding the contract method 0xaaa13cc2.
//
// Solidity: function precalculatedWrapperAddress(uint32 originNetwork, address originTokenAddress, string name, string symbol, uint8 decimals) view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) PrecalculatedWrapperAddress(originNetwork uint32, originTokenAddress common.Address, name string, symbol string, decimals uint8) (common.Address, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.PrecalculatedWrapperAddress(&_Bridgel2sovereignchainpessimistic.CallOpts, originNetwork, originTokenAddress, name, symbol, decimals)
}

// PrecalculatedWrapperAddress is a free data retrieval call binding the contract method 0xaaa13cc2.
//
// Solidity: function precalculatedWrapperAddress(uint32 originNetwork, address originTokenAddress, string name, string symbol, uint8 decimals) view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) PrecalculatedWrapperAddress(originNetwork uint32, originTokenAddress common.Address, name string, symbol string, decimals uint8) (common.Address, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.PrecalculatedWrapperAddress(&_Bridgel2sovereignchainpessimistic.CallOpts, originNetwork, originTokenAddress, name, symbol, decimals)
}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) TokenInfoToWrappedToken(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "tokenInfoToWrappedToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) TokenInfoToWrappedToken(arg0 [32]byte) (common.Address, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.TokenInfoToWrappedToken(&_Bridgel2sovereignchainpessimistic.CallOpts, arg0)
}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) TokenInfoToWrappedToken(arg0 [32]byte) (common.Address, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.TokenInfoToWrappedToken(&_Bridgel2sovereignchainpessimistic.CallOpts, arg0)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) VerifyMerkleProof(opts *bind.CallOpts, leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "verifyMerkleProof", leafHash, smtProof, index, root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.VerifyMerkleProof(&_Bridgel2sovereignchainpessimistic.CallOpts, leafHash, smtProof, index, root)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.VerifyMerkleProof(&_Bridgel2sovereignchainpessimistic.CallOpts, leafHash, smtProof, index, root)
}

// WrappedAddressIsNotMintable is a free data retrieval call binding the contract method 0xc0f49163.
//
// Solidity: function wrappedAddressIsNotMintable(address wrappedAddress) view returns(bool isNotMintable)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) WrappedAddressIsNotMintable(opts *bind.CallOpts, wrappedAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "wrappedAddressIsNotMintable", wrappedAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WrappedAddressIsNotMintable is a free data retrieval call binding the contract method 0xc0f49163.
//
// Solidity: function wrappedAddressIsNotMintable(address wrappedAddress) view returns(bool isNotMintable)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) WrappedAddressIsNotMintable(wrappedAddress common.Address) (bool, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.WrappedAddressIsNotMintable(&_Bridgel2sovereignchainpessimistic.CallOpts, wrappedAddress)
}

// WrappedAddressIsNotMintable is a free data retrieval call binding the contract method 0xc0f49163.
//
// Solidity: function wrappedAddressIsNotMintable(address wrappedAddress) view returns(bool isNotMintable)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) WrappedAddressIsNotMintable(wrappedAddress common.Address) (bool, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.WrappedAddressIsNotMintable(&_Bridgel2sovereignchainpessimistic.CallOpts, wrappedAddress)
}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address ) view returns(uint32 originNetwork, address originTokenAddress)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCaller) WrappedTokenToTokenInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	var out []interface{}
	err := _Bridgel2sovereignchainpessimistic.contract.Call(opts, &out, "wrappedTokenToTokenInfo", arg0)

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
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) WrappedTokenToTokenInfo(arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.WrappedTokenToTokenInfo(&_Bridgel2sovereignchainpessimistic.CallOpts, arg0)
}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address ) view returns(uint32 originNetwork, address originTokenAddress)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticCallerSession) WrappedTokenToTokenInfo(arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.WrappedTokenToTokenInfo(&_Bridgel2sovereignchainpessimistic.CallOpts, arg0)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactor) BridgeAsset(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.contract.Transact(opts, "bridgeAsset", destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) BridgeAsset(destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.BridgeAsset(&_Bridgel2sovereignchainpessimistic.TransactOpts, destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactorSession) BridgeAsset(destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.BridgeAsset(&_Bridgel2sovereignchainpessimistic.TransactOpts, destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactor) BridgeMessage(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.contract.Transact(opts, "bridgeMessage", destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) BridgeMessage(destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.BridgeMessage(&_Bridgel2sovereignchainpessimistic.TransactOpts, destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactorSession) BridgeMessage(destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.BridgeMessage(&_Bridgel2sovereignchainpessimistic.TransactOpts, destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactor) BridgeMessageWETH(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.contract.Transact(opts, "bridgeMessageWETH", destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) BridgeMessageWETH(destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.BridgeMessageWETH(&_Bridgel2sovereignchainpessimistic.TransactOpts, destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactorSession) BridgeMessageWETH(destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.BridgeMessageWETH(&_Bridgel2sovereignchainpessimistic.TransactOpts, destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactor) ClaimAsset(opts *bind.TransactOpts, smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.contract.Transact(opts, "claimAsset", smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) ClaimAsset(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.ClaimAsset(&_Bridgel2sovereignchainpessimistic.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactorSession) ClaimAsset(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.ClaimAsset(&_Bridgel2sovereignchainpessimistic.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactor) ClaimMessage(opts *bind.TransactOpts, smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.contract.Transact(opts, "claimMessage", smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) ClaimMessage(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.ClaimMessage(&_Bridgel2sovereignchainpessimistic.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactorSession) ClaimMessage(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.ClaimMessage(&_Bridgel2sovereignchainpessimistic.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// Initialize is a paid mutator transaction binding the contract method 0x8c0dd470.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address _sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactor) Initialize(opts *bind.TransactOpts, _networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, _sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.contract.Transact(opts, "initialize", _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, _sovereignWETHAddress, _sovereignWETHAddressIsNotMintable)
}

// Initialize is a paid mutator transaction binding the contract method 0x8c0dd470.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address _sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) Initialize(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, _sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.Initialize(&_Bridgel2sovereignchainpessimistic.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, _sovereignWETHAddress, _sovereignWETHAddressIsNotMintable)
}

// Initialize is a paid mutator transaction binding the contract method 0x8c0dd470.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address _sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactorSession) Initialize(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, _sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.Initialize(&_Bridgel2sovereignchainpessimistic.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, _sovereignWETHAddress, _sovereignWETHAddressIsNotMintable)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes ) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactor) Initialize0(opts *bind.TransactOpts, arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.contract.Transact(opts, "initialize0", arg0, arg1, arg2, arg3, arg4, arg5)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes ) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) Initialize0(arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.Initialize0(&_Bridgel2sovereignchainpessimistic.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes ) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactorSession) Initialize0(arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.Initialize0(&_Bridgel2sovereignchainpessimistic.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// MigrateLegacyToken is a paid mutator transaction binding the contract method 0x9e76158f.
//
// Solidity: function migrateLegacyToken(address legacyTokenAddress, uint256 amount) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactor) MigrateLegacyToken(opts *bind.TransactOpts, legacyTokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.contract.Transact(opts, "migrateLegacyToken", legacyTokenAddress, amount)
}

// MigrateLegacyToken is a paid mutator transaction binding the contract method 0x9e76158f.
//
// Solidity: function migrateLegacyToken(address legacyTokenAddress, uint256 amount) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) MigrateLegacyToken(legacyTokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.MigrateLegacyToken(&_Bridgel2sovereignchainpessimistic.TransactOpts, legacyTokenAddress, amount)
}

// MigrateLegacyToken is a paid mutator transaction binding the contract method 0x9e76158f.
//
// Solidity: function migrateLegacyToken(address legacyTokenAddress, uint256 amount) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactorSession) MigrateLegacyToken(legacyTokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.MigrateLegacyToken(&_Bridgel2sovereignchainpessimistic.TransactOpts, legacyTokenAddress, amount)
}

// RemoveLegacySovereignTokenAddress is a paid mutator transaction binding the contract method 0xb4586962.
//
// Solidity: function removeLegacySovereignTokenAddress(address legacySovereignTokenAddress) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactor) RemoveLegacySovereignTokenAddress(opts *bind.TransactOpts, legacySovereignTokenAddress common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.contract.Transact(opts, "removeLegacySovereignTokenAddress", legacySovereignTokenAddress)
}

// RemoveLegacySovereignTokenAddress is a paid mutator transaction binding the contract method 0xb4586962.
//
// Solidity: function removeLegacySovereignTokenAddress(address legacySovereignTokenAddress) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) RemoveLegacySovereignTokenAddress(legacySovereignTokenAddress common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.RemoveLegacySovereignTokenAddress(&_Bridgel2sovereignchainpessimistic.TransactOpts, legacySovereignTokenAddress)
}

// RemoveLegacySovereignTokenAddress is a paid mutator transaction binding the contract method 0xb4586962.
//
// Solidity: function removeLegacySovereignTokenAddress(address legacySovereignTokenAddress) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactorSession) RemoveLegacySovereignTokenAddress(legacySovereignTokenAddress common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.RemoveLegacySovereignTokenAddress(&_Bridgel2sovereignchainpessimistic.TransactOpts, legacySovereignTokenAddress)
}

// SetBridgeManager is a paid mutator transaction binding the contract method 0xeabd372a.
//
// Solidity: function setBridgeManager(address _bridgeManager) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactor) SetBridgeManager(opts *bind.TransactOpts, _bridgeManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.contract.Transact(opts, "setBridgeManager", _bridgeManager)
}

// SetBridgeManager is a paid mutator transaction binding the contract method 0xeabd372a.
//
// Solidity: function setBridgeManager(address _bridgeManager) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) SetBridgeManager(_bridgeManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.SetBridgeManager(&_Bridgel2sovereignchainpessimistic.TransactOpts, _bridgeManager)
}

// SetBridgeManager is a paid mutator transaction binding the contract method 0xeabd372a.
//
// Solidity: function setBridgeManager(address _bridgeManager) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactorSession) SetBridgeManager(_bridgeManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.SetBridgeManager(&_Bridgel2sovereignchainpessimistic.TransactOpts, _bridgeManager)
}

// SetMultipleSovereignTokenAddress is a paid mutator transaction binding the contract method 0x57cfbee3.
//
// Solidity: function setMultipleSovereignTokenAddress(uint32[] originNetworks, address[] originTokenAddresses, address[] sovereignTokenAddresses, bool[] isNotMintable) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactor) SetMultipleSovereignTokenAddress(opts *bind.TransactOpts, originNetworks []uint32, originTokenAddresses []common.Address, sovereignTokenAddresses []common.Address, isNotMintable []bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.contract.Transact(opts, "setMultipleSovereignTokenAddress", originNetworks, originTokenAddresses, sovereignTokenAddresses, isNotMintable)
}

// SetMultipleSovereignTokenAddress is a paid mutator transaction binding the contract method 0x57cfbee3.
//
// Solidity: function setMultipleSovereignTokenAddress(uint32[] originNetworks, address[] originTokenAddresses, address[] sovereignTokenAddresses, bool[] isNotMintable) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) SetMultipleSovereignTokenAddress(originNetworks []uint32, originTokenAddresses []common.Address, sovereignTokenAddresses []common.Address, isNotMintable []bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.SetMultipleSovereignTokenAddress(&_Bridgel2sovereignchainpessimistic.TransactOpts, originNetworks, originTokenAddresses, sovereignTokenAddresses, isNotMintable)
}

// SetMultipleSovereignTokenAddress is a paid mutator transaction binding the contract method 0x57cfbee3.
//
// Solidity: function setMultipleSovereignTokenAddress(uint32[] originNetworks, address[] originTokenAddresses, address[] sovereignTokenAddresses, bool[] isNotMintable) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactorSession) SetMultipleSovereignTokenAddress(originNetworks []uint32, originTokenAddresses []common.Address, sovereignTokenAddresses []common.Address, isNotMintable []bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.SetMultipleSovereignTokenAddress(&_Bridgel2sovereignchainpessimistic.TransactOpts, originNetworks, originTokenAddresses, sovereignTokenAddresses, isNotMintable)
}

// SetSovereignWETHAddress is a paid mutator transaction binding the contract method 0xbf130d7f.
//
// Solidity: function setSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactor) SetSovereignWETHAddress(opts *bind.TransactOpts, sovereignWETHTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.contract.Transact(opts, "setSovereignWETHAddress", sovereignWETHTokenAddress, isNotMintable)
}

// SetSovereignWETHAddress is a paid mutator transaction binding the contract method 0xbf130d7f.
//
// Solidity: function setSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) SetSovereignWETHAddress(sovereignWETHTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.SetSovereignWETHAddress(&_Bridgel2sovereignchainpessimistic.TransactOpts, sovereignWETHTokenAddress, isNotMintable)
}

// SetSovereignWETHAddress is a paid mutator transaction binding the contract method 0xbf130d7f.
//
// Solidity: function setSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactorSession) SetSovereignWETHAddress(sovereignWETHTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.SetSovereignWETHAddress(&_Bridgel2sovereignchainpessimistic.TransactOpts, sovereignWETHTokenAddress, isNotMintable)
}

// UnsetMultipleClaimedBitmap is a paid mutator transaction binding the contract method 0x8781a5c5.
//
// Solidity: function unsetMultipleClaimedBitmap(uint32[] leafIndexes, uint32[] sourceBridgeNetworks) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactor) UnsetMultipleClaimedBitmap(opts *bind.TransactOpts, leafIndexes []uint32, sourceBridgeNetworks []uint32) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.contract.Transact(opts, "unsetMultipleClaimedBitmap", leafIndexes, sourceBridgeNetworks)
}

// UnsetMultipleClaimedBitmap is a paid mutator transaction binding the contract method 0x8781a5c5.
//
// Solidity: function unsetMultipleClaimedBitmap(uint32[] leafIndexes, uint32[] sourceBridgeNetworks) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) UnsetMultipleClaimedBitmap(leafIndexes []uint32, sourceBridgeNetworks []uint32) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.UnsetMultipleClaimedBitmap(&_Bridgel2sovereignchainpessimistic.TransactOpts, leafIndexes, sourceBridgeNetworks)
}

// UnsetMultipleClaimedBitmap is a paid mutator transaction binding the contract method 0x8781a5c5.
//
// Solidity: function unsetMultipleClaimedBitmap(uint32[] leafIndexes, uint32[] sourceBridgeNetworks) returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactorSession) UnsetMultipleClaimedBitmap(leafIndexes []uint32, sourceBridgeNetworks []uint32) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.UnsetMultipleClaimedBitmap(&_Bridgel2sovereignchainpessimistic.TransactOpts, leafIndexes, sourceBridgeNetworks)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactor) UpdateGlobalExitRoot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.contract.Transact(opts, "updateGlobalExitRoot")
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticSession) UpdateGlobalExitRoot() (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.UpdateGlobalExitRoot(&_Bridgel2sovereignchainpessimistic.TransactOpts)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticTransactorSession) UpdateGlobalExitRoot() (*types.Transaction, error) {
	return _Bridgel2sovereignchainpessimistic.Contract.UpdateGlobalExitRoot(&_Bridgel2sovereignchainpessimistic.TransactOpts)
}

// Bridgel2sovereignchainpessimisticBridgeEventIterator is returned from FilterBridgeEvent and is used to iterate over the raw logs and unpacked data for BridgeEvent events raised by the Bridgel2sovereignchainpessimistic contract.
type Bridgel2sovereignchainpessimisticBridgeEventIterator struct {
	Event *Bridgel2sovereignchainpessimisticBridgeEvent // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainpessimisticBridgeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainpessimisticBridgeEvent)
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
		it.Event = new(Bridgel2sovereignchainpessimisticBridgeEvent)
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
func (it *Bridgel2sovereignchainpessimisticBridgeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainpessimisticBridgeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainpessimisticBridgeEvent represents a BridgeEvent event raised by the Bridgel2sovereignchainpessimistic contract.
type Bridgel2sovereignchainpessimisticBridgeEvent struct {
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
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) FilterBridgeEvent(opts *bind.FilterOpts) (*Bridgel2sovereignchainpessimisticBridgeEventIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainpessimistic.contract.FilterLogs(opts, "BridgeEvent")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainpessimisticBridgeEventIterator{contract: _Bridgel2sovereignchainpessimistic.contract, event: "BridgeEvent", logs: logs, sub: sub}, nil
}

// WatchBridgeEvent is a free log subscription operation binding the contract event 0x501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b.
//
// Solidity: event BridgeEvent(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata, uint32 depositCount)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) WatchBridgeEvent(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainpessimisticBridgeEvent) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainpessimistic.contract.WatchLogs(opts, "BridgeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainpessimisticBridgeEvent)
				if err := _Bridgel2sovereignchainpessimistic.contract.UnpackLog(event, "BridgeEvent", log); err != nil {
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
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) ParseBridgeEvent(log types.Log) (*Bridgel2sovereignchainpessimisticBridgeEvent, error) {
	event := new(Bridgel2sovereignchainpessimisticBridgeEvent)
	if err := _Bridgel2sovereignchainpessimistic.contract.UnpackLog(event, "BridgeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainpessimisticClaimEventIterator is returned from FilterClaimEvent and is used to iterate over the raw logs and unpacked data for ClaimEvent events raised by the Bridgel2sovereignchainpessimistic contract.
type Bridgel2sovereignchainpessimisticClaimEventIterator struct {
	Event *Bridgel2sovereignchainpessimisticClaimEvent // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainpessimisticClaimEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainpessimisticClaimEvent)
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
		it.Event = new(Bridgel2sovereignchainpessimisticClaimEvent)
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
func (it *Bridgel2sovereignchainpessimisticClaimEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainpessimisticClaimEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainpessimisticClaimEvent represents a ClaimEvent event raised by the Bridgel2sovereignchainpessimistic contract.
type Bridgel2sovereignchainpessimisticClaimEvent struct {
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
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) FilterClaimEvent(opts *bind.FilterOpts) (*Bridgel2sovereignchainpessimisticClaimEventIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainpessimistic.contract.FilterLogs(opts, "ClaimEvent")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainpessimisticClaimEventIterator{contract: _Bridgel2sovereignchainpessimistic.contract, event: "ClaimEvent", logs: logs, sub: sub}, nil
}

// WatchClaimEvent is a free log subscription operation binding the contract event 0x1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d.
//
// Solidity: event ClaimEvent(uint256 globalIndex, uint32 originNetwork, address originAddress, address destinationAddress, uint256 amount)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) WatchClaimEvent(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainpessimisticClaimEvent) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainpessimistic.contract.WatchLogs(opts, "ClaimEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainpessimisticClaimEvent)
				if err := _Bridgel2sovereignchainpessimistic.contract.UnpackLog(event, "ClaimEvent", log); err != nil {
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
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) ParseClaimEvent(log types.Log) (*Bridgel2sovereignchainpessimisticClaimEvent, error) {
	event := new(Bridgel2sovereignchainpessimisticClaimEvent)
	if err := _Bridgel2sovereignchainpessimistic.contract.UnpackLog(event, "ClaimEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainpessimisticEmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the Bridgel2sovereignchainpessimistic contract.
type Bridgel2sovereignchainpessimisticEmergencyStateActivatedIterator struct {
	Event *Bridgel2sovereignchainpessimisticEmergencyStateActivated // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainpessimisticEmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainpessimisticEmergencyStateActivated)
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
		it.Event = new(Bridgel2sovereignchainpessimisticEmergencyStateActivated)
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
func (it *Bridgel2sovereignchainpessimisticEmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainpessimisticEmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainpessimisticEmergencyStateActivated represents a EmergencyStateActivated event raised by the Bridgel2sovereignchainpessimistic contract.
type Bridgel2sovereignchainpessimisticEmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*Bridgel2sovereignchainpessimisticEmergencyStateActivatedIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainpessimistic.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainpessimisticEmergencyStateActivatedIterator{contract: _Bridgel2sovereignchainpessimistic.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainpessimisticEmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainpessimistic.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainpessimisticEmergencyStateActivated)
				if err := _Bridgel2sovereignchainpessimistic.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
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
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) ParseEmergencyStateActivated(log types.Log) (*Bridgel2sovereignchainpessimisticEmergencyStateActivated, error) {
	event := new(Bridgel2sovereignchainpessimisticEmergencyStateActivated)
	if err := _Bridgel2sovereignchainpessimistic.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainpessimisticEmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the Bridgel2sovereignchainpessimistic contract.
type Bridgel2sovereignchainpessimisticEmergencyStateDeactivatedIterator struct {
	Event *Bridgel2sovereignchainpessimisticEmergencyStateDeactivated // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainpessimisticEmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainpessimisticEmergencyStateDeactivated)
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
		it.Event = new(Bridgel2sovereignchainpessimisticEmergencyStateDeactivated)
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
func (it *Bridgel2sovereignchainpessimisticEmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainpessimisticEmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainpessimisticEmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the Bridgel2sovereignchainpessimistic contract.
type Bridgel2sovereignchainpessimisticEmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*Bridgel2sovereignchainpessimisticEmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainpessimistic.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainpessimisticEmergencyStateDeactivatedIterator{contract: _Bridgel2sovereignchainpessimistic.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainpessimisticEmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainpessimistic.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainpessimisticEmergencyStateDeactivated)
				if err := _Bridgel2sovereignchainpessimistic.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
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
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) ParseEmergencyStateDeactivated(log types.Log) (*Bridgel2sovereignchainpessimisticEmergencyStateDeactivated, error) {
	event := new(Bridgel2sovereignchainpessimisticEmergencyStateDeactivated)
	if err := _Bridgel2sovereignchainpessimistic.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainpessimisticInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bridgel2sovereignchainpessimistic contract.
type Bridgel2sovereignchainpessimisticInitializedIterator struct {
	Event *Bridgel2sovereignchainpessimisticInitialized // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainpessimisticInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainpessimisticInitialized)
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
		it.Event = new(Bridgel2sovereignchainpessimisticInitialized)
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
func (it *Bridgel2sovereignchainpessimisticInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainpessimisticInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainpessimisticInitialized represents a Initialized event raised by the Bridgel2sovereignchainpessimistic contract.
type Bridgel2sovereignchainpessimisticInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) FilterInitialized(opts *bind.FilterOpts) (*Bridgel2sovereignchainpessimisticInitializedIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainpessimistic.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainpessimisticInitializedIterator{contract: _Bridgel2sovereignchainpessimistic.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainpessimisticInitialized) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainpessimistic.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainpessimisticInitialized)
				if err := _Bridgel2sovereignchainpessimistic.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) ParseInitialized(log types.Log) (*Bridgel2sovereignchainpessimisticInitialized, error) {
	event := new(Bridgel2sovereignchainpessimisticInitialized)
	if err := _Bridgel2sovereignchainpessimistic.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainpessimisticMigrateLegacyTokenIterator is returned from FilterMigrateLegacyToken and is used to iterate over the raw logs and unpacked data for MigrateLegacyToken events raised by the Bridgel2sovereignchainpessimistic contract.
type Bridgel2sovereignchainpessimisticMigrateLegacyTokenIterator struct {
	Event *Bridgel2sovereignchainpessimisticMigrateLegacyToken // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainpessimisticMigrateLegacyTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainpessimisticMigrateLegacyToken)
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
		it.Event = new(Bridgel2sovereignchainpessimisticMigrateLegacyToken)
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
func (it *Bridgel2sovereignchainpessimisticMigrateLegacyTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainpessimisticMigrateLegacyTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainpessimisticMigrateLegacyToken represents a MigrateLegacyToken event raised by the Bridgel2sovereignchainpessimistic contract.
type Bridgel2sovereignchainpessimisticMigrateLegacyToken struct {
	Sender              common.Address
	LegacyTokenAddress  common.Address
	UpdatedTokenAddress common.Address
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMigrateLegacyToken is a free log retrieval operation binding the contract event 0xb7f8fd4d1faf9b2929dc269f59c53e3a2bccc44e9950f33a568fcbcb37eb69a9.
//
// Solidity: event MigrateLegacyToken(address sender, address legacyTokenAddress, address updatedTokenAddress, uint256 amount)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) FilterMigrateLegacyToken(opts *bind.FilterOpts) (*Bridgel2sovereignchainpessimisticMigrateLegacyTokenIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainpessimistic.contract.FilterLogs(opts, "MigrateLegacyToken")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainpessimisticMigrateLegacyTokenIterator{contract: _Bridgel2sovereignchainpessimistic.contract, event: "MigrateLegacyToken", logs: logs, sub: sub}, nil
}

// WatchMigrateLegacyToken is a free log subscription operation binding the contract event 0xb7f8fd4d1faf9b2929dc269f59c53e3a2bccc44e9950f33a568fcbcb37eb69a9.
//
// Solidity: event MigrateLegacyToken(address sender, address legacyTokenAddress, address updatedTokenAddress, uint256 amount)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) WatchMigrateLegacyToken(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainpessimisticMigrateLegacyToken) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainpessimistic.contract.WatchLogs(opts, "MigrateLegacyToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainpessimisticMigrateLegacyToken)
				if err := _Bridgel2sovereignchainpessimistic.contract.UnpackLog(event, "MigrateLegacyToken", log); err != nil {
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
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) ParseMigrateLegacyToken(log types.Log) (*Bridgel2sovereignchainpessimisticMigrateLegacyToken, error) {
	event := new(Bridgel2sovereignchainpessimisticMigrateLegacyToken)
	if err := _Bridgel2sovereignchainpessimistic.contract.UnpackLog(event, "MigrateLegacyToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainpessimisticNewWrappedTokenIterator is returned from FilterNewWrappedToken and is used to iterate over the raw logs and unpacked data for NewWrappedToken events raised by the Bridgel2sovereignchainpessimistic contract.
type Bridgel2sovereignchainpessimisticNewWrappedTokenIterator struct {
	Event *Bridgel2sovereignchainpessimisticNewWrappedToken // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainpessimisticNewWrappedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainpessimisticNewWrappedToken)
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
		it.Event = new(Bridgel2sovereignchainpessimisticNewWrappedToken)
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
func (it *Bridgel2sovereignchainpessimisticNewWrappedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainpessimisticNewWrappedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainpessimisticNewWrappedToken represents a NewWrappedToken event raised by the Bridgel2sovereignchainpessimistic contract.
type Bridgel2sovereignchainpessimisticNewWrappedToken struct {
	OriginNetwork       uint32
	OriginTokenAddress  common.Address
	WrappedTokenAddress common.Address
	Metadata            []byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewWrappedToken is a free log retrieval operation binding the contract event 0x490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a39.
//
// Solidity: event NewWrappedToken(uint32 originNetwork, address originTokenAddress, address wrappedTokenAddress, bytes metadata)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) FilterNewWrappedToken(opts *bind.FilterOpts) (*Bridgel2sovereignchainpessimisticNewWrappedTokenIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainpessimistic.contract.FilterLogs(opts, "NewWrappedToken")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainpessimisticNewWrappedTokenIterator{contract: _Bridgel2sovereignchainpessimistic.contract, event: "NewWrappedToken", logs: logs, sub: sub}, nil
}

// WatchNewWrappedToken is a free log subscription operation binding the contract event 0x490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a39.
//
// Solidity: event NewWrappedToken(uint32 originNetwork, address originTokenAddress, address wrappedTokenAddress, bytes metadata)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) WatchNewWrappedToken(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainpessimisticNewWrappedToken) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainpessimistic.contract.WatchLogs(opts, "NewWrappedToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainpessimisticNewWrappedToken)
				if err := _Bridgel2sovereignchainpessimistic.contract.UnpackLog(event, "NewWrappedToken", log); err != nil {
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
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) ParseNewWrappedToken(log types.Log) (*Bridgel2sovereignchainpessimisticNewWrappedToken, error) {
	event := new(Bridgel2sovereignchainpessimisticNewWrappedToken)
	if err := _Bridgel2sovereignchainpessimistic.contract.UnpackLog(event, "NewWrappedToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainpessimisticRemoveLegacySovereignTokenAddressIterator is returned from FilterRemoveLegacySovereignTokenAddress and is used to iterate over the raw logs and unpacked data for RemoveLegacySovereignTokenAddress events raised by the Bridgel2sovereignchainpessimistic contract.
type Bridgel2sovereignchainpessimisticRemoveLegacySovereignTokenAddressIterator struct {
	Event *Bridgel2sovereignchainpessimisticRemoveLegacySovereignTokenAddress // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainpessimisticRemoveLegacySovereignTokenAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainpessimisticRemoveLegacySovereignTokenAddress)
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
		it.Event = new(Bridgel2sovereignchainpessimisticRemoveLegacySovereignTokenAddress)
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
func (it *Bridgel2sovereignchainpessimisticRemoveLegacySovereignTokenAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainpessimisticRemoveLegacySovereignTokenAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainpessimisticRemoveLegacySovereignTokenAddress represents a RemoveLegacySovereignTokenAddress event raised by the Bridgel2sovereignchainpessimistic contract.
type Bridgel2sovereignchainpessimisticRemoveLegacySovereignTokenAddress struct {
	SovereignTokenAddress common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRemoveLegacySovereignTokenAddress is a free log retrieval operation binding the contract event 0xc2ae0bd0ec0fd0352bfe5bacac49637af342c1e40f1b80a7f74440dc7fe3f063.
//
// Solidity: event RemoveLegacySovereignTokenAddress(address sovereignTokenAddress)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) FilterRemoveLegacySovereignTokenAddress(opts *bind.FilterOpts) (*Bridgel2sovereignchainpessimisticRemoveLegacySovereignTokenAddressIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainpessimistic.contract.FilterLogs(opts, "RemoveLegacySovereignTokenAddress")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainpessimisticRemoveLegacySovereignTokenAddressIterator{contract: _Bridgel2sovereignchainpessimistic.contract, event: "RemoveLegacySovereignTokenAddress", logs: logs, sub: sub}, nil
}

// WatchRemoveLegacySovereignTokenAddress is a free log subscription operation binding the contract event 0xc2ae0bd0ec0fd0352bfe5bacac49637af342c1e40f1b80a7f74440dc7fe3f063.
//
// Solidity: event RemoveLegacySovereignTokenAddress(address sovereignTokenAddress)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) WatchRemoveLegacySovereignTokenAddress(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainpessimisticRemoveLegacySovereignTokenAddress) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainpessimistic.contract.WatchLogs(opts, "RemoveLegacySovereignTokenAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainpessimisticRemoveLegacySovereignTokenAddress)
				if err := _Bridgel2sovereignchainpessimistic.contract.UnpackLog(event, "RemoveLegacySovereignTokenAddress", log); err != nil {
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
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) ParseRemoveLegacySovereignTokenAddress(log types.Log) (*Bridgel2sovereignchainpessimisticRemoveLegacySovereignTokenAddress, error) {
	event := new(Bridgel2sovereignchainpessimisticRemoveLegacySovereignTokenAddress)
	if err := _Bridgel2sovereignchainpessimistic.contract.UnpackLog(event, "RemoveLegacySovereignTokenAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainpessimisticSetBridgeManagerIterator is returned from FilterSetBridgeManager and is used to iterate over the raw logs and unpacked data for SetBridgeManager events raised by the Bridgel2sovereignchainpessimistic contract.
type Bridgel2sovereignchainpessimisticSetBridgeManagerIterator struct {
	Event *Bridgel2sovereignchainpessimisticSetBridgeManager // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainpessimisticSetBridgeManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainpessimisticSetBridgeManager)
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
		it.Event = new(Bridgel2sovereignchainpessimisticSetBridgeManager)
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
func (it *Bridgel2sovereignchainpessimisticSetBridgeManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainpessimisticSetBridgeManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainpessimisticSetBridgeManager represents a SetBridgeManager event raised by the Bridgel2sovereignchainpessimistic contract.
type Bridgel2sovereignchainpessimisticSetBridgeManager struct {
	BridgeManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetBridgeManager is a free log retrieval operation binding the contract event 0x32cf74f8a6d5f88593984d2cd52be5592bfa6884f5896175801a5069ef09cd67.
//
// Solidity: event SetBridgeManager(address bridgeManager)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) FilterSetBridgeManager(opts *bind.FilterOpts) (*Bridgel2sovereignchainpessimisticSetBridgeManagerIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainpessimistic.contract.FilterLogs(opts, "SetBridgeManager")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainpessimisticSetBridgeManagerIterator{contract: _Bridgel2sovereignchainpessimistic.contract, event: "SetBridgeManager", logs: logs, sub: sub}, nil
}

// WatchSetBridgeManager is a free log subscription operation binding the contract event 0x32cf74f8a6d5f88593984d2cd52be5592bfa6884f5896175801a5069ef09cd67.
//
// Solidity: event SetBridgeManager(address bridgeManager)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) WatchSetBridgeManager(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainpessimisticSetBridgeManager) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainpessimistic.contract.WatchLogs(opts, "SetBridgeManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainpessimisticSetBridgeManager)
				if err := _Bridgel2sovereignchainpessimistic.contract.UnpackLog(event, "SetBridgeManager", log); err != nil {
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
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) ParseSetBridgeManager(log types.Log) (*Bridgel2sovereignchainpessimisticSetBridgeManager, error) {
	event := new(Bridgel2sovereignchainpessimisticSetBridgeManager)
	if err := _Bridgel2sovereignchainpessimistic.contract.UnpackLog(event, "SetBridgeManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainpessimisticSetSovereignTokenAddressIterator is returned from FilterSetSovereignTokenAddress and is used to iterate over the raw logs and unpacked data for SetSovereignTokenAddress events raised by the Bridgel2sovereignchainpessimistic contract.
type Bridgel2sovereignchainpessimisticSetSovereignTokenAddressIterator struct {
	Event *Bridgel2sovereignchainpessimisticSetSovereignTokenAddress // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainpessimisticSetSovereignTokenAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainpessimisticSetSovereignTokenAddress)
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
		it.Event = new(Bridgel2sovereignchainpessimisticSetSovereignTokenAddress)
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
func (it *Bridgel2sovereignchainpessimisticSetSovereignTokenAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainpessimisticSetSovereignTokenAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainpessimisticSetSovereignTokenAddress represents a SetSovereignTokenAddress event raised by the Bridgel2sovereignchainpessimistic contract.
type Bridgel2sovereignchainpessimisticSetSovereignTokenAddress struct {
	OriginNetwork         uint32
	OriginTokenAddress    common.Address
	SovereignTokenAddress common.Address
	IsNotMintable         bool
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetSovereignTokenAddress is a free log retrieval operation binding the contract event 0xdbe8a5da6a7a916d9adfda9160167a0f8a3da415ee6610e810e753853597fce7.
//
// Solidity: event SetSovereignTokenAddress(uint32 originNetwork, address originTokenAddress, address sovereignTokenAddress, bool isNotMintable)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) FilterSetSovereignTokenAddress(opts *bind.FilterOpts) (*Bridgel2sovereignchainpessimisticSetSovereignTokenAddressIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainpessimistic.contract.FilterLogs(opts, "SetSovereignTokenAddress")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainpessimisticSetSovereignTokenAddressIterator{contract: _Bridgel2sovereignchainpessimistic.contract, event: "SetSovereignTokenAddress", logs: logs, sub: sub}, nil
}

// WatchSetSovereignTokenAddress is a free log subscription operation binding the contract event 0xdbe8a5da6a7a916d9adfda9160167a0f8a3da415ee6610e810e753853597fce7.
//
// Solidity: event SetSovereignTokenAddress(uint32 originNetwork, address originTokenAddress, address sovereignTokenAddress, bool isNotMintable)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) WatchSetSovereignTokenAddress(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainpessimisticSetSovereignTokenAddress) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainpessimistic.contract.WatchLogs(opts, "SetSovereignTokenAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainpessimisticSetSovereignTokenAddress)
				if err := _Bridgel2sovereignchainpessimistic.contract.UnpackLog(event, "SetSovereignTokenAddress", log); err != nil {
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
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) ParseSetSovereignTokenAddress(log types.Log) (*Bridgel2sovereignchainpessimisticSetSovereignTokenAddress, error) {
	event := new(Bridgel2sovereignchainpessimisticSetSovereignTokenAddress)
	if err := _Bridgel2sovereignchainpessimistic.contract.UnpackLog(event, "SetSovereignTokenAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainpessimisticSetSovereignWETHAddressIterator is returned from FilterSetSovereignWETHAddress and is used to iterate over the raw logs and unpacked data for SetSovereignWETHAddress events raised by the Bridgel2sovereignchainpessimistic contract.
type Bridgel2sovereignchainpessimisticSetSovereignWETHAddressIterator struct {
	Event *Bridgel2sovereignchainpessimisticSetSovereignWETHAddress // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainpessimisticSetSovereignWETHAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainpessimisticSetSovereignWETHAddress)
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
		it.Event = new(Bridgel2sovereignchainpessimisticSetSovereignWETHAddress)
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
func (it *Bridgel2sovereignchainpessimisticSetSovereignWETHAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainpessimisticSetSovereignWETHAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainpessimisticSetSovereignWETHAddress represents a SetSovereignWETHAddress event raised by the Bridgel2sovereignchainpessimistic contract.
type Bridgel2sovereignchainpessimisticSetSovereignWETHAddress struct {
	SovereignWETHTokenAddress common.Address
	IsNotMintable             bool
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterSetSovereignWETHAddress is a free log retrieval operation binding the contract event 0xc7318b7ed6ba4f2908a3de396d8ab49b1dadb55db5b55123247a401f29ff8d82.
//
// Solidity: event SetSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) FilterSetSovereignWETHAddress(opts *bind.FilterOpts) (*Bridgel2sovereignchainpessimisticSetSovereignWETHAddressIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainpessimistic.contract.FilterLogs(opts, "SetSovereignWETHAddress")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainpessimisticSetSovereignWETHAddressIterator{contract: _Bridgel2sovereignchainpessimistic.contract, event: "SetSovereignWETHAddress", logs: logs, sub: sub}, nil
}

// WatchSetSovereignWETHAddress is a free log subscription operation binding the contract event 0xc7318b7ed6ba4f2908a3de396d8ab49b1dadb55db5b55123247a401f29ff8d82.
//
// Solidity: event SetSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) WatchSetSovereignWETHAddress(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainpessimisticSetSovereignWETHAddress) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainpessimistic.contract.WatchLogs(opts, "SetSovereignWETHAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainpessimisticSetSovereignWETHAddress)
				if err := _Bridgel2sovereignchainpessimistic.contract.UnpackLog(event, "SetSovereignWETHAddress", log); err != nil {
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
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) ParseSetSovereignWETHAddress(log types.Log) (*Bridgel2sovereignchainpessimisticSetSovereignWETHAddress, error) {
	event := new(Bridgel2sovereignchainpessimisticSetSovereignWETHAddress)
	if err := _Bridgel2sovereignchainpessimistic.contract.UnpackLog(event, "SetSovereignWETHAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainpessimisticUnsetClaimIterator is returned from FilterUnsetClaim and is used to iterate over the raw logs and unpacked data for UnsetClaim events raised by the Bridgel2sovereignchainpessimistic contract.
type Bridgel2sovereignchainpessimisticUnsetClaimIterator struct {
	Event *Bridgel2sovereignchainpessimisticUnsetClaim // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainpessimisticUnsetClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainpessimisticUnsetClaim)
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
		it.Event = new(Bridgel2sovereignchainpessimisticUnsetClaim)
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
func (it *Bridgel2sovereignchainpessimisticUnsetClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainpessimisticUnsetClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainpessimisticUnsetClaim represents a UnsetClaim event raised by the Bridgel2sovereignchainpessimistic contract.
type Bridgel2sovereignchainpessimisticUnsetClaim struct {
	LeafIndex           uint32
	SourceBridgeNetwork uint32
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUnsetClaim is a free log retrieval operation binding the contract event 0x4a402ac607e71571d0543be8fcccc358a4df62dcd019a1e7f7e98ca6175e8f2a.
//
// Solidity: event UnsetClaim(uint32 leafIndex, uint32 sourceBridgeNetwork)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) FilterUnsetClaim(opts *bind.FilterOpts) (*Bridgel2sovereignchainpessimisticUnsetClaimIterator, error) {

	logs, sub, err := _Bridgel2sovereignchainpessimistic.contract.FilterLogs(opts, "UnsetClaim")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainpessimisticUnsetClaimIterator{contract: _Bridgel2sovereignchainpessimistic.contract, event: "UnsetClaim", logs: logs, sub: sub}, nil
}

// WatchUnsetClaim is a free log subscription operation binding the contract event 0x4a402ac607e71571d0543be8fcccc358a4df62dcd019a1e7f7e98ca6175e8f2a.
//
// Solidity: event UnsetClaim(uint32 leafIndex, uint32 sourceBridgeNetwork)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) WatchUnsetClaim(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainpessimisticUnsetClaim) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchainpessimistic.contract.WatchLogs(opts, "UnsetClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainpessimisticUnsetClaim)
				if err := _Bridgel2sovereignchainpessimistic.contract.UnpackLog(event, "UnsetClaim", log); err != nil {
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

// ParseUnsetClaim is a log parse operation binding the contract event 0x4a402ac607e71571d0543be8fcccc358a4df62dcd019a1e7f7e98ca6175e8f2a.
//
// Solidity: event UnsetClaim(uint32 leafIndex, uint32 sourceBridgeNetwork)
func (_Bridgel2sovereignchainpessimistic *Bridgel2sovereignchainpessimisticFilterer) ParseUnsetClaim(log types.Log) (*Bridgel2sovereignchainpessimisticUnsetClaim, error) {
	event := new(Bridgel2sovereignchainpessimisticUnsetClaim)
	if err := _Bridgel2sovereignchainpessimistic.contract.UnpackLog(event, "UnsetClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

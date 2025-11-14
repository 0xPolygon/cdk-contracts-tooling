// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridgemessagereceivermock

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

// BridgemessagereceivermockMetaData contains all meta data concerning the Bridgemessagereceivermock contract.
var BridgemessagereceivermockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractAgglayerBridge\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"}],\"name\":\"MessageReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"UpdateParameters\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractAgglayerBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onMessageReceived\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"claimData1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bridgeAsset\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"claimData2\",\"type\":\"bytes\"}],\"name\":\"testClaim\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"msmtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"msmtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"mglobalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mmainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"mrollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moriginNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"moriginAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"mdestinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"mdestinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mamount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"mmetadata\",\"type\":\"bytes\"}],\"name\":\"updateParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523461006257610019610014610139565b61015a565b610021610067565b611a6861016082396080518181816104cf015281816108bc0152818161163d0152818161170c015281816117dd0152818161186f01526119110152611a6890f35b61006d565b60405190565b600080fd5b601f801991011690565b634e487b7160e01b600052604160045260246000fd5b9061009c90610072565b810190811060018060401b038211176100b457604052565b61007c565b906100cc6100c5610067565b9283610092565b565b600080fd5b60018060a01b031690565b6100e7906100d3565b90565b6100f3906100de565b90565b6100ff816100ea565b0361010657565b600080fd5b90505190610118826100f6565b565b90602082820312610134576101319160000161010b565b90565b6100ce565b610157611bc88038038061014c816100b9565b92833981019061011a565b90565b60805256fe6101806040526004361015610014575b610580565b61001f60003561005e565b80631806b5f21461005957806381119ab214610054578063862e4a0c1461004f5763a3c573eb0361000f5761054b565b610493565b6103ce565b610224565b60e01c90565b60405190565b600080fd5b600080fd5b60018060a01b031690565b61008890610074565b90565b6100948161007f565b0361009b57565b600080fd5b905035906100ad8261008b565b565b63ffffffff1690565b6100c1816100af565b036100c857565b600080fd5b905035906100da826100b8565b565b600080fd5b600080fd5b601f801991011690565b634e487b7160e01b600052604160045260246000fd5b90610110906100e6565b810190811067ffffffffffffffff82111761012a57604052565b6100f0565b9061014261013b610064565b9283610106565b565b67ffffffffffffffff81116101625761015e6020916100e6565b0190565b6100f0565b90826000939282370152565b9092919261018861018382610144565b61012f565b938185526020850190828401116101a4576101a292610167565b565b6100e1565b9080601f830112156101c7578160206101c493359101610173565b90565b6100dc565b91606083830312610219576101e482600085016100a0565b926101f283602083016100cd565b92604082013567ffffffffffffffff81116102145761021192016101a9565b90565b61006f565b61006a565b60000190565b6102386102323660046101cc565b916108b4565b610240610064565b8061024a8161021e565b0390f35b600080fd5b600080fd5b91906020800283011161026757565b610253565b90565b6102788161026c565b0361027f57565b600080fd5b905035906102918261026f565b565b90565b61029f81610293565b036102a657565b600080fd5b905035906102b882610296565b565b600080fd5b909182601f830112156102f95781359167ffffffffffffffff83116102f45760200192600183028401116102ef57565b610253565b6102ba565b6100dc565b919091610920818403126103c9576103198360008301610258565b92610328816104008401610258565b92610337826108008501610284565b926103468361082083016102ab565b926103558161084084016102ab565b926103648261086085016100cd565b926103738361088083016100a0565b92610382816108a084016100cd565b92610391826108c085016100a0565b926103a0836108e08301610284565b9261090082013567ffffffffffffffff81116103c4576103c092016102bf565b9091565b61006f565b61006a565b3461040c576103f66103e13660046102fe565b9a999099989198979297969396959495610ec3565b6103fe610064565b806104088161021e565b0390f35b61024e565b9160608383031261048e57600083013567ffffffffffffffff8111610489578261043c9185016101a9565b92602081013567ffffffffffffffff8111610484578361045d9183016101a9565b92604082013567ffffffffffffffff811161047f5761047c92016101a9565b90565b61006f565b61006f565b61006f565b61006a565b6104a76104a1366004610411565b916115e5565b6104af610064565b806104b98161021e565b0390f35b60009103126104c857565b61006a565b7f000000000000000000000000000000000000000000000000000000000000000090565b90565b61050861050361050d92610074565b6104f1565b610074565b90565b610519906104f4565b90565b61052590610510565b90565b6105319061051c565b9052565b919061054990600060208501940190610528565b565b3461057b5761055b3660046104bd565b6105776105666104cd565b61056e610064565b91829182610535565b0390f35b61024e565b600080fd5b60001c90565b90565b61059a61059f91610585565b61058b565b90565b6105ac905461058e565b90565b90565b6105be6105c391610585565b6105af565b90565b6105d090546105b2565b90565b60c01c90565b63ffffffff1690565b6105ee6105f3916105d3565b6105d9565b90565b61060090546105e2565b90565b60018060a01b031690565b61061a61061f91610585565b610603565b90565b61062c905461060e565b90565b600080fd5b60e01b90565b600091031261064557565b61006a565b50602090565b905090565b90565b61066190610293565b9052565b9061067281602093610658565b0190565b61068090546105b2565b90565b60010190565b6106a561069f6106988361064a565b8094610650565b91610655565b6000915b8383106106b65750505050565b6106d36106cd6001926106c885610676565b610665565b92610683565b920191906106a9565b6106e59061026c565b9052565b6106f290610293565b9052565b6106ff906100af565b9052565b61070c9061007f565b9052565b634e487b7160e01b600052602260045260246000fd5b9060016002830492168015610746575b602083101461074157565b610710565b91607f1691610736565b60209181520190565b600052602060002090565b906000929180549061077f61077883610726565b8094610750565b916001811690816000146107d8575060011461079b575b505050565b6107a89192939450610759565b916000925b8184106107c05750500190388080610796565b600181602092959395548486015201910192906107ad565b92949550505060ff1916825215156020020190388080610796565b989694929099979593916109208a019a60008b0161081091610689565b6104008a0161081e91610689565b610800890161082c916106dc565b610820880161083a916106e9565b6108408701610848916106e9565b6108608601610856916106f6565b610880850161086491610703565b6108a08401610872916106f6565b6108c0830161088091610703565b6108e0820161088e916106dc565b8082039061090001526108a091610764565b90565b6108ab610064565b3d6000823e3d90fd5b9091506108e07f000000000000000000000000000000000000000000000000000000000000000061051c565b9063f5efcd7960006020926108f560406105a2565b9561090060416105c6565b9161090b60426105c6565b919061091760436105f6565b906109226044610622565b9261092d60456105a2565b946046968b3b156109a657610940610064565b9c8d9b61094d8d9c610634565b8c5260048c019a61095d9b6107f3565b03815a6000948591f180156109a157610974575b50565b6109949060003d811161099a575b61098c8183610106565b81019061063a565b38610971565b503d610982565b6108a3565b61062f565b600190818003010490565b1b90565b919060086109d69102916109d0600019846109b6565b926109b6565b9181191691161790565b6109e990610293565b90565b6109f590610585565b90565b9190610a0e610a09610a16936109e0565b6109ec565b9083546109ba565b9055565b600090565b610a3191610a2b610a1a565b916109f8565b565b5b818110610a3f575050565b80610a4d6000600193610a1f565b01610a34565b1c90565b9091828110610a66575b505050565b610a84610a7e610a78610a8f956109ab565b926109ab565b92610655565b918201910190610a33565b388080610a61565b90680100000000000000008111610abd5781610ab5610abb9361064a565b90610a57565b565b6100f0565b50602090565b35610ad281610296565b90565b90565b610ae182610ac2565b9167ffffffffffffffff8311610b4557610b0f610b09600192610b048686610a97565b610ad5565b92610655565b92049160005b838110610b225750505050565b6001906020610b38610b3386610ac8565b6109ec565b9401938184015501610b15565b6100f0565b90610b5491610ad8565b565b60001b90565b90610b6960001991610b56565b9181191691161790565b610b87610b82610b8c9261026c565b6104f1565b61026c565b90565b90565b90610ba7610ba2610bae92610b73565b610b8f565b8254610b5c565b9055565b90610bc7610bc2610bce926109e0565b6109ec565b8254610b5c565b9055565b90610be163ffffffff91610b56565b9181191691161790565b610bff610bfa610c04926100af565b6104f1565b6100af565b90565b90565b90610c1f610c1a610c2692610beb565b610c07565b8254610bd2565b9055565b60201b90565b90610c46640100000000600160c01b0391610c2a565b9181191691161790565b610c5990610510565b90565b90565b90610c74610c6f610c7b92610c50565b610c5c565b8254610c30565b9055565b60c01b90565b90610c9763ffffffff60c01b91610c7f565b9181191691161790565b90610cb6610cb1610cbd92610beb565b610c07565b8254610c85565b9055565b90610cd260018060a01b0391610b56565b9181191691161790565b90610cf1610cec610cf892610c50565b610c5c565b8254610cc1565b9055565b5090565b601f602091010490565b9190610d20610d1b610d2893610b73565b610b8f565b9083546109ba565b9055565b600090565b610d4391610d3d610d2c565b91610d0a565b565b5b818110610d51575050565b80610d5f6000600193610d31565b01610d46565b9190601f8111610d75575b505050565b610d81610da693610759565b906020610d8d84610d00565b83019310610dae575b610d9f90610d00565b0190610d45565b388080610d70565b9150610d9f81929050610d96565b90610dcd9060001990600802610a53565b191690565b81610ddc91610dbc565b906002021790565b91610def9082610cfc565b9067ffffffffffffffff8211610eb157610e1382610e0d8554610726565b85610d65565b600090601f8311600114610e4857918091610e3793600092610e3c575b5050610dd2565b90555b565b90915001353880610e30565b601f19831691610e5785610759565b9260005b818110610e9957509160029391856001969410610e7f575b50505002019055610e3a565b610e8f910135601f841690610dbc565b9055388080610e73565b91936020600181928787013581550195019201610e5b565b6100f0565b90610ec19291610de4565b565b610f369895610f13610f2896610f0c610f3d9f9e9d9b97610f05610f2198610efe610f2f9e99610ef7610f1a9a6000610b4a565b6020610b4a565b6040610b92565b6041610bb2565b6042610bb2565b6043610c0a565b6043610c5f565b6043610ca1565b6044610cdc565b6045610b92565b6046610eb6565b7f9d226db03d4d6614ea01926ce8a588879492a2681b9684eb655b1470d32d4b9e610f66610064565b80610f708161021e565b0390a1565b67ffffffffffffffff8111610f8a5760200290565b6100f0565b90505190610f9c82610296565b565b90929192610fb3610fae82610f75565b61012f565b936020859202830192818411610feb57915b838310610fd25750505050565b60208091610fe08486610f8f565b815201920191610fc5565b610253565b9080601f8301121561100b5761100891602090610f9e565b90565b6100dc565b9050519061101d8261026f565b565b9050519061102c826100b8565b565b61103790610074565b90565b6110438161102e565b0361104a57565b600080fd5b9050519061105c8261103a565b565b60005b838110611072575050906000910152565b806020918301518185015201611061565b9092919261109861109382610144565b61012f565b938185526020850190828401116110b4576110b29261105e565b565b6100e1565b9080601f830112156110d7578160206110d493519101611083565b90565b6100dc565b919091610920818403126111a6576110f78360008301610ff0565b92611106816104008401610ff0565b92611115826108008501611010565b92611124836108208301610f8f565b92611133816108408401610f8f565b9261114282610860850161101f565b9261115183610880830161104f565b92611160816108a0840161101f565b9261116f826108c0850161104f565b9261117e836108e08301611010565b9261090082015167ffffffffffffffff81116111a15761119e92016110b9565b90565b61006f565b61006a565b5190565b6111b890610510565b90565b50602090565b90565b60200190565b6111e66111e06111d9836111bb565b8094610650565b916111c1565b6000915b8383106111f75750505050565b61120d6112076001928451610665565b926111c4565b920191906111ea565b61123561123e6020936112439361122c816111ab565b93848093610750565b9586910161105e565b6100e6565b0190565b989694929099979593916109208a019a60008b01611264916111ca565b6104008a01611272916111ca565b6108008901611280916106dc565b610820880161128e916106e9565b610840870161129c916106e9565b61086086016112aa916106f6565b61088085016112b891610703565b6108a084016112c6916106f6565b6108c083016112d491610703565b6108e082016112e2916106dc565b8082039061090001526112f491611216565b90565b90565b61130e611309611313926112f7565b6104f1565b6100af565b90565b61131f906112fa565b9052565b989694929099979593916109208a019a60008b01611340916111ca565b6104008a0161134e916111ca565b610800890161135c916106dc565b610820880161136a916106e9565b6108408701611378916106e9565b6108608601611386916106f6565b610880850161139491610703565b6108a084016113a291611316565b6108c083016113b091610703565b6108e082016113be916106dc565b8082039061090001526113d091611216565b90565b906113e56113e083610144565b61012f565b918252565b606090565b3d60001461140c576114003d6113d3565b903d6000602084013e5b565b6114146113ea565b9061140a565b151590565b60209181520190565b60007f44657374696e6174696f6e4e6574776f726b496e76616c696400000000000000910152565b61145d601960209261141f565b61146681611428565b0190565b6114809060208101906000818303910152611450565b90565b1561148a57565b611492610064565b62461bcd60e51b8152806114a86004820161146a565b0390fd5b6114b58161141a565b036114bc57565b600080fd5b905051906114ce826114ac565b565b9160c083830312611547576114e8826000850161101f565b926114f6836020830161104f565b926115048160408401611010565b92611512826060850161104f565b9261152083608083016114c1565b9260a082015167ffffffffffffffff81116115425761153f92016110b9565b90565b61006f565b61006a565b6115559061141a565b9052565b926115b796946115966115a09261158c6115aa9699959961158260c08a019b60008b01906106f6565b6020890190610703565b60408701906106dc565b6060850190610703565b608083019061154c565b60a0818403910152611216565b90565b156115c157565b600080fd5b906020828203126115e0576115dd9160000161101f565b90565b61006a565b6116359161160661162f9260206115fb826111ab565b8183010191016110dc565b60e09d979c9692959a949899939d5260c0526101605261010052610140529792949693946111af565b966111af565b6080526116617f000000000000000000000000000000000000000000000000000000000000000061051c565b61014051610100519085938561012052610160519089908b6080519260c0519460e05196893b15611a2d5760009a6116bf9961169b610064565b60a0526116ab63f5efcd79610634565b60a051526101205192600460a05101611247565b9060a0519160a0519003908360a051915af18015611a285761183a6117c1600096956117ab88806118659b60046117806020611771859e61182b9e6117cf9e6117d69e6119f2575b6117307f000000000000000000000000000000000000000000000000000000000000000061051c565b97610140519361010051939291610160519190916103e8936080519560c0519760e0519961175c610064565b9e8f9d8e0163f5efcd7960e01b815201611323565b60208201810382520382610106565b82602082019151925af16117926113ef565b506117a561179f8a61141a565b9161141a565b14611483565b60206117b6826111ab565b8183010191016114d0565b9893969197909294966111af565b93966111af565b90966118017f000000000000000000000000000000000000000000000000000000000000000061051c565b973496949792909192611812610064565b98899660046020890163cd58657960e01b815201611559565b60208201810382520383610106565b602082019151925af161184b6113ef565b5061185f611859600161141a565b9161141a565b146115ba565b6118a960206118937f000000000000000000000000000000000000000000000000000000000000000061051c565b63bab161bf906118a1610064565b938492610634565b825281806118b96004820161021e565b03915afa80156119ed5761190992611902926118eb926119c1575b5060206118e0826111ab565b8183010191016110dc565b9496999b9192959b9a90939a9897999b969a6111af565b94916111af565b9192936119357f000000000000000000000000000000000000000000000000000000000000000061051c565b9963f5efcd799897999b96909192939495968b3b156119bc57611956610064565b9c8d9b6119638d9c610634565b8c5260048c019a6119739b611247565b03815a6000948591f180156119b75761198a575b50565b6119aa9060003d81116119b0575b6119a28183610106565b81019061063a565b38611987565b503d611998565b6108a3565b61062f565b6119e19060203d81116119e6575b6119d98183610106565b8101906115c6565b6118d4565b503d6119cf565b6108a3565b883d8111611a1c575b80611a0b611a179260a051610106565b60a0510160a05161063a565b611707565b50611a173d90506119fb565b6108a3565b61062f56fea2646970667358221220e73d170973c57875b8606c7794888a2af5a60083169f7404124e74ab7c09cf4a64736f6c634300081c0033",
}

// BridgemessagereceivermockABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgemessagereceivermockMetaData.ABI instead.
var BridgemessagereceivermockABI = BridgemessagereceivermockMetaData.ABI

// BridgemessagereceivermockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgemessagereceivermockMetaData.Bin instead.
var BridgemessagereceivermockBin = BridgemessagereceivermockMetaData.Bin

// DeployBridgemessagereceivermock deploys a new Ethereum contract, binding an instance of Bridgemessagereceivermock to it.
func DeployBridgemessagereceivermock(auth *bind.TransactOpts, backend bind.ContractBackend, _bridgeAddress common.Address) (common.Address, *types.Transaction, *Bridgemessagereceivermock, error) {
	parsed, err := BridgemessagereceivermockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgemessagereceivermockBin), backend, _bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridgemessagereceivermock{BridgemessagereceivermockCaller: BridgemessagereceivermockCaller{contract: contract}, BridgemessagereceivermockTransactor: BridgemessagereceivermockTransactor{contract: contract}, BridgemessagereceivermockFilterer: BridgemessagereceivermockFilterer{contract: contract}}, nil
}

// Bridgemessagereceivermock is an auto generated Go binding around an Ethereum contract.
type Bridgemessagereceivermock struct {
	BridgemessagereceivermockCaller     // Read-only binding to the contract
	BridgemessagereceivermockTransactor // Write-only binding to the contract
	BridgemessagereceivermockFilterer   // Log filterer for contract events
}

// BridgemessagereceivermockCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgemessagereceivermockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgemessagereceivermockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgemessagereceivermockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgemessagereceivermockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgemessagereceivermockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgemessagereceivermockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgemessagereceivermockSession struct {
	Contract     *Bridgemessagereceivermock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BridgemessagereceivermockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgemessagereceivermockCallerSession struct {
	Contract *BridgemessagereceivermockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// BridgemessagereceivermockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgemessagereceivermockTransactorSession struct {
	Contract     *BridgemessagereceivermockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// BridgemessagereceivermockRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgemessagereceivermockRaw struct {
	Contract *Bridgemessagereceivermock // Generic contract binding to access the raw methods on
}

// BridgemessagereceivermockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgemessagereceivermockCallerRaw struct {
	Contract *BridgemessagereceivermockCaller // Generic read-only contract binding to access the raw methods on
}

// BridgemessagereceivermockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgemessagereceivermockTransactorRaw struct {
	Contract *BridgemessagereceivermockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgemessagereceivermock creates a new instance of Bridgemessagereceivermock, bound to a specific deployed contract.
func NewBridgemessagereceivermock(address common.Address, backend bind.ContractBackend) (*Bridgemessagereceivermock, error) {
	contract, err := bindBridgemessagereceivermock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridgemessagereceivermock{BridgemessagereceivermockCaller: BridgemessagereceivermockCaller{contract: contract}, BridgemessagereceivermockTransactor: BridgemessagereceivermockTransactor{contract: contract}, BridgemessagereceivermockFilterer: BridgemessagereceivermockFilterer{contract: contract}}, nil
}

// NewBridgemessagereceivermockCaller creates a new read-only instance of Bridgemessagereceivermock, bound to a specific deployed contract.
func NewBridgemessagereceivermockCaller(address common.Address, caller bind.ContractCaller) (*BridgemessagereceivermockCaller, error) {
	contract, err := bindBridgemessagereceivermock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgemessagereceivermockCaller{contract: contract}, nil
}

// NewBridgemessagereceivermockTransactor creates a new write-only instance of Bridgemessagereceivermock, bound to a specific deployed contract.
func NewBridgemessagereceivermockTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgemessagereceivermockTransactor, error) {
	contract, err := bindBridgemessagereceivermock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgemessagereceivermockTransactor{contract: contract}, nil
}

// NewBridgemessagereceivermockFilterer creates a new log filterer instance of Bridgemessagereceivermock, bound to a specific deployed contract.
func NewBridgemessagereceivermockFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgemessagereceivermockFilterer, error) {
	contract, err := bindBridgemessagereceivermock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgemessagereceivermockFilterer{contract: contract}, nil
}

// bindBridgemessagereceivermock binds a generic wrapper to an already deployed contract.
func bindBridgemessagereceivermock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgemessagereceivermockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridgemessagereceivermock *BridgemessagereceivermockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridgemessagereceivermock.Contract.BridgemessagereceivermockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridgemessagereceivermock *BridgemessagereceivermockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgemessagereceivermock.Contract.BridgemessagereceivermockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridgemessagereceivermock *BridgemessagereceivermockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridgemessagereceivermock.Contract.BridgemessagereceivermockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridgemessagereceivermock *BridgemessagereceivermockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridgemessagereceivermock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridgemessagereceivermock *BridgemessagereceivermockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgemessagereceivermock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridgemessagereceivermock *BridgemessagereceivermockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridgemessagereceivermock.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Bridgemessagereceivermock *BridgemessagereceivermockCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgemessagereceivermock.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Bridgemessagereceivermock *BridgemessagereceivermockSession) BridgeAddress() (common.Address, error) {
	return _Bridgemessagereceivermock.Contract.BridgeAddress(&_Bridgemessagereceivermock.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Bridgemessagereceivermock *BridgemessagereceivermockCallerSession) BridgeAddress() (common.Address, error) {
	return _Bridgemessagereceivermock.Contract.BridgeAddress(&_Bridgemessagereceivermock.CallOpts)
}

// OnMessageReceived is a paid mutator transaction binding the contract method 0x1806b5f2.
//
// Solidity: function onMessageReceived(address originAddress, uint32 originNetwork, bytes data) payable returns()
func (_Bridgemessagereceivermock *BridgemessagereceivermockTransactor) OnMessageReceived(opts *bind.TransactOpts, originAddress common.Address, originNetwork uint32, data []byte) (*types.Transaction, error) {
	return _Bridgemessagereceivermock.contract.Transact(opts, "onMessageReceived", originAddress, originNetwork, data)
}

// OnMessageReceived is a paid mutator transaction binding the contract method 0x1806b5f2.
//
// Solidity: function onMessageReceived(address originAddress, uint32 originNetwork, bytes data) payable returns()
func (_Bridgemessagereceivermock *BridgemessagereceivermockSession) OnMessageReceived(originAddress common.Address, originNetwork uint32, data []byte) (*types.Transaction, error) {
	return _Bridgemessagereceivermock.Contract.OnMessageReceived(&_Bridgemessagereceivermock.TransactOpts, originAddress, originNetwork, data)
}

// OnMessageReceived is a paid mutator transaction binding the contract method 0x1806b5f2.
//
// Solidity: function onMessageReceived(address originAddress, uint32 originNetwork, bytes data) payable returns()
func (_Bridgemessagereceivermock *BridgemessagereceivermockTransactorSession) OnMessageReceived(originAddress common.Address, originNetwork uint32, data []byte) (*types.Transaction, error) {
	return _Bridgemessagereceivermock.Contract.OnMessageReceived(&_Bridgemessagereceivermock.TransactOpts, originAddress, originNetwork, data)
}

// TestClaim is a paid mutator transaction binding the contract method 0x862e4a0c.
//
// Solidity: function testClaim(bytes claimData1, bytes bridgeAsset, bytes claimData2) payable returns()
func (_Bridgemessagereceivermock *BridgemessagereceivermockTransactor) TestClaim(opts *bind.TransactOpts, claimData1 []byte, bridgeAsset []byte, claimData2 []byte) (*types.Transaction, error) {
	return _Bridgemessagereceivermock.contract.Transact(opts, "testClaim", claimData1, bridgeAsset, claimData2)
}

// TestClaim is a paid mutator transaction binding the contract method 0x862e4a0c.
//
// Solidity: function testClaim(bytes claimData1, bytes bridgeAsset, bytes claimData2) payable returns()
func (_Bridgemessagereceivermock *BridgemessagereceivermockSession) TestClaim(claimData1 []byte, bridgeAsset []byte, claimData2 []byte) (*types.Transaction, error) {
	return _Bridgemessagereceivermock.Contract.TestClaim(&_Bridgemessagereceivermock.TransactOpts, claimData1, bridgeAsset, claimData2)
}

// TestClaim is a paid mutator transaction binding the contract method 0x862e4a0c.
//
// Solidity: function testClaim(bytes claimData1, bytes bridgeAsset, bytes claimData2) payable returns()
func (_Bridgemessagereceivermock *BridgemessagereceivermockTransactorSession) TestClaim(claimData1 []byte, bridgeAsset []byte, claimData2 []byte) (*types.Transaction, error) {
	return _Bridgemessagereceivermock.Contract.TestClaim(&_Bridgemessagereceivermock.TransactOpts, claimData1, bridgeAsset, claimData2)
}

// UpdateParameters is a paid mutator transaction binding the contract method 0x81119ab2.
//
// Solidity: function updateParameters(bytes32[32] msmtProofLocalExitRoot, bytes32[32] msmtProofRollupExitRoot, uint256 mglobalIndex, bytes32 mmainnetExitRoot, bytes32 mrollupExitRoot, uint32 moriginNetwork, address moriginAddress, uint32 mdestinationNetwork, address mdestinationAddress, uint256 mamount, bytes mmetadata) returns()
func (_Bridgemessagereceivermock *BridgemessagereceivermockTransactor) UpdateParameters(opts *bind.TransactOpts, msmtProofLocalExitRoot [32][32]byte, msmtProofRollupExitRoot [32][32]byte, mglobalIndex *big.Int, mmainnetExitRoot [32]byte, mrollupExitRoot [32]byte, moriginNetwork uint32, moriginAddress common.Address, mdestinationNetwork uint32, mdestinationAddress common.Address, mamount *big.Int, mmetadata []byte) (*types.Transaction, error) {
	return _Bridgemessagereceivermock.contract.Transact(opts, "updateParameters", msmtProofLocalExitRoot, msmtProofRollupExitRoot, mglobalIndex, mmainnetExitRoot, mrollupExitRoot, moriginNetwork, moriginAddress, mdestinationNetwork, mdestinationAddress, mamount, mmetadata)
}

// UpdateParameters is a paid mutator transaction binding the contract method 0x81119ab2.
//
// Solidity: function updateParameters(bytes32[32] msmtProofLocalExitRoot, bytes32[32] msmtProofRollupExitRoot, uint256 mglobalIndex, bytes32 mmainnetExitRoot, bytes32 mrollupExitRoot, uint32 moriginNetwork, address moriginAddress, uint32 mdestinationNetwork, address mdestinationAddress, uint256 mamount, bytes mmetadata) returns()
func (_Bridgemessagereceivermock *BridgemessagereceivermockSession) UpdateParameters(msmtProofLocalExitRoot [32][32]byte, msmtProofRollupExitRoot [32][32]byte, mglobalIndex *big.Int, mmainnetExitRoot [32]byte, mrollupExitRoot [32]byte, moriginNetwork uint32, moriginAddress common.Address, mdestinationNetwork uint32, mdestinationAddress common.Address, mamount *big.Int, mmetadata []byte) (*types.Transaction, error) {
	return _Bridgemessagereceivermock.Contract.UpdateParameters(&_Bridgemessagereceivermock.TransactOpts, msmtProofLocalExitRoot, msmtProofRollupExitRoot, mglobalIndex, mmainnetExitRoot, mrollupExitRoot, moriginNetwork, moriginAddress, mdestinationNetwork, mdestinationAddress, mamount, mmetadata)
}

// UpdateParameters is a paid mutator transaction binding the contract method 0x81119ab2.
//
// Solidity: function updateParameters(bytes32[32] msmtProofLocalExitRoot, bytes32[32] msmtProofRollupExitRoot, uint256 mglobalIndex, bytes32 mmainnetExitRoot, bytes32 mrollupExitRoot, uint32 moriginNetwork, address moriginAddress, uint32 mdestinationNetwork, address mdestinationAddress, uint256 mamount, bytes mmetadata) returns()
func (_Bridgemessagereceivermock *BridgemessagereceivermockTransactorSession) UpdateParameters(msmtProofLocalExitRoot [32][32]byte, msmtProofRollupExitRoot [32][32]byte, mglobalIndex *big.Int, mmainnetExitRoot [32]byte, mrollupExitRoot [32]byte, moriginNetwork uint32, moriginAddress common.Address, mdestinationNetwork uint32, mdestinationAddress common.Address, mamount *big.Int, mmetadata []byte) (*types.Transaction, error) {
	return _Bridgemessagereceivermock.Contract.UpdateParameters(&_Bridgemessagereceivermock.TransactOpts, msmtProofLocalExitRoot, msmtProofRollupExitRoot, mglobalIndex, mmainnetExitRoot, mrollupExitRoot, moriginNetwork, moriginAddress, mdestinationNetwork, mdestinationAddress, mamount, mmetadata)
}

// BridgemessagereceivermockMessageReceivedIterator is returned from FilterMessageReceived and is used to iterate over the raw logs and unpacked data for MessageReceived events raised by the Bridgemessagereceivermock contract.
type BridgemessagereceivermockMessageReceivedIterator struct {
	Event *BridgemessagereceivermockMessageReceived // Event containing the contract specifics and raw log

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
func (it *BridgemessagereceivermockMessageReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgemessagereceivermockMessageReceived)
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
		it.Event = new(BridgemessagereceivermockMessageReceived)
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
func (it *BridgemessagereceivermockMessageReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgemessagereceivermockMessageReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgemessagereceivermockMessageReceived represents a MessageReceived event raised by the Bridgemessagereceivermock contract.
type BridgemessagereceivermockMessageReceived struct {
	DestinationAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMessageReceived is a free log retrieval operation binding the contract event 0xdf9f4a3ac608a3edf2b45dafa2b30a40073df2a24c06756d4a68210b7de0a8b8.
//
// Solidity: event MessageReceived(address destinationAddress)
func (_Bridgemessagereceivermock *BridgemessagereceivermockFilterer) FilterMessageReceived(opts *bind.FilterOpts) (*BridgemessagereceivermockMessageReceivedIterator, error) {

	logs, sub, err := _Bridgemessagereceivermock.contract.FilterLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return &BridgemessagereceivermockMessageReceivedIterator{contract: _Bridgemessagereceivermock.contract, event: "MessageReceived", logs: logs, sub: sub}, nil
}

// WatchMessageReceived is a free log subscription operation binding the contract event 0xdf9f4a3ac608a3edf2b45dafa2b30a40073df2a24c06756d4a68210b7de0a8b8.
//
// Solidity: event MessageReceived(address destinationAddress)
func (_Bridgemessagereceivermock *BridgemessagereceivermockFilterer) WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *BridgemessagereceivermockMessageReceived) (event.Subscription, error) {

	logs, sub, err := _Bridgemessagereceivermock.contract.WatchLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgemessagereceivermockMessageReceived)
				if err := _Bridgemessagereceivermock.contract.UnpackLog(event, "MessageReceived", log); err != nil {
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

// ParseMessageReceived is a log parse operation binding the contract event 0xdf9f4a3ac608a3edf2b45dafa2b30a40073df2a24c06756d4a68210b7de0a8b8.
//
// Solidity: event MessageReceived(address destinationAddress)
func (_Bridgemessagereceivermock *BridgemessagereceivermockFilterer) ParseMessageReceived(log types.Log) (*BridgemessagereceivermockMessageReceived, error) {
	event := new(BridgemessagereceivermockMessageReceived)
	if err := _Bridgemessagereceivermock.contract.UnpackLog(event, "MessageReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgemessagereceivermockUpdateParametersIterator is returned from FilterUpdateParameters and is used to iterate over the raw logs and unpacked data for UpdateParameters events raised by the Bridgemessagereceivermock contract.
type BridgemessagereceivermockUpdateParametersIterator struct {
	Event *BridgemessagereceivermockUpdateParameters // Event containing the contract specifics and raw log

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
func (it *BridgemessagereceivermockUpdateParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgemessagereceivermockUpdateParameters)
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
		it.Event = new(BridgemessagereceivermockUpdateParameters)
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
func (it *BridgemessagereceivermockUpdateParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgemessagereceivermockUpdateParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgemessagereceivermockUpdateParameters represents a UpdateParameters event raised by the Bridgemessagereceivermock contract.
type BridgemessagereceivermockUpdateParameters struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpdateParameters is a free log retrieval operation binding the contract event 0x9d226db03d4d6614ea01926ce8a588879492a2681b9684eb655b1470d32d4b9e.
//
// Solidity: event UpdateParameters()
func (_Bridgemessagereceivermock *BridgemessagereceivermockFilterer) FilterUpdateParameters(opts *bind.FilterOpts) (*BridgemessagereceivermockUpdateParametersIterator, error) {

	logs, sub, err := _Bridgemessagereceivermock.contract.FilterLogs(opts, "UpdateParameters")
	if err != nil {
		return nil, err
	}
	return &BridgemessagereceivermockUpdateParametersIterator{contract: _Bridgemessagereceivermock.contract, event: "UpdateParameters", logs: logs, sub: sub}, nil
}

// WatchUpdateParameters is a free log subscription operation binding the contract event 0x9d226db03d4d6614ea01926ce8a588879492a2681b9684eb655b1470d32d4b9e.
//
// Solidity: event UpdateParameters()
func (_Bridgemessagereceivermock *BridgemessagereceivermockFilterer) WatchUpdateParameters(opts *bind.WatchOpts, sink chan<- *BridgemessagereceivermockUpdateParameters) (event.Subscription, error) {

	logs, sub, err := _Bridgemessagereceivermock.contract.WatchLogs(opts, "UpdateParameters")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgemessagereceivermockUpdateParameters)
				if err := _Bridgemessagereceivermock.contract.UnpackLog(event, "UpdateParameters", log); err != nil {
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

// ParseUpdateParameters is a log parse operation binding the contract event 0x9d226db03d4d6614ea01926ce8a588879492a2681b9684eb655b1470d32d4b9e.
//
// Solidity: event UpdateParameters()
func (_Bridgemessagereceivermock *BridgemessagereceivermockFilterer) ParseUpdateParameters(log types.Log) (*BridgemessagereceivermockUpdateParameters, error) {
	event := new(BridgemessagereceivermockUpdateParameters)
	if err := _Bridgemessagereceivermock.contract.UnpackLog(event, "UpdateParameters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

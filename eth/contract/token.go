// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FCTokenABI is the input ABI used to generate the binding from.
const FCTokenABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_initiSuppy\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"transcationId\",\"type\":\"uint8\"}],\"name\":\"CreateTranscation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accountsMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowers\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8[]\",\"name\":\"b\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256\",\"name\":\"len\",\"type\":\"uint256\"}],\"name\":\"approveWithArray\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"}],\"name\":\"createMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"getCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingTranscation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"multiSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proportion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"}],\"name\":\"removeAllower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"}],\"name\":\"removeSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"}],\"name\":\"setAllower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"setPropertion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"startTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transcationId\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transcations\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"signatureCounts\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FCTokenFuncSigs maps the 4-byte function signature to its string representation.
var FCTokenFuncSigs = map[string]string{
	"f2a40db8": "accounts(uint256)",
	"e946cfd4": "accountsMap(address)",
	"dd62ed3e": "allowance(address,address)",
	"1de4d19d": "allowers(address)",
	"095ea7b3": "approve(address,uint256)",
	"275334ae": "approveWithArray(uint8[],uint256)",
	"70a08231": "balanceOf(address)",
	"a6f2ae3a": "buy()",
	"8f75c003": "createMiner(address,string)",
	"313ce567": "decimals()",
	"a98e4e77": "getAccountCount()",
	"12065fe0": "getBalance()",
	"615715a2": "getCounts(uint8)",
	"519ed728": "getPendingTranscation()",
	"3458810a": "multiSignature(uint8)",
	"06fdde03": "name()",
	"1c53c280": "nodes(uint256)",
	"8da5cb5b": "owner()",
	"5b4246d4": "proportion()",
	"b869abf7": "removeAllower(address)",
	"0e316ab7": "removeSigner(address)",
	"4679a9e1": "setAllower(address)",
	"675dcedf": "setPropertion(uint256)",
	"6c19e783": "setSigner(address)",
	"0536b723": "signer(address)",
	"604374f0": "startTransfer(address,uint256)",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"87019cd4": "transcationId()",
	"78e8691d": "transcations(uint256)",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// FCTokenBin is the compiled bytecode used for deploying new contracts.
var FCTokenBin = "0x60806040523480156200001157600080fd5b5060405162001f7338038062001f7383398101604081905262000034916200021f565b835162000049906000906020870190620000c6565b5082516200005f906001906020860190620000c6565b506002805460ff191660ff84169081179091556200007f90600a620002f3565b6200008b9082620003e8565b600381905533600081815260066020526040902091909155600480546001600160a01b03191690911790555050606460055550620004739050565b828054620000d4906200040a565b90600052602060002090601f016020900481019282620000f8576000855562000143565b82601f106200011357805160ff191683800117855562000143565b8280016001018555821562000143579182015b828111156200014357825182559160200191906001019062000126565b506200015192915062000155565b5090565b5b8082111562000151576000815560010162000156565b600082601f8301126200017d578081fd5b81516001600160401b03808211156200019a576200019a6200045d565b604051601f8301601f19908116603f01168101908282118183101715620001c557620001c56200045d565b81604052838152602092508683858801011115620001e1578485fd5b8491505b83821015620002045785820183015181830184015290820190620001e5565b838211156200021557848385830101525b9695505050505050565b6000806000806080858703121562000235578384fd5b84516001600160401b03808211156200024c578586fd5b6200025a888389016200016c565b9550602087015191508082111562000270578485fd5b506200027f878288016200016c565b935050604085015160ff8116811462000296578283fd5b6060959095015193969295505050565b80825b6001808611620002ba5750620002ea565b818704821115620002cf57620002cf62000447565b80861615620002dd57918102915b9490941c938002620002a9565b94509492505050565b60006200030460001984846200030b565b9392505050565b6000826200031c5750600162000304565b816200032b5750600062000304565b81600181146200034457600281146200034f5762000383565b600191505062000304565b60ff84111562000363576200036362000447565b6001841b9150848211156200037c576200037c62000447565b5062000304565b5060208310610133831016604e8410600b8410161715620003bb575081810a83811115620003b557620003b562000447565b62000304565b620003ca8484846001620002a6565b808604821115620003df57620003df62000447565b02949350505050565b600081600019048311821515161562000405576200040562000447565b500290565b600181811c908216806200041f57607f821691505b602082108114156200044157634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b611af080620004836000396000f3fe6080604052600436106101e35760003560e01c8063615715a21161010257806395d89b4111610095578063b869abf711610064578063b869abf7146105db578063dd62ed3e146105fb578063e946cfd414610633578063f2a40db814610673576101e3565b806395d89b4114610589578063a6f2ae3a1461059e578063a9059cbb146105a6578063a98e4e77146105c6576101e3565b806378e8691d116100d157806378e8691d1461049757806387019cd4146105175780638da5cb5b146105315780638f75c00314610569576101e3565b8063615715a21461040a578063675dcedf1461042a5780636c19e7831461044a57806370a082311461046a576101e3565b806323b872dd1161017a5780634679a9e1116101495780634679a9e11461039f578063519ed728146103bf5780635b4246d4146103d4578063604374f0146103ea576101e3565b806323b872dd14610325578063275334ae14610345578063313ce567146103655780633458810a1461037f576101e3565b806312065fe0116101b657806312065fe01461029357806318160ddd146102bf5780631c53c280146102d55780631de4d19d146102f5576101e3565b80630536b723146101e857806306fdde031461022f578063095ea7b3146102515780630e316ab714610273575b600080fd5b3480156101f457600080fd5b5061021861020336600461157a565b60086020526000908152604090205460ff1681565b60405160ff90911681526020015b60405180910390f35b34801561023b57600080fd5b50610244610693565b60405161022691906117b8565b34801561025d57600080fd5b5061027161026c3660046116a8565b610721565b005b34801561027f57600080fd5b5061027161028e36600461157a565b6107fd565b34801561029f57600080fd5b50336000908152600660205260409020545b604051908152602001610226565b3480156102cb57600080fd5b506102b160035481565b3480156102e157600080fd5b506102446102f0366004611786565b610835565b34801561030157600080fd5b5061021861031036600461157a565b60096020526000908152604090205460ff1681565b34801561033157600080fd5b506102716103403660046115cd565b610860565b34801561035157600080fd5b506102716103603660046116d1565b61093c565b34801561037157600080fd5b506002546102189060ff1681565b34801561038b57600080fd5b5061027161039a36600461179e565b610ce1565b3480156103ab57600080fd5b506102716103ba36600461157a565b610e5f565b3480156103cb57600080fd5b50600d546102b1565b3480156103e057600080fd5b506102b160055481565b3480156103f657600080fd5b506102716104053660046116a8565b610e9a565b34801561041657600080fd5b506102b161042536600461179e565b610fe8565b34801561043657600080fd5b50610271610445366004611786565b611006565b34801561045657600080fd5b5061027161046536600461157a565b611037565b34801561047657600080fd5b506102b161048536600461157a565b60066020526000908152604090205481565b3480156104a357600080fd5b506104e96104b2366004611786565b600f6020526000908152604090208054600182015460028301546003909301546001600160a01b0392831693919092169160ff1684565b604080516001600160a01b0395861681529490931660208501529183015260ff166060820152608001610226565b34801561052357600080fd5b50600e546102189060ff1681565b34801561053d57600080fd5b50600454610551906001600160a01b031681565b6040516001600160a01b039091168152602001610226565b34801561057557600080fd5b50610271610584366004611608565b611072565b34801561059557600080fd5b50610244611166565b6102b1611173565b3480156105b257600080fd5b506102716105c13660046116a8565b61127c565b3480156105d257600080fd5b50600a546102b1565b3480156105e757600080fd5b506102716105f636600461157a565b611287565b34801561060757600080fd5b506102b161061636600461159b565b600760209081526000928352604080842090915290825290205481565b34801561063f57600080fd5b5061066361064e36600461157a565b600c6020526000908152604090205460ff1681565b6040519015158152602001610226565b34801561067f57600080fd5b5061055161068e366004611786565b6112bf565b600080546106a0906119ee565b80601f01602080910402602001604051908101604052809291908181526020018280546106cc906119ee565b80156107195780601f106106ee57610100808354040283529160200191610719565b820191906000526020600020905b8154815290600101906020018083116106fc57829003601f168201915b505050505081565b3360009081526009602052604090205460ff1660011461074057600080fd5b6001600160a01b03821661075357600080fd5b6004546001600160a01b031660009081526006602052604090205481111561077a57600080fd5b6004546001600160a01b039081166000908152600760209081526040808320938616835292905290812080548392906107b490849061183c565b90915550506040518181526001600160a01b0383169033907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259060200160405180910390a35050565b6004546001600160a01b0316331461081457600080fd5b6001600160a01b03166000908152600860205260409020805460ff19169055565b600b818154811061084557600080fd5b9060005260206000200160009150905080546106a0906119ee565b6001600160a01b03821661087357600080fd5b6004546001600160a01b031660009081526007602090815260408083203384529091529020548111156108a557600080fd5b6001600160a01b0383166000908152600760209081526040808320338452909152812080548392906108d89084906119c0565b90915550506001600160a01b038316600090815260066020526040812080548392906109059084906119c0565b90915550506001600160a01b0382166000908152600660205260408120805483929061093290849061183c565b9091555050505050565b600082511161094a57600080fd5b3360009081526009602052604090205460ff1660011461096957600080fd5b6000811161097657600080fd5b60008167ffffffffffffffff81111561099f57634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156109c8578160200160208202803683370190505b5083519091506000905b8015610ab8576000856109e66001846119c0565b81518110610a0457634e487b7160e01b600052603260045260246000fd5b6020026020010151905060005b60088160ff161015610aa35760008160ff166001901b8360ff16161115610a91578060ff16838851610a4391906119c0565b610a4e9060086119a1565b610a58919061183c565b858581518110610a7857634e487b7160e01b600052603260045260246000fd5b602090810291909101015283610a8d81611a29565b9450505b80610a9b81611a44565b915050610a11565b50508080610ab0906119d7565b9150506109d2565b508151600a90811115610ac9575081515b60008167ffffffffffffffff811115610af257634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610b1b578160200160208202803683370190505b50905060005b82811015610cd85760004442604051602001610b47929190918252602082015260400190565b6040516020818303038152906040528051906020012060001c90506000865182610b719190611a64565b90506000878281518110610b9557634e487b7160e01b600052603260045260246000fd5b602002602001015190506000805b85811015610bf55782878281518110610bcc57634e487b7160e01b600052603260045260246000fd5b60200260200101511415610be35760019150610bf5565b80610bed81611a29565b915050610ba3565b5080610ccf5781868681518110610c1c57634e487b7160e01b600052603260045260246000fd5b6020026020010181815250506000600a8381548110610c4b57634e487b7160e01b600052603260045260246000fd5b6000918252602090912001546002546001600160a01b039091169150610c759060ff16600a6118d3565b610c809060016119a1565b6004546001600160a01b03908116600090815260076020908152604080832093861683529290529081208054909190610cba90849061183c565b90915550869050610cca81611a29565b965050505b50505050610b21565b50505050505050565b6004546001600160a01b0316331480610d0c57503360009081526008602052604090205460ff166001145b610d1557600080fd5b60ff81166000908152600f6020526040902080546001600160a01b0316610d3b57600080fd5b80546001600160a01b0316331415610d5257600080fd5b33600090815260048201602052604090205460ff1660011415610d7457600080fd5b3360009081526004820160205260408120805460ff1916600117905560038201805460ff1691610da383611a44565b82546101009290920a60ff818102199093169183160217909155600383015460029116109050610e5b576002810154306000908152600660205260409020541015610ded57600080fd5b60018101546002820154610e0e9130916001600160a01b03909116906112e9565b610e1a600d6000611491565b60ff82166000908152600f6020526040812080546001600160a01b0319908116825560018201805490911690556002810191909155600301805460ff191690555b5050565b6004546001600160a01b03163314610e7657600080fd5b6001600160a01b03166000908152600960205260409020805460ff19166001179055565b6001600160a01b038216610ead57600080fd5b33600090815260066020526040902054811115610ec957600080fd5b610ed43330836112e9565b600e54610ee59060ff166001611854565b600e805460ff92831660ff19918216811783556000908152600f6020908152604080832080546001600160a01b0319908116339081178355600183810180546001600160a01b038e16941684179055600284018b90556003840180549098169097558754600d805498890181559096528487047fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb5018054968a16601f9098166101000a978802978a02199096169690961790945594548151948552918401929092529082018590529092166060830152907fe1dd58c97af942b2e36ca0fc977e17197d6a8008fda5381780e291337e8ac9959060800160405180910390a1505050565b60ff8082166000908152600f6020526040902060030154165b919050565b3360009081526009602052604090205460ff1660011461102557600080fd5b6000811161103257600080fd5b600555565b6004546001600160a01b0316331461104e57600080fd5b6001600160a01b03166000908152600860205260409020805460ff19166001179055565b6001600160a01b03821661108557600080fd5b805160341461109357600080fd5b6001600160a01b0382166000908152600c602052604090205460ff1680151560011461116157600b805460018101825560009190915282516110fc917f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db9019060208501906114b9565b50600a805460018082019092557fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a80180546001600160a01b0319166001600160a01b0386169081179091556000908152600c60205260409020805460ff191690911790555b505050565b600180546106a0906119ee565b6005546002546000918291670de0b6b3a764000091906111979060ff16600a6118d3565b6111a190346119a1565b6111ab91906119a1565b6111b59190611879565b6004546001600160a01b031660009081526006602052604090205490915081908111156111e157600080fd5b6004546001600160a01b03166000908152600660205260408120805483929061120b9084906119c0565b9091555050336000908152600660205260408120805483929061122f90849061183c565b909155505060045460405182815233916001600160a01b0316907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a391505090565b610e5b3383836112e9565b6004546001600160a01b0316331461129e57600080fd5b6001600160a01b03166000908152600960205260409020805460ff19169055565b600a81815481106112cf57600080fd5b6000918252602090912001546001600160a01b0316905081565b6004546001600160a01b031633141561130157600080fd5b6001600160a01b03831660009081526006602052604090205481111561132657600080fd5b6001600160a01b03821661133957600080fd5b6001600160a01b03821660009081526006602052604090205461135c828261183c565b1161136657600080fd5b6001600160a01b0380831660009081526006602052604080822054928616825281205490916113949161183c565b6001600160a01b0385166000908152600660205260408120805492935084929091906113c19084906119c0565b90915550506001600160a01b038316600090815260066020526040812080548492906113ee90849061183c565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161143a91815260200190565b60405180910390a36001600160a01b0380841660009081526006602052604080822054928716825290205482916114709161183c565b1461148b57634e487b7160e01b600052600160045260246000fd5b50505050565b50805460008255601f0160209004906000526020600020908101906114b6919061153d565b50565b8280546114c5906119ee565b90600052602060002090601f0160209004810192826114e7576000855561152d565b82601f1061150057805160ff191683800117855561152d565b8280016001018555821561152d579182015b8281111561152d578251825591602001919060010190611512565b5061153992915061153d565b5090565b5b80821115611539576000815560010161153e565b80356001600160a01b038116811461100157600080fd5b803560ff8116811461100157600080fd5b60006020828403121561158b578081fd5b61159482611552565b9392505050565b600080604083850312156115ad578081fd5b6115b683611552565b91506115c460208401611552565b90509250929050565b6000806000606084860312156115e1578081fd5b6115ea84611552565b92506115f860208501611552565b9150604084013590509250925092565b6000806040838503121561161a578182fd5b61162383611552565b915060208084013567ffffffffffffffff80821115611640578384fd5b818601915086601f830112611653578384fd5b81358181111561166557611665611aa4565b611677601f8201601f1916850161180b565b9150808252878482850101111561168c578485fd5b8084840185840137810190920192909252919491935090915050565b600080604083850312156116ba578182fd5b6116c383611552565b946020939093013593505050565b600080604083850312156116e3578182fd5b823567ffffffffffffffff808211156116fa578384fd5b818501915085601f83011261170d578384fd5b813560208282111561172157611721611aa4565b8160051b925061173281840161180b565b8281528181019085830185870184018b101561174c578889fd5b8896505b848710156117755761176181611569565b835260019690960195918301918301611750565b509997909101359750505050505050565b600060208284031215611797578081fd5b5035919050565b6000602082840312156117af578081fd5b61159482611569565b6000602080835283518082850152825b818110156117e4578581018301518582016040015282016117c8565b818111156117f55783604083870101525b50601f01601f1916929092016040019392505050565b604051601f8201601f1916810167ffffffffffffffff8111828210171561183457611834611aa4565b604052919050565b6000821982111561184f5761184f611a78565b500190565b600060ff821660ff84168060ff0382111561187157611871611a78565b019392505050565b60008261188857611888611a8e565b500490565b80825b600180861161189f57506118ca565b8187048211156118b1576118b1611a78565b808616156118be57918102915b9490941c938002611890565b94509492505050565b600061159460001984846000826118ec57506001611594565b816118f957506000611594565b816001811461190f576002811461191957611946565b6001915050611594565b60ff84111561192a5761192a611a78565b6001841b91508482111561194057611940611a78565b50611594565b5060208310610133831016604e8410600b8410161715611979575081810a8381111561197457611974611a78565b611594565b611986848484600161188d565b80860482111561199857611998611a78565b02949350505050565b60008160001904831182151516156119bb576119bb611a78565b500290565b6000828210156119d2576119d2611a78565b500390565b6000816119e6576119e6611a78565b506000190190565b600181811c90821680611a0257607f821691505b60208210811415611a2357634e487b7160e01b600052602260045260246000fd5b50919050565b6000600019821415611a3d57611a3d611a78565b5060010190565b600060ff821660ff811415611a5b57611a5b611a78565b60010192915050565b600082611a7357611a73611a8e565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052604160045260246000fdfea2646970667358221220a5ef1ff4eacd19028cacdce49fe73c020c73f317e30bbad8547bd97f5dc10fab64736f6c63430008030033"

// DeployFCToken deploys a new Ethereum contract, binding an instance of FCToken to it.
func DeployFCToken(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, _decimals uint8, _initiSuppy *big.Int) (common.Address, *types.Transaction, *FCToken, error) {
	parsed, err := abi.JSON(strings.NewReader(FCTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FCTokenBin), backend, _name, _symbol, _decimals, _initiSuppy)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FCToken{FCTokenCaller: FCTokenCaller{contract: contract}, FCTokenTransactor: FCTokenTransactor{contract: contract}, FCTokenFilterer: FCTokenFilterer{contract: contract}}, nil
}

// FCToken is an auto generated Go binding around an Ethereum contract.
type FCToken struct {
	FCTokenCaller     // Read-only binding to the contract
	FCTokenTransactor // Write-only binding to the contract
	FCTokenFilterer   // Log filterer for contract events
}

// FCTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type FCTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FCTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FCTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FCTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FCTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FCTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FCTokenSession struct {
	Contract     *FCToken          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FCTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FCTokenCallerSession struct {
	Contract *FCTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FCTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FCTokenTransactorSession struct {
	Contract     *FCTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FCTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type FCTokenRaw struct {
	Contract *FCToken // Generic contract binding to access the raw methods on
}

// FCTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FCTokenCallerRaw struct {
	Contract *FCTokenCaller // Generic read-only contract binding to access the raw methods on
}

// FCTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FCTokenTransactorRaw struct {
	Contract *FCTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFCToken creates a new instance of FCToken, bound to a specific deployed contract.
func NewFCToken(address common.Address, backend bind.ContractBackend) (*FCToken, error) {
	contract, err := bindFCToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FCToken{FCTokenCaller: FCTokenCaller{contract: contract}, FCTokenTransactor: FCTokenTransactor{contract: contract}, FCTokenFilterer: FCTokenFilterer{contract: contract}}, nil
}

// NewFCTokenCaller creates a new read-only instance of FCToken, bound to a specific deployed contract.
func NewFCTokenCaller(address common.Address, caller bind.ContractCaller) (*FCTokenCaller, error) {
	contract, err := bindFCToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FCTokenCaller{contract: contract}, nil
}

// NewFCTokenTransactor creates a new write-only instance of FCToken, bound to a specific deployed contract.
func NewFCTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*FCTokenTransactor, error) {
	contract, err := bindFCToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FCTokenTransactor{contract: contract}, nil
}

// NewFCTokenFilterer creates a new log filterer instance of FCToken, bound to a specific deployed contract.
func NewFCTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*FCTokenFilterer, error) {
	contract, err := bindFCToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FCTokenFilterer{contract: contract}, nil
}

// bindFCToken binds a generic wrapper to an already deployed contract.
func bindFCToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FCTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FCToken *FCTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FCToken.Contract.FCTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FCToken *FCTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FCToken.Contract.FCTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FCToken *FCTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FCToken.Contract.FCTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FCToken *FCTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FCToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FCToken *FCTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FCToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FCToken *FCTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FCToken.Contract.contract.Transact(opts, method, params...)
}

// Accounts is a free data retrieval call binding the contract method 0xf2a40db8.
//
// Solidity: function accounts(uint256 ) view returns(address)
func (_FCToken *FCTokenCaller) Accounts(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FCToken.contract.Call(opts, &out, "accounts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Accounts is a free data retrieval call binding the contract method 0xf2a40db8.
//
// Solidity: function accounts(uint256 ) view returns(address)
func (_FCToken *FCTokenSession) Accounts(arg0 *big.Int) (common.Address, error) {
	return _FCToken.Contract.Accounts(&_FCToken.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0xf2a40db8.
//
// Solidity: function accounts(uint256 ) view returns(address)
func (_FCToken *FCTokenCallerSession) Accounts(arg0 *big.Int) (common.Address, error) {
	return _FCToken.Contract.Accounts(&_FCToken.CallOpts, arg0)
}

// AccountsMap is a free data retrieval call binding the contract method 0xe946cfd4.
//
// Solidity: function accountsMap(address ) view returns(bool)
func (_FCToken *FCTokenCaller) AccountsMap(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FCToken.contract.Call(opts, &out, "accountsMap", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AccountsMap is a free data retrieval call binding the contract method 0xe946cfd4.
//
// Solidity: function accountsMap(address ) view returns(bool)
func (_FCToken *FCTokenSession) AccountsMap(arg0 common.Address) (bool, error) {
	return _FCToken.Contract.AccountsMap(&_FCToken.CallOpts, arg0)
}

// AccountsMap is a free data retrieval call binding the contract method 0xe946cfd4.
//
// Solidity: function accountsMap(address ) view returns(bool)
func (_FCToken *FCTokenCallerSession) AccountsMap(arg0 common.Address) (bool, error) {
	return _FCToken.Contract.AccountsMap(&_FCToken.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_FCToken *FCTokenCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FCToken.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_FCToken *FCTokenSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _FCToken.Contract.Allowance(&_FCToken.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_FCToken *FCTokenCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _FCToken.Contract.Allowance(&_FCToken.CallOpts, arg0, arg1)
}

// Allowers is a free data retrieval call binding the contract method 0x1de4d19d.
//
// Solidity: function allowers(address ) view returns(uint8)
func (_FCToken *FCTokenCaller) Allowers(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _FCToken.contract.Call(opts, &out, "allowers", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Allowers is a free data retrieval call binding the contract method 0x1de4d19d.
//
// Solidity: function allowers(address ) view returns(uint8)
func (_FCToken *FCTokenSession) Allowers(arg0 common.Address) (uint8, error) {
	return _FCToken.Contract.Allowers(&_FCToken.CallOpts, arg0)
}

// Allowers is a free data retrieval call binding the contract method 0x1de4d19d.
//
// Solidity: function allowers(address ) view returns(uint8)
func (_FCToken *FCTokenCallerSession) Allowers(arg0 common.Address) (uint8, error) {
	return _FCToken.Contract.Allowers(&_FCToken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_FCToken *FCTokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FCToken.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_FCToken *FCTokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _FCToken.Contract.BalanceOf(&_FCToken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_FCToken *FCTokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _FCToken.Contract.BalanceOf(&_FCToken.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FCToken *FCTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FCToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FCToken *FCTokenSession) Decimals() (uint8, error) {
	return _FCToken.Contract.Decimals(&_FCToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FCToken *FCTokenCallerSession) Decimals() (uint8, error) {
	return _FCToken.Contract.Decimals(&_FCToken.CallOpts)
}

// GetAccountCount is a free data retrieval call binding the contract method 0xa98e4e77.
//
// Solidity: function getAccountCount() view returns(uint256)
func (_FCToken *FCTokenCaller) GetAccountCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FCToken.contract.Call(opts, &out, "getAccountCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountCount is a free data retrieval call binding the contract method 0xa98e4e77.
//
// Solidity: function getAccountCount() view returns(uint256)
func (_FCToken *FCTokenSession) GetAccountCount() (*big.Int, error) {
	return _FCToken.Contract.GetAccountCount(&_FCToken.CallOpts)
}

// GetAccountCount is a free data retrieval call binding the contract method 0xa98e4e77.
//
// Solidity: function getAccountCount() view returns(uint256)
func (_FCToken *FCTokenCallerSession) GetAccountCount() (*big.Int, error) {
	return _FCToken.Contract.GetAccountCount(&_FCToken.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_FCToken *FCTokenCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FCToken.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_FCToken *FCTokenSession) GetBalance() (*big.Int, error) {
	return _FCToken.Contract.GetBalance(&_FCToken.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_FCToken *FCTokenCallerSession) GetBalance() (*big.Int, error) {
	return _FCToken.Contract.GetBalance(&_FCToken.CallOpts)
}

// GetCounts is a free data retrieval call binding the contract method 0x615715a2.
//
// Solidity: function getCounts(uint8 id) view returns(uint256)
func (_FCToken *FCTokenCaller) GetCounts(opts *bind.CallOpts, id uint8) (*big.Int, error) {
	var out []interface{}
	err := _FCToken.contract.Call(opts, &out, "getCounts", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCounts is a free data retrieval call binding the contract method 0x615715a2.
//
// Solidity: function getCounts(uint8 id) view returns(uint256)
func (_FCToken *FCTokenSession) GetCounts(id uint8) (*big.Int, error) {
	return _FCToken.Contract.GetCounts(&_FCToken.CallOpts, id)
}

// GetCounts is a free data retrieval call binding the contract method 0x615715a2.
//
// Solidity: function getCounts(uint8 id) view returns(uint256)
func (_FCToken *FCTokenCallerSession) GetCounts(id uint8) (*big.Int, error) {
	return _FCToken.Contract.GetCounts(&_FCToken.CallOpts, id)
}

// GetPendingTranscation is a free data retrieval call binding the contract method 0x519ed728.
//
// Solidity: function getPendingTranscation() view returns(uint256)
func (_FCToken *FCTokenCaller) GetPendingTranscation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FCToken.contract.Call(opts, &out, "getPendingTranscation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingTranscation is a free data retrieval call binding the contract method 0x519ed728.
//
// Solidity: function getPendingTranscation() view returns(uint256)
func (_FCToken *FCTokenSession) GetPendingTranscation() (*big.Int, error) {
	return _FCToken.Contract.GetPendingTranscation(&_FCToken.CallOpts)
}

// GetPendingTranscation is a free data retrieval call binding the contract method 0x519ed728.
//
// Solidity: function getPendingTranscation() view returns(uint256)
func (_FCToken *FCTokenCallerSession) GetPendingTranscation() (*big.Int, error) {
	return _FCToken.Contract.GetPendingTranscation(&_FCToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FCToken *FCTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FCToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FCToken *FCTokenSession) Name() (string, error) {
	return _FCToken.Contract.Name(&_FCToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FCToken *FCTokenCallerSession) Name() (string, error) {
	return _FCToken.Contract.Name(&_FCToken.CallOpts)
}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) view returns(string)
func (_FCToken *FCTokenCaller) Nodes(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _FCToken.contract.Call(opts, &out, "nodes", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) view returns(string)
func (_FCToken *FCTokenSession) Nodes(arg0 *big.Int) (string, error) {
	return _FCToken.Contract.Nodes(&_FCToken.CallOpts, arg0)
}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) view returns(string)
func (_FCToken *FCTokenCallerSession) Nodes(arg0 *big.Int) (string, error) {
	return _FCToken.Contract.Nodes(&_FCToken.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FCToken *FCTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FCToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FCToken *FCTokenSession) Owner() (common.Address, error) {
	return _FCToken.Contract.Owner(&_FCToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FCToken *FCTokenCallerSession) Owner() (common.Address, error) {
	return _FCToken.Contract.Owner(&_FCToken.CallOpts)
}

// Proportion is a free data retrieval call binding the contract method 0x5b4246d4.
//
// Solidity: function proportion() view returns(uint256)
func (_FCToken *FCTokenCaller) Proportion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FCToken.contract.Call(opts, &out, "proportion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Proportion is a free data retrieval call binding the contract method 0x5b4246d4.
//
// Solidity: function proportion() view returns(uint256)
func (_FCToken *FCTokenSession) Proportion() (*big.Int, error) {
	return _FCToken.Contract.Proportion(&_FCToken.CallOpts)
}

// Proportion is a free data retrieval call binding the contract method 0x5b4246d4.
//
// Solidity: function proportion() view returns(uint256)
func (_FCToken *FCTokenCallerSession) Proportion() (*big.Int, error) {
	return _FCToken.Contract.Proportion(&_FCToken.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x0536b723.
//
// Solidity: function signer(address ) view returns(uint8)
func (_FCToken *FCTokenCaller) Signer(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _FCToken.contract.Call(opts, &out, "signer", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x0536b723.
//
// Solidity: function signer(address ) view returns(uint8)
func (_FCToken *FCTokenSession) Signer(arg0 common.Address) (uint8, error) {
	return _FCToken.Contract.Signer(&_FCToken.CallOpts, arg0)
}

// Signer is a free data retrieval call binding the contract method 0x0536b723.
//
// Solidity: function signer(address ) view returns(uint8)
func (_FCToken *FCTokenCallerSession) Signer(arg0 common.Address) (uint8, error) {
	return _FCToken.Contract.Signer(&_FCToken.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FCToken *FCTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FCToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FCToken *FCTokenSession) Symbol() (string, error) {
	return _FCToken.Contract.Symbol(&_FCToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FCToken *FCTokenCallerSession) Symbol() (string, error) {
	return _FCToken.Contract.Symbol(&_FCToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FCToken *FCTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FCToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FCToken *FCTokenSession) TotalSupply() (*big.Int, error) {
	return _FCToken.Contract.TotalSupply(&_FCToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FCToken *FCTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _FCToken.Contract.TotalSupply(&_FCToken.CallOpts)
}

// TranscationId is a free data retrieval call binding the contract method 0x87019cd4.
//
// Solidity: function transcationId() view returns(uint8)
func (_FCToken *FCTokenCaller) TranscationId(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FCToken.contract.Call(opts, &out, "transcationId")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TranscationId is a free data retrieval call binding the contract method 0x87019cd4.
//
// Solidity: function transcationId() view returns(uint8)
func (_FCToken *FCTokenSession) TranscationId() (uint8, error) {
	return _FCToken.Contract.TranscationId(&_FCToken.CallOpts)
}

// TranscationId is a free data retrieval call binding the contract method 0x87019cd4.
//
// Solidity: function transcationId() view returns(uint8)
func (_FCToken *FCTokenCallerSession) TranscationId() (uint8, error) {
	return _FCToken.Contract.TranscationId(&_FCToken.CallOpts)
}

// Transcations is a free data retrieval call binding the contract method 0x78e8691d.
//
// Solidity: function transcations(uint256 ) view returns(address from, address to, uint256 amount, uint8 signatureCounts)
func (_FCToken *FCTokenCaller) Transcations(opts *bind.CallOpts, arg0 *big.Int) (struct {
	From            common.Address
	To              common.Address
	Amount          *big.Int
	SignatureCounts uint8
}, error) {
	var out []interface{}
	err := _FCToken.contract.Call(opts, &out, "transcations", arg0)

	outstruct := new(struct {
		From            common.Address
		To              common.Address
		Amount          *big.Int
		SignatureCounts uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.From = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.To = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SignatureCounts = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// Transcations is a free data retrieval call binding the contract method 0x78e8691d.
//
// Solidity: function transcations(uint256 ) view returns(address from, address to, uint256 amount, uint8 signatureCounts)
func (_FCToken *FCTokenSession) Transcations(arg0 *big.Int) (struct {
	From            common.Address
	To              common.Address
	Amount          *big.Int
	SignatureCounts uint8
}, error) {
	return _FCToken.Contract.Transcations(&_FCToken.CallOpts, arg0)
}

// Transcations is a free data retrieval call binding the contract method 0x78e8691d.
//
// Solidity: function transcations(uint256 ) view returns(address from, address to, uint256 amount, uint8 signatureCounts)
func (_FCToken *FCTokenCallerSession) Transcations(arg0 *big.Int) (struct {
	From            common.Address
	To              common.Address
	Amount          *big.Int
	SignatureCounts uint8
}, error) {
	return _FCToken.Contract.Transcations(&_FCToken.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _delegatee, uint256 _value) returns()
func (_FCToken *FCTokenTransactor) Approve(opts *bind.TransactOpts, _delegatee common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FCToken.contract.Transact(opts, "approve", _delegatee, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _delegatee, uint256 _value) returns()
func (_FCToken *FCTokenSession) Approve(_delegatee common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FCToken.Contract.Approve(&_FCToken.TransactOpts, _delegatee, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _delegatee, uint256 _value) returns()
func (_FCToken *FCTokenTransactorSession) Approve(_delegatee common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FCToken.Contract.Approve(&_FCToken.TransactOpts, _delegatee, _value)
}

// ApproveWithArray is a paid mutator transaction binding the contract method 0x275334ae.
//
// Solidity: function approveWithArray(uint8[] b, uint256 len) returns()
func (_FCToken *FCTokenTransactor) ApproveWithArray(opts *bind.TransactOpts, b []uint8, len *big.Int) (*types.Transaction, error) {
	return _FCToken.contract.Transact(opts, "approveWithArray", b, len)
}

// ApproveWithArray is a paid mutator transaction binding the contract method 0x275334ae.
//
// Solidity: function approveWithArray(uint8[] b, uint256 len) returns()
func (_FCToken *FCTokenSession) ApproveWithArray(b []uint8, len *big.Int) (*types.Transaction, error) {
	return _FCToken.Contract.ApproveWithArray(&_FCToken.TransactOpts, b, len)
}

// ApproveWithArray is a paid mutator transaction binding the contract method 0x275334ae.
//
// Solidity: function approveWithArray(uint8[] b, uint256 len) returns()
func (_FCToken *FCTokenTransactorSession) ApproveWithArray(b []uint8, len *big.Int) (*types.Transaction, error) {
	return _FCToken.Contract.ApproveWithArray(&_FCToken.TransactOpts, b, len)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() payable returns(uint256 amount)
func (_FCToken *FCTokenTransactor) Buy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FCToken.contract.Transact(opts, "buy")
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() payable returns(uint256 amount)
func (_FCToken *FCTokenSession) Buy() (*types.Transaction, error) {
	return _FCToken.Contract.Buy(&_FCToken.TransactOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() payable returns(uint256 amount)
func (_FCToken *FCTokenTransactorSession) Buy() (*types.Transaction, error) {
	return _FCToken.Contract.Buy(&_FCToken.TransactOpts)
}

// CreateMiner is a paid mutator transaction binding the contract method 0x8f75c003.
//
// Solidity: function createMiner(address miner, string nodeId) returns()
func (_FCToken *FCTokenTransactor) CreateMiner(opts *bind.TransactOpts, miner common.Address, nodeId string) (*types.Transaction, error) {
	return _FCToken.contract.Transact(opts, "createMiner", miner, nodeId)
}

// CreateMiner is a paid mutator transaction binding the contract method 0x8f75c003.
//
// Solidity: function createMiner(address miner, string nodeId) returns()
func (_FCToken *FCTokenSession) CreateMiner(miner common.Address, nodeId string) (*types.Transaction, error) {
	return _FCToken.Contract.CreateMiner(&_FCToken.TransactOpts, miner, nodeId)
}

// CreateMiner is a paid mutator transaction binding the contract method 0x8f75c003.
//
// Solidity: function createMiner(address miner, string nodeId) returns()
func (_FCToken *FCTokenTransactorSession) CreateMiner(miner common.Address, nodeId string) (*types.Transaction, error) {
	return _FCToken.Contract.CreateMiner(&_FCToken.TransactOpts, miner, nodeId)
}

// MultiSignature is a paid mutator transaction binding the contract method 0x3458810a.
//
// Solidity: function multiSignature(uint8 id) returns()
func (_FCToken *FCTokenTransactor) MultiSignature(opts *bind.TransactOpts, id uint8) (*types.Transaction, error) {
	return _FCToken.contract.Transact(opts, "multiSignature", id)
}

// MultiSignature is a paid mutator transaction binding the contract method 0x3458810a.
//
// Solidity: function multiSignature(uint8 id) returns()
func (_FCToken *FCTokenSession) MultiSignature(id uint8) (*types.Transaction, error) {
	return _FCToken.Contract.MultiSignature(&_FCToken.TransactOpts, id)
}

// MultiSignature is a paid mutator transaction binding the contract method 0x3458810a.
//
// Solidity: function multiSignature(uint8 id) returns()
func (_FCToken *FCTokenTransactorSession) MultiSignature(id uint8) (*types.Transaction, error) {
	return _FCToken.Contract.MultiSignature(&_FCToken.TransactOpts, id)
}

// RemoveAllower is a paid mutator transaction binding the contract method 0xb869abf7.
//
// Solidity: function removeAllower(address s) returns()
func (_FCToken *FCTokenTransactor) RemoveAllower(opts *bind.TransactOpts, s common.Address) (*types.Transaction, error) {
	return _FCToken.contract.Transact(opts, "removeAllower", s)
}

// RemoveAllower is a paid mutator transaction binding the contract method 0xb869abf7.
//
// Solidity: function removeAllower(address s) returns()
func (_FCToken *FCTokenSession) RemoveAllower(s common.Address) (*types.Transaction, error) {
	return _FCToken.Contract.RemoveAllower(&_FCToken.TransactOpts, s)
}

// RemoveAllower is a paid mutator transaction binding the contract method 0xb869abf7.
//
// Solidity: function removeAllower(address s) returns()
func (_FCToken *FCTokenTransactorSession) RemoveAllower(s common.Address) (*types.Transaction, error) {
	return _FCToken.Contract.RemoveAllower(&_FCToken.TransactOpts, s)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x0e316ab7.
//
// Solidity: function removeSigner(address s) returns()
func (_FCToken *FCTokenTransactor) RemoveSigner(opts *bind.TransactOpts, s common.Address) (*types.Transaction, error) {
	return _FCToken.contract.Transact(opts, "removeSigner", s)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x0e316ab7.
//
// Solidity: function removeSigner(address s) returns()
func (_FCToken *FCTokenSession) RemoveSigner(s common.Address) (*types.Transaction, error) {
	return _FCToken.Contract.RemoveSigner(&_FCToken.TransactOpts, s)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x0e316ab7.
//
// Solidity: function removeSigner(address s) returns()
func (_FCToken *FCTokenTransactorSession) RemoveSigner(s common.Address) (*types.Transaction, error) {
	return _FCToken.Contract.RemoveSigner(&_FCToken.TransactOpts, s)
}

// SetAllower is a paid mutator transaction binding the contract method 0x4679a9e1.
//
// Solidity: function setAllower(address s) returns()
func (_FCToken *FCTokenTransactor) SetAllower(opts *bind.TransactOpts, s common.Address) (*types.Transaction, error) {
	return _FCToken.contract.Transact(opts, "setAllower", s)
}

// SetAllower is a paid mutator transaction binding the contract method 0x4679a9e1.
//
// Solidity: function setAllower(address s) returns()
func (_FCToken *FCTokenSession) SetAllower(s common.Address) (*types.Transaction, error) {
	return _FCToken.Contract.SetAllower(&_FCToken.TransactOpts, s)
}

// SetAllower is a paid mutator transaction binding the contract method 0x4679a9e1.
//
// Solidity: function setAllower(address s) returns()
func (_FCToken *FCTokenTransactorSession) SetAllower(s common.Address) (*types.Transaction, error) {
	return _FCToken.Contract.SetAllower(&_FCToken.TransactOpts, s)
}

// SetPropertion is a paid mutator transaction binding the contract method 0x675dcedf.
//
// Solidity: function setPropertion(uint256 num) returns()
func (_FCToken *FCTokenTransactor) SetPropertion(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _FCToken.contract.Transact(opts, "setPropertion", num)
}

// SetPropertion is a paid mutator transaction binding the contract method 0x675dcedf.
//
// Solidity: function setPropertion(uint256 num) returns()
func (_FCToken *FCTokenSession) SetPropertion(num *big.Int) (*types.Transaction, error) {
	return _FCToken.Contract.SetPropertion(&_FCToken.TransactOpts, num)
}

// SetPropertion is a paid mutator transaction binding the contract method 0x675dcedf.
//
// Solidity: function setPropertion(uint256 num) returns()
func (_FCToken *FCTokenTransactorSession) SetPropertion(num *big.Int) (*types.Transaction, error) {
	return _FCToken.Contract.SetPropertion(&_FCToken.TransactOpts, num)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address s) returns()
func (_FCToken *FCTokenTransactor) SetSigner(opts *bind.TransactOpts, s common.Address) (*types.Transaction, error) {
	return _FCToken.contract.Transact(opts, "setSigner", s)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address s) returns()
func (_FCToken *FCTokenSession) SetSigner(s common.Address) (*types.Transaction, error) {
	return _FCToken.Contract.SetSigner(&_FCToken.TransactOpts, s)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address s) returns()
func (_FCToken *FCTokenTransactorSession) SetSigner(s common.Address) (*types.Transaction, error) {
	return _FCToken.Contract.SetSigner(&_FCToken.TransactOpts, s)
}

// StartTransfer is a paid mutator transaction binding the contract method 0x604374f0.
//
// Solidity: function startTransfer(address _to, uint256 _amount) returns()
func (_FCToken *FCTokenTransactor) StartTransfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FCToken.contract.Transact(opts, "startTransfer", _to, _amount)
}

// StartTransfer is a paid mutator transaction binding the contract method 0x604374f0.
//
// Solidity: function startTransfer(address _to, uint256 _amount) returns()
func (_FCToken *FCTokenSession) StartTransfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FCToken.Contract.StartTransfer(&_FCToken.TransactOpts, _to, _amount)
}

// StartTransfer is a paid mutator transaction binding the contract method 0x604374f0.
//
// Solidity: function startTransfer(address _to, uint256 _amount) returns()
func (_FCToken *FCTokenTransactorSession) StartTransfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FCToken.Contract.StartTransfer(&_FCToken.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_FCToken *FCTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FCToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_FCToken *FCTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FCToken.Contract.Transfer(&_FCToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_FCToken *FCTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FCToken.Contract.Transfer(&_FCToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns()
func (_FCToken *FCTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FCToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns()
func (_FCToken *FCTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FCToken.Contract.TransferFrom(&_FCToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns()
func (_FCToken *FCTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FCToken.Contract.TransferFrom(&_FCToken.TransactOpts, _from, _to, _value)
}

// FCTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the FCToken contract.
type FCTokenApprovalIterator struct {
	Event *FCTokenApproval // Event containing the contract specifics and raw log

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
func (it *FCTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FCTokenApproval)
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
		it.Event = new(FCTokenApproval)
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
func (it *FCTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FCTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FCTokenApproval represents a Approval event raised by the FCToken contract.
type FCTokenApproval struct {
	From      common.Address
	Delegatee common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed from, address indexed delegatee, uint256 value)
func (_FCToken *FCTokenFilterer) FilterApproval(opts *bind.FilterOpts, from []common.Address, delegatee []common.Address) (*FCTokenApprovalIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _FCToken.contract.FilterLogs(opts, "Approval", fromRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return &FCTokenApprovalIterator{contract: _FCToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed from, address indexed delegatee, uint256 value)
func (_FCToken *FCTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FCTokenApproval, from []common.Address, delegatee []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _FCToken.contract.WatchLogs(opts, "Approval", fromRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FCTokenApproval)
				if err := _FCToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed from, address indexed delegatee, uint256 value)
func (_FCToken *FCTokenFilterer) ParseApproval(log types.Log) (*FCTokenApproval, error) {
	event := new(FCTokenApproval)
	if err := _FCToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FCTokenCreateTranscationIterator is returned from FilterCreateTranscation and is used to iterate over the raw logs and unpacked data for CreateTranscation events raised by the FCToken contract.
type FCTokenCreateTranscationIterator struct {
	Event *FCTokenCreateTranscation // Event containing the contract specifics and raw log

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
func (it *FCTokenCreateTranscationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FCTokenCreateTranscation)
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
		it.Event = new(FCTokenCreateTranscation)
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
func (it *FCTokenCreateTranscationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FCTokenCreateTranscationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FCTokenCreateTranscation represents a CreateTranscation event raised by the FCToken contract.
type FCTokenCreateTranscation struct {
	From          common.Address
	To            common.Address
	Amount        *big.Int
	TranscationId uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCreateTranscation is a free log retrieval operation binding the contract event 0xe1dd58c97af942b2e36ca0fc977e17197d6a8008fda5381780e291337e8ac995.
//
// Solidity: event CreateTranscation(address from, address to, uint256 amount, uint8 transcationId)
func (_FCToken *FCTokenFilterer) FilterCreateTranscation(opts *bind.FilterOpts) (*FCTokenCreateTranscationIterator, error) {

	logs, sub, err := _FCToken.contract.FilterLogs(opts, "CreateTranscation")
	if err != nil {
		return nil, err
	}
	return &FCTokenCreateTranscationIterator{contract: _FCToken.contract, event: "CreateTranscation", logs: logs, sub: sub}, nil
}

// WatchCreateTranscation is a free log subscription operation binding the contract event 0xe1dd58c97af942b2e36ca0fc977e17197d6a8008fda5381780e291337e8ac995.
//
// Solidity: event CreateTranscation(address from, address to, uint256 amount, uint8 transcationId)
func (_FCToken *FCTokenFilterer) WatchCreateTranscation(opts *bind.WatchOpts, sink chan<- *FCTokenCreateTranscation) (event.Subscription, error) {

	logs, sub, err := _FCToken.contract.WatchLogs(opts, "CreateTranscation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FCTokenCreateTranscation)
				if err := _FCToken.contract.UnpackLog(event, "CreateTranscation", log); err != nil {
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

// ParseCreateTranscation is a log parse operation binding the contract event 0xe1dd58c97af942b2e36ca0fc977e17197d6a8008fda5381780e291337e8ac995.
//
// Solidity: event CreateTranscation(address from, address to, uint256 amount, uint8 transcationId)
func (_FCToken *FCTokenFilterer) ParseCreateTranscation(log types.Log) (*FCTokenCreateTranscation, error) {
	event := new(FCTokenCreateTranscation)
	if err := _FCToken.contract.UnpackLog(event, "CreateTranscation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FCTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the FCToken contract.
type FCTokenTransferIterator struct {
	Event *FCTokenTransfer // Event containing the contract specifics and raw log

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
func (it *FCTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FCTokenTransfer)
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
		it.Event = new(FCTokenTransfer)
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
func (it *FCTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FCTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FCTokenTransfer represents a Transfer event raised by the FCToken contract.
type FCTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FCToken *FCTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FCTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FCToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FCTokenTransferIterator{contract: _FCToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FCToken *FCTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FCTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FCToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FCTokenTransfer)
				if err := _FCToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FCToken *FCTokenFilterer) ParseTransfer(log types.Log) (*FCTokenTransfer, error) {
	event := new(FCTokenTransfer)
	if err := _FCToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

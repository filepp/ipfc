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
const FCTokenABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_initiSuppy\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"transcationId\",\"type\":\"uint8\"}],\"name\":\"CreateTranscation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accountsMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowers\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8[]\",\"name\":\"b\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256\",\"name\":\"len\",\"type\":\"uint256\"}],\"name\":\"approveWithArray\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"}],\"name\":\"createMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"getCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingTranscation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"multiSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"}],\"name\":\"removeAllower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"}],\"name\":\"removeSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"}],\"name\":\"setAllower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"startTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transcationId\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transcations\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"signatureCounts\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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
	"b869abf7": "removeAllower(address)",
	"0e316ab7": "removeSigner(address)",
	"4679a9e1": "setAllower(address)",
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
var FCTokenBin = "0x60806040523480156200001157600080fd5b5060405162001d8438038062001d8483398101604081905262000034916200021a565b835162000049906000906020870190620000c1565b5082516200005f906001906020860190620000c1565b506002805460ff191660ff84169081179091556200007f90600a620002ee565b6200008b9082620003e3565b600381905533600081815260056020526040902091909155600480546001600160a01b0319169091179055506200046e92505050565b828054620000cf9062000405565b90600052602060002090601f016020900481019282620000f357600085556200013e565b82601f106200010e57805160ff19168380011785556200013e565b828001600101855582156200013e579182015b828111156200013e57825182559160200191906001019062000121565b506200014c92915062000150565b5090565b5b808211156200014c576000815560010162000151565b600082601f83011262000178578081fd5b81516001600160401b038082111562000195576200019562000458565b604051601f8301601f19908116603f01168101908282118183101715620001c057620001c062000458565b81604052838152602092508683858801011115620001dc578485fd5b8491505b83821015620001ff5785820183015181830184015290820190620001e0565b838211156200021057848385830101525b9695505050505050565b6000806000806080858703121562000230578384fd5b84516001600160401b038082111562000247578586fd5b620002558883890162000167565b955060208701519150808211156200026b578485fd5b506200027a8782880162000167565b935050604085015160ff8116811462000291578283fd5b6060959095015193969295505050565b80825b6001808611620002b55750620002e5565b818704821115620002ca57620002ca62000442565b80861615620002d857918102915b9490941c938002620002a4565b94509492505050565b6000620002ff600019848462000306565b9392505050565b6000826200031757506001620002ff565b816200032657506000620002ff565b81600181146200033f57600281146200034a576200037e565b6001915050620002ff565b60ff8411156200035e576200035e62000442565b6001841b91508482111562000377576200037762000442565b50620002ff565b5060208310610133831016604e8410600b8410161715620003b6575081810a83811115620003b057620003b062000442565b620002ff565b620003c58484846001620002a1565b808604821115620003da57620003da62000442565b02949350505050565b600081600019048311821515161562000400576200040062000442565b500290565b600181811c908216806200041a57607f821691505b602082108114156200043c57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b611906806200047e6000396000f3fe6080604052600436106101cd5760003560e01c8063615715a2116100f757806395d89b4111610095578063b869abf711610064578063b869abf71461058f578063dd62ed3e146105af578063e946cfd4146105e7578063f2a40db814610627576101cd565b806395d89b411461053d578063a6f2ae3a14610552578063a9059cbb1461055a578063a98e4e771461057a576101cd565b806378e8691d116100d157806378e8691d1461044b57806387019cd4146104cb5780638da5cb5b146104e55780638f75c0031461051d576101cd565b8063615715a2146103de5780636c19e783146103fe57806370a082311461041e576101cd565b80631de4d19d1161016f5780633458810a1161013e5780633458810a146103695780634679a9e114610389578063519ed728146103a9578063604374f0146103be576101cd565b80631de4d19d146102df57806323b872dd1461030f578063275334ae1461032f578063313ce5671461034f576101cd565b80630e316ab7116101ab5780630e316ab71461025d57806312065fe01461027d57806318160ddd146102a95780631c53c280146102bf576101cd565b80630536b723146101d257806306fdde0314610219578063095ea7b31461023b575b600080fd5b3480156101de57600080fd5b506102026101ed3660046114a4565b60076020526000908152604090205460ff1681565b60405160ff90911681526020015b60405180910390f35b34801561022557600080fd5b5061022e610647565b60405161021091906116e2565b34801561024757600080fd5b5061025b6102563660046115d2565b6106d5565b005b34801561026957600080fd5b5061025b6102783660046114a4565b6107b1565b34801561028957600080fd5b50336000908152600560205260409020545b604051908152602001610210565b3480156102b557600080fd5b5061029b60035481565b3480156102cb57600080fd5b5061022e6102da3660046116b0565b6107e9565b3480156102eb57600080fd5b506102026102fa3660046114a4565b60086020526000908152604090205460ff1681565b34801561031b57600080fd5b5061025b61032a3660046114f7565b610814565b34801561033b57600080fd5b5061025b61034a3660046115fb565b6108f0565b34801561035b57600080fd5b506002546102029060ff1681565b34801561037557600080fd5b5061025b6103843660046116c8565b610c6d565b34801561039557600080fd5b5061025b6103a43660046114a4565b610deb565b3480156103b557600080fd5b50600c5461029b565b3480156103ca57600080fd5b5061025b6103d93660046115d2565b610e26565b3480156103ea57600080fd5b5061029b6103f93660046116c8565b610f74565b34801561040a57600080fd5b5061025b6104193660046114a4565b610f92565b34801561042a57600080fd5b5061029b6104393660046114a4565b60056020526000908152604090205481565b34801561045757600080fd5b5061049d6104663660046116b0565b600e6020526000908152604090208054600182015460028301546003909301546001600160a01b0392831693919092169160ff1684565b604080516001600160a01b0395861681529490931660208501529183015260ff166060820152608001610210565b3480156104d757600080fd5b50600d546102029060ff1681565b3480156104f157600080fd5b50600454610505906001600160a01b031681565b6040516001600160a01b039091168152602001610210565b34801561052957600080fd5b5061025b610538366004611532565b610fcd565b34801561054957600080fd5b5061022e6110c1565b61029b6110ce565b34801561056657600080fd5b5061025b6105753660046115d2565b6111a6565b34801561058657600080fd5b5060095461029b565b34801561059b57600080fd5b5061025b6105aa3660046114a4565b6111b1565b3480156105bb57600080fd5b5061029b6105ca3660046114c5565b600660209081526000928352604080842090915290825290205481565b3480156105f357600080fd5b506106176106023660046114a4565b600b6020526000908152604090205460ff1681565b6040519015158152602001610210565b34801561063357600080fd5b506105056106423660046116b0565b6111e9565b6000805461065490611804565b80601f016020809104026020016040519081016040528092919081815260200182805461068090611804565b80156106cd5780601f106106a2576101008083540402835291602001916106cd565b820191906000526020600020905b8154815290600101906020018083116106b057829003601f168201915b505050505081565b3360009081526008602052604090205460ff166001146106f457600080fd5b6001600160a01b03821661070757600080fd5b6004546001600160a01b031660009081526005602052604090205481111561072e57600080fd5b6004546001600160a01b03908116600090815260066020908152604080832093861683529290529081208054839290610768908490611766565b90915550506040518181526001600160a01b0383169033907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259060200160405180910390a35050565b6004546001600160a01b031633146107c857600080fd5b6001600160a01b03166000908152600760205260409020805460ff19169055565b600a81815481106107f957600080fd5b90600052602060002001600091509050805461065490611804565b6001600160a01b03821661082757600080fd5b6004546001600160a01b0316600090815260066020908152604080832033845290915290205481111561085957600080fd5b6001600160a01b03831660009081526006602090815260408083203384529091528120805483929061088c9084906117d6565b90915550506001600160a01b038316600090815260056020526040812080548392906108b99084906117d6565b90915550506001600160a01b038216600090815260056020526040812080548392906108e6908490611766565b9091555050505050565b60008251116108fe57600080fd5b3360009081526008602052604090205460ff1660011461091d57600080fd5b6000811161092a57600080fd5b60008167ffffffffffffffff81111561095357634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561097c578160200160208202803683370190505b5083519091506000905b8015610a6c5760008561099a6001846117d6565b815181106109b857634e487b7160e01b600052603260045260246000fd5b6020026020010151905060005b60088160ff161015610a575760008160ff166001901b8360ff16161115610a45578060ff168388516109f791906117d6565b610a029060086117b7565b610a0c9190611766565b858581518110610a2c57634e487b7160e01b600052603260045260246000fd5b602090810291909101015283610a418161183f565b9450505b80610a4f8161185a565b9150506109c5565b50508080610a64906117ed565b915050610986565b508151600a90811115610a7d575081515b60008167ffffffffffffffff811115610aa657634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610acf578160200160208202803683370190505b50905060005b82811015610c645760004442604051602001610afb929190918252602082015260400190565b6040516020818303038152906040528051906020012060001c90506000865182610b25919061187a565b90506000878281518110610b4957634e487b7160e01b600052603260045260246000fd5b602002602001015190506000805b85811015610ba95782878281518110610b8057634e487b7160e01b600052603260045260246000fd5b60200260200101511415610b975760019150610ba9565b80610ba18161183f565b915050610b57565b5080610c5b5781868681518110610bd057634e487b7160e01b600052603260045260246000fd5b602002602001018181525050600060098381548110610bff57634e487b7160e01b600052603260045260246000fd5b60009182526020808320909101546004546001600160a01b039081168452600683526040808520919092168085529252822080549193506001929091610c46908490611766565b90915550869050610c568161183f565b965050505b50505050610ad5565b50505050505050565b6004546001600160a01b0316331480610c9857503360009081526007602052604090205460ff166001145b610ca157600080fd5b60ff81166000908152600e6020526040902080546001600160a01b0316610cc757600080fd5b80546001600160a01b0316331415610cde57600080fd5b33600090815260048201602052604090205460ff1660011415610d0057600080fd5b3360009081526004820160205260408120805460ff1916600117905560038201805460ff1691610d2f8361185a565b82546101009290920a60ff818102199093169183160217909155600383015460029116109050610de7576002810154306000908152600560205260409020541015610d7957600080fd5b60018101546002820154610d9a9130916001600160a01b0390911690611213565b610da6600c60006113bb565b60ff82166000908152600e6020526040812080546001600160a01b0319908116825560018201805490911690556002810191909155600301805460ff191690555b5050565b6004546001600160a01b03163314610e0257600080fd5b6001600160a01b03166000908152600860205260409020805460ff19166001179055565b6001600160a01b038216610e3957600080fd5b33600090815260056020526040902054811115610e5557600080fd5b610e60333083611213565b600d54610e719060ff16600161177e565b600d805460ff92831660ff19918216811783556000908152600e6020908152604080832080546001600160a01b0319908116339081178355600183810180546001600160a01b038e16941684179055600284018b90556003840180549098169097558754600c805498890181559096528487047fdf6966c971051c3d54ec59162606531493a51404a002842f56009d7e5cf4a8c7018054968a16601f9098166101000a978802978a02199096169690961790945594548151948552918401929092529082018590529092166060830152907fe1dd58c97af942b2e36ca0fc977e17197d6a8008fda5381780e291337e8ac9959060800160405180910390a1505050565b60ff8082166000908152600e6020526040902060030154165b919050565b6004546001600160a01b03163314610fa957600080fd5b6001600160a01b03166000908152600760205260409020805460ff19166001179055565b6001600160a01b038216610fe057600080fd5b8051603414610fee57600080fd5b6001600160a01b0382166000908152600b602052604090205460ff168015156001146110bc57600a80546001810182556000919091528251611057917fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a8019060208501906113e3565b506009805460018082019092557f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af0180546001600160a01b0319166001600160a01b0386169081179091556000908152600b60205260409020805460ff191690911790555b505050565b6001805461065490611804565b6000806110e2662386f26fc10000346117a3565b6004546001600160a01b031660009081526005602052604090205490915081111561110c57600080fd5b6004546001600160a01b0316600090815260056020526040812080548392906111369084906117d6565b9091555050336000908152600560205260408120805483929061115a908490611766565b909155505060045460405182815233916001600160a01b0316907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a3905090565b610de7338383611213565b6004546001600160a01b031633146111c857600080fd5b6001600160a01b03166000908152600860205260409020805460ff19169055565b600981815481106111f957600080fd5b6000918252602090912001546001600160a01b0316905081565b6004546001600160a01b031633141561122b57600080fd5b6001600160a01b03831660009081526005602052604090205481111561125057600080fd5b6001600160a01b03821661126357600080fd5b6001600160a01b0382166000908152600560205260409020546112868282611766565b1161129057600080fd5b6001600160a01b0380831660009081526005602052604080822054928616825281205490916112be91611766565b6001600160a01b0385166000908152600560205260408120805492935084929091906112eb9084906117d6565b90915550506001600160a01b03831660009081526005602052604081208054849290611318908490611766565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161136491815260200190565b60405180910390a36001600160a01b03808416600090815260056020526040808220549287168252902054829161139a91611766565b146113b557634e487b7160e01b600052600160045260246000fd5b50505050565b50805460008255601f0160209004906000526020600020908101906113e09190611467565b50565b8280546113ef90611804565b90600052602060002090601f0160209004810192826114115760008555611457565b82601f1061142a57805160ff1916838001178555611457565b82800160010185558215611457579182015b8281111561145757825182559160200191906001019061143c565b50611463929150611467565b5090565b5b808211156114635760008155600101611468565b80356001600160a01b0381168114610f8d57600080fd5b803560ff81168114610f8d57600080fd5b6000602082840312156114b5578081fd5b6114be8261147c565b9392505050565b600080604083850312156114d7578081fd5b6114e08361147c565b91506114ee6020840161147c565b90509250929050565b60008060006060848603121561150b578081fd5b6115148461147c565b92506115226020850161147c565b9150604084013590509250925092565b60008060408385031215611544578182fd5b61154d8361147c565b915060208084013567ffffffffffffffff8082111561156a578384fd5b818601915086601f83011261157d578384fd5b81358181111561158f5761158f6118ba565b6115a1601f8201601f19168501611735565b915080825287848285010111156115b6578485fd5b8084840185840137810190920192909252919491935090915050565b600080604083850312156115e4578182fd5b6115ed8361147c565b946020939093013593505050565b6000806040838503121561160d578182fd5b823567ffffffffffffffff80821115611624578384fd5b818501915085601f830112611637578384fd5b813560208282111561164b5761164b6118ba565b8160051b925061165c818401611735565b8281528181019085830185870184018b1015611676578889fd5b8896505b8487101561169f5761168b81611493565b83526001969096019591830191830161167a565b509997909101359750505050505050565b6000602082840312156116c1578081fd5b5035919050565b6000602082840312156116d9578081fd5b6114be82611493565b6000602080835283518082850152825b8181101561170e578581018301518582016040015282016116f2565b8181111561171f5783604083870101525b50601f01601f1916929092016040019392505050565b604051601f8201601f1916810167ffffffffffffffff8111828210171561175e5761175e6118ba565b604052919050565b600082198211156117795761177961188e565b500190565b600060ff821660ff84168060ff0382111561179b5761179b61188e565b019392505050565b6000826117b2576117b26118a4565b500490565b60008160001904831182151516156117d1576117d161188e565b500290565b6000828210156117e8576117e861188e565b500390565b6000816117fc576117fc61188e565b506000190190565b600181811c9082168061181857607f821691505b6020821081141561183957634e487b7160e01b600052602260045260246000fd5b50919050565b60006000198214156118535761185361188e565b5060010190565b600060ff821660ff8114156118715761187161188e565b60010192915050565b600082611889576118896118a4565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052604160045260246000fdfea2646970667358221220146c61c6b8381bb4db93bb308b31af345bdc1b23729a36705eea23e349ac210464736f6c63430008030033"

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

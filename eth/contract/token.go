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
const FCTokenABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_initiSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"transcationId\",\"type\":\"uint8\"}],\"name\":\"CreateTranscation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destorys\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Destory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Freeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Unfreeze\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvalBanlance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_spenders\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_values\",\"type\":\"uint256[]\"}],\"name\":\"approveWthArray\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"destory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"destoryFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"freeze\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"freezeOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"getCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingTranscation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"multiSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"}],\"name\":\"removeSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"startTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transcationId\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transcations\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"signatureCounts\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"unfreeze\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FCTokenFuncSigs maps the 4-byte function signature to its string representation.
var FCTokenFuncSigs = map[string]string{
	"75247a58": "approval(address,uint256)",
	"0307892a": "approvalBanlance(address,address)",
	"43272bc0": "approveWthArray(address[],uint256[])",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"908e049b": "destory(uint256)",
	"a9dece5b": "destoryFrom(address,uint256)",
	"24bce60c": "freeze(address,uint256)",
	"cd4217c1": "freezeOf(address)",
	"12065fe0": "getBalance()",
	"615715a2": "getCounts(uint8)",
	"519ed728": "getPendingTranscation()",
	"3458810a": "multiSignature(uint8)",
	"06fdde03": "name()",
	"8da5cb5b": "owner()",
	"0e316ab7": "removeSigner(address)",
	"6c19e783": "setSigner(address)",
	"0536b723": "signer(address)",
	"604374f0": "startTransfer(address,uint256)",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"87019cd4": "transcationId()",
	"78e8691d": "transcations(uint256)",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"7b46b80b": "unfreeze(address,uint256)",
}

// FCTokenBin is the compiled bytecode used for deploying new contracts.
var FCTokenBin = "0x60806040523480156200001157600080fd5b50604051620019373803806200193783398101604081905262000034916200021a565b835162000049906000906020870190620000c1565b5082516200005f906001906020860190620000c1565b506002805460ff191660ff84169081179091556200007f90600a620002ee565b6200008b9082620003e3565b600381905533600081815260056020526040902091909155600480546001600160a01b0319169091179055506200046e92505050565b828054620000cf9062000405565b90600052602060002090601f016020900481019282620000f357600085556200013e565b82601f106200010e57805160ff19168380011785556200013e565b828001600101855582156200013e579182015b828111156200013e57825182559160200191906001019062000121565b506200014c92915062000150565b5090565b5b808211156200014c576000815560010162000151565b600082601f83011262000178578081fd5b81516001600160401b038082111562000195576200019562000458565b604051601f8301601f19908116603f01168101908282118183101715620001c057620001c062000458565b81604052838152602092508683858801011115620001dc578485fd5b8491505b83821015620001ff5785820183015181830184015290820190620001e0565b838211156200021057848385830101525b9695505050505050565b6000806000806080858703121562000230578384fd5b84516001600160401b038082111562000247578586fd5b620002558883890162000167565b955060208701519150808211156200026b578485fd5b506200027a8782880162000167565b935050604085015160ff8116811462000291578283fd5b6060959095015193969295505050565b80825b6001808611620002b55750620002e5565b818704821115620002ca57620002ca62000442565b80861615620002d857918102915b9490941c938002620002a4565b94509492505050565b6000620002ff600019848462000306565b9392505050565b6000826200031757506001620002ff565b816200032657506000620002ff565b81600181146200033f57600281146200034a576200037e565b6001915050620002ff565b60ff8411156200035e576200035e62000442565b6001841b91508482111562000377576200037762000442565b50620002ff565b5060208310610133831016604e8410600b8410161715620003b6575081810a83811115620003b057620003b062000442565b620002ff565b620003c58484846001620002a1565b808604821115620003da57620003da62000442565b02949350505050565b600081600019048311821515161562000400576200040062000442565b500290565b600181811c908216806200041a57607f821691505b602082108114156200043c57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6114b9806200047e6000396000f3fe608060405234801561001057600080fd5b506004361061018e5760003560e01c8063615715a2116100de57806387019cd41161009757806395d89b411161007157806395d89b41146103fc578063a9059cbb14610404578063a9dece5b14610417578063cd4217c11461042a5761018e565b806387019cd4146103b15780638da5cb5b146103be578063908e049b146103e95761018e565b8063615715a2146102d25780636c19e783146102e557806370a08231146102f857806375247a581461031857806378e8691d1461032b5780637b46b80b1461039e5761018e565b806323b872dd1161014b5780633458810a116101255780633458810a1461029157806343272bc0146102a4578063519ed728146102b7578063604374f0146102bf5761018e565b806323b872dd1461024e57806324bce60c14610261578063313ce567146102845761018e565b80630307892a146101935780630536b723146101d157806306fdde03146102065780630e316ab71461021b57806312065fe01461023057806318160ddd14610245575b600080fd5b6101be6101a1366004611155565b600660209081526000928352604080842090915290825290205481565b6040519081526020015b60405180910390f35b6101f46101df366004611134565b60076020526000908152604090205460ff1681565b60405160ff90911681526020016101c8565b61020e61044a565b6040516101c891906112e5565b61022e610229366004611134565b6104d8565b005b336000908152600560205260409020546101be565b6101be60035481565b61022e61025c366004611187565b610510565b61027461026f3660046111c2565b6105b2565b60405190151581526020016101c8565b6002546101f49060ff1681565b61022e61029f3660046112c4565b6106c4565b6102746102b23660046111eb565b610842565b6009546101be565b61022e6102cd3660046111c2565b6109be565b6101be6102e03660046112c4565b610b0c565b61022e6102f3366004611134565b610b2a565b6101be610306366004611134565b60056020526000908152604090205481565b61022e6103263660046111c2565b610b65565b6103706103393660046112ac565b600b6020526000908152604090208054600182015460028301546003909301546001600160a01b0392831693919092169160ff1684565b604080516001600160a01b0395861681529490931660208501529183015260ff1660608201526080016101c8565b6102746103ac3660046111c2565b610c34565b600a546101f49060ff1681565b6004546103d1906001600160a01b031681565b6040516001600160a01b0390911681526020016101c8565b61022e6103f73660046112ac565b610d36565b61020e610dcb565b61022e6104123660046111c2565b61018e565b61022e6104253660046111c2565b610dd8565b6101be610438366004611134565b60086020526000908152604090205481565b60008054610457906113e1565b80601f0160208091040260200160405190810160405280929190818152602001828054610483906113e1565b80156104d05780601f106104a5576101008083540402835291602001916104d0565b820191906000526020600020905b8154815290600101906020018083116104b357829003601f168201915b505050505081565b6004546001600160a01b031633146104ef57600080fd5b6001600160a01b03166000908152600760205260409020805460ff19169055565b6001600160a01b0383161580159061053057506001600160a01b03821615155b61053957600080fd5b6001600160a01b038316600090815260066020908152604080832033845290915290205481111561056957600080fd5b6001600160a01b03831660009081526006602090815260408083203384529091528120805483929061059c9084906113ca565b909155506105ad9050838383610ee1565b505050565b6004546000906001600160a01b031633146105cc57600080fd5b6001600160a01b0383166105df57600080fd5b600082116105ec57600080fd5b6001600160a01b038316600090815260056020526040902054821061061057600080fd5b6001600160a01b0383166000908152600560205260409020546106349083906113ca565b6001600160a01b03841660009081526005602090815260408083209390935560089052205461066490839061138d565b6001600160a01b038416600081815260086020526040908190209290925590517ff97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e0906106b39085815260200190565b60405180910390a250600192915050565b6004546001600160a01b03163314806106ef57503360009081526007602052604090205460ff166001145b6106f857600080fd5b60ff81166000908152600b6020526040902080546001600160a01b031661071e57600080fd5b80546001600160a01b031633141561073557600080fd5b33600090815260048201602052604090205460ff166001141561075757600080fd5b3360009081526004820160205260408120805460ff1916600117905560038201805460ff169161078683611437565b82546101009290920a60ff81810219909316918316021790915560038301546002911610905061083e5760028101543060009081526005602052604090205410156107d057600080fd5b600181015460028201546107f19130916001600160a01b0390911690610ee1565b6107fd6009600061106d565b60ff82166000908152600b6020526040812080546001600160a01b0319908116825560018201805490911690556002810191909155600301805460ff191690555b5050565b3360009081526007602052604081205460ff166001148061086d57506004546001600160a01b031633145b61087657600080fd5b60008251118015610888575060008351115b61089157600080fd5b60005b83518110156109b4578281815181106108bd57634e487b7160e01b600052603260045260246000fd5b602002602001015160056000336001600160a01b03166001600160a01b031681526020019081526020016000205410156108f657600080fd5b82818151811061091657634e487b7160e01b600052603260045260246000fd5b602002602001015160066000336001600160a01b03166001600160a01b03168152602001908152602001600020600086848151811061096557634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020600082825461099c919061138d565b909155508190506109ac8161141c565b915050610894565b5060019392505050565b6001600160a01b0382166109d157600080fd5b336000908152600560205260409020548111156109ed57600080fd5b6109f8333083610ee1565b600a54610a099060ff1660016113a5565b600a805460ff92831660ff19918216811783556000908152600b6020908152604080832080546001600160a01b0319908116339081178355600183810180546001600160a01b038e16941684179055600284018b905560038401805490981690975587546009805498890181559096528487047f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af018054968a16601f9098166101000a978802978a02199096169690961790945594548151948552918401929092529082018590529092166060830152907fe1dd58c97af942b2e36ca0fc977e17197d6a8008fda5381780e291337e8ac9959060800160405180910390a1505050565b60ff8082166000908152600b6020526040902060030154165b919050565b6004546001600160a01b03163314610b4157600080fd5b6001600160a01b03166000908152600760205260409020805460ff19166001179055565b3360009081526007602052604090205460ff1660011480610b9057506004546001600160a01b031633145b610b9957600080fd5b6001600160a01b038216610bac57600080fd5b33600090815260056020526040902054811115610bc857600080fd5b3360008181526006602090815260408083206001600160a01b0387168085529083529281902085905580519384529083019190915281018290527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925906060015b60405180910390a15050565b6004546000906001600160a01b03163314610c4e57600080fd5b6001600160a01b038316610c6157600080fd5b6001600160a01b038316600090815260086020526040902054821115610c8657600080fd5b60008211610c9357600080fd5b6001600160a01b038316600090815260086020526040902054610cb79083906113ca565b6001600160a01b038416600090815260086020908152604080832093909355600590522054610ce790839061138d565b6001600160a01b038416600081815260056020526040908190209290925590517f2cfce4af01bcb9d6cf6c84ee1b7c491100b8695368264146a94d71e10a63083f906106b39085815260200190565b33600090815260056020526040902054811115610d5257600080fd5b3360009081526005602052604081208054839290610d719084906113ca565b925050819055508060036000828254610d8a91906113ca565b909155505060408051338152602081018390527f34609bf2c0fbf14a094ca0db1f52ae152d2d5b98b3ed6205eefe08674bed81b2910160405180910390a150565b60018054610457906113e1565b6001600160a01b0382166000908152600660209081526040808320338452909152902054811115610e0857600080fd5b6001600160a01b038216600090815260056020526040902054811115610e2d57600080fd5b6001600160a01b03821660009081526005602052604081208054839290610e559084906113ca565b90915550506001600160a01b038216600090815260066020908152604080832033845290915281208054839290610e8d9084906113ca565b925050819055508060036000828254610ea691906113ca565b909155505060408051338152602081018390527f34609bf2c0fbf14a094ca0db1f52ae152d2d5b98b3ed6205eefe08674bed81b29101610c28565b6001600160a01b038316600090815260056020526040902054811115610f0657600080fd5b6001600160a01b038216610f1957600080fd5b6001600160a01b038216600090815260056020526040902054610f3c828261138d565b11610f4657600080fd5b6001600160a01b038083166000908152600560205260408082205492861682528120549091610f749161138d565b6001600160a01b038516600090815260056020526040812080549293508492909190610fa19084906113ca565b90915550506001600160a01b03831660009081526005602052604081208054849290610fce90849061138d565b9091555050604080516001600160a01b038087168252851660208201529081018390527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060600160405180910390a16001600160a01b03808416600090815260056020526040808220549287168252902054829161104c9161138d565b1461106757634e487b7160e01b600052600160045260246000fd5b50505050565b50805460008255601f0160209004906000526020600020908101906110929190611095565b50565b5b808211156110aa5760008155600101611096565b5090565b80356001600160a01b0381168114610b2557600080fd5b600082601f8301126110d5578081fd5b813560206110ea6110e583611369565b611338565b80838252828201915082860187848660051b8901011115611109578586fd5b855b858110156111275781358452928401929084019060010161110b565b5090979650505050505050565b600060208284031215611145578081fd5b61114e826110ae565b9392505050565b60008060408385031215611167578081fd5b611170836110ae565b915061117e602084016110ae565b90509250929050565b60008060006060848603121561119b578081fd5b6111a4846110ae565b92506111b2602085016110ae565b9150604084013590509250925092565b600080604083850312156111d4578182fd5b6111dd836110ae565b946020939093013593505050565b600080604083850312156111fd578182fd5b823567ffffffffffffffff80821115611214578384fd5b818501915085601f830112611227578384fd5b813560206112376110e583611369565b8083825282820191508286018a848660051b8901011115611256578889fd5b8896505b8487101561127f5761126b816110ae565b83526001969096019591830191830161125a565b5096505086013592505080821115611295578283fd5b506112a2858286016110c5565b9150509250929050565b6000602082840312156112bd578081fd5b5035919050565b6000602082840312156112d5578081fd5b813560ff8116811461114e578182fd5b6000602080835283518082850152825b81811015611311578581018301518582016040015282016112f5565b818111156113225783604083870101525b50601f01601f1916929092016040019392505050565b604051601f8201601f1916810167ffffffffffffffff811182821017156113615761136161146d565b604052919050565b600067ffffffffffffffff8211156113835761138361146d565b5060051b60200190565b600082198211156113a0576113a0611457565b500190565b600060ff821660ff84168060ff038211156113c2576113c2611457565b019392505050565b6000828210156113dc576113dc611457565b500390565b600181811c908216806113f557607f821691505b6020821081141561141657634e487b7160e01b600052602260045260246000fd5b50919050565b600060001982141561143057611430611457565b5060010190565b600060ff821660ff81141561144e5761144e611457565b60010192915050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fdfea2646970667358221220042086995b98e799717743f67a7a9c71e1c7a97e157ae7f7ab6a965901a2ca5564736f6c63430008030033"

// DeployFCToken deploys a new Ethereum contract, binding an instance of FCToken to it.
func DeployFCToken(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, _decimals uint8, _initiSupply *big.Int) (common.Address, *types.Transaction, *FCToken, error) {
	parsed, err := abi.JSON(strings.NewReader(FCTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FCTokenBin), backend, _name, _symbol, _decimals, _initiSupply)
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

// ApprovalBanlance is a free data retrieval call binding the contract method 0x0307892a.
//
// Solidity: function approvalBanlance(address , address ) view returns(uint256)
func (_FCToken *FCTokenCaller) ApprovalBanlance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FCToken.contract.Call(opts, &out, "approvalBanlance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApprovalBanlance is a free data retrieval call binding the contract method 0x0307892a.
//
// Solidity: function approvalBanlance(address , address ) view returns(uint256)
func (_FCToken *FCTokenSession) ApprovalBanlance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _FCToken.Contract.ApprovalBanlance(&_FCToken.CallOpts, arg0, arg1)
}

// ApprovalBanlance is a free data retrieval call binding the contract method 0x0307892a.
//
// Solidity: function approvalBanlance(address , address ) view returns(uint256)
func (_FCToken *FCTokenCallerSession) ApprovalBanlance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _FCToken.Contract.ApprovalBanlance(&_FCToken.CallOpts, arg0, arg1)
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

// FreezeOf is a free data retrieval call binding the contract method 0xcd4217c1.
//
// Solidity: function freezeOf(address ) view returns(uint256)
func (_FCToken *FCTokenCaller) FreezeOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FCToken.contract.Call(opts, &out, "freezeOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FreezeOf is a free data retrieval call binding the contract method 0xcd4217c1.
//
// Solidity: function freezeOf(address ) view returns(uint256)
func (_FCToken *FCTokenSession) FreezeOf(arg0 common.Address) (*big.Int, error) {
	return _FCToken.Contract.FreezeOf(&_FCToken.CallOpts, arg0)
}

// FreezeOf is a free data retrieval call binding the contract method 0xcd4217c1.
//
// Solidity: function freezeOf(address ) view returns(uint256)
func (_FCToken *FCTokenCallerSession) FreezeOf(arg0 common.Address) (*big.Int, error) {
	return _FCToken.Contract.FreezeOf(&_FCToken.CallOpts, arg0)
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

// Approval is a paid mutator transaction binding the contract method 0x75247a58.
//
// Solidity: function approval(address _delegatee, uint256 _value) returns()
func (_FCToken *FCTokenTransactor) Approval(opts *bind.TransactOpts, _delegatee common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FCToken.contract.Transact(opts, "approval", _delegatee, _value)
}

// Approval is a paid mutator transaction binding the contract method 0x75247a58.
//
// Solidity: function approval(address _delegatee, uint256 _value) returns()
func (_FCToken *FCTokenSession) Approval(_delegatee common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FCToken.Contract.Approval(&_FCToken.TransactOpts, _delegatee, _value)
}

// Approval is a paid mutator transaction binding the contract method 0x75247a58.
//
// Solidity: function approval(address _delegatee, uint256 _value) returns()
func (_FCToken *FCTokenTransactorSession) Approval(_delegatee common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FCToken.Contract.Approval(&_FCToken.TransactOpts, _delegatee, _value)
}

// ApproveWthArray is a paid mutator transaction binding the contract method 0x43272bc0.
//
// Solidity: function approveWthArray(address[] _spenders, uint256[] _values) returns(bool success)
func (_FCToken *FCTokenTransactor) ApproveWthArray(opts *bind.TransactOpts, _spenders []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _FCToken.contract.Transact(opts, "approveWthArray", _spenders, _values)
}

// ApproveWthArray is a paid mutator transaction binding the contract method 0x43272bc0.
//
// Solidity: function approveWthArray(address[] _spenders, uint256[] _values) returns(bool success)
func (_FCToken *FCTokenSession) ApproveWthArray(_spenders []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _FCToken.Contract.ApproveWthArray(&_FCToken.TransactOpts, _spenders, _values)
}

// ApproveWthArray is a paid mutator transaction binding the contract method 0x43272bc0.
//
// Solidity: function approveWthArray(address[] _spenders, uint256[] _values) returns(bool success)
func (_FCToken *FCTokenTransactorSession) ApproveWthArray(_spenders []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _FCToken.Contract.ApproveWthArray(&_FCToken.TransactOpts, _spenders, _values)
}

// Destory is a paid mutator transaction binding the contract method 0x908e049b.
//
// Solidity: function destory(uint256 _value) returns()
func (_FCToken *FCTokenTransactor) Destory(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _FCToken.contract.Transact(opts, "destory", _value)
}

// Destory is a paid mutator transaction binding the contract method 0x908e049b.
//
// Solidity: function destory(uint256 _value) returns()
func (_FCToken *FCTokenSession) Destory(_value *big.Int) (*types.Transaction, error) {
	return _FCToken.Contract.Destory(&_FCToken.TransactOpts, _value)
}

// Destory is a paid mutator transaction binding the contract method 0x908e049b.
//
// Solidity: function destory(uint256 _value) returns()
func (_FCToken *FCTokenTransactorSession) Destory(_value *big.Int) (*types.Transaction, error) {
	return _FCToken.Contract.Destory(&_FCToken.TransactOpts, _value)
}

// DestoryFrom is a paid mutator transaction binding the contract method 0xa9dece5b.
//
// Solidity: function destoryFrom(address _from, uint256 _value) returns()
func (_FCToken *FCTokenTransactor) DestoryFrom(opts *bind.TransactOpts, _from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FCToken.contract.Transact(opts, "destoryFrom", _from, _value)
}

// DestoryFrom is a paid mutator transaction binding the contract method 0xa9dece5b.
//
// Solidity: function destoryFrom(address _from, uint256 _value) returns()
func (_FCToken *FCTokenSession) DestoryFrom(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FCToken.Contract.DestoryFrom(&_FCToken.TransactOpts, _from, _value)
}

// DestoryFrom is a paid mutator transaction binding the contract method 0xa9dece5b.
//
// Solidity: function destoryFrom(address _from, uint256 _value) returns()
func (_FCToken *FCTokenTransactorSession) DestoryFrom(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FCToken.Contract.DestoryFrom(&_FCToken.TransactOpts, _from, _value)
}

// Freeze is a paid mutator transaction binding the contract method 0x24bce60c.
//
// Solidity: function freeze(address who, uint256 _value) returns(bool success)
func (_FCToken *FCTokenTransactor) Freeze(opts *bind.TransactOpts, who common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FCToken.contract.Transact(opts, "freeze", who, _value)
}

// Freeze is a paid mutator transaction binding the contract method 0x24bce60c.
//
// Solidity: function freeze(address who, uint256 _value) returns(bool success)
func (_FCToken *FCTokenSession) Freeze(who common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FCToken.Contract.Freeze(&_FCToken.TransactOpts, who, _value)
}

// Freeze is a paid mutator transaction binding the contract method 0x24bce60c.
//
// Solidity: function freeze(address who, uint256 _value) returns(bool success)
func (_FCToken *FCTokenTransactorSession) Freeze(who common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FCToken.Contract.Freeze(&_FCToken.TransactOpts, who, _value)
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

// Unfreeze is a paid mutator transaction binding the contract method 0x7b46b80b.
//
// Solidity: function unfreeze(address who, uint256 _value) returns(bool success)
func (_FCToken *FCTokenTransactor) Unfreeze(opts *bind.TransactOpts, who common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FCToken.contract.Transact(opts, "unfreeze", who, _value)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x7b46b80b.
//
// Solidity: function unfreeze(address who, uint256 _value) returns(bool success)
func (_FCToken *FCTokenSession) Unfreeze(who common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FCToken.Contract.Unfreeze(&_FCToken.TransactOpts, who, _value)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x7b46b80b.
//
// Solidity: function unfreeze(address who, uint256 _value) returns(bool success)
func (_FCToken *FCTokenTransactorSession) Unfreeze(who common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FCToken.Contract.Unfreeze(&_FCToken.TransactOpts, who, _value)
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
// Solidity: event Approval(address from, address delegatee, uint256 value)
func (_FCToken *FCTokenFilterer) FilterApproval(opts *bind.FilterOpts) (*FCTokenApprovalIterator, error) {

	logs, sub, err := _FCToken.contract.FilterLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return &FCTokenApprovalIterator{contract: _FCToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address from, address delegatee, uint256 value)
func (_FCToken *FCTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FCTokenApproval) (event.Subscription, error) {

	logs, sub, err := _FCToken.contract.WatchLogs(opts, "Approval")
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
// Solidity: event Approval(address from, address delegatee, uint256 value)
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

// FCTokenDestoryIterator is returned from FilterDestory and is used to iterate over the raw logs and unpacked data for Destory events raised by the FCToken contract.
type FCTokenDestoryIterator struct {
	Event *FCTokenDestory // Event containing the contract specifics and raw log

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
func (it *FCTokenDestoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FCTokenDestory)
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
		it.Event = new(FCTokenDestory)
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
func (it *FCTokenDestoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FCTokenDestoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FCTokenDestory represents a Destory event raised by the FCToken contract.
type FCTokenDestory struct {
	Destorys common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDestory is a free log retrieval operation binding the contract event 0x34609bf2c0fbf14a094ca0db1f52ae152d2d5b98b3ed6205eefe08674bed81b2.
//
// Solidity: event Destory(address destorys, uint256 value)
func (_FCToken *FCTokenFilterer) FilterDestory(opts *bind.FilterOpts) (*FCTokenDestoryIterator, error) {

	logs, sub, err := _FCToken.contract.FilterLogs(opts, "Destory")
	if err != nil {
		return nil, err
	}
	return &FCTokenDestoryIterator{contract: _FCToken.contract, event: "Destory", logs: logs, sub: sub}, nil
}

// WatchDestory is a free log subscription operation binding the contract event 0x34609bf2c0fbf14a094ca0db1f52ae152d2d5b98b3ed6205eefe08674bed81b2.
//
// Solidity: event Destory(address destorys, uint256 value)
func (_FCToken *FCTokenFilterer) WatchDestory(opts *bind.WatchOpts, sink chan<- *FCTokenDestory) (event.Subscription, error) {

	logs, sub, err := _FCToken.contract.WatchLogs(opts, "Destory")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FCTokenDestory)
				if err := _FCToken.contract.UnpackLog(event, "Destory", log); err != nil {
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

// ParseDestory is a log parse operation binding the contract event 0x34609bf2c0fbf14a094ca0db1f52ae152d2d5b98b3ed6205eefe08674bed81b2.
//
// Solidity: event Destory(address destorys, uint256 value)
func (_FCToken *FCTokenFilterer) ParseDestory(log types.Log) (*FCTokenDestory, error) {
	event := new(FCTokenDestory)
	if err := _FCToken.contract.UnpackLog(event, "Destory", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FCTokenFreezeIterator is returned from FilterFreeze and is used to iterate over the raw logs and unpacked data for Freeze events raised by the FCToken contract.
type FCTokenFreezeIterator struct {
	Event *FCTokenFreeze // Event containing the contract specifics and raw log

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
func (it *FCTokenFreezeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FCTokenFreeze)
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
		it.Event = new(FCTokenFreeze)
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
func (it *FCTokenFreezeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FCTokenFreezeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FCTokenFreeze represents a Freeze event raised by the FCToken contract.
type FCTokenFreeze struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterFreeze is a free log retrieval operation binding the contract event 0xf97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e0.
//
// Solidity: event Freeze(address indexed from, uint256 value)
func (_FCToken *FCTokenFilterer) FilterFreeze(opts *bind.FilterOpts, from []common.Address) (*FCTokenFreezeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _FCToken.contract.FilterLogs(opts, "Freeze", fromRule)
	if err != nil {
		return nil, err
	}
	return &FCTokenFreezeIterator{contract: _FCToken.contract, event: "Freeze", logs: logs, sub: sub}, nil
}

// WatchFreeze is a free log subscription operation binding the contract event 0xf97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e0.
//
// Solidity: event Freeze(address indexed from, uint256 value)
func (_FCToken *FCTokenFilterer) WatchFreeze(opts *bind.WatchOpts, sink chan<- *FCTokenFreeze, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _FCToken.contract.WatchLogs(opts, "Freeze", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FCTokenFreeze)
				if err := _FCToken.contract.UnpackLog(event, "Freeze", log); err != nil {
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

// ParseFreeze is a log parse operation binding the contract event 0xf97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e0.
//
// Solidity: event Freeze(address indexed from, uint256 value)
func (_FCToken *FCTokenFilterer) ParseFreeze(log types.Log) (*FCTokenFreeze, error) {
	event := new(FCTokenFreeze)
	if err := _FCToken.contract.UnpackLog(event, "Freeze", log); err != nil {
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
// Solidity: event Transfer(address from, address to, uint256 value)
func (_FCToken *FCTokenFilterer) FilterTransfer(opts *bind.FilterOpts) (*FCTokenTransferIterator, error) {

	logs, sub, err := _FCToken.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &FCTokenTransferIterator{contract: _FCToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address from, address to, uint256 value)
func (_FCToken *FCTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FCTokenTransfer) (event.Subscription, error) {

	logs, sub, err := _FCToken.contract.WatchLogs(opts, "Transfer")
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
// Solidity: event Transfer(address from, address to, uint256 value)
func (_FCToken *FCTokenFilterer) ParseTransfer(log types.Log) (*FCTokenTransfer, error) {
	event := new(FCTokenTransfer)
	if err := _FCToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FCTokenUnfreezeIterator is returned from FilterUnfreeze and is used to iterate over the raw logs and unpacked data for Unfreeze events raised by the FCToken contract.
type FCTokenUnfreezeIterator struct {
	Event *FCTokenUnfreeze // Event containing the contract specifics and raw log

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
func (it *FCTokenUnfreezeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FCTokenUnfreeze)
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
		it.Event = new(FCTokenUnfreeze)
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
func (it *FCTokenUnfreezeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FCTokenUnfreezeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FCTokenUnfreeze represents a Unfreeze event raised by the FCToken contract.
type FCTokenUnfreeze struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnfreeze is a free log retrieval operation binding the contract event 0x2cfce4af01bcb9d6cf6c84ee1b7c491100b8695368264146a94d71e10a63083f.
//
// Solidity: event Unfreeze(address indexed from, uint256 value)
func (_FCToken *FCTokenFilterer) FilterUnfreeze(opts *bind.FilterOpts, from []common.Address) (*FCTokenUnfreezeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _FCToken.contract.FilterLogs(opts, "Unfreeze", fromRule)
	if err != nil {
		return nil, err
	}
	return &FCTokenUnfreezeIterator{contract: _FCToken.contract, event: "Unfreeze", logs: logs, sub: sub}, nil
}

// WatchUnfreeze is a free log subscription operation binding the contract event 0x2cfce4af01bcb9d6cf6c84ee1b7c491100b8695368264146a94d71e10a63083f.
//
// Solidity: event Unfreeze(address indexed from, uint256 value)
func (_FCToken *FCTokenFilterer) WatchUnfreeze(opts *bind.WatchOpts, sink chan<- *FCTokenUnfreeze, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _FCToken.contract.WatchLogs(opts, "Unfreeze", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FCTokenUnfreeze)
				if err := _FCToken.contract.UnpackLog(event, "Unfreeze", log); err != nil {
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

// ParseUnfreeze is a log parse operation binding the contract event 0x2cfce4af01bcb9d6cf6c84ee1b7c491100b8695368264146a94d71e10a63083f.
//
// Solidity: event Unfreeze(address indexed from, uint256 value)
func (_FCToken *FCTokenFilterer) ParseUnfreeze(log types.Log) (*FCTokenUnfreeze, error) {
	event := new(FCTokenUnfreeze)
	if err := _FCToken.contract.UnpackLog(event, "Unfreeze", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

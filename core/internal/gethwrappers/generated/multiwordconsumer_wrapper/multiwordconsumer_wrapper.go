// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package multiwordconsumer_wrapper

import (
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated"
)

var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

const MultiWordConsumerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_link\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_specId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"price\",\"type\":\"bytes\"}],\"name\":\"RequestFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"usd\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"eur\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"jpy\",\"type\":\"bytes32\"}],\"name\":\"RequestMultipleFulfilled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"}],\"name\":\"addExternalRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_payment\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"_callbackFunctionId\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"_expiration\",\"type\":\"uint256\"}],\"name\":\"cancelRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPrice\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eur\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_price\",\"type\":\"bytes\"}],\"name\":\"fulfillBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_usd\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_eur\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_jpy\",\"type\":\"bytes32\"}],\"name\":\"fulfillMultipleParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"jpy\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_currency\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_payment\",\"type\":\"uint256\"}],\"name\":\"requestEthereumPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_currency\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_payment\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_callback\",\"type\":\"address\"}],\"name\":\"requestEthereumPriceByCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_currency\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_payment\",\"type\":\"uint256\"}],\"name\":\"requestMultipleParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_urlUSD\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_pathUSD\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_urlEUR\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_pathEUR\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_urlJPY\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_pathJPY\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_payment\",\"type\":\"uint256\"}],\"name\":\"requestMultipleParametersWithCustomURLs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_specId\",\"type\":\"bytes32\"}],\"name\":\"setSpecID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usd\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

var MultiWordConsumerBin = "0x6080604052600160045534801561001557600080fd5b50604051611bce380380611bce8339818101604052606081101561003857600080fd5b508051602082015160409092015190919061005283610066565b61005b82610088565b600655506100aa9050565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b600380546001600160a01b0319166001600160a01b0392909216919091179055565b611b15806100b96000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80639d1b464a1161008c578063d63a6ccd11610066578063d63a6ccd146107aa578063e89855ba146107b2578063e8d5359d1461085a578063faa3676114610893576100ea565b80639d1b464a14610651578063a856ff6b146106ce578063c2fb8523146106fd576100ea565b806374961d4d116100c857806374961d4d1461019557806383db5cbc146102565780638dc654a2146102fe578063938649e514610306576100ea565b8063501fdd5d146100ef5780635591a6081461010e5780637439ae591461017b575b600080fd5b61010c6004803603602081101561010557600080fd5b503561089b565b005b61010c600480360360a081101561012457600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813516906020810135906040810135907fffffffff0000000000000000000000000000000000000000000000000000000060608201351690608001356108a0565b610183610967565b60408051918252519081900360200190f35b61010c600480360360608110156101ab57600080fd5b8101906020810181356401000000008111156101c657600080fd5b8201836020820111156101d857600080fd5b803590602001918460018302840111640100000000831117156101fa57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020013573ffffffffffffffffffffffffffffffffffffffff1661096d565b61010c6004803603604081101561026c57600080fd5b81019060208101813564010000000081111561028757600080fd5b82018360208201111561029957600080fd5b803590602001918460018302840111640100000000831117156102bb57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610996915050565b61010c6109a5565b61010c600480360360e081101561031c57600080fd5b81019060208101813564010000000081111561033757600080fd5b82018360208201111561034957600080fd5b8035906020019184600183028401116401000000008311171561036b57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156103be57600080fd5b8201836020820111156103d057600080fd5b803590602001918460018302840111640100000000831117156103f257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561044557600080fd5b82018360208201111561045757600080fd5b8035906020019184600183028401116401000000008311171561047957600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156104cc57600080fd5b8201836020820111156104de57600080fd5b8035906020019184600183028401116401000000008311171561050057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561055357600080fd5b82018360208201111561056557600080fd5b8035906020019184600183028401116401000000008311171561058757600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156105da57600080fd5b8201836020820111156105ec57600080fd5b8035906020019184600183028401116401000000008311171561060e57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610b6f915050565b610659610d1d565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561069357818101518382015260200161067b565b50505050905090810190601f1680156106c05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61010c600480360360808110156106e457600080fd5b5080359060208101359060408101359060600135610dc9565b61010c6004803603604081101561071357600080fd5b8135919081019060408101602082013564010000000081111561073557600080fd5b82018360208201111561074757600080fd5b8035906020019184600183028401116401000000008311171561076957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610eec945050505050565b61018361109e565b61010c600480360360408110156107c857600080fd5b8101906020810181356401000000008111156107e357600080fd5b8201836020820111156107f557600080fd5b8035906020019184600183028401116401000000008311171561081757600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506110a4915050565b61010c6004803603604081101561087057600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356110c6565b6101836110d0565b600655565b604080517f6ee4d55300000000000000000000000000000000000000000000000000000000815260048101869052602481018590527fffffffff0000000000000000000000000000000000000000000000000000000084166044820152606481018390529051869173ffffffffffffffffffffffffffffffffffffffff831691636ee4d5539160848082019260009290919082900301818387803b15801561094757600080fd5b505af115801561095b573d6000803e3d6000fd5b50505050505050505050565b60095481565b60006109836006548363c2fb852360e01b6110d6565b905061098f81846110fb565b5050505050565b6109a182823061096d565b5050565b60006109af61112b565b90508073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb338373ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015610a3557600080fd5b505afa158015610a49573d6000803e3d6000fd5b505050506040513d6020811015610a5f57600080fd5b5051604080517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b16815273ffffffffffffffffffffffffffffffffffffffff909316600484015260248301919091525160448083019260209291908290030181600087803b158015610ad557600080fd5b505af1158015610ae9573d6000803e3d6000fd5b505050506040513d6020811015610aff57600080fd5b5051610b6c57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f556e61626c6520746f207472616e736665720000000000000000000000000000604482015290519081900360640190fd5b50565b6000610b856006543063a856ff6b60e01b6110d6565b60408051808201909152600681527f75726c55534400000000000000000000000000000000000000000000000000006020820152909150610bc89082908a611147565b60408051808201909152600781527f70617468555344000000000000000000000000000000000000000000000000006020820152610c0890829089611147565b60408051808201909152600681527f75726c45555200000000000000000000000000000000000000000000000000006020820152610c4890829088611147565b60408051808201909152600781527f70617468455552000000000000000000000000000000000000000000000000006020820152610c8890829087611147565b60408051808201909152600681527f75726c4a505900000000000000000000000000000000000000000000000000006020820152610cc890829086611147565b60408051808201909152600781527f706174684a5059000000000000000000000000000000000000000000000000006020820152610d0890829085611147565b610d1281836110fb565b505050505050505050565b6007805460408051602060026001851615610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190941693909304601f81018490048402820184019092528181529291830182828015610dc15780601f10610d9657610100808354040283529160200191610dc1565b820191906000526020600020905b815481529060010190602001808311610da457829003601f168201915b505050505081565b600084815260056020526040902054849073ffffffffffffffffffffffffffffffffffffffff163314610e47576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526028815260200180611ae16028913960400191505060405180910390fd5b60008181526005602052604080822080547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555182917f7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a91a28284867f0ec0c13e44aa04198947078cb990660252870dd3363f4c4bb3cc780f808dabbe856040518082815260200191505060405180910390a450600892909255600955600a5550565b600082815260056020526040902054829073ffffffffffffffffffffffffffffffffffffffff163314610f6a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526028815260200180611ae16028913960400191505060405180910390fd5b60008181526005602052604080822080547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555182917f7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a91a2816040518082805190602001908083835b6020831061101357805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe09092019160209182019101610fd6565b5181516020939093036101000a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01801990911692169190911790526040519201829003822093508692507f1a111c5dcf9a71088bd5e1797fdfaf399fec2afbb24aca247e4e3e9f4b61df919160009150a381516110989060079060208501906119cd565b50505050565b60085481565b60006110ba6006543063a856ff6b60e01b6110d6565b905061109881836110fb565b6109a1828261116a565b600a5481565b6110de611a59565b6110e6611a59565b6110f281868686611251565b95945050505050565b6003546000906111229073ffffffffffffffffffffffffffffffffffffffff1684846112b3565b90505b92915050565b60025473ffffffffffffffffffffffffffffffffffffffff1690565b608083015161115690836112eb565b608083015161116590826112eb565b505050565b600081815260056020526040902054819073ffffffffffffffffffffffffffffffffffffffff16156111fd57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f5265717565737420697320616c72656164792070656e64696e67000000000000604482015290519081900360640190fd5b50600090815260056020526040902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b611259611a59565b6112698560800151610100611302565b505091835273ffffffffffffffffffffffffffffffffffffffff1660208301527fffffffff0000000000000000000000000000000000000000000000000000000016604082015290565b60006112e384848460027f6de879d60000000000000000000000000000000000000000000000000000000061133c565b949350505050565b6112f882600383516116ef565b61116582826117c3565b61130a611a8e565b602082061561131f5760208206602003820191505b506020828101829052604080518085526000815290920101905290565b6004546040805130606090811b60208084019190915260348084018690528451808503909101815260549093018452825192810192909220908801939093526000838152600590915281812080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8a16179055905182917fb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af991a2600082600080886000015189602001518a604001518b606001518a8d6080015160000151604051602401808973ffffffffffffffffffffffffffffffffffffffff1681526020018881526020018781526020018673ffffffffffffffffffffffffffffffffffffffff168152602001857bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156114c65781810151838201526020016114ae565b50505050905090810190601f1680156114f35780820380516001836020036101000a031916815260200191505b509950505050505050505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050509050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634000aea08887846040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561160d5781810151838201526020016115f5565b50505050905090810190601f16801561163a5780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b15801561165b57600080fd5b505af115801561166f573d6000803e3d6000fd5b505050506040513d602081101561168557600080fd5b50516116dc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180611abe6023913960400191505060405180910390fd5b5060048054600101905595945050505050565b601781116117105761170a8360e0600585901b1683176117dd565b50611165565b60ff811161173a5761172d836018611fe0600586901b16176117dd565b5061170a838260016117f5565b61ffff811161176557611758836019611fe0600586901b16176117dd565b5061170a838260026117f5565b63ffffffff81116117925761178583601a611fe0600586901b16176117dd565b5061170a838260046117f5565b67ffffffffffffffff8111611165576117b683601b611fe0600586901b16176117dd565b50611098838260086117f5565b6117cb611a8e565b6111228384600001515184855161180e565b6117e5611a8e565b61112283846000015151846118f6565b6117fd611a8e565b6112e3848560000151518585611941565b611816611a8e565b825182111561182457600080fd5b8460200151828501111561184e5761184e85611846876020015187860161199f565b6002026119b6565b60008086518051876020830101935080888701111561186d5787860182525b505050602084015b602084106118b257805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe09093019260209182019101611875565b5181517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60208690036101000a019081169019919091161790525083949350505050565b6118fe611a8e565b8360200151831061191a5761191a8485602001516002026119b6565b835180516020858301018481535080851415611937576001810182525b5093949350505050565b611949611a8e565b8460200151848301111561196657611966858584016002026119b6565b60006001836101000a0390508551838682010185831982511617815250805184870111156119945783860181525b509495945050505050565b6000818311156119b0575081611125565b50919050565b81516119c28383611302565b5061109883826117c3565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282611a035760008555611a49565b82601f10611a1c57805160ff1916838001178555611a49565b82800160010185558215611a49579182015b82811115611a49578251825591602001919060010190611a2e565b50611a55929150611aa8565b5090565b6040805160a081018252600080825260208201819052918101829052606081019190915260808101611a89611a8e565b905290565b604051806040016040528060608152602001600081525090565b5b80821115611a555760008155600101611aa956fe756e61626c6520746f207472616e73666572416e6443616c6c20746f206f7261636c65536f75726365206d75737420626520746865206f7261636c65206f66207468652072657175657374a164736f6c6343000706000a"

func DeployMultiWordConsumer(auth *bind.TransactOpts, backend bind.ContractBackend, _link common.Address, _oracle common.Address, _specId [32]byte) (common.Address, *types.Transaction, *MultiWordConsumer, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiWordConsumerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MultiWordConsumerBin), backend, _link, _oracle, _specId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MultiWordConsumer{MultiWordConsumerCaller: MultiWordConsumerCaller{contract: contract}, MultiWordConsumerTransactor: MultiWordConsumerTransactor{contract: contract}, MultiWordConsumerFilterer: MultiWordConsumerFilterer{contract: contract}}, nil
}

type MultiWordConsumer struct {
	address common.Address
	abi     abi.ABI
	MultiWordConsumerCaller
	MultiWordConsumerTransactor
	MultiWordConsumerFilterer
}

type MultiWordConsumerCaller struct {
	contract *bind.BoundContract
}

type MultiWordConsumerTransactor struct {
	contract *bind.BoundContract
}

type MultiWordConsumerFilterer struct {
	contract *bind.BoundContract
}

type MultiWordConsumerSession struct {
	Contract     *MultiWordConsumer
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MultiWordConsumerCallerSession struct {
	Contract *MultiWordConsumerCaller
	CallOpts bind.CallOpts
}

type MultiWordConsumerTransactorSession struct {
	Contract     *MultiWordConsumerTransactor
	TransactOpts bind.TransactOpts
}

type MultiWordConsumerRaw struct {
	Contract *MultiWordConsumer
}

type MultiWordConsumerCallerRaw struct {
	Contract *MultiWordConsumerCaller
}

type MultiWordConsumerTransactorRaw struct {
	Contract *MultiWordConsumerTransactor
}

func NewMultiWordConsumer(address common.Address, backend bind.ContractBackend) (*MultiWordConsumer, error) {
	abi, err := abi.JSON(strings.NewReader(MultiWordConsumerABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMultiWordConsumer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiWordConsumer{address: address, abi: abi, MultiWordConsumerCaller: MultiWordConsumerCaller{contract: contract}, MultiWordConsumerTransactor: MultiWordConsumerTransactor{contract: contract}, MultiWordConsumerFilterer: MultiWordConsumerFilterer{contract: contract}}, nil
}

func NewMultiWordConsumerCaller(address common.Address, caller bind.ContractCaller) (*MultiWordConsumerCaller, error) {
	contract, err := bindMultiWordConsumer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiWordConsumerCaller{contract: contract}, nil
}

func NewMultiWordConsumerTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiWordConsumerTransactor, error) {
	contract, err := bindMultiWordConsumer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiWordConsumerTransactor{contract: contract}, nil
}

func NewMultiWordConsumerFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiWordConsumerFilterer, error) {
	contract, err := bindMultiWordConsumer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiWordConsumerFilterer{contract: contract}, nil
}

func bindMultiWordConsumer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiWordConsumerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_MultiWordConsumer *MultiWordConsumerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiWordConsumer.Contract.MultiWordConsumerCaller.contract.Call(opts, result, method, params...)
}

func (_MultiWordConsumer *MultiWordConsumerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiWordConsumer.Contract.MultiWordConsumerTransactor.contract.Transfer(opts)
}

func (_MultiWordConsumer *MultiWordConsumerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiWordConsumer.Contract.MultiWordConsumerTransactor.contract.Transact(opts, method, params...)
}

func (_MultiWordConsumer *MultiWordConsumerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiWordConsumer.Contract.contract.Call(opts, result, method, params...)
}

func (_MultiWordConsumer *MultiWordConsumerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiWordConsumer.Contract.contract.Transfer(opts)
}

func (_MultiWordConsumer *MultiWordConsumerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiWordConsumer.Contract.contract.Transact(opts, method, params...)
}

func (_MultiWordConsumer *MultiWordConsumerCaller) CurrentPrice(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _MultiWordConsumer.contract.Call(opts, &out, "currentPrice")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MultiWordConsumer *MultiWordConsumerSession) CurrentPrice() ([]byte, error) {
	return _MultiWordConsumer.Contract.CurrentPrice(&_MultiWordConsumer.CallOpts)
}

func (_MultiWordConsumer *MultiWordConsumerCallerSession) CurrentPrice() ([]byte, error) {
	return _MultiWordConsumer.Contract.CurrentPrice(&_MultiWordConsumer.CallOpts)
}

func (_MultiWordConsumer *MultiWordConsumerCaller) Eur(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MultiWordConsumer.contract.Call(opts, &out, "eur")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_MultiWordConsumer *MultiWordConsumerSession) Eur() ([32]byte, error) {
	return _MultiWordConsumer.Contract.Eur(&_MultiWordConsumer.CallOpts)
}

func (_MultiWordConsumer *MultiWordConsumerCallerSession) Eur() ([32]byte, error) {
	return _MultiWordConsumer.Contract.Eur(&_MultiWordConsumer.CallOpts)
}

func (_MultiWordConsumer *MultiWordConsumerCaller) Jpy(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MultiWordConsumer.contract.Call(opts, &out, "jpy")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_MultiWordConsumer *MultiWordConsumerSession) Jpy() ([32]byte, error) {
	return _MultiWordConsumer.Contract.Jpy(&_MultiWordConsumer.CallOpts)
}

func (_MultiWordConsumer *MultiWordConsumerCallerSession) Jpy() ([32]byte, error) {
	return _MultiWordConsumer.Contract.Jpy(&_MultiWordConsumer.CallOpts)
}

func (_MultiWordConsumer *MultiWordConsumerCaller) Usd(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MultiWordConsumer.contract.Call(opts, &out, "usd")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_MultiWordConsumer *MultiWordConsumerSession) Usd() ([32]byte, error) {
	return _MultiWordConsumer.Contract.Usd(&_MultiWordConsumer.CallOpts)
}

func (_MultiWordConsumer *MultiWordConsumerCallerSession) Usd() ([32]byte, error) {
	return _MultiWordConsumer.Contract.Usd(&_MultiWordConsumer.CallOpts)
}

func (_MultiWordConsumer *MultiWordConsumerTransactor) AddExternalRequest(opts *bind.TransactOpts, _oracle common.Address, _requestId [32]byte) (*types.Transaction, error) {
	return _MultiWordConsumer.contract.Transact(opts, "addExternalRequest", _oracle, _requestId)
}

func (_MultiWordConsumer *MultiWordConsumerSession) AddExternalRequest(_oracle common.Address, _requestId [32]byte) (*types.Transaction, error) {
	return _MultiWordConsumer.Contract.AddExternalRequest(&_MultiWordConsumer.TransactOpts, _oracle, _requestId)
}

func (_MultiWordConsumer *MultiWordConsumerTransactorSession) AddExternalRequest(_oracle common.Address, _requestId [32]byte) (*types.Transaction, error) {
	return _MultiWordConsumer.Contract.AddExternalRequest(&_MultiWordConsumer.TransactOpts, _oracle, _requestId)
}

func (_MultiWordConsumer *MultiWordConsumerTransactor) CancelRequest(opts *bind.TransactOpts, _oracle common.Address, _requestId [32]byte, _payment *big.Int, _callbackFunctionId [4]byte, _expiration *big.Int) (*types.Transaction, error) {
	return _MultiWordConsumer.contract.Transact(opts, "cancelRequest", _oracle, _requestId, _payment, _callbackFunctionId, _expiration)
}

func (_MultiWordConsumer *MultiWordConsumerSession) CancelRequest(_oracle common.Address, _requestId [32]byte, _payment *big.Int, _callbackFunctionId [4]byte, _expiration *big.Int) (*types.Transaction, error) {
	return _MultiWordConsumer.Contract.CancelRequest(&_MultiWordConsumer.TransactOpts, _oracle, _requestId, _payment, _callbackFunctionId, _expiration)
}

func (_MultiWordConsumer *MultiWordConsumerTransactorSession) CancelRequest(_oracle common.Address, _requestId [32]byte, _payment *big.Int, _callbackFunctionId [4]byte, _expiration *big.Int) (*types.Transaction, error) {
	return _MultiWordConsumer.Contract.CancelRequest(&_MultiWordConsumer.TransactOpts, _oracle, _requestId, _payment, _callbackFunctionId, _expiration)
}

func (_MultiWordConsumer *MultiWordConsumerTransactor) FulfillBytes(opts *bind.TransactOpts, _requestId [32]byte, _price []byte) (*types.Transaction, error) {
	return _MultiWordConsumer.contract.Transact(opts, "fulfillBytes", _requestId, _price)
}

func (_MultiWordConsumer *MultiWordConsumerSession) FulfillBytes(_requestId [32]byte, _price []byte) (*types.Transaction, error) {
	return _MultiWordConsumer.Contract.FulfillBytes(&_MultiWordConsumer.TransactOpts, _requestId, _price)
}

func (_MultiWordConsumer *MultiWordConsumerTransactorSession) FulfillBytes(_requestId [32]byte, _price []byte) (*types.Transaction, error) {
	return _MultiWordConsumer.Contract.FulfillBytes(&_MultiWordConsumer.TransactOpts, _requestId, _price)
}

func (_MultiWordConsumer *MultiWordConsumerTransactor) FulfillMultipleParameters(opts *bind.TransactOpts, _requestId [32]byte, _usd [32]byte, _eur [32]byte, _jpy [32]byte) (*types.Transaction, error) {
	return _MultiWordConsumer.contract.Transact(opts, "fulfillMultipleParameters", _requestId, _usd, _eur, _jpy)
}

func (_MultiWordConsumer *MultiWordConsumerSession) FulfillMultipleParameters(_requestId [32]byte, _usd [32]byte, _eur [32]byte, _jpy [32]byte) (*types.Transaction, error) {
	return _MultiWordConsumer.Contract.FulfillMultipleParameters(&_MultiWordConsumer.TransactOpts, _requestId, _usd, _eur, _jpy)
}

func (_MultiWordConsumer *MultiWordConsumerTransactorSession) FulfillMultipleParameters(_requestId [32]byte, _usd [32]byte, _eur [32]byte, _jpy [32]byte) (*types.Transaction, error) {
	return _MultiWordConsumer.Contract.FulfillMultipleParameters(&_MultiWordConsumer.TransactOpts, _requestId, _usd, _eur, _jpy)
}

func (_MultiWordConsumer *MultiWordConsumerTransactor) RequestEthereumPrice(opts *bind.TransactOpts, _currency string, _payment *big.Int) (*types.Transaction, error) {
	return _MultiWordConsumer.contract.Transact(opts, "requestEthereumPrice", _currency, _payment)
}

func (_MultiWordConsumer *MultiWordConsumerSession) RequestEthereumPrice(_currency string, _payment *big.Int) (*types.Transaction, error) {
	return _MultiWordConsumer.Contract.RequestEthereumPrice(&_MultiWordConsumer.TransactOpts, _currency, _payment)
}

func (_MultiWordConsumer *MultiWordConsumerTransactorSession) RequestEthereumPrice(_currency string, _payment *big.Int) (*types.Transaction, error) {
	return _MultiWordConsumer.Contract.RequestEthereumPrice(&_MultiWordConsumer.TransactOpts, _currency, _payment)
}

func (_MultiWordConsumer *MultiWordConsumerTransactor) RequestEthereumPriceByCallback(opts *bind.TransactOpts, _currency string, _payment *big.Int, _callback common.Address) (*types.Transaction, error) {
	return _MultiWordConsumer.contract.Transact(opts, "requestEthereumPriceByCallback", _currency, _payment, _callback)
}

func (_MultiWordConsumer *MultiWordConsumerSession) RequestEthereumPriceByCallback(_currency string, _payment *big.Int, _callback common.Address) (*types.Transaction, error) {
	return _MultiWordConsumer.Contract.RequestEthereumPriceByCallback(&_MultiWordConsumer.TransactOpts, _currency, _payment, _callback)
}

func (_MultiWordConsumer *MultiWordConsumerTransactorSession) RequestEthereumPriceByCallback(_currency string, _payment *big.Int, _callback common.Address) (*types.Transaction, error) {
	return _MultiWordConsumer.Contract.RequestEthereumPriceByCallback(&_MultiWordConsumer.TransactOpts, _currency, _payment, _callback)
}

func (_MultiWordConsumer *MultiWordConsumerTransactor) RequestMultipleParameters(opts *bind.TransactOpts, _currency string, _payment *big.Int) (*types.Transaction, error) {
	return _MultiWordConsumer.contract.Transact(opts, "requestMultipleParameters", _currency, _payment)
}

func (_MultiWordConsumer *MultiWordConsumerSession) RequestMultipleParameters(_currency string, _payment *big.Int) (*types.Transaction, error) {
	return _MultiWordConsumer.Contract.RequestMultipleParameters(&_MultiWordConsumer.TransactOpts, _currency, _payment)
}

func (_MultiWordConsumer *MultiWordConsumerTransactorSession) RequestMultipleParameters(_currency string, _payment *big.Int) (*types.Transaction, error) {
	return _MultiWordConsumer.Contract.RequestMultipleParameters(&_MultiWordConsumer.TransactOpts, _currency, _payment)
}

func (_MultiWordConsumer *MultiWordConsumerTransactor) RequestMultipleParametersWithCustomURLs(opts *bind.TransactOpts, _urlUSD string, _pathUSD string, _urlEUR string, _pathEUR string, _urlJPY string, _pathJPY string, _payment *big.Int) (*types.Transaction, error) {
	return _MultiWordConsumer.contract.Transact(opts, "requestMultipleParametersWithCustomURLs", _urlUSD, _pathUSD, _urlEUR, _pathEUR, _urlJPY, _pathJPY, _payment)
}

func (_MultiWordConsumer *MultiWordConsumerSession) RequestMultipleParametersWithCustomURLs(_urlUSD string, _pathUSD string, _urlEUR string, _pathEUR string, _urlJPY string, _pathJPY string, _payment *big.Int) (*types.Transaction, error) {
	return _MultiWordConsumer.Contract.RequestMultipleParametersWithCustomURLs(&_MultiWordConsumer.TransactOpts, _urlUSD, _pathUSD, _urlEUR, _pathEUR, _urlJPY, _pathJPY, _payment)
}

func (_MultiWordConsumer *MultiWordConsumerTransactorSession) RequestMultipleParametersWithCustomURLs(_urlUSD string, _pathUSD string, _urlEUR string, _pathEUR string, _urlJPY string, _pathJPY string, _payment *big.Int) (*types.Transaction, error) {
	return _MultiWordConsumer.Contract.RequestMultipleParametersWithCustomURLs(&_MultiWordConsumer.TransactOpts, _urlUSD, _pathUSD, _urlEUR, _pathEUR, _urlJPY, _pathJPY, _payment)
}

func (_MultiWordConsumer *MultiWordConsumerTransactor) SetSpecID(opts *bind.TransactOpts, _specId [32]byte) (*types.Transaction, error) {
	return _MultiWordConsumer.contract.Transact(opts, "setSpecID", _specId)
}

func (_MultiWordConsumer *MultiWordConsumerSession) SetSpecID(_specId [32]byte) (*types.Transaction, error) {
	return _MultiWordConsumer.Contract.SetSpecID(&_MultiWordConsumer.TransactOpts, _specId)
}

func (_MultiWordConsumer *MultiWordConsumerTransactorSession) SetSpecID(_specId [32]byte) (*types.Transaction, error) {
	return _MultiWordConsumer.Contract.SetSpecID(&_MultiWordConsumer.TransactOpts, _specId)
}

func (_MultiWordConsumer *MultiWordConsumerTransactor) WithdrawLink(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiWordConsumer.contract.Transact(opts, "withdrawLink")
}

func (_MultiWordConsumer *MultiWordConsumerSession) WithdrawLink() (*types.Transaction, error) {
	return _MultiWordConsumer.Contract.WithdrawLink(&_MultiWordConsumer.TransactOpts)
}

func (_MultiWordConsumer *MultiWordConsumerTransactorSession) WithdrawLink() (*types.Transaction, error) {
	return _MultiWordConsumer.Contract.WithdrawLink(&_MultiWordConsumer.TransactOpts)
}

type MultiWordConsumerChainlinkCancelledIterator struct {
	Event *MultiWordConsumerChainlinkCancelled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiWordConsumerChainlinkCancelledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiWordConsumerChainlinkCancelled)
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

	select {
	case log := <-it.logs:
		it.Event = new(MultiWordConsumerChainlinkCancelled)
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

func (it *MultiWordConsumerChainlinkCancelledIterator) Error() error {
	return it.fail
}

func (it *MultiWordConsumerChainlinkCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiWordConsumerChainlinkCancelled struct {
	Id  [32]byte
	Raw types.Log
}

func (_MultiWordConsumer *MultiWordConsumerFilterer) FilterChainlinkCancelled(opts *bind.FilterOpts, id [][32]byte) (*MultiWordConsumerChainlinkCancelledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MultiWordConsumer.contract.FilterLogs(opts, "ChainlinkCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return &MultiWordConsumerChainlinkCancelledIterator{contract: _MultiWordConsumer.contract, event: "ChainlinkCancelled", logs: logs, sub: sub}, nil
}

func (_MultiWordConsumer *MultiWordConsumerFilterer) WatchChainlinkCancelled(opts *bind.WatchOpts, sink chan<- *MultiWordConsumerChainlinkCancelled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MultiWordConsumer.contract.WatchLogs(opts, "ChainlinkCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiWordConsumerChainlinkCancelled)
				if err := _MultiWordConsumer.contract.UnpackLog(event, "ChainlinkCancelled", log); err != nil {
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

func (_MultiWordConsumer *MultiWordConsumerFilterer) ParseChainlinkCancelled(log types.Log) (*MultiWordConsumerChainlinkCancelled, error) {
	event := new(MultiWordConsumerChainlinkCancelled)
	if err := _MultiWordConsumer.contract.UnpackLog(event, "ChainlinkCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MultiWordConsumerChainlinkFulfilledIterator struct {
	Event *MultiWordConsumerChainlinkFulfilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiWordConsumerChainlinkFulfilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiWordConsumerChainlinkFulfilled)
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

	select {
	case log := <-it.logs:
		it.Event = new(MultiWordConsumerChainlinkFulfilled)
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

func (it *MultiWordConsumerChainlinkFulfilledIterator) Error() error {
	return it.fail
}

func (it *MultiWordConsumerChainlinkFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiWordConsumerChainlinkFulfilled struct {
	Id  [32]byte
	Raw types.Log
}

func (_MultiWordConsumer *MultiWordConsumerFilterer) FilterChainlinkFulfilled(opts *bind.FilterOpts, id [][32]byte) (*MultiWordConsumerChainlinkFulfilledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MultiWordConsumer.contract.FilterLogs(opts, "ChainlinkFulfilled", idRule)
	if err != nil {
		return nil, err
	}
	return &MultiWordConsumerChainlinkFulfilledIterator{contract: _MultiWordConsumer.contract, event: "ChainlinkFulfilled", logs: logs, sub: sub}, nil
}

func (_MultiWordConsumer *MultiWordConsumerFilterer) WatchChainlinkFulfilled(opts *bind.WatchOpts, sink chan<- *MultiWordConsumerChainlinkFulfilled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MultiWordConsumer.contract.WatchLogs(opts, "ChainlinkFulfilled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiWordConsumerChainlinkFulfilled)
				if err := _MultiWordConsumer.contract.UnpackLog(event, "ChainlinkFulfilled", log); err != nil {
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

func (_MultiWordConsumer *MultiWordConsumerFilterer) ParseChainlinkFulfilled(log types.Log) (*MultiWordConsumerChainlinkFulfilled, error) {
	event := new(MultiWordConsumerChainlinkFulfilled)
	if err := _MultiWordConsumer.contract.UnpackLog(event, "ChainlinkFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MultiWordConsumerChainlinkRequestedIterator struct {
	Event *MultiWordConsumerChainlinkRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiWordConsumerChainlinkRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiWordConsumerChainlinkRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(MultiWordConsumerChainlinkRequested)
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

func (it *MultiWordConsumerChainlinkRequestedIterator) Error() error {
	return it.fail
}

func (it *MultiWordConsumerChainlinkRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiWordConsumerChainlinkRequested struct {
	Id  [32]byte
	Raw types.Log
}

func (_MultiWordConsumer *MultiWordConsumerFilterer) FilterChainlinkRequested(opts *bind.FilterOpts, id [][32]byte) (*MultiWordConsumerChainlinkRequestedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MultiWordConsumer.contract.FilterLogs(opts, "ChainlinkRequested", idRule)
	if err != nil {
		return nil, err
	}
	return &MultiWordConsumerChainlinkRequestedIterator{contract: _MultiWordConsumer.contract, event: "ChainlinkRequested", logs: logs, sub: sub}, nil
}

func (_MultiWordConsumer *MultiWordConsumerFilterer) WatchChainlinkRequested(opts *bind.WatchOpts, sink chan<- *MultiWordConsumerChainlinkRequested, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MultiWordConsumer.contract.WatchLogs(opts, "ChainlinkRequested", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiWordConsumerChainlinkRequested)
				if err := _MultiWordConsumer.contract.UnpackLog(event, "ChainlinkRequested", log); err != nil {
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

func (_MultiWordConsumer *MultiWordConsumerFilterer) ParseChainlinkRequested(log types.Log) (*MultiWordConsumerChainlinkRequested, error) {
	event := new(MultiWordConsumerChainlinkRequested)
	if err := _MultiWordConsumer.contract.UnpackLog(event, "ChainlinkRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MultiWordConsumerRequestFulfilledIterator struct {
	Event *MultiWordConsumerRequestFulfilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiWordConsumerRequestFulfilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiWordConsumerRequestFulfilled)
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

	select {
	case log := <-it.logs:
		it.Event = new(MultiWordConsumerRequestFulfilled)
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

func (it *MultiWordConsumerRequestFulfilledIterator) Error() error {
	return it.fail
}

func (it *MultiWordConsumerRequestFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiWordConsumerRequestFulfilled struct {
	RequestId [32]byte
	Price     common.Hash
	Raw       types.Log
}

func (_MultiWordConsumer *MultiWordConsumerFilterer) FilterRequestFulfilled(opts *bind.FilterOpts, requestId [][32]byte, price [][]byte) (*MultiWordConsumerRequestFulfilledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _MultiWordConsumer.contract.FilterLogs(opts, "RequestFulfilled", requestIdRule, priceRule)
	if err != nil {
		return nil, err
	}
	return &MultiWordConsumerRequestFulfilledIterator{contract: _MultiWordConsumer.contract, event: "RequestFulfilled", logs: logs, sub: sub}, nil
}

func (_MultiWordConsumer *MultiWordConsumerFilterer) WatchRequestFulfilled(opts *bind.WatchOpts, sink chan<- *MultiWordConsumerRequestFulfilled, requestId [][32]byte, price [][]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _MultiWordConsumer.contract.WatchLogs(opts, "RequestFulfilled", requestIdRule, priceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiWordConsumerRequestFulfilled)
				if err := _MultiWordConsumer.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
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

func (_MultiWordConsumer *MultiWordConsumerFilterer) ParseRequestFulfilled(log types.Log) (*MultiWordConsumerRequestFulfilled, error) {
	event := new(MultiWordConsumerRequestFulfilled)
	if err := _MultiWordConsumer.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MultiWordConsumerRequestMultipleFulfilledIterator struct {
	Event *MultiWordConsumerRequestMultipleFulfilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiWordConsumerRequestMultipleFulfilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiWordConsumerRequestMultipleFulfilled)
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

	select {
	case log := <-it.logs:
		it.Event = new(MultiWordConsumerRequestMultipleFulfilled)
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

func (it *MultiWordConsumerRequestMultipleFulfilledIterator) Error() error {
	return it.fail
}

func (it *MultiWordConsumerRequestMultipleFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiWordConsumerRequestMultipleFulfilled struct {
	RequestId [32]byte
	Usd       [32]byte
	Eur       [32]byte
	Jpy       [32]byte
	Raw       types.Log
}

func (_MultiWordConsumer *MultiWordConsumerFilterer) FilterRequestMultipleFulfilled(opts *bind.FilterOpts, requestId [][32]byte, usd [][32]byte, eur [][32]byte) (*MultiWordConsumerRequestMultipleFulfilledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var usdRule []interface{}
	for _, usdItem := range usd {
		usdRule = append(usdRule, usdItem)
	}
	var eurRule []interface{}
	for _, eurItem := range eur {
		eurRule = append(eurRule, eurItem)
	}

	logs, sub, err := _MultiWordConsumer.contract.FilterLogs(opts, "RequestMultipleFulfilled", requestIdRule, usdRule, eurRule)
	if err != nil {
		return nil, err
	}
	return &MultiWordConsumerRequestMultipleFulfilledIterator{contract: _MultiWordConsumer.contract, event: "RequestMultipleFulfilled", logs: logs, sub: sub}, nil
}

func (_MultiWordConsumer *MultiWordConsumerFilterer) WatchRequestMultipleFulfilled(opts *bind.WatchOpts, sink chan<- *MultiWordConsumerRequestMultipleFulfilled, requestId [][32]byte, usd [][32]byte, eur [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var usdRule []interface{}
	for _, usdItem := range usd {
		usdRule = append(usdRule, usdItem)
	}
	var eurRule []interface{}
	for _, eurItem := range eur {
		eurRule = append(eurRule, eurItem)
	}

	logs, sub, err := _MultiWordConsumer.contract.WatchLogs(opts, "RequestMultipleFulfilled", requestIdRule, usdRule, eurRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiWordConsumerRequestMultipleFulfilled)
				if err := _MultiWordConsumer.contract.UnpackLog(event, "RequestMultipleFulfilled", log); err != nil {
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

func (_MultiWordConsumer *MultiWordConsumerFilterer) ParseRequestMultipleFulfilled(log types.Log) (*MultiWordConsumerRequestMultipleFulfilled, error) {
	event := new(MultiWordConsumerRequestMultipleFulfilled)
	if err := _MultiWordConsumer.contract.UnpackLog(event, "RequestMultipleFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_MultiWordConsumer *MultiWordConsumer) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _MultiWordConsumer.abi.Events["ChainlinkCancelled"].ID:
		return _MultiWordConsumer.ParseChainlinkCancelled(log)
	case _MultiWordConsumer.abi.Events["ChainlinkFulfilled"].ID:
		return _MultiWordConsumer.ParseChainlinkFulfilled(log)
	case _MultiWordConsumer.abi.Events["ChainlinkRequested"].ID:
		return _MultiWordConsumer.ParseChainlinkRequested(log)
	case _MultiWordConsumer.abi.Events["RequestFulfilled"].ID:
		return _MultiWordConsumer.ParseRequestFulfilled(log)
	case _MultiWordConsumer.abi.Events["RequestMultipleFulfilled"].ID:
		return _MultiWordConsumer.ParseRequestMultipleFulfilled(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (MultiWordConsumerChainlinkCancelled) Topic() common.Hash {
	return common.HexToHash("0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5")
}

func (MultiWordConsumerChainlinkFulfilled) Topic() common.Hash {
	return common.HexToHash("0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a")
}

func (MultiWordConsumerChainlinkRequested) Topic() common.Hash {
	return common.HexToHash("0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9")
}

func (MultiWordConsumerRequestFulfilled) Topic() common.Hash {
	return common.HexToHash("0x1a111c5dcf9a71088bd5e1797fdfaf399fec2afbb24aca247e4e3e9f4b61df91")
}

func (MultiWordConsumerRequestMultipleFulfilled) Topic() common.Hash {
	return common.HexToHash("0x0ec0c13e44aa04198947078cb990660252870dd3363f4c4bb3cc780f808dabbe")
}

func (_MultiWordConsumer *MultiWordConsumer) Address() common.Address {
	return _MultiWordConsumer.address
}

type MultiWordConsumerInterface interface {
	CurrentPrice(opts *bind.CallOpts) ([]byte, error)

	Eur(opts *bind.CallOpts) ([32]byte, error)

	Jpy(opts *bind.CallOpts) ([32]byte, error)

	Usd(opts *bind.CallOpts) ([32]byte, error)

	AddExternalRequest(opts *bind.TransactOpts, _oracle common.Address, _requestId [32]byte) (*types.Transaction, error)

	CancelRequest(opts *bind.TransactOpts, _oracle common.Address, _requestId [32]byte, _payment *big.Int, _callbackFunctionId [4]byte, _expiration *big.Int) (*types.Transaction, error)

	FulfillBytes(opts *bind.TransactOpts, _requestId [32]byte, _price []byte) (*types.Transaction, error)

	FulfillMultipleParameters(opts *bind.TransactOpts, _requestId [32]byte, _usd [32]byte, _eur [32]byte, _jpy [32]byte) (*types.Transaction, error)

	RequestEthereumPrice(opts *bind.TransactOpts, _currency string, _payment *big.Int) (*types.Transaction, error)

	RequestEthereumPriceByCallback(opts *bind.TransactOpts, _currency string, _payment *big.Int, _callback common.Address) (*types.Transaction, error)

	RequestMultipleParameters(opts *bind.TransactOpts, _currency string, _payment *big.Int) (*types.Transaction, error)

	RequestMultipleParametersWithCustomURLs(opts *bind.TransactOpts, _urlUSD string, _pathUSD string, _urlEUR string, _pathEUR string, _urlJPY string, _pathJPY string, _payment *big.Int) (*types.Transaction, error)

	SetSpecID(opts *bind.TransactOpts, _specId [32]byte) (*types.Transaction, error)

	WithdrawLink(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterChainlinkCancelled(opts *bind.FilterOpts, id [][32]byte) (*MultiWordConsumerChainlinkCancelledIterator, error)

	WatchChainlinkCancelled(opts *bind.WatchOpts, sink chan<- *MultiWordConsumerChainlinkCancelled, id [][32]byte) (event.Subscription, error)

	ParseChainlinkCancelled(log types.Log) (*MultiWordConsumerChainlinkCancelled, error)

	FilterChainlinkFulfilled(opts *bind.FilterOpts, id [][32]byte) (*MultiWordConsumerChainlinkFulfilledIterator, error)

	WatchChainlinkFulfilled(opts *bind.WatchOpts, sink chan<- *MultiWordConsumerChainlinkFulfilled, id [][32]byte) (event.Subscription, error)

	ParseChainlinkFulfilled(log types.Log) (*MultiWordConsumerChainlinkFulfilled, error)

	FilterChainlinkRequested(opts *bind.FilterOpts, id [][32]byte) (*MultiWordConsumerChainlinkRequestedIterator, error)

	WatchChainlinkRequested(opts *bind.WatchOpts, sink chan<- *MultiWordConsumerChainlinkRequested, id [][32]byte) (event.Subscription, error)

	ParseChainlinkRequested(log types.Log) (*MultiWordConsumerChainlinkRequested, error)

	FilterRequestFulfilled(opts *bind.FilterOpts, requestId [][32]byte, price [][]byte) (*MultiWordConsumerRequestFulfilledIterator, error)

	WatchRequestFulfilled(opts *bind.WatchOpts, sink chan<- *MultiWordConsumerRequestFulfilled, requestId [][32]byte, price [][]byte) (event.Subscription, error)

	ParseRequestFulfilled(log types.Log) (*MultiWordConsumerRequestFulfilled, error)

	FilterRequestMultipleFulfilled(opts *bind.FilterOpts, requestId [][32]byte, usd [][32]byte, eur [][32]byte) (*MultiWordConsumerRequestMultipleFulfilledIterator, error)

	WatchRequestMultipleFulfilled(opts *bind.WatchOpts, sink chan<- *MultiWordConsumerRequestMultipleFulfilled, requestId [][32]byte, usd [][32]byte, eur [][32]byte) (event.Subscription, error)

	ParseRequestMultipleFulfilled(log types.Log) (*MultiWordConsumerRequestMultipleFulfilled, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}

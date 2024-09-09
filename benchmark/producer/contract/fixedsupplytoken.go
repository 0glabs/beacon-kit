// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// Erc20testMetaData contains all meta data concerning the Erc20test contract.
var Erc20testMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferAnyERC20Token\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenOwner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600080546001600160a01b031916331790556040805180820190915260058082527f464958454400000000000000000000000000000000000000000000000000000060209092019182526100679160029161012a565b5060408051808201909152601a8082527f4578616d706c6520466978656420537570706c7920546f6b656e00000000000060209092019182526100ac9160039161012a565b5060048054601260ff19909116179081905560ff16600a0a620f4240026005819055600080546001600160a01b0390811682526006602090815260408084208590558354815195865290519216937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929081900390910190a36101c5565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061016b57805160ff1916838001178555610198565b82800160010185558215610198579182015b8281111561019857825182559160200191906001019061017d565b506101a49291506101a8565b5090565b6101c291905b808211156101a457600081556001016101ae565b90565b610b53806101d46000396000f3fe6080604052600436106100e85760003560e01c80638da5cb5b1161008a578063d4ee1d9011610059578063d4ee1d90146103ea578063dc39d06d146103ff578063dd62ed3e14610438578063f2fde38b14610473576100e8565b80638da5cb5b146102a357806395d89b41146102d4578063a9059cbb146102e9578063cae9ca5114610322576100e8565b806323b872dd116100c657806323b872dd146101eb578063313ce5671461022e57806370a082311461025957806379ba50971461028c576100e8565b806306fdde03146100ed578063095ea7b31461017757806318160ddd146101c4575b600080fd5b3480156100f957600080fd5b506101026104a6565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561013c578181015183820152602001610124565b50505050905090810190601f1680156101695780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561018357600080fd5b506101b06004803603604081101561019a57600080fd5b506001600160a01b038135169060200135610534565b604080519115158252519081900360200190f35b3480156101d057600080fd5b506101d961059b565b60408051918252519081900360200190f35b3480156101f757600080fd5b506101b06004803603606081101561020e57600080fd5b506001600160a01b038135811691602081013590911690604001356105de565b34801561023a57600080fd5b506102436106e9565b6040805160ff9092168252519081900360200190f35b34801561026557600080fd5b506101d96004803603602081101561027c57600080fd5b50356001600160a01b03166106f2565b34801561029857600080fd5b506102a161070d565b005b3480156102af57600080fd5b506102b8610788565b604080516001600160a01b039092168252519081900360200190f35b3480156102e057600080fd5b50610102610797565b3480156102f557600080fd5b506101b06004803603604081101561030c57600080fd5b506001600160a01b0381351690602001356107ef565b34801561032e57600080fd5b506101b06004803603606081101561034557600080fd5b6001600160a01b038235169160208101359181019060608101604082013564010000000081111561037557600080fd5b82018360208201111561038757600080fd5b803590602001918460018302840111640100000000831117156103a957600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061089f945050505050565b3480156103f657600080fd5b506102b86109ea565b34801561040b57600080fd5b506101b06004803603604081101561042257600080fd5b506001600160a01b0381351690602001356109f9565b34801561044457600080fd5b506101d96004803603604081101561045b57600080fd5b506001600160a01b0381358116916020013516610a9e565b34801561047f57600080fd5b506102a16004803603602081101561049657600080fd5b50356001600160a01b0316610ac9565b6003805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561052c5780601f106105015761010080835404028352916020019161052c565b820191906000526020600020905b81548152906001019060200180831161050f57829003601f168201915b505050505081565b3360008181526007602090815260408083206001600160a01b038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a35060015b92915050565b600080805260066020527f54cdd369e4e8a8515e52ca72ec816c2101831ad1f18bf44102ed171459c9b4f8546005546105d99163ffffffff610b0216565b905090565b6001600160a01b038316600090815260066020526040812054610607908363ffffffff610b0216565b6001600160a01b0385166000908152600660209081526040808320939093556007815282822033835290522054610644908363ffffffff610b0216565b6001600160a01b038086166000908152600760209081526040808320338452825280832094909455918616815260069091522054610688908363ffffffff610b1716565b6001600160a01b0380851660008181526006602090815260409182902094909455805186815290519193928816927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a35060019392505050565b60045460ff1681565b6001600160a01b031660009081526006602052604090205490565b6001546001600160a01b0316331461072457600080fd5b600154600080546040516001600160a01b0393841693909116917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a360018054600080546001600160a01b03199081166001600160a01b03841617909155169055565b6000546001600160a01b031681565b6002805460408051602060018416156101000260001901909316849004601f8101849004840282018401909252818152929183018282801561052c5780601f106105015761010080835404028352916020019161052c565b3360009081526006602052604081205461080f908363ffffffff610b0216565b33600090815260066020526040808220929092556001600160a01b03851681522054610841908363ffffffff610b1716565b6001600160a01b0384166000818152600660209081526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b3360008181526007602090815260408083206001600160a01b038816808552908352818420879055815187815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a3604051600160e01b638f4ffcb102815233600482018181526024830186905230604484018190526080606485019081528651608486015286516001600160a01b038a1695638f4ffcb195948a94938a939192909160a490910190602085019080838360005b83811015610979578181015183820152602001610961565b50505050905090810190601f1680156109a65780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b1580156109c857600080fd5b505af11580156109dc573d6000803e3d6000fd5b506001979650505050505050565b6001546001600160a01b031681565b600080546001600160a01b03163314610a1157600080fd5b6000805460408051600160e01b63a9059cbb0281526001600160a01b0392831660048201526024810186905290519186169263a9059cbb926044808401936020939083900390910190829087803b158015610a6b57600080fd5b505af1158015610a7f573d6000803e3d6000fd5b505050506040513d6020811015610a9557600080fd5b50519392505050565b6001600160a01b03918216600090815260076020908152604080832093909416825291909152205490565b6000546001600160a01b03163314610ae057600080fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b600082821115610b1157600080fd5b50900390565b8181018281101561059557600080fdfea165627a7a72305820a6aa9231cc520fbbc3b891faeeddc44729a7f8a4449fd4b5ef98392af7b07cc90029",
}

// Erc20testABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20testMetaData.ABI instead.
var Erc20testABI = Erc20testMetaData.ABI

// Erc20testBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Erc20testMetaData.Bin instead.
var Erc20testBin = Erc20testMetaData.Bin

// DeployErc20test deploys a new Ethereum contract, binding an instance of Erc20test to it.
func DeployErc20test(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Erc20test, error) {
	parsed, err := Erc20testMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Erc20testBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc20test{Erc20testCaller: Erc20testCaller{contract: contract}, Erc20testTransactor: Erc20testTransactor{contract: contract}, Erc20testFilterer: Erc20testFilterer{contract: contract}}, nil
}

// Erc20test is an auto generated Go binding around an Ethereum contract.
type Erc20test struct {
	Erc20testCaller     // Read-only binding to the contract
	Erc20testTransactor // Write-only binding to the contract
	Erc20testFilterer   // Log filterer for contract events
}

// Erc20testCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20testCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20testTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20testTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20testFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20testFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20testSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20testSession struct {
	Contract     *Erc20test        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20testCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20testCallerSession struct {
	Contract *Erc20testCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Erc20testTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20testTransactorSession struct {
	Contract     *Erc20testTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Erc20testRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20testRaw struct {
	Contract *Erc20test // Generic contract binding to access the raw methods on
}

// Erc20testCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20testCallerRaw struct {
	Contract *Erc20testCaller // Generic read-only contract binding to access the raw methods on
}

// Erc20testTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20testTransactorRaw struct {
	Contract *Erc20testTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20test creates a new instance of Erc20test, bound to a specific deployed contract.
func NewErc20test(address common.Address, backend bind.ContractBackend) (*Erc20test, error) {
	contract, err := bindErc20test(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20test{Erc20testCaller: Erc20testCaller{contract: contract}, Erc20testTransactor: Erc20testTransactor{contract: contract}, Erc20testFilterer: Erc20testFilterer{contract: contract}}, nil
}

// NewErc20testCaller creates a new read-only instance of Erc20test, bound to a specific deployed contract.
func NewErc20testCaller(address common.Address, caller bind.ContractCaller) (*Erc20testCaller, error) {
	contract, err := bindErc20test(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20testCaller{contract: contract}, nil
}

// NewErc20testTransactor creates a new write-only instance of Erc20test, bound to a specific deployed contract.
func NewErc20testTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc20testTransactor, error) {
	contract, err := bindErc20test(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20testTransactor{contract: contract}, nil
}

// NewErc20testFilterer creates a new log filterer instance of Erc20test, bound to a specific deployed contract.
func NewErc20testFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc20testFilterer, error) {
	contract, err := bindErc20test(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20testFilterer{contract: contract}, nil
}

// bindErc20test binds a generic wrapper to an already deployed contract.
func bindErc20test(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc20testMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20test *Erc20testRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20test.Contract.Erc20testCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20test *Erc20testRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20test.Contract.Erc20testTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20test *Erc20testRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20test.Contract.Erc20testTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20test *Erc20testCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20test.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20test *Erc20testTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20test.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20test *Erc20testTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20test.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) view returns(uint256 remaining)
func (_Erc20test *Erc20testCaller) Allowance(opts *bind.CallOpts, tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20test.contract.Call(opts, &out, "allowance", tokenOwner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) view returns(uint256 remaining)
func (_Erc20test *Erc20testSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc20test.Contract.Allowance(&_Erc20test.CallOpts, tokenOwner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) view returns(uint256 remaining)
func (_Erc20test *Erc20testCallerSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc20test.Contract.Allowance(&_Erc20test.CallOpts, tokenOwner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256 balance)
func (_Erc20test *Erc20testCaller) BalanceOf(opts *bind.CallOpts, tokenOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20test.contract.Call(opts, &out, "balanceOf", tokenOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256 balance)
func (_Erc20test *Erc20testSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _Erc20test.Contract.BalanceOf(&_Erc20test.CallOpts, tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256 balance)
func (_Erc20test *Erc20testCallerSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _Erc20test.Contract.BalanceOf(&_Erc20test.CallOpts, tokenOwner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20test *Erc20testCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Erc20test.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20test *Erc20testSession) Decimals() (uint8, error) {
	return _Erc20test.Contract.Decimals(&_Erc20test.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20test *Erc20testCallerSession) Decimals() (uint8, error) {
	return _Erc20test.Contract.Decimals(&_Erc20test.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20test *Erc20testCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc20test.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20test *Erc20testSession) Name() (string, error) {
	return _Erc20test.Contract.Name(&_Erc20test.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20test *Erc20testCallerSession) Name() (string, error) {
	return _Erc20test.Contract.Name(&_Erc20test.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_Erc20test *Erc20testCaller) NewOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20test.contract.Call(opts, &out, "newOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_Erc20test *Erc20testSession) NewOwner() (common.Address, error) {
	return _Erc20test.Contract.NewOwner(&_Erc20test.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_Erc20test *Erc20testCallerSession) NewOwner() (common.Address, error) {
	return _Erc20test.Contract.NewOwner(&_Erc20test.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc20test *Erc20testCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20test.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc20test *Erc20testSession) Owner() (common.Address, error) {
	return _Erc20test.Contract.Owner(&_Erc20test.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc20test *Erc20testCallerSession) Owner() (common.Address, error) {
	return _Erc20test.Contract.Owner(&_Erc20test.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20test *Erc20testCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc20test.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20test *Erc20testSession) Symbol() (string, error) {
	return _Erc20test.Contract.Symbol(&_Erc20test.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20test *Erc20testCallerSession) Symbol() (string, error) {
	return _Erc20test.Contract.Symbol(&_Erc20test.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20test *Erc20testCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20test.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20test *Erc20testSession) TotalSupply() (*big.Int, error) {
	return _Erc20test.Contract.TotalSupply(&_Erc20test.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20test *Erc20testCallerSession) TotalSupply() (*big.Int, error) {
	return _Erc20test.Contract.TotalSupply(&_Erc20test.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Erc20test *Erc20testTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20test.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Erc20test *Erc20testSession) AcceptOwnership() (*types.Transaction, error) {
	return _Erc20test.Contract.AcceptOwnership(&_Erc20test.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Erc20test *Erc20testTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Erc20test.Contract.AcceptOwnership(&_Erc20test.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 tokens) returns(bool success)
func (_Erc20test *Erc20testTransactor) Approve(opts *bind.TransactOpts, spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Erc20test.contract.Transact(opts, "approve", spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 tokens) returns(bool success)
func (_Erc20test *Erc20testSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Erc20test.Contract.Approve(&_Erc20test.TransactOpts, spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 tokens) returns(bool success)
func (_Erc20test *Erc20testTransactorSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Erc20test.Contract.Approve(&_Erc20test.TransactOpts, spender, tokens)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 tokens, bytes data) returns(bool success)
func (_Erc20test *Erc20testTransactor) ApproveAndCall(opts *bind.TransactOpts, spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc20test.contract.Transact(opts, "approveAndCall", spender, tokens, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 tokens, bytes data) returns(bool success)
func (_Erc20test *Erc20testSession) ApproveAndCall(spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc20test.Contract.ApproveAndCall(&_Erc20test.TransactOpts, spender, tokens, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 tokens, bytes data) returns(bool success)
func (_Erc20test *Erc20testTransactorSession) ApproveAndCall(spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc20test.Contract.ApproveAndCall(&_Erc20test.TransactOpts, spender, tokens, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 tokens) returns(bool success)
func (_Erc20test *Erc20testTransactor) Transfer(opts *bind.TransactOpts, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Erc20test.contract.Transact(opts, "transfer", to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 tokens) returns(bool success)
func (_Erc20test *Erc20testSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Erc20test.Contract.Transfer(&_Erc20test.TransactOpts, to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 tokens) returns(bool success)
func (_Erc20test *Erc20testTransactorSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Erc20test.Contract.Transfer(&_Erc20test.TransactOpts, to, tokens)
}

// TransferAnyERC20Token is a paid mutator transaction binding the contract method 0xdc39d06d.
//
// Solidity: function transferAnyERC20Token(address tokenAddress, uint256 tokens) returns(bool success)
func (_Erc20test *Erc20testTransactor) TransferAnyERC20Token(opts *bind.TransactOpts, tokenAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Erc20test.contract.Transact(opts, "transferAnyERC20Token", tokenAddress, tokens)
}

// TransferAnyERC20Token is a paid mutator transaction binding the contract method 0xdc39d06d.
//
// Solidity: function transferAnyERC20Token(address tokenAddress, uint256 tokens) returns(bool success)
func (_Erc20test *Erc20testSession) TransferAnyERC20Token(tokenAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Erc20test.Contract.TransferAnyERC20Token(&_Erc20test.TransactOpts, tokenAddress, tokens)
}

// TransferAnyERC20Token is a paid mutator transaction binding the contract method 0xdc39d06d.
//
// Solidity: function transferAnyERC20Token(address tokenAddress, uint256 tokens) returns(bool success)
func (_Erc20test *Erc20testTransactorSession) TransferAnyERC20Token(tokenAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Erc20test.Contract.TransferAnyERC20Token(&_Erc20test.TransactOpts, tokenAddress, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_Erc20test *Erc20testTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Erc20test.contract.Transact(opts, "transferFrom", from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_Erc20test *Erc20testSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Erc20test.Contract.TransferFrom(&_Erc20test.TransactOpts, from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_Erc20test *Erc20testTransactorSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Erc20test.Contract.TransferFrom(&_Erc20test.TransactOpts, from, to, tokens)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Erc20test *Erc20testTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Erc20test.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Erc20test *Erc20testSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Erc20test.Contract.TransferOwnership(&_Erc20test.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Erc20test *Erc20testTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Erc20test.Contract.TransferOwnership(&_Erc20test.TransactOpts, _newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Erc20test *Erc20testTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Erc20test.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Erc20test *Erc20testSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Erc20test.Contract.Fallback(&_Erc20test.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Erc20test *Erc20testTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Erc20test.Contract.Fallback(&_Erc20test.TransactOpts, calldata)
}

// Erc20testApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc20test contract.
type Erc20testApprovalIterator struct {
	Event *Erc20testApproval // Event containing the contract specifics and raw log

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
func (it *Erc20testApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20testApproval)
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
		it.Event = new(Erc20testApproval)
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
func (it *Erc20testApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20testApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20testApproval represents a Approval event raised by the Erc20test contract.
type Erc20testApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Erc20test *Erc20testFilterer) FilterApproval(opts *bind.FilterOpts, tokenOwner []common.Address, spender []common.Address) (*Erc20testApprovalIterator, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20test.contract.FilterLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Erc20testApprovalIterator{contract: _Erc20test.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Erc20test *Erc20testFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc20testApproval, tokenOwner []common.Address, spender []common.Address) (event.Subscription, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20test.contract.WatchLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20testApproval)
				if err := _Erc20test.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Erc20test *Erc20testFilterer) ParseApproval(log types.Log) (*Erc20testApproval, error) {
	event := new(Erc20testApproval)
	if err := _Erc20test.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20testOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Erc20test contract.
type Erc20testOwnershipTransferredIterator struct {
	Event *Erc20testOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Erc20testOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20testOwnershipTransferred)
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
		it.Event = new(Erc20testOwnershipTransferred)
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
func (it *Erc20testOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20testOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20testOwnershipTransferred represents a OwnershipTransferred event raised by the Erc20test contract.
type Erc20testOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _from, address indexed _to)
func (_Erc20test *Erc20testFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*Erc20testOwnershipTransferredIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Erc20test.contract.FilterLogs(opts, "OwnershipTransferred", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &Erc20testOwnershipTransferredIterator{contract: _Erc20test.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _from, address indexed _to)
func (_Erc20test *Erc20testFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Erc20testOwnershipTransferred, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Erc20test.contract.WatchLogs(opts, "OwnershipTransferred", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20testOwnershipTransferred)
				if err := _Erc20test.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _from, address indexed _to)
func (_Erc20test *Erc20testFilterer) ParseOwnershipTransferred(log types.Log) (*Erc20testOwnershipTransferred, error) {
	event := new(Erc20testOwnershipTransferred)
	if err := _Erc20test.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20testTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc20test contract.
type Erc20testTransferIterator struct {
	Event *Erc20testTransfer // Event containing the contract specifics and raw log

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
func (it *Erc20testTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20testTransfer)
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
		it.Event = new(Erc20testTransfer)
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
func (it *Erc20testTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20testTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20testTransfer represents a Transfer event raised by the Erc20test contract.
type Erc20testTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Erc20test *Erc20testFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Erc20testTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20test.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc20testTransferIterator{contract: _Erc20test.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Erc20test *Erc20testFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc20testTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20test.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20testTransfer)
				if err := _Erc20test.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Erc20test *Erc20testFilterer) ParseTransfer(log types.Log) (*Erc20testTransfer, error) {
	event := new(Erc20testTransfer)
	if err := _Erc20test.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

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
)

// StoreUser is an auto generated low-level Go binding around an user-defined struct.
type StoreUser struct {
	Username  string
	Direction common.Address
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"direction\",\"type\":\"address\"}],\"name\":\"Unsubscribe\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"list\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"direction\",\"type\":\"address\"}],\"internalType\":\"structStore.User[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundMoneyUser\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_username\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unsubscribe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506002600160008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506113b0806100c66000396000f3fe60806040526004361061003f5760003560e01c80630f560cd714610044578063cd0731ac1461006f578063f2c298be14610079578063fcae448414610095575b600080fd5b34801561005057600080fd5b506100596100ac565b6040516100669190610a72565b60405180910390f35b6100776101f3565b005b610093600480360381019061008e9190610bdd565b610367565b005b3480156100a157600080fd5b506100aa61053f565b005b60606002805480602002602001604051908101604052809291908181526020016000905b828210156101ea578382906000526020600020906002020160405180604001604052908160008201805461010390610c55565b80601f016020809104026020016040519081016040528092919081815260200182805461012f90610c55565b801561017c5780601f106101515761010080835404028352916020019161017c565b820191906000526020600020905b81548152906001019060200180831161015f57829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525050815260200190600101906100d0565b50505050905090565b6003600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414610275576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026c90610ce3565b60405180910390fd5b60003373ffffffffffffffffffffffffffffffffffffffff163460405161029b90610d34565b60006040518083038185875af1925050503d80600081146102d8576040519150601f19603f3d011682016040523d82523d6000602084013e6102dd565b606091505b5050905080610321576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031890610d95565b60405180910390fd5b6002600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020505050565b600260405180604001604052808381526020013373ffffffffffffffffffffffffffffffffffffffff16815250908060018154018082558091505060019003906000526020600020906002020160009091909190915060008201518160000190816103d29190610f6b565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050506000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541461049e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161049590611089565b60405180910390fd5b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008154809291906104ee906110d8565b91905055506509184e72a00034101561053c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053390610d95565b60405180910390fd5b50565b600060018060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054146105c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b99061116c565b60405180910390fd5b60005b600280549050811015610661573373ffffffffffffffffffffffffffffffffffffffff16600282815481106105fd576105fc61118c565b5b906000526020600020906002020160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361064e578091505b8080610659906110d8565b9150506105c5565b5060008190505b600160028054905061067a91906111bb565b81101561076357600260018261069091906111ef565b815481106106a1576106a061118c565b5b9060005260206000209060020201600282815481106106c3576106c261118c565b5b9060005260206000209060020201600082018160000190816106e59190611239565b506001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550905050808061075b906110d8565b915050610668565b50600280548061077657610775611321565b5b60019003818190600052602060002090600202016000808201600061079b9190610845565b6001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055505090556003600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507feed050308c603899d7397c26bdccda0810c3ccc6e9730a8a10c452b522f8edf43360405161083a919061135f565b60405180910390a150565b50805461085190610c55565b6000825580601f106108635750610882565b601f0160209004906000526020600020908101906108819190610885565b5b50565b5b8082111561089e576000816000905550600101610886565b5090565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b838110156109085780820151818401526020810190506108ed565b60008484015250505050565b6000601f19601f8301169050919050565b6000610930826108ce565b61093a81856108d9565b935061094a8185602086016108ea565b61095381610914565b840191505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006109898261095e565b9050919050565b6109998161097e565b82525050565b600060408301600083015184820360008601526109bc8282610925565b91505060208301516109d16020860182610990565b508091505092915050565b60006109e8838361099f565b905092915050565b6000602082019050919050565b6000610a08826108a2565b610a1281856108ad565b935083602082028501610a24856108be565b8060005b85811015610a605784840389528151610a4185826109dc565b9450610a4c836109f0565b925060208a01995050600181019050610a28565b50829750879550505050505092915050565b60006020820190508181036000830152610a8c81846109fd565b905092915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610aea82610914565b810181811067ffffffffffffffff82111715610b0957610b08610ab2565b5b80604052505050565b6000610b1c610a94565b9050610b288282610ae1565b919050565b600067ffffffffffffffff821115610b4857610b47610ab2565b5b610b5182610914565b9050602081019050919050565b82818337600083830152505050565b6000610b80610b7b84610b2d565b610b12565b905082815260208101848484011115610b9c57610b9b610aad565b5b610ba7848285610b5e565b509392505050565b600082601f830112610bc457610bc3610aa8565b5b8135610bd4848260208601610b6d565b91505092915050565b600060208284031215610bf357610bf2610a9e565b5b600082013567ffffffffffffffff811115610c1157610c10610aa3565b5b610c1d84828501610baf565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610c6d57607f821691505b602082108103610c8057610c7f610c26565b5b50919050565b600082825260208201905092915050565b7f4661696c656420746f20726566756e6400000000000000000000000000000000600082015250565b6000610ccd601083610c86565b9150610cd882610c97565b602082019050919050565b60006020820190508181036000830152610cfc81610cc0565b9050919050565b600081905092915050565b50565b6000610d1e600083610d03565b9150610d2982610d0e565b600082019050919050565b6000610d3f82610d11565b9150819050919050565b7f4661696c656420746f2073656e64204574686572000000000000000000000000600082015250565b6000610d7f601483610c86565b9150610d8a82610d49565b602082019050919050565b60006020820190508181036000830152610dae81610d72565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610e177fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610dda565b610e218683610dda565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000610e68610e63610e5e84610e39565b610e43565b610e39565b9050919050565b6000819050919050565b610e8283610e4d565b610e96610e8e82610e6f565b848454610de7565b825550505050565b600090565b610eab610e9e565b610eb6818484610e79565b505050565b5b81811015610eda57610ecf600082610ea3565b600181019050610ebc565b5050565b601f821115610f1f57610ef081610db5565b610ef984610dca565b81016020851015610f08578190505b610f1c610f1485610dca565b830182610ebb565b50505b505050565b600082821c905092915050565b6000610f4260001984600802610f24565b1980831691505092915050565b6000610f5b8383610f31565b9150826002028217905092915050565b610f74826108ce565b67ffffffffffffffff811115610f8d57610f8c610ab2565b5b610f978254610c55565b610fa2828285610ede565b600060209050601f831160018114610fd55760008415610fc3578287015190505b610fcd8582610f4f565b865550611035565b601f198416610fe386610db5565b60005b8281101561100b57848901518255600182019150602085019450602081019050610fe6565b868310156110285784890151611024601f891682610f31565b8355505b6001600288020188555050505b505050505050565b7f5265676973746572206661696c65640000000000000000000000000000000000600082015250565b6000611073600f83610c86565b915061107e8261103d565b602082019050919050565b600060208201905081810360008301526110a281611066565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006110e382610e39565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611115576111146110a9565b5b600182019050919050565b7f556e737562736372696265206661696c65640000000000000000000000000000600082015250565b6000611156601283610c86565b915061116182611120565b602082019050919050565b6000602082019050818103600083015261118581611149565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006111c682610e39565b91506111d183610e39565b92508282039050818111156111e9576111e86110a9565b5b92915050565b60006111fa82610e39565b915061120583610e39565b925082820190508082111561121d5761121c6110a9565b5b92915050565b60008154905061123281610c55565b9050919050565b81810361124757505061131f565b61125082611223565b67ffffffffffffffff81111561126957611268610ab2565b5b6112738254610c55565b61127e828285610ede565b6000601f8311600181146112ad576000841561129b578287015490505b6112a58582610f4f565b865550611318565b601f1984166112bb87610db5565b96506112c686610db5565b60005b828110156112ee578489015482556001820191506001850194506020810190506112c9565b8683101561130b5784890154611307601f891682610f31565b8355505b6001600288020188555050505b5050505050505b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6113598161097e565b82525050565b60006020820190506113746000830184611350565b9291505056fea2646970667358221220c2b490f0ef056bf8397c70906f0663fafd64626833327b9e32eefa9fc01ca27864736f6c63430008110033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,address)[])
func (_Contract *ContractCaller) List(opts *bind.CallOpts) ([]StoreUser, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "list")

	if err != nil {
		return *new([]StoreUser), err
	}

	out0 := *abi.ConvertType(out[0], new([]StoreUser)).(*[]StoreUser)

	return out0, err

}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,address)[])
func (_Contract *ContractSession) List() ([]StoreUser, error) {
	return _Contract.Contract.List(&_Contract.CallOpts)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,address)[])
func (_Contract *ContractCallerSession) List() ([]StoreUser, error) {
	return _Contract.Contract.List(&_Contract.CallOpts)
}

// RefundMoneyUser is a paid mutator transaction binding the contract method 0xcd0731ac.
//
// Solidity: function refundMoneyUser() payable returns()
func (_Contract *ContractTransactor) RefundMoneyUser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "refundMoneyUser")
}

// RefundMoneyUser is a paid mutator transaction binding the contract method 0xcd0731ac.
//
// Solidity: function refundMoneyUser() payable returns()
func (_Contract *ContractSession) RefundMoneyUser() (*types.Transaction, error) {
	return _Contract.Contract.RefundMoneyUser(&_Contract.TransactOpts)
}

// RefundMoneyUser is a paid mutator transaction binding the contract method 0xcd0731ac.
//
// Solidity: function refundMoneyUser() payable returns()
func (_Contract *ContractTransactorSession) RefundMoneyUser() (*types.Transaction, error) {
	return _Contract.Contract.RefundMoneyUser(&_Contract.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string _username) payable returns()
func (_Contract *ContractTransactor) Register(opts *bind.TransactOpts, _username string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "register", _username)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string _username) payable returns()
func (_Contract *ContractSession) Register(_username string) (*types.Transaction, error) {
	return _Contract.Contract.Register(&_Contract.TransactOpts, _username)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string _username) payable returns()
func (_Contract *ContractTransactorSession) Register(_username string) (*types.Transaction, error) {
	return _Contract.Contract.Register(&_Contract.TransactOpts, _username)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xfcae4484.
//
// Solidity: function unsubscribe() returns()
func (_Contract *ContractTransactor) Unsubscribe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unsubscribe")
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xfcae4484.
//
// Solidity: function unsubscribe() returns()
func (_Contract *ContractSession) Unsubscribe() (*types.Transaction, error) {
	return _Contract.Contract.Unsubscribe(&_Contract.TransactOpts)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xfcae4484.
//
// Solidity: function unsubscribe() returns()
func (_Contract *ContractTransactorSession) Unsubscribe() (*types.Transaction, error) {
	return _Contract.Contract.Unsubscribe(&_Contract.TransactOpts)
}

// ContractUnsubscribeIterator is returned from FilterUnsubscribe and is used to iterate over the raw logs and unpacked data for Unsubscribe events raised by the Contract contract.
type ContractUnsubscribeIterator struct {
	Event *ContractUnsubscribe // Event containing the contract specifics and raw log

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
func (it *ContractUnsubscribeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUnsubscribe)
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
		it.Event = new(ContractUnsubscribe)
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
func (it *ContractUnsubscribeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUnsubscribeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUnsubscribe represents a Unsubscribe event raised by the Contract contract.
type ContractUnsubscribe struct {
	Direction common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnsubscribe is a free log retrieval operation binding the contract event 0xeed050308c603899d7397c26bdccda0810c3ccc6e9730a8a10c452b522f8edf4.
//
// Solidity: event Unsubscribe(address direction)
func (_Contract *ContractFilterer) FilterUnsubscribe(opts *bind.FilterOpts) (*ContractUnsubscribeIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Unsubscribe")
	if err != nil {
		return nil, err
	}
	return &ContractUnsubscribeIterator{contract: _Contract.contract, event: "Unsubscribe", logs: logs, sub: sub}, nil
}

// WatchUnsubscribe is a free log subscription operation binding the contract event 0xeed050308c603899d7397c26bdccda0810c3ccc6e9730a8a10c452b522f8edf4.
//
// Solidity: event Unsubscribe(address direction)
func (_Contract *ContractFilterer) WatchUnsubscribe(opts *bind.WatchOpts, sink chan<- *ContractUnsubscribe) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Unsubscribe")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUnsubscribe)
				if err := _Contract.contract.UnpackLog(event, "Unsubscribe", log); err != nil {
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

// ParseUnsubscribe is a log parse operation binding the contract event 0xeed050308c603899d7397c26bdccda0810c3ccc6e9730a8a10c452b522f8edf4.
//
// Solidity: event Unsubscribe(address direction)
func (_Contract *ContractFilterer) ParseUnsubscribe(log types.Log) (*ContractUnsubscribe, error) {
	event := new(ContractUnsubscribe)
	if err := _Contract.contract.UnpackLog(event, "Unsubscribe", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

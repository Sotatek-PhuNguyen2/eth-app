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
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506002600160008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506113b7806100c66000396000f3fe60806040526004361061003f5760003560e01c80630f560cd714610044578063cd0731ac1461006f578063f2c298be14610079578063fcae448414610095575b600080fd5b34801561005057600080fd5b506100596100ac565b6040516100669190610a79565b60405180910390f35b6100776101f3565b005b610093600480360381019061008e9190610be4565b61036e565b005b3480156100a157600080fd5b506100aa610546565b005b60606002805480602002602001604051908101604052809291908181526020016000905b828210156101ea578382906000526020600020906002020160405180604001604052908160008201805461010390610c5c565b80601f016020809104026020016040519081016040528092919081815260200182805461012f90610c5c565b801561017c5780601f106101515761010080835404028352916020019161017c565b820191906000526020600020905b81548152906001019060200180831161015f57829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525050815260200190600101906100d0565b50505050905090565b6003600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414610275576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026c90610cea565b60405180910390fd5b60003373ffffffffffffffffffffffffffffffffffffffff1664e8d4a510006040516102a090610d3b565b60006040518083038185875af1925050503d80600081146102dd576040519150601f19603f3d011682016040523d82523d6000602084013e6102e2565b606091505b5050905080610326576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031d90610d9c565b60405180910390fd5b6002600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050565b600260405180604001604052808381526020013373ffffffffffffffffffffffffffffffffffffffff16815250908060018154018082558091505060019003906000526020600020906002020160009091909190915060008201518160000190816103d99190610f72565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050506000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054146104a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161049c90611090565b60405180910390fd5b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008154809291906104f5906110df565b91905055506509184e72a000341015610543576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053a90610d9c565b60405180910390fd5b50565b600060018060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054146105c9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105c090611173565b60405180910390fd5b60005b600280549050811015610668573373ffffffffffffffffffffffffffffffffffffffff166002828154811061060457610603611193565b5b906000526020600020906002020160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610655578091505b8080610660906110df565b9150506105cc565b5060008190505b600160028054905061068191906111c2565b81101561076a57600260018261069791906111f6565b815481106106a8576106a7611193565b5b9060005260206000209060020201600282815481106106ca576106c9611193565b5b9060005260206000209060020201600082018160000190816106ec9190611240565b506001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050508080610762906110df565b91505061066f565b50600280548061077d5761077c611328565b5b6001900381819060005260206000209060020201600080820160006107a2919061084c565b6001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055505090556003600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507feed050308c603899d7397c26bdccda0810c3ccc6e9730a8a10c452b522f8edf4336040516108419190611366565b60405180910390a150565b50805461085890610c5c565b6000825580601f1061086a5750610889565b601f016020900490600052602060002090810190610888919061088c565b5b50565b5b808211156108a557600081600090555060010161088d565b5090565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561090f5780820151818401526020810190506108f4565b60008484015250505050565b6000601f19601f8301169050919050565b6000610937826108d5565b61094181856108e0565b93506109518185602086016108f1565b61095a8161091b565b840191505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061099082610965565b9050919050565b6109a081610985565b82525050565b600060408301600083015184820360008601526109c3828261092c565b91505060208301516109d86020860182610997565b508091505092915050565b60006109ef83836109a6565b905092915050565b6000602082019050919050565b6000610a0f826108a9565b610a1981856108b4565b935083602082028501610a2b856108c5565b8060005b85811015610a675784840389528151610a4885826109e3565b9450610a53836109f7565b925060208a01995050600181019050610a2f565b50829750879550505050505092915050565b60006020820190508181036000830152610a938184610a04565b905092915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610af18261091b565b810181811067ffffffffffffffff82111715610b1057610b0f610ab9565b5b80604052505050565b6000610b23610a9b565b9050610b2f8282610ae8565b919050565b600067ffffffffffffffff821115610b4f57610b4e610ab9565b5b610b588261091b565b9050602081019050919050565b82818337600083830152505050565b6000610b87610b8284610b34565b610b19565b905082815260208101848484011115610ba357610ba2610ab4565b5b610bae848285610b65565b509392505050565b600082601f830112610bcb57610bca610aaf565b5b8135610bdb848260208601610b74565b91505092915050565b600060208284031215610bfa57610bf9610aa5565b5b600082013567ffffffffffffffff811115610c1857610c17610aaa565b5b610c2484828501610bb6565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610c7457607f821691505b602082108103610c8757610c86610c2d565b5b50919050565b600082825260208201905092915050565b7f4661696c656420746f20726566756e6400000000000000000000000000000000600082015250565b6000610cd4601083610c8d565b9150610cdf82610c9e565b602082019050919050565b60006020820190508181036000830152610d0381610cc7565b9050919050565b600081905092915050565b50565b6000610d25600083610d0a565b9150610d3082610d15565b600082019050919050565b6000610d4682610d18565b9150819050919050565b7f4661696c656420746f2073656e64204574686572000000000000000000000000600082015250565b6000610d86601483610c8d565b9150610d9182610d50565b602082019050919050565b60006020820190508181036000830152610db581610d79565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610e1e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610de1565b610e288683610de1565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000610e6f610e6a610e6584610e40565b610e4a565b610e40565b9050919050565b6000819050919050565b610e8983610e54565b610e9d610e9582610e76565b848454610dee565b825550505050565b600090565b610eb2610ea5565b610ebd818484610e80565b505050565b5b81811015610ee157610ed6600082610eaa565b600181019050610ec3565b5050565b601f821115610f2657610ef781610dbc565b610f0084610dd1565b81016020851015610f0f578190505b610f23610f1b85610dd1565b830182610ec2565b50505b505050565b600082821c905092915050565b6000610f4960001984600802610f2b565b1980831691505092915050565b6000610f628383610f38565b9150826002028217905092915050565b610f7b826108d5565b67ffffffffffffffff811115610f9457610f93610ab9565b5b610f9e8254610c5c565b610fa9828285610ee5565b600060209050601f831160018114610fdc5760008415610fca578287015190505b610fd48582610f56565b86555061103c565b601f198416610fea86610dbc565b60005b8281101561101257848901518255600182019150602085019450602081019050610fed565b8683101561102f578489015161102b601f891682610f38565b8355505b6001600288020188555050505b505050505050565b7f5265676973746572206661696c65640000000000000000000000000000000000600082015250565b600061107a600f83610c8d565b915061108582611044565b602082019050919050565b600060208201905081810360008301526110a98161106d565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006110ea82610e40565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361111c5761111b6110b0565b5b600182019050919050565b7f556e737562736372696265206661696c65640000000000000000000000000000600082015250565b600061115d601283610c8d565b915061116882611127565b602082019050919050565b6000602082019050818103600083015261118c81611150565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006111cd82610e40565b91506111d883610e40565b92508282039050818111156111f0576111ef6110b0565b5b92915050565b600061120182610e40565b915061120c83610e40565b9250828201905080821115611224576112236110b0565b5b92915050565b60008154905061123981610c5c565b9050919050565b81810361124e575050611326565b6112578261122a565b67ffffffffffffffff8111156112705761126f610ab9565b5b61127a8254610c5c565b611285828285610ee5565b6000601f8311600181146112b457600084156112a2578287015490505b6112ac8582610f56565b86555061131f565b601f1984166112c287610dbc565b96506112cd86610dbc565b60005b828110156112f5578489015482556001820191506001850194506020810190506112d0565b86831015611312578489015461130e601f891682610f38565b8355505b6001600288020188555050505b5050505050505b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b61136081610985565b82525050565b600060208201905061137b6000830184611357565b9291505056fea2646970667358221220c5f07847737435960481d0823305e817a6a1a3b376ef03f292ef734ccfde726e64736f6c63430008110033",
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

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package todo

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

// TodoTask is an auto generated low-level Go binding around an user-defined struct.
type TodoTask struct {
	MStatus bool
	MData   string
}

// TodoMetaData contains all meta data concerning the Todo contract.
var TodoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"m_status\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"m_data\",\"type\":\"string\"}],\"internalType\":\"structTodo.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"list\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"m_status\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"m_data\",\"type\":\"string\"}],\"internalType\":\"structTodo.Task[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610f74806100616000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630f560cd7146100675780634cc82215146100855780638da5cb5b146100a15780639507d39a146100bf578063b0c8f9dc146100ef578063f745630f1461010b575b600080fd5b61006f610127565b60405161007c9190610a2b565b60405180910390f35b61009f600480360381019061009a9190610a97565b61028d565b005b6100a961040c565b6040516100b69190610b05565b60405180910390f35b6100d960048036038101906100d49190610a97565b610432565b6040516100e69190610b5d565b60405180910390f35b61010960048036038101906101049190610cb4565b610575565b005b61012560048036038101906101209190610cfd565b610654565b005b60603373ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461018357600080fd5b6000805480602002602001604051908101604052809291908181526020016000905b8282101561028457838290600052602060002090600202016040518060400160405290816000820160009054906101000a900460ff161515151581526020016001820180546101f390610d88565b80601f016020809104026020016040519081016040528092919081815260200182805461021f90610d88565b801561026c5780601f106102415761010080835404028352916020019161026c565b820191906000526020600020905b81548152906001019060200180831161024f57829003601f168201915b505050505081525050815260200190600101906101a5565b50505050905090565b3373ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146102e757600080fd5b60008190505b60016000805490506102ff9190610de8565b8110156103b95760006001826103159190610e1c565b8154811061032657610325610e72565b5b90600052602060002090600202016000828154811061034857610347610e72565b5b90600052602060002090600202016000820160009054906101000a900460ff168160000160006101000a81548160ff021916908315150217905550600182018160010190805461039790610d88565b6103a29291906106ec565b5090505080806103b190610ea1565b9150506102ed565b5060008054806103cc576103cb610ee9565b5b6001900381819060005260206000209060020201600080820160006101000a81549060ff02191690556001820160006104059190610779565b5050905550565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61043a6107b9565b3373ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461049457600080fd5b600082815481106104a8576104a7610e72565b5b90600052602060002090600202016040518060400160405290816000820160009054906101000a900460ff161515151581526020016001820180546104ec90610d88565b80601f016020809104026020016040519081016040528092919081815260200182805461051890610d88565b80156105655780601f1061053a57610100808354040283529160200191610565565b820191906000526020600020905b81548152906001019060200180831161054857829003601f168201915b5050505050815250509050919050565b3373ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146105cf57600080fd5b6000604051806040016040528060001515815260200183815250908060018154018082558091505060019003906000526020600020906002020160009091909190915060008201518160000160006101000a81548160ff021916908315150217905550602082015181600101908051906020019061064e9291906107d5565b50505050565b3373ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146106ae57600080fd5b80600083815481106106c3576106c2610e72565b5b906000526020600020906002020160010190805190602001906106e79291906107d5565b505050565b8280546106f890610d88565b90600052602060002090601f01602090048101928261071a5760008555610768565b82601f1061072b5780548555610768565b8280016001018555821561076857600052602060002091601f016020900482015b8281111561076757825482559160010191906001019061074c565b5b509050610775919061085b565b5090565b50805461078590610d88565b6000825580601f1061079757506107b6565b601f0160209004906000526020600020908101906107b5919061085b565b5b50565b6040518060400160405280600015158152602001606081525090565b8280546107e190610d88565b90600052602060002090601f016020900481019282610803576000855561084a565b82601f1061081c57805160ff191683800117855561084a565b8280016001018555821561084a579182015b8281111561084957825182559160200191906001019061082e565b5b509050610857919061085b565b5090565b5b8082111561087457600081600090555060010161085c565b5090565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60008115159050919050565b6108b9816108a4565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b838110156108f95780820151818401526020810190506108de565b83811115610908576000848401525b50505050565b6000601f19601f8301169050919050565b600061092a826108bf565b61093481856108ca565b93506109448185602086016108db565b61094d8161090e565b840191505092915050565b600060408301600083015161097060008601826108b0565b5060208301518482036020860152610988828261091f565b9150508091505092915050565b60006109a18383610958565b905092915050565b6000602082019050919050565b60006109c182610878565b6109cb8185610883565b9350836020820285016109dd85610894565b8060005b85811015610a1957848403895281516109fa8582610995565b9450610a05836109a9565b925060208a019950506001810190506109e1565b50829750879550505050505092915050565b60006020820190508181036000830152610a4581846109b6565b905092915050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610a7481610a61565b8114610a7f57600080fd5b50565b600081359050610a9181610a6b565b92915050565b600060208284031215610aad57610aac610a57565b5b6000610abb84828501610a82565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610aef82610ac4565b9050919050565b610aff81610ae4565b82525050565b6000602082019050610b1a6000830184610af6565b92915050565b6000604083016000830151610b3860008601826108b0565b5060208301518482036020860152610b50828261091f565b9150508091505092915050565b60006020820190508181036000830152610b778184610b20565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610bc18261090e565b810181811067ffffffffffffffff82111715610be057610bdf610b89565b5b80604052505050565b6000610bf3610a4d565b9050610bff8282610bb8565b919050565b600067ffffffffffffffff821115610c1f57610c1e610b89565b5b610c288261090e565b9050602081019050919050565b82818337600083830152505050565b6000610c57610c5284610c04565b610be9565b905082815260208101848484011115610c7357610c72610b84565b5b610c7e848285610c35565b509392505050565b600082601f830112610c9b57610c9a610b7f565b5b8135610cab848260208601610c44565b91505092915050565b600060208284031215610cca57610cc9610a57565b5b600082013567ffffffffffffffff811115610ce857610ce7610a5c565b5b610cf484828501610c86565b91505092915050565b60008060408385031215610d1457610d13610a57565b5b6000610d2285828601610a82565b925050602083013567ffffffffffffffff811115610d4357610d42610a5c565b5b610d4f85828601610c86565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610da057607f821691505b602082108103610db357610db2610d59565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610df382610a61565b9150610dfe83610a61565b925082821015610e1157610e10610db9565b5b828203905092915050565b6000610e2782610a61565b9150610e3283610a61565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610e6757610e66610db9565b5b828201905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000610eac82610a61565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610ede57610edd610db9565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea26469706673582212205041bb6e1bac6cfcbdb54085d0c9066dd77929e41a229e33cf15031690f5ecbb64736f6c637828302e382e31342d646576656c6f702e323032322e352e31322b636f6d6d69742e30633066663466630059",
}

// TodoABI is the input ABI used to generate the binding from.
// Deprecated: Use TodoMetaData.ABI instead.
var TodoABI = TodoMetaData.ABI

// TodoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TodoMetaData.Bin instead.
var TodoBin = TodoMetaData.Bin

// DeployTodo deploys a new Ethereum contract, binding an instance of Todo to it.
func DeployTodo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Todo, error) {
	parsed, err := TodoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TodoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// Todo is an auto generated Go binding around an Ethereum contract.
type Todo struct {
	TodoCaller     // Read-only binding to the contract
	TodoTransactor // Write-only binding to the contract
	TodoFilterer   // Log filterer for contract events
}

// TodoCaller is an auto generated read-only Go binding around an Ethereum contract.
type TodoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TodoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TodoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TodoSession struct {
	Contract     *Todo             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TodoCallerSession struct {
	Contract *TodoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TodoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TodoTransactorSession struct {
	Contract     *TodoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoRaw is an auto generated low-level Go binding around an Ethereum contract.
type TodoRaw struct {
	Contract *Todo // Generic contract binding to access the raw methods on
}

// TodoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TodoCallerRaw struct {
	Contract *TodoCaller // Generic read-only contract binding to access the raw methods on
}

// TodoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TodoTransactorRaw struct {
	Contract *TodoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTodo creates a new instance of Todo, bound to a specific deployed contract.
func NewTodo(address common.Address, backend bind.ContractBackend) (*Todo, error) {
	contract, err := bindTodo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// NewTodoCaller creates a new read-only instance of Todo, bound to a specific deployed contract.
func NewTodoCaller(address common.Address, caller bind.ContractCaller) (*TodoCaller, error) {
	contract, err := bindTodo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TodoCaller{contract: contract}, nil
}

// NewTodoTransactor creates a new write-only instance of Todo, bound to a specific deployed contract.
func NewTodoTransactor(address common.Address, transactor bind.ContractTransactor) (*TodoTransactor, error) {
	contract, err := bindTodo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TodoTransactor{contract: contract}, nil
}

// NewTodoFilterer creates a new log filterer instance of Todo, bound to a specific deployed contract.
func NewTodoFilterer(address common.Address, filterer bind.ContractFilterer) (*TodoFilterer, error) {
	contract, err := bindTodo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TodoFilterer{contract: contract}, nil
}

// bindTodo binds a generic wrapper to an already deployed contract.
func bindTodo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TodoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.TodoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 id) view returns((bool,string))
func (_Todo *TodoCaller) Get(opts *bind.CallOpts, id *big.Int) (TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "get", id)

	if err != nil {
		return *new(TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new(TodoTask)).(*TodoTask)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 id) view returns((bool,string))
func (_Todo *TodoSession) Get(id *big.Int) (TodoTask, error) {
	return _Todo.Contract.Get(&_Todo.CallOpts, id)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 id) view returns((bool,string))
func (_Todo *TodoCallerSession) Get(id *big.Int) (TodoTask, error) {
	return _Todo.Contract.Get(&_Todo.CallOpts, id)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((bool,string)[])
func (_Todo *TodoCaller) List(opts *bind.CallOpts) ([]TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "list")

	if err != nil {
		return *new([]TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new([]TodoTask)).(*[]TodoTask)

	return out0, err

}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((bool,string)[])
func (_Todo *TodoSession) List() ([]TodoTask, error) {
	return _Todo.Contract.List(&_Todo.CallOpts)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((bool,string)[])
func (_Todo *TodoCallerSession) List() ([]TodoTask, error) {
	return _Todo.Contract.List(&_Todo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoSession) Owner() (common.Address, error) {
	return _Todo.Contract.Owner(&_Todo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoCallerSession) Owner() (common.Address, error) {
	return _Todo.Contract.Owner(&_Todo.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string data) returns()
func (_Todo *TodoTransactor) Add(opts *bind.TransactOpts, data string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "add", data)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string data) returns()
func (_Todo *TodoSession) Add(data string) (*types.Transaction, error) {
	return _Todo.Contract.Add(&_Todo.TransactOpts, data)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string data) returns()
func (_Todo *TodoTransactorSession) Add(data string) (*types.Transaction, error) {
	return _Todo.Contract.Add(&_Todo.TransactOpts, data)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 id) returns()
func (_Todo *TodoTransactor) Remove(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "remove", id)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 id) returns()
func (_Todo *TodoSession) Remove(id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Remove(&_Todo.TransactOpts, id)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 id) returns()
func (_Todo *TodoTransactorSession) Remove(id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Remove(&_Todo.TransactOpts, id)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 id, string data) returns()
func (_Todo *TodoTransactor) Update(opts *bind.TransactOpts, id *big.Int, data string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "update", id, data)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 id, string data) returns()
func (_Todo *TodoSession) Update(id *big.Int, data string) (*types.Transaction, error) {
	return _Todo.Contract.Update(&_Todo.TransactOpts, id, data)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 id, string data) returns()
func (_Todo *TodoTransactorSession) Update(id *big.Int, data string) (*types.Transaction, error) {
	return _Todo.Contract.Update(&_Todo.TransactOpts, id, data)
}

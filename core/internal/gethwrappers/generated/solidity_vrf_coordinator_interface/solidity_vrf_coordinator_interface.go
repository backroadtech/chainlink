// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package solidity_vrf_coordinator_interface

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// LinkTokenInterfaceABI is the input ABI used to generate the binding from.
const LinkTokenInterfaceABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"decimalPlaces\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalTokensIssued\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// LinkTokenInterfaceFuncSigs maps the 4-byte function signature to its string representation.
var LinkTokenInterfaceFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"66188463": "decreaseApproval(address,uint256)",
	"d73dd623": "increaseApproval(address,uint256)",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"4000aea0": "transferAndCall(address,uint256,bytes)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// LinkTokenInterface is an auto generated Go binding around an Ethereum contract.
type LinkTokenInterface struct {
	LinkTokenInterfaceCaller     // Read-only binding to the contract
	LinkTokenInterfaceTransactor // Write-only binding to the contract
	LinkTokenInterfaceFilterer   // Log filterer for contract events
}

// LinkTokenInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type LinkTokenInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LinkTokenInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LinkTokenInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LinkTokenInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LinkTokenInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LinkTokenInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LinkTokenInterfaceSession struct {
	Contract     *LinkTokenInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LinkTokenInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LinkTokenInterfaceCallerSession struct {
	Contract *LinkTokenInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// LinkTokenInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LinkTokenInterfaceTransactorSession struct {
	Contract     *LinkTokenInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// LinkTokenInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type LinkTokenInterfaceRaw struct {
	Contract *LinkTokenInterface // Generic contract binding to access the raw methods on
}

// LinkTokenInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LinkTokenInterfaceCallerRaw struct {
	Contract *LinkTokenInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// LinkTokenInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LinkTokenInterfaceTransactorRaw struct {
	Contract *LinkTokenInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLinkTokenInterface creates a new instance of LinkTokenInterface, bound to a specific deployed contract.
func NewLinkTokenInterface(address common.Address, backend bind.ContractBackend) (*LinkTokenInterface, error) {
	contract, err := bindLinkTokenInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LinkTokenInterface{LinkTokenInterfaceCaller: LinkTokenInterfaceCaller{contract: contract}, LinkTokenInterfaceTransactor: LinkTokenInterfaceTransactor{contract: contract}, LinkTokenInterfaceFilterer: LinkTokenInterfaceFilterer{contract: contract}}, nil
}

// NewLinkTokenInterfaceCaller creates a new read-only instance of LinkTokenInterface, bound to a specific deployed contract.
func NewLinkTokenInterfaceCaller(address common.Address, caller bind.ContractCaller) (*LinkTokenInterfaceCaller, error) {
	contract, err := bindLinkTokenInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LinkTokenInterfaceCaller{contract: contract}, nil
}

// NewLinkTokenInterfaceTransactor creates a new write-only instance of LinkTokenInterface, bound to a specific deployed contract.
func NewLinkTokenInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*LinkTokenInterfaceTransactor, error) {
	contract, err := bindLinkTokenInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LinkTokenInterfaceTransactor{contract: contract}, nil
}

// NewLinkTokenInterfaceFilterer creates a new log filterer instance of LinkTokenInterface, bound to a specific deployed contract.
func NewLinkTokenInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*LinkTokenInterfaceFilterer, error) {
	contract, err := bindLinkTokenInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LinkTokenInterfaceFilterer{contract: contract}, nil
}

// bindLinkTokenInterface binds a generic wrapper to an already deployed contract.
func bindLinkTokenInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LinkTokenInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LinkTokenInterface *LinkTokenInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LinkTokenInterface.Contract.LinkTokenInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LinkTokenInterface *LinkTokenInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.LinkTokenInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LinkTokenInterface *LinkTokenInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.LinkTokenInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LinkTokenInterface *LinkTokenInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LinkTokenInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LinkTokenInterface *LinkTokenInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LinkTokenInterface *LinkTokenInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256 remaining)
func (_LinkTokenInterface *LinkTokenInterfaceCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LinkTokenInterface.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256 remaining)
func (_LinkTokenInterface *LinkTokenInterfaceSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LinkTokenInterface.Contract.Allowance(&_LinkTokenInterface.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256 remaining)
func (_LinkTokenInterface *LinkTokenInterfaceCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LinkTokenInterface.Contract.Allowance(&_LinkTokenInterface.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_LinkTokenInterface *LinkTokenInterfaceCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LinkTokenInterface.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_LinkTokenInterface *LinkTokenInterfaceSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LinkTokenInterface.Contract.BalanceOf(&_LinkTokenInterface.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_LinkTokenInterface *LinkTokenInterfaceCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LinkTokenInterface.Contract.BalanceOf(&_LinkTokenInterface.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8 decimalPlaces)
func (_LinkTokenInterface *LinkTokenInterfaceCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _LinkTokenInterface.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8 decimalPlaces)
func (_LinkTokenInterface *LinkTokenInterfaceSession) Decimals() (uint8, error) {
	return _LinkTokenInterface.Contract.Decimals(&_LinkTokenInterface.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8 decimalPlaces)
func (_LinkTokenInterface *LinkTokenInterfaceCallerSession) Decimals() (uint8, error) {
	return _LinkTokenInterface.Contract.Decimals(&_LinkTokenInterface.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string tokenName)
func (_LinkTokenInterface *LinkTokenInterfaceCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _LinkTokenInterface.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string tokenName)
func (_LinkTokenInterface *LinkTokenInterfaceSession) Name() (string, error) {
	return _LinkTokenInterface.Contract.Name(&_LinkTokenInterface.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string tokenName)
func (_LinkTokenInterface *LinkTokenInterfaceCallerSession) Name() (string, error) {
	return _LinkTokenInterface.Contract.Name(&_LinkTokenInterface.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string tokenSymbol)
func (_LinkTokenInterface *LinkTokenInterfaceCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _LinkTokenInterface.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string tokenSymbol)
func (_LinkTokenInterface *LinkTokenInterfaceSession) Symbol() (string, error) {
	return _LinkTokenInterface.Contract.Symbol(&_LinkTokenInterface.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string tokenSymbol)
func (_LinkTokenInterface *LinkTokenInterfaceCallerSession) Symbol() (string, error) {
	return _LinkTokenInterface.Contract.Symbol(&_LinkTokenInterface.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256 totalTokensIssued)
func (_LinkTokenInterface *LinkTokenInterfaceCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LinkTokenInterface.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256 totalTokensIssued)
func (_LinkTokenInterface *LinkTokenInterfaceSession) TotalSupply() (*big.Int, error) {
	return _LinkTokenInterface.Contract.TotalSupply(&_LinkTokenInterface.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256 totalTokensIssued)
func (_LinkTokenInterface *LinkTokenInterfaceCallerSession) TotalSupply() (*big.Int, error) {
	return _LinkTokenInterface.Contract.TotalSupply(&_LinkTokenInterface.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool success)
func (_LinkTokenInterface *LinkTokenInterfaceTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool success)
func (_LinkTokenInterface *LinkTokenInterfaceSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.Approve(&_LinkTokenInterface.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool success)
func (_LinkTokenInterface *LinkTokenInterfaceTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.Approve(&_LinkTokenInterface.TransactOpts, spender, value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address spender, uint256 addedValue) returns(bool success)
func (_LinkTokenInterface *LinkTokenInterfaceTransactor) DecreaseApproval(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.contract.Transact(opts, "decreaseApproval", spender, addedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address spender, uint256 addedValue) returns(bool success)
func (_LinkTokenInterface *LinkTokenInterfaceSession) DecreaseApproval(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.DecreaseApproval(&_LinkTokenInterface.TransactOpts, spender, addedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address spender, uint256 addedValue) returns(bool success)
func (_LinkTokenInterface *LinkTokenInterfaceTransactorSession) DecreaseApproval(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.DecreaseApproval(&_LinkTokenInterface.TransactOpts, spender, addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address spender, uint256 subtractedValue) returns()
func (_LinkTokenInterface *LinkTokenInterfaceTransactor) IncreaseApproval(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.contract.Transact(opts, "increaseApproval", spender, subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address spender, uint256 subtractedValue) returns()
func (_LinkTokenInterface *LinkTokenInterfaceSession) IncreaseApproval(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.IncreaseApproval(&_LinkTokenInterface.TransactOpts, spender, subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address spender, uint256 subtractedValue) returns()
func (_LinkTokenInterface *LinkTokenInterfaceTransactorSession) IncreaseApproval(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.IncreaseApproval(&_LinkTokenInterface.TransactOpts, spender, subtractedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool success)
func (_LinkTokenInterface *LinkTokenInterfaceTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool success)
func (_LinkTokenInterface *LinkTokenInterfaceSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.Transfer(&_LinkTokenInterface.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool success)
func (_LinkTokenInterface *LinkTokenInterfaceTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.Transfer(&_LinkTokenInterface.TransactOpts, to, value)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool success)
func (_LinkTokenInterface *LinkTokenInterfaceTransactor) TransferAndCall(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _LinkTokenInterface.contract.Transact(opts, "transferAndCall", to, value, data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool success)
func (_LinkTokenInterface *LinkTokenInterfaceSession) TransferAndCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.TransferAndCall(&_LinkTokenInterface.TransactOpts, to, value, data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool success)
func (_LinkTokenInterface *LinkTokenInterfaceTransactorSession) TransferAndCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.TransferAndCall(&_LinkTokenInterface.TransactOpts, to, value, data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool success)
func (_LinkTokenInterface *LinkTokenInterfaceTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool success)
func (_LinkTokenInterface *LinkTokenInterfaceSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.TransferFrom(&_LinkTokenInterface.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool success)
func (_LinkTokenInterface *LinkTokenInterfaceTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.TransferFrom(&_LinkTokenInterface.TransactOpts, from, to, value)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e06291179dac2cbce3b0ddb79fbdf77ddd37b1a3ef709d3fe8250aa7a4b0c12164736f6c63430006020033"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// VRFABI is the input ABI used to generate the binding from.
const VRFABI = "[]"

// VRFBin is the compiled bytecode used for deploying new contracts.
var VRFBin = "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea2646970667358221220546e951fb8ad1b411342104a2f15fc5f8a5c6d4d9a0da164463af437a5405ca264736f6c63430006020033"

// DeployVRF deploys a new Ethereum contract, binding an instance of VRF to it.
func DeployVRF(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VRF, error) {
	parsed, err := abi.JSON(strings.NewReader(VRFABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VRFBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRF{VRFCaller: VRFCaller{contract: contract}, VRFTransactor: VRFTransactor{contract: contract}, VRFFilterer: VRFFilterer{contract: contract}}, nil
}

// VRF is an auto generated Go binding around an Ethereum contract.
type VRF struct {
	VRFCaller     // Read-only binding to the contract
	VRFTransactor // Write-only binding to the contract
	VRFFilterer   // Log filterer for contract events
}

// VRFCaller is an auto generated read-only Go binding around an Ethereum contract.
type VRFCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VRFTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VRFFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VRFSession struct {
	Contract     *VRF              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VRFCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VRFCallerSession struct {
	Contract *VRFCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VRFTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VRFTransactorSession struct {
	Contract     *VRFTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VRFRaw is an auto generated low-level Go binding around an Ethereum contract.
type VRFRaw struct {
	Contract *VRF // Generic contract binding to access the raw methods on
}

// VRFCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VRFCallerRaw struct {
	Contract *VRFCaller // Generic read-only contract binding to access the raw methods on
}

// VRFTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VRFTransactorRaw struct {
	Contract *VRFTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVRF creates a new instance of VRF, bound to a specific deployed contract.
func NewVRF(address common.Address, backend bind.ContractBackend) (*VRF, error) {
	contract, err := bindVRF(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRF{VRFCaller: VRFCaller{contract: contract}, VRFTransactor: VRFTransactor{contract: contract}, VRFFilterer: VRFFilterer{contract: contract}}, nil
}

// NewVRFCaller creates a new read-only instance of VRF, bound to a specific deployed contract.
func NewVRFCaller(address common.Address, caller bind.ContractCaller) (*VRFCaller, error) {
	contract, err := bindVRF(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCaller{contract: contract}, nil
}

// NewVRFTransactor creates a new write-only instance of VRF, bound to a specific deployed contract.
func NewVRFTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFTransactor, error) {
	contract, err := bindVRF(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFTransactor{contract: contract}, nil
}

// NewVRFFilterer creates a new log filterer instance of VRF, bound to a specific deployed contract.
func NewVRFFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFFilterer, error) {
	contract, err := bindVRF(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFFilterer{contract: contract}, nil
}

// bindVRF binds a generic wrapper to an already deployed contract.
func bindVRF(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VRFABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRF *VRFRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VRF.Contract.VRFCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRF *VRFRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRF.Contract.VRFTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRF *VRFRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRF.Contract.VRFTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRF *VRFCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VRF.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRF *VRFTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRF.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRF *VRFTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRF.Contract.contract.Transact(opts, method, params...)
}

// VRFConsumerBaseABI is the input ABI used to generate the binding from.
const VRFConsumerBaseABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vrfCoordinator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_link\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"randomness\",\"type\":\"uint256\"}],\"name\":\"fulfillRandomness_\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_seed\",\"type\":\"uint256\"}],\"name\":\"requestRandomness\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// VRFConsumerBaseFuncSigs maps the 4-byte function signature to its string representation.
var VRFConsumerBaseFuncSigs = map[string]string{
	"0441c5fb": "fulfillRandomness_(bytes32,uint256)",
	"9e317f12": "nonces(bytes32)",
	"dc6cfe10": "requestRandomness(bytes32,uint256,uint256)",
}

// VRFConsumerBase is an auto generated Go binding around an Ethereum contract.
type VRFConsumerBase struct {
	VRFConsumerBaseCaller     // Read-only binding to the contract
	VRFConsumerBaseTransactor // Write-only binding to the contract
	VRFConsumerBaseFilterer   // Log filterer for contract events
}

// VRFConsumerBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type VRFConsumerBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFConsumerBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VRFConsumerBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFConsumerBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VRFConsumerBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFConsumerBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VRFConsumerBaseSession struct {
	Contract     *VRFConsumerBase  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VRFConsumerBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VRFConsumerBaseCallerSession struct {
	Contract *VRFConsumerBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// VRFConsumerBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VRFConsumerBaseTransactorSession struct {
	Contract     *VRFConsumerBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// VRFConsumerBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type VRFConsumerBaseRaw struct {
	Contract *VRFConsumerBase // Generic contract binding to access the raw methods on
}

// VRFConsumerBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VRFConsumerBaseCallerRaw struct {
	Contract *VRFConsumerBaseCaller // Generic read-only contract binding to access the raw methods on
}

// VRFConsumerBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VRFConsumerBaseTransactorRaw struct {
	Contract *VRFConsumerBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVRFConsumerBase creates a new instance of VRFConsumerBase, bound to a specific deployed contract.
func NewVRFConsumerBase(address common.Address, backend bind.ContractBackend) (*VRFConsumerBase, error) {
	contract, err := bindVRFConsumerBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFConsumerBase{VRFConsumerBaseCaller: VRFConsumerBaseCaller{contract: contract}, VRFConsumerBaseTransactor: VRFConsumerBaseTransactor{contract: contract}, VRFConsumerBaseFilterer: VRFConsumerBaseFilterer{contract: contract}}, nil
}

// NewVRFConsumerBaseCaller creates a new read-only instance of VRFConsumerBase, bound to a specific deployed contract.
func NewVRFConsumerBaseCaller(address common.Address, caller bind.ContractCaller) (*VRFConsumerBaseCaller, error) {
	contract, err := bindVRFConsumerBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFConsumerBaseCaller{contract: contract}, nil
}

// NewVRFConsumerBaseTransactor creates a new write-only instance of VRFConsumerBase, bound to a specific deployed contract.
func NewVRFConsumerBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFConsumerBaseTransactor, error) {
	contract, err := bindVRFConsumerBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFConsumerBaseTransactor{contract: contract}, nil
}

// NewVRFConsumerBaseFilterer creates a new log filterer instance of VRFConsumerBase, bound to a specific deployed contract.
func NewVRFConsumerBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFConsumerBaseFilterer, error) {
	contract, err := bindVRFConsumerBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFConsumerBaseFilterer{contract: contract}, nil
}

// bindVRFConsumerBase binds a generic wrapper to an already deployed contract.
func bindVRFConsumerBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VRFConsumerBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRFConsumerBase *VRFConsumerBaseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VRFConsumerBase.Contract.VRFConsumerBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRFConsumerBase *VRFConsumerBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFConsumerBase.Contract.VRFConsumerBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRFConsumerBase *VRFConsumerBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFConsumerBase.Contract.VRFConsumerBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRFConsumerBase *VRFConsumerBaseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VRFConsumerBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRFConsumerBase *VRFConsumerBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFConsumerBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRFConsumerBase *VRFConsumerBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFConsumerBase.Contract.contract.Transact(opts, method, params...)
}

// Nonces is a free data retrieval call binding the contract method 0x9e317f12.
//
// Solidity: function nonces(bytes32 ) constant returns(uint256)
func (_VRFConsumerBase *VRFConsumerBaseCaller) Nonces(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VRFConsumerBase.contract.Call(opts, out, "nonces", arg0)
	return *ret0, err
}

// Nonces is a free data retrieval call binding the contract method 0x9e317f12.
//
// Solidity: function nonces(bytes32 ) constant returns(uint256)
func (_VRFConsumerBase *VRFConsumerBaseSession) Nonces(arg0 [32]byte) (*big.Int, error) {
	return _VRFConsumerBase.Contract.Nonces(&_VRFConsumerBase.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x9e317f12.
//
// Solidity: function nonces(bytes32 ) constant returns(uint256)
func (_VRFConsumerBase *VRFConsumerBaseCallerSession) Nonces(arg0 [32]byte) (*big.Int, error) {
	return _VRFConsumerBase.Contract.Nonces(&_VRFConsumerBase.CallOpts, arg0)
}

// FulfillRandomness is a paid mutator transaction binding the contract method 0x0441c5fb.
//
// Solidity: function fulfillRandomness_(bytes32 requestId, uint256 randomness) returns()
func (_VRFConsumerBase *VRFConsumerBaseTransactor) FulfillRandomness(opts *bind.TransactOpts, requestId [32]byte, randomness *big.Int) (*types.Transaction, error) {
	return _VRFConsumerBase.contract.Transact(opts, "fulfillRandomness_", requestId, randomness)
}

// FulfillRandomness is a paid mutator transaction binding the contract method 0x0441c5fb.
//
// Solidity: function fulfillRandomness_(bytes32 requestId, uint256 randomness) returns()
func (_VRFConsumerBase *VRFConsumerBaseSession) FulfillRandomness(requestId [32]byte, randomness *big.Int) (*types.Transaction, error) {
	return _VRFConsumerBase.Contract.FulfillRandomness(&_VRFConsumerBase.TransactOpts, requestId, randomness)
}

// FulfillRandomness is a paid mutator transaction binding the contract method 0x0441c5fb.
//
// Solidity: function fulfillRandomness_(bytes32 requestId, uint256 randomness) returns()
func (_VRFConsumerBase *VRFConsumerBaseTransactorSession) FulfillRandomness(requestId [32]byte, randomness *big.Int) (*types.Transaction, error) {
	return _VRFConsumerBase.Contract.FulfillRandomness(&_VRFConsumerBase.TransactOpts, requestId, randomness)
}

// RequestRandomness is a paid mutator transaction binding the contract method 0xdc6cfe10.
//
// Solidity: function requestRandomness(bytes32 _keyHash, uint256 _fee, uint256 _seed) returns(bytes32 requestId)
func (_VRFConsumerBase *VRFConsumerBaseTransactor) RequestRandomness(opts *bind.TransactOpts, _keyHash [32]byte, _fee *big.Int, _seed *big.Int) (*types.Transaction, error) {
	return _VRFConsumerBase.contract.Transact(opts, "requestRandomness", _keyHash, _fee, _seed)
}

// RequestRandomness is a paid mutator transaction binding the contract method 0xdc6cfe10.
//
// Solidity: function requestRandomness(bytes32 _keyHash, uint256 _fee, uint256 _seed) returns(bytes32 requestId)
func (_VRFConsumerBase *VRFConsumerBaseSession) RequestRandomness(_keyHash [32]byte, _fee *big.Int, _seed *big.Int) (*types.Transaction, error) {
	return _VRFConsumerBase.Contract.RequestRandomness(&_VRFConsumerBase.TransactOpts, _keyHash, _fee, _seed)
}

// RequestRandomness is a paid mutator transaction binding the contract method 0xdc6cfe10.
//
// Solidity: function requestRandomness(bytes32 _keyHash, uint256 _fee, uint256 _seed) returns(bytes32 requestId)
func (_VRFConsumerBase *VRFConsumerBaseTransactorSession) RequestRandomness(_keyHash [32]byte, _fee *big.Int, _seed *big.Int) (*types.Transaction, error) {
	return _VRFConsumerBase.Contract.RequestRandomness(&_VRFConsumerBase.TransactOpts, _keyHash, _fee, _seed)
}

// VRFCoordinatorABI is the input ABI used to generate the binding from.
const VRFCoordinatorABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_link\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"NewServiceAgreement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"jobID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"RandomnessRequest\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"callbacks\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"callbackContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"randomnessFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"fulfillRandomnessRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_publicKey\",\"type\":\"uint256[2]\"}],\"name\":\"hashOfKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"_publicProvingKey\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32\",\"name\":\"_jobID\",\"type\":\"bytes32\"}],\"name\":\"registerProvingKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"serviceAgreements\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"vRFOracle\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"jobID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawableTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VRFCoordinatorFuncSigs maps the 4-byte function signature to its string representation.
var VRFCoordinatorFuncSigs = map[string]string{
	"21f36509": "callbacks(bytes32)",
	"5e1c1059": "fulfillRandomnessRequest(bytes)",
	"caf70c4a": "hashOfKey(uint256[2])",
	"a4c0ed36": "onTokenTransfer(address,uint256,bytes)",
	"d8340209": "registerProvingKey(uint256,address,uint256[2],bytes32)",
	"75d35070": "serviceAgreements(bytes32)",
	"f3fef3a3": "withdraw(address,uint256)",
	"006f6ad0": "withdrawableTokens(address)",
}

// VRFCoordinatorBin is the compiled bytecode used for deploying new contracts.
var VRFCoordinatorBin = "0x608060405234801561001057600080fd5b50604051611a16380380611a168339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b03199092169190911790556119b1806100656000396000f3fe608060405234801561001057600080fd5b50600436106100875760003560e01c8063a4c0ed361161005b578063a4c0ed36146101e0578063caf70c4a1461029d578063d8340209146102e8578063f3fef3a31461031f57610087565b80626f6ad01461008c57806321f36509146100c45780635e1c10591461010957806375d35070146101c3575b600080fd5b6100b2600480360360208110156100a257600080fd5b50356001600160a01b031661034b565b60408051918252519081900360200190f35b6100e1600480360360208110156100da57600080fd5b503561035d565b604080516001600160a01b039094168452602084019290925282820152519081900360600190f35b6101af6004803603602081101561011f57600080fd5b81019060208101813564010000000081111561013a57600080fd5b82018360208201111561014c57600080fd5b8035906020019184600183028401116401000000008311171561016e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610389945050505050565b604080519115158252519081900360200190f35b6100e1600480360360208110156101d957600080fd5b50356105d2565b61029b600480360360608110156101f657600080fd5b6001600160a01b038235169160208101359181019060608101604082013564010000000081111561022657600080fd5b82018360208201111561023857600080fd5b8035906020019184600183028401116401000000008311171561025a57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506105fe945050505050565b005b6100b2600480360360408110156102b357600080fd5b6040805180820182529183019291818301918390600290839083908082843760009201919091525091945061068c9350505050565b61029b600480360360a08110156102fe57600080fd5b508035906001600160a01b03602082013516906040810190608001356106e1565b61029b6004803603604081101561033557600080fd5b506001600160a01b038135169060200135610851565b60036020526000908152604090205481565b60016020819052600091825260409091208054918101546002909101546001600160a01b039092169183565b6000610393611896565b5060e0820151602083019060006103a98361068c565b905060006103b78284610973565b90506103c16118b4565b50600081815260016020818152604092839020835160608101855281546001600160a01b031680825293820154928101929092526002015492810192909252610451576040805162461bcd60e51b815260206004820152601860248201527f6e6f20636f72726573706f6e64696e6720726571756573740000000000000000604482015290519081900360640190fd5b600061045c886109a2565b600085815260026020908152604080832054868301516001600160a01b0390911680855260039093529220549293509161049b9163ffffffff610aeb16565b6001600160a01b038083166000908152600360209081526040808320949094558351602481018990526044808201889052855180830390910181526064909101855290810180516001600160e01b0316630441c5fb60e01b178152875194518251939592948694931692859282918083835b6020831061052c5780518252601f19909201916020918201910161050d565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d806000811461058e576040519150601f19603f3d011682016040523d82523d6000602084013e610593565b606091505b50506000978852600160208190526040892080546001600160a01b0319168155908101899055600201979097555094985050505050505050505b919050565b60026020819052600091825260409091208054600182015491909201546001600160a01b039092169183565b6000546001600160a01b03163314610653576040805162461bcd60e51b815260206004820152601360248201527226bab9ba103ab9b2902624a725903a37b5b2b760691b604482015290519081900360640190fd5b60008082806020019051604081101561066b57600080fd5b508051602090910151909250905061068582828688610b4c565b5050505050565b6000816040516020018082600260200280838360005b838110156106ba5781810151838201526020016106a2565b50505050905001915050604051602081830303815290604052805190602001209050919050565b60408051808201825260009161071091908590600290839083908082843760009201919091525061068c915050565b6000818152600260205260409020549091506001600160a01b0316801561077e576040805162461bcd60e51b815260206004820152601960248201527f706c656173652072656769737465722061206e6577206b657900000000000000604482015290519081900360640190fd5b6001600160a01b0385166107d9576040805162461bcd60e51b815260206004820152601760248201527f5f6f7261636c65206d757374206e6f7420626520307830000000000000000000604482015290519081900360640190fd5b60008281526002602081815260409283902080546001600160a01b0319166001600160a01b038a1617815560018101879055909101889055815184815290810188905281517fae189157e0628c1e62315e9179156e1ea10e90e9c15060002f7021e907dc2cfe929181900390910190a1505050505050565b3360009081526003602052604090205481908111156108b7576040805162461bcd60e51b815260206004820181905260248201527f63616e2774207769746864726177206d6f7265207468616e2062616c616e6365604482015290519081900360640190fd5b336000908152600360205260409020546108d7908363ffffffff610d0416565b336000908152600360209081526040808320939093558154835163a9059cbb60e01b81526001600160a01b038881166004830152602482018890529451949091169363a9059cbb93604480840194938390030190829087803b15801561093c57600080fd5b505af1158015610950573d6000803e3d6000fd5b505050506040513d602081101561096657600080fd5b505161096e57fe5b505050565b604080516020808201859052818301849052825180830384018152606090920190925280519101205b92915050565b60006101a08251146109f0576040805162461bcd60e51b81526020600482015260126024820152710eee4dedcce40e0e4dedecc40d8cadccee8d60731b604482015290519081900360640190fd5b6109f8611896565b610a00611896565b610a086118de565b6000610a12611896565b610a1a611896565b6000888060200190516101a0811015610a3257600080fd5b5060e0810151610180820151919850604089019750608089019650945061010088019350610140880192509050610a85878787600060200201518860016020020151896002602002015189898989610d61565b6003866040516020018083815260200182600260200280838360005b83811015610ab9578181015183820152602001610aa1565b50505050905001925050506040516020818303038152906040528051906020012060001c975050505050505050919050565b600082820183811015610b45576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b81846002600082815260200190815260200160002060020154821015610bb0576040805162461bcd60e51b815260206004820152601460248201527310995b1bddc81859dc995959081c185e5b595b9d60621b604482015290519081900360640190fd5b60008681526004602090815260408083206001600160a01b038716845290915281205490610be088888785610fae565b90506000610bee8983610973565b6000818152600160205260409020549091506001600160a01b031615610c1057fe5b600081815260016020818152604080842080546001600160a01b038c166001600160a01b0319909116811782558185018d905560029182018890558e8652908352938190209092015482518d815291820186905281830193909352606081018a905290517fd241d78a52145a5d1d1ff002e32ec15cdc395631bcee66246650c2429dfaccaa9181900360800190a260008981526004602090815260408083206001600160a01b038a168452909152902054610cd290600163ffffffff610aeb16565b6000998a52600460209081526040808c206001600160a01b039099168c52979052959098209490945550505050505050565b600082821115610d5b576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b610d6a89610ff5565b610dbb576040805162461bcd60e51b815260206004820152601a60248201527f7075626c6963206b6579206973206e6f74206f6e206375727665000000000000604482015290519081900360640190fd5b610dc488610ff5565b610e0d576040805162461bcd60e51b815260206004820152601560248201527467616d6d61206973206e6f74206f6e20637572766560581b604482015290519081900360640190fd5b610e1683610ff5565b610e67576040805162461bcd60e51b815260206004820152601d60248201527f6347616d6d615769746e657373206973206e6f74206f6e206375727665000000604482015290519081900360640190fd5b610e7082610ff5565b610ec1576040805162461bcd60e51b815260206004820152601c60248201527f73486173685769746e657373206973206e6f74206f6e20637572766500000000604482015290519081900360640190fd5b610ecd878a888761101f565b610f1e576040805162461bcd60e51b815260206004820152601a60248201527f6164647228632a706b2b732a6729e289a05f755769746e657373000000000000604482015290519081900360640190fd5b610f26611896565b610f308a8761114d565b9050610f3a611896565b610f49898b878b8689896111f0565b90506000610f5a838d8d8a866112fb565b9050808a14610fa0576040805162461bcd60e51b815260206004820152600d60248201526c34b73b30b634b210383937b7b360991b604482015290519081900360640190fd5b505050505050505050505050565b60408051602080820196909652808201949094526001600160a01b039290921660608401526080808401919091528151808403909101815260a09092019052805191012090565b60208101516000906401000003d0199080096110188360005b6020020151611404565b1492915050565b60006001600160a01b03821661106a576040805162461bcd60e51b815260206004820152600b60248201526a626164207769746e65737360a81b604482015290519081900360640190fd5b60208401516000906001161561108157601c611084565b601b5b9050600070014551231950b75fc4402da1732fc9bebe1985876000602002015109865170014551231950b75fc4402da1732fc9bebe1991820392506000919089098751604080516000808252602082810180855288905260ff8916838501526060830194909452608082018590529151939450909260019260a0808401939192601f1981019281900390910190855afa158015611125573d6000803e3d6000fd5b5050604051601f1901516001600160a01b039081169088161495505050505050949350505050565b611155611896565b6111b3600184846040516020018084815260200183600260200280838360005b8381101561118d578181015183820152602001611175565b505050509050018281526020019350505050604051602081830303815290604052611428565b90505b6111bf81610ff5565b61099c5780516040805160208181019390935281518082039093018352810190526111e990611428565b90506111b6565b6111f8611896565b825186516401000003d01991900306611258576040805162461bcd60e51b815260206004820152601e60248201527f706f696e747320696e2073756d206d7573742062652064697374696e63740000604482015290519081900360640190fd5b611263878988611476565b61129e5760405162461bcd60e51b81526004018080602001828103825260218152602001806119396021913960400191505060405180910390fd5b6112a9848685611476565b6112e45760405162461bcd60e51b815260040180806020018281038252602281526020018061195a6022913960400191505060405180910390fd5b6112ef868484611596565b98975050505050505050565b6000600286868685876040516020018087815260200186600260200280838360005b8381101561133557818101518382015260200161131d565b5050505090500185600260200280838360005b83811015611360578181015183820152602001611348565b5050505090500184600260200280838360005b8381101561138b578181015183820152602001611373565b5050505090500183600260200280838360005b838110156113b657818101518382015260200161139e565b50505050905001826001600160a01b03166001600160a01b031660601b815260140196505050505050506040516020818303038152906040528051906020012060001c905095945050505050565b6000806401000003d01980848509840990506401000003d019600782089392505050565b611430611896565b6114398261165c565b815261144e61144982600061100e565b611697565b6020820181905260029006600114156105cd576020810180516401000003d019039052919050565b60008261148257600080fd5b835160208501516000906001161561149b57601c61149e565b601b5b9050600070014551231950b75fc4402da1732fc9bebe19838709604080516000808252602080830180855282905260ff871683850152606083018890526080830185905292519394509260019260a0808401939192601f1981019281900390910190855afa158015611514573d6000803e3d6000fd5b5050506020604051035190506000866040516020018082600260200280838360005b8381101561154e578181015183820152602001611536565b505050509050019150506040516020818303038152906040528051906020012060001c9050806001600160a01b0316826001600160a01b031614955050505050509392505050565b61159e611896565b8351602080860151855191860151600093849384936115bf939091906116ad565b919450925090506401000003d019858209600114611624576040805162461bcd60e51b815260206004820152601960248201527f696e765a206d75737420626520696e7665727365206f66207a00000000000000604482015290519081900360640190fd5b60405180604001604052806401000003d0198061163d57fe5b87860981526020016401000003d0198785099052979650505050505050565b805160208201205b6401000003d01981106105cd57604080516020808201939093528151808203840181529082019091528051910120611664565b600061099c8263400000f4600160fe1b0361178d565b60008080600180826401000003d019896401000003d019038808905060006401000003d0198b6401000003d019038a08905060006116ed83838585611829565b90985090506116fe88828e8861184d565b909850905061170f88828c8761184d565b909850905060006117228d878b8561184d565b909850905061173388828686611829565b909850905061174488828e8961184d565b9098509050818114611779576401000003d019818a0998506401000003d01982890997506401000003d019818309965061177d565b8196505b5050505050509450945094915050565b6000806117986118fc565b6020808252818101819052604082015260608101859052608081018490526401000003d01960a08201526117ca61191a565b60208160c0846005600019fa92508261181f576040805162461bcd60e51b81526020600482015260126024820152716269674d6f64457870206661696c7572652160701b604482015290519081900360640190fd5b5195945050505050565b6000806401000003d0198487096401000003d0198487099097909650945050505050565b600080806401000003d019878509905060006401000003d01987876401000003d019030990506401000003d0198183086401000003d01986890990999098509650505050505050565b60405180604001604052806002906020820280388339509192915050565b604051806060016040528060006001600160a01b0316815260200160008152602001600081525090565b60405180606001604052806003906020820280388339509192915050565b6040518060c001604052806006906020820280388339509192915050565b6040518060200160405280600190602082028038833950919291505056fe4669727374206d756c7469706c69636174696f6e20636865636b206661696c65645365636f6e64206d756c7469706c69636174696f6e20636865636b206661696c6564a26469706673582212201032d19b3ab072e8659465a821d803c716232be6a59db71919878764c1e9fe1764736f6c63430006020033"

// DeployVRFCoordinator deploys a new Ethereum contract, binding an instance of VRFCoordinator to it.
func DeployVRFCoordinator(auth *bind.TransactOpts, backend bind.ContractBackend, _link common.Address) (common.Address, *types.Transaction, *VRFCoordinator, error) {
	parsed, err := abi.JSON(strings.NewReader(VRFCoordinatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VRFCoordinatorBin), backend, _link)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFCoordinator{VRFCoordinatorCaller: VRFCoordinatorCaller{contract: contract}, VRFCoordinatorTransactor: VRFCoordinatorTransactor{contract: contract}, VRFCoordinatorFilterer: VRFCoordinatorFilterer{contract: contract}}, nil
}

// VRFCoordinator is an auto generated Go binding around an Ethereum contract.
type VRFCoordinator struct {
	VRFCoordinatorCaller     // Read-only binding to the contract
	VRFCoordinatorTransactor // Write-only binding to the contract
	VRFCoordinatorFilterer   // Log filterer for contract events
}

// VRFCoordinatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type VRFCoordinatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFCoordinatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VRFCoordinatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFCoordinatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VRFCoordinatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFCoordinatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VRFCoordinatorSession struct {
	Contract     *VRFCoordinator   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VRFCoordinatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VRFCoordinatorCallerSession struct {
	Contract *VRFCoordinatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// VRFCoordinatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VRFCoordinatorTransactorSession struct {
	Contract     *VRFCoordinatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VRFCoordinatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type VRFCoordinatorRaw struct {
	Contract *VRFCoordinator // Generic contract binding to access the raw methods on
}

// VRFCoordinatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VRFCoordinatorCallerRaw struct {
	Contract *VRFCoordinatorCaller // Generic read-only contract binding to access the raw methods on
}

// VRFCoordinatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VRFCoordinatorTransactorRaw struct {
	Contract *VRFCoordinatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVRFCoordinator creates a new instance of VRFCoordinator, bound to a specific deployed contract.
func NewVRFCoordinator(address common.Address, backend bind.ContractBackend) (*VRFCoordinator, error) {
	contract, err := bindVRFCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinator{VRFCoordinatorCaller: VRFCoordinatorCaller{contract: contract}, VRFCoordinatorTransactor: VRFCoordinatorTransactor{contract: contract}, VRFCoordinatorFilterer: VRFCoordinatorFilterer{contract: contract}}, nil
}

// NewVRFCoordinatorCaller creates a new read-only instance of VRFCoordinator, bound to a specific deployed contract.
func NewVRFCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*VRFCoordinatorCaller, error) {
	contract, err := bindVRFCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorCaller{contract: contract}, nil
}

// NewVRFCoordinatorTransactor creates a new write-only instance of VRFCoordinator, bound to a specific deployed contract.
func NewVRFCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFCoordinatorTransactor, error) {
	contract, err := bindVRFCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorTransactor{contract: contract}, nil
}

// NewVRFCoordinatorFilterer creates a new log filterer instance of VRFCoordinator, bound to a specific deployed contract.
func NewVRFCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFCoordinatorFilterer, error) {
	contract, err := bindVRFCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorFilterer{contract: contract}, nil
}

// bindVRFCoordinator binds a generic wrapper to an already deployed contract.
func bindVRFCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VRFCoordinatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRFCoordinator *VRFCoordinatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VRFCoordinator.Contract.VRFCoordinatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRFCoordinator *VRFCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.VRFCoordinatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRFCoordinator *VRFCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.VRFCoordinatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRFCoordinator *VRFCoordinatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VRFCoordinator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRFCoordinator *VRFCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRFCoordinator *VRFCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.contract.Transact(opts, method, params...)
}

// Callbacks is a free data retrieval call binding the contract method 0x21f36509.
//
// Solidity: function callbacks(bytes32 ) constant returns(address callbackContract, uint256 randomnessFee, uint256 seed)
func (_VRFCoordinator *VRFCoordinatorCaller) Callbacks(opts *bind.CallOpts, arg0 [32]byte) (struct {
	CallbackContract common.Address
	RandomnessFee    *big.Int
	Seed             *big.Int
}, error) {
	ret := new(struct {
		CallbackContract common.Address
		RandomnessFee    *big.Int
		Seed             *big.Int
	})
	out := ret
	err := _VRFCoordinator.contract.Call(opts, out, "callbacks", arg0)
	return *ret, err
}

// Callbacks is a free data retrieval call binding the contract method 0x21f36509.
//
// Solidity: function callbacks(bytes32 ) constant returns(address callbackContract, uint256 randomnessFee, uint256 seed)
func (_VRFCoordinator *VRFCoordinatorSession) Callbacks(arg0 [32]byte) (struct {
	CallbackContract common.Address
	RandomnessFee    *big.Int
	Seed             *big.Int
}, error) {
	return _VRFCoordinator.Contract.Callbacks(&_VRFCoordinator.CallOpts, arg0)
}

// Callbacks is a free data retrieval call binding the contract method 0x21f36509.
//
// Solidity: function callbacks(bytes32 ) constant returns(address callbackContract, uint256 randomnessFee, uint256 seed)
func (_VRFCoordinator *VRFCoordinatorCallerSession) Callbacks(arg0 [32]byte) (struct {
	CallbackContract common.Address
	RandomnessFee    *big.Int
	Seed             *big.Int
}, error) {
	return _VRFCoordinator.Contract.Callbacks(&_VRFCoordinator.CallOpts, arg0)
}

// HashOfKey is a free data retrieval call binding the contract method 0xcaf70c4a.
//
// Solidity: function hashOfKey(uint256[2] _publicKey) constant returns(bytes32)
func (_VRFCoordinator *VRFCoordinatorCaller) HashOfKey(opts *bind.CallOpts, _publicKey [2]*big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _VRFCoordinator.contract.Call(opts, out, "hashOfKey", _publicKey)
	return *ret0, err
}

// HashOfKey is a free data retrieval call binding the contract method 0xcaf70c4a.
//
// Solidity: function hashOfKey(uint256[2] _publicKey) constant returns(bytes32)
func (_VRFCoordinator *VRFCoordinatorSession) HashOfKey(_publicKey [2]*big.Int) ([32]byte, error) {
	return _VRFCoordinator.Contract.HashOfKey(&_VRFCoordinator.CallOpts, _publicKey)
}

// HashOfKey is a free data retrieval call binding the contract method 0xcaf70c4a.
//
// Solidity: function hashOfKey(uint256[2] _publicKey) constant returns(bytes32)
func (_VRFCoordinator *VRFCoordinatorCallerSession) HashOfKey(_publicKey [2]*big.Int) ([32]byte, error) {
	return _VRFCoordinator.Contract.HashOfKey(&_VRFCoordinator.CallOpts, _publicKey)
}

// ServiceAgreements is a free data retrieval call binding the contract method 0x75d35070.
//
// Solidity: function serviceAgreements(bytes32 ) constant returns(address vRFOracle, bytes32 jobID, uint256 fee)
func (_VRFCoordinator *VRFCoordinatorCaller) ServiceAgreements(opts *bind.CallOpts, arg0 [32]byte) (struct {
	VRFOracle common.Address
	JobID     [32]byte
	Fee       *big.Int
}, error) {
	ret := new(struct {
		VRFOracle common.Address
		JobID     [32]byte
		Fee       *big.Int
	})
	out := ret
	err := _VRFCoordinator.contract.Call(opts, out, "serviceAgreements", arg0)
	return *ret, err
}

// ServiceAgreements is a free data retrieval call binding the contract method 0x75d35070.
//
// Solidity: function serviceAgreements(bytes32 ) constant returns(address vRFOracle, bytes32 jobID, uint256 fee)
func (_VRFCoordinator *VRFCoordinatorSession) ServiceAgreements(arg0 [32]byte) (struct {
	VRFOracle common.Address
	JobID     [32]byte
	Fee       *big.Int
}, error) {
	return _VRFCoordinator.Contract.ServiceAgreements(&_VRFCoordinator.CallOpts, arg0)
}

// ServiceAgreements is a free data retrieval call binding the contract method 0x75d35070.
//
// Solidity: function serviceAgreements(bytes32 ) constant returns(address vRFOracle, bytes32 jobID, uint256 fee)
func (_VRFCoordinator *VRFCoordinatorCallerSession) ServiceAgreements(arg0 [32]byte) (struct {
	VRFOracle common.Address
	JobID     [32]byte
	Fee       *big.Int
}, error) {
	return _VRFCoordinator.Contract.ServiceAgreements(&_VRFCoordinator.CallOpts, arg0)
}

// WithdrawableTokens is a free data retrieval call binding the contract method 0x006f6ad0.
//
// Solidity: function withdrawableTokens(address ) constant returns(uint256)
func (_VRFCoordinator *VRFCoordinatorCaller) WithdrawableTokens(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VRFCoordinator.contract.Call(opts, out, "withdrawableTokens", arg0)
	return *ret0, err
}

// WithdrawableTokens is a free data retrieval call binding the contract method 0x006f6ad0.
//
// Solidity: function withdrawableTokens(address ) constant returns(uint256)
func (_VRFCoordinator *VRFCoordinatorSession) WithdrawableTokens(arg0 common.Address) (*big.Int, error) {
	return _VRFCoordinator.Contract.WithdrawableTokens(&_VRFCoordinator.CallOpts, arg0)
}

// WithdrawableTokens is a free data retrieval call binding the contract method 0x006f6ad0.
//
// Solidity: function withdrawableTokens(address ) constant returns(uint256)
func (_VRFCoordinator *VRFCoordinatorCallerSession) WithdrawableTokens(arg0 common.Address) (*big.Int, error) {
	return _VRFCoordinator.Contract.WithdrawableTokens(&_VRFCoordinator.CallOpts, arg0)
}

// FulfillRandomnessRequest is a paid mutator transaction binding the contract method 0x5e1c1059.
//
// Solidity: function fulfillRandomnessRequest(bytes _proof) returns(bool)
func (_VRFCoordinator *VRFCoordinatorTransactor) FulfillRandomnessRequest(opts *bind.TransactOpts, _proof []byte) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "fulfillRandomnessRequest", _proof)
}

// FulfillRandomnessRequest is a paid mutator transaction binding the contract method 0x5e1c1059.
//
// Solidity: function fulfillRandomnessRequest(bytes _proof) returns(bool)
func (_VRFCoordinator *VRFCoordinatorSession) FulfillRandomnessRequest(_proof []byte) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.FulfillRandomnessRequest(&_VRFCoordinator.TransactOpts, _proof)
}

// FulfillRandomnessRequest is a paid mutator transaction binding the contract method 0x5e1c1059.
//
// Solidity: function fulfillRandomnessRequest(bytes _proof) returns(bool)
func (_VRFCoordinator *VRFCoordinatorTransactorSession) FulfillRandomnessRequest(_proof []byte) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.FulfillRandomnessRequest(&_VRFCoordinator.TransactOpts, _proof)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address _sender, uint256 _fee, bytes _data) returns()
func (_VRFCoordinator *VRFCoordinatorTransactor) OnTokenTransfer(opts *bind.TransactOpts, _sender common.Address, _fee *big.Int, _data []byte) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "onTokenTransfer", _sender, _fee, _data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address _sender, uint256 _fee, bytes _data) returns()
func (_VRFCoordinator *VRFCoordinatorSession) OnTokenTransfer(_sender common.Address, _fee *big.Int, _data []byte) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.OnTokenTransfer(&_VRFCoordinator.TransactOpts, _sender, _fee, _data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address _sender, uint256 _fee, bytes _data) returns()
func (_VRFCoordinator *VRFCoordinatorTransactorSession) OnTokenTransfer(_sender common.Address, _fee *big.Int, _data []byte) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.OnTokenTransfer(&_VRFCoordinator.TransactOpts, _sender, _fee, _data)
}

// RegisterProvingKey is a paid mutator transaction binding the contract method 0xd8340209.
//
// Solidity: function registerProvingKey(uint256 _fee, address _oracle, uint256[2] _publicProvingKey, bytes32 _jobID) returns()
func (_VRFCoordinator *VRFCoordinatorTransactor) RegisterProvingKey(opts *bind.TransactOpts, _fee *big.Int, _oracle common.Address, _publicProvingKey [2]*big.Int, _jobID [32]byte) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "registerProvingKey", _fee, _oracle, _publicProvingKey, _jobID)
}

// RegisterProvingKey is a paid mutator transaction binding the contract method 0xd8340209.
//
// Solidity: function registerProvingKey(uint256 _fee, address _oracle, uint256[2] _publicProvingKey, bytes32 _jobID) returns()
func (_VRFCoordinator *VRFCoordinatorSession) RegisterProvingKey(_fee *big.Int, _oracle common.Address, _publicProvingKey [2]*big.Int, _jobID [32]byte) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.RegisterProvingKey(&_VRFCoordinator.TransactOpts, _fee, _oracle, _publicProvingKey, _jobID)
}

// RegisterProvingKey is a paid mutator transaction binding the contract method 0xd8340209.
//
// Solidity: function registerProvingKey(uint256 _fee, address _oracle, uint256[2] _publicProvingKey, bytes32 _jobID) returns()
func (_VRFCoordinator *VRFCoordinatorTransactorSession) RegisterProvingKey(_fee *big.Int, _oracle common.Address, _publicProvingKey [2]*big.Int, _jobID [32]byte) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.RegisterProvingKey(&_VRFCoordinator.TransactOpts, _fee, _oracle, _publicProvingKey, _jobID)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _recipient, uint256 _amount) returns()
func (_VRFCoordinator *VRFCoordinatorTransactor) Withdraw(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "withdraw", _recipient, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _recipient, uint256 _amount) returns()
func (_VRFCoordinator *VRFCoordinatorSession) Withdraw(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.Withdraw(&_VRFCoordinator.TransactOpts, _recipient, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _recipient, uint256 _amount) returns()
func (_VRFCoordinator *VRFCoordinatorTransactorSession) Withdraw(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.Withdraw(&_VRFCoordinator.TransactOpts, _recipient, _amount)
}

// VRFCoordinatorNewServiceAgreementIterator is returned from FilterNewServiceAgreement and is used to iterate over the raw logs and unpacked data for NewServiceAgreement events raised by the VRFCoordinator contract.
type VRFCoordinatorNewServiceAgreementIterator struct {
	Event *VRFCoordinatorNewServiceAgreement // Event containing the contract specifics and raw log

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
func (it *VRFCoordinatorNewServiceAgreementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorNewServiceAgreement)
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
		it.Event = new(VRFCoordinatorNewServiceAgreement)
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
func (it *VRFCoordinatorNewServiceAgreementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VRFCoordinatorNewServiceAgreementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VRFCoordinatorNewServiceAgreement represents a NewServiceAgreement event raised by the VRFCoordinator contract.
type VRFCoordinatorNewServiceAgreement struct {
	KeyHash [32]byte
	Fee     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewServiceAgreement is a free log retrieval operation binding the contract event 0xae189157e0628c1e62315e9179156e1ea10e90e9c15060002f7021e907dc2cfe.
//
// Solidity: event NewServiceAgreement(bytes32 keyHash, uint256 fee)
func (_VRFCoordinator *VRFCoordinatorFilterer) FilterNewServiceAgreement(opts *bind.FilterOpts) (*VRFCoordinatorNewServiceAgreementIterator, error) {

	logs, sub, err := _VRFCoordinator.contract.FilterLogs(opts, "NewServiceAgreement")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorNewServiceAgreementIterator{contract: _VRFCoordinator.contract, event: "NewServiceAgreement", logs: logs, sub: sub}, nil
}

// WatchNewServiceAgreement is a free log subscription operation binding the contract event 0xae189157e0628c1e62315e9179156e1ea10e90e9c15060002f7021e907dc2cfe.
//
// Solidity: event NewServiceAgreement(bytes32 keyHash, uint256 fee)
func (_VRFCoordinator *VRFCoordinatorFilterer) WatchNewServiceAgreement(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorNewServiceAgreement) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinator.contract.WatchLogs(opts, "NewServiceAgreement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VRFCoordinatorNewServiceAgreement)
				if err := _VRFCoordinator.contract.UnpackLog(event, "NewServiceAgreement", log); err != nil {
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

// ParseNewServiceAgreement is a log parse operation binding the contract event 0xae189157e0628c1e62315e9179156e1ea10e90e9c15060002f7021e907dc2cfe.
//
// Solidity: event NewServiceAgreement(bytes32 keyHash, uint256 fee)
func (_VRFCoordinator *VRFCoordinatorFilterer) ParseNewServiceAgreement(log types.Log) (*VRFCoordinatorNewServiceAgreement, error) {
	event := new(VRFCoordinatorNewServiceAgreement)
	if err := _VRFCoordinator.contract.UnpackLog(event, "NewServiceAgreement", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VRFCoordinatorRandomnessRequestIterator is returned from FilterRandomnessRequest and is used to iterate over the raw logs and unpacked data for RandomnessRequest events raised by the VRFCoordinator contract.
type VRFCoordinatorRandomnessRequestIterator struct {
	Event *VRFCoordinatorRandomnessRequest // Event containing the contract specifics and raw log

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
func (it *VRFCoordinatorRandomnessRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorRandomnessRequest)
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
		it.Event = new(VRFCoordinatorRandomnessRequest)
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
func (it *VRFCoordinatorRandomnessRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VRFCoordinatorRandomnessRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VRFCoordinatorRandomnessRequest represents a RandomnessRequest event raised by the VRFCoordinator contract.
type VRFCoordinatorRandomnessRequest struct {
	KeyHash [32]byte
	Seed    *big.Int
	JobID   [32]byte
	Sender  common.Address
	Fee     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRandomnessRequest is a free log retrieval operation binding the contract event 0xd241d78a52145a5d1d1ff002e32ec15cdc395631bcee66246650c2429dfaccaa.
//
// Solidity: event RandomnessRequest(bytes32 keyHash, uint256 seed, bytes32 indexed jobID, address sender, uint256 fee)
func (_VRFCoordinator *VRFCoordinatorFilterer) FilterRandomnessRequest(opts *bind.FilterOpts, jobID [][32]byte) (*VRFCoordinatorRandomnessRequestIterator, error) {

	var jobIDRule []interface{}
	for _, jobIDItem := range jobID {
		jobIDRule = append(jobIDRule, jobIDItem)
	}

	logs, sub, err := _VRFCoordinator.contract.FilterLogs(opts, "RandomnessRequest", jobIDRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorRandomnessRequestIterator{contract: _VRFCoordinator.contract, event: "RandomnessRequest", logs: logs, sub: sub}, nil
}

// WatchRandomnessRequest is a free log subscription operation binding the contract event 0xd241d78a52145a5d1d1ff002e32ec15cdc395631bcee66246650c2429dfaccaa.
//
// Solidity: event RandomnessRequest(bytes32 keyHash, uint256 seed, bytes32 indexed jobID, address sender, uint256 fee)
func (_VRFCoordinator *VRFCoordinatorFilterer) WatchRandomnessRequest(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorRandomnessRequest, jobID [][32]byte) (event.Subscription, error) {

	var jobIDRule []interface{}
	for _, jobIDItem := range jobID {
		jobIDRule = append(jobIDRule, jobIDItem)
	}

	logs, sub, err := _VRFCoordinator.contract.WatchLogs(opts, "RandomnessRequest", jobIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VRFCoordinatorRandomnessRequest)
				if err := _VRFCoordinator.contract.UnpackLog(event, "RandomnessRequest", log); err != nil {
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

// ParseRandomnessRequest is a log parse operation binding the contract event 0xd241d78a52145a5d1d1ff002e32ec15cdc395631bcee66246650c2429dfaccaa.
//
// Solidity: event RandomnessRequest(bytes32 keyHash, uint256 seed, bytes32 indexed jobID, address sender, uint256 fee)
func (_VRFCoordinator *VRFCoordinatorFilterer) ParseRandomnessRequest(log types.Log) (*VRFCoordinatorRandomnessRequest, error) {
	event := new(VRFCoordinatorRandomnessRequest)
	if err := _VRFCoordinator.contract.UnpackLog(event, "RandomnessRequest", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VRFRequestIDBaseABI is the input ABI used to generate the binding from.
const VRFRequestIDBaseABI = "[]"

// VRFRequestIDBaseBin is the compiled bytecode used for deploying new contracts.
var VRFRequestIDBaseBin = "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea26469706673582212208bba9a435c1a7d0c09c1cfe66c3894331981a5e1d83789837919a4272efc89af64736f6c63430006020033"

// DeployVRFRequestIDBase deploys a new Ethereum contract, binding an instance of VRFRequestIDBase to it.
func DeployVRFRequestIDBase(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VRFRequestIDBase, error) {
	parsed, err := abi.JSON(strings.NewReader(VRFRequestIDBaseABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VRFRequestIDBaseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFRequestIDBase{VRFRequestIDBaseCaller: VRFRequestIDBaseCaller{contract: contract}, VRFRequestIDBaseTransactor: VRFRequestIDBaseTransactor{contract: contract}, VRFRequestIDBaseFilterer: VRFRequestIDBaseFilterer{contract: contract}}, nil
}

// VRFRequestIDBase is an auto generated Go binding around an Ethereum contract.
type VRFRequestIDBase struct {
	VRFRequestIDBaseCaller     // Read-only binding to the contract
	VRFRequestIDBaseTransactor // Write-only binding to the contract
	VRFRequestIDBaseFilterer   // Log filterer for contract events
}

// VRFRequestIDBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type VRFRequestIDBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFRequestIDBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VRFRequestIDBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFRequestIDBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VRFRequestIDBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFRequestIDBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VRFRequestIDBaseSession struct {
	Contract     *VRFRequestIDBase // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VRFRequestIDBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VRFRequestIDBaseCallerSession struct {
	Contract *VRFRequestIDBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// VRFRequestIDBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VRFRequestIDBaseTransactorSession struct {
	Contract     *VRFRequestIDBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// VRFRequestIDBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type VRFRequestIDBaseRaw struct {
	Contract *VRFRequestIDBase // Generic contract binding to access the raw methods on
}

// VRFRequestIDBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VRFRequestIDBaseCallerRaw struct {
	Contract *VRFRequestIDBaseCaller // Generic read-only contract binding to access the raw methods on
}

// VRFRequestIDBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VRFRequestIDBaseTransactorRaw struct {
	Contract *VRFRequestIDBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVRFRequestIDBase creates a new instance of VRFRequestIDBase, bound to a specific deployed contract.
func NewVRFRequestIDBase(address common.Address, backend bind.ContractBackend) (*VRFRequestIDBase, error) {
	contract, err := bindVRFRequestIDBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFRequestIDBase{VRFRequestIDBaseCaller: VRFRequestIDBaseCaller{contract: contract}, VRFRequestIDBaseTransactor: VRFRequestIDBaseTransactor{contract: contract}, VRFRequestIDBaseFilterer: VRFRequestIDBaseFilterer{contract: contract}}, nil
}

// NewVRFRequestIDBaseCaller creates a new read-only instance of VRFRequestIDBase, bound to a specific deployed contract.
func NewVRFRequestIDBaseCaller(address common.Address, caller bind.ContractCaller) (*VRFRequestIDBaseCaller, error) {
	contract, err := bindVRFRequestIDBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFRequestIDBaseCaller{contract: contract}, nil
}

// NewVRFRequestIDBaseTransactor creates a new write-only instance of VRFRequestIDBase, bound to a specific deployed contract.
func NewVRFRequestIDBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFRequestIDBaseTransactor, error) {
	contract, err := bindVRFRequestIDBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFRequestIDBaseTransactor{contract: contract}, nil
}

// NewVRFRequestIDBaseFilterer creates a new log filterer instance of VRFRequestIDBase, bound to a specific deployed contract.
func NewVRFRequestIDBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFRequestIDBaseFilterer, error) {
	contract, err := bindVRFRequestIDBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFRequestIDBaseFilterer{contract: contract}, nil
}

// bindVRFRequestIDBase binds a generic wrapper to an already deployed contract.
func bindVRFRequestIDBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VRFRequestIDBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRFRequestIDBase *VRFRequestIDBaseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VRFRequestIDBase.Contract.VRFRequestIDBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRFRequestIDBase *VRFRequestIDBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFRequestIDBase.Contract.VRFRequestIDBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRFRequestIDBase *VRFRequestIDBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFRequestIDBase.Contract.VRFRequestIDBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRFRequestIDBase *VRFRequestIDBaseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VRFRequestIDBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRFRequestIDBase *VRFRequestIDBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFRequestIDBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRFRequestIDBase *VRFRequestIDBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFRequestIDBase.Contract.contract.Transact(opts, method, params...)
}

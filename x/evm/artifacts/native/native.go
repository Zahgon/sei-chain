// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package native

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

// NativeMetaData contains all meta data concerning the Native contract.
var NativeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BankPrecompile\",\"outputs\":[{\"internalType\":\"contractIBank\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ddecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denom\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nname\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ssymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NativeABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeMetaData.ABI instead.
var NativeABI = NativeMetaData.ABI

// Native is an auto generated Go binding around an Ethereum contract.
type Native struct {
	NativeCaller     // Read-only binding to the contract
	NativeTransactor // Write-only binding to the contract
	NativeFilterer   // Log filterer for contract events
}

// NativeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeSession struct {
	Contract     *Native           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NativeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeCallerSession struct {
	Contract *NativeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NativeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeTransactorSession struct {
	Contract     *NativeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NativeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeRaw struct {
	Contract *Native // Generic contract binding to access the raw methods on
}

// NativeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeCallerRaw struct {
	Contract *NativeCaller // Generic read-only contract binding to access the raw methods on
}

// NativeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeTransactorRaw struct {
	Contract *NativeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNative creates a new instance of Native, bound to a specific deployed contract.
func NewNative(address common.Address, backend bind.ContractBackend) (*Native, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewNativeCaller creates a new read-only instance of Native, bound to a specific deployed contract.
func NewNativeCaller(address common.Address, caller bind.ContractCaller) (*NativeCaller, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewNativeTransactor creates a new write-only instance of Native, bound to a specific deployed contract.
func NewNativeTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTransactor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewNativeFilterer creates a new log filterer instance of Native, bound to a specific deployed contract.
func NewNativeFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeFilterer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// bindNative binds a generic wrapper to an already deployed contract.
func bindNative(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Native *NativeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Native *NativeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transact invokes the (paid) contract method with params as input values.
func (_Native *NativeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Native *NativeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Native *NativeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transact invokes the (paid) contract method with params as input values.
func (_Native *NativeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BankPrecompile is a free data retrieval call binding the contract method 0x566732c1.
//
// Solidity: function BankPrecompile() view returns(address)
func (_Native *NativeCaller) BankPrecompile(opts *bind.CallOpts) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// BankPrecompile is a free data retrieval call binding the contract method 0x566732c1.
//
// Solidity: function BankPrecompile() view returns(address)
func (_Native *NativeSession) BankPrecompile() (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// BankPrecompile is a free data retrieval call binding the contract method 0x566732c1.
//
// Solidity: function BankPrecompile() view returns(address)
func (_Native *NativeCallerSession) BankPrecompile() (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Native *NativeCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Native *NativeSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Native *NativeCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Native *NativeCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Native *NativeSession) BalanceOf(account common.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Native *NativeCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ddecimals is a free data retrieval call binding the contract method 0xa8ad11e4.
//
// Solidity: function ddecimals() view returns(uint8)
func (_Native *NativeCaller) Ddecimals(opts *bind.CallOpts) (uint8, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Ddecimals is a free data retrieval call binding the contract method 0xa8ad11e4.
//
// Solidity: function ddecimals() view returns(uint8)
func (_Native *NativeSession) Ddecimals() (uint8, error) { _ = "STUB: not implemented"; return 0, nil }

// Ddecimals is a free data retrieval call binding the contract method 0xa8ad11e4.
//
// Solidity: function ddecimals() view returns(uint8)
func (_Native *NativeCallerSession) Ddecimals() (uint8, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Native *NativeCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Native *NativeSession) Decimals() (uint8, error) { _ = "STUB: not implemented"; return 0, nil }

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Native *NativeCallerSession) Decimals() (uint8, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Denom is a free data retrieval call binding the contract method 0xc370b042.
//
// Solidity: function denom() view returns(string)
func (_Native *NativeCaller) Denom(opts *bind.CallOpts) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Denom is a free data retrieval call binding the contract method 0xc370b042.
//
// Solidity: function denom() view returns(string)
func (_Native *NativeSession) Denom() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Denom is a free data retrieval call binding the contract method 0xc370b042.
//
// Solidity: function denom() view returns(string)
func (_Native *NativeCallerSession) Denom() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Native *NativeCaller) Name(opts *bind.CallOpts) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Native *NativeSession) Name() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Native *NativeCallerSession) Name() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Nname is a free data retrieval call binding the contract method 0x8a0989f5.
//
// Solidity: function nname() view returns(string)
func (_Native *NativeCaller) Nname(opts *bind.CallOpts) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Nname is a free data retrieval call binding the contract method 0x8a0989f5.
//
// Solidity: function nname() view returns(string)
func (_Native *NativeSession) Nname() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Nname is a free data retrieval call binding the contract method 0x8a0989f5.
//
// Solidity: function nname() view returns(string)
func (_Native *NativeCallerSession) Nname() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Ssymbol is a free data retrieval call binding the contract method 0x9e10aa24.
//
// Solidity: function ssymbol() view returns(string)
func (_Native *NativeCaller) Ssymbol(opts *bind.CallOpts) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Ssymbol is a free data retrieval call binding the contract method 0x9e10aa24.
//
// Solidity: function ssymbol() view returns(string)
func (_Native *NativeSession) Ssymbol() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Ssymbol is a free data retrieval call binding the contract method 0x9e10aa24.
//
// Solidity: function ssymbol() view returns(string)
func (_Native *NativeCallerSession) Ssymbol() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Native *NativeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Native *NativeSession) Symbol() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Native *NativeCallerSession) Symbol() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Native *NativeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Native *NativeSession) TotalSupply() (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Native *NativeCallerSession) TotalSupply() (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Native *NativeTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Native *NativeSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Native *NativeTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Native *NativeTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Native *NativeSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Native *NativeTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Native *NativeTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Native *NativeSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Native *NativeTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NativeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Native contract.
type NativeApprovalIterator struct {
	Event *NativeApproval // Event containing the contract specifics and raw log

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
func (it *NativeApprovalIterator) Next() bool {
	_ = "STUB: not implemented"
	// If the iterator failed, stop iterating
	return false
}

// If the iterator completed, deliver directly whatever's available

// Iterator still in progress, wait for either a data or an error event

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeApprovalIterator) Error() error {
	_ = "STUB: not implemented"

	// Close terminates the iteration process, releasing any pending underlying
	// resources.
	return nil
}

func (it *NativeApprovalIterator) Close() error { _ = "STUB: not implemented"; return nil }

// NativeApproval represents a Approval event raised by the Native contract.
type NativeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Native *NativeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*NativeApprovalIterator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Native *NativeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NativeApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {
	_ = "STUB: not implemented"
	return *new(event.Subscription), nil
}

// New log arrived, parse the event and forward to the user

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Native *NativeFilterer) ParseApproval(log types.Log) (*NativeApproval, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NativeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Native contract.
type NativeTransferIterator struct {
	Event *NativeTransfer // Event containing the contract specifics and raw log

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
func (it *NativeTransferIterator) Next() bool {
	_ = "STUB: not implemented"
	// If the iterator failed, stop iterating
	return false
}

// If the iterator completed, deliver directly whatever's available

// Iterator still in progress, wait for either a data or an error event

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTransferIterator) Error() error {
	_ = "STUB: not implemented"

	// Close terminates the iteration process, releasing any pending underlying
	// resources.
	return nil
}

func (it *NativeTransferIterator) Close() error { _ = "STUB: not implemented"; return nil }

// NativeTransfer represents a Transfer event raised by the Native contract.
type NativeTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Native *NativeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NativeTransferIterator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Native *NativeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NativeTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {
	_ = "STUB: not implemented"
	return *new(event.Subscription), nil
}

// New log arrived, parse the event and forward to the user

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Native *NativeFilterer) ParseTransfer(log types.Log) (*NativeTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

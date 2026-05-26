// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package echo

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

// EchoMetaData contains all meta data concerning the Echo contract.
var EchoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"echo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"setTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"timestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// EchoABI is the input ABI used to generate the binding from.
// Deprecated: Use EchoMetaData.ABI instead.
var EchoABI = EchoMetaData.ABI

// Echo is an auto generated Go binding around an Ethereum contract.
type Echo struct {
	EchoCaller     // Read-only binding to the contract
	EchoTransactor // Write-only binding to the contract
	EchoFilterer   // Log filterer for contract events
}

// EchoCaller is an auto generated read-only Go binding around an Ethereum contract.
type EchoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EchoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EchoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EchoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EchoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EchoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EchoSession struct {
	Contract     *Echo             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EchoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EchoCallerSession struct {
	Contract *EchoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EchoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EchoTransactorSession struct {
	Contract     *EchoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EchoRaw is an auto generated low-level Go binding around an Ethereum contract.
type EchoRaw struct {
	Contract *Echo // Generic contract binding to access the raw methods on
}

// EchoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EchoCallerRaw struct {
	Contract *EchoCaller // Generic read-only contract binding to access the raw methods on
}

// EchoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EchoTransactorRaw struct {
	Contract *EchoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEcho creates a new instance of Echo, bound to a specific deployed contract.
func NewEcho(address common.Address, backend bind.ContractBackend) (*Echo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewEchoCaller creates a new read-only instance of Echo, bound to a specific deployed contract.
func NewEchoCaller(address common.Address, caller bind.ContractCaller) (*EchoCaller, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewEchoTransactor creates a new write-only instance of Echo, bound to a specific deployed contract.
func NewEchoTransactor(address common.Address, transactor bind.ContractTransactor) (*EchoTransactor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewEchoFilterer creates a new log filterer instance of Echo, bound to a specific deployed contract.
func NewEchoFilterer(address common.Address, filterer bind.ContractFilterer) (*EchoFilterer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// bindEcho binds a generic wrapper to an already deployed contract.
func bindEcho(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Echo *EchoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Echo *EchoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transact invokes the (paid) contract method with params as input values.
func (_Echo *EchoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Echo *EchoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Echo *EchoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transact invokes the (paid) contract method with params as input values.
func (_Echo *EchoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Echo is a free data retrieval call binding the contract method 0x6279e43c.
//
// Solidity: function echo(uint256 value) pure returns(uint256)
func (_Echo *EchoCaller) Echo(opts *bind.CallOpts, value *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Echo is a free data retrieval call binding the contract method 0x6279e43c.
//
// Solidity: function echo(uint256 value) pure returns(uint256)
func (_Echo *EchoSession) Echo(value *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Echo is a free data retrieval call binding the contract method 0x6279e43c.
//
// Solidity: function echo(uint256 value) pure returns(uint256)
func (_Echo *EchoCallerSession) Echo(value *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Timestamps is a free data retrieval call binding the contract method 0x8bc33af3.
//
// Solidity: function timestamps(uint256 ) view returns(uint256)
func (_Echo *EchoCaller) Timestamps(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Timestamps is a free data retrieval call binding the contract method 0x8bc33af3.
//
// Solidity: function timestamps(uint256 ) view returns(uint256)
func (_Echo *EchoSession) Timestamps(arg0 *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Timestamps is a free data retrieval call binding the contract method 0x8bc33af3.
//
// Solidity: function timestamps(uint256 ) view returns(uint256)
func (_Echo *EchoCallerSession) Timestamps(arg0 *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetTime is a paid mutator transaction binding the contract method 0x3beb26c4.
//
// Solidity: function setTime(uint256 epoch) returns()
func (_Echo *EchoTransactor) SetTime(opts *bind.TransactOpts, epoch *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetTime is a paid mutator transaction binding the contract method 0x3beb26c4.
//
// Solidity: function setTime(uint256 epoch) returns()
func (_Echo *EchoSession) SetTime(epoch *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetTime is a paid mutator transaction binding the contract method 0x3beb26c4.
//
// Solidity: function setTime(uint256 epoch) returns()
func (_Echo *EchoTransactorSession) SetTime(epoch *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cw1155

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

// Cw1155MetaData contains all meta data concerning the Cw1155 contract.
var Cw1155MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Cw1155Address_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC1155InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valuesLength\",\"type\":\"uint256\"}],\"name\":\"ERC1155InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC1155MissingApprovalForAll\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"ERC2981InvalidDefaultRoyalty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC2981InvalidDefaultRoyaltyReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"ERC2981InvalidTokenRoyalty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC2981InvalidTokenRoyaltyReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"}],\"name\":\"NotImplemented\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"}],\"name\":\"NotImplementedOnCosmwasmContract\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AddrPrecompile\",\"outputs\":[{\"internalType\":\"contractIAddr\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Cw1155Address\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"JsonPrecompile\",\"outputs\":[{\"internalType\":\"contractIJson\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WasmdPrecompile\",\"outputs\":[{\"internalType\":\"contractIWasmd\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"burnBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Cw1155ABI is the input ABI used to generate the binding from.
// Deprecated: Use Cw1155MetaData.ABI instead.
var Cw1155ABI = Cw1155MetaData.ABI

// Cw1155 is an auto generated Go binding around an Ethereum contract.
type Cw1155 struct {
	Cw1155Caller     // Read-only binding to the contract
	Cw1155Transactor // Write-only binding to the contract
	Cw1155Filterer   // Log filterer for contract events
}

// Cw1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type Cw1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Cw1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Cw1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Cw1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Cw1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Cw1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Cw1155Session struct {
	Contract     *Cw1155           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Cw1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Cw1155CallerSession struct {
	Contract *Cw1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Cw1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Cw1155TransactorSession struct {
	Contract     *Cw1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Cw1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type Cw1155Raw struct {
	Contract *Cw1155 // Generic contract binding to access the raw methods on
}

// Cw1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Cw1155CallerRaw struct {
	Contract *Cw1155Caller // Generic read-only contract binding to access the raw methods on
}

// Cw1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Cw1155TransactorRaw struct {
	Contract *Cw1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCw1155 creates a new instance of Cw1155, bound to a specific deployed contract.
func NewCw1155(address common.Address, backend bind.ContractBackend) (*Cw1155, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewCw1155Caller creates a new read-only instance of Cw1155, bound to a specific deployed contract.
func NewCw1155Caller(address common.Address, caller bind.ContractCaller) (*Cw1155Caller, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewCw1155Transactor creates a new write-only instance of Cw1155, bound to a specific deployed contract.
func NewCw1155Transactor(address common.Address, transactor bind.ContractTransactor) (*Cw1155Transactor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewCw1155Filterer creates a new log filterer instance of Cw1155, bound to a specific deployed contract.
func NewCw1155Filterer(address common.Address, filterer bind.ContractFilterer) (*Cw1155Filterer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// bindCw1155 binds a generic wrapper to an already deployed contract.
func bindCw1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cw1155 *Cw1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cw1155 *Cw1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cw1155 *Cw1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cw1155 *Cw1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cw1155 *Cw1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cw1155 *Cw1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddrPrecompile is a free data retrieval call binding the contract method 0xc2aed302.
//
// Solidity: function AddrPrecompile() view returns(address)
func (_Cw1155 *Cw1155Caller) AddrPrecompile(opts *bind.CallOpts) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// AddrPrecompile is a free data retrieval call binding the contract method 0xc2aed302.
//
// Solidity: function AddrPrecompile() view returns(address)
func (_Cw1155 *Cw1155Session) AddrPrecompile() (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// AddrPrecompile is a free data retrieval call binding the contract method 0xc2aed302.
//
// Solidity: function AddrPrecompile() view returns(address)
func (_Cw1155 *Cw1155CallerSession) AddrPrecompile() (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// Cw1155Address is a free data retrieval call binding the contract method 0xb98933a0.
//
// Solidity: function Cw1155Address() view returns(string)
func (_Cw1155 *Cw1155Caller) Cw1155Address(opts *bind.CallOpts) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Cw1155Address is a free data retrieval call binding the contract method 0xb98933a0.
//
// Solidity: function Cw1155Address() view returns(string)
func (_Cw1155 *Cw1155Session) Cw1155Address() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Cw1155Address is a free data retrieval call binding the contract method 0xb98933a0.
//
// Solidity: function Cw1155Address() view returns(string)
func (_Cw1155 *Cw1155CallerSession) Cw1155Address() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// JsonPrecompile is a free data retrieval call binding the contract method 0xde4725cc.
//
// Solidity: function JsonPrecompile() view returns(address)
func (_Cw1155 *Cw1155Caller) JsonPrecompile(opts *bind.CallOpts) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// JsonPrecompile is a free data retrieval call binding the contract method 0xde4725cc.
//
// Solidity: function JsonPrecompile() view returns(address)
func (_Cw1155 *Cw1155Session) JsonPrecompile() (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// JsonPrecompile is a free data retrieval call binding the contract method 0xde4725cc.
//
// Solidity: function JsonPrecompile() view returns(address)
func (_Cw1155 *Cw1155CallerSession) JsonPrecompile() (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// WasmdPrecompile is a free data retrieval call binding the contract method 0xf00b0255.
//
// Solidity: function WasmdPrecompile() view returns(address)
func (_Cw1155 *Cw1155Caller) WasmdPrecompile(opts *bind.CallOpts) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// WasmdPrecompile is a free data retrieval call binding the contract method 0xf00b0255.
//
// Solidity: function WasmdPrecompile() view returns(address)
func (_Cw1155 *Cw1155Session) WasmdPrecompile() (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// WasmdPrecompile is a free data retrieval call binding the contract method 0xf00b0255.
//
// Solidity: function WasmdPrecompile() view returns(address)
func (_Cw1155 *Cw1155CallerSession) WasmdPrecompile() (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Cw1155 *Cw1155Caller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Cw1155 *Cw1155Session) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Cw1155 *Cw1155CallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[] balances)
func (_Cw1155 *Cw1155Caller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[] balances)
func (_Cw1155 *Cw1155Session) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[] balances)
func (_Cw1155 *Cw1155CallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_Cw1155 *Cw1155Caller) Exists(opts *bind.CallOpts, id *big.Int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_Cw1155 *Cw1155Session) Exists(id *big.Int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_Cw1155 *Cw1155CallerSession) Exists(id *big.Int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Cw1155 *Cw1155Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Cw1155 *Cw1155Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Cw1155 *Cw1155CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cw1155 *Cw1155Caller) Name(opts *bind.CallOpts) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cw1155 *Cw1155Session) Name() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cw1155 *Cw1155CallerSession) Name() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_Cw1155 *Cw1155Caller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil, nil
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_Cw1155 *Cw1155Session) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil, nil
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_Cw1155 *Cw1155CallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil, nil
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Cw1155 *Cw1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Cw1155 *Cw1155Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Cw1155 *Cw1155CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cw1155 *Cw1155Caller) Symbol(opts *bind.CallOpts) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cw1155 *Cw1155Session) Symbol() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cw1155 *Cw1155CallerSession) Symbol() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cw1155 *Cw1155Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cw1155 *Cw1155Session) TotalSupply() (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cw1155 *Cw1155CallerSession) TotalSupply() (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_Cw1155 *Cw1155Caller) TotalSupply0(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_Cw1155 *Cw1155Session) TotalSupply0(id *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_Cw1155 *Cw1155CallerSession) TotalSupply0(id *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_Cw1155 *Cw1155Caller) Uri(opts *bind.CallOpts, id *big.Int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_Cw1155 *Cw1155Session) Uri(id *big.Int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_Cw1155 *Cw1155CallerSession) Uri(id *big.Int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 amount) returns()
func (_Cw1155 *Cw1155Transactor) Burn(opts *bind.TransactOpts, account common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 amount) returns()
func (_Cw1155 *Cw1155Session) Burn(account common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 amount) returns()
func (_Cw1155 *Cw1155TransactorSession) Burn(account common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] amounts) returns()
func (_Cw1155 *Cw1155Transactor) BurnBatch(opts *bind.TransactOpts, account common.Address, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] amounts) returns()
func (_Cw1155 *Cw1155Session) BurnBatch(account common.Address, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] amounts) returns()
func (_Cw1155 *Cw1155TransactorSession) BurnBatch(account common.Address, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Cw1155 *Cw1155Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Cw1155 *Cw1155Session) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Cw1155 *Cw1155TransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Cw1155 *Cw1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Cw1155 *Cw1155Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Cw1155 *Cw1155TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Cw1155 *Cw1155Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Cw1155 *Cw1155Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Cw1155 *Cw1155TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cw1155ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Cw1155 contract.
type Cw1155ApprovalForAllIterator struct {
	Event *Cw1155ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Cw1155ApprovalForAllIterator) Next() bool {
	_ = "STUB: not implemented"
	// If the iterator failed, stop iterating
	return false
}

// If the iterator completed, deliver directly whatever's available

// Iterator still in progress, wait for either a data or an error event

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Cw1155ApprovalForAllIterator) Error() error {
	_ = "STUB: not implemented"

	// Close terminates the iteration process, releasing any pending underlying
	// resources.
	return nil
}

func (it *Cw1155ApprovalForAllIterator) Close() error { _ = "STUB: not implemented"; return nil }

// Cw1155ApprovalForAll represents a ApprovalForAll event raised by the Cw1155 contract.
type Cw1155ApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Cw1155 *Cw1155Filterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*Cw1155ApprovalForAllIterator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Cw1155 *Cw1155Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Cw1155ApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {
	_ = "STUB: not implemented"
	return *new(event.Subscription), nil
}

// New log arrived, parse the event and forward to the user

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Cw1155 *Cw1155Filterer) ParseApprovalForAll(log types.Log) (*Cw1155ApprovalForAll, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cw1155TransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Cw1155 contract.
type Cw1155TransferBatchIterator struct {
	Event *Cw1155TransferBatch // Event containing the contract specifics and raw log

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
func (it *Cw1155TransferBatchIterator) Next() bool {
	_ = "STUB: not implemented"
	// If the iterator failed, stop iterating
	return false
}

// If the iterator completed, deliver directly whatever's available

// Iterator still in progress, wait for either a data or an error event

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Cw1155TransferBatchIterator) Error() error {
	_ = "STUB: not implemented"

	// Close terminates the iteration process, releasing any pending underlying
	// resources.
	return nil
}

func (it *Cw1155TransferBatchIterator) Close() error { _ = "STUB: not implemented"; return nil }

// Cw1155TransferBatch represents a TransferBatch event raised by the Cw1155 contract.
type Cw1155TransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Cw1155 *Cw1155Filterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*Cw1155TransferBatchIterator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Cw1155 *Cw1155Filterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *Cw1155TransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {
	_ = "STUB: not implemented"
	return *new(event.Subscription), nil
}

// New log arrived, parse the event and forward to the user

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Cw1155 *Cw1155Filterer) ParseTransferBatch(log types.Log) (*Cw1155TransferBatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cw1155TransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Cw1155 contract.
type Cw1155TransferSingleIterator struct {
	Event *Cw1155TransferSingle // Event containing the contract specifics and raw log

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
func (it *Cw1155TransferSingleIterator) Next() bool {
	_ = "STUB: not implemented"
	// If the iterator failed, stop iterating
	return false
}

// If the iterator completed, deliver directly whatever's available

// Iterator still in progress, wait for either a data or an error event

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Cw1155TransferSingleIterator) Error() error {
	_ = "STUB: not implemented"

	// Close terminates the iteration process, releasing any pending underlying
	// resources.
	return nil
}

func (it *Cw1155TransferSingleIterator) Close() error { _ = "STUB: not implemented"; return nil }

// Cw1155TransferSingle represents a TransferSingle event raised by the Cw1155 contract.
type Cw1155TransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Cw1155 *Cw1155Filterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*Cw1155TransferSingleIterator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Cw1155 *Cw1155Filterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *Cw1155TransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {
	_ = "STUB: not implemented"
	return *new(event.Subscription), nil
}

// New log arrived, parse the event and forward to the user

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Cw1155 *Cw1155Filterer) ParseTransferSingle(log types.Log) (*Cw1155TransferSingle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cw1155URIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Cw1155 contract.
type Cw1155URIIterator struct {
	Event *Cw1155URI // Event containing the contract specifics and raw log

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
func (it *Cw1155URIIterator) Next() bool {
	_ = "STUB: not implemented"
	// If the iterator failed, stop iterating
	return false
}

// If the iterator completed, deliver directly whatever's available

// Iterator still in progress, wait for either a data or an error event

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Cw1155URIIterator) Error() error {
	_ = "STUB: not implemented"

	// Close terminates the iteration process, releasing any pending underlying
	// resources.
	return nil
}

func (it *Cw1155URIIterator) Close() error { _ = "STUB: not implemented"; return nil }

// Cw1155URI represents a URI event raised by the Cw1155 contract.
type Cw1155URI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Cw1155 *Cw1155Filterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*Cw1155URIIterator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Cw1155 *Cw1155Filterer) WatchURI(opts *bind.WatchOpts, sink chan<- *Cw1155URI, id []*big.Int) (event.Subscription, error) {
	_ = "STUB: not implemented"
	return *new(event.Subscription), nil
}

// New log arrived, parse the event and forward to the user

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Cw1155 *Cw1155Filterer) ParseURI(log types.Log) (*Cw1155URI, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

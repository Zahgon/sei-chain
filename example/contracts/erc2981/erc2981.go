// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc2981

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

// Erc2981MetaData contains all meta data concerning the Erc2981 contract.
var Erc2981MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"ERC2981InvalidDefaultRoyalty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC2981InvalidDefaultRoyaltyReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"ERC2981InvalidTokenRoyalty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC2981InvalidTokenRoyaltyReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"setDefaultRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Erc2981ABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc2981MetaData.ABI instead.
var Erc2981ABI = Erc2981MetaData.ABI

// Erc2981 is an auto generated Go binding around an Ethereum contract.
type Erc2981 struct {
	Erc2981Caller     // Read-only binding to the contract
	Erc2981Transactor // Write-only binding to the contract
	Erc2981Filterer   // Log filterer for contract events
}

// Erc2981Caller is an auto generated read-only Go binding around an Ethereum contract.
type Erc2981Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc2981Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc2981Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc2981Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc2981Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc2981Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc2981Session struct {
	Contract     *Erc2981          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc2981CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc2981CallerSession struct {
	Contract *Erc2981Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// Erc2981TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc2981TransactorSession struct {
	Contract     *Erc2981Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Erc2981Raw is an auto generated low-level Go binding around an Ethereum contract.
type Erc2981Raw struct {
	Contract *Erc2981 // Generic contract binding to access the raw methods on
}

// Erc2981CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc2981CallerRaw struct {
	Contract *Erc2981Caller // Generic read-only contract binding to access the raw methods on
}

// Erc2981TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc2981TransactorRaw struct {
	Contract *Erc2981Transactor // Generic write-only contract binding to access the raw methods on
}

// NewErc2981 creates a new instance of Erc2981, bound to a specific deployed contract.
func NewErc2981(address common.Address, backend bind.ContractBackend) (*Erc2981, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewErc2981Caller creates a new read-only instance of Erc2981, bound to a specific deployed contract.
func NewErc2981Caller(address common.Address, caller bind.ContractCaller) (*Erc2981Caller, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewErc2981Transactor creates a new write-only instance of Erc2981, bound to a specific deployed contract.
func NewErc2981Transactor(address common.Address, transactor bind.ContractTransactor) (*Erc2981Transactor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewErc2981Filterer creates a new log filterer instance of Erc2981, bound to a specific deployed contract.
func NewErc2981Filterer(address common.Address, filterer bind.ContractFilterer) (*Erc2981Filterer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// bindErc2981 binds a generic wrapper to an already deployed contract.
func bindErc2981(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc2981 *Erc2981Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc2981 *Erc2981Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc2981 *Erc2981Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc2981 *Erc2981CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc2981 *Erc2981TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc2981 *Erc2981TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Erc2981 *Erc2981Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Erc2981 *Erc2981Session) BalanceOf(owner common.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Erc2981 *Erc2981CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Erc2981 *Erc2981Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Erc2981 *Erc2981Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Erc2981 *Erc2981CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Erc2981 *Erc2981Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Erc2981 *Erc2981Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Erc2981 *Erc2981CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc2981 *Erc2981Caller) Name(opts *bind.CallOpts) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc2981 *Erc2981Session) Name() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc2981 *Erc2981CallerSession) Name() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Erc2981 *Erc2981Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Erc2981 *Erc2981Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Erc2981 *Erc2981CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_Erc2981 *Erc2981Caller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil, nil
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_Erc2981 *Erc2981Session) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil, nil
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_Erc2981 *Erc2981CallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil, nil
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Erc2981 *Erc2981Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Erc2981 *Erc2981Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Erc2981 *Erc2981CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc2981 *Erc2981Caller) Symbol(opts *bind.CallOpts) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc2981 *Erc2981Session) Symbol() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc2981 *Erc2981CallerSession) Symbol() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Erc2981 *Erc2981Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Erc2981 *Erc2981Session) TokenURI(tokenId *big.Int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Erc2981 *Erc2981CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Erc2981 *Erc2981Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Erc2981 *Erc2981Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Erc2981 *Erc2981TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc2981 *Erc2981Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc2981 *Erc2981Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc2981 *Erc2981TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Erc2981 *Erc2981Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Erc2981 *Erc2981Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Erc2981 *Erc2981TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc2981 *Erc2981Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc2981 *Erc2981Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc2981 *Erc2981TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0xb03b74aa.
//
// Solidity: function setDefaultRoyalty(address receiver) returns()
func (_Erc2981 *Erc2981Transactor) SetDefaultRoyalty(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0xb03b74aa.
//
// Solidity: function setDefaultRoyalty(address receiver) returns()
func (_Erc2981 *Erc2981Session) SetDefaultRoyalty(receiver common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0xb03b74aa.
//
// Solidity: function setDefaultRoyalty(address receiver) returns()
func (_Erc2981 *Erc2981TransactorSession) SetDefaultRoyalty(receiver common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc2981 *Erc2981Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc2981 *Erc2981Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc2981 *Erc2981TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Erc2981ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc2981 contract.
type Erc2981ApprovalIterator struct {
	Event *Erc2981Approval // Event containing the contract specifics and raw log

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
func (it *Erc2981ApprovalIterator) Next() bool {
	_ = "STUB: not implemented"
	// If the iterator failed, stop iterating
	return false
}

// If the iterator completed, deliver directly whatever's available

// Iterator still in progress, wait for either a data or an error event

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc2981ApprovalIterator) Error() error {
	_ = "STUB: not implemented"

	// Close terminates the iteration process, releasing any pending underlying
	// resources.
	return nil
}

func (it *Erc2981ApprovalIterator) Close() error { _ = "STUB: not implemented"; return nil }

// Erc2981Approval represents a Approval event raised by the Erc2981 contract.
type Erc2981Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Erc2981 *Erc2981Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*Erc2981ApprovalIterator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Erc2981 *Erc2981Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc2981Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {
	_ = "STUB: not implemented"
	return *new(event.Subscription), nil
}

// New log arrived, parse the event and forward to the user

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Erc2981 *Erc2981Filterer) ParseApproval(log types.Log) (*Erc2981Approval, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Erc2981ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Erc2981 contract.
type Erc2981ApprovalForAllIterator struct {
	Event *Erc2981ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Erc2981ApprovalForAllIterator) Next() bool {
	_ = "STUB: not implemented"
	// If the iterator failed, stop iterating
	return false
}

// If the iterator completed, deliver directly whatever's available

// Iterator still in progress, wait for either a data or an error event

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc2981ApprovalForAllIterator) Error() error {
	_ = "STUB: not implemented"

	// Close terminates the iteration process, releasing any pending underlying
	// resources.
	return nil
}

func (it *Erc2981ApprovalForAllIterator) Close() error { _ = "STUB: not implemented"; return nil }

// Erc2981ApprovalForAll represents a ApprovalForAll event raised by the Erc2981 contract.
type Erc2981ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Erc2981 *Erc2981Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*Erc2981ApprovalForAllIterator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Erc2981 *Erc2981Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Erc2981ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {
	_ = "STUB: not implemented"
	return *new(event.Subscription), nil
}

// New log arrived, parse the event and forward to the user

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Erc2981 *Erc2981Filterer) ParseApprovalForAll(log types.Log) (*Erc2981ApprovalForAll, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Erc2981TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc2981 contract.
type Erc2981TransferIterator struct {
	Event *Erc2981Transfer // Event containing the contract specifics and raw log

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
func (it *Erc2981TransferIterator) Next() bool {
	_ = "STUB: not implemented"
	// If the iterator failed, stop iterating
	return false
}

// If the iterator completed, deliver directly whatever's available

// Iterator still in progress, wait for either a data or an error event

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc2981TransferIterator) Error() error {
	_ = "STUB: not implemented"

	// Close terminates the iteration process, releasing any pending underlying
	// resources.
	return nil
}

func (it *Erc2981TransferIterator) Close() error { _ = "STUB: not implemented"; return nil }

// Erc2981Transfer represents a Transfer event raised by the Erc2981 contract.
type Erc2981Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Erc2981 *Erc2981Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*Erc2981TransferIterator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Erc2981 *Erc2981Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc2981Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {
	_ = "STUB: not implemented"
	return *new(event.Subscription), nil
}

// New log arrived, parse the event and forward to the user

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Erc2981 *Erc2981Filterer) ParseTransfer(log types.Log) (*Erc2981Transfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cw721

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

// Cw721MetaData contains all meta data concerning the Cw721 contract.
var Cw721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Cw721Address_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"ERC2981InvalidDefaultRoyalty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC2981InvalidDefaultRoyaltyReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"ERC2981InvalidTokenRoyalty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC2981InvalidTokenRoyaltyReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"}],\"name\":\"NotImplemented\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"}],\"name\":\"NotImplementedOnCosmwasmContract\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AddrPrecompile\",\"outputs\":[{\"internalType\":\"contractIAddr\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Cw721Address\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"JsonPrecompile\",\"outputs\":[{\"internalType\":\"contractIJson\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WasmdPrecompile\",\"outputs\":[{\"internalType\":\"contractIWasmd\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Cw721ABI is the input ABI used to generate the binding from.
// Deprecated: Use Cw721MetaData.ABI instead.
var Cw721ABI = Cw721MetaData.ABI

// Cw721 is an auto generated Go binding around an Ethereum contract.
type Cw721 struct {
	Cw721Caller     // Read-only binding to the contract
	Cw721Transactor // Write-only binding to the contract
	Cw721Filterer   // Log filterer for contract events
}

// Cw721Caller is an auto generated read-only Go binding around an Ethereum contract.
type Cw721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Cw721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Cw721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Cw721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Cw721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Cw721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Cw721Session struct {
	Contract     *Cw721            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Cw721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Cw721CallerSession struct {
	Contract *Cw721Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Cw721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Cw721TransactorSession struct {
	Contract     *Cw721Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Cw721Raw is an auto generated low-level Go binding around an Ethereum contract.
type Cw721Raw struct {
	Contract *Cw721 // Generic contract binding to access the raw methods on
}

// Cw721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Cw721CallerRaw struct {
	Contract *Cw721Caller // Generic read-only contract binding to access the raw methods on
}

// Cw721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Cw721TransactorRaw struct {
	Contract *Cw721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCw721 creates a new instance of Cw721, bound to a specific deployed contract.
func NewCw721(address common.Address, backend bind.ContractBackend) (*Cw721, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewCw721Caller creates a new read-only instance of Cw721, bound to a specific deployed contract.
func NewCw721Caller(address common.Address, caller bind.ContractCaller) (*Cw721Caller, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewCw721Transactor creates a new write-only instance of Cw721, bound to a specific deployed contract.
func NewCw721Transactor(address common.Address, transactor bind.ContractTransactor) (*Cw721Transactor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewCw721Filterer creates a new log filterer instance of Cw721, bound to a specific deployed contract.
func NewCw721Filterer(address common.Address, filterer bind.ContractFilterer) (*Cw721Filterer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// bindCw721 binds a generic wrapper to an already deployed contract.
func bindCw721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cw721 *Cw721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cw721 *Cw721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cw721 *Cw721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cw721 *Cw721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cw721 *Cw721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cw721 *Cw721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddrPrecompile is a free data retrieval call binding the contract method 0xc2aed302.
//
// Solidity: function AddrPrecompile() view returns(address)
func (_Cw721 *Cw721Caller) AddrPrecompile(opts *bind.CallOpts) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// AddrPrecompile is a free data retrieval call binding the contract method 0xc2aed302.
//
// Solidity: function AddrPrecompile() view returns(address)
func (_Cw721 *Cw721Session) AddrPrecompile() (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// AddrPrecompile is a free data retrieval call binding the contract method 0xc2aed302.
//
// Solidity: function AddrPrecompile() view returns(address)
func (_Cw721 *Cw721CallerSession) AddrPrecompile() (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// Cw721Address is a free data retrieval call binding the contract method 0x5c4aead7.
//
// Solidity: function Cw721Address() view returns(string)
func (_Cw721 *Cw721Caller) Cw721Address(opts *bind.CallOpts) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Cw721Address is a free data retrieval call binding the contract method 0x5c4aead7.
//
// Solidity: function Cw721Address() view returns(string)
func (_Cw721 *Cw721Session) Cw721Address() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Cw721Address is a free data retrieval call binding the contract method 0x5c4aead7.
//
// Solidity: function Cw721Address() view returns(string)
func (_Cw721 *Cw721CallerSession) Cw721Address() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// JsonPrecompile is a free data retrieval call binding the contract method 0xde4725cc.
//
// Solidity: function JsonPrecompile() view returns(address)
func (_Cw721 *Cw721Caller) JsonPrecompile(opts *bind.CallOpts) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// JsonPrecompile is a free data retrieval call binding the contract method 0xde4725cc.
//
// Solidity: function JsonPrecompile() view returns(address)
func (_Cw721 *Cw721Session) JsonPrecompile() (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// JsonPrecompile is a free data retrieval call binding the contract method 0xde4725cc.
//
// Solidity: function JsonPrecompile() view returns(address)
func (_Cw721 *Cw721CallerSession) JsonPrecompile() (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// WasmdPrecompile is a free data retrieval call binding the contract method 0xf00b0255.
//
// Solidity: function WasmdPrecompile() view returns(address)
func (_Cw721 *Cw721Caller) WasmdPrecompile(opts *bind.CallOpts) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// WasmdPrecompile is a free data retrieval call binding the contract method 0xf00b0255.
//
// Solidity: function WasmdPrecompile() view returns(address)
func (_Cw721 *Cw721Session) WasmdPrecompile() (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// WasmdPrecompile is a free data retrieval call binding the contract method 0xf00b0255.
//
// Solidity: function WasmdPrecompile() view returns(address)
func (_Cw721 *Cw721CallerSession) WasmdPrecompile() (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Cw721 *Cw721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Cw721 *Cw721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Cw721 *Cw721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Cw721 *Cw721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Cw721 *Cw721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Cw721 *Cw721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Cw721 *Cw721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Cw721 *Cw721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Cw721 *Cw721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cw721 *Cw721Caller) Name(opts *bind.CallOpts) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cw721 *Cw721Session) Name() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cw721 *Cw721CallerSession) Name() (string, error) { _ = "STUB: not implemented"; return "", nil }

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Cw721 *Cw721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Cw721 *Cw721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Cw721 *Cw721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_Cw721 *Cw721Caller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil, nil
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_Cw721 *Cw721Session) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil, nil
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_Cw721 *Cw721CallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil, nil
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Cw721 *Cw721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Cw721 *Cw721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Cw721 *Cw721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cw721 *Cw721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cw721 *Cw721Session) Symbol() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cw721 *Cw721CallerSession) Symbol() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 ) view returns(uint256)
func (_Cw721 *Cw721Caller) TokenByIndex(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 ) view returns(uint256)
func (_Cw721 *Cw721Session) TokenByIndex(arg0 *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 ) view returns(uint256)
func (_Cw721 *Cw721CallerSession) TokenByIndex(arg0 *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address , uint256 ) view returns(uint256)
func (_Cw721 *Cw721Caller) TokenOfOwnerByIndex(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address , uint256 ) view returns(uint256)
func (_Cw721 *Cw721Session) TokenOfOwnerByIndex(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address , uint256 ) view returns(uint256)
func (_Cw721 *Cw721CallerSession) TokenOfOwnerByIndex(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Cw721 *Cw721Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Cw721 *Cw721Session) TokenURI(tokenId *big.Int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Cw721 *Cw721CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cw721 *Cw721Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cw721 *Cw721Session) TotalSupply() (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cw721 *Cw721CallerSession) TotalSupply() (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address approved, uint256 tokenId) returns()
func (_Cw721 *Cw721Transactor) Approve(opts *bind.TransactOpts, approved common.Address, tokenId *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address approved, uint256 tokenId) returns()
func (_Cw721 *Cw721Session) Approve(approved common.Address, tokenId *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address approved, uint256 tokenId) returns()
func (_Cw721 *Cw721TransactorSession) Approve(approved common.Address, tokenId *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Cw721 *Cw721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Cw721 *Cw721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Cw721 *Cw721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Cw721 *Cw721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Cw721 *Cw721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Cw721 *Cw721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Cw721 *Cw721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Cw721 *Cw721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Cw721 *Cw721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Cw721 *Cw721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Cw721 *Cw721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Cw721 *Cw721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cw721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Cw721 contract.
type Cw721ApprovalIterator struct {
	Event *Cw721Approval // Event containing the contract specifics and raw log

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
func (it *Cw721ApprovalIterator) Next() bool {
	_ = "STUB: not implemented"
	// If the iterator failed, stop iterating
	return false
}

// If the iterator completed, deliver directly whatever's available

// Iterator still in progress, wait for either a data or an error event

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Cw721ApprovalIterator) Error() error {
	_ = "STUB: not implemented"

	// Close terminates the iteration process, releasing any pending underlying
	// resources.
	return nil
}

func (it *Cw721ApprovalIterator) Close() error { _ = "STUB: not implemented"; return nil }

// Cw721Approval represents a Approval event raised by the Cw721 contract.
type Cw721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Cw721 *Cw721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*Cw721ApprovalIterator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Cw721 *Cw721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Cw721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {
	_ = "STUB: not implemented"
	return *new(event.Subscription), nil
}

// New log arrived, parse the event and forward to the user

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Cw721 *Cw721Filterer) ParseApproval(log types.Log) (*Cw721Approval, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cw721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Cw721 contract.
type Cw721ApprovalForAllIterator struct {
	Event *Cw721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Cw721ApprovalForAllIterator) Next() bool {
	_ = "STUB: not implemented"
	// If the iterator failed, stop iterating
	return false
}

// If the iterator completed, deliver directly whatever's available

// Iterator still in progress, wait for either a data or an error event

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Cw721ApprovalForAllIterator) Error() error {
	_ = "STUB: not implemented"

	// Close terminates the iteration process, releasing any pending underlying
	// resources.
	return nil
}

func (it *Cw721ApprovalForAllIterator) Close() error { _ = "STUB: not implemented"; return nil }

// Cw721ApprovalForAll represents a ApprovalForAll event raised by the Cw721 contract.
type Cw721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Cw721 *Cw721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*Cw721ApprovalForAllIterator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Cw721 *Cw721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Cw721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {
	_ = "STUB: not implemented"
	return *new(event.Subscription), nil
}

// New log arrived, parse the event and forward to the user

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Cw721 *Cw721Filterer) ParseApprovalForAll(log types.Log) (*Cw721ApprovalForAll, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cw721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Cw721 contract.
type Cw721TransferIterator struct {
	Event *Cw721Transfer // Event containing the contract specifics and raw log

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
func (it *Cw721TransferIterator) Next() bool {
	_ = "STUB: not implemented"
	// If the iterator failed, stop iterating
	return false
}

// If the iterator completed, deliver directly whatever's available

// Iterator still in progress, wait for either a data or an error event

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Cw721TransferIterator) Error() error {
	_ = "STUB: not implemented"

	// Close terminates the iteration process, releasing any pending underlying
	// resources.
	return nil
}

func (it *Cw721TransferIterator) Close() error { _ = "STUB: not implemented"; return nil }

// Cw721Transfer represents a Transfer event raised by the Cw721 contract.
type Cw721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Cw721 *Cw721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*Cw721TransferIterator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Cw721 *Cw721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Cw721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {
	_ = "STUB: not implemented"
	return *new(event.Subscription), nil
}

// New log arrived, parse the event and forward to the user

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Cw721 *Cw721Filterer) ParseTransfer(log types.Log) (*Cw721Transfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

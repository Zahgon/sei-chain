package types

import (
	"github.com/gogo/protobuf/proto"

	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

var (
	_ AccountI                           = (*BaseAccount)(nil)
	_ GenesisAccount                     = (*BaseAccount)(nil)
	_ codectypes.UnpackInterfacesMessage = (*BaseAccount)(nil)
	_ GenesisAccount                     = (*ModuleAccount)(nil)
	_ ModuleAccountI                     = (*ModuleAccount)(nil)
)

// NewBaseAccount creates a new BaseAccount object
func NewBaseAccount(address sdk.AccAddress, pubKey cryptotypes.PubKey, accountNumber, sequence uint64) *BaseAccount {
	_ = "STUB: not implemented"
	return nil
}

// ProtoBaseAccount - a prototype function for BaseAccount
func ProtoBaseAccount() AccountI {
	_ = "STUB: not implemented"
	return *

	// NewBaseAccountWithAddress - returns a new base account with a given address
	// leaving AccountNumber and Sequence to zero.
	new(AccountI)
}

func NewBaseAccountWithAddress(addr sdk.AccAddress) *BaseAccount {
	_ = "STUB: not implemented"
	return nil
}

// GetAddress - Implements sdk.AccountI.
func (acc BaseAccount) GetAddress() sdk.AccAddress {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}

// SetAddress - Implements sdk.AccountI.
func (acc *BaseAccount) SetAddress(addr sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// GetPubKey - Implements sdk.AccountI.
func (acc BaseAccount) GetPubKey() (pk cryptotypes.PubKey) {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PubKey)
}

// SetPubKey - Implements sdk.AccountI.
func (acc *BaseAccount) SetPubKey(pubKey cryptotypes.PubKey) error {
	_ = "STUB: not implemented"
	return nil
}

// GetAccountNumber - Implements AccountI
func (acc BaseAccount) GetAccountNumber() uint64 { _ = "STUB: not implemented"; return 0 }

// SetAccountNumber - Implements AccountI
func (acc *BaseAccount) SetAccountNumber(accNumber uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// GetSequence - Implements sdk.AccountI.
func (acc BaseAccount) GetSequence() uint64 { _ = "STUB: not implemented"; return 0 }

// SetSequence - Implements sdk.AccountI.
func (acc *BaseAccount) SetSequence(seq uint64) error { _ = "STUB: not implemented"; return nil }

// Validate checks for errors on the account fields
func (acc BaseAccount) Validate() error { _ = "STUB: not implemented"; return nil }

func (acc BaseAccount) String() string { _ = "STUB: not implemented"; return "" }

// MarshalYAML returns the YAML representation of an account.
func (acc BaseAccount) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (acc BaseAccount) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// NewModuleAddress creates an AccAddress from the hash of the module's name
func NewModuleAddress(name string) sdk.AccAddress {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}

// NewEmptyModuleAccount creates a empty ModuleAccount from a string
func NewEmptyModuleAccount(name string, permissions ...string) *ModuleAccount {
	_ = "STUB: not implemented"
	return nil
}

// NewModuleAccount creates a new ModuleAccount instance
func NewModuleAccount(ba *BaseAccount, name string, permissions ...string) *ModuleAccount {
	_ = "STUB: not implemented"
	return nil
}

// HasPermission returns whether or not the module account has permission.
func (ma ModuleAccount) HasPermission(permission string) bool {
	_ = "STUB: not implemented"
	return false
}

// GetName returns the the name of the holder's module
func (ma ModuleAccount) GetName() string {
	_ = "STUB: not implemented"

	// GetPermissions returns permissions granted to the module account
	return ""
}

func (ma ModuleAccount) GetPermissions() []string { _ = "STUB: not implemented"; return nil }

// SetPubKey - Implements AccountI
func (ma ModuleAccount) SetPubKey(pubKey cryptotypes.PubKey) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate checks for errors on the account fields
func (ma ModuleAccount) Validate() error { _ = "STUB: not implemented"; return nil }

type moduleAccountPretty struct {
	Address       sdk.AccAddress `json:"address" yaml:"address"`
	PubKey        string         `json:"public_key" yaml:"public_key"`
	AccountNumber uint64         `json:"account_number" yaml:"account_number"`
	Sequence      uint64         `json:"sequence" yaml:"sequence"`
	Name          string         `json:"name" yaml:"name"`
	Permissions   []string       `json:"permissions" yaml:"permissions"`
}

func (ma ModuleAccount) String() string { _ = "STUB: not implemented"; return "" }

// MarshalYAML returns the YAML representation of a ModuleAccount.
func (ma ModuleAccount) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalJSON returns the JSON representation of a ModuleAccount.
func (ma ModuleAccount) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON unmarshals raw JSON bytes into a ModuleAccount.
func (ma *ModuleAccount) UnmarshalJSON(bz []byte) error { _ = "STUB: not implemented"; return nil }

// AccountI is an interface used to store coins at a given address within state.
// It presumes a notion of sequence numbers for replay protection,
// a notion of account numbers for replay protection for previously pruned accounts,
// and a pubkey for authentication purposes.
//
// Many complex conditions can be used in the concrete struct which implements AccountI.
type AccountI interface {
	proto.Message

	GetAddress() sdk.AccAddress
	SetAddress(sdk.AccAddress) error // errors if already set.

	GetPubKey() cryptotypes.PubKey // can return nil.
	SetPubKey(cryptotypes.PubKey) error

	GetAccountNumber() uint64
	SetAccountNumber(uint64) error

	GetSequence() uint64
	SetSequence(uint64) error

	// Ensure that account implements stringer
	String() string
}

// ModuleAccountI defines an account interface for modules that hold tokens in
// an escrow.
type ModuleAccountI interface {
	AccountI

	GetName() string
	GetPermissions() []string
	HasPermission(string) bool
}

// GenesisAccounts defines a slice of GenesisAccount objects
type GenesisAccounts []GenesisAccount

// Contains returns true if the given address exists in a slice of GenesisAccount
// objects.
func (ga GenesisAccounts) Contains(addr sdk.Address) bool { _ = "STUB: not implemented"; return false }

// GenesisAccount defines a genesis account that embeds an AccountI with validation capabilities.
type GenesisAccount interface {
	AccountI

	Validate() error
}

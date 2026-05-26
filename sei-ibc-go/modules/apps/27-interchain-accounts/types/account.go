package types

import (
	"regexp"

	crypto "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	authtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/types"
)

var (
	_ authtypes.GenesisAccount = (*InterchainAccount)(nil)
	_ InterchainAccountI       = (*InterchainAccount)(nil)
)

// DefaultMaxAddrLength defines the default maximum character length used in validation of addresses
var DefaultMaxAddrLength = 128

// isValidAddr defines a regular expression to check if the provided string consists of
// strictly alphanumeric characters and is non empty.
var isValidAddr = regexp.MustCompile("^[a-zA-Z0-9]+$").MatchString

// InterchainAccountI wraps the authtypes.AccountI interface
type InterchainAccountI interface {
	authtypes.AccountI
}

// interchainAccountPretty defines an unexported struct used for encoding the InterchainAccount details
type interchainAccountPretty struct {
	Address       sdk.AccAddress `json:"address" yaml:"address"`
	PubKey        string         `json:"public_key" yaml:"public_key"`
	AccountNumber uint64         `json:"account_number" yaml:"account_number"`
	Sequence      uint64         `json:"sequence" yaml:"sequence"`
	AccountOwner  string         `json:"account_owner" yaml:"account_owner"`
}

// GenerateAddress returns an sdk.AccAddress derived using the provided module account address and connection and port identifiers.
// The sdk.AccAddress returned is a sub-address of the module account, using the host chain connection ID and controller chain's port ID as the derivation key
// Deprecated: this function is deprecated! Please use GenerateUniqueAddress in favour of GenerateAddress
func GenerateAddress(moduleAccAddr sdk.AccAddress, connectionID, portID string) sdk.AccAddress {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}

// GenerateUniqueAddress returns an sdk.AccAddress derived using a host module account address, host connection ID, the controller portID,
// the current block app hash, and the current block data hash. The sdk.AccAddress returned is a sub-address of the host module account.
func GenerateUniqueAddress(ctx sdk.Context, connectionID, portID string) sdk.AccAddress {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}

// ValidateAccountAddress performs basic validation of interchain account addresses, enforcing constraints
// on address length and character set
func ValidateAccountAddress(addr string) error { _ = "STUB: not implemented"; return nil }

// NewInterchainAccount creates and returns a new InterchainAccount type
func NewInterchainAccount(ba *authtypes.BaseAccount, accountOwner string) *InterchainAccount {
	_ = "STUB: not implemented"
	return nil
}

// SetPubKey implements the authtypes.AccountI interface
func (ia InterchainAccount) SetPubKey(pubKey crypto.PubKey) error {
	_ = "STUB: not implemented"
	return nil
}

// SetSequence implements the authtypes.AccountI interface
func (ia InterchainAccount) SetSequence(seq uint64) error { _ = "STUB: not implemented"; return nil }

// Validate implements basic validation of the InterchainAccount
func (ia InterchainAccount) Validate() error { _ = "STUB: not implemented"; return nil }

// String returns a string representation of the InterchainAccount
func (ia InterchainAccount) String() string { _ = "STUB: not implemented"; return "" }

// MarshalYAML returns the YAML representation of the InterchainAccount
func (ia InterchainAccount) MarshalYAML() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalJSON returns the JSON representation of the InterchainAccount
func (ia InterchainAccount) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalJSON unmarshals raw JSON bytes into the InterchainAccount
func (ia *InterchainAccount) UnmarshalJSON(bz []byte) error { _ = "STUB: not implemented"; return nil }

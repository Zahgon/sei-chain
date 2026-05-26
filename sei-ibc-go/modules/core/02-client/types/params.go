package types

import (
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

var (
	// DefaultAllowedClients are "06-solomachine" and "07-tendermint"
	DefaultAllowedClients = []string{exported.Solomachine, exported.Tendermint}

	// KeyAllowedClients is store's key for AllowedClients Params
	KeyAllowedClients = []byte("AllowedClients")
)

// ParamKeyTable type declaration for parameters
func ParamKeyTable() paramtypes.KeyTable {
	_ = "STUB: not implemented"
	return *new(paramtypes.KeyTable)
}

// NewParams creates a new parameter configuration for the ibc client module
func NewParams(allowedClients ...string) Params { _ = "STUB: not implemented"; return *new(Params) }

// DefaultParams is the default parameter configuration for the ibc-client module
func DefaultParams() Params { _ = "STUB: not implemented"; return *new(Params) }

// Validate all ibc-client module parameters
func (p Params) Validate() error { _ = "STUB: not implemented"; return nil }

// ParamSetPairs implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	_ = "STUB: not implemented"
	return *new(paramtypes.ParamSetPairs)
}

// IsAllowedClient checks if the given client type is registered on the allowlist.
func (p Params) IsAllowedClient(clientType string) bool { _ = "STUB: not implemented"; return false }

func validateClients(i interface{}) error { _ = "STUB: not implemented"; return nil }

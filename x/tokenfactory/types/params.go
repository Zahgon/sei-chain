package types

import (
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
)

// DefaultDenomAllowListMaxSize default denom allowlist max size and can be overridden by governance proposal.
const DefaultDenomAllowListMaxSize = 2000

// ParamKeyTable ParamTable for tokenfactory module.
func ParamKeyTable() paramtypes.KeyTable {
	_ = "STUB: not implemented"
	return *new(paramtypes.KeyTable)
}

// DefaultParams default tokenfactory module parameters.
func DefaultParams() Params { _ = "STUB: not implemented"; return *new(Params) }

// Validate validate params.
func (p Params) Validate() error { _ = "STUB: not implemented"; return nil }

// ParamSetPairs Implements params.ParamSet.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	_ = "STUB: not implemented"
	return *new(paramtypes.ParamSetPairs)
}

// validateDenomAllowListMaxSize validates a parameter value is within a valid range.
func validateDenomAllowListMaxSize(i interface{}) error { _ = "STUB: not implemented"; return nil }

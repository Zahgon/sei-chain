package types

import (
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	_ = "STUB: not implemented"
	return *new(paramtypes.KeyTable)
}

// NewParams creates a new Params instance
func NewParams() Params {
	_ = "STUB: not implemented"

	// DefaultParams returns a default set of parameters
	return *new(Params)
}

func DefaultParams() Params {
	_ = "STUB: not implemented"

	// ParamSetPairs get the params.ParamSet
	return *new(Params)
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	_ = "STUB: not implemented"
	return *new(paramtypes.ParamSetPairs)
}

// Validate validates the set of params
func (p Params) Validate() error { _ = "STUB: not implemented"; return nil }

// String implements the Stringer interface.
func (p Params) String() string { _ = "STUB: not implemented"; return "" }

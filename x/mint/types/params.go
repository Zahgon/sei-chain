package types

import (
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
)

// Parameter store keys
var (
	KeyMintDenom            = []byte("MintDenom")
	KeyTokenReleaseSchedule = []byte("TokenReleaseSchedule")
)

// ParamTable for minting module.
func ParamKeyTable() paramtypes.KeyTable {
	_ = "STUB: not implemented"
	return *new(paramtypes.KeyTable)
}

func NewParams(
	mintDenom string, tokenReleaseSchedule []ScheduledTokenRelease,
) Params {
	_ = "STUB: not implemented"
	return *new(Params)
}

// default minting module parameters
func DefaultParams() Params { _ = "STUB: not implemented"; return *new(Params) }

// validate params
func (p Params) Validate() error { _ = "STUB: not implemented"; return nil }

// String implements the Stringer interface.
func (p Params) String() string { _ = "STUB: not implemented"; return "" }

func (p Version2Params) String() string { _ = "STUB: not implemented"; return "" }

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	_ = "STUB: not implemented"
	return *new(paramtypes.ParamSetPairs)
}

// Used for v2 -> v3 migration
func (p *Version2Params) ParamSetPairs() paramtypes.ParamSetPairs {
	_ = "STUB: not implemented"
	return *new(paramtypes.ParamSetPairs)
}

func validateMintDenom(i interface{}) error { _ = "STUB: not implemented"; return nil }

func SortTokenReleaseCalendar(tokenReleaseSchedule []ScheduledTokenRelease) []ScheduledTokenRelease {
	_ = "STUB: not implemented"
	return nil
}

func validateTokenReleaseSchedule(i interface{}) error { _ = "STUB: not implemented"; return nil }

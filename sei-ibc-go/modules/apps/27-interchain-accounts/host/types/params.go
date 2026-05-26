package types

import (
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
)

const (
	// DefaultHostEnabled is the default value for the host param (set to true)
	DefaultHostEnabled = true
)

var (
	// KeyHostEnabled is the store key for HostEnabled Params
	KeyHostEnabled = []byte("HostEnabled")
	// KeyAllowMessages is the store key for the AllowMessages Params
	KeyAllowMessages = []byte("AllowMessages")
)

// ParamKeyTable type declaration for parameters
func ParamKeyTable() paramtypes.KeyTable {
	_ = "STUB: not implemented"
	return *new(paramtypes.KeyTable)
}

// NewParams creates a new parameter configuration for the host submodule
func NewParams(enableHost bool, allowMsgs []string) Params {
	_ = "STUB: not implemented"
	return *new(Params)
}

// DefaultParams is the default parameter configuration for the host submodule
func DefaultParams() Params { _ = "STUB: not implemented"; return *new(Params) }

// Validate validates all host submodule parameters
func (p Params) Validate() error { _ = "STUB: not implemented"; return nil }

// ParamSetPairs implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	_ = "STUB: not implemented"
	return *new(paramtypes.ParamSetPairs)
}

func validateEnabled(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateAllowlist(i interface{}) error { _ = "STUB: not implemented"; return nil }

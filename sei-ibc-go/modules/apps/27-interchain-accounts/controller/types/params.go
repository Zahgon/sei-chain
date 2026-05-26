package types

import (
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
)

const (
	// DefaultControllerEnabled is the default value for the controller param (set to true)
	DefaultControllerEnabled = true
)

// KeyControllerEnabled is the store key for ControllerEnabled Params
var KeyControllerEnabled = []byte("ControllerEnabled")

// ParamKeyTable type declaration for parameters
func ParamKeyTable() paramtypes.KeyTable {
	_ = "STUB: not implemented"
	return *new(paramtypes.KeyTable)
}

// NewParams creates a new parameter configuration for the controller submodule
func NewParams(enableController bool) Params { _ = "STUB: not implemented"; return *new(Params) }

// DefaultParams is the default parameter configuration for the controller submodule
func DefaultParams() Params { _ = "STUB: not implemented"; return *new(Params) }

// Validate validates all controller submodule parameters
func (p Params) Validate() error { _ = "STUB: not implemented"; return nil }

// ParamSetPairs implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	_ = "STUB: not implemented"
	return *new(paramtypes.ParamSetPairs)
}

func validateEnabled(i interface{}) error { _ = "STUB: not implemented"; return nil }

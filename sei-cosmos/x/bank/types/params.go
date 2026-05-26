package types

import (
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
)

const (
	// DefaultSendEnabled enabled
	DefaultSendEnabled = true
)

var (
	// KeySendEnabled is store's key for SendEnabled Params
	KeySendEnabled = []byte("SendEnabled")
	// KeyDefaultSendEnabled is store's key for the DefaultSendEnabled option
	KeyDefaultSendEnabled = []byte("DefaultSendEnabled")
)

// ParamKeyTable for bank module.
func ParamKeyTable() paramtypes.KeyTable {
	_ = "STUB: not implemented"
	return *new(paramtypes.KeyTable)
}

// NewParams creates a new parameter configuration for the bank module
func NewParams(defaultSendEnabled bool, sendEnabledParams SendEnabledParams) Params {
	_ = "STUB: not implemented"
	return *new(Params)
}

// DefaultParams is the default parameter configuration for the bank module
func DefaultParams() Params { _ = "STUB: not implemented"; return *new(Params) }

// The default send enabled value allows send transfers for all coin denoms

// Validate all bank module parameters
func (p Params) Validate() error { _ = "STUB: not implemented"; return nil }

// String implements the Stringer interface.
func (p Params) String() string { _ = "STUB: not implemented"; return "" }

// SendEnabledDenom returns true if the given denom is enabled for sending
func (p Params) SendEnabledDenom(denom string) bool { _ = "STUB: not implemented"; return false }

// SetSendEnabledParam returns an updated set of Parameters with the given denom
// send enabled flag set.
func (p Params) SetSendEnabledParam(denom string, sendEnabled bool) Params {
	_ = "STUB: not implemented"
	return *new(Params)
}

// ParamSetPairs implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	_ = "STUB: not implemented"
	return *new(paramtypes.ParamSetPairs)
}

// SendEnabledParams is a collection of parameters indicating if a coin denom is enabled for sending
type SendEnabledParams []*SendEnabled

func validateSendEnabledParams(i interface{}) error { _ = "STUB: not implemented"; return nil }

// ensure each denom is only registered one time.

// NewSendEnabled creates a new SendEnabled object
// The denom may be left empty to control the global default setting of send_enabled
func NewSendEnabled(denom string, sendEnabled bool) *SendEnabled {
	_ = "STUB: not implemented"
	return nil
}

// String implements stringer insterface
func (se SendEnabled) String() string { _ = "STUB: not implemented"; return "" }

func validateSendEnabled(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateIsBool(i interface{}) error { _ = "STUB: not implemented"; return nil }

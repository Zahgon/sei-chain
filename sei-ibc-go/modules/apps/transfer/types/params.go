package types

import (
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
)

const (
	// DefaultSendEnabled enabled
	DefaultSendEnabled = true
	// DefaultReceiveEnabled enabled
	DefaultReceiveEnabled = true
)

var (
	// KeySendEnabled is store's key for SendEnabled Params
	KeySendEnabled = []byte("SendEnabled")
	// KeyReceiveEnabled is store's key for ReceiveEnabled Params
	KeyReceiveEnabled = []byte("ReceiveEnabled")
)

// ParamKeyTable type declaration for parameters
func ParamKeyTable() paramtypes.KeyTable {
	_ = "STUB: not implemented"
	return *new(paramtypes.KeyTable)
}

// NewParams creates a new parameter configuration for the ibc transfer module
func NewParams(enableSend, enableReceive bool) Params {
	_ = "STUB: not implemented"
	return *new(Params)
}

// DefaultParams is the default parameter configuration for the ibc-transfer module
func DefaultParams() Params { _ = "STUB: not implemented"; return *new(Params) }

// Validate all ibc-transfer module parameters
func (p Params) Validate() error { _ = "STUB: not implemented"; return nil }

// ParamSetPairs implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	_ = "STUB: not implemented"
	return *new(paramtypes.ParamSetPairs)
}

func validateEnabled(i interface{}) error { _ = "STUB: not implemented"; return nil }

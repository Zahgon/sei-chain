package types

import (
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
)

var (
	KeyInboundEnabled  = []byte("InboundEnabled")
	KeyOutboundEnabled = []byte("OutboundEnabled")
)

type Params struct {
	InboundEnabled  bool `json:"inbound_enabled" yaml:"inbound_enabled"`
	OutboundEnabled bool `json:"outbound_enabled" yaml:"outbound_enabled"`
}

// ParamKeyTable for the ibc core module params
func ParamKeyTable() paramtypes.KeyTable {
	_ = "STUB: not implemented"
	return *new(paramtypes.KeyTable)
}

func NewParams(inbound, outbound bool) Params { _ = "STUB: not implemented"; return *new(Params) }

func DefaultParams() Params { _ = "STUB: not implemented"; return *new(Params) }

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	_ = "STUB: not implemented"
	return *new(paramtypes.ParamSetPairs)
}

func (p Params) Validate() error { _ = "STUB: not implemented"; return nil }

func validateBool(i interface{}) error { _ = "STUB: not implemented"; return nil }

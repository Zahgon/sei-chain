package types

import (
	"time"

	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
)

// DefaultTimePerBlock is the default value for maximum expected time per block (in nanoseconds).
const DefaultTimePerBlock = 30 * time.Second

// KeyMaxExpectedTimePerBlock is store's key for MaxExpectedTimePerBlock parameter
var KeyMaxExpectedTimePerBlock = []byte("MaxExpectedTimePerBlock")

// ParamKeyTable type declaration for parameters
func ParamKeyTable() paramtypes.KeyTable {
	_ = "STUB: not implemented"
	return *new(paramtypes.KeyTable)
}

// NewParams creates a new parameter configuration for the ibc connection module
func NewParams(timePerBlock uint64) Params { _ = "STUB: not implemented"; return *new(Params) }

// DefaultParams is the default parameter configuration for the ibc connection module
func DefaultParams() Params { _ = "STUB: not implemented"; return *new(Params) }

// Validate ensures MaxExpectedTimePerBlock is non-zero
func (p Params) Validate() error { _ = "STUB: not implemented"; return nil }

// ParamSetPairs implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	_ = "STUB: not implemented"
	return *new(paramtypes.ParamSetPairs)
}

func validateParams(i interface{}) error { _ = "STUB: not implemented"; return nil }

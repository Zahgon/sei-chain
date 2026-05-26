package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// invariant route
type InvarRoute struct {
	ModuleName string
	Route      string
	Invar      sdk.Invariant
}

// NewInvarRoute - create an InvarRoute object
func NewInvarRoute(moduleName, route string, invar sdk.Invariant) InvarRoute {
	_ = "STUB: not implemented"
	return *new(InvarRoute)
}

// get the full invariance route
func (i InvarRoute) FullRoute() string { _ = "STUB: not implemented"; return "" }

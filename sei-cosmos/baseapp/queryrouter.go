package baseapp

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

type QueryRouter struct {
	routes map[string]sdk.Querier
}

var _ sdk.QueryRouter = NewQueryRouter()

// NewQueryRouter returns a reference to a new QueryRouter.
func NewQueryRouter() *QueryRouter { _ = "STUB: not implemented"; return nil }

// AddRoute adds a query path to the router with a given Querier. It will panic
// if a duplicate route is given. The route must be alphanumeric.
func (qrt *QueryRouter) AddRoute(route string, q sdk.Querier) sdk.QueryRouter {
	_ = "STUB: not implemented"
	return *new(sdk.QueryRouter)
}

// paths are only the final extensions!
// Needed to ensure erroneous queries don't get into the state machine.

// Route returns the Querier for a given query route path.
func (qrt *QueryRouter) Route(path string) sdk.Querier {
	_ = "STUB: not implemented"
	return *new(sdk.Querier)
}

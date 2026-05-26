package baseapp

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

type Router struct {
	routes map[string]sdk.Handler
}

var _ sdk.Router = NewRouter()

// NewRouter returns a reference to a new router.
func NewRouter() *Router { _ = "STUB: not implemented"; return nil }

// AddRoute adds a route path to the router with a given handler. The route must
// be alphanumeric.
func (rtr *Router) AddRoute(route sdk.Route) sdk.Router {
	_ = "STUB: not implemented"
	return *new(sdk.Router)
}

// Route returns a handler for a given route path.
//
// TODO: Handle expressive matches.
func (rtr *Router) Route(_ sdk.Context, path string) sdk.Handler {
	_ = "STUB: not implemented"
	return *new(sdk.Handler)
}

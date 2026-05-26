package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/evidence/exported"
)

type (
	// Handler defines an agnostic Evidence handler. The handler is responsible
	// for executing all corresponding business logic necessary for verifying the
	// evidence as valid. In addition, the Handler may execute any necessary
	// slashing and potential jailing.
	Handler func(sdk.Context, exported.Evidence) error

	// Router defines a contract for which any Evidence handling module must
	// implement in order to route Evidence to registered Handlers.
	Router interface {
		AddRoute(r string, h Handler) Router
		HasRoute(r string) bool
		GetRoute(path string) Handler
		Seal()
		Sealed() bool
	}

	router struct {
		routes map[string]Handler
		sealed bool
	}
)

func NewRouter() Router { _ = "STUB: not implemented"; return *new(Router) }

// Seal prevents the router from any subsequent route handlers to be registered.
// Seal will panic if called more than once.
func (rtr *router) Seal() { _ = "STUB: not implemented"; return }

// Sealed returns a boolean signifying if the Router is sealed or not.
func (rtr router) Sealed() bool {
	_ = "STUB: not implemented"

	// AddRoute adds a governance handler for a given path. It returns the Router
	// so AddRoute calls can be linked. It will panic if the router is sealed.
	return false
}

func (rtr *router) AddRoute(path string, h Handler) Router {
	_ = "STUB: not implemented"
	return *new(Router)
}

// HasRoute returns true if the router has a path registered or false otherwise.
func (rtr *router) HasRoute(path string) bool { _ = "STUB: not implemented"; return false }

// GetRoute returns a Handler for a given path.
func (rtr *router) GetRoute(path string) Handler { _ = "STUB: not implemented"; return *new(Handler) }

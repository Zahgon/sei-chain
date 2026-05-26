package types

var _ Router = (*router)(nil)

// Router implements a governance Handler router.
//
// TODO: Use generic router (ref #3976).
type Router interface {
	AddRoute(r string, h Handler) (rtr Router)
	HasRoute(r string) bool
	GetRoute(path string) (h Handler)
	Seal()
}

type router struct {
	routes map[string]Handler
	sealed bool
}

// NewRouter creates a new Router interface instance
func NewRouter() Router { _ = "STUB: not implemented"; return *new(Router) }

// Seal seals the router which prohibits any subsequent route handlers to be
// added. Seal will panic if called more than once.
func (rtr *router) Seal() { _ = "STUB: not implemented"; return }

// AddRoute adds a governance handler for a given path. It returns the Router
// so AddRoute calls can be linked. It will panic if the router is sealed.
func (rtr *router) AddRoute(path string, h Handler) Router {
	_ = "STUB: not implemented"
	return *new(Router)
}

// HasRoute returns true if the router has a path registered or false otherwise.
func (rtr *router) HasRoute(path string) bool { _ = "STUB: not implemented"; return false }

// GetRoute returns a Handler for a given path.
func (rtr *router) GetRoute(path string) Handler { _ = "STUB: not implemented"; return *new(Handler) }

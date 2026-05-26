package types

// The router is a map from module name to the IBCModule
// which contains all the module-defined callbacks required by ICS-26
type Router struct {
	routes map[string]IBCModule
	sealed bool
}

func NewRouter() *Router { _ = "STUB: not implemented"; return nil }

// Seal prevents the Router from any subsequent route handlers to be registered.
// Seal will panic if called more than once.
func (rtr *Router) Seal() { _ = "STUB: not implemented"; return }

// Sealed returns a boolean signifying if the Router is sealed or not.
func (rtr Router) Sealed() bool {
	_ = "STUB: not implemented"

	// AddRoute adds IBCModule for a given module name. It returns the Router
	// so AddRoute calls can be linked. It will panic if the Router is sealed.
	return false
}

func (rtr *Router) AddRoute(module string, cbs IBCModule) *Router {
	_ = "STUB: not implemented"
	return nil
}

// HasRoute returns true if the Router has a module registered or false otherwise.
func (rtr *Router) HasRoute(module string) bool { _ = "STUB: not implemented"; return false }

// GetRoute returns a IBCModule for a given module.
func (rtr *Router) GetRoute(module string) (IBCModule, bool) {
	_ = "STUB: not implemented"
	return *new(IBCModule), false
}

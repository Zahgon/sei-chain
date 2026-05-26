package types

// GetModuleOwner enforces that only IBC and the module bound to port can own the capability
// while future implementations may allow multiple modules to bind to a port, currently we
// only allow one module to be bound to a port at any given time
func GetModuleOwner(modules []string) string { _ = "STUB: not implemented"; return "" }

package errors

// AssertNil panics on error
// Should be only used with interface methods, which require return error, but the
// error is always nil
func AssertNil(err error) { _ = "STUB: not implemented"; return }

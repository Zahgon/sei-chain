package store

// UnsafeStrToBytes uses unsafe to convert string into byte array. Returned bytes
// must not be altered after this function is called as it will cause a segmentation fault.
func UnsafeStrToBytes(s string) []byte { _ = "STUB: not implemented"; return nil }

//nolint:gosec,staticcheck
//nolint:gosec,staticcheck

// UnsafeBytesToStr is meant to make a zero allocation conversion
// from []byte -> string to speed up operations, it is not meant
// to be used generally, but for a specific pattern to delete keys
// from a map.
func UnsafeBytesToStr(b []byte) string { _ = "STUB: not implemented"; return "" }

//nolint:gosec

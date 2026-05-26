package utils

// Clone returns a copy of b[:len(b)].
// The result may have additional unused capacity.
// Clone(nil) returns nil.
func Clone(b []byte) []byte { _ = "STUB: not implemented"; return nil }

// Equal reports whether a and b
// are the same length and contain the same bytes.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []byte) bool {
	_ = "STUB: not implemented"
	// Neither cmd/compile nor gccgo allocates for these string conversions.
	return false
}

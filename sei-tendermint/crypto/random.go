package crypto

import (
	"io"
)

// This only uses the OS's randomness
func CRandBytes(numBytes int) []byte { _ = "STUB: not implemented"; return nil }

// Returns a crand.Reader.
func CReader() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

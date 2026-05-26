package tmhash

import (
	"crypto/sha256"
	"hash"
)

const (
	Size = sha256.Size
)

// New returns a new hash.Hash.
func New() hash.Hash {
	_ = "STUB: not implemented"
	return *

	// Sum returns the SHA256 of the bz.
	new(hash.Hash)
}

func Sum(bz []byte) []byte { _ = "STUB: not implemented"; return nil }

package v65

import (
	"crypto/ecdsa"
	"math/big"
)

func newPublicKey(x, y *big.Int) *ecdsa.PublicKey {
	_ = "STUB: not implemented"
	// Check if the given coordinates are valid
	return nil
}

// Verify verifies the given signature (r, s) for the given hash and public key (x, y).
// It returns true if the signature is valid, false otherwise.
func Verify(hash []byte, r, s, x, y *big.Int) bool {
	_ = "STUB: not implemented"
	// Check if the signature is valid
	return false
}

// Create the public key format

// Check if they are invalid public key coordinates

// Verify the signature with the public key,
// then return true if it's valid, false otherwise

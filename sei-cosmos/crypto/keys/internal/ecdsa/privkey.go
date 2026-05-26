package ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"math/big"
)

// p256Order returns the curve order for the secp256r1 curve
// NOTE: this is specific to the secp256r1/P256 curve,
// and not taken from the domain params for the key itself
// (which would be a more generic approach for all EC).
var p256Order = elliptic.P256().Params().N

// p256HalfOrder returns half the curve order
// a bit shift of 1 to the right (Rsh) is equivalent
// to division by 2, only faster.
var p256HalfOrder = new(big.Int).Rsh(p256Order, 1)

// IsSNormalized returns true for the integer sigS if sigS falls in
// lower half of the curve order
func IsSNormalized(sigS *big.Int) bool { _ = "STUB: not implemented"; return false }

// NormalizeS will invert the s value if not already in the lower half
// of curve order value
func NormalizeS(sigS *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

// signatureRaw will serialize signature to R || S.
// R, S are padded to 32 bytes respectively.
// code roughly copied from secp256k1_nocgo.go
func signatureRaw(r *big.Int, s *big.Int) []byte { _ = "STUB: not implemented"; return nil }

// 0 pad the byte arrays from the left if they aren't big enough.

// GenPrivKey generates a new secp256r1 private key. It uses operating
// system randomness.
func GenPrivKey(curve elliptic.Curve) (PrivKey, error) {
	_ = "STUB: not implemented"
	return *new(PrivKey), nil
}

type PrivKey struct {
	ecdsa.PrivateKey
}

// PubKey returns ECDSA public key associated with this private key.
func (sk *PrivKey) PubKey() PubKey { _ = "STUB: not implemented"; return *new(PubKey) }

// Bytes serialize the private key using big-endian.
func (sk *PrivKey) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// Sign hashes and signs the message using ECDSA. Implements SDK
// PrivKey interface.
// NOTE: this now calls the ecdsa Sign function
// (not method!) directly as the s value of the signature is needed to
// low-s normalize the signature value
// See issue: https://github.com/cosmos/cosmos-sdk/issues/9723
// It then raw encodes the signature as two fixed width 32-byte values
// concatenated, reusing the code copied from secp256k1_nocgo.go
func (sk *PrivKey) Sign(msg []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// String returns a string representation of the public key based on the curveName.
func (sk *PrivKey) String(name string) string { _ = "STUB: not implemented"; return "" }

// MarshalTo implements proto.Marshaler interface.
func (sk *PrivKey) MarshalTo(dAtA []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Unmarshal implements proto.Marshaler interface.
func (sk *PrivKey) Unmarshal(bz []byte, curve elliptic.Curve, expectedSize int) error {
	_ = "STUB: not implemented"
	return nil
}

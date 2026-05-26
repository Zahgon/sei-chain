package ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"math/big"

	tmcrypto "github.com/sei-protocol/sei-chain/sei-tendermint/crypto"
)

// signatureFromBytes function roughly copied from secp256k1_nocgo.go
// Read Signature struct from R || S. Caller needs to ensure that
// len(sigStr) == 64.
func signatureFromBytes(sigStr []byte) *signature { _ = "STUB: not implemented"; return nil }

// signature holds the r and s values of an ECDSA signature.
type signature struct {
	R, S *big.Int
}

type PubKey struct {
	ecdsa.PublicKey

	// cache
	address tmcrypto.Address
}

// Address gets the address associated with a pubkey. If no address exists, it returns a newly created ADR-28 address
// for ECDSA keys.
// protoName is a concrete proto structure id.
func (pk *PubKey) Address(protoName string) tmcrypto.Address {
	_ = "STUB: not implemented"
	return *new(tmcrypto.Address)
}

// Bytes returns the byte representation of the public key using a compressed form
// specified in section 4.3.6 of ANSI X9.62 with first byte being the curve type.
func (pk *PubKey) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// VerifySignature checks if sig is a valid ECDSA signature for msg.
// This includes checking for low-s normalized signatures
// where the s integer component of the signature is in the
// lower half of the curve order
// 7/21/21 - expects raw encoded signature (fixed-width 64-bytes, R || S)
func (pk *PubKey) VerifySignature(msg []byte, sig []byte) bool {
	_ = "STUB: not implemented"

	// check length for raw signature
	// which is two 32-byte padded big.Ints
	// concatenated
	// NOT DER!
	return false
}

// String returns a string representation of the public key based on the curveName.
func (pk *PubKey) String(curveName string) string { _ = "STUB: not implemented"; return "" }

// **** Proto Marshaler ****

// MarshalTo implements proto.Marshaler interface.
func (pk *PubKey) MarshalTo(dAtA []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Unmarshal implements proto.Marshaler interface.
func (pk *PubKey) Unmarshal(bz []byte, curve elliptic.Curve, expectedSize int) error {
	_ = "STUB: not implemented"
	return nil
}

package ed25519

import (
	tmbytes "github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"
)

// Bytes converts the public key or signature to bytes.
func (k PublicKey) Bytes() []byte {
	_ = "STUB: not implemented"

	// Bytes converts the signature to bytes.
	return nil
}

func (s Signature) Bytes() []byte {
	_ = "STUB: not implemented"

	// PublicKeyFromBytes constructs a public key from bytes.
	return nil
}

func PublicKeyFromBytes(key []byte) (PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(PublicKey), nil
}

// SignatureFromBytes constructs a signature from bytes.
func SignatureFromBytes(sig []byte) (Signature, error) {
	_ = "STUB: not implemented"
	return *new(Signature), nil
}

// String returns a string representation.
func (k PublicKey) String() string { _ = "STUB: not implemented"; return "" }

// String returns a log-safe representation of the secret key.
func (k SecretKey) String() string { _ = "STUB: not implemented"; return "" }

// String returns a string representation.
func (s Signature) String() string { _ = "STUB: not implemented"; return "" }

// GoString returns a strings representation.
func (k PublicKey) GoString() string {
	_ = "STUB: not implemented"

	// GoString returns a log-safe representation of the secret key.
	return ""
}

func (k SecretKey) GoString() string {
	_ = "STUB: not implemented"

	// GoString returns a strings representation.
	return ""
}

func (s Signature) GoString() string {
	_ = "STUB: not implemented"

	// PublicKeyFromString constructs a public key from a string representation.
	return ""
}

func PublicKeyFromString(s string) (PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(PublicKey), nil
}

// Address is the SHA256-20 of the raw pubkey bytes.
func (k PublicKey) Address() tmbytes.HexBytes {
	_ = "STUB: not implemented"
	return *new(tmbytes.HexBytes)
}

package internal

import (
	"github.com/oasisprotocol/curve25519-voi/primitives/sr25519"

	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto"
)

const (
	// PubKeySize is the size of a sr25519 public key in bytes.
	PubKeySize = sr25519.PublicKeySize

	// SignatureSize is the size of a sr25519 signature in bytes.
	SignatureSize = sr25519.SignatureSize
)

// PubKey implements crypto.PubKey.
type PubKey []byte

// Address is the SHA256-20 of the raw pubkey bytes.
func (pubKey PubKey) Address() crypto.Address {
	_ = "STUB: not implemented"
	return *new(crypto.Address)
}

// Bytes returns the PubKey byte format.
func (pubKey PubKey) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (pubKey PubKey) Equals(other PubKey) bool { _ = "STUB: not implemented"; return false }

func (pubKey PubKey) VerifySignature(msg []byte, sigBytes []byte) bool {
	_ = "STUB: not implemented"
	return false
}

func (pubKey PubKey) String() string { _ = "STUB: not implemented"; return "" }

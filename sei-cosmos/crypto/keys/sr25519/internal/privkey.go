package internal

import (
	"io"

	"github.com/oasisprotocol/curve25519-voi/primitives/sr25519"
)

var signingCtx = sr25519.NewSigningContext([]byte{})

const KeyType = "sr25519"

// PrivKey implements crypto.PrivKey.
type PrivKey struct {
	msk sr25519.MiniSecretKey
	kp  *sr25519.KeyPair
}

func (privKey PrivKey) Type() string {
	_ = "STUB: not implemented"

	// Bytes returns the byte-encoded PrivKey.
	return ""
}

func (privKey PrivKey) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// Sign produces a signature on the provided message.
func (privKey PrivKey) Sign(msg []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// PubKey gets the corresponding public key from the private key.
//
// Panics if the private key is not initialized.
func (privKey PrivKey) PubKey() PubKey { _ = "STUB: not implemented"; return *new(PubKey) }

// Equals - you probably don't need to use this.
// Runs in constant time based on length of the keys.
func (privKey PrivKey) Equals(other PrivKey) bool { _ = "STUB: not implemented"; return false }

func (privKey PrivKey) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"

	// Handle uninitialized private keys gracefully.
	return nil, nil
}

func (privKey *PrivKey) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// GenPrivKey generates a new sr25519 private key.
// It uses OS randomness in conjunction with the current global random seed
// in tendermint/libs/common to generate the private key.
func GenPrivKey() PrivKey { _ = "STUB: not implemented"; return *new(PrivKey) }

func genPrivKey(rng io.Reader) PrivKey { _ = "STUB: not implemented"; return *new(PrivKey) }

// GenPrivKeyFromSecret hashes the secret with SHA2, and uses
// that 32 byte output to create the private key.
// NOTE: secret should be the output of a KDF like bcrypt,
// if it's derived from user input.
func GenPrivKeyFromSecret(secret []byte) PrivKey { _ = "STUB: not implemented"; return *new(PrivKey) }

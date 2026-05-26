package ed25519

import (
	"io"

	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
)

//-------------------------------------

const (
	PrivKeyName = "tendermint/PrivKeyEd25519"
	PubKeyName  = "tendermint/PubKeyEd25519"
	// PubKeySize is is the size, in bytes, of public keys as used in this package.
	PubKeySize = 32
	// PrivKeySize is the size, in bytes, of private keys as used in this package.
	PrivKeySize = 64
	// Size of an Edwards25519 signature. Namely the size of a compressed
	// Edwards25519 point, and a field element. Both of which are 32 bytes.
	SignatureSize = 64
	// SeedSize is the size, in bytes, of private key seeds. These are the
	// private key representations used by RFC 8032.
	SeedSize = 32

	keyType = "ed25519"
)

var _ cryptotypes.PrivKey = &PrivKey{}
var _ codec.AminoMarshaler = &PrivKey{}

// Bytes returns the privkey byte format.
func (privKey *PrivKey) Bytes() []byte {
	_ = "STUB: not implemented"

	// Sign produces a signature on the provided message.
	// This assumes the privkey is wellformed in the golang format.
	// The first 32 bytes should be random,
	// corresponding to the normal ed25519 private key.
	// The latter 32 bytes should be the compressed public key.
	// If these conditions aren't met, Sign will panic or produce an
	// incorrect signature.
	return nil
}

func (privKey *PrivKey) Sign(msg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PubKey gets the corresponding public key from the private key.
//
// Panics if the private key is not initialized.
func (privKey *PrivKey) PubKey() cryptotypes.PubKey {
	_ = "STUB: not implemented"
	// If the latter 32 bytes of the privkey are all zero, privkey is not
	// initialized.
	return *new(cryptotypes.PubKey)
}

// Equals - you probably don't need to use this.
// Runs in constant time based on length of the keys.
func (privKey *PrivKey) Equals(other cryptotypes.LedgerPrivKey) bool {
	_ = "STUB: not implemented"
	return false
}

func (privKey *PrivKey) Type() string {
	_ = "STUB: not implemented"

	// MarshalAmino overrides Amino binary marshalling.
	return ""
}

func (privKey PrivKey) MarshalAmino() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// UnmarshalAmino overrides Amino binary marshalling.
		nil
}

func (privKey *PrivKey) UnmarshalAmino(bz []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalAminoJSON overrides Amino JSON marshalling.
func (privKey PrivKey) MarshalAminoJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	// When we marshal to Amino JSON, we don't marshal the "key" field itself,
	// just its contents (i.e. the key bytes).
	return nil, nil
}

// UnmarshalAminoJSON overrides Amino JSON marshalling.
func (privKey *PrivKey) UnmarshalAminoJSON(bz []byte) error { _ = "STUB: not implemented"; return nil }

// GenPrivKey generates a new ed25519 private key. These ed25519 keys must not
// be used in SDK apps except in a tendermint validator context.
// It uses OS randomness in conjunction with the current global random seed
// in tendermint/libs/common to generate the private key.
func GenPrivKey() *PrivKey { _ = "STUB: not implemented"; return nil }

// genPrivKey generates a new ed25519 private key using the provided reader.
func genPrivKey(rand io.Reader) *PrivKey { _ = "STUB: not implemented"; return nil }

// GenPrivKeyFromSecret hashes the secret with SHA2, and uses
// that 32 byte output to create the private key.
// NOTE: ed25519 keys must not be used in SDK apps except in a tendermint validator context.
// NOTE: secret should be the output of a KDF like bcrypt,
// if it's derived from user input.
func GenPrivKeyFromSecret(secret []byte) *PrivKey { _ = "STUB: not implemented"; return nil }

// Not Ripemd160 because we want 32 bytes.

//-------------------------------------

var _ cryptotypes.PubKey = &PubKey{}
var _ codec.AminoMarshaler = &PubKey{}

const TruncatedSize = 20

// Address is the SHA256-20 of the raw pubkey bytes.
// It doesn't implement ADR-28 addresses and it must not be used
// in SDK except in a tendermint validator context.
func (pubKey *PubKey) Address() crypto.Address {
	_ = "STUB: not implemented"
	return *new(crypto.Address)
}

// For ADR-28 compatible address we would need to
// return address.Hash(proto.MessageName(pubKey), pubKey.Key)

// Bytes returns the PubKey byte format.
func (pubKey *PubKey) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (pubKey *PubKey) VerifySignature(msg []byte, sig []byte) bool {
	_ = "STUB: not implemented"
	// make sure we use the same algorithm to sign
	return false
}

// uses https://github.com/hdevalence/ed25519consensus.Verify to comply with zip215 verification rules

// String returns Hex representation of a pubkey with it's type
func (pubKey *PubKey) String() string { _ = "STUB: not implemented"; return "" }

func (pubKey *PubKey) Type() string { _ = "STUB: not implemented"; return "" }

func (pubKey *PubKey) Equals(other cryptotypes.PubKey) bool {
	_ = "STUB: not implemented"
	return false
}

// MarshalAmino overrides Amino binary marshalling.
func (pubKey PubKey) MarshalAmino() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// UnmarshalAmino overrides Amino binary marshalling.
		nil
}

func (pubKey *PubKey) UnmarshalAmino(bz []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalAminoJSON overrides Amino JSON marshalling.
func (pubKey PubKey) MarshalAminoJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	// When we marshal to Amino JSON, we don't marshal the "key" field itself,
	// just its contents (i.e. the key bytes).
	return nil, nil
}

// UnmarshalAminoJSON overrides Amino JSON marshalling.
func (pubKey *PubKey) UnmarshalAminoJSON(bz []byte) error { _ = "STUB: not implemented"; return nil }

package secp256k1

import (
	"io"
	"math/big"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto"
	//nolint:gosec,staticcheck // necessary for Bitcoin/Cosmos address derivation standard
)

var _ cryptotypes.PrivKey = &PrivKey{}
var _ codec.AminoMarshaler = &PrivKey{}

const (
	PrivKeySize = 32
	keyType     = "secp256k1"
	PrivKeyName = "tendermint/PrivKeySecp256k1"
	PubKeyName  = "tendermint/PubKeySecp256k1"
)

// Bytes returns the byte representation of the Private Key.
func (privKey *PrivKey) Bytes() []byte {
	_ = "STUB: not implemented"

	// PubKey performs the point-scalar multiplication from the privKey on the
	// generator point to get the pubkey.
	return nil
}

func (privKey *PrivKey) PubKey() cryptotypes.PubKey {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PubKey)
}

// Equals - you probably don't need to use this.
// Runs in constant time based on length of the
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

// GenPrivKey generates a new ECDSA private key on curve secp256k1 private key.
// It uses OS randomness to generate the private key.
func GenPrivKey() *PrivKey { _ = "STUB: not implemented"; return nil }

// genPrivKey generates a new secp256k1 private key using the provided reader.
func genPrivKey(rand io.Reader) []byte { _ = "STUB: not implemented"; return nil }

// break if we found a valid point (i.e. > 0 and < N == curverOrder)

var one = new(big.Int).SetInt64(1)

// GenPrivKeyFromSecret hashes the secret with SHA2, and uses
// that 32 byte output to create the private key.
//
// It makes sure the private key is a valid field element by setting:
//
// c = sha256(secret)
// k = (c mod (n − 1)) + 1, where n = curve order.
//
// NOTE: secret should be the output of a KDF like bcrypt,
// if it's derived from user input.
func GenPrivKeyFromSecret(secret []byte) *PrivKey { _ = "STUB: not implemented"; return nil }

// to guarantee that we have a valid field element, we use the approach of:
// "Suite B Implementer’s Guide to FIPS 186-3", A.2.1
// https://apps.nsa.gov/iaarchive/library/ia-guidance/ia-solutions-for-classified/algorithm-guidance/suite-b-implementers-guide-to-fips-186-3-ecdsa.cfm
// see also https://github.com/golang/go/blob/0380c9ad38843d523d9c9804fe300cb7edd7cd3c/src/crypto/ecdsa/ecdsa.go#L89-L101

// copy feB over to fixed 32 byte privKey32 and pad (if necessary)

//-------------------------------------

var _ cryptotypes.PubKey = &PubKey{}
var _ codec.AminoMarshaler = &PubKey{}

// PubKeySize comprises 32 bytes for one field element
// (the x-coordinate), plus one byte for the parity of the y-coordinate.
const PubKeySize = 33

// Address returns a Bitcoin style addresses: RIPEMD160(SHA256(pubkey))
func (pubKey *PubKey) Address() crypto.Address {
	_ = "STUB: not implemented"
	return *new(crypto.Address)
}

//nolint:gosec // RIPEMD-160 is required by the Bitcoin/Cosmos address derivation standard, not used for general-purpose hashing
// does not error

// Bytes returns the pubkey byte format.
func (pubKey *PubKey) Bytes() []byte { _ = "STUB: not implemented"; return nil }

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

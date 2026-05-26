package crypto

import (
	"crypto/sha256"

	ed25519 "github.com/sei-protocol/sei-chain/sei-tendermint/crypto/ed25519"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"
)

const (
	// HashSize is the size in bytes of an AddressHash.
	HashSize = sha256.Size

	// AddressSize is the size of a pubkey address.
	AddressSize = 20
)

// An address is a []byte, but hex-encoded even in JSON.
// []byte leaves us the option to change the address length.
// Use an alias so Unmarshal methods (with ptr receivers) are available too.
type Address = bytes.HexBytes

// AddressHash computes a truncated SHA-256 hash of bz for use as
// a peer address.
//
// See: https://docs.tendermint.com/master/spec/core/data_structures.html#address
func AddressHash(bz []byte) Address { _ = "STUB: not implemented"; return *new(Address) }

type Hash [sha256.Size]byte

func (h Hash) Bytes() bytes.HexBytes { _ = "STUB: not implemented"; return *new(bytes.HexBytes) }

// Checksum returns the SHA256 of the bz.
func Checksum(bz []byte) Hash { _ = "STUB: not implemented"; return *new(Hash) }

type PubKey = ed25519.PublicKey
type PrivKey = ed25519.SecretKey
type Sig = ed25519.Signature
type BatchVerifier = ed25519.BatchVerifier
type ErrBadSig = ed25519.ErrBadSig

func SigFromBytes(raw []byte) (Sig, error) { _ = "STUB: not implemented"; return *new(Sig), nil }

func NewBatchVerifier() *BatchVerifier { _ = "STUB: not implemented"; return nil }

package secp256r1

import (
	tmcrypto "github.com/sei-protocol/sei-chain/sei-tendermint/crypto"

	ecdsa "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/keys/internal/ecdsa"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
)

// String implements proto.Message interface.
func (m *PubKey) String() string { _ = "STUB: not implemented"; return "" }

// Bytes implements SDK PubKey interface.
func (m *PubKey) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// Equals implements SDK PubKey interface.
func (m *PubKey) Equals(other cryptotypes.PubKey) bool { _ = "STUB: not implemented"; return false }

// Address implements SDK PubKey interface.
func (m *PubKey) Address() tmcrypto.Address {
	_ = "STUB: not implemented"
	return *new(tmcrypto.Address)
}

// Type returns key type name. Implements SDK PubKey interface.
func (m *PubKey) Type() string {
	_ = "STUB: not implemented"

	// VerifySignature implements SDK PubKey interface.
	return ""
}

func (m *PubKey) VerifySignature(msg []byte, sig []byte) bool {
	_ = "STUB: not implemented"
	return false
}

type ecdsaPK struct {
	ecdsa.PubKey
}

// Size implements proto.Marshaler interface
func (pk *ecdsaPK) Size() int { _ = "STUB: not implemented"; return 0 }

// Unmarshal implements proto.Marshaler interface
func (pk *ecdsaPK) Unmarshal(bz []byte) error { _ = "STUB: not implemented"; return nil }

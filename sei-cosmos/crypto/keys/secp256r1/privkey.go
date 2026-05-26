package secp256r1

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/keys/internal/ecdsa"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
)

// GenPrivKey generates a new secp256r1 private key. It uses operating system randomness.
func GenPrivKey() (*PrivKey, error) { _ = "STUB: not implemented"; return nil, nil }

// PubKey implements SDK PrivKey interface.
func (m *PrivKey) PubKey() cryptotypes.PubKey {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PubKey)
}

// String implements SDK proto.Message interface.
func (m *PrivKey) String() string { _ = "STUB: not implemented"; return "" }

// Type returns key type name. Implements SDK PrivKey interface.
func (m *PrivKey) Type() string {
	_ = "STUB: not implemented"

	// Sign hashes and signs the message usign ECDSA. Implements sdk.PrivKey interface.
	return ""
}

func (m *PrivKey) Sign(msg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Bytes serialize the private key.
		nil
}

func (m *PrivKey) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// Equals implements SDK PrivKey interface.
func (m *PrivKey) Equals(other cryptotypes.LedgerPrivKey) bool {
	_ = "STUB: not implemented"
	return false
}

type ecdsaSK struct {
	ecdsa.PrivKey
}

// Size implements proto.Marshaler interface
func (sk *ecdsaSK) Size() int { _ = "STUB: not implemented"; return 0 }

// Unmarshal implements proto.Marshaler interface
func (sk *ecdsaSK) Unmarshal(bz []byte) error { _ = "STUB: not implemented"; return nil }

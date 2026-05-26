package sr25519

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/keys/sr25519/internal"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
)

const (
	PrivKeySize = 32
	PrivKeyName = "tendermint/PrivKeySr25519"
)

type PrivKey struct {
	internal.PrivKey
}

// type conversion
func (m *PrivKey) PubKey() cryptotypes.PubKey {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PubKey)
}

// type conversion
func (m *PrivKey) Equals(other cryptotypes.LedgerPrivKey) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *PrivKey) ProtoMessage() { _ = "STUB: not implemented"; return }

func (m *PrivKey) Reset() { _ = "STUB: not implemented"; return }

func (m *PrivKey) String() string { _ = "STUB: not implemented"; return "" }

func GenPrivKey() *PrivKey { _ = "STUB: not implemented"; return nil }

func GenPrivKeyFromSecret(secret []byte) *PrivKey { _ = "STUB: not implemented"; return nil }

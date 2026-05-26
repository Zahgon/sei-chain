package sr25519

import (
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto"
)

const PubKeyName = "tendermint/PubKeySr25519"

func (m *PubKey) Equals(other cryptotypes.PubKey) bool { _ = "STUB: not implemented"; return false }

func (m *PubKey) Address() crypto.Address { _ = "STUB: not implemented"; return *new(crypto.Address) }

func (m PubKey) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (m PubKey) String() string { _ = "STUB: not implemented"; return "" }

func (m PubKey) Type() string { _ = "STUB: not implemented"; return "" }

func (m PubKey) VerifySignature(msg []byte, sigBytes []byte) bool {
	_ = "STUB: not implemented"
	return false
}

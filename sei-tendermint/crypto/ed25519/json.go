package ed25519

import (
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/jsontypes"
)

const SecretKeyName = "tendermint/PrivKeyEd25519" //nolint:gosec
const PublicKeyName = "tendermint/PubKeyEd25519"
const KeyType = "ed25519"

func init() {
	jsontypes.MustRegister(PublicKey{})
	jsontypes.MustRegister(SecretKey{})
}

func (k SecretKey) TypeTag() string { _ = "STUB: not implemented"; return "" }
func (k SecretKey) Type() string    { _ = "STUB: not implemented"; return "" }

func (k PublicKey) TypeTag() string { _ = "STUB: not implemented"; return "" }
func (k PublicKey) Type() string {
	_ = "STUB: not implemented"

	// WARNING: this is very BAD that one can leak a secret by embedding
	// a private key in some struct and then calling json.Marshal on it.
	// TODO(gprusak): get rid of it.
	return ""
}

func (k SecretKey) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (k *SecretKey) UnmarshalJSON(j []byte) error { _ = "STUB: not implemented"; return nil }

func (k PublicKey) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (k *PublicKey) UnmarshalJSON(j []byte) error { _ = "STUB: not implemented"; return nil }

func (s Signature) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Signature) UnmarshalJSON(j []byte) error { _ = "STUB: not implemented"; return nil }

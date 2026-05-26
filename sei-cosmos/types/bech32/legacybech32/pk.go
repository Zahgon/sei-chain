// Deprecated: The module provides legacy bech32 functions which will be removed in a future
// release.
package legacybech32

import (
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
)

// TODO: when removing this package remove:
// + sdk:config.GetBech32AccountPubPrefix (and other related functions)
// + Bech32PrefixAccAddr and other related constants

// Deprecated: Bech32PubKeyType defines a string type alias for a Bech32 public key type.
type Bech32PubKeyType string

// Bech32 conversion constants
const (
	AccPK  Bech32PubKeyType = "accpub"
	ValPK  Bech32PubKeyType = "valpub"
	ConsPK Bech32PubKeyType = "conspub"
)

// Deprecated: MarshalPubKey returns a Bech32 encoded string containing the appropriate
// prefix based on the key type provided for a given PublicKey.
func MarshalPubKey(pkt Bech32PubKeyType, pubkey cryptotypes.PubKey) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Deprecated: MustMarshalPubKey calls MarshalPubKey and panics on error.
func MustMarshalPubKey(pkt Bech32PubKeyType, pubkey cryptotypes.PubKey) string {
	_ = "STUB: not implemented"
	return ""
}

func getPrefix(pkt Bech32PubKeyType) string { _ = "STUB: not implemented"; return "" }

// Deprecated: UnmarshalPubKey returns a PublicKey from a bech32-encoded PublicKey with
// a given key type.
func UnmarshalPubKey(pkt Bech32PubKeyType, pubkeyStr string) (cryptotypes.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PubKey), nil
}

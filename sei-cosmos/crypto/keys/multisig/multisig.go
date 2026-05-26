package multisig

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	multisigtypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types/multisig"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/tx/signing"
)

var _ multisigtypes.PubKey = &LegacyAminoPubKey{}
var _ types.UnpackInterfacesMessage = &LegacyAminoPubKey{}

// NewLegacyAminoPubKey returns a new LegacyAminoPubKey.
// Multisig can be constructed with multiple same keys - it will increase the power of
// the owner of that key (he will still need to add multiple signatures in the right order).
// Panics if len(pubKeys) < k or 0 >= k.
func NewLegacyAminoPubKey(threshold int, pubKeys []cryptotypes.PubKey) *LegacyAminoPubKey {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // threshold is validated positive above and practically small

// Address implements cryptotypes.PubKey Address method
func (m *LegacyAminoPubKey) Address() cryptotypes.Address {
	_ = "STUB: not implemented"
	return *new(cryptotypes.Address)
}

// Bytes returns the proto encoded version of the LegacyAminoPubKey
func (m *LegacyAminoPubKey) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// VerifyMultisignature implements the multisigtypes.PubKey VerifyMultisignature method.
// The signatures must be added in an order corresponding to the public keys order in
// LegacyAminoPubKey. It's OK to have multiple same keys in the multisig - it will increase
// the power of the owner of that key - in that case the signer will still need to append
// multiple same signatures in the right order.
func (m *LegacyAminoPubKey) VerifyMultisignature(getSignBytes multisigtypes.GetSignBytesFunc, sig *signing.MultiSignatureData) error {
	_ = "STUB: not implemented"
	return nil
}

// ensure bit array is the correct size

// ensure size of signature list

// ensure at least k signatures are set

// index in the list of signatures which we are concerned with.

// VerifySignature implements cryptotypes.PubKey VerifySignature method,
// it panics because it can't handle MultiSignatureData
// cf. https://github.com/cosmos/cosmos-sdk/issues/7109#issuecomment-686329936
func (m *LegacyAminoPubKey) VerifySignature(msg []byte, sig []byte) bool {
	_ = "STUB: not implemented"
	return false

	// GetPubKeys implements the PubKey.GetPubKeys method
}

func (m *LegacyAminoPubKey) GetPubKeys() []cryptotypes.PubKey {
	_ = "STUB: not implemented"
	return nil
}

// Equals returns true if m and other both have the same number of keys, and
// all constituent keys are the same, and in the same order.
func (m *LegacyAminoPubKey) Equals(key cryptotypes.PubKey) bool {
	_ = "STUB: not implemented"
	return false
}

// GetThreshold implements the PubKey.GetThreshold method
func (m *LegacyAminoPubKey) GetThreshold() uint { _ = "STUB: not implemented"; return 0 }

// Type returns multisig type
func (m *LegacyAminoPubKey) Type() string { _ = "STUB: not implemented"; return "" }

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (m *LegacyAminoPubKey) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

func packPubKeys(pubKeys []cryptotypes.PubKey) ([]*types.Any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

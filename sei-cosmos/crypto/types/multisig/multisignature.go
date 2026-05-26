package multisig

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/tx/signing"
)

// AminoMultisignature is used to represent amino multi-signatures for StdTx's.
// It is assumed that all signatures were made with SIGN_MODE_LEGACY_AMINO_JSON.
// Sigs is a list of signatures, sorted by corresponding index.
type AminoMultisignature struct {
	BitArray *types.CompactBitArray
	Sigs     [][]byte
}

// NewMultisig returns a new MultiSignatureData
func NewMultisig(n int) *signing.MultiSignatureData { _ = "STUB: not implemented"; return nil }

// GetIndex returns the index of pk in keys. Returns -1 if not found
func getIndex(pk types.PubKey, keys []types.PubKey) int { _ = "STUB: not implemented"; return 0 }

// AddSignature adds a signature to the multisig, at the corresponding index. The index must
// represent the pubkey index in the LegacyAmingPubKey structure, which verifies this signature.
// If the signature already exists, replace it.
func AddSignature(mSig *signing.MultiSignatureData, sig signing.SignatureData, index int) {
	_ = "STUB: not implemented"
	return
}

// Signature already exists, just replace the value there

// Optimization if the index is the greatest index

// Expand slice by one with a dummy element, move all elements after i
// over by one, then place the new signature in that gap.

// AddSignatureFromPubKey adds a signature to the multisig, at the index in
// keys corresponding to the provided pubkey.
func AddSignatureFromPubKey(mSig *signing.MultiSignatureData, sig signing.SignatureData, pubkey types.PubKey, keys []types.PubKey) error {
	_ = "STUB: not implemented"
	return nil
}

func AddSignatureV2(mSig *signing.MultiSignatureData, sig signing.SignatureV2, keys []types.PubKey) error {
	_ = "STUB: not implemented"
	return nil
}

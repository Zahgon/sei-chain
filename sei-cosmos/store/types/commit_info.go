package types

import (
	tmcrypto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/crypto"
)

// GetHash returns the GetHash from the CommitID.
// This is used in CommitInfo.Hash()
//
// When we commit to this in a merkle proof, we create a map of storeInfo.Name -> storeInfo.GetHash()
// and build a merkle proof from that.
// This is then chained with the substore proof, so we prove the root hash from the substore before this
// and need to pass that (unmodified) as the leaf value of the multistore proof.
func (si StoreInfo) GetHash() []byte { _ = "STUB: not implemented"; return nil }

func (ci CommitInfo) toMap() map[string][]byte { _ = "STUB: not implemented"; return nil }

// Hash returns the simple merkle root hash of the stores sorted by name.
func (ci CommitInfo) Hash() []byte {
	_ = "STUB: not implemented"
	// we need a special case for empty set, as SimpleProofsFromMap requires at least one entry
	return nil
}

func (ci CommitInfo) ProofOp(storeName string) tmcrypto.ProofOp {
	_ = "STUB: not implemented"
	return *new(tmcrypto.ProofOp)
}

// convert merkle.SimpleProof to CommitmentProof

func (ci CommitInfo) CommitID() CommitID { _ = "STUB: not implemented"; return *new(CommitID) }

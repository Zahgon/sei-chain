package types

import (
	crypto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/crypto"
)

// ConvertProofs converts crypto.ProofOps into MerkleProof
func ConvertProofs(tmProof *crypto.ProofOps) (MerkleProof, error) {
	_ = "STUB: not implemented"
	return *new(MerkleProof), nil
}

// Unmarshal all proof ops to CommitmentProof

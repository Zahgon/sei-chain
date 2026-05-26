/*
*
This implements the client side functions as specified in
https://github.com/cosmos/ics/tree/master/spec/ics-023-vector-commitments

In particular:

	// Assumes ExistenceProof
	type verifyMembership = (root: CommitmentRoot, proof: CommitmentProof, key: Key, value: Value) => boolean

	// Assumes NonExistenceProof
	type verifyNonMembership = (root: CommitmentRoot, proof: CommitmentProof, key: Key) => boolean

	// Assumes BatchProof - required ExistenceProofs may be a subset of all items proven
	type batchVerifyMembership = (root: CommitmentRoot, proof: CommitmentProof, items: Map<Key, Value>) => boolean

	// Assumes BatchProof - required NonExistenceProofs may be a subset of all items proven
	type batchVerifyNonMembership = (root: CommitmentRoot, proof: CommitmentProof, keys: Set<Key>) => boolean

We make an adjustment to accept a Spec to ensure the provided proof is in the format of the expected merkle store.
This can avoid an range of attacks on fake preimages, as we need to be careful on how to map key, value -> leaf
and determine neighbors
*/
package ics23

// CommitmentRoot is a byte slice that represents the merkle root of a tree that can be used to validate proofs
type CommitmentRoot []byte

// VerifyMembership returns true iff
// proof is (contains) an ExistenceProof for the given key and value AND
// calculating the root for the ExistenceProof matches the provided CommitmentRoot
func VerifyMembership(spec *ProofSpec, root CommitmentRoot, proof *CommitmentProof, key []byte, value []byte) bool {
	_ = "STUB: not implemented"
	// decompress it before running code (no-op if not compressed)
	return false
}

// VerifyNonMembership returns true iff
// proof is (contains) a NonExistenceProof
// both left and right sub-proofs are valid existence proofs (see above) or nil
// left and right proofs are neighbors (or left/right most if one is nil)
// provided key is between the keys of the two proofs
func VerifyNonMembership(spec *ProofSpec, root CommitmentRoot, proof *CommitmentProof, key []byte) bool {
	_ = "STUB: not implemented"
	// decompress it before running code (no-op if not compressed)
	return false
}

// BatchVerifyMembership will ensure all items are also proven by the CommitmentProof (which should be a BatchProof,
// unless there is one item, when a ExistenceProof may work)
func BatchVerifyMembership(spec *ProofSpec, root CommitmentRoot, proof *CommitmentProof, items map[string][]byte) bool {
	_ = "STUB: not implemented"
	// decompress it before running code (no-op if not compressed) - once for batch
	return false
}

// BatchVerifyNonMembership will ensure all items are also proven to not be in the Commitment by the CommitmentProof
// (which should be a BatchProof, unless there is one item, when a NonExistenceProof may work)
func BatchVerifyNonMembership(spec *ProofSpec, root CommitmentRoot, proof *CommitmentProof, keys [][]byte) bool {
	_ = "STUB: not implemented"
	// decompress it before running code (no-op if not compressed) - once for batch
	return false
}

// CombineProofs takes a number of commitment proofs (simple or batch) and
// converts them into a batch and compresses them.
//
// This is designed for proof generation libraries to create efficient batches
func CombineProofs(proofs []*CommitmentProof) (*CommitmentProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getExistProofForKey(proof *CommitmentProof, key []byte) *ExistenceProof {
	_ = "STUB: not implemented"
	return nil
}

func getNonExistProofForKey(proof *CommitmentProof, key []byte) *NonExistenceProof {
	_ = "STUB: not implemented"
	return nil
}

func isLeft(left *ExistenceProof, key []byte) bool { _ = "STUB: not implemented"; return false }

func isRight(right *ExistenceProof, key []byte) bool { _ = "STUB: not implemented"; return false }

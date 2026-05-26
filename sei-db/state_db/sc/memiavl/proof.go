package memiavl

import (
	ics23 "github.com/confio/ics23/go"
)

// ProofInnerNode represents an inner node in an IAVL proof path.
type ProofInnerNode struct {
	Height  int8   `json:"height"`
	Size    int64  `json:"size"`
	Version int64  `json:"version"`
	Left    []byte `json:"left"`
	Right   []byte `json:"right"`
}

// PathToLeaf represents an inner path to a leaf node in an IAVL tree.
type PathToLeaf []ProofInnerNode

/*
GetMembershipProof will produce a CommitmentProof that the given key (and queries value) exists in the iavl tree.
If the key doesn't exist in the tree, this will return an error.
*/
func (t *Tree) GetMembershipProof(key []byte) (*ics23.CommitmentProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// VerifyMembership returns true iff proof is an ExistenceProof for the given key.
func (t *Tree) VerifyMembership(proof *ics23.CommitmentProof, key []byte) bool {
	_ = "STUB: not implemented"
	return false
}

/*
GetNonMembershipProof will produce a CommitmentProof that the given key doesn't exist in the iavl tree.
If the key exists in the tree, this will return an error.
*/
func (t *Tree) GetNonMembershipProof(key []byte) (*ics23.CommitmentProof, error) {
	_ = "STUB: not implemented"
	// idx is one node right of what we want....
	return nil, nil
}

// this will be nil if nothing right of the queried key

// VerifyNonMembership returns true iff proof is a NonExistenceProof for the given key.
func (t *Tree) VerifyNonMembership(proof *ics23.CommitmentProof, key []byte) bool {
	_ = "STUB: not implemented"
	return false
}

// createExistenceProof will get the proof from the tree and convert the proof into a valid
// existence proof, if that's what it is.
func (t *Tree) createExistenceProof(key []byte) (*ics23.ExistenceProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertLeafOp(version int64) *ics23.LeafOp {
	_ = "STUB: not implemented"
	// this is adapted from iavl/proof.go:proofLeafNode.Hash()
	return nil
}

// we cannot get the proofInnerNode type, so we need to do the whole path in one function
func convertInnerOps(path PathToLeaf) []*ics23.InnerOp { _ = "STUB: not implemented"; return nil }

// lengthByte is the length prefix prepended to each of the sha256 sub-hashes

// we need to go in reverse order, iavl starts from root to leaf,
// we want to go up from the leaf to the root

// this is adapted from iavl/proof.go:proofInnerNode.Hash()

// length prefixed left side

// prepend the length prefix for child

// prepend the length prefix for child

// length-prefixed right side

func convertVarIntToBytes(orig int64) []byte { _ = "STUB: not implemented"; return nil }

func pathToLeaf(node Node, key []byte) (PathToLeaf, Node, error) {
	_ = "STUB: not implemented"
	return *new(PathToLeaf), *new(Node), nil
}

// left side

// right side

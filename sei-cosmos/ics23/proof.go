package ics23

// IavlSpec constrains the format from proofs-iavl (iavl merkle proofs)
var IavlSpec = &ProofSpec{
	LeafSpec: &LeafOp{
		Prefix:       []byte{0},
		PrehashKey:   HashOp_NO_HASH,
		Hash:         HashOp_SHA256,
		PrehashValue: HashOp_SHA256,
		Length:       LengthOp_VAR_PROTO,
	},
	InnerSpec: &InnerSpec{
		ChildOrder:      []int32{0, 1},
		MinPrefixLength: 4,
		MaxPrefixLength: 12,
		ChildSize:       33, // (with length byte)
		EmptyChild:      nil,
		Hash:            HashOp_SHA256,
	},
}

// SmtSpec constrains the format for SMT proofs (as implemented by github.com/celestiaorg/smt)
var SmtSpec = &ProofSpec{
	LeafSpec: &LeafOp{
		Hash:         HashOp_SHA256,
		PrehashKey:   HashOp_NO_HASH,
		PrehashValue: HashOp_SHA256,
		Length:       LengthOp_NO_PREFIX,
		Prefix:       []byte{0},
	},
	InnerSpec: &InnerSpec{
		ChildOrder:      []int32{0, 1},
		ChildSize:       32,
		MinPrefixLength: 1,
		MaxPrefixLength: 1,
		EmptyChild:      make([]byte, 32),
		Hash:            HashOp_SHA256,
	},
	MaxDepth: 256,
}

func encodeVarintProto(l int) []byte {
	_ = "STUB: not implemented"
	// avoid multiple allocs for normal case
	return nil
}

// Calculate determines the root hash that matches a given Commitment proof
// by type switching and calculating root based on proof type
// NOTE: Calculate will return the first calculated root in the proof,
// you must validate that all other embedded ExistenceProofs commit to the same root.
// This can be done with the Verify method
func (p *CommitmentProof) Calculate() (CommitmentRoot, error) {
	_ = "STUB: not implemented"
	return *new(CommitmentRoot), nil
}

// Verify does all checks to ensure this proof proves this key, value -> root
// and matches the spec.
func (p *ExistenceProof) Verify(spec *ProofSpec, root CommitmentRoot, key []byte, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Calculate determines the root hash that matches the given proof.
// You must validate the result is what you have in a header.
// Returns error if the calculations cannot be performed.
func (p *ExistenceProof) Calculate() (CommitmentRoot, error) {
	_ = "STUB: not implemented"
	return *new(CommitmentRoot), nil
}

func (p *ExistenceProof) calculate(spec *ProofSpec) (CommitmentRoot, error) {
	_ = "STUB: not implemented"
	return *new(CommitmentRoot), nil
}

// leaf step takes the key and value as input

// the rest just take the output of the last step (reducing it)

func decompressEntry(entry *CompressedBatchEntry, lookup []*InnerOp) *BatchEntry {
	_ = "STUB: not implemented"
	return nil
}

// Calculate determines the root hash that matches the given nonexistence rpoog.
// You must validate the result is what you have in a header.
// Returns error if the calculations cannot be performed.
func (p *NonExistenceProof) Calculate() (CommitmentRoot, error) {
	_ = "STUB: not implemented"
	// A Nonexist proof may have left or right proof nil
	return *new(CommitmentRoot), nil
}

// CheckAgainstSpec will verify the leaf and all path steps are in the format defined in spec
func (p *ExistenceProof) CheckAgainstSpec(spec *ProofSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// Verify does all checks to ensure the proof has valid non-existence proofs,
// and they ensure the given key is not in the CommitmentState
func (p *NonExistenceProof) Verify(spec *ProofSpec, root CommitmentRoot, key []byte) error {
	_ = "STUB: not implemented"
	// ensure the existence proofs are valid
	return nil
}

// If both proofs are missing, this is not a valid proof

// Ensure in valid range

// in the middle

// IsLeftMost returns true if this is the left-most path in the tree, excluding placeholder (empty child) nodes
func IsLeftMost(spec *InnerSpec, path []*InnerOp) bool { _ = "STUB: not implemented"; return false }

// ensure every step has a prefix and suffix defined to be leftmost, unless it is a placeholder node

// IsRightMost returns true if this is the left-most path in the tree, excluding placeholder (empty child) nodes
func IsRightMost(spec *InnerSpec, path []*InnerOp) bool { _ = "STUB: not implemented"; return false }

// ensure every step has a prefix and suffix defined to be rightmost, unless it is a placeholder node

// IsLeftNeighbor returns true if `right` is the next possible path right of `left`
//
//	Find the common suffix from the Left.Path and Right.Path and remove it. We have LPath and RPath now, which must be neighbors.
//	Validate that LPath[len-1] is the left neighbor of RPath[len-1]
//	For step in LPath[0..len-1], validate step is right-most node
//	For step in RPath[0..len-1], validate step is left-most node
func IsLeftNeighbor(spec *InnerSpec, left []*InnerOp, right []*InnerOp) bool {
	_ = "STUB: not implemented"
	// count common tail (from end, near root)
	return false
}

// now topleft and topright are the first divergent nodes
// make sure they are left and right of each other

// left and right are remaining children below the split,
// ensure left child is the rightmost path, and visa versa

// isLeftStep assumes left and right have common parents
// checks if left is exactly one slot to the left of right
func isLeftStep(spec *InnerSpec, left *InnerOp, right *InnerOp) bool {
	_ = "STUB: not implemented"
	return false
}

// TODO: is it possible there are empty (nil) children???

// checks if an op has the expected padding
func hasPadding(op *InnerOp, minPrefix, maxPrefix, suffix int) bool {
	_ = "STUB: not implemented"
	return false
}

// getPadding determines prefix and suffix with the given spec and position in the tree
func getPadding(spec *InnerSpec, branch int32) (minPrefix, maxPrefix, suffix int) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

// count how many children are in the prefix

// count how many children are in the suffix

// leftBranchesAreEmpty returns true if the padding bytes correspond to all empty siblings
// on the left side of a branch, ie. it's a valid placeholder on a leftmost path
func leftBranchesAreEmpty(spec *InnerSpec, op *InnerOp) bool {
	_ = "STUB: not implemented"
	return false
}

// count branches to left of this

// compare prefix with the expected number of empty branches

// rightBranchesAreEmpty returns true if the padding bytes correspond to all empty siblings
// on the right side of a branch, ie. it's a valid placeholder on a rightmost path
func rightBranchesAreEmpty(spec *InnerSpec, op *InnerOp) bool {
	_ = "STUB: not implemented"
	return false
}

// count branches to right of this one

// compare suffix with the expected number of empty branches

// sanity check

// getPosition checks where the branch is in the order and returns
// the index of this branch
func getPosition(order []int32, branch int32) int { _ = "STUB: not implemented"; return 0 }

// This will look at the proof and determine which order it is...
// So we can see if it is branch 0, 1, 2 etc... to determine neighbors
func orderFromPadding(spec *InnerSpec, inner *InnerOp) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// over-declares equality, which we cosnider fine for now.
func (p *ProofSpec) SpecEquals(spec *ProofSpec) bool { _ = "STUB: not implemented"; return false }

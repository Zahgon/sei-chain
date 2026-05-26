package types

import (
	ics23 "github.com/confio/ics23/go"
	tmcrypto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/crypto"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

// var representing the proofspecs for a SDK chain
var sdkSpecs = []*ics23.ProofSpec{ics23.IavlSpec, ics23.TendermintSpec}

// ICS 023 Merkle Types Implementation
//
// This file defines Merkle commitment types that implements ICS 023.

// Merkle proof implementation of the Proof interface
// Applied on SDK-based IBC implementation
var _ exported.Root = (*MerkleRoot)(nil)

// GetSDKSpecs is a getter function for the proofspecs of an sdk chain
func GetSDKSpecs() []*ics23.ProofSpec {
	_ = "STUB: not implemented"

	// NewMerkleRoot constructs a new MerkleRoot
	return nil
}

func NewMerkleRoot(hash []byte) MerkleRoot { _ = "STUB: not implemented"; return *new(MerkleRoot) }

// GetHash implements RootI interface
func (mr MerkleRoot) GetHash() []byte {
	_ = "STUB: not implemented"

	// Empty returns true if the root is empty
	return nil
}

func (mr MerkleRoot) Empty() bool { _ = "STUB: not implemented"; return false }

var _ exported.Prefix = (*MerklePrefix)(nil)

// NewMerklePrefix constructs new MerklePrefix instance
func NewMerklePrefix(keyPrefix []byte) MerklePrefix {
	_ = "STUB: not implemented"
	return *new(MerklePrefix)
}

// Bytes returns the key prefix bytes
func (mp MerklePrefix) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// Empty returns true if the prefix is empty
func (mp MerklePrefix) Empty() bool { _ = "STUB: not implemented"; return false }

var _ exported.Path = (*MerklePath)(nil)

// NewMerklePath creates a new MerklePath instance
// The keys must be passed in from root-to-leaf order
func NewMerklePath(keyPath ...string) MerklePath {
	_ = "STUB: not implemented"
	return *new(MerklePath)
}

// String implements fmt.Stringer.
// This represents the path in the same way the tendermint KeyPath will
// represent a key path. The backslashes partition the key path into
// the respective stores they belong to.
func (mp MerklePath) String() string { _ = "STUB: not implemented"; return "" }

// Pretty returns the unescaped path of the URL string.
// This function will unescape any backslash within a particular store key.
// This makes the keypath more human-readable while removing information
// about the exact partitions in the key path.
func (mp MerklePath) Pretty() string { _ = "STUB: not implemented"; return "" }

// GetKey will return a byte representation of the key
// after URL escaping the key element
func (mp MerklePath) GetKey(i uint64) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Empty returns true if the path is empty
func (mp MerklePath) Empty() bool { _ = "STUB: not implemented"; return false }

// ApplyPrefix constructs a new commitment path from the arguments. It prepends the prefix key
// with the given path.
func ApplyPrefix(prefix exported.Prefix, path MerklePath) (MerklePath, error) {
	_ = "STUB: not implemented"
	return *new(MerklePath), nil
}

var _ exported.Proof = (*MerkleProof)(nil)

// VerifyMembership verifies the membership pf a merkle proof against the given root, path, and value.
func (proof MerkleProof) VerifyMembership(specs []*ics23.ProofSpec, root exported.Root, path exported.Path, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyMembership specific argument validation

// Since every proof in chain is a membership proof we can use verifyChainedMembershipProof from index 0
// to validate entire proof

// VerifyNonMembership verifies the absence of a merkle proof against the given root and path.
// VerifyNonMembership verifies a chained proof where the absence of a given path is proven
// at the lowest subtree and then each subtree's inclusion is proved up to the final root.
func (proof MerkleProof) VerifyNonMembership(specs []*ics23.ProofSpec, root exported.Root, path exported.Path) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyNonMembership specific argument validation

// VerifyNonMembership will verify the absence of key in lowest subtree, and then chain inclusion proofs
// of all subroots up to final root

// #nosec G115 -- length is checked above to be non-empty

// Verify chained membership proof starting from index 1 with value = subroot

// BatchVerifyMembership verifies a group of key value pairs against the given root
// NOTE: Currently left unimplemented as it is unused
func (proof MerkleProof) BatchVerifyMembership(specs []*ics23.ProofSpec, root exported.Root, path exported.Path, items map[string][]byte) error {
	_ = "STUB: not implemented"
	return nil
}

// BatchVerifyNonMembership verifies absence of a group of keys against the given root
// NOTE: Currently left unimplemented as it is unused
func (proof MerkleProof) BatchVerifyNonMembership(specs []*ics23.ProofSpec, root exported.Root, path exported.Path, items [][]byte) error {
	_ = "STUB: not implemented"
	return nil
}

// verifyChainedMembershipProof takes a list of proofs and specs and verifies each proof sequentially ensuring that the value is committed to
// by first proof and each subsequent subroot is committed to by the next subroot and checking that the final calculated root is equal to the given roothash.
// The proofs and specs are passed in from lowest subtree to the highest subtree, but the keys are passed in from highest subtree to lowest.
// The index specifies what index to start chaining the membership proofs, this is useful since the lowest proof may not be a membership proof, thus we
// will want to start the membership proof chaining from index 1 with value being the lowest subroot
func verifyChainedMembershipProof(root []byte, specs []*ics23.ProofSpec, proofs []*ics23.CommitmentProof, keys MerklePath, value []byte, index int) error {
	_ = "STUB: not implemented"
	return nil
}

// Initialize subroot to value since the proofs list may be empty.
// This may happen if this call is verifying intermediate proofs after the lowest proof has been executed.
// In this case, there may be no intermediate proofs to verify and we just check that lowest proof root equals final root

// Since keys are passed in from highest to lowest, we must grab their indices in reverse order
// from the proofs and specs which are lowest to highest

// #nosec G115 -- keyIndex is bounds checked above

// verify membership of the proof at this index with appropriate key and value

// Set value to subroot so that we verify next proof in chain commits to this subroot

// Check that chained proof root equals passed-in root

// blankMerkleProof and blankProofOps will be used to compare against their zero values,
// and are declared as globals to avoid having to unnecessarily re-allocate on every comparison.
var (
	blankMerkleProof = &MerkleProof{}
	blankProofOps    = &tmcrypto.ProofOps{}
)

// Empty returns true if the root is empty
func (proof *MerkleProof) Empty() bool { _ = "STUB: not implemented"; return false }

// ValidateBasic checks if the proof is empty.
func (proof MerkleProof) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// validateVerificationArgs verifies the proof arguments are valid
func (proof MerkleProof) validateVerificationArgs(specs []*ics23.ProofSpec, root exported.Root) error {
	_ = "STUB: not implemented"
	return nil
}

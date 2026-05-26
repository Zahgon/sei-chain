package memiavl

import (
	"io"
)

// Node interface encapsulate the interface of both PersistedNode and MemNode.
type Node interface {
	Height() uint8
	IsLeaf() bool
	Size() int64
	Version() uint32
	Key() []byte
	Value() []byte
	Left() Node
	Right() Node
	Hash() []byte

	// SafeHash returns byte slice that's safe to retain
	SafeHash() []byte

	// PersistedNode clone a new node, MemNode modify in place
	Mutate(version, cowVersion uint32) *MemNode

	// Get query the value for a key, it's put into interface because a specialized implementation is more efficient.
	Get(key []byte) ([]byte, uint32)
	GetByIndex(uint32) ([]byte, []byte)
}

// setRecursive do set operation.
// it always do modification and return new `MemNode`, even if the value is the same.
// also returns if it's an update or insertion, if updated, the tree height and balance is not changed.
func setRecursive(node Node, key, value []byte, version, cowVersion uint32) (*MemNode, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// removeRecursive returns:
// - (nil, origNode, nil) -> nothing changed in subtree
// - (value, nil, newKey) -> leaf node is removed
// - (value, new node, newKey) -> subtree changed
func removeRecursive(node Node, key []byte, version, cowVersion uint32) ([]byte, Node, []byte) {
	_ = "STUB: not implemented"
	return nil, *new(Node), nil
}

// Writes the node's hash to the given `io.Writer`. This function recursively calls
// children to update hashes.
func writeHashBytes(node Node, w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Key is not written for inner nodes, unlike writeBytes.

// Indirection needed to provide proofs without values.
// (e.g. ProofLeafNode.ValueHash)

// HashNode computes the hash of the node.
func HashNode(node Node) []byte { _ = "STUB: not implemented"; return nil }

// VerifyHash compare node's cached hash with computed one
func VerifyHash(node Node) bool { _ = "STUB: not implemented"; return false }

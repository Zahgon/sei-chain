package memiavl

import (
	"io"
	"sync/atomic"
)

type MemNode struct {
	height  uint8
	version uint32
	size    int64
	key     []byte
	value   []byte
	left    Node
	right   Node
	hash    []byte
}

var _ Node = (*MemNode)(nil)
var (
	TotalMemNodeSize  = atomic.Int64{}
	TotalNumOfMemNode = atomic.Int64{}
)

func newBranchNode(
	height uint8,
	size int64,
	version uint32,
	key []byte,
	left Node,
	right Node) *MemNode {
	_ = "STUB: not implemented"
	return nil
}

func newLeafNode(key, value []byte, version uint32) *MemNode { _ = "STUB: not implemented"; return nil }

func IncrementMemNodeSize(node *MemNode) {
	_ = "STUB: not implemented"

	// struct itself (includes slice headers and interface headers)
	return
}

// backing arrays for the slices (bytes)

func (node *MemNode) Height() uint8 { _ = "STUB: not implemented"; return 0 }

func (node *MemNode) IsLeaf() bool { _ = "STUB: not implemented"; return false }

func (node *MemNode) Size() int64 { _ = "STUB: not implemented"; return 0 }

func (node *MemNode) Version() uint32 { _ = "STUB: not implemented"; return 0 }

func (node *MemNode) Key() []byte { _ = "STUB: not implemented"; return nil }

func (node *MemNode) Value() []byte { _ = "STUB: not implemented"; return nil }

func (node *MemNode) Left() Node { _ = "STUB: not implemented"; return *new(Node) }

func (node *MemNode) Right() Node {
	_ = "STUB: not implemented"

	// Mutate clones the node if it's version is smaller than or equal to cowVersion, otherwise modify in-place
	return *new(Node)
}

func (node *MemNode) Mutate(version, cowVersion uint32) *MemNode {
	_ = "STUB: not implemented"
	return nil
}

func (node *MemNode) SafeHash() []byte {
	_ = "STUB: not implemented"

	// Computes the hash of the node without computing its descendants. Must be
	// called on nodes which have descendant node hashes already computed.
	return nil
}

func (node *MemNode) Hash() []byte { _ = "STUB: not implemented"; return nil }

func (node *MemNode) updateHeightSize() { _ = "STUB: not implemented"; return }

func (node *MemNode) calcBalance() int { _ = "STUB: not implemented"; return 0 }

func calcBalance(node Node) int { _ = "STUB: not implemented"; return 0 }

// Invariant: node is returned by `Mutate(version)`.
//
//	   S               L
//	  / \      =>     / \
//	 L                   S
//	/ \                 / \
//	  LR               LR
func (node *MemNode) rotateRight(version, cowVersion uint32) *MemNode {
	_ = "STUB: not implemented"
	return nil
}

// Invariant: node is returned by `Mutate(version, cowVersion)`.
//
//	 S              R
//	/ \     =>     / \
//	    R         S
//	   / \       / \
//	 RL             RL
func (node *MemNode) rotateLeft(version, cowVersion uint32) *MemNode {
	_ = "STUB: not implemented"
	return nil
}

// Invariant: node is returned by `Mutate(version, cowVersion)`.
func (node *MemNode) reBalance(version, cowVersion uint32) *MemNode {
	_ = "STUB: not implemented"
	return nil
}

// left left

// left right

// right right

// right left

// nothing changed

func (node *MemNode) Get(key []byte) ([]byte, uint32) { _ = "STUB: not implemented"; return nil, 0 }

func (node *MemNode) GetByIndex(index uint32) ([]byte, []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EncodeBytes writes a varint length-prefixed byte slice to the writer,
// it's used for hash computation, must be compactible with the official IAVL implementation.
func EncodeBytes(w io.Writer, bz []byte) error { _ = "STUB: not implemented"; return nil }

func maxUInt8(a, b uint8) uint8 { _ = "STUB: not implemented"; return 0 }

package memiavl

import (
	"crypto/sha256"
)

const (
	OffsetHeight   = 0
	OffsetPreTrees = OffsetHeight + 1
	OffsetVersion  = OffsetHeight + 4
	OffsetSize     = OffsetVersion + 4
	OffsetKeyLeaf  = OffsetSize + 4

	OffsetHash          = OffsetKeyLeaf + 4
	SizeHash            = sha256.Size
	SizeNodeWithoutHash = OffsetHash
	SizeNode            = SizeNodeWithoutHash + SizeHash

	OffsetLeafVersion   = 0
	OffsetLeafKeyLen    = OffsetLeafVersion + 4
	OffsetLeafKeyOffset = OffsetLeafKeyLen + 4
	OffsetLeafHash      = OffsetLeafKeyOffset + 8
	SizeLeafWithoutHash = OffsetLeafHash
	SizeLeaf            = SizeLeafWithoutHash + SizeHash
)

// PersistedNode is backed by serialized byte array, usually mmap-ed from disk file.
// Encoding format (all integers are encoded in little endian):
//
// Branch node:
// - height    : 1
// - preTrees  : 1
// - _padding  : 2
// - version   : 4
// - size      : 4
// - key node  : 4  // node index of the smallest leaf in right branch
// - hash      : 32
// Leaf node:
// - version    : 4
// - key len    : 4
// - key offset : 8
// - hash       : 32
type PersistedNode struct {
	snapshot *Snapshot
	isLeaf   bool
	index    uint32
}

var _ Node = PersistedNode{}

func (node PersistedNode) branchNode() NodeLayout {
	_ = "STUB: not implemented"
	return *new(NodeLayout)
}

func (node PersistedNode) leafNode() LeafLayout { _ = "STUB: not implemented"; return *new(LeafLayout) }

func (node PersistedNode) Height() uint8 { _ = "STUB: not implemented"; return 0 }

func (node PersistedNode) IsLeaf() bool { _ = "STUB: not implemented"; return false }

func (node PersistedNode) Version() uint32 { _ = "STUB: not implemented"; return 0 }

func (node PersistedNode) Size() int64 { _ = "STUB: not implemented"; return 0 }

func (node PersistedNode) Key() []byte { _ = "STUB: not implemented"; return nil }

// Value returns nil for non-leaf node.
func (node PersistedNode) Value() []byte { _ = "STUB: not implemented"; return nil }

// Left result is not defined for leaf nodes.
func (node PersistedNode) Left() Node { _ = "STUB: not implemented"; return *new(Node) }

// Right result is not defined for leaf nodes.
func (node PersistedNode) Right() Node { _ = "STUB: not implemented"; return *new(Node) }

func (node PersistedNode) SafeHash() []byte { _ = "STUB: not implemented"; return nil }

func (node PersistedNode) Hash() []byte { _ = "STUB: not implemented"; return nil }

func (node PersistedNode) Mutate(version, _ uint32) *MemNode { _ = "STUB: not implemented"; return nil }

func (node PersistedNode) Get(key []byte) ([]byte, uint32) {
	_ = "STUB: not implemented"
	return nil, 0
}

// binary search in the leaf node array

//nolint:gosec

//nolint:gosec

// return the next index if the key is greater than all keys in the node

func (node PersistedNode) GetByIndex(leafIndex uint32) ([]byte, []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getStartLeaf returns the index of the first leaf in the node.
//
// > start leaf = pre leaves
// >            = pre branches + pre trees
// >            = total branches - sub branches + pre trees
// >            = (index + 1) - (size - 1) + preTrees
// >            = index + 2 - size + preTrees
func getStartLeaf(index, size, preTrees uint32) uint32 { _ = "STUB: not implemented"; return 0 }

// getEndLeaf returns the index of the last leaf in the node.
//
// > end leaf = start leaf + size - 1
// >          = (index + 2 - size + preTrees) + size - 1
// >          = index + 1 + preTrees
func getEndLeaf(index, preTrees uint32) uint32 { _ = "STUB: not implemented"; return 0 }

// getLeftBranch returns the index of the left branch of the node.
//
// > left branch = pre branches + left branches - 1
// >             = (total branches - sub branches) + (left leaves - 1) - 1
// >             = (total branches - sub branches) + (key leaf - start leaf - 1) - 1
// >             = (index+1 - (size-1)) + (key leaf - (index + 2 - size + preTrees) - 1) - 1
// >             = (index - size + 2) + key leaf - index - 2 + size - preTrees - 2
// >             = key leaf - preTrees - 2
func getLeftBranch(keyLeaf, preTrees uint32) uint32 { _ = "STUB: not implemented"; return 0 }

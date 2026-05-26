//go:build nativebyteorder

package memiavl

import (
	"unsafe"
)

func init() {
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)

	if buf != [2]byte{0xCD, 0xAB} {
		panic("native byte order is not little endian, please build without nativebyteorder")
	}
}

type NodeLayout = *nodeLayout

// Nodes is a continuously stored IAVL nodes
type Nodes struct {
	nodes []nodeLayout
}

func NewNodes(buf []byte) (Nodes, error) {
	_ = "STUB: not implemented"
	// check alignment and size of the buffer
	return *new(Nodes), nil
}

func (nodes Nodes) Node(i uint32) NodeLayout {
	_ = "STUB: not implemented"
	return *

	// see comment of `PersistedNode`
	new(NodeLayout)
}

type nodeLayout struct {
	data [4]uint32
	hash [32]byte
}

func (node *nodeLayout) Height() uint8 { _ = "STUB: not implemented"; return 0 }

func (node NodeLayout) PreTrees() uint8 { _ = "STUB: not implemented"; return 0 }

func (node *nodeLayout) Version() uint32 { _ = "STUB: not implemented"; return 0 }

func (node *nodeLayout) Size() uint32 { _ = "STUB: not implemented"; return 0 }

func (node *nodeLayout) KeyLeaf() uint32 { _ = "STUB: not implemented"; return 0 }

func (node *nodeLayout) Hash() []byte { _ = "STUB: not implemented"; return nil }

type LeafLayout = *leafLayout

// Nodes is a continuously stored IAVL nodes
type Leaves struct {
	leaves []leafLayout
}

func NewLeaves(buf []byte) (Leaves, error) {
	_ = "STUB: not implemented"
	// check alignment and size of the buffer
	return *new(Leaves), nil
}

func (leaves Leaves) Leaf(i uint32) LeafLayout { _ = "STUB: not implemented"; return *new(LeafLayout) }

type leafLayout struct {
	version   uint32
	keyLen    uint32
	keyOffset uint64
	hash      [32]byte
}

func (leaf *leafLayout) Version() uint32 { _ = "STUB: not implemented"; return 0 }

func (leaf *leafLayout) KeyLength() uint32 { _ = "STUB: not implemented"; return 0 }

func (leaf *leafLayout) KeyOffset() uint64 { _ = "STUB: not implemented"; return 0 }

func (leaf *leafLayout) Hash() []byte { _ = "STUB: not implemented"; return nil }

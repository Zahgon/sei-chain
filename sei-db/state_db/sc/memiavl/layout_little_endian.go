//go:build !nativebyteorder

package memiavl

// Nodes is a continuously stored IAVL nodes
type Nodes struct {
	data []byte
}

func NewNodes(data []byte) (Nodes, error) { _ = "STUB: not implemented"; return *new(Nodes), nil }

func (nodes Nodes) Node(i uint32) NodeLayout { _ = "STUB: not implemented"; return *new(NodeLayout) }

// see comment of `PersistedNode`
type NodeLayout struct {
	data *[SizeNode]byte
}

func (node NodeLayout) Height() uint8 { _ = "STUB: not implemented"; return 0 }

func (node NodeLayout) PreTrees() uint8 { _ = "STUB: not implemented"; return 0 }

func (node NodeLayout) Version() uint32 { _ = "STUB: not implemented"; return 0 }

func (node NodeLayout) Size() uint32 { _ = "STUB: not implemented"; return 0 }

func (node NodeLayout) KeyLeaf() uint32 { _ = "STUB: not implemented"; return 0 }

func (node NodeLayout) Hash() []byte { _ = "STUB: not implemented"; return nil }

// Leaves is a continuously stored IAVL nodes
type Leaves struct {
	data []byte
}

func NewLeaves(data []byte) (Leaves, error) { _ = "STUB: not implemented"; return *new(Leaves), nil }

func (leaves Leaves) Leaf(i uint32) LeafLayout { _ = "STUB: not implemented"; return *new(LeafLayout) }

type LeafLayout struct {
	data *[SizeLeaf]byte
}

func (leaf LeafLayout) Version() uint32 { _ = "STUB: not implemented"; return 0 }

func (leaf LeafLayout) KeyLength() uint32 { _ = "STUB: not implemented"; return 0 }

func (leaf LeafLayout) KeyOffset() uint64 { _ = "STUB: not implemented"; return 0 }

func (leaf LeafLayout) Hash() []byte { _ = "STUB: not implemented"; return nil }

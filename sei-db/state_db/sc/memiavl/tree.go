package memiavl

import (
	"crypto/sha256"
	"sync"

	ics23 "github.com/confio/ics23/go"

	"github.com/sei-protocol/sei-chain/sei-db/proto"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/types"
	dbm "github.com/tendermint/tm-db"
)

var _ types.CommitKVStore = (*Tree)(nil)
var emptyHash = sha256.New().Sum(nil)

// Tree verify change sets by replay them to rebuild iavl tree and verify the root hashes
type Tree struct {
	version uint32
	// root node of empty tree is represented as `nil`
	root     Node
	snapshot *Snapshot

	initialVersion, cowVersion uint32

	// when true, the get and iterator methods could return a slice pointing to mmaped blob files.
	zeroCopy bool

	// sync.RWMutex is used to protect the tree for thread safety during snapshot reload
	mtx *sync.RWMutex

	pendingChanges chan proto.ChangeSet
	pendingWg      *sync.WaitGroup
}

// NewEmptyTree creates an empty tree at an arbitrary version.
func NewEmptyTree(version uint64, initialVersion uint32) *Tree {
	_ = "STUB: not implemented"
	return nil
}

// no need to copy if the tree is not backed by snapshot

// New creates an empty tree at genesis version
func New(_ int) *Tree { _ = "STUB: not implemented"; return nil }

// NewWithInitialVersion creates an empty tree with initial-version,
// it happens when a new store created at the middle of the chain.
func NewWithInitialVersion(initialVersion uint32) *Tree { _ = "STUB: not implemented"; return nil }

// NewFromSnapshot mmap the blob files and create the root node.
func NewFromSnapshot(snapshot *Snapshot, opts Options) *Tree { _ = "STUB: not implemented"; return nil }

func (t *Tree) SetZeroCopy(zeroCopy bool) { _ = "STUB: not implemented"; return }

func (t *Tree) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (t *Tree) SetInitialVersion(initialVersion int64) error { _ = "STUB: not implemented"; return nil }

// Copy returns a concurrent-safe snapshot. Acquires the underlying *Snapshot
// so background rewrites can't unmap it while the copy is live; callers must
// call Close on the returned tree to release the ref.
func (t *Tree) Copy() *Tree { _ = "STUB: not implemented"; return nil }

// protect the existing `MemNode`s from get modified in-place

// ApplyChangeSet apply the change set of a whole version, and update hashes.
func (t *Tree) ApplyChangeSet(changeSet proto.ChangeSet) { _ = "STUB: not implemented"; return }

func (t *Tree) ApplyChangeSetAsync(changeSet proto.ChangeSet) { _ = "STUB: not implemented"; return }

func (t *Tree) StartBackgroundWrite() { _ = "STUB: not implemented"; return }

func (t *Tree) WaitToCompleteAsyncWrite() { _ = "STUB: not implemented"; return }

func (t *Tree) Set(key, value []byte) { _ = "STUB: not implemented"; return }

// the value could be nil when replaying changes from write-ahead-log because of protobuf decoding

func (t *Tree) Remove(key []byte) { _ = "STUB: not implemented"; return }

// SaveVersion increases the version number and optionally updates the hashes
func (t *Tree) SaveVersion(updateHash bool) ([]byte, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Version returns the current tree version
func (t *Tree) Version() int64 { _ = "STUB: not implemented"; return 0 }

// RootHash updates the hashes and return the current root hash,
// it clones the persisted node's bytes, so the returned bytes is safe to retain.
func (t *Tree) RootHash() []byte { _ = "STUB: not implemented"; return nil }

func (t *Tree) GetWithIndex(key []byte) (int64, []byte) { _ = "STUB: not implemented"; return 0, nil }

func (t *Tree) GetByIndex(index int64) ([]byte, []byte) { _ = "STUB: not implemented"; return nil, nil }

func (t *Tree) Get(key []byte) []byte { _ = "STUB: not implemented"; return nil }

func (t *Tree) Has(key []byte) bool { _ = "STUB: not implemented"; return false }

func (t *Tree) Iterator(start, end []byte, ascending bool) dbm.Iterator {
	_ = "STUB: not implemented"
	return *new(dbm.Iterator)
}

// ScanPostOrder scans the tree in post-order, and call the callback function on each node.
// If the callback function returns false, the scan will be stopped.
func (t *Tree) ScanPostOrder(callback func(node Node) bool) { _ = "STUB: not implemented"; return }

type stackEntry struct {
	node     Node
	expanded bool
}

// Export returns a snapshot of the tree which won't be corrupted by further modifications on the main tree.
func (t *Tree) Export() *Exporter { _ = "STUB: not implemented"; return nil }

// snapshot export algorithm is more efficient

// do normal post-order traversal export

func (t *Tree) Close() error { _ = "STUB: not implemented"; return nil }

// ReplaceWith is used during reload to replace the current tree with the newly loaded snapshot
func (t *Tree) ReplaceWith(other *Tree) error { _ = "STUB: not implemented"; return nil }

// nextVersionU32 is compatible with existing golang iavl implementation.
// see: https://github.com/cosmos/iavl/pull/660
func nextVersionU32(v uint32, initialVersion uint32) uint32 { _ = "STUB: not implemented"; return 0 }

// GetProof takes a key for creating existence or absence proof and returns the
// appropriate merkle.Proof. Since this must be called after querying for the value, this function should never error
// Thus, it will panic on error rather than returning it
func (t *Tree) GetProof(key []byte) *ics23.CommitmentProof { _ = "STUB: not implemented"; return nil }

// value was found

// sanity check: If value was found, membership proof must be creatable

// value wasn't found

// sanity check: If value wasn't found, nonmembership proof must be creatable

package migration

import (
	ics23 "github.com/confio/ics23/go"
	"github.com/sei-protocol/sei-chain/sei-db/proto"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/types"
	db "github.com/tendermint/tm-db"
)

// rootHashSize matches the digest length used by the other CommitKVStore
// implementations in this codebase: memiavl returns sha256 (32 B) and flatkv
// returns Blake3-256 (32 B).
const rootHashSize = 32

var _ types.CommitKVStore = (*RouterCommitKVStore)(nil)

// RouterCommitKVStore adapts a [Router] (which is keyed by store name on every
// call) to the store-name-less [types.CommitKVStore] interface by binding the
// view to a single module store name.
//
// The CommitKVStore interface does not return errors. Any error returned by the
// underlying router is therefore surfaced as a panic. This is a short-term
// limitation; the long-term plan is to plumb errors through the interface.
type RouterCommitKVStore struct {
	router          Router
	storeName       string
	versionProvider func() int64
}

func NewRouterCommitKVStore(
	router Router,
	storeName string,
	versionProvider func() int64,
) *RouterCommitKVStore {
	_ = "STUB: not implemented"
	return nil
}

// Close is illegal during the standard CommitKVStore lifecycle for this type:
// the wrapped Router is owned by the caller and must outlive this view.
func (r *RouterCommitKVStore) Close() error { _ = "STUB: not implemented"; return nil }

func (r *RouterCommitKVStore) Get(key []byte) []byte { _ = "STUB: not implemented"; return nil }

func (r *RouterCommitKVStore) Has(key []byte) bool { _ = "STUB: not implemented"; return false }

func (r *RouterCommitKVStore) Set(key []byte, value []byte) { _ = "STUB: not implemented"; return }

func (r *RouterCommitKVStore) Remove(key []byte) { _ = "STUB: not implemented"; return }

// applyOne dispatches a single KV change as a one-pair NamedChangeSet through
// the router, panicking on any router error.
func (r *RouterCommitKVStore) applyOne(pair *proto.KVPair) { _ = "STUB: not implemented"; return }

func (r *RouterCommitKVStore) Iterator(start []byte, end []byte, ascending bool) db.Iterator {
	_ = "STUB: not implemented"
	return *new(db.Iterator)
}

func (r *RouterCommitKVStore) GetProof(key []byte) *ics23.CommitmentProof {
	_ = "STUB: not implemented"
	return nil
}

// RootHash is a placeholder that returns a fresh zeroed 32-byte slice on every
// call. The CommitKVStore contract permits callers to mutate the returned
// slice, so a fresh allocation is required to keep the placeholder safe.
//
// TODO: revisit before shipping to production once the production usage of
// RootHash() across this code path is understood.
func (r *RouterCommitKVStore) RootHash() []byte { _ = "STUB: not implemented"; return nil }

func (r *RouterCommitKVStore) Version() int64 { _ = "STUB: not implemented"; return 0 }

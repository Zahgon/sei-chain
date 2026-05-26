package memiavl

import (
	ics23 "github.com/confio/ics23/go"
	dbm "github.com/tendermint/tm-db"

	"github.com/sei-protocol/sei-chain/sei-db/proto"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/types"
)

var _ types.Committer = (*CommitStore)(nil)

type CommitStore struct {
	db      *DB
	opts    Options
	homeDir string
}

func NewCommitStore(homeDir string, config Config) *CommitStore {
	_ = "STUB: not implemented"
	return nil
}

// Embed the config directly

// Disable zero copy to avoid segfault during historical read

func (cs *CommitStore) Initialize(initialStores []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs *CommitStore) SetInitialVersion(initialVersion int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs *CommitStore) Rollback(targetVersion int64) error {
	_ = "STUB: not implemented"
	// Close existing resources
	return nil
}

// LoadVersion loads the specified version of the database.
// If copyExisting is true, creates a read-only copy for querying.
func (cs *CommitStore) LoadVersion(targetVersion int64, readOnly bool) (types.Committer, error) {
	_ = "STUB: not implemented"
	return *new(types.Committer), nil
}

// Create a read-only copy via NewCommitStore.

// Close existing resources

// Copy returns an O(1) memiavl snapshot; COW nodes are shared with the live store.
func (cs *CommitStore) Copy() types.Committer {
	_ = "STUB: not implemented"
	return *new(types.Committer)
}

// ReleaseSnapshotRefs releases refs held by a copied in-memory snapshot without
// closing DB-level resources shared with the live store.
func (cs *CommitStore) ReleaseSnapshotRefs() error { _ = "STUB: not implemented"; return nil }

func (cs *CommitStore) Commit() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (cs *CommitStore) Version() int64 { _ = "STUB: not implemented"; return 0 }

// GetLatestVersion returns the highest version durably written to the
// changelog WAL on disk. Note that with AsyncCommitBuffer > 0,
// wal.Write returns before the entry is durable (see sei-db/wal/wal.go,
// "Do not wait for the write to be durable"), so this value can lag the
// in-memory MultiTree by one or more commits while async writes drain.
// Callers that need the just-committed version should use cs.Version()
// or cs.LastCommitInfo().Version, both of which read the in-memory
// tree. The lag is harmless for current production callers
// (rootmulti.NewStore and rootmulti.LastCommitID), which only invoke
// GetLatestVersion before LoadVersion has opened the DB; in that
// pre-load state nothing is in memory anyway.
func (cs *CommitStore) GetLatestVersion() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (cs *CommitStore) ApplyChangeSets(changesets []*proto.NamedChangeSet) error {
	_ = "STUB: not implemented"
	return nil
}

// Apply to tree

func (cs *CommitStore) ApplyUpgrades(upgrades []*proto.TreeNameUpgrade) error {
	_ = "STUB: not implemented"
	return nil
}

// Apply to tree

func (cs *CommitStore) WorkingCommitInfo() *proto.CommitInfo { _ = "STUB: not implemented"; return nil }

func (cs *CommitStore) LastCommitInfo() *proto.CommitInfo { _ = "STUB: not implemented"; return nil }

func (cs *CommitStore) GetChildStoreByName(name string) types.CommitKVStore {
	_ = "STUB: not implemented"
	return *new(types.CommitKVStore)
}

// Return an explicitly nil interface (not a typed-nil *Tree wrapped in an
// interface), so callers can compare the result against nil.

func (cs *CommitStore) Exporter(version int64) (types.Exporter, error) {
	_ = "STUB: not implemented"
	return *new(types.Exporter), nil
}

func (cs *CommitStore) Importer(version int64) (types.Importer, error) {
	_ = "STUB: not implemented"
	return *new(types.Importer), nil
}

func (cs *CommitStore) Close() error { _ = "STUB: not implemented"; return nil }

func (cs *CommitStore) Get(store string, key []byte) (value []byte, ok bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (cs *CommitStore) GetProof(store string, key []byte) (*ics23.CommitmentProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs *CommitStore) Has(store string, key []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cs *CommitStore) Iterator(store string, start []byte, end []byte, ascending bool) (dbm.Iterator, error) {
	_ = "STUB: not implemented"
	return *new(dbm.Iterator), nil
}

// Get the underlying memiavl DB.
func (cs *CommitStore) GetDB() *DB { _ = "STUB: not implemented"; return nil }

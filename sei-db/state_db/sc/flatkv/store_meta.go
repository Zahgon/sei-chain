package flatkv

import (
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/flatkv/ktype"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/flatkv/lthash"
)

// versionToBytes encodes a non-negative version as 8-byte big-endian.
// Panics on negative input to catch programming errors early.
// Only called from internal commit/test paths — never with untrusted input.
func versionToBytes(v int64) []byte { _ = "STUB: not implemented"; return nil }

//nolint:gosec // guarded above

// loadLocalMeta loads per-DB metadata by reading separate keys.
func loadLocalMeta(db types.KeyValueDB) (*ktype.LocalMeta, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // version won't exceed int64 max

// writeLocalMetaToBatch writes per-DB metadata (version + LtHash) as separate keys.
func writeLocalMetaToBatch(batch types.Batch, version int64, ltHash *lthash.LtHash) error {
	_ = "STUB: not implemented"
	return nil
}

// loadGlobalVersion reads the global committed version from metadata DB.
// Returns 0 if not found (fresh start).
func (s *CommitStore) loadGlobalVersion() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:gosec // overflow checked above

// loadGlobalLtHash reads the global committed LtHash from metadata DB.
// Returns nil if not found (fresh start).
func (s *CommitStore) loadGlobalLtHash() (*lthash.LtHash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// commitGlobalMetadata atomically commits global version and global LtHash
// to metadata DB. Per-DB LtHashes are stored in each DB's LocalMeta
// (committed atomically with data in commitBatches).
func (s *CommitStore) commitGlobalMetadata(version int64, hash *lthash.LtHash) error {
	_ = "STUB: not implemented"
	return nil
}

// newPerDBLtHashMap returns a map with a fresh zero LtHash for each data DB.
func newPerDBLtHashMap() map[string]*lthash.LtHash { _ = "STUB: not implemented"; return nil }

// SetInitialVersion seeds the store so that the next Commit produces
// initialVersion. Mirrors memiavl.DB.SetInitialVersion: only valid on a
// truly fresh store (committedVersion == 0 and no prior commits), rejected
// on read-only stores, and persists durably across restart.
//
// Implementation notes:
//   - We persist version = initialVersion - 1 to both the global metadata DB
//     and every per-DB LocalMeta. Commit() does `version := committedVersion + 1`,
//     so the next commit will return initialVersion.
//   - Write order is "global first, per-DB second" so that any partial-write
//     crash recovers as "fresh store" (loadGlobalMetadata lowers the global
//     watermark to the minimum per-DB watermark; per-DB at 0 forces global
//     back to 0). A retry with the same initialVersion is idempotent.
//   - LtHashes stay at their zero values (lthash.New()) — a freshly seeded
//     store has no data, so committed/working LtHashes remain the identity.
func (s *CommitStore) SetInitialVersion(initialVersion int64) error {
	_ = "STUB: not implemented"
	return nil
}

// GetLatestVersion returns the latest committed version persisted under
// dir without holding an open *CommitStore. Mirrors memiavl.GetLatestVersion
// in role: a side-channel for callers that need the on-disk watermark
// before LoadVersion has run (e.g. the rootmulti sanity check at
// process startup). Returns 0 when the store has never been opened or
// has no commits yet.
//
// The truth source is MetaVersionKey in working/metadata. The working
// dir survives across restarts and is updated on every Commit, so this
// matches the precision of memiavl.GetLatestVersion (which reads the
// WAL tail). It must not be called concurrently with a running
// CommitStore on dir, because the underlying PebbleDB takes an
// exclusive file lock.
func GetLatestVersion(dir string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:gosec // overflow checked above

// GetLatestVersion returns the latest committed version. When the store
// is open, the in-memory committed watermark is authoritative; before
// LoadVersion has run, it falls back to the free-standing on-disk
// helper. Either path returns 0 on a fresh store.
func (s *CommitStore) GetLatestVersion() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

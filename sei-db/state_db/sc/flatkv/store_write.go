package flatkv

import (
	"github.com/sei-protocol/sei-chain/sei-db/common/keys"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/flatkv/ktype"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/flatkv/lthash"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/flatkv/vtype"
)

// Commit persists buffered writes and advances the version.
// Protocol: WAL → per-DB batch (with LocalMeta) → flush → update metaDB.
// On crash, catchup replays WAL to recover incomplete commits.
func (s *CommitStore) Commit() (version int64, err error) { _ = "STUB: not implemented"; return 0, nil }

// Auto-increment version

// Step 1: Write Changelog (WAL) - source of truth (always sync)

// Step 2: Commit to each DB (data + LocalMeta.CommittedVersion atomically)

// Step 3: Persist global metadata to metadata DB.
// This must succeed before we update in-memory state; otherwise a
// metadataDB write failure would leave committedVersion advanced while
// the caller sees an error, making the store's internal state
// inconsistent. Per-DB data is already committed (Step 2) and the WAL
// (Step 1) is the source of truth, so a restart will self-heal via
// catchup even if we fail here.

// Step 4: Update in-memory committed state (only after metadata persisted)

// Step 5: Clear pending buffers

// Periodic snapshot so WAL stays bounded and restarts are fast.

// Best-effort WAL truncation, throttled to amortize ReadDir cost.

// flushAllDBs flushes all DBs in parallel.
func (s *CommitStore) flushAllDBs() error { _ = "STUB: not implemented"; return nil }

func (s *CommitStore) clearPendingWrites() { _ = "STUB: not implemented"; return }

// commitBatches commits pending writes to their respective DBs atomically.
// Each DB batch includes LocalMeta update for crash recovery.
// Batches are built serially, then committed in parallel.
// Also called by catchup to replay WAL without re-writing changelog.
func (s *CommitStore) commitBatches(version int64) error { _ = "STUB: not implemented"; return nil }

// Commit all batches in parallel.

// Update in-memory local meta after all commits succeed.

func prepareBatch[T vtype.VType](
	db types.KeyValueDB,
	writes map[string]T,
	version int64,
	localMeta *ktype.LocalMeta,
	ltHash *lthash.LtHash,
	dbName string,
) (types.Batch, error) {
	_ = "STUB: not implemented"
	return *new(types.Batch), nil
}

// collectPendingReads partitions keys from changeMaps into those already
// buffered in pendingWrites (copied to old) and those needing a DB read
// (returned as a BatchGetResult map).
func collectPendingReads[T vtype.VType](
	pendingWrites map[string]T,
	old map[string]T,
	changeMaps ...map[string][]byte,
) map[string]types.BatchGetResult {
	_ = "STUB: not implemented"
	return nil
}

// deserializeBatchResults converts raw BatchGetResults into typed values.
func deserializeBatchResults[T vtype.VType](
	batch map[string]types.BatchGetResult,
	old map[string]T,
	deserialize func([]byte) (T, error),
	dbName string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// rawKVPair is a raw physical key/value pair as stored on disk.
type rawKVPair struct {
	Key   []byte
	Value []byte
}

// FinalizeImport persists per-DB metadata (version + LtHash) and global
// metadata after all import data has been written. This must be called
// exactly once at the end of an import to make the data durable across restarts.
func (s *CommitStore) FinalizeImport(version int64) error { _ = "STUB: not implemented"; return nil }

// batchReadOldValues returns the prior value for every key in changesByType.
// Pending writes are resolved from memory; the rest are batch-read from disk
// in parallel.
func (s *CommitStore) batchReadOldValues(changesByType map[keys.EVMKeyKind]map[string][]byte) (
	storageOld map[string]*vtype.StorageData,
	accountOld map[string]*vtype.AccountData,
	codeOld map[string]*vtype.CodeData,
	legacyOld map[string]*vtype.LegacyData,
	err error,
) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil
}

// TODO: add balance changeMap when balance key is supported.

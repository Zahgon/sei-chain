package flatkv

import (
	"context"
	"errors"
	"time"

	"github.com/zbiljic/go-filelock"

	"github.com/sei-protocol/sei-chain/sei-db/common/metrics"
	"github.com/sei-protocol/sei-chain/sei-db/common/threading"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/dbcache"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/pebbledb"
	seidbtypes "github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
	"github.com/sei-protocol/sei-chain/sei-db/proto"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/flatkv/config"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/flatkv/ktype"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/flatkv/lthash"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/flatkv/vtype"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/types"
	"github.com/sei-protocol/sei-chain/sei-db/wal"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("db", "state-db", "sc", "flatkv")

const (
	// Top-level directory names
	flatkvRootDir = "flatkv"
	changelogDir  = "changelog"
	lockFileName  = "LOCK"

	// DB subdirectories (inside each snapshot)
	accountDBDir = "account"
	codeDBDir    = "code"
	storageDBDir = "storage"
	legacyDBDir  = "legacy"
	metadataDir  = "metadata"

	// Suffixes for atomic directory operations
	tmpSuffix      = "-tmp"
	removingSuffix = "-removing"

	readOnlyDirPrefix = "readonly-"

	flatkvMeterName = "seidb_flatkv"
)

// dataDBDirs lists all data DB directory names (used for per-DB LtHash iteration).
var dataDBDirs = []string{accountDBDir, codeDBDir, storageDBDir, legacyDBDir}

// InitializeDataDirectories sets the DataDir for each nested PebbleDB config
// that does not already have one, using DataDir as the base path. The DBs live
// under the working directory: <DataDir>/working/<subdir>.
func InitializeDataDirectories(c *config.Config) { _ = "STUB: not implemented"; return }

func applyPebbleMetricsConfig(c *config.Config) {
	_ = "STUB: not implemented"
	// Keep a single FlatKV-level knob for Pebble internal metrics. Per-DB
	// EnableMetrics values are intentionally overwritten here.
	return
}

// CommitStore implements flatkv.Store for EVM state storage.
// NOT thread-safe; callers must serialize all operations.
type CommitStore struct {
	ctx    context.Context
	cancel context.CancelFunc
	config config.Config
	dbDir  string

	// Five separate PebbleDB instances.
	// Physical key format: "module/" + type_prefix + stripped_key.
	metadataDB seidbtypes.KeyValueDB // Global version + LtHash watermark
	accountDB  seidbtypes.KeyValueDB // "evm/"+0x0a+addr(20) → vtype.AccountData
	codeDB     seidbtypes.KeyValueDB // "evm/"+0x07+addr(20) → vtype.CodeData
	storageDB  seidbtypes.KeyValueDB // "evm/"+0x03+addr(20)||slot(32) → vtype.StorageData
	legacyDB   seidbtypes.KeyValueDB // "module/"+key → vtype.LegacyData

	// Per-DB committed version, keyed by DB dir name (e.g. accountDBDir).
	localMeta map[string]*ktype.LocalMeta

	// LtHash state for integrity checking
	committedVersion int64
	committedLtHash  *lthash.LtHash
	workingLtHash    *lthash.LtHash

	// Per-DB working LTHash tracking. Authoritative copies live in each
	// DB's LocalMeta (atomically committed with data). On startup the
	// working hashes are loaded from LocalMeta.
	perDBWorkingLtHash map[string]*lthash.LtHash

	// Pending writes buffer
	accountWrites map[string]*vtype.AccountData
	codeWrites    map[string]*vtype.CodeData
	storageWrites map[string]*vtype.StorageData
	legacyWrites  map[string]*vtype.LegacyData

	changelog wal.ChangelogWAL

	// Changes to feed into the WAL at the next commit.
	pendingChangeSets []*proto.NamedChangeSet

	lastSnapshotTime time.Time

	// File lock prevents multiple processes from opening the same DB.
	fileLock filelock.TryLockerSafe

	phaseTimer *metrics.PhaseTimer

	// readOnly marks stores opened via LoadVersion(..., true).
	readOnly bool

	readOnlyWorkDir string // Temp working dir for readonly store; removed by Close.

	// A work pool for reading from the DBs.
	//
	// Uses a fixed-size pool.
	readPool threading.Pool

	// A work pool for miscellaneous operations that are neither computationally intensive nor IO bound.
	//
	// Uses an elasticly-sized pool, so it is safe to submit tasks that have dependencies on other tasks in the pool.
	miscPool threading.Pool
}

var _ Store = (*CommitStore)(nil)

// dataDBs returns the four data PebbleDB instances in fixed iteration order:
// accountDB, codeDB, storageDB, legacyDB. metadataDB is excluded.
func (s *CommitStore) dataDBs() []seidbtypes.KeyValueDB { _ = "STUB: not implemented"; return nil }

type namedDB struct {
	dir string
	db  seidbtypes.KeyValueDB
}

// namedDataDBs returns the four data DBs paired with their directory names.
func (s *CommitStore) namedDataDBs() []namedDB { _ = "STUB: not implemented"; return nil }

// routePhysicalKey maps a physical DB key to its target database.
// Non-EVM modules are routed to legacyDB; EVM keys are routed by kind.
func (s *CommitStore) routePhysicalKey(physicalKey []byte) (seidbtypes.KeyValueDB, error) {
	_ = "STUB: not implemented"
	return *new(seidbtypes.KeyValueDB), nil
}

// NewCommitStore creates a new (unopened) FlatKV commit store.
// Call LoadVersion to open and initialize.
func NewCommitStore(
	ctx context.Context,
	cfg *config.Config,
) (*CommitStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// resetPools recreates the context and thread pools after a full Close().
func (s *CommitStore) resetPools() { _ = "STUB: not implemented"; return }

func (s *CommitStore) flatkvDir() string { _ = "STUB: not implemented"; return "" }

var errReadOnly = errors.New("flatkv: store is read-only")

// LoadVersion opens the database at the given version (0 = latest).
// When readOnly is true an isolated, read-only CommitStore is returned;
// the caller must Close it when done.
func (s *CommitStore) LoadVersion(targetVersion int64, readOnly bool) (opened Store, retErr error) {
	_ = "STUB: not implemented"
	return *new(Store), nil
}

// Track whether we acquire the lock in this call so we can release it
// on any error path (open() won't track a pre-held lock).

// Acquire lock before mutating the current symlink to prevent
// a race with another process observing an unintended baseline.

// Force a fresh working dir clone: the working dir may contain data
// beyond targetVersion from a previous open-to-latest.

// loadVersionReadOnly creates an isolated, read-only CommitStore at the
// requested version. If the writer lock has not yet been acquired (e.g. the
// store was freshly constructed), CleanupOrphanedReadOnlyDirs is called
// lazily to acquire it and clean up any leftover directories. When the lock
// is acquired lazily, ownership is transferred to the returned clone so that
// closing the clone releases it; this prevents leaking the lock when the
// caller never explicitly closes the parent store.
func (s *CommitStore) loadVersionReadOnly(targetVersion int64) (_ Store, retErr error) {
	_ = "STUB: not implemented"
	return *new(Store), nil
}

// Transfer the lazily-acquired lock to the clone so that ro.Close()
// releases it, preventing a leak when the parent is never closed.

// openReadOnly opens PebbleDBs in readOnlyWorkDir, replays the WAL to
// targetVersion, then closes the WAL and marks the store as read-only.
// It never modifies the global "current" symlink.
func (s *CommitStore) openReadOnly(targetVersion int64) error {
	_ = "STUB: not implemented"
	return nil
}

// openTo opens all DBs and catches up via WAL to the given version.
//   - 0  -> replay to end of WAL (latest).
//   - >0 -> replay up to (and including) that version.
func (s *CommitStore) openTo(catchupTarget int64) error { _ = "STUB: not implemented"; return nil }

// open opens all database instances.
//
// Layout:
//
//	flatkv/
//	  current -> snapshot-NNNNN
//	  snapshot-NNNNN/{account,code,...}/  (immutable)
//	  working/{account,code,...}/          (mutable clone)
//	  changelog/                           (WAL, shared)
//
// The baseline snapshot is cloned into working/ on every open so that
// PebbleDB writes never mutate snapshot directories. On first run,
// existing flat DB directories are migrated into a snapshot.
func (s *CommitStore) open() (retErr error) { _ = "STUB: not implemented"; return nil }

func (s *CommitStore) acquireFileLock(dir string) error { _ = "STUB: not implemented"; return nil }

// openPebbleDB creates the directory at cfg.DataDir and opens a PebbleDB instance.
func (s *CommitStore) openPebbleDB(cfg *pebbledb.PebbleDBConfig, cacheCfg *dbcache.CacheConfig) (seidbtypes.KeyValueDB, error) {
	_ = "STUB: not implemented"
	return *new(seidbtypes.KeyValueDB), nil
}

// openDBs opens all PebbleDBs from dbDir and optionally the changelog WAL
// from changelogRoot. On failure all already-opened handles are closed.
func (s *CommitStore) openDBs(dbDir, changelogRoot string) (retErr error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *CommitStore) loadGlobalMetadata() error { _ = "STUB: not implemented"; return nil }

// Load per-DB LtHashes from each DB's LocalMeta (already loaded in openDBs).
// If any DB's version is behind the global version (partial commit or
// corruption), lower committedVersion so catchup replays from there.

// clearChangelog closes the WAL, removes its directory, and reopens an empty
// WAL. Used by Rollback when the target version predates all WAL entries and
// the entire log must be discarded to prevent re-application on restart.
func (s *CommitStore) clearChangelog() error { _ = "STUB: not implemented"; return nil }

func (s *CommitStore) Version() int64 { _ = "STUB: not implemented"; return 0 }

// RootHash returns the Blake3-256 digest of the working LtHash.
func (s *CommitStore) RootHash() []byte { _ = "STUB: not implemented"; return nil }

// CommittedRootHash returns the Blake3-256 digest of the last committed LtHash.
func (s *CommitStore) CommittedRootHash() []byte { _ = "STUB: not implemented"; return nil }

func (s *CommitStore) Importer(version int64) (types.Importer, error) {
	_ = "STUB: not implemented"
	return *new(types.Importer), nil
}

// rootmulti.Restore closes the store before creating an importer.
// Close() cancels the context (killing pools), so recreate them
// before reopening the DBs.

// resetForImport purges all existing data so that a subsequent import
// produces a clean store containing only the snapshot being restored.
// Without this, keys that exist locally but were deleted in the remote
// snapshot would survive the import, producing a mixed stale state.
func (s *CommitStore) resetForImport() error { _ = "STUB: not implemented"; return nil }

// rootmulti.Restore calls Close() (which releases the file lock)
// before calling Importer(). Re-acquire the lock before mutating
// the data directory so no other process can interfere.

// Reopen from a pristine empty state. open() will load metadata
// from the empty DB (a no-op), then we reset in-memory state below.

func (s *CommitStore) GetPhaseTimer() *metrics.PhaseTimer { _ = "STUB: not implemented"; return nil }

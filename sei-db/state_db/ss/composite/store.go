package composite

import (
	"sync"

	"github.com/sei-protocol/sei-chain/sei-db/config"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
	"github.com/sei-protocol/sei-chain/sei-db/proto"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/ss/pruning"
	"github.com/sei-protocol/sei-chain/sei-db/wal"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("db", "state-db", "ss", "composite")

// Compile-time check.
var _ types.StateStore = (*CompositeStateStore)(nil)

// CompositeStateStore routes operations between Cosmos_SS and EVM_SS.
// Both are db_engine.StateStore; the composite itself also implements db_engine.StateStore.
type CompositeStateStore struct {
	cosmosStore    types.StateStore // CosmosStateStore wrapping MVCC DB
	evmStore       types.StateStore // EVMStateStore wrapping sub MVCC DBs (nil if disabled)
	pruningManager *pruning.Manager
	config         config.StateStoreConfig
	closeOnce      sync.Once
	closeErr       error
}

// NewCompositeStateStore creates a new composite state store.
// Backend (PebbleDB or RocksDB) is resolved at compile time via build-tag-gated files in db_engine/backend.
func NewCompositeStateStore(
	ssConfig config.StateStoreConfig,
	homeDir string,
) (*CompositeStateStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Runs before the DB is opened so a rejection leaves no empty dir behind.

// Catches a dir-exists-but-DB-empty case the directory check can't see.

// Mismatched earliest versions = DBs from different snapshots; reads would diverge.

// ssHasData: checks both latest and earliest because state-sync restore only sets earliest.
func ssHasData(ss types.StateStore) bool { _ = "STUB: not implemented"; return false }

// validateEVMSSDirectory rejects enabling evm-ss-split on a populated Cosmos SS
// when the EVM SS dir is missing or empty (i.e. flipping the flag without state sync).
func validateEVMSSDirectory(cosmosStore types.StateStore, evmDir string) error {
	_ = "STUB: not implemented"
	return nil
}

// fresh node, nothing to diverge from

// validateEVMSSPreRecovery rejects an opened-but-empty EVM SS against a populated Cosmos SS.
func (s *CompositeStateStore) validateEVMSSPreRecovery() error {
	_ = "STUB: not implemented"
	return nil
}

// validateEVMSSPostRecovery rejects mismatched earliest versions between the two SS DBs.
func (s *CompositeStateStore) validateEVMSSPostRecovery() error {
	_ = "STUB: not implemented"
	return nil
}

func (s *CompositeStateStore) StartPruning() { _ = "STUB: not implemented"; return }

// evmRouted returns true when the key should be served from the EVM backend.
// If evmStore is open at all, EVMSplit was true at startup and the backend is
// the sole home for EVM data — routing to cosmos would return wrong/empty.
func (s *CompositeStateStore) evmRouted(storeKey string) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *CompositeStateStore) Get(storeKey string, version int64, key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *CompositeStateStore) Has(storeKey string, version int64, key []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *CompositeStateStore) Iterator(storeKey string, version int64, start, end []byte) (types.DBIterator, error) {
	_ = "STUB: not implemented"
	return *new(types.DBIterator), nil
}

func (s *CompositeStateStore) ReverseIterator(storeKey string, version int64, start, end []byte) (types.DBIterator, error) {
	_ = "STUB: not implemented"
	return *new(types.DBIterator), nil
}

func (s *CompositeStateStore) RawIterate(storeKey string, fn func([]byte, []byte, int64) bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *CompositeStateStore) GetLatestVersion() int64 { _ = "STUB: not implemented"; return 0 }

func (s *CompositeStateStore) GetEarliestVersion() int64 { _ = "STUB: not implemented"; return 0 }

func (s *CompositeStateStore) Close() error { _ = "STUB: not implemented"; return nil }

// =============================================================================
// Write path
// =============================================================================

func (s *CompositeStateStore) SetLatestVersion(version int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *CompositeStateStore) SetEarliestVersion(version int64, ignoreVersion bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *CompositeStateStore) ApplyChangesetSync(version int64, changesets []*proto.NamedChangeSet) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *CompositeStateStore) ApplyChangesetAsync(version int64, changesets []*proto.NamedChangeSet) error {
	_ = "STUB: not implemented"
	return nil
}

func filterEVMChangesets(changesets []*proto.NamedChangeSet) []*proto.NamedChangeSet {
	_ = "STUB: not implemented"
	return nil
}

func stripEVMFromChangesets(changesets []*proto.NamedChangeSet) []*proto.NamedChangeSet {
	_ = "STUB: not implemented"
	return nil
}

// convertFlatKVNodes transforms a single FlatKV physical-key snapshot node
// into one or more SS nodes by stripping the module prefix from the key,
// deserializing the vtype metadata from the value, and (for merged account
// rows) splitting into separate nonce and codeHash nodes.
//
// For EVM-specific keys (account, storage, code) the output StoreKey is "evm".
// For legacy keys the original module name is preserved so they route back to
// the correct Cosmos SS module.
func convertFlatKVNodes(node types.SnapshotNode) ([]types.SnapshotNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *CompositeStateStore) Import(version int64, ch <-chan types.SnapshotNode) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *CompositeStateStore) Prune(version int64) error { _ = "STUB: not implemented"; return nil }

// =============================================================================
// Recovery
// =============================================================================

func RecoverCompositeStateStore(
	changelogPath string,
	compositeStore *CompositeStateStore,
) error {
	_ = "STUB: not implemented"
	return nil
}

type WALEntryHandler func(entry proto.ChangelogEntry) error

func ReplayWAL(
	changelogPath string,
	fromVersion int64,
	toVersion int64,
	handler WALEntryHandler,
) error {
	_ = "STUB: not implemented"
	return nil
}

func findReplayStartOffset(streamHandler wal.ChangelogWAL, firstOffset, lastOffset uint64, targetVersion int64) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

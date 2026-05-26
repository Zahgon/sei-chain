package migration

import (
	ics23 "github.com/confio/ics23/go"
	"github.com/sei-protocol/sei-chain/sei-db/proto"
	"github.com/sei-protocol/seilog"
	db "github.com/tendermint/tm-db"
)

var logger = seilog.NewLogger("db", "state-db", "sc", "migration")

var _ Router = (*MigrationManager)(nil)

// MigrationManager handles migration from one database to another,
// routing reads and writes during the course of the migration.
//
// MigrationManager is NOT safe for concurrent use; wrap it with
// NewThreadSafeRouter to serialize external callers. BuildRouter wraps
// each Router it returns automatically.
type MigrationManager struct {
	// For reading values out of the old database.
	oldDBReader DBReader

	// For writing values to the old database.
	oldDBWriter DBWriter

	// For reading values out of the new database.
	newDBReader DBReader

	// For writing values to the new database.
	newDBWriter DBWriter

	// For iterating through key-value pairs to migrate in the old
	// database.
	iterator MigrationIterator

	// The boundary of the migration. All keys to the left of (or equal
	// to) the boundary are considered migrated. Reaches
	// MigrationBoundaryComplete on the final block of the migration.
	boundary MigrationBoundary

	// The number of key-value pairs to migrate after each write operation.
	migrationBatchSize int

	// The version we want to migrate to.
	targetVersion uint64

	// Optional metrics sink. May be nil; all calls on this field go
	// through nil-safe methods on *MigrationMetrics.
	metrics *MigrationMetrics
}

// Handles the migration of data from one database to another.
func NewMigrationManager(
	// The number of key-value pairs to migrate after each write operation. Must be > 0.
	migrationBatchSize int,
	// The migration version the stored data is expected to be at on entry. If no prior migration
	// version is stored in the DB, startVersion should be 0.
	startVersion uint64,
	// The migration version after the migration is complete.
	// Must be strictly greater than startVersion.
	targetVersion uint64,
	// For reading values out of the old database.
	oldDBReader DBReader,
	// For writing values to the old database.
	oldDBWriter DBWriter,
	// For reading values out of the new database.
	newDBReader DBReader,
	// For writing values to the new database.
	newDBWriter DBWriter,
	// For iterating through key-value pairs to migrate in the old database.
	iterator MigrationIterator,
	// Optional metrics sink. Pass nil to disable metric emission.
	metrics *MigrationMetrics,
) (*MigrationManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Migration metadata is owned exclusively by the new DB (flatkv).

// The final block of the migration wrote MigrationVersionKey =
// targetVersion and deleted MigrationBoundaryKey atomically, so
// there is no boundary on disk to read. Come up in passthrough:
// every read routes to the new DB via IsMigrated, every write
// takes the post-completion early-return in ApplyChangeSets,
// and the iterator's Complete short-circuit keeps it inert.

// readMigrationBoundary reads the current migration boundary from the new
// database, or returns MigrationBoundaryNotStarted if none is stored yet.
func readMigrationBoundary(newDBReader DBReader) (MigrationBoundary, error) {
	_ = "STUB: not implemented"
	return *new(MigrationBoundary), nil
}

// readVersionFromDB reads MigrationVersionKey from the given DB's
// MigrationStore, returning (version, present, error). An absent key is
// reported as (0, false, nil) so the caller can distinguish "not set"
// from "explicitly zero".
//
// This helper takes a raw DBReader rather than going through
// MigrationManager.Read because the MigrationStore is reserved for
// internal use and MigrationManager.Read rejects reads against it.
func readVersionFromDB(reader DBReader) (uint64, bool, error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

// IsAtVersion reports whether the DB reached by reader is currently at the
// given migration version. An absent MigrationVersionKey is interpreted as
// version 0.
func IsAtVersion(reader DBReader, version uint64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Read a value from the database. If the requested value is migrated,
// read it from the new database. Otherwise, read it from the old
// database. After the boundary has reached MigrationBoundaryComplete on
// the final block of the migration, IsMigrated returns true for every
// key, so all reads route to the new DB.
//
// Reads targeting MigrationStore are rejected with an error: that store
// is reserved for the manager's own bookkeeping.
//
// Not safe for concurrent use; wrap with NewThreadSafeRouter.
func (m *MigrationManager) Read(store string, key []byte) ([]byte, bool, error) {
	_ = "STUB: not implemented"
	return nil,

		// The migration module is reserved for internal use, do not permit outer scope reads from it.
		false, nil
}

// This key has already been migrated, read it from the new DB.

// This key has not been migrated, read it from the old DB.

// ApplyChangeSets applies a batch of change sets to the database.
//
// Not safe for concurrent use; wrap with NewThreadSafeRouter.
func (m *MigrationManager) ApplyChangeSets(changesets []*proto.NamedChangeSet) error {
	_ = "STUB: not implemented"
	return nil
}

// The migration module is reserved for internal use, do not permit outer scope writes to it.

// Migration is complete; forward the caller's writes to the new DB only.

// Get the next batch of keys to migrate.

// Pairs destined for each DB, grouped by store name and keyed by KVPair.Key.

// Create change sets that move the values to migrate from the old DB to the new DB.

// Write the value to the new DB.

// Delete the value from the old DB.

// For each pair in the original change sets, route to the appropriate database.
// These must overwrite migrated values, so it's important to do this after we've collected
// the change set for the migrated values.

// On the final block of the migration, update the migration version and delete the boundary.

// Mirror the on-disk version bump in the in-memory metric so the
// version gauge and the boundary-snapshot loop see the
// completion at the same moment the DB does.

// On every other block of the migration, update the boundary.

// putPair inserts pair into dest under (store, pair.Key), creating the inner
// map on demand. Later writes to the same (store, key) overwrite earlier ones.
func putPair(dest map[string]map[string]*proto.KVPair, store string, pair *proto.KVPair) {
	_ = "STUB: not implemented"
	return
}

// flattenPairsByStore collapses a store-keyed map of (key -> KVPair) into one
// NamedChangeSet per store, with stores and pairs emitted in sorted order for
// deterministic downstream writes.
func flattenPairsByStore(pairsByStore map[string]map[string]*proto.KVPair) []*proto.NamedChangeSet {
	_ = "STUB: not implemented"
	return nil
}

// GetProof implements [Router].
func (m *MigrationManager) GetProof(store string, key []byte) (*ics23.CommitmentProof, error) {
	_ = "STUB: not implemented"
	// We won't be able to serve state proofs for flatKV until we implement BUD proofs.
	return nil, nil
}

// Iterator implements [Router].
func (m *MigrationManager) Iterator(store string, start []byte, end []byte, ascending bool) (db.Iterator, error) {
	_ = "STUB: not implemented"
	// Eventually we will implement iteration for some modules within FlatKV, but never for the evm/ module.
	// Since we're migrating the evm/ module first, implementing iteration for FlatKV is not a blocker.
	return *new(db.Iterator), nil
}

// BuildRoute returns a Route that dispatches the given module names to
// this MigrationManager. Reads, writes, iteration and proof requests
// for those modules will all flow through this migration manager.
//
// Module names must be unique; NewRoute's validation rules apply. The
// returned Route may be passed to NewModuleRouter alongside other
// Routes to compose multi-database setups.
func (m *MigrationManager) BuildRoute(moduleNames ...string) (*Route, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

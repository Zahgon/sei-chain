package sstest

import (
	"github.com/stretchr/testify/suite"

	"github.com/sei-protocol/sei-chain/sei-db/config"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
)

const (
	storeKey1 = "store1"
	storeKey2 = "store2"
	pebbledb  = "pebbledb"
)

// BaseStorageTestSuite defines a reusable test suite for all storage backends.
// It contains tests that work with any comparer (both MVCCComparer and DefaultComparer).
// For tests that require MVCCComparer's iterator functionality, use StorageTestSuite instead.
type BaseStorageTestSuite struct {
	suite.Suite

	NewDB          func(dir string, config config.StateStoreConfig) (types.StateStore, error)
	EmptyBatchSize int
	Config         config.StateStoreConfig
}

// StorageTestSuite extends BaseStorageTestSuite with tests that require MVCCComparer.
// This includes all iterator-related tests and prune tests that use iteration internally.
type StorageTestSuite struct {
	BaseStorageTestSuite
}

func (s *BaseStorageTestSuite) TestDatabaseClose_IsIdempotent() { _ = "STUB: not implemented"; return }

// The close operation on an already closed db should be a noop, because:
//  1. it avoids panic in cases where the db closure may be called multiple times in the call stack.
//  2. it is harmless to make it idempotent and adheres with Go SDK's IO package Close behavior.

func (s *BaseStorageTestSuite) TestDatabaseLatestVersion() { _ = "STUB: not implemented"; return }

// Test even after closing and reopening, the latest version is maintained

func (s *BaseStorageTestSuite) TestDatabaseVersionedKeys() { _ = "STUB: not implemented"; return }

func (s *BaseStorageTestSuite) TestDatabaseGetVersionedKey() { _ = "STUB: not implemented"; return }

// store a key at version 1

// assume chain progresses to version 10 w/o any changes to key

// chain progresses to version 11 with an update to key

// chain progresses to version 15 with a delete to key

// all queries up to version 14 should return the latest value

// all queries after version 15 should return nil

func (s *BaseStorageTestSuite) TestDatabaseVersionZero() {
	_ = "STUB: not implemented"
	// Db should write all keys at version 0 at version 1
	return
}

// Get at version 0 should return error

// Retrieve each key at version 1

func (s *BaseStorageTestSuite) TestDatabaseApplyChangeset() { _ = "STUB: not implemented"; return }

// Deletes

// Ensure ApplyChangesetSync(version, []*NamedChangeSet{moduleA, moduleB, ...})
// only bumps latest version after all module writes are persisted.
func (s *BaseStorageTestSuite) TestApplyChangesetSyncAtomicAcrossModules() {
	_ = "STUB: not implemented"
	return
}

// Prepare two modules' changesets at the same version

// Latest version should be 0 before apply

// Apply both modules in one sync call

// After ApplyChangesetSync returns, latest version must reflect completion

// And both modules' data must be readable at version 1

func (s *StorageTestSuite) TestDatabaseIteratorEmptyDomain() { _ = "STUB: not implemented"; return }

func (s *StorageTestSuite) TestDatabaseIteratorClose() { _ = "STUB: not implemented"; return }

// Close is idempotent

func (s *StorageTestSuite) TestDatabaseIteratorDomain() { _ = "STUB: not implemented"; return }

func (s *StorageTestSuite) TestDatabaseIterator() { _ = "STUB: not implemented"; return }

// iterator without an end key over multiple versions

// seek past domain, which should make the iterator invalid and produce an error

// iterator with a start and end domain over multiple versions

// seek past domain, which should make the iterator invalid and produce an error

// start must be <= end

func (s *StorageTestSuite) TestDatabaseIteratorRangedDeletes() { _ = "STUB: not implemented"; return }

// there should only be one valid key in the iterator -- key001

func (s *StorageTestSuite) TestDatabaseIteratorDeletes() { _ = "STUB: not implemented"; return }

// there should be only one valid key in the iterator

// there should be two valid keys in the iterator

func (s *StorageTestSuite) TestDatabaseIteratorMultiVersion() { _ = "STUB: not implemented"; return }

// for versions 50-100, only update even keys

// All keys should be present; All odd keys should have a value that reflects
// version 49, and all even keys should have a value that reflects the desired
// version, 69.

// Tests bug where iterator loops continuously
func (s *StorageTestSuite) TestDatabaseBugInitialReverseIteration() {
	_ = "STUB: not implemented"
	return
}

// Forward Iteration
// Less than iterator version

func (s *StorageTestSuite) TestDatabaseBugInitialForwardIteration() {
	_ = "STUB: not implemented"
	return
}

// Forward Iteration
// Less than iterator version

func (s *StorageTestSuite) TestDatabaseBugInitialForwardIterationHigher() {
	_ = "STUB: not implemented"
	return
}

// Less than iterator version

func (s *StorageTestSuite) TestDatabaseBugInitialReverseIterationHigher() {
	_ = "STUB: not implemented"
	return
}

// Reverse Iteration
// Less than iterator version

func (s *StorageTestSuite) TestDatabaseIteratorNoDomain() { _ = "STUB: not implemented"; return }

// create an iterator over the entire domain

func (s *StorageTestSuite) TestDatabasePrune() { _ = "STUB: not implemented"; return }

// Verify earliest version is 0

// Verify latest version is accessible

// prune the first 25 versions

// Verify earliest version is 26 (first 25 pruned)

// Verify metadata keys are still accessible after pruning
// (this verifies metadata keys weren't corrupted during pruning)

// Ensure all keys are no longer present up to and including version 25 and
// all keys are present after version 25.

// prune the latest version which should prune the entire dataset

// Verify earliest version is 51 (first 50 pruned)

// Verify metadata is still intact after full pruning

func (s *BaseStorageTestSuite) TestDatabasePruneAndTombstone() { _ = "STUB: not implemented"; return }

// write a key at three different versions 1, 100 and 200

// prune version 150

func (s *StorageTestSuite) TestDatabasePruneKeepRecent() { _ = "STUB: not implemented"; return }

// write a key at three different versions 1, 100 and 200

// prune version 50

// ensure queries for versions 50 and older return nil

// ensure the value previously at version 1 is still there for queries greater than 50

// ensure the correct value at a greater height

// prune latest height and ensure we have the previous version when querying above it

func (s *BaseStorageTestSuite) TestDatabasePruneKeepLastVersion() {
	_ = "STUB: not implemented"
	// Only test KeepLastVersion = false for pebbledb backend
	// NOTE: KeepLastVersion is always true and will be removed in future
	return
}

// Update config to set KeepLastVersion to false

// prune version 150

// Verify that all keys before prune height are deleted

// Verify keys after prune height can be retrieved

// Now reset KeepLastVersion to true and verify latest version of key exists

// prune version 150

// Can still retrieve those keys because KeepLastVersion is true

func (s *StorageTestSuite) TestDatabaseReverseIterator() { _ = "STUB: not implemented"; return }

// reverse iterator without an end key

// seek past domain, which should make the iterator invalid and produce an error

// reverse iterator with with a start and end domain

// seek past domain, which should make the iterator invalid and produce an error

// start must be <= end

func (s *BaseStorageTestSuite) TestParallelWrites() { _ = "STUB: not implemented"; return }

// start 10 goroutines that write to the database

// start the goroutines

// check that all the data is there

func (s *StorageTestSuite) TestParallelWriteAndPruning() { _ = "STUB: not implemented"; return }

// start a goroutine that write to the database

// start a goroutine that prunes the database

// wait for the goroutines

// check if the data is pruned

func (s *StorageTestSuite) TestDatabaseParallelDeleteIteration() { _ = "STUB: not implemented"; return }

// start a goroutine that deletes from the database at latest Version

// start a goroutine that iterates over the database

// iterator without an end key over multiple versions

// seek past domain, which should make the iterator invalid and produce an error

// wait for the goroutines

// Verify deletes

func (s *BaseStorageTestSuite) TestDatabaseParallelWriteDelete() { _ = "STUB: not implemented"; return }

// start a goroutine that writes to the database at latestVersion

// Apply changeset for each key separately

// start a goroutine that deletes from the database

// Apply changeset for each key separately

// wait for the goroutines

// Verify writes and deletes on latest version

func (s *StorageTestSuite) TestParallelIterationAndPruning() { _ = "STUB: not implemented"; return }

// start a goroutine that prunes the database

// start a goroutine that iterates over the database

// iterator without an end key over multiple versions

// seek past domain, which should make the iterator invalid and produce an error

// wait for the goroutines

// Ensure all keys are no longer present up to latestVersion - 20 and
// all keys are present after

func (s *StorageTestSuite) TestDatabaseParallelIterationVersions() {
	_ = "STUB: not implemented"
	return
}

// start multiple goroutines that iterate over different version of the database

// seek past domain, which should make the iterator invalid and produce an error

// wait for the goroutines

func (s *BaseStorageTestSuite) TestDatabaseImport() { _ = "STUB: not implemented"; return }

// TestDatabaseReverseIteratorPrefixIsolation Verifies that ReverseIterator(nil, nil) is clamped to the caller's prefix
// via prefixEnd()/UpperBound and does **not** spill into the next module.
func (s *StorageTestSuite) TestDatabaseReverseIteratorPrefixIsolation() {
	_ = "STUB: not implemented"
	return
}

// store1 : key000-key009
// store2 : key000-key009   (different prefix, same suffixes)

// ---------- nil / nil reverse scan on store1 ----------

// We should see exactly the 10 keys from store1, in reverse order,
// and we should *never* see a key that belongs to store2.

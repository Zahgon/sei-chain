// Package composite provides a unified commit store that coordinates
// between Cosmos (memiavl) and EVM (flatkv) committers.
package composite

import (
	"context"
	"sync/atomic"

	ics23 "github.com/confio/ics23/go"
	"github.com/sei-protocol/sei-chain/sei-db/config"
	"github.com/sei-protocol/sei-chain/sei-db/proto"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/flatkv"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/memiavl"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/migration"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/types"
	"github.com/sei-protocol/seilog"
	db "github.com/tendermint/tm-db"
)

var logger = seilog.NewLogger("db", "state-db", "sc", "composite")

// For backward compatibility purpose reuse current interface
var _ types.Committer = (*CompositeCommitStore)(nil)

// CompositeCommitStore manages multiple commit store backends (Cosmos/memiavl and FlatKV)
// and routes operations based on the configured migration strategy.
type CompositeCommitStore struct {
	// The memIAVL backend. Will be nil after all data is migrated to flatkv.
	memIAVL *memiavl.CommitStore

	// The flatKV backend. Will be nil if migration to flatKV has not yet started.
	flatKV flatkv.Store

	// Manages routing of traffic between the memiavl and flatkv backends.
	// Built (and rebuilt) inside LoadVersion against the just-opened
	// backends so that lazily-eager constructors like
	// NewMemiavlMigrationIterator see a non-nil memiavl DB.
	router migration.Router

	// ctx is the constructor's context. Each invocation of buildRouter
	// derives a per-router child context from it and stores the
	// corresponding cancel function in routerCancel; cancelling that
	// child stops any background goroutines owned by the current
	// router (today: the MigrationMetrics boundary-snapshot loop)
	// without affecting any unrelated work that shares cs.ctx.
	ctx context.Context

	// routerCancel cancels the child context handed to the current
	// router. Called before installing a new router on reload, and on
	// Close. Nil before the first LoadVersion and after Close.
	routerCancel context.CancelFunc

	// homeDir is the base directory for the store
	homeDir string

	// config holds the store configuration
	config config.StateCommitConfig

	// latticeAppendLatched is a sticky one-way flag: once it transitions
	// to true, LastCommitInfo and WorkingCommitInfo unconditionally
	// append the evm_lattice StoreInfo without consulting the on-disk
	// migration metadata again. The flag protects the AppHash continuity
	// invariant for a live MemiavlOnly -> MigrateEVM transition: while the
	// migration boundary on flatkv is still NotStarted, the lattice must
	// be suppressed so the post-restart LastCommitInfo matches the
	// pre-restart memiavl-only AppHash at the same height. Once the
	// boundary advances (or the migration completes), the gate latches
	// and subsequent calls skip the flatkv read. See shouldAppendLatticeHash.
	latticeAppendLatched atomic.Bool
}

// NewCompositeCommitStore creates a new composite commit store.
// Note: The store is NOT opened yet. Call LoadVersion to open and initialize the DBs.
// This matches the memiavl.NewCommitStore pattern.
func NewCompositeCommitStore(
	ctx context.Context,
	homeDir string,
	cfg config.StateCommitConfig,
) (*CompositeCommitStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initialize records the set of child store names that should exist on
// the memiavl backend the first time it is opened. In mixed-DB modes
// names must be members of keys.MemIAVLStoreKeys.
func (cs *CompositeCommitStore) Initialize(initialStores []string) error {
	_ = "STUB: not implemented"
	return nil
}

// validateInitialStores enforces the rules described on Initialize.
func validateInitialStores(mode config.WriteMode, initialStores []string) error {
	_ = "STUB: not implemented"
	return nil
}

// CleanupCrashArtifacts removes temporary/orphaned files left by a
// previous process crash (e.g. FlatKV readonly-* working directories).
// Must be called once at process startup, before any read-only clones
// are created. Any writer lock acquired during cleanup is retained for
// the subsequent LoadVersion(..., false) call.
func (cs *CompositeCommitStore) CleanupCrashArtifacts() error {
	_ = "STUB: not implemented"
	return nil
}

// SetInitialVersion seeds every active backend so that the next Commit
// produces initialVersion. Called from cosmos-sdk BaseApp.InitChain on
// fresh genesis.
func (cs *CompositeCommitStore) SetInitialVersion(initialVersion int64) error {
	_ = "STUB: not implemented"
	return nil
}

// LoadVersion opens the database at the given version (0 = latest).
// When readOnly is true an isolated composite store is returned.
func (cs *CompositeCommitStore) LoadVersion(targetVersion int64, readOnly bool) (committer types.Committer, retErr error) {
	_ = "STUB: not implemented"
	return *new(types.Committer), nil
}

// Defensive: in practice memiavl always returns
// *CommitStore, but if some future implementation does not,
// close whatever was returned so we do not leak it.

// Build a per-handle composite with its own router. Without
// this the read-only handle has cs.router == nil and every
// read-side method nil-dereferences on first call. The new
// composite inherits cs.ctx so cancellation of the parent
// context cascades, but buildRouter installs its own child
// cancel so closing this handle does not affect the parent.

// Reassign the freshly-loaded backends. flatkv.Store.LoadVersion
// is documented to return the receiver on the writable path, but
// the field is an interface (tests inject mocks via cs.flatKV =
// mock); honoring the return value future-proofs against an
// implementation that returns a swapped instance.

// Migration-entry seeding: turning on a non-MemiavlOnly mode on a
// chain that has been running on MemiavlOnly leaves memiavl at
// version N while flatkv starts fresh at version 0. Bring flatkv
// into lockstep so the next composite commit produces matching
// versions on both backends. Only runs at load-latest; targeted
// loads stay strict so a mismatch is surfaced loudly.

// When loading latest (targetVersion==0), a crash between the
// sequential cosmos and EVM commits can leave the backends at
// different versions. Detect the mismatch and roll the ahead
// backend back so both restart from a consistent point.

// buildRouter constructs the migration router against the currently-opened
// backends and assigns it to cs.router. Must be called after memIAVL and
// flatKV (if any) have been opened via LoadVersion.
func (cs *CompositeCommitStore) buildRouter() error { _ = "STUB: not implemented"; return nil }

// ApplyChangeSets applies changesets to the appropriate backends based on config.
func (cs *CompositeCommitStore) ApplyChangeSets(changesets []*proto.NamedChangeSet) error {
	_ = "STUB: not implemented"
	return nil
}

// ApplyUpgrades applies store upgrades (only applicable to memIAVL Cosmos backend). Data in
// flatKV is not affected by this method.
func (cs *CompositeCommitStore) ApplyUpgrades(upgrades []*proto.TreeNameUpgrade) error {
	_ = "STUB: not implemented"
	return nil
}

// Commit commits the current state to all active backends
func (cs *CompositeCommitStore) Commit() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// reconcileVersions checks whether the cosmos and EVM backends are at the
// same version after loading latest. A crash between the sequential Commit
// calls can leave one backend one version ahead. When a mismatch is found
// and both backends have committed at least once (version > 0), the ahead
// backend is rolled back to the behind version. Rollback truncates the WAL
// so the correction survives subsequent restarts.
func (cs *CompositeCommitStore) reconcileVersions() error { _ = "STUB: not implemented"; return nil }

// Nothing to reconcile if one of the backends is not present.

// Skip reconciliation when either backend is at version 0 (fresh
// initialization / migration), since that is not a crash artifact.

// Version returns the current version
func (cs *CompositeCommitStore) Version() int64 { _ = "STUB: not implemented"; return 0 }

// GetLatestVersion returns the highest committed version.
func (cs *CompositeCommitStore) GetLatestVersion() (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// shouldAppendLatticeHash reports whether LastCommitInfo and
// WorkingCommitInfo should append the synthetic evm_lattice StoreInfo to
// the commit info.
//
// The composite store contributes an evm_lattice entry to every commit
// info whenever the flatkv backend is participating in the AppHash. The
// one exception is the brief NotStarted window of a live
// MemiavlOnly -> MigrateEVM transition: between the time flatkv is
// opened (LoadVersion seeds it to memiavl's version) and the first
// post-transition commit (which advances the migration boundary), the
// chain's stored AppHash at the just-loaded height still reflects the
// memiavl-only era. Adding an evm_lattice entry at that height would
// silently change the AppHash that Tendermint already accepted and
// fail the handshake.
//
// To preserve continuity exactly through that window — and not a moment
// longer — the gate consults the on-disk migration metadata on flatkv:
//
//   - flatKV == nil (MemiavlOnly): never append; flatkv is not part of
//     the merkle root at all.
//   - WriteMode != MigrateEVM (EVMMigrated, MigrateAllButBank,
//     AllMigratedButBank, MigrateBank, FlatKVOnly, TestOnlyDualWrite):
//     always append. These modes either entered with the lattice baked
//     into their genesis or descend from a flatkv-bearing predecessor
//     that already committed it; there is no memiavl-only prior
//     AppHash to be inconsistent with. By design no operator will jump
//     a memiavl-only chain straight into one of these modes.
//   - MigrateEVM: append iff the migration has progressed past
//     MigrationNotStarted. We treat the boundary as "started" if
//     MigrationBoundaryKey is present and decodes to any status other
//     than MigrationNotStarted, OR if MigrationVersionKey is present.
//     The latter is what survives a completion block: the manager
//     atomically deletes MigrationBoundaryKey and writes
//     MigrationVersionKey, so checking both keys covers the entire
//     post-NotStarted lifecycle.
//
// The result is sticky once true. After the very first observation
// that the gate has opened, latticeAppendLatched is set and subsequent
// calls return immediately without re-reading flatkv. This both avoids
// per-call DB work on the hot commit-info path and guarantees a
// consistent answer across the completion block on which the on-disk
// signal hops from MigrationBoundaryKey to MigrationVersionKey.
func (cs *CompositeCommitStore) shouldAppendLatticeHash() bool {
	_ = "STUB: not implemented"
	return false
}

// Consensus-critical: a corrupt boundary record means we
// cannot tell whether the lattice should be in the AppHash.
// Failing loud is the only safe option.

// appendEvmLatticeHash returns a new CommitInfo with the EVM lattice hash
// appended, without mutating the original. Returns the original unchanged
// when flatKV is not present.
func (cs *CompositeCommitStore) appendEvmLatticeHash(ci *proto.CommitInfo, evmHash []byte) *proto.CommitInfo {
	_ = "STUB: not implemented"
	return nil
}

// WorkingCommitInfo returns the working commit info
func (cs *CompositeCommitStore) WorkingCommitInfo() *proto.CommitInfo {
	_ = "STUB: not implemented"
	return nil
}

// LastCommitInfo returns the last commit info
func (cs *CompositeCommitStore) LastCommitInfo() *proto.CommitInfo {
	_ = "STUB: not implemented"
	return nil
}

// GetChildStoreByName returns the underlying child store by module name.
// Panics if the store name is not supported by the current write mode.
//
// The reserved migration.MigrationStore tree is always rejected,
// regardless of mode: it is owned by the migration workflow.
func (cs *CompositeCommitStore) GetChildStoreByName(name string) types.CommitKVStore {
	_ = "STUB: not implemented"
	return *new(types.CommitKVStore)
}

// In MemiavlOnly mode, check to see if the tree exists. Required to support legacy test apps
// that use non-standard store names.

// FlatKV only mode can support arbitrary store names. Otherwise, require the store to be in the canonical list.

// Copy returns an in-memory snapshot, or nil when flatkv is engaged
// (no in-memory primitive; a partial snapshot would miss EVM state).
func (cs *CompositeCommitStore) Copy() types.Committer {
	_ = "STUB: not implemented"
	return *new(types.Committer)
}

// ReleaseSnapshotRefs releases refs held by a copied in-memory snapshot without
// closing DB-level resources shared with the live store.
func (cs *CompositeCommitStore) ReleaseSnapshotRefs() error { _ = "STUB: not implemented"; return nil }

// Rollback rolls back to the specified version
func (cs *CompositeCommitStore) Rollback(targetVersion int64) error {
	_ = "STUB: not implemented"
	return nil
}

// Exporter returns an exporter for state sync
func (cs *CompositeCommitStore) Exporter(version int64) (types.Exporter, error) {
	_ = "STUB: not implemented"
	return *new(types.Exporter), nil
}

// Importer returns an importer for state sync
func (cs *CompositeCommitStore) Importer(version int64) (types.Importer, error) {
	_ = "STUB: not implemented"
	return *new(types.Importer), nil
}

// Close closes all backends
func (cs *CompositeCommitStore) Close() error { _ = "STUB: not implemented"; return nil }

func (cs *CompositeCommitStore) Get(store string, key []byte) (value []byte, ok bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (cs *CompositeCommitStore) GetProof(store string, key []byte) (*ics23.CommitmentProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs *CompositeCommitStore) Has(store string, key []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cs *CompositeCommitStore) Iterator(store string, start []byte, end []byte, ascending bool) (db.Iterator, error) {
	_ = "STUB: not implemented"
	return *new(db.Iterator), nil
}

package operations

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/flatkv"
	sctypes "github.com/sei-protocol/sei-chain/sei-db/state_db/sc/types"
	"github.com/spf13/cobra"
)

// translatorBatchSize bounds how many memiavl key/value pairs we hand to a
// single flatkv.ImportTranslator.Translate call. Batching amortizes the
// per-call classifyAndPrefix map allocations across many keys without
// growing ImportTranslator's account-buffer memory beyond what an unbatched
// stream would already need.
//
// Distinct from flatkv.importBatchSize, which is the per-DB-worker flush
// threshold (in already-translated physical pairs); the two constants tune
// different stages of the pipeline.
const translatorBatchSize = 2048

// ImportFlatKVFromMemiavlCmd imports selected memiavl modules into FlatKV.
//
// Initial production scope is intentionally narrow: only the evm module is
// accepted. Non-EVM modules remain in memiavl and are not copied into FlatKV.
// Importing resets FlatKV and replaces it with the selected memiavl data; the
// CLI refuses to run over existing FlatKV data unless --force is supplied.
func ImportFlatKVFromMemiavlCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func resolveSeiHome(homeDir, dataDir string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func normalizeImportModules(modules []string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// importerErr surfaces any pipeline error the FlatKV importer's worker
// goroutines have already recorded, so the import loop can fail-fast
// between exporter reads instead of waiting until Close. The anonymous
// interface assertion (rather than a concrete *flatkv.KVImporter type
// switch) lets any future Importer impl opt into mid-stream error
// reporting just by adding Err() error to its method set, without
// touching this helper.
func importerErr(importer sctypes.Importer) error { _ = "STUB: not implemented"; return nil }

// emitPairs forwards translator output to the FlatKV importer, returning the
// number of pairs written.
func emitPairs(importer sctypes.Importer, pairs []flatkv.PhysicalKVPair, height int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

func importMemiavlModulesToFlatKV(ctx context.Context, homeDir string, modules []string, height int64, force bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Refuse mismatched heights. If we wrote FlatKV at H < memiavlLatest,
// the next GIGA_STORAGE startup would call
// CompositeCommitStore.reconcileVersions (see
// sei-db/state_db/sc/composite/store.go) and silently roll memiavl
// back to H, truncating every cosmos block in (H, memiavlLatest].
// H > memiavlLatest is unreachable in practice (the memiavl exporter
// would error a few lines below) but caught here for a clearer
// message. Operators who genuinely want a non-latest H must first
// roll memiavl back to H themselves; this CLI deliberately does NOT
// roll memiavl back on their behalf because "import" is a one-way,
// abortable operation and should never be a hidden gateway into a
// destructive cosmos rollback.

//nolint:gosec // height range checked above

// On the failure path we must NOT finalize: KVImporter.Close otherwise
// commits whatever pairs were already buffered, leaving FlatKV at the
// target version with only a partial copy of the source state. Route
// errors through Abort instead, which records the failure on the
// importer and then drains workers without writing a snapshot. On the
// success path the explicit Close below has already run, so the
// deferred Close here is just an idempotent safety net.

// err path: do NOT call Close, which would finalize the partial
// import (see KVImporter.Close docstring). If the type assertion
// fails (future Importer impl), leave the pipeline to GC -- a
// leak strictly beats silently committing a half-imported snapshot.

// acceptCurrent caches whether the current module (batch.Name) is in
// moduleSet so the per-pair SnapshotNode arm doesn't repeat the map
// lookup for every key emitted by the exporter. It's recomputed once
// per module switch in the `case string:` arm below.

// AddModule takes the source module name (here the memiavl
// module being read), not the destination store name. On
// *flatkv.KVImporter this is currently a no-op, but
// telemetry-/log-bearing implementations downstream will
// attribute the import to batch.Name rather than
// hard-coding it to "flatkv".

// EVM-only choke point. normalizeImportModules already rejects
// non-EVM module names at the CLI boundary, so today this skip
// is defense-in-depth. If a future expansion adds another
// module to the allow-list, this `continue` is what keeps that
// module's pairs out of the importer -- the flatkv store does
// not have a routing path for non-EVM physical keys yet, and
// silently accepting them would land them in the legacyDB
// bucket. Any allow-list change MUST be paired with a flatkv
// routePhysicalKey extension; otherwise leave this skip alone.

// MemiavlLatestVersionCmd is the read-only companion to ImportFlatKVFromMemiavlCmd:
// it reports the latest committed memiavl version of a stopped node so an
// orchestration script can pick a single import height across a multi-validator
// cluster. Lives in this file (rather than a standalone *_cmd.go) because it
// shares resolveSeiHome with the import command and exists solely to support
// that workflow -- see integration_test/contracts/import_flatkv_evm_cluster.sh
// for the call site, which reads each validator's version after pkill, picks
// the minimum, and rolls back any node that committed extra blocks before
// running the offline import.
func MemiavlLatestVersionCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

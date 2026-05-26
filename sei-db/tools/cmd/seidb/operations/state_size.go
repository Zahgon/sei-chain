package operations

import (
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/memiavl"
	"github.com/sei-protocol/sei-chain/sei-db/tools/utils"
	"github.com/spf13/cobra"
)

func StateSizeCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// FlatKV integration: optional FlatKV data directory. When non-empty (or
// when a sibling flatkv/ dir is auto-detected next to --db-dir) the tool
// also scans FlatKV and folds the result into the same console output
// and the same DynamoDB batch as the memIAVL module rows.

// DynamoDB export flags

func executeStateSize(cmd *cobra.Command, _ []string) { _ = "STUB: not implemented"; return }

// Optionally scan FlatKV at the same requested height. We only bother
// when --module is empty or "evm" because FlatKV in production holds
// only evm keys (anything else is bucketed into the "legacy" DB).

// resolveFlatKVDir returns the FlatKV directory to scan, if any.
//
//   - if --flatkv-dir was supplied explicitly, it is returned as-is (the
//     caller is responsible for the path being valid).
//   - otherwise the tool auto-detects a sibling "flatkv/" directory next
//     to --db-dir (e.g. <home>/data/committer.db -> <home>/data/flatkv),
//     which is the standard layout on a seid shadow node. Returns "" if
//     no such sibling exists.
func resolveFlatKVDir(flatkvDir, dbDir string) string { _ = "STUB: not implemented"; return "" }

// maybeCollectFlatKV resolves the FlatKV directory and, if present and the
// caller's --module filter is compatible, opens a read-only clone of the
// FlatKV store and scans it.
//
// On any error (dir missing, snapshot unavailable, open failure) we log the
// reason and return a nil result so the memIAVL path still succeeds. FlatKV
// analysis is strictly additive; failing here must never take down the
// existing state-size workflow.
func maybeCollectFlatKV(flatkvDir, dbDir, module string, height int64) (*FlatKVStateSizeResult, int64) {
	_ = "STUB: not implemented"
	return nil, 0
}

// collectModuleStats collects all the statistics for a module
func collectModuleStats(tree *memiavl.Tree, moduleName string) *ModuleResult {
	_ = "STUB: not implemented"
	return nil
}

// Scan the tree to collect statistics

//nolint:gosec

// Handle EVM contract analysis

// Progress every 10M keys. The largest module (evm) holds
// hundreds of millions of leaves; a 1M-per-line cadence here
// was drowning the actual report in 700+ progress lines.

// Limit to top 100 contracts by total size

// limitToTopContracts keeps only the top N contracts by total size
func limitToTopContracts(contracts map[string]*utils.ContractSizeEntry, limit int) map[string]*utils.ContractSizeEntry {
	_ = "STUB: not implemented"
	return nil
}

// Convert to slice for sorting

// Sort by total size in descending order

// Keep only top N

// ModuleResult holds the complete analysis results for a single module
type ModuleResult struct {
	ModuleName     string
	TotalNumKeys   uint64
	TotalKeySize   uint64
	TotalValueSize uint64
	TotalSize      uint64
	PrefixSizes    map[string]*utils.PrefixSize
	ContractSizes  map[string]*utils.ContractSizeEntry
}

// collectAllModuleData scans all modules and collects statistics in memory
func collectAllModuleData(module string, db *memiavl.DB) map[string]*ModuleResult {
	_ = "STUB: not implemented"
	return nil
}

// Collect statistics directly into ModuleResult

// Store in memory (result is already a ModuleResult)

// exportResultsToDynamoDB exports the collected memIAVL module results plus
// any additional pre-built analyses (e.g. FlatKV) to DynamoDB as a single
// batch. The metadata latest-height record is keyed off the memIAVL height
// because it remains the canonical "observation height" even when FlatKV
// resolved to a slightly older snapshot.
func exportResultsToDynamoDB(
	moduleResults map[string]*ModuleResult,
	extras []*utils.StateSizeAnalysis,
	height int64,
	tableName, awsRegion string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// printResultsToConsole prints the collected results to console.
//
// PrefixSizes is keyed by hex prefix byte (e.g. "03", "0A"), not by module
// name, so previous code that indexed this map with the module name always
// panicked on nil deref the first time the console path was taken. We now
// marshal the entire map per module, which is what "prefix breakdown" was
// always meant to surface.
//
// Modules are emitted in alphabetical order so successive runs produce
// diffable output. The top-contracts table is skipped for modules without
// any 0x03 entries to avoid printing empty table headers for every non-evm
// module.
func printResultsToConsole(moduleResults map[string]*ModuleResult) {
	_ = "STUB: not implemented"
	return
}

// createStateSizeAnalysis creates a new StateSizeAnalysis from ModuleResult
func createStateSizeAnalysis(blockHeight int64, moduleName string, result *ModuleResult) *utils.StateSizeAnalysis {
	_ = "STUB: not implemented"
	// Convert raw data to JSON strings for DynamoDB storage
	return nil
}

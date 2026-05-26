package operations

import (
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/flatkv"
	"github.com/sei-protocol/sei-chain/sei-db/tools/utils"
)

// flatkvAnalysisModuleName is the logical module name used for the FlatKV
// row in the shared DynamoDB state-size table. Consumers key off this name
// to distinguish FlatKV from memIAVL module rows.
const flatkvAnalysisModuleName = "flatkv"

// FlatKVStateSizeResult holds the complete analysis of a FlatKV store.
type FlatKVStateSizeResult struct {
	// Total holds the aggregate size stats across every physical row.
	Total FlatKVDBSize

	// Per-DB breakdown (account, code, storage, legacy).
	DBSizes map[string]*FlatKVDBSize

	// Top EVM contracts by storage size.
	ContractSizes map[string]*utils.ContractSizeEntry
}

// FlatKVDBSize holds size stats for one logical DB.
type FlatKVDBSize struct {
	NumKeys   uint64
	KeySize   uint64
	ValueSize uint64
	TotalSize uint64
}

// collectFlatKVStateSize iterates every physical row in the FlatKV store and
// aggregates size stats per logical DB, plus a top-100 EVM contract table.
func collectFlatKVStateSize(store *flatkv.CommitStore) (*FlatKVStateSizeResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// classifyFlatKVPhysicalKey determines which logical DB a physical key
// belongs to. Physical format: "<module>/" + type_prefix_byte + stripped_key.
// Non-evm modules and evm keys with an unrecognised type prefix are bucketed
// into "legacy". The kind switch mirrors CommitStore.routePhysicalKey so the
// classification stays in sync with FlatKV's actual write routing.
func classifyFlatKVPhysicalKey(key []byte) string { _ = "STUB: not implemented"; return "" }

// extractFlatKVContractAddress extracts the hex address from an evm storage
// physical key. Physical format: "evm/" + 0x03 + addr(20) + slot(32).
func extractFlatKVContractAddress(key []byte) string { _ = "STUB: not implemented"; return "" }

func limitFlatKVTopContracts(contracts map[string]*utils.ContractSizeEntry, limit int) map[string]*utils.ContractSizeEntry {
	_ = "STUB: not implemented"
	return nil
}

// printFlatKVResults prints a FlatKV section to stdout formatted to match
// the surrounding memIAVL module output from state_size.go.
func printFlatKVResults(r *FlatKVStateSizeResult, height int64) { _ = "STUB: not implemented"; return }

// flatkvStateSizeAnalysis packages a FlatKV scan result as a
// *utils.StateSizeAnalysis so it can be pushed to DynamoDB alongside the
// memIAVL module rows in a single batch.
//
// The per-DB breakdown is rendered into PrefixBreakdown using the same JSON
// map shape memIAVL uses ({"<bucket>": PrefixSize}), so downstream consumers
// can parse both module types uniformly.
func flatkvStateSizeAnalysis(r *FlatKVStateSizeResult, height int64) *utils.StateSizeAnalysis {
	_ = "STUB: not implemented"
	return nil
}

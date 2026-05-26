package operations

import (
	"bufio"
	"os"

	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/flatkv"
	"github.com/spf13/cobra"
)

const (
	flatkvBucketAccount = "account"
	flatkvBucketCode    = "code"
	flatkvBucketStorage = "storage"
	flatkvBucketLegacy  = "legacy"
)

// flatkvBucketOrder lists the logical bucket names in the same order
// RawGlobalIterator returns them (account → code → storage → legacy). Keeping
// this as the single source of truth lets us loop once for both CLI
// validation and per-bucket file allocation.
var flatkvBucketOrder = []string{flatkvBucketAccount, flatkvBucketCode, flatkvBucketStorage, flatkvBucketLegacy}

// DumpFlatKVCmd dumps every (physical key, value) pair of a FlatKV store
// into per-bucket files, formatted to match dump-iavl so the same diff
// tooling works on both.
func DumpFlatKVCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func executeDumpFlatKV(cmd *cobra.Command, _ []string) { _ = "STUB: not implemented"; return }

func isFlatKVBucket(name string) bool { _ = "STUB: not implemented"; return false }

// DumpFlatKVData opens a read-only clone of a FlatKV store at the requested
// version and writes every (physical key, value) pair into per-bucket files
// under outputDir. Each file mirrors the dump-iavl format so downstream
// diff tooling can be shared:
//
//	Bucket <name> at version <V>
//	Key: <HEX>, Value: <HEX>
//	...
//
// Physical keys are emitted verbatim, including their "<module>/" + type
// prefix header, because they are not byte-for-byte comparable with
// memIAVL logical keys anyway (different type prefixes per domain). The
// FlatKV metadataDB and the per-DB _meta/* rows are intentionally excluded:
// they are internal bookkeeping and RawGlobalIterator already filters the
// per-DB ones for us.
func DumpFlatKVData(dbDir, outputDir string, height int64, bucket string) error {
	_ = "STUB: not implemented"
	return nil
}

// dumpFlatKVFromStore is the core scan+write path, split out so tests can
// exercise it against an in-memory store without going through the
// snapshot clone machinery used by the CLI.
func dumpFlatKVFromStore(store *flatkv.CommitStore, outputDir string, version int64, bucket string) error {
	_ = "STUB: not implemented"
	return nil
}

func flushAndCloseBucketWriters(files map[string]*os.File, writers map[string]*bufio.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

// openBucketWriters creates per-bucket output files inside outputDir. When
// bucket != "" only that bucket's writer is populated; unselected buckets
// are absent from the returned maps, which the scan loop treats as "skip
// writes for this key but keep iterating" (the iterator is sequential and
// cannot cheaply skip an entire sub-DB without package-private access).
func openBucketWriters(outputDir string, version int64, bucket string) (map[string]*os.File, map[string]*bufio.Writer, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

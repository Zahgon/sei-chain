package bench

import (
	"time"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
	"github.com/sei-protocol/sei-chain/sei-db/tools/utils"
)

// writeToDBConcurrently generates random write load against the db
// Given kv pairs (randomly shuffled), the version, batch size, it will spin up `concurrency` goroutines
// each of which is assigned to a portion of the kv data and writes to db in `batchSize` batches.
// It maintains a `latencies` channel which aggregates all the latencies
func writeToDBConcurrently(db types.StateStore, allKVs []utils.KeyValuePair, concurrency int, version int64, batchSize int) []time.Duration {
	_ = "STUB: not implemented"
	return nil
}

// Add key-value pairs to the batch up to batchSize

// No store key for benchmarks

// BenchmarkDBWrite measures random write performance of the db
// Given an input dir containing all the raw kv data, it writes to the db one version after another
func BenchmarkDBWrite(db types.StateStore, inputKVDir string, numVersions int, concurrency int, batchSize int) {
	_ = "STUB: not implemented"
	return
}

// Write each version sequentially

// Write shuffled entries to RocksDB concurrently

// Latencies per version

// Log throughput

// readFromDBConcurrently generates random read load against the db
// Given kv pairs (randomly shuffled), numVersions, it will spin up `concurrency` goroutines
// that randomly select a version, key and query the db.
// It only performs `maxOps“ random reads and maintains a `latencies` channel which aggregates all the latencies.
func readFromDBConcurrently(db types.StateStore, allKVs []utils.KeyValuePair, numVersions int, concurrency int, maxOps int64) []time.Duration {
	_ = "STUB: not implemented"
	return nil
}

// Each goroutine will handle reading a subset of kv pairs

// Randomly pick a version and retrieve its column family handle

// Randomly pick a key-value pair to read

// No store key for benchmarks

// BenchmarkDBRead measures random read performance of the db
// Given an input dir containing all the raw kv data, it generates random read load and measures performance.
func BenchmarkDBRead(db types.StateStore, inputKVDir string, numVersions int, concurrency int, maxOps int64) {
	_ = "STUB: not implemented"
	return
}

// Log throughput

// Sort latencies for percentile calculations

// Calculate average latency

// forwardIterateDBConcurrently generates forward iteration load against the db
// Given kv pairs (randomly shuffled), numVersions, it will spin up `concurrency` goroutines
// that randomly select a version, key, seeks to that key and starts a forward iteration for at most `numIterationSteps` steps.
// It only performs `maxOps“ forward iterations and maintains a `latencies` channel which aggregates all the latencies.
func forwardIterateDBConcurrently(db types.StateStore, allKVs []utils.KeyValuePair, numVersions int, concurrency int, numIterationSteps int, maxOps int64) ([]time.Duration, int) {
	_ = "STUB: not implemented"
	return nil, 0
}

// Randomly pick a version and retrieve its column family handle

// Randomly pick a key-value pair to seek to

// No end key since we iterate for fixed numIterationSteps steps

// BenchmarkDBForwardIteration measures forward iteration performance of the db
// Given an input dir containing all the raw kv data, it selects a random key, forward iterates and measures performance.
func BenchmarkDBForwardIteration(db types.StateStore, inputKVDir string, numVersions int, concurrency int, maxOps int64, iterationSteps int) {
	_ = "STUB: not implemented"
	return
}

// Log throughput

// Calculate average latency

// reverseIterateDBConcurrently generates reverse iteration load against the db
// Given kv pairs (randomly shuffled), numVersions, it will spin up `concurrency` goroutines
// that randomly select a version, key, seeks to that key and starts a reverse iteration for at most `numIterationSteps` steps.
// It only performs `maxOps“ reverse iterations and maintains a `latencies` channel which aggregates all the latencies.
func reverseIterateDBConcurrently(db types.StateStore, allKVs []utils.KeyValuePair, numVersions int, concurrency int, numIterationSteps int, maxOps int64) ([]time.Duration, int) {
	_ = "STUB: not implemented"
	return nil, 0
}

// Randomly pick a version and retrieve its column family handle

// Randomly pick a key-value pair to seek to

// No start key since we iterate for fixed numIterationSteps steps

// BenchmarkDBReverseIteration measures reverse iteration performance of the db
// Given an input dir containing all the raw kv data, it selects a random key, reverse iterates and measures performance.
func BenchmarkDBReverseIteration(db types.StateStore, inputKVDir string, numVersions int, concurrency int, maxOps int64, iterationSteps int) {
	_ = "STUB: not implemented"
	return
}

// Log throughput

// Calculate average latency

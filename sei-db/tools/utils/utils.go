package utils

import (
	"io"
	"os"
	"sync"
	"time"

	dbm "github.com/tendermint/tm-db"
)

type KeyValuePair struct {
	Key   []byte `json:"key"`
	Value []byte `json:"value"`
}

// Opens application db
func OpenDB(dir string) (dbm.DB, error) { _ = "STUB: not implemented"; return *new(dbm.DB), nil }

// TODO: doesn't work on windows!

// Reads raw keys / values from a file
func ReadKVEntriesFromFile(filename string) ([]KeyValuePair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readByteSlice(r io.Reader) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// NOTE: Assumes latencies is sorted
// CalculatePercentile calculates latencies percentile
func CalculatePercentile(latencies []time.Duration, percentile float64) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Picks random file from input kv dir and updates processedFiles Map with it
func PickRandomKVFile(inputKVDir string, processedFiles *sync.Map) string {
	_ = "STUB: not implemented"
	return ""
}

func ListAllFiles(dir string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// Extract file nams from input KV dir

func LoadAndShuffleKV(inputDir string, concurrency int) ([]KeyValuePair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Start worker goroutines

// Safely append the kvEntries to allKVs

// Send file names to filesChan

// Wait for all workers to finish

func CreateFile(outputDir string, fileName string) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

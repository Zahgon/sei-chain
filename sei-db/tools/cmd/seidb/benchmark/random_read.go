package benchmark

import (
	"github.com/spf13/cobra"
)

func DBRandomReadCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func executeRandomRead(cmd *cobra.Command, args []string) { _ = "STUB: not implemented"; return }

// BenchmarkRead read latencies and throughput of db backend
func DBRandomRead(inputKVDir string, numVersions int, outputDir string, dbBackend string, concurrency int, maxOps int64) {
	_ = "STUB: not implemented"
	// Create output directory
	return
}

// Iterate over files in directory

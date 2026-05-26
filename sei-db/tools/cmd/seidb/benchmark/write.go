package benchmark

import (
	"github.com/spf13/cobra"
)

func DBWriteCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func executeWrite(cmd *cobra.Command, args []string) { _ = "STUB: not implemented"; return }

// BenchmarkWrite write latencies and throughput of db backend
func DBWrite(inputKVDir string, numVersions int, outputDir string, dbBackend string, concurrency int, batchSize int) {
	_ = "STUB: not implemented"
	// Create output directory
	return
}

// Iterate over files in directory

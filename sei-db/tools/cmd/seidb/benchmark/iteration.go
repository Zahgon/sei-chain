package benchmark

import (
	"github.com/spf13/cobra"
)

func DBIterationCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func executeForwardIteration(cmd *cobra.Command, args []string) { _ = "STUB: not implemented"; return }

// BenchmarkDBIteration read latencies and throughput of db backend
func DBIteration(inputKVDir string, numVersions int, outputDir string, dbBackend string, concurrency int, maxOps int64, iterationSteps int) {
	_ = "STUB: not implemented"
	// Iterate over db at directory
	return
}

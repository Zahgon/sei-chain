package benchmark

import (
	"github.com/spf13/cobra"
)

func DBReverseIterationCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func executeReverseIteration(cmd *cobra.Command, args []string) { _ = "STUB: not implemented"; return }

// BenchmarkDBReverseIteration reverse iteration performance of db backend
func DBReverseIteration(inputKVDir string, numVersions int, outputDir string, dbBackend string, concurrency int, maxOps int64, iterationSteps int) {
	_ = "STUB: not implemented"
	// Reverse Iterate over db at directory
	return
}

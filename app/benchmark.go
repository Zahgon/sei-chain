package app

import (
	"context"
	"time"

	"github.com/sei-protocol/sei-chain/app/benchmark"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// InitBenchmark initializes the benchmark system with the configured scenarios.
// This is called during app initialization when the benchmark build tag is enabled.
func (app *App) InitBenchmark(ctx context.Context, chainID string, evmChainID int64) {
	_ = "STUB: not implemented"
	// Defensive check: prevent benchmarking on live chains
	return
}

// ProcessBenchmarkReceipts extracts receipts from the block and forwards them to
// the benchmark system for deployment tracking during the setup phase.
func (app *App) ProcessBenchmarkReceipts(ctx sdk.Context) { _ = "STUB: not implemented"; return }

// BenchmarkLogger returns the benchmark logger for recording timing metrics.
// Returns nil if benchmark mode is not enabled.
func (app *App) BenchmarkLogger() *benchmark.Logger { _ = "STUB: not implemented"; return nil }

// RecordBenchmarkCommitTime records the commit duration for benchmark metrics.
func (app *App) RecordBenchmarkCommitTime(duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

// StartBenchmarkBlockProcessing marks the start of block processing for timing.
func (app *App) StartBenchmarkBlockProcessing() { _ = "STUB: not implemented"; return }

// EndBenchmarkBlockProcessing marks the end of block processing for timing.
func (app *App) EndBenchmarkBlockProcessing() { _ = "STUB: not implemented"; return }

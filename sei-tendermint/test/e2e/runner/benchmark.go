package main

import (
	"context"
	"time"

	e2e "github.com/sei-protocol/sei-chain/sei-tendermint/test/e2e/pkg"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("tendermint", "test", "e2e", "runner")

// Benchmark is a simple function for fetching, calculating and printing
// the following metrics:
// 1. Average block production time
// 2. Block interval standard deviation
// 3. Max block interval (slowest block)
// 4. Min block interval (fastest block)
//
// Metrics are based of the `benchmarkLength`, the amount of consecutive blocks
// sampled from in the testnet
func Benchmark(ctx context.Context, testnet *e2e.Testnet, benchmarkLength int64) error {
	_ = "STUB: not implemented"
	return nil
}

// wait for the length of the benchmark period in blocks to pass. We allow 5 seconds for each block
// which should be sufficient.

// fetch a sample of blocks

// slice into time intervals and collate data

// populate data

// print and return

func (t *testnetStats) populateTxns(blocks []*types.BlockMeta) { _ = "STUB: not implemented"; return }

type testnetStats struct {
	startHeight int64
	endHeight   int64

	benchmarkLength int64
	numtxns         int64
	totalTime       time.Duration
	// average time to produce a block
	mean time.Duration
	// standard deviation of block production
	std float64
	// longest time to produce a block
	max time.Duration
	// shortest time to produce a block
	min time.Duration
}

func (t *testnetStats) getReportJSON(net *e2e.Testnet) string { _ = "STUB: not implemented"; return "" }

func (t *testnetStats) String() string { _ = "STUB: not implemented"; return "" }

// fetchBlockChainSample waits for `benchmarkLength` amount of blocks to pass, fetching
// all of the headers for these blocks from an archive node and returning it.
func fetchBlockChainSample(ctx context.Context, testnet *e2e.Testnet, benchmarkLength int64) ([]*types.BlockMeta, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Find the first archive node
}

// find the latest height

// Fetch blocks

// fetch the blockchain metas. Currently we can only fetch 20 at a time

// we receive blocks in descending order so we have to add them in reverse

func splitIntoBlockIntervals(blocks []*types.BlockMeta) []time.Duration {
	_ = "STUB: not implemented"
	return nil
}

// skip the first block

func extractTestnetStats(intervals []time.Duration) testnetStats {
	_ = "STUB: not implemented"
	return *new(testnetStats)
}

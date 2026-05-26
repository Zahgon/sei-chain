package cryptosim

import (
	"context"
)

// A builder for blocks of transactions.
type blockBuilder struct {
	ctx context.Context

	config *CryptoSimConfig

	// Metrics for the benchmark.
	metrics *CryptosimMetrics

	// Produces random data.
	dataGenerator *DataGenerator

	// Blocks are sent to this channel.
	blocksChan chan *block

	// The next block number to be used.
	nextBlockNumber int64
}

// Asyncronously produces blocks of transactions.
func NewBlockBuilder(
	ctx context.Context,
	config *CryptoSimConfig,
	metrics *CryptosimMetrics,
	dataGenerator *DataGenerator,
) *blockBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Starts the block builder. This should not be called until all other threads are done using the data generator,
// as the data generator is not thread-safe.
func (b *blockBuilder) Start() {
	_ = "STUB: not implemented"

	// Builds blocks and sends them to the blocks channel.
	return
}

func (b *blockBuilder) mainLoop() { _ = "STUB: not implemented"; return }

func (b *blockBuilder) buildBlock() *block { _ = "STUB: not implemented"; return nil }

//nolint:gosec
//nolint:gosec

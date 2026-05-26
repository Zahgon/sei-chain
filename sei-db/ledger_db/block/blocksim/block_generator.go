package blocksim

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-db/common/rand"
	"github.com/sei-protocol/sei-chain/sei-db/ledger_db/block"
)

const (
	blockHashType = 'b'
	txHashType    = 't'
)

// Asynchronously generates random blocks and feeds them into a channel.
type BlockGenerator struct {
	ctx    context.Context
	config *BlocksimConfig
	rand   *rand.CannedRandom

	// The next block height to be assigned.
	nextHeight uint64

	// Generated blocks are sent to this channel.
	blocksChan chan *block.BinaryBlock
}

// Creates a new BlockGenerator and immediately starts its background goroutine.
// The generator stops when the context is cancelled.
func NewBlockGenerator(
	ctx context.Context,
	config *BlocksimConfig,
	rng *rand.CannedRandom,
	startHeight uint64,
) *BlockGenerator {
	_ = "STUB: not implemented"
	return nil
}

// NextBlock blocks until the next generated block is available and returns it.
// Returns nil if the context has been cancelled and no more blocks will be produced.
func (g *BlockGenerator) NextBlock() *block.BinaryBlock { _ = "STUB: not implemented"; return nil }

func (g *BlockGenerator) mainLoop() { _ = "STUB: not implemented"; return }

func (g *BlockGenerator) buildBlock() *block.BinaryBlock { _ = "STUB: not implemented"; return nil }

//nolint:gosec

//nolint:gosec
//nolint:gosec

//nolint:gosec
//nolint:gosec

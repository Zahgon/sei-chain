package memblockdb

import (
	"context"
	"sync"

	"github.com/sei-protocol/sei-chain/sei-db/ledger_db/block"
)

// Shared backing store, keyed by path in test builders to simulate restarts.
type memBlockDBData struct {
	mu             sync.RWMutex
	blocksByHash   map[string]*block.BinaryBlock
	blocksByHeight map[uint64]*block.BinaryBlock
	txByHash       map[string]*block.BinaryTransaction
	lowestHeight   uint64
	highestHeight  uint64
	hasBlocks      bool
}

// An in-memory implementation of the BlockDB interface. Useful as a test fixture to sanity check
// test flows.
type memBlockDB struct {
	data *memBlockDBData
}

// NewMemBlockDB creates an in-memory BlockDB suitable for testing and benchmarks.
func NewMemBlockDB() block.BlockDB { _ = "STUB: not implemented"; return *new(block.BlockDB) }

func (m *memBlockDB) WriteBlock(_ context.Context, blk *block.BinaryBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *memBlockDB) Flush(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (m *memBlockDB) GetBlockByHash(_ context.Context, hash []byte) (*block.BinaryBlock, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (m *memBlockDB) GetBlockByHeight(_ context.Context, height uint64) (*block.BinaryBlock, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (m *memBlockDB) GetTransactionByHash(_ context.Context, hash []byte) (*block.BinaryTransaction, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (m *memBlockDB) Prune(_ context.Context, lowestHeightToKeep uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *memBlockDB) GetLowestBlockHeight(_ context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *memBlockDB) GetHighestBlockHeight(_ context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *memBlockDB) Close(_ context.Context) error { _ = "STUB: not implemented"; return nil }

package producer

import (
	"context"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/consensus"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/mempool"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

// Config is the config of the block scope.
type Config struct {
	MaxGasPerBlock   uint64
	MaxTxsPerBlock   uint64
	MaxTxsPerSecond  utils.Option[uint64]
	MempoolSize      uint64
	BlockInterval    time.Duration
	AllowEmptyBlocks bool
}

// minTxGas is the minimum gas cost of an evm tx.
const minTxGas = 21000

func (c *Config) maxTxsPerBlock() uint64 { _ = "STUB: not implemented"; return 0 }

// MaxGasPerBlockI64 returns MaxGasPerBlock clamped to the int64 range.
// Config validation only enforces > 0 (sei-tendermint/config/autobahn.go),
// so a misconfigured chain with a value above math.MaxInt64 can't silently
// overflow when consumed by APIs that take int64 (the mempool's ReapLimits,
// the RPC layer's ConsensusParamUpdates.Block.MaxGas). Centralizing the
// clamp here means callers pick this up by name instead of repeating
// utils.Clamp[int64] at every site, and any future change to the clamp
// rule (or the underlying field type) lives in one place.
func (c *Config) MaxGasPerBlockI64() int64 { _ = "STUB: not implemented"; return 0 }

// State is the block producer state.
type State struct {
	cfg       *Config
	txMempool *mempool.TxMempool
	// consensus state to which published blocks will be reported.
	consensus *consensus.State
}

// NewState constructs a new block producer state.
// Returns an error if the current node is NOT a producer.
func NewState(cfg *Config, txMempool *mempool.TxMempool, consensus *consensus.State) *State {
	_ = "STUB: not implemented"
	return nil
}

// makePayload constructs payload for the next produced block.
// It waits for any transactions OR until `cfg.BlockInterval` passes.
func (s *State) makePayload(ctx context.Context) (*types.Payload, error) {
	_ = "STUB: not implemented"
	// Wait for transactions. We give up and produce an empty block if mempool is empty for
	// cfg.BlockInterval.
	return nil, nil
}

// If the context has been cancelled though, we just fail.

// TODO: ReapMaxTxsBytesMaxGas does not handle corner cases correctly rn, which actually
// can produce negative total gas. Fixing it right away might be backward incompatible afaict,
// so we leave it as is for now.
// nolint:gosec

// This should never happen: we construct the payload from correctly sized data.

// nextPayload constructs the payload for the next block.
// Wrapper of makePayload which ensures that the block is not empty (if required).
func (s *State) nextPayload(ctx context.Context) (*types.Payload, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run runs the background tasks of the producer state.
func (s *State) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Construct blocks from mempool.

// nolint:gosec

package evmrpc

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	gethtracers "github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/sei-protocol/seilog"

	"github.com/sei-protocol/sei-chain/x/evm/keeper"
)

var bakerLogger = seilog.NewLogger("evmrpc", "trace-baker")

// blockTracer is the subset of *gethtracers.API the baker uses.
type blockTracer interface {
	TraceBlockByNumber(ctx context.Context, number rpc.BlockNumber, config *gethtracers.TraceConfig) ([]*gethtracers.TxTraceResult, error)
}

// TraceBaker re-runs committed blocks through the tracer in background workers
// and writes the JSON to a TraceDB. Enqueue is non-blocking; misses fall
// through to live re-execution.
type TraceBaker struct {
	tracersAPI    blockTracer
	cache         *keeper.TraceDB
	tracers       []string
	bakeTimeout   time.Duration
	tipFn         func() int64
	windowBlocks  int64
	pruneInterval time.Duration

	queue    chan int64
	progress chan int64
	workers  int

	closeOnce sync.Once
	done      chan struct{}
	wg        sync.WaitGroup

	dropped, baked, failed uint64 // atomic
}

type TraceBakerConfig struct {
	Workers       int           // re-execution goroutines (default 1)
	QueueSize     int           // bounds in-flight heights (default 4096); drops on full
	Tracers       []string      // tracers to bake per block (default ["callTracer"])
	BakeTimeout   time.Duration // per-(block,tracer) timeout (default 60s)
	TipFn         func() int64  // chain tip; enables catch-up + prune when set
	WindowBlocks  int64         // catch-up cap and rolling prune window (0 disables prune)
	PruneInterval time.Duration // prune tick (default 1m)
}

// StartTraceBakerForDebugAPI wires a baker against api and starts it.
// Returns nil when the keeper has no TraceDB.
func StartTraceBakerForDebugAPI(api *DebugAPI, cfg TraceBakerConfig) *TraceBaker {
	_ = "STUB: not implemented"
	return nil
}

func NewTraceBaker(api *gethtracers.API, cache *keeper.TraceDB, cfg TraceBakerConfig) *TraceBaker {
	_ = "STUB: not implemented"
	return nil
}

func (b *TraceBaker) Start() { _ = "STUB: not implemented"; return }

// Stop signals goroutines to exit and waits for them to drain. Idempotent.
// Doesn't close b.queue so concurrent Enqueue calls can't panic.
func (b *TraceBaker) Stop() { _ = "STUB: not implemented"; return }

// Enqueue is non-blocking; drops on a full queue. Dropped blocks fall through
// to live re-execution at debug_trace time.
func (b *TraceBaker) Enqueue(height int64) { _ = "STUB: not implemented"; return }

func (b *TraceBaker) DroppedCount() uint64 { _ = "STUB: not implemented"; return 0 }
func (b *TraceBaker) BakedCount() uint64   { _ = "STUB: not implemented"; return 0 }
func (b *TraceBaker) FailedCount() uint64  { _ = "STUB: not implemented"; return 0 }

func (b *TraceBaker) workerLoop() { _ = "STUB: not implemented"; return }

func (b *TraceBaker) bakeBlock(height int64) { _ = "STUB: not implemented"; return }

func (b *TraceBaker) bakeBlockOneTracer(height int64, tracer string) bool {
	_ = "STUB: not implemented"
	return false
}

// Skip empty blocks: json.Marshal(nil) is "null", live path returns [].

func (b *TraceBaker) progressLoop(last int64) { _ = "STUB: not implemented"; return }

func advanceContiguous(last int64, doneHeights map[int64]struct{}) int64 {
	_ = "STUB: not implemented"
	return 0
}

func (b *TraceBaker) progressGapSkipTo(last int64, doneHeights map[int64]struct{}) int64 {
	_ = "STUB: not implemented"
	return 0
}

// catchUpLoop bakes blocks committed since the last successful run, bounded
// by WindowBlocks so a long-stopped node doesn't bake from genesis.
func (b *TraceBaker) catchUpLoop(last int64) { _ = "STUB: not implemented"; return }

func (b *TraceBaker) readStartingLastBaked() (int64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (b *TraceBaker) startingLastBaked(last int64) int64 { _ = "STUB: not implemented"; return 0 }

func (b *TraceBaker) windowFloor(tip int64) int64 { _ = "STUB: not implemented"; return 0 }

// pruneLoop deletes rows older than the configured window every PruneInterval.
func (b *TraceBaker) pruneLoop() { _ = "STUB: not implemented"; return }

func encodeTraceResult(v interface{}) (json.RawMessage, error) {
	_ = "STUB: not implemented"
	return *new(json.RawMessage), nil
}

package keeper

import (
	"encoding/json"
	"sync"

	"github.com/cockroachdb/pebble/v2"
	"github.com/ethereum/go-ethereum/common"
)

// TraceDB stores pre-computed debug_trace results in a pebble db at
// <home>/data/trace_db. Two keyspaces, both height-leading so one range
// delete per prefix prunes a window:
//
//	ts/<height,8>/<tracerLen,1><tracer>/<txHash,32>   per-tx
//	tb/<height,8>/<tracerLen,1><tracer>               per-block (pre-encoded array)
type TraceDB struct {
	db *pebble.DB

	enqMu    sync.Mutex
	enqueuer TraceEnqueuer
}

const (
	traceDBPrefix       = "ts/"
	traceDBBlockPrefix  = "tb/"
	traceDBLastBakedKey = "meta/last_baked_height"
)

func NewTraceDB(homeDir string) (*TraceDB, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *TraceDB) Close() error { _ = "STUB: not implemented"; return nil }

// Drain the baker before closing pebble so workers don't write to a closed db.

func traceDBKey(height int64, tracer string, txHash common.Hash) []byte {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec

func traceDBBlockKey(height int64, tracer string) []byte { _ = "STUB: not implemented"; return nil }

//nolint:gosec

func (c *TraceDB) Put(height int64, tracer string, txHash common.Hash, value json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *TraceDB) Get(height int64, tracer string, txHash common.Hash) (json.RawMessage, bool, error) {
	_ = "STUB: not implemented"
	return *new(json.RawMessage), false, nil
}

func (c *TraceDB) PutBlock(height int64, tracer string, value json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *TraceDB) GetBlock(height int64, tracer string) (json.RawMessage, bool, error) {
	_ = "STUB: not implemented"
	return *new(json.RawMessage), false, nil
}

// SetLastBakedHeight records the highest fully-processed block. Atomic max:
// out-of-order workers can't roll it back.
func (c *TraceDB) SetLastBakedHeight(h int64) error { _ = "STUB: not implemented"; return nil }

//nolint:gosec

func (c *TraceDB) LastBakedHeight() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *TraceDB) lastBakedHeightUnlocked() (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//nolint:gosec

// Prune deletes per-tx and per-block rows with height < belowHeight.
func (c *TraceDB) Prune(belowHeight int64) error { _ = "STUB: not implemented"; return nil }

//nolint:gosec

// TraceEnqueuer is implemented by the trace baker.
type TraceEnqueuer interface {
	Enqueue(height int64)
	Stop()
}

func (c *TraceDB) SetTraceEnqueuer(e TraceEnqueuer) { _ = "STUB: not implemented"; return }

func (c *TraceDB) Enqueue(height int64) { _ = "STUB: not implemented"; return }

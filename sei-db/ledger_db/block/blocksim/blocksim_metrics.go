package blocksim

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-db/common/metrics"
	"github.com/sei-protocol/sei-chain/sei-db/ledger_db/block"
	"go.opentelemetry.io/otel/metric"
)

const blocksimMeterName = "blocksim"

// BlocksimMetrics holds OpenTelemetry metrics for the blocksim benchmark.
type BlocksimMetrics struct {
	ctx context.Context

	blocksWrittenTotal       metric.Int64Counter
	transactionsWrittenTotal metric.Int64Counter
	bytesWrittenTotal        metric.Int64Counter
	pruneCallsTotal          metric.Int64Counter
	flushCallsTotal          metric.Int64Counter

	lowestBlockHeight  metric.Int64Gauge
	highestBlockHeight metric.Int64Gauge
	blockSizeBytes     metric.Int64Gauge

	mainThreadPhase *metrics.PhaseTimer
}

// NewBlocksimMetrics creates metrics for the blocksim benchmark using the
// global OTel MeterProvider. The caller must configure the MeterProvider
// before calling this.
func NewBlocksimMetrics(ctx context.Context, config *BlocksimConfig) *BlocksimMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (m *BlocksimMetrics) recordBlockSize(config *BlocksimConfig) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec

// StartBlockDBPolling launches a background goroutine that periodically queries
// the database for the current lowest and highest block heights, recording them
// as gauge metrics. The goroutine exits when ctx is cancelled.
func (m *BlocksimMetrics) StartBlockDBPolling(ctx context.Context, db block.BlockDB, intervalSeconds int) {
	_ = "STUB: not implemented"
	return
}

func (m *BlocksimMetrics) pollBlockHeights(ctx context.Context, db block.BlockDB) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec

//nolint:gosec

func (m *BlocksimMetrics) ReportBlockWritten(transactionCount int64, byteCount int64) {
	_ = "STUB: not implemented"
	return
}

func (m *BlocksimMetrics) ReportPrune() { _ = "STUB: not implemented"; return }

func (m *BlocksimMetrics) ReportFlush() { _ = "STUB: not implemented"; return }

func (m *BlocksimMetrics) SetMainThreadPhase(phase string) { _ = "STUB: not implemented"; return }

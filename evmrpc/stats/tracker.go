package stats

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/sei-protocol/seilog"
)

// Global tracker state
var (
	logger = seilog.NewLogger("evmrpc", "stats")

	httpTracker *tracker
	wsTracker   *tracker
)

type apiEvent struct {
	Method    string
	Duration  time.Duration
	Success   bool
	StartTime time.Time
	EndTime   time.Time
}

// periodStats holds aggregated stats for a time period.
type periodStats struct {
	periodStart  time.Time
	totalEvents  int
	totalSuccess int
	methodData   map[string]*methodStats
}

// methodStats holds per-method aggregated stats.
type methodStats struct {
	count        int
	successCount int
	totalLatency time.Duration
	maxLatency   time.Duration
}

// InitRPCTracker initializes the HTTP/RPC tracker.
func InitRPCTracker(ctx context.Context, interval time.Duration) { _ = "STUB: not implemented"; return }

// InitWSTracker initializes the WebSocket tracker.
func InitWSTracker(ctx context.Context, interval time.Duration) { _ = "STUB: not implemented"; return }

type tracker struct {
	logger   *slog.Logger
	ch       chan apiEvent
	interval time.Duration
	ctx      context.Context
	cancel   context.CancelFunc
	wg       sync.WaitGroup

	// Simple current period tracking
	mu            sync.RWMutex
	currentPeriod *periodStats
	connType      string
}

// newTracker creates a new stats tracker.
func newTracker(ctx context.Context,
	connType string, interval time.Duration) *tracker {
	_ = "STUB: not implemented"
	return nil
}

// run processes events continuously and reports periods on interval
func (t *tracker) run() {
	_ = "STUB: not implemented"

	// Report stats every interval.
	return
}

// Report current period before stopping

// Process event immediately

// Report current period and start fresh

// processEvent aggregates an event into the current period
func (t *tracker) processEvent(event apiEvent) { _ = "STUB: not implemented"; return }

// Truncate event end time to period boundary (use completion time for period attribution)

// Check if we need to rotate periods

// apiEvent is in a new period, so report the current period and start fresh

// Initialize current period if needed (using event timestamp)

// Update overall stats

// Get or create method stats

// Update method stats

// reportCurrentPeriod reports the current period and starts a new one
func (t *tracker) reportCurrentPeriod() { _ = "STUB: not implemented"; return }

// If no current period, create an empty one for this interval

// Report the current period

// Start a new period

// reportPeriodLocked logs the stats for a completed period (assumes lock is held)
func (t *tracker) reportPeriodLocked(period *periodStats) {
	_ = "STUB: not implemented"
	// Always log overall stats, even for periods with no events
	return
}

// Log overall stats for periods with no requests

// No method stats to report

// Calculate overall success rate

// Log overall stats for this period

// Log per-method stats for this period

// Calculate average latency in milliseconds with 2 decimal places

// TrackMessage tracks a JSON-RPC method call with timing information
func (t *tracker) TrackMessage(method string, connectionType string, startTime time.Time, success bool) {
	_ = "STUB: not implemented"
	return

	// Gracefully handle nil tracker
}

// Drop on overflow to not block RPC - log at debug level to avoid spam

func (t *tracker) Stop() {
	_ = "STUB: not implemented"
	// Cancel context to stop the goroutine
	return
}

// Wait for goroutine to finish

// RecordAPIInvocation is a simple entry point for recording API calls.
// It uses the appropriate tracker based on connection type.
// InitRPCTracker and InitWSTracker must be called first from server creation.
func RecordAPIInvocation(method string, connectionType string, startTime time.Time, success bool) {
	_ = "STUB: not implemented"
	return
}

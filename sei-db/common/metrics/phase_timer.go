package metrics

import (
	"time"

	"go.opentelemetry.io/otel/metric"
)

// PhaseTimerFactory constructs shared OTel metrics and builds independent
// PhaseTimer instances. Use Build() to create a timer for each thread.
type PhaseTimerFactory struct {
	phaseDurationTotal metric.Float64Counter
	phaseLatency       metric.Float64Histogram
	timerName          string
}

// NewPhaseTimerFactory creates a factory that records to the given meter with the
// specified timer name (e.g., "main_thread" or "transaction"). Metric names are
// {timerName}_phase_duration_seconds_total and {timerName}_phase_latency_seconds
// to match existing Grafana dashboards.
func NewPhaseTimerFactory(meter metric.Meter, timerName string) *PhaseTimerFactory {
	_ = "STUB: not implemented"
	return nil
}

// NewPhaseTimer creates a factory and builds a single PhaseTimer. Convenient when
// only one timer is needed (e.g., for a single-threaded main loop).
func NewPhaseTimer(meter metric.Meter, timerName string) *PhaseTimer {
	_ = "STUB: not implemented"
	return nil
}

// Build returns a new PhaseTimer that records to this factory's metrics.
// Each timer has independent phase state; safe for use by different threads.
func (f *PhaseTimerFactory) Build() *PhaseTimer { _ = "STUB: not implemented"; return nil }

// PhaseTimer records time spent in phases (e.g., "executing", "finalizing").
// Call SetPhase when transitioning to a new phase; latency is calculated from the
// previous transition. Not safe for concurrent use on a single instance.
type PhaseTimer struct {
	phaseDurationTotal  metric.Float64Counter
	phaseLatency        metric.Float64Histogram
	lastPhase           string
	lastPhaseChangeTime time.Time
}

// SetPhase records a transition to a new phase.
func (p *PhaseTimer) SetPhase(phase string) { _ = "STUB: not implemented"; return }

// Reset ends the current phase (capturing its metrics) and clears the phase state.
func (p *PhaseTimer) Reset() { _ = "STUB: not implemented"; return }

package telemetry

import (
	"time"

	metrics "github.com/armon/go-metrics"
)

// Common metric key constants
const (
	MetricKeyBeginBlocker = "begin_blocker"
	MetricKeyMidBlocker   = "mid_blocker"
	MetricKeyEndBlocker   = "end_blocker"
	MetricLabelNameModule = "module"
	MessageCount          = "message"
	TxCount               = "transaction"
)

// NewLabel creates a new instance of Label with name and value
func NewLabel(name, value string) metrics.Label {
	_ = "STUB: not implemented"
	return *new(metrics.Label)
}

// ModuleMeasureSince provides a short hand method for emitting a time measure
// metric for a module with a given set of keys. If any global labels are defined,
// they will be added to the module label.
func ModuleMeasureSince(module string, start time.Time, keys ...string) {
	_ = "STUB: not implemented"
	return
}

// ModuleSetGauge provides a short hand method for emitting a gauge metric for a
// module with a given set of keys. If any global labels are defined, they will
// be added to the module label.
func ModuleSetGauge(module string, val float32, keys ...string) { _ = "STUB: not implemented"; return }

// IncrCounter provides a wrapper functionality for emitting a counter metric with
// global labels (if any).
func IncrCounter(val float32, keys ...string) { _ = "STUB: not implemented"; return }

// IncrCounterWithLabels provides a wrapper functionality for emitting a counter
// metric with global labels (if any) along with the provided labels.
func IncrCounterWithLabels(keys []string, val float32, labels []metrics.Label) {
	_ = "STUB: not implemented"
	return
}

// SetGauge provides a wrapper functionality for emitting a gauge metric with
// global labels (if any).
func SetGauge(val float32, keys ...string) { _ = "STUB: not implemented"; return }

// SetGaugeWithLabels provides a wrapper functionality for emitting a gauge
// metric with global labels (if any) along with the provided labels.
func SetGaugeWithLabels(keys []string, val float32, labels []metrics.Label) {
	_ = "STUB: not implemented"
	return
}

// MeasureSince provides a wrapper functionality for emitting a a time measure
// metric with global labels (if any).
func MeasureSince(start time.Time, keys ...string) { _ = "STUB: not implemented"; return }

// MeasureSinceWithLabels provides a wrapper functionality for emitting a a time measure
// metric with custom labels (if any)
func MeasureSinceWithLabels(keys []string, start time.Time, labels []metrics.Label) {
	_ = "STUB: not implemented"
	return
}

// Measure Validator slashing events
// validator_slashed
func IncrValidatorSlashedCounter(validator string, slashingType string) {
	_ = "STUB: not implemented"
	return
}

// Measures throughput
// Metric Name:
//
//	sei_throughput_<metric_name>
func MeasureThroughputSinceWithLabels(metricName string, labels []metrics.Label, start time.Time) {
	_ = "STUB: not implemented"
	return
}

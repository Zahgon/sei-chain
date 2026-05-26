package metrics

import (
	"context"

	"go.opentelemetry.io/otel/metric"
)

// MonitoredDir describes a directory whose on-disk size should be tracked.
type MonitoredDir struct {
	// Name is the metric-name component, e.g. "data_dir" produces
	// "{prefix}_data_dir_size_bytes".
	Name string
	// Path is the filesystem path to monitor.
	Path string
	// TrackAvailableSpace also emits "{prefix}_{Name}_available_bytes".
	TrackAvailableSpace bool
}

// StartSystemMetrics creates OTel instruments for system-level metrics and
// spawns background goroutines that poll them periodically. All goroutines
// exit when ctx is cancelled. If intervalSeconds <= 0 the call is a no-op.
//
// Metrics created (where {p} = prefix, {d} = dir.Name):
//
//	{p}_{d}_size_bytes            for each dir
//	{p}_{d}_available_bytes       for dirs with TrackAvailableSpace
//	{p}_uptime_seconds
//	{p}_process_read_bytes_total  (Linux only)
//	{p}_process_write_bytes_total (Linux only)
//	{p}_process_read_count_total  (Linux only)
//	{p}_process_write_count_total (Linux only)
func StartSystemMetrics(ctx context.Context, prefix string, intervalSeconds int, dirs []MonitoredDir) {
	_ = "STUB: not implemented"
	return
}

// startPeriodicSampling runs sampleFn immediately and then every intervalSeconds
// in a background goroutine. The goroutine exits when ctx is cancelled.
func startPeriodicSampling(ctx context.Context, intervalSeconds int, sampleFn func()) {
	_ = "STUB: not implemented"
	return
}

// startUptimeSampling records elapsed seconds since now into gauge once per second.
func startUptimeSampling(ctx context.Context, gauge metric.Float64Gauge) {
	_ = "STUB: not implemented"
	return
}

// startProcessIOSampling tracks process-level I/O counters via gopsutil.
// Skipped on darwin where gopsutil does not implement IOCounters.
func startProcessIOSampling(ctx context.Context, meter metric.Meter, prefix string, intervalSeconds int) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec

// measureDirSize walks dir and returns the sum of all regular file sizes.
func measureDirSize(dir string) int64 { _ = "STUB: not implemented"; return 0 }

// measureAvailableBytes returns the available bytes on the filesystem containing dir.
func measureAvailableBytes(dir string) int64 { _ = "STUB: not implemented"; return 0 }

//nolint:gosec

func uint64ToInt64Clamped(v uint64) int64 { _ = "STUB: not implemented"; return 0 }

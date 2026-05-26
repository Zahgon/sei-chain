package metrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

// SetupOtelPrometheus configures the global OTel MeterProvider backed by a
// Prometheus registry. Returns the registry (for HTTP serving) and a shutdown
// function that flushes the provider.
func SetupOtelPrometheus() (*prometheus.Registry, func(context.Context) error, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// StartMetricsServer serves /metrics from the given gatherer on addr.
// If addr is empty the call is a no-op. The server shuts down when ctx is
// cancelled.
func StartMetricsServer(ctx context.Context, gatherer prometheus.Gatherer, addr string) {
	_ = "STUB: not implemented"
	return
}

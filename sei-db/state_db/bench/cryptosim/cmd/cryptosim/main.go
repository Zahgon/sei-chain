package main

import (
	"context"
	"fmt"
	"os"

	"github.com/prometheus/client_golang/prometheus"
)

// setupOtelPrometheus configures the global OTel MeterProvider to export to Prometheus.
// Returns the registry (for HTTP serving) and a shutdown function.
func setupOtelPrometheus() (*prometheus.Registry, func(context.Context) error, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// No namespace: instrument names (e.g. cryptosim_blocks_finalized_total) are used as-is for Grafana compatibility.

// startMetricsServer serves /metrics from the given gatherer. Shuts down when ctx is cancelled.
func startMetricsServer(ctx context.Context, gatherer prometheus.Gatherer, addr string) {
	_ = "STUB: not implemented"
	return
}

// Run the cryptosim benchmark.
func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error { _ = "STUB: not implemented"; return nil }

// Configure OTel to export to Prometheus before creating cryptosim (metrics use global provider).

// Start metrics HTTP server after cryptosim setup (metrics are populated).

// Toggle suspend/resume on Enter when enabled

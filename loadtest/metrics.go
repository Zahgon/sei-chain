package main

import (
	"net/http"

	"github.com/sei-protocol/sei-chain/sei-cosmos/telemetry"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

var (
	ltMeter = otel.Meter("loadtest")

	loadtestMetrics = struct {
		produceCount metric.Int64Counter
		consumeCount metric.Int64Counter
		tps          metric.Float64Gauge
	}{
		produceCount: must(ltMeter.Int64Counter(
			"produce",
			metric.WithDescription("Number of transactions produced by message type"),
			metric.WithUnit("{count}"),
		)),
		consumeCount: must(ltMeter.Int64Counter(
			"consume",
			metric.WithDescription("Number of transactions consumed by message type"),
			metric.WithUnit("{count}"),
		)),
		tps: must(ltMeter.Float64Gauge(
			"tps",
			metric.WithDescription("Transactions per second by message type"),
			metric.WithUnit("{tps}"),
		)),
	}
)

func must[V any](v V, err error) V { _ = "STUB: not implemented"; return *new(V) }

const (
	defaultListenAddress = "0.0.0.0"
	defaultMetricsPort   = 9696
)

type MetricsServer struct {
	metrics *telemetry.Metrics
	server  *http.Server
}

func (s *MetricsServer) metricsHandler(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *MetricsServer) StartMetricsClient(config Config) { _ = "STUB: not implemented"; return }

func (s *MetricsServer) healthzHandler(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}

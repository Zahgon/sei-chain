package data

import (
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/component-base/metrics/prometheusextension"
)

var _ prometheus.Collector = (*State)(nil)

type latencyMetric struct {
	*prometheusextension.WeightedHistogramVec
}

type resourceLatencyMetric struct {
	Receive prometheusextension.WeightedObserver
	Execute prometheusextension.WeightedObserver
	Prune   prometheusextension.WeightedObserver
}

func newLatencyMetric() latencyMetric { _ = "STUB: not implemented"; return *new(latencyMetric) }

func (m latencyMetric) resource(resource string) resourceLatencyMetric {
	_ = "STUB: not implemented"
	return *new(resourceLatencyMetric)
}

type dataMetrics struct {
	Base   latencyMetric
	Blocks resourceLatencyMetric
	Txs    resourceLatencyMetric
}

func newDataMetrics() *dataMetrics { _ = "STUB: not implemented"; return nil }

// Describe from prometheus.Collector.
func (s *State) Describe(chan<- *prometheus.Desc) {
	_ = "STUB: not implemented"

	// Collect from prometheus.Collector.
	return
}

func (s *State) Collect(m chan<- prometheus.Metric) { _ = "STUB: not implemented"; return }

package keeper

import (
	"github.com/prometheus/client_golang/prometheus"
	wasmvmtypes "github.com/sei-protocol/sei-chain/sei-wasmvm/types"
)

const (
	labelPinned = "pinned"
	labelMemory = "memory"
	labelFs     = "fs"
)

// metricSource source of wasmvm metrics
type metricSource interface {
	GetMetrics() (*wasmvmtypes.Metrics, error)
}

var _ prometheus.Collector = (*WasmVMMetricsCollector)(nil)

// WasmVMMetricsCollector custom metrics collector to be used with Prometheus
type WasmVMMetricsCollector struct {
	source             metricSource
	CacheHitsDescr     *prometheus.Desc
	CacheMissesDescr   *prometheus.Desc
	CacheElementsDescr *prometheus.Desc
	CacheSizeDescr     *prometheus.Desc
}

// NewWasmVMMetricsCollector constructor
func NewWasmVMMetricsCollector(s metricSource) *WasmVMMetricsCollector {
	_ = "STUB: not implemented"
	return nil
}

// Register registers all metrics
func (p *WasmVMMetricsCollector) Register(r prometheus.Registerer) {
	_ = "STUB: not implemented"

	// Describe sends the super-set of all possible descriptors of metrics
	return
}

func (p *WasmVMMetricsCollector) Describe(descs chan<- *prometheus.Desc) {
	_ = "STUB: not implemented"
	return
}

// Collect is called by the Prometheus registry when collecting metrics.
func (p *WasmVMMetricsCollector) Collect(c chan<- prometheus.Metric) {
	_ = "STUB: not implemented"
	return
}

// Node about fs metrics:
// The number of elements and the size of elements in the file system cache cannot easily be obtained.
// We had to either scan the whole directory of potentially thousands of files or track the values when files are added or removed.
// Such a tracking would need to be on disk such that the values are not cleared when the node is restarted.

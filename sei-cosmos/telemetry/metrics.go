package telemetry

import (
	metrics "github.com/armon/go-metrics"
	metricsprom "github.com/armon/go-metrics/prometheus"
)

// globalLabels defines the set of global labels that will be applied to all
// metrics emitted using the telemetry package function wrappers.
var globalLabels = []metrics.Label{}

// Metrics supported format types.
const (
	FormatDefault    = ""
	FormatPrometheus = "prometheus"
	FormatText       = "text"
)

// Config defines the configuration options for application telemetry.
type Config struct {
	// Prefixed with keys to separate services
	ServiceName string `mapstructure:"service-name"`

	// Enabled enables the application telemetry functionality. When enabled,
	// an in-memory sink is also enabled by default. Operators may also enabled
	// other sinks such as Prometheus.
	Enabled bool `mapstructure:"enabled"`

	// Enable prefixing gauge values with hostname
	EnableHostname bool `mapstructure:"enable-hostname"`

	// Enable adding hostname to labels
	EnableHostnameLabel bool `mapstructure:"enable-hostname-label"`

	// Enable adding service to labels
	EnableServiceLabel bool `mapstructure:"enable-service-label"`

	// PrometheusRetentionTime, when positive, enables a Prometheus metrics sink.
	// It defines the retention duration in seconds.
	PrometheusRetentionTime int64 `mapstructure:"prometheus-retention-time"`

	// GlobalLabels defines a global set of name/value label tuples applied to all
	// metrics emitted using the wrapper functions defined in telemetry package.
	//
	// Example:
	// [["chain_id", "cosmoshub-1"]]
	GlobalLabels [][]string `mapstructure:"global-labels"`
}

// Metrics defines a wrapper around application telemetry functionality. It allows
// metrics to be gathered at any point in time. When creating a Metrics object,
// internally, a global metrics is registered with a set of sinks as configured
// by the operator. In addition to the sinks, when a process gets a SIGUSR1, a
// dump of formatted recent metrics will be sent to STDERR.
type Metrics struct {
	memSink           *metrics.InmemSink
	prometheusEnabled bool
}

// GatherResponse is the response type of registered metrics
type GatherResponse struct {
	Metrics     []byte
	ContentType string
}

// New creates a new instance of Metrics
func New(cfg Config) (*Metrics, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Metrics) setupPrometheus(cfg Config) (*metricsprom.PrometheusSink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Default definitions, this allows Prometheus to persist metrics between scrapes instead
// not reporting them if they are not updated. Please use this only if needed as it
// will mean the metrics are stored in memory.

// Gather collects all registered metrics and returns a GatherResponse where the
// metrics are encoded depending on the type. Metrics are either encoded via
// Prometheus or JSON if in-memory.
func (m *Metrics) Gather(format string) (GatherResponse, error) {
	_ = "STUB: not implemented"
	return *new(GatherResponse), nil
}

func (m *Metrics) gatherPrometheus() (GatherResponse, error) {
	_ = "STUB: not implemented"
	return *new(GatherResponse), nil
}

func (m *Metrics) gatherGeneric() (GatherResponse, error) {
	_ = "STUB: not implemented"
	return *new(GatherResponse), nil
}

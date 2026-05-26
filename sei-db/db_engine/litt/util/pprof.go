package util

import (
	"log/slog"

	_ "net/http/pprof" //nolint:gosec // pprof endpoint is intentional for profiling
)

type PprofProfiler struct {
	logger   *slog.Logger
	httpPort string
}

func NewPprofProfiler(httpPort string, logger *slog.Logger) *PprofProfiler {
	_ = "STUB: not implemented"
	return nil
}

// Start the pprof server
func (p *PprofProfiler) Start() { _ = "STUB: not implemented"; return }

package evmrpc

import (
	"net/http"
)

type wsConnectionHandler struct {
	underlying http.Handler
}

func (h *wsConnectionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// TODO(PLT-326): remove legacy dual-emit once dashboards are migrated to evmrpc_* OTEL metrics. Use metrics.wsConnectionCount instead.

func NewWSConnectionHandler(handler http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

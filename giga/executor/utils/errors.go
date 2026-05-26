package utils

const (
	// GigaAbortCodespace and GigaAbortCode identify a sentinel ResponseDeliverTx
	// that signals the caller to fall back to the v2 execution path.
	GigaAbortCodespace = "giga"
	GigaAbortCode      = uint32(1)
	GigaAbortInfo      = "giga_fallback_to_v2"
)

// ShouldExecutionAbort checks if the given error is an AbortError that should
// cause Giga execution to abort and fall back to standard execution.
func ShouldExecutionAbort(err error) bool { _ = "STUB: not implemented"; return false }

package util

import (
	"log/slog"
)

// GetLogger returns a logger for use in tests.
// The logger always includes source information and logs at debug level.
func GetLogger() *slog.Logger { _ = "STUB: not implemented"; return nil }

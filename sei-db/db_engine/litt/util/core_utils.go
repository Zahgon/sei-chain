package util

import (
	"io"
	"log/slog"
)

// CloseLogOnError attempts to close the given io.Closer and logs an error if it fails.
// Meant to be called in a defer statement: defer CloseLogOnError(c, "nameOfResourceToClose", log).
func CloseLogOnError(c io.Closer, name string, log *slog.Logger) { _ = "STUB: not implemented"; return }

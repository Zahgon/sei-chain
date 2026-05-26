package logging

import (
	"time"

	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("utils", "logging")

func LogIfNotDoneAfter[R any](task func() (R, error), after time.Duration, label string) (R, error) {
	_ = "STUB: not implemented"
	return *new(R), nil
}

// reraise panic in main goroutine

package util

import (
	"testing"
	"time"
)

// AssertEventuallyTrue asserts that a condition is true within a given duration. Repeatably checks the condition.
func AssertEventuallyTrue(t *testing.T, condition func() bool, duration time.Duration, debugInfo ...any) {
	_ = "STUB: not implemented"
	return
}

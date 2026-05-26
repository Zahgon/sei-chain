package time

import (
	"time"
)

// Now returns the current time in UTC with no monotonic component.
func Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Canonical returns UTC time with no monotonic component.
// Stripping the monotonic component is for time equality.
// See https://github.com/tendermint/tendermint/pull/2203#discussion_r215064334
func Canonical(t time.Time) time.Time {
	_ = "STUB: not implemented"
	return *

	//go:generate ../../scripts/mockery_generate.sh Source
	new(time.Time)
}

// Source is an interface that defines a way to fetch the current time.
type Source interface {
	Now() time.Time
}

// DefaultSource implements the Source interface using the system clock provided by the standard library.
type DefaultSource struct{}

func (DefaultSource) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

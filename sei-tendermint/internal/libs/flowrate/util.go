//
// Written by Maxim Khitrov (November 2012)
//

package flowrate

import (
	"time"
)

// clockRate is the resolution and precision of clock().
const clockRate = 20 * time.Millisecond

// clock returns a low resolution timestamp relative to the process start time.
func clock(startAt time.Time) time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// clockRound returns d rounded to the nearest clockRate increment.
func clockRound(d time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// round returns x rounded to the nearest int64 (non-negative values only).
func round(x float64) int64 { _ = "STUB: not implemented"; return 0 }

// Percent represents a percentage in increments of 1/1000th of a percent.
type Percent uint32

// percentOf calculates what percent of the total is x.
func percentOf(x, total float64) Percent { _ = "STUB: not implemented"; return *new(Percent) }

//nolint:gosec // bounds checked above against MaxUint32

func (p Percent) Float() float64 { _ = "STUB: not implemented"; return 0 }

func (p Percent) String() string { _ = "STUB: not implemented"; return "" }

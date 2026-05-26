package util

import "time"

// ToMilliseconds converts the given duration to milliseconds. Unlike duration.Milliseconds(), this function returns
// a float64 with nanosecond precision (at least, as much precision as floating point numbers allow).
func ToMilliseconds(duration time.Duration) float64 { _ = "STUB: not implemented"; return 0 }

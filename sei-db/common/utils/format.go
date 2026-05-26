package utils

import (
	"time"
)

// Int64Commas formats n with commas as thousands separators (e.g., 1000000 -> "1,000,000").
func Int64Commas(n int64) string { _ = "STUB: not implemented"; return "" }

// FormatNumberFloat64 formats f with commas in the integer part and the given number of decimal places.
// Special values (NaN, Inf) are formatted as strconv.FormatFloat would.
func FormatNumberFloat64(f float64, decimals int) string { _ = "STUB: not implemented"; return "" }

// FormatBytes formats a byte count using the most appropriate binary unit
// (e.g. 1536 -> "1.50 KiB", 2097152 -> "2.00 MiB").
func FormatBytes(b int64) string { _ = "STUB: not implemented"; return "" }

// FormatDuration formats d using the most appropriate unit (days, hours, minutes, seconds, ms, us, ns).
func FormatDuration(d time.Duration, decimals int) string { _ = "STUB: not implemented"; return "" }

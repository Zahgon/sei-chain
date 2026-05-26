package cryptosim

import (
	"time"
)

// BytesToHex returns a lowercase hex string with 0x prefix, suitable for printing binary keys or addresses.
func BytesToHex(b []byte) string { _ = "STUB: not implemented"; return "" }

// Get the key for the account ID counter in the database.
// Uses EVMKeyCode with padded keyBytes; EVMKeyNonce requires 20-byte addresses and
// non-standard lengths are routed to EVMKeyLegacy which FlatKV ignores.
func AccountIDCounterKey() []byte { _ = "STUB: not implemented"; return nil }

// Get the key for the ERC20 contract ID counter in the database.
func Erc20IDCounterKey() []byte { _ = "STUB: not implemented"; return nil }

// Get the key for the block number counter in the database.
func BlockNumberCounterKey() []byte { _ = "STUB: not implemented"; return nil }

// paddedCounterKey pads the string to AddressLen bytes for use with EVM key builders.
func paddedCounterKey(s string) []byte { _ = "STUB: not implemented"; return nil }

// int64Commas formats n with commas as thousands separators (e.g., 1000000 -> "1,000,000").
func int64Commas(n int64) string { _ = "STUB: not implemented"; return "" }

// formatNumberFloat64 formats f with commas in the integer part and the given number of decimal places.
// Special values (NaN, Inf) are formatted as strconv.FormatFloat would.
func formatNumberFloat64(f float64, decimals int) string { _ = "STUB: not implemented"; return "" }

// Handle negative sign - we'll need to format the abs and prepend minus

// formatDuration formats d using the most appropriate unit (days, hours, minutes, seconds, ms, µs, ns).
func formatDuration(d time.Duration, decimals int) string { _ = "STUB: not implemented"; return "" }

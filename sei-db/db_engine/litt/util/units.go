package util

// the name of a unit step
type unitStep struct {
	// name of the unit step
	name string
	// multiply by this number to get the previous unit step. For example, if this unit is "KiB", the step is 1024.
	// Taking a number of kilobytes and multiplying by 1024 gives you the number of bytes.
	multiple uint64
}

// byteUnits is a list of units for bytes, in increasing order of size.
var byteSteps = []unitStep{
	{"bytes", 1},
	{"KiB", 1024},
	{"MiB", 1024},
	{"GiB", 1024},
	{"TiB", 1024},
	{"PiB", 1024},
	{"EiB", 1024},
}

var timeSteps = []unitStep{
	{"ns", 1},
	{"μs", 1000},
	{"ms", 1000},
	{"s", 1000},
	{"minutes", 60},
	{"hours", 60},
	{"days", 24},
	{"years", 365}, // I don't care that this is imprecise, I refuse to mess with leap years.
}

// prettyPrintUnit formats a quantity in a human-readable way using the provided unit steps. The quantity
// is assumed to be in the smallest supported unit (e.g., bytes, nanoseconds, etc.).
func prettyPrintUnit(quantity uint64, steps []unitStep) string {
	_ = "STUB: not implemented"
	return ""
}

// Edge case, print without a decimal point if we have the smallest unit.

// We've found the appropriate unit.

// PrettyPrintBytes formats a byte count into a human-readable string with appropriate units.
func PrettyPrintBytes(bytes uint64) string { _ = "STUB: not implemented"; return "" }

// PrettyPrintTime formats a time duration in nanoseconds into a human-readable string with appropriate units.
func PrettyPrintTime(nanoseconds uint64) string { _ = "STUB: not implemented"; return "" }

// CommaOMatic converts a number into string representation with commas for thousands, millions, etc.
func CommaOMatic(value uint64) string { _ = "STUB: not implemented"; return "" }

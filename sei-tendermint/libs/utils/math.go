package utils

import (
	"golang.org/x/exp/constraints"
)

// Bits returns the number of bits of the given integer type.
func Bits[I constraints.Integer]() uintptr { _ = "STUB: not implemented"; return 0 }

// Max returns the maximal value of the given integer type.
func Max[I constraints.Integer]() I {
	_ = "STUB: not implemented"

	// Returns true iff I is a signed integer type.
	return *new(I)
}

func Signed[I constraints.Integer]() bool {
	_ = "STUB: not implemented"

	// Min returns the minimal value of the given integer type.
	return false
}

func Min[I constraints.Integer]() I { _ = "STUB: not implemented"; return *new(I) }

// SafeCast casts between integer types, checking for overflows.
func SafeCast[To, From constraints.Integer](v From) (x To, ok bool) {
	_ = "STUB: not implemented"

	// This can be further optimized by:
	// * making compiler detect if From -> To conversion is always safe
	// * making compiler detect if the parity check is necessary
	return *new(To), false
}

// Clamp converts an integer to another integer type clamping it to the target types' [min,max] range
// in case of overflow.
func Clamp[To, From constraints.Integer](v From) To { _ = "STUB: not implemented"; return *new(To) }

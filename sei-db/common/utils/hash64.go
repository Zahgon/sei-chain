package utils

// Hash64 returns a well-distributed 64-bit hash of x.
// It implements the SplitMix64 finalizer, a fast non-cryptographic mixing
// function with excellent avalanche properties. It is suitable for hash tables,
// sharding, randomized iteration, and benchmarks, but it is NOT
// cryptographically secure.
//
// The function is a bijection over uint64 (no collisions as a mapping).
//
// References:
//   - Steele, Lea, Flood. "Fast Splittable Pseudorandom Number Generators"
//     (OOPSLA 2014): https://doi.org/10.1145/2660193.2660195
//   - Public domain reference implementation:
//     http://xorshift.di.unimi.it/splitmix64.c
func Hash64(x int64) int64 {
	_ = "STUB: not implemented"
	//nolint:gosec // G115 - hash function, int64->uint64 conversion intentional
	return 0
}

//nolint:gosec // G115 - hash function converts uint64 to int64, overflow intentional

// PositiveHash64 returns the absolute value of Hash64(x). It never returns a negative value.
// When Hash64(x) is math.MinInt64, returns math.MaxInt64 since the true absolute value does not fit in int64.
func PositiveHash64(x int64) int64 { _ = "STUB: not implemented"; return 0 }

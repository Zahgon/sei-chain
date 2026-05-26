package rand

// CannedRandom provides pre-generated randomness for benchmarking.
// It contains a buffer of random bytes that it reuses to avoid generating lots of random numbers
// at runtime.
//
// The goal is to avoid accidentally creating CPU hotspots in the benchmark framework. We want
// to exercise the DB, not the code that feeds it randomness.
//
// This utility is not thread safe.
//
// In case it requires saying, this utility is NOT a suitable source cryptographically secure random numbers.
type CannedRandom struct {
	// Pre-generated buffer of random bytes.
	buffer []byte

	// Controls the next index to read from the buffer.
	index int64
}

// NewCannedRandom creates a new CannedRandom.
func NewCannedRandom(
	// The size of the buffer to create. A slice of this size is instantiated, so avoid setting this too large.
	// Do not set this too small though, as the buffer will panic if the requested number of bytes is greater
	// than the buffer size.
	bufferSize int,
	// The seed to use to generate the random bytes. This utility provides deterministic random numbers
	// given the same seed.
	seed int64,
) *CannedRandom {
	_ = "STUB: not implemented"
	return nil
}

// Adjust the buffer size so that (bufferSize % 8) == 1. This way when the index wraps around, it won't align
// perfectly with the buffer, giving us a longer runway before we repeat the exact same sequence of bytes.
// The expression maps (bufferSize%8) -> add: 0->1, 1->0, 2->7, 3->6, 4->5, 5->4, 6->3, 7->2.

// Clone creates a copy of the CannedRandom. On its own, a single CannedRandom is not thread safe. A cloned copy
// is however thread safe, with respect to the original CannedRandom and other cloned copies. Cloning a CannedRandom
// is cheap, and does not require significant additional memory.
//
// If randomizeOffset is true, the clone's read position is set to a new offset derived from the source's buffer,
// ensuring that the clone does not produce random numbers in lockstep with the original. Successive calls to
// Clone(true) yield clones with different offsets.
func (cr *CannedRandom) Clone(randomizeOffset bool) *CannedRandom {
	_ = "STUB: not implemented"
	return nil
}

// Reset the index of the CannedRandom to the beginning of the buffer.
func (cr *CannedRandom) Reset() {
	_ = "STUB: not implemented"

	// Returns a slice of random bytes.
	//
	// Returned slice is NOT safe to modify. If modification is required, the caller should make a copy of the slice.
	return
}

func (cr *CannedRandom) Bytes(count int) []byte { _ = "STUB: not implemented"; return nil }

// Returns a slice of random bytes from a given seed. Bytes are deterministic given the same seed.
//
// Unlike most CannedRandom methods, SeededBytes is safe for concurrent use: it only reads
// from the immutable buffer and does not advance the internal index.
//
// Returned slice is NOT safe to modify. If modification is required, the caller should make a copy of the slice.
func (cr *CannedRandom) SeededBytes(count int, seed int64) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Generate a random-ish int64.
func (cr *CannedRandom) Int64() int64 { _ = "STUB: not implemented"; return 0 }

//nolint:gosec // G115 - benchmark uses deterministic non-crypto randomness, overflow acceptable

// Add 8 to the index to skip the 8 bytes we just read.

// Int64Range returns a random int64 in [min, max). Min is inclusive, max is exclusive.
// If min == max, returns min.
func (cr *CannedRandom) Int64Range(min int64, max int64) int64 { _ = "STUB: not implemented"; return 0 }

//nolint:gosec // G115 - benchmark uses deterministic non-crypto randomness, overflow acceptable

// Float64 returns a random float64 in the range [0.0, 1.0].
// It uses Int64() internally and converts the result to the proper range.
func (cr *CannedRandom) Float64() float64 {
	_ = "STUB: not implemented"
	//nolint:gosec // G115 - benchmark uses deterministic non-crypto randomness, overflow acceptable
	return 0
}

// Bool returns a random boolean.
func (cr *CannedRandom) Bool() bool { _ = "STUB: not implemented"; return false }

// Address generates a deterministic byte sequence suitable for simulating keys.
// For the same input arguments, a canned random generator with the same seed and buffer size
// will produce the same output.
//
// The first keys.AddressLen bytes have the following shape (eth-style address):
//
//	1 byte addressType
//	8 bytes of random data
//	8 bytes containing the ID
//	the remainder is filled with random data
//
// The ID is not at the beginning so that adjacent IDs will not appear close to each other
// if keys are sorted lexicographically. If size > keys.AddressLen, the remainder is filled with
// additional deterministic random bytes seeded by id.
func (cr *CannedRandom) Address(
	// A one-char byte descriptor. Allows keys for different types to have different values
	// even if they have the same ID.
	addressType uint8,
	// A unique ID for the key.
	id int64,
	// Total size in bytes. Must be at least keys.AddressLen.
	size int,
) []byte {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // G115 - id is from benchmark data, overflow acceptable for address generation

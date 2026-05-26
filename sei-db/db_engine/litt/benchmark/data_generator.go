package benchmark

import (
	"sync"
)

// DataGenerator is responsible for generating key-value pairs to be inserted into the database, for the sake of
// benchmarking.
type DataGenerator struct {
	// Pool of random number generators
	randPool *sync.Pool

	// A pool of randomness. Used to generate values.
	dataPool []byte

	// The seed that determines the key/value pairs generated.
	seed int64
}

// NewDataGenerator builds a data generator instance.
func NewDataGenerator(seed int64, poolSize uint64) *DataGenerator {
	_ = "STUB: not implemented"
	return nil
}

// Key generates a new key. The value is deterministic for the same index and seed.
func (g *DataGenerator) Key(index uint64) []byte { _ = "STUB: not implemented"; return nil }

//nolint:gosec // deterministic test seeding

// Value generates a new value. The value is deterministic for the same index, seed, and value size.
func (g *DataGenerator) Value(index uint64, valueLength uint64) []byte {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // deterministic test seeding

// Special case: we don't have enough data in the pool to satisfy the request.
// For the sake of completeness, just generate the data if this happens.
// This shouldn't be encountered for sane configurations (i.e. with a pool size much larger than value sizes).

//nolint:gosec // valueLength bounded by len(dataPool)
//nolint:gosec // valueLength bounded by len(dataPool)

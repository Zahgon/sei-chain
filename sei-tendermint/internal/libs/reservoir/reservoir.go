package reservoir

import (
	"cmp"
	"math/rand/v2"
	"sync"
	"time"
)

// lastPercentileCacheTTL is the duration for which a cached percentile value is
// considered valid if no new percentile p is asked for.
const lastPercentileCacheTTL = 5 * time.Second

// Sampler maintains a thread-safe reservoir of size k for ordered items of
// type T, allowing random sampling from a stream of unknown length and
// percentile queries on the current samples.
//
// It uses Vitter's Algorithm R for reservoir sampling (see
// https://en.wikipedia.org/wiki/Reservoir_sampling#Algorithm_R).
//
// The zero value is not usable; use New to create a Sampler.
type Sampler[T cmp.Ordered] struct {
	size    int
	samples []T
	seen    int64
	mu      sync.Mutex
	rng     *rand.Rand

	// Caching of the last calculated percentile and its value.
	// See Percentile() for details.
	lastVal       T         // last sample at percentile lastP
	lastCalc      time.Time // zero if never calculated
	dirtySinceAdd bool      // true if Add() happened after the last calculation
	p             float64
}

func New[T cmp.Ordered](size int, p float64, rng *rand.Rand) *Sampler[T] {
	_ = "STUB: not implemented"
	return nil
}

// Clamp p to [0.0, 1.0]

// Add inserts an item into the reservoir with correct probability.
func (s *Sampler[T]) Add(item T) { _ = "STUB: not implemented"; return }

// Seen returns the number of items observed so far.
func (s *Sampler[T]) Seen() int64 { _ = "STUB: not implemented"; return 0 }

// Percentile returns the nearest-rank percentile value.
// Recalculation rules:
//   - If p changed since last call: recompute immediately.
//   - If new data arrived since last calc: recompute only if >= 5s have passed since last calc.
//   - Otherwise, return cached value.
func (s *Sampler[T]) Percentile() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

// If we have a cached value and p is unchanged:

// No new data, cached value is valid.

// Compute nearest-rank percentile.

// Clamp index to [0, n-1].

func nonDeterministicSeed() (uint64, uint64) { _ = "STUB: not implemented"; return 0, 0 }

package utils

import (
	"bytes"
	"context"
	"math/big"
	"math/rand"
	"reflect"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/protobuf/testing/protocmp"
)

// ReadOnly - if a struct embeds ReadOnly,
// its private fields will be compared by TestEqual.
type ReadOnly struct{}

// isReadOnly returns true if t embeds ReadOnly.
func isReadOnly(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func cmpComparer[T any, PT interface {
	Cmp(b *T) int
	*T
}](a PT, b PT) bool {
	_ = "STUB: not implemented"
	return false
}

var cmpOpts = []cmp.Option{
	protocmp.Transform(),
	cmp.Exporter(isReadOnly),
	cmpopts.EquateEmpty(),
	// Optimization for comparing slices of bytes.
	// Applies iff any of the slices is non-empty to avoid collision with EquateEmpty.
	cmp.FilterValues(func(x, y []byte) bool { return len(x) > 0 || len(y) > 0 }, cmp.Comparer(bytes.Equal)),
	cmp.Comparer(cmpComparer[big.Int]),
}

func OrPanic(err error) { _ = "STUB: not implemented"; return }

func OrPanic1[T any](v T, err error) T {
	_ = "STUB: not implemented"
	return *

	// TestDiff generates a human-readable diff between two objects.
	new(T)
}

func TestDiff[T any](want, got T) error { _ = "STUB: not implemented"; return nil }

// TestEqual is a more robust replacement for reflect.DeepEqual for tests.
func TestEqual[T any](a, b T) bool { _ = "STUB: not implemented"; return false }

// Thread-safe wrapper of rand.Rand.
type Rng struct{ inner *Mutex[*rand.Rand] }

func (rng Rng) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (rng Rng) Int63() int64 { _ = "STUB: not implemented"; return 0 }

func (rng Rng) Uint64() uint64 { _ = "STUB: not implemented"; return 0 }

func (rng Rng) Int() int { _ = "STUB: not implemented"; return 0 }

func (rng Rng) Intn(n int) int { _ = "STUB: not implemented"; return 0 }

func (rng Rng) Int63n(n int64) int64 { _ = "STUB: not implemented"; return 0 }

func (rng Rng) Shuffle(n int, swap func(i, j int)) { _ = "STUB: not implemented"; return }

// Split returns a new random number splitted from the given one.
// It should be used to provide deterministic rngs to independent goroutines.
// This is a very primitive splitting, known to result with dependent randomness.
// If that ever causes a problem, we can switch to SplitMix.
func (rng Rng) Split() Rng { _ = "STUB: not implemented"; return *new(Rng) }

// TestRng returns a deterministic random number generator.
func TestRng() Rng { _ = "STUB: not implemented"; return *new(Rng) }

func TestRngFromSeed(seed int64) Rng { _ = "STUB: not implemented"; return *new(Rng) }

func GenBool(rng Rng) bool { _ = "STUB: not implemented"; return false }

var alphanum = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// GenString generates a random string of length n.
func GenString(rng Rng, n int) string { _ = "STUB: not implemented"; return "" }

// Shuffle reorders the elements of s uniformly at random.
func Shuffle[T any](rng Rng, s []T) { _ = "STUB: not implemented"; return }

// GenBytes generates a random byte slice.
func GenBytes(rng Rng, n int) []byte { _ = "STUB: not implemented"; return nil }

// GenF is a function which generates T.
type GenF[T any] = func(rng Rng) T

// GenSlice generates a slice of small random length.
func GenSlice[T any](rng Rng, gen GenF[T]) []T { _ = "STUB: not implemented"; return nil }

// GenSliceN generates a slice of n elements.
func GenSliceN[T any](rng Rng, n int, gen GenF[T]) []T { _ = "STUB: not implemented"; return nil }

// GenMap generates a map of small random length.
func GenMap[K comparable, V any](rng Rng, genK GenF[K], genV GenF[V]) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

// GenMapN generates a map of n elements.
func GenMapN[K comparable, V any](rng Rng, n int, genK GenF[K], genV GenF[V]) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

// GenTimestamp generates a random timestamp.
func GenTimestamp(rng Rng) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Test tests whether reencoding a value is an identity operation.
func (c *ProtoConv[T, P]) Test(want T) error { _ = "STUB: not implemented"; return nil }

// IgnoreAfterCancel silently drops the error if the context is already canceled.
// Should be used for background tasks in tests, which cannot be guaranteed to exit gracefully.
// For example - if you have a tcp connection, then during cleanup one end will disconnect faster than the other,
// causing a race condition between context cancellation and disconnection error.
func IgnoreAfterCancel(ctx context.Context, err error) error { _ = "STUB: not implemented"; return nil }

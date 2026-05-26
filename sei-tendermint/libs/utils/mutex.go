package utils

import (
	"context"
	"iter"
	"sync"
	"sync/atomic"
)

// Mutex guards access to object of type T.
type Mutex[T any] struct {
	mu    sync.Mutex
	value T
}

// NewMutex creates a new Mutex with given object.
func NewMutex[T any](value T) (m Mutex[T]) {
	_ = "STUB: not implemented"

	// nolint:nakedret
	return nil
}

// Lock returns an iterator which locks the mutex and yields the guarded object.
// The mutex is unlocked when the iterator is done.
// If the mutex is nil, the iterator is a no-op.
func (m *Mutex[T]) Lock() iter.Seq[T] { _ = "STUB: not implemented"; return nil }

// Mutex guards access to object of type T.
type RWMutex[T any] struct {
	mu    sync.RWMutex
	value T
}

// NewMutex creates a new Mutex with given object.
func NewRWMutex[T any](value T) (m RWMutex[T]) {
	_ = "STUB: not implemented"

	// nolint:nakedret
	return nil
}

// Lock returns an iterator which locks the mutex and yields the guarded object.
// The mutex is unlocked when the iterator is done.
// If the mutex is nil, the iterator is a no-op.
func (m *RWMutex[T]) Lock() iter.Seq[T] { _ = "STUB: not implemented"; return nil }

// RLock returns an iterator which locks the mutex FOR READ and yields the guarded object.
// The mutex is unlocked when the iterator is done.
// If the mutex is nil, the iterator is a no-op.
func (m *RWMutex[T]) RLock() iter.Seq[T] { _ = "STUB: not implemented"; return nil }

// version of the value stored in an atomic watch.
type version[T any] struct {
	updated chan struct{}
	value   T
}

// newVersion constructs a new active version.
func newVersion[T any](value T) *version[T] { _ = "STUB: not implemented"; return nil }

type atomicWatch[T any] struct {
	ptr atomic.Pointer[version[T]]
}

type AtomicSend[T any] struct{ atomicWatch[T] }

func (w *AtomicSend[T]) Subscribe() AtomicRecv[T] { _ = "STUB: not implemented"; return nil }

// NewAtomicWatch creates a new AtomicWatch with the given initial value.
func NewAtomicSend[T any](value T) (w AtomicSend[T]) { _ = "STUB: not implemented"; return nil }

// nolint:nakedret

// Store updates the value of the atomic watch.
func (w *AtomicSend[T]) Store(value T) { _ = "STUB: not implemented"; return }

// AtomicRecv is a read-only reference to AtomicWatch.
type AtomicRecv[T any] struct{ *atomicWatch[T] }

// Load returns the current value of the atomic watch.
// Does not do any locking.
func (w *atomicWatch[T]) Load() T {
	_ = "STUB: not implemented"
	return *

	// Wait waits for the value of the atomic watch to satisfy the predicate.
	// Does not do any locking.
	new(T)
}

func (w *atomicWatch[T]) Wait(ctx context.Context, pred func(T) bool) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// Iter executes sequentially the function f on each value of the atomic watch.
// Context passed to f is canceled when the next value is available.
// Exits when the returned error is different from nil and context.Canceled,
// or when the context passed to Iter is canceled (after f exits).
func (w *atomicWatch[T]) Iter(ctx context.Context, f func(ctx context.Context, v T) error) error {
	_ = "STUB: not implemented"
	return nil
}

// WatchCtrl controls the locked object in a Watch.
// It is provided only in the iterator returned by Lock().
// Should NOT be stored anywhere.
type WatchCtrl struct {
	mu      sync.Mutex
	updated chan struct{}
}

// Watch stores a value of type T.
// Essentially a mutex, that can be awaited for updates.
type Watch[T any] struct {
	ctrl WatchCtrl
	val  T
}

// NewWatch constructs a new watch with the given value.
// Note that value in the watch cannot be changed, so T
// should be a pointer type if updates are required.
func NewWatch[T any](val T) Watch[T] { _ = "STUB: not implemented"; return nil }

// Wait waits for the value in the watch to be updated.
// Should be called only after locking the watch, i.e. within Lock() iterator.
// It unlocks -> waits for the update -> locks again.
func (c *WatchCtrl) Wait(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// WaitUntil waits for the value in the watch to satisfy the predicate.
// Should be called only after locking the watch, i.e. within Lock() iterator.
// The predicate is evaluated under the lock, so it can access the guarded object.
func (c *WatchCtrl) WaitUntil(ctx context.Context, pred func() bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Updated signals waiters that the value in the watch has been updated.
func (c *WatchCtrl) Updated() { _ = "STUB: not implemented"; return }

// Lock returns an iterator which locks the watch and yields the guarded object.
// The watch is unlocked when the iterator is done.
// If the watch is nil, the iterator is a no-op.
// Additionally the WatchCtrl object is provided to the yield function:
// * to unlock -> wait for the update -> lock again, call ctrl.Wait(ctx)
// * to signal an update, call ctrl.Updated().
func (w *Watch[T]) Lock() iter.Seq2[T, *WatchCtrl] { _ = "STUB: not implemented"; return nil }

// MonitorWatchUpdates calls f and checks if it has updated the watch.
func MonitorWatchUpdates[T any](w *Watch[T], f func()) bool {
	_ = "STUB: not implemented"
	return false
}

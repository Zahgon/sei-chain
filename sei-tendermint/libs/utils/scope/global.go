package scope

import (
	"context"
)

// GlobalHandle is a handle to a task spawned via SpawnGlobal.
type GlobalHandle[T any] struct {
	cancel context.CancelFunc
	done   chan struct{}
	res    T
}

// SpawnGlobal spawns a task in a global context.
// Use with care, as it is not tied to any scope and must be terminated manually by calling Terminate().
// The task does not return an error, because there is no canonical way to handle it.
// Can be used as an intermediate step when migrating code to use scopes.
func SpawnGlobal[T any](task func(ctx context.Context) T) *GlobalHandle[T] {
	_ = "STUB: not implemented"
	return nil
}

// WhileRunning restricts ctx to the lifetime of the task.
// WARNING: If the task is already finished, it SKIPs running f and returns context.Canceled.
func (h *GlobalHandle[T]) WhileRunning(ctx context.Context, f func(ctx context.Context) error) error {
	_ = "STUB: not implemented"
	return nil
}

// WhileRunning1 is like WhileRunning but for functions returning a value.
func WhileRunning1[R any, T any](ctx context.Context, h *GlobalHandle[T], f func(ctx context.Context) (R, error)) (res R, err error) {
	_ = "STUB: not implemented"
	// We need to set the error outside the closure, because
	// h.WhileRunning() may return context.Canceled if the task is already finished.
	return *new(R), nil
}

// Join awaits tasks completion.
func (h *GlobalHandle[T]) Join(ctx context.Context) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// Terminate cancels the task and waits for it to finish.
// Returns the task's result.
func (h *GlobalHandle[T]) Terminate() T { _ = "STUB: not implemented"; return *new(T) }

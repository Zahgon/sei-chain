package utils

import (
	"context"
)

// Semaphore provides a way to bound concurrenct access to a resource.
type Semaphore struct {
	ch chan struct{}
}

// NewSemaphore constructs a new semaphore with n permits.
func NewSemaphore(n int) *Semaphore { _ = "STUB: not implemented"; return nil }

// Acquire acquires a permit from the semaphore.
// Blocks until a permit is available.
func (s *Semaphore) Acquire(ctx context.Context) (release func(), err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

package scope

import (
	"sync"
	"sync/atomic"
)

type parallelScope struct {
	wg  sync.WaitGroup
	err atomic.Pointer[error]
}

// ParallelScope is a scope which doesn't require cancellation token,
// just parallelization.
type ParallelScope struct{ *parallelScope }

// Spawn spawns a new task in the scope.
func (s *parallelScope) Spawn(t func() error) { _ = "STUB: not implemented"; return }

// Parallel executes a function in parallel scope.
// Compared to Run, it does not allow for early cancellation,
// therefore is suitable for non-blocking computations.
// Returns the first error returned by any of the spawned tasks.
// Waits until all the tasks complete, before returning.
func Parallel(main func(ParallelScope) error) error { _ = "STUB: not implemented"; return nil }

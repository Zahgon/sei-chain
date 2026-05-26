package threading

import "sync"

var _ Pool = (*fixedPool)(nil)

// fixedPool is a pool of workers that can be used to execute tasks concurrently.
// More efficient than spawning large numbers of short lived goroutines.
type fixedPool struct {
	workQueue chan func()
	wg        sync.WaitGroup
	closeOnce sync.Once
	closed    bool
}

// Create a new work pool.
func NewFixedPool(
	// The name of the work pool. Used for metrics.
	name string,
	// The number of workers to create.
	workers int,
	// The size of the work queue. Once full, Submit will block until a slot is available.
	queueSize int,
) Pool {
	_ = "STUB: not implemented"
	return *new(Pool)
}

func (fp *fixedPool) Submit(task func()) { _ = "STUB: not implemented"; return }

func (fp *fixedPool) Close() { _ = "STUB: not implemented"; return }

func (fp *fixedPool) worker() { _ = "STUB: not implemented"; return }

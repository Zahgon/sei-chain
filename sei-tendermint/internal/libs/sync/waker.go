package sync

// Waker is used to wake up a sleeper when some event occurs. It debounces
// multiple wakeup calls occurring between each sleep, and wakeups are
// non-blocking to avoid having to coordinate goroutines.
type Waker struct {
	wakeCh chan struct{}
}

// NewWaker creates a new Waker.
func NewWaker() *Waker { _ = "STUB: not implemented"; return nil }

// buffer used for debouncing

// Sleep returns a channel that blocks until Wake() is called.
func (w *Waker) Sleep() <-chan struct{} {
	_ = "STUB: not implemented"

	// Wake wakes up the sleeper.
	return nil
}

func (w *Waker) Wake() {
	_ = "STUB: not implemented"
	// A non-blocking send with a size 1 buffer ensures that we never block, and
	// that we queue up at most a single wakeup call between each Sleep().
	return
}

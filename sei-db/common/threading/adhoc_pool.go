package threading

import "sync"

var _ Pool = (*adHocPool)(nil)

// adHocPool is a Pool that runs each task in a new goroutine.
// Intended for use in unit tests or where performance is not important.
type adHocPool struct {
	wg     sync.WaitGroup
	closed bool
}

// NewAdHocPool creates a Pool that runs each submitted task in a one-off goroutine.
func NewAdHocPool() Pool { _ = "STUB: not implemented"; return *new(Pool) }

func (p *adHocPool) Submit(task func()) { _ = "STUB: not implemented"; return }

func (p *adHocPool) Close() { _ = "STUB: not implemented"; return }

package consensus

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

// TimeoutTicker is a timer that schedules timeouts
// conditional on the height/round/step in the timeoutInfo.
// The timeoutInfo.Duration may be non-positive.
type TimeoutTicker interface {
	Run(context.Context) error
	Chan() <-chan timeoutInfo       // on which to receive a timeout
	ScheduleTimeout(ti timeoutInfo) // reset the timer
}

// timeoutTicker wraps time.Timer,
// scheduling timeouts only for greater height/round/step
// than what it's already seen.
// Timeouts are scheduled along the tickChan,
// and fired on the tockChan.
type timeoutTicker struct {
	tick     utils.Mutex[*utils.AtomicSend[utils.Option[timeoutInfo]]] // for scheduling timeouts
	tockChan chan timeoutInfo                                          // for notifying about them
}

// NewTimeoutTicker returns a new TimeoutTicker.
func NewTimeoutTicker() TimeoutTicker { _ = "STUB: not implemented"; return *new(TimeoutTicker) }

// Chan returns a channel on which timeouts are sent.
func (t *timeoutTicker) Chan() <-chan timeoutInfo {
	_ = "STUB: not implemented"

	// ScheduleTimeout schedules a new timeout, which replaces the previous one.
	// Noop if a timeout for a later height/round/step has been already scheduled.
	return nil
}

func (t *timeoutTicker) ScheduleTimeout(newti timeoutInfo) { _ = "STUB: not implemented"; return }

func (t *timeoutTicker) tickSubscribe() utils.AtomicRecv[utils.Option[timeoutInfo]] {
	_ = "STUB: not implemented"
	return nil
}

// timers are interupted and replaced by new ticks from later steps
// timeouts of 0 on the tickChan will be immediately relayed to the tockChan
func (t *timeoutTicker) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// last fired timeout

// Task measuring timeouts. Owns tockSend.

// Task reporting timeouts via channel. Owns tockRecv.
// TODO(gprusak): it would be better to expose t.tock directly,
// however the receiving task doesn't support receiving from AtomicWatch yet.

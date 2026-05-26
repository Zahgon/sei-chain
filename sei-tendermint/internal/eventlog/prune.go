package eventlog

import (
	"time"
)

// checkPrune checks whether the log has exceeded its boundaries of size or
// age, and if so prunes the log and updates the head.
func (lg *Log) checkPrune(head *logEntry, size int, age time.Duration) error {
	_ = "STUB: not implemented"
	// To avoid potentially re-pruning for every event, don't trigger an age
	// prune until we're at least this far beyond the designated size.
	return nil
}

// no pruning is needed

// We exceeded the size cap. In this case, age does not matter: count off
// the newest items and drop the unconsumed tail. Note that we prune by a
// fraction rather than an absolute amount so that we only have to prune
// for size occasionally.

// TODO(creachadair): We may want to spill dropped events to secondary
// storage rather than dropping them. The size cap is meant as a safety
// valve against unexpected extremes, but if a network has "expected"
// spikes that nevertheless exceed any safe buffer size (e.g., Osmosis
// epochs), we may want to have a fallback so that we don't lose events
// that would otherwise fall within the window.

// We did not exceed the size cap, but some items are too old.

// Note that when we update the head after pruning, we do not need to signal
// any waiters; pruning never adds new material to the log so anyone waiting
// should continue doing so until a subsequent Add occurs.

// pruneSize returns a new log state by pruning head to newSize.
// Precondition: newSize ≤ len(head).
func (lg *Log) pruneSize(head *logEntry, newSize int) (logState, error) {
	_ = "STUB: not implemented"
	// Special case for size 0 to simplify the logic below.
	return *new(logState), nil
}

// drop everything

// Initialize: New head has the same item as the old head.
// new head
// new tail (last copied cons)

// pruneAge returns a new log state by pruning items older than the window
// prior to the head element.
func (lg *Log) pruneAge(head *logEntry) logState { _ = "STUB: not implemented"; return *new(logState) }

// all remaining items are older than the window

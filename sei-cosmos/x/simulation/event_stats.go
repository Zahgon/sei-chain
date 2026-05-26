package simulation

import (
	"io"
)

// EventStats defines an object that keeps a tally of each event that has occurred
// during a simulation.
type EventStats map[string]map[string]map[string]int

// NewEventStats creates a new empty EventStats object
func NewEventStats() EventStats {
	_ = "STUB: not implemented"
	return *

	// Tally increases the count of a simulation event.
	new(EventStats)
}

func (es EventStats) Tally(route, op, evResult string) { _ = "STUB: not implemented"; return }

// Print the event stats in JSON format.
func (es EventStats) Print(w io.Writer) { _ = "STUB: not implemented"; return }

// ExportJSON saves the event stats as a JSON file on a given path
func (es EventStats) ExportJSON(path string) { _ = "STUB: not implemented"; return }

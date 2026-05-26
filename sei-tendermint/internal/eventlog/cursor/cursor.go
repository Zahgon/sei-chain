// Package cursor implements time-ordered item cursors for an event log.
package cursor

import (
	"time"
)

// A Source produces cursors based on a time index generator and a sequence
// counter. A zero-valued Source is ready for use with defaults as described.
type Source struct {
	// This function is called to produce the current time index.
	// If nil, it defaults to time.Now().UnixNano().
	TimeIndex func() int64

	// The current counter value used for sequence number generation.  It is
	// incremented in-place each time a cursor is generated.
	Counter int64
}

func (s *Source) timeIndex() int64 { _ = "STUB: not implemented"; return 0 }

func (s *Source) nextCounter() int64 { _ = "STUB: not implemented"; return 0 }

// Cursor produces a fresh cursor from s at the current time index and counter.
func (s *Source) Cursor() Cursor { _ = "STUB: not implemented"; return *new(Cursor) }

//nolint:gosec // timeIndex returns a non-negative value
//nolint:gosec // masked

// A Cursor is a unique identifier for an item in a time-ordered event log.
// It is safe to copy and compare cursors by value.
type Cursor struct {
	timestamp uint64 // ns since Unix epoch
	sequence  uint16 // sequence number
}

// Before reports whether c is prior to o in time ordering. This comparison
// ignores sequence numbers.
func (c Cursor) Before(o Cursor) bool { _ = "STUB: not implemented"; return false }

// Diff returns the time duration between c and o. The duration is negative if
// c is before o in time order.
func (c Cursor) Diff(o Cursor) time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

//nolint:gosec // timestamps are ns since epoch; values within valid range for time.Duration

// IsZero reports whether c is the zero cursor.
func (c Cursor) IsZero() bool { _ = "STUB: not implemented"; return false }

// MarshalText implements the encoding.TextMarshaler interface.
// A zero cursor marshals as "", otherwise the format used by the String method.
func (c Cursor) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// An empty text unmarshals without error to a zero cursor.
func (c *Cursor) UnmarshalText(data []byte) error { _ = "STUB: not implemented"; return nil }

// set zero

// String returns a printable text representation of a cursor.
func (c Cursor) String() string { _ = "STUB: not implemented"; return "" }

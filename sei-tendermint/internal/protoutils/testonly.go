package protoutils

import (
	"iter"
)

// Test tests whether reencoding a value is an identity operation.
func (c *Conv[T, P]) Test(want T) error { _ = "STUB: not implemented"; return nil }

// Check that Decode does not panic on any malformed version of p.
// The malformed values might or might not be parseable - here we
// only check that Decode does not panic.

// Iterates over copies of msg with exactly one transitive field set to nil.
// This simulates situations in which a malicious proto value has been received.
// Note that setting to nil a single entry of a slice representing a repeated field
// is NOT a feasible malformed message (proto.Unmarshal would never return such a result).
func iterMalformed[M Message](msg M) iter.Seq[M] { _ = "STUB: not implemented"; return nil }

// Clear the field, then yield a clone of the top level message, then set the field to a clone we did beforehand.
// We cannot v := Get -> Clear -> Set(v), because Get for repeated values returns a reference (i.e. Clear would destroy it).

// Iterate recursively in case the field was a message/contained messages.

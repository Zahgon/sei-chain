package wal

import (
	"github.com/sei-protocol/sei-chain/sei-db/proto"
)

// ChangelogWAL is a type alias for a WAL specialized for ChangelogEntry.
type ChangelogWAL = GenericWAL[proto.ChangelogEntry]

// NewChangelogWAL creates a new WAL for ChangelogEntry.
// This is a convenience wrapper that handles serialization automatically.
func NewChangelogWAL(dir string, config Config) (ChangelogWAL, error) {
	_ = "STUB: not implemented"
	return *new(ChangelogWAL), nil
}

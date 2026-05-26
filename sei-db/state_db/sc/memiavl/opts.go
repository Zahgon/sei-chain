package memiavl

import (
	"time"
)

// Options contains all settings for opening a memiavl database.
// It embeds Config for the configurable settings and adds runtime-specific options.
type Options struct {
	Config // Embedded config for all configurable settings

	// Dir is the directory path for the memiavl database
	Dir string
	// CreateIfMissing creates the database if it doesn't exist
	CreateIfMissing bool
	// InitialVersion is the initial version number
	InitialVersion uint32
	// ReadOnly opens the database in read-only mode
	ReadOnly bool
	// InitialStores are the initial store names when initializing an empty instance
	InitialStores []string
	// ZeroCopy if true, get and iterator methods return slices pointing to mmaped blob files
	ZeroCopy bool
	// LoadForOverwriting if true, rollbacks the state by truncating versions after TargetVersion
	LoadForOverwriting bool
	// OnlyAllowExportOnSnapshotVersion restricts export to snapshot versions only
	OnlyAllowExportOnSnapshotVersion bool

	// snapshotMinTimeIntervalDuration is the converted Duration from Config.SnapshotMinTimeInterval
	// This is populated by FillDefaults()
	snapshotMinTimeIntervalDuration time.Duration
}

// SnapshotMinTimeDuration returns the minimum time interval between snapshots as a Duration.
// Call FillDefaults() before using this method.
func (opts *Options) SnapshotMinTimeDuration() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (opts Options) Validate() error { _ = "STUB: not implemented"; return nil }

func (opts *Options) FillDefaults() { _ = "STUB: not implemented"; return }

// SnapshotWriterLimit controls tree concurrency but not I/O rate (use SnapshotWriteRateMBps for that)

// Convert SnapshotMinTimeInterval (seconds) to Duration

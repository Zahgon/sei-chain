package benchmark

import (
	"math/rand"
	"time"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/util"
)

// CohortFileExtension is the file extension used for cohort files.
const CohortFileExtension = ".cohort"

// CohortSwapFileExtension is the file extension used for cohort swap files. Used to atomically update cohort files.
const CohortSwapFileExtension = CohortFileExtension + util.SwapFileExtension

/* The lifecycle of a cohort:

    +-----+     +-----------+     +----------+     +---------+
    | new | --> | exhausted | --> | complete | --> | expired |
    +-----+     +-----------+     +----------+     +---------+
       |              |
       v              |
    +-----------+     |
    | abandoned | <---|
    +-----------+

- new: the cohort was just created and is currently being used to supply keys for writing.
- exhausted: all keys in the cohort have been scheduled for writing, but the DB may not have ingested them all yet.
- complete: all keys in the cohort have been written to the DB and are safe to read.
- abandoned: before becoming complete, the benchmark was restarted. It will never be thread safe to read or write
              any keys in this cohort.
- expired: the cohort has been marked as complete, but it can no longer be read because the TTL has expired
            (or is about to expire).
*/

// A Cohort is a grouping of key-value pairs used for benchmarking.
//
// If a benchmark wants to read values, it must somehow figure out which keys have been written to the database.
// If it wants to verify the validity of the data it reads, it must also be able to determine the correct value
// that should be associated with any particular key, and it must also be able to determine when keys are
// expected to be removed from the database due to TTL expiration.
//
// Tracking the sort of metadata required to do reads in a benchmark is not a trivial thing, especially when
// the scale of the benchmark is large (i.e. tens or hundreds of millions of keys over weeks or months of time).
// Storing this information in memory is simply not plausible, and storing it on disk requires database scale similar
// to what LittDB is handling, unless we are clever about it. A "cohort" is that clever mechanism. Each cohort tracks a
// large collection of key-value pairs in the database, and it does it in a way that uses very little disk space.
//
// Key-value pairs each have unique indices, and knowing the index of a key-value pair allows the data to be
// regenerated deterministically. All key-value pairs in a cohort have sequential indices. A single cohort can
// track multiple gigabytes worth of key-value pairs, but on disk it only requires a few dozen bytes of data.
type Cohort struct {
	// The directory where the cohort file is stored.
	parentDirectory string

	// The unique ID of this cohort.
	cohortIndex uint64

	// The index of the first key-value pair in the cohort.
	lowKeyIndex uint64

	// The index of the last key-value pair in the cohort.
	highKeyIndex uint64

	// The size of the values written in this cohort.
	valueSize uint64

	// The next available index to be written. Only relevant for a new cohort that is currently being written to
	// the DB. This value is undefined for cohorts that have been completely written or loaded from disk. This value
	// is NOT serialized to disk.
	nextKeyIndex uint64

	// True iff all key-value pairs in the cohort have been written to the database.
	allValuesWritten bool

	// A timestamp that is guaranteed to come before the first value in the cohort is written to the database.
	firstValueTimestamp time.Time

	// True iff the cohort has been loaded from disk. This value is NOT serialized to disk.
	loadedFromDisk bool

	// Whether fsync mode is enabled. Disable for faster unit tests.
	fsync bool
}

// NewCohort creates a new cohort with the given index range.
func NewCohort(
	parentDirectory string,
	cohortIndex uint64,
	lowIndex uint64,
	highIndex uint64,
	valueSize uint64,
	fsync bool) (*Cohort, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoadCohort loads a cohort from the given path.
func LoadCohort(path string) (*Cohort, error) { _ = "STUB: not implemented"; return nil, nil }

// Cohort file names are in the format "X.cohort", where X is the cohort index.
// Replacing ".cohort" with an empty string gives us the cohort index in string form.

//nolint:gosec // path within cohort directory

//nolint:gosec // path within cohort directory

// NextCohort creates the next cohort in the sequence with the given number of keys.
func (c *Cohort) NextCohort(keyCount uint64, valueSize uint64) (*Cohort, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CohortIndex returns the index of the cohort.
func (c *Cohort) CohortIndex() uint64 { _ = "STUB: not implemented"; return 0 }

// LowKeyIndex returns the index of the first key in the cohort.
func (c *Cohort) LowKeyIndex() uint64 { _ = "STUB: not implemented"; return 0 }

// HighKeyIndex returns the index of the last key in the cohort.
func (c *Cohort) HighKeyIndex() uint64 { _ = "STUB: not implemented"; return 0 }

func (c *Cohort) ValueSize() uint64 {
	_ = "STUB: not implemented"

	// FirstValueTimestamp returns the timestamp of the first value in the cohort.
	return 0
}

func (c *Cohort) FirstValueTimestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// IsComplete returns true if all key-value pairs in the cohort have been written to the database. Only complete
// cohorts are safe to read from.
func (c *Cohort) IsComplete() bool { _ = "STUB: not implemented"; return false }

// IsExhausted returns true if the cohort has been exhausted, i.e. it has produced all keys for writing that it is
// capable of producing. Once exhausted, a cohort should be marked as completed once all key-value pairs have been
// written to the database, thus making all keys in the cohort safe to read.
func (c *Cohort) IsExhausted() bool { _ = "STUB: not implemented"; return false }

// IsLoadedFromDisk returns true if the cohort has been loaded from disk.
func (c *Cohort) IsLoadedFromDisk() bool { _ = "STUB: not implemented"; return false }

// GetKeyIndexForWriting gets the next key to be written to the database.
func (c *Cohort) GetKeyIndexForWriting() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// GetKeyIndexForReading gets a random key from the cohort that is safe to read. This function should only be called
// after the cohort has been marked as complete.
func (c *Cohort) GetKeyIndexForReading(rand *rand.Rand) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// sanity check

// MarkComplete marks that all key-value pairs in the cohort have been written to the database. Once done,
// all key-value pairs in the cohort become safe to read, so long as the cohort has not yet expired. A cohort
// is said to have expired when it is possible that at least one key in the cohort may be deleted from the DB
// due to the TTL.
func (c *Cohort) MarkComplete() error { _ = "STUB: not implemented"; return nil }

// Path returns the file path of the cohort file.
func (c *Cohort) Path() string { _ = "STUB: not implemented"; return "" }

// Write the data in this cohort to its file on disk. When this method returns, the cohort file is guaranteed to be
// crash durable.
func (c *Cohort) Write() error { _ = "STUB: not implemented"; return nil }

// serialize serializes the cohort to a byte array.
func (c *Cohort) serialize() []byte {
	_ = "STUB: not implemented"
	// Data size:
	//   - cohortIndex (8 bytes)
	//   - lowKeyIndex (8 bytes)
	//   - highKeyIndex (8 bytes)
	//   - valueSize (8 bytes)
	//   - firstValueTimestamp (8 bytes)
	//   - allValuesWritten (1 byte)
	//
	// Total: 41 bytes
	return nil
}

//nolint:gosec // wall-clock seconds non-negative

func (c *Cohort) deserialize(data []byte) error { _ = "STUB: not implemented"; return nil }

//nolint:gosec // wall-clock seconds fit int64

// IsExpired returns true if the cohort has expired (i.e. it is no longer safe to read).
func (c *Cohort) IsExpired(now time.Time, maxAge time.Duration) bool {
	_ = "STUB: not implemented"
	return false
}

// Incomplete cohorts loaded from disk are instantly expired.

// A cohort currently in the process of being written can't expire.

// Delete the associated cohort file.
func (c *Cohort) Delete() error { _ = "STUB: not implemented"; return nil }

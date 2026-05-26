// Crash-safe A/B file persistence.
//
// # A/B File Strategy
//
// We use an A/B file pair (<prefix>_a.pb/<prefix>_b.pb) instead of the traditional
// temp-file-then-rename approach:
//
//   - Traditional approach: write to temp file, fsync, rename to target, fsync directory
//   - A/B approach: alternate writes between <prefix>_a.pb and <prefix>_b.pb, fsync after write
//
// Why A/B files?
//
//   - Traditional temp+rename requires directory sync after file sync to ensure
//     the rename (directory entry update) is durable — that's an extra disk operation
//   - With A/B files, we only need one file sync per write
//   - Directory sync is only needed when a file is first created (at most twice total)
//   - Safety comes from redundancy: while writing to A, B is untouched; while writing
//     to B, A is untouched. A crash only corrupts the file being written.
//   - On load, we read both files and pick the one with the higher seq
//
// # Recovery Behavior
//
//   - Fresh start (files don't exist): Returns ErrNoData
//   - One file corrupt (e.g. crash during write): Uses the other file; logged at WARN
//   - Both files corrupt: Returns error (real data loss)
//   - OS-level errors (permission denied, I/O): Propagated immediately
//
// # Write Behavior
//
//   - State directory must already exist (we do not create it).
//   - Writes are synchronous (fsync after each write).
//   - Writes are idempotent, so retries on next state change are safe.
//   - Seq is only advanced after a successful write (rollback on failure).
package persist

import (
	"errors"
	"hash/crc32"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/protoutils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

// A/B file suffixes.
const (
	suffixA = "_a.pb"
	suffixB = "_b.pb"
)

var crc32c = crc32.MakeTable(crc32.Castagnoli)

const (
	crcSize    = 4                 // CRC32-C prefix length
	seqSize    = 8                 // uint64 little-endian
	headerSize = crcSize + seqSize // file header: [4-byte CRC32-C BE][8-byte seq LE]
)

// ErrNoData is returned by loadPersisted when no persisted files exist for the prefix.
var ErrNoData = errors.New("no persisted data")

// ErrCorrupt indicates that a persisted file exists but contains invalid data
// (e.g. partially written during a crash). loadPersisted tolerates one corrupt
// file and falls back to the other A/B copy. OS-level errors (permission denied,
// I/O errors) are NOT wrapped with ErrCorrupt and cause loadPersisted to fail.
var ErrCorrupt = errors.New("corrupt persisted data")

// dataWithSeq is the unit stored in each A/B file: a sequence number and a proto payload.
type dataWithSeq struct {
	seq  uint64
	data []byte // nil on fresh start
}

// Persister[T] is a strongly-typed persister for a proto message type.
type Persister[T protoutils.Message] interface {
	Persist(T) error
}

type noopPersister[T protoutils.Message] struct{}

func (noopPersister[T]) Persist(T) error {
	_ = "STUB: not implemented"

	// newNoOpPersister returns a Persister that silently discards all writes.
	return nil
}

func newNoOpPersister[T protoutils.Message]() Persister[T] { _ = "STUB: not implemented"; return nil }

// abPersister writes data to A/B files with automatic seq management.
// File format: [4-byte CRC32-C BE] [8-byte seq LE] [proto-marshalled message].
// Only created when config has a state dir; dir is always a valid path.
// File selection is derived from seq: odd seq → A, even seq → B.
type abPersister[T protoutils.Message] struct {
	dir    string
	prefix string
	seq    uint64
}

// NewPersister creates a crash-safe persister for the given directory and prefix.
// When dir is None, returns a no-op persister that discards all writes.
// When dir is Some, it must already exist and be a directory; returns error otherwise.
// Also returns the loaded message (None on fresh start or no-op) for the caller to use.
// This encapsulates all on-disk format details (A/B files, seq wrapper, proto marshal) in one place.
func NewPersister[T protoutils.Message](dir utils.Option[string], prefix string) (Persister[T], utils.Option[T], error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Probe writability by creating and removing a temp file. Checking permission
// bits is not reliable: Windows emulates Unix bits from ACLs, root bypasses
// permission checks on Unix, and group/other write bits are easy to miss.

// Ensure both A/B files exist and are writable so Persist never creates new
// directory entries. Empty files are treated as non-existent by loadFile,
// so they won't interfere with loading on restart.

//nolint:gosec // path is stateDir + hardcoded suffix; not user-controlled

// Sync directory to make file entries durable (harmless if files already existed).
//nolint:gosec // d is operator-configured stateDir; not user-controlled

// Persist writes a proto message to persistent storage. Not safe for concurrent use.
func (w *abPersister[T]) Persist(msg T) error { _ = "STUB: not implemented"; return nil }

// Odd seq → A, even seq → B.

// loadFile reads a single A/B file and returns its contents as a dataWithSeq.
// Returns os.ErrNotExist when the file does not exist.
// Returns ErrCorrupt on CRC mismatch or truncated header.
// OS-level errors (permission denied, I/O) are returned unwrapped.
func loadFile(stateDir, filename string) (dataWithSeq, error) {
	_ = "STUB: not implemented"
	return *new(dataWithSeq), nil
}

//nolint:gosec // path is constructed from operator-configured stateDir + hardcoded filename suffix; no user-controlled input

// Empty files are created by NewPersister to pre-populate directory entries.

// loadPersisted loads persisted data for the given directory and prefix.
// Tries both A and B files; if one is corrupt (e.g. crash during write), the other is used
// so the validator can restart. Returns ErrNoData when no persisted files exist (use errors.Is).
// Returns other error only when both files fail to load or state is inconsistent (same seq).
func loadPersisted(dir string, prefix string) (dataWithSeq, error) {
	_ = "STUB: not implemented"
	return *new(dataWithSeq), nil
}

// Fail fast on OS-level errors (permission denied, I/O errors).
// Only ErrNotExist (fresh start) and ErrCorrupt (crash mid-write) are tolerable.

// writeAndSync atomically replaces path contents with data (O_TRUNC) and fsyncs.
// Used by WAL persistence (blocks, commitqcs).
func writeAndSync(path string, data []byte) error { _ = "STUB: not implemented"; return nil }

//nolint:gosec // path is stateDir + hardcoded suffix; not user-controlled

// writeFile writes an A/B state file: [4-byte CRC32-C BE][8-byte seq LE][proto data].
// Encodes seq and computes CRC internally; writes chunks directly to avoid
// copying data into an intermediate buffer. The file is fsynced before return.
func writeFile(path string, d dataWithSeq) error { _ = "STUB: not implemented"; return nil }

//nolint:gosec // path is stateDir + hardcoded suffix; not user-controlled

// hash.Hash.Write never returns an error.

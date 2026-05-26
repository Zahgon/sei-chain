package segment

import (
	"bufio"
	"log/slog"
	"os"
	"sync/atomic"
)

// ValuesFileExtension is the file extension for the values file. This file contains the values for the data
// segment. Value files are written in the form "X-Y.values", where X is the segment index and Y is the shard number.
const ValuesFileExtension = ".values"

// valueFile represents a file that stores values.
type valueFile struct {
	// The logger for the value file.
	logger *slog.Logger

	// The segment index.
	index uint32

	// The shard number of this value file.
	shard uint8

	// Path data for the segment file.
	segmentPath *SegmentPath

	// The file wrapped by the writer. If the file is sealed, this value is nil.
	file *os.File

	// The writer for the file. If the file is sealed, this value is nil.
	writer *bufio.Writer

	// The current size of the file in bytes. Includes both flushed and unflushed data.
	size uint64

	// The current size of the file, only including flushed data. Protects against reads of partially written values.
	flushedSize atomic.Uint64

	// Whether fsync mode is enabled. If fsync mode is enabled, then each flush operation will invoke the OS fsync
	// operation before returning. An fsync operation is required to ensure that data is not sitting in OS level
	// in-memory buffers (otherwise, an OS crash may lead to data loss). This option is provided for testing,
	// as many test scenarios do lots of tiny writes and flushes, and this workload is MUCH slower with fsync
	// mode enabled. In production, fsync mode should always be enabled.
	fsync bool
}

// createValueFile creates a new value file.
func createValueFile(
	logger *slog.Logger,
	index uint32,
	shard uint8,
	segmentPath *SegmentPath,
	fsync bool,
) (*valueFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Open the file for writing.
//nolint:gosec // path validated by segment manager

// loadValueFile loads a value file from disk. It looks for the file in the given parent directories until it finds
// the file. If the file is not found, it returns an error.
func loadValueFile(
	logger *slog.Logger,
	index uint32,
	shard uint8,
	segmentPaths []*SegmentPath) (*valueFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // file size is non-negative

// getValueFileIndex returns the index of the value file from the file name. Value file names have the form
// "X-Y.values", where X is the segment index and Y is the shard number.
func getValueFileIndex(fileName string) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:gosec // segment index fits uint32

// getValueFileShard returns the shard number of the value file from the file name. Value file names have the form
// "X-Y.values", where X is the segment index and Y is the shard number.
func getValueFileShard(fileName string) (uint8, error) { _ = "STUB: not implemented"; return 0, nil }

// Size returns the size of the value file in bytes.
func (v *valueFile) Size() uint64 {
	_ = "STUB: not implemented"

	// name returns the name of the value file.
	return 0
}

func (v *valueFile) name() string { _ = "STUB: not implemented"; return "" }

// path returns the path to the value file.
func (v *valueFile) path() string { _ = "STUB: not implemented"; return "" }

// read reads a value from the value file.
func (v *valueFile) read(firstByteIndex uint32) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // path validated by segment manager

// Read the length of the value.

// Read the value itself.

//nolint:gosec // bytesRead bounded by length

// write writes a value to the value file, returning the index of the first byte written.
func (v *valueFile) write(value []byte) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

// No matter what, we can't start a new value if its first byte would be beyond position 2^32.
// This is because we only have 32 bits in an address to store the position of a value's first byte.

// First, write the length of the value.
//nolint:gosec // value length fits uint32

// Then, write the value itself.

//nolint:gosec // value length non-negative

// flush writes all unflushed data to disk.
func (v *valueFile) flush() error { _ = "STUB: not implemented"; return nil }

// It is now safe to read the flushed bytes directly from the file.

// seal seals the value file.
func (v *valueFile) seal() error { _ = "STUB: not implemented"; return nil }

// snapshot creates a hard link to the file in the snapshot directory, and a soft link to the hard linked file in the
// soft link directory. Requires that the file is sealed and that snapshotting is enabled.
func (v *valueFile) snapshot() error { _ = "STUB: not implemented"; return nil }

// delete deletes the value file.
func (v *valueFile) delete() error { _ = "STUB: not implemented"; return nil }

// As an extra safety check, make it so that all future reads fail before they do I/O.

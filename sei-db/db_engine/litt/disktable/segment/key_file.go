package segment

import (
	"bufio"
	"log/slog"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/types"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/util"
)

// KeyFileExtension is the file extension for the keys file. This file contains the keys for the data segment,
// and is used for performing garbage collection on the keymap. It can also be used to rebuild the keymap.
const KeyFileExtension = ".keys"

// KeyFileSwapExtension is the file extension for the keys swap file. This file is used to atomically
// update key files.
const KeyFileSwapExtension = KeyFileExtension + util.SwapFileExtension

// keyFile tracks the keys in a segment. It is used to do garbage collection on the keymap.
//
// This struct is NOT goroutine safe. It is unsafe to concurrently call write, flush, or seal on the same key file.
// It is not safe to read a key file until it is sealed. Once sealed, read only operations are goroutine safe.
type keyFile struct {
	// The logger for the key file.
	logger *slog.Logger

	// The segment index.
	index uint32

	// Path data for the segment file.
	segmentPath *SegmentPath

	// The writer for the file. If the file is sealed, this value is nil.
	writer *bufio.Writer

	// The size of the key file in bytes.
	size uint64

	// The segment version. Determines serialization format.
	segmentVersion SegmentVersion

	// If true, then this key file is intended to replace another key file. It is written to a temporary
	// file, and then atomically renamed to the final file name.
	swap bool
}

// newKeyFile creates a new key file.
func createKeyFile(
	logger *slog.Logger,
	index uint32,
	segmentPath *SegmentPath,
	swap bool,
) (*keyFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // path validated by segment manager

// loadKeyFile loads the key file from disk, looking in the given parent directories until it finds the file.
// If the file is not found, it returns an error.
func loadKeyFile(
	logger *slog.Logger,
	index uint32,
	segmentPaths []*SegmentPath,
	segmentVersion SegmentVersion,
) (*keyFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // file size is non-negative

// Size returns the size of the key file in bytes.
func (k *keyFile) Size() uint64 {
	_ = "STUB: not implemented"

	// name returns the name of the key file.
	return 0
}

func (k *keyFile) name() string { _ = "STUB: not implemented"; return "" }

// path returns the full path to the key file.
func (k *keyFile) path() string { _ = "STUB: not implemented"; return "" }

// atomicSwap atomically replaces the key file, replacing the old one.
func (k *keyFile) atomicSwap(sync bool) error { _ = "STUB: not implemented"; return nil }

// write writes a key to the key file.
func (k *keyFile) write(scopedKey *types.ScopedKey) error { _ = "STUB: not implemented"; return nil }

// Write the length of the key.
//nolint:gosec // key length fits uint32

// Write the key itself.

// Write the serialized address (which includes the shard ID and value size).

//nolint:gosec // sizes are non-negative
/* uint32 size of key */

// getKeyFileIndex returns the index of the key file from the file name. Key file names have the form "X.keys",
// where X is the segment index.
func getKeyFileIndex(fileName string) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:gosec // segment index fits uint32

// flush flushes the key file to disk.
func (k *keyFile) flush() error { _ = "STUB: not implemented"; return nil }

// seal seals the key file, preventing further writes.
func (k *keyFile) seal() error { _ = "STUB: not implemented"; return nil }

// readKeys reads all keys from the key file. This method returns an error if the key file is not sealed.
// If there are keys that were only partially written (i.e. keys being written when the process crashed), then
// those keys may not be returned. If a key is returned, it is guaranteed to be "whole" (i.e. a partial key will
// never be returned).
func (k *keyFile) readKeys() ([]*types.ScopedKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Key files are small as long as key length is sane. Safe to read the whole file into memory.

// We need at least 4 bytes to read the length of the key.
//nolint:staticcheck // QF1006
// There are fewer than 4 bytes left in the file.

// We need to read the key, as well as the serialized address (which embeds the shard ID and value size).

// There are insufficient bytes left in the file to read the key and address.

// This can happen if there is a crash while we are writing to the key file.
// Recoverable, but best to note the event in the logs.

// snapshot creates a hard link to the file in the snapshot directory, and a soft link to the hard linked file in the
// soft link directory. Requires that the file is sealed and that snapshotting is enabled.
func (k *keyFile) snapshot() error { _ = "STUB: not implemented"; return nil }

// delete deletes the key file. If this key_file is a snapshot file (i.e. it is backed by a symlink), this method will
// also delete the file pointed to by the symlink.
func (k *keyFile) delete() error { _ = "STUB: not implemented"; return nil }

// isSealed returns true if the key file is sealed, and false otherwise.
func (k *keyFile) isSealed() bool { _ = "STUB: not implemented"; return false }

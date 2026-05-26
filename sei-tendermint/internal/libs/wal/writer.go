package wal

import (
	"bufio"
	"os"
)

type logWriter struct {
	file      *os.File
	buf       *bufio.Writer
	bytesSize int64
}

// Returns the size of the file, ignoring the last truncated entry.
// WARNING it needs to read the whole file.
func realFileSize(path string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func sync(path string) error { _ = "STUB: not implemented"; return nil }

func openLogWriter(path string) (res *logWriter, resErr error) {
	_ = "STUB: not implemented"
	// Read the whole file and find the non-corrupted prefix.
	return nil, nil
}

// Sync the directory containing the file:
// realFileSize() may have created a file if it didn't exist.
// In that case we need the directory synced, so that file's inode
// is not lost in case of crash.

// Truncate the file to non-corrupted prefix and sync.
// It is still not 100% corruption-proof,
// but we would need a rolling checksums to fix that,
// which would require redesigning the WAL format.

func (w *logWriter) AppendEntry(entry []byte) (err error) { _ = "STUB: not implemented"; return nil }

//nolint:gosec // WAL entries are bounded by max message size; no overflow risk

func (w *logWriter) Sync() (err error) { _ = "STUB: not implemented"; return nil }

// Close unconditionally releases all the resources.
func (w *logWriter) Close() { _ = "STUB: not implemented"; return }

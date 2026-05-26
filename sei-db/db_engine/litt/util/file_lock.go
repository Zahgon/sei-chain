package util

import (
	"log/slog"
	"os"
)

// FileLock represents a file-based lock
type FileLock struct {
	logger *slog.Logger
	path   string
	file   *os.File
}

// IsProcessAlive checks if a process with the given PID is still running
func IsProcessAlive(pid int) bool { _ = "STUB: not implemented"; return false }

// Send signal 0 to check if process exists
// This doesn't actually send a signal, just checks if we can send one

// Check the specific error

// No such process

// Permission denied, but process exists

// Other error, assume process exists to be safe

// Unknown error, assume process exists to be safe

// parseLockFile parses a lock file and returns the PID if valid
func parseLockFile(path string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:gosec // caller-supplied lock file path

// NewFileLock attempts to create a lock file at the specified path. Fails if another process has already created a
// lock file. Useful for situations where a process wants to hold a mutual exclusion lock on a resource.
// The caller is responsible for calling Release() to release the lock.
func NewFileLock(logger *slog.Logger, path string, fsync bool) (*FileLock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Try to create the lock file exclusively (O_EXCL ensures it fails if file exists)
//nolint:gosec // caller-supplied lock file path

// Lock file exists, check if it's stale

// Process is dead, remove stale lock file and try again

// Try to create the lock file again
//nolint:gosec // caller-supplied lock file path

// Process is still alive, cannot acquire lock

//nolint:gosec // caller-supplied lock file path

// Cannot parse lock file, treat as existing lock with debug info

//nolint:gosec // caller-supplied lock file path

// Write process ID and timestamp to the lock file for debugging

// Close and remove the file if we can't write to it

// Close and remove the file if we can't sync it

// Release releases the file lock by closing and removing the lock file.
// This is a no-op if the lock is already released.
func (fl *FileLock) Release() { _ = "STUB: not implemented"; return }

// Close the file first

// Remove the lock file

// Path returns the path of the lock file
func (fl *FileLock) Path() string {
	_ = "STUB: not implemented"

	// Create a lock on multiple directories. Returns a function that can be used to release all locks.
	return ""
}

func LockDirectories(
	logger *slog.Logger,
	directories []string,
	lockFileName string,
	fsync bool) (func(), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Release all previously acquired locks before returning an error

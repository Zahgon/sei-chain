/*
		Write-ahead log using files of bounded size for storage.
	  It appends entries to the <headPath> file until it reaches the limit size.
		Then it renames it to <headPath>.<sequential number> (a tail file) and creates new empty <headPath> file.
		It uses flock on an empty <HeadPath>.lock file to ensure exclusive access to the log files.

		Dir/
		- <HeadPath>.000   // First rolled file
		- <HeadPath>.001   // Second rolled file
		- ...
		- <HeadPath>       // Head file.
		- <HeadPath>.lock  // File used as a mutex.
*/
package wal

import (
	"errors"
	"os"

	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

const headerSize int64 = 8

const filePerms = os.FileMode(0600)

var ErrClosed error = errors.New("WAL closed")

type Config struct {
	FileSizeLimit  int64
	TotalSizeLimit int64
}

func DefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

// 10MB
// 1GB

func lockPath(headPath string) string { _ = "STUB: not implemented"; return "" }

func openLockFile(headPath string) (*os.File, error) { _ = "STUB: not implemented"; return nil, nil }

// logInner is a non-threadsafe inner implementation of Log.
// It is invalidated whenever ANY of the method call returns an error.
// Log protects access to the invalidated logInner to avoid misuse.
type logInner struct {
	cfg      *Config
	lockFile *os.File
	view     *logView
	writer   *logWriter
}

// ReadFile reads the whole log file at a given offset.
func (i *logInner) ReadFile(fileOffset int) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If head is requested, we need first sync it do disk.

func (i *logInner) Append(entry []byte) (err error) { _ = "STUB: not implemented"; return nil }

// Sync and close head.

// Move head to tail.

// Reopen head.

func (i *logInner) Size() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// Close releases all resources unconditionally.
// It invalidates logInner object.
func (i *logInner) Close() {
	_ = "STUB: not implemented"
	// Best effort syncing at close. No guarantees.
	return
}

// non-threadsafe WAL.
// Automatically closes the WAL if any operation returns an error.
// Locks the WAL files while opened.
type Log struct {
	inner utils.Option[*logInner]
}

func OpenLog(headPath string, cfg *Config) (*Log, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *Log) MinOffset() int { _ = "STUB: not implemented"; return 0 }

// ReadFile reads all entries from a file at a given offset.
// Available offsets are from range [MinOffset(),0]
func (l *Log) ReadFile(fileOffset int) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Append appends entry to the log atomically.
// You need to call Sync afterwards to ensure that the entry is persisted.
func (l *Log) Append(entry []byte) (err error) { _ = "STUB: not implemented"; return nil }

// Sync writes all buffered data to disk and calls fsync to ensure persistence.
func (l *Log) Sync() (err error) { _ = "STUB: not implemented"; return nil }

// Returns the total size of the log in bytes.
func (l *Log) Size() (res int64, err error) { _ = "STUB: not implemented"; return 0, nil }

// Close releases all resources unconditionally.
func (l *Log) Close() { _ = "STUB: not implemented"; return }

// Closes Log iff *err!=nil.
func (l *Log) closeOnErr(err *error) { _ = "STUB: not implemented"; return }

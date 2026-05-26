package wal

import (
	"bufio"
	"errors"
	"hash/crc32"
	"os"
)

var errEOF = errors.New("EOF")
var errCorrupted = errors.New("file corrupted")

var crc32c = crc32.MakeTable(crc32.Castagnoli)

type logReader struct {
	file      *os.File
	buf       *bufio.Reader
	bytesLeft int64
}

func openLogReader(path string) (*logReader, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *logReader) read(n int64) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *logReader) ReadEntry() (data []byte, err error) {
	_ = "STUB: not implemented"
	// Locking files on filesystem level is not really supported.
	// Therefore it is always possible that the file gets modified while we read it.
	// Hence we return a custom EOF error, so that it is distinguishable from EOF
	// returned by the file system.
	return nil, nil
}

// Close unconditionally releases all the resources.
func (r *logReader) Close() { _ = "STUB: not implemented"; return }

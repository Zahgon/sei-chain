package memiavl

import (
	"os"

	"github.com/ledgerwatch/erigon-lib/mmap"
)

// MmapFile manage the resources of a mmap-ed file
type MmapFile struct {
	file *os.File
	data []byte
	// mmap handle for windows (this is used to close mmap)
	handle *[mmap.MaxMapSize]byte
}

// NewMmap opens the file and creates a read-only memory mapping with MADV_RANDOM.
// MADV_RANDOM disables kernel readahead, which is optimal for the B+ tree random
// access patterns used during WAL replay and normal serving.
// Page cache warming is handled separately by prefetchSnapshot() using file I/O.
func NewMmap(path string) (*MmapFile, error) { _ = "STUB: not implemented"; return nil, nil }

// newMmapFile is the shared implementation that creates an mmap without any madvise hints.
func newMmapFile(path string) (*MmapFile, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *MmapFile) PrepareForRandomRead() { _ = "STUB: not implemented"; return }

// Switch to RANDOM access mode to disable readahead for random access patterns

// Close closes the file and mmap handles
func (m *MmapFile) Close() error { _ = "STUB: not implemented"; return nil }

// Data returns the mmap-ed buffer
func (m *MmapFile) Data() []byte { _ = "STUB: not implemented"; return nil }

func Mmap(f *os.File) ([]byte, *[mmap.MaxMapSize]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

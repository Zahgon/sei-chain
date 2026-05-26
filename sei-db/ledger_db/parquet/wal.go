package parquet

import (
	dbwal "github.com/sei-protocol/sei-chain/sei-db/wal"
)

// WALEntry represents a batch of receipts for a single block in the WAL.
type WALEntry struct {
	BlockNumber uint64
	Receipts    [][]byte
}

// encodeWALEntry encodes a WALEntry to binary format:
// [blockNumber:8][numReceipts:4][len1:4][receipt1]...[lenN:4][receiptN]
func encodeWALEntry(entry WALEntry) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// blockNumber + numReceipts

// length prefix + data

// decodeWALEntry decodes a binary WALEntry.
func decodeWALEntry(data []byte) (WALEntry, error) {
	_ = "STUB: not implemented"
	return *new(WALEntry), nil
}

func validateUint32Int(value int, field string) error { _ = "STUB: not implemented"; return nil }

func putUint32FromInt(dst []byte, value int) { _ = "STUB: not implemented"; return }

// NewWAL creates a new WAL for parquet receipts.
func NewWAL(dir string) (dbwal.GenericWAL[WALEntry], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Allow the WAL to be fully emptied after rotation/truncation.

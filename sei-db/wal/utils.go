package wal

import (
	"github.com/tidwall/wal"

	seidbproto "github.com/sei-protocol/sei-chain/sei-db/proto"
)

func LogPath(dir string) string { _ = "STUB: not implemented"; return "" }

// GetLastIndex returns the last written index of the replay log
func GetLastIndex(dir string) (index uint64, err error) { _ = "STUB: not implemented"; return 0, nil }

// truncateCorruptedTail truncates the corrupted tail
func truncateCorruptedTail(path string, format wal.LogFormat) error {
	_ = "STUB: not implemented"
	return nil
}

// loadNextJSONEntry loads json data like {"index":number,"data":string}
func loadNextJSONEntry(data []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:gosec

// loadNextBinaryEntry loads binary data like data_size + data
func loadNextBinaryEntry(data []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func channelBatchRecv[T any](ch <-chan T) []T {
	_ = "STUB: not implemented"
	// block if channel is empty
	return nil
}

// channel is closed

func MockKVPairs(kvPairs ...string) []*seidbproto.KVPair { _ = "STUB: not implemented"; return nil }

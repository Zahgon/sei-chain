//go:build rocksdbBackend
// +build rocksdbBackend

package mvcc

import (
	"github.com/linxGnu/grocksdb"
)

// CreateTSComparator should behavior identical with RocksDB builtin timestamp comparator.
// We also use the same builtin comparator name so the builtin tools `ldb`/`sst_dump`
// can work with the database.
func CreateTSComparator() *grocksdb.Comparator { _ = "STUB: not implemented"; return nil }

// compareTS compares timestamp as little endian encoded integers.
//
// NOTICE: The behavior must be identical to RocksDB builtin comparator
// "leveldb.BytewiseComparator.u64ts".
func compareTS(bz1 []byte, bz2 []byte) int { _ = "STUB: not implemented"; return 0 }

// compare compares two internal keys with timestamp suffix, larger timestamp
// comes first.
//
// NOTICE: The behavior must be identical to RocksDB builtin comparator
// "leveldb.BytewiseComparator.u64ts".
func compare(a []byte, b []byte) int { _ = "STUB: not implemented"; return 0 }

// Compare timestamp. For the same user key with different timestamps, larger
// (newer) timestamp comes first, which means seek operation will try to find
// a version less than or equal to the target version.

// compareWithoutTS compares two internal keys without the timestamp part.
//
// NOTICE: the behavior must be identical to RocksDB builtin comparator
// "leveldb.BytewiseComparator.u64ts".
func compareWithoutTS(a []byte, aHasTS bool, b []byte, bHasTS bool) int {
	_ = "STUB: not implemented"
	return 0
}

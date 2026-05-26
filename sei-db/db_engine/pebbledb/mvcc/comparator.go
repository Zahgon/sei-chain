package mvcc

import (
	"bytes"
	"fmt"

	"github.com/cockroachdb/pebble/v2"
)

// MVCCComparer returns a PebbleDB Comparer with encoding and decoding routines
// for MVCC control, used to compare and store versioned keys.
//
// Note: This Comparer implementation is largely based on PebbleDB's internal
// MVCC example, which can be found here:
// https://github.com/cockroachdb/pebble/blob/master/cmd/pebble/mvcc.go
var MVCCComparer = &pebble.Comparer{
	Name: "ss_pebbledb_comparator",

	Compare: MVCCKeyCompare,

	AbbreviatedKey: func(k []byte) uint64 {
		key, _, ok := SplitMVCCKey(k)
		if !ok {
			return 0
		}

		return pebble.DefaultComparer.AbbreviatedKey(key)
	},

	Equal: func(a, b []byte) bool {
		return MVCCKeyCompare(a, b) == 0
	},

	Separator: func(dst, a, b []byte) []byte {
		aKey, _, ok := SplitMVCCKey(a)
		if !ok {
			return append(dst, a...)
		}

		bKey, _, ok := SplitMVCCKey(b)
		if !ok {
			return append(dst, a...)
		}

		// if the keys are the same just return a
		if bytes.Equal(aKey, bKey) {
			return append(dst, a...)
		}

		n := len(dst)

		// MVCC key comparison uses bytes.Compare on the roachpb.Key, which is the
		// same semantics as pebble.DefaultComparer, so reuse the latter's Separator
		// implementation.
		dst = pebble.DefaultComparer.Separator(dst, aKey, bKey)

		// Did we pick a separator different than aKey? If we did not, we can't do
		// better than a.
		buf := dst[n:]
		if bytes.Equal(aKey, buf) {
			return append(dst[:n], a...)
		}

		// The separator is > aKey, so we only need to add the timestamp sentinel.
		return append(dst, 0)
	},

	ImmediateSuccessor: func(dst, a []byte) []byte {
		// The key `a` is guaranteed to be a bare prefix: It's a key without a version
		// — just a trailing 0-byte to signify the length of the version. For example
		// the user key "foo" is encoded as: "foo\0". We need to encode the immediate
		// successor to "foo", which in the natural byte ordering is "foo\0". Append
		// a single additional zero, to encode the user key "foo\0" with a zero-length
		// version.
		return append(append(dst, a...), 0)
	},

	Successor: func(dst, a []byte) []byte {
		aKey, _, ok := SplitMVCCKey(a)
		if !ok {
			return append(dst, a...)
		}

		n := len(dst)

		// MVCC key comparison uses bytes.Compare on the roachpb.Key, which is the
		// same semantics as pebble.DefaultComparer, so reuse the latter's Successor
		// implementation.
		dst = pebble.DefaultComparer.Successor(dst, aKey)

		// Did we pick a successor different than aKey? If we did not, we can't do
		// better than a.
		buf := dst[n:]
		if bytes.Equal(aKey, buf) {
			return append(dst[:n], a...)
		}

		// The successor is > aKey, so we only need to add the timestamp sentinel.
		return append(dst, 0)
	},

	FormatKey: func(k []byte) fmt.Formatter {
		return mvccKeyFormatter{key: k}
	},

	Split: func(k []byte) int {
		key, _, ok := SplitMVCCKey(k)
		if !ok {
			return len(k)
		}

		// This matches the behavior of libroach/KeyPrefix. RocksDB requires that
		// keys generated via a SliceTransform be comparable with normal encoded
		// MVCC keys. Encoded MVCC keys have a suffix indicating the number of
		// bytes of timestamp data. MVCC keys without a timestamp have a suffix of
		// 0. We're careful in EncodeKey to make sure that the user-key always has
		// a trailing 0. If there is no timestamp this falls out naturally. If
		// there is a timestamp we prepend a 0 to the encoded timestamp data.
		return len(key) + 1
	},

	ComparePointSuffixes: pebble.DefaultComparer.ComparePointSuffixes,

	CompareRangeSuffixes: pebble.DefaultComparer.CompareRangeSuffixes,
}

type mvccKeyFormatter struct {
	key []byte
}

func (f mvccKeyFormatter) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

// SplitMVCCKey accepts an MVCC key and returns the "user" key, the MVCC version,
// and a boolean indicating if the provided key is an MVCC key.
//
// Note, internally, we must make a copy of the provided mvccKey argument, which
// typically comes from the Key() method as it's not safe.
func SplitMVCCKey(mvccKey []byte) (key, version []byte, ok bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

// MVCCKeyCompare compares two MVCC keys.
func MVCCKeyCompare(a, b []byte) int { _ = "STUB: not implemented"; return 0 }

// This should never happen unless there is some sort of corruption of
// the keys. This is a little bizarre, but the behavior exactly matches
// engine/db.cc:DBComparator.

// Compute the index of the separator between the key and the timestamp.

// This should never happen unless there is some sort of corruption of
// the keys. This is a little bizarre, but the behavior exactly matches
// engine/db.cc:DBComparator.

// compare the "user key" part of the key

// compare the timestamp part of the key

// MVCCEncode dispatches between the descending and ascending encoders based on
// the mode flag. Descending-mode is used for fresh DBs created by this build;
// ascending-mode preserves compatibility with legacy DBs written by the
// previous ascending-version build.
func MVCCEncode(key []byte, version int64, descending bool) []byte {
	_ = "STUB: not implemented"
	return nil
}

// MVCCEncodeDescending encodes an MVCC key with the version encoded in
// descending byte order so newer versions sort before older ones for the same
// logical key.
//
// <key>\x00[<version>]<#version-bytes>
func MVCCEncodeDescending(key []byte, version int64) (dst []byte) {
	_ = "STUB: not implemented"
	return nil
}

// MVCCEncodeAscending encodes an MVCC key with the version encoded in
// ascending byte order. This matches the legacy on-disk format used by main.
//
// <key>\x00[<version>]<#version-bytes>
func MVCCEncodeAscending(key []byte, version int64) (dst []byte) {
	_ = "STUB: not implemented"
	return nil
}

// encodeUint64Descending encodes the uint64 value in descending order so newer
// versions sort before older versions for the same logical key.
func encodeUint64Descending(dst []byte, v uint64) []byte { _ = "STUB: not implemented"; return nil }

// decodeUint64Descending decodes a descending-encoded int64 from the input
// buffer and returns the original ascending version value.
func decodeUint64Descending(b []byte) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// encodeUint64Ascending encodes the uint64 value using a big-endian 8 byte
// representation. The bytes are appended to the supplied buffer and
// the final buffer is returned.
func encodeUint64Ascending(dst []byte, v uint64) []byte { _ = "STUB: not implemented"; return nil }

// decodeUint64Ascending decodes a int64 from the input buffer, treating
// the input as a big-endian 8 byte uint64 representation. The decoded int64 is
// returned.
func decodeUint64Ascending(b []byte) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

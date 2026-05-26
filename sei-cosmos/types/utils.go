package types

import (
	"time"

	dbm "github.com/tendermint/tm-db"
)

var (
	// This is set at compile time. Could be cleveldb, defaults is goleveldb.
	DBBackend = ""
	backend   = dbm.GoLevelDBBackend
)

func init() {
	if len(DBBackend) != 0 {
		backend = dbm.BackendType(DBBackend)
	}
}

// SortedJSON takes any JSON and returns it sorted by keys. Also, all white-spaces
// are removed.
// This method can be used to canonicalize JSON to be returned by GetSignBytes,
// e.g. for the ledger integration.
// If the passed JSON isn't valid it will return an error.
func SortJSON(toSortJSON []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MustSortJSON is like SortJSON but panic if an error occurs, e.g., if
// the passed JSON isn't valid.
func MustSortJSON(toSortJSON []byte) []byte { _ = "STUB: not implemented"; return nil }

// Uint64ToBigEndian - marshals uint64 to a bigendian byte slice so it can be sorted
func Uint64ToBigEndian(i uint64) []byte { _ = "STUB: not implemented"; return nil }

// BigEndianToUint64 returns an uint64 from big endian encoded bytes. If encoding
// is empty, zero is returned.
func BigEndianToUint64(bz []byte) uint64 { _ = "STUB: not implemented"; return 0 }

// Slight modification of the RFC3339Nano but it right pads all zeros and drops the time zone info
const SortableTimeFormat = "2006-01-02T15:04:05.000000000"

// Formats a time.Time into a []byte that can be sorted
func FormatTimeBytes(t time.Time) []byte { _ = "STUB: not implemented"; return nil }

// Parses a []byte encoded using FormatTimeKey back into a time.Time
func ParseTimeBytes(bz []byte) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// NewLevelDB instantiate a new LevelDB instance according to DBBackend.
func NewLevelDB(name, dir string) (db dbm.DB, err error) {
	_ = "STUB: not implemented"
	return *new(dbm.DB), nil
}

// copy bytes
func CopyBytes(bz []byte) (ret []byte) { _ = "STUB: not implemented"; return nil }

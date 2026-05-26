package disktable

import (
	"log/slog"
	"sync/atomic"
	"time"
)

const tableMetadataSerializationVersion = 0
const TableMetadataFileName = "table.metadata"

// tableMetadataSize is the on-disk byte size of the table metadata file:
//   - 4 bytes: serialization version
//   - 8 bytes: TTL (nanoseconds)
//   - 1 byte: sharding factor (capped at litt.MaxShardingFactor = 255)
const tableMetadataSize = 13

// tableMetadata contains table data that is preserved across restarts.
type tableMetadata struct {
	logger *slog.Logger

	tableDirectory string

	// the table's TTL, accessed/modified by concurrent goroutines
	ttl atomic.Pointer[time.Duration]

	// the table's sharding factor, accessed/modified by concurrent goroutines
	shardingFactor atomic.Uint32

	// If true, metadata writes will be atomic. Should be set to true in production, but can be set to false
	// to speed up unit tests.
	fsync bool
}

// newTableMetadata creates a new table metadata object.
func newTableMetadata(
	logger *slog.Logger,
	tableDirectory string,
	ttl time.Duration,
	shardingFactor uint8,
	fsync bool) (*tableMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// loadTableMetadata loads the table metadata from disk.
func loadTableMetadata(logger *slog.Logger, tableDirectory string) (*tableMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // path within table directory

// Size returns the size of the table metadata file in bytes.
func (t *tableMetadata) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// GetTTL returns the time-to-live for the table.
func (t *tableMetadata) GetTTL() time.Duration {
	_ = "STUB: not implemented"
	return *

	// SetTTL sets the time-to-live for the table.
	new(time.Duration)
}

func (t *tableMetadata) SetTTL(ttl time.Duration) error { _ = "STUB: not implemented"; return nil }

// GetShardingFactor returns the sharding factor for the table. Capped at litt.MaxShardingFactor (255) so the value
func (t *tableMetadata) GetShardingFactor() uint8 { _ = "STUB: not implemented"; return 0 }

//nolint:gosec // bounded to uint8 by SetShardingFactor / deserialize

// SetShardingFactor sets the sharding factor for the table.
func (t *tableMetadata) SetShardingFactor(shardingFactor uint8) error {
	_ = "STUB: not implemented"
	return nil
}

// Store atomically stores the table metadata to disk.
func (t *tableMetadata) write() error { _ = "STUB: not implemented"; return nil }

// serialize serializes the table metadata to a byte slice.
func (t *tableMetadata) serialize() []byte { _ = "STUB: not implemented"; return nil }

// Write the version.

// Write the TTL.

//nolint:gosec // serialized as time.Duration

// Write the sharding factor. Storing this in a single byte makes it structurally impossible for the on-disk
// shard count to exceed litt.MaxShardingFactor (255).

// deserialize deserializes the table metadata from a byte slice.
func deserialize(data []byte) (*tableMetadata, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:gosec // serialized as time.Duration

// delete deletes the table metadata from disk.
func (t *tableMetadata) delete() error { _ = "STUB: not implemented"; return nil }

// path returns the path to the table metadata file.
func metadataPath(tableDirectory string) string { _ = "STUB: not implemented"; return "" }

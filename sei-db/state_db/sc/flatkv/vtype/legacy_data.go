package vtype

type LegacyDataVersion uint8

// DO NOT CHANGE VERSION VALUES!!! Adding new versions is ok, but historical versions should never be removed/changed.
const (
	LegacyDataVersion0 LegacyDataVersion = 0
)

/*
Serialization schema for LegacyData version 0:

| Version | Block Height | Value    |
|---------|--------------|----------|
| 1 byte  | 8 bytes      | variable |

Data is stored in big-endian order. Value is variable length.
*/

const (
	legacyVersionStart     = 0
	legacyBlockHeightStart = legacyVersionStart + VersionLength
	legacyValueStart       = legacyBlockHeightStart + BlockHeightLength
	legacyHeaderLength     = VersionLength + BlockHeightLength
)

var _ VType = (*LegacyData)(nil)

// Used for encapsulating and serializing legacy data in the FlatKV legacy database.
//
// This data structure is not threadsafe. Values passed into and values received from this data structure
// are not safe to modify without first copying them.
type LegacyData struct {
	version     LegacyDataVersion
	blockHeight int64
	value       []byte
	isDelete    bool
}

// Create a new LegacyData with the given value.
func NewLegacyData() *LegacyData { _ = "STUB: not implemented"; return nil }

// Serialize the legacy data to a byte slice.
func (l *LegacyData) Serialize() []byte { _ = "STUB: not implemented"; return nil }

//nolint:gosec

// Deserialize the legacy data from the given byte slice.
func DeserializeLegacyData(data []byte) (*LegacyData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec

// Get the serialization version for this LegacyData instance.
func (l *LegacyData) GetSerializationVersion() LegacyDataVersion {
	_ = "STUB: not implemented"
	return *new(LegacyDataVersion)
}

// Get the block height when this legacy entry was last modified.
func (l *LegacyData) GetBlockHeight() int64 { _ = "STUB: not implemented"; return 0 }

// Get the legacy value.
func (l *LegacyData) GetValue() []byte { _ = "STUB: not implemented"; return nil }

// Set the block height when this legacy entry was last modified/touched. Returns self (or a new LegacyData if nil).
func (l *LegacyData) SetBlockHeight(blockHeight int64) *LegacyData {
	_ = "STUB: not implemented"
	return nil
}

// Set the legacy value. Returns self (or a new LegacyData if nil).
// Clears the delete flag — an explicit SetValue is a write, not a deletion,
// even when value is empty ([]byte{} is a valid Cosmos module value).
func (l *LegacyData) SetValue(value []byte) *LegacyData { _ = "STUB: not implemented"; return nil }

// MarkDeleted flags this entry for physical key removal at commit time.
// The stored value is irrelevant once marked; IsDelete() will return true.
func (l *LegacyData) MarkDeleted() *LegacyData { _ = "STUB: not implemented"; return nil }

// IsDelete reports whether this entry represents a deletion.
// Uses an explicit flag rather than value-length inference so that empty
// values ([]byte{}) written by Cosmos modules are not misinterpreted as
// deletions.
func (l *LegacyData) IsDelete() bool { _ = "STUB: not implemented"; return false }

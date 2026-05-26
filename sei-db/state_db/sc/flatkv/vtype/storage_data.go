package vtype

type StorageDataVersion uint8

// DO NOT CHANGE VERSION VALUES!!! Adding new versions is ok, but historical versions should never be removed/changed.
const (
	StorageDataVersion0 StorageDataVersion = 0
)

/*
Serialization schema for StorageData version 0:

| Version | Block Height | Value    |
|---------|--------------|----------|
| 1 byte  | 8 bytes      | 32 bytes |

Data is stored in big-endian order.
*/

const (
	storageVersionStart     = 0
	storageBlockHeightStart = storageVersionStart + VersionLength
	storageValueStart       = storageBlockHeightStart + BlockHeightLength

	storageDataLength = VersionLength + BlockHeightLength + StorageValueLength
)

var _ VType = (*StorageData)(nil)

// Used for encapsulating and serializing storage slot data in the FlatKV storage database.
//
// This data structure is not threadsafe. Values passed into and values received from this data structure
// are not safe to modify without first copying them.
type StorageData struct {
	data []byte
}

// Create a new StorageData initialized to all 0s.
func NewStorageData() *StorageData { _ = "STUB: not implemented"; return nil }

// Serialize the storage data to a byte slice.
//
// The returned byte slice is not safe to modify without first copying it.
func (s *StorageData) Serialize() []byte { _ = "STUB: not implemented"; return nil }

// Deserialize the storage data from the given byte slice.
func DeserializeStorageData(data []byte) (*StorageData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get the serialization version for this StorageData instance.
func (s *StorageData) GetSerializationVersion() StorageDataVersion {
	_ = "STUB: not implemented"
	return *new(StorageDataVersion)
}

// Get the block height when this storage slot was last modified.
func (s *StorageData) GetBlockHeight() int64 { _ = "STUB: not implemented"; return 0 }

//nolint:gosec // block height is always within int64 range

// Get the storage slot value.
func (s *StorageData) GetValue() *[32]byte { _ = "STUB: not implemented"; return nil }

// Check if this storage data signifies a deletion operation. A deletion operation is automatically
// performed when the value is all 0s (with the exception of the serialization version and block height).
func (s *StorageData) IsDelete() bool { _ = "STUB: not implemented"; return false }

// Set the block height when this storage slot was last modified/touched. Returns self (or a new StorageData if nil).
func (s *StorageData) SetBlockHeight(blockHeight int64) *StorageData {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // block height is always non-negative

// Set the storage slot value. Returns self (or a new StorageData if nil).
func (s *StorageData) SetValue(value *[32]byte) *StorageData { _ = "STUB: not implemented"; return nil }

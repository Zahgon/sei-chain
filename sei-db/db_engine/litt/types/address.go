package types

// AddressSerializedSize is the on-disk size of a serialized Address in bytes.
// Layout: index(4) | offset(4) | shardID(1) | valueSize(4)
const AddressSerializedSize = 13

// Address describes the location of a value on disk.
//
// An Address identifies the file the value lives in (Index), the byte offset of the value's length prefix
// within that file (Offset), the shard within the segment that owns the value (ShardID), and the size of
// the value itself in bytes (ValueSize).
type Address struct {
	// index is the segment index that owns the value. Combined with the shardID, it identifies the value file
	// that contains the value's bytes.
	index uint32
	// offset is the byte position of the value's length prefix within the shard's value file. The value's
	// bytes immediately follow the 4-byte length prefix.
	offset uint32
	// shardID is the index of the shard within the segment that holds the value. Encoded as a single byte,
	// which caps the maximum sharding factor at 255 (see litt.MaxShardingFactor).
	shardID uint8
	// valueSize is the length of the value in bytes (not counting the 4-byte length prefix on disk).
	valueSize uint32
}

// NewAddress creates a new Address.
func NewAddress(index uint32, offset uint32, shardID uint8, valueSize uint32) Address {
	_ = "STUB: not implemented"
	return *new(Address)
}

// DeserializeAddress converts a byte slice to an Address. The slice must be exactly AddressSerializedSize bytes.
func DeserializeAddress(bytes []byte) (Address, error) {
	_ = "STUB: not implemented"
	return *new(Address), nil
}

// Index returns the segment index of the value.
func (a Address) Index() uint32 {
	_ = "STUB: not implemented"

	// Offset returns the byte offset of the value within its shard's value file.
	return 0
}

func (a Address) Offset() uint32 {
	_ = "STUB: not implemented"

	// ShardID returns the shard within the segment that owns the value.
	return 0
}

func (a Address) ShardID() uint8 {
	_ = "STUB: not implemented"

	// ValueSize returns the size of the value in bytes.
	return 0
}

func (a Address) ValueSize() uint32 {
	_ = "STUB: not implemented"

	// String returns a string representation of the address.
	return 0
}

func (a Address) String() string { _ = "STUB: not implemented"; return "" }

// Serialize converts the address to a byte slice of length AddressSerializedSize.
func (a Address) Serialize() []byte { _ = "STUB: not implemented"; return nil }

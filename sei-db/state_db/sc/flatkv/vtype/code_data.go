package vtype

type CodeDataVersion uint8

// DO NOT CHANGE VERSION VALUES!!! Adding new versions is ok, but historical versions should never be removed/changed.
const (
	CodeDataVersion0 CodeDataVersion = 0
)

/*
Serialization schema for CodeData version 0:

| Version | Block Height | Bytecode     |
|---------|--------------|--------------|
| 1 byte  | 8 bytes      | variable     |

Data is stored in big-endian order. Bytecode is variable length.
*/

const (
	codeVersionStart     = 0
	codeBlockHeightStart = codeVersionStart + VersionLength

	codeBytecodeStart = codeBlockHeightStart + BlockHeightLength
)

var _ VType = (*CodeData)(nil)

// Used for encapsulating and serializing contract bytecode in the FlatKV code database.
//
// This data structure is not threadsafe. Values passed into and values received from this data structure
// are not safe to modify without first copying them.
type CodeData struct {
	version     CodeDataVersion
	blockHeight int64
	bytecode    []byte
}

// Create a new CodeData with the given bytecode.
func NewCodeData() *CodeData { _ = "STUB: not implemented"; return nil }

// Serialize the code data to a byte slice.
func (c *CodeData) Serialize() []byte { _ = "STUB: not implemented"; return nil }

//nolint:gosec

// Deserialize the code data from the given byte slice.
func DeserializeCodeData(data []byte) (*CodeData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec

// Get the serialization version for this CodeData instance.
func (c *CodeData) GetSerializationVersion() CodeDataVersion {
	_ = "STUB: not implemented"
	return *new(CodeDataVersion)
}

// Get the block height when this code was last modified.
func (c *CodeData) GetBlockHeight() int64 { _ = "STUB: not implemented"; return 0 }

// Get the contract bytecode.
func (c *CodeData) GetBytecode() []byte { _ = "STUB: not implemented"; return nil }

// Set the contract bytecode. Returns self (or a new CodeData if nil).
func (c *CodeData) SetBytecode(bytecode []byte) *CodeData { _ = "STUB: not implemented"; return nil }

// Check if this code data signifies a deletion operation. A deletion operation is automatically
// performed when the bytecode is empty (with the exception of the serialization version and block height).
func (c *CodeData) IsDelete() bool { _ = "STUB: not implemented"; return false }

// Set the block height when this code was last modified/touched. Returns self (or a new CodeData if nil).
func (c *CodeData) SetBlockHeight(blockHeight int64) *CodeData {
	_ = "STUB: not implemented"
	return nil
}

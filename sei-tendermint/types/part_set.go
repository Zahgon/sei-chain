package types

import (
	"bytes"
	"errors"
	"io"
	"sync"

	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto/merkle"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/bits"
	tmbytes "github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	"github.com/sei-protocol/seilog"
)

var (
	logger = seilog.NewLogger("tendermint", "types")

	ErrPartSetUnexpectedIndex = errors.New("error part set unexpected index")
	ErrPartSetInvalidProof    = errors.New("error part set invalid proof")
)

type Part struct {
	Index uint32           `json:"index"`
	Bytes tmbytes.HexBytes `json:"bytes"`
	Proof merkle.Proof     `json:"proof"`
}

// ValidateBasic performs basic validation.
func (part *Part) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String returns a string representation of Part.
//
// See StringIndented.
func (part *Part) String() string { _ = "STUB: not implemented"; return "" }

// StringIndented returns an indented Part.
//
// See merkle.Proof#StringIndented
func (part *Part) StringIndented(indent string) string { _ = "STUB: not implemented"; return "" }

func (part *Part) ToProto() *tmproto.Part { _ = "STUB: not implemented"; return nil }

func PartFromProto(pb *tmproto.Part) (*Part, error) { _ = "STUB: not implemented"; return nil, nil }

//-------------------------------------

type PartSetHeader struct {
	Total uint32           `json:"total"` // BlockPartsCount
	Hash  tmbytes.HexBytes `json:"hash"`
}

// String returns a string representation of PartSetHeader.
//
// 1. total number of parts
// 2. first 6 bytes of the hash
func (psh PartSetHeader) String() string { _ = "STUB: not implemented"; return "" }

func (psh PartSetHeader) IsZero() bool { _ = "STUB: not implemented"; return false }

func (psh PartSetHeader) Equals(other PartSetHeader) bool { _ = "STUB: not implemented"; return false }

// ValidateBasic performs basic validation.
func (psh PartSetHeader) ValidateBasic() error {
	_ = "STUB: not implemented"
	// Hash can be empty in case of POLBlockID.PartSetHeader in Proposal.
	return nil
}

// Check memory limits before acquiring lock or setting any state

// ToProto converts PartSetHeader to protobuf
func (psh *PartSetHeader) ToProto() tmproto.PartSetHeader {
	_ = "STUB: not implemented"
	return *new(tmproto.PartSetHeader)
}

// FromProto sets a protobuf PartSetHeader to the given pointer
func PartSetHeaderFromProto(ppsh *tmproto.PartSetHeader) (*PartSetHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProtoPartSetHeaderIsZero is similar to the IsZero function for
// PartSetHeader, but for the Protobuf representation.
func ProtoPartSetHeaderIsZero(ppsh *tmproto.PartSetHeader) bool {
	_ = "STUB: not implemented"
	return false
}

//-------------------------------------

type PartSet struct {
	total uint32
	hash  []byte

	mtx           sync.Mutex
	parts         []*Part
	partsBitArray *bits.BitArray
	count         uint32
	// a count of the total size (in bytes). Used to ensure that the
	// part set doesn't exceed the maximum block bytes
	byteSize int64
}

// Returns an immutable, full PartSet from the data bytes.
// The data bytes are split into "partSize" chunks, and merkle tree computed.
// CONTRACT: partSize is greater than zero.
func NewPartSetFromData(data []byte, partSize uint32) *PartSet {
	_ = "STUB: not implemented"
	// divide data into 4kb parts.
	return nil
}

//nolint:gosec // data length is bounded by block size limits; no overflow risk

//nolint:gosec // total fits in int since it's derived from block-bounded data

//nolint:gosec // partSize is small (4KB); product fits in uint32

//nolint:gosec // i < total which fits in int

// Compute merkle proofs

// Returns an empty PartSet ready to be populated.
func NewPartSetFromHeader(header PartSetHeader) *PartSet { _ = "STUB: not implemented"; return nil }

// Minimal safe size
// Keep original hash for compatibility

func (ps *PartSet) Header() PartSetHeader { _ = "STUB: not implemented"; return *new(PartSetHeader) }

func (ps *PartSet) HasHeader(header PartSetHeader) bool { _ = "STUB: not implemented"; return false }

func (ps *PartSet) BitArray() *bits.BitArray { _ = "STUB: not implemented"; return nil }

func (ps *PartSet) Hash() []byte { _ = "STUB: not implemented"; return nil }

func (ps *PartSet) HashesTo(hash []byte) bool { _ = "STUB: not implemented"; return false }

func (ps *PartSet) Count() uint32 { _ = "STUB: not implemented"; return 0 }

func (ps *PartSet) ByteSize() int64 { _ = "STUB: not implemented"; return 0 }

func (ps *PartSet) Total() uint32 { _ = "STUB: not implemented"; return 0 }

func (ps *PartSet) AddPart(part *Part) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Invalid part index

// If part already exists, return false.

// Check hash proof

// Add part

func (ps *PartSet) GetPart(index int) *Part { _ = "STUB: not implemented"; return nil }

func (ps *PartSet) IsComplete() bool { _ = "STUB: not implemented"; return false }

func (ps *PartSet) GetReader() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

type PartSetReader struct {
	i      int
	parts  []*Part
	reader *bytes.Reader
}

func NewPartSetReader(parts []*Part) *PartSetReader { _ = "STUB: not implemented"; return nil }

func (psr *PartSetReader) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// StringShort returns a short version of String.
//
// (Count of Total)
func (ps *PartSet) StringShort() string { _ = "STUB: not implemented"; return "" }

func (ps *PartSet) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

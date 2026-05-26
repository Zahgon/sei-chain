package bits

import (
	"regexp"
	"sync"

	tmprotobits "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/libs/bits"
)

// BitArray is a thread-safe implementation of a bit array.
type BitArray struct {
	mtx   sync.Mutex
	Bits  int      `json:"bits"`  // NOTE: persisted via reflect, must be exported
	Elems []uint64 `json:"elems"` // NOTE: persisted via reflect, must be exported
}

// NewBitArray returns a new bit array.
// It returns nil if the number of bits is zero.
func NewBitArray(bits int) *BitArray { _ = "STUB: not implemented"; return nil }

// reset changes size of BitArray to `bits` and re-allocates (zeroed) data buffer
func (bA *BitArray) reset(bits int) { _ = "STUB: not implemented"; return }

// Size returns the number of bits in the bitarray
func (bA *BitArray) Size() int { _ = "STUB: not implemented"; return 0 }

// GetIndex returns the bit at index i within the bit array.
// The behavior is undefined if i >= bA.Bits
func (bA *BitArray) GetIndex(i int) bool { _ = "STUB: not implemented"; return false }

func (bA *BitArray) getIndex(i int) bool { _ = "STUB: not implemented"; return false }

//nolint:gosec // i is bounds-checked above; i%64 is always in [0, 63]

// SetIndex sets the bit at index i within the bit array.
// This method returns false if i is out of range of the BitArray.
func (bA *BitArray) SetIndex(i int, v bool) bool { _ = "STUB: not implemented"; return false }

func (bA *BitArray) setIndex(i int, v bool) bool { _ = "STUB: not implemented"; return false }

//nolint:gosec // i is bounds-checked above; i%64 is always in [0, 63]

//nolint:gosec // i is bounds-checked above; i%64 is always in [0, 63]

// Copy returns a copy of the provided bit array.
func (bA *BitArray) Copy() *BitArray { _ = "STUB: not implemented"; return nil }

func (bA *BitArray) copy() *BitArray { _ = "STUB: not implemented"; return nil }

func (bA *BitArray) copyBits(bits int) *BitArray { _ = "STUB: not implemented"; return nil }

// Or returns a bit array resulting from a bitwise OR of the two bit arrays.
// If the two bit-arrys have different lengths, Or right-pads the smaller of the two bit-arrays with zeroes.
// Thus the size of the return value is the maximum of the two provided bit arrays.
func (bA *BitArray) Or(o *BitArray) *BitArray { _ = "STUB: not implemented"; return nil }

// And returns a bit array resulting from a bitwise AND of the two bit arrays.
// If the two bit-arrys have different lengths, this truncates the larger of the two bit-arrays from the right.
// Thus the size of the return value is the minimum of the two provided bit arrays.
func (bA *BitArray) And(o *BitArray) *BitArray { _ = "STUB: not implemented"; return nil }

func (bA *BitArray) and(o *BitArray) *BitArray { _ = "STUB: not implemented"; return nil }

// Not returns a bit array resulting from a bitwise Not of the provided bit array.
func (bA *BitArray) Not() *BitArray { _ = "STUB: not implemented"; return nil }

// Degenerate

func (bA *BitArray) not() *BitArray { _ = "STUB: not implemented"; return nil }

// Sub subtracts the two bit-arrays bitwise, without carrying the bits.
// Note that carryless subtraction of a - b is (a and not b).
// The output is the same as bA, regardless of o's size.
// If bA is longer than o, o is right padded with zeroes
func (bA *BitArray) Sub(o *BitArray) *BitArray { _ = "STUB: not implemented"; return nil }

// TODO: Decide if we should do 1's complement here?

// output is the same size as bA

// Only iterate to the minimum size between the two.
// If o is longer, those bits are ignored.
// If bA is longer, then skipping those iterations is equivalent
// to right padding with 0's

// &^ is and not in golang

// PickRandom returns a random index for a set bit in the bit array.
// If there is no such value, it returns 0, false.
// It uses math/rand's global randomness Source to get this index.
func (bA *BitArray) PickRandom() (int, bool) { _ = "STUB: not implemented"; return 0, false }

// no bits set to true

// NOTE: using the default math/rand might result in somewhat
// amount of determinism here. It would be possible to use
// rand.New(rand.NewSeed(time.Now().Unix())).Intn() to
// counteract this possibility if it proved to be material.
//
// nolint:gosec // G404: Use of weak random number generator

func (bA *BitArray) getTrueIndices() []int { _ = "STUB: not implemented"; return nil }

// set all true indices

//nolint:gosec // j is in [0, 63]; always safe for uint64

// handle last element

//nolint:gosec // i is in [0, 63]; always safe for uint64

// String returns a string representation of BitArray: BA{<bit-string>},
// where <bit-string> is a sequence of 'x' (1) and '_' (0).
// The <bit-string> includes spaces and newlines to help people.
// For a simple sequence of 'x' and '_' characters with no spaces or newlines,
// see the MarshalJSON() method.
// Example: "BA{_x_}" or "nil-BitArray" for nil.
func (bA *BitArray) String() string { _ = "STUB: not implemented"; return "" }

// StringIndented returns the same thing as String(), but applies the indent
// at every 10th bit, and twice at every 50th bit.
func (bA *BitArray) StringIndented(indent string) string { _ = "STUB: not implemented"; return "" }

func (bA *BitArray) stringIndented(indent string) string { _ = "STUB: not implemented"; return "" }

// Update sets the bA's bits to be that of the other bit array.
// The copying begins from the begin of both bit arrays.
func (bA *BitArray) Update(o *BitArray) { _ = "STUB: not implemented"; return }

// MarshalJSON implements json.Marshaler interface by marshaling bit array
// using a custom format: a string of '-' or 'x' where 'x' denotes the 1 bit.
func (bA *BitArray) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

var bitArrayJSONRegexp = regexp.MustCompile(`\A"([_x]*)"\z`)

// UnmarshalJSON implements json.Unmarshaler interface by unmarshaling a custom
// JSON description.
func (bA *BitArray) UnmarshalJSON(bz []byte) error { _ = "STUB: not implemented"; return nil }

// This is required e.g. for encoding/json when decoding
// into a pointer with pre-allocated BitArray.

// Validate 'b'.

// ToProto converts BitArray to protobuf. It returns nil if BitArray is
// nil/empty.
func (bA *BitArray) ToProto() *tmprotobits.BitArray { _ = "STUB: not implemented"; return nil }

// empty

// FromProto sets BitArray to the given protoBitArray. It returns an error if
// protoBitArray is invalid.
func (bA *BitArray) FromProto(protoBitArray *tmprotobits.BitArray) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate protoBitArray.

// #[32bit]
// prevent overflow on 32bit systems

func numElems(bits int) int { _ = "STUB: not implemented"; return 0 }

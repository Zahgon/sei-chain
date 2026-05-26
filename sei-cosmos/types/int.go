package types

import (
	"encoding"
	"math/big"
	"testing"
)

const maxBitLen = 256

func newIntegerFromString(s string) (*big.Int, bool) { _ = "STUB: not implemented"; return nil, false }

func equal(i *big.Int, i2 *big.Int) bool { _ = "STUB: not implemented"; return false }

func gt(i *big.Int, i2 *big.Int) bool { _ = "STUB: not implemented"; return false }

func gte(i *big.Int, i2 *big.Int) bool { _ = "STUB: not implemented"; return false }

func lt(i *big.Int, i2 *big.Int) bool { _ = "STUB: not implemented"; return false }

func lte(i *big.Int, i2 *big.Int) bool { _ = "STUB: not implemented"; return false }

func add(i *big.Int, i2 *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

func sub(i *big.Int, i2 *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

func mul(i *big.Int, i2 *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

func div(i *big.Int, i2 *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

func mod(i *big.Int, i2 *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

func neg(i *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

func abs(i *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

func min(i *big.Int, i2 *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

func max(i *big.Int, i2 *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

func unmarshalText(i *big.Int, text string) error { _ = "STUB: not implemented"; return nil }

var _ CustomProtobufType = (*Int)(nil)

// Int wraps big.Int with a 257 bit range bound
// Checks overflow, underflow and division by zero
// Exists in range from -(2^256 - 1) to 2^256 - 1
type Int struct {
	i *big.Int
}

// BigInt converts Int to big.Int
func (i Int) BigInt() *big.Int { _ = "STUB: not implemented"; return nil }

// IsNil returns true if Int is uninitialized
func (i Int) IsNil() bool {
	_ = "STUB: not implemented"

	// NewInt constructs Int from int64
	return false
}

func NewInt(n int64) Int {
	_ = "STUB: not implemented"
	return *

	// NewIntFromUint64 constructs an Int from a uint64.
	new(Int)
}

func NewIntFromUint64(n uint64) Int { _ = "STUB: not implemented"; return *new(Int) }

// NewIntFromBigInt constructs Int from big.Int. If the provided big.Int is nil,
// it returns an empty instance. This function panics if the bit length is > 256.
func NewIntFromBigInt(i *big.Int) Int { _ = "STUB: not implemented"; return *new(Int) }

// NewIntFromString constructs Int from string
func NewIntFromString(s string) (res Int, ok bool) {
	_ = "STUB: not implemented"
	return *new(Int), false
}

// Check overflow

// NewIntWithDecimal constructs Int with decimal
// Result value is n*10^dec
func NewIntWithDecimal(n int64, dec int) Int { _ = "STUB: not implemented"; return *new(Int) }

// Check overflow

// ZeroInt returns Int value with zero
func ZeroInt() Int {
	_ = "STUB: not implemented"
	return *

	// OneInt returns Int value with one
	new(Int)
}

func OneInt() Int {
	_ = "STUB: not implemented"
	return *

	// ToDec converts Int to Dec
	new(Int)
}

func (i Int) ToDec() Dec {
	_ = "STUB: not implemented"
	return *

	// Int64 converts Int to int64
	// Panics if the value is out of range
	new(Dec)
}

func (i Int) Int64() int64 { _ = "STUB: not implemented"; return 0 }

// IsInt64 returns true if Int64() not panics
func (i Int) IsInt64() bool { _ = "STUB: not implemented"; return false }

// Uint64 converts Int to uint64
// Panics if the value is out of range
func (i Int) Uint64() uint64 { _ = "STUB: not implemented"; return 0 }

// IsUint64 returns true if Uint64() not panics
func (i Int) IsUint64() bool { _ = "STUB: not implemented"; return false }

// IsZero returns true if Int is zero
func (i Int) IsZero() bool { _ = "STUB: not implemented"; return false }

// IsNegative returns true if Int is negative
func (i Int) IsNegative() bool { _ = "STUB: not implemented"; return false }

// IsPositive returns true if Int is positive
func (i Int) IsPositive() bool { _ = "STUB: not implemented"; return false }

// Sign returns sign of Int
func (i Int) Sign() int {
	_ = "STUB: not implemented"

	// Equal compares two Ints
	return 0
}

func (i Int) Equal(i2 Int) bool { _ = "STUB: not implemented"; return false }

// GT returns true if first Int is greater than second
func (i Int) GT(i2 Int) bool { _ = "STUB: not implemented"; return false }

// GTE returns true if receiver Int is greater than or equal to the parameter
// Int.
func (i Int) GTE(i2 Int) bool { _ = "STUB: not implemented"; return false }

// LT returns true if first Int is lesser than second
func (i Int) LT(i2 Int) bool { _ = "STUB: not implemented"; return false }

// LTE returns true if first Int is less than or equal to second
func (i Int) LTE(i2 Int) bool { _ = "STUB: not implemented"; return false }

// Add adds Int from another
func (i Int) Add(i2 Int) (res Int) {
	_ = "STUB: not implemented"
	return *

	// Check overflow
	new(Int)
}

// AddRaw adds int64 to Int
func (i Int) AddRaw(i2 int64) Int {
	_ = "STUB: not implemented"
	return *

	// Sub subtracts Int from another
	new(Int)
}

func (i Int) Sub(i2 Int) (res Int) {
	_ = "STUB: not implemented"
	return *

	// Check overflow
	new(Int)
}

// SubRaw subtracts int64 from Int
func (i Int) SubRaw(i2 int64) Int {
	_ = "STUB: not implemented"
	return *

	// Mul multiples two Ints
	new(Int)
}

func (i Int) Mul(i2 Int) (res Int) {
	_ = "STUB: not implemented"
	// Check overflow
	return *new(Int)
}

// Check overflow if sign of both are same

// MulRaw multiplies Int and int64
func (i Int) MulRaw(i2 int64) Int {
	_ = "STUB: not implemented"
	return *

	// Quo divides Int with Int
	new(Int)
}

func (i Int) Quo(i2 Int) (res Int) {
	_ = "STUB: not implemented"
	// Check division-by-zero
	return *new(Int)
}

// QuoRaw returns the quotient of Int division by int64.
func (i Int) QuoRaw(i2 int64) Int {
	_ = "STUB: not implemented"
	return *

	// Mod returns remainder after dividing with Int
	new(Int)
}

func (i Int) Mod(i2 Int) Int { _ = "STUB: not implemented"; return *new(Int) }

// ModRaw returns remainder after dividing with int64
func (i Int) ModRaw(i2 int64) Int {
	_ = "STUB: not implemented"
	return *

	// Neg negates Int
	new(Int)
}

func (i Int) Neg() (res Int) {
	_ = "STUB: not implemented"
	return *

	// Abs returns the absolute value of Int.
	new(Int)
}

func (i Int) Abs() Int {
	_ = "STUB: not implemented"
	return *

	// return the minimum of the ints
	new(Int)
}

func MinInt(i1, i2 Int) Int { _ = "STUB: not implemented"; return *new(Int) }

// MaxInt returns the maximum between two integers.
func MaxInt(i, i2 Int) Int { _ = "STUB: not implemented"; return *new(Int) }

// String returns the string representation of the Int.
func (i Int) String() string { _ = "STUB: not implemented"; return "" }

// MarshalJSON defines custom encoding scheme
func (i Int) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	// Necessary since default Uint initialization has i.i as nil
	return nil, nil
}

// UnmarshalJSON defines custom decoding scheme
func (i *Int) UnmarshalJSON(bz []byte) error {
	_ = "STUB: not implemented"
	// Necessary since default Int initialization has i.i as nil
	return nil
}

// MarshalJSON for custom encoding scheme
// Must be encoded as a string for JSON precision
func marshalJSON(i encoding.TextMarshaler) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalJSON for custom decoding scheme
// Must be encoded as a string for JSON precision
func unmarshalJSON(i *big.Int, bz []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalYAML returns the YAML representation.
func (i Int) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil,

		// Marshal implements the gogo proto custom type interface.
		nil
}

func (i Int) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalTo implements the gogo proto custom type interface.
func (i *Int) MarshalTo(data []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// The value 0

// Unmarshal implements the gogo proto custom type interface.
func (i *Int) Unmarshal(data []byte) error {
	_ = "STUB: not implemented"
	// maxBitLen is 256, which requires ~78 decimal digits, so 100 is a safe upper
	// bound (with room for sign).
	return nil
}

// Size implements the gogo proto custom type interface.
func (i *Int) Size() int { _ = "STUB: not implemented"; return 0 }

// Override Amino binary serialization by proxying to protobuf.
func (i Int) MarshalAmino() ([]byte, error)   { _ = "STUB: not implemented"; return nil, nil }
func (i *Int) UnmarshalAmino(bz []byte) error { _ = "STUB: not implemented"; return nil }

// intended to be used with require/assert:  require.True(IntEq(...))
func IntEq(t *testing.T, exp, got Int) (*testing.T, bool, string, string, string) {
	_ = "STUB: not implemented"
	return nil, false, "", "", ""
}

func (ip IntProto) String() string { _ = "STUB: not implemented"; return "" }

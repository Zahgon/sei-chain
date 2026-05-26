package types

import (
	"math/big"
)

// Uint wraps integer with 256 bit range bound
// Checks overflow, underflow and division by zero
// Exists in range from 0 to 2^256-1
type Uint struct {
	i *big.Int
}

// BigInt converts Uint to big.Int
func (u Uint) BigInt() *big.Int { _ = "STUB: not implemented"; return nil }

// NewUintFromBigUint constructs Uint from big.Uint
func NewUintFromBigInt(i *big.Int) Uint { _ = "STUB: not implemented"; return *new(Uint) }

// NewUint constructs Uint from uint64.
func NewUint(n uint64) Uint { _ = "STUB: not implemented"; return *new(Uint) }

// NewUintFromString constructs Uint from string
func NewUintFromString(s string) Uint { _ = "STUB: not implemented"; return *new(Uint) }

// ZeroUint returns unsigned zero.
func ZeroUint() Uint {
	_ = "STUB: not implemented"
	return *

	// OneUint returns Uint value with one.
	new(Uint)
}

func OneUint() Uint { _ = "STUB: not implemented"; return *new(Uint) }

var _ CustomProtobufType = (*Uint)(nil)

// Uint64 converts Uint to uint64
// Panics if the value is out of range
func (u Uint) Uint64() uint64 { _ = "STUB: not implemented"; return 0 }

// IsZero returns 1 if the uint equals to 0.
func (u Uint) IsZero() bool { _ = "STUB: not implemented"; return false }

// Equal compares two Uints
func (u Uint) Equal(u2 Uint) bool { _ = "STUB: not implemented"; return false }

// GT returns true if first Uint is greater than second
func (u Uint) GT(u2 Uint) bool { _ = "STUB: not implemented"; return false }

// GTE returns true if first Uint is greater than second
func (u Uint) GTE(u2 Uint) bool { _ = "STUB: not implemented"; return false }

// LT returns true if first Uint is lesser than second
func (u Uint) LT(u2 Uint) bool { _ = "STUB: not implemented"; return false }

// LTE returns true if first Uint is lesser than or equal to the second
func (u Uint) LTE(u2 Uint) bool {
	_ = "STUB: not implemented"

	// Add adds Uint from another
	return false
}

func (u Uint) Add(u2 Uint) Uint { _ = "STUB: not implemented"; return *new(Uint) }

// Add convert uint64 and add it to Uint
func (u Uint) AddUint64(u2 uint64) Uint {
	_ = "STUB: not implemented"
	return *

	// Sub subtracts another Uint.
	new(Uint)
}

func (u Uint) Sub(u2 Uint) Uint { _ = "STUB: not implemented"; return *new(Uint) }

// SubUint64 subtracts a uint64 from Uint.
func (u Uint) SubUint64(u2 uint64) Uint {
	_ = "STUB: not implemented"
	return *

	// Mul multiplies two Uints
	new(Uint)
}

func (u Uint) Mul(u2 Uint) (res Uint) { _ = "STUB: not implemented"; return *new(Uint) }

// Mul multiplies two Uints
func (u Uint) MulUint64(u2 uint64) (res Uint) {
	_ = "STUB: not implemented"
	return *

	// Quo divides Uint with Uint
	new(Uint)
}

func (u Uint) Quo(u2 Uint) Uint { _ = "STUB: not implemented"; return *new(Uint) }

// Mod returns remainder after dividing with Uint
func (u Uint) Mod(u2 Uint) Uint { _ = "STUB: not implemented"; return *new(Uint) }

// Incr increments the Uint by one.
func (u Uint) Incr() Uint {
	_ = "STUB: not implemented"
	return *

	// Decr decrements the Uint by one.
	// Decr will panic if the Uint is zero.
	new(Uint)
}

func (u Uint) Decr() Uint {
	_ = "STUB: not implemented"
	return *

	// QuoUint64 divides Uint by a uint64.
	new(Uint)
}

func (u Uint) QuoUint64(u2 uint64) Uint {
	_ = "STUB: not implemented"
	return *

	// MinUint returns the minimum of the Uints.
	new(Uint)
}

func MinUint(u1, u2 Uint) Uint { _ = "STUB: not implemented"; return *new(Uint) }

// Return the maximum of the Uints
func MaxUint(u1, u2 Uint) Uint { _ = "STUB: not implemented"; return *new(Uint) }

// String returns the string representation of Uint.
func (u Uint) String() string {
	_ = "STUB: not implemented"

	// MarshalJSON defines custom encoding scheme
	return ""
}

func (u Uint) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	// Necessary since default Uint initialization has i.i as nil
	return nil, nil
}

// UnmarshalJSON defines custom decoding scheme
func (u *Uint) UnmarshalJSON(bz []byte) error {
	_ = "STUB: not implemented"
	// Necessary since default Uint initialization has i.i as nil
	return nil
}

// Marshal implements the gogo proto custom type interface.
func (u Uint) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalTo implements the gogo proto custom type interface.
func (u *Uint) MarshalTo(data []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// The value 0

// Unmarshal implements the gogo proto custom type interface.
func (u *Uint) Unmarshal(data []byte) error {
	_ = "STUB: not implemented"
	// maxBitLen is 256, which requires ~78 decimal digits, so 100 is a safe upper
	// bound.
	return nil
}

// Use ZeroUint, not Uint{}: a nil *big.Int breaks String(), BigInt(), etc.

// Size implements the gogo proto custom type interface.
func (u *Uint) Size() int { _ = "STUB: not implemented"; return 0 }

// Override Amino binary serialization by proxying to protobuf.
func (u Uint) MarshalAmino() ([]byte, error)   { _ = "STUB: not implemented"; return nil, nil }
func (u *Uint) UnmarshalAmino(bz []byte) error { _ = "STUB: not implemented"; return nil }

// UintOverflow returns true if a given unsigned integer overflows and false
// otherwise.
func UintOverflow(i *big.Int) error { _ = "STUB: not implemented"; return nil }

// ParseUint reads a string-encoded Uint value and return a Uint.
func ParseUint(s string) (Uint, error) { _ = "STUB: not implemented"; return *new(Uint), nil }

func checkNewUint(i *big.Int) (Uint, error) { _ = "STUB: not implemented"; return *new(Uint), nil }

// RelativePow raises x to the power of n, where x (and the result, z) are scaled by factor b
// for example, RelativePow(210, 2, 100) = 441 (2.1^2 = 4.41)
func RelativePow(x Uint, n Uint, b Uint) (z Uint) { _ = "STUB: not implemented"; return *new(Uint) }

// 0^0 = 1

// otherwise 0^a = 0

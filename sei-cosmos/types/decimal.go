package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"testing"
)

var _ CustomProtobufType = (*Dec)(nil)

// NOTE: never use new(Dec) or else we will panic unmarshalling into the
// nil embedded big.Int
type Dec struct {
	i *big.Int
}

const (
	// number of decimal places
	Precision = 18

	// bits required to represent the fractional precision.
	// DecimalPrecisionBits = ceil(log2(10^Precision) - 1) = 60 for Precision=18.
	DecimalPrecisionBits = 60

	// intBitLen is the bit length used by sdk.Int (currently 256-bit)
	intBitLen = 256

	// maxDecBitLen is the maximum allowed bit length for Dec values.
	// It is derived instead of hard-coded so that future changes to Precision
	// or intBitLen automatically propagate.
	// Example with current constants: 256 + 60 − 1 = 315.
	maxDecBitLen = intBitLen + DecimalPrecisionBits - 1

	// max number of iterations in ApproxRoot function
	maxApproxRootIterations = 100
)

var (
	precisionReuse       = new(big.Int).Exp(big.NewInt(10), big.NewInt(Precision), nil)
	fivePrecision        = new(big.Int).Quo(precisionReuse, big.NewInt(2))
	precisionMultipliers []*big.Int
	zeroInt              = big.NewInt(0)
	oneInt               = big.NewInt(1)
	tenInt               = big.NewInt(10)
)

// Decimal errors
var (
	ErrEmptyDecimalStr      = errors.New("decimal string cannot be empty")
	ErrInvalidDecimalLength = errors.New("invalid decimal length")
	ErrInvalidDecimalStr    = errors.New("invalid decimal string")
)

// Set precision multipliers
func init() {
	precisionMultipliers = make([]*big.Int, Precision+1)
	for i := 0; i <= Precision; i++ {
		precisionMultipliers[i] = calcPrecisionMultiplier(int64(i))
	}
}

func precisionInt() *big.Int { _ = "STUB: not implemented"; return nil }

func ZeroDec() Dec     { _ = "STUB: not implemented"; return *new(Dec) }
func OneDec() Dec      { _ = "STUB: not implemented"; return *new(Dec) }
func SmallestDec() Dec { _ = "STUB: not implemented"; return *new(Dec) }

// calculate the precision multiplier
func calcPrecisionMultiplier(prec int64) *big.Int { _ = "STUB: not implemented"; return nil }

// get the precision multiplier, do not mutate result
func precisionMultiplier(prec int64) *big.Int { _ = "STUB: not implemented"; return nil }

// NewDec creates a new Dec from integer assuming whole number.
func NewDec(i int64) Dec { _ = "STUB: not implemented"; return *new(Dec) }

// NewDecWithPrec creates a new Dec from integer with decimal place at prec.
// CONTRACT: prec <= Precision
func NewDecWithPrec(i, prec int64) Dec { _ = "STUB: not implemented"; return *new(Dec) }

// create a new Dec from big integer assuming whole numbers
// CONTRACT: prec <= Precision
func NewDecFromBigInt(i *big.Int) Dec { _ = "STUB: not implemented"; return *new(Dec) }

// create a new Dec from big integer assuming whole numbers
// CONTRACT: prec <= Precision
func NewDecFromBigIntWithPrec(i *big.Int, prec int64) Dec {
	_ = "STUB: not implemented"
	return *new(Dec)
}

// NewDecFromInt creates a new Dec from Int assuming whole numbers.
// CONTRACT: prec <= Precision
func NewDecFromInt(i Int) Dec { _ = "STUB: not implemented"; return *new(Dec) }

// NewDecFromIntWithPrec creates a new Dec from Int with decimal place at prec.
// CONTRACT: prec <= Precision
func NewDecFromIntWithPrec(i Int, prec int64) Dec { _ = "STUB: not implemented"; return *new(Dec) }

// create a decimal from an input decimal string.
// valid must come in the form:
//
//	(-) whole integers (.) decimal integers
//
// examples of acceptable input include:
//
//	-123.456
//	456.7890
//	345
//	-456789
//
// NOTE - An error will return if more decimal places
// are provided in the string than the constant Precision.
//
// CONTRACT - This function does not mutate the input str.
func NewDecFromStr(str string) (Dec, error) { _ = "STUB: not implemented"; return *new(Dec), nil }

// first extract any negative symbol

// has a decimal place

// add some extra zero's to correct to the Precision factor

// base 10

// Decimal from string, panic on error
func MustNewDecFromStr(s string) Dec { _ = "STUB: not implemented"; return *new(Dec) }

func (d Dec) IsNil() bool       { _ = "STUB: not implemented"; return false }     // is decimal nil
func (d Dec) IsZero() bool      { _ = "STUB: not implemented"; return false }     // is equal to zero
func (d Dec) IsNegative() bool  { _ = "STUB: not implemented"; return false }     // is negative
func (d Dec) IsPositive() bool  { _ = "STUB: not implemented"; return false }     // is positive
func (d Dec) Equal(d2 Dec) bool { _ = "STUB: not implemented"; return false }     // equal decimals
func (d Dec) GT(d2 Dec) bool    { _ = "STUB: not implemented"; return false }     // greater than
func (d Dec) GTE(d2 Dec) bool   { _ = "STUB: not implemented"; return false }     // greater than or equal
func (d Dec) LT(d2 Dec) bool    { _ = "STUB: not implemented"; return false }     // less than
func (d Dec) LTE(d2 Dec) bool   { _ = "STUB: not implemented"; return false }     // less than or equal
func (d Dec) Neg() Dec          { _ = "STUB: not implemented"; return *new(Dec) } // reverse the decimal sign
func (d Dec) Abs() Dec          { _ = "STUB: not implemented"; return *new(Dec) } // absolute value

// BigInt returns a copy of the underlying big.Int.
func (d Dec) BigInt() *big.Int { _ = "STUB: not implemented"; return nil }

// Add returns the sum of two Dec values.
func (d Dec) Add(d2 Dec) Dec { _ = "STUB: not implemented"; return *new(Dec) }

// Sub returns the difference of two Dec values.
func (d Dec) Sub(d2 Dec) Dec { _ = "STUB: not implemented"; return *new(Dec) }

// Mul returns the product of two Dec values with rounding.
func (d Dec) Mul(d2 Dec) Dec { _ = "STUB: not implemented"; return *new(Dec) }

// multiplication truncate
func (d Dec) MulTruncate(d2 Dec) Dec { _ = "STUB: not implemented"; return *new(Dec) }

// multiplication
func (d Dec) MulInt(i Int) Dec { _ = "STUB: not implemented"; return *new(Dec) }

// MulInt64 - multiplication with int64
func (d Dec) MulInt64(i int64) Dec { _ = "STUB: not implemented"; return *new(Dec) }

// quotient
func (d Dec) Quo(d2 Dec) Dec {
	_ = "STUB: not implemented"
	// multiply precision twice
	return *new(Dec)
}

// quotient truncate
func (d Dec) QuoTruncate(d2 Dec) Dec {
	_ = "STUB: not implemented"
	// multiply precision twice
	return *new(Dec)
}

// quotient, round up
func (d Dec) QuoRoundUp(d2 Dec) Dec {
	_ = "STUB: not implemented"
	// multiply precision twice
	return *new(Dec)
}

// quotient
func (d Dec) QuoInt(i Int) Dec { _ = "STUB: not implemented"; return *new(Dec) }

// QuoInt64 - quotient with int64
func (d Dec) QuoInt64(i int64) Dec { _ = "STUB: not implemented"; return *new(Dec) }

// ApproxRoot returns an approximate estimation of a Dec's positive real nth root
// using Newton's method (where n is positive). The algorithm starts with some guess and
// computes the sequence of improved guesses until an answer converges to an
// approximate answer.  It returns `|d|.ApproxRoot() * -1` if input is negative.
// A maximum number of 100 iterations is used a backup boundary condition for
// cases where the answer never converges enough to satisfy the main condition.
func (d Dec) ApproxRoot(root uint64) (guess Dec, err error) {
	_ = "STUB: not implemented"
	return *new(Dec), nil
}

// Power returns a the result of raising to a positive integer power
func (d Dec) Power(power uint64) Dec { _ = "STUB: not implemented"; return *new(Dec) }

// ApproxSqrt is a wrapper around ApproxRoot for the common special case
// of finding the square root of a number. It returns -(sqrt(abs(d)) if input is negative.
func (d Dec) ApproxSqrt() (Dec, error) {
	_ = "STUB: not implemented"
	return *

	// is integer, e.g. decimals are zero
	new(Dec), nil
}

func (d Dec) IsInteger() bool { _ = "STUB: not implemented"; return false }

// format decimal state
func (d Dec) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (d Dec) String() string { _ = "STUB: not implemented"; return "" }

// TODO: Remove trailing zeros
// case 1, purely decimal

// 0. prefix

// set relevant digits to 0

// set final digits

// inputSize + 1 to account for the decimal point that is being added

// pre-decimal digits
// decimal point
// post-decimal digits

// Float64 returns the float64 representation of a Dec.
// Will return the error if the conversion failed.
func (d Dec) Float64() (float64, error) { _ = "STUB: not implemented"; return 0, nil }

// MustFloat64 returns the float64 representation of a Dec.
// Would panic if the conversion failed.
func (d Dec) MustFloat64() float64 { _ = "STUB: not implemented"; return 0 }

//     ____
//  __|    |__   "chop 'em
//       ` \     round!"
// ___||  ~  _     -bankers
// |         |      __
// |       | |   __|__|__
// |_____:  /   | $$$    |
//              |________|

// Remove a Precision amount of rightmost digits and perform bankers rounding
// on the remainder (gaussian rounding) on the digits which have been removed.
//
// Mutates the input. Use the non-mutative version if that is undesired
func chopPrecisionAndRound(d *big.Int) *big.Int {
	_ = "STUB: not implemented"
	// remove the negative and add it back when returning
	return nil
}

// make d positive, compute chopped value, and then un-mutate d

// get the truncated quotient and remainder

// remainder is zero

// bankers rounding must take place
// always round to an even number

func chopPrecisionAndRoundUp(d *big.Int) *big.Int {
	_ = "STUB: not implemented"
	// remove the negative and add it back when returning
	return nil
}

// make d positive, compute chopped value, and then un-mutate d

// truncate since d is negative...

// get the truncated quotient and remainder

// remainder is zero

func chopPrecisionAndRoundNonMutative(d *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

// RoundInt64 rounds the decimal using bankers rounding
func (d Dec) RoundInt64() int64 { _ = "STUB: not implemented"; return 0 }

// RoundInt round the decimal using bankers rounding
func (d Dec) RoundInt() Int { _ = "STUB: not implemented"; return *new(Int) }

// chopPrecisionAndTruncate is similar to chopPrecisionAndRound,
// but always rounds down. It does not mutate the input.
func chopPrecisionAndTruncate(d *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

// TruncateInt64 truncates the decimals from the number and returns an int64
func (d Dec) TruncateInt64() int64 { _ = "STUB: not implemented"; return 0 }

// TruncateInt truncates the decimals from the number and returns an Int
func (d Dec) TruncateInt() Int { _ = "STUB: not implemented"; return *new(Int) }

// TruncateDec truncates the decimals from the number and returns a Dec
func (d Dec) TruncateDec() Dec { _ = "STUB: not implemented"; return *new(Dec) }

// Ceil returns the smallest interger value (as a decimal) that is greater than
// or equal to the given decimal.
func (d Dec) Ceil() Dec { _ = "STUB: not implemented"; return *new(Dec) }

// no need to round with a zero remainder regardless of sign

// MaxSortableDec is the largest Dec that can be passed into SortableDecBytes()
// Its negative form is the least Dec that can be passed in.
var MaxSortableDec = OneDec().Quo(SmallestDec())

// ValidSortableDec ensures that a Dec is within the sortable bounds,
// a Dec can't have a precision of less than 10^-18.
// Max sortable decimal was set to the reciprocal of SmallestDec.
func ValidSortableDec(dec Dec) bool { _ = "STUB: not implemented"; return false }

// SortableDecBytes returns a byte slice representation of a Dec that can be sorted.
// Left and right pads with 0s so there are 18 digits to left and right of the decimal point.
// For this reason, there is a maximum and minimum value for this, enforced by ValidSortableDec.
func SortableDecBytes(dec Dec) []byte { _ = "STUB: not implemented"; return nil }

// Instead of adding an extra byte to all sortable decs in order to handle max sortable, we just
// makes its bytes be "max" which comes after all numbers in ASCIIbetical order

// For the same reason, we make the bytes of minimum sortable dec be --, which comes before all numbers.

// We move the negative sign to the front of all the left padded 0s, to make negative numbers come before positive numbers

// reuse nil values
var nilJSON []byte

func init() {
	empty := new(big.Int)
	bz, _ := empty.MarshalText()
	nilJSON, _ = json.Marshal(string(bz))
}

// MarshalJSON marshals the decimal
func (d Dec) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON defines custom decoding scheme
func (d *Dec) UnmarshalJSON(bz []byte) error { _ = "STUB: not implemented"; return nil }

// TODO: Reuse dec allocation

// MarshalYAML returns the YAML representation.
func (d Dec) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil,

		// Marshal implements the gogo proto custom type interface.
		nil
}

func (d Dec) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalTo implements the gogo proto custom type interface.
func (d *Dec) MarshalTo(data []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Unmarshal implements the gogo proto custom type interface.
func (d *Dec) Unmarshal(data []byte) error {
	_ = "STUB: not implemented"
	// The maximum valid Dec value requires ~95 decimal digits (315 bits), so 100 is
	// a safe upper bound.
	return nil
}

// Use ZeroDec, not Dec{}: a nil *big.Int breaks Equal, arithmetic, etc.

// Size implements the gogo proto custom type interface.
func (d *Dec) Size() int { _ = "STUB: not implemented"; return 0 }

// Override Amino binary serialization by proxying to protobuf.
func (d Dec) MarshalAmino() ([]byte, error)   { _ = "STUB: not implemented"; return nil, nil }
func (d *Dec) UnmarshalAmino(bz []byte) error { _ = "STUB: not implemented"; return nil }

func (dp DecProto) String() string { _ = "STUB: not implemented"; return "" }

// helpers

// test if two decimal arrays are equal
func DecsEqual(d1s, d2s []Dec) bool { _ = "STUB: not implemented"; return false }

// minimum decimal between two
func MinDec(d1, d2 Dec) Dec { _ = "STUB: not implemented"; return *new(Dec) }

// maximum decimal between two
func MaxDec(d1, d2 Dec) Dec { _ = "STUB: not implemented"; return *new(Dec) }

// IsInValidRange returns true if the decimal's underlying big.Int is within the valid range.
func (d Dec) IsInValidRange() bool { _ = "STUB: not implemented"; return false }

// Use maxDecBitLen (315 bits) to align with the official Cosmos SDK implementation.
// 315 bits can cover all values within (2^256−1)×10^18 − 1,
// so bitLen ≤ maxDecBitLen ensures alignment with the 256-bit boundary of sdk.Int while also supporting 18-decimal-place precision.

// assertInValidRange panics if the decimal is out of the valid range
func (d Dec) assertInValidRange() { _ = "STUB: not implemented"; return }

// intended to be used with require/assert:  require.True(DecEq(...))
func DecEq(t *testing.T, exp, got Dec) (*testing.T, bool, string, string, string) {
	_ = "STUB: not implemented"
	return nil, false, "", "", ""
}

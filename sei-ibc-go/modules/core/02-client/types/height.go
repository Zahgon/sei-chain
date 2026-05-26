package types

import (
	"regexp"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

var _ exported.Height = (*Height)(nil)

// IsRevisionFormat checks if a chainID is in the format required for parsing revisions
// The chainID must be in the form: `{chainID}-{revision}`.
// 24-host may enforce stricter checks on chainID
var IsRevisionFormat = regexp.MustCompile(`^.*[^\n-]-{1}[1-9][0-9]*$`).MatchString

// ZeroHeight is a helper function which returns an uninitialized height.
func ZeroHeight() Height {
	_ = "STUB: not implemented"

	// NewHeight is a constructor for the IBC height type
	return *new(Height)
}

func NewHeight(revisionNumber, revisionHeight uint64) Height {
	_ = "STUB: not implemented"
	return *new(Height)
}

// GetRevisionNumber returns the revision-number of the height
func (h Height) GetRevisionNumber() uint64 { _ = "STUB: not implemented"; return 0 }

// GetRevisionHeight returns the revision-height of the height
func (h Height) GetRevisionHeight() uint64 { _ = "STUB: not implemented"; return 0 }

// Compare implements a method to compare two heights. When comparing two heights a, b
// we can call a.Compare(b) which will return
// -1 if a < b
// 0  if a = b
// 1  if a > b
//
// It first compares based on revision numbers, whichever has the higher revision number is the higher height
// If revision number is the same, then the revision height is compared
func (h Height) Compare(other exported.Height) int64 { _ = "STUB: not implemented"; return 0 }

// LT Helper comparison function returns true if h < other
func (h Height) LT(other exported.Height) bool { _ = "STUB: not implemented"; return false }

// LTE Helper comparison function returns true if h <= other
func (h Height) LTE(other exported.Height) bool { _ = "STUB: not implemented"; return false }

// GT Helper comparison function returns true if h > other
func (h Height) GT(other exported.Height) bool { _ = "STUB: not implemented"; return false }

// GTE Helper comparison function returns true if h >= other
func (h Height) GTE(other exported.Height) bool { _ = "STUB: not implemented"; return false }

// EQ Helper comparison function returns true if h == other
func (h Height) EQ(other exported.Height) bool { _ = "STUB: not implemented"; return false }

// String returns a string representation of Height
func (h Height) String() string { _ = "STUB: not implemented"; return "" }

// Decrement will return a new height with the RevisionHeight decremented
// If the RevisionHeight is already at lowest value (1), then false success flag is returend
func (h Height) Decrement() (decremented exported.Height, success bool) {
	_ = "STUB: not implemented"
	return *new(exported.Height), false
}

// Increment will return a height with the same revision number but an
// incremented revision height
func (h Height) Increment() exported.Height {
	_ = "STUB: not implemented"
	return *new(exported.Height)
}

// IsZero returns true if height revision and revision-height are both 0
func (h Height) IsZero() bool { _ = "STUB: not implemented"; return false }

// MustParseHeight will attempt to parse a string representation of a height and panic if
// parsing fails.
func MustParseHeight(heightStr string) Height { _ = "STUB: not implemented"; return *new(Height) }

// ParseHeight is a utility function that takes a string representation of the height
// and returns a Height struct
func ParseHeight(heightStr string) (Height, error) {
	_ = "STUB: not implemented"
	return *new(Height), nil
}

// SetRevisionNumber takes a chainID in valid revision format and swaps the revision number
// in the chainID with the given revision number.
func SetRevisionNumber(chainID string, revision uint64) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// swap out revision number with given revision

// ParseChainID is a utility function that returns an revision number from the given ChainID.
// ParseChainID attempts to parse a chain id in the format: `{chainID}-{revision}`
// and return the revisionnumber as a uint64.
// If the chainID is not in the expected format, a default revision value of 0 is returned.
func ParseChainID(chainID string) uint64 { _ = "STUB: not implemented"; return 0 }

// chainID is not in revision format, return 0 as default

// sanity check: error should always be nil since regex only allows numbers in last element

// GetSelfHeight is a utility function that returns self height given context
// Revision number is retrieved from ctx.ChainID()
func GetSelfHeight(ctx sdk.Context) Height { _ = "STUB: not implemented"; return *new(Height) }

// #nosec G115 -- block height is checked above to be non-negative

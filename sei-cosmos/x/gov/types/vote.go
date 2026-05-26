package types

import (
	"fmt"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// NewVote creates a new Vote instance
func NewVote(proposalID uint64, voter sdk.AccAddress, options WeightedVoteOptions) Vote {
	_ = "STUB: not implemented"
	return *new(Vote)
}

func (v Vote) String() string { _ = "STUB: not implemented"; return "" }

// Votes is a collection of Vote objects
type Votes []Vote

// Equal returns true if two slices (order-dependant) of votes are equal.
func (v Votes) Equal(other Votes) bool { _ = "STUB: not implemented"; return false }

func (v Votes) String() string { _ = "STUB: not implemented"; return "" }

// Empty returns whether a vote is empty.
func (v Vote) Empty() bool { _ = "STUB: not implemented"; return false }

// NewNonSplitVoteOption creates a single option vote with weight 1
func NewNonSplitVoteOption(option VoteOption) WeightedVoteOptions {
	_ = "STUB: not implemented"
	return *new(WeightedVoteOptions)
}

func (v WeightedVoteOption) String() string { _ = "STUB: not implemented"; return "" }

// WeightedVoteOptions describes array of WeightedVoteOptions
type WeightedVoteOptions []WeightedVoteOption

func (v WeightedVoteOptions) String() (out string) { _ = "STUB: not implemented"; return "" }

// ValidWeightedVoteOption returns true if the sub vote is valid and false otherwise.
func ValidWeightedVoteOption(option WeightedVoteOption) bool {
	_ = "STUB: not implemented"
	return false
}

// VoteOptionFromString returns a VoteOption from a string. It returns an error
// if the string is invalid.
func VoteOptionFromString(str string) (VoteOption, error) {
	_ = "STUB: not implemented"
	return *new(VoteOption), nil
}

// WeightedVoteOptionsFromString returns weighted vote options from string. It returns an error
// if the string is invalid.
func WeightedVoteOptionsFromString(str string) (WeightedVoteOptions, error) {
	_ = "STUB: not implemented"
	return *new(WeightedVoteOptions), nil
}

// ValidVoteOption returns true if the vote option is valid and false otherwise.
func ValidVoteOption(option VoteOption) bool { _ = "STUB: not implemented"; return false }

// Marshal needed for protobuf compatibility.
func (vo VoteOption) Marshal() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Unmarshal needed for protobuf compatibility.
		nil
}

func (vo *VoteOption) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }

// Format implements the fmt.Formatter interface.
func (vo VoteOption) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

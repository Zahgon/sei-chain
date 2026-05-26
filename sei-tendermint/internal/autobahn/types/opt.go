package types

import (
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

// NextOpt defaults to 0.
func NextOpt[I ~uint64, T interface{ Next() I }](mv utils.Option[T]) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// ProposalOpt extracts optional proposal from optional value.
func ProposalOpt[P any, T interface{ Proposal() P }](mv utils.Option[T]) utils.Option[P] {
	_ = "STUB: not implemented"
	return nil
}

// NextIndexOpt defaults to 0.
func NextIndexOpt[T interface{ Index() RoadIndex }](mv utils.Option[T]) RoadIndex {
	_ = "STUB: not implemented"
	return *new(RoadIndex)
}

// NextViewOpt defaults to {0,0}.
func NextViewOpt[T interface{ View() View }](mv utils.Option[T]) View {
	_ = "STUB: not implemented"
	return *new(View)
}

// LaneRangeOpt defaults to an empty initial range.
func LaneRangeOpt[T interface {
	LaneRange(lane LaneID) *LaneRange
}](mv utils.Option[T], lane LaneID) *LaneRange {
	_ = "STUB: not implemented"
	return nil
}

// GlobalRangeOpt defaults to an empty initial range.
func GlobalRangeOpt[T interface {
	GlobalRange(c *Committee) GlobalRange
}](mv utils.Option[T], c *Committee) GlobalRange {
	_ = "STUB: not implemented"
	return *new(GlobalRange)
}

// AppOpt defaults to None.
func AppOpt[T interface {
	App() utils.Option[*AppProposal]
}](mv utils.Option[T]) utils.Option[*AppProposal] {
	_ = "STUB: not implemented"
	return nil
}

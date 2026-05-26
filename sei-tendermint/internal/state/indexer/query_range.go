package indexer

import (
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/pubsub/query/syntax"
)

// QueryRanges defines a mapping between a composite event key and a QueryRange.
//
// e.g.account.number => queryRange{lowerBound: 1, upperBound: 5}
type QueryRanges map[string]QueryRange

// QueryRange defines a range within a query condition.
type QueryRange struct {
	LowerBound        interface{} // int || time.Time
	UpperBound        interface{} // int || time.Time
	Key               string
	IncludeLowerBound bool
	IncludeUpperBound bool
}

// AnyBound returns either the lower bound if non-nil, otherwise the upper bound.
func (qr QueryRange) AnyBound() interface{} { _ = "STUB: not implemented"; return nil }

// LowerBoundValue returns the value for the lower bound. If the lower bound is
// nil, nil will be returned.
func (qr QueryRange) LowerBoundValue() interface{} { _ = "STUB: not implemented"; return nil }

// UpperBoundValue returns the value for the upper bound. If the upper bound is
// nil, nil will be returned.
func (qr QueryRange) UpperBoundValue() interface{} { _ = "STUB: not implemented"; return nil }

// LookForRanges returns a mapping of QueryRanges and the matching indexes in
// the provided query conditions.
func LookForRanges(conditions []syntax.Condition) (ranges QueryRanges, indexes []int) {
	_ = "STUB: not implemented"
	return *new(QueryRanges), nil
}

// IsRangeOperation returns a boolean signifying if a query Operator is a range
// operation or not.
func IsRangeOperation(op syntax.Token) bool { _ = "STUB: not implemented"; return false }

func conditionArg(c syntax.Condition) interface{} { _ = "STUB: not implemented"; return nil }

// string

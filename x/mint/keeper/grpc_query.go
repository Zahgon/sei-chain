package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/x/mint/types"
)

var _ types.QueryServer = Querier{}

// Querier defines a wrapper around the x/mint keeper providing gRPC method
// handlers.
type Querier struct {
	Keeper
}

func NewQuerier(k Keeper) Querier {
	_ = "STUB: not implemented"
	return *

	// Params returns params of the mint module.
	new(Querier)
}

func (q Querier) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns the most last mint state
func (q Querier) Minter(c context.Context, _ *types.QueryMinterRequest) (*types.QueryMinterResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types/proposal"
)

var _ proposal.QueryServer = Keeper{}

// Params returns subspace params
func (k Keeper) Params(c context.Context, req *proposal.QueryParamsRequest) (*proposal.QueryParamsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

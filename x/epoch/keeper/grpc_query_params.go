package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/x/epoch/types"
)

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

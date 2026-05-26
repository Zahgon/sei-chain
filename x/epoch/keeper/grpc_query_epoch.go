package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/x/epoch/types"
)

func (k Keeper) Epoch(c context.Context, _ *types.QueryEpochRequest) (*types.QueryEpochResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

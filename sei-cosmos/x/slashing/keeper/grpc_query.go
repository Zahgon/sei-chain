package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-cosmos/x/slashing/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k Keeper) SigningInfo(c context.Context, req *types.QuerySigningInfoRequest) (*types.QuerySigningInfoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k Keeper) SigningInfos(c context.Context, req *types.QuerySigningInfosRequest) (*types.QuerySigningInfosResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

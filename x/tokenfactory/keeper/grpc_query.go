package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/x/tokenfactory/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Params(ctx context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k Keeper) DenomAuthorityMetadata(ctx context.Context, req *types.QueryDenomAuthorityMetadataRequest) (*types.QueryDenomAuthorityMetadataResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k Keeper) DenomsFromCreator(ctx context.Context, req *types.QueryDenomsFromCreatorRequest) (*types.QueryDenomsFromCreatorResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DenomMetadata implements Query/DenomMetadata gRPC method.
func (k Keeper) DenomMetadata(c context.Context, req *types.QueryDenomMetadataRequest) (*types.QueryDenomMetadataResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k Keeper) DenomAllowList(c context.Context, req *types.QueryDenomAllowListRequest) (*types.QueryDenomAllowListResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

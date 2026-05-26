package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/27-interchain-accounts/controller/types"
)

var _ types.QueryServer = Keeper{}

// InterchainAccount implements the Query/InterchainAccount gRPC method
func (k Keeper) InterchainAccount(goCtx context.Context, req *types.QueryInterchainAccountRequest) (*types.QueryInterchainAccountResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Params implements the Query/Params gRPC method
func (k Keeper) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

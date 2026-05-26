package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/types"
)

var _ types.QueryServer = BaseKeeper{}

// Balance implements the Query/Balance gRPC method
func (k BaseKeeper) Balance(ctx context.Context, req *types.QueryBalanceRequest) (*types.QueryBalanceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AllBalances implements the Query/AllBalances gRPC method
func (k BaseKeeper) AllBalances(ctx context.Context, req *types.QueryAllBalancesRequest) (*types.QueryAllBalancesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SpendableBalances implements a gRPC query handler for retrieving an account's
// spendable balances.
func (k BaseKeeper) SpendableBalances(ctx context.Context, req *types.QuerySpendableBalancesRequest) (*types.QuerySpendableBalancesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TotalSupply implements the Query/TotalSupply gRPC method
func (k BaseKeeper) TotalSupply(ctx context.Context, req *types.QueryTotalSupplyRequest) (*types.QueryTotalSupplyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SupplyOf implements the Query/SupplyOf gRPC method
func (k BaseKeeper) SupplyOf(c context.Context, req *types.QuerySupplyOfRequest) (*types.QuerySupplyOfResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Params implements the gRPC service handler for querying x/bank parameters.
func (k BaseKeeper) Params(ctx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DenomsMetadata implements Query/DenomsMetadata gRPC method.
func (k BaseKeeper) DenomsMetadata(c context.Context, req *types.QueryDenomsMetadataRequest) (*types.QueryDenomsMetadataResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DenomMetadata implements Query/DenomMetadata gRPC method.
func (k BaseKeeper) DenomMetadata(c context.Context, req *types.QueryDenomMetadataRequest) (*types.QueryDenomMetadataResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

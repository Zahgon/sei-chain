package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-cosmos/x/authz"
)

var _ authz.QueryServer = Keeper{}

// Authorizations implements the Query/Grants gRPC method.
func (k Keeper) Grants(c context.Context, req *authz.QueryGrantsRequest) (*authz.QueryGrantsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GranterGrants implements the Query/GranterGrants gRPC method.
func (k Keeper) GranterGrants(c context.Context, req *authz.QueryGranterGrantsRequest) (*authz.QueryGranterGrantsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GranteeGrants implements the Query/GranteeGrants gRPC method.
func (k Keeper) GranteeGrants(c context.Context, req *authz.QueryGranteeGrantsRequest) (*authz.QueryGranteeGrantsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

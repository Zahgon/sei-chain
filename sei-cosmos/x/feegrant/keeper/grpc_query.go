package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-cosmos/x/feegrant"
)

var _ feegrant.QueryServer = Keeper{}

// Allowance returns fee granted to the grantee by the granter.
func (q Keeper) Allowance(c context.Context, req *feegrant.QueryAllowanceRequest) (*feegrant.QueryAllowanceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Allowances queries all the allowances granted to the given grantee.
func (q Keeper) Allowances(c context.Context, req *feegrant.QueryAllowancesRequest) (*feegrant.QueryAllowancesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AllowancesByGranter queries all the allowances granted by the given granter
func (q Keeper) AllowancesByGranter(c context.Context, req *feegrant.QueryAllowancesByGranterRequest) (*feegrant.QueryAllowancesByGranterResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseAddressesFromFeeAllowanceKey expects the full key including the prefix.

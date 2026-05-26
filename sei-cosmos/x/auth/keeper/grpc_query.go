package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/types"
)

var _ types.QueryServer = AccountKeeper{}

func (ak AccountKeeper) Accounts(c context.Context, req *types.QueryAccountsRequest) (*types.QueryAccountsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Account returns account details based on address
func (ak AccountKeeper) Account(c context.Context, req *types.QueryAccountRequest) (*types.QueryAccountResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Params returns parameters of auth module
func (ak AccountKeeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ak AccountKeeper) NextAccountNumber(c context.Context, req *types.QueryNextAccountNumberRequest) (*types.QueryNextAccountNumberResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

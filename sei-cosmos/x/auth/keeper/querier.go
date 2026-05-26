package keeper

import (
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// NewQuerier creates a querier for auth REST endpoints
func NewQuerier(k AccountKeeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	_ = "STUB: not implemented"
	return *new(sdk.Querier)
}

func queryAccount(ctx sdk.Context, req abci.RequestQuery, k AccountKeeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryParams(ctx sdk.Context, k AccountKeeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

package keeper

import (
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// NewQuerier creates a new querier for slashing clients.
func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	_ = "STUB: not implemented"
	return *new(sdk.Querier)
}

func queryParams(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func querySigningInfo(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func querySigningInfos(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

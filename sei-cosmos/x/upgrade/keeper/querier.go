package keeper

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"

	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// NewQuerier creates a querier for upgrade cli and REST endpoints
func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	_ = "STUB: not implemented"
	return *new(sdk.Querier)
}

func queryCurrent(ctx sdk.Context, _ abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryApplied(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // bounds checked above

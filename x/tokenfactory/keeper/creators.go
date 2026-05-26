package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func (k Keeper) addDenomFromCreator(ctx sdk.Context, creator, denom string) {
	_ = "STUB: not implemented"
	return
}

func (k Keeper) getDenomsFromCreator(ctx sdk.Context, creator string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (k Keeper) GetAllDenomsIterator(ctx sdk.Context) sdk.Iterator {
	_ = "STUB: not implemented"
	return *new(sdk.Iterator)
}

package store

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func GetCachedContext(ctx sdk.Context) (sdk.Context, sdk.CacheMultiStore) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), *new(sdk.CacheMultiStore)
}

package testutil

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// DefaultContext creates a sdk.Context with a fresh MemDB that can be used in tests.
func DefaultContext(key sdk.StoreKey, tkey sdk.StoreKey) sdk.Context {
	_ = "STUB: not implemented"
	return *new(sdk.Context)
}

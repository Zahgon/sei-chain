package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

type MintHooks interface {
	AfterDistributeMintedCoin(ctx sdk.Context, mintedCoin sdk.Coin)
}

var _ MintHooks = MultiMintHooks{}

// combine multiple mint hooks, all hook functions are run in array sequence.
type MultiMintHooks []MintHooks

func NewMultiMintHooks(hooks ...MintHooks) MultiMintHooks {
	_ = "STUB: not implemented"
	return *new(MultiMintHooks)
}

func (h MultiMintHooks) AfterDistributeMintedCoin(ctx sdk.Context, mintedCoin sdk.Coin) {
	_ = "STUB: not implemented"
	return
}

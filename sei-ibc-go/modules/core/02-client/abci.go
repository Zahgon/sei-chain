package client

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/keeper"
)

// BeginBlocker updates an existing localhost client with the latest block height.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) { _ = "STUB: not implemented"; return }

// Once we are at the last block this chain will commit, set the upgraded consensus state
// so that IBC clients can use the last NextValidatorsHash as a trusted kernel for verifying
// headers on the next version of the chain.
// Set the time to the last block time of the current chain.
// In order for a client to upgrade successfully, the first block of the new chain must be committed
// within the trusting period of the last block time on this chain.

// update the localhost client with the latest block height

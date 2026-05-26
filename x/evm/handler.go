package evm

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	govtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"

	"github.com/sei-protocol/sei-chain/x/evm/keeper"
)

func NewHandler(k *keeper.Keeper) sdk.Handler { _ = "STUB: not implemented"; return *new(sdk.Handler) }

func NewProposalHandler(k keeper.Keeper) govtypes.Handler {
	_ = "STUB: not implemented"
	return *new(govtypes.Handler)
}

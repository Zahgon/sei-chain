package distribution

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution/keeper"
	govtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler { _ = "STUB: not implemented"; return *new(sdk.Handler) }

func NewCommunityPoolSpendProposalHandler(k keeper.Keeper) govtypes.Handler {
	_ = "STUB: not implemented"
	return *new(govtypes.Handler)
}

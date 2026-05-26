package tokenfactory

import (
	govtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"

	"github.com/sei-protocol/sei-chain/x/tokenfactory/keeper"
)

func NewProposalHandler(_ keeper.Keeper) govtypes.Handler {
	_ = "STUB: not implemented"
	return *new(govtypes.Handler)
}

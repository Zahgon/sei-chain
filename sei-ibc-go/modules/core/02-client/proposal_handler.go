package client

import (
	govtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/keeper"
)

// NewClientProposalHandler defines the 02-client proposal handler
func NewClientProposalHandler(k keeper.Keeper) govtypes.Handler {
	_ = "STUB: not implemented"
	return *new(govtypes.Handler)
}

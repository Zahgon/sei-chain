package params

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	govtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/params/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types/proposal"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("cosmos", "x", "params")

// NewParamChangeProposalHandler creates a new governance Handler for a ParamChangeProposal
func NewParamChangeProposalHandler(k keeper.Keeper) govtypes.Handler {
	_ = "STUB: not implemented"
	return *new(govtypes.Handler)
}

func handleParameterChangeProposal(ctx sdk.Context, k keeper.Keeper, p *proposal.ParameterChangeProposal) error {
	_ = "STUB: not implemented"
	return nil
}

package upgrade

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	govtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade/types"
)

// NewSoftwareUpgradeProposalHandler creates a governance handler to manage new proposal types.
// It enables SoftwareUpgradeProposal to propose an Upgrade, and CancelSoftwareUpgradeProposal
// to abort a previously voted upgrade.
func NewSoftwareUpgradeProposalHandler(k keeper.Keeper) govtypes.Handler {
	_ = "STUB: not implemented"
	return *new(govtypes.Handler)
}

func handleSoftwareUpgradeProposal(ctx sdk.Context, k keeper.Keeper, p *types.SoftwareUpgradeProposal) error {
	_ = "STUB: not implemented"
	return nil
}

func handleCancelSoftwareUpgradeProposal(ctx sdk.Context, k keeper.Keeper, _ *types.CancelSoftwareUpgradeProposal) error {
	_ = "STUB: not implemented"
	return nil
}

package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
)

// GetDepositParams returns the current DepositParams from the global param store
func (keeper Keeper) GetDepositParams(ctx sdk.Context) types.DepositParams {
	_ = "STUB: not implemented"
	return *new(types.DepositParams)
}

// GetVotingParams returns the current VotingParams from the global param store
func (keeper Keeper) GetVotingParams(ctx sdk.Context) types.VotingParams {
	_ = "STUB: not implemented"
	return *new(types.VotingParams)
}

// GetTallyParams returns the current TallyParam from the global param store
func (keeper Keeper) GetTallyParams(ctx sdk.Context) types.TallyParams {
	_ = "STUB: not implemented"
	return *new(types.TallyParams)
}

// SetDepositParams sets DepositParams to the global param store
func (keeper Keeper) SetDepositParams(ctx sdk.Context, depositParams types.DepositParams) {
	_ = "STUB: not implemented"
	return
}

// SetVotingParams sets VotingParams to the global param store
func (keeper Keeper) SetVotingParams(ctx sdk.Context, votingParams types.VotingParams) {
	_ = "STUB: not implemented"
	return
}

// SetTallyParams sets TallyParams to the global param store
func (keeper Keeper) SetTallyParams(ctx sdk.Context, tallyParams types.TallyParams) {
	_ = "STUB: not implemented"
	return
}

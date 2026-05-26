package staking

import (
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/types"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
)

// InitGenesis sets the pool and parameters for the provided keeper.  For each
// validator in data, it sets that validator in the keeper along with manually
// setting the indexes. In addition, it also sets any delegations found in
// data. Finally, it updates the bonded validators.
// Returns final validator set after applying all declaration and delegations
func InitGenesis(
	ctx sdk.Context, keeper keeper.Keeper, accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper, data *types.GenesisState,
) (res []abci.ValidatorUpdate) {
	_ = "STUB: not implemented"
	return nil
}

// We need to pretend to be "n blocks before genesis", where "n" is the
// validator update delay, so that e.g. slashing periods are correctly
// initialized for the validator set e.g. with a one-block offset - the
// first TM block is at height 1, so state updates applied from
// genesis.json are in block 0.

// Manually set indices for the first time

// Call the creation hook if not exported

// update timeslice if necessary

// Call the before-creation hook if not exported

// Call the after-modification hook if not exported

// check if the unbonded and bonded pools accounts exists

// TODO remove with genesis 2-phases refactor https://github.com/cosmos/cosmos-sdk/issues/2862

// if balance is different from bonded coins panic because genesis is most likely malformed

// if balance is different from non bonded coins panic because genesis is most likely malformed

// don't need to run Tendermint updates if we exported

// keep the next-val-set offset, use the last power for the first block

// ExportGenesis returns a GenesisState for a given context and keeper. The
// GenesisState will contain the pool, params, validators, and bonds found in
// the keeper.
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) *types.GenesisState {
	_ = "STUB: not implemented"
	return nil
}

// WriteValidators returns a slice of bonded genesis validators.
func WriteValidators(ctx sdk.Context, keeper keeper.Keeper) (vals []tmtypes.GenesisValidator, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ValidateGenesis validates the provided staking genesis state to ensure the
// expected invariants holds. (i.e. params in correct bounds, no duplicate validators)
func ValidateGenesis(data *types.GenesisState) error { _ = "STUB: not implemented"; return nil }

func validateGenesisStateValidators(validators []types.Validator) error {
	_ = "STUB: not implemented"
	return nil
}

package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
)

// ValidatorSetSource is a subset of the staking keeper
type ValidatorSetSource interface {
	ApplyAndReturnValidatorSetUpdates(sdk.Context) (updates []abci.ValidatorUpdate, err error)
}

// InitGenesis sets supply information for genesis.
//
// CONTRACT: all types of accounts must have been already initialized/created
func InitGenesis(ctx sdk.Context, keeper *Keeper, data types.GenesisState, stakingKeeper ValidatorSetSource, msgHandler sdk.Handler) ([]abci.ValidatorUpdate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// #nosec G115 -- loop index i is always non-negative

// sanity check seq values

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper *Keeper) *types.GenesisState {
	_ = "STUB: not implemented"
	return nil
}

// redact contract info

const GENSIS_STATE_STREAM_BUF_THRESHOLD = 50000

func ExportGenesisStream(ctx sdk.Context, keeper *Keeper) <-chan *types.GenesisState {
	_ = "STUB: not implemented"
	return nil
}

// Needs to be first because there are invariant checks when importing that need sequences info

// redact contract info

// flush any remaining state

package keeper

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/x/tokenfactory/types"

	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
)

type (
	Keeper struct {
		storeKey sdk.StoreKey

		paramSpace paramtypes.Subspace

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
		distrKeeper   types.DistrKeeper
	}
)

// NewKeeper returns a new instance of the x/tokenfactory keeper
func NewKeeper(
	_ codec.Codec,
	storeKey sdk.StoreKey,
	paramSpace paramtypes.Subspace,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	distrKeeper types.DistrKeeper,
) Keeper {
	_ = "STUB: not implemented"
	return *new(Keeper)
}

// GetDenomPrefixStore returns the substore for a specific denom
func (k Keeper) GetDenomPrefixStore(ctx sdk.Context, denom string) sdk.KVStore {
	_ = "STUB: not implemented"
	return *new(sdk.KVStore)
}

// GetCreatorPrefixStore returns the substore for a specific creator address
func (k Keeper) GetCreatorPrefixStore(ctx sdk.Context, creator string) sdk.KVStore {
	_ = "STUB: not implemented"
	return *new(sdk.KVStore)
}

// GetCreatorsPrefixStore returns the substore that contains a list of creators
func (k Keeper) GetCreatorsPrefixStore(ctx sdk.Context) sdk.KVStore {
	_ = "STUB: not implemented"
	return *new(sdk.KVStore)
}

// CreateModuleAccount creates a module account with minting and burning capabilities
// This account isn't intended to store any coins,
// it purely mints and burns them on behalf of the admin of respective denoms,
// and sends to the relevant address.
func (k Keeper) CreateModuleAccount(ctx sdk.Context) { _ = "STUB: not implemented"; return }

func (k Keeper) GetDenomAllowListMaxSize(ctx sdk.Context) uint32 {
	_ = "STUB: not implemented"
	return 0
}

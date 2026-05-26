package keeper

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
)

// Implements ValidatorSet interface
var _ types.ValidatorSet = Keeper{}

// Implements DelegationSet interface
var _ types.DelegationSet = Keeper{}

// keeper of the staking store
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.BinaryCodec
	authKeeper types.AccountKeeper
	bankKeeper types.BankKeeper
	hooks      types.StakingHooks
	paramstore paramtypes.Subspace
}

// NewKeeper creates a new staking Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec, key sdk.StoreKey, ak types.AccountKeeper, bk types.BankKeeper,
	ps paramtypes.Subspace,
) Keeper {
	_ = "STUB: not implemented"
	// set KeyTable if it has not already been set
	return *new(Keeper)
}

// ensure bonded and not bonded module accounts are set

func (k Keeper) GetStoreKey() sdk.StoreKey {
	_ = "STUB: not implemented"

	// Set the validator hooks
	return *new(sdk.StoreKey)
}

func (k *Keeper) SetHooks(sh types.StakingHooks) *Keeper { _ = "STUB: not implemented"; return nil }

// Load the last total validator power.
func (k Keeper) GetLastTotalPower(ctx sdk.Context) sdk.Int {
	_ = "STUB: not implemented"
	return *new(sdk.Int)
}

// Set the last total validator power.
func (k Keeper) SetLastTotalPower(ctx sdk.Context, power sdk.Int) {
	_ = "STUB: not implemented"
	return
}

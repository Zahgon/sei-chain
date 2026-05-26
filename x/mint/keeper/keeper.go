package keeper

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
	epochTypes "github.com/sei-protocol/sei-chain/x/epoch/types"
	"github.com/sei-protocol/sei-chain/x/mint/types"
)

// Keeper of the mint store
type Keeper struct {
	cdc              codec.BinaryCodec
	storeKey         sdk.StoreKey
	paramSpace       paramtypes.Subspace
	stakingKeeper    types.StakingKeeper
	bankKeeper       types.BankKeeper
	hooks            types.MintHooks
	feeCollectorName string
}

// NewKeeper creates a new mint Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec, key sdk.StoreKey, paramSpace paramtypes.Subspace,
	sk types.StakingKeeper, ak types.AccountKeeper, bk types.BankKeeper,
	_ types.EpochKeeper, feeCollectorName string,
) Keeper {
	_ = "STUB: not implemented"
	// ensure mint module account is set
	return *new(Keeper)
}

// set KeyTable if it has not already been set

// Set the mint hooks.
func (k *Keeper) SetHooks(h types.MintHooks) *Keeper { _ = "STUB: not implemented"; return nil }

// get the minter
func (k Keeper) GetMinter(ctx sdk.Context) (minter types.Minter) {
	_ = "STUB: not implemented"
	return *new(types.Minter)
}

// set the minter
func (k Keeper) SetMinter(ctx sdk.Context, minter types.Minter) { _ = "STUB: not implemented"; return }

// GetParams returns the total set of minting parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	_ = "STUB: not implemented"
	return *new(types.Params)
}

// SetParams sets the total set of minting parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) { _ = "STUB: not implemented"; return }

// StakingTokenSupply implements an alias call to the underlying staking keeper's
func (k Keeper) StakingTokenSupply(ctx sdk.Context) sdk.Int {
	_ = "STUB: not implemented"
	return *new(sdk.Int)
}

// BondedRatio implements an alias call to the underlying staking keeper's
func (k Keeper) BondedRatio(ctx sdk.Context) sdk.Dec {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// MintCoins implements an alias call to the underlying supply keeper's
// MintCoins to be used in BeginBlocker.
func (k Keeper) MintCoins(ctx sdk.Context, newCoins sdk.Coins) error {
	_ = "STUB: not implemented"
	return nil

	// skip as no coins need to be minted
}

// AddCollectedFees implements an alias call to the underlying supply keeper's
// AddCollectedFees to be used in BeginBlocker.
func (k Keeper) AddCollectedFees(ctx sdk.Context, fees sdk.Coins) error {
	_ = "STUB: not implemented"
	return nil
}

// GetProportions gets the balance of the `MintedDenom` from minted coins and returns coins according to the `AllocationRatio`.
func (k Keeper) GetOrUpdateLatestMinter(
	ctx sdk.Context,
	epoch epochTypes.Epoch,
) types.Minter {
	_ = "STUB: not implemented"
	return *new(types.Minter)
}

// There's still an ongoing release (> 0 remaining amount or same start date) or there's no release scheduled

func (k Keeper) GetCdc() codec.BinaryCodec {
	_ = "STUB: not implemented"
	return *new(codec.BinaryCodec)
}

func (k Keeper) GetStoreKey() sdk.StoreKey { _ = "STUB: not implemented"; return *new(sdk.StoreKey) }

func (k Keeper) GetParamSpace() paramtypes.Subspace {
	_ = "STUB: not implemented"
	return *new(paramtypes.Subspace)
}

func (k *Keeper) SetParamSpace(subspace paramtypes.Subspace) { _ = "STUB: not implemented"; return }

func GetNextScheduledTokenRelease(
	epoch epochTypes.Epoch,
	tokenReleaseSchedule []types.ScheduledTokenRelease,
	currentMinter types.Minter,
) *types.ScheduledTokenRelease {
	_ = "STUB: not implemented"
	return nil
}

// This should not happen as the scheduled release date is validated when the param is updated

// If epoch is after the currentScheduled date and it's after the current release

package keeper

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/ante"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/feegrant"
)

// Keeper manages state of all fee grants, as well as calculating approval.
// It must have a codec with all available allowances registered.
type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   sdk.StoreKey
	authKeeper feegrant.AccountKeeper
}

var _ ante.FeegrantKeeper = &Keeper{}

// NewKeeper creates a fee grant Keeper
func NewKeeper(cdc codec.BinaryCodec, storeKey sdk.StoreKey, ak feegrant.AccountKeeper) Keeper {
	_ = "STUB: not implemented"
	return *new(Keeper)
}

// GrantAllowance creates a new grant
func (k Keeper) GrantAllowance(ctx sdk.Context, granter, grantee sdk.AccAddress, feeAllowance feegrant.FeeAllowanceI) error {
	_ = "STUB: not implemented"

	// create the account if it is not in account state
	return nil
}

// revokeAllowance removes an existing grant
func (k Keeper) revokeAllowance(ctx sdk.Context, granter, grantee sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// GetAllowance returns the allowance between the granter and grantee.
// If there is none, it returns nil, nil.
// Returns an error on parsing issues
func (k Keeper) GetAllowance(ctx sdk.Context, granter, grantee sdk.AccAddress) (feegrant.FeeAllowanceI, error) {
	_ = "STUB: not implemented"
	return *new(feegrant.FeeAllowanceI), nil
}

// getGrant returns entire grant between both accounts
func (k Keeper) getGrant(ctx sdk.Context, granter sdk.AccAddress, grantee sdk.AccAddress) (*feegrant.Grant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IterateAllFeeAllowances iterates over all the grants in the store.
// Callback to get all data, returns true to stop, false to keep reading
// Calling this without pagination is very expensive and only designed for export genesis
func (k Keeper) IterateAllFeeAllowances(ctx sdk.Context, cb func(grant feegrant.Grant) bool) error {
	_ = "STUB: not implemented"
	return nil
}

// UseGrantedFees will try to pay the given fee from the granter's account as requested by the grantee
func (k Keeper) UseGrantedFees(ctx sdk.Context, granter, grantee sdk.AccAddress, fee sdk.Coins, msgs []sdk.Msg) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignoring the `revokeFeeAllowance` error, because the user has enough grants to perform this transaction.

// if fee allowance is accepted, store the updated state of the allowance

func emitUseGrantEvent(ctx sdk.Context, granter, grantee string) { _ = "STUB: not implemented"; return }

// InitGenesis will initialize the keeper from a *previously validated* GenesisState
func (k Keeper) InitGenesis(ctx sdk.Context, data *feegrant.GenesisState) error {
	_ = "STUB: not implemented"
	return nil
}

// ExportGenesis will dump the contents of the keeper into a serializable GenesisState.
func (k Keeper) ExportGenesis(ctx sdk.Context) (*feegrant.GenesisState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	banktypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/types"

	"github.com/sei-protocol/sei-chain/x/tokenfactory/types"
)

// CreateDenom creates a new token denom with the given subdenom.
func (k Keeper) CreateDenom(ctx sdk.Context, creatorAddr string, subdenom string) (newTokenDenom string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Runs CreateDenom logic after the charge and all denom validation has been handled.
// Made into a second function for genesis initialization.
func (k Keeper) createDenomAfterValidation(ctx sdk.Context, creatorAddr string, denom string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// The following is necessary for x/bank denom validation

func (k Keeper) validateCreateDenom(ctx sdk.Context, creatorAddr string, subdenom string) (newTokenDenom string, err error) {
	_ = "STUB: not implemented"
	// Temporary check until IBC bug is sorted out
	return "", nil
}

func (k Keeper) validateUpdateDenom(ctx sdk.Context, msg *types.MsgUpdateDenom) (tokenDenom string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (k Keeper) validateAllowListSize(ctx sdk.Context, allowList *banktypes.AllowList) error {
	_ = "STUB: not implemented"
	return nil
}

func (k Keeper) validateAllowList(ctx sdk.Context, allowList *banktypes.AllowList) error {
	_ = "STUB: not implemented"
	return nil
}

// validate all addresses in the allow list are bech32

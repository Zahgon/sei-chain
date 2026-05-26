package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// RegisterInterchainAccount attempts to create a new account using the provided address and
// stores it in state keyed by the provided connection and port identifiers
// If an account for the provided address already exists this function returns early (no-op)
// NOTE: This function is deprecated!
func (k Keeper) RegisterInterchainAccount(ctx sdk.Context, connectionID, controllerPortID string, accAddress sdk.AccAddress) {
	_ = "STUB: not implemented"
	return
}

// createInterchainAccount creates a new interchain account. An address is generated using the host connectionID, the controller portID,
// and block dependent information. An error is returned if an account already exists for the generated account.
// An interchain account type is set in the account keeper and the interchain account address mapping is updated.
func (k Keeper) createInterchainAccount(ctx sdk.Context, connectionID, controllerPortID string) (sdk.AccAddress, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), nil
}

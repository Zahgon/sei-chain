package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	capabilitytypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/types"

	channeltypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/04-channel/types"
)

// OnChanOpenInit performs basic validation of channel initialization.
// The channel order must be ORDERED, the counterparty port identifier
// must be the host chain representation as defined in the types package,
// the channel version must be equal to the version in the types package,
// there must not be an active channel for the specfied port identifier,
// and the interchain accounts module must be able to claim the channel
// capability.
func (k Keeper) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanOpenAck sets the active channel for the interchain account/owner pair
// and stores the associated interchain account address in state keyed by it's corresponding port identifier
func (k Keeper) OnChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID string,
	counterpartyVersion string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanCloseConfirm removes the active channel stored in state
func (k Keeper) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

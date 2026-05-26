package transfer

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	capabilitytypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/transfer/keeper"
	channeltypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/04-channel/types"
	ibcexported "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

// IBCModule implements the ICS26 interface for transfer given the transfer keeper.
type IBCModule struct {
	keeper keeper.Keeper
}

// NewIBCModule creates a new IBCModule given the keeper
func NewIBCModule(k keeper.Keeper) IBCModule { _ = "STUB: not implemented"; return *new(IBCModule) }

// ValidateTransferChannelParams does validation of a newly created transfer channel. A transfer
// channel must be UNORDERED, use the correct port (by default 'transfer'), and use the current
// supported version. Only 2^32 channels are allowed to be created.
func ValidateTransferChannelParams(
	ctx sdk.Context,
	keeper keeper.Keeper,
	order channeltypes.Order,
	portID string,
	channelID string,
) error {
	_ = "STUB: not implemented"
	// NOTE: for escrow address security only 2^32 channels are allowed to be created
	// Issue: https://github.com/cosmos/cosmos-sdk/issues/7737
	return nil
}

// Require portID is the portID transfer module is bound to

// OnChanOpenInit implements the IBCModule interface
func (im IBCModule) OnChanOpenInit(
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

// Claim channel capability passed back by IBC module

// OnChanOpenTry implements the IBCModule interface.
func (im IBCModule) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Module may have already claimed capability in OnChanOpenInit in the case of crossing hellos
// (ie chainA and chainB both call ChanOpenInit before one of them calls ChanOpenTry)
// If module can already authenticate the capability then module already owns it so we don't need to claim
// Otherwise, module does not have channel capability and we must claim it from IBC

// Only claim channel capability passed back by IBC module if we do not already own it

// OnChanOpenAck implements the IBCModule interface
func (im IBCModule) OnChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID string,
	_ string,
	counterpartyVersion string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanOpenConfirm implements the IBCModule interface
func (im IBCModule) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	_ = "STUB: not implemented"

	// OnChanCloseInit implements the IBCModule interface
	return nil
}

func (im IBCModule) OnChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	_ = "STUB: not implemented"
	// Disallow user-initiated channel closing for transfer channels
	return nil
}

// OnChanCloseConfirm implements the IBCModule interface
func (im IBCModule) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	_ = "STUB: not implemented"

	// OnRecvPacket implements the IBCModule interface. A successful acknowledgement
	// is returned if the packet data is successfully decoded and the receive application
	// logic returns without error.
	return nil
}

func (im IBCModule) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	_ = "STUB: not implemented"
	return *new(ibcexported.Acknowledgement)
}

// only attempt the application logic if the packet data
// was successfully decoded

// NOTE: acknowledgement will be written synchronously during IBC handler execution.

// OnAcknowledgementPacket implements the IBCModule interface
func (im IBCModule) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	_ = "STUB: not implemented"
	return nil
}

// OnTimeoutPacket implements the IBCModule interface
func (im IBCModule) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	_ = "STUB: not implemented"
	return nil
}

// refund tokens

package mock

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	capabilitytypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/types"

	channeltypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/04-channel/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

// IBCModule implements the ICS26 callbacks for testing/mock.
type IBCModule struct {
	appModule *AppModule
	IBCApp    *MockIBCApp // base application of an IBC middleware stack
}

// NewIBCModule creates a new IBCModule given the underlying mock IBC application and scopedKeeper.
func NewIBCModule(appModule *AppModule, app *MockIBCApp) IBCModule {
	_ = "STUB: not implemented"
	return *new(IBCModule)
}

// OnChanOpenInit implements the IBCModule interface.
func (im IBCModule) OnChanOpenInit(
	ctx sdk.Context, order channeltypes.Order, connectionHops []string, portID string,
	channelID string, chanCap *capabilitytypes.Capability, counterparty channeltypes.Counterparty, version string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Claim channel capability passed back by IBC module

// OnChanOpenTry implements the IBCModule interface.
func (im IBCModule) OnChanOpenTry(
	ctx sdk.Context, order channeltypes.Order, connectionHops []string, portID string,
	channelID string, chanCap *capabilitytypes.Capability, counterparty channeltypes.Counterparty, counterpartyVersion string,
) (version string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Claim channel capability passed back by IBC module

// OnChanOpenAck implements the IBCModule interface.
func (im IBCModule) OnChanOpenAck(ctx sdk.Context, portID string, channelID string, counterpartyChannelID string, counterpartyVersion string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanOpenConfirm implements the IBCModule interface.
func (im IBCModule) OnChanOpenConfirm(ctx sdk.Context, portID, channelID string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanCloseInit implements the IBCModule interface.
func (im IBCModule) OnChanCloseInit(ctx sdk.Context, portID, channelID string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanCloseConfirm implements the IBCModule interface.
func (im IBCModule) OnChanCloseConfirm(ctx sdk.Context, portID, channelID string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnRecvPacket implements the IBCModule interface.
func (im IBCModule) OnRecvPacket(ctx sdk.Context, packet channeltypes.Packet, relayer sdk.AccAddress) exported.Acknowledgement {
	_ = "STUB: not implemented"
	return *new(exported.Acknowledgement)
}

// set state by claiming capability to check if revert happens return

// application callback called twice on same packet sequence
// must never occur

// OnAcknowledgementPacket implements the IBCModule interface.
func (im IBCModule) OnAcknowledgementPacket(ctx sdk.Context, packet channeltypes.Packet, acknowledgement []byte, relayer sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// application callback called twice on same packet sequence
// must never occur

// OnTimeoutPacket implements the IBCModule interface.
func (im IBCModule) OnTimeoutPacket(ctx sdk.Context, packet channeltypes.Packet, relayer sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// application callback called twice on same packet sequence
// must never occur

// GetMockRecvCanaryCapabilityName generates a capability name for testing OnRecvPacket functionality.
func GetMockRecvCanaryCapabilityName(packet channeltypes.Packet) string {
	_ = "STUB: not implemented"
	return ""
}

// GetMockAckCanaryCapabilityName generates a capability name for OnAcknowledgementPacket functionality.
func GetMockAckCanaryCapabilityName(packet channeltypes.Packet) string {
	_ = "STUB: not implemented"
	return ""
}

// GetMockTimeoutCanaryCapabilityName generates a capability name for OnTimeoutacket functionality.
func GetMockTimeoutCanaryCapabilityName(packet channeltypes.Packet) string {
	_ = "STUB: not implemented"
	return ""
}

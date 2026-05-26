package wasmtesting

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	capabilitytypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/types"
	channeltypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/04-channel/types"
	ibcexported "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
)

type MockChannelKeeper struct {
	GetChannelFn          func(ctx sdk.Context, srcPort, srcChan string) (channel channeltypes.Channel, found bool)
	GetNextSequenceSendFn func(ctx sdk.Context, portID, channelID string) (uint64, bool)
	SendPacketFn          func(ctx sdk.Context, channelCap *capabilitytypes.Capability, packet ibcexported.PacketI) error
	ChanCloseInitFn       func(ctx sdk.Context, portID, channelID string, chanCap *capabilitytypes.Capability) error
	GetAllChannelsFn      func(ctx sdk.Context) []channeltypes.IdentifiedChannel
	IterateChannelsFn     func(ctx sdk.Context, cb func(channeltypes.IdentifiedChannel) bool)
	SetChannelFn          func(ctx sdk.Context, portID, channelID string, channel channeltypes.Channel)
}

func (m *MockChannelKeeper) GetChannel(ctx sdk.Context, srcPort, srcChan string) (channel channeltypes.Channel, found bool) {
	_ = "STUB: not implemented"
	return *new(channeltypes.Channel), false
}

func (m *MockChannelKeeper) GetAllChannels(ctx sdk.Context) []channeltypes.IdentifiedChannel {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockChannelKeeper) GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (m *MockChannelKeeper) SendPacket(ctx sdk.Context, channelCap *capabilitytypes.Capability, packet ibcexported.PacketI) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockChannelKeeper) ChanCloseInit(ctx sdk.Context, portID, channelID string, chanCap *capabilitytypes.Capability) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockChannelKeeper) IterateChannels(ctx sdk.Context, cb func(channeltypes.IdentifiedChannel) bool) {
	_ = "STUB: not implemented"
	return
}

func (m *MockChannelKeeper) SetChannel(ctx sdk.Context, portID, channelID string, channel channeltypes.Channel) {
	_ = "STUB: not implemented"
	return
}

func MockChannelKeeperIterator(s []channeltypes.IdentifiedChannel) func(ctx sdk.Context, cb func(channeltypes.IdentifiedChannel) bool) {
	_ = "STUB: not implemented"
	return nil
}

type MockCapabilityKeeper struct {
	GetCapabilityFn          func(ctx sdk.Context, name string) (*capabilitytypes.Capability, bool)
	ClaimCapabilityFn        func(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error
	AuthenticateCapabilityFn func(ctx sdk.Context, capability *capabilitytypes.Capability, name string) bool
}

func (m MockCapabilityKeeper) GetCapability(ctx sdk.Context, name string) (*capabilitytypes.Capability, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m MockCapabilityKeeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m MockCapabilityKeeper) AuthenticateCapability(ctx sdk.Context, capability *capabilitytypes.Capability, name string) bool {
	_ = "STUB: not implemented"
	return false
}

var _ types.ICS20TransferPortSource = &MockIBCTransferKeeper{}

type MockIBCTransferKeeper struct {
	GetPortFn func(ctx sdk.Context) string
}

func (m MockIBCTransferKeeper) GetPort(ctx sdk.Context) string {
	_ = "STUB: not implemented"
	return ""
}

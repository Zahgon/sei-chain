package keeper

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	capabilitykeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/keeper"
	capabilitytypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/types"
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
	db "github.com/tendermint/tm-db"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/04-channel/types"
	porttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/05-port/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

var _ porttypes.ICS4Wrapper = Keeper{}

// Keeper defines the IBC channel keeper
type Keeper struct {
	// implements gRPC QueryServer interface
	types.QueryServer

	storeKey         sdk.StoreKey
	cdc              codec.BinaryCodec
	paramSpace       paramtypes.Subspace
	clientKeeper     types.ClientKeeper
	connectionKeeper types.ConnectionKeeper
	portKeeper       types.PortKeeper
	scopedKeeper     capabilitykeeper.ScopedKeeper
}

// NewKeeper creates a new IBC channel Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec, key sdk.StoreKey, paramSpace paramtypes.Subspace,
	clientKeeper types.ClientKeeper, connectionKeeper types.ConnectionKeeper,
	portKeeper types.PortKeeper, scopedKeeper capabilitykeeper.ScopedKeeper,
) Keeper {
	_ = "STUB: not implemented"
	return *new(Keeper)
}

// KeyOutboundEnabled is the param key for outbound enabled
var KeyOutboundEnabled = []byte("OutboundEnabled")

// KeyInboundEnabled is the param key for inbound enabled
var KeyInboundEnabled = []byte("InboundEnabled")

// IsOutboundEnabled returns true if outbound IBC is enabled.
func (k Keeper) IsOutboundEnabled(ctx sdk.Context) bool { _ = "STUB: not implemented"; return false }

// IsInboundEnabled returns true if inbound IBC is enabled.
func (k Keeper) IsInboundEnabled(ctx sdk.Context) bool { _ = "STUB: not implemented"; return false }

// GenerateChannelIdentifier returns the next channel identifier.
func (k Keeper) GenerateChannelIdentifier(ctx sdk.Context) string {
	_ = "STUB: not implemented"
	return ""
}

// GetChannel returns a channel with a particular identifier binded to a specific port
func (k Keeper) GetChannel(ctx sdk.Context, portID, channelID string) (types.Channel, bool) {
	_ = "STUB: not implemented"
	return *new(types.Channel), false
}

// SetChannel sets a channel to the store
func (k Keeper) SetChannel(ctx sdk.Context, portID, channelID string, channel types.Channel) {
	_ = "STUB: not implemented"
	return
}

// GetNextChannelSequence gets the next channel sequence from the store.
func (k Keeper) GetNextChannelSequence(ctx sdk.Context) uint64 { _ = "STUB: not implemented"; return 0 }

// SetNextChannelSequence sets the next channel sequence to the store.
func (k Keeper) SetNextChannelSequence(ctx sdk.Context, sequence uint64) {
	_ = "STUB: not implemented"
	return
}

// GetNextSequenceSend gets a channel's next send sequence from the store
func (k Keeper) GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// SetNextSequenceSend sets a channel's next send sequence to the store
func (k Keeper) SetNextSequenceSend(ctx sdk.Context, portID, channelID string, sequence uint64) {
	_ = "STUB: not implemented"
	return
}

// GetNextSequenceRecv gets a channel's next receive sequence from the store
func (k Keeper) GetNextSequenceRecv(ctx sdk.Context, portID, channelID string) (uint64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// SetNextSequenceRecv sets a channel's next receive sequence to the store
func (k Keeper) SetNextSequenceRecv(ctx sdk.Context, portID, channelID string, sequence uint64) {
	_ = "STUB: not implemented"
	return
}

// GetNextSequenceAck gets a channel's next ack sequence from the store
func (k Keeper) GetNextSequenceAck(ctx sdk.Context, portID, channelID string) (uint64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// SetNextSequenceAck sets a channel's next ack sequence to the store
func (k Keeper) SetNextSequenceAck(ctx sdk.Context, portID, channelID string, sequence uint64) {
	_ = "STUB: not implemented"
	return
}

// GetPacketReceipt gets a packet receipt from the store
func (k Keeper) GetPacketReceipt(ctx sdk.Context, portID, channelID string, sequence uint64) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// SetPacketReceipt sets an empty packet receipt to the store
func (k Keeper) SetPacketReceipt(ctx sdk.Context, portID, channelID string, sequence uint64) {
	_ = "STUB: not implemented"
	return
}

// GetPacketCommitment gets the packet commitment hash from the store
func (k Keeper) GetPacketCommitment(ctx sdk.Context, portID, channelID string, sequence uint64) []byte {
	_ = "STUB: not implemented"
	return nil
}

// HasPacketCommitment returns true if the packet commitment exists
func (k Keeper) HasPacketCommitment(ctx sdk.Context, portID, channelID string, sequence uint64) bool {
	_ = "STUB: not implemented"
	return false
}

// SetPacketCommitment sets the packet commitment hash to the store
func (k Keeper) SetPacketCommitment(ctx sdk.Context, portID, channelID string, sequence uint64, commitmentHash []byte) {
	_ = "STUB: not implemented"
	return
}

func (k Keeper) deletePacketCommitment(ctx sdk.Context, portID, channelID string, sequence uint64) {
	_ = "STUB: not implemented"
	return
}

// SetPacketAcknowledgement sets the packet ack hash to the store
func (k Keeper) SetPacketAcknowledgement(ctx sdk.Context, portID, channelID string, sequence uint64, ackHash []byte) {
	_ = "STUB: not implemented"
	return
}

// GetPacketAcknowledgement gets the packet ack hash from the store
func (k Keeper) GetPacketAcknowledgement(ctx sdk.Context, portID, channelID string, sequence uint64) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// HasPacketAcknowledgement check if the packet ack hash is already on the store
func (k Keeper) HasPacketAcknowledgement(ctx sdk.Context, portID, channelID string, sequence uint64) bool {
	_ = "STUB: not implemented"
	return false
}

// IteratePacketSequence provides an iterator over all send, receive or ack sequences.
// For each sequence, cb will be called. If the cb returns true, the iterator
// will close and stop.
func (k Keeper) IteratePacketSequence(ctx sdk.Context, iterator db.Iterator, cb func(portID, channelID string, sequence uint64) bool) {
	_ = "STUB: not implemented"
	return
}

// return if the key is not a channel key

// GetAllPacketSendSeqs returns all stored next send sequences.
func (k Keeper) GetAllPacketSendSeqs(ctx sdk.Context) (seqs []types.PacketSequence) {
	_ = "STUB: not implemented"
	return nil
}

// GetAllPacketRecvSeqs returns all stored next recv sequences.
func (k Keeper) GetAllPacketRecvSeqs(ctx sdk.Context) (seqs []types.PacketSequence) {
	_ = "STUB: not implemented"
	return nil
}

// GetAllPacketAckSeqs returns all stored next acknowledgements sequences.
func (k Keeper) GetAllPacketAckSeqs(ctx sdk.Context) (seqs []types.PacketSequence) {
	_ = "STUB: not implemented"
	return nil
}

// IteratePacketCommitment provides an iterator over all PacketCommitment objects. For each
// packet commitment, cb will be called. If the cb returns true, the iterator will close
// and stop.
func (k Keeper) IteratePacketCommitment(ctx sdk.Context, cb func(portID, channelID string, sequence uint64, hash []byte) bool) {
	_ = "STUB: not implemented"
	return
}

// GetAllPacketCommitments returns all stored PacketCommitments objects.
func (k Keeper) GetAllPacketCommitments(ctx sdk.Context) (commitments []types.PacketState) {
	_ = "STUB: not implemented"
	return nil
}

// IteratePacketCommitmentAtChannel provides an iterator over all PacketCommmitment objects
// at a specified channel. For each packet commitment, cb will be called. If the cb returns
// true, the iterator will close and stop.
func (k Keeper) IteratePacketCommitmentAtChannel(ctx sdk.Context, portID, channelID string, cb func(_, _ string, sequence uint64, hash []byte) bool) {
	_ = "STUB: not implemented"
	return
}

// GetAllPacketCommitmentsAtChannel returns all stored PacketCommitments objects for a specified
// port ID and channel ID.
func (k Keeper) GetAllPacketCommitmentsAtChannel(ctx sdk.Context, portID, channelID string) (commitments []types.PacketState) {
	_ = "STUB: not implemented"
	return nil
}

// IteratePacketReceipt provides an iterator over all PacketReceipt objects. For each
// receipt, cb will be called. If the cb returns true, the iterator will close
// and stop.
func (k Keeper) IteratePacketReceipt(ctx sdk.Context, cb func(portID, channelID string, sequence uint64, receipt []byte) bool) {
	_ = "STUB: not implemented"
	return
}

// GetAllPacketReceipts returns all stored PacketReceipt objects.
func (k Keeper) GetAllPacketReceipts(ctx sdk.Context) (receipts []types.PacketState) {
	_ = "STUB: not implemented"
	return nil
}

// IteratePacketAcknowledgement provides an iterator over all PacketAcknowledgement objects. For each
// aknowledgement, cb will be called. If the cb returns true, the iterator will close
// and stop.
func (k Keeper) IteratePacketAcknowledgement(ctx sdk.Context, cb func(portID, channelID string, sequence uint64, hash []byte) bool) {
	_ = "STUB: not implemented"
	return
}

// GetAllPacketAcks returns all stored PacketAcknowledgements objects.
func (k Keeper) GetAllPacketAcks(ctx sdk.Context) (acks []types.PacketState) {
	_ = "STUB: not implemented"
	return nil
}

// IterateChannels provides an iterator over all Channel objects. For each
// Channel, cb will be called. If the cb returns true, the iterator will close
// and stop.
func (k Keeper) IterateChannels(ctx sdk.Context, cb func(types.IdentifiedChannel) bool) {
	_ = "STUB: not implemented"
	return
}

// GetAllChannels returns all stored Channel objects.
func (k Keeper) GetAllChannels(ctx sdk.Context) (channels []types.IdentifiedChannel) {
	_ = "STUB: not implemented"
	return nil
}

// GetChannelClientState returns the associated client state with its ID, from a port and channel identifier.
func (k Keeper) GetChannelClientState(ctx sdk.Context, portID, channelID string) (string, exported.ClientState, error) {
	_ = "STUB: not implemented"
	return "", *new(exported.ClientState), nil
}

// GetConnection wraps the connection keeper's GetConnection function.
func (k Keeper) GetConnection(ctx sdk.Context, connectionID string) (exported.ConnectionI, error) {
	_ = "STUB: not implemented"
	return *new(exported.ConnectionI), nil
}

// GetChannelConnection returns the connection ID and state associated with the given port and channel identifier.
func (k Keeper) GetChannelConnection(ctx sdk.Context, portID, channelID string) (string, exported.ConnectionI, error) {
	_ = "STUB: not implemented"
	return "", *new(exported.ConnectionI), nil
}

// LookupModuleByChannel will return the IBCModule along with the capability associated with a given channel defined by its portID and channelID
func (k Keeper) LookupModuleByChannel(ctx sdk.Context, portID, channelID string) (string, *capabilitytypes.Capability, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// common functionality for IteratePacketCommitment and IteratePacketAcknowledgement
func (k Keeper) iterateHashes(_ sdk.Context, iterator db.Iterator, cb func(portID, channelID string, sequence uint64, hash []byte) bool) {
	_ = "STUB: not implemented"
	return
}

package p2p

import (
	"context"
	"testing"

	"github.com/gogo/protobuf/proto"
	gogotypes "github.com/gogo/protobuf/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// Message is a simple message containing a string-typed Value field.
type TestMessage = gogotypes.StringValue

func NodeInSlice(id types.NodeID, ids []types.NodeID) bool { _ = "STUB: not implemented"; return false }

// Network sets up an in-memory network that can be used for high-level P2P
// testing. It creates an arbitrary number of nodes that are connected to each
// other, and can open channels across all nodes with custom reactors.
type TestNetwork struct {
	nodes utils.Mutex[map[types.NodeID]*TestNode]
}

// NetworkOptions is an argument structure to parameterize the
// MakeNetwork function.
type TestNetworkOptions struct {
	NumNodes int
	NodeOpts TestNodeOptions
}

type TestNodeOptions struct {
	MaxConnected   utils.Option[int]
	PexOnHandshake bool
	SelfAddress    bool
}

func TestAddress(r *Router) NodeAddress { _ = "STUB: not implemented"; return *new(NodeAddress) }

// MakeNetwork creates a test network with the given number of nodes and
// connects them to each other.
func MakeTestNetwork(t *testing.T, opts TestNetworkOptions) *TestNetwork {
	_ = "STUB: not implemented"
	return nil
}

func (n *TestNetwork) Nodes() []*TestNode {
	_ = "STUB: not implemented"
	//nolint:prealloc
	return nil
}

func (n *TestNetwork) ConnectCycle(ctx context.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Start starts the network and connects all nodes to each other.
func (n *TestNetwork) Start(t *testing.T) {
	_ = "STUB: not implemented"

	// Populate peer managers.
	return
}

// nodes <i already connected

// NodeIDs returns the network's node IDs.
func (n *TestNetwork) NodeIDs() []types.NodeID { _ = "STUB: not implemented"; return nil }

//nolint:prealloc

// MakeChannels makes a channel on all nodes and returns them, automatically
// doing error checks and cleanups.
func TestMakeChannels[T proto.Message](
	t *testing.T,
	n *TestNetwork,
	chDesc ChannelDescriptor[T],
) map[types.NodeID]*Channel[T] {
	_ = "STUB: not implemented"
	return nil
}

// MakeChannelsNoCleanup makes a channel on all nodes and returns them,
// automatically doing error checks. The caller must ensure proper cleanup of
// all the channels.
func TestMakeChannelsNoCleanup[T proto.Message](
	t *testing.T,
	n *TestNetwork,
	chDesc ChannelDescriptor[T],
) map[types.NodeID]*Channel[T] {
	_ = "STUB: not implemented"
	return nil
}

func (n *TestNetwork) Node(id types.NodeID) *TestNode { _ = "STUB: not implemented"; return nil }

// RandomNode returns a random node.
func (n *TestNetwork) RandomNode() *TestNode { _ = "STUB: not implemented"; return nil }

// nolint:gosec

// Peers returns a node's peers (i.e. everyone except itself).
func (n *TestNetwork) Peers(id types.NodeID) []*TestNode { _ = "STUB: not implemented"; return nil }

// Remove removes a node from the network, stopping it and waiting for all other
// nodes to pick up the disconnection.
func (n *TestNetwork) Remove(t *testing.T, id types.NodeID) { _ = "STUB: not implemented"; return }

//nolint:prealloc

// Node is a node in a Network, with a Router and a PeerManager.
type TestNode struct {
	NodeID      types.NodeID
	NodeInfo    types.NodeInfo
	NodeAddress NodeAddress
	PrivKey     NodeSecretKey
	Router      *Router
}

// Waits for the specific connection to get disconnected.
func (n *TestNode) WaitForDisconnect(ctx context.Context, conn *ConnV2) {
	_ = "STUB: not implemented"
	return
}

func (n *TestNode) WaitForConns(ctx context.Context, wantPeers int) {
	_ = "STUB: not implemented"
	return
}

func (n *TestNode) WaitForConnAndGet(ctx context.Context, target types.NodeID) (conn *ConnV2) {
	_ = "STUB: not implemented"
	return nil
}

func (n *TestNode) WaitForConn(ctx context.Context, target types.NodeID, status bool) {
	_ = "STUB: not implemented"
	return
}

func (n *TestNode) Connect(ctx context.Context, target *TestNode) {
	_ = "STUB: not implemented"
	return
}

func (n *TestNode) Disconnect(ctx context.Context, target types.NodeID) {
	_ = "STUB: not implemented"
	return
}

// MakeNode creates a new Node configured for the network with a
// running peer manager, but does not add it to the existing
// network. Callers are responsible for updating peering relationships.
func (n *TestNetwork) MakeNode(t *testing.T, opts TestNodeOptions) *TestNode {
	_ = "STUB: not implemented"
	return nil
}

// Endpoint has been allocated via tcp.TestReserveAddr(), so it is NOT IP(v4/v6)Unspecified.

// MakeChannel opens a channel, with automatic error handling and cleanup. On
// test cleanup, it also checks that the channel is empty, to make sure
// all expected messages have been asserted.
func TestMakeChannel[T proto.Message](
	t *testing.T,
	n *TestNode,
	chDesc ChannelDescriptor[T],
) *Channel[T] {
	_ = "STUB: not implemented"
	return nil
}

// MakeChannelNoCleanup opens a channel, with automatic error handling. The
// caller must ensure proper cleanup of the channel.
func TestMakeChannelNoCleanup[T proto.Message](t *testing.T, n *TestNode, chDesc ChannelDescriptor[T]) *Channel[T] {
	_ = "STUB: not implemented"
	return nil
}

// MakePeerUpdates opens a peer update subscription, with automatic cleanup.
// It checks that all updates have been consumed during cleanup.
func (n *TestNode) MakePeerUpdates() *PeerUpdatesRecv { _ = "STUB: not implemented"; return nil }

func MakeTestChannelDesc(chID ChannelID) ChannelDescriptor[*TestMessage] {
	_ = "STUB: not implemented"
	return nil
}

// RequireEmpty requires that the given channel is empty.
func RequireEmpty[T proto.Message](t *testing.T, channels ...*Channel[T]) {
	_ = "STUB: not implemented"
	return
}

// RequireReceive requires that the given envelope is received on the channel.
func RequireReceive[T proto.Message](t *testing.T, channel *Channel[T], want RecvMsg[T]) {
	_ = "STUB: not implemented"
	return
}

// RequireReceiveUnordered requires that the given envelopes are all received on
// the channel, ignoring order.
func RequireReceiveUnordered[T proto.Message](t *testing.T, channel *Channel[T], want []RecvMsg[T]) {
	_ = "STUB: not implemented"
	return
}

// RequireUpdate requires that a PeerUpdates subscription yields the given update.
func RequireUpdate(t *testing.T, recv *PeerUpdatesRecv, expect PeerUpdate) {
	_ = "STUB: not implemented"
	return
}

// RequireUpdates requires that a PeerUpdates subscription yields the given updates
// in the given order.
func RequireUpdates(t *testing.T, recv *PeerUpdatesRecv, expect []PeerUpdate) {
	_ = "STUB: not implemented"
	return
}

package p2p

import (
	"context"
	"net/netip"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p/conn"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

const queueBufferDefault = 1024

type InvalidEndpointErr struct{ error }

func toChannelIDs(bytes []byte) ChannelIDSet { _ = "STUB: not implemented"; return *new(ChannelIDSet) }

type ChannelIDSet map[ChannelID]struct{}

func (cs ChannelIDSet) Contains(id ChannelID) bool { _ = "STUB: not implemented"; return false }

// Connection implements Connection for Transport.
type ConnV2 struct {
	PeerConnInfo
	sendQueue *Queue[sendMsg]
	mconn     *conn.MConnection
}

func (c *ConnV2) Info() PeerConnInfo { _ = "STUB: not implemented"; return *new(PeerConnInfo) }

func (r *Router) connSendRoutine(ctx context.Context, conn *ConnV2) error {
	_ = "STUB: not implemented"
	return nil
}

// receivePeer receives inbound messages from a peer, deserializes them and
// passes them on to the appropriate channel.
func (r *Router) connRecvRoutine(ctx context.Context, conn *ConnV2) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(gprusak): verify if this is a misbehavior, and drop the peer if it is.

// Priority is not used since all messages in this queue are from the same channel.

func (r *Router) runConn(ctx context.Context, hConn *handshakedConn, peerInfo types.NodeInfo, dialAddr utils.Option[NodeAddress]) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConnV2) String() string           { _ = "STUB: not implemented"; return "" }
func (c *ConnV2) LocalEndpoint() Endpoint  { _ = "STUB: not implemented"; return *new(Endpoint) }
func (c *ConnV2) RemoteEndpoint() Endpoint { _ = "STUB: not implemented"; return *new(Endpoint) }
func (c *ConnV2) Close()                   { _ = "STUB: not implemented"; return }

// Endpoint represents a transport connection endpoint, either local or remote.
// It is a TCP endpoint address.
type Endpoint struct{ netip.AddrPort }

// NewEndpoint constructs an Endpoint from a types.NetAddress structure.
func ResolveEndpoint(addr string) (Endpoint, error) {
	_ = "STUB: not implemented"
	return *new(Endpoint), nil
}

// NodeAddress converts the endpoint into a NodeAddress for the given node ID.
func (e Endpoint) NodeAddress(nodeID types.NodeID) NodeAddress {
	_ = "STUB: not implemented"
	return *new(NodeAddress)
}

// Validate validates the endpoint.
func (e Endpoint) Validate() error { _ = "STUB: not implemented"; return nil }

package p2p

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils/im"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("tendermint", "internal", "p2p")

type PeerConnInfo struct {
	ID               types.NodeID
	Channels         ChannelIDSet
	DialedAddr       utils.Option[NodeAddress]
	SelfDeclaredAddr utils.Option[NodeAddress]
}

func (i PeerConnInfo) connID() connID { _ = "STUB: not implemented"; return *new(connID) }

type peerConn interface {
	comparable
	Info() PeerConnInfo
	Close()
}

type connSet[C peerConn] = im.Map[connID, C]

func GetAny[C peerConn](conns connSet[C], id types.NodeID) (C, bool) {
	_ = "STUB: not implemented"
	return *new(C), false
}

func GetAll[C peerConn](cs connSet[C], id types.NodeID) []C { _ = "STUB: not implemented"; return nil }

type peerManagerInner[C peerConn] struct {
	isPersistent map[types.NodeID]bool
	conns        utils.AtomicSend[connSet[C]]
	regular      *poolManager
	persistent   *poolManager
	lastDialPool *poolManager
}

func (i *peerManagerInner[C]) poolByID(id types.NodeID) *poolManager {
	_ = "STUB: not implemented"
	return nil
}

// PeerManager manages connections and addresses of potential peers.
// PeerManager may trigger disconnects by calling conn.Close() in case of conflicting connections.
// Possible lifecycles of the connection are as follows:
// For outbound connections:
// * StartDial() -> [dialing] -> Connected(conn) -> [communicate] -> Disconnected(conn)
// * StartDial() -> [dialing] -> DialFailed(addr)
// For inbound connections:
// * Connected(conn) -> [communicate] -> Disconnected(conn)
// For adding new peer addrs, call AddAddrs().
type peerManager[C peerConn] struct {
	selfID          types.NodeID
	options         *RouterOptions
	isBlockSyncPeer map[types.NodeID]bool
	isPrivate       map[types.NodeID]bool

	inner utils.Watch[*peerManagerInner[C]]
	// Receiver of the inner.conns. It is copyable and allows accessing connections
	// without taking lock on inner.
	conns utils.AtomicRecv[connSet[C]]
}

func (p *peerManager[C]) LogState() { _ = "STUB: not implemented"; return }

func newPeerManager[C peerConn](selfID types.NodeID, options *RouterOptions) *peerManager[C] {
	_ = "STUB: not implemented"
	return nil
}

// We do not allow multiple addresses for the same peer in the peer manager any more.
// It would be backward incompatible to invalidate configs with multiple addresses per peer.
// Instead we just log an error to indicate that some addresses have been ignored.

func (m *peerManager[C]) Conns() connSet[C] { _ = "STUB: not implemented"; return nil }

// PushPex registers address list received from sender in the pex table.
// Address list replaces the previous address list received from that sender
// (every sender has a bounded capacity in peermanager).
// The addresses on the list are expected to be fresh, ideally they should be addresses
// of the current peers of the sender. This property allows us to quickly prune stale
// addresses. PeerManager keeps address list from every connected peer and a small
// "extra" cache for senders which are not connected to facilitate random local search.
// If any of the addresses is invalid (does not parse), the whole slice is rejected.
// Addresses to persistent peers are ignored, since they are populated in constructor.
func (m *peerManager[C]) PushPex(sender utils.Option[types.NodeID], addrs []NodeAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// pex data is indexed by senders which are connected peers.
// Other pex data is restricted to a small unindexed cache.
// Therefore we downgrade sender to None, if it is not a connected peer.

func (m *peerManager[C]) PushUpgradePermit() { _ = "STUB: not implemented"; return }

// StartDial waits until there is a address available for dialing.
// Returns a collection of addresses known for this peer.
// On success, it marks the peer as dialing and this peer won't be available
// for dialing until DialFailed is called.
func (m *peerManager[C]) StartDial(ctx context.Context) ([]NodeAddress, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Start with pool which has NOT dialed previously (for fairness).

// DialFailed notifies the peer manager that dialing addresses of id has failed.
func (m *peerManager[C]) DialFailed(id types.NodeID) { _ = "STUB: not implemented"; return }

// DialFailed will fail if id was not marked as dialing.

// Connected adds conn to the connections pool.
// Connected peer won't be available for dialing until disconnect (we don't need duplicate connections).
// May close and drop a duplicate connection already present in the pool.
// Returns an error if the connection should be rejected.
func (m *peerManager[C]) Connected(conn C) error { _ = "STUB: not implemented"; return nil }

// Notify the pool.

// Update the connection set.

// Check if pool requested a disconnect.

// Insert new connection.

// Disconnected removes conn from the connection pool.
// Noop if conn was not in the connection pool.
// conn.PeerInfo().NodeID peer is available for dialing again.
func (m *peerManager[C]) Disconnected(conn C) { _ = "STUB: not implemented"; return }

// It is fine to call Disconnected for conn which is not present.

// Notify pool about disconnect.
// Panic is OK, because inconsistency between conns and pool would be a bug.

// Evict closes connection to id.
func (m *peerManager[C]) Evict(id types.NodeID) { _ = "STUB: not implemented"; return }

func (m *peerManager[C]) IsBlockSyncPeer(id types.NodeID) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *peerManager[C]) Advertise() []NodeAddress { _ = "STUB: not implemented"; return nil }

// Advertise your own address.

// Prioritize dialed addresses of outbound connections.

// Fallback to self-declared addresses of inbound connections.

// All addresses in pools.
// Used by net_info endpoint, which is used by integration tests and for debugging.
func (m *peerManager[C]) AllAddrs() []NodeAddress { _ = "STUB: not implemented"; return nil }

// Infos of connections in the pool.
func (m *peerManager[C]) ConnInfos() []PeerConnInfo { _ = "STUB: not implemented"; return nil }

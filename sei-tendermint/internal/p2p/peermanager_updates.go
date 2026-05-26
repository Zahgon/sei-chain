package p2p

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// PeerUpdatesRecv.
// NOT THREAD-SAFE.
type peerUpdatesRecv[C peerConn] struct {
	recv utils.AtomicRecv[connSet[C]]
	last map[types.NodeID]struct{}
}

// PeerUpdate is a peer update event sent via PeerUpdates.
type PeerUpdate struct {
	NodeID   types.NodeID
	Status   PeerStatus
	Channels ChannelIDSet
}

func (s *peerUpdatesRecv[C]) Recv(ctx context.Context) (PeerUpdate, error) {
	_ = "STUB: not implemented"
	return *new(PeerUpdate), nil
}

// Check for disconnected peers.

// Check for connected peers.

func (m *peerManager[C]) Subscribe() *peerUpdatesRecv[C] { _ = "STUB: not implemented"; return nil }

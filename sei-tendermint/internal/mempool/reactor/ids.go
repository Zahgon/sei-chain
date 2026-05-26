package reactor

import (
	"math"
	"sync"

	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

const MaxActiveIDs = math.MaxUint16

type IDs struct {
	mtx       sync.RWMutex
	peerMap   map[types.NodeID]uint16
	nextID    uint16              // assumes that a node will never have over 65536 active peers
	activeIDs map[uint16]struct{} // used to check if a given peerID key is used
}

func NewMempoolIDs() *IDs { _ = "STUB: not implemented"; return nil }

// reserve UnknownPeerID for mempoolReactor.BroadcastTx

// ReserveForPeer searches for the next unused ID and assigns it to the provided
// peer.
func (ids *IDs) ReserveForPeer(peerID types.NodeID) { _ = "STUB: not implemented"; return }

// the peer has been reserved

// Reclaim returns the ID reserved for the peer back to unused pool.
func (ids *IDs) Reclaim(peerID types.NodeID) { _ = "STUB: not implemented"; return }

// GetForPeer returns an ID reserved for the peer.
func (ids *IDs) GetForPeer(peerID types.NodeID) uint16 { _ = "STUB: not implemented"; return 0 }

// nextPeerID returns the next unused peer ID to use. We assume that the mutex
// is already held.
func (ids *IDs) nextPeerID() uint16 { _ = "STUB: not implemented"; return 0 }

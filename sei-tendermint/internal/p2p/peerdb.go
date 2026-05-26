package p2p

import (
	"iter"
	"time"

	"github.com/google/btree"
	dbm "github.com/tendermint/tm-db"

	p2pproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/p2p"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// peerInfoFromProto converts a Protobuf PeerInfo message to a peerInfo,
// erroring if the data is invalid.
func peerDBRowFromProto(msg *p2pproto.PeerInfo) (peerDBRow, error) {
	_ = "STUB: not implemented"
	return *new(peerDBRow), nil
}

func peerDBRowFromBytes(buf []byte) (peerDBRow, error) {
	_ = "STUB: not implemented"
	return *new(peerDBRow), nil
}

// ToProto converts the peerInfo to p2pproto.PeerInfo for database storage. The
// Protobuf type only contains persisted fields, while ephemeral fields are
// discarded. The returned message may contain pointers to original data, since
// it is expected to be serialized immediately.
func (r peerDBRow) ToProto() *p2pproto.PeerInfo { _ = "STUB: not implemented"; return nil }

// Database key prefixes.
const (
	prefixPeerInfo int64 = 1
)

// keyPeerInfo generates a peerInfo database key.
func keyPeerInfo(id types.NodeID) []byte { _ = "STUB: not implemented"; return nil }

// keyPeerInfoRange generates start/end keys for the entire peerInfo key range.
func keyPeerInfoRange() ([]byte, []byte) { _ = "STUB: not implemented"; return nil, nil }

type peerDBRow struct {
	LastConnected time.Time
	Addr          NodeAddress
}

func (a peerDBRow) Compare(b peerDBRow) int { _ = "STUB: not implemented"; return 0 }

type peerDB struct {
	db              dbm.DB
	maxRows         int
	byNodeID        map[types.NodeID]peerDBRow
	byLastConnected *btree.BTreeG[peerDBRow]
}

func newPeerDB(db dbm.DB, maxRows int) (*peerDB, error) { _ = "STUB: not implemented"; return nil, nil }

// Prune invalid data.

func (db *peerDB) Close() { _ = "STUB: not implemented"; return }

func (db *peerDB) All() iter.Seq[NodeAddress] { _ = "STUB: not implemented"; return nil }

func (db *peerDB) Insert(addr NodeAddress, lastConnected time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *peerDB) truncate() error { _ = "STUB: not implemented"; return nil }

package types

import (
	"encoding/json"

	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto"
)

//------------------------------------------------------------------------------
// Persistent peer ID
// TODO: encrypt on disk

// NodeKey is the persistent peer key.
// It contains the nodes private key for authentication.
type NodeKey crypto.PrivKey

type nodeKeyJSON struct {
	ID      NodeID          `json:"id"`
	PrivKey json.RawMessage `json:"priv_key"`
}

func (nk NodeKey) ID() NodeID { _ = "STUB: not implemented"; return *new(NodeID) }

func (nk NodeKey) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (nk *NodeKey) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// PubKey returns the peer's PubKey
func (nk NodeKey) PubKey() crypto.PubKey { _ = "STUB: not implemented"; return *new(crypto.PubKey) }

// SaveAs persists the NodeKey to filePath.
// It also writes a node_pubkey.txt file in the same directory containing the
// public key in "node:ed25519:public:<hex>" format for use in autobahn config generation.
func (nk NodeKey) SaveAs(filePath string) error { _ = "STUB: not implemented"; return nil }

// Write pubkey in autobahn-compatible format alongside the key file.
// TODO: use p2p.NodePublicKey.String() directly to avoid duplicating the "node:" prefix.

// LoadOrGenNodeKey attempts to load the NodeKey from the given filePath. If
// the file does not exist, it generates and saves a new NodeKey.
func LoadOrGenNodeKey(filePath string) (NodeKey, error) {
	_ = "STUB: not implemented"
	return *new(NodeKey), nil
}

// GenNodeKey generates a new node key.
func GenNodeKey() NodeKey { _ = "STUB: not implemented"; return *new(NodeKey) }

// LoadNodeKey loads NodeKey located in filePath.
func LoadNodeKey(filePath string) (NodeKey, error) {
	_ = "STUB: not implemented"
	return *new(NodeKey), nil
}

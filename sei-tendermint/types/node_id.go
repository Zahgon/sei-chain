package types

import (
	"regexp"

	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto"
)

// NodeIDByteLength is the length of a crypto.Address. Currently only 20.
// FIXME: support other length addresses?
const NodeIDByteLength = crypto.AddressSize

// reNodeID is a regexp for valid node IDs.
var reNodeID = regexp.MustCompile(`^[0-9a-f]{40}$`)

// NodeID is a hex-encoded crypto.Address. It must be lowercased
// (for uniqueness) and of length 2*NodeIDByteLength.
type NodeID string

// NewNodeID returns a lowercased (normalized) NodeID, or errors if the
// node ID is invalid.
func NewNodeID(nodeID string) (NodeID, error) { _ = "STUB: not implemented"; return *new(NodeID), nil }

// IDAddressString returns id@hostPort. It strips the leading
// protocol from protocolHostPort if it exists.
func (id NodeID) AddressString(protocolHostPort string) string {
	_ = "STUB: not implemented"
	return ""
}

// NodeIDFromPubKey creates a node ID from a given PubKey address.
func NodeIDFromPubKey(pubKey crypto.PubKey) NodeID { _ = "STUB: not implemented"; return *new(NodeID) }

// Bytes converts the node ID to its binary byte representation.
func (id NodeID) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Validate validates the NodeID.
func (id NodeID) Validate() error { _ = "STUB: not implemented"; return nil }

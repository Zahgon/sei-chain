package factory

import (
	"testing"

	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// NodeID returns a valid NodeID based on an inputted string
func NodeID(t *testing.T, str string) types.NodeID {
	_ = "STUB: not implemented"
	return *new(types.NodeID)
}

// RandomNodeID returns a randomly generated valid NodeID
func RandomNodeID(t *testing.T) types.NodeID { _ = "STUB: not implemented"; return *new(types.NodeID) }

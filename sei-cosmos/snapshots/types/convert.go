package types

import (
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
)

// Converts an ABCI snapshot to a snapshot. Mainly to decode the SDK metadata.
func SnapshotFromABCI(in *abci.Snapshot) (Snapshot, error) {
	_ = "STUB: not implemented"
	return *new(Snapshot), nil
}

// Converts a Snapshot to its ABCI representation. Mainly to encode the SDK metadata.
func (s Snapshot) ToABCI() (abci.Snapshot, error) {
	_ = "STUB: not implemented"
	return *new(abci.Snapshot), nil
}

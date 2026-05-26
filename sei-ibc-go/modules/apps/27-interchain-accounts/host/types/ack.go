package types

import (
	channeltypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/04-channel/types"
)

const (
	// ackErrorString defines a string constant included in error acknowledgements
	// NOTE: Changing this const is state machine breaking as acknowledgements are written into state
	ackErrorString = "error handling packet on host chain: see events for details"
)

// NewErrorAcknowledgement returns a deterministic error string which may be used in
// the packet acknowledgement.
func NewErrorAcknowledgement(err error) channeltypes.Acknowledgement {
	_ = "STUB: not implemented"
	// the ABCI code is included in the abcitypes.ResponseDeliverTx hash
	// constructed in Tendermint and is therefore determinstic
	return *new(channeltypes.Acknowledgement)
}

// discard non-deterministic codespace and log values

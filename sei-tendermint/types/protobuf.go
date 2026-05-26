package types

import (
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
)

//-------------------------------------------------------

// TM2PB is used for converting Tendermint ABCI to protobuf ABCI.
// UNSTABLE
var TM2PB = tm2pb{}

type tm2pb struct{}

func (tm2pb) Validator(val *Validator) abci.Validator {
	_ = "STUB: not implemented"
	return *new(abci.Validator)
}

func (tm2pb) ValidatorUpdate(val *Validator) abci.ValidatorUpdate {
	_ = "STUB: not implemented"
	return *new(abci.ValidatorUpdate)
}

func (tm2pb) ValidatorUpdates(vals *ValidatorSet) []abci.ValidatorUpdate {
	_ = "STUB: not implemented"
	return nil
}

//----------------------------------------------------------------------------

// PB2TM is used for converting protobuf ABCI to Tendermint ABCI.
// UNSTABLE
var PB2TM = pb2tm{}

type pb2tm struct{}

func (pb2tm) ValidatorUpdates(vals []abci.ValidatorUpdate) ([]*Validator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

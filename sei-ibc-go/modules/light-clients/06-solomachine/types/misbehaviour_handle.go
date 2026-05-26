package types

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

// CheckMisbehaviourAndUpdateState determines whether or not the currently registered
// public key signed over two different messages with the same sequence. If this is true
// the client state is updated to a frozen status.
// NOTE: Misbehaviour is not tracked for previous public keys, a solo machine may update to
// a new public key before the misbehaviour is processed. Therefore, misbehaviour is data
// order processing dependent.
func (cs ClientState) CheckMisbehaviourAndUpdateState(
	ctx sdk.Context,
	cdc codec.BinaryCodec,
	clientStore sdk.KVStore,
	misbehaviour exported.Misbehaviour,
) (exported.ClientState, error) {
	_ = "STUB: not implemented"
	return *new(exported.ClientState), nil
}

// NOTE: a check that the misbehaviour message data are not equal is done by
// misbehaviour.ValidateBasic which is called by the 02-client keeper.

// verify first signature

// verify second signature

// verifySignatureAndData verifies that the currently registered public key has signed
// over the provided data and that the data is valid. The data is valid if it can be
// unmarshaled into the specified data type.
func verifySignatureAndData(cdc codec.BinaryCodec, clientState ClientState, misbehaviour *Misbehaviour, sigAndData *SignatureAndData) error {
	_ = "STUB: not implemented"
	// do not check misbehaviour timestamp since we want to allow processing of past misbehaviour
	return nil
}

// ensure data can be unmarshaled to the specified data type

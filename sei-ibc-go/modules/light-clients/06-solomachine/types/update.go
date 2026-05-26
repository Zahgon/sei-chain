package types

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

// CheckHeaderAndUpdateState checks if the provided header is valid and updates
// the consensus state if appropriate. It returns an error if:
// - the header provided is not parseable to a solo machine header
// - the header sequence does not match the current sequence
// - the header timestamp is less than the consensus state timestamp
// - the currently registered public key did not provide the update signature
func (cs ClientState) CheckHeaderAndUpdateState(
	ctx sdk.Context, cdc codec.BinaryCodec, clientStore sdk.KVStore,
	header exported.Header,
) (exported.ClientState, exported.ConsensusState, error) {
	_ = "STUB: not implemented"
	return *new(exported.ClientState), *new(exported.ConsensusState), nil
}

// checkHeader checks if the Solo Machine update signature is valid.
func checkHeader(cdc codec.BinaryCodec, clientState *ClientState, header *Header) error {
	_ = "STUB: not implemented"
	// assert update sequence is current sequence
	return nil
}

// assert update timestamp is not less than current consensus state timestamp

// assert currently registered public key signed over the new public key with correct sequence

// update the consensus state to the new public key and an incremented sequence
func update(clientState *ClientState, header *Header) (*ClientState, *ConsensusState) {
	_ = "STUB: not implemented"
	return nil, nil
}

// increment sequence number

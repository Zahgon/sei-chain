package types

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

// CheckSubstituteAndUpdateState verifies that the subject is allowed to be updated by
// a governance proposal and that the substitute client is a solo machine.
// It will update the consensus state to the substitute's consensus state and
// the sequence to the substitute's current sequence. An error is returned if
// the client has been disallowed to be updated by a governance proposal,
// the substitute is not a solo machine, or the current public key equals
// the new public key.
func (cs ClientState) CheckSubstituteAndUpdateState(
	ctx sdk.Context, cdc codec.BinaryCodec, subjectClientStore,
	_ sdk.KVStore, substituteClient exported.ClientState,
) (exported.ClientState, error) {
	_ = "STUB: not implemented"
	return *new(exported.ClientState), nil
}

// update to substitute parameters

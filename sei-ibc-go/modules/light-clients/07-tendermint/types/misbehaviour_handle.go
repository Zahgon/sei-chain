package types

import (
	"time"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

// CheckMisbehaviourAndUpdateState determines whether or not two conflicting
// headers at the same height would have convinced the light client.
//
// NOTE: consensusState1 is the trusted consensus state that corresponds to the TrustedHeight
// of misbehaviour.Header1
// Similarly, consensusState2 is the trusted consensus state that corresponds
// to misbehaviour.Header2
// Misbehaviour sets frozen height to {0, 1} since it is only used as a boolean value (zero or non-zero).
func (cs ClientState) CheckMisbehaviourAndUpdateState(
	ctx sdk.Context,
	cdc codec.BinaryCodec,
	clientStore sdk.KVStore,
	misbehaviour exported.Misbehaviour,
) (exported.ClientState, error) {
	_ = "STUB: not implemented"
	return *new(exported.ClientState), nil
}

// The status of the client is checked in 02-client

// if heights are equal check that this is valid misbehaviour of a fork
// otherwise if heights are unequal check that this is valid misbehavior of BFT time violation

// Ensure that Commit Hashes are different

// Header1 is at greater height than Header2, therefore Header1 time must be less than or equal to
// Header2 time in order to be valid misbehaviour (violation of monotonic time).

// Regardless of the type of misbehaviour, ensure that both headers are valid and would have been accepted by light-client

// Retrieve trusted consensus states for each Header in misbehaviour
// and unmarshal from clientStore

// Get consensus bytes from clientStore

// Get consensus bytes from clientStore

// Check the validity of the two conflicting headers against their respective
// trusted consensus states
// NOTE: header height and commitment root assertions are checked in
// misbehaviour.ValidateBasic by the client keeper and msg.ValidateBasic
// by the base application.

// checkMisbehaviourHeader checks that a Header in Misbehaviour is valid misbehaviour given
// a trusted ConsensusState
func checkMisbehaviourHeader(
	clientState *ClientState, consState *ConsensusState, header *Header, currentTimestamp time.Time,
) error {
	_ = "STUB: not implemented"
	return nil
}

// check the trusted fields for the header against ConsensusState

// assert that the age of the trusted consensus state is not older than the trusting period

// If chainID is in revision format, then set revision number of chainID with the revision number
// of the misbehaviour header

// - ValidatorSet must have TrustLevel similarity with trusted FromValidatorSet
// - ValidatorSets on both headers are valid given the last trusted ValidatorSet

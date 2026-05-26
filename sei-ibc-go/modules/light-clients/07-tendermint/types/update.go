package types

import (
	"time"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

// CheckHeaderAndUpdateState checks if the provided header is valid, and if valid it will:
// create the consensus state for the header.Height
// and update the client state if the header height is greater than the latest client state height
// It returns an error if:
// - the client or header provided are not parseable to tendermint types
// - the header is invalid
// - header height is less than or equal to the trusted header height
// - header revision is not equal to trusted header revision
// - header valset commit verification fails
// - header timestamp is past the trusting period in relation to the consensus state
// - header timestamp is less than or equal to the consensus state timestamp
//
// UpdateClient may be used to either create a consensus state for:
// - a future height greater than the latest client state height
// - a past height that was skipped during bisection
// If we are updating to a past height, a consensus state is created for that height to be persisted in client store
// If we are updating to a future height, the consensus state is created and the client state is updated to reflect
// the new latest height
// UpdateClient must only be used to update within a single revision, thus header revision number and trusted height's revision
// number must be the same. To update to a new revision, use a separate upgrade path
// Tendermint client validity checking uses the bisection algorithm described
// in the [Tendermint spec](https://github.com/tendermint/spec/blob/master/spec/consensus/light-client.md).
//
// Misbehaviour Detection:
// UpdateClient will detect implicit misbehaviour by enforcing certain invariants on any new update call and will return a frozen client.
// 1. Any valid update that creates a different consensus state for an already existing height is evidence of misbehaviour and will freeze client.
// 2. Any valid update that breaks time monotonicity with respect to its neighboring consensus states is evidence of misbehaviour and will freeze client.
// Misbehaviour sets frozen height to {0, 1} since it is only used as a boolean value (zero or non-zero).
//
// Pruning:
// UpdateClient will additionally retrieve the earliest consensus state for this clientID and check if it is expired. If it is,
// that consensus state will be pruned from store along with all associated metadata. This will prevent the client store from
// becoming bloated with expired consensus states that can no longer be used for updates and packet verification.
func (cs ClientState) CheckHeaderAndUpdateState(
	ctx sdk.Context, cdc codec.BinaryCodec, clientStore sdk.KVStore,
	header exported.Header,
) (exported.ClientState, exported.ConsensusState, error) {
	_ = "STUB: not implemented"
	return *new(exported.ClientState), *new(exported.ConsensusState), nil
}

// Check if the Client store already has a consensus state for the header's height
// If the consensus state exists, and it matches the header then we return early
// since header has already been submitted in a previous UpdateClient.

// This header has already been submitted and the necessary state is already stored
// in client store, thus we can return early without further validation.

// A consensus state already exists for this height, but it does not match the provided header.
// Thus, we must check that this header is valid, and if so we will freeze the client.

// get consensus state from clientStore

// Header is different from existing consensus state and also valid, so freeze the client and return

// Check that consensus state timestamps are monotonic

// if previous consensus state exists, check consensus state time is greater than previous consensus state time
// if previous consensus state is not before current consensus state, freeze the client and return.

// if next consensus state exists, check consensus state time is less than next consensus state time
// if next consensus state is not after current consensus state, freeze the client and return.

// Check the earliest consensus state to see if it is expired, if so then set the prune height
// so that we can delete consensus state and all associated metadata.

// this error should never occur

// if pruneHeight is set, delete consensus state and metadata

// checkTrustedHeader checks that consensus state matches trusted fields of Header
func checkTrustedHeader(header *Header, consState *ConsensusState) error {
	_ = "STUB: not implemented"
	return nil
}

// assert that trustedVals is NextValidators of last trusted header
// to do this, we check that trustedVals.Hash() == consState.NextValidatorsHash

// checkValidity checks if the Tendermint header is valid.
// CONTRACT: consState.Height == header.TrustedHeight
func checkValidity(
	clientState *ClientState, consState *ConsensusState,
	header *Header, currentTimestamp time.Time,
) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateClient only accepts updates with a header at the same revision
// as the trusted consensus state

// assert header height is newer than consensus state

// If chainID is in revision format, then set revision number of chainID with the revision number
// of the header we are verifying
// This is useful if the update is at a previous revision rather than an update to the latest revision
// of the client.
// The chainID must be set correctly for the previous revision before attempting verification.
// Updates for previous revisions are not supported if the chainID is not in revision format.

// Construct a trusted header using the fields in consensus state
// Only Height, Time, and NextValidatorsHash are necessary for verification

// #nosec G115 --- checked above

// Verify next header with the passed-in trustedVals
// - asserts trusting period not passed
// - assert header timestamp is not past the trusting period
// - assert header timestamp is past latest stored consensus state timestamp
// - assert that a TrustLevel proportion of TrustedValidators signed new Commit

// update the consensus state from a new header and set processed time metadata
func update(ctx sdk.Context, clientStore sdk.KVStore, clientState *ClientState, header *Header) (*ClientState, *ConsensusState) {
	_ = "STUB: not implemented"
	return nil, nil
}

// set metadata for this consensus state

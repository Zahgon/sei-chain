package types

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

// CheckSubstituteAndUpdateState will try to update the client with the state of the
// substitute.
//
// AllowUpdateAfterMisbehaviour and AllowUpdateAfterExpiry have been deprecated.
// Please see ADR 026 for more information.
//
// The following must always be true:
//   - The substitute client is the same type as the subject client
//   - The subject and substitute client states match in all parameters (expect frozen height, latest height, and chain-id)
//
// In case 1) before updating the client, the client will be unfrozen by resetting
// the FrozenHeight to the zero Height.
func (cs ClientState) CheckSubstituteAndUpdateState(
	ctx sdk.Context, cdc codec.BinaryCodec, subjectClientStore,
	substituteClientStore sdk.KVStore, substituteClient exported.ClientState,
) (exported.ClientState, error) {
	_ = "STUB: not implemented"
	return *new(exported.ClientState), nil
}

// unfreeze the client

// copy consensus states and processed time from substitute to subject
// starting from initial height and ending on the latest height (inclusive)

// set metadata stored for the substitute consensus state

// set new trusting period based on the substitute client state

// no validation is necessary since the substitute is verified to be Active
// in 02-client.

// IsMatchingClientState returns true if all the client state parameters match
// except for frozen height, latest height, trusting period, chain-id.
func IsMatchingClientState(subject, substitute ClientState) bool {
	_ = "STUB: not implemented"
	// zero out parameters which do not need to match
	return false
}

// sets both sets of flags to true as these flags have been DEPRECATED, see ADR-026 for more information

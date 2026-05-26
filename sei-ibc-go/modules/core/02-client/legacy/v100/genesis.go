package v100

import (
	"time"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

// MigrateGenesis accepts exported v1.0.0 IBC client genesis file and migrates it to:
//
// - Update solo machine client state protobuf definition (v1 to v2)
// - Remove all solo machine consensus states
// - Remove all expired tendermint consensus states
// - Adds ProcessedHeight and Iteration keys for unexpired tendermint consensus states
func MigrateGenesis(cdc codec.BinaryCodec, clientGenState *types.GenesisState, genesisBlockTime time.Time, selfHeight exported.Height) (*types.GenesisState, error) {
	_ = "STUB: not implemented"
	// To prune the consensus states, we will create new clientsConsensus
	// and clientsMetadata. These slices will be filled up with consensus states
	// which should not be pruned. No solo machine consensus states should be added
	// and only unexpired consensus states for tendermint clients will be added.
	// The metadata keys for unexpired consensus states will be added to clientsMetadata
	return nil, nil
}

// update solo machine client state defintions

// iterate consensus states by client

// look for consensus states for the current client

// remove all consensus states for the solo machine
// do not add to new clientsConsensus

// only add non expired consensus states to new clientsConsensus

// collect unexpired consensus states

// if we found at least one unexpired consensus state, create a clientConsensusState
// and add it to clientsConsensus

// collect metadata for unexpired consensus states

// remove all expired tendermint consensus state metadata by adding only
// unexpired consensus state metadata

// look for metadata for current client

// obtain height for consensus state being pruned

// iterate through metadata and find metadata for current unexpired height
// only unexpired consensus state metadata should be added

// the previous version of IBC only contained the processed time metadata
// if we find the processed time metadata for an unexpired height, add the
// iteration key and processed height keys.

// set the processed height using the current self height
// this is safe, it may cause delays in packet processing if there
// is a non zero connection delay time

// processed time

// if we have metadata for unexipred consensus states, add it to consensusMetadata

package v100

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	smtypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/light-clients/06-solomachine/types"
	ibctmtypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/light-clients/07-tendermint/types"
)

// MigrateStore performs in-place store migrations from SDK v0.40 of the IBC module to v1.0.0 of ibc-go.
// The migration includes:
//
// - Migrating solo machine client states from v1 to v2 protobuf definition
// - Pruning all solo machine consensus states
// - Pruning expired tendermint consensus states
// - Adds ProcessedHeight and Iteration keys for unexpired tendermint consensus states
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, cdc codec.BinaryCodec) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// collect all clients

// key is clients/{clientid}/clientState
// Thus, keySplit[1] is clientID

// update solomachine in store

// add iteration keys so pruning will be successful

// migrateSolomachine migrates the solomachine from v1 to v2 solo machine protobuf definition.
func migrateSolomachine(clientState *ClientState) *smtypes.ClientState {
	_ = "STUB: not implemented"
	return nil
}

// pruneSolomachineConsensusStates removes all solomachine consensus states from the
// client store.
func pruneSolomachineConsensusStates(clientStore sdk.KVStore) { _ = "STUB: not implemented"; return }

// key is in the format "consensusStates/<height>"

// collect consensus states to be pruned

// delete all consensus states

// addConsensusMetadata adds the iteration key and processed height for all tendermint consensus states
// These keys were not included in the previous release of the IBC module. Adding the iteration keys allows
// for pruning iteration.
func addConsensusMetadata(ctx sdk.Context, clientStore sdk.KVStore, cdc codec.BinaryCodec, clientState *ibctmtypes.ClientState) error {
	_ = "STUB: not implemented"
	return nil
}

// consensus key is in the format "consensusStates/<height>"

// set the iteration key and processed height
// these keys were not included in the SDK v0.42.0 release

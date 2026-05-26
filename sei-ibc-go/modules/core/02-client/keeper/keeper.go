package keeper

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
	upgradetypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

// KeyInboundEnabled is the param key for inbound enabled
var KeyInboundEnabled = []byte("InboundEnabled")

// Keeper represents a type that grants read and write permissions to any client
// state information
type Keeper struct {
	storeKey      sdk.StoreKey
	cdc           codec.BinaryCodec
	paramSpace    paramtypes.Subspace
	stakingKeeper types.StakingKeeper
	upgradeKeeper types.UpgradeKeeper
}

// NewKeeper creates a new NewKeeper instance
func NewKeeper(cdc codec.BinaryCodec, key sdk.StoreKey, paramSpace paramtypes.Subspace, sk types.StakingKeeper, uk types.UpgradeKeeper) Keeper {
	_ = "STUB: not implemented"
	// set KeyTable if it has not already been set
	return *new(Keeper)
}

// IsInboundEnabled returns true if inbound IBC is enabled.
func (k Keeper) IsInboundEnabled(ctx sdk.Context) bool { _ = "STUB: not implemented"; return false }

// GenerateClientIdentifier returns the next client identifier.
func (k Keeper) GenerateClientIdentifier(ctx sdk.Context, clientType string) string {
	_ = "STUB: not implemented"
	return ""
}

// GetClientState gets a particular client from the store
func (k Keeper) GetClientState(ctx sdk.Context, clientID string) (exported.ClientState, bool) {
	_ = "STUB: not implemented"
	return *new(exported.ClientState), false
}

// SetClientState sets a particular Client to the store
func (k Keeper) SetClientState(ctx sdk.Context, clientID string, clientState exported.ClientState) {
	_ = "STUB: not implemented"
	return
}

// GetClientConsensusState gets the stored consensus state from a client at a given height.
func (k Keeper) GetClientConsensusState(ctx sdk.Context, clientID string, height exported.Height) (exported.ConsensusState, bool) {
	_ = "STUB: not implemented"
	return *new(exported.ConsensusState), false
}

// SetClientConsensusState sets a ConsensusState to a particular client at the given
// height
func (k Keeper) SetClientConsensusState(ctx sdk.Context, clientID string, height exported.Height, consensusState exported.ConsensusState) {
	_ = "STUB: not implemented"
	return
}

// GetNextClientSequence gets the next client sequence from the store.
func (k Keeper) GetNextClientSequence(ctx sdk.Context) uint64 { _ = "STUB: not implemented"; return 0 }

// SetNextClientSequence sets the next client sequence to the store.
func (k Keeper) SetNextClientSequence(ctx sdk.Context, sequence uint64) {
	_ = "STUB: not implemented"
	return
}

// IterateConsensusStates provides an iterator over all stored consensus states.
// objects. For each State object, cb will be called. If the cb returns true,
// the iterator will close and stop.
func (k Keeper) IterateConsensusStates(ctx sdk.Context, cb func(clientID string, cs types.ConsensusStateWithHeight) bool) {
	_ = "STUB: not implemented"
	return
}

// consensus key is in the format "clients/<clientID>/consensusStates/<height>"

// GetAllGenesisClients returns all the clients in state with their client ids returned as IdentifiedClientState
func (k Keeper) GetAllGenesisClients(ctx sdk.Context) types.IdentifiedClientStates {
	_ = "STUB: not implemented"
	return *new(types.IdentifiedClientStates)
}

// GetAllClientMetadata will take a list of IdentifiedClientState and return a list
// of IdentifiedGenesisMetadata necessary for exporting and importing client metadata
// into the client store.
func (k Keeper) GetAllClientMetadata(ctx sdk.Context, genClients []types.IdentifiedClientState) ([]types.IdentifiedGenesisMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetAllClientMetadata takes a list of IdentifiedGenesisMetadata and stores all of the metadata in the client store at the appropriate paths.
func (k Keeper) SetAllClientMetadata(ctx sdk.Context, genMetadata []types.IdentifiedGenesisMetadata) {
	_ = "STUB: not implemented"
	return
}

// create client store

// set all metadata kv pairs in client store

// GetAllConsensusStates returns all stored client consensus states.
func (k Keeper) GetAllConsensusStates(ctx sdk.Context) types.ClientsConsensusStates {
	_ = "STUB: not implemented"
	return *new(types.ClientsConsensusStates)
}

// HasClientConsensusState returns if keeper has a ConsensusState for a particular
// client at the given height
func (k Keeper) HasClientConsensusState(ctx sdk.Context, clientID string, height exported.Height) bool {
	_ = "STUB: not implemented"
	return false
}

// GetLatestClientConsensusState gets the latest ConsensusState stored for a given client
func (k Keeper) GetLatestClientConsensusState(ctx sdk.Context, clientID string) (exported.ConsensusState, bool) {
	_ = "STUB: not implemented"
	return *new(exported.ConsensusState), false
}

// GetSelfConsensusState introspects the (self) past historical info at a given height
// and returns the expected consensus state at that height.
// For now, can only retrieve self consensus states for the current revision
func (k Keeper) GetSelfConsensusState(ctx sdk.Context, height exported.Height) (exported.ConsensusState, error) {
	_ = "STUB: not implemented"
	return *new(exported.ConsensusState), nil
}

// check that height revision matches chainID revision

// #nosec G115 -- revision height is bounds checked above

// ValidateSelfClient validates the client parameters for a client of the running chain
// This function is only used to validate the client state the counterparty stores for this chain
// Client must be in same revision as the executing chain
func (k Keeper) ValidateSelfClient(ctx sdk.Context, clientState exported.ClientState) error {
	_ = "STUB: not implemented"
	return nil
}

// client must be in the same revision as executing chain

// #nosec G115 -- block height is checked above to be non-negative

// For now, SDK IBC implementation assumes that upgrade path (if defined) is defined by SDK upgrade module

// GetUpgradePlan executes the upgrade keeper GetUpgradePlan function.
func (k Keeper) GetUpgradePlan(ctx sdk.Context) (plan upgradetypes.Plan, havePlan bool) {
	_ = "STUB: not implemented"
	return *new(upgradetypes.Plan), false
}

// GetUpgradedClient executes the upgrade keeper GetUpgradeClient function.
func (k Keeper) GetUpgradedClient(ctx sdk.Context, planHeight int64) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// GetUpgradedConsensusState returns the upgraded consensus state
func (k Keeper) GetUpgradedConsensusState(ctx sdk.Context, planHeight int64) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// SetUpgradedConsensusState executes the upgrade keeper SetUpgradedConsensusState function.
func (k Keeper) SetUpgradedConsensusState(ctx sdk.Context, planHeight int64, bz []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// IterateClients provides an iterator over all stored light client State
// objects. For each State object, cb will be called. If the cb returns true,
// the iterator will close and stop.
func (k Keeper) IterateClients(ctx sdk.Context, cb func(clientID string, cs exported.ClientState) bool) {
	_ = "STUB: not implemented"
	return
}

// key is ibc/{clientid}/clientState
// Thus, keySplit[1] is clientID

// GetAllClients returns all stored light client State objects.
func (k Keeper) GetAllClients(ctx sdk.Context) (states []exported.ClientState) {
	_ = "STUB: not implemented"
	return nil
}

// ClientStore returns isolated prefix store for each client so they can read/write in separate
// namespace without being able to read/write other client's data
func (k Keeper) ClientStore(ctx sdk.Context, clientID string) sdk.KVStore {
	_ = "STUB: not implemented"
	return *new(sdk.KVStore)
}

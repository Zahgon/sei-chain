package state

import (
	"github.com/google/orderedcode"
	dbm "github.com/tendermint/tm-db"

	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	tmstate "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/state"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

const (
	// persist validators every valSetCheckpointInterval blocks to avoid
	// LoadValidators taking too much time.
	// https://github.com/tendermint/tendermint/pull/3438
	// 100000 results in ~ 100ms to get 100 validators (see BenchmarkLoadValidators)
	valSetCheckpointInterval = 100000
)

//------------------------------------------------------------------------

// key prefixes
// NB: Before modifying these, cross-check them with those in
// * internal/store/store.go    [0..4, 13]
// * internal/state/store.go    [5..8, 14]
// * internal/evidence/pool.go  [9..10]
// * light/store/db/db.go       [11..12]
// TODO(thane): Move all these to their own package.
// TODO: what about these (they already collide):
// * scripts/scmigrate/migrate.go [3]
// * internal/p2p/peermanager.go  [1]
const (
	// prefixes are unique across all tm db's
	prefixValidators             = int64(5)
	prefixConsensusParams        = int64(6)
	prefixABCIResponses          = int64(7) // deprecated in v0.36
	prefixState                  = int64(8)
	prefixFinalizeBlockResponses = int64(14)
)

func encodeKey(prefix int64, height int64) []byte { _ = "STUB: not implemented"; return nil }

func validatorsKey(height int64) []byte { _ = "STUB: not implemented"; return nil }

func consensusParamsKey(height int64) []byte { _ = "STUB: not implemented"; return nil }

func abciResponsesKey(height int64) []byte { _ = "STUB: not implemented"; return nil }

func finalizeBlockResponsesKey(height int64) []byte { _ = "STUB: not implemented"; return nil }

// stateKey should never change after being set in init()
var stateKey []byte

func init() {
	var err error
	stateKey, err = orderedcode.Append(nil, prefixState)
	if err != nil {
		panic(err)
	}
}

//----------------------

//go:generate ../../scripts/mockery_generate.sh Store

// Store defines the state store interface
//
// It is used to retrieve current state and save and load ABCI responses,
// validators and consensus parameters
type Store interface {
	// Load loads the current state of the blockchain
	Load() (State, error)
	// LoadValidators loads the validator set at a given height
	LoadValidators(int64) (*types.ValidatorSet, error)
	// LoadFinalizeBlockResponses loads the responses to FinalizeBlock for a given height
	LoadFinalizeBlockResponses(int64) (*abci.ResponseFinalizeBlock, error)
	// LoadConsensusParams loads the consensus params for a given height
	LoadConsensusParams(int64) (types.ConsensusParams, error)
	// Save overwrites the previous state with the updated one
	Save(State) error
	// SaveFinalizeBlockResponses saves responses to FinalizeBlock for a given height
	SaveFinalizeBlockResponses(int64, *abci.ResponseFinalizeBlock) error
	// SaveValidatorSet saves the validator set at a given height
	SaveValidatorSets(int64, int64, *types.ValidatorSet) error
	// Bootstrap is used for bootstrapping state when not starting from a initial height.
	Bootstrap(State) error
	// PruneStates takes the height from which to prune up to (exclusive)
	PruneStates(int64) error
	// Close closes the connection with the database
	Close() error
}

// dbStore wraps a db (github.com/tendermint/tm-db)
type dbStore struct {
	db dbm.DB
}

var _ Store = (*dbStore)(nil)

// NewStore creates the dbStore of the state pkg.
func NewStore(db dbm.DB) Store {
	_ = "STUB: not implemented"

	// LoadState loads the State from the database.
	return *new(Store)
}

func (store dbStore) Load() (State, error) { _ = "STUB: not implemented"; return *new(State), nil }

func (store dbStore) loadState(key []byte) (state State, err error) {
	_ = "STUB: not implemented"
	return *new(State), nil
}

// DATA HAS BEEN CORRUPTED OR THE SPEC HAS CHANGED

// Save persists the State, the ValidatorsInfo, and the ConsensusParamsInfo to the database.
// This flushes the writes (e.g. calls SetSync).
func (store dbStore) Save(state State) error { _ = "STUB: not implemented"; return nil }

func (store dbStore) save(state State, key []byte) error { _ = "STUB: not implemented"; return nil }

// If first block, save validators for the block.

// This extra logic due to Tendermint validator set changes being delayed 1 block.
// It may get overwritten due to InitChain validator updates.

// Save next validators.

// Save next consensus params.

// fmt.Printf("Tendermint State Saved height=%d hash=%X lastResultHash=%X\n", state.LastBlockHeight, state.AppHash, state.LastResultsHash)

// BootstrapState saves a new state, used e.g. by state sync when starting from non-zero height.
func (store dbStore) Bootstrap(state State) error { _ = "STUB: not implemented"; return nil }

// PruneStates deletes states up to the height specified (exclusive). It is not
// guaranteed to delete all states, since the last checkpointed state and states being pointed to by
// e.g. `LastHeightChanged` must remain. The state at retain height must also exist.
// Pruning is done in descending order.
func (store dbStore) PruneStates(retainHeight int64) error { _ = "STUB: not implemented"; return nil }

// NOTE: We need to prune consensus params first because the validator
// sets have always one extra height. If validator sets were pruned first
// we could get a situation where we prune up to the last validator set
// yet don't have the respective consensus params at that height and thus
// return an error

// pruneValidatorSets calls a reverse iterator from base height to retain height (exclusive), deleting
// all validator sets in between. Due to the fact that most validator sets stored reference an earlier
// validator set, it is likely that there will remain one validator set left after pruning.
func (store dbStore) pruneValidatorSets(retainHeight int64) error {
	_ = "STUB: not implemented"
	return nil
}

// We will prune up to the validator set at the given "height". As we don't save validator sets every
// height but only when they change or at a check point, it is likely that the validator set at the height
// we prune to is empty and thus dependent on the validator set saved at a previous height. We must find
// that validator set and make sure it is not pruned.

// if this is not equal to the retain height, prune from the retain height to the height above
// the last saved validator set. This way we can skip over the dependent validator set.

// prune all the validators sets up to last saved validator set

// pruneConsensusParams calls a reverse iterator from base height to retain height batch deleting
// all consensus params in between. If the consensus params at the new base height is dependent
// on a prior height then this will keep that lower height too.
func (store dbStore) pruneConsensusParams(retainHeight int64) error {
	_ = "STUB: not implemented"
	return nil
}

// As we don't save the consensus params at every height, only when there is a consensus params change,
// we must not prune (or save) the last consensus params that the consensus params info at height
// is dependent on.

// sanity check that the consensus params at the last height it was changed is there

// prune the params above the height with which it last changed and below the retain height.

// prune all the consensus params up to either the last height the params changed or if the params
// last changed at the retain height, then up to the retain height.

// pruneFinalizeBlockResponses calls a reverse iterator from base height to retain height
// batch deleting all responses to FinalizeBlock, and legacy ABCI responses, in between
func (store dbStore) pruneFinalizeBlockResponses(height int64) error {
	_ = "STUB: not implemented"
	return nil
}

// Remove any stale legacy ABCI responses

// pruneRange is a generic function for deleting a range of keys in reverse order.
// we keep filling up batches of at most 1000 keys, perform a deletion and continue until
// we have gone through all the keys in the range. This avoids doing any writes whilst
// iterating.
func (store dbStore) pruneRange(start []byte, end []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// iterate until the last batch of the pruning range in which case we will perform a
// write sync

// fill a new batch of keys for deletion over the remainding range

// reverseBatchDelete runs a reverse iterator (from end to start) filling up a batch until either
// (a) the iterator reaches the start or (b) the iterator has added a 1000 keys (this avoids the
// batch from growing too large)
func (store dbStore) reverseBatchDelete(batch dbm.Batch, start, end []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// avoid batches growing too large by capping them

//------------------------------------------------------------------------

// LoadFinalizeBlockResponses loads the responses to FinalizeBlock for the
// given height from the database. If not found,
// ErrNoFinalizeBlockResponsesForHeight is returned.
//
// This is useful for recovering from crashes where we called app.Commit
// and before we called s.Save(). It can also be used to produce Merkle
// proofs of the result of txs.
func (store dbStore) LoadFinalizeBlockResponses(height int64) (*abci.ResponseFinalizeBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DATA HAS BEEN CORRUPTED OR THE SPEC HAS CHANGED

// TODO: ensure that buf is completely read.

// SaveFinalizeBlockResponses persists to the database the responses to FinalizeBlock.
// This is useful in case we crash after app.Commit and before s.Save().
// Responses are indexed by height so they can also be loaded later to produce
// Merkle proofs.
//
// Exposed for testing.
func (store dbStore) SaveFinalizeBlockResponses(height int64, finalizeBlockResponses *abci.ResponseFinalizeBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (store dbStore) saveFinalizeBlockResponses(height int64, finalizeBlockResponses *abci.ResponseFinalizeBlock) error {
	_ = "STUB: not implemented"
	return nil
}

// strip nil values,

// SaveValidatorSets is used to save the validator set over multiple heights.
// It is exposed so that a backfill operation during state sync can populate
// the store with the necessary amount of validator sets to verify any evidence
// it may encounter.
func (store dbStore) SaveValidatorSets(lowerHeight, upperHeight int64, vals *types.ValidatorSet) error {
	_ = "STUB: not implemented"
	return nil
}

// batch together all the validator sets from lowerHeight to upperHeight

//-----------------------------------------------------------------------------

// LoadValidators loads the ValidatorSet for a given height.
// Returns ErrNoValSetForHeight if the validator set can't be found for this height.
func (store dbStore) LoadValidators(height int64) (*types.ValidatorSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// mutate

func lastStoredHeightFor(height, lastHeightChanged int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

func (store dbStore) LoadValidatorsInfo(height int64) (*tmstate.ValidatorsInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CONTRACT: Returned ValidatorsInfo can be mutated.
func loadValidatorsInfo(db dbm.DB, height int64) (*tmstate.ValidatorsInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DATA HAS BEEN CORRUPTED OR THE SPEC HAS CHANGED

// TODO: ensure that buf is completely read.

// saveValidatorsInfo persists the validator set.
//
// `height` is the effective height for which the validator is responsible for
// signing. It should be called from s.Save(), right before the state itself is
// persisted.
func (store dbStore) saveValidatorsInfo(
	height, lastHeightChanged int64,
	valSet *types.ValidatorSet,
	batch dbm.Batch,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Only persist validator set if it was updated or checkpoint height (see
// valSetCheckpointInterval) is reached.

//-----------------------------------------------------------------------------

// ConsensusParamsInfo represents the latest consensus params, or the last height it changed

// Allocate empty Consensus params at compile time to avoid multiple allocations during runtime
var (
	empty   = types.ConsensusParams{}
	emptypb = tmproto.ConsensusParams{}
)

// LoadConsensusParams loads the ConsensusParams for a given height.
func (store dbStore) LoadConsensusParams(height int64) (types.ConsensusParams, error) {
	_ = "STUB: not implemented"
	return *new(types.ConsensusParams), nil
}

func (store dbStore) loadConsensusParamsInfo(height int64) (*tmstate.ConsensusParamsInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DATA HAS BEEN CORRUPTED OR THE SPEC HAS CHANGED

// TODO: ensure that buf is completely read.

// saveConsensusParamsInfo persists the consensus params for the next block to disk.
// It should be called from s.Save(), right before the state itself is persisted.
// If the consensus params did not change after processing the latest block,
// only the last height for which they changed is persisted.
func (store dbStore) saveConsensusParamsInfo(
	nextHeight, changeHeight int64,
	params types.ConsensusParams,
	batch dbm.Batch,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (store dbStore) Close() error { _ = "STUB: not implemented"; return nil }

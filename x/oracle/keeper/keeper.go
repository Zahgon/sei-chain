package keeper

import (
	"sync"

	"github.com/sei-protocol/sei-chain/utils/datastructures"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	paramstypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"

	"github.com/sei-protocol/sei-chain/x/oracle/types"
)

// Keeper of the oracle store
type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   sdk.StoreKey
	memKey     sdk.StoreKey
	paramSpace paramstypes.Subspace

	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	distrKeeper   types.DistributionKeeper
	StakingKeeper types.StakingKeeper

	spamPreventionCounterMtxMap *datastructures.TypedSyncMap[string, *sync.Mutex]

	distrName string
}

// NewKeeper constructs a new keeper for oracle
func NewKeeper(cdc codec.BinaryCodec, storeKey sdk.StoreKey, memKey sdk.StoreKey,
	paramspace paramstypes.Subspace, accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper, distrKeeper types.DistributionKeeper,
	stakingKeeper types.StakingKeeper, distrName string,
) Keeper {
	_ = "STUB: not implemented"
	// ensure oracle module account is set
	return *new(Keeper)
}

// set KeyTable if it has not already been set

//-----------------------------------
// ExchangeRate logic

func (k Keeper) GetBaseExchangeRate(ctx sdk.Context, denom string) (sdk.Dec, sdk.Int, int64, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Dec), *new(sdk.Int), 0, nil
}

func (k Keeper) SetBaseExchangeRate(ctx sdk.Context, denom string, exchangeRate sdk.Dec) {
	_ = "STUB: not implemented"
	return
}

func (k Keeper) SetBaseExchangeRateWithEvent(ctx sdk.Context, denom string, exchangeRate sdk.Dec) {
	_ = "STUB: not implemented"
	return
}

func (k Keeper) DeleteBaseExchangeRate(ctx sdk.Context, denom string) {
	_ = "STUB: not implemented"
	return
}

func (k Keeper) IterateBaseExchangeRates(ctx sdk.Context, handler func(denom string, exchangeRate types.OracleExchangeRate) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

func (k Keeper) RemoveExcessFeeds(ctx sdk.Context) {
	_ = "STUB: not implemented"
	// get actives
	return
}

// get vote targets

// remove vote targets from actives

// compare

// clear exchange rates

//-----------------------------------
// Oracle delegation logic

// GetFeederDelegation gets the account address that the validator operator delegated oracle vote rights to
func (k Keeper) GetFeederDelegation(ctx sdk.Context, operator sdk.ValAddress) sdk.AccAddress {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}

// By default the right is delegated to the validator itself

// SetFeederDelegation sets the account address that the validator operator delegated oracle vote rights to
func (k Keeper) SetFeederDelegation(ctx sdk.Context, operator sdk.ValAddress, delegatedFeeder sdk.AccAddress) {
	_ = "STUB: not implemented"
	return
}

// IterateFeederDelegations iterates over the feed delegates and performs a callback function.
func (k Keeper) IterateFeederDelegations(ctx sdk.Context,
	handler func(delegator sdk.ValAddress, delegate sdk.AccAddress) (stop bool),
) {
	_ = "STUB: not implemented"
	return
}

//-----------------------------------
// Miss counter logic

// GetVotePenaltyCounter retrieves the # of vote periods missed and abstained in this oracle slash window
func (k Keeper) GetVotePenaltyCounter(ctx sdk.Context, operator sdk.ValAddress) types.VotePenaltyCounter {
	_ = "STUB: not implemented"
	return *new(types.VotePenaltyCounter)
}

// By default the empty counter has values of 0

// SetVotePenaltyCounter updates the # of vote periods missed in this oracle slash window
func (k Keeper) SetVotePenaltyCounter(ctx sdk.Context, operator sdk.ValAddress, missCount, abstainCount, successCount uint64) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec
//nolint:gosec
//nolint:gosec
// TODO(PLT-336): remove once oracle_vote_penalty_count verified

func (k Keeper) IncrementMissCount(ctx sdk.Context, operator sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

func (k Keeper) IncrementAbstainCount(ctx sdk.Context, operator sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

func (k Keeper) IncrementSuccessCount(ctx sdk.Context, operator sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

func (k Keeper) GetMissCount(ctx sdk.Context, operator sdk.ValAddress) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func (k Keeper) GetAbstainCount(ctx sdk.Context, operator sdk.ValAddress) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func (k Keeper) GetSuccessCount(ctx sdk.Context, operator sdk.ValAddress) uint64 {
	_ = "STUB: not implemented"
	return 0
}

// DeleteVotePenaltyCounter removes miss counter for the validator
func (k Keeper) DeleteVotePenaltyCounter(ctx sdk.Context, operator sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

// TODO(PLT-336): remove once oracle_vote_penalty_count verified

// IterateVotePenaltyCounters iterates over the miss counters and performs a callback function.
func (k Keeper) IterateVotePenaltyCounters(ctx sdk.Context,
	handler func(operator sdk.ValAddress, votePenaltyCounter types.VotePenaltyCounter) (stop bool),
) {
	_ = "STUB: not implemented"
	return
}

//-----------------------------------
// AggregateExchangeRateVote logic

// GetAggregateExchangeRateVote retrieves an oracle vote from the store
func (k Keeper) GetAggregateExchangeRateVote(ctx sdk.Context, voter sdk.ValAddress) (aggregateVote types.AggregateExchangeRateVote, err error) {
	_ = "STUB: not implemented"
	return *new(types.AggregateExchangeRateVote), nil
}

// SetAggregateExchangeRateVote adds an oracle aggregate vote to the store
func (k Keeper) SetAggregateExchangeRateVote(ctx sdk.Context, voter sdk.ValAddress, vote types.AggregateExchangeRateVote) {
	_ = "STUB: not implemented"
	return
}

// DeleteAggregateExchangeRateVote deletes an oracle vote from the store
func (k Keeper) DeleteAggregateExchangeRateVote(ctx sdk.Context, voter sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

// IterateAggregateExchangeRateVotes iterates rate over votes in the store
func (k Keeper) IterateAggregateExchangeRateVotes(ctx sdk.Context, handler func(voterAddr sdk.ValAddress, aggregateVote types.AggregateExchangeRateVote) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

func (k Keeper) GetVoteTarget(ctx sdk.Context, denom string) (types.Denom, error) {
	_ = "STUB: not implemented"
	return *new(types.Denom), nil
}

func (k Keeper) SetVoteTarget(ctx sdk.Context, denom string) { _ = "STUB: not implemented"; return }

func (k Keeper) IterateVoteTargets(ctx sdk.Context, handler func(denom string, denomInfo types.Denom) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

func (k Keeper) ClearVoteTargets(ctx sdk.Context) { _ = "STUB: not implemented"; return }

func (k Keeper) getAllKeysForPrefix(store sdk.KVStore, prefix []byte) [][]byte {
	_ = "STUB: not implemented"
	return nil
}

// ValidateFeeder return the given feeder is allowed to feed the message or not
func (k Keeper) ValidateFeeder(ctx sdk.Context, feederAddr sdk.AccAddress, validatorAddr sdk.ValAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// Check that the given validator exists

func (k Keeper) GetPriceSnapshot(ctx sdk.Context, timestamp int64) types.PriceSnapshot {
	_ = "STUB: not implemented"
	return *new(types.PriceSnapshot)
}

//nolint:gosec

func (k Keeper) SetPriceSnapshot(ctx sdk.Context, snapshot types.PriceSnapshot) {
	_ = "STUB: not implemented"
	// shouldn't be used directly, use "add" instead for individual price snapshots
	return
}

//nolint:gosec

func (k Keeper) AddPriceSnapshot(ctx sdk.Context, snapshot types.PriceSnapshot) {
	_ = "STUB: not implemented"
	return
}

// Sanity check to make sure LookbackDuration can be converted to int64
// Lookback duration should never get this large

// Check

// we need to evict old snapshots (except for one that is out of range)

// delete the previous out of range snapshot

// update last out of range snapshot

func (k Keeper) IteratePriceSnapshots(ctx sdk.Context, handler func(snapshot types.PriceSnapshot) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

func (k Keeper) IteratePriceSnapshotsReverse(ctx sdk.Context, keyPrefix []byte, handler func(snapshot types.PriceSnapshot) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

func (k Keeper) DeletePriceSnapshot(ctx sdk.Context, timestamp int64) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec

func (k Keeper) CalculateTwaps(ctx sdk.Context, lookbackSeconds uint64) (types.OracleTwaps, error) {
	_ = "STUB: not implemented"
	return *new(types.OracleTwaps), nil
}

// get targets - only calculate for the targets

//nolint:gosec

//nolint:gosec

// update time traversed to represent current snapshot
// replace SnapshotTimestamp with lookback duration bounding

// iterate through denoms in the snapshot
// if we find a new one, we have to setup the TWAP calc for that one

// set up the TWAP info for a denom

// get the denom specific TWAP data

// calculate the new Time Weighted Sum for the denom exchange rate
// we calculate a weighted sum of exchange rates previously by multiplying each exchange rate by time interval that it was active
// then we divide by the overall time in the lookback window, which gives us the time weighted average

// set the denom TWAP data

// iterate over all denoms with TWAP data

// divide the denom time weighed sum by denom duration

func (k Keeper) ValidateLookbackSeconds(ctx sdk.Context, lookbackSeconds uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (k Keeper) CheckAndSetSpamPreventionCounter(ctx sdk.Context, validatorAddr sdk.ValAddress) error {
	_ = "STUB: not implemented"
	return nil
}

func (k Keeper) getSpamPreventionCounter(ctx sdk.Context, validatorAddr sdk.ValAddress) int64 {
	_ = "STUB: not implemented"
	return 0
}

//nolint:gosec

func (k Keeper) setSpamPreventionCounter(ctx sdk.Context, validatorAddr sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

// nolint:gosec

func (k Keeper) GetStoreKey() sdk.StoreKey { _ = "STUB: not implemented"; return *new(sdk.StoreKey) }

func (k Keeper) GetCdc() codec.BinaryCodec {
	_ = "STUB: not implemented"
	return *new(codec.BinaryCodec)
}

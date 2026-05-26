package oracle

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/x/oracle/keeper"
	"github.com/sei-protocol/sei-chain/x/oracle/types"
)

// Tally calculates the median and returns it. Sets the set of voters to be rewarded, i.e. voted within
// a reasonable spread from the weighted median to the store
// CONTRACT: pb must be sorted
func Tally(_ sdk.Context, pb types.ExchangeRateBallot, rewardBand sdk.Dec, validatorClaimMap map[string]types.Claim) (weightedMedian sdk.Dec) {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// Filter ballot winners

func UpdateDidVote(ctx sdk.Context, ballot types.ExchangeRateBallot, validatorClaimMap map[string]types.Claim) {
	_ = "STUB: not implemented"
	return
}

// because we can't actually effectively calculate a reward band for the below threshold votes,
// simply voting here counts as a win++. However, the validator still needs to be within band
// for over threshold denoms too so as to not be counted as "miss"

// ballot for the asset is passing the threshold amount of voting power
func ballotIsPassing(ballot types.ExchangeRateBallot, thresholdVotes sdk.Int) (sdk.Int, bool) {
	_ = "STUB: not implemented"
	return *new(sdk.Int), false
}

// choose reference denom with the highest voter turnout
// If the voting power of the two denominations is the same,
// select reference denom in alphabetical order.
func pickReferenceDenom(ctx sdk.Context, k keeper.Keeper, voteTargets map[string]types.Denom, voteMap map[string]types.ExchangeRateBallot) (referenceDenom string, belowThresholdVoteMap map[string]types.ExchangeRateBallot) {
	_ = "STUB: not implemented"
	return "", nil
}

// If denom is not in the voteTargets, or the ballot for it has failed, then skip
// and remove it from voteMap for iteration efficiency

// If the ballot is not passed, remove it from the voteTargets array
// to prevent slashing validators who did valid vote.

// add assets below threshold to separate map for tally evaluation

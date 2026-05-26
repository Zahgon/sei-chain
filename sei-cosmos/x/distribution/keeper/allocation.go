package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	stakingtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
)

// AllocateTokens handles distribution of the collected fees
// bondedVotes is a list of (validator address, validator voted on last block flag) for all
// validators in the bonded set.
func (k Keeper) AllocateTokens(
	ctx sdk.Context, sumPreviousPrecommitPower, totalPreviousPower int64,
	previousProposer sdk.ConsAddress, bondedVotes []abci.VoteInfo,
) {
	_ = "STUB: not implemented"

	// fetch and clear the collected fees for distribution, since this is
	// called in BeginBlock, collected fees will be from the previous block
	// (and distributed to the previous proposer)
	return
}

// transfer collected fees to the distribution module account

// temporary workaround to keep CanWithdrawInvariant happy
// general discussions here: https://github.com/cosmos/cosmos-sdk/issues/2906#issuecomment-441867634

// calculate fraction votes

// calculate previous proposer reward

// pay previous proposer

// previous proposer can be unknown if say, the unbonding period is 1 block, so
// e.g. a validator undelegates at block X, it's removed entirely by
// block X+1's endblock, then X+2 we need to refer to the previous
// proposer for X+1, but we've forgotten about them.

// calculate fraction allocated to validators

// allocate tokens proportionally to voting power
//
// TODO: Consider parallelizing later
//
// Ref: https://github.com/cosmos/cosmos-sdk/pull/3099#discussion_r246276376

// TODO: Consider micro-slashing for missing votes.
//
// Ref: https://github.com/cosmos/cosmos-sdk/issues/2525#issuecomment-430838701

// allocate community funding

// AllocateTokensToValidator allocate tokens to a particular validator,
// splitting according to commission.
func (k Keeper) AllocateTokensToValidator(ctx sdk.Context, val stakingtypes.ValidatorI, tokens sdk.DecCoins) {
	_ = "STUB: not implemented"
	// split tokens between validator and delegators according to commission
	return
}

// update current commission

// update current rewards

// update outstanding rewards

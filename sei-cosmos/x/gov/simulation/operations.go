package simulation

import (
	"math/rand"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	simtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/simulation"
)

var initialProposalID = uint64(100000000000000)

// Simulation operation weights constants
const (
	OpWeightMsgDeposit      = "op_weight_msg_deposit"       //nolint:gosec
	OpWeightMsgVote         = "op_weight_msg_vote"          //nolint:gosec
	OpWeightMsgVoteWeighted = "op_weight_msg_weighted_vote" //nolint:gosec
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec, ak types.AccountKeeper,
	bk types.BankKeeper, k keeper.Keeper, wContents []simtypes.WeightedProposalContent,
) simulation.WeightedOperations {
	_ = "STUB: not implemented"
	return *new(simulation.WeightedOperations)
}

// generate the weighted operations for the proposal contents

// pin variable

// SimulateMsgSubmitProposal simulates creating a msg Submit Proposal
// voting on the proposal, and subsequently slashing the proposal. It is implemented using
// future operations.
func SimulateMsgSubmitProposal(
	ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper, contentSim simtypes.ContentSimulatorFn,
) simtypes.Operation {
	_ = "STUB: not implemented"
	// The states are:
	// column 1: All validators vote
	// column 2: 90% vote
	// column 3: 75% vote
	// column 4: 40% vote
	// column 5: 15% vote
	// column 6: noone votes
	// All columns sum to 100 for simplicity, values chosen by @valardragon semi-arbitrarily,
	// feel free to change.
	return *new(simtypes.Operation)
}

// 1) submit proposal now

// get the submitted proposal ID

// 2) Schedule operations for votes
// 2.1) first pick a number of people to vote.

// 2.2) select who votes and when

// didntVote := whoVotes[numVotes:]

//nolint:gosec // voting period seconds is a small positive config value

//nolint:gosec // proposal IDs are sequential small values

// SimulateMsgDeposit generates a MsgDeposit with random values.
func SimulateMsgDeposit(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

// SimulateMsgVote generates a MsgVote with random values.
func SimulateMsgVote(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

func operationSimulateMsgVote(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper,
	simAccount simtypes.Account, proposalIDInt int64) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

// SimulateMsgVoteWeighted generates a MsgVoteWeighted with random values.
func SimulateMsgVoteWeighted(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

func operationSimulateMsgVoteWeighted(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper,
	simAccount simtypes.Account, proposalIDInt int64) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

// Pick a random deposit with a random denomination with a
// deposit amount between (0, min(balance, minDepositAmount))
// This is to simulate multiple users depositing to get the
// proposal above the minimum deposit amount
func randomDeposit(r *rand.Rand, ctx sdk.Context,
	ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper, addr sdk.AccAddress,
) (deposit sdk.Coins, skip bool, err error) {
	_ = "STUB: not implemented"
	return *new(sdk.Coins), false, nil
}

// skip

// Pick a random proposal ID between the initial proposal ID
// (defined in gov GenesisState) and the latest proposal ID
// that matches a given Status.
// It does not provide a default ID.
func randomProposalID(r *rand.Rand, k keeper.Keeper,
	ctx sdk.Context, status types.ProposalStatus) (proposalID uint64, found bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// select a random ID between [initialProposalID, proposalID]

//#nosec G115 -- bounds checked above

// This is called on the first call to this function
// in order to update the global variable

// Pick a random voting option
func randomVotingOption(r *rand.Rand) types.VoteOption {
	_ = "STUB: not implemented"
	return *new(types.VoteOption)
}

// Pick a random weighted voting options
func randomWeightedVotingOptions(r *rand.Rand) types.WeightedVoteOptions {
	_ = "STUB: not implemented"
	return *new(types.WeightedVoteOptions)
}

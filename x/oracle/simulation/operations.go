package simulation

// DONTCOVER

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	simtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/simulation"

	"github.com/sei-protocol/sei-chain/x/oracle/keeper"
	"github.com/sei-protocol/sei-chain/x/oracle/types"
)

// Simulation operation weights constants
//
//nolint:gosec
const (
	OpWeightMsgAggregateExchangeRateVote = "op_weight_msg_exchange_rate_aggregate_vote"
	OpWeightMsgDelegateFeedConsent       = "op_weight_msg_exchange_feed_consent"
)

var voteHashMap = make(map[string]string)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams,
	cdc codec.JSONCodec,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simulation.WeightedOperations {
	_ = "STUB: not implemented"
	return *new(simulation.WeightedOperations)
}

// SimulateMsgAggregateExchangeRateVote generates a MsgAggregateExchangeRateVote with random values.
// nolint: funlen
func SimulateMsgAggregateExchangeRateVote(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

// ensure the validator exists

// ensure vote hash exists

// SimulateMsgDelegateFeedConsent generates a MsgDelegateFeedConsent with random values.
// nolint: funlen
func SimulateMsgDelegateFeedConsent(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

// ensure the validator exists

// ensure the target address is not a validator

package simulation

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	simtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/simulation"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/slashing/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/slashing/types"
	stakingkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/keeper"
)

// Simulation operation weights constants
const (
	OpWeightMsgUnjail = "op_weight_msg_unjail" //nolint:gosec
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec, ak types.AccountKeeper,
	bk types.BankKeeper, k keeper.Keeper, sk types.StakingKeeper,
) simulation.WeightedOperations {
	_ = "STUB: not implemented"
	return *new(simulation.WeightedOperations)
}

// SimulateMsgUnjail generates a MsgUnjail with random values
func SimulateMsgUnjail(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper, sk stakingkeeper.Keeper) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

// skip

// skip

// TODO: due to this condition this message is almost, if not always, skipped !

// skip

// skip

// result should fail if:
// - validator cannot be unjailed due to tombstone
// - validator is still in jailed period
// - self delegation too low

// msg failed as expected

package simulation

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	simtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/simulation"
	stakingkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/keeper"
)

// Simulation operation weights constants
const (
	OpWeightMsgSetWithdrawAddress          = "op_weight_msg_set_withdraw_address"          //nolint:gosec
	OpWeightMsgWithdrawDelegationReward    = "op_weight_msg_withdraw_delegation_reward"    //nolint:gosec
	OpWeightMsgWithdrawValidatorCommission = "op_weight_msg_withdraw_validator_commission" //nolint:gosec
	OpWeightMsgFundCommunityPool           = "op_weight_msg_fund_community_pool"           //nolint:gosec
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(appParams simtypes.AppParams, cdc codec.JSONCodec, ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper, sk types.StakingKeeper) simulation.WeightedOperations {
	_ = "STUB: not implemented"
	return *new(simulation.WeightedOperations)
}

// SimulateMsgSetWithdrawAddress generates a MsgSetWithdrawAddress with random values.
func SimulateMsgSetWithdrawAddress(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

// SimulateMsgWithdrawDelegatorReward generates a MsgWithdrawDelegatorReward with random values.
func SimulateMsgWithdrawDelegatorReward(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper, sk stakingkeeper.Keeper) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

// SimulateMsgWithdrawValidatorCommission generates a MsgWithdrawValidatorCommission with random values.
func SimulateMsgWithdrawValidatorCommission(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper, sk stakingkeeper.Keeper) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

// SimulateMsgFundCommunityPool simulates MsgFundCommunityPool execution where
// a random account sends a random amount of its funds to the community pool.
func SimulateMsgFundCommunityPool(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper, sk stakingkeeper.Keeper) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

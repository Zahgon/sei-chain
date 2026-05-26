package simulation

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	simtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/simulation"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
)

// Simulation operation weights constants
const (
	OpWeightMsgCreateValidator = "op_weight_msg_create_validator" //nolint:gosec
	OpWeightMsgEditValidator   = "op_weight_msg_edit_validator"   //nolint:gosec
	OpWeightMsgDelegate        = "op_weight_msg_delegate"         //nolint:gosec
	OpWeightMsgUndelegate      = "op_weight_msg_undelegate"       //nolint:gosec
	OpWeightMsgBeginRedelegate = "op_weight_msg_begin_redelegate" //nolint:gosec
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec, ak types.AccountKeeper,
	bk types.BankKeeper, k keeper.Keeper,
) simulation.WeightedOperations {
	_ = "STUB: not implemented"
	return *new(simulation.WeightedOperations)
}

// SimulateMsgCreateValidator generates a MsgCreateValidator with random values
func SimulateMsgCreateValidator(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

// ensure the validator doesn't exist already

// SimulateMsgEditValidator generates a MsgEditValidator with random values
func SimulateMsgEditValidator(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

// skip as the commission is invalid

// SimulateMsgDelegate generates a MsgDelegate with random values
func SimulateMsgDelegate(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

// SimulateMsgUndelegate generates a MsgUndelegate with random values
func SimulateMsgUndelegate(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

// get random validator

// get random delegator from validator

// need to retrieve the simulation account associated with delegation to retrieve PrivKey

// if simaccount.PrivKey == nil, delegation address does not exist in accs. Return error

// SimulateMsgBeginRedelegate generates a MsgBeginRedelegate with random values
func SimulateMsgBeginRedelegate(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

// get random source validator

// get random delegator from src validator

// skip

// get random destination validator

// check if the shares truncate to zero

// skip

// need to retrieve the simulation account associated with delegation to retrieve PrivKey

// if simaccount.PrivKey == nil, delegation address does not exist in accs. Return error

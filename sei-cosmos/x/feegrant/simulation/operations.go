package simulation

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	simtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/feegrant"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/feegrant/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/simulation"
)

// Simulation operation weights constants
const (
	OpWeightMsgGrantAllowance  = "op_weight_msg_grant_fee_allowance"    //nolint:gosec
	OpWeightMsgRevokeAllowance = "op_weight_msg_grant_revoke_allowance" //nolint:gosec
)

var (
	TypeMsgGrantAllowance  = sdk.MsgTypeURL(&feegrant.MsgGrantAllowance{})
	TypeMsgRevokeAllowance = sdk.MsgTypeURL(&feegrant.MsgRevokeAllowance{})
)

func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	ak feegrant.AccountKeeper, bk feegrant.BankKeeper, k keeper.Keeper,
) simulation.WeightedOperations {
	_ = "STUB: not implemented"
	return *new(simulation.WeightedOperations)
}

// SimulateMsgGrantAllowance generates MsgGrantAllowance with random values.
func SimulateMsgGrantAllowance(ak feegrant.AccountKeeper, bk feegrant.BankKeeper, k keeper.Keeper) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

// SimulateMsgRevokeAllowance generates a MsgRevokeAllowance with random values.
func SimulateMsgRevokeAllowance(ak feegrant.AccountKeeper, bk feegrant.BankKeeper, k keeper.Keeper) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

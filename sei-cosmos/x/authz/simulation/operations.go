package simulation

import (
	"math/rand"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	cdctypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	simtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/authz"

	"github.com/sei-protocol/sei-chain/sei-cosmos/x/authz/keeper"

	"github.com/sei-protocol/sei-chain/sei-cosmos/x/simulation"
)

// authz message types
var (
	TypeMsgGrant  = sdk.MsgTypeURL(&authz.MsgGrant{})
	TypeMsgRevoke = sdk.MsgTypeURL(&authz.MsgRevoke{})
	TypeMsgExec   = sdk.MsgTypeURL(&authz.MsgExec{})
)

// Simulation operation weights constants
const (
	OpWeightMsgGrant = "op_weight_msg_grant"   //nolint:gosec
	OpWeightRevoke   = "op_weight_msg_revoke"  //nolint:gosec
	OpWeightExec     = "op_weight_msg_execute" //nolint:gosec
)

// authz operations weights
const (
	WeightGrant  = 100
	WeightRevoke = 90
	WeightExec   = 90
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec, ak authz.AccountKeeper, bk authz.BankKeeper, k keeper.Keeper, appCdc cdctypes.AnyUnpacker) simulation.WeightedOperations {
	_ = "STUB: not implemented"
	return *new(simulation.WeightedOperations)
}

// SimulateMsgGrant generates a MsgGrant with random values.
func SimulateMsgGrant(ak authz.AccountKeeper, bk authz.BankKeeper, _ keeper.Keeper) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

func generateRandomAuthorization(r *rand.Rand, spendLimit sdk.Coins) authz.Authorization {
	_ = "STUB: not implemented"
	return *new(authz.Authorization)
}

// SimulateMsgRevoke generates a MsgRevoke with random values.
func SimulateMsgRevoke(ak authz.AccountKeeper, bk authz.BankKeeper, k keeper.Keeper) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

// SimulateMsgExec generates a MsgExec with random values.
func SimulateMsgExec(ak authz.AccountKeeper, bk authz.BankKeeper, k keeper.Keeper, cdc cdctypes.AnyUnpacker) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

// Check send_enabled status of each sent coin denom

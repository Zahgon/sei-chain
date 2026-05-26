package simulation

import (
	"math/rand"

	"github.com/sei-protocol/sei-chain/sei-cosmos/baseapp"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	simtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/simulation"
)

// Simulation operation weights constants
const (
	OpWeightMsgSend      = "op_weight_msg_send"      //nolint:gosec
	OpWeightMsgMultiSend = "op_weight_msg_multisend" //nolint:gosec
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec, ak types.AccountKeeper, bk keeper.Keeper,
) simulation.WeightedOperations {
	_ = "STUB: not implemented"
	return *new(simulation.WeightedOperations)
}

// SimulateMsgSend tests and runs a single msg send where both
// accounts already exist.
func SimulateMsgSend(ak types.AccountKeeper, bk keeper.Keeper) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

// Check send_enabled status of each coin denom

// SimulateMsgSendToModuleAccount tests and runs a single msg send where both
// accounts already exist.
func SimulateMsgSendToModuleAccount(ak types.AccountKeeper, bk keeper.Keeper, moduleAccCount int) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

// Check send_enabled status of each coin denom

// sendMsgSend sends a transaction with a MsgSend from a provided random account.
func sendMsgSend(
	r *rand.Rand, app *baseapp.BaseApp, bk keeper.Keeper, ak types.AccountKeeper,
	msg *types.MsgSend, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
) error {
	_ = "STUB: not implemented"
	return nil
}

// SimulateMsgMultiSend tests and runs a single msg multisend, with randomized, capped number of inputs/outputs.
// all accounts in msg fields exist in state
func SimulateMsgMultiSend(ak types.AccountKeeper, bk keeper.Keeper) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

// random number of inputs/outputs between [1, 3]

// collect signer privKeys

// use map to check if address already exists as input

// generate random input fields, ignore to address

// make sure account is fresh and not used in previous input

// set input address in used address map

// set signer privkey

// set next input and accumulate total sent coins

// Check send_enabled status of each sent coin denom

// split total sent coins into random subsets for output

// take random subset of remaining coins for output
// and update remaining coins

// remove any output that has no coins

// continue onto next coin

// SimulateMsgMultiSendToModuleAccount sends coins to Module Accounts
func SimulateMsgMultiSendToModuleAccount(ak types.AccountKeeper, bk keeper.Keeper, moduleAccCount int) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

// collect signer privKeys

// split total sent coins into random subsets for output

// take random subset of remaining coins for output
// and update remaining coins

// remove any output that has no coins

// continue onto next coin

// sendMsgMultiSend sends a transaction with a MsgMultiSend from a provided random
// account.
func sendMsgMultiSend(
	r *rand.Rand, app *baseapp.BaseApp, bk keeper.Keeper, ak types.AccountKeeper,
	msg *types.MsgMultiSend, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
) error {
	_ = "STUB: not implemented"
	return nil
}

// feePayer is the first signer, i.e. first input address

// randomSendFields returns the sender and recipient simulation accounts as well
// as the transferred amount.
func randomSendFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, bk keeper.Keeper, ak types.AccountKeeper,
) (simtypes.Account, simtypes.Account, sdk.Coins, bool) {
	_ = "STUB: not implemented"
	return *new(simtypes.Account), *new(simtypes.Account), *new(sdk.Coins), false
}

// disallow sending money to yourself

func getModuleAccounts(ak types.AccountKeeper, ctx sdk.Context, moduleAccCount int) []simtypes.Account {
	_ = "STUB: not implemented"
	return nil
}

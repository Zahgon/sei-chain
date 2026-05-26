package upgrade

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("cosmos", "x", "upgrade")

// BeginBlock will check if there is a scheduled plan and if it is ready to be executed.
// If the current height is in the provided set of heights to skip, it will skip and clear the upgrade plan.
// If it is ready, it will execute it if the handler is installed, and panic/abort otherwise.
// If the plan is not ready, it will ensure the handler is not registered too early (and abort otherwise).
//
// The purpose is to ensure the binary is switched EXACTLY at the desired block, and to allow
// a migration to be executed if needed upon this switch (migration defined in the new binary)
// skipUpgradeHeightArray is a set of block heights for which the upgrade must be skipped
func BeginBlocker(k keeper.Keeper, ctx sdk.Context) { _ = "STUB: not implemented"; return }

// This check will make sure that we are using a valid binary.
// It'll panic in these cases if there is no upgrade handler registered for the last applied upgrade.
// 1. If there is no scheduled upgrade.
// 2. If the plan is not ready.
// 3. If the plan is ready and skip upgrade height is set for current height.

// If the plan's block height has passed, then it must be the executed version
// All major and minor releases are REQUIRED to execute on the scheduled block height

// If skip upgrade has been set for current height, we clear the upgrade plan

// If we don't have an upgrade handler for this upgrade name, then we need to shutdown

// If running a pending minor release, apply the upgrade if handler is present
// Minor releases are allowed to run before the scheduled upgrade height, but not required to.

// if not yet present, then emit a scheduled log (every 100 blocks, to reduce logs)

// if we have a handler for a non-minor upgrade, that means it updated too early and must stop

// panicUpgradeNeeded shuts down the node and prints a message that the upgrade needs to be applied.
func panicUpgradeNeeded(k keeper.Keeper, ctx sdk.Context, plan types.Plan) {
	_ = "STUB: not implemented"
	// Write the upgrade info to disk. The UpgradeStoreLoader uses this info to perform or skip
	// store migrations.
	return
}

//nolint:staticcheck // SA1019: not worth fixing

// Emit the raw upgrade message to stderr so that cosmovisor's log scanner
// can match it. Structured log handlers (JSON, text) escape the inner
// double-quotes, which breaks the cosmovisor upgrade regex.

func applyUpgrade(k keeper.Keeper, ctx sdk.Context, plan types.Plan) {
	_ = "STUB: not implemented"
	return
}

// skipUpgrade logs a message that the upgrade has been skipped and clears the upgrade plan.
func skipUpgrade(k keeper.Keeper, ctx sdk.Context, plan types.Plan) {
	_ = "STUB: not implemented"
	return
}

// BuildUpgradeNeededMsg prints the message that notifies that an upgrade is needed.
func BuildUpgradeNeededMsg(plan types.Plan) string { _ = "STUB: not implemented"; return "" }

// BuildUpgradeScheduledMsg prints upgrade scheduled message
func BuildUpgradeScheduledMsg(plan types.Plan) string { _ = "STUB: not implemented"; return "" }

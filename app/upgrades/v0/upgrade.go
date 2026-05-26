package v0

import (
	"github.com/sei-protocol/sei-chain/app/upgrades"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm"
	wasmkeeper "github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/keeper"
)

const (
	UpgradeName = "v0"
)

// HardForkUpgradeHandler defines an example hard fork handler that will be
// executed during BeginBlock at a target height and chain-ID.
type HardForkUpgradeHandler struct {
	TargetHeight  int64
	TargetChainID string
	WasmKeeper    wasm.Keeper
}

func NewHardForkUpgradeHandler(height int64, chainID string, wk wasm.Keeper) upgrades.HardForkHandler {
	_ = "STUB: not implemented"
	return *new(upgrades.HardForkHandler)
}

func (h HardForkUpgradeHandler) GetName() string { _ = "STUB: not implemented"; return "" }

func (h HardForkUpgradeHandler) GetTargetChainID() string { _ = "STUB: not implemented"; return "" }

func (h HardForkUpgradeHandler) GetTargetHeight() int64 { _ = "STUB: not implemented"; return 0 }

func (h HardForkUpgradeHandler) ExecuteHandler(ctx sdk.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// If other contract need to be migrated, create functions for them and pass
// the govKeeper to them.

func (h HardForkUpgradeHandler) migrateGringotts(ctx sdk.Context, govKeeper *wasmkeeper.PermissionedKeeper) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Fill in the appropriate fields (contractAddr, newCodeID, and msg) here!

// Note: Since we're using a GovPermissionKeeper, the caller is not used/required,
// since the authz policy will automatically allow the migration.

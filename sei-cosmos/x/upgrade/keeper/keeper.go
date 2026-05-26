package keeper

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	store "github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
	xp "github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade/exported"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade/types"
)

// UpgradeInfoFileName file to store upgrade information
const UpgradeInfoFileName string = "upgrade-info.json"

// upgrade defines a comparable structure for sorting upgrades.
type upgrade struct {
	Name        string
	BlockHeight int64
}

type Keeper struct {
	homePath           string                          // root directory of app config
	skipUpgradeHeights map[int64]bool                  // map of heights to skip for an upgrade
	storeKey           sdk.StoreKey                    // key to access x/upgrade store
	cdc                codec.BinaryCodec               // App-wide binary codec
	upgradeHandlers    map[string]types.UpgradeHandler // map of plan name to upgrade handler
	versionSetter      xp.ProtocolVersionSetter        // implements setting the protocol version field on BaseApp
	downgradeVerified  bool                            // tells if we've already sanity checked that this binary version isn't being used against an old state.
}

// NewKeeper constructs an upgrade Keeper which requires the following arguments:
// skipUpgradeHeights - map of heights to skip an upgrade
// storeKey - a store key with which to access upgrade's store
// cdc - the app-wide binary codec
// homePath - root directory of the application's config
// vs - the interface implemented by baseapp which allows setting baseapp's protocol version field
func NewKeeper(skipUpgradeHeights map[int64]bool, storeKey sdk.StoreKey, cdc codec.BinaryCodec, homePath string, vs xp.ProtocolVersionSetter) Keeper {
	_ = "STUB: not implemented"
	return *new(Keeper)
}

// SetUpgradeHandler sets an UpgradeHandler for the upgrade specified by name. This handler will be called when the upgrade
// with this name is applied. In order for an upgrade with the given name to proceed, a handler for this upgrade
// must be set even if it is a no-op function.
func (k Keeper) SetUpgradeHandler(name string, upgradeHandler types.UpgradeHandler) {
	_ = "STUB: not implemented"
	return
}

// setProtocolVersion sets the protocol version to state
func (k Keeper) setProtocolVersion(ctx sdk.Context, v uint64) { _ = "STUB: not implemented"; return }

// getProtocolVersion gets the protocol version from state
func (k Keeper) getProtocolVersion(ctx sdk.Context) uint64 { _ = "STUB: not implemented"; return 0 }

// default value

// SetModuleVersionMap saves a given version map to state
func (k Keeper) SetModuleVersionMap(ctx sdk.Context, vm module.VersionMap) {
	_ = "STUB: not implemented"
	return
}

// Even though the underlying store (cachekv) store is sorted, we still
// prefer a deterministic iteration order of the map, to avoid undesired
// surprises if we ever change stores.

// GetModuleVersionMap returns a map of key module name and value module consensus version
// as defined in ADR-041.
func (k Keeper) GetModuleVersionMap(ctx sdk.Context) module.VersionMap {
	_ = "STUB: not implemented"
	return *new(module.VersionMap)
}

// first byte is prefix key, so we remove it here

// GetModuleVersions gets a slice of module consensus versions
func (k Keeper) GetModuleVersions(ctx sdk.Context) []*types.ModuleVersion {
	_ = "STUB: not implemented"
	return nil
}

// gets the version for a given module, and returns true if it exists, false otherwise
func (k Keeper) getModuleVersion(ctx sdk.Context, name string) (uint64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// ScheduleUpgrade schedules an upgrade based on the specified plan.
// If there is another Plan already scheduled, it will overwrite it
// (implicitly cancelling the current plan)
// ScheduleUpgrade will also write the upgraded client to the upgraded client path
// if an upgraded client is specified in the plan
func (k Keeper) ScheduleUpgrade(ctx sdk.Context, plan types.Plan) error {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: allow for the possibility of chains to schedule upgrades in begin block of the same block
// as a strategy for emergency hard fork recoveries

// clear any old IBC state stored by previous plan

// SetUpgradedClient sets the expected upgraded client for the next version of this chain at the last height the current chain will commit.
func (k Keeper) SetUpgradedClient(ctx sdk.Context, planHeight int64, bz []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// GetUpgradedClient gets the expected upgraded client for the next version of this chain
func (k Keeper) GetUpgradedClient(ctx sdk.Context, height int64) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// SetUpgradedConsensusState set the expected upgraded consensus state for the next version of this chain
// using the last height committed on this chain.
func (k Keeper) SetUpgradedConsensusState(ctx sdk.Context, planHeight int64, bz []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// GetUpgradedConsensusState set the expected upgraded consensus state for the next version of this chain
func (k Keeper) GetUpgradedConsensusState(ctx sdk.Context, lastHeight int64) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// GetLastCompletedUpgrade returns the last applied upgrade name and height.
func (k Keeper) GetLastCompletedUpgrade(ctx sdk.Context) (string, int64) {
	_ = "STUB: not implemented"
	return "", 0
}

//nolint:gosec // stored by SetDone from block heights which are always non-negative

// parseDoneKey - split upgrade name from the done key
func parseDoneKey(key []byte) string { _ = "STUB: not implemented"; return "" }

// GetDoneHeight returns the height at which the given upgrade was executed
func (k Keeper) GetDoneHeight(ctx sdk.Context, name string) int64 {
	_ = "STUB: not implemented"
	return 0
}

//nolint:gosec // stored by SetDone from block heights which are always non-negative

func (k Keeper) IsUpgradeActiveAtHeight(ctx sdk.Context, name string, height int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (k Keeper) GetClosestUpgrade(ctx sdk.Context, height int64) (string, int64) {
	_ = "STUB: not implemented"
	return "", 0
}

//nolint:gosec // stored by SetDone from block heights which are always non-negative

// ClearIBCState clears any planned IBC state
func (k Keeper) ClearIBCState(ctx sdk.Context, lastHeight int64) {
	_ = "STUB: not implemented"
	// delete IBC client and consensus state from store if this is IBC plan
	return
}

// ClearUpgradePlan clears any schedule upgrade and associated IBC states.
func (k Keeper) ClearUpgradePlan(ctx sdk.Context) {
	_ = "STUB: not implemented"
	// clear IBC states everytime upgrade plan is removed
	return
}

// GetUpgradePlan returns the currently scheduled Plan if any, setting havePlan to true if there is a scheduled
// upgrade or false if there is none
func (k Keeper) GetUpgradePlan(ctx sdk.Context) (plan types.Plan, havePlan bool) {
	_ = "STUB: not implemented"
	return *new(types.Plan), false
}

// SetDone marks this upgrade name as being done so the name can't be reused accidentally
func (k Keeper) SetDone(ctx sdk.Context, name string) { _ = "STUB: not implemented"; return }

//nolint:gosec // block heights are always non-negative

// HasHandler returns true iff there is a handler registered for this name
func (k Keeper) HasHandler(name string) bool { _ = "STUB: not implemented"; return false }

// ApplyUpgrade will execute the handler associated with the Plan and mark the plan as done.
func (k Keeper) ApplyUpgrade(ctx sdk.Context, plan types.Plan) { _ = "STUB: not implemented"; return }

// incremement the protocol version and set it in state and baseapp

// set protocol version on BaseApp

// Must clear IBC state after upgrade is applied as it is stored separately from the upgrade plan.
// This will prevent resubmission of upgrade msg after upgrade is already completed.

// IsSkipHeight checks if the given height is part of skipUpgradeHeights
func (k Keeper) IsSkipHeight(height int64) bool { _ = "STUB: not implemented"; return false }

// DumpUpgradeInfoToDisk writes upgrade information to UpgradeInfoFileName. The function
// doesn't save the `Plan.Info` data, hence it won't support auto download functionality
// by cosmvisor.
// NOTE: this function will be update in the next release.
func (k Keeper) DumpUpgradeInfoToDisk(height int64, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated: DumpUpgradeInfoWithInfoToDisk writes upgrade information to UpgradeInfoFileName.
// `info` should be provided and contain Plan.Info data in order to support
// auto download functionality by cosmovisor and other tools using upgrade-info.json
// (GetUpgradeInfoPath()) file.
func (k Keeper) DumpUpgradeInfoWithInfoToDisk(height int64, name string, info string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetUpgradeInfoPath returns the upgrade info file path
func (k Keeper) GetUpgradeInfoPath() (string, error) { _ = "STUB: not implemented"; return "", nil }

// getHomeDir returns the height at which the given upgrade was executed
func (k Keeper) getHomeDir() string {
	_ = "STUB: not implemented"

	// ReadUpgradeInfoFromDisk returns the name and height of the upgrade which is
	// written to disk by the old binary when panicking. An error is returned if
	// the upgrade path directory cannot be created or if the file exists and
	// cannot be read or if the upgrade info fails to unmarshal.
	return ""
}

func (k Keeper) ReadUpgradeInfoFromDisk() (store.UpgradeInfo, error) {
	_ = "STUB: not implemented"
	return *new(store.UpgradeInfo), nil
}

// if file does not exist, assume there are no upgrades

// upgradeInfo is stripped types.Plan structure used to dump upgrade plan data.
type upgradeInfo struct {
	// Name has types.Plan.Name value
	Name string `json:"name,omitempty"`
	// Height has types.Plan.Height value
	Height int64 `json:"height,omitempty"`
	// Info has types.Plan.Info value
	Info string `json:"info,omitempty"`
}

// SetDowngradeVerified updates downgradeVerified.
func (k *Keeper) SetDowngradeVerified(v bool) { _ = "STUB: not implemented"; return }

// DowngradeVerified returns downgradeVerified.
func (k Keeper) DowngradeVerified() bool { _ = "STUB: not implemented"; return false }

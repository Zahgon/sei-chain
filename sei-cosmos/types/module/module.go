/*
Package module contains application module patterns and associated "manager" functionality.
The module pattern has been broken down by:
  - independent module functionality (AppModuleBasic)
  - inter-dependent module genesis functionality (AppModuleGenesis)
  - inter-dependent module simulation functionality (AppModuleSimulation)
  - inter-dependent module full functionality (AppModule)

inter-dependent module functionality is module functionality which somehow
depends on other modules, typically through the module keeper.  Many of the
module keepers are dependent on each other, thus in order to access the full
set of module functionality we need to define all the keepers/params-store/keys
etc. This full set of advanced functionality is defined by the AppModule interface.

Independent module functions are separated to allow for the construction of the
basic application structures required early on in the application definition
and used to enable the definition of full module functionality later in the
process. This separation is necessary, however we still want to allow for a
high level pattern for modules to follow - for instance, such that we don't
have to manually register all of the codecs for all the modules. This basic
procedure as well as other basic patterns are handled through the use of
BasicManager.

Lastly the interface for genesis functionality (AppModuleGenesis) has been
separated out from full module functionality (AppModule) so that modules which
are only used for genesis can take advantage of the Module patterns without
needlessly defining many placeholder functions
*/
package module

import (
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/seilog"
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	genesistypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/genesis"
)

var logger = seilog.NewLogger("cosmos", "types", "module")

// AppModuleBasic is the standard form for basic non-dependant elements of an application module.
type AppModuleBasic interface {
	Name() string
	RegisterLegacyAminoCodec(*codec.LegacyAmino)
	RegisterInterfaces(codectypes.InterfaceRegistry)

	DefaultGenesis(codec.JSONCodec) json.RawMessage
	ValidateGenesis(codec.JSONCodec, client.TxEncodingConfig, json.RawMessage) error
	ValidateGenesisStream(codec.JSONCodec, client.TxEncodingConfig, <-chan json.RawMessage) error

	// client functionality
	RegisterRESTRoutes(client.Context, *mux.Router)
	RegisterGRPCGatewayRoutes(client.Context, *runtime.ServeMux)
	GetTxCmd() *cobra.Command
	GetQueryCmd() *cobra.Command
}

// BasicManager is a collection of AppModuleBasic
type BasicManager map[string]AppModuleBasic

// NewBasicManager creates a new BasicManager object
func NewBasicManager(modules ...AppModuleBasic) BasicManager {
	_ = "STUB: not implemented"
	return *new(BasicManager)
}

// RegisterLegacyAminoCodec registers all module codecs
func (bm BasicManager) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	_ = "STUB: not implemented"
	return
}

// RegisterInterfaces registers all module interface types
func (bm BasicManager) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	_ = "STUB: not implemented"
	return
}

// DefaultGenesis provides default genesis information for all modules
func (bm BasicManager) DefaultGenesis(cdc codec.JSONCodec) map[string]json.RawMessage {
	_ = "STUB: not implemented"
	return nil
}

// ValidateGenesis performs genesis state validation for all modules
func (bm BasicManager) ValidateGenesis(cdc codec.JSONCodec, txEncCfg client.TxEncodingConfig, genesis map[string]json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateGenesisStream performs genesis state validation on all modules in a streaming fashion.
func (bm BasicManager) ValidateGenesisStream(cdc codec.JSONCodec, txEncCfg client.TxEncodingConfig, moduleName string, genesisCh <-chan json.RawMessage, doneCh <-chan struct{}, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

// RegisterRESTRoutes registers all module rest routes
func (bm BasicManager) RegisterRESTRoutes(clientCtx client.Context, rtr *mux.Router) {
	_ = "STUB: not implemented"
	return
}

// RegisterGRPCGatewayRoutes registers all module rest routes
func (bm BasicManager) RegisterGRPCGatewayRoutes(clientCtx client.Context, rtr *runtime.ServeMux) {
	_ = "STUB: not implemented"
	return
}

// AddTxCommands adds all tx commands to the rootTxCmd.
//
// TODO: Remove clientCtx argument.
// REF: https://github.com/cosmos/cosmos-sdk/issues/6571
func (bm BasicManager) AddTxCommands(rootTxCmd *cobra.Command) { _ = "STUB: not implemented"; return }

// AddQueryCommands adds all query commands to the rootQueryCmd.
//
// TODO: Remove clientCtx argument.
// REF: https://github.com/cosmos/cosmos-sdk/issues/6571
func (bm BasicManager) AddQueryCommands(rootQueryCmd *cobra.Command) {
	_ = "STUB: not implemented"
	return
}

// AppModuleGenesis is the standard form for an application module genesis functions
type AppModuleGenesis interface {
	AppModuleBasic

	InitGenesis(sdk.Context, codec.JSONCodec, json.RawMessage) []abci.ValidatorUpdate
	ExportGenesis(sdk.Context, codec.JSONCodec) json.RawMessage
	ExportGenesisStream(ctx sdk.Context, cdc codec.JSONCodec) <-chan json.RawMessage
}

// AppModule is the standard form for an application module
type AppModule interface {
	AppModuleGenesis

	// registers
	RegisterInvariants(sdk.InvariantRegistry)

	// Deprecated: use RegisterServices
	Route() sdk.Route

	// Deprecated: use RegisterServices
	QuerierRoute() string

	// Deprecated: use RegisterServices
	LegacyQuerierHandler(*codec.LegacyAmino) sdk.Querier

	// RegisterServices allows a module to register services
	RegisterServices(Configurator)

	// ConsensusVersion is a sequence number for state-breaking change of the
	// module. It should be incremented on each consensus-breaking change
	// introduced by the module. To avoid wrong/empty versions, the initial version
	// should be set to 1.
	ConsensusVersion() uint64
}

// BeginBlockAppModule is an extension interface that contains information about the AppModule and BeginBlock.
type BeginBlockAppModule interface {
	AppModule
	BeginBlock(sdk.Context, abci.RequestBeginBlock)
}

type MidBlockAppModule interface {
	AppModule
	MidBlock(sdk.Context, int64)
}

// EndBlockAppModule is an extension interface that contains information about the AppModule and EndBlock.
type EndBlockAppModule interface {
	AppModule
	EndBlock(sdk.Context, abci.RequestEndBlock) []abci.ValidatorUpdate
}

// GenesisOnlyAppModule is an AppModule that only has import/export functionality
type GenesisOnlyAppModule struct {
	AppModuleGenesis
}

// NewGenesisOnlyAppModule creates a new GenesisOnlyAppModule object
func NewGenesisOnlyAppModule(amg AppModuleGenesis) AppModule {
	_ = "STUB: not implemented"
	return *new(AppModule)
}

// RegisterInvariants is a placeholder function register no invariants
func (GenesisOnlyAppModule) RegisterInvariants(_ sdk.InvariantRegistry) {
	_ = "STUB: not implemented"

	// Route empty module message route
	return
}

func (GenesisOnlyAppModule) Route() sdk.Route {
	_ = "STUB: not implemented"

	// QuerierRoute returns an empty module querier route
	return *new(sdk.Route)
}

func (GenesisOnlyAppModule) QuerierRoute() string {
	_ = "STUB: not implemented"

	// LegacyQuerierHandler returns an empty module querier
	return ""
}

func (gam GenesisOnlyAppModule) LegacyQuerierHandler(*codec.LegacyAmino) sdk.Querier {
	_ = "STUB: not implemented"

	// RegisterServices registers all services.
	return *new(sdk.Querier)
}

func (gam GenesisOnlyAppModule) RegisterServices(Configurator) {
	_ = "STUB: not implemented"

	// ConsensusVersion implements AppModule/ConsensusVersion.
	return
}

func (gam GenesisOnlyAppModule) ConsensusVersion() uint64 {
	_ = "STUB: not implemented"

	// Manager defines a module manager that provides the high level utility for managing and executing
	// operations for a group of modules
	return 0
}

type Manager struct {
	Modules            map[string]AppModule
	OrderInitGenesis   []string
	OrderExportGenesis []string
	OrderMidBlockers   []string
	OrderMigrations    []string
}

// NewManager creates a new Manager object
func NewManager(modules ...AppModule) *Manager { _ = "STUB: not implemented"; return nil }

// SetOrderInitGenesis sets the order of init genesis calls
func (m *Manager) SetOrderInitGenesis(moduleNames ...string) { _ = "STUB: not implemented"; return }

// SetOrderExportGenesis sets the order of export genesis calls
func (m *Manager) SetOrderExportGenesis(moduleNames ...string) { _ = "STUB: not implemented"; return }

// SetOrderMidBlockers sets the order of set mid-blocker calls
func (m *Manager) SetOrderMidBlockers(moduleNames ...string) { _ = "STUB: not implemented"; return }

// SetOrderMigrations sets the order of migrations to be run. If not set
// then migrations will be run with an order defined in `DefaultMigrationsOrder`.
func (m *Manager) SetOrderMigrations(moduleNames ...string) { _ = "STUB: not implemented"; return }

// RegisterInvariants registers all module invariants
func (m *Manager) RegisterInvariants(ir sdk.InvariantRegistry) { _ = "STUB: not implemented"; return }

// RegisterRoutes registers all module routes and module querier routes
func (m *Manager) RegisterRoutes(router sdk.Router, queryRouter sdk.QueryRouter, legacyQuerierCdc *codec.LegacyAmino) {
	_ = "STUB: not implemented"
	return
}

// RegisterServices registers all module services
func (m *Manager) RegisterServices(cfg Configurator) { _ = "STUB: not implemented"; return }

type AppState struct {
	Module string          `json:"module"`
	Data   json.RawMessage `json:"data"`
}

type ModuleState struct {
	AppState AppState `json:"app_state"`
}

func parseModule(jsonStr string) (*ModuleState, error) { _ = "STUB: not implemented"; return nil, nil }

// InitGenesis performs init genesis functionality for modules
func (m *Manager) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, genesisData map[string]json.RawMessage, genesisImportConfig genesistypes.GenesisImportConfig) abci.ResponseInitChain {
	_ = "STUB: not implemented"
	return *new(abci.ResponseInitChain)
}

// use these validator updates if provided, the module manager assumes
// only one module will update the validator set

// ExportGenesis performs export genesis functionality for modules
func (m *Manager) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) map[string]json.RawMessage {
	_ = "STUB: not implemented"
	return nil
}

func (m *Manager) ProcessGenesisPerModule(ctx sdk.Context, cdc codec.JSONCodec, process func(string, json.RawMessage) error) error {
	_ = "STUB: not implemented"
	// It's important that we use OrderInitGenesis here instead of OrderExportGenesis because the order of exporting
	// doesn't matter much but the order of importing does due to invariant checks and how we are streaming the genesis
	// file here
	return nil
}

// assertNoForgottenModules checks that we didn't forget any modules in the
// SetOrder* functions.
func (m *Manager) assertNoForgottenModules(setOrderFnName string, moduleNames []string) {
	_ = "STUB: not implemented"
	return
}

// MigrationHandler is the migration function that each module registers.
type MigrationHandler func(sdk.Context) error

// VersionMap is a map of moduleName -> version
type VersionMap map[string]uint64

// RunMigrations performs in-place store migrations for all modules. This
// function MUST be called insde an x/upgrade UpgradeHandler.
//
// Recall that in an upgrade handler, the `fromVM` VersionMap is retrieved from
// x/upgrade's store, and the function needs to return the target VersionMap
// that will in turn be persisted to the x/upgrade's store. In general,
// returning RunMigrations should be enough:
//
// Example:
//
//	cfg := module.NewConfigurator(...)
//	app.UpgradeKeeper.SetUpgradeHandler("my-plan", func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
//	    return app.mm.RunMigrations(ctx, cfg, fromVM)
//	})
//
// Internally, RunMigrations will perform the following steps:
// - create an `updatedVM` VersionMap of module with their latest ConsensusVersion
// - make a diff of `fromVM` and `udpatedVM`, and for each module:
//   - if the module's `fromVM` version is less than its `updatedVM` version,
//     then run in-place store migrations for that module between those versions.
//   - if the module does not exist in the `fromVM` (which means that it's a new module,
//     because it was not in the previous x/upgrade's store), then run
//     `InitGenesis` on that module.
//
// - return the `updatedVM` to be persisted in the x/upgrade's store.
//
// Migrations are run in an order defined by `Manager.OrderMigrations` or (if not set) defined by
// `DefaultMigrationsOrder` function.
//
// As an app developer, if you wish to skip running InitGenesis for your new
// module "foo", you need to manually pass a `fromVM` argument to this function
// foo's module version set to its latest ConsensusVersion. That way, the diff
// between the function's `fromVM` and `udpatedVM` will be empty, hence not
// running anything for foo.
//
// Example:
//
//	cfg := module.NewConfigurator(...)
//	app.UpgradeKeeper.SetUpgradeHandler("my-plan", func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
//	    // Assume "foo" is a new module.
//	    // `fromVM` is fetched from existing x/upgrade store. Since foo didn't exist
//	    // before this upgrade, `v, exists := fromVM["foo"]; exists == false`, and RunMigration will by default
//	    // run InitGenesis on foo.
//	    // To skip running foo's InitGenesis, you need set `fromVM`'s foo to its latest
//	    // consensus version:
//	    fromVM["foo"] = foo.AppModule{}.ConsensusVersion()
//
//	    return app.mm.RunMigrations(ctx, cfg, fromVM)
//	})
//
// Please also refer to docs/core/upgrade.md for more information.
func (m Manager) RunMigrations(ctx sdk.Context, cfg Configurator, fromVM VersionMap) (VersionMap, error) {
	_ = "STUB: not implemented"
	return *new(VersionMap), nil
}

// Only run migrations when the module exists in the fromVM.
// Run InitGenesis otherwise.
//
// the module won't exist in the fromVM in two cases:
// 1. A new module is added. In this case we run InitGenesis with an
// empty genesis state.
// 2. An existing chain is upgrading to v043 for the first time. In this case,
// all modules have yet to be added to x/upgrade's VersionMap store.

// Currently, the only implementator of Configurator (the interface)
// is configurator (the struct).

// The module manager assumes only one module will update the
// validator set, and that it will not be by a new module.

// EndBlock performs end block functionality for all modules. It creates a
// child context with an event manager to aggregate events emitted from all
// modules.
func (m *Manager) MidBlock(ctx sdk.Context, height int64) []abci.Event {
	_ = "STUB: not implemented"
	return nil
}

// GetVersionMap gets consensus version from all modules
func (m *Manager) GetVersionMap() VersionMap { _ = "STUB: not implemented"; return *new(VersionMap) }

// ModuleNames returns list of all module names, without any particular order.
func (m *Manager) ModuleNames() []string { _ = "STUB: not implemented"; return nil }

// DefaultMigrationsOrder returns a default migrations order: ascending alphabetical by module name,
// except x/auth which will run last, see:
// https://github.com/cosmos/cosmos-sdk/issues/10591
func DefaultMigrationsOrder(modules []string) []string { _ = "STUB: not implemented"; return nil }

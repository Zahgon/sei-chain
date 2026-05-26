package crisis

import (
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/crisis/keeper"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// Module init related flags
const (
	FlagSkipGenesisInvariants = "x-crisis-skip-assert-invariants"
)

// AppModuleBasic defines the basic application module used by the crisis module.
type AppModuleBasic struct{}

// Name returns the crisis module's name.
func (AppModuleBasic) Name() string { _ = "STUB: not implemented"; return "" }

// RegisterLegacyAminoCodec registers the crisis module's types on the given LegacyAmino codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	_ = "STUB: not implemented"
	return
}

// DefaultGenesis returns default genesis state as raw bytes for the crisis
// module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

// ValidateGenesis performs genesis state validation for the crisis module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func (am AppModuleBasic) ValidateGenesisStream(cdc codec.JSONCodec, config client.TxEncodingConfig, genesisCh <-chan json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterRESTRoutes registers no REST routes for the crisis module.
func (AppModuleBasic) RegisterRESTRoutes(_ client.Context, _ *mux.Router) {
	_ = "STUB: not implemented"

	// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the capability module.
	return
}

func (AppModuleBasic) RegisterGRPCGatewayRoutes(_ client.Context, _ *runtime.ServeMux) {
	_ = "STUB: not implemented"

	// GetTxCmd returns the root tx command for the crisis module.
	return
}

func (b AppModuleBasic) GetTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetQueryCmd returns no root query command for the crisis module.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	_ = "STUB: not implemented"

	// RegisterInterfaces registers interfaces and implementations of the crisis
	// module.
	return nil
}

func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	_ = "STUB: not implemented"
	return
}

// AppModule implements an application module for the crisis module.
type AppModule struct {
	AppModuleBasic

	// NOTE: We store a reference to the keeper here so that after a module
	// manager is created, the invariants can be properly registered and
	// executed.
	keeper *keeper.Keeper

	skipGenesisInvariants bool
}

// NewAppModule creates a new AppModule object. If initChainAssertInvariants is set,
// we will call keeper.AssertInvariants during InitGenesis (it may take a significant time)
// - which doesn't impact the chain security unless 66+% of validators have a wrongly
// modified genesis file.
func NewAppModule(keeper *keeper.Keeper, skipGenesisInvariants bool) AppModule {
	_ = "STUB: not implemented"
	return *new(AppModule)
}

// AddModuleInitFlags implements servertypes.ModuleInitFlags interface.
func AddModuleInitFlags(startCmd *cobra.Command) { _ = "STUB: not implemented"; return }

// Name returns the crisis module's name.
func (AppModule) Name() string { _ = "STUB: not implemented"; return "" }

// RegisterInvariants performs a no-op.
func (AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {
	_ = "STUB: not implemented"

	// Route returns the message routing key for the crisis module.
	return
}

func (am AppModule) Route() sdk.Route { _ = "STUB: not implemented"; return *new(sdk.Route) }

// QuerierRoute returns no querier route.
func (AppModule) QuerierRoute() string {
	_ = "STUB: not implemented"

	// LegacyQuerierHandler returns no sdk.Querier.
	return ""
}

func (AppModule) LegacyQuerierHandler(*codec.LegacyAmino) sdk.Querier {
	_ = "STUB: not implemented"

	// RegisterServices registers module services.
	return *new(sdk.Querier)
}

func (am AppModule) RegisterServices(cfg module.Configurator) { _ = "STUB: not implemented"; return }

// InitGenesis performs genesis initialization for the crisis module. It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	_ = "STUB: not implemented"
	return nil
}

// ExportGenesis returns the exported genesis state as raw bytes for the crisis
// module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func (am AppModule) ExportGenesisStream(ctx sdk.Context, cdc codec.JSONCodec) <-chan json.RawMessage {
	_ = "STUB: not implemented"
	return nil
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { _ = "STUB: not implemented"; return 0 }

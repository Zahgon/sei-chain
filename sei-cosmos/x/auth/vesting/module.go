package vesting

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
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/vesting/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic defines the basic application module used by the sub-vesting
// module. The module itself contain no special logic or state other than message
// handling.
type AppModuleBasic struct{}

// Name returns the module's name.
func (AppModuleBasic) Name() string { _ = "STUB: not implemented"; return "" }

// RegisterCodec registers the module's types with the given codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	_ = "STUB: not implemented"
	return
}

// RegisterInterfaces registers the module's interfaces and implementations with
// the given interface registry.
func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	_ = "STUB: not implemented"
	return
}

// DefaultGenesis returns the module's default genesis state as raw bytes.
func (AppModuleBasic) DefaultGenesis(_ codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"
	return *

	// ValidateGenesis performs genesis state validation. Currently, this is a no-op.
	new(json.RawMessage)
}

func (AppModuleBasic) ValidateGenesis(_ codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func (am AppModuleBasic) ValidateGenesisStream(cdc codec.JSONCodec, config client.TxEncodingConfig, genesisCh <-chan json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterRESTRoutes registers module's REST handlers. Currently, this is a no-op.
func (AppModuleBasic) RegisterRESTRoutes(_ client.Context, _ *mux.Router) {
	_ = "STUB: not implemented"

	// RegisterGRPCGatewayRoutes registers the module's gRPC Gateway routes. Currently, this
	// is a no-op.
	return
}

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(_ client.Context, _ *runtime.ServeMux) {
	_ = "STUB: not implemented"

	// GetTxCmd returns the root tx command for the auth module.
	return
}

func (AppModuleBasic) GetTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetQueryCmd returns the module's root query command. Currently, this is a no-op.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	_ = "STUB: not implemented"

	// AppModule extends the AppModuleBasic implementation by implementing the
	// AppModule interface.
	return nil
}

type AppModule struct {
	AppModuleBasic

	accountKeeper keeper.AccountKeeper
	bankKeeper    types.BankKeeper
}

func NewAppModule(ak keeper.AccountKeeper, bk types.BankKeeper) AppModule {
	_ = "STUB: not implemented"
	return *new(AppModule)
}

// RegisterInvariants performs a no-op; there are no invariants to enforce.
func (AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {
	_ = "STUB: not implemented"

	// Route returns the module's message router and handler.
	return
}

func (am AppModule) Route() sdk.Route { _ = "STUB: not implemented"; return *new(sdk.Route) }

// QuerierRoute returns an empty string as the module contains no query
// functionality.
func (AppModule) QuerierRoute() string {
	_ = "STUB: not implemented"

	// RegisterServices registers module services.
	return ""
}

func (am AppModule) RegisterServices(cfg module.Configurator) { _ = "STUB: not implemented"; return }

// LegacyQuerierHandler performs a no-op.
func (am AppModule) LegacyQuerierHandler(_ *codec.LegacyAmino) sdk.Querier {
	_ = "STUB: not implemented"

	// InitGenesis performs a no-op.
	return *new(sdk.Querier)
}

func (am AppModule) InitGenesis(_ sdk.Context, _ codec.JSONCodec, _ json.RawMessage) []abci.ValidatorUpdate {
	_ = "STUB: not implemented"
	return nil
}

// ExportGenesis is always empty, as InitGenesis does nothing either.
func (am AppModule) ExportGenesis(_ sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func (am AppModule) ExportGenesisStream(ctx sdk.Context, cdc codec.JSONCodec) <-chan json.RawMessage {
	_ = "STUB: not implemented"
	return nil
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { _ = "STUB: not implemented"; return 0 }

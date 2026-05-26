package ica

import (
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/27-interchain-accounts/controller"
	controllerkeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/27-interchain-accounts/controller/keeper"
	controllertypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/27-interchain-accounts/controller/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/27-interchain-accounts/host"
	hostkeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/27-interchain-accounts/host/keeper"
	hosttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/27-interchain-accounts/host/types"
	porttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/05-port/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}

	_ porttypes.IBCModule = controller.IBCModule{}
	_ porttypes.IBCModule = host.IBCModule{}
)

// AppModuleBasic is the IBC interchain accounts AppModuleBasic
type AppModuleBasic struct{}

// Name implements AppModuleBasic interface
func (AppModuleBasic) Name() string { _ = "STUB: not implemented"; return "" }

// RegisterLegacyAminoCodec implements AppModuleBasic.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	_ = "STUB: not implemented"

	// RegisterInterfaces registers module concrete types into protobuf Any
	return
}

func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	_ = "STUB: not implemented"
	return
}

// DefaultGenesis returns default genesis state as raw bytes for the IBC
// interchain accounts module
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

// ValidateGenesis performs genesis state validation for the IBC interchain acounts module
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func (am AppModuleBasic) ValidateGenesisStream(cdc codec.JSONCodec, config client.TxEncodingConfig, genesisCh <-chan json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterRESTRoutes implements AppModuleBasic interface
func (AppModuleBasic) RegisterRESTRoutes(ctx client.Context, rtr *mux.Router) {
	_ = "STUB: not implemented"

	// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the interchain accounts module.
	return
}

func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	_ = "STUB: not implemented"
	return
}

// GetTxCmd implements AppModuleBasic interface
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	_ = "STUB: not implemented"

	// GetQueryCmd implements AppModuleBasic interface
	return nil
}

func (AppModuleBasic) GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// AppModule is the application module for the IBC interchain accounts module
type AppModule struct {
	AppModuleBasic
	controllerKeeper *controllerkeeper.Keeper
	hostKeeper       *hostkeeper.Keeper
}

// NewAppModule creates a new IBC interchain accounts module
func NewAppModule(controllerKeeper *controllerkeeper.Keeper, hostKeeper *hostkeeper.Keeper) AppModule {
	_ = "STUB: not implemented"
	return *new(AppModule)
}

// InitModule will initialize the interchain accounts moudule. It should only be
// called once and as an alternative to InitGenesis.
func (am AppModule) InitModule(ctx sdk.Context, controllerParams controllertypes.Params, hostParams hosttypes.Params) {
	_ = "STUB: not implemented"
	return
}

// RegisterInvariants implements the AppModule interface
func (AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
	_ = "STUB: not implemented"

	// Route implements the AppModule interface
	return
}

func (AppModule) Route() sdk.Route { _ = "STUB: not implemented"; return *new(sdk.Route) }

// NewHandler implements the AppModule interface
func (AppModule) NewHandler() sdk.Handler {
	_ = "STUB: not implemented"

	// QuerierRoute implements the AppModule interface
	return *new(sdk.Handler)
}

func (AppModule) QuerierRoute() string { _ = "STUB: not implemented"; return "" }

// LegacyQuerierHandler implements the AppModule interface
func (am AppModule) LegacyQuerierHandler(legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	_ = "STUB: not implemented"

	// RegisterServices registers module services
	return *new(sdk.Querier)
}

func (am AppModule) RegisterServices(cfg module.Configurator) { _ = "STUB: not implemented"; return }

// InitGenesis performs genesis initialization for the interchain accounts module.
// It returns no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	_ = "STUB: not implemented"
	return nil
}

// ExportGenesis returns the exported genesis state as raw bytes for the interchain accounts module
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func (am AppModule) ExportGenesisStream(ctx sdk.Context, cdc codec.JSONCodec) <-chan json.RawMessage {
	_ = "STUB: not implemented"
	return nil
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 {
	_ = "STUB: not implemented"

	// BeginBlock implements the AppModule interface
	return 0
}

func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {
	_ = "STUB: not implemented"

	// EndBlock implements the AppModule interface
	return
}

func (am AppModule) EndBlock(ctx sdk.Context, req abci.RequestEndBlock) []abci.ValidatorUpdate {
	_ = "STUB: not implemented"
	return nil
}

package authz

import (
	"encoding/json"
	"math/rand"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/spf13/cobra"

	sdkclient "github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	cdctypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
	simtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/authz"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/authz/keeper"
)

var (
	_ module.AppModule           = AppModule{}
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.AppModuleSimulation = AppModule{}
)

// AppModuleBasic defines the basic application module used by the authz module.
type AppModuleBasic struct {
	cdc codec.Codec
}

// Name returns the authz module's name.
func (AppModuleBasic) Name() string { _ = "STUB: not implemented"; return "" }

// RegisterServices registers a gRPC query service to respond to the
// module-specific gRPC queries.
func (am AppModule) RegisterServices(cfg module.Configurator) { _ = "STUB: not implemented"; return }

// RegisterLegacyAminoCodec registers the authz module's types for the given codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	_ = "STUB: not implemented"

	// RegisterInterfaces registers the authz module's interface types
	return
}

func (AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	_ = "STUB: not implemented"
	return
}

// DefaultGenesis returns default genesis state as raw bytes for the authz
// module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

// ValidateGenesis performs genesis state validation for the authz module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config sdkclient.TxEncodingConfig, bz json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func (am AppModuleBasic) ValidateGenesisStream(cdc codec.JSONCodec, config sdkclient.TxEncodingConfig, genesisCh <-chan json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterRESTRoutes registers the REST routes for the authz module.
func (AppModuleBasic) RegisterRESTRoutes(clientCtx sdkclient.Context, r *mux.Router) {
	_ = "STUB: not implemented"

	// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the authz module.
	return
}

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx sdkclient.Context, mux *runtime.ServeMux) {
	_ = "STUB: not implemented"
	return
}

// GetQueryCmd returns the cli query commands for the authz module
func (AppModuleBasic) GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetTxCmd returns the transaction commands for the authz module
func (AppModuleBasic) GetTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// AppModule implements the sdk.AppModule interface
type AppModule struct {
	AppModuleBasic
	keeper        keeper.Keeper
	accountKeeper authz.AccountKeeper
	bankKeeper    authz.BankKeeper
	registry      cdctypes.InterfaceRegistry
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec, keeper keeper.Keeper, ak authz.AccountKeeper, bk authz.BankKeeper, registry cdctypes.InterfaceRegistry) AppModule {
	_ = "STUB: not implemented"
	return *new(AppModule)
}

// Name returns the authz module's name.
func (AppModule) Name() string { _ = "STUB: not implemented"; return "" }

// RegisterInvariants does nothing, there are no invariants to enforce
func (AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {
	_ = "STUB: not implemented"

	// Route returns the message routing key for the staking module.
	return
}

func (am AppModule) Route() sdk.Route { _ = "STUB: not implemented"; return *new(sdk.Route) }

func (am AppModule) NewHandler() sdk.Handler {
	_ = "STUB: not implemented"

	// QuerierRoute returns the route we respond to for abci queries
	return *new(sdk.Handler)
}

func (AppModule) QuerierRoute() string {
	_ = "STUB: not implemented"

	// LegacyQuerierHandler returns the authz module sdk.Querier.
	return ""
}

func (am AppModule) LegacyQuerierHandler(legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	_ = "STUB: not implemented"

	// InitGenesis performs genesis initialization for the authz module. It returns
	// no validator updates.
	return *new(sdk.Querier)
}

func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	_ = "STUB: not implemented"
	return nil
}

// ExportGenesis returns the exported genesis state as raw bytes for the authz
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
func (AppModule) ConsensusVersion() uint64 {
	_ = "STUB: not implemented"

	// ____________________________________________________________________________
	return 0
}

// AppModuleSimulation functions

// GenerateGenesisState creates a randomized GenState of the authz module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	_ = "STUB: not implemented"
	return
}

// ProposalContents returns all the authz content functions used to
// simulate governance proposals.
func (am AppModule) ProposalContents(simState module.SimulationState) []simtypes.WeightedProposalContent {
	_ = "STUB: not implemented"

	// RandomizedParams creates randomized authz param changes for the simulator.
	return nil
}

func (AppModule) RandomizedParams(r *rand.Rand) []simtypes.ParamChange {
	_ = "STUB: not implemented"

	// RegisterStoreDecoder registers a decoder for authz module's types
	return nil
}

func (am AppModule) RegisterStoreDecoder(sdr sdk.StoreDecoderRegistry) {
	_ = "STUB: not implemented"
	return
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	_ = "STUB: not implemented"
	return nil
}

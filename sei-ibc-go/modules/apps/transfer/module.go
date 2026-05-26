package transfer

import (
	"encoding/json"
	"math/rand"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
	simtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/transfer/keeper"
	porttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/05-port/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
	_ porttypes.IBCModule   = IBCModule{}
)

// AppModuleBasic is the IBC Transfer AppModuleBasic
type AppModuleBasic struct{}

// Name implements AppModuleBasic interface
func (AppModuleBasic) Name() string { _ = "STUB: not implemented"; return "" }

// RegisterLegacyAminoCodec implements AppModuleBasic interface
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	_ = "STUB: not implemented"
	return
}

// RegisterInterfaces registers module concrete types into protobuf Any.
func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	_ = "STUB: not implemented"
	return
}

// DefaultGenesis returns default genesis state as raw bytes for the ibc
// transfer module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

// ValidateGenesis performs genesis state validation for the ibc transfer module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateGenesisStream performs genesis state validation for the ibc transfer module in a streaming fashion.
func (am AppModuleBasic) ValidateGenesisStream(cdc codec.JSONCodec, config client.TxEncodingConfig, genesisCh <-chan json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterRESTRoutes implements AppModuleBasic interface
func (AppModuleBasic) RegisterRESTRoutes(clientCtx client.Context, rtr *mux.Router) {
	_ = "STUB: not implemented"

	// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the ibc-transfer module.
	return
}

func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	_ = "STUB: not implemented"
	return
}

// GetTxCmd implements AppModuleBasic interface
func (AppModuleBasic) GetTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetQueryCmd implements AppModuleBasic interface
func (AppModuleBasic) GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// AppModule represents the AppModule for this module
type AppModule struct {
	AppModuleBasic
	keeper keeper.Keeper
}

// NewAppModule creates a new 20-transfer module
func NewAppModule(k keeper.Keeper) AppModule { _ = "STUB: not implemented"; return *new(AppModule) }

// RegisterInvariants implements the AppModule interface
func (AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
	_ = "STUB: not implemented"
	// TODO

	// Route implements the AppModule interface
	return
}

func (am AppModule) Route() sdk.Route {
	_ = "STUB: not implemented"

	// QuerierRoute implements the AppModule interface
	return *new(sdk.Route)
}

func (AppModule) QuerierRoute() string { _ = "STUB: not implemented"; return "" }

// LegacyQuerierHandler implements the AppModule interface
func (am AppModule) LegacyQuerierHandler(*codec.LegacyAmino) sdk.Querier {
	_ = "STUB: not implemented"

	// RegisterServices registers module services.
	return *new(sdk.Querier)
}

func (am AppModule) RegisterServices(cfg module.Configurator) { _ = "STUB: not implemented"; return }

// InitGenesis performs genesis initialization for the ibc-transfer module. It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	_ = "STUB: not implemented"
	return nil
}

// ExportGenesis returns the exported genesis state as raw bytes for the ibc-transfer
// module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

// ExportGenesisStream returns the exported genesis state as raw bytes for the ibc-transfer module in a
// streaming fashion.
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

// AppModuleSimulation functions

// GenerateGenesisState creates a randomized GenState of the transfer module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	_ = "STUB: not implemented"
	return
}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	_ = "STUB: not implemented"

	// RandomizedParams creates randomized ibc-transfer param changes for the simulator.
	return nil
}

func (AppModule) RandomizedParams(r *rand.Rand) []simtypes.ParamChange {
	_ = "STUB: not implemented"
	return nil
}

// RegisterStoreDecoder registers a decoder for transfer module's types
func (am AppModule) RegisterStoreDecoder(sdr sdk.StoreDecoderRegistry) {
	_ = "STUB: not implemented"
	return
}

// WeightedOperations returns the all the transfer module operations with their respective weights.
func (am AppModule) WeightedOperations(_ module.SimulationState) []simtypes.WeightedOperation {
	_ = "STUB: not implemented"
	return nil
}

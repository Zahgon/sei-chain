/*
The tokenfactory module allows any account to create a new token with
the name `factory/{creator address}/{subdenom}`.

- Mint and burn user denom to and form any account
- Create a transfer of their denom between any two accounts
- Change the admin. In the future, more admin capabilities may be added.
*/
package tokenfactory

import (
	"encoding/json"
	"math/rand"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	cdctypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/spf13/cobra"

	simtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
	"github.com/sei-protocol/sei-chain/x/tokenfactory/keeper"
	"github.com/sei-protocol/sei-chain/x/tokenfactory/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// ----------------------------------------------------------------------------
// AppModuleBasic
// ----------------------------------------------------------------------------

// AppModuleBasic implements the AppModuleBasic interface for the capability module.
type AppModuleBasic struct{}

func NewAppModuleBasic() AppModuleBasic {
	_ = "STUB: not implemented"
	return *

	// Name returns the x/tokenfactory module's name.
	new(AppModuleBasic)
}

func (AppModuleBasic) Name() string { _ = "STUB: not implemented"; return "" }

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	_ = "STUB: not implemented"
	return
}

// RegisterInterfaces registers the module's interface types
func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	_ = "STUB: not implemented"
	return
}

// DefaultGenesis returns the x/tokenfactory module's default genesis state.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

// ValidateGenesis performs genesis state validation for the x/tokenfactory module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterRESTRoutes registers the capability module's REST service handlers.
func (AppModuleBasic) RegisterRESTRoutes(_ client.Context, _ *mux.Router) {
	_ = "STUB: not implemented"

	// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module.
	return
}

func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	_ = "STUB: not implemented"
	return
}

// GetTxCmd returns the x/tokenfactory module's root tx command.
func (a AppModuleBasic) GetTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetQueryCmd returns the x/tokenfactory module's root query command.
func (AppModuleBasic) GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// ----------------------------------------------------------------------------
// AppModule
// ----------------------------------------------------------------------------

// AppModule implements the AppModule interface for the capability module.
type AppModule struct {
	AppModuleBasic

	keeper        keeper.Keeper
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
}

func NewAppModule(
	keeper keeper.Keeper,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) AppModule {
	_ = "STUB: not implemented"
	return *new(AppModule)
}

// Name returns the x/tokenfactory module's name.
func (am AppModule) Name() string { _ = "STUB: not implemented"; return "" }

// Route returns the x/tokenfactory module's message routing key.
func (am AppModule) Route() sdk.Route {
	_ = "STUB: not implemented"

	// QuerierRoute returns the x/tokenfactory module's query routing key.
	return *new(sdk.Route)
}

func (AppModule) QuerierRoute() string { _ = "STUB: not implemented"; return "" }

// LegacyQuerierHandler returns the x/tokenfactory module's Querier.
func (am AppModule) LegacyQuerierHandler(_ *codec.LegacyAmino) sdk.Querier {
	_ = "STUB: not implemented"

	// RegisterServices registers a GRPC query service to respond to the
	// module-specific GRPC queries.
	return *new(sdk.Querier)
}

func (am AppModule) RegisterServices(cfg module.Configurator) { _ = "STUB: not implemented"; return }

// RegisterInvariants registers the x/tokenfactory module's invariants.
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {
	_ = "STUB: not implemented"

	// InitGenesis performs the x/tokenfactory module's genesis initialization. It
	// returns no validator updates.
	return
}

func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) []abci.ValidatorUpdate {
	_ = "STUB: not implemented"
	return nil
}

// ExportGenesis returns the x/tokenfactory module's exported genesis state as raw
// JSON bytes.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

// ExportGenesisStream returns the tokenfactory module's exported genesis state as raw JSON bytes in a streaming fashion.
func (am AppModule) ExportGenesisStream(ctx sdk.Context, cdc codec.JSONCodec) <-chan json.RawMessage {
	_ = "STUB: not implemented"
	return nil
}

// ValidateGenesisStream performs genesis state validation for the x/tokenfactory module in a streaming fashion.
func (am AppModuleBasic) ValidateGenesisStream(cdc codec.JSONCodec, config client.TxEncodingConfig, genesisCh <-chan json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// ConsensusVersion implements ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 {
	_ = "STUB: not implemented"

	// ___________________________________________________________________________
	return 0
}

// AppModuleSimulation functions

// GenerateGenesisState creates a randomized GenState of the tokenfactory module.
func (am AppModule) GenerateGenesisState(simState *module.SimulationState) {
	_ = "STUB: not implemented"
	return
}

// ProposalContents doesn't return any content functions for governance proposals.
func (am AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	_ = "STUB: not implemented"

	// RandomizedParams creates randomized txfees param changes for the simulator.
	return nil
}

func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	_ = "STUB: not implemented"

	// RegisterStoreDecoder registers a decoder for tokenfactory module's types
	return nil
}

func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {
	_ = "STUB: not implemented"

	// WeightedOperations returns simulator module operations with their respective weights.
	return
}

func (am AppModule) WeightedOperations(_ module.SimulationState) []simtypes.WeightedOperation {
	_ = "STUB: not implemented"
	return nil
}

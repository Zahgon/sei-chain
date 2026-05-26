package wasm

import (
	"encoding/json"
	"math/rand"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	cdctypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	servertypes "github.com/sei-protocol/sei-chain/sei-cosmos/server/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
	simtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
	simKeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/simulation"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/keeper"
	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// Module init related flags
const (
	flagWasmMemoryCacheSize    = "wasm.memory_cache_size"
	flagWasmQueryGasLimit      = "wasm.query_gas_limit"
	flagWasmSimulationGasLimit = "wasm.simulation_gas_limit"
)

// AppModuleBasic defines the basic application module used by the wasm module.
type AppModuleBasic struct{}

func (b AppModuleBasic) RegisterLegacyAminoCodec(amino *codec.LegacyAmino) {
	_ = "STUB: not implemented" //nolint:staticcheck
	return
}

func (b AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, serveMux *runtime.ServeMux) {
	_ = "STUB: not implemented"
	return
}

// Name returns the wasm module's name.
func (AppModuleBasic) Name() string {
	_ = "STUB: not implemented"

	// DefaultGenesis returns default genesis state as raw bytes for the wasm
	// module.
	return ""
}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

// ValidateGenesis performs genesis state validation for the wasm module.
func (b AppModuleBasic) ValidateGenesis(marshaler codec.JSONCodec, config client.TxEncodingConfig, message json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func (am AppModuleBasic) ValidateGenesisStream(cdc codec.JSONCodec, config client.TxEncodingConfig, genesisCh <-chan json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterRESTRoutes registers the REST routes for the wasm module.
func (AppModuleBasic) RegisterRESTRoutes(cliCtx client.Context, rtr *mux.Router) {
	_ = "STUB: not implemented"
	return
}

// GetTxCmd returns the root tx command for the wasm module.
func (b AppModuleBasic) GetTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetQueryCmd returns no root query command for the wasm module.
func (b AppModuleBasic) GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// RegisterInterfaces implements InterfaceModule
func (b AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	_ = "STUB: not implemented"
	return
}

// ____________________________________________________________________________

// AppModule implements an application module for the wasm module.
type AppModule struct {
	AppModuleBasic
	cdc                codec.Codec
	keeper             *Keeper
	validatorSetSource keeper.ValidatorSetSource
	accountKeeper      types.AccountKeeper // for simulation
	bankKeeper         simKeeper.BankKeeper
}

// ConsensusVersion is a sequence number for state-breaking change of the
// module. It should be incremented on each consensus-breaking change
// introduced by the module. To avoid wrong/empty versions, the initial version
// should be set to 1.
func (AppModule) ConsensusVersion() uint64 {
	_ = "STUB: not implemented"

	// NewAppModule creates a new AppModule object
	return 0
}

func NewAppModule(
	cdc codec.Codec,
	keeper *Keeper,
	validatorSetSource keeper.ValidatorSetSource,
	ak types.AccountKeeper,
	bk simKeeper.BankKeeper,
) AppModule {
	_ = "STUB: not implemented"
	return *new(AppModule)
}

func (am AppModule) RegisterServices(cfg module.Configurator) { _ = "STUB: not implemented"; return }

func (am AppModule) LegacyQuerierHandler(amino *codec.LegacyAmino) sdk.Querier {
	_ = "STUB: not implemented" //nolint:staticcheck
	return *new(sdk.Querier)
}

// RegisterInvariants registers the wasm module invariants.
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
	_ = "STUB: not implemented"

	// Route returns the message routing key for the wasm module.
	return
}

func (am AppModule) Route() sdk.Route { _ = "STUB: not implemented"; return *new(sdk.Route) }

// QuerierRoute returns the wasm module's querier route name.
func (AppModule) QuerierRoute() string { _ = "STUB: not implemented"; return "" }

// InitGenesis performs genesis initialization for the wasm module. It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	_ = "STUB: not implemented"
	return nil
}

// ExportGenesis returns the exported genesis state as raw bytes for the wasm
// module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func (am AppModule) ExportGenesisStream(ctx sdk.Context, cdc codec.JSONCodec) <-chan json.RawMessage {
	_ = "STUB: not implemented"
	return nil
}

// ____________________________________________________________________________
// AppModuleSimulation functions

// GenerateGenesisState creates a randomized GenState of the bank module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	_ = "STUB: not implemented"
	return
}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(simState module.SimulationState) []simtypes.WeightedProposalContent {
	_ = "STUB: not implemented"

	// RandomizedParams creates randomized bank param changes for the simulator.
	return nil
}

func (am AppModule) RandomizedParams(r *rand.Rand) []simtypes.ParamChange {
	_ = "STUB: not implemented"
	return nil
}

// RegisterStoreDecoder registers a decoder for supply module's types
func (am AppModule) RegisterStoreDecoder(sdr sdk.StoreDecoderRegistry) {
	_ = "STUB: not implemented"

	// WeightedOperations returns the all the gov module operations with their respective weights.
	return
}

func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	_ = "STUB: not implemented"
	return nil
}

// ____________________________________________________________________________
// AddModuleInitFlags implements servertypes.ModuleInitFlags interface.
func AddModuleInitFlags(startCmd *cobra.Command) { _ = "STUB: not implemented"; return }

// ReadWasmConfig reads the wasm specifig configuration
func ReadWasmConfig(opts servertypes.AppOptions) (types.WasmConfig, error) {
	_ = "STUB: not implemented"
	return *new(types.WasmConfig), nil
}

// non empty string set

// attach contract debugging to global "trace" flag

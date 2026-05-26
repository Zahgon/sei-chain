package genutil

import (
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	cdctypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/genutil/types"
)

var (
	_ module.AppModuleGenesis = AppModule{}
	_ module.AppModuleBasic   = AppModuleBasic{}
)

// AppModuleBasic defines the basic application module used by the genutil module.
type AppModuleBasic struct{}

// Name returns the genutil module's name.
func (AppModuleBasic) Name() string { _ = "STUB: not implemented"; return "" }

// RegisterLegacyAminoCodec registers the genutil module's types on the given LegacyAmino codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	_ = "STUB: not implemented"

	// RegisterInterfaces registers the module's interface types
	return
}

func (b AppModuleBasic) RegisterInterfaces(_ cdctypes.InterfaceRegistry) {
	_ = "STUB: not implemented"

	// DefaultGenesis returns default genesis state as raw bytes for the genutil
	// module.
	return
}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

// ValidateGenesis performs genesis state validation for the genutil module.
func (b AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, txEncodingConfig client.TxEncodingConfig, bz json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func (am AppModuleBasic) ValidateGenesisStream(cdc codec.JSONCodec, config client.TxEncodingConfig, genesisCh <-chan json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterRESTRoutes registers the REST routes for the genutil module.
func (AppModuleBasic) RegisterRESTRoutes(_ client.Context, _ *mux.Router) {
	_ = "STUB: not implemented"

	// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the genutil module.
	return
}

func (AppModuleBasic) RegisterGRPCGatewayRoutes(_ client.Context, _ *runtime.ServeMux) {
	_ = "STUB: not implemented"

	// GetTxCmd returns no root tx command for the genutil module.
	return
}

func (AppModuleBasic) GetTxCmd() *cobra.Command {
	_ = "STUB: not implemented"

	// GetQueryCmd returns no root query command for the genutil module.
	return nil
}

func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	_ = "STUB: not implemented"

	// AppModule implements an application module for the genutil module.
	return nil
}

type AppModule struct {
	AppModuleBasic

	accountKeeper    types.AccountKeeper
	stakingKeeper    types.StakingKeeper
	deliverTx        deliverTxfn
	txEncodingConfig client.TxEncodingConfig
}

// NewAppModule creates a new AppModule object
func NewAppModule(accountKeeper types.AccountKeeper,
	stakingKeeper types.StakingKeeper, deliverTx deliverTxfn,
	txEncodingConfig client.TxEncodingConfig,
) module.AppModule {
	_ = "STUB: not implemented"
	return *new(module.AppModule)
}

// InitGenesis performs genesis initialization for the genutil module.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	_ = "STUB: not implemented"
	return nil
}

// ExportGenesis returns the exported genesis state as raw bytes for the genutil
// module.
func (am AppModule) ExportGenesis(_ sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func (am AppModule) ExportGenesisStream(ctx sdk.Context, cdc codec.JSONCodec) <-chan json.RawMessage {
	_ = "STUB: not implemented"
	return nil
}

func (am AppModule) ValidateGenesisStream(cdc codec.JSONCodec, config client.TxEncodingConfig, genesisCh <-chan json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { _ = "STUB: not implemented"; return 0 }

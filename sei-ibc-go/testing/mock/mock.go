package mock

import (
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
	capabilitytypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/types"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/spf13/cobra"

	channeltypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/04-channel/types"
	porttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/05-port/types"
)

const (
	ModuleName = "mock"

	PortID  = ModuleName
	Version = "mock-version"
)

var (
	MockAcknowledgement             = channeltypes.NewResultAcknowledgement([]byte("mock acknowledgement"))
	MockFailAcknowledgement         = channeltypes.NewErrorAcknowledgement("mock failed acknowledgement")
	MockPacketData                  = []byte("mock packet data")
	MockFailPacketData              = []byte("mock failed packet data")
	MockAsyncPacketData             = []byte("mock async packet data")
	MockRecvCanaryCapabilityName    = "mock receive canary capability name"
	MockAckCanaryCapabilityName     = "mock acknowledgement canary capability name"
	MockTimeoutCanaryCapabilityName = "mock timeout canary capability name"
)

var _ porttypes.IBCModule = IBCModule{}

// Expected Interface
// PortKeeper defines the expected IBC port keeper
type PortKeeper interface {
	BindPort(ctx sdk.Context, portID string) *capabilitytypes.Capability
	IsBound(ctx sdk.Context, portID string) bool
}

// AppModuleBasic is the mock AppModuleBasic.
type AppModuleBasic struct{}

// Name implements AppModuleBasic interface.
func (AppModuleBasic) Name() string {
	_ = "STUB: not implemented"

	// RegisterLegacyAminoCodec implements AppModuleBasic interface.
	return ""
}

func (AppModuleBasic) RegisterLegacyAminoCodec(*codec.LegacyAmino) {
	_ = "STUB: not implemented"

	// RegisterInterfaces implements AppModuleBasic interface.
	return
}

func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	_ = "STUB: not implemented"

	// DefaultGenesis implements AppModuleBasic interface.
	return
}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"

	// ValidateGenesis implements the AppModuleBasic interface.
	return *new(json.RawMessage)
}

func (AppModuleBasic) ValidateGenesis(codec.JSONCodec, client.TxEncodingConfig, json.RawMessage) error {
	_ = "STUB: not implemented"

	// ValidateGenesisStream implements the AppModuleBasic interface.
	return nil
}

func (am AppModuleBasic) ValidateGenesisStream(cdc codec.JSONCodec, config client.TxEncodingConfig, genesisCh <-chan json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterRESTRoutes implements AppModuleBasic interface.
func (AppModuleBasic) RegisterRESTRoutes(clientCtx client.Context, rtr *mux.Router) {
	_ = "STUB: not implemented"

	// RegisterGRPCGatewayRoutes implements AppModuleBasic interface.
	return
}

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(_ client.Context, _ *runtime.ServeMux) {
	_ = "STUB: not implemented"

	// GetTxCmd implements AppModuleBasic interface.
	return
}

func (AppModuleBasic) GetTxCmd() *cobra.Command {
	_ = "STUB: not implemented"

	// GetQueryCmd implements AppModuleBasic interface.
	return nil
}

func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	_ = "STUB: not implemented"

	// AppModule represents the AppModule for the mock module.
	return nil
}

type AppModule struct {
	AppModuleBasic
	ibcApps    []*MockIBCApp
	portKeeper PortKeeper
}

// NewAppModule returns a mock AppModule instance.
func NewAppModule(pk PortKeeper) AppModule { _ = "STUB: not implemented"; return *new(AppModule) }

// RegisterInvariants implements the AppModule interface.
func (AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
	_ = "STUB: not implemented"

	// Route implements the AppModule interface.
	return
}

func (am AppModule) Route() sdk.Route { _ = "STUB: not implemented"; return *new(sdk.Route) }

// QuerierRoute implements the AppModule interface.
func (AppModule) QuerierRoute() string {
	_ = "STUB: not implemented"

	// LegacyQuerierHandler implements the AppModule interface.
	return ""
}

func (am AppModule) LegacyQuerierHandler(*codec.LegacyAmino) sdk.Querier {
	_ = "STUB: not implemented"

	// RegisterServices implements the AppModule interface.
	return *new(sdk.Querier)
}

func (am AppModule) RegisterServices(module.Configurator) {
	_ = "STUB: not implemented"

	// InitGenesis implements the AppModule interface.
	return
}

func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	_ = "STUB: not implemented"
	return nil
}

// bind mock portID

// ExportGenesis implements the AppModule interface.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"

	// ExportGenesisStream implements the AppModule interface.
	return *new(json.RawMessage)
}

func (am AppModule) ExportGenesisStream(ctx sdk.Context, cdc codec.JSONCodec) <-chan json.RawMessage {
	_ = "STUB: not implemented"

	// ConsensusVersion implements AppModule/ConsensusVersion.
	return nil
}

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

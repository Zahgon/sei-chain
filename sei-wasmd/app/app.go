package app

import (
	"io"
	"os"

	"github.com/sei-protocol/sei-chain/sei-cosmos/baseapp"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/server/api"
	"github.com/sei-protocol/sei-chain/sei-cosmos/server/config"
	servertypes "github.com/sei-protocol/sei-chain/sei-cosmos/server/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth"
	authkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/keeper"
	authtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/vesting"
	authzkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/authz/keeper"
	authzmodule "github.com/sei-protocol/sei-chain/sei-cosmos/x/authz/module"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/bank"
	bankkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/capability"
	capabilitykeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/crisis"
	crisiskeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/crisis/keeper"
	distr "github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution"
	distrclient "github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution/client"
	distrkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution/keeper"
	distrtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/evidence"
	evidencekeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/evidence/keeper"
	feegrantkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/feegrant/keeper"
	feegrantmodule "github.com/sei-protocol/sei-chain/sei-cosmos/x/feegrant/module"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/genutil"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/gov"
	govkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/keeper"
	govtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/params"
	paramsclient "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/client"
	paramskeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/keeper"
	paramstypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/slashing"
	slashingkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/slashing/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/staking"
	stakingkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/keeper"
	stakingtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade"
	upgradeclient "github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade/client"
	upgradekeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade/keeper"
	icacontrollerkeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/27-interchain-accounts/controller/keeper"
	icahostkeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/27-interchain-accounts/host/keeper"
	icatypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/27-interchain-accounts/types"
	transfer "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/transfer"
	ibctransferkeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/transfer/types"
	ibc "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core"
	ibcclientclient "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/client"
	ibckeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/keeper"
	tmcfg "github.com/sei-protocol/sei-chain/sei-tendermint/config"
	"github.com/sei-protocol/sei-chain/x/mint"
	mintkeeper "github.com/sei-protocol/sei-chain/x/mint/keeper"
	minttypes "github.com/sei-protocol/sei-chain/x/mint/types"
	"github.com/sei-protocol/seilog"

	"github.com/gorilla/mux"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	dbm "github.com/tendermint/tm-db"

	wasmappparams "github.com/sei-protocol/sei-chain/sei-wasmd/app/params"
	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm"
	wasmclient "github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/client"

	// unnamed import of statik for swagger UI support
	_ "github.com/sei-protocol/sei-chain/sei-cosmos/client/docs/statik"
)

const appName = "WasmApp"

// We pull these out so we can set them with LDFLAGS in the Makefile
var (
	logger = seilog.NewLogger("wasmd", "app")

	NodeDir      = ".wasmd"
	Bech32Prefix = "wasm"

	// If EnabledSpecificProposals is "", and this is "true", then enable all x/wasm proposals.
	// If EnabledSpecificProposals is "", and this is not "true", then disable all x/wasm proposals.
	ProposalsEnabled = "false"
	// If set to non-empty string it must be comma-separated list of values that are all a subset
	// of "EnableAllProposals" (takes precedence over ProposalsEnabled)
	// https://github.com/CosmWasm/wasmd/blob/02a54d33ff2c064f3539ae12d75d027d9c665f05/x/wasm/internal/types/proposal.go#L28-L34
	EnableSpecificProposals = ""
)

// GetEnabledProposals parses the ProposalsEnabled / EnableSpecificProposals values to
// produce a list of enabled proposals to pass into wasmd app.
func GetEnabledProposals() []wasm.ProposalType { _ = "STUB: not implemented"; return nil }

// These constants are derived from the above variables.
// These are the ones we will want to use in the code, based on
// any overrides above
var (
	// DefaultNodeHome default home directories for wasmd
	DefaultNodeHome = os.ExpandEnv("$HOME/") + NodeDir

	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address
	Bech32PrefixAccAddr = Bech32Prefix
	// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key
	Bech32PrefixAccPub = Bech32Prefix + sdk.PrefixPublic
	// Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address
	Bech32PrefixValAddr = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator
	// Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key
	Bech32PrefixValPub = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator + sdk.PrefixPublic
	// Bech32PrefixConsAddr defines the Bech32 prefix of a consensus node address
	Bech32PrefixConsAddr = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus
	// Bech32PrefixConsPub defines the Bech32 prefix of a consensus node public key
	Bech32PrefixConsPub = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus + sdk.PrefixPublic
)

var (
	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(
			append(
				wasmclient.ProposalHandlers,
				paramsclient.ProposalHandler,
				distrclient.ProposalHandler,
				upgradeclient.ProposalHandler,
				upgradeclient.CancelProposalHandler,
				ibcclientclient.UpdateClientProposalHandler,
				ibcclientclient.UpgradeProposalHandler,
			)...,
		),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		ibc.AppModuleBasic{},
		transfer.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		vesting.AppModuleBasic{},
		wasm.AppModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:     nil,
		distrtypes.ModuleName:          nil,
		minttypes.ModuleName:           {authtypes.Minter},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:            {authtypes.Burner},
		ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
		icatypes.ModuleName:            nil,
		wasm.ModuleName:                {authtypes.Burner},
	}
)

var (
	_ servertypes.Application = (*WasmApp)(nil)
)

// WasmApp extended ABCI application
type WasmApp struct {
	*baseapp.BaseApp
	legacyAmino       *codec.LegacyAmino //nolint:staticcheck
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*sdk.KVStoreKey
	tkeys   map[string]*sdk.TransientStoreKey
	memKeys map[string]*sdk.MemoryStoreKey

	// keepers
	accountKeeper       authkeeper.AccountKeeper
	bankKeeper          bankkeeper.Keeper
	capabilityKeeper    *capabilitykeeper.Keeper
	stakingKeeper       stakingkeeper.Keeper
	slashingKeeper      slashingkeeper.Keeper
	mintKeeper          mintkeeper.Keeper
	distrKeeper         distrkeeper.Keeper
	govKeeper           govkeeper.Keeper
	crisisKeeper        crisiskeeper.Keeper
	upgradeKeeper       upgradekeeper.Keeper
	paramsKeeper        paramskeeper.Keeper
	evidenceKeeper      evidencekeeper.Keeper
	ibcKeeper           *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	icaControllerKeeper icacontrollerkeeper.Keeper
	icaHostKeeper       icahostkeeper.Keeper
	transferKeeper      ibctransferkeeper.Keeper
	feeGrantKeeper      feegrantkeeper.Keeper
	authzKeeper         authzkeeper.Keeper
	wasmKeeper          wasm.Keeper

	scopedIBCKeeper           capabilitykeeper.ScopedKeeper
	scopedICAHostKeeper       capabilitykeeper.ScopedKeeper
	scopedICAControllerKeeper capabilitykeeper.ScopedKeeper
	scopedTransferKeeper      capabilitykeeper.ScopedKeeper
	scopedWasmKeeper          capabilitykeeper.ScopedKeeper

	// the module manager
	mm *module.Manager

	// module configurator
	configurator module.Configurator

	txDecoder sdk.TxDecoder
}

// NewWasmApp returns a reference to an initialized WasmApp.
func NewWasmApp(
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	tmConfig *tmcfg.Config,
	encodingConfig wasmappparams.EncodingConfig,
	enabledProposals []wasm.ProposalType,
	appOpts servertypes.AppOptions,
	wasmOpts []wasm.Option,
	baseAppOptions ...func(*baseapp.BaseApp),
) *WasmApp {
	_ = "STUB: not implemented"
	return nil
}

// set the BaseApp's parameter store

// add capability keeper and ScopeToModule for ibc module

// add keepers

// register the staking hooks
// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks

// register the proposal types

// Create Transfer Keepers

// may be replaced with middleware such as ics29 fee

// For wasmd we use the demo controller from https://github.com/cosmos/interchain-accounts but see notes below
// Note: please do your research before using this in production app, this is a demo and not an officially
// supported IBC team implementation. Do your own research before using it.
// You will likely want to swap out the second argument with your own reviewed and maintained ica auth module

// create evidence keeper with router

// The last arguments can contain custom message handlers, and custom query handlers,
// if we want to allow any custom callbacks

// Create static IBC router, add app routes, then set and seal it

// The gov proposal types can be individually enabled

/****  Module Options ****/

// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
// we prefer to be more strict in what arguments the modules expect.

// NOTE: Any module instantiated in the module manager that is later modified
// must be passed by reference here.

// always be last to make sure that it checks for all invariants and not only part of them

// NOTE: The genutils module must occur after staking so that pools are
// properly initialized with tokens from genesis accounts.
// NOTE: Capability module must occur first so that it can initialize any capabilities
// so that other modules that want to create or claim capabilities afterwards in InitChain
// can do so safely.
// NOTE: wasm module should be at the end as it can call other module functionality direct or via message dispatching during
// genesis phase. For example bank transfer, auth account check, staking, ...

// additional non simd modules

// wasm after ibc transfer

// Uncomment if you want to set a custom migration order here.
// app.mm.SetOrderMigrations(custom order)

// initialize stores

// must be before Loading version
// requires the snapshot store to be created and registered as a BaseAppOption
// see cmd/wasmd/root.go: 206 - 214 approx

// Initialize pinned codes in wasmvm as they are not persisted there

// Name returns the name of the App
func (app *WasmApp) Name() string { _ = "STUB: not implemented"; return "" }

func (app *WasmApp) ProcessProposalHandler(ctx sdk.Context, req *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *WasmApp) FinalizeBlocker(ctx sdk.Context, req *abci.RequestFinalizeBlock) (*abci.ResponseFinalizeBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// InitChainer application update at chain initialization
func (app *WasmApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	_ = "STUB: not implemented"
	return *new(abci.ResponseInitChain)
}

func (app *WasmApp) EndBlocker(ctx sdk.Context) []abci.ValidatorUpdate {
	_ = "STUB: not implemented"
	return nil
}

// LoadHeight loads a particular height
func (app *WasmApp) LoadHeight(height int64) error { _ = "STUB: not implemented"; return nil }

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *WasmApp) ModuleAccountAddrs() map[string]bool { _ = "STUB: not implemented"; return nil }

// LegacyAmino returns legacy amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *WasmApp) LegacyAmino() *codec.LegacyAmino {
	_ = "STUB: not implemented" //nolint:staticcheck
	return nil
}

// getSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *WasmApp) getSubspace(moduleName string) paramstypes.Subspace {
	_ = "STUB: not implemented"
	return *new(paramstypes.Subspace)
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *WasmApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	_ = "STUB: not implemented"
	return
}

// Register legacy tx routes.

// Register new tx routes from grpc-gateway.

// Register new tendermint queries routes from grpc-gateway.

// Register legacy and grpc-gateway routes for all modules.

// register swagger API from root so that other applications can override easily

// RegisterTxService implements the Application.RegisterLocalServices method.
func (app *WasmApp) RegisterLocalServices(node client.LocalClient, txConfig client.TxConfig) {
	_ = "STUB: not implemented"
	return
}

func (app *WasmApp) AppCodec() codec.Codec { _ = "STUB: not implemented"; return *new(codec.Codec) }

func (app *WasmApp) GetCapabilityKeeper() *capabilitykeeper.Keeper {
	_ = "STUB: not implemented"
	return nil
}

func (app *WasmApp) GetDistrKeeper() *distrkeeper.Keeper { _ = "STUB: not implemented"; return nil }

func (app *WasmApp) GetSlashingKeeper() *slashingkeeper.Keeper {
	_ = "STUB: not implemented"
	return nil
}

func (app *WasmApp) GetEvidenceKeeper() *evidencekeeper.Keeper {
	_ = "STUB: not implemented"
	return nil

	// RegisterSwaggerAPI registers swagger route with API Server
}

func RegisterSwaggerAPI(rtr *mux.Router) { _ = "STUB: not implemented"; return }

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string { _ = "STUB: not implemented"; return nil }

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey sdk.StoreKey) paramskeeper.Keeper {
	_ = "STUB: not implemented"
	return *new(paramskeeper.Keeper)
}

package simapp

import (
	"io"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	tmcfg "github.com/sei-protocol/sei-chain/sei-tendermint/config"
	dbm "github.com/tendermint/tm-db"

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
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/bank"
	bankkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/keeper"
	banktypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/capability"
	capabilitykeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/keeper"
	simappparams "github.com/sei-protocol/sei-chain/sei-ibc-go/testing/simapp/params"

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
	ica "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/27-interchain-accounts"
	icacontrollerkeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/27-interchain-accounts/controller/keeper"
	icahostkeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/27-interchain-accounts/host/keeper"
	icatypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/27-interchain-accounts/types"
	transfer "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/transfer"
	ibctransferkeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/transfer/types"
	ibc "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core"
	ibcclientclient "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/client"
	ibckeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/keeper"
	ibcmock "github.com/sei-protocol/sei-chain/sei-ibc-go/testing/mock"

	authzkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/authz/keeper"
	authzmodule "github.com/sei-protocol/sei-chain/sei-cosmos/x/authz/module"

	// unnamed import of statik for swagger UI support
	_ "github.com/sei-protocol/sei-chain/sei-cosmos/client/docs/statik"
)

const appName = "SimApp"

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(
			paramsclient.ProposalHandler, distrclient.ProposalHandler, upgradeclient.ProposalHandler, upgradeclient.CancelProposalHandler,
			ibcclientclient.UpdateClientProposalHandler, ibcclientclient.UpgradeProposalHandler,
		),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		ibc.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		transfer.AppModuleBasic{},
		ibcmock.AppModuleBasic{},
		ica.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		vesting.AppModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:     nil,
		distrtypes.ModuleName:          nil,
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:            {authtypes.Burner},
		ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
		banktypes.ModuleName:           {authtypes.Minter, authtypes.Burner},
		icatypes.ModuleName:            nil,
	}
)

var (
	_ App                     = (*SimApp)(nil)
	_ servertypes.Application = (*SimApp)(nil)
)

// SimApp extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type SimApp struct {
	*baseapp.BaseApp
	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*sdk.KVStoreKey
	tkeys   map[string]*sdk.TransientStoreKey
	memKeys map[string]*sdk.MemoryStoreKey

	// keepers
	AccountKeeper       authkeeper.AccountKeeper
	BankKeeper          bankkeeper.Keeper
	CapabilityKeeper    *capabilitykeeper.Keeper
	StakingKeeper       stakingkeeper.Keeper
	SlashingKeeper      slashingkeeper.Keeper
	DistrKeeper         distrkeeper.Keeper
	GovKeeper           govkeeper.Keeper
	CrisisKeeper        crisiskeeper.Keeper
	UpgradeKeeper       upgradekeeper.Keeper
	ParamsKeeper        paramskeeper.Keeper
	AuthzKeeper         authzkeeper.Keeper
	IBCKeeper           *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	ICAControllerKeeper icacontrollerkeeper.Keeper
	ICAHostKeeper       icahostkeeper.Keeper
	EvidenceKeeper      evidencekeeper.Keeper
	TransferKeeper      ibctransferkeeper.Keeper
	FeeGrantKeeper      feegrantkeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper           capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper      capabilitykeeper.ScopedKeeper
	ScopedICAControllerKeeper capabilitykeeper.ScopedKeeper
	ScopedICAHostKeeper       capabilitykeeper.ScopedKeeper
	ScopedIBCMockKeeper       capabilitykeeper.ScopedKeeper
	ScopedICAMockKeeper       capabilitykeeper.ScopedKeeper

	// make IBC modules public for test purposes
	// these modules are never directly routed to by the IBC Router
	ICAAuthModule ibcmock.IBCModule

	// the module manager
	mm *module.Manager

	// the configurator
	configurator module.Configurator

	txDecoder sdk.TxDecoder
}

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, ".simapp")
}

// NewSimApp returns a reference to an initialized SimApp.
func NewSimApp(
	db dbm.DB, traceStore io.Writer, loadLatest bool, skipUpgradeHeights map[int64]bool,
	homePath string, invCheckPeriod uint, tmConfig *tmcfg.Config, encodingConfig simappparams.EncodingConfig,
	appOpts servertypes.AppOptions, baseAppOptions ...func(*baseapp.BaseApp),
) *SimApp {
	_ = "STUB: not implemented"
	return nil
}

// set the BaseApp's parameter store

// add capability keeper and ScopeToModule for ibc module

// NOTE: the IBC mock keeper and application module is used only for testing core IBC. Do
// not replicate if you do not need to test core IBC or light clients.

// seal capability keeper after scoping modules

// add keepers

// register the staking hooks
// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks

// Create IBC Keeper

// register the proposal types

// Create Transfer Keepers

// NOTE: the IBC mock keeper and application module is used only for testing core IBC. Do
// not replicate if you do not need to test core IBC or light clients.

// may be replaced with middleware such as ics29 fee

// initialize ICA module with mock module as the authentication module on the controller side

// Create static IBC router, add app routes, then set and seal it

// ica with mock auth module stack route to ica (top level of middleware stack)

// create evidence keeper with router

// If evidence needs to be handled for the app, set routes in router here and seal

/****  Module Options ****/

// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
// we prefer to be more strict in what arguments the modules expect.

// NOTE: Any module instantiated in the module manager that is later modified
// must be passed by reference here.

// NOTE: The genutils module must occur after staking so that pools are
// properly initialized with tokens from genesis accounts.
// NOTE: Capability module must occur first so that it can initialize any capabilities
// so that other modules that want to create or claim capabilities afterwards in InitChain
// can do so safely.

// add test gRPC service for testing gRPC queries in isolation

// initialize stores

// initialize BaseApp

// NOTE: the IBC mock keeper and application module is used only for testing core IBC. Do
// note replicate if you do not need to test core IBC or light clients.

// Name returns the name of the App
func (app *SimApp) Name() string { _ = "STUB: not implemented"; return "" }

func (app *SimApp) FinalizeBlocker(ctx sdk.Context, req *abci.RequestFinalizeBlock) (*abci.ResponseFinalizeBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BeginBlocker application updates every begin block
func (app *SimApp) BeginBlocker(ctx sdk.Context) { _ = "STUB: not implemented"; return }

// EndBlocker application updates every end block
func (app *SimApp) EndBlocker(ctx sdk.Context) []abci.ValidatorUpdate {
	_ = "STUB: not implemented"
	return nil
}

// InitChainer application update at chain initialization
func (app *SimApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	_ = "STUB: not implemented"
	return *new(abci.ResponseInitChain)
}

// LoadHeight loads a particular height
func (app *SimApp) LoadHeight(height int64) error { _ = "STUB: not implemented"; return nil }

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *SimApp) ModuleAccountAddrs() map[string]bool { _ = "STUB: not implemented"; return nil }

// GetModuleManager returns the app module manager
// NOTE: used for testing purposes
func (app *SimApp) GetModuleManager() *module.Manager {
	_ = "STUB: not implemented"

	// LegacyAmino returns SimApp's amino codec.
	//
	// NOTE: This is solely to be used for testing purposes as it may be desirable
	// for modules to register their own custom testing types.
	return nil
}

func (app *SimApp) LegacyAmino() *codec.LegacyAmino { _ = "STUB: not implemented"; return nil }

// AppCodec returns SimApp's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *SimApp) AppCodec() codec.Codec {
	_ = "STUB: not implemented"
	return *

	// InterfaceRegistry returns SimApp's InterfaceRegistry
	new(codec.Codec)
}

func (app *SimApp) InterfaceRegistry() types.InterfaceRegistry {
	_ = "STUB: not implemented"
	return *new(types.InterfaceRegistry)
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *SimApp) GetKey(storeKey string) *sdk.KVStoreKey { _ = "STUB: not implemented"; return nil }

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *SimApp) GetTKey(storeKey string) *sdk.TransientStoreKey {
	_ = "STUB: not implemented"
	return nil

	// GetMemKey returns the MemStoreKey for the provided mem key.
	//
	// NOTE: This is solely used for testing purposes.
}

func (app *SimApp) GetMemKey(storeKey string) *sdk.MemoryStoreKey {
	_ = "STUB: not implemented"
	return nil
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *SimApp) GetSubspace(moduleName string) paramstypes.Subspace {
	_ = "STUB: not implemented"
	return *new(paramstypes.Subspace)
}

// TestingApp functions

// GetBaseApp implements the TestingApp interface.
func (app *SimApp) GetBaseApp() *baseapp.BaseApp {
	_ = "STUB: not implemented"

	// GetStakingKeeper implements the TestingApp interface.
	return nil
}

func (app *SimApp) GetStakingKeeper() stakingkeeper.Keeper {
	_ = "STUB: not implemented"
	return *

	// GetIBCKeeper implements the TestingApp interface.
	new(stakingkeeper.Keeper)
}

func (app *SimApp) GetIBCKeeper() *ibckeeper.Keeper { _ = "STUB: not implemented"; return nil }

// GetScopedIBCKeeper implements the TestingApp interface.
func (app *SimApp) GetScopedIBCKeeper() capabilitykeeper.ScopedKeeper {
	_ = "STUB: not implemented"
	return *new(capabilitykeeper.ScopedKeeper)
}

// GetTxConfig implements the TestingApp interface.
func (app *SimApp) GetTxConfig() client.TxConfig {
	_ = "STUB: not implemented"
	return *new(client.TxConfig)
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *SimApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	_ = "STUB: not implemented"
	return
}

// Register legacy tx routes.

// Register new tx routes from grpc-gateway.

// Register new tendermint queries routes from grpc-gateway.

// Register legacy and grpc-gateway routes for all modules.

// register swagger API from root so that other applications can override easily

// RegisterTxService implements the Application.RegisterLocalServices method.
func (app *SimApp) RegisterLocalServices(node client.LocalClient, txConfig client.TxConfig) {
	_ = "STUB: not implemented"
	return
}

// RegisterSwaggerAPI registers swagger route with API Server
func RegisterSwaggerAPI(ctx client.Context, rtr *mux.Router) { _ = "STUB: not implemented"; return }

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string { _ = "STUB: not implemented"; return nil }

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey sdk.StoreKey) paramskeeper.Keeper {
	_ = "STUB: not implemented"
	return *new(paramskeeper.Keeper)
}

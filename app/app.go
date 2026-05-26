package app

import (
	"context"
	"io"
	"math/big"
	"os"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	ethparams "github.com/ethereum/go-ethereum/params"
	"github.com/sei-protocol/sei-chain/admin"
	"github.com/sei-protocol/sei-chain/sei-cosmos/baseapp"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/server/api"
	"github.com/sei-protocol/sei-chain/sei-cosmos/server/config"
	servertypes "github.com/sei-protocol/sei-chain/sei-cosmos/server/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	genesistypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/genesis"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
	"github.com/sei-protocol/sei-chain/sei-cosmos/version"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth"
	authkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/keeper"
	authtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/vesting"
	authzkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/authz/keeper"
	authzmodule "github.com/sei-protocol/sei-chain/sei-cosmos/x/authz/module"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/bank"
	bankkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/keeper"
	banktypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/capability"
	capabilitykeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/keeper"
	capabilitytypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/crisis"
	crisiskeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/crisis/keeper"
	distr "github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution"
	distrkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution/keeper"
	distrtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/evidence"
	evidencekeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/evidence/keeper"
	evidencetypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/evidence/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/feegrant"
	feegrantkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/feegrant/keeper"
	feegrantmodule "github.com/sei-protocol/sei-chain/sei-cosmos/x/feegrant/module"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/genutil"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/gov"
	govclient "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/client"
	govkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/keeper"
	govtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/params"
	paramskeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/keeper"
	paramstypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/slashing"
	slashingkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/slashing/keeper"
	slashingtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/slashing/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/staking"
	stakingkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/keeper"
	stakingtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade"
	upgradekeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade/keeper"
	upgradetypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade/types"
	seidb "github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
	"github.com/sei-protocol/seilog"
	"google.golang.org/grpc"

	"github.com/gorilla/mux"
	"github.com/sei-protocol/sei-chain/app/benchmark"
	"github.com/sei-protocol/sei-chain/app/legacyabci"
	appparams "github.com/sei-protocol/sei-chain/app/params"
	"github.com/sei-protocol/sei-chain/app/upgrades"
	"github.com/sei-protocol/sei-chain/evmrpc"
	evmrpcconfig "github.com/sei-protocol/sei-chain/evmrpc/config"
	putils "github.com/sei-protocol/sei-chain/precompiles/utils"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/transfer"
	ibctransferkeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/transfer/types"
	ibc "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core"
	ibchost "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/24-host"
	ibckeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/keeper"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	tmcfg "github.com/sei-protocol/sei-chain/sei-tendermint/config"
	tmutils "github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"

	epochmodule "github.com/sei-protocol/sei-chain/x/epoch"
	epochmodulekeeper "github.com/sei-protocol/sei-chain/x/epoch/keeper"
	epochmoduletypes "github.com/sei-protocol/sei-chain/x/epoch/types"
	"github.com/sei-protocol/sei-chain/x/evm"
	evmkeeper "github.com/sei-protocol/sei-chain/x/evm/keeper"
	evmtypes "github.com/sei-protocol/sei-chain/x/evm/types"
	"github.com/sei-protocol/sei-chain/x/mint"
	mintkeeper "github.com/sei-protocol/sei-chain/x/mint/keeper"
	minttypes "github.com/sei-protocol/sei-chain/x/mint/types"
	oraclemodule "github.com/sei-protocol/sei-chain/x/oracle"
	oraclekeeper "github.com/sei-protocol/sei-chain/x/oracle/keeper"
	oracletypes "github.com/sei-protocol/sei-chain/x/oracle/types"
	tokenfactorymodule "github.com/sei-protocol/sei-chain/x/tokenfactory"
	tokenfactorykeeper "github.com/sei-protocol/sei-chain/x/tokenfactory/keeper"
	tokenfactorytypes "github.com/sei-protocol/sei-chain/x/tokenfactory/types"
	dbm "github.com/tendermint/tm-db"

	// this line is used by starport scaffolding # stargate/app/moduleImport

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm"
	wasmtypes "github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"

	// unnamed import of statik for openapi/swagger UI support
	_ "github.com/sei-protocol/sei-chain/docs/swagger"
	receipt "github.com/sei-protocol/sei-chain/sei-db/ledger_db/receipt"

	gigabankkeeper "github.com/sei-protocol/sei-chain/giga/deps/xbank/keeper"
	gigaevmkeeper "github.com/sei-protocol/sei-chain/giga/deps/xevm/keeper"
)

// this line is used by starport scaffolding # stargate/wasm/app/enabledProposals
func getGovProposalHandlers() []govclient.ProposalHandler { _ = "STUB: not implemented"; return nil }

//nolint:gocritic // ignore: appending to a slice is OK

// this line is used by starport scaffolding # stargate/app/govProposalHandler

var (
	logger = seilog.NewLogger("app")

	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// upgradePanicRe matches upgrade panic messages using Cosmovisor-compatible regex
	// Matches multiple upgrade-related panic patterns:
	// 1. UPGRADE "name" NEEDED at height: 123 (or height123)
	// 2. Wrong app version X, upgrade handler is missing for name upgrade plan
	// 3. BINARY UPDATED BEFORE TRIGGER! UPGRADE "name"
	upgradePanicRe = regexp.MustCompile(`^(UPGRADE "[^"]+" NEEDED at height:?\s*\d+|Wrong app version \d+, upgrade handler is missing for .+ upgrade plan|BINARY UPDATED BEFORE TRIGGER! UPGRADE "[^"]+")`)

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(getGovProposalHandlers()...),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		ibc.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		transfer.AppModuleBasic{},
		vesting.AppModuleBasic{},
		oraclemodule.AppModuleBasic{},
		evm.AppModuleBasic{},
		wasm.AppModuleBasic{},
		epochmodule.AppModuleBasic{},
		tokenfactorymodule.AppModuleBasic{},
		// this line is used by starport scaffolding # stargate/app/moduleBasic
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
		oracletypes.ModuleName:         nil,
		wasm.ModuleName:                {authtypes.Burner},
		evmtypes.ModuleName:            {authtypes.Minter, authtypes.Burner},
		tokenfactorytypes.ModuleName:   {authtypes.Minter, authtypes.Burner},
		// this line is used by starport scaffolding # stargate/app/maccPerms
	}

	allowedReceivingModAcc = map[string]bool{
		oracletypes.ModuleName: true,
	}

	// kvStoreKeyNames is the canonical, in-order list of module KV store
	// names mounted on the SeiDB / memiavl backend. It is the single source
	// of truth consumed by sdk.NewKVStoreKeys in app.New, and is cross-
	// checked against keys.MemIAVLStoreKeys (sei-db/common/keys) by tests
	// in this package. Adding or removing an entry here MUST be matched in
	// sei-db/common/keys/store_keys.go.
	kvStoreKeyNames = []string{
		authtypes.StoreKey, authzkeeper.StoreKey, banktypes.StoreKey, stakingtypes.StoreKey,
		minttypes.StoreKey, distrtypes.StoreKey, slashingtypes.StoreKey,
		govtypes.StoreKey, paramstypes.StoreKey, ibchost.StoreKey, upgradetypes.StoreKey, feegrant.StoreKey,
		evidencetypes.StoreKey, ibctransfertypes.StoreKey, capabilitytypes.StoreKey, oracletypes.StoreKey,
		evmtypes.StoreKey, wasm.StoreKey,
		epochmoduletypes.StoreKey,
		tokenfactorytypes.StoreKey,
		// this line is used by starport scaffolding # stargate/app/storeKey
	}

	// WasmProposalsEnabled enables all x/wasm proposals when it's value is "true"
	// and EnableSpecificWasmProposals is empty. Otherwise, all x/wasm proposals
	// are disabled.
	// Used as a flag to turn it on and off
	WasmProposalsEnabled = "true"

	// EnableSpecificWasmProposals, if set, must be comma-separated list of values
	// that are all a subset of "EnableAllProposals", which takes precedence over
	// WasmProposalsEnabled.
	//
	// See: https://github.com/CosmWasm/wasmd/blob/02a54d33ff2c064f3539ae12d75d027d9c665f05/x/wasm/internal/types/proposal.go#L28-L34
	EnableSpecificWasmProposals = ""

	// EmptyWasmOpts defines a type alias for a list of wasm options.
	EmptyWasmOpts []wasm.Option

	// Boolean to only emit seid version and git commit metric once per chain initialization
	EmittedSeidVersionMetric = false
	// EnableOCC allows tests to override default OCC enablement behavior
	EnableOCC       = true
	EmptyAppOptions []AppOption
)

var (
	_ servertypes.Application = (*App)(nil)
)

const (
	MinGasEVMTx = 21000

	// NewHeadsNotifierCapacity bounds the in-process eth_newHeads
	// notifier buffer. Capacity 1 pairs with the notifier's
	// overwrite-on-full semantics: if a consumer lags, the latest head
	// always wins and stale heads are dropped. Anything larger only
	// buffers staleness — newHeads subscribers care about the current
	// head, not a backlog.
	NewHeadsNotifierCapacity = 1
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+AppName)
}

// GetWasmEnabledProposals parses the WasmProposalsEnabled and
// EnableSpecificWasmProposals values to produce a list of enabled proposals to
// pass into the application.
func GetWasmEnabledProposals() []wasm.ProposalType { _ = "STUB: not implemented"; return nil }

// App extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
// gigaBlockCache holds block-constant values that are identical for all txs in a block.
// Populated once before block execution, read-only during parallel execution, cleared after.
type gigaBlockCache struct {
	chainID     *big.Int
	blockCtx    vm.BlockContext
	chainConfig *ethparams.ChainConfig
	baseFee     *big.Int
}

func newGigaBlockCache(ctx sdk.Context, keeper *gigaevmkeeper.Keeper) (*gigaBlockCache, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type App struct {
	*baseapp.BaseApp

	cdc               *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*sdk.KVStoreKey
	tkeys   map[string]*sdk.TransientStoreKey
	memKeys map[string]*sdk.MemoryStoreKey

	// keepers
	AccountKeeper    authkeeper.AccountKeeper
	AuthzKeeper      authzkeeper.Keeper
	BankKeeper       bankkeeper.Keeper
	GigaBankKeeper   *gigabankkeeper.BaseKeeper
	CapabilityKeeper *capabilitykeeper.Keeper
	StakingKeeper    stakingkeeper.Keeper
	SlashingKeeper   slashingkeeper.Keeper
	MintKeeper       mintkeeper.Keeper
	DistrKeeper      distrkeeper.Keeper
	GovKeeper        govkeeper.Keeper
	CrisisKeeper     crisiskeeper.Keeper
	UpgradeKeeper    upgradekeeper.Keeper
	ParamsKeeper     paramskeeper.Keeper
	IBCKeeper        *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	EvidenceKeeper   evidencekeeper.Keeper
	TransferKeeper   ibctransferkeeper.Keeper
	FeeGrantKeeper   feegrantkeeper.Keeper
	WasmKeeper       wasm.Keeper
	OracleKeeper     oraclekeeper.Keeper
	EvmKeeper        evmkeeper.Keeper
	GigaEvmKeeper    gigaevmkeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper      capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper capabilitykeeper.ScopedKeeper
	ScopedWasmKeeper     capabilitykeeper.ScopedKeeper

	EpochKeeper epochmodulekeeper.Keeper

	TokenFactoryKeeper tokenfactorykeeper.Keeper

	BeginBlockKeepers legacyabci.BeginBlockKeepers
	EndBlockKeepers   legacyabci.EndBlockKeepers
	CheckTxKeepers    legacyabci.CheckTxKeepers
	DeliverTxKeepers  legacyabci.DeliverTxKeepers

	// mm is the module manager
	mm *module.Manager

	// sm is the simulation manager
	sm *module.SimulationManager

	configurator module.Configurator

	optimisticProcessingInfo      OptimisticProcessingInfo
	optimisticProcessingInfoMutex sync.RWMutex

	txDecoder         sdk.TxDecoder
	AnteHandler       sdk.AnteHandler
	TracerAnteHandler sdk.AnteHandler

	versionInfo version.Info

	// Stores mapping counter name to counter value
	metricCounter *map[string]float32

	mounter func()

	HardForkManager *upgrades.HardForkManager

	encodingConfig       appparams.EncodingConfig
	legacyEncodingConfig appparams.EncodingConfig
	evmRPCConfig         evmrpcconfig.Config
	// blockHeaderNotifier is non-nil only when Autobahn is enabled. It
	// owns the FinalizeBlock→Commit pairing for eth_subscribe("newHeads"):
	// FinalizeBlocker calls Stash with the (hash, header, response)
	// tuple, App.Commit calls PublishStashed after a successful
	// BaseApp.Commit, and FinalizeBlocker entry calls ClearStash to
	// defend against stale tuples from prior failed commits or
	// non-stashing return paths (EthReplay/EthBlockTest).
	blockHeaderNotifier   tmutils.Option[*evmrpc.BlockHeaderNotifier]
	adminConfig           admin.Config
	adminServer           *grpc.Server
	lightInvarianceConfig LightInvarianceConfig

	genesisImportConfig genesistypes.GenesisImportConfig

	stateStore   seidb.StateStore
	receiptStore receipt.ReceiptStore

	forkInitializer func(sdk.Context)

	httpServerStartSignal     chan struct{}
	wsServerStartSignal       chan struct{}
	httpServerStartSignalSent bool
	wsServerStartSignalSent   bool

	txPrioritizer sdk.TxPrioritizer

	benchmarkManager *benchmark.Manager

	// GigaExecutorEnabled controls whether to use the Giga executor.
	GigaExecutorEnabled bool
	// GigaOCCEnabled controls whether to use OCC with the Giga executor
	GigaOCCEnabled bool
}

type AppOption func(*App)

// New returns a reference to an initialized blockchain app
func New(
	db dbm.DB,
	traceStore io.Writer,
	_ bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	enableCustomEVMPrecompiles bool,
	tmConfig *tmcfg.Config,
	encodingConfig appparams.EncodingConfig,
	enabledProposals []wasm.ProposalType,
	appOpts servertypes.AppOptions,
	wasmOpts []wasm.Option,
	appOptions []AppOption,
	baseAppOptions ...func(*baseapp.BaseApp),
) *App {
	_ = "STUB: not implemented"
	return nil
}

// Bind OTEL metrics provider once at application construction

// set the BaseApp's parameter store

// add capability keeper and ScopeToModule for ibc module

// grant capabilities for the ibc and ibc-transfer modules

// this line is used by starport scaffolding # stargate/app/scopedKeeper

// add keepers

// register the staking hooks
// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks

// ... other modules keepers

// Create IBC Keeper

// Create Transfer Keepers

// Create evidence Keeper for to register the IBC light client misbehaviour evidence route

// If evidence needs to be handled for the app, set routes in router here and seal

// The last arguments can contain custom message handlers, and custom query handlers,
// if we want to allow any custom callbacks

// TODO: remove the mode gate and always construct the notifier once
// non-Autobahn is also producer-wired to feed it (i.e. a listener
// invocation in sei-tendermint/internal/state/execution.go next to
// PublishEventNewBlockHeader). That switch is blocked on parity work:
// the legacy event-bus path encodes a real Tendermint Header (real
// parentHash/receiptsRoot/transactionsRoot, pre-execution stateRoot)
// while encodeCommittedBlock zeroes those fields and uses
// post-execution stateRoot. We need to either verify both encoders
// produce identical headers for the same block under legacy, or
// reconcile the encoder so swapping the consumer is a no-op for
// non-Autobahn subscribers. Until that's verified, keep this gate
// so non-Autobahn newHeads semantics are unchanged by this PR.

// Read Giga Executor config

// evmone is loaded best-effort

// register the proposal types

// this line is used by starport scaffolding # stargate/app/keeperDefinition

// Create static IBC router, add transfer route, then set and seal it

// this line is used by starport scaffolding # ibc/app/router

/****  Module Options ****/

// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
// we prefer to be more strict in what arguments the modules expect.

// NOTE: Any module instantiated in the module manager that is later modified
// must be passed by reference here.

// this line is used by starport scaffolding # stargate/app/appModule

// NOTE: The genutils module must occur after staking so that pools are
// properly initialized with tokens from genesis accounts.
// NOTE: Capability module must occur first so that it can initialize any capabilities
// so that other modules that want to create or claim capabilities afterwards in InitChain
// can do so safely.

// this line is used by starport scaffolding # stargate/app/initGenesis

// create the simulation manager and define the order of the modules for deterministic simulations

// this line is used by starport scaffolding # stargate/app/appModule

// initialize stores

// initialize BaseApp

// BatchVerifier:   app.batchVerifier,

// benchmarkEnabled is enabled via build flag (make install-bench)

// Register snapshot extensions to enable state-sync for wasm.

// Create hard fork manager and register all hard fork upgrade handlers. Note,
// when creating the manager, BaseApp must already be instantiated.
//
// example: app.HardForkManager.RegisterHandler(myHandler)

// HandlePreCommit happens right before the block is committed
func (app *App) HandlePreCommit(ctx sdk.Context) error { _ = "STUB: not implemented"; return nil }

// Close closes all items that needs closing (called by baseapp)
func (app *App) HandleClose() error {
	_ = "STUB: not implemented"

	// Close trace db so its WAL is flushed; baker writes use NoSync.
	return nil
}

// Close receipt store

// Stop admin gRPC server

// Note: stateStore (ssStore) is already closed by cms.Close() in BaseApp.Close()
// No need to close it again here.

// Add (or remove) keepers when they are introduced / removed in different versions
func (app *App) SetStoreUpgradeHandlers() { _ = "STUB: not implemented"; return }

// configure store loader that checks if version == upgradeHeight and applies store upgrades

// configure store loader that checks if version == upgradeHeight and applies store upgrades

// configure store loader that checks if version == upgradeHeight and applies store upgrades

// configure store loader that checks if version == upgradeHeight and applies store upgrades

// configure store loader that checks if version == upgradeHeight and applies store upgrades

// configure store loader that checks if version == upgradeHeight and applies store upgrades

// configure store loader that checks if version == upgradeHeight and applies store upgrades

// configure store loader that checks if version == upgradeHeight and applies store upgrades

// AppName returns the name of the App
func (app *App) Name() string { _ = "STUB: not implemented"; return "" }

// GetBaseApp returns the base app of the application
func (app *App) GetBaseApp() *baseapp.BaseApp {
	_ = "STUB: not implemented"

	// GetStateStore returns the state store of the application
	return nil
}

func (app *App) GetStateStore() seidb.StateStore {
	_ = "STUB: not implemented"
	return *

	// MidBlocker application updates every mid block
	new(seidb.StateStore)
}

func (app *App) MidBlocker(ctx sdk.Context, height int64) []abci.Event {
	_ = "STUB: not implemented"
	return nil
}

// InitChainer application update at chain initialization
func (app *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	_ = "STUB: not implemented"
	return *new(abci.ResponseInitChain)
}

func (app *App) GetOptimisticProcessingInfo() OptimisticProcessingInfo {
	_ = "STUB: not implemented"
	return *new(OptimisticProcessingInfo)
}

func (app *App) ClearOptimisticProcessingInfo() { _ = "STUB: not implemented"; return }

func (app *App) ProcessProposalHandler(ctx sdk.Context, req *abci.RequestProcessProposal) (resp *abci.ResponseProcessProposal, err error) {
	_ = "STUB: not implemented"
	// Start block processing timing (ends at FinalizeBlock)
	return nil, nil
}

// TODO(PLT-327): remove once app_failed_total_gas_wanted_check_total verified

// Use the clean context for gas validation only. We cannot reassign
// ctx because ProcessBlock writes to ctx's store downstream.

// Invariant: at this point nil entries in typedTxs are proto decode failures only.
// EVM preprocessing runs later inside ProcessBlock; do not reorder that before this check.

// TODO(PLT-327): remove once app_failed_total_gas_wanted_check_total verified

// ProcessBlock has panic recovery and returns error for any processing failures
// All panics (including GetSigners) are handled in ProcessBlock, not affecting proposal acceptance

// ProcessBlock failed (including GetSigners panics), mark as aborted

// ProcessBlock succeeded, store results

// Optimistic processing already running, check if hash matches

func (app *App) FinalizeBlocker(ctx sdk.Context, req *abci.RequestFinalizeBlock) (*abci.ResponseFinalizeBlock, error) {
	_ = "STUB: not implemented"
	// Drop any leftover stash so only the current FinalizeBlock can be
	// published by the next Commit. Defends against return paths that
	// don't Stash (EthReplay/EthBlockTest) and prior Commit failures.
	return nil, nil
}

// End block processing timing (started at ProcessProposal)

// Process receipts for benchmark deployment tracking

// Get all optimistic processing info atomically

// Get the final state atomically after completion

// TODO(PLT-327): remove once app_optimistic_processing_total verified

// TODO(PLT-327): remove once app_optimistic_processing_total verified

func (app *App) DeliverTxWithResult(ctx sdk.Context, tx []byte, typedTx sdk.Tx) *abci.ExecTxResult {
	_ = "STUB: not implemented"
	return nil
}

// Check if transaction is gasless before recording metrics
// perf optimization: skip gasless check for obviously non-gasless transaction types

// Only do expensive validation for potentially gasless transactions

// ErrAggregateVoteExist is expected when checking gasless status after tx processing
// since oracle votes will now exist in state. We know it was gasless, skip metrics.

// If we can't determine if it's gasless, record metrics to maintain existing behavior

// Skip metrics for confirmed gasless transactions

// Record metrics for non-gasless transactions
// TODO(PLT-327): remove once app_tx_gas_total verified
// TODO(PLT-327): remove once app_tx_gas_total verified

func (app *App) ProcessTxsSynchronousV2(ctx sdk.Context, txs [][]byte, typedTxs []sdk.Tx) []*abci.ExecTxResult {
	_ = "STUB: not implemented"
	return nil
}

// TODO(PLT-327): remove once app_block_process_duration_seconds verified

// TODO(PLT-327): remove once app_tx_process_type_total verified

func (app *App) ProcessTxsSynchronousGiga(ctx sdk.Context, txs [][]byte, typedTxs []sdk.Tx) []*abci.ExecTxResult {
	_ = "STUB: not implemented"
	return nil
}

// TODO(PLT-327): remove once app_block_process_duration_seconds verified

// Cache block-level constants (identical for all txs in this block).

// If not an EVM tx, fall back to v2 processing

// Execute EVM transaction through giga executor with panic recovery.
// Matches V2's recover behavior in legacyabci/deliver_tx.go.

// IIFE (immediately-invoked function) to scope defer/recover to this tx only,
// allowing the loop to continue processing subsequent transactions after a panic.

// Handle panics by type (matches V2's recovery middleware in baseapp/recovery.go)

// For other panics (e.g., nil deref from malformed protobuf), log and return ErrPanic

// Validate Cosmos SDK envelope (memo, timeoutHeight, signerInfos, etc.)
// This prevents consensus divergence if a malicious proposer includes invalid envelope fields.

// Check if this is a fail-fast error (Cosmos precompile interop detected)

// TODO(PLT-327): remove once app_tx_process_type_total verified

func (app *App) shouldProcessSingleRecipientEVMTransfersSynchronously(typedTxs []sdk.Tx) bool {
	_ = "STUB: not implemented"
	return false
}

// cacheContext returns a new context based off of the provided context with
// a branched multi-store.
func (app *App) CacheContext(ctx sdk.Context) (sdk.Context, sdk.CacheMultiStore) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), *new(sdk.CacheMultiStore)
}

// ExecuteTxsConcurrently calls the appropriate function for processing transacitons
func (app *App) ExecuteTxsConcurrently(ctx sdk.Context, txs [][]byte, typedTxs []sdk.Tx) ([]*abci.ExecTxResult, sdk.Context) {
	_ = "STUB: not implemented"
	return nil, *new(sdk.Context)
}

func (app *App) GetDeliverTxEntry(ctx sdk.Context, txIndex int, bz []byte, tx sdk.Tx) (res *sdk.DeliverTxEntry) {
	_ = "STUB: not implemented"
	return nil
}

// ProcessTXsWithOCCV2 runs the transactions concurrently via OCC, using the V2 executor
func (app *App) ProcessTXsWithOCCV2(ctx sdk.Context, txs [][]byte, typedTxs []sdk.Tx) ([]*abci.ExecTxResult, sdk.Context) {
	_ = "STUB: not implemented"
	return nil, *new(sdk.Context)
}

// TODO(PLT-327): remove once app_tx_process_type_total verified

// Check if transaction is gasless before recording gas metrics

// perf optimization: skip gasless check for obviously non-gasless transaction types

// Only do expensive validation for potentially gasless transactions

// ErrAggregateVoteExist is expected when checking gasless status after tx processing
// since oracle votes will now exist in state. We know it was gasless, skip metrics.

// If we can't determine if it's gasless, record metrics to maintain existing behavior

// TODO(PLT-327): remove once app_tx_gas_total verified
// TODO(PLT-327): remove once app_tx_gas_total verified

// ProcessTXsWithOCCGiga runs the transactions concurrently via OCC, using the Giga executor
func (app *App) ProcessTXsWithOCCGiga(ctx sdk.Context, txs [][]byte, typedTxs []sdk.Tx) ([]*abci.ExecTxResult, sdk.Context) {
	_ = "STUB: not implemented"
	return nil, *new(sdk.Context)
}

// Oops! This isn't "all EVM txs, then all Cosmos txs" - we need to fallback to V2.

// Run EVM txs against a cache so we can discard all changes on fallback.

// Cache block-level constants (identical for all txs in this block).
// Must use evmCtx (not ctx) because giga KV stores are registered in CacheContext.

// Create OCC scheduler with giga executor deliverTx capturing the cache.

// TODO(PLT-327): remove once app_giga_fallback_to_v2_total verified

// Discard all EVM changes by skipping cache writes, then re-run all txs via DeliverTx.

// Commit EVM cache to main store before processing non-EVM txs.

// ProcessBlock executes block transactions. If preDecoded is non-nil and len(preDecoded)==len(txs),
// those decoded transactions are reused (bytes are not decoded again); EVM preprocessing still runs
// on the block context.
func (app *App) ProcessBlock(ctx sdk.Context, txs [][]byte, req *BlockProcessRequest, lastCommit abci.CommitInfo, simulate bool, preDecoded []sdk.Tx) (events []abci.Event, txResults []*abci.ExecTxResult, endBlockResp abci.ResponseEndBlock, err error) {
	_ = "STUB: not implemented"
	return nil, nil, *new(abci.ResponseEndBlock), nil
}

// Re-panic for upgrade-related panics to allow proper upgrade mechanism

// Re-panic to trigger upgrade mechanism

// nil for non-EVM txs

// Execute all transactions

// Flush giga stores so WriteDeferredBalances (which uses the standard BankKeeper)
// can see balance changes made by the giga executor via GigaBankKeeper.

// Finalize all Bank Module Transfers here so that events are included

// Sum up total used per block only for evm transactions

// executeEVMTxWithGigaExecutor executes a single EVM transaction using the giga executor.
// The sender address is recovered directly from the transaction signature - no Cosmos SDK ante handlers needed.
func (app *App) executeEVMTxWithGigaExecutor(ctx sdk.Context, msg *evmtypes.MsgEVMTransaction, cache *gigaBlockCache) (*abci.ExecTxResult, error) {
	_ = "STUB: not implemented"
	// Get the Ethereum transaction from the message
	return nil, nil
}

// Recover sender using the same logic as preprocess.go (version-based signer selection)

// Run validation checks (fee/nonce/balance - stateless checks done earlier)

// Prepare context for EVM transaction (set infinite gas meter like original flow)

// Validation failed - bump nonce via keeper if it was valid (matches V2's DeliverTxCallback
// behavior where nonce is incremented even on fee validation failures).
// For successful txs, the nonce is bumped by the EVM during execution.

// V2 reports intrinsic gas as gasUsed even on validation failure (for metrics),
// but no actual balance is deducted

//nolint:gosec
//nolint:gosec

// Unassociated addresses require balance migration (iterating all balances),
// which giga's cachekv doesn't support. Fall back to v2 for this tx.

// Create state DB for this transaction (only for valid transactions)

// Pre-charge gas fee (like V2's ante handler), then execute with feeAlreadyCharged=true.
// V2 charges fees in the ante handler, then runs the EVM with feeAlreadyCharged=true
// which skips buyGas/refundGas/coinbase. Without this, GasUsed differs between Giga
// and V2, causing LastResultsHash → AppHash divergence.

// Get gas pool (mutated per tx, cannot be cached)

// Use cached block-level constants

// Create Giga executor VM

// Execute with feeAlreadyCharged=true — matching V2's msg_server behavior

// Match V2 error handling: bump nonce, commit fee deduction, track surplus

// stateDB.Finalize is not expected to fail in practice. If it
// does, the nonce bump above may not have been persisted, so per
// the receipt-iff-nonce-bumped invariant we cannot claim the tx
// happened: skip the receipt + deferred-info writes and return.

//nolint:gosec

// Receipt-iff-nonce-bumped invariant: this tx bumped the sender's
// nonce on the line above, so it must produce a receipt. State-
// transition errors land here when Execute() bails before any
// opcode ran (notably EIP-7623's floor-data-gas check, which
// happens inside go-ethereum's Execute() rather than the Sei
// antehandler). Without an explicit WriteReceipt the receipt
// store stays empty for this tx hash — Giga's
// AppendToEvmTxDeferredInfo call below doesn't propagate the
// error, so EndBlock's synthetic-receipt path skips it — and
// eth_getTransactionReceipt returns null forever, hanging any
// client that polls for it.

// EIP-1559 effective gas price (not GasFeeCap)

//nolint:gosec

// Check if the execution hit a fail-fast precompile (Cosmos interop detected)
// Return the error to the caller so it can handle accordingly (e.g., fallback to standard execution)

// Finalize state changes — captures surplus (fee deduction + execution balance changes)

// Write receipt

// Create core.Message from ethTx for WriteReceipt.
// GasPrice must be the EIP-1559 effective gas price (min(baseFee+tip,
// maxFee)) — that's what the chain actually charges (see line 1866)
// and what the receipt's EffectiveGasPrice field needs to report.
// ethTx.GasPrice() returns GasFeeCap for dynamic-fee txs, which puts
// the wrong value on the receipt and breaks EIP-1559 clients.

// Append deferred info for EndBlock processing

// Determine result code based on VM error

// GasWanted should be set to the transaction's gas limit to match standard executor behavior.
// This is critical for LastResultsHash computation which uses Code, Data, GasWanted, and GasUsed.
//nolint:gosec // G115: safe, Gas() won't exceed int64 max
//nolint:gosec // G115: safe, UsedGas won't exceed int64 max

// Build Data field to match standard executor format.
// Standard path wraps MsgEVMTransactionResponse in TxMsgData.
// This is critical for LastResultsHash to match.

// Wrap in TxMsgData like the standard path does

// gigaDeliverTx is the OCC-compatible deliverTx function for the giga executor.
// makeGigaDeliverTx returns an OCC-compatible deliverTx callback that captures the given
// block cache, avoiding mutable state on App for cache lifecycle management.
func (app *App) makeGigaDeliverTx(cache *gigaBlockCache) func(sdk.Context, abci.RequestDeliverTxV2, sdk.Tx, [32]byte) abci.ResponseDeliverTx {
	_ = "STUB: not implemented"
	return nil
}

// Handle panics as V2 does (matches baseapp/recovery.go middleware chain)

// For other panics (e.g., nil deref from malformed protobuf), log and return ErrPanic

// Validate Cosmos SDK envelope (memo, timeoutHeight, signerInfos, etc.)
// This prevents consensus divergence if a malicious proposer includes invalid envelope fields.

// Check if this is a fail-fast error (Cosmos precompile interop detected)

// Return a sentinel response so the caller can fall back to v2.

func (app *App) GetEVMMsg(tx sdk.Tx) (res *evmtypes.MsgEVMTransaction) {
	_ = "STUB: not implemented"
	return nil
}

func (app *App) DecodeTxBytesConcurrently(ctx context.Context, txs [][]byte) ([]sdk.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// not needed on Go 1.22+

// Bail early if another goroutine already failed.

// finalizeDecodedEVMSlot runs EVM Preprocess for a decoded tx at idx. typedTx must be non-nil EVM.
// Panics and preprocess errors clear typedTxs[idx].
func (app *App) finalizeDecodedEVMSlot(ctx sdk.Context, idx int, typedTx sdk.Tx, typedTxs []sdk.Tx) {
	_ = "STUB: not implemented"
	return
}

// FinalizeDecodedTransactionsConcurrently runs EVM preprocessing on typedTxs in place.
// Non-EVM entries are unchanged; failed preprocessing clears the slot (nil).
func (app *App) FinalizeDecodedTransactionsConcurrently(ctx sdk.Context, typedTxs []sdk.Tx) {
	_ = "STUB: not implemented"
	return
}

// DecodeTransactionsConcurrently decodes each tx and runs EVM preprocessing in one goroutine per index.
// Failed decodes, decode panics, and failed preprocessing clear the slot (nil), matching prior behavior.
func (app *App) DecodeTransactionsConcurrently(ctx sdk.Context, txs [][]byte) []sdk.Tx {
	_ = "STUB: not implemented"
	return nil
}

func (app *App) getFinalizeBlockResponse(appHash []byte, events []abci.Event, txResults []*abci.ExecTxResult, endBlockResp abci.ResponseEndBlock) abci.ResponseFinalizeBlock {
	_ = "STUB: not implemented"
	return *new(abci.ResponseFinalizeBlock)
}

// LoadHeight loads a particular height
func (app *App) LoadHeight(height int64) error { _ = "STUB: not implemented"; return nil }

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *App) ModuleAccountAddrs() map[string]bool { _ = "STUB: not implemented"; return nil }

// LegacyAmino returns SimApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) LegacyAmino() *codec.LegacyAmino { _ = "STUB: not implemented"; return nil }

func (app *App) GetValidators() []abci.ValidatorUpdate {
	_ = "STUB: not implemented"
	// AUTOBAHN: After InitChain but before the first Commit, the committed
	// store is empty — staking params don't exist, so reading from committed
	// store panics in MaxValidators. Use DeliverContext when available at
	// height 0, since it has the uncommitted staking state from InitChain.
	// CometBFT consensus never hits this because its handshaker commits
	// after InitChain before any block processing begins.
	return nil
}

// AppCodec returns an app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) AppCodec() codec.Codec {
	_ = "STUB: not implemented"
	return *

	// InterfaceRegistry returns an InterfaceRegistry
	new(codec.Codec)
}

func (app *App) InterfaceRegistry() types.InterfaceRegistry {
	_ = "STUB: not implemented"
	return *new(types.InterfaceRegistry)
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetKey(storeKey string) *sdk.KVStoreKey { _ = "STUB: not implemented"; return nil }

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetTKey(storeKey string) *sdk.TransientStoreKey {
	_ = "STUB: not implemented"
	return nil

	// GetMemKey returns the MemStoreKey for the provided mem key.
	//
	// NOTE: This is solely used for testing purposes.
}

func (app *App) GetMemKey(storeKey string) *sdk.MemoryStoreKey {
	_ = "STUB: not implemented"
	return nil
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetSubspace(moduleName string) paramstypes.Subspace {
	_ = "STUB: not implemented"
	return *new(paramstypes.Subspace)
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *App) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	_ = "STUB: not implemented"
	return
}

// Register legacy tx routes.

// Register new tx routes from grpc-gateway.

// Register new tendermint queries routes from grpc-gateway.

// Register legacy and grpc-gateway routes for all modules.

// register swagger API from root so that other applications can override easily

func (app *App) RPCContextProvider(i int64) sdk.Context {
	_ = "STUB: not implemented"
	return *new(sdk.Context)
}

// Populate ConsensusParams on the RPC ctx. Neither GetCheckCtx nor
// CreateQueryContext sets it (only the tx-execution path does, via
// getContextForTx), so without this, ctx.ConsensusParams() returns
// nil from any RPC handler. evmrpc/block.go's gasLimit and
// evmrpc/info.go's gas-used-ratio rely on it.

// SnapshotAwareRPCContextProvider builds SDK contexts from in-memory memiavl
// snapshots; falls back to RPCContextProvider on miss or unsupported backend.
func (app *App) SnapshotAwareRPCContextProvider() evmrpc.TraceContextProvider {
	_ = "STUB: not implemented"
	return *new(evmrpc.TraceContextProvider)
}

// RegisterTendermintService implements the Application.RegisterLocalServices method.
func (app *App) RegisterLocalServices(node client.LocalClient, txConfig client.TxConfig) {
	_ = "STUB: not implemented"
	return
}

// use current for post v6.0.6 heights

// RegisterSwaggerAPI registers swagger route with API Server
func RegisterSwaggerAPI(rtr *mux.Router) { _ = "STUB: not implemented"; return }

// checkTotalBlockGas checks that the block gas limit is not exceeded by our best estimate of
// the total gas by the txs in the block. The gas of a tx is either the gas estimate if it's an EVM tx,
// or the gas wanted if it's a Cosmos tx. typedTxs must align with proposal order (nil = decode failure).
func (app *App) checkTotalBlockGas(ctx sdk.Context, typedTxs []sdk.Tx) (_result bool) {
	_ = "STUB: not implemented"
	return false
}

// MsgEVMTransaction cannot be gasless under IsTxGasless (only oracle vote / MsgAssociate).
// Skip keeper-backed IsTxGasless for valid single-message EVM txs; still run it when the tx
// is not EVM or EVM classification failed (e.g. multi-msg with an EVM message).

// Unexpected panic: reject the entire proposal.

// Business-logic errors (e.g. duplicate votes): keep going, tx is treated as non-gasless.

// EVM classification failed (e.g. multi-msg containing an EVM message); such a tx won't be
// processed and so contributes no gas to the block.

// Non-fee tx won't be processed and thus won't consume gas. Skipping.

// Overflow guards: gasWanted must fit in int64, and adding it to either accumulator
// must not wrap uint64.
//nolint:gosec

// Prefer the gas estimate when it's a valid EVM estimate (>= MinGasEVMTx) and not
// inflated above gasWanted; otherwise charge full gasWanted.

//nolint:gosec
//nolint:gosec

//nolint:gosec

// isExpectedGaslessMetricsError reports whether err is the well-known oracle
// duplicate-vote error that we deliberately tolerate when collecting
// gasless-tx metrics. errors.Is handles properly-wrapped chains; the substring
// fallback covers chains that lost sentinel identity via %s/%v wrapping.
func isExpectedGaslessMetricsError(err error) bool { _ = "STUB: not implemented"; return false }

// couldBeGaslessTransaction is a fast heuristic that returns true when tx
// might be gasless and a full IsTxGasless keeper check is therefore worth
// running. It MUST be a conservative over-approximation: returning false for
// a tx that is actually gasless would cause its gas to be counted against
// the block limit, producing incorrect gas accounting (and in the worst case
// rejecting an otherwise-valid block).
func (app *App) couldBeGaslessTransaction(tx sdk.Tx) bool { _ = "STUB: not implemented"; return false }

func (app *App) GetTxConfig() client.TxConfig {
	_ = "STUB: not implemented"
	return *new(client.TxConfig)
}

func (app *App) GetLegacyTxConfig() client.TxConfig {
	_ = "STUB: not implemented"
	return *new(client.TxConfig)
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string { _ = "STUB: not implemented"; return nil }

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey sdk.StoreKey) paramskeeper.Keeper {
	_ = "STUB: not implemented"
	return *new(paramskeeper.Keeper)
}

// this line is used by starport scaffolding # stargate/app/paramSubspace

// SimulationManager implements the SimulationApp interface
func (app *App) SimulationManager() *module.SimulationManager {
	_ = "STUB: not implemented"
	return nil
}

func (app *App) BlacklistedAccAddrs() map[string]bool { _ = "STUB: not implemented"; return nil }

func (app *App) GetPrecompileKeepers() putils.Keepers {
	_ = "STUB: not implemented"
	return *new(putils.Keepers)
}

// test-only
func (app *App) SetTxDecoder(txDecoder sdk.TxDecoder) { _ = "STUB: not implemented"; return }

func (app *App) inplacetestnetInitializer(pk cryptotypes.PubKey) error {
	_ = "STUB: not implemented"
	return nil
}

// gigaValidationResult holds the result of EVM transaction validation.
type gigaValidationResult struct {
	err          *abci.ExecTxResult // nil if validation passed
	bumpNonce    bool               // true if tx nonce matches expected nonce
	currentNonce uint64             // the expected nonce at time of validation
	baseFee      *big.Int           // the base fee used for validation
}

// validateGigaEVMTx validates an EVM tx for fee, nonce, and stateless checks.
// Note: Cosmos envelope checks (chain ID, intrinsic gas, etc.) are done earlier via EvmStatelessChecks.
//
// This function handles checks from V2's EVMFeeCheckDecorator + go-ethereum's StatelessChecks:
//  1. Fee cap checks
//  2. Nonce validity (including overflow guard)
//  3. Sender EOA check (unless delegated via EIP-7702)
//  4. Gas fee/tip cap bit length checks
//  5. Tip <= fee cap check
//  6. Set-code tx validation
//  7. Balance check
func (app *App) validateGigaEVMTx(
	ctx sdk.Context,
	ethTx *ethtypes.Transaction,
	sender common.Address,
	seiAddr sdk.AccAddress,
	isAssociated bool,
) gigaValidationResult {
	_ = "STUB: not implemented"
	return *new(gigaValidationResult)
}

// Check nonce validity - determines if we bump nonce on fee/balance failures

// Fee cap below base fee

// Fee cap below minimum fee

// ========================================================================
// go-ethereum StatelessChecks (matches V2's EVMFeeCheckDecorator call to st.StatelessChecks())
// ========================================================================

// Nonce checks (too high, too low, overflow guard)

// Nonce overflow guard (currentNonce + 1 would overflow)

// Sender must be EOA unless delegated (EIP-7702)

// GasFeeCap bit length must be <= 256

// GasTipCap bit length must be <= 256

// GasTipCap must be <= GasFeeCap

// Set-code tx (EIP-7702) validation

// Set-code tx must not be contract creation

// Set-code tx auth list must be non-empty

// ========================================================================
// Balance check (matches V2's st.BuyGas())
// ========================================================================

// Insufficient balance for gas + value

// Include cast address balance for unassociated addresses (matches V2 PreprocessDecorator)

// All checks passed

func init() {
	// override max wasm size to 2MB
	wasmtypes.MaxWasmSize = 2 * 1024 * 1024
}

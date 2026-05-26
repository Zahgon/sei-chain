package keeper

import (
	"testing"

	"github.com/sei-protocol/sei-chain/sei-cosmos/baseapp"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth"
	authkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/bank"
	bankkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/capability"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/crisis"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution"
	distrclient "github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution/client"
	distributionkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/evidence"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/gov"
	govkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/params"
	paramsclient "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/slashing"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/staking"
	stakingkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/keeper"
	stakingtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade"
	upgradeclient "github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade/client"
	ibckeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/keeper"
	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto"
	"github.com/sei-protocol/sei-chain/x/mint"
	dbm "github.com/tendermint/tm-db"

	wasmappparams "github.com/sei-protocol/sei-chain/sei-wasmd/app/params"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
)

var ModuleBasics = module.NewBasicManager(
	auth.AppModuleBasic{},
	bank.AppModuleBasic{},
	capability.AppModuleBasic{},
	staking.AppModuleBasic{},
	mint.AppModuleBasic{},
	distribution.AppModuleBasic{},
	gov.NewAppModuleBasic(
		paramsclient.ProposalHandler, distrclient.ProposalHandler, upgradeclient.ProposalHandler,
	),
	params.AppModuleBasic{},
	crisis.AppModuleBasic{},
	slashing.AppModuleBasic{},
	upgrade.AppModuleBasic{},
	evidence.AppModuleBasic{},
)

func MakeTestCodec(t testing.TB) codec.Codec { _ = "STUB: not implemented"; return *new(codec.Codec) }

func MakeEncodingConfig(_ testing.TB) wasmappparams.EncodingConfig {
	_ = "STUB: not implemented"
	return *new(wasmappparams.EncodingConfig)
}

// add wasmd types

var TestingStakeParams = stakingtypes.Params{
	UnbondingTime:                      100,
	MaxValidators:                      10,
	MaxEntries:                         10,
	HistoricalEntries:                  10,
	MinCommissionRate:                  sdk.MustNewDecFromStr("0.05"),
	MaxVotingPowerRatio:                sdk.MustNewDecFromStr("0.5"),
	MaxVotingPowerEnforcementThreshold: sdk.NewInt(10000),
	BondDenom:                          "stake",
}

type TestFaucet struct {
	t                testing.TB
	bankKeeper       bankkeeper.Keeper
	sender           sdk.AccAddress
	balance          sdk.Coins
	minterModuleName string
}

func NewTestFaucet(t testing.TB, ctx sdk.Context, bankKeeper bankkeeper.Keeper, minterModuleName string, initialAmount ...sdk.Coin) *TestFaucet {
	_ = "STUB: not implemented"
	return nil
}

func (f *TestFaucet) Mint(parentCtx sdk.Context, addr sdk.AccAddress, amounts ...sdk.Coin) {
	_ = "STUB: not implemented"
	return
}

// discard all faucet related events

func (f *TestFaucet) Fund(parentCtx sdk.Context, receiver sdk.AccAddress, amounts ...sdk.Coin) {
	_ = "STUB: not implemented"
	return
}

// ensure faucet is always filled

// discard all faucet related events

func (f *TestFaucet) NewFundedAccount(ctx sdk.Context, amounts ...sdk.Coin) sdk.AccAddress {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}

type TestKeepers struct {
	AccountKeeper  authkeeper.AccountKeeper
	StakingKeeper  stakingkeeper.Keeper
	DistKeeper     distributionkeeper.Keeper
	BankKeeper     bankkeeper.Keeper
	GovKeeper      govkeeper.Keeper
	ContractKeeper types.ContractOpsKeeper
	WasmKeeper     *Keeper
	IBCKeeper      *ibckeeper.Keeper
	Router         *baseapp.Router
	EncodingConfig wasmappparams.EncodingConfig
	Faucet         *TestFaucet
	MultiStore     sdk.CommitMultiStore
}

// CreateDefaultTestInput common settings for CreateTestInput
func CreateDefaultTestInput(t testing.TB) (sdk.Context, TestKeepers) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), *new(TestKeepers)
}

// CreateTestInput encoders can be nil to accept the defaults, or set it to override some of the message handlers (like default)
func CreateTestInput(t testing.TB, isCheckTx bool, supportedFeatures string, opts ...Option) (sdk.Context, TestKeepers) {
	_ = "STUB: not implemented"
	// Load default wasm config
	return *new(sdk.Context), *new(TestKeepers)
}

// encoders can be nil to accept the defaults, or set it to override some of the message handlers (like default)
func createTestInput(
	t testing.TB,
	isCheckTx bool,
	supportedFeatures string,
	wasmConfig types.WasmConfig,
	db dbm.DB,
	opts ...Option,
) (sdk.Context, TestKeepers) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), *new(TestKeepers)
}

// module account permissions

// target store

// prototype

// set genesis items required for distribution

// set some funds ot pay out validatores, based on code from:
// https://github.com/cosmos/cosmos-sdk/blob/fea231556aee4d549d7551a6190389c4328194eb/x/distribution/keeper/keeper_test.go#L50-L57

// add wasm handler so we can loop-back (contracts calling contracts)

// minimal module set that we use for message/ query tests

// TestHandler returns a wasm handler for tests (to avoid circular imports)
func TestHandler(k types.ContractOpsKeeper) sdk.Handler {
	_ = "STUB: not implemented"
	return *new(sdk.Handler)
}

func handleStoreCode(ctx sdk.Context, k types.ContractOpsKeeper, msg *types.MsgStoreCode) (*sdk.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func handleInstantiate(ctx sdk.Context, k types.ContractOpsKeeper, msg *types.MsgInstantiateContract) (*sdk.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func handleExecute(ctx sdk.Context, k types.ContractOpsKeeper, msg *types.MsgExecuteContract) (*sdk.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func RandomAccountAddress(_ testing.TB) sdk.AccAddress {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}

func RandomBech32AccountAddress(t testing.TB) string { _ = "STUB: not implemented"; return "" }

type ExampleContract struct {
	InitialAmount sdk.Coins
	Creator       crypto.PrivKey
	CreatorAddr   sdk.AccAddress
	CodeID        uint64
}

func StoreHackatomExampleContract(t testing.TB, ctx sdk.Context, keepers TestKeepers) ExampleContract {
	_ = "STUB: not implemented"
	return *new(ExampleContract)
}

func StoreBurnerExampleContract(t testing.TB, ctx sdk.Context, keepers TestKeepers) ExampleContract {
	_ = "STUB: not implemented"
	return *new(ExampleContract)
}

func StoreIBCReflectContract(t testing.TB, ctx sdk.Context, keepers TestKeepers) ExampleContract {
	_ = "STUB: not implemented"
	return *new(ExampleContract)
}

func StoreReflectContract(t testing.TB, ctx sdk.Context, keepers TestKeepers) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func StoreExampleContract(t testing.TB, ctx sdk.Context, keepers TestKeepers, wasmFile string) ExampleContract {
	_ = "STUB: not implemented"
	return *new(ExampleContract)
}

var wasmIdent = []byte("\x00\x61\x73\x6D")

type ExampleContractInstance struct {
	ExampleContract
	Contract sdk.AccAddress
}

// SeedNewContractInstance sets the mock wasmerEngine in keeper and calls store + instantiate to init the contract's metadata
func SeedNewContractInstance(t testing.TB, ctx sdk.Context, keepers TestKeepers, mock types.WasmerEngine) ExampleContractInstance {
	_ = "STUB: not implemented"
	return *new(ExampleContractInstance)
}

// StoreRandomContract sets the mock wasmerEngine in keeper and calls store
func StoreRandomContract(t testing.TB, ctx sdk.Context, keepers TestKeepers, mock types.WasmerEngine) ExampleContract {
	_ = "STUB: not implemented"
	return *new(ExampleContract)
}

func StoreRandomContractWithAccessConfig(
	t testing.TB, ctx sdk.Context,
	keepers TestKeepers,
	mock types.WasmerEngine,
	cfg *types.AccessConfig,
) ExampleContract {
	_ = "STUB: not implemented"
	return *new(ExampleContract)
}

//nolint:gocritic

type HackatomExampleInstance struct {
	ExampleContract
	Contract        sdk.AccAddress
	Verifier        crypto.PrivKey
	VerifierAddr    sdk.AccAddress
	Beneficiary     crypto.PrivKey
	BeneficiaryAddr sdk.AccAddress
}

// InstantiateHackatomExampleContract load and instantiate the "./testdata/hackatom.wasm" contract
func InstantiateHackatomExampleContract(t testing.TB, ctx sdk.Context, keepers TestKeepers) HackatomExampleInstance {
	_ = "STUB: not implemented"
	return *new(HackatomExampleInstance)
}

type HackatomExampleInitMsg struct {
	Verifier    sdk.AccAddress `json:"verifier"`
	Beneficiary sdk.AccAddress `json:"beneficiary"`
}

func (m HackatomExampleInitMsg) GetBytes(t testing.TB) []byte {
	_ = "STUB: not implemented"
	return nil
}

type IBCReflectExampleInstance struct {
	Contract      sdk.AccAddress
	Admin         sdk.AccAddress
	CodeID        uint64
	ReflectCodeID uint64
}

// InstantiateIBCReflectContract load and instantiate the "./testdata/ibc_reflect.wasm" contract
func InstantiateIBCReflectContract(t testing.TB, ctx sdk.Context, keepers TestKeepers) IBCReflectExampleInstance {
	_ = "STUB: not implemented"
	return *new(IBCReflectExampleInstance)
}

type IBCReflectInitMsg struct {
	ReflectCodeID uint64 `json:"reflect_code_id"`
}

func (m IBCReflectInitMsg) GetBytes(t testing.TB) []byte { _ = "STUB: not implemented"; return nil }

type BurnerExampleInitMsg struct {
	Payout sdk.AccAddress `json:"payout"`
}

func (m BurnerExampleInitMsg) GetBytes(t testing.TB) []byte { _ = "STUB: not implemented"; return nil }

func fundAccounts(t testing.TB, ctx sdk.Context, am authkeeper.AccountKeeper, bank bankkeeper.Keeper, addr sdk.AccAddress, coins sdk.Coins) {
	_ = "STUB: not implemented"
	return
}

var keyCounter uint64

// we need to make this deterministic (same every test run), as encoded address size and thus gas cost,
// depends on the actual bytes (due to ugly CanonicalAddress encoding)
func keyPubAddr() (crypto.PrivKey, crypto.PubKey, sdk.AccAddress) {
	_ = "STUB: not implemented"
	return *new(crypto.PrivKey), *new(crypto.PubKey), *new(sdk.AccAddress)
}

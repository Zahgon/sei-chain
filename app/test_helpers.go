package app

import (
	"testing"
	"time"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	receipt "github.com/sei-protocol/sei-chain/sei-db/ledger_db/receipt"
	"github.com/stretchr/testify/suite"

	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/types"
	dbm "github.com/tendermint/tm-db"

	bam "github.com/sei-protocol/sei-chain/sei-cosmos/baseapp"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	authtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/types"
	banktypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/types"
	stakingtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
)

const TestContract = "TEST"
const TestUser = "sei1jdppe6fnj2q7hjsepty5crxtrryzhuqsjrj95y"

type TestTx struct {
	msgs []sdk.Msg
}

func NewTestTx(msgs []sdk.Msg) TestTx { _ = "STUB: not implemented"; return *new(TestTx) }

func (t TestTx) GetMsgs() []sdk.Msg { _ = "STUB: not implemented"; return nil }

func (t TestTx) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (t TestTx) GetGasEstimate() uint64 { _ = "STUB: not implemented"; return 0 }

type TestAppOpts struct {
	UseSc          bool
	EnableGiga     bool
	EnableGigaOCC  bool
	ReceiptBackend string // e.g. "parquet" to use parquet receipt store; empty = default (pebble)
}

func (t TestAppOpts) Get(s string) interface{} { _ = "STUB: not implemented"; return nil }

// Disable snapshot creation in tests to avoid background goroutines
// that are not relevant to the test logic

// 0 = disabled

// Disable EVM HTTP and WebSocket servers in tests to avoid port conflicts
// when multiple tests run in parallel (all would try to bind to port 8545)

type TestWrapper struct {
	suite.Suite

	App *App
	Ctx sdk.Context

	// hasT tracks whether SetT was called, so we know if we can use Require()
	hasT bool
}

func NewTestWrapper(tb testing.TB, tm time.Time, valPub cryptotypes.PubKey, enableEVMCustomPrecompiles bool, baseAppOptions ...func(*bam.BaseApp)) *TestWrapper {
	_ = "STUB: not implemented"
	return nil
}

func NewTestWrapperWithSc(t *testing.T, tm time.Time, valPub cryptotypes.PubKey, enableEVMCustomPrecompiles bool, baseAppOptions ...func(*bam.BaseApp)) *TestWrapper {
	_ = "STUB: not implemented"
	return nil
}

func NewGigaTestWrapper(t *testing.T, tm time.Time, valPub cryptotypes.PubKey, enableEVMCustomPrecompiles bool, useOcc bool, baseAppOptions ...func(*bam.BaseApp)) *TestWrapper {
	_ = "STUB: not implemented"
	return nil
}

// NewGigaTestWrapperWithRegularStore creates a test wrapper that runs Giga executor
// but uses regular KVStore instead of GigaKVStore. This is for debugging - it isolates
// whether issues are in the Giga executor logic vs the GigaKVStore layer.
//
// How it works:
// - Creates app with UseSc=true but EnableGiga=false (so GigaKVStore is NOT registered)
// - Manually enables Giga executor flags on the app
// - Sets GigaEvmKeeper.UseRegularStore=true so it uses ctx.KVStore instead of ctx.GigaKVStore
func NewGigaTestWrapperWithRegularStore(t *testing.T, tm time.Time, valPub cryptotypes.PubKey, enableEVMCustomPrecompiles bool, useOcc bool, baseAppOptions ...func(*bam.BaseApp)) *TestWrapper {
	_ = "STUB: not implemented"
	// Create wrapper with Sc but WITHOUT EnableGiga - this means GigaKVStore won't be registered
	return nil
}

// Manually enable Giga executor on the app

// Configure GigaEvmKeeper to use regular KVStore instead of GigaKVStore

// Configure GigaBankKeeper to use regular KVStore instead of GigaKVStore

// Initialize evmone VM if not already initialized (best effort)

// Init genesis for GigaEvmKeeper (now uses regular KVStore)

func newTestWrapper(tb testing.TB, tm time.Time, valPub cryptotypes.PubKey, enableEVMCustomPrecompiles bool, UseSc bool, testAppOpts TestAppOpts, baseAppOptions ...func(*bam.BaseApp)) *TestWrapper {
	_ = "STUB: not implemented"
	return nil
}

// requireNoError handles errors for both tests and benchmarks.
// For tests (when SetT was called), it uses Require().NoError().
// For benchmarks (when SetT wasn't called), it panics on error.
func (s *TestWrapper) requireNoError(err error) { _ = "STUB: not implemented"; return }

// Benchmark context - panic on error

func (s *TestWrapper) FundAcc(acc sdk.AccAddress, amounts sdk.Coins) {
	_ = "STUB: not implemented"
	return
}

func (s *TestWrapper) setupValidator(bondStatus stakingtypes.BondStatus, valPub cryptotypes.PubKey) sdk.ValAddress {
	_ = "STUB: not implemented"
	return *new(sdk.ValAddress)
}

// Create validator message and handle it

func (s *TestWrapper) BeginBlock() { _ = "STUB: not implemented"; return }

func (s *TestWrapper) EndBlock() { _ = "STUB: not implemented"; return }

func setupReceiptStore(storeKey sdk.StoreKey) (receipt.ReceiptStore, error) {
	_ = "STUB: not implemented"
	// Create a unique temporary directory per test process to avoid Pebble DB lock conflicts
	return *new(receipt.ReceiptStore), nil
}

// No min retain blocks in test

func SetupWithDefaultHome(isCheckTx bool, enableEVMCustomPrecompiles bool, overrideWasmGasMultiplier bool, baseAppOptions ...func(*bam.BaseApp)) (res *App) {
	_ = "STUB: not implemented"
	return nil
}

func SetupWithAppOptsAndDefaultHome(isCheckTx bool, appOpts TestAppOpts, enableEVMCustomPrecompiles bool, overrideWasmGasMultiplier bool, baseAppOptions ...func(*bam.BaseApp)) (res *App) {
	_ = "STUB: not implemented"
	return nil
}

// TODO: remove once init chain works with SC

func Setup(tb testing.TB, isCheckTx bool, enableEVMCustomPrecompiles bool, overrideWasmGasMultiplier bool, baseAppOptions ...func(*bam.BaseApp)) (res *App) {
	_ = "STUB: not implemented"
	return nil
}

func SetupWithDB(tb testing.TB, db dbm.DB, isCheckTx bool, enableEVMCustomPrecompiles bool, overrideWasmGasMultiplier bool, baseAppOptions ...func(*bam.BaseApp)) (res *App) {
	_ = "STUB: not implemented"
	return nil
}

// SetupWithScReceiptFromOpts is like SetupWithSc but does not inject a receipt store via AppOption.
// The receipt store is created inside New() from testAppOpts (e.g. testAppOpts.ReceiptBackend = "parquet").
// Use this to test the full app path with rs-backend from config.
func SetupWithScReceiptFromOpts(t *testing.T, isCheckTx bool, enableEVMCustomPrecompiles bool, testAppOpts TestAppOpts, baseAppOptions ...func(*bam.BaseApp)) (res *App) {
	_ = "STUB: not implemented"
	return nil
}

// no options: receipt store is created from testAppOpts inside New()

func SetupWithSc(t *testing.T, isCheckTx bool, enableEVMCustomPrecompiles bool, testAppOpts TestAppOpts, baseAppOptions ...func(*bam.BaseApp)) (res *App) {
	_ = "STUB: not implemented"
	return nil
}

// TODO: remove once init chain works with SC

func SetupTestingAppWithLevelDb(t *testing.T, isCheckTx bool, enableEVMCustomPrecompiles bool) (*App, func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DefaultConsensusParams defines the default Tendermint consensus params used in
// SimApp testing.
var DefaultConsensusParams = &tmproto.ConsensusParams{
	Block: &tmproto.BlockParams{
		MaxBytes: 200000,
		MaxGas:   100000000,
	},
	Evidence: &tmproto.EvidenceParams{
		MaxAgeNumBlocks: 302400,
		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
		MaxBytes:        10000,
	},
	Validator: &tmproto.ValidatorParams{
		PubKeyTypes: []string{
			tmtypes.ABCIPubKeyTypeEd25519,
		},
	},
}

func setup(t *testing.T, withGenesis bool, invCheckPeriod uint) (*App, GenesisState) {
	_ = "STUB: not implemented"
	return nil, *new(GenesisState)
}

// SetupWithGenesisValSet initializes a new SimApp with a validator set and genesis accounts
// that also act as delegators. For simplicity, each validator is bonded with a delegation
// of one consensus engine unit (10^6) in the default token of the simapp from first genesis
// account. A Nop logger is set in SimApp.
func SetupWithGenesisValSet(t *testing.T, valSet *tmtypes.ValidatorSet, genAccs []authtypes.GenesisAccount, balances ...banktypes.Balance) *App {
	_ = "STUB: not implemented"
	return nil
}

// set genesis accounts

// set validators and delegations

// add genesis acc tokens and delegated tokens to total supply

// add bonded amount to bonded pool module account

// update total supply

// init chain will set the validator set and initialize the genesis accounts

// commit genesis changes

// SetupWithGenesisAccounts initializes a new SimApp with the provided genesis
// accounts and possible balances.
func SetupWithGenesisAccounts(t *testing.T, genAccs []authtypes.GenesisAccount, balances ...banktypes.Balance) *App {
	_ = "STUB: not implemented"
	return nil
}

type GenerateAccountStrategy func(int) []sdk.AccAddress

// createRandomAccounts is a strategy used by addTestAddrs() in order to generated addresses in random order.
func createRandomAccounts(accNum int) []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// createIncrementalAccounts is a strategy used by addTestAddrs() in order to generated addresses in ascending order.
func createIncrementalAccounts(accNum int) []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// start at 100 so we can make up to 999 test addresses with valid test addresses

// base address string

// adding on final two digits to make addresses unique

// AddTestAddrsFromPubKeys adds the addresses into the SimApp providing only the public keys.
func AddTestAddrsFromPubKeys(app *App, ctx sdk.Context, pubKeys []cryptotypes.PubKey, accAmt sdk.Int) {
	_ = "STUB: not implemented"
	return
}

// AddTestAddrs constructs and returns accNum amount of accounts with an
// initial balance of accAmt in random order
func AddTestAddrs(app *App, ctx sdk.Context, accNum int, accAmt sdk.Int) []sdk.AccAddress {
	_ = "STUB: not implemented"
	return nil
}

// AddTestAddrs constructs and returns accNum amount of accounts with an
// initial balance of accAmt in random order
func AddTestAddrsIncremental(app *App, ctx sdk.Context, accNum int, accAmt sdk.Int) []sdk.AccAddress {
	_ = "STUB: not implemented"
	return nil
}

func addTestAddrs(app *App, ctx sdk.Context, accNum int, accAmt sdk.Int, strategy GenerateAccountStrategy) []sdk.AccAddress {
	_ = "STUB: not implemented"
	return nil
}

func initAccountWithCoins(app *App, ctx sdk.Context, addr sdk.AccAddress, coins sdk.Coins) {
	_ = "STUB: not implemented"
	return
}

// ConvertAddrsToValAddrs converts the provided addresses to ValAddress.
func ConvertAddrsToValAddrs(addrs []sdk.AccAddress) []sdk.ValAddress {
	_ = "STUB: not implemented"
	return nil
}

func TestAddr(addr string, bech string) (sdk.AccAddress, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), nil
}

func GenTx(gen client.TxConfig, msgs []sdk.Msg, feeAmt sdk.Coins, gas uint64, chainID string, accNums, accSeqs []uint64, priv ...cryptotypes.PrivKey) (sdk.Tx, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Tx), nil
}

// create a random length memo

// 1st round: set SignatureV2 with empty signatures, to set correct
// signer infos.

// 2nd round: once all signer infos are set, every signer can sign.

func SignCheckDeliver(
	t *testing.T, txCfg client.TxConfig, app *bam.BaseApp, header tmproto.Header, msgs []sdk.Msg,
	accNums, accSeqs []uint64, expSimPass, expPass bool, priv ...cryptotypes.PrivKey,
) (sdk.GasInfo, *sdk.Result, error) {
	_ = "STUB: not implemented"
	return *new(sdk.GasInfo), nil, nil
}

// Must simulate now as CheckTx doesn't run Msgs anymore

func GenSequenceOfTxs(txGen client.TxConfig, msgs []sdk.Msg, accNums []uint64, initSeqNums []uint64, numToGenerate int, priv ...cryptotypes.PrivKey) ([]sdk.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func incrementAllSequenceNumbers(initSeqNums []uint64) { _ = "STUB: not implemented"; return }

// CheckBalance checks the balance of an account.
func CheckBalance(t *testing.T, app *App, addr sdk.AccAddress, balances sdk.Coins) {
	_ = "STUB: not implemented"
	return
}

// CreateTestPubKeys returns a total of numPubKeys public keys in ascending order.
func CreateTestPubKeys(numPubKeys int) []cryptotypes.PubKey { _ = "STUB: not implemented"; return nil }

// start at 10 to avoid changing 1 to 01, 2 to 02, etc

// base pubkey string
// adding on final two digits to make pubkeys unique

// NewPubKeyFromHex returns a PubKey from a hex string.
func NewPubKeyFromHex(pk string) (res cryptotypes.PubKey) {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PubKey)
}

const DefaultGenTxGas = 10000000

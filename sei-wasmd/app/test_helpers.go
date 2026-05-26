package app

import (
	"testing"
	"time"

	"github.com/sei-protocol/sei-chain/sei-cosmos/testutil"
	ibckeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/keeper"

	bam "github.com/sei-protocol/sei-chain/sei-cosmos/baseapp"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	authtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/types"
	bankkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/keeper"
	banktypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/types"
	capabilitykeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/keeper"
	distrkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution/keeper"
	evidencekeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/evidence/keeper"
	slashingkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/slashing/keeper"
	stakingkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/keeper"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/types"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm"
)

// DefaultConsensusParams defines the default Tendermint consensus params used in
// WasmApp testing.
var DefaultConsensusParams = &tmproto.ConsensusParams{
	Block: &tmproto.BlockParams{
		MaxBytes: 8000000,
		MaxGas:   1234000000,
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

func setup(t testing.TB, withGenesis bool, invCheckPeriod uint, opts ...wasm.Option) (*WasmApp, GenesisState) {
	_ = "STUB: not implemented"
	return nil, *new(GenesisState)
}

// SetupWithGenesisValSet initializes a new WasmApp with a validator set and genesis accounts
// that also act as delegators. For simplicity, each validator is bonded with a delegation
// of one consensus engine unit (10^6) in the default token of the WasmApp from first genesis
// account. A Nop logger is set in WasmApp.
func SetupWithGenesisValSet(t *testing.T, chainID string, valSet *tmtypes.ValidatorSet, genAccs []authtypes.GenesisAccount, opts []wasm.Option, balances ...banktypes.Balance) *WasmApp {
	_ = "STUB: not implemented"
	return nil
}

// set genesis accounts

// set validators and delegations

// add genesis acc tokens and delegated tokens to total supply

// add bonded amount to bonded pool module account

// update total supply

// init chain will set the validator set and initialize the genesis accounts

// This line is necessary due to the capability module which has a map as well as store usages,
// and because the initChain runs 3 times (deliver, prepare, process proposal),
// we need to make sure to first commit the last one so the proper values are persisted from the
// memory map into the store.

// commit genesis changes

// SetupWithEmptyStore setup a wasmd app instance with empty DB
func SetupWithEmptyStore(t testing.TB) *WasmApp { _ = "STUB: not implemented"; return nil }

type GenerateAccountStrategy func(int) []sdk.AccAddress

// createRandomAccounts is a strategy used by addTestAddrs() in order to generated addresses in random order.
func createRandomAccounts(accNum int) []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// createIncrementalAccounts is a strategy used by addTestAddrs() in order to generated addresses in ascending order.
func createIncrementalAccounts(accNum int) []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// start at 100 so we can make up to 999 test addresses with valid test addresses

// base address string

// adding on final two digits to make addresses unique

// AddTestAddrsFromPubKeys adds the addresses into the WasmApp providing only the public keys.
func AddTestAddrsFromPubKeys(app *WasmApp, ctx sdk.Context, pubKeys []cryptotypes.PubKey, accAmt sdk.Int) {
	_ = "STUB: not implemented"
	return
}

// AddTestAddrs constructs and returns accNum amount of accounts with an
// initial balance of accAmt in random order
func AddTestAddrs(app *WasmApp, ctx sdk.Context, accNum int, accAmt sdk.Int) []sdk.AccAddress {
	_ = "STUB: not implemented"
	return nil
}

// AddTestAddrs constructs and returns accNum amount of accounts with an
// initial balance of accAmt in random order
func AddTestAddrsIncremental(app *WasmApp, ctx sdk.Context, accNum int, accAmt sdk.Int) []sdk.AccAddress {
	_ = "STUB: not implemented"
	return nil
}

func addTestAddrs(app *WasmApp, ctx sdk.Context, accNum int, accAmt sdk.Int, strategy GenerateAccountStrategy) []sdk.AccAddress {
	_ = "STUB: not implemented"
	return nil
}

// fill all the addresses with some coins, set the loose pool tokens simultaneously

func initAccountWithCoins(app *WasmApp, ctx sdk.Context, addr sdk.AccAddress, coins sdk.Coins) {
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

// CheckBalance checks the balance of an account.
func CheckBalance(t *testing.T, app *WasmApp, addr sdk.AccAddress, balances sdk.Coins) {
	_ = "STUB: not implemented"
	return
}

const DefaultGas = 1200000

// SignCheckDeliver checks a generated signed transaction and simulates a
// block commitment with the given transaction. A test assertion is made using
// the parameter 'expPass' against the result. A corresponding result is
// returned.
func SignCheckDeliver(
	t *testing.T, txCfg client.TxConfig, app *bam.BaseApp, header tmproto.Header, msgs []sdk.Msg,
	chainID string, accNums, accSeqs []uint64, expSimPass, expPass bool, priv ...cryptotypes.PrivKey,
) (sdk.GasInfo, *sdk.Result, error) {
	_ = "STUB: not implemented"
	return *new(sdk.GasInfo), nil, nil
}

// Must simulate now as CheckTx doesn't run Msgs anymore

// Simulate a sending a transaction and committing a block

// SignAndDeliver signs and delivers a transaction. No simulation occurs as the
// ibc testing package causes checkState and deliverState to diverge in block time.
func SignAndDeliver(
	t *testing.T, txCfg client.TxConfig, app *bam.BaseApp,
	ibcKeeper *ibckeeper.Keeper, stakingKeeper stakingkeeper.Keeper, capabilityKeeper *capabilitykeeper.Keeper, distrKeeper *distrkeeper.Keeper, slashingKeeper *slashingkeeper.Keeper, evidenceKeeper *evidencekeeper.Keeper, header tmproto.Header, msgs []sdk.Msg,
	chainID string, accNums, accSeqs []uint64, expSimPass, expPass bool, priv ...cryptotypes.PrivKey,
) (sdk.GasInfo, *sdk.Result, error) {
	_ = "STUB: not implemented"
	return *new(sdk.GasInfo), nil, nil
}

// Simulate a sending a transaction and committing a block

// GenSequenceOfTxs generates a set of signed transactions of messages, such
// that they differ only by having the sequence numbers incremented between
// every transaction.
func GenSequenceOfTxs(txGen client.TxConfig, msgs []sdk.Msg, accNums []uint64, initSeqNums []uint64, numToGenerate int, priv ...cryptotypes.PrivKey) ([]sdk.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func incrementAllSequenceNumbers(initSeqNums []uint64) { _ = "STUB: not implemented"; return }

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

// EmptyBaseAppOptions is a stub implementing AppOptions
type EmptyBaseAppOptions = testutil.TestAppOpts

// FundAccount is a utility function that funds an account by minting and
// sending the coins to the address. This should be used for testing purposes
// only!
//
// Instead of using the mint module account, which has the
// permission of minting, create a "faucet" account. (@fdymylja)
func FundAccount(bankKeeper bankkeeper.Keeper, ctx sdk.Context, addr sdk.AccAddress, amounts sdk.Coins) error {
	_ = "STUB: not implemented"
	return nil
}

// FundModuleAccount is a utility function that funds a module account by
// minting and sending the coins to the address. This should be used for testing
// purposes only!
//
// Instead of using the mint module account, which has the
// permission of minting, create a "faucet" account. (@fdymylja)
func FundModuleAccount(bankKeeper bankkeeper.Keeper, ctx sdk.Context, recipientMod string, amounts sdk.Coins) error {
	_ = "STUB: not implemented"
	return nil
}

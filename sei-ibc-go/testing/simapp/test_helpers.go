package simapp

import (
	"testing"
	"time"

	bam "github.com/sei-protocol/sei-chain/sei-cosmos/baseapp"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	authtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/types"
	banktypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/types"
	capabilitykeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/keeper"
	distrkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution/keeper"
	evidencekeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/evidence/keeper"
	slashingkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/slashing/keeper"
	stakingkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/keeper"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/types"

	ibckeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/keeper"
)

// DefaultConsensusParams defines the default Tendermint consensus params used in
// SimApp testing.
var DefaultConsensusParams = &tmproto.ConsensusParams{
	Block: &tmproto.BlockParams{
		MaxBytes: 200000,
		MaxGas:   2000000,
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

func setup(withGenesis bool, invCheckPeriod uint) (*SimApp, GenesisState) {
	_ = "STUB: not implemented"
	return nil, *new(GenesisState)
}

// Setup initializes a new SimApp. A Nop logger is set in SimApp.
func Setup(isCheckTx bool) *SimApp { _ = "STUB: not implemented"; return nil }

// init chain must be called to stop deliverState from being nil

// Initialize the chain

// SetupWithGenesisValSet initializes a new SimApp with a validator set and genesis accounts
// that also act as delegators. For simplicity, each validator is bonded with a delegation
// of one consensus engine unit (10^6) in the default token of the simapp from first genesis
// account. A Nop logger is set in SimApp.
func SetupWithGenesisValSet(t *testing.T, valSet *tmtypes.ValidatorSet, genAccs []authtypes.GenesisAccount, chainID string, powerReduction sdk.Int, balances ...banktypes.Balance) *SimApp {
	_ = "STUB: not implemented"
	return nil
}

// set genesis accounts

// set validators and delegations

// add bonded amount to bonded pool module account

// set validators and delegations

// update total supply

// init chain will set the validator set and initialize the genesis accounts

// This line is necessary due to the capability module which has a map as well as store usages,
// and because the initChain runs 3 times (deliver, prepare, process proposal),
// we need to make sure to first commit the last one so the proper values are persisted from the
// memory map into the store.

// commit genesis changes

type GenerateAccountStrategy func(int) []sdk.AccAddress

// createRandomAccounts is a strategy used by addTestAddrs() in order to generated addresses in random order.
func createRandomAccounts(accNum int) []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// createIncrementalAccounts is a strategy used by addTestAddrs() in order to generated addresses in ascending order.
func createIncrementalAccounts(accNum int) []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// start at 100 so we can make up to 999 test addresses with valid test addresses

// base address string

// adding on final two digits to make addresses unique

// AddTestAddrsFromPubKeys adds the addresses into the SimApp providing only the public keys.
func AddTestAddrsFromPubKeys(app *SimApp, ctx sdk.Context, pubKeys []cryptotypes.PubKey, accAmt sdk.Int) {
	_ = "STUB: not implemented"
	return
}

// AddTestAddrs constructs and returns accNum amount of accounts with an
// initial balance of accAmt in random order
func AddTestAddrs(app *SimApp, ctx sdk.Context, accNum int, accAmt sdk.Int) []sdk.AccAddress {
	_ = "STUB: not implemented"
	return nil
}

// AddTestAddrs constructs and returns accNum amount of accounts with an
// initial balance of accAmt in random order
func AddTestAddrsIncremental(app *SimApp, ctx sdk.Context, accNum int, accAmt sdk.Int) []sdk.AccAddress {
	_ = "STUB: not implemented"
	return nil
}

func addTestAddrs(app *SimApp, ctx sdk.Context, accNum int, accAmt sdk.Int, strategy GenerateAccountStrategy) []sdk.AccAddress {
	_ = "STUB: not implemented"
	return nil
}

func initAccountWithCoins(app *SimApp, ctx sdk.Context, addr sdk.AccAddress, coins sdk.Coins) {
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
func CheckBalance(t *testing.T, app *SimApp, addr sdk.AccAddress, balances sdk.Coins) {
	_ = "STUB: not implemented"
	return
}

// SignAndDeliver signs and delivers a transaction. No simulation occurs as the
// ibc testing package causes checkState and deliverState to diverge in block time.
func SignAndDeliver(
	t *testing.T, txCfg client.TxConfig, app *bam.BaseApp, ibcKeeper *ibckeeper.Keeper, stakingKeeper stakingkeeper.Keeper, capabilityKeeper *capabilitykeeper.Keeper, distrKeeper *distrkeeper.Keeper, slashingKeeper *slashingkeeper.Keeper, evidenceKeeper *evidencekeeper.Keeper, header tmproto.Header, msgs []sdk.Msg,
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

// EmptyAppOptions is a stub implementing AppOptions
type EmptyAppOptions struct{}

// Get implements AppOptions
func (ao EmptyAppOptions) Get(o string) interface{} { _ = "STUB: not implemented"; return nil }

// FundAccount is a utility function that funds an account by minting and sending the coins to the address
// TODO(fdymylja): instead of using the mint module account, which has the permission of minting, create a "faucet" account
func FundAccount(app *SimApp, ctx sdk.Context, addr sdk.AccAddress, amounts sdk.Coins) error {
	_ = "STUB: not implemented"
	return nil
}

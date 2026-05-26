package utils

import (
	"crypto/ecdsa"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/store"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	wasmkeeper "github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/keeper"

	"github.com/sei-protocol/sei-chain/app"
)

// ignoreStoreKeys are store keys that are not compared
var ignoredStoreKeys = map[string]struct{}{
	"mem_capability": {},
	"epoch":          {},
	"deferredcache":  {},
}

type TestMessage struct {
	Msg       sdk.Msg
	Type      string
	EVMSigner TestAcct
	IsEVM     bool
}

type TestContext struct {
	Ctx            sdk.Context
	CodeID         uint64
	CW20CodeID     uint64
	Validator      TestAcct
	TestAccounts   []TestAcct
	ContractKeeper *wasmkeeper.PermissionedKeeper
	TestApp        *app.App
	CW20Addrs      []string
}

type TestAcct struct {
	ValidatorAddress sdk.ValAddress
	AccountAddress   sdk.AccAddress
	PrivateKey       cryptotypes.PrivKey
	PublicKey        cryptotypes.PubKey
	EvmAddress       common.Address
	EvmSigner        ethtypes.Signer
	EvmPrivateKey    *ecdsa.PrivateKey
}

func NewTestAccounts(count int) []TestAcct { _ = "STUB: not implemented"; return nil }

func NewSigner() TestAcct { _ = "STUB: not implemented"; return *new(TestAcct) }

func Funds(amount int64) sdk.Coins { _ = "STUB: not implemented"; return *new(sdk.Coins) }

func panicIfErr(err error) { _ = "STUB: not implemented"; return }

func addressToValAddress(addr sdk.AccAddress) sdk.ValAddress {
	_ = "STUB: not implemented"
	return *new(sdk.ValAddress)
}

// deployCW20Token deploys a CW20 token contract for testing
func deployCW20Token(tCtx *TestContext, i int) (string, error) {
	_ = "STUB: not implemented"
	// CW20 instantiate message with initial balances for the test accounts
	return "", nil
}

// Execute the message to instantiate the contract

// NewTestContext initializes a new TestContext with a new app and a new contract
func NewTestContext(tb testing.TB, testAccts []TestAcct, blockTime time.Time, workers int, occEnabled bool) *TestContext {
	_ = "STUB: not implemented"
	return nil
}

// deploy a contract so we can use it

// Upload the CW20 contract

// ToTxBytes converts test messages to transaction bytes.
// This includes signing, encoding, and state preparation (funding accounts, updating sequences).
func ToTxBytes(testCtx *TestContext, msgs []*TestMessage) [][]byte {
	_ = "STUB: not implemented"
	return nil
}

func toTxBytes(testCtx *TestContext, msgs []*TestMessage) [][]byte {
	_ = "STUB: not implemented"
	return nil
}

// fund account so it has funds

// RunWithOCC runs the given messages with OCC enabled, number of workers is configured via context
func RunWithOCC(testCtx *TestContext, msgs []*TestMessage) ([]types.Event, []*types.ExecTxResult, types.ResponseEndBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil, *new(types.ResponseEndBlock), nil
}

// RunWithoutOCC runs the given messages without OCC enabled
func RunWithoutOCC(testCtx *TestContext, msgs []*TestMessage) ([]types.Event, []*types.ExecTxResult, types.ResponseEndBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil, *new(types.ResponseEndBlock), nil
}

func runTxs(testCtx *TestContext, msgs []*TestMessage, occ bool) ([]types.Event, []*types.ExecTxResult, types.ResponseEndBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil, *new(types.ResponseEndBlock), nil
}

// ProcessBlockDirect calls ProcessBlock directly with pre-prepared transaction bytes.
// This is useful for benchmarks where you want to measure only ProcessBlock execution time,
// excluding the overhead of transaction encoding, signing, and state preparation.
func ProcessBlockDirect(testCtx *TestContext, txs [][]byte, occ bool) ([]types.Event, []*types.ExecTxResult, types.ResponseEndBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil, *new(types.ResponseEndBlock), nil
}

func JoinMsgs(msgsList ...[]*TestMessage) []*TestMessage { _ = "STUB: not implemented"; return nil }

func Shuffle(msgs []*TestMessage) []*TestMessage { _ = "STUB: not implemented"; return nil }

func CompareStores(t *testing.T, storeKey sdk.StoreKey, expected store.KVStore, actual store.KVStore, testName string) {
	_ = "STUB: not implemented"
	return
}

// Iterate over the expected store

// Ensure the key exists in the actual store

// Compare the values for the current key

// Move to the next key in the actual store for the upcoming iteration

// Ensure there are no extra keys in the actual store

package kvstore

import (
	"context"
	"sync"

	"github.com/sei-protocol/seilog"
	dbm "github.com/tendermint/tm-db"

	"github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/proxy"
	cryptoproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/crypto"
)

var (
	logger = seilog.NewLogger("tendermint", "abci", "example", "kvstore")

	stateKey        = []byte("stateKey")
	kvPairPrefixKey = []byte("kvPairKey:")

	ProtocolVersion uint64 = 0x1
)

const ValidatorSetChangePrefix = "val:"

type State struct {
	db      dbm.DB
	Size    int64  `json:"size"`
	Height  int64  `json:"height"`
	AppHash []byte `json:"app_hash"`
}

func loadState(db dbm.DB) State { _ = "STUB: not implemented"; return *new(State) }

func saveState(state State) { _ = "STUB: not implemented"; return }

func prefixKey(key []byte) []byte { _ = "STUB: not implemented"; return nil }

//---------------------------------------------------

var _ types.Application = (*Application)(nil)

type Application struct {
	types.BaseApplication
	mu           sync.Mutex
	state        State
	RetainBlocks int64 // blocks to retain after commit (via ResponseCommit.RetainHeight)

	// validator set
	ValUpdates         []types.ValidatorUpdate
	valAddrToPubKeyMap map[string]cryptoproto.PublicKey
}

func NewApplication() *Application { _ = "STUB: not implemented"; return nil }

func NewProxy() *proxy.Proxy { _ = "STUB: not implemented"; return nil }

func (app *Application) InitChain(_ context.Context, req *types.RequestInitChain) (*types.ResponseInitChain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *Application) Info(_ context.Context, req *types.RequestInfo) (*types.ResponseInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *Application) GetValidators() []types.ValidatorUpdate {
	_ = "STUB: not implemented"
	return nil

	// tx is either "val:pubkey!power" or "key=value" or just arbitrary bytes
}

func (app *Application) handleTx(tx []byte) *types.ExecTxResult {
	_ = "STUB: not implemented"
	// if it starts with "val:", update the validator set
	// format is "val:pubkey!power"
	return nil
}

// update validators in the merkle tree
// and in app.ValUpdates

func (app *Application) Close() error { _ = "STUB: not implemented"; return nil }

func (app *Application) FinalizeBlock(_ context.Context, req *types.RequestFinalizeBlock) (*types.ResponseFinalizeBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// reset valset changes

// Punish validators who committed equivocation.

// Using a memdb - just return the big endian size of the db

func (*Application) CheckTx(_ context.Context, req *types.RequestCheckTxV2) *types.ResponseCheckTxV2 {
	_ = "STUB: not implemented"
	return nil
}

func (app *Application) Commit(_ context.Context) (*types.ResponseCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns an associated value or nil if missing.
func (app *Application) Query(_ context.Context, reqQuery *types.RequestQuery) (*types.ResponseQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*Application) ProcessProposal(_ context.Context, req *types.RequestProcessProposal) (*types.ResponseProcessProposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//---------------------------------------------
// update validators

func (app *Application) SetValidators(validators []types.ValidatorUpdate) {
	_ = "STUB: not implemented"
	return
}

func (app *Application) Validators() (validators []types.ValidatorUpdate) {
	_ = "STUB: not implemented"
	return nil
}

func MakeValSetChangeTx(pubkey cryptoproto.PublicKey, power int64) []byte {
	_ = "STUB: not implemented"
	return nil
}

func isValidatorTx(tx []byte) bool { _ = "STUB: not implemented"; return false }

// format is "val:pubkey!power"
// pubkey is a base64-encoded 32-byte ed25519 key
func (app *Application) execValidatorTx(tx []byte) *types.ExecTxResult {
	_ = "STUB: not implemented"
	return nil
}

//  get the pubkey and power

// decode the pubkey

// decode the power

// update

// add, update, or remove a validator
func (app *Application) updateValidator(v types.ValidatorUpdate) *types.ExecTxResult {
	_ = "STUB: not implemented"
	return nil
}

// remove validator

// add or update validator

// we only update the changes array if we successfully updated the tree

// -----------------------------
// prepare proposal machinery

const PreparePrefix = "prepare"

func isPrepareTx(tx []byte) bool { _ = "STUB: not implemented"; return false }

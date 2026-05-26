package mock

import (
	"context"
	"encoding/json"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	storetypes "github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// NewApp creates a simple mock kvstore app for testing. It should work
// similar to a real app. Make sure rootDir is empty before running the test,
// in order to guarantee consistent results
func NewApp(rootDir string) (abci.Application, error) {
	_ = "STUB: not implemented"
	return *new(abci.Application), nil
}

// Capabilities key to access the main KVStore.

// Create BaseApp.

// Set mounts for BaseApp's MultiStore.

// Load latest version.

// KVStoreHandler is a simple handler that takes kvstoreTx and writes
// them to the db
func KVStoreHandler(storeKey sdk.StoreKey) sdk.Handler {
	_ = "STUB: not implemented"
	return *new(sdk.Handler)
}

// tx is already unmarshalled

// basic KV structure
type KV struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// What Genesis JSON is formatted as
type GenesisJSON struct {
	Values []KV `json:"values"`
}

// InitChainer returns a function that can initialize the chain
// with key/value pairs
func InitChainer(key sdk.StoreKey) func(sdk.Context, abci.RequestInitChain) abci.ResponseInitChain {
	_ = "STUB: not implemented"
	return nil
}

// TODO https://github.com/cosmos/cosmos-sdk/issues/468
// return sdk.ErrGenesisParse("").TraceCause(err, "")

// AppGenState can be passed into InitCmd, returns a static string of a few
// key-values that can be parsed by InitChainer
func AppGenState(_ *codec.LegacyAmino, _ types.GenesisDoc, _ []json.RawMessage) (appState json.
	RawMessage, err error) {
	_ = "STUB: not implemented"
	return *new(json.
		RawMessage), nil
}

// AppGenStateEmpty returns an empty transaction state for mocking.
func AppGenStateEmpty(_ *codec.LegacyAmino, _ types.GenesisDoc, _ []json.RawMessage) (
	appState json.RawMessage, err error) {
	_ = "STUB: not implemented"
	return *new(json.RawMessage), nil
}

// Manually write the handlers for this custom message
type MsgServer interface {
	Test(ctx context.Context, msg *kvstoreTx) (*sdk.Result, error)
}

type MsgServerImpl struct {
	capKeyMainStore *storetypes.KVStoreKey
}

func (m MsgServerImpl) Test(ctx context.Context, msg *kvstoreTx) (*sdk.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

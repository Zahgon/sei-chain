package tests

import (
	"sync/atomic"
	"testing"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/sei-protocol/sei-chain/app"
	"github.com/sei-protocol/sei-chain/evmrpc"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
)

const testAddr = "127.0.0.1"

var portProvider = atomic.Int32{}

func init() {
	portProvider.Store(7800)
}

type TestServer struct {
	evmrpc.EVMServer
	port int

	mockClient *MockClient
	app        *app.App
}

func (ts TestServer) Run(r func(port int)) { _ = "STUB: not implemented"; return }

func (ts TestServer) SetupBlocks(blocks [][][]byte, initializer ...func(sdk.Context, *app.App)) {
	_ = "STUB: not implemented"
	return
}

func initializeApp(
	t *testing.T,
	chainID string,
	initializer ...func(sdk.Context, *app.App),
) (*app.App, *abci.ResponseFinalizeBlock) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SetupTestServer(
	t *testing.T,
	blocks [][][]byte,
	initializer ...func(sdk.Context, *app.App),
) TestServer {
	_ = "STUB: not implemented"
	return *new(TestServer)
}

// for i, txRes := range res.TxResults {
// 	fmt.Printf("tx %d: %s\n", i, txRes.Log)
// }

func SetupMockPacificTestServer(t *testing.T, initializer func(*app.App, *MockClient) sdk.Context) TestServer {
	_ = "STUB: not implemented"
	return *new(TestServer)
}

// seed mock client with genesis block results so latest height queries work

func setupTestServer(
	a *app.App,
	ctxProvider func(int64) sdk.Context,
	mockClient *MockClient,
) TestServer {
	_ = "STUB: not implemented"
	return *new(TestServer)
}

func sendRequestWithNamespace(namespace string, port int, method string, params ...interface{}) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func formatParam(p interface{}) string { _ = "STUB: not implemented"; return "" }

func signAndEncodeTx(txData ethtypes.TxData, mnemonic string) []byte {
	_ = "STUB: not implemented"
	return nil
}

func encodeEvmTx(txData ethtypes.TxData, signed *ethtypes.Transaction) []byte {
	_ = "STUB: not implemented"
	return nil
}

func signAndEncodeCosmosTx(msg sdk.Msg, mnemonic string, acctN uint64, seq uint64) []byte {
	_ = "STUB: not implemented"
	return nil
}

// TODO: pass in testing.T and assert no error instead

func encodeCosmosTx(tx sdk.Tx) []byte { _ = "STUB: not implemented"; return nil }

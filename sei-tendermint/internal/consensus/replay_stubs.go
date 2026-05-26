package consensus

import (
	"context"

	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/proxy"
)

//-----------------------------------------------------------------------------
// mockProxyApp uses Responses to FinalizeBlock to give the right results.
//
// Useful because we don't want to call Commit() twice for the same block on
// the real app.

func newMockProxyApp(
	appHash []byte,
	finalizeBlockResponses *abci.ResponseFinalizeBlock,
) *proxy.Proxy {
	_ = "STUB: not implemented"
	return nil
}

type mockProxyApp struct {
	abci.BaseApplication

	appHash                []byte
	txCount                int
	finalizeBlockResponses *abci.ResponseFinalizeBlock
}

func (mock *mockProxyApp) FinalizeBlock(_ context.Context, req *abci.RequestFinalizeBlock) (*abci.ResponseFinalizeBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mock *mockProxyApp) Commit(context.Context) (*abci.ResponseCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

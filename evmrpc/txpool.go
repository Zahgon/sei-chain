package evmrpc

import (
	"context"

	"github.com/ethereum/go-ethereum/export"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/evm/keeper"
)

type TxPoolAPI struct {
	tmClient         client.LocalClient
	keeper           *keeper.Keeper
	ctxProvider      func(int64) sdk.Context
	txConfigProvider func(int64) client.TxConfig
	txPoolConfig     *TxPoolConfig
	connectionType   ConnectionType
}

type TxPoolConfig struct {
	maxNumTxs int
}

// NewTxPoolConfig creates a new TxPoolConfig primarily for tests.
func NewTxPoolConfig(maxNumTxs int) *TxPoolConfig { _ = "STUB: not implemented"; return nil }

func NewTxPoolAPI(tmClient client.LocalClient, k *keeper.Keeper, ctxProvider func(int64) sdk.Context, txConfigProvider func(int64) client.TxConfig, txPoolConfig *TxPoolConfig, connectionType ConnectionType) *TxPoolAPI {
	_ = "STUB: not implemented"
	return nil
}

// Content returns the content of the txpool.
// for now, we put all unconfirmed txs in pending and none in queued.
func (t *TxPoolAPI) Content(ctx context.Context) (result map[string]map[string]map[string]*export.RPCTransaction, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// not an evm tx

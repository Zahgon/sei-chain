package evmrpc

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/evm/keeper"
)

type NetAPI struct {
	tmClient       client.LocalClient
	keeper         *keeper.Keeper
	ctxProvider    func(int64) sdk.Context
	connectionType ConnectionType
}

func NewNetAPI(tmClient client.LocalClient, k *keeper.Keeper, ctxProvider func(int64) sdk.Context, connectionType ConnectionType) *NetAPI {
	_ = "STUB: not implemented"
	return nil
}

func (i *NetAPI) Version(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

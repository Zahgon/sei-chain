package evmrpc

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/evm/keeper"
)

type AssociationAPI struct {
	tmClient         client.LocalClient
	keeper           *keeper.Keeper
	ctxProvider      func(int64) sdk.Context
	txConfigProvider func(int64) client.TxConfig
	sendAPI          *SendAPI
	connectionType   ConnectionType
	watermarks       *WatermarkManager
}

func NewAssociationAPI(
	tmClient client.LocalClient,
	k *keeper.Keeper,
	ctxProvider func(int64) sdk.Context,
	txConfigProvider func(int64) client.TxConfig,
	sendAPI *SendAPI,
	connectionType ConnectionType,
	watermarks *WatermarkManager,
) *AssociationAPI {
	_ = "STUB: not implemented"
	return nil
}

type AssociateRequest struct {
	R             string `json:"r"`
	S             string `json:"s"`
	V             string `json:"v"`
	CustomMessage string `json:"custom_message"`
}

func (t *AssociationAPI) Associate(ctx context.Context, req *AssociateRequest) (returnErr error) {
	_ = "STUB: not implemented"
	return nil
}

func (t *AssociationAPI) GetSeiAddress(ctx context.Context, ethAddress common.Address) (result string, returnErr error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (t *AssociationAPI) GetEVMAddress(ctx context.Context, seiAddress string) (result string, returnErr error) {
	_ = "STUB: not implemented"
	return "", nil
}

func decodeHexString(hexString string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *AssociationAPI) GetCosmosTx(ctx context.Context, ethHash common.Hash) (result string, returnErr error) {
	_ = "STUB: not implemented"
	return "", nil
}

//nolint:gosec

func (t *AssociationAPI) GetEvmTx(ctx context.Context, cosmosHash string) (result string, returnErr error) {
	_ = "STUB: not implemented"
	return "", nil
}

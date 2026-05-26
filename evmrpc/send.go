package evmrpc

import (
	"context"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/export"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/sei-protocol/sei-chain/app/legacyabci"
	"github.com/sei-protocol/sei-chain/sei-cosmos/baseapp"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/evm/keeper"
)

type SendAPI struct {
	tmClient         client.LocalClient
	txConfigProvider func(int64) client.TxConfig
	sendConfig       *SendConfig
	keeper           *keeper.Keeper
	ctxProvider      func(int64) sdk.Context
	homeDir          string
	backend          *Backend
	connectionType   ConnectionType
}

type SendConfig struct {
	slow bool
}

func NewSendAPI(
	tmClient client.LocalClient,
	txConfigProvider func(int64) client.TxConfig,
	sendConfig *SendConfig,
	k *keeper.Keeper,
	beginBlockKeepers legacyabci.BeginBlockKeepers,
	ctxProvider func(int64) sdk.Context,
	homeDir string,
	simulateConfig *SimulateConfig,
	app *baseapp.BaseApp,
	antehandler sdk.AnteHandler,
	connectionType ConnectionType,
	globalBlockCache BlockCache,
	cacheCreationMutex *sync.Mutex,
	watermarks *WatermarkManager,
) *SendAPI {
	_ = "STUB: not implemented"
	return nil
}

func (s *SendAPI) SendRawTransaction(ctx context.Context, input hexutil.Bytes) (hash common.Hash, err error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// getSender fails for AccessListTx, in which case we are not able to proxy or simulate,
// but we still need to handle it.

// HTTP transport pooling already happens globally underneath net/http, so
// creating a fresh RPC client per proxied request is fine here. If we
// start proxying over WebSocket, we'll need explicit custom pooling since
// the underlying TCP connection lifecycle is strictly bound to Dial -> Close calls.

// if issue simulating, fallback to gas limit
// simulation requires sender.

func getSender(tx *ethtypes.Transaction, chainID *big.Int) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

func (s *SendAPI) simulateTx(ctx context.Context, sender common.Address, tx *ethtypes.Transaction) (estimate uint64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *SendAPI) SignTransaction(ctx context.Context, args apitypes.SendTxArgs, _ *string) (result *export.SignTransactionResult, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SendAPI) SendTransaction(ctx context.Context, args export.TransactionArgs) (result common.Hash, returnErr error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

func (s *SendAPI) signTransaction(unsignedTx *ethtypes.Transaction, from string) (*ethtypes.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

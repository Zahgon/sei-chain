package tx

import (
	"context"

	ctypes "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// QueryTxsByEvents performs a search for transactions for a given set of events
// via the Tendermint RPC. An event takes the form of:
// "{eventAttribute}.{attributeKey} = '{attributeValue}'". Each event is
// concatenated with an 'AND' operand. It returns a slice of Info object
// containing txs and metadata. An error is returned if the query fails.
// If an empty string is provided it will order txs by asc
func QueryTxsByEvents(ctx context.Context, node client.Client, txConfig client.TxConfig, events []string, page, limit int, orderBy string) (*sdk.SearchTxsResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// XXX: implement ANY

// TODO: this may not always need to be proven
// https://github.com/cosmos/cosmos-sdk/issues/6807

//nolint:gosec // TotalCount from Tendermint is non-negative
//nolint:gosec // len() is always non-negative
//nolint:gosec // validated positive above
//nolint:gosec // validated positive above

// QueryTx queries for a single transaction by a hash string in hex format. An
// error is returned if the transaction does not exist or cannot be queried.
func QueryTx(ctx context.Context, node client.Client, txConfig client.TxConfig, hashHexStr string) (*sdk.TxResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//TODO: this may not always need to be proven
// https://github.com/cosmos/cosmos-sdk/issues/6807

// formatTxResults parses the indexed txs into a slice of TxResponse objects.
func formatTxResults(txConfig client.TxConfig, resTxs []*ctypes.ResultTx, resBlocks map[int64]*ctypes.ResultBlock) ([]*sdk.TxResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getBlocksForTxResults(ctx context.Context, node client.Client, resTxs []*ctypes.ResultTx) (map[int64]*ctypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mkTxResult(txConfig client.TxConfig, resTx *ctypes.ResultTx, resBlock *ctypes.ResultBlock) (*sdk.TxResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deprecated: this interface is used only internally for scenario we are
// deprecating (StdTxConfig support)
type intoAny interface {
	AsAny() *codectypes.Any
}

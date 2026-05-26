package client

import (
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	tmbytes "github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// GetNode returns an RPC client. If the context's client is not defined, an
// error is returned.
func (ctx Context) GetNode() (Client, error) { _ = "STUB: not implemented"; return *new(Client), nil }

// Query performs a query to a Tendermint node with the provided path.
// It returns the result and height of the query upon success or an error if
// the query fails.
func (ctx Context) Query(path string) ([]byte, int64, error) {
	_ = "STUB: not implemented"
	return nil,

		// QueryWithData performs a query to a Tendermint node with the provided path
		// and a data payload. It returns the result and height of the query upon success
		// or an error if the query fails.
		0, nil
}

func (ctx Context) QueryWithData(path string, data []byte) ([]byte, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0,

		// QueryStore performs a query to a Tendermint node with the provided key and
		// store name. It returns the result and height of the query upon success
		// or an error if the query fails.
		nil
}

func (ctx Context) QueryStore(key tmbytes.HexBytes, storeName string) ([]byte, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// QueryABCI performs a query to a Tendermint node with the provide RequestQuery.
// It returns the ResultQuery obtained from the query. The height used to perform
// the query is the RequestQuery Height if it is non-zero, otherwise the context
// height is used.
func (ctx Context) QueryABCI(req abci.RequestQuery) (abci.ResponseQuery, error) {
	_ = "STUB: not implemented"
	return *new(abci.ResponseQuery), nil
}

// GetFromAddress returns the from address from the context's name.
func (ctx Context) GetFromAddress() sdk.AccAddress {
	_ = "STUB: not implemented"
	return *

	// GetFeeGranterAddress returns the fee granter address from the context
	new(sdk.AccAddress)
}

func (ctx Context) GetFeeGranterAddress() sdk.AccAddress {
	_ = "STUB: not implemented"
	return *

	// GetFromName returns the key name for the current context.
	new(sdk.AccAddress)
}

func (ctx Context) GetFromName() string { _ = "STUB: not implemented"; return "" }

func (ctx Context) queryABCI(req abci.RequestQuery) (abci.ResponseQuery, error) {
	_ = "STUB: not implemented"
	return *new(abci.ResponseQuery), nil
}

// fallback on the context height

// data from trusted node or subspace query doesn't need verification

func sdkErrorToGRPCError(resp abci.ResponseQuery) error { _ = "STUB: not implemented"; return nil }

// query performs a query to a Tendermint node with the provided store name
// and path. It returns the result and height of the query upon success
// or an error if the query fails.
func (ctx Context) query(path string, key tmbytes.HexBytes) ([]byte, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// queryStore performs a query to a Tendermint node with the provided a store
// name and path. It returns the result and height of the query upon success
// or an error if the query fails.
func (ctx Context) queryStore(key tmbytes.HexBytes, storeName, endPath string) ([]byte, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// isQueryStoreWithProof expects a format like /<queryType>/<storeName>/<subpath>
// queryType must be "store" and subpath must be "key" to require a proof.
func isQueryStoreWithProof(path string) bool { _ = "STUB: not implemented"; return false }

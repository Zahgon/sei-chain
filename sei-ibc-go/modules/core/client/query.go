package client

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"

	clienttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
)

// QueryTendermintProof performs an ABCI query with the given key and returns
// the value of the query, the proto encoded merkle proof, and the height of
// the Tendermint block containing the state root. The desired tendermint height
// to perform the query should be set in the client context. The query will be
// performed at one below this height (at the IAVL version) in order to obtain
// the correct merkle proof. Proof queries at height less than or equal to 2 are
// not supported. Queries with a client context height of 0 will perform a query
// at the lastest state available.
// Issue: https://github.com/cosmos/cosmos-sdk/issues/6567
func QueryTendermintProof(clientCtx client.Context, key []byte) ([]byte, []byte, clienttypes.Height, error) {
	_ = "STUB: not implemented"
	return nil,

		// ABCI queries at heights 1, 2 or less than or equal to 0 are not supported.
		// Base app does not support queries for height less than or equal to 1.
		// Therefore, a query at height 2 would be equivalent to a query at height 3.
		// A height of 0 will query with the lastest state.
		nil, *new(clienttypes.Height), nil
}

// Use the IAVL height if a valid tendermint height is passed in.
// A height of 0 will query with the latest state.

// #nosec G115 --- overflow checked above

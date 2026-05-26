package client

import (
	rpchttp "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/client/http"
	"github.com/spf13/pflag"

	"github.com/sei-protocol/sei-chain/sei-cosmos/types/query"
)

// Paginate returns the correct starting and ending index for a paginated query,
// given that client provides a desired page and limit of objects and the handler
// provides the total number of objects. The start page is assumed to be 1-indexed.
// If the start page is invalid, non-positive values are returned signaling the
// request is invalid; it returns non-positive values if limit is non-positive and
// defLimit is negative.
func Paginate(numObjs, page, limit, defLimit int) (start, end int) {
	_ = "STUB: not implemented"

	// invalid start page
	return 0, 0
}

// fallback to default limit if supplied limit is invalid

// invalid default limit

// page is out of bounds

// ReadPageRequest reads and builds the necessary page request flags for pagination.
func ReadPageRequest(flagSet *pflag.FlagSet) (*query.PageRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewClientFromNode sets up Client implementation that communicates with a Tendermint node over
// JSON RPC and WebSockets
// TODO: We might not need to manually append `/websocket`:
// https://github.com/cosmos/cosmos-sdk/issues/8986
func NewClientFromNode(nodeURI string) (*rpchttp.HTTP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

package utils

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
	ibctmtypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/light-clients/07-tendermint/types"
)

// QueryClientState returns a client state. If prove is true, it performs an ABCI store query
// in order to retrieve the merkle proof. Otherwise, it uses the gRPC query client.
func QueryClientState(
	clientCtx client.Context, clientID string, prove bool,
) (*types.QueryClientStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// QueryClientStateABCI queries the store to get the light client state and a merkle proof.
func QueryClientStateABCI(
	clientCtx client.Context, clientID string,
) (*types.QueryClientStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if client exists

// QueryConsensusState returns a consensus state. If prove is true, it performs an ABCI store
// query in order to retrieve the merkle proof. Otherwise, it uses the gRPC query client.
func QueryConsensusState(
	clientCtx client.Context, clientID string, height exported.Height, prove, latestHeight bool,
) (*types.QueryConsensusStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// QueryConsensusStateABCI queries the store to get the consensus state of a light client and a
// merkle proof of its existence or non-existence.
func QueryConsensusStateABCI(
	clientCtx client.Context, clientID string, height exported.Height,
) (*types.QueryConsensusStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if consensus state exists

// QueryTendermintHeader takes a client context and returns the appropriate
// tendermint header
func QueryTendermintHeader(clientCtx client.Context) (ibctmtypes.Header, int64, error) {
	_ = "STUB: not implemented"
	return *new(ibctmtypes.Header), 0, nil
}

// QuerySelfConsensusState takes a client context and returns the appropriate
// tendermint consensus state
func QuerySelfConsensusState(clientCtx client.Context) (*ibctmtypes.ConsensusState, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

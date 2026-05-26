package utils

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"

	clienttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/03-connection/types"
	commitmenttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/23-commitment/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

// QueryConnection returns a connection end.
// If prove is true, it performs an ABCI store query in order to retrieve the merkle proof. Otherwise,
// it uses the gRPC query client.
func QueryConnection(
	clientCtx client.Context, connectionID string, prove bool,
) (*types.QueryConnectionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryConnectionABCI(clientCtx client.Context, connectionID string) (*types.QueryConnectionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if connection exists

// QueryClientConnections queries the connection paths registered for a particular client.
// If prove is true, it performs an ABCI store query in order to retrieve the merkle proof. Otherwise,
// it uses the gRPC query client.
func QueryClientConnections(
	clientCtx client.Context, clientID string, prove bool,
) (*types.QueryClientConnectionsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryClientConnectionsABCI(clientCtx client.Context, clientID string) (*types.QueryClientConnectionsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if connection paths exist

// QueryConnectionClientState returns the ClientState of a connection end. If
// prove is true, it performs an ABCI store query in order to retrieve the
// merkle proof. Otherwise, it uses the gRPC query client.
func QueryConnectionClientState(
	clientCtx client.Context, connectionID string, prove bool,
) (*types.QueryConnectionClientStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// use client state returned from ABCI query in case query height differs

// QueryConnectionConsensusState returns the ConsensusState of a connection end. If
// prove is true, it performs an ABCI store query in order to retrieve the
// merkle proof. Otherwise, it uses the gRPC query client.
func QueryConnectionConsensusState(
	clientCtx client.Context, connectionID string, height clienttypes.Height, prove bool,
) (*types.QueryConnectionConsensusStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseClientState unmarshals a cmd input argument from a JSON string to a client state
// If the input is not a JSON, it looks for a path to the JSON file
func ParseClientState(cdc *codec.LegacyAmino, arg string) (exported.ClientState, error) {
	_ = "STUB: not implemented"
	return *new(exported.ClientState), nil
}

// check for file path if JSON input is not provided

// ParsePrefix unmarshals an cmd input argument from a JSON string to a commitment
// Prefix. If the input is not a JSON, it looks for a path to the JSON file.
func ParsePrefix(cdc *codec.LegacyAmino, arg string) (commitmenttypes.MerklePrefix, error) {
	_ = "STUB: not implemented"
	return *new(commitmenttypes.MerklePrefix), nil
}

// check for file path if JSON input is not provided

// ParseProof unmarshals a cmd input argument from a JSON string to a commitment
// Proof. If the input is not a JSON, it looks for a path to the JSON file. It
// then marshals the commitment proof into a proto encoded byte array.
func ParseProof(cdc *codec.LegacyAmino, arg string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check for file path if JSON input is not provided

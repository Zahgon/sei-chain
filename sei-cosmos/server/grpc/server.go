package grpc

import (
	"google.golang.org/grpc"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/server/types"
)

// StartGRPCServer starts a gRPC server on the given address.
func StartGRPCServer(clientCtx client.Context, app types.Application, address string) (*grpc.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// reflection allows consumers to build dynamic clients that can write
// to any cosmos-sdk application without relying on application packages at compile time

// Reflection allows external clients to see what services and methods
// the gRPC server exposes.

// assume server started successfully

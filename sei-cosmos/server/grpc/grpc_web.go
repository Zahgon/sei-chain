package grpc

import (
	"net/http"

	"google.golang.org/grpc"

	"github.com/sei-protocol/sei-chain/sei-cosmos/server/config"
)

// StartGRPCWeb starts a gRPC-Web server on the given address.
func StartGRPCWeb(grpcSrv *grpc.Server, config config.Config) (*http.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// assume server started successfully

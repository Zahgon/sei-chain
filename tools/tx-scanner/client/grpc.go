package client

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/client/grpc/tmservice"
	txtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/tx"
	"google.golang.org/grpc"
)

var GrpcConn *grpc.ClientConn

func InitializeGRPCClient(targetEndpoint string, port int) { _ = "STUB: not implemented"; return }

// Use default insecure if we don't have credentials setup

// spin up goroutine for monitoring and reconnect purposes

func GetTmServiceClient() tmservice.ServiceClient {
	_ = "STUB: not implemented"
	return *new(tmservice.ServiceClient)
}

func GetTxClient() txtypes.ServiceClient {
	_ = "STUB: not implemented"
	return *new(txtypes.ServiceClient)
}

package grpc

import (
	"context"

	"github.com/sei-protocol/seilog"

	"google.golang.org/grpc"

	"github.com/sei-protocol/sei-chain/sei-tendermint/config"
)

var logger = seilog.NewLogger("tendermint", "privval", "grpc")

// DefaultDialOptions constructs a list of grpc dial options
func DefaultDialOptions(
	extraOpts ...grpc.DialOption,
) []grpc.DialOption {
	_ = "STUB: not implemented"
	return nil
}

// 50 * 100ms = 5s total

// Default 5Mb

// send pings every 10 seconds if there is no activity
// wait 2 seconds for ping ack before considering the connection dead

func GenerateTLS(certPath, keyPath, ca string) grpc.DialOption {
	_ = "STUB: not implemented"
	return *new(grpc.DialOption)
}

// DialRemoteSigner is  a generalized function to dial the gRPC server.
func DialRemoteSigner(
	ctx context.Context,
	cfg *config.PrivValidatorConfig,
	chainID string,
	usePrometheus bool,
) (*SignerClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

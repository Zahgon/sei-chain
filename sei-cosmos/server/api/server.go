package api

import (
	"net"
	"sync"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sei-protocol/seilog"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/server/config"
	"github.com/sei-protocol/sei-chain/sei-cosmos/telemetry"

	// unnamed import of statik for swagger UI support
	_ "github.com/sei-protocol/sei-chain/sei-cosmos/client/docs/statik"
)

var logger = seilog.NewLogger("tendermint", "server", "api")

// Server defines the server's API interface.
type Server struct {
	Router            *mux.Router
	GRPCGatewayRouter *runtime.ServeMux
	ClientCtx         client.Context

	metrics *telemetry.Metrics
	// Start() is blocking and generally called from a separate goroutine.
	// Close() can be called asynchronously and access shared memory
	// via the listener. Therefore, we sync access to Start and Close with
	// this mutex to avoid data races.
	mtx      sync.Mutex
	listener net.Listener
}

// CustomGRPCHeaderMatcher for mapping request headers to
// GRPC metadata.
// HTTP headers that start with 'Grpc-Metadata-' are automatically mapped to
// gRPC metadata after removing prefix 'Grpc-Metadata-'. We can use this
// CustomGRPCHeaderMatcher if headers don't start with `Grpc-Metadata-`
func CustomGRPCHeaderMatcher(key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func New(clientCtx client.Context) *Server {
	_ = "STUB: not implemented"
	// The default JSON marshaller used by the gRPC-Gateway is unable to marshal non-nullable non-scalar fields.
	// Using the gogo/gateway package with the gRPC-Gateway WithMarshaler option fixes the scalar field marshalling issue.
	return nil
}

// Custom marshaler option is required for gogo proto

// This is necessary to get error details properly
// marshalled in unary requests.

// Custom header matcher for mapping request headers to
// GRPC metadata

// Start starts the API server. Internally, the API server leverages Tendermint's
// JSON RPC server. Configuration options are provided via config.APIConfig
// and are delegated to the Tendermint JSON RPC server. The process is
// non-blocking, so an external signal handler must be used.
func (s *Server) Start(cfg config.Config, apiMetrics *telemetry.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // config value, validated at startup
//nolint:gosec // config value, validated at startup
//nolint:gosec // config value, validated at startup
//nolint:gosec // config value, validated at startup

// Close closes the API server.
func (s *Server) Close() error { _ = "STUB: not implemented"; return nil }

func (s *Server) registerGRPCGatewayRoutes() { _ = "STUB: not implemented"; return }

func (s *Server) registerMetrics() { _ = "STUB: not implemented"; return }

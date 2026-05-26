package admin

import (
	"github.com/sei-protocol/seilog"
	"google.golang.org/grpc"
)

var logger = seilog.NewLogger("admin")

// StartServer creates and starts a dedicated admin gRPC server on the given
// loopback address. Returns the server so the caller can stop it on shutdown.
func StartServer(address string) (*grpc.Server, error) { _ = "STUB: not implemented"; return nil, nil }

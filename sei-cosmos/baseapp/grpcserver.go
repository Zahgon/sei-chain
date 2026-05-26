package baseapp

import (
	gogogrpc "github.com/gogo/protobuf/grpc"
)

// GRPCQueryRouter returns the GRPCQueryRouter of a BaseApp.
func (app *BaseApp) GRPCQueryRouter() *GRPCQueryRouter { _ = "STUB: not implemented"; return nil }

// RegisterGRPCServer registers gRPC services directly with the gRPC server.
func (app *BaseApp) RegisterGRPCServer(server gogogrpc.Server) {
	_ = "STUB: not implemented"
	// Define an interceptor for all gRPC queries: this interceptor will create
	// a new sdk.Context, and pass it into the query handler.
	return
}

// If there's some metadata in the context, retrieve it.

// Get height header from the request context, if present.

// Create the sdk.Context. Passing false as 2nd arg, as we can't
// actually support proofs with gRPC right now.

// Add relevant gRPC headers

// If height was not set in the request, set it to the latest

// Attach the sdk.Context into the gRPC's context.Context.

// Loop through all services and methods, add the interceptor, and register
// the service.

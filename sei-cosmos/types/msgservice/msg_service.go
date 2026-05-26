package msgservice

import (
	"context"

	"google.golang.org/grpc"

	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
)

// RegisterMsgServiceDesc registers all type_urls from Msg services described
// in `sd` into the registry.
func RegisterMsgServiceDesc(registry codectypes.InterfaceRegistry, sd *grpc.ServiceDesc) {
	_ = "STUB: not implemented"
	// Adds a top-level type_url based on the Msg service name.
	return
}

// NOTE: This is how we pull the concrete request type for each handler for registering in the InterfaceRegistry.
// This approach is maybe a bit hacky, but less hacky than reflecting on the handler object itself.
// We use a no-op interceptor to avoid actually calling into the handler itself.

// We panic here because there is no other alternative and the app cannot be initialized correctly
// this should only happen if there is a problem with code generation in which case the app won't
// work correctly anyway.

// gRPC NOOP interceptor
func noopInterceptor(_ context.Context, _ interface{}, _ *grpc.UnaryServerInfo, _ grpc.UnaryHandler) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

package baseapp

import (
	"context"

	gogogrpc "github.com/gogo/protobuf/grpc"
	"google.golang.org/grpc"

	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// MsgServiceRouter routes fully-qualified Msg service methods to their handler.
type MsgServiceRouter struct {
	interfaceRegistry codectypes.InterfaceRegistry
	routes            map[string]MsgServiceHandler
}

var _ gogogrpc.Server = &MsgServiceRouter{}

// NewMsgServiceRouter creates a new MsgServiceRouter.
func NewMsgServiceRouter() *MsgServiceRouter { _ = "STUB: not implemented"; return nil }

// MsgServiceHandler defines a function type which handles Msg service message.
type MsgServiceHandler = func(ctx sdk.Context, req sdk.Msg) (*sdk.Result, error)

// Handler returns the MsgServiceHandler for a given msg or nil if not found.
func (msr *MsgServiceRouter) Handler(msg sdk.Msg) MsgServiceHandler {
	_ = "STUB: not implemented"
	return *new(MsgServiceHandler)
}

// HandlerByTypeURL returns the MsgServiceHandler for a given query route path or nil
// if not found.
func (msr *MsgServiceRouter) HandlerByTypeURL(typeURL string) MsgServiceHandler {
	_ = "STUB: not implemented"
	return *new(MsgServiceHandler)
}

// RegisterService implements the gRPC Server.RegisterService method. sd is a gRPC
// service description, handler is an object which implements that gRPC service.
//
// This function PANICs:
//   - if it is called before the service `Msg`s have been registered using
//     RegisterInterfaces,
//   - or if a service is being registered twice.
func (msr *MsgServiceRouter) RegisterService(sd *grpc.ServiceDesc, handler interface{}) {
	_ = "STUB: not implemented"
	// Adds a top-level query handler based on the gRPC service name.
	return
}

// NOTE: This is how we pull the concrete request type for each handler for registering in the InterfaceRegistry.
// This approach is maybe a bit hacky, but less hacky than reflecting on the handler object itself.
// We use a no-op interceptor to avoid actually calling into the handler itself.

// We panic here because there is no other alternative and the app cannot be initialized correctly
// this should only happen if there is a problem with code generation in which case the app won't
// work correctly anyway.

// Check that the service Msg fully-qualified method name has already
// been registered (via RegisterInterfaces). If the user registers a
// service without registering according service Msg type, there might be
// some unexpected behavior down the road. Since we can't return an error
// (`Server.RegisterService` interface restriction) we panic (at startup).

// Check that each service is only registered once. If a service is
// registered more than once, then we should error. Since we can't
// return an error (`Server.RegisterService` interface restriction) we
// panic (at startup).

// Call the method handler from the service description with the handler object.
// We don't do any decoding here because the decoding was already done.

// SetInterfaceRegistry sets the interface registry for the router.
func (msr *MsgServiceRouter) SetInterfaceRegistry(interfaceRegistry codectypes.InterfaceRegistry) {
	_ = "STUB: not implemented"
	return
}

func noopDecoder(_ interface{}) error { _ = "STUB: not implemented"; return nil }
func noopInterceptor(_ context.Context, _ interface{}, _ *grpc.UnaryServerInfo, _ grpc.UnaryHandler) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type CoinInterface interface {
	GetAmount() sdk.Coin
}

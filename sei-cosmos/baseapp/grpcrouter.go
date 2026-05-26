package baseapp

import (
	gogogrpc "github.com/gogo/protobuf/grpc"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/encoding/proto"

	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

var protoCodec = encoding.GetCodec(proto.Name)

// GRPCQueryRouter routes ABCI Query requests to GRPC handlers
type GRPCQueryRouter struct {
	routes            map[string]GRPCQueryHandler
	interfaceRegistry codectypes.InterfaceRegistry
	serviceData       []serviceData
}

// serviceData represents a gRPC service, along with its handler.
type serviceData struct {
	serviceDesc *grpc.ServiceDesc
	handler     interface{}
}

var _ gogogrpc.Server = &GRPCQueryRouter{}

// NewGRPCQueryRouter creates a new GRPCQueryRouter
func NewGRPCQueryRouter() *GRPCQueryRouter { _ = "STUB: not implemented"; return nil }

// GRPCQueryHandler defines a function type which handles ABCI Query requests
// using gRPC
type GRPCQueryHandler = func(ctx sdk.Context, req abci.RequestQuery) (abci.ResponseQuery, error)

// Route returns the GRPCQueryHandler for a given query route path or nil
// if not found
func (qrt *GRPCQueryRouter) Route(path string) GRPCQueryHandler {
	_ = "STUB: not implemented"
	return *new(GRPCQueryHandler)
}

// RegisterService implements the gRPC Server.RegisterService method. sd is a gRPC
// service description, handler is an object which implements that gRPC service/
//
// This functions PANICS:
// - if a protobuf service is registered twice.
func (qrt *GRPCQueryRouter) RegisterService(sd *grpc.ServiceDesc, handler interface{}) {
	_ = "STUB: not implemented"
	// adds a top-level query handler based on the gRPC service name
	return
}

// Check that each service is only registered once. If a service is
// registered more than once, then we should error. Since we can't
// return an error (`Server.RegisterService` interface restriction) we
// panic (at startup).

// call the method handler from the service description with the handler object,
// a wrapped sdk.Context with proto-unmarshaled data from the ABCI request data

// proto marshal the result bytes

// return the result bytes as the response value

// SetInterfaceRegistry sets the interface registry for the router. This will
// also register the interface reflection gRPC service.
func (qrt *GRPCQueryRouter) SetInterfaceRegistry(interfaceRegistry codectypes.InterfaceRegistry) {
	_ = "STUB: not implemented"
	return
}

// Once we have an interface registry, we can register the interface
// registry reflection gRPC service.

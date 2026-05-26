package reflection

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
)

type reflectionServiceServer struct {
	interfaceRegistry types.InterfaceRegistry
}

// NewReflectionServiceServer creates a new reflectionServiceServer.
func NewReflectionServiceServer(interfaceRegistry types.InterfaceRegistry) ReflectionServiceServer {
	_ = "STUB: not implemented"
	return *new(ReflectionServiceServer)
}

var _ ReflectionServiceServer = (*reflectionServiceServer)(nil)

// ListAllInterfaces implements the ListAllInterfaces method of the
// ReflectionServiceServer interface.
func (r reflectionServiceServer) ListAllInterfaces(_ context.Context, _ *ListAllInterfacesRequest) (*ListAllInterfacesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListImplementations implements the ListImplementations method of the
// ReflectionServiceServer interface.
func (r reflectionServiceServer) ListImplementations(_ context.Context, req *ListImplementationsRequest) (*ListImplementationsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

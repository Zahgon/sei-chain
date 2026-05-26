package baseapp

import (
	gocontext "context"

	gogogrpc "github.com/gogo/protobuf/grpc"
	"google.golang.org/grpc"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// QueryServiceTestHelper provides a helper for making grpc query service
// rpc calls in unit tests. It implements both the grpc Server and ClientConn
// interfaces needed to register a query service server and create a query
// service client.
type QueryServiceTestHelper struct {
	*GRPCQueryRouter
	Ctx sdk.Context
}

var (
	_ gogogrpc.Server     = &QueryServiceTestHelper{}
	_ gogogrpc.ClientConn = &QueryServiceTestHelper{}
)

// NewQueryServerTestHelper creates a new QueryServiceTestHelper that wraps
// the provided sdk.Context
func NewQueryServerTestHelper(ctx sdk.Context, interfaceRegistry types.InterfaceRegistry) *QueryServiceTestHelper {
	_ = "STUB: not implemented"
	return nil
}

// Invoke implements the grpc ClientConn.Invoke method
func (q *QueryServiceTestHelper) Invoke(_ gocontext.Context, method string, args, reply interface{}, _ ...grpc.CallOption) error {
	_ = "STUB: not implemented"
	return nil
}

// NewStream implements the grpc ClientConn.NewStream method
func (q *QueryServiceTestHelper) NewStream(gocontext.Context, *grpc.StreamDesc, string, ...grpc.CallOption) (grpc.ClientStream, error) {
	_ = "STUB: not implemented"
	return *new(grpc.ClientStream), nil
}

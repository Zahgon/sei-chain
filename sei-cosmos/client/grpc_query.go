package client

import (
	gocontext "context"

	gogogrpc "github.com/gogo/protobuf/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/encoding/proto"
)

var _ gogogrpc.ClientConn = Context{}

var protoCodec = encoding.GetCodec(proto.Name)

// Invoke implements the grpc ClientConn.Invoke method
func (ctx Context) Invoke(grpcCtx gocontext.Context, method string, req, reply interface{}, opts ...grpc.CallOption) error {
	_ = "STUB: not implemented"
	// Two things can happen here:
	// 1. either we're broadcasting a Tx, in which call we call Tendermint's broadcast endpoint directly,
	// 2. or we are querying for state, in which case we call ABCI's Query.
	return nil
}

// In both cases, we don't allow empty request args (it will panic unexpectedly).

// Case 1. Broadcasting a Tx.

// Case 2. Querying state.

// parse height header

// Create header metadata. For now the headers contain:
// - block height
// We then parse all the call options, if the call option is a
// HeaderCallOption, then we manually set the value of that header to the
// metadata.

// NewStream implements the grpc ClientConn.NewStream method
func (Context) NewStream(gocontext.Context, *grpc.StreamDesc, string, ...grpc.CallOption) (grpc.ClientStream, error) {
	_ = "STUB: not implemented"
	return *new(grpc.ClientStream), nil
}

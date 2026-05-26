package tx

import (
	"context"

	gogogrpc "github.com/gogo/protobuf/grpc" // nolint: staticcheck
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	txtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/tx"
)

// baseAppSimulateFn is the signature of the Baseapp#Simulate function.
type baseAppSimulateFn func(txBytes []byte) (sdk.GasInfo, *sdk.Result, error)

// txServer is the server for the protobuf Tx service.
type txServer struct {
	node              client.LocalClient
	txConfig          client.TxConfig
	simulate          baseAppSimulateFn
	interfaceRegistry codectypes.InterfaceRegistry
}

// NewTxServer creates a new Tx service server.
func NewTxServer(node client.LocalClient, txConfig client.TxConfig, simulate baseAppSimulateFn, interfaceRegistry codectypes.InterfaceRegistry) txtypes.ServiceServer {
	_ = "STUB: not implemented"
	return *new(txtypes.ServiceServer)
}

var _ txtypes.ServiceServer = txServer{}

const (
	eventFormat = "{eventType}.{eventAttribute}={value}"
)

// TxsByEvents implements the ServiceServer.TxsByEvents RPC method.
func (s txServer) GetTxsEvent(ctx context.Context, req *txtypes.GetTxsEventRequest) (*txtypes.GetTxsEventResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create a proto codec, we need it to unmarshal the tx bytes.

// Simulate implements the ServiceServer.Simulate RPC method.
func (s txServer) Simulate(ctx context.Context, req *txtypes.SimulateRequest) (*txtypes.SimulateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This block is for backwards-compatibility.
// We used to support passing a `Tx` in req. But if we do that, sig
// verification might not pass, because the .Marshal() below might not
// be the same marshaling done by the client.

// GetTx implements the ServiceServer.GetTx RPC method.
func (s txServer) GetTx(ctx context.Context, req *txtypes.GetTxRequest) (*txtypes.GetTxResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO We should also check the proof flag in gRPC header.
// https://github.com/cosmos/cosmos-sdk/issues/7036.

// protoTxProvider is a type which can provide a proto transaction. It is a
// workaround to get access to the wrapper TxBuilder's method GetProtoTx().
// ref: https://github.com/cosmos/cosmos-sdk/issues/10347
type protoTxProvider interface {
	GetProtoTx() *txtypes.Tx
}

// GetBlockWithTxs returns a block with decoded txs.
func (s txServer) GetBlockWithTxs(ctx context.Context, req *txtypes.GetBlockWithTxsRequest) (*txtypes.GetBlockWithTxsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s txServer) BroadcastTx(ctx context.Context, req *txtypes.BroadcastTxRequest) (*txtypes.BroadcastTxResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RegisterTxService registers the tx service on the gRPC router.
func RegisterTxService(
	qrt gogogrpc.Server,
	node client.LocalClient,
	txConfig client.TxConfig,
	simulateFn baseAppSimulateFn,
	interfaceRegistry codectypes.InterfaceRegistry,
) {
	_ = "STUB: not implemented"
	return
}

// RegisterGRPCGatewayRoutes mounts the tx service's GRPC-gateway routes on the
// given Mux.
func RegisterGRPCGatewayRoutes(clientConn gogogrpc.ClientConn, mux *runtime.ServeMux) {
	_ = "STUB: not implemented"
	return
}

func parseOrderBy(orderBy txtypes.OrderBy) string { _ = "STUB: not implemented"; return "" }

// Defaults to Tendermint's default, which is `asc` now.

package tmservice

import (
	"context"

	gogogrpc "github.com/gogo/protobuf/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
)

func GetProtoBlock(ctx context.Context, node client.Client, height *int64) (tmproto.BlockID, *tmproto.Block, error) {
	_ = "STUB: not implemented"
	return *new(tmproto.BlockID), nil, nil
}

// This is the struct that we will implement all the handlers on.
type queryServer struct {
	node              client.Client
	interfaceRegistry codectypes.InterfaceRegistry
}

var _ ServiceServer = queryServer{}
var _ codectypes.UnpackInterfacesMessage = &GetLatestValidatorSetResponse{}

// NewQueryServer creates a new tendermint query server.
func NewQueryServer(node client.Client, interfaceRegistry codectypes.InterfaceRegistry) ServiceServer {
	_ = "STUB: not implemented"
	return *new(ServiceServer)
}

// GetSyncing implements ServiceServer.GetSyncing
func (s queryServer) GetSyncing(ctx context.Context, _ *GetSyncingRequest) (*GetSyncingResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetLatestBlock implements ServiceServer.GetLatestBlock
func (s queryServer) GetLatestBlock(ctx context.Context, _ *GetLatestBlockRequest) (*GetLatestBlockResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetBlockByHeight implements ServiceServer.GetBlockByHeight
func (s queryServer) GetBlockByHeight(ctx context.Context, req *GetBlockByHeightRequest) (*GetBlockByHeightResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetLatestValidatorSet implements ServiceServer.GetLatestValidatorSet
func (s queryServer) GetLatestValidatorSet(ctx context.Context, req *GetLatestValidatorSetRequest) (*GetLatestValidatorSetResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *GetLatestValidatorSetResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// GetValidatorSetByHeight implements ServiceServer.GetValidatorSetByHeight
func (s queryServer) GetValidatorSetByHeight(ctx context.Context, req *GetValidatorSetByHeightRequest) (*GetValidatorSetByHeightResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validatorsOutput(ctx context.Context, node client.Client, height *int64, page, limit int) (*GetLatestValidatorSetResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetNodeInfo implements ServiceServer.GetNodeInfo
func (s queryServer) GetNodeInfo(ctx context.Context, req *GetNodeInfoRequest) (*GetNodeInfoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RegisterTendermintService registers the tendermint queries on the gRPC router.
func RegisterTendermintService(
	qrt gogogrpc.Server,
	node client.LocalClient,
	interfaceRegistry codectypes.InterfaceRegistry,
) {
	_ = "STUB: not implemented"
	return
}

// RegisterGRPCGatewayRoutes mounts the tendermint service's GRPC-gateway routes on the
// given Mux.
func RegisterGRPCGatewayRoutes(clientConn gogogrpc.ClientConn, mux *runtime.ServeMux) {
	_ = "STUB: not implemented"
	return
}

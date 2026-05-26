package v2alpha1

import (
	"context"

	"google.golang.org/grpc"

	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

type Config struct {
	SigningModes      map[string]int32
	ChainID           string
	SdkConfig         *sdk.Config
	InterfaceRegistry codectypes.InterfaceRegistry
}

// Register registers the cosmos sdk reflection service
// to the provided *grpc.Server given a Config
func Register(srv *grpc.Server, conf Config) error { _ = "STUB: not implemented"; return nil }

type reflectionServiceServer struct {
	desc *AppDescriptor
}

func (r reflectionServiceServer) GetAuthnDescriptor(_ context.Context, _ *GetAuthnDescriptorRequest) (*GetAuthnDescriptorResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r reflectionServiceServer) GetChainDescriptor(_ context.Context, _ *GetChainDescriptorRequest) (*GetChainDescriptorResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r reflectionServiceServer) GetCodecDescriptor(_ context.Context, _ *GetCodecDescriptorRequest) (*GetCodecDescriptorResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r reflectionServiceServer) GetConfigurationDescriptor(_ context.Context, _ *GetConfigurationDescriptorRequest) (*GetConfigurationDescriptorResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r reflectionServiceServer) GetQueryServicesDescriptor(_ context.Context, _ *GetQueryServicesDescriptorRequest) (*GetQueryServicesDescriptorResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r reflectionServiceServer) GetTxDescriptor(_ context.Context, _ *GetTxDescriptorRequest) (*GetTxDescriptorResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newReflectionServiceServer(grpcSrv *grpc.Server, conf Config) (reflectionServiceServer, error) {
	_ = "STUB: not implemented"
	// set chain descriptor
	return *new(reflectionServiceServer), nil
}

// set configuration descriptor

// set codec descriptor

// set query service descriptor

// set deliver descriptor

// newCodecDescriptor describes the codec given the codectypes.InterfaceRegistry
func newCodecDescriptor(ir codectypes.InterfaceRegistry) (*CodecDescriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NOTE(fdymylja): this could be filled, but it won't be filled as of now
// doing this would require us to fully rebuild in a (dependency) transitive way the proto
// registry of the supported proto.Messages for the application, this could be easily
// done if we weren't relying on gogoproto which does not allow us to iterate over the
// registry. Achieving this right now would mean to start slowly building descriptors
// getting their files dependencies, building those dependencies then rebuilding the
// descriptor builder. It's too much work as of now.

func newQueryServiceDescriptor(srv *grpc.Server) *QueryServicesDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func newTxDescriptor(ir codectypes.InterfaceRegistry) (*TxDescriptor, error) {
	_ = "STUB: not implemented"
	// get base tx type name
	return nil, nil
}

// get msgs

// process sdk.Msg

func newAuthnDescriptor(signingModes map[string]int32) *AuthnDescriptor {
	_ = "STUB: not implemented"
	return nil
}

// NOTE(fdymylja): this cannot be filled as of now, auth and the sdk itself don't support as of now
// a service which allows to get authentication metadata for the provided sign mode.

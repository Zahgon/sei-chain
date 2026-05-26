package privval

import (
	"context"
	"sync"

	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/service"
	privvalproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/privval"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// ValidationRequestHandlerFunc handles different remoteSigner requests
type ValidationRequestHandlerFunc func(
	ctx context.Context,
	privVal types.PrivValidator,
	requestMessage privvalproto.Message,
	chainID string) (privvalproto.Message, error)

type SignerServer struct {
	service.BaseService

	endpoint *SignerDialerEndpoint
	chainID  string
	privVal  types.PrivValidator

	handlerMtx               sync.Mutex
	validationRequestHandler ValidationRequestHandlerFunc
}

func NewSignerServer(endpoint *SignerDialerEndpoint, chainID string, privVal types.PrivValidator) *SignerServer {
	_ = "STUB: not implemented"
	return nil
}

// OnStart implements service.Service.
func (ss *SignerServer) OnStart(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// OnStop implements service.Service.
func (ss *SignerServer) OnStop() { _ = "STUB: not implemented"; return }

// SetRequestHandler override the default function that is used to service requests
func (ss *SignerServer) SetRequestHandler(validationRequestHandler ValidationRequestHandlerFunc) {
	_ = "STUB: not implemented"
	return
}

func (ss *SignerServer) servicePendingRequest(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// Ignore error from closing.

// limit the scope of the lock

// todo

// only log the error; we'll reply with an error in res

func (ss *SignerServer) serviceLoop(ctx context.Context) { _ = "STUB: not implemented"; return }

package rpc

import (
	"context"
	"reflect"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p/conn"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p/mux"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/protoutils"
	"golang.org/x/time/rate"
)

type InBytes uint64
type InMsgs uint64

type Msg[M protoutils.Message] struct {
	MsgSize InBytes
	Window  InMsgs
}

type Limit struct {
	Rate       rate.Limit
	Concurrent uint64
}

type RPC[API any, Req, Resp protoutils.Message] struct {
	Kind  mux.StreamKind
	Limit Limit
	Req   Msg[Req]
	Resp  Msg[Resp]
}

type rpcConfig struct {
	limit Limit
}

type service map[mux.StreamKind]*rpcConfig

func (s service) muxServerConfig() *mux.Config { _ = "STUB: not implemented"; return nil }

var registry = map[reflect.Type]service{}

func (s service) muxClientConfig() *mux.Config { _ = "STUB: not implemented"; return nil }

func Register[API any, Req, Resp protoutils.Message](kind mux.StreamKind, limit Limit, req Msg[Req], resp Msg[Resp]) *RPC[API, Req, Resp] {
	_ = "STUB: not implemented"
	return nil
}

// Simplification: we allow the same number of streams in each direction.

type Server[API any] struct{ mux *mux.Mux }

func NewServer[API any]() Server[API] { _ = "STUB: not implemented"; return nil }

func (s Server[API]) Run(ctx context.Context, conn conn.Conn) error {
	_ = "STUB: not implemented"
	return nil
}

type Client[API any] struct{ mux *mux.Mux }

func NewClient[API any]() Client[API] { _ = "STUB: not implemented"; return nil }

func (c Client[API]) Run(ctx context.Context, conn conn.Conn) error {
	_ = "STUB: not implemented"
	return nil

	// TODO: add client-size rate limiting.
}

func (r *RPC[API, Req, Resp]) Call(ctx context.Context, client Client[API]) (Stream[Req, Resp], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RPC[API, Req, Resp]) Serve(ctx context.Context, server Server[API], handler func(context.Context, Stream[Resp, Req]) error) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // Concurrent is a bounded config value

//nolint:gosec // MsgSize and Window are validated config values

type Stream[SendT, RecvT protoutils.Message] struct{ inner *mux.Stream }

func (s Stream[SendT, RecvT]) Close() { _ = "STUB: not implemented"; return }
func (s Stream[SendT, RecvT]) Send(ctx context.Context, msg SendT) error {
	_ = "STUB: not implemented"
	return nil
}

func (s Stream[SendT, RecvT]) Recv(ctx context.Context) (RecvT, error) {
	_ = "STUB: not implemented"
	return *new(RecvT), nil
}

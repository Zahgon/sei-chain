package p2p

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p/conn"

	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

type handshakedConn struct {
	conn *conn.SecretConnection
	msg  *handshakeMsg
}

func handshake(ctx context.Context, c conn.Conn, key NodeSecretKey, spec handshakeSpec) (*handshakedConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// handshake handshakes with a peer, validating the peer's information. If
// dialAddr is given, we check that the peer's info matches it.
// Closes the tcpConn if case of any error.
func exchangeNodeInfo(ctx context.Context, hConn *handshakedConn, nodeInfo types.NodeInfo) (types.NodeInfo, error) {
	_ = "STUB: not implemented"
	return *new(types.NodeInfo), nil
}

// Marshalling should always succeed.

//nolint:gosec // MaxNodeInfoSize() returns a small positive constant

// Authenticate the peer first.

// Validate the received info.

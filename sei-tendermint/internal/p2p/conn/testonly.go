package conn

import (
	"context"
	"net/netip"

	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

type buf struct {
	addr                netip.AddrPort
	data                [1500]byte
	begin, end, flushed uint64
}

func (b *buf) capacity() int { _ = "STUB: not implemented"; return 0 }

//nolint:gosec // b.end-b.begin is bounded by len(b.data) which fits in int

func (b *buf) push(data []byte) int { _ = "STUB: not implemented"; return 0 }

//nolint:gosec // n is non-negative, derived from min of two non-negative values
//nolint:gosec // len(b.data) is always non-negative

//nolint:gosec // n is non-negative

func (b *buf) pop(data []byte) int { _ = "STUB: not implemented"; return 0 }

//nolint:gosec // flushed-begin represents buffered data size, expected to fit in int
//nolint:gosec // n is non-negative, derived from min of two non-negative values
//nolint:gosec // len(b.data) is always non-negative

//nolint:gosec // n is non-negative

var _ Conn = (*TestConn)(nil)

type TestConn struct {
	write *utils.Watch[*buf]
	read  *utils.Watch[*buf]
}

func (c *TestConn) Read(ctx context.Context, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *TestConn) Write(ctx context.Context, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *TestConn) Flush(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *TestConn) Close() { _ = "STUB: not implemented"; return }

func (c *TestConn) LocalAddr() netip.AddrPort {
	_ = "STUB: not implemented"
	return *new(netip.AddrPort)
}

func (c *TestConn) RemoteAddr() netip.AddrPort {
	_ = "STUB: not implemented"
	return *new(netip.AddrPort)
}

func NewTestConn() (*TestConn, *TestConn) { _ = "STUB: not implemented"; return nil, nil }

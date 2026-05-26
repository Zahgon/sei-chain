package conn

import (
	"context"
	"net/netip"
)

type Conn interface {
	LocalAddr() netip.AddrPort
	RemoteAddr() netip.AddrPort
	Read(ctx context.Context, data []byte) error
	Write(ctx context.Context, data []byte) error
	Flush(ctx context.Context) error
	Close()
}

func IsDisconnect(err error) bool { _ = "STUB: not implemented"; return false }

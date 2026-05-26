package tcp

import (
	"context"
	"net"
	"net/netip"
	"sync/atomic"

	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

type call struct {
	data []byte
	done chan bool
}

type Conn struct {
	writes chan call
	reads  chan call
	errors chan error
	conn   *net.TCPConn
}

func (c Conn) Read(ctx context.Context, data []byte) error { _ = "STUB: not implemented"; return nil }

// close the read half.
// wait for data ownership to be returned.

// wait for context to finish.

func (c Conn) Write(ctx context.Context, data []byte) error { _ = "STUB: not implemented"; return nil }

// close the write half.
// wait for data ownership to be returned.

// wait for context to finish.

func (c Conn) Flush(_ context.Context) error { _ = "STUB: not implemented"; return nil }
func (c Conn) Close()                        { _ = "STUB: not implemented"; return }

func (c Conn) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c Conn) LocalAddr() netip.AddrPort { _ = "STUB: not implemented"; return *new(netip.AddrPort) }

func (c Conn) RemoteAddr() netip.AddrPort { _ = "STUB: not implemented"; return *new(netip.AddrPort) }

// reserverAddrs is a global register of reserved ports.
//   - Some(fd) indicates that the port is not currently in use.
//     fd is the socket bound to the addr, which guards the port from being allocated to different process.
//   - None indicates that the port is currently in use.
//     Calling Listen() for this addr will result in error, until the current listener closes.
var reservedAddrs = utils.NewMutex(map[netip.AddrPort]utils.Option[int]{})

// IPv4Loopback returns the IPv4 loopback address.
func IPv4Loopback() netip.Addr { _ = "STUB: not implemented"; return *new(netip.Addr) }

func Dial(ctx context.Context, addr netip.AddrPort) (Conn, error) {
	_ = "STUB: not implemented"
	return *new(Conn), nil
}

type HostPort struct {
	Hostname string
	Port     uint16
}

func (hp HostPort) String() string { _ = "STUB: not implemented"; return "" }

// MarshalText implements the encoding.TextMarshaler interface.
func (hp HostPort) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (hp *HostPort) UnmarshalText(b []byte) error { _ = "STUB: not implemented"; return nil }

func ParseHostPort(hp string) (HostPort, error) {
	_ = "STUB: not implemented"
	return *new(HostPort), nil
}

func (hp HostPort) Resolve(ctx context.Context) ([]netip.AddrPort, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Listener struct {
	reserved atomic.Pointer[netip.AddrPort]
	inner    *net.TCPListener
}

func testBind(addr netip.AddrPort) int { _ = "STUB: not implemented"; return 0 }

// NONBLOCK and CLOEXEC for consistency with net.ListenConfig.Listen().

// NOTE: linux allows sharing REUSEPORT port across 0.0.0.0 and 127.0.0.1, macOS does not.
// NOTE: linux distributes incoming connections across REUSEPORT listeners,
// macOS doesn't care that the socket is not listening yet and doesn't even use round-robin.

func (l *Listener) Close() error {
	_ = "STUB: not implemented"
	// We use reserved to check if the listener is holding ownership of
	// a reserved port. Ownership is released with the first Close() call.
	return nil
}

// We close under lock to avoid the following race scenario:
// 1. old listener releases port
// 2. new listener acquires port
// 3. port is dialed (old listener still open)
// 4. old listener closes.

// Accepts an incoming TCP connection.
// Closes the listener if ctx is done before a connection is accepted.
func (l *Listener) AcceptOrClose(ctx context.Context) (Conn, error) {
	_ = "STUB: not implemented"
	return *new(Conn), nil
}

// Early error check. Close listener to terminate Accept.
// This task guarantees that either err of res are set (possibly both).

// If there were no error, then res contains an open connection.

// Otherwise close the listener (for consistency), and close the connection (if established).

// Listen opens a TCP listener on the given address.
// It takes into account the reserved addresses (in tests) and sets the SO_REUSEPORT.
// nolint: contextcheck
func Listen(addr netip.AddrPort) (*Listener, error) { _ = "STUB: not implemented"; return nil, nil }

// nolint:lll

// Backlog has to be large enough, so that test dials succeed on the first try.

// net.FileListener duplicates fd.

// Passing the background context is ok, because Listen is
// non-blocking if it doesn't need to resolve the address
// against a DNS server.

// TestReserveAddr (testonly) reserves a localhost port in ephemeral range to open a TCP listener on it.
// Reservation prevents race conditions with other processes.
func TestReserveAddr() netip.AddrPort { _ = "STUB: not implemented"; return *new(netip.AddrPort) }

// TestReservePort (testonly) reserves a port on the given ip in ephemeral range to open a TCP listener on it.
// Reservation prevents race conditions with other processes.
func TestReservePort(ip netip.Addr) netip.AddrPort {
	_ = "STUB: not implemented"
	return *new(netip.AddrPort)
}

//nolint:gosec // OS-assigned port is always in valid uint16 range [0, 65535]

//nolint:gosec // OS-assigned port is always in valid uint16 range [0, 65535]

func TestPipe() (Conn, Conn) { _ = "STUB: not implemented"; return *new(Conn), *new(Conn) }

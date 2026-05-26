package p2p

import (
	"net/netip"
	"sync"
	"time"
)

type connTracker struct {
	cache       map[netip.Addr]uint
	lastConnect map[netip.Addr]time.Time
	mutex       sync.RWMutex
	max         uint
	window      time.Duration
}

func newConnTracker(max uint, window time.Duration) *connTracker {
	_ = "STUB: not implemented"
	return nil
}

func (rat *connTracker) Len() int { _ = "STUB: not implemented"; return 0 }

func (rat *connTracker) AddConn(addrPort netip.AddrPort) error {
	_ = "STUB: not implemented"
	return nil
}

// if there is already at least one connection, check to
// see if it was established before within the window,
// and error if so.

func (rat *connTracker) RemoveConn(addrPort netip.AddrPort) { _ = "STUB: not implemented"; return }

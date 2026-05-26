package p2p

import (
	"net"

	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// ErrFilterTimeout indicates that a filter operation timed out.
type ErrFilterTimeout struct{}

func (e ErrFilterTimeout) Error() string { _ = "STUB: not implemented"; return "" }

// ErrRejected indicates that a Peer was rejected carrying additional
// information as to the reason.
type ErrRejected struct {
	addr              NodeAddress
	conn              net.Conn
	err               error
	id                types.NodeID
	isAuthFailure     bool
	isDuplicate       bool
	isFiltered        bool
	isIncompatible    bool
	isNodeInfoInvalid bool
	isSelf            bool
}

// Addr returns the NetAddress for the rejected Peer.
func (e ErrRejected) Addr() NodeAddress { _ = "STUB: not implemented"; return *new(NodeAddress) }

func (e ErrRejected) Error() string { _ = "STUB: not implemented"; return "" }

// IsAuthFailure when Peer authentication was unsuccessful.
func (e ErrRejected) IsAuthFailure() bool { _ = "STUB: not implemented"; return false }

// IsDuplicate when Peer ID or IP are present already.
func (e ErrRejected) IsDuplicate() bool { _ = "STUB: not implemented"; return false }

// IsFiltered when Peer ID or IP was filtered.
func (e ErrRejected) IsFiltered() bool {
	_ = "STUB: not implemented"

	// IsIncompatible when Peer NodeInfo is not compatible with our own.
	return false
}

func (e ErrRejected) IsIncompatible() bool { _ = "STUB: not implemented"; return false }

// IsNodeInfoInvalid when the sent NodeInfo is not valid.
func (e ErrRejected) IsNodeInfoInvalid() bool { _ = "STUB: not implemented"; return false }

// IsSelf when Peer is our own node.
func (e ErrRejected) IsSelf() bool {
	_ = "STUB: not implemented"

	// ErrSwitchDuplicatePeerID to be raised when a peer is connecting with a known
	// ID.
	return false
}

type ErrSwitchDuplicatePeerID struct {
	ID types.NodeID
}

func (e ErrSwitchDuplicatePeerID) Error() string { _ = "STUB: not implemented"; return "" }

// ErrSwitchDuplicatePeerIP to be raised whena a peer is connecting with a known
// IP.
type ErrSwitchDuplicatePeerIP struct {
	IP net.IP
}

func (e ErrSwitchDuplicatePeerIP) Error() string { _ = "STUB: not implemented"; return "" }

// ErrSwitchConnectToSelf to be raised when trying to connect to itself.
type ErrSwitchConnectToSelf struct {
	Addr *NodeAddress
}

func (e ErrSwitchConnectToSelf) Error() string { _ = "STUB: not implemented"; return "" }

type ErrSwitchAuthenticationFailure struct {
	Dialed *NodeAddress
	Got    types.NodeID
}

func (e ErrSwitchAuthenticationFailure) Error() string { _ = "STUB: not implemented"; return "" }

// ErrTransportClosed is raised when the Transport has been closed.
type ErrTransportClosed struct{}

func (e ErrTransportClosed) Error() string { _ = "STUB: not implemented"; return "" }

//-------------------------------------------------------------------

type ErrNetAddressNoID struct {
	Addr string
}

func (e ErrNetAddressNoID) Error() string { _ = "STUB: not implemented"; return "" }

type ErrNetAddressInvalid struct {
	Addr string
	Err  error
}

func (e ErrNetAddressInvalid) Error() string { _ = "STUB: not implemented"; return "" }

// ErrCurrentlyDialingOrExistingAddress indicates that we're currently
// dialing this address or it belongs to an existing peer.
type ErrCurrentlyDialingOrExistingAddress struct {
	Addr string
}

func (e ErrCurrentlyDialingOrExistingAddress) Error() string { _ = "STUB: not implemented"; return "" }

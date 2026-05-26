package types

import (
	"net/netip"

	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"
	tmp2p "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/p2p"
)

const (
	maxNodeInfoSize = 10240 // 10KB
	maxNumChannels  = 16    // plenty of room for upgrades, for now
)

// Max size of the NodeInfo struct
func MaxNodeInfoSize() int { _ = "STUB: not implemented"; return 0 }

// ProtocolVersion contains the protocol versions for the software.
type ProtocolVersion struct {
	P2P   uint64 `json:"p2p,string"`
	Block uint64 `json:"block,string"`
	App   uint64 `json:"app,string"`
}

//-------------------------------------------------------------

// NodeInfo is the basic node information exchanged
// between two peers during the Tendermint P2P handshake.
type NodeInfo struct {
	ProtocolVersion ProtocolVersion `json:"protocol_version"`

	// Authenticate
	NodeID NodeID `json:"id"` // authenticated identifier
	// TODO(gprusak): for some reason ListenAddr is unused. Why do we have it?
	ListenAddr string `json:"listen_addr"` // accepting incoming

	// Check compatibility.
	// Channels are HexBytes so easier to read as JSON
	Network string `json:"network"` // network/chain ID
	Version string `json:"version"` // major.minor.revision
	// FIXME: This should be changed to uint16 to be consistent with the updated channel type
	Channels bytes.HexBytes `json:"channels"` // channels this node knows about

	// ASCIIText fields
	Moniker string        `json:"moniker"` // arbitrary moniker
	Other   NodeInfoOther `json:"other"`   // other application specific data
}

// NodeInfoOther is the misc. applcation specific data
type NodeInfoOther struct {
	TxIndex    string `json:"tx_index"`
	RPCAddress string `json:"rpc_address"`
}

// ID returns the node's peer ID.
func (info NodeInfo) ID() NodeID {
	_ = "STUB: not implemented"

	// Validate checks the self-reported NodeInfo is safe.
	// It returns an error if there
	// are too many Channels, if there are any duplicate Channels,
	// if the ListenAddr is malformed, or if the ListenAddr is a host name
	// that can not be resolved to some IP.
	// TODO: constraints for Moniker/Other? Or is that for the UI ?
	// JAE: It needs to be done on the client, but to prevent ambiguous
	// unicode characters, maybe it's worth sanitizing it here.
	// In the future we might want to validate these, once we have a
	// name-resolution system up.
	// International clients could then use punycode (or we could use
	// url-encoding), and we just need to be careful with how we handle that in our
	// clients. (e.g. off by default).
	return *new(NodeID)
}

func (info NodeInfo) Validate() error { _ = "STUB: not implemented"; return nil }

// Validate Version

// Validate Channels - ensure max and check for duplicates.

// Validate Other.

// XXX: Should we be more strict about address formats?

// CompatibleWith checks if two NodeInfo are compatible with each other.
// CONTRACT: two nodes are compatible if the Block version and network match
// and they have at least one channel in common.
func (info NodeInfo) CompatibleWith(other NodeInfo) error { _ = "STUB: not implemented"; return nil }

// nodes must be on the same network

// AddChannel is used by the router when a channel is opened to add it to the node info
func (info *NodeInfo) AddChannel(channel uint16) {
	_ = "STUB: not implemented"
	// check that the channel doesn't already exist
	return
}

func (info NodeInfo) ToProto() *tmp2p.NodeInfo { _ = "STUB: not implemented"; return nil }

func NodeInfoFromProto(pb *tmp2p.NodeInfo) (NodeInfo, error) {
	_ = "STUB: not implemented"
	return *new(NodeInfo), nil
}

// ResolveAddressString reads an address string, and returns the IP
// address and port information, returning an error for any validation
// errors.
func ResolveAddressString(addr string) (netip.AddrPort, error) {
	_ = "STUB: not implemented"
	return *new(netip.AddrPort), nil
}

// ResolveTCPAddr returns IPv6-embedded IPv4 addresses for no reason.

func removeProtocolIfDefined(addr string) string { _ = "STUB: not implemented"; return "" }

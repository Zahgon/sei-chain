package p2p

import (
	"context"
	"regexp"
	"strings"

	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

const defaultPort uint16 = 26657

var (
	// stringHasScheme tries to detect URLs with schemes. It looks for a : before a / (if any).
	stringHasScheme = func(str string) bool {
		return strings.Contains(str, "://")
	}

	// reSchemeIsHost tries to detect URLs where the scheme part is instead a
	// hostname, i.e. of the form "host:80/path" where host: is a hostname.
	reSchemeIsHost = regexp.MustCompile(`^[^/:]+:\d+(/|$)`)
)

// NodeAddress is a node address URL. It differs from a transport Endpoint in
// that it contains the node's ID, and that the hostname migth be either an IP or a DNS address.
type NodeAddress struct {
	NodeID   types.NodeID
	Hostname string
	Port     uint16
}

// ParseNodeAddress parses a node address URL into a NodeAddress, normalizing
// and validating it.
func ParseNodeAddress(urlString string) (NodeAddress, error) {
	_ = "STUB: not implemented"
	// url.Parse requires a scheme, so if it fails to parse a scheme-less URL
	// we try to apply a default scheme.
	return *new(NodeAddress), nil
}

// Otherwise, just parse a normal networked URL.

// For some reasons, missing or 0 port on parsing is interpretented as the default port.

// Resolve resolves a NodeAddress into a set of Endpoints, by expanding
// out a DNS hostname to IP addresses.
func (a NodeAddress) Resolve(ctx context.Context) ([]Endpoint, error) {
	_ = "STUB: not implemented"
	// LookIP for some reason returns IPv6-embedded addresses.
	return nil, nil
}

// String formats the address as a URL string.
func (a NodeAddress) String() string { _ = "STUB: not implemented"; return "" }

// Validate validates a NodeAddress.
func (a NodeAddress) Validate() error { _ = "STUB: not implemented"; return nil }

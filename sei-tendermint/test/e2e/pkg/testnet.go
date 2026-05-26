// nolint: gosec
package e2e

import (
	"math/rand"
	"net"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto"
	rpchttp "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/client/http"
)

const (
	randomSeed     int64  = 2308084734268
	proxyPortFirst uint32 = 5701
	networkIPv4           = "10.186.73.0/24"
	networkIPv6           = "fd80:b10c::/48"
)

type Mode string
type Protocol string
type Perturbation string

const (
	ModeValidator Mode = "validator"
	ModeFull      Mode = "full"
	ModeLight     Mode = "light"
	ModeSeed      Mode = "seed"

	ProtocolBuiltin Protocol = "builtin"
	ProtocolFile    Protocol = "file"
	ProtocolGRPC    Protocol = "grpc"
	ProtocolTCP     Protocol = "tcp"
	ProtocolUNIX    Protocol = "unix"

	PerturbationDisconnect Perturbation = "disconnect"
	PerturbationKill       Perturbation = "kill"
	PerturbationPause      Perturbation = "pause"
	PerturbationRestart    Perturbation = "restart"

	EvidenceAgeHeight int64         = 7
	EvidenceAgeTime   time.Duration = 500 * time.Millisecond

	StateSyncP2P      = "p2p"
	StateSyncRPC      = "rpc"
	StateSyncDisabled = ""
)

// Testnet represents a single testnet.
type Testnet struct {
	Name                   string
	File                   string
	Dir                    string
	IP                     *net.IPNet
	InitialHeight          int64
	InitialState           map[string]string
	Validators             map[*Node]int64
	ValidatorUpdates       map[int64]map[*Node]int64
	Nodes                  []*Node
	KeyType                string
	Evidence               int
	LogLevel               string
	TxSize                 int
	ABCIProtocol           Protocol
	PrepareProposalDelayMS int
	ProcessProposalDelayMS int
	CheckTxDelayMS         int
	FinalizeBlockDelayMS   int
}

// Node represents a Tendermint node in a testnet.
type Node struct {
	Name             string
	Testnet          *Testnet
	Mode             Mode
	PrivvalKey       crypto.PrivKey
	NodeKey          crypto.PrivKey
	IP               net.IP
	ProxyPort        uint32
	StartAt          int64
	Mempool          string
	StateSync        string
	Database         string
	PrivvalProtocol  Protocol
	PersistInterval  uint64
	SnapshotInterval uint64
	RetainBlocks     uint64
	Seeds            []*Node
	PersistentPeers  []*Node
	Perturbations    []Perturbation
	LogLevel         string
	QueueType        string
	HasStarted       bool
}

// LoadTestnet loads a testnet from a manifest file, using the filename to
// determine the testnet name and directory (from the basename of the file).
// The testnet generation must be deterministic, since it is generated
// separately by the runner and the test cases. For this reason, testnets use a
// random seed to generate e.g. keys.
func LoadTestnet(file string) (*Testnet, error) { _ = "STUB: not implemented"; return nil, nil }

// Set up resource generators. These must be deterministic.

// Set up nodes, in alphabetical order (IPs and ports get same order).

// normalize to 0 for initial nodes, since code expects this

// We do a second pass to set up seeds and persistent peers, which allows graph cycles.

// If there are no seeds or persistent peers specified, default to persistent
// connections to all other full nodes.

// Set up genesis validators. If not specified explicitly, use all validator nodes.

// Set up validator updates.

// Validate validates a testnet.
func (t Testnet) Validate() error { _ = "STUB: not implemented"; return nil }

// Validate validates a node.
func (n Node) Validate(testnet Testnet) error { _ = "STUB: not implemented"; return nil }

// LookupNode looks up a node by name. For now, simply do a linear search.
func (t Testnet) LookupNode(name string) *Node { _ = "STUB: not implemented"; return nil }

// ArchiveNodes returns a list of archive nodes that start at the initial height
// and contain the entire blockchain history. They are used e.g. as light client
// RPC servers.
func (t Testnet) ArchiveNodes() []*Node { _ = "STUB: not implemented"; return nil }

// IPv6 returns true if the testnet is an IPv6 network.
func (t Testnet) IPv6() bool { _ = "STUB: not implemented"; return false }

// HasPerturbations returns whether the network has any perturbations.
func (t Testnet) HasPerturbations() bool { _ = "STUB: not implemented"; return false }

// Address returns a P2P endpoint address for the node.
func (n Node) AddressP2P(withID bool) string { _ = "STUB: not implemented"; return "" }

// IPv6 addresses must be wrapped in [] to avoid conflict with : port separator

// Address returns an RPC endpoint address for the node.
func (n Node) AddressRPC() string { _ = "STUB: not implemented"; return "" }

// IPv6 addresses must be wrapped in [] to avoid conflict with : port separator

// Client returns an RPC client for a node.
func (n Node) Client() (*rpchttp.HTTP, error) { _ = "STUB: not implemented"; return nil, nil }

// Stateless returns true if the node is either a seed node or a light node
func (n Node) Stateless() bool { _ = "STUB: not implemented"; return false }

// keyGenerator generates pseudorandom Ed25519 keys based on a seed.
type keyGenerator struct {
	random *rand.Rand
}

func newKeyGenerator(seed int64) *keyGenerator { _ = "STUB: not implemented"; return nil }

func (g *keyGenerator) Generate(keyType string) crypto.PrivKey {
	_ = "STUB: not implemented"
	return *new(crypto.PrivKey)
}

// this shouldn't happen

// portGenerator generates local Docker proxy ports for each node.
type portGenerator struct {
	nextPort uint32
}

func newPortGenerator(firstPort uint32) *portGenerator { _ = "STUB: not implemented"; return nil }

func (g *portGenerator) Next() uint32 { _ = "STUB: not implemented"; return 0 }

// ipGenerator generates sequential IP addresses for each node, using a random
// network address.
type ipGenerator struct {
	network *net.IPNet
	nextIP  net.IP
}

func newIPGenerator(network *net.IPNet) *ipGenerator { _ = "STUB: not implemented"; return nil }

// Skip network and gateway addresses

func (g *ipGenerator) Network() *net.IPNet { _ = "STUB: not implemented"; return nil }

func (g *ipGenerator) Next() net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

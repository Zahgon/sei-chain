package p2p

import (
	"context"

	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p/conn"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/service"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils/tcp"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

// the maximum amount of addresses that can be included in a PEX batch.
const MaxPexAddrs = 100

type errBadNetwork struct{ error }

type PeerManager = peerManager[*ConnV2]
type PeerUpdatesRecv = peerUpdatesRecv[*ConnV2]
type ConnSet = connSet[*ConnV2]

// Router manages peer connections and routes messages between peers and channels.
type Router struct {
	*service.BaseService

	metrics *Metrics
	lc      *metricsLabelCache

	options     *RouterOptions
	privKey     NodeSecretKey
	peerManager *PeerManager

	peerDB           utils.Watch[*peerDB]
	nodeInfoProducer func() *types.NodeInfo

	channels utils.RWMutex[map[ChannelID]*channel]
	giga     utils.Option[*GigaRouter]

	started chan struct{}
}

func (r *Router) getChannelDescs() []*conn.ChannelDescriptor { _ = "STUB: not implemented"; return nil }

// NewRouter creates a new Router.
func NewRouter(
	metrics *Metrics,
	privKey NodeSecretKey,
	nodeInfoProducer func() *types.NodeInfo,
	db dbm.DB,
	options *RouterOptions,
) (*Router, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 100 is arbitrary - we need some bound, otherwise peerDB will
// maintain the whole connection history without pruning.
// 100 is more or less an upper bound on how many concurrent
// connections sei-v2 can effectively handle currently.

// initialAddrs will stay around util pex table fills the whole "extra" cache.

func (r *Router) Endpoint() Endpoint { _ = "STUB: not implemented"; return *new(Endpoint) }

func (r *Router) WaitForStart(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *Router) AddAddrs(sender types.NodeID, addrs []NodeAddress) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Router) Subscribe() *PeerUpdatesRecv { _ = "STUB: not implemented"; return nil }

func (r *Router) Connected(id types.NodeID) bool { _ = "STUB: not implemented"; return false }

func (r *Router) Advertise(maxAddrs int) []NodeAddress { _ = "STUB: not implemented"; return nil }

func (r *Router) ConnInfos() []PeerConnInfo { _ = "STUB: not implemented"; return nil }
func (r *Router) AllAddrs() []NodeAddress   { _ = "STUB: not implemented"; return nil }

// Giga returns the GigaRouter if Autobahn is enabled, None otherwise.
// Consumers (e.g. the /status RPC handler) use this to reach Autobahn-specific
// state like the last committed block number.
func (r *Router) Giga() utils.Option[*GigaRouter] {
	_ = "STUB: not implemented"

	// OpenChannel opens a new channel for the given message type.
	return nil
}

func OpenChannel[T gogoproto.Message](r *Router, chDesc ChannelDescriptor[T]) (*Channel[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// add the channel to the nodeInfo if it's not already there.

func (r *Router) acceptPeersRoutine(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// signal that we are listening

// Spawn a goroutine per connection.

// Listener has to send pex data, so that dialer can learn about more peers in
// case listener does not have capacity for new connections.
// Dialer also could potentially send pex data, but there is no benefit from doing so:
// - if listener is full, then it won't use the new data and it won't gossip it further either, since only verified data is gossiped.
// - if it is not full, then the connection will be established and pex data will be sent the regular way using PEX protocol.

func (r *Router) dialPeersRoutine(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Task feeding the upgrade permit to peer manager.

// Since the connection is not established yet, the handshake pex data
// will end up in a bounded cache, rather than main index. That's fine because
// we use the handshake pex data only for a local search,
// which is not supposed to be exhaustive.

// storePeersRoutine periodically snapshots the current connection set to disk,
// so that peers are immediately rediscovered on restart.
func (r *Router) storePeersRoutine(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Mark connections as still available.

func (r *Router) metricsRoutine(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Evict reports a peer misbehavior and forces peer to be disconnected.
func (r *Router) Evict(id types.NodeID, err error) { _ = "STUB: not implemented"; return }

func (r *Router) IsBlockSyncPeer(id types.NodeID) bool { _ = "STUB: not implemented"; return false }

// dial connects to a peer by dialing it.
func (r *Router) dial(ctx context.Context, addrs []NodeAddress) (_ tcp.Conn, err error) {
	_ = "STUB: not implemented"
	return *new(tcp.Conn), nil
}

// Resolve addresses in parallel. No errors expected,
// just resolve as many addresses as possible within timeout.

func (r *Router) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// OnStart implements service.Service.
func (r *Router) OnStart(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// OnStop implements service.Service.
func (r *Router) OnStop() { _ = "STUB: not implemented"; return }

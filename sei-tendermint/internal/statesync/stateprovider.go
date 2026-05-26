package statesync

import (
	"context"
	"sync"
	"time"

	"github.com/sei-protocol/seilog"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p"
	sm "github.com/sei-protocol/sei-chain/sei-tendermint/internal/state"
	"github.com/sei-protocol/sei-chain/sei-tendermint/light"
	lightprovider "github.com/sei-protocol/sei-chain/sei-tendermint/light/provider"
	pb "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/statesync"
	rpchttp "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/client/http"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

const (
	consensusParamsResponseTimeout = 5 * time.Second
)

//go:generate ../../scripts/mockery_generate.sh StateProvider

var logger = seilog.NewLogger("tendermint", "internal", "statesync")

// StateProvider is a provider of trusted state data for bootstrapping a node. This refers
// to the state.State object, not the state machine. There are two implementations. One
// uses the P2P layer and the other uses the RPC layer. Both use light client verification.
type StateProvider interface {
	// AppHash returns the app hash after the given height has been committed.
	AppHash(ctx context.Context, height uint64) ([]byte, error)
	// Commit returns the commit at the given height.
	Commit(ctx context.Context, height uint64) (*types.Commit, error)
	// State returns a state object at the given height.
	State(ctx context.Context, height uint64) (sm.State, error)
}

type stateProviderRPC struct {
	sync.Mutex              // light.Client is not concurrency-safe
	lc                      *light.Client
	initialHeight           int64
	providers               map[lightprovider.Provider]string
	verifyLightBlockTimeout time.Duration
}

// NewRPCStateProvider creates a new StateProvider using a light client and RPC clients.
func NewRPCStateProvider(
	ctx context.Context,
	chainID string,
	initialHeight int64,
	verifyLightBlockTimeout time.Duration,
	servers []string,
	trustOptions light.TrustOptions,
	blacklistTTL time.Duration,
) (StateProvider, error) {
	_ = "STUB: not implemented"
	return *new(StateProvider), nil
}

// We store the RPC addresses keyed by provider, so we can find the address of the primary
// provider used by the light client and use it to fetch consensus parameters.

func (s *stateProviderRPC) verifyLightBlockAtHeight(ctx context.Context, height uint64, ts time.Time) (*types.LightBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // height validated by Message.Validate() upstream

// AppHash implements part of StateProvider. It calls the application to verify the
// light blocks at heights h+1 and h+2 and, if verification succeeds, reports the app
// hash for the block at height h+1 which correlates to the state at height h.
func (s *stateProviderRPC) AppHash(ctx context.Context, height uint64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// We have to fetch the next height, which contains the app hash for the previous height.
		nil
}

// We also try to fetch the blocks at H+2, since we need these
// when building the state while restoring the snapshot. This avoids the race
// condition where we try to restore a snapshot before H+2 exists.

// Commit implements StateProvider.
func (s *stateProviderRPC) Commit(ctx context.Context, height uint64) (*types.Commit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// State implements StateProvider.
func (s *stateProviderRPC) State(ctx context.Context, height uint64) (sm.State, error) {
	_ = "STUB: not implemented"
	return *new(sm.State), nil
}

// The snapshot height maps onto the state heights as follows:
//
// height: last block, i.e. the snapshotted height
// height+1: current block, i.e. the first block we'll process after the snapshot
// height+2: next block, i.e. the second block after the snapshot
//
// We need to fetch the NextValidators from height+2 because if the application changed
// the validator set at the snapshot height then this only takes effect at height+2.

// We'll also need to fetch consensus params via RPC, using light client verification.

// rpcClient sets up a new RPC client
func rpcClient(server string) (*rpchttp.HTTP, error) { _ = "STUB: not implemented"; return nil, nil }

type StateProviderP2P struct {
	sync.Mutex              // light.Client is not concurrency-safe
	lc                      *light.Client
	initialHeight           int64
	paramsSendCh            *p2p.Channel[*pb.Message]
	paramsRecvCh            chan types.ConsensusParams
	verifyLightBlockTimeout time.Duration
}

// NewP2PStateProvider creates a light client state
// provider but uses a dispatcher connected to the P2P layer
func NewP2PStateProvider(
	ctx context.Context,
	chainID string,
	initialHeight int64,
	verifyLightBlockTimeout time.Duration,
	providers []lightprovider.Provider,
	trustOptions light.TrustOptions,
	paramsSendCh *p2p.Channel[*pb.Message],
	blacklistTTL time.Duration,
) (StateProvider, error) {
	_ = "STUB: not implemented"
	return *new(StateProvider), nil
}

func (s *StateProviderP2P) verifyLightBlockAtHeight(ctx context.Context, height uint64, ts time.Time) (*types.LightBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // height validated by Message.Validate() upstream

// AppHash implements StateProvider.
func (s *StateProviderP2P) AppHash(ctx context.Context, height uint64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// We have to fetch the next height, which contains the app hash for the previous height.
		nil
}

// We also try to fetch the blocks at H+2, since we need these
// when building the state while restoring the snapshot. This avoids the race
// condition where we try to restore a snapshot before H+2 exists.

// Commit implements StateProvider.
func (s *StateProviderP2P) Commit(ctx context.Context, height uint64) (*types.Commit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// State implements StateProvider.
func (s *StateProviderP2P) State(ctx context.Context, height uint64) (sm.State, error) {
	_ = "STUB: not implemented"
	return *new(sm.State), nil
}

// The snapshot height maps onto the state heights as follows:
//
// height: last block, i.e. the snapshotted height
// height+1: current block, i.e. the first block we'll process after the snapshot
// height+2: next block, i.e. the second block after the snapshot
//
// We need to fetch the NextValidators from height+2 because if the application changed
// the validator set at the snapshot height then this only takes effect at height+2.

// We'll also need to fetch consensus params via P2P.

// validate the consensus params

// set the last height changed to the current height

// AddProvider dynamically adds a peer as a new witness. A limit of 6 providers is kept as a
// heuristic. Too many overburdens the network and too little compromises the second layer of security.
func (s *StateProviderP2P) AddProvider(p lightprovider.Provider) { _ = "STUB: not implemented"; return }

// RemoveProviderByID removes a peer from the light client's witness list.
func (s *StateProviderP2P) RemoveProviderByID(ID types.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

// Providers returns the list of providers (useful for tests)
func (s *StateProviderP2P) Providers() []lightprovider.Provider {
	_ = "STUB: not implemented"
	return nil
}

func (s *StateProviderP2P) ParamsRecvCh() chan types.ConsensusParams {
	_ = "STUB: not implemented"
	return nil

	// consensusParams sends requests for consensus parameters to all witnesses
	// in parallel, retrying with increasing backoff until a response is
	// received or the context is canceled.
	//
	// For each witness, a goroutine sends a parameter request, retrying periodically
	// if no response is obtained, with increasing intervals. It returns the
	// consensus parameters upon receiving a response, or an error if the context is canceled.
}

func (s *StateProviderP2P) consensusParams(ctx context.Context, height int64) (types.ConsensusParams, error) {
	_ = "STUB: not implemented"
	return *new(types.ConsensusParams), nil
}

//nolint:gosec // height is a validated positive block height

// jitter+backoff the retry loop

//nolint:gosec

package evmrpc

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/types"
	"github.com/sei-protocol/sei-chain/x/evm/keeper"
)

const SleepInterval = 5 * time.Second
const NewHeadsListenerBuffer = 10

type SubscriptionAPI struct {
	tmClient            client.LocalClient
	subscriptionManager *SubscriptionManager
	subscriptonConfig   *SubscriptionConfig

	logFetcher          *LogFetcher
	newHeadListenersMtx *sync.RWMutex
	newHeadListeners    map[rpc.ID]chan map[string]interface{}
	connectionType      ConnectionType
}

type SubscriptionConfig struct {
	subscriptionCapacity int
	newHeadLimit         uint64
}

func NewSubscriptionAPI(tmClient client.LocalClient, k *keeper.Keeper, ctxProvider func(int64) sdk.Context, logFetcher *LogFetcher, subscriptionConfig *SubscriptionConfig, filterConfig *FilterConfig, connectionType ConnectionType, blockHeaderNotifier *BlockHeaderNotifier) *SubscriptionAPI {
	_ = "STUB: not implemented"
	return nil
}

// subscriptionManager is only constructed for the legacy
// event-bus path below; under Autobahn the notifier feeds the
// fan-out directly and the manager is unused.

// Autobahn (and any future direct-channel) path. The producer
// pushes one event per committed block; there is no Tendermint
// event-bus subscription.

// Legacy CometBFT path: subscribe to the Tendermint event bus.

func (a *SubscriptionAPI) runNewHeadsFromNotifier(notifier *BlockHeaderNotifier, k *keeper.Keeper, ctxProvider func(int64) sdk.Context) {
	_ = "STUB: not implemented"
	return
}

// Defend against a misbehaving producer. OnBlockCommitted's
// contract requires non-nil header/response, but a single bad
// event must not kill the fan-out goroutine for all subscribers.

// Source gasLimit from the active SDK ConsensusParams rather than
// evt.response.ConsensusParamUpdates: the latter is only populated
// on actual updates (nil for nearly every block). See block.go's
// GetBlockByNumber for the same pattern + rationale.

// pickHeadBaseFee returns the baseFeePerGas to attach to the eth_newHeads
// notification for the block at `height`. Mirrors block.go's
// GetBlockByNumber: GetNextBaseFeePerGas(ctx_at_N) is the fee for N+1, so
// we call it on the *parent* ctx (height-1). Genesis (height 1) has no
// parent; return the configured default min fee instead.
//
// `getNextBaseFee` is a function pointer rather than a *keeper.Keeper
// method so tests can inject a fake without needing a full keeper.
func pickHeadBaseFee(getNextBaseFee func(sdk.Context) sdk.Dec, ctxProvider func(int64) sdk.Context, height int64) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

func (a *SubscriptionAPI) broadcastNewHead(ethHeader map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}

func handleListener(c chan map[string]interface{}, ethHeader map[string]interface{}) bool {
	_ = "STUB: not implemented"
	// if the channel is already closed, sending to it/closing it will panic
	return false
}

// this path is hit when the buffer is full, meaning that the subscriber is not consuming
// fast enough

func (a *SubscriptionAPI) NewHeads(ctx context.Context) (s *rpc.Subscription, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// might have already been closed

func (a *SubscriptionAPI) Logs(ctx context.Context, filter *filters.FilterCriteria) (s *rpc.Subscription, _err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// create empty filter if filter does not exist

// when fromBlock is 0 and toBlock is latest, adjust the filter
// to unbounded filter

// set to latest block height
// set to nil to continue listening

// Track subscription metrics

const SubscriberPrefix = "evm.rpc."

type SubscriberID uint64

type SubInfo struct {
	Query          string
	SubscriptionCh <-chan coretypes.ResultEvent
}

type SubscriptionManager struct {
	subMu            sync.Mutex
	NextID           SubscriberID
	SubscriptionInfo map[SubscriberID]SubInfo
	tmClient         client.LocalClient
}

func NewSubscriptionManager(tmClient client.LocalClient) *SubscriptionManager {
	_ = "STUB: not implemented"
	return nil
}

func (s *SubscriptionManager) Subscribe(ctx context.Context, q *QueryBuilder, limit int) (SubscriberID, <-chan coretypes.ResultEvent, error) {
	_ = "STUB: not implemented"
	return *new(SubscriberID), nil, nil
}

// ignore deprecation here since the new endpoint does not support polling
//nolint:staticcheck

func (s *SubscriptionManager) Unsubscribe(ctx context.Context, id SubscriberID) error {
	_ = "STUB: not implemented"
	return nil
}

// ignore deprecation here since the new endpoint does not support polling
//nolint:staticcheck

// encodeCommittedBlock builds the eth_newHeads payload for an Autobahn-
// committed block. It differs from encodeTmHeader in two notable ways:
//
//  1. "hash" is the explicit Autobahn block-header hash from evt.hash
//     (the same value the EVM receipt store records as blockHash). See
//     blockHeaderEvent's doc for the rationale.
//  2. parentHash, receiptsRoot, and transactionsRoot are zero. The
//     Autobahn block-execution path does not compute a Tendermint-style
//     hash chain (LastBlockID / LastResultsHash / DataHash), so there is
//     nothing meaningful to surface for those fields. Subscribers that
//     chain-validate the head stream will need a different mechanism
//     under Autobahn.
//
// stateRoot is taken from evt.response.AppHash (the AppHash produced by
// finalizing *this* block). evt.header.AppHash would be wrong: by
// Tendermint convention Header.AppHash holds the result of the previous
// block, not the current one, so the producer leaves it unset.
//
// gasLimit is read by the caller from the active SDK ConsensusParams
// (see runNewHeadsFromNotifier); ConsensusParamUpdates on the response
// would be nil for the vast majority of blocks.
func encodeCommittedBlock(evt blockHeaderEvent, baseFee *big.Int, gasLimit int64) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// TODO(autobahn): TxResult.GasUsed can be wrong for ante-failing EVM
// txs; block.go (GetBlockByNumber) sums receipt.GasUsed for that
// reason. We approximate here to keep newHeads cheap; subscribers
// needing exact gas should call eth_getBlockByNumber.

// inapplicable to Sei
// inapplicable to Sei
//nolint:gosec
//nolint:gosec
// TODO(autobahn): derive from receipts so newHeads subscribers can pre-filter logs

// inapplicable to Sei

// see function doc
// see function doc
// inapplicable to Sei

//nolint:gosec
// see function doc
// inapplicable to Sei
// inapplicable to Sei
// inapplicable to Sei

// inapplicable to Sei
// inapplicable to Sei

func encodeTmHeader(
	header tmtypes.EventDataNewBlockHeader,
	baseFee *big.Int,
) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec

// inapplicable to Sei
// inapplicable to Sei

//nolint:gosec
// inapplicable to Sei

// inapplicable to Sei

// inapplicable to Sei

//nolint:gosec

// inapplicable to Sei
// inapplicable to Sei
// inapplicable to Sei

// inapplicable to Sei

// inapplicable to Sei
// inapplicable to Sei
